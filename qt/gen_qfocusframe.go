package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFocusFrame_chain uintptr
	addr_QFocusFrame_new uintptr
	addr_QFocusFrame_new2 uintptr
	addr_QFocusFrame_tr uintptr
	addr_QFocusFrame_setWidget uintptr
	addr_QFocusFrame_widget uintptr
	addr_QFocusFrame_setTabOrder uintptr
	addr_QFocusFrame_mouseGrabber uintptr
	addr_QFocusFrame_keyboardGrabber uintptr
	addr_QFocusFrame_find uintptr
	addr_QFocusFrame_createWindowContainer uintptr
	addr_QFocusFrame_updateMicroFocus uintptr
	addr_QFocusFrame_create uintptr
	addr_QFocusFrame_destroy uintptr
	addr_QFocusFrame_focusNextChild uintptr
	addr_QFocusFrame_focusPreviousChild uintptr
	addr_QFocusFrame_connect uintptr
	addr_QFocusFrame_connect2 uintptr
	addr_QFocusFrame_sender uintptr
	addr_QFocusFrame_senderSignalIndex uintptr
	addr_QFocusFrame_receivers uintptr
	addr_QFocusFrame_isSignalConnected uintptr
	pg_QFocusFrame_devicePixelRatioFScale func() float64
	pg_QFocusFrame_encodeMetricF func(int32, float64) int32
	pg_QFocusFrame_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QFocusFrame_tr2 uintptr
	addr_QFocusFrame_tr3 uintptr
	addr_QFocusFrame_createWindowContainer2 uintptr
	addr_QFocusFrame_createWindowContainer3 uintptr
	addr_QFocusFrame_updateMicroFocus2 uintptr
	addr_QFocusFrame_create2 uintptr
	addr_QFocusFrame_create3 uintptr
	addr_QFocusFrame_create4 uintptr
	addr_QFocusFrame_destroy2 uintptr
	addr_QFocusFrame_destroy3 uintptr
	addr_QFocusFrame_connect3 uintptr
	addr_QFocusFrame_connect4 uintptr
	addr_QFocusFrame_metaObject_default uintptr
	addr_QFocusFrame_event_default uintptr
	addr_QFocusFrame_eventFilter_default uintptr
	addr_QFocusFrame_paintEvent_default uintptr
	addr_QFocusFrame_initStyleOption_default uintptr
	addr_QFocusFrame_devType_default uintptr
	addr_QFocusFrame_setVisible_default uintptr
	addr_QFocusFrame_sizeHint_default uintptr
	addr_QFocusFrame_minimumSizeHint_default uintptr
	addr_QFocusFrame_heightForWidth_default uintptr
	addr_QFocusFrame_hasHeightForWidth_default uintptr
	addr_QFocusFrame_paintEngine_default uintptr
	addr_QFocusFrame_mousePressEvent_default uintptr
	addr_QFocusFrame_mouseReleaseEvent_default uintptr
	addr_QFocusFrame_mouseDoubleClickEvent_default uintptr
	addr_QFocusFrame_mouseMoveEvent_default uintptr
	addr_QFocusFrame_wheelEvent_default uintptr
	addr_QFocusFrame_keyPressEvent_default uintptr
	addr_QFocusFrame_keyReleaseEvent_default uintptr
	addr_QFocusFrame_focusInEvent_default uintptr
	addr_QFocusFrame_focusOutEvent_default uintptr
	addr_QFocusFrame_enterEvent_default uintptr
	addr_QFocusFrame_leaveEvent_default uintptr
	addr_QFocusFrame_moveEvent_default uintptr
	addr_QFocusFrame_resizeEvent_default uintptr
	addr_QFocusFrame_closeEvent_default uintptr
	addr_QFocusFrame_contextMenuEvent_default uintptr
	addr_QFocusFrame_tabletEvent_default uintptr
	addr_QFocusFrame_actionEvent_default uintptr
	addr_QFocusFrame_dragEnterEvent_default uintptr
	addr_QFocusFrame_dragMoveEvent_default uintptr
	addr_QFocusFrame_dragLeaveEvent_default uintptr
	addr_QFocusFrame_dropEvent_default uintptr
	addr_QFocusFrame_showEvent_default uintptr
	addr_QFocusFrame_hideEvent_default uintptr
	addr_QFocusFrame_changeEvent_default uintptr
	addr_QFocusFrame_metric_default uintptr
	addr_QFocusFrame_initPainter_default uintptr
	addr_QFocusFrame_redirected_default uintptr
	addr_QFocusFrame_sharedPainter_default uintptr
	addr_QFocusFrame_inputMethodEvent_default uintptr
	addr_QFocusFrame_inputMethodQuery_default uintptr
	addr_QFocusFrame_focusNextPrevChild_default uintptr
	addr_QFocusFrame_timerEvent_default uintptr
	addr_QFocusFrame_childEvent_default uintptr
	addr_QFocusFrame_customEvent_default uintptr
	addr_QFocusFrame_connectNotify_default uintptr
	addr_QFocusFrame_disconnectNotify_default uintptr
	pg_QFocusFrame_connect_windowTitleChanged func(uintptr, int64)
	pg_QFocusFrame_connect_windowIconChanged func(uintptr, int64)
	pg_QFocusFrame_connect_windowIconTextChanged func(uintptr, int64)
	pg_QFocusFrame_connect_customContextMenuRequested func(uintptr, int64)
	pg_QFocusFrame_connect_destroyed func(uintptr, int64)
	pg_QFocusFrame_connect_destroyed2 func(uintptr, int64)
	pg_QFocusFrame_connect_objectNameChanged func(uintptr, int64)
	pg_QFocusFrame__disconnect func(int64)
	pg_QFocusFrame__set_override func(uintptr, int32)
	pg_QFocusFrame_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFocusFrame_chain, "QFocusFrame_chain")
	symaddr(&addr_QFocusFrame_new, "QFocusFrame_new")
	symaddr(&addr_QFocusFrame_new2, "QFocusFrame_new2")
	symaddr(&addr_QFocusFrame_tr, "QFocusFrame_tr")
	symaddr(&addr_QFocusFrame_setWidget, "QFocusFrame_setWidget")
	symaddr(&addr_QFocusFrame_widget, "QFocusFrame_widget")
	symaddr(&addr_QFocusFrame_setTabOrder, "QFocusFrame_setTabOrder")
	symaddr(&addr_QFocusFrame_mouseGrabber, "QFocusFrame_mouseGrabber")
	symaddr(&addr_QFocusFrame_keyboardGrabber, "QFocusFrame_keyboardGrabber")
	symaddr(&addr_QFocusFrame_find, "QFocusFrame_find")
	symaddr(&addr_QFocusFrame_createWindowContainer, "QFocusFrame_createWindowContainer")
	symaddr(&addr_QFocusFrame_updateMicroFocus, "QFocusFrame_updateMicroFocus")
	symaddr(&addr_QFocusFrame_create, "QFocusFrame_create")
	symaddr(&addr_QFocusFrame_destroy, "QFocusFrame_destroy")
	symaddr(&addr_QFocusFrame_focusNextChild, "QFocusFrame_focusNextChild")
	symaddr(&addr_QFocusFrame_focusPreviousChild, "QFocusFrame_focusPreviousChild")
	symaddr(&addr_QFocusFrame_connect, "QFocusFrame_connect")
	symaddr(&addr_QFocusFrame_connect2, "QFocusFrame_connect2")
	symaddr(&addr_QFocusFrame_sender, "QFocusFrame_sender")
	symaddr(&addr_QFocusFrame_senderSignalIndex, "QFocusFrame_senderSignalIndex")
	symaddr(&addr_QFocusFrame_receivers, "QFocusFrame_receivers")
	symaddr(&addr_QFocusFrame_isSignalConnected, "QFocusFrame_isSignalConnected")
	reg(&pg_QFocusFrame_devicePixelRatioFScale, "QFocusFrame_devicePixelRatioFScale")
	reg(&pg_QFocusFrame_encodeMetricF, "QFocusFrame_encodeMetricF")
	reg(&pg_QFocusFrame_getDecodedMetricF, "QFocusFrame_getDecodedMetricF")
	symaddr(&addr_QFocusFrame_tr2, "QFocusFrame_tr2")
	symaddr(&addr_QFocusFrame_tr3, "QFocusFrame_tr3")
	symaddr(&addr_QFocusFrame_createWindowContainer2, "QFocusFrame_createWindowContainer2")
	symaddr(&addr_QFocusFrame_createWindowContainer3, "QFocusFrame_createWindowContainer3")
	symaddr(&addr_QFocusFrame_updateMicroFocus2, "QFocusFrame_updateMicroFocus2")
	symaddr(&addr_QFocusFrame_create2, "QFocusFrame_create2")
	symaddr(&addr_QFocusFrame_create3, "QFocusFrame_create3")
	symaddr(&addr_QFocusFrame_create4, "QFocusFrame_create4")
	symaddr(&addr_QFocusFrame_destroy2, "QFocusFrame_destroy2")
	symaddr(&addr_QFocusFrame_destroy3, "QFocusFrame_destroy3")
	symaddr(&addr_QFocusFrame_connect3, "QFocusFrame_connect3")
	symaddr(&addr_QFocusFrame_connect4, "QFocusFrame_connect4")
	symaddr(&addr_QFocusFrame_metaObject_default, "QFocusFrame_metaObject_default")
	symaddr(&addr_QFocusFrame_event_default, "QFocusFrame_event_default")
	symaddr(&addr_QFocusFrame_eventFilter_default, "QFocusFrame_eventFilter_default")
	symaddr(&addr_QFocusFrame_paintEvent_default, "QFocusFrame_paintEvent_default")
	symaddr(&addr_QFocusFrame_initStyleOption_default, "QFocusFrame_initStyleOption_default")
	symaddr(&addr_QFocusFrame_devType_default, "QFocusFrame_devType_default")
	symaddr(&addr_QFocusFrame_setVisible_default, "QFocusFrame_setVisible_default")
	symaddr(&addr_QFocusFrame_sizeHint_default, "QFocusFrame_sizeHint_default")
	symaddr(&addr_QFocusFrame_minimumSizeHint_default, "QFocusFrame_minimumSizeHint_default")
	symaddr(&addr_QFocusFrame_heightForWidth_default, "QFocusFrame_heightForWidth_default")
	symaddr(&addr_QFocusFrame_hasHeightForWidth_default, "QFocusFrame_hasHeightForWidth_default")
	symaddr(&addr_QFocusFrame_paintEngine_default, "QFocusFrame_paintEngine_default")
	symaddr(&addr_QFocusFrame_mousePressEvent_default, "QFocusFrame_mousePressEvent_default")
	symaddr(&addr_QFocusFrame_mouseReleaseEvent_default, "QFocusFrame_mouseReleaseEvent_default")
	symaddr(&addr_QFocusFrame_mouseDoubleClickEvent_default, "QFocusFrame_mouseDoubleClickEvent_default")
	symaddr(&addr_QFocusFrame_mouseMoveEvent_default, "QFocusFrame_mouseMoveEvent_default")
	symaddr(&addr_QFocusFrame_wheelEvent_default, "QFocusFrame_wheelEvent_default")
	symaddr(&addr_QFocusFrame_keyPressEvent_default, "QFocusFrame_keyPressEvent_default")
	symaddr(&addr_QFocusFrame_keyReleaseEvent_default, "QFocusFrame_keyReleaseEvent_default")
	symaddr(&addr_QFocusFrame_focusInEvent_default, "QFocusFrame_focusInEvent_default")
	symaddr(&addr_QFocusFrame_focusOutEvent_default, "QFocusFrame_focusOutEvent_default")
	symaddr(&addr_QFocusFrame_enterEvent_default, "QFocusFrame_enterEvent_default")
	symaddr(&addr_QFocusFrame_leaveEvent_default, "QFocusFrame_leaveEvent_default")
	symaddr(&addr_QFocusFrame_moveEvent_default, "QFocusFrame_moveEvent_default")
	symaddr(&addr_QFocusFrame_resizeEvent_default, "QFocusFrame_resizeEvent_default")
	symaddr(&addr_QFocusFrame_closeEvent_default, "QFocusFrame_closeEvent_default")
	symaddr(&addr_QFocusFrame_contextMenuEvent_default, "QFocusFrame_contextMenuEvent_default")
	symaddr(&addr_QFocusFrame_tabletEvent_default, "QFocusFrame_tabletEvent_default")
	symaddr(&addr_QFocusFrame_actionEvent_default, "QFocusFrame_actionEvent_default")
	symaddr(&addr_QFocusFrame_dragEnterEvent_default, "QFocusFrame_dragEnterEvent_default")
	symaddr(&addr_QFocusFrame_dragMoveEvent_default, "QFocusFrame_dragMoveEvent_default")
	symaddr(&addr_QFocusFrame_dragLeaveEvent_default, "QFocusFrame_dragLeaveEvent_default")
	symaddr(&addr_QFocusFrame_dropEvent_default, "QFocusFrame_dropEvent_default")
	symaddr(&addr_QFocusFrame_showEvent_default, "QFocusFrame_showEvent_default")
	symaddr(&addr_QFocusFrame_hideEvent_default, "QFocusFrame_hideEvent_default")
	symaddr(&addr_QFocusFrame_changeEvent_default, "QFocusFrame_changeEvent_default")
	symaddr(&addr_QFocusFrame_metric_default, "QFocusFrame_metric_default")
	symaddr(&addr_QFocusFrame_initPainter_default, "QFocusFrame_initPainter_default")
	symaddr(&addr_QFocusFrame_redirected_default, "QFocusFrame_redirected_default")
	symaddr(&addr_QFocusFrame_sharedPainter_default, "QFocusFrame_sharedPainter_default")
	symaddr(&addr_QFocusFrame_inputMethodEvent_default, "QFocusFrame_inputMethodEvent_default")
	symaddr(&addr_QFocusFrame_inputMethodQuery_default, "QFocusFrame_inputMethodQuery_default")
	symaddr(&addr_QFocusFrame_focusNextPrevChild_default, "QFocusFrame_focusNextPrevChild_default")
	symaddr(&addr_QFocusFrame_timerEvent_default, "QFocusFrame_timerEvent_default")
	symaddr(&addr_QFocusFrame_childEvent_default, "QFocusFrame_childEvent_default")
	symaddr(&addr_QFocusFrame_customEvent_default, "QFocusFrame_customEvent_default")
	symaddr(&addr_QFocusFrame_connectNotify_default, "QFocusFrame_connectNotify_default")
	symaddr(&addr_QFocusFrame_disconnectNotify_default, "QFocusFrame_disconnectNotify_default")
	reg(&pg_QFocusFrame_connect_windowTitleChanged, "QFocusFrame_connect_windowTitleChanged")
	reg(&pg_QFocusFrame_connect_windowIconChanged, "QFocusFrame_connect_windowIconChanged")
	reg(&pg_QFocusFrame_connect_windowIconTextChanged, "QFocusFrame_connect_windowIconTextChanged")
	reg(&pg_QFocusFrame_connect_customContextMenuRequested, "QFocusFrame_connect_customContextMenuRequested")
	reg(&pg_QFocusFrame_connect_destroyed, "QFocusFrame_connect_destroyed")
	reg(&pg_QFocusFrame_connect_destroyed2, "QFocusFrame_connect_destroyed2")
	reg(&pg_QFocusFrame_connect_objectNameChanged, "QFocusFrame_connect_objectNameChanged")
	reg(&pg_QFocusFrame__disconnect, "QFocusFrame__disconnect")
	reg(&pg_QFocusFrame__set_override, "QFocusFrame__set_override")
	reg(&pg_QFocusFrame_delete, "QFocusFrame_delete")
}

