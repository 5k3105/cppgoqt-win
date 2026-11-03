package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUndoView_chain uintptr
	addr_QUndoView_new uintptr
	addr_QUndoView_new2 uintptr
	addr_QUndoView_new3 uintptr
	addr_QUndoView_new4 uintptr
	addr_QUndoView_new5 uintptr
	addr_QUndoView_new6 uintptr
	addr_QUndoView_tr uintptr
	addr_QUndoView_stack uintptr
	addr_QUndoView_group uintptr
	addr_QUndoView_setEmptyLabel uintptr
	addr_QUndoView_emptyLabel uintptr
	addr_QUndoView_setCleanIcon uintptr
	addr_QUndoView_cleanIcon uintptr
	addr_QUndoView_setStack uintptr
	addr_QUndoView_setGroup uintptr
	addr_QUndoView_resizeContents uintptr
	addr_QUndoView_contentsSize uintptr
	addr_QUndoView_rectForIndex uintptr
	addr_QUndoView_setPositionForIndex uintptr
	addr_QUndoView_state uintptr
	addr_QUndoView_scheduleDelayedItemsLayout uintptr
	addr_QUndoView_executeDelayedItemsLayout uintptr
	addr_QUndoView_setDirtyRegion uintptr
	addr_QUndoView_scrollDirtyRegion uintptr
	addr_QUndoView_dirtyRegionOffset uintptr
	addr_QUndoView_startAutoScroll uintptr
	addr_QUndoView_stopAutoScroll uintptr
	addr_QUndoView_doAutoScroll uintptr
	addr_QUndoView_dropIndicatorPosition uintptr
	addr_QUndoView_setViewportMargins uintptr
	addr_QUndoView_setViewportMargins2 uintptr
	addr_QUndoView_viewportMargins uintptr
	addr_QUndoView_drawFrame uintptr
	addr_QUndoView_setTabOrder uintptr
	addr_QUndoView_mouseGrabber uintptr
	addr_QUndoView_keyboardGrabber uintptr
	addr_QUndoView_find uintptr
	addr_QUndoView_createWindowContainer uintptr
	addr_QUndoView_updateMicroFocus uintptr
	addr_QUndoView_create uintptr
	addr_QUndoView_destroy uintptr
	addr_QUndoView_focusNextChild uintptr
	addr_QUndoView_focusPreviousChild uintptr
	addr_QUndoView_connect uintptr
	addr_QUndoView_connect2 uintptr
	addr_QUndoView_sender uintptr
	addr_QUndoView_senderSignalIndex uintptr
	addr_QUndoView_receivers uintptr
	addr_QUndoView_isSignalConnected uintptr
	pg_QUndoView_devicePixelRatioFScale func() float64
	pg_QUndoView_encodeMetricF func(int32, float64) int32
	pg_QUndoView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QUndoView_tr2 uintptr
	addr_QUndoView_tr3 uintptr
	addr_QUndoView_createWindowContainer2 uintptr
	addr_QUndoView_createWindowContainer3 uintptr
	addr_QUndoView_updateMicroFocus2 uintptr
	addr_QUndoView_create2 uintptr
	addr_QUndoView_create3 uintptr
	addr_QUndoView_create4 uintptr
	addr_QUndoView_destroy2 uintptr
	addr_QUndoView_destroy3 uintptr
	addr_QUndoView_connect3 uintptr
	addr_QUndoView_connect4 uintptr
	addr_QUndoView_metaObject_default uintptr
	addr_QUndoView_visualRect_default uintptr
	addr_QUndoView_scrollTo_default uintptr
	addr_QUndoView_indexAt_default uintptr
	addr_QUndoView_doItemsLayout_default uintptr
	addr_QUndoView_reset_default uintptr
	addr_QUndoView_setRootIndex_default uintptr
	addr_QUndoView_event_default uintptr
	addr_QUndoView_scrollContentsBy_default uintptr
	addr_QUndoView_rowsInserted_default uintptr
	addr_QUndoView_rowsAboutToBeRemoved_default uintptr
	addr_QUndoView_mouseMoveEvent_default uintptr
	addr_QUndoView_mouseReleaseEvent_default uintptr
	addr_QUndoView_wheelEvent_default uintptr
	addr_QUndoView_timerEvent_default uintptr
	addr_QUndoView_resizeEvent_default uintptr
	addr_QUndoView_dragMoveEvent_default uintptr
	addr_QUndoView_dragLeaveEvent_default uintptr
	addr_QUndoView_dropEvent_default uintptr
	addr_QUndoView_startDrag_default uintptr
	addr_QUndoView_initViewItemOption_default uintptr
	addr_QUndoView_paintEvent_default uintptr
	addr_QUndoView_horizontalOffset_default uintptr
	addr_QUndoView_verticalOffset_default uintptr
	addr_QUndoView_moveCursor_default uintptr
	addr_QUndoView_setSelection_default uintptr
	addr_QUndoView_visualRegionForSelection_default uintptr
	addr_QUndoView_selectedIndexes_default uintptr
	addr_QUndoView_updateGeometries_default uintptr
	addr_QUndoView_isIndexHidden_default uintptr
	addr_QUndoView_selectionChanged_default uintptr
	addr_QUndoView_currentChanged_default uintptr
	addr_QUndoView_viewportSizeHint_default uintptr
	addr_QUndoView_setModel_default uintptr
	addr_QUndoView_setSelectionModel_default uintptr
	addr_QUndoView_keyboardSearch_default uintptr
	addr_QUndoView_sizeHintForRow_default uintptr
	addr_QUndoView_sizeHintForColumn_default uintptr
	addr_QUndoView_itemDelegateForIndex_default uintptr
	addr_QUndoView_inputMethodQuery_default uintptr
	addr_QUndoView_selectAll_default uintptr
	addr_QUndoView_updateEditorData_default uintptr
	addr_QUndoView_updateEditorGeometries_default uintptr
	addr_QUndoView_verticalScrollbarAction_default uintptr
	addr_QUndoView_horizontalScrollbarAction_default uintptr
	addr_QUndoView_verticalScrollbarValueChanged_default uintptr
	addr_QUndoView_horizontalScrollbarValueChanged_default uintptr
	addr_QUndoView_closeEditor_default uintptr
	addr_QUndoView_commitData_default uintptr
	addr_QUndoView_editorDestroyed_default uintptr
	addr_QUndoView_edit_default uintptr
	addr_QUndoView_selectionCommand_default uintptr
	addr_QUndoView_focusNextPrevChild_default uintptr
	addr_QUndoView_viewportEvent_default uintptr
	addr_QUndoView_mousePressEvent_default uintptr
	addr_QUndoView_mouseDoubleClickEvent_default uintptr
	addr_QUndoView_dragEnterEvent_default uintptr
	addr_QUndoView_focusInEvent_default uintptr
	addr_QUndoView_focusOutEvent_default uintptr
	addr_QUndoView_keyPressEvent_default uintptr
	addr_QUndoView_inputMethodEvent_default uintptr
	addr_QUndoView_eventFilter_default uintptr
	addr_QUndoView_minimumSizeHint_default uintptr
	addr_QUndoView_sizeHint_default uintptr
	addr_QUndoView_setupViewport_default uintptr
	addr_QUndoView_contextMenuEvent_default uintptr
	addr_QUndoView_changeEvent_default uintptr
	addr_QUndoView_initStyleOption_default uintptr
	addr_QUndoView_devType_default uintptr
	addr_QUndoView_setVisible_default uintptr
	addr_QUndoView_heightForWidth_default uintptr
	addr_QUndoView_hasHeightForWidth_default uintptr
	addr_QUndoView_paintEngine_default uintptr
	addr_QUndoView_keyReleaseEvent_default uintptr
	addr_QUndoView_enterEvent_default uintptr
	addr_QUndoView_leaveEvent_default uintptr
	addr_QUndoView_moveEvent_default uintptr
	addr_QUndoView_closeEvent_default uintptr
	addr_QUndoView_tabletEvent_default uintptr
	addr_QUndoView_actionEvent_default uintptr
	addr_QUndoView_showEvent_default uintptr
	addr_QUndoView_hideEvent_default uintptr
	addr_QUndoView_metric_default uintptr
	addr_QUndoView_initPainter_default uintptr
	addr_QUndoView_redirected_default uintptr
	addr_QUndoView_sharedPainter_default uintptr
	addr_QUndoView_childEvent_default uintptr
	addr_QUndoView_customEvent_default uintptr
	addr_QUndoView_connectNotify_default uintptr
	addr_QUndoView_disconnectNotify_default uintptr
	pg_QUndoView_connect_indexesMoved func(uintptr, int64)
	pg_QUndoView_connect_pressed func(uintptr, int64)
	pg_QUndoView_connect_clicked func(uintptr, int64)
	pg_QUndoView_connect_doubleClicked func(uintptr, int64)
	pg_QUndoView_connect_activated func(uintptr, int64)
	pg_QUndoView_connect_entered func(uintptr, int64)
	pg_QUndoView_connect_viewportEntered func(uintptr, int64)
	pg_QUndoView_connect_iconSizeChanged func(uintptr, int64)
	pg_QUndoView_connect_windowTitleChanged func(uintptr, int64)
	pg_QUndoView_connect_windowIconChanged func(uintptr, int64)
	pg_QUndoView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QUndoView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QUndoView_connect_destroyed func(uintptr, int64)
	pg_QUndoView_connect_destroyed2 func(uintptr, int64)
	pg_QUndoView_connect_objectNameChanged func(uintptr, int64)
	pg_QUndoView__disconnect func(int64)
	pg_QUndoView__set_override func(uintptr, int32)
	pg_QUndoView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUndoView_chain, "QUndoView_chain")
	symaddr(&addr_QUndoView_new, "QUndoView_new")
	symaddr(&addr_QUndoView_new2, "QUndoView_new2")
	symaddr(&addr_QUndoView_new3, "QUndoView_new3")
	symaddr(&addr_QUndoView_new4, "QUndoView_new4")
	symaddr(&addr_QUndoView_new5, "QUndoView_new5")
	symaddr(&addr_QUndoView_new6, "QUndoView_new6")
	symaddr(&addr_QUndoView_tr, "QUndoView_tr")
	symaddr(&addr_QUndoView_stack, "QUndoView_stack")
	symaddr(&addr_QUndoView_group, "QUndoView_group")
	symaddr(&addr_QUndoView_setEmptyLabel, "QUndoView_setEmptyLabel")
	symaddr(&addr_QUndoView_emptyLabel, "QUndoView_emptyLabel")
	symaddr(&addr_QUndoView_setCleanIcon, "QUndoView_setCleanIcon")
	symaddr(&addr_QUndoView_cleanIcon, "QUndoView_cleanIcon")
	symaddr(&addr_QUndoView_setStack, "QUndoView_setStack")
	symaddr(&addr_QUndoView_setGroup, "QUndoView_setGroup")
	symaddr(&addr_QUndoView_resizeContents, "QUndoView_resizeContents")
	symaddr(&addr_QUndoView_contentsSize, "QUndoView_contentsSize")
	symaddr(&addr_QUndoView_rectForIndex, "QUndoView_rectForIndex")
	symaddr(&addr_QUndoView_setPositionForIndex, "QUndoView_setPositionForIndex")
	symaddr(&addr_QUndoView_state, "QUndoView_state")
	symaddr(&addr_QUndoView_scheduleDelayedItemsLayout, "QUndoView_scheduleDelayedItemsLayout")
	symaddr(&addr_QUndoView_executeDelayedItemsLayout, "QUndoView_executeDelayedItemsLayout")
	symaddr(&addr_QUndoView_setDirtyRegion, "QUndoView_setDirtyRegion")
	symaddr(&addr_QUndoView_scrollDirtyRegion, "QUndoView_scrollDirtyRegion")
	symaddr(&addr_QUndoView_dirtyRegionOffset, "QUndoView_dirtyRegionOffset")
	symaddr(&addr_QUndoView_startAutoScroll, "QUndoView_startAutoScroll")
	symaddr(&addr_QUndoView_stopAutoScroll, "QUndoView_stopAutoScroll")
	symaddr(&addr_QUndoView_doAutoScroll, "QUndoView_doAutoScroll")
	symaddr(&addr_QUndoView_dropIndicatorPosition, "QUndoView_dropIndicatorPosition")
	symaddr(&addr_QUndoView_setViewportMargins, "QUndoView_setViewportMargins")
	symaddr(&addr_QUndoView_setViewportMargins2, "QUndoView_setViewportMargins2")
	symaddr(&addr_QUndoView_viewportMargins, "QUndoView_viewportMargins")
	symaddr(&addr_QUndoView_drawFrame, "QUndoView_drawFrame")
	symaddr(&addr_QUndoView_setTabOrder, "QUndoView_setTabOrder")
	symaddr(&addr_QUndoView_mouseGrabber, "QUndoView_mouseGrabber")
	symaddr(&addr_QUndoView_keyboardGrabber, "QUndoView_keyboardGrabber")
	symaddr(&addr_QUndoView_find, "QUndoView_find")
	symaddr(&addr_QUndoView_createWindowContainer, "QUndoView_createWindowContainer")
	symaddr(&addr_QUndoView_updateMicroFocus, "QUndoView_updateMicroFocus")
	symaddr(&addr_QUndoView_create, "QUndoView_create")
	symaddr(&addr_QUndoView_destroy, "QUndoView_destroy")
	symaddr(&addr_QUndoView_focusNextChild, "QUndoView_focusNextChild")
	symaddr(&addr_QUndoView_focusPreviousChild, "QUndoView_focusPreviousChild")
	symaddr(&addr_QUndoView_connect, "QUndoView_connect")
	symaddr(&addr_QUndoView_connect2, "QUndoView_connect2")
	symaddr(&addr_QUndoView_sender, "QUndoView_sender")
	symaddr(&addr_QUndoView_senderSignalIndex, "QUndoView_senderSignalIndex")
	symaddr(&addr_QUndoView_receivers, "QUndoView_receivers")
	symaddr(&addr_QUndoView_isSignalConnected, "QUndoView_isSignalConnected")
	reg(&pg_QUndoView_devicePixelRatioFScale, "QUndoView_devicePixelRatioFScale")
	reg(&pg_QUndoView_encodeMetricF, "QUndoView_encodeMetricF")
	reg(&pg_QUndoView_getDecodedMetricF, "QUndoView_getDecodedMetricF")
	symaddr(&addr_QUndoView_tr2, "QUndoView_tr2")
	symaddr(&addr_QUndoView_tr3, "QUndoView_tr3")
	symaddr(&addr_QUndoView_createWindowContainer2, "QUndoView_createWindowContainer2")
	symaddr(&addr_QUndoView_createWindowContainer3, "QUndoView_createWindowContainer3")
	symaddr(&addr_QUndoView_updateMicroFocus2, "QUndoView_updateMicroFocus2")
	symaddr(&addr_QUndoView_create2, "QUndoView_create2")
	symaddr(&addr_QUndoView_create3, "QUndoView_create3")
	symaddr(&addr_QUndoView_create4, "QUndoView_create4")
	symaddr(&addr_QUndoView_destroy2, "QUndoView_destroy2")
	symaddr(&addr_QUndoView_destroy3, "QUndoView_destroy3")
	symaddr(&addr_QUndoView_connect3, "QUndoView_connect3")
	symaddr(&addr_QUndoView_connect4, "QUndoView_connect4")
	symaddr(&addr_QUndoView_metaObject_default, "QUndoView_metaObject_default")
	symaddr(&addr_QUndoView_visualRect_default, "QUndoView_visualRect_default")
	symaddr(&addr_QUndoView_scrollTo_default, "QUndoView_scrollTo_default")
	symaddr(&addr_QUndoView_indexAt_default, "QUndoView_indexAt_default")
	symaddr(&addr_QUndoView_doItemsLayout_default, "QUndoView_doItemsLayout_default")
	symaddr(&addr_QUndoView_reset_default, "QUndoView_reset_default")
	symaddr(&addr_QUndoView_setRootIndex_default, "QUndoView_setRootIndex_default")
	symaddr(&addr_QUndoView_event_default, "QUndoView_event_default")
	symaddr(&addr_QUndoView_scrollContentsBy_default, "QUndoView_scrollContentsBy_default")
	symaddr(&addr_QUndoView_rowsInserted_default, "QUndoView_rowsInserted_default")
	symaddr(&addr_QUndoView_rowsAboutToBeRemoved_default, "QUndoView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QUndoView_mouseMoveEvent_default, "QUndoView_mouseMoveEvent_default")
	symaddr(&addr_QUndoView_mouseReleaseEvent_default, "QUndoView_mouseReleaseEvent_default")
	symaddr(&addr_QUndoView_wheelEvent_default, "QUndoView_wheelEvent_default")
	symaddr(&addr_QUndoView_timerEvent_default, "QUndoView_timerEvent_default")
	symaddr(&addr_QUndoView_resizeEvent_default, "QUndoView_resizeEvent_default")
	symaddr(&addr_QUndoView_dragMoveEvent_default, "QUndoView_dragMoveEvent_default")
	symaddr(&addr_QUndoView_dragLeaveEvent_default, "QUndoView_dragLeaveEvent_default")
	symaddr(&addr_QUndoView_dropEvent_default, "QUndoView_dropEvent_default")
	symaddr(&addr_QUndoView_startDrag_default, "QUndoView_startDrag_default")
	symaddr(&addr_QUndoView_initViewItemOption_default, "QUndoView_initViewItemOption_default")
	symaddr(&addr_QUndoView_paintEvent_default, "QUndoView_paintEvent_default")
	symaddr(&addr_QUndoView_horizontalOffset_default, "QUndoView_horizontalOffset_default")
	symaddr(&addr_QUndoView_verticalOffset_default, "QUndoView_verticalOffset_default")
	symaddr(&addr_QUndoView_moveCursor_default, "QUndoView_moveCursor_default")
	symaddr(&addr_QUndoView_setSelection_default, "QUndoView_setSelection_default")
	symaddr(&addr_QUndoView_visualRegionForSelection_default, "QUndoView_visualRegionForSelection_default")
	symaddr(&addr_QUndoView_selectedIndexes_default, "QUndoView_selectedIndexes_default")
	symaddr(&addr_QUndoView_updateGeometries_default, "QUndoView_updateGeometries_default")
	symaddr(&addr_QUndoView_isIndexHidden_default, "QUndoView_isIndexHidden_default")
	symaddr(&addr_QUndoView_selectionChanged_default, "QUndoView_selectionChanged_default")
	symaddr(&addr_QUndoView_currentChanged_default, "QUndoView_currentChanged_default")
	symaddr(&addr_QUndoView_viewportSizeHint_default, "QUndoView_viewportSizeHint_default")
	symaddr(&addr_QUndoView_setModel_default, "QUndoView_setModel_default")
	symaddr(&addr_QUndoView_setSelectionModel_default, "QUndoView_setSelectionModel_default")
	symaddr(&addr_QUndoView_keyboardSearch_default, "QUndoView_keyboardSearch_default")
	symaddr(&addr_QUndoView_sizeHintForRow_default, "QUndoView_sizeHintForRow_default")
	symaddr(&addr_QUndoView_sizeHintForColumn_default, "QUndoView_sizeHintForColumn_default")
	symaddr(&addr_QUndoView_itemDelegateForIndex_default, "QUndoView_itemDelegateForIndex_default")
	symaddr(&addr_QUndoView_inputMethodQuery_default, "QUndoView_inputMethodQuery_default")
	symaddr(&addr_QUndoView_selectAll_default, "QUndoView_selectAll_default")
	symaddr(&addr_QUndoView_updateEditorData_default, "QUndoView_updateEditorData_default")
	symaddr(&addr_QUndoView_updateEditorGeometries_default, "QUndoView_updateEditorGeometries_default")
	symaddr(&addr_QUndoView_verticalScrollbarAction_default, "QUndoView_verticalScrollbarAction_default")
	symaddr(&addr_QUndoView_horizontalScrollbarAction_default, "QUndoView_horizontalScrollbarAction_default")
	symaddr(&addr_QUndoView_verticalScrollbarValueChanged_default, "QUndoView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QUndoView_horizontalScrollbarValueChanged_default, "QUndoView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QUndoView_closeEditor_default, "QUndoView_closeEditor_default")
	symaddr(&addr_QUndoView_commitData_default, "QUndoView_commitData_default")
	symaddr(&addr_QUndoView_editorDestroyed_default, "QUndoView_editorDestroyed_default")
	symaddr(&addr_QUndoView_edit_default, "QUndoView_edit_default")
	symaddr(&addr_QUndoView_selectionCommand_default, "QUndoView_selectionCommand_default")
	symaddr(&addr_QUndoView_focusNextPrevChild_default, "QUndoView_focusNextPrevChild_default")
	symaddr(&addr_QUndoView_viewportEvent_default, "QUndoView_viewportEvent_default")
	symaddr(&addr_QUndoView_mousePressEvent_default, "QUndoView_mousePressEvent_default")
	symaddr(&addr_QUndoView_mouseDoubleClickEvent_default, "QUndoView_mouseDoubleClickEvent_default")
	symaddr(&addr_QUndoView_dragEnterEvent_default, "QUndoView_dragEnterEvent_default")
	symaddr(&addr_QUndoView_focusInEvent_default, "QUndoView_focusInEvent_default")
	symaddr(&addr_QUndoView_focusOutEvent_default, "QUndoView_focusOutEvent_default")
	symaddr(&addr_QUndoView_keyPressEvent_default, "QUndoView_keyPressEvent_default")
	symaddr(&addr_QUndoView_inputMethodEvent_default, "QUndoView_inputMethodEvent_default")
	symaddr(&addr_QUndoView_eventFilter_default, "QUndoView_eventFilter_default")
	symaddr(&addr_QUndoView_minimumSizeHint_default, "QUndoView_minimumSizeHint_default")
	symaddr(&addr_QUndoView_sizeHint_default, "QUndoView_sizeHint_default")
	symaddr(&addr_QUndoView_setupViewport_default, "QUndoView_setupViewport_default")
	symaddr(&addr_QUndoView_contextMenuEvent_default, "QUndoView_contextMenuEvent_default")
	symaddr(&addr_QUndoView_changeEvent_default, "QUndoView_changeEvent_default")
	symaddr(&addr_QUndoView_initStyleOption_default, "QUndoView_initStyleOption_default")
	symaddr(&addr_QUndoView_devType_default, "QUndoView_devType_default")
	symaddr(&addr_QUndoView_setVisible_default, "QUndoView_setVisible_default")
	symaddr(&addr_QUndoView_heightForWidth_default, "QUndoView_heightForWidth_default")
	symaddr(&addr_QUndoView_hasHeightForWidth_default, "QUndoView_hasHeightForWidth_default")
	symaddr(&addr_QUndoView_paintEngine_default, "QUndoView_paintEngine_default")
	symaddr(&addr_QUndoView_keyReleaseEvent_default, "QUndoView_keyReleaseEvent_default")
	symaddr(&addr_QUndoView_enterEvent_default, "QUndoView_enterEvent_default")
	symaddr(&addr_QUndoView_leaveEvent_default, "QUndoView_leaveEvent_default")
	symaddr(&addr_QUndoView_moveEvent_default, "QUndoView_moveEvent_default")
	symaddr(&addr_QUndoView_closeEvent_default, "QUndoView_closeEvent_default")
	symaddr(&addr_QUndoView_tabletEvent_default, "QUndoView_tabletEvent_default")
	symaddr(&addr_QUndoView_actionEvent_default, "QUndoView_actionEvent_default")
	symaddr(&addr_QUndoView_showEvent_default, "QUndoView_showEvent_default")
	symaddr(&addr_QUndoView_hideEvent_default, "QUndoView_hideEvent_default")
	symaddr(&addr_QUndoView_metric_default, "QUndoView_metric_default")
	symaddr(&addr_QUndoView_initPainter_default, "QUndoView_initPainter_default")
	symaddr(&addr_QUndoView_redirected_default, "QUndoView_redirected_default")
	symaddr(&addr_QUndoView_sharedPainter_default, "QUndoView_sharedPainter_default")
	symaddr(&addr_QUndoView_childEvent_default, "QUndoView_childEvent_default")
	symaddr(&addr_QUndoView_customEvent_default, "QUndoView_customEvent_default")
	symaddr(&addr_QUndoView_connectNotify_default, "QUndoView_connectNotify_default")
	symaddr(&addr_QUndoView_disconnectNotify_default, "QUndoView_disconnectNotify_default")
	reg(&pg_QUndoView_connect_indexesMoved, "QUndoView_connect_indexesMoved")
	reg(&pg_QUndoView_connect_pressed, "QUndoView_connect_pressed")
	reg(&pg_QUndoView_connect_clicked, "QUndoView_connect_clicked")
	reg(&pg_QUndoView_connect_doubleClicked, "QUndoView_connect_doubleClicked")
	reg(&pg_QUndoView_connect_activated, "QUndoView_connect_activated")
	reg(&pg_QUndoView_connect_entered, "QUndoView_connect_entered")
	reg(&pg_QUndoView_connect_viewportEntered, "QUndoView_connect_viewportEntered")
	reg(&pg_QUndoView_connect_iconSizeChanged, "QUndoView_connect_iconSizeChanged")
	reg(&pg_QUndoView_connect_windowTitleChanged, "QUndoView_connect_windowTitleChanged")
	reg(&pg_QUndoView_connect_windowIconChanged, "QUndoView_connect_windowIconChanged")
	reg(&pg_QUndoView_connect_windowIconTextChanged, "QUndoView_connect_windowIconTextChanged")
	reg(&pg_QUndoView_connect_customContextMenuRequested, "QUndoView_connect_customContextMenuRequested")
	reg(&pg_QUndoView_connect_destroyed, "QUndoView_connect_destroyed")
	reg(&pg_QUndoView_connect_destroyed2, "QUndoView_connect_destroyed2")
	reg(&pg_QUndoView_connect_objectNameChanged, "QUndoView_connect_objectNameChanged")
	reg(&pg_QUndoView__disconnect, "QUndoView__disconnect")
	reg(&pg_QUndoView__set_override, "QUndoView__set_override")
	reg(&pg_QUndoView_delete, "QUndoView_delete")
}

