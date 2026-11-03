package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDialog_chain uintptr
	addr_QDialog_new uintptr
	addr_QDialog_new2 uintptr
	addr_QDialog_new3 uintptr
	addr_QDialog_tr uintptr
	addr_QDialog_result uintptr
	addr_QDialog_setSizeGripEnabled uintptr
	addr_QDialog_isSizeGripEnabled uintptr
	addr_QDialog_setModal uintptr
	addr_QDialog_setResult uintptr
	addr_QDialog_finished uintptr
	addr_QDialog_accepted uintptr
	addr_QDialog_rejected uintptr
	addr_QDialog_open uintptr
	addr_QDialog_exec uintptr
	addr_QDialog_done uintptr
	addr_QDialog_accept uintptr
	addr_QDialog_reject uintptr
	addr_QDialog_adjustPosition uintptr
	addr_QDialog_setTabOrder uintptr
	addr_QDialog_mouseGrabber uintptr
	addr_QDialog_keyboardGrabber uintptr
	addr_QDialog_find uintptr
	addr_QDialog_createWindowContainer uintptr
	addr_QDialog_updateMicroFocus uintptr
	addr_QDialog_create uintptr
	addr_QDialog_destroy uintptr
	addr_QDialog_focusNextChild uintptr
	addr_QDialog_focusPreviousChild uintptr
	addr_QDialog_connect uintptr
	addr_QDialog_connect2 uintptr
	addr_QDialog_sender uintptr
	addr_QDialog_senderSignalIndex uintptr
	addr_QDialog_receivers uintptr
	addr_QDialog_isSignalConnected uintptr
	pg_QDialog_devicePixelRatioFScale func() float64
	pg_QDialog_encodeMetricF func(int32, float64) int32
	pg_QDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDialog_tr2 uintptr
	addr_QDialog_tr3 uintptr
	addr_QDialog_createWindowContainer2 uintptr
	addr_QDialog_createWindowContainer3 uintptr
	addr_QDialog_updateMicroFocus2 uintptr
	addr_QDialog_create2 uintptr
	addr_QDialog_create3 uintptr
	addr_QDialog_create4 uintptr
	addr_QDialog_destroy2 uintptr
	addr_QDialog_destroy3 uintptr
	addr_QDialog_connect3 uintptr
	addr_QDialog_connect4 uintptr
	addr_QDialog_open_default uintptr
	addr_QDialog_exec_default uintptr
	addr_QDialog_done_default uintptr
	addr_QDialog_accept_default uintptr
	addr_QDialog_reject_default uintptr
	addr_QDialog_metaObject_default uintptr
	addr_QDialog_setVisible_default uintptr
	addr_QDialog_sizeHint_default uintptr
	addr_QDialog_minimumSizeHint_default uintptr
	addr_QDialog_keyPressEvent_default uintptr
	addr_QDialog_closeEvent_default uintptr
	addr_QDialog_showEvent_default uintptr
	addr_QDialog_resizeEvent_default uintptr
	addr_QDialog_contextMenuEvent_default uintptr
	addr_QDialog_eventFilter_default uintptr
	addr_QDialog_devType_default uintptr
	addr_QDialog_heightForWidth_default uintptr
	addr_QDialog_hasHeightForWidth_default uintptr
	addr_QDialog_paintEngine_default uintptr
	addr_QDialog_event_default uintptr
	addr_QDialog_mousePressEvent_default uintptr
	addr_QDialog_mouseReleaseEvent_default uintptr
	addr_QDialog_mouseDoubleClickEvent_default uintptr
	addr_QDialog_mouseMoveEvent_default uintptr
	addr_QDialog_wheelEvent_default uintptr
	addr_QDialog_keyReleaseEvent_default uintptr
	addr_QDialog_focusInEvent_default uintptr
	addr_QDialog_focusOutEvent_default uintptr
	addr_QDialog_enterEvent_default uintptr
	addr_QDialog_leaveEvent_default uintptr
	addr_QDialog_paintEvent_default uintptr
	addr_QDialog_moveEvent_default uintptr
	addr_QDialog_tabletEvent_default uintptr
	addr_QDialog_actionEvent_default uintptr
	addr_QDialog_dragEnterEvent_default uintptr
	addr_QDialog_dragMoveEvent_default uintptr
	addr_QDialog_dragLeaveEvent_default uintptr
	addr_QDialog_dropEvent_default uintptr
	addr_QDialog_hideEvent_default uintptr
	addr_QDialog_changeEvent_default uintptr
	addr_QDialog_metric_default uintptr
	addr_QDialog_initPainter_default uintptr
	addr_QDialog_redirected_default uintptr
	addr_QDialog_sharedPainter_default uintptr
	addr_QDialog_inputMethodEvent_default uintptr
	addr_QDialog_inputMethodQuery_default uintptr
	addr_QDialog_focusNextPrevChild_default uintptr
	addr_QDialog_timerEvent_default uintptr
	addr_QDialog_childEvent_default uintptr
	addr_QDialog_customEvent_default uintptr
	addr_QDialog_connectNotify_default uintptr
	addr_QDialog_disconnectNotify_default uintptr
	pg_QDialog_connect_finished func(uintptr, int64)
	pg_QDialog_connect_accepted func(uintptr, int64)
	pg_QDialog_connect_rejected func(uintptr, int64)
	pg_QDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDialog_connect_destroyed func(uintptr, int64)
	pg_QDialog_connect_destroyed2 func(uintptr, int64)
	pg_QDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QDialog__disconnect func(int64)
	pg_QDialog__set_override func(uintptr, int32)
	pg_QDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDialog_chain, "QDialog_chain")
	symaddr(&addr_QDialog_new, "QDialog_new")
	symaddr(&addr_QDialog_new2, "QDialog_new2")
	symaddr(&addr_QDialog_new3, "QDialog_new3")
	symaddr(&addr_QDialog_tr, "QDialog_tr")
	symaddr(&addr_QDialog_result, "QDialog_result")
	symaddr(&addr_QDialog_setSizeGripEnabled, "QDialog_setSizeGripEnabled")
	symaddr(&addr_QDialog_isSizeGripEnabled, "QDialog_isSizeGripEnabled")
	symaddr(&addr_QDialog_setModal, "QDialog_setModal")
	symaddr(&addr_QDialog_setResult, "QDialog_setResult")
	symaddr(&addr_QDialog_finished, "QDialog_finished")
	symaddr(&addr_QDialog_accepted, "QDialog_accepted")
	symaddr(&addr_QDialog_rejected, "QDialog_rejected")
	symaddr(&addr_QDialog_open, "QDialog_open")
	symaddr(&addr_QDialog_exec, "QDialog_exec")
	symaddr(&addr_QDialog_done, "QDialog_done")
	symaddr(&addr_QDialog_accept, "QDialog_accept")
	symaddr(&addr_QDialog_reject, "QDialog_reject")
	symaddr(&addr_QDialog_adjustPosition, "QDialog_adjustPosition")
	symaddr(&addr_QDialog_setTabOrder, "QDialog_setTabOrder")
	symaddr(&addr_QDialog_mouseGrabber, "QDialog_mouseGrabber")
	symaddr(&addr_QDialog_keyboardGrabber, "QDialog_keyboardGrabber")
	symaddr(&addr_QDialog_find, "QDialog_find")
	symaddr(&addr_QDialog_createWindowContainer, "QDialog_createWindowContainer")
	symaddr(&addr_QDialog_updateMicroFocus, "QDialog_updateMicroFocus")
	symaddr(&addr_QDialog_create, "QDialog_create")
	symaddr(&addr_QDialog_destroy, "QDialog_destroy")
	symaddr(&addr_QDialog_focusNextChild, "QDialog_focusNextChild")
	symaddr(&addr_QDialog_focusPreviousChild, "QDialog_focusPreviousChild")
	symaddr(&addr_QDialog_connect, "QDialog_connect")
	symaddr(&addr_QDialog_connect2, "QDialog_connect2")
	symaddr(&addr_QDialog_sender, "QDialog_sender")
	symaddr(&addr_QDialog_senderSignalIndex, "QDialog_senderSignalIndex")
	symaddr(&addr_QDialog_receivers, "QDialog_receivers")
	symaddr(&addr_QDialog_isSignalConnected, "QDialog_isSignalConnected")
	reg(&pg_QDialog_devicePixelRatioFScale, "QDialog_devicePixelRatioFScale")
	reg(&pg_QDialog_encodeMetricF, "QDialog_encodeMetricF")
	reg(&pg_QDialog_getDecodedMetricF, "QDialog_getDecodedMetricF")
	symaddr(&addr_QDialog_tr2, "QDialog_tr2")
	symaddr(&addr_QDialog_tr3, "QDialog_tr3")
	symaddr(&addr_QDialog_createWindowContainer2, "QDialog_createWindowContainer2")
	symaddr(&addr_QDialog_createWindowContainer3, "QDialog_createWindowContainer3")
	symaddr(&addr_QDialog_updateMicroFocus2, "QDialog_updateMicroFocus2")
	symaddr(&addr_QDialog_create2, "QDialog_create2")
	symaddr(&addr_QDialog_create3, "QDialog_create3")
	symaddr(&addr_QDialog_create4, "QDialog_create4")
	symaddr(&addr_QDialog_destroy2, "QDialog_destroy2")
	symaddr(&addr_QDialog_destroy3, "QDialog_destroy3")
	symaddr(&addr_QDialog_connect3, "QDialog_connect3")
	symaddr(&addr_QDialog_connect4, "QDialog_connect4")
	symaddr(&addr_QDialog_open_default, "QDialog_open_default")
	symaddr(&addr_QDialog_exec_default, "QDialog_exec_default")
	symaddr(&addr_QDialog_done_default, "QDialog_done_default")
	symaddr(&addr_QDialog_accept_default, "QDialog_accept_default")
	symaddr(&addr_QDialog_reject_default, "QDialog_reject_default")
	symaddr(&addr_QDialog_metaObject_default, "QDialog_metaObject_default")
	symaddr(&addr_QDialog_setVisible_default, "QDialog_setVisible_default")
	symaddr(&addr_QDialog_sizeHint_default, "QDialog_sizeHint_default")
	symaddr(&addr_QDialog_minimumSizeHint_default, "QDialog_minimumSizeHint_default")
	symaddr(&addr_QDialog_keyPressEvent_default, "QDialog_keyPressEvent_default")
	symaddr(&addr_QDialog_closeEvent_default, "QDialog_closeEvent_default")
	symaddr(&addr_QDialog_showEvent_default, "QDialog_showEvent_default")
	symaddr(&addr_QDialog_resizeEvent_default, "QDialog_resizeEvent_default")
	symaddr(&addr_QDialog_contextMenuEvent_default, "QDialog_contextMenuEvent_default")
	symaddr(&addr_QDialog_eventFilter_default, "QDialog_eventFilter_default")
	symaddr(&addr_QDialog_devType_default, "QDialog_devType_default")
	symaddr(&addr_QDialog_heightForWidth_default, "QDialog_heightForWidth_default")
	symaddr(&addr_QDialog_hasHeightForWidth_default, "QDialog_hasHeightForWidth_default")
	symaddr(&addr_QDialog_paintEngine_default, "QDialog_paintEngine_default")
	symaddr(&addr_QDialog_event_default, "QDialog_event_default")
	symaddr(&addr_QDialog_mousePressEvent_default, "QDialog_mousePressEvent_default")
	symaddr(&addr_QDialog_mouseReleaseEvent_default, "QDialog_mouseReleaseEvent_default")
	symaddr(&addr_QDialog_mouseDoubleClickEvent_default, "QDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QDialog_mouseMoveEvent_default, "QDialog_mouseMoveEvent_default")
	symaddr(&addr_QDialog_wheelEvent_default, "QDialog_wheelEvent_default")
	symaddr(&addr_QDialog_keyReleaseEvent_default, "QDialog_keyReleaseEvent_default")
	symaddr(&addr_QDialog_focusInEvent_default, "QDialog_focusInEvent_default")
	symaddr(&addr_QDialog_focusOutEvent_default, "QDialog_focusOutEvent_default")
	symaddr(&addr_QDialog_enterEvent_default, "QDialog_enterEvent_default")
	symaddr(&addr_QDialog_leaveEvent_default, "QDialog_leaveEvent_default")
	symaddr(&addr_QDialog_paintEvent_default, "QDialog_paintEvent_default")
	symaddr(&addr_QDialog_moveEvent_default, "QDialog_moveEvent_default")
	symaddr(&addr_QDialog_tabletEvent_default, "QDialog_tabletEvent_default")
	symaddr(&addr_QDialog_actionEvent_default, "QDialog_actionEvent_default")
	symaddr(&addr_QDialog_dragEnterEvent_default, "QDialog_dragEnterEvent_default")
	symaddr(&addr_QDialog_dragMoveEvent_default, "QDialog_dragMoveEvent_default")
	symaddr(&addr_QDialog_dragLeaveEvent_default, "QDialog_dragLeaveEvent_default")
	symaddr(&addr_QDialog_dropEvent_default, "QDialog_dropEvent_default")
	symaddr(&addr_QDialog_hideEvent_default, "QDialog_hideEvent_default")
	symaddr(&addr_QDialog_changeEvent_default, "QDialog_changeEvent_default")
	symaddr(&addr_QDialog_metric_default, "QDialog_metric_default")
	symaddr(&addr_QDialog_initPainter_default, "QDialog_initPainter_default")
	symaddr(&addr_QDialog_redirected_default, "QDialog_redirected_default")
	symaddr(&addr_QDialog_sharedPainter_default, "QDialog_sharedPainter_default")
	symaddr(&addr_QDialog_inputMethodEvent_default, "QDialog_inputMethodEvent_default")
	symaddr(&addr_QDialog_inputMethodQuery_default, "QDialog_inputMethodQuery_default")
	symaddr(&addr_QDialog_focusNextPrevChild_default, "QDialog_focusNextPrevChild_default")
	symaddr(&addr_QDialog_timerEvent_default, "QDialog_timerEvent_default")
	symaddr(&addr_QDialog_childEvent_default, "QDialog_childEvent_default")
	symaddr(&addr_QDialog_customEvent_default, "QDialog_customEvent_default")
	symaddr(&addr_QDialog_connectNotify_default, "QDialog_connectNotify_default")
	symaddr(&addr_QDialog_disconnectNotify_default, "QDialog_disconnectNotify_default")
	reg(&pg_QDialog_connect_finished, "QDialog_connect_finished")
	reg(&pg_QDialog_connect_accepted, "QDialog_connect_accepted")
	reg(&pg_QDialog_connect_rejected, "QDialog_connect_rejected")
	reg(&pg_QDialog_connect_windowTitleChanged, "QDialog_connect_windowTitleChanged")
	reg(&pg_QDialog_connect_windowIconChanged, "QDialog_connect_windowIconChanged")
	reg(&pg_QDialog_connect_windowIconTextChanged, "QDialog_connect_windowIconTextChanged")
	reg(&pg_QDialog_connect_customContextMenuRequested, "QDialog_connect_customContextMenuRequested")
	reg(&pg_QDialog_connect_destroyed, "QDialog_connect_destroyed")
	reg(&pg_QDialog_connect_destroyed2, "QDialog_connect_destroyed2")
	reg(&pg_QDialog_connect_objectNameChanged, "QDialog_connect_objectNameChanged")
	reg(&pg_QDialog__disconnect, "QDialog__disconnect")
	reg(&pg_QDialog__set_override, "QDialog__set_override")
	reg(&pg_QDialog_delete, "QDialog_delete")
}

