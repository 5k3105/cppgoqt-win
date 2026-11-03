package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSpinBox_chain uintptr
	addr_QSpinBox_new uintptr
	addr_QSpinBox_new2 uintptr
	addr_QSpinBox_tr uintptr
	addr_QSpinBox_value uintptr
	addr_QSpinBox_prefix uintptr
	addr_QSpinBox_setPrefix uintptr
	addr_QSpinBox_suffix uintptr
	addr_QSpinBox_setSuffix uintptr
	addr_QSpinBox_cleanText uintptr
	addr_QSpinBox_singleStep uintptr
	addr_QSpinBox_setSingleStep uintptr
	addr_QSpinBox_minimum uintptr
	addr_QSpinBox_setMinimum uintptr
	addr_QSpinBox_maximum uintptr
	addr_QSpinBox_setMaximum uintptr
	addr_QSpinBox_setRange uintptr
	addr_QSpinBox_stepType uintptr
	addr_QSpinBox_setStepType uintptr
	addr_QSpinBox_displayIntegerBase uintptr
	addr_QSpinBox_setDisplayIntegerBase uintptr
	addr_QSpinBox_setValue uintptr
	addr_QSpinBox_valueChanged uintptr
	addr_QSpinBox_textChanged uintptr
	addr_QSpinBox_lineEdit uintptr
	addr_QSpinBox_setLineEdit uintptr
	addr_QSpinBox_setTabOrder uintptr
	addr_QSpinBox_mouseGrabber uintptr
	addr_QSpinBox_keyboardGrabber uintptr
	addr_QSpinBox_find uintptr
	addr_QSpinBox_createWindowContainer uintptr
	addr_QSpinBox_updateMicroFocus uintptr
	addr_QSpinBox_create uintptr
	addr_QSpinBox_destroy uintptr
	addr_QSpinBox_focusNextChild uintptr
	addr_QSpinBox_focusPreviousChild uintptr
	addr_QSpinBox_connect uintptr
	addr_QSpinBox_connect2 uintptr
	addr_QSpinBox_sender uintptr
	addr_QSpinBox_senderSignalIndex uintptr
	addr_QSpinBox_receivers uintptr
	addr_QSpinBox_isSignalConnected uintptr
	pg_QSpinBox_devicePixelRatioFScale func() float64
	pg_QSpinBox_encodeMetricF func(int32, float64) int32
	pg_QSpinBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QSpinBox_tr2 uintptr
	addr_QSpinBox_tr3 uintptr
	addr_QSpinBox_createWindowContainer2 uintptr
	addr_QSpinBox_createWindowContainer3 uintptr
	addr_QSpinBox_updateMicroFocus2 uintptr
	addr_QSpinBox_create2 uintptr
	addr_QSpinBox_create3 uintptr
	addr_QSpinBox_create4 uintptr
	addr_QSpinBox_destroy2 uintptr
	addr_QSpinBox_destroy3 uintptr
	addr_QSpinBox_connect3 uintptr
	addr_QSpinBox_connect4 uintptr
	addr_QSpinBox_metaObject_default uintptr
	addr_QSpinBox_event_default uintptr
	addr_QSpinBox_validate_default uintptr
	addr_QSpinBox_valueFromText_default uintptr
	addr_QSpinBox_textFromValue_default uintptr
	addr_QSpinBox_fixup_default uintptr
	addr_QSpinBox_sizeHint_default uintptr
	addr_QSpinBox_minimumSizeHint_default uintptr
	addr_QSpinBox_inputMethodQuery_default uintptr
	addr_QSpinBox_stepBy_default uintptr
	addr_QSpinBox_clear_default uintptr
	addr_QSpinBox_resizeEvent_default uintptr
	addr_QSpinBox_keyPressEvent_default uintptr
	addr_QSpinBox_keyReleaseEvent_default uintptr
	addr_QSpinBox_wheelEvent_default uintptr
	addr_QSpinBox_focusInEvent_default uintptr
	addr_QSpinBox_focusOutEvent_default uintptr
	addr_QSpinBox_contextMenuEvent_default uintptr
	addr_QSpinBox_changeEvent_default uintptr
	addr_QSpinBox_closeEvent_default uintptr
	addr_QSpinBox_hideEvent_default uintptr
	addr_QSpinBox_mousePressEvent_default uintptr
	addr_QSpinBox_mouseReleaseEvent_default uintptr
	addr_QSpinBox_mouseMoveEvent_default uintptr
	addr_QSpinBox_timerEvent_default uintptr
	addr_QSpinBox_paintEvent_default uintptr
	addr_QSpinBox_showEvent_default uintptr
	addr_QSpinBox_initStyleOption_default uintptr
	addr_QSpinBox_stepEnabled_default uintptr
	addr_QSpinBox_devType_default uintptr
	addr_QSpinBox_setVisible_default uintptr
	addr_QSpinBox_heightForWidth_default uintptr
	addr_QSpinBox_hasHeightForWidth_default uintptr
	addr_QSpinBox_paintEngine_default uintptr
	addr_QSpinBox_mouseDoubleClickEvent_default uintptr
	addr_QSpinBox_enterEvent_default uintptr
	addr_QSpinBox_leaveEvent_default uintptr
	addr_QSpinBox_moveEvent_default uintptr
	addr_QSpinBox_tabletEvent_default uintptr
	addr_QSpinBox_actionEvent_default uintptr
	addr_QSpinBox_dragEnterEvent_default uintptr
	addr_QSpinBox_dragMoveEvent_default uintptr
	addr_QSpinBox_dragLeaveEvent_default uintptr
	addr_QSpinBox_dropEvent_default uintptr
	addr_QSpinBox_metric_default uintptr
	addr_QSpinBox_initPainter_default uintptr
	addr_QSpinBox_redirected_default uintptr
	addr_QSpinBox_sharedPainter_default uintptr
	addr_QSpinBox_inputMethodEvent_default uintptr
	addr_QSpinBox_focusNextPrevChild_default uintptr
	addr_QSpinBox_eventFilter_default uintptr
	addr_QSpinBox_childEvent_default uintptr
	addr_QSpinBox_customEvent_default uintptr
	addr_QSpinBox_connectNotify_default uintptr
	addr_QSpinBox_disconnectNotify_default uintptr
	pg_QSpinBox_connect_valueChanged func(uintptr, int64)
	pg_QSpinBox_connect_textChanged func(uintptr, int64)
	pg_QSpinBox_connect_editingFinished func(uintptr, int64)
	pg_QSpinBox_connect_returnPressed func(uintptr, int64)
	pg_QSpinBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QSpinBox_connect_windowIconChanged func(uintptr, int64)
	pg_QSpinBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSpinBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSpinBox_connect_destroyed func(uintptr, int64)
	pg_QSpinBox_connect_destroyed2 func(uintptr, int64)
	pg_QSpinBox_connect_objectNameChanged func(uintptr, int64)
	pg_QSpinBox__disconnect func(int64)
	pg_QSpinBox__set_override func(uintptr, int32)
	pg_QSpinBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSpinBox_chain, "QSpinBox_chain")
	symaddr(&addr_QSpinBox_new, "QSpinBox_new")
	symaddr(&addr_QSpinBox_new2, "QSpinBox_new2")
	symaddr(&addr_QSpinBox_tr, "QSpinBox_tr")
	symaddr(&addr_QSpinBox_value, "QSpinBox_value")
	symaddr(&addr_QSpinBox_prefix, "QSpinBox_prefix")
	symaddr(&addr_QSpinBox_setPrefix, "QSpinBox_setPrefix")
	symaddr(&addr_QSpinBox_suffix, "QSpinBox_suffix")
	symaddr(&addr_QSpinBox_setSuffix, "QSpinBox_setSuffix")
	symaddr(&addr_QSpinBox_cleanText, "QSpinBox_cleanText")
	symaddr(&addr_QSpinBox_singleStep, "QSpinBox_singleStep")
	symaddr(&addr_QSpinBox_setSingleStep, "QSpinBox_setSingleStep")
	symaddr(&addr_QSpinBox_minimum, "QSpinBox_minimum")
	symaddr(&addr_QSpinBox_setMinimum, "QSpinBox_setMinimum")
	symaddr(&addr_QSpinBox_maximum, "QSpinBox_maximum")
	symaddr(&addr_QSpinBox_setMaximum, "QSpinBox_setMaximum")
	symaddr(&addr_QSpinBox_setRange, "QSpinBox_setRange")
	symaddr(&addr_QSpinBox_stepType, "QSpinBox_stepType")
	symaddr(&addr_QSpinBox_setStepType, "QSpinBox_setStepType")
	symaddr(&addr_QSpinBox_displayIntegerBase, "QSpinBox_displayIntegerBase")
	symaddr(&addr_QSpinBox_setDisplayIntegerBase, "QSpinBox_setDisplayIntegerBase")
	symaddr(&addr_QSpinBox_setValue, "QSpinBox_setValue")
	symaddr(&addr_QSpinBox_valueChanged, "QSpinBox_valueChanged")
	symaddr(&addr_QSpinBox_textChanged, "QSpinBox_textChanged")
	symaddr(&addr_QSpinBox_lineEdit, "QSpinBox_lineEdit")
	symaddr(&addr_QSpinBox_setLineEdit, "QSpinBox_setLineEdit")
	symaddr(&addr_QSpinBox_setTabOrder, "QSpinBox_setTabOrder")
	symaddr(&addr_QSpinBox_mouseGrabber, "QSpinBox_mouseGrabber")
	symaddr(&addr_QSpinBox_keyboardGrabber, "QSpinBox_keyboardGrabber")
	symaddr(&addr_QSpinBox_find, "QSpinBox_find")
	symaddr(&addr_QSpinBox_createWindowContainer, "QSpinBox_createWindowContainer")
	symaddr(&addr_QSpinBox_updateMicroFocus, "QSpinBox_updateMicroFocus")
	symaddr(&addr_QSpinBox_create, "QSpinBox_create")
	symaddr(&addr_QSpinBox_destroy, "QSpinBox_destroy")
	symaddr(&addr_QSpinBox_focusNextChild, "QSpinBox_focusNextChild")
	symaddr(&addr_QSpinBox_focusPreviousChild, "QSpinBox_focusPreviousChild")
	symaddr(&addr_QSpinBox_connect, "QSpinBox_connect")
	symaddr(&addr_QSpinBox_connect2, "QSpinBox_connect2")
	symaddr(&addr_QSpinBox_sender, "QSpinBox_sender")
	symaddr(&addr_QSpinBox_senderSignalIndex, "QSpinBox_senderSignalIndex")
	symaddr(&addr_QSpinBox_receivers, "QSpinBox_receivers")
	symaddr(&addr_QSpinBox_isSignalConnected, "QSpinBox_isSignalConnected")
	reg(&pg_QSpinBox_devicePixelRatioFScale, "QSpinBox_devicePixelRatioFScale")
	reg(&pg_QSpinBox_encodeMetricF, "QSpinBox_encodeMetricF")
	reg(&pg_QSpinBox_getDecodedMetricF, "QSpinBox_getDecodedMetricF")
	symaddr(&addr_QSpinBox_tr2, "QSpinBox_tr2")
	symaddr(&addr_QSpinBox_tr3, "QSpinBox_tr3")
	symaddr(&addr_QSpinBox_createWindowContainer2, "QSpinBox_createWindowContainer2")
	symaddr(&addr_QSpinBox_createWindowContainer3, "QSpinBox_createWindowContainer3")
	symaddr(&addr_QSpinBox_updateMicroFocus2, "QSpinBox_updateMicroFocus2")
	symaddr(&addr_QSpinBox_create2, "QSpinBox_create2")
	symaddr(&addr_QSpinBox_create3, "QSpinBox_create3")
	symaddr(&addr_QSpinBox_create4, "QSpinBox_create4")
	symaddr(&addr_QSpinBox_destroy2, "QSpinBox_destroy2")
	symaddr(&addr_QSpinBox_destroy3, "QSpinBox_destroy3")
	symaddr(&addr_QSpinBox_connect3, "QSpinBox_connect3")
	symaddr(&addr_QSpinBox_connect4, "QSpinBox_connect4")
	symaddr(&addr_QSpinBox_metaObject_default, "QSpinBox_metaObject_default")
	symaddr(&addr_QSpinBox_event_default, "QSpinBox_event_default")
	symaddr(&addr_QSpinBox_validate_default, "QSpinBox_validate_default")
	symaddr(&addr_QSpinBox_valueFromText_default, "QSpinBox_valueFromText_default")
	symaddr(&addr_QSpinBox_textFromValue_default, "QSpinBox_textFromValue_default")
	symaddr(&addr_QSpinBox_fixup_default, "QSpinBox_fixup_default")
	symaddr(&addr_QSpinBox_sizeHint_default, "QSpinBox_sizeHint_default")
	symaddr(&addr_QSpinBox_minimumSizeHint_default, "QSpinBox_minimumSizeHint_default")
	symaddr(&addr_QSpinBox_inputMethodQuery_default, "QSpinBox_inputMethodQuery_default")
	symaddr(&addr_QSpinBox_stepBy_default, "QSpinBox_stepBy_default")
	symaddr(&addr_QSpinBox_clear_default, "QSpinBox_clear_default")
	symaddr(&addr_QSpinBox_resizeEvent_default, "QSpinBox_resizeEvent_default")
	symaddr(&addr_QSpinBox_keyPressEvent_default, "QSpinBox_keyPressEvent_default")
	symaddr(&addr_QSpinBox_keyReleaseEvent_default, "QSpinBox_keyReleaseEvent_default")
	symaddr(&addr_QSpinBox_wheelEvent_default, "QSpinBox_wheelEvent_default")
	symaddr(&addr_QSpinBox_focusInEvent_default, "QSpinBox_focusInEvent_default")
	symaddr(&addr_QSpinBox_focusOutEvent_default, "QSpinBox_focusOutEvent_default")
	symaddr(&addr_QSpinBox_contextMenuEvent_default, "QSpinBox_contextMenuEvent_default")
	symaddr(&addr_QSpinBox_changeEvent_default, "QSpinBox_changeEvent_default")
	symaddr(&addr_QSpinBox_closeEvent_default, "QSpinBox_closeEvent_default")
	symaddr(&addr_QSpinBox_hideEvent_default, "QSpinBox_hideEvent_default")
	symaddr(&addr_QSpinBox_mousePressEvent_default, "QSpinBox_mousePressEvent_default")
	symaddr(&addr_QSpinBox_mouseReleaseEvent_default, "QSpinBox_mouseReleaseEvent_default")
	symaddr(&addr_QSpinBox_mouseMoveEvent_default, "QSpinBox_mouseMoveEvent_default")
	symaddr(&addr_QSpinBox_timerEvent_default, "QSpinBox_timerEvent_default")
	symaddr(&addr_QSpinBox_paintEvent_default, "QSpinBox_paintEvent_default")
	symaddr(&addr_QSpinBox_showEvent_default, "QSpinBox_showEvent_default")
	symaddr(&addr_QSpinBox_initStyleOption_default, "QSpinBox_initStyleOption_default")
	symaddr(&addr_QSpinBox_stepEnabled_default, "QSpinBox_stepEnabled_default")
	symaddr(&addr_QSpinBox_devType_default, "QSpinBox_devType_default")
	symaddr(&addr_QSpinBox_setVisible_default, "QSpinBox_setVisible_default")
	symaddr(&addr_QSpinBox_heightForWidth_default, "QSpinBox_heightForWidth_default")
	symaddr(&addr_QSpinBox_hasHeightForWidth_default, "QSpinBox_hasHeightForWidth_default")
	symaddr(&addr_QSpinBox_paintEngine_default, "QSpinBox_paintEngine_default")
	symaddr(&addr_QSpinBox_mouseDoubleClickEvent_default, "QSpinBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QSpinBox_enterEvent_default, "QSpinBox_enterEvent_default")
	symaddr(&addr_QSpinBox_leaveEvent_default, "QSpinBox_leaveEvent_default")
	symaddr(&addr_QSpinBox_moveEvent_default, "QSpinBox_moveEvent_default")
	symaddr(&addr_QSpinBox_tabletEvent_default, "QSpinBox_tabletEvent_default")
	symaddr(&addr_QSpinBox_actionEvent_default, "QSpinBox_actionEvent_default")
	symaddr(&addr_QSpinBox_dragEnterEvent_default, "QSpinBox_dragEnterEvent_default")
	symaddr(&addr_QSpinBox_dragMoveEvent_default, "QSpinBox_dragMoveEvent_default")
	symaddr(&addr_QSpinBox_dragLeaveEvent_default, "QSpinBox_dragLeaveEvent_default")
	symaddr(&addr_QSpinBox_dropEvent_default, "QSpinBox_dropEvent_default")
	symaddr(&addr_QSpinBox_metric_default, "QSpinBox_metric_default")
	symaddr(&addr_QSpinBox_initPainter_default, "QSpinBox_initPainter_default")
	symaddr(&addr_QSpinBox_redirected_default, "QSpinBox_redirected_default")
	symaddr(&addr_QSpinBox_sharedPainter_default, "QSpinBox_sharedPainter_default")
	symaddr(&addr_QSpinBox_inputMethodEvent_default, "QSpinBox_inputMethodEvent_default")
	symaddr(&addr_QSpinBox_focusNextPrevChild_default, "QSpinBox_focusNextPrevChild_default")
	symaddr(&addr_QSpinBox_eventFilter_default, "QSpinBox_eventFilter_default")
	symaddr(&addr_QSpinBox_childEvent_default, "QSpinBox_childEvent_default")
	symaddr(&addr_QSpinBox_customEvent_default, "QSpinBox_customEvent_default")
	symaddr(&addr_QSpinBox_connectNotify_default, "QSpinBox_connectNotify_default")
	symaddr(&addr_QSpinBox_disconnectNotify_default, "QSpinBox_disconnectNotify_default")
	reg(&pg_QSpinBox_connect_valueChanged, "QSpinBox_connect_valueChanged")
	reg(&pg_QSpinBox_connect_textChanged, "QSpinBox_connect_textChanged")
	reg(&pg_QSpinBox_connect_editingFinished, "QSpinBox_connect_editingFinished")
	reg(&pg_QSpinBox_connect_returnPressed, "QSpinBox_connect_returnPressed")
	reg(&pg_QSpinBox_connect_windowTitleChanged, "QSpinBox_connect_windowTitleChanged")
	reg(&pg_QSpinBox_connect_windowIconChanged, "QSpinBox_connect_windowIconChanged")
	reg(&pg_QSpinBox_connect_windowIconTextChanged, "QSpinBox_connect_windowIconTextChanged")
	reg(&pg_QSpinBox_connect_customContextMenuRequested, "QSpinBox_connect_customContextMenuRequested")
	reg(&pg_QSpinBox_connect_destroyed, "QSpinBox_connect_destroyed")
	reg(&pg_QSpinBox_connect_destroyed2, "QSpinBox_connect_destroyed2")
	reg(&pg_QSpinBox_connect_objectNameChanged, "QSpinBox_connect_objectNameChanged")
	reg(&pg_QSpinBox__disconnect, "QSpinBox__disconnect")
	reg(&pg_QSpinBox__set_override, "QSpinBox__set_override")
	reg(&pg_QSpinBox_delete, "QSpinBox_delete")
}

