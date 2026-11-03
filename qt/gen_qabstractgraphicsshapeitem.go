package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAbstractGraphicsShapeItem_base func(uintptr) uintptr
	addr_QAbstractGraphicsShapeItem_pen uintptr
	addr_QAbstractGraphicsShapeItem_setPen uintptr
	addr_QAbstractGraphicsShapeItem_brush uintptr
	addr_QAbstractGraphicsShapeItem_setBrush uintptr
	addr_QAbstractGraphicsShapeItem_updateMicroFocus uintptr
	addr_QAbstractGraphicsShapeItem_addToIndex uintptr
	addr_QAbstractGraphicsShapeItem_removeFromIndex uintptr
	addr_QAbstractGraphicsShapeItem_prepareGeometryChange uintptr
	addr_QAbstractGraphicsShapeItem_isObscuredBy_default uintptr
	addr_QAbstractGraphicsShapeItem_opaqueArea_default uintptr
	addr_QAbstractGraphicsShapeItem_advance_default uintptr
	addr_QAbstractGraphicsShapeItem_shape_default uintptr
	addr_QAbstractGraphicsShapeItem_contains_default uintptr
	addr_QAbstractGraphicsShapeItem_collidesWithItem_default uintptr
	addr_QAbstractGraphicsShapeItem_collidesWithPath_default uintptr
	addr_QAbstractGraphicsShapeItem_type_default uintptr
	addr_QAbstractGraphicsShapeItem_sceneEventFilter_default uintptr
	addr_QAbstractGraphicsShapeItem_sceneEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_contextMenuEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_dragEnterEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_dragLeaveEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_dragMoveEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_dropEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_focusInEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_focusOutEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_hoverEnterEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_hoverMoveEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_hoverLeaveEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_keyPressEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_keyReleaseEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_mousePressEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_mouseMoveEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_mouseReleaseEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_wheelEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_inputMethodEvent_default uintptr
	addr_QAbstractGraphicsShapeItem_inputMethodQuery_default uintptr
	addr_QAbstractGraphicsShapeItem_itemChange_default uintptr
	addr_QAbstractGraphicsShapeItem_supportsExtension_default uintptr
	addr_QAbstractGraphicsShapeItem_setExtension_default uintptr
	addr_QAbstractGraphicsShapeItem_extension_default uintptr
	pg_QAbstractGraphicsShapeItem__set_override func(uintptr, int32)
	pg_QAbstractGraphicsShapeItem__Go_new func(int64) uintptr
	pg_QAbstractGraphicsShapeItem_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractGraphicsShapeItem_base, "QAbstractGraphicsShapeItem_base")
	symaddr(&addr_QAbstractGraphicsShapeItem_pen, "QAbstractGraphicsShapeItem_pen")
	symaddr(&addr_QAbstractGraphicsShapeItem_setPen, "QAbstractGraphicsShapeItem_setPen")
	symaddr(&addr_QAbstractGraphicsShapeItem_brush, "QAbstractGraphicsShapeItem_brush")
	symaddr(&addr_QAbstractGraphicsShapeItem_setBrush, "QAbstractGraphicsShapeItem_setBrush")
	symaddr(&addr_QAbstractGraphicsShapeItem_updateMicroFocus, "QAbstractGraphicsShapeItem_updateMicroFocus")
	symaddr(&addr_QAbstractGraphicsShapeItem_addToIndex, "QAbstractGraphicsShapeItem_addToIndex")
	symaddr(&addr_QAbstractGraphicsShapeItem_removeFromIndex, "QAbstractGraphicsShapeItem_removeFromIndex")
	symaddr(&addr_QAbstractGraphicsShapeItem_prepareGeometryChange, "QAbstractGraphicsShapeItem_prepareGeometryChange")
	symaddr(&addr_QAbstractGraphicsShapeItem_isObscuredBy_default, "QAbstractGraphicsShapeItem_isObscuredBy_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_opaqueArea_default, "QAbstractGraphicsShapeItem_opaqueArea_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_advance_default, "QAbstractGraphicsShapeItem_advance_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_shape_default, "QAbstractGraphicsShapeItem_shape_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_contains_default, "QAbstractGraphicsShapeItem_contains_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_collidesWithItem_default, "QAbstractGraphicsShapeItem_collidesWithItem_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_collidesWithPath_default, "QAbstractGraphicsShapeItem_collidesWithPath_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_type_default, "QAbstractGraphicsShapeItem_type_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_sceneEventFilter_default, "QAbstractGraphicsShapeItem_sceneEventFilter_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_sceneEvent_default, "QAbstractGraphicsShapeItem_sceneEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_contextMenuEvent_default, "QAbstractGraphicsShapeItem_contextMenuEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_dragEnterEvent_default, "QAbstractGraphicsShapeItem_dragEnterEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_dragLeaveEvent_default, "QAbstractGraphicsShapeItem_dragLeaveEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_dragMoveEvent_default, "QAbstractGraphicsShapeItem_dragMoveEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_dropEvent_default, "QAbstractGraphicsShapeItem_dropEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_focusInEvent_default, "QAbstractGraphicsShapeItem_focusInEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_focusOutEvent_default, "QAbstractGraphicsShapeItem_focusOutEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_hoverEnterEvent_default, "QAbstractGraphicsShapeItem_hoverEnterEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_hoverMoveEvent_default, "QAbstractGraphicsShapeItem_hoverMoveEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_hoverLeaveEvent_default, "QAbstractGraphicsShapeItem_hoverLeaveEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_keyPressEvent_default, "QAbstractGraphicsShapeItem_keyPressEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_keyReleaseEvent_default, "QAbstractGraphicsShapeItem_keyReleaseEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_mousePressEvent_default, "QAbstractGraphicsShapeItem_mousePressEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_mouseMoveEvent_default, "QAbstractGraphicsShapeItem_mouseMoveEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_mouseReleaseEvent_default, "QAbstractGraphicsShapeItem_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default, "QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_wheelEvent_default, "QAbstractGraphicsShapeItem_wheelEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_inputMethodEvent_default, "QAbstractGraphicsShapeItem_inputMethodEvent_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_inputMethodQuery_default, "QAbstractGraphicsShapeItem_inputMethodQuery_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_itemChange_default, "QAbstractGraphicsShapeItem_itemChange_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_supportsExtension_default, "QAbstractGraphicsShapeItem_supportsExtension_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_setExtension_default, "QAbstractGraphicsShapeItem_setExtension_default")
	symaddr(&addr_QAbstractGraphicsShapeItem_extension_default, "QAbstractGraphicsShapeItem_extension_default")
	reg(&pg_QAbstractGraphicsShapeItem__set_override, "QAbstractGraphicsShapeItem__set_override")
	reg(&pg_QAbstractGraphicsShapeItem__Go_new, "QAbstractGraphicsShapeItem__Go_new")
	reg(&pg_QAbstractGraphicsShapeItem_delete, "QAbstractGraphicsShapeItem_delete")
}

