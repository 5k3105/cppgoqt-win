package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCommonStyle_chain uintptr
	addr_QCommonStyle_new uintptr
	addr_QCommonStyle_tr uintptr
	addr_QCommonStyle_visualRect uintptr
	addr_QCommonStyle_visualPos uintptr
	addr_QCommonStyle_sliderPositionFromValue uintptr
	addr_QCommonStyle_sliderValueFromPosition uintptr
	addr_QCommonStyle_visualAlignment uintptr
	addr_QCommonStyle_alignedRect uintptr
	addr_QCommonStyle_connect uintptr
	addr_QCommonStyle_connect2 uintptr
	addr_QCommonStyle_sender uintptr
	addr_QCommonStyle_senderSignalIndex uintptr
	addr_QCommonStyle_receivers uintptr
	addr_QCommonStyle_isSignalConnected uintptr
	addr_QCommonStyle_tr2 uintptr
	addr_QCommonStyle_tr3 uintptr
	addr_QCommonStyle_sliderPositionFromValue2 uintptr
	addr_QCommonStyle_sliderValueFromPosition2 uintptr
	addr_QCommonStyle_connect3 uintptr
	addr_QCommonStyle_connect4 uintptr
	addr_QCommonStyle_metaObject_default uintptr
	addr_QCommonStyle_drawPrimitive_default uintptr
	addr_QCommonStyle_drawControl_default uintptr
	addr_QCommonStyle_subElementRect_default uintptr
	addr_QCommonStyle_drawComplexControl_default uintptr
	addr_QCommonStyle_hitTestComplexControl_default uintptr
	addr_QCommonStyle_subControlRect_default uintptr
	addr_QCommonStyle_sizeFromContents_default uintptr
	addr_QCommonStyle_pixelMetric_default uintptr
	addr_QCommonStyle_styleHint_default uintptr
	addr_QCommonStyle_standardIcon_default uintptr
	addr_QCommonStyle_standardPixmap_default uintptr
	addr_QCommonStyle_generatedIconPixmap_default uintptr
	addr_QCommonStyle_layoutSpacing_default uintptr
	addr_QCommonStyle_polish_default uintptr
	addr_QCommonStyle_unpolish_default uintptr
	addr_QCommonStyle_itemTextRect_default uintptr
	addr_QCommonStyle_itemPixmapRect_default uintptr
	addr_QCommonStyle_drawItemText_default uintptr
	addr_QCommonStyle_drawItemPixmap_default uintptr
	addr_QCommonStyle_standardPalette_default uintptr
	addr_QCommonStyle_event_default uintptr
	addr_QCommonStyle_eventFilter_default uintptr
	addr_QCommonStyle_timerEvent_default uintptr
	addr_QCommonStyle_childEvent_default uintptr
	addr_QCommonStyle_customEvent_default uintptr
	addr_QCommonStyle_connectNotify_default uintptr
	addr_QCommonStyle_disconnectNotify_default uintptr
	pg_QCommonStyle_connect_destroyed func(uintptr, int64)
	pg_QCommonStyle_connect_destroyed2 func(uintptr, int64)
	pg_QCommonStyle_connect_objectNameChanged func(uintptr, int64)
	pg_QCommonStyle__disconnect func(int64)
	pg_QCommonStyle__set_override func(uintptr, int32)
	pg_QCommonStyle_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCommonStyle_chain, "QCommonStyle_chain")
	symaddr(&addr_QCommonStyle_new, "QCommonStyle_new")
	symaddr(&addr_QCommonStyle_tr, "QCommonStyle_tr")
	symaddr(&addr_QCommonStyle_visualRect, "QCommonStyle_visualRect")
	symaddr(&addr_QCommonStyle_visualPos, "QCommonStyle_visualPos")
	symaddr(&addr_QCommonStyle_sliderPositionFromValue, "QCommonStyle_sliderPositionFromValue")
	symaddr(&addr_QCommonStyle_sliderValueFromPosition, "QCommonStyle_sliderValueFromPosition")
	symaddr(&addr_QCommonStyle_visualAlignment, "QCommonStyle_visualAlignment")
	symaddr(&addr_QCommonStyle_alignedRect, "QCommonStyle_alignedRect")
	symaddr(&addr_QCommonStyle_connect, "QCommonStyle_connect")
	symaddr(&addr_QCommonStyle_connect2, "QCommonStyle_connect2")
	symaddr(&addr_QCommonStyle_sender, "QCommonStyle_sender")
	symaddr(&addr_QCommonStyle_senderSignalIndex, "QCommonStyle_senderSignalIndex")
	symaddr(&addr_QCommonStyle_receivers, "QCommonStyle_receivers")
	symaddr(&addr_QCommonStyle_isSignalConnected, "QCommonStyle_isSignalConnected")
	symaddr(&addr_QCommonStyle_tr2, "QCommonStyle_tr2")
	symaddr(&addr_QCommonStyle_tr3, "QCommonStyle_tr3")
	symaddr(&addr_QCommonStyle_sliderPositionFromValue2, "QCommonStyle_sliderPositionFromValue2")
	symaddr(&addr_QCommonStyle_sliderValueFromPosition2, "QCommonStyle_sliderValueFromPosition2")
	symaddr(&addr_QCommonStyle_connect3, "QCommonStyle_connect3")
	symaddr(&addr_QCommonStyle_connect4, "QCommonStyle_connect4")
	symaddr(&addr_QCommonStyle_metaObject_default, "QCommonStyle_metaObject_default")
	symaddr(&addr_QCommonStyle_drawPrimitive_default, "QCommonStyle_drawPrimitive_default")
	symaddr(&addr_QCommonStyle_drawControl_default, "QCommonStyle_drawControl_default")
	symaddr(&addr_QCommonStyle_subElementRect_default, "QCommonStyle_subElementRect_default")
	symaddr(&addr_QCommonStyle_drawComplexControl_default, "QCommonStyle_drawComplexControl_default")
	symaddr(&addr_QCommonStyle_hitTestComplexControl_default, "QCommonStyle_hitTestComplexControl_default")
	symaddr(&addr_QCommonStyle_subControlRect_default, "QCommonStyle_subControlRect_default")
	symaddr(&addr_QCommonStyle_sizeFromContents_default, "QCommonStyle_sizeFromContents_default")
	symaddr(&addr_QCommonStyle_pixelMetric_default, "QCommonStyle_pixelMetric_default")
	symaddr(&addr_QCommonStyle_styleHint_default, "QCommonStyle_styleHint_default")
	symaddr(&addr_QCommonStyle_standardIcon_default, "QCommonStyle_standardIcon_default")
	symaddr(&addr_QCommonStyle_standardPixmap_default, "QCommonStyle_standardPixmap_default")
	symaddr(&addr_QCommonStyle_generatedIconPixmap_default, "QCommonStyle_generatedIconPixmap_default")
	symaddr(&addr_QCommonStyle_layoutSpacing_default, "QCommonStyle_layoutSpacing_default")
	symaddr(&addr_QCommonStyle_polish_default, "QCommonStyle_polish_default")
	symaddr(&addr_QCommonStyle_unpolish_default, "QCommonStyle_unpolish_default")
	symaddr(&addr_QCommonStyle_itemTextRect_default, "QCommonStyle_itemTextRect_default")
	symaddr(&addr_QCommonStyle_itemPixmapRect_default, "QCommonStyle_itemPixmapRect_default")
	symaddr(&addr_QCommonStyle_drawItemText_default, "QCommonStyle_drawItemText_default")
	symaddr(&addr_QCommonStyle_drawItemPixmap_default, "QCommonStyle_drawItemPixmap_default")
	symaddr(&addr_QCommonStyle_standardPalette_default, "QCommonStyle_standardPalette_default")
	symaddr(&addr_QCommonStyle_event_default, "QCommonStyle_event_default")
	symaddr(&addr_QCommonStyle_eventFilter_default, "QCommonStyle_eventFilter_default")
	symaddr(&addr_QCommonStyle_timerEvent_default, "QCommonStyle_timerEvent_default")
	symaddr(&addr_QCommonStyle_childEvent_default, "QCommonStyle_childEvent_default")
	symaddr(&addr_QCommonStyle_customEvent_default, "QCommonStyle_customEvent_default")
	symaddr(&addr_QCommonStyle_connectNotify_default, "QCommonStyle_connectNotify_default")
	symaddr(&addr_QCommonStyle_disconnectNotify_default, "QCommonStyle_disconnectNotify_default")
	reg(&pg_QCommonStyle_connect_destroyed, "QCommonStyle_connect_destroyed")
	reg(&pg_QCommonStyle_connect_destroyed2, "QCommonStyle_connect_destroyed2")
	reg(&pg_QCommonStyle_connect_objectNameChanged, "QCommonStyle_connect_objectNameChanged")
	reg(&pg_QCommonStyle__disconnect, "QCommonStyle__disconnect")
	reg(&pg_QCommonStyle__set_override, "QCommonStyle__set_override")
	reg(&pg_QCommonStyle_delete, "QCommonStyle_delete")
}

