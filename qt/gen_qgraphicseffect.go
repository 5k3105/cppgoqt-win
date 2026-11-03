package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsEffect_base func(uintptr) uintptr
	addr_QGraphicsEffect_tr uintptr
	addr_QGraphicsEffect_boundingRectFor uintptr
	addr_QGraphicsEffect_boundingRect uintptr
	addr_QGraphicsEffect_isEnabled uintptr
	addr_QGraphicsEffect_setEnabled uintptr
	addr_QGraphicsEffect_update uintptr
	addr_QGraphicsEffect_enabledChanged uintptr
	addr_QGraphicsEffect_updateBoundingRect uintptr
	addr_QGraphicsEffect_sourceIsPixmap uintptr
	addr_QGraphicsEffect_sourceBoundingRect uintptr
	addr_QGraphicsEffect_drawSource uintptr
	addr_QGraphicsEffect_sourcePixmap uintptr
	addr_QGraphicsEffect_connect uintptr
	addr_QGraphicsEffect_connect2 uintptr
	addr_QGraphicsEffect_sender uintptr
	addr_QGraphicsEffect_senderSignalIndex uintptr
	addr_QGraphicsEffect_receivers uintptr
	addr_QGraphicsEffect_isSignalConnected uintptr
	addr_QGraphicsEffect_tr2 uintptr
	addr_QGraphicsEffect_tr3 uintptr
	addr_QGraphicsEffect_sourceBoundingRect2 uintptr
	addr_QGraphicsEffect_sourcePixmap2 uintptr
	addr_QGraphicsEffect_sourcePixmap3 uintptr
	addr_QGraphicsEffect_sourcePixmap4 uintptr
	addr_QGraphicsEffect_connect3 uintptr
	addr_QGraphicsEffect_connect4 uintptr
	addr_QGraphicsEffect_boundingRectFor_default uintptr
	addr_QGraphicsEffect_metaObject_default uintptr
	addr_QGraphicsEffect_sourceChanged_default uintptr
	addr_QGraphicsEffect_event_default uintptr
	addr_QGraphicsEffect_eventFilter_default uintptr
	addr_QGraphicsEffect_timerEvent_default uintptr
	addr_QGraphicsEffect_childEvent_default uintptr
	addr_QGraphicsEffect_customEvent_default uintptr
	addr_QGraphicsEffect_connectNotify_default uintptr
	addr_QGraphicsEffect_disconnectNotify_default uintptr
	pg_QGraphicsEffect_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsEffect_connect_destroyed func(uintptr, int64)
	pg_QGraphicsEffect_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsEffect_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsEffect__disconnect func(int64)
	pg_QGraphicsEffect__set_override func(uintptr, int32)
	pg_QGraphicsEffect__Go_new func(int64, uintptr) uintptr
	pg_QGraphicsEffect_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsEffect_base, "QGraphicsEffect_base")
	symaddr(&addr_QGraphicsEffect_tr, "QGraphicsEffect_tr")
	symaddr(&addr_QGraphicsEffect_boundingRectFor, "QGraphicsEffect_boundingRectFor")
	symaddr(&addr_QGraphicsEffect_boundingRect, "QGraphicsEffect_boundingRect")
	symaddr(&addr_QGraphicsEffect_isEnabled, "QGraphicsEffect_isEnabled")
	symaddr(&addr_QGraphicsEffect_setEnabled, "QGraphicsEffect_setEnabled")
	symaddr(&addr_QGraphicsEffect_update, "QGraphicsEffect_update")
	symaddr(&addr_QGraphicsEffect_enabledChanged, "QGraphicsEffect_enabledChanged")
	symaddr(&addr_QGraphicsEffect_updateBoundingRect, "QGraphicsEffect_updateBoundingRect")
	symaddr(&addr_QGraphicsEffect_sourceIsPixmap, "QGraphicsEffect_sourceIsPixmap")
	symaddr(&addr_QGraphicsEffect_sourceBoundingRect, "QGraphicsEffect_sourceBoundingRect")
	symaddr(&addr_QGraphicsEffect_drawSource, "QGraphicsEffect_drawSource")
	symaddr(&addr_QGraphicsEffect_sourcePixmap, "QGraphicsEffect_sourcePixmap")
	symaddr(&addr_QGraphicsEffect_connect, "QGraphicsEffect_connect")
	symaddr(&addr_QGraphicsEffect_connect2, "QGraphicsEffect_connect2")
	symaddr(&addr_QGraphicsEffect_sender, "QGraphicsEffect_sender")
	symaddr(&addr_QGraphicsEffect_senderSignalIndex, "QGraphicsEffect_senderSignalIndex")
	symaddr(&addr_QGraphicsEffect_receivers, "QGraphicsEffect_receivers")
	symaddr(&addr_QGraphicsEffect_isSignalConnected, "QGraphicsEffect_isSignalConnected")
	symaddr(&addr_QGraphicsEffect_tr2, "QGraphicsEffect_tr2")
	symaddr(&addr_QGraphicsEffect_tr3, "QGraphicsEffect_tr3")
	symaddr(&addr_QGraphicsEffect_sourceBoundingRect2, "QGraphicsEffect_sourceBoundingRect2")
	symaddr(&addr_QGraphicsEffect_sourcePixmap2, "QGraphicsEffect_sourcePixmap2")
	symaddr(&addr_QGraphicsEffect_sourcePixmap3, "QGraphicsEffect_sourcePixmap3")
	symaddr(&addr_QGraphicsEffect_sourcePixmap4, "QGraphicsEffect_sourcePixmap4")
	symaddr(&addr_QGraphicsEffect_connect3, "QGraphicsEffect_connect3")
	symaddr(&addr_QGraphicsEffect_connect4, "QGraphicsEffect_connect4")
	symaddr(&addr_QGraphicsEffect_boundingRectFor_default, "QGraphicsEffect_boundingRectFor_default")
	symaddr(&addr_QGraphicsEffect_metaObject_default, "QGraphicsEffect_metaObject_default")
	symaddr(&addr_QGraphicsEffect_sourceChanged_default, "QGraphicsEffect_sourceChanged_default")
	symaddr(&addr_QGraphicsEffect_event_default, "QGraphicsEffect_event_default")
	symaddr(&addr_QGraphicsEffect_eventFilter_default, "QGraphicsEffect_eventFilter_default")
	symaddr(&addr_QGraphicsEffect_timerEvent_default, "QGraphicsEffect_timerEvent_default")
	symaddr(&addr_QGraphicsEffect_childEvent_default, "QGraphicsEffect_childEvent_default")
	symaddr(&addr_QGraphicsEffect_customEvent_default, "QGraphicsEffect_customEvent_default")
	symaddr(&addr_QGraphicsEffect_connectNotify_default, "QGraphicsEffect_connectNotify_default")
	symaddr(&addr_QGraphicsEffect_disconnectNotify_default, "QGraphicsEffect_disconnectNotify_default")
	reg(&pg_QGraphicsEffect_connect_enabledChanged, "QGraphicsEffect_connect_enabledChanged")
	reg(&pg_QGraphicsEffect_connect_destroyed, "QGraphicsEffect_connect_destroyed")
	reg(&pg_QGraphicsEffect_connect_destroyed2, "QGraphicsEffect_connect_destroyed2")
	reg(&pg_QGraphicsEffect_connect_objectNameChanged, "QGraphicsEffect_connect_objectNameChanged")
	reg(&pg_QGraphicsEffect__disconnect, "QGraphicsEffect__disconnect")
	reg(&pg_QGraphicsEffect__set_override, "QGraphicsEffect__set_override")
	reg(&pg_QGraphicsEffect__Go_new, "QGraphicsEffect__Go_new")
	reg(&pg_QGraphicsEffect_delete, "QGraphicsEffect_delete")
}

