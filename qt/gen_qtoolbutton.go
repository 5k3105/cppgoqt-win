package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QToolButton_chain uintptr
	addr_QToolButton_new uintptr
	addr_QToolButton_new2 uintptr
	addr_QToolButton_tr uintptr
	addr_QToolButton_toolButtonStyle uintptr
	addr_QToolButton_arrowType uintptr
	addr_QToolButton_setArrowType uintptr
	addr_QToolButton_setMenu uintptr
	addr_QToolButton_menu uintptr
	addr_QToolButton_setPopupMode uintptr
	addr_QToolButton_popupMode uintptr
	addr_QToolButton_defaultAction uintptr
	addr_QToolButton_setAutoRaise uintptr
	addr_QToolButton_autoRaise uintptr
	addr_QToolButton_showMenu uintptr
	addr_QToolButton_setToolButtonStyle uintptr
	addr_QToolButton_setDefaultAction uintptr
	addr_QToolButton_triggered uintptr
	addr_QToolButton_setTabOrder uintptr
	addr_QToolButton_mouseGrabber uintptr
	addr_QToolButton_keyboardGrabber uintptr
	addr_QToolButton_find uintptr
	addr_QToolButton_createWindowContainer uintptr
	addr_QToolButton_updateMicroFocus uintptr
	addr_QToolButton_create uintptr
	addr_QToolButton_destroy uintptr
	addr_QToolButton_focusNextChild uintptr
	addr_QToolButton_focusPreviousChild uintptr
	addr_QToolButton_connect uintptr
	addr_QToolButton_connect2 uintptr
	addr_QToolButton_sender uintptr
	addr_QToolButton_senderSignalIndex uintptr
	addr_QToolButton_receivers uintptr
	addr_QToolButton_isSignalConnected uintptr
	pg_QToolButton_devicePixelRatioFScale func() float64
	pg_QToolButton_encodeMetricF func(int32, float64) int32
	pg_QToolButton_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QToolButton_tr2 uintptr
	addr_QToolButton_tr3 uintptr
	addr_QToolButton_createWindowContainer2 uintptr
	addr_QToolButton_createWindowContainer3 uintptr
	addr_QToolButton_updateMicroFocus2 uintptr
	addr_QToolButton_create2 uintptr
	addr_QToolButton_create3 uintptr
	addr_QToolButton_create4 uintptr
	addr_QToolButton_destroy2 uintptr
	addr_QToolButton_destroy3 uintptr
	addr_QToolButton_connect3 uintptr
	addr_QToolButton_connect4 uintptr
	addr_QToolButton_metaObject_default uintptr
	addr_QToolButton_sizeHint_default uintptr
	addr_QToolButton_minimumSizeHint_default uintptr
	addr_QToolButton_event_default uintptr
	addr_QToolButton_mousePressEvent_default uintptr
	addr_QToolButton_mouseReleaseEvent_default uintptr
	addr_QToolButton_paintEvent_default uintptr
	addr_QToolButton_actionEvent_default uintptr
	addr_QToolButton_enterEvent_default uintptr
	addr_QToolButton_leaveEvent_default uintptr
	addr_QToolButton_timerEvent_default uintptr
	addr_QToolButton_changeEvent_default uintptr
	addr_QToolButton_hitButton_default uintptr
	addr_QToolButton_checkStateSet_default uintptr
	addr_QToolButton_nextCheckState_default uintptr
	addr_QToolButton_initStyleOption_default uintptr
	addr_QToolButton_keyPressEvent_default uintptr
	addr_QToolButton_keyReleaseEvent_default uintptr
	addr_QToolButton_mouseMoveEvent_default uintptr
	addr_QToolButton_focusInEvent_default uintptr
	addr_QToolButton_focusOutEvent_default uintptr
	addr_QToolButton_devType_default uintptr
	addr_QToolButton_setVisible_default uintptr
	addr_QToolButton_heightForWidth_default uintptr
	addr_QToolButton_hasHeightForWidth_default uintptr
	addr_QToolButton_paintEngine_default uintptr
	addr_QToolButton_mouseDoubleClickEvent_default uintptr
	addr_QToolButton_wheelEvent_default uintptr
	addr_QToolButton_moveEvent_default uintptr
	addr_QToolButton_resizeEvent_default uintptr
	addr_QToolButton_closeEvent_default uintptr
	addr_QToolButton_contextMenuEvent_default uintptr
	addr_QToolButton_tabletEvent_default uintptr
	addr_QToolButton_dragEnterEvent_default uintptr
	addr_QToolButton_dragMoveEvent_default uintptr
	addr_QToolButton_dragLeaveEvent_default uintptr
	addr_QToolButton_dropEvent_default uintptr
	addr_QToolButton_showEvent_default uintptr
	addr_QToolButton_hideEvent_default uintptr
	addr_QToolButton_metric_default uintptr
	addr_QToolButton_initPainter_default uintptr
	addr_QToolButton_redirected_default uintptr
	addr_QToolButton_sharedPainter_default uintptr
	addr_QToolButton_inputMethodEvent_default uintptr
	addr_QToolButton_inputMethodQuery_default uintptr
	addr_QToolButton_focusNextPrevChild_default uintptr
	addr_QToolButton_eventFilter_default uintptr
	addr_QToolButton_childEvent_default uintptr
	addr_QToolButton_customEvent_default uintptr
	addr_QToolButton_connectNotify_default uintptr
	addr_QToolButton_disconnectNotify_default uintptr
	pg_QToolButton_connect_triggered func(uintptr, int64)
	pg_QToolButton_connect_pressed func(uintptr, int64)
	pg_QToolButton_connect_released func(uintptr, int64)
	pg_QToolButton_connect_clicked func(uintptr, int64)
	pg_QToolButton_connect_clicked2 func(uintptr, int64)
	pg_QToolButton_connect_toggled func(uintptr, int64)
	pg_QToolButton_connect_windowTitleChanged func(uintptr, int64)
	pg_QToolButton_connect_windowIconChanged func(uintptr, int64)
	pg_QToolButton_connect_windowIconTextChanged func(uintptr, int64)
	pg_QToolButton_connect_customContextMenuRequested func(uintptr, int64)
	pg_QToolButton_connect_destroyed func(uintptr, int64)
	pg_QToolButton_connect_destroyed2 func(uintptr, int64)
	pg_QToolButton_connect_objectNameChanged func(uintptr, int64)
	pg_QToolButton__disconnect func(int64)
	pg_QToolButton__set_override func(uintptr, int32)
	pg_QToolButton_delete func(uintptr)
)

