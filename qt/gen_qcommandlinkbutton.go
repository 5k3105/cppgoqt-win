package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCommandLinkButton_chain uintptr
	addr_QCommandLinkButton_new uintptr
	addr_QCommandLinkButton_new2 uintptr
	addr_QCommandLinkButton_new3 uintptr
	addr_QCommandLinkButton_new4 uintptr
	addr_QCommandLinkButton_new5 uintptr
	addr_QCommandLinkButton_new6 uintptr
	addr_QCommandLinkButton_tr uintptr
	addr_QCommandLinkButton_description uintptr
	addr_QCommandLinkButton_setDescription uintptr
	addr_QCommandLinkButton_setTabOrder uintptr
	addr_QCommandLinkButton_mouseGrabber uintptr
	addr_QCommandLinkButton_keyboardGrabber uintptr
	addr_QCommandLinkButton_find uintptr
	addr_QCommandLinkButton_createWindowContainer uintptr
	addr_QCommandLinkButton_updateMicroFocus uintptr
	addr_QCommandLinkButton_create uintptr
	addr_QCommandLinkButton_destroy uintptr
	addr_QCommandLinkButton_focusNextChild uintptr
	addr_QCommandLinkButton_focusPreviousChild uintptr
	addr_QCommandLinkButton_connect uintptr
	addr_QCommandLinkButton_connect2 uintptr
	addr_QCommandLinkButton_sender uintptr
	addr_QCommandLinkButton_senderSignalIndex uintptr
	addr_QCommandLinkButton_receivers uintptr
	addr_QCommandLinkButton_isSignalConnected uintptr
	pg_QCommandLinkButton_devicePixelRatioFScale func() float64
	pg_QCommandLinkButton_encodeMetricF func(int32, float64) int32
	pg_QCommandLinkButton_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QCommandLinkButton_tr2 uintptr
	addr_QCommandLinkButton_tr3 uintptr
	addr_QCommandLinkButton_createWindowContainer2 uintptr
	addr_QCommandLinkButton_createWindowContainer3 uintptr
	addr_QCommandLinkButton_updateMicroFocus2 uintptr
	addr_QCommandLinkButton_create2 uintptr
	addr_QCommandLinkButton_create3 uintptr
	addr_QCommandLinkButton_create4 uintptr
	addr_QCommandLinkButton_destroy2 uintptr
	addr_QCommandLinkButton_destroy3 uintptr
	addr_QCommandLinkButton_connect3 uintptr
	addr_QCommandLinkButton_connect4 uintptr
	addr_QCommandLinkButton_metaObject_default uintptr
	addr_QCommandLinkButton_sizeHint_default uintptr
	addr_QCommandLinkButton_heightForWidth_default uintptr
	addr_QCommandLinkButton_minimumSizeHint_default uintptr
	addr_QCommandLinkButton_initStyleOption_default uintptr
	addr_QCommandLinkButton_event_default uintptr
	addr_QCommandLinkButton_paintEvent_default uintptr
	addr_QCommandLinkButton_keyPressEvent_default uintptr
	addr_QCommandLinkButton_focusInEvent_default uintptr
	addr_QCommandLinkButton_focusOutEvent_default uintptr
	addr_QCommandLinkButton_mouseMoveEvent_default uintptr
	addr_QCommandLinkButton_hitButton_default uintptr
	addr_QCommandLinkButton_checkStateSet_default uintptr
	addr_QCommandLinkButton_nextCheckState_default uintptr
	addr_QCommandLinkButton_keyReleaseEvent_default uintptr
	addr_QCommandLinkButton_mousePressEvent_default uintptr
	addr_QCommandLinkButton_mouseReleaseEvent_default uintptr
	addr_QCommandLinkButton_changeEvent_default uintptr
	addr_QCommandLinkButton_timerEvent_default uintptr
	addr_QCommandLinkButton_devType_default uintptr
	addr_QCommandLinkButton_setVisible_default uintptr
	addr_QCommandLinkButton_hasHeightForWidth_default uintptr
	addr_QCommandLinkButton_paintEngine_default uintptr
	addr_QCommandLinkButton_mouseDoubleClickEvent_default uintptr
	addr_QCommandLinkButton_wheelEvent_default uintptr
	addr_QCommandLinkButton_enterEvent_default uintptr
	addr_QCommandLinkButton_leaveEvent_default uintptr
	addr_QCommandLinkButton_moveEvent_default uintptr
	addr_QCommandLinkButton_resizeEvent_default uintptr
	addr_QCommandLinkButton_closeEvent_default uintptr
	addr_QCommandLinkButton_contextMenuEvent_default uintptr
	addr_QCommandLinkButton_tabletEvent_default uintptr
	addr_QCommandLinkButton_actionEvent_default uintptr
	addr_QCommandLinkButton_dragEnterEvent_default uintptr
	addr_QCommandLinkButton_dragMoveEvent_default uintptr
	addr_QCommandLinkButton_dragLeaveEvent_default uintptr
	addr_QCommandLinkButton_dropEvent_default uintptr
	addr_QCommandLinkButton_showEvent_default uintptr
	addr_QCommandLinkButton_hideEvent_default uintptr
	addr_QCommandLinkButton_metric_default uintptr
	addr_QCommandLinkButton_initPainter_default uintptr
	addr_QCommandLinkButton_redirected_default uintptr
	addr_QCommandLinkButton_sharedPainter_default uintptr
	addr_QCommandLinkButton_inputMethodEvent_default uintptr
	addr_QCommandLinkButton_inputMethodQuery_default uintptr
	addr_QCommandLinkButton_focusNextPrevChild_default uintptr
	addr_QCommandLinkButton_eventFilter_default uintptr
	addr_QCommandLinkButton_childEvent_default uintptr
	addr_QCommandLinkButton_customEvent_default uintptr
	addr_QCommandLinkButton_connectNotify_default uintptr
	addr_QCommandLinkButton_disconnectNotify_default uintptr
	pg_QCommandLinkButton_connect_pressed func(uintptr, int64)
	pg_QCommandLinkButton_connect_released func(uintptr, int64)
	pg_QCommandLinkButton_connect_clicked func(uintptr, int64)
	pg_QCommandLinkButton_connect_clicked2 func(uintptr, int64)
	pg_QCommandLinkButton_connect_toggled func(uintptr, int64)
	pg_QCommandLinkButton_connect_windowTitleChanged func(uintptr, int64)
	pg_QCommandLinkButton_connect_windowIconChanged func(uintptr, int64)
	pg_QCommandLinkButton_connect_windowIconTextChanged func(uintptr, int64)
	pg_QCommandLinkButton_connect_customContextMenuRequested func(uintptr, int64)
	pg_QCommandLinkButton_connect_destroyed func(uintptr, int64)
	pg_QCommandLinkButton_connect_destroyed2 func(uintptr, int64)
	pg_QCommandLinkButton_connect_objectNameChanged func(uintptr, int64)
	pg_QCommandLinkButton__disconnect func(int64)
	pg_QCommandLinkButton__set_override func(uintptr, int32)
	pg_QCommandLinkButton_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCommandLinkButton_chain, "QCommandLinkButton_chain")
	symaddr(&addr_QCommandLinkButton_new, "QCommandLinkButton_new")
	symaddr(&addr_QCommandLinkButton_new2, "QCommandLinkButton_new2")
	symaddr(&addr_QCommandLinkButton_new3, "QCommandLinkButton_new3")
	symaddr(&addr_QCommandLinkButton_new4, "QCommandLinkButton_new4")
	symaddr(&addr_QCommandLinkButton_new5, "QCommandLinkButton_new5")
	symaddr(&addr_QCommandLinkButton_new6, "QCommandLinkButton_new6")
	symaddr(&addr_QCommandLinkButton_tr, "QCommandLinkButton_tr")
	symaddr(&addr_QCommandLinkButton_description, "QCommandLinkButton_description")
	symaddr(&addr_QCommandLinkButton_setDescription, "QCommandLinkButton_setDescription")
	symaddr(&addr_QCommandLinkButton_setTabOrder, "QCommandLinkButton_setTabOrder")
	symaddr(&addr_QCommandLinkButton_mouseGrabber, "QCommandLinkButton_mouseGrabber")
	symaddr(&addr_QCommandLinkButton_keyboardGrabber, "QCommandLinkButton_keyboardGrabber")
	symaddr(&addr_QCommandLinkButton_find, "QCommandLinkButton_find")
	symaddr(&addr_QCommandLinkButton_createWindowContainer, "QCommandLinkButton_createWindowContainer")
	symaddr(&addr_QCommandLinkButton_updateMicroFocus, "QCommandLinkButton_updateMicroFocus")
	symaddr(&addr_QCommandLinkButton_create, "QCommandLinkButton_create")
	symaddr(&addr_QCommandLinkButton_destroy, "QCommandLinkButton_destroy")
	symaddr(&addr_QCommandLinkButton_focusNextChild, "QCommandLinkButton_focusNextChild")
	symaddr(&addr_QCommandLinkButton_focusPreviousChild, "QCommandLinkButton_focusPreviousChild")
	symaddr(&addr_QCommandLinkButton_connect, "QCommandLinkButton_connect")
	symaddr(&addr_QCommandLinkButton_connect2, "QCommandLinkButton_connect2")
	symaddr(&addr_QCommandLinkButton_sender, "QCommandLinkButton_sender")
	symaddr(&addr_QCommandLinkButton_senderSignalIndex, "QCommandLinkButton_senderSignalIndex")
	symaddr(&addr_QCommandLinkButton_receivers, "QCommandLinkButton_receivers")
	symaddr(&addr_QCommandLinkButton_isSignalConnected, "QCommandLinkButton_isSignalConnected")
	reg(&pg_QCommandLinkButton_devicePixelRatioFScale, "QCommandLinkButton_devicePixelRatioFScale")
	reg(&pg_QCommandLinkButton_encodeMetricF, "QCommandLinkButton_encodeMetricF")
	reg(&pg_QCommandLinkButton_getDecodedMetricF, "QCommandLinkButton_getDecodedMetricF")
	symaddr(&addr_QCommandLinkButton_tr2, "QCommandLinkButton_tr2")
	symaddr(&addr_QCommandLinkButton_tr3, "QCommandLinkButton_tr3")
	symaddr(&addr_QCommandLinkButton_createWindowContainer2, "QCommandLinkButton_createWindowContainer2")
	symaddr(&addr_QCommandLinkButton_createWindowContainer3, "QCommandLinkButton_createWindowContainer3")
	symaddr(&addr_QCommandLinkButton_updateMicroFocus2, "QCommandLinkButton_updateMicroFocus2")
	symaddr(&addr_QCommandLinkButton_create2, "QCommandLinkButton_create2")
	symaddr(&addr_QCommandLinkButton_create3, "QCommandLinkButton_create3")
	symaddr(&addr_QCommandLinkButton_create4, "QCommandLinkButton_create4")
	symaddr(&addr_QCommandLinkButton_destroy2, "QCommandLinkButton_destroy2")
	symaddr(&addr_QCommandLinkButton_destroy3, "QCommandLinkButton_destroy3")
	symaddr(&addr_QCommandLinkButton_connect3, "QCommandLinkButton_connect3")
	symaddr(&addr_QCommandLinkButton_connect4, "QCommandLinkButton_connect4")
	symaddr(&addr_QCommandLinkButton_metaObject_default, "QCommandLinkButton_metaObject_default")
	symaddr(&addr_QCommandLinkButton_sizeHint_default, "QCommandLinkButton_sizeHint_default")
	symaddr(&addr_QCommandLinkButton_heightForWidth_default, "QCommandLinkButton_heightForWidth_default")
	symaddr(&addr_QCommandLinkButton_minimumSizeHint_default, "QCommandLinkButton_minimumSizeHint_default")
	symaddr(&addr_QCommandLinkButton_initStyleOption_default, "QCommandLinkButton_initStyleOption_default")
	symaddr(&addr_QCommandLinkButton_event_default, "QCommandLinkButton_event_default")
	symaddr(&addr_QCommandLinkButton_paintEvent_default, "QCommandLinkButton_paintEvent_default")
	symaddr(&addr_QCommandLinkButton_keyPressEvent_default, "QCommandLinkButton_keyPressEvent_default")
	symaddr(&addr_QCommandLinkButton_focusInEvent_default, "QCommandLinkButton_focusInEvent_default")
	symaddr(&addr_QCommandLinkButton_focusOutEvent_default, "QCommandLinkButton_focusOutEvent_default")
	symaddr(&addr_QCommandLinkButton_mouseMoveEvent_default, "QCommandLinkButton_mouseMoveEvent_default")
	symaddr(&addr_QCommandLinkButton_hitButton_default, "QCommandLinkButton_hitButton_default")
	symaddr(&addr_QCommandLinkButton_checkStateSet_default, "QCommandLinkButton_checkStateSet_default")
	symaddr(&addr_QCommandLinkButton_nextCheckState_default, "QCommandLinkButton_nextCheckState_default")
	symaddr(&addr_QCommandLinkButton_keyReleaseEvent_default, "QCommandLinkButton_keyReleaseEvent_default")
	symaddr(&addr_QCommandLinkButton_mousePressEvent_default, "QCommandLinkButton_mousePressEvent_default")
	symaddr(&addr_QCommandLinkButton_mouseReleaseEvent_default, "QCommandLinkButton_mouseReleaseEvent_default")
	symaddr(&addr_QCommandLinkButton_changeEvent_default, "QCommandLinkButton_changeEvent_default")
	symaddr(&addr_QCommandLinkButton_timerEvent_default, "QCommandLinkButton_timerEvent_default")
	symaddr(&addr_QCommandLinkButton_devType_default, "QCommandLinkButton_devType_default")
	symaddr(&addr_QCommandLinkButton_setVisible_default, "QCommandLinkButton_setVisible_default")
	symaddr(&addr_QCommandLinkButton_hasHeightForWidth_default, "QCommandLinkButton_hasHeightForWidth_default")
	symaddr(&addr_QCommandLinkButton_paintEngine_default, "QCommandLinkButton_paintEngine_default")
	symaddr(&addr_QCommandLinkButton_mouseDoubleClickEvent_default, "QCommandLinkButton_mouseDoubleClickEvent_default")
	symaddr(&addr_QCommandLinkButton_wheelEvent_default, "QCommandLinkButton_wheelEvent_default")
	symaddr(&addr_QCommandLinkButton_enterEvent_default, "QCommandLinkButton_enterEvent_default")
	symaddr(&addr_QCommandLinkButton_leaveEvent_default, "QCommandLinkButton_leaveEvent_default")
	symaddr(&addr_QCommandLinkButton_moveEvent_default, "QCommandLinkButton_moveEvent_default")
	symaddr(&addr_QCommandLinkButton_resizeEvent_default, "QCommandLinkButton_resizeEvent_default")
	symaddr(&addr_QCommandLinkButton_closeEvent_default, "QCommandLinkButton_closeEvent_default")
	symaddr(&addr_QCommandLinkButton_contextMenuEvent_default, "QCommandLinkButton_contextMenuEvent_default")
	symaddr(&addr_QCommandLinkButton_tabletEvent_default, "QCommandLinkButton_tabletEvent_default")
	symaddr(&addr_QCommandLinkButton_actionEvent_default, "QCommandLinkButton_actionEvent_default")
	symaddr(&addr_QCommandLinkButton_dragEnterEvent_default, "QCommandLinkButton_dragEnterEvent_default")
	symaddr(&addr_QCommandLinkButton_dragMoveEvent_default, "QCommandLinkButton_dragMoveEvent_default")
	symaddr(&addr_QCommandLinkButton_dragLeaveEvent_default, "QCommandLinkButton_dragLeaveEvent_default")
	symaddr(&addr_QCommandLinkButton_dropEvent_default, "QCommandLinkButton_dropEvent_default")
	symaddr(&addr_QCommandLinkButton_showEvent_default, "QCommandLinkButton_showEvent_default")
	symaddr(&addr_QCommandLinkButton_hideEvent_default, "QCommandLinkButton_hideEvent_default")
	symaddr(&addr_QCommandLinkButton_metric_default, "QCommandLinkButton_metric_default")
	symaddr(&addr_QCommandLinkButton_initPainter_default, "QCommandLinkButton_initPainter_default")
	symaddr(&addr_QCommandLinkButton_redirected_default, "QCommandLinkButton_redirected_default")
	symaddr(&addr_QCommandLinkButton_sharedPainter_default, "QCommandLinkButton_sharedPainter_default")
	symaddr(&addr_QCommandLinkButton_inputMethodEvent_default, "QCommandLinkButton_inputMethodEvent_default")
	symaddr(&addr_QCommandLinkButton_inputMethodQuery_default, "QCommandLinkButton_inputMethodQuery_default")
	symaddr(&addr_QCommandLinkButton_focusNextPrevChild_default, "QCommandLinkButton_focusNextPrevChild_default")
	symaddr(&addr_QCommandLinkButton_eventFilter_default, "QCommandLinkButton_eventFilter_default")
	symaddr(&addr_QCommandLinkButton_childEvent_default, "QCommandLinkButton_childEvent_default")
	symaddr(&addr_QCommandLinkButton_customEvent_default, "QCommandLinkButton_customEvent_default")
	symaddr(&addr_QCommandLinkButton_connectNotify_default, "QCommandLinkButton_connectNotify_default")
	symaddr(&addr_QCommandLinkButton_disconnectNotify_default, "QCommandLinkButton_disconnectNotify_default")
	reg(&pg_QCommandLinkButton_connect_pressed, "QCommandLinkButton_connect_pressed")
	reg(&pg_QCommandLinkButton_connect_released, "QCommandLinkButton_connect_released")
	reg(&pg_QCommandLinkButton_connect_clicked, "QCommandLinkButton_connect_clicked")
	reg(&pg_QCommandLinkButton_connect_clicked2, "QCommandLinkButton_connect_clicked2")
	reg(&pg_QCommandLinkButton_connect_toggled, "QCommandLinkButton_connect_toggled")
	reg(&pg_QCommandLinkButton_connect_windowTitleChanged, "QCommandLinkButton_connect_windowTitleChanged")
	reg(&pg_QCommandLinkButton_connect_windowIconChanged, "QCommandLinkButton_connect_windowIconChanged")
	reg(&pg_QCommandLinkButton_connect_windowIconTextChanged, "QCommandLinkButton_connect_windowIconTextChanged")
	reg(&pg_QCommandLinkButton_connect_customContextMenuRequested, "QCommandLinkButton_connect_customContextMenuRequested")
	reg(&pg_QCommandLinkButton_connect_destroyed, "QCommandLinkButton_connect_destroyed")
	reg(&pg_QCommandLinkButton_connect_destroyed2, "QCommandLinkButton_connect_destroyed2")
	reg(&pg_QCommandLinkButton_connect_objectNameChanged, "QCommandLinkButton_connect_objectNameChanged")
	reg(&pg_QCommandLinkButton__disconnect, "QCommandLinkButton__disconnect")
	reg(&pg_QCommandLinkButton__set_override, "QCommandLinkButton__set_override")
	reg(&pg_QCommandLinkButton_delete, "QCommandLinkButton_delete")
}

