package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDial_chain uintptr
	addr_QDial_new uintptr
	addr_QDial_new2 uintptr
	addr_QDial_tr uintptr
	addr_QDial_wrapping uintptr
	addr_QDial_notchSize uintptr
	pg_QDial_setNotchTarget func(uintptr, float64)
	pg_QDial_notchTarget func(uintptr) float64
	addr_QDial_notchesVisible uintptr
	addr_QDial_setNotchesVisible uintptr
	addr_QDial_setWrapping uintptr
	addr_QDial_setRepeatAction uintptr
	addr_QDial_repeatAction uintptr
	addr_QDial_setTabOrder uintptr
	addr_QDial_mouseGrabber uintptr
	addr_QDial_keyboardGrabber uintptr
	addr_QDial_find uintptr
	addr_QDial_createWindowContainer uintptr
	addr_QDial_updateMicroFocus uintptr
	addr_QDial_create uintptr
	addr_QDial_destroy uintptr
	addr_QDial_focusNextChild uintptr
	addr_QDial_focusPreviousChild uintptr
	addr_QDial_connect uintptr
	addr_QDial_connect2 uintptr
	addr_QDial_sender uintptr
	addr_QDial_senderSignalIndex uintptr
	addr_QDial_receivers uintptr
	addr_QDial_isSignalConnected uintptr
	pg_QDial_devicePixelRatioFScale func() float64
	pg_QDial_encodeMetricF func(int32, float64) int32
	pg_QDial_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDial_tr2 uintptr
	addr_QDial_tr3 uintptr
	addr_QDial_setRepeatAction2 uintptr
	addr_QDial_setRepeatAction3 uintptr
	addr_QDial_createWindowContainer2 uintptr
	addr_QDial_createWindowContainer3 uintptr
	addr_QDial_updateMicroFocus2 uintptr
	addr_QDial_create2 uintptr
	addr_QDial_create3 uintptr
	addr_QDial_create4 uintptr
	addr_QDial_destroy2 uintptr
	addr_QDial_destroy3 uintptr
	addr_QDial_connect3 uintptr
	addr_QDial_connect4 uintptr
	addr_QDial_metaObject_default uintptr
	addr_QDial_sizeHint_default uintptr
	addr_QDial_minimumSizeHint_default uintptr
	addr_QDial_event_default uintptr
	addr_QDial_resizeEvent_default uintptr
	addr_QDial_paintEvent_default uintptr
	addr_QDial_mousePressEvent_default uintptr
	addr_QDial_mouseReleaseEvent_default uintptr
	addr_QDial_mouseMoveEvent_default uintptr
	addr_QDial_sliderChange_default uintptr
	addr_QDial_initStyleOption_default uintptr
	addr_QDial_keyPressEvent_default uintptr
	addr_QDial_timerEvent_default uintptr
	addr_QDial_wheelEvent_default uintptr
	addr_QDial_changeEvent_default uintptr
	addr_QDial_devType_default uintptr
	addr_QDial_setVisible_default uintptr
	addr_QDial_heightForWidth_default uintptr
	addr_QDial_hasHeightForWidth_default uintptr
	addr_QDial_paintEngine_default uintptr
	addr_QDial_mouseDoubleClickEvent_default uintptr
	addr_QDial_keyReleaseEvent_default uintptr
	addr_QDial_focusInEvent_default uintptr
	addr_QDial_focusOutEvent_default uintptr
	addr_QDial_enterEvent_default uintptr
	addr_QDial_leaveEvent_default uintptr
	addr_QDial_moveEvent_default uintptr
	addr_QDial_closeEvent_default uintptr
	addr_QDial_contextMenuEvent_default uintptr
	addr_QDial_tabletEvent_default uintptr
	addr_QDial_actionEvent_default uintptr
	addr_QDial_dragEnterEvent_default uintptr
	addr_QDial_dragMoveEvent_default uintptr
	addr_QDial_dragLeaveEvent_default uintptr
	addr_QDial_dropEvent_default uintptr
	addr_QDial_showEvent_default uintptr
	addr_QDial_hideEvent_default uintptr
	addr_QDial_metric_default uintptr
	addr_QDial_initPainter_default uintptr
	addr_QDial_redirected_default uintptr
	addr_QDial_sharedPainter_default uintptr
	addr_QDial_inputMethodEvent_default uintptr
	addr_QDial_inputMethodQuery_default uintptr
	addr_QDial_focusNextPrevChild_default uintptr
	addr_QDial_eventFilter_default uintptr
	addr_QDial_childEvent_default uintptr
	addr_QDial_customEvent_default uintptr
	addr_QDial_connectNotify_default uintptr
	addr_QDial_disconnectNotify_default uintptr
	pg_QDial_connect_valueChanged func(uintptr, int64)
	pg_QDial_connect_sliderPressed func(uintptr, int64)
	pg_QDial_connect_sliderMoved func(uintptr, int64)
	pg_QDial_connect_sliderReleased func(uintptr, int64)
	pg_QDial_connect_rangeChanged func(uintptr, int64)
	pg_QDial_connect_actionTriggered func(uintptr, int64)
	pg_QDial_connect_windowTitleChanged func(uintptr, int64)
	pg_QDial_connect_windowIconChanged func(uintptr, int64)
	pg_QDial_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDial_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDial_connect_destroyed func(uintptr, int64)
	pg_QDial_connect_destroyed2 func(uintptr, int64)
	pg_QDial_connect_objectNameChanged func(uintptr, int64)
	pg_QDial__disconnect func(int64)
	pg_QDial__set_override func(uintptr, int32)
	pg_QDial_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDial_chain, "QDial_chain")
	symaddr(&addr_QDial_new, "QDial_new")
	symaddr(&addr_QDial_new2, "QDial_new2")
	symaddr(&addr_QDial_tr, "QDial_tr")
	symaddr(&addr_QDial_wrapping, "QDial_wrapping")
	symaddr(&addr_QDial_notchSize, "QDial_notchSize")
	reg(&pg_QDial_setNotchTarget, "QDial_setNotchTarget")
	reg(&pg_QDial_notchTarget, "QDial_notchTarget")
	symaddr(&addr_QDial_notchesVisible, "QDial_notchesVisible")
	symaddr(&addr_QDial_setNotchesVisible, "QDial_setNotchesVisible")
	symaddr(&addr_QDial_setWrapping, "QDial_setWrapping")
	symaddr(&addr_QDial_setRepeatAction, "QDial_setRepeatAction")
	symaddr(&addr_QDial_repeatAction, "QDial_repeatAction")
	symaddr(&addr_QDial_setTabOrder, "QDial_setTabOrder")
	symaddr(&addr_QDial_mouseGrabber, "QDial_mouseGrabber")
	symaddr(&addr_QDial_keyboardGrabber, "QDial_keyboardGrabber")
	symaddr(&addr_QDial_find, "QDial_find")
	symaddr(&addr_QDial_createWindowContainer, "QDial_createWindowContainer")
	symaddr(&addr_QDial_updateMicroFocus, "QDial_updateMicroFocus")
	symaddr(&addr_QDial_create, "QDial_create")
	symaddr(&addr_QDial_destroy, "QDial_destroy")
	symaddr(&addr_QDial_focusNextChild, "QDial_focusNextChild")
	symaddr(&addr_QDial_focusPreviousChild, "QDial_focusPreviousChild")
	symaddr(&addr_QDial_connect, "QDial_connect")
	symaddr(&addr_QDial_connect2, "QDial_connect2")
	symaddr(&addr_QDial_sender, "QDial_sender")
	symaddr(&addr_QDial_senderSignalIndex, "QDial_senderSignalIndex")
	symaddr(&addr_QDial_receivers, "QDial_receivers")
	symaddr(&addr_QDial_isSignalConnected, "QDial_isSignalConnected")
	reg(&pg_QDial_devicePixelRatioFScale, "QDial_devicePixelRatioFScale")
	reg(&pg_QDial_encodeMetricF, "QDial_encodeMetricF")
	reg(&pg_QDial_getDecodedMetricF, "QDial_getDecodedMetricF")
	symaddr(&addr_QDial_tr2, "QDial_tr2")
	symaddr(&addr_QDial_tr3, "QDial_tr3")
	symaddr(&addr_QDial_setRepeatAction2, "QDial_setRepeatAction2")
	symaddr(&addr_QDial_setRepeatAction3, "QDial_setRepeatAction3")
	symaddr(&addr_QDial_createWindowContainer2, "QDial_createWindowContainer2")
	symaddr(&addr_QDial_createWindowContainer3, "QDial_createWindowContainer3")
	symaddr(&addr_QDial_updateMicroFocus2, "QDial_updateMicroFocus2")
	symaddr(&addr_QDial_create2, "QDial_create2")
	symaddr(&addr_QDial_create3, "QDial_create3")
	symaddr(&addr_QDial_create4, "QDial_create4")
	symaddr(&addr_QDial_destroy2, "QDial_destroy2")
	symaddr(&addr_QDial_destroy3, "QDial_destroy3")
	symaddr(&addr_QDial_connect3, "QDial_connect3")
	symaddr(&addr_QDial_connect4, "QDial_connect4")
	symaddr(&addr_QDial_metaObject_default, "QDial_metaObject_default")
	symaddr(&addr_QDial_sizeHint_default, "QDial_sizeHint_default")
	symaddr(&addr_QDial_minimumSizeHint_default, "QDial_minimumSizeHint_default")
	symaddr(&addr_QDial_event_default, "QDial_event_default")
	symaddr(&addr_QDial_resizeEvent_default, "QDial_resizeEvent_default")
	symaddr(&addr_QDial_paintEvent_default, "QDial_paintEvent_default")
	symaddr(&addr_QDial_mousePressEvent_default, "QDial_mousePressEvent_default")
	symaddr(&addr_QDial_mouseReleaseEvent_default, "QDial_mouseReleaseEvent_default")
	symaddr(&addr_QDial_mouseMoveEvent_default, "QDial_mouseMoveEvent_default")
	symaddr(&addr_QDial_sliderChange_default, "QDial_sliderChange_default")
	symaddr(&addr_QDial_initStyleOption_default, "QDial_initStyleOption_default")
	symaddr(&addr_QDial_keyPressEvent_default, "QDial_keyPressEvent_default")
	symaddr(&addr_QDial_timerEvent_default, "QDial_timerEvent_default")
	symaddr(&addr_QDial_wheelEvent_default, "QDial_wheelEvent_default")
	symaddr(&addr_QDial_changeEvent_default, "QDial_changeEvent_default")
	symaddr(&addr_QDial_devType_default, "QDial_devType_default")
	symaddr(&addr_QDial_setVisible_default, "QDial_setVisible_default")
	symaddr(&addr_QDial_heightForWidth_default, "QDial_heightForWidth_default")
	symaddr(&addr_QDial_hasHeightForWidth_default, "QDial_hasHeightForWidth_default")
	symaddr(&addr_QDial_paintEngine_default, "QDial_paintEngine_default")
	symaddr(&addr_QDial_mouseDoubleClickEvent_default, "QDial_mouseDoubleClickEvent_default")
	symaddr(&addr_QDial_keyReleaseEvent_default, "QDial_keyReleaseEvent_default")
	symaddr(&addr_QDial_focusInEvent_default, "QDial_focusInEvent_default")
	symaddr(&addr_QDial_focusOutEvent_default, "QDial_focusOutEvent_default")
	symaddr(&addr_QDial_enterEvent_default, "QDial_enterEvent_default")
	symaddr(&addr_QDial_leaveEvent_default, "QDial_leaveEvent_default")
	symaddr(&addr_QDial_moveEvent_default, "QDial_moveEvent_default")
	symaddr(&addr_QDial_closeEvent_default, "QDial_closeEvent_default")
	symaddr(&addr_QDial_contextMenuEvent_default, "QDial_contextMenuEvent_default")
	symaddr(&addr_QDial_tabletEvent_default, "QDial_tabletEvent_default")
	symaddr(&addr_QDial_actionEvent_default, "QDial_actionEvent_default")
	symaddr(&addr_QDial_dragEnterEvent_default, "QDial_dragEnterEvent_default")
	symaddr(&addr_QDial_dragMoveEvent_default, "QDial_dragMoveEvent_default")
	symaddr(&addr_QDial_dragLeaveEvent_default, "QDial_dragLeaveEvent_default")
	symaddr(&addr_QDial_dropEvent_default, "QDial_dropEvent_default")
	symaddr(&addr_QDial_showEvent_default, "QDial_showEvent_default")
	symaddr(&addr_QDial_hideEvent_default, "QDial_hideEvent_default")
	symaddr(&addr_QDial_metric_default, "QDial_metric_default")
	symaddr(&addr_QDial_initPainter_default, "QDial_initPainter_default")
	symaddr(&addr_QDial_redirected_default, "QDial_redirected_default")
	symaddr(&addr_QDial_sharedPainter_default, "QDial_sharedPainter_default")
	symaddr(&addr_QDial_inputMethodEvent_default, "QDial_inputMethodEvent_default")
	symaddr(&addr_QDial_inputMethodQuery_default, "QDial_inputMethodQuery_default")
	symaddr(&addr_QDial_focusNextPrevChild_default, "QDial_focusNextPrevChild_default")
	symaddr(&addr_QDial_eventFilter_default, "QDial_eventFilter_default")
	symaddr(&addr_QDial_childEvent_default, "QDial_childEvent_default")
	symaddr(&addr_QDial_customEvent_default, "QDial_customEvent_default")
	symaddr(&addr_QDial_connectNotify_default, "QDial_connectNotify_default")
	symaddr(&addr_QDial_disconnectNotify_default, "QDial_disconnectNotify_default")
	reg(&pg_QDial_connect_valueChanged, "QDial_connect_valueChanged")
	reg(&pg_QDial_connect_sliderPressed, "QDial_connect_sliderPressed")
	reg(&pg_QDial_connect_sliderMoved, "QDial_connect_sliderMoved")
	reg(&pg_QDial_connect_sliderReleased, "QDial_connect_sliderReleased")
	reg(&pg_QDial_connect_rangeChanged, "QDial_connect_rangeChanged")
	reg(&pg_QDial_connect_actionTriggered, "QDial_connect_actionTriggered")
	reg(&pg_QDial_connect_windowTitleChanged, "QDial_connect_windowTitleChanged")
	reg(&pg_QDial_connect_windowIconChanged, "QDial_connect_windowIconChanged")
	reg(&pg_QDial_connect_windowIconTextChanged, "QDial_connect_windowIconTextChanged")
	reg(&pg_QDial_connect_customContextMenuRequested, "QDial_connect_customContextMenuRequested")
	reg(&pg_QDial_connect_destroyed, "QDial_connect_destroyed")
	reg(&pg_QDial_connect_destroyed2, "QDial_connect_destroyed2")
	reg(&pg_QDial_connect_objectNameChanged, "QDial_connect_objectNameChanged")
	reg(&pg_QDial__disconnect, "QDial__disconnect")
	reg(&pg_QDial__set_override, "QDial__set_override")
	reg(&pg_QDial_delete, "QDial_delete")
}

