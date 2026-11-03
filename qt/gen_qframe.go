package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFrame_chain uintptr
	addr_QFrame_new uintptr
	addr_QFrame_new2 uintptr
	addr_QFrame_new3 uintptr
	addr_QFrame_tr uintptr
	addr_QFrame_frameStyle uintptr
	addr_QFrame_setFrameStyle uintptr
	addr_QFrame_frameWidth uintptr
	addr_QFrame_frameShape uintptr
	addr_QFrame_setFrameShape uintptr
	addr_QFrame_frameShadow uintptr
	addr_QFrame_setFrameShadow uintptr
	addr_QFrame_lineWidth uintptr
	addr_QFrame_setLineWidth uintptr
	addr_QFrame_midLineWidth uintptr
	addr_QFrame_setMidLineWidth uintptr
	addr_QFrame_frameRect uintptr
	addr_QFrame_setFrameRect uintptr
	addr_QFrame_drawFrame uintptr
	addr_QFrame_setTabOrder uintptr
	addr_QFrame_mouseGrabber uintptr
	addr_QFrame_keyboardGrabber uintptr
	addr_QFrame_find uintptr
	addr_QFrame_createWindowContainer uintptr
	addr_QFrame_updateMicroFocus uintptr
	addr_QFrame_create uintptr
	addr_QFrame_destroy uintptr
	addr_QFrame_focusNextChild uintptr
	addr_QFrame_focusPreviousChild uintptr
	addr_QFrame_connect uintptr
	addr_QFrame_connect2 uintptr
	addr_QFrame_sender uintptr
	addr_QFrame_senderSignalIndex uintptr
	addr_QFrame_receivers uintptr
	addr_QFrame_isSignalConnected uintptr
	pg_QFrame_devicePixelRatioFScale func() float64
	pg_QFrame_encodeMetricF func(int32, float64) int32
	pg_QFrame_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QFrame_tr2 uintptr
	addr_QFrame_tr3 uintptr
	addr_QFrame_createWindowContainer2 uintptr
	addr_QFrame_createWindowContainer3 uintptr
	addr_QFrame_updateMicroFocus2 uintptr
	addr_QFrame_create2 uintptr
	addr_QFrame_create3 uintptr
	addr_QFrame_create4 uintptr
	addr_QFrame_destroy2 uintptr
	addr_QFrame_destroy3 uintptr
	addr_QFrame_connect3 uintptr
	addr_QFrame_connect4 uintptr
	addr_QFrame_metaObject_default uintptr
	addr_QFrame_sizeHint_default uintptr
	addr_QFrame_event_default uintptr
	addr_QFrame_paintEvent_default uintptr
	addr_QFrame_changeEvent_default uintptr
	addr_QFrame_initStyleOption_default uintptr
	addr_QFrame_devType_default uintptr
	addr_QFrame_setVisible_default uintptr
	addr_QFrame_minimumSizeHint_default uintptr
	addr_QFrame_heightForWidth_default uintptr
	addr_QFrame_hasHeightForWidth_default uintptr
	addr_QFrame_paintEngine_default uintptr
	addr_QFrame_mousePressEvent_default uintptr
	addr_QFrame_mouseReleaseEvent_default uintptr
	addr_QFrame_mouseDoubleClickEvent_default uintptr
	addr_QFrame_mouseMoveEvent_default uintptr
	addr_QFrame_wheelEvent_default uintptr
	addr_QFrame_keyPressEvent_default uintptr
	addr_QFrame_keyReleaseEvent_default uintptr
	addr_QFrame_focusInEvent_default uintptr
	addr_QFrame_focusOutEvent_default uintptr
	addr_QFrame_enterEvent_default uintptr
	addr_QFrame_leaveEvent_default uintptr
	addr_QFrame_moveEvent_default uintptr
	addr_QFrame_resizeEvent_default uintptr
	addr_QFrame_closeEvent_default uintptr
	addr_QFrame_contextMenuEvent_default uintptr
	addr_QFrame_tabletEvent_default uintptr
	addr_QFrame_actionEvent_default uintptr
	addr_QFrame_dragEnterEvent_default uintptr
	addr_QFrame_dragMoveEvent_default uintptr
	addr_QFrame_dragLeaveEvent_default uintptr
	addr_QFrame_dropEvent_default uintptr
	addr_QFrame_showEvent_default uintptr
	addr_QFrame_hideEvent_default uintptr
	addr_QFrame_metric_default uintptr
	addr_QFrame_initPainter_default uintptr
	addr_QFrame_redirected_default uintptr
	addr_QFrame_sharedPainter_default uintptr
	addr_QFrame_inputMethodEvent_default uintptr
	addr_QFrame_inputMethodQuery_default uintptr
	addr_QFrame_focusNextPrevChild_default uintptr
	addr_QFrame_eventFilter_default uintptr
	addr_QFrame_timerEvent_default uintptr
	addr_QFrame_childEvent_default uintptr
	addr_QFrame_customEvent_default uintptr
	addr_QFrame_connectNotify_default uintptr
	addr_QFrame_disconnectNotify_default uintptr
	pg_QFrame_connect_windowTitleChanged func(uintptr, int64)
	pg_QFrame_connect_windowIconChanged func(uintptr, int64)
	pg_QFrame_connect_windowIconTextChanged func(uintptr, int64)
	pg_QFrame_connect_customContextMenuRequested func(uintptr, int64)
	pg_QFrame_connect_destroyed func(uintptr, int64)
	pg_QFrame_connect_destroyed2 func(uintptr, int64)
	pg_QFrame_connect_objectNameChanged func(uintptr, int64)
	pg_QFrame__disconnect func(int64)
	pg_QFrame__set_override func(uintptr, int32)
	pg_QFrame_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFrame_chain, "QFrame_chain")
	symaddr(&addr_QFrame_new, "QFrame_new")
	symaddr(&addr_QFrame_new2, "QFrame_new2")
	symaddr(&addr_QFrame_new3, "QFrame_new3")
	symaddr(&addr_QFrame_tr, "QFrame_tr")
	symaddr(&addr_QFrame_frameStyle, "QFrame_frameStyle")
	symaddr(&addr_QFrame_setFrameStyle, "QFrame_setFrameStyle")
	symaddr(&addr_QFrame_frameWidth, "QFrame_frameWidth")
	symaddr(&addr_QFrame_frameShape, "QFrame_frameShape")
	symaddr(&addr_QFrame_setFrameShape, "QFrame_setFrameShape")
	symaddr(&addr_QFrame_frameShadow, "QFrame_frameShadow")
	symaddr(&addr_QFrame_setFrameShadow, "QFrame_setFrameShadow")
	symaddr(&addr_QFrame_lineWidth, "QFrame_lineWidth")
	symaddr(&addr_QFrame_setLineWidth, "QFrame_setLineWidth")
	symaddr(&addr_QFrame_midLineWidth, "QFrame_midLineWidth")
	symaddr(&addr_QFrame_setMidLineWidth, "QFrame_setMidLineWidth")
	symaddr(&addr_QFrame_frameRect, "QFrame_frameRect")
	symaddr(&addr_QFrame_setFrameRect, "QFrame_setFrameRect")
	symaddr(&addr_QFrame_drawFrame, "QFrame_drawFrame")
	symaddr(&addr_QFrame_setTabOrder, "QFrame_setTabOrder")
	symaddr(&addr_QFrame_mouseGrabber, "QFrame_mouseGrabber")
	symaddr(&addr_QFrame_keyboardGrabber, "QFrame_keyboardGrabber")
	symaddr(&addr_QFrame_find, "QFrame_find")
	symaddr(&addr_QFrame_createWindowContainer, "QFrame_createWindowContainer")
	symaddr(&addr_QFrame_updateMicroFocus, "QFrame_updateMicroFocus")
	symaddr(&addr_QFrame_create, "QFrame_create")
	symaddr(&addr_QFrame_destroy, "QFrame_destroy")
	symaddr(&addr_QFrame_focusNextChild, "QFrame_focusNextChild")
	symaddr(&addr_QFrame_focusPreviousChild, "QFrame_focusPreviousChild")
	symaddr(&addr_QFrame_connect, "QFrame_connect")
	symaddr(&addr_QFrame_connect2, "QFrame_connect2")
	symaddr(&addr_QFrame_sender, "QFrame_sender")
	symaddr(&addr_QFrame_senderSignalIndex, "QFrame_senderSignalIndex")
	symaddr(&addr_QFrame_receivers, "QFrame_receivers")
	symaddr(&addr_QFrame_isSignalConnected, "QFrame_isSignalConnected")
	reg(&pg_QFrame_devicePixelRatioFScale, "QFrame_devicePixelRatioFScale")
	reg(&pg_QFrame_encodeMetricF, "QFrame_encodeMetricF")
	reg(&pg_QFrame_getDecodedMetricF, "QFrame_getDecodedMetricF")
	symaddr(&addr_QFrame_tr2, "QFrame_tr2")
	symaddr(&addr_QFrame_tr3, "QFrame_tr3")
	symaddr(&addr_QFrame_createWindowContainer2, "QFrame_createWindowContainer2")
	symaddr(&addr_QFrame_createWindowContainer3, "QFrame_createWindowContainer3")
	symaddr(&addr_QFrame_updateMicroFocus2, "QFrame_updateMicroFocus2")
	symaddr(&addr_QFrame_create2, "QFrame_create2")
	symaddr(&addr_QFrame_create3, "QFrame_create3")
	symaddr(&addr_QFrame_create4, "QFrame_create4")
	symaddr(&addr_QFrame_destroy2, "QFrame_destroy2")
	symaddr(&addr_QFrame_destroy3, "QFrame_destroy3")
	symaddr(&addr_QFrame_connect3, "QFrame_connect3")
	symaddr(&addr_QFrame_connect4, "QFrame_connect4")
	symaddr(&addr_QFrame_metaObject_default, "QFrame_metaObject_default")
	symaddr(&addr_QFrame_sizeHint_default, "QFrame_sizeHint_default")
	symaddr(&addr_QFrame_event_default, "QFrame_event_default")
	symaddr(&addr_QFrame_paintEvent_default, "QFrame_paintEvent_default")
	symaddr(&addr_QFrame_changeEvent_default, "QFrame_changeEvent_default")
	symaddr(&addr_QFrame_initStyleOption_default, "QFrame_initStyleOption_default")
	symaddr(&addr_QFrame_devType_default, "QFrame_devType_default")
	symaddr(&addr_QFrame_setVisible_default, "QFrame_setVisible_default")
	symaddr(&addr_QFrame_minimumSizeHint_default, "QFrame_minimumSizeHint_default")
	symaddr(&addr_QFrame_heightForWidth_default, "QFrame_heightForWidth_default")
	symaddr(&addr_QFrame_hasHeightForWidth_default, "QFrame_hasHeightForWidth_default")
	symaddr(&addr_QFrame_paintEngine_default, "QFrame_paintEngine_default")
	symaddr(&addr_QFrame_mousePressEvent_default, "QFrame_mousePressEvent_default")
	symaddr(&addr_QFrame_mouseReleaseEvent_default, "QFrame_mouseReleaseEvent_default")
	symaddr(&addr_QFrame_mouseDoubleClickEvent_default, "QFrame_mouseDoubleClickEvent_default")
	symaddr(&addr_QFrame_mouseMoveEvent_default, "QFrame_mouseMoveEvent_default")
	symaddr(&addr_QFrame_wheelEvent_default, "QFrame_wheelEvent_default")
	symaddr(&addr_QFrame_keyPressEvent_default, "QFrame_keyPressEvent_default")
	symaddr(&addr_QFrame_keyReleaseEvent_default, "QFrame_keyReleaseEvent_default")
	symaddr(&addr_QFrame_focusInEvent_default, "QFrame_focusInEvent_default")
	symaddr(&addr_QFrame_focusOutEvent_default, "QFrame_focusOutEvent_default")
	symaddr(&addr_QFrame_enterEvent_default, "QFrame_enterEvent_default")
	symaddr(&addr_QFrame_leaveEvent_default, "QFrame_leaveEvent_default")
	symaddr(&addr_QFrame_moveEvent_default, "QFrame_moveEvent_default")
	symaddr(&addr_QFrame_resizeEvent_default, "QFrame_resizeEvent_default")
	symaddr(&addr_QFrame_closeEvent_default, "QFrame_closeEvent_default")
	symaddr(&addr_QFrame_contextMenuEvent_default, "QFrame_contextMenuEvent_default")
	symaddr(&addr_QFrame_tabletEvent_default, "QFrame_tabletEvent_default")
	symaddr(&addr_QFrame_actionEvent_default, "QFrame_actionEvent_default")
	symaddr(&addr_QFrame_dragEnterEvent_default, "QFrame_dragEnterEvent_default")
	symaddr(&addr_QFrame_dragMoveEvent_default, "QFrame_dragMoveEvent_default")
	symaddr(&addr_QFrame_dragLeaveEvent_default, "QFrame_dragLeaveEvent_default")
	symaddr(&addr_QFrame_dropEvent_default, "QFrame_dropEvent_default")
	symaddr(&addr_QFrame_showEvent_default, "QFrame_showEvent_default")
	symaddr(&addr_QFrame_hideEvent_default, "QFrame_hideEvent_default")
	symaddr(&addr_QFrame_metric_default, "QFrame_metric_default")
	symaddr(&addr_QFrame_initPainter_default, "QFrame_initPainter_default")
	symaddr(&addr_QFrame_redirected_default, "QFrame_redirected_default")
	symaddr(&addr_QFrame_sharedPainter_default, "QFrame_sharedPainter_default")
	symaddr(&addr_QFrame_inputMethodEvent_default, "QFrame_inputMethodEvent_default")
	symaddr(&addr_QFrame_inputMethodQuery_default, "QFrame_inputMethodQuery_default")
	symaddr(&addr_QFrame_focusNextPrevChild_default, "QFrame_focusNextPrevChild_default")
	symaddr(&addr_QFrame_eventFilter_default, "QFrame_eventFilter_default")
	symaddr(&addr_QFrame_timerEvent_default, "QFrame_timerEvent_default")
	symaddr(&addr_QFrame_childEvent_default, "QFrame_childEvent_default")
	symaddr(&addr_QFrame_customEvent_default, "QFrame_customEvent_default")
	symaddr(&addr_QFrame_connectNotify_default, "QFrame_connectNotify_default")
	symaddr(&addr_QFrame_disconnectNotify_default, "QFrame_disconnectNotify_default")
	reg(&pg_QFrame_connect_windowTitleChanged, "QFrame_connect_windowTitleChanged")
	reg(&pg_QFrame_connect_windowIconChanged, "QFrame_connect_windowIconChanged")
	reg(&pg_QFrame_connect_windowIconTextChanged, "QFrame_connect_windowIconTextChanged")
	reg(&pg_QFrame_connect_customContextMenuRequested, "QFrame_connect_customContextMenuRequested")
	reg(&pg_QFrame_connect_destroyed, "QFrame_connect_destroyed")
	reg(&pg_QFrame_connect_destroyed2, "QFrame_connect_destroyed2")
	reg(&pg_QFrame_connect_objectNameChanged, "QFrame_connect_objectNameChanged")
	reg(&pg_QFrame__disconnect, "QFrame__disconnect")
	reg(&pg_QFrame__set_override, "QFrame__set_override")
	reg(&pg_QFrame_delete, "QFrame_delete")
}

