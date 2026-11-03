#!/usr/bin/env bash
# build-qt-static.sh — build a static qtbase for mingw-w64, the Qt that build-dll.sh links into the
# shim DLL.
#
#   QTBASE_SRC=<qtbase-everywhere-src-6.11.1> QT_STATIC=<install folder> ./qt-static/build-qt-static.sh
#
# Needs cmake, ninja, the mingw-w64 cross compilers, and a host Qt 6.11.1 (its moc/rcc tools;
# QT_HOST_PATH, default /usr). The version must match the bindings (6.11.1). Settings: QTBASE_SRC
# and QT_STATIC (required), QT_HOST_PATH, BUILD (build folder, default <QT_STATIC>-build), JOBS.
set -euo pipefail
here=$(cd "$(dirname "$0")" && pwd)

src=${QTBASE_SRC:?set QTBASE_SRC to the qtbase 6.11.1 source folder (qtbase-everywhere-src-6.11.1)}
prefix=${QT_STATIC:?set QT_STATIC to the folder to install the static Qt into}
build=${BUILD:-$prefix-build}
jobs=${JOBS:-2}

cmake -S "$src" -B "$build" -G Ninja \
	-DCMAKE_TOOLCHAIN_FILE="$here/mingw-w64-x86_64.cmake" \
	-DCMAKE_BUILD_TYPE=Release \
	-DBUILD_SHARED_LIBS=OFF \
	-DQT_HOST_PATH="${QT_HOST_PATH:-/usr}" \
	-DQT_BUILD_EXAMPLES=OFF \
	-DQT_BUILD_TESTS=OFF \
	-DCMAKE_INSTALL_PREFIX="$prefix"
cmake --build "$build" --parallel "$jobs"
cmake --install "$build"
echo ">> static Qt installed in $prefix"
