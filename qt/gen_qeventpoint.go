package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QEventPoint_new uintptr
	addr_QEventPoint_new2 uintptr
	addr_QEventPoint_new3 uintptr
	addr_QEventPoint_new4 uintptr
	addr_QEventPoint_new5 uintptr
	addr_QEventPoint_OperatorAssign uintptr
	addr_QEventPoint_OperatorEqual uintptr
	addr_QEventPoint_OperatorNotEqual uintptr
	addr_QEventPoint_swap uintptr
	addr_QEventPoint_position uintptr
	addr_QEventPoint_pressPosition uintptr
	addr_QEventPoint_grabPosition uintptr
	addr_QEventPoint_lastPosition uintptr
	addr_QEventPoint_scenePosition uintptr
	addr_QEventPoint_scenePressPosition uintptr
	addr_QEventPoint_sceneGrabPosition uintptr
	addr_QEventPoint_sceneLastPosition uintptr
	addr_QEventPoint_globalPosition uintptr
	addr_QEventPoint_globalPressPosition uintptr
	addr_QEventPoint_globalGrabPosition uintptr
	addr_QEventPoint_globalLastPosition uintptr
	addr_QEventPoint_normalizedPosition uintptr
	addr_QEventPoint_pos uintptr
	addr_QEventPoint_startPos uintptr
	addr_QEventPoint_scenePos uintptr
	addr_QEventPoint_startScenePos uintptr
	addr_QEventPoint_screenPos uintptr
	addr_QEventPoint_startScreenPos uintptr
	addr_QEventPoint_startNormalizedPos uintptr
	addr_QEventPoint_normalizedPos uintptr
	addr_QEventPoint_lastPos uintptr
	addr_QEventPoint_lastScenePos uintptr
	addr_QEventPoint_lastScreenPos uintptr
	addr_QEventPoint_lastNormalizedPos uintptr
	addr_QEventPoint_velocity uintptr
	addr_QEventPoint_state uintptr
	addr_QEventPoint_device uintptr
	addr_QEventPoint_id uintptr
	addr_QEventPoint_uniqueId uintptr
	addr_QEventPoint_timestamp uintptr
	addr_QEventPoint_lastTimestamp uintptr
	addr_QEventPoint_pressTimestamp uintptr
	pg_QEventPoint_timeHeld func(uintptr) float64
	pg_QEventPoint_pressure func(uintptr) float64
	pg_QEventPoint_rotation func(uintptr) float64
	addr_QEventPoint_ellipseDiameters uintptr
	addr_QEventPoint_isAccepted uintptr
	addr_QEventPoint_setAccepted uintptr
	addr_QEventPoint_setAccepted2 uintptr
	pg_QEventPoint_delete func(uintptr)
)

