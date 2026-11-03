package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QListView_chain uintptr
	addr_QListView_new uintptr
	addr_QListView_new2 uintptr
	addr_QListView_tr uintptr
	addr_QListView_setMovement uintptr
	addr_QListView_movement uintptr
	addr_QListView_setFlow uintptr
	addr_QListView_flow uintptr
	addr_QListView_setWrapping uintptr
	addr_QListView_isWrapping uintptr
	addr_QListView_setResizeMode uintptr
	addr_QListView_resizeMode uintptr
	addr_QListView_setLayoutMode uintptr
	addr_QListView_layoutMode uintptr
	addr_QListView_setSpacing uintptr
	addr_QListView_spacing uintptr
	addr_QListView_setBatchSize uintptr
	addr_QListView_batchSize uintptr
	addr_QListView_setGridSize uintptr
	addr_QListView_gridSize uintptr
	addr_QListView_setViewMode uintptr
	addr_QListView_viewMode uintptr
	addr_QListView_clearPropertyFlags uintptr
	addr_QListView_isRowHidden uintptr
	addr_QListView_setRowHidden uintptr
	addr_QListView_setModelColumn uintptr
	addr_QListView_modelColumn uintptr
	addr_QListView_setUniformItemSizes uintptr
	addr_QListView_uniformItemSizes uintptr
	addr_QListView_setWordWrap uintptr
	addr_QListView_wordWrap uintptr
	addr_QListView_setSelectionRectVisible uintptr
	addr_QListView_isSelectionRectVisible uintptr
	addr_QListView_setItemAlignment uintptr
	addr_QListView_itemAlignment uintptr
	addr_QListView_indexesMoved uintptr
	addr_QListView_resizeContents uintptr
	addr_QListView_contentsSize uintptr
	addr_QListView_rectForIndex uintptr
	addr_QListView_setPositionForIndex uintptr
	addr_QListView_state uintptr
	addr_QListView_scheduleDelayedItemsLayout uintptr
	addr_QListView_executeDelayedItemsLayout uintptr
	addr_QListView_setDirtyRegion uintptr
	addr_QListView_scrollDirtyRegion uintptr
	addr_QListView_dirtyRegionOffset uintptr
	addr_QListView_startAutoScroll uintptr
	addr_QListView_stopAutoScroll uintptr
	addr_QListView_doAutoScroll uintptr
	addr_QListView_dropIndicatorPosition uintptr
	addr_QListView_setViewportMargins uintptr
	addr_QListView_setViewportMargins2 uintptr
	addr_QListView_viewportMargins uintptr
	addr_QListView_drawFrame uintptr
	addr_QListView_setTabOrder uintptr
	addr_QListView_mouseGrabber uintptr
	addr_QListView_keyboardGrabber uintptr
	addr_QListView_find uintptr
	addr_QListView_createWindowContainer uintptr
	addr_QListView_updateMicroFocus uintptr
	addr_QListView_create uintptr
	addr_QListView_destroy uintptr
	addr_QListView_focusNextChild uintptr
	addr_QListView_focusPreviousChild uintptr
	addr_QListView_connect uintptr
	addr_QListView_connect2 uintptr
	addr_QListView_sender uintptr
	addr_QListView_senderSignalIndex uintptr
	addr_QListView_receivers uintptr
	addr_QListView_isSignalConnected uintptr
	pg_QListView_devicePixelRatioFScale func() float64
	pg_QListView_encodeMetricF func(int32, float64) int32
	pg_QListView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QListView_tr2 uintptr
	addr_QListView_tr3 uintptr
	addr_QListView_createWindowContainer2 uintptr
	addr_QListView_createWindowContainer3 uintptr
	addr_QListView_updateMicroFocus2 uintptr
	addr_QListView_create2 uintptr
	addr_QListView_create3 uintptr
	addr_QListView_create4 uintptr
	addr_QListView_destroy2 uintptr
	addr_QListView_destroy3 uintptr
	addr_QListView_connect3 uintptr
	addr_QListView_connect4 uintptr
	addr_QListView_metaObject_default uintptr
	addr_QListView_visualRect_default uintptr
	addr_QListView_scrollTo_default uintptr
	addr_QListView_indexAt_default uintptr
	addr_QListView_doItemsLayout_default uintptr
	addr_QListView_reset_default uintptr
	addr_QListView_setRootIndex_default uintptr
	addr_QListView_event_default uintptr
	addr_QListView_scrollContentsBy_default uintptr
	addr_QListView_rowsInserted_default uintptr
	addr_QListView_rowsAboutToBeRemoved_default uintptr
	addr_QListView_mouseMoveEvent_default uintptr
	addr_QListView_mouseReleaseEvent_default uintptr
	addr_QListView_wheelEvent_default uintptr
	addr_QListView_timerEvent_default uintptr
	addr_QListView_resizeEvent_default uintptr
	addr_QListView_dragMoveEvent_default uintptr
	addr_QListView_dragLeaveEvent_default uintptr
	addr_QListView_dropEvent_default uintptr
	addr_QListView_startDrag_default uintptr
	addr_QListView_initViewItemOption_default uintptr
	addr_QListView_paintEvent_default uintptr
	addr_QListView_horizontalOffset_default uintptr
	addr_QListView_verticalOffset_default uintptr
	addr_QListView_moveCursor_default uintptr
	addr_QListView_setSelection_default uintptr
	addr_QListView_visualRegionForSelection_default uintptr
	addr_QListView_selectedIndexes_default uintptr
	addr_QListView_updateGeometries_default uintptr
	addr_QListView_isIndexHidden_default uintptr
	addr_QListView_selectionChanged_default uintptr
	addr_QListView_currentChanged_default uintptr
	addr_QListView_viewportSizeHint_default uintptr
	addr_QListView_setModel_default uintptr
	addr_QListView_setSelectionModel_default uintptr
	addr_QListView_keyboardSearch_default uintptr
	addr_QListView_sizeHintForRow_default uintptr
	addr_QListView_sizeHintForColumn_default uintptr
	addr_QListView_itemDelegateForIndex_default uintptr
	addr_QListView_inputMethodQuery_default uintptr
	addr_QListView_selectAll_default uintptr
	addr_QListView_updateEditorData_default uintptr
	addr_QListView_updateEditorGeometries_default uintptr
	addr_QListView_verticalScrollbarAction_default uintptr
	addr_QListView_horizontalScrollbarAction_default uintptr
	addr_QListView_verticalScrollbarValueChanged_default uintptr
	addr_QListView_horizontalScrollbarValueChanged_default uintptr
	addr_QListView_closeEditor_default uintptr
	addr_QListView_commitData_default uintptr
	addr_QListView_editorDestroyed_default uintptr
	addr_QListView_edit_default uintptr
	addr_QListView_selectionCommand_default uintptr
	addr_QListView_focusNextPrevChild_default uintptr
	addr_QListView_viewportEvent_default uintptr
	addr_QListView_mousePressEvent_default uintptr
	addr_QListView_mouseDoubleClickEvent_default uintptr
	addr_QListView_dragEnterEvent_default uintptr
	addr_QListView_focusInEvent_default uintptr
	addr_QListView_focusOutEvent_default uintptr
	addr_QListView_keyPressEvent_default uintptr
	addr_QListView_inputMethodEvent_default uintptr
	addr_QListView_eventFilter_default uintptr
	addr_QListView_minimumSizeHint_default uintptr
	addr_QListView_sizeHint_default uintptr
	addr_QListView_setupViewport_default uintptr
	addr_QListView_contextMenuEvent_default uintptr
	addr_QListView_changeEvent_default uintptr
	addr_QListView_initStyleOption_default uintptr
	addr_QListView_devType_default uintptr
	addr_QListView_setVisible_default uintptr
	addr_QListView_heightForWidth_default uintptr
	addr_QListView_hasHeightForWidth_default uintptr
	addr_QListView_paintEngine_default uintptr
	addr_QListView_keyReleaseEvent_default uintptr
	addr_QListView_enterEvent_default uintptr
	addr_QListView_leaveEvent_default uintptr
	addr_QListView_moveEvent_default uintptr
	addr_QListView_closeEvent_default uintptr
	addr_QListView_tabletEvent_default uintptr
	addr_QListView_actionEvent_default uintptr
	addr_QListView_showEvent_default uintptr
	addr_QListView_hideEvent_default uintptr
	addr_QListView_metric_default uintptr
	addr_QListView_initPainter_default uintptr
	addr_QListView_redirected_default uintptr
	addr_QListView_sharedPainter_default uintptr
	addr_QListView_childEvent_default uintptr
	addr_QListView_customEvent_default uintptr
	addr_QListView_connectNotify_default uintptr
	addr_QListView_disconnectNotify_default uintptr
	pg_QListView_connect_indexesMoved func(uintptr, int64)
	pg_QListView_connect_pressed func(uintptr, int64)
	pg_QListView_connect_clicked func(uintptr, int64)
	pg_QListView_connect_doubleClicked func(uintptr, int64)
	pg_QListView_connect_activated func(uintptr, int64)
	pg_QListView_connect_entered func(uintptr, int64)
	pg_QListView_connect_viewportEntered func(uintptr, int64)
	pg_QListView_connect_iconSizeChanged func(uintptr, int64)
	pg_QListView_connect_windowTitleChanged func(uintptr, int64)
	pg_QListView_connect_windowIconChanged func(uintptr, int64)
	pg_QListView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QListView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QListView_connect_destroyed func(uintptr, int64)
	pg_QListView_connect_destroyed2 func(uintptr, int64)
	pg_QListView_connect_objectNameChanged func(uintptr, int64)
	pg_QListView__disconnect func(int64)
	pg_QListView__set_override func(uintptr, int32)
	pg_QListView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QListView_chain, "QListView_chain")
	symaddr(&addr_QListView_new, "QListView_new")
	symaddr(&addr_QListView_new2, "QListView_new2")
	symaddr(&addr_QListView_tr, "QListView_tr")
	symaddr(&addr_QListView_setMovement, "QListView_setMovement")
	symaddr(&addr_QListView_movement, "QListView_movement")
	symaddr(&addr_QListView_setFlow, "QListView_setFlow")
	symaddr(&addr_QListView_flow, "QListView_flow")
	symaddr(&addr_QListView_setWrapping, "QListView_setWrapping")
	symaddr(&addr_QListView_isWrapping, "QListView_isWrapping")
	symaddr(&addr_QListView_setResizeMode, "QListView_setResizeMode")
	symaddr(&addr_QListView_resizeMode, "QListView_resizeMode")
	symaddr(&addr_QListView_setLayoutMode, "QListView_setLayoutMode")
	symaddr(&addr_QListView_layoutMode, "QListView_layoutMode")
	symaddr(&addr_QListView_setSpacing, "QListView_setSpacing")
	symaddr(&addr_QListView_spacing, "QListView_spacing")
	symaddr(&addr_QListView_setBatchSize, "QListView_setBatchSize")
	symaddr(&addr_QListView_batchSize, "QListView_batchSize")
	symaddr(&addr_QListView_setGridSize, "QListView_setGridSize")
	symaddr(&addr_QListView_gridSize, "QListView_gridSize")
	symaddr(&addr_QListView_setViewMode, "QListView_setViewMode")
	symaddr(&addr_QListView_viewMode, "QListView_viewMode")
	symaddr(&addr_QListView_clearPropertyFlags, "QListView_clearPropertyFlags")
	symaddr(&addr_QListView_isRowHidden, "QListView_isRowHidden")
	symaddr(&addr_QListView_setRowHidden, "QListView_setRowHidden")
	symaddr(&addr_QListView_setModelColumn, "QListView_setModelColumn")
	symaddr(&addr_QListView_modelColumn, "QListView_modelColumn")
	symaddr(&addr_QListView_setUniformItemSizes, "QListView_setUniformItemSizes")
	symaddr(&addr_QListView_uniformItemSizes, "QListView_uniformItemSizes")
	symaddr(&addr_QListView_setWordWrap, "QListView_setWordWrap")
	symaddr(&addr_QListView_wordWrap, "QListView_wordWrap")
	symaddr(&addr_QListView_setSelectionRectVisible, "QListView_setSelectionRectVisible")
	symaddr(&addr_QListView_isSelectionRectVisible, "QListView_isSelectionRectVisible")
	symaddr(&addr_QListView_setItemAlignment, "QListView_setItemAlignment")
	symaddr(&addr_QListView_itemAlignment, "QListView_itemAlignment")
	symaddr(&addr_QListView_indexesMoved, "QListView_indexesMoved")
	symaddr(&addr_QListView_resizeContents, "QListView_resizeContents")
	symaddr(&addr_QListView_contentsSize, "QListView_contentsSize")
	symaddr(&addr_QListView_rectForIndex, "QListView_rectForIndex")
	symaddr(&addr_QListView_setPositionForIndex, "QListView_setPositionForIndex")
	symaddr(&addr_QListView_state, "QListView_state")
	symaddr(&addr_QListView_scheduleDelayedItemsLayout, "QListView_scheduleDelayedItemsLayout")
	symaddr(&addr_QListView_executeDelayedItemsLayout, "QListView_executeDelayedItemsLayout")
	symaddr(&addr_QListView_setDirtyRegion, "QListView_setDirtyRegion")
	symaddr(&addr_QListView_scrollDirtyRegion, "QListView_scrollDirtyRegion")
	symaddr(&addr_QListView_dirtyRegionOffset, "QListView_dirtyRegionOffset")
	symaddr(&addr_QListView_startAutoScroll, "QListView_startAutoScroll")
	symaddr(&addr_QListView_stopAutoScroll, "QListView_stopAutoScroll")
	symaddr(&addr_QListView_doAutoScroll, "QListView_doAutoScroll")
	symaddr(&addr_QListView_dropIndicatorPosition, "QListView_dropIndicatorPosition")
	symaddr(&addr_QListView_setViewportMargins, "QListView_setViewportMargins")
	symaddr(&addr_QListView_setViewportMargins2, "QListView_setViewportMargins2")
	symaddr(&addr_QListView_viewportMargins, "QListView_viewportMargins")
	symaddr(&addr_QListView_drawFrame, "QListView_drawFrame")
	symaddr(&addr_QListView_setTabOrder, "QListView_setTabOrder")
	symaddr(&addr_QListView_mouseGrabber, "QListView_mouseGrabber")
	symaddr(&addr_QListView_keyboardGrabber, "QListView_keyboardGrabber")
	symaddr(&addr_QListView_find, "QListView_find")
	symaddr(&addr_QListView_createWindowContainer, "QListView_createWindowContainer")
	symaddr(&addr_QListView_updateMicroFocus, "QListView_updateMicroFocus")
	symaddr(&addr_QListView_create, "QListView_create")
	symaddr(&addr_QListView_destroy, "QListView_destroy")
	symaddr(&addr_QListView_focusNextChild, "QListView_focusNextChild")
	symaddr(&addr_QListView_focusPreviousChild, "QListView_focusPreviousChild")
	symaddr(&addr_QListView_connect, "QListView_connect")
	symaddr(&addr_QListView_connect2, "QListView_connect2")
	symaddr(&addr_QListView_sender, "QListView_sender")
	symaddr(&addr_QListView_senderSignalIndex, "QListView_senderSignalIndex")
	symaddr(&addr_QListView_receivers, "QListView_receivers")
	symaddr(&addr_QListView_isSignalConnected, "QListView_isSignalConnected")
	reg(&pg_QListView_devicePixelRatioFScale, "QListView_devicePixelRatioFScale")
	reg(&pg_QListView_encodeMetricF, "QListView_encodeMetricF")
	reg(&pg_QListView_getDecodedMetricF, "QListView_getDecodedMetricF")
	symaddr(&addr_QListView_tr2, "QListView_tr2")
	symaddr(&addr_QListView_tr3, "QListView_tr3")
	symaddr(&addr_QListView_createWindowContainer2, "QListView_createWindowContainer2")
	symaddr(&addr_QListView_createWindowContainer3, "QListView_createWindowContainer3")
	symaddr(&addr_QListView_updateMicroFocus2, "QListView_updateMicroFocus2")
	symaddr(&addr_QListView_create2, "QListView_create2")
	symaddr(&addr_QListView_create3, "QListView_create3")
	symaddr(&addr_QListView_create4, "QListView_create4")
	symaddr(&addr_QListView_destroy2, "QListView_destroy2")
	symaddr(&addr_QListView_destroy3, "QListView_destroy3")
	symaddr(&addr_QListView_connect3, "QListView_connect3")
	symaddr(&addr_QListView_connect4, "QListView_connect4")
	symaddr(&addr_QListView_metaObject_default, "QListView_metaObject_default")
	symaddr(&addr_QListView_visualRect_default, "QListView_visualRect_default")
	symaddr(&addr_QListView_scrollTo_default, "QListView_scrollTo_default")
	symaddr(&addr_QListView_indexAt_default, "QListView_indexAt_default")
	symaddr(&addr_QListView_doItemsLayout_default, "QListView_doItemsLayout_default")
	symaddr(&addr_QListView_reset_default, "QListView_reset_default")
	symaddr(&addr_QListView_setRootIndex_default, "QListView_setRootIndex_default")
	symaddr(&addr_QListView_event_default, "QListView_event_default")
	symaddr(&addr_QListView_scrollContentsBy_default, "QListView_scrollContentsBy_default")
	symaddr(&addr_QListView_rowsInserted_default, "QListView_rowsInserted_default")
	symaddr(&addr_QListView_rowsAboutToBeRemoved_default, "QListView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QListView_mouseMoveEvent_default, "QListView_mouseMoveEvent_default")
	symaddr(&addr_QListView_mouseReleaseEvent_default, "QListView_mouseReleaseEvent_default")
	symaddr(&addr_QListView_wheelEvent_default, "QListView_wheelEvent_default")
	symaddr(&addr_QListView_timerEvent_default, "QListView_timerEvent_default")
	symaddr(&addr_QListView_resizeEvent_default, "QListView_resizeEvent_default")
	symaddr(&addr_QListView_dragMoveEvent_default, "QListView_dragMoveEvent_default")
	symaddr(&addr_QListView_dragLeaveEvent_default, "QListView_dragLeaveEvent_default")
	symaddr(&addr_QListView_dropEvent_default, "QListView_dropEvent_default")
	symaddr(&addr_QListView_startDrag_default, "QListView_startDrag_default")
	symaddr(&addr_QListView_initViewItemOption_default, "QListView_initViewItemOption_default")
	symaddr(&addr_QListView_paintEvent_default, "QListView_paintEvent_default")
	symaddr(&addr_QListView_horizontalOffset_default, "QListView_horizontalOffset_default")
	symaddr(&addr_QListView_verticalOffset_default, "QListView_verticalOffset_default")
	symaddr(&addr_QListView_moveCursor_default, "QListView_moveCursor_default")
	symaddr(&addr_QListView_setSelection_default, "QListView_setSelection_default")
	symaddr(&addr_QListView_visualRegionForSelection_default, "QListView_visualRegionForSelection_default")
	symaddr(&addr_QListView_selectedIndexes_default, "QListView_selectedIndexes_default")
	symaddr(&addr_QListView_updateGeometries_default, "QListView_updateGeometries_default")
	symaddr(&addr_QListView_isIndexHidden_default, "QListView_isIndexHidden_default")
	symaddr(&addr_QListView_selectionChanged_default, "QListView_selectionChanged_default")
	symaddr(&addr_QListView_currentChanged_default, "QListView_currentChanged_default")
	symaddr(&addr_QListView_viewportSizeHint_default, "QListView_viewportSizeHint_default")
	symaddr(&addr_QListView_setModel_default, "QListView_setModel_default")
	symaddr(&addr_QListView_setSelectionModel_default, "QListView_setSelectionModel_default")
	symaddr(&addr_QListView_keyboardSearch_default, "QListView_keyboardSearch_default")
	symaddr(&addr_QListView_sizeHintForRow_default, "QListView_sizeHintForRow_default")
	symaddr(&addr_QListView_sizeHintForColumn_default, "QListView_sizeHintForColumn_default")
	symaddr(&addr_QListView_itemDelegateForIndex_default, "QListView_itemDelegateForIndex_default")
	symaddr(&addr_QListView_inputMethodQuery_default, "QListView_inputMethodQuery_default")
	symaddr(&addr_QListView_selectAll_default, "QListView_selectAll_default")
	symaddr(&addr_QListView_updateEditorData_default, "QListView_updateEditorData_default")
	symaddr(&addr_QListView_updateEditorGeometries_default, "QListView_updateEditorGeometries_default")
	symaddr(&addr_QListView_verticalScrollbarAction_default, "QListView_verticalScrollbarAction_default")
	symaddr(&addr_QListView_horizontalScrollbarAction_default, "QListView_horizontalScrollbarAction_default")
	symaddr(&addr_QListView_verticalScrollbarValueChanged_default, "QListView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QListView_horizontalScrollbarValueChanged_default, "QListView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QListView_closeEditor_default, "QListView_closeEditor_default")
	symaddr(&addr_QListView_commitData_default, "QListView_commitData_default")
	symaddr(&addr_QListView_editorDestroyed_default, "QListView_editorDestroyed_default")
	symaddr(&addr_QListView_edit_default, "QListView_edit_default")
	symaddr(&addr_QListView_selectionCommand_default, "QListView_selectionCommand_default")
	symaddr(&addr_QListView_focusNextPrevChild_default, "QListView_focusNextPrevChild_default")
	symaddr(&addr_QListView_viewportEvent_default, "QListView_viewportEvent_default")
	symaddr(&addr_QListView_mousePressEvent_default, "QListView_mousePressEvent_default")
	symaddr(&addr_QListView_mouseDoubleClickEvent_default, "QListView_mouseDoubleClickEvent_default")
	symaddr(&addr_QListView_dragEnterEvent_default, "QListView_dragEnterEvent_default")
	symaddr(&addr_QListView_focusInEvent_default, "QListView_focusInEvent_default")
	symaddr(&addr_QListView_focusOutEvent_default, "QListView_focusOutEvent_default")
	symaddr(&addr_QListView_keyPressEvent_default, "QListView_keyPressEvent_default")
	symaddr(&addr_QListView_inputMethodEvent_default, "QListView_inputMethodEvent_default")
	symaddr(&addr_QListView_eventFilter_default, "QListView_eventFilter_default")
	symaddr(&addr_QListView_minimumSizeHint_default, "QListView_minimumSizeHint_default")
	symaddr(&addr_QListView_sizeHint_default, "QListView_sizeHint_default")
	symaddr(&addr_QListView_setupViewport_default, "QListView_setupViewport_default")
	symaddr(&addr_QListView_contextMenuEvent_default, "QListView_contextMenuEvent_default")
	symaddr(&addr_QListView_changeEvent_default, "QListView_changeEvent_default")
	symaddr(&addr_QListView_initStyleOption_default, "QListView_initStyleOption_default")
	symaddr(&addr_QListView_devType_default, "QListView_devType_default")
	symaddr(&addr_QListView_setVisible_default, "QListView_setVisible_default")
	symaddr(&addr_QListView_heightForWidth_default, "QListView_heightForWidth_default")
	symaddr(&addr_QListView_hasHeightForWidth_default, "QListView_hasHeightForWidth_default")
	symaddr(&addr_QListView_paintEngine_default, "QListView_paintEngine_default")
	symaddr(&addr_QListView_keyReleaseEvent_default, "QListView_keyReleaseEvent_default")
	symaddr(&addr_QListView_enterEvent_default, "QListView_enterEvent_default")
	symaddr(&addr_QListView_leaveEvent_default, "QListView_leaveEvent_default")
	symaddr(&addr_QListView_moveEvent_default, "QListView_moveEvent_default")
	symaddr(&addr_QListView_closeEvent_default, "QListView_closeEvent_default")
	symaddr(&addr_QListView_tabletEvent_default, "QListView_tabletEvent_default")
	symaddr(&addr_QListView_actionEvent_default, "QListView_actionEvent_default")
	symaddr(&addr_QListView_showEvent_default, "QListView_showEvent_default")
	symaddr(&addr_QListView_hideEvent_default, "QListView_hideEvent_default")
	symaddr(&addr_QListView_metric_default, "QListView_metric_default")
	symaddr(&addr_QListView_initPainter_default, "QListView_initPainter_default")
	symaddr(&addr_QListView_redirected_default, "QListView_redirected_default")
	symaddr(&addr_QListView_sharedPainter_default, "QListView_sharedPainter_default")
	symaddr(&addr_QListView_childEvent_default, "QListView_childEvent_default")
	symaddr(&addr_QListView_customEvent_default, "QListView_customEvent_default")
	symaddr(&addr_QListView_connectNotify_default, "QListView_connectNotify_default")
	symaddr(&addr_QListView_disconnectNotify_default, "QListView_disconnectNotify_default")
	reg(&pg_QListView_connect_indexesMoved, "QListView_connect_indexesMoved")
	reg(&pg_QListView_connect_pressed, "QListView_connect_pressed")
	reg(&pg_QListView_connect_clicked, "QListView_connect_clicked")
	reg(&pg_QListView_connect_doubleClicked, "QListView_connect_doubleClicked")
	reg(&pg_QListView_connect_activated, "QListView_connect_activated")
	reg(&pg_QListView_connect_entered, "QListView_connect_entered")
	reg(&pg_QListView_connect_viewportEntered, "QListView_connect_viewportEntered")
	reg(&pg_QListView_connect_iconSizeChanged, "QListView_connect_iconSizeChanged")
	reg(&pg_QListView_connect_windowTitleChanged, "QListView_connect_windowTitleChanged")
	reg(&pg_QListView_connect_windowIconChanged, "QListView_connect_windowIconChanged")
	reg(&pg_QListView_connect_windowIconTextChanged, "QListView_connect_windowIconTextChanged")
	reg(&pg_QListView_connect_customContextMenuRequested, "QListView_connect_customContextMenuRequested")
	reg(&pg_QListView_connect_destroyed, "QListView_connect_destroyed")
	reg(&pg_QListView_connect_destroyed2, "QListView_connect_destroyed2")
	reg(&pg_QListView_connect_objectNameChanged, "QListView_connect_objectNameChanged")
	reg(&pg_QListView__disconnect, "QListView__disconnect")
	reg(&pg_QListView__set_override, "QListView__set_override")
	reg(&pg_QListView_delete, "QListView_delete")
}

