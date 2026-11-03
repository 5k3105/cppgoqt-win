package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsLineItem_base func(uintptr) uintptr
	addr_QGraphicsLineItem_new uintptr
	addr_QGraphicsLineItem_new2 uintptr
	pg_QGraphicsLineItem_new3 func(int64, float64, float64, float64, float64) uintptr
	addr_QGraphicsLineItem_new4 uintptr
	addr_QGraphicsLineItem_new5 uintptr
	pg_QGraphicsLineItem_new6 func(int64, float64, float64, float64, float64, uintptr) uintptr
	addr_QGraphicsLineItem_pen uintptr
	addr_QGraphicsLineItem_setPen uintptr
	addr_QGraphicsLineItem_line uintptr
	addr_QGraphicsLineItem_setLine uintptr
	pg_QGraphicsLineItem_setLine2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsLineItem_updateMicroFocus uintptr
	addr_QGraphicsLineItem_addToIndex uintptr
	addr_QGraphicsLineItem_removeFromIndex uintptr
	addr_QGraphicsLineItem_prepareGeometryChange uintptr
	addr_QGraphicsLineItem_boundingRect_default uintptr
	addr_QGraphicsLineItem_shape_default uintptr
	addr_QGraphicsLineItem_contains_default uintptr
	addr_QGraphicsLineItem_paint_default uintptr
	addr_QGraphicsLineItem_isObscuredBy_default uintptr
	addr_QGraphicsLineItem_opaqueArea_default uintptr
	addr_QGraphicsLineItem_type_default uintptr
	addr_QGraphicsLineItem_supportsExtension_default uintptr
	addr_QGraphicsLineItem_setExtension_default uintptr
	addr_QGraphicsLineItem_extension_default uintptr
	addr_QGraphicsLineItem_advance_default uintptr
	addr_QGraphicsLineItem_collidesWithItem_default uintptr
	addr_QGraphicsLineItem_collidesWithPath_default uintptr
	addr_QGraphicsLineItem_sceneEventFilter_default uintptr
	addr_QGraphicsLineItem_sceneEvent_default uintptr
	addr_QGraphicsLineItem_contextMenuEvent_default uintptr
	addr_QGraphicsLineItem_dragEnterEvent_default uintptr
	addr_QGraphicsLineItem_dragLeaveEvent_default uintptr
	addr_QGraphicsLineItem_dragMoveEvent_default uintptr
	addr_QGraphicsLineItem_dropEvent_default uintptr
	addr_QGraphicsLineItem_focusInEvent_default uintptr
	addr_QGraphicsLineItem_focusOutEvent_default uintptr
	addr_QGraphicsLineItem_hoverEnterEvent_default uintptr
	addr_QGraphicsLineItem_hoverMoveEvent_default uintptr
	addr_QGraphicsLineItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsLineItem_keyPressEvent_default uintptr
	addr_QGraphicsLineItem_keyReleaseEvent_default uintptr
	addr_QGraphicsLineItem_mousePressEvent_default uintptr
	addr_QGraphicsLineItem_mouseMoveEvent_default uintptr
	addr_QGraphicsLineItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsLineItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsLineItem_wheelEvent_default uintptr
	addr_QGraphicsLineItem_inputMethodEvent_default uintptr
	addr_QGraphicsLineItem_inputMethodQuery_default uintptr
	addr_QGraphicsLineItem_itemChange_default uintptr
	pg_QGraphicsLineItem__set_override func(uintptr, int32)
	pg_QGraphicsLineItem_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsLineItem_base, "QGraphicsLineItem_base")
	symaddr(&addr_QGraphicsLineItem_new, "QGraphicsLineItem_new")
	symaddr(&addr_QGraphicsLineItem_new2, "QGraphicsLineItem_new2")
	reg(&pg_QGraphicsLineItem_new3, "QGraphicsLineItem_new3")
	symaddr(&addr_QGraphicsLineItem_new4, "QGraphicsLineItem_new4")
	symaddr(&addr_QGraphicsLineItem_new5, "QGraphicsLineItem_new5")
	reg(&pg_QGraphicsLineItem_new6, "QGraphicsLineItem_new6")
	symaddr(&addr_QGraphicsLineItem_pen, "QGraphicsLineItem_pen")
	symaddr(&addr_QGraphicsLineItem_setPen, "QGraphicsLineItem_setPen")
	symaddr(&addr_QGraphicsLineItem_line, "QGraphicsLineItem_line")
	symaddr(&addr_QGraphicsLineItem_setLine, "QGraphicsLineItem_setLine")
	reg(&pg_QGraphicsLineItem_setLine2, "QGraphicsLineItem_setLine2")
	symaddr(&addr_QGraphicsLineItem_updateMicroFocus, "QGraphicsLineItem_updateMicroFocus")
	symaddr(&addr_QGraphicsLineItem_addToIndex, "QGraphicsLineItem_addToIndex")
	symaddr(&addr_QGraphicsLineItem_removeFromIndex, "QGraphicsLineItem_removeFromIndex")
	symaddr(&addr_QGraphicsLineItem_prepareGeometryChange, "QGraphicsLineItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsLineItem_boundingRect_default, "QGraphicsLineItem_boundingRect_default")
	symaddr(&addr_QGraphicsLineItem_shape_default, "QGraphicsLineItem_shape_default")
	symaddr(&addr_QGraphicsLineItem_contains_default, "QGraphicsLineItem_contains_default")
	symaddr(&addr_QGraphicsLineItem_paint_default, "QGraphicsLineItem_paint_default")
	symaddr(&addr_QGraphicsLineItem_isObscuredBy_default, "QGraphicsLineItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsLineItem_opaqueArea_default, "QGraphicsLineItem_opaqueArea_default")
	symaddr(&addr_QGraphicsLineItem_type_default, "QGraphicsLineItem_type_default")
	symaddr(&addr_QGraphicsLineItem_supportsExtension_default, "QGraphicsLineItem_supportsExtension_default")
	symaddr(&addr_QGraphicsLineItem_setExtension_default, "QGraphicsLineItem_setExtension_default")
	symaddr(&addr_QGraphicsLineItem_extension_default, "QGraphicsLineItem_extension_default")
	symaddr(&addr_QGraphicsLineItem_advance_default, "QGraphicsLineItem_advance_default")
	symaddr(&addr_QGraphicsLineItem_collidesWithItem_default, "QGraphicsLineItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsLineItem_collidesWithPath_default, "QGraphicsLineItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsLineItem_sceneEventFilter_default, "QGraphicsLineItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsLineItem_sceneEvent_default, "QGraphicsLineItem_sceneEvent_default")
	symaddr(&addr_QGraphicsLineItem_contextMenuEvent_default, "QGraphicsLineItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsLineItem_dragEnterEvent_default, "QGraphicsLineItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsLineItem_dragLeaveEvent_default, "QGraphicsLineItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsLineItem_dragMoveEvent_default, "QGraphicsLineItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsLineItem_dropEvent_default, "QGraphicsLineItem_dropEvent_default")
	symaddr(&addr_QGraphicsLineItem_focusInEvent_default, "QGraphicsLineItem_focusInEvent_default")
	symaddr(&addr_QGraphicsLineItem_focusOutEvent_default, "QGraphicsLineItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsLineItem_hoverEnterEvent_default, "QGraphicsLineItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsLineItem_hoverMoveEvent_default, "QGraphicsLineItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsLineItem_hoverLeaveEvent_default, "QGraphicsLineItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsLineItem_keyPressEvent_default, "QGraphicsLineItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsLineItem_keyReleaseEvent_default, "QGraphicsLineItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsLineItem_mousePressEvent_default, "QGraphicsLineItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsLineItem_mouseMoveEvent_default, "QGraphicsLineItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsLineItem_mouseReleaseEvent_default, "QGraphicsLineItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsLineItem_mouseDoubleClickEvent_default, "QGraphicsLineItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsLineItem_wheelEvent_default, "QGraphicsLineItem_wheelEvent_default")
	symaddr(&addr_QGraphicsLineItem_inputMethodEvent_default, "QGraphicsLineItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsLineItem_inputMethodQuery_default, "QGraphicsLineItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsLineItem_itemChange_default, "QGraphicsLineItem_itemChange_default")
	reg(&pg_QGraphicsLineItem__set_override, "QGraphicsLineItem__set_override")
	reg(&pg_QGraphicsLineItem_delete, "QGraphicsLineItem_delete")
}

type QGraphicsLineItem struct {
	h uintptr
	*QGraphicsItem
}

func bindQGraphicsLineItem(h uintptr) *QGraphicsLineItem {
	if h == 0 {
		return nil
	}
	return &QGraphicsLineItem{h: h, QGraphicsItem: bindQGraphicsItem(pg_QGraphicsLineItem_base(h))}
}

func newownedQGraphicsLineItem(h uintptr) *QGraphicsLineItem {
	o := bindQGraphicsLineItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsLineItem).Delete)
	}
	return o
}

