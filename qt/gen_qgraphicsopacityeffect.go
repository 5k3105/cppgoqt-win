package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsOpacityEffect_chain uintptr
	addr_QGraphicsOpacityEffect_new uintptr
	addr_QGraphicsOpacityEffect_new2 uintptr
	addr_QGraphicsOpacityEffect_tr uintptr
	pg_QGraphicsOpacityEffect_opacity func(uintptr) float64
	addr_QGraphicsOpacityEffect_opacityMask uintptr
	pg_QGraphicsOpacityEffect_setOpacity func(uintptr, float64)
	addr_QGraphicsOpacityEffect_setOpacityMask uintptr
	pg_QGraphicsOpacityEffect_opacityChanged func(uintptr, float64)
	addr_QGraphicsOpacityEffect_opacityMaskChanged uintptr
	addr_QGraphicsOpacityEffect_updateBoundingRect uintptr
	addr_QGraphicsOpacityEffect_sourceIsPixmap uintptr
	addr_QGraphicsOpacityEffect_sourceBoundingRect uintptr
	addr_QGraphicsOpacityEffect_drawSource uintptr
	addr_QGraphicsOpacityEffect_sourcePixmap uintptr
	addr_QGraphicsOpacityEffect_connect uintptr
	addr_QGraphicsOpacityEffect_connect2 uintptr
	addr_QGraphicsOpacityEffect_sender uintptr
	addr_QGraphicsOpacityEffect_senderSignalIndex uintptr
	addr_QGraphicsOpacityEffect_receivers uintptr
	addr_QGraphicsOpacityEffect_isSignalConnected uintptr
	addr_QGraphicsOpacityEffect_tr2 uintptr
	addr_QGraphicsOpacityEffect_tr3 uintptr
	addr_QGraphicsOpacityEffect_sourceBoundingRect2 uintptr
	addr_QGraphicsOpacityEffect_sourcePixmap2 uintptr
	addr_QGraphicsOpacityEffect_sourcePixmap3 uintptr
	addr_QGraphicsOpacityEffect_sourcePixmap4 uintptr
	addr_QGraphicsOpacityEffect_connect3 uintptr
	addr_QGraphicsOpacityEffect_connect4 uintptr
	addr_QGraphicsOpacityEffect_metaObject_default uintptr
	addr_QGraphicsOpacityEffect_draw_default uintptr
	addr_QGraphicsOpacityEffect_boundingRectFor_default uintptr
	addr_QGraphicsOpacityEffect_sourceChanged_default uintptr
	addr_QGraphicsOpacityEffect_event_default uintptr
	addr_QGraphicsOpacityEffect_eventFilter_default uintptr
	addr_QGraphicsOpacityEffect_timerEvent_default uintptr
	addr_QGraphicsOpacityEffect_childEvent_default uintptr
	addr_QGraphicsOpacityEffect_customEvent_default uintptr
	addr_QGraphicsOpacityEffect_connectNotify_default uintptr
	addr_QGraphicsOpacityEffect_disconnectNotify_default uintptr
	pg_QGraphicsOpacityEffect_connect_opacityChanged func(uintptr, int64)
	pg_QGraphicsOpacityEffect_connect_opacityMaskChanged func(uintptr, int64)
	pg_QGraphicsOpacityEffect_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsOpacityEffect_connect_destroyed func(uintptr, int64)
	pg_QGraphicsOpacityEffect_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsOpacityEffect_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsOpacityEffect__disconnect func(int64)
	pg_QGraphicsOpacityEffect__set_override func(uintptr, int32)
	pg_QGraphicsOpacityEffect_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsOpacityEffect_chain, "QGraphicsOpacityEffect_chain")
	symaddr(&addr_QGraphicsOpacityEffect_new, "QGraphicsOpacityEffect_new")
	symaddr(&addr_QGraphicsOpacityEffect_new2, "QGraphicsOpacityEffect_new2")
	symaddr(&addr_QGraphicsOpacityEffect_tr, "QGraphicsOpacityEffect_tr")
	reg(&pg_QGraphicsOpacityEffect_opacity, "QGraphicsOpacityEffect_opacity")
	symaddr(&addr_QGraphicsOpacityEffect_opacityMask, "QGraphicsOpacityEffect_opacityMask")
	reg(&pg_QGraphicsOpacityEffect_setOpacity, "QGraphicsOpacityEffect_setOpacity")
	symaddr(&addr_QGraphicsOpacityEffect_setOpacityMask, "QGraphicsOpacityEffect_setOpacityMask")
	reg(&pg_QGraphicsOpacityEffect_opacityChanged, "QGraphicsOpacityEffect_opacityChanged")
	symaddr(&addr_QGraphicsOpacityEffect_opacityMaskChanged, "QGraphicsOpacityEffect_opacityMaskChanged")
	symaddr(&addr_QGraphicsOpacityEffect_updateBoundingRect, "QGraphicsOpacityEffect_updateBoundingRect")
	symaddr(&addr_QGraphicsOpacityEffect_sourceIsPixmap, "QGraphicsOpacityEffect_sourceIsPixmap")
	symaddr(&addr_QGraphicsOpacityEffect_sourceBoundingRect, "QGraphicsOpacityEffect_sourceBoundingRect")
	symaddr(&addr_QGraphicsOpacityEffect_drawSource, "QGraphicsOpacityEffect_drawSource")
	symaddr(&addr_QGraphicsOpacityEffect_sourcePixmap, "QGraphicsOpacityEffect_sourcePixmap")
	symaddr(&addr_QGraphicsOpacityEffect_connect, "QGraphicsOpacityEffect_connect")
	symaddr(&addr_QGraphicsOpacityEffect_connect2, "QGraphicsOpacityEffect_connect2")
	symaddr(&addr_QGraphicsOpacityEffect_sender, "QGraphicsOpacityEffect_sender")
	symaddr(&addr_QGraphicsOpacityEffect_senderSignalIndex, "QGraphicsOpacityEffect_senderSignalIndex")
	symaddr(&addr_QGraphicsOpacityEffect_receivers, "QGraphicsOpacityEffect_receivers")
	symaddr(&addr_QGraphicsOpacityEffect_isSignalConnected, "QGraphicsOpacityEffect_isSignalConnected")
	symaddr(&addr_QGraphicsOpacityEffect_tr2, "QGraphicsOpacityEffect_tr2")
	symaddr(&addr_QGraphicsOpacityEffect_tr3, "QGraphicsOpacityEffect_tr3")
	symaddr(&addr_QGraphicsOpacityEffect_sourceBoundingRect2, "QGraphicsOpacityEffect_sourceBoundingRect2")
	symaddr(&addr_QGraphicsOpacityEffect_sourcePixmap2, "QGraphicsOpacityEffect_sourcePixmap2")
	symaddr(&addr_QGraphicsOpacityEffect_sourcePixmap3, "QGraphicsOpacityEffect_sourcePixmap3")
	symaddr(&addr_QGraphicsOpacityEffect_sourcePixmap4, "QGraphicsOpacityEffect_sourcePixmap4")
	symaddr(&addr_QGraphicsOpacityEffect_connect3, "QGraphicsOpacityEffect_connect3")
	symaddr(&addr_QGraphicsOpacityEffect_connect4, "QGraphicsOpacityEffect_connect4")
	symaddr(&addr_QGraphicsOpacityEffect_metaObject_default, "QGraphicsOpacityEffect_metaObject_default")
	symaddr(&addr_QGraphicsOpacityEffect_draw_default, "QGraphicsOpacityEffect_draw_default")
	symaddr(&addr_QGraphicsOpacityEffect_boundingRectFor_default, "QGraphicsOpacityEffect_boundingRectFor_default")
	symaddr(&addr_QGraphicsOpacityEffect_sourceChanged_default, "QGraphicsOpacityEffect_sourceChanged_default")
	symaddr(&addr_QGraphicsOpacityEffect_event_default, "QGraphicsOpacityEffect_event_default")
	symaddr(&addr_QGraphicsOpacityEffect_eventFilter_default, "QGraphicsOpacityEffect_eventFilter_default")
	symaddr(&addr_QGraphicsOpacityEffect_timerEvent_default, "QGraphicsOpacityEffect_timerEvent_default")
	symaddr(&addr_QGraphicsOpacityEffect_childEvent_default, "QGraphicsOpacityEffect_childEvent_default")
	symaddr(&addr_QGraphicsOpacityEffect_customEvent_default, "QGraphicsOpacityEffect_customEvent_default")
	symaddr(&addr_QGraphicsOpacityEffect_connectNotify_default, "QGraphicsOpacityEffect_connectNotify_default")
	symaddr(&addr_QGraphicsOpacityEffect_disconnectNotify_default, "QGraphicsOpacityEffect_disconnectNotify_default")
	reg(&pg_QGraphicsOpacityEffect_connect_opacityChanged, "QGraphicsOpacityEffect_connect_opacityChanged")
	reg(&pg_QGraphicsOpacityEffect_connect_opacityMaskChanged, "QGraphicsOpacityEffect_connect_opacityMaskChanged")
	reg(&pg_QGraphicsOpacityEffect_connect_enabledChanged, "QGraphicsOpacityEffect_connect_enabledChanged")
	reg(&pg_QGraphicsOpacityEffect_connect_destroyed, "QGraphicsOpacityEffect_connect_destroyed")
	reg(&pg_QGraphicsOpacityEffect_connect_destroyed2, "QGraphicsOpacityEffect_connect_destroyed2")
	reg(&pg_QGraphicsOpacityEffect_connect_objectNameChanged, "QGraphicsOpacityEffect_connect_objectNameChanged")
	reg(&pg_QGraphicsOpacityEffect__disconnect, "QGraphicsOpacityEffect__disconnect")
	reg(&pg_QGraphicsOpacityEffect__set_override, "QGraphicsOpacityEffect__set_override")
	reg(&pg_QGraphicsOpacityEffect_delete, "QGraphicsOpacityEffect_delete")
}

