package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProxyStyle_chain uintptr
	addr_QProxyStyle_new uintptr
	addr_QProxyStyle_new2 uintptr
	addr_QProxyStyle_new3 uintptr
	addr_QProxyStyle_tr uintptr
	addr_QProxyStyle_baseStyle uintptr
	addr_QProxyStyle_setBaseStyle uintptr
	addr_QProxyStyle_visualRect uintptr
	addr_QProxyStyle_visualPos uintptr
	addr_QProxyStyle_sliderPositionFromValue uintptr
	addr_QProxyStyle_sliderValueFromPosition uintptr
	addr_QProxyStyle_visualAlignment uintptr
	addr_QProxyStyle_alignedRect uintptr
	addr_QProxyStyle_connect uintptr
	addr_QProxyStyle_connect2 uintptr
	addr_QProxyStyle_sender uintptr
	addr_QProxyStyle_senderSignalIndex uintptr
	addr_QProxyStyle_receivers uintptr
	addr_QProxyStyle_isSignalConnected uintptr
	addr_QProxyStyle_tr2 uintptr
	addr_QProxyStyle_tr3 uintptr
	addr_QProxyStyle_sliderPositionFromValue2 uintptr
	addr_QProxyStyle_sliderValueFromPosition2 uintptr
	addr_QProxyStyle_connect3 uintptr
	addr_QProxyStyle_connect4 uintptr
	addr_QProxyStyle_metaObject_default uintptr
	addr_QProxyStyle_drawPrimitive_default uintptr
	addr_QProxyStyle_drawControl_default uintptr
	addr_QProxyStyle_drawComplexControl_default uintptr
	addr_QProxyStyle_drawItemText_default uintptr
	addr_QProxyStyle_drawItemPixmap_default uintptr
	addr_QProxyStyle_sizeFromContents_default uintptr
	addr_QProxyStyle_subElementRect_default uintptr
	addr_QProxyStyle_subControlRect_default uintptr
	addr_QProxyStyle_itemTextRect_default uintptr
	addr_QProxyStyle_itemPixmapRect_default uintptr
	addr_QProxyStyle_hitTestComplexControl_default uintptr
	addr_QProxyStyle_styleHint_default uintptr
	addr_QProxyStyle_pixelMetric_default uintptr
	addr_QProxyStyle_layoutSpacing_default uintptr
	addr_QProxyStyle_standardIcon_default uintptr
	addr_QProxyStyle_standardPixmap_default uintptr
	addr_QProxyStyle_generatedIconPixmap_default uintptr
	addr_QProxyStyle_standardPalette_default uintptr
	addr_QProxyStyle_polish_default uintptr
	addr_QProxyStyle_unpolish_default uintptr
	addr_QProxyStyle_event_default uintptr
	addr_QProxyStyle_eventFilter_default uintptr
	addr_QProxyStyle_timerEvent_default uintptr
	addr_QProxyStyle_childEvent_default uintptr
	addr_QProxyStyle_customEvent_default uintptr
	addr_QProxyStyle_connectNotify_default uintptr
	addr_QProxyStyle_disconnectNotify_default uintptr
	pg_QProxyStyle_connect_destroyed func(uintptr, int64)
	pg_QProxyStyle_connect_destroyed2 func(uintptr, int64)
	pg_QProxyStyle_connect_objectNameChanged func(uintptr, int64)
	pg_QProxyStyle__disconnect func(int64)
	pg_QProxyStyle__set_override func(uintptr, int32)
	pg_QProxyStyle_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProxyStyle_chain, "QProxyStyle_chain")
	symaddr(&addr_QProxyStyle_new, "QProxyStyle_new")
	symaddr(&addr_QProxyStyle_new2, "QProxyStyle_new2")
	symaddr(&addr_QProxyStyle_new3, "QProxyStyle_new3")
	symaddr(&addr_QProxyStyle_tr, "QProxyStyle_tr")
	symaddr(&addr_QProxyStyle_baseStyle, "QProxyStyle_baseStyle")
	symaddr(&addr_QProxyStyle_setBaseStyle, "QProxyStyle_setBaseStyle")
	symaddr(&addr_QProxyStyle_visualRect, "QProxyStyle_visualRect")
	symaddr(&addr_QProxyStyle_visualPos, "QProxyStyle_visualPos")
	symaddr(&addr_QProxyStyle_sliderPositionFromValue, "QProxyStyle_sliderPositionFromValue")
	symaddr(&addr_QProxyStyle_sliderValueFromPosition, "QProxyStyle_sliderValueFromPosition")
	symaddr(&addr_QProxyStyle_visualAlignment, "QProxyStyle_visualAlignment")
	symaddr(&addr_QProxyStyle_alignedRect, "QProxyStyle_alignedRect")
	symaddr(&addr_QProxyStyle_connect, "QProxyStyle_connect")
	symaddr(&addr_QProxyStyle_connect2, "QProxyStyle_connect2")
	symaddr(&addr_QProxyStyle_sender, "QProxyStyle_sender")
	symaddr(&addr_QProxyStyle_senderSignalIndex, "QProxyStyle_senderSignalIndex")
	symaddr(&addr_QProxyStyle_receivers, "QProxyStyle_receivers")
	symaddr(&addr_QProxyStyle_isSignalConnected, "QProxyStyle_isSignalConnected")
	symaddr(&addr_QProxyStyle_tr2, "QProxyStyle_tr2")
	symaddr(&addr_QProxyStyle_tr3, "QProxyStyle_tr3")
	symaddr(&addr_QProxyStyle_sliderPositionFromValue2, "QProxyStyle_sliderPositionFromValue2")
	symaddr(&addr_QProxyStyle_sliderValueFromPosition2, "QProxyStyle_sliderValueFromPosition2")
	symaddr(&addr_QProxyStyle_connect3, "QProxyStyle_connect3")
	symaddr(&addr_QProxyStyle_connect4, "QProxyStyle_connect4")
	symaddr(&addr_QProxyStyle_metaObject_default, "QProxyStyle_metaObject_default")
	symaddr(&addr_QProxyStyle_drawPrimitive_default, "QProxyStyle_drawPrimitive_default")
	symaddr(&addr_QProxyStyle_drawControl_default, "QProxyStyle_drawControl_default")
	symaddr(&addr_QProxyStyle_drawComplexControl_default, "QProxyStyle_drawComplexControl_default")
	symaddr(&addr_QProxyStyle_drawItemText_default, "QProxyStyle_drawItemText_default")
	symaddr(&addr_QProxyStyle_drawItemPixmap_default, "QProxyStyle_drawItemPixmap_default")
	symaddr(&addr_QProxyStyle_sizeFromContents_default, "QProxyStyle_sizeFromContents_default")
	symaddr(&addr_QProxyStyle_subElementRect_default, "QProxyStyle_subElementRect_default")
	symaddr(&addr_QProxyStyle_subControlRect_default, "QProxyStyle_subControlRect_default")
	symaddr(&addr_QProxyStyle_itemTextRect_default, "QProxyStyle_itemTextRect_default")
	symaddr(&addr_QProxyStyle_itemPixmapRect_default, "QProxyStyle_itemPixmapRect_default")
	symaddr(&addr_QProxyStyle_hitTestComplexControl_default, "QProxyStyle_hitTestComplexControl_default")
	symaddr(&addr_QProxyStyle_styleHint_default, "QProxyStyle_styleHint_default")
	symaddr(&addr_QProxyStyle_pixelMetric_default, "QProxyStyle_pixelMetric_default")
	symaddr(&addr_QProxyStyle_layoutSpacing_default, "QProxyStyle_layoutSpacing_default")
	symaddr(&addr_QProxyStyle_standardIcon_default, "QProxyStyle_standardIcon_default")
	symaddr(&addr_QProxyStyle_standardPixmap_default, "QProxyStyle_standardPixmap_default")
	symaddr(&addr_QProxyStyle_generatedIconPixmap_default, "QProxyStyle_generatedIconPixmap_default")
	symaddr(&addr_QProxyStyle_standardPalette_default, "QProxyStyle_standardPalette_default")
	symaddr(&addr_QProxyStyle_polish_default, "QProxyStyle_polish_default")
	symaddr(&addr_QProxyStyle_unpolish_default, "QProxyStyle_unpolish_default")
	symaddr(&addr_QProxyStyle_event_default, "QProxyStyle_event_default")
	symaddr(&addr_QProxyStyle_eventFilter_default, "QProxyStyle_eventFilter_default")
	symaddr(&addr_QProxyStyle_timerEvent_default, "QProxyStyle_timerEvent_default")
	symaddr(&addr_QProxyStyle_childEvent_default, "QProxyStyle_childEvent_default")
	symaddr(&addr_QProxyStyle_customEvent_default, "QProxyStyle_customEvent_default")
	symaddr(&addr_QProxyStyle_connectNotify_default, "QProxyStyle_connectNotify_default")
	symaddr(&addr_QProxyStyle_disconnectNotify_default, "QProxyStyle_disconnectNotify_default")
	reg(&pg_QProxyStyle_connect_destroyed, "QProxyStyle_connect_destroyed")
	reg(&pg_QProxyStyle_connect_destroyed2, "QProxyStyle_connect_destroyed2")
	reg(&pg_QProxyStyle_connect_objectNameChanged, "QProxyStyle_connect_objectNameChanged")
	reg(&pg_QProxyStyle__disconnect, "QProxyStyle__disconnect")
	reg(&pg_QProxyStyle__set_override, "QProxyStyle__set_override")
	reg(&pg_QProxyStyle_delete, "QProxyStyle_delete")
}

