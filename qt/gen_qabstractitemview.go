package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractItemView_chain uintptr
	addr_QAbstractItemView_tr uintptr
	addr_QAbstractItemView_setModel uintptr
	addr_QAbstractItemView_model uintptr
	addr_QAbstractItemView_setSelectionModel uintptr
	addr_QAbstractItemView_selectionModel uintptr
	addr_QAbstractItemView_setItemDelegate uintptr
	addr_QAbstractItemView_itemDelegate uintptr
	addr_QAbstractItemView_setSelectionMode uintptr
	addr_QAbstractItemView_selectionMode uintptr
	addr_QAbstractItemView_setSelectionBehavior uintptr
	addr_QAbstractItemView_selectionBehavior uintptr
	addr_QAbstractItemView_currentIndex uintptr
	addr_QAbstractItemView_rootIndex uintptr
	addr_QAbstractItemView_setEditTriggers uintptr
	addr_QAbstractItemView_editTriggers uintptr
	addr_QAbstractItemView_setVerticalScrollMode uintptr
	addr_QAbstractItemView_verticalScrollMode uintptr
	addr_QAbstractItemView_resetVerticalScrollMode uintptr
	addr_QAbstractItemView_setHorizontalScrollMode uintptr
	addr_QAbstractItemView_horizontalScrollMode uintptr
	addr_QAbstractItemView_resetHorizontalScrollMode uintptr
	addr_QAbstractItemView_setAutoScroll uintptr
	addr_QAbstractItemView_hasAutoScroll uintptr
	addr_QAbstractItemView_setAutoScrollMargin uintptr
	addr_QAbstractItemView_autoScrollMargin uintptr
	addr_QAbstractItemView_setTabKeyNavigation uintptr
	addr_QAbstractItemView_tabKeyNavigation uintptr
	addr_QAbstractItemView_setDropIndicatorShown uintptr
	addr_QAbstractItemView_showDropIndicator uintptr
	addr_QAbstractItemView_setDragEnabled uintptr
	addr_QAbstractItemView_dragEnabled uintptr
	addr_QAbstractItemView_setDragDropOverwriteMode uintptr
	addr_QAbstractItemView_dragDropOverwriteMode uintptr
	addr_QAbstractItemView_setDragDropMode uintptr
	addr_QAbstractItemView_dragDropMode uintptr
	addr_QAbstractItemView_setDefaultDropAction uintptr
	addr_QAbstractItemView_defaultDropAction uintptr
	addr_QAbstractItemView_setAlternatingRowColors uintptr
	addr_QAbstractItemView_alternatingRowColors uintptr
	addr_QAbstractItemView_setIconSize uintptr
	addr_QAbstractItemView_iconSize uintptr
	addr_QAbstractItemView_setTextElideMode uintptr
	addr_QAbstractItemView_textElideMode uintptr
	addr_QAbstractItemView_keyboardSearch uintptr
	addr_QAbstractItemView_visualRect uintptr
	addr_QAbstractItemView_scrollTo uintptr
	addr_QAbstractItemView_indexAt uintptr
	addr_QAbstractItemView_sizeHintForIndex uintptr
	addr_QAbstractItemView_sizeHintForRow uintptr
	addr_QAbstractItemView_sizeHintForColumn uintptr
	addr_QAbstractItemView_updateThreshold uintptr
	addr_QAbstractItemView_setUpdateThreshold uintptr
	addr_QAbstractItemView_keyboardSearchFlags uintptr
	addr_QAbstractItemView_setKeyboardSearchFlags uintptr
	addr_QAbstractItemView_openPersistentEditor uintptr
	addr_QAbstractItemView_closePersistentEditor uintptr
	addr_QAbstractItemView_isPersistentEditorOpen uintptr
	addr_QAbstractItemView_setIndexWidget uintptr
	addr_QAbstractItemView_indexWidget uintptr
	addr_QAbstractItemView_setItemDelegateForRow uintptr
	addr_QAbstractItemView_itemDelegateForRow uintptr
	addr_QAbstractItemView_setItemDelegateForColumn uintptr
	addr_QAbstractItemView_itemDelegateForColumn uintptr
	addr_QAbstractItemView_itemDelegate2 uintptr
	addr_QAbstractItemView_itemDelegateForIndex uintptr
	addr_QAbstractItemView_reset uintptr
	addr_QAbstractItemView_setRootIndex uintptr
	addr_QAbstractItemView_doItemsLayout uintptr
	addr_QAbstractItemView_selectAll uintptr
	addr_QAbstractItemView_edit uintptr
	addr_QAbstractItemView_clearSelection uintptr
	addr_QAbstractItemView_setCurrentIndex uintptr
	addr_QAbstractItemView_scrollToTop uintptr
	addr_QAbstractItemView_scrollToBottom uintptr
	addr_QAbstractItemView_update uintptr
	addr_QAbstractItemView_pressed uintptr
	addr_QAbstractItemView_clicked uintptr
	addr_QAbstractItemView_doubleClicked uintptr
	addr_QAbstractItemView_activated uintptr
	addr_QAbstractItemView_entered uintptr
	addr_QAbstractItemView_viewportEntered uintptr
	addr_QAbstractItemView_iconSizeChanged uintptr
	addr_QAbstractItemView_state uintptr
	addr_QAbstractItemView_scheduleDelayedItemsLayout uintptr
	addr_QAbstractItemView_executeDelayedItemsLayout uintptr
	addr_QAbstractItemView_setDirtyRegion uintptr
	addr_QAbstractItemView_scrollDirtyRegion uintptr
	addr_QAbstractItemView_dirtyRegionOffset uintptr
	addr_QAbstractItemView_startAutoScroll uintptr
	addr_QAbstractItemView_stopAutoScroll uintptr
	addr_QAbstractItemView_doAutoScroll uintptr
	addr_QAbstractItemView_dropIndicatorPosition uintptr
	addr_QAbstractItemView_setViewportMargins uintptr
	addr_QAbstractItemView_setViewportMargins2 uintptr
	addr_QAbstractItemView_viewportMargins uintptr
	addr_QAbstractItemView_drawFrame uintptr
	addr_QAbstractItemView_setTabOrder uintptr
	addr_QAbstractItemView_mouseGrabber uintptr
	addr_QAbstractItemView_keyboardGrabber uintptr
	addr_QAbstractItemView_find uintptr
	addr_QAbstractItemView_createWindowContainer uintptr
	addr_QAbstractItemView_updateMicroFocus uintptr
	addr_QAbstractItemView_create uintptr
	addr_QAbstractItemView_destroy uintptr
	addr_QAbstractItemView_focusNextChild uintptr
	addr_QAbstractItemView_focusPreviousChild uintptr
	addr_QAbstractItemView_connect uintptr
	addr_QAbstractItemView_connect2 uintptr
	addr_QAbstractItemView_sender uintptr
	addr_QAbstractItemView_senderSignalIndex uintptr
	addr_QAbstractItemView_receivers uintptr
	addr_QAbstractItemView_isSignalConnected uintptr
	pg_QAbstractItemView_devicePixelRatioFScale func() float64
	pg_QAbstractItemView_encodeMetricF func(int32, float64) int32
	pg_QAbstractItemView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QAbstractItemView_tr2 uintptr
	addr_QAbstractItemView_tr3 uintptr
	addr_QAbstractItemView_createWindowContainer2 uintptr
	addr_QAbstractItemView_createWindowContainer3 uintptr
	addr_QAbstractItemView_updateMicroFocus2 uintptr
	addr_QAbstractItemView_create2 uintptr
	addr_QAbstractItemView_create3 uintptr
	addr_QAbstractItemView_create4 uintptr
	addr_QAbstractItemView_destroy2 uintptr
	addr_QAbstractItemView_destroy3 uintptr
	addr_QAbstractItemView_connect3 uintptr
	addr_QAbstractItemView_connect4 uintptr
	addr_QAbstractItemView_setModel_default uintptr
	addr_QAbstractItemView_setSelectionModel_default uintptr
	addr_QAbstractItemView_keyboardSearch_default uintptr
	addr_QAbstractItemView_sizeHintForRow_default uintptr
	addr_QAbstractItemView_sizeHintForColumn_default uintptr
	addr_QAbstractItemView_itemDelegateForIndex_default uintptr
	addr_QAbstractItemView_reset_default uintptr
	addr_QAbstractItemView_setRootIndex_default uintptr
	addr_QAbstractItemView_doItemsLayout_default uintptr
	addr_QAbstractItemView_selectAll_default uintptr
	addr_QAbstractItemView_edit_default uintptr
	addr_QAbstractItemView_metaObject_default uintptr
	addr_QAbstractItemView_inputMethodQuery_default uintptr
	addr_QAbstractItemView_rowsInserted_default uintptr
	addr_QAbstractItemView_rowsAboutToBeRemoved_default uintptr
	addr_QAbstractItemView_selectionChanged_default uintptr
	addr_QAbstractItemView_currentChanged_default uintptr
	addr_QAbstractItemView_updateEditorData_default uintptr
	addr_QAbstractItemView_updateEditorGeometries_default uintptr
	addr_QAbstractItemView_updateGeometries_default uintptr
	addr_QAbstractItemView_verticalScrollbarAction_default uintptr
	addr_QAbstractItemView_horizontalScrollbarAction_default uintptr
	addr_QAbstractItemView_verticalScrollbarValueChanged_default uintptr
	addr_QAbstractItemView_horizontalScrollbarValueChanged_default uintptr
	addr_QAbstractItemView_closeEditor_default uintptr
	addr_QAbstractItemView_commitData_default uintptr
	addr_QAbstractItemView_editorDestroyed_default uintptr
	addr_QAbstractItemView_selectedIndexes_default uintptr
	addr_QAbstractItemView_selectionCommand_default uintptr
	addr_QAbstractItemView_startDrag_default uintptr
	addr_QAbstractItemView_initViewItemOption_default uintptr
	addr_QAbstractItemView_focusNextPrevChild_default uintptr
	addr_QAbstractItemView_event_default uintptr
	addr_QAbstractItemView_viewportEvent_default uintptr
	addr_QAbstractItemView_mousePressEvent_default uintptr
	addr_QAbstractItemView_mouseMoveEvent_default uintptr
	addr_QAbstractItemView_mouseReleaseEvent_default uintptr
	addr_QAbstractItemView_mouseDoubleClickEvent_default uintptr
	addr_QAbstractItemView_dragEnterEvent_default uintptr
	addr_QAbstractItemView_dragMoveEvent_default uintptr
	addr_QAbstractItemView_dragLeaveEvent_default uintptr
	addr_QAbstractItemView_dropEvent_default uintptr
	addr_QAbstractItemView_focusInEvent_default uintptr
	addr_QAbstractItemView_focusOutEvent_default uintptr
	addr_QAbstractItemView_keyPressEvent_default uintptr
	addr_QAbstractItemView_resizeEvent_default uintptr
	addr_QAbstractItemView_timerEvent_default uintptr
	addr_QAbstractItemView_inputMethodEvent_default uintptr
	addr_QAbstractItemView_eventFilter_default uintptr
	addr_QAbstractItemView_viewportSizeHint_default uintptr
	addr_QAbstractItemView_minimumSizeHint_default uintptr
	addr_QAbstractItemView_sizeHint_default uintptr
	addr_QAbstractItemView_setupViewport_default uintptr
	addr_QAbstractItemView_paintEvent_default uintptr
	addr_QAbstractItemView_wheelEvent_default uintptr
	addr_QAbstractItemView_contextMenuEvent_default uintptr
	addr_QAbstractItemView_scrollContentsBy_default uintptr
	addr_QAbstractItemView_changeEvent_default uintptr
	addr_QAbstractItemView_initStyleOption_default uintptr
	addr_QAbstractItemView_devType_default uintptr
	addr_QAbstractItemView_setVisible_default uintptr
	addr_QAbstractItemView_heightForWidth_default uintptr
	addr_QAbstractItemView_hasHeightForWidth_default uintptr
	addr_QAbstractItemView_paintEngine_default uintptr
	addr_QAbstractItemView_keyReleaseEvent_default uintptr
	addr_QAbstractItemView_enterEvent_default uintptr
	addr_QAbstractItemView_leaveEvent_default uintptr
	addr_QAbstractItemView_moveEvent_default uintptr
	addr_QAbstractItemView_closeEvent_default uintptr
	addr_QAbstractItemView_tabletEvent_default uintptr
	addr_QAbstractItemView_actionEvent_default uintptr
	addr_QAbstractItemView_showEvent_default uintptr
	addr_QAbstractItemView_hideEvent_default uintptr
	addr_QAbstractItemView_metric_default uintptr
	addr_QAbstractItemView_initPainter_default uintptr
	addr_QAbstractItemView_redirected_default uintptr
	addr_QAbstractItemView_sharedPainter_default uintptr
	addr_QAbstractItemView_childEvent_default uintptr
	addr_QAbstractItemView_customEvent_default uintptr
	addr_QAbstractItemView_connectNotify_default uintptr
	addr_QAbstractItemView_disconnectNotify_default uintptr
	pg_QAbstractItemView_connect_pressed func(uintptr, int64)
	pg_QAbstractItemView_connect_clicked func(uintptr, int64)
	pg_QAbstractItemView_connect_doubleClicked func(uintptr, int64)
	pg_QAbstractItemView_connect_activated func(uintptr, int64)
	pg_QAbstractItemView_connect_entered func(uintptr, int64)
	pg_QAbstractItemView_connect_viewportEntered func(uintptr, int64)
	pg_QAbstractItemView_connect_iconSizeChanged func(uintptr, int64)
	pg_QAbstractItemView_connect_windowTitleChanged func(uintptr, int64)
	pg_QAbstractItemView_connect_windowIconChanged func(uintptr, int64)
	pg_QAbstractItemView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QAbstractItemView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QAbstractItemView_connect_destroyed func(uintptr, int64)
	pg_QAbstractItemView_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractItemView_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractItemView__disconnect func(int64)
	pg_QAbstractItemView__set_override func(uintptr, int32)
	pg_QAbstractItemView__Go_new func(int64) uintptr
	pg_QAbstractItemView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractItemView_chain, "QAbstractItemView_chain")
	symaddr(&addr_QAbstractItemView_tr, "QAbstractItemView_tr")
	symaddr(&addr_QAbstractItemView_setModel, "QAbstractItemView_setModel")
	symaddr(&addr_QAbstractItemView_model, "QAbstractItemView_model")
	symaddr(&addr_QAbstractItemView_setSelectionModel, "QAbstractItemView_setSelectionModel")
	symaddr(&addr_QAbstractItemView_selectionModel, "QAbstractItemView_selectionModel")
	symaddr(&addr_QAbstractItemView_setItemDelegate, "QAbstractItemView_setItemDelegate")
	symaddr(&addr_QAbstractItemView_itemDelegate, "QAbstractItemView_itemDelegate")
	symaddr(&addr_QAbstractItemView_setSelectionMode, "QAbstractItemView_setSelectionMode")
	symaddr(&addr_QAbstractItemView_selectionMode, "QAbstractItemView_selectionMode")
	symaddr(&addr_QAbstractItemView_setSelectionBehavior, "QAbstractItemView_setSelectionBehavior")
	symaddr(&addr_QAbstractItemView_selectionBehavior, "QAbstractItemView_selectionBehavior")
	symaddr(&addr_QAbstractItemView_currentIndex, "QAbstractItemView_currentIndex")
	symaddr(&addr_QAbstractItemView_rootIndex, "QAbstractItemView_rootIndex")
	symaddr(&addr_QAbstractItemView_setEditTriggers, "QAbstractItemView_setEditTriggers")
	symaddr(&addr_QAbstractItemView_editTriggers, "QAbstractItemView_editTriggers")
	symaddr(&addr_QAbstractItemView_setVerticalScrollMode, "QAbstractItemView_setVerticalScrollMode")
	symaddr(&addr_QAbstractItemView_verticalScrollMode, "QAbstractItemView_verticalScrollMode")
	symaddr(&addr_QAbstractItemView_resetVerticalScrollMode, "QAbstractItemView_resetVerticalScrollMode")
	symaddr(&addr_QAbstractItemView_setHorizontalScrollMode, "QAbstractItemView_setHorizontalScrollMode")
	symaddr(&addr_QAbstractItemView_horizontalScrollMode, "QAbstractItemView_horizontalScrollMode")
	symaddr(&addr_QAbstractItemView_resetHorizontalScrollMode, "QAbstractItemView_resetHorizontalScrollMode")
	symaddr(&addr_QAbstractItemView_setAutoScroll, "QAbstractItemView_setAutoScroll")
	symaddr(&addr_QAbstractItemView_hasAutoScroll, "QAbstractItemView_hasAutoScroll")
	symaddr(&addr_QAbstractItemView_setAutoScrollMargin, "QAbstractItemView_setAutoScrollMargin")
	symaddr(&addr_QAbstractItemView_autoScrollMargin, "QAbstractItemView_autoScrollMargin")
	symaddr(&addr_QAbstractItemView_setTabKeyNavigation, "QAbstractItemView_setTabKeyNavigation")
	symaddr(&addr_QAbstractItemView_tabKeyNavigation, "QAbstractItemView_tabKeyNavigation")
	symaddr(&addr_QAbstractItemView_setDropIndicatorShown, "QAbstractItemView_setDropIndicatorShown")
	symaddr(&addr_QAbstractItemView_showDropIndicator, "QAbstractItemView_showDropIndicator")
	symaddr(&addr_QAbstractItemView_setDragEnabled, "QAbstractItemView_setDragEnabled")
	symaddr(&addr_QAbstractItemView_dragEnabled, "QAbstractItemView_dragEnabled")
	symaddr(&addr_QAbstractItemView_setDragDropOverwriteMode, "QAbstractItemView_setDragDropOverwriteMode")
	symaddr(&addr_QAbstractItemView_dragDropOverwriteMode, "QAbstractItemView_dragDropOverwriteMode")
	symaddr(&addr_QAbstractItemView_setDragDropMode, "QAbstractItemView_setDragDropMode")
	symaddr(&addr_QAbstractItemView_dragDropMode, "QAbstractItemView_dragDropMode")
	symaddr(&addr_QAbstractItemView_setDefaultDropAction, "QAbstractItemView_setDefaultDropAction")
	symaddr(&addr_QAbstractItemView_defaultDropAction, "QAbstractItemView_defaultDropAction")
	symaddr(&addr_QAbstractItemView_setAlternatingRowColors, "QAbstractItemView_setAlternatingRowColors")
	symaddr(&addr_QAbstractItemView_alternatingRowColors, "QAbstractItemView_alternatingRowColors")
	symaddr(&addr_QAbstractItemView_setIconSize, "QAbstractItemView_setIconSize")
	symaddr(&addr_QAbstractItemView_iconSize, "QAbstractItemView_iconSize")
	symaddr(&addr_QAbstractItemView_setTextElideMode, "QAbstractItemView_setTextElideMode")
	symaddr(&addr_QAbstractItemView_textElideMode, "QAbstractItemView_textElideMode")
	symaddr(&addr_QAbstractItemView_keyboardSearch, "QAbstractItemView_keyboardSearch")
	symaddr(&addr_QAbstractItemView_visualRect, "QAbstractItemView_visualRect")
	symaddr(&addr_QAbstractItemView_scrollTo, "QAbstractItemView_scrollTo")
	symaddr(&addr_QAbstractItemView_indexAt, "QAbstractItemView_indexAt")
	symaddr(&addr_QAbstractItemView_sizeHintForIndex, "QAbstractItemView_sizeHintForIndex")
	symaddr(&addr_QAbstractItemView_sizeHintForRow, "QAbstractItemView_sizeHintForRow")
	symaddr(&addr_QAbstractItemView_sizeHintForColumn, "QAbstractItemView_sizeHintForColumn")
	symaddr(&addr_QAbstractItemView_updateThreshold, "QAbstractItemView_updateThreshold")
	symaddr(&addr_QAbstractItemView_setUpdateThreshold, "QAbstractItemView_setUpdateThreshold")
	symaddr(&addr_QAbstractItemView_keyboardSearchFlags, "QAbstractItemView_keyboardSearchFlags")
	symaddr(&addr_QAbstractItemView_setKeyboardSearchFlags, "QAbstractItemView_setKeyboardSearchFlags")
	symaddr(&addr_QAbstractItemView_openPersistentEditor, "QAbstractItemView_openPersistentEditor")
	symaddr(&addr_QAbstractItemView_closePersistentEditor, "QAbstractItemView_closePersistentEditor")
	symaddr(&addr_QAbstractItemView_isPersistentEditorOpen, "QAbstractItemView_isPersistentEditorOpen")
	symaddr(&addr_QAbstractItemView_setIndexWidget, "QAbstractItemView_setIndexWidget")
	symaddr(&addr_QAbstractItemView_indexWidget, "QAbstractItemView_indexWidget")
	symaddr(&addr_QAbstractItemView_setItemDelegateForRow, "QAbstractItemView_setItemDelegateForRow")
	symaddr(&addr_QAbstractItemView_itemDelegateForRow, "QAbstractItemView_itemDelegateForRow")
	symaddr(&addr_QAbstractItemView_setItemDelegateForColumn, "QAbstractItemView_setItemDelegateForColumn")
	symaddr(&addr_QAbstractItemView_itemDelegateForColumn, "QAbstractItemView_itemDelegateForColumn")
	symaddr(&addr_QAbstractItemView_itemDelegate2, "QAbstractItemView_itemDelegate2")
	symaddr(&addr_QAbstractItemView_itemDelegateForIndex, "QAbstractItemView_itemDelegateForIndex")
	symaddr(&addr_QAbstractItemView_reset, "QAbstractItemView_reset")
	symaddr(&addr_QAbstractItemView_setRootIndex, "QAbstractItemView_setRootIndex")
	symaddr(&addr_QAbstractItemView_doItemsLayout, "QAbstractItemView_doItemsLayout")
	symaddr(&addr_QAbstractItemView_selectAll, "QAbstractItemView_selectAll")
	symaddr(&addr_QAbstractItemView_edit, "QAbstractItemView_edit")
	symaddr(&addr_QAbstractItemView_clearSelection, "QAbstractItemView_clearSelection")
	symaddr(&addr_QAbstractItemView_setCurrentIndex, "QAbstractItemView_setCurrentIndex")
	symaddr(&addr_QAbstractItemView_scrollToTop, "QAbstractItemView_scrollToTop")
	symaddr(&addr_QAbstractItemView_scrollToBottom, "QAbstractItemView_scrollToBottom")
	symaddr(&addr_QAbstractItemView_update, "QAbstractItemView_update")
	symaddr(&addr_QAbstractItemView_pressed, "QAbstractItemView_pressed")
	symaddr(&addr_QAbstractItemView_clicked, "QAbstractItemView_clicked")
	symaddr(&addr_QAbstractItemView_doubleClicked, "QAbstractItemView_doubleClicked")
	symaddr(&addr_QAbstractItemView_activated, "QAbstractItemView_activated")
	symaddr(&addr_QAbstractItemView_entered, "QAbstractItemView_entered")
	symaddr(&addr_QAbstractItemView_viewportEntered, "QAbstractItemView_viewportEntered")
	symaddr(&addr_QAbstractItemView_iconSizeChanged, "QAbstractItemView_iconSizeChanged")
	symaddr(&addr_QAbstractItemView_state, "QAbstractItemView_state")
	symaddr(&addr_QAbstractItemView_scheduleDelayedItemsLayout, "QAbstractItemView_scheduleDelayedItemsLayout")
	symaddr(&addr_QAbstractItemView_executeDelayedItemsLayout, "QAbstractItemView_executeDelayedItemsLayout")
	symaddr(&addr_QAbstractItemView_setDirtyRegion, "QAbstractItemView_setDirtyRegion")
	symaddr(&addr_QAbstractItemView_scrollDirtyRegion, "QAbstractItemView_scrollDirtyRegion")
	symaddr(&addr_QAbstractItemView_dirtyRegionOffset, "QAbstractItemView_dirtyRegionOffset")
	symaddr(&addr_QAbstractItemView_startAutoScroll, "QAbstractItemView_startAutoScroll")
	symaddr(&addr_QAbstractItemView_stopAutoScroll, "QAbstractItemView_stopAutoScroll")
	symaddr(&addr_QAbstractItemView_doAutoScroll, "QAbstractItemView_doAutoScroll")
	symaddr(&addr_QAbstractItemView_dropIndicatorPosition, "QAbstractItemView_dropIndicatorPosition")
	symaddr(&addr_QAbstractItemView_setViewportMargins, "QAbstractItemView_setViewportMargins")
	symaddr(&addr_QAbstractItemView_setViewportMargins2, "QAbstractItemView_setViewportMargins2")
	symaddr(&addr_QAbstractItemView_viewportMargins, "QAbstractItemView_viewportMargins")
	symaddr(&addr_QAbstractItemView_drawFrame, "QAbstractItemView_drawFrame")
	symaddr(&addr_QAbstractItemView_setTabOrder, "QAbstractItemView_setTabOrder")
	symaddr(&addr_QAbstractItemView_mouseGrabber, "QAbstractItemView_mouseGrabber")
	symaddr(&addr_QAbstractItemView_keyboardGrabber, "QAbstractItemView_keyboardGrabber")
	symaddr(&addr_QAbstractItemView_find, "QAbstractItemView_find")
	symaddr(&addr_QAbstractItemView_createWindowContainer, "QAbstractItemView_createWindowContainer")
	symaddr(&addr_QAbstractItemView_updateMicroFocus, "QAbstractItemView_updateMicroFocus")
	symaddr(&addr_QAbstractItemView_create, "QAbstractItemView_create")
	symaddr(&addr_QAbstractItemView_destroy, "QAbstractItemView_destroy")
	symaddr(&addr_QAbstractItemView_focusNextChild, "QAbstractItemView_focusNextChild")
	symaddr(&addr_QAbstractItemView_focusPreviousChild, "QAbstractItemView_focusPreviousChild")
	symaddr(&addr_QAbstractItemView_connect, "QAbstractItemView_connect")
	symaddr(&addr_QAbstractItemView_connect2, "QAbstractItemView_connect2")
	symaddr(&addr_QAbstractItemView_sender, "QAbstractItemView_sender")
	symaddr(&addr_QAbstractItemView_senderSignalIndex, "QAbstractItemView_senderSignalIndex")
	symaddr(&addr_QAbstractItemView_receivers, "QAbstractItemView_receivers")
	symaddr(&addr_QAbstractItemView_isSignalConnected, "QAbstractItemView_isSignalConnected")
	reg(&pg_QAbstractItemView_devicePixelRatioFScale, "QAbstractItemView_devicePixelRatioFScale")
	reg(&pg_QAbstractItemView_encodeMetricF, "QAbstractItemView_encodeMetricF")
	reg(&pg_QAbstractItemView_getDecodedMetricF, "QAbstractItemView_getDecodedMetricF")
	symaddr(&addr_QAbstractItemView_tr2, "QAbstractItemView_tr2")
	symaddr(&addr_QAbstractItemView_tr3, "QAbstractItemView_tr3")
	symaddr(&addr_QAbstractItemView_createWindowContainer2, "QAbstractItemView_createWindowContainer2")
	symaddr(&addr_QAbstractItemView_createWindowContainer3, "QAbstractItemView_createWindowContainer3")
	symaddr(&addr_QAbstractItemView_updateMicroFocus2, "QAbstractItemView_updateMicroFocus2")
	symaddr(&addr_QAbstractItemView_create2, "QAbstractItemView_create2")
	symaddr(&addr_QAbstractItemView_create3, "QAbstractItemView_create3")
	symaddr(&addr_QAbstractItemView_create4, "QAbstractItemView_create4")
	symaddr(&addr_QAbstractItemView_destroy2, "QAbstractItemView_destroy2")
	symaddr(&addr_QAbstractItemView_destroy3, "QAbstractItemView_destroy3")
	symaddr(&addr_QAbstractItemView_connect3, "QAbstractItemView_connect3")
	symaddr(&addr_QAbstractItemView_connect4, "QAbstractItemView_connect4")
	symaddr(&addr_QAbstractItemView_setModel_default, "QAbstractItemView_setModel_default")
	symaddr(&addr_QAbstractItemView_setSelectionModel_default, "QAbstractItemView_setSelectionModel_default")
	symaddr(&addr_QAbstractItemView_keyboardSearch_default, "QAbstractItemView_keyboardSearch_default")
	symaddr(&addr_QAbstractItemView_sizeHintForRow_default, "QAbstractItemView_sizeHintForRow_default")
	symaddr(&addr_QAbstractItemView_sizeHintForColumn_default, "QAbstractItemView_sizeHintForColumn_default")
	symaddr(&addr_QAbstractItemView_itemDelegateForIndex_default, "QAbstractItemView_itemDelegateForIndex_default")
	symaddr(&addr_QAbstractItemView_reset_default, "QAbstractItemView_reset_default")
	symaddr(&addr_QAbstractItemView_setRootIndex_default, "QAbstractItemView_setRootIndex_default")
	symaddr(&addr_QAbstractItemView_doItemsLayout_default, "QAbstractItemView_doItemsLayout_default")
	symaddr(&addr_QAbstractItemView_selectAll_default, "QAbstractItemView_selectAll_default")
	symaddr(&addr_QAbstractItemView_edit_default, "QAbstractItemView_edit_default")
	symaddr(&addr_QAbstractItemView_metaObject_default, "QAbstractItemView_metaObject_default")
	symaddr(&addr_QAbstractItemView_inputMethodQuery_default, "QAbstractItemView_inputMethodQuery_default")
	symaddr(&addr_QAbstractItemView_rowsInserted_default, "QAbstractItemView_rowsInserted_default")
	symaddr(&addr_QAbstractItemView_rowsAboutToBeRemoved_default, "QAbstractItemView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QAbstractItemView_selectionChanged_default, "QAbstractItemView_selectionChanged_default")
	symaddr(&addr_QAbstractItemView_currentChanged_default, "QAbstractItemView_currentChanged_default")
	symaddr(&addr_QAbstractItemView_updateEditorData_default, "QAbstractItemView_updateEditorData_default")
	symaddr(&addr_QAbstractItemView_updateEditorGeometries_default, "QAbstractItemView_updateEditorGeometries_default")
	symaddr(&addr_QAbstractItemView_updateGeometries_default, "QAbstractItemView_updateGeometries_default")
	symaddr(&addr_QAbstractItemView_verticalScrollbarAction_default, "QAbstractItemView_verticalScrollbarAction_default")
	symaddr(&addr_QAbstractItemView_horizontalScrollbarAction_default, "QAbstractItemView_horizontalScrollbarAction_default")
	symaddr(&addr_QAbstractItemView_verticalScrollbarValueChanged_default, "QAbstractItemView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QAbstractItemView_horizontalScrollbarValueChanged_default, "QAbstractItemView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QAbstractItemView_closeEditor_default, "QAbstractItemView_closeEditor_default")
	symaddr(&addr_QAbstractItemView_commitData_default, "QAbstractItemView_commitData_default")
	symaddr(&addr_QAbstractItemView_editorDestroyed_default, "QAbstractItemView_editorDestroyed_default")
	symaddr(&addr_QAbstractItemView_selectedIndexes_default, "QAbstractItemView_selectedIndexes_default")
	symaddr(&addr_QAbstractItemView_selectionCommand_default, "QAbstractItemView_selectionCommand_default")
	symaddr(&addr_QAbstractItemView_startDrag_default, "QAbstractItemView_startDrag_default")
	symaddr(&addr_QAbstractItemView_initViewItemOption_default, "QAbstractItemView_initViewItemOption_default")
	symaddr(&addr_QAbstractItemView_focusNextPrevChild_default, "QAbstractItemView_focusNextPrevChild_default")
	symaddr(&addr_QAbstractItemView_event_default, "QAbstractItemView_event_default")
	symaddr(&addr_QAbstractItemView_viewportEvent_default, "QAbstractItemView_viewportEvent_default")
	symaddr(&addr_QAbstractItemView_mousePressEvent_default, "QAbstractItemView_mousePressEvent_default")
	symaddr(&addr_QAbstractItemView_mouseMoveEvent_default, "QAbstractItemView_mouseMoveEvent_default")
	symaddr(&addr_QAbstractItemView_mouseReleaseEvent_default, "QAbstractItemView_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractItemView_mouseDoubleClickEvent_default, "QAbstractItemView_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractItemView_dragEnterEvent_default, "QAbstractItemView_dragEnterEvent_default")
	symaddr(&addr_QAbstractItemView_dragMoveEvent_default, "QAbstractItemView_dragMoveEvent_default")
	symaddr(&addr_QAbstractItemView_dragLeaveEvent_default, "QAbstractItemView_dragLeaveEvent_default")
	symaddr(&addr_QAbstractItemView_dropEvent_default, "QAbstractItemView_dropEvent_default")
	symaddr(&addr_QAbstractItemView_focusInEvent_default, "QAbstractItemView_focusInEvent_default")
	symaddr(&addr_QAbstractItemView_focusOutEvent_default, "QAbstractItemView_focusOutEvent_default")
	symaddr(&addr_QAbstractItemView_keyPressEvent_default, "QAbstractItemView_keyPressEvent_default")
	symaddr(&addr_QAbstractItemView_resizeEvent_default, "QAbstractItemView_resizeEvent_default")
	symaddr(&addr_QAbstractItemView_timerEvent_default, "QAbstractItemView_timerEvent_default")
	symaddr(&addr_QAbstractItemView_inputMethodEvent_default, "QAbstractItemView_inputMethodEvent_default")
	symaddr(&addr_QAbstractItemView_eventFilter_default, "QAbstractItemView_eventFilter_default")
	symaddr(&addr_QAbstractItemView_viewportSizeHint_default, "QAbstractItemView_viewportSizeHint_default")
	symaddr(&addr_QAbstractItemView_minimumSizeHint_default, "QAbstractItemView_minimumSizeHint_default")
	symaddr(&addr_QAbstractItemView_sizeHint_default, "QAbstractItemView_sizeHint_default")
	symaddr(&addr_QAbstractItemView_setupViewport_default, "QAbstractItemView_setupViewport_default")
	symaddr(&addr_QAbstractItemView_paintEvent_default, "QAbstractItemView_paintEvent_default")
	symaddr(&addr_QAbstractItemView_wheelEvent_default, "QAbstractItemView_wheelEvent_default")
	symaddr(&addr_QAbstractItemView_contextMenuEvent_default, "QAbstractItemView_contextMenuEvent_default")
	symaddr(&addr_QAbstractItemView_scrollContentsBy_default, "QAbstractItemView_scrollContentsBy_default")
	symaddr(&addr_QAbstractItemView_changeEvent_default, "QAbstractItemView_changeEvent_default")
	symaddr(&addr_QAbstractItemView_initStyleOption_default, "QAbstractItemView_initStyleOption_default")
	symaddr(&addr_QAbstractItemView_devType_default, "QAbstractItemView_devType_default")
	symaddr(&addr_QAbstractItemView_setVisible_default, "QAbstractItemView_setVisible_default")
	symaddr(&addr_QAbstractItemView_heightForWidth_default, "QAbstractItemView_heightForWidth_default")
	symaddr(&addr_QAbstractItemView_hasHeightForWidth_default, "QAbstractItemView_hasHeightForWidth_default")
	symaddr(&addr_QAbstractItemView_paintEngine_default, "QAbstractItemView_paintEngine_default")
	symaddr(&addr_QAbstractItemView_keyReleaseEvent_default, "QAbstractItemView_keyReleaseEvent_default")
	symaddr(&addr_QAbstractItemView_enterEvent_default, "QAbstractItemView_enterEvent_default")
	symaddr(&addr_QAbstractItemView_leaveEvent_default, "QAbstractItemView_leaveEvent_default")
	symaddr(&addr_QAbstractItemView_moveEvent_default, "QAbstractItemView_moveEvent_default")
	symaddr(&addr_QAbstractItemView_closeEvent_default, "QAbstractItemView_closeEvent_default")
	symaddr(&addr_QAbstractItemView_tabletEvent_default, "QAbstractItemView_tabletEvent_default")
	symaddr(&addr_QAbstractItemView_actionEvent_default, "QAbstractItemView_actionEvent_default")
	symaddr(&addr_QAbstractItemView_showEvent_default, "QAbstractItemView_showEvent_default")
	symaddr(&addr_QAbstractItemView_hideEvent_default, "QAbstractItemView_hideEvent_default")
	symaddr(&addr_QAbstractItemView_metric_default, "QAbstractItemView_metric_default")
	symaddr(&addr_QAbstractItemView_initPainter_default, "QAbstractItemView_initPainter_default")
	symaddr(&addr_QAbstractItemView_redirected_default, "QAbstractItemView_redirected_default")
	symaddr(&addr_QAbstractItemView_sharedPainter_default, "QAbstractItemView_sharedPainter_default")
	symaddr(&addr_QAbstractItemView_childEvent_default, "QAbstractItemView_childEvent_default")
	symaddr(&addr_QAbstractItemView_customEvent_default, "QAbstractItemView_customEvent_default")
	symaddr(&addr_QAbstractItemView_connectNotify_default, "QAbstractItemView_connectNotify_default")
	symaddr(&addr_QAbstractItemView_disconnectNotify_default, "QAbstractItemView_disconnectNotify_default")
	reg(&pg_QAbstractItemView_connect_pressed, "QAbstractItemView_connect_pressed")
	reg(&pg_QAbstractItemView_connect_clicked, "QAbstractItemView_connect_clicked")
	reg(&pg_QAbstractItemView_connect_doubleClicked, "QAbstractItemView_connect_doubleClicked")
	reg(&pg_QAbstractItemView_connect_activated, "QAbstractItemView_connect_activated")
	reg(&pg_QAbstractItemView_connect_entered, "QAbstractItemView_connect_entered")
	reg(&pg_QAbstractItemView_connect_viewportEntered, "QAbstractItemView_connect_viewportEntered")
	reg(&pg_QAbstractItemView_connect_iconSizeChanged, "QAbstractItemView_connect_iconSizeChanged")
	reg(&pg_QAbstractItemView_connect_windowTitleChanged, "QAbstractItemView_connect_windowTitleChanged")
	reg(&pg_QAbstractItemView_connect_windowIconChanged, "QAbstractItemView_connect_windowIconChanged")
	reg(&pg_QAbstractItemView_connect_windowIconTextChanged, "QAbstractItemView_connect_windowIconTextChanged")
	reg(&pg_QAbstractItemView_connect_customContextMenuRequested, "QAbstractItemView_connect_customContextMenuRequested")
	reg(&pg_QAbstractItemView_connect_destroyed, "QAbstractItemView_connect_destroyed")
	reg(&pg_QAbstractItemView_connect_destroyed2, "QAbstractItemView_connect_destroyed2")
	reg(&pg_QAbstractItemView_connect_objectNameChanged, "QAbstractItemView_connect_objectNameChanged")
	reg(&pg_QAbstractItemView__disconnect, "QAbstractItemView__disconnect")
	reg(&pg_QAbstractItemView__set_override, "QAbstractItemView__set_override")
	reg(&pg_QAbstractItemView__Go_new, "QAbstractItemView__Go_new")
	reg(&pg_QAbstractItemView_delete, "QAbstractItemView_delete")
}