func init() {
	symaddr(&addr_QToolButton_chain, "QToolButton_chain")
	symaddr(&addr_QToolButton_new, "QToolButton_new")
	symaddr(&addr_QToolButton_new2, "QToolButton_new2")
	symaddr(&addr_QToolButton_tr, "QToolButton_tr")
	symaddr(&addr_QToolButton_toolButtonStyle, "QToolButton_toolButtonStyle")
	symaddr(&addr_QToolButton_arrowType, "QToolButton_arrowType")
	symaddr(&addr_QToolButton_setArrowType, "QToolButton_setArrowType")
	symaddr(&addr_QToolButton_setMenu, "QToolButton_setMenu")
	symaddr(&addr_QToolButton_menu, "QToolButton_menu")
	symaddr(&addr_QToolButton_setPopupMode, "QToolButton_setPopupMode")
	symaddr(&addr_QToolButton_popupMode, "QToolButton_popupMode")
	symaddr(&addr_QToolButton_defaultAction, "QToolButton_defaultAction")
	symaddr(&addr_QToolButton_setAutoRaise, "QToolButton_setAutoRaise")
	symaddr(&addr_QToolButton_autoRaise, "QToolButton_autoRaise")
	symaddr(&addr_QToolButton_showMenu, "QToolButton_showMenu")
	symaddr(&addr_QToolButton_setToolButtonStyle, "QToolButton_setToolButtonStyle")
	symaddr(&addr_QToolButton_setDefaultAction, "QToolButton_setDefaultAction")
	symaddr(&addr_QToolButton_triggered, "QToolButton_triggered")
	symaddr(&addr_QToolButton_setTabOrder, "QToolButton_setTabOrder")
	symaddr(&addr_QToolButton_mouseGrabber, "QToolButton_mouseGrabber")
	symaddr(&addr_QToolButton_keyboardGrabber, "QToolButton_keyboardGrabber")
	symaddr(&addr_QToolButton_find, "QToolButton_find")
	symaddr(&addr_QToolButton_createWindowContainer, "QToolButton_createWindowContainer")
	symaddr(&addr_QToolButton_updateMicroFocus, "QToolButton_updateMicroFocus")
	symaddr(&addr_QToolButton_create, "QToolButton_create")
	symaddr(&addr_QToolButton_destroy, "QToolButton_destroy")
	symaddr(&addr_QToolButton_focusNextChild, "QToolButton_focusNextChild")
	symaddr(&addr_QToolButton_focusPreviousChild, "QToolButton_focusPreviousChild")
	symaddr(&addr_QToolButton_connect, "QToolButton_connect")
	symaddr(&addr_QToolButton_connect2, "QToolButton_connect2")
	symaddr(&addr_QToolButton_sender, "QToolButton_sender")
	symaddr(&addr_QToolButton_senderSignalIndex, "QToolButton_senderSignalIndex")
	symaddr(&addr_QToolButton_receivers, "QToolButton_receivers")
	symaddr(&addr_QToolButton_isSignalConnected, "QToolButton_isSignalConnected")
	reg(&pg_QToolButton_devicePixelRatioFScale, "QToolButton_devicePixelRatioFScale")
	reg(&pg_QToolButton_encodeMetricF, "QToolButton_encodeMetricF")
	reg(&pg_QToolButton_getDecodedMetricF, "QToolButton_getDecodedMetricF")
	symaddr(&addr_QToolButton_tr2, "QToolButton_tr2")
	symaddr(&addr_QToolButton_tr3, "QToolButton_tr3")
	symaddr(&addr_QToolButton_createWindowContainer2, "QToolButton_createWindowContainer2")
	symaddr(&addr_QToolButton_createWindowContainer3, "QToolButton_createWindowContainer3")
	symaddr(&addr_QToolButton_updateMicroFocus2, "QToolButton_updateMicroFocus2")
	symaddr(&addr_QToolButton_create2, "QToolButton_create2")
	symaddr(&addr_QToolButton_create3, "QToolButton_create3")
	symaddr(&addr_QToolButton_create4, "QToolButton_create4")
	symaddr(&addr_QToolButton_destroy2, "QToolButton_destroy2")
	symaddr(&addr_QToolButton_destroy3, "QToolButton_destroy3")
	symaddr(&addr_QToolButton_connect3, "QToolButton_connect3")
	symaddr(&addr_QToolButton_connect4, "QToolButton_connect4")
	symaddr(&addr_QToolButton_metaObject_default, "QToolButton_metaObject_default")
	symaddr(&addr_QToolButton_sizeHint_default, "QToolButton_sizeHint_default")
	symaddr(&addr_QToolButton_minimumSizeHint_default, "QToolButton_minimumSizeHint_default")
	symaddr(&addr_QToolButton_event_default, "QToolButton_event_default")
	symaddr(&addr_QToolButton_mousePressEvent_default, "QToolButton_mousePressEvent_default")
	symaddr(&addr_QToolButton_mouseReleaseEvent_default, "QToolButton_mouseReleaseEvent_default")
	symaddr(&addr_QToolButton_paintEvent_default, "QToolButton_paintEvent_default")
	symaddr(&addr_QToolButton_actionEvent_default, "QToolButton_actionEvent_default")
	symaddr(&addr_QToolButton_enterEvent_default, "QToolButton_enterEvent_default")
	symaddr(&addr_QToolButton_leaveEvent_default, "QToolButton_leaveEvent_default")
	symaddr(&addr_QToolButton_timerEvent_default, "QToolButton_timerEvent_default")
	symaddr(&addr_QToolButton_changeEvent_default, "QToolButton_changeEvent_default")
	symaddr(&addr_QToolButton_hitButton_default, "QToolButton_hitButton_default")
	symaddr(&addr_QToolButton_checkStateSet_default, "QToolButton_checkStateSet_default")
	symaddr(&addr_QToolButton_nextCheckState_default, "QToolButton_nextCheckState_default")
	symaddr(&addr_QToolButton_initStyleOption_default, "QToolButton_initStyleOption_default")
	symaddr(&addr_QToolButton_keyPressEvent_default, "QToolButton_keyPressEvent_default")
	symaddr(&addr_QToolButton_keyReleaseEvent_default, "QToolButton_keyReleaseEvent_default")
	symaddr(&addr_QToolButton_mouseMoveEvent_default, "QToolButton_mouseMoveEvent_default")
	symaddr(&addr_QToolButton_focusInEvent_default, "QToolButton_focusInEvent_default")
	symaddr(&addr_QToolButton_focusOutEvent_default, "QToolButton_focusOutEvent_default")
	symaddr(&addr_QToolButton_devType_default, "QToolButton_devType_default")
	symaddr(&addr_QToolButton_setVisible_default, "QToolButton_setVisible_default")
	symaddr(&addr_QToolButton_heightForWidth_default, "QToolButton_heightForWidth_default")
	symaddr(&addr_QToolButton_hasHeightForWidth_default, "QToolButton_hasHeightForWidth_default")
	symaddr(&addr_QToolButton_paintEngine_default, "QToolButton_paintEngine_default")
	symaddr(&addr_QToolButton_mouseDoubleClickEvent_default, "QToolButton_mouseDoubleClickEvent_default")
	symaddr(&addr_QToolButton_wheelEvent_default, "QToolButton_wheelEvent_default")
	symaddr(&addr_QToolButton_moveEvent_default, "QToolButton_moveEvent_default")
	symaddr(&addr_QToolButton_resizeEvent_default, "QToolButton_resizeEvent_default")
	symaddr(&addr_QToolButton_closeEvent_default, "QToolButton_closeEvent_default")
	symaddr(&addr_QToolButton_contextMenuEvent_default, "QToolButton_contextMenuEvent_default")
	symaddr(&addr_QToolButton_tabletEvent_default, "QToolButton_tabletEvent_default")
	symaddr(&addr_QToolButton_dragEnterEvent_default, "QToolButton_dragEnterEvent_default")
	symaddr(&addr_QToolButton_dragMoveEvent_default, "QToolButton_dragMoveEvent_default")
	symaddr(&addr_QToolButton_dragLeaveEvent_default, "QToolButton_dragLeaveEvent_default")
	symaddr(&addr_QToolButton_dropEvent_default, "QToolButton_dropEvent_default")
	symaddr(&addr_QToolButton_showEvent_default, "QToolButton_showEvent_default")
	symaddr(&addr_QToolButton_hideEvent_default, "QToolButton_hideEvent_default")
	symaddr(&addr_QToolButton_metric_default, "QToolButton_metric_default")
	symaddr(&addr_QToolButton_initPainter_default, "QToolButton_initPainter_default")
	symaddr(&addr_QToolButton_redirected_default, "QToolButton_redirected_default")
	symaddr(&addr_QToolButton_sharedPainter_default, "QToolButton_sharedPainter_default")
	symaddr(&addr_QToolButton_inputMethodEvent_default, "QToolButton_inputMethodEvent_default")
	symaddr(&addr_QToolButton_inputMethodQuery_default, "QToolButton_inputMethodQuery_default")
	symaddr(&addr_QToolButton_focusNextPrevChild_default, "QToolButton_focusNextPrevChild_default")
	symaddr(&addr_QToolButton_eventFilter_default, "QToolButton_eventFilter_default")
	symaddr(&addr_QToolButton_childEvent_default, "QToolButton_childEvent_default")
	symaddr(&addr_QToolButton_customEvent_default, "QToolButton_customEvent_default")
	symaddr(&addr_QToolButton_connectNotify_default, "QToolButton_connectNotify_default")
	symaddr(&addr_QToolButton_disconnectNotify_default, "QToolButton_disconnectNotify_default")
	reg(&pg_QToolButton_connect_triggered, "QToolButton_connect_triggered")
	reg(&pg_QToolButton_connect_pressed, "QToolButton_connect_pressed")
	reg(&pg_QToolButton_connect_released, "QToolButton_connect_released")
	reg(&pg_QToolButton_connect_clicked, "QToolButton_connect_clicked")
	reg(&pg_QToolButton_connect_clicked2, "QToolButton_connect_clicked2")
	reg(&pg_QToolButton_connect_toggled, "QToolButton_connect_toggled")
	reg(&pg_QToolButton_connect_windowTitleChanged, "QToolButton_connect_windowTitleChanged")
	reg(&pg_QToolButton_connect_windowIconChanged, "QToolButton_connect_windowIconChanged")
	reg(&pg_QToolButton_connect_windowIconTextChanged, "QToolButton_connect_windowIconTextChanged")
	reg(&pg_QToolButton_connect_customContextMenuRequested, "QToolButton_connect_customContextMenuRequested")
	reg(&pg_QToolButton_connect_destroyed, "QToolButton_connect_destroyed")
	reg(&pg_QToolButton_connect_destroyed2, "QToolButton_connect_destroyed2")
	reg(&pg_QToolButton_connect_objectNameChanged, "QToolButton_connect_objectNameChanged")
	reg(&pg_QToolButton__disconnect, "QToolButton__disconnect")
	reg(&pg_QToolButton__set_override, "QToolButton__set_override")
	reg(&pg_QToolButton_delete, "QToolButton_delete")
}

