package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTableWidget_chain uintptr
	addr_QTableWidget_new uintptr
	addr_QTableWidget_new2 uintptr
	addr_QTableWidget_new3 uintptr
	addr_QTableWidget_new4 uintptr
	addr_QTableWidget_tr uintptr
	addr_QTableWidget_setRowCount uintptr
	addr_QTableWidget_rowCount uintptr
	addr_QTableWidget_setColumnCount uintptr
	addr_QTableWidget_columnCount uintptr
	addr_QTableWidget_row uintptr
	addr_QTableWidget_column uintptr
	addr_QTableWidget_item uintptr
	addr_QTableWidget_setItem uintptr
	addr_QTableWidget_takeItem uintptr
	addr_QTableWidget_items uintptr
	addr_QTableWidget_indexFromItem uintptr
	addr_QTableWidget_itemFromIndex uintptr
	addr_QTableWidget_verticalHeaderItem uintptr
	addr_QTableWidget_setVerticalHeaderItem uintptr
	addr_QTableWidget_takeVerticalHeaderItem uintptr
	addr_QTableWidget_horizontalHeaderItem uintptr
	addr_QTableWidget_setHorizontalHeaderItem uintptr
	addr_QTableWidget_takeHorizontalHeaderItem uintptr
	addr_QTableWidget_setVerticalHeaderLabels uintptr
	addr_QTableWidget_setHorizontalHeaderLabels uintptr
	addr_QTableWidget_currentRow uintptr
	addr_QTableWidget_currentColumn uintptr
	addr_QTableWidget_currentItem uintptr
	addr_QTableWidget_setCurrentItem uintptr
	addr_QTableWidget_setCurrentItem2 uintptr
	addr_QTableWidget_setCurrentCell uintptr
	addr_QTableWidget_setCurrentCell2 uintptr
	addr_QTableWidget_sortItems uintptr
	addr_QTableWidget_setSortingEnabled uintptr
	addr_QTableWidget_isSortingEnabled uintptr
	addr_QTableWidget_editItem uintptr
	addr_QTableWidget_openPersistentEditor uintptr
	addr_QTableWidget_closePersistentEditor uintptr
	addr_QTableWidget_isPersistentEditorOpen uintptr
	addr_QTableWidget_cellWidget uintptr
	addr_QTableWidget_setCellWidget uintptr
	addr_QTableWidget_removeCellWidget uintptr
	addr_QTableWidget_setRangeSelected uintptr
	addr_QTableWidget_selectedRanges uintptr
	addr_QTableWidget_selectedItems uintptr
	addr_QTableWidget_findItems uintptr
	addr_QTableWidget_visualRow uintptr
	addr_QTableWidget_visualColumn uintptr
	addr_QTableWidget_itemAt uintptr
	addr_QTableWidget_itemAt2 uintptr
	addr_QTableWidget_visualItemRect uintptr
	addr_QTableWidget_itemPrototype uintptr
	addr_QTableWidget_setItemPrototype uintptr
	addr_QTableWidget_supportedDragActions uintptr
	addr_QTableWidget_setSupportedDragActions uintptr
	addr_QTableWidget_scrollToItem uintptr
	addr_QTableWidget_insertRow uintptr
	addr_QTableWidget_insertColumn uintptr
	addr_QTableWidget_removeRow uintptr
	addr_QTableWidget_removeColumn uintptr
	addr_QTableWidget_clear uintptr
	addr_QTableWidget_clearContents uintptr
	addr_QTableWidget_itemPressed uintptr
	addr_QTableWidget_itemClicked uintptr
	addr_QTableWidget_itemDoubleClicked uintptr
	addr_QTableWidget_itemActivated uintptr
	addr_QTableWidget_itemEntered uintptr
	addr_QTableWidget_itemChanged uintptr
	addr_QTableWidget_currentItemChanged uintptr
	addr_QTableWidget_itemSelectionChanged uintptr
	addr_QTableWidget_cellPressed uintptr
	addr_QTableWidget_cellClicked uintptr
	addr_QTableWidget_cellDoubleClicked uintptr
	addr_QTableWidget_cellActivated uintptr
	addr_QTableWidget_cellEntered uintptr
	addr_QTableWidget_cellChanged uintptr
	addr_QTableWidget_currentCellChanged uintptr
	addr_QTableWidget_rowMoved uintptr
	addr_QTableWidget_columnMoved uintptr
	addr_QTableWidget_rowResized uintptr
	addr_QTableWidget_columnResized uintptr
	addr_QTableWidget_rowCountChanged uintptr
	addr_QTableWidget_columnCountChanged uintptr
	addr_QTableWidget_state uintptr
	addr_QTableWidget_scheduleDelayedItemsLayout uintptr
	addr_QTableWidget_executeDelayedItemsLayout uintptr
	addr_QTableWidget_setDirtyRegion uintptr
	addr_QTableWidget_scrollDirtyRegion uintptr
	addr_QTableWidget_dirtyRegionOffset uintptr
	addr_QTableWidget_startAutoScroll uintptr
	addr_QTableWidget_stopAutoScroll uintptr
	addr_QTableWidget_doAutoScroll uintptr
	addr_QTableWidget_dropIndicatorPosition uintptr
	addr_QTableWidget_setViewportMargins uintptr
	addr_QTableWidget_setViewportMargins2 uintptr
	addr_QTableWidget_viewportMargins uintptr
	addr_QTableWidget_drawFrame uintptr
	addr_QTableWidget_setTabOrder uintptr
	addr_QTableWidget_mouseGrabber uintptr
	addr_QTableWidget_keyboardGrabber uintptr
	addr_QTableWidget_find uintptr
	addr_QTableWidget_createWindowContainer uintptr
	addr_QTableWidget_updateMicroFocus uintptr
	addr_QTableWidget_create uintptr
	addr_QTableWidget_destroy uintptr
	addr_QTableWidget_focusNextChild uintptr
	addr_QTableWidget_focusPreviousChild uintptr
	addr_QTableWidget_connect uintptr
	addr_QTableWidget_connect2 uintptr
	addr_QTableWidget_sender uintptr
	addr_QTableWidget_senderSignalIndex uintptr
	addr_QTableWidget_receivers uintptr
	addr_QTableWidget_isSignalConnected uintptr
	pg_QTableWidget_devicePixelRatioFScale func() float64
	pg_QTableWidget_encodeMetricF func(int32, float64) int32
	pg_QTableWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTableWidget_tr2 uintptr
	addr_QTableWidget_tr3 uintptr
	addr_QTableWidget_sortItems2 uintptr
	addr_QTableWidget_scrollToItem2 uintptr
	addr_QTableWidget_createWindowContainer2 uintptr
	addr_QTableWidget_createWindowContainer3 uintptr
	addr_QTableWidget_updateMicroFocus2 uintptr
	addr_QTableWidget_create2 uintptr
	addr_QTableWidget_create3 uintptr
	addr_QTableWidget_create4 uintptr
	addr_QTableWidget_destroy2 uintptr
	addr_QTableWidget_destroy3 uintptr
	addr_QTableWidget_connect3 uintptr
	addr_QTableWidget_connect4 uintptr
	addr_QTableWidget_metaObject_default uintptr
	addr_QTableWidget_event_default uintptr
	addr_QTableWidget_mimeTypes_default uintptr
	addr_QTableWidget_dropMimeData_default uintptr
	addr_QTableWidget_supportedDropActions_default uintptr
	addr_QTableWidget_dropEvent_default uintptr
	addr_QTableWidget_setRootIndex_default uintptr
	addr_QTableWidget_setSelectionModel_default uintptr
	addr_QTableWidget_doItemsLayout_default uintptr
	addr_QTableWidget_visualRect_default uintptr
	addr_QTableWidget_scrollTo_default uintptr
	addr_QTableWidget_indexAt_default uintptr
	addr_QTableWidget_scrollContentsBy_default uintptr
	addr_QTableWidget_initViewItemOption_default uintptr
	addr_QTableWidget_paintEvent_default uintptr
	addr_QTableWidget_timerEvent_default uintptr
	addr_QTableWidget_horizontalOffset_default uintptr
	addr_QTableWidget_verticalOffset_default uintptr
	addr_QTableWidget_moveCursor_default uintptr
	addr_QTableWidget_setSelection_default uintptr
	addr_QTableWidget_visualRegionForSelection_default uintptr
	addr_QTableWidget_selectedIndexes_default uintptr
	addr_QTableWidget_updateGeometries_default uintptr
	addr_QTableWidget_viewportSizeHint_default uintptr
	addr_QTableWidget_sizeHintForRow_default uintptr
	addr_QTableWidget_sizeHintForColumn_default uintptr
	addr_QTableWidget_verticalScrollbarAction_default uintptr
	addr_QTableWidget_horizontalScrollbarAction_default uintptr
	addr_QTableWidget_isIndexHidden_default uintptr
	addr_QTableWidget_selectionChanged_default uintptr
	addr_QTableWidget_currentChanged_default uintptr
	addr_QTableWidget_keyboardSearch_default uintptr
	addr_QTableWidget_itemDelegateForIndex_default uintptr
	addr_QTableWidget_inputMethodQuery_default uintptr
	addr_QTableWidget_reset_default uintptr
	addr_QTableWidget_selectAll_default uintptr
	addr_QTableWidget_rowsInserted_default uintptr
	addr_QTableWidget_rowsAboutToBeRemoved_default uintptr
	addr_QTableWidget_updateEditorData_default uintptr
	addr_QTableWidget_updateEditorGeometries_default uintptr
	addr_QTableWidget_verticalScrollbarValueChanged_default uintptr
	addr_QTableWidget_horizontalScrollbarValueChanged_default uintptr
	addr_QTableWidget_closeEditor_default uintptr
	addr_QTableWidget_commitData_default uintptr
	addr_QTableWidget_editorDestroyed_default uintptr
	addr_QTableWidget_edit_default uintptr
	addr_QTableWidget_selectionCommand_default uintptr
	addr_QTableWidget_startDrag_default uintptr
	addr_QTableWidget_focusNextPrevChild_default uintptr
	addr_QTableWidget_viewportEvent_default uintptr
	addr_QTableWidget_mousePressEvent_default uintptr
	addr_QTableWidget_mouseMoveEvent_default uintptr
	addr_QTableWidget_mouseReleaseEvent_default uintptr
	addr_QTableWidget_mouseDoubleClickEvent_default uintptr
	addr_QTableWidget_dragEnterEvent_default uintptr
	addr_QTableWidget_dragMoveEvent_default uintptr
	addr_QTableWidget_dragLeaveEvent_default uintptr
	addr_QTableWidget_focusInEvent_default uintptr
	addr_QTableWidget_focusOutEvent_default uintptr
	addr_QTableWidget_keyPressEvent_default uintptr
	addr_QTableWidget_resizeEvent_default uintptr
	addr_QTableWidget_inputMethodEvent_default uintptr
	addr_QTableWidget_eventFilter_default uintptr
	addr_QTableWidget_minimumSizeHint_default uintptr
	addr_QTableWidget_sizeHint_default uintptr
	addr_QTableWidget_setupViewport_default uintptr
	addr_QTableWidget_wheelEvent_default uintptr
	addr_QTableWidget_contextMenuEvent_default uintptr
	addr_QTableWidget_changeEvent_default uintptr
	addr_QTableWidget_initStyleOption_default uintptr
	addr_QTableWidget_devType_default uintptr
	addr_QTableWidget_setVisible_default uintptr
	addr_QTableWidget_heightForWidth_default uintptr
	addr_QTableWidget_hasHeightForWidth_default uintptr
	addr_QTableWidget_paintEngine_default uintptr
	addr_QTableWidget_keyReleaseEvent_default uintptr
	addr_QTableWidget_enterEvent_default uintptr
	addr_QTableWidget_leaveEvent_default uintptr
	addr_QTableWidget_moveEvent_default uintptr
	addr_QTableWidget_closeEvent_default uintptr
	addr_QTableWidget_tabletEvent_default uintptr
	addr_QTableWidget_actionEvent_default uintptr
	addr_QTableWidget_showEvent_default uintptr
	addr_QTableWidget_hideEvent_default uintptr
	addr_QTableWidget_metric_default uintptr
	addr_QTableWidget_initPainter_default uintptr
	addr_QTableWidget_redirected_default uintptr
	addr_QTableWidget_sharedPainter_default uintptr
	addr_QTableWidget_childEvent_default uintptr
	addr_QTableWidget_customEvent_default uintptr
	addr_QTableWidget_connectNotify_default uintptr
	addr_QTableWidget_disconnectNotify_default uintptr
	pg_QTableWidget_connect_itemPressed func(uintptr, int64)
	pg_QTableWidget_connect_itemClicked func(uintptr, int64)
	pg_QTableWidget_connect_itemDoubleClicked func(uintptr, int64)
	pg_QTableWidget_connect_itemActivated func(uintptr, int64)
	pg_QTableWidget_connect_itemEntered func(uintptr, int64)
	pg_QTableWidget_connect_itemChanged func(uintptr, int64)
	pg_QTableWidget_connect_currentItemChanged func(uintptr, int64)
	pg_QTableWidget_connect_itemSelectionChanged func(uintptr, int64)
	pg_QTableWidget_connect_cellPressed func(uintptr, int64)
	pg_QTableWidget_connect_cellClicked func(uintptr, int64)
	pg_QTableWidget_connect_cellDoubleClicked func(uintptr, int64)
	pg_QTableWidget_connect_cellActivated func(uintptr, int64)
	pg_QTableWidget_connect_cellEntered func(uintptr, int64)
	pg_QTableWidget_connect_cellChanged func(uintptr, int64)
	pg_QTableWidget_connect_currentCellChanged func(uintptr, int64)
	pg_QTableWidget_connect_pressed func(uintptr, int64)
	pg_QTableWidget_connect_clicked func(uintptr, int64)
	pg_QTableWidget_connect_doubleClicked func(uintptr, int64)
	pg_QTableWidget_connect_activated func(uintptr, int64)
	pg_QTableWidget_connect_entered func(uintptr, int64)
	pg_QTableWidget_connect_viewportEntered func(uintptr, int64)
	pg_QTableWidget_connect_iconSizeChanged func(uintptr, int64)
	pg_QTableWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QTableWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QTableWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTableWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTableWidget_connect_destroyed func(uintptr, int64)
	pg_QTableWidget_connect_destroyed2 func(uintptr, int64)
	pg_QTableWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QTableWidget__disconnect func(int64)
	pg_QTableWidget__set_override func(uintptr, int32)
	pg_QTableWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTableWidget_chain, "QTableWidget_chain")
	symaddr(&addr_QTableWidget_new, "QTableWidget_new")
	symaddr(&addr_QTableWidget_new2, "QTableWidget_new2")
	symaddr(&addr_QTableWidget_new3, "QTableWidget_new3")
	symaddr(&addr_QTableWidget_new4, "QTableWidget_new4")
	symaddr(&addr_QTableWidget_tr, "QTableWidget_tr")
	symaddr(&addr_QTableWidget_setRowCount, "QTableWidget_setRowCount")
	symaddr(&addr_QTableWidget_rowCount, "QTableWidget_rowCount")
	symaddr(&addr_QTableWidget_setColumnCount, "QTableWidget_setColumnCount")
	symaddr(&addr_QTableWidget_columnCount, "QTableWidget_columnCount")
	symaddr(&addr_QTableWidget_row, "QTableWidget_row")
	symaddr(&addr_QTableWidget_column, "QTableWidget_column")
	symaddr(&addr_QTableWidget_item, "QTableWidget_item")
	symaddr(&addr_QTableWidget_setItem, "QTableWidget_setItem")
	symaddr(&addr_QTableWidget_takeItem, "QTableWidget_takeItem")
	symaddr(&addr_QTableWidget_items, "QTableWidget_items")
	symaddr(&addr_QTableWidget_indexFromItem, "QTableWidget_indexFromItem")
	symaddr(&addr_QTableWidget_itemFromIndex, "QTableWidget_itemFromIndex")
	symaddr(&addr_QTableWidget_verticalHeaderItem, "QTableWidget_verticalHeaderItem")
	symaddr(&addr_QTableWidget_setVerticalHeaderItem, "QTableWidget_setVerticalHeaderItem")
	symaddr(&addr_QTableWidget_takeVerticalHeaderItem, "QTableWidget_takeVerticalHeaderItem")
	symaddr(&addr_QTableWidget_horizontalHeaderItem, "QTableWidget_horizontalHeaderItem")
	symaddr(&addr_QTableWidget_setHorizontalHeaderItem, "QTableWidget_setHorizontalHeaderItem")
	symaddr(&addr_QTableWidget_takeHorizontalHeaderItem, "QTableWidget_takeHorizontalHeaderItem")
	symaddr(&addr_QTableWidget_setVerticalHeaderLabels, "QTableWidget_setVerticalHeaderLabels")
	symaddr(&addr_QTableWidget_setHorizontalHeaderLabels, "QTableWidget_setHorizontalHeaderLabels")
	symaddr(&addr_QTableWidget_currentRow, "QTableWidget_currentRow")
	symaddr(&addr_QTableWidget_currentColumn, "QTableWidget_currentColumn")
	symaddr(&addr_QTableWidget_currentItem, "QTableWidget_currentItem")
	symaddr(&addr_QTableWidget_setCurrentItem, "QTableWidget_setCurrentItem")
	symaddr(&addr_QTableWidget_setCurrentItem2, "QTableWidget_setCurrentItem2")
	symaddr(&addr_QTableWidget_setCurrentCell, "QTableWidget_setCurrentCell")
	symaddr(&addr_QTableWidget_setCurrentCell2, "QTableWidget_setCurrentCell2")
	symaddr(&addr_QTableWidget_sortItems, "QTableWidget_sortItems")
	symaddr(&addr_QTableWidget_setSortingEnabled, "QTableWidget_setSortingEnabled")
	symaddr(&addr_QTableWidget_isSortingEnabled, "QTableWidget_isSortingEnabled")
	symaddr(&addr_QTableWidget_editItem, "QTableWidget_editItem")
	symaddr(&addr_QTableWidget_openPersistentEditor, "QTableWidget_openPersistentEditor")
	symaddr(&addr_QTableWidget_closePersistentEditor, "QTableWidget_closePersistentEditor")
	symaddr(&addr_QTableWidget_isPersistentEditorOpen, "QTableWidget_isPersistentEditorOpen")
	symaddr(&addr_QTableWidget_cellWidget, "QTableWidget_cellWidget")
	symaddr(&addr_QTableWidget_setCellWidget, "QTableWidget_setCellWidget")
	symaddr(&addr_QTableWidget_removeCellWidget, "QTableWidget_removeCellWidget")
	symaddr(&addr_QTableWidget_setRangeSelected, "QTableWidget_setRangeSelected")
	symaddr(&addr_QTableWidget_selectedRanges, "QTableWidget_selectedRanges")
	symaddr(&addr_QTableWidget_selectedItems, "QTableWidget_selectedItems")
	symaddr(&addr_QTableWidget_findItems, "QTableWidget_findItems")
	symaddr(&addr_QTableWidget_visualRow, "QTableWidget_visualRow")
	symaddr(&addr_QTableWidget_visualColumn, "QTableWidget_visualColumn")
	symaddr(&addr_QTableWidget_itemAt, "QTableWidget_itemAt")
	symaddr(&addr_QTableWidget_itemAt2, "QTableWidget_itemAt2")
	symaddr(&addr_QTableWidget_visualItemRect, "QTableWidget_visualItemRect")
	symaddr(&addr_QTableWidget_itemPrototype, "QTableWidget_itemPrototype")
	symaddr(&addr_QTableWidget_setItemPrototype, "QTableWidget_setItemPrototype")
	symaddr(&addr_QTableWidget_supportedDragActions, "QTableWidget_supportedDragActions")
	symaddr(&addr_QTableWidget_setSupportedDragActions, "QTableWidget_setSupportedDragActions")
	symaddr(&addr_QTableWidget_scrollToItem, "QTableWidget_scrollToItem")
	symaddr(&addr_QTableWidget_insertRow, "QTableWidget_insertRow")
	symaddr(&addr_QTableWidget_insertColumn, "QTableWidget_insertColumn")
	symaddr(&addr_QTableWidget_removeRow, "QTableWidget_removeRow")
	symaddr(&addr_QTableWidget_removeColumn, "QTableWidget_removeColumn")
	symaddr(&addr_QTableWidget_clear, "QTableWidget_clear")
	symaddr(&addr_QTableWidget_clearContents, "QTableWidget_clearContents")
	symaddr(&addr_QTableWidget_itemPressed, "QTableWidget_itemPressed")
	symaddr(&addr_QTableWidget_itemClicked, "QTableWidget_itemClicked")
	symaddr(&addr_QTableWidget_itemDoubleClicked, "QTableWidget_itemDoubleClicked")
	symaddr(&addr_QTableWidget_itemActivated, "QTableWidget_itemActivated")
	symaddr(&addr_QTableWidget_itemEntered, "QTableWidget_itemEntered")
	symaddr(&addr_QTableWidget_itemChanged, "QTableWidget_itemChanged")
	symaddr(&addr_QTableWidget_currentItemChanged, "QTableWidget_currentItemChanged")
	symaddr(&addr_QTableWidget_itemSelectionChanged, "QTableWidget_itemSelectionChanged")
	symaddr(&addr_QTableWidget_cellPressed, "QTableWidget_cellPressed")
	symaddr(&addr_QTableWidget_cellClicked, "QTableWidget_cellClicked")
	symaddr(&addr_QTableWidget_cellDoubleClicked, "QTableWidget_cellDoubleClicked")
	symaddr(&addr_QTableWidget_cellActivated, "QTableWidget_cellActivated")
	symaddr(&addr_QTableWidget_cellEntered, "QTableWidget_cellEntered")
	symaddr(&addr_QTableWidget_cellChanged, "QTableWidget_cellChanged")
	symaddr(&addr_QTableWidget_currentCellChanged, "QTableWidget_currentCellChanged")
	symaddr(&addr_QTableWidget_rowMoved, "QTableWidget_rowMoved")
	symaddr(&addr_QTableWidget_columnMoved, "QTableWidget_columnMoved")
	symaddr(&addr_QTableWidget_rowResized, "QTableWidget_rowResized")
	symaddr(&addr_QTableWidget_columnResized, "QTableWidget_columnResized")
	symaddr(&addr_QTableWidget_rowCountChanged, "QTableWidget_rowCountChanged")
	symaddr(&addr_QTableWidget_columnCountChanged, "QTableWidget_columnCountChanged")
	symaddr(&addr_QTableWidget_state, "QTableWidget_state")
	symaddr(&addr_QTableWidget_scheduleDelayedItemsLayout, "QTableWidget_scheduleDelayedItemsLayout")
	symaddr(&addr_QTableWidget_executeDelayedItemsLayout, "QTableWidget_executeDelayedItemsLayout")
	symaddr(&addr_QTableWidget_setDirtyRegion, "QTableWidget_setDirtyRegion")
	symaddr(&addr_QTableWidget_scrollDirtyRegion, "QTableWidget_scrollDirtyRegion")
	symaddr(&addr_QTableWidget_dirtyRegionOffset, "QTableWidget_dirtyRegionOffset")
	symaddr(&addr_QTableWidget_startAutoScroll, "QTableWidget_startAutoScroll")
	symaddr(&addr_QTableWidget_stopAutoScroll, "QTableWidget_stopAutoScroll")
	symaddr(&addr_QTableWidget_doAutoScroll, "QTableWidget_doAutoScroll")
	symaddr(&addr_QTableWidget_dropIndicatorPosition, "QTableWidget_dropIndicatorPosition")
	symaddr(&addr_QTableWidget_setViewportMargins, "QTableWidget_setViewportMargins")
	symaddr(&addr_QTableWidget_setViewportMargins2, "QTableWidget_setViewportMargins2")
	symaddr(&addr_QTableWidget_viewportMargins, "QTableWidget_viewportMargins")
	symaddr(&addr_QTableWidget_drawFrame, "QTableWidget_drawFrame")
	symaddr(&addr_QTableWidget_setTabOrder, "QTableWidget_setTabOrder")
	symaddr(&addr_QTableWidget_mouseGrabber, "QTableWidget_mouseGrabber")
	symaddr(&addr_QTableWidget_keyboardGrabber, "QTableWidget_keyboardGrabber")
	symaddr(&addr_QTableWidget_find, "QTableWidget_find")
	symaddr(&addr_QTableWidget_createWindowContainer, "QTableWidget_createWindowContainer")
	symaddr(&addr_QTableWidget_updateMicroFocus, "QTableWidget_updateMicroFocus")
	symaddr(&addr_QTableWidget_create, "QTableWidget_create")
	symaddr(&addr_QTableWidget_destroy, "QTableWidget_destroy")
	symaddr(&addr_QTableWidget_focusNextChild, "QTableWidget_focusNextChild")
	symaddr(&addr_QTableWidget_focusPreviousChild, "QTableWidget_focusPreviousChild")
	symaddr(&addr_QTableWidget_connect, "QTableWidget_connect")
	symaddr(&addr_QTableWidget_connect2, "QTableWidget_connect2")
	symaddr(&addr_QTableWidget_sender, "QTableWidget_sender")
	symaddr(&addr_QTableWidget_senderSignalIndex, "QTableWidget_senderSignalIndex")
	symaddr(&addr_QTableWidget_receivers, "QTableWidget_receivers")
	symaddr(&addr_QTableWidget_isSignalConnected, "QTableWidget_isSignalConnected")
	reg(&pg_QTableWidget_devicePixelRatioFScale, "QTableWidget_devicePixelRatioFScale")
	reg(&pg_QTableWidget_encodeMetricF, "QTableWidget_encodeMetricF")
	reg(&pg_QTableWidget_getDecodedMetricF, "QTableWidget_getDecodedMetricF")
	symaddr(&addr_QTableWidget_tr2, "QTableWidget_tr2")
	symaddr(&addr_QTableWidget_tr3, "QTableWidget_tr3")
	symaddr(&addr_QTableWidget_sortItems2, "QTableWidget_sortItems2")
	symaddr(&addr_QTableWidget_scrollToItem2, "QTableWidget_scrollToItem2")
	symaddr(&addr_QTableWidget_createWindowContainer2, "QTableWidget_createWindowContainer2")
	symaddr(&addr_QTableWidget_createWindowContainer3, "QTableWidget_createWindowContainer3")
	symaddr(&addr_QTableWidget_updateMicroFocus2, "QTableWidget_updateMicroFocus2")
	symaddr(&addr_QTableWidget_create2, "QTableWidget_create2")
	symaddr(&addr_QTableWidget_create3, "QTableWidget_create3")
	symaddr(&addr_QTableWidget_create4, "QTableWidget_create4")
	symaddr(&addr_QTableWidget_destroy2, "QTableWidget_destroy2")
	symaddr(&addr_QTableWidget_destroy3, "QTableWidget_destroy3")
	symaddr(&addr_QTableWidget_connect3, "QTableWidget_connect3")
	symaddr(&addr_QTableWidget_connect4, "QTableWidget_connect4")
	symaddr(&addr_QTableWidget_metaObject_default, "QTableWidget_metaObject_default")
	symaddr(&addr_QTableWidget_event_default, "QTableWidget_event_default")
	symaddr(&addr_QTableWidget_mimeTypes_default, "QTableWidget_mimeTypes_default")
	symaddr(&addr_QTableWidget_dropMimeData_default, "QTableWidget_dropMimeData_default")
	symaddr(&addr_QTableWidget_supportedDropActions_default, "QTableWidget_supportedDropActions_default")
	symaddr(&addr_QTableWidget_dropEvent_default, "QTableWidget_dropEvent_default")
	symaddr(&addr_QTableWidget_setRootIndex_default, "QTableWidget_setRootIndex_default")
	symaddr(&addr_QTableWidget_setSelectionModel_default, "QTableWidget_setSelectionModel_default")
	symaddr(&addr_QTableWidget_doItemsLayout_default, "QTableWidget_doItemsLayout_default")
	symaddr(&addr_QTableWidget_visualRect_default, "QTableWidget_visualRect_default")
	symaddr(&addr_QTableWidget_scrollTo_default, "QTableWidget_scrollTo_default")
	symaddr(&addr_QTableWidget_indexAt_default, "QTableWidget_indexAt_default")
	symaddr(&addr_QTableWidget_scrollContentsBy_default, "QTableWidget_scrollContentsBy_default")
	symaddr(&addr_QTableWidget_initViewItemOption_default, "QTableWidget_initViewItemOption_default")
	symaddr(&addr_QTableWidget_paintEvent_default, "QTableWidget_paintEvent_default")
	symaddr(&addr_QTableWidget_timerEvent_default, "QTableWidget_timerEvent_default")
	symaddr(&addr_QTableWidget_horizontalOffset_default, "QTableWidget_horizontalOffset_default")
	symaddr(&addr_QTableWidget_verticalOffset_default, "QTableWidget_verticalOffset_default")
	symaddr(&addr_QTableWidget_moveCursor_default, "QTableWidget_moveCursor_default")
	symaddr(&addr_QTableWidget_setSelection_default, "QTableWidget_setSelection_default")
	symaddr(&addr_QTableWidget_visualRegionForSelection_default, "QTableWidget_visualRegionForSelection_default")
	symaddr(&addr_QTableWidget_selectedIndexes_default, "QTableWidget_selectedIndexes_default")
	symaddr(&addr_QTableWidget_updateGeometries_default, "QTableWidget_updateGeometries_default")
	symaddr(&addr_QTableWidget_viewportSizeHint_default, "QTableWidget_viewportSizeHint_default")
	symaddr(&addr_QTableWidget_sizeHintForRow_default, "QTableWidget_sizeHintForRow_default")
	symaddr(&addr_QTableWidget_sizeHintForColumn_default, "QTableWidget_sizeHintForColumn_default")
	symaddr(&addr_QTableWidget_verticalScrollbarAction_default, "QTableWidget_verticalScrollbarAction_default")
	symaddr(&addr_QTableWidget_horizontalScrollbarAction_default, "QTableWidget_horizontalScrollbarAction_default")
	symaddr(&addr_QTableWidget_isIndexHidden_default, "QTableWidget_isIndexHidden_default")
	symaddr(&addr_QTableWidget_selectionChanged_default, "QTableWidget_selectionChanged_default")
	symaddr(&addr_QTableWidget_currentChanged_default, "QTableWidget_currentChanged_default")
	symaddr(&addr_QTableWidget_keyboardSearch_default, "QTableWidget_keyboardSearch_default")
	symaddr(&addr_QTableWidget_itemDelegateForIndex_default, "QTableWidget_itemDelegateForIndex_default")
	symaddr(&addr_QTableWidget_inputMethodQuery_default, "QTableWidget_inputMethodQuery_default")
	symaddr(&addr_QTableWidget_reset_default, "QTableWidget_reset_default")
	symaddr(&addr_QTableWidget_selectAll_default, "QTableWidget_selectAll_default")
	symaddr(&addr_QTableWidget_rowsInserted_default, "QTableWidget_rowsInserted_default")
	symaddr(&addr_QTableWidget_rowsAboutToBeRemoved_default, "QTableWidget_rowsAboutToBeRemoved_default")
	symaddr(&addr_QTableWidget_updateEditorData_default, "QTableWidget_updateEditorData_default")
	symaddr(&addr_QTableWidget_updateEditorGeometries_default, "QTableWidget_updateEditorGeometries_default")
	symaddr(&addr_QTableWidget_verticalScrollbarValueChanged_default, "QTableWidget_verticalScrollbarValueChanged_default")
	symaddr(&addr_QTableWidget_horizontalScrollbarValueChanged_default, "QTableWidget_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QTableWidget_closeEditor_default, "QTableWidget_closeEditor_default")
	symaddr(&addr_QTableWidget_commitData_default, "QTableWidget_commitData_default")
	symaddr(&addr_QTableWidget_editorDestroyed_default, "QTableWidget_editorDestroyed_default")
	symaddr(&addr_QTableWidget_edit_default, "QTableWidget_edit_default")
	symaddr(&addr_QTableWidget_selectionCommand_default, "QTableWidget_selectionCommand_default")
	symaddr(&addr_QTableWidget_startDrag_default, "QTableWidget_startDrag_default")
	symaddr(&addr_QTableWidget_focusNextPrevChild_default, "QTableWidget_focusNextPrevChild_default")
	symaddr(&addr_QTableWidget_viewportEvent_default, "QTableWidget_viewportEvent_default")
	symaddr(&addr_QTableWidget_mousePressEvent_default, "QTableWidget_mousePressEvent_default")
	symaddr(&addr_QTableWidget_mouseMoveEvent_default, "QTableWidget_mouseMoveEvent_default")
	symaddr(&addr_QTableWidget_mouseReleaseEvent_default, "QTableWidget_mouseReleaseEvent_default")
	symaddr(&addr_QTableWidget_mouseDoubleClickEvent_default, "QTableWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QTableWidget_dragEnterEvent_default, "QTableWidget_dragEnterEvent_default")
	symaddr(&addr_QTableWidget_dragMoveEvent_default, "QTableWidget_dragMoveEvent_default")
	symaddr(&addr_QTableWidget_dragLeaveEvent_default, "QTableWidget_dragLeaveEvent_default")
	symaddr(&addr_QTableWidget_focusInEvent_default, "QTableWidget_focusInEvent_default")
	symaddr(&addr_QTableWidget_focusOutEvent_default, "QTableWidget_focusOutEvent_default")
	symaddr(&addr_QTableWidget_keyPressEvent_default, "QTableWidget_keyPressEvent_default")
	symaddr(&addr_QTableWidget_resizeEvent_default, "QTableWidget_resizeEvent_default")
	symaddr(&addr_QTableWidget_inputMethodEvent_default, "QTableWidget_inputMethodEvent_default")
	symaddr(&addr_QTableWidget_eventFilter_default, "QTableWidget_eventFilter_default")
	symaddr(&addr_QTableWidget_minimumSizeHint_default, "QTableWidget_minimumSizeHint_default")
	symaddr(&addr_QTableWidget_sizeHint_default, "QTableWidget_sizeHint_default")
	symaddr(&addr_QTableWidget_setupViewport_default, "QTableWidget_setupViewport_default")
	symaddr(&addr_QTableWidget_wheelEvent_default, "QTableWidget_wheelEvent_default")
	symaddr(&addr_QTableWidget_contextMenuEvent_default, "QTableWidget_contextMenuEvent_default")
	symaddr(&addr_QTableWidget_changeEvent_default, "QTableWidget_changeEvent_default")
	symaddr(&addr_QTableWidget_initStyleOption_default, "QTableWidget_initStyleOption_default")
	symaddr(&addr_QTableWidget_devType_default, "QTableWidget_devType_default")
	symaddr(&addr_QTableWidget_setVisible_default, "QTableWidget_setVisible_default")
	symaddr(&addr_QTableWidget_heightForWidth_default, "QTableWidget_heightForWidth_default")
	symaddr(&addr_QTableWidget_hasHeightForWidth_default, "QTableWidget_hasHeightForWidth_default")
	symaddr(&addr_QTableWidget_paintEngine_default, "QTableWidget_paintEngine_default")
	symaddr(&addr_QTableWidget_keyReleaseEvent_default, "QTableWidget_keyReleaseEvent_default")
	symaddr(&addr_QTableWidget_enterEvent_default, "QTableWidget_enterEvent_default")
	symaddr(&addr_QTableWidget_leaveEvent_default, "QTableWidget_leaveEvent_default")
	symaddr(&addr_QTableWidget_moveEvent_default, "QTableWidget_moveEvent_default")
	symaddr(&addr_QTableWidget_closeEvent_default, "QTableWidget_closeEvent_default")
	symaddr(&addr_QTableWidget_tabletEvent_default, "QTableWidget_tabletEvent_default")
	symaddr(&addr_QTableWidget_actionEvent_default, "QTableWidget_actionEvent_default")
	symaddr(&addr_QTableWidget_showEvent_default, "QTableWidget_showEvent_default")
	symaddr(&addr_QTableWidget_hideEvent_default, "QTableWidget_hideEvent_default")
	symaddr(&addr_QTableWidget_metric_default, "QTableWidget_metric_default")
	symaddr(&addr_QTableWidget_initPainter_default, "QTableWidget_initPainter_default")
	symaddr(&addr_QTableWidget_redirected_default, "QTableWidget_redirected_default")
	symaddr(&addr_QTableWidget_sharedPainter_default, "QTableWidget_sharedPainter_default")
	symaddr(&addr_QTableWidget_childEvent_default, "QTableWidget_childEvent_default")
	symaddr(&addr_QTableWidget_customEvent_default, "QTableWidget_customEvent_default")
	symaddr(&addr_QTableWidget_connectNotify_default, "QTableWidget_connectNotify_default")
	symaddr(&addr_QTableWidget_disconnectNotify_default, "QTableWidget_disconnectNotify_default")
	reg(&pg_QTableWidget_connect_itemPressed, "QTableWidget_connect_itemPressed")
	reg(&pg_QTableWidget_connect_itemClicked, "QTableWidget_connect_itemClicked")
	reg(&pg_QTableWidget_connect_itemDoubleClicked, "QTableWidget_connect_itemDoubleClicked")
	reg(&pg_QTableWidget_connect_itemActivated, "QTableWidget_connect_itemActivated")
	reg(&pg_QTableWidget_connect_itemEntered, "QTableWidget_connect_itemEntered")
	reg(&pg_QTableWidget_connect_itemChanged, "QTableWidget_connect_itemChanged")
	reg(&pg_QTableWidget_connect_currentItemChanged, "QTableWidget_connect_currentItemChanged")
	reg(&pg_QTableWidget_connect_itemSelectionChanged, "QTableWidget_connect_itemSelectionChanged")
	reg(&pg_QTableWidget_connect_cellPressed, "QTableWidget_connect_cellPressed")
	reg(&pg_QTableWidget_connect_cellClicked, "QTableWidget_connect_cellClicked")
	reg(&pg_QTableWidget_connect_cellDoubleClicked, "QTableWidget_connect_cellDoubleClicked")
	reg(&pg_QTableWidget_connect_cellActivated, "QTableWidget_connect_cellActivated")
	reg(&pg_QTableWidget_connect_cellEntered, "QTableWidget_connect_cellEntered")
	reg(&pg_QTableWidget_connect_cellChanged, "QTableWidget_connect_cellChanged")
	reg(&pg_QTableWidget_connect_currentCellChanged, "QTableWidget_connect_currentCellChanged")
	reg(&pg_QTableWidget_connect_pressed, "QTableWidget_connect_pressed")
	reg(&pg_QTableWidget_connect_clicked, "QTableWidget_connect_clicked")
	reg(&pg_QTableWidget_connect_doubleClicked, "QTableWidget_connect_doubleClicked")
	reg(&pg_QTableWidget_connect_activated, "QTableWidget_connect_activated")
	reg(&pg_QTableWidget_connect_entered, "QTableWidget_connect_entered")
	reg(&pg_QTableWidget_connect_viewportEntered, "QTableWidget_connect_viewportEntered")
	reg(&pg_QTableWidget_connect_iconSizeChanged, "QTableWidget_connect_iconSizeChanged")
	reg(&pg_QTableWidget_connect_windowTitleChanged, "QTableWidget_connect_windowTitleChanged")
	reg(&pg_QTableWidget_connect_windowIconChanged, "QTableWidget_connect_windowIconChanged")
	reg(&pg_QTableWidget_connect_windowIconTextChanged, "QTableWidget_connect_windowIconTextChanged")
	reg(&pg_QTableWidget_connect_customContextMenuRequested, "QTableWidget_connect_customContextMenuRequested")
	reg(&pg_QTableWidget_connect_destroyed, "QTableWidget_connect_destroyed")
	reg(&pg_QTableWidget_connect_destroyed2, "QTableWidget_connect_destroyed2")
	reg(&pg_QTableWidget_connect_objectNameChanged, "QTableWidget_connect_objectNameChanged")
	reg(&pg_QTableWidget__disconnect, "QTableWidget__disconnect")
	reg(&pg_QTableWidget__set_override, "QTableWidget__set_override")
	reg(&pg_QTableWidget_delete, "QTableWidget_delete")
}

type QTableWidget struct {
	h uintptr
	*QTableView
}

func bindQTableWidget(h uintptr) *QTableWidget {
	if h == 0 {
		return nil
	}
	var __c [6]uintptr
	purego.SyscallN(addr_QTableWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTableWidget{h: h, QTableView: &QTableView{h: __c[0], QAbstractItemView: &QAbstractItemView{h: __c[1], QAbstractScrollArea: &QAbstractScrollArea{h: __c[2], QFrame: &QFrame{h: __c[3], QWidget: &QWidget{h: __c[4], QObject: &QObject{h: __c[5]}}}}}}}
}

func newownedQTableWidget(h uintptr) *QTableWidget { return bindQTableWidget(h) }

func (this *QTableWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTableWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTableWidget(h unsafe.Pointer) *QTableWidget { return bindQTableWidget(uintptr(h)) }
func UnsafeNewQTableWidgetOwned(h unsafe.Pointer) *QTableWidget { return newownedQTableWidget(uintptr(h)) }

func (this *QTableWidget) UnsafeRelease() {
	this.h = 0
	if this.QTableView != nil {
		this.QTableView.UnsafeRelease()
	}
}

func NewQTableWidget(parent *QWidget) *QTableWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTableWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidget_virtIdx, QTableWidget_enable)
	return __o
}

func NewQTableWidget2() *QTableWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidget_new2, uintptr(__goid))
	__o := newownedQTableWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidget_virtIdx, QTableWidget_enable)
	return __o
}

