package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsProxyWidget_chain uintptr
	addr_QGraphicsProxyWidget_new uintptr
	addr_QGraphicsProxyWidget_new2 uintptr
	addr_QGraphicsProxyWidget_new3 uintptr
	addr_QGraphicsProxyWidget_tr uintptr
	addr_QGraphicsProxyWidget_setWidget uintptr
	addr_QGraphicsProxyWidget_widget uintptr
	addr_QGraphicsProxyWidget_subWidgetRect uintptr
	addr_QGraphicsProxyWidget_createProxyForChildWidget uintptr
	addr_QGraphicsProxyWidget_newProxyWidget uintptr
	addr_QGraphicsProxyWidget_setTabOrder uintptr
	addr_QGraphicsProxyWidget_updateMicroFocus uintptr
	addr_QGraphicsProxyWidget_connect uintptr
	addr_QGraphicsProxyWidget_connect2 uintptr
	addr_QGraphicsProxyWidget_sender uintptr
	addr_QGraphicsProxyWidget_senderSignalIndex uintptr
	addr_QGraphicsProxyWidget_receivers uintptr
	addr_QGraphicsProxyWidget_isSignalConnected uintptr
	addr_QGraphicsProxyWidget_addToIndex uintptr
	addr_QGraphicsProxyWidget_removeFromIndex uintptr
	addr_QGraphicsProxyWidget_prepareGeometryChange uintptr
	addr_QGraphicsProxyWidget_setGraphicsItem uintptr
	addr_QGraphicsProxyWidget_setOwnedByLayout uintptr
	addr_QGraphicsProxyWidget_tr2 uintptr
	addr_QGraphicsProxyWidget_tr3 uintptr
	addr_QGraphicsProxyWidget_connect3 uintptr
	addr_QGraphicsProxyWidget_connect4 uintptr
	addr_QGraphicsProxyWidget_metaObject_default uintptr
	addr_QGraphicsProxyWidget_setGeometry_default uintptr
	addr_QGraphicsProxyWidget_paint_default uintptr
	addr_QGraphicsProxyWidget_type_default uintptr
	addr_QGraphicsProxyWidget_itemChange_default uintptr
	addr_QGraphicsProxyWidget_event_default uintptr
	addr_QGraphicsProxyWidget_eventFilter_default uintptr
	addr_QGraphicsProxyWidget_showEvent_default uintptr
	addr_QGraphicsProxyWidget_hideEvent_default uintptr
	addr_QGraphicsProxyWidget_contextMenuEvent_default uintptr
	addr_QGraphicsProxyWidget_dragEnterEvent_default uintptr
	addr_QGraphicsProxyWidget_dragLeaveEvent_default uintptr
	addr_QGraphicsProxyWidget_dragMoveEvent_default uintptr
	addr_QGraphicsProxyWidget_dropEvent_default uintptr
	addr_QGraphicsProxyWidget_hoverEnterEvent_default uintptr
	addr_QGraphicsProxyWidget_hoverLeaveEvent_default uintptr
	addr_QGraphicsProxyWidget_hoverMoveEvent_default uintptr
	addr_QGraphicsProxyWidget_grabMouseEvent_default uintptr
	addr_QGraphicsProxyWidget_ungrabMouseEvent_default uintptr
	addr_QGraphicsProxyWidget_mouseMoveEvent_default uintptr
	addr_QGraphicsProxyWidget_mousePressEvent_default uintptr
	addr_QGraphicsProxyWidget_mouseReleaseEvent_default uintptr
	addr_QGraphicsProxyWidget_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsProxyWidget_wheelEvent_default uintptr
	addr_QGraphicsProxyWidget_keyPressEvent_default uintptr
	addr_QGraphicsProxyWidget_keyReleaseEvent_default uintptr
	addr_QGraphicsProxyWidget_focusInEvent_default uintptr
	addr_QGraphicsProxyWidget_focusOutEvent_default uintptr
	addr_QGraphicsProxyWidget_focusNextPrevChild_default uintptr
	addr_QGraphicsProxyWidget_inputMethodQuery_default uintptr
	addr_QGraphicsProxyWidget_inputMethodEvent_default uintptr
	addr_QGraphicsProxyWidget_sizeHint_default uintptr
	addr_QGraphicsProxyWidget_resizeEvent_default uintptr
	addr_QGraphicsProxyWidget_getContentsMargins_default uintptr
	addr_QGraphicsProxyWidget_paintWindowFrame_default uintptr
	addr_QGraphicsProxyWidget_boundingRect_default uintptr
	addr_QGraphicsProxyWidget_shape_default uintptr
	addr_QGraphicsProxyWidget_initStyleOption_default uintptr
	addr_QGraphicsProxyWidget_updateGeometry_default uintptr
	addr_QGraphicsProxyWidget_propertyChange_default uintptr
	addr_QGraphicsProxyWidget_sceneEvent_default uintptr
	addr_QGraphicsProxyWidget_windowFrameEvent_default uintptr
	addr_QGraphicsProxyWidget_windowFrameSectionAt_default uintptr
	addr_QGraphicsProxyWidget_changeEvent_default uintptr
	addr_QGraphicsProxyWidget_closeEvent_default uintptr
	addr_QGraphicsProxyWidget_moveEvent_default uintptr
	addr_QGraphicsProxyWidget_polishEvent_default uintptr
	addr_QGraphicsProxyWidget_grabKeyboardEvent_default uintptr
	addr_QGraphicsProxyWidget_ungrabKeyboardEvent_default uintptr
	addr_QGraphicsProxyWidget_timerEvent_default uintptr
	addr_QGraphicsProxyWidget_childEvent_default uintptr
	addr_QGraphicsProxyWidget_customEvent_default uintptr
	addr_QGraphicsProxyWidget_connectNotify_default uintptr
	addr_QGraphicsProxyWidget_disconnectNotify_default uintptr
	addr_QGraphicsProxyWidget_advance_default uintptr
	addr_QGraphicsProxyWidget_contains_default uintptr
	addr_QGraphicsProxyWidget_collidesWithItem_default uintptr
	addr_QGraphicsProxyWidget_collidesWithPath_default uintptr
	addr_QGraphicsProxyWidget_isObscuredBy_default uintptr
	addr_QGraphicsProxyWidget_opaqueArea_default uintptr
	addr_QGraphicsProxyWidget_sceneEventFilter_default uintptr
	addr_QGraphicsProxyWidget_supportsExtension_default uintptr
	addr_QGraphicsProxyWidget_setExtension_default uintptr
	addr_QGraphicsProxyWidget_extension_default uintptr
	addr_QGraphicsProxyWidget_isEmpty_default uintptr
	pg_QGraphicsProxyWidget_connect_geometryChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_layoutChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_parentChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_opacityChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_visibleChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_xChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_yChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_zChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_rotationChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_scaleChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_childrenChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_widthChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_heightChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_destroyed func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsProxyWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsProxyWidget__disconnect func(int64)
	pg_QGraphicsProxyWidget__set_override func(uintptr, int32)
	pg_QGraphicsProxyWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsProxyWidget_chain, "QGraphicsProxyWidget_chain")
	symaddr(&addr_QGraphicsProxyWidget_new, "QGraphicsProxyWidget_new")
	symaddr(&addr_QGraphicsProxyWidget_new2, "QGraphicsProxyWidget_new2")
	symaddr(&addr_QGraphicsProxyWidget_new3, "QGraphicsProxyWidget_new3")
	symaddr(&addr_QGraphicsProxyWidget_tr, "QGraphicsProxyWidget_tr")
	symaddr(&addr_QGraphicsProxyWidget_setWidget, "QGraphicsProxyWidget_setWidget")
	symaddr(&addr_QGraphicsProxyWidget_widget, "QGraphicsProxyWidget_widget")
	symaddr(&addr_QGraphicsProxyWidget_subWidgetRect, "QGraphicsProxyWidget_subWidgetRect")
	symaddr(&addr_QGraphicsProxyWidget_createProxyForChildWidget, "QGraphicsProxyWidget_createProxyForChildWidget")
	symaddr(&addr_QGraphicsProxyWidget_newProxyWidget, "QGraphicsProxyWidget_newProxyWidget")
	symaddr(&addr_QGraphicsProxyWidget_setTabOrder, "QGraphicsProxyWidget_setTabOrder")
	symaddr(&addr_QGraphicsProxyWidget_updateMicroFocus, "QGraphicsProxyWidget_updateMicroFocus")
	symaddr(&addr_QGraphicsProxyWidget_connect, "QGraphicsProxyWidget_connect")
	symaddr(&addr_QGraphicsProxyWidget_connect2, "QGraphicsProxyWidget_connect2")
	symaddr(&addr_QGraphicsProxyWidget_sender, "QGraphicsProxyWidget_sender")
	symaddr(&addr_QGraphicsProxyWidget_senderSignalIndex, "QGraphicsProxyWidget_senderSignalIndex")
	symaddr(&addr_QGraphicsProxyWidget_receivers, "QGraphicsProxyWidget_receivers")
	symaddr(&addr_QGraphicsProxyWidget_isSignalConnected, "QGraphicsProxyWidget_isSignalConnected")
	symaddr(&addr_QGraphicsProxyWidget_addToIndex, "QGraphicsProxyWidget_addToIndex")
	symaddr(&addr_QGraphicsProxyWidget_removeFromIndex, "QGraphicsProxyWidget_removeFromIndex")
	symaddr(&addr_QGraphicsProxyWidget_prepareGeometryChange, "QGraphicsProxyWidget_prepareGeometryChange")
	symaddr(&addr_QGraphicsProxyWidget_setGraphicsItem, "QGraphicsProxyWidget_setGraphicsItem")
	symaddr(&addr_QGraphicsProxyWidget_setOwnedByLayout, "QGraphicsProxyWidget_setOwnedByLayout")
	symaddr(&addr_QGraphicsProxyWidget_tr2, "QGraphicsProxyWidget_tr2")
	symaddr(&addr_QGraphicsProxyWidget_tr3, "QGraphicsProxyWidget_tr3")
	symaddr(&addr_QGraphicsProxyWidget_connect3, "QGraphicsProxyWidget_connect3")
	symaddr(&addr_QGraphicsProxyWidget_connect4, "QGraphicsProxyWidget_connect4")
	symaddr(&addr_QGraphicsProxyWidget_metaObject_default, "QGraphicsProxyWidget_metaObject_default")
	symaddr(&addr_QGraphicsProxyWidget_setGeometry_default, "QGraphicsProxyWidget_setGeometry_default")
	symaddr(&addr_QGraphicsProxyWidget_paint_default, "QGraphicsProxyWidget_paint_default")
	symaddr(&addr_QGraphicsProxyWidget_type_default, "QGraphicsProxyWidget_type_default")
	symaddr(&addr_QGraphicsProxyWidget_itemChange_default, "QGraphicsProxyWidget_itemChange_default")
	symaddr(&addr_QGraphicsProxyWidget_event_default, "QGraphicsProxyWidget_event_default")
	symaddr(&addr_QGraphicsProxyWidget_eventFilter_default, "QGraphicsProxyWidget_eventFilter_default")
	symaddr(&addr_QGraphicsProxyWidget_showEvent_default, "QGraphicsProxyWidget_showEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_hideEvent_default, "QGraphicsProxyWidget_hideEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_contextMenuEvent_default, "QGraphicsProxyWidget_contextMenuEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_dragEnterEvent_default, "QGraphicsProxyWidget_dragEnterEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_dragLeaveEvent_default, "QGraphicsProxyWidget_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_dragMoveEvent_default, "QGraphicsProxyWidget_dragMoveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_dropEvent_default, "QGraphicsProxyWidget_dropEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_hoverEnterEvent_default, "QGraphicsProxyWidget_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_hoverLeaveEvent_default, "QGraphicsProxyWidget_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_hoverMoveEvent_default, "QGraphicsProxyWidget_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_grabMouseEvent_default, "QGraphicsProxyWidget_grabMouseEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_ungrabMouseEvent_default, "QGraphicsProxyWidget_ungrabMouseEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_mouseMoveEvent_default, "QGraphicsProxyWidget_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_mousePressEvent_default, "QGraphicsProxyWidget_mousePressEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_mouseReleaseEvent_default, "QGraphicsProxyWidget_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_mouseDoubleClickEvent_default, "QGraphicsProxyWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_wheelEvent_default, "QGraphicsProxyWidget_wheelEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_keyPressEvent_default, "QGraphicsProxyWidget_keyPressEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_keyReleaseEvent_default, "QGraphicsProxyWidget_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_focusInEvent_default, "QGraphicsProxyWidget_focusInEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_focusOutEvent_default, "QGraphicsProxyWidget_focusOutEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_focusNextPrevChild_default, "QGraphicsProxyWidget_focusNextPrevChild_default")
	symaddr(&addr_QGraphicsProxyWidget_inputMethodQuery_default, "QGraphicsProxyWidget_inputMethodQuery_default")
	symaddr(&addr_QGraphicsProxyWidget_inputMethodEvent_default, "QGraphicsProxyWidget_inputMethodEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_sizeHint_default, "QGraphicsProxyWidget_sizeHint_default")
	symaddr(&addr_QGraphicsProxyWidget_resizeEvent_default, "QGraphicsProxyWidget_resizeEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_getContentsMargins_default, "QGraphicsProxyWidget_getContentsMargins_default")
	symaddr(&addr_QGraphicsProxyWidget_paintWindowFrame_default, "QGraphicsProxyWidget_paintWindowFrame_default")
	symaddr(&addr_QGraphicsProxyWidget_boundingRect_default, "QGraphicsProxyWidget_boundingRect_default")
	symaddr(&addr_QGraphicsProxyWidget_shape_default, "QGraphicsProxyWidget_shape_default")
	symaddr(&addr_QGraphicsProxyWidget_initStyleOption_default, "QGraphicsProxyWidget_initStyleOption_default")
	symaddr(&addr_QGraphicsProxyWidget_updateGeometry_default, "QGraphicsProxyWidget_updateGeometry_default")
	symaddr(&addr_QGraphicsProxyWidget_propertyChange_default, "QGraphicsProxyWidget_propertyChange_default")
	symaddr(&addr_QGraphicsProxyWidget_sceneEvent_default, "QGraphicsProxyWidget_sceneEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_windowFrameEvent_default, "QGraphicsProxyWidget_windowFrameEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_windowFrameSectionAt_default, "QGraphicsProxyWidget_windowFrameSectionAt_default")
	symaddr(&addr_QGraphicsProxyWidget_changeEvent_default, "QGraphicsProxyWidget_changeEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_closeEvent_default, "QGraphicsProxyWidget_closeEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_moveEvent_default, "QGraphicsProxyWidget_moveEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_polishEvent_default, "QGraphicsProxyWidget_polishEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_grabKeyboardEvent_default, "QGraphicsProxyWidget_grabKeyboardEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_ungrabKeyboardEvent_default, "QGraphicsProxyWidget_ungrabKeyboardEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_timerEvent_default, "QGraphicsProxyWidget_timerEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_childEvent_default, "QGraphicsProxyWidget_childEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_customEvent_default, "QGraphicsProxyWidget_customEvent_default")
	symaddr(&addr_QGraphicsProxyWidget_connectNotify_default, "QGraphicsProxyWidget_connectNotify_default")
	symaddr(&addr_QGraphicsProxyWidget_disconnectNotify_default, "QGraphicsProxyWidget_disconnectNotify_default")
	symaddr(&addr_QGraphicsProxyWidget_advance_default, "QGraphicsProxyWidget_advance_default")
	symaddr(&addr_QGraphicsProxyWidget_contains_default, "QGraphicsProxyWidget_contains_default")
	symaddr(&addr_QGraphicsProxyWidget_collidesWithItem_default, "QGraphicsProxyWidget_collidesWithItem_default")
	symaddr(&addr_QGraphicsProxyWidget_collidesWithPath_default, "QGraphicsProxyWidget_collidesWithPath_default")
	symaddr(&addr_QGraphicsProxyWidget_isObscuredBy_default, "QGraphicsProxyWidget_isObscuredBy_default")
	symaddr(&addr_QGraphicsProxyWidget_opaqueArea_default, "QGraphicsProxyWidget_opaqueArea_default")
	symaddr(&addr_QGraphicsProxyWidget_sceneEventFilter_default, "QGraphicsProxyWidget_sceneEventFilter_default")
	symaddr(&addr_QGraphicsProxyWidget_supportsExtension_default, "QGraphicsProxyWidget_supportsExtension_default")
	symaddr(&addr_QGraphicsProxyWidget_setExtension_default, "QGraphicsProxyWidget_setExtension_default")
	symaddr(&addr_QGraphicsProxyWidget_extension_default, "QGraphicsProxyWidget_extension_default")
	symaddr(&addr_QGraphicsProxyWidget_isEmpty_default, "QGraphicsProxyWidget_isEmpty_default")
	reg(&pg_QGraphicsProxyWidget_connect_geometryChanged, "QGraphicsProxyWidget_connect_geometryChanged")
	reg(&pg_QGraphicsProxyWidget_connect_layoutChanged, "QGraphicsProxyWidget_connect_layoutChanged")
	reg(&pg_QGraphicsProxyWidget_connect_parentChanged, "QGraphicsProxyWidget_connect_parentChanged")
	reg(&pg_QGraphicsProxyWidget_connect_opacityChanged, "QGraphicsProxyWidget_connect_opacityChanged")
	reg(&pg_QGraphicsProxyWidget_connect_visibleChanged, "QGraphicsProxyWidget_connect_visibleChanged")
	reg(&pg_QGraphicsProxyWidget_connect_enabledChanged, "QGraphicsProxyWidget_connect_enabledChanged")
	reg(&pg_QGraphicsProxyWidget_connect_xChanged, "QGraphicsProxyWidget_connect_xChanged")
	reg(&pg_QGraphicsProxyWidget_connect_yChanged, "QGraphicsProxyWidget_connect_yChanged")
	reg(&pg_QGraphicsProxyWidget_connect_zChanged, "QGraphicsProxyWidget_connect_zChanged")
	reg(&pg_QGraphicsProxyWidget_connect_rotationChanged, "QGraphicsProxyWidget_connect_rotationChanged")
	reg(&pg_QGraphicsProxyWidget_connect_scaleChanged, "QGraphicsProxyWidget_connect_scaleChanged")
	reg(&pg_QGraphicsProxyWidget_connect_childrenChanged, "QGraphicsProxyWidget_connect_childrenChanged")
	reg(&pg_QGraphicsProxyWidget_connect_widthChanged, "QGraphicsProxyWidget_connect_widthChanged")
	reg(&pg_QGraphicsProxyWidget_connect_heightChanged, "QGraphicsProxyWidget_connect_heightChanged")
	reg(&pg_QGraphicsProxyWidget_connect_destroyed, "QGraphicsProxyWidget_connect_destroyed")
	reg(&pg_QGraphicsProxyWidget_connect_destroyed2, "QGraphicsProxyWidget_connect_destroyed2")
	reg(&pg_QGraphicsProxyWidget_connect_objectNameChanged, "QGraphicsProxyWidget_connect_objectNameChanged")
	reg(&pg_QGraphicsProxyWidget__disconnect, "QGraphicsProxyWidget__disconnect")
	reg(&pg_QGraphicsProxyWidget__set_override, "QGraphicsProxyWidget__set_override")
	reg(&pg_QGraphicsProxyWidget_delete, "QGraphicsProxyWidget_delete")
}

type QGraphicsProxyWidget struct {
	h uintptr
	*QGraphicsWidget
}

func bindQGraphicsProxyWidget(h uintptr) *QGraphicsProxyWidget {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QGraphicsProxyWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsProxyWidget{h: h, QGraphicsWidget: &QGraphicsWidget{h: __c[0], QGraphicsObject: &QGraphicsObject{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQGraphicsProxyWidget(h uintptr) *QGraphicsProxyWidget { return bindQGraphicsProxyWidget(h) }

func (this *QGraphicsProxyWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsProxyWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsProxyWidget(h unsafe.Pointer) *QGraphicsProxyWidget { return bindQGraphicsProxyWidget(uintptr(h)) }
func UnsafeNewQGraphicsProxyWidgetOwned(h unsafe.Pointer) *QGraphicsProxyWidget { return newownedQGraphicsProxyWidget(uintptr(h)) }

func (this *QGraphicsProxyWidget) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsWidget != nil {
		this.QGraphicsWidget.UnsafeRelease()
	}
}

func NewQGraphicsProxyWidget() *QGraphicsProxyWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_new, uintptr(__goid))
	__o := newownedQGraphicsProxyWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsProxyWidget_virtIdx, QGraphicsProxyWidget_enable)
	return __o
}

func NewQGraphicsProxyWidget2(parent *QGraphicsItem) *QGraphicsProxyWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsProxyWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsProxyWidget_virtIdx, QGraphicsProxyWidget_enable)
	return __o
}

func NewQGraphicsProxyWidget3(parent *QGraphicsItem, wFlags WindowType) *QGraphicsProxyWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(wFlags))
	__o := newownedQGraphicsProxyWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsProxyWidget_virtIdx, QGraphicsProxyWidget_enable)
	return __o
}

func QGraphicsProxyWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsProxyWidget) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsProxyWidget) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QGraphicsProxyWidget) SubWidgetRect(widget *QWidget) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_subWidgetRect, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return newownedQRectF(__r)
}

func (this *QGraphicsProxyWidget) CreateProxyForChildWidget(child *QWidget) *QGraphicsProxyWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_createProxyForChildWidget, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return bindQGraphicsProxyWidget(__r)
}

func (this *QGraphicsProxyWidget) NewProxyWidget(arg1 *QWidget) *QGraphicsProxyWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_newProxyWidget, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQGraphicsProxyWidget(__r)
}

func QGraphicsProxyWidget_SetTabOrder(first *QGraphicsWidget, second *QGraphicsWidget) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setTabOrder, uintptr(first.UnsafePointer()), uintptr(second.UnsafePointer()))
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)
}

func (this *QGraphicsProxyWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsProxyWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsProxyWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsProxyWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsProxyWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsProxyWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsProxyWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsProxyWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QGraphicsProxyWidget) AddToIndex() {
	purego.SyscallN(addr_QGraphicsProxyWidget_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsProxyWidget_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsProxyWidget_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsProxyWidget) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func QGraphicsProxyWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsProxyWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsProxyWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsProxyWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsProxyWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsProxyWidget) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsProxyWidget) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsProxyWidget_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsProxyWidget) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsProxyWidget) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsProxyWidget) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsProxyWidget) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsProxyWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) GrabMouseEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_grabMouseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) UngrabMouseEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_ungrabMouseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsProxyWidget) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsProxyWidget) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) SizeHintDefault(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_sizeHint_default, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsProxyWidget) ResizeEventDefault(event *QGraphicsSceneResizeEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsProxyWidget_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsProxyWidget) PaintWindowFrameDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsProxyWidget_paintWindowFrame_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsProxyWidget) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsProxyWidget) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsProxyWidget) InitStyleOptionDefault(option *QStyleOption) {
	purego.SyscallN(addr_QGraphicsProxyWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QGraphicsProxyWidget) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsProxyWidget_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) PropertyChangeDefault(propertyName string, value *QVariant) *QVariant {
	var propertyName_p uintptr
	if len(propertyName) > 0 {
		propertyName_p = uintptr(unsafe.Pointer(unsafe.StringData(propertyName)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_propertyChange_default, this.h, propertyName_p, uintptr(len(propertyName)), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(propertyName)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsProxyWidget) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsProxyWidget) WindowFrameEventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_windowFrameEvent_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QGraphicsProxyWidget) WindowFrameSectionAtDefault(pos *QPointF) WindowFrameSection {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_windowFrameSectionAt_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return WindowFrameSection(__r)
}

func (this *QGraphicsProxyWidget) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) MoveEventDefault(event *QGraphicsSceneMoveEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) PolishEventDefault() {
	purego.SyscallN(addr_QGraphicsProxyWidget_polishEvent_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) GrabKeyboardEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_grabKeyboardEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) UngrabKeyboardEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_ungrabKeyboardEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsProxyWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsProxyWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsProxyWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsProxyWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsProxyWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsProxyWidget) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsProxyWidget_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsProxyWidget) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsProxyWidget) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsProxyWidget) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsProxyWidget) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsProxyWidget) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsProxyWidget) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsProxyWidget) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsProxyWidget) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsProxyWidget_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsProxyWidget) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsProxyWidget) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsProxyWidget_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsProxyWidget) OnGeometryChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_geometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnParentChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_parentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnOpacityChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnEnabledChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnXChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnYChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnZChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_zChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnRotationChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_rotationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_scaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnChildrenChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_childrenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnWidthChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnHeightChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsProxyWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsProxyWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsProxyWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsProxyWidget) Disconnect(handle int64) {
	pg_QGraphicsProxyWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsProxyWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 5684820848726969227:
		return 1
	case 6083584261053556819:
		return 2
	case -6371403872082420691:
		return 3
	case -1888572168794878830:
		return 4
	case 1237236456312824991:
		return 5
	case 5217623514630858297:
		return 6
	case -6825369174751916:
		return 7
	case -199718707876719129:
		return 8
	case -6698043222227006015:
		return 9
	case 206122715919917589:
		return 10
	case -3139904158373725006:
		return 11
	case -1781545137352745978:
		return 12
	case 9107058682754591716:
		return 13
	case 6932941943906681679:
		return 14
	case 7944408482203603984:
		return 15
	case -2745191522144316396:
		return 16
	case 7928460550911621078:
		return 17
	case -5662911227992581871:
		return 18
	case -4255494004099806581:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case 1919988340934511316:
		return 23
	case 1023791977679793185:
		return 24
	case 772598475950050299:
		return 25
	case 975043767033552552:
		return 26
	case -385596010576048993:
		return 27
	case 5898328517505784191:
		return 28
	case 7345853097340709638:
		return 29
	case 5438819661942843222:
		return 30
	case 729690261457684101:
		return 31
	case -4302917285655014661:
		return 32
	case 164760286202646972:
		return 33
	case -1640122132753285616:
		return 34
	case 2510008782785038831:
		return 35
	case 7938019069247444294:
		return 36
	case -488693769771308531:
		return 37
	case 8082455443689927668:
		return 38
	case 5289089528078387296:
		return 39
	case -2846371507838587071:
		return 40
	case -7362375632910814166:
		return 41
	case 8834866722574844732:
		return 42
	case 4959376063826165943:
		return 43
	case -1989937050019464783:
		return 44
	case -9161949150489525396:
		return 45
	case -214704830098031580:
		return 46
	case 6423297148758141226:
		return 47
	case 2912496686457383679:
		return 48
	case -5409927581812374498:
		return 49
	case -1185315367526857833:
		return 50
	case 3769868229123868474:
		return 51
	case 7379029982170503008:
		return 52
	case -568369772982215962:
		return 53
	case -4830836341480073475:
		return 54
	case 5645366511944179768:
		return 55
	case 6469736184362593173:
		return 56
	case 6636575182718753853:
		return 57
	case -7215030271914903671:
		return 58
	case 399370970023989435:
		return 59
	case -4584203218241215725:
		return 60
	case -5729166443651306418:
		return 61
	case -6730181576106171580:
		return 62
	case 2270616358835251226:
		return 63
	case -626822430251098498:
		return 64
	}
	return -1
}