type QToolButton struct {
	h uintptr
	*QAbstractButton
}

func bindQToolButton(h uintptr) *QToolButton {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QToolButton_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QToolButton{h: h, QAbstractButton: &QAbstractButton{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQToolButton(h uintptr) *QToolButton { return bindQToolButton(h) }

func (this *QToolButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QToolButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQToolButton(h unsafe.Pointer) *QToolButton { return bindQToolButton(uintptr(h)) }
func UnsafeNewQToolButtonOwned(h unsafe.Pointer) *QToolButton { return newownedQToolButton(uintptr(h)) }

func (this *QToolButton) UnsafeRelease() {
	this.h = 0
	if this.QAbstractButton != nil {
		this.QAbstractButton.UnsafeRelease()
	}
}

func NewQToolButton(parent *QWidget) *QToolButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolButton_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQToolButton(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolButton_virtIdx, QToolButton_enable)
	return __o
}

func NewQToolButton2() *QToolButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolButton_new2, uintptr(__goid))
	__o := newownedQToolButton(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolButton_virtIdx, QToolButton_enable)
	return __o
}

func QToolButton_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolButton_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QToolButton) ToolButtonStyle() ToolButtonStyle {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_toolButtonStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolButtonStyle(__r)
}

func (this *QToolButton) ArrowType() ArrowType {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_arrowType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ArrowType(__r)
}