type QAbstractGraphicsShapeItem struct {
	h uintptr
	*QGraphicsItem
}

func bindQAbstractGraphicsShapeItem(h uintptr) *QAbstractGraphicsShapeItem {
	if h == 0 {
		return nil
	}
	return &QAbstractGraphicsShapeItem{h: h, QGraphicsItem: bindQGraphicsItem(pg_QAbstractGraphicsShapeItem_base(h))}
}

func newownedQAbstractGraphicsShapeItem(h uintptr) *QAbstractGraphicsShapeItem {
	o := bindQAbstractGraphicsShapeItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractGraphicsShapeItem).Delete)
	}
	return o
}

func (this *QAbstractGraphicsShapeItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractGraphicsShapeItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractGraphicsShapeItem(h unsafe.Pointer) *QAbstractGraphicsShapeItem { return bindQAbstractGraphicsShapeItem(uintptr(h)) }
func UnsafeNewQAbstractGraphicsShapeItemOwned(h unsafe.Pointer) *QAbstractGraphicsShapeItem { return newownedQAbstractGraphicsShapeItem(uintptr(h)) }

func (this *QAbstractGraphicsShapeItem) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsItem != nil {
		this.QGraphicsItem.UnsafeRelease()
	}
}

func (this *QAbstractGraphicsShapeItem) Pen() *QPen {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_pen, this.h)
	runtime.KeepAlive(this)
	return newownedQPen(__r)
}

func (this *QAbstractGraphicsShapeItem) SetPen(pen *QPen) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_setPen, this.h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
}

