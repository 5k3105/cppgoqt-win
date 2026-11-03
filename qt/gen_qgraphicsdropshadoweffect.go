package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsDropShadowEffect_chain uintptr
	addr_QGraphicsDropShadowEffect_new uintptr
	addr_QGraphicsDropShadowEffect_new2 uintptr
	addr_QGraphicsDropShadowEffect_tr uintptr
	addr_QGraphicsDropShadowEffect_offset uintptr
	pg_QGraphicsDropShadowEffect_xOffset func(uintptr) float64
	pg_QGraphicsDropShadowEffect_yOffset func(uintptr) float64
	pg_QGraphicsDropShadowEffect_blurRadius func(uintptr) float64
	addr_QGraphicsDropShadowEffect_color uintptr
	addr_QGraphicsDropShadowEffect_setOffset uintptr
	pg_QGraphicsDropShadowEffect_setOffset2 func(uintptr, float64, float64)
	pg_QGraphicsDropShadowEffect_setOffset3 func(uintptr, float64)
	pg_QGraphicsDropShadowEffect_setXOffset func(uintptr, float64)
	pg_QGraphicsDropShadowEffect_setYOffset func(uintptr, float64)
	pg_QGraphicsDropShadowEffect_setBlurRadius func(uintptr, float64)
	addr_QGraphicsDropShadowEffect_setColor uintptr
	addr_QGraphicsDropShadowEffect_offsetChanged uintptr
	pg_QGraphicsDropShadowEffect_blurRadiusChanged func(uintptr, float64)
	addr_QGraphicsDropShadowEffect_colorChanged uintptr
	addr_QGraphicsDropShadowEffect_updateBoundingRect uintptr
	addr_QGraphicsDropShadowEffect_sourceIsPixmap uintptr
	addr_QGraphicsDropShadowEffect_sourceBoundingRect uintptr
	addr_QGraphicsDropShadowEffect_drawSource uintptr
	addr_QGraphicsDropShadowEffect_sourcePixmap uintptr
	addr_QGraphicsDropShadowEffect_connect uintptr
	addr_QGraphicsDropShadowEffect_connect2 uintptr
	addr_QGraphicsDropShadowEffect_sender uintptr
	addr_QGraphicsDropShadowEffect_senderSignalIndex uintptr
	addr_QGraphicsDropShadowEffect_receivers uintptr
	addr_QGraphicsDropShadowEffect_isSignalConnected uintptr
	addr_QGraphicsDropShadowEffect_tr2 uintptr
	addr_QGraphicsDropShadowEffect_tr3 uintptr
	addr_QGraphicsDropShadowEffect_sourceBoundingRect2 uintptr
	addr_QGraphicsDropShadowEffect_sourcePixmap2 uintptr
	addr_QGraphicsDropShadowEffect_sourcePixmap3 uintptr
	addr_QGraphicsDropShadowEffect_sourcePixmap4 uintptr
	addr_QGraphicsDropShadowEffect_connect3 uintptr
	addr_QGraphicsDropShadowEffect_connect4 uintptr
	addr_QGraphicsDropShadowEffect_metaObject_default uintptr
	addr_QGraphicsDropShadowEffect_boundingRectFor_default uintptr
	addr_QGraphicsDropShadowEffect_draw_default uintptr
	addr_QGraphicsDropShadowEffect_sourceChanged_default uintptr
	addr_QGraphicsDropShadowEffect_event_default uintptr
	addr_QGraphicsDropShadowEffect_eventFilter_default uintptr
	addr_QGraphicsDropShadowEffect_timerEvent_default uintptr
	addr_QGraphicsDropShadowEffect_childEvent_default uintptr
	addr_QGraphicsDropShadowEffect_customEvent_default uintptr
	addr_QGraphicsDropShadowEffect_connectNotify_default uintptr
	addr_QGraphicsDropShadowEffect_disconnectNotify_default uintptr
	pg_QGraphicsDropShadowEffect_connect_offsetChanged func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_blurRadiusChanged func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_colorChanged func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_destroyed func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsDropShadowEffect_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsDropShadowEffect__disconnect func(int64)
	pg_QGraphicsDropShadowEffect__set_override func(uintptr, int32)
	pg_QGraphicsDropShadowEffect_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsDropShadowEffect_chain, "QGraphicsDropShadowEffect_chain")
	symaddr(&addr_QGraphicsDropShadowEffect_new, "QGraphicsDropShadowEffect_new")
	symaddr(&addr_QGraphicsDropShadowEffect_new2, "QGraphicsDropShadowEffect_new2")
	symaddr(&addr_QGraphicsDropShadowEffect_tr, "QGraphicsDropShadowEffect_tr")
	symaddr(&addr_QGraphicsDropShadowEffect_offset, "QGraphicsDropShadowEffect_offset")
	reg(&pg_QGraphicsDropShadowEffect_xOffset, "QGraphicsDropShadowEffect_xOffset")
	reg(&pg_QGraphicsDropShadowEffect_yOffset, "QGraphicsDropShadowEffect_yOffset")
	reg(&pg_QGraphicsDropShadowEffect_blurRadius, "QGraphicsDropShadowEffect_blurRadius")
	symaddr(&addr_QGraphicsDropShadowEffect_color, "QGraphicsDropShadowEffect_color")
	symaddr(&addr_QGraphicsDropShadowEffect_setOffset, "QGraphicsDropShadowEffect_setOffset")
	reg(&pg_QGraphicsDropShadowEffect_setOffset2, "QGraphicsDropShadowEffect_setOffset2")
	reg(&pg_QGraphicsDropShadowEffect_setOffset3, "QGraphicsDropShadowEffect_setOffset3")
	reg(&pg_QGraphicsDropShadowEffect_setXOffset, "QGraphicsDropShadowEffect_setXOffset")
	reg(&pg_QGraphicsDropShadowEffect_setYOffset, "QGraphicsDropShadowEffect_setYOffset")
	reg(&pg_QGraphicsDropShadowEffect_setBlurRadius, "QGraphicsDropShadowEffect_setBlurRadius")
	symaddr(&addr_QGraphicsDropShadowEffect_setColor, "QGraphicsDropShadowEffect_setColor")
	symaddr(&addr_QGraphicsDropShadowEffect_offsetChanged, "QGraphicsDropShadowEffect_offsetChanged")
	reg(&pg_QGraphicsDropShadowEffect_blurRadiusChanged, "QGraphicsDropShadowEffect_blurRadiusChanged")
	symaddr(&addr_QGraphicsDropShadowEffect_colorChanged, "QGraphicsDropShadowEffect_colorChanged")
	symaddr(&addr_QGraphicsDropShadowEffect_updateBoundingRect, "QGraphicsDropShadowEffect_updateBoundingRect")
	symaddr(&addr_QGraphicsDropShadowEffect_sourceIsPixmap, "QGraphicsDropShadowEffect_sourceIsPixmap")
	symaddr(&addr_QGraphicsDropShadowEffect_sourceBoundingRect, "QGraphicsDropShadowEffect_sourceBoundingRect")
	symaddr(&addr_QGraphicsDropShadowEffect_drawSource, "QGraphicsDropShadowEffect_drawSource")
	symaddr(&addr_QGraphicsDropShadowEffect_sourcePixmap, "QGraphicsDropShadowEffect_sourcePixmap")
	symaddr(&addr_QGraphicsDropShadowEffect_connect, "QGraphicsDropShadowEffect_connect")
	symaddr(&addr_QGraphicsDropShadowEffect_connect2, "QGraphicsDropShadowEffect_connect2")
	symaddr(&addr_QGraphicsDropShadowEffect_sender, "QGraphicsDropShadowEffect_sender")
	symaddr(&addr_QGraphicsDropShadowEffect_senderSignalIndex, "QGraphicsDropShadowEffect_senderSignalIndex")
	symaddr(&addr_QGraphicsDropShadowEffect_receivers, "QGraphicsDropShadowEffect_receivers")
	symaddr(&addr_QGraphicsDropShadowEffect_isSignalConnected, "QGraphicsDropShadowEffect_isSignalConnected")
	symaddr(&addr_QGraphicsDropShadowEffect_tr2, "QGraphicsDropShadowEffect_tr2")
	symaddr(&addr_QGraphicsDropShadowEffect_tr3, "QGraphicsDropShadowEffect_tr3")
	symaddr(&addr_QGraphicsDropShadowEffect_sourceBoundingRect2, "QGraphicsDropShadowEffect_sourceBoundingRect2")
	symaddr(&addr_QGraphicsDropShadowEffect_sourcePixmap2, "QGraphicsDropShadowEffect_sourcePixmap2")
	symaddr(&addr_QGraphicsDropShadowEffect_sourcePixmap3, "QGraphicsDropShadowEffect_sourcePixmap3")
	symaddr(&addr_QGraphicsDropShadowEffect_sourcePixmap4, "QGraphicsDropShadowEffect_sourcePixmap4")
	symaddr(&addr_QGraphicsDropShadowEffect_connect3, "QGraphicsDropShadowEffect_connect3")
	symaddr(&addr_QGraphicsDropShadowEffect_connect4, "QGraphicsDropShadowEffect_connect4")
	symaddr(&addr_QGraphicsDropShadowEffect_metaObject_default, "QGraphicsDropShadowEffect_metaObject_default")
	symaddr(&addr_QGraphicsDropShadowEffect_boundingRectFor_default, "QGraphicsDropShadowEffect_boundingRectFor_default")
	symaddr(&addr_QGraphicsDropShadowEffect_draw_default, "QGraphicsDropShadowEffect_draw_default")
	symaddr(&addr_QGraphicsDropShadowEffect_sourceChanged_default, "QGraphicsDropShadowEffect_sourceChanged_default")
	symaddr(&addr_QGraphicsDropShadowEffect_event_default, "QGraphicsDropShadowEffect_event_default")
	symaddr(&addr_QGraphicsDropShadowEffect_eventFilter_default, "QGraphicsDropShadowEffect_eventFilter_default")
	symaddr(&addr_QGraphicsDropShadowEffect_timerEvent_default, "QGraphicsDropShadowEffect_timerEvent_default")
	symaddr(&addr_QGraphicsDropShadowEffect_childEvent_default, "QGraphicsDropShadowEffect_childEvent_default")
	symaddr(&addr_QGraphicsDropShadowEffect_customEvent_default, "QGraphicsDropShadowEffect_customEvent_default")
	symaddr(&addr_QGraphicsDropShadowEffect_connectNotify_default, "QGraphicsDropShadowEffect_connectNotify_default")
	symaddr(&addr_QGraphicsDropShadowEffect_disconnectNotify_default, "QGraphicsDropShadowEffect_disconnectNotify_default")
	reg(&pg_QGraphicsDropShadowEffect_connect_offsetChanged, "QGraphicsDropShadowEffect_connect_offsetChanged")
	reg(&pg_QGraphicsDropShadowEffect_connect_blurRadiusChanged, "QGraphicsDropShadowEffect_connect_blurRadiusChanged")
	reg(&pg_QGraphicsDropShadowEffect_connect_colorChanged, "QGraphicsDropShadowEffect_connect_colorChanged")
	reg(&pg_QGraphicsDropShadowEffect_connect_enabledChanged, "QGraphicsDropShadowEffect_connect_enabledChanged")
	reg(&pg_QGraphicsDropShadowEffect_connect_destroyed, "QGraphicsDropShadowEffect_connect_destroyed")
	reg(&pg_QGraphicsDropShadowEffect_connect_destroyed2, "QGraphicsDropShadowEffect_connect_destroyed2")
	reg(&pg_QGraphicsDropShadowEffect_connect_objectNameChanged, "QGraphicsDropShadowEffect_connect_objectNameChanged")
	reg(&pg_QGraphicsDropShadowEffect__disconnect, "QGraphicsDropShadowEffect__disconnect")
	reg(&pg_QGraphicsDropShadowEffect__set_override, "QGraphicsDropShadowEffect__set_override")
	reg(&pg_QGraphicsDropShadowEffect_delete, "QGraphicsDropShadowEffect_delete")
}

