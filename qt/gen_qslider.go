package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSlider_chain uintptr
	addr_QSlider_new uintptr
	addr_QSlider_new2 uintptr
	addr_QSlider_new3 uintptr
	addr_QSlider_new4 uintptr
	addr_QSlider_tr uintptr
	addr_QSlider_setTickPosition uintptr
	addr_QSlider_tickPosition uintptr
	addr_QSlider_setTickInterval uintptr
	addr_QSlider_tickInterval uintptr
	addr_QSlider_setRepeatAction uintptr
	addr_QSlider_repeatAction uintptr
	addr_QSlider_setTabOrder uintptr
	addr_QSlider_mouseGrabber uintptr
	addr_QSlider_keyboardGrabber uintptr
	addr_QSlider_find uintptr
	addr_QSlider_createWindowContainer uintptr
	addr_QSlider_updateMicroFocus uintptr
	addr_QSlider_create uintptr
	addr_QSlider_destroy uintptr
	addr_QSlider_focusNextChild uintptr
	addr_QSlider_focusPreviousChild uintptr
	addr_QSlider_connect uintptr
	addr_QSlider_connect2 uintptr
	addr_QSlider_sender uintptr
	addr_QSlider_senderSignalIndex uintptr
	addr_QSlider_receivers uintptr
	addr_QSlider_isSignalConnected uintptr
	pg_QSlider_devicePixelRatioFScale func() float64
	pg_QSlider_encodeMetricF func(int32, float64) int32
	pg_QSlider_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QSlider_tr2 uintptr
	addr_QSlider_tr3 uintptr
	addr_QSlider_setRepeatAction2 uintptr
	addr_QSlider_setRepeatAction3 uintptr
	addr_QSlider_createWindowContainer2 uintptr
	addr_QSlider_createWindowContainer3 uintptr
	addr_QSlider_updateMicroFocus2 uintptr
	addr_QSlider_create2 uintptr
	addr_QSlider_create3 uintptr
	addr_QSlider_create4 uintptr
	addr_QSlider_destroy2 uintptr
	addr_QSlider_destroy3 uintptr
	addr_QSlider_connect3 uintptr
	addr_QSlider_connect4 uintptr
	addr_QSlider_metaObject_default uintptr
	addr_QSlider_sizeHint_default uintptr
	addr_QSlider_minimumSizeHint_default uintptr
	addr_QSlider_event_default uintptr
	addr_QSlider_paintEvent_default uintptr
	addr_QSlider_mousePressEvent_default uintptr
	addr_QSlider_mouseReleaseEvent_default uintptr
	addr_QSlider_mouseMoveEvent_default uintptr
	addr_QSlider_initStyleOption_default uintptr
	addr_QSlider_sliderChange_default uintptr
	addr_QSlider_keyPressEvent_default uintptr
	addr_QSlider_timerEvent_default uintptr
	addr_QSlider_wheelEvent_default uintptr
	addr_QSlider_changeEvent_default uintptr
	addr_QSlider_devType_default uintptr
	addr_QSlider_setVisible_default uintptr
	addr_QSlider_heightForWidth_default uintptr
	addr_QSlider_hasHeightForWidth_default uintptr
	addr_QSlider_paintEngine_default uintptr
	addr_QSlider_mouseDoubleClickEvent_default uintptr
	addr_QSlider_keyReleaseEvent_default uintptr
	addr_QSlider_focusInEvent_default uintptr
	addr_QSlider_focusOutEvent_default uintptr
	addr_QSlider_enterEvent_default uintptr
	addr_QSlider_leaveEvent_default uintptr
	addr_QSlider_moveEvent_default uintptr
	addr_QSlider_resizeEvent_default uintptr
	addr_QSlider_closeEvent_default uintptr
	addr_QSlider_contextMenuEvent_default uintptr
	addr_QSlider_tabletEvent_default uintptr
	addr_QSlider_actionEvent_default uintptr
	addr_QSlider_dragEnterEvent_default uintptr
	addr_QSlider_dragMoveEvent_default uintptr
	addr_QSlider_dragLeaveEvent_default uintptr
	addr_QSlider_dropEvent_default uintptr
	addr_QSlider_showEvent_default uintptr
	addr_QSlider_hideEvent_default uintptr
	addr_QSlider_metric_default uintptr
	addr_QSlider_initPainter_default uintptr
	addr_QSlider_redirected_default uintptr
	addr_QSlider_sharedPainter_default uintptr
	addr_QSlider_inputMethodEvent_default uintptr
	addr_QSlider_inputMethodQuery_default uintptr
	addr_QSlider_focusNextPrevChild_default uintptr
	addr_QSlider_eventFilter_default uintptr
	addr_QSlider_childEvent_default uintptr
	addr_QSlider_customEvent_default uintptr
	addr_QSlider_connectNotify_default uintptr
	addr_QSlider_disconnectNotify_default uintptr
	pg_QSlider_connect_valueChanged func(uintptr, int64)
	pg_QSlider_connect_sliderPressed func(uintptr, int64)
	pg_QSlider_connect_sliderMoved func(uintptr, int64)
	pg_QSlider_connect_sliderReleased func(uintptr, int64)
	pg_QSlider_connect_rangeChanged func(uintptr, int64)
	pg_QSlider_connect_actionTriggered func(uintptr, int64)
	pg_QSlider_connect_windowTitleChanged func(uintptr, int64)
	pg_QSlider_connect_windowIconChanged func(uintptr, int64)
	pg_QSlider_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSlider_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSlider_connect_destroyed func(uintptr, int64)
	pg_QSlider_connect_destroyed2 func(uintptr, int64)
	pg_QSlider_connect_objectNameChanged func(uintptr, int64)
	pg_QSlider__disconnect func(int64)
	pg_QSlider__set_override func(uintptr, int32)
	pg_QSlider_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSlider_chain, "QSlider_chain")
	symaddr(&addr_QSlider_new, "QSlider_new")
	symaddr(&addr_QSlider_new2, "QSlider_new2")
	symaddr(&addr_QSlider_new3, "QSlider_new3")
	symaddr(&addr_QSlider_new4, "QSlider_new4")
	symaddr(&addr_QSlider_tr, "QSlider_tr")
	symaddr(&addr_QSlider_setTickPosition, "QSlider_setTickPosition")
	symaddr(&addr_QSlider_tickPosition, "QSlider_tickPosition")
	symaddr(&addr_QSlider_setTickInterval, "QSlider_setTickInterval")
	symaddr(&addr_QSlider_tickInterval, "QSlider_tickInterval")
	symaddr(&addr_QSlider_setRepeatAction, "QSlider_setRepeatAction")
	symaddr(&addr_QSlider_repeatAction, "QSlider_repeatAction")
	symaddr(&addr_QSlider_setTabOrder, "QSlider_setTabOrder")
	symaddr(&addr_QSlider_mouseGrabber, "QSlider_mouseGrabber")
	symaddr(&addr_QSlider_keyboardGrabber, "QSlider_keyboardGrabber")
	symaddr(&addr_QSlider_find, "QSlider_find")
	symaddr(&addr_QSlider_createWindowContainer, "QSlider_createWindowContainer")
	symaddr(&addr_QSlider_updateMicroFocus, "QSlider_updateMicroFocus")
	symaddr(&addr_QSlider_create, "QSlider_create")
	symaddr(&addr_QSlider_destroy, "QSlider_destroy")
	symaddr(&addr_QSlider_focusNextChild, "QSlider_focusNextChild")
	symaddr(&addr_QSlider_focusPreviousChild, "QSlider_focusPreviousChild")
	symaddr(&addr_QSlider_connect, "QSlider_connect")
	symaddr(&addr_QSlider_connect2, "QSlider_connect2")
	symaddr(&addr_QSlider_sender, "QSlider_sender")
	symaddr(&addr_QSlider_senderSignalIndex, "QSlider_senderSignalIndex")
	symaddr(&addr_QSlider_receivers, "QSlider_receivers")
	symaddr(&addr_QSlider_isSignalConnected, "QSlider_isSignalConnected")
	reg(&pg_QSlider_devicePixelRatioFScale, "QSlider_devicePixelRatioFScale")
	reg(&pg_QSlider_encodeMetricF, "QSlider_encodeMetricF")
	reg(&pg_QSlider_getDecodedMetricF, "QSlider_getDecodedMetricF")
	symaddr(&addr_QSlider_tr2, "QSlider_tr2")
	symaddr(&addr_QSlider_tr3, "QSlider_tr3")
	symaddr(&addr_QSlider_setRepeatAction2, "QSlider_setRepeatAction2")
	symaddr(&addr_QSlider_setRepeatAction3, "QSlider_setRepeatAction3")
	symaddr(&addr_QSlider_createWindowContainer2, "QSlider_createWindowContainer2")
	symaddr(&addr_QSlider_createWindowContainer3, "QSlider_createWindowContainer3")
	symaddr(&addr_QSlider_updateMicroFocus2, "QSlider_updateMicroFocus2")
	symaddr(&addr_QSlider_create2, "QSlider_create2")
	symaddr(&addr_QSlider_create3, "QSlider_create3")
	symaddr(&addr_QSlider_create4, "QSlider_create4")
	symaddr(&addr_QSlider_destroy2, "QSlider_destroy2")
	symaddr(&addr_QSlider_destroy3, "QSlider_destroy3")
	symaddr(&addr_QSlider_connect3, "QSlider_connect3")
	symaddr(&addr_QSlider_connect4, "QSlider_connect4")
	symaddr(&addr_QSlider_metaObject_default, "QSlider_metaObject_default")
	symaddr(&addr_QSlider_sizeHint_default, "QSlider_sizeHint_default")
	symaddr(&addr_QSlider_minimumSizeHint_default, "QSlider_minimumSizeHint_default")
	symaddr(&addr_QSlider_event_default, "QSlider_event_default")
	symaddr(&addr_QSlider_paintEvent_default, "QSlider_paintEvent_default")
	symaddr(&addr_QSlider_mousePressEvent_default, "QSlider_mousePressEvent_default")
	symaddr(&addr_QSlider_mouseReleaseEvent_default, "QSlider_mouseReleaseEvent_default")
	symaddr(&addr_QSlider_mouseMoveEvent_default, "QSlider_mouseMoveEvent_default")
	symaddr(&addr_QSlider_initStyleOption_default, "QSlider_initStyleOption_default")
	symaddr(&addr_QSlider_sliderChange_default, "QSlider_sliderChange_default")
	symaddr(&addr_QSlider_keyPressEvent_default, "QSlider_keyPressEvent_default")
	symaddr(&addr_QSlider_timerEvent_default, "QSlider_timerEvent_default")
	symaddr(&addr_QSlider_wheelEvent_default, "QSlider_wheelEvent_default")
	symaddr(&addr_QSlider_changeEvent_default, "QSlider_changeEvent_default")
	symaddr(&addr_QSlider_devType_default, "QSlider_devType_default")
	symaddr(&addr_QSlider_setVisible_default, "QSlider_setVisible_default")
	symaddr(&addr_QSlider_heightForWidth_default, "QSlider_heightForWidth_default")
	symaddr(&addr_QSlider_hasHeightForWidth_default, "QSlider_hasHeightForWidth_default")
	symaddr(&addr_QSlider_paintEngine_default, "QSlider_paintEngine_default")
	symaddr(&addr_QSlider_mouseDoubleClickEvent_default, "QSlider_mouseDoubleClickEvent_default")
	symaddr(&addr_QSlider_keyReleaseEvent_default, "QSlider_keyReleaseEvent_default")
	symaddr(&addr_QSlider_focusInEvent_default, "QSlider_focusInEvent_default")
	symaddr(&addr_QSlider_focusOutEvent_default, "QSlider_focusOutEvent_default")
	symaddr(&addr_QSlider_enterEvent_default, "QSlider_enterEvent_default")
	symaddr(&addr_QSlider_leaveEvent_default, "QSlider_leaveEvent_default")
	symaddr(&addr_QSlider_moveEvent_default, "QSlider_moveEvent_default")
	symaddr(&addr_QSlider_resizeEvent_default, "QSlider_resizeEvent_default")
	symaddr(&addr_QSlider_closeEvent_default, "QSlider_closeEvent_default")
	symaddr(&addr_QSlider_contextMenuEvent_default, "QSlider_contextMenuEvent_default")
	symaddr(&addr_QSlider_tabletEvent_default, "QSlider_tabletEvent_default")
	symaddr(&addr_QSlider_actionEvent_default, "QSlider_actionEvent_default")
	symaddr(&addr_QSlider_dragEnterEvent_default, "QSlider_dragEnterEvent_default")
	symaddr(&addr_QSlider_dragMoveEvent_default, "QSlider_dragMoveEvent_default")
	symaddr(&addr_QSlider_dragLeaveEvent_default, "QSlider_dragLeaveEvent_default")
	symaddr(&addr_QSlider_dropEvent_default, "QSlider_dropEvent_default")
	symaddr(&addr_QSlider_showEvent_default, "QSlider_showEvent_default")
	symaddr(&addr_QSlider_hideEvent_default, "QSlider_hideEvent_default")
	symaddr(&addr_QSlider_metric_default, "QSlider_metric_default")
	symaddr(&addr_QSlider_initPainter_default, "QSlider_initPainter_default")
	symaddr(&addr_QSlider_redirected_default, "QSlider_redirected_default")
	symaddr(&addr_QSlider_sharedPainter_default, "QSlider_sharedPainter_default")
	symaddr(&addr_QSlider_inputMethodEvent_default, "QSlider_inputMethodEvent_default")
	symaddr(&addr_QSlider_inputMethodQuery_default, "QSlider_inputMethodQuery_default")
	symaddr(&addr_QSlider_focusNextPrevChild_default, "QSlider_focusNextPrevChild_default")
	symaddr(&addr_QSlider_eventFilter_default, "QSlider_eventFilter_default")
	symaddr(&addr_QSlider_childEvent_default, "QSlider_childEvent_default")
	symaddr(&addr_QSlider_customEvent_default, "QSlider_customEvent_default")
	symaddr(&addr_QSlider_connectNotify_default, "QSlider_connectNotify_default")
	symaddr(&addr_QSlider_disconnectNotify_default, "QSlider_disconnectNotify_default")
	reg(&pg_QSlider_connect_valueChanged, "QSlider_connect_valueChanged")
	reg(&pg_QSlider_connect_sliderPressed, "QSlider_connect_sliderPressed")
	reg(&pg_QSlider_connect_sliderMoved, "QSlider_connect_sliderMoved")
	reg(&pg_QSlider_connect_sliderReleased, "QSlider_connect_sliderReleased")
	reg(&pg_QSlider_connect_rangeChanged, "QSlider_connect_rangeChanged")
	reg(&pg_QSlider_connect_actionTriggered, "QSlider_connect_actionTriggered")
	reg(&pg_QSlider_connect_windowTitleChanged, "QSlider_connect_windowTitleChanged")
	reg(&pg_QSlider_connect_windowIconChanged, "QSlider_connect_windowIconChanged")
	reg(&pg_QSlider_connect_windowIconTextChanged, "QSlider_connect_windowIconTextChanged")
	reg(&pg_QSlider_connect_customContextMenuRequested, "QSlider_connect_customContextMenuRequested")
	reg(&pg_QSlider_connect_destroyed, "QSlider_connect_destroyed")
	reg(&pg_QSlider_connect_destroyed2, "QSlider_connect_destroyed2")
	reg(&pg_QSlider_connect_objectNameChanged, "QSlider_connect_objectNameChanged")
	reg(&pg_QSlider__disconnect, "QSlider__disconnect")
	reg(&pg_QSlider__set_override, "QSlider__set_override")
	reg(&pg_QSlider_delete, "QSlider_delete")
}

