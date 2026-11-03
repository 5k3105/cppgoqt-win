package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsColorizeEffect_chain uintptr
	addr_QGraphicsColorizeEffect_new uintptr
	addr_QGraphicsColorizeEffect_new2 uintptr
	addr_QGraphicsColorizeEffect_tr uintptr
	addr_QGraphicsColorizeEffect_color uintptr
	pg_QGraphicsColorizeEffect_strength func(uintptr) float64
	addr_QGraphicsColorizeEffect_setColor uintptr
	pg_QGraphicsColorizeEffect_setStrength func(uintptr, float64)
	addr_QGraphicsColorizeEffect_colorChanged uintptr
	pg_QGraphicsColorizeEffect_strengthChanged func(uintptr, float64)
	addr_QGraphicsColorizeEffect_updateBoundingRect uintptr
	addr_QGraphicsColorizeEffect_sourceIsPixmap uintptr
	addr_QGraphicsColorizeEffect_sourceBoundingRect uintptr
	addr_QGraphicsColorizeEffect_drawSource uintptr
	addr_QGraphicsColorizeEffect_sourcePixmap uintptr
	addr_QGraphicsColorizeEffect_connect uintptr
	addr_QGraphicsColorizeEffect_connect2 uintptr
	addr_QGraphicsColorizeEffect_sender uintptr
	addr_QGraphicsColorizeEffect_senderSignalIndex uintptr
	addr_QGraphicsColorizeEffect_receivers uintptr
	addr_QGraphicsColorizeEffect_isSignalConnected uintptr
	addr_QGraphicsColorizeEffect_tr2 uintptr
	addr_QGraphicsColorizeEffect_tr3 uintptr
	addr_QGraphicsColorizeEffect_sourceBoundingRect2 uintptr
	addr_QGraphicsColorizeEffect_sourcePixmap2 uintptr
	addr_QGraphicsColorizeEffect_sourcePixmap3 uintptr
	addr_QGraphicsColorizeEffect_sourcePixmap4 uintptr
	addr_QGraphicsColorizeEffect_connect3 uintptr
	addr_QGraphicsColorizeEffect_connect4 uintptr
	addr_QGraphicsColorizeEffect_metaObject_default uintptr
	addr_QGraphicsColorizeEffect_draw_default uintptr
	addr_QGraphicsColorizeEffect_boundingRectFor_default uintptr
	addr_QGraphicsColorizeEffect_sourceChanged_default uintptr
	addr_QGraphicsColorizeEffect_event_default uintptr
	addr_QGraphicsColorizeEffect_eventFilter_default uintptr
	addr_QGraphicsColorizeEffect_timerEvent_default uintptr
	addr_QGraphicsColorizeEffect_childEvent_default uintptr
	addr_QGraphicsColorizeEffect_customEvent_default uintptr
	addr_QGraphicsColorizeEffect_connectNotify_default uintptr
	addr_QGraphicsColorizeEffect_disconnectNotify_default uintptr
	pg_QGraphicsColorizeEffect_connect_colorChanged func(uintptr, int64)
	pg_QGraphicsColorizeEffect_connect_strengthChanged func(uintptr, int64)
	pg_QGraphicsColorizeEffect_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsColorizeEffect_connect_destroyed func(uintptr, int64)
	pg_QGraphicsColorizeEffect_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsColorizeEffect_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsColorizeEffect__disconnect func(int64)
	pg_QGraphicsColorizeEffect__set_override func(uintptr, int32)
	pg_QGraphicsColorizeEffect_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsColorizeEffect_chain, "QGraphicsColorizeEffect_chain")
	symaddr(&addr_QGraphicsColorizeEffect_new, "QGraphicsColorizeEffect_new")
	symaddr(&addr_QGraphicsColorizeEffect_new2, "QGraphicsColorizeEffect_new2")
	symaddr(&addr_QGraphicsColorizeEffect_tr, "QGraphicsColorizeEffect_tr")
	symaddr(&addr_QGraphicsColorizeEffect_color, "QGraphicsColorizeEffect_color")
	reg(&pg_QGraphicsColorizeEffect_strength, "QGraphicsColorizeEffect_strength")
	symaddr(&addr_QGraphicsColorizeEffect_setColor, "QGraphicsColorizeEffect_setColor")
	reg(&pg_QGraphicsColorizeEffect_setStrength, "QGraphicsColorizeEffect_setStrength")
	symaddr(&addr_QGraphicsColorizeEffect_colorChanged, "QGraphicsColorizeEffect_colorChanged")
	reg(&pg_QGraphicsColorizeEffect_strengthChanged, "QGraphicsColorizeEffect_strengthChanged")
	symaddr(&addr_QGraphicsColorizeEffect_updateBoundingRect, "QGraphicsColorizeEffect_updateBoundingRect")
	symaddr(&addr_QGraphicsColorizeEffect_sourceIsPixmap, "QGraphicsColorizeEffect_sourceIsPixmap")
	symaddr(&addr_QGraphicsColorizeEffect_sourceBoundingRect, "QGraphicsColorizeEffect_sourceBoundingRect")
	symaddr(&addr_QGraphicsColorizeEffect_drawSource, "QGraphicsColorizeEffect_drawSource")
	symaddr(&addr_QGraphicsColorizeEffect_sourcePixmap, "QGraphicsColorizeEffect_sourcePixmap")
	symaddr(&addr_QGraphicsColorizeEffect_connect, "QGraphicsColorizeEffect_connect")
	symaddr(&addr_QGraphicsColorizeEffect_connect2, "QGraphicsColorizeEffect_connect2")
	symaddr(&addr_QGraphicsColorizeEffect_sender, "QGraphicsColorizeEffect_sender")
	symaddr(&addr_QGraphicsColorizeEffect_senderSignalIndex, "QGraphicsColorizeEffect_senderSignalIndex")
	symaddr(&addr_QGraphicsColorizeEffect_receivers, "QGraphicsColorizeEffect_receivers")
	symaddr(&addr_QGraphicsColorizeEffect_isSignalConnected, "QGraphicsColorizeEffect_isSignalConnected")
	symaddr(&addr_QGraphicsColorizeEffect_tr2, "QGraphicsColorizeEffect_tr2")
	symaddr(&addr_QGraphicsColorizeEffect_tr3, "QGraphicsColorizeEffect_tr3")
	symaddr(&addr_QGraphicsColorizeEffect_sourceBoundingRect2, "QGraphicsColorizeEffect_sourceBoundingRect2")
	symaddr(&addr_QGraphicsColorizeEffect_sourcePixmap2, "QGraphicsColorizeEffect_sourcePixmap2")
	symaddr(&addr_QGraphicsColorizeEffect_sourcePixmap3, "QGraphicsColorizeEffect_sourcePixmap3")
	symaddr(&addr_QGraphicsColorizeEffect_sourcePixmap4, "QGraphicsColorizeEffect_sourcePixmap4")
	symaddr(&addr_QGraphicsColorizeEffect_connect3, "QGraphicsColorizeEffect_connect3")
	symaddr(&addr_QGraphicsColorizeEffect_connect4, "QGraphicsColorizeEffect_connect4")
	symaddr(&addr_QGraphicsColorizeEffect_metaObject_default, "QGraphicsColorizeEffect_metaObject_default")
	symaddr(&addr_QGraphicsColorizeEffect_draw_default, "QGraphicsColorizeEffect_draw_default")
	symaddr(&addr_QGraphicsColorizeEffect_boundingRectFor_default, "QGraphicsColorizeEffect_boundingRectFor_default")
	symaddr(&addr_QGraphicsColorizeEffect_sourceChanged_default, "QGraphicsColorizeEffect_sourceChanged_default")
	symaddr(&addr_QGraphicsColorizeEffect_event_default, "QGraphicsColorizeEffect_event_default")
	symaddr(&addr_QGraphicsColorizeEffect_eventFilter_default, "QGraphicsColorizeEffect_eventFilter_default")
	symaddr(&addr_QGraphicsColorizeEffect_timerEvent_default, "QGraphicsColorizeEffect_timerEvent_default")
	symaddr(&addr_QGraphicsColorizeEffect_childEvent_default, "QGraphicsColorizeEffect_childEvent_default")
	symaddr(&addr_QGraphicsColorizeEffect_customEvent_default, "QGraphicsColorizeEffect_customEvent_default")
	symaddr(&addr_QGraphicsColorizeEffect_connectNotify_default, "QGraphicsColorizeEffect_connectNotify_default")
	symaddr(&addr_QGraphicsColorizeEffect_disconnectNotify_default, "QGraphicsColorizeEffect_disconnectNotify_default")
	reg(&pg_QGraphicsColorizeEffect_connect_colorChanged, "QGraphicsColorizeEffect_connect_colorChanged")
	reg(&pg_QGraphicsColorizeEffect_connect_strengthChanged, "QGraphicsColorizeEffect_connect_strengthChanged")
	reg(&pg_QGraphicsColorizeEffect_connect_enabledChanged, "QGraphicsColorizeEffect_connect_enabledChanged")
	reg(&pg_QGraphicsColorizeEffect_connect_destroyed, "QGraphicsColorizeEffect_connect_destroyed")
	reg(&pg_QGraphicsColorizeEffect_connect_destroyed2, "QGraphicsColorizeEffect_connect_destroyed2")
	reg(&pg_QGraphicsColorizeEffect_connect_objectNameChanged, "QGraphicsColorizeEffect_connect_objectNameChanged")
	reg(&pg_QGraphicsColorizeEffect__disconnect, "QGraphicsColorizeEffect__disconnect")
	reg(&pg_QGraphicsColorizeEffect__set_override, "QGraphicsColorizeEffect__set_override")
	reg(&pg_QGraphicsColorizeEffect_delete, "QGraphicsColorizeEffect_delete")
}