type QGraphicsEffect struct {
	h uintptr
	*QObject
}

func bindQGraphicsEffect(h uintptr) *QGraphicsEffect {
	if h == 0 {
		return nil
	}
	return &QGraphicsEffect{h: h, QObject: bindQObject(pg_QGraphicsEffect_base(h))}
}

func newownedQGraphicsEffect(h uintptr) *QGraphicsEffect { return bindQGraphicsEffect(h) }

func (this *QGraphicsEffect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsEffect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsEffect(h unsafe.Pointer) *QGraphicsEffect { return bindQGraphicsEffect(uintptr(h)) }
func UnsafeNewQGraphicsEffectOwned(h unsafe.Pointer) *QGraphicsEffect { return newownedQGraphicsEffect(uintptr(h)) }

func (this *QGraphicsEffect) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QGraphicsEffect_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsEffect) BoundingRectFor(sourceRect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_boundingRectFor, this.h, uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceRect)
	return newownedQRectF(__r)
}

func (this *QGraphicsEffect) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsEffect) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsEffect) SetEnabled(enable bool) {
	purego.SyscallN(addr_QGraphicsEffect_setEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEffect) Update() {
	purego.SyscallN(addr_QGraphicsEffect_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEffect) EnabledChanged(enabled bool) {
	purego.SyscallN(addr_QGraphicsEffect_enabledChanged, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEffect) UpdateBoundingRect() {
	purego.SyscallN(addr_QGraphicsEffect_updateBoundingRect, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEffect) SourceIsPixmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourceIsPixmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsEffect) SourceBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourceBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsEffect) DrawSource(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsEffect_drawSource, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsEffect) SourcePixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourcePixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QGraphicsEffect_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsEffect_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsEffect) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsEffect) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsEffect) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsEffect) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsEffect_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsEffect_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsEffect) SourceBoundingRectWithSystem(system CoordinateSystem) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourceBoundingRect2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsEffect) SourcePixmapWithSystem(system CoordinateSystem) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourcePixmap2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsEffect) SourcePixmap2(system CoordinateSystem, offset *QPoint) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourcePixmap3, this.h, uintptr(system), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func (this *QGraphicsEffect) SourcePixmap3(system CoordinateSystem, offset *QPoint, mode QGraphicsEffect__PixmapPadMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_sourcePixmap4, this.h, uintptr(system), uintptr(offset.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func QGraphicsEffect_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsEffect_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsEffect) BoundingRectForDefault(sourceRect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_boundingRectFor_default, this.h, uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceRect)
	return newownedQRectF(__r)
}