func (this *QToolButton) SetArrowType(type_ ArrowType) {
	purego.SyscallN(addr_QToolButton_setArrowType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QToolButton) SetMenu(menu *QMenu) {
	purego.SyscallN(addr_QToolButton_setMenu, this.h, uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menu)
}

func (this *QToolButton) Menu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QToolButton_menu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QToolButton) SetPopupMode(mode QToolButton__ToolButtonPopupMode) {
	purego.SyscallN(addr_QToolButton_setPopupMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QToolButton) PopupMode() QToolButton__ToolButtonPopupMode {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_popupMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QToolButton__ToolButtonPopupMode(__r)
}

func (this *QToolButton) DefaultAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolButton_defaultAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QToolButton) SetAutoRaise(enable bool) {
	purego.SyscallN(addr_QToolButton_setAutoRaise, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QToolButton) AutoRaise() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_autoRaise, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolButton) ShowMenu() {
	purego.SyscallN(addr_QToolButton_showMenu, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) SetToolButtonStyle(style ToolButtonStyle) {
	purego.SyscallN(addr_QToolButton_setToolButtonStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QToolButton) SetDefaultAction(arg1 *QAction) {
	purego.SyscallN(addr_QToolButton_setDefaultAction, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) Triggered(arg1 *QAction) {
	purego.SyscallN(addr_QToolButton_triggered, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QToolButton_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QToolButton_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QToolButton_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_mouseGrabber)
	return bindQWidget(__r)
}

func QToolButton_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_keyboardGrabber)
	return bindQWidget(__r)
}

func QToolButton_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QToolButton_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QToolButton) UpdateMicroFocus() {
	purego.SyscallN(addr_QToolButton_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) Create() {
	purego.SyscallN(addr_QToolButton_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) Destroy() {
	purego.SyscallN(addr_QToolButton_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolButton) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QToolButton_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolButton_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolButton_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolButton_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolButton) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QToolButton_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QToolButton) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolButton) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QToolButton_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QToolButton) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QToolButton_DevicePixelRatioFScale() float64 {
	__r := pg_QToolButton_devicePixelRatioFScale()
	return __r
}

