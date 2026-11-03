#!/usr/bin/env bash
# regenerate.sh — regenerate qt/ (the Go bindings and the C++ shims) with cppgo from the Windows Qt
# headers: cppgo's Windows pipeline, stages 1-3 (frontend, IR, purego/LLP64 generation).
#
#   CPPGO=<cppgo checkout> QT_STATIC=<static Qt 6.11.1 mingw install> ./regenerate.sh
#
# Needs clang/LLVM with dev headers (cppgo's frontend), Go, and the static Qt from qt-static/ (its
# headers are the Windows headers). The shims include Qt headers relative to <QT_STATIC>/include
# (CPPGO_INCLUDE_ROOTS), so they compile wherever build-dll.sh is pointed at a static Qt.
set -euo pipefail
here=$(cd "$(dirname "$0")" && pwd)

cppgo=${CPPGO:?set CPPGO to a cppgo checkout (https://github.com/5k3105/cppgo)}
QT=${QT_STATIC:?set QT_STATIC to the static Qt install (the folder with include/)}
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
cflags="-I$QT/include -I$QT/include/QtCore -I$QT/include/QtGui -I$QT/include/QtWidgets -I$QT/mkspecs/win32-g++ -DQT_CORE_LIB -DQT_GUI_LIB -DQT_WIDGETS_LIB"

cd "$cppgo"
echo ">> [1/3] cppgo frontend"
LLVM_INC="$(llvm-config --includedir 2>/dev/null || true)"
LLVM_LIB="$(llvm-config --libdir 2>/dev/null || true)"
clang++ -std=c++17 -fno-rtti ${LLVM_INC:+-I"$LLVM_INC"} ${LLVM_LIB:+-L"$LLVM_LIB"} \
	frontend/cppir2.cpp -o "$tmp/cppir2" -lclang-cpp -lLLVM

echo ">> [2/3] IR from the Windows Qt headers"
echo "#include <QtWidgets/QtWidgets>" > "$tmp/umbrella_qt.h"
"$tmp/cppir2" "$tmp/umbrella_qt.h" "QtCore,QtGui,QtWidgets" --target=x86_64-w64-mingw32 $cflags > "$tmp/qtbase-win.json"

echo ">> [3/3] generate (purego backend, LLP64)"
go build -o "$tmp/cppgogen" ./cmd/gen
rm -rf "$here/qt"
CPPGO_BACKEND=purego CPPGO_DATAMODEL=llp64 CPPGO_INCLUDE_ROOTS="$QT/include" \
	"$tmp/cppgogen" "$here/qt" qt - all "$tmp/qtbase-win.json"

# the two files cppgo does not generate: the DLL's embed/self-extract and the static plugin imports
cp "$here/extra/aaa_bootstrap_windows.go" "$here/qt/"
cp "$here/extra/windows_plugins.cpp" "$here/qt/cshim/"
echo ">> regenerated qt/ — now build the shim DLL: QT_STATIC=$QT ./build-dll.sh"