type QUndoView struct {
	h uintptr
	*QListView
}

func bindQUndoView(h uintptr) *QUndoView {
	if h == 0 {
		return nil
	}
	var __c [6]uintptr
	purego.SyscallN(addr_QUndoView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QUndoView{h: h, QListView: &QListView{h: __c[0], QAbstractItemView: &QAbstractItemView{h: __c[1], QAbstractScrollArea: &QAbstractScrollArea{h: __c[2], QFrame: &QFrame{h: __c[3], QWidget: &QWidget{h: __c[4], QObject: &QObject{h: __c[5]}}}}}}}
}

func newownedQUndoView(h uintptr) *QUndoView { return bindQUndoView(h) }

func (this *QUndoView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUndoView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUndoView(h unsafe.Pointer) *QUndoView { return bindQUndoView(uintptr(h)) }
func UnsafeNewQUndoViewOwned(h unsafe.Pointer) *QUndoView { return newownedQUndoView(uintptr(h)) }

func (this *QUndoView) UnsafeRelease() {
	this.h = 0
	if this.QListView != nil {
		this.QListView.UnsafeRelease()
	}
}

func NewQUndoView(parent *QWidget) *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func NewQUndoView2() *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new2, uintptr(__goid))
	__o := newownedQUndoView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func NewQUndoView3(stack *QUndoStack) *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new3, uintptr(__goid), uintptr(stack.UnsafePointer()))
	__o := newownedQUndoView(__r)
	runtime.KeepAlive(stack)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func NewQUndoView4(group *QUndoGroup) *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new4, uintptr(__goid), uintptr(group.UnsafePointer()))
	__o := newownedQUndoView(__r)
	runtime.KeepAlive(group)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func NewQUndoView5(stack *QUndoStack, parent *QWidget) *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new5, uintptr(__goid), uintptr(stack.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoView(__r)
	runtime.KeepAlive(stack)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func NewQUndoView6(group *QUndoGroup, parent *QWidget) *QUndoView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoView_new6, uintptr(__goid), uintptr(group.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoView(__r)
	runtime.KeepAlive(group)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoView_virtIdx, QUndoView_enable)
	return __o
}