func (this *QGraphicsLineItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsLineItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsLineItem(h unsafe.Pointer) *QGraphicsLineItem { return bindQGraphicsLineItem(uintptr(h)) }
func UnsafeNewQGraphicsLineItemOwned(h unsafe.Pointer) *QGraphicsLineItem { return newownedQGraphicsLineItem(uintptr(h)) }

func (this *QGraphicsLineItem) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsItem != nil {
		this.QGraphicsItem.UnsafeRelease()
	}
}

func NewQGraphicsLineItem() *QGraphicsLineItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_new, uintptr(__goid))
	__o := newownedQGraphicsLineItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func NewQGraphicsLineItem2(line *QLineF) *QGraphicsLineItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_new2, uintptr(__goid), uintptr(line.UnsafePointer()))
	__o := newownedQGraphicsLineItem(__r)
	runtime.KeepAlive(line)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func NewQGraphicsLineItem3(x1 float64, y1 float64, x2 float64, y2 float64) *QGraphicsLineItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsLineItem(pg_QGraphicsLineItem_new3(__goid, x1, y1, x2, y2))
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func NewQGraphicsLineItem4(parent *QGraphicsItem) *QGraphicsLineItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsLineItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func NewQGraphicsLineItem5(line *QLineF, parent *QGraphicsItem) *QGraphicsLineItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_new5, uintptr(__goid), uintptr(line.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsLineItem(__r)
	runtime.KeepAlive(line)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func NewQGraphicsLineItem6(x1 float64, y1 float64, x2 float64, y2 float64, parent *QGraphicsItem) *QGraphicsLineItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsLineItem(pg_QGraphicsLineItem_new6(__goid, x1, y1, x2, y2, uintptr(parent.UnsafePointer())))
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLineItem_virtIdx, QGraphicsLineItem_enable)
	return __o
}

