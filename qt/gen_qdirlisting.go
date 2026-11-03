package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDirListing_new uintptr
	addr_QDirListing_new2 uintptr
	addr_QDirListing_new3 uintptr
	addr_QDirListing_new4 uintptr
	addr_QDirListing_swap uintptr
	addr_QDirListing_iteratorPath uintptr
	addr_QDirListing_iteratorFlags uintptr
	addr_QDirListing_nameFilters uintptr
	addr_QDirListing_end uintptr
	addr_QDirListing_cend uintptr
	addr_QDirListing_constEnd uintptr
	pg_QDirListing_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDirListing_new, "QDirListing_new")
	symaddr(&addr_QDirListing_new2, "QDirListing_new2")
	symaddr(&addr_QDirListing_new3, "QDirListing_new3")
	symaddr(&addr_QDirListing_new4, "QDirListing_new4")
	symaddr(&addr_QDirListing_swap, "QDirListing_swap")
	symaddr(&addr_QDirListing_iteratorPath, "QDirListing_iteratorPath")
	symaddr(&addr_QDirListing_iteratorFlags, "QDirListing_iteratorFlags")
	symaddr(&addr_QDirListing_nameFilters, "QDirListing_nameFilters")
	symaddr(&addr_QDirListing_end, "QDirListing_end")
	symaddr(&addr_QDirListing_cend, "QDirListing_cend")
	symaddr(&addr_QDirListing_constEnd, "QDirListing_constEnd")
	reg(&pg_QDirListing_delete, "QDirListing_delete")
}

type QDirListing struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDirListing) UnsafeGoID() int64     { return this.vid }
func (this *QDirListing) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDirListing(h uintptr) *QDirListing {
	if h == 0 {
		return nil
	}
	return &QDirListing{h: h}
}

func newownedQDirListing(h uintptr) *QDirListing {
	o := bindQDirListing(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDirListing).Delete)
	}
	return o
}

func (this *QDirListing) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDirListing) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDirListing(h unsafe.Pointer) *QDirListing { return bindQDirListing(uintptr(h)) }
func UnsafeNewQDirListingOwned(h unsafe.Pointer) *QDirListing { return newownedQDirListing(uintptr(h)) }

func (this *QDirListing) UnsafeRelease() { this.h = 0 }

func NewQDirListing(path string) *QDirListing {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirListing_new, path_p, uintptr(len(path)))
	__o := newownedQDirListing(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirListing2(path string, nameFilters []string) *QDirListing {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QDirListing_new2, path_p, uintptr(len(path)), nameFilters_p, uintptr(len(nameFilters)))
	__o := newownedQDirListing(__r)
	runtime.KeepAlive(path)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	return __o
}

func NewQDirListing3(path string, flags QDirListing__IteratorFlag) *QDirListing {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirListing_new3, path_p, uintptr(len(path)), uintptr(flags))
	__o := newownedQDirListing(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirListing4(path string, nameFilters []string, flags QDirListing__IteratorFlag) *QDirListing {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QDirListing_new4, path_p, uintptr(len(path)), nameFilters_p, uintptr(len(nameFilters)), uintptr(flags))
	__o := newownedQDirListing(__r)
	runtime.KeepAlive(path)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	return __o
}

func (this *QDirListing) Swap(other *QDirListing) {
	purego.SyscallN(addr_QDirListing_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QDirListing) IteratorPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing_iteratorPath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing) IteratorFlags() QDirListing__IteratorFlag {
	__r0, _, _ := purego.SyscallN(addr_QDirListing_iteratorFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDirListing__IteratorFlag(__r)
}

func (this *QDirListing) NameFilters() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing_nameFilters, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDirListing) End() *QDirListing__sentinel {
	__r, _, _ := purego.SyscallN(addr_QDirListing_end, this.h)
	runtime.KeepAlive(this)
	return newownedQDirListing__sentinel(__r)
}

func (this *QDirListing) Cend() *QDirListing__sentinel {
	__r, _, _ := purego.SyscallN(addr_QDirListing_cend, this.h)
	runtime.KeepAlive(this)
	return newownedQDirListing__sentinel(__r)
}

func (this *QDirListing) ConstEnd() *QDirListing__sentinel {
	__r, _, _ := purego.SyscallN(addr_QDirListing_constEnd, this.h)
	runtime.KeepAlive(this)
	return newownedQDirListing__sentinel(__r)
}

func (this *QDirListing) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDirListing_delete(this.h)
	}
	this.UnsafeRelease()
}
