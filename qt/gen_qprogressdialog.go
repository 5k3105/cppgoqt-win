package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProgressDialog_chain uintptr
	addr_QProgressDialog_new uintptr
	addr_QProgressDialog_new2 uintptr
	addr_QProgressDialog_new3 uintptr
	addr_QProgressDialog_new4 uintptr
	addr_QProgressDialog_new5 uintptr
	addr_QProgressDialog_new6 uintptr
	addr_QProgressDialog_tr uintptr
	addr_QProgressDialog_setLabel uintptr
	addr_QProgressDialog_setCancelButton uintptr
	addr_QProgressDialog_setBar uintptr
	addr_QProgressDialog_wasCanceled uintptr
	addr_QProgressDialog_minimum uintptr
	addr_QProgressDialog_maximum uintptr
	addr_QProgressDialog_value uintptr
	addr_QProgressDialog_labelText uintptr
	addr_QProgressDialog_minimumDuration uintptr
	addr_QProgressDialog_setAutoReset uintptr
	addr_QProgressDialog_autoReset uintptr
	addr_QProgressDialog_setAutoClose uintptr
	addr_QProgressDialog_autoClose uintptr
	addr_QProgressDialog_open uintptr
	addr_QProgressDialog_cancel uintptr
	addr_QProgressDialog_reset uintptr
	addr_QProgressDialog_setMaximum uintptr
	addr_QProgressDialog_setMinimum uintptr
	addr_QProgressDialog_setRange uintptr
	addr_QProgressDialog_setValue uintptr
	addr_QProgressDialog_setLabelText uintptr
	addr_QProgressDialog_setCancelButtonText uintptr
	addr_QProgressDialog_setMinimumDuration uintptr
	addr_QProgressDialog_canceled uintptr
	addr_QProgressDialog_forceShow uintptr
	addr_QProgressDialog_adjustPosition uintptr
	addr_QProgressDialog_setTabOrder uintptr
	addr_QProgressDialog_mouseGrabber uintptr
	addr_QProgressDialog_keyboardGrabber uintptr
	addr_QProgressDialog_find uintptr
	addr_QProgressDialog_createWindowContainer uintptr
	addr_QProgressDialog_updateMicroFocus uintptr
	addr_QProgressDialog_create uintptr
	addr_QProgressDialog_destroy uintptr
	addr_QProgressDialog_focusNextChild uintptr
	addr_QProgressDialog_focusPreviousChild uintptr
	addr_QProgressDialog_connect uintptr
	addr_QProgressDialog_connect2 uintptr
	addr_QProgressDialog_sender uintptr
	addr_QProgressDialog_senderSignalIndex uintptr
	addr_QProgressDialog_receivers uintptr
	addr_QProgressDialog_isSignalConnected uintptr
	pg_QProgressDialog_devicePixelRatioFScale func() float64
	pg_QProgressDialog_encodeMetricF func(int32, float64) int32
	pg_QProgressDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QProgressDialog_tr2 uintptr
	addr_QProgressDialog_tr3 uintptr
	addr_QProgressDialog_createWindowContainer2 uintptr
	addr_QProgressDialog_createWindowContainer3 uintptr
	addr_QProgressDialog_updateMicroFocus2 uintptr
	addr_QProgressDialog_create2 uintptr
	addr_QProgressDialog_create3 uintptr
	addr_QProgressDialog_create4 uintptr
	addr_QProgressDialog_destroy2 uintptr
	addr_QProgressDialog_destroy3 uintptr
	addr_QProgressDialog_connect3 uintptr
	addr_QProgressDialog_connect4 uintptr
	addr_QProgressDialog_metaObject_default uintptr
	addr_QProgressDialog_sizeHint_default uintptr
	addr_QProgressDialog_resizeEvent_default uintptr
	addr_QProgressDialog_closeEvent_default uintptr
	addr_QProgressDialog_changeEvent_default uintptr
	addr_QProgressDialog_showEvent_default uintptr
	addr_QProgressDialog_setVisible_default uintptr
	addr_QProgressDialog_minimumSizeHint_default uintptr
	addr_QProgressDialog_exec_default uintptr
	addr_QProgressDialog_done_default uintptr
	addr_QProgressDialog_accept_default uintptr
	addr_QProgressDialog_reject_default uintptr
	addr_QProgressDialog_keyPressEvent_default uintptr
	addr_QProgressDialog_contextMenuEvent_default uintptr
	addr_QProgressDialog_eventFilter_default uintptr
	addr_QProgressDialog_devType_default uintptr
	addr_QProgressDialog_heightForWidth_default uintptr
	addr_QProgressDialog_hasHeightForWidth_default uintptr
	addr_QProgressDialog_paintEngine_default uintptr
	addr_QProgressDialog_event_default uintptr
	addr_QProgressDialog_mousePressEvent_default uintptr
	addr_QProgressDialog_mouseReleaseEvent_default uintptr
	addr_QProgressDialog_mouseDoubleClickEvent_default uintptr
	addr_QProgressDialog_mouseMoveEvent_default uintptr
	addr_QProgressDialog_wheelEvent_default uintptr
	addr_QProgressDialog_keyReleaseEvent_default uintptr
	addr_QProgressDialog_focusInEvent_default uintptr
	addr_QProgressDialog_focusOutEvent_default uintptr
	addr_QProgressDialog_enterEvent_default uintptr
	addr_QProgressDialog_leaveEvent_default uintptr
	addr_QProgressDialog_paintEvent_default uintptr
	addr_QProgressDialog_moveEvent_default uintptr
	addr_QProgressDialog_tabletEvent_default uintptr
	addr_QProgressDialog_actionEvent_default uintptr
	addr_QProgressDialog_dragEnterEvent_default uintptr
	addr_QProgressDialog_dragMoveEvent_default uintptr
	addr_QProgressDialog_dragLeaveEvent_default uintptr
	addr_QProgressDialog_dropEvent_default uintptr
	addr_QProgressDialog_hideEvent_default uintptr
	addr_QProgressDialog_metric_default uintptr
	addr_QProgressDialog_initPainter_default uintptr
	addr_QProgressDialog_redirected_default uintptr
	addr_QProgressDialog_sharedPainter_default uintptr
	addr_QProgressDialog_inputMethodEvent_default uintptr
	addr_QProgressDialog_inputMethodQuery_default uintptr
	addr_QProgressDialog_focusNextPrevChild_default uintptr
	addr_QProgressDialog_timerEvent_default uintptr
	addr_QProgressDialog_childEvent_default uintptr
	addr_QProgressDialog_customEvent_default uintptr
	addr_QProgressDialog_connectNotify_default uintptr
	addr_QProgressDialog_disconnectNotify_default uintptr
	pg_QProgressDialog_connect_canceled func(uintptr, int64)
	pg_QProgressDialog_connect_finished func(uintptr, int64)
	pg_QProgressDialog_connect_accepted func(uintptr, int64)
	pg_QProgressDialog_connect_rejected func(uintptr, int64)
	pg_QProgressDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QProgressDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QProgressDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QProgressDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QProgressDialog_connect_destroyed func(uintptr, int64)
	pg_QProgressDialog_connect_destroyed2 func(uintptr, int64)
	pg_QProgressDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QProgressDialog__disconnect func(int64)
	pg_QProgressDialog__set_override func(uintptr, int32)
	pg_QProgressDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProgressDialog_chain, "QProgressDialog_chain")
	symaddr(&addr_QProgressDialog_new, "QProgressDialog_new")
	symaddr(&addr_QProgressDialog_new2, "QProgressDialog_new2")
	symaddr(&addr_QProgressDialog_new3, "QProgressDialog_new3")
	symaddr(&addr_QProgressDialog_new4, "QProgressDialog_new4")
	symaddr(&addr_QProgressDialog_new5, "QProgressDialog_new5")
	symaddr(&addr_QProgressDialog_new6, "QProgressDialog_new6")
	symaddr(&addr_QProgressDialog_tr, "QProgressDialog_tr")
	symaddr(&addr_QProgressDialog_setLabel, "QProgressDialog_setLabel")
	symaddr(&addr_QProgressDialog_setCancelButton, "QProgressDialog_setCancelButton")
	symaddr(&addr_QProgressDialog_setBar, "QProgressDialog_setBar")
	symaddr(&addr_QProgressDialog_wasCanceled, "QProgressDialog_wasCanceled")
	symaddr(&addr_QProgressDialog_minimum, "QProgressDialog_minimum")
	symaddr(&addr_QProgressDialog_maximum, "QProgressDialog_maximum")
	symaddr(&addr_QProgressDialog_value, "QProgressDialog_value")
	symaddr(&addr_QProgressDialog_labelText, "QProgressDialog_labelText")
	symaddr(&addr_QProgressDialog_minimumDuration, "QProgressDialog_minimumDuration")
	symaddr(&addr_QProgressDialog_setAutoReset, "QProgressDialog_setAutoReset")
	symaddr(&addr_QProgressDialog_autoReset, "QProgressDialog_autoReset")
	symaddr(&addr_QProgressDialog_setAutoClose, "QProgressDialog_setAutoClose")
	symaddr(&addr_QProgressDialog_autoClose, "QProgressDialog_autoClose")
	symaddr(&addr_QProgressDialog_open, "QProgressDialog_open")
	symaddr(&addr_QProgressDialog_cancel, "QProgressDialog_cancel")
	symaddr(&addr_QProgressDialog_reset, "QProgressDialog_reset")
	symaddr(&addr_QProgressDialog_setMaximum, "QProgressDialog_setMaximum")
	symaddr(&addr_QProgressDialog_setMinimum, "QProgressDialog_setMinimum")
	symaddr(&addr_QProgressDialog_setRange, "QProgressDialog_setRange")
	symaddr(&addr_QProgressDialog_setValue, "QProgressDialog_setValue")
	symaddr(&addr_QProgressDialog_setLabelText, "QProgressDialog_setLabelText")
	symaddr(&addr_QProgressDialog_setCancelButtonText, "QProgressDialog_setCancelButtonText")
	symaddr(&addr_QProgressDialog_setMinimumDuration, "QProgressDialog_setMinimumDuration")
	symaddr(&addr_QProgressDialog_canceled, "QProgressDialog_canceled")
	symaddr(&addr_QProgressDialog_forceShow, "QProgressDialog_forceShow")
	symaddr(&addr_QProgressDialog_adjustPosition, "QProgressDialog_adjustPosition")
	symaddr(&addr_QProgressDialog_setTabOrder, "QProgressDialog_setTabOrder")
	symaddr(&addr_QProgressDialog_mouseGrabber, "QProgressDialog_mouseGrabber")
	symaddr(&addr_QProgressDialog_keyboardGrabber, "QProgressDialog_keyboardGrabber")
	symaddr(&addr_QProgressDialog_find, "QProgressDialog_find")
	symaddr(&addr_QProgressDialog_createWindowContainer, "QProgressDialog_createWindowContainer")
	symaddr(&addr_QProgressDialog_updateMicroFocus, "QProgressDialog_updateMicroFocus")
	symaddr(&addr_QProgressDialog_create, "QProgressDialog_create")
	symaddr(&addr_QProgressDialog_destroy, "QProgressDialog_destroy")
	symaddr(&addr_QProgressDialog_focusNextChild, "QProgressDialog_focusNextChild")
	symaddr(&addr_QProgressDialog_focusPreviousChild, "QProgressDialog_focusPreviousChild")
	symaddr(&addr_QProgressDialog_connect, "QProgressDialog_connect")
	symaddr(&addr_QProgressDialog_connect2, "QProgressDialog_connect2")
	symaddr(&addr_QProgressDialog_sender, "QProgressDialog_sender")
	symaddr(&addr_QProgressDialog_senderSignalIndex, "QProgressDialog_senderSignalIndex")
	symaddr(&addr_QProgressDialog_receivers, "QProgressDialog_receivers")
	symaddr(&addr_QProgressDialog_isSignalConnected, "QProgressDialog_isSignalConnected")
	reg(&pg_QProgressDialog_devicePixelRatioFScale, "QProgressDialog_devicePixelRatioFScale")
	reg(&pg_QProgressDialog_encodeMetricF, "QProgressDialog_encodeMetricF")
	reg(&pg_QProgressDialog_getDecodedMetricF, "QProgressDialog_getDecodedMetricF")
	symaddr(&addr_QProgressDialog_tr2, "QProgressDialog_tr2")
	symaddr(&addr_QProgressDialog_tr3, "QProgressDialog_tr3")
	symaddr(&addr_QProgressDialog_createWindowContainer2, "QProgressDialog_createWindowContainer2")
	symaddr(&addr_QProgressDialog_createWindowContainer3, "QProgressDialog_createWindowContainer3")
	symaddr(&addr_QProgressDialog_updateMicroFocus2, "QProgressDialog_updateMicroFocus2")
	symaddr(&addr_QProgressDialog_create2, "QProgressDialog_create2")
	symaddr(&addr_QProgressDialog_create3, "QProgressDialog_create3")
	symaddr(&addr_QProgressDialog_create4, "QProgressDialog_create4")
	symaddr(&addr_QProgressDialog_destroy2, "QProgressDialog_destroy2")
	symaddr(&addr_QProgressDialog_destroy3, "QProgressDialog_destroy3")
	symaddr(&addr_QProgressDialog_connect3, "QProgressDialog_connect3")
	symaddr(&addr_QProgressDialog_connect4, "QProgressDialog_connect4")
	symaddr(&addr_QProgressDialog_metaObject_default, "QProgressDialog_metaObject_default")
	symaddr(&addr_QProgressDialog_sizeHint_default, "QProgressDialog_sizeHint_default")
	symaddr(&addr_QProgressDialog_resizeEvent_default, "QProgressDialog_resizeEvent_default")
	symaddr(&addr_QProgressDialog_closeEvent_default, "QProgressDialog_closeEvent_default")
	symaddr(&addr_QProgressDialog_changeEvent_default, "QProgressDialog_changeEvent_default")
	symaddr(&addr_QProgressDialog_showEvent_default, "QProgressDialog_showEvent_default")
	symaddr(&addr_QProgressDialog_setVisible_default, "QProgressDialog_setVisible_default")
	symaddr(&addr_QProgressDialog_minimumSizeHint_default, "QProgressDialog_minimumSizeHint_default")
	symaddr(&addr_QProgressDialog_exec_default, "QProgressDialog_exec_default")
	symaddr(&addr_QProgressDialog_done_default, "QProgressDialog_done_default")
	symaddr(&addr_QProgressDialog_accept_default, "QProgressDialog_accept_default")
	symaddr(&addr_QProgressDialog_reject_default, "QProgressDialog_reject_default")
	symaddr(&addr_QProgressDialog_keyPressEvent_default, "QProgressDialog_keyPressEvent_default")
	symaddr(&addr_QProgressDialog_contextMenuEvent_default, "QProgressDialog_contextMenuEvent_default")
	symaddr(&addr_QProgressDialog_eventFilter_default, "QProgressDialog_eventFilter_default")
	symaddr(&addr_QProgressDialog_devType_default, "QProgressDialog_devType_default")
	symaddr(&addr_QProgressDialog_heightForWidth_default, "QProgressDialog_heightForWidth_default")
	symaddr(&addr_QProgressDialog_hasHeightForWidth_default, "QProgressDialog_hasHeightForWidth_default")
	symaddr(&addr_QProgressDialog_paintEngine_default, "QProgressDialog_paintEngine_default")
	symaddr(&addr_QProgressDialog_event_default, "QProgressDialog_event_default")
	symaddr(&addr_QProgressDialog_mousePressEvent_default, "QProgressDialog_mousePressEvent_default")
	symaddr(&addr_QProgressDialog_mouseReleaseEvent_default, "QProgressDialog_mouseReleaseEvent_default")
	symaddr(&addr_QProgressDialog_mouseDoubleClickEvent_default, "QProgressDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QProgressDialog_mouseMoveEvent_default, "QProgressDialog_mouseMoveEvent_default")
	symaddr(&addr_QProgressDialog_wheelEvent_default, "QProgressDialog_wheelEvent_default")
	symaddr(&addr_QProgressDialog_keyReleaseEvent_default, "QProgressDialog_keyReleaseEvent_default")
	symaddr(&addr_QProgressDialog_focusInEvent_default, "QProgressDialog_focusInEvent_default")
	symaddr(&addr_QProgressDialog_focusOutEvent_default, "QProgressDialog_focusOutEvent_default")
	symaddr(&addr_QProgressDialog_enterEvent_default, "QProgressDialog_enterEvent_default")
	symaddr(&addr_QProgressDialog_leaveEvent_default, "QProgressDialog_leaveEvent_default")
	symaddr(&addr_QProgressDialog_paintEvent_default, "QProgressDialog_paintEvent_default")
	symaddr(&addr_QProgressDialog_moveEvent_default, "QProgressDialog_moveEvent_default")
	symaddr(&addr_QProgressDialog_tabletEvent_default, "QProgressDialog_tabletEvent_default")
	symaddr(&addr_QProgressDialog_actionEvent_default, "QProgressDialog_actionEvent_default")
	symaddr(&addr_QProgressDialog_dragEnterEvent_default, "QProgressDialog_dragEnterEvent_default")
	symaddr(&addr_QProgressDialog_dragMoveEvent_default, "QProgressDialog_dragMoveEvent_default")
	symaddr(&addr_QProgressDialog_dragLeaveEvent_default, "QProgressDialog_dragLeaveEvent_default")
	symaddr(&addr_QProgressDialog_dropEvent_default, "QProgressDialog_dropEvent_default")
	symaddr(&addr_QProgressDialog_hideEvent_default, "QProgressDialog_hideEvent_default")
	symaddr(&addr_QProgressDialog_metric_default, "QProgressDialog_metric_default")
	symaddr(&addr_QProgressDialog_initPainter_default, "QProgressDialog_initPainter_default")
	symaddr(&addr_QProgressDialog_redirected_default, "QProgressDialog_redirected_default")
	symaddr(&addr_QProgressDialog_sharedPainter_default, "QProgressDialog_sharedPainter_default")
	symaddr(&addr_QProgressDialog_inputMethodEvent_default, "QProgressDialog_inputMethodEvent_default")
	symaddr(&addr_QProgressDialog_inputMethodQuery_default, "QProgressDialog_inputMethodQuery_default")
	symaddr(&addr_QProgressDialog_focusNextPrevChild_default, "QProgressDialog_focusNextPrevChild_default")
	symaddr(&addr_QProgressDialog_timerEvent_default, "QProgressDialog_timerEvent_default")
	symaddr(&addr_QProgressDialog_childEvent_default, "QProgressDialog_childEvent_default")
	symaddr(&addr_QProgressDialog_customEvent_default, "QProgressDialog_customEvent_default")
	symaddr(&addr_QProgressDialog_connectNotify_default, "QProgressDialog_connectNotify_default")
	symaddr(&addr_QProgressDialog_disconnectNotify_default, "QProgressDialog_disconnectNotify_default")
	reg(&pg_QProgressDialog_connect_canceled, "QProgressDialog_connect_canceled")
	reg(&pg_QProgressDialog_connect_finished, "QProgressDialog_connect_finished")
	reg(&pg_QProgressDialog_connect_accepted, "QProgressDialog_connect_accepted")
	reg(&pg_QProgressDialog_connect_rejected, "QProgressDialog_connect_rejected")
	reg(&pg_QProgressDialog_connect_windowTitleChanged, "QProgressDialog_connect_windowTitleChanged")
	reg(&pg_QProgressDialog_connect_windowIconChanged, "QProgressDialog_connect_windowIconChanged")
	reg(&pg_QProgressDialog_connect_windowIconTextChanged, "QProgressDialog_connect_windowIconTextChanged")
	reg(&pg_QProgressDialog_connect_customContextMenuRequested, "QProgressDialog_connect_customContextMenuRequested")
	reg(&pg_QProgressDialog_connect_destroyed, "QProgressDialog_connect_destroyed")
	reg(&pg_QProgressDialog_connect_destroyed2, "QProgressDialog_connect_destroyed2")
	reg(&pg_QProgressDialog_connect_objectNameChanged, "QProgressDialog_connect_objectNameChanged")
	reg(&pg_QProgressDialog__disconnect, "QProgressDialog__disconnect")
	reg(&pg_QProgressDialog__set_override, "QProgressDialog__set_override")
	reg(&pg_QProgressDialog_delete, "QProgressDialog_delete")
}

