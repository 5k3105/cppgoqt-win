package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsSimpleTextItem_chain uintptr
	addr_QGraphicsSimpleTextItem_new uintptr
	addr_QGraphicsSimpleTextItem_new2 uintptr
	addr_QGraphicsSimpleTextItem_new3 uintptr
	addr_QGraphicsSimpleTextItem_new4 uintptr
	addr_QGraphicsSimpleTextItem_setText uintptr
	addr_QGraphicsSimpleTextItem_text uintptr
	addr_QGraphicsSimpleTextItem_setFont uintptr
	addr_QGraphicsSimpleTextItem_font uintptr
	addr_QGraphicsSimpleTextItem_updateMicroFocus uintptr
	addr_QGraphicsSimpleTextItem_addToIndex uintptr
	addr_QGraphicsSimpleTextItem_removeFromIndex uintptr
	addr_QGraphicsSimpleTextItem_prepareGeometryChange uintptr
	addr_QGraphicsSimpleTextItem_boundingRect_default uintptr
	addr_QGraphicsSimpleTextItem_shape_default uintptr
	addr_QGraphicsSimpleTextItem_contains_default uintptr
	addr_QGraphicsSimpleTextItem_paint_default uintptr
	addr_QGraphicsSimpleTextItem_isObscuredBy_default uintptr
	addr_QGraphicsSimpleTextItem_opaqueArea_default uintptr
	addr_QGraphicsSimpleTextItem_type_default uintptr
	addr_QGraphicsSimpleTextItem_supportsExtension_default uintptr
	addr_QGraphicsSimpleTextItem_setExtension_default uintptr
	addr_QGraphicsSimpleTextItem_extension_default uintptr
	addr_QGraphicsSimpleTextItem_advance_default uintptr
	addr_QGraphicsSimpleTextItem_collidesWithItem_default uintptr
	addr_QGraphicsSimpleTextItem_collidesWithPath_default uintptr
	addr_QGraphicsSimpleTextItem_sceneEventFilter_default uintptr
	addr_QGraphicsSimpleTextItem_sceneEvent_default uintptr
	addr_QGraphicsSimpleTextItem_contextMenuEvent_default uintptr
	addr_QGraphicsSimpleTextItem_dragEnterEvent_default uintptr
	addr_QGraphicsSimpleTextItem_dragLeaveEvent_default uintptr
	addr_QGraphicsSimpleTextItem_dragMoveEvent_default uintptr
	addr_QGraphicsSimpleTextItem_dropEvent_default uintptr
	addr_QGraphicsSimpleTextItem_focusInEvent_default uintptr
	addr_QGraphicsSimpleTextItem_focusOutEvent_default uintptr
	addr_QGraphicsSimpleTextItem_hoverEnterEvent_default uintptr
	addr_QGraphicsSimpleTextItem_hoverMoveEvent_default uintptr
	addr_QGraphicsSimpleTextItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsSimpleTextItem_keyPressEvent_default uintptr
	addr_QGraphicsSimpleTextItem_keyReleaseEvent_default uintptr
	addr_QGraphicsSimpleTextItem_mousePressEvent_default uintptr
	addr_QGraphicsSimpleTextItem_mouseMoveEvent_default uintptr
	addr_QGraphicsSimpleTextItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsSimpleTextItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsSimpleTextItem_wheelEvent_default uintptr
	addr_QGraphicsSimpleTextItem_inputMethodEvent_default uintptr
	addr_QGraphicsSimpleTextItem_inputMethodQuery_default uintptr
	addr_QGraphicsSimpleTextItem_itemChange_default uintptr
	pg_QGraphicsSimpleTextItem__set_override func(uintptr, int32)
	pg_QGraphicsSimpleTextItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsSimpleTextItem_chain, "QGraphicsSimpleTextItem_chain")
	symaddr(&addr_QGraphicsSimpleTextItem_new, "QGraphicsSimpleTextItem_new")
	symaddr(&addr_QGraphicsSimpleTextItem_new2, "QGraphicsSimpleTextItem_new2")
	symaddr(&addr_QGraphicsSimpleTextItem_new3, "QGraphicsSimpleTextItem_new3")
	symaddr(&addr_QGraphicsSimpleTextItem_new4, "QGraphicsSimpleTextItem_new4")
	symaddr(&addr_QGraphicsSimpleTextItem_setText, "QGraphicsSimpleTextItem_setText")
	symaddr(&addr_QGraphicsSimpleTextItem_text, "QGraphicsSimpleTextItem_text")
	symaddr(&addr_QGraphicsSimpleTextItem_setFont, "QGraphicsSimpleTextItem_setFont")
	symaddr(&addr_QGraphicsSimpleTextItem_font, "QGraphicsSimpleTextItem_font")
	symaddr(&addr_QGraphicsSimpleTextItem_updateMicroFocus, "QGraphicsSimpleTextItem_updateMicroFocus")
	symaddr(&addr_QGraphicsSimpleTextItem_addToIndex, "QGraphicsSimpleTextItem_addToIndex")
	symaddr(&addr_QGraphicsSimpleTextItem_removeFromIndex, "QGraphicsSimpleTextItem_removeFromIndex")
	symaddr(&addr_QGraphicsSimpleTextItem_prepareGeometryChange, "QGraphicsSimpleTextItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsSimpleTextItem_boundingRect_default, "QGraphicsSimpleTextItem_boundingRect_default")
	symaddr(&addr_QGraphicsSimpleTextItem_shape_default, "QGraphicsSimpleTextItem_shape_default")
	symaddr(&addr_QGraphicsSimpleTextItem_contains_default, "QGraphicsSimpleTextItem_contains_default")
	symaddr(&addr_QGraphicsSimpleTextItem_paint_default, "QGraphicsSimpleTextItem_paint_default")
	symaddr(&addr_QGraphicsSimpleTextItem_isObscuredBy_default, "QGraphicsSimpleTextItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsSimpleTextItem_opaqueArea_default, "QGraphicsSimpleTextItem_opaqueArea_default")
	symaddr(&addr_QGraphicsSimpleTextItem_type_default, "QGraphicsSimpleTextItem_type_default")
	symaddr(&addr_QGraphicsSimpleTextItem_supportsExtension_default, "QGraphicsSimpleTextItem_supportsExtension_default")
	symaddr(&addr_QGraphicsSimpleTextItem_setExtension_default, "QGraphicsSimpleTextItem_setExtension_default")
	symaddr(&addr_QGraphicsSimpleTextItem_extension_default, "QGraphicsSimpleTextItem_extension_default")
	symaddr(&addr_QGraphicsSimpleTextItem_advance_default, "QGraphicsSimpleTextItem_advance_default")
	symaddr(&addr_QGraphicsSimpleTextItem_collidesWithItem_default, "QGraphicsSimpleTextItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsSimpleTextItem_collidesWithPath_default, "QGraphicsSimpleTextItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsSimpleTextItem_sceneEventFilter_default, "QGraphicsSimpleTextItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsSimpleTextItem_sceneEvent_default, "QGraphicsSimpleTextItem_sceneEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_contextMenuEvent_default, "QGraphicsSimpleTextItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_dragEnterEvent_default, "QGraphicsSimpleTextItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_dragLeaveEvent_default, "QGraphicsSimpleTextItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_dragMoveEvent_default, "QGraphicsSimpleTextItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_dropEvent_default, "QGraphicsSimpleTextItem_dropEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_focusInEvent_default, "QGraphicsSimpleTextItem_focusInEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_focusOutEvent_default, "QGraphicsSimpleTextItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_hoverEnterEvent_default, "QGraphicsSimpleTextItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_hoverMoveEvent_default, "QGraphicsSimpleTextItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_hoverLeaveEvent_default, "QGraphicsSimpleTextItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_keyPressEvent_default, "QGraphicsSimpleTextItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_keyReleaseEvent_default, "QGraphicsSimpleTextItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_mousePressEvent_default, "QGraphicsSimpleTextItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_mouseMoveEvent_default, "QGraphicsSimpleTextItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_mouseReleaseEvent_default, "QGraphicsSimpleTextItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_mouseDoubleClickEvent_default, "QGraphicsSimpleTextItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_wheelEvent_default, "QGraphicsSimpleTextItem_wheelEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_inputMethodEvent_default, "QGraphicsSimpleTextItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsSimpleTextItem_inputMethodQuery_default, "QGraphicsSimpleTextItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsSimpleTextItem_itemChange_default, "QGraphicsSimpleTextItem_itemChange_default")
	reg(&pg_QGraphicsSimpleTextItem__set_override, "QGraphicsSimpleTextItem__set_override")
	reg(&pg_QGraphicsSimpleTextItem_delete, "QGraphicsSimpleTextItem_delete")
}

type QGraphicsSimpleTextItem struct {
	h uintptr
	*QAbstractGraphicsShapeItem
}

func bindQGraphicsSimpleTextItem(h uintptr) *QGraphicsSimpleTextItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsSimpleTextItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsSimpleTextItem{h: h, QAbstractGraphicsShapeItem: &QAbstractGraphicsShapeItem{h: __c[0], QGraphicsItem: &QGraphicsItem{h: __c[1]}}}
}

func newownedQGraphicsSimpleTextItem(h uintptr) *QGraphicsSimpleTextItem {
	o := bindQGraphicsSimpleTextItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSimpleTextItem).Delete)
	}
	return o
}

func (this *QGraphicsSimpleTextItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSimpleTextItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSimpleTextItem(h unsafe.Pointer) *QGraphicsSimpleTextItem { return bindQGraphicsSimpleTextItem(uintptr(h)) }
func UnsafeNewQGraphicsSimpleTextItemOwned(h unsafe.Pointer) *QGraphicsSimpleTextItem { return newownedQGraphicsSimpleTextItem(uintptr(h)) }

func (this *QGraphicsSimpleTextItem) UnsafeRelease() {
	this.h = 0
	if this.QAbstractGraphicsShapeItem != nil {
		this.QAbstractGraphicsShapeItem.UnsafeRelease()
	}
}

func NewQGraphicsSimpleTextItem() *QGraphicsSimpleTextItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_new, uintptr(__goid))
	__o := newownedQGraphicsSimpleTextItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSimpleTextItem_virtIdx, QGraphicsSimpleTextItem_enable)
	return __o
}

func NewQGraphicsSimpleTextItem2(text string) *QGraphicsSimpleTextItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQGraphicsSimpleTextItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSimpleTextItem_virtIdx, QGraphicsSimpleTextItem_enable)
	return __o
}

func NewQGraphicsSimpleTextItem3(parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsSimpleTextItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSimpleTextItem_virtIdx, QGraphicsSimpleTextItem_enable)
	return __o
}

