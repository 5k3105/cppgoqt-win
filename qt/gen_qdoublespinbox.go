package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDoubleSpinBox_chain uintptr
	addr_QDoubleSpinBox_new uintptr
	addr_QDoubleSpinBox_new2 uintptr
	addr_QDoubleSpinBox_tr uintptr
	pg_QDoubleSpinBox_value func(uintptr) float64
	addr_QDoubleSpinBox_prefix uintptr
	addr_QDoubleSpinBox_setPrefix uintptr
	addr_QDoubleSpinBox_suffix uintptr
	addr_QDoubleSpinBox_setSuffix uintptr
	addr_QDoubleSpinBox_cleanText uintptr
	pg_QDoubleSpinBox_singleStep func(uintptr) float64
	pg_QDoubleSpinBox_setSingleStep func(uintptr, float64)
	pg_QDoubleSpinBox_minimum func(uintptr) float64
	pg_QDoubleSpinBox_setMinimum func(uintptr, float64)
	pg_QDoubleSpinBox_maximum func(uintptr) float64
	pg_QDoubleSpinBox_setMaximum func(uintptr, float64)
	pg_QDoubleSpinBox_setRange func(uintptr, float64, float64)
	addr_QDoubleSpinBox_stepType uintptr
	addr_QDoubleSpinBox_setStepType uintptr
	addr_QDoubleSpinBox_decimals uintptr
	addr_QDoubleSpinBox_setDecimals uintptr
	pg_QDoubleSpinBox_valueFromText func(uintptr, string, int64) float64
	pg_QDoubleSpinBox_textFromValue func(uintptr, float64, *int64) uintptr
	pg_QDoubleSpinBox_setValue func(uintptr, float64)
	pg_QDoubleSpinBox_valueChanged func(uintptr, float64)
	addr_QDoubleSpinBox_textChanged uintptr
	addr_QDoubleSpinBox_lineEdit uintptr
	addr_QDoubleSpinBox_setLineEdit uintptr
	addr_QDoubleSpinBox_setTabOrder uintptr
	addr_QDoubleSpinBox_mouseGrabber uintptr
	addr_QDoubleSpinBox_keyboardGrabber uintptr
	addr_QDoubleSpinBox_find uintptr
	addr_QDoubleSpinBox_createWindowContainer uintptr
	addr_QDoubleSpinBox_updateMicroFocus uintptr
	addr_QDoubleSpinBox_create uintptr
	addr_QDoubleSpinBox_destroy uintptr
	addr_QDoubleSpinBox_focusNextChild uintptr
	addr_QDoubleSpinBox_focusPreviousChild uintptr
	addr_QDoubleSpinBox_connect uintptr
	addr_QDoubleSpinBox_connect2 uintptr
	addr_QDoubleSpinBox_sender uintptr
	addr_QDoubleSpinBox_senderSignalIndex uintptr
	addr_QDoubleSpinBox_receivers uintptr
	addr_QDoubleSpinBox_isSignalConnected uintptr
	pg_QDoubleSpinBox_devicePixelRatioFScale func() float64
	pg_QDoubleSpinBox_encodeMetricF func(int32, float64) int32
	pg_QDoubleSpinBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDoubleSpinBox_tr2 uintptr
	addr_QDoubleSpinBox_tr3 uintptr
	addr_QDoubleSpinBox_createWindowContainer2 uintptr
	addr_QDoubleSpinBox_createWindowContainer3 uintptr
	addr_QDoubleSpinBox_updateMicroFocus2 uintptr
	addr_QDoubleSpinBox_create2 uintptr
	addr_QDoubleSpinBox_create3 uintptr
	addr_QDoubleSpinBox_create4 uintptr
	addr_QDoubleSpinBox_destroy2 uintptr
	addr_QDoubleSpinBox_destroy3 uintptr
	addr_QDoubleSpinBox_connect3 uintptr
	addr_QDoubleSpinBox_connect4 uintptr
	pg_QDoubleSpinBox_valueFromText_default func(uintptr, string, int64) float64
	pg_QDoubleSpinBox_textFromValue_default func(uintptr, float64, *int64) uintptr
	addr_QDoubleSpinBox_metaObject_default uintptr
	addr_QDoubleSpinBox_validate_default uintptr
	addr_QDoubleSpinBox_fixup_default uintptr
	addr_QDoubleSpinBox_sizeHint_default uintptr
	addr_QDoubleSpinBox_minimumSizeHint_default uintptr
	addr_QDoubleSpinBox_event_default uintptr
	addr_QDoubleSpinBox_inputMethodQuery_default uintptr
	addr_QDoubleSpinBox_stepBy_default uintptr
	addr_QDoubleSpinBox_clear_default uintptr
	addr_QDoubleSpinBox_resizeEvent_default uintptr
	addr_QDoubleSpinBox_keyPressEvent_default uintptr
	addr_QDoubleSpinBox_keyReleaseEvent_default uintptr
	addr_QDoubleSpinBox_wheelEvent_default uintptr
	addr_QDoubleSpinBox_focusInEvent_default uintptr
	addr_QDoubleSpinBox_focusOutEvent_default uintptr
	addr_QDoubleSpinBox_contextMenuEvent_default uintptr
	addr_QDoubleSpinBox_changeEvent_default uintptr
	addr_QDoubleSpinBox_closeEvent_default uintptr
	addr_QDoubleSpinBox_hideEvent_default uintptr
	addr_QDoubleSpinBox_mousePressEvent_default uintptr
	addr_QDoubleSpinBox_mouseReleaseEvent_default uintptr
	addr_QDoubleSpinBox_mouseMoveEvent_default uintptr
	addr_QDoubleSpinBox_timerEvent_default uintptr
	addr_QDoubleSpinBox_paintEvent_default uintptr
	addr_QDoubleSpinBox_showEvent_default uintptr
	addr_QDoubleSpinBox_initStyleOption_default uintptr
	addr_QDoubleSpinBox_stepEnabled_default uintptr
	addr_QDoubleSpinBox_devType_default uintptr
	addr_QDoubleSpinBox_setVisible_default uintptr
	addr_QDoubleSpinBox_heightForWidth_default uintptr
	addr_QDoubleSpinBox_hasHeightForWidth_default uintptr
	addr_QDoubleSpinBox_paintEngine_default uintptr
	addr_QDoubleSpinBox_mouseDoubleClickEvent_default uintptr
	addr_QDoubleSpinBox_enterEvent_default uintptr
	addr_QDoubleSpinBox_leaveEvent_default uintptr
	addr_QDoubleSpinBox_moveEvent_default uintptr
	addr_QDoubleSpinBox_tabletEvent_default uintptr
	addr_QDoubleSpinBox_actionEvent_default uintptr
	addr_QDoubleSpinBox_dragEnterEvent_default uintptr
	addr_QDoubleSpinBox_dragMoveEvent_default uintptr
	addr_QDoubleSpinBox_dragLeaveEvent_default uintptr
	addr_QDoubleSpinBox_dropEvent_default uintptr
	addr_QDoubleSpinBox_metric_default uintptr
	addr_QDoubleSpinBox_initPainter_default uintptr
	addr_QDoubleSpinBox_redirected_default uintptr
	addr_QDoubleSpinBox_sharedPainter_default uintptr
	addr_QDoubleSpinBox_inputMethodEvent_default uintptr
	addr_QDoubleSpinBox_focusNextPrevChild_default uintptr
	addr_QDoubleSpinBox_eventFilter_default uintptr
	addr_QDoubleSpinBox_childEvent_default uintptr
	addr_QDoubleSpinBox_customEvent_default uintptr
	addr_QDoubleSpinBox_connectNotify_default uintptr
	addr_QDoubleSpinBox_disconnectNotify_default uintptr
	pg_QDoubleSpinBox_connect_valueChanged func(uintptr, int64)
	pg_QDoubleSpinBox_connect_textChanged func(uintptr, int64)
	pg_QDoubleSpinBox_connect_editingFinished func(uintptr, int64)
	pg_QDoubleSpinBox_connect_returnPressed func(uintptr, int64)
	pg_QDoubleSpinBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QDoubleSpinBox_connect_windowIconChanged func(uintptr, int64)
	pg_QDoubleSpinBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDoubleSpinBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDoubleSpinBox_connect_destroyed func(uintptr, int64)
	pg_QDoubleSpinBox_connect_destroyed2 func(uintptr, int64)
	pg_QDoubleSpinBox_connect_objectNameChanged func(uintptr, int64)
	pg_QDoubleSpinBox__disconnect func(int64)
	pg_QDoubleSpinBox__set_override func(uintptr, int32)
	pg_QDoubleSpinBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDoubleSpinBox_chain, "QDoubleSpinBox_chain")
	symaddr(&addr_QDoubleSpinBox_new, "QDoubleSpinBox_new")
	symaddr(&addr_QDoubleSpinBox_new2, "QDoubleSpinBox_new2")
	symaddr(&addr_QDoubleSpinBox_tr, "QDoubleSpinBox_tr")
	reg(&pg_QDoubleSpinBox_value, "QDoubleSpinBox_value")
	symaddr(&addr_QDoubleSpinBox_prefix, "QDoubleSpinBox_prefix")
	symaddr(&addr_QDoubleSpinBox_setPrefix, "QDoubleSpinBox_setPrefix")
	symaddr(&addr_QDoubleSpinBox_suffix, "QDoubleSpinBox_suffix")
	symaddr(&addr_QDoubleSpinBox_setSuffix, "QDoubleSpinBox_setSuffix")
	symaddr(&addr_QDoubleSpinBox_cleanText, "QDoubleSpinBox_cleanText")
	reg(&pg_QDoubleSpinBox_singleStep, "QDoubleSpinBox_singleStep")
	reg(&pg_QDoubleSpinBox_setSingleStep, "QDoubleSpinBox_setSingleStep")
	reg(&pg_QDoubleSpinBox_minimum, "QDoubleSpinBox_minimum")
	reg(&pg_QDoubleSpinBox_setMinimum, "QDoubleSpinBox_setMinimum")
	reg(&pg_QDoubleSpinBox_maximum, "QDoubleSpinBox_maximum")
	reg(&pg_QDoubleSpinBox_setMaximum, "QDoubleSpinBox_setMaximum")
	reg(&pg_QDoubleSpinBox_setRange, "QDoubleSpinBox_setRange")
	symaddr(&addr_QDoubleSpinBox_stepType, "QDoubleSpinBox_stepType")
	symaddr(&addr_QDoubleSpinBox_setStepType, "QDoubleSpinBox_setStepType")
	symaddr(&addr_QDoubleSpinBox_decimals, "QDoubleSpinBox_decimals")
	symaddr(&addr_QDoubleSpinBox_setDecimals, "QDoubleSpinBox_setDecimals")
	reg(&pg_QDoubleSpinBox_valueFromText, "QDoubleSpinBox_valueFromText")
	reg(&pg_QDoubleSpinBox_textFromValue, "QDoubleSpinBox_textFromValue")
	reg(&pg_QDoubleSpinBox_setValue, "QDoubleSpinBox_setValue")
	reg(&pg_QDoubleSpinBox_valueChanged, "QDoubleSpinBox_valueChanged")
	symaddr(&addr_QDoubleSpinBox_textChanged, "QDoubleSpinBox_textChanged")
	symaddr(&addr_QDoubleSpinBox_lineEdit, "QDoubleSpinBox_lineEdit")
	symaddr(&addr_QDoubleSpinBox_setLineEdit, "QDoubleSpinBox_setLineEdit")
	symaddr(&addr_QDoubleSpinBox_setTabOrder, "QDoubleSpinBox_setTabOrder")
	symaddr(&addr_QDoubleSpinBox_mouseGrabber, "QDoubleSpinBox_mouseGrabber")
	symaddr(&addr_QDoubleSpinBox_keyboardGrabber, "QDoubleSpinBox_keyboardGrabber")
	symaddr(&addr_QDoubleSpinBox_find, "QDoubleSpinBox_find")
	symaddr(&addr_QDoubleSpinBox_createWindowContainer, "QDoubleSpinBox_createWindowContainer")
	symaddr(&addr_QDoubleSpinBox_updateMicroFocus, "QDoubleSpinBox_updateMicroFocus")
	symaddr(&addr_QDoubleSpinBox_create, "QDoubleSpinBox_create")
	symaddr(&addr_QDoubleSpinBox_destroy, "QDoubleSpinBox_destroy")
	symaddr(&addr_QDoubleSpinBox_focusNextChild, "QDoubleSpinBox_focusNextChild")
	symaddr(&addr_QDoubleSpinBox_focusPreviousChild, "QDoubleSpinBox_focusPreviousChild")
	symaddr(&addr_QDoubleSpinBox_connect, "QDoubleSpinBox_connect")
	symaddr(&addr_QDoubleSpinBox_connect2, "QDoubleSpinBox_connect2")
	symaddr(&addr_QDoubleSpinBox_sender, "QDoubleSpinBox_sender")
	symaddr(&addr_QDoubleSpinBox_senderSignalIndex, "QDoubleSpinBox_senderSignalIndex")
	symaddr(&addr_QDoubleSpinBox_receivers, "QDoubleSpinBox_receivers")
	symaddr(&addr_QDoubleSpinBox_isSignalConnected, "QDoubleSpinBox_isSignalConnected")
	reg(&pg_QDoubleSpinBox_devicePixelRatioFScale, "QDoubleSpinBox_devicePixelRatioFScale")
	reg(&pg_QDoubleSpinBox_encodeMetricF, "QDoubleSpinBox_encodeMetricF")
	reg(&pg_QDoubleSpinBox_getDecodedMetricF, "QDoubleSpinBox_getDecodedMetricF")
	symaddr(&addr_QDoubleSpinBox_tr2, "QDoubleSpinBox_tr2")
	symaddr(&addr_QDoubleSpinBox_tr3, "QDoubleSpinBox_tr3")
	symaddr(&addr_QDoubleSpinBox_createWindowContainer2, "QDoubleSpinBox_createWindowContainer2")
	symaddr(&addr_QDoubleSpinBox_createWindowContainer3, "QDoubleSpinBox_createWindowContainer3")
	symaddr(&addr_QDoubleSpinBox_updateMicroFocus2, "QDoubleSpinBox_updateMicroFocus2")
	symaddr(&addr_QDoubleSpinBox_create2, "QDoubleSpinBox_create2")
	symaddr(&addr_QDoubleSpinBox_create3, "QDoubleSpinBox_create3")
	symaddr(&addr_QDoubleSpinBox_create4, "QDoubleSpinBox_create4")
	symaddr(&addr_QDoubleSpinBox_destroy2, "QDoubleSpinBox_destroy2")
	symaddr(&addr_QDoubleSpinBox_destroy3, "QDoubleSpinBox_destroy3")
	symaddr(&addr_QDoubleSpinBox_connect3, "QDoubleSpinBox_connect3")
	symaddr(&addr_QDoubleSpinBox_connect4, "QDoubleSpinBox_connect4")
	reg(&pg_QDoubleSpinBox_valueFromText_default, "QDoubleSpinBox_valueFromText_default")
	reg(&pg_QDoubleSpinBox_textFromValue_default, "QDoubleSpinBox_textFromValue_default")
	symaddr(&addr_QDoubleSpinBox_metaObject_default, "QDoubleSpinBox_metaObject_default")
	symaddr(&addr_QDoubleSpinBox_validate_default, "QDoubleSpinBox_validate_default")
	symaddr(&addr_QDoubleSpinBox_fixup_default, "QDoubleSpinBox_fixup_default")
	symaddr(&addr_QDoubleSpinBox_sizeHint_default, "QDoubleSpinBox_sizeHint_default")
	symaddr(&addr_QDoubleSpinBox_minimumSizeHint_default, "QDoubleSpinBox_minimumSizeHint_default")
	symaddr(&addr_QDoubleSpinBox_event_default, "QDoubleSpinBox_event_default")
	symaddr(&addr_QDoubleSpinBox_inputMethodQuery_default, "QDoubleSpinBox_inputMethodQuery_default")
	symaddr(&addr_QDoubleSpinBox_stepBy_default, "QDoubleSpinBox_stepBy_default")
	symaddr(&addr_QDoubleSpinBox_clear_default, "QDoubleSpinBox_clear_default")
	symaddr(&addr_QDoubleSpinBox_resizeEvent_default, "QDoubleSpinBox_resizeEvent_default")
	symaddr(&addr_QDoubleSpinBox_keyPressEvent_default, "QDoubleSpinBox_keyPressEvent_default")
	symaddr(&addr_QDoubleSpinBox_keyReleaseEvent_default, "QDoubleSpinBox_keyReleaseEvent_default")
	symaddr(&addr_QDoubleSpinBox_wheelEvent_default, "QDoubleSpinBox_wheelEvent_default")
	symaddr(&addr_QDoubleSpinBox_focusInEvent_default, "QDoubleSpinBox_focusInEvent_default")
	symaddr(&addr_QDoubleSpinBox_focusOutEvent_default, "QDoubleSpinBox_focusOutEvent_default")
	symaddr(&addr_QDoubleSpinBox_contextMenuEvent_default, "QDoubleSpinBox_contextMenuEvent_default")
	symaddr(&addr_QDoubleSpinBox_changeEvent_default, "QDoubleSpinBox_changeEvent_default")
	symaddr(&addr_QDoubleSpinBox_closeEvent_default, "QDoubleSpinBox_closeEvent_default")
	symaddr(&addr_QDoubleSpinBox_hideEvent_default, "QDoubleSpinBox_hideEvent_default")
	symaddr(&addr_QDoubleSpinBox_mousePressEvent_default, "QDoubleSpinBox_mousePressEvent_default")
	symaddr(&addr_QDoubleSpinBox_mouseReleaseEvent_default, "QDoubleSpinBox_mouseReleaseEvent_default")
	symaddr(&addr_QDoubleSpinBox_mouseMoveEvent_default, "QDoubleSpinBox_mouseMoveEvent_default")
	symaddr(&addr_QDoubleSpinBox_timerEvent_default, "QDoubleSpinBox_timerEvent_default")
	symaddr(&addr_QDoubleSpinBox_paintEvent_default, "QDoubleSpinBox_paintEvent_default")
	symaddr(&addr_QDoubleSpinBox_showEvent_default, "QDoubleSpinBox_showEvent_default")
	symaddr(&addr_QDoubleSpinBox_initStyleOption_default, "QDoubleSpinBox_initStyleOption_default")
	symaddr(&addr_QDoubleSpinBox_stepEnabled_default, "QDoubleSpinBox_stepEnabled_default")
	symaddr(&addr_QDoubleSpinBox_devType_default, "QDoubleSpinBox_devType_default")
	symaddr(&addr_QDoubleSpinBox_setVisible_default, "QDoubleSpinBox_setVisible_default")
	symaddr(&addr_QDoubleSpinBox_heightForWidth_default, "QDoubleSpinBox_heightForWidth_default")
	symaddr(&addr_QDoubleSpinBox_hasHeightForWidth_default, "QDoubleSpinBox_hasHeightForWidth_default")
	symaddr(&addr_QDoubleSpinBox_paintEngine_default, "QDoubleSpinBox_paintEngine_default")
	symaddr(&addr_QDoubleSpinBox_mouseDoubleClickEvent_default, "QDoubleSpinBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QDoubleSpinBox_enterEvent_default, "QDoubleSpinBox_enterEvent_default")
	symaddr(&addr_QDoubleSpinBox_leaveEvent_default, "QDoubleSpinBox_leaveEvent_default")
	symaddr(&addr_QDoubleSpinBox_moveEvent_default, "QDoubleSpinBox_moveEvent_default")
	symaddr(&addr_QDoubleSpinBox_tabletEvent_default, "QDoubleSpinBox_tabletEvent_default")
	symaddr(&addr_QDoubleSpinBox_actionEvent_default, "QDoubleSpinBox_actionEvent_default")
	symaddr(&addr_QDoubleSpinBox_dragEnterEvent_default, "QDoubleSpinBox_dragEnterEvent_default")
	symaddr(&addr_QDoubleSpinBox_dragMoveEvent_default, "QDoubleSpinBox_dragMoveEvent_default")
	symaddr(&addr_QDoubleSpinBox_dragLeaveEvent_default, "QDoubleSpinBox_dragLeaveEvent_default")
	symaddr(&addr_QDoubleSpinBox_dropEvent_default, "QDoubleSpinBox_dropEvent_default")
	symaddr(&addr_QDoubleSpinBox_metric_default, "QDoubleSpinBox_metric_default")
	symaddr(&addr_QDoubleSpinBox_initPainter_default, "QDoubleSpinBox_initPainter_default")
	symaddr(&addr_QDoubleSpinBox_redirected_default, "QDoubleSpinBox_redirected_default")
	symaddr(&addr_QDoubleSpinBox_sharedPainter_default, "QDoubleSpinBox_sharedPainter_default")
	symaddr(&addr_QDoubleSpinBox_inputMethodEvent_default, "QDoubleSpinBox_inputMethodEvent_default")
	symaddr(&addr_QDoubleSpinBox_focusNextPrevChild_default, "QDoubleSpinBox_focusNextPrevChild_default")
	symaddr(&addr_QDoubleSpinBox_eventFilter_default, "QDoubleSpinBox_eventFilter_default")
	symaddr(&addr_QDoubleSpinBox_childEvent_default, "QDoubleSpinBox_childEvent_default")
	symaddr(&addr_QDoubleSpinBox_customEvent_default, "QDoubleSpinBox_customEvent_default")
	symaddr(&addr_QDoubleSpinBox_connectNotify_default, "QDoubleSpinBox_connectNotify_default")
	symaddr(&addr_QDoubleSpinBox_disconnectNotify_default, "QDoubleSpinBox_disconnectNotify_default")
	reg(&pg_QDoubleSpinBox_connect_valueChanged, "QDoubleSpinBox_connect_valueChanged")
	reg(&pg_QDoubleSpinBox_connect_textChanged, "QDoubleSpinBox_connect_textChanged")
	reg(&pg_QDoubleSpinBox_connect_editingFinished, "QDoubleSpinBox_connect_editingFinished")
	reg(&pg_QDoubleSpinBox_connect_returnPressed, "QDoubleSpinBox_connect_returnPressed")
	reg(&pg_QDoubleSpinBox_connect_windowTitleChanged, "QDoubleSpinBox_connect_windowTitleChanged")
	reg(&pg_QDoubleSpinBox_connect_windowIconChanged, "QDoubleSpinBox_connect_windowIconChanged")
	reg(&pg_QDoubleSpinBox_connect_windowIconTextChanged, "QDoubleSpinBox_connect_windowIconTextChanged")
	reg(&pg_QDoubleSpinBox_connect_customContextMenuRequested, "QDoubleSpinBox_connect_customContextMenuRequested")
	reg(&pg_QDoubleSpinBox_connect_destroyed, "QDoubleSpinBox_connect_destroyed")
	reg(&pg_QDoubleSpinBox_connect_destroyed2, "QDoubleSpinBox_connect_destroyed2")
	reg(&pg_QDoubleSpinBox_connect_objectNameChanged, "QDoubleSpinBox_connect_objectNameChanged")
	reg(&pg_QDoubleSpinBox__disconnect, "QDoubleSpinBox__disconnect")
	reg(&pg_QDoubleSpinBox__set_override, "QDoubleSpinBox__set_override")
	reg(&pg_QDoubleSpinBox_delete, "QDoubleSpinBox_delete")
}

