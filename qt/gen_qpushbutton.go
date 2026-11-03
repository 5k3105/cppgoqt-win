package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPushButton_chain uintptr
	addr_QPushButton_new uintptr
	addr_QPushButton_new2 uintptr
	addr_QPushButton_new3 uintptr
	addr_QPushButton_new4 uintptr
	addr_QPushButton_new5 uintptr
	addr_QPushButton_new6 uintptr
	addr_QPushButton_tr uintptr
	addr_QPushButton_autoDefault uintptr
	addr_QPushButton_setAutoDefault uintptr
	addr_QPushButton_isDefault uintptr
	addr_QPushButton_setDefault uintptr
	addr_QPushButton_setMenu uintptr
	addr_QPushButton_menu uintptr
	addr_QPushButton_setFlat uintptr
	addr_QPushButton_isFlat uintptr
	addr_QPushButton_showMenu uintptr
	addr_QPushButton_setTabOrder uintptr
	addr_QPushButton_mouseGrabber uintptr
	addr_QPushButton_keyboardGrabber uintptr
	addr_QPushButton_find uintptr
	addr_QPushButton_createWindowContainer uintptr
	addr_QPushButton_updateMicroFocus uintptr
	addr_QPushButton_create uintptr
	addr_QPushButton_destroy uintptr
	addr_QPushButton_focusNextChild uintptr
	addr_QPushButton_focusPreviousChild uintptr
	addr_QPushButton_connect uintptr
	addr_QPushButton_connect2 uintptr
	addr_QPushButton_sender uintptr
	addr_QPushButton_senderSignalIndex uintptr
	addr_QPushButton_receivers uintptr
	addr_QPushButton_isSignalConnected uintptr
	pg_QPushButton_devicePixelRatioFScale func() float64
	pg_QPushButton_encodeMetricF func(int32, float64) int32
	pg_QPushButton_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QPushButton_tr2 uintptr
	addr_QPushButton_tr3 uintptr
	addr_QPushButton_createWindowContainer2 uintptr
	addr_QPushButton_createWindowContainer3 uintptr
	addr_QPushButton_updateMicroFocus2 uintptr
	addr_QPushButton_create2 uintptr
	addr_QPushButton_create3 uintptr
	addr_QPushButton_create4 uintptr
	addr_QPushButton_destroy2 uintptr
	addr_QPushButton_destroy3 uintptr
	addr_QPushButton_connect3 uintptr
	addr_QPushButton_connect4 uintptr
	addr_QPushButton_metaObject_default uintptr
	addr_QPushButton_sizeHint_default uintptr
	addr_QPushButton_minimumSizeHint_default uintptr
	addr_QPushButton_event_default uintptr
	addr_QPushButton_paintEvent_default uintptr
	addr_QPushButton_keyPressEvent_default uintptr
	addr_QPushButton_focusInEvent_default uintptr
	addr_QPushButton_focusOutEvent_default uintptr
	addr_QPushButton_mouseMoveEvent_default uintptr
	addr_QPushButton_initStyleOption_default uintptr
	addr_QPushButton_hitButton_default uintptr
	addr_QPushButton_checkStateSet_default uintptr
	addr_QPushButton_nextCheckState_default uintptr
	addr_QPushButton_keyReleaseEvent_default uintptr
	addr_QPushButton_mousePressEvent_default uintptr
	addr_QPushButton_mouseReleaseEvent_default uintptr
	addr_QPushButton_changeEvent_default uintptr
	addr_QPushButton_timerEvent_default uintptr
	addr_QPushButton_devType_default uintptr
	addr_QPushButton_setVisible_default uintptr
	addr_QPushButton_heightForWidth_default uintptr
	addr_QPushButton_hasHeightForWidth_default uintptr
	addr_QPushButton_paintEngine_default uintptr
	addr_QPushButton_mouseDoubleClickEvent_default uintptr
	addr_QPushButton_wheelEvent_default uintptr
	addr_QPushButton_enterEvent_default uintptr
	addr_QPushButton_leaveEvent_default uintptr
	addr_QPushButton_moveEvent_default uintptr
	addr_QPushButton_resizeEvent_default uintptr
	addr_QPushButton_closeEvent_default uintptr
	addr_QPushButton_contextMenuEvent_default uintptr
	addr_QPushButton_tabletEvent_default uintptr
	addr_QPushButton_actionEvent_default uintptr
	addr_QPushButton_dragEnterEvent_default uintptr
	addr_QPushButton_dragMoveEvent_default uintptr
	addr_QPushButton_dragLeaveEvent_default uintptr
	addr_QPushButton_dropEvent_default uintptr
	addr_QPushButton_showEvent_default uintptr
	addr_QPushButton_hideEvent_default uintptr
	addr_QPushButton_metric_default uintptr
	addr_QPushButton_initPainter_default uintptr
	addr_QPushButton_redirected_default uintptr
	addr_QPushButton_sharedPainter_default uintptr
	addr_QPushButton_inputMethodEvent_default uintptr
	addr_QPushButton_inputMethodQuery_default uintptr
	addr_QPushButton_focusNextPrevChild_default uintptr
	addr_QPushButton_eventFilter_default uintptr
	addr_QPushButton_childEvent_default uintptr
	addr_QPushButton_customEvent_default uintptr
	addr_QPushButton_connectNotify_default uintptr
	addr_QPushButton_disconnectNotify_default uintptr
	pg_QPushButton_connect_pressed func(uintptr, int64)
	pg_QPushButton_connect_released func(uintptr, int64)
	pg_QPushButton_connect_clicked func(uintptr, int64)
	pg_QPushButton_connect_clicked2 func(uintptr, int64)
	pg_QPushButton_connect_toggled func(uintptr, int64)
	pg_QPushButton_connect_windowTitleChanged func(uintptr, int64)
	pg_QPushButton_connect_windowIconChanged func(uintptr, int64)
	pg_QPushButton_connect_windowIconTextChanged func(uintptr, int64)
	pg_QPushButton_connect_customContextMenuRequested func(uintptr, int64)
	pg_QPushButton_connect_destroyed func(uintptr, int64)
	pg_QPushButton_connect_destroyed2 func(uintptr, int64)
	pg_QPushButton_connect_objectNameChanged func(uintptr, int64)
	pg_QPushButton__disconnect func(int64)
	pg_QPushButton__set_override func(uintptr, int32)
	pg_QPushButton_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPushButton_chain, "QPushButton_chain")
	symaddr(&addr_QPushButton_new, "QPushButton_new")
	symaddr(&addr_QPushButton_new2, "QPushButton_new2")
	symaddr(&addr_QPushButton_new3, "QPushButton_new3")
	symaddr(&addr_QPushButton_new4, "QPushButton_new4")
	symaddr(&addr_QPushButton_new5, "QPushButton_new5")
	symaddr(&addr_QPushButton_new6, "QPushButton_new6")
	symaddr(&addr_QPushButton_tr, "QPushButton_tr")
	symaddr(&addr_QPushButton_autoDefault, "QPushButton_autoDefault")
	symaddr(&addr_QPushButton_setAutoDefault, "QPushButton_setAutoDefault")
	symaddr(&addr_QPushButton_isDefault, "QPushButton_isDefault")
	symaddr(&addr_QPushButton_setDefault, "QPushButton_setDefault")
	symaddr(&addr_QPushButton_setMenu, "QPushButton_setMenu")
	symaddr(&addr_QPushButton_menu, "QPushButton_menu")
	symaddr(&addr_QPushButton_setFlat, "QPushButton_setFlat")
	symaddr(&addr_QPushButton_isFlat, "QPushButton_isFlat")
	symaddr(&addr_QPushButton_showMenu, "QPushButton_showMenu")
	symaddr(&addr_QPushButton_setTabOrder, "QPushButton_setTabOrder")
	symaddr(&addr_QPushButton_mouseGrabber, "QPushButton_mouseGrabber")
	symaddr(&addr_QPushButton_keyboardGrabber, "QPushButton_keyboardGrabber")
	symaddr(&addr_QPushButton_find, "QPushButton_find")
	symaddr(&addr_QPushButton_createWindowContainer, "QPushButton_createWindowContainer")
	symaddr(&addr_QPushButton_updateMicroFocus, "QPushButton_updateMicroFocus")
	symaddr(&addr_QPushButton_create, "QPushButton_create")
	symaddr(&addr_QPushButton_destroy, "QPushButton_destroy")
	symaddr(&addr_QPushButton_focusNextChild, "QPushButton_focusNextChild")
	symaddr(&addr_QPushButton_focusPreviousChild, "QPushButton_focusPreviousChild")
	symaddr(&addr_QPushButton_connect, "QPushButton_connect")
	symaddr(&addr_QPushButton_connect2, "QPushButton_connect2")
	symaddr(&addr_QPushButton_sender, "QPushButton_sender")
	symaddr(&addr_QPushButton_senderSignalIndex, "QPushButton_senderSignalIndex")
	symaddr(&addr_QPushButton_receivers, "QPushButton_receivers")
	symaddr(&addr_QPushButton_isSignalConnected, "QPushButton_isSignalConnected")
	reg(&pg_QPushButton_devicePixelRatioFScale, "QPushButton_devicePixelRatioFScale")
	reg(&pg_QPushButton_encodeMetricF, "QPushButton_encodeMetricF")
	reg(&pg_QPushButton_getDecodedMetricF, "QPushButton_getDecodedMetricF")
	symaddr(&addr_QPushButton_tr2, "QPushButton_tr2")
	symaddr(&addr_QPushButton_tr3, "QPushButton_tr3")
	symaddr(&addr_QPushButton_createWindowContainer2, "QPushButton_createWindowContainer2")
	symaddr(&addr_QPushButton_createWindowContainer3, "QPushButton_createWindowContainer3")
	symaddr(&addr_QPushButton_updateMicroFocus2, "QPushButton_updateMicroFocus2")
	symaddr(&addr_QPushButton_create2, "QPushButton_create2")
	symaddr(&addr_QPushButton_create3, "QPushButton_create3")
	symaddr(&addr_QPushButton_create4, "QPushButton_create4")
	symaddr(&addr_QPushButton_destroy2, "QPushButton_destroy2")
	symaddr(&addr_QPushButton_destroy3, "QPushButton_destroy3")
	symaddr(&addr_QPushButton_connect3, "QPushButton_connect3")
	symaddr(&addr_QPushButton_connect4, "QPushButton_connect4")
	symaddr(&addr_QPushButton_metaObject_default, "QPushButton_metaObject_default")
	symaddr(&addr_QPushButton_sizeHint_default, "QPushButton_sizeHint_default")
	symaddr(&addr_QPushButton_minimumSizeHint_default, "QPushButton_minimumSizeHint_default")
	symaddr(&addr_QPushButton_event_default, "QPushButton_event_default")
	symaddr(&addr_QPushButton_paintEvent_default, "QPushButton_paintEvent_default")
	symaddr(&addr_QPushButton_keyPressEvent_default, "QPushButton_keyPressEvent_default")
	symaddr(&addr_QPushButton_focusInEvent_default, "QPushButton_focusInEvent_default")
	symaddr(&addr_QPushButton_focusOutEvent_default, "QPushButton_focusOutEvent_default")
	symaddr(&addr_QPushButton_mouseMoveEvent_default, "QPushButton_mouseMoveEvent_default")
	symaddr(&addr_QPushButton_initStyleOption_default, "QPushButton_initStyleOption_default")
	symaddr(&addr_QPushButton_hitButton_default, "QPushButton_hitButton_default")
	symaddr(&addr_QPushButton_checkStateSet_default, "QPushButton_checkStateSet_default")
	symaddr(&addr_QPushButton_nextCheckState_default, "QPushButton_nextCheckState_default")
	symaddr(&addr_QPushButton_keyReleaseEvent_default, "QPushButton_keyReleaseEvent_default")
	symaddr(&addr_QPushButton_mousePressEvent_default, "QPushButton_mousePressEvent_default")
	symaddr(&addr_QPushButton_mouseReleaseEvent_default, "QPushButton_mouseReleaseEvent_default")
	symaddr(&addr_QPushButton_changeEvent_default, "QPushButton_changeEvent_default")
	symaddr(&addr_QPushButton_timerEvent_default, "QPushButton_timerEvent_default")
	symaddr(&addr_QPushButton_devType_default, "QPushButton_devType_default")
	symaddr(&addr_QPushButton_setVisible_default, "QPushButton_setVisible_default")
	symaddr(&addr_QPushButton_heightForWidth_default, "QPushButton_heightForWidth_default")
	symaddr(&addr_QPushButton_hasHeightForWidth_default, "QPushButton_hasHeightForWidth_default")
	symaddr(&addr_QPushButton_paintEngine_default, "QPushButton_paintEngine_default")
	symaddr(&addr_QPushButton_mouseDoubleClickEvent_default, "QPushButton_mouseDoubleClickEvent_default")
	symaddr(&addr_QPushButton_wheelEvent_default, "QPushButton_wheelEvent_default")
	symaddr(&addr_QPushButton_enterEvent_default, "QPushButton_enterEvent_default")
	symaddr(&addr_QPushButton_leaveEvent_default, "QPushButton_leaveEvent_default")
	symaddr(&addr_QPushButton_moveEvent_default, "QPushButton_moveEvent_default")
	symaddr(&addr_QPushButton_resizeEvent_default, "QPushButton_resizeEvent_default")
	symaddr(&addr_QPushButton_closeEvent_default, "QPushButton_closeEvent_default")
	symaddr(&addr_QPushButton_contextMenuEvent_default, "QPushButton_contextMenuEvent_default")
	symaddr(&addr_QPushButton_tabletEvent_default, "QPushButton_tabletEvent_default")
	symaddr(&addr_QPushButton_actionEvent_default, "QPushButton_actionEvent_default")
	symaddr(&addr_QPushButton_dragEnterEvent_default, "QPushButton_dragEnterEvent_default")
	symaddr(&addr_QPushButton_dragMoveEvent_default, "QPushButton_dragMoveEvent_default")
	symaddr(&addr_QPushButton_dragLeaveEvent_default, "QPushButton_dragLeaveEvent_default")
	symaddr(&addr_QPushButton_dropEvent_default, "QPushButton_dropEvent_default")
	symaddr(&addr_QPushButton_showEvent_default, "QPushButton_showEvent_default")
	symaddr(&addr_QPushButton_hideEvent_default, "QPushButton_hideEvent_default")
	symaddr(&addr_QPushButton_metric_default, "QPushButton_metric_default")
	symaddr(&addr_QPushButton_initPainter_default, "QPushButton_initPainter_default")
	symaddr(&addr_QPushButton_redirected_default, "QPushButton_redirected_default")
	symaddr(&addr_QPushButton_sharedPainter_default, "QPushButton_sharedPainter_default")
	symaddr(&addr_QPushButton_inputMethodEvent_default, "QPushButton_inputMethodEvent_default")
	symaddr(&addr_QPushButton_inputMethodQuery_default, "QPushButton_inputMethodQuery_default")
	symaddr(&addr_QPushButton_focusNextPrevChild_default, "QPushButton_focusNextPrevChild_default")
	symaddr(&addr_QPushButton_eventFilter_default, "QPushButton_eventFilter_default")
	symaddr(&addr_QPushButton_childEvent_default, "QPushButton_childEvent_default")
	symaddr(&addr_QPushButton_customEvent_default, "QPushButton_customEvent_default")
	symaddr(&addr_QPushButton_connectNotify_default, "QPushButton_connectNotify_default")
	symaddr(&addr_QPushButton_disconnectNotify_default, "QPushButton_disconnectNotify_default")
	reg(&pg_QPushButton_connect_pressed, "QPushButton_connect_pressed")
	reg(&pg_QPushButton_connect_released, "QPushButton_connect_released")
	reg(&pg_QPushButton_connect_clicked, "QPushButton_connect_clicked")
	reg(&pg_QPushButton_connect_clicked2, "QPushButton_connect_clicked2")
	reg(&pg_QPushButton_connect_toggled, "QPushButton_connect_toggled")
	reg(&pg_QPushButton_connect_windowTitleChanged, "QPushButton_connect_windowTitleChanged")
	reg(&pg_QPushButton_connect_windowIconChanged, "QPushButton_connect_windowIconChanged")
	reg(&pg_QPushButton_connect_windowIconTextChanged, "QPushButton_connect_windowIconTextChanged")
	reg(&pg_QPushButton_connect_customContextMenuRequested, "QPushButton_connect_customContextMenuRequested")
	reg(&pg_QPushButton_connect_destroyed, "QPushButton_connect_destroyed")
	reg(&pg_QPushButton_connect_destroyed2, "QPushButton_connect_destroyed2")
	reg(&pg_QPushButton_connect_objectNameChanged, "QPushButton_connect_objectNameChanged")
	reg(&pg_QPushButton__disconnect, "QPushButton__disconnect")
	reg(&pg_QPushButton__set_override, "QPushButton__set_override")
	reg(&pg_QPushButton_delete, "QPushButton_delete")
}