func NewQTableWidget3(rows int, columns int) *QTableWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidget_new3, uintptr(__goid), uintptr(rows), uintptr(columns))
	__o := newownedQTableWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidget_virtIdx, QTableWidget_enable)
	return __o
}

func NewQTableWidget4(rows int, columns int, parent *QWidget) *QTableWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidget_new4, uintptr(__goid), uintptr(rows), uintptr(columns), uintptr(parent.UnsafePointer()))
	__o := newownedQTableWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidget_virtIdx, QTableWidget_enable)
	return __o
}

func QTableWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTableWidget) SetRowCount(rows int) {
	purego.SyscallN(addr_QTableWidget_setRowCount, this.h, uintptr(rows))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) SetColumnCount(columns int) {
	purego.SyscallN(addr_QTableWidget_setColumnCount, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) Row(item *QTableWidgetItem) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_row, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return int(__r)
}

func (this *QTableWidget) Column(item *QTableWidgetItem) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_column, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return int(__r)
}

func (this *QTableWidget) Item(row int, column int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_item, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetItem(row int, column int, item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_setItem, this.h, uintptr(row), uintptr(column), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) TakeItem(row int, column int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_takeItem, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) Items(data *QMimeData) []*QTableWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_items, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	__out := make([]*QTableWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTableWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableWidget) IndexFromItem(item *QTableWidgetItem) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_indexFromItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQModelIndex(__r)
}

