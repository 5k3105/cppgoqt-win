package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QScrollBar_chain uintptr
	addr_QScrollBar_new uintptr
	addr_QScrollBar_new2 uintptr
	addr_QScrollBar_new3 uintptr
	addr_QScrollBar_new4 uintptr
	addr_QScrollBar_tr uintptr
	addr_QScrollBar_createStandardContextMenu uintptr
	addr_QScrollBar_setRepeatAction uintptr
	addr_QScrollBar_repeatAction uintptr
	addr_QScrollBar_setTabOrder uintptr
	addr_QScrollBar_mouseGrabber uintptr
	addr_QScrollBar_keyboardGrabber uintptr
	addr_QScrollBar_find uintptr
	addr_QScrollBar_createWindowContainer uintptr
	addr_QScrollBar_updateMicroFocus uintptr
	addr_QScrollBar_create uintptr
	addr_QScrollBar_destroy uintptr
	addr_QScrollBar_focusNextChild uintptr
	addr_QScrollBar_focusPreviousChild uintptr
	addr_QScrollBar_connect uintptr
	addr_QScrollBar_connect2 uintptr
	addr_QScrollBar_sender uintptr
	addr_QScrollBar_senderSignalIndex uintptr
	addr_QScrollBar_receivers uintptr
	addr_QScrollBar_isSignalConnected uintptr
	pg_QScrollBar_devicePixelRatioFScale func() float64
	pg_QScrollBar_encodeMetricF func(int32, float64) int32
	pg_QScrollBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QScrollBar_tr2 uintptr
	addr_QScrollBar_tr3 uintptr
	addr_QScrollBar_setRepeatAction2 uintptr
	addr_QScrollBar_setRepeatAction3 uintptr
	addr_QScrollBar_createWindowContainer2 uintptr
	addr_QScrollBar_createWindowContainer3 uintptr
	addr_QScrollBar_updateMicroFocus2 uintptr
	addr_QScrollBar_create2 uintptr
	addr_QScrollBar_create3 uintptr
	addr_QScrollBar_create4 uintptr
	addr_QScrollBar_destroy2 uintptr
	addr_QScrollBar_destroy3 uintptr
	addr_QScrollBar_connect3 uintptr
	addr_QScrollBar_connect4 uintptr
	addr_QScrollBar_metaObject_default uintptr
	addr_QScrollBar_sizeHint_default uintptr
	addr_QScrollBar_event_default uintptr
	addr_QScrollBar_wheelEvent_default uintptr
	addr_QScrollBar_paintEvent_default uintptr
	addr_QScrollBar_mousePressEvent_default uintptr
	addr_QScrollBar_mouseReleaseEvent_default uintptr
	addr_QScrollBar_mouseMoveEvent_default uintptr
	addr_QScrollBar_hideEvent_default uintptr
	addr_QScrollBar_sliderChange_default uintptr
	addr_QScrollBar_contextMenuEvent_default uintptr
	addr_QScrollBar_initStyleOption_default uintptr
	addr_QScrollBar_keyPressEvent_default uintptr
	addr_QScrollBar_timerEvent_default uintptr
	addr_QScrollBar_changeEvent_default uintptr
	addr_QScrollBar_devType_default uintptr
	addr_QScrollBar_setVisible_default uintptr
	addr_QScrollBar_minimumSizeHint_default uintptr
	addr_QScrollBar_heightForWidth_default uintptr
	addr_QScrollBar_hasHeightForWidth_default uintptr
	addr_QScrollBar_paintEngine_default uintptr
	addr_QScrollBar_mouseDoubleClickEvent_default uintptr
	addr_QScrollBar_keyReleaseEvent_default uintptr
	addr_QScrollBar_focusInEvent_default uintptr
	addr_QScrollBar_focusOutEvent_default uintptr
	addr_QScrollBar_enterEvent_default uintptr
	addr_QScrollBar_leaveEvent_default uintptr
	addr_QScrollBar_moveEvent_default uintptr
	addr_QScrollBar_resizeEvent_default uintptr
	addr_QScrollBar_closeEvent_default uintptr
	addr_QScrollBar_tabletEvent_default uintptr
	addr_QScrollBar_actionEvent_default uintptr
	addr_QScrollBar_dragEnterEvent_default uintptr
	addr_QScrollBar_dragMoveEvent_default uintptr
	addr_QScrollBar_dragLeaveEvent_default uintptr
	addr_QScrollBar_dropEvent_default uintptr
	addr_QScrollBar_showEvent_default uintptr
	addr_QScrollBar_metric_default uintptr
	addr_QScrollBar_initPainter_default uintptr
	addr_QScrollBar_redirected_default uintptr
	addr_QScrollBar_sharedPainter_default uintptr
	addr_QScrollBar_inputMethodEvent_default uintptr
	addr_QScrollBar_inputMethodQuery_default uintptr
	addr_QScrollBar_focusNextPrevChild_default uintptr
	addr_QScrollBar_eventFilter_default uintptr
	addr_QScrollBar_childEvent_default uintptr
	addr_QScrollBar_customEvent_default uintptr
	addr_QScrollBar_connectNotify_default uintptr
	addr_QScrollBar_disconnectNotify_default uintptr
	pg_QScrollBar_connect_valueChanged func(uintptr, int64)
	pg_QScrollBar_connect_sliderPressed func(uintptr, int64)
	pg_QScrollBar_connect_sliderMoved func(uintptr, int64)
	pg_QScrollBar_connect_sliderReleased func(uintptr, int64)
	pg_QScrollBar_connect_rangeChanged func(uintptr, int64)
	pg_QScrollBar_connect_actionTriggered func(uintptr, int64)
	pg_QScrollBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QScrollBar_connect_windowIconChanged func(uintptr, int64)
	pg_QScrollBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QScrollBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QScrollBar_connect_destroyed func(uintptr, int64)
	pg_QScrollBar_connect_destroyed2 func(uintptr, int64)
	pg_QScrollBar_connect_objectNameChanged func(uintptr, int64)
	pg_QScrollBar__disconnect func(int64)
	pg_QScrollBar__set_override func(uintptr, int32)
	pg_QScrollBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QScrollBar_chain, "QScrollBar_chain")
	symaddr(&addr_QScrollBar_new, "QScrollBar_new")
	symaddr(&addr_QScrollBar_new2, "QScrollBar_new2")
	symaddr(&addr_QScrollBar_new3, "QScrollBar_new3")
	symaddr(&addr_QScrollBar_new4, "QScrollBar_new4")
	symaddr(&addr_QScrollBar_tr, "QScrollBar_tr")
	symaddr(&addr_QScrollBar_createStandardContextMenu, "QScrollBar_createStandardContextMenu")
	symaddr(&addr_QScrollBar_setRepeatAction, "QScrollBar_setRepeatAction")
	symaddr(&addr_QScrollBar_repeatAction, "QScrollBar_repeatAction")
	symaddr(&addr_QScrollBar_setTabOrder, "QScrollBar_setTabOrder")
	symaddr(&addr_QScrollBar_mouseGrabber, "QScrollBar_mouseGrabber")
	symaddr(&addr_QScrollBar_keyboardGrabber, "QScrollBar_keyboardGrabber")
	symaddr(&addr_QScrollBar_find, "QScrollBar_find")
	symaddr(&addr_QScrollBar_createWindowContainer, "QScrollBar_createWindowContainer")
	symaddr(&addr_QScrollBar_updateMicroFocus, "QScrollBar_updateMicroFocus")
	symaddr(&addr_QScrollBar_create, "QScrollBar_create")
	symaddr(&addr_QScrollBar_destroy, "QScrollBar_destroy")
	symaddr(&addr_QScrollBar_focusNextChild, "QScrollBar_focusNextChild")
	symaddr(&addr_QScrollBar_focusPreviousChild, "QScrollBar_focusPreviousChild")
	symaddr(&addr_QScrollBar_connect, "QScrollBar_connect")
	symaddr(&addr_QScrollBar_connect2, "QScrollBar_connect2")
	symaddr(&addr_QScrollBar_sender, "QScrollBar_sender")
	symaddr(&addr_QScrollBar_senderSignalIndex, "QScrollBar_senderSignalIndex")
	symaddr(&addr_QScrollBar_receivers, "QScrollBar_receivers")
	symaddr(&addr_QScrollBar_isSignalConnected, "QScrollBar_isSignalConnected")
	reg(&pg_QScrollBar_devicePixelRatioFScale, "QScrollBar_devicePixelRatioFScale")
	reg(&pg_QScrollBar_encodeMetricF, "QScrollBar_encodeMetricF")
	reg(&pg_QScrollBar_getDecodedMetricF, "QScrollBar_getDecodedMetricF")
	symaddr(&addr_QScrollBar_tr2, "QScrollBar_tr2")
	symaddr(&addr_QScrollBar_tr3, "QScrollBar_tr3")
	symaddr(&addr_QScrollBar_setRepeatAction2, "QScrollBar_setRepeatAction2")
	symaddr(&addr_QScrollBar_setRepeatAction3, "QScrollBar_setRepeatAction3")
	symaddr(&addr_QScrollBar_createWindowContainer2, "QScrollBar_createWindowContainer2")
	symaddr(&addr_QScrollBar_createWindowContainer3, "QScrollBar_createWindowContainer3")
	symaddr(&addr_QScrollBar_updateMicroFocus2, "QScrollBar_updateMicroFocus2")
	symaddr(&addr_QScrollBar_create2, "QScrollBar_create2")
	symaddr(&addr_QScrollBar_create3, "QScrollBar_create3")
	symaddr(&addr_QScrollBar_create4, "QScrollBar_create4")
	symaddr(&addr_QScrollBar_destroy2, "QScrollBar_destroy2")
	symaddr(&addr_QScrollBar_destroy3, "QScrollBar_destroy3")
	symaddr(&addr_QScrollBar_connect3, "QScrollBar_connect3")
	symaddr(&addr_QScrollBar_connect4, "QScrollBar_connect4")
	symaddr(&addr_QScrollBar_metaObject_default, "QScrollBar_metaObject_default")
	symaddr(&addr_QScrollBar_sizeHint_default, "QScrollBar_sizeHint_default")
	symaddr(&addr_QScrollBar_event_default, "QScrollBar_event_default")
	symaddr(&addr_QScrollBar_wheelEvent_default, "QScrollBar_wheelEvent_default")
	symaddr(&addr_QScrollBar_paintEvent_default, "QScrollBar_paintEvent_default")
	symaddr(&addr_QScrollBar_mousePressEvent_default, "QScrollBar_mousePressEvent_default")
	symaddr(&addr_QScrollBar_mouseReleaseEvent_default, "QScrollBar_mouseReleaseEvent_default")
	symaddr(&addr_QScrollBar_mouseMoveEvent_default, "QScrollBar_mouseMoveEvent_default")
	symaddr(&addr_QScrollBar_hideEvent_default, "QScrollBar_hideEvent_default")
	symaddr(&addr_QScrollBar_sliderChange_default, "QScrollBar_sliderChange_default")
	symaddr(&addr_QScrollBar_contextMenuEvent_default, "QScrollBar_contextMenuEvent_default")
	symaddr(&addr_QScrollBar_initStyleOption_default, "QScrollBar_initStyleOption_default")
	symaddr(&addr_QScrollBar_keyPressEvent_default, "QScrollBar_keyPressEvent_default")
	symaddr(&addr_QScrollBar_timerEvent_default, "QScrollBar_timerEvent_default")
	symaddr(&addr_QScrollBar_changeEvent_default, "QScrollBar_changeEvent_default")
	symaddr(&addr_QScrollBar_devType_default, "QScrollBar_devType_default")
	symaddr(&addr_QScrollBar_setVisible_default, "QScrollBar_setVisible_default")
	symaddr(&addr_QScrollBar_minimumSizeHint_default, "QScrollBar_minimumSizeHint_default")
	symaddr(&addr_QScrollBar_heightForWidth_default, "QScrollBar_heightForWidth_default")
	symaddr(&addr_QScrollBar_hasHeightForWidth_default, "QScrollBar_hasHeightForWidth_default")
	symaddr(&addr_QScrollBar_paintEngine_default, "QScrollBar_paintEngine_default")
	symaddr(&addr_QScrollBar_mouseDoubleClickEvent_default, "QScrollBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QScrollBar_keyReleaseEvent_default, "QScrollBar_keyReleaseEvent_default")
	symaddr(&addr_QScrollBar_focusInEvent_default, "QScrollBar_focusInEvent_default")
	symaddr(&addr_QScrollBar_focusOutEvent_default, "QScrollBar_focusOutEvent_default")
	symaddr(&addr_QScrollBar_enterEvent_default, "QScrollBar_enterEvent_default")
	symaddr(&addr_QScrollBar_leaveEvent_default, "QScrollBar_leaveEvent_default")
	symaddr(&addr_QScrollBar_moveEvent_default, "QScrollBar_moveEvent_default")
	symaddr(&addr_QScrollBar_resizeEvent_default, "QScrollBar_resizeEvent_default")
	symaddr(&addr_QScrollBar_closeEvent_default, "QScrollBar_closeEvent_default")
	symaddr(&addr_QScrollBar_tabletEvent_default, "QScrollBar_tabletEvent_default")
	symaddr(&addr_QScrollBar_actionEvent_default, "QScrollBar_actionEvent_default")
	symaddr(&addr_QScrollBar_dragEnterEvent_default, "QScrollBar_dragEnterEvent_default")
	symaddr(&addr_QScrollBar_dragMoveEvent_default, "QScrollBar_dragMoveEvent_default")
	symaddr(&addr_QScrollBar_dragLeaveEvent_default, "QScrollBar_dragLeaveEvent_default")
	symaddr(&addr_QScrollBar_dropEvent_default, "QScrollBar_dropEvent_default")
	symaddr(&addr_QScrollBar_showEvent_default, "QScrollBar_showEvent_default")
	symaddr(&addr_QScrollBar_metric_default, "QScrollBar_metric_default")
	symaddr(&addr_QScrollBar_initPainter_default, "QScrollBar_initPainter_default")
	symaddr(&addr_QScrollBar_redirected_default, "QScrollBar_redirected_default")
	symaddr(&addr_QScrollBar_sharedPainter_default, "QScrollBar_sharedPainter_default")
	symaddr(&addr_QScrollBar_inputMethodEvent_default, "QScrollBar_inputMethodEvent_default")
	symaddr(&addr_QScrollBar_inputMethodQuery_default, "QScrollBar_inputMethodQuery_default")
	symaddr(&addr_QScrollBar_focusNextPrevChild_default, "QScrollBar_focusNextPrevChild_default")
	symaddr(&addr_QScrollBar_eventFilter_default, "QScrollBar_eventFilter_default")
	symaddr(&addr_QScrollBar_childEvent_default, "QScrollBar_childEvent_default")
	symaddr(&addr_QScrollBar_customEvent_default, "QScrollBar_customEvent_default")
	symaddr(&addr_QScrollBar_connectNotify_default, "QScrollBar_connectNotify_default")
	symaddr(&addr_QScrollBar_disconnectNotify_default, "QScrollBar_disconnectNotify_default")
	reg(&pg_QScrollBar_connect_valueChanged, "QScrollBar_connect_valueChanged")
	reg(&pg_QScrollBar_connect_sliderPressed, "QScrollBar_connect_sliderPressed")
	reg(&pg_QScrollBar_connect_sliderMoved, "QScrollBar_connect_sliderMoved")
	reg(&pg_QScrollBar_connect_sliderReleased, "QScrollBar_connect_sliderReleased")
	reg(&pg_QScrollBar_connect_rangeChanged, "QScrollBar_connect_rangeChanged")
	reg(&pg_QScrollBar_connect_actionTriggered, "QScrollBar_connect_actionTriggered")
	reg(&pg_QScrollBar_connect_windowTitleChanged, "QScrollBar_connect_windowTitleChanged")
	reg(&pg_QScrollBar_connect_windowIconChanged, "QScrollBar_connect_windowIconChanged")
	reg(&pg_QScrollBar_connect_windowIconTextChanged, "QScrollBar_connect_windowIconTextChanged")
	reg(&pg_QScrollBar_connect_customContextMenuRequested, "QScrollBar_connect_customContextMenuRequested")
	reg(&pg_QScrollBar_connect_destroyed, "QScrollBar_connect_destroyed")
	reg(&pg_QScrollBar_connect_destroyed2, "QScrollBar_connect_destroyed2")
	reg(&pg_QScrollBar_connect_objectNameChanged, "QScrollBar_connect_objectNameChanged")
	reg(&pg_QScrollBar__disconnect, "QScrollBar__disconnect")
	reg(&pg_QScrollBar__set_override, "QScrollBar__set_override")
	reg(&pg_QScrollBar_delete, "QScrollBar_delete")
}

