package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionSlider_chain uintptr
	addr_QStyleOptionSlider_new uintptr
	addr_QStyleOptionSlider_new2 uintptr
	addr_QStyleOptionSlider_OperatorAssign uintptr
	addr_QStyleOptionSlider_orientation_get uintptr
	addr_QStyleOptionSlider_orientation_set uintptr
	addr_QStyleOptionSlider_minimum_get uintptr
	addr_QStyleOptionSlider_minimum_set uintptr
	addr_QStyleOptionSlider_maximum_get uintptr
	addr_QStyleOptionSlider_maximum_set uintptr
	addr_QStyleOptionSlider_tickPosition_get uintptr
	addr_QStyleOptionSlider_tickPosition_set uintptr
	addr_QStyleOptionSlider_tickInterval_get uintptr
	addr_QStyleOptionSlider_tickInterval_set uintptr
	addr_QStyleOptionSlider_upsideDown_get uintptr
	addr_QStyleOptionSlider_upsideDown_set uintptr
	addr_QStyleOptionSlider_sliderPosition_get uintptr
	addr_QStyleOptionSlider_sliderPosition_set uintptr
	addr_QStyleOptionSlider_sliderValue_get uintptr
	addr_QStyleOptionSlider_sliderValue_set uintptr
	addr_QStyleOptionSlider_singleStep_get uintptr
	addr_QStyleOptionSlider_singleStep_set uintptr
	addr_QStyleOptionSlider_pageStep_get uintptr
	addr_QStyleOptionSlider_pageStep_set uintptr
	pg_QStyleOptionSlider_notchTarget_get func(uintptr) float64
	pg_QStyleOptionSlider_notchTarget_set func(uintptr, float64)
	addr_QStyleOptionSlider_dialWrapping_get uintptr
	addr_QStyleOptionSlider_dialWrapping_set uintptr
	addr_QStyleOptionSlider_keyboardModifiers_get uintptr
	addr_QStyleOptionSlider_keyboardModifiers_set uintptr
	pg_QStyleOptionSlider_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionSlider_chain, "QStyleOptionSlider_chain")
	symaddr(&addr_QStyleOptionSlider_new, "QStyleOptionSlider_new")
	symaddr(&addr_QStyleOptionSlider_new2, "QStyleOptionSlider_new2")
	symaddr(&addr_QStyleOptionSlider_OperatorAssign, "QStyleOptionSlider_OperatorAssign")
	symaddr(&addr_QStyleOptionSlider_orientation_get, "QStyleOptionSlider_orientation_get")
	symaddr(&addr_QStyleOptionSlider_orientation_set, "QStyleOptionSlider_orientation_set")
	symaddr(&addr_QStyleOptionSlider_minimum_get, "QStyleOptionSlider_minimum_get")
	symaddr(&addr_QStyleOptionSlider_minimum_set, "QStyleOptionSlider_minimum_set")
	symaddr(&addr_QStyleOptionSlider_maximum_get, "QStyleOptionSlider_maximum_get")
	symaddr(&addr_QStyleOptionSlider_maximum_set, "QStyleOptionSlider_maximum_set")
	symaddr(&addr_QStyleOptionSlider_tickPosition_get, "QStyleOptionSlider_tickPosition_get")
	symaddr(&addr_QStyleOptionSlider_tickPosition_set, "QStyleOptionSlider_tickPosition_set")
	symaddr(&addr_QStyleOptionSlider_tickInterval_get, "QStyleOptionSlider_tickInterval_get")
	symaddr(&addr_QStyleOptionSlider_tickInterval_set, "QStyleOptionSlider_tickInterval_set")
	symaddr(&addr_QStyleOptionSlider_upsideDown_get, "QStyleOptionSlider_upsideDown_get")
	symaddr(&addr_QStyleOptionSlider_upsideDown_set, "QStyleOptionSlider_upsideDown_set")
	symaddr(&addr_QStyleOptionSlider_sliderPosition_get, "QStyleOptionSlider_sliderPosition_get")
	symaddr(&addr_QStyleOptionSlider_sliderPosition_set, "QStyleOptionSlider_sliderPosition_set")
	symaddr(&addr_QStyleOptionSlider_sliderValue_get, "QStyleOptionSlider_sliderValue_get")
	symaddr(&addr_QStyleOptionSlider_sliderValue_set, "QStyleOptionSlider_sliderValue_set")
	symaddr(&addr_QStyleOptionSlider_singleStep_get, "QStyleOptionSlider_singleStep_get")
	symaddr(&addr_QStyleOptionSlider_singleStep_set, "QStyleOptionSlider_singleStep_set")
	symaddr(&addr_QStyleOptionSlider_pageStep_get, "QStyleOptionSlider_pageStep_get")
	symaddr(&addr_QStyleOptionSlider_pageStep_set, "QStyleOptionSlider_pageStep_set")
	reg(&pg_QStyleOptionSlider_notchTarget_get, "QStyleOptionSlider_notchTarget_get")
	reg(&pg_QStyleOptionSlider_notchTarget_set, "QStyleOptionSlider_notchTarget_set")
	symaddr(&addr_QStyleOptionSlider_dialWrapping_get, "QStyleOptionSlider_dialWrapping_get")
	symaddr(&addr_QStyleOptionSlider_dialWrapping_set, "QStyleOptionSlider_dialWrapping_set")
	symaddr(&addr_QStyleOptionSlider_keyboardModifiers_get, "QStyleOptionSlider_keyboardModifiers_get")
	symaddr(&addr_QStyleOptionSlider_keyboardModifiers_set, "QStyleOptionSlider_keyboardModifiers_set")
	reg(&pg_QStyleOptionSlider_delete, "QStyleOptionSlider_delete")
}