type QCommandLinkButton struct {
	h uintptr
	*QPushButton
}

func bindQCommandLinkButton(h uintptr) *QCommandLinkButton {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QCommandLinkButton_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QCommandLinkButton{h: h, QPushButton: &QPushButton{h: __c[0], QAbstractButton: &QAbstractButton{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQCommandLinkButton(h uintptr) *QCommandLinkButton { return bindQCommandLinkButton(h) }

func (this *QCommandLinkButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCommandLinkButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCommandLinkButton(h unsafe.Pointer) *QCommandLinkButton { return bindQCommandLinkButton(uintptr(h)) }
func UnsafeNewQCommandLinkButtonOwned(h unsafe.Pointer) *QCommandLinkButton { return newownedQCommandLinkButton(uintptr(h)) }

func (this *QCommandLinkButton) UnsafeRelease() {
	this.h = 0
	if this.QPushButton != nil {
		this.QPushButton.UnsafeRelease()
	}
}

func NewQCommandLinkButton(parent *QWidget) *QCommandLinkButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQCommandLinkButton(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func NewQCommandLinkButton2() *QCommandLinkButton {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new2, uintptr(__goid))
	__o := newownedQCommandLinkButton(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func NewQCommandLinkButton3(text string) *QCommandLinkButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQCommandLinkButton(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func NewQCommandLinkButton4(text string, description string) *QCommandLinkButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new4, uintptr(__goid), text_p, uintptr(len(text)), description_p, uintptr(len(description)))
	__o := newownedQCommandLinkButton(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(description)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func NewQCommandLinkButton5(text string, parent *QWidget) *QCommandLinkButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new5, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQCommandLinkButton(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func NewQCommandLinkButton6(text string, description string, parent *QWidget) *QCommandLinkButton {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_new6, uintptr(__goid), text_p, uintptr(len(text)), description_p, uintptr(len(description)), uintptr(parent.UnsafePointer()))
	__o := newownedQCommandLinkButton(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(description)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCommandLinkButton_virtIdx, QCommandLinkButton_enable)
	return __o
}

func QCommandLinkButton_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QCommandLinkButton) Description() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_description, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCommandLinkButton) SetDescription(description string) {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QCommandLinkButton_setDescription, this.h, description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(description)
}

func QCommandLinkButton_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QCommandLinkButton_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QCommandLinkButton_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_mouseGrabber)
	return bindQWidget(__r)
}

func QCommandLinkButton_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_keyboardGrabber)
	return bindQWidget(__r)
}

func QCommandLinkButton_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QCommandLinkButton_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QCommandLinkButton) UpdateMicroFocus() {
	purego.SyscallN(addr_QCommandLinkButton_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) Create() {
	purego.SyscallN(addr_QCommandLinkButton_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) Destroy() {
	purego.SyscallN(addr_QCommandLinkButton_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCommandLinkButton) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QCommandLinkButton_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCommandLinkButton_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCommandLinkButton) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QCommandLinkButton) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCommandLinkButton) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QCommandLinkButton) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QCommandLinkButton_DevicePixelRatioFScale() float64 {
	__r := pg_QCommandLinkButton_devicePixelRatioFScale()
	return __r
}

