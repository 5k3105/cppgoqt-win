package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QErrorMessage_chain uintptr
	addr_QErrorMessage_new uintptr
	addr_QErrorMessage_new2 uintptr
	addr_QErrorMessage_tr uintptr
	addr_QErrorMessage_qtHandler uintptr
	addr_QErrorMessage_showMessage uintptr
	addr_QErrorMessage_showMessage2 uintptr
	addr_QErrorMessage_adjustPosition uintptr
	addr_QErrorMessage_setTabOrder uintptr
	addr_QErrorMessage_mouseGrabber uintptr
	addr_QErrorMessage_keyboardGrabber uintptr
	addr_QErrorMessage_find uintptr
	addr_QErrorMessage_createWindowContainer uintptr
	addr_QErrorMessage_updateMicroFocus uintptr
	addr_QErrorMessage_create uintptr
	addr_QErrorMessage_destroy uintptr
	addr_QErrorMessage_focusNextChild uintptr
	addr_QErrorMessage_focusPreviousChild uintptr
	addr_QErrorMessage_connect uintptr
	addr_QErrorMessage_connect2 uintptr
	addr_QErrorMessage_sender uintptr
	addr_QErrorMessage_senderSignalIndex uintptr
	addr_QErrorMessage_receivers uintptr
	addr_QErrorMessage_isSignalConnected uintptr
	pg_QErrorMessage_devicePixelRatioFScale func() float64
	pg_QErrorMessage_encodeMetricF func(int32, float64) int32
	pg_QErrorMessage_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QErrorMessage_tr2 uintptr
	addr_QErrorMessage_tr3 uintptr
	addr_QErrorMessage_createWindowContainer2 uintptr
	addr_QErrorMessage_createWindowContainer3 uintptr
	addr_QErrorMessage_updateMicroFocus2 uintptr
	addr_QErrorMessage_create2 uintptr
	addr_QErrorMessage_create3 uintptr
	addr_QErrorMessage_create4 uintptr
	addr_QErrorMessage_destroy2 uintptr
	addr_QErrorMessage_destroy3 uintptr
	addr_QErrorMessage_connect3 uintptr
	addr_QErrorMessage_connect4 uintptr
	addr_QErrorMessage_metaObject_default uintptr
	addr_QErrorMessage_done_default uintptr
	addr_QErrorMessage_changeEvent_default uintptr
	addr_QErrorMessage_setVisible_default uintptr
	addr_QErrorMessage_sizeHint_default uintptr
	addr_QErrorMessage_minimumSizeHint_default uintptr
	addr_QErrorMessage_open_default uintptr
	addr_QErrorMessage_exec_default uintptr
	addr_QErrorMessage_accept_default uintptr
	addr_QErrorMessage_reject_default uintptr
	addr_QErrorMessage_keyPressEvent_default uintptr
	addr_QErrorMessage_closeEvent_default uintptr
	addr_QErrorMessage_showEvent_default uintptr
	addr_QErrorMessage_resizeEvent_default uintptr
	addr_QErrorMessage_contextMenuEvent_default uintptr
	addr_QErrorMessage_eventFilter_default uintptr
	addr_QErrorMessage_devType_default uintptr
	addr_QErrorMessage_heightForWidth_default uintptr
	addr_QErrorMessage_hasHeightForWidth_default uintptr
	addr_QErrorMessage_paintEngine_default uintptr
	addr_QErrorMessage_event_default uintptr
	addr_QErrorMessage_mousePressEvent_default uintptr
	addr_QErrorMessage_mouseReleaseEvent_default uintptr
	addr_QErrorMessage_mouseDoubleClickEvent_default uintptr
	addr_QErrorMessage_mouseMoveEvent_default uintptr
	addr_QErrorMessage_wheelEvent_default uintptr
	addr_QErrorMessage_keyReleaseEvent_default uintptr
	addr_QErrorMessage_focusInEvent_default uintptr
	addr_QErrorMessage_focusOutEvent_default uintptr
	addr_QErrorMessage_enterEvent_default uintptr
	addr_QErrorMessage_leaveEvent_default uintptr
	addr_QErrorMessage_paintEvent_default uintptr
	addr_QErrorMessage_moveEvent_default uintptr
	addr_QErrorMessage_tabletEvent_default uintptr
	addr_QErrorMessage_actionEvent_default uintptr
	addr_QErrorMessage_dragEnterEvent_default uintptr
	addr_QErrorMessage_dragMoveEvent_default uintptr
	addr_QErrorMessage_dragLeaveEvent_default uintptr
	addr_QErrorMessage_dropEvent_default uintptr
	addr_QErrorMessage_hideEvent_default uintptr
	addr_QErrorMessage_metric_default uintptr
	addr_QErrorMessage_initPainter_default uintptr
	addr_QErrorMessage_redirected_default uintptr
	addr_QErrorMessage_sharedPainter_default uintptr
	addr_QErrorMessage_inputMethodEvent_default uintptr
	addr_QErrorMessage_inputMethodQuery_default uintptr
	addr_QErrorMessage_focusNextPrevChild_default uintptr
	addr_QErrorMessage_timerEvent_default uintptr
	addr_QErrorMessage_childEvent_default uintptr
	addr_QErrorMessage_customEvent_default uintptr
	addr_QErrorMessage_connectNotify_default uintptr
	addr_QErrorMessage_disconnectNotify_default uintptr
	pg_QErrorMessage_connect_finished func(uintptr, int64)
	pg_QErrorMessage_connect_accepted func(uintptr, int64)
	pg_QErrorMessage_connect_rejected func(uintptr, int64)
	pg_QErrorMessage_connect_windowTitleChanged func(uintptr, int64)
	pg_QErrorMessage_connect_windowIconChanged func(uintptr, int64)
	pg_QErrorMessage_connect_windowIconTextChanged func(uintptr, int64)
	pg_QErrorMessage_connect_customContextMenuRequested func(uintptr, int64)
	pg_QErrorMessage_connect_destroyed func(uintptr, int64)
	pg_QErrorMessage_connect_destroyed2 func(uintptr, int64)
	pg_QErrorMessage_connect_objectNameChanged func(uintptr, int64)
	pg_QErrorMessage__disconnect func(int64)
	pg_QErrorMessage__set_override func(uintptr, int32)
	pg_QErrorMessage_delete func(uintptr)
)

