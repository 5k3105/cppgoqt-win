package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCheckBox_chain uintptr
	addr_QCheckBox_new uintptr
	addr_QCheckBox_new2 uintptr
	addr_QCheckBox_new3 uintptr
	addr_QCheckBox_new4 uintptr
	addr_QCheckBox_tr uintptr
	addr_QCheckBox_setTristate uintptr
	addr_QCheckBox_isTristate uintptr
	addr_QCheckBox_checkState uintptr
	addr_QCheckBox_setCheckState uintptr
	addr_QCheckBox_stateChanged uintptr
	addr_QCheckBox_checkStateChanged uintptr
	addr_QCheckBox_setTabOrder uintptr
	addr_QCheckBox_mouseGrabber uintptr
	addr_QCheckBox_keyboardGrabber uintptr
	addr_QCheckBox_find uintptr
	addr_QCheckBox_createWindowContainer uintptr
	addr_QCheckBox_updateMicroFocus uintptr
	addr_QCheckBox_create uintptr
	addr_QCheckBox_destroy uintptr
	addr_QCheckBox_focusNextChild uintptr
	addr_QCheckBox_focusPreviousChild uintptr
	addr_QCheckBox_connect uintptr
	addr_QCheckBox_connect2 uintptr
	addr_QCheckBox_sender uintptr
	addr_QCheckBox_senderSignalIndex uintptr
	addr_QCheckBox_receivers uintptr
	addr_QCheckBox_isSignalConnected uintptr
	pg_QCheckBox_devicePixelRatioFScale func() float64
	pg_QCheckBox_encodeMetricF func(int32, float64) int32
	pg_QCheckBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QCheckBox_tr2 uintptr
	addr_QCheckBox_tr3 uintptr
	addr_QCheckBox_setTristate2 uintptr
	addr_QCheckBox_createWindowContainer2 uintptr
	addr_QCheckBox_createWindowContainer3 uintptr
	addr_QCheckBox_updateMicroFocus2 uintptr
	addr_QCheckBox_create2 uintptr
	addr_QCheckBox_create3 uintptr
	addr_QCheckBox_create4 uintptr
	addr_QCheckBox_destroy2 uintptr
	addr_QCheckBox_destroy3 uintptr
	addr_QCheckBox_connect3 uintptr
	addr_QCheckBox_connect4 uintptr
	addr_QCheckBox_metaObject_default uintptr
	addr_QCheckBox_sizeHint_default uintptr
	addr_QCheckBox_minimumSizeHint_default uintptr
	addr_QCheckBox_event_default uintptr
	addr_QCheckBox_hitButton_default uintptr
	addr_QCheckBox_checkStateSet_default uintptr
	addr_QCheckBox_nextCheckState_default uintptr
	addr_QCheckBox_paintEvent_default uintptr
	addr_QCheckBox_mouseMoveEvent_default uintptr
	addr_QCheckBox_initStyleOption_default uintptr
	addr_QCheckBox_keyPressEvent_default uintptr
	addr_QCheckBox_keyReleaseEvent_default uintptr
	addr_QCheckBox_mousePressEvent_default uintptr
	addr_QCheckBox_mouseReleaseEvent_default uintptr
	addr_QCheckBox_focusInEvent_default uintptr
	addr_QCheckBox_focusOutEvent_default uintptr
	addr_QCheckBox_changeEvent_default uintptr
	addr_QCheckBox_timerEvent_default uintptr
	addr_QCheckBox_devType_default uintptr
	addr_QCheckBox_setVisible_default uintptr
	addr_QCheckBox_heightForWidth_default uintptr
	addr_QCheckBox_hasHeightForWidth_default uintptr
	addr_QCheckBox_paintEngine_default uintptr
	addr_QCheckBox_mouseDoubleClickEvent_default uintptr
	addr_QCheckBox_wheelEvent_default uintptr
	addr_QCheckBox_enterEvent_default uintptr
	addr_QCheckBox_leaveEvent_default uintptr
	addr_QCheckBox_moveEvent_default uintptr
	addr_QCheckBox_resizeEvent_default uintptr
	addr_QCheckBox_closeEvent_default uintptr
	addr_QCheckBox_contextMenuEvent_default uintptr
	addr_QCheckBox_tabletEvent_default uintptr
	addr_QCheckBox_actionEvent_default uintptr
	addr_QCheckBox_dragEnterEvent_default uintptr
	addr_QCheckBox_dragMoveEvent_default uintptr
	addr_QCheckBox_dragLeaveEvent_default uintptr
	addr_QCheckBox_dropEvent_default uintptr
	addr_QCheckBox_showEvent_default uintptr
	addr_QCheckBox_hideEvent_default uintptr
	addr_QCheckBox_metric_default uintptr
	addr_QCheckBox_initPainter_default uintptr
	addr_QCheckBox_redirected_default uintptr
	addr_QCheckBox_sharedPainter_default uintptr
	addr_QCheckBox_inputMethodEvent_default uintptr
	addr_QCheckBox_inputMethodQuery_default uintptr
	addr_QCheckBox_focusNextPrevChild_default uintptr
	addr_QCheckBox_eventFilter_default uintptr
	addr_QCheckBox_childEvent_default uintptr
	addr_QCheckBox_customEvent_default uintptr
	addr_QCheckBox_connectNotify_default uintptr
	addr_QCheckBox_disconnectNotify_default uintptr
	pg_QCheckBox_connect_stateChanged func(uintptr, int64)
	pg_QCheckBox_connect_checkStateChanged func(uintptr, int64)
	pg_QCheckBox_connect_pressed func(uintptr, int64)
	pg_QCheckBox_connect_released func(uintptr, int64)
	pg_QCheckBox_connect_clicked func(uintptr, int64)
	pg_QCheckBox_connect_clicked2 func(uintptr, int64)
	pg_QCheckBox_connect_toggled func(uintptr, int64)
	pg_QCheckBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QCheckBox_connect_windowIconChanged func(uintptr, int64)
	pg_QCheckBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QCheckBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QCheckBox_connect_destroyed func(uintptr, int64)
	pg_QCheckBox_connect_destroyed2 func(uintptr, int64)
	pg_QCheckBox_connect_objectNameChanged func(uintptr, int64)
	pg_QCheckBox__disconnect func(int64)
	pg_QCheckBox__set_override func(uintptr, int32)
	pg_QCheckBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCheckBox_chain, "QCheckBox_chain")
	symaddr(&addr_QCheckBox_new, "QCheckBox_new")
	symaddr(&addr_QCheckBox_new2, "QCheckBox_new2")
	symaddr(&addr_QCheckBox_new3, "QCheckBox_new3")
	symaddr(&addr_QCheckBox_new4, "QCheckBox_new4")
	symaddr(&addr_QCheckBox_tr, "QCheckBox_tr")
	symaddr(&addr_QCheckBox_setTristate, "QCheckBox_setTristate")
	symaddr(&addr_QCheckBox_isTristate, "QCheckBox_isTristate")
	symaddr(&addr_QCheckBox_checkState, "QCheckBox_checkState")
	symaddr(&addr_QCheckBox_setCheckState, "QCheckBox_setCheckState")
	symaddr(&addr_QCheckBox_stateChanged, "QCheckBox_stateChanged")
	symaddr(&addr_QCheckBox_checkStateChanged, "QCheckBox_checkStateChanged")
	symaddr(&addr_QCheckBox_setTabOrder, "QCheckBox_setTabOrder")
	symaddr(&addr_QCheckBox_mouseGrabber, "QCheckBox_mouseGrabber")
	symaddr(&addr_QCheckBox_keyboardGrabber, "QCheckBox_keyboardGrabber")
	symaddr(&addr_QCheckBox_find, "QCheckBox_find")
	symaddr(&addr_QCheckBox_createWindowContainer, "QCheckBox_createWindowContainer")
	symaddr(&addr_QCheckBox_updateMicroFocus, "QCheckBox_updateMicroFocus")
	symaddr(&addr_QCheckBox_create, "QCheckBox_create")
	symaddr(&addr_QCheckBox_destroy, "QCheckBox_destroy")
	symaddr(&addr_QCheckBox_focusNextChild, "QCheckBox_focusNextChild")
	symaddr(&addr_QCheckBox_focusPreviousChild, "QCheckBox_focusPreviousChild")
	symaddr(&addr_QCheckBox_connect, "QCheckBox_connect")
	symaddr(&addr_QCheckBox_connect2, "QCheckBox_connect2")
	symaddr(&addr_QCheckBox_sender, "QCheckBox_sender")
	symaddr(&addr_QCheckBox_senderSignalIndex, "QCheckBox_senderSignalIndex")
	symaddr(&addr_QCheckBox_receivers, "QCheckBox_receivers")
	symaddr(&addr_QCheckBox_isSignalConnected, "QCheckBox_isSignalConnected")
	reg(&pg_QCheckBox_devicePixelRatioFScale, "QCheckBox_devicePixelRatioFScale")
	reg(&pg_QCheckBox_encodeMetricF, "QCheckBox_encodeMetricF")
	reg(&pg_QCheckBox_getDecodedMetricF, "QCheckBox_getDecodedMetricF")
	symaddr(&addr_QCheckBox_tr2, "QCheckBox_tr2")
	symaddr(&addr_QCheckBox_tr3, "QCheckBox_tr3")
	symaddr(&addr_QCheckBox_setTristate2, "QCheckBox_setTristate2")
	symaddr(&addr_QCheckBox_createWindowContainer2, "QCheckBox_createWindowContainer2")
	symaddr(&addr_QCheckBox_createWindowContainer3, "QCheckBox_createWindowContainer3")
	symaddr(&addr_QCheckBox_updateMicroFocus2, "QCheckBox_updateMicroFocus2")
	symaddr(&addr_QCheckBox_create2, "QCheckBox_create2")
	symaddr(&addr_QCheckBox_create3, "QCheckBox_create3")
	symaddr(&addr_QCheckBox_create4, "QCheckBox_create4")
	symaddr(&addr_QCheckBox_destroy2, "QCheckBox_destroy2")
	symaddr(&addr_QCheckBox_destroy3, "QCheckBox_destroy3")
	symaddr(&addr_QCheckBox_connect3, "QCheckBox_connect3")
	symaddr(&addr_QCheckBox_connect4, "QCheckBox_connect4")
	symaddr(&addr_QCheckBox_metaObject_default, "QCheckBox_metaObject_default")
	symaddr(&addr_QCheckBox_sizeHint_default, "QCheckBox_sizeHint_default")
	symaddr(&addr_QCheckBox_minimumSizeHint_default, "QCheckBox_minimumSizeHint_default")
	symaddr(&addr_QCheckBox_event_default, "QCheckBox_event_default")
	symaddr(&addr_QCheckBox_hitButton_default, "QCheckBox_hitButton_default")
	symaddr(&addr_QCheckBox_checkStateSet_default, "QCheckBox_checkStateSet_default")
	symaddr(&addr_QCheckBox_nextCheckState_default, "QCheckBox_nextCheckState_default")
	symaddr(&addr_QCheckBox_paintEvent_default, "QCheckBox_paintEvent_default")
	symaddr(&addr_QCheckBox_mouseMoveEvent_default, "QCheckBox_mouseMoveEvent_default")
	symaddr(&addr_QCheckBox_initStyleOption_default, "QCheckBox_initStyleOption_default")
	symaddr(&addr_QCheckBox_keyPressEvent_default, "QCheckBox_keyPressEvent_default")
	symaddr(&addr_QCheckBox_keyReleaseEvent_default, "QCheckBox_keyReleaseEvent_default")
	symaddr(&addr_QCheckBox_mousePressEvent_default, "QCheckBox_mousePressEvent_default")
	symaddr(&addr_QCheckBox_mouseReleaseEvent_default, "QCheckBox_mouseReleaseEvent_default")
	symaddr(&addr_QCheckBox_focusInEvent_default, "QCheckBox_focusInEvent_default")
	symaddr(&addr_QCheckBox_focusOutEvent_default, "QCheckBox_focusOutEvent_default")
	symaddr(&addr_QCheckBox_changeEvent_default, "QCheckBox_changeEvent_default")
	symaddr(&addr_QCheckBox_timerEvent_default, "QCheckBox_timerEvent_default")
	symaddr(&addr_QCheckBox_devType_default, "QCheckBox_devType_default")
	symaddr(&addr_QCheckBox_setVisible_default, "QCheckBox_setVisible_default")
	symaddr(&addr_QCheckBox_heightForWidth_default, "QCheckBox_heightForWidth_default")
	symaddr(&addr_QCheckBox_hasHeightForWidth_default, "QCheckBox_hasHeightForWidth_default")
	symaddr(&addr_QCheckBox_paintEngine_default, "QCheckBox_paintEngine_default")
	symaddr(&addr_QCheckBox_mouseDoubleClickEvent_default, "QCheckBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QCheckBox_wheelEvent_default, "QCheckBox_wheelEvent_default")
	symaddr(&addr_QCheckBox_enterEvent_default, "QCheckBox_enterEvent_default")
	symaddr(&addr_QCheckBox_leaveEvent_default, "QCheckBox_leaveEvent_default")
	symaddr(&addr_QCheckBox_moveEvent_default, "QCheckBox_moveEvent_default")
	symaddr(&addr_QCheckBox_resizeEvent_default, "QCheckBox_resizeEvent_default")
	symaddr(&addr_QCheckBox_closeEvent_default, "QCheckBox_closeEvent_default")
	symaddr(&addr_QCheckBox_contextMenuEvent_default, "QCheckBox_contextMenuEvent_default")
	symaddr(&addr_QCheckBox_tabletEvent_default, "QCheckBox_tabletEvent_default")
	symaddr(&addr_QCheckBox_actionEvent_default, "QCheckBox_actionEvent_default")
	symaddr(&addr_QCheckBox_dragEnterEvent_default, "QCheckBox_dragEnterEvent_default")
	symaddr(&addr_QCheckBox_dragMoveEvent_default, "QCheckBox_dragMoveEvent_default")
	symaddr(&addr_QCheckBox_dragLeaveEvent_default, "QCheckBox_dragLeaveEvent_default")
	symaddr(&addr_QCheckBox_dropEvent_default, "QCheckBox_dropEvent_default")
	symaddr(&addr_QCheckBox_showEvent_default, "QCheckBox_showEvent_default")
	symaddr(&addr_QCheckBox_hideEvent_default, "QCheckBox_hideEvent_default")
	symaddr(&addr_QCheckBox_metric_default, "QCheckBox_metric_default")
	symaddr(&addr_QCheckBox_initPainter_default, "QCheckBox_initPainter_default")
	symaddr(&addr_QCheckBox_redirected_default, "QCheckBox_redirected_default")
	symaddr(&addr_QCheckBox_sharedPainter_default, "QCheckBox_sharedPainter_default")
	symaddr(&addr_QCheckBox_inputMethodEvent_default, "QCheckBox_inputMethodEvent_default")
	symaddr(&addr_QCheckBox_inputMethodQuery_default, "QCheckBox_inputMethodQuery_default")
	symaddr(&addr_QCheckBox_focusNextPrevChild_default, "QCheckBox_focusNextPrevChild_default")
	symaddr(&addr_QCheckBox_eventFilter_default, "QCheckBox_eventFilter_default")
	symaddr(&addr_QCheckBox_childEvent_default, "QCheckBox_childEvent_default")
	symaddr(&addr_QCheckBox_customEvent_default, "QCheckBox_customEvent_default")
	symaddr(&addr_QCheckBox_connectNotify_default, "QCheckBox_connectNotify_default")
	symaddr(&addr_QCheckBox_disconnectNotify_default, "QCheckBox_disconnectNotify_default")
	reg(&pg_QCheckBox_connect_stateChanged, "QCheckBox_connect_stateChanged")
	reg(&pg_QCheckBox_connect_checkStateChanged, "QCheckBox_connect_checkStateChanged")
	reg(&pg_QCheckBox_connect_pressed, "QCheckBox_connect_pressed")
	reg(&pg_QCheckBox_connect_released, "QCheckBox_connect_released")
	reg(&pg_QCheckBox_connect_clicked, "QCheckBox_connect_clicked")
	reg(&pg_QCheckBox_connect_clicked2, "QCheckBox_connect_clicked2")
	reg(&pg_QCheckBox_connect_toggled, "QCheckBox_connect_toggled")
	reg(&pg_QCheckBox_connect_windowTitleChanged, "QCheckBox_connect_windowTitleChanged")
	reg(&pg_QCheckBox_connect_windowIconChanged, "QCheckBox_connect_windowIconChanged")
	reg(&pg_QCheckBox_connect_windowIconTextChanged, "QCheckBox_connect_windowIconTextChanged")
	reg(&pg_QCheckBox_connect_customContextMenuRequested, "QCheckBox_connect_customContextMenuRequested")
	reg(&pg_QCheckBox_connect_destroyed, "QCheckBox_connect_destroyed")
	reg(&pg_QCheckBox_connect_destroyed2, "QCheckBox_connect_destroyed2")
	reg(&pg_QCheckBox_connect_objectNameChanged, "QCheckBox_connect_objectNameChanged")
	reg(&pg_QCheckBox__disconnect, "QCheckBox__disconnect")
	reg(&pg_QCheckBox__set_override, "QCheckBox__set_override")
	reg(&pg_QCheckBox_delete, "QCheckBox_delete")
}