type QSlider struct {
	h uintptr
	*QAbstractSlider
}

func bindQSlider(h uintptr) *QSlider {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSlider_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSlider{h: h, QAbstractSlider: &QAbstractSlider{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSlider(h uintptr) *QSlider { return bindQSlider(h) }

func (this *QSlider) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSlider) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSlider(h unsafe.Pointer) *QSlider { return bindQSlider(uintptr(h)) }
func UnsafeNewQSliderOwned(h unsafe.Pointer) *QSlider { return newownedQSlider(uintptr(h)) }

func (this *QSlider) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSlider != nil {
		this.QAbstractSlider.UnsafeRelease()
	}
}

func NewQSlider(parent *QWidget) *QSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSlider_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSlider(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSlider_virtIdx, QSlider_enable)
	return __o
}

func NewQSlider2() *QSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSlider_new2, uintptr(__goid))
	__o := newownedQSlider(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSlider_virtIdx, QSlider_enable)
	return __o
}

func NewQSlider3(orientation Orientation) *QSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSlider_new3, uintptr(__goid), uintptr(orientation))
	__o := newownedQSlider(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSlider_virtIdx, QSlider_enable)
	return __o
}

func NewQSlider4(orientation Orientation, parent *QWidget) *QSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSlider_new4, uintptr(__goid), uintptr(orientation), uintptr(parent.UnsafePointer()))
	__o := newownedQSlider(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSlider_virtIdx, QSlider_enable)
	return __o
}