func init() {
	symaddr(&addr_QErrorMessage_chain, "QErrorMessage_chain")
	symaddr(&addr_QErrorMessage_new, "QErrorMessage_new")
	symaddr(&addr_QErrorMessage_new2, "QErrorMessage_new2")
	symaddr(&addr_QErrorMessage_tr, "QErrorMessage_tr")
	symaddr(&addr_QErrorMessage_qtHandler, "QErrorMessage_qtHandler")
	symaddr(&addr_QErrorMessage_showMessage, "QErrorMessage_showMessage")
	symaddr(&addr_QErrorMessage_showMessage2, "QErrorMessage_showMessage2")
	symaddr(&addr_QErrorMessage_adjustPosition, "QErrorMessage_adjustPosition")
	symaddr(&addr_QErrorMessage_setTabOrder, "QErrorMessage_setTabOrder")
	symaddr(&addr_QErrorMessage_mouseGrabber, "QErrorMessage_mouseGrabber")
	symaddr(&addr_QErrorMessage_keyboardGrabber, "QErrorMessage_keyboardGrabber")
	symaddr(&addr_QErrorMessage_find, "QErrorMessage_find")
	symaddr(&addr_QErrorMessage_createWindowContainer, "QErrorMessage_createWindowContainer")
	symaddr(&addr_QErrorMessage_updateMicroFocus, "QErrorMessage_updateMicroFocus")
	symaddr(&addr_QErrorMessage_create, "QErrorMessage_create")
	symaddr(&addr_QErrorMessage_destroy, "QErrorMessage_destroy")
	symaddr(&addr_QErrorMessage_focusNextChild, "QErrorMessage_focusNextChild")
	symaddr(&addr_QErrorMessage_focusPreviousChild, "QErrorMessage_focusPreviousChild")
	symaddr(&addr_QErrorMessage_connect, "QErrorMessage_connect")
	symaddr(&addr_QErrorMessage_connect2, "QErrorMessage_connect2")
	symaddr(&addr_QErrorMessage_sender, "QErrorMessage_sender")
	symaddr(&addr_QErrorMessage_senderSignalIndex, "QErrorMessage_senderSignalIndex")
	symaddr(&addr_QErrorMessage_receivers, "QErrorMessage_receivers")
	symaddr(&addr_QErrorMessage_isSignalConnected, "QErrorMessage_isSignalConnected")
	reg(&pg_QErrorMessage_devicePixelRatioFScale, "QErrorMessage_devicePixelRatioFScale")
	reg(&pg_QErrorMessage_encodeMetricF, "QErrorMessage_encodeMetricF")
	reg(&pg_QErrorMessage_getDecodedMetricF, "QErrorMessage_getDecodedMetricF")
	symaddr(&addr_QErrorMessage_tr2, "QErrorMessage_tr2")
	symaddr(&addr_QErrorMessage_tr3, "QErrorMessage_tr3")
	symaddr(&addr_QErrorMessage_createWindowContainer2, "QErrorMessage_createWindowContainer2")
	symaddr(&addr_QErrorMessage_createWindowContainer3, "QErrorMessage_createWindowContainer3")
	symaddr(&addr_QErrorMessage_updateMicroFocus2, "QErrorMessage_updateMicroFocus2")
	symaddr(&addr_QErrorMessage_create2, "QErrorMessage_create2")
	symaddr(&addr_QErrorMessage_create3, "QErrorMessage_create3")
	symaddr(&addr_QErrorMessage_create4, "QErrorMessage_create4")
	symaddr(&addr_QErrorMessage_destroy2, "QErrorMessage_destroy2")
	symaddr(&addr_QErrorMessage_destroy3, "QErrorMessage_destroy3")
	symaddr(&addr_QErrorMessage_connect3, "QErrorMessage_connect3")
	symaddr(&addr_QErrorMessage_connect4, "QErrorMessage_connect4")
	symaddr(&addr_QErrorMessage_metaObject_default, "QErrorMessage_metaObject_default")
	symaddr(&addr_QErrorMessage_done_default, "QErrorMessage_done_default")
	symaddr(&addr_QErrorMessage_changeEvent_default, "QErrorMessage_changeEvent_default")
	symaddr(&addr_QErrorMessage_setVisible_default, "QErrorMessage_setVisible_default")
	symaddr(&addr_QErrorMessage_sizeHint_default, "QErrorMessage_sizeHint_default")
	symaddr(&addr_QErrorMessage_minimumSizeHint_default, "QErrorMessage_minimumSizeHint_default")
	symaddr(&addr_QErrorMessage_open_default, "QErrorMessage_open_default")
	symaddr(&addr_QErrorMessage_exec_default, "QErrorMessage_exec_default")
	symaddr(&addr_QErrorMessage_accept_default, "QErrorMessage_accept_default")
	symaddr(&addr_QErrorMessage_reject_default, "QErrorMessage_reject_default")
	symaddr(&addr_QErrorMessage_keyPressEvent_default, "QErrorMessage_keyPressEvent_default")
	symaddr(&addr_QErrorMessage_closeEvent_default, "QErrorMessage_closeEvent_default")
	symaddr(&addr_QErrorMessage_showEvent_default, "QErrorMessage_showEvent_default")
	symaddr(&addr_QErrorMessage_resizeEvent_default, "QErrorMessage_resizeEvent_default")
	symaddr(&addr_QErrorMessage_contextMenuEvent_default, "QErrorMessage_contextMenuEvent_default")
	symaddr(&addr_QErrorMessage_eventFilter_default, "QErrorMessage_eventFilter_default")
	symaddr(&addr_QErrorMessage_devType_default, "QErrorMessage_devType_default")
	symaddr(&addr_QErrorMessage_heightForWidth_default, "QErrorMessage_heightForWidth_default")
	symaddr(&addr_QErrorMessage_hasHeightForWidth_default, "QErrorMessage_hasHeightForWidth_default")
	symaddr(&addr_QErrorMessage_paintEngine_default, "QErrorMessage_paintEngine_default")
	symaddr(&addr_QErrorMessage_event_default, "QErrorMessage_event_default")
	symaddr(&addr_QErrorMessage_mousePressEvent_default, "QErrorMessage_mousePressEvent_default")
	symaddr(&addr_QErrorMessage_mouseReleaseEvent_default, "QErrorMessage_mouseReleaseEvent_default")
	symaddr(&addr_QErrorMessage_mouseDoubleClickEvent_default, "QErrorMessage_mouseDoubleClickEvent_default")
	symaddr(&addr_QErrorMessage_mouseMoveEvent_default, "QErrorMessage_mouseMoveEvent_default")
	symaddr(&addr_QErrorMessage_wheelEvent_default, "QErrorMessage_wheelEvent_default")
	symaddr(&addr_QErrorMessage_keyReleaseEvent_default, "QErrorMessage_keyReleaseEvent_default")
	symaddr(&addr_QErrorMessage_focusInEvent_default, "QErrorMessage_focusInEvent_default")
	symaddr(&addr_QErrorMessage_focusOutEvent_default, "QErrorMessage_focusOutEvent_default")
	symaddr(&addr_QErrorMessage_enterEvent_default, "QErrorMessage_enterEvent_default")
	symaddr(&addr_QErrorMessage_leaveEvent_default, "QErrorMessage_leaveEvent_default")
	symaddr(&addr_QErrorMessage_paintEvent_default, "QErrorMessage_paintEvent_default")
	symaddr(&addr_QErrorMessage_moveEvent_default, "QErrorMessage_moveEvent_default")
	symaddr(&addr_QErrorMessage_tabletEvent_default, "QErrorMessage_tabletEvent_default")
	symaddr(&addr_QErrorMessage_actionEvent_default, "QErrorMessage_actionEvent_default")
	symaddr(&addr_QErrorMessage_dragEnterEvent_default, "QErrorMessage_dragEnterEvent_default")
	symaddr(&addr_QErrorMessage_dragMoveEvent_default, "QErrorMessage_dragMoveEvent_default")
	symaddr(&addr_QErrorMessage_dragLeaveEvent_default, "QErrorMessage_dragLeaveEvent_default")
	symaddr(&addr_QErrorMessage_dropEvent_default, "QErrorMessage_dropEvent_default")
	symaddr(&addr_QErrorMessage_hideEvent_default, "QErrorMessage_hideEvent_default")
	symaddr(&addr_QErrorMessage_metric_default, "QErrorMessage_metric_default")
	symaddr(&addr_QErrorMessage_initPainter_default, "QErrorMessage_initPainter_default")
	symaddr(&addr_QErrorMessage_redirected_default, "QErrorMessage_redirected_default")
	symaddr(&addr_QErrorMessage_sharedPainter_default, "QErrorMessage_sharedPainter_default")
	symaddr(&addr_QErrorMessage_inputMethodEvent_default, "QErrorMessage_inputMethodEvent_default")
	symaddr(&addr_QErrorMessage_inputMethodQuery_default, "QErrorMessage_inputMethodQuery_default")
	symaddr(&addr_QErrorMessage_focusNextPrevChild_default, "QErrorMessage_focusNextPrevChild_default")
	symaddr(&addr_QErrorMessage_timerEvent_default, "QErrorMessage_timerEvent_default")
	symaddr(&addr_QErrorMessage_childEvent_default, "QErrorMessage_childEvent_default")
	symaddr(&addr_QErrorMessage_customEvent_default, "QErrorMessage_customEvent_default")
	symaddr(&addr_QErrorMessage_connectNotify_default, "QErrorMessage_connectNotify_default")
	symaddr(&addr_QErrorMessage_disconnectNotify_default, "QErrorMessage_disconnectNotify_default")
	reg(&pg_QErrorMessage_connect_finished, "QErrorMessage_connect_finished")
	reg(&pg_QErrorMessage_connect_accepted, "QErrorMessage_connect_accepted")
	reg(&pg_QErrorMessage_connect_rejected, "QErrorMessage_connect_rejected")
	reg(&pg_QErrorMessage_connect_windowTitleChanged, "QErrorMessage_connect_windowTitleChanged")
	reg(&pg_QErrorMessage_connect_windowIconChanged, "QErrorMessage_connect_windowIconChanged")
	reg(&pg_QErrorMessage_connect_windowIconTextChanged, "QErrorMessage_connect_windowIconTextChanged")
	reg(&pg_QErrorMessage_connect_customContextMenuRequested, "QErrorMessage_connect_customContextMenuRequested")
	reg(&pg_QErrorMessage_connect_destroyed, "QErrorMessage_connect_destroyed")
	reg(&pg_QErrorMessage_connect_destroyed2, "QErrorMessage_connect_destroyed2")
	reg(&pg_QErrorMessage_connect_objectNameChanged, "QErrorMessage_connect_objectNameChanged")
	reg(&pg_QErrorMessage__disconnect, "QErrorMessage__disconnect")
	reg(&pg_QErrorMessage__set_override, "QErrorMessage__set_override")
	reg(&pg_QErrorMessage_delete, "QErrorMessage_delete")
}

