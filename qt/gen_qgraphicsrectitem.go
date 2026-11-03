package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsRectItem_chain uintptr
	addr_QGraphicsRectItem_new uintptr
	addr_QGraphicsRectItem_new2 uintptr
	pg_QGraphicsRectItem_new3 func(int64, float64, float64, float64, float64) uintptr
	addr_QGraphicsRectItem_new4 uintptr
	addr_QGraphicsRectItem_new5 uintptr
	pg_QGraphicsRectItem_new6 func(int64, float64, float64, float64, float64, uintptr) uintptr
	addr_QGraphicsRectItem_rect uintptr
	addr_QGraphicsRectItem_setRect uintptr
	pg_QGraphicsRectItem_setRect2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsRectItem_updateMicroFocus uintptr
	addr_QGraphicsRectItem_addToIndex uintptr
	addr_QGraphicsRectItem_removeFromIndex uintptr
	addr_QGraphicsRectItem_prepareGeometryChange uintptr
	addr_QGraphicsRectItem_boundingRect_default uintptr
	addr_QGraphicsRectItem_shape_default uintptr
	addr_QGraphicsRectItem_contains_default uintptr
	addr_QGraphicsRectItem_paint_default uintptr
	addr_QGraphicsRectItem_isObscuredBy_default uintptr
	addr_QGraphicsRectItem_opaqueArea_default uintptr
	addr_QGraphicsRectItem_type_default uintptr
	addr_QGraphicsRectItem_supportsExtension_default uintptr
	addr_QGraphicsRectItem_setExtension_default uintptr
	addr_QGraphicsRectItem_extension_default uintptr
	addr_QGraphicsRectItem_advance_default uintptr
	addr_QGraphicsRectItem_collidesWithItem_default uintptr
	addr_QGraphicsRectItem_collidesWithPath_default uintptr
	addr_QGraphicsRectItem_sceneEventFilter_default uintptr
	addr_QGraphicsRectItem_sceneEvent_default uintptr
	addr_QGraphicsRectItem_contextMenuEvent_default uintptr
	addr_QGraphicsRectItem_dragEnterEvent_default uintptr
	addr_QGraphicsRectItem_dragLeaveEvent_default uintptr
	addr_QGraphicsRectItem_dragMoveEvent_default uintptr
	addr_QGraphicsRectItem_dropEvent_default uintptr
	addr_QGraphicsRectItem_focusInEvent_default uintptr
	addr_QGraphicsRectItem_focusOutEvent_default uintptr
	addr_QGraphicsRectItem_hoverEnterEvent_default uintptr
	addr_QGraphicsRectItem_hoverMoveEvent_default uintptr
	addr_QGraphicsRectItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsRectItem_keyPressEvent_default uintptr
	addr_QGraphicsRectItem_keyReleaseEvent_default uintptr
	addr_QGraphicsRectItem_mousePressEvent_default uintptr
	addr_QGraphicsRectItem_mouseMoveEvent_default uintptr
	addr_QGraphicsRectItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsRectItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsRectItem_wheelEvent_default uintptr
	addr_QGraphicsRectItem_inputMethodEvent_default uintptr
	addr_QGraphicsRectItem_inputMethodQuery_default uintptr
	addr_QGraphicsRectItem_itemChange_default uintptr
	pg_QGraphicsRectItem__set_override func(uintptr, int32)
	pg_QGraphicsRectItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsRectItem_chain, "QGraphicsRectItem_chain")
	symaddr(&addr_QGraphicsRectItem_new, "QGraphicsRectItem_new")
	symaddr(&addr_QGraphicsRectItem_new2, "QGraphicsRectItem_new2")
	reg(&pg_QGraphicsRectItem_new3, "QGraphicsRectItem_new3")
	symaddr(&addr_QGraphicsRectItem_new4, "QGraphicsRectItem_new4")
	symaddr(&addr_QGraphicsRectItem_new5, "QGraphicsRectItem_new5")
	reg(&pg_QGraphicsRectItem_new6, "QGraphicsRectItem_new6")
	symaddr(&addr_QGraphicsRectItem_rect, "QGraphicsRectItem_rect")
	symaddr(&addr_QGraphicsRectItem_setRect, "QGraphicsRectItem_setRect")
	reg(&pg_QGraphicsRectItem_setRect2, "QGraphicsRectItem_setRect2")
	symaddr(&addr_QGraphicsRectItem_updateMicroFocus, "QGraphicsRectItem_updateMicroFocus")
	symaddr(&addr_QGraphicsRectItem_addToIndex, "QGraphicsRectItem_addToIndex")
	symaddr(&addr_QGraphicsRectItem_removeFromIndex, "QGraphicsRectItem_removeFromIndex")
	symaddr(&addr_QGraphicsRectItem_prepareGeometryChange, "QGraphicsRectItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsRectItem_boundingRect_default, "QGraphicsRectItem_boundingRect_default")
	symaddr(&addr_QGraphicsRectItem_shape_default, "QGraphicsRectItem_shape_default")
	symaddr(&addr_QGraphicsRectItem_contains_default, "QGraphicsRectItem_contains_default")
	symaddr(&addr_QGraphicsRectItem_paint_default, "QGraphicsRectItem_paint_default")
	symaddr(&addr_QGraphicsRectItem_isObscuredBy_default, "QGraphicsRectItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsRectItem_opaqueArea_default, "QGraphicsRectItem_opaqueArea_default")
	symaddr(&addr_QGraphicsRectItem_type_default, "QGraphicsRectItem_type_default")
	symaddr(&addr_QGraphicsRectItem_supportsExtension_default, "QGraphicsRectItem_supportsExtension_default")
	symaddr(&addr_QGraphicsRectItem_setExtension_default, "QGraphicsRectItem_setExtension_default")
	symaddr(&addr_QGraphicsRectItem_extension_default, "QGraphicsRectItem_extension_default")
	symaddr(&addr_QGraphicsRectItem_advance_default, "QGraphicsRectItem_advance_default")
	symaddr(&addr_QGraphicsRectItem_collidesWithItem_default, "QGraphicsRectItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsRectItem_collidesWithPath_default, "QGraphicsRectItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsRectItem_sceneEventFilter_default, "QGraphicsRectItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsRectItem_sceneEvent_default, "QGraphicsRectItem_sceneEvent_default")
	symaddr(&addr_QGraphicsRectItem_contextMenuEvent_default, "QGraphicsRectItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsRectItem_dragEnterEvent_default, "QGraphicsRectItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsRectItem_dragLeaveEvent_default, "QGraphicsRectItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsRectItem_dragMoveEvent_default, "QGraphicsRectItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsRectItem_dropEvent_default, "QGraphicsRectItem_dropEvent_default")
	symaddr(&addr_QGraphicsRectItem_focusInEvent_default, "QGraphicsRectItem_focusInEvent_default")
	symaddr(&addr_QGraphicsRectItem_focusOutEvent_default, "QGraphicsRectItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsRectItem_hoverEnterEvent_default, "QGraphicsRectItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsRectItem_hoverMoveEvent_default, "QGraphicsRectItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsRectItem_hoverLeaveEvent_default, "QGraphicsRectItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsRectItem_keyPressEvent_default, "QGraphicsRectItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsRectItem_keyReleaseEvent_default, "QGraphicsRectItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsRectItem_mousePressEvent_default, "QGraphicsRectItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsRectItem_mouseMoveEvent_default, "QGraphicsRectItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsRectItem_mouseReleaseEvent_default, "QGraphicsRectItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsRectItem_mouseDoubleClickEvent_default, "QGraphicsRectItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsRectItem_wheelEvent_default, "QGraphicsRectItem_wheelEvent_default")
	symaddr(&addr_QGraphicsRectItem_inputMethodEvent_default, "QGraphicsRectItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsRectItem_inputMethodQuery_default, "QGraphicsRectItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsRectItem_itemChange_default, "QGraphicsRectItem_itemChange_default")
	reg(&pg_QGraphicsRectItem__set_override, "QGraphicsRectItem__set_override")
	reg(&pg_QGraphicsRectItem_delete, "QGraphicsRectItem_delete")
}

type QGraphicsRectItem struct {
	h uintptr
	*QAbstractGraphicsShapeItem
}

func bindQGraphicsRectItem(h uintptr) *QGraphicsRectItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsRectItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsRectItem{h: h, QAbstractGraphicsShapeItem: &QAbstractGraphicsShapeItem{h: __c[0], QGraphicsItem: &QGraphicsItem{h: __c[1]}}}
}

func newownedQGraphicsRectItem(h uintptr) *QGraphicsRectItem {
	o := bindQGraphicsRectItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsRectItem).Delete)
	}
	return o
}