type QListView struct {
	h uintptr
	*QAbstractItemView
}

func bindQListView(h uintptr) *QListView {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QListView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QListView{h: h, QAbstractItemView: &QAbstractItemView{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQListView(h uintptr) *QListView { return bindQListView(h) }

func (this *QListView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QListView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQListView(h unsafe.Pointer) *QListView { return bindQListView(uintptr(h)) }
func UnsafeNewQListViewOwned(h unsafe.Pointer) *QListView { return newownedQListView(uintptr(h)) }

func (this *QListView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemView != nil {
		this.QAbstractItemView.UnsafeRelease()
	}
}

func NewQListView(parent *QWidget) *QListView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQListView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListView_virtIdx, QListView_enable)
	return __o
}

func NewQListView2() *QListView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListView_new2, uintptr(__goid))
	__o := newownedQListView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListView_virtIdx, QListView_enable)
	return __o
}

func QListView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QListView) SetMovement(movement QListView__Movement) {
	purego.SyscallN(addr_QListView_setMovement, this.h, uintptr(movement))
	runtime.KeepAlive(this)
}

func (this *QListView) Movement() QListView__Movement {
	__r0, _, _ := purego.SyscallN(addr_QListView_movement, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QListView__Movement(__r)
}

func (this *QListView) SetFlow(flow QListView__Flow) {
	purego.SyscallN(addr_QListView_setFlow, this.h, uintptr(flow))
	runtime.KeepAlive(this)
}

func (this *QListView) Flow() QListView__Flow {
	__r0, _, _ := purego.SyscallN(addr_QListView_flow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QListView__Flow(__r)
}

func (this *QListView) SetWrapping(enable bool) {
	purego.SyscallN(addr_QListView_setWrapping, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QListView) IsWrapping() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_isWrapping, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) SetResizeMode(mode QListView__ResizeMode) {
	purego.SyscallN(addr_QListView_setResizeMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QListView) ResizeMode() QListView__ResizeMode {
	__r0, _, _ := purego.SyscallN(addr_QListView_resizeMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QListView__ResizeMode(__r)
}

func (this *QListView) SetLayoutMode(mode QListView__LayoutMode) {
	purego.SyscallN(addr_QListView_setLayoutMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QListView) LayoutMode() QListView__LayoutMode {
	__r0, _, _ := purego.SyscallN(addr_QListView_layoutMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QListView__LayoutMode(__r)
}

func (this *QListView) SetSpacing(space int) {
	purego.SyscallN(addr_QListView_setSpacing, this.h, uintptr(space))
	runtime.KeepAlive(this)
}

func (this *QListView) Spacing() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_spacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) SetBatchSize(batchSize int) {
	purego.SyscallN(addr_QListView_setBatchSize, this.h, uintptr(batchSize))
	runtime.KeepAlive(this)
}

func (this *QListView) BatchSize() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_batchSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) SetGridSize(size *QSize) {
	purego.SyscallN(addr_QListView_setGridSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QListView) GridSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListView_gridSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListView) SetViewMode(mode QListView__ViewMode) {
	purego.SyscallN(addr_QListView_setViewMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QListView) ViewMode() QListView__ViewMode {
	__r0, _, _ := purego.SyscallN(addr_QListView_viewMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QListView__ViewMode(__r)
}

func (this *QListView) ClearPropertyFlags() {
	purego.SyscallN(addr_QListView_clearPropertyFlags, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) IsRowHidden(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_isRowHidden, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) SetRowHidden(row int, hide bool) {
	purego.SyscallN(addr_QListView_setRowHidden, this.h, uintptr(row), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QListView) SetModelColumn(column int) {
	purego.SyscallN(addr_QListView_setModelColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QListView) ModelColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_modelColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) SetUniformItemSizes(enable bool) {
	purego.SyscallN(addr_QListView_setUniformItemSizes, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QListView) UniformItemSizes() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_uniformItemSizes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) SetWordWrap(on bool) {
	purego.SyscallN(addr_QListView_setWordWrap, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QListView) WordWrap() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_wordWrap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) SetSelectionRectVisible(show bool) {
	purego.SyscallN(addr_QListView_setSelectionRectVisible, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QListView) IsSelectionRectVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_isSelectionRectVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) SetItemAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QListView_setItemAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QListView) ItemAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QListView_itemAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QListView) IndexesMoved(indexes []*QModelIndex) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QListView_indexesMoved, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
}

func (this *QListView) ResizeContents(width int, height int) {
	purego.SyscallN(addr_QListView_resizeContents, this.h, uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QListView) ContentsSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListView_contentsSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListView) RectForIndex(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QListView_rectForIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QListView) SetPositionForIndex(position *QPoint, index *QModelIndex) {
	purego.SyscallN(addr_QListView_setPositionForIndex, this.h, uintptr(position.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	runtime.KeepAlive(index)
}

func (this *QListView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QListView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QListView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QListView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QListView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QListView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QListView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QListView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QListView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QListView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QListView) StartAutoScroll() {
	purego.SyscallN(addr_QListView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) StopAutoScroll() {
	purego.SyscallN(addr_QListView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) DoAutoScroll() {
	purego.SyscallN(addr_QListView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QListView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QListView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QListView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QListView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QListView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QListView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QListView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QListView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QListView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QListView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QListView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QListView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_mouseGrabber)
	return bindQWidget(__r)
}

func QListView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_keyboardGrabber)
	return bindQWidget(__r)
}

func QListView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QListView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QListView) UpdateMicroFocus() {
	purego.SyscallN(addr_QListView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) Create() {
	purego.SyscallN(addr_QListView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) Destroy() {
	purego.SyscallN(addr_QListView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QListView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QListView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QListView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QListView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QListView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QListView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QListView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QListView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QListView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QListView_DevicePixelRatioFScale() float64 {
	__r := pg_QListView_devicePixelRatioFScale()
	return __r
}

func QListView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QListView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QListView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QListView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QListView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QListView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QListView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QListView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QListView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QListView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QListView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QListView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QListView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QListView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QListView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QListView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QListView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QListView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QListView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QListView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QListView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QListView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QListView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QListView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QListView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QListView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QListView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QListView) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QListView_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QListView) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QListView_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QListView) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QListView_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QListView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QListView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) ResetDefault() {
	purego.SyscallN(addr_QListView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QListView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QListView) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QListView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QListView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QListView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QListView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QListView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QListView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QListView) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QListView_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QListView_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QListView_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QListView_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QListView_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QListView_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QListView_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) DropEventDefault(e *QDropEvent) {
	purego.SyscallN(addr_QListView_dropEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QListView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QListView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QListView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QListView) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QListView_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QListView) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QListView_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QListView) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QListView_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QListView) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QListView_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QListView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QListView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QListView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QListView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QListView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QListView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QListView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QListView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QListView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QListView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QListView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QListView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QListView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QListView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QListView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QListView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QListView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QListView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QListView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QListView) SelectAllDefault() {
	purego.SyscallN(addr_QListView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QListView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QListView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QListView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QListView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QListView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QListView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QListView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QListView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QListView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QListView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QListView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QListView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QListView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QListView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QListView) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QListView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QListView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QListView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QListView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QListView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QListView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QListView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QListView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QListView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QListView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QListView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QListView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QListView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QListView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QListView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QListView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QListView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QListView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QListView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QListView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QListView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QListView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QListView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QListView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QListView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QListView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QListView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QListView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QListView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QListView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QListView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QListView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QListView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QListView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QListView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QListView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QListView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QListView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QListView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QListView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QListView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QListView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QListView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QListView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QListView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QListView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QListView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QListView) OnIndexesMoved(fn func([]*QModelIndex)) int64 {
	id := registerCallback(func(indexes uintptr, indexes_len int64) { fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()) })
	pg_QListView_connect_indexesMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QListView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QListView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QListView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QListView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QListView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QListView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QListView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QListView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QListView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QListView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QListView) Disconnect(handle int64) {
	pg_QListView__disconnect(handle)
	unregisterCallback(handle)
}

func QListView_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -1267696851865372997:
		return 1
	case -4832340528568269947:
		return 2
	case -7252823847285716802:
		return 3
	case -1888629738259754754:
		return 4
	case 1086335023529244512:
		return 5
	case -1574077111192992763:
		return 6
	case 1237236456312824991:
		return 7
	case -6118001150126564745:
		return 8
	case 4472434241583495620:
		return 9
	case -8151049665126058157:
		return 10
	case -4255494004099806581:
		return 11
	case -4415920049297795213:
		return 12
	case 1919988340934511316:
		return 13
	case -5409927581812374498:
		return 14
	case -4302917285655014661:
		return 15
	case -1781545137352745978:
		return 16
	case -3139904158373725006:
		return 17
	case 9107058682754591716:
		return 18
	case -8112366604639385491:
		return 19
	case -7450787277202094556:
		return 20
	case 5290442163167458465:
		return 21
	case -1500310024801917802:
		return 22
	case 5218707810677921016:
		return 23
	case -158523882194220890:
		return 24
	case -1539266838375037305:
		return 25
	case 7063795720810552074:
		return 26
	case 740869982805721908:
		return 27
	case 1022771470898460576:
		return 28
	case 6057989667089378859:
		return 29
	case -5892205462145366901:
		return 30
	case -6978583972112316034:
		return 31
	case 503579067252048239:
		return 32
	case 4283922507359547360:
		return 33
	case -4689322347236366780:
		return 34
	case -5041452885022114200:
		return 35
	case -4084029023708556876:
		return 36
	case -3666792239577137638:
		return 37
	case -3811438228839363824:
		return 38
	case 7345853097340709638:
		return 39
	case 8758113809152574300:
		return 40
	case 9029941737538200621:
		return 41
	case -8726300139121449691:
		return 42
	case -1155555122391250623:
		return 43
	case -785551405282638793:
		return 44
	case 6366953527806415524:
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
	case 5898328517505784191:
		return 52
	case -1971465653061861519:
		return 53
	case -332090898078543319:
		return 54
	case -4029073905697023887:
		return 55
	case 206122715919917589:
		return 56
	case 975043767033552552:
		return 57
	case -385596010576048993:
		return 58
	case 1023791977679793185:
		return 59
	case 5438819661942843222:
		return 60
	case 5217623514630858297:
		return 61
	case 7049264599837532491:
		return 62
	case 729690261457684101:
		return 63
	case 4531996381736730688:
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

func QListView_enable(h uintptr, idx int32) { pg_QListView__set_override(h, idx) }

func (this *QListView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QListView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QListView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QListView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QListView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QListView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QListView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QListView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QListView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QListView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QListView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QListView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QListView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QListView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QListView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QListView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QListView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QListView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QListView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(e uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(e))) }) // slot "dropEvent"
}

func (this *QListView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QListView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QListView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QListView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QListView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QListView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QListView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QListView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QListView) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QListView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QListView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QListView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QListView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QListView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QListView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QListView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QListView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QListView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QListView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QListView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QListView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QListView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QListView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QListView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QListView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QListView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QListView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QListView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QListView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QListView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QListView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QListView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QListView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QListView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QListView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QListView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QListView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QListView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QListView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QListView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QListView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QListView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QListView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QListView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QListView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QListView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QListView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QListView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QListView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QListView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QListView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QListView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QListView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QListView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QListView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QListView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QListView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QListView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QListView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QListView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QListView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QListView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QListView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QListView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QListView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QListView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QListView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QListView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QListView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QListView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QListView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QListView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QListView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QListView_delete(this.h)
	}
	this.UnsafeRelease()
}