func (this *QTableWidget) ItemFromIndex(index *QModelIndex) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_itemFromIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) VerticalHeaderItem(row int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_verticalHeaderItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetVerticalHeaderItem(row int, item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_setVerticalHeaderItem, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) TakeVerticalHeaderItem(row int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_takeVerticalHeaderItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) HorizontalHeaderItem(column int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_horizontalHeaderItem, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetHorizontalHeaderItem(column int, item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_setHorizontalHeaderItem, this.h, uintptr(column), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) TakeHorizontalHeaderItem(column int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_takeHorizontalHeaderItem, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetVerticalHeaderLabels(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QTableWidget_setVerticalHeaderLabels, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QTableWidget) SetHorizontalHeaderLabels(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QTableWidget_setHorizontalHeaderLabels, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QTableWidget) CurrentRow() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_currentRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) CurrentColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_currentColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) CurrentItem() *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_currentItem, this.h)
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetCurrentItem(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_setCurrentItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) SetCurrentItem2(item *QTableWidgetItem, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTableWidget_setCurrentItem2, this.h, uintptr(item.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) SetCurrentCell(row int, column int) {
	purego.SyscallN(addr_QTableWidget_setCurrentCell, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SetCurrentCell2(row int, column int, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTableWidget_setCurrentCell2, this.h, uintptr(row), uintptr(column), uintptr(command))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SortItems(column int) {
	purego.SyscallN(addr_QTableWidget_sortItems, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SetSortingEnabled(enable bool) {
	purego.SyscallN(addr_QTableWidget_setSortingEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) IsSortingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_isSortingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableWidget) EditItem(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_editItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) OpenPersistentEditor(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_openPersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ClosePersistentEditor(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_closePersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) IsPersistentEditorOpen(item *QTableWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_isPersistentEditorOpen, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QTableWidget) CellWidget(row int, column int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_cellWidget, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QTableWidget) SetCellWidget(row int, column int, widget *QWidget) {
	purego.SyscallN(addr_QTableWidget_setCellWidget, this.h, uintptr(row), uintptr(column), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QTableWidget) RemoveCellWidget(row int, column int) {
	purego.SyscallN(addr_QTableWidget_removeCellWidget, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SetRangeSelected(range_ *QTableWidgetSelectionRange, select_ bool) {
	purego.SyscallN(addr_QTableWidget_setRangeSelected, this.h, uintptr(range_.UnsafePointer()), uintptr(b2i(select_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(range_)
}

func (this *QTableWidget) SelectedRanges() []*QTableWidgetSelectionRange {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_selectedRanges, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTableWidgetSelectionRange, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTableWidgetSelectionRange(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableWidget) SelectedItems() []*QTableWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_selectedItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTableWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTableWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableWidget) FindItems(text string, flags MatchFlag) []*QTableWidgetItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_findItems, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QTableWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTableWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableWidget) VisualRow(logicalRow int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_visualRow, this.h, uintptr(logicalRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) VisualColumn(logicalColumn int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_visualColumn, this.h, uintptr(logicalColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) ItemAt(p *QPoint) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_itemAt, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) ItemAt2(x int, y int) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_itemAt2, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) VisualItemRect(item *QTableWidgetItem) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_visualItemRect, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRect(__r)
}

func (this *QTableWidget) ItemPrototype() *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_itemPrototype, this.h)
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidget) SetItemPrototype(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_setItemPrototype, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) SupportedDragActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_supportedDragActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTableWidget) SetSupportedDragActions(actions DropAction) {
	purego.SyscallN(addr_QTableWidget_setSupportedDragActions, this.h, uintptr(actions))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ScrollToItem(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_scrollToItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) InsertRow(row int) {
	purego.SyscallN(addr_QTableWidget_insertRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) InsertColumn(column int) {
	purego.SyscallN(addr_QTableWidget_insertColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RemoveRow(row int) {
	purego.SyscallN(addr_QTableWidget_removeRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RemoveColumn(column int) {
	purego.SyscallN(addr_QTableWidget_removeColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Clear() {
	purego.SyscallN(addr_QTableWidget_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ClearContents() {
	purego.SyscallN(addr_QTableWidget_clearContents, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ItemPressed(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemPressed, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ItemClicked(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemClicked, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ItemDoubleClicked(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemDoubleClicked, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ItemActivated(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemActivated, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ItemEntered(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemEntered, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) ItemChanged(item *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_itemChanged, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTableWidget) CurrentItemChanged(current *QTableWidgetItem, previous *QTableWidgetItem) {
	purego.SyscallN(addr_QTableWidget_currentItemChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTableWidget) ItemSelectionChanged() {
	purego.SyscallN(addr_QTableWidget_itemSelectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellPressed(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellPressed, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellClicked(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellClicked, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellDoubleClicked(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellDoubleClicked, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellActivated(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellActivated, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellEntered(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellEntered, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CellChanged(row int, column int) {
	purego.SyscallN(addr_QTableWidget_cellChanged, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CurrentCellChanged(currentRow int, currentColumn int, previousRow int, previousColumn int) {
	purego.SyscallN(addr_QTableWidget_currentCellChanged, this.h, uintptr(currentRow), uintptr(currentColumn), uintptr(previousRow), uintptr(previousColumn))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RowMoved(row int, oldIndex int, newIndex int) {
	purego.SyscallN(addr_QTableWidget_rowMoved, this.h, uintptr(row), uintptr(oldIndex), uintptr(newIndex))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ColumnMoved(column int, oldIndex int, newIndex int) {
	purego.SyscallN(addr_QTableWidget_columnMoved, this.h, uintptr(column), uintptr(oldIndex), uintptr(newIndex))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RowResized(row int, oldHeight int, newHeight int) {
	purego.SyscallN(addr_QTableWidget_rowResized, this.h, uintptr(row), uintptr(oldHeight), uintptr(newHeight))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ColumnResized(column int, oldWidth int, newWidth int) {
	purego.SyscallN(addr_QTableWidget_columnResized, this.h, uintptr(column), uintptr(oldWidth), uintptr(newWidth))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RowCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTableWidget_rowCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ColumnCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTableWidget_columnCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QTableWidget) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QTableWidget_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QTableWidget_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QTableWidget_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QTableWidget) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QTableWidget_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTableWidget) StartAutoScroll() {
	purego.SyscallN(addr_QTableWidget_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) StopAutoScroll() {
	purego.SyscallN(addr_QTableWidget_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) DoAutoScroll() {
	purego.SyscallN(addr_QTableWidget_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QTableWidget) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTableWidget_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTableWidget_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTableWidget) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTableWidget) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTableWidget_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTableWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTableWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTableWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QTableWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QTableWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTableWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTableWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QTableWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Create() {
	purego.SyscallN(addr_QTableWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Destroy() {
	purego.SyscallN(addr_QTableWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTableWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTableWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTableWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTableWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTableWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTableWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTableWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QTableWidget_devicePixelRatioFScale()
	return __r
}

func QTableWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTableWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTableWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTableWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTableWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTableWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTableWidget) SortItems2(column int, order SortOrder) {
	purego.SyscallN(addr_QTableWidget_sortItems2, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ScrollToItem2(item *QTableWidgetItem, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTableWidget_scrollToItem2, this.h, uintptr(item.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func QTableWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTableWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTableWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTableWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTableWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTableWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTableWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTableWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTableWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTableWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTableWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTableWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTableWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTableWidget) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QTableWidget) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QTableWidget) DropMimeDataDefault(row int, column int, data *QMimeData, action DropAction) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_dropMimeData_default, this.h, uintptr(row), uintptr(column), uintptr(data.UnsafePointer()), uintptr(action))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QTableWidget) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTableWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTableWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QTableWidget_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTableWidget) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QTableWidget_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QTableWidget) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QTableWidget_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QTableWidget) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTableWidget_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTableWidget) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QTableWidget) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTableWidget_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QTableWidget_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTableWidget) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QTableWidget_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTableWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTableWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTableWidget) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTableWidget_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QTableWidget) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QTableWidget) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidget_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableWidget) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QTableWidget_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableWidget) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTableWidget_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTableWidget_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTableWidget) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QTableWidget_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QTableWidget) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QTableWidget_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTableWidget) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QTableWidget_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QTableWidget) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QTableWidget) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTableWidget) ResetDefault() {
	purego.SyscallN(addr_QTableWidget_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) SelectAllDefault() {
	purego.SyscallN(addr_QTableWidget_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTableWidget_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTableWidget) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTableWidget_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTableWidget) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QTableWidget_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QTableWidget_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableWidget) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTableWidget_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTableWidget_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QTableWidget_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableWidget) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QTableWidget_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableWidget) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QTableWidget_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableWidget) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableWidget) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QTableWidget) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QTableWidget_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableWidget) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTableWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTableWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTableWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTableWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTableWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTableWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QTableWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QTableWidget_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableWidget) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTableWidget_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTableWidget) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QTableWidget_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableWidget) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QTableWidget_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QTableWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableWidget) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTableWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTableWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTableWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTableWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTableWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTableWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTableWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTableWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTableWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTableWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTableWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTableWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QTableWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTableWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTableWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTableWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTableWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTableWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTableWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTableWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTableWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTableWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTableWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTableWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTableWidget) OnItemPressed(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemClicked(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemDoubleClicked(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemActivated(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemEntered(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemChanged(fn func(*QTableWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(item))) })
	pg_QTableWidget_connect_itemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCurrentItemChanged(fn func(*QTableWidgetItem, *QTableWidgetItem)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQTableWidgetItem(unsafe.Pointer(current)), UnsafeNewQTableWidgetItem(unsafe.Pointer(previous))) })
	pg_QTableWidget_connect_currentItemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnItemSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTableWidget_connect_itemSelectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellPressed(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellClicked(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellDoubleClicked(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellActivated(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellEntered(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCellChanged(fn func(int, int)) int64 {
	id := registerCallback(func(row int32, column int32) { fn(int(row), int(column)) })
	pg_QTableWidget_connect_cellChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCurrentCellChanged(fn func(int, int, int, int)) int64 {
	id := registerCallback(func(currentRow int32, currentColumn int32, previousRow int32, previousColumn int32) { fn(int(currentRow), int(currentColumn), int(previousRow), int(previousColumn)) })
	pg_QTableWidget_connect_currentCellChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableWidget_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableWidget_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableWidget_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableWidget_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableWidget_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTableWidget_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QTableWidget_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTableWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTableWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTableWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTableWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTableWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTableWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTableWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableWidget) Disconnect(handle int64) {
	pg_QTableWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QTableWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5464467312394730418:
		return 2
	case -9049994684305682212:
		return 3
	case -7784642176186821205:
		return 4
	case 9107058682754591716:
		return 5
	case -1574077111192992763:
		return 6
	case -4689322347236366780:
		return 7
	case -1888629738259754754:
		return 8
	case -1267696851865372997:
		return 9
	case -4832340528568269947:
		return 10
	case -7252823847285716802:
		return 11
	case -6118001150126564745:
		return 12
	case -7450787277202094556:
		return 13
	case 5290442163167458465:
		return 14
	case -5409927581812374498:
		return 15
	case -1500310024801917802:
		return 16
	case 5218707810677921016:
		return 17
	case -158523882194220890:
		return 18
	case -1539266838375037305:
		return 19
	case 7063795720810552074:
		return 20
	case 740869982805721908:
		return 21
	case 1022771470898460576:
		return 22
	case 503579067252048239:
		return 23
	case -4084029023708556876:
		return 24
	case -3666792239577137638:
		return 25
	case -1155555122391250623:
		return 26
	case -785551405282638793:
		return 27
	case 6057989667089378859:
		return 28
	case -5892205462145366901:
		return 29
	case -6978583972112316034:
		return 30
	case -5041452885022114200:
		return 31
	case -3811438228839363824:
		return 32
	case 7345853097340709638:
		return 33
	case 1086335023529244512:
		return 34
	case 8758113809152574300:
		return 35
	case 4472434241583495620:
		return 36
	case -8151049665126058157:
		return 37
	case 9029941737538200621:
		return 38
	case -8726300139121449691:
		return 39
	case 6366953527806415524:
		return 40
	case 1786430177580376614:
		return 41
	case 8648484264628816806:
		return 42
	case -6021383853323633948:
		return 43
	case 6979418566323745651:
		return 44
	case 4206689019978932209:
		return 45
	case -2447490284072398302:
		return 46
	case -8112366604639385491:
		return 47
	case 5898328517505784191:
		return 48
	case -1971465653061861519:
		return 49
	case -332090898078543319:
		return 50
	case -4255494004099806581:
		return 51
	case -4415920049297795213:
		return 52
	case -4029073905697023887:
		return 53
	case 206122715919917589:
		return 54
	case -1781545137352745978:
		return 55
	case -3139904158373725006:
		return 56
	case 975043767033552552:
		return 57
	case -385596010576048993:
		return 58
	case 1023791977679793185:
		return 59
	case -4302917285655014661:
		return 60
	case 5438819661942843222:
		return 61
	case 5217623514630858297:
		return 62
	case 7049264599837532491:
		return 63
	case 729690261457684101:
		return 64
	case 4531996381736730688:
		return 65
	case 1919988340934511316:
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

func QTableWidget_enable(h uintptr, idx int32) { pg_QTableWidget__set_override(h, idx) }

func (this *QTableWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTableWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QTableWidget) OnMimeTypes(fn func() []string) {
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

func (this *QTableWidget) OnDropMimeData(fn func(int, int, *QMimeData, DropAction) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(row int32, column int32, data uintptr, action int32) int32 { return b2i(fn(int(row), int(column), UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action))) }) // slot "dropMimeData"
}

func (this *QTableWidget) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QTableWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTableWidget) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QTableWidget) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QTableWidget) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QTableWidget) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QTableWidget) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QTableWidget) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QTableWidget) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTableWidget) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QTableWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QTableWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTableWidget) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QTableWidget) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QTableWidget) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QTableWidget) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QTableWidget) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QTableWidget) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QTableWidget) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QTableWidget) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTableWidget) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QTableWidget) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QTableWidget) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QTableWidget) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QTableWidget) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QTableWidget) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QTableWidget) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QTableWidget) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QTableWidget) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QTableWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTableWidget) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QTableWidget) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QTableWidget) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QTableWidget) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QTableWidget) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QTableWidget) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QTableWidget) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QTableWidget) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QTableWidget) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QTableWidget) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QTableWidget) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QTableWidget) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QTableWidget) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QTableWidget) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QTableWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTableWidget) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QTableWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QTableWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QTableWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QTableWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTableWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTableWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTableWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTableWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTableWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTableWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QTableWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QTableWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QTableWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTableWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTableWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTableWidget) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTableWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QTableWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QTableWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QTableWidget) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTableWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTableWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTableWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTableWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTableWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTableWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTableWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTableWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTableWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTableWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTableWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTableWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTableWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QTableWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTableWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTableWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTableWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTableWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTableWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTableWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTableWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTableWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTableWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTableWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTableWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
