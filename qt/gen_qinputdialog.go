package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QInputDialog_chain uintptr
	addr_QInputDialog_new uintptr
	addr_QInputDialog_new2 uintptr
	addr_QInputDialog_new3 uintptr
	addr_QInputDialog_tr uintptr
	addr_QInputDialog_setInputMode uintptr
	addr_QInputDialog_inputMode uintptr
	addr_QInputDialog_setLabelText uintptr
	addr_QInputDialog_labelText uintptr
	addr_QInputDialog_setOption uintptr
	addr_QInputDialog_testOption uintptr
	addr_QInputDialog_setOptions uintptr
	addr_QInputDialog_options uintptr
	addr_QInputDialog_setTextValue uintptr
	addr_QInputDialog_textValue uintptr
	addr_QInputDialog_setTextEchoMode uintptr
	addr_QInputDialog_textEchoMode uintptr
	addr_QInputDialog_setComboBoxEditable uintptr
	addr_QInputDialog_isComboBoxEditable uintptr
	addr_QInputDialog_setComboBoxItems uintptr
	addr_QInputDialog_comboBoxItems uintptr
	addr_QInputDialog_setIntValue uintptr
	addr_QInputDialog_intValue uintptr
	addr_QInputDialog_setIntMinimum uintptr
	addr_QInputDialog_intMinimum uintptr
	addr_QInputDialog_setIntMaximum uintptr
	addr_QInputDialog_intMaximum uintptr
	addr_QInputDialog_setIntRange uintptr
	addr_QInputDialog_setIntStep uintptr
	addr_QInputDialog_intStep uintptr
	pg_QInputDialog_setDoubleValue func(uintptr, float64)
	pg_QInputDialog_doubleValue func(uintptr) float64
	pg_QInputDialog_setDoubleMinimum func(uintptr, float64)
	pg_QInputDialog_doubleMinimum func(uintptr) float64
	pg_QInputDialog_setDoubleMaximum func(uintptr, float64)
	pg_QInputDialog_doubleMaximum func(uintptr) float64
	pg_QInputDialog_setDoubleRange func(uintptr, float64, float64)
	addr_QInputDialog_setDoubleDecimals uintptr
	addr_QInputDialog_doubleDecimals uintptr
	addr_QInputDialog_setOkButtonText uintptr
	addr_QInputDialog_okButtonText uintptr
	addr_QInputDialog_setCancelButtonText uintptr
	addr_QInputDialog_cancelButtonText uintptr
	addr_QInputDialog_open uintptr
	addr_QInputDialog_getText uintptr
	addr_QInputDialog_getMultiLineText uintptr
	addr_QInputDialog_getItem uintptr
	addr_QInputDialog_getInt uintptr
	pg_QInputDialog_getDouble func(uintptr, string, int64, string, int64) float64
	pg_QInputDialog_setDoubleStep func(uintptr, float64)
	pg_QInputDialog_doubleStep func(uintptr) float64
	addr_QInputDialog_textValueChanged uintptr
	addr_QInputDialog_textValueSelected uintptr
	addr_QInputDialog_intValueChanged uintptr
	addr_QInputDialog_intValueSelected uintptr
	pg_QInputDialog_doubleValueChanged func(uintptr, float64)
	pg_QInputDialog_doubleValueSelected func(uintptr, float64)
	addr_QInputDialog_adjustPosition uintptr
	addr_QInputDialog_setTabOrder uintptr
	addr_QInputDialog_mouseGrabber uintptr
	addr_QInputDialog_keyboardGrabber uintptr
	addr_QInputDialog_find uintptr
	addr_QInputDialog_createWindowContainer uintptr
	addr_QInputDialog_updateMicroFocus uintptr
	addr_QInputDialog_create uintptr
	addr_QInputDialog_destroy uintptr
	addr_QInputDialog_focusNextChild uintptr
	addr_QInputDialog_focusPreviousChild uintptr
	addr_QInputDialog_connect uintptr
	addr_QInputDialog_connect2 uintptr
	addr_QInputDialog_sender uintptr
	addr_QInputDialog_senderSignalIndex uintptr
	addr_QInputDialog_receivers uintptr
	addr_QInputDialog_isSignalConnected uintptr
	pg_QInputDialog_devicePixelRatioFScale func() float64
	pg_QInputDialog_encodeMetricF func(int32, float64) int32
	pg_QInputDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QInputDialog_tr2 uintptr
	addr_QInputDialog_tr3 uintptr
	addr_QInputDialog_setOption2 uintptr
	addr_QInputDialog_getText2 uintptr
	addr_QInputDialog_getText3 uintptr
	addr_QInputDialog_getText4 uintptr
	addr_QInputDialog_getText5 uintptr
	addr_QInputDialog_getText6 uintptr
	addr_QInputDialog_getMultiLineText2 uintptr
	addr_QInputDialog_getMultiLineText3 uintptr
	addr_QInputDialog_getMultiLineText4 uintptr
	addr_QInputDialog_getMultiLineText5 uintptr
	addr_QInputDialog_getItem2 uintptr
	addr_QInputDialog_getItem3 uintptr
	addr_QInputDialog_getItem4 uintptr
	addr_QInputDialog_getItem5 uintptr
	addr_QInputDialog_getItem6 uintptr
	addr_QInputDialog_getInt2 uintptr
	addr_QInputDialog_getInt3 uintptr
	addr_QInputDialog_getInt4 uintptr
	addr_QInputDialog_getInt5 uintptr
	addr_QInputDialog_getInt6 uintptr
	addr_QInputDialog_getInt7 uintptr
	pg_QInputDialog_getDouble2 func(uintptr, string, int64, string, int64, float64) float64
	pg_QInputDialog_getDouble3 func(uintptr, string, int64, string, int64, float64, float64) float64
	pg_QInputDialog_getDouble4 func(uintptr, string, int64, string, int64, float64, float64, float64) float64
	pg_QInputDialog_getDouble5 func(uintptr, string, int64, string, int64, float64, float64, float64, int32) float64
	pg_QInputDialog_getDouble6 func(uintptr, string, int64, string, int64, float64, float64, float64, int32, uintptr) float64
	pg_QInputDialog_getDouble7 func(uintptr, string, int64, string, int64, float64, float64, float64, int32, uintptr, int32) float64
	pg_QInputDialog_getDouble8 func(uintptr, string, int64, string, int64, float64, float64, float64, int32, uintptr, int32, float64) float64
	addr_QInputDialog_createWindowContainer2 uintptr
	addr_QInputDialog_createWindowContainer3 uintptr
	addr_QInputDialog_updateMicroFocus2 uintptr
	addr_QInputDialog_create2 uintptr
	addr_QInputDialog_create3 uintptr
	addr_QInputDialog_create4 uintptr
	addr_QInputDialog_destroy2 uintptr
	addr_QInputDialog_destroy3 uintptr
	addr_QInputDialog_connect3 uintptr
	addr_QInputDialog_connect4 uintptr
	addr_QInputDialog_metaObject_default uintptr
	addr_QInputDialog_minimumSizeHint_default uintptr
	addr_QInputDialog_sizeHint_default uintptr
	addr_QInputDialog_setVisible_default uintptr
	addr_QInputDialog_done_default uintptr
	addr_QInputDialog_exec_default uintptr
	addr_QInputDialog_accept_default uintptr
	addr_QInputDialog_reject_default uintptr
	addr_QInputDialog_keyPressEvent_default uintptr
	addr_QInputDialog_closeEvent_default uintptr
	addr_QInputDialog_showEvent_default uintptr
	addr_QInputDialog_resizeEvent_default uintptr
	addr_QInputDialog_contextMenuEvent_default uintptr
	addr_QInputDialog_eventFilter_default uintptr
	addr_QInputDialog_devType_default uintptr
	addr_QInputDialog_heightForWidth_default uintptr
	addr_QInputDialog_hasHeightForWidth_default uintptr
	addr_QInputDialog_paintEngine_default uintptr
	addr_QInputDialog_event_default uintptr
	addr_QInputDialog_mousePressEvent_default uintptr
	addr_QInputDialog_mouseReleaseEvent_default uintptr
	addr_QInputDialog_mouseDoubleClickEvent_default uintptr
	addr_QInputDialog_mouseMoveEvent_default uintptr
	addr_QInputDialog_wheelEvent_default uintptr
	addr_QInputDialog_keyReleaseEvent_default uintptr
	addr_QInputDialog_focusInEvent_default uintptr
	addr_QInputDialog_focusOutEvent_default uintptr
	addr_QInputDialog_enterEvent_default uintptr
	addr_QInputDialog_leaveEvent_default uintptr
	addr_QInputDialog_paintEvent_default uintptr
	addr_QInputDialog_moveEvent_default uintptr
	addr_QInputDialog_tabletEvent_default uintptr
	addr_QInputDialog_actionEvent_default uintptr
	addr_QInputDialog_dragEnterEvent_default uintptr
	addr_QInputDialog_dragMoveEvent_default uintptr
	addr_QInputDialog_dragLeaveEvent_default uintptr
	addr_QInputDialog_dropEvent_default uintptr
	addr_QInputDialog_hideEvent_default uintptr
	addr_QInputDialog_changeEvent_default uintptr
	addr_QInputDialog_metric_default uintptr
	addr_QInputDialog_initPainter_default uintptr
	addr_QInputDialog_redirected_default uintptr
	addr_QInputDialog_sharedPainter_default uintptr
	addr_QInputDialog_inputMethodEvent_default uintptr
	addr_QInputDialog_inputMethodQuery_default uintptr
	addr_QInputDialog_focusNextPrevChild_default uintptr
	addr_QInputDialog_timerEvent_default uintptr
	addr_QInputDialog_childEvent_default uintptr
	addr_QInputDialog_customEvent_default uintptr
	addr_QInputDialog_connectNotify_default uintptr
	addr_QInputDialog_disconnectNotify_default uintptr
	pg_QInputDialog_connect_textValueChanged func(uintptr, int64)
	pg_QInputDialog_connect_textValueSelected func(uintptr, int64)
	pg_QInputDialog_connect_intValueChanged func(uintptr, int64)
	pg_QInputDialog_connect_intValueSelected func(uintptr, int64)
	pg_QInputDialog_connect_doubleValueChanged func(uintptr, int64)
	pg_QInputDialog_connect_doubleValueSelected func(uintptr, int64)
	pg_QInputDialog_connect_finished func(uintptr, int64)
	pg_QInputDialog_connect_accepted func(uintptr, int64)
	pg_QInputDialog_connect_rejected func(uintptr, int64)
	pg_QInputDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QInputDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QInputDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QInputDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QInputDialog_connect_destroyed func(uintptr, int64)
	pg_QInputDialog_connect_destroyed2 func(uintptr, int64)
	pg_QInputDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QInputDialog__disconnect func(int64)
	pg_QInputDialog__set_override func(uintptr, int32)
	pg_QInputDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QInputDialog_chain, "QInputDialog_chain")
	symaddr(&addr_QInputDialog_new, "QInputDialog_new")
	symaddr(&addr_QInputDialog_new2, "QInputDialog_new2")
	symaddr(&addr_QInputDialog_new3, "QInputDialog_new3")
	symaddr(&addr_QInputDialog_tr, "QInputDialog_tr")
	symaddr(&addr_QInputDialog_setInputMode, "QInputDialog_setInputMode")
	symaddr(&addr_QInputDialog_inputMode, "QInputDialog_inputMode")
	symaddr(&addr_QInputDialog_setLabelText, "QInputDialog_setLabelText")
	symaddr(&addr_QInputDialog_labelText, "QInputDialog_labelText")
	symaddr(&addr_QInputDialog_setOption, "QInputDialog_setOption")
	symaddr(&addr_QInputDialog_testOption, "QInputDialog_testOption")
	symaddr(&addr_QInputDialog_setOptions, "QInputDialog_setOptions")
	symaddr(&addr_QInputDialog_options, "QInputDialog_options")
	symaddr(&addr_QInputDialog_setTextValue, "QInputDialog_setTextValue")
	symaddr(&addr_QInputDialog_textValue, "QInputDialog_textValue")
	symaddr(&addr_QInputDialog_setTextEchoMode, "QInputDialog_setTextEchoMode")
	symaddr(&addr_QInputDialog_textEchoMode, "QInputDialog_textEchoMode")
	symaddr(&addr_QInputDialog_setComboBoxEditable, "QInputDialog_setComboBoxEditable")
	symaddr(&addr_QInputDialog_isComboBoxEditable, "QInputDialog_isComboBoxEditable")
	symaddr(&addr_QInputDialog_setComboBoxItems, "QInputDialog_setComboBoxItems")
	symaddr(&addr_QInputDialog_comboBoxItems, "QInputDialog_comboBoxItems")
	symaddr(&addr_QInputDialog_setIntValue, "QInputDialog_setIntValue")
	symaddr(&addr_QInputDialog_intValue, "QInputDialog_intValue")
	symaddr(&addr_QInputDialog_setIntMinimum, "QInputDialog_setIntMinimum")
	symaddr(&addr_QInputDialog_intMinimum, "QInputDialog_intMinimum")
	symaddr(&addr_QInputDialog_setIntMaximum, "QInputDialog_setIntMaximum")
	symaddr(&addr_QInputDialog_intMaximum, "QInputDialog_intMaximum")
	symaddr(&addr_QInputDialog_setIntRange, "QInputDialog_setIntRange")
	symaddr(&addr_QInputDialog_setIntStep, "QInputDialog_setIntStep")
	symaddr(&addr_QInputDialog_intStep, "QInputDialog_intStep")
	reg(&pg_QInputDialog_setDoubleValue, "QInputDialog_setDoubleValue")
	reg(&pg_QInputDialog_doubleValue, "QInputDialog_doubleValue")
	reg(&pg_QInputDialog_setDoubleMinimum, "QInputDialog_setDoubleMinimum")
	reg(&pg_QInputDialog_doubleMinimum, "QInputDialog_doubleMinimum")
	reg(&pg_QInputDialog_setDoubleMaximum, "QInputDialog_setDoubleMaximum")
	reg(&pg_QInputDialog_doubleMaximum, "QInputDialog_doubleMaximum")
	reg(&pg_QInputDialog_setDoubleRange, "QInputDialog_setDoubleRange")
	symaddr(&addr_QInputDialog_setDoubleDecimals, "QInputDialog_setDoubleDecimals")
	symaddr(&addr_QInputDialog_doubleDecimals, "QInputDialog_doubleDecimals")
	symaddr(&addr_QInputDialog_setOkButtonText, "QInputDialog_setOkButtonText")
	symaddr(&addr_QInputDialog_okButtonText, "QInputDialog_okButtonText")
	symaddr(&addr_QInputDialog_setCancelButtonText, "QInputDialog_setCancelButtonText")
	symaddr(&addr_QInputDialog_cancelButtonText, "QInputDialog_cancelButtonText")
	symaddr(&addr_QInputDialog_open, "QInputDialog_open")
	symaddr(&addr_QInputDialog_getText, "QInputDialog_getText")
	symaddr(&addr_QInputDialog_getMultiLineText, "QInputDialog_getMultiLineText")
	symaddr(&addr_QInputDialog_getItem, "QInputDialog_getItem")
	symaddr(&addr_QInputDialog_getInt, "QInputDialog_getInt")
	reg(&pg_QInputDialog_getDouble, "QInputDialog_getDouble")
	reg(&pg_QInputDialog_setDoubleStep, "QInputDialog_setDoubleStep")
	reg(&pg_QInputDialog_doubleStep, "QInputDialog_doubleStep")
	symaddr(&addr_QInputDialog_textValueChanged, "QInputDialog_textValueChanged")
	symaddr(&addr_QInputDialog_textValueSelected, "QInputDialog_textValueSelected")
	symaddr(&addr_QInputDialog_intValueChanged, "QInputDialog_intValueChanged")
	symaddr(&addr_QInputDialog_intValueSelected, "QInputDialog_intValueSelected")
	reg(&pg_QInputDialog_doubleValueChanged, "QInputDialog_doubleValueChanged")
	reg(&pg_QInputDialog_doubleValueSelected, "QInputDialog_doubleValueSelected")
	symaddr(&addr_QInputDialog_adjustPosition, "QInputDialog_adjustPosition")
	symaddr(&addr_QInputDialog_setTabOrder, "QInputDialog_setTabOrder")
	symaddr(&addr_QInputDialog_mouseGrabber, "QInputDialog_mouseGrabber")
	symaddr(&addr_QInputDialog_keyboardGrabber, "QInputDialog_keyboardGrabber")
	symaddr(&addr_QInputDialog_find, "QInputDialog_find")
	symaddr(&addr_QInputDialog_createWindowContainer, "QInputDialog_createWindowContainer")
	symaddr(&addr_QInputDialog_updateMicroFocus, "QInputDialog_updateMicroFocus")
	symaddr(&addr_QInputDialog_create, "QInputDialog_create")
	symaddr(&addr_QInputDialog_destroy, "QInputDialog_destroy")
	symaddr(&addr_QInputDialog_focusNextChild, "QInputDialog_focusNextChild")
	symaddr(&addr_QInputDialog_focusPreviousChild, "QInputDialog_focusPreviousChild")
	symaddr(&addr_QInputDialog_connect, "QInputDialog_connect")
	symaddr(&addr_QInputDialog_connect2, "QInputDialog_connect2")
	symaddr(&addr_QInputDialog_sender, "QInputDialog_sender")
	symaddr(&addr_QInputDialog_senderSignalIndex, "QInputDialog_senderSignalIndex")
	symaddr(&addr_QInputDialog_receivers, "QInputDialog_receivers")
	symaddr(&addr_QInputDialog_isSignalConnected, "QInputDialog_isSignalConnected")
	reg(&pg_QInputDialog_devicePixelRatioFScale, "QInputDialog_devicePixelRatioFScale")
	reg(&pg_QInputDialog_encodeMetricF, "QInputDialog_encodeMetricF")
	reg(&pg_QInputDialog_getDecodedMetricF, "QInputDialog_getDecodedMetricF")
	symaddr(&addr_QInputDialog_tr2, "QInputDialog_tr2")
	symaddr(&addr_QInputDialog_tr3, "QInputDialog_tr3")
	symaddr(&addr_QInputDialog_setOption2, "QInputDialog_setOption2")
	symaddr(&addr_QInputDialog_getText2, "QInputDialog_getText2")
	symaddr(&addr_QInputDialog_getText3, "QInputDialog_getText3")
	symaddr(&addr_QInputDialog_getText4, "QInputDialog_getText4")
	symaddr(&addr_QInputDialog_getText5, "QInputDialog_getText5")
	symaddr(&addr_QInputDialog_getText6, "QInputDialog_getText6")
	symaddr(&addr_QInputDialog_getMultiLineText2, "QInputDialog_getMultiLineText2")
	symaddr(&addr_QInputDialog_getMultiLineText3, "QInputDialog_getMultiLineText3")
	symaddr(&addr_QInputDialog_getMultiLineText4, "QInputDialog_getMultiLineText4")
	symaddr(&addr_QInputDialog_getMultiLineText5, "QInputDialog_getMultiLineText5")
	symaddr(&addr_QInputDialog_getItem2, "QInputDialog_getItem2")
	symaddr(&addr_QInputDialog_getItem3, "QInputDialog_getItem3")
	symaddr(&addr_QInputDialog_getItem4, "QInputDialog_getItem4")
	symaddr(&addr_QInputDialog_getItem5, "QInputDialog_getItem5")
	symaddr(&addr_QInputDialog_getItem6, "QInputDialog_getItem6")
	symaddr(&addr_QInputDialog_getInt2, "QInputDialog_getInt2")
	symaddr(&addr_QInputDialog_getInt3, "QInputDialog_getInt3")
	symaddr(&addr_QInputDialog_getInt4, "QInputDialog_getInt4")
	symaddr(&addr_QInputDialog_getInt5, "QInputDialog_getInt5")
	symaddr(&addr_QInputDialog_getInt6, "QInputDialog_getInt6")
	symaddr(&addr_QInputDialog_getInt7, "QInputDialog_getInt7")
	reg(&pg_QInputDialog_getDouble2, "QInputDialog_getDouble2")
	reg(&pg_QInputDialog_getDouble3, "QInputDialog_getDouble3")
	reg(&pg_QInputDialog_getDouble4, "QInputDialog_getDouble4")
	reg(&pg_QInputDialog_getDouble5, "QInputDialog_getDouble5")
	reg(&pg_QInputDialog_getDouble6, "QInputDialog_getDouble6")
	reg(&pg_QInputDialog_getDouble7, "QInputDialog_getDouble7")
	reg(&pg_QInputDialog_getDouble8, "QInputDialog_getDouble8")
	symaddr(&addr_QInputDialog_createWindowContainer2, "QInputDialog_createWindowContainer2")
	symaddr(&addr_QInputDialog_createWindowContainer3, "QInputDialog_createWindowContainer3")
	symaddr(&addr_QInputDialog_updateMicroFocus2, "QInputDialog_updateMicroFocus2")
	symaddr(&addr_QInputDialog_create2, "QInputDialog_create2")
	symaddr(&addr_QInputDialog_create3, "QInputDialog_create3")
	symaddr(&addr_QInputDialog_create4, "QInputDialog_create4")
	symaddr(&addr_QInputDialog_destroy2, "QInputDialog_destroy2")
	symaddr(&addr_QInputDialog_destroy3, "QInputDialog_destroy3")
	symaddr(&addr_QInputDialog_connect3, "QInputDialog_connect3")
	symaddr(&addr_QInputDialog_connect4, "QInputDialog_connect4")
	symaddr(&addr_QInputDialog_metaObject_default, "QInputDialog_metaObject_default")
	symaddr(&addr_QInputDialog_minimumSizeHint_default, "QInputDialog_minimumSizeHint_default")
	symaddr(&addr_QInputDialog_sizeHint_default, "QInputDialog_sizeHint_default")
	symaddr(&addr_QInputDialog_setVisible_default, "QInputDialog_setVisible_default")
	symaddr(&addr_QInputDialog_done_default, "QInputDialog_done_default")
	symaddr(&addr_QInputDialog_exec_default, "QInputDialog_exec_default")
	symaddr(&addr_QInputDialog_accept_default, "QInputDialog_accept_default")
	symaddr(&addr_QInputDialog_reject_default, "QInputDialog_reject_default")
	symaddr(&addr_QInputDialog_keyPressEvent_default, "QInputDialog_keyPressEvent_default")
	symaddr(&addr_QInputDialog_closeEvent_default, "QInputDialog_closeEvent_default")
	symaddr(&addr_QInputDialog_showEvent_default, "QInputDialog_showEvent_default")
	symaddr(&addr_QInputDialog_resizeEvent_default, "QInputDialog_resizeEvent_default")
	symaddr(&addr_QInputDialog_contextMenuEvent_default, "QInputDialog_contextMenuEvent_default")
	symaddr(&addr_QInputDialog_eventFilter_default, "QInputDialog_eventFilter_default")
	symaddr(&addr_QInputDialog_devType_default, "QInputDialog_devType_default")
	symaddr(&addr_QInputDialog_heightForWidth_default, "QInputDialog_heightForWidth_default")
	symaddr(&addr_QInputDialog_hasHeightForWidth_default, "QInputDialog_hasHeightForWidth_default")
	symaddr(&addr_QInputDialog_paintEngine_default, "QInputDialog_paintEngine_default")
	symaddr(&addr_QInputDialog_event_default, "QInputDialog_event_default")
	symaddr(&addr_QInputDialog_mousePressEvent_default, "QInputDialog_mousePressEvent_default")
	symaddr(&addr_QInputDialog_mouseReleaseEvent_default, "QInputDialog_mouseReleaseEvent_default")
	symaddr(&addr_QInputDialog_mouseDoubleClickEvent_default, "QInputDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QInputDialog_mouseMoveEvent_default, "QInputDialog_mouseMoveEvent_default")
	symaddr(&addr_QInputDialog_wheelEvent_default, "QInputDialog_wheelEvent_default")
	symaddr(&addr_QInputDialog_keyReleaseEvent_default, "QInputDialog_keyReleaseEvent_default")
	symaddr(&addr_QInputDialog_focusInEvent_default, "QInputDialog_focusInEvent_default")
	symaddr(&addr_QInputDialog_focusOutEvent_default, "QInputDialog_focusOutEvent_default")
	symaddr(&addr_QInputDialog_enterEvent_default, "QInputDialog_enterEvent_default")
	symaddr(&addr_QInputDialog_leaveEvent_default, "QInputDialog_leaveEvent_default")
	symaddr(&addr_QInputDialog_paintEvent_default, "QInputDialog_paintEvent_default")
	symaddr(&addr_QInputDialog_moveEvent_default, "QInputDialog_moveEvent_default")
	symaddr(&addr_QInputDialog_tabletEvent_default, "QInputDialog_tabletEvent_default")
	symaddr(&addr_QInputDialog_actionEvent_default, "QInputDialog_actionEvent_default")
	symaddr(&addr_QInputDialog_dragEnterEvent_default, "QInputDialog_dragEnterEvent_default")
	symaddr(&addr_QInputDialog_dragMoveEvent_default, "QInputDialog_dragMoveEvent_default")
	symaddr(&addr_QInputDialog_dragLeaveEvent_default, "QInputDialog_dragLeaveEvent_default")
	symaddr(&addr_QInputDialog_dropEvent_default, "QInputDialog_dropEvent_default")
	symaddr(&addr_QInputDialog_hideEvent_default, "QInputDialog_hideEvent_default")
	symaddr(&addr_QInputDialog_changeEvent_default, "QInputDialog_changeEvent_default")
	symaddr(&addr_QInputDialog_metric_default, "QInputDialog_metric_default")
	symaddr(&addr_QInputDialog_initPainter_default, "QInputDialog_initPainter_default")
	symaddr(&addr_QInputDialog_redirected_default, "QInputDialog_redirected_default")
	symaddr(&addr_QInputDialog_sharedPainter_default, "QInputDialog_sharedPainter_default")
	symaddr(&addr_QInputDialog_inputMethodEvent_default, "QInputDialog_inputMethodEvent_default")
	symaddr(&addr_QInputDialog_inputMethodQuery_default, "QInputDialog_inputMethodQuery_default")
	symaddr(&addr_QInputDialog_focusNextPrevChild_default, "QInputDialog_focusNextPrevChild_default")
	symaddr(&addr_QInputDialog_timerEvent_default, "QInputDialog_timerEvent_default")
	symaddr(&addr_QInputDialog_childEvent_default, "QInputDialog_childEvent_default")
	symaddr(&addr_QInputDialog_customEvent_default, "QInputDialog_customEvent_default")
	symaddr(&addr_QInputDialog_connectNotify_default, "QInputDialog_connectNotify_default")
	symaddr(&addr_QInputDialog_disconnectNotify_default, "QInputDialog_disconnectNotify_default")
	reg(&pg_QInputDialog_connect_textValueChanged, "QInputDialog_connect_textValueChanged")
	reg(&pg_QInputDialog_connect_textValueSelected, "QInputDialog_connect_textValueSelected")
	reg(&pg_QInputDialog_connect_intValueChanged, "QInputDialog_connect_intValueChanged")
	reg(&pg_QInputDialog_connect_intValueSelected, "QInputDialog_connect_intValueSelected")
	reg(&pg_QInputDialog_connect_doubleValueChanged, "QInputDialog_connect_doubleValueChanged")
	reg(&pg_QInputDialog_connect_doubleValueSelected, "QInputDialog_connect_doubleValueSelected")
	reg(&pg_QInputDialog_connect_finished, "QInputDialog_connect_finished")
	reg(&pg_QInputDialog_connect_accepted, "QInputDialog_connect_accepted")
	reg(&pg_QInputDialog_connect_rejected, "QInputDialog_connect_rejected")
	reg(&pg_QInputDialog_connect_windowTitleChanged, "QInputDialog_connect_windowTitleChanged")
	reg(&pg_QInputDialog_connect_windowIconChanged, "QInputDialog_connect_windowIconChanged")
	reg(&pg_QInputDialog_connect_windowIconTextChanged, "QInputDialog_connect_windowIconTextChanged")
	reg(&pg_QInputDialog_connect_customContextMenuRequested, "QInputDialog_connect_customContextMenuRequested")
	reg(&pg_QInputDialog_connect_destroyed, "QInputDialog_connect_destroyed")
	reg(&pg_QInputDialog_connect_destroyed2, "QInputDialog_connect_destroyed2")
	reg(&pg_QInputDialog_connect_objectNameChanged, "QInputDialog_connect_objectNameChanged")
	reg(&pg_QInputDialog__disconnect, "QInputDialog__disconnect")
	reg(&pg_QInputDialog__set_override, "QInputDialog__set_override")
	reg(&pg_QInputDialog_delete, "QInputDialog_delete")
}