func QToolButton_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QToolButton_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QToolButton) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QToolButton_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QToolButton_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolButton_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolButton_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolButton_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolButton_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QToolButton_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolButton_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QToolButton) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QToolButton_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QToolButton) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QToolButton_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QToolButton) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QToolButton_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolButton) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QToolButton_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolButton) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QToolButton_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolButton) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QToolButton_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QToolButton_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolButton_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolButton_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolButton_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolButton) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QToolButton_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QToolButton) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolButton_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolButton) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolButton_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolButton) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QToolButton) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QToolButton_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QToolButton_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QToolButton_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) ActionEventDefault(arg1 *QActionEvent) {
	purego.SyscallN(addr_QToolButton_actionEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) EnterEventDefault(arg1 *QEnterEvent) {
	purego.SyscallN(addr_QToolButton_enterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) LeaveEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QToolButton_leaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QToolButton_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QToolButton_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) HitButtonDefault(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_hitButton_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QToolButton) CheckStateSetDefault() {
	purego.SyscallN(addr_QToolButton_checkStateSet_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) NextCheckStateDefault() {
	purego.SyscallN(addr_QToolButton_nextCheckState_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolButton) InitStyleOptionDefault(option *QStyleOptionToolButton) {
	purego.SyscallN(addr_QToolButton_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QToolButton) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QToolButton_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolButton) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QToolButton_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolButton) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QToolButton_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolButton) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QToolButton_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolButton) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QToolButton_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolButton) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolButton) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QToolButton_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QToolButton) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolButton) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolButton) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QToolButton_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QToolButton) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolButton_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QToolButton_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QToolButton_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QToolButton_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QToolButton_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QToolButton_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QToolButton_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QToolButton_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QToolButton_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QToolButton_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QToolButton_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QToolButton_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QToolButton_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolButton) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QToolButton_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QToolButton) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QToolButton_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QToolButton) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QToolButton_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QToolButton) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QToolButton_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolButton) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QToolButton_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QToolButton) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolButton) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolButton_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QToolButton) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QToolButton_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolButton_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolButton) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolButton_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolButton) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolButton_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolButton) OnTriggered(fn func(*QAction)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(arg1))) })
	pg_QToolButton_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolButton_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolButton_connect_released(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolButton_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QToolButton_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QToolButton_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QToolButton_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QToolButton_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QToolButton_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QToolButton_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolButton_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QToolButton_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QToolButton_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolButton) Disconnect(handle int64) {
	pg_QToolButton__disconnect(handle)
	unregisterCallback(handle)
}

