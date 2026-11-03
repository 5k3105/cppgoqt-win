package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QComboBox_chain uintptr
	addr_QComboBox_new uintptr
	addr_QComboBox_new2 uintptr
	addr_QComboBox_tr uintptr
	addr_QComboBox_maxVisibleItems uintptr
	addr_QComboBox_setMaxVisibleItems uintptr
	addr_QComboBox_count uintptr
	addr_QComboBox_setMaxCount uintptr
	addr_QComboBox_maxCount uintptr
	addr_QComboBox_duplicatesEnabled uintptr
	addr_QComboBox_setDuplicatesEnabled uintptr
	addr_QComboBox_setFrame uintptr
	addr_QComboBox_hasFrame uintptr
	addr_QComboBox_findText uintptr
	addr_QComboBox_findData uintptr
	addr_QComboBox_insertPolicy uintptr
	addr_QComboBox_setInsertPolicy uintptr
	addr_QComboBox_sizeAdjustPolicy uintptr
	addr_QComboBox_setSizeAdjustPolicy uintptr
	addr_QComboBox_minimumContentsLength uintptr
	addr_QComboBox_setMinimumContentsLength uintptr
	addr_QComboBox_iconSize uintptr
	addr_QComboBox_setIconSize uintptr
	addr_QComboBox_setPlaceholderText uintptr
	addr_QComboBox_placeholderText uintptr
	addr_QComboBox_isEditable uintptr
	addr_QComboBox_setEditable uintptr
	addr_QComboBox_setLineEdit uintptr
	addr_QComboBox_lineEdit uintptr
	addr_QComboBox_setValidator uintptr
	addr_QComboBox_validator uintptr
	addr_QComboBox_setCompleter uintptr
	addr_QComboBox_completer uintptr
	addr_QComboBox_itemDelegate uintptr
	addr_QComboBox_setItemDelegate uintptr
	addr_QComboBox_model uintptr
	addr_QComboBox_setModel uintptr
	addr_QComboBox_rootModelIndex uintptr
	addr_QComboBox_setRootModelIndex uintptr
	addr_QComboBox_modelColumn uintptr
	addr_QComboBox_setModelColumn uintptr
	addr_QComboBox_labelDrawingMode uintptr
	addr_QComboBox_setLabelDrawingMode uintptr
	addr_QComboBox_currentIndex uintptr
	addr_QComboBox_currentText uintptr
	addr_QComboBox_currentData uintptr
	addr_QComboBox_itemText uintptr
	addr_QComboBox_itemIcon uintptr
	addr_QComboBox_itemData uintptr
	addr_QComboBox_addItem uintptr
	addr_QComboBox_addItem2 uintptr
	addr_QComboBox_addItems uintptr
	addr_QComboBox_insertItem uintptr
	addr_QComboBox_insertItem2 uintptr
	addr_QComboBox_insertItems uintptr
	addr_QComboBox_insertSeparator uintptr
	addr_QComboBox_removeItem uintptr
	addr_QComboBox_setItemText uintptr
	addr_QComboBox_setItemIcon uintptr
	addr_QComboBox_setItemData uintptr
	addr_QComboBox_view uintptr
	addr_QComboBox_setView uintptr
	addr_QComboBox_showPopup uintptr
	addr_QComboBox_hidePopup uintptr
	addr_QComboBox_inputMethodQuery uintptr
	addr_QComboBox_inputMethodQuery2 uintptr
	addr_QComboBox_clear uintptr
	addr_QComboBox_clearEditText uintptr
	addr_QComboBox_setEditText uintptr
	addr_QComboBox_setCurrentIndex uintptr
	addr_QComboBox_setCurrentText uintptr
	addr_QComboBox_editTextChanged uintptr
	addr_QComboBox_activated uintptr
	addr_QComboBox_textActivated uintptr
	addr_QComboBox_highlighted uintptr
	addr_QComboBox_textHighlighted uintptr
	addr_QComboBox_currentIndexChanged uintptr
	addr_QComboBox_currentTextChanged uintptr
	addr_QComboBox_setTabOrder uintptr
	addr_QComboBox_mouseGrabber uintptr
	addr_QComboBox_keyboardGrabber uintptr
	addr_QComboBox_find uintptr
	addr_QComboBox_createWindowContainer uintptr
	addr_QComboBox_updateMicroFocus uintptr
	addr_QComboBox_create uintptr
	addr_QComboBox_destroy uintptr
	addr_QComboBox_focusNextChild uintptr
	addr_QComboBox_focusPreviousChild uintptr
	addr_QComboBox_connect uintptr
	addr_QComboBox_connect2 uintptr
	addr_QComboBox_sender uintptr
	addr_QComboBox_senderSignalIndex uintptr
	addr_QComboBox_receivers uintptr
	addr_QComboBox_isSignalConnected uintptr
	pg_QComboBox_devicePixelRatioFScale func() float64
	pg_QComboBox_encodeMetricF func(int32, float64) int32
	pg_QComboBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QComboBox_tr2 uintptr
	addr_QComboBox_tr3 uintptr
	addr_QComboBox_findText2 uintptr
	addr_QComboBox_findData2 uintptr
	addr_QComboBox_findData3 uintptr
	addr_QComboBox_currentData2 uintptr
	addr_QComboBox_itemData2 uintptr
	addr_QComboBox_addItem3 uintptr
	addr_QComboBox_addItem4 uintptr
	addr_QComboBox_insertItem3 uintptr
	addr_QComboBox_insertItem4 uintptr
	addr_QComboBox_setItemData2 uintptr
	addr_QComboBox_createWindowContainer2 uintptr
	addr_QComboBox_createWindowContainer3 uintptr
	addr_QComboBox_updateMicroFocus2 uintptr
	addr_QComboBox_create2 uintptr
	addr_QComboBox_create3 uintptr
	addr_QComboBox_create4 uintptr
	addr_QComboBox_destroy2 uintptr
	addr_QComboBox_destroy3 uintptr
	addr_QComboBox_connect3 uintptr
	addr_QComboBox_connect4 uintptr
	addr_QComboBox_setModel_default uintptr
	addr_QComboBox_showPopup_default uintptr
	addr_QComboBox_hidePopup_default uintptr
	addr_QComboBox_inputMethodQuery_default uintptr
	addr_QComboBox_inputMethodQuery_default2 uintptr
	addr_QComboBox_metaObject_default uintptr
	addr_QComboBox_sizeHint_default uintptr
	addr_QComboBox_minimumSizeHint_default uintptr
	addr_QComboBox_event_default uintptr
	addr_QComboBox_focusInEvent_default uintptr
	addr_QComboBox_focusOutEvent_default uintptr
	addr_QComboBox_changeEvent_default uintptr
	addr_QComboBox_resizeEvent_default uintptr
	addr_QComboBox_paintEvent_default uintptr
	addr_QComboBox_showEvent_default uintptr
	addr_QComboBox_hideEvent_default uintptr
	addr_QComboBox_mousePressEvent_default uintptr
	addr_QComboBox_mouseReleaseEvent_default uintptr
	addr_QComboBox_keyPressEvent_default uintptr
	addr_QComboBox_keyReleaseEvent_default uintptr
	addr_QComboBox_wheelEvent_default uintptr
	addr_QComboBox_contextMenuEvent_default uintptr
	addr_QComboBox_inputMethodEvent_default uintptr
	addr_QComboBox_initStyleOption_default uintptr
	addr_QComboBox_devType_default uintptr
	addr_QComboBox_setVisible_default uintptr
	addr_QComboBox_heightForWidth_default uintptr
	addr_QComboBox_hasHeightForWidth_default uintptr
	addr_QComboBox_paintEngine_default uintptr
	addr_QComboBox_mouseDoubleClickEvent_default uintptr
	addr_QComboBox_mouseMoveEvent_default uintptr
	addr_QComboBox_enterEvent_default uintptr
	addr_QComboBox_leaveEvent_default uintptr
	addr_QComboBox_moveEvent_default uintptr
	addr_QComboBox_closeEvent_default uintptr
	addr_QComboBox_tabletEvent_default uintptr
	addr_QComboBox_actionEvent_default uintptr
	addr_QComboBox_dragEnterEvent_default uintptr
	addr_QComboBox_dragMoveEvent_default uintptr
	addr_QComboBox_dragLeaveEvent_default uintptr
	addr_QComboBox_dropEvent_default uintptr
	addr_QComboBox_metric_default uintptr
	addr_QComboBox_initPainter_default uintptr
	addr_QComboBox_redirected_default uintptr
	addr_QComboBox_sharedPainter_default uintptr
	addr_QComboBox_focusNextPrevChild_default uintptr
	addr_QComboBox_eventFilter_default uintptr
	addr_QComboBox_timerEvent_default uintptr
	addr_QComboBox_childEvent_default uintptr
	addr_QComboBox_customEvent_default uintptr
	addr_QComboBox_connectNotify_default uintptr
	addr_QComboBox_disconnectNotify_default uintptr
	pg_QComboBox_connect_editTextChanged func(uintptr, int64)
	pg_QComboBox_connect_activated func(uintptr, int64)
	pg_QComboBox_connect_textActivated func(uintptr, int64)
	pg_QComboBox_connect_highlighted func(uintptr, int64)
	pg_QComboBox_connect_textHighlighted func(uintptr, int64)
	pg_QComboBox_connect_currentIndexChanged func(uintptr, int64)
	pg_QComboBox_connect_currentTextChanged func(uintptr, int64)
	pg_QComboBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QComboBox_connect_windowIconChanged func(uintptr, int64)
	pg_QComboBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QComboBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QComboBox_connect_destroyed func(uintptr, int64)
	pg_QComboBox_connect_destroyed2 func(uintptr, int64)
	pg_QComboBox_connect_objectNameChanged func(uintptr, int64)
	pg_QComboBox__disconnect func(int64)
	pg_QComboBox__set_override func(uintptr, int32)
	pg_QComboBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QComboBox_chain, "QComboBox_chain")
	symaddr(&addr_QComboBox_new, "QComboBox_new")
	symaddr(&addr_QComboBox_new2, "QComboBox_new2")
	symaddr(&addr_QComboBox_tr, "QComboBox_tr")
	symaddr(&addr_QComboBox_maxVisibleItems, "QComboBox_maxVisibleItems")
	symaddr(&addr_QComboBox_setMaxVisibleItems, "QComboBox_setMaxVisibleItems")
	symaddr(&addr_QComboBox_count, "QComboBox_count")
	symaddr(&addr_QComboBox_setMaxCount, "QComboBox_setMaxCount")
	symaddr(&addr_QComboBox_maxCount, "QComboBox_maxCount")
	symaddr(&addr_QComboBox_duplicatesEnabled, "QComboBox_duplicatesEnabled")
	symaddr(&addr_QComboBox_setDuplicatesEnabled, "QComboBox_setDuplicatesEnabled")
	symaddr(&addr_QComboBox_setFrame, "QComboBox_setFrame")
	symaddr(&addr_QComboBox_hasFrame, "QComboBox_hasFrame")
	symaddr(&addr_QComboBox_findText, "QComboBox_findText")
	symaddr(&addr_QComboBox_findData, "QComboBox_findData")
	symaddr(&addr_QComboBox_insertPolicy, "QComboBox_insertPolicy")
	symaddr(&addr_QComboBox_setInsertPolicy, "QComboBox_setInsertPolicy")
	symaddr(&addr_QComboBox_sizeAdjustPolicy, "QComboBox_sizeAdjustPolicy")
	symaddr(&addr_QComboBox_setSizeAdjustPolicy, "QComboBox_setSizeAdjustPolicy")
	symaddr(&addr_QComboBox_minimumContentsLength, "QComboBox_minimumContentsLength")
	symaddr(&addr_QComboBox_setMinimumContentsLength, "QComboBox_setMinimumContentsLength")
	symaddr(&addr_QComboBox_iconSize, "QComboBox_iconSize")
	symaddr(&addr_QComboBox_setIconSize, "QComboBox_setIconSize")
	symaddr(&addr_QComboBox_setPlaceholderText, "QComboBox_setPlaceholderText")
	symaddr(&addr_QComboBox_placeholderText, "QComboBox_placeholderText")
	symaddr(&addr_QComboBox_isEditable, "QComboBox_isEditable")
	symaddr(&addr_QComboBox_setEditable, "QComboBox_setEditable")
	symaddr(&addr_QComboBox_setLineEdit, "QComboBox_setLineEdit")
	symaddr(&addr_QComboBox_lineEdit, "QComboBox_lineEdit")
	symaddr(&addr_QComboBox_setValidator, "QComboBox_setValidator")
	symaddr(&addr_QComboBox_validator, "QComboBox_validator")
	symaddr(&addr_QComboBox_setCompleter, "QComboBox_setCompleter")
	symaddr(&addr_QComboBox_completer, "QComboBox_completer")
	symaddr(&addr_QComboBox_itemDelegate, "QComboBox_itemDelegate")
	symaddr(&addr_QComboBox_setItemDelegate, "QComboBox_setItemDelegate")
	symaddr(&addr_QComboBox_model, "QComboBox_model")
	symaddr(&addr_QComboBox_setModel, "QComboBox_setModel")
	symaddr(&addr_QComboBox_rootModelIndex, "QComboBox_rootModelIndex")
	symaddr(&addr_QComboBox_setRootModelIndex, "QComboBox_setRootModelIndex")
	symaddr(&addr_QComboBox_modelColumn, "QComboBox_modelColumn")
	symaddr(&addr_QComboBox_setModelColumn, "QComboBox_setModelColumn")
	symaddr(&addr_QComboBox_labelDrawingMode, "QComboBox_labelDrawingMode")
	symaddr(&addr_QComboBox_setLabelDrawingMode, "QComboBox_setLabelDrawingMode")
	symaddr(&addr_QComboBox_currentIndex, "QComboBox_currentIndex")
	symaddr(&addr_QComboBox_currentText, "QComboBox_currentText")
	symaddr(&addr_QComboBox_currentData, "QComboBox_currentData")
	symaddr(&addr_QComboBox_itemText, "QComboBox_itemText")
	symaddr(&addr_QComboBox_itemIcon, "QComboBox_itemIcon")
	symaddr(&addr_QComboBox_itemData, "QComboBox_itemData")
	symaddr(&addr_QComboBox_addItem, "QComboBox_addItem")
	symaddr(&addr_QComboBox_addItem2, "QComboBox_addItem2")
	symaddr(&addr_QComboBox_addItems, "QComboBox_addItems")
	symaddr(&addr_QComboBox_insertItem, "QComboBox_insertItem")
	symaddr(&addr_QComboBox_insertItem2, "QComboBox_insertItem2")
	symaddr(&addr_QComboBox_insertItems, "QComboBox_insertItems")
	symaddr(&addr_QComboBox_insertSeparator, "QComboBox_insertSeparator")
	symaddr(&addr_QComboBox_removeItem, "QComboBox_removeItem")
	symaddr(&addr_QComboBox_setItemText, "QComboBox_setItemText")
	symaddr(&addr_QComboBox_setItemIcon, "QComboBox_setItemIcon")
	symaddr(&addr_QComboBox_setItemData, "QComboBox_setItemData")
	symaddr(&addr_QComboBox_view, "QComboBox_view")
	symaddr(&addr_QComboBox_setView, "QComboBox_setView")
	symaddr(&addr_QComboBox_showPopup, "QComboBox_showPopup")
	symaddr(&addr_QComboBox_hidePopup, "QComboBox_hidePopup")
	symaddr(&addr_QComboBox_inputMethodQuery, "QComboBox_inputMethodQuery")
	symaddr(&addr_QComboBox_inputMethodQuery2, "QComboBox_inputMethodQuery2")
	symaddr(&addr_QComboBox_clear, "QComboBox_clear")
	symaddr(&addr_QComboBox_clearEditText, "QComboBox_clearEditText")
	symaddr(&addr_QComboBox_setEditText, "QComboBox_setEditText")
	symaddr(&addr_QComboBox_setCurrentIndex, "QComboBox_setCurrentIndex")
	symaddr(&addr_QComboBox_setCurrentText, "QComboBox_setCurrentText")
	symaddr(&addr_QComboBox_editTextChanged, "QComboBox_editTextChanged")
	symaddr(&addr_QComboBox_activated, "QComboBox_activated")
	symaddr(&addr_QComboBox_textActivated, "QComboBox_textActivated")
	symaddr(&addr_QComboBox_highlighted, "QComboBox_highlighted")
	symaddr(&addr_QComboBox_textHighlighted, "QComboBox_textHighlighted")
	symaddr(&addr_QComboBox_currentIndexChanged, "QComboBox_currentIndexChanged")
	symaddr(&addr_QComboBox_currentTextChanged, "QComboBox_currentTextChanged")
	symaddr(&addr_QComboBox_setTabOrder, "QComboBox_setTabOrder")
	symaddr(&addr_QComboBox_mouseGrabber, "QComboBox_mouseGrabber")
	symaddr(&addr_QComboBox_keyboardGrabber, "QComboBox_keyboardGrabber")
	symaddr(&addr_QComboBox_find, "QComboBox_find")
	symaddr(&addr_QComboBox_createWindowContainer, "QComboBox_createWindowContainer")
	symaddr(&addr_QComboBox_updateMicroFocus, "QComboBox_updateMicroFocus")
	symaddr(&addr_QComboBox_create, "QComboBox_create")
	symaddr(&addr_QComboBox_destroy, "QComboBox_destroy")
	symaddr(&addr_QComboBox_focusNextChild, "QComboBox_focusNextChild")
	symaddr(&addr_QComboBox_focusPreviousChild, "QComboBox_focusPreviousChild")
	symaddr(&addr_QComboBox_connect, "QComboBox_connect")
	symaddr(&addr_QComboBox_connect2, "QComboBox_connect2")
	symaddr(&addr_QComboBox_sender, "QComboBox_sender")
	symaddr(&addr_QComboBox_senderSignalIndex, "QComboBox_senderSignalIndex")
	symaddr(&addr_QComboBox_receivers, "QComboBox_receivers")
	symaddr(&addr_QComboBox_isSignalConnected, "QComboBox_isSignalConnected")
	reg(&pg_QComboBox_devicePixelRatioFScale, "QComboBox_devicePixelRatioFScale")
	reg(&pg_QComboBox_encodeMetricF, "QComboBox_encodeMetricF")
	reg(&pg_QComboBox_getDecodedMetricF, "QComboBox_getDecodedMetricF")
	symaddr(&addr_QComboBox_tr2, "QComboBox_tr2")
	symaddr(&addr_QComboBox_tr3, "QComboBox_tr3")
	symaddr(&addr_QComboBox_findText2, "QComboBox_findText2")
	symaddr(&addr_QComboBox_findData2, "QComboBox_findData2")
	symaddr(&addr_QComboBox_findData3, "QComboBox_findData3")
	symaddr(&addr_QComboBox_currentData2, "QComboBox_currentData2")
	symaddr(&addr_QComboBox_itemData2, "QComboBox_itemData2")
	symaddr(&addr_QComboBox_addItem3, "QComboBox_addItem3")
	symaddr(&addr_QComboBox_addItem4, "QComboBox_addItem4")
	symaddr(&addr_QComboBox_insertItem3, "QComboBox_insertItem3")
	symaddr(&addr_QComboBox_insertItem4, "QComboBox_insertItem4")
	symaddr(&addr_QComboBox_setItemData2, "QComboBox_setItemData2")
	symaddr(&addr_QComboBox_createWindowContainer2, "QComboBox_createWindowContainer2")
	symaddr(&addr_QComboBox_createWindowContainer3, "QComboBox_createWindowContainer3")
	symaddr(&addr_QComboBox_updateMicroFocus2, "QComboBox_updateMicroFocus2")
	symaddr(&addr_QComboBox_create2, "QComboBox_create2")
	symaddr(&addr_QComboBox_create3, "QComboBox_create3")
	symaddr(&addr_QComboBox_create4, "QComboBox_create4")
	symaddr(&addr_QComboBox_destroy2, "QComboBox_destroy2")
	symaddr(&addr_QComboBox_destroy3, "QComboBox_destroy3")
	symaddr(&addr_QComboBox_connect3, "QComboBox_connect3")
	symaddr(&addr_QComboBox_connect4, "QComboBox_connect4")
	symaddr(&addr_QComboBox_setModel_default, "QComboBox_setModel_default")
	symaddr(&addr_QComboBox_showPopup_default, "QComboBox_showPopup_default")
	symaddr(&addr_QComboBox_hidePopup_default, "QComboBox_hidePopup_default")
	symaddr(&addr_QComboBox_inputMethodQuery_default, "QComboBox_inputMethodQuery_default")
	symaddr(&addr_QComboBox_inputMethodQuery_default2, "QComboBox_inputMethodQuery_default2")
	symaddr(&addr_QComboBox_metaObject_default, "QComboBox_metaObject_default")
	symaddr(&addr_QComboBox_sizeHint_default, "QComboBox_sizeHint_default")
	symaddr(&addr_QComboBox_minimumSizeHint_default, "QComboBox_minimumSizeHint_default")
	symaddr(&addr_QComboBox_event_default, "QComboBox_event_default")
	symaddr(&addr_QComboBox_focusInEvent_default, "QComboBox_focusInEvent_default")
	symaddr(&addr_QComboBox_focusOutEvent_default, "QComboBox_focusOutEvent_default")
	symaddr(&addr_QComboBox_changeEvent_default, "QComboBox_changeEvent_default")
	symaddr(&addr_QComboBox_resizeEvent_default, "QComboBox_resizeEvent_default")
	symaddr(&addr_QComboBox_paintEvent_default, "QComboBox_paintEvent_default")
	symaddr(&addr_QComboBox_showEvent_default, "QComboBox_showEvent_default")
	symaddr(&addr_QComboBox_hideEvent_default, "QComboBox_hideEvent_default")
	symaddr(&addr_QComboBox_mousePressEvent_default, "QComboBox_mousePressEvent_default")
	symaddr(&addr_QComboBox_mouseReleaseEvent_default, "QComboBox_mouseReleaseEvent_default")
	symaddr(&addr_QComboBox_keyPressEvent_default, "QComboBox_keyPressEvent_default")
	symaddr(&addr_QComboBox_keyReleaseEvent_default, "QComboBox_keyReleaseEvent_default")
	symaddr(&addr_QComboBox_wheelEvent_default, "QComboBox_wheelEvent_default")
	symaddr(&addr_QComboBox_contextMenuEvent_default, "QComboBox_contextMenuEvent_default")
	symaddr(&addr_QComboBox_inputMethodEvent_default, "QComboBox_inputMethodEvent_default")
	symaddr(&addr_QComboBox_initStyleOption_default, "QComboBox_initStyleOption_default")
	symaddr(&addr_QComboBox_devType_default, "QComboBox_devType_default")
	symaddr(&addr_QComboBox_setVisible_default, "QComboBox_setVisible_default")
	symaddr(&addr_QComboBox_heightForWidth_default, "QComboBox_heightForWidth_default")
	symaddr(&addr_QComboBox_hasHeightForWidth_default, "QComboBox_hasHeightForWidth_default")
	symaddr(&addr_QComboBox_paintEngine_default, "QComboBox_paintEngine_default")
	symaddr(&addr_QComboBox_mouseDoubleClickEvent_default, "QComboBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QComboBox_mouseMoveEvent_default, "QComboBox_mouseMoveEvent_default")
	symaddr(&addr_QComboBox_enterEvent_default, "QComboBox_enterEvent_default")
	symaddr(&addr_QComboBox_leaveEvent_default, "QComboBox_leaveEvent_default")
	symaddr(&addr_QComboBox_moveEvent_default, "QComboBox_moveEvent_default")
	symaddr(&addr_QComboBox_closeEvent_default, "QComboBox_closeEvent_default")
	symaddr(&addr_QComboBox_tabletEvent_default, "QComboBox_tabletEvent_default")
	symaddr(&addr_QComboBox_actionEvent_default, "QComboBox_actionEvent_default")
	symaddr(&addr_QComboBox_dragEnterEvent_default, "QComboBox_dragEnterEvent_default")
	symaddr(&addr_QComboBox_dragMoveEvent_default, "QComboBox_dragMoveEvent_default")
	symaddr(&addr_QComboBox_dragLeaveEvent_default, "QComboBox_dragLeaveEvent_default")
	symaddr(&addr_QComboBox_dropEvent_default, "QComboBox_dropEvent_default")
	symaddr(&addr_QComboBox_metric_default, "QComboBox_metric_default")
	symaddr(&addr_QComboBox_initPainter_default, "QComboBox_initPainter_default")
	symaddr(&addr_QComboBox_redirected_default, "QComboBox_redirected_default")
	symaddr(&addr_QComboBox_sharedPainter_default, "QComboBox_sharedPainter_default")
	symaddr(&addr_QComboBox_focusNextPrevChild_default, "QComboBox_focusNextPrevChild_default")
	symaddr(&addr_QComboBox_eventFilter_default, "QComboBox_eventFilter_default")
	symaddr(&addr_QComboBox_timerEvent_default, "QComboBox_timerEvent_default")
	symaddr(&addr_QComboBox_childEvent_default, "QComboBox_childEvent_default")
	symaddr(&addr_QComboBox_customEvent_default, "QComboBox_customEvent_default")
	symaddr(&addr_QComboBox_connectNotify_default, "QComboBox_connectNotify_default")
	symaddr(&addr_QComboBox_disconnectNotify_default, "QComboBox_disconnectNotify_default")
	reg(&pg_QComboBox_connect_editTextChanged, "QComboBox_connect_editTextChanged")
	reg(&pg_QComboBox_connect_activated, "QComboBox_connect_activated")
	reg(&pg_QComboBox_connect_textActivated, "QComboBox_connect_textActivated")
	reg(&pg_QComboBox_connect_highlighted, "QComboBox_connect_highlighted")
	reg(&pg_QComboBox_connect_textHighlighted, "QComboBox_connect_textHighlighted")
	reg(&pg_QComboBox_connect_currentIndexChanged, "QComboBox_connect_currentIndexChanged")
	reg(&pg_QComboBox_connect_currentTextChanged, "QComboBox_connect_currentTextChanged")
	reg(&pg_QComboBox_connect_windowTitleChanged, "QComboBox_connect_windowTitleChanged")
	reg(&pg_QComboBox_connect_windowIconChanged, "QComboBox_connect_windowIconChanged")
	reg(&pg_QComboBox_connect_windowIconTextChanged, "QComboBox_connect_windowIconTextChanged")
	reg(&pg_QComboBox_connect_customContextMenuRequested, "QComboBox_connect_customContextMenuRequested")
	reg(&pg_QComboBox_connect_destroyed, "QComboBox_connect_destroyed")
	reg(&pg_QComboBox_connect_destroyed2, "QComboBox_connect_destroyed2")
	reg(&pg_QComboBox_connect_objectNameChanged, "QComboBox_connect_objectNameChanged")
	reg(&pg_QComboBox__disconnect, "QComboBox__disconnect")
	reg(&pg_QComboBox__set_override, "QComboBox__set_override")
	reg(&pg_QComboBox_delete, "QComboBox_delete")
}

