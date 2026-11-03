package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionMenuItemV2_chain uintptr
	addr_QStyleOptionMenuItemV2_new uintptr
	addr_QStyleOptionMenuItemV2_new2 uintptr
	addr_QStyleOptionMenuItemV2_OperatorAssign uintptr
	addr_QStyleOptionMenuItemV2_mouseDown_get uintptr
	addr_QStyleOptionMenuItemV2_mouseDown_set uintptr
	addr_QStyleOptionMenuItemV2_unused_get uintptr
	addr_QStyleOptionMenuItemV2_unused_set uintptr
	pg_QStyleOptionMenuItemV2_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionMenuItemV2_chain, "QStyleOptionMenuItemV2_chain")
	symaddr(&addr_QStyleOptionMenuItemV2_new, "QStyleOptionMenuItemV2_new")
	symaddr(&addr_QStyleOptionMenuItemV2_new2, "QStyleOptionMenuItemV2_new2")
	symaddr(&addr_QStyleOptionMenuItemV2_OperatorAssign, "QStyleOptionMenuItemV2_OperatorAssign")
	symaddr(&addr_QStyleOptionMenuItemV2_mouseDown_get, "QStyleOptionMenuItemV2_mouseDown_get")
	symaddr(&addr_QStyleOptionMenuItemV2_mouseDown_set, "QStyleOptionMenuItemV2_mouseDown_set")
	symaddr(&addr_QStyleOptionMenuItemV2_unused_get, "QStyleOptionMenuItemV2_unused_get")
	symaddr(&addr_QStyleOptionMenuItemV2_unused_set, "QStyleOptionMenuItemV2_unused_set")
	reg(&pg_QStyleOptionMenuItemV2_delete, "QStyleOptionMenuItemV2_delete")
}

type QStyleOptionMenuItemV2 struct {
	h uintptr
	*QStyleOptionMenuItem
}

func bindQStyleOptionMenuItemV2(h uintptr) *QStyleOptionMenuItemV2 {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionMenuItemV2_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionMenuItemV2{h: h, QStyleOptionMenuItem: &QStyleOptionMenuItem{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionMenuItemV2(h uintptr) *QStyleOptionMenuItemV2 {
	o := bindQStyleOptionMenuItemV2(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionMenuItemV2).Delete)
	}
	return o
}

func (this *QStyleOptionMenuItemV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionMenuItemV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionMenuItemV2(h unsafe.Pointer) *QStyleOptionMenuItemV2 { return bindQStyleOptionMenuItemV2(uintptr(h)) }
func UnsafeNewQStyleOptionMenuItemV2Owned(h unsafe.Pointer) *QStyleOptionMenuItemV2 { return newownedQStyleOptionMenuItemV2(uintptr(h)) }

func (this *QStyleOptionMenuItemV2) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionMenuItem != nil {
		this.QStyleOptionMenuItem.UnsafeRelease()
	}
}

func NewQStyleOptionMenuItemV2() *QStyleOptionMenuItemV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItemV2_new)
	__o := newownedQStyleOptionMenuItemV2(__r)
	return __o
}

func NewQStyleOptionMenuItemV22(other *QStyleOptionMenuItemV2) *QStyleOptionMenuItemV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItemV2_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionMenuItemV2(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionMenuItemV2) OperatorAssign(arg1 *QStyleOptionMenuItemV2) *QStyleOptionMenuItemV2 {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionMenuItemV2_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionMenuItemV2(__r)
}

func (this *QStyleOptionMenuItemV2) MouseDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItemV2_mouseDown_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionMenuItemV2) SetMouseDown(mouseDown bool) {
	purego.SyscallN(addr_QStyleOptionMenuItemV2_mouseDown_set, this.h, uintptr(b2i(mouseDown)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItemV2) Unused() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionMenuItemV2_unused_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionMenuItemV2) SetUnused(unused int) {
	purego.SyscallN(addr_QStyleOptionMenuItemV2_unused_set, this.h, uintptr(unused))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionMenuItemV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionMenuItemV2_delete(this.h)
	}
	this.UnsafeRelease()
}
