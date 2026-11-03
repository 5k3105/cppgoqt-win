package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QItemSelectionRange_new uintptr
	addr_QItemSelectionRange_new2 uintptr
	addr_QItemSelectionRange_new3 uintptr
	addr_QItemSelectionRange_new4 uintptr
	addr_QItemSelectionRange_swap uintptr
	addr_QItemSelectionRange_top uintptr
	addr_QItemSelectionRange_left uintptr
	addr_QItemSelectionRange_bottom uintptr
	addr_QItemSelectionRange_right uintptr
	addr_QItemSelectionRange_width uintptr
	addr_QItemSelectionRange_height uintptr
	addr_QItemSelectionRange_topLeft uintptr
	addr_QItemSelectionRange_bottomRight uintptr
	addr_QItemSelectionRange_parent uintptr
	addr_QItemSelectionRange_model uintptr
	addr_QItemSelectionRange_contains uintptr
	addr_QItemSelectionRange_contains2 uintptr
	addr_QItemSelectionRange_intersects uintptr
	addr_QItemSelectionRange_intersected uintptr
	addr_QItemSelectionRange_isValid uintptr
	addr_QItemSelectionRange_isEmpty uintptr
	addr_QItemSelectionRange_indexes uintptr
	addr_QItemSelectionRange_OperatorAssign uintptr
	addr_QItemSelectionRange_OperatorEqual uintptr
	addr_QItemSelectionRange_OperatorNotEqual uintptr
	pg_QItemSelectionRange_delete func(uintptr)
)

func init() {
	symaddr(&addr_QItemSelectionRange_new, "QItemSelectionRange_new")
	symaddr(&addr_QItemSelectionRange_new2, "QItemSelectionRange_new2")
	symaddr(&addr_QItemSelectionRange_new3, "QItemSelectionRange_new3")
	symaddr(&addr_QItemSelectionRange_new4, "QItemSelectionRange_new4")
	symaddr(&addr_QItemSelectionRange_swap, "QItemSelectionRange_swap")
	symaddr(&addr_QItemSelectionRange_top, "QItemSelectionRange_top")
	symaddr(&addr_QItemSelectionRange_left, "QItemSelectionRange_left")
	symaddr(&addr_QItemSelectionRange_bottom, "QItemSelectionRange_bottom")
	symaddr(&addr_QItemSelectionRange_right, "QItemSelectionRange_right")
	symaddr(&addr_QItemSelectionRange_width, "QItemSelectionRange_width")
	symaddr(&addr_QItemSelectionRange_height, "QItemSelectionRange_height")
	symaddr(&addr_QItemSelectionRange_topLeft, "QItemSelectionRange_topLeft")
	symaddr(&addr_QItemSelectionRange_bottomRight, "QItemSelectionRange_bottomRight")
	symaddr(&addr_QItemSelectionRange_parent, "QItemSelectionRange_parent")
	symaddr(&addr_QItemSelectionRange_model, "QItemSelectionRange_model")
	symaddr(&addr_QItemSelectionRange_contains, "QItemSelectionRange_contains")
	symaddr(&addr_QItemSelectionRange_contains2, "QItemSelectionRange_contains2")
	symaddr(&addr_QItemSelectionRange_intersects, "QItemSelectionRange_intersects")
	symaddr(&addr_QItemSelectionRange_intersected, "QItemSelectionRange_intersected")
	symaddr(&addr_QItemSelectionRange_isValid, "QItemSelectionRange_isValid")
	symaddr(&addr_QItemSelectionRange_isEmpty, "QItemSelectionRange_isEmpty")
	symaddr(&addr_QItemSelectionRange_indexes, "QItemSelectionRange_indexes")
	symaddr(&addr_QItemSelectionRange_OperatorAssign, "QItemSelectionRange_OperatorAssign")
	symaddr(&addr_QItemSelectionRange_OperatorEqual, "QItemSelectionRange_OperatorEqual")
	symaddr(&addr_QItemSelectionRange_OperatorNotEqual, "QItemSelectionRange_OperatorNotEqual")
	reg(&pg_QItemSelectionRange_delete, "QItemSelectionRange_delete")
}

type QItemSelectionRange struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QItemSelectionRange) UnsafeGoID() int64     { return this.vid }
func (this *QItemSelectionRange) UnsafeSetGoID(id int64) { this.vid = id }

func bindQItemSelectionRange(h uintptr) *QItemSelectionRange {
	if h == 0 {
		return nil
	}
	return &QItemSelectionRange{h: h}
}

func newownedQItemSelectionRange(h uintptr) *QItemSelectionRange {
	o := bindQItemSelectionRange(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QItemSelectionRange).Delete)
	}
	return o
}

func (this *QItemSelectionRange) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemSelectionRange) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemSelectionRange(h unsafe.Pointer) *QItemSelectionRange { return bindQItemSelectionRange(uintptr(h)) }
func UnsafeNewQItemSelectionRangeOwned(h unsafe.Pointer) *QItemSelectionRange { return newownedQItemSelectionRange(uintptr(h)) }

func (this *QItemSelectionRange) UnsafeRelease() { this.h = 0 }

func NewQItemSelectionRange() *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_new)
	__o := newownedQItemSelectionRange(__r)
	return __o
}

func NewQItemSelectionRange2(topL *QModelIndex, bottomR *QModelIndex) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_new2, uintptr(topL.UnsafePointer()), uintptr(bottomR.UnsafePointer()))
	__o := newownedQItemSelectionRange(__r)
	runtime.KeepAlive(topL)
	runtime.KeepAlive(bottomR)
	return __o
}

func NewQItemSelectionRange3(index *QModelIndex) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_new3, uintptr(index.UnsafePointer()))
	__o := newownedQItemSelectionRange(__r)
	runtime.KeepAlive(index)
	return __o
}

func NewQItemSelectionRange4(arg1 *QItemSelectionRange) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQItemSelectionRange(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QItemSelectionRange) Swap(other *QItemSelectionRange) {
	purego.SyscallN(addr_QItemSelectionRange_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QItemSelectionRange) Top() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_top, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) Left() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_left, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) Bottom() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_bottom, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) Right() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_right, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionRange) TopLeft() *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_topLeft, this.h)
	runtime.KeepAlive(this)
	return bindQPersistentModelIndex(__r)
}

func (this *QItemSelectionRange) BottomRight() *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_bottomRight, this.h)
	runtime.KeepAlive(this)
	return bindQPersistentModelIndex(__r)
}

func (this *QItemSelectionRange) Parent() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_parent, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QItemSelectionRange) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QItemSelectionRange) Contains(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_contains, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QItemSelectionRange) Contains2(row int, column int, parentIndex *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_contains2, this.h, uintptr(row), uintptr(column), uintptr(parentIndex.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parentIndex)
	return __r != 0
}

func (this *QItemSelectionRange) Intersects(other *QItemSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_intersects, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QItemSelectionRange) Intersected(other *QItemSelectionRange) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_intersected, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQItemSelectionRange(__r)
}

func (this *QItemSelectionRange) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionRange) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionRange) Indexes() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_indexes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelectionRange) OperatorAssign(arg1 *QItemSelectionRange) *QItemSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionRange_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQItemSelectionRange(__r)
}

func (this *QItemSelectionRange) OperatorEqual(rhs *QItemSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QItemSelectionRange) OperatorNotEqual(rhs *QItemSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionRange_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QItemSelectionRange) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemSelectionRange_delete(this.h)
	}
	this.UnsafeRelease()
}
