package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyle_base func(uintptr) uintptr
	addr_QStyle_tr uintptr
	addr_QStyle_name uintptr
	addr_QStyle_polish uintptr
	addr_QStyle_unpolish uintptr
	addr_QStyle_polish2 uintptr
	addr_QStyle_unpolish2 uintptr
	addr_QStyle_polish3 uintptr
	addr_QStyle_itemTextRect uintptr
	addr_QStyle_itemPixmapRect uintptr
	addr_QStyle_drawItemText uintptr
	addr_QStyle_drawItemPixmap uintptr
	addr_QStyle_standardPalette uintptr
	addr_QStyle_drawPrimitive uintptr
	addr_QStyle_drawControl uintptr
	addr_QStyle_subElementRect uintptr
	addr_QStyle_drawComplexControl uintptr
	addr_QStyle_hitTestComplexControl uintptr
	addr_QStyle_subControlRect uintptr
	addr_QStyle_pixelMetric uintptr
	addr_QStyle_sizeFromContents uintptr
	addr_QStyle_styleHint uintptr
	addr_QStyle_standardPixmap uintptr
	addr_QStyle_standardIcon uintptr
	addr_QStyle_generatedIconPixmap uintptr
	addr_QStyle_visualRect uintptr
	addr_QStyle_visualPos uintptr
	addr_QStyle_sliderPositionFromValue uintptr
	addr_QStyle_sliderValueFromPosition uintptr
	addr_QStyle_visualAlignment uintptr
	addr_QStyle_alignedRect uintptr
	addr_QStyle_layoutSpacing uintptr
	addr_QStyle_combinedLayoutSpacing uintptr
	addr_QStyle_proxy uintptr
	addr_QStyle_connect uintptr
	addr_QStyle_connect2 uintptr
	addr_QStyle_sender uintptr
	addr_QStyle_senderSignalIndex uintptr
	addr_QStyle_receivers uintptr
	addr_QStyle_isSignalConnected uintptr
	addr_QStyle_tr2 uintptr
	addr_QStyle_tr3 uintptr
	addr_QStyle_sliderPositionFromValue2 uintptr
	addr_QStyle_sliderValueFromPosition2 uintptr
	addr_QStyle_combinedLayoutSpacing2 uintptr
	addr_QStyle_combinedLayoutSpacing3 uintptr
	addr_QStyle_connect3 uintptr
	addr_QStyle_connect4 uintptr
	addr_QStyle_polish_default uintptr
	addr_QStyle_unpolish_default uintptr
	addr_QStyle_polish_default2 uintptr
	addr_QStyle_unpolish_default2 uintptr
	addr_QStyle_polish_default3 uintptr
	addr_QStyle_itemTextRect_default uintptr
	addr_QStyle_itemPixmapRect_default uintptr
	addr_QStyle_drawItemText_default uintptr
	addr_QStyle_drawItemPixmap_default uintptr
	addr_QStyle_standardPalette_default uintptr
	addr_QStyle_metaObject_default uintptr
	addr_QStyle_event_default uintptr
	addr_QStyle_eventFilter_default uintptr
	addr_QStyle_timerEvent_default uintptr
	addr_QStyle_childEvent_default uintptr
	addr_QStyle_customEvent_default uintptr
	addr_QStyle_connectNotify_default uintptr
	addr_QStyle_disconnectNotify_default uintptr
	pg_QStyle_connect_destroyed func(uintptr, int64)
	pg_QStyle_connect_destroyed2 func(uintptr, int64)
	pg_QStyle_connect_objectNameChanged func(uintptr, int64)
	pg_QStyle__disconnect func(int64)
	pg_QStyle__set_override func(uintptr, int32)
	pg_QStyle__Go_new func(int64) uintptr
	pg_QStyle_delete func(uintptr)
)

