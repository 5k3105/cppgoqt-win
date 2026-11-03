package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsItemAnimation_base func(uintptr) uintptr
	addr_QGraphicsItemAnimation_new uintptr
	addr_QGraphicsItemAnimation_new2 uintptr
	addr_QGraphicsItemAnimation_tr uintptr
	addr_QGraphicsItemAnimation_item uintptr
	addr_QGraphicsItemAnimation_setItem uintptr
	addr_QGraphicsItemAnimation_timeLine uintptr
	addr_QGraphicsItemAnimation_setTimeLine uintptr
	pg_QGraphicsItemAnimation_posAt func(uintptr, float64) uintptr
	pg_QGraphicsItemAnimation_setPosAt func(uintptr, float64, uintptr)
	pg_QGraphicsItemAnimation_transformAt func(uintptr, float64) uintptr
	pg_QGraphicsItemAnimation_rotationAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_setRotationAt func(uintptr, float64, float64)
	pg_QGraphicsItemAnimation_xTranslationAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_yTranslationAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_setTranslationAt func(uintptr, float64, float64, float64)
	pg_QGraphicsItemAnimation_verticalScaleAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_horizontalScaleAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_setScaleAt func(uintptr, float64, float64, float64)
	pg_QGraphicsItemAnimation_verticalShearAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_horizontalShearAt func(uintptr, float64) float64
	pg_QGraphicsItemAnimation_setShearAt func(uintptr, float64, float64, float64)
	addr_QGraphicsItemAnimation_clear uintptr
	pg_QGraphicsItemAnimation_setStep func(uintptr, float64)
	addr_QGraphicsItemAnimation_connect uintptr
	addr_QGraphicsItemAnimation_connect2 uintptr
	addr_QGraphicsItemAnimation_sender uintptr
	addr_QGraphicsItemAnimation_senderSignalIndex uintptr
	addr_QGraphicsItemAnimation_receivers uintptr
	addr_QGraphicsItemAnimation_isSignalConnected uintptr
	addr_QGraphicsItemAnimation_tr2 uintptr
	addr_QGraphicsItemAnimation_tr3 uintptr
	addr_QGraphicsItemAnimation_connect3 uintptr
	addr_QGraphicsItemAnimation_connect4 uintptr
	addr_QGraphicsItemAnimation_metaObject_default uintptr
	pg_QGraphicsItemAnimation_beforeAnimationStep_default func(uintptr, float64)
	pg_QGraphicsItemAnimation_afterAnimationStep_default func(uintptr, float64)
	addr_QGraphicsItemAnimation_event_default uintptr
	addr_QGraphicsItemAnimation_eventFilter_default uintptr
	addr_QGraphicsItemAnimation_timerEvent_default uintptr
	addr_QGraphicsItemAnimation_childEvent_default uintptr
	addr_QGraphicsItemAnimation_customEvent_default uintptr
	addr_QGraphicsItemAnimation_connectNotify_default uintptr
	addr_QGraphicsItemAnimation_disconnectNotify_default uintptr
	pg_QGraphicsItemAnimation_connect_destroyed func(uintptr, int64)
	pg_QGraphicsItemAnimation_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsItemAnimation_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsItemAnimation__disconnect func(int64)
	pg_QGraphicsItemAnimation__set_override func(uintptr, int32)
	pg_QGraphicsItemAnimation_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsItemAnimation_base, "QGraphicsItemAnimation_base")
	symaddr(&addr_QGraphicsItemAnimation_new, "QGraphicsItemAnimation_new")
	symaddr(&addr_QGraphicsItemAnimation_new2, "QGraphicsItemAnimation_new2")
	symaddr(&addr_QGraphicsItemAnimation_tr, "QGraphicsItemAnimation_tr")
	symaddr(&addr_QGraphicsItemAnimation_item, "QGraphicsItemAnimation_item")
	symaddr(&addr_QGraphicsItemAnimation_setItem, "QGraphicsItemAnimation_setItem")
	symaddr(&addr_QGraphicsItemAnimation_timeLine, "QGraphicsItemAnimation_timeLine")
	symaddr(&addr_QGraphicsItemAnimation_setTimeLine, "QGraphicsItemAnimation_setTimeLine")
	reg(&pg_QGraphicsItemAnimation_posAt, "QGraphicsItemAnimation_posAt")
	reg(&pg_QGraphicsItemAnimation_setPosAt, "QGraphicsItemAnimation_setPosAt")
	reg(&pg_QGraphicsItemAnimation_transformAt, "QGraphicsItemAnimation_transformAt")
	reg(&pg_QGraphicsItemAnimation_rotationAt, "QGraphicsItemAnimation_rotationAt")
	reg(&pg_QGraphicsItemAnimation_setRotationAt, "QGraphicsItemAnimation_setRotationAt")
	reg(&pg_QGraphicsItemAnimation_xTranslationAt, "QGraphicsItemAnimation_xTranslationAt")
	reg(&pg_QGraphicsItemAnimation_yTranslationAt, "QGraphicsItemAnimation_yTranslationAt")
	reg(&pg_QGraphicsItemAnimation_setTranslationAt, "QGraphicsItemAnimation_setTranslationAt")
	reg(&pg_QGraphicsItemAnimation_verticalScaleAt, "QGraphicsItemAnimation_verticalScaleAt")
	reg(&pg_QGraphicsItemAnimation_horizontalScaleAt, "QGraphicsItemAnimation_horizontalScaleAt")
	reg(&pg_QGraphicsItemAnimation_setScaleAt, "QGraphicsItemAnimation_setScaleAt")
	reg(&pg_QGraphicsItemAnimation_verticalShearAt, "QGraphicsItemAnimation_verticalShearAt")
	reg(&pg_QGraphicsItemAnimation_horizontalShearAt, "QGraphicsItemAnimation_horizontalShearAt")
	reg(&pg_QGraphicsItemAnimation_setShearAt, "QGraphicsItemAnimation_setShearAt")
	symaddr(&addr_QGraphicsItemAnimation_clear, "QGraphicsItemAnimation_clear")
	reg(&pg_QGraphicsItemAnimation_setStep, "QGraphicsItemAnimation_setStep")
	symaddr(&addr_QGraphicsItemAnimation_connect, "QGraphicsItemAnimation_connect")
	symaddr(&addr_QGraphicsItemAnimation_connect2, "QGraphicsItemAnimation_connect2")
	symaddr(&addr_QGraphicsItemAnimation_sender, "QGraphicsItemAnimation_sender")
	symaddr(&addr_QGraphicsItemAnimation_senderSignalIndex, "QGraphicsItemAnimation_senderSignalIndex")
	symaddr(&addr_QGraphicsItemAnimation_receivers, "QGraphicsItemAnimation_receivers")
	symaddr(&addr_QGraphicsItemAnimation_isSignalConnected, "QGraphicsItemAnimation_isSignalConnected")
	symaddr(&addr_QGraphicsItemAnimation_tr2, "QGraphicsItemAnimation_tr2")
	symaddr(&addr_QGraphicsItemAnimation_tr3, "QGraphicsItemAnimation_tr3")
	symaddr(&addr_QGraphicsItemAnimation_connect3, "QGraphicsItemAnimation_connect3")
	symaddr(&addr_QGraphicsItemAnimation_connect4, "QGraphicsItemAnimation_connect4")
	symaddr(&addr_QGraphicsItemAnimation_metaObject_default, "QGraphicsItemAnimation_metaObject_default")
	reg(&pg_QGraphicsItemAnimation_beforeAnimationStep_default, "QGraphicsItemAnimation_beforeAnimationStep_default")
	reg(&pg_QGraphicsItemAnimation_afterAnimationStep_default, "QGraphicsItemAnimation_afterAnimationStep_default")
	symaddr(&addr_QGraphicsItemAnimation_event_default, "QGraphicsItemAnimation_event_default")
	symaddr(&addr_QGraphicsItemAnimation_eventFilter_default, "QGraphicsItemAnimation_eventFilter_default")
	symaddr(&addr_QGraphicsItemAnimation_timerEvent_default, "QGraphicsItemAnimation_timerEvent_default")
	symaddr(&addr_QGraphicsItemAnimation_childEvent_default, "QGraphicsItemAnimation_childEvent_default")
	symaddr(&addr_QGraphicsItemAnimation_customEvent_default, "QGraphicsItemAnimation_customEvent_default")
	symaddr(&addr_QGraphicsItemAnimation_connectNotify_default, "QGraphicsItemAnimation_connectNotify_default")
	symaddr(&addr_QGraphicsItemAnimation_disconnectNotify_default, "QGraphicsItemAnimation_disconnectNotify_default")
	reg(&pg_QGraphicsItemAnimation_connect_destroyed, "QGraphicsItemAnimation_connect_destroyed")
	reg(&pg_QGraphicsItemAnimation_connect_destroyed2, "QGraphicsItemAnimation_connect_destroyed2")
	reg(&pg_QGraphicsItemAnimation_connect_objectNameChanged, "QGraphicsItemAnimation_connect_objectNameChanged")
	reg(&pg_QGraphicsItemAnimation__disconnect, "QGraphicsItemAnimation__disconnect")
	reg(&pg_QGraphicsItemAnimation__set_override, "QGraphicsItemAnimation__set_override")
	reg(&pg_QGraphicsItemAnimation_delete, "QGraphicsItemAnimation_delete")
}