func QSlider_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSlider_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSlider) SetTickPosition(position QSlider__TickPosition) {
	purego.SyscallN(addr_QSlider_setTickPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QSlider) TickPosition() QSlider__TickPosition {
	__r0, _, _ := purego.SyscallN(addr_QSlider_tickPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSlider__TickPosition(__r)
}

func (this *QSlider) SetTickInterval(ti int) {
	purego.SyscallN(addr_QSlider_setTickInterval, this.h, uintptr(ti))
	runtime.KeepAlive(this)
}

func (this *QSlider) TickInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QSlider_tickInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSlider) SetRepeatAction(action QAbstractSlider__SliderAction) {
	purego.SyscallN(addr_QSlider_setRepeatAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QSlider) RepeatAction() QAbstractSlider__SliderAction {
	__r0, _, _ := purego.SyscallN(addr_QSlider_repeatAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSlider__SliderAction(__r)
}

func QSlider_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSlider_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSlider_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_mouseGrabber)
	return bindQWidget(__r)
}

func QSlider_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_keyboardGrabber)
	return bindQWidget(__r)
}

func QSlider_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSlider_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QSlider) UpdateMicroFocus() {
	purego.SyscallN(addr_QSlider_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QSlider) Create() {
	purego.SyscallN(addr_QSlider_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QSlider) Destroy() {
	purego.SyscallN(addr_QSlider_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QSlider) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSlider) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSlider_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSlider_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSlider_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSlider_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSlider) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSlider_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSlider) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSlider_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSlider) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSlider_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSlider) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSlider_DevicePixelRatioFScale() float64 {
	__r := pg_QSlider_devicePixelRatioFScale()
	return __r
}