func init() {
	symaddr(&addr_QEventPoint_new, "QEventPoint_new")
	symaddr(&addr_QEventPoint_new2, "QEventPoint_new2")
	symaddr(&addr_QEventPoint_new3, "QEventPoint_new3")
	symaddr(&addr_QEventPoint_new4, "QEventPoint_new4")
	symaddr(&addr_QEventPoint_new5, "QEventPoint_new5")
	symaddr(&addr_QEventPoint_OperatorAssign, "QEventPoint_OperatorAssign")
	symaddr(&addr_QEventPoint_OperatorEqual, "QEventPoint_OperatorEqual")
	symaddr(&addr_QEventPoint_OperatorNotEqual, "QEventPoint_OperatorNotEqual")
	symaddr(&addr_QEventPoint_swap, "QEventPoint_swap")
	symaddr(&addr_QEventPoint_position, "QEventPoint_position")
	symaddr(&addr_QEventPoint_pressPosition, "QEventPoint_pressPosition")
	symaddr(&addr_QEventPoint_grabPosition, "QEventPoint_grabPosition")
	symaddr(&addr_QEventPoint_lastPosition, "QEventPoint_lastPosition")
	symaddr(&addr_QEventPoint_scenePosition, "QEventPoint_scenePosition")
	symaddr(&addr_QEventPoint_scenePressPosition, "QEventPoint_scenePressPosition")
	symaddr(&addr_QEventPoint_sceneGrabPosition, "QEventPoint_sceneGrabPosition")
	symaddr(&addr_QEventPoint_sceneLastPosition, "QEventPoint_sceneLastPosition")
	symaddr(&addr_QEventPoint_globalPosition, "QEventPoint_globalPosition")
	symaddr(&addr_QEventPoint_globalPressPosition, "QEventPoint_globalPressPosition")
	symaddr(&addr_QEventPoint_globalGrabPosition, "QEventPoint_globalGrabPosition")
	symaddr(&addr_QEventPoint_globalLastPosition, "QEventPoint_globalLastPosition")
	symaddr(&addr_QEventPoint_normalizedPosition, "QEventPoint_normalizedPosition")
	symaddr(&addr_QEventPoint_pos, "QEventPoint_pos")
	symaddr(&addr_QEventPoint_startPos, "QEventPoint_startPos")
	symaddr(&addr_QEventPoint_scenePos, "QEventPoint_scenePos")
	symaddr(&addr_QEventPoint_startScenePos, "QEventPoint_startScenePos")
	symaddr(&addr_QEventPoint_screenPos, "QEventPoint_screenPos")
	symaddr(&addr_QEventPoint_startScreenPos, "QEventPoint_startScreenPos")
	symaddr(&addr_QEventPoint_startNormalizedPos, "QEventPoint_startNormalizedPos")
	symaddr(&addr_QEventPoint_normalizedPos, "QEventPoint_normalizedPos")
	symaddr(&addr_QEventPoint_lastPos, "QEventPoint_lastPos")
	symaddr(&addr_QEventPoint_lastScenePos, "QEventPoint_lastScenePos")
	symaddr(&addr_QEventPoint_lastScreenPos, "QEventPoint_lastScreenPos")
	symaddr(&addr_QEventPoint_lastNormalizedPos, "QEventPoint_lastNormalizedPos")
	symaddr(&addr_QEventPoint_velocity, "QEventPoint_velocity")
	symaddr(&addr_QEventPoint_state, "QEventPoint_state")
	symaddr(&addr_QEventPoint_device, "QEventPoint_device")
	symaddr(&addr_QEventPoint_id, "QEventPoint_id")
	symaddr(&addr_QEventPoint_uniqueId, "QEventPoint_uniqueId")
	symaddr(&addr_QEventPoint_timestamp, "QEventPoint_timestamp")
	symaddr(&addr_QEventPoint_lastTimestamp, "QEventPoint_lastTimestamp")
	symaddr(&addr_QEventPoint_pressTimestamp, "QEventPoint_pressTimestamp")
	reg(&pg_QEventPoint_timeHeld, "QEventPoint_timeHeld")
	reg(&pg_QEventPoint_pressure, "QEventPoint_pressure")
	reg(&pg_QEventPoint_rotation, "QEventPoint_rotation")
	symaddr(&addr_QEventPoint_ellipseDiameters, "QEventPoint_ellipseDiameters")
	symaddr(&addr_QEventPoint_isAccepted, "QEventPoint_isAccepted")
	symaddr(&addr_QEventPoint_setAccepted, "QEventPoint_setAccepted")
	symaddr(&addr_QEventPoint_setAccepted2, "QEventPoint_setAccepted2")
	reg(&pg_QEventPoint_delete, "QEventPoint_delete")
}

type QEventPoint struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QEventPoint) UnsafeGoID() int64     { return this.vid }
func (this *QEventPoint) UnsafeSetGoID(id int64) { this.vid = id }

func bindQEventPoint(h uintptr) *QEventPoint {
	if h == 0 {
		return nil
	}
	return &QEventPoint{h: h}
}

func newownedQEventPoint(h uintptr) *QEventPoint {
	o := bindQEventPoint(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QEventPoint).Delete)
	}
	return o
}

func (this *QEventPoint) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEventPoint) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEventPoint(h unsafe.Pointer) *QEventPoint { return bindQEventPoint(uintptr(h)) }
func UnsafeNewQEventPointOwned(h unsafe.Pointer) *QEventPoint { return newownedQEventPoint(uintptr(h)) }

