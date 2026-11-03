package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsTextItem_chain uintptr
	addr_QGraphicsTextItem_new uintptr
	addr_QGraphicsTextItem_new2 uintptr
	addr_QGraphicsTextItem_new3 uintptr
	addr_QGraphicsTextItem_new4 uintptr
	addr_QGraphicsTextItem_tr uintptr
	addr_QGraphicsTextItem_toHtml uintptr
	addr_QGraphicsTextItem_setHtml uintptr
	addr_QGraphicsTextItem_toPlainText uintptr
	addr_QGraphicsTextItem_setPlainText uintptr
	addr_QGraphicsTextItem_font uintptr
	addr_QGraphicsTextItem_setFont uintptr
	addr_QGraphicsTextItem_setDefaultTextColor uintptr
	addr_QGraphicsTextItem_defaultTextColor uintptr
	pg_QGraphicsTextItem_setTextWidth func(uintptr, float64)
	pg_QGraphicsTextItem_textWidth func(uintptr) float64
	addr_QGraphicsTextItem_adjustSize uintptr
	addr_QGraphicsTextItem_setDocument uintptr
	addr_QGraphicsTextItem_document uintptr
	addr_QGraphicsTextItem_setTextInteractionFlags uintptr
	addr_QGraphicsTextItem_textInteractionFlags uintptr
	addr_QGraphicsTextItem_setTabChangesFocus uintptr
	addr_QGraphicsTextItem_tabChangesFocus uintptr
	addr_QGraphicsTextItem_setOpenExternalLinks uintptr
	addr_QGraphicsTextItem_openExternalLinks uintptr
	addr_QGraphicsTextItem_setTextCursor uintptr
	addr_QGraphicsTextItem_textCursor uintptr
	addr_QGraphicsTextItem_linkActivated uintptr
	addr_QGraphicsTextItem_linkHovered uintptr
	addr_QGraphicsTextItem_updateMicroFocus uintptr
	addr_QGraphicsTextItem_connect uintptr
	addr_QGraphicsTextItem_connect2 uintptr
	addr_QGraphicsTextItem_sender uintptr
	addr_QGraphicsTextItem_senderSignalIndex uintptr
	addr_QGraphicsTextItem_receivers uintptr
	addr_QGraphicsTextItem_isSignalConnected uintptr
	addr_QGraphicsTextItem_addToIndex uintptr
	addr_QGraphicsTextItem_removeFromIndex uintptr
	addr_QGraphicsTextItem_prepareGeometryChange uintptr
	addr_QGraphicsTextItem_tr2 uintptr
	addr_QGraphicsTextItem_tr3 uintptr
	addr_QGraphicsTextItem_connect3 uintptr
	addr_QGraphicsTextItem_connect4 uintptr
	addr_QGraphicsTextItem_metaObject_default uintptr
	addr_QGraphicsTextItem_boundingRect_default uintptr
	addr_QGraphicsTextItem_shape_default uintptr
	addr_QGraphicsTextItem_contains_default uintptr
	addr_QGraphicsTextItem_paint_default uintptr
	addr_QGraphicsTextItem_isObscuredBy_default uintptr
	addr_QGraphicsTextItem_opaqueArea_default uintptr
	addr_QGraphicsTextItem_type_default uintptr
	addr_QGraphicsTextItem_sceneEvent_default uintptr
	addr_QGraphicsTextItem_mousePressEvent_default uintptr
	addr_QGraphicsTextItem_mouseMoveEvent_default uintptr
	addr_QGraphicsTextItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsTextItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsTextItem_contextMenuEvent_default uintptr
	addr_QGraphicsTextItem_keyPressEvent_default uintptr
	addr_QGraphicsTextItem_keyReleaseEvent_default uintptr
	addr_QGraphicsTextItem_focusInEvent_default uintptr
	addr_QGraphicsTextItem_focusOutEvent_default uintptr
	addr_QGraphicsTextItem_dragEnterEvent_default uintptr
	addr_QGraphicsTextItem_dragLeaveEvent_default uintptr
	addr_QGraphicsTextItem_dragMoveEvent_default uintptr
	addr_QGraphicsTextItem_dropEvent_default uintptr
	addr_QGraphicsTextItem_inputMethodEvent_default uintptr
	addr_QGraphicsTextItem_hoverEnterEvent_default uintptr
	addr_QGraphicsTextItem_hoverMoveEvent_default uintptr
	addr_QGraphicsTextItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsTextItem_inputMethodQuery_default uintptr
	addr_QGraphicsTextItem_supportsExtension_default uintptr
	addr_QGraphicsTextItem_setExtension_default uintptr
	addr_QGraphicsTextItem_extension_default uintptr
	addr_QGraphicsTextItem_event_default uintptr
	addr_QGraphicsTextItem_eventFilter_default uintptr
	addr_QGraphicsTextItem_timerEvent_default uintptr
	addr_QGraphicsTextItem_childEvent_default uintptr
	addr_QGraphicsTextItem_customEvent_default uintptr
	addr_QGraphicsTextItem_connectNotify_default uintptr
	addr_QGraphicsTextItem_disconnectNotify_default uintptr
	addr_QGraphicsTextItem_advance_default uintptr
	addr_QGraphicsTextItem_collidesWithItem_default uintptr
	addr_QGraphicsTextItem_collidesWithPath_default uintptr
	addr_QGraphicsTextItem_sceneEventFilter_default uintptr
	addr_QGraphicsTextItem_wheelEvent_default uintptr
	addr_QGraphicsTextItem_itemChange_default uintptr
	pg_QGraphicsTextItem_connect_linkActivated func(uintptr, int64)
	pg_QGraphicsTextItem_connect_linkHovered func(uintptr, int64)
	pg_QGraphicsTextItem_connect_parentChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_opacityChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_visibleChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_xChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_yChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_zChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_rotationChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_scaleChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_childrenChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_widthChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_heightChanged func(uintptr, int64)
	pg_QGraphicsTextItem_connect_destroyed func(uintptr, int64)
	pg_QGraphicsTextItem_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsTextItem_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsTextItem__disconnect func(int64)
	pg_QGraphicsTextItem__set_override func(uintptr, int32)
	pg_QGraphicsTextItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsTextItem_chain, "QGraphicsTextItem_chain")
	symaddr(&addr_QGraphicsTextItem_new, "QGraphicsTextItem_new")
	symaddr(&addr_QGraphicsTextItem_new2, "QGraphicsTextItem_new2")
	symaddr(&addr_QGraphicsTextItem_new3, "QGraphicsTextItem_new3")
	symaddr(&addr_QGraphicsTextItem_new4, "QGraphicsTextItem_new4")
	symaddr(&addr_QGraphicsTextItem_tr, "QGraphicsTextItem_tr")
	symaddr(&addr_QGraphicsTextItem_toHtml, "QGraphicsTextItem_toHtml")
	symaddr(&addr_QGraphicsTextItem_setHtml, "QGraphicsTextItem_setHtml")
	symaddr(&addr_QGraphicsTextItem_toPlainText, "QGraphicsTextItem_toPlainText")
	symaddr(&addr_QGraphicsTextItem_setPlainText, "QGraphicsTextItem_setPlainText")
	symaddr(&addr_QGraphicsTextItem_font, "QGraphicsTextItem_font")
	symaddr(&addr_QGraphicsTextItem_setFont, "QGraphicsTextItem_setFont")
	symaddr(&addr_QGraphicsTextItem_setDefaultTextColor, "QGraphicsTextItem_setDefaultTextColor")
	symaddr(&addr_QGraphicsTextItem_defaultTextColor, "QGraphicsTextItem_defaultTextColor")
	reg(&pg_QGraphicsTextItem_setTextWidth, "QGraphicsTextItem_setTextWidth")
	reg(&pg_QGraphicsTextItem_textWidth, "QGraphicsTextItem_textWidth")
	symaddr(&addr_QGraphicsTextItem_adjustSize, "QGraphicsTextItem_adjustSize")
	symaddr(&addr_QGraphicsTextItem_setDocument, "QGraphicsTextItem_setDocument")
	symaddr(&addr_QGraphicsTextItem_document, "QGraphicsTextItem_document")
	symaddr(&addr_QGraphicsTextItem_setTextInteractionFlags, "QGraphicsTextItem_setTextInteractionFlags")
	symaddr(&addr_QGraphicsTextItem_textInteractionFlags, "QGraphicsTextItem_textInteractionFlags")
	symaddr(&addr_QGraphicsTextItem_setTabChangesFocus, "QGraphicsTextItem_setTabChangesFocus")
	symaddr(&addr_QGraphicsTextItem_tabChangesFocus, "QGraphicsTextItem_tabChangesFocus")
	symaddr(&addr_QGraphicsTextItem_setOpenExternalLinks, "QGraphicsTextItem_setOpenExternalLinks")
	symaddr(&addr_QGraphicsTextItem_openExternalLinks, "QGraphicsTextItem_openExternalLinks")
	symaddr(&addr_QGraphicsTextItem_setTextCursor, "QGraphicsTextItem_setTextCursor")
	symaddr(&addr_QGraphicsTextItem_textCursor, "QGraphicsTextItem_textCursor")
	symaddr(&addr_QGraphicsTextItem_linkActivated, "QGraphicsTextItem_linkActivated")
	symaddr(&addr_QGraphicsTextItem_linkHovered, "QGraphicsTextItem_linkHovered")
	symaddr(&addr_QGraphicsTextItem_updateMicroFocus, "QGraphicsTextItem_updateMicroFocus")
	symaddr(&addr_QGraphicsTextItem_connect, "QGraphicsTextItem_connect")
	symaddr(&addr_QGraphicsTextItem_connect2, "QGraphicsTextItem_connect2")
	symaddr(&addr_QGraphicsTextItem_sender, "QGraphicsTextItem_sender")
	symaddr(&addr_QGraphicsTextItem_senderSignalIndex, "QGraphicsTextItem_senderSignalIndex")
	symaddr(&addr_QGraphicsTextItem_receivers, "QGraphicsTextItem_receivers")
	symaddr(&addr_QGraphicsTextItem_isSignalConnected, "QGraphicsTextItem_isSignalConnected")
	symaddr(&addr_QGraphicsTextItem_addToIndex, "QGraphicsTextItem_addToIndex")
	symaddr(&addr_QGraphicsTextItem_removeFromIndex, "QGraphicsTextItem_removeFromIndex")
	symaddr(&addr_QGraphicsTextItem_prepareGeometryChange, "QGraphicsTextItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsTextItem_tr2, "QGraphicsTextItem_tr2")
	symaddr(&addr_QGraphicsTextItem_tr3, "QGraphicsTextItem_tr3")
	symaddr(&addr_QGraphicsTextItem_connect3, "QGraphicsTextItem_connect3")
	symaddr(&addr_QGraphicsTextItem_connect4, "QGraphicsTextItem_connect4")
	symaddr(&addr_QGraphicsTextItem_metaObject_default, "QGraphicsTextItem_metaObject_default")
	symaddr(&addr_QGraphicsTextItem_boundingRect_default, "QGraphicsTextItem_boundingRect_default")
	symaddr(&addr_QGraphicsTextItem_shape_default, "QGraphicsTextItem_shape_default")
	symaddr(&addr_QGraphicsTextItem_contains_default, "QGraphicsTextItem_contains_default")
	symaddr(&addr_QGraphicsTextItem_paint_default, "QGraphicsTextItem_paint_default")
	symaddr(&addr_QGraphicsTextItem_isObscuredBy_default, "QGraphicsTextItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsTextItem_opaqueArea_default, "QGraphicsTextItem_opaqueArea_default")
	symaddr(&addr_QGraphicsTextItem_type_default, "QGraphicsTextItem_type_default")
	symaddr(&addr_QGraphicsTextItem_sceneEvent_default, "QGraphicsTextItem_sceneEvent_default")
	symaddr(&addr_QGraphicsTextItem_mousePressEvent_default, "QGraphicsTextItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsTextItem_mouseMoveEvent_default, "QGraphicsTextItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsTextItem_mouseReleaseEvent_default, "QGraphicsTextItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsTextItem_mouseDoubleClickEvent_default, "QGraphicsTextItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsTextItem_contextMenuEvent_default, "QGraphicsTextItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsTextItem_keyPressEvent_default, "QGraphicsTextItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsTextItem_keyReleaseEvent_default, "QGraphicsTextItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsTextItem_focusInEvent_default, "QGraphicsTextItem_focusInEvent_default")
	symaddr(&addr_QGraphicsTextItem_focusOutEvent_default, "QGraphicsTextItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsTextItem_dragEnterEvent_default, "QGraphicsTextItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsTextItem_dragLeaveEvent_default, "QGraphicsTextItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsTextItem_dragMoveEvent_default, "QGraphicsTextItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsTextItem_dropEvent_default, "QGraphicsTextItem_dropEvent_default")
	symaddr(&addr_QGraphicsTextItem_inputMethodEvent_default, "QGraphicsTextItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsTextItem_hoverEnterEvent_default, "QGraphicsTextItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsTextItem_hoverMoveEvent_default, "QGraphicsTextItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsTextItem_hoverLeaveEvent_default, "QGraphicsTextItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsTextItem_inputMethodQuery_default, "QGraphicsTextItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsTextItem_supportsExtension_default, "QGraphicsTextItem_supportsExtension_default")
	symaddr(&addr_QGraphicsTextItem_setExtension_default, "QGraphicsTextItem_setExtension_default")
	symaddr(&addr_QGraphicsTextItem_extension_default, "QGraphicsTextItem_extension_default")
	symaddr(&addr_QGraphicsTextItem_event_default, "QGraphicsTextItem_event_default")
	symaddr(&addr_QGraphicsTextItem_eventFilter_default, "QGraphicsTextItem_eventFilter_default")
	symaddr(&addr_QGraphicsTextItem_timerEvent_default, "QGraphicsTextItem_timerEvent_default")
	symaddr(&addr_QGraphicsTextItem_childEvent_default, "QGraphicsTextItem_childEvent_default")
	symaddr(&addr_QGraphicsTextItem_customEvent_default, "QGraphicsTextItem_customEvent_default")
	symaddr(&addr_QGraphicsTextItem_connectNotify_default, "QGraphicsTextItem_connectNotify_default")
	symaddr(&addr_QGraphicsTextItem_disconnectNotify_default, "QGraphicsTextItem_disconnectNotify_default")
	symaddr(&addr_QGraphicsTextItem_advance_default, "QGraphicsTextItem_advance_default")
	symaddr(&addr_QGraphicsTextItem_collidesWithItem_default, "QGraphicsTextItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsTextItem_collidesWithPath_default, "QGraphicsTextItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsTextItem_sceneEventFilter_default, "QGraphicsTextItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsTextItem_wheelEvent_default, "QGraphicsTextItem_wheelEvent_default")
	symaddr(&addr_QGraphicsTextItem_itemChange_default, "QGraphicsTextItem_itemChange_default")
	reg(&pg_QGraphicsTextItem_connect_linkActivated, "QGraphicsTextItem_connect_linkActivated")
	reg(&pg_QGraphicsTextItem_connect_linkHovered, "QGraphicsTextItem_connect_linkHovered")
	reg(&pg_QGraphicsTextItem_connect_parentChanged, "QGraphicsTextItem_connect_parentChanged")
	reg(&pg_QGraphicsTextItem_connect_opacityChanged, "QGraphicsTextItem_connect_opacityChanged")
	reg(&pg_QGraphicsTextItem_connect_visibleChanged, "QGraphicsTextItem_connect_visibleChanged")
	reg(&pg_QGraphicsTextItem_connect_enabledChanged, "QGraphicsTextItem_connect_enabledChanged")
	reg(&pg_QGraphicsTextItem_connect_xChanged, "QGraphicsTextItem_connect_xChanged")
	reg(&pg_QGraphicsTextItem_connect_yChanged, "QGraphicsTextItem_connect_yChanged")
	reg(&pg_QGraphicsTextItem_connect_zChanged, "QGraphicsTextItem_connect_zChanged")
	reg(&pg_QGraphicsTextItem_connect_rotationChanged, "QGraphicsTextItem_connect_rotationChanged")
	reg(&pg_QGraphicsTextItem_connect_scaleChanged, "QGraphicsTextItem_connect_scaleChanged")
	reg(&pg_QGraphicsTextItem_connect_childrenChanged, "QGraphicsTextItem_connect_childrenChanged")
	reg(&pg_QGraphicsTextItem_connect_widthChanged, "QGraphicsTextItem_connect_widthChanged")
	reg(&pg_QGraphicsTextItem_connect_heightChanged, "QGraphicsTextItem_connect_heightChanged")
	reg(&pg_QGraphicsTextItem_connect_destroyed, "QGraphicsTextItem_connect_destroyed")
	reg(&pg_QGraphicsTextItem_connect_destroyed2, "QGraphicsTextItem_connect_destroyed2")
	reg(&pg_QGraphicsTextItem_connect_objectNameChanged, "QGraphicsTextItem_connect_objectNameChanged")
	reg(&pg_QGraphicsTextItem__disconnect, "QGraphicsTextItem__disconnect")
	reg(&pg_QGraphicsTextItem__set_override, "QGraphicsTextItem__set_override")
	reg(&pg_QGraphicsTextItem_delete, "QGraphicsTextItem_delete")
}

type QGraphicsTextItem struct {
	h uintptr
	*QGraphicsObject
}

func bindQGraphicsTextItem(h uintptr) *QGraphicsTextItem {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsTextItem_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsTextItem{h: h, QGraphicsObject: &QGraphicsObject{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsTextItem(h uintptr) *QGraphicsTextItem { return bindQGraphicsTextItem(h) }

func (this *QGraphicsTextItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsTextItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsTextItem(h unsafe.Pointer) *QGraphicsTextItem { return bindQGraphicsTextItem(uintptr(h)) }
func UnsafeNewQGraphicsTextItemOwned(h unsafe.Pointer) *QGraphicsTextItem { return newownedQGraphicsTextItem(uintptr(h)) }

func (this *QGraphicsTextItem) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsObject != nil {
		this.QGraphicsObject.UnsafeRelease()
	}
}

func NewQGraphicsTextItem() *QGraphicsTextItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_new, uintptr(__goid))
	__o := newownedQGraphicsTextItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsTextItem_virtIdx, QGraphicsTextItem_enable)
	return __o
}

func NewQGraphicsTextItem2(text string) *QGraphicsTextItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQGraphicsTextItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsTextItem_virtIdx, QGraphicsTextItem_enable)
	return __o
}

func NewQGraphicsTextItem3(parent *QGraphicsItem) *QGraphicsTextItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsTextItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsTextItem_virtIdx, QGraphicsTextItem_enable)
	return __o
}

