package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLCDNumber_chain uintptr
	addr_QLCDNumber_new uintptr
	addr_QLCDNumber_new2 uintptr
	addr_QLCDNumber_new3 uintptr
	addr_QLCDNumber_new4 uintptr
	addr_QLCDNumber_tr uintptr
	addr_QLCDNumber_smallDecimalPoint uintptr
	addr_QLCDNumber_digitCount uintptr
	addr_QLCDNumber_setDigitCount uintptr
	pg_QLCDNumber_checkOverflow func(uintptr, float64) int32
	addr_QLCDNumber_checkOverflow2 uintptr
	addr_QLCDNumber_mode uintptr
	addr_QLCDNumber_setMode uintptr
	addr_QLCDNumber_segmentStyle uintptr
	addr_QLCDNumber_setSegmentStyle uintptr
	pg_QLCDNumber_value func(uintptr) float64
	addr_QLCDNumber_intValue uintptr
	addr_QLCDNumber_display uintptr
	addr_QLCDNumber_display2 uintptr
	pg_QLCDNumber_display3 func(uintptr, float64)
	addr_QLCDNumber_setHexMode uintptr
	addr_QLCDNumber_setDecMode uintptr
	addr_QLCDNumber_setOctMode uintptr
	addr_QLCDNumber_setBinMode uintptr
	addr_QLCDNumber_setSmallDecimalPoint uintptr
	addr_QLCDNumber_overflow uintptr
	addr_QLCDNumber_drawFrame uintptr
	addr_QLCDNumber_setTabOrder uintptr
	addr_QLCDNumber_mouseGrabber uintptr
	addr_QLCDNumber_keyboardGrabber uintptr
	addr_QLCDNumber_find uintptr
	addr_QLCDNumber_createWindowContainer uintptr
	addr_QLCDNumber_updateMicroFocus uintptr
	addr_QLCDNumber_create uintptr
	addr_QLCDNumber_destroy uintptr
	addr_QLCDNumber_focusNextChild uintptr
	addr_QLCDNumber_focusPreviousChild uintptr
	addr_QLCDNumber_connect uintptr
	addr_QLCDNumber_connect2 uintptr
	addr_QLCDNumber_sender uintptr
	addr_QLCDNumber_senderSignalIndex uintptr
	addr_QLCDNumber_receivers uintptr
	addr_QLCDNumber_isSignalConnected uintptr
	pg_QLCDNumber_devicePixelRatioFScale func() float64
	pg_QLCDNumber_encodeMetricF func(int32, float64) int32
	pg_QLCDNumber_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QLCDNumber_tr2 uintptr
	addr_QLCDNumber_tr3 uintptr
	addr_QLCDNumber_createWindowContainer2 uintptr
	addr_QLCDNumber_createWindowContainer3 uintptr
	addr_QLCDNumber_updateMicroFocus2 uintptr
	addr_QLCDNumber_create2 uintptr
	addr_QLCDNumber_create3 uintptr
	addr_QLCDNumber_create4 uintptr
	addr_QLCDNumber_destroy2 uintptr
	addr_QLCDNumber_destroy3 uintptr
	addr_QLCDNumber_connect3 uintptr
	addr_QLCDNumber_connect4 uintptr
	addr_QLCDNumber_metaObject_default uintptr
	addr_QLCDNumber_sizeHint_default uintptr
	addr_QLCDNumber_event_default uintptr
	addr_QLCDNumber_paintEvent_default uintptr
	addr_QLCDNumber_changeEvent_default uintptr
	addr_QLCDNumber_initStyleOption_default uintptr
	addr_QLCDNumber_devType_default uintptr
	addr_QLCDNumber_setVisible_default uintptr
	addr_QLCDNumber_minimumSizeHint_default uintptr
	addr_QLCDNumber_heightForWidth_default uintptr
	addr_QLCDNumber_hasHeightForWidth_default uintptr
	addr_QLCDNumber_paintEngine_default uintptr
	addr_QLCDNumber_mousePressEvent_default uintptr
	addr_QLCDNumber_mouseReleaseEvent_default uintptr
	addr_QLCDNumber_mouseDoubleClickEvent_default uintptr
	addr_QLCDNumber_mouseMoveEvent_default uintptr
	addr_QLCDNumber_wheelEvent_default uintptr
	addr_QLCDNumber_keyPressEvent_default uintptr
	addr_QLCDNumber_keyReleaseEvent_default uintptr
	addr_QLCDNumber_focusInEvent_default uintptr
	addr_QLCDNumber_focusOutEvent_default uintptr
	addr_QLCDNumber_enterEvent_default uintptr
	addr_QLCDNumber_leaveEvent_default uintptr
	addr_QLCDNumber_moveEvent_default uintptr
	addr_QLCDNumber_resizeEvent_default uintptr
	addr_QLCDNumber_closeEvent_default uintptr
	addr_QLCDNumber_contextMenuEvent_default uintptr
	addr_QLCDNumber_tabletEvent_default uintptr
	addr_QLCDNumber_actionEvent_default uintptr
	addr_QLCDNumber_dragEnterEvent_default uintptr
	addr_QLCDNumber_dragMoveEvent_default uintptr
	addr_QLCDNumber_dragLeaveEvent_default uintptr
	addr_QLCDNumber_dropEvent_default uintptr
	addr_QLCDNumber_showEvent_default uintptr
	addr_QLCDNumber_hideEvent_default uintptr
	addr_QLCDNumber_metric_default uintptr
	addr_QLCDNumber_initPainter_default uintptr
	addr_QLCDNumber_redirected_default uintptr
	addr_QLCDNumber_sharedPainter_default uintptr
	addr_QLCDNumber_inputMethodEvent_default uintptr
	addr_QLCDNumber_inputMethodQuery_default uintptr
	addr_QLCDNumber_focusNextPrevChild_default uintptr
	addr_QLCDNumber_eventFilter_default uintptr
	addr_QLCDNumber_timerEvent_default uintptr
	addr_QLCDNumber_childEvent_default uintptr
	addr_QLCDNumber_customEvent_default uintptr
	addr_QLCDNumber_connectNotify_default uintptr
	addr_QLCDNumber_disconnectNotify_default uintptr
	pg_QLCDNumber_connect_overflow func(uintptr, int64)
	pg_QLCDNumber_connect_windowTitleChanged func(uintptr, int64)
	pg_QLCDNumber_connect_windowIconChanged func(uintptr, int64)
	pg_QLCDNumber_connect_windowIconTextChanged func(uintptr, int64)
	pg_QLCDNumber_connect_customContextMenuRequested func(uintptr, int64)
	pg_QLCDNumber_connect_destroyed func(uintptr, int64)
	pg_QLCDNumber_connect_destroyed2 func(uintptr, int64)
	pg_QLCDNumber_connect_objectNameChanged func(uintptr, int64)
	pg_QLCDNumber__disconnect func(int64)
	pg_QLCDNumber__set_override func(uintptr, int32)
	pg_QLCDNumber_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLCDNumber_chain, "QLCDNumber_chain")
	symaddr(&addr_QLCDNumber_new, "QLCDNumber_new")
	symaddr(&addr_QLCDNumber_new2, "QLCDNumber_new2")
	symaddr(&addr_QLCDNumber_new3, "QLCDNumber_new3")
	symaddr(&addr_QLCDNumber_new4, "QLCDNumber_new4")
	symaddr(&addr_QLCDNumber_tr, "QLCDNumber_tr")
	symaddr(&addr_QLCDNumber_smallDecimalPoint, "QLCDNumber_smallDecimalPoint")
	symaddr(&addr_QLCDNumber_digitCount, "QLCDNumber_digitCount")
	symaddr(&addr_QLCDNumber_setDigitCount, "QLCDNumber_setDigitCount")
	reg(&pg_QLCDNumber_checkOverflow, "QLCDNumber_checkOverflow")
	symaddr(&addr_QLCDNumber_checkOverflow2, "QLCDNumber_checkOverflow2")
	symaddr(&addr_QLCDNumber_mode, "QLCDNumber_mode")
	symaddr(&addr_QLCDNumber_setMode, "QLCDNumber_setMode")
	symaddr(&addr_QLCDNumber_segmentStyle, "QLCDNumber_segmentStyle")
	symaddr(&addr_QLCDNumber_setSegmentStyle, "QLCDNumber_setSegmentStyle")
	reg(&pg_QLCDNumber_value, "QLCDNumber_value")
	symaddr(&addr_QLCDNumber_intValue, "QLCDNumber_intValue")
	symaddr(&addr_QLCDNumber_display, "QLCDNumber_display")
	symaddr(&addr_QLCDNumber_display2, "QLCDNumber_display2")
	reg(&pg_QLCDNumber_display3, "QLCDNumber_display3")
	symaddr(&addr_QLCDNumber_setHexMode, "QLCDNumber_setHexMode")
	symaddr(&addr_QLCDNumber_setDecMode, "QLCDNumber_setDecMode")
	symaddr(&addr_QLCDNumber_setOctMode, "QLCDNumber_setOctMode")
	symaddr(&addr_QLCDNumber_setBinMode, "QLCDNumber_setBinMode")
	symaddr(&addr_QLCDNumber_setSmallDecimalPoint, "QLCDNumber_setSmallDecimalPoint")
	symaddr(&addr_QLCDNumber_overflow, "QLCDNumber_overflow")
	symaddr(&addr_QLCDNumber_drawFrame, "QLCDNumber_drawFrame")
	symaddr(&addr_QLCDNumber_setTabOrder, "QLCDNumber_setTabOrder")
	symaddr(&addr_QLCDNumber_mouseGrabber, "QLCDNumber_mouseGrabber")
	symaddr(&addr_QLCDNumber_keyboardGrabber, "QLCDNumber_keyboardGrabber")
	symaddr(&addr_QLCDNumber_find, "QLCDNumber_find")
	symaddr(&addr_QLCDNumber_createWindowContainer, "QLCDNumber_createWindowContainer")
	symaddr(&addr_QLCDNumber_updateMicroFocus, "QLCDNumber_updateMicroFocus")
	symaddr(&addr_QLCDNumber_create, "QLCDNumber_create")
	symaddr(&addr_QLCDNumber_destroy, "QLCDNumber_destroy")
	symaddr(&addr_QLCDNumber_focusNextChild, "QLCDNumber_focusNextChild")
	symaddr(&addr_QLCDNumber_focusPreviousChild, "QLCDNumber_focusPreviousChild")
	symaddr(&addr_QLCDNumber_connect, "QLCDNumber_connect")
	symaddr(&addr_QLCDNumber_connect2, "QLCDNumber_connect2")
	symaddr(&addr_QLCDNumber_sender, "QLCDNumber_sender")
	symaddr(&addr_QLCDNumber_senderSignalIndex, "QLCDNumber_senderSignalIndex")
	symaddr(&addr_QLCDNumber_receivers, "QLCDNumber_receivers")
	symaddr(&addr_QLCDNumber_isSignalConnected, "QLCDNumber_isSignalConnected")
	reg(&pg_QLCDNumber_devicePixelRatioFScale, "QLCDNumber_devicePixelRatioFScale")
	reg(&pg_QLCDNumber_encodeMetricF, "QLCDNumber_encodeMetricF")
	reg(&pg_QLCDNumber_getDecodedMetricF, "QLCDNumber_getDecodedMetricF")
	symaddr(&addr_QLCDNumber_tr2, "QLCDNumber_tr2")
	symaddr(&addr_QLCDNumber_tr3, "QLCDNumber_tr3")
	symaddr(&addr_QLCDNumber_createWindowContainer2, "QLCDNumber_createWindowContainer2")
	symaddr(&addr_QLCDNumber_createWindowContainer3, "QLCDNumber_createWindowContainer3")
	symaddr(&addr_QLCDNumber_updateMicroFocus2, "QLCDNumber_updateMicroFocus2")
	symaddr(&addr_QLCDNumber_create2, "QLCDNumber_create2")
	symaddr(&addr_QLCDNumber_create3, "QLCDNumber_create3")
	symaddr(&addr_QLCDNumber_create4, "QLCDNumber_create4")
	symaddr(&addr_QLCDNumber_destroy2, "QLCDNumber_destroy2")
	symaddr(&addr_QLCDNumber_destroy3, "QLCDNumber_destroy3")
	symaddr(&addr_QLCDNumber_connect3, "QLCDNumber_connect3")
	symaddr(&addr_QLCDNumber_connect4, "QLCDNumber_connect4")
	symaddr(&addr_QLCDNumber_metaObject_default, "QLCDNumber_metaObject_default")
	symaddr(&addr_QLCDNumber_sizeHint_default, "QLCDNumber_sizeHint_default")
	symaddr(&addr_QLCDNumber_event_default, "QLCDNumber_event_default")
	symaddr(&addr_QLCDNumber_paintEvent_default, "QLCDNumber_paintEvent_default")
	symaddr(&addr_QLCDNumber_changeEvent_default, "QLCDNumber_changeEvent_default")
	symaddr(&addr_QLCDNumber_initStyleOption_default, "QLCDNumber_initStyleOption_default")
	symaddr(&addr_QLCDNumber_devType_default, "QLCDNumber_devType_default")
	symaddr(&addr_QLCDNumber_setVisible_default, "QLCDNumber_setVisible_default")
	symaddr(&addr_QLCDNumber_minimumSizeHint_default, "QLCDNumber_minimumSizeHint_default")
	symaddr(&addr_QLCDNumber_heightForWidth_default, "QLCDNumber_heightForWidth_default")
	symaddr(&addr_QLCDNumber_hasHeightForWidth_default, "QLCDNumber_hasHeightForWidth_default")
	symaddr(&addr_QLCDNumber_paintEngine_default, "QLCDNumber_paintEngine_default")
	symaddr(&addr_QLCDNumber_mousePressEvent_default, "QLCDNumber_mousePressEvent_default")
	symaddr(&addr_QLCDNumber_mouseReleaseEvent_default, "QLCDNumber_mouseReleaseEvent_default")
	symaddr(&addr_QLCDNumber_mouseDoubleClickEvent_default, "QLCDNumber_mouseDoubleClickEvent_default")
	symaddr(&addr_QLCDNumber_mouseMoveEvent_default, "QLCDNumber_mouseMoveEvent_default")
	symaddr(&addr_QLCDNumber_wheelEvent_default, "QLCDNumber_wheelEvent_default")
	symaddr(&addr_QLCDNumber_keyPressEvent_default, "QLCDNumber_keyPressEvent_default")
	symaddr(&addr_QLCDNumber_keyReleaseEvent_default, "QLCDNumber_keyReleaseEvent_default")
	symaddr(&addr_QLCDNumber_focusInEvent_default, "QLCDNumber_focusInEvent_default")
	symaddr(&addr_QLCDNumber_focusOutEvent_default, "QLCDNumber_focusOutEvent_default")
	symaddr(&addr_QLCDNumber_enterEvent_default, "QLCDNumber_enterEvent_default")
	symaddr(&addr_QLCDNumber_leaveEvent_default, "QLCDNumber_leaveEvent_default")
	symaddr(&addr_QLCDNumber_moveEvent_default, "QLCDNumber_moveEvent_default")
	symaddr(&addr_QLCDNumber_resizeEvent_default, "QLCDNumber_resizeEvent_default")
	symaddr(&addr_QLCDNumber_closeEvent_default, "QLCDNumber_closeEvent_default")
	symaddr(&addr_QLCDNumber_contextMenuEvent_default, "QLCDNumber_contextMenuEvent_default")
	symaddr(&addr_QLCDNumber_tabletEvent_default, "QLCDNumber_tabletEvent_default")
	symaddr(&addr_QLCDNumber_actionEvent_default, "QLCDNumber_actionEvent_default")
	symaddr(&addr_QLCDNumber_dragEnterEvent_default, "QLCDNumber_dragEnterEvent_default")
	symaddr(&addr_QLCDNumber_dragMoveEvent_default, "QLCDNumber_dragMoveEvent_default")
	symaddr(&addr_QLCDNumber_dragLeaveEvent_default, "QLCDNumber_dragLeaveEvent_default")
	symaddr(&addr_QLCDNumber_dropEvent_default, "QLCDNumber_dropEvent_default")
	symaddr(&addr_QLCDNumber_showEvent_default, "QLCDNumber_showEvent_default")
	symaddr(&addr_QLCDNumber_hideEvent_default, "QLCDNumber_hideEvent_default")
	symaddr(&addr_QLCDNumber_metric_default, "QLCDNumber_metric_default")
	symaddr(&addr_QLCDNumber_initPainter_default, "QLCDNumber_initPainter_default")
	symaddr(&addr_QLCDNumber_redirected_default, "QLCDNumber_redirected_default")
	symaddr(&addr_QLCDNumber_sharedPainter_default, "QLCDNumber_sharedPainter_default")
	symaddr(&addr_QLCDNumber_inputMethodEvent_default, "QLCDNumber_inputMethodEvent_default")
	symaddr(&addr_QLCDNumber_inputMethodQuery_default, "QLCDNumber_inputMethodQuery_default")
	symaddr(&addr_QLCDNumber_focusNextPrevChild_default, "QLCDNumber_focusNextPrevChild_default")
	symaddr(&addr_QLCDNumber_eventFilter_default, "QLCDNumber_eventFilter_default")
	symaddr(&addr_QLCDNumber_timerEvent_default, "QLCDNumber_timerEvent_default")
	symaddr(&addr_QLCDNumber_childEvent_default, "QLCDNumber_childEvent_default")
	symaddr(&addr_QLCDNumber_customEvent_default, "QLCDNumber_customEvent_default")
	symaddr(&addr_QLCDNumber_connectNotify_default, "QLCDNumber_connectNotify_default")
	symaddr(&addr_QLCDNumber_disconnectNotify_default, "QLCDNumber_disconnectNotify_default")
	reg(&pg_QLCDNumber_connect_overflow, "QLCDNumber_connect_overflow")
	reg(&pg_QLCDNumber_connect_windowTitleChanged, "QLCDNumber_connect_windowTitleChanged")
	reg(&pg_QLCDNumber_connect_windowIconChanged, "QLCDNumber_connect_windowIconChanged")
	reg(&pg_QLCDNumber_connect_windowIconTextChanged, "QLCDNumber_connect_windowIconTextChanged")
	reg(&pg_QLCDNumber_connect_customContextMenuRequested, "QLCDNumber_connect_customContextMenuRequested")
	reg(&pg_QLCDNumber_connect_destroyed, "QLCDNumber_connect_destroyed")
	reg(&pg_QLCDNumber_connect_destroyed2, "QLCDNumber_connect_destroyed2")
	reg(&pg_QLCDNumber_connect_objectNameChanged, "QLCDNumber_connect_objectNameChanged")
	reg(&pg_QLCDNumber__disconnect, "QLCDNumber__disconnect")
	reg(&pg_QLCDNumber__set_override, "QLCDNumber__set_override")
	reg(&pg_QLCDNumber_delete, "QLCDNumber_delete")
}

