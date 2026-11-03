package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleValueInterface_currentValue uintptr
	addr_QAccessibleValueInterface_setCurrentValue uintptr
	addr_QAccessibleValueInterface_maximumValue uintptr
	addr_QAccessibleValueInterface_minimumValue uintptr
	addr_QAccessibleValueInterface_minimumStepSize uintptr
	addr_QAccessibleValueInterface_OperatorAssign uintptr
	pg_QAccessibleValueInterface__set_override func(uintptr, int32)
	pg_QAccessibleValueInterface__Go_new func(int64) uintptr
	pg_QAccessibleValueInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleValueInterface_currentValue, "QAccessibleValueInterface_currentValue")
	symaddr(&addr_QAccessibleValueInterface_setCurrentValue, "QAccessibleValueInterface_setCurrentValue")
	symaddr(&addr_QAccessibleValueInterface_maximumValue, "QAccessibleValueInterface_maximumValue")
	symaddr(&addr_QAccessibleValueInterface_minimumValue, "QAccessibleValueInterface_minimumValue")
	symaddr(&addr_QAccessibleValueInterface_minimumStepSize, "QAccessibleValueInterface_minimumStepSize")
	symaddr(&addr_QAccessibleValueInterface_OperatorAssign, "QAccessibleValueInterface_OperatorAssign")
	reg(&pg_QAccessibleValueInterface__set_override, "QAccessibleValueInterface__set_override")
	reg(&pg_QAccessibleValueInterface__Go_new, "QAccessibleValueInterface__Go_new")
	reg(&pg_QAccessibleValueInterface_delete, "QAccessibleValueInterface_delete")
}

type QAccessibleValueInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleValueInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleValueInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleValueInterface(h uintptr) *QAccessibleValueInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleValueInterface{h: h}
}

func newownedQAccessibleValueInterface(h uintptr) *QAccessibleValueInterface {
	o := bindQAccessibleValueInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleValueInterface).Delete)
	}
	return o
}

func (this *QAccessibleValueInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleValueInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleValueInterface(h unsafe.Pointer) *QAccessibleValueInterface { return bindQAccessibleValueInterface(uintptr(h)) }
func UnsafeNewQAccessibleValueInterfaceOwned(h unsafe.Pointer) *QAccessibleValueInterface { return newownedQAccessibleValueInterface(uintptr(h)) }

func (this *QAccessibleValueInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleValueInterface) CurrentValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueInterface_currentValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleValueInterface) SetCurrentValue(value *QVariant) {
	purego.SyscallN(addr_QAccessibleValueInterface_setCurrentValue, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QAccessibleValueInterface) MaximumValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueInterface_maximumValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleValueInterface) MinimumValue() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueInterface_minimumValue, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleValueInterface) MinimumStepSize() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueInterface_minimumStepSize, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleValueInterface) OperatorAssign(arg1 *QAccessibleValueInterface) *QAccessibleValueInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleValueInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleValueInterface(__r)
}

func QAccessibleValueInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 9131170551987814307:
		return 0
	case 3595292167320780357:
		return 1
	case 3939360635140333998:
		return 2
	case -2855573697932810260:
		return 3
	case -373090913906449796:
		return 4
	}
	return -1
}

func QAccessibleValueInterface_enable(h uintptr, idx int32) { pg_QAccessibleValueInterface__set_override(h, idx) }

func NewQAccessibleValueInterface() *QAccessibleValueInterface {
	id := nextVirtID()
	o := newownedQAccessibleValueInterface(pg_QAccessibleValueInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleValueInterface_virtIdx, QAccessibleValueInterface_enable)
	return o
}

func (this *QAccessibleValueInterface) OnCurrentValue(fn func() *QVariant) {
	registerVirt(this.UnsafeGoID(), 9131170551987814307, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "currentValue"
}

func (this *QAccessibleValueInterface) OnSetCurrentValue(fn func(*QVariant)) {
	registerVirt(this.UnsafeGoID(), 3595292167320780357, func(value uintptr) { fn(UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "setCurrentValue"
}

func (this *QAccessibleValueInterface) OnMaximumValue(fn func() *QVariant) {
	registerVirt(this.UnsafeGoID(), 3939360635140333998, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumValue"
}

func (this *QAccessibleValueInterface) OnMinimumValue(fn func() *QVariant) {
	registerVirt(this.UnsafeGoID(), -2855573697932810260, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumValue"
}

func (this *QAccessibleValueInterface) OnMinimumStepSize(fn func() *QVariant) {
	registerVirt(this.UnsafeGoID(), -373090913906449796, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumStepSize"
}

func (this *QAccessibleValueInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleValueInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