type QComboBox struct {
	h uintptr
	*QWidget
}

func bindQComboBox(h uintptr) *QComboBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QComboBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QComboBox{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQComboBox(h uintptr) *QComboBox { return bindQComboBox(h) }

func (this *QComboBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QComboBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQComboBox(h unsafe.Pointer) *QComboBox { return bindQComboBox(uintptr(h)) }
func UnsafeNewQComboBoxOwned(h unsafe.Pointer) *QComboBox { return newownedQComboBox(uintptr(h)) }

func (this *QComboBox) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQComboBox(parent *QWidget) *QComboBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QComboBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQComboBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QComboBox_virtIdx, QComboBox_enable)
	return __o
}

func NewQComboBox2() *QComboBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QComboBox_new2, uintptr(__goid))
	__o := newownedQComboBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QComboBox_virtIdx, QComboBox_enable)
	return __o
}

func QComboBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QComboBox) MaxVisibleItems() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_maxVisibleItems, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	purego.SyscallN(addr_QComboBox_setMaxVisibleItems, this.h, uintptr(maxItems))
	runtime.KeepAlive(this)
}

func (this *QComboBox) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) SetMaxCount(max int) {
	purego.SyscallN(addr_QComboBox_setMaxCount, this.h, uintptr(max))
	runtime.KeepAlive(this)
}

