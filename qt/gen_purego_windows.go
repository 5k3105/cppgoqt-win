//go:build windows

package qt

import "syscall"

// shimLibName is the default shim DLL (searched via the standard DLL search order;
// $CPPGO_LIB_QT overrides with an explicit path). DLLs are per-module by construction —
// no RTLD_GLOBAL interposition risk on Windows.
const shimLibName = "shim_qt.dll"

func openShimLib(path string) (uintptr, error) {
	h, err := syscall.LoadLibrary(path)
	return uintptr(h), err
}

func shimSym(name string) (uintptr, error) {
	return syscall.GetProcAddress(syscall.Handle(_lib), name)
}