type QDial struct {
	h uintptr
	*QAbstractSlider
}

func bindQDial(h uintptr) *QDial {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QDial_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDial{h: h, QAbstractSlider: &QAbstractSlider{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQDial(h uintptr) *QDial { return bindQDial(h) }

func (this *QDial) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDial) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDial(h unsafe.Pointer) *QDial { return bindQDial(uintptr(h)) }
func UnsafeNewQDialOwned(h unsafe.Pointer) *QDial { return newownedQDial(uintptr(h)) }

func (this *QDial) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSlider != nil {
		this.QAbstractSlider.UnsafeRelease()
	}
}

func NewQDial(parent *QWidget) *QDial {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDial_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDial(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDial_virtIdx, QDial_enable)
	return __o
}

func NewQDial2() *QDial {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDial_new2, uintptr(__goid))
	__o := newownedQDial(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDial_virtIdx, QDial_enable)
	return __o
}

func QDial_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDial_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDial) Wrapping() bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_wrapping, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDial) NotchSize() int {
	__r0, _, _ := purego.SyscallN(addr_QDial_notchSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDial) SetNotchTarget(target float64) {
	pg_QDial_setNotchTarget(this.h, target)
	runtime.KeepAlive(this)
}

func (this *QDial) NotchTarget() float64 {
	__r := pg_QDial_notchTarget(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDial) NotchesVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_notchesVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDial) SetNotchesVisible(visible bool) {
	purego.SyscallN(addr_QDial_setNotchesVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDial) SetWrapping(on bool) {
	purego.SyscallN(addr_QDial_setWrapping, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QDial) SetRepeatAction(action QAbstractSlider__SliderAction) {
	purego.SyscallN(addr_QDial_setRepeatAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QDial) RepeatAction() QAbstractSlider__SliderAction {
	__r0, _, _ := purego.SyscallN(addr_QDial_repeatAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSlider__SliderAction(__r)
}

func QDial_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDial_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDial_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_mouseGrabber)
	return bindQWidget(__r)
}

func QDial_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_keyboardGrabber)
	return bindQWidget(__r)
}

func QDial_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDial_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDial) UpdateMicroFocus() {
	purego.SyscallN(addr_QDial_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDial) Create() {
	purego.SyscallN(addr_QDial_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDial) Destroy() {
	purego.SyscallN(addr_QDial_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDial) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDial) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDial_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDial_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDial_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDial_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDial) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDial_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDial) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDial_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDial) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDial_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDial) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDial_DevicePixelRatioFScale() float64 {
	__r := pg_QDial_devicePixelRatioFScale()
	return __r
}