type QGraphicsColorizeEffect struct {
	h uintptr
	*QGraphicsEffect
}

func bindQGraphicsColorizeEffect(h uintptr) *QGraphicsColorizeEffect {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsColorizeEffect_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsColorizeEffect{h: h, QGraphicsEffect: &QGraphicsEffect{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsColorizeEffect(h uintptr) *QGraphicsColorizeEffect { return bindQGraphicsColorizeEffect(h) }

func (this *QGraphicsColorizeEffect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsColorizeEffect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsColorizeEffect(h unsafe.Pointer) *QGraphicsColorizeEffect { return bindQGraphicsColorizeEffect(uintptr(h)) }
func UnsafeNewQGraphicsColorizeEffectOwned(h unsafe.Pointer) *QGraphicsColorizeEffect { return newownedQGraphicsColorizeEffect(uintptr(h)) }

func (this *QGraphicsColorizeEffect) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsEffect != nil {
		this.QGraphicsEffect.UnsafeRelease()
	}
}

func NewQGraphicsColorizeEffect() *QGraphicsColorizeEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_new, uintptr(__goid))
	__o := newownedQGraphicsColorizeEffect(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsColorizeEffect_virtIdx, QGraphicsColorizeEffect_enable)
	return __o
}

func NewQGraphicsColorizeEffect2(parent *QObject) *QGraphicsColorizeEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsColorizeEffect(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsColorizeEffect_virtIdx, QGraphicsColorizeEffect_enable)
	return __o
}

