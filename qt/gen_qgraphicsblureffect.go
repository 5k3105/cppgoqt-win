package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsBlurEffect_chain uintptr
	addr_QGraphicsBlurEffect_new uintptr
	addr_QGraphicsBlurEffect_new2 uintptr
	addr_QGraphicsBlurEffect_tr uintptr
	pg_QGraphicsBlurEffect_blurRadius func(uintptr) float64
	addr_QGraphicsBlurEffect_blurHints uintptr
	pg_QGraphicsBlurEffect_setBlurRadius func(uintptr, float64)
	addr_QGraphicsBlurEffect_setBlurHints uintptr
	pg_QGraphicsBlurEffect_blurRadiusChanged func(uintptr, float64)
	addr_QGraphicsBlurEffect_blurHintsChanged uintptr
	addr_QGraphicsBlurEffect_updateBoundingRect uintptr
	addr_QGraphicsBlurEffect_sourceIsPixmap uintptr
	addr_QGraphicsBlurEffect_sourceBoundingRect uintptr
	addr_QGraphicsBlurEffect_drawSource uintptr
	addr_QGraphicsBlurEffect_sourcePixmap uintptr
	addr_QGraphicsBlurEffect_connect uintptr
	addr_QGraphicsBlurEffect_connect2 uintptr
	addr_QGraphicsBlurEffect_sender uintptr
	addr_QGraphicsBlurEffect_senderSignalIndex uintptr
	addr_QGraphicsBlurEffect_receivers uintptr
	addr_QGraphicsBlurEffect_isSignalConnected uintptr
	addr_QGraphicsBlurEffect_tr2 uintptr
	addr_QGraphicsBlurEffect_tr3 uintptr
	addr_QGraphicsBlurEffect_sourceBoundingRect2 uintptr
	addr_QGraphicsBlurEffect_sourcePixmap2 uintptr
	addr_QGraphicsBlurEffect_sourcePixmap3 uintptr
	addr_QGraphicsBlurEffect_sourcePixmap4 uintptr
	addr_QGraphicsBlurEffect_connect3 uintptr
	addr_QGraphicsBlurEffect_connect4 uintptr
	addr_QGraphicsBlurEffect_metaObject_default uintptr
	addr_QGraphicsBlurEffect_boundingRectFor_default uintptr
	addr_QGraphicsBlurEffect_draw_default uintptr
	addr_QGraphicsBlurEffect_sourceChanged_default uintptr
	addr_QGraphicsBlurEffect_event_default uintptr
	addr_QGraphicsBlurEffect_eventFilter_default uintptr
	addr_QGraphicsBlurEffect_timerEvent_default uintptr
	addr_QGraphicsBlurEffect_childEvent_default uintptr
	addr_QGraphicsBlurEffect_customEvent_default uintptr
	addr_QGraphicsBlurEffect_connectNotify_default uintptr
	addr_QGraphicsBlurEffect_disconnectNotify_default uintptr
	pg_QGraphicsBlurEffect_connect_blurRadiusChanged func(uintptr, int64)
	pg_QGraphicsBlurEffect_connect_blurHintsChanged func(uintptr, int64)
	pg_QGraphicsBlurEffect_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsBlurEffect_connect_destroyed func(uintptr, int64)
	pg_QGraphicsBlurEffect_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsBlurEffect_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsBlurEffect__disconnect func(int64)
	pg_QGraphicsBlurEffect__set_override func(uintptr, int32)
	pg_QGraphicsBlurEffect_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsBlurEffect_chain, "QGraphicsBlurEffect_chain")
	symaddr(&addr_QGraphicsBlurEffect_new, "QGraphicsBlurEffect_new")
	symaddr(&addr_QGraphicsBlurEffect_new2, "QGraphicsBlurEffect_new2")
	symaddr(&addr_QGraphicsBlurEffect_tr, "QGraphicsBlurEffect_tr")
	reg(&pg_QGraphicsBlurEffect_blurRadius, "QGraphicsBlurEffect_blurRadius")
	symaddr(&addr_QGraphicsBlurEffect_blurHints, "QGraphicsBlurEffect_blurHints")
	reg(&pg_QGraphicsBlurEffect_setBlurRadius, "QGraphicsBlurEffect_setBlurRadius")
	symaddr(&addr_QGraphicsBlurEffect_setBlurHints, "QGraphicsBlurEffect_setBlurHints")
	reg(&pg_QGraphicsBlurEffect_blurRadiusChanged, "QGraphicsBlurEffect_blurRadiusChanged")
	symaddr(&addr_QGraphicsBlurEffect_blurHintsChanged, "QGraphicsBlurEffect_blurHintsChanged")
	symaddr(&addr_QGraphicsBlurEffect_updateBoundingRect, "QGraphicsBlurEffect_updateBoundingRect")
	symaddr(&addr_QGraphicsBlurEffect_sourceIsPixmap, "QGraphicsBlurEffect_sourceIsPixmap")
	symaddr(&addr_QGraphicsBlurEffect_sourceBoundingRect, "QGraphicsBlurEffect_sourceBoundingRect")
	symaddr(&addr_QGraphicsBlurEffect_drawSource, "QGraphicsBlurEffect_drawSource")
	symaddr(&addr_QGraphicsBlurEffect_sourcePixmap, "QGraphicsBlurEffect_sourcePixmap")
	symaddr(&addr_QGraphicsBlurEffect_connect, "QGraphicsBlurEffect_connect")
	symaddr(&addr_QGraphicsBlurEffect_connect2, "QGraphicsBlurEffect_connect2")
	symaddr(&addr_QGraphicsBlurEffect_sender, "QGraphicsBlurEffect_sender")
	symaddr(&addr_QGraphicsBlurEffect_senderSignalIndex, "QGraphicsBlurEffect_senderSignalIndex")
	symaddr(&addr_QGraphicsBlurEffect_receivers, "QGraphicsBlurEffect_receivers")
	symaddr(&addr_QGraphicsBlurEffect_isSignalConnected, "QGraphicsBlurEffect_isSignalConnected")
	symaddr(&addr_QGraphicsBlurEffect_tr2, "QGraphicsBlurEffect_tr2")
	symaddr(&addr_QGraphicsBlurEffect_tr3, "QGraphicsBlurEffect_tr3")
	symaddr(&addr_QGraphicsBlurEffect_sourceBoundingRect2, "QGraphicsBlurEffect_sourceBoundingRect2")
	symaddr(&addr_QGraphicsBlurEffect_sourcePixmap2, "QGraphicsBlurEffect_sourcePixmap2")
	symaddr(&addr_QGraphicsBlurEffect_sourcePixmap3, "QGraphicsBlurEffect_sourcePixmap3")
	symaddr(&addr_QGraphicsBlurEffect_sourcePixmap4, "QGraphicsBlurEffect_sourcePixmap4")
	symaddr(&addr_QGraphicsBlurEffect_connect3, "QGraphicsBlurEffect_connect3")
	symaddr(&addr_QGraphicsBlurEffect_connect4, "QGraphicsBlurEffect_connect4")
	symaddr(&addr_QGraphicsBlurEffect_metaObject_default, "QGraphicsBlurEffect_metaObject_default")
	symaddr(&addr_QGraphicsBlurEffect_boundingRectFor_default, "QGraphicsBlurEffect_boundingRectFor_default")
	symaddr(&addr_QGraphicsBlurEffect_draw_default, "QGraphicsBlurEffect_draw_default")
	symaddr(&addr_QGraphicsBlurEffect_sourceChanged_default, "QGraphicsBlurEffect_sourceChanged_default")
	symaddr(&addr_QGraphicsBlurEffect_event_default, "QGraphicsBlurEffect_event_default")
	symaddr(&addr_QGraphicsBlurEffect_eventFilter_default, "QGraphicsBlurEffect_eventFilter_default")
	symaddr(&addr_QGraphicsBlurEffect_timerEvent_default, "QGraphicsBlurEffect_timerEvent_default")
	symaddr(&addr_QGraphicsBlurEffect_childEvent_default, "QGraphicsBlurEffect_childEvent_default")
	symaddr(&addr_QGraphicsBlurEffect_customEvent_default, "QGraphicsBlurEffect_customEvent_default")
	symaddr(&addr_QGraphicsBlurEffect_connectNotify_default, "QGraphicsBlurEffect_connectNotify_default")
	symaddr(&addr_QGraphicsBlurEffect_disconnectNotify_default, "QGraphicsBlurEffect_disconnectNotify_default")
	reg(&pg_QGraphicsBlurEffect_connect_blurRadiusChanged, "QGraphicsBlurEffect_connect_blurRadiusChanged")
	reg(&pg_QGraphicsBlurEffect_connect_blurHintsChanged, "QGraphicsBlurEffect_connect_blurHintsChanged")
	reg(&pg_QGraphicsBlurEffect_connect_enabledChanged, "QGraphicsBlurEffect_connect_enabledChanged")
	reg(&pg_QGraphicsBlurEffect_connect_destroyed, "QGraphicsBlurEffect_connect_destroyed")
	reg(&pg_QGraphicsBlurEffect_connect_destroyed2, "QGraphicsBlurEffect_connect_destroyed2")
	reg(&pg_QGraphicsBlurEffect_connect_objectNameChanged, "QGraphicsBlurEffect_connect_objectNameChanged")
	reg(&pg_QGraphicsBlurEffect__disconnect, "QGraphicsBlurEffect__disconnect")
	reg(&pg_QGraphicsBlurEffect__set_override, "QGraphicsBlurEffect__set_override")
	reg(&pg_QGraphicsBlurEffect_delete, "QGraphicsBlurEffect_delete")
}

