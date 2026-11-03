package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTreeView_chain uintptr
	addr_QTreeView_new uintptr
	addr_QTreeView_new2 uintptr
	addr_QTreeView_tr uintptr
	addr_QTreeView_header uintptr
	addr_QTreeView_setHeader uintptr
	addr_QTreeView_autoExpandDelay uintptr
	addr_QTreeView_setAutoExpandDelay uintptr
	addr_QTreeView_indentation uintptr
	addr_QTreeView_setIndentation uintptr
	addr_QTreeView_resetIndentation uintptr
	addr_QTreeView_rootIsDecorated uintptr
	addr_QTreeView_setRootIsDecorated uintptr
	addr_QTreeView_uniformRowHeights uintptr
	addr_QTreeView_setUniformRowHeights uintptr
	addr_QTreeView_itemsExpandable uintptr
	addr_QTreeView_setItemsExpandable uintptr
	addr_QTreeView_expandsOnDoubleClick uintptr
	addr_QTreeView_setExpandsOnDoubleClick uintptr
	addr_QTreeView_columnViewportPosition uintptr
	addr_QTreeView_columnWidth uintptr
	addr_QTreeView_setColumnWidth uintptr
	addr_QTreeView_columnAt uintptr
	addr_QTreeView_isColumnHidden uintptr
	addr_QTreeView_setColumnHidden uintptr
	addr_QTreeView_isHeaderHidden uintptr
	addr_QTreeView_setHeaderHidden uintptr
	addr_QTreeView_isRowHidden uintptr
	addr_QTreeView_setRowHidden uintptr
	addr_QTreeView_isFirstColumnSpanned uintptr
	addr_QTreeView_setFirstColumnSpanned uintptr
	addr_QTreeView_isExpanded uintptr
	addr_QTreeView_setExpanded uintptr
	addr_QTreeView_setSortingEnabled uintptr
	addr_QTreeView_isSortingEnabled uintptr
	addr_QTreeView_setAnimated uintptr
	addr_QTreeView_isAnimated uintptr
	addr_QTreeView_setAllColumnsShowFocus uintptr
	addr_QTreeView_allColumnsShowFocus uintptr
	addr_QTreeView_setWordWrap uintptr
	addr_QTreeView_wordWrap uintptr
	addr_QTreeView_setTreePosition uintptr
	addr_QTreeView_treePosition uintptr
	addr_QTreeView_indexAbove uintptr
	addr_QTreeView_indexBelow uintptr
	addr_QTreeView_expanded uintptr
	addr_QTreeView_collapsed uintptr
	addr_QTreeView_hideColumn uintptr
	addr_QTreeView_showColumn uintptr
	addr_QTreeView_expand uintptr
	addr_QTreeView_collapse uintptr
	addr_QTreeView_resizeColumnToContents uintptr
	addr_QTreeView_sortByColumn uintptr
	addr_QTreeView_expandAll uintptr
	addr_QTreeView_expandRecursively uintptr
	addr_QTreeView_collapseAll uintptr
	addr_QTreeView_expandToDepth uintptr
	addr_QTreeView_columnResized uintptr
	addr_QTreeView_columnCountChanged uintptr
	addr_QTreeView_columnMoved uintptr
	addr_QTreeView_reexpand uintptr
	addr_QTreeView_rowsRemoved uintptr
	addr_QTreeView_drawTree uintptr
	addr_QTreeView_indexRowSizeHint uintptr
	addr_QTreeView_rowHeight uintptr
	addr_QTreeView_state uintptr
	addr_QTreeView_scheduleDelayedItemsLayout uintptr
	addr_QTreeView_executeDelayedItemsLayout uintptr
	addr_QTreeView_setDirtyRegion uintptr
	addr_QTreeView_scrollDirtyRegion uintptr
	addr_QTreeView_dirtyRegionOffset uintptr
	addr_QTreeView_startAutoScroll uintptr
	addr_QTreeView_stopAutoScroll uintptr
	addr_QTreeView_doAutoScroll uintptr
	addr_QTreeView_dropIndicatorPosition uintptr
	addr_QTreeView_setViewportMargins uintptr
	addr_QTreeView_setViewportMargins2 uintptr
	addr_QTreeView_viewportMargins uintptr
	addr_QTreeView_drawFrame uintptr
	addr_QTreeView_setTabOrder uintptr
	addr_QTreeView_mouseGrabber uintptr
	addr_QTreeView_keyboardGrabber uintptr
	addr_QTreeView_find uintptr
	addr_QTreeView_createWindowContainer uintptr
	addr_QTreeView_updateMicroFocus uintptr
	addr_QTreeView_create uintptr
	addr_QTreeView_destroy uintptr
	addr_QTreeView_focusNextChild uintptr
	addr_QTreeView_focusPreviousChild uintptr
	addr_QTreeView_connect uintptr
	addr_QTreeView_connect2 uintptr
	addr_QTreeView_sender uintptr
	addr_QTreeView_senderSignalIndex uintptr
	addr_QTreeView_receivers uintptr
	addr_QTreeView_isSignalConnected uintptr
	pg_QTreeView_devicePixelRatioFScale func() float64
	pg_QTreeView_encodeMetricF func(int32, float64) int32
	pg_QTreeView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTreeView_tr2 uintptr
	addr_QTreeView_tr3 uintptr
	addr_QTreeView_expandRecursively2 uintptr
	addr_QTreeView_createWindowContainer2 uintptr
	addr_QTreeView_createWindowContainer3 uintptr
	addr_QTreeView_updateMicroFocus2 uintptr
	addr_QTreeView_create2 uintptr
	addr_QTreeView_create3 uintptr
	addr_QTreeView_create4 uintptr
	addr_QTreeView_destroy2 uintptr
	addr_QTreeView_destroy3 uintptr
	addr_QTreeView_connect3 uintptr
	addr_QTreeView_connect4 uintptr
	addr_QTreeView_metaObject_default uintptr
	addr_QTreeView_setModel_default uintptr
	addr_QTreeView_setRootIndex_default uintptr
	addr_QTreeView_setSelectionModel_default uintptr
	addr_QTreeView_keyboardSearch_default uintptr
	addr_QTreeView_visualRect_default uintptr
	addr_QTreeView_scrollTo_default uintptr
	addr_QTreeView_indexAt_default uintptr
	addr_QTreeView_doItemsLayout_default uintptr
	addr_QTreeView_reset_default uintptr
	addr_QTreeView_selectAll_default uintptr
	addr_QTreeView_verticalScrollbarValueChanged_default uintptr
	addr_QTreeView_scrollContentsBy_default uintptr
	addr_QTreeView_rowsInserted_default uintptr
	addr_QTreeView_rowsAboutToBeRemoved_default uintptr
	addr_QTreeView_moveCursor_default uintptr
	addr_QTreeView_horizontalOffset_default uintptr
	addr_QTreeView_verticalOffset_default uintptr
	addr_QTreeView_setSelection_default uintptr
	addr_QTreeView_visualRegionForSelection_default uintptr
	addr_QTreeView_selectedIndexes_default uintptr
	addr_QTreeView_changeEvent_default uintptr
	addr_QTreeView_timerEvent_default uintptr
	addr_QTreeView_paintEvent_default uintptr
	addr_QTreeView_drawRow_default uintptr
	addr_QTreeView_drawBranches_default uintptr
	addr_QTreeView_mousePressEvent_default uintptr
	addr_QTreeView_mouseReleaseEvent_default uintptr
	addr_QTreeView_mouseDoubleClickEvent_default uintptr
	addr_QTreeView_mouseMoveEvent_default uintptr
	addr_QTreeView_keyPressEvent_default uintptr
	addr_QTreeView_dragMoveEvent_default uintptr
	addr_QTreeView_viewportEvent_default uintptr
	addr_QTreeView_updateGeometries_default uintptr
	addr_QTreeView_viewportSizeHint_default uintptr
	addr_QTreeView_sizeHintForColumn_default uintptr
	addr_QTreeView_horizontalScrollbarAction_default uintptr
	addr_QTreeView_isIndexHidden_default uintptr
	addr_QTreeView_selectionChanged_default uintptr
	addr_QTreeView_currentChanged_default uintptr
	addr_QTreeView_sizeHintForRow_default uintptr
	addr_QTreeView_itemDelegateForIndex_default uintptr
	addr_QTreeView_inputMethodQuery_default uintptr
	addr_QTreeView_updateEditorData_default uintptr
	addr_QTreeView_updateEditorGeometries_default uintptr
	addr_QTreeView_verticalScrollbarAction_default uintptr
	addr_QTreeView_horizontalScrollbarValueChanged_default uintptr
	addr_QTreeView_closeEditor_default uintptr
	addr_QTreeView_commitData_default uintptr
	addr_QTreeView_editorDestroyed_default uintptr
	addr_QTreeView_edit_default uintptr
	addr_QTreeView_selectionCommand_default uintptr
	addr_QTreeView_startDrag_default uintptr
	addr_QTreeView_initViewItemOption_default uintptr
	addr_QTreeView_focusNextPrevChild_default uintptr
	addr_QTreeView_event_default uintptr
	addr_QTreeView_dragEnterEvent_default uintptr
	addr_QTreeView_dragLeaveEvent_default uintptr
	addr_QTreeView_dropEvent_default uintptr
	addr_QTreeView_focusInEvent_default uintptr
	addr_QTreeView_focusOutEvent_default uintptr
	addr_QTreeView_resizeEvent_default uintptr
	addr_QTreeView_inputMethodEvent_default uintptr
	addr_QTreeView_eventFilter_default uintptr
	addr_QTreeView_minimumSizeHint_default uintptr
	addr_QTreeView_sizeHint_default uintptr
	addr_QTreeView_setupViewport_default uintptr
	addr_QTreeView_wheelEvent_default uintptr
	addr_QTreeView_contextMenuEvent_default uintptr
	addr_QTreeView_initStyleOption_default uintptr
	addr_QTreeView_devType_default uintptr
	addr_QTreeView_setVisible_default uintptr
	addr_QTreeView_heightForWidth_default uintptr
	addr_QTreeView_hasHeightForWidth_default uintptr
	addr_QTreeView_paintEngine_default uintptr
	addr_QTreeView_keyReleaseEvent_default uintptr
	addr_QTreeView_enterEvent_default uintptr
	addr_QTreeView_leaveEvent_default uintptr
	addr_QTreeView_moveEvent_default uintptr
	addr_QTreeView_closeEvent_default uintptr
	addr_QTreeView_tabletEvent_default uintptr
	addr_QTreeView_actionEvent_default uintptr
	addr_QTreeView_showEvent_default uintptr
	addr_QTreeView_hideEvent_default uintptr
	addr_QTreeView_metric_default uintptr
	addr_QTreeView_initPainter_default uintptr
	addr_QTreeView_redirected_default uintptr
	addr_QTreeView_sharedPainter_default uintptr
	addr_QTreeView_childEvent_default uintptr
	addr_QTreeView_customEvent_default uintptr
	addr_QTreeView_connectNotify_default uintptr
	addr_QTreeView_disconnectNotify_default uintptr
	pg_QTreeView_connect_expanded func(uintptr, int64)
	pg_QTreeView_connect_collapsed func(uintptr, int64)
	pg_QTreeView_connect_pressed func(uintptr, int64)
	pg_QTreeView_connect_clicked func(uintptr, int64)
	pg_QTreeView_connect_doubleClicked func(uintptr, int64)
	pg_QTreeView_connect_activated func(uintptr, int64)
	pg_QTreeView_connect_entered func(uintptr, int64)
	pg_QTreeView_connect_viewportEntered func(uintptr, int64)
	pg_QTreeView_connect_iconSizeChanged func(uintptr, int64)
	pg_QTreeView_connect_windowTitleChanged func(uintptr, int64)
	pg_QTreeView_connect_windowIconChanged func(uintptr, int64)
	pg_QTreeView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTreeView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTreeView_connect_destroyed func(uintptr, int64)
	pg_QTreeView_connect_destroyed2 func(uintptr, int64)
	pg_QTreeView_connect_objectNameChanged func(uintptr, int64)
	pg_QTreeView__disconnect func(int64)
	pg_QTreeView__set_override func(uintptr, int32)
	pg_QTreeView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTreeView_chain, "QTreeView_chain")
	symaddr(&addr_QTreeView_new, "QTreeView_new")
	symaddr(&addr_QTreeView_new2, "QTreeView_new2")
	symaddr(&addr_QTreeView_tr, "QTreeView_tr")
	symaddr(&addr_QTreeView_header, "QTreeView_header")
	symaddr(&addr_QTreeView_setHeader, "QTreeView_setHeader")
	symaddr(&addr_QTreeView_autoExpandDelay, "QTreeView_autoExpandDelay")
	symaddr(&addr_QTreeView_setAutoExpandDelay, "QTreeView_setAutoExpandDelay")
	symaddr(&addr_QTreeView_indentation, "QTreeView_indentation")
	symaddr(&addr_QTreeView_setIndentation, "QTreeView_setIndentation")
	symaddr(&addr_QTreeView_resetIndentation, "QTreeView_resetIndentation")
	symaddr(&addr_QTreeView_rootIsDecorated, "QTreeView_rootIsDecorated")
	symaddr(&addr_QTreeView_setRootIsDecorated, "QTreeView_setRootIsDecorated")
	symaddr(&addr_QTreeView_uniformRowHeights, "QTreeView_uniformRowHeights")
	symaddr(&addr_QTreeView_setUniformRowHeights, "QTreeView_setUniformRowHeights")
	symaddr(&addr_QTreeView_itemsExpandable, "QTreeView_itemsExpandable")
	symaddr(&addr_QTreeView_setItemsExpandable, "QTreeView_setItemsExpandable")
	symaddr(&addr_QTreeView_expandsOnDoubleClick, "QTreeView_expandsOnDoubleClick")
	symaddr(&addr_QTreeView_setExpandsOnDoubleClick, "QTreeView_setExpandsOnDoubleClick")
	symaddr(&addr_QTreeView_columnViewportPosition, "QTreeView_columnViewportPosition")
	symaddr(&addr_QTreeView_columnWidth, "QTreeView_columnWidth")
	symaddr(&addr_QTreeView_setColumnWidth, "QTreeView_setColumnWidth")
	symaddr(&addr_QTreeView_columnAt, "QTreeView_columnAt")
	symaddr(&addr_QTreeView_isColumnHidden, "QTreeView_isColumnHidden")
	symaddr(&addr_QTreeView_setColumnHidden, "QTreeView_setColumnHidden")
	symaddr(&addr_QTreeView_isHeaderHidden, "QTreeView_isHeaderHidden")
	symaddr(&addr_QTreeView_setHeaderHidden, "QTreeView_setHeaderHidden")
	symaddr(&addr_QTreeView_isRowHidden, "QTreeView_isRowHidden")
	symaddr(&addr_QTreeView_setRowHidden, "QTreeView_setRowHidden")
	symaddr(&addr_QTreeView_isFirstColumnSpanned, "QTreeView_isFirstColumnSpanned")
	symaddr(&addr_QTreeView_setFirstColumnSpanned, "QTreeView_setFirstColumnSpanned")
	symaddr(&addr_QTreeView_isExpanded, "QTreeView_isExpanded")
	symaddr(&addr_QTreeView_setExpanded, "QTreeView_setExpanded")
	symaddr(&addr_QTreeView_setSortingEnabled, "QTreeView_setSortingEnabled")
	symaddr(&addr_QTreeView_isSortingEnabled, "QTreeView_isSortingEnabled")
	symaddr(&addr_QTreeView_setAnimated, "QTreeView_setAnimated")
	symaddr(&addr_QTreeView_isAnimated, "QTreeView_isAnimated")
	symaddr(&addr_QTreeView_setAllColumnsShowFocus, "QTreeView_setAllColumnsShowFocus")
	symaddr(&addr_QTreeView_allColumnsShowFocus, "QTreeView_allColumnsShowFocus")
	symaddr(&addr_QTreeView_setWordWrap, "QTreeView_setWordWrap")
	symaddr(&addr_QTreeView_wordWrap, "QTreeView_wordWrap")
	symaddr(&addr_QTreeView_setTreePosition, "QTreeView_setTreePosition")
	symaddr(&addr_QTreeView_treePosition, "QTreeView_treePosition")
	symaddr(&addr_QTreeView_indexAbove, "QTreeView_indexAbove")
	symaddr(&addr_QTreeView_indexBelow, "QTreeView_indexBelow")
	symaddr(&addr_QTreeView_expanded, "QTreeView_expanded")
	symaddr(&addr_QTreeView_collapsed, "QTreeView_collapsed")
	symaddr(&addr_QTreeView_hideColumn, "QTreeView_hideColumn")
	symaddr(&addr_QTreeView_showColumn, "QTreeView_showColumn")
	symaddr(&addr_QTreeView_expand, "QTreeView_expand")
	symaddr(&addr_QTreeView_collapse, "QTreeView_collapse")
	symaddr(&addr_QTreeView_resizeColumnToContents, "QTreeView_resizeColumnToContents")
	symaddr(&addr_QTreeView_sortByColumn, "QTreeView_sortByColumn")
	symaddr(&addr_QTreeView_expandAll, "QTreeView_expandAll")
	symaddr(&addr_QTreeView_expandRecursively, "QTreeView_expandRecursively")
	symaddr(&addr_QTreeView_collapseAll, "QTreeView_collapseAll")
	symaddr(&addr_QTreeView_expandToDepth, "QTreeView_expandToDepth")
	symaddr(&addr_QTreeView_columnResized, "QTreeView_columnResized")
	symaddr(&addr_QTreeView_columnCountChanged, "QTreeView_columnCountChanged")
	symaddr(&addr_QTreeView_columnMoved, "QTreeView_columnMoved")
	symaddr(&addr_QTreeView_reexpand, "QTreeView_reexpand")
	symaddr(&addr_QTreeView_rowsRemoved, "QTreeView_rowsRemoved")
	symaddr(&addr_QTreeView_drawTree, "QTreeView_drawTree")
	symaddr(&addr_QTreeView_indexRowSizeHint, "QTreeView_indexRowSizeHint")
	symaddr(&addr_QTreeView_rowHeight, "QTreeView_rowHeight")
	symaddr(&addr_QTreeView_state, "QTreeView_state")
	symaddr(&addr_QTreeView_scheduleDelayedItemsLayout, "QTreeView_scheduleDelayedItemsLayout")
	symaddr(&addr_QTreeView_executeDelayedItemsLayout, "QTreeView_executeDelayedItemsLayout")
	symaddr(&addr_QTreeView_setDirtyRegion, "QTreeView_setDirtyRegion")
	symaddr(&addr_QTreeView_scrollDirtyRegion, "QTreeView_scrollDirtyRegion")
	symaddr(&addr_QTreeView_dirtyRegionOffset, "QTreeView_dirtyRegionOffset")
	symaddr(&addr_QTreeView_startAutoScroll, "QTreeView_startAutoScroll")
	symaddr(&addr_QTreeView_stopAutoScroll, "QTreeView_stopAutoScroll")
	symaddr(&addr_QTreeView_doAutoScroll, "QTreeView_doAutoScroll")
	symaddr(&addr_QTreeView_dropIndicatorPosition, "QTreeView_dropIndicatorPosition")
	symaddr(&addr_QTreeView_setViewportMargins, "QTreeView_setViewportMargins")
	symaddr(&addr_QTreeView_setViewportMargins2, "QTreeView_setViewportMargins2")
	symaddr(&addr_QTreeView_viewportMargins, "QTreeView_viewportMargins")
	symaddr(&addr_QTreeView_drawFrame, "QTreeView_drawFrame")
	symaddr(&addr_QTreeView_setTabOrder, "QTreeView_setTabOrder")
	symaddr(&addr_QTreeView_mouseGrabber, "QTreeView_mouseGrabber")
	symaddr(&addr_QTreeView_keyboardGrabber, "QTreeView_keyboardGrabber")
	symaddr(&addr_QTreeView_find, "QTreeView_find")
	symaddr(&addr_QTreeView_createWindowContainer, "QTreeView_createWindowContainer")
	symaddr(&addr_QTreeView_updateMicroFocus, "QTreeView_updateMicroFocus")
	symaddr(&addr_QTreeView_create, "QTreeView_create")
	symaddr(&addr_QTreeView_destroy, "QTreeView_destroy")
	symaddr(&addr_QTreeView_focusNextChild, "QTreeView_focusNextChild")
	symaddr(&addr_QTreeView_focusPreviousChild, "QTreeView_focusPreviousChild")
	symaddr(&addr_QTreeView_connect, "QTreeView_connect")
	symaddr(&addr_QTreeView_connect2, "QTreeView_connect2")
	symaddr(&addr_QTreeView_sender, "QTreeView_sender")
	symaddr(&addr_QTreeView_senderSignalIndex, "QTreeView_senderSignalIndex")
	symaddr(&addr_QTreeView_receivers, "QTreeView_receivers")
	symaddr(&addr_QTreeView_isSignalConnected, "QTreeView_isSignalConnected")
	reg(&pg_QTreeView_devicePixelRatioFScale, "QTreeView_devicePixelRatioFScale")
	reg(&pg_QTreeView_encodeMetricF, "QTreeView_encodeMetricF")
	reg(&pg_QTreeView_getDecodedMetricF, "QTreeView_getDecodedMetricF")
	symaddr(&addr_QTreeView_tr2, "QTreeView_tr2")
	symaddr(&addr_QTreeView_tr3, "QTreeView_tr3")
	symaddr(&addr_QTreeView_expandRecursively2, "QTreeView_expandRecursively2")
	symaddr(&addr_QTreeView_createWindowContainer2, "QTreeView_createWindowContainer2")
	symaddr(&addr_QTreeView_createWindowContainer3, "QTreeView_createWindowContainer3")
	symaddr(&addr_QTreeView_updateMicroFocus2, "QTreeView_updateMicroFocus2")
	symaddr(&addr_QTreeView_create2, "QTreeView_create2")
	symaddr(&addr_QTreeView_create3, "QTreeView_create3")
	symaddr(&addr_QTreeView_create4, "QTreeView_create4")
	symaddr(&addr_QTreeView_destroy2, "QTreeView_destroy2")
	symaddr(&addr_QTreeView_destroy3, "QTreeView_destroy3")
	symaddr(&addr_QTreeView_connect3, "QTreeView_connect3")
	symaddr(&addr_QTreeView_connect4, "QTreeView_connect4")
	symaddr(&addr_QTreeView_metaObject_default, "QTreeView_metaObject_default")
	symaddr(&addr_QTreeView_setModel_default, "QTreeView_setModel_default")
	symaddr(&addr_QTreeView_setRootIndex_default, "QTreeView_setRootIndex_default")
	symaddr(&addr_QTreeView_setSelectionModel_default, "QTreeView_setSelectionModel_default")
	symaddr(&addr_QTreeView_keyboardSearch_default, "QTreeView_keyboardSearch_default")
	symaddr(&addr_QTreeView_visualRect_default, "QTreeView_visualRect_default")
	symaddr(&addr_QTreeView_scrollTo_default, "QTreeView_scrollTo_default")
	symaddr(&addr_QTreeView_indexAt_default, "QTreeView_indexAt_default")
	symaddr(&addr_QTreeView_doItemsLayout_default, "QTreeView_doItemsLayout_default")
	symaddr(&addr_QTreeView_reset_default, "QTreeView_reset_default")
	symaddr(&addr_QTreeView_selectAll_default, "QTreeView_selectAll_default")
	symaddr(&addr_QTreeView_verticalScrollbarValueChanged_default, "QTreeView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QTreeView_scrollContentsBy_default, "QTreeView_scrollContentsBy_default")
	symaddr(&addr_QTreeView_rowsInserted_default, "QTreeView_rowsInserted_default")
	symaddr(&addr_QTreeView_rowsAboutToBeRemoved_default, "QTreeView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QTreeView_moveCursor_default, "QTreeView_moveCursor_default")
	symaddr(&addr_QTreeView_horizontalOffset_default, "QTreeView_horizontalOffset_default")
	symaddr(&addr_QTreeView_verticalOffset_default, "QTreeView_verticalOffset_default")
	symaddr(&addr_QTreeView_setSelection_default, "QTreeView_setSelection_default")
	symaddr(&addr_QTreeView_visualRegionForSelection_default, "QTreeView_visualRegionForSelection_default")
	symaddr(&addr_QTreeView_selectedIndexes_default, "QTreeView_selectedIndexes_default")
	symaddr(&addr_QTreeView_changeEvent_default, "QTreeView_changeEvent_default")
	symaddr(&addr_QTreeView_timerEvent_default, "QTreeView_timerEvent_default")
	symaddr(&addr_QTreeView_paintEvent_default, "QTreeView_paintEvent_default")
	symaddr(&addr_QTreeView_drawRow_default, "QTreeView_drawRow_default")
	symaddr(&addr_QTreeView_drawBranches_default, "QTreeView_drawBranches_default")
	symaddr(&addr_QTreeView_mousePressEvent_default, "QTreeView_mousePressEvent_default")
	symaddr(&addr_QTreeView_mouseReleaseEvent_default, "QTreeView_mouseReleaseEvent_default")
	symaddr(&addr_QTreeView_mouseDoubleClickEvent_default, "QTreeView_mouseDoubleClickEvent_default")
	symaddr(&addr_QTreeView_mouseMoveEvent_default, "QTreeView_mouseMoveEvent_default")
	symaddr(&addr_QTreeView_keyPressEvent_default, "QTreeView_keyPressEvent_default")
	symaddr(&addr_QTreeView_dragMoveEvent_default, "QTreeView_dragMoveEvent_default")
	symaddr(&addr_QTreeView_viewportEvent_default, "QTreeView_viewportEvent_default")
	symaddr(&addr_QTreeView_updateGeometries_default, "QTreeView_updateGeometries_default")
	symaddr(&addr_QTreeView_viewportSizeHint_default, "QTreeView_viewportSizeHint_default")
	symaddr(&addr_QTreeView_sizeHintForColumn_default, "QTreeView_sizeHintForColumn_default")
	symaddr(&addr_QTreeView_horizontalScrollbarAction_default, "QTreeView_horizontalScrollbarAction_default")
	symaddr(&addr_QTreeView_isIndexHidden_default, "QTreeView_isIndexHidden_default")
	symaddr(&addr_QTreeView_selectionChanged_default, "QTreeView_selectionChanged_default")
	symaddr(&addr_QTreeView_currentChanged_default, "QTreeView_currentChanged_default")
	symaddr(&addr_QTreeView_sizeHintForRow_default, "QTreeView_sizeHintForRow_default")
	symaddr(&addr_QTreeView_itemDelegateForIndex_default, "QTreeView_itemDelegateForIndex_default")
	symaddr(&addr_QTreeView_inputMethodQuery_default, "QTreeView_inputMethodQuery_default")
	symaddr(&addr_QTreeView_updateEditorData_default, "QTreeView_updateEditorData_default")
	symaddr(&addr_QTreeView_updateEditorGeometries_default, "QTreeView_updateEditorGeometries_default")
	symaddr(&addr_QTreeView_verticalScrollbarAction_default, "QTreeView_verticalScrollbarAction_default")
	symaddr(&addr_QTreeView_horizontalScrollbarValueChanged_default, "QTreeView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QTreeView_closeEditor_default, "QTreeView_closeEditor_default")
	symaddr(&addr_QTreeView_commitData_default, "QTreeView_commitData_default")
	symaddr(&addr_QTreeView_editorDestroyed_default, "QTreeView_editorDestroyed_default")
	symaddr(&addr_QTreeView_edit_default, "QTreeView_edit_default")
	symaddr(&addr_QTreeView_selectionCommand_default, "QTreeView_selectionCommand_default")
	symaddr(&addr_QTreeView_startDrag_default, "QTreeView_startDrag_default")
	symaddr(&addr_QTreeView_initViewItemOption_default, "QTreeView_initViewItemOption_default")
	symaddr(&addr_QTreeView_focusNextPrevChild_default, "QTreeView_focusNextPrevChild_default")
	symaddr(&addr_QTreeView_event_default, "QTreeView_event_default")
	symaddr(&addr_QTreeView_dragEnterEvent_default, "QTreeView_dragEnterEvent_default")
	symaddr(&addr_QTreeView_dragLeaveEvent_default, "QTreeView_dragLeaveEvent_default")
	symaddr(&addr_QTreeView_dropEvent_default, "QTreeView_dropEvent_default")
	symaddr(&addr_QTreeView_focusInEvent_default, "QTreeView_focusInEvent_default")
	symaddr(&addr_QTreeView_focusOutEvent_default, "QTreeView_focusOutEvent_default")
	symaddr(&addr_QTreeView_resizeEvent_default, "QTreeView_resizeEvent_default")
	symaddr(&addr_QTreeView_inputMethodEvent_default, "QTreeView_inputMethodEvent_default")
	symaddr(&addr_QTreeView_eventFilter_default, "QTreeView_eventFilter_default")
	symaddr(&addr_QTreeView_minimumSizeHint_default, "QTreeView_minimumSizeHint_default")
	symaddr(&addr_QTreeView_sizeHint_default, "QTreeView_sizeHint_default")
	symaddr(&addr_QTreeView_setupViewport_default, "QTreeView_setupViewport_default")
	symaddr(&addr_QTreeView_wheelEvent_default, "QTreeView_wheelEvent_default")
	symaddr(&addr_QTreeView_contextMenuEvent_default, "QTreeView_contextMenuEvent_default")
	symaddr(&addr_QTreeView_initStyleOption_default, "QTreeView_initStyleOption_default")
	symaddr(&addr_QTreeView_devType_default, "QTreeView_devType_default")
	symaddr(&addr_QTreeView_setVisible_default, "QTreeView_setVisible_default")
	symaddr(&addr_QTreeView_heightForWidth_default, "QTreeView_heightForWidth_default")
	symaddr(&addr_QTreeView_hasHeightForWidth_default, "QTreeView_hasHeightForWidth_default")
	symaddr(&addr_QTreeView_paintEngine_default, "QTreeView_paintEngine_default")
	symaddr(&addr_QTreeView_keyReleaseEvent_default, "QTreeView_keyReleaseEvent_default")
	symaddr(&addr_QTreeView_enterEvent_default, "QTreeView_enterEvent_default")
	symaddr(&addr_QTreeView_leaveEvent_default, "QTreeView_leaveEvent_default")
	symaddr(&addr_QTreeView_moveEvent_default, "QTreeView_moveEvent_default")
	symaddr(&addr_QTreeView_closeEvent_default, "QTreeView_closeEvent_default")
	symaddr(&addr_QTreeView_tabletEvent_default, "QTreeView_tabletEvent_default")
	symaddr(&addr_QTreeView_actionEvent_default, "QTreeView_actionEvent_default")
	symaddr(&addr_QTreeView_showEvent_default, "QTreeView_showEvent_default")
	symaddr(&addr_QTreeView_hideEvent_default, "QTreeView_hideEvent_default")
	symaddr(&addr_QTreeView_metric_default, "QTreeView_metric_default")
	symaddr(&addr_QTreeView_initPainter_default, "QTreeView_initPainter_default")
	symaddr(&addr_QTreeView_redirected_default, "QTreeView_redirected_default")
	symaddr(&addr_QTreeView_sharedPainter_default, "QTreeView_sharedPainter_default")
	symaddr(&addr_QTreeView_childEvent_default, "QTreeView_childEvent_default")
	symaddr(&addr_QTreeView_customEvent_default, "QTreeView_customEvent_default")
	symaddr(&addr_QTreeView_connectNotify_default, "QTreeView_connectNotify_default")
	symaddr(&addr_QTreeView_disconnectNotify_default, "QTreeView_disconnectNotify_default")
	reg(&pg_QTreeView_connect_expanded, "QTreeView_connect_expanded")
	reg(&pg_QTreeView_connect_collapsed, "QTreeView_connect_collapsed")
	reg(&pg_QTreeView_connect_pressed, "QTreeView_connect_pressed")
	reg(&pg_QTreeView_connect_clicked, "QTreeView_connect_clicked")
	reg(&pg_QTreeView_connect_doubleClicked, "QTreeView_connect_doubleClicked")
	reg(&pg_QTreeView_connect_activated, "QTreeView_connect_activated")
	reg(&pg_QTreeView_connect_entered, "QTreeView_connect_entered")
	reg(&pg_QTreeView_connect_viewportEntered, "QTreeView_connect_viewportEntered")
	reg(&pg_QTreeView_connect_iconSizeChanged, "QTreeView_connect_iconSizeChanged")
	reg(&pg_QTreeView_connect_windowTitleChanged, "QTreeView_connect_windowTitleChanged")
	reg(&pg_QTreeView_connect_windowIconChanged, "QTreeView_connect_windowIconChanged")
	reg(&pg_QTreeView_connect_windowIconTextChanged, "QTreeView_connect_windowIconTextChanged")
	reg(&pg_QTreeView_connect_customContextMenuRequested, "QTreeView_connect_customContextMenuRequested")
	reg(&pg_QTreeView_connect_destroyed, "QTreeView_connect_destroyed")
	reg(&pg_QTreeView_connect_destroyed2, "QTreeView_connect_destroyed2")
	reg(&pg_QTreeView_connect_objectNameChanged, "QTreeView_connect_objectNameChanged")
	reg(&pg_QTreeView__disconnect, "QTreeView__disconnect")
	reg(&pg_QTreeView__set_override, "QTreeView__set_override")
	reg(&pg_QTreeView_delete, "QTreeView_delete")
}

type QTreeView struct {
	h uintptr
	*QAbstractItemView
}

func bindQTreeView(h uintptr) *QTreeView {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QTreeView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTreeView{h: h, QAbstractItemView: &QAbstractItemView{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQTreeView(h uintptr) *QTreeView { return bindQTreeView(h) }

func (this *QTreeView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTreeView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTreeView(h unsafe.Pointer) *QTreeView { return bindQTreeView(uintptr(h)) }
func UnsafeNewQTreeViewOwned(h unsafe.Pointer) *QTreeView { return newownedQTreeView(uintptr(h)) }

func (this *QTreeView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemView != nil {
		this.QAbstractItemView.UnsafeRelease()
	}
}

func NewQTreeView(parent *QWidget) *QTreeView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTreeView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeView_virtIdx, QTreeView_enable)
	return __o
}

func NewQTreeView2() *QTreeView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeView_new2, uintptr(__goid))
	__o := newownedQTreeView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeView_virtIdx, QTreeView_enable)
	return __o
}

func QTreeView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTreeView) Header() *QHeaderView {
	__r, _, _ := purego.SyscallN(addr_QTreeView_header, this.h)
	runtime.KeepAlive(this)
	return bindQHeaderView(__r)
}

