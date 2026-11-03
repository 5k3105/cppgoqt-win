package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionHeaderV2_chain uintptr
	addr_QStyleOptionHeaderV2_new uintptr
	addr_QStyleOptionHeaderV2_new2 uintptr
	addr_QStyleOptionHeaderV2_OperatorAssign uintptr
	addr_QStyleOptionHeaderV2_textElideMode_get uintptr
	addr_QStyleOptionHeaderV2_textElideMode_set uintptr
	addr_QStyleOptionHeaderV2_isSectionDragTarget_get uintptr
	addr_QStyleOptionHeaderV2_isSectionDragTarget_set uintptr
	addr_QStyleOptionHeaderV2_unused_get uintptr
	addr_QStyleOptionHeaderV2_unused_set uintptr
	pg_QStyleOptionHeaderV2_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionHeaderV2_chain, "QStyleOptionHeaderV2_chain")
	symaddr(&addr_QStyleOptionHeaderV2_new, "QStyleOptionHeaderV2_new")
	symaddr(&addr_QStyleOptionHeaderV2_new2, "QStyleOptionHeaderV2_new2")
	symaddr(&addr_QStyleOptionHeaderV2_OperatorAssign, "QStyleOptionHeaderV2_OperatorAssign")
	symaddr(&addr_QStyleOptionHeaderV2_textElideMode_get, "QStyleOptionHeaderV2_textElideMode_get")
	symaddr(&addr_QStyleOptionHeaderV2_textElideMode_set, "QStyleOptionHeaderV2_textElideMode_set")
	symaddr(&addr_QStyleOptionHeaderV2_isSectionDragTarget_get, "QStyleOptionHeaderV2_isSectionDragTarget_get")
	symaddr(&addr_QStyleOptionHeaderV2_isSectionDragTarget_set, "QStyleOptionHeaderV2_isSectionDragTarget_set")
	symaddr(&addr_QStyleOptionHeaderV2_unused_get, "QStyleOptionHeaderV2_unused_get")
	symaddr(&addr_QStyleOptionHeaderV2_unused_set, "QStyleOptionHeaderV2_unused_set")
	reg(&pg_QStyleOptionHeaderV2_delete, "QStyleOptionHeaderV2_delete")
}

type QStyleOptionHeaderV2 struct {
	h uintptr
	*QStyleOptionHeader
}

func bindQStyleOptionHeaderV2(h uintptr) *QStyleOptionHeaderV2 {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionHeaderV2_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionHeaderV2{h: h, QStyleOptionHeader: &QStyleOptionHeader{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionHeaderV2(h uintptr) *QStyleOptionHeaderV2 {
	o := bindQStyleOptionHeaderV2(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionHeaderV2).Delete)
	}
	return o
}

func (this *QStyleOptionHeaderV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionHeaderV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionHeaderV2(h unsafe.Pointer) *QStyleOptionHeaderV2 { return bindQStyleOptionHeaderV2(uintptr(h)) }
func UnsafeNewQStyleOptionHeaderV2Owned(h unsafe.Pointer) *QStyleOptionHeaderV2 { return newownedQStyleOptionHeaderV2(uintptr(h)) }

func (this *QStyleOptionHeaderV2) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionHeader != nil {
		this.QStyleOptionHeader.UnsafeRelease()
	}
}

func NewQStyleOptionHeaderV2() *QStyleOptionHeaderV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_new)
	__o := newownedQStyleOptionHeaderV2(__r)
	return __o
}

func NewQStyleOptionHeaderV22(other *QStyleOptionHeaderV2) *QStyleOptionHeaderV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionHeaderV2(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionHeaderV2) OperatorAssign(arg1 *QStyleOptionHeaderV2) *QStyleOptionHeaderV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionHeaderV2(__r)
}

func (this *QStyleOptionHeaderV2) TextElideMode() TextElideMode {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_textElideMode_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextElideMode(__r)
}

func (this *QStyleOptionHeaderV2) SetTextElideMode(textElideMode TextElideMode) {
	purego.SyscallN(addr_QStyleOptionHeaderV2_textElideMode_set, this.h, uintptr(textElideMode))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeaderV2) IsSectionDragTarget() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_isSectionDragTarget_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionHeaderV2) SetIsSectionDragTarget(isSectionDragTarget bool) {
	purego.SyscallN(addr_QStyleOptionHeaderV2_isSectionDragTarget_set, this.h, uintptr(b2i(isSectionDragTarget)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeaderV2) Unused() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeaderV2_unused_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionHeaderV2) SetUnused(unused int) {
	purego.SyscallN(addr_QStyleOptionHeaderV2_unused_set, this.h, uintptr(unused))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeaderV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionHeaderV2_delete(this.h)
	}
	this.UnsafeRelease()
}
