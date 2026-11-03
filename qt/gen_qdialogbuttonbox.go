package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDialogButtonBox_chain uintptr
	addr_QDialogButtonBox_new uintptr
	addr_QDialogButtonBox_new2 uintptr
	addr_QDialogButtonBox_new3 uintptr
	addr_QDialogButtonBox_new4 uintptr
	addr_QDialogButtonBox_new5 uintptr
	addr_QDialogButtonBox_new6 uintptr
	addr_QDialogButtonBox_new7 uintptr
	addr_QDialogButtonBox_new8 uintptr
	addr_QDialogButtonBox_tr uintptr
	addr_QDialogButtonBox_setOrientation uintptr
	addr_QDialogButtonBox_orientation uintptr
	addr_QDialogButtonBox_addButton uintptr
	addr_QDialogButtonBox_addButton2 uintptr
	addr_QDialogButtonBox_addButton3 uintptr
	addr_QDialogButtonBox_removeButton uintptr
	addr_QDialogButtonBox_clear uintptr
	addr_QDialogButtonBox_buttons uintptr
	addr_QDialogButtonBox_buttonRole uintptr
	addr_QDialogButtonBox_setStandardButtons uintptr
	addr_QDialogButtonBox_standardButtons uintptr
	addr_QDialogButtonBox_standardButton uintptr
	addr_QDialogButtonBox_button uintptr
	addr_QDialogButtonBox_setCenterButtons uintptr
	addr_QDialogButtonBox_centerButtons uintptr
	addr_QDialogButtonBox_clicked uintptr
	addr_QDialogButtonBox_accepted uintptr
	addr_QDialogButtonBox_helpRequested uintptr
	addr_QDialogButtonBox_rejected uintptr
	addr_QDialogButtonBox_setTabOrder uintptr
	addr_QDialogButtonBox_mouseGrabber uintptr
	addr_QDialogButtonBox_keyboardGrabber uintptr
	addr_QDialogButtonBox_find uintptr
	addr_QDialogButtonBox_createWindowContainer uintptr
	addr_QDialogButtonBox_updateMicroFocus uintptr
	addr_QDialogButtonBox_create uintptr
	addr_QDialogButtonBox_destroy uintptr
	addr_QDialogButtonBox_focusNextChild uintptr
	addr_QDialogButtonBox_focusPreviousChild uintptr
	addr_QDialogButtonBox_connect uintptr
	addr_QDialogButtonBox_connect2 uintptr
	addr_QDialogButtonBox_sender uintptr
	addr_QDialogButtonBox_senderSignalIndex uintptr
	addr_QDialogButtonBox_receivers uintptr
	addr_QDialogButtonBox_isSignalConnected uintptr
	pg_QDialogButtonBox_devicePixelRatioFScale func() float64
	pg_QDialogButtonBox_encodeMetricF func(int32, float64) int32
	pg_QDialogButtonBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDialogButtonBox_tr2 uintptr
	addr_QDialogButtonBox_tr3 uintptr
	addr_QDialogButtonBox_createWindowContainer2 uintptr
	addr_QDialogButtonBox_createWindowContainer3 uintptr
	addr_QDialogButtonBox_updateMicroFocus2 uintptr
	addr_QDialogButtonBox_create2 uintptr
	addr_QDialogButtonBox_create3 uintptr
	addr_QDialogButtonBox_create4 uintptr
	addr_QDialogButtonBox_destroy2 uintptr
	addr_QDialogButtonBox_destroy3 uintptr
	addr_QDialogButtonBox_connect3 uintptr
	addr_QDialogButtonBox_connect4 uintptr
	addr_QDialogButtonBox_metaObject_default uintptr
	addr_QDialogButtonBox_changeEvent_default uintptr
	addr_QDialogButtonBox_event_default uintptr
	addr_QDialogButtonBox_devType_default uintptr
	addr_QDialogButtonBox_setVisible_default uintptr
	addr_QDialogButtonBox_sizeHint_default uintptr
	addr_QDialogButtonBox_minimumSizeHint_default uintptr
	addr_QDialogButtonBox_heightForWidth_default uintptr
	addr_QDialogButtonBox_hasHeightForWidth_default uintptr
	addr_QDialogButtonBox_paintEngine_default uintptr
	addr_QDialogButtonBox_mousePressEvent_default uintptr
	addr_QDialogButtonBox_mouseReleaseEvent_default uintptr
	addr_QDialogButtonBox_mouseDoubleClickEvent_default uintptr
	addr_QDialogButtonBox_mouseMoveEvent_default uintptr
	addr_QDialogButtonBox_wheelEvent_default uintptr
	addr_QDialogButtonBox_keyPressEvent_default uintptr
	addr_QDialogButtonBox_keyReleaseEvent_default uintptr
	addr_QDialogButtonBox_focusInEvent_default uintptr
	addr_QDialogButtonBox_focusOutEvent_default uintptr
	addr_QDialogButtonBox_enterEvent_default uintptr
	addr_QDialogButtonBox_leaveEvent_default uintptr
	addr_QDialogButtonBox_paintEvent_default uintptr
	addr_QDialogButtonBox_moveEvent_default uintptr
	addr_QDialogButtonBox_resizeEvent_default uintptr
	addr_QDialogButtonBox_closeEvent_default uintptr
	addr_QDialogButtonBox_contextMenuEvent_default uintptr
	addr_QDialogButtonBox_tabletEvent_default uintptr
	addr_QDialogButtonBox_actionEvent_default uintptr
	addr_QDialogButtonBox_dragEnterEvent_default uintptr
	addr_QDialogButtonBox_dragMoveEvent_default uintptr
	addr_QDialogButtonBox_dragLeaveEvent_default uintptr
	addr_QDialogButtonBox_dropEvent_default uintptr
	addr_QDialogButtonBox_showEvent_default uintptr
	addr_QDialogButtonBox_hideEvent_default uintptr
	addr_QDialogButtonBox_metric_default uintptr
	addr_QDialogButtonBox_initPainter_default uintptr
	addr_QDialogButtonBox_redirected_default uintptr
	addr_QDialogButtonBox_sharedPainter_default uintptr
	addr_QDialogButtonBox_inputMethodEvent_default uintptr
	addr_QDialogButtonBox_inputMethodQuery_default uintptr
	addr_QDialogButtonBox_focusNextPrevChild_default uintptr
	addr_QDialogButtonBox_eventFilter_default uintptr
	addr_QDialogButtonBox_timerEvent_default uintptr
	addr_QDialogButtonBox_childEvent_default uintptr
	addr_QDialogButtonBox_customEvent_default uintptr
	addr_QDialogButtonBox_connectNotify_default uintptr
	addr_QDialogButtonBox_disconnectNotify_default uintptr
	pg_QDialogButtonBox_connect_clicked func(uintptr, int64)
	pg_QDialogButtonBox_connect_accepted func(uintptr, int64)
	pg_QDialogButtonBox_connect_helpRequested func(uintptr, int64)
	pg_QDialogButtonBox_connect_rejected func(uintptr, int64)
	pg_QDialogButtonBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QDialogButtonBox_connect_windowIconChanged func(uintptr, int64)
	pg_QDialogButtonBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDialogButtonBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDialogButtonBox_connect_destroyed func(uintptr, int64)
	pg_QDialogButtonBox_connect_destroyed2 func(uintptr, int64)
	pg_QDialogButtonBox_connect_objectNameChanged func(uintptr, int64)
	pg_QDialogButtonBox__disconnect func(int64)
	pg_QDialogButtonBox__set_override func(uintptr, int32)
	pg_QDialogButtonBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDialogButtonBox_chain, "QDialogButtonBox_chain")
	symaddr(&addr_QDialogButtonBox_new, "QDialogButtonBox_new")
	symaddr(&addr_QDialogButtonBox_new2, "QDialogButtonBox_new2")
	symaddr(&addr_QDialogButtonBox_new3, "QDialogButtonBox_new3")
	symaddr(&addr_QDialogButtonBox_new4, "QDialogButtonBox_new4")
	symaddr(&addr_QDialogButtonBox_new5, "QDialogButtonBox_new5")
	symaddr(&addr_QDialogButtonBox_new6, "QDialogButtonBox_new6")
	symaddr(&addr_QDialogButtonBox_new7, "QDialogButtonBox_new7")
	symaddr(&addr_QDialogButtonBox_new8, "QDialogButtonBox_new8")
	symaddr(&addr_QDialogButtonBox_tr, "QDialogButtonBox_tr")
	symaddr(&addr_QDialogButtonBox_setOrientation, "QDialogButtonBox_setOrientation")
	symaddr(&addr_QDialogButtonBox_orientation, "QDialogButtonBox_orientation")
	symaddr(&addr_QDialogButtonBox_addButton, "QDialogButtonBox_addButton")
	symaddr(&addr_QDialogButtonBox_addButton2, "QDialogButtonBox_addButton2")
	symaddr(&addr_QDialogButtonBox_addButton3, "QDialogButtonBox_addButton3")
	symaddr(&addr_QDialogButtonBox_removeButton, "QDialogButtonBox_removeButton")
	symaddr(&addr_QDialogButtonBox_clear, "QDialogButtonBox_clear")
	symaddr(&addr_QDialogButtonBox_buttons, "QDialogButtonBox_buttons")
	symaddr(&addr_QDialogButtonBox_buttonRole, "QDialogButtonBox_buttonRole")
	symaddr(&addr_QDialogButtonBox_setStandardButtons, "QDialogButtonBox_setStandardButtons")
	symaddr(&addr_QDialogButtonBox_standardButtons, "QDialogButtonBox_standardButtons")
	symaddr(&addr_QDialogButtonBox_standardButton, "QDialogButtonBox_standardButton")
	symaddr(&addr_QDialogButtonBox_button, "QDialogButtonBox_button")
	symaddr(&addr_QDialogButtonBox_setCenterButtons, "QDialogButtonBox_setCenterButtons")
	symaddr(&addr_QDialogButtonBox_centerButtons, "QDialogButtonBox_centerButtons")
	symaddr(&addr_QDialogButtonBox_clicked, "QDialogButtonBox_clicked")
	symaddr(&addr_QDialogButtonBox_accepted, "QDialogButtonBox_accepted")
	symaddr(&addr_QDialogButtonBox_helpRequested, "QDialogButtonBox_helpRequested")
	symaddr(&addr_QDialogButtonBox_rejected, "QDialogButtonBox_rejected")
	symaddr(&addr_QDialogButtonBox_setTabOrder, "QDialogButtonBox_setTabOrder")
	symaddr(&addr_QDialogButtonBox_mouseGrabber, "QDialogButtonBox_mouseGrabber")
	symaddr(&addr_QDialogButtonBox_keyboardGrabber, "QDialogButtonBox_keyboardGrabber")
	symaddr(&addr_QDialogButtonBox_find, "QDialogButtonBox_find")
	symaddr(&addr_QDialogButtonBox_createWindowContainer, "QDialogButtonBox_createWindowContainer")
	symaddr(&addr_QDialogButtonBox_updateMicroFocus, "QDialogButtonBox_updateMicroFocus")
	symaddr(&addr_QDialogButtonBox_create, "QDialogButtonBox_create")
	symaddr(&addr_QDialogButtonBox_destroy, "QDialogButtonBox_destroy")
	symaddr(&addr_QDialogButtonBox_focusNextChild, "QDialogButtonBox_focusNextChild")
	symaddr(&addr_QDialogButtonBox_focusPreviousChild, "QDialogButtonBox_focusPreviousChild")
	symaddr(&addr_QDialogButtonBox_connect, "QDialogButtonBox_connect")
	symaddr(&addr_QDialogButtonBox_connect2, "QDialogButtonBox_connect2")
	symaddr(&addr_QDialogButtonBox_sender, "QDialogButtonBox_sender")
	symaddr(&addr_QDialogButtonBox_senderSignalIndex, "QDialogButtonBox_senderSignalIndex")
	symaddr(&addr_QDialogButtonBox_receivers, "QDialogButtonBox_receivers")
	symaddr(&addr_QDialogButtonBox_isSignalConnected, "QDialogButtonBox_isSignalConnected")
	reg(&pg_QDialogButtonBox_devicePixelRatioFScale, "QDialogButtonBox_devicePixelRatioFScale")
	reg(&pg_QDialogButtonBox_encodeMetricF, "QDialogButtonBox_encodeMetricF")
	reg(&pg_QDialogButtonBox_getDecodedMetricF, "QDialogButtonBox_getDecodedMetricF")
	symaddr(&addr_QDialogButtonBox_tr2, "QDialogButtonBox_tr2")
	symaddr(&addr_QDialogButtonBox_tr3, "QDialogButtonBox_tr3")
	symaddr(&addr_QDialogButtonBox_createWindowContainer2, "QDialogButtonBox_createWindowContainer2")
	symaddr(&addr_QDialogButtonBox_createWindowContainer3, "QDialogButtonBox_createWindowContainer3")
	symaddr(&addr_QDialogButtonBox_updateMicroFocus2, "QDialogButtonBox_updateMicroFocus2")
	symaddr(&addr_QDialogButtonBox_create2, "QDialogButtonBox_create2")
	symaddr(&addr_QDialogButtonBox_create3, "QDialogButtonBox_create3")
	symaddr(&addr_QDialogButtonBox_create4, "QDialogButtonBox_create4")
	symaddr(&addr_QDialogButtonBox_destroy2, "QDialogButtonBox_destroy2")
	symaddr(&addr_QDialogButtonBox_destroy3, "QDialogButtonBox_destroy3")
	symaddr(&addr_QDialogButtonBox_connect3, "QDialogButtonBox_connect3")
	symaddr(&addr_QDialogButtonBox_connect4, "QDialogButtonBox_connect4")
	symaddr(&addr_QDialogButtonBox_metaObject_default, "QDialogButtonBox_metaObject_default")
	symaddr(&addr_QDialogButtonBox_changeEvent_default, "QDialogButtonBox_changeEvent_default")
	symaddr(&addr_QDialogButtonBox_event_default, "QDialogButtonBox_event_default")
	symaddr(&addr_QDialogButtonBox_devType_default, "QDialogButtonBox_devType_default")
	symaddr(&addr_QDialogButtonBox_setVisible_default, "QDialogButtonBox_setVisible_default")
	symaddr(&addr_QDialogButtonBox_sizeHint_default, "QDialogButtonBox_sizeHint_default")
	symaddr(&addr_QDialogButtonBox_minimumSizeHint_default, "QDialogButtonBox_minimumSizeHint_default")
	symaddr(&addr_QDialogButtonBox_heightForWidth_default, "QDialogButtonBox_heightForWidth_default")
	symaddr(&addr_QDialogButtonBox_hasHeightForWidth_default, "QDialogButtonBox_hasHeightForWidth_default")
	symaddr(&addr_QDialogButtonBox_paintEngine_default, "QDialogButtonBox_paintEngine_default")
	symaddr(&addr_QDialogButtonBox_mousePressEvent_default, "QDialogButtonBox_mousePressEvent_default")
	symaddr(&addr_QDialogButtonBox_mouseReleaseEvent_default, "QDialogButtonBox_mouseReleaseEvent_default")
	symaddr(&addr_QDialogButtonBox_mouseDoubleClickEvent_default, "QDialogButtonBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QDialogButtonBox_mouseMoveEvent_default, "QDialogButtonBox_mouseMoveEvent_default")
	symaddr(&addr_QDialogButtonBox_wheelEvent_default, "QDialogButtonBox_wheelEvent_default")
	symaddr(&addr_QDialogButtonBox_keyPressEvent_default, "QDialogButtonBox_keyPressEvent_default")
	symaddr(&addr_QDialogButtonBox_keyReleaseEvent_default, "QDialogButtonBox_keyReleaseEvent_default")
	symaddr(&addr_QDialogButtonBox_focusInEvent_default, "QDialogButtonBox_focusInEvent_default")
	symaddr(&addr_QDialogButtonBox_focusOutEvent_default, "QDialogButtonBox_focusOutEvent_default")
	symaddr(&addr_QDialogButtonBox_enterEvent_default, "QDialogButtonBox_enterEvent_default")
	symaddr(&addr_QDialogButtonBox_leaveEvent_default, "QDialogButtonBox_leaveEvent_default")
	symaddr(&addr_QDialogButtonBox_paintEvent_default, "QDialogButtonBox_paintEvent_default")
	symaddr(&addr_QDialogButtonBox_moveEvent_default, "QDialogButtonBox_moveEvent_default")
	symaddr(&addr_QDialogButtonBox_resizeEvent_default, "QDialogButtonBox_resizeEvent_default")
	symaddr(&addr_QDialogButtonBox_closeEvent_default, "QDialogButtonBox_closeEvent_default")
	symaddr(&addr_QDialogButtonBox_contextMenuEvent_default, "QDialogButtonBox_contextMenuEvent_default")
	symaddr(&addr_QDialogButtonBox_tabletEvent_default, "QDialogButtonBox_tabletEvent_default")
	symaddr(&addr_QDialogButtonBox_actionEvent_default, "QDialogButtonBox_actionEvent_default")
	symaddr(&addr_QDialogButtonBox_dragEnterEvent_default, "QDialogButtonBox_dragEnterEvent_default")
	symaddr(&addr_QDialogButtonBox_dragMoveEvent_default, "QDialogButtonBox_dragMoveEvent_default")
	symaddr(&addr_QDialogButtonBox_dragLeaveEvent_default, "QDialogButtonBox_dragLeaveEvent_default")
	symaddr(&addr_QDialogButtonBox_dropEvent_default, "QDialogButtonBox_dropEvent_default")
	symaddr(&addr_QDialogButtonBox_showEvent_default, "QDialogButtonBox_showEvent_default")
	symaddr(&addr_QDialogButtonBox_hideEvent_default, "QDialogButtonBox_hideEvent_default")
	symaddr(&addr_QDialogButtonBox_metric_default, "QDialogButtonBox_metric_default")
	symaddr(&addr_QDialogButtonBox_initPainter_default, "QDialogButtonBox_initPainter_default")
	symaddr(&addr_QDialogButtonBox_redirected_default, "QDialogButtonBox_redirected_default")
	symaddr(&addr_QDialogButtonBox_sharedPainter_default, "QDialogButtonBox_sharedPainter_default")
	symaddr(&addr_QDialogButtonBox_inputMethodEvent_default, "QDialogButtonBox_inputMethodEvent_default")
	symaddr(&addr_QDialogButtonBox_inputMethodQuery_default, "QDialogButtonBox_inputMethodQuery_default")
	symaddr(&addr_QDialogButtonBox_focusNextPrevChild_default, "QDialogButtonBox_focusNextPrevChild_default")
	symaddr(&addr_QDialogButtonBox_eventFilter_default, "QDialogButtonBox_eventFilter_default")
	symaddr(&addr_QDialogButtonBox_timerEvent_default, "QDialogButtonBox_timerEvent_default")
	symaddr(&addr_QDialogButtonBox_childEvent_default, "QDialogButtonBox_childEvent_default")
	symaddr(&addr_QDialogButtonBox_customEvent_default, "QDialogButtonBox_customEvent_default")
	symaddr(&addr_QDialogButtonBox_connectNotify_default, "QDialogButtonBox_connectNotify_default")
	symaddr(&addr_QDialogButtonBox_disconnectNotify_default, "QDialogButtonBox_disconnectNotify_default")
	reg(&pg_QDialogButtonBox_connect_clicked, "QDialogButtonBox_connect_clicked")
	reg(&pg_QDialogButtonBox_connect_accepted, "QDialogButtonBox_connect_accepted")
	reg(&pg_QDialogButtonBox_connect_helpRequested, "QDialogButtonBox_connect_helpRequested")
	reg(&pg_QDialogButtonBox_connect_rejected, "QDialogButtonBox_connect_rejected")
	reg(&pg_QDialogButtonBox_connect_windowTitleChanged, "QDialogButtonBox_connect_windowTitleChanged")
	reg(&pg_QDialogButtonBox_connect_windowIconChanged, "QDialogButtonBox_connect_windowIconChanged")
	reg(&pg_QDialogButtonBox_connect_windowIconTextChanged, "QDialogButtonBox_connect_windowIconTextChanged")
	reg(&pg_QDialogButtonBox_connect_customContextMenuRequested, "QDialogButtonBox_connect_customContextMenuRequested")
	reg(&pg_QDialogButtonBox_connect_destroyed, "QDialogButtonBox_connect_destroyed")
	reg(&pg_QDialogButtonBox_connect_destroyed2, "QDialogButtonBox_connect_destroyed2")
	reg(&pg_QDialogButtonBox_connect_objectNameChanged, "QDialogButtonBox_connect_objectNameChanged")
	reg(&pg_QDialogButtonBox__disconnect, "QDialogButtonBox__disconnect")
	reg(&pg_QDialogButtonBox__set_override, "QDialogButtonBox__set_override")
	reg(&pg_QDialogButtonBox_delete, "QDialogButtonBox_delete")
}

