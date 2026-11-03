package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTabletEvent_chain uintptr
	pg_QTabletEvent_new func(int32, uintptr, uintptr, uintptr, float64, float32, float32, float32, float64, float32, int32, int32, int32) uintptr
	addr_QTabletEvent_pos uintptr
	addr_QTabletEvent_globalPos uintptr
	addr_QTabletEvent_posF uintptr
	addr_QTabletEvent_globalPosF uintptr
	addr_QTabletEvent_x uintptr
	addr_QTabletEvent_y uintptr
	addr_QTabletEvent_globalX uintptr
	addr_QTabletEvent_globalY uintptr
	pg_QTabletEvent_hiResGlobalX func(uintptr) float64
	pg_QTabletEvent_hiResGlobalY func(uintptr) float64
	addr_QTabletEvent_uniqueId uintptr
	pg_QTabletEvent_pressure func(uintptr) float64
	pg_QTabletEvent_rotation func(uintptr) float64
	pg_QTabletEvent_z func(uintptr) float64
	pg_QTabletEvent_tangentialPressure func(uintptr) float64
	pg_QTabletEvent_xTilt func(uintptr) float64
	pg_QTabletEvent_yTilt func(uintptr) float64
	addr_QTabletEvent_registerEventType uintptr
	addr_QTabletEvent_registerEventType2 uintptr
	pg_QTabletEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTabletEvent_chain, "QTabletEvent_chain")
	reg(&pg_QTabletEvent_new, "QTabletEvent_new")
	symaddr(&addr_QTabletEvent_pos, "QTabletEvent_pos")
	symaddr(&addr_QTabletEvent_globalPos, "QTabletEvent_globalPos")
	symaddr(&addr_QTabletEvent_posF, "QTabletEvent_posF")
	symaddr(&addr_QTabletEvent_globalPosF, "QTabletEvent_globalPosF")
	symaddr(&addr_QTabletEvent_x, "QTabletEvent_x")
	symaddr(&addr_QTabletEvent_y, "QTabletEvent_y")
	symaddr(&addr_QTabletEvent_globalX, "QTabletEvent_globalX")
	symaddr(&addr_QTabletEvent_globalY, "QTabletEvent_globalY")
	reg(&pg_QTabletEvent_hiResGlobalX, "QTabletEvent_hiResGlobalX")
	reg(&pg_QTabletEvent_hiResGlobalY, "QTabletEvent_hiResGlobalY")
	symaddr(&addr_QTabletEvent_uniqueId, "QTabletEvent_uniqueId")
	reg(&pg_QTabletEvent_pressure, "QTabletEvent_pressure")
	reg(&pg_QTabletEvent_rotation, "QTabletEvent_rotation")
	reg(&pg_QTabletEvent_z, "QTabletEvent_z")
	reg(&pg_QTabletEvent_tangentialPressure, "QTabletEvent_tangentialPressure")
	reg(&pg_QTabletEvent_xTilt, "QTabletEvent_xTilt")
	reg(&pg_QTabletEvent_yTilt, "QTabletEvent_yTilt")
	symaddr(&addr_QTabletEvent_registerEventType, "QTabletEvent_registerEventType")
	symaddr(&addr_QTabletEvent_registerEventType2, "QTabletEvent_registerEventType2")
	reg(&pg_QTabletEvent_delete, "QTabletEvent_delete")
}

type QTabletEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQTabletEvent(h uintptr) *QTabletEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QTabletEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTabletEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQTabletEvent(h uintptr) *QTabletEvent {
	o := bindQTabletEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTabletEvent).Delete)
	}
	return o
}

func (this *QTabletEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTabletEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTabletEvent(h unsafe.Pointer) *QTabletEvent { return bindQTabletEvent(uintptr(h)) }
func UnsafeNewQTabletEventOwned(h unsafe.Pointer) *QTabletEvent { return newownedQTabletEvent(uintptr(h)) }

func (this *QTabletEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQTabletEvent(t QEvent__Type, device *QPointingDevice, pos *QPointF, globalPos *QPointF, pressure float64, xTilt float32, yTilt float32, tangentialPressure float32, rotation float64, z float32, keyState KeyboardModifier, button MouseButton, buttons MouseButton) *QTabletEvent {
	__o := newownedQTabletEvent(pg_QTabletEvent_new(int32(t), uintptr(device.UnsafePointer()), uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), pressure, xTilt, yTilt, tangentialPressure, rotation, z, int32(keyState), int32(button), int32(buttons)))
	runtime.KeepAlive(device)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	return __o
}

func (this *QTabletEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTabletEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTabletEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTabletEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTabletEvent) PosF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTabletEvent_posF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTabletEvent) GlobalPosF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTabletEvent_globalPosF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QTabletEvent) X() int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabletEvent) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabletEvent) GlobalX() int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_globalX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabletEvent) GlobalY() int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_globalY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabletEvent) HiResGlobalX() float64 {
	__r := pg_QTabletEvent_hiResGlobalX(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) HiResGlobalY() float64 {
	__r := pg_QTabletEvent_hiResGlobalY(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) UniqueId() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_uniqueId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) Pressure() float64 {
	__r := pg_QTabletEvent_pressure(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) Rotation() float64 {
	__r := pg_QTabletEvent_rotation(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) Z() float64 {
	__r := pg_QTabletEvent_z(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) TangentialPressure() float64 {
	__r := pg_QTabletEvent_tangentialPressure(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) XTilt() float64 {
	__r := pg_QTabletEvent_xTilt(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTabletEvent) YTilt() float64 {
	__r := pg_QTabletEvent_yTilt(this.h)
	runtime.KeepAlive(this)
	return __r
}

func QTabletEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QTabletEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QTabletEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QTabletEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTabletEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