func init() {
	reg(&pg_QStyle_base, "QStyle_base")
	symaddr(&addr_QStyle_tr, "QStyle_tr")
	symaddr(&addr_QStyle_name, "QStyle_name")
	symaddr(&addr_QStyle_polish, "QStyle_polish")
	symaddr(&addr_QStyle_unpolish, "QStyle_unpolish")
	symaddr(&addr_QStyle_polish2, "QStyle_polish2")
	symaddr(&addr_QStyle_unpolish2, "QStyle_unpolish2")
	symaddr(&addr_QStyle_polish3, "QStyle_polish3")
	symaddr(&addr_QStyle_itemTextRect, "QStyle_itemTextRect")
	symaddr(&addr_QStyle_itemPixmapRect, "QStyle_itemPixmapRect")
	symaddr(&addr_QStyle_drawItemText, "QStyle_drawItemText")
	symaddr(&addr_QStyle_drawItemPixmap, "QStyle_drawItemPixmap")
	symaddr(&addr_QStyle_standardPalette, "QStyle_standardPalette")
	symaddr(&addr_QStyle_drawPrimitive, "QStyle_drawPrimitive")
	symaddr(&addr_QStyle_drawControl, "QStyle_drawControl")
	symaddr(&addr_QStyle_subElementRect, "QStyle_subElementRect")
	symaddr(&addr_QStyle_drawComplexControl, "QStyle_drawComplexControl")
	symaddr(&addr_QStyle_hitTestComplexControl, "QStyle_hitTestComplexControl")
	symaddr(&addr_QStyle_subControlRect, "QStyle_subControlRect")
	symaddr(&addr_QStyle_pixelMetric, "QStyle_pixelMetric")
	symaddr(&addr_QStyle_sizeFromContents, "QStyle_sizeFromContents")
	symaddr(&addr_QStyle_styleHint, "QStyle_styleHint")
	symaddr(&addr_QStyle_standardPixmap, "QStyle_standardPixmap")
	symaddr(&addr_QStyle_standardIcon, "QStyle_standardIcon")
	symaddr(&addr_QStyle_generatedIconPixmap, "QStyle_generatedIconPixmap")
	symaddr(&addr_QStyle_visualRect, "QStyle_visualRect")
	symaddr(&addr_QStyle_visualPos, "QStyle_visualPos")
	symaddr(&addr_QStyle_sliderPositionFromValue, "QStyle_sliderPositionFromValue")
	symaddr(&addr_QStyle_sliderValueFromPosition, "QStyle_sliderValueFromPosition")
	symaddr(&addr_QStyle_visualAlignment, "QStyle_visualAlignment")
	symaddr(&addr_QStyle_alignedRect, "QStyle_alignedRect")
	symaddr(&addr_QStyle_layoutSpacing, "QStyle_layoutSpacing")
	symaddr(&addr_QStyle_combinedLayoutSpacing, "QStyle_combinedLayoutSpacing")
	symaddr(&addr_QStyle_proxy, "QStyle_proxy")
	symaddr(&addr_QStyle_connect, "QStyle_connect")
	symaddr(&addr_QStyle_connect2, "QStyle_connect2")
	symaddr(&addr_QStyle_sender, "QStyle_sender")
	symaddr(&addr_QStyle_senderSignalIndex, "QStyle_senderSignalIndex")
	symaddr(&addr_QStyle_receivers, "QStyle_receivers")
	symaddr(&addr_QStyle_isSignalConnected, "QStyle_isSignalConnected")
	symaddr(&addr_QStyle_tr2, "QStyle_tr2")
	symaddr(&addr_QStyle_tr3, "QStyle_tr3")
	symaddr(&addr_QStyle_sliderPositionFromValue2, "QStyle_sliderPositionFromValue2")
	symaddr(&addr_QStyle_sliderValueFromPosition2, "QStyle_sliderValueFromPosition2")
	symaddr(&addr_QStyle_combinedLayoutSpacing2, "QStyle_combinedLayoutSpacing2")
	symaddr(&addr_QStyle_combinedLayoutSpacing3, "QStyle_combinedLayoutSpacing3")
	symaddr(&addr_QStyle_connect3, "QStyle_connect3")
	symaddr(&addr_QStyle_connect4, "QStyle_connect4")
	symaddr(&addr_QStyle_polish_default, "QStyle_polish_default")
	symaddr(&addr_QStyle_unpolish_default, "QStyle_unpolish_default")
	symaddr(&addr_QStyle_polish_default2, "QStyle_polish_default2")
	symaddr(&addr_QStyle_unpolish_default2, "QStyle_unpolish_default2")
	symaddr(&addr_QStyle_polish_default3, "QStyle_polish_default3")
	symaddr(&addr_QStyle_itemTextRect_default, "QStyle_itemTextRect_default")
	symaddr(&addr_QStyle_itemPixmapRect_default, "QStyle_itemPixmapRect_default")
	symaddr(&addr_QStyle_drawItemText_default, "QStyle_drawItemText_default")
	symaddr(&addr_QStyle_drawItemPixmap_default, "QStyle_drawItemPixmap_default")
	symaddr(&addr_QStyle_standardPalette_default, "QStyle_standardPalette_default")
	symaddr(&addr_QStyle_metaObject_default, "QStyle_metaObject_default")
	symaddr(&addr_QStyle_event_default, "QStyle_event_default")
	symaddr(&addr_QStyle_eventFilter_default, "QStyle_eventFilter_default")
	symaddr(&addr_QStyle_timerEvent_default, "QStyle_timerEvent_default")
	symaddr(&addr_QStyle_childEvent_default, "QStyle_childEvent_default")
	symaddr(&addr_QStyle_customEvent_default, "QStyle_customEvent_default")
	symaddr(&addr_QStyle_connectNotify_default, "QStyle_connectNotify_default")
	symaddr(&addr_QStyle_disconnectNotify_default, "QStyle_disconnectNotify_default")
	reg(&pg_QStyle_connect_destroyed, "QStyle_connect_destroyed")
	reg(&pg_QStyle_connect_destroyed2, "QStyle_connect_destroyed2")
	reg(&pg_QStyle_connect_objectNameChanged, "QStyle_connect_objectNameChanged")
	reg(&pg_QStyle__disconnect, "QStyle__disconnect")
	reg(&pg_QStyle__set_override, "QStyle__set_override")
	reg(&pg_QStyle__Go_new, "QStyle__Go_new")
	reg(&pg_QStyle_delete, "QStyle_delete")
}

