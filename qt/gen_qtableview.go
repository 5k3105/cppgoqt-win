package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTableView_chain uintptr
	addr_QTableView_new uintptr
	addr_QTableView_new2 uintptr
	addr_QTableView_tr uintptr
	addr_QTableView_horizontalHeader uintptr
	addr_QTableView_verticalHeader uintptr
	addr_QTableView_setHorizontalHeader uintptr
	addr_QTableView_setVerticalHeader uintptr
	addr_QTableView_rowViewportPosition uintptr
	addr_QTableView_rowAt uintptr
	addr_QTableView_setRowHeight uintptr
	addr_QTableView_rowHeight uintptr
	addr_QTableView_columnViewportPosition uintptr
	addr_QTableView_columnAt uintptr
	addr_QTableView_setColumnWidth uintptr
	addr_QTableView_columnWidth uintptr
	addr_QTableView_isRowHidden uintptr
	addr_QTableView_setRowHidden uintptr
	addr_QTableView_isColumnHidden uintptr
	addr_QTableView_setColumnHidden uintptr
	addr_QTableView_setSortingEnabled uintptr
	addr_QTableView_isSortingEnabled uintptr
	addr_QTableView_showGrid uintptr
	addr_QTableView_gridStyle uintptr
	addr_QTableView_setGridStyle uintptr
	addr_QTableView_setWordWrap uintptr
	addr_QTableView_wordWrap uintptr
	addr_QTableView_setCornerButtonEnabled uintptr
	addr_QTableView_isCornerButtonEnabled uintptr
	addr_QTableView_setSpan uintptr
	addr_QTableView_rowSpan uintptr
	addr_QTableView_columnSpan uintptr
	addr_QTableView_clearSpans uintptr
	addr_QTableView_selectRow uintptr
	addr_QTableView_selectColumn uintptr
	addr_QTableView_hideRow uintptr
	addr_QTableView_hideColumn uintptr
	addr_QTableView_showRow uintptr
	addr_QTableView_showColumn uintptr
	addr_QTableView_resizeRowToContents uintptr
	addr_QTableView_resizeRowsToContents uintptr
	addr_QTableView_resizeColumnToContents uintptr
	addr_QTableView_resizeColumnsToContents uintptr
	addr_QTableView_sortByColumn uintptr
	addr_QTableView_setShowGrid uintptr
	addr_QTableView_rowMoved uintptr
	addr_QTableView_columnMoved uintptr
	addr_QTableView_rowResized uintptr
	addr_QTableView_columnResized uintptr
	addr_QTableView_rowCountChanged uintptr
	addr_QTableView_columnCountChanged uintptr
	addr_QTableView_state uintptr
	addr_QTableView_scheduleDelayedItemsLayout uintptr
	addr_QTableView_executeDelayedItemsLayout uintptr
	addr_QTableView_setDirtyRegion uintptr
	addr_QTableView_scrollDirtyRegion uintptr
	addr_QTableView_dirtyRegionOffset uintptr
	addr_QTableView_startAutoScroll uintptr
	addr_QTableView_stopAutoScroll uintptr
	addr_QTableView_doAutoScroll uintptr
	addr_QTableView_dropIndicatorPosition uintptr
	addr_QTableView_setViewportMargins uintptr
	addr_QTableView_setViewportMargins2 uintptr
	addr_QTableView_viewportMargins uintptr
	addr_QTableView_drawFrame uintptr
	addr_QTableView_setTabOrder uintptr
	addr_QTableView_mouseGrabber uintptr
	addr_QTableView_keyboardGrabber uintptr
	addr_QTableView_find uintptr
	addr_QTableView_createWindowContainer uintptr
	addr_QTableView_updateMicroFocus uintptr
	addr_QTableView_create uintptr
	addr_QTableView_destroy uintptr
	addr_QTableView_focusNextChild uintptr
	addr_QTableView_focusPreviousChild uintptr
	addr_QTableView_connect uintptr
	addr_QTableView_connect2 uintptr
	addr_QTableView_sender uintptr
	addr_QTableView_senderSignalIndex uintptr
	addr_QTableView_receivers uintptr
	addr_QTableView_isSignalConnected uintptr
	pg_QTableView_devicePixelRatioFScale func() float64
	pg_QTableView_encodeMetricF func(int32, float64) int32
	pg_QTableView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTableView_tr2 uintptr
	addr_QTableView_tr3 uintptr
	addr_QTableView_createWindowContainer2 uintptr
	addr_QTableView_createWindowContainer3 uintptr
	addr_QTableView_updateMicroFocus2 uintptr
	addr_QTableView_create2 uintptr
	addr_QTableView_create3 uintptr
	addr_QTableView_create4 uintptr
	addr_QTableView_destroy2 uintptr
	addr_QTableView_destroy3 uintptr
	addr_QTableView_connect3 uintptr
	addr_QTableView_connect4 uintptr
	addr_QTableView_metaObject_default uintptr
	addr_QTableView_setModel_default uintptr
	addr_QTableView_setRootIndex_default uintptr
	addr_QTableView_setSelectionModel_default uintptr
	addr_QTableView_doItemsLayout_default uintptr
	addr_QTableView_visualRect_default uintptr
	addr_QTableView_scrollTo_default uintptr
	addr_QTableView_indexAt_default uintptr
	addr_QTableView_scrollContentsBy_default uintptr
	addr_QTableView_initViewItemOption_default uintptr
	addr_QTableView_paintEvent_default uintptr
	addr_QTableView_timerEvent_default uintptr
	addr_QTableView_dropEvent_default uintptr
	addr_QTableView_horizontalOffset_default uintptr
	addr_QTableView_verticalOffset_default uintptr
	addr_QTableView_moveCursor_default uintptr
	addr_QTableView_setSelection_default uintptr
	addr_QTableView_visualRegionForSelection_default uintptr
	addr_QTableView_selectedIndexes_default uintptr
	addr_QTableView_updateGeometries_default uintptr
	addr_QTableView_viewportSizeHint_default uintptr
	addr_QTableView_sizeHintForRow_default uintptr
	addr_QTableView_sizeHintForColumn_default uintptr
	addr_QTableView_verticalScrollbarAction_default uintptr
	addr_QTableView_horizontalScrollbarAction_default uintptr
	addr_QTableView_isIndexHidden_default uintptr
	addr_QTableView_selectionChanged_default uintptr
	addr_QTableView_currentChanged_default uintptr
	addr_QTableView_keyboardSearch_default uintptr
	addr_QTableView_itemDelegateForIndex_default uintptr
	addr_QTableView_inputMethodQuery_default uintptr
	addr_QTableView_reset_default uintptr
	addr_QTableView_selectAll_default uintptr
	addr_QTableView_rowsInserted_default uintptr
	addr_QTableView_rowsAboutToBeRemoved_default uintptr
	addr_QTableView_updateEditorData_default uintptr
	addr_QTableView_updateEditorGeometries_default uintptr
	addr_QTableView_verticalScrollbarValueChanged_default uintptr
	addr_QTableView_horizontalScrollbarValueChanged_default uintptr
	addr_QTableView_closeEditor_default uintptr
	addr_QTableView_commitData_default uintptr
	addr_QTableView_editorDestroyed_default uintptr
	addr_QTableView_edit_default uintptr
	addr_QTableView_selectionCommand_default uintptr
	addr_QTableView_startDrag_default uintptr
	addr_QTableView_focusNextPrevChild_default uintptr
	addr_QTableView_event_default uintptr
	addr_QTableView_viewportEvent_default uintptr
	addr_QTableView_mousePressEvent_default uintptr
	addr_QTableView_mouseMoveEvent_default uintptr
	addr_QTableView_mouseReleaseEvent_default uintptr
	addr_QTableView_mouseDoubleClickEvent_default uintptr
	addr_QTableView_dragEnterEvent_default uintptr
	addr_QTableView_dragMoveEvent_default uintptr
	addr_QTableView_dragLeaveEvent_default uintptr
	addr_QTableView_focusInEvent_default uintptr
	addr_QTableView_focusOutEvent_default uintptr
	addr_QTableView_keyPressEvent_default uintptr
	addr_QTableView_resizeEvent_default uintptr
	addr_QTableView_inputMethodEvent_default uintptr
	addr_QTableView_eventFilter_default uintptr
	addr_QTableView_minimumSizeHint_default uintptr
	addr_QTableView_sizeHint_default uintptr
	addr_QTableView_setupViewport_default uintptr
	addr_QTableView_wheelEvent_default uintptr
	addr_QTableView_contextMenuEvent_default uintptr
	addr_QTableView_changeEvent_default uintptr
	addr_QTableView_initStyleOption_default uintptr
	addr_QTableView_devType_default uintptr
	addr_QTableView_setVisible_default uintptr
	addr_QTableView_heightForWidth_default uintptr
	addr_QTableView_hasHeightForWidth_default uintptr
	addr_QTableView_paintEngine_default uintptr
	addr_QTableView_keyReleaseEvent_default uintptr
	addr_QTableView_enterEvent_default uintptr
	addr_QTableView_leaveEvent_default uintptr
	addr_QTableView_moveEvent_default uintptr
	addr_QTableView_closeEvent_default uintptr
	addr_QTableView_tabletEvent_default uintptr
	addr_QTableView_actionEvent_default uintptr
	addr_QTableView_showEvent_default uintptr
	addr_QTableView_hideEvent_default uintptr
	addr_QTableView_metric_default uintptr
	addr_QTableView_initPainter_default uintptr
	addr_QTableView_redirected_default uintptr
	addr_QTableView_sharedPainter_default uintptr
	addr_QTableView_childEvent_default uintptr
	addr_QTableView_customEvent_default uintptr
	addr_QTableView_connectNotify_default uintptr
	addr_QTableView_disconnectNotify_default uintptr
	pg_QTableView_connect_pressed func(uintptr, int64)
	pg_QTableView_connect_clicked func(uintptr, int64)
	pg_QTableView_connect_doubleClicked func(uintptr, int64)
	pg_QTableView_connect_activated func(uintptr, int64)
	pg_QTableView_connect_entered func(uintptr, int64)
	pg_QTableView_connect_viewportEntered func(uintptr, int64)
	pg_QTableView_connect_iconSizeChanged func(uintptr, int64)
	pg_QTableView_connect_windowTitleChanged func(uintptr, int64)
	pg_QTableView_connect_windowIconChanged func(uintptr, int64)
	pg_QTableView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTableView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTableView_connect_destroyed func(uintptr, int64)
	pg_QTableView_connect_destroyed2 func(uintptr, int64)
	pg_QTableView_connect_objectNameChanged func(uintptr, int64)
	pg_QTableView__disconnect func(int64)
	pg_QTableView__set_override func(uintptr, int32)
	pg_QTableView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTableView_chain, "QTableView_chain")
	symaddr(&addr_QTableView_new, "QTableView_new")
	symaddr(&addr_QTableView_new2, "QTableView_new2")
	symaddr(&addr_QTableView_tr, "QTableView_tr")
	symaddr(&addr_QTableView_horizontalHeader, "QTableView_horizontalHeader")
	symaddr(&addr_QTableView_verticalHeader, "QTableView_verticalHeader")
	symaddr(&addr_QTableView_setHorizontalHeader, "QTableView_setHorizontalHeader")
	symaddr(&addr_QTableView_setVerticalHeader, "QTableView_setVerticalHeader")
	symaddr(&addr_QTableView_rowViewportPosition, "QTableView_rowViewportPosition")
	symaddr(&addr_QTableView_rowAt, "QTableView_rowAt")
	symaddr(&addr_QTableView_setRowHeight, "QTableView_setRowHeight")
	symaddr(&addr_QTableView_rowHeight, "QTableView_rowHeight")
	symaddr(&addr_QTableView_columnViewportPosition, "QTableView_columnViewportPosition")
	symaddr(&addr_QTableView_columnAt, "QTableView_columnAt")
	symaddr(&addr_QTableView_setColumnWidth, "QTableView_setColumnWidth")
	symaddr(&addr_QTableView_columnWidth, "QTableView_columnWidth")
	symaddr(&addr_QTableView_isRowHidden, "QTableView_isRowHidden")
	symaddr(&addr_QTableView_setRowHidden, "QTableView_setRowHidden")
	symaddr(&addr_QTableView_isColumnHidden, "QTableView_isColumnHidden")
	symaddr(&addr_QTableView_setColumnHidden, "QTableView_setColumnHidden")
	symaddr(&addr_QTableView_setSortingEnabled, "QTableView_setSortingEnabled")
	symaddr(&addr_QTableView_isSortingEnabled, "QTableView_isSortingEnabled")
	symaddr(&addr_QTableView_showGrid, "QTableView_showGrid")
	symaddr(&addr_QTableView_gridStyle, "QTableView_gridStyle")
	symaddr(&addr_QTableView_setGridStyle, "QTableView_setGridStyle")
	symaddr(&addr_QTableView_setWordWrap, "QTableView_setWordWrap")
	symaddr(&addr_QTableView_wordWrap, "QTableView_wordWrap")
	symaddr(&addr_QTableView_setCornerButtonEnabled, "QTableView_setCornerButtonEnabled")
	symaddr(&addr_QTableView_isCornerButtonEnabled, "QTableView_isCornerButtonEnabled")
	symaddr(&addr_QTableView_setSpan, "QTableView_setSpan")
	symaddr(&addr_QTableView_rowSpan, "QTableView_rowSpan")
	symaddr(&addr_QTableView_columnSpan, "QTableView_columnSpan")
	symaddr(&addr_QTableView_clearSpans, "QTableView_clearSpans")
	symaddr(&addr_QTableView_selectRow, "QTableView_selectRow")
	symaddr(&addr_QTableView_selectColumn, "QTableView_selectColumn")
	symaddr(&addr_QTableView_hideRow, "QTableView_hideRow")
	symaddr(&addr_QTableView_hideColumn, "QTableView_hideColumn")
	symaddr(&addr_QTableView_showRow, "QTableView_showRow")
	symaddr(&addr_QTableView_showColumn, "QTableView_showColumn")
	symaddr(&addr_QTableView_resizeRowToContents, "QTableView_resizeRowToContents")
	symaddr(&addr_QTableView_resizeRowsToContents, "QTableView_resizeRowsToContents")
	symaddr(&addr_QTableView_resizeColumnToContents, "QTableView_resizeColumnToContents")
	symaddr(&addr_QTableView_resizeColumnsToContents, "QTableView_resizeColumnsToContents")
	symaddr(&addr_QTableView_sortByColumn, "QTableView_sortByColumn")
	symaddr(&addr_QTableView_setShowGrid, "QTableView_setShowGrid")
	symaddr(&addr_QTableView_rowMoved, "QTableView_rowMoved")
	symaddr(&addr_QTableView_columnMoved, "QTableView_columnMoved")
	symaddr(&addr_QTableView_rowResized, "QTableView_rowResized")
	symaddr(&addr_QTableView_columnResized, "QTableView_columnResized")
	symaddr(&addr_QTableView_rowCountChanged, "QTableView_rowCountChanged")
	symaddr(&addr_QTableView_columnCountChanged, "QTableView_columnCountChanged")
	symaddr(&addr_QTableView_state, "QTableView_state")
	symaddr(&addr_QTableView_scheduleDelayedItemsLayout, "QTableView_scheduleDelayedItemsLayout")
	symaddr(&addr_QTableView_executeDelayedItemsLayout, "QTableView_executeDelayedItemsLayout")
	symaddr(&addr_QTableView_setDirtyRegion, "QTableView_setDirtyRegion")
	symaddr(&addr_QTableView_scrollDirtyRegion, "QTableView_scrollDirtyRegion")
	symaddr(&addr_QTableView_dirtyRegionOffset, "QTableView_dirtyRegionOffset")
	symaddr(&addr_QTableView_startAutoScroll, "QTableView_startAutoScroll")
	symaddr(&addr_QTableView_stopAutoScroll, "QTableView_stopAutoScroll")
	symaddr(&addr_QTableView_doAutoScroll, "QTableView_doAutoScroll")
	symaddr(&addr_QTableView_dropIndicatorPosition, "QTableView_dropIndicatorPosition")
	symaddr(&addr_QTableView_setViewportMargins, "QTableView_setViewportMargins")
	symaddr(&addr_QTableView_setViewportMargins2, "QTableView_setViewportMargins2")
	symaddr(&addr_QTableView_viewportMargins, "QTableView_viewportMargins")
	symaddr(&addr_QTableView_drawFrame, "QTableView_drawFrame")
	symaddr(&addr_QTableView_setTabOrder, "QTableView_setTabOrder")
	symaddr(&addr_QTableView_mouseGrabber, "QTableView_mouseGrabber")
	symaddr(&addr_QTableView_keyboardGrabber, "QTableView_keyboardGrabber")
	symaddr(&addr_QTableView_find, "QTableView_find")
	symaddr(&addr_QTableView_createWindowContainer, "QTableView_createWindowContainer")
	symaddr(&addr_QTableView_updateMicroFocus, "QTableView_updateMicroFocus")
	symaddr(&addr_QTableView_create, "QTableView_create")
	symaddr(&addr_QTableView_destroy, "QTableView_destroy")
	symaddr(&addr_QTableView_focusNextChild, "QTableView_focusNextChild")
	symaddr(&addr_QTableView_focusPreviousChild, "QTableView_focusPreviousChild")
	symaddr(&addr_QTableView_connect, "QTableView_connect")
	symaddr(&addr_QTableView_connect2, "QTableView_connect2")
	symaddr(&addr_QTableView_sender, "QTableView_sender")
	symaddr(&addr_QTableView_senderSignalIndex, "QTableView_senderSignalIndex")
	symaddr(&addr_QTableView_receivers, "QTableView_receivers")
	symaddr(&addr_QTableView_isSignalConnected, "QTableView_isSignalConnected")
	reg(&pg_QTableView_devicePixelRatioFScale, "QTableView_devicePixelRatioFScale")
	reg(&pg_QTableView_encodeMetricF, "QTableView_encodeMetricF")
	reg(&pg_QTableView_getDecodedMetricF, "QTableView_getDecodedMetricF")
	symaddr(&addr_QTableView_tr2, "QTableView_tr2")
	symaddr(&addr_QTableView_tr3, "QTableView_tr3")
	symaddr(&addr_QTableView_createWindowContainer2, "QTableView_createWindowContainer2")
	symaddr(&addr_QTableView_createWindowContainer3, "QTableView_createWindowContainer3")
	symaddr(&addr_QTableView_updateMicroFocus2, "QTableView_updateMicroFocus2")
	symaddr(&addr_QTableView_create2, "QTableView_create2")
	symaddr(&addr_QTableView_create3, "QTableView_create3")
	symaddr(&addr_QTableView_create4, "QTableView_create4")
	symaddr(&addr_QTableView_destroy2, "QTableView_destroy2")
	symaddr(&addr_QTableView_destroy3, "QTableView_destroy3")
	symaddr(&addr_QTableView_connect3, "QTableView_connect3")
	symaddr(&addr_QTableView_connect4, "QTableView_connect4")
	symaddr(&addr_QTableView_metaObject_default, "QTableView_metaObject_default")
	symaddr(&addr_QTableView_setModel_default, "QTableView_setModel_default")
	symaddr(&addr_QTableView_setRootIndex_default, "QTableView_setRootIndex_default")
	symaddr(&addr_QTableView_setSelectionModel_default, "QTableView_setSelectionModel_default")
	symaddr(&addr_QTableView_doItemsLayout_default, "QTableView_doItemsLayout_default")
	symaddr(&addr_QTableView_visualRect_default, "QTableView_visualRect_default")
	symaddr(&addr_QTableView_scrollTo_default, "QTableView_scrollTo_default")
	symaddr(&addr_QTableView_indexAt_default, "QTableView_indexAt_default")
	symaddr(&addr_QTableView_scrollContentsBy_default, "QTableView_scrollContentsBy_default")
	symaddr(&addr_QTableView_initViewItemOption_default, "QTableView_initViewItemOption_default")
	symaddr(&addr_QTableView_paintEvent_default, "QTableView_paintEvent_default")
	symaddr(&addr_QTableView_timerEvent_default, "QTableView_timerEvent_default")
	symaddr(&addr_QTableView_dropEvent_default, "QTableView_dropEvent_default")
	symaddr(&addr_QTableView_horizontalOffset_default, "QTableView_horizontalOffset_default")
	symaddr(&addr_QTableView_verticalOffset_default, "QTableView_verticalOffset_default")
	symaddr(&addr_QTableView_moveCursor_default, "QTableView_moveCursor_default")
	symaddr(&addr_QTableView_setSelection_default, "QTableView_setSelection_default")
	symaddr(&addr_QTableView_visualRegionForSelection_default, "QTableView_visualRegionForSelection_default")
	symaddr(&addr_QTableView_selectedIndexes_default, "QTableView_selectedIndexes_default")
	symaddr(&addr_QTableView_updateGeometries_default, "QTableView_updateGeometries_default")
	symaddr(&addr_QTableView_viewportSizeHint_default, "QTableView_viewportSizeHint_default")
	symaddr(&addr_QTableView_sizeHintForRow_default, "QTableView_sizeHintForRow_default")
	symaddr(&addr_QTableView_sizeHintForColumn_default, "QTableView_sizeHintForColumn_default")
	symaddr(&addr_QTableView_verticalScrollbarAction_default, "QTableView_verticalScrollbarAction_default")
	symaddr(&addr_QTableView_horizontalScrollbarAction_default, "QTableView_horizontalScrollbarAction_default")
	symaddr(&addr_QTableView_isIndexHidden_default, "QTableView_isIndexHidden_default")
	symaddr(&addr_QTableView_selectionChanged_default, "QTableView_selectionChanged_default")
	symaddr(&addr_QTableView_currentChanged_default, "QTableView_currentChanged_default")
	symaddr(&addr_QTableView_keyboardSearch_default, "QTableView_keyboardSearch_default")
	symaddr(&addr_QTableView_itemDelegateForIndex_default, "QTableView_itemDelegateForIndex_default")
	symaddr(&addr_QTableView_inputMethodQuery_default, "QTableView_inputMethodQuery_default")
	symaddr(&addr_QTableView_reset_default, "QTableView_reset_default")
	symaddr(&addr_QTableView_selectAll_default, "QTableView_selectAll_default")
	symaddr(&addr_QTableView_rowsInserted_default, "QTableView_rowsInserted_default")
	symaddr(&addr_QTableView_rowsAboutToBeRemoved_default, "QTableView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QTableView_updateEditorData_default, "QTableView_updateEditorData_default")
	symaddr(&addr_QTableView_updateEditorGeometries_default, "QTableView_updateEditorGeometries_default")
	symaddr(&addr_QTableView_verticalScrollbarValueChanged_default, "QTableView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QTableView_horizontalScrollbarValueChanged_default, "QTableView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QTableView_closeEditor_default, "QTableView_closeEditor_default")
	symaddr(&addr_QTableView_commitData_default, "QTableView_commitData_default")
	symaddr(&addr_QTableView_editorDestroyed_default, "QTableView_editorDestroyed_default")
	symaddr(&addr_QTableView_edit_default, "QTableView_edit_default")
	symaddr(&addr_QTableView_selectionCommand_default, "QTableView_selectionCommand_default")
	symaddr(&addr_QTableView_startDrag_default, "QTableView_startDrag_default")
	symaddr(&addr_QTableView_focusNextPrevChild_default, "QTableView_focusNextPrevChild_default")
	symaddr(&addr_QTableView_event_default, "QTableView_event_default")
	symaddr(&addr_QTableView_viewportEvent_default, "QTableView_viewportEvent_default")
	symaddr(&addr_QTableView_mousePressEvent_default, "QTableView_mousePressEvent_default")
	symaddr(&addr_QTableView_mouseMoveEvent_default, "QTableView_mouseMoveEvent_default")
	symaddr(&addr_QTableView_mouseReleaseEvent_default, "QTableView_mouseReleaseEvent_default")
	symaddr(&addr_QTableView_mouseDoubleClickEvent_default, "QTableView_mouseDoubleClickEvent_default")
	symaddr(&addr_QTableView_dragEnterEvent_default, "QTableView_dragEnterEvent_default")
	symaddr(&addr_QTableView_dragMoveEvent_default, "QTableView_dragMoveEvent_default")
	symaddr(&addr_QTableView_dragLeaveEvent_default, "QTableView_dragLeaveEvent_default")
	symaddr(&addr_QTableView_focusInEvent_default, "QTableView_focusInEvent_default")
	symaddr(&addr_QTableView_focusOutEvent_default, "QTableView_focusOutEvent_default")
	symaddr(&addr_QTableView_keyPressEvent_default, "QTableView_keyPressEvent_default")
	symaddr(&addr_QTableView_resizeEvent_default, "QTableView_resizeEvent_default")
	symaddr(&addr_QTableView_inputMethodEvent_default, "QTableView_inputMethodEvent_default")
	symaddr(&addr_QTableView_eventFilter_default, "QTableView_eventFilter_default")
	symaddr(&addr_QTableView_minimumSizeHint_default, "QTableView_minimumSizeHint_default")
	symaddr(&addr_QTableView_sizeHint_default, "QTableView_sizeHint_default")
	symaddr(&addr_QTableView_setupViewport_default, "QTableView_setupViewport_default")
	symaddr(&addr_QTableView_wheelEvent_default, "QTableView_wheelEvent_default")
	symaddr(&addr_QTableView_contextMenuEvent_default, "QTableView_contextMenuEvent_default")
	symaddr(&addr_QTableView_changeEvent_default, "QTableView_changeEvent_default")
	symaddr(&addr_QTableView_initStyleOption_default, "QTableView_initStyleOption_default")
	symaddr(&addr_QTableView_devType_default, "QTableView_devType_default")
	symaddr(&addr_QTableView_setVisible_default, "QTableView_setVisible_default")
	symaddr(&addr_QTableView_heightForWidth_default, "QTableView_heightForWidth_default")
	symaddr(&addr_QTableView_hasHeightForWidth_default, "QTableView_hasHeightForWidth_default")
	symaddr(&addr_QTableView_paintEngine_default, "QTableView_paintEngine_default")
	symaddr(&addr_QTableView_keyReleaseEvent_default, "QTableView_keyReleaseEvent_default")
	symaddr(&addr_QTableView_enterEvent_default, "QTableView_enterEvent_default")
	symaddr(&addr_QTableView_leaveEvent_default, "QTableView_leaveEvent_default")
	symaddr(&addr_QTableView_moveEvent_default, "QTableView_moveEvent_default")
	symaddr(&addr_QTableView_closeEvent_default, "QTableView_closeEvent_default")
	symaddr(&addr_QTableView_tabletEvent_default, "QTableView_tabletEvent_default")
	symaddr(&addr_QTableView_actionEvent_default, "QTableView_actionEvent_default")
	symaddr(&addr_QTableView_showEvent_default, "QTableView_showEvent_default")
	symaddr(&addr_QTableView_hideEvent_default, "QTableView_hideEvent_default")
	symaddr(&addr_QTableView_metric_default, "QTableView_metric_default")
	symaddr(&addr_QTableView_initPainter_default, "QTableView_initPainter_default")
	symaddr(&addr_QTableView_redirected_default, "QTableView_redirected_default")
	symaddr(&addr_QTableView_sharedPainter_default, "QTableView_sharedPainter_default")
	symaddr(&addr_QTableView_childEvent_default, "QTableView_childEvent_default")
	symaddr(&addr_QTableView_customEvent_default, "QTableView_customEvent_default")
	symaddr(&addr_QTableView_connectNotify_default, "QTableView_connectNotify_default")
	symaddr(&addr_QTableView_disconnectNotify_default, "QTableView_disconnectNotify_default")
	reg(&pg_QTableView_connect_pressed, "QTableView_connect_pressed")
	reg(&pg_QTableView_connect_clicked, "QTableView_connect_clicked")
	reg(&pg_QTableView_connect_doubleClicked, "QTableView_connect_doubleClicked")
	reg(&pg_QTableView_connect_activated, "QTableView_connect_activated")
	reg(&pg_QTableView_connect_entered, "QTableView_connect_entered")
	reg(&pg_QTableView_connect_viewportEntered, "QTableView_connect_viewportEntered")
	reg(&pg_QTableView_connect_iconSizeChanged, "QTableView_connect_iconSizeChanged")
	reg(&pg_QTableView_connect_windowTitleChanged, "QTableView_connect_windowTitleChanged")
	reg(&pg_QTableView_connect_windowIconChanged, "QTableView_connect_windowIconChanged")
	reg(&pg_QTableView_connect_windowIconTextChanged, "QTableView_connect_windowIconTextChanged")
	reg(&pg_QTableView_connect_customContextMenuRequested, "QTableView_connect_customContextMenuRequested")
	reg(&pg_QTableView_connect_destroyed, "QTableView_connect_destroyed")
	reg(&pg_QTableView_connect_destroyed2, "QTableView_connect_destroyed2")
	reg(&pg_QTableView_connect_objectNameChanged, "QTableView_connect_objectNameChanged")
	reg(&pg_QTableView__disconnect, "QTableView__disconnect")
	reg(&pg_QTableView__set_override, "QTableView__set_override")
	reg(&pg_QTableView_delete, "QTableView_delete")
}