func QSlider_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSlider_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QSlider) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QSlider_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QSlider_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSlider_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSlider_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSlider_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSlider) SetRepeatAction2(action QAbstractSlider__SliderAction, thresholdTime int) {
	purego.SyscallN(addr_QSlider_setRepeatAction2, this.h, uintptr(action), uintptr(thresholdTime))
	runtime.KeepAlive(this)
}

func (this *QSlider) SetRepeatAction3(action QAbstractSlider__SliderAction, thresholdTime int, repeatTime int) {
	purego.SyscallN(addr_QSlider_setRepeatAction3, this.h, uintptr(action), uintptr(thresholdTime), uintptr(repeatTime))
	runtime.KeepAlive(this)
}

func QSlider_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSlider_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSlider_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QSlider) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QSlider_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QSlider) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QSlider_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSlider) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QSlider_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QSlider) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QSlider_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QSlider) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QSlider_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QSlider) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QSlider_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QSlider_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSlider_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSlider_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSlider_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSlider) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSlider_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSlider) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSlider_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSlider) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSlider_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSlider) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSlider) PaintEventDefault(ev *QPaintEvent) {
	purego.SyscallN(addr_QSlider_paintEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QSlider) MousePressEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QSlider_mousePressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QSlider) MouseReleaseEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QSlider_mouseReleaseEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QSlider) MouseMoveEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QSlider_mouseMoveEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QSlider) InitStyleOptionDefault(option *QStyleOptionSlider) {
	purego.SyscallN(addr_QSlider_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QSlider) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	purego.SyscallN(addr_QSlider_sliderChange_default, this.h, uintptr(change))
	runtime.KeepAlive(this)
}