func (this *QComboBox) MaxCount() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_maxCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) DuplicatesEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_duplicatesEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	purego.SyscallN(addr_QComboBox_setDuplicatesEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) SetFrame(arg1 bool) {
	purego.SyscallN(addr_QComboBox_setFrame, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) HasFrame() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_hasFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) FindText(text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QComboBox_findText, this.h, text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QComboBox) FindData(data *QVariant) int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_findData, this.h, uintptr(data.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return int(__r)
}

func (this *QComboBox) InsertPolicy() QComboBox__InsertPolicy {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_insertPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QComboBox__InsertPolicy(__r)
}

func (this *QComboBox) SetInsertPolicy(policy QComboBox__InsertPolicy) {
	purego.SyscallN(addr_QComboBox_setInsertPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QComboBox) SizeAdjustPolicy() QComboBox__SizeAdjustPolicy {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_sizeAdjustPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QComboBox__SizeAdjustPolicy(__r)
}

func (this *QComboBox) SetSizeAdjustPolicy(policy QComboBox__SizeAdjustPolicy) {
	purego.SyscallN(addr_QComboBox_setSizeAdjustPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QComboBox) MinimumContentsLength() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_minimumContentsLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) SetMinimumContentsLength(characters int) {
	purego.SyscallN(addr_QComboBox_setMinimumContentsLength, this.h, uintptr(characters))
	runtime.KeepAlive(this)
}

func (this *QComboBox) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QComboBox_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QComboBox) SetIconSize(size *QSize) {
	purego.SyscallN(addr_QComboBox_setIconSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QComboBox) SetPlaceholderText(placeholderText string) {
	var placeholderText_p uintptr
	if len(placeholderText) > 0 {
		placeholderText_p = uintptr(unsafe.Pointer(unsafe.StringData(placeholderText)))
	}
	purego.SyscallN(addr_QComboBox_setPlaceholderText, this.h, placeholderText_p, uintptr(len(placeholderText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(placeholderText)
}

func (this *QComboBox) PlaceholderText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_placeholderText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QComboBox) IsEditable() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_isEditable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) SetEditable(editable bool) {
	purego.SyscallN(addr_QComboBox_setEditable, this.h, uintptr(b2i(editable)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QComboBox_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func (this *QComboBox) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QComboBox_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QComboBox) SetValidator(v *QValidator) {
	purego.SyscallN(addr_QComboBox_setValidator, this.h, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v)
}

func (this *QComboBox) Validator() *QValidator {
	__r, _, _ := purego.SyscallN(addr_QComboBox_validator, this.h)
	runtime.KeepAlive(this)
	return bindQValidator(__r)
}

func (this *QComboBox) SetCompleter(c *QCompleter) {
	purego.SyscallN(addr_QComboBox_setCompleter, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QComboBox) Completer() *QCompleter {
	__r, _, _ := purego.SyscallN(addr_QComboBox_completer, this.h)
	runtime.KeepAlive(this)
	return bindQCompleter(__r)
}

func (this *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QComboBox_itemDelegate, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QComboBox) SetItemDelegate(delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QComboBox_setItemDelegate, this.h, uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QComboBox) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QComboBox_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QComboBox) SetModel(model *QAbstractItemModel) {
	purego.SyscallN(addr_QComboBox_setModel, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QComboBox) RootModelIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QComboBox_rootModelIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QComboBox) SetRootModelIndex(index *QModelIndex) {
	purego.SyscallN(addr_QComboBox_setRootModelIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QComboBox) ModelColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_modelColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) SetModelColumn(visibleColumn int) {
	purego.SyscallN(addr_QComboBox_setModelColumn, this.h, uintptr(visibleColumn))
	runtime.KeepAlive(this)
}

func (this *QComboBox) LabelDrawingMode() QComboBox__LabelDrawingMode {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_labelDrawingMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QComboBox__LabelDrawingMode(__r)
}

func (this *QComboBox) SetLabelDrawingMode(labelDrawing QComboBox__LabelDrawingMode) {
	purego.SyscallN(addr_QComboBox_setLabelDrawingMode, this.h, uintptr(labelDrawing))
	runtime.KeepAlive(this)
}

func (this *QComboBox) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) CurrentText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_currentText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QComboBox) CurrentData() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_currentData, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) ItemText(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_itemText, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QComboBox) ItemIcon(index int) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QComboBox_itemIcon, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QComboBox) ItemData(index int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_itemData, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) AddItem(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_addItem, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QComboBox) AddItem2(icon *QIcon, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_addItem2, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
}