func NewQGraphicsSimpleTextItem4(text string, parent *QGraphicsItem) *QGraphicsSimpleTextItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsSimpleTextItem(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsSimpleTextItem_virtIdx, QGraphicsSimpleTextItem_enable)
	return __o
}

func (this *QGraphicsSimpleTextItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QGraphicsSimpleTextItem_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QGraphicsSimpleTextItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsSimpleTextItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QGraphicsSimpleTextItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QGraphicsSimpleTextItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsSimpleTextItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsSimpleTextItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsSimpleTextItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsSimpleTextItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsSimpleTextItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsSimpleTextItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsSimpleTextItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsSimpleTextItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsSimpleTextItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsSimpleTextItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsSimpleTextItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsSimpleTextItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsSimpleTextItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsSimpleTextItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsSimpleTextItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsSimpleTextItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSimpleTextItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func QGraphicsSimpleTextItem_virtIdx(slot int64) int32 {
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

func QGraphicsSimpleTextItem_enable(h uintptr, idx int32) { pg_QGraphicsSimpleTextItem__set_override(h, idx) }

func (this *QGraphicsSimpleTextItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsSimpleTextItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsSimpleTextItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsSimpleTextItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsSimpleTextItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsSimpleTextItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsSimpleTextItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsSimpleTextItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsSimpleTextItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsSimpleTextItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsSimpleTextItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsSimpleTextItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsSimpleTextItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsSimpleTextItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsSimpleTextItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsSimpleTextItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsSimpleTextItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsSimpleTextItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsSimpleTextItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsSimpleTextItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsSimpleTextItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsSimpleTextItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsSimpleTextItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsSimpleTextItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsSimpleTextItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsSimpleTextItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsSimpleTextItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsSimpleTextItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsSimpleTextItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsSimpleTextItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsSimpleTextItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsSimpleTextItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsSimpleTextItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsSimpleTextItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsSimpleTextItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsSimpleTextItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSimpleTextItem_delete(this.h)
	}
	this.UnsafeRelease()
}