type QTableView struct {
	h uintptr
	*QAbstractItemView
}

func bindQTableView(h uintptr) *QTableView {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QTableView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTableView{h: h, QAbstractItemView: &QAbstractItemView{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQTableView(h uintptr) *QTableView { return bindQTableView(h) }

func (this *QTableView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTableView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTableView(h unsafe.Pointer) *QTableView { return bindQTableView(uintptr(h)) }
func UnsafeNewQTableViewOwned(h unsafe.Pointer) *QTableView { return newownedQTableView(uintptr(h)) }

func (this *QTableView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemView != nil {
		this.QAbstractItemView.UnsafeRelease()
	}
}

func NewQTableView(parent *QWidget) *QTableView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTableView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableView_virtIdx, QTableView_enable)
	return __o
}

func NewQTableView2() *QTableView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableView_new2, uintptr(__goid))
	__o := newownedQTableView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableView_virtIdx, QTableView_enable)
	return __o
}

func QTableView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTableView) HorizontalHeader() *QHeaderView {
	__r, _, _ := purego.SyscallN(addr_QTableView_horizontalHeader, this.h)
	runtime.KeepAlive(this)
	return bindQHeaderView(__r)
}

func (this *QTableView) VerticalHeader() *QHeaderView {
	__r, _, _ := purego.SyscallN(addr_QTableView_verticalHeader, this.h)
	runtime.KeepAlive(this)
	return bindQHeaderView(__r)
}