func QToolButton_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case -332090898078543319:
		return 4
	case -4415920049297795213:
		return 5
	case 5290442163167458465:
		return 6
	case 1697444140088555581:
		return 7
	case 6316270346668237895:
		return 8
	case -551070009623711640:
		return 9
	case -5409927581812374498:
		return 10
	case 4959376063826165943:
		return 11
	case 5736040353629178930:
		return 12
	case 5747806663224476200:
		return 13
	case 1474647914901213797:
		return 14
	case -488693769771308531:
		return 15
	case 1023791977679793185:
		return 16
	case 772598475950050299:
		return 17
	case -4255494004099806581:
		return 18
	case 975043767033552552:
		return 19
	case -385596010576048993:
		return 20
	case -8659679893868524044:
		return 21
	case 2110182719375245575:
		return 22
	case 4311534060495737339:
		return 23
	case -8266762397995423915:
		return 24
	case 6398007404756346381:
		return 25
	case -4029073905697023887:
		return 26
	case 1919988340934511316:
		return 27
	case -9161949150489525396:
		return 28
	case -4302917285655014661:
		return 29
	case -1989937050019464783:
		return 30
	case -6698043222227006015:
		return 31
	case -212068866860701141:
		return 32
	case 206122715919917589:
		return 33
	case -1781545137352745978:
		return 34
	case -3139904158373725006:
		return 35
	case 9107058682754591716:
		return 36
	case -6825369174751916:
		return 37
	case -199718707876719129:
		return 38
	case -2039166586246443706:
		return 39
	case 427629649267311777:
		return 40
	case 6477500296980050210:
		return 41
	case 4127240633310655048:
		return 42
	case 5792771328741774991:
		return 43
	case 5438819661942843222:
		return 44
	case 7345853097340709638:
		return 45
	case 5898328517505784191:
		return 46
	case 5217623514630858297:
		return 47
	case -1185315367526857833:
		return 48
	case 3769868229123868474:
		return 49
	case 7379029982170503008:
		return 50
	case -568369772982215962:
		return 51
	}
	return -1
}