type QPushButton struct {
	h uintptr
	*QAbstractButton
}

func bindQPushButton(h uintptr) *QPushButton {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QPushButton_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPushButton{h: h, QAbstractButton: &QAbstractButton{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQPushButton(h uintptr) *QPushButton { return bindQPushButton(h) }

func (this *QPushButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPushButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPushButton(h unsafe.Pointer) *QPushButton { return bindQPushButton(uintptr(h)) }
func UnsafeNewQPushButtonOwned(h unsafe.Pointer) *QPushButton { return newownedQPushButton(uintptr(h)) }

func (this *QPushButton) UnsafeRelease() {
	this.h = 0
	if this.QAbstractButton != nil {
		this.QAbstractButton.UnsafeRelease()
	}
}

func NewQPushButton(parent *QWidget) *QPushButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPushButton_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPushButton(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func NewQPushButton2() *QPushButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPushButton_new2, uintptr(__goid))
	__o := newownedQPushButton(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func NewQPushButton3(text string) *QPushButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPushButton_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQPushButton(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func NewQPushButton4(icon *QIcon, text string) *QPushButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPushButton_new4, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__o := newownedQPushButton(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func NewQPushButton5(text string, parent *QWidget) *QPushButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPushButton_new5, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQPushButton(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func NewQPushButton6(icon *QIcon, text string, parent *QWidget) *QPushButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPushButton_new6, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQPushButton(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPushButton_virtIdx, QPushButton_enable)
	return __o
}

func QPushButton_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPushButton_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPushButton) AutoDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_autoDefault, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) SetAutoDefault(arg1 bool) {
	purego.SyscallN(addr_QPushButton_setAutoDefault, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) IsDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_isDefault, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) SetDefault(arg1 bool) {
	purego.SyscallN(addr_QPushButton_setDefault, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) SetMenu(menu *QMenu) {
	purego.SyscallN(addr_QPushButton_setMenu, this.h, uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menu)
}

func (this *QPushButton) Menu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QPushButton_menu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QPushButton) SetFlat(arg1 bool) {
	purego.SyscallN(addr_QPushButton_setFlat, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) IsFlat() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_isFlat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) ShowMenu() {
	purego.SyscallN(addr_QPushButton_showMenu, this.h)
	runtime.KeepAlive(this)
}

func QPushButton_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QPushButton_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QPushButton_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_mouseGrabber)
	return bindQWidget(__r)
}

func QPushButton_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_keyboardGrabber)
	return bindQWidget(__r)
}

func QPushButton_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QPushButton_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QPushButton) UpdateMicroFocus() {
	purego.SyscallN(addr_QPushButton_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QPushButton) Create() {
	purego.SyscallN(addr_QPushButton_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QPushButton) Destroy() {
	purego.SyscallN(addr_QPushButton_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QPushButton) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QPushButton_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPushButton_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPushButton_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPushButton_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPushButton) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPushButton_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPushButton) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPushButton) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPushButton_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPushButton) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPushButton_DevicePixelRatioFScale() float64 {
	__r := pg_QPushButton_devicePixelRatioFScale()
	return __r
}

