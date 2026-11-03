package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsPixmapItem_base func(uintptr) uintptr
	addr_QGraphicsPixmapItem_new uintptr
	addr_QGraphicsPixmapItem_new2 uintptr
	addr_QGraphicsPixmapItem_new3 uintptr
	addr_QGraphicsPixmapItem_new4 uintptr
	addr_QGraphicsPixmapItem_pixmap uintptr
	addr_QGraphicsPixmapItem_setPixmap uintptr
	addr_QGraphicsPixmapItem_transformationMode uintptr
	addr_QGraphicsPixmapItem_setTransformationMode uintptr
	addr_QGraphicsPixmapItem_offset uintptr
	addr_QGraphicsPixmapItem_setOffset uintptr
	pg_QGraphicsPixmapItem_setOffset2 func(uintptr, float64, float64)
	addr_QGraphicsPixmapItem_shapeMode uintptr
	addr_QGraphicsPixmapItem_setShapeMode uintptr
	addr_QGraphicsPixmapItem_updateMicroFocus uintptr
	addr_QGraphicsPixmapItem_addToIndex uintptr
	addr_QGraphicsPixmapItem_removeFromIndex uintptr
	addr_QGraphicsPixmapItem_prepareGeometryChange uintptr
	addr_QGraphicsPixmapItem_boundingRect_default uintptr
	addr_QGraphicsPixmapItem_shape_default uintptr
	addr_QGraphicsPixmapItem_contains_default uintptr
	addr_QGraphicsPixmapItem_paint_default uintptr
	addr_QGraphicsPixmapItem_isObscuredBy_default uintptr
	addr_QGraphicsPixmapItem_opaqueArea_default uintptr
	addr_QGraphicsPixmapItem_type_default uintptr
	addr_QGraphicsPixmapItem_supportsExtension_default uintptr
	addr_QGraphicsPixmapItem_setExtension_default uintptr
	addr_QGraphicsPixmapItem_extension_default uintptr
	addr_QGraphicsPixmapItem_advance_default uintptr
	addr_QGraphicsPixmapItem_collidesWithItem_default uintptr
	addr_QGraphicsPixmapItem_collidesWithPath_default uintptr
	addr_QGraphicsPixmapItem_sceneEventFilter_default uintptr
	addr_QGraphicsPixmapItem_sceneEvent_default uintptr
	addr_QGraphicsPixmapItem_contextMenuEvent_default uintptr
	addr_QGraphicsPixmapItem_dragEnterEvent_default uintptr
	addr_QGraphicsPixmapItem_dragLeaveEvent_default uintptr
	addr_QGraphicsPixmapItem_dragMoveEvent_default uintptr
	addr_QGraphicsPixmapItem_dropEvent_default uintptr
	addr_QGraphicsPixmapItem_focusInEvent_default uintptr
	addr_QGraphicsPixmapItem_focusOutEvent_default uintptr
	addr_QGraphicsPixmapItem_hoverEnterEvent_default uintptr
	addr_QGraphicsPixmapItem_hoverMoveEvent_default uintptr
	addr_QGraphicsPixmapItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsPixmapItem_keyPressEvent_default uintptr
	addr_QGraphicsPixmapItem_keyReleaseEvent_default uintptr
	addr_QGraphicsPixmapItem_mousePressEvent_default uintptr
	addr_QGraphicsPixmapItem_mouseMoveEvent_default uintptr
	addr_QGraphicsPixmapItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsPixmapItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsPixmapItem_wheelEvent_default uintptr
	addr_QGraphicsPixmapItem_inputMethodEvent_default uintptr
	addr_QGraphicsPixmapItem_inputMethodQuery_default uintptr
	addr_QGraphicsPixmapItem_itemChange_default uintptr
	pg_QGraphicsPixmapItem__set_override func(uintptr, int32)
	pg_QGraphicsPixmapItem_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsPixmapItem_base, "QGraphicsPixmapItem_base")
	symaddr(&addr_QGraphicsPixmapItem_new, "QGraphicsPixmapItem_new")
	symaddr(&addr_QGraphicsPixmapItem_new2, "QGraphicsPixmapItem_new2")
	symaddr(&addr_QGraphicsPixmapItem_new3, "QGraphicsPixmapItem_new3")
	symaddr(&addr_QGraphicsPixmapItem_new4, "QGraphicsPixmapItem_new4")
	symaddr(&addr_QGraphicsPixmapItem_pixmap, "QGraphicsPixmapItem_pixmap")
	symaddr(&addr_QGraphicsPixmapItem_setPixmap, "QGraphicsPixmapItem_setPixmap")
	symaddr(&addr_QGraphicsPixmapItem_transformationMode, "QGraphicsPixmapItem_transformationMode")
	symaddr(&addr_QGraphicsPixmapItem_setTransformationMode, "QGraphicsPixmapItem_setTransformationMode")
	symaddr(&addr_QGraphicsPixmapItem_offset, "QGraphicsPixmapItem_offset")
	symaddr(&addr_QGraphicsPixmapItem_setOffset, "QGraphicsPixmapItem_setOffset")
	reg(&pg_QGraphicsPixmapItem_setOffset2, "QGraphicsPixmapItem_setOffset2")
	symaddr(&addr_QGraphicsPixmapItem_shapeMode, "QGraphicsPixmapItem_shapeMode")
	symaddr(&addr_QGraphicsPixmapItem_setShapeMode, "QGraphicsPixmapItem_setShapeMode")
	symaddr(&addr_QGraphicsPixmapItem_updateMicroFocus, "QGraphicsPixmapItem_updateMicroFocus")
	symaddr(&addr_QGraphicsPixmapItem_addToIndex, "QGraphicsPixmapItem_addToIndex")
	symaddr(&addr_QGraphicsPixmapItem_removeFromIndex, "QGraphicsPixmapItem_removeFromIndex")
	symaddr(&addr_QGraphicsPixmapItem_prepareGeometryChange, "QGraphicsPixmapItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsPixmapItem_boundingRect_default, "QGraphicsPixmapItem_boundingRect_default")
	symaddr(&addr_QGraphicsPixmapItem_shape_default, "QGraphicsPixmapItem_shape_default")
	symaddr(&addr_QGraphicsPixmapItem_contains_default, "QGraphicsPixmapItem_contains_default")
	symaddr(&addr_QGraphicsPixmapItem_paint_default, "QGraphicsPixmapItem_paint_default")
	symaddr(&addr_QGraphicsPixmapItem_isObscuredBy_default, "QGraphicsPixmapItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsPixmapItem_opaqueArea_default, "QGraphicsPixmapItem_opaqueArea_default")
	symaddr(&addr_QGraphicsPixmapItem_type_default, "QGraphicsPixmapItem_type_default")
	symaddr(&addr_QGraphicsPixmapItem_supportsExtension_default, "QGraphicsPixmapItem_supportsExtension_default")
	symaddr(&addr_QGraphicsPixmapItem_setExtension_default, "QGraphicsPixmapItem_setExtension_default")
	symaddr(&addr_QGraphicsPixmapItem_extension_default, "QGraphicsPixmapItem_extension_default")
	symaddr(&addr_QGraphicsPixmapItem_advance_default, "QGraphicsPixmapItem_advance_default")
	symaddr(&addr_QGraphicsPixmapItem_collidesWithItem_default, "QGraphicsPixmapItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsPixmapItem_collidesWithPath_default, "QGraphicsPixmapItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsPixmapItem_sceneEventFilter_default, "QGraphicsPixmapItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsPixmapItem_sceneEvent_default, "QGraphicsPixmapItem_sceneEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_contextMenuEvent_default, "QGraphicsPixmapItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_dragEnterEvent_default, "QGraphicsPixmapItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_dragLeaveEvent_default, "QGraphicsPixmapItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_dragMoveEvent_default, "QGraphicsPixmapItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_dropEvent_default, "QGraphicsPixmapItem_dropEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_focusInEvent_default, "QGraphicsPixmapItem_focusInEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_focusOutEvent_default, "QGraphicsPixmapItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_hoverEnterEvent_default, "QGraphicsPixmapItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_hoverMoveEvent_default, "QGraphicsPixmapItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_hoverLeaveEvent_default, "QGraphicsPixmapItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_keyPressEvent_default, "QGraphicsPixmapItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_keyReleaseEvent_default, "QGraphicsPixmapItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_mousePressEvent_default, "QGraphicsPixmapItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_mouseMoveEvent_default, "QGraphicsPixmapItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_mouseReleaseEvent_default, "QGraphicsPixmapItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_mouseDoubleClickEvent_default, "QGraphicsPixmapItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_wheelEvent_default, "QGraphicsPixmapItem_wheelEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_inputMethodEvent_default, "QGraphicsPixmapItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsPixmapItem_inputMethodQuery_default, "QGraphicsPixmapItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsPixmapItem_itemChange_default, "QGraphicsPixmapItem_itemChange_default")
	reg(&pg_QGraphicsPixmapItem__set_override, "QGraphicsPixmapItem__set_override")
	reg(&pg_QGraphicsPixmapItem_delete, "QGraphicsPixmapItem_delete")
}

type QGraphicsPixmapItem struct {
	h uintptr
	*QGraphicsItem
}

func bindQGraphicsPixmapItem(h uintptr) *QGraphicsPixmapItem {
	if h == 0 {
		return nil
	}
	return &QGraphicsPixmapItem{h: h, QGraphicsItem: bindQGraphicsItem(pg_QGraphicsPixmapItem_base(h))}
}

func newownedQGraphicsPixmapItem(h uintptr) *QGraphicsPixmapItem {
	o := bindQGraphicsPixmapItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsPixmapItem).Delete)
	}
	return o
}