type QFocusFrame struct {
	h uintptr
	*QWidget
}

func bindQFocusFrame(h uintptr) *QFocusFrame {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QFocusFrame_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFocusFrame{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQFocusFrame(h uintptr) *QFocusFrame { return bindQFocusFrame(h) }

func (this *QFocusFrame) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFocusFrame) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFocusFrame(h unsafe.Pointer) *QFocusFrame { return bindQFocusFrame(uintptr(h)) }
func UnsafeNewQFocusFrameOwned(h unsafe.Pointer) *QFocusFrame { return newownedQFocusFrame(uintptr(h)) }

func (this *QFocusFrame) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQFocusFrame(parent *QWidget) *QFocusFrame {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFocusFrame(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFocusFrame_virtIdx, QFocusFrame_enable)
	return __o
}

func NewQFocusFrame2() *QFocusFrame {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_new2, uintptr(__goid))
	__o := newownedQFocusFrame(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFocusFrame_virtIdx, QFocusFrame_enable)
	return __o
}

func QFocusFrame_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFocusFrame) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QFocusFrame_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFocusFrame) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func QFocusFrame_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QFocusFrame_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFocusFrame_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_mouseGrabber)
	return bindQWidget(__r)
}

func QFocusFrame_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_keyboardGrabber)
	return bindQWidget(__r)
}