type QLCDNumber struct {
	h uintptr
	*QFrame
}

func bindQLCDNumber(h uintptr) *QLCDNumber {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QLCDNumber_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QLCDNumber{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQLCDNumber(h uintptr) *QLCDNumber { return bindQLCDNumber(h) }

func (this *QLCDNumber) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLCDNumber) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLCDNumber(h unsafe.Pointer) *QLCDNumber { return bindQLCDNumber(uintptr(h)) }
func UnsafeNewQLCDNumberOwned(h unsafe.Pointer) *QLCDNumber { return newownedQLCDNumber(uintptr(h)) }

func (this *QLCDNumber) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQLCDNumber(parent *QWidget) *QLCDNumber {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQLCDNumber(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLCDNumber_virtIdx, QLCDNumber_enable)
	return __o
}

func NewQLCDNumber2() *QLCDNumber {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_new2, uintptr(__goid))
	__o := newownedQLCDNumber(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLCDNumber_virtIdx, QLCDNumber_enable)
	return __o
}

func NewQLCDNumber3(numDigits uint) *QLCDNumber {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_new3, uintptr(__goid), uintptr(numDigits))
	__o := newownedQLCDNumber(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLCDNumber_virtIdx, QLCDNumber_enable)
	return __o
}

func NewQLCDNumber4(numDigits uint, parent *QWidget) *QLCDNumber {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_new4, uintptr(__goid), uintptr(numDigits), uintptr(parent.UnsafePointer()))
	__o := newownedQLCDNumber(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLCDNumber_virtIdx, QLCDNumber_enable)
	return __o
}

