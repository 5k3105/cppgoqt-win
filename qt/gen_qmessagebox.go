package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMessageBox_chain uintptr
	addr_QMessageBox_new uintptr
	addr_QMessageBox_new2 uintptr
	addr_QMessageBox_new3 uintptr
	addr_QMessageBox_new4 uintptr
	addr_QMessageBox_new5 uintptr
	addr_QMessageBox_new6 uintptr
	addr_QMessageBox_new7 uintptr
	addr_QMessageBox_new8 uintptr
	addr_QMessageBox_new9 uintptr
	addr_QMessageBox_tr uintptr
	addr_QMessageBox_addButton uintptr
	addr_QMessageBox_addButton2 uintptr
	addr_QMessageBox_addButton3 uintptr
	addr_QMessageBox_removeButton uintptr
	addr_QMessageBox_open uintptr
	addr_QMessageBox_buttons uintptr
	addr_QMessageBox_buttonRole uintptr
	addr_QMessageBox_setStandardButtons uintptr
	addr_QMessageBox_standardButtons uintptr
	addr_QMessageBox_standardButton uintptr
	addr_QMessageBox_button uintptr
	addr_QMessageBox_defaultButton uintptr
	addr_QMessageBox_setDefaultButton uintptr
	addr_QMessageBox_setDefaultButton2 uintptr
	addr_QMessageBox_escapeButton uintptr
	addr_QMessageBox_setEscapeButton uintptr
	addr_QMessageBox_setEscapeButton2 uintptr
	addr_QMessageBox_clickedButton uintptr
	addr_QMessageBox_text uintptr
	addr_QMessageBox_setText uintptr
	addr_QMessageBox_icon uintptr
	addr_QMessageBox_setIcon uintptr
	addr_QMessageBox_iconPixmap uintptr
	addr_QMessageBox_setIconPixmap uintptr
	addr_QMessageBox_textFormat uintptr
	addr_QMessageBox_setTextFormat uintptr
	addr_QMessageBox_setTextInteractionFlags uintptr
	addr_QMessageBox_textInteractionFlags uintptr
	addr_QMessageBox_setCheckBox uintptr
	addr_QMessageBox_checkBox uintptr
	addr_QMessageBox_setOption uintptr
	addr_QMessageBox_testOption uintptr
	addr_QMessageBox_setOptions uintptr
	addr_QMessageBox_options uintptr
	addr_QMessageBox_information uintptr
	addr_QMessageBox_information2 uintptr
	addr_QMessageBox_question uintptr
	addr_QMessageBox_question2 uintptr
	addr_QMessageBox_warning uintptr
	addr_QMessageBox_warning2 uintptr
	addr_QMessageBox_critical uintptr
	addr_QMessageBox_critical2 uintptr
	addr_QMessageBox_about uintptr
	addr_QMessageBox_aboutQt uintptr
	addr_QMessageBox_information3 uintptr
	addr_QMessageBox_information4 uintptr
	addr_QMessageBox_question3 uintptr
	addr_QMessageBox_question4 uintptr
	addr_QMessageBox_warning3 uintptr
	addr_QMessageBox_warning4 uintptr
	addr_QMessageBox_critical3 uintptr
	addr_QMessageBox_critical4 uintptr
	addr_QMessageBox_buttonText uintptr
	addr_QMessageBox_setButtonText uintptr
	addr_QMessageBox_informativeText uintptr
	addr_QMessageBox_setInformativeText uintptr
	addr_QMessageBox_detailedText uintptr
	addr_QMessageBox_setDetailedText uintptr
	addr_QMessageBox_setWindowTitle uintptr
	addr_QMessageBox_setWindowModality uintptr
	addr_QMessageBox_standardIcon uintptr
	addr_QMessageBox_buttonClicked uintptr
	addr_QMessageBox_adjustPosition uintptr
	addr_QMessageBox_setTabOrder uintptr
	addr_QMessageBox_mouseGrabber uintptr
	addr_QMessageBox_keyboardGrabber uintptr
	addr_QMessageBox_find uintptr
	addr_QMessageBox_createWindowContainer uintptr
	addr_QMessageBox_updateMicroFocus uintptr
	addr_QMessageBox_create uintptr
	addr_QMessageBox_destroy uintptr
	addr_QMessageBox_focusNextChild uintptr
	addr_QMessageBox_focusPreviousChild uintptr
	addr_QMessageBox_connect uintptr
	addr_QMessageBox_connect2 uintptr
	addr_QMessageBox_sender uintptr
	addr_QMessageBox_senderSignalIndex uintptr
	addr_QMessageBox_receivers uintptr
	addr_QMessageBox_isSignalConnected uintptr
	pg_QMessageBox_devicePixelRatioFScale func() float64
	pg_QMessageBox_encodeMetricF func(int32, float64) int32
	pg_QMessageBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMessageBox_tr2 uintptr
	addr_QMessageBox_tr3 uintptr
	addr_QMessageBox_setOption2 uintptr
	addr_QMessageBox_information5 uintptr
	addr_QMessageBox_information6 uintptr
	addr_QMessageBox_information7 uintptr
	addr_QMessageBox_question5 uintptr
	addr_QMessageBox_question6 uintptr
	addr_QMessageBox_warning5 uintptr
	addr_QMessageBox_warning6 uintptr
	addr_QMessageBox_critical5 uintptr
	addr_QMessageBox_critical6 uintptr
	addr_QMessageBox_aboutQt2 uintptr
	addr_QMessageBox_information8 uintptr
	addr_QMessageBox_information9 uintptr
	addr_QMessageBox_information10 uintptr
	addr_QMessageBox_information11 uintptr
	addr_QMessageBox_information12 uintptr
	addr_QMessageBox_information13 uintptr
	addr_QMessageBox_question7 uintptr
	addr_QMessageBox_question8 uintptr
	addr_QMessageBox_question9 uintptr
	addr_QMessageBox_question10 uintptr
	addr_QMessageBox_question11 uintptr
	addr_QMessageBox_question12 uintptr
	addr_QMessageBox_warning7 uintptr
	addr_QMessageBox_warning8 uintptr
	addr_QMessageBox_warning9 uintptr
	addr_QMessageBox_warning10 uintptr
	addr_QMessageBox_warning11 uintptr
	addr_QMessageBox_critical7 uintptr
	addr_QMessageBox_critical8 uintptr
	addr_QMessageBox_critical9 uintptr
	addr_QMessageBox_critical10 uintptr
	addr_QMessageBox_critical11 uintptr
	addr_QMessageBox_createWindowContainer2 uintptr
	addr_QMessageBox_createWindowContainer3 uintptr
	addr_QMessageBox_updateMicroFocus2 uintptr
	addr_QMessageBox_create2 uintptr
	addr_QMessageBox_create3 uintptr
	addr_QMessageBox_create4 uintptr
	addr_QMessageBox_destroy2 uintptr
	addr_QMessageBox_destroy3 uintptr
	addr_QMessageBox_connect3 uintptr
	addr_QMessageBox_connect4 uintptr
	addr_QMessageBox_metaObject_default uintptr
	addr_QMessageBox_event_default uintptr
	addr_QMessageBox_resizeEvent_default uintptr
	addr_QMessageBox_showEvent_default uintptr
	addr_QMessageBox_closeEvent_default uintptr
	addr_QMessageBox_keyPressEvent_default uintptr
	addr_QMessageBox_changeEvent_default uintptr
	addr_QMessageBox_setVisible_default uintptr
	addr_QMessageBox_sizeHint_default uintptr
	addr_QMessageBox_minimumSizeHint_default uintptr
	addr_QMessageBox_exec_default uintptr
	addr_QMessageBox_done_default uintptr
	addr_QMessageBox_accept_default uintptr
	addr_QMessageBox_reject_default uintptr
	addr_QMessageBox_contextMenuEvent_default uintptr
	addr_QMessageBox_eventFilter_default uintptr
	addr_QMessageBox_devType_default uintptr
	addr_QMessageBox_heightForWidth_default uintptr
	addr_QMessageBox_hasHeightForWidth_default uintptr
	addr_QMessageBox_paintEngine_default uintptr
	addr_QMessageBox_mousePressEvent_default uintptr
	addr_QMessageBox_mouseReleaseEvent_default uintptr
	addr_QMessageBox_mouseDoubleClickEvent_default uintptr
	addr_QMessageBox_mouseMoveEvent_default uintptr
	addr_QMessageBox_wheelEvent_default uintptr
	addr_QMessageBox_keyReleaseEvent_default uintptr
	addr_QMessageBox_focusInEvent_default uintptr
	addr_QMessageBox_focusOutEvent_default uintptr
	addr_QMessageBox_enterEvent_default uintptr
	addr_QMessageBox_leaveEvent_default uintptr
	addr_QMessageBox_paintEvent_default uintptr
	addr_QMessageBox_moveEvent_default uintptr
	addr_QMessageBox_tabletEvent_default uintptr
	addr_QMessageBox_actionEvent_default uintptr
	addr_QMessageBox_dragEnterEvent_default uintptr
	addr_QMessageBox_dragMoveEvent_default uintptr
	addr_QMessageBox_dragLeaveEvent_default uintptr
	addr_QMessageBox_dropEvent_default uintptr
	addr_QMessageBox_hideEvent_default uintptr
	addr_QMessageBox_metric_default uintptr
	addr_QMessageBox_initPainter_default uintptr
	addr_QMessageBox_redirected_default uintptr
	addr_QMessageBox_sharedPainter_default uintptr
	addr_QMessageBox_inputMethodEvent_default uintptr
	addr_QMessageBox_inputMethodQuery_default uintptr
	addr_QMessageBox_focusNextPrevChild_default uintptr
	addr_QMessageBox_timerEvent_default uintptr
	addr_QMessageBox_childEvent_default uintptr
	addr_QMessageBox_customEvent_default uintptr
	addr_QMessageBox_connectNotify_default uintptr
	addr_QMessageBox_disconnectNotify_default uintptr
	pg_QMessageBox_connect_buttonClicked func(uintptr, int64)
	pg_QMessageBox_connect_finished func(uintptr, int64)
	pg_QMessageBox_connect_accepted func(uintptr, int64)
	pg_QMessageBox_connect_rejected func(uintptr, int64)
	pg_QMessageBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QMessageBox_connect_windowIconChanged func(uintptr, int64)
	pg_QMessageBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMessageBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMessageBox_connect_destroyed func(uintptr, int64)
	pg_QMessageBox_connect_destroyed2 func(uintptr, int64)
	pg_QMessageBox_connect_objectNameChanged func(uintptr, int64)
	pg_QMessageBox__disconnect func(int64)
	pg_QMessageBox__set_override func(uintptr, int32)
	pg_QMessageBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMessageBox_chain, "QMessageBox_chain")
	symaddr(&addr_QMessageBox_new, "QMessageBox_new")
	symaddr(&addr_QMessageBox_new2, "QMessageBox_new2")
	symaddr(&addr_QMessageBox_new3, "QMessageBox_new3")
	symaddr(&addr_QMessageBox_new4, "QMessageBox_new4")
	symaddr(&addr_QMessageBox_new5, "QMessageBox_new5")
	symaddr(&addr_QMessageBox_new6, "QMessageBox_new6")
	symaddr(&addr_QMessageBox_new7, "QMessageBox_new7")
	symaddr(&addr_QMessageBox_new8, "QMessageBox_new8")
	symaddr(&addr_QMessageBox_new9, "QMessageBox_new9")
	symaddr(&addr_QMessageBox_tr, "QMessageBox_tr")
	symaddr(&addr_QMessageBox_addButton, "QMessageBox_addButton")
	symaddr(&addr_QMessageBox_addButton2, "QMessageBox_addButton2")
	symaddr(&addr_QMessageBox_addButton3, "QMessageBox_addButton3")
	symaddr(&addr_QMessageBox_removeButton, "QMessageBox_removeButton")
	symaddr(&addr_QMessageBox_open, "QMessageBox_open")
	symaddr(&addr_QMessageBox_buttons, "QMessageBox_buttons")
	symaddr(&addr_QMessageBox_buttonRole, "QMessageBox_buttonRole")
	symaddr(&addr_QMessageBox_setStandardButtons, "QMessageBox_setStandardButtons")
	symaddr(&addr_QMessageBox_standardButtons, "QMessageBox_standardButtons")
	symaddr(&addr_QMessageBox_standardButton, "QMessageBox_standardButton")
	symaddr(&addr_QMessageBox_button, "QMessageBox_button")
	symaddr(&addr_QMessageBox_defaultButton, "QMessageBox_defaultButton")
	symaddr(&addr_QMessageBox_setDefaultButton, "QMessageBox_setDefaultButton")
	symaddr(&addr_QMessageBox_setDefaultButton2, "QMessageBox_setDefaultButton2")
	symaddr(&addr_QMessageBox_escapeButton, "QMessageBox_escapeButton")
	symaddr(&addr_QMessageBox_setEscapeButton, "QMessageBox_setEscapeButton")
	symaddr(&addr_QMessageBox_setEscapeButton2, "QMessageBox_setEscapeButton2")
	symaddr(&addr_QMessageBox_clickedButton, "QMessageBox_clickedButton")
	symaddr(&addr_QMessageBox_text, "QMessageBox_text")
	symaddr(&addr_QMessageBox_setText, "QMessageBox_setText")
	symaddr(&addr_QMessageBox_icon, "QMessageBox_icon")
	symaddr(&addr_QMessageBox_setIcon, "QMessageBox_setIcon")
	symaddr(&addr_QMessageBox_iconPixmap, "QMessageBox_iconPixmap")
	symaddr(&addr_QMessageBox_setIconPixmap, "QMessageBox_setIconPixmap")
	symaddr(&addr_QMessageBox_textFormat, "QMessageBox_textFormat")
	symaddr(&addr_QMessageBox_setTextFormat, "QMessageBox_setTextFormat")
	symaddr(&addr_QMessageBox_setTextInteractionFlags, "QMessageBox_setTextInteractionFlags")
	symaddr(&addr_QMessageBox_textInteractionFlags, "QMessageBox_textInteractionFlags")
	symaddr(&addr_QMessageBox_setCheckBox, "QMessageBox_setCheckBox")
	symaddr(&addr_QMessageBox_checkBox, "QMessageBox_checkBox")
	symaddr(&addr_QMessageBox_setOption, "QMessageBox_setOption")
	symaddr(&addr_QMessageBox_testOption, "QMessageBox_testOption")
	symaddr(&addr_QMessageBox_setOptions, "QMessageBox_setOptions")
	symaddr(&addr_QMessageBox_options, "QMessageBox_options")
	symaddr(&addr_QMessageBox_information, "QMessageBox_information")
	symaddr(&addr_QMessageBox_information2, "QMessageBox_information2")
	symaddr(&addr_QMessageBox_question, "QMessageBox_question")
	symaddr(&addr_QMessageBox_question2, "QMessageBox_question2")
	symaddr(&addr_QMessageBox_warning, "QMessageBox_warning")
	symaddr(&addr_QMessageBox_warning2, "QMessageBox_warning2")
	symaddr(&addr_QMessageBox_critical, "QMessageBox_critical")
	symaddr(&addr_QMessageBox_critical2, "QMessageBox_critical2")
	symaddr(&addr_QMessageBox_about, "QMessageBox_about")
	symaddr(&addr_QMessageBox_aboutQt, "QMessageBox_aboutQt")
	symaddr(&addr_QMessageBox_information3, "QMessageBox_information3")
	symaddr(&addr_QMessageBox_information4, "QMessageBox_information4")
	symaddr(&addr_QMessageBox_question3, "QMessageBox_question3")
	symaddr(&addr_QMessageBox_question4, "QMessageBox_question4")
	symaddr(&addr_QMessageBox_warning3, "QMessageBox_warning3")
	symaddr(&addr_QMessageBox_warning4, "QMessageBox_warning4")
	symaddr(&addr_QMessageBox_critical3, "QMessageBox_critical3")
	symaddr(&addr_QMessageBox_critical4, "QMessageBox_critical4")
	symaddr(&addr_QMessageBox_buttonText, "QMessageBox_buttonText")
	symaddr(&addr_QMessageBox_setButtonText, "QMessageBox_setButtonText")
	symaddr(&addr_QMessageBox_informativeText, "QMessageBox_informativeText")
	symaddr(&addr_QMessageBox_setInformativeText, "QMessageBox_setInformativeText")
	symaddr(&addr_QMessageBox_detailedText, "QMessageBox_detailedText")
	symaddr(&addr_QMessageBox_setDetailedText, "QMessageBox_setDetailedText")
	symaddr(&addr_QMessageBox_setWindowTitle, "QMessageBox_setWindowTitle")
	symaddr(&addr_QMessageBox_setWindowModality, "QMessageBox_setWindowModality")
	symaddr(&addr_QMessageBox_standardIcon, "QMessageBox_standardIcon")
	symaddr(&addr_QMessageBox_buttonClicked, "QMessageBox_buttonClicked")
	symaddr(&addr_QMessageBox_adjustPosition, "QMessageBox_adjustPosition")
	symaddr(&addr_QMessageBox_setTabOrder, "QMessageBox_setTabOrder")
	symaddr(&addr_QMessageBox_mouseGrabber, "QMessageBox_mouseGrabber")
	symaddr(&addr_QMessageBox_keyboardGrabber, "QMessageBox_keyboardGrabber")
	symaddr(&addr_QMessageBox_find, "QMessageBox_find")
	symaddr(&addr_QMessageBox_createWindowContainer, "QMessageBox_createWindowContainer")
	symaddr(&addr_QMessageBox_updateMicroFocus, "QMessageBox_updateMicroFocus")
	symaddr(&addr_QMessageBox_create, "QMessageBox_create")
	symaddr(&addr_QMessageBox_destroy, "QMessageBox_destroy")
	symaddr(&addr_QMessageBox_focusNextChild, "QMessageBox_focusNextChild")
	symaddr(&addr_QMessageBox_focusPreviousChild, "QMessageBox_focusPreviousChild")
	symaddr(&addr_QMessageBox_connect, "QMessageBox_connect")
	symaddr(&addr_QMessageBox_connect2, "QMessageBox_connect2")
	symaddr(&addr_QMessageBox_sender, "QMessageBox_sender")
	symaddr(&addr_QMessageBox_senderSignalIndex, "QMessageBox_senderSignalIndex")
	symaddr(&addr_QMessageBox_receivers, "QMessageBox_receivers")
	symaddr(&addr_QMessageBox_isSignalConnected, "QMessageBox_isSignalConnected")
	reg(&pg_QMessageBox_devicePixelRatioFScale, "QMessageBox_devicePixelRatioFScale")
	reg(&pg_QMessageBox_encodeMetricF, "QMessageBox_encodeMetricF")
	reg(&pg_QMessageBox_getDecodedMetricF, "QMessageBox_getDecodedMetricF")
	symaddr(&addr_QMessageBox_tr2, "QMessageBox_tr2")
	symaddr(&addr_QMessageBox_tr3, "QMessageBox_tr3")
	symaddr(&addr_QMessageBox_setOption2, "QMessageBox_setOption2")
	symaddr(&addr_QMessageBox_information5, "QMessageBox_information5")
	symaddr(&addr_QMessageBox_information6, "QMessageBox_information6")
	symaddr(&addr_QMessageBox_information7, "QMessageBox_information7")
	symaddr(&addr_QMessageBox_question5, "QMessageBox_question5")
	symaddr(&addr_QMessageBox_question6, "QMessageBox_question6")
	symaddr(&addr_QMessageBox_warning5, "QMessageBox_warning5")
	symaddr(&addr_QMessageBox_warning6, "QMessageBox_warning6")
	symaddr(&addr_QMessageBox_critical5, "QMessageBox_critical5")
	symaddr(&addr_QMessageBox_critical6, "QMessageBox_critical6")
	symaddr(&addr_QMessageBox_aboutQt2, "QMessageBox_aboutQt2")
	symaddr(&addr_QMessageBox_information8, "QMessageBox_information8")
	symaddr(&addr_QMessageBox_information9, "QMessageBox_information9")
	symaddr(&addr_QMessageBox_information10, "QMessageBox_information10")
	symaddr(&addr_QMessageBox_information11, "QMessageBox_information11")
	symaddr(&addr_QMessageBox_information12, "QMessageBox_information12")
	symaddr(&addr_QMessageBox_information13, "QMessageBox_information13")
	symaddr(&addr_QMessageBox_question7, "QMessageBox_question7")
	symaddr(&addr_QMessageBox_question8, "QMessageBox_question8")
	symaddr(&addr_QMessageBox_question9, "QMessageBox_question9")
	symaddr(&addr_QMessageBox_question10, "QMessageBox_question10")
	symaddr(&addr_QMessageBox_question11, "QMessageBox_question11")
	symaddr(&addr_QMessageBox_question12, "QMessageBox_question12")
	symaddr(&addr_QMessageBox_warning7, "QMessageBox_warning7")
	symaddr(&addr_QMessageBox_warning8, "QMessageBox_warning8")
	symaddr(&addr_QMessageBox_warning9, "QMessageBox_warning9")
	symaddr(&addr_QMessageBox_warning10, "QMessageBox_warning10")
	symaddr(&addr_QMessageBox_warning11, "QMessageBox_warning11")
	symaddr(&addr_QMessageBox_critical7, "QMessageBox_critical7")
	symaddr(&addr_QMessageBox_critical8, "QMessageBox_critical8")
	symaddr(&addr_QMessageBox_critical9, "QMessageBox_critical9")
	symaddr(&addr_QMessageBox_critical10, "QMessageBox_critical10")
	symaddr(&addr_QMessageBox_critical11, "QMessageBox_critical11")
	symaddr(&addr_QMessageBox_createWindowContainer2, "QMessageBox_createWindowContainer2")
	symaddr(&addr_QMessageBox_createWindowContainer3, "QMessageBox_createWindowContainer3")
	symaddr(&addr_QMessageBox_updateMicroFocus2, "QMessageBox_updateMicroFocus2")
	symaddr(&addr_QMessageBox_create2, "QMessageBox_create2")
	symaddr(&addr_QMessageBox_create3, "QMessageBox_create3")
	symaddr(&addr_QMessageBox_create4, "QMessageBox_create4")
	symaddr(&addr_QMessageBox_destroy2, "QMessageBox_destroy2")
	symaddr(&addr_QMessageBox_destroy3, "QMessageBox_destroy3")
	symaddr(&addr_QMessageBox_connect3, "QMessageBox_connect3")
	symaddr(&addr_QMessageBox_connect4, "QMessageBox_connect4")
	symaddr(&addr_QMessageBox_metaObject_default, "QMessageBox_metaObject_default")
	symaddr(&addr_QMessageBox_event_default, "QMessageBox_event_default")
	symaddr(&addr_QMessageBox_resizeEvent_default, "QMessageBox_resizeEvent_default")
	symaddr(&addr_QMessageBox_showEvent_default, "QMessageBox_showEvent_default")
	symaddr(&addr_QMessageBox_closeEvent_default, "QMessageBox_closeEvent_default")
	symaddr(&addr_QMessageBox_keyPressEvent_default, "QMessageBox_keyPressEvent_default")
	symaddr(&addr_QMessageBox_changeEvent_default, "QMessageBox_changeEvent_default")
	symaddr(&addr_QMessageBox_setVisible_default, "QMessageBox_setVisible_default")
	symaddr(&addr_QMessageBox_sizeHint_default, "QMessageBox_sizeHint_default")
	symaddr(&addr_QMessageBox_minimumSizeHint_default, "QMessageBox_minimumSizeHint_default")
	symaddr(&addr_QMessageBox_exec_default, "QMessageBox_exec_default")
	symaddr(&addr_QMessageBox_done_default, "QMessageBox_done_default")
	symaddr(&addr_QMessageBox_accept_default, "QMessageBox_accept_default")
	symaddr(&addr_QMessageBox_reject_default, "QMessageBox_reject_default")
	symaddr(&addr_QMessageBox_contextMenuEvent_default, "QMessageBox_contextMenuEvent_default")
	symaddr(&addr_QMessageBox_eventFilter_default, "QMessageBox_eventFilter_default")
	symaddr(&addr_QMessageBox_devType_default, "QMessageBox_devType_default")
	symaddr(&addr_QMessageBox_heightForWidth_default, "QMessageBox_heightForWidth_default")
	symaddr(&addr_QMessageBox_hasHeightForWidth_default, "QMessageBox_hasHeightForWidth_default")
	symaddr(&addr_QMessageBox_paintEngine_default, "QMessageBox_paintEngine_default")
	symaddr(&addr_QMessageBox_mousePressEvent_default, "QMessageBox_mousePressEvent_default")
	symaddr(&addr_QMessageBox_mouseReleaseEvent_default, "QMessageBox_mouseReleaseEvent_default")
	symaddr(&addr_QMessageBox_mouseDoubleClickEvent_default, "QMessageBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QMessageBox_mouseMoveEvent_default, "QMessageBox_mouseMoveEvent_default")
	symaddr(&addr_QMessageBox_wheelEvent_default, "QMessageBox_wheelEvent_default")
	symaddr(&addr_QMessageBox_keyReleaseEvent_default, "QMessageBox_keyReleaseEvent_default")
	symaddr(&addr_QMessageBox_focusInEvent_default, "QMessageBox_focusInEvent_default")
	symaddr(&addr_QMessageBox_focusOutEvent_default, "QMessageBox_focusOutEvent_default")
	symaddr(&addr_QMessageBox_enterEvent_default, "QMessageBox_enterEvent_default")
	symaddr(&addr_QMessageBox_leaveEvent_default, "QMessageBox_leaveEvent_default")
	symaddr(&addr_QMessageBox_paintEvent_default, "QMessageBox_paintEvent_default")
	symaddr(&addr_QMessageBox_moveEvent_default, "QMessageBox_moveEvent_default")
	symaddr(&addr_QMessageBox_tabletEvent_default, "QMessageBox_tabletEvent_default")
	symaddr(&addr_QMessageBox_actionEvent_default, "QMessageBox_actionEvent_default")
	symaddr(&addr_QMessageBox_dragEnterEvent_default, "QMessageBox_dragEnterEvent_default")
	symaddr(&addr_QMessageBox_dragMoveEvent_default, "QMessageBox_dragMoveEvent_default")
	symaddr(&addr_QMessageBox_dragLeaveEvent_default, "QMessageBox_dragLeaveEvent_default")
	symaddr(&addr_QMessageBox_dropEvent_default, "QMessageBox_dropEvent_default")
	symaddr(&addr_QMessageBox_hideEvent_default, "QMessageBox_hideEvent_default")
	symaddr(&addr_QMessageBox_metric_default, "QMessageBox_metric_default")
	symaddr(&addr_QMessageBox_initPainter_default, "QMessageBox_initPainter_default")
	symaddr(&addr_QMessageBox_redirected_default, "QMessageBox_redirected_default")
	symaddr(&addr_QMessageBox_sharedPainter_default, "QMessageBox_sharedPainter_default")
	symaddr(&addr_QMessageBox_inputMethodEvent_default, "QMessageBox_inputMethodEvent_default")
	symaddr(&addr_QMessageBox_inputMethodQuery_default, "QMessageBox_inputMethodQuery_default")
	symaddr(&addr_QMessageBox_focusNextPrevChild_default, "QMessageBox_focusNextPrevChild_default")
	symaddr(&addr_QMessageBox_timerEvent_default, "QMessageBox_timerEvent_default")
	symaddr(&addr_QMessageBox_childEvent_default, "QMessageBox_childEvent_default")
	symaddr(&addr_QMessageBox_customEvent_default, "QMessageBox_customEvent_default")
	symaddr(&addr_QMessageBox_connectNotify_default, "QMessageBox_connectNotify_default")
	symaddr(&addr_QMessageBox_disconnectNotify_default, "QMessageBox_disconnectNotify_default")
	reg(&pg_QMessageBox_connect_buttonClicked, "QMessageBox_connect_buttonClicked")
	reg(&pg_QMessageBox_connect_finished, "QMessageBox_connect_finished")
	reg(&pg_QMessageBox_connect_accepted, "QMessageBox_connect_accepted")
	reg(&pg_QMessageBox_connect_rejected, "QMessageBox_connect_rejected")
	reg(&pg_QMessageBox_connect_windowTitleChanged, "QMessageBox_connect_windowTitleChanged")
	reg(&pg_QMessageBox_connect_windowIconChanged, "QMessageBox_connect_windowIconChanged")
	reg(&pg_QMessageBox_connect_windowIconTextChanged, "QMessageBox_connect_windowIconTextChanged")
	reg(&pg_QMessageBox_connect_customContextMenuRequested, "QMessageBox_connect_customContextMenuRequested")
	reg(&pg_QMessageBox_connect_destroyed, "QMessageBox_connect_destroyed")
	reg(&pg_QMessageBox_connect_destroyed2, "QMessageBox_connect_destroyed2")
	reg(&pg_QMessageBox_connect_objectNameChanged, "QMessageBox_connect_objectNameChanged")
	reg(&pg_QMessageBox__disconnect, "QMessageBox__disconnect")
	reg(&pg_QMessageBox__set_override, "QMessageBox__set_override")
	reg(&pg_QMessageBox_delete, "QMessageBox_delete")
}