func QGraphicsProxyWidget_enable(h uintptr, idx int32) { pg_QGraphicsProxyWidget__set_override(h, idx) }

func (this *QGraphicsProxyWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsProxyWidget) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsProxyWidget) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsProxyWidget) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsProxyWidget) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsProxyWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsProxyWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsProxyWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QGraphicsProxyWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QGraphicsProxyWidget) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsProxyWidget) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsProxyWidget) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsProxyWidget) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsProxyWidget) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsProxyWidget) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsProxyWidget) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsProxyWidget) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsProxyWidget) OnGrabMouseEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 7928460550911621078, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "grabMouseEvent"
}

func (this *QGraphicsProxyWidget) OnUngrabMouseEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -5662911227992581871, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "ungrabMouseEvent"
}

func (this *QGraphicsProxyWidget) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsProxyWidget) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsProxyWidget) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsProxyWidget) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsProxyWidget) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsProxyWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsProxyWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsProxyWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsProxyWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsProxyWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QGraphicsProxyWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsProxyWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsProxyWidget) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsProxyWidget) OnResizeEvent(fn func(*QGraphicsSceneResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQGraphicsSceneResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QGraphicsProxyWidget) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsProxyWidget) OnPaintWindowFrame(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -1640122132753285616, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paintWindowFrame"
}