func QCommandLinkButton_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QCommandLinkButton_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QCommandLinkButton) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QCommandLinkButton_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QCommandLinkButton_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCommandLinkButton_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCommandLinkButton_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QCommandLinkButton_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QCommandLinkButton) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QCommandLinkButton_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QCommandLinkButton_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QCommandLinkButton_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QCommandLinkButton_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QCommandLinkButton_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QCommandLinkButton_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QCommandLinkButton_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCommandLinkButton_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCommandLinkButton) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QCommandLinkButton) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCommandLinkButton) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCommandLinkButton) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCommandLinkButton) InitStyleOptionDefault(option *QStyleOptionButton) {
	purego.SyscallN(addr_QCommandLinkButton_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QCommandLinkButton) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QCommandLinkButton) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QCommandLinkButton_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QCommandLinkButton_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QCommandLinkButton_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QCommandLinkButton_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QCommandLinkButton_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) HitButtonDefault(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_hitButton_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QCommandLinkButton) CheckStateSetDefault() {
	purego.SyscallN(addr_QCommandLinkButton_checkStateSet_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) NextCheckStateDefault() {
	purego.SyscallN(addr_QCommandLinkButton_nextCheckState_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QCommandLinkButton_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCommandLinkButton) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QCommandLinkButton_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCommandLinkButton) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QCommandLinkButton_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCommandLinkButton) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QCommandLinkButton_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCommandLinkButton) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QCommandLinkButton_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCommandLinkButton) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCommandLinkButton) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QCommandLinkButton_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QCommandLinkButton) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCommandLinkButton) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QCommandLinkButton) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCommandLinkButton_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QCommandLinkButton_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QCommandLinkButton_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCommandLinkButton_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QCommandLinkButton_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QCommandLinkButton_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QCommandLinkButton_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QCommandLinkButton_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QCommandLinkButton_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QCommandLinkButton_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QCommandLinkButton_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QCommandLinkButton_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QCommandLinkButton_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QCommandLinkButton_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QCommandLinkButton_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QCommandLinkButton_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCommandLinkButton) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QCommandLinkButton_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QCommandLinkButton) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QCommandLinkButton) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QCommandLinkButton) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QCommandLinkButton_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCommandLinkButton) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCommandLinkButton_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCommandLinkButton) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCommandLinkButton) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLinkButton_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCommandLinkButton) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QCommandLinkButton_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCommandLinkButton_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCommandLinkButton) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCommandLinkButton_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCommandLinkButton) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCommandLinkButton_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCommandLinkButton) OnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCommandLinkButton_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCommandLinkButton_connect_released(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCommandLinkButton_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QCommandLinkButton_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QCommandLinkButton_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QCommandLinkButton_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QCommandLinkButton_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QCommandLinkButton_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QCommandLinkButton_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCommandLinkButton_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCommandLinkButton_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCommandLinkButton_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCommandLinkButton) Disconnect(handle int64) {
	pg_QCommandLinkButton__disconnect(handle)
	unregisterCallback(handle)
}