func QUndoView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QUndoView) Stack() *QUndoStack {
	__r, _, _ := purego.SyscallN(addr_QUndoView_stack, this.h)
	runtime.KeepAlive(this)
	return bindQUndoStack(__r)
}

func (this *QUndoView) Group() *QUndoGroup {
	__r, _, _ := purego.SyscallN(addr_QUndoView_group, this.h)
	runtime.KeepAlive(this)
	return bindQUndoGroup(__r)
}

func (this *QUndoView) SetEmptyLabel(label string) {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	purego.SyscallN(addr_QUndoView_setEmptyLabel, this.h, label_p, uintptr(len(label)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
}

func (this *QUndoView) EmptyLabel() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoView_emptyLabel, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoView) SetCleanIcon(icon *QIcon) {
	purego.SyscallN(addr_QUndoView_setCleanIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QUndoView) CleanIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QUndoView_cleanIcon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QUndoView) SetStack(stack *QUndoStack) {
	purego.SyscallN(addr_QUndoView_setStack, this.h, uintptr(stack.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stack)
}

func (this *QUndoView) SetGroup(group *QUndoGroup) {
	purego.SyscallN(addr_QUndoView_setGroup, this.h, uintptr(group.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(group)
}

func (this *QUndoView) ResizeContents(width int, height int) {
	purego.SyscallN(addr_QUndoView_resizeContents, this.h, uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QUndoView) ContentsSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QUndoView_contentsSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QUndoView) RectForIndex(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QUndoView_rectForIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QUndoView) SetPositionForIndex(position *QPoint, index *QModelIndex) {
	purego.SyscallN(addr_QUndoView_setPositionForIndex, this.h, uintptr(position.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	runtime.KeepAlive(index)
}

func (this *QUndoView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QUndoView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QUndoView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QUndoView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QUndoView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QUndoView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QUndoView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QUndoView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QUndoView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QUndoView) StartAutoScroll() {
	purego.SyscallN(addr_QUndoView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) StopAutoScroll() {
	purego.SyscallN(addr_QUndoView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) DoAutoScroll() {
	purego.SyscallN(addr_QUndoView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QUndoView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QUndoView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QUndoView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QUndoView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QUndoView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QUndoView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QUndoView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QUndoView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QUndoView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QUndoView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QUndoView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_mouseGrabber)
	return bindQWidget(__r)
}

func QUndoView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_keyboardGrabber)
	return bindQWidget(__r)
}

func QUndoView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QUndoView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QUndoView) UpdateMicroFocus() {
	purego.SyscallN(addr_QUndoView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) Create() {
	purego.SyscallN(addr_QUndoView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) Destroy() {
	purego.SyscallN(addr_QUndoView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QUndoView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QUndoView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QUndoView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QUndoView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QUndoView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QUndoView_DevicePixelRatioFScale() float64 {
	__r := pg_QUndoView_devicePixelRatioFScale()
	return __r
}

func QUndoView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QUndoView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QUndoView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QUndoView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QUndoView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QUndoView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QUndoView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QUndoView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QUndoView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QUndoView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QUndoView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QUndoView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QUndoView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QUndoView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QUndoView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QUndoView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QUndoView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QUndoView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QUndoView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QUndoView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QUndoView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QUndoView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QUndoView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QUndoView) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QUndoView_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QUndoView) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QUndoView_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QUndoView) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QUndoView_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QUndoView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QUndoView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) ResetDefault() {
	purego.SyscallN(addr_QUndoView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QUndoView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QUndoView) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QUndoView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QUndoView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QUndoView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QUndoView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QUndoView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QUndoView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QUndoView) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QUndoView_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QUndoView_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QUndoView_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QUndoView_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QUndoView_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QUndoView_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QUndoView_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) DropEventDefault(e *QDropEvent) {
	purego.SyscallN(addr_QUndoView_dropEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QUndoView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QUndoView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QUndoView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QUndoView) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QUndoView_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QUndoView) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QUndoView_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QUndoView) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QUndoView_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QUndoView) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QUndoView_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QUndoView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QUndoView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QUndoView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QUndoView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QUndoView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QUndoView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QUndoView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QUndoView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QUndoView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QUndoView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QUndoView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QUndoView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QUndoView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QUndoView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QUndoView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QUndoView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QUndoView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QUndoView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QUndoView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QUndoView) SelectAllDefault() {
	purego.SyscallN(addr_QUndoView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QUndoView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QUndoView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QUndoView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QUndoView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QUndoView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QUndoView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QUndoView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QUndoView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QUndoView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QUndoView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QUndoView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QUndoView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QUndoView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QUndoView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QUndoView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QUndoView) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QUndoView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QUndoView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QUndoView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QUndoView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QUndoView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QUndoView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QUndoView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QUndoView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QUndoView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QUndoView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QUndoView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QUndoView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QUndoView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QUndoView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QUndoView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QUndoView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QUndoView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QUndoView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QUndoView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QUndoView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QUndoView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QUndoView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QUndoView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QUndoView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QUndoView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QUndoView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QUndoView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QUndoView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QUndoView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QUndoView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QUndoView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QUndoView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QUndoView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QUndoView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QUndoView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QUndoView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QUndoView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QUndoView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QUndoView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QUndoView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QUndoView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QUndoView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoView) OnIndexesMoved(fn func([]*QModelIndex)) int64 {
	id := registerCallback(func(indexes uintptr, indexes_len int64) { fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()) })
	pg_QUndoView_connect_indexesMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QUndoView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QUndoView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QUndoView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QUndoView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QUndoView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QUndoView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QUndoView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QUndoView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QUndoView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QUndoView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QUndoView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QUndoView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QUndoView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QUndoView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoView) Disconnect(handle int64) {
	pg_QUndoView__disconnect(handle)
	unregisterCallback(handle)
}