type QGraphicsOpacityEffect struct {
	h uintptr
	*QGraphicsEffect
}

func bindQGraphicsOpacityEffect(h uintptr) *QGraphicsOpacityEffect {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsOpacityEffect_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsOpacityEffect{h: h, QGraphicsEffect: &QGraphicsEffect{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsOpacityEffect(h uintptr) *QGraphicsOpacityEffect { return bindQGraphicsOpacityEffect(h) }

func (this *QGraphicsOpacityEffect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsOpacityEffect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsOpacityEffect(h unsafe.Pointer) *QGraphicsOpacityEffect { return bindQGraphicsOpacityEffect(uintptr(h)) }
func UnsafeNewQGraphicsOpacityEffectOwned(h unsafe.Pointer) *QGraphicsOpacityEffect { return newownedQGraphicsOpacityEffect(uintptr(h)) }

func (this *QGraphicsOpacityEffect) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsEffect != nil {
		this.QGraphicsEffect.UnsafeRelease()
	}
}

func NewQGraphicsOpacityEffect() *QGraphicsOpacityEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_new, uintptr(__goid))
	__o := newownedQGraphicsOpacityEffect(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsOpacityEffect_virtIdx, QGraphicsOpacityEffect_enable)
	return __o
}

func NewQGraphicsOpacityEffect2(parent *QObject) *QGraphicsOpacityEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsOpacityEffect(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsOpacityEffect_virtIdx, QGraphicsOpacityEffect_enable)
	return __o
}