func QLCDNumber_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QLCDNumber) SmallDecimalPoint() bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_smallDecimalPoint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) DigitCount() int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_digitCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) SetDigitCount(nDigits int) {
	purego.SyscallN(addr_QLCDNumber_setDigitCount, this.h, uintptr(nDigits))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) CheckOverflow(num float64) bool {
	__r := pg_QLCDNumber_checkOverflow(this.h, num)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) CheckOverflowWithNum(num int) bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_checkOverflow2, this.h, uintptr(num))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) Mode() QLCDNumber__Mode {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_mode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLCDNumber__Mode(__r)
}

func (this *QLCDNumber) SetMode(arg1 QLCDNumber__Mode) {
	purego.SyscallN(addr_QLCDNumber_setMode, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SegmentStyle() QLCDNumber__SegmentStyle {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_segmentStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLCDNumber__SegmentStyle(__r)
}

func (this *QLCDNumber) SetSegmentStyle(arg1 QLCDNumber__SegmentStyle) {
	purego.SyscallN(addr_QLCDNumber_setSegmentStyle, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Value() float64 {
	__r := pg_QLCDNumber_value(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLCDNumber) IntValue() int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_intValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) Display(str string) {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	purego.SyscallN(addr_QLCDNumber_display, this.h, str_p, uintptr(len(str)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
}

func (this *QLCDNumber) DisplayWithNum(num int) {
	purego.SyscallN(addr_QLCDNumber_display2, this.h, uintptr(num))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Display2(num float64) {
	pg_QLCDNumber_display3(this.h, num)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SetHexMode() {
	purego.SyscallN(addr_QLCDNumber_setHexMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SetDecMode() {
	purego.SyscallN(addr_QLCDNumber_setDecMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SetOctMode() {
	purego.SyscallN(addr_QLCDNumber_setOctMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SetBinMode() {
	purego.SyscallN(addr_QLCDNumber_setBinMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) SetSmallDecimalPoint(arg1 bool) {
	purego.SyscallN(addr_QLCDNumber_setSmallDecimalPoint, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Overflow() {
	purego.SyscallN(addr_QLCDNumber_overflow, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QLCDNumber_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QLCDNumber_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QLCDNumber_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QLCDNumber_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_mouseGrabber)
	return bindQWidget(__r)
}

func QLCDNumber_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_keyboardGrabber)
	return bindQWidget(__r)
}

func QLCDNumber_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QLCDNumber_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QLCDNumber) UpdateMicroFocus() {
	purego.SyscallN(addr_QLCDNumber_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Create() {
	purego.SyscallN(addr_QLCDNumber_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Destroy() {
	purego.SyscallN(addr_QLCDNumber_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QLCDNumber_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLCDNumber_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLCDNumber) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QLCDNumber) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QLCDNumber) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QLCDNumber_DevicePixelRatioFScale() float64 {
	__r := pg_QLCDNumber_devicePixelRatioFScale()
	return __r
}

func QLCDNumber_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QLCDNumber_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QLCDNumber) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QLCDNumber_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QLCDNumber_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLCDNumber_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLCDNumber_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QLCDNumber_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QLCDNumber) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QLCDNumber_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QLCDNumber_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QLCDNumber_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QLCDNumber_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QLCDNumber_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QLCDNumber_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QLCDNumber_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLCDNumber_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLCDNumber) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QLCDNumber) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLCDNumber) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QLCDNumber) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QLCDNumber_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLCDNumber) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QLCDNumber_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLCDNumber) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QLCDNumber_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QLCDNumber) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QLCDNumber_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QLCDNumber) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLCDNumber) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QLCDNumber) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QLCDNumber_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QLCDNumber_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QLCDNumber_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QLCDNumber_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QLCDNumber_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QLCDNumber_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QLCDNumber_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QLCDNumber_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QLCDNumber_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QLCDNumber_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLCDNumber_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QLCDNumber_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QLCDNumber_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QLCDNumber_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QLCDNumber_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QLCDNumber_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QLCDNumber_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QLCDNumber_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QLCDNumber_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QLCDNumber_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QLCDNumber_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QLCDNumber_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QLCDNumber_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLCDNumber) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QLCDNumber_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QLCDNumber) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QLCDNumber) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QLCDNumber) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QLCDNumber_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLCDNumber) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLCDNumber_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QLCDNumber) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLCDNumber) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLCDNumber_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLCDNumber) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QLCDNumber_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QLCDNumber_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLCDNumber_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLCDNumber) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLCDNumber_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLCDNumber) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLCDNumber_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLCDNumber) OnOverflow(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLCDNumber_connect_overflow(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QLCDNumber_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QLCDNumber_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QLCDNumber_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QLCDNumber_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLCDNumber_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QLCDNumber_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QLCDNumber_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLCDNumber) Disconnect(handle int64) {
	pg_QLCDNumber__disconnect(handle)
	unregisterCallback(handle)
}

func QLCDNumber_virtIdx(slot int64) int32 {
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

func QLCDNumber_enable(h uintptr, idx int32) { pg_QLCDNumber__set_override(h, idx) }

func (this *QLCDNumber) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QLCDNumber) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QLCDNumber) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QLCDNumber) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QLCDNumber) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QLCDNumber) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QLCDNumber) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QLCDNumber) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QLCDNumber) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QLCDNumber) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QLCDNumber) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QLCDNumber) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QLCDNumber) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QLCDNumber) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QLCDNumber) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QLCDNumber) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QLCDNumber) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QLCDNumber) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QLCDNumber) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QLCDNumber) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QLCDNumber) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QLCDNumber) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QLCDNumber) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QLCDNumber) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QLCDNumber) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QLCDNumber) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QLCDNumber) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QLCDNumber) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QLCDNumber) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QLCDNumber) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QLCDNumber) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QLCDNumber) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QLCDNumber) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QLCDNumber) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QLCDNumber) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QLCDNumber) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QLCDNumber) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QLCDNumber) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QLCDNumber) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QLCDNumber) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QLCDNumber) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QLCDNumber) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QLCDNumber) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QLCDNumber) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QLCDNumber) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QLCDNumber) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QLCDNumber) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QLCDNumber) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QLCDNumber) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QLCDNumber) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLCDNumber_delete(this.h)
	}
	this.UnsafeRelease()
}