type QProgressDialog struct {
	h uintptr
	*QDialog
}

func bindQProgressDialog(h uintptr) *QProgressDialog {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QProgressDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QProgressDialog{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQProgressDialog(h uintptr) *QProgressDialog { return bindQProgressDialog(h) }

func (this *QProgressDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProgressDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProgressDialog(h unsafe.Pointer) *QProgressDialog { return bindQProgressDialog(uintptr(h)) }
func UnsafeNewQProgressDialogOwned(h unsafe.Pointer) *QProgressDialog { return newownedQProgressDialog(uintptr(h)) }

func (this *QProgressDialog) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQProgressDialog(parent *QWidget) *QProgressDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQProgressDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func NewQProgressDialog2() *QProgressDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new2, uintptr(__goid))
	__o := newownedQProgressDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func NewQProgressDialog3(labelText string, cancelButtonText string, minimum int, maximum int) *QProgressDialog {
	__goid := nextVirtID()
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	var cancelButtonText_p uintptr
	if len(cancelButtonText) > 0 {
		cancelButtonText_p = uintptr(unsafe.Pointer(unsafe.StringData(cancelButtonText)))
	}
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new3, uintptr(__goid), labelText_p, uintptr(len(labelText)), cancelButtonText_p, uintptr(len(cancelButtonText)), uintptr(minimum), uintptr(maximum))
	__o := newownedQProgressDialog(__r)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(cancelButtonText)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func NewQProgressDialog4(parent *QWidget, flags WindowType) *QProgressDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new4, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQProgressDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func NewQProgressDialog5(labelText string, cancelButtonText string, minimum int, maximum int, parent *QWidget) *QProgressDialog {
	__goid := nextVirtID()
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	var cancelButtonText_p uintptr
	if len(cancelButtonText) > 0 {
		cancelButtonText_p = uintptr(unsafe.Pointer(unsafe.StringData(cancelButtonText)))
	}
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new5, uintptr(__goid), labelText_p, uintptr(len(labelText)), cancelButtonText_p, uintptr(len(cancelButtonText)), uintptr(minimum), uintptr(maximum), uintptr(parent.UnsafePointer()))
	__o := newownedQProgressDialog(__r)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(cancelButtonText)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func NewQProgressDialog6(labelText string, cancelButtonText string, minimum int, maximum int, parent *QWidget, flags WindowType) *QProgressDialog {
	__goid := nextVirtID()
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	var cancelButtonText_p uintptr
	if len(cancelButtonText) > 0 {
		cancelButtonText_p = uintptr(unsafe.Pointer(unsafe.StringData(cancelButtonText)))
	}
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_new6, uintptr(__goid), labelText_p, uintptr(len(labelText)), cancelButtonText_p, uintptr(len(cancelButtonText)), uintptr(minimum), uintptr(maximum), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQProgressDialog(__r)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(cancelButtonText)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressDialog_virtIdx, QProgressDialog_enable)
	return __o
}

func QProgressDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QProgressDialog) SetLabel(label *QLabel) {
	purego.SyscallN(addr_QProgressDialog_setLabel, this.h, uintptr(label.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
}

func (this *QProgressDialog) SetCancelButton(button *QPushButton) {
	purego.SyscallN(addr_QProgressDialog_setCancelButton, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QProgressDialog) SetBar(bar *QProgressBar) {
	purego.SyscallN(addr_QProgressDialog_setBar, this.h, uintptr(bar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bar)
}

func (this *QProgressDialog) WasCanceled() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_wasCanceled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_minimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_maximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) Value() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) LabelText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_labelText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProgressDialog) MinimumDuration() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_minimumDuration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) SetAutoReset(reset bool) {
	purego.SyscallN(addr_QProgressDialog_setAutoReset, this.h, uintptr(b2i(reset)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) AutoReset() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_autoReset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) SetAutoClose(close_ bool) {
	purego.SyscallN(addr_QProgressDialog_setAutoClose, this.h, uintptr(b2i(close_)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) AutoClose() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_autoClose, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QProgressDialog_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func (this *QProgressDialog) Cancel() {
	purego.SyscallN(addr_QProgressDialog_cancel, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Reset() {
	purego.SyscallN(addr_QProgressDialog_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) SetMaximum(maximum int) {
	purego.SyscallN(addr_QProgressDialog_setMaximum, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) SetMinimum(minimum int) {
	purego.SyscallN(addr_QProgressDialog_setMinimum, this.h, uintptr(minimum))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	purego.SyscallN(addr_QProgressDialog_setRange, this.h, uintptr(minimum), uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) SetValue(progress int) {
	purego.SyscallN(addr_QProgressDialog_setValue, this.h, uintptr(progress))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) SetLabelText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QProgressDialog_setLabelText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QProgressDialog) SetCancelButtonText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QProgressDialog_setCancelButtonText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QProgressDialog) SetMinimumDuration(ms int) {
	purego.SyscallN(addr_QProgressDialog_setMinimumDuration, this.h, uintptr(ms))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Canceled() {
	purego.SyscallN(addr_QProgressDialog_canceled, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) ForceShow() {
	purego.SyscallN(addr_QProgressDialog_forceShow, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QProgressDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QProgressDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QProgressDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QProgressDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QProgressDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QProgressDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QProgressDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QProgressDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QProgressDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Create() {
	purego.SyscallN(addr_QProgressDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Destroy() {
	purego.SyscallN(addr_QProgressDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QProgressDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProgressDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProgressDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QProgressDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QProgressDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QProgressDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QProgressDialog_devicePixelRatioFScale()
	return __r
}

func QProgressDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QProgressDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QProgressDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QProgressDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QProgressDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProgressDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProgressDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QProgressDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QProgressDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QProgressDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QProgressDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QProgressDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QProgressDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QProgressDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QProgressDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QProgressDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProgressDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProgressDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QProgressDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QProgressDialog) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QProgressDialog_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QProgressDialog_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProgressDialog_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QProgressDialog_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QProgressDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QProgressDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) DoneDefault(arg1 int) {
	purego.SyscallN(addr_QProgressDialog_done_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) AcceptDefault() {
	purego.SyscallN(addr_QProgressDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) RejectDefault() {
	purego.SyscallN(addr_QProgressDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QProgressDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QProgressDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressDialog) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QProgressDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QProgressDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QProgressDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QProgressDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QProgressDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QProgressDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QProgressDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QProgressDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProgressDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QProgressDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QProgressDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QProgressDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QProgressDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QProgressDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QProgressDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QProgressDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QProgressDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QProgressDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QProgressDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QProgressDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QProgressDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QProgressDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QProgressDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QProgressDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QProgressDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QProgressDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QProgressDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProgressDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProgressDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProgressDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProgressDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProgressDialog) OnCanceled(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProgressDialog_connect_canceled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QProgressDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProgressDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProgressDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QProgressDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QProgressDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QProgressDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QProgressDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProgressDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QProgressDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QProgressDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressDialog) Disconnect(handle int64) {
	pg_QProgressDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QProgressDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case -4302917285655014661:
		return 2
	case -1989937050019464783:
		return 3
	case 4959376063826165943:
		return 4
	case -6825369174751916:
		return 5
	case 2110182719375245575:
		return 6
	case 7049264599837532491:
		return 7
	case -1745322561779819684:
		return 8
	case -2570997490807640431:
		return 9
	case 9106131928185084681:
		return 10
	case 2875220963396170150:
		return 11
	case 1023791977679793185:
		return 12
	case -6698043222227006015:
		return 13
	case 5217623514630858297:
		return 14
	case -8659679893868524044:
		return 15
	case 4311534060495737339:
		return 16
	case -8266762397995423915:
		return 17
	case 6398007404756346381:
		return 18
	case 1237236456312824991:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case -4255494004099806581:
		return 23
	case 1919988340934511316:
		return 24
	case 772598475950050299:
		return 25
	case 975043767033552552:
		return 26
	case -385596010576048993:
		return 27
	case 6316270346668237895:
		return 28
	case -551070009623711640:
		return 29
	case 5290442163167458465:
		return 30
	case -9161949150489525396:
		return 31
	case -212068866860701141:
		return 32
	case 1697444140088555581:
		return 33
	case 206122715919917589:
		return 34
	case -1781545137352745978:
		return 35
	case -3139904158373725006:
		return 36
	case 9107058682754591716:
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
	case -5409927581812374498:
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

func QProgressDialog_enable(h uintptr, idx int32) { pg_QProgressDialog__set_override(h, idx) }

func (this *QProgressDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QProgressDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QProgressDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QProgressDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QProgressDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QProgressDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QProgressDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QProgressDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QProgressDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QProgressDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(arg1 int32) { fn(int(arg1)) }) // slot "done"
}

func (this *QProgressDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QProgressDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QProgressDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QProgressDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QProgressDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QProgressDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QProgressDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QProgressDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QProgressDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QProgressDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QProgressDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QProgressDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QProgressDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QProgressDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QProgressDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QProgressDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QProgressDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QProgressDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QProgressDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QProgressDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QProgressDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QProgressDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QProgressDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QProgressDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QProgressDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QProgressDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QProgressDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QProgressDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QProgressDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QProgressDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QProgressDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QProgressDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QProgressDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QProgressDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QProgressDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QProgressDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QProgressDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QProgressDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QProgressDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QProgressDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QProgressDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QProgressDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QProgressDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProgressDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