type QFrame struct {
	h uintptr
	*QWidget
}

func bindQFrame(h uintptr) *QFrame {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QFrame_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFrame{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQFrame(h uintptr) *QFrame { return bindQFrame(h) }

func (this *QFrame) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFrame) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFrame(h unsafe.Pointer) *QFrame { return bindQFrame(uintptr(h)) }
func UnsafeNewQFrameOwned(h unsafe.Pointer) *QFrame { return newownedQFrame(uintptr(h)) }

func (this *QFrame) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQFrame(parent *QWidget) *QFrame {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFrame_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFrame(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFrame_virtIdx, QFrame_enable)
	return __o
}

func NewQFrame2() *QFrame {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFrame_new2, uintptr(__goid))
	__o := newownedQFrame(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFrame_virtIdx, QFrame_enable)
	return __o
}

func NewQFrame3(parent *QWidget, f WindowType) *QFrame {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFrame_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQFrame(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFrame_virtIdx, QFrame_enable)
	return __o
}

func QFrame_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFrame_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFrame) FrameStyle() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_frameStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) SetFrameStyle(arg1 int) {
	purego.SyscallN(addr_QFrame_setFrameStyle, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) FrameWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_frameWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) FrameShape() QFrame__Shape {
	__r0, _, _ := purego.SyscallN(addr_QFrame_frameShape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFrame__Shape(__r)
}

func (this *QFrame) SetFrameShape(arg1 QFrame__Shape) {
	purego.SyscallN(addr_QFrame_setFrameShape, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) FrameShadow() QFrame__Shadow {
	__r0, _, _ := purego.SyscallN(addr_QFrame_frameShadow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFrame__Shadow(__r)
}

func (this *QFrame) SetFrameShadow(arg1 QFrame__Shadow) {
	purego.SyscallN(addr_QFrame_setFrameShadow, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_lineWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) SetLineWidth(arg1 int) {
	purego.SyscallN(addr_QFrame_setLineWidth, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) MidLineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_midLineWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) SetMidLineWidth(arg1 int) {
	purego.SyscallN(addr_QFrame_setMidLineWidth, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) FrameRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QFrame_frameRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QFrame) SetFrameRect(arg1 *QRect) {
	purego.SyscallN(addr_QFrame_setFrameRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFrame) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QFrame_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QFrame_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QFrame_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFrame_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_mouseGrabber)
	return bindQWidget(__r)
}

func QFrame_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_keyboardGrabber)
	return bindQWidget(__r)
}

func QFrame_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QFrame_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QFrame) UpdateMicroFocus() {
	purego.SyscallN(addr_QFrame_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QFrame) Create() {
	purego.SyscallN(addr_QFrame_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QFrame) Destroy() {
	purego.SyscallN(addr_QFrame_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QFrame) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFrame) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFrame_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFrame_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFrame_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFrame_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFrame) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFrame_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFrame) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFrame_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFrame) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFrame_DevicePixelRatioFScale() float64 {
	__r := pg_QFrame_devicePixelRatioFScale()
	return __r
}