func QDial_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDial_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDial) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDial_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDial_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDial_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDial_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDial_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QDial) SetRepeatAction2(action QAbstractSlider__SliderAction, thresholdTime int) {
	purego.SyscallN(addr_QDial_setRepeatAction2, this.h, uintptr(action), uintptr(thresholdTime))
	runtime.KeepAlive(this)
}

func (this *QDial) SetRepeatAction3(action QAbstractSlider__SliderAction, thresholdTime int, repeatTime int) {
	purego.SyscallN(addr_QDial_setRepeatAction3, this.h, uintptr(action), uintptr(thresholdTime), uintptr(repeatTime))
	runtime.KeepAlive(this)
}

func QDial_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDial_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDial_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDial) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDial_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDial) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDial_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDial) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDial_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDial) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDial_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDial) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDial_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDial) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDial_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDial_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDial_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDial_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDial_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDial) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDial_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDial) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDial_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDial) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDial_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDial) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QDial) ResizeEventDefault(re *QResizeEvent) {
	purego.SyscallN(addr_QDial_resizeEvent_default, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
}

func (this *QDial) PaintEventDefault(pe *QPaintEvent) {
	purego.SyscallN(addr_QDial_paintEvent_default, this.h, uintptr(pe.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pe)
}

func (this *QDial) MousePressEventDefault(me *QMouseEvent) {
	purego.SyscallN(addr_QDial_mousePressEvent_default, this.h, uintptr(me.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(me)
}

func (this *QDial) MouseReleaseEventDefault(me *QMouseEvent) {
	purego.SyscallN(addr_QDial_mouseReleaseEvent_default, this.h, uintptr(me.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(me)
}

func (this *QDial) MouseMoveEventDefault(me *QMouseEvent) {
	purego.SyscallN(addr_QDial_mouseMoveEvent_default, this.h, uintptr(me.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(me)
}

func (this *QDial) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	purego.SyscallN(addr_QDial_sliderChange_default, this.h, uintptr(change))
	runtime.KeepAlive(this)
}

func (this *QDial) InitStyleOptionDefault(option *QStyleOptionSlider) {
	purego.SyscallN(addr_QDial_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QDial) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QDial_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QDial) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QDial_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDial) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QDial_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QDial) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QDial_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QDial) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDial_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDial) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDial_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDial) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDial_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDial) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDial) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDial_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDial) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDial_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDial_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDial_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDial_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDial_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDial_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDial_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDial_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDial_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDial_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDial_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDial_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDial_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDial_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDial_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDial_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDial_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDial_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDial) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDial_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDial) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDial_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDial) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDial_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDial) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDial_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDial) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDial_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDial) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDial) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDial_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDial) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDial_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDial_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDial) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDial_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDial) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDial_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDial) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QDial_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnSliderPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDial_connect_sliderPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnSliderMoved(fn func(int)) int64 {
	id := registerCallback(func(position int32) { fn(int(position)) })
	pg_QDial_connect_sliderMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnSliderReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDial_connect_sliderReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnRangeChanged(fn func(int, int)) int64 {
	id := registerCallback(func(min int32, max int32) { fn(int(min), int(max)) })
	pg_QDial_connect_rangeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnActionTriggered(fn func(int)) int64 {
	id := registerCallback(func(action int32) { fn(int(action)) })
	pg_QDial_connect_actionTriggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDial_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDial_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDial_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDial_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDial_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDial_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDial_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDial) Disconnect(handle int64) {
	pg_QDial__disconnect(handle)
	unregisterCallback(handle)
}

func QDial_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case -4302917285655014661:
		return 4
	case 5290442163167458465:
		return 5
	case -332090898078543319:
		return 6
	case -4415920049297795213:
		return 7
	case -4255494004099806581:
		return 8
	case 4658628703300230988:
		return 9
	case -488693769771308531:
		return 10
	case 1023791977679793185:
		return 11
	case -5409927581812374498:
		return 12
	case 1919988340934511316:
		return 13
	case 4959376063826165943:
		return 14
	case -8659679893868524044:
		return 15
	case 2110182719375245575:
		return 16
	case 4311534060495737339:
		return 17
	case -8266762397995423915:
		return 18
	case 6398007404756346381:
		return 19
	case -4029073905697023887:
		return 20
	case 772598475950050299:
		return 21
	case 975043767033552552:
		return 22
	case -385596010576048993:
		return 23
	case 6316270346668237895:
		return 24
	case -551070009623711640:
		return 25
	case -9161949150489525396:
		return 26
	case -1989937050019464783:
		return 27
	case -6698043222227006015:
		return 28
	case -212068866860701141:
		return 29
	case 1697444140088555581:
		return 30
	case 206122715919917589:
		return 31
	case -1781545137352745978:
		return 32
	case -3139904158373725006:
		return 33
	case 9107058682754591716:
		return 34
	case -6825369174751916:
		return 35
	case -199718707876719129:
		return 36
	case -2039166586246443706:
		return 37
	case 427629649267311777:
		return 38
	case 6477500296980050210:
		return 39
	case 4127240633310655048:
		return 40
	case 5792771328741774991:
		return 41
	case 5438819661942843222:
		return 42
	case 7345853097340709638:
		return 43
	case 5898328517505784191:
		return 44
	case 5217623514630858297:
		return 45
	case -1185315367526857833:
		return 46
	case 3769868229123868474:
		return 47
	case 7379029982170503008:
		return 48
	case -568369772982215962:
		return 49
	}
	return -1
}