type QStyle struct {
	h uintptr
	*QObject
}

func bindQStyle(h uintptr) *QStyle {
	if h == 0 {
		return nil
	}
	return &QStyle{h: h, QObject: bindQObject(pg_QStyle_base(h))}
}

func newownedQStyle(h uintptr) *QStyle { return bindQStyle(h) }

func (this *QStyle) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyle) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyle(h unsafe.Pointer) *QStyle { return bindQStyle(uintptr(h)) }
func UnsafeNewQStyleOwned(h unsafe.Pointer) *QStyle { return newownedQStyle(uintptr(h)) }

func (this *QStyle) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QStyle_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyle_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStyle) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyle_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyle) Polish(widget *QWidget) {
	purego.SyscallN(addr_QStyle_polish, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStyle) Unpolish(widget *QWidget) {
	purego.SyscallN(addr_QStyle_unpolish, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStyle) PolishWithApplication(application *QApplication) {
	purego.SyscallN(addr_QStyle_polish2, this.h, uintptr(application.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(application)
}

func (this *QStyle) UnpolishWithApplication(application *QApplication) {
	purego.SyscallN(addr_QStyle_unpolish2, this.h, uintptr(application.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(application)
}

func (this *QStyle) PolishWithPalette(palette *QPalette) {
	purego.SyscallN(addr_QStyle_polish3, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QStyle) ItemTextRect(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QStyle_itemTextRect, this.h, uintptr(fm.UnsafePointer()), uintptr(r.UnsafePointer()), uintptr(flags), uintptr(b2i(enabled)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fm)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QStyle) ItemPixmapRect(r *QRect, flags int, pixmap *QPixmap) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_itemPixmapRect, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	return newownedQRect(__r)
}

func (this *QStyle) DrawItemText(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyle_drawItemText, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)), uintptr(textRole))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QStyle) DrawItemPixmap(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	purego.SyscallN(addr_QStyle_drawItemPixmap, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pixmap)
}

func (this *QStyle) StandardPalette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QStyle_standardPalette, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QStyle) DrawPrimitive(pe QStyle__PrimitiveElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	purego.SyscallN(addr_QStyle_drawPrimitive, this.h, uintptr(pe), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(w)
}

func (this *QStyle) DrawControl(element QStyle__ControlElement, opt *QStyleOption, p *QPainter, w *QWidget) {
	purego.SyscallN(addr_QStyle_drawControl, this.h, uintptr(element), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(w)
}

func (this *QStyle) SubElementRect(subElement QStyle__SubElement, option *QStyleOption, widget *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_subElementRect, this.h, uintptr(subElement), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return newownedQRect(__r)
}

func (this *QStyle) DrawComplexControl(cc QStyle__ComplexControl, opt *QStyleOptionComplex, p *QPainter, widget *QWidget) {
	purego.SyscallN(addr_QStyle_drawComplexControl, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(p.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(p)
	runtime.KeepAlive(widget)
}

func (this *QStyle) HitTestComplexControl(cc QStyle__ComplexControl, opt *QStyleOptionComplex, pt *QPoint, widget *QWidget) QStyle__SubControl {
	__r0, _, _ := purego.SyscallN(addr_QStyle_hitTestComplexControl, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(pt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(pt)
	runtime.KeepAlive(widget)
	return QStyle__SubControl(__r)
}

func (this *QStyle) SubControlRect(cc QStyle__ComplexControl, opt *QStyleOptionComplex, sc QStyle__SubControl, widget *QWidget) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_subControlRect, this.h, uintptr(cc), uintptr(opt.UnsafePointer()), uintptr(sc), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQRect(__r)
}

func (this *QStyle) PixelMetric(metric QStyle__PixelMetric, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_pixelMetric, this.h, uintptr(metric), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStyle) SizeFromContents(ct QStyle__ContentsType, opt *QStyleOption, contentsSize *QSize, w *QWidget) *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyle_sizeFromContents, this.h, uintptr(ct), uintptr(opt.UnsafePointer()), uintptr(contentsSize.UnsafePointer()), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(contentsSize)
	runtime.KeepAlive(w)
	return newownedQSize(__r)
}

func (this *QStyle) StyleHint(stylehint QStyle__StyleHint, opt *QStyleOption, widget *QWidget, returnData *QStyleHintReturn) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_styleHint, this.h, uintptr(stylehint), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()), uintptr(returnData.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(returnData)
	return int(__r)
}

func (this *QStyle) StandardPixmap(standardPixmap QStyle__StandardPixmap, opt *QStyleOption, widget *QWidget) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QStyle_standardPixmap, this.h, uintptr(standardPixmap), uintptr(opt.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
	runtime.KeepAlive(widget)
	return newownedQPixmap(__r)
}

func (this *QStyle) StandardIcon(standardIcon QStyle__StandardPixmap, option *QStyleOption, widget *QWidget) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyle_standardIcon, this.h, uintptr(standardIcon), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return newownedQIcon(__r)
}

func (this *QStyle) GeneratedIconPixmap(iconMode QIcon__Mode, pixmap *QPixmap, opt *QStyleOption) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QStyle_generatedIconPixmap, this.h, uintptr(iconMode), uintptr(pixmap.UnsafePointer()), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(opt)
	return newownedQPixmap(__r)
}

func QStyle_VisualRect(direction LayoutDirection, boundingRect *QRect, logicalRect *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_visualRect, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalRect.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalRect)
	return newownedQRect(__r)
}

func QStyle_VisualPos(direction LayoutDirection, boundingRect *QRect, logicalPos *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QStyle_visualPos, uintptr(direction), uintptr(boundingRect.UnsafePointer()), uintptr(logicalPos.UnsafePointer()))
	runtime.KeepAlive(boundingRect)
	runtime.KeepAlive(logicalPos)
	return newownedQPoint(__r)
}

func QStyle_SliderPositionFromValue(min int, max int, val int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_sliderPositionFromValue, uintptr(min), uintptr(max), uintptr(val), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QStyle_SliderValueFromPosition(min int, max int, pos int, space int) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_sliderValueFromPosition, uintptr(min), uintptr(max), uintptr(pos), uintptr(space))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QStyle_VisualAlignment(direction LayoutDirection, alignment AlignmentFlag) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyle_visualAlignment, uintptr(direction), uintptr(alignment))
	__r := int32(uint32(__r0))
	return AlignmentFlag(__r)
}