func (this *QGraphicsRectItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsRectItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsRectItem(h unsafe.Pointer) *QGraphicsRectItem { return bindQGraphicsRectItem(uintptr(h)) }
func UnsafeNewQGraphicsRectItemOwned(h unsafe.Pointer) *QGraphicsRectItem { return newownedQGraphicsRectItem(uintptr(h)) }

func (this *QGraphicsRectItem) UnsafeRelease() {
	this.h = 0
	if this.QAbstractGraphicsShapeItem != nil {
		this.QAbstractGraphicsShapeItem.UnsafeRelease()
	}
}

func NewQGraphicsRectItem() *QGraphicsRectItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_new, uintptr(__goid))
	__o := newownedQGraphicsRectItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func NewQGraphicsRectItem2(rect *QRectF) *QGraphicsRectItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_new2, uintptr(__goid), uintptr(rect.UnsafePointer()))
	__o := newownedQGraphicsRectItem(__r)
	runtime.KeepAlive(rect)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func NewQGraphicsRectItem3(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsRectItem(pg_QGraphicsRectItem_new3(__goid, x, y, w, h))
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func NewQGraphicsRectItem4(parent *QGraphicsItem) *QGraphicsRectItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsRectItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func NewQGraphicsRectItem5(rect *QRectF, parent *QGraphicsItem) *QGraphicsRectItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_new5, uintptr(__goid), uintptr(rect.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsRectItem(__r)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func NewQGraphicsRectItem6(x float64, y float64, w float64, h float64, parent *QGraphicsItem) *QGraphicsRectItem {
	__goid := nextVirtID()
	__o := newownedQGraphicsRectItem(pg_QGraphicsRectItem_new6(__goid, x, y, w, h, uintptr(parent.UnsafePointer())))
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsRectItem_virtIdx, QGraphicsRectItem_enable)
	return __o
}

func (this *QGraphicsRectItem) Rect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsRectItem) SetRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsRectItem_setRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsRectItem) SetRect2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsRectItem_setRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsRectItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsRectItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsRectItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsRectItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsRectItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsRectItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsRectItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsRectItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsRectItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsRectItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsRectItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsRectItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsRectItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsRectItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsRectItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsRectItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsRectItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsRectItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsRectItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsRectItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsRectItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsRectItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsRectItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsRectItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsRectItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsRectItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsRectItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsRectItem_virtIdx(slot int64) int32 {
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

func QGraphicsRectItem_enable(h uintptr, idx int32) { pg_QGraphicsRectItem__set_override(h, idx) }

func (this *QGraphicsRectItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsRectItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsRectItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsRectItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsRectItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsRectItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsRectItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsRectItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsRectItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsRectItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsRectItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsRectItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsRectItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsRectItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsRectItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsRectItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsRectItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsRectItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsRectItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsRectItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsRectItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsRectItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsRectItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsRectItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsRectItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsRectItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsRectItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsRectItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsRectItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsRectItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsRectItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsRectItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsRectItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsRectItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsRectItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsRectItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsRectItem_delete(this.h)
	}
	this.UnsafeRelease()
}