func (this *QGraphicsPixmapItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsPixmapItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsPixmapItem(h unsafe.Pointer) *QGraphicsPixmapItem { return bindQGraphicsPixmapItem(uintptr(h)) }
func UnsafeNewQGraphicsPixmapItemOwned(h unsafe.Pointer) *QGraphicsPixmapItem { return newownedQGraphicsPixmapItem(uintptr(h)) }

func (this *QGraphicsPixmapItem) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsItem != nil {
		this.QGraphicsItem.UnsafeRelease()
	}
}

func NewQGraphicsPixmapItem() *QGraphicsPixmapItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_new, uintptr(__goid))
	__o := newownedQGraphicsPixmapItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPixmapItem_virtIdx, QGraphicsPixmapItem_enable)
	return __o
}

func NewQGraphicsPixmapItem2(pixmap *QPixmap) *QGraphicsPixmapItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_new2, uintptr(__goid), uintptr(pixmap.UnsafePointer()))
	__o := newownedQGraphicsPixmapItem(__r)
	runtime.KeepAlive(pixmap)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPixmapItem_virtIdx, QGraphicsPixmapItem_enable)
	return __o
}

func NewQGraphicsPixmapItem3(parent *QGraphicsItem) *QGraphicsPixmapItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPixmapItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPixmapItem_virtIdx, QGraphicsPixmapItem_enable)
	return __o
}