type QGraphicsDropShadowEffect struct {
	h uintptr
	*QGraphicsEffect
}

func bindQGraphicsDropShadowEffect(h uintptr) *QGraphicsDropShadowEffect {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsDropShadowEffect_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsDropShadowEffect{h: h, QGraphicsEffect: &QGraphicsEffect{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsDropShadowEffect(h uintptr) *QGraphicsDropShadowEffect { return bindQGraphicsDropShadowEffect(h) }

func (this *QGraphicsDropShadowEffect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsDropShadowEffect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsDropShadowEffect(h unsafe.Pointer) *QGraphicsDropShadowEffect { return bindQGraphicsDropShadowEffect(uintptr(h)) }
func UnsafeNewQGraphicsDropShadowEffectOwned(h unsafe.Pointer) *QGraphicsDropShadowEffect { return newownedQGraphicsDropShadowEffect(uintptr(h)) }

func (this *QGraphicsDropShadowEffect) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsEffect != nil {
		this.QGraphicsEffect.UnsafeRelease()
	}
}

func NewQGraphicsDropShadowEffect() *QGraphicsDropShadowEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_new, uintptr(__goid))
	__o := newownedQGraphicsDropShadowEffect(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsDropShadowEffect_virtIdx, QGraphicsDropShadowEffect_enable)
	return __o
}

func NewQGraphicsDropShadowEffect2(parent *QObject) *QGraphicsDropShadowEffect {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsDropShadowEffect(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsDropShadowEffect_virtIdx, QGraphicsDropShadowEffect_enable)
	return __o
}

func QGraphicsDropShadowEffect_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsDropShadowEffect) Offset() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_offset, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsDropShadowEffect) XOffset() float64 {
	__r := pg_QGraphicsDropShadowEffect_xOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsDropShadowEffect) YOffset() float64 {
	__r := pg_QGraphicsDropShadowEffect_yOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsDropShadowEffect) BlurRadius() float64 {
	__r := pg_QGraphicsDropShadowEffect_blurRadius(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsDropShadowEffect) Color() *QColor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_color, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QGraphicsDropShadowEffect) SetOffset(ofs *QPointF) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_setOffset, this.h, uintptr(ofs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ofs)
}

