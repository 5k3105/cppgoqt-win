package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPageRanges_new uintptr
	addr_QPageRanges_new2 uintptr
	addr_QPageRanges_OperatorAssign uintptr
	addr_QPageRanges_swap uintptr
	addr_QPageRanges_addPage uintptr
	addr_QPageRanges_addRange uintptr
	addr_QPageRanges_toRangeList uintptr
	addr_QPageRanges_clear uintptr
	addr_QPageRanges_toString uintptr
	addr_QPageRanges_fromString uintptr
	addr_QPageRanges_contains uintptr
	addr_QPageRanges_isEmpty uintptr
	addr_QPageRanges_firstPage uintptr
	addr_QPageRanges_lastPage uintptr
	addr_QPageRanges_detach uintptr
	addr_QPageRanges_OperatorEqual uintptr
	addr_QPageRanges_OperatorNotEqual uintptr
	pg_QPageRanges_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPageRanges_new, "QPageRanges_new")
	symaddr(&addr_QPageRanges_new2, "QPageRanges_new2")
	symaddr(&addr_QPageRanges_OperatorAssign, "QPageRanges_OperatorAssign")
	symaddr(&addr_QPageRanges_swap, "QPageRanges_swap")
	symaddr(&addr_QPageRanges_addPage, "QPageRanges_addPage")
	symaddr(&addr_QPageRanges_addRange, "QPageRanges_addRange")
	symaddr(&addr_QPageRanges_toRangeList, "QPageRanges_toRangeList")
	symaddr(&addr_QPageRanges_clear, "QPageRanges_clear")
	symaddr(&addr_QPageRanges_toString, "QPageRanges_toString")
	symaddr(&addr_QPageRanges_fromString, "QPageRanges_fromString")
	symaddr(&addr_QPageRanges_contains, "QPageRanges_contains")
	symaddr(&addr_QPageRanges_isEmpty, "QPageRanges_isEmpty")
	symaddr(&addr_QPageRanges_firstPage, "QPageRanges_firstPage")
	symaddr(&addr_QPageRanges_lastPage, "QPageRanges_lastPage")
	symaddr(&addr_QPageRanges_detach, "QPageRanges_detach")
	symaddr(&addr_QPageRanges_OperatorEqual, "QPageRanges_OperatorEqual")
	symaddr(&addr_QPageRanges_OperatorNotEqual, "QPageRanges_OperatorNotEqual")
	reg(&pg_QPageRanges_delete, "QPageRanges_delete")
}

type QPageRanges struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPageRanges) UnsafeGoID() int64     { return this.vid }
func (this *QPageRanges) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPageRanges(h uintptr) *QPageRanges {
	if h == 0 {
		return nil
	}
	return &QPageRanges{h: h}
}

func newownedQPageRanges(h uintptr) *QPageRanges {
	o := bindQPageRanges(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPageRanges).Delete)
	}
	return o
}

func (this *QPageRanges) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPageRanges) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPageRanges(h unsafe.Pointer) *QPageRanges { return bindQPageRanges(uintptr(h)) }
func UnsafeNewQPageRangesOwned(h unsafe.Pointer) *QPageRanges { return newownedQPageRanges(uintptr(h)) }

func (this *QPageRanges) UnsafeRelease() { this.h = 0 }

func NewQPageRanges() *QPageRanges {
	__r, _, _ := purego.SyscallN(addr_QPageRanges_new)
	__o := newownedQPageRanges(__r)
	return __o
}

func NewQPageRanges2(other *QPageRanges) *QPageRanges {
	__r, _, _ := purego.SyscallN(addr_QPageRanges_new2, uintptr(other.UnsafePointer()))
	__o := newownedQPageRanges(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QPageRanges) OperatorAssign(other *QPageRanges) *QPageRanges {
	__r, _, _ := purego.SyscallN(addr_QPageRanges_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPageRanges(__r)
}

func (this *QPageRanges) Swap(other *QPageRanges) {
	purego.SyscallN(addr_QPageRanges_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPageRanges) AddPage(pageNumber int) {
	purego.SyscallN(addr_QPageRanges_addPage, this.h, uintptr(pageNumber))
	runtime.KeepAlive(this)
}

func (this *QPageRanges) AddRange(from int, to int) {
	purego.SyscallN(addr_QPageRanges_addRange, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QPageRanges) ToRangeList() []*QPageRanges__Range {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageRanges_toRangeList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPageRanges__Range, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPageRanges__Range(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPageRanges) Clear() {
	purego.SyscallN(addr_QPageRanges_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QPageRanges) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageRanges_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QPageRanges_FromString(ranges string) *QPageRanges {
	var ranges_p uintptr
	if len(ranges) > 0 {
		ranges_p = uintptr(unsafe.Pointer(unsafe.StringData(ranges)))
	}
	__r, _, _ := purego.SyscallN(addr_QPageRanges_fromString, ranges_p, uintptr(len(ranges)))
	runtime.KeepAlive(ranges)
	return newownedQPageRanges(__r)
}

func (this *QPageRanges) Contains(pageNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_contains, this.h, uintptr(pageNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageRanges) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageRanges) FirstPage() int {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_firstPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPageRanges) LastPage() int {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_lastPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPageRanges) Detach() {
	purego.SyscallN(addr_QPageRanges_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QPageRanges) OperatorEqual(rhs *QPageRanges) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageRanges) OperatorNotEqual(rhs *QPageRanges) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageRanges_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageRanges) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPageRanges_delete(this.h)
	}
	this.UnsafeRelease()
}