type QMessageBox struct {
	h uintptr
	*QDialog
}

func bindQMessageBox(h uintptr) *QMessageBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QMessageBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMessageBox{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQMessageBox(h uintptr) *QMessageBox { return bindQMessageBox(h) }

func (this *QMessageBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMessageBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMessageBox(h unsafe.Pointer) *QMessageBox { return bindQMessageBox(uintptr(h)) }
func UnsafeNewQMessageBoxOwned(h unsafe.Pointer) *QMessageBox { return newownedQMessageBox(uintptr(h)) }

func (this *QMessageBox) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQMessageBox(parent *QWidget) *QMessageBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox2() *QMessageBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new2, uintptr(__goid))
	__o := newownedQMessageBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox3(icon QMessageBox__Icon, title string, text string) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new3, uintptr(__goid), uintptr(icon), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox4(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new4, uintptr(__goid), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(icon), uintptr(button0), uintptr(button1), uintptr(button2))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox5(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new5, uintptr(__goid), uintptr(icon), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox6(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent *QWidget) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new6, uintptr(__goid), uintptr(icon), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(parent.UnsafePointer()))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox7(icon QMessageBox__Icon, title string, text string, buttons QMessageBox__StandardButton, parent *QWidget, flags WindowType) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new7, uintptr(__goid), uintptr(icon), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox8(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int, parent *QWidget) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new8, uintptr(__goid), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(icon), uintptr(button0), uintptr(button1), uintptr(button2), uintptr(parent.UnsafePointer()))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func NewQMessageBox9(title string, text string, icon QMessageBox__Icon, button0 int, button1 int, button2 int, parent *QWidget, f WindowType) *QMessageBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_new9, uintptr(__goid), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(icon), uintptr(button0), uintptr(button1), uintptr(button2), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQMessageBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMessageBox_virtIdx, QMessageBox_enable)
	return __o
}

func QMessageBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMessageBox) AddButton(button *QAbstractButton, role QMessageBox__ButtonRole) {
	purego.SyscallN(addr_QMessageBox_addButton, this.h, uintptr(button.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QMessageBox) AddButton2(text string, role QMessageBox__ButtonRole) *QPushButton {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMessageBox_addButton2, this.h, text_p, uintptr(len(text)), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQPushButton(__r)
}

func (this *QMessageBox) AddButtonWithButton(button QMessageBox__StandardButton) *QPushButton {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_addButton3, this.h, uintptr(button))
	runtime.KeepAlive(this)
	return bindQPushButton(__r)
}

func (this *QMessageBox) RemoveButton(button *QAbstractButton) {
	purego.SyscallN(addr_QMessageBox_removeButton, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QMessageBox) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QMessageBox_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func (this *QMessageBox) Buttons() []*QAbstractButton {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_buttons, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAbstractButton, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAbstractButton(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMessageBox) ButtonRole(button *QAbstractButton) QMessageBox__ButtonRole {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_buttonRole, this.h, uintptr(button.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
	return QMessageBox__ButtonRole(__r)
}

func (this *QMessageBox) SetStandardButtons(buttons QMessageBox__StandardButton) {
	purego.SyscallN(addr_QMessageBox_setStandardButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) StandardButtons() QMessageBox__StandardButton {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_standardButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMessageBox__StandardButton(__r)
}

func (this *QMessageBox) StandardButton(button *QAbstractButton) QMessageBox__StandardButton {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_standardButton, this.h, uintptr(button.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
	return QMessageBox__StandardButton(__r)
}

func (this *QMessageBox) Button(which QMessageBox__StandardButton) *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_button, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QMessageBox) DefaultButton() *QPushButton {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_defaultButton, this.h)
	runtime.KeepAlive(this)
	return bindQPushButton(__r)
}

func (this *QMessageBox) SetDefaultButton(button *QPushButton) {
	purego.SyscallN(addr_QMessageBox_setDefaultButton, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QMessageBox) SetDefaultButtonWithButton(button QMessageBox__StandardButton) {
	purego.SyscallN(addr_QMessageBox_setDefaultButton2, this.h, uintptr(button))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) EscapeButton() *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_escapeButton, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QMessageBox) SetEscapeButton(button *QAbstractButton) {
	purego.SyscallN(addr_QMessageBox_setEscapeButton, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QMessageBox) SetEscapeButtonWithButton(button QMessageBox__StandardButton) {
	purego.SyscallN(addr_QMessageBox_setEscapeButton2, this.h, uintptr(button))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) ClickedButton() *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_clickedButton, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QMessageBox) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageBox) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMessageBox_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QMessageBox) Icon() QMessageBox__Icon {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_icon, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMessageBox__Icon(__r)
}

func (this *QMessageBox) SetIcon(arg1 QMessageBox__Icon) {
	purego.SyscallN(addr_QMessageBox_setIcon, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) IconPixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_iconPixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QMessageBox) SetIconPixmap(pixmap *QPixmap) {
	purego.SyscallN(addr_QMessageBox_setIconPixmap, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QMessageBox) TextFormat() TextFormat {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_textFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextFormat(__r)
}

func (this *QMessageBox) SetTextFormat(format TextFormat) {
	purego.SyscallN(addr_QMessageBox_setTextFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) SetTextInteractionFlags(flags TextInteractionFlag) {
	purego.SyscallN(addr_QMessageBox_setTextInteractionFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) TextInteractionFlags() TextInteractionFlag {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_textInteractionFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextInteractionFlag(__r)
}

func (this *QMessageBox) SetCheckBox(cb *QCheckBox) {
	purego.SyscallN(addr_QMessageBox_setCheckBox, this.h, uintptr(cb.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cb)
}

func (this *QMessageBox) CheckBox() *QCheckBox {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_checkBox, this.h)
	runtime.KeepAlive(this)
	return bindQCheckBox(__r)
}

func (this *QMessageBox) SetOption(option QMessageBox__Option) {
	purego.SyscallN(addr_QMessageBox_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) TestOption(option QMessageBox__Option) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMessageBox) SetOptions(options QMessageBox__Option) {
	purego.SyscallN(addr_QMessageBox_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Options() QMessageBox__Option {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMessageBox__Option(__r)
}

func QMessageBox_Information(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Information2(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Question(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Question2(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Warning(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Warning2(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Critical(parent *QWidget, title string, text string) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Critical2(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_About(parent *QWidget, title string, text string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMessageBox_about, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
}

func QMessageBox_AboutQt(parent *QWidget) {
	purego.SyscallN(addr_QMessageBox_aboutQt, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
}

func QMessageBox_Information3(parent *QWidget, title string, text string, button0 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Information4(parent *QWidget, title string, text string, button0Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	return int(__r)
}

func QMessageBox_Question3(parent *QWidget, title string, text string, button0 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Question4(parent *QWidget, title string, text string, button0Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	return int(__r)
}

func QMessageBox_Warning3(parent *QWidget, title string, text string, button0 int, button1 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Warning4(parent *QWidget, title string, text string, button0Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	return int(__r)
}

func QMessageBox_Critical3(parent *QWidget, title string, text string, button0 int, button1 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical3, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Critical4(parent *QWidget, title string, text string, button0Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical4, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	return int(__r)
}

func (this *QMessageBox) ButtonText(button int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_buttonText, this.h, uintptr(button), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageBox) SetButtonText(button int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMessageBox_setButtonText, this.h, uintptr(button), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QMessageBox) InformativeText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_informativeText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageBox) SetInformativeText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMessageBox_setInformativeText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QMessageBox) DetailedText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_detailedText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageBox) SetDetailedText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMessageBox_setDetailedText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QMessageBox) SetWindowTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QMessageBox_setWindowTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QMessageBox) SetWindowModality(windowModality WindowModality) {
	purego.SyscallN(addr_QMessageBox_setWindowModality, this.h, uintptr(windowModality))
	runtime.KeepAlive(this)
}

func QMessageBox_StandardIcon(icon QMessageBox__Icon) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_standardIcon, uintptr(icon))
	return newownedQPixmap(__r)
}

func (this *QMessageBox) ButtonClicked(button *QAbstractButton) {
	purego.SyscallN(addr_QMessageBox_buttonClicked, this.h, uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QMessageBox) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QMessageBox_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QMessageBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMessageBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMessageBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_mouseGrabber)
	return bindQWidget(__r)
}

func QMessageBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QMessageBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMessageBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMessageBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QMessageBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Create() {
	purego.SyscallN(addr_QMessageBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Destroy() {
	purego.SyscallN(addr_QMessageBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMessageBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMessageBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMessageBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMessageBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMessageBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMessageBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMessageBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMessageBox_DevicePixelRatioFScale() float64 {
	__r := pg_QMessageBox_devicePixelRatioFScale()
	return __r
}

func QMessageBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMessageBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMessageBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMessageBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMessageBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMessageBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMessageBox) SetOption2(option QMessageBox__Option, on bool) {
	purego.SyscallN(addr_QMessageBox_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QMessageBox_Information5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Information6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(defaultButton))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Information7(parent *QWidget, title string, text string, button0 QMessageBox__StandardButton, button1 QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information7, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Question5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Question6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(defaultButton))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Warning5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Warning6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(defaultButton))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Critical5(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical5, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_Critical6(parent *QWidget, title string, text string, buttons QMessageBox__StandardButton, defaultButton QMessageBox__StandardButton) QMessageBox__StandardButton {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical6, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(buttons), uintptr(defaultButton))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return QMessageBox__StandardButton(__r)
}

func QMessageBox_AboutQt2(parent *QWidget, title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QMessageBox_aboutQt2, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
}

func QMessageBox_Information8(parent *QWidget, title string, text string, button0 int, button1 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information8, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Information9(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information9, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1), uintptr(button2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Information10(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information10, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	return int(__r)
}

func QMessageBox_Information11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information11, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Information12(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information12, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Information13(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_information13, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber), uintptr(escapeButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Question7(parent *QWidget, title string, text string, button0 int, button1 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question7, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Question8(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question8, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1), uintptr(button2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Question9(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question9, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	return int(__r)
}

func QMessageBox_Question10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question10, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Question11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question11, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Question12(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_question12, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber), uintptr(escapeButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Warning7(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning7, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1), uintptr(button2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Warning8(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning8, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	return int(__r)
}

func QMessageBox_Warning9(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning9, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Warning10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning10, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Warning11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_warning11, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber), uintptr(escapeButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Critical7(parent *QWidget, title string, text string, button0 int, button1 int, button2 int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical7, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), uintptr(button0), uintptr(button1), uintptr(button2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	return int(__r)
}

func QMessageBox_Critical8(parent *QWidget, title string, text string, button0Text string, button1Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical8, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	return int(__r)
}

func QMessageBox_Critical9(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical9, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Critical10(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical10, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_Critical11(parent *QWidget, title string, text string, button0Text string, button1Text string, button2Text string, defaultButtonNumber int, escapeButtonNumber int) int {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var button0Text_p uintptr
	if len(button0Text) > 0 {
		button0Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button0Text)))
	}
	var button1Text_p uintptr
	if len(button1Text) > 0 {
		button1Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button1Text)))
	}
	var button2Text_p uintptr
	if len(button2Text) > 0 {
		button2Text_p = uintptr(unsafe.Pointer(unsafe.StringData(button2Text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_critical11, uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), text_p, uintptr(len(text)), button0Text_p, uintptr(len(button0Text)), button1Text_p, uintptr(len(button1Text)), button2Text_p, uintptr(len(button2Text)), uintptr(defaultButtonNumber), uintptr(escapeButtonNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	runtime.KeepAlive(text)
	runtime.KeepAlive(button0Text)
	runtime.KeepAlive(button1Text)
	runtime.KeepAlive(button2Text)
	return int(__r)
}

func QMessageBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMessageBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMessageBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMessageBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMessageBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMessageBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMessageBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMessageBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMessageBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMessageBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMessageBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMessageBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMessageBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMessageBox) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QMessageBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QMessageBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QMessageBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QMessageBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMessageBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMessageBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMessageBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMessageBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMessageBox) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageBox) DoneDefault(arg1 int) {
	purego.SyscallN(addr_QMessageBox_done_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMessageBox) AcceptDefault() {
	purego.SyscallN(addr_QMessageBox_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageBox) RejectDefault() {
	purego.SyscallN(addr_QMessageBox_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageBox) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QMessageBox_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMessageBox) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QMessageBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMessageBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMessageBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMessageBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMessageBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMessageBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMessageBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QMessageBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMessageBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMessageBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMessageBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QMessageBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMessageBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QMessageBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QMessageBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMessageBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QMessageBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QMessageBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QMessageBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QMessageBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QMessageBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QMessageBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMessageBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMessageBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMessageBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMessageBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMessageBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMessageBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMessageBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMessageBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMessageBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QMessageBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMessageBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMessageBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMessageBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMessageBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMessageBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMessageBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMessageBox) OnButtonClicked(fn func(*QAbstractButton)) int64 {
	id := registerCallback(func(button uintptr) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(button))) })
	pg_QMessageBox_connect_buttonClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QMessageBox_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMessageBox_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMessageBox_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMessageBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMessageBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMessageBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMessageBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMessageBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMessageBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMessageBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMessageBox) Disconnect(handle int64) {
	pg_QMessageBox__disconnect(handle)
	unregisterCallback(handle)
}

func QMessageBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case -4302917285655014661:
		return 2
	case -6825369174751916:
		return 3
	case -1989937050019464783:
		return 4
	case 1023791977679793185:
		return 5
	case 4959376063826165943:
		return 6
	case 2110182719375245575:
		return 7
	case 729690261457684101:
		return 8
	case 7049264599837532491:
		return 9
	case -1745322561779819684:
		return 10
	case -2570997490807640431:
		return 11
	case 9106131928185084681:
		return 12
	case 2875220963396170150:
		return 13
	case -6698043222227006015:
		return 14
	case 5217623514630858297:
		return 15
	case -8659679893868524044:
		return 16
	case 4311534060495737339:
		return 17
	case -8266762397995423915:
		return 18
	case 6398007404756346381:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case -4255494004099806581:
		return 23
	case 1919988340934511316:
		return 24
	case 772598475950050299:
		return 25
	case 975043767033552552:
		return 26
	case -385596010576048993:
		return 27
	case 6316270346668237895:
		return 28
	case -551070009623711640:
		return 29
	case 5290442163167458465:
		return 30
	case -9161949150489525396:
		return 31
	case -212068866860701141:
		return 32
	case 1697444140088555581:
		return 33
	case 206122715919917589:
		return 34
	case -1781545137352745978:
		return 35
	case -3139904158373725006:
		return 36
	case 9107058682754591716:
		return 37
	case -199718707876719129:
		return 38
	case -2039166586246443706:
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

func QMessageBox_enable(h uintptr, idx int32) { pg_QMessageBox__set_override(h, idx) }

func (this *QMessageBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMessageBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QMessageBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QMessageBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QMessageBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QMessageBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QMessageBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QMessageBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMessageBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMessageBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMessageBox) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QMessageBox) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(arg1 int32) { fn(int(arg1)) }) // slot "done"
}

func (this *QMessageBox) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QMessageBox) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QMessageBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QMessageBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QMessageBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMessageBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMessageBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMessageBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMessageBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QMessageBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QMessageBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMessageBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QMessageBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QMessageBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMessageBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QMessageBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QMessageBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QMessageBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QMessageBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QMessageBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QMessageBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMessageBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QMessageBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QMessageBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QMessageBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QMessageBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QMessageBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QMessageBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMessageBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMessageBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMessageBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMessageBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMessageBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMessageBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMessageBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMessageBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QMessageBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMessageBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMessageBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMessageBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMessageBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMessageBox_delete(this.h)
	}
	this.UnsafeRelease()
}