func (this *QComboBox) AddItems(texts []string) {
	texts_arr := make([]uintptr, len(texts))
	for __i, __e := range texts {
		texts_arr[__i] = pgCString(__e)
	}
	var texts_p uintptr
	if len(texts_arr) > 0 {
		texts_p = uintptr(unsafe.Pointer(&texts_arr[0]))
	}
	purego.SyscallN(addr_QComboBox_addItems, this.h, texts_p, uintptr(len(texts)))
	runtime.KeepAlive(this)
	for __i := range texts_arr {
		pgFree(texts_arr[__i])
	}
	runtime.KeepAlive(texts_arr)
}

func (this *QComboBox) InsertItem(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_insertItem, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QComboBox) InsertItem2(index int, icon *QIcon, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_insertItem2, this.h, uintptr(index), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
}

func (this *QComboBox) InsertItems(index int, texts []string) {
	texts_arr := make([]uintptr, len(texts))
	for __i, __e := range texts {
		texts_arr[__i] = pgCString(__e)
	}
	var texts_p uintptr
	if len(texts_arr) > 0 {
		texts_p = uintptr(unsafe.Pointer(&texts_arr[0]))
	}
	purego.SyscallN(addr_QComboBox_insertItems, this.h, uintptr(index), texts_p, uintptr(len(texts)))
	runtime.KeepAlive(this)
	for __i := range texts_arr {
		pgFree(texts_arr[__i])
	}
	runtime.KeepAlive(texts_arr)
}