type QSpinBox struct {
	h uintptr
	*QAbstractSpinBox
}

func bindQSpinBox(h uintptr) *QSpinBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSpinBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSpinBox{h: h, QAbstractSpinBox: &QAbstractSpinBox{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSpinBox(h uintptr) *QSpinBox { return bindQSpinBox(h) }

func (this *QSpinBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSpinBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSpinBox(h unsafe.Pointer) *QSpinBox { return bindQSpinBox(uintptr(h)) }
func UnsafeNewQSpinBoxOwned(h unsafe.Pointer) *QSpinBox { return newownedQSpinBox(uintptr(h)) }

func (this *QSpinBox) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSpinBox != nil {
		this.QAbstractSpinBox.UnsafeRelease()
	}
}

func NewQSpinBox(parent *QWidget) *QSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSpinBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSpinBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSpinBox_virtIdx, QSpinBox_enable)
	return __o
}

func NewQSpinBox2() *QSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSpinBox_new2, uintptr(__goid))
	__o := newownedQSpinBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSpinBox_virtIdx, QSpinBox_enable)
	return __o
}

func QSpinBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSpinBox) Value() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) Prefix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_prefix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSpinBox) SetPrefix(prefix string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QSpinBox_setPrefix, this.h, prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func (this *QSpinBox) Suffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_suffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSpinBox) SetSuffix(suffix string) {
	var suffix_p uintptr
	if len(suffix) > 0 {
		suffix_p = uintptr(unsafe.Pointer(unsafe.StringData(suffix)))
	}
	purego.SyscallN(addr_QSpinBox_setSuffix, this.h, suffix_p, uintptr(len(suffix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(suffix)
}

func (this *QSpinBox) CleanText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_cleanText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSpinBox) SingleStep() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_singleStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) SetSingleStep(val int) {
	purego.SyscallN(addr_QSpinBox_setSingleStep, this.h, uintptr(val))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_minimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) SetMinimum(min int) {
	purego.SyscallN(addr_QSpinBox_setMinimum, this.h, uintptr(min))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_maximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) SetMaximum(max int) {
	purego.SyscallN(addr_QSpinBox_setMaximum, this.h, uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) SetRange(min int, max int) {
	purego.SyscallN(addr_QSpinBox_setRange, this.h, uintptr(min), uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) StepType() QAbstractSpinBox__StepType {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_stepType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepType(__r)
}

func (this *QSpinBox) SetStepType(stepType QAbstractSpinBox__StepType) {
	purego.SyscallN(addr_QSpinBox_setStepType, this.h, uintptr(stepType))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) DisplayIntegerBase() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_displayIntegerBase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	purego.SyscallN(addr_QSpinBox_setDisplayIntegerBase, this.h, uintptr(base))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) SetValue(val int) {
	purego.SyscallN(addr_QSpinBox_setValue, this.h, uintptr(val))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) ValueChanged(arg1 int) {
	purego.SyscallN(addr_QSpinBox_valueChanged, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) TextChanged(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QSpinBox_textChanged, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSpinBox) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QSpinBox) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QSpinBox_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func QSpinBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSpinBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSpinBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_mouseGrabber)
	return bindQWidget(__r)
}

func QSpinBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QSpinBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSpinBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QSpinBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QSpinBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Create() {
	purego.SyscallN(addr_QSpinBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Destroy() {
	purego.SyscallN(addr_QSpinBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QSpinBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSpinBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSpinBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSpinBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSpinBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSpinBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSpinBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSpinBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSpinBox_DevicePixelRatioFScale() float64 {
	__r := pg_QSpinBox_devicePixelRatioFScale()
	return __r
}

func QSpinBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSpinBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QSpinBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QSpinBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QSpinBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSpinBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSpinBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSpinBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QSpinBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QSpinBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QSpinBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QSpinBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QSpinBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QSpinBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QSpinBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QSpinBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSpinBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSpinBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSpinBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSpinBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSpinBox) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QSpinBox) ValueFromTextDefault(text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_valueFromText_default, this.h, text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QSpinBox) TextFromValueDefault(val int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSpinBox_textFromValue_default, this.h, uintptr(val), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSpinBox) FixupDefault(str *QString) {
	purego.SyscallN(addr_QSpinBox_fixup_default, this.h, uintptr(str.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
}

func (this *QSpinBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSpinBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSpinBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSpinBox) StepByDefault(steps int) {
	purego.SyscallN(addr_QSpinBox_stepBy_default, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) ClearDefault() {
	purego.SyscallN(addr_QSpinBox_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QSpinBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QSpinBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSpinBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSpinBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QSpinBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSpinBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSpinBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QSpinBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSpinBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QSpinBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QSpinBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSpinBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSpinBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSpinBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSpinBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QSpinBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QSpinBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) InitStyleOptionDefault(option *QStyleOptionSpinBox) {
	purego.SyscallN(addr_QSpinBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QSpinBox) StepEnabledDefault() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_stepEnabled_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QSpinBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QSpinBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QSpinBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSpinBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QSpinBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSpinBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QSpinBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSpinBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QSpinBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QSpinBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QSpinBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QSpinBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QSpinBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QSpinBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QSpinBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSpinBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QSpinBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QSpinBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QSpinBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QSpinBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QSpinBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QSpinBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSpinBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSpinBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSpinBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSpinBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSpinBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSpinBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSpinBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSpinBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSpinBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSpinBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSpinBox) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QSpinBox_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QSpinBox_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSpinBox_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSpinBox_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSpinBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSpinBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSpinBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSpinBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSpinBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSpinBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSpinBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSpinBox) Disconnect(handle int64) {
	pg_QSpinBox__disconnect(handle)
	unregisterCallback(handle)
}

func QSpinBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 6784179835071232555:
		return 2
	case 5050482105304000557:
		return 3
	case -9083693956578760701:
		return 4
	case -4855206404746966715:
		return 5
	case 729690261457684101:
		return 6
	case 7049264599837532491:
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

func QSpinBox_enable(h uintptr, idx int32) { pg_QSpinBox__set_override(h, idx) }

func (this *QSpinBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSpinBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSpinBox) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QSpinBox) OnValueFromText(fn func(string) int) {
	registerVirt(this.UnsafeGoID(), 5050482105304000557, func(text uintptr, text_len int64) int32 { return int32(fn(pgStrNoFree(text, text_len))) }) // slot "valueFromText"
}

func (this *QSpinBox) OnTextFromValue(fn func(int) string) {
	registerVirt(this.UnsafeGoID(), -9083693956578760701, func(val int32) (uintptr, int64) {
		__s := fn(int(val))
		return pgCString(__s), int64(len(__s))
	}) // slot "textFromValue"
}

func (this *QSpinBox) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(str uintptr) { fn(UnsafeNewQString(unsafe.Pointer(str))) }) // slot "fixup"
}

func (this *QSpinBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QSpinBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QSpinBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QSpinBox) OnStepBy(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 7085718643589954634, func(steps int32) { fn(int(steps)) }) // slot "stepBy"
}

func (this *QSpinBox) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QSpinBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QSpinBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QSpinBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QSpinBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QSpinBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QSpinBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QSpinBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QSpinBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QSpinBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QSpinBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QSpinBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QSpinBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QSpinBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QSpinBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSpinBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QSpinBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QSpinBox) OnInitStyleOption(fn func(*QStyleOptionSpinBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSpinBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QSpinBox) OnStepEnabled(fn func() QAbstractSpinBox__StepEnabledFlag) {
	registerVirt(this.UnsafeGoID(), 5257289574516412244, func() int32 { return int32(fn()) }) // slot "stepEnabled"
}

func (this *QSpinBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QSpinBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QSpinBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QSpinBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QSpinBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QSpinBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QSpinBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QSpinBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QSpinBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QSpinBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QSpinBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QSpinBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QSpinBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QSpinBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QSpinBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QSpinBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QSpinBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QSpinBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QSpinBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QSpinBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QSpinBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QSpinBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QSpinBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSpinBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSpinBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSpinBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSpinBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSpinBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSpinBox_delete(this.h)
	}
	this.UnsafeRelease()
}