func QStyle_AlignedRect(direction LayoutDirection, alignment AlignmentFlag, size *QSize, rectangle *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_alignedRect, uintptr(direction), uintptr(alignment), uintptr(size.UnsafePointer()), uintptr(rectangle.UnsafePointer()))
	runtime.KeepAlive(size)
	runtime.KeepAlive(rectangle)
	return newownedQRect(__r)
}

func (this *QStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_layoutSpacing, this.h, uintptr(control1), uintptr(control2), uintptr(orientation), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStyle) CombinedLayoutSpacing(controls1 QSizePolicy__ControlType, controls2 QSizePolicy__ControlType, orientation Orientation) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_combinedLayoutSpacing, this.h, uintptr(controls1), uintptr(controls2), uintptr(orientation))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyle) Proxy() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QStyle_proxy, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func QStyle_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyle_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyle_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyle_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStyle) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStyle_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStyle) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyle) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStyle_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStyle) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyle_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStyle_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyle_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyle_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyle_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyle_SliderPositionFromValue2(min int, max int, val int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_sliderPositionFromValue2, uintptr(min), uintptr(max), uintptr(val), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QStyle_SliderValueFromPosition2(min int, max int, pos int, space int, upsideDown bool) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_sliderValueFromPosition2, uintptr(min), uintptr(max), uintptr(pos), uintptr(space), uintptr(b2i(upsideDown)))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QStyle) CombinedLayoutSpacing2(controls1 QSizePolicy__ControlType, controls2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_combinedLayoutSpacing2, this.h, uintptr(controls1), uintptr(controls2), uintptr(orientation), uintptr(option.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	return int(__r)
}

func (this *QStyle) CombinedLayoutSpacing3(controls1 QSizePolicy__ControlType, controls2 QSizePolicy__ControlType, orientation Orientation, option *QStyleOption, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStyle_combinedLayoutSpacing3, this.h, uintptr(controls1), uintptr(controls2), uintptr(orientation), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
	return int(__r)
}

func QStyle_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyle_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyle_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyle_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStyle) PolishDefault(widget *QWidget) {
	purego.SyscallN(addr_QStyle_polish_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStyle) UnpolishDefault(widget *QWidget) {
	purego.SyscallN(addr_QStyle_unpolish_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStyle) PolishWithApplicationDefault(application *QApplication) {
	purego.SyscallN(addr_QStyle_polish_default2, this.h, uintptr(application.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(application)
}

func (this *QStyle) UnpolishWithApplicationDefault(application *QApplication) {
	purego.SyscallN(addr_QStyle_unpolish_default2, this.h, uintptr(application.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(application)
}

func (this *QStyle) PolishWithPaletteDefault(palette *QPalette) {
	purego.SyscallN(addr_QStyle_polish_default3, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QStyle) ItemTextRectDefault(fm *QFontMetrics, r *QRect, flags int, enabled bool, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QStyle_itemTextRect_default, this.h, uintptr(fm.UnsafePointer()), uintptr(r.UnsafePointer()), uintptr(flags), uintptr(b2i(enabled)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fm)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QStyle) ItemPixmapRectDefault(r *QRect, flags int, pixmap *QPixmap) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyle_itemPixmapRect_default, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
	return newownedQRect(__r)
}

func (this *QStyle) DrawItemTextDefault(painter *QPainter, rect *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyle_drawItemText_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)), uintptr(textRole))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QStyle) DrawItemPixmapDefault(painter *QPainter, rect *QRect, alignment int, pixmap *QPixmap) {
	purego.SyscallN(addr_QStyle_drawItemPixmap_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pixmap)
}

func (this *QStyle) StandardPaletteDefault() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QStyle_standardPalette_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QStyle) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStyle_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStyle) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyle_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStyle) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyle_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStyle) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStyle_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyle) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStyle_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyle) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStyle_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyle) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStyle_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStyle) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStyle_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStyle) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStyle_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyle) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStyle_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyle) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStyle_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyle) Disconnect(handle int64) {
	pg_QStyle__disconnect(handle)
	unregisterCallback(handle)
}