func QDial_enable(h uintptr, idx int32) { pg_QDial__set_override(h, idx) }

func (this *QDial) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDial) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDial) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDial) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QDial) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(re uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(re))) }) // slot "resizeEvent"
}

func (this *QDial) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(pe uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(pe))) }) // slot "paintEvent"
}

func (this *QDial) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(me uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(me))) }) // slot "mousePressEvent"
}

func (this *QDial) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(me uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(me))) }) // slot "mouseReleaseEvent"
}

func (this *QDial) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(me uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(me))) }) // slot "mouseMoveEvent"
}

func (this *QDial) OnSliderChange(fn func(QAbstractSlider__SliderChange)) {
	registerVirt(this.UnsafeGoID(), 4658628703300230988, func(change int32) { fn(QAbstractSlider__SliderChange(change)) }) // slot "sliderChange"
}

func (this *QDial) OnInitStyleOption(fn func(*QStyleOptionSlider)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSlider(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QDial) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QDial) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QDial) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QDial) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QDial) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDial) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDial) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDial) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDial) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDial) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDial) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDial) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDial) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDial) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDial) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDial) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDial) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDial) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDial) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDial) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDial) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDial) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDial) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDial) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDial) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDial) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDial) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDial) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDial) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDial) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDial) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDial) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDial) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDial) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDial) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDial) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDial) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDial) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDial) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDial) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDial_delete(this.h)
	}
	this.UnsafeRelease()
}