func QFrame_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QFrame_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QFrame) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QFrame_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QFrame_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFrame_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFrame_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFrame_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFrame_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QFrame_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFrame_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QFrame) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QFrame_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QFrame) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QFrame_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFrame) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QFrame_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QFrame) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QFrame_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QFrame) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QFrame_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QFrame) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QFrame_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QFrame_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFrame_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFrame_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFrame_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFrame) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFrame_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFrame) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFrame_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFrame) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QFrame) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QFrame_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFrame) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QFrame_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFrame) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QFrame_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QFrame) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QFrame_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QFrame) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFrame_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFrame) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFrame) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QFrame_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QFrame) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFrame_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFrame_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFrame_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFrame_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QFrame_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFrame_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFrame_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFrame_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFrame_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QFrame_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFrame_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QFrame_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QFrame_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QFrame_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QFrame_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QFrame_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QFrame_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QFrame_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QFrame_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QFrame_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QFrame_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QFrame_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QFrame_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QFrame_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFrame) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QFrame_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QFrame) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QFrame_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QFrame) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QFrame_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QFrame) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QFrame_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFrame) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFrame_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFrame) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFrame) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFrame_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFrame) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFrame_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFrame_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFrame_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFrame) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFrame_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFrame) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFrame_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFrame) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QFrame_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QFrame_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QFrame_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QFrame_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFrame_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFrame_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFrame_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFrame) Disconnect(handle int64) {
	pg_QFrame__disconnect(handle)
	unregisterCallback(handle)
}