type QAbstractItemView struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQAbstractItemView(h uintptr) *QAbstractItemView {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QAbstractItemView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractItemView{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQAbstractItemView(h uintptr) *QAbstractItemView { return bindQAbstractItemView(h) }

func (this *QAbstractItemView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractItemView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractItemView(h unsafe.Pointer) *QAbstractItemView { return bindQAbstractItemView(uintptr(h)) }
func UnsafeNewQAbstractItemViewOwned(h unsafe.Pointer) *QAbstractItemView { return newownedQAbstractItemView(uintptr(h)) }

func (this *QAbstractItemView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func QAbstractItemView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractItemView) SetModel(model *QAbstractItemModel) {
	purego.SyscallN(addr_QAbstractItemView_setModel, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QAbstractItemView) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QAbstractItemView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QAbstractItemView_setSelectionModel, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QAbstractItemView) SelectionModel() *QItemSelectionModel {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_selectionModel, this.h)
	runtime.KeepAlive(this)
	return bindQItemSelectionModel(__r)
}

func (this *QAbstractItemView) SetItemDelegate(delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QAbstractItemView_setItemDelegate, this.h, uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegate, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	purego.SyscallN(addr_QAbstractItemView_setSelectionMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_selectionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__SelectionMode(__r)
}

func (this *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	purego.SyscallN(addr_QAbstractItemView_setSelectionBehavior, this.h, uintptr(behavior))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_selectionBehavior, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__SelectionBehavior(__r)
}

func (this *QAbstractItemView) CurrentIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_currentIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemView) RootIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_rootIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemView) SetEditTriggers(triggers QAbstractItemView__EditTrigger) {
	purego.SyscallN(addr_QAbstractItemView_setEditTriggers, this.h, uintptr(triggers))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) EditTriggers() QAbstractItemView__EditTrigger {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_editTriggers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__EditTrigger(__r)
}

func (this *QAbstractItemView) SetVerticalScrollMode(mode QAbstractItemView__ScrollMode) {
	purego.SyscallN(addr_QAbstractItemView_setVerticalScrollMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) VerticalScrollMode() QAbstractItemView__ScrollMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_verticalScrollMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__ScrollMode(__r)
}

func (this *QAbstractItemView) ResetVerticalScrollMode() {
	purego.SyscallN(addr_QAbstractItemView_resetVerticalScrollMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetHorizontalScrollMode(mode QAbstractItemView__ScrollMode) {
	purego.SyscallN(addr_QAbstractItemView_setHorizontalScrollMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) HorizontalScrollMode() QAbstractItemView__ScrollMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_horizontalScrollMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__ScrollMode(__r)
}

func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	purego.SyscallN(addr_QAbstractItemView_resetHorizontalScrollMode, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	purego.SyscallN(addr_QAbstractItemView_setAutoScroll, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) HasAutoScroll() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_hasAutoScroll, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	purego.SyscallN(addr_QAbstractItemView_setAutoScrollMargin, this.h, uintptr(margin))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) AutoScrollMargin() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_autoScrollMargin, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	purego.SyscallN(addr_QAbstractItemView_setTabKeyNavigation, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) TabKeyNavigation() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_tabKeyNavigation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	purego.SyscallN(addr_QAbstractItemView_setDropIndicatorShown, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) ShowDropIndicator() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_showDropIndicator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	purego.SyscallN(addr_QAbstractItemView_setDragEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DragEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_dragEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	purego.SyscallN(addr_QAbstractItemView_setDragDropOverwriteMode, this.h, uintptr(b2i(overwrite)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_dragDropOverwriteMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetDragDropMode(behavior QAbstractItemView__DragDropMode) {
	purego.SyscallN(addr_QAbstractItemView_setDragDropMode, this.h, uintptr(behavior))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DragDropMode() QAbstractItemView__DragDropMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_dragDropMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DragDropMode(__r)
}

func (this *QAbstractItemView) SetDefaultDropAction(dropAction DropAction) {
	purego.SyscallN(addr_QAbstractItemView_setDefaultDropAction, this.h, uintptr(dropAction))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DefaultDropAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_defaultDropAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	purego.SyscallN(addr_QAbstractItemView_setAlternatingRowColors, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) AlternatingRowColors() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_alternatingRowColors, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) SetIconSize(size *QSize) {
	purego.SyscallN(addr_QAbstractItemView_setIconSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QAbstractItemView) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractItemView) SetTextElideMode(mode TextElideMode) {
	purego.SyscallN(addr_QAbstractItemView_setTextElideMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) TextElideMode() TextElideMode {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_textElideMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextElideMode(__r)
}

func (this *QAbstractItemView) KeyboardSearch(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QAbstractItemView_keyboardSearch, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QAbstractItemView) VisualRect(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_visualRect, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QAbstractItemView) ScrollTo(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QAbstractItemView_scrollTo, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) IndexAt(point *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_indexAt, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemView) SizeHintForIndex(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHintForIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractItemView) SizeHintForRow(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHintForRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) SizeHintForColumn(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHintForColumn, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) UpdateThreshold() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_updateThreshold, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) SetUpdateThreshold(threshold int) {
	purego.SyscallN(addr_QAbstractItemView_setUpdateThreshold, this.h, uintptr(threshold))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) KeyboardSearchFlags() MatchFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_keyboardSearchFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MatchFlag(__r)
}

func (this *QAbstractItemView) SetKeyboardSearchFlags(searchFlags MatchFlag) {
	purego.SyscallN(addr_QAbstractItemView_setKeyboardSearchFlags, this.h, uintptr(searchFlags))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) OpenPersistentEditor(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_openPersistentEditor, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) ClosePersistentEditor(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_closePersistentEditor, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) IsPersistentEditorOpen(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_isPersistentEditorOpen, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemView) SetIndexWidget(index *QModelIndex, widget *QWidget) {
	purego.SyscallN(addr_QAbstractItemView_setIndexWidget, this.h, uintptr(index.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(widget)
}

func (this *QAbstractItemView) IndexWidget(index *QModelIndex) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_indexWidget, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQWidget(__r)
}

func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QAbstractItemView_setItemDelegateForRow, this.h, uintptr(row), uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegateForRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QAbstractItemView_setItemDelegateForColumn, this.h, uintptr(column), uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegateForColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) ItemDelegateWithIndex(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegate2, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegateForIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) Reset() {
	purego.SyscallN(addr_QAbstractItemView_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetRootIndex(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_setRootIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) DoItemsLayout() {
	purego.SyscallN(addr_QAbstractItemView_doItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SelectAll() {
	purego.SyscallN(addr_QAbstractItemView_selectAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Edit(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_edit, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) ClearSelection() {
	purego.SyscallN(addr_QAbstractItemView_clearSelection, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetCurrentIndex(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_setCurrentIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) ScrollToTop() {
	purego.SyscallN(addr_QAbstractItemView_scrollToTop, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) ScrollToBottom() {
	purego.SyscallN(addr_QAbstractItemView_scrollToBottom, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Update(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_update, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) Pressed(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_pressed, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) Clicked(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_clicked, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) DoubleClicked(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_doubleClicked, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) Activated(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_activated, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) Entered(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_entered, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) ViewportEntered() {
	purego.SyscallN(addr_QAbstractItemView_viewportEntered, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) IconSizeChanged(size *QSize) {
	purego.SyscallN(addr_QAbstractItemView_iconSizeChanged, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QAbstractItemView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QAbstractItemView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QAbstractItemView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QAbstractItemView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QAbstractItemView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QAbstractItemView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QAbstractItemView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QAbstractItemView) StartAutoScroll() {
	purego.SyscallN(addr_QAbstractItemView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) StopAutoScroll() {
	purego.SyscallN(addr_QAbstractItemView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DoAutoScroll() {
	purego.SyscallN(addr_QAbstractItemView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QAbstractItemView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QAbstractItemView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QAbstractItemView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QAbstractItemView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QAbstractItemView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QAbstractItemView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QAbstractItemView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QAbstractItemView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QAbstractItemView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_mouseGrabber)
	return bindQWidget(__r)
}

func QAbstractItemView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_keyboardGrabber)
	return bindQWidget(__r)
}

func QAbstractItemView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QAbstractItemView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QAbstractItemView) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractItemView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Create() {
	purego.SyscallN(addr_QAbstractItemView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Destroy() {
	purego.SyscallN(addr_QAbstractItemView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAbstractItemView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractItemView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractItemView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractItemView_DevicePixelRatioFScale() float64 {
	__r := pg_QAbstractItemView_devicePixelRatioFScale()
	return __r
}

func QAbstractItemView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QAbstractItemView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QAbstractItemView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QAbstractItemView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QAbstractItemView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractItemView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractItemView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QAbstractItemView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QAbstractItemView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QAbstractItemView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QAbstractItemView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QAbstractItemView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QAbstractItemView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QAbstractItemView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QAbstractItemView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QAbstractItemView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QAbstractItemView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QAbstractItemView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QAbstractItemView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QAbstractItemView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QAbstractItemView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QAbstractItemView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QAbstractItemView) ResetDefault() {
	purego.SyscallN(addr_QAbstractItemView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QAbstractItemView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) SelectAllDefault() {
	purego.SyscallN(addr_QAbstractItemView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) EditDefault(index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_edit_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractItemView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractItemView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QAbstractItemView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QAbstractItemView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QAbstractItemView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QAbstractItemView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QAbstractItemView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QAbstractItemView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QAbstractItemView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QAbstractItemView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QAbstractItemView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QAbstractItemView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QAbstractItemView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QAbstractItemView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QAbstractItemView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QAbstractItemView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QAbstractItemView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QAbstractItemView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QAbstractItemView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QAbstractItemView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractItemView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QAbstractItemView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QAbstractItemView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QAbstractItemView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QAbstractItemView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractItemView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractItemView_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractItemView_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractItemView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QAbstractItemView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QAbstractItemView_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QAbstractItemView_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QAbstractItemView_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractItemView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractItemView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractItemView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QAbstractItemView_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractItemView_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractItemView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractItemView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractItemView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractItemView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QAbstractItemView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QAbstractItemView) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QAbstractItemView_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractItemView) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QAbstractItemView_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractItemView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QAbstractItemView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractItemView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QAbstractItemView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QAbstractItemView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractItemView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QAbstractItemView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QAbstractItemView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QAbstractItemView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QAbstractItemView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractItemView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QAbstractItemView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractItemView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QAbstractItemView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QAbstractItemView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QAbstractItemView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QAbstractItemView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QAbstractItemView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QAbstractItemView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QAbstractItemView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QAbstractItemView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QAbstractItemView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QAbstractItemView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractItemView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractItemView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QAbstractItemView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QAbstractItemView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QAbstractItemView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QAbstractItemView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QAbstractItemView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QAbstractItemView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QAbstractItemView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QAbstractItemView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QAbstractItemView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QAbstractItemView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractItemView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractItemView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemView) Disconnect(handle int64) {
	pg_QAbstractItemView__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractItemView_virtIdx(slot int64) int32 {
	switch slot {
	case -1267696851865372997:
		return 0
	case -4832340528568269947:
		return 1
	case -7252823847285716802:
		return 2
	case -158523882194220890:
		return 3
	case -1500310024801917802:
		return 4
	case 5218707810677921016:
		return 5
	case 6057989667089378859:
		return 6
	case -1539266838375037305:
		return 7
	case 7063795720810552074:
		return 8
	case 1607084412591819503:
		return 9
	case 4283922507359547360:
		return 10
	case -4689322347236366780:
		return 11
	case -5041452885022114200:
		return 12
	case -4084029023708556876:
		return 13
	case -3666792239577137638:
		return 14
	case -3811438228839363824:
		return 15
	case 7345853097340709638:
		return 16
	case 1086335023529244512:
		return 17
	case -1574077111192992763:
		return 18
	case -1888629738259754754:
		return 19
	case 8758113809152574300:
		return 20
	case 4472434241583495620:
		return 21
	case -8151049665126058157:
		return 22
	case -5892205462145366901:
		return 23
	case -6978583972112316034:
		return 24
	case 9029941737538200621:
		return 25
	case -8726300139121449691:
		return 26
	case 1022771470898460576:
		return 27
	case -1155555122391250623:
		return 28
	case -785551405282638793:
		return 29
	case 6366953527806415524:
		return 30
	case 1786430177580376614:
		return 31
	case 8648484264628816806:
		return 32
	case -6021383853323633948:
		return 33
	case 6979418566323745651:
		return 34
	case 740869982805721908:
		return 35
	case 4206689019978932209:
		return 36
	case -2447490284072398302:
		return 37
	case -8112366604639385491:
		return 38
	case -7450787277202094556:
		return 39
	case 5898328517505784191:
		return 40
	case 1237236456312824991:
		return 41
	case -1971465653061861519:
		return 42
	case -332090898078543319:
		return 43
	case -4255494004099806581:
		return 44
	case -4415920049297795213:
		return 45
	case -4029073905697023887:
		return 46
	case 206122715919917589:
		return 47
	case -1781545137352745978:
		return 48
	case -3139904158373725006:
		return 49
	case 9107058682754591716:
		return 50
	case 975043767033552552:
		return 51
	case -385596010576048993:
		return 52
	case 1023791977679793185:
		return 53
	case -4302917285655014661:
		return 54
	case -5409927581812374498:
		return 55
	case 5438819661942843222:
		return 56
	case 5217623514630858297:
		return 57
	case 503579067252048239:
		return 58
	case 7049264599837532491:
		return 59
	case 729690261457684101:
		return 60
	case 4531996381736730688:
		return 61
	case 5290442163167458465:
		return 62
	case 1919988340934511316:
		return 63
	case -6698043222227006015:
		return 64
	case -6118001150126564745:
		return 65
	case 4959376063826165943:
		return 66
	case -488693769771308531:
		return 67
	case -8659679893868524044:
		return 68
	case 2110182719375245575:
		return 69
	case 4311534060495737339:
		return 70
	case -8266762397995423915:
		return 71
	case 6398007404756346381:
		return 72
	case 772598475950050299:
		return 73
	case 6316270346668237895:
		return 74
	case -551070009623711640:
		return 75
	case -9161949150489525396:
		return 76
	case -1989937050019464783:
		return 77
	case -212068866860701141:
		return 78
	case 1697444140088555581:
		return 79
	case -6825369174751916:
		return 80
	case -199718707876719129:
		return 81
	case -2039166586246443706:
		return 82
	case 427629649267311777:
		return 83
	case 6477500296980050210:
		return 84
	case 4127240633310655048:
		return 85
	case 5792771328741774991:
		return 86
	case -1185315367526857833:
		return 87
	case 3769868229123868474:
		return 88
	case 7379029982170503008:
		return 89
	case -568369772982215962:
		return 90
	}
	return -1
}

func QAbstractItemView_enable(h uintptr, idx int32) { pg_QAbstractItemView__set_override(h, idx) }

func NewQAbstractItemView() *QAbstractItemView {
	id := nextVirtID()
	o := newownedQAbstractItemView(pg_QAbstractItemView__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractItemView_virtIdx, QAbstractItemView_enable)
	return o
}

func (this *QAbstractItemView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QAbstractItemView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QAbstractItemView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(point uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(point))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QAbstractItemView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QAbstractItemView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QAbstractItemView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QAbstractItemView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QAbstractItemView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QAbstractItemView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QAbstractItemView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractItemView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QAbstractItemView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QAbstractItemView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QAbstractItemView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QAbstractItemView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QAbstractItemView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QAbstractItemView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractItemView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QAbstractItemView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QAbstractItemView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QAbstractItemView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QAbstractItemView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QAbstractItemView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QAbstractItemView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QAbstractItemView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QAbstractItemView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QAbstractItemView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QAbstractItemView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QAbstractItemView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QAbstractItemView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QAbstractItemView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QAbstractItemView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QAbstractItemView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QAbstractItemView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QAbstractItemView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QAbstractItemView) OnSelectedIndexes(fn func() []*QModelIndex) {
	registerVirt(this.UnsafeGoID(), 740869982805721908, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "selectedIndexes"
}

func (this *QAbstractItemView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QAbstractItemView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QAbstractItemView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QAbstractItemView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QAbstractItemView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QAbstractItemView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractItemView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QAbstractItemView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QAbstractItemView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractItemView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractItemView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractItemView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QAbstractItemView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QAbstractItemView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractItemView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QAbstractItemView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QAbstractItemView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QAbstractItemView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QAbstractItemView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QAbstractItemView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractItemView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QAbstractItemView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractItemView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QAbstractItemView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QAbstractItemView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractItemView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QAbstractItemView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QAbstractItemView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QAbstractItemView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QAbstractItemView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QAbstractItemView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QAbstractItemView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QAbstractItemView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QAbstractItemView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QAbstractItemView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QAbstractItemView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QAbstractItemView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QAbstractItemView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractItemView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QAbstractItemView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QAbstractItemView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QAbstractItemView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QAbstractItemView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QAbstractItemView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QAbstractItemView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QAbstractItemView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QAbstractItemView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QAbstractItemView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QAbstractItemView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QAbstractItemView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QAbstractItemView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QAbstractItemView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractItemView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractItemView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractItemView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractItemView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractItemView_delete(this.h)
	}
	this.UnsafeRelease()
}
