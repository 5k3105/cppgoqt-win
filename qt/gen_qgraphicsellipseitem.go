package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsEllipseItem_chain uintptr
	addr_QGraphicsEllipseItem_new uintptr
	addr_QGraphicsEllipseItem_new2 uintptr
	pg_QGraphicsEllipseItem_new3 func(int64, float64, float64, float64, float64) uintptr
	addr_QGraphicsEllipseItem_new4 uintptr
	addr_QGraphicsEllipseItem_new5 uintptr
	pg_QGraphicsEllipseItem_new6 func(int64, float64, float64, float64, float64, uintptr) uintptr
	addr_QGraphicsEllipseItem_rect uintptr
	addr_QGraphicsEllipseItem_setRect uintptr
	pg_QGraphicsEllipseItem_setRect2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsEllipseItem_startAngle uintptr
	addr_QGraphicsEllipseItem_setStartAngle uintptr
	addr_QGraphicsEllipseItem_spanAngle uintptr
	addr_QGraphicsEllipseItem_setSpanAngle uintptr
	addr_QGraphicsEllipseItem_updateMicroFocus uintptr
	addr_QGraphicsEllipseItem_addToIndex uintptr
	addr_QGraphicsEllipseItem_removeFromIndex uintptr
	addr_QGraphicsEllipseItem_prepareGeometryChange uintptr
	addr_QGraphicsEllipseItem_boundingRect_default uintptr
	addr_QGraphicsEllipseItem_shape_default uintptr
	addr_QGraphicsEllipseItem_contains_default uintptr
	addr_QGraphicsEllipseItem_paint_default uintptr
	addr_QGraphicsEllipseItem_isObscuredBy_default uintptr
	addr_QGraphicsEllipseItem_opaqueArea_default uintptr
	addr_QGraphicsEllipseItem_type_default uintptr
	addr_QGraphicsEllipseItem_supportsExtension_default uintptr
	addr_QGraphicsEllipseItem_setExtension_default uintptr
	addr_QGraphicsEllipseItem_extension_default uintptr
	addr_QGraphicsEllipseItem_advance_default uintptr
	addr_QGraphicsEllipseItem_collidesWithItem_default uintptr
	addr_QGraphicsEllipseItem_collidesWithPath_default uintptr
	addr_QGraphicsEllipseItem_sceneEventFilter_default uintptr
	addr_QGraphicsEllipseItem_sceneEvent_default uintptr
	addr_QGraphicsEllipseItem_contextMenuEvent_default uintptr
	addr_QGraphicsEllipseItem_dragEnterEvent_default uintptr
	addr_QGraphicsEllipseItem_dragLeaveEvent_default uintptr
	addr_QGraphicsEllipseItem_dragMoveEvent_default uintptr
	addr_QGraphicsEllipseItem_dropEvent_default uintptr
	addr_QGraphicsEllipseItem_focusInEvent_default uintptr
	addr_QGraphicsEllipseItem_focusOutEvent_default uintptr
	addr_QGraphicsEllipseItem_hoverEnterEvent_default uintptr
	addr_QGraphicsEllipseItem_hoverMoveEvent_default uintptr
	addr_QGraphicsEllipseItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsEllipseItem_keyPressEvent_default uintptr
	addr_QGraphicsEllipseItem_keyReleaseEvent_default uintptr
	addr_QGraphicsEllipseItem_mousePressEvent_default uintptr
	addr_QGraphicsEllipseItem_mouseMoveEvent_default uintptr
	addr_QGraphicsEllipseItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsEllipseItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsEllipseItem_wheelEvent_default uintptr
	addr_QGraphicsEllipseItem_inputMethodEvent_default uintptr
	addr_QGraphicsEllipseItem_inputMethodQuery_default uintptr
	addr_QGraphicsEllipseItem_itemChange_default uintptr
	pg_QGraphicsEllipseItem__set_override func(uintptr, int32)
	pg_QGraphicsEllipseItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsEllipseItem_chain, "QGraphicsEllipseItem_chain")
	symaddr(&addr_QGraphicsEllipseItem_new, "QGraphicsEllipseItem_new")
	symaddr(&addr_QGraphicsEllipseItem_new2, "QGraphicsEllipseItem_new2")
	reg(&pg_QGraphicsEllipseItem_new3, "QGraphicsEllipseItem_new3")
	symaddr(&addr_QGraphicsEllipseItem_new4, "QGraphicsEllipseItem_new4")
	symaddr(&addr_QGraphicsEllipseItem_new5, "QGraphicsEllipseItem_new5")
	reg(&pg_QGraphicsEllipseItem_new6, "QGraphicsEllipseItem_new6")
	symaddr(&addr_QGraphicsEllipseItem_rect, "QGraphicsEllipseItem_rect")
	symaddr(&addr_QGraphicsEllipseItem_setRect, "QGraphicsEllipseItem_setRect")
	reg(&pg_QGraphicsEllipseItem_setRect2, "QGraphicsEllipseItem_setRect2")
	symaddr(&addr_QGraphicsEllipseItem_startAngle, "QGraphicsEllipseItem_startAngle")
	symaddr(&addr_QGraphicsEllipseItem_setStartAngle, "QGraphicsEllipseItem_setStartAngle")
	symaddr(&addr_QGraphicsEllipseItem_spanAngle, "QGraphicsEllipseItem_spanAngle")
	symaddr(&addr_QGraphicsEllipseItem_setSpanAngle, "QGraphicsEllipseItem_setSpanAngle")
	symaddr(&addr_QGraphicsEllipseItem_updateMicroFocus, "QGraphicsEllipseItem_updateMicroFocus")
	symaddr(&addr_QGraphicsEllipseItem_addToIndex, "QGraphicsEllipseItem_addToIndex")
	symaddr(&addr_QGraphicsEllipseItem_removeFromIndex, "QGraphicsEllipseItem_removeFromIndex")
	symaddr(&addr_QGraphicsEllipseItem_prepareGeometryChange, "QGraphicsEllipseItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsEllipseItem_boundingRect_default, "QGraphicsEllipseItem_boundingRect_default")
	symaddr(&addr_QGraphicsEllipseItem_shape_default, "QGraphicsEllipseItem_shape_default")
	symaddr(&addr_QGraphicsEllipseItem_contains_default, "QGraphicsEllipseItem_contains_default")
	symaddr(&addr_QGraphicsEllipseItem_paint_default, "QGraphicsEllipseItem_paint_default")
	symaddr(&addr_QGraphicsEllipseItem_isObscuredBy_default, "QGraphicsEllipseItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsEllipseItem_opaqueArea_default, "QGraphicsEllipseItem_opaqueArea_default")
	symaddr(&addr_QGraphicsEllipseItem_type_default, "QGraphicsEllipseItem_type_default")
	symaddr(&addr_QGraphicsEllipseItem_supportsExtension_default, "QGraphicsEllipseItem_supportsExtension_default")
	symaddr(&addr_QGraphicsEllipseItem_setExtension_default, "QGraphicsEllipseItem_setExtension_default")
	symaddr(&addr_QGraphicsEllipseItem_extension_default, "QGraphicsEllipseItem_extension_default")
	symaddr(&addr_QGraphicsEllipseItem_advance_default, "QGraphicsEllipseItem_advance_default")
	symaddr(&addr_QGraphicsEllipseItem_collidesWithItem_default, "QGraphicsEllipseItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsEllipseItem_collidesWithPath_default, "QGraphicsEllipseItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsEllipseItem_sceneEventFilter_default, "QGraphicsEllipseItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsEllipseItem_sceneEvent_default, "QGraphicsEllipseItem_sceneEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_contextMenuEvent_default, "QGraphicsEllipseItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_dragEnterEvent_default, "QGraphicsEllipseItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_dragLeaveEvent_default, "QGraphicsEllipseItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_dragMoveEvent_default, "QGraphicsEllipseItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_dropEvent_default, "QGraphicsEllipseItem_dropEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_focusInEvent_default, "QGraphicsEllipseItem_focusInEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_focusOutEvent_default, "QGraphicsEllipseItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_hoverEnterEvent_default, "QGraphicsEllipseItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_hoverMoveEvent_default, "QGraphicsEllipseItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_hoverLeaveEvent_default, "QGraphicsEllipseItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_keyPressEvent_default, "QGraphicsEllipseItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_keyReleaseEvent_default, "QGraphicsEllipseItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_mousePressEvent_default, "QGraphicsEllipseItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_mouseMoveEvent_default, "QGraphicsEllipseItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_mouseReleaseEvent_default, "QGraphicsEllipseItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_mouseDoubleClickEvent_default, "QGraphicsEllipseItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_wheelEvent_default, "QGraphicsEllipseItem_wheelEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_inputMethodEvent_default, "QGraphicsEllipseItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsEllipseItem_inputMethodQuery_default, "QGraphicsEllipseItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsEllipseItem_itemChange_default, "QGraphicsEllipseItem_itemChange_default")
	reg(&pg_QGraphicsEllipseItem__set_override, "QGraphicsEllipseItem__set_override")
	reg(&pg_QGraphicsEllipseItem_delete, "QGraphicsEllipseItem_delete")
}

type QGraphicsEllipseItem struct {
	h uintptr
	*QAbstractGraphicsShapeItem
}

func bindQGraphicsEllipseItem(h uintptr) *QGraphicsEllipseItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsEllipseItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsEllipseItem{h: h, QAbstractGraphicsShapeItem: &QAbstractGraphicsShapeItem{h: __c[0], QGraphicsItem: &QGraphicsItem{h: __c[1]}}}
}

func newownedQGraphicsEllipseItem(h uintptr) *QGraphicsEllipseItem {
	o := bindQGraphicsEllipseItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsEllipseItem).Delete)
	}
	return o
}