type QGraphicsItemAnimation struct {
	h uintptr
	*QObject
}

func bindQGraphicsItemAnimation(h uintptr) *QGraphicsItemAnimation {
	if h == 0 {
		return nil
	}
	return &QGraphicsItemAnimation{h: h, QObject: bindQObject(pg_QGraphicsItemAnimation_base(h))}
}

func newownedQGraphicsItemAnimation(h uintptr) *QGraphicsItemAnimation { return bindQGraphicsItemAnimation(h) }

func (this *QGraphicsItemAnimation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsItemAnimation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsItemAnimation(h unsafe.Pointer) *QGraphicsItemAnimation { return bindQGraphicsItemAnimation(uintptr(h)) }
func UnsafeNewQGraphicsItemAnimationOwned(h unsafe.Pointer) *QGraphicsItemAnimation { return newownedQGraphicsItemAnimation(uintptr(h)) }

func (this *QGraphicsItemAnimation) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQGraphicsItemAnimation() *QGraphicsItemAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_new, uintptr(__goid))
	__o := newownedQGraphicsItemAnimation(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsItemAnimation_virtIdx, QGraphicsItemAnimation_enable)
	return __o
}

func NewQGraphicsItemAnimation2(parent *QObject) *QGraphicsItemAnimation {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsItemAnimation(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsItemAnimation_virtIdx, QGraphicsItemAnimation_enable)
	return __o
}

