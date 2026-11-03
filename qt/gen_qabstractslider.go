package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractSlider_chain uintptr
	addr_QAbstractSlider_new uintptr
	addr_QAbstractSlider_new2 uintptr
	addr_QAbstractSlider_tr uintptr
	addr_QAbstractSlider_orientation uintptr
	addr_QAbstractSlider_setMinimum uintptr
	addr_QAbstractSlider_minimum uintptr
	addr_QAbstractSlider_setMaximum uintptr
	addr_QAbstractSlider_maximum uintptr
	addr_QAbstractSlider_setSingleStep uintptr
	addr_QAbstractSlider_singleStep uintptr
	addr_QAbstractSlider_setPageStep uintptr
	addr_QAbstractSlider_pageStep uintptr
	addr_QAbstractSlider_setTracking uintptr
	addr_QAbstractSlider_hasTracking uintptr
	addr_QAbstractSlider_setSliderDown uintptr
	addr_QAbstractSlider_isSliderDown uintptr
	addr_QAbstractSlider_setSliderPosition uintptr
	addr_QAbstractSlider_sliderPosition uintptr
	addr_QAbstractSlider_setInvertedAppearance uintptr
	addr_QAbstractSlider_invertedAppearance uintptr
	addr_QAbstractSlider_setInvertedControls uintptr
	addr_QAbstractSlider_invertedControls uintptr
	addr_QAbstractSlider_value uintptr
	addr_QAbstractSlider_triggerAction uintptr
	addr_QAbstractSlider_setValue uintptr
	addr_QAbstractSlider_setOrientation uintptr
	addr_QAbstractSlider_setRange uintptr
	addr_QAbstractSlider_valueChanged uintptr
	addr_QAbstractSlider_sliderPressed uintptr
	addr_QAbstractSlider_sliderMoved uintptr
	addr_QAbstractSlider_sliderReleased uintptr
	addr_QAbstractSlider_rangeChanged uintptr
	addr_QAbstractSlider_actionTriggered uintptr
	addr_QAbstractSlider_setRepeatAction uintptr
	addr_QAbstractSlider_repeatAction uintptr
	addr_QAbstractSlider_setTabOrder uintptr
	addr_QAbstractSlider_mouseGrabber uintptr
	addr_QAbstractSlider_keyboardGrabber uintptr
	addr_QAbstractSlider_find uintptr
	addr_QAbstractSlider_createWindowContainer uintptr
	addr_QAbstractSlider_updateMicroFocus uintptr
	addr_QAbstractSlider_create uintptr
	addr_QAbstractSlider_destroy uintptr
	addr_QAbstractSlider_focusNextChild uintptr
	addr_QAbstractSlider_focusPreviousChild uintptr
	addr_QAbstractSlider_connect uintptr
	addr_QAbstractSlider_connect2 uintptr
	addr_QAbstractSlider_sender uintptr
	addr_QAbstractSlider_senderSignalIndex uintptr
	addr_QAbstractSlider_receivers uintptr
	addr_QAbstractSlider_isSignalConnected uintptr
	pg_QAbstractSlider_devicePixelRatioFScale func() float64
	pg_QAbstractSlider_encodeMetricF func(int32, float64) int32
	pg_QAbstractSlider_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QAbstractSlider_tr2 uintptr
	addr_QAbstractSlider_tr3 uintptr
	addr_QAbstractSlider_setRepeatAction2 uintptr
	addr_QAbstractSlider_setRepeatAction3 uintptr
	addr_QAbstractSlider_createWindowContainer2 uintptr
	addr_QAbstractSlider_createWindowContainer3 uintptr
	addr_QAbstractSlider_updateMicroFocus2 uintptr
	addr_QAbstractSlider_create2 uintptr
	addr_QAbstractSlider_create3 uintptr
	addr_QAbstractSlider_create4 uintptr
	addr_QAbstractSlider_destroy2 uintptr
	addr_QAbstractSlider_destroy3 uintptr
	addr_QAbstractSlider_connect3 uintptr
	addr_QAbstractSlider_connect4 uintptr
	addr_QAbstractSlider_metaObject_default uintptr
	addr_QAbstractSlider_event_default uintptr
	addr_QAbstractSlider_sliderChange_default uintptr
	addr_QAbstractSlider_keyPressEvent_default uintptr
	addr_QAbstractSlider_timerEvent_default uintptr
	addr_QAbstractSlider_wheelEvent_default uintptr
	addr_QAbstractSlider_changeEvent_default uintptr
	addr_QAbstractSlider_devType_default uintptr
	addr_QAbstractSlider_setVisible_default uintptr
	addr_QAbstractSlider_sizeHint_default uintptr
	addr_QAbstractSlider_minimumSizeHint_default uintptr
	addr_QAbstractSlider_heightForWidth_default uintptr
	addr_QAbstractSlider_hasHeightForWidth_default uintptr
	addr_QAbstractSlider_paintEngine_default uintptr
	addr_QAbstractSlider_mousePressEvent_default uintptr
	addr_QAbstractSlider_mouseReleaseEvent_default uintptr
	addr_QAbstractSlider_mouseDoubleClickEvent_default uintptr
	addr_QAbstractSlider_mouseMoveEvent_default uintptr
	addr_QAbstractSlider_keyReleaseEvent_default uintptr
	addr_QAbstractSlider_focusInEvent_default uintptr
	addr_QAbstractSlider_focusOutEvent_default uintptr
	addr_QAbstractSlider_enterEvent_default uintptr
	addr_QAbstractSlider_leaveEvent_default uintptr
	addr_QAbstractSlider_paintEvent_default uintptr
	addr_QAbstractSlider_moveEvent_default uintptr
	addr_QAbstractSlider_resizeEvent_default uintptr
	addr_QAbstractSlider_closeEvent_default uintptr
	addr_QAbstractSlider_contextMenuEvent_default uintptr
	addr_QAbstractSlider_tabletEvent_default uintptr
	addr_QAbstractSlider_actionEvent_default uintptr
	addr_QAbstractSlider_dragEnterEvent_default uintptr
	addr_QAbstractSlider_dragMoveEvent_default uintptr
	addr_QAbstractSlider_dragLeaveEvent_default uintptr
	addr_QAbstractSlider_dropEvent_default uintptr
	addr_QAbstractSlider_showEvent_default uintptr
	addr_QAbstractSlider_hideEvent_default uintptr
	addr_QAbstractSlider_metric_default uintptr
	addr_QAbstractSlider_initPainter_default uintptr
	addr_QAbstractSlider_redirected_default uintptr
	addr_QAbstractSlider_sharedPainter_default uintptr
	addr_QAbstractSlider_inputMethodEvent_default uintptr
	addr_QAbstractSlider_inputMethodQuery_default uintptr
	addr_QAbstractSlider_focusNextPrevChild_default uintptr
	addr_QAbstractSlider_eventFilter_default uintptr
	addr_QAbstractSlider_childEvent_default uintptr
	addr_QAbstractSlider_customEvent_default uintptr
	addr_QAbstractSlider_connectNotify_default uintptr
	addr_QAbstractSlider_disconnectNotify_default uintptr
	pg_QAbstractSlider_connect_valueChanged func(uintptr, int64)
	pg_QAbstractSlider_connect_sliderPressed func(uintptr, int64)
	pg_QAbstractSlider_connect_sliderMoved func(uintptr, int64)
	pg_QAbstractSlider_connect_sliderReleased func(uintptr, int64)
	pg_QAbstractSlider_connect_rangeChanged func(uintptr, int64)
	pg_QAbstractSlider_connect_actionTriggered func(uintptr, int64)
	pg_QAbstractSlider_connect_windowTitleChanged func(uintptr, int64)
	pg_QAbstractSlider_connect_windowIconChanged func(uintptr, int64)
	pg_QAbstractSlider_connect_windowIconTextChanged func(uintptr, int64)
	pg_QAbstractSlider_connect_customContextMenuRequested func(uintptr, int64)
	pg_QAbstractSlider_connect_destroyed func(uintptr, int64)
	pg_QAbstractSlider_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractSlider_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractSlider__disconnect func(int64)
	pg_QAbstractSlider__set_override func(uintptr, int32)
	pg_QAbstractSlider_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractSlider_chain, "QAbstractSlider_chain")
	symaddr(&addr_QAbstractSlider_new, "QAbstractSlider_new")
	symaddr(&addr_QAbstractSlider_new2, "QAbstractSlider_new2")
	symaddr(&addr_QAbstractSlider_tr, "QAbstractSlider_tr")
	symaddr(&addr_QAbstractSlider_orientation, "QAbstractSlider_orientation")
	symaddr(&addr_QAbstractSlider_setMinimum, "QAbstractSlider_setMinimum")
	symaddr(&addr_QAbstractSlider_minimum, "QAbstractSlider_minimum")
	symaddr(&addr_QAbstractSlider_setMaximum, "QAbstractSlider_setMaximum")
	symaddr(&addr_QAbstractSlider_maximum, "QAbstractSlider_maximum")
	symaddr(&addr_QAbstractSlider_setSingleStep, "QAbstractSlider_setSingleStep")
	symaddr(&addr_QAbstractSlider_singleStep, "QAbstractSlider_singleStep")
	symaddr(&addr_QAbstractSlider_setPageStep, "QAbstractSlider_setPageStep")
	symaddr(&addr_QAbstractSlider_pageStep, "QAbstractSlider_pageStep")
	symaddr(&addr_QAbstractSlider_setTracking, "QAbstractSlider_setTracking")
	symaddr(&addr_QAbstractSlider_hasTracking, "QAbstractSlider_hasTracking")
	symaddr(&addr_QAbstractSlider_setSliderDown, "QAbstractSlider_setSliderDown")
	symaddr(&addr_QAbstractSlider_isSliderDown, "QAbstractSlider_isSliderDown")
	symaddr(&addr_QAbstractSlider_setSliderPosition, "QAbstractSlider_setSliderPosition")
	symaddr(&addr_QAbstractSlider_sliderPosition, "QAbstractSlider_sliderPosition")
	symaddr(&addr_QAbstractSlider_setInvertedAppearance, "QAbstractSlider_setInvertedAppearance")
	symaddr(&addr_QAbstractSlider_invertedAppearance, "QAbstractSlider_invertedAppearance")
	symaddr(&addr_QAbstractSlider_setInvertedControls, "QAbstractSlider_setInvertedControls")
	symaddr(&addr_QAbstractSlider_invertedControls, "QAbstractSlider_invertedControls")
	symaddr(&addr_QAbstractSlider_value, "QAbstractSlider_value")
	symaddr(&addr_QAbstractSlider_triggerAction, "QAbstractSlider_triggerAction")
	symaddr(&addr_QAbstractSlider_setValue, "QAbstractSlider_setValue")
	symaddr(&addr_QAbstractSlider_setOrientation, "QAbstractSlider_setOrientation")
	symaddr(&addr_QAbstractSlider_setRange, "QAbstractSlider_setRange")
	symaddr(&addr_QAbstractSlider_valueChanged, "QAbstractSlider_valueChanged")
	symaddr(&addr_QAbstractSlider_sliderPressed, "QAbstractSlider_sliderPressed")
	symaddr(&addr_QAbstractSlider_sliderMoved, "QAbstractSlider_sliderMoved")
	symaddr(&addr_QAbstractSlider_sliderReleased, "QAbstractSlider_sliderReleased")
	symaddr(&addr_QAbstractSlider_rangeChanged, "QAbstractSlider_rangeChanged")
	symaddr(&addr_QAbstractSlider_actionTriggered, "QAbstractSlider_actionTriggered")
	symaddr(&addr_QAbstractSlider_setRepeatAction, "QAbstractSlider_setRepeatAction")
	symaddr(&addr_QAbstractSlider_repeatAction, "QAbstractSlider_repeatAction")
	symaddr(&addr_QAbstractSlider_setTabOrder, "QAbstractSlider_setTabOrder")
	symaddr(&addr_QAbstractSlider_mouseGrabber, "QAbstractSlider_mouseGrabber")
	symaddr(&addr_QAbstractSlider_keyboardGrabber, "QAbstractSlider_keyboardGrabber")
	symaddr(&addr_QAbstractSlider_find, "QAbstractSlider_find")
	symaddr(&addr_QAbstractSlider_createWindowContainer, "QAbstractSlider_createWindowContainer")
	symaddr(&addr_QAbstractSlider_updateMicroFocus, "QAbstractSlider_updateMicroFocus")
	symaddr(&addr_QAbstractSlider_create, "QAbstractSlider_create")
	symaddr(&addr_QAbstractSlider_destroy, "QAbstractSlider_destroy")
	symaddr(&addr_QAbstractSlider_focusNextChild, "QAbstractSlider_focusNextChild")
	symaddr(&addr_QAbstractSlider_focusPreviousChild, "QAbstractSlider_focusPreviousChild")
	symaddr(&addr_QAbstractSlider_connect, "QAbstractSlider_connect")
	symaddr(&addr_QAbstractSlider_connect2, "QAbstractSlider_connect2")
	symaddr(&addr_QAbstractSlider_sender, "QAbstractSlider_sender")
	symaddr(&addr_QAbstractSlider_senderSignalIndex, "QAbstractSlider_senderSignalIndex")
	symaddr(&addr_QAbstractSlider_receivers, "QAbstractSlider_receivers")
	symaddr(&addr_QAbstractSlider_isSignalConnected, "QAbstractSlider_isSignalConnected")
	reg(&pg_QAbstractSlider_devicePixelRatioFScale, "QAbstractSlider_devicePixelRatioFScale")
	reg(&pg_QAbstractSlider_encodeMetricF, "QAbstractSlider_encodeMetricF")
	reg(&pg_QAbstractSlider_getDecodedMetricF, "QAbstractSlider_getDecodedMetricF")
	symaddr(&addr_QAbstractSlider_tr2, "QAbstractSlider_tr2")
	symaddr(&addr_QAbstractSlider_tr3, "QAbstractSlider_tr3")
	symaddr(&addr_QAbstractSlider_setRepeatAction2, "QAbstractSlider_setRepeatAction2")
	symaddr(&addr_QAbstractSlider_setRepeatAction3, "QAbstractSlider_setRepeatAction3")
	symaddr(&addr_QAbstractSlider_createWindowContainer2, "QAbstractSlider_createWindowContainer2")
	symaddr(&addr_QAbstractSlider_createWindowContainer3, "QAbstractSlider_createWindowContainer3")
	symaddr(&addr_QAbstractSlider_updateMicroFocus2, "QAbstractSlider_updateMicroFocus2")
	symaddr(&addr_QAbstractSlider_create2, "QAbstractSlider_create2")
	symaddr(&addr_QAbstractSlider_create3, "QAbstractSlider_create3")
	symaddr(&addr_QAbstractSlider_create4, "QAbstractSlider_create4")
	symaddr(&addr_QAbstractSlider_destroy2, "QAbstractSlider_destroy2")
	symaddr(&addr_QAbstractSlider_destroy3, "QAbstractSlider_destroy3")
	symaddr(&addr_QAbstractSlider_connect3, "QAbstractSlider_connect3")
	symaddr(&addr_QAbstractSlider_connect4, "QAbstractSlider_connect4")
	symaddr(&addr_QAbstractSlider_metaObject_default, "QAbstractSlider_metaObject_default")
	symaddr(&addr_QAbstractSlider_event_default, "QAbstractSlider_event_default")
	symaddr(&addr_QAbstractSlider_sliderChange_default, "QAbstractSlider_sliderChange_default")
	symaddr(&addr_QAbstractSlider_keyPressEvent_default, "QAbstractSlider_keyPressEvent_default")
	symaddr(&addr_QAbstractSlider_timerEvent_default, "QAbstractSlider_timerEvent_default")
	symaddr(&addr_QAbstractSlider_wheelEvent_default, "QAbstractSlider_wheelEvent_default")
	symaddr(&addr_QAbstractSlider_changeEvent_default, "QAbstractSlider_changeEvent_default")
	symaddr(&addr_QAbstractSlider_devType_default, "QAbstractSlider_devType_default")
	symaddr(&addr_QAbstractSlider_setVisible_default, "QAbstractSlider_setVisible_default")
	symaddr(&addr_QAbstractSlider_sizeHint_default, "QAbstractSlider_sizeHint_default")
	symaddr(&addr_QAbstractSlider_minimumSizeHint_default, "QAbstractSlider_minimumSizeHint_default")
	symaddr(&addr_QAbstractSlider_heightForWidth_default, "QAbstractSlider_heightForWidth_default")
	symaddr(&addr_QAbstractSlider_hasHeightForWidth_default, "QAbstractSlider_hasHeightForWidth_default")
	symaddr(&addr_QAbstractSlider_paintEngine_default, "QAbstractSlider_paintEngine_default")
	symaddr(&addr_QAbstractSlider_mousePressEvent_default, "QAbstractSlider_mousePressEvent_default")
	symaddr(&addr_QAbstractSlider_mouseReleaseEvent_default, "QAbstractSlider_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractSlider_mouseDoubleClickEvent_default, "QAbstractSlider_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractSlider_mouseMoveEvent_default, "QAbstractSlider_mouseMoveEvent_default")
	symaddr(&addr_QAbstractSlider_keyReleaseEvent_default, "QAbstractSlider_keyReleaseEvent_default")
	symaddr(&addr_QAbstractSlider_focusInEvent_default, "QAbstractSlider_focusInEvent_default")
	symaddr(&addr_QAbstractSlider_focusOutEvent_default, "QAbstractSlider_focusOutEvent_default")
	symaddr(&addr_QAbstractSlider_enterEvent_default, "QAbstractSlider_enterEvent_default")
	symaddr(&addr_QAbstractSlider_leaveEvent_default, "QAbstractSlider_leaveEvent_default")
	symaddr(&addr_QAbstractSlider_paintEvent_default, "QAbstractSlider_paintEvent_default")
	symaddr(&addr_QAbstractSlider_moveEvent_default, "QAbstractSlider_moveEvent_default")
	symaddr(&addr_QAbstractSlider_resizeEvent_default, "QAbstractSlider_resizeEvent_default")
	symaddr(&addr_QAbstractSlider_closeEvent_default, "QAbstractSlider_closeEvent_default")
	symaddr(&addr_QAbstractSlider_contextMenuEvent_default, "QAbstractSlider_contextMenuEvent_default")
	symaddr(&addr_QAbstractSlider_tabletEvent_default, "QAbstractSlider_tabletEvent_default")
	symaddr(&addr_QAbstractSlider_actionEvent_default, "QAbstractSlider_actionEvent_default")
	symaddr(&addr_QAbstractSlider_dragEnterEvent_default, "QAbstractSlider_dragEnterEvent_default")
	symaddr(&addr_QAbstractSlider_dragMoveEvent_default, "QAbstractSlider_dragMoveEvent_default")
	symaddr(&addr_QAbstractSlider_dragLeaveEvent_default, "QAbstractSlider_dragLeaveEvent_default")
	symaddr(&addr_QAbstractSlider_dropEvent_default, "QAbstractSlider_dropEvent_default")
	symaddr(&addr_QAbstractSlider_showEvent_default, "QAbstractSlider_showEvent_default")
	symaddr(&addr_QAbstractSlider_hideEvent_default, "QAbstractSlider_hideEvent_default")
	symaddr(&addr_QAbstractSlider_metric_default, "QAbstractSlider_metric_default")
	symaddr(&addr_QAbstractSlider_initPainter_default, "QAbstractSlider_initPainter_default")
	symaddr(&addr_QAbstractSlider_redirected_default, "QAbstractSlider_redirected_default")
	symaddr(&addr_QAbstractSlider_sharedPainter_default, "QAbstractSlider_sharedPainter_default")
	symaddr(&addr_QAbstractSlider_inputMethodEvent_default, "QAbstractSlider_inputMethodEvent_default")
	symaddr(&addr_QAbstractSlider_inputMethodQuery_default, "QAbstractSlider_inputMethodQuery_default")
	symaddr(&addr_QAbstractSlider_focusNextPrevChild_default, "QAbstractSlider_focusNextPrevChild_default")
	symaddr(&addr_QAbstractSlider_eventFilter_default, "QAbstractSlider_eventFilter_default")
	symaddr(&addr_QAbstractSlider_childEvent_default, "QAbstractSlider_childEvent_default")
	symaddr(&addr_QAbstractSlider_customEvent_default, "QAbstractSlider_customEvent_default")
	symaddr(&addr_QAbstractSlider_connectNotify_default, "QAbstractSlider_connectNotify_default")
	symaddr(&addr_QAbstractSlider_disconnectNotify_default, "QAbstractSlider_disconnectNotify_default")
	reg(&pg_QAbstractSlider_connect_valueChanged, "QAbstractSlider_connect_valueChanged")
	reg(&pg_QAbstractSlider_connect_sliderPressed, "QAbstractSlider_connect_sliderPressed")
	reg(&pg_QAbstractSlider_connect_sliderMoved, "QAbstractSlider_connect_sliderMoved")
	reg(&pg_QAbstractSlider_connect_sliderReleased, "QAbstractSlider_connect_sliderReleased")
	reg(&pg_QAbstractSlider_connect_rangeChanged, "QAbstractSlider_connect_rangeChanged")
	reg(&pg_QAbstractSlider_connect_actionTriggered, "QAbstractSlider_connect_actionTriggered")
	reg(&pg_QAbstractSlider_connect_windowTitleChanged, "QAbstractSlider_connect_windowTitleChanged")
	reg(&pg_QAbstractSlider_connect_windowIconChanged, "QAbstractSlider_connect_windowIconChanged")
	reg(&pg_QAbstractSlider_connect_windowIconTextChanged, "QAbstractSlider_connect_windowIconTextChanged")
	reg(&pg_QAbstractSlider_connect_customContextMenuRequested, "QAbstractSlider_connect_customContextMenuRequested")
	reg(&pg_QAbstractSlider_connect_destroyed, "QAbstractSlider_connect_destroyed")
	reg(&pg_QAbstractSlider_connect_destroyed2, "QAbstractSlider_connect_destroyed2")
	reg(&pg_QAbstractSlider_connect_objectNameChanged, "QAbstractSlider_connect_objectNameChanged")
	reg(&pg_QAbstractSlider__disconnect, "QAbstractSlider__disconnect")
	reg(&pg_QAbstractSlider__set_override, "QAbstractSlider__set_override")
	reg(&pg_QAbstractSlider_delete, "QAbstractSlider_delete")
}