func (this *QGraphicsEffect) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEffect_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsEffect) SourceChangedDefault(flags QGraphicsEffect__ChangeFlag) {
	purego.SyscallN(addr_QGraphicsEffect_sourceChanged_default, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEffect) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsEffect) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEffect_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsEffect) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsEffect_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEffect) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsEffect_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEffect) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsEffect_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEffect) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsEffect_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsEffect) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsEffect_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsEffect) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QGraphicsEffect_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsEffect) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsEffect_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsEffect) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsEffect_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsEffect) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsEffect_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsEffect) Disconnect(handle int64) {
	pg_QGraphicsEffect__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsEffect_virtIdx(slot int64) int32 {
	switch slot {
	case -1044652152830004333:
		return 0
	case 1607084412591819503:
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

func QGraphicsEffect_enable(h uintptr, idx int32) { pg_QGraphicsEffect__set_override(h, idx) }

func NewQGraphicsEffect() *QGraphicsEffect {
	id := nextVirtID()
	o := newownedQGraphicsEffect(pg_QGraphicsEffect__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsEffect_virtIdx, QGraphicsEffect_enable)
	return o
}

func NewQGraphicsEffect2(parent *QObject) *QGraphicsEffect {
	id := nextVirtID()
	o := newownedQGraphicsEffect(pg_QGraphicsEffect__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsEffect_virtIdx, QGraphicsEffect_enable)
	return o
}

func (this *QGraphicsEffect) OnDraw(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -1044652152830004333, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "draw"
}

func (this *QGraphicsEffect) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsEffect) OnBoundingRectFor(fn func(*QRectF) *QRectF) {
	registerVirt(this.UnsafeGoID(), 5831744872130794138, func(sourceRect uintptr) uintptr { return uintptr(fn(UnsafeNewQRectF(unsafe.Pointer(sourceRect))).UnsafePointer()) }) // slot "boundingRectFor"
}

func (this *QGraphicsEffect) OnSourceChanged(fn func(QGraphicsEffect__ChangeFlag)) {
	registerVirt(this.UnsafeGoID(), -4520368087989481956, func(flags int32) { fn(QGraphicsEffect__ChangeFlag(flags)) }) // slot "sourceChanged"
}

func (this *QGraphicsEffect) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsEffect) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsEffect) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsEffect) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsEffect) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsEffect) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsEffect) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsEffect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsEffect_delete(this.h)
	}
	this.UnsafeRelease()
}