func QGraphicsColorizeEffect_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsColorizeEffect) Color() *QColor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_color, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QGraphicsColorizeEffect) Strength() float64 {
	__r := pg_QGraphicsColorizeEffect_strength(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsColorizeEffect) SetColor(c *QColor) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_setColor, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QGraphicsColorizeEffect) SetStrength(strength float64) {
	pg_QGraphicsColorizeEffect_setStrength(this.h, strength)
	runtime.KeepAlive(this)
}

func (this *QGraphicsColorizeEffect) ColorChanged(color *QColor) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_colorChanged, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QGraphicsColorizeEffect) StrengthChanged(strength float64) {
	pg_QGraphicsColorizeEffect_strengthChanged(this.h, strength)
	runtime.KeepAlive(this)
}

func (this *QGraphicsColorizeEffect) UpdateBoundingRect() {
	purego.SyscallN(addr_QGraphicsColorizeEffect_updateBoundingRect, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsColorizeEffect) SourceIsPixmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourceIsPixmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsColorizeEffect) SourceBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourceBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsColorizeEffect) DrawSource(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_drawSource, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsColorizeEffect) SourcePixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourcePixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QGraphicsColorizeEffect_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsColorizeEffect_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsColorizeEffect) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsColorizeEffect) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsColorizeEffect) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsColorizeEffect) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsColorizeEffect_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsColorizeEffect_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsColorizeEffect) SourceBoundingRectWithSystem(system CoordinateSystem) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourceBoundingRect2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsColorizeEffect) SourcePixmapWithSystem(system CoordinateSystem) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourcePixmap2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsColorizeEffect) SourcePixmap2(system CoordinateSystem, offset *QPoint) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourcePixmap3, this.h, uintptr(system), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func (this *QGraphicsColorizeEffect) SourcePixmap3(system CoordinateSystem, offset *QPoint, mode QGraphicsEffect__PixmapPadMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_sourcePixmap4, this.h, uintptr(system), uintptr(offset.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func QGraphicsColorizeEffect_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsColorizeEffect_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsColorizeEffect) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsColorizeEffect) DrawDefault(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_draw_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsColorizeEffect) BoundingRectForDefault(sourceRect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_boundingRectFor_default, this.h, uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceRect)
	return newownedQRectF(__r)
}

