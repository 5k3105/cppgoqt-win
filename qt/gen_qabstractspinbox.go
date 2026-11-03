package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractSpinBox_chain uintptr
	addr_QAbstractSpinBox_new uintptr
	addr_QAbstractSpinBox_new2 uintptr
	addr_QAbstractSpinBox_tr uintptr
	addr_QAbstractSpinBox_buttonSymbols uintptr
	addr_QAbstractSpinBox_setButtonSymbols uintptr
	addr_QAbstractSpinBox_setCorrectionMode uintptr
	addr_QAbstractSpinBox_correctionMode uintptr
	addr_QAbstractSpinBox_hasAcceptableInput uintptr
	addr_QAbstractSpinBox_text uintptr
	addr_QAbstractSpinBox_specialValueText uintptr
	addr_QAbstractSpinBox_setSpecialValueText uintptr
	addr_QAbstractSpinBox_wrapping uintptr
	addr_QAbstractSpinBox_setWrapping uintptr
	addr_QAbstractSpinBox_setReadOnly uintptr
	addr_QAbstractSpinBox_isReadOnly uintptr
	addr_QAbstractSpinBox_setKeyboardTracking uintptr
	addr_QAbstractSpinBox_keyboardTracking uintptr
	addr_QAbstractSpinBox_setAlignment uintptr
	addr_QAbstractSpinBox_alignment uintptr
	addr_QAbstractSpinBox_setFrame uintptr
	addr_QAbstractSpinBox_hasFrame uintptr
	addr_QAbstractSpinBox_setAccelerated uintptr
	addr_QAbstractSpinBox_isAccelerated uintptr
	addr_QAbstractSpinBox_setGroupSeparatorShown uintptr
	addr_QAbstractSpinBox_isGroupSeparatorShown uintptr
	addr_QAbstractSpinBox_interpretText uintptr
	addr_QAbstractSpinBox_validate uintptr
	addr_QAbstractSpinBox_fixup uintptr
	addr_QAbstractSpinBox_stepBy uintptr
	addr_QAbstractSpinBox_stepUp uintptr
	addr_QAbstractSpinBox_stepDown uintptr
	addr_QAbstractSpinBox_selectAll uintptr
	addr_QAbstractSpinBox_clear uintptr
	addr_QAbstractSpinBox_lineEdit uintptr
	addr_QAbstractSpinBox_setLineEdit uintptr
	addr_QAbstractSpinBox_editingFinished uintptr
	addr_QAbstractSpinBox_returnPressed uintptr
	addr_QAbstractSpinBox_setTabOrder uintptr
	addr_QAbstractSpinBox_mouseGrabber uintptr
	addr_QAbstractSpinBox_keyboardGrabber uintptr
	addr_QAbstractSpinBox_find uintptr
	addr_QAbstractSpinBox_createWindowContainer uintptr
	addr_QAbstractSpinBox_updateMicroFocus uintptr
	addr_QAbstractSpinBox_create uintptr
	addr_QAbstractSpinBox_destroy uintptr
	addr_QAbstractSpinBox_focusNextChild uintptr
	addr_QAbstractSpinBox_focusPreviousChild uintptr
	addr_QAbstractSpinBox_connect uintptr
	addr_QAbstractSpinBox_connect2 uintptr
	addr_QAbstractSpinBox_sender uintptr
	addr_QAbstractSpinBox_senderSignalIndex uintptr
	addr_QAbstractSpinBox_receivers uintptr
	addr_QAbstractSpinBox_isSignalConnected uintptr
	pg_QAbstractSpinBox_devicePixelRatioFScale func() float64
	pg_QAbstractSpinBox_encodeMetricF func(int32, float64) int32
	pg_QAbstractSpinBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QAbstractSpinBox_tr2 uintptr
	addr_QAbstractSpinBox_tr3 uintptr
	addr_QAbstractSpinBox_createWindowContainer2 uintptr
	addr_QAbstractSpinBox_createWindowContainer3 uintptr
	addr_QAbstractSpinBox_updateMicroFocus2 uintptr
	addr_QAbstractSpinBox_create2 uintptr
	addr_QAbstractSpinBox_create3 uintptr
	addr_QAbstractSpinBox_create4 uintptr
	addr_QAbstractSpinBox_destroy2 uintptr
	addr_QAbstractSpinBox_destroy3 uintptr
	addr_QAbstractSpinBox_connect3 uintptr
	addr_QAbstractSpinBox_connect4 uintptr
	addr_QAbstractSpinBox_validate_default uintptr
	addr_QAbstractSpinBox_fixup_default uintptr
	addr_QAbstractSpinBox_stepBy_default uintptr
	addr_QAbstractSpinBox_clear_default uintptr
	addr_QAbstractSpinBox_metaObject_default uintptr
	addr_QAbstractSpinBox_sizeHint_default uintptr
	addr_QAbstractSpinBox_minimumSizeHint_default uintptr
	addr_QAbstractSpinBox_event_default uintptr
	addr_QAbstractSpinBox_inputMethodQuery_default uintptr
	addr_QAbstractSpinBox_resizeEvent_default uintptr
	addr_QAbstractSpinBox_keyPressEvent_default uintptr
	addr_QAbstractSpinBox_keyReleaseEvent_default uintptr
	addr_QAbstractSpinBox_wheelEvent_default uintptr
	addr_QAbstractSpinBox_focusInEvent_default uintptr
	addr_QAbstractSpinBox_focusOutEvent_default uintptr
	addr_QAbstractSpinBox_contextMenuEvent_default uintptr
	addr_QAbstractSpinBox_changeEvent_default uintptr
	addr_QAbstractSpinBox_closeEvent_default uintptr
	addr_QAbstractSpinBox_hideEvent_default uintptr
	addr_QAbstractSpinBox_mousePressEvent_default uintptr
	addr_QAbstractSpinBox_mouseReleaseEvent_default uintptr
	addr_QAbstractSpinBox_mouseMoveEvent_default uintptr
	addr_QAbstractSpinBox_timerEvent_default uintptr
	addr_QAbstractSpinBox_paintEvent_default uintptr
	addr_QAbstractSpinBox_showEvent_default uintptr
	addr_QAbstractSpinBox_initStyleOption_default uintptr
	addr_QAbstractSpinBox_stepEnabled_default uintptr
	addr_QAbstractSpinBox_devType_default uintptr
	addr_QAbstractSpinBox_setVisible_default uintptr
	addr_QAbstractSpinBox_heightForWidth_default uintptr
	addr_QAbstractSpinBox_hasHeightForWidth_default uintptr
	addr_QAbstractSpinBox_paintEngine_default uintptr
	addr_QAbstractSpinBox_mouseDoubleClickEvent_default uintptr
	addr_QAbstractSpinBox_enterEvent_default uintptr
	addr_QAbstractSpinBox_leaveEvent_default uintptr
	addr_QAbstractSpinBox_moveEvent_default uintptr
	addr_QAbstractSpinBox_tabletEvent_default uintptr
	addr_QAbstractSpinBox_actionEvent_default uintptr
	addr_QAbstractSpinBox_dragEnterEvent_default uintptr
	addr_QAbstractSpinBox_dragMoveEvent_default uintptr
	addr_QAbstractSpinBox_dragLeaveEvent_default uintptr
	addr_QAbstractSpinBox_dropEvent_default uintptr
	addr_QAbstractSpinBox_metric_default uintptr
	addr_QAbstractSpinBox_initPainter_default uintptr
	addr_QAbstractSpinBox_redirected_default uintptr
	addr_QAbstractSpinBox_sharedPainter_default uintptr
	addr_QAbstractSpinBox_inputMethodEvent_default uintptr
	addr_QAbstractSpinBox_focusNextPrevChild_default uintptr
	addr_QAbstractSpinBox_eventFilter_default uintptr
	addr_QAbstractSpinBox_childEvent_default uintptr
	addr_QAbstractSpinBox_customEvent_default uintptr
	addr_QAbstractSpinBox_connectNotify_default uintptr
	addr_QAbstractSpinBox_disconnectNotify_default uintptr
	pg_QAbstractSpinBox_connect_editingFinished func(uintptr, int64)
	pg_QAbstractSpinBox_connect_returnPressed func(uintptr, int64)
	pg_QAbstractSpinBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QAbstractSpinBox_connect_windowIconChanged func(uintptr, int64)
	pg_QAbstractSpinBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QAbstractSpinBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QAbstractSpinBox_connect_destroyed func(uintptr, int64)
	pg_QAbstractSpinBox_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractSpinBox_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractSpinBox__disconnect func(int64)
	pg_QAbstractSpinBox__set_override func(uintptr, int32)
	pg_QAbstractSpinBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractSpinBox_chain, "QAbstractSpinBox_chain")
	symaddr(&addr_QAbstractSpinBox_new, "QAbstractSpinBox_new")
	symaddr(&addr_QAbstractSpinBox_new2, "QAbstractSpinBox_new2")
	symaddr(&addr_QAbstractSpinBox_tr, "QAbstractSpinBox_tr")
	symaddr(&addr_QAbstractSpinBox_buttonSymbols, "QAbstractSpinBox_buttonSymbols")
	symaddr(&addr_QAbstractSpinBox_setButtonSymbols, "QAbstractSpinBox_setButtonSymbols")
	symaddr(&addr_QAbstractSpinBox_setCorrectionMode, "QAbstractSpinBox_setCorrectionMode")
	symaddr(&addr_QAbstractSpinBox_correctionMode, "QAbstractSpinBox_correctionMode")
	symaddr(&addr_QAbstractSpinBox_hasAcceptableInput, "QAbstractSpinBox_hasAcceptableInput")
	symaddr(&addr_QAbstractSpinBox_text, "QAbstractSpinBox_text")
	symaddr(&addr_QAbstractSpinBox_specialValueText, "QAbstractSpinBox_specialValueText")
	symaddr(&addr_QAbstractSpinBox_setSpecialValueText, "QAbstractSpinBox_setSpecialValueText")
	symaddr(&addr_QAbstractSpinBox_wrapping, "QAbstractSpinBox_wrapping")
	symaddr(&addr_QAbstractSpinBox_setWrapping, "QAbstractSpinBox_setWrapping")
	symaddr(&addr_QAbstractSpinBox_setReadOnly, "QAbstractSpinBox_setReadOnly")
	symaddr(&addr_QAbstractSpinBox_isReadOnly, "QAbstractSpinBox_isReadOnly")
	symaddr(&addr_QAbstractSpinBox_setKeyboardTracking, "QAbstractSpinBox_setKeyboardTracking")
	symaddr(&addr_QAbstractSpinBox_keyboardTracking, "QAbstractSpinBox_keyboardTracking")
	symaddr(&addr_QAbstractSpinBox_setAlignment, "QAbstractSpinBox_setAlignment")
	symaddr(&addr_QAbstractSpinBox_alignment, "QAbstractSpinBox_alignment")
	symaddr(&addr_QAbstractSpinBox_setFrame, "QAbstractSpinBox_setFrame")
	symaddr(&addr_QAbstractSpinBox_hasFrame, "QAbstractSpinBox_hasFrame")
	symaddr(&addr_QAbstractSpinBox_setAccelerated, "QAbstractSpinBox_setAccelerated")
	symaddr(&addr_QAbstractSpinBox_isAccelerated, "QAbstractSpinBox_isAccelerated")
	symaddr(&addr_QAbstractSpinBox_setGroupSeparatorShown, "QAbstractSpinBox_setGroupSeparatorShown")
	symaddr(&addr_QAbstractSpinBox_isGroupSeparatorShown, "QAbstractSpinBox_isGroupSeparatorShown")
	symaddr(&addr_QAbstractSpinBox_interpretText, "QAbstractSpinBox_interpretText")
	symaddr(&addr_QAbstractSpinBox_validate, "QAbstractSpinBox_validate")
	symaddr(&addr_QAbstractSpinBox_fixup, "QAbstractSpinBox_fixup")
	symaddr(&addr_QAbstractSpinBox_stepBy, "QAbstractSpinBox_stepBy")
	symaddr(&addr_QAbstractSpinBox_stepUp, "QAbstractSpinBox_stepUp")
	symaddr(&addr_QAbstractSpinBox_stepDown, "QAbstractSpinBox_stepDown")
	symaddr(&addr_QAbstractSpinBox_selectAll, "QAbstractSpinBox_selectAll")
	symaddr(&addr_QAbstractSpinBox_clear, "QAbstractSpinBox_clear")
	symaddr(&addr_QAbstractSpinBox_lineEdit, "QAbstractSpinBox_lineEdit")
	symaddr(&addr_QAbstractSpinBox_setLineEdit, "QAbstractSpinBox_setLineEdit")
	symaddr(&addr_QAbstractSpinBox_editingFinished, "QAbstractSpinBox_editingFinished")
	symaddr(&addr_QAbstractSpinBox_returnPressed, "QAbstractSpinBox_returnPressed")
	symaddr(&addr_QAbstractSpinBox_setTabOrder, "QAbstractSpinBox_setTabOrder")
	symaddr(&addr_QAbstractSpinBox_mouseGrabber, "QAbstractSpinBox_mouseGrabber")
	symaddr(&addr_QAbstractSpinBox_keyboardGrabber, "QAbstractSpinBox_keyboardGrabber")
	symaddr(&addr_QAbstractSpinBox_find, "QAbstractSpinBox_find")
	symaddr(&addr_QAbstractSpinBox_createWindowContainer, "QAbstractSpinBox_createWindowContainer")
	symaddr(&addr_QAbstractSpinBox_updateMicroFocus, "QAbstractSpinBox_updateMicroFocus")
	symaddr(&addr_QAbstractSpinBox_create, "QAbstractSpinBox_create")
	symaddr(&addr_QAbstractSpinBox_destroy, "QAbstractSpinBox_destroy")
	symaddr(&addr_QAbstractSpinBox_focusNextChild, "QAbstractSpinBox_focusNextChild")
	symaddr(&addr_QAbstractSpinBox_focusPreviousChild, "QAbstractSpinBox_focusPreviousChild")
	symaddr(&addr_QAbstractSpinBox_connect, "QAbstractSpinBox_connect")
	symaddr(&addr_QAbstractSpinBox_connect2, "QAbstractSpinBox_connect2")
	symaddr(&addr_QAbstractSpinBox_sender, "QAbstractSpinBox_sender")
	symaddr(&addr_QAbstractSpinBox_senderSignalIndex, "QAbstractSpinBox_senderSignalIndex")
	symaddr(&addr_QAbstractSpinBox_receivers, "QAbstractSpinBox_receivers")
	symaddr(&addr_QAbstractSpinBox_isSignalConnected, "QAbstractSpinBox_isSignalConnected")
	reg(&pg_QAbstractSpinBox_devicePixelRatioFScale, "QAbstractSpinBox_devicePixelRatioFScale")
	reg(&pg_QAbstractSpinBox_encodeMetricF, "QAbstractSpinBox_encodeMetricF")
	reg(&pg_QAbstractSpinBox_getDecodedMetricF, "QAbstractSpinBox_getDecodedMetricF")
	symaddr(&addr_QAbstractSpinBox_tr2, "QAbstractSpinBox_tr2")
	symaddr(&addr_QAbstractSpinBox_tr3, "QAbstractSpinBox_tr3")
	symaddr(&addr_QAbstractSpinBox_createWindowContainer2, "QAbstractSpinBox_createWindowContainer2")
	symaddr(&addr_QAbstractSpinBox_createWindowContainer3, "QAbstractSpinBox_createWindowContainer3")
	symaddr(&addr_QAbstractSpinBox_updateMicroFocus2, "QAbstractSpinBox_updateMicroFocus2")
	symaddr(&addr_QAbstractSpinBox_create2, "QAbstractSpinBox_create2")
	symaddr(&addr_QAbstractSpinBox_create3, "QAbstractSpinBox_create3")
	symaddr(&addr_QAbstractSpinBox_create4, "QAbstractSpinBox_create4")
	symaddr(&addr_QAbstractSpinBox_destroy2, "QAbstractSpinBox_destroy2")
	symaddr(&addr_QAbstractSpinBox_destroy3, "QAbstractSpinBox_destroy3")
	symaddr(&addr_QAbstractSpinBox_connect3, "QAbstractSpinBox_connect3")
	symaddr(&addr_QAbstractSpinBox_connect4, "QAbstractSpinBox_connect4")
	symaddr(&addr_QAbstractSpinBox_validate_default, "QAbstractSpinBox_validate_default")
	symaddr(&addr_QAbstractSpinBox_fixup_default, "QAbstractSpinBox_fixup_default")
	symaddr(&addr_QAbstractSpinBox_stepBy_default, "QAbstractSpinBox_stepBy_default")
	symaddr(&addr_QAbstractSpinBox_clear_default, "QAbstractSpinBox_clear_default")
	symaddr(&addr_QAbstractSpinBox_metaObject_default, "QAbstractSpinBox_metaObject_default")
	symaddr(&addr_QAbstractSpinBox_sizeHint_default, "QAbstractSpinBox_sizeHint_default")
	symaddr(&addr_QAbstractSpinBox_minimumSizeHint_default, "QAbstractSpinBox_minimumSizeHint_default")
	symaddr(&addr_QAbstractSpinBox_event_default, "QAbstractSpinBox_event_default")
	symaddr(&addr_QAbstractSpinBox_inputMethodQuery_default, "QAbstractSpinBox_inputMethodQuery_default")
	symaddr(&addr_QAbstractSpinBox_resizeEvent_default, "QAbstractSpinBox_resizeEvent_default")
	symaddr(&addr_QAbstractSpinBox_keyPressEvent_default, "QAbstractSpinBox_keyPressEvent_default")
	symaddr(&addr_QAbstractSpinBox_keyReleaseEvent_default, "QAbstractSpinBox_keyReleaseEvent_default")
	symaddr(&addr_QAbstractSpinBox_wheelEvent_default, "QAbstractSpinBox_wheelEvent_default")
	symaddr(&addr_QAbstractSpinBox_focusInEvent_default, "QAbstractSpinBox_focusInEvent_default")
	symaddr(&addr_QAbstractSpinBox_focusOutEvent_default, "QAbstractSpinBox_focusOutEvent_default")
	symaddr(&addr_QAbstractSpinBox_contextMenuEvent_default, "QAbstractSpinBox_contextMenuEvent_default")
	symaddr(&addr_QAbstractSpinBox_changeEvent_default, "QAbstractSpinBox_changeEvent_default")
	symaddr(&addr_QAbstractSpinBox_closeEvent_default, "QAbstractSpinBox_closeEvent_default")
	symaddr(&addr_QAbstractSpinBox_hideEvent_default, "QAbstractSpinBox_hideEvent_default")
	symaddr(&addr_QAbstractSpinBox_mousePressEvent_default, "QAbstractSpinBox_mousePressEvent_default")
	symaddr(&addr_QAbstractSpinBox_mouseReleaseEvent_default, "QAbstractSpinBox_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractSpinBox_mouseMoveEvent_default, "QAbstractSpinBox_mouseMoveEvent_default")
	symaddr(&addr_QAbstractSpinBox_timerEvent_default, "QAbstractSpinBox_timerEvent_default")
	symaddr(&addr_QAbstractSpinBox_paintEvent_default, "QAbstractSpinBox_paintEvent_default")
	symaddr(&addr_QAbstractSpinBox_showEvent_default, "QAbstractSpinBox_showEvent_default")
	symaddr(&addr_QAbstractSpinBox_initStyleOption_default, "QAbstractSpinBox_initStyleOption_default")
	symaddr(&addr_QAbstractSpinBox_stepEnabled_default, "QAbstractSpinBox_stepEnabled_default")
	symaddr(&addr_QAbstractSpinBox_devType_default, "QAbstractSpinBox_devType_default")
	symaddr(&addr_QAbstractSpinBox_setVisible_default, "QAbstractSpinBox_setVisible_default")
	symaddr(&addr_QAbstractSpinBox_heightForWidth_default, "QAbstractSpinBox_heightForWidth_default")
	symaddr(&addr_QAbstractSpinBox_hasHeightForWidth_default, "QAbstractSpinBox_hasHeightForWidth_default")
	symaddr(&addr_QAbstractSpinBox_paintEngine_default, "QAbstractSpinBox_paintEngine_default")
	symaddr(&addr_QAbstractSpinBox_mouseDoubleClickEvent_default, "QAbstractSpinBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractSpinBox_enterEvent_default, "QAbstractSpinBox_enterEvent_default")
	symaddr(&addr_QAbstractSpinBox_leaveEvent_default, "QAbstractSpinBox_leaveEvent_default")
	symaddr(&addr_QAbstractSpinBox_moveEvent_default, "QAbstractSpinBox_moveEvent_default")
	symaddr(&addr_QAbstractSpinBox_tabletEvent_default, "QAbstractSpinBox_tabletEvent_default")
	symaddr(&addr_QAbstractSpinBox_actionEvent_default, "QAbstractSpinBox_actionEvent_default")
	symaddr(&addr_QAbstractSpinBox_dragEnterEvent_default, "QAbstractSpinBox_dragEnterEvent_default")
	symaddr(&addr_QAbstractSpinBox_dragMoveEvent_default, "QAbstractSpinBox_dragMoveEvent_default")
	symaddr(&addr_QAbstractSpinBox_dragLeaveEvent_default, "QAbstractSpinBox_dragLeaveEvent_default")
	symaddr(&addr_QAbstractSpinBox_dropEvent_default, "QAbstractSpinBox_dropEvent_default")
	symaddr(&addr_QAbstractSpinBox_metric_default, "QAbstractSpinBox_metric_default")
	symaddr(&addr_QAbstractSpinBox_initPainter_default, "QAbstractSpinBox_initPainter_default")
	symaddr(&addr_QAbstractSpinBox_redirected_default, "QAbstractSpinBox_redirected_default")
	symaddr(&addr_QAbstractSpinBox_sharedPainter_default, "QAbstractSpinBox_sharedPainter_default")
	symaddr(&addr_QAbstractSpinBox_inputMethodEvent_default, "QAbstractSpinBox_inputMethodEvent_default")
	symaddr(&addr_QAbstractSpinBox_focusNextPrevChild_default, "QAbstractSpinBox_focusNextPrevChild_default")
	symaddr(&addr_QAbstractSpinBox_eventFilter_default, "QAbstractSpinBox_eventFilter_default")
	symaddr(&addr_QAbstractSpinBox_childEvent_default, "QAbstractSpinBox_childEvent_default")
	symaddr(&addr_QAbstractSpinBox_customEvent_default, "QAbstractSpinBox_customEvent_default")
	symaddr(&addr_QAbstractSpinBox_connectNotify_default, "QAbstractSpinBox_connectNotify_default")
	symaddr(&addr_QAbstractSpinBox_disconnectNotify_default, "QAbstractSpinBox_disconnectNotify_default")
	reg(&pg_QAbstractSpinBox_connect_editingFinished, "QAbstractSpinBox_connect_editingFinished")
	reg(&pg_QAbstractSpinBox_connect_returnPressed, "QAbstractSpinBox_connect_returnPressed")
	reg(&pg_QAbstractSpinBox_connect_windowTitleChanged, "QAbstractSpinBox_connect_windowTitleChanged")
	reg(&pg_QAbstractSpinBox_connect_windowIconChanged, "QAbstractSpinBox_connect_windowIconChanged")
	reg(&pg_QAbstractSpinBox_connect_windowIconTextChanged, "QAbstractSpinBox_connect_windowIconTextChanged")
	reg(&pg_QAbstractSpinBox_connect_customContextMenuRequested, "QAbstractSpinBox_connect_customContextMenuRequested")
	reg(&pg_QAbstractSpinBox_connect_destroyed, "QAbstractSpinBox_connect_destroyed")
	reg(&pg_QAbstractSpinBox_connect_destroyed2, "QAbstractSpinBox_connect_destroyed2")
	reg(&pg_QAbstractSpinBox_connect_objectNameChanged, "QAbstractSpinBox_connect_objectNameChanged")
	reg(&pg_QAbstractSpinBox__disconnect, "QAbstractSpinBox__disconnect")
	reg(&pg_QAbstractSpinBox__set_override, "QAbstractSpinBox__set_override")
	reg(&pg_QAbstractSpinBox_delete, "QAbstractSpinBox_delete")
}

type QAbstractSpinBox struct {
	h uintptr
	*QWidget
}

func bindQAbstractSpinBox(h uintptr) *QAbstractSpinBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractSpinBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractSpinBox{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractSpinBox(h uintptr) *QAbstractSpinBox { return bindQAbstractSpinBox(h) }

func (this *QAbstractSpinBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractSpinBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractSpinBox(h unsafe.Pointer) *QAbstractSpinBox { return bindQAbstractSpinBox(uintptr(h)) }
func UnsafeNewQAbstractSpinBoxOwned(h unsafe.Pointer) *QAbstractSpinBox { return newownedQAbstractSpinBox(uintptr(h)) }

func (this *QAbstractSpinBox) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQAbstractSpinBox(parent *QWidget) *QAbstractSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAbstractSpinBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractSpinBox_virtIdx, QAbstractSpinBox_enable)
	return __o
}

func NewQAbstractSpinBox2() *QAbstractSpinBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_new2, uintptr(__goid))
	__o := newownedQAbstractSpinBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractSpinBox_virtIdx, QAbstractSpinBox_enable)
	return __o
}

func QAbstractSpinBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractSpinBox) ButtonSymbols() QAbstractSpinBox__ButtonSymbols {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_buttonSymbols, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__ButtonSymbols(__r)
}

func (this *QAbstractSpinBox) SetButtonSymbols(bs QAbstractSpinBox__ButtonSymbols) {
	purego.SyscallN(addr_QAbstractSpinBox_setButtonSymbols, this.h, uintptr(bs))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) SetCorrectionMode(cm QAbstractSpinBox__CorrectionMode) {
	purego.SyscallN(addr_QAbstractSpinBox_setCorrectionMode, this.h, uintptr(cm))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) CorrectionMode() QAbstractSpinBox__CorrectionMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_correctionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__CorrectionMode(__r)
}

func (this *QAbstractSpinBox) HasAcceptableInput() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_hasAcceptableInput, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAbstractSpinBox) SpecialValueText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_specialValueText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAbstractSpinBox) SetSpecialValueText(txt string) {
	var txt_p uintptr
	if len(txt) > 0 {
		txt_p = uintptr(unsafe.Pointer(unsafe.StringData(txt)))
	}
	purego.SyscallN(addr_QAbstractSpinBox_setSpecialValueText, this.h, txt_p, uintptr(len(txt)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(txt)
}

func (this *QAbstractSpinBox) Wrapping() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_wrapping, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) SetWrapping(w bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setWrapping, this.h, uintptr(b2i(w)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) SetReadOnly(r bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setReadOnly, this.h, uintptr(b2i(r)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) SetKeyboardTracking(kt bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setKeyboardTracking, this.h, uintptr(b2i(kt)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) KeyboardTracking() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_keyboardTracking, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) SetAlignment(flag AlignmentFlag) {
	purego.SyscallN(addr_QAbstractSpinBox_setAlignment, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QAbstractSpinBox) SetFrame(arg1 bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setFrame, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) HasFrame() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_hasFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) SetAccelerated(on bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setAccelerated, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) IsAccelerated() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_isAccelerated, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) SetGroupSeparatorShown(shown bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setGroupSeparatorShown, this.h, uintptr(b2i(shown)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) IsGroupSeparatorShown() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_isGroupSeparatorShown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) InterpretText() {
	purego.SyscallN(addr_QAbstractSpinBox_interpretText, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Validate(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_validate, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QAbstractSpinBox) Fixup(input *QString) {
	purego.SyscallN(addr_QAbstractSpinBox_fixup, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QAbstractSpinBox) StepBy(steps int) {
	purego.SyscallN(addr_QAbstractSpinBox_stepBy, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) StepUp() {
	purego.SyscallN(addr_QAbstractSpinBox_stepUp, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) StepDown() {
	purego.SyscallN(addr_QAbstractSpinBox_stepDown, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) SelectAll() {
	purego.SyscallN(addr_QAbstractSpinBox_selectAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Clear() {
	purego.SyscallN(addr_QAbstractSpinBox_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QAbstractSpinBox) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QAbstractSpinBox_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func (this *QAbstractSpinBox) EditingFinished() {
	purego.SyscallN(addr_QAbstractSpinBox_editingFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) ReturnPressed() {
	purego.SyscallN(addr_QAbstractSpinBox_returnPressed, this.h)
	runtime.KeepAlive(this)
}

func QAbstractSpinBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QAbstractSpinBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QAbstractSpinBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_mouseGrabber)
	return bindQWidget(__r)
}

func QAbstractSpinBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QAbstractSpinBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QAbstractSpinBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QAbstractSpinBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractSpinBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Create() {
	purego.SyscallN(addr_QAbstractSpinBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Destroy() {
	purego.SyscallN(addr_QAbstractSpinBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAbstractSpinBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractSpinBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractSpinBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractSpinBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSpinBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractSpinBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractSpinBox_DevicePixelRatioFScale() float64 {
	__r := pg_QAbstractSpinBox_devicePixelRatioFScale()
	return __r
}

func QAbstractSpinBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QAbstractSpinBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QAbstractSpinBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QAbstractSpinBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QAbstractSpinBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractSpinBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractSpinBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QAbstractSpinBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QAbstractSpinBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QAbstractSpinBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QAbstractSpinBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QAbstractSpinBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QAbstractSpinBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QAbstractSpinBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QAbstractSpinBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QAbstractSpinBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractSpinBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractSpinBox) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QAbstractSpinBox) FixupDefault(input *QString) {
	purego.SyscallN(addr_QAbstractSpinBox_fixup_default, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QAbstractSpinBox) StepByDefault(steps int) {
	purego.SyscallN(addr_QAbstractSpinBox_stepBy_default, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) ClearDefault() {
	purego.SyscallN(addr_QAbstractSpinBox_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractSpinBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractSpinBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractSpinBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractSpinBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractSpinBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) InitStyleOptionDefault(option *QStyleOptionSpinBox) {
	purego.SyscallN(addr_QAbstractSpinBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QAbstractSpinBox) StepEnabledDefault() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_stepEnabled_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QAbstractSpinBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSpinBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QAbstractSpinBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAbstractSpinBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSpinBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QAbstractSpinBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractSpinBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QAbstractSpinBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QAbstractSpinBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QAbstractSpinBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QAbstractSpinBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QAbstractSpinBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractSpinBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractSpinBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractSpinBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractSpinBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractSpinBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractSpinBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractSpinBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractSpinBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractSpinBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractSpinBox) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSpinBox_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSpinBox_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QAbstractSpinBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QAbstractSpinBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QAbstractSpinBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QAbstractSpinBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractSpinBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractSpinBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractSpinBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractSpinBox) Disconnect(handle int64) {
	pg_QAbstractSpinBox__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractSpinBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case 7345853097340709638:
		return 4
	case 6784179835071232555:
		return 5
	case -4855206404746966715:
		return 6
	case 7085718643589954634:
		return 7
	case -778567940572639070:
		return 8
	case -4302917285655014661:
		return 9
	case 1023791977679793185:
		return 10
	case 772598475950050299:
		return 11
	case 1919988340934511316:
		return 12
	case 975043767033552552:
		return 13
	case -385596010576048993:
		return 14
	case -6698043222227006015:
		return 15
	case 4959376063826165943:
		return 16
	case -1989937050019464783:
		return 17
	case -199718707876719129:
		return 18
	case -332090898078543319:
		return 19
	case -4415920049297795213:
		return 20
	case -4255494004099806581:
		return 21
	case -5409927581812374498:
		return 22
	case 5290442163167458465:
		return 23
	case -6825369174751916:
		return 24
	case -488693769771308531:
		return 25
	case 5257289574516412244:
		return 26
	case -8659679893868524044:
		return 27
	case 2110182719375245575:
		return 28
	case 4311534060495737339:
		return 29
	case -8266762397995423915:
		return 30
	case 6398007404756346381:
		return 31
	case -4029073905697023887:
		return 32
	case 6316270346668237895:
		return 33
	case -551070009623711640:
		return 34
	case -9161949150489525396:
		return 35
	case -212068866860701141:
		return 36
	case 1697444140088555581:
		return 37
	case 206122715919917589:
		return 38
	case -1781545137352745978:
		return 39
	case -3139904158373725006:
		return 40
	case 9107058682754591716:
		return 41
	case -2039166586246443706:
		return 42
	case 427629649267311777:
		return 43
	case 6477500296980050210:
		return 44
	case 4127240633310655048:
		return 45
	case 5792771328741774991:
		return 46
	case 5438819661942843222:
		return 47
	case 5898328517505784191:
		return 48
	case 5217623514630858297:
		return 49
	case -1185315367526857833:
		return 50
	case 3769868229123868474:
		return 51
	case 7379029982170503008:
		return 52
	case -568369772982215962:
		return 53
	}
	return -1
}

func QAbstractSpinBox_enable(h uintptr, idx int32) { pg_QAbstractSpinBox__set_override(h, idx) }

func (this *QAbstractSpinBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractSpinBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractSpinBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QAbstractSpinBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractSpinBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractSpinBox) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QAbstractSpinBox) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(input uintptr) { fn(UnsafeNewQString(unsafe.Pointer(input))) }) // slot "fixup"
}

func (this *QAbstractSpinBox) OnStepBy(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 7085718643589954634, func(steps int32) { fn(int(steps)) }) // slot "stepBy"
}

func (this *QAbstractSpinBox) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QAbstractSpinBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QAbstractSpinBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QAbstractSpinBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractSpinBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QAbstractSpinBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QAbstractSpinBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QAbstractSpinBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QAbstractSpinBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QAbstractSpinBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QAbstractSpinBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QAbstractSpinBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QAbstractSpinBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractSpinBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractSpinBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractSpinBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QAbstractSpinBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QAbstractSpinBox) OnInitStyleOption(fn func(*QStyleOptionSpinBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSpinBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QAbstractSpinBox) OnStepEnabled(fn func() QAbstractSpinBox__StepEnabledFlag) {
	registerVirt(this.UnsafeGoID(), 5257289574516412244, func() int32 { return int32(fn()) }) // slot "stepEnabled"
}

func (this *QAbstractSpinBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QAbstractSpinBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QAbstractSpinBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QAbstractSpinBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QAbstractSpinBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QAbstractSpinBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractSpinBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QAbstractSpinBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QAbstractSpinBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QAbstractSpinBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QAbstractSpinBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QAbstractSpinBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QAbstractSpinBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QAbstractSpinBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractSpinBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QAbstractSpinBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QAbstractSpinBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QAbstractSpinBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QAbstractSpinBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QAbstractSpinBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QAbstractSpinBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QAbstractSpinBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QAbstractSpinBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractSpinBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractSpinBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractSpinBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractSpinBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractSpinBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractSpinBox_delete(this.h)
	}
	this.UnsafeRelease()
}