func QCommandLinkButton_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 4311534060495737339:
		return 2
	case 7049264599837532491:
		return 3
	case -488693769771308531:
		return 4
	case 1237236456312824991:
		return 5
	case 5290442163167458465:
		return 6
	case 1023791977679793185:
		return 7
	case 975043767033552552:
		return 8
	case -385596010576048993:
		return 9
	case -4255494004099806581:
		return 10
	case 5736040353629178930:
		return 11
	case 5747806663224476200:
		return 12
	case 1474647914901213797:
		return 13
	case 772598475950050299:
		return 14
	case -332090898078543319:
		return 15
	case -4415920049297795213:
		return 16
	case 4959376063826165943:
		return 17
	case -5409927581812374498:
		return 18
	case -8659679893868524044:
		return 19
	case 2110182719375245575:
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

func QCommandLinkButton_enable(h uintptr, idx int32) { pg_QCommandLinkButton__set_override(h, idx) }

func (this *QCommandLinkButton) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QCommandLinkButton) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QCommandLinkButton) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QCommandLinkButton) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QCommandLinkButton) OnInitStyleOption(fn func(*QStyleOptionButton)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionButton(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QCommandLinkButton) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QCommandLinkButton) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QCommandLinkButton) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QCommandLinkButton) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QCommandLinkButton) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QCommandLinkButton) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QCommandLinkButton) OnHitButton(fn func(*QPoint) bool) {
	registerVirt(this.UnsafeGoID(), 5736040353629178930, func(pos uintptr) int32 { return b2i(fn(UnsafeNewQPoint(unsafe.Pointer(pos)))) }) // slot "hitButton"
}