func (this *QGraphicsDropShadowEffect) SetOffset2(dx float64, dy float64) {
	pg_QGraphicsDropShadowEffect_setOffset2(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SetOffsetWith(d float64) {
	pg_QGraphicsDropShadowEffect_setOffset3(this.h, d)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SetXOffset(dx float64) {
	pg_QGraphicsDropShadowEffect_setXOffset(this.h, dx)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SetYOffset(dy float64) {
	pg_QGraphicsDropShadowEffect_setYOffset(this.h, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SetBlurRadius(blurRadius float64) {
	pg_QGraphicsDropShadowEffect_setBlurRadius(this.h, blurRadius)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SetColor(color *QColor) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_setColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QGraphicsDropShadowEffect) OffsetChanged(offset *QPointF) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_offsetChanged, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QGraphicsDropShadowEffect) BlurRadiusChanged(blurRadius float64) {
	pg_QGraphicsDropShadowEffect_blurRadiusChanged(this.h, blurRadius)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) ColorChanged(color *QColor) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_colorChanged, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QGraphicsDropShadowEffect) UpdateBoundingRect() {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_updateBoundingRect, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) SourceIsPixmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourceIsPixmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsDropShadowEffect) SourceBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourceBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsDropShadowEffect) DrawSource(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_drawSource, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsDropShadowEffect) SourcePixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourcePixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QGraphicsDropShadowEffect_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsDropShadowEffect_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsDropShadowEffect) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsDropShadowEffect) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsDropShadowEffect) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsDropShadowEffect) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsDropShadowEffect_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsDropShadowEffect_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsDropShadowEffect) SourceBoundingRectWithSystem(system CoordinateSystem) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourceBoundingRect2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsDropShadowEffect) SourcePixmapWithSystem(system CoordinateSystem) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourcePixmap2, this.h, uintptr(system))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsDropShadowEffect) SourcePixmap2(system CoordinateSystem, offset *QPoint) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourcePixmap3, this.h, uintptr(system), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func (this *QGraphicsDropShadowEffect) SourcePixmap3(system CoordinateSystem, offset *QPoint, mode QGraphicsEffect__PixmapPadMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_sourcePixmap4, this.h, uintptr(system), uintptr(offset.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPixmap(__r)
}

func QGraphicsDropShadowEffect_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsDropShadowEffect_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsDropShadowEffect) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsDropShadowEffect) BoundingRectForDefault(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_boundingRectFor_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsDropShadowEffect) DrawDefault(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_draw_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsDropShadowEffect) SourceChangedDefault(flags QGraphicsEffect__ChangeFlag) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_sourceChanged_default, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsDropShadowEffect) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsDropShadowEffect) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsDropShadowEffect_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsDropShadowEffect) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsDropShadowEffect) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsDropShadowEffect) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsDropShadowEffect) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsDropShadowEffect) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsDropShadowEffect_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsDropShadowEffect) OnOffsetChanged(fn func(*QPointF)) int64 {
	id := registerCallback(func(offset uintptr) { fn(UnsafeNewQPointF(unsafe.Pointer(offset))) })
	pg_QGraphicsDropShadowEffect_connect_offsetChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnBlurRadiusChanged(fn func(float64)) int64 {
	id := registerCallback(func(blurRadius uintptr) { fn(math.Float64frombits(uint64(blurRadius))) })
	pg_QGraphicsDropShadowEffect_connect_blurRadiusChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnColorChanged(fn func(*QColor)) int64 {
	id := registerCallback(func(color uintptr) { fn(UnsafeNewQColor(unsafe.Pointer(color))) })
	pg_QGraphicsDropShadowEffect_connect_colorChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QGraphicsDropShadowEffect_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsDropShadowEffect_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsDropShadowEffect_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsDropShadowEffect_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsDropShadowEffect) Disconnect(handle int64) {
	pg_QGraphicsDropShadowEffect__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsDropShadowEffect_virtIdx(slot int64) int32 {
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

func QGraphicsDropShadowEffect_enable(h uintptr, idx int32) { pg_QGraphicsDropShadowEffect__set_override(h, idx) }

func (this *QGraphicsDropShadowEffect) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsDropShadowEffect) OnBoundingRectFor(fn func(*QRectF) *QRectF) {
	registerVirt(this.UnsafeGoID(), 5831744872130794138, func(rect uintptr) uintptr { return uintptr(fn(UnsafeNewQRectF(unsafe.Pointer(rect))).UnsafePointer()) }) // slot "boundingRectFor"
}

func (this *QGraphicsDropShadowEffect) OnDraw(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -1044652152830004333, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "draw"
}

func (this *QGraphicsDropShadowEffect) OnSourceChanged(fn func(QGraphicsEffect__ChangeFlag)) {
	registerVirt(this.UnsafeGoID(), -4520368087989481956, func(flags int32) { fn(QGraphicsEffect__ChangeFlag(flags)) }) // slot "sourceChanged"
}

func (this *QGraphicsDropShadowEffect) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsDropShadowEffect) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsDropShadowEffect) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsDropShadowEffect) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsDropShadowEffect) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsDropShadowEffect) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsDropShadowEffect) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsDropShadowEffect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsDropShadowEffect_delete(this.h)
	}
	this.UnsafeRelease()
}