type QProxyStyle struct {
	h uintptr
	*QCommonStyle
}

func bindQProxyStyle(h uintptr) *QProxyStyle {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QProxyStyle_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QProxyStyle{h: h, QCommonStyle: &QCommonStyle{h: __c[0], QStyle: &QStyle{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQProxyStyle(h uintptr) *QProxyStyle { return bindQProxyStyle(h) }

func (this *QProxyStyle) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProxyStyle) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProxyStyle(h unsafe.Pointer) *QProxyStyle { return bindQProxyStyle(uintptr(h)) }
func UnsafeNewQProxyStyleOwned(h unsafe.Pointer) *QProxyStyle { return newownedQProxyStyle(uintptr(h)) }

func (this *QProxyStyle) UnsafeRelease() {
	this.h = 0
	if this.QCommonStyle != nil {
		this.QCommonStyle.UnsafeRelease()
	}
}

func NewQProxyStyle() *QProxyStyle {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_new, uintptr(__goid))
	__o := newownedQProxyStyle(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProxyStyle_virtIdx, QProxyStyle_enable)
	return __o
}

func NewQProxyStyle2(key string) *QProxyStyle {
	__goid := nextVirtID()
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_new2, uintptr(__goid), key_p, uintptr(len(key)))
	__o := newownedQProxyStyle(__r)
	runtime.KeepAlive(key)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProxyStyle_virtIdx, QProxyStyle_enable)
	return __o
}