type QErrorMessage struct {
	h uintptr
	*QDialog
}

func bindQErrorMessage(h uintptr) *QErrorMessage {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QErrorMessage_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QErrorMessage{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQErrorMessage(h uintptr) *QErrorMessage { return bindQErrorMessage(h) }

func (this *QErrorMessage) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QErrorMessage) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQErrorMessage(h unsafe.Pointer) *QErrorMessage { return bindQErrorMessage(uintptr(h)) }
func UnsafeNewQErrorMessageOwned(h unsafe.Pointer) *QErrorMessage { return newownedQErrorMessage(uintptr(h)) }

func (this *QErrorMessage) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQErrorMessage(parent *QWidget) *QErrorMessage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQErrorMessage(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QErrorMessage_virtIdx, QErrorMessage_enable)
	return __o
}

func NewQErrorMessage2() *QErrorMessage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_new2, uintptr(__goid))
	__o := newownedQErrorMessage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QErrorMessage_virtIdx, QErrorMessage_enable)
	return __o
}

func QErrorMessage_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QErrorMessage_QtHandler() *QErrorMessage {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_qtHandler)
	return bindQErrorMessage(__r)
}

func (this *QErrorMessage) ShowMessage(message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QErrorMessage_showMessage, this.h, message_p, uintptr(len(message)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func (this *QErrorMessage) ShowMessage2(message string, type_ string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	var type__p uintptr
	if len(type_) > 0 {
		type__p = uintptr(unsafe.Pointer(unsafe.StringData(type_)))
	}
	purego.SyscallN(addr_QErrorMessage_showMessage2, this.h, message_p, uintptr(len(message)), type__p, uintptr(len(type_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
	runtime.KeepAlive(type_)
}

func (this *QErrorMessage) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QErrorMessage_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QErrorMessage_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QErrorMessage_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QErrorMessage_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_mouseGrabber)
	return bindQWidget(__r)
}

func QErrorMessage_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_keyboardGrabber)
	return bindQWidget(__r)
}

func QErrorMessage_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QErrorMessage_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QErrorMessage) UpdateMicroFocus() {
	purego.SyscallN(addr_QErrorMessage_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) Create() {
	purego.SyscallN(addr_QErrorMessage_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) Destroy() {
	purego.SyscallN(addr_QErrorMessage_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QErrorMessage) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QErrorMessage_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QErrorMessage_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QErrorMessage) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QErrorMessage) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QErrorMessage) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QErrorMessage) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QErrorMessage_DevicePixelRatioFScale() float64 {
	__r := pg_QErrorMessage_devicePixelRatioFScale()
	return __r
}