func (this *QGraphicsEllipseItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsEllipseItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsEllipseItem(h unsafe.Pointer) *QGraphicsEllipseItem { return bindQGraphicsEllipseItem(uintptr(h)) }
func UnsafeNewQGraphicsEllipseItemOwned(h unsafe.Pointer) *QGraphicsEllipseItem { return newownedQGraphicsEllipseItem(uintptr(h)) }

func (this *QGraphicsEllipseItem) UnsafeRelease() {
	this.h = 0
	if this.QAbstractGraphicsShapeItem != nil {
		this.QAbstractGraphicsShapeItem.UnsafeRelease()
	}
}

func NewQGraphicsEllipseItem() *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_new, uintptr(__goid))
	__o := newownedQGraphicsEllipseItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func NewQGraphicsEllipseItem2(rect *QRectF) *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_new2, uintptr(__goid), uintptr(rect.UnsafePointer()))
	__o := newownedQGraphicsEllipseItem(__r)
	runtime.KeepAlive(rect)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func NewQGraphicsEllipseItem3(x float64, y float64, w float64, h float64) *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsEllipseItem(pg_QGraphicsEllipseItem_new3(__goid, x, y, w, h))
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func NewQGraphicsEllipseItem4(parent *QGraphicsItem) *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsEllipseItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func NewQGraphicsEllipseItem5(rect *QRectF, parent *QGraphicsItem) *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_new5, uintptr(__goid), uintptr(rect.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsEllipseItem(__r)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func NewQGraphicsEllipseItem6(x float64, y float64, w float64, h float64, parent *QGraphicsItem) *QGraphicsEllipseItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsEllipseItem(pg_QGraphicsEllipseItem_new6(__goid, x, y, w, h, uintptr(parent.UnsafePointer())))
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsEllipseItem_virtIdx, QGraphicsEllipseItem_enable)
	return __o
}

