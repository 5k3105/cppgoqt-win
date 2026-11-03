package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsPolygonItem_chain uintptr
	addr_QGraphicsPolygonItem_new uintptr
	addr_QGraphicsPolygonItem_new2 uintptr
	addr_QGraphicsPolygonItem_new3 uintptr
	addr_QGraphicsPolygonItem_new4 uintptr
	addr_QGraphicsPolygonItem_polygon uintptr
	addr_QGraphicsPolygonItem_setPolygon uintptr
	addr_QGraphicsPolygonItem_fillRule uintptr
	addr_QGraphicsPolygonItem_setFillRule uintptr
	addr_QGraphicsPolygonItem_updateMicroFocus uintptr
	addr_QGraphicsPolygonItem_addToIndex uintptr
	addr_QGraphicsPolygonItem_removeFromIndex uintptr
	addr_QGraphicsPolygonItem_prepareGeometryChange uintptr
	addr_QGraphicsPolygonItem_boundingRect_default uintptr
	addr_QGraphicsPolygonItem_shape_default uintptr
	addr_QGraphicsPolygonItem_contains_default uintptr
	addr_QGraphicsPolygonItem_paint_default uintptr
	addr_QGraphicsPolygonItem_isObscuredBy_default uintptr
	addr_QGraphicsPolygonItem_opaqueArea_default uintptr
	addr_QGraphicsPolygonItem_type_default uintptr
	addr_QGraphicsPolygonItem_supportsExtension_default uintptr
	addr_QGraphicsPolygonItem_setExtension_default uintptr
	addr_QGraphicsPolygonItem_extension_default uintptr
	addr_QGraphicsPolygonItem_advance_default uintptr
	addr_QGraphicsPolygonItem_collidesWithItem_default uintptr
	addr_QGraphicsPolygonItem_collidesWithPath_default uintptr
	addr_QGraphicsPolygonItem_sceneEventFilter_default uintptr
	addr_QGraphicsPolygonItem_sceneEvent_default uintptr
	addr_QGraphicsPolygonItem_contextMenuEvent_default uintptr
	addr_QGraphicsPolygonItem_dragEnterEvent_default uintptr
	addr_QGraphicsPolygonItem_dragLeaveEvent_default uintptr
	addr_QGraphicsPolygonItem_dragMoveEvent_default uintptr
	addr_QGraphicsPolygonItem_dropEvent_default uintptr
	addr_QGraphicsPolygonItem_focusInEvent_default uintptr
	addr_QGraphicsPolygonItem_focusOutEvent_default uintptr
	addr_QGraphicsPolygonItem_hoverEnterEvent_default uintptr
	addr_QGraphicsPolygonItem_hoverMoveEvent_default uintptr
	addr_QGraphicsPolygonItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsPolygonItem_keyPressEvent_default uintptr
	addr_QGraphicsPolygonItem_keyReleaseEvent_default uintptr
	addr_QGraphicsPolygonItem_mousePressEvent_default uintptr
	addr_QGraphicsPolygonItem_mouseMoveEvent_default uintptr
	addr_QGraphicsPolygonItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsPolygonItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsPolygonItem_wheelEvent_default uintptr
	addr_QGraphicsPolygonItem_inputMethodEvent_default uintptr
	addr_QGraphicsPolygonItem_inputMethodQuery_default uintptr
	addr_QGraphicsPolygonItem_itemChange_default uintptr
	pg_QGraphicsPolygonItem__set_override func(uintptr, int32)
	pg_QGraphicsPolygonItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsPolygonItem_chain, "QGraphicsPolygonItem_chain")
	symaddr(&addr_QGraphicsPolygonItem_new, "QGraphicsPolygonItem_new")
	symaddr(&addr_QGraphicsPolygonItem_new2, "QGraphicsPolygonItem_new2")
	symaddr(&addr_QGraphicsPolygonItem_new3, "QGraphicsPolygonItem_new3")
	symaddr(&addr_QGraphicsPolygonItem_new4, "QGraphicsPolygonItem_new4")
	symaddr(&addr_QGraphicsPolygonItem_polygon, "QGraphicsPolygonItem_polygon")
	symaddr(&addr_QGraphicsPolygonItem_setPolygon, "QGraphicsPolygonItem_setPolygon")
	symaddr(&addr_QGraphicsPolygonItem_fillRule, "QGraphicsPolygonItem_fillRule")
	symaddr(&addr_QGraphicsPolygonItem_setFillRule, "QGraphicsPolygonItem_setFillRule")
	symaddr(&addr_QGraphicsPolygonItem_updateMicroFocus, "QGraphicsPolygonItem_updateMicroFocus")
	symaddr(&addr_QGraphicsPolygonItem_addToIndex, "QGraphicsPolygonItem_addToIndex")
	symaddr(&addr_QGraphicsPolygonItem_removeFromIndex, "QGraphicsPolygonItem_removeFromIndex")
	symaddr(&addr_QGraphicsPolygonItem_prepareGeometryChange, "QGraphicsPolygonItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsPolygonItem_boundingRect_default, "QGraphicsPolygonItem_boundingRect_default")
	symaddr(&addr_QGraphicsPolygonItem_shape_default, "QGraphicsPolygonItem_shape_default")
	symaddr(&addr_QGraphicsPolygonItem_contains_default, "QGraphicsPolygonItem_contains_default")
	symaddr(&addr_QGraphicsPolygonItem_paint_default, "QGraphicsPolygonItem_paint_default")
	symaddr(&addr_QGraphicsPolygonItem_isObscuredBy_default, "QGraphicsPolygonItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsPolygonItem_opaqueArea_default, "QGraphicsPolygonItem_opaqueArea_default")
	symaddr(&addr_QGraphicsPolygonItem_type_default, "QGraphicsPolygonItem_type_default")
	symaddr(&addr_QGraphicsPolygonItem_supportsExtension_default, "QGraphicsPolygonItem_supportsExtension_default")
	symaddr(&addr_QGraphicsPolygonItem_setExtension_default, "QGraphicsPolygonItem_setExtension_default")
	symaddr(&addr_QGraphicsPolygonItem_extension_default, "QGraphicsPolygonItem_extension_default")
	symaddr(&addr_QGraphicsPolygonItem_advance_default, "QGraphicsPolygonItem_advance_default")
	symaddr(&addr_QGraphicsPolygonItem_collidesWithItem_default, "QGraphicsPolygonItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsPolygonItem_collidesWithPath_default, "QGraphicsPolygonItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsPolygonItem_sceneEventFilter_default, "QGraphicsPolygonItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsPolygonItem_sceneEvent_default, "QGraphicsPolygonItem_sceneEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_contextMenuEvent_default, "QGraphicsPolygonItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_dragEnterEvent_default, "QGraphicsPolygonItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_dragLeaveEvent_default, "QGraphicsPolygonItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_dragMoveEvent_default, "QGraphicsPolygonItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_dropEvent_default, "QGraphicsPolygonItem_dropEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_focusInEvent_default, "QGraphicsPolygonItem_focusInEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_focusOutEvent_default, "QGraphicsPolygonItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_hoverEnterEvent_default, "QGraphicsPolygonItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_hoverMoveEvent_default, "QGraphicsPolygonItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_hoverLeaveEvent_default, "QGraphicsPolygonItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_keyPressEvent_default, "QGraphicsPolygonItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_keyReleaseEvent_default, "QGraphicsPolygonItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_mousePressEvent_default, "QGraphicsPolygonItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_mouseMoveEvent_default, "QGraphicsPolygonItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_mouseReleaseEvent_default, "QGraphicsPolygonItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_mouseDoubleClickEvent_default, "QGraphicsPolygonItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_wheelEvent_default, "QGraphicsPolygonItem_wheelEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_inputMethodEvent_default, "QGraphicsPolygonItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsPolygonItem_inputMethodQuery_default, "QGraphicsPolygonItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsPolygonItem_itemChange_default, "QGraphicsPolygonItem_itemChange_default")
	reg(&pg_QGraphicsPolygonItem__set_override, "QGraphicsPolygonItem__set_override")
	reg(&pg_QGraphicsPolygonItem_delete, "QGraphicsPolygonItem_delete")
}

type QGraphicsPolygonItem struct {
	h uintptr
	*QAbstractGraphicsShapeItem
}

func bindQGraphicsPolygonItem(h uintptr) *QGraphicsPolygonItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsPolygonItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsPolygonItem{h: h, QAbstractGraphicsShapeItem: &QAbstractGraphicsShapeItem{h: __c[0], QGraphicsItem: &QGraphicsItem{h: __c[1]}}}
}

func newownedQGraphicsPolygonItem(h uintptr) *QGraphicsPolygonItem {
	o := bindQGraphicsPolygonItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsPolygonItem).Delete)
	}
	return o
}