func QUndoView_virtIdx(slot int64) int32 {
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

func QUndoView_enable(h uintptr, idx int32) { pg_QUndoView__set_override(h, idx) }

func (this *QUndoView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QUndoView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QUndoView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QUndoView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QUndoView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QUndoView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QUndoView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QUndoView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QUndoView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QUndoView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QUndoView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QUndoView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QUndoView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QUndoView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QUndoView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QUndoView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QUndoView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QUndoView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QUndoView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(e uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(e))) }) // slot "dropEvent"
}

func (this *QUndoView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QUndoView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QUndoView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QUndoView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QUndoView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QUndoView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QUndoView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QUndoView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QUndoView) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QUndoView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QUndoView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QUndoView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QUndoView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QUndoView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QUndoView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QUndoView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QUndoView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QUndoView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QUndoView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QUndoView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QUndoView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QUndoView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QUndoView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QUndoView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QUndoView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QUndoView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QUndoView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QUndoView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QUndoView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QUndoView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QUndoView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QUndoView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QUndoView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QUndoView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QUndoView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QUndoView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QUndoView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QUndoView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QUndoView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QUndoView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QUndoView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QUndoView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QUndoView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QUndoView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QUndoView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QUndoView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QUndoView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QUndoView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QUndoView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QUndoView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QUndoView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QUndoView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QUndoView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QUndoView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QUndoView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QUndoView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QUndoView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QUndoView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QUndoView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QUndoView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QUndoView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QUndoView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QUndoView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QUndoView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QUndoView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QUndoView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QUndoView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QUndoView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QUndoView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QUndoView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QUndoView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QUndoView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QUndoView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUndoView_delete(this.h)
	}
	this.UnsafeRelease()
}