func QFrame_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 1237236456312824991:
		return 2
	case 5290442163167458465:
		return 3
	case 4959376063826165943:
		return 4
	case -488693769771308531:
		return 5
	case -8659679893868524044:
		return 6
	case 2110182719375245575:
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
	case 427629649267311777:
		return 36
	case 6477500296980050210:
		return 37
	case 4127240633310655048:
		return 38
	case 5792771328741774991:
		return 39
	case 5438819661942843222:
		return 40
	case 7345853097340709638:
		return 41
	case 5898328517505784191:
		return 42
	case 5217623514630858297:
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

func QFrame_enable(h uintptr, idx int32) { pg_QFrame__set_override(h, idx) }

func (this *QFrame) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFrame) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFrame) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QFrame) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QFrame) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QFrame) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QFrame) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QFrame) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QFrame) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QFrame) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QFrame) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFrame) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QFrame) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QFrame) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QFrame) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QFrame) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QFrame) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QFrame) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QFrame) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QFrame) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QFrame) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QFrame) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QFrame) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QFrame) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QFrame) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QFrame) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QFrame) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QFrame) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QFrame) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QFrame) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QFrame) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QFrame) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QFrame) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QFrame) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QFrame) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QFrame) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QFrame) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QFrame) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QFrame) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QFrame) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QFrame) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QFrame) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QFrame) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QFrame) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFrame) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFrame) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFrame) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFrame) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFrame) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFrame) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFrame_delete(this.h)
	}
	this.UnsafeRelease()
}
