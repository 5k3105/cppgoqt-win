package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsPathItem_chain uintptr
	addr_QGraphicsPathItem_new uintptr
	addr_QGraphicsPathItem_new2 uintptr
	addr_QGraphicsPathItem_new3 uintptr
	addr_QGraphicsPathItem_new4 uintptr
	addr_QGraphicsPathItem_path uintptr
	addr_QGraphicsPathItem_setPath uintptr
	addr_QGraphicsPathItem_updateMicroFocus uintptr
	addr_QGraphicsPathItem_addToIndex uintptr
	addr_QGraphicsPathItem_removeFromIndex uintptr
	addr_QGraphicsPathItem_prepareGeometryChange uintptr
	addr_QGraphicsPathItem_boundingRect_default uintptr
	addr_QGraphicsPathItem_shape_default uintptr
	addr_QGraphicsPathItem_contains_default uintptr
	addr_QGraphicsPathItem_paint_default uintptr
	addr_QGraphicsPathItem_isObscuredBy_default uintptr
	addr_QGraphicsPathItem_opaqueArea_default uintptr
	addr_QGraphicsPathItem_type_default uintptr
	addr_QGraphicsPathItem_supportsExtension_default uintptr
	addr_QGraphicsPathItem_setExtension_default uintptr
	addr_QGraphicsPathItem_extension_default uintptr
	addr_QGraphicsPathItem_advance_default uintptr
	addr_QGraphicsPathItem_collidesWithItem_default uintptr
	addr_QGraphicsPathItem_collidesWithPath_default uintptr
	addr_QGraphicsPathItem_sceneEventFilter_default uintptr
	addr_QGraphicsPathItem_sceneEvent_default uintptr
	addr_QGraphicsPathItem_contextMenuEvent_default uintptr
	addr_QGraphicsPathItem_dragEnterEvent_default uintptr
	addr_QGraphicsPathItem_dragLeaveEvent_default uintptr
	addr_QGraphicsPathItem_dragMoveEvent_default uintptr
	addr_QGraphicsPathItem_dropEvent_default uintptr
	addr_QGraphicsPathItem_focusInEvent_default uintptr
	addr_QGraphicsPathItem_focusOutEvent_default uintptr
	addr_QGraphicsPathItem_hoverEnterEvent_default uintptr
	addr_QGraphicsPathItem_hoverMoveEvent_default uintptr
	addr_QGraphicsPathItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsPathItem_keyPressEvent_default uintptr
	addr_QGraphicsPathItem_keyReleaseEvent_default uintptr
	addr_QGraphicsPathItem_mousePressEvent_default uintptr
	addr_QGraphicsPathItem_mouseMoveEvent_default uintptr
	addr_QGraphicsPathItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsPathItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsPathItem_wheelEvent_default uintptr
	addr_QGraphicsPathItem_inputMethodEvent_default uintptr
	addr_QGraphicsPathItem_inputMethodQuery_default uintptr
	addr_QGraphicsPathItem_itemChange_default uintptr
	pg_QGraphicsPathItem__set_override func(uintptr, int32)
	pg_QGraphicsPathItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsPathItem_chain, "QGraphicsPathItem_chain")
	symaddr(&addr_QGraphicsPathItem_new, "QGraphicsPathItem_new")
	symaddr(&addr_QGraphicsPathItem_new2, "QGraphicsPathItem_new2")
	symaddr(&addr_QGraphicsPathItem_new3, "QGraphicsPathItem_new3")
	symaddr(&addr_QGraphicsPathItem_new4, "QGraphicsPathItem_new4")
	symaddr(&addr_QGraphicsPathItem_path, "QGraphicsPathItem_path")
	symaddr(&addr_QGraphicsPathItem_setPath, "QGraphicsPathItem_setPath")
	symaddr(&addr_QGraphicsPathItem_updateMicroFocus, "QGraphicsPathItem_updateMicroFocus")
	symaddr(&addr_QGraphicsPathItem_addToIndex, "QGraphicsPathItem_addToIndex")
	symaddr(&addr_QGraphicsPathItem_removeFromIndex, "QGraphicsPathItem_removeFromIndex")
	symaddr(&addr_QGraphicsPathItem_prepareGeometryChange, "QGraphicsPathItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsPathItem_boundingRect_default, "QGraphicsPathItem_boundingRect_default")
	symaddr(&addr_QGraphicsPathItem_shape_default, "QGraphicsPathItem_shape_default")
	symaddr(&addr_QGraphicsPathItem_contains_default, "QGraphicsPathItem_contains_default")
	symaddr(&addr_QGraphicsPathItem_paint_default, "QGraphicsPathItem_paint_default")
	symaddr(&addr_QGraphicsPathItem_isObscuredBy_default, "QGraphicsPathItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsPathItem_opaqueArea_default, "QGraphicsPathItem_opaqueArea_default")
	symaddr(&addr_QGraphicsPathItem_type_default, "QGraphicsPathItem_type_default")
	symaddr(&addr_QGraphicsPathItem_supportsExtension_default, "QGraphicsPathItem_supportsExtension_default")
	symaddr(&addr_QGraphicsPathItem_setExtension_default, "QGraphicsPathItem_setExtension_default")
	symaddr(&addr_QGraphicsPathItem_extension_default, "QGraphicsPathItem_extension_default")
	symaddr(&addr_QGraphicsPathItem_advance_default, "QGraphicsPathItem_advance_default")
	symaddr(&addr_QGraphicsPathItem_collidesWithItem_default, "QGraphicsPathItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsPathItem_collidesWithPath_default, "QGraphicsPathItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsPathItem_sceneEventFilter_default, "QGraphicsPathItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsPathItem_sceneEvent_default, "QGraphicsPathItem_sceneEvent_default")
	symaddr(&addr_QGraphicsPathItem_contextMenuEvent_default, "QGraphicsPathItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsPathItem_dragEnterEvent_default, "QGraphicsPathItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsPathItem_dragLeaveEvent_default, "QGraphicsPathItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsPathItem_dragMoveEvent_default, "QGraphicsPathItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsPathItem_dropEvent_default, "QGraphicsPathItem_dropEvent_default")
	symaddr(&addr_QGraphicsPathItem_focusInEvent_default, "QGraphicsPathItem_focusInEvent_default")
	symaddr(&addr_QGraphicsPathItem_focusOutEvent_default, "QGraphicsPathItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsPathItem_hoverEnterEvent_default, "QGraphicsPathItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsPathItem_hoverMoveEvent_default, "QGraphicsPathItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsPathItem_hoverLeaveEvent_default, "QGraphicsPathItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsPathItem_keyPressEvent_default, "QGraphicsPathItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsPathItem_keyReleaseEvent_default, "QGraphicsPathItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsPathItem_mousePressEvent_default, "QGraphicsPathItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsPathItem_mouseMoveEvent_default, "QGraphicsPathItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsPathItem_mouseReleaseEvent_default, "QGraphicsPathItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsPathItem_mouseDoubleClickEvent_default, "QGraphicsPathItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsPathItem_wheelEvent_default, "QGraphicsPathItem_wheelEvent_default")
	symaddr(&addr_QGraphicsPathItem_inputMethodEvent_default, "QGraphicsPathItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsPathItem_inputMethodQuery_default, "QGraphicsPathItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsPathItem_itemChange_default, "QGraphicsPathItem_itemChange_default")
	reg(&pg_QGraphicsPathItem__set_override, "QGraphicsPathItem__set_override")
	reg(&pg_QGraphicsPathItem_delete, "QGraphicsPathItem_delete")
}

type QGraphicsPathItem struct {
	h uintptr
	*QAbstractGraphicsShapeItem
}

func bindQGraphicsPathItem(h uintptr) *QGraphicsPathItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsPathItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsPathItem{h: h, QAbstractGraphicsShapeItem: &QAbstractGraphicsShapeItem{h: __c[0], QGraphicsItem: &QGraphicsItem{h: __c[1]}}}
}

func newownedQGraphicsPathItem(h uintptr) *QGraphicsPathItem {
	o := bindQGraphicsPathItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsPathItem).Delete)
	}
	return o
}