type QScrollBar struct {
	h uintptr
	*QAbstractSlider
}

func bindQScrollBar(h uintptr) *QScrollBar {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QScrollBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QScrollBar{h: h, QAbstractSlider: &QAbstractSlider{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQScrollBar(h uintptr) *QScrollBar { return bindQScrollBar(h) }

func (this *QScrollBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScrollBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScrollBar(h unsafe.Pointer) *QScrollBar { return bindQScrollBar(uintptr(h)) }
func UnsafeNewQScrollBarOwned(h unsafe.Pointer) *QScrollBar { return newownedQScrollBar(uintptr(h)) }

func (this *QScrollBar) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSlider != nil {
		this.QAbstractSlider.UnsafeRelease()
	}
}

func NewQScrollBar(parent *QWidget) *QScrollBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQScrollBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollBar_virtIdx, QScrollBar_enable)
	return __o
}

func NewQScrollBar2() *QScrollBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollBar_new2, uintptr(__goid))
	__o := newownedQScrollBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollBar_virtIdx, QScrollBar_enable)
	return __o
}

func NewQScrollBar3(arg1 Orientation) *QScrollBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollBar_new3, uintptr(__goid), uintptr(arg1))
	__o := newownedQScrollBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollBar_virtIdx, QScrollBar_enable)
	return __o
}