func NewQGraphicsPixmapItem4(pixmap *QPixmap, parent *QGraphicsItem) *QGraphicsPixmapItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_new4, uintptr(__goid), uintptr(pixmap.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPixmapItem(__r)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPixmapItem_virtIdx, QGraphicsPixmapItem_enable)
	return __o
}

func (this *QGraphicsPixmapItem) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_pixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QGraphicsPixmapItem) SetPixmap(pixmap *QPixmap) {
	purego.SyscallN(addr_QGraphicsPixmapItem_setPixmap, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QGraphicsPixmapItem) TransformationMode() TransformationMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_transformationMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TransformationMode(__r)
}

func (this *QGraphicsPixmapItem) SetTransformationMode(mode TransformationMode) {
	purego.SyscallN(addr_QGraphicsPixmapItem_setTransformationMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) Offset() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_offset, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsPixmapItem) SetOffset(offset *QPointF) {
	purego.SyscallN(addr_QGraphicsPixmapItem_setOffset, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QGraphicsPixmapItem) SetOffset2(x float64, y float64) {
	pg_QGraphicsPixmapItem_setOffset2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) ShapeMode() QGraphicsPixmapItem__ShapeMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_shapeMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsPixmapItem__ShapeMode(__r)
}

func (this *QGraphicsPixmapItem) SetShapeMode(mode QGraphicsPixmapItem__ShapeMode) {
	purego.SyscallN(addr_QGraphicsPixmapItem_setShapeMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsPixmapItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsPixmapItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsPixmapItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsPixmapItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsPixmapItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPixmapItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsPixmapItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsPixmapItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsPixmapItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsPixmapItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPixmapItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsPixmapItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsPixmapItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsPixmapItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsPixmapItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsPixmapItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsPixmapItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPixmapItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsPixmapItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsPixmapItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPixmapItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPixmapItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsPixmapItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPixmapItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsPixmapItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPixmapItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsPixmapItem_virtIdx(slot int64) int32 {
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

func QGraphicsPixmapItem_enable(h uintptr, idx int32) { pg_QGraphicsPixmapItem__set_override(h, idx) }

func (this *QGraphicsPixmapItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsPixmapItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsPixmapItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsPixmapItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsPixmapItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsPixmapItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsPixmapItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsPixmapItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsPixmapItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsPixmapItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsPixmapItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsPixmapItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsPixmapItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsPixmapItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsPixmapItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsPixmapItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsPixmapItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsPixmapItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsPixmapItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsPixmapItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsPixmapItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsPixmapItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsPixmapItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsPixmapItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsPixmapItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsPixmapItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsPixmapItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsPixmapItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsPixmapItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsPixmapItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsPixmapItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsPixmapItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsPixmapItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsPixmapItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsPixmapItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsPixmapItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsPixmapItem_delete(this.h)
	}
	this.UnsafeRelease()
}