func QGraphicsOpacityEffect_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsOpacityEffect) Opacity() float64 {
	__r := pg_QGraphicsOpacityEffect_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsOpacityEffect) OpacityMask() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_opacityMask, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QGraphicsOpacityEffect) SetOpacity(opacity float64) {
	pg_QGraphicsOpacityEffect_setOpacity(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsOpacityEffect) SetOpacityMask(mask *QBrush) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_setOpacityMask, this.h, uintptr(mask.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mask)
}

func (this *QGraphicsOpacityEffect) OpacityChanged(opacity float64) {
	pg_QGraphicsOpacityEffect_opacityChanged(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsOpacityEffect) OpacityMaskChanged(mask *QBrush) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_opacityMaskChanged, this.h, uintptr(mask.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mask)
}

func (this *QGraphicsOpacityEffect) UpdateBoundingRect() {
	purego.SyscallN(addr_QGraphicsOpacityEffect_updateBoundingRect, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsOpacityEffect) SourceIsPixmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourceIsPixmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsOpacityEffect) SourceBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourceBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsOpacityEffect) DrawSource(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_drawSource, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsOpacityEffect) SourcePixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourcePixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QGraphicsOpacityEffect_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsOpacityEffect_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsOpacityEffect) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsOpacityEffect) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsOpacityEffect) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsOpacityEffect) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsOpacityEffect_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsOpacityEffect_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsOpacityEffect) SourceBoundingRectWithSystem(system CoordinateSystem) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourceBoundingRect2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsOpacityEffect) SourcePixmapWithSystem(system CoordinateSystem) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourcePixmap2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsOpacityEffect) SourcePixmap2(system CoordinateSystem, offset *QPoint) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourcePixmap3, this.h, uintptr(system), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func (this *QGraphicsOpacityEffect) SourcePixmap3(system CoordinateSystem, offset *QPoint, mode QGraphicsEffect__PixmapPadMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_sourcePixmap4, this.h, uintptr(system), uintptr(offset.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func QGraphicsOpacityEffect_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsOpacityEffect_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsOpacityEffect) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsOpacityEffect) DrawDefault(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_draw_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsOpacityEffect) BoundingRectForDefault(sourceRect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_boundingRectFor_default, this.h, uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceRect)
	return newownedQRectF(__r)
}