type QCommonStyle struct {
	h uintptr
	*QStyle
}

func bindQCommonStyle(h uintptr) *QCommonStyle {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QCommonStyle_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QCommonStyle{h: h, QStyle: &QStyle{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQCommonStyle(h uintptr) *QCommonStyle { return bindQCommonStyle(h) }

func (this *QCommonStyle) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCommonStyle) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCommonStyle(h unsafe.Pointer) *QCommonStyle { return bindQCommonStyle(uintptr(h)) }
func UnsafeNewQCommonStyleOwned(h unsafe.Pointer) *QCommonStyle { return newownedQCommonStyle(uintptr(h)) }

func (this *QCommonStyle) UnsafeRelease() {
	this.h = 0
	if this.QStyle != nil {
		this.QStyle.UnsafeRelease()
	}
}

func NewQCommonStyle() *QCommonStyle {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_new, uintptr(__goid))
	__o := newownedQCommonStyle(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommonStyle_virtIdx, QCommonStyle_enable)
	return __o
}

func QCommonStyle_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QCommonStyle_VisualRect(direction LayoutDirection, boundingRect *QRect, logicalRect *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_visualRect, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalRect.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalRect)
	return newownedQRect(__r)
}

func QCommonStyle_VisualPos(direction LayoutDirection, boundingRect *QRect, logicalPos *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_visualPos, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalPos.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalPos)
	return newownedQPoint(__r)
}