func (this *QGraphicsEllipseItem) Rect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsEllipseItem) SetRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsEllipseItem_setRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsEllipseItem) SetRect2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsEllipseItem_setRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) StartAngle() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_startAngle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsEllipseItem) SetStartAngle(angle int) {
	purego.SyscallN(addr_QGraphicsEllipseItem_setStartAngle, this.h, uintptr(angle))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) SpanAngle() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_spanAngle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsEllipseItem) SetSpanAngle(angle int) {
	purego.SyscallN(addr_QGraphicsEllipseItem_setSpanAngle, this.h, uintptr(angle))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsEllipseItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsEllipseItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsEllipseItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsEllipseItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsEllipseItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsEllipseItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsEllipseItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsEllipseItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsEllipseItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsEllipseItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsEllipseItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsEllipseItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsEllipseItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsEllipseItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsEllipseItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsEllipseItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsEllipseItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsEllipseItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsEllipseItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsEllipseItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsEllipseItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsEllipseItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsEllipseItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsEllipseItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsEllipseItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsEllipseItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsEllipseItem_virtIdx(slot int64) int32 {
	switch slot {
	case 2510008782785038831:
		return 0
	case 7938019069247444294:
		return 1
	case 5645366511944179768:
		return 2
	case 6083584261053556819:
		return 3
	case -7215030271914903671:
		return 4
	case 399370970023989435:
		return 5
	case -6371403872082420691:
		return 6
	case -5729166443651306418:
		return 7
	case -6730181576106171580:
		return 8
	case 2270616358835251226:
		return 9
	case -4830836341480073475:
		return 10
	case 6469736184362593173:
		return 11
	case 6636575182718753853:
		return 12
	case -4584203218241215725:
		return 13
	case -2846371507838587071:
		return 14
	case -6698043222227006015:
		return 15
	case 206122715919917589:
		return 16
	case -3139904158373725006:
		return 17
	case -1781545137352745978:
		return 18
	case 9107058682754591716:
		return 19
	case 975043767033552552:
		return 20
	case -385596010576048993:
		return 21
	case 6932941943906681679:
		return 22
	case -2745191522144316396:
		return 23
	case 7944408482203603984:
		return 24
	case 1023791977679793185:
		return 25
	case 772598475950050299:
		return 26
	case -332090898078543319:
		return 27
	case -4255494004099806581:
		return 28
	case -4415920049297795213:
		return 29
	case -4029073905697023887:
		return 30
	case 1919988340934511316:
		return 31
	case 5438819661942843222:
		return 32
	case 7345853097340709638:
		return 33
	case -1888572168794878830:
		return 34
	}
	return -1
}

func QGraphicsEllipseItem_enable(h uintptr, idx int32) { pg_QGraphicsEllipseItem__set_override(h, idx) }

func (this *QGraphicsEllipseItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsEllipseItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsEllipseItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsEllipseItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsEllipseItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsEllipseItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsEllipseItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsEllipseItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsEllipseItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsEllipseItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsEllipseItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsEllipseItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsEllipseItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsEllipseItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsEllipseItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsEllipseItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsEllipseItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsEllipseItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsEllipseItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsEllipseItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsEllipseItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsEllipseItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsEllipseItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsEllipseItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsEllipseItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsEllipseItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsEllipseItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsEllipseItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsEllipseItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsEllipseItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsEllipseItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsEllipseItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsEllipseItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsEllipseItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsEllipseItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsEllipseItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsEllipseItem_delete(this.h)
	}
	this.UnsafeRelease()
}