type QDialogButtonBox struct {
	h uintptr
	*QWidget
}

func bindQDialogButtonBox(h uintptr) *QDialogButtonBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QDialogButtonBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDialogButtonBox{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQDialogButtonBox(h uintptr) *QDialogButtonBox { return bindQDialogButtonBox(h) }

func (this *QDialogButtonBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDialogButtonBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDialogButtonBox(h unsafe.Pointer) *QDialogButtonBox { return bindQDialogButtonBox(uintptr(h)) }
func UnsafeNewQDialogButtonBoxOwned(h unsafe.Pointer) *QDialogButtonBox { return newownedQDialogButtonBox(uintptr(h)) }

func (this *QDialogButtonBox) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQDialogButtonBox(parent *QWidget) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDialogButtonBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox2() *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new2, uintptr(__goid))
	__o := newownedQDialogButtonBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox3(orientation Orientation) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new3, uintptr(__goid), uintptr(orientation))
	__o := newownedQDialogButtonBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox4(buttons QDialogButtonBox__StandardButton) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new4, uintptr(__goid), uintptr(buttons))
	__o := newownedQDialogButtonBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox5(buttons QDialogButtonBox__StandardButton, orientation Orientation) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new5, uintptr(__goid), uintptr(buttons), uintptr(orientation))
	__o := newownedQDialogButtonBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox6(orientation Orientation, parent *QWidget) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new6, uintptr(__goid), uintptr(orientation), uintptr(parent.UnsafePointer()))
	__o := newownedQDialogButtonBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox7(buttons QDialogButtonBox__StandardButton, parent *QWidget) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new7, uintptr(__goid), uintptr(buttons), uintptr(parent.UnsafePointer()))
	__o := newownedQDialogButtonBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func NewQDialogButtonBox8(buttons QDialogButtonBox__StandardButton, orientation Orientation, parent *QWidget) *QDialogButtonBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_new8, uintptr(__goid), uintptr(buttons), uintptr(orientation), uintptr(parent.UnsafePointer()))
	__o := newownedQDialogButtonBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialogButtonBox_virtIdx, QDialogButtonBox_enable)
	return __o
}

func QDialogButtonBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDialogButtonBox) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QDialogButtonBox_setOrientation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QDialogButtonBox) AddButton(button *QAbstractButton, role QDialogButtonBox__ButtonRole) {
	purego.SyscallN(addr_QDialogButtonBox_addButton, this.h, uintptr(button.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QDialogButtonBox) AddButton2(text string, role QDialogButtonBox__ButtonRole) *QPushButton {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_addButton2, this.h, text_p, uintptr(len(text)), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQPushButton(__r)
}

func (this *QDialogButtonBox) AddButtonWithButton(button QDialogButtonBox__StandardButton) *QPushButton {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_addButton3, this.h, uintptr(button))
	runtime.KeepAlive(this)
	return bindQPushButton(__r)
}

func (this *QDialogButtonBox) RemoveButton(button *QAbstractButton) {
	purego.SyscallN(addr_QDialogButtonBox_removeButton, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QDialogButtonBox) Clear() {
	purego.SyscallN(addr_QDialogButtonBox_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Buttons() []*QAbstractButton {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_buttons, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAbstractButton, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAbstractButton(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDialogButtonBox) ButtonRole(button *QAbstractButton) QDialogButtonBox__ButtonRole {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_buttonRole, this.h, uintptr(button.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
	return QDialogButtonBox__ButtonRole(__r)
}

func (this *QDialogButtonBox) SetStandardButtons(buttons QDialogButtonBox__StandardButton) {
	purego.SyscallN(addr_QDialogButtonBox_setStandardButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) StandardButtons() QDialogButtonBox__StandardButton {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_standardButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDialogButtonBox__StandardButton(__r)
}

func (this *QDialogButtonBox) StandardButton(button *QAbstractButton) QDialogButtonBox__StandardButton {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_standardButton, this.h, uintptr(button.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
	return QDialogButtonBox__StandardButton(__r)
}

func (this *QDialogButtonBox) Button(which QDialogButtonBox__StandardButton) *QPushButton {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_button, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return bindQPushButton(__r)
}

func (this *QDialogButtonBox) SetCenterButtons(center bool) {
	purego.SyscallN(addr_QDialogButtonBox_setCenterButtons, this.h, uintptr(b2i(center)))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) CenterButtons() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_centerButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialogButtonBox) Clicked(button *QAbstractButton) {
	purego.SyscallN(addr_QDialogButtonBox_clicked, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QDialogButtonBox) Accepted() {
	purego.SyscallN(addr_QDialogButtonBox_accepted, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) HelpRequested() {
	purego.SyscallN(addr_QDialogButtonBox_helpRequested, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Rejected() {
	purego.SyscallN(addr_QDialogButtonBox_rejected, this.h)
	runtime.KeepAlive(this)
}

func QDialogButtonBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDialogButtonBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDialogButtonBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_mouseGrabber)
	return bindQWidget(__r)
}

func QDialogButtonBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QDialogButtonBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDialogButtonBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDialogButtonBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QDialogButtonBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Create() {
	purego.SyscallN(addr_QDialogButtonBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Destroy() {
	purego.SyscallN(addr_QDialogButtonBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialogButtonBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDialogButtonBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDialogButtonBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDialogButtonBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDialogButtonBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialogButtonBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDialogButtonBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDialogButtonBox_DevicePixelRatioFScale() float64 {
	__r := pg_QDialogButtonBox_devicePixelRatioFScale()
	return __r
}

func QDialogButtonBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDialogButtonBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDialogButtonBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDialogButtonBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDialogButtonBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDialogButtonBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDialogButtonBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDialogButtonBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDialogButtonBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDialogButtonBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDialogButtonBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDialogButtonBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDialogButtonBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDialogButtonBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDialogButtonBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDialogButtonBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDialogButtonBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDialogButtonBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDialogButtonBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDialogButtonBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDialogButtonBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialogButtonBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDialogButtonBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDialogButtonBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDialogButtonBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDialogButtonBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialogButtonBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialogButtonBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDialogButtonBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialogButtonBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialogButtonBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialogButtonBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialogButtonBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDialogButtonBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDialogButtonBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDialogButtonBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDialogButtonBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDialogButtonBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDialogButtonBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDialogButtonBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDialogButtonBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDialogButtonBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QDialogButtonBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDialogButtonBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDialogButtonBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDialogButtonBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDialogButtonBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDialogButtonBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDialogButtonBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDialogButtonBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDialogButtonBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDialogButtonBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDialogButtonBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialogButtonBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDialogButtonBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDialogButtonBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDialogButtonBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDialogButtonBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDialogButtonBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialogButtonBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDialogButtonBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDialogButtonBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialogButtonBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialogButtonBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDialogButtonBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDialogButtonBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDialogButtonBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDialogButtonBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialogButtonBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDialogButtonBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDialogButtonBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDialogButtonBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDialogButtonBox) OnClicked(fn func(*QAbstractButton)) int64 {
	id := registerCallback(func(button uintptr) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(button))) })
	pg_QDialogButtonBox_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialogButtonBox_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnHelpRequested(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialogButtonBox_connect_helpRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialogButtonBox_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDialogButtonBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDialogButtonBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDialogButtonBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDialogButtonBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialogButtonBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDialogButtonBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDialogButtonBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialogButtonBox) Disconnect(handle int64) {
	pg_QDialogButtonBox__disconnect(handle)
	unregisterCallback(handle)
}

func QDialogButtonBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 4959376063826165943:
		return 1
	case 1237236456312824991:
		return 2
	case -8659679893868524044:
		return 3
	case 2110182719375245575:
		return 4
	case 729690261457684101:
		return 5
	case 7049264599837532491:
		return 6
	case 4311534060495737339:
		return 7
	case -8266762397995423915:
		return 8
	case 6398007404756346381:
		return 9
	case -332090898078543319:
		return 10
	case -4415920049297795213:
		return 11
	case -4029073905697023887:
		return 12
	case -4255494004099806581:
		return 13
	case 1919988340934511316:
		return 14
	case 1023791977679793185:
		return 15
	case 772598475950050299:
		return 16
	case 975043767033552552:
		return 17
	case -385596010576048993:
		return 18
	case 6316270346668237895:
		return 19
	case -551070009623711640:
		return 20
	case 5290442163167458465:
		return 21
	case -9161949150489525396:
		return 22
	case -4302917285655014661:
		return 23
	case -1989937050019464783:
		return 24
	case -6698043222227006015:
		return 25
	case -212068866860701141:
		return 26
	case 1697444140088555581:
		return 27
	case 206122715919917589:
		return 28
	case -1781545137352745978:
		return 29
	case -3139904158373725006:
		return 30
	case 9107058682754591716:
		return 31
	case -6825369174751916:
		return 32
	case -199718707876719129:
		return 33
	case -2039166586246443706:
		return 34
	case 427629649267311777:
		return 35
	case 6477500296980050210:
		return 36
	case 4127240633310655048:
		return 37
	case 5792771328741774991:
		return 38
	case 5438819661942843222:
		return 39
	case 7345853097340709638:
		return 40
	case 5898328517505784191:
		return 41
	case 5217623514630858297:
		return 42
	case -5409927581812374498:
		return 43
	case -1185315367526857833:
		return 44
	case 3769868229123868474:
		return 45
	case 7379029982170503008:
		return 46
	case -568369772982215962:
		return 47
	}
	return -1
}

func QDialogButtonBox_enable(h uintptr, idx int32) { pg_QDialogButtonBox__set_override(h, idx) }

func (this *QDialogButtonBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDialogButtonBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QDialogButtonBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDialogButtonBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDialogButtonBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDialogButtonBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDialogButtonBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDialogButtonBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDialogButtonBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDialogButtonBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDialogButtonBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDialogButtonBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDialogButtonBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDialogButtonBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDialogButtonBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDialogButtonBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QDialogButtonBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDialogButtonBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDialogButtonBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDialogButtonBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDialogButtonBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDialogButtonBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDialogButtonBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDialogButtonBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QDialogButtonBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDialogButtonBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDialogButtonBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDialogButtonBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDialogButtonBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDialogButtonBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDialogButtonBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDialogButtonBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDialogButtonBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDialogButtonBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDialogButtonBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDialogButtonBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDialogButtonBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDialogButtonBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDialogButtonBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDialogButtonBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDialogButtonBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDialogButtonBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDialogButtonBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDialogButtonBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDialogButtonBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDialogButtonBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDialogButtonBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDialogButtonBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDialogButtonBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDialogButtonBox_delete(this.h)
	}
	this.UnsafeRelease()
}
