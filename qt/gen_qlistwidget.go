package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QListWidget_chain uintptr
	addr_QListWidget_new uintptr
	addr_QListWidget_new2 uintptr
	addr_QListWidget_tr uintptr
	addr_QListWidget_item uintptr
	addr_QListWidget_row uintptr
	addr_QListWidget_insertItem uintptr
	addr_QListWidget_insertItem2 uintptr
	addr_QListWidget_insertItems uintptr
	addr_QListWidget_addItem uintptr
	addr_QListWidget_addItem2 uintptr
	addr_QListWidget_addItems uintptr
	addr_QListWidget_takeItem uintptr
	addr_QListWidget_count uintptr
	addr_QListWidget_currentItem uintptr
	addr_QListWidget_setCurrentItem uintptr
	addr_QListWidget_setCurrentItem2 uintptr
	addr_QListWidget_currentRow uintptr
	addr_QListWidget_setCurrentRow uintptr
	addr_QListWidget_setCurrentRow2 uintptr
	addr_QListWidget_itemAt uintptr
	addr_QListWidget_itemAt2 uintptr
	addr_QListWidget_visualItemRect uintptr
	addr_QListWidget_sortItems uintptr
	addr_QListWidget_setSortingEnabled uintptr
	addr_QListWidget_isSortingEnabled uintptr
	addr_QListWidget_editItem uintptr
	addr_QListWidget_openPersistentEditor uintptr
	addr_QListWidget_closePersistentEditor uintptr
	addr_QListWidget_isPersistentEditorOpen uintptr
	addr_QListWidget_itemWidget uintptr
	addr_QListWidget_setItemWidget uintptr
	addr_QListWidget_removeItemWidget uintptr
	addr_QListWidget_selectedItems uintptr
	addr_QListWidget_findItems uintptr
	addr_QListWidget_items uintptr
	addr_QListWidget_indexFromItem uintptr
	addr_QListWidget_itemFromIndex uintptr
	addr_QListWidget_setSupportedDragActions uintptr
	addr_QListWidget_supportedDragActions uintptr
	addr_QListWidget_scrollToItem uintptr
	addr_QListWidget_clear uintptr
	addr_QListWidget_itemPressed uintptr
	addr_QListWidget_itemClicked uintptr
	addr_QListWidget_itemDoubleClicked uintptr
	addr_QListWidget_itemActivated uintptr
	addr_QListWidget_itemEntered uintptr
	addr_QListWidget_itemChanged uintptr
	addr_QListWidget_currentItemChanged uintptr
	addr_QListWidget_currentTextChanged uintptr
	addr_QListWidget_currentRowChanged uintptr
	addr_QListWidget_itemSelectionChanged uintptr
	addr_QListWidget_resizeContents uintptr
	addr_QListWidget_contentsSize uintptr
	addr_QListWidget_rectForIndex uintptr
	addr_QListWidget_setPositionForIndex uintptr
	addr_QListWidget_state uintptr
	addr_QListWidget_scheduleDelayedItemsLayout uintptr
	addr_QListWidget_executeDelayedItemsLayout uintptr
	addr_QListWidget_setDirtyRegion uintptr
	addr_QListWidget_scrollDirtyRegion uintptr
	addr_QListWidget_dirtyRegionOffset uintptr
	addr_QListWidget_startAutoScroll uintptr
	addr_QListWidget_stopAutoScroll uintptr
	addr_QListWidget_doAutoScroll uintptr
	addr_QListWidget_dropIndicatorPosition uintptr
	addr_QListWidget_setViewportMargins uintptr
	addr_QListWidget_setViewportMargins2 uintptr
	addr_QListWidget_viewportMargins uintptr
	addr_QListWidget_drawFrame uintptr
	addr_QListWidget_setTabOrder uintptr
	addr_QListWidget_mouseGrabber uintptr
	addr_QListWidget_keyboardGrabber uintptr
	addr_QListWidget_find uintptr
	addr_QListWidget_createWindowContainer uintptr
	addr_QListWidget_updateMicroFocus uintptr
	addr_QListWidget_create uintptr
	addr_QListWidget_destroy uintptr
	addr_QListWidget_focusNextChild uintptr
	addr_QListWidget_focusPreviousChild uintptr
	addr_QListWidget_connect uintptr
	addr_QListWidget_connect2 uintptr
	addr_QListWidget_sender uintptr
	addr_QListWidget_senderSignalIndex uintptr
	addr_QListWidget_receivers uintptr
	addr_QListWidget_isSignalConnected uintptr
	pg_QListWidget_devicePixelRatioFScale func() float64
	pg_QListWidget_encodeMetricF func(int32, float64) int32
	pg_QListWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QListWidget_tr2 uintptr
	addr_QListWidget_tr3 uintptr
	addr_QListWidget_sortItems2 uintptr
	addr_QListWidget_scrollToItem2 uintptr
	addr_QListWidget_createWindowContainer2 uintptr
	addr_QListWidget_createWindowContainer3 uintptr
	addr_QListWidget_updateMicroFocus2 uintptr
	addr_QListWidget_create2 uintptr
	addr_QListWidget_create3 uintptr
	addr_QListWidget_create4 uintptr
	addr_QListWidget_destroy2 uintptr
	addr_QListWidget_destroy3 uintptr
	addr_QListWidget_connect3 uintptr
	addr_QListWidget_connect4 uintptr
	addr_QListWidget_metaObject_default uintptr
	addr_QListWidget_setSelectionModel_default uintptr
	addr_QListWidget_dropEvent_default uintptr
	addr_QListWidget_event_default uintptr
	addr_QListWidget_mimeTypes_default uintptr
	addr_QListWidget_dropMimeData_default uintptr
	addr_QListWidget_supportedDropActions_default uintptr
	addr_QListWidget_visualRect_default uintptr
	addr_QListWidget_scrollTo_default uintptr
	addr_QListWidget_indexAt_default uintptr
	addr_QListWidget_doItemsLayout_default uintptr
	addr_QListWidget_reset_default uintptr
	addr_QListWidget_setRootIndex_default uintptr
	addr_QListWidget_scrollContentsBy_default uintptr
	addr_QListWidget_rowsInserted_default uintptr
	addr_QListWidget_rowsAboutToBeRemoved_default uintptr
	addr_QListWidget_mouseMoveEvent_default uintptr
	addr_QListWidget_mouseReleaseEvent_default uintptr
	addr_QListWidget_wheelEvent_default uintptr
	addr_QListWidget_timerEvent_default uintptr
	addr_QListWidget_resizeEvent_default uintptr
	addr_QListWidget_dragMoveEvent_default uintptr
	addr_QListWidget_dragLeaveEvent_default uintptr
	addr_QListWidget_startDrag_default uintptr
	addr_QListWidget_initViewItemOption_default uintptr
	addr_QListWidget_paintEvent_default uintptr
	addr_QListWidget_horizontalOffset_default uintptr
	addr_QListWidget_verticalOffset_default uintptr
	addr_QListWidget_moveCursor_default uintptr
	addr_QListWidget_setSelection_default uintptr
	addr_QListWidget_visualRegionForSelection_default uintptr
	addr_QListWidget_selectedIndexes_default uintptr
	addr_QListWidget_updateGeometries_default uintptr
	addr_QListWidget_isIndexHidden_default uintptr
	addr_QListWidget_selectionChanged_default uintptr
	addr_QListWidget_currentChanged_default uintptr
	addr_QListWidget_viewportSizeHint_default uintptr
	addr_QListWidget_keyboardSearch_default uintptr
	addr_QListWidget_sizeHintForRow_default uintptr
	addr_QListWidget_sizeHintForColumn_default uintptr
	addr_QListWidget_itemDelegateForIndex_default uintptr
	addr_QListWidget_inputMethodQuery_default uintptr
	addr_QListWidget_selectAll_default uintptr
	addr_QListWidget_updateEditorData_default uintptr
	addr_QListWidget_updateEditorGeometries_default uintptr
	addr_QListWidget_verticalScrollbarAction_default uintptr
	addr_QListWidget_horizontalScrollbarAction_default uintptr
	addr_QListWidget_verticalScrollbarValueChanged_default uintptr
	addr_QListWidget_horizontalScrollbarValueChanged_default uintptr
	addr_QListWidget_closeEditor_default uintptr
	addr_QListWidget_commitData_default uintptr
	addr_QListWidget_editorDestroyed_default uintptr
	addr_QListWidget_edit_default uintptr
	addr_QListWidget_selectionCommand_default uintptr
	addr_QListWidget_focusNextPrevChild_default uintptr
	addr_QListWidget_viewportEvent_default uintptr
	addr_QListWidget_mousePressEvent_default uintptr
	addr_QListWidget_mouseDoubleClickEvent_default uintptr
	addr_QListWidget_dragEnterEvent_default uintptr
	addr_QListWidget_focusInEvent_default uintptr
	addr_QListWidget_focusOutEvent_default uintptr
	addr_QListWidget_keyPressEvent_default uintptr
	addr_QListWidget_inputMethodEvent_default uintptr
	addr_QListWidget_eventFilter_default uintptr
	addr_QListWidget_minimumSizeHint_default uintptr
	addr_QListWidget_sizeHint_default uintptr
	addr_QListWidget_setupViewport_default uintptr
	addr_QListWidget_contextMenuEvent_default uintptr
	addr_QListWidget_changeEvent_default uintptr
	addr_QListWidget_initStyleOption_default uintptr
	addr_QListWidget_devType_default uintptr
	addr_QListWidget_setVisible_default uintptr
	addr_QListWidget_heightForWidth_default uintptr
	addr_QListWidget_hasHeightForWidth_default uintptr
	addr_QListWidget_paintEngine_default uintptr
	addr_QListWidget_keyReleaseEvent_default uintptr
	addr_QListWidget_enterEvent_default uintptr
	addr_QListWidget_leaveEvent_default uintptr
	addr_QListWidget_moveEvent_default uintptr
	addr_QListWidget_closeEvent_default uintptr
	addr_QListWidget_tabletEvent_default uintptr
	addr_QListWidget_actionEvent_default uintptr
	addr_QListWidget_showEvent_default uintptr
	addr_QListWidget_hideEvent_default uintptr
	addr_QListWidget_metric_default uintptr
	addr_QListWidget_initPainter_default uintptr
	addr_QListWidget_redirected_default uintptr
	addr_QListWidget_sharedPainter_default uintptr
	addr_QListWidget_childEvent_default uintptr
	addr_QListWidget_customEvent_default uintptr
	addr_QListWidget_connectNotify_default uintptr
	addr_QListWidget_disconnectNotify_default uintptr
	pg_QListWidget_connect_itemPressed func(uintptr, int64)
	pg_QListWidget_connect_itemClicked func(uintptr, int64)
	pg_QListWidget_connect_itemDoubleClicked func(uintptr, int64)
	pg_QListWidget_connect_itemActivated func(uintptr, int64)
	pg_QListWidget_connect_itemEntered func(uintptr, int64)
	pg_QListWidget_connect_itemChanged func(uintptr, int64)
	pg_QListWidget_connect_currentItemChanged func(uintptr, int64)
	pg_QListWidget_connect_currentTextChanged func(uintptr, int64)
	pg_QListWidget_connect_currentRowChanged func(uintptr, int64)
	pg_QListWidget_connect_itemSelectionChanged func(uintptr, int64)
	pg_QListWidget_connect_indexesMoved func(uintptr, int64)
	pg_QListWidget_connect_pressed func(uintptr, int64)
	pg_QListWidget_connect_clicked func(uintptr, int64)
	pg_QListWidget_connect_doubleClicked func(uintptr, int64)
	pg_QListWidget_connect_activated func(uintptr, int64)
	pg_QListWidget_connect_entered func(uintptr, int64)
	pg_QListWidget_connect_viewportEntered func(uintptr, int64)
	pg_QListWidget_connect_iconSizeChanged func(uintptr, int64)
	pg_QListWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QListWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QListWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QListWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QListWidget_connect_destroyed func(uintptr, int64)
	pg_QListWidget_connect_destroyed2 func(uintptr, int64)
	pg_QListWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QListWidget__disconnect func(int64)
	pg_QListWidget__set_override func(uintptr, int32)
	pg_QListWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QListWidget_chain, "QListWidget_chain")
	symaddr(&addr_QListWidget_new, "QListWidget_new")
	symaddr(&addr_QListWidget_new2, "QListWidget_new2")
	symaddr(&addr_QListWidget_tr, "QListWidget_tr")
	symaddr(&addr_QListWidget_item, "QListWidget_item")
	symaddr(&addr_QListWidget_row, "QListWidget_row")
	symaddr(&addr_QListWidget_insertItem, "QListWidget_insertItem")
	symaddr(&addr_QListWidget_insertItem2, "QListWidget_insertItem2")
	symaddr(&addr_QListWidget_insertItems, "QListWidget_insertItems")
	symaddr(&addr_QListWidget_addItem, "QListWidget_addItem")
	symaddr(&addr_QListWidget_addItem2, "QListWidget_addItem2")
	symaddr(&addr_QListWidget_addItems, "QListWidget_addItems")
	symaddr(&addr_QListWidget_takeItem, "QListWidget_takeItem")
	symaddr(&addr_QListWidget_count, "QListWidget_count")
	symaddr(&addr_QListWidget_currentItem, "QListWidget_currentItem")
	symaddr(&addr_QListWidget_setCurrentItem, "QListWidget_setCurrentItem")
	symaddr(&addr_QListWidget_setCurrentItem2, "QListWidget_setCurrentItem2")
	symaddr(&addr_QListWidget_currentRow, "QListWidget_currentRow")
	symaddr(&addr_QListWidget_setCurrentRow, "QListWidget_setCurrentRow")
	symaddr(&addr_QListWidget_setCurrentRow2, "QListWidget_setCurrentRow2")
	symaddr(&addr_QListWidget_itemAt, "QListWidget_itemAt")
	symaddr(&addr_QListWidget_itemAt2, "QListWidget_itemAt2")
	symaddr(&addr_QListWidget_visualItemRect, "QListWidget_visualItemRect")
	symaddr(&addr_QListWidget_sortItems, "QListWidget_sortItems")
	symaddr(&addr_QListWidget_setSortingEnabled, "QListWidget_setSortingEnabled")
	symaddr(&addr_QListWidget_isSortingEnabled, "QListWidget_isSortingEnabled")
	symaddr(&addr_QListWidget_editItem, "QListWidget_editItem")
	symaddr(&addr_QListWidget_openPersistentEditor, "QListWidget_openPersistentEditor")
	symaddr(&addr_QListWidget_closePersistentEditor, "QListWidget_closePersistentEditor")
	symaddr(&addr_QListWidget_isPersistentEditorOpen, "QListWidget_isPersistentEditorOpen")
	symaddr(&addr_QListWidget_itemWidget, "QListWidget_itemWidget")
	symaddr(&addr_QListWidget_setItemWidget, "QListWidget_setItemWidget")
	symaddr(&addr_QListWidget_removeItemWidget, "QListWidget_removeItemWidget")
	symaddr(&addr_QListWidget_selectedItems, "QListWidget_selectedItems")
	symaddr(&addr_QListWidget_findItems, "QListWidget_findItems")
	symaddr(&addr_QListWidget_items, "QListWidget_items")
	symaddr(&addr_QListWidget_indexFromItem, "QListWidget_indexFromItem")
	symaddr(&addr_QListWidget_itemFromIndex, "QListWidget_itemFromIndex")
	symaddr(&addr_QListWidget_setSupportedDragActions, "QListWidget_setSupportedDragActions")
	symaddr(&addr_QListWidget_supportedDragActions, "QListWidget_supportedDragActions")
	symaddr(&addr_QListWidget_scrollToItem, "QListWidget_scrollToItem")
	symaddr(&addr_QListWidget_clear, "QListWidget_clear")
	symaddr(&addr_QListWidget_itemPressed, "QListWidget_itemPressed")
	symaddr(&addr_QListWidget_itemClicked, "QListWidget_itemClicked")
	symaddr(&addr_QListWidget_itemDoubleClicked, "QListWidget_itemDoubleClicked")
	symaddr(&addr_QListWidget_itemActivated, "QListWidget_itemActivated")
	symaddr(&addr_QListWidget_itemEntered, "QListWidget_itemEntered")
	symaddr(&addr_QListWidget_itemChanged, "QListWidget_itemChanged")
	symaddr(&addr_QListWidget_currentItemChanged, "QListWidget_currentItemChanged")
	symaddr(&addr_QListWidget_currentTextChanged, "QListWidget_currentTextChanged")
	symaddr(&addr_QListWidget_currentRowChanged, "QListWidget_currentRowChanged")
	symaddr(&addr_QListWidget_itemSelectionChanged, "QListWidget_itemSelectionChanged")
	symaddr(&addr_QListWidget_resizeContents, "QListWidget_resizeContents")
	symaddr(&addr_QListWidget_contentsSize, "QListWidget_contentsSize")
	symaddr(&addr_QListWidget_rectForIndex, "QListWidget_rectForIndex")
	symaddr(&addr_QListWidget_setPositionForIndex, "QListWidget_setPositionForIndex")
	symaddr(&addr_QListWidget_state, "QListWidget_state")
	symaddr(&addr_QListWidget_scheduleDelayedItemsLayout, "QListWidget_scheduleDelayedItemsLayout")
	symaddr(&addr_QListWidget_executeDelayedItemsLayout, "QListWidget_executeDelayedItemsLayout")
	symaddr(&addr_QListWidget_setDirtyRegion, "QListWidget_setDirtyRegion")
	symaddr(&addr_QListWidget_scrollDirtyRegion, "QListWidget_scrollDirtyRegion")
	symaddr(&addr_QListWidget_dirtyRegionOffset, "QListWidget_dirtyRegionOffset")
	symaddr(&addr_QListWidget_startAutoScroll, "QListWidget_startAutoScroll")
	symaddr(&addr_QListWidget_stopAutoScroll, "QListWidget_stopAutoScroll")
	symaddr(&addr_QListWidget_doAutoScroll, "QListWidget_doAutoScroll")
	symaddr(&addr_QListWidget_dropIndicatorPosition, "QListWidget_dropIndicatorPosition")
	symaddr(&addr_QListWidget_setViewportMargins, "QListWidget_setViewportMargins")
	symaddr(&addr_QListWidget_setViewportMargins2, "QListWidget_setViewportMargins2")
	symaddr(&addr_QListWidget_viewportMargins, "QListWidget_viewportMargins")
	symaddr(&addr_QListWidget_drawFrame, "QListWidget_drawFrame")
	symaddr(&addr_QListWidget_setTabOrder, "QListWidget_setTabOrder")
	symaddr(&addr_QListWidget_mouseGrabber, "QListWidget_mouseGrabber")
	symaddr(&addr_QListWidget_keyboardGrabber, "QListWidget_keyboardGrabber")
	symaddr(&addr_QListWidget_find, "QListWidget_find")
	symaddr(&addr_QListWidget_createWindowContainer, "QListWidget_createWindowContainer")
	symaddr(&addr_QListWidget_updateMicroFocus, "QListWidget_updateMicroFocus")
	symaddr(&addr_QListWidget_create, "QListWidget_create")
	symaddr(&addr_QListWidget_destroy, "QListWidget_destroy")
	symaddr(&addr_QListWidget_focusNextChild, "QListWidget_focusNextChild")
	symaddr(&addr_QListWidget_focusPreviousChild, "QListWidget_focusPreviousChild")
	symaddr(&addr_QListWidget_connect, "QListWidget_connect")
	symaddr(&addr_QListWidget_connect2, "QListWidget_connect2")
	symaddr(&addr_QListWidget_sender, "QListWidget_sender")
	symaddr(&addr_QListWidget_senderSignalIndex, "QListWidget_senderSignalIndex")
	symaddr(&addr_QListWidget_receivers, "QListWidget_receivers")
	symaddr(&addr_QListWidget_isSignalConnected, "QListWidget_isSignalConnected")
	reg(&pg_QListWidget_devicePixelRatioFScale, "QListWidget_devicePixelRatioFScale")
	reg(&pg_QListWidget_encodeMetricF, "QListWidget_encodeMetricF")
	reg(&pg_QListWidget_getDecodedMetricF, "QListWidget_getDecodedMetricF")
	symaddr(&addr_QListWidget_tr2, "QListWidget_tr2")
	symaddr(&addr_QListWidget_tr3, "QListWidget_tr3")
	symaddr(&addr_QListWidget_sortItems2, "QListWidget_sortItems2")
	symaddr(&addr_QListWidget_scrollToItem2, "QListWidget_scrollToItem2")
	symaddr(&addr_QListWidget_createWindowContainer2, "QListWidget_createWindowContainer2")
	symaddr(&addr_QListWidget_createWindowContainer3, "QListWidget_createWindowContainer3")
	symaddr(&addr_QListWidget_updateMicroFocus2, "QListWidget_updateMicroFocus2")
	symaddr(&addr_QListWidget_create2, "QListWidget_create2")
	symaddr(&addr_QListWidget_create3, "QListWidget_create3")
	symaddr(&addr_QListWidget_create4, "QListWidget_create4")
	symaddr(&addr_QListWidget_destroy2, "QListWidget_destroy2")
	symaddr(&addr_QListWidget_destroy3, "QListWidget_destroy3")
	symaddr(&addr_QListWidget_connect3, "QListWidget_connect3")
	symaddr(&addr_QListWidget_connect4, "QListWidget_connect4")
	symaddr(&addr_QListWidget_metaObject_default, "QListWidget_metaObject_default")
	symaddr(&addr_QListWidget_setSelectionModel_default, "QListWidget_setSelectionModel_default")
	symaddr(&addr_QListWidget_dropEvent_default, "QListWidget_dropEvent_default")
	symaddr(&addr_QListWidget_event_default, "QListWidget_event_default")
	symaddr(&addr_QListWidget_mimeTypes_default, "QListWidget_mimeTypes_default")
	symaddr(&addr_QListWidget_dropMimeData_default, "QListWidget_dropMimeData_default")
	symaddr(&addr_QListWidget_supportedDropActions_default, "QListWidget_supportedDropActions_default")
	symaddr(&addr_QListWidget_visualRect_default, "QListWidget_visualRect_default")
	symaddr(&addr_QListWidget_scrollTo_default, "QListWidget_scrollTo_default")
	symaddr(&addr_QListWidget_indexAt_default, "QListWidget_indexAt_default")
	symaddr(&addr_QListWidget_doItemsLayout_default, "QListWidget_doItemsLayout_default")
	symaddr(&addr_QListWidget_reset_default, "QListWidget_reset_default")
	symaddr(&addr_QListWidget_setRootIndex_default, "QListWidget_setRootIndex_default")
	symaddr(&addr_QListWidget_scrollContentsBy_default, "QListWidget_scrollContentsBy_default")
	symaddr(&addr_QListWidget_rowsInserted_default, "QListWidget_rowsInserted_default")
	symaddr(&addr_QListWidget_rowsAboutToBeRemoved_default, "QListWidget_rowsAboutToBeRemoved_default")
	symaddr(&addr_QListWidget_mouseMoveEvent_default, "QListWidget_mouseMoveEvent_default")
	symaddr(&addr_QListWidget_mouseReleaseEvent_default, "QListWidget_mouseReleaseEvent_default")
	symaddr(&addr_QListWidget_wheelEvent_default, "QListWidget_wheelEvent_default")
	symaddr(&addr_QListWidget_timerEvent_default, "QListWidget_timerEvent_default")
	symaddr(&addr_QListWidget_resizeEvent_default, "QListWidget_resizeEvent_default")
	symaddr(&addr_QListWidget_dragMoveEvent_default, "QListWidget_dragMoveEvent_default")
	symaddr(&addr_QListWidget_dragLeaveEvent_default, "QListWidget_dragLeaveEvent_default")
	symaddr(&addr_QListWidget_startDrag_default, "QListWidget_startDrag_default")
	symaddr(&addr_QListWidget_initViewItemOption_default, "QListWidget_initViewItemOption_default")
	symaddr(&addr_QListWidget_paintEvent_default, "QListWidget_paintEvent_default")
	symaddr(&addr_QListWidget_horizontalOffset_default, "QListWidget_horizontalOffset_default")
	symaddr(&addr_QListWidget_verticalOffset_default, "QListWidget_verticalOffset_default")
	symaddr(&addr_QListWidget_moveCursor_default, "QListWidget_moveCursor_default")
	symaddr(&addr_QListWidget_setSelection_default, "QListWidget_setSelection_default")
	symaddr(&addr_QListWidget_visualRegionForSelection_default, "QListWidget_visualRegionForSelection_default")
	symaddr(&addr_QListWidget_selectedIndexes_default, "QListWidget_selectedIndexes_default")
	symaddr(&addr_QListWidget_updateGeometries_default, "QListWidget_updateGeometries_default")
	symaddr(&addr_QListWidget_isIndexHidden_default, "QListWidget_isIndexHidden_default")
	symaddr(&addr_QListWidget_selectionChanged_default, "QListWidget_selectionChanged_default")
	symaddr(&addr_QListWidget_currentChanged_default, "QListWidget_currentChanged_default")
	symaddr(&addr_QListWidget_viewportSizeHint_default, "QListWidget_viewportSizeHint_default")
	symaddr(&addr_QListWidget_keyboardSearch_default, "QListWidget_keyboardSearch_default")
	symaddr(&addr_QListWidget_sizeHintForRow_default, "QListWidget_sizeHintForRow_default")
	symaddr(&addr_QListWidget_sizeHintForColumn_default, "QListWidget_sizeHintForColumn_default")
	symaddr(&addr_QListWidget_itemDelegateForIndex_default, "QListWidget_itemDelegateForIndex_default")
	symaddr(&addr_QListWidget_inputMethodQuery_default, "QListWidget_inputMethodQuery_default")
	symaddr(&addr_QListWidget_selectAll_default, "QListWidget_selectAll_default")
	symaddr(&addr_QListWidget_updateEditorData_default, "QListWidget_updateEditorData_default")
	symaddr(&addr_QListWidget_updateEditorGeometries_default, "QListWidget_updateEditorGeometries_default")
	symaddr(&addr_QListWidget_verticalScrollbarAction_default, "QListWidget_verticalScrollbarAction_default")
	symaddr(&addr_QListWidget_horizontalScrollbarAction_default, "QListWidget_horizontalScrollbarAction_default")
	symaddr(&addr_QListWidget_verticalScrollbarValueChanged_default, "QListWidget_verticalScrollbarValueChanged_default")
	symaddr(&addr_QListWidget_horizontalScrollbarValueChanged_default, "QListWidget_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QListWidget_closeEditor_default, "QListWidget_closeEditor_default")
	symaddr(&addr_QListWidget_commitData_default, "QListWidget_commitData_default")
	symaddr(&addr_QListWidget_editorDestroyed_default, "QListWidget_editorDestroyed_default")
	symaddr(&addr_QListWidget_edit_default, "QListWidget_edit_default")
	symaddr(&addr_QListWidget_selectionCommand_default, "QListWidget_selectionCommand_default")
	symaddr(&addr_QListWidget_focusNextPrevChild_default, "QListWidget_focusNextPrevChild_default")
	symaddr(&addr_QListWidget_viewportEvent_default, "QListWidget_viewportEvent_default")
	symaddr(&addr_QListWidget_mousePressEvent_default, "QListWidget_mousePressEvent_default")
	symaddr(&addr_QListWidget_mouseDoubleClickEvent_default, "QListWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QListWidget_dragEnterEvent_default, "QListWidget_dragEnterEvent_default")
	symaddr(&addr_QListWidget_focusInEvent_default, "QListWidget_focusInEvent_default")
	symaddr(&addr_QListWidget_focusOutEvent_default, "QListWidget_focusOutEvent_default")
	symaddr(&addr_QListWidget_keyPressEvent_default, "QListWidget_keyPressEvent_default")
	symaddr(&addr_QListWidget_inputMethodEvent_default, "QListWidget_inputMethodEvent_default")
	symaddr(&addr_QListWidget_eventFilter_default, "QListWidget_eventFilter_default")
	symaddr(&addr_QListWidget_minimumSizeHint_default, "QListWidget_minimumSizeHint_default")
	symaddr(&addr_QListWidget_sizeHint_default, "QListWidget_sizeHint_default")
	symaddr(&addr_QListWidget_setupViewport_default, "QListWidget_setupViewport_default")
	symaddr(&addr_QListWidget_contextMenuEvent_default, "QListWidget_contextMenuEvent_default")
	symaddr(&addr_QListWidget_changeEvent_default, "QListWidget_changeEvent_default")
	symaddr(&addr_QListWidget_initStyleOption_default, "QListWidget_initStyleOption_default")
	symaddr(&addr_QListWidget_devType_default, "QListWidget_devType_default")
	symaddr(&addr_QListWidget_setVisible_default, "QListWidget_setVisible_default")
	symaddr(&addr_QListWidget_heightForWidth_default, "QListWidget_heightForWidth_default")
	symaddr(&addr_QListWidget_hasHeightForWidth_default, "QListWidget_hasHeightForWidth_default")
	symaddr(&addr_QListWidget_paintEngine_default, "QListWidget_paintEngine_default")
	symaddr(&addr_QListWidget_keyReleaseEvent_default, "QListWidget_keyReleaseEvent_default")
	symaddr(&addr_QListWidget_enterEvent_default, "QListWidget_enterEvent_default")
	symaddr(&addr_QListWidget_leaveEvent_default, "QListWidget_leaveEvent_default")
	symaddr(&addr_QListWidget_moveEvent_default, "QListWidget_moveEvent_default")
	symaddr(&addr_QListWidget_closeEvent_default, "QListWidget_closeEvent_default")
	symaddr(&addr_QListWidget_tabletEvent_default, "QListWidget_tabletEvent_default")
	symaddr(&addr_QListWidget_actionEvent_default, "QListWidget_actionEvent_default")
	symaddr(&addr_QListWidget_showEvent_default, "QListWidget_showEvent_default")
	symaddr(&addr_QListWidget_hideEvent_default, "QListWidget_hideEvent_default")
	symaddr(&addr_QListWidget_metric_default, "QListWidget_metric_default")
	symaddr(&addr_QListWidget_initPainter_default, "QListWidget_initPainter_default")
	symaddr(&addr_QListWidget_redirected_default, "QListWidget_redirected_default")
	symaddr(&addr_QListWidget_sharedPainter_default, "QListWidget_sharedPainter_default")
	symaddr(&addr_QListWidget_childEvent_default, "QListWidget_childEvent_default")
	symaddr(&addr_QListWidget_customEvent_default, "QListWidget_customEvent_default")
	symaddr(&addr_QListWidget_connectNotify_default, "QListWidget_connectNotify_default")
	symaddr(&addr_QListWidget_disconnectNotify_default, "QListWidget_disconnectNotify_default")
	reg(&pg_QListWidget_connect_itemPressed, "QListWidget_connect_itemPressed")
	reg(&pg_QListWidget_connect_itemClicked, "QListWidget_connect_itemClicked")
	reg(&pg_QListWidget_connect_itemDoubleClicked, "QListWidget_connect_itemDoubleClicked")
	reg(&pg_QListWidget_connect_itemActivated, "QListWidget_connect_itemActivated")
	reg(&pg_QListWidget_connect_itemEntered, "QListWidget_connect_itemEntered")
	reg(&pg_QListWidget_connect_itemChanged, "QListWidget_connect_itemChanged")
	reg(&pg_QListWidget_connect_currentItemChanged, "QListWidget_connect_currentItemChanged")
	reg(&pg_QListWidget_connect_currentTextChanged, "QListWidget_connect_currentTextChanged")
	reg(&pg_QListWidget_connect_currentRowChanged, "QListWidget_connect_currentRowChanged")
	reg(&pg_QListWidget_connect_itemSelectionChanged, "QListWidget_connect_itemSelectionChanged")
	reg(&pg_QListWidget_connect_indexesMoved, "QListWidget_connect_indexesMoved")
	reg(&pg_QListWidget_connect_pressed, "QListWidget_connect_pressed")
	reg(&pg_QListWidget_connect_clicked, "QListWidget_connect_clicked")
	reg(&pg_QListWidget_connect_doubleClicked, "QListWidget_connect_doubleClicked")
	reg(&pg_QListWidget_connect_activated, "QListWidget_connect_activated")
	reg(&pg_QListWidget_connect_entered, "QListWidget_connect_entered")
	reg(&pg_QListWidget_connect_viewportEntered, "QListWidget_connect_viewportEntered")
	reg(&pg_QListWidget_connect_iconSizeChanged, "QListWidget_connect_iconSizeChanged")
	reg(&pg_QListWidget_connect_windowTitleChanged, "QListWidget_connect_windowTitleChanged")
	reg(&pg_QListWidget_connect_windowIconChanged, "QListWidget_connect_windowIconChanged")
	reg(&pg_QListWidget_connect_windowIconTextChanged, "QListWidget_connect_windowIconTextChanged")
	reg(&pg_QListWidget_connect_customContextMenuRequested, "QListWidget_connect_customContextMenuRequested")
	reg(&pg_QListWidget_connect_destroyed, "QListWidget_connect_destroyed")
	reg(&pg_QListWidget_connect_destroyed2, "QListWidget_connect_destroyed2")
	reg(&pg_QListWidget_connect_objectNameChanged, "QListWidget_connect_objectNameChanged")
	reg(&pg_QListWidget__disconnect, "QListWidget__disconnect")
	reg(&pg_QListWidget__set_override, "QListWidget__set_override")
	reg(&pg_QListWidget_delete, "QListWidget_delete")
}

type QListWidget struct {
	h uintptr
	*QListView
}

func bindQListWidget(h uintptr) *QListWidget {
	if h == 0 {
		return nil
	}
	var __c [6]uintptr
	purego.SyscallN(addr_QListWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QListWidget{h: h, QListView: &QListView{h: __c[0], QAbstractItemView: &QAbstractItemView{h: __c[1], QAbstractScrollArea: &QAbstractScrollArea{h: __c[2], QFrame: &QFrame{h: __c[3], QWidget: &QWidget{h: __c[4], QObject: &QObject{h: __c[5]}}}}}}}
}

func newownedQListWidget(h uintptr) *QListWidget { return bindQListWidget(h) }

func (this *QListWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QListWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQListWidget(h unsafe.Pointer) *QListWidget { return bindQListWidget(uintptr(h)) }
func UnsafeNewQListWidgetOwned(h unsafe.Pointer) *QListWidget { return newownedQListWidget(uintptr(h)) }

func (this *QListWidget) UnsafeRelease() {
	this.h = 0
	if this.QListView != nil {
		this.QListView.UnsafeRelease()
	}
}

func NewQListWidget(parent *QWidget) *QListWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQListWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidget_virtIdx, QListWidget_enable)
	return __o
}

func NewQListWidget2() *QListWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidget_new2, uintptr(__goid))
	__o := newownedQListWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidget_virtIdx, QListWidget_enable)
	return __o
}

func QListWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QListWidget) Item(row int) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_item, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) Row(item *QListWidgetItem) int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_row, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return int(__r)
}

func (this *QListWidget) InsertItem(row int, item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_insertItem, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) InsertItem2(row int, label string) {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	purego.SyscallN(addr_QListWidget_insertItem2, this.h, uintptr(row), label_p, uintptr(len(label)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
}

func (this *QListWidget) InsertItems(row int, labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QListWidget_insertItems, this.h, uintptr(row), labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QListWidget) AddItem(label string) {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	purego.SyscallN(addr_QListWidget_addItem, this.h, label_p, uintptr(len(label)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
}

func (this *QListWidget) AddItemWithItem(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_addItem2, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) AddItems(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QListWidget_addItems, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QListWidget) TakeItem(row int) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_takeItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) CurrentItem() *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_currentItem, this.h)
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) SetCurrentItem(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_setCurrentItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) SetCurrentItem2(item *QListWidgetItem, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QListWidget_setCurrentItem2, this.h, uintptr(item.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) CurrentRow() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_currentRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) SetCurrentRow(row int) {
	purego.SyscallN(addr_QListWidget_setCurrentRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QListWidget) SetCurrentRow2(row int, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QListWidget_setCurrentRow2, this.h, uintptr(row), uintptr(command))
	runtime.KeepAlive(this)
}

func (this *QListWidget) ItemAt(p *QPoint) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_itemAt, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) ItemAt2(x int, y int) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_itemAt2, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) VisualItemRect(item *QListWidgetItem) *QRect {
	__r, _, _ := purego.SyscallN(addr_QListWidget_visualItemRect, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRect(__r)
}

func (this *QListWidget) SortItems() {
	purego.SyscallN(addr_QListWidget_sortItems, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) SetSortingEnabled(enable bool) {
	purego.SyscallN(addr_QListWidget_setSortingEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QListWidget) IsSortingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_isSortingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidget) EditItem(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_editItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) OpenPersistentEditor(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_openPersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ClosePersistentEditor(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_closePersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) IsPersistentEditorOpen(item *QListWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_isPersistentEditorOpen, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QListWidget) ItemWidget(item *QListWidgetItem) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_itemWidget, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQWidget(__r)
}

func (this *QListWidget) SetItemWidget(item *QListWidgetItem, widget *QWidget) {
	purego.SyscallN(addr_QListWidget_setItemWidget, this.h, uintptr(item.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(widget)
}

func (this *QListWidget) RemoveItemWidget(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_removeItemWidget, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) SelectedItems() []*QListWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_selectedItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QListWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQListWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QListWidget) FindItems(text string, flags MatchFlag) []*QListWidgetItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_findItems, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QListWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQListWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QListWidget) Items(data *QMimeData) []*QListWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_items, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	__out := make([]*QListWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQListWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QListWidget) IndexFromItem(item *QListWidgetItem) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QListWidget_indexFromItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQModelIndex(__r)
}

func (this *QListWidget) ItemFromIndex(index *QModelIndex) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidget_itemFromIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQListWidgetItem(__r)
}

