#!/usr/bin/env bash
# build-dll.sh — compile the generated C++ shims (qt/cshim) and link them against a static Qt into
# qt/shim_qt.dll, the DLL the Go package embeds (qt/aaa_bootstrap_windows.go).
#
#   QT_STATIC=<static Qt 6.11.1 mingw install> ./build-dll.sh
#
# Needs the mingw-w64 cross compiler (x86_64-w64-mingw32-g++) and a static Qt built for mingw with
# the same Qt version the bindings were generated against (qt-static/ shows how). The DLL then
# depends only on Windows system DLLs: Qt, its platform and style plugins, and the mingw runtimes
# are all linked in.
#
# Settings: QT_STATIC (required), CXX (default x86_64-w64-mingw32-g++), JOBS (parallel compiles,
# default 2: each unit is large).
set -euo pipefail
cd "$(dirname "$0")"

QT=${QT_STATIC:?set QT_STATIC to the static Qt install (the folder with include/, lib/, plugins/)}
CXX=${CXX:-x86_64-w64-mingw32-g++}
JOBS=${JOBS:-2}
command -v "$CXX" >/dev/null || { echo "build-dll.sh: $CXX not found (install mingw-w64's g++)" >&2; exit 1; }
for f in lib/libQt6Widgets.a plugins/platforms/libqwindows.a plugins/styles/libqmodernwindowsstyle.a; do
	[ -f "$QT/$f" ] || { echo "build-dll.sh: $QT/$f is missing: QT_STATIC is not a static mingw Qt (see qt-static/)" >&2; exit 1; }
done

cflags=(-std=c++17 -O0 -g0 -I"$QT/include" -I"$QT/include/QtCore" -I"$QT/include/QtGui" -I"$QT/include/QtWidgets"
	-I"$QT/mkspecs/win32-g++" -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB -Iqt -include qt/cppgo_pch.h)

obj=build/obj
mkdir -p "$obj"
start=$(date +%s)
echo ">> compiling $(ls qt/cshim/*.cpp | wc -l) shim units ($JOBS at a time)"
export CXX obj
printf '%s\n' qt/cshim/*.cpp | xargs -P "$JOBS" -I{} sh -c '
	src=$1; shift
	o="$obj/$(basename "$src").o"
	[ "$o" -nt "$src" ] && exit 0  # up to date
	"$CXX" "$@" -c "$src" -o "$o" || { echo "build-dll.sh: compiling $src failed" >&2; exit 255; }
' _ {} "${cflags[@]}"

# the resource objects the platform and style plugins need (rcc init objects), not every module's
rcc=$(ls "$QT"/lib/objects-Release/*/.qt/rcc/*_init.cpp.obj | grep -E '/(QWindowsIntegrationPlugin|Gui|Widgets)_resources_')

echo ">> linking qt/shim_qt.dll"
# --exclude-libs,ALL: export only the shim's extern "C" functions (re-exporting every static Qt
# symbol overflows the DLL export table). -static: bake in the mingw runtimes and winpthread.
"$CXX" -shared -static -static-libgcc -static-libstdc++ -Wl,--exclude-libs,ALL -o qt/shim_qt.dll \
	"$obj"/*.o $rcc \
	-Wl,--start-group \
	"$QT"/plugins/styles/libqmodernwindowsstyle.a \
	"$QT"/plugins/platforms/libqwindows.a \
	"$QT"/lib/libQt6Widgets.a \
	"$QT"/lib/libQt6OpenGL.a \
	"$QT"/lib/libQt6Gui.a \
	"$QT"/lib/libQt6Core.a \
	"$QT"/lib/libQt6BundledHarfbuzz.a \
	"$QT"/lib/libQt6BundledFreetype.a \
	"$QT"/lib/libQt6BundledLibpng.a \
	"$QT"/lib/libQt6BundledPcre2.a \
	"$QT"/lib/libQt6BundledZLIB.a \
	-Wl,--end-group \
	-ldwmapi -limm32 -loleaut32 -lsetupapi -lshlwapi -lwinspool -lwtsapi32 -lshcore -lcomdlg32 \
	-ld3d9 -ld3d11 -ldxgi -ldxguid -ld3d12 -lgdi32 -luxtheme -ld2d1 -ldwrite \
	-lsynchronization -lmpr -luserenv -ladvapi32 -lauthz -lkernel32 -lnetapi32 -lntdll \
	-lole32 -lruntimeobject -lshell32 -luser32 -lversion -lwinmm -lws2_32 -luuid

echo ">> built qt/shim_qt.dll ($(du -h qt/shim_qt.dll | cut -f1)) in $(( $(date +%s) - start ))s"
if command -v x86_64-w64-mingw32-objdump >/dev/null; then
	echo ">> it needs: $(x86_64-w64-mingw32-objdump -p qt/shim_qt.dll | awk '/DLL Name/ {printf "%s ", $3}')"
fi