func NewQScrollBar4(arg1 Orientation, parent *QWidget) *QScrollBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollBar_new4, uintptr(__goid), uintptr(arg1), uintptr(parent.UnsafePointer()))
	__o := newownedQScrollBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollBar_virtIdx, QScrollBar_enable)
	return __o
}

func QScrollBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QScrollBar) CreateStandardContextMenu(position *QPoint) *QMenu {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_createStandardContextMenu, this.h, uintptr(position.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	return bindQMenu(__r)
}

func (this *QScrollBar) SetRepeatAction(action QAbstractSlider__SliderAction) {
	purego.SyscallN(addr_QScrollBar_setRepeatAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) RepeatAction() QAbstractSlider__SliderAction {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_repeatAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSlider__SliderAction(__r)
}

func QScrollBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QScrollBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QScrollBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_mouseGrabber)
	return bindQWidget(__r)
}

func QScrollBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QScrollBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QScrollBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QScrollBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QScrollBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollBar) Create() {
	purego.SyscallN(addr_QScrollBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollBar) Destroy() {
	purego.SyscallN(addr_QScrollBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QScrollBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScrollBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScrollBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScrollBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QScrollBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QScrollBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QScrollBar_DevicePixelRatioFScale() float64 {
	__r := pg_QScrollBar_devicePixelRatioFScale()
	return __r
}

func QScrollBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QScrollBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QScrollBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QScrollBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QScrollBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QScrollBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QScrollBar) SetRepeatAction2(action QAbstractSlider__SliderAction, thresholdTime int) {
	purego.SyscallN(addr_QScrollBar_setRepeatAction2, this.h, uintptr(action), uintptr(thresholdTime))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) SetRepeatAction3(action QAbstractSlider__SliderAction, thresholdTime int, repeatTime int) {
	purego.SyscallN(addr_QScrollBar_setRepeatAction3, this.h, uintptr(action), uintptr(thresholdTime), uintptr(repeatTime))
	runtime.KeepAlive(this)
}

func QScrollBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QScrollBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QScrollBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QScrollBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QScrollBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QScrollBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QScrollBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QScrollBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QScrollBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QScrollBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScrollBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScrollBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScrollBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QScrollBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScrollBar) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QScrollBar) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QScrollBar_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QScrollBar_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollBar_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollBar_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollBar_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) HideEventDefault(arg1 *QHideEvent) {
	purego.SyscallN(addr_QScrollBar_hideEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	purego.SyscallN(addr_QScrollBar_sliderChange_default, this.h, uintptr(change))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QScrollBar_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) InitStyleOptionDefault(option *QStyleOptionSlider) {
	purego.SyscallN(addr_QScrollBar_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QScrollBar) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QScrollBar_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QScrollBar) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QScrollBar_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QScrollBar_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QScrollBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QScrollBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QScrollBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScrollBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QScrollBar) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QScrollBar_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QScrollBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QScrollBar_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QScrollBar_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QScrollBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QScrollBar_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QScrollBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QScrollBar_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QScrollBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QScrollBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QScrollBar_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QScrollBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QScrollBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QScrollBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QScrollBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QScrollBar_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QScrollBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QScrollBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QScrollBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QScrollBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QScrollBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QScrollBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QScrollBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollBar) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollBar_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QScrollBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QScrollBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QScrollBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QScrollBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QScrollBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QScrollBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QScrollBar) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QScrollBar_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnSliderPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScrollBar_connect_sliderPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnSliderMoved(fn func(int)) int64 {
	id := registerCallback(func(position int32) { fn(int(position)) })
	pg_QScrollBar_connect_sliderMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnSliderReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScrollBar_connect_sliderReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnRangeChanged(fn func(int, int)) int64 {
	id := registerCallback(func(min int32, max int32) { fn(int(min), int(max)) })
	pg_QScrollBar_connect_rangeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnActionTriggered(fn func(int)) int64 {
	id := registerCallback(func(action int32) { fn(int(action)) })
	pg_QScrollBar_connect_actionTriggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QScrollBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QScrollBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QScrollBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QScrollBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScrollBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QScrollBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QScrollBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollBar) Disconnect(handle int64) {
	pg_QScrollBar__disconnect(handle)
	unregisterCallback(handle)
}

func QScrollBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 1237236456312824991:
		return 2
	case 1919988340934511316:
		return 3
	case 5290442163167458465:
		return 4
	case -332090898078543319:
		return 5
	case -4415920049297795213:
		return 6
	case -4255494004099806581:
		return 7
	case -199718707876719129:
		return 8
	case 4658628703300230988:
		return 9
	case -6698043222227006015:
		return 10
	case -488693769771308531:
		return 11
	case 1023791977679793185:
		return 12
	case -5409927581812374498:
		return 13
	case 4959376063826165943:
		return 14
	case -8659679893868524044:
		return 15
	case 2110182719375245575:
		return 16
	case 7049264599837532491:
		return 17
	case 4311534060495737339:
		return 18
	case -8266762397995423915:
		return 19
	case 6398007404756346381:
		return 20
	case -4029073905697023887:
		return 21
	case 772598475950050299:
		return 22
	case 975043767033552552:
		return 23
	case -385596010576048993:
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
	case -212068866860701141:
		return 30
	case 1697444140088555581:
		return 31
	case 206122715919917589:
		return 32
	case -1781545137352745978:
		return 33
	case -3139904158373725006:
		return 34
	case 9107058682754591716:
		return 35
	case -6825369174751916:
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

func QScrollBar_enable(h uintptr, idx int32) { pg_QScrollBar__set_override(h, idx) }

func (this *QScrollBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QScrollBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QScrollBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QScrollBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QScrollBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QScrollBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QScrollBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QScrollBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QScrollBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(arg1 uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(arg1))) }) // slot "hideEvent"
}

func (this *QScrollBar) OnSliderChange(fn func(QAbstractSlider__SliderChange)) {
	registerVirt(this.UnsafeGoID(), 4658628703300230988, func(change int32) { fn(QAbstractSlider__SliderChange(change)) }) // slot "sliderChange"
}

func (this *QScrollBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QScrollBar) OnInitStyleOption(fn func(*QStyleOptionSlider)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSlider(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QScrollBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QScrollBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QScrollBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QScrollBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QScrollBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QScrollBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QScrollBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QScrollBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QScrollBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QScrollBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QScrollBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QScrollBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QScrollBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QScrollBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QScrollBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QScrollBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QScrollBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QScrollBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QScrollBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QScrollBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QScrollBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QScrollBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QScrollBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QScrollBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QScrollBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QScrollBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QScrollBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QScrollBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QScrollBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QScrollBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QScrollBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QScrollBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QScrollBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QScrollBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QScrollBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QScrollBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QScrollBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QScrollBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QScrollBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScrollBar_delete(this.h)
	}
	this.UnsafeRelease()
}