func QPushButton_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPushButton_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPushButton) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QPushButton_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QPushButton_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPushButton_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPushButton_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPushButton_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPushButton_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QPushButton_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPushButton_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QPushButton) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QPushButton_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QPushButton) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QPushButton_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPushButton) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QPushButton_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QPushButton_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QPushButton_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QPushButton_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QPushButton_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPushButton_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPushButton_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPushButton_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPushButton) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPushButton_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPushButton) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPushButton_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPushButton) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPushButton_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPushButton) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QPushButton) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QPushButton_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QPushButton_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QPushButton_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QPushButton_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QPushButton_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) InitStyleOptionDefault(option *QStyleOptionButton) {
	purego.SyscallN(addr_QPushButton_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QPushButton) HitButtonDefault(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_hitButton_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QPushButton) CheckStateSetDefault() {
	purego.SyscallN(addr_QPushButton_checkStateSet_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QPushButton) NextCheckStateDefault() {
	purego.SyscallN(addr_QPushButton_nextCheckState_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QPushButton) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QPushButton_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPushButton) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPushButton_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPushButton) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPushButton_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPushButton) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QPushButton_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPushButton) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QPushButton_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPushButton) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPushButton) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QPushButton_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QPushButton) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPushButton) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPushButton_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPushButton) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QPushButton_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QPushButton_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QPushButton_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPushButton_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QPushButton_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QPushButton_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QPushButton_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QPushButton_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QPushButton_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QPushButton_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QPushButton_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QPushButton_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QPushButton_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QPushButton_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QPushButton_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QPushButton_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPushButton) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QPushButton_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QPushButton) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPushButton_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QPushButton) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPushButton_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QPushButton) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QPushButton_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPushButton) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPushButton_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPushButton) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPushButton) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPushButton_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPushButton) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPushButton_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPushButton_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPushButton) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPushButton_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPushButton) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPushButton_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPushButton) OnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPushButton_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPushButton_connect_released(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPushButton_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QPushButton_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QPushButton_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QPushButton_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QPushButton_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QPushButton_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QPushButton_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPushButton_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPushButton_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPushButton_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPushButton) Disconnect(handle int64) {
	pg_QPushButton__disconnect(handle)
	unregisterCallback(handle)
}