func (this *QTableView) SetHorizontalHeader(header *QHeaderView) {
	purego.SyscallN(addr_QTableView_setHorizontalHeader, this.h, uintptr(header.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(header)
}

func (this *QTableView) SetVerticalHeader(header *QHeaderView) {
	purego.SyscallN(addr_QTableView_setVerticalHeader, this.h, uintptr(header.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(header)
}

func (this *QTableView) RowViewportPosition(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_rowViewportPosition, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) RowAt(y int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_rowAt, this.h, uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) SetRowHeight(row int, height int) {
	purego.SyscallN(addr_QTableView_setRowHeight, this.h, uintptr(row), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QTableView) RowHeight(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_rowHeight, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) ColumnViewportPosition(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_columnViewportPosition, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) ColumnAt(x int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_columnAt, this.h, uintptr(x))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) SetColumnWidth(column int, width int) {
	purego.SyscallN(addr_QTableView_setColumnWidth, this.h, uintptr(column), uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QTableView) ColumnWidth(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_columnWidth, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) IsRowHidden(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isRowHidden, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) SetRowHidden(row int, hide bool) {
	purego.SyscallN(addr_QTableView_setRowHidden, this.h, uintptr(row), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTableView) IsColumnHidden(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isColumnHidden, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) SetColumnHidden(column int, hide bool) {
	purego.SyscallN(addr_QTableView_setColumnHidden, this.h, uintptr(column), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTableView) SetSortingEnabled(enable bool) {
	purego.SyscallN(addr_QTableView_setSortingEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTableView) IsSortingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isSortingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) ShowGrid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_showGrid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) GridStyle() PenStyle {
	__r0, _, _ := purego.SyscallN(addr_QTableView_gridStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenStyle(__r)
}

func (this *QTableView) SetGridStyle(style PenStyle) {
	purego.SyscallN(addr_QTableView_setGridStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTableView) SetWordWrap(on bool) {
	purego.SyscallN(addr_QTableView_setWordWrap, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QTableView) WordWrap() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_wordWrap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) SetCornerButtonEnabled(enable bool) {
	purego.SyscallN(addr_QTableView_setCornerButtonEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTableView) IsCornerButtonEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isCornerButtonEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) SetSpan(row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QTableView_setSpan, this.h, uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
}

func (this *QTableView) RowSpan(row int, column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_rowSpan, this.h, uintptr(row), uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) ColumnSpan(row int, column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_columnSpan, this.h, uintptr(row), uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) ClearSpans() {
	purego.SyscallN(addr_QTableView_clearSpans, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) SelectRow(row int) {
	purego.SyscallN(addr_QTableView_selectRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableView) SelectColumn(column int) {
	purego.SyscallN(addr_QTableView_selectColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableView) HideRow(row int) {
	purego.SyscallN(addr_QTableView_hideRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableView) HideColumn(column int) {
	purego.SyscallN(addr_QTableView_hideColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableView) ShowRow(row int) {
	purego.SyscallN(addr_QTableView_showRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableView) ShowColumn(column int) {
	purego.SyscallN(addr_QTableView_showColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableView) ResizeRowToContents(row int) {
	purego.SyscallN(addr_QTableView_resizeRowToContents, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QTableView) ResizeRowsToContents() {
	purego.SyscallN(addr_QTableView_resizeRowsToContents, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) ResizeColumnToContents(column int) {
	purego.SyscallN(addr_QTableView_resizeColumnToContents, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTableView) ResizeColumnsToContents() {
	purego.SyscallN(addr_QTableView_resizeColumnsToContents, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) SortByColumn(column int, order SortOrder) {
	purego.SyscallN(addr_QTableView_sortByColumn, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTableView) SetShowGrid(show bool) {
	purego.SyscallN(addr_QTableView_setShowGrid, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QTableView) RowMoved(row int, oldIndex int, newIndex int) {
	purego.SyscallN(addr_QTableView_rowMoved, this.h, uintptr(row), uintptr(oldIndex), uintptr(newIndex))
	runtime.KeepAlive(this)
}

func (this *QTableView) ColumnMoved(column int, oldIndex int, newIndex int) {
	purego.SyscallN(addr_QTableView_columnMoved, this.h, uintptr(column), uintptr(oldIndex), uintptr(newIndex))
	runtime.KeepAlive(this)
}

func (this *QTableView) RowResized(row int, oldHeight int, newHeight int) {
	purego.SyscallN(addr_QTableView_rowResized, this.h, uintptr(row), uintptr(oldHeight), uintptr(newHeight))
	runtime.KeepAlive(this)
}

func (this *QTableView) ColumnResized(column int, oldWidth int, newWidth int) {
	purego.SyscallN(addr_QTableView_columnResized, this.h, uintptr(column), uintptr(oldWidth), uintptr(newWidth))
	runtime.KeepAlive(this)
}

func (this *QTableView) RowCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTableView_rowCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTableView) ColumnCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTableView_columnCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTableView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QTableView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QTableView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QTableView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QTableView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QTableView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QTableView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QTableView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTableView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTableView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTableView) StartAutoScroll() {
	purego.SyscallN(addr_QTableView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) StopAutoScroll() {
	purego.SyscallN(addr_QTableView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) DoAutoScroll() {
	purego.SyscallN(addr_QTableView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QTableView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QTableView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTableView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTableView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTableView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTableView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTableView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTableView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTableView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTableView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTableView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTableView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_mouseGrabber)
	return bindQWidget(__r)
}

func QTableView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_keyboardGrabber)
	return bindQWidget(__r)
}

func QTableView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTableView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTableView) UpdateMicroFocus() {
	purego.SyscallN(addr_QTableView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) Create() {
	purego.SyscallN(addr_QTableView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) Destroy() {
	purego.SyscallN(addr_QTableView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTableView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTableView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTableView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTableView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTableView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTableView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTableView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTableView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTableView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTableView_DevicePixelRatioFScale() float64 {
	__r := pg_QTableView_devicePixelRatioFScale()
	return __r
}

func QTableView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTableView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTableView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTableView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTableView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTableView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTableView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTableView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTableView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTableView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTableView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTableView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTableView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTableView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTableView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTableView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTableView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTableView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTableView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTableView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTableView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTableView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTableView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTableView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTableView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTableView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QTableView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QTableView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QTableView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTableView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QTableView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QTableView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QTableView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTableView_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QTableView) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTableView_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTableView) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTableView_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QTableView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTableView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTableView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QTableView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTableView) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QTableView_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTableView) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTableView_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTableView_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTableView_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTableView) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTableView_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QTableView) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QTableView_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QTableView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTableView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QTableView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTableView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTableView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTableView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTableView) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTableView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QTableView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QTableView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QTableView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTableView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QTableView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QTableView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QTableView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QTableView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTableView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTableView) ResetDefault() {
	purego.SyscallN(addr_QTableView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) SelectAllDefault() {
	purego.SyscallN(addr_QTableView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTableView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTableView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTableView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTableView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QTableView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QTableView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTableView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTableView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTableView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTableView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTableView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QTableView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QTableView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QTableView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTableView) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QTableView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QTableView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QTableView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QTableView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableView_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableView_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTableView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTableView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTableView_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTableView_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTableView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTableView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTableView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QTableView_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QTableView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTableView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTableView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTableView) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QTableView_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QTableView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QTableView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTableView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTableView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTableView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTableView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTableView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTableView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTableView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTableView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTableView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTableView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTableView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTableView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTableView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTableView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QTableView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTableView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTableView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTableView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTableView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTableView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTableView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTableView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTableView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTableView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTableView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTableView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTableView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTableView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTableView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTableView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTableView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTableView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QTableView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTableView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTableView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTableView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTableView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTableView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTableView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTableView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTableView) Disconnect(handle int64) {
	pg_QTableView__disconnect(handle)
	unregisterCallback(handle)
}

func QTableView_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 4283922507359547360:
		return 1
	case -1574077111192992763:
		return 2
	case -4689322347236366780:
		return 3
	case -1888629738259754754:
		return 4
	case -1267696851865372997:
		return 5
	case -4832340528568269947:
		return 6
	case -7252823847285716802:
		return 7
	case -6118001150126564745:
		return 8
	case -7450787277202094556:
		return 9
	case 5290442163167458465:
		return 10
	case -5409927581812374498:
		return 11
	case 9107058682754591716:
		return 12
	case -1500310024801917802:
		return 13
	case 5218707810677921016:
		return 14
	case -158523882194220890:
		return 15
	case -1539266838375037305:
		return 16
	case 7063795720810552074:
		return 17
	case 740869982805721908:
		return 18
	case 1022771470898460576:
		return 19
	case 503579067252048239:
		return 20
	case -4084029023708556876:
		return 21
	case -3666792239577137638:
		return 22
	case -1155555122391250623:
		return 23
	case -785551405282638793:
		return 24
	case 6057989667089378859:
		return 25
	case -5892205462145366901:
		return 26
	case -6978583972112316034:
		return 27
	case -5041452885022114200:
		return 28
	case -3811438228839363824:
		return 29
	case 7345853097340709638:
		return 30
	case 1086335023529244512:
		return 31
	case 8758113809152574300:
		return 32
	case 4472434241583495620:
		return 33
	case -8151049665126058157:
		return 34
	case 9029941737538200621:
		return 35
	case -8726300139121449691:
		return 36
	case 6366953527806415524:
		return 37
	case 1786430177580376614:
		return 38
	case 8648484264628816806:
		return 39
	case -6021383853323633948:
		return 40
	case 6979418566323745651:
		return 41
	case 4206689019978932209:
		return 42
	case -2447490284072398302:
		return 43
	case -8112366604639385491:
		return 44
	case 5898328517505784191:
		return 45
	case 1237236456312824991:
		return 46
	case -1971465653061861519:
		return 47
	case -332090898078543319:
		return 48
	case -4255494004099806581:
		return 49
	case -4415920049297795213:
		return 50
	case -4029073905697023887:
		return 51
	case 206122715919917589:
		return 52
	case -1781545137352745978:
		return 53
	case -3139904158373725006:
		return 54
	case 975043767033552552:
		return 55
	case -385596010576048993:
		return 56
	case 1023791977679793185:
		return 57
	case -4302917285655014661:
		return 58
	case 5438819661942843222:
		return 59
	case 5217623514630858297:
		return 60
	case 7049264599837532491:
		return 61
	case 729690261457684101:
		return 62
	case 4531996381736730688:
		return 63
	case 1919988340934511316:
		return 64
	case -6698043222227006015:
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

func QTableView_enable(h uintptr, idx int32) { pg_QTableView__set_override(h, idx) }

func (this *QTableView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTableView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QTableView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QTableView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QTableView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QTableView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QTableView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QTableView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QTableView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTableView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QTableView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QTableView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTableView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTableView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QTableView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QTableView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QTableView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QTableView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QTableView) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QTableView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QTableView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTableView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QTableView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QTableView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QTableView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QTableView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QTableView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QTableView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QTableView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QTableView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QTableView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTableView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QTableView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QTableView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QTableView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QTableView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QTableView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QTableView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QTableView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QTableView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QTableView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QTableView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QTableView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QTableView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QTableView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QTableView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTableView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTableView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QTableView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QTableView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QTableView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QTableView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTableView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTableView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTableView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTableView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTableView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTableView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QTableView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QTableView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QTableView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTableView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTableView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTableView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTableView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QTableView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QTableView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QTableView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTableView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTableView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTableView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTableView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTableView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTableView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTableView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTableView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTableView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTableView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTableView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTableView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTableView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QTableView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTableView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTableView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTableView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTableView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTableView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTableView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTableView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTableView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTableView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTableView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTableView_delete(this.h)
	}
	this.UnsafeRelease()
}