type QCheckBox struct {
	h uintptr
	*QAbstractButton
}

func bindQCheckBox(h uintptr) *QCheckBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QCheckBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QCheckBox{h: h, QAbstractButton: &QAbstractButton{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQCheckBox(h uintptr) *QCheckBox { return bindQCheckBox(h) }

func (this *QCheckBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCheckBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCheckBox(h unsafe.Pointer) *QCheckBox { return bindQCheckBox(uintptr(h)) }
func UnsafeNewQCheckBoxOwned(h unsafe.Pointer) *QCheckBox { return newownedQCheckBox(uintptr(h)) }

func (this *QCheckBox) UnsafeRelease() {
	this.h = 0
	if this.QAbstractButton != nil {
		this.QAbstractButton.UnsafeRelease()
	}
}

func NewQCheckBox(parent *QWidget) *QCheckBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCheckBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQCheckBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCheckBox_virtIdx, QCheckBox_enable)
	return __o
}

func NewQCheckBox2() *QCheckBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCheckBox_new2, uintptr(__goid))
	__o := newownedQCheckBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCheckBox_virtIdx, QCheckBox_enable)
	return __o
}

func NewQCheckBox3(text string) *QCheckBox {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QCheckBox_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQCheckBox(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCheckBox_virtIdx, QCheckBox_enable)
	return __o
}