func NewQProxyStyle3(style *QStyle) *QProxyStyle {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_new3, uintptr(__goid), uintptr(style.UnsafePointer()))
	__o := newownedQProxyStyle(__r)
	runtime.KeepAlive(style)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProxyStyle_virtIdx, QProxyStyle_enable)
	return __o
}

func QProxyStyle_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QProxyStyle) BaseStyle() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_baseStyle, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func (this *QProxyStyle) SetBaseStyle(style *QStyle) {
	purego.SyscallN(addr_QProxyStyle_setBaseStyle, this.h, uintptr(style.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(style)
}

func QProxyStyle_VisualRect(direction LayoutDirection, boundingRect *QRect, logicalRect *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_visualRect, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalRect.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalRect)
	return newownedQRect(__r)
}

func QProxyStyle_VisualPos(direction LayoutDirection, boundingRect *QRect, logicalPos *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_visualPos, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalPos.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalPos)
	return newownedQPoint(__r)
}

func QProxyStyle_SliderPositionFromValue(min int, max int, val int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_sliderPositionFromValue, uintptr(min), uintptr(max), uintptr(val), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QProxyStyle_SliderValueFromPosition(min int, max int, pos int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_sliderValueFromPosition, uintptr(min), uintptr(max), uintptr(pos), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QProxyStyle_VisualAlignment(direction LayoutDirection, alignment AlignmentFlag) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_visualAlignment, uintptr(direction), uintptr(alignment))
	__r := int32(uint32(__r0))
	return AlignmentFlag(__r)
}

