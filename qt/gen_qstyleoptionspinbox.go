package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionSpinBox_chain uintptr
	addr_QStyleOptionSpinBox_new uintptr
	addr_QStyleOptionSpinBox_new2 uintptr
	addr_QStyleOptionSpinBox_OperatorAssign uintptr
	addr_QStyleOptionSpinBox_buttonSymbols_get uintptr
	addr_QStyleOptionSpinBox_buttonSymbols_set uintptr
	addr_QStyleOptionSpinBox_stepEnabled_get uintptr
	addr_QStyleOptionSpinBox_stepEnabled_set uintptr
	addr_QStyleOptionSpinBox_frame_get uintptr
	addr_QStyleOptionSpinBox_frame_set uintptr
	pg_QStyleOptionSpinBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionSpinBox_chain, "QStyleOptionSpinBox_chain")
	symaddr(&addr_QStyleOptionSpinBox_new, "QStyleOptionSpinBox_new")
	symaddr(&addr_QStyleOptionSpinBox_new2, "QStyleOptionSpinBox_new2")
	symaddr(&addr_QStyleOptionSpinBox_OperatorAssign, "QStyleOptionSpinBox_OperatorAssign")
	symaddr(&addr_QStyleOptionSpinBox_buttonSymbols_get, "QStyleOptionSpinBox_buttonSymbols_get")
	symaddr(&addr_QStyleOptionSpinBox_buttonSymbols_set, "QStyleOptionSpinBox_buttonSymbols_set")
	symaddr(&addr_QStyleOptionSpinBox_stepEnabled_get, "QStyleOptionSpinBox_stepEnabled_get")
	symaddr(&addr_QStyleOptionSpinBox_stepEnabled_set, "QStyleOptionSpinBox_stepEnabled_set")
	symaddr(&addr_QStyleOptionSpinBox_frame_get, "QStyleOptionSpinBox_frame_get")
	symaddr(&addr_QStyleOptionSpinBox_frame_set, "QStyleOptionSpinBox_frame_set")
	reg(&pg_QStyleOptionSpinBox_delete, "QStyleOptionSpinBox_delete")
}

type QStyleOptionSpinBox struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionSpinBox(h uintptr) *QStyleOptionSpinBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionSpinBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionSpinBox{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionSpinBox(h uintptr) *QStyleOptionSpinBox {
	o := bindQStyleOptionSpinBox(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionSpinBox).Delete)
	}
	return o
}

func (this *QStyleOptionSpinBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionSpinBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionSpinBox(h unsafe.Pointer) *QStyleOptionSpinBox { return bindQStyleOptionSpinBox(uintptr(h)) }
func UnsafeNewQStyleOptionSpinBoxOwned(h unsafe.Pointer) *QStyleOptionSpinBox { return newownedQStyleOptionSpinBox(uintptr(h)) }

func (this *QStyleOptionSpinBox) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionSpinBox() *QStyleOptionSpinBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_new)
	__o := newownedQStyleOptionSpinBox(__r)
	return __o
}

func NewQStyleOptionSpinBox2(other *QStyleOptionSpinBox) *QStyleOptionSpinBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionSpinBox(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionSpinBox) OperatorAssign(arg1 *QStyleOptionSpinBox) *QStyleOptionSpinBox {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionSpinBox(__r)
}

func (this *QStyleOptionSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_buttonSymbols_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__ButtonSymbols(__r)
}

func (this *QStyleOptionSpinBox) SetButtonSymbols(buttonSymbols QAbstractSpinBox__ButtonSymbols) {
	purego.SyscallN(addr_QStyleOptionSpinBox_buttonSymbols_set, this.h, uintptr(buttonSymbols))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSpinBox) StepEnabled() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_stepEnabled_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QStyleOptionSpinBox) SetStepEnabled(stepEnabled QAbstractSpinBox__StepEnabledFlag) {
	purego.SyscallN(addr_QStyleOptionSpinBox_stepEnabled_set, this.h, uintptr(stepEnabled))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSpinBox) Frame() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSpinBox_frame_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionSpinBox) SetFrame(frame bool) {
	purego.SyscallN(addr_QStyleOptionSpinBox_frame_set, this.h, uintptr(b2i(frame)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSpinBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionSpinBox_delete(this.h)
	}
	this.UnsafeRelease()
}