func (this *QCommandLinkButton) OnCheckStateSet(fn func()) {
	registerVirt(this.UnsafeGoID(), 5747806663224476200, func() { fn() }) // slot "checkStateSet"
}

func (this *QCommandLinkButton) OnNextCheckState(fn func()) {
	registerVirt(this.UnsafeGoID(), 1474647914901213797, func() { fn() }) // slot "nextCheckState"
}

func (this *QCommandLinkButton) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QCommandLinkButton) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QCommandLinkButton) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QCommandLinkButton) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QCommandLinkButton) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QCommandLinkButton) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QCommandLinkButton) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QCommandLinkButton) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QCommandLinkButton) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QCommandLinkButton) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QCommandLinkButton) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QCommandLinkButton) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QCommandLinkButton) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QCommandLinkButton) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QCommandLinkButton) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QCommandLinkButton) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QCommandLinkButton) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QCommandLinkButton) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QCommandLinkButton) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QCommandLinkButton) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QCommandLinkButton) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QCommandLinkButton) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QCommandLinkButton) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QCommandLinkButton) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QCommandLinkButton) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QCommandLinkButton) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QCommandLinkButton) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QCommandLinkButton) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QCommandLinkButton) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QCommandLinkButton) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QCommandLinkButton) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QCommandLinkButton) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QCommandLinkButton) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QCommandLinkButton) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QCommandLinkButton) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QCommandLinkButton) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QCommandLinkButton) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QCommandLinkButton) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QCommandLinkButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCommandLinkButton_delete(this.h)
	}
	this.UnsafeRelease()
}