type QDoubleSpinBox struct {
	h uintptr
	*QAbstractSpinBox
}

func bindQDoubleSpinBox(h uintptr) *QDoubleSpinBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QDoubleSpinBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDoubleSpinBox{h: h, QAbstractSpinBox: &QAbstractSpinBox{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQDoubleSpinBox(h uintptr) *QDoubleSpinBox { return bindQDoubleSpinBox(h) }

func (this *QDoubleSpinBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDoubleSpinBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDoubleSpinBox(h unsafe.Pointer) *QDoubleSpinBox { return bindQDoubleSpinBox(uintptr(h)) }
func UnsafeNewQDoubleSpinBoxOwned(h unsafe.Pointer) *QDoubleSpinBox { return newownedQDoubleSpinBox(uintptr(h)) }

func (this *QDoubleSpinBox) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSpinBox != nil {
		this.QAbstractSpinBox.UnsafeRelease()
	}
}

func NewQDoubleSpinBox(parent *QWidget) *QDoubleSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDoubleSpinBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleSpinBox_virtIdx, QDoubleSpinBox_enable)
	return __o
}

func NewQDoubleSpinBox2() *QDoubleSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_new2, uintptr(__goid))
	__o := newownedQDoubleSpinBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleSpinBox_virtIdx, QDoubleSpinBox_enable)
	return __o
}

func QDoubleSpinBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) Value() float64 {
	__r := pg_QDoubleSpinBox_value(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) Prefix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_prefix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) SetPrefix(prefix string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QDoubleSpinBox_setPrefix, this.h, prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func (this *QDoubleSpinBox) Suffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_suffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) SetSuffix(suffix string) {
	var suffix_p uintptr
	if len(suffix) > 0 {
		suffix_p = uintptr(unsafe.Pointer(unsafe.StringData(suffix)))
	}
	purego.SyscallN(addr_QDoubleSpinBox_setSuffix, this.h, suffix_p, uintptr(len(suffix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(suffix)
}

func (this *QDoubleSpinBox) CleanText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_cleanText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) SingleStep() float64 {
	__r := pg_QDoubleSpinBox_singleStep(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	pg_QDoubleSpinBox_setSingleStep(this.h, val)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Minimum() float64 {
	__r := pg_QDoubleSpinBox_minimum(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) SetMinimum(min float64) {
	pg_QDoubleSpinBox_setMinimum(this.h, min)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Maximum() float64 {
	__r := pg_QDoubleSpinBox_maximum(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) SetMaximum(max float64) {
	pg_QDoubleSpinBox_setMaximum(this.h, max)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	pg_QDoubleSpinBox_setRange(this.h, min, max)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) StepType() QAbstractSpinBox__StepType {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_stepType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepType(__r)
}

func (this *QDoubleSpinBox) SetStepType(stepType QAbstractSpinBox__StepType) {
	purego.SyscallN(addr_QDoubleSpinBox_setStepType, this.h, uintptr(stepType))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Decimals() int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_decimals, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleSpinBox) SetDecimals(prec int) {
	purego.SyscallN(addr_QDoubleSpinBox_setDecimals, this.h, uintptr(prec))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) ValueFromText(text string) float64 {
	__r := pg_QDoubleSpinBox_valueFromText(this.h, text, int64(len(text)))
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) TextFromValue(val float64) string {
	var __n int64
	__r := pg_QDoubleSpinBox_textFromValue(this.h, val, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) SetValue(val float64) {
	pg_QDoubleSpinBox_setValue(this.h, val)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) ValueChanged(arg1 float64) {
	pg_QDoubleSpinBox_valueChanged(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) TextChanged(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QDoubleSpinBox_textChanged, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDoubleSpinBox) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QDoubleSpinBox) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QDoubleSpinBox_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func QDoubleSpinBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDoubleSpinBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDoubleSpinBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_mouseGrabber)
	return bindQWidget(__r)
}

func QDoubleSpinBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QDoubleSpinBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDoubleSpinBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDoubleSpinBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QDoubleSpinBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Create() {
	purego.SyscallN(addr_QDoubleSpinBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Destroy() {
	purego.SyscallN(addr_QDoubleSpinBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDoubleSpinBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDoubleSpinBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDoubleSpinBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDoubleSpinBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDoubleSpinBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleSpinBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDoubleSpinBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDoubleSpinBox_DevicePixelRatioFScale() float64 {
	__r := pg_QDoubleSpinBox_devicePixelRatioFScale()
	return __r
}

func QDoubleSpinBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDoubleSpinBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDoubleSpinBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDoubleSpinBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDoubleSpinBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDoubleSpinBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDoubleSpinBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDoubleSpinBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDoubleSpinBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDoubleSpinBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDoubleSpinBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDoubleSpinBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDoubleSpinBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDoubleSpinBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDoubleSpinBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDoubleSpinBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDoubleSpinBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDoubleSpinBox) ValueFromTextDefault(text string) float64 {
	__r := pg_QDoubleSpinBox_valueFromText_default(this.h, text, int64(len(text)))
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleSpinBox) TextFromValueDefault(val float64) string {
	var __n int64
	__r := pg_QDoubleSpinBox_textFromValue_default(this.h, val, &__n)
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDoubleSpinBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDoubleSpinBox) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QDoubleSpinBox) FixupDefault(str *QString) {
	purego.SyscallN(addr_QDoubleSpinBox_fixup_default, this.h, uintptr(str.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
}

func (this *QDoubleSpinBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDoubleSpinBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDoubleSpinBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDoubleSpinBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDoubleSpinBox) StepByDefault(steps int) {
	purego.SyscallN(addr_QDoubleSpinBox_stepBy_default, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) ClearDefault() {
	purego.SyscallN(addr_QDoubleSpinBox_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) InitStyleOptionDefault(option *QStyleOptionSpinBox) {
	purego.SyscallN(addr_QDoubleSpinBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QDoubleSpinBox) StepEnabledDefault() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_stepEnabled_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QDoubleSpinBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleSpinBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDoubleSpinBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDoubleSpinBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleSpinBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDoubleSpinBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDoubleSpinBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleSpinBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDoubleSpinBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDoubleSpinBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDoubleSpinBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDoubleSpinBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDoubleSpinBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDoubleSpinBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDoubleSpinBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleSpinBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDoubleSpinBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDoubleSpinBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleSpinBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDoubleSpinBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDoubleSpinBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDoubleSpinBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDoubleSpinBox) OnValueChanged(fn func(float64)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(math.Float64frombits(uint64(arg1))) })
	pg_QDoubleSpinBox_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QDoubleSpinBox_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDoubleSpinBox_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDoubleSpinBox_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDoubleSpinBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDoubleSpinBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDoubleSpinBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDoubleSpinBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDoubleSpinBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDoubleSpinBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDoubleSpinBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleSpinBox) Disconnect(handle int64) {
	pg_QDoubleSpinBox__disconnect(handle)
	unregisterCallback(handle)
}

func QDoubleSpinBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6784179835071232555:
		return 1
	case 5050482105304000557:
		return 2
	case -9083693956578760701:
		return 3
	case -4855206404746966715:
		return 4
	case 729690261457684101:
		return 5
	case 7049264599837532491:
		return 6
	case 1237236456312824991:
		return 7
	case 7345853097340709638:
		return 8
	case 7085718643589954634:
		return 9
	case -778567940572639070:
		return 10
	case -4302917285655014661:
		return 11
	case 1023791977679793185:
		return 12
	case 772598475950050299:
		return 13
	case 1919988340934511316:
		return 14
	case 975043767033552552:
		return 15
	case -385596010576048993:
		return 16
	case -6698043222227006015:
		return 17
	case 4959376063826165943:
		return 18
	case -1989937050019464783:
		return 19
	case -199718707876719129:
		return 20
	case -332090898078543319:
		return 21
	case -4415920049297795213:
		return 22
	case -4255494004099806581:
		return 23
	case -5409927581812374498:
		return 24
	case 5290442163167458465:
		return 25
	case -6825369174751916:
		return 26
	case -488693769771308531:
		return 27
	case 5257289574516412244:
		return 28
	case -8659679893868524044:
		return 29
	case 2110182719375245575:
		return 30
	case 4311534060495737339:
		return 31
	case -8266762397995423915:
		return 32
	case 6398007404756346381:
		return 33
	case -4029073905697023887:
		return 34
	case 6316270346668237895:
		return 35
	case -551070009623711640:
		return 36
	case -9161949150489525396:
		return 37
	case -212068866860701141:
		return 38
	case 1697444140088555581:
		return 39
	case 206122715919917589:
		return 40
	case -1781545137352745978:
		return 41
	case -3139904158373725006:
		return 42
	case 9107058682754591716:
		return 43
	case -2039166586246443706:
		return 44
	case 427629649267311777:
		return 45
	case 6477500296980050210:
		return 46
	case 4127240633310655048:
		return 47
	case 5792771328741774991:
		return 48
	case 5438819661942843222:
		return 49
	case 5898328517505784191:
		return 50
	case 5217623514630858297:
		return 51
	case -1185315367526857833:
		return 52
	case 3769868229123868474:
		return 53
	case 7379029982170503008:
		return 54
	case -568369772982215962:
		return 55
	}
	return -1
}

func QDoubleSpinBox_enable(h uintptr, idx int32) { pg_QDoubleSpinBox__set_override(h, idx) }

func (this *QDoubleSpinBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDoubleSpinBox) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QDoubleSpinBox) OnValueFromText(fn func(string) float64) {
	registerVirt(this.UnsafeGoID(), 5050482105304000557, func(text uintptr, text_len int64) uintptr { return uintptr(math.Float64bits(fn(pgStrNoFree(text, text_len)))) }) // slot "valueFromText"
}

func (this *QDoubleSpinBox) OnTextFromValue(fn func(float64) string) {
	registerVirt(this.UnsafeGoID(), -9083693956578760701, func(val uintptr) (uintptr, int64) {
		__s := fn(math.Float64frombits(uint64(val)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textFromValue"
}

func (this *QDoubleSpinBox) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(str uintptr) { fn(UnsafeNewQString(unsafe.Pointer(str))) }) // slot "fixup"
}

func (this *QDoubleSpinBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDoubleSpinBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDoubleSpinBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDoubleSpinBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDoubleSpinBox) OnStepBy(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 7085718643589954634, func(steps int32) { fn(int(steps)) }) // slot "stepBy"
}

func (this *QDoubleSpinBox) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QDoubleSpinBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QDoubleSpinBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QDoubleSpinBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDoubleSpinBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDoubleSpinBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDoubleSpinBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDoubleSpinBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDoubleSpinBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QDoubleSpinBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDoubleSpinBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDoubleSpinBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDoubleSpinBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDoubleSpinBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDoubleSpinBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDoubleSpinBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDoubleSpinBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDoubleSpinBox) OnInitStyleOption(fn func(*QStyleOptionSpinBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSpinBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QDoubleSpinBox) OnStepEnabled(fn func() QAbstractSpinBox__StepEnabledFlag) {
	registerVirt(this.UnsafeGoID(), 5257289574516412244, func() int32 { return int32(fn()) }) // slot "stepEnabled"
}

func (this *QDoubleSpinBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDoubleSpinBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDoubleSpinBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDoubleSpinBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDoubleSpinBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDoubleSpinBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDoubleSpinBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDoubleSpinBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDoubleSpinBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDoubleSpinBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDoubleSpinBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDoubleSpinBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDoubleSpinBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDoubleSpinBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDoubleSpinBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDoubleSpinBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDoubleSpinBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDoubleSpinBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDoubleSpinBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDoubleSpinBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDoubleSpinBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDoubleSpinBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDoubleSpinBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDoubleSpinBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDoubleSpinBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDoubleSpinBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDoubleSpinBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDoubleSpinBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDoubleSpinBox_delete(this.h)
	}
	this.UnsafeRelease()
}