func (this *QComboBox) InsertSeparator(index int) {
	purego.SyscallN(addr_QComboBox_insertSeparator, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) RemoveItem(index int) {
	purego.SyscallN(addr_QComboBox_removeItem, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) SetItemText(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_setItemText, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QComboBox) SetItemIcon(index int, icon *QIcon) {
	purego.SyscallN(addr_QComboBox_setItemIcon, this.h, uintptr(index), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QComboBox) SetItemData(index int, value *QVariant) {
	purego.SyscallN(addr_QComboBox_setItemData, this.h, uintptr(index), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QComboBox) View() *QAbstractItemView {
	__r, _, _ := purego.SyscallN(addr_QComboBox_view, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemView(__r)
}

func (this *QComboBox) SetView(itemView *QAbstractItemView) {
	purego.SyscallN(addr_QComboBox_setView, this.h, uintptr(itemView.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(itemView)
}

func (this *QComboBox) ShowPopup() {
	purego.SyscallN(addr_QComboBox_showPopup, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) HidePopup() {
	purego.SyscallN(addr_QComboBox_hidePopup, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) InputMethodQuery(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_inputMethodQuery, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) InputMethodQuery2(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_inputMethodQuery2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QComboBox) Clear() {
	purego.SyscallN(addr_QComboBox_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) ClearEditText() {
	purego.SyscallN(addr_QComboBox_clearEditText, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) SetEditText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_setEditText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QComboBox) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QComboBox_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) SetCurrentText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_setCurrentText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QComboBox) EditTextChanged(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QComboBox_editTextChanged, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QComboBox) Activated(index int) {
	purego.SyscallN(addr_QComboBox_activated, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) TextActivated(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QComboBox_textActivated, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QComboBox) Highlighted(index int) {
	purego.SyscallN(addr_QComboBox_highlighted, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) TextHighlighted(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QComboBox_textHighlighted, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QComboBox) CurrentIndexChanged(index int) {
	purego.SyscallN(addr_QComboBox_currentIndexChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QComboBox) CurrentTextChanged(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QComboBox_currentTextChanged, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QComboBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QComboBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QComboBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_mouseGrabber)
	return bindQWidget(__r)
}

func QComboBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QComboBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QComboBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QComboBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QComboBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) Create() {
	purego.SyscallN(addr_QComboBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) Destroy() {
	purego.SyscallN(addr_QComboBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QComboBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QComboBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QComboBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QComboBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QComboBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QComboBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QComboBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QComboBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QComboBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QComboBox_DevicePixelRatioFScale() float64 {
	__r := pg_QComboBox_devicePixelRatioFScale()
	return __r
}

func QComboBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QComboBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QComboBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QComboBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QComboBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QComboBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QComboBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QComboBox) FindText2(text string, flags MatchFlag) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QComboBox_findText2, this.h, text_p, uintptr(len(text)), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QComboBox) FindData2(data *QVariant, role int) int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_findData2, this.h, uintptr(data.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return int(__r)
}

func (this *QComboBox) FindData3(data *QVariant, role int, flags MatchFlag) int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_findData3, this.h, uintptr(data.UnsafePointer()), uintptr(role), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return int(__r)
}

func (this *QComboBox) CurrentDataWithRole(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_currentData2, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) ItemData2(index int, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_itemData2, this.h, uintptr(index), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) AddItem3(text string, userData *QVariant) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_addItem3, this.h, text_p, uintptr(len(text)), uintptr(userData.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(userData)
}

func (this *QComboBox) AddItem4(icon *QIcon, text string, userData *QVariant) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_addItem4, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(userData.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(userData)
}

func (this *QComboBox) InsertItem3(index int, text string, userData *QVariant) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_insertItem3, this.h, uintptr(index), text_p, uintptr(len(text)), uintptr(userData.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(userData)
}

func (this *QComboBox) InsertItem4(index int, icon *QIcon, text string, userData *QVariant) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QComboBox_insertItem4, this.h, uintptr(index), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(userData.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(userData)
}

func (this *QComboBox) SetItemData2(index int, value *QVariant, role int) {
	purego.SyscallN(addr_QComboBox_setItemData2, this.h, uintptr(index), uintptr(value.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func QComboBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QComboBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QComboBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QComboBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QComboBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QComboBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QComboBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QComboBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QComboBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QComboBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QComboBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QComboBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QComboBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QComboBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QComboBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QComboBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QComboBox) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QComboBox_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QComboBox) ShowPopupDefault() {
	purego.SyscallN(addr_QComboBox_showPopup_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) HidePopupDefault() {
	purego.SyscallN(addr_QComboBox_hidePopup_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QComboBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QComboBox) InputMethodQuery2Default(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QComboBox_inputMethodQuery_default2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QComboBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QComboBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QComboBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QComboBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QComboBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QComboBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QComboBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QComboBox) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QComboBox_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QComboBox_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QComboBox_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QComboBox_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QComboBox_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) ShowEventDefault(e *QShowEvent) {
	purego.SyscallN(addr_QComboBox_showEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) HideEventDefault(e *QHideEvent) {
	purego.SyscallN(addr_QComboBox_hideEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QComboBox_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QComboBox_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QComboBox_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QComboBox_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QComboBox_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) ContextMenuEventDefault(e *QContextMenuEvent) {
	purego.SyscallN(addr_QComboBox_contextMenuEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QComboBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QComboBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QComboBox) InitStyleOptionDefault(option *QStyleOptionComboBox) {
	purego.SyscallN(addr_QComboBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QComboBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QComboBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QComboBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QComboBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QComboBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QComboBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QComboBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QComboBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QComboBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QComboBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QComboBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QComboBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QComboBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QComboBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QComboBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QComboBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QComboBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QComboBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QComboBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QComboBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QComboBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QComboBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QComboBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QComboBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QComboBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QComboBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QComboBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QComboBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QComboBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QComboBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QComboBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QComboBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QComboBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QComboBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QComboBox) OnEditTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QComboBox_connect_editTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnActivated(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QComboBox_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnTextActivated(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QComboBox_connect_textActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnHighlighted(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QComboBox_connect_highlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnTextHighlighted(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QComboBox_connect_textHighlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnCurrentIndexChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QComboBox_connect_currentIndexChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnCurrentTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QComboBox_connect_currentTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QComboBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QComboBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QComboBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QComboBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QComboBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QComboBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QComboBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QComboBox) Disconnect(handle int64) {
	pg_QComboBox__disconnect(handle)
	unregisterCallback(handle)
}

func QComboBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 4283922507359547360:
		return 1
	case 729690261457684101:
		return 2
	case 7049264599837532491:
		return 3
	case -977583612921719482:
		return 4
	case 3308350117881191909:
		return 5
	case 1237236456312824991:
		return 6
	case 7345853097340709638:
		return 7
	case 975043767033552552:
		return 8
	case -385596010576048993:
		return 9
	case 4959376063826165943:
		return 10
	case -4302917285655014661:
		return 11
	case 5290442163167458465:
		return 12
	case -6825369174751916:
		return 13
	case -199718707876719129:
		return 14
	case -332090898078543319:
		return 15
	case -4415920049297795213:
		return 16
	case 1023791977679793185:
		return 17
	case 772598475950050299:
		return 18
	case 1919988340934511316:
		return 19
	case -6698043222227006015:
		return 20
	case 5438819661942843222:
		return 21
	case -488693769771308531:
		return 22
	case -8659679893868524044:
		return 23
	case 2110182719375245575:
		return 24
	case 4311534060495737339:
		return 25
	case -8266762397995423915:
		return 26
	case 6398007404756346381:
		return 27
	case -4029073905697023887:
		return 28
	case -4255494004099806581:
		return 29
	case 6316270346668237895:
		return 30
	case -551070009623711640:
		return 31
	case -9161949150489525396:
		return 32
	case -1989937050019464783:
		return 33
	case -212068866860701141:
		return 34
	case 1697444140088555581:
		return 35
	case 206122715919917589:
		return 36
	case -1781545137352745978:
		return 37
	case -3139904158373725006:
		return 38
	case 9107058682754591716:
		return 39
	case -2039166586246443706:
		return 40
	case 427629649267311777:
		return 41
	case 6477500296980050210:
		return 42
	case 4127240633310655048:
		return 43
	case 5792771328741774991:
		return 44
	case 5898328517505784191:
		return 45
	case 5217623514630858297:
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

func QComboBox_enable(h uintptr, idx int32) { pg_QComboBox__set_override(h, idx) }

func (this *QComboBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QComboBox) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QComboBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QComboBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QComboBox) OnShowPopup(fn func()) {
	registerVirt(this.UnsafeGoID(), -977583612921719482, func() { fn() }) // slot "showPopup"
}

func (this *QComboBox) OnHidePopup(fn func()) {
	registerVirt(this.UnsafeGoID(), 3308350117881191909, func() { fn() }) // slot "hidePopup"
}

func (this *QComboBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QComboBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QComboBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QComboBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QComboBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QComboBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QComboBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QComboBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(e uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(e))) }) // slot "showEvent"
}

func (this *QComboBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(e uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(e))) }) // slot "hideEvent"
}

func (this *QComboBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QComboBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QComboBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QComboBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QComboBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QComboBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(e uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(e))) }) // slot "contextMenuEvent"
}

func (this *QComboBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QComboBox) OnInitStyleOption(fn func(*QStyleOptionComboBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionComboBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QComboBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QComboBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QComboBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QComboBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QComboBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QComboBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QComboBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QComboBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QComboBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QComboBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QComboBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QComboBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QComboBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QComboBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QComboBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QComboBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QComboBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QComboBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QComboBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QComboBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QComboBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QComboBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QComboBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QComboBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QComboBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QComboBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QComboBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QComboBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QComboBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QComboBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QComboBox_delete(this.h)
	}
	this.UnsafeRelease()
}