func (this *QGraphicsPathItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsPathItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsPathItem(h unsafe.Pointer) *QGraphicsPathItem { return bindQGraphicsPathItem(uintptr(h)) }
func UnsafeNewQGraphicsPathItemOwned(h unsafe.Pointer) *QGraphicsPathItem { return newownedQGraphicsPathItem(uintptr(h)) }

func (this *QGraphicsPathItem) UnsafeRelease() {
	this.h = 0
	if this.QAbstractGraphicsShapeItem != nil {
		this.QAbstractGraphicsShapeItem.UnsafeRelease()
	}
}

func NewQGraphicsPathItem() *QGraphicsPathItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_new, uintptr(__goid))
	__o := newownedQGraphicsPathItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPathItem_virtIdx, QGraphicsPathItem_enable)
	return __o
}

func NewQGraphicsPathItem2(path *QPainterPath) *QGraphicsPathItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_new2, uintptr(__goid), uintptr(path.UnsafePointer()))
	__o := newownedQGraphicsPathItem(__r)
	runtime.KeepAlive(path)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPathItem_virtIdx, QGraphicsPathItem_enable)
	return __o
}

func NewQGraphicsPathItem3(parent *QGraphicsItem) *QGraphicsPathItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPathItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPathItem_virtIdx, QGraphicsPathItem_enable)
	return __o
}

func NewQGraphicsPathItem4(path *QPainterPath, parent *QGraphicsItem) *QGraphicsPathItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_new4, uintptr(__goid), uintptr(path.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsPathItem(__r)
	runtime.KeepAlive(path)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsPathItem_virtIdx, QGraphicsPathItem_enable)
	return __o
}

func (this *QGraphicsPathItem) Path() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_path, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPathItem) SetPath(path *QPainterPath) {
	purego.SyscallN(addr_QGraphicsPathItem_setPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QGraphicsPathItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsPathItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPathItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsPathItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPathItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsPathItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPathItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsPathItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsPathItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsPathItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPathItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsPathItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsPathItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsPathItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsPathItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsPathItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsPathItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsPathItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsPathItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsPathItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsPathItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsPathItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsPathItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsPathItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsPathItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPathItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsPathItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsPathItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsPathItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsPathItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsPathItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsPathItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsPathItem_virtIdx(slot int64) int32 {
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

func QGraphicsPathItem_enable(h uintptr, idx int32) { pg_QGraphicsPathItem__set_override(h, idx) }

func (this *QGraphicsPathItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsPathItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsPathItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsPathItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsPathItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsPathItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsPathItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsPathItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsPathItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsPathItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsPathItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsPathItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsPathItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsPathItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsPathItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsPathItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsPathItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsPathItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsPathItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsPathItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsPathItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsPathItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsPathItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsPathItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsPathItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsPathItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsPathItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsPathItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsPathItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsPathItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsPathItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsPathItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsPathItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsPathItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsPathItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsPathItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsPathItem_delete(this.h)
	}
	this.UnsafeRelease()
}