func (this *QListWidget) SetSupportedDragActions(actions DropAction) {
	purego.SyscallN(addr_QListWidget_setSupportedDragActions, this.h, uintptr(actions))
	runtime.KeepAlive(this)
}

func (this *QListWidget) SupportedDragActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_supportedDragActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QListWidget) ScrollToItem(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_scrollToItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) Clear() {
	purego.SyscallN(addr_QListWidget_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) ItemPressed(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemPressed, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ItemClicked(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemClicked, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ItemDoubleClicked(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemDoubleClicked, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ItemActivated(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemActivated, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ItemEntered(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemEntered, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) ItemChanged(item *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_itemChanged, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QListWidget) CurrentItemChanged(current *QListWidgetItem, previous *QListWidgetItem) {
	purego.SyscallN(addr_QListWidget_currentItemChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QListWidget) CurrentTextChanged(currentText string) {
	var currentText_p uintptr
	if len(currentText) > 0 {
		currentText_p = uintptr(unsafe.Pointer(unsafe.StringData(currentText)))
	}
	purego.SyscallN(addr_QListWidget_currentTextChanged, this.h, currentText_p, uintptr(len(currentText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(currentText)
}

func (this *QListWidget) CurrentRowChanged(currentRow int) {
	purego.SyscallN(addr_QListWidget_currentRowChanged, this.h, uintptr(currentRow))
	runtime.KeepAlive(this)
}

func (this *QListWidget) ItemSelectionChanged() {
	purego.SyscallN(addr_QListWidget_itemSelectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) ResizeContents(width int, height int) {
	purego.SyscallN(addr_QListWidget_resizeContents, this.h, uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QListWidget) ContentsSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListWidget_contentsSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListWidget) RectForIndex(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QListWidget_rectForIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QListWidget) SetPositionForIndex(position *QPoint, index *QModelIndex) {
	purego.SyscallN(addr_QListWidget_setPositionForIndex, this.h, uintptr(position.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	runtime.KeepAlive(index)
}

func (this *QListWidget) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QListWidget) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QListWidget_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QListWidget_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QListWidget_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QListWidget) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QListWidget_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QListWidget) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QListWidget_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QListWidget) StartAutoScroll() {
	purego.SyscallN(addr_QListWidget_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) StopAutoScroll() {
	purego.SyscallN(addr_QListWidget_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) DoAutoScroll() {
	purego.SyscallN(addr_QListWidget_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QListWidget) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QListWidget_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QListWidget) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QListWidget_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QListWidget) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QListWidget_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QListWidget) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QListWidget_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QListWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QListWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QListWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QListWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QListWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QListWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QListWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QListWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) Create() {
	purego.SyscallN(addr_QListWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) Destroy() {
	purego.SyscallN(addr_QListWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QListWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QListWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QListWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QListWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QListWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QListWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QListWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QListWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QListWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QListWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QListWidget_devicePixelRatioFScale()
	return __r
}

func QListWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QListWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QListWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QListWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QListWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QListWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QListWidget) SortItemsWithOrder(order SortOrder) {
	purego.SyscallN(addr_QListWidget_sortItems2, this.h, uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QListWidget) ScrollToItem2(item *QListWidgetItem, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QListWidget_scrollToItem2, this.h, uintptr(item.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func QListWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QListWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QListWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QListWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QListWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QListWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QListWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QListWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QListWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QListWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QListWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QListWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QListWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QListWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QListWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QListWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QListWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QListWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QListWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QListWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QListWidget) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QListWidget_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QListWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QListWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QListWidget) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QListWidget) DropMimeDataDefault(index int, data *QMimeData, action DropAction) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_dropMimeData_default, this.h, uintptr(index), uintptr(data.UnsafePointer()), uintptr(action))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QListWidget) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QListWidget) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QListWidget_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QListWidget) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QListWidget_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QListWidget) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QListWidget_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QListWidget) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QListWidget_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) ResetDefault() {
	purego.SyscallN(addr_QListWidget_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QListWidget_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QListWidget) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QListWidget_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QListWidget) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QListWidget_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QListWidget) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QListWidget_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QListWidget) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QListWidget_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QListWidget_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QListWidget_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QListWidget_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QListWidget_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QListWidget_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QListWidget_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QListWidget_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QListWidget) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QListWidget_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QListWidget) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QListWidget_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListWidget) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QListWidget_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QListWidget) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QListWidget_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QListWidget) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QListWidget_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QListWidget) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidget_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QListWidget) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QListWidget_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QListWidget) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QListWidget_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QListWidget) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QListWidget_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QListWidget) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListWidget_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListWidget) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QListWidget_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QListWidget) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QListWidget_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QListWidget) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QListWidget_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QListWidget) SelectAllDefault() {
	purego.SyscallN(addr_QListWidget_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QListWidget_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QListWidget_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListWidget) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QListWidget_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QListWidget) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QListWidget_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QListWidget) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QListWidget_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QListWidget) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QListWidget_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QListWidget) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QListWidget_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListWidget) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QListWidget_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListWidget) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QListWidget_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListWidget) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListWidget) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QListWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidget) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QListWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QListWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QListWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QListWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QListWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QListWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QListWidget_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListWidget) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QListWidget_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QListWidget) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QListWidget_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QListWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QListWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QListWidget) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QListWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QListWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QListWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QListWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QListWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QListWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QListWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QListWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QListWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QListWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QListWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QListWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QListWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QListWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QListWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QListWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QListWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QListWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QListWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QListWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QListWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QListWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QListWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QListWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QListWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QListWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QListWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QListWidget) OnItemPressed(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemClicked(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemDoubleClicked(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemActivated(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemEntered(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemChanged(fn func(*QListWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(item))) })
	pg_QListWidget_connect_itemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnCurrentItemChanged(fn func(*QListWidgetItem, *QListWidgetItem)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQListWidgetItem(unsafe.Pointer(current)), UnsafeNewQListWidgetItem(unsafe.Pointer(previous))) })
	pg_QListWidget_connect_currentItemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnCurrentTextChanged(fn func(string)) int64 {
	id := registerCallback(func(currentText uintptr, currentText_len int64) { fn(pgStrNoFree(currentText, currentText_len)) })
	pg_QListWidget_connect_currentTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnCurrentRowChanged(fn func(int)) int64 {
	id := registerCallback(func(currentRow int32) { fn(int(currentRow)) })
	pg_QListWidget_connect_currentRowChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnItemSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QListWidget_connect_itemSelectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnIndexesMoved(fn func([]*QModelIndex)) int64 {
	id := registerCallback(func(indexes uintptr, indexes_len int64) { fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()) })
	pg_QListWidget_connect_indexesMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListWidget_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListWidget_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListWidget_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListWidget_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListWidget_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QListWidget_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QListWidget_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QListWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QListWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QListWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QListWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QListWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QListWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QListWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListWidget) Disconnect(handle int64) {
	pg_QListWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QListWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -4689322347236366780:
		return 1
	case 9107058682754591716:
		return 2
	case 1237236456312824991:
		return 3
	case 5464467312394730418:
		return 4
	case -9049994684305682212:
		return 5
	case -7784642176186821205:
		return 6
	case -1267696851865372997:
		return 7
	case -4832340528568269947:
		return 8
	case -7252823847285716802:
		return 9
	case -1888629738259754754:
		return 10
	case 1086335023529244512:
		return 11
	case -1574077111192992763:
		return 12
	case -6118001150126564745:
		return 13
	case 4472434241583495620:
		return 14
	case -8151049665126058157:
		return 15
	case -4255494004099806581:
		return 16
	case -4415920049297795213:
		return 17
	case 1919988340934511316:
		return 18
	case -5409927581812374498:
		return 19
	case -4302917285655014661:
		return 20
	case -1781545137352745978:
		return 21
	case -3139904158373725006:
		return 22
	case -8112366604639385491:
		return 23
	case -7450787277202094556:
		return 24
	case 5290442163167458465:
		return 25
	case -1500310024801917802:
		return 26
	case 5218707810677921016:
		return 27
	case -158523882194220890:
		return 28
	case -1539266838375037305:
		return 29
	case 7063795720810552074:
		return 30
	case 740869982805721908:
		return 31
	case 1022771470898460576:
		return 32
	case 6057989667089378859:
		return 33
	case -5892205462145366901:
		return 34
	case -6978583972112316034:
		return 35
	case 503579067252048239:
		return 36
	case -5041452885022114200:
		return 37
	case -4084029023708556876:
		return 38
	case -3666792239577137638:
		return 39
	case -3811438228839363824:
		return 40
	case 7345853097340709638:
		return 41
	case 8758113809152574300:
		return 42
	case 9029941737538200621:
		return 43
	case -8726300139121449691:
		return 44
	case -1155555122391250623:
		return 45
	case -785551405282638793:
		return 46
	case 6366953527806415524:
		return 47
	case 1786430177580376614:
		return 48
	case 8648484264628816806:
		return 49
	case -6021383853323633948:
		return 50
	case 6979418566323745651:
		return 51
	case 4206689019978932209:
		return 52
	case -2447490284072398302:
		return 53
	case 5898328517505784191:
		return 54
	case -1971465653061861519:
		return 55
	case -332090898078543319:
		return 56
	case -4029073905697023887:
		return 57
	case 206122715919917589:
		return 58
	case 975043767033552552:
		return 59
	case -385596010576048993:
		return 60
	case 1023791977679793185:
		return 61
	case 5438819661942843222:
		return 62
	case 5217623514630858297:
		return 63
	case 7049264599837532491:
		return 64
	case 729690261457684101:
		return 65
	case 4531996381736730688:
		return 66
	case -6698043222227006015:
		return 67
	case 4959376063826165943:
		return 68
	case -488693769771308531:
		return 69
	case -8659679893868524044:
		return 70
	case 2110182719375245575:
		return 71
	case 4311534060495737339:
		return 72
	case -8266762397995423915:
		return 73
	case 6398007404756346381:
		return 74
	case 772598475950050299:
		return 75
	case 6316270346668237895:
		return 76
	case -551070009623711640:
		return 77
	case -9161949150489525396:
		return 78
	case -1989937050019464783:
		return 79
	case -212068866860701141:
		return 80
	case 1697444140088555581:
		return 81
	case -6825369174751916:
		return 82
	case -199718707876719129:
		return 83
	case -2039166586246443706:
		return 84
	case 427629649267311777:
		return 85
	case 6477500296980050210:
		return 86
	case 4127240633310655048:
		return 87
	case 5792771328741774991:
		return 88
	case -1185315367526857833:
		return 89
	case 3769868229123868474:
		return 90
	case 7379029982170503008:
		return 91
	case -568369772982215962:
		return 92
	}
	return -1
}

func QListWidget_enable(h uintptr, idx int32) { pg_QListWidget__set_override(h, idx) }

func (this *QListWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QListWidget) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QListWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QListWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QListWidget) OnMimeTypes(fn func() []string) {
	registerVirt(this.UnsafeGoID(), 5464467312394730418, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "mimeTypes"
}

func (this *QListWidget) OnDropMimeData(fn func(int, *QMimeData, DropAction) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(index int32, data uintptr, action int32) int32 { return b2i(fn(int(index), UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action))) }) // slot "dropMimeData"
}

func (this *QListWidget) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QListWidget) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QListWidget) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QListWidget) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QListWidget) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QListWidget) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QListWidget) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QListWidget) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QListWidget) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QListWidget) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QListWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QListWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QListWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QListWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QListWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QListWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QListWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QListWidget) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QListWidget) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QListWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QListWidget) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QListWidget) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QListWidget) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QListWidget) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QListWidget) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QListWidget) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QListWidget) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QListWidget) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QListWidget) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QListWidget) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QListWidget) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QListWidget) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QListWidget) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QListWidget) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QListWidget) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QListWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QListWidget) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QListWidget) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QListWidget) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QListWidget) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QListWidget) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QListWidget) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QListWidget) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QListWidget) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QListWidget) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QListWidget) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QListWidget) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QListWidget) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QListWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QListWidget) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QListWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QListWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QListWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QListWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QListWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QListWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QListWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QListWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QListWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QListWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QListWidget) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QListWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QListWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QListWidget) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QListWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QListWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QListWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QListWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QListWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QListWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QListWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QListWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QListWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QListWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QListWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QListWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QListWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QListWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QListWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QListWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QListWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QListWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QListWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QListWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QListWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QListWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QListWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QListWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QListWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