func (this *QAbstractGraphicsShapeItem) Brush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_brush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QAbstractGraphicsShapeItem) SetBrush(brush *QBrush) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_setBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QAbstractGraphicsShapeItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractGraphicsShapeItem) AddToIndex() {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractGraphicsShapeItem) RemoveFromIndex() {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractGraphicsShapeItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractGraphicsShapeItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QAbstractGraphicsShapeItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QAbstractGraphicsShapeItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QAbstractGraphicsShapeItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractGraphicsShapeItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractGraphicsShapeItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractGraphicsShapeItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QAbstractGraphicsShapeItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractGraphicsShapeItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QAbstractGraphicsShapeItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QAbstractGraphicsShapeItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractGraphicsShapeItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func QAbstractGraphicsShapeItem_virtIdx(slot int64) int32 {
	switch slot {
	case 2510008782785038831:
		return 0
	case 6083584261053556819:
		return 1
	case -7215030271914903671:
		return 2
	case 399370970023989435:
		return 3
	case -4830836341480073475:
		return 4
	case 7938019069247444294:
		return 5
	case 5645366511944179768:
		return 6
	case 6469736184362593173:
		return 7
	case 6636575182718753853:
		return 8
	case -6371403872082420691:
		return 9
	case -4584203218241215725:
		return 10
	case -2846371507838587071:
		return 11
	case -6698043222227006015:
		return 12
	case 206122715919917589:
		return 13
	case -3139904158373725006:
		return 14
	case -1781545137352745978:
		return 15
	case 9107058682754591716:
		return 16
	case 975043767033552552:
		return 17
	case -385596010576048993:
		return 18
	case 6932941943906681679:
		return 19
	case -2745191522144316396:
		return 20
	case 7944408482203603984:
		return 21
	case 1023791977679793185:
		return 22
	case 772598475950050299:
		return 23
	case -332090898078543319:
		return 24
	case -4255494004099806581:
		return 25
	case -4415920049297795213:
		return 26
	case -4029073905697023887:
		return 27
	case 1919988340934511316:
		return 28
	case 5438819661942843222:
		return 29
	case 7345853097340709638:
		return 30
	case -1888572168794878830:
		return 31
	case -5729166443651306418:
		return 32
	case -6730181576106171580:
		return 33
	case 2270616358835251226:
		return 34
	}
	return -1
}

func QAbstractGraphicsShapeItem_enable(h uintptr, idx int32) { pg_QAbstractGraphicsShapeItem__set_override(h, idx) }

func NewQAbstractGraphicsShapeItem() *QAbstractGraphicsShapeItem {
	id := nextVirtID()
	o := newownedQAbstractGraphicsShapeItem(pg_QAbstractGraphicsShapeItem__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractGraphicsShapeItem_virtIdx, QAbstractGraphicsShapeItem_enable)
	return o
}

func (this *QAbstractGraphicsShapeItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QAbstractGraphicsShapeItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QAbstractGraphicsShapeItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QAbstractGraphicsShapeItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QAbstractGraphicsShapeItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QAbstractGraphicsShapeItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QAbstractGraphicsShapeItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QAbstractGraphicsShapeItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QAbstractGraphicsShapeItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QAbstractGraphicsShapeItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QAbstractGraphicsShapeItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QAbstractGraphicsShapeItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QAbstractGraphicsShapeItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QAbstractGraphicsShapeItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QAbstractGraphicsShapeItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractGraphicsShapeItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QAbstractGraphicsShapeItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QAbstractGraphicsShapeItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QAbstractGraphicsShapeItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QAbstractGraphicsShapeItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QAbstractGraphicsShapeItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QAbstractGraphicsShapeItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QAbstractGraphicsShapeItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QAbstractGraphicsShapeItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractGraphicsShapeItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QAbstractGraphicsShapeItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractGraphicsShapeItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractGraphicsShapeItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractGraphicsShapeItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QAbstractGraphicsShapeItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QAbstractGraphicsShapeItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractGraphicsShapeItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QAbstractGraphicsShapeItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QAbstractGraphicsShapeItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QAbstractGraphicsShapeItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QAbstractGraphicsShapeItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractGraphicsShapeItem_delete(this.h)
	}
	this.UnsafeRelease()
}
