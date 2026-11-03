package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStandardPaths_writableLocation uintptr
	addr_QStandardPaths_standardLocations uintptr
	addr_QStandardPaths_locate uintptr
	addr_QStandardPaths_locateAll uintptr
	addr_QStandardPaths_displayName uintptr
	addr_QStandardPaths_findExecutable uintptr
	addr_QStandardPaths_setTestModeEnabled uintptr
	addr_QStandardPaths_isTestModeEnabled uintptr
	addr_QStandardPaths_locate2 uintptr
	addr_QStandardPaths_locateAll2 uintptr
	addr_QStandardPaths_findExecutable2 uintptr
)

func init() {
	symaddr(&addr_QStandardPaths_writableLocation, "QStandardPaths_writableLocation")
	symaddr(&addr_QStandardPaths_standardLocations, "QStandardPaths_standardLocations")
	symaddr(&addr_QStandardPaths_locate, "QStandardPaths_locate")
	symaddr(&addr_QStandardPaths_locateAll, "QStandardPaths_locateAll")
	symaddr(&addr_QStandardPaths_displayName, "QStandardPaths_displayName")
	symaddr(&addr_QStandardPaths_findExecutable, "QStandardPaths_findExecutable")
	symaddr(&addr_QStandardPaths_setTestModeEnabled, "QStandardPaths_setTestModeEnabled")
	symaddr(&addr_QStandardPaths_isTestModeEnabled, "QStandardPaths_isTestModeEnabled")
	symaddr(&addr_QStandardPaths_locate2, "QStandardPaths_locate2")
	symaddr(&addr_QStandardPaths_locateAll2, "QStandardPaths_locateAll2")
	symaddr(&addr_QStandardPaths_findExecutable2, "QStandardPaths_findExecutable2")
}

type QStandardPaths struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStandardPaths) UnsafeGoID() int64     { return this.vid }
func (this *QStandardPaths) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStandardPaths(h uintptr) *QStandardPaths {
	if h == 0 {
		return nil
	}
	return &QStandardPaths{h: h}
}

func newownedQStandardPaths(h uintptr) *QStandardPaths { return bindQStandardPaths(h) }

func (this *QStandardPaths) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStandardPaths) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStandardPaths(h unsafe.Pointer) *QStandardPaths { return bindQStandardPaths(uintptr(h)) }
func UnsafeNewQStandardPathsOwned(h unsafe.Pointer) *QStandardPaths { return newownedQStandardPaths(uintptr(h)) }

func (this *QStandardPaths) UnsafeRelease() { this.h = 0 }

func QStandardPaths_WritableLocation(type_ QStandardPaths__StandardLocation) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_writableLocation, uintptr(type_), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QStandardPaths_StandardLocations(type_ QStandardPaths__StandardLocation) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_standardLocations, uintptr(type_), uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStandardPaths_Locate(type_ QStandardPaths__StandardLocation, fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_locate, uintptr(type_), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func QStandardPaths_LocateAll(type_ QStandardPaths__StandardLocation, fileName string) []string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_locateAll, uintptr(type_), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStandardPaths_DisplayName(type_ QStandardPaths__StandardLocation) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_displayName, uintptr(type_), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QStandardPaths_FindExecutable(executableName string) string {
	var executableName_p uintptr
	if len(executableName) > 0 {
		executableName_p = uintptr(unsafe.Pointer(unsafe.StringData(executableName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_findExecutable, executableName_p, uintptr(len(executableName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(executableName)
	return pgStr(__r, __n)
}

func QStandardPaths_SetTestModeEnabled(testMode bool) {
	purego.SyscallN(addr_QStandardPaths_setTestModeEnabled, uintptr(b2i(testMode)))
}

func QStandardPaths_IsTestModeEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardPaths_isTestModeEnabled)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QStandardPaths_Locate2(type_ QStandardPaths__StandardLocation, fileName string, options QStandardPaths__LocateOption) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_locate2, uintptr(type_), fileName_p, uintptr(len(fileName)), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func QStandardPaths_LocateAll2(type_ QStandardPaths__StandardLocation, fileName string, options QStandardPaths__LocateOption) []string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_locateAll2, uintptr(type_), fileName_p, uintptr(len(fileName)), uintptr(options), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStandardPaths_FindExecutable2(executableName string, paths []string) string {
	var executableName_p uintptr
	if len(executableName) > 0 {
		executableName_p = uintptr(unsafe.Pointer(unsafe.StringData(executableName)))
	}
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardPaths_findExecutable2, executableName_p, uintptr(len(executableName)), paths_p, uintptr(len(paths)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(executableName)
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
	return pgStr(__r, __n)
}

func (this *QStandardPaths) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