func (this *QGraphicsPolygonItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsPolygonItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsPolygonItem(h unsafe.Pointer) *QGraphicsPolygonItem { return bindQGraphicsPolygonItem(uintptr(h)) }
func UnsafeNewQGraphicsPolygonItemOwned(h unsafe.Pointer) *QGraphicsPolygonItem { return newownedQGraphicsPolygonItem(uintptr(h)) }

func (this *QGraphicsPolygonItem) UnsafeRelease() {
	this.h = 0
	if this.QAbstractGraphicsShapeItem != nil {
		this.QAbstractGraphicsShapeItem.UnsafeRelease()
	}
}

func NewQGraphicsPolygonItem() *QGraphicsPolygonItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_new, uintptr(__goid))
	__o := newownedQGraphicsPolygonItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPolygonItem_virtIdx, QGraphicsPolygonItem_enable)
	return __o
}

func NewQGraphicsPolygonItem2(polygon *QPolygonF) *QGraphicsPolygonItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_new2, uintptr(__goid), uintptr(polygon.UnsafePointer()))
	__o := newownedQGraphicsPolygonItem(__r)
	runtime.KeepAlive(polygon)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPolygonItem_virtIdx, QGraphicsPolygonItem_enable)
	return __o
}

func NewQGraphicsPolygonItem3(parent *QGraphicsItem) *QGraphicsPolygonItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPolygonItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPolygonItem_virtIdx, QGraphicsPolygonItem_enable)
	return __o
}