func QCommonStyle_SliderPositionFromValue(min int, max int, val int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_sliderPositionFromValue, uintptr(min), uintptr(max), uintptr(val), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCommonStyle_SliderValueFromPosition(min int, max int, pos int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_sliderValueFromPosition, uintptr(min), uintptr(max), uintptr(pos), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCommonStyle_VisualAlignment(direction LayoutDirection, alignment AlignmentFlag) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_visualAlignment, uintptr(direction), uintptr(alignment))
	__r := int32(uint32(__r0))
	return AlignmentFlag(__r)
}

func QCommonStyle_AlignedRect(direction LayoutDirection, alignment AlignmentFlag, size *QSize, rectangle *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_alignedRect, uintptr(direction), uintptr(alignment), uintptr(size.UnsafePointer()), uintptr(rectangle.UnsafePointer()))
	runtime.KeepAlive(size)
	runtime.KeepAlive(rectangle)
	return newownedQRect(__r)
}

func QCommonStyle_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCommonStyle_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCommonStyle) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QCommonStyle) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCommonStyle) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QCommonStyle) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QCommonStyle_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCommonStyle_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCommonStyle_SliderPositionFromValue2(min int, max int, val int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_sliderPositionFromValue2, uintptr(min), uintptr(max), uintptr(val), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCommonStyle_SliderValueFromPosition2(min int, max int, pos int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_sliderValueFromPosition2, uintptr(min), uintptr(max), uintptr(pos), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCommonStyle_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCommonStyle_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCommonStyle) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QCommonStyle) DrawPrimitiveDefault(pe QStyle__PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	purego.SyscallN(addr_QCommonStyle_drawPrimitive_default, this.h, uintptr(pe), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(w)
}

func (this *QCommonStyle) DrawControlDefault(element QStyle__ControlElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	purego.SyscallN(addr_QCommonStyle_drawControl_default, this.h, uintptr(element), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(w)
}

func (this *QCommonStyle) SubElementRectDefault(r QStyle__SubElement, opt *QStyleOption, widget *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_subElementRect_default, this.h, uintptr(r), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQRect(__r)
}

func (this *QCommonStyle) DrawComplexControlDefault(cc QStyle__ComplexControl, opt *QStyleOptionComplex, p *QPainter, w *QWidget) {
	purego.SyscallN(addr_QCommonStyle_drawComplexControl_default, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(w)
}

func (this *QCommonStyle) HitTestComplexControlDefault(cc QStyle__ComplexControl, opt *QStyleOptionComplex, pt *QPoint, w *QWidget) QStyle__SubControl {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_hitTestComplexControl_default, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(pt.UnsafePointer()), uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(pt)
	runtime.KeepAlive(w)
	return QStyle__SubControl(__r)
}

func (this *QCommonStyle) SubControlRectDefault(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, w *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_subControlRect_default, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(sc), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(w)
	return newownedQRect(__r)
}

func (this *QCommonStyle) SizeFromContentsDefault(ct QStyle__ContentsType, opt *QStyleOption, contentsSize *QSize, widget *QWidget) *QSize {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_sizeFromContents_default, this.h, uintptr(ct), uintptr(opt.UnsafePointer()), uintptr(contentsSize.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(contentsSize)
	runtime.KeepAlive(widget)
	return newownedQSize(__r)
}

func (this *QCommonStyle) PixelMetricDefault(m QStyle__PixelMetric, opt *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_pixelMetric_default, this.h, uintptr(m), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QCommonStyle) StyleHintDefault(sh QStyle__StyleHint, opt *QStyleOption, w *QWidget, shret *QStyleHintReturn) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_styleHint_default, this.h, uintptr(sh), uintptr(opt.UnsafePointer()), uintptr(w.UnsafePointer()), uintptr(shret.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(w)
	runtime.KeepAlive(shret)
	return int(__r)
}

func (this *QCommonStyle) StandardIconDefault(standardIcon QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_standardIcon_default, this.h, uintptr(standardIcon), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQIcon(__r)
}

func (this *QCommonStyle) StandardPixmapDefault(sp QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_standardPixmap_default, this.h, uintptr(sp), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQPixmap(__r)
}

func (this *QCommonStyle) GeneratedIconPixmapDefault(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_generatedIconPixmap_default, this.h, uintptr(iconMode), uintptr(pixmap.UnsafePointer()), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(opt)
	return newownedQPixmap(__r)
}

func (this *QCommonStyle) LayoutSpacingDefault(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_layoutSpacing_default, this.h, uintptr(control1), uintptr(control2), uintptr(orientation), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QCommonStyle) PolishDefault(arg1 *QPalette) {
	purego.SyscallN(addr_QCommonStyle_polish_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommonStyle) UnpolishDefault(widget *QWidget) {
	purego.SyscallN(addr_QCommonStyle_unpolish_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QCommonStyle) ItemTextRectDefault(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_itemTextRect_default, this.h, uintptr(fm.UnsafePointer()), uintptr(r.UnsafePointer()), uintptr(flags), uintptr(b2i(enabled)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fm)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QCommonStyle) ItemPixmapRectDefault(r *QRect, flags int, pixmap *QPixmap) *QRect {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_itemPixmapRect_default, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	return newownedQRect(__r)
}

func (this *QCommonStyle) DrawItemTextDefault(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QCommonStyle_drawItemText_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)), uintptr(textRole))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QCommonStyle) DrawItemPixmapDefault(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	purego.SyscallN(addr_QCommonStyle_drawItemPixmap_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pixmap)
}

func (this *QCommonStyle) StandardPaletteDefault() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QCommonStyle_standardPalette_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QCommonStyle) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCommonStyle) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommonStyle_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCommonStyle) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QCommonStyle_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommonStyle) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QCommonStyle_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommonStyle) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCommonStyle_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommonStyle) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCommonStyle_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCommonStyle) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCommonStyle_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCommonStyle) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCommonStyle_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommonStyle) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCommonStyle_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommonStyle) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCommonStyle_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommonStyle) Disconnect(handle int64) {
	pg_QCommonStyle__disconnect(handle)
	unregisterCallback(handle)
}