type QDialog struct {
	h uintptr
	*QWidget
}

func bindQDialog(h uintptr) *QDialog {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDialog{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQDialog(h uintptr) *QDialog { return bindQDialog(h) }

func (this *QDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDialog(h unsafe.Pointer) *QDialog { return bindQDialog(uintptr(h)) }
func UnsafeNewQDialogOwned(h unsafe.Pointer) *QDialog { return newownedQDialog(uintptr(h)) }

func (this *QDialog) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQDialog(parent *QWidget) *QDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialog_virtIdx, QDialog_enable)
	return __o
}

func NewQDialog2() *QDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialog_new2, uintptr(__goid))
	__o := newownedQDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialog_virtIdx, QDialog_enable)
	return __o
}

func NewQDialog3(parent *QWidget, f WindowType) *QDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDialog_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDialog_virtIdx, QDialog_enable)
	return __o
}

func QDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDialog) Result() int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_result, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) SetSizeGripEnabled(arg1 bool) {
	purego.SyscallN(addr_QDialog_setSizeGripEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QDialog) IsSizeGripEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_isSizeGripEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialog) SetModal(modal bool) {
	purego.SyscallN(addr_QDialog_setModal, this.h, uintptr(b2i(modal)))
	runtime.KeepAlive(this)
}

