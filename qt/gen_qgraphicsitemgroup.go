package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsItemGroup_base func(uintptr) uintptr
	addr_QGraphicsItemGroup_new uintptr
	addr_QGraphicsItemGroup_new2 uintptr
	addr_QGraphicsItemGroup_addToGroup uintptr
	addr_QGraphicsItemGroup_removeFromGroup uintptr
	addr_QGraphicsItemGroup_updateMicroFocus uintptr
	addr_QGraphicsItemGroup_addToIndex uintptr
	addr_QGraphicsItemGroup_removeFromIndex uintptr
	addr_QGraphicsItemGroup_prepareGeometryChange uintptr
	addr_QGraphicsItemGroup_boundingRect_default uintptr
	addr_QGraphicsItemGroup_paint_default uintptr
	addr_QGraphicsItemGroup_isObscuredBy_default uintptr
	addr_QGraphicsItemGroup_opaqueArea_default uintptr
	addr_QGraphicsItemGroup_type_default uintptr
	addr_QGraphicsItemGroup_advance_default uintptr
	addr_QGraphicsItemGroup_shape_default uintptr
	addr_QGraphicsItemGroup_contains_default uintptr
	addr_QGraphicsItemGroup_collidesWithItem_default uintptr
	addr_QGraphicsItemGroup_collidesWithPath_default uintptr
	addr_QGraphicsItemGroup_sceneEventFilter_default uintptr
	addr_QGraphicsItemGroup_sceneEvent_default uintptr
	addr_QGraphicsItemGroup_contextMenuEvent_default uintptr
	addr_QGraphicsItemGroup_dragEnterEvent_default uintptr
	addr_QGraphicsItemGroup_dragLeaveEvent_default uintptr
	addr_QGraphicsItemGroup_dragMoveEvent_default uintptr
	addr_QGraphicsItemGroup_dropEvent_default uintptr
	addr_QGraphicsItemGroup_focusInEvent_default uintptr
	addr_QGraphicsItemGroup_focusOutEvent_default uintptr
	addr_QGraphicsItemGroup_hoverEnterEvent_default uintptr
	addr_QGraphicsItemGroup_hoverMoveEvent_default uintptr
	addr_QGraphicsItemGroup_hoverLeaveEvent_default uintptr
	addr_QGraphicsItemGroup_keyPressEvent_default uintptr
	addr_QGraphicsItemGroup_keyReleaseEvent_default uintptr
	addr_QGraphicsItemGroup_mousePressEvent_default uintptr
	addr_QGraphicsItemGroup_mouseMoveEvent_default uintptr
	addr_QGraphicsItemGroup_mouseReleaseEvent_default uintptr
	addr_QGraphicsItemGroup_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsItemGroup_wheelEvent_default uintptr
	addr_QGraphicsItemGroup_inputMethodEvent_default uintptr
	addr_QGraphicsItemGroup_inputMethodQuery_default uintptr
	addr_QGraphicsItemGroup_itemChange_default uintptr
	addr_QGraphicsItemGroup_supportsExtension_default uintptr
	addr_QGraphicsItemGroup_setExtension_default uintptr
	addr_QGraphicsItemGroup_extension_default uintptr
	pg_QGraphicsItemGroup__set_override func(uintptr, int32)
	pg_QGraphicsItemGroup_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsItemGroup_base, "QGraphicsItemGroup_base")
	symaddr(&addr_QGraphicsItemGroup_new, "QGraphicsItemGroup_new")
	symaddr(&addr_QGraphicsItemGroup_new2, "QGraphicsItemGroup_new2")
	symaddr(&addr_QGraphicsItemGroup_addToGroup, "QGraphicsItemGroup_addToGroup")
	symaddr(&addr_QGraphicsItemGroup_removeFromGroup, "QGraphicsItemGroup_removeFromGroup")
	symaddr(&addr_QGraphicsItemGroup_updateMicroFocus, "QGraphicsItemGroup_updateMicroFocus")
	symaddr(&addr_QGraphicsItemGroup_addToIndex, "QGraphicsItemGroup_addToIndex")
	symaddr(&addr_QGraphicsItemGroup_removeFromIndex, "QGraphicsItemGroup_removeFromIndex")
	symaddr(&addr_QGraphicsItemGroup_prepareGeometryChange, "QGraphicsItemGroup_prepareGeometryChange")
	symaddr(&addr_QGraphicsItemGroup_boundingRect_default, "QGraphicsItemGroup_boundingRect_default")
	symaddr(&addr_QGraphicsItemGroup_paint_default, "QGraphicsItemGroup_paint_default")
	symaddr(&addr_QGraphicsItemGroup_isObscuredBy_default, "QGraphicsItemGroup_isObscuredBy_default")
	symaddr(&addr_QGraphicsItemGroup_opaqueArea_default, "QGraphicsItemGroup_opaqueArea_default")
	symaddr(&addr_QGraphicsItemGroup_type_default, "QGraphicsItemGroup_type_default")
	symaddr(&addr_QGraphicsItemGroup_advance_default, "QGraphicsItemGroup_advance_default")
	symaddr(&addr_QGraphicsItemGroup_shape_default, "QGraphicsItemGroup_shape_default")
	symaddr(&addr_QGraphicsItemGroup_contains_default, "QGraphicsItemGroup_contains_default")
	symaddr(&addr_QGraphicsItemGroup_collidesWithItem_default, "QGraphicsItemGroup_collidesWithItem_default")
	symaddr(&addr_QGraphicsItemGroup_collidesWithPath_default, "QGraphicsItemGroup_collidesWithPath_default")
	symaddr(&addr_QGraphicsItemGroup_sceneEventFilter_default, "QGraphicsItemGroup_sceneEventFilter_default")
	symaddr(&addr_QGraphicsItemGroup_sceneEvent_default, "QGraphicsItemGroup_sceneEvent_default")
	symaddr(&addr_QGraphicsItemGroup_contextMenuEvent_default, "QGraphicsItemGroup_contextMenuEvent_default")
	symaddr(&addr_QGraphicsItemGroup_dragEnterEvent_default, "QGraphicsItemGroup_dragEnterEvent_default")
	symaddr(&addr_QGraphicsItemGroup_dragLeaveEvent_default, "QGraphicsItemGroup_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsItemGroup_dragMoveEvent_default, "QGraphicsItemGroup_dragMoveEvent_default")
	symaddr(&addr_QGraphicsItemGroup_dropEvent_default, "QGraphicsItemGroup_dropEvent_default")
	symaddr(&addr_QGraphicsItemGroup_focusInEvent_default, "QGraphicsItemGroup_focusInEvent_default")
	symaddr(&addr_QGraphicsItemGroup_focusOutEvent_default, "QGraphicsItemGroup_focusOutEvent_default")
	symaddr(&addr_QGraphicsItemGroup_hoverEnterEvent_default, "QGraphicsItemGroup_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsItemGroup_hoverMoveEvent_default, "QGraphicsItemGroup_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsItemGroup_hoverLeaveEvent_default, "QGraphicsItemGroup_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsItemGroup_keyPressEvent_default, "QGraphicsItemGroup_keyPressEvent_default")
	symaddr(&addr_QGraphicsItemGroup_keyReleaseEvent_default, "QGraphicsItemGroup_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsItemGroup_mousePressEvent_default, "QGraphicsItemGroup_mousePressEvent_default")
	symaddr(&addr_QGraphicsItemGroup_mouseMoveEvent_default, "QGraphicsItemGroup_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsItemGroup_mouseReleaseEvent_default, "QGraphicsItemGroup_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsItemGroup_mouseDoubleClickEvent_default, "QGraphicsItemGroup_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsItemGroup_wheelEvent_default, "QGraphicsItemGroup_wheelEvent_default")
	symaddr(&addr_QGraphicsItemGroup_inputMethodEvent_default, "QGraphicsItemGroup_inputMethodEvent_default")
	symaddr(&addr_QGraphicsItemGroup_inputMethodQuery_default, "QGraphicsItemGroup_inputMethodQuery_default")
	symaddr(&addr_QGraphicsItemGroup_itemChange_default, "QGraphicsItemGroup_itemChange_default")
	symaddr(&addr_QGraphicsItemGroup_supportsExtension_default, "QGraphicsItemGroup_supportsExtension_default")
	symaddr(&addr_QGraphicsItemGroup_setExtension_default, "QGraphicsItemGroup_setExtension_default")
	symaddr(&addr_QGraphicsItemGroup_extension_default, "QGraphicsItemGroup_extension_default")
	reg(&pg_QGraphicsItemGroup__set_override, "QGraphicsItemGroup__set_override")
	reg(&pg_QGraphicsItemGroup_delete, "QGraphicsItemGroup_delete")
}

type QGraphicsItemGroup struct {
	h uintptr
	*QGraphicsItem
}

func bindQGraphicsItemGroup(h uintptr) *QGraphicsItemGroup {
	if h == 0 {
		return nil
	}
	return &QGraphicsItemGroup{h: h, QGraphicsItem: bindQGraphicsItem(pg_QGraphicsItemGroup_base(h))}
}

func newownedQGraphicsItemGroup(h uintptr) *QGraphicsItemGroup {
	o := bindQGraphicsItemGroup(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsItemGroup).Delete)
	}
	return o
}

func (this *QGraphicsItemGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsItemGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsItemGroup(h unsafe.Pointer) *QGraphicsItemGroup { return bindQGraphicsItemGroup(uintptr(h)) }
func UnsafeNewQGraphicsItemGroupOwned(h unsafe.Pointer) *QGraphicsItemGroup { return newownedQGraphicsItemGroup(uintptr(h)) }

func (this *QGraphicsItemGroup) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsItem != nil {
		this.QGraphicsItem.UnsafeRelease()
	}
}

func NewQGraphicsItemGroup() *QGraphicsItemGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_new, uintptr(__goid))
	__o := newownedQGraphicsItemGroup(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsItemGroup_virtIdx, QGraphicsItemGroup_enable)
	return __o
}