func QFocusFrame_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QFocusFrame_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QFocusFrame) UpdateMicroFocus() {
	purego.SyscallN(addr_QFocusFrame_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) Create() {
	purego.SyscallN(addr_QFocusFrame_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) Destroy() {
	purego.SyscallN(addr_QFocusFrame_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFocusFrame) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFocusFrame_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFocusFrame_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFocusFrame) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFocusFrame) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFocusFrame) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFocusFrame) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFocusFrame_DevicePixelRatioFScale() float64 {
	__r := pg_QFocusFrame_devicePixelRatioFScale()
	return __r
}

func QFocusFrame_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QFocusFrame_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QFocusFrame) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QFocusFrame_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QFocusFrame_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFocusFrame_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFocusFrame_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QFocusFrame_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QFocusFrame) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QFocusFrame_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QFocusFrame_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QFocusFrame_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QFocusFrame_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QFocusFrame_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QFocusFrame_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QFocusFrame_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFocusFrame_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFocusFrame) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFocusFrame) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QFocusFrame) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QFocusFrame) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QFocusFrame_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFocusFrame) InitStyleOptionDefault(option *QStyleOption) {
	purego.SyscallN(addr_QFocusFrame_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QFocusFrame) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFocusFrame) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QFocusFrame_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QFocusFrame) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFocusFrame) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFocusFrame) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFocusFrame) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFocusFrame) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QFocusFrame) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFocusFrame_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFocusFrame_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFocusFrame_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFocusFrame_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QFocusFrame_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFocusFrame_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFocusFrame_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFocusFrame_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFocusFrame_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QFocusFrame_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFocusFrame_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QFocusFrame_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QFocusFrame_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QFocusFrame_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QFocusFrame_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QFocusFrame_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QFocusFrame_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QFocusFrame_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QFocusFrame_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QFocusFrame_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QFocusFrame_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QFocusFrame_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QFocusFrame_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QFocusFrame_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFocusFrame) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFocusFrame) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QFocusFrame_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QFocusFrame) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QFocusFrame) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QFocusFrame) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QFocusFrame_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFocusFrame) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFocusFrame_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFocusFrame) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QFocusFrame_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFocusFrame) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFocusFrame_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFocusFrame_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFocusFrame_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFocusFrame) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFocusFrame_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFocusFrame) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFocusFrame_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFocusFrame) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QFocusFrame_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QFocusFrame_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QFocusFrame_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QFocusFrame_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFocusFrame_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFocusFrame_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFocusFrame_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFocusFrame) Disconnect(handle int64) {
	pg_QFocusFrame__disconnect(handle)
	unregisterCallback(handle)
}