func QPushButton_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case 5290442163167458465:
		return 4
	case 1023791977679793185:
		return 5
	case 975043767033552552:
		return 6
	case -385596010576048993:
		return 7
	case -4255494004099806581:
		return 8
	case -488693769771308531:
		return 9
	case 5736040353629178930:
		return 10
	case 5747806663224476200:
		return 11
	case 1474647914901213797:
		return 12
	case 772598475950050299:
		return 13
	case -332090898078543319:
		return 14
	case -4415920049297795213:
		return 15
	case 4959376063826165943:
		return 16
	case -5409927581812374498:
		return 17
	case -8659679893868524044:
		return 18
	case 2110182719375245575:
		return 19
	case 4311534060495737339:
		return 20
	case -8266762397995423915:
		return 21
	case 6398007404756346381:
		return 22
	case -4029073905697023887:
		return 23
	case 1919988340934511316:
		return 24
	case 6316270346668237895:
		return 25
	case -551070009623711640:
		return 26
	case -9161949150489525396:
		return 27
	case -4302917285655014661:
		return 28
	case -1989937050019464783:
		return 29
	case -6698043222227006015:
		return 30
	case -212068866860701141:
		return 31
	case 1697444140088555581:
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

func QPushButton_enable(h uintptr, idx int32) { pg_QPushButton__set_override(h, idx) }