type QStyleOptionSlider struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionSlider(h uintptr) *QStyleOptionSlider {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionSlider_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionSlider{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionSlider(h uintptr) *QStyleOptionSlider {
	o := bindQStyleOptionSlider(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionSlider).Delete)
	}
	return o
}

func (this *QStyleOptionSlider) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionSlider) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionSlider(h unsafe.Pointer) *QStyleOptionSlider { return bindQStyleOptionSlider(uintptr(h)) }
func UnsafeNewQStyleOptionSliderOwned(h unsafe.Pointer) *QStyleOptionSlider { return newownedQStyleOptionSlider(uintptr(h)) }

func (this *QStyleOptionSlider) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionSlider() *QStyleOptionSlider {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSlider_new)
	__o := newownedQStyleOptionSlider(__r)
	return __o
}

func NewQStyleOptionSlider2(other *QStyleOptionSlider) *QStyleOptionSlider {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSlider_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionSlider(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionSlider) OperatorAssign(arg1 *QStyleOptionSlider) *QStyleOptionSlider {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSlider_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionSlider(__r)
}

func (this *QStyleOptionSlider) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_orientation_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QStyleOptionSlider) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QStyleOptionSlider_orientation_set, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_minimum_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetMinimum(minimum int) {
	purego.SyscallN(addr_QStyleOptionSlider_minimum_set, this.h, uintptr(minimum))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_maximum_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetMaximum(maximum int) {
	purego.SyscallN(addr_QStyleOptionSlider_maximum_set, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) TickPosition() QSlider__TickPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_tickPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSlider__TickPosition(__r)
}

func (this *QStyleOptionSlider) SetTickPosition(tickPosition QSlider__TickPosition) {
	purego.SyscallN(addr_QStyleOptionSlider_tickPosition_set, this.h, uintptr(tickPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) TickInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_tickInterval_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetTickInterval(tickInterval int) {
	purego.SyscallN(addr_QStyleOptionSlider_tickInterval_set, this.h, uintptr(tickInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) UpsideDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_upsideDown_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionSlider) SetUpsideDown(upsideDown bool) {
	purego.SyscallN(addr_QStyleOptionSlider_upsideDown_set, this.h, uintptr(b2i(upsideDown)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) SliderPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_sliderPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetSliderPosition(sliderPosition int) {
	purego.SyscallN(addr_QStyleOptionSlider_sliderPosition_set, this.h, uintptr(sliderPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) SliderValue() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_sliderValue_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetSliderValue(sliderValue int) {
	purego.SyscallN(addr_QStyleOptionSlider_sliderValue_set, this.h, uintptr(sliderValue))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) SingleStep() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_singleStep_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetSingleStep(singleStep int) {
	purego.SyscallN(addr_QStyleOptionSlider_singleStep_set, this.h, uintptr(singleStep))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) PageStep() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_pageStep_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionSlider) SetPageStep(pageStep int) {
	purego.SyscallN(addr_QStyleOptionSlider_pageStep_set, this.h, uintptr(pageStep))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) NotchTarget() float64 {
	__r := pg_QStyleOptionSlider_notchTarget_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStyleOptionSlider) SetNotchTarget(notchTarget float64) {
	pg_QStyleOptionSlider_notchTarget_set(this.h, notchTarget)
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) DialWrapping() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_dialWrapping_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionSlider) SetDialWrapping(dialWrapping bool) {
	purego.SyscallN(addr_QStyleOptionSlider_dialWrapping_set, this.h, uintptr(b2i(dialWrapping)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) KeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSlider_keyboardModifiers_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QStyleOptionSlider) SetKeyboardModifiers(keyboardModifiers KeyboardModifier) {
	purego.SyscallN(addr_QStyleOptionSlider_keyboardModifiers_set, this.h, uintptr(keyboardModifiers))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSlider) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionSlider_delete(this.h)
	}
	this.UnsafeRelease()
}