func QErrorMessage_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QErrorMessage_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QErrorMessage) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QErrorMessage_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QErrorMessage_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QErrorMessage_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QErrorMessage_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QErrorMessage_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QErrorMessage) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QErrorMessage_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QErrorMessage_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QErrorMessage_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QErrorMessage_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QErrorMessage_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QErrorMessage_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QErrorMessage_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QErrorMessage_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QErrorMessage) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QErrorMessage) DoneDefault(arg1 int) {
	purego.SyscallN(addr_QErrorMessage_done_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QErrorMessage_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QErrorMessage) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QErrorMessage_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QErrorMessage) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QErrorMessage) OpenDefault() {
	purego.SyscallN(addr_QErrorMessage_open_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QErrorMessage) AcceptDefault() {
	purego.SyscallN(addr_QErrorMessage_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) RejectDefault() {
	purego.SyscallN(addr_QErrorMessage_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QErrorMessage) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QErrorMessage_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QErrorMessage_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QErrorMessage_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QErrorMessage_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QErrorMessage_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QErrorMessage) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QErrorMessage) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QErrorMessage) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QErrorMessage) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QErrorMessage) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QErrorMessage) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QErrorMessage_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QErrorMessage_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QErrorMessage_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QErrorMessage_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QErrorMessage_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QErrorMessage_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QErrorMessage_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QErrorMessage_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QErrorMessage_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QErrorMessage_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QErrorMessage_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QErrorMessage_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QErrorMessage_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QErrorMessage_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QErrorMessage_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QErrorMessage_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QErrorMessage_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QErrorMessage_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QErrorMessage_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QErrorMessage) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QErrorMessage_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QErrorMessage) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QErrorMessage) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QErrorMessage) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QErrorMessage_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QErrorMessage) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QErrorMessage_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QErrorMessage) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QErrorMessage_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QErrorMessage) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QErrorMessage_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QErrorMessage_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QErrorMessage_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QErrorMessage) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QErrorMessage_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QErrorMessage) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QErrorMessage_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QErrorMessage) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QErrorMessage_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QErrorMessage_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QErrorMessage_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QErrorMessage_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QErrorMessage_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QErrorMessage_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QErrorMessage_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QErrorMessage_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QErrorMessage_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QErrorMessage_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QErrorMessage) Disconnect(handle int64) {
	pg_QErrorMessage__disconnect(handle)
	unregisterCallback(handle)
}