func (this *QGraphicsColorizeEffect) SourceChangedDefault(flags QGraphicsEffect__ChangeFlag) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_sourceChanged_default, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsColorizeEffect) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsColorizeEffect) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsColorizeEffect_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsColorizeEffect) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsColorizeEffect) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsColorizeEffect) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsColorizeEffect) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsColorizeEffect) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsColorizeEffect_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsColorizeEffect) OnColorChanged(fn func(*QColor)) int64 {
	id := registerCallback(func(color uintptr) { fn(UnsafeNewQColor(unsafe.Pointer(color))) })
	pg_QGraphicsColorizeEffect_connect_colorChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) OnStrengthChanged(fn func(float64)) int64 {
	id := registerCallback(func(strength uintptr) { fn(math.Float64frombits(uint64(strength))) })
	pg_QGraphicsColorizeEffect_connect_strengthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QGraphicsColorizeEffect_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsColorizeEffect_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsColorizeEffect_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsColorizeEffect_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsColorizeEffect) Disconnect(handle int64) {
	pg_QGraphicsColorizeEffect__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsColorizeEffect_virtIdx(slot int64) int32 {
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

func QGraphicsColorizeEffect_enable(h uintptr, idx int32) { pg_QGraphicsColorizeEffect__set_override(h, idx) }

func (this *QGraphicsColorizeEffect) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsColorizeEffect) OnDraw(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -1044652152830004333, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "draw"
}

func (this *QGraphicsColorizeEffect) OnBoundingRectFor(fn func(*QRectF) *QRectF) {
	registerVirt(this.UnsafeGoID(), 5831744872130794138, func(sourceRect uintptr) uintptr { return uintptr(fn(UnsafeNewQRectF(unsafe.Pointer(sourceRect))).UnsafePointer()) }) // slot "boundingRectFor"
}

func (this *QGraphicsColorizeEffect) OnSourceChanged(fn func(QGraphicsEffect__ChangeFlag)) {
	registerVirt(this.UnsafeGoID(), -4520368087989481956, func(flags int32) { fn(QGraphicsEffect__ChangeFlag(flags)) }) // slot "sourceChanged"
}

func (this *QGraphicsColorizeEffect) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsColorizeEffect) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsColorizeEffect) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsColorizeEffect) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsColorizeEffect) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsColorizeEffect) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsColorizeEffect) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsColorizeEffect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsColorizeEffect_delete(this.h)
	}
	this.UnsafeRelease()
}