func QFocusFrame_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case 5290442163167458465:
		return 3
	case -488693769771308531:
		return 4
	case -8659679893868524044:
		return 5
	case 2110182719375245575:
		return 6
	case 729690261457684101:
		return 7
	case 7049264599837532491:
		return 8
	case 4311534060495737339:
		return 9
	case -8266762397995423915:
		return 10
	case 6398007404756346381:
		return 11
	case -332090898078543319:
		return 12
	case -4415920049297795213:
		return 13
	case -4029073905697023887:
		return 14
	case -4255494004099806581:
		return 15
	case 1919988340934511316:
		return 16
	case 1023791977679793185:
		return 17
	case 772598475950050299:
		return 18
	case 975043767033552552:
		return 19
	case -385596010576048993:
		return 20
	case 6316270346668237895:
		return 21
	case -551070009623711640:
		return 22
	case -9161949150489525396:
		return 23
	case -4302917285655014661:
		return 24
	case -1989937050019464783:
		return 25
	case -6698043222227006015:
		return 26
	case -212068866860701141:
		return 27
	case 1697444140088555581:
		return 28
	case 206122715919917589:
		return 29
	case -1781545137352745978:
		return 30
	case -3139904158373725006:
		return 31
	case 9107058682754591716:
		return 32
	case -6825369174751916:
		return 33
	case -199718707876719129:
		return 34
	case -2039166586246443706:
		return 35
	case 4959376063826165943:
		return 36
	case 427629649267311777:
		return 37
	case 6477500296980050210:
		return 38
	case 4127240633310655048:
		return 39
	case 5792771328741774991:
		return 40
	case 5438819661942843222:
		return 41
	case 7345853097340709638:
		return 42
	case 5898328517505784191:
		return 43
	case -5409927581812374498:
		return 44
	case -1185315367526857833:
		return 45
	case 3769868229123868474:
		return 46
	case 7379029982170503008:
		return 47
	case -568369772982215962:
		return 48
	}
	return -1
}