type QInputDialog struct {
	h uintptr
	*QDialog
}

func bindQInputDialog(h uintptr) *QInputDialog {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QInputDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QInputDialog{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQInputDialog(h uintptr) *QInputDialog { return bindQInputDialog(h) }

func (this *QInputDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputDialog(h unsafe.Pointer) *QInputDialog { return bindQInputDialog(uintptr(h)) }
func UnsafeNewQInputDialogOwned(h unsafe.Pointer) *QInputDialog { return newownedQInputDialog(uintptr(h)) }

func (this *QInputDialog) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQInputDialog(parent *QWidget) *QInputDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQInputDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDialog_virtIdx, QInputDialog_enable)
	return __o
}

func NewQInputDialog2() *QInputDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputDialog_new2, uintptr(__goid))
	__o := newownedQInputDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDialog_virtIdx, QInputDialog_enable)
	return __o
}

func NewQInputDialog3(parent *QWidget, flags WindowType) *QInputDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputDialog_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQInputDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDialog_virtIdx, QInputDialog_enable)
	return __o
}

func QInputDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QInputDialog) SetInputMode(mode QInputDialog__InputMode) {
	purego.SyscallN(addr_QInputDialog_setInputMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) InputMode() QInputDialog__InputMode {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_inputMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputDialog__InputMode(__r)
}

func (this *QInputDialog) SetLabelText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_setLabelText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) LabelText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_labelText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDialog) SetOption(option QInputDialog__InputDialogOption) {
	purego.SyscallN(addr_QInputDialog_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) TestOption(option QInputDialog__InputDialogOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDialog) SetOptions(options QInputDialog__InputDialogOption) {
	purego.SyscallN(addr_QInputDialog_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Options() QInputDialog__InputDialogOption {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputDialog__InputDialogOption(__r)
}

func (this *QInputDialog) SetTextValue(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_setTextValue, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) TextValue() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_textValue, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDialog) SetTextEchoMode(mode QLineEdit__EchoMode) {
	purego.SyscallN(addr_QInputDialog_setTextEchoMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) TextEchoMode() QLineEdit__EchoMode {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_textEchoMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLineEdit__EchoMode(__r)
}

func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	purego.SyscallN(addr_QInputDialog_setComboBoxEditable, this.h, uintptr(b2i(editable)))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IsComboBoxEditable() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_isComboBoxEditable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDialog) SetComboBoxItems(items []string) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QInputDialog_setComboBoxItems, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
}

func (this *QInputDialog) ComboBoxItems() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_comboBoxItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QInputDialog) SetIntValue(value int) {
	purego.SyscallN(addr_QInputDialog_setIntValue, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IntValue() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_intValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) SetIntMinimum(min int) {
	purego.SyscallN(addr_QInputDialog_setIntMinimum, this.h, uintptr(min))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IntMinimum() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_intMinimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) SetIntMaximum(max int) {
	purego.SyscallN(addr_QInputDialog_setIntMaximum, this.h, uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IntMaximum() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_intMaximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) SetIntRange(min int, max int) {
	purego.SyscallN(addr_QInputDialog_setIntRange, this.h, uintptr(min), uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) SetIntStep(step int) {
	purego.SyscallN(addr_QInputDialog_setIntStep, this.h, uintptr(step))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IntStep() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_intStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) SetDoubleValue(value float64) {
	pg_QInputDialog_setDoubleValue(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleValue() float64 {
	__r := pg_QInputDialog_doubleValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QInputDialog) SetDoubleMinimum(min float64) {
	pg_QInputDialog_setDoubleMinimum(this.h, min)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleMinimum() float64 {
	__r := pg_QInputDialog_doubleMinimum(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QInputDialog) SetDoubleMaximum(max float64) {
	pg_QInputDialog_setDoubleMaximum(this.h, max)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleMaximum() float64 {
	__r := pg_QInputDialog_doubleMaximum(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	pg_QInputDialog_setDoubleRange(this.h, min, max)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	purego.SyscallN(addr_QInputDialog_setDoubleDecimals, this.h, uintptr(decimals))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleDecimals() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_doubleDecimals, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) SetOkButtonText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_setOkButtonText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) OkButtonText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_okButtonText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDialog) SetCancelButtonText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_setCancelButtonText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) CancelButtonText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_cancelButtonText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDialog) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QInputDialog_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func QInputDialog_GetText(parent *QWidget, title string, label string) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return pgStr(__r, __n)
}

func QInputDialog_GetMultiLineText(parent *QWidget, title string, label string) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getMultiLineText, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return pgStr(__r, __n)
}

func QInputDialog_GetItem(parent *QWidget, title string, label string, items []string) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n)
}

func QInputDialog_GetInt(parent *QWidget, title string, label string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r)
}