func NewQGraphicsTextItem4(text string, parent *QGraphicsItem) *QGraphicsTextItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsTextItem(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsTextItem_virtIdx, QGraphicsTextItem_enable)
	return __o
}

func QGraphicsTextItem_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsTextItem) ToHtml() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_toHtml, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsTextItem) SetHtml(html string) {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	purego.SyscallN(addr_QGraphicsTextItem_setHtml, this.h, html_p, uintptr(len(html)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(html)
}

func (this *QGraphicsTextItem) ToPlainText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_toPlainText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsTextItem) SetPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QGraphicsTextItem_setPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QGraphicsTextItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QGraphicsTextItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QGraphicsTextItem_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QGraphicsTextItem) SetDefaultTextColor(c *QColor) {
	purego.SyscallN(addr_QGraphicsTextItem_setDefaultTextColor, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QGraphicsTextItem) DefaultTextColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_defaultTextColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QGraphicsTextItem) SetTextWidth(width float64) {
	pg_QGraphicsTextItem_setTextWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) TextWidth() float64 {
	__r := pg_QGraphicsTextItem_textWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsTextItem) AdjustSize() {
	purego.SyscallN(addr_QGraphicsTextItem_adjustSize, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) SetDocument(document *QTextDocument) {
	purego.SyscallN(addr_QGraphicsTextItem_setDocument, this.h, uintptr(document.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(document)
}

func (this *QGraphicsTextItem) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QGraphicsTextItem) SetTextInteractionFlags(flags TextInteractionFlag) {
	purego.SyscallN(addr_QGraphicsTextItem_setTextInteractionFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) TextInteractionFlags() TextInteractionFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_textInteractionFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextInteractionFlag(__r)
}

func (this *QGraphicsTextItem) SetTabChangesFocus(b bool) {
	purego.SyscallN(addr_QGraphicsTextItem_setTabChangesFocus, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) TabChangesFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_tabChangesFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsTextItem) SetOpenExternalLinks(open bool) {
	purego.SyscallN(addr_QGraphicsTextItem_setOpenExternalLinks, this.h, uintptr(b2i(open)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) OpenExternalLinks() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_openExternalLinks, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsTextItem) SetTextCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QGraphicsTextItem_setTextCursor, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QGraphicsTextItem) TextCursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_textCursor, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QGraphicsTextItem) LinkActivated(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QGraphicsTextItem_linkActivated, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGraphicsTextItem) LinkHovered(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QGraphicsTextItem_linkHovered, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGraphicsTextItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsTextItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsTextItem_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsTextItem_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsTextItem) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsTextItem) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsTextItem) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsTextItem) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QGraphicsTextItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsTextItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsTextItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsTextItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsTextItem_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsTextItem_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsTextItem_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsTextItem_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsTextItem) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsTextItem) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsTextItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsTextItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsTextItem) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsTextItem_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsTextItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsTextItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsTextItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsTextItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsTextItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsTextItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsTextItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsTextItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsTextItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsTextItem) EventDefault(ev *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_event_default, this.h, uintptr(ev.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
	return __r != 0
}

func (this *QGraphicsTextItem) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsTextItem) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsTextItem_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsTextItem) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsTextItem_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsTextItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsTextItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsTextItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsTextItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsTextItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsTextItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsTextItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsTextItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsTextItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsTextItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsTextItem) OnLinkActivated(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QGraphicsTextItem_connect_linkActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnLinkHovered(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QGraphicsTextItem_connect_linkHovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnParentChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_parentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnOpacityChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnEnabledChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnXChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnYChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnZChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_zChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnRotationChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_rotationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_scaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnChildrenChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_childrenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnWidthChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnHeightChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsTextItem_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsTextItem_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsTextItem_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsTextItem) Disconnect(handle int64) {
	pg_QGraphicsTextItem__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsTextItem_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 2510008782785038831:
		return 1
	case 7938019069247444294:
		return 2
	case 5645366511944179768:
		return 3
	case 6083584261053556819:
		return 4
	case -7215030271914903671:
		return 5
	case 399370970023989435:
		return 6
	case -6371403872082420691:
		return 7
	case -2846371507838587071:
		return 8
	case -332090898078543319:
		return 9
	case -4255494004099806581:
		return 10
	case -4415920049297795213:
		return 11
	case -4029073905697023887:
		return 12
	case -6698043222227006015:
		return 13
	case 1023791977679793185:
		return 14
	case 772598475950050299:
		return 15
	case 975043767033552552:
		return 16
	case -385596010576048993:
		return 17
	case 206122715919917589:
		return 18
	case -3139904158373725006:
		return 19
	case -1781545137352745978:
		return 20
	case 9107058682754591716:
		return 21
	case 5438819661942843222:
		return 22
	case 6932941943906681679:
		return 23
	case -2745191522144316396:
		return 24
	case 7944408482203603984:
		return 25
	case 7345853097340709638:
		return 26
	case -5729166443651306418:
		return 27
	case -6730181576106171580:
		return 28
	case 2270616358835251226:
		return 29
	case 1237236456312824991:
		return 30
	case 5217623514630858297:
		return 31
	case -5409927581812374498:
		return 32
	case -1185315367526857833:
		return 33
	case 3769868229123868474:
		return 34
	case 7379029982170503008:
		return 35
	case -568369772982215962:
		return 36
	case -4830836341480073475:
		return 37
	case 6469736184362593173:
		return 38
	case 6636575182718753853:
		return 39
	case -4584203218241215725:
		return 40
	case 1919988340934511316:
		return 41
	case -1888572168794878830:
		return 42
	}
	return -1
}

func QGraphicsTextItem_enable(h uintptr, idx int32) { pg_QGraphicsTextItem__set_override(h, idx) }

func (this *QGraphicsTextItem) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsTextItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsTextItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsTextItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsTextItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsTextItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsTextItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsTextItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsTextItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsTextItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsTextItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsTextItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsTextItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsTextItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsTextItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsTextItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsTextItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsTextItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsTextItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsTextItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsTextItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsTextItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsTextItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsTextItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsTextItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsTextItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsTextItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsTextItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsTextItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsTextItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsTextItem) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(ev uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(ev)))) }) // slot "event"
}

func (this *QGraphicsTextItem) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsTextItem) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsTextItem) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsTextItem) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsTextItem) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsTextItem) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsTextItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsTextItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsTextItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsTextItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsTextItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsTextItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsTextItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsTextItem_delete(this.h)
	}
	this.UnsafeRelease()
}
