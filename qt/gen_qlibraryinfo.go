package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLibraryInfo_build uintptr
	addr_QLibraryInfo_isDebugBuild uintptr
	addr_QLibraryInfo_isSharedBuild uintptr
	addr_QLibraryInfo_version uintptr
	addr_QLibraryInfo_path uintptr
	addr_QLibraryInfo_paths uintptr
	addr_QLibraryInfo_location uintptr
	addr_QLibraryInfo_platformPluginArguments uintptr
	pg_QLibraryInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLibraryInfo_build, "QLibraryInfo_build")
	symaddr(&addr_QLibraryInfo_isDebugBuild, "QLibraryInfo_isDebugBuild")
	symaddr(&addr_QLibraryInfo_isSharedBuild, "QLibraryInfo_isSharedBuild")
	symaddr(&addr_QLibraryInfo_version, "QLibraryInfo_version")
	symaddr(&addr_QLibraryInfo_path, "QLibraryInfo_path")
	symaddr(&addr_QLibraryInfo_paths, "QLibraryInfo_paths")
	symaddr(&addr_QLibraryInfo_location, "QLibraryInfo_location")
	symaddr(&addr_QLibraryInfo_platformPluginArguments, "QLibraryInfo_platformPluginArguments")
	reg(&pg_QLibraryInfo_delete, "QLibraryInfo_delete")
}

type QLibraryInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLibraryInfo) UnsafeGoID() int64     { return this.vid }
func (this *QLibraryInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLibraryInfo(h uintptr) *QLibraryInfo {
	if h == 0 {
		return nil
	}
	return &QLibraryInfo{h: h}
}

func newownedQLibraryInfo(h uintptr) *QLibraryInfo {
	o := bindQLibraryInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLibraryInfo).Delete)
	}
	return o
}

func (this *QLibraryInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLibraryInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLibraryInfo(h unsafe.Pointer) *QLibraryInfo { return bindQLibraryInfo(uintptr(h)) }
func UnsafeNewQLibraryInfoOwned(h unsafe.Pointer) *QLibraryInfo { return newownedQLibraryInfo(uintptr(h)) }

func (this *QLibraryInfo) UnsafeRelease() { this.h = 0 }

func QLibraryInfo_Build() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_build, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLibraryInfo_IsDebugBuild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLibraryInfo_isDebugBuild)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QLibraryInfo_IsSharedBuild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLibraryInfo_isSharedBuild)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QLibraryInfo_Version() *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_version)
	return newownedQVersionNumber(__r)
}

func QLibraryInfo_Path(p QLibraryInfo__LibraryPath) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_path, uintptr(p), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLibraryInfo_Paths(p QLibraryInfo__LibraryPath) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_paths, uintptr(p), uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QLibraryInfo_Location(location QLibraryInfo__LibraryPath) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_location, uintptr(location), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QLibraryInfo_PlatformPluginArguments(platformName string) []string {
	var platformName_p uintptr
	if len(platformName) > 0 {
		platformName_p = uintptr(unsafe.Pointer(unsafe.StringData(platformName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibraryInfo_platformPluginArguments, platformName_p, uintptr(len(platformName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(platformName)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QLibraryInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLibraryInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