func QFocusFrame_enable(h uintptr, idx int32) { pg_QFocusFrame__set_override(h, idx) }

func (this *QFocusFrame) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFocusFrame) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QFocusFrame) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QFocusFrame) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QFocusFrame) OnInitStyleOption(fn func(*QStyleOption)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOption(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QFocusFrame) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QFocusFrame) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QFocusFrame) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFocusFrame) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QFocusFrame) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QFocusFrame) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFocusFrame) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QFocusFrame) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QFocusFrame) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QFocusFrame) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QFocusFrame) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QFocusFrame) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QFocusFrame) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QFocusFrame) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QFocusFrame) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QFocusFrame) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QFocusFrame) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QFocusFrame) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QFocusFrame) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QFocusFrame) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QFocusFrame) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QFocusFrame) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QFocusFrame) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QFocusFrame) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QFocusFrame) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QFocusFrame) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QFocusFrame) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QFocusFrame) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QFocusFrame) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QFocusFrame) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QFocusFrame) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QFocusFrame) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QFocusFrame) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QFocusFrame) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QFocusFrame) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QFocusFrame) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QFocusFrame) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QFocusFrame) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QFocusFrame) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QFocusFrame) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFocusFrame) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFocusFrame) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFocusFrame) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFocusFrame) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFocusFrame) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFocusFrame_delete(this.h)
	}
	this.UnsafeRelease()
}