func QCommonStyle_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6098275085281943360:
		return 1
	case -8697913287309583788:
		return 2
	case -1580195784239416471:
		return 3
	case -778167327566211030:
		return 4
	case 7058885155526130727:
		return 5
	case -8244607819117461060:
		return 6
	case -3316981301817172730:
		return 7
	case -7621686327863150821:
		return 8
	case 2017029407462380579:
		return 9
	case -1870022100809264807:
		return 10
	case -5288619860977043091:
		return 11
	case -4387834948099211678:
		return 12
	case -4874522738816254324:
		return 13
	case -4633126903937721620:
		return 14
	case 4263633720130086714:
		return 15
	case 4263634819641714925:
		return 16
	case -5547720835653541585:
		return 17
	case 265826155645830215:
		return 18
	case 8974830244642632823:
		return 19
	case 3847657773274941773:
		return 20
	case 2315857538243927747:
		return 21
	case -3740415202673381007:
		return 22
	case 7278964765692882119:
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

func QCommonStyle_enable(h uintptr, idx int32) { pg_QCommonStyle__set_override(h, idx) }

func (this *QCommonStyle) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QCommonStyle) OnDrawPrimitive(fn func(QStyle__PrimitiveElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6098275085281943360, func(pe int32, opt uintptr, p uintptr, w uintptr) { fn(QStyle__PrimitiveElement(pe), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(w))) }) // slot "drawPrimitive"
}