func (this *QSlider) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QSlider_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QSlider) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QSlider_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSlider) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QSlider_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QSlider) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QSlider_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QSlider) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QSlider_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSlider) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QSlider_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QSlider) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QSlider_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSlider) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSlider) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QSlider_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QSlider) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSlider_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSlider_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSlider_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSlider_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QSlider_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSlider_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QSlider_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QSlider_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QSlider_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QSlider_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QSlider_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QSlider_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QSlider_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QSlider_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QSlider_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QSlider_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QSlider_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QSlider_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QSlider_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSlider) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QSlider_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QSlider) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QSlider_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QSlider) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QSlider_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QSlider) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QSlider_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSlider) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSlider_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSlider) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSlider) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSlider_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSlider) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSlider_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSlider_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSlider) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSlider_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSlider) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSlider_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSlider) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QSlider_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnSliderPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSlider_connect_sliderPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnSliderMoved(fn func(int)) int64 {
	id := registerCallback(func(position int32) { fn(int(position)) })
	pg_QSlider_connect_sliderMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnSliderReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSlider_connect_sliderReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnRangeChanged(fn func(int, int)) int64 {
	id := registerCallback(func(min int32, max int32) { fn(int(min), int(max)) })
	pg_QSlider_connect_rangeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnActionTriggered(fn func(int)) int64 {
	id := registerCallback(func(action int32) { fn(int(action)) })
	pg_QSlider_connect_actionTriggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSlider_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSlider_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSlider_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSlider_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSlider_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSlider_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSlider_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSlider) Disconnect(handle int64) {
	pg_QSlider__disconnect(handle)
	unregisterCallback(handle)
}

func QSlider_virtIdx(slot int64) int32 {
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
	case -332090898078543319:
		return 5
	case -4415920049297795213:
		return 6
	case -4255494004099806581:
		return 7
	case -488693769771308531:
		return 8
	case 4658628703300230988:
		return 9
	case 1023791977679793185:
		return 10
	case -5409927581812374498:
		return 11
	case 1919988340934511316:
		return 12
	case 4959376063826165943:
		return 13
	case -8659679893868524044:
		return 14
	case 2110182719375245575:
		return 15
	case 4311534060495737339:
		return 16
	case -8266762397995423915:
		return 17
	case 6398007404756346381:
		return 18
	case -4029073905697023887:
		return 19
	case 772598475950050299:
		return 20
	case 975043767033552552:
		return 21
	case -385596010576048993:
		return 22
	case 6316270346668237895:
		return 23
	case -551070009623711640:
		return 24
	case -9161949150489525396:
		return 25
	case -4302917285655014661:
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

func QSlider_enable(h uintptr, idx int32) { pg_QSlider__set_override(h, idx) }

func (this *QSlider) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSlider) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QSlider) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QSlider) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSlider) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(ev uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(ev))) }) // slot "paintEvent"
}

func (this *QSlider) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mousePressEvent"
}

func (this *QSlider) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseReleaseEvent"
}

func (this *QSlider) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseMoveEvent"
}

func (this *QSlider) OnInitStyleOption(fn func(*QStyleOptionSlider)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSlider(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QSlider) OnSliderChange(fn func(QAbstractSlider__SliderChange)) {
	registerVirt(this.UnsafeGoID(), 4658628703300230988, func(change int32) { fn(QAbstractSlider__SliderChange(change)) }) // slot "sliderChange"
}

func (this *QSlider) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QSlider) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QSlider) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QSlider) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QSlider) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QSlider) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QSlider) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QSlider) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QSlider) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QSlider) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QSlider) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QSlider) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QSlider) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QSlider) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QSlider) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QSlider) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QSlider) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QSlider) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QSlider) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QSlider) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QSlider) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QSlider) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QSlider) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QSlider) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QSlider) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QSlider) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QSlider) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QSlider) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QSlider) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QSlider) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QSlider) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QSlider) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QSlider) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QSlider) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QSlider) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QSlider) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSlider) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSlider) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSlider) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSlider) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSlider) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSlider_delete(this.h)
	}
	this.UnsafeRelease()
}