func QInputDialog_GetDouble(parent *QWidget, title string, label string) float64 {
	__r := pg_QInputDialog_getDouble(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)))
	runtime.KeepAlive(parent)
	return __r
}

func (this *QInputDialog) SetDoubleStep(step float64) {
	pg_QInputDialog_setDoubleStep(this.h, step)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleStep() float64 {
	__r := pg_QInputDialog_doubleStep(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QInputDialog) TextValueChanged(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_textValueChanged, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) TextValueSelected(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QInputDialog_textValueSelected, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QInputDialog) IntValueChanged(value int) {
	purego.SyscallN(addr_QInputDialog_intValueChanged, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) IntValueSelected(value int) {
	purego.SyscallN(addr_QInputDialog_intValueSelected, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleValueChanged(value float64) {
	pg_QInputDialog_doubleValueChanged(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoubleValueSelected(value float64) {
	pg_QInputDialog_doubleValueSelected(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QInputDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QInputDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QInputDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QInputDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QInputDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QInputDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QInputDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QInputDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QInputDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Create() {
	purego.SyscallN(addr_QInputDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Destroy() {
	purego.SyscallN(addr_QInputDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QInputDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QInputDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QInputDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QInputDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QInputDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QInputDialog_devicePixelRatioFScale()
	return __r
}

func QInputDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QInputDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QInputDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QInputDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QInputDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QInputDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QInputDialog) SetOption2(option QInputDialog__InputDialogOption, on bool) {
	purego.SyscallN(addr_QInputDialog_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QInputDialog_GetText2(parent *QWidget, title string, label string, echo QLineEdit__EchoMode) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(echo), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return pgStr(__r, __n)
}

func QInputDialog_GetText3(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(echo), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n)
}

func QInputDialog_GetText4(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(echo), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetText5(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string, flags WindowType) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(echo), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetText6(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string, flags WindowType, inputMethodHints InputMethodHint) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getText6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(echo), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(inputMethodHints), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetMultiLineText2(parent *QWidget, title string, label string, text string) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getMultiLineText2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n)
}

func QInputDialog_GetMultiLineText3(parent *QWidget, title string, label string, text string) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getMultiLineText3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetMultiLineText4(parent *QWidget, title string, label string, text string, flags WindowType) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getMultiLineText4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetMultiLineText5(parent *QWidget, title string, label string, text string, flags WindowType, inputMethodHints InputMethodHint) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getMultiLineText5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(inputMethodHints), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	runtime.KeepAlive(text)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetItem2(parent *QWidget, title string, label string, items []string, current int) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(current), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n)
}

func QInputDialog_GetItem3(parent *QWidget, title string, label string, items []string, current int, editable bool) string {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(current), uintptr(b2i(editable)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n)
}

func QInputDialog_GetItem4(parent *QWidget, title string, label string, items []string, current int, editable bool) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(current), uintptr(b2i(editable)), uintptr(unsafe.Pointer(&ok_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetItem5(parent *QWidget, title string, label string, items []string, current int, editable bool, flags WindowType) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(current), uintptr(b2i(editable)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetItem6(parent *QWidget, title string, label string, items []string, current int, editable bool, flags WindowType, inputMethodHints InputMethodHint) (string, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = pgCString(__e)
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	var ok_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDialog_getItem6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), items_p, uintptr(len(items)), uintptr(current), uintptr(b2i(editable)), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags), uintptr(inputMethodHints), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	for __i := range items_arr {
		pgFree(items_arr[__i])
	}
	runtime.KeepAlive(items_arr)
	return pgStr(__r, __n), ok_out != 0
}

func QInputDialog_GetInt2(parent *QWidget, title string, label string, value int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r)
}

func QInputDialog_GetInt3(parent *QWidget, title string, label string, value int, minValue int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value), uintptr(minValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r)
}

func QInputDialog_GetInt4(parent *QWidget, title string, label string, value int, minValue int, maxValue int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value), uintptr(minValue), uintptr(maxValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r)
}

func QInputDialog_GetInt5(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value), uintptr(minValue), uintptr(maxValue), uintptr(step))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r)
}

func QInputDialog_GetInt6(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int) (int, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value), uintptr(minValue), uintptr(maxValue), uintptr(step), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r), ok_out != 0
}

func QInputDialog_GetInt7(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int, flags WindowType) (int, bool) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_getInt7, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), label_p, uintptr(len(label)), uintptr(value), uintptr(minValue), uintptr(maxValue), uintptr(step), uintptr(unsafe.Pointer(&ok_out)), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(label)
	return int(__r), ok_out != 0
}