func (this *QPushButton) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPushButton) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QPushButton) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QPushButton) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QPushButton) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QPushButton) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QPushButton) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QPushButton) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QPushButton) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QPushButton) OnInitStyleOption(fn func(*QStyleOptionButton)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionButton(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QPushButton) OnHitButton(fn func(*QPoint) bool) {
	registerVirt(this.UnsafeGoID(), 5736040353629178930, func(pos uintptr) int32 { return b2i(fn(UnsafeNewQPoint(unsafe.Pointer(pos)))) }) // slot "hitButton"
}

func (this *QPushButton) OnCheckStateSet(fn func()) {
	registerVirt(this.UnsafeGoID(), 5747806663224476200, func() { fn() }) // slot "checkStateSet"
}

func (this *QPushButton) OnNextCheckState(fn func()) {
	registerVirt(this.UnsafeGoID(), 1474647914901213797, func() { fn() }) // slot "nextCheckState"
}

func (this *QPushButton) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QPushButton) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QPushButton) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QPushButton) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QPushButton) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QPushButton) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QPushButton) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QPushButton) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QPushButton) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QPushButton) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QPushButton) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QPushButton) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QPushButton) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QPushButton) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QPushButton) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QPushButton) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QPushButton) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QPushButton) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QPushButton) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QPushButton) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QPushButton) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QPushButton) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QPushButton) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QPushButton) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QPushButton) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QPushButton) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QPushButton) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QPushButton) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QPushButton) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QPushButton) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QPushButton) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QPushButton) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QPushButton) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QPushButton) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QPushButton) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPushButton) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPushButton) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPushButton) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPushButton) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPushButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPushButton_delete(this.h)
	}
	this.UnsafeRelease()
}