func QStyle_virtIdx(slot int64) int32 {
	switch slot {
	case 6098275085281943360:
		return 0
	case -8697913287309583788:
		return 1
	case -1580195784239416471:
		return 2
	case -778167327566211030:
		return 3
	case 7058885155526130727:
		return 4
	case -8244607819117461060:
		return 5
	case -7621686327863150821:
		return 6
	case -3316981301817172730:
		return 7
	case 2017029407462380579:
		return 8
	case -5288619860977043091:
		return 9
	case -1870022100809264807:
		return 10
	case -4387834948099211678:
		return 11
	case -4874522738816254324:
		return 12
	case 1607084412591819503:
		return 13
	case -4633126903937721620:
		return 14
	case -5547720835653541585:
		return 15
	case 4263633720130086714:
		return 16
	case 265826155645830215:
		return 17
	case 4263634819641714925:
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

func QStyle_enable(h uintptr, idx int32) { pg_QStyle__set_override(h, idx) }

func NewQStyle() *QStyle {
	id := nextVirtID()
	o := newownedQStyle(pg_QStyle__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QStyle_virtIdx, QStyle_enable)
	return o
}

func (this *QStyle) OnDrawPrimitive(fn func(QStyle__PrimitiveElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6098275085281943360, func(pe int32, opt uintptr, p uintptr, w uintptr) { fn(QStyle__PrimitiveElement(pe), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(w))) }) // slot "drawPrimitive"
}

func (this *QStyle) OnDrawControl(fn func(QStyle__ControlElement, *QStyleOption, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -8697913287309583788, func(element int32, opt uintptr, p uintptr, w uintptr) { fn(QStyle__ControlElement(element), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(w))) }) // slot "drawControl"
}