type QAbstractSlider struct {
	h uintptr
	*QWidget
}

func bindQAbstractSlider(h uintptr) *QAbstractSlider {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractSlider_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractSlider{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractSlider(h uintptr) *QAbstractSlider { return bindQAbstractSlider(h) }

func (this *QAbstractSlider) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractSlider) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractSlider(h unsafe.Pointer) *QAbstractSlider { return bindQAbstractSlider(uintptr(h)) }
func UnsafeNewQAbstractSliderOwned(h unsafe.Pointer) *QAbstractSlider { return newownedQAbstractSlider(uintptr(h)) }

func (this *QAbstractSlider) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQAbstractSlider(parent *QWidget) *QAbstractSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAbstractSlider(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractSlider_virtIdx, QAbstractSlider_enable)
	return __o
}

func NewQAbstractSlider2() *QAbstractSlider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_new2, uintptr(__goid))
	__o := newownedQAbstractSlider(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractSlider_virtIdx, QAbstractSlider_enable)
	return __o
}

func QAbstractSlider_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractSlider) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QAbstractSlider) SetMinimum(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setMinimum, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_minimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetMaximum(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setMaximum, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_maximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetSingleStep(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setSingleStep, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SingleStep() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_singleStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetPageStep(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setPageStep, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) PageStep() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_pageStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetTracking(enable bool) {
	purego.SyscallN(addr_QAbstractSlider_setTracking, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) HasTracking() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_hasTracking, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) SetSliderDown(arg1 bool) {
	purego.SyscallN(addr_QAbstractSlider_setSliderDown, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) IsSliderDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_isSliderDown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) SetSliderPosition(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setSliderPosition, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SliderPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_sliderPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetInvertedAppearance(arg1 bool) {
	purego.SyscallN(addr_QAbstractSlider_setInvertedAppearance, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) InvertedAppearance() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_invertedAppearance, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) SetInvertedControls(arg1 bool) {
	purego.SyscallN(addr_QAbstractSlider_setInvertedControls, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) InvertedControls() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_invertedControls, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) Value() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) TriggerAction(action QAbstractSlider__SliderAction) {
	purego.SyscallN(addr_QAbstractSlider_triggerAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SetValue(arg1 int) {
	purego.SyscallN(addr_QAbstractSlider_setValue, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SetOrientation(arg1 Orientation) {
	purego.SyscallN(addr_QAbstractSlider_setOrientation, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SetRange(min int, max int) {
	purego.SyscallN(addr_QAbstractSlider_setRange, this.h, uintptr(min), uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) ValueChanged(value int) {
	purego.SyscallN(addr_QAbstractSlider_valueChanged, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SliderPressed() {
	purego.SyscallN(addr_QAbstractSlider_sliderPressed, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SliderMoved(position int) {
	purego.SyscallN(addr_QAbstractSlider_sliderMoved, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SliderReleased() {
	purego.SyscallN(addr_QAbstractSlider_sliderReleased, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) RangeChanged(min int, max int) {
	purego.SyscallN(addr_QAbstractSlider_rangeChanged, this.h, uintptr(min), uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) ActionTriggered(action int) {
	purego.SyscallN(addr_QAbstractSlider_actionTriggered, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SetRepeatAction(action QAbstractSlider__SliderAction) {
	purego.SyscallN(addr_QAbstractSlider_setRepeatAction, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) RepeatAction() QAbstractSlider__SliderAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_repeatAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSlider__SliderAction(__r)
}

func QAbstractSlider_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QAbstractSlider_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QAbstractSlider_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_mouseGrabber)
	return bindQWidget(__r)
}

func QAbstractSlider_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_keyboardGrabber)
	return bindQWidget(__r)
}

func QAbstractSlider_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QAbstractSlider_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QAbstractSlider) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractSlider_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Create() {
	purego.SyscallN(addr_QAbstractSlider_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Destroy() {
	purego.SyscallN(addr_QAbstractSlider_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAbstractSlider_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractSlider_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractSlider) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractSlider) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractSlider) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractSlider_DevicePixelRatioFScale() float64 {
	__r := pg_QAbstractSlider_devicePixelRatioFScale()
	return __r
}

func QAbstractSlider_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QAbstractSlider_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QAbstractSlider) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QAbstractSlider_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QAbstractSlider_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractSlider_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractSlider) SetRepeatAction2(action QAbstractSlider__SliderAction, thresholdTime int) {
	purego.SyscallN(addr_QAbstractSlider_setRepeatAction2, this.h, uintptr(action), uintptr(thresholdTime))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SetRepeatAction3(action QAbstractSlider__SliderAction, thresholdTime int, repeatTime int) {
	purego.SyscallN(addr_QAbstractSlider_setRepeatAction3, this.h, uintptr(action), uintptr(thresholdTime), uintptr(repeatTime))
	runtime.KeepAlive(this)
}

func QAbstractSlider_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QAbstractSlider_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QAbstractSlider) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QAbstractSlider_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QAbstractSlider_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QAbstractSlider_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QAbstractSlider_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QAbstractSlider_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QAbstractSlider_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QAbstractSlider_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractSlider_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractSlider) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractSlider) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QAbstractSlider) SliderChangeDefault(change QAbstractSlider__SliderChange) {
	purego.SyscallN(addr_QAbstractSlider_sliderChange_default, this.h, uintptr(change))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QAbstractSlider_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QAbstractSlider) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QAbstractSlider_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractSlider) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QAbstractSlider_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractSlider) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QAbstractSlider_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractSlider) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QAbstractSlider_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSlider) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractSlider) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractSlider) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QAbstractSlider) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSlider_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSlider_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSlider_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSlider_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractSlider_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractSlider_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractSlider_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QAbstractSlider_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractSlider_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QAbstractSlider_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QAbstractSlider_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QAbstractSlider_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QAbstractSlider_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QAbstractSlider_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QAbstractSlider_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QAbstractSlider_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QAbstractSlider_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QAbstractSlider_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QAbstractSlider_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QAbstractSlider_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QAbstractSlider_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QAbstractSlider_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSlider) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QAbstractSlider_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QAbstractSlider) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QAbstractSlider) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QAbstractSlider) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractSlider_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractSlider) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractSlider_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractSlider) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSlider) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSlider_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractSlider) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractSlider_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractSlider_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSlider) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractSlider_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractSlider) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractSlider_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractSlider) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QAbstractSlider_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnSliderPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSlider_connect_sliderPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnSliderMoved(fn func(int)) int64 {
	id := registerCallback(func(position int32) { fn(int(position)) })
	pg_QAbstractSlider_connect_sliderMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnSliderReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSlider_connect_sliderReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnRangeChanged(fn func(int, int)) int64 {
	id := registerCallback(func(min int32, max int32) { fn(int(min), int(max)) })
	pg_QAbstractSlider_connect_rangeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnActionTriggered(fn func(int)) int64 {
	id := registerCallback(func(action int32) { fn(int(action)) })
	pg_QAbstractSlider_connect_actionTriggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QAbstractSlider_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QAbstractSlider_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QAbstractSlider_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QAbstractSlider_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSlider_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractSlider_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractSlider_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSlider) Disconnect(handle int64) {
	pg_QAbstractSlider__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractSlider_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 4658628703300230988:
		return 2
	case 1023791977679793185:
		return 3
	case -5409927581812374498:
		return 4
	case 1919988340934511316:
		return 5
	case 4959376063826165943:
		return 6
	case -8659679893868524044:
		return 7
	case 2110182719375245575:
		return 8
	case 729690261457684101:
		return 9
	case 7049264599837532491:
		return 10
	case 4311534060495737339:
		return 11
	case -8266762397995423915:
		return 12
	case 6398007404756346381:
		return 13
	case -332090898078543319:
		return 14
	case -4415920049297795213:
		return 15
	case -4029073905697023887:
		return 16
	case -4255494004099806581:
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
	case 5290442163167458465:
		return 23
	case -9161949150489525396:
		return 24
	case -4302917285655014661:
		return 25
	case -1989937050019464783:
		return 26
	case -6698043222227006015:
		return 27
	case -212068866860701141:
		return 28
	case 1697444140088555581:
		return 29
	case 206122715919917589:
		return 30
	case -1781545137352745978:
		return 31
	case -3139904158373725006:
		return 32
	case 9107058682754591716:
		return 33
	case -6825369174751916:
		return 34
	case -199718707876719129:
		return 35
	case -2039166586246443706:
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
	case 5217623514630858297:
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

func QAbstractSlider_enable(h uintptr, idx int32) { pg_QAbstractSlider__set_override(h, idx) }

func (this *QAbstractSlider) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractSlider) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QAbstractSlider) OnSliderChange(fn func(QAbstractSlider__SliderChange)) {
	registerVirt(this.UnsafeGoID(), 4658628703300230988, func(change int32) { fn(QAbstractSlider__SliderChange(change)) }) // slot "sliderChange"
}

func (this *QAbstractSlider) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QAbstractSlider) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QAbstractSlider) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QAbstractSlider) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QAbstractSlider) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QAbstractSlider) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QAbstractSlider) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractSlider) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QAbstractSlider) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QAbstractSlider) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QAbstractSlider) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QAbstractSlider) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QAbstractSlider) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractSlider) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractSlider) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractSlider) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractSlider) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QAbstractSlider) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QAbstractSlider) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QAbstractSlider) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QAbstractSlider) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QAbstractSlider) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QAbstractSlider) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QAbstractSlider) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QAbstractSlider) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QAbstractSlider) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QAbstractSlider) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QAbstractSlider) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QAbstractSlider) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QAbstractSlider) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractSlider) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QAbstractSlider) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QAbstractSlider) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QAbstractSlider) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QAbstractSlider) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QAbstractSlider) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QAbstractSlider) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QAbstractSlider) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QAbstractSlider) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QAbstractSlider) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractSlider) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QAbstractSlider) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractSlider) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractSlider) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractSlider) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractSlider) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractSlider) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractSlider_delete(this.h)
	}
	this.UnsafeRelease()
}