func QToolButton_enable(h uintptr, idx int32) { pg_QToolButton__set_override(h, idx) }

func (this *QToolButton) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QToolButton) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QToolButton) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QToolButton) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QToolButton) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QToolButton) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QToolButton) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QToolButton) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(arg1 uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(arg1))) }) // slot "actionEvent"
}

func (this *QToolButton) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(arg1 uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(arg1))) }) // slot "enterEvent"
}

func (this *QToolButton) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "leaveEvent"
}

func (this *QToolButton) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QToolButton) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QToolButton) OnHitButton(fn func(*QPoint) bool) {
	registerVirt(this.UnsafeGoID(), 5736040353629178930, func(pos uintptr) int32 { return b2i(fn(UnsafeNewQPoint(unsafe.Pointer(pos)))) }) // slot "hitButton"
}

func (this *QToolButton) OnCheckStateSet(fn func()) {
	registerVirt(this.UnsafeGoID(), 5747806663224476200, func() { fn() }) // slot "checkStateSet"
}

func (this *QToolButton) OnNextCheckState(fn func()) {
	registerVirt(this.UnsafeGoID(), 1474647914901213797, func() { fn() }) // slot "nextCheckState"
}

func (this *QToolButton) OnInitStyleOption(fn func(*QStyleOptionToolButton)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionToolButton(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QToolButton) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QToolButton) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QToolButton) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QToolButton) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QToolButton) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QToolButton) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QToolButton) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QToolButton) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QToolButton) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QToolButton) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QToolButton) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QToolButton) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QToolButton) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QToolButton) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QToolButton) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QToolButton) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QToolButton) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QToolButton) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QToolButton) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QToolButton) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QToolButton) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QToolButton) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QToolButton) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QToolButton) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QToolButton) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QToolButton) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QToolButton) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QToolButton) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QToolButton) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QToolButton) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QToolButton) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QToolButton) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QToolButton) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QToolButton) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QToolButton) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QToolButton) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QToolButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QToolButton_delete(this.h)
	}
	this.UnsafeRelease()
}