func QProxyStyle_AlignedRect(direction LayoutDirection, alignment AlignmentFlag, size *QSize, rectangle *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_alignedRect, uintptr(direction), uintptr(alignment), uintptr(size.UnsafePointer()), uintptr(rectangle.UnsafePointer()))
	runtime.KeepAlive(size)
	runtime.KeepAlive(rectangle)
	return newownedQRect(__r)
}

func QProxyStyle_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProxyStyle_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProxyStyle) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QProxyStyle) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProxyStyle) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QProxyStyle) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QProxyStyle_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProxyStyle_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProxyStyle_SliderPositionFromValue2(min int, max int, val int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_sliderPositionFromValue2, uintptr(min), uintptr(max), uintptr(val), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QProxyStyle_SliderValueFromPosition2(min int, max int, pos int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_sliderValueFromPosition2, uintptr(min), uintptr(max), uintptr(pos), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QProxyStyle_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProxyStyle_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProxyStyle) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QProxyStyle) DrawPrimitiveDefault(element QStyle__PrimitiveElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	purego.SyscallN(addr_QProxyStyle_drawPrimitive_default, this.h, uintptr(element), uintptr(option.UnsafePointer()), uintptr(painter.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(widget)
}

func (this *QProxyStyle) DrawControlDefault(element QStyle__ControlElement, option *QStyleOption, painter *QPainter, widget *QWidget) {
	purego.SyscallN(addr_QProxyStyle_drawControl_default, this.h, uintptr(element), uintptr(option.UnsafePointer()), uintptr(painter.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(widget)
}

func (this *QProxyStyle) DrawComplexControlDefault(control QStyle__ComplexControl, option *QStyleOptionComplex, painter *QPainter, widget *QWidget) {
	purego.SyscallN(addr_QProxyStyle_drawComplexControl_default, this.h, uintptr(control), uintptr(option.UnsafePointer()), uintptr(painter.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(widget)
}

func (this *QProxyStyle) DrawItemTextDefault(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QProxyStyle_drawItemText_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)), uintptr(textRole))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QProxyStyle) DrawItemPixmapDefault(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	purego.SyscallN(addr_QProxyStyle_drawItemPixmap_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pixmap)
}

func (this *QProxyStyle) SizeFromContentsDefault(type_ QStyle__ContentsType, option *QStyleOption, size *QSize, widget *QWidget) *QSize {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_sizeFromContents_default, this.h, uintptr(type_), uintptr(option.UnsafePointer()), uintptr(size.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(size)
	runtime.KeepAlive(widget)
	return newownedQSize(__r)
}

func (this *QProxyStyle) SubElementRectDefault(element QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_subElementRect_default, this.h, uintptr(element), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return newownedQRect(__r)
}

func (this *QProxyStyle) SubControlRectDefault(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_subControlRect_default, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(sc), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQRect(__r)
}

func (this *QProxyStyle) ItemTextRectDefault(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_itemTextRect_default, this.h, uintptr(fm.UnsafePointer()), uintptr(r.UnsafePointer()), uintptr(flags), uintptr(b2i(enabled)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fm)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QProxyStyle) ItemPixmapRectDefault(r *QRect, flags int, pixmap *QPixmap) *QRect {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_itemPixmapRect_default, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	return newownedQRect(__r)
}

func (this *QProxyStyle) HitTestComplexControlDefault(control QStyle__ComplexControl, option *QStyleOptionComplex, pos *QPoint, widget *QWidget) QStyle__SubControl {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_hitTestComplexControl_default, this.h, uintptr(control), uintptr(option.UnsafePointer()), uintptr(pos.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(widget)
	return QStyle__SubControl(__r)
}

func (this *QProxyStyle) StyleHintDefault(hint QStyle__StyleHint, option *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_styleHint_default, this.h, uintptr(hint), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()), uintptr(returnData.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(returnData)
	return int(__r)
}

func (this *QProxyStyle) PixelMetricDefault(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_pixelMetric_default, this.h, uintptr(metric), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QProxyStyle) LayoutSpacingDefault(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_layoutSpacing_default, this.h, uintptr(control1), uintptr(control2), uintptr(orientation), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QProxyStyle) StandardIconDefault(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_standardIcon_default, this.h, uintptr(standardIcon), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return newownedQIcon(__r)
}

func (this *QProxyStyle) StandardPixmapDefault(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_standardPixmap_default, this.h, uintptr(standardPixmap), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQPixmap(__r)
}

func (this *QProxyStyle) GeneratedIconPixmapDefault(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_generatedIconPixmap_default, this.h, uintptr(iconMode), uintptr(pixmap.UnsafePointer()), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(opt)
	return newownedQPixmap(__r)
}

func (this *QProxyStyle) StandardPaletteDefault() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QProxyStyle_standardPalette_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QProxyStyle) PolishDefault(widget *QWidget) {
	purego.SyscallN(addr_QProxyStyle_polish_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QProxyStyle) UnpolishDefault(widget *QWidget) {
	purego.SyscallN(addr_QProxyStyle_unpolish_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QProxyStyle) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QProxyStyle) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProxyStyle_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QProxyStyle) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QProxyStyle_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProxyStyle) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QProxyStyle_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProxyStyle) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProxyStyle_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProxyStyle) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProxyStyle_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProxyStyle) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProxyStyle_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProxyStyle) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProxyStyle_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProxyStyle) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QProxyStyle_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProxyStyle) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QProxyStyle_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProxyStyle) Disconnect(handle int64) {
	pg_QProxyStyle__disconnect(handle)
	unregisterCallback(handle)
}