func QErrorMessage_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -2570997490807640431:
		return 1
	case 4959376063826165943:
		return 2
	case 2110182719375245575:
		return 3
	case 729690261457684101:
		return 4
	case 7049264599837532491:
		return 5
	case -554057428128861975:
		return 6
	case -1745322561779819684:
		return 7
	case 9106131928185084681:
		return 8
	case 2875220963396170150:
		return 9
	case 1023791977679793185:
		return 10
	case -1989937050019464783:
		return 11
	case -6825369174751916:
		return 12
	case -4302917285655014661:
		return 13
	case -6698043222227006015:
		return 14
	case 5217623514630858297:
		return 15
	case -8659679893868524044:
		return 16
	case 4311534060495737339:
		return 17
	case -8266762397995423915:
		return 18
	case 6398007404756346381:
		return 19
	case 1237236456312824991:
		return 20
	case -332090898078543319:
		return 21
	case -4415920049297795213:
		return 22
	case -4029073905697023887:
		return 23
	case -4255494004099806581:
		return 24
	case 1919988340934511316:
		return 25
	case 772598475950050299:
		return 26
	case 975043767033552552:
		return 27
	case -385596010576048993:
		return 28
	case 6316270346668237895:
		return 29
	case -551070009623711640:
		return 30
	case 5290442163167458465:
		return 31
	case -9161949150489525396:
		return 32
	case -212068866860701141:
		return 33
	case 1697444140088555581:
		return 34
	case 206122715919917589:
		return 35
	case -1781545137352745978:
		return 36
	case -3139904158373725006:
		return 37
	case 9107058682754591716:
		return 38
	case -199718707876719129:
		return 39
	case -2039166586246443706:
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