func NewQGraphicsItemGroup2(parent *QGraphicsItem) *QGraphicsItemGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsItemGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsItemGroup_virtIdx, QGraphicsItemGroup_enable)
	return __o
}

func (this *QGraphicsItemGroup) AddToGroup(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItemGroup_addToGroup, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsItemGroup) RemoveFromGroup(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItemGroup_removeFromGroup, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsItemGroup) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsItemGroup_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemGroup) AddToIndex() {
	purego.SyscallN(addr_QGraphicsItemGroup_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemGroup) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsItemGroup_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemGroup) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsItemGroup_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemGroup) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItemGroup) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsItemGroup_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsItemGroup) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsItemGroup) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItemGroup) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsItemGroup) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsItemGroup_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItemGroup) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItemGroup) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsItemGroup) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsItemGroup) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsItemGroup) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItemGroup) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItemGroup) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsItemGroup_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItemGroup) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsItemGroup) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsItemGroup) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItemGroup) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsItemGroup_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsItemGroup) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItemGroup_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func QGraphicsItemGroup_virtIdx(slot int64) int32 {
	switch slot {
	case 2510008782785038831:
		return 0
	case 6083584261053556819:
		return 1
	case -7215030271914903671:
		return 2
	case 399370970023989435:
		return 3
	case -6371403872082420691:
		return 4
	case -4830836341480073475:
		return 5
	case 7938019069247444294:
		return 6
	case 5645366511944179768:
		return 7
	case 6469736184362593173:
		return 8
	case 6636575182718753853:
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

func QGraphicsItemGroup_enable(h uintptr, idx int32) { pg_QGraphicsItemGroup__set_override(h, idx) }

func (this *QGraphicsItemGroup) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsItemGroup) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsItemGroup) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsItemGroup) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsItemGroup) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsItemGroup) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsItemGroup) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsItemGroup) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsItemGroup) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsItemGroup) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsItemGroup) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsItemGroup) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsItemGroup) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsItemGroup) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsItemGroup) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsItemGroup) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsItemGroup) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsItemGroup) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsItemGroup) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsItemGroup) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsItemGroup) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsItemGroup) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsItemGroup) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsItemGroup) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsItemGroup) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsItemGroup) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsItemGroup) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsItemGroup) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsItemGroup) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsItemGroup) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsItemGroup) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsItemGroup) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsItemGroup) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsItemGroup) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsItemGroup) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsItemGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsItemGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