func QProxyStyle_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6098275085281943360:
		return 1
	case -8697913287309583788:
		return 2
	case -778167327566211030:
		return 3
	case 2315857538243927747:
		return 4
	case -3740415202673381007:
		return 5
	case -3316981301817172730:
		return 6
	case -1580195784239416471:
		return 7
	case -8244607819117461060:
		return 8
	case 8974830244642632823:
		return 9
	case 3847657773274941773:
		return 10
	case 7058885155526130727:
		return 11
	case 2017029407462380579:
		return 12
	case -7621686327863150821:
		return 13
	case -4874522738816254324:
		return 14
	case -1870022100809264807:
		return 15
	case -5288619860977043091:
		return 16
	case -4387834948099211678:
		return 17
	case 7278964765692882119:
		return 18
	case -4633126903937721620:
		return 19
	case 4263633720130086714:
		return 20
	case 4263634819641714925:
		return 21
	case -5547720835653541585:
		return 22
	case 265826155645830215:
		return 23
	case 1237236456312824991:
		return 24
	case 5217623514630858297:
		return 25
	case -5409927581812374498:
		return 26
	case -1185315367526857833:
		return 27
	case 3769868229123868474:
		return 28
	case 7379029982170503008:
		return 29
	case -568369772982215962:
		return 30
	}
	return -1
}