func (this *QCommonStyle) OnDrawControl(fn func(QStyle__ControlElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -8697913287309583788, func(element int32, opt uintptr, p uintptr, w uintptr) { fn(QStyle__ControlElement(element), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(w))) }) // slot "drawControl"
}

func (this *QCommonStyle) OnSubElementRect(fn func(QStyle__SubElement, *QStyleOption, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -1580195784239416471, func(r int32, opt uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__SubElement(r), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "subElementRect"
}

func (this *QCommonStyle) OnDrawComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -778167327566211030, func(cc int32, opt uintptr, p uintptr, w uintptr) { fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(w))) }) // slot "drawComplexControl"
}

func (this *QCommonStyle) OnHitTestComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPoint, *QWidget) QStyle__SubControl) {
	registerVirt(this.UnsafeGoID(), 7058885155526130727, func(cc int32, opt uintptr, pt uintptr, w uintptr) int32 { return int32(fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), UnsafeNewQPoint(unsafe.Pointer(pt)), UnsafeNewQWidget(unsafe.Pointer(w)))) }) // slot "hitTestComplexControl"
}

func (this *QCommonStyle) OnSubControlRect(fn func(QStyle__ComplexControl, *QStyleOptionComplex, QStyle__SubControl, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -8244607819117461060, func(cc int32, opt uintptr, sc int32, w uintptr) uintptr { return uintptr(fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), QStyle__SubControl(sc), UnsafeNewQWidget(unsafe.Pointer(w))).UnsafePointer()) }) // slot "subControlRect"
}