func (this *QTreeView) SetHeader(header *QHeaderView) {
	purego.SyscallN(addr_QTreeView_setHeader, this.h, uintptr(header.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(header)
}

func (this *QTreeView) AutoExpandDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_autoExpandDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) SetAutoExpandDelay(delay int) {
	purego.SyscallN(addr_QTreeView_setAutoExpandDelay, this.h, uintptr(delay))
	runtime.KeepAlive(this)
}

func (this *QTreeView) Indentation() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_indentation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) SetIndentation(i int) {
	purego.SyscallN(addr_QTreeView_setIndentation, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ResetIndentation() {
	purego.SyscallN(addr_QTreeView_resetIndentation, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) RootIsDecorated() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_rootIsDecorated, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetRootIsDecorated(show bool) {
	purego.SyscallN(addr_QTreeView_setRootIsDecorated, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) UniformRowHeights() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_uniformRowHeights, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetUniformRowHeights(uniform bool) {
	purego.SyscallN(addr_QTreeView_setUniformRowHeights, this.h, uintptr(b2i(uniform)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ItemsExpandable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_itemsExpandable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetItemsExpandable(enable bool) {
	purego.SyscallN(addr_QTreeView_setItemsExpandable, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ExpandsOnDoubleClick() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_expandsOnDoubleClick, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetExpandsOnDoubleClick(enable bool) {
	purego.SyscallN(addr_QTreeView_setExpandsOnDoubleClick, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ColumnViewportPosition(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_columnViewportPosition, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) ColumnWidth(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_columnWidth, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) SetColumnWidth(column int, width int) {
	purego.SyscallN(addr_QTreeView_setColumnWidth, this.h, uintptr(column), uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ColumnAt(x int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_columnAt, this.h, uintptr(x))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) IsColumnHidden(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isColumnHidden, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetColumnHidden(column int, hide bool) {
	purego.SyscallN(addr_QTreeView_setColumnHidden, this.h, uintptr(column), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) IsHeaderHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isHeaderHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetHeaderHidden(hide bool) {
	purego.SyscallN(addr_QTreeView_setHeaderHidden, this.h, uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) IsRowHidden(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isRowHidden, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTreeView) SetRowHidden(row int, parent *QModelIndex, hide bool) {
	purego.SyscallN(addr_QTreeView_setRowHidden, this.h, uintptr(row), uintptr(parent.UnsafePointer()), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeView) IsFirstColumnSpanned(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isFirstColumnSpanned, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTreeView) SetFirstColumnSpanned(row int, parent *QModelIndex, span bool) {
	purego.SyscallN(addr_QTreeView_setFirstColumnSpanned, this.h, uintptr(row), uintptr(parent.UnsafePointer()), uintptr(b2i(span)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeView) IsExpanded(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isExpanded, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTreeView) SetExpanded(index *QModelIndex, expand bool) {
	purego.SyscallN(addr_QTreeView_setExpanded, this.h, uintptr(index.UnsafePointer()), uintptr(b2i(expand)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) SetSortingEnabled(enable bool) {
	purego.SyscallN(addr_QTreeView_setSortingEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) IsSortingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isSortingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetAnimated(enable bool) {
	purego.SyscallN(addr_QTreeView_setAnimated, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) IsAnimated() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isAnimated, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetAllColumnsShowFocus(enable bool) {
	purego.SyscallN(addr_QTreeView_setAllColumnsShowFocus, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) AllColumnsShowFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_allColumnsShowFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetWordWrap(on bool) {
	purego.SyscallN(addr_QTreeView_setWordWrap, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) WordWrap() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_wordWrap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) SetTreePosition(logicalIndex int) {
	purego.SyscallN(addr_QTreeView_setTreePosition, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QTreeView) TreePosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_treePosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) IndexAbove(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeView_indexAbove, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QTreeView) IndexBelow(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeView_indexBelow, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QTreeView) Expanded(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_expanded, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) Collapsed(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_collapsed, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) HideColumn(column int) {
	purego.SyscallN(addr_QTreeView_hideColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ShowColumn(column int) {
	purego.SyscallN(addr_QTreeView_showColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTreeView) Expand(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_expand, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) Collapse(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_collapse, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) ResizeColumnToContents(column int) {
	purego.SyscallN(addr_QTreeView_resizeColumnToContents, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QTreeView) SortByColumn(column int, order SortOrder) {
	purego.SyscallN(addr_QTreeView_sortByColumn, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ExpandAll() {
	purego.SyscallN(addr_QTreeView_expandAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) ExpandRecursively(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_expandRecursively, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) CollapseAll() {
	purego.SyscallN(addr_QTreeView_collapseAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) ExpandToDepth(depth int) {
	purego.SyscallN(addr_QTreeView_expandToDepth, this.h, uintptr(depth))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ColumnResized(column int, oldSize int, newSize int) {
	purego.SyscallN(addr_QTreeView_columnResized, this.h, uintptr(column), uintptr(oldSize), uintptr(newSize))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ColumnCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTreeView_columnCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ColumnMoved() {
	purego.SyscallN(addr_QTreeView_columnMoved, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) Reexpand() {
	purego.SyscallN(addr_QTreeView_reexpand, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) RowsRemoved(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTreeView_rowsRemoved, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeView) DrawTree(painter *QPainter, region *QRegion) {
	purego.SyscallN(addr_QTreeView_drawTree, this.h, uintptr(painter.UnsafePointer()), uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(region)
}

func (this *QTreeView) IndexRowSizeHint(index *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_indexRowSizeHint, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return int(__r)
}

func (this *QTreeView) RowHeight(index *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_rowHeight, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return int(__r)
}

func (this *QTreeView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QTreeView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QTreeView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QTreeView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QTreeView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QTreeView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QTreeView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTreeView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTreeView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTreeView) StartAutoScroll() {
	purego.SyscallN(addr_QTreeView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) StopAutoScroll() {
	purego.SyscallN(addr_QTreeView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) DoAutoScroll() {
	purego.SyscallN(addr_QTreeView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QTreeView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTreeView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTreeView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTreeView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTreeView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTreeView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTreeView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTreeView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTreeView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTreeView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTreeView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_mouseGrabber)
	return bindQWidget(__r)
}

func QTreeView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_keyboardGrabber)
	return bindQWidget(__r)
}

func QTreeView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTreeView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTreeView) UpdateMicroFocus() {
	purego.SyscallN(addr_QTreeView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) Create() {
	purego.SyscallN(addr_QTreeView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) Destroy() {
	purego.SyscallN(addr_QTreeView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTreeView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTreeView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTreeView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTreeView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTreeView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTreeView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTreeView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTreeView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTreeView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTreeView_DevicePixelRatioFScale() float64 {
	__r := pg_QTreeView_devicePixelRatioFScale()
	return __r
}

func QTreeView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTreeView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTreeView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTreeView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTreeView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTreeView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTreeView) ExpandRecursively2(index *QModelIndex, depth int) {
	purego.SyscallN(addr_QTreeView_expandRecursively2, this.h, uintptr(index.UnsafePointer()), uintptr(depth))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func QTreeView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTreeView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTreeView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTreeView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTreeView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTreeView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTreeView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTreeView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTreeView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTreeView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTreeView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTreeView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTreeView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTreeView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTreeView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTreeView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTreeView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTreeView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QTreeView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QTreeView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QTreeView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QTreeView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QTreeView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QTreeView) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTreeView_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QTreeView) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTreeView_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeView) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeView_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QTreeView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QTreeView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) ResetDefault() {
	purego.SyscallN(addr_QTreeView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) SelectAllDefault() {
	purego.SyscallN(addr_QTreeView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTreeView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTreeView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTreeView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTreeView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTreeView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTreeView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeView_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTreeView) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTreeView_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QTreeView) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QTreeView_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QTreeView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeView) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeView_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTreeView_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QTreeView_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) DrawRowDefault(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_drawRow_default, this.h, uintptr(painter.UnsafePointer()), uintptr(options.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(options)
	runtime.KeepAlive(index)
}

func (this *QTreeView) DrawBranchesDefault(painter *QPainter, rect *QRect, index *QModelIndex) {
	purego.SyscallN(addr_QTreeView_drawBranches_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(index)
}

func (this *QTreeView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeView_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeView_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTreeView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTreeView_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QTreeView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTreeView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTreeView) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTreeView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QTreeView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QTreeView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QTreeView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTreeView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QTreeView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QTreeView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTreeView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTreeView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QTreeView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QTreeView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTreeView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTreeView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTreeView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTreeView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QTreeView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QTreeView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QTreeView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeView) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QTreeView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QTreeView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QTreeView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QTreeView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTreeView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTreeView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTreeView_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTreeView_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTreeView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTreeView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QTreeView_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QTreeView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTreeView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTreeView) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QTreeView_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTreeView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QTreeView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTreeView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTreeView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTreeView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTreeView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTreeView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTreeView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTreeView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTreeView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTreeView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTreeView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTreeView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTreeView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTreeView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QTreeView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTreeView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTreeView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTreeView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTreeView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTreeView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTreeView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTreeView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTreeView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTreeView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTreeView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTreeView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTreeView) OnExpanded(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_expanded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnCollapsed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_collapsed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTreeView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QTreeView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTreeView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTreeView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTreeView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTreeView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTreeView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTreeView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTreeView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeView) Disconnect(handle int64) {
	pg_QTreeView__disconnect(handle)
	unregisterCallback(handle)
}

func QTreeView_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 4283922507359547360:
		return 1
	case -1574077111192992763:
		return 2
	case -4689322347236366780:
		return 3
	case -5041452885022114200:
		return 4
	case -1267696851865372997:
		return 5
	case -4832340528568269947:
		return 6
	case -7252823847285716802:
		return 7
	case -1888629738259754754:
		return 8
	case 1086335023529244512:
		return 9
	case 8758113809152574300:
		return 10
	case 6366953527806415524:
		return 11
	case -6118001150126564745:
		return 12
	case 4472434241583495620:
		return 13
	case -8151049665126058157:
		return 14
	case -158523882194220890:
		return 15
	case -1500310024801917802:
		return 16
	case 5218707810677921016:
		return 17
	case -1539266838375037305:
		return 18
	case 7063795720810552074:
		return 19
	case 740869982805721908:
		return 20
	case 4959376063826165943:
		return 21
	case -5409927581812374498:
		return 22
	case 5290442163167458465:
		return 23
	case -3950604704165452999:
		return 24
	case 6369455307610057755:
		return 25
	case -332090898078543319:
		return 26
	case -4415920049297795213:
		return 27
	case -4029073905697023887:
		return 28
	case -4255494004099806581:
		return 29
	case 1023791977679793185:
		return 30
	case -1781545137352745978:
		return 31
	case -1971465653061861519:
		return 32
	case 1022771470898460576:
		return 33
	case 503579067252048239:
		return 34
	case -3666792239577137638:
		return 35
	case -785551405282638793:
		return 36
	case 6057989667089378859:
		return 37
	case -5892205462145366901:
		return 38
	case -6978583972112316034:
		return 39
	case -4084029023708556876:
		return 40
	case -3811438228839363824:
		return 41
	case 7345853097340709638:
		return 42
	case 9029941737538200621:
		return 43
	case -8726300139121449691:
		return 44
	case -1155555122391250623:
		return 45
	case 1786430177580376614:
		return 46
	case 8648484264628816806:
		return 47
	case -6021383853323633948:
		return 48
	case 6979418566323745651:
		return 49
	case 4206689019978932209:
		return 50
	case -2447490284072398302:
		return 51
	case -8112366604639385491:
		return 52
	case -7450787277202094556:
		return 53
	case 5898328517505784191:
		return 54
	case 1237236456312824991:
		return 55
	case 206122715919917589:
		return 56
	case -3139904158373725006:
		return 57
	case 9107058682754591716:
		return 58
	case 975043767033552552:
		return 59
	case -385596010576048993:
		return 60
	case -4302917285655014661:
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
	case 1919988340934511316:
		return 67
	case -6698043222227006015:
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

func QTreeView_enable(h uintptr, idx int32) { pg_QTreeView__set_override(h, idx) }

func (this *QTreeView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTreeView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QTreeView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QTreeView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QTreeView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QTreeView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QTreeView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QTreeView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QTreeView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QTreeView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QTreeView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QTreeView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QTreeView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTreeView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QTreeView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QTreeView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QTreeView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QTreeView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QTreeView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QTreeView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QTreeView) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QTreeView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QTreeView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTreeView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QTreeView) OnDrawRow(fn func(*QPainter, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -3950604704165452999, func(painter uintptr, options uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(options)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "drawRow"
}

func (this *QTreeView) OnDrawBranches(fn func(*QPainter, *QRect, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 6369455307610057755, func(painter uintptr, rect uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "drawBranches"
}

func (this *QTreeView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QTreeView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QTreeView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTreeView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QTreeView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QTreeView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTreeView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QTreeView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QTreeView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTreeView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QTreeView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QTreeView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QTreeView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QTreeView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QTreeView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QTreeView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QTreeView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTreeView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QTreeView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QTreeView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QTreeView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QTreeView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QTreeView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QTreeView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QTreeView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QTreeView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QTreeView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QTreeView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QTreeView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTreeView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTreeView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTreeView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTreeView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTreeView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTreeView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTreeView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QTreeView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QTreeView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTreeView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTreeView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTreeView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTreeView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QTreeView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QTreeView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTreeView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTreeView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTreeView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTreeView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTreeView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTreeView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTreeView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTreeView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTreeView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTreeView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTreeView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTreeView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTreeView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QTreeView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTreeView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTreeView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTreeView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTreeView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTreeView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTreeView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTreeView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTreeView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTreeView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTreeView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTreeView_delete(this.h)
	}
	this.UnsafeRelease()
}