func NewQGraphicsPolygonItem4(polygon *QPolygonF, parent *QGraphicsItem) *QGraphicsPolygonItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_new4, uintptr(__goid), uintptr(polygon.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPolygonItem(__r)
	runtime.KeepAlive(polygon)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPolygonItem_virtIdx, QGraphicsPolygonItem_enable)
	return __o
}

func (this *QGraphicsPolygonItem) Polygon() *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_polygon, this.h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsPolygonItem) SetPolygon(polygon *QPolygonF) {
	purego.SyscallN(addr_QGraphicsPolygonItem_setPolygon, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QGraphicsPolygonItem) FillRule() FillRule {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_fillRule, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return FillRule(__r)
}

func (this *QGraphicsPolygonItem) SetFillRule(rule FillRule) {
	purego.SyscallN(addr_QGraphicsPolygonItem_setFillRule, this.h, uintptr(rule))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsPolygonItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsPolygonItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsPolygonItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsPolygonItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsPolygonItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPolygonItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsPolygonItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsPolygonItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsPolygonItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsPolygonItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPolygonItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsPolygonItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsPolygonItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsPolygonItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsPolygonItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsPolygonItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsPolygonItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPolygonItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsPolygonItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsPolygonItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPolygonItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPolygonItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsPolygonItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPolygonItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsPolygonItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPolygonItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsPolygonItem_virtIdx(slot int64) int32 {
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

func QGraphicsPolygonItem_enable(h uintptr, idx int32) { pg_QGraphicsPolygonItem__set_override(h, idx) }

func (this *QGraphicsPolygonItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsPolygonItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsPolygonItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsPolygonItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsPolygonItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsPolygonItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsPolygonItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsPolygonItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsPolygonItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsPolygonItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsPolygonItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsPolygonItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsPolygonItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsPolygonItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsPolygonItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsPolygonItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsPolygonItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsPolygonItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsPolygonItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsPolygonItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsPolygonItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsPolygonItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsPolygonItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsPolygonItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsPolygonItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsPolygonItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsPolygonItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsPolygonItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsPolygonItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsPolygonItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsPolygonItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsPolygonItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsPolygonItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsPolygonItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsPolygonItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsPolygonItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsPolygonItem_delete(this.h)
	}
	this.UnsafeRelease()
}