func (this *QCommonStyle) OnSizeFromContents(fn func(QStyle__ContentsType, *QStyleOption, *QSize, *QWidget) *QSize) {
	registerVirt(this.UnsafeGoID(), -3316981301817172730, func(ct int32, opt uintptr, contentsSize uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__ContentsType(ct), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQSize(unsafe.Pointer(contentsSize)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "sizeFromContents"
}

func (this *QCommonStyle) OnPixelMetric(fn func(QStyle__PixelMetric, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7621686327863150821, func(m int32, opt uintptr, widget uintptr) int32 { return int32(fn(QStyle__PixelMetric(m), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "pixelMetric"
}

func (this *QCommonStyle) OnStyleHint(fn func(QStyle__StyleHint, *QStyleOption, *QWidget, *QStyleHintReturn) int) {
	registerVirt(this.UnsafeGoID(), 2017029407462380579, func(sh int32, opt uintptr, w uintptr, shret uintptr) int32 { return int32(fn(QStyle__StyleHint(sh), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(w)), UnsafeNewQStyleHintReturn(unsafe.Pointer(shret)))) }) // slot "styleHint"
}

func (this *QCommonStyle) OnStandardIcon(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QIcon) {
	registerVirt(this.UnsafeGoID(), -1870022100809264807, func(standardIcon int32, opt uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(standardIcon), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardIcon"
}

func (this *QCommonStyle) OnStandardPixmap(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -5288619860977043091, func(sp int32, opt uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(sp), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardPixmap"
}

func (this *QCommonStyle) OnGeneratedIconPixmap(fn func(QIcon__Mode, *QPixmap, *QStyleOption) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -4387834948099211678, func(iconMode int32, pixmap uintptr, opt uintptr) uintptr { return uintptr(fn(QIcon__Mode(iconMode), UnsafeNewQPixmap(unsafe.Pointer(pixmap)), UnsafeNewQStyleOption(unsafe.Pointer(opt))).UnsafePointer()) }) // slot "generatedIconPixmap"
}

func (this *QCommonStyle) OnLayoutSpacing(fn func(QSizePolicy__ControlType, QSizePolicy__ControlType, Orientation, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -4874522738816254324, func(control1 int32, control2 int32, orientation int32, option uintptr, widget uintptr) int32 { return int32(fn(QSizePolicy__ControlType(control1), QSizePolicy__ControlType(control2), Orientation(orientation), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "layoutSpacing"
}

func (this *QCommonStyle) OnPolish(fn func(*QPalette)) {
	registerVirt(this.UnsafeGoID(), -4633126903937721620, func(arg1 uintptr) { fn(UnsafeNewQPalette(unsafe.Pointer(arg1))) }) // slot "polish"
}

func (this *QCommonStyle) OnPolish2(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 4263633720130086714, func(app uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(app))) }) // slot "polish2"
}

func (this *QCommonStyle) OnPolish3(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4263634819641714925, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "polish3"
}

func (this *QCommonStyle) OnUnpolish(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -5547720835653541585, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "unpolish"
}

func (this *QCommonStyle) OnUnpolish2(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 265826155645830215, func(application uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(application))) }) // slot "unpolish2"
}

func (this *QCommonStyle) OnItemTextRect(fn func(*QFontMetrics, *QRect, int, bool, string) *QRect) {
	registerVirt(this.UnsafeGoID(), 8974830244642632823, func(fm uintptr, r uintptr, flags int32, enabled int32, text uintptr, text_len int64) uintptr { return uintptr(fn(UnsafeNewQFontMetrics(unsafe.Pointer(fm)), UnsafeNewQRect(unsafe.Pointer(r)), int(flags), enabled != 0, pgStrNoFree(text, text_len)).UnsafePointer()) }) // slot "itemTextRect"
}

func (this *QCommonStyle) OnItemPixmapRect(fn func(*QRect, int, *QPixmap) *QRect) {
	registerVirt(this.UnsafeGoID(), 3847657773274941773, func(r uintptr, flags int32, pixmap uintptr) uintptr { return uintptr(fn(UnsafeNewQRect(unsafe.Pointer(r)), int(flags), UnsafeNewQPixmap(unsafe.Pointer(pixmap))).UnsafePointer()) }) // slot "itemPixmapRect"
}

func (this *QCommonStyle) OnDrawItemText(fn func(*QPainter, *QRect, int, *QPalette, bool, string, QPalette__ColorRole)) {
	registerVirt(this.UnsafeGoID(), 2315857538243927747, func(painter uintptr, rect uintptr, flags int32, pal uintptr, enabled int32, text uintptr, text_len int64, textRole int32) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(flags), UnsafeNewQPalette(unsafe.Pointer(pal)), enabled != 0, pgStrNoFree(text, text_len), QPalette__ColorRole(textRole)) }) // slot "drawItemText"
}

func (this *QCommonStyle) OnDrawItemPixmap(fn func(*QPainter, *QRect, int, *QPixmap)) {
	registerVirt(this.UnsafeGoID(), -3740415202673381007, func(painter uintptr, rect uintptr, alignment int32, pixmap uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(alignment), UnsafeNewQPixmap(unsafe.Pointer(pixmap))) }) // slot "drawItemPixmap"
}

func (this *QCommonStyle) OnStandardPalette(fn func() *QPalette) {
	registerVirt(this.UnsafeGoID(), 7278964765692882119, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "standardPalette"
}

func (this *QCommonStyle) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QCommonStyle) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QCommonStyle) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QCommonStyle) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QCommonStyle) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QCommonStyle) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QCommonStyle) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QCommonStyle) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCommonStyle_delete(this.h)
	}
	this.UnsafeRelease()
}