func (this *QEventPoint) UnsafeRelease() { this.h = 0 }

func NewQEventPoint() *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_new)
	__o := newownedQEventPoint(__r)
	return __o
}

func NewQEventPoint2(pointId int, state QEventPoint__State, scenePosition *QPointF, globalPosition *QPointF) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_new2, uintptr(pointId), uintptr(state), uintptr(scenePosition.UnsafePointer()), uintptr(globalPosition.UnsafePointer()))
	__o := newownedQEventPoint(__r)
	runtime.KeepAlive(scenePosition)
	runtime.KeepAlive(globalPosition)
	return __o
}

func NewQEventPoint3(other *QEventPoint) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_new3, uintptr(other.UnsafePointer()))
	__o := newownedQEventPoint(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQEventPoint4(id int) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_new4, uintptr(id))
	__o := newownedQEventPoint(__r)
	return __o
}

func NewQEventPoint5(id int, device *QPointingDevice) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_new5, uintptr(id), uintptr(device.UnsafePointer()))
	__o := newownedQEventPoint(__r)
	runtime.KeepAlive(device)
	return __o
}

func (this *QEventPoint) OperatorAssign(other *QEventPoint) *QEventPoint {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQEventPoint(__r)
}

func (this *QEventPoint) OperatorEqual(other *QEventPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QEventPoint) OperatorNotEqual(other *QEventPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QEventPoint) Swap(other *QEventPoint) {
	purego.SyscallN(addr_QEventPoint_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QEventPoint) Position() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) PressPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_pressPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) GrabPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_grabPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) LastPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_lastPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) ScenePosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_scenePosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) ScenePressPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_scenePressPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) SceneGrabPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_sceneGrabPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) SceneLastPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_sceneLastPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) GlobalPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_globalPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) GlobalPressPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_globalPressPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) GlobalGrabPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_globalGrabPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) GlobalLastPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_globalLastPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) NormalizedPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_normalizedPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) StartPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_startPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) StartScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_startScenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) ScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_screenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) StartScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_startScreenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) StartNormalizedPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_startNormalizedPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) NormalizedPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_normalizedPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) LastPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_lastPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) LastScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_lastScenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) LastScreenPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_lastScreenPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) LastNormalizedPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_lastNormalizedPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QEventPoint) Velocity() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_velocity, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QEventPoint) State() QEventPoint__State {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QEventPoint__State(__r)
}

func (this *QEventPoint) Device() *QPointingDevice {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_device, this.h)
	runtime.KeepAlive(this)
	return bindQPointingDevice(__r)
}

func (this *QEventPoint) Id() int {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QEventPoint) UniqueId() *QPointingDeviceUniqueId {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_uniqueId, this.h)
	runtime.KeepAlive(this)
	return newownedQPointingDeviceUniqueId(__r)
}

func (this *QEventPoint) Timestamp() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_timestamp, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QEventPoint) LastTimestamp() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_lastTimestamp, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QEventPoint) PressTimestamp() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_pressTimestamp, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QEventPoint) TimeHeld() float64 {
	__r := pg_QEventPoint_timeHeld(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEventPoint) Pressure() float64 {
	__r := pg_QEventPoint_pressure(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEventPoint) Rotation() float64 {
	__r := pg_QEventPoint_rotation(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEventPoint) EllipseDiameters() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QEventPoint_ellipseDiameters, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QEventPoint) IsAccepted() bool {
	__r0, _, _ := purego.SyscallN(addr_QEventPoint_isAccepted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEventPoint) SetAccepted() {
	purego.SyscallN(addr_QEventPoint_setAccepted, this.h)
	runtime.KeepAlive(this)
}

func (this *QEventPoint) SetAcceptedWithAccepted(accepted bool) {
	purego.SyscallN(addr_QEventPoint_setAccepted2, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QEventPoint) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEventPoint_delete(this.h)
	}
	this.UnsafeRelease()
}