func QProxyStyle_enable(h uintptr, idx int32) { pg_QProxyStyle__set_override(h, idx) }

func (this *QProxyStyle) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QProxyStyle) OnDrawPrimitive(fn func(QStyle__PrimitiveElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6098275085281943360, func(element int32, option uintptr, painter uintptr, widget uintptr) { fn(QStyle__PrimitiveElement(element), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "drawPrimitive"
}

func (this *QProxyStyle) OnDrawControl(fn func(QStyle__ControlElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -8697913287309583788, func(element int32, option uintptr, painter uintptr, widget uintptr) { fn(QStyle__ControlElement(element), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "drawControl"
}

func (this *QProxyStyle) OnDrawComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -778167327566211030, func(control int32, option uintptr, painter uintptr, widget uintptr) { fn(QStyle__ComplexControl(control), UnsafeNewQStyleOptionComplex(unsafe.Pointer(option)), UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "drawComplexControl"
}

func (this *QProxyStyle) OnDrawItemText(fn func(*QPainter, *QRect, int, *QPalette, bool, string, QPalette__ColorRole)) {
	registerVirt(this.UnsafeGoID(), 2315857538243927747, func(painter uintptr, rect uintptr, flags int32, pal uintptr, enabled int32, text uintptr, text_len int64, textRole int32) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(flags), UnsafeNewQPalette(unsafe.Pointer(pal)), enabled != 0, pgStrNoFree(text, text_len), QPalette__ColorRole(textRole)) }) // slot "drawItemText"
}

func (this *QProxyStyle) OnDrawItemPixmap(fn func(*QPainter, *QRect, int, *QPixmap)) {
	registerVirt(this.UnsafeGoID(), -3740415202673381007, func(painter uintptr, rect uintptr, alignment int32, pixmap uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(alignment), UnsafeNewQPixmap(unsafe.Pointer(pixmap))) }) // slot "drawItemPixmap"
}

func (this *QProxyStyle) OnSizeFromContents(fn func(QStyle__ContentsType, *QStyleOption, *QSize, *QWidget) *QSize) {
	registerVirt(this.UnsafeGoID(), -3316981301817172730, func(type_ int32, option uintptr, size uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__ContentsType(type_), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQSize(unsafe.Pointer(size)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "sizeFromContents"
}

func (this *QProxyStyle) OnSubElementRect(fn func(QStyle__SubElement, *QStyleOption, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -1580195784239416471, func(element int32, option uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__SubElement(element), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "subElementRect"
}

func (this *QProxyStyle) OnSubControlRect(fn func(QStyle__ComplexControl, *QStyleOptionComplex, QStyle__SubControl, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -8244607819117461060, func(cc int32, opt uintptr, sc int32, widget uintptr) uintptr { return uintptr(fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), QStyle__SubControl(sc), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "subControlRect"
}

func (this *QProxyStyle) OnItemTextRect(fn func(*QFontMetrics, *QRect, int, bool, string) *QRect) {
	registerVirt(this.UnsafeGoID(), 8974830244642632823, func(fm uintptr, r uintptr, flags int32, enabled int32, text uintptr, text_len int64) uintptr { return uintptr(fn(UnsafeNewQFontMetrics(unsafe.Pointer(fm)), UnsafeNewQRect(unsafe.Pointer(r)), int(flags), enabled != 0, pgStrNoFree(text, text_len)).UnsafePointer()) }) // slot "itemTextRect"
}

func (this *QProxyStyle) OnItemPixmapRect(fn func(*QRect, int, *QPixmap) *QRect) {
	registerVirt(this.UnsafeGoID(), 3847657773274941773, func(r uintptr, flags int32, pixmap uintptr) uintptr { return uintptr(fn(UnsafeNewQRect(unsafe.Pointer(r)), int(flags), UnsafeNewQPixmap(unsafe.Pointer(pixmap))).UnsafePointer()) }) // slot "itemPixmapRect"
}

func (this *QProxyStyle) OnHitTestComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPoint, *QWidget) QStyle__SubControl) {
	registerVirt(this.UnsafeGoID(), 7058885155526130727, func(control int32, option uintptr, pos uintptr, widget uintptr) int32 { return int32(fn(QStyle__ComplexControl(control), UnsafeNewQStyleOptionComplex(unsafe.Pointer(option)), UnsafeNewQPoint(unsafe.Pointer(pos)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "hitTestComplexControl"
}

func (this *QProxyStyle) OnStyleHint(fn func(QStyle__StyleHint, *QStyleOption, *QWidget, *QStyleHintReturn) int) {
	registerVirt(this.UnsafeGoID(), 2017029407462380579, func(hint int32, option uintptr, widget uintptr, returnData uintptr) int32 { return int32(fn(QStyle__StyleHint(hint), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)), UnsafeNewQStyleHintReturn(unsafe.Pointer(returnData)))) }) // slot "styleHint"
}

func (this *QProxyStyle) OnPixelMetric(fn func(QStyle__PixelMetric, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7621686327863150821, func(metric int32, option uintptr, widget uintptr) int32 { return int32(fn(QStyle__PixelMetric(metric), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "pixelMetric"
}

func (this *QProxyStyle) OnLayoutSpacing(fn func(QSizePolicy__ControlType, QSizePolicy__ControlType, Orientation, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -4874522738816254324, func(control1 int32, control2 int32, orientation int32, option uintptr, widget uintptr) int32 { return int32(fn(QSizePolicy__ControlType(control1), QSizePolicy__ControlType(control2), Orientation(orientation), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "layoutSpacing"
}

func (this *QProxyStyle) OnStandardIcon(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QIcon) {
	registerVirt(this.UnsafeGoID(), -1870022100809264807, func(standardIcon int32, option uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(standardIcon), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardIcon"
}

func (this *QProxyStyle) OnStandardPixmap(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -5288619860977043091, func(standardPixmap int32, opt uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(standardPixmap), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardPixmap"
}

func (this *QProxyStyle) OnGeneratedIconPixmap(fn func(QIcon__Mode, *QPixmap, *QStyleOption) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -4387834948099211678, func(iconMode int32, pixmap uintptr, opt uintptr) uintptr { return uintptr(fn(QIcon__Mode(iconMode), UnsafeNewQPixmap(unsafe.Pointer(pixmap)), UnsafeNewQStyleOption(unsafe.Pointer(opt))).UnsafePointer()) }) // slot "generatedIconPixmap"
}

func (this *QProxyStyle) OnStandardPalette(fn func() *QPalette) {
	registerVirt(this.UnsafeGoID(), 7278964765692882119, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "standardPalette"
}

func (this *QProxyStyle) OnPolish(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -4633126903937721620, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "polish"
}

func (this *QProxyStyle) OnPolish2(fn func(*QPalette)) {
	registerVirt(this.UnsafeGoID(), 4263633720130086714, func(pal uintptr) { fn(UnsafeNewQPalette(unsafe.Pointer(pal))) }) // slot "polish2"
}

func (this *QProxyStyle) OnPolish3(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 4263634819641714925, func(app uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(app))) }) // slot "polish3"
}

func (this *QProxyStyle) OnUnpolish(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -5547720835653541585, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "unpolish"
}

func (this *QProxyStyle) OnUnpolish2(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 265826155645830215, func(app uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(app))) }) // slot "unpolish2"
}

func (this *QProxyStyle) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QProxyStyle) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QProxyStyle) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QProxyStyle) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QProxyStyle) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QProxyStyle) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QProxyStyle) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QProxyStyle) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProxyStyle_delete(this.h)
	}
	this.UnsafeRelease()
}