func (this *QStyle) OnSubElementRect(fn func(QStyle__SubElement, *QStyleOption, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -1580195784239416471, func(subElement int32, option uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__SubElement(subElement), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "subElementRect"
}

func (this *QStyle) OnDrawComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPainter, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -778167327566211030, func(cc int32, opt uintptr, p uintptr, widget uintptr) { fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), UnsafeNewQPainter(unsafe.Pointer(p)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "drawComplexControl"
}

func (this *QStyle) OnHitTestComplexControl(fn func(QStyle__ComplexControl, *QStyleOptionComplex, *QPoint, *QWidget) QStyle__SubControl) {
	registerVirt(this.UnsafeGoID(), 7058885155526130727, func(cc int32, opt uintptr, pt uintptr, widget uintptr) int32 { return int32(fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), UnsafeNewQPoint(unsafe.Pointer(pt)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "hitTestComplexControl"
}

func (this *QStyle) OnSubControlRect(fn func(QStyle__ComplexControl, *QStyleOptionComplex, QStyle__SubControl, *QWidget) *QRect) {
	registerVirt(this.UnsafeGoID(), -8244607819117461060, func(cc int32, opt uintptr, sc int32, widget uintptr) uintptr { return uintptr(fn(QStyle__ComplexControl(cc), UnsafeNewQStyleOptionComplex(unsafe.Pointer(opt)), QStyle__SubControl(sc), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "subControlRect"
}

func (this *QStyle) OnPixelMetric(fn func(QStyle__PixelMetric, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7621686327863150821, func(metric int32, option uintptr, widget uintptr) int32 { return int32(fn(QStyle__PixelMetric(metric), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "pixelMetric"
}

func (this *QStyle) OnSizeFromContents(fn func(QStyle__ContentsType, *QStyleOption, *QSize, *QWidget) *QSize) {
	registerVirt(this.UnsafeGoID(), -3316981301817172730, func(ct int32, opt uintptr, contentsSize uintptr, w uintptr) uintptr { return uintptr(fn(QStyle__ContentsType(ct), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQSize(unsafe.Pointer(contentsSize)), UnsafeNewQWidget(unsafe.Pointer(w))).UnsafePointer()) }) // slot "sizeFromContents"
}

func (this *QStyle) OnStyleHint(fn func(QStyle__StyleHint, *QStyleOption, *QWidget, *QStyleHintReturn) int) {
	registerVirt(this.UnsafeGoID(), 2017029407462380579, func(stylehint int32, opt uintptr, widget uintptr, returnData uintptr) int32 { return int32(fn(QStyle__StyleHint(stylehint), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget)), UnsafeNewQStyleHintReturn(unsafe.Pointer(returnData)))) }) // slot "styleHint"
}

func (this *QStyle) OnStandardPixmap(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -5288619860977043091, func(standardPixmap int32, opt uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(standardPixmap), UnsafeNewQStyleOption(unsafe.Pointer(opt)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardPixmap"
}

func (this *QStyle) OnStandardIcon(fn func(QStyle__StandardPixmap, *QStyleOption, *QWidget) *QIcon) {
	registerVirt(this.UnsafeGoID(), -1870022100809264807, func(standardIcon int32, option uintptr, widget uintptr) uintptr { return uintptr(fn(QStyle__StandardPixmap(standardIcon), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))).UnsafePointer()) }) // slot "standardIcon"
}

func (this *QStyle) OnGeneratedIconPixmap(fn func(QIcon__Mode, *QPixmap, *QStyleOption) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -4387834948099211678, func(iconMode int32, pixmap uintptr, opt uintptr) uintptr { return uintptr(fn(QIcon__Mode(iconMode), UnsafeNewQPixmap(unsafe.Pointer(pixmap)), UnsafeNewQStyleOption(unsafe.Pointer(opt))).UnsafePointer()) }) // slot "generatedIconPixmap"
}

func (this *QStyle) OnLayoutSpacing(fn func(QSizePolicy__ControlType, QSizePolicy__ControlType, Orientation, *QStyleOption, *QWidget) int) {
	registerVirt(this.UnsafeGoID(), -4874522738816254324, func(control1 int32, control2 int32, orientation int32, option uintptr, widget uintptr) int32 { return int32(fn(QSizePolicy__ControlType(control1), QSizePolicy__ControlType(control2), Orientation(orientation), UnsafeNewQStyleOption(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget)))) }) // slot "layoutSpacing"
}

func (this *QStyle) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStyle) OnPolish(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -4633126903937721620, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "polish"
}

func (this *QStyle) OnUnpolish(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -5547720835653541585, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "unpolish"
}

func (this *QStyle) OnPolish2(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 4263633720130086714, func(application uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(application))) }) // slot "polish2"
}

func (this *QStyle) OnUnpolish2(fn func(*QApplication)) {
	registerVirt(this.UnsafeGoID(), 265826155645830215, func(application uintptr) { fn(UnsafeNewQApplication(unsafe.Pointer(application))) }) // slot "unpolish2"
}

func (this *QStyle) OnPolish3(fn func(*QPalette)) {
	registerVirt(this.UnsafeGoID(), 4263634819641714925, func(palette uintptr) { fn(UnsafeNewQPalette(unsafe.Pointer(palette))) }) // slot "polish3"
}

func (this *QStyle) OnItemTextRect(fn func(*QFontMetrics, *QRect, int, bool, string) *QRect) {
	registerVirt(this.UnsafeGoID(), 8974830244642632823, func(fm uintptr, r uintptr, flags int32, enabled int32, text uintptr, text_len int64) uintptr { return uintptr(fn(UnsafeNewQFontMetrics(unsafe.Pointer(fm)), UnsafeNewQRect(unsafe.Pointer(r)), int(flags), enabled != 0, pgStrNoFree(text, text_len)).UnsafePointer()) }) // slot "itemTextRect"
}

func (this *QStyle) OnItemPixmapRect(fn func(*QRect, int, *QPixmap) *QRect) {
	registerVirt(this.UnsafeGoID(), 3847657773274941773, func(r uintptr, flags int32, pixmap uintptr) uintptr { return uintptr(fn(UnsafeNewQRect(unsafe.Pointer(r)), int(flags), UnsafeNewQPixmap(unsafe.Pointer(pixmap))).UnsafePointer()) }) // slot "itemPixmapRect"
}

func (this *QStyle) OnDrawItemText(fn func(*QPainter, *QRect, int, *QPalette, bool, string, QPalette__ColorRole)) {
	registerVirt(this.UnsafeGoID(), 2315857538243927747, func(painter uintptr, rect uintptr, flags int32, pal uintptr, enabled int32, text uintptr, text_len int64, textRole int32) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(flags), UnsafeNewQPalette(unsafe.Pointer(pal)), enabled != 0, pgStrNoFree(text, text_len), QPalette__ColorRole(textRole)) }) // slot "drawItemText"
}

func (this *QStyle) OnDrawItemPixmap(fn func(*QPainter, *QRect, int, *QPixmap)) {
	registerVirt(this.UnsafeGoID(), -3740415202673381007, func(painter uintptr, rect uintptr, alignment int32, pixmap uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), int(alignment), UnsafeNewQPixmap(unsafe.Pointer(pixmap))) }) // slot "drawItemPixmap"
}

func (this *QStyle) OnStandardPalette(fn func() *QPalette) {
	registerVirt(this.UnsafeGoID(), 7278964765692882119, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "standardPalette"
}

func (this *QStyle) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStyle) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStyle) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStyle) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStyle) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStyle) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStyle) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStyle) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyle_delete(this.h)
	}
	this.UnsafeRelease()
}