func NewQCheckBox4(text string, parent *QWidget) *QCheckBox {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QCheckBox_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQCheckBox(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCheckBox_virtIdx, QCheckBox_enable)
	return __o
}

func QCheckBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCheckBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QCheckBox) SetTristate() {
	purego.SyscallN(addr_QCheckBox_setTristate, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) IsTristate() bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_isTristate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCheckBox) CheckState() CheckState {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_checkState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QCheckBox) SetCheckState(state CheckState) {
	purego.SyscallN(addr_QCheckBox_setCheckState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) StateChanged(arg1 int) {
	purego.SyscallN(addr_QCheckBox_stateChanged, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) CheckStateChanged(arg1 CheckState) {
	purego.SyscallN(addr_QCheckBox_checkStateChanged, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func QCheckBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QCheckBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QCheckBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_mouseGrabber)
	return bindQWidget(__r)
}

func QCheckBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QCheckBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QCheckBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QCheckBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QCheckBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) Create() {
	purego.SyscallN(addr_QCheckBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) Destroy() {
	purego.SyscallN(addr_QCheckBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCheckBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QCheckBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCheckBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCheckBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCheckBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QCheckBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCheckBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QCheckBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QCheckBox_DevicePixelRatioFScale() float64 {
	__r := pg_QCheckBox_devicePixelRatioFScale()
	return __r
}

func QCheckBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QCheckBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QCheckBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QCheckBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QCheckBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCheckBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCheckBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCheckBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QCheckBox) SetTristateWithBool(y bool) {
	purego.SyscallN(addr_QCheckBox_setTristate2, this.h, uintptr(b2i(y)))
	runtime.KeepAlive(this)
}

func QCheckBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QCheckBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QCheckBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QCheckBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QCheckBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QCheckBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QCheckBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QCheckBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QCheckBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QCheckBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCheckBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCheckBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCheckBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QCheckBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCheckBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCheckBox) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QCheckBox) HitButtonDefault(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_hitButton_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QCheckBox) CheckStateSetDefault() {
	purego.SyscallN(addr_QCheckBox_checkStateSet_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) NextCheckStateDefault() {
	purego.SyscallN(addr_QCheckBox_nextCheckState_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QCheckBox) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QCheckBox_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCheckBox) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QCheckBox_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCheckBox) InitStyleOptionDefault(option *QStyleOptionButton) {
	purego.SyscallN(addr_QCheckBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QCheckBox) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QCheckBox_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QCheckBox_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QCheckBox_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QCheckBox_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QCheckBox_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QCheckBox_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QCheckBox_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QCheckBox_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QCheckBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCheckBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QCheckBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QCheckBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCheckBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCheckBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QCheckBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCheckBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QCheckBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QCheckBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCheckBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QCheckBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QCheckBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QCheckBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QCheckBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QCheckBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QCheckBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QCheckBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QCheckBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QCheckBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QCheckBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QCheckBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QCheckBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCheckBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QCheckBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QCheckBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QCheckBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QCheckBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QCheckBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCheckBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCheckBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCheckBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCheckBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCheckBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCheckBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QCheckBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCheckBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCheckBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCheckBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCheckBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCheckBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCheckBox) OnStateChanged(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QCheckBox_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnCheckStateChanged(fn func(CheckState)) int64 {
	id := registerCallback(func(arg1 int32) { fn(CheckState(arg1)) })
	pg_QCheckBox_connect_checkStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCheckBox_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCheckBox_connect_released(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCheckBox_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QCheckBox_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QCheckBox_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QCheckBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QCheckBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QCheckBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QCheckBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCheckBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCheckBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCheckBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCheckBox) Disconnect(handle int64) {
	pg_QCheckBox__disconnect(handle)
	unregisterCallback(handle)
}

func QCheckBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case 5736040353629178930:
		return 4
	case 5747806663224476200:
		return 5
	case 1474647914901213797:
		return 6
	case 5290442163167458465:
		return 7
	case -4255494004099806581:
		return 8
	case -488693769771308531:
		return 9
	case 1023791977679793185:
		return 10
	case 772598475950050299:
		return 11
	case -332090898078543319:
		return 12
	case -4415920049297795213:
		return 13
	case 975043767033552552:
		return 14
	case -385596010576048993:
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

func QCheckBox_enable(h uintptr, idx int32) { pg_QCheckBox__set_override(h, idx) }

func (this *QCheckBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QCheckBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QCheckBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QCheckBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QCheckBox) OnHitButton(fn func(*QPoint) bool) {
	registerVirt(this.UnsafeGoID(), 5736040353629178930, func(pos uintptr) int32 { return b2i(fn(UnsafeNewQPoint(unsafe.Pointer(pos)))) }) // slot "hitButton"
}

func (this *QCheckBox) OnCheckStateSet(fn func()) {
	registerVirt(this.UnsafeGoID(), 5747806663224476200, func() { fn() }) // slot "checkStateSet"
}

func (this *QCheckBox) OnNextCheckState(fn func()) {
	registerVirt(this.UnsafeGoID(), 1474647914901213797, func() { fn() }) // slot "nextCheckState"
}

func (this *QCheckBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QCheckBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QCheckBox) OnInitStyleOption(fn func(*QStyleOptionButton)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionButton(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QCheckBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QCheckBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QCheckBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QCheckBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QCheckBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QCheckBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QCheckBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QCheckBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QCheckBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QCheckBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QCheckBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QCheckBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QCheckBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QCheckBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QCheckBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QCheckBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QCheckBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QCheckBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QCheckBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QCheckBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QCheckBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QCheckBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QCheckBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QCheckBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QCheckBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QCheckBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QCheckBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QCheckBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QCheckBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QCheckBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QCheckBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QCheckBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QCheckBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QCheckBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QCheckBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QCheckBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QCheckBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QCheckBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QCheckBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QCheckBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QCheckBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QCheckBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QCheckBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCheckBox_delete(this.h)
	}
	this.UnsafeRelease()
}