func QGraphicsItemAnimation_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsItemAnimation) Item() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_item, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItemAnimation) SetItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItemAnimation_setItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsItemAnimation) TimeLine() *QTimeLine {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_timeLine, this.h)
	runtime.KeepAlive(this)
	return bindQTimeLine(__r)
}

func (this *QGraphicsItemAnimation) SetTimeLine(timeLine *QTimeLine) {
	purego.SyscallN(addr_QGraphicsItemAnimation_setTimeLine, this.h, uintptr(timeLine.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timeLine)
}

func (this *QGraphicsItemAnimation) PosAt(step float64) *QPointF {
	__r := pg_QGraphicsItemAnimation_posAt(this.h, step)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItemAnimation) SetPosAt(step float64, pos *QPointF) {
	pg_QGraphicsItemAnimation_setPosAt(this.h, step, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsItemAnimation) TransformAt(step float64) *QTransform {
	__r := pg_QGraphicsItemAnimation_transformAt(this.h, step)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsItemAnimation) RotationAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_rotationAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) SetRotationAt(step float64, angle float64) {
	pg_QGraphicsItemAnimation_setRotationAt(this.h, step, angle)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) XTranslationAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_xTranslationAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) YTranslationAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_yTranslationAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) SetTranslationAt(step float64, dx float64, dy float64) {
	pg_QGraphicsItemAnimation_setTranslationAt(this.h, step, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) VerticalScaleAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_verticalScaleAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) HorizontalScaleAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_horizontalScaleAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) SetScaleAt(step float64, sx float64, sy float64) {
	pg_QGraphicsItemAnimation_setScaleAt(this.h, step, sx, sy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) VerticalShearAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_verticalShearAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) HorizontalShearAt(step float64) float64 {
	__r := pg_QGraphicsItemAnimation_horizontalShearAt(this.h, step)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItemAnimation) SetShearAt(step float64, sh float64, sv float64) {
	pg_QGraphicsItemAnimation_setShearAt(this.h, step, sh, sv)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) Clear() {
	purego.SyscallN(addr_QGraphicsItemAnimation_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) SetStep(x float64) {
	pg_QGraphicsItemAnimation_setStep(this.h, x)
	runtime.KeepAlive(this)
}

func QGraphicsItemAnimation_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsItemAnimation_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsItemAnimation) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsItemAnimation) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsItemAnimation) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsItemAnimation) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsItemAnimation_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsItemAnimation_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsItemAnimation_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsItemAnimation_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsItemAnimation) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsItemAnimation) BeforeAnimationStepDefault(step float64) {
	pg_QGraphicsItemAnimation_beforeAnimationStep_default(this.h, step)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) AfterAnimationStepDefault(step float64) {
	pg_QGraphicsItemAnimation_afterAnimationStep_default(this.h, step)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemAnimation) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItemAnimation) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemAnimation_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItemAnimation) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsItemAnimation_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemAnimation) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsItemAnimation_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemAnimation) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsItemAnimation_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemAnimation) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsItemAnimation_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsItemAnimation) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsItemAnimation_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsItemAnimation) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsItemAnimation_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsItemAnimation) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsItemAnimation_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsItemAnimation) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsItemAnimation_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsItemAnimation) Disconnect(handle int64) {
	pg_QGraphicsItemAnimation__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsItemAnimation_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -1838142790170279106:
		return 1
	case 7260939011549207179:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
		return 4
	case -5409927581812374498:
		return 5
	case -1185315367526857833:
		return 6
	case 3769868229123868474:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	}
	return -1
}

func QGraphicsItemAnimation_enable(h uintptr, idx int32) { pg_QGraphicsItemAnimation__set_override(h, idx) }

func (this *QGraphicsItemAnimation) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsItemAnimation) OnBeforeAnimationStep(fn func(float64)) {
	registerVirt(this.UnsafeGoID(), -1838142790170279106, func(step uintptr) { fn(math.Float64frombits(uint64(step))) }) // slot "beforeAnimationStep"
}

func (this *QGraphicsItemAnimation) OnAfterAnimationStep(fn func(float64)) {
	registerVirt(this.UnsafeGoID(), 7260939011549207179, func(step uintptr) { fn(math.Float64frombits(uint64(step))) }) // slot "afterAnimationStep"
}

func (this *QGraphicsItemAnimation) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsItemAnimation) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsItemAnimation) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsItemAnimation) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsItemAnimation) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsItemAnimation) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsItemAnimation) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsItemAnimation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsItemAnimation_delete(this.h)
	}
	this.UnsafeRelease()
}