func (this *QGraphicsLineItem) Pen() *QPen {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_pen, this.h)
	runtime.KeepAlive(this)
	return newownedQPen(__r)
}

func (this *QGraphicsLineItem) SetPen(pen *QPen) {
	purego.SyscallN(addr_QGraphicsLineItem_setPen, this.h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
}

func (this *QGraphicsLineItem) Line() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_line, this.h)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QGraphicsLineItem) SetLine(line *QLineF) {
	purego.SyscallN(addr_QGraphicsLineItem_setLine, this.h, uintptr(line.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
}

func (this *QGraphicsLineItem) SetLine2(x1 float64, y1 float64, x2 float64, y2 float64) {
	pg_QGraphicsLineItem_setLine2(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsLineItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsLineItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsLineItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsLineItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsLineItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsLineItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsLineItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsLineItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsLineItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsLineItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsLineItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsLineItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsLineItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsLineItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsLineItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsLineItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsLineItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLineItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsLineItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsLineItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsLineItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLineItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsLineItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsLineItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsLineItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsLineItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLineItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsLineItem_virtIdx(slot int64) int32 {
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

func QGraphicsLineItem_enable(h uintptr, idx int32) { pg_QGraphicsLineItem__set_override(h, idx) }

func (this *QGraphicsLineItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsLineItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsLineItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsLineItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsLineItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsLineItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsLineItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsLineItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsLineItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsLineItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsLineItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsLineItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsLineItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsLineItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsLineItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsLineItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsLineItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsLineItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsLineItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsLineItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsLineItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsLineItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsLineItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsLineItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsLineItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsLineItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsLineItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsLineItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsLineItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsLineItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsLineItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsLineItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsLineItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsLineItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsLineItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsLineItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsLineItem_delete(this.h)
	}
	this.UnsafeRelease()
}