func (this *QGraphicsOpacityEffect) SourceChangedDefault(flags QGraphicsEffect__ChangeFlag) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_sourceChanged_default, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsOpacityEffect) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsOpacityEffect) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsOpacityEffect_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsOpacityEffect) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsOpacityEffect) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsOpacityEffect) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsOpacityEffect) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsOpacityEffect) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsOpacityEffect_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsOpacityEffect) OnOpacityChanged(fn func(float64)) int64 {
	id := registerCallback(func(opacity uintptr) { fn(math.Float64frombits(uint64(opacity))) })
	pg_QGraphicsOpacityEffect_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) OnOpacityMaskChanged(fn func(*QBrush)) int64 {
	id := registerCallback(func(mask uintptr) { fn(UnsafeNewQBrush(unsafe.Pointer(mask))) })
	pg_QGraphicsOpacityEffect_connect_opacityMaskChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QGraphicsOpacityEffect_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsOpacityEffect_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsOpacityEffect_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsOpacityEffect_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsOpacityEffect) Disconnect(handle int64) {
	pg_QGraphicsOpacityEffect__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsOpacityEffect_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -1044652152830004333:
		return 1
	case 5831744872130794138:
		return 2
	case -4520368087989481956:
		return 3
	case 1237236456312824991:
		return 4
	case 5217623514630858297:
		return 5
	case -5409927581812374498:
		return 6
	case -1185315367526857833:
		return 7
	case 3769868229123868474:
		return 8
	case 7379029982170503008:
		return 9
	case -568369772982215962:
		return 10
	}
	return -1
}

func QGraphicsOpacityEffect_enable(h uintptr, idx int32) { pg_QGraphicsOpacityEffect__set_override(h, idx) }

func (this *QGraphicsOpacityEffect) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsOpacityEffect) OnDraw(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -1044652152830004333, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "draw"
}

func (this *QGraphicsOpacityEffect) OnBoundingRectFor(fn func(*QRectF) *QRectF) {
	registerVirt(this.UnsafeGoID(), 5831744872130794138, func(sourceRect uintptr) uintptr { return uintptr(fn(UnsafeNewQRectF(unsafe.Pointer(sourceRect))).UnsafePointer()) }) // slot "boundingRectFor"
}

func (this *QGraphicsOpacityEffect) OnSourceChanged(fn func(QGraphicsEffect__ChangeFlag)) {
	registerVirt(this.UnsafeGoID(), -4520368087989481956, func(flags int32) { fn(QGraphicsEffect__ChangeFlag(flags)) }) // slot "sourceChanged"
}

func (this *QGraphicsOpacityEffect) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsOpacityEffect) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsOpacityEffect) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsOpacityEffect) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsOpacityEffect) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsOpacityEffect) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsOpacityEffect) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsOpacityEffect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsOpacityEffect_delete(this.h)
	}
	this.UnsafeRelease()
}