func (this *QGraphicsProxyWidget) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsProxyWidget) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsProxyWidget) OnInitStyleOption(fn func(*QStyleOption)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOption(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QGraphicsProxyWidget) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsProxyWidget) OnPropertyChange(fn func(string, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 5289089528078387296, func(propertyName uintptr, propertyName_len int64, value uintptr) uintptr { return uintptr(fn(pgStrNoFree(propertyName, propertyName_len), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "propertyChange"
}

func (this *QGraphicsProxyWidget) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsProxyWidget) OnWindowFrameEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -7362375632910814166, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "windowFrameEvent"
}

func (this *QGraphicsProxyWidget) OnWindowFrameSectionAt(fn func(*QPointF) WindowFrameSection) {
	registerVirt(this.UnsafeGoID(), 8834866722574844732, func(pos uintptr) int32 { return int32(fn(UnsafeNewQPointF(unsafe.Pointer(pos)))) }) // slot "windowFrameSectionAt"
}

func (this *QGraphicsProxyWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QGraphicsProxyWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QGraphicsProxyWidget) OnMoveEvent(fn func(*QGraphicsSceneMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QGraphicsProxyWidget) OnPolishEvent(fn func()) {
	registerVirt(this.UnsafeGoID(), -214704830098031580, func() { fn() }) // slot "polishEvent"
}

func (this *QGraphicsProxyWidget) OnGrabKeyboardEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 6423297148758141226, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "grabKeyboardEvent"
}

func (this *QGraphicsProxyWidget) OnUngrabKeyboardEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 2912496686457383679, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "ungrabKeyboardEvent"
}

func (this *QGraphicsProxyWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsProxyWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsProxyWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsProxyWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsProxyWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsProxyWidget) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsProxyWidget) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsProxyWidget) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsProxyWidget) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsProxyWidget) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsProxyWidget) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsProxyWidget) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsProxyWidget) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsProxyWidget) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsProxyWidget) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsProxyWidget) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsProxyWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsProxyWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