type QGraphicsBlurEffect struct {
	h uintptr
	*QGraphicsEffect
}

func bindQGraphicsBlurEffect(h uintptr) *QGraphicsBlurEffect {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsBlurEffect_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsBlurEffect{h: h, QGraphicsEffect: &QGraphicsEffect{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsBlurEffect(h uintptr) *QGraphicsBlurEffect { return bindQGraphicsBlurEffect(h) }

func (this *QGraphicsBlurEffect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsBlurEffect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsBlurEffect(h unsafe.Pointer) *QGraphicsBlurEffect { return bindQGraphicsBlurEffect(uintptr(h)) }
func UnsafeNewQGraphicsBlurEffectOwned(h unsafe.Pointer) *QGraphicsBlurEffect { return newownedQGraphicsBlurEffect(uintptr(h)) }

func (this *QGraphicsBlurEffect) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsEffect != nil {
		this.QGraphicsEffect.UnsafeRelease()
	}
}

func NewQGraphicsBlurEffect() *QGraphicsBlurEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_new, uintptr(__goid))
	__o := newownedQGraphicsBlurEffect(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsBlurEffect_virtIdx, QGraphicsBlurEffect_enable)
	return __o
}

func NewQGraphicsBlurEffect2(parent *QObject) *QGraphicsBlurEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsBlurEffect(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsBlurEffect_virtIdx, QGraphicsBlurEffect_enable)
	return __o
}