func QInputDialog_GetDouble2(parent *QWidget, title string, label string, value float64) float64 {
	__r := pg_QInputDialog_getDouble2(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value)
	runtime.KeepAlive(parent)
	return __r
}

func QInputDialog_GetDouble3(parent *QWidget, title string, label string, value float64, minValue float64) float64 {
	__r := pg_QInputDialog_getDouble3(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue)
	runtime.KeepAlive(parent)
	return __r
}

func QInputDialog_GetDouble4(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64) float64 {
	__r := pg_QInputDialog_getDouble4(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue, maxValue)
	runtime.KeepAlive(parent)
	return __r
}

func QInputDialog_GetDouble5(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int) float64 {
	__r := pg_QInputDialog_getDouble5(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue, maxValue, int32(decimals))
	runtime.KeepAlive(parent)
	return __r
}

func QInputDialog_GetDouble6(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int) (float64, bool) {
	var ok_out int32
	__r := pg_QInputDialog_getDouble6(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue, maxValue, int32(decimals), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(parent)
	return __r, ok_out != 0
}

func QInputDialog_GetDouble7(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int, flags WindowType) (float64, bool) {
	var ok_out int32
	__r := pg_QInputDialog_getDouble7(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue, maxValue, int32(decimals), uintptr(unsafe.Pointer(&ok_out)), int32(flags))
	runtime.KeepAlive(parent)
	return __r, ok_out != 0
}

func QInputDialog_GetDouble8(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int, flags WindowType, step float64) (float64, bool) {
	var ok_out int32
	__r := pg_QInputDialog_getDouble8(uintptr(parent.UnsafePointer()), title, int64(len(title)), label, int64(len(label)), value, minValue, maxValue, int32(decimals), uintptr(unsafe.Pointer(&ok_out)), int32(flags), step)
	runtime.KeepAlive(parent)
	return __r, ok_out != 0
}

func QInputDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QInputDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QInputDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QInputDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QInputDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QInputDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QInputDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QInputDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QInputDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QInputDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QInputDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QInputDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QInputDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QInputDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QInputDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) DoneDefault(result int) {
	purego.SyscallN(addr_QInputDialog_done_default, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QInputDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) AcceptDefault() {
	purego.SyscallN(addr_QInputDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) RejectDefault() {
	purego.SyscallN(addr_QInputDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QInputDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QInputDialog_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QInputDialog_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QInputDialog_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QInputDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QInputDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QInputDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QInputDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QInputDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QInputDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QInputDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QInputDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QInputDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QInputDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QInputDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QInputDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QInputDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QInputDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QInputDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QInputDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QInputDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QInputDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QInputDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QInputDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QInputDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QInputDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QInputDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QInputDialog_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QInputDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QInputDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QInputDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QInputDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QInputDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QInputDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QInputDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QInputDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QInputDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QInputDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QInputDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QInputDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QInputDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QInputDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QInputDialog) OnTextValueChanged(fn func(string)) int64 {
	id := registerCallback(func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) })
	pg_QInputDialog_connect_textValueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnTextValueSelected(fn func(string)) int64 {
	id := registerCallback(func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) })
	pg_QInputDialog_connect_textValueSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnIntValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QInputDialog_connect_intValueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnIntValueSelected(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QInputDialog_connect_intValueSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnDoubleValueChanged(fn func(float64)) int64 {
	id := registerCallback(func(value uintptr) { fn(math.Float64frombits(uint64(value))) })
	pg_QInputDialog_connect_doubleValueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnDoubleValueSelected(fn func(float64)) int64 {
	id := registerCallback(func(value uintptr) { fn(math.Float64frombits(uint64(value))) })
	pg_QInputDialog_connect_doubleValueSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QInputDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QInputDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QInputDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QInputDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QInputDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QInputDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QInputDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDialog) Disconnect(handle int64) {
	pg_QInputDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QInputDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7049264599837532491:
		return 1
	case 729690261457684101:
		return 2
	case 2110182719375245575:
		return 3
	case -2570997490807640431:
		return 4
	case -1745322561779819684:
		return 5
	case 9106131928185084681:
		return 6
	case 2875220963396170150:
		return 7
	case 1023791977679793185:
		return 8
	case -1989937050019464783:
		return 9
	case -6825369174751916:
		return 10
	case -4302917285655014661:
		return 11
	case -6698043222227006015:
		return 12
	case 5217623514630858297:
		return 13
	case -8659679893868524044:
		return 14
	case 4311534060495737339:
		return 15
	case -8266762397995423915:
		return 16
	case 6398007404756346381:
		return 17
	case 1237236456312824991:
		return 18
	case -332090898078543319:
		return 19
	case -4415920049297795213:
		return 20
	case -4029073905697023887:
		return 21
	case -4255494004099806581:
		return 22
	case 1919988340934511316:
		return 23
	case 772598475950050299:
		return 24
	case 975043767033552552:
		return 25
	case -385596010576048993:
		return 26
	case 6316270346668237895:
		return 27
	case -551070009623711640:
		return 28
	case 5290442163167458465:
		return 29
	case -9161949150489525396:
		return 30
	case -212068866860701141:
		return 31
	case 1697444140088555581:
		return 32
	case 206122715919917589:
		return 33
	case -1781545137352745978:
		return 34
	case -3139904158373725006:
		return 35
	case 9107058682754591716:
		return 36
	case -199718707876719129:
		return 37
	case -2039166586246443706:
		return 38
	case 4959376063826165943:
		return 39
	case 427629649267311777:
		return 40
	case 6477500296980050210:
		return 41
	case 4127240633310655048:
		return 42
	case 5792771328741774991:
		return 43
	case 5438819661942843222:
		return 44
	case 7345853097340709638:
		return 45
	case 5898328517505784191:
		return 46
	case -5409927581812374498:
		return 47
	case -1185315367526857833:
		return 48
	case 3769868229123868474:
		return 49
	case 7379029982170503008:
		return 50
	case -568369772982215962:
		return 51
	}
	return -1
}

func QInputDialog_enable(h uintptr, idx int32) { pg_QInputDialog__set_override(h, idx) }

func (this *QInputDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QInputDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QInputDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QInputDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QInputDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(result int32) { fn(int(result)) }) // slot "done"
}

func (this *QInputDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QInputDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QInputDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QInputDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QInputDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QInputDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QInputDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QInputDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QInputDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QInputDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QInputDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QInputDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QInputDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QInputDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QInputDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QInputDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QInputDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QInputDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QInputDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QInputDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QInputDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QInputDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QInputDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QInputDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QInputDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QInputDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QInputDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QInputDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QInputDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QInputDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QInputDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QInputDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QInputDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QInputDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QInputDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QInputDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QInputDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QInputDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QInputDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QInputDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QInputDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QInputDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QInputDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QInputDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QInputDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QInputDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QInputDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QInputDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