func (this *QDialog) SetResult(r int) {
	purego.SyscallN(addr_QDialog_setResult, this.h, uintptr(r))
	runtime.KeepAlive(this)
}

func (this *QDialog) Finished(result int) {
	purego.SyscallN(addr_QDialog_finished, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QDialog) Accepted() {
	purego.SyscallN(addr_QDialog_accepted, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Rejected() {
	purego.SyscallN(addr_QDialog_rejected, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Open() {
	purego.SyscallN(addr_QDialog_open, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_exec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) Done(arg1 int) {
	purego.SyscallN(addr_QDialog_done, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDialog) Accept() {
	purego.SyscallN(addr_QDialog_accept, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Reject() {
	purego.SyscallN(addr_QDialog_reject, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Create() {
	purego.SyscallN(addr_QDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) Destroy() {
	purego.SyscallN(addr_QDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QDialog_devicePixelRatioFScale()
	return __r
}

func QDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDialog) OpenDefault() {
	purego.SyscallN(addr_QDialog_open_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) DoneDefault(arg1 int) {
	purego.SyscallN(addr_QDialog_done_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDialog) AcceptDefault() {
	purego.SyscallN(addr_QDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) RejectDefault() {
	purego.SyscallN(addr_QDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QDialog_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QDialog_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QDialog_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QDialog_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDialog) Disconnect(handle int64) {
	pg_QDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 2110182719375245575:
		return 1
	case 729690261457684101:
		return 2
	case 7049264599837532491:
		return 3
	case -554057428128861975:
		return 4
	case -1745322561779819684:
		return 5
	case -2570997490807640431:
		return 6
	case 9106131928185084681:
		return 7
	case 2875220963396170150:
		return 8
	case 1023791977679793185:
		return 9
	case -1989937050019464783:
		return 10
	case -6825369174751916:
		return 11
	case -4302917285655014661:
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
	case 4959376063826165943:
		return 40
	case 427629649267311777:
		return 41
	case 6477500296980050210:
		return 42
	case 4127240633310655048:
		return 43
	case 5792771328741774991:
		return 44
	case 5438819661942843222:
		return 45
	case 7345853097340709638:
		return 46
	case 5898328517505784191:
		return 47
	case -5409927581812374498:
		return 48
	case -1185315367526857833:
		return 49
	case 3769868229123868474:
		return 50
	case 7379029982170503008:
		return 51
	case -568369772982215962:
		return 52
	}
	return -1
}

func QDialog_enable(h uintptr, idx int32) { pg_QDialog__set_override(h, idx) }

func (this *QDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDialog) OnOpen(fn func()) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func() { fn() }) // slot "open"
}

func (this *QDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(arg1 int32) { fn(int(arg1)) }) // slot "done"
}

func (this *QDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