func QGraphicsBlurEffect_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsBlurEffect) BlurRadius() float64 {
	__r := pg_QGraphicsBlurEffect_blurRadius(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsBlurEffect) BlurHints() QGraphicsBlurEffect__BlurHint {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_blurHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsBlurEffect__BlurHint(__r)
}

func (this *QGraphicsBlurEffect) SetBlurRadius(blurRadius float64) {
	pg_QGraphicsBlurEffect_setBlurRadius(this.h, blurRadius)
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) SetBlurHints(hints QGraphicsBlurEffect__BlurHint) {
	purego.SyscallN(addr_QGraphicsBlurEffect_setBlurHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) BlurRadiusChanged(blurRadius float64) {
	pg_QGraphicsBlurEffect_blurRadiusChanged(this.h, blurRadius)
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) BlurHintsChanged(hints QGraphicsBlurEffect__BlurHint) {
	purego.SyscallN(addr_QGraphicsBlurEffect_blurHintsChanged, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) UpdateBoundingRect() {
	purego.SyscallN(addr_QGraphicsBlurEffect_updateBoundingRect, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) SourceIsPixmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourceIsPixmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsBlurEffect) SourceBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourceBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsBlurEffect) DrawSource(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsBlurEffect_drawSource, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsBlurEffect) SourcePixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourcePixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QGraphicsBlurEffect_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsBlurEffect_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsBlurEffect) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsBlurEffect) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsBlurEffect) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsBlurEffect) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsBlurEffect_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsBlurEffect_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsBlurEffect) SourceBoundingRectWithSystem(system CoordinateSystem) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourceBoundingRect2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsBlurEffect) SourcePixmapWithSystem(system CoordinateSystem) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourcePixmap2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsBlurEffect) SourcePixmap2(system CoordinateSystem, offset *QPoint) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourcePixmap3, this.h, uintptr(system), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func (this *QGraphicsBlurEffect) SourcePixmap3(system CoordinateSystem, offset *QPoint, mode QGraphicsEffect__PixmapPadMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_sourcePixmap4, this.h, uintptr(system), uintptr(offset.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func QGraphicsBlurEffect_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsBlurEffect_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsBlurEffect) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsBlurEffect) BoundingRectForDefault(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_boundingRectFor_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsBlurEffect) DrawDefault(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsBlurEffect_draw_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsBlurEffect) SourceChangedDefault(flags QGraphicsEffect__ChangeFlag) {
	purego.SyscallN(addr_QGraphicsBlurEffect_sourceChanged_default, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsBlurEffect) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsBlurEffect) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsBlurEffect_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsBlurEffect) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsBlurEffect_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsBlurEffect) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsBlurEffect_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsBlurEffect) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsBlurEffect_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsBlurEffect) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsBlurEffect_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsBlurEffect) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsBlurEffect_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsBlurEffect) OnBlurRadiusChanged(fn func(float64)) int64 {
	id := registerCallback(func(blurRadius uintptr) { fn(math.Float64frombits(uint64(blurRadius))) })
	pg_QGraphicsBlurEffect_connect_blurRadiusChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) OnBlurHintsChanged(fn func(QGraphicsBlurEffect__BlurHint)) int64 {
	id := registerCallback(func(hints int32) { fn(QGraphicsBlurEffect__BlurHint(hints)) })
	pg_QGraphicsBlurEffect_connect_blurHintsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QGraphicsBlurEffect_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsBlurEffect_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsBlurEffect_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsBlurEffect_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsBlurEffect) Disconnect(handle int64) {
	pg_QGraphicsBlurEffect__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsBlurEffect_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 5831744872130794138:
		return 1
	case -1044652152830004333:
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

func QGraphicsBlurEffect_enable(h uintptr, idx int32) { pg_QGraphicsBlurEffect__set_override(h, idx) }

func (this *QGraphicsBlurEffect) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsBlurEffect) OnBoundingRectFor(fn func(*QRectF) *QRectF) {
	registerVirt(this.UnsafeGoID(), 5831744872130794138, func(rect uintptr) uintptr { return uintptr(fn(UnsafeNewQRectF(unsafe.Pointer(rect))).UnsafePointer()) }) // slot "boundingRectFor"
}

func (this *QGraphicsBlurEffect) OnDraw(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -1044652152830004333, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "draw"
}

func (this *QGraphicsBlurEffect) OnSourceChanged(fn func(QGraphicsEffect__ChangeFlag)) {
	registerVirt(this.UnsafeGoID(), -4520368087989481956, func(flags int32) { fn(QGraphicsEffect__ChangeFlag(flags)) }) // slot "sourceChanged"
}

func (this *QGraphicsBlurEffect) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsBlurEffect) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsBlurEffect) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsBlurEffect) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsBlurEffect) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsBlurEffect) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsBlurEffect) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsBlurEffect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsBlurEffect_delete(this.h)
	}
	this.UnsafeRelease()
}