func QErrorMessage_enable(h uintptr, idx int32) { pg_QErrorMessage__set_override(h, idx) }

func (this *QErrorMessage) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QErrorMessage) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(arg1 int32) { fn(int(arg1)) }) // slot "done"
}

func (this *QErrorMessage) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QErrorMessage) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QErrorMessage) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QErrorMessage) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QErrorMessage) OnOpen(fn func()) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func() { fn() }) // slot "open"
}

func (this *QErrorMessage) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QErrorMessage) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QErrorMessage) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QErrorMessage) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QErrorMessage) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QErrorMessage) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QErrorMessage) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QErrorMessage) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QErrorMessage) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QErrorMessage) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QErrorMessage) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QErrorMessage) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QErrorMessage) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QErrorMessage) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QErrorMessage) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QErrorMessage) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QErrorMessage) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QErrorMessage) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QErrorMessage) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QErrorMessage) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QErrorMessage) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QErrorMessage) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QErrorMessage) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QErrorMessage) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QErrorMessage) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QErrorMessage) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QErrorMessage) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QErrorMessage) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QErrorMessage) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QErrorMessage) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QErrorMessage) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QErrorMessage) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QErrorMessage) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QErrorMessage) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QErrorMessage) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QErrorMessage) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QErrorMessage) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QErrorMessage) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QErrorMessage) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QErrorMessage) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QErrorMessage) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QErrorMessage) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QErrorMessage) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QErrorMessage) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QErrorMessage) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QErrorMessage) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QErrorMessage) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QErrorMessage_delete(this.h)
	}
	this.UnsafeRelease()
}
