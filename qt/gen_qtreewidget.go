package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTreeWidget_chain uintptr
	addr_QTreeWidget_new uintptr
	addr_QTreeWidget_new2 uintptr
	addr_QTreeWidget_tr uintptr
	addr_QTreeWidget_columnCount uintptr
	addr_QTreeWidget_setColumnCount uintptr
	addr_QTreeWidget_invisibleRootItem uintptr
	addr_QTreeWidget_topLevelItem uintptr
	addr_QTreeWidget_topLevelItemCount uintptr
	addr_QTreeWidget_insertTopLevelItem uintptr
	addr_QTreeWidget_addTopLevelItem uintptr
	addr_QTreeWidget_takeTopLevelItem uintptr
	addr_QTreeWidget_indexOfTopLevelItem uintptr
	addr_QTreeWidget_insertTopLevelItems uintptr
	addr_QTreeWidget_addTopLevelItems uintptr
	addr_QTreeWidget_headerItem uintptr
	addr_QTreeWidget_setHeaderItem uintptr
	addr_QTreeWidget_setHeaderLabels uintptr
	addr_QTreeWidget_setHeaderLabel uintptr
	addr_QTreeWidget_currentItem uintptr
	addr_QTreeWidget_currentColumn uintptr
	addr_QTreeWidget_setCurrentItem uintptr
	addr_QTreeWidget_setCurrentItem2 uintptr
	addr_QTreeWidget_setCurrentItem3 uintptr
	addr_QTreeWidget_itemAt uintptr
	addr_QTreeWidget_itemAt2 uintptr
	addr_QTreeWidget_visualItemRect uintptr
	addr_QTreeWidget_sortColumn uintptr
	addr_QTreeWidget_sortItems uintptr
	addr_QTreeWidget_editItem uintptr
	addr_QTreeWidget_openPersistentEditor uintptr
	addr_QTreeWidget_closePersistentEditor uintptr
	addr_QTreeWidget_isPersistentEditorOpen uintptr
	addr_QTreeWidget_itemWidget uintptr
	addr_QTreeWidget_setItemWidget uintptr
	addr_QTreeWidget_removeItemWidget uintptr
	addr_QTreeWidget_selectedItems uintptr
	addr_QTreeWidget_findItems uintptr
	addr_QTreeWidget_itemAbove uintptr
	addr_QTreeWidget_itemBelow uintptr
	addr_QTreeWidget_indexFromItem uintptr
	addr_QTreeWidget_itemFromIndex uintptr
	addr_QTreeWidget_supportedDragActions uintptr
	addr_QTreeWidget_setSupportedDragActions uintptr
	addr_QTreeWidget_scrollToItem uintptr
	addr_QTreeWidget_expandItem uintptr
	addr_QTreeWidget_collapseItem uintptr
	addr_QTreeWidget_clear uintptr
	addr_QTreeWidget_itemPressed uintptr
	addr_QTreeWidget_itemClicked uintptr
	addr_QTreeWidget_itemDoubleClicked uintptr
	addr_QTreeWidget_itemActivated uintptr
	addr_QTreeWidget_itemEntered uintptr
	addr_QTreeWidget_itemChanged uintptr
	addr_QTreeWidget_itemExpanded uintptr
	addr_QTreeWidget_itemCollapsed uintptr
	addr_QTreeWidget_currentItemChanged uintptr
	addr_QTreeWidget_itemSelectionChanged uintptr
	addr_QTreeWidget_columnResized uintptr
	addr_QTreeWidget_columnCountChanged uintptr
	addr_QTreeWidget_columnMoved uintptr
	addr_QTreeWidget_reexpand uintptr
	addr_QTreeWidget_rowsRemoved uintptr
	addr_QTreeWidget_drawTree uintptr
	addr_QTreeWidget_indexRowSizeHint uintptr
	addr_QTreeWidget_rowHeight uintptr
	addr_QTreeWidget_state uintptr
	addr_QTreeWidget_scheduleDelayedItemsLayout uintptr
	addr_QTreeWidget_executeDelayedItemsLayout uintptr
	addr_QTreeWidget_setDirtyRegion uintptr
	addr_QTreeWidget_scrollDirtyRegion uintptr
	addr_QTreeWidget_dirtyRegionOffset uintptr
	addr_QTreeWidget_startAutoScroll uintptr
	addr_QTreeWidget_stopAutoScroll uintptr
	addr_QTreeWidget_doAutoScroll uintptr
	addr_QTreeWidget_dropIndicatorPosition uintptr
	addr_QTreeWidget_setViewportMargins uintptr
	addr_QTreeWidget_setViewportMargins2 uintptr
	addr_QTreeWidget_viewportMargins uintptr
	addr_QTreeWidget_drawFrame uintptr
	addr_QTreeWidget_setTabOrder uintptr
	addr_QTreeWidget_mouseGrabber uintptr
	addr_QTreeWidget_keyboardGrabber uintptr
	addr_QTreeWidget_find uintptr
	addr_QTreeWidget_createWindowContainer uintptr
	addr_QTreeWidget_updateMicroFocus uintptr
	addr_QTreeWidget_create uintptr
	addr_QTreeWidget_destroy uintptr
	addr_QTreeWidget_focusNextChild uintptr
	addr_QTreeWidget_focusPreviousChild uintptr
	addr_QTreeWidget_connect uintptr
	addr_QTreeWidget_connect2 uintptr
	addr_QTreeWidget_sender uintptr
	addr_QTreeWidget_senderSignalIndex uintptr
	addr_QTreeWidget_receivers uintptr
	addr_QTreeWidget_isSignalConnected uintptr
	pg_QTreeWidget_devicePixelRatioFScale func() float64
	pg_QTreeWidget_encodeMetricF func(int32, float64) int32
	pg_QTreeWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTreeWidget_tr2 uintptr
	addr_QTreeWidget_tr3 uintptr
	addr_QTreeWidget_editItem2 uintptr
	addr_QTreeWidget_openPersistentEditor2 uintptr
	addr_QTreeWidget_closePersistentEditor2 uintptr
	addr_QTreeWidget_isPersistentEditorOpen2 uintptr
	addr_QTreeWidget_findItems2 uintptr
	addr_QTreeWidget_indexFromItem2 uintptr
	addr_QTreeWidget_scrollToItem2 uintptr
	addr_QTreeWidget_createWindowContainer2 uintptr
	addr_QTreeWidget_createWindowContainer3 uintptr
	addr_QTreeWidget_updateMicroFocus2 uintptr
	addr_QTreeWidget_create2 uintptr
	addr_QTreeWidget_create3 uintptr
	addr_QTreeWidget_create4 uintptr
	addr_QTreeWidget_destroy2 uintptr
	addr_QTreeWidget_destroy3 uintptr
	addr_QTreeWidget_connect3 uintptr
	addr_QTreeWidget_connect4 uintptr
	addr_QTreeWidget_metaObject_default uintptr
	addr_QTreeWidget_setSelectionModel_default uintptr
	addr_QTreeWidget_event_default uintptr
	addr_QTreeWidget_mimeTypes_default uintptr
	addr_QTreeWidget_dropMimeData_default uintptr
	addr_QTreeWidget_supportedDropActions_default uintptr
	addr_QTreeWidget_dropEvent_default uintptr
	addr_QTreeWidget_setRootIndex_default uintptr
	addr_QTreeWidget_keyboardSearch_default uintptr
	addr_QTreeWidget_visualRect_default uintptr
	addr_QTreeWidget_scrollTo_default uintptr
	addr_QTreeWidget_indexAt_default uintptr
	addr_QTreeWidget_doItemsLayout_default uintptr
	addr_QTreeWidget_reset_default uintptr
	addr_QTreeWidget_selectAll_default uintptr
	addr_QTreeWidget_verticalScrollbarValueChanged_default uintptr
	addr_QTreeWidget_scrollContentsBy_default uintptr
	addr_QTreeWidget_rowsInserted_default uintptr
	addr_QTreeWidget_rowsAboutToBeRemoved_default uintptr
	addr_QTreeWidget_moveCursor_default uintptr
	addr_QTreeWidget_horizontalOffset_default uintptr
	addr_QTreeWidget_verticalOffset_default uintptr
	addr_QTreeWidget_setSelection_default uintptr
	addr_QTreeWidget_visualRegionForSelection_default uintptr
	addr_QTreeWidget_selectedIndexes_default uintptr
	addr_QTreeWidget_changeEvent_default uintptr
	addr_QTreeWidget_timerEvent_default uintptr
	addr_QTreeWidget_paintEvent_default uintptr
	addr_QTreeWidget_drawRow_default uintptr
	addr_QTreeWidget_drawBranches_default uintptr
	addr_QTreeWidget_mousePressEvent_default uintptr
	addr_QTreeWidget_mouseReleaseEvent_default uintptr
	addr_QTreeWidget_mouseDoubleClickEvent_default uintptr
	addr_QTreeWidget_mouseMoveEvent_default uintptr
	addr_QTreeWidget_keyPressEvent_default uintptr
	addr_QTreeWidget_dragMoveEvent_default uintptr
	addr_QTreeWidget_viewportEvent_default uintptr
	addr_QTreeWidget_updateGeometries_default uintptr
	addr_QTreeWidget_viewportSizeHint_default uintptr
	addr_QTreeWidget_sizeHintForColumn_default uintptr
	addr_QTreeWidget_horizontalScrollbarAction_default uintptr
	addr_QTreeWidget_isIndexHidden_default uintptr
	addr_QTreeWidget_selectionChanged_default uintptr
	addr_QTreeWidget_currentChanged_default uintptr
	addr_QTreeWidget_sizeHintForRow_default uintptr
	addr_QTreeWidget_itemDelegateForIndex_default uintptr
	addr_QTreeWidget_inputMethodQuery_default uintptr
	addr_QTreeWidget_updateEditorData_default uintptr
	addr_QTreeWidget_updateEditorGeometries_default uintptr
	addr_QTreeWidget_verticalScrollbarAction_default uintptr
	addr_QTreeWidget_horizontalScrollbarValueChanged_default uintptr
	addr_QTreeWidget_closeEditor_default uintptr
	addr_QTreeWidget_commitData_default uintptr
	addr_QTreeWidget_editorDestroyed_default uintptr
	addr_QTreeWidget_edit_default uintptr
	addr_QTreeWidget_selectionCommand_default uintptr
	addr_QTreeWidget_startDrag_default uintptr
	addr_QTreeWidget_initViewItemOption_default uintptr
	addr_QTreeWidget_focusNextPrevChild_default uintptr
	addr_QTreeWidget_dragEnterEvent_default uintptr
	addr_QTreeWidget_dragLeaveEvent_default uintptr
	addr_QTreeWidget_focusInEvent_default uintptr
	addr_QTreeWidget_focusOutEvent_default uintptr
	addr_QTreeWidget_resizeEvent_default uintptr
	addr_QTreeWidget_inputMethodEvent_default uintptr
	addr_QTreeWidget_eventFilter_default uintptr
	addr_QTreeWidget_minimumSizeHint_default uintptr
	addr_QTreeWidget_sizeHint_default uintptr
	addr_QTreeWidget_setupViewport_default uintptr
	addr_QTreeWidget_wheelEvent_default uintptr
	addr_QTreeWidget_contextMenuEvent_default uintptr
	addr_QTreeWidget_initStyleOption_default uintptr
	addr_QTreeWidget_devType_default uintptr
	addr_QTreeWidget_setVisible_default uintptr
	addr_QTreeWidget_heightForWidth_default uintptr
	addr_QTreeWidget_hasHeightForWidth_default uintptr
	addr_QTreeWidget_paintEngine_default uintptr
	addr_QTreeWidget_keyReleaseEvent_default uintptr
	addr_QTreeWidget_enterEvent_default uintptr
	addr_QTreeWidget_leaveEvent_default uintptr
	addr_QTreeWidget_moveEvent_default uintptr
	addr_QTreeWidget_closeEvent_default uintptr
	addr_QTreeWidget_tabletEvent_default uintptr
	addr_QTreeWidget_actionEvent_default uintptr
	addr_QTreeWidget_showEvent_default uintptr
	addr_QTreeWidget_hideEvent_default uintptr
	addr_QTreeWidget_metric_default uintptr
	addr_QTreeWidget_initPainter_default uintptr
	addr_QTreeWidget_redirected_default uintptr
	addr_QTreeWidget_sharedPainter_default uintptr
	addr_QTreeWidget_childEvent_default uintptr
	addr_QTreeWidget_customEvent_default uintptr
	addr_QTreeWidget_connectNotify_default uintptr
	addr_QTreeWidget_disconnectNotify_default uintptr
	pg_QTreeWidget_connect_itemPressed func(uintptr, int64)
	pg_QTreeWidget_connect_itemClicked func(uintptr, int64)
	pg_QTreeWidget_connect_itemDoubleClicked func(uintptr, int64)
	pg_QTreeWidget_connect_itemActivated func(uintptr, int64)
	pg_QTreeWidget_connect_itemEntered func(uintptr, int64)
	pg_QTreeWidget_connect_itemChanged func(uintptr, int64)
	pg_QTreeWidget_connect_itemExpanded func(uintptr, int64)
	pg_QTreeWidget_connect_itemCollapsed func(uintptr, int64)
	pg_QTreeWidget_connect_currentItemChanged func(uintptr, int64)
	pg_QTreeWidget_connect_itemSelectionChanged func(uintptr, int64)
	pg_QTreeWidget_connect_expanded func(uintptr, int64)
	pg_QTreeWidget_connect_collapsed func(uintptr, int64)
	pg_QTreeWidget_connect_pressed func(uintptr, int64)
	pg_QTreeWidget_connect_clicked func(uintptr, int64)
	pg_QTreeWidget_connect_doubleClicked func(uintptr, int64)
	pg_QTreeWidget_connect_activated func(uintptr, int64)
	pg_QTreeWidget_connect_entered func(uintptr, int64)
	pg_QTreeWidget_connect_viewportEntered func(uintptr, int64)
	pg_QTreeWidget_connect_iconSizeChanged func(uintptr, int64)
	pg_QTreeWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QTreeWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QTreeWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTreeWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTreeWidget_connect_destroyed func(uintptr, int64)
	pg_QTreeWidget_connect_destroyed2 func(uintptr, int64)
	pg_QTreeWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QTreeWidget__disconnect func(int64)
	pg_QTreeWidget__set_override func(uintptr, int32)
	pg_QTreeWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTreeWidget_chain, "QTreeWidget_chain")
	symaddr(&addr_QTreeWidget_new, "QTreeWidget_new")
	symaddr(&addr_QTreeWidget_new2, "QTreeWidget_new2")
	symaddr(&addr_QTreeWidget_tr, "QTreeWidget_tr")
	symaddr(&addr_QTreeWidget_columnCount, "QTreeWidget_columnCount")
	symaddr(&addr_QTreeWidget_setColumnCount, "QTreeWidget_setColumnCount")
	symaddr(&addr_QTreeWidget_invisibleRootItem, "QTreeWidget_invisibleRootItem")
	symaddr(&addr_QTreeWidget_topLevelItem, "QTreeWidget_topLevelItem")
	symaddr(&addr_QTreeWidget_topLevelItemCount, "QTreeWidget_topLevelItemCount")
	symaddr(&addr_QTreeWidget_insertTopLevelItem, "QTreeWidget_insertTopLevelItem")
	symaddr(&addr_QTreeWidget_addTopLevelItem, "QTreeWidget_addTopLevelItem")
	symaddr(&addr_QTreeWidget_takeTopLevelItem, "QTreeWidget_takeTopLevelItem")
	symaddr(&addr_QTreeWidget_indexOfTopLevelItem, "QTreeWidget_indexOfTopLevelItem")
	symaddr(&addr_QTreeWidget_insertTopLevelItems, "QTreeWidget_insertTopLevelItems")
	symaddr(&addr_QTreeWidget_addTopLevelItems, "QTreeWidget_addTopLevelItems")
	symaddr(&addr_QTreeWidget_headerItem, "QTreeWidget_headerItem")
	symaddr(&addr_QTreeWidget_setHeaderItem, "QTreeWidget_setHeaderItem")
	symaddr(&addr_QTreeWidget_setHeaderLabels, "QTreeWidget_setHeaderLabels")
	symaddr(&addr_QTreeWidget_setHeaderLabel, "QTreeWidget_setHeaderLabel")
	symaddr(&addr_QTreeWidget_currentItem, "QTreeWidget_currentItem")
	symaddr(&addr_QTreeWidget_currentColumn, "QTreeWidget_currentColumn")
	symaddr(&addr_QTreeWidget_setCurrentItem, "QTreeWidget_setCurrentItem")
	symaddr(&addr_QTreeWidget_setCurrentItem2, "QTreeWidget_setCurrentItem2")
	symaddr(&addr_QTreeWidget_setCurrentItem3, "QTreeWidget_setCurrentItem3")
	symaddr(&addr_QTreeWidget_itemAt, "QTreeWidget_itemAt")
	symaddr(&addr_QTreeWidget_itemAt2, "QTreeWidget_itemAt2")
	symaddr(&addr_QTreeWidget_visualItemRect, "QTreeWidget_visualItemRect")
	symaddr(&addr_QTreeWidget_sortColumn, "QTreeWidget_sortColumn")
	symaddr(&addr_QTreeWidget_sortItems, "QTreeWidget_sortItems")
	symaddr(&addr_QTreeWidget_editItem, "QTreeWidget_editItem")
	symaddr(&addr_QTreeWidget_openPersistentEditor, "QTreeWidget_openPersistentEditor")
	symaddr(&addr_QTreeWidget_closePersistentEditor, "QTreeWidget_closePersistentEditor")
	symaddr(&addr_QTreeWidget_isPersistentEditorOpen, "QTreeWidget_isPersistentEditorOpen")
	symaddr(&addr_QTreeWidget_itemWidget, "QTreeWidget_itemWidget")
	symaddr(&addr_QTreeWidget_setItemWidget, "QTreeWidget_setItemWidget")
	symaddr(&addr_QTreeWidget_removeItemWidget, "QTreeWidget_removeItemWidget")
	symaddr(&addr_QTreeWidget_selectedItems, "QTreeWidget_selectedItems")
	symaddr(&addr_QTreeWidget_findItems, "QTreeWidget_findItems")
	symaddr(&addr_QTreeWidget_itemAbove, "QTreeWidget_itemAbove")
	symaddr(&addr_QTreeWidget_itemBelow, "QTreeWidget_itemBelow")
	symaddr(&addr_QTreeWidget_indexFromItem, "QTreeWidget_indexFromItem")
	symaddr(&addr_QTreeWidget_itemFromIndex, "QTreeWidget_itemFromIndex")
	symaddr(&addr_QTreeWidget_supportedDragActions, "QTreeWidget_supportedDragActions")
	symaddr(&addr_QTreeWidget_setSupportedDragActions, "QTreeWidget_setSupportedDragActions")
	symaddr(&addr_QTreeWidget_scrollToItem, "QTreeWidget_scrollToItem")
	symaddr(&addr_QTreeWidget_expandItem, "QTreeWidget_expandItem")
	symaddr(&addr_QTreeWidget_collapseItem, "QTreeWidget_collapseItem")
	symaddr(&addr_QTreeWidget_clear, "QTreeWidget_clear")
	symaddr(&addr_QTreeWidget_itemPressed, "QTreeWidget_itemPressed")
	symaddr(&addr_QTreeWidget_itemClicked, "QTreeWidget_itemClicked")
	symaddr(&addr_QTreeWidget_itemDoubleClicked, "QTreeWidget_itemDoubleClicked")
	symaddr(&addr_QTreeWidget_itemActivated, "QTreeWidget_itemActivated")
	symaddr(&addr_QTreeWidget_itemEntered, "QTreeWidget_itemEntered")
	symaddr(&addr_QTreeWidget_itemChanged, "QTreeWidget_itemChanged")
	symaddr(&addr_QTreeWidget_itemExpanded, "QTreeWidget_itemExpanded")
	symaddr(&addr_QTreeWidget_itemCollapsed, "QTreeWidget_itemCollapsed")
	symaddr(&addr_QTreeWidget_currentItemChanged, "QTreeWidget_currentItemChanged")
	symaddr(&addr_QTreeWidget_itemSelectionChanged, "QTreeWidget_itemSelectionChanged")
	symaddr(&addr_QTreeWidget_columnResized, "QTreeWidget_columnResized")
	symaddr(&addr_QTreeWidget_columnCountChanged, "QTreeWidget_columnCountChanged")
	symaddr(&addr_QTreeWidget_columnMoved, "QTreeWidget_columnMoved")
	symaddr(&addr_QTreeWidget_reexpand, "QTreeWidget_reexpand")
	symaddr(&addr_QTreeWidget_rowsRemoved, "QTreeWidget_rowsRemoved")
	symaddr(&addr_QTreeWidget_drawTree, "QTreeWidget_drawTree")
	symaddr(&addr_QTreeWidget_indexRowSizeHint, "QTreeWidget_indexRowSizeHint")
	symaddr(&addr_QTreeWidget_rowHeight, "QTreeWidget_rowHeight")
	symaddr(&addr_QTreeWidget_state, "QTreeWidget_state")
	symaddr(&addr_QTreeWidget_scheduleDelayedItemsLayout, "QTreeWidget_scheduleDelayedItemsLayout")
	symaddr(&addr_QTreeWidget_executeDelayedItemsLayout, "QTreeWidget_executeDelayedItemsLayout")
	symaddr(&addr_QTreeWidget_setDirtyRegion, "QTreeWidget_setDirtyRegion")
	symaddr(&addr_QTreeWidget_scrollDirtyRegion, "QTreeWidget_scrollDirtyRegion")
	symaddr(&addr_QTreeWidget_dirtyRegionOffset, "QTreeWidget_dirtyRegionOffset")
	symaddr(&addr_QTreeWidget_startAutoScroll, "QTreeWidget_startAutoScroll")
	symaddr(&addr_QTreeWidget_stopAutoScroll, "QTreeWidget_stopAutoScroll")
	symaddr(&addr_QTreeWidget_doAutoScroll, "QTreeWidget_doAutoScroll")
	symaddr(&addr_QTreeWidget_dropIndicatorPosition, "QTreeWidget_dropIndicatorPosition")
	symaddr(&addr_QTreeWidget_setViewportMargins, "QTreeWidget_setViewportMargins")
	symaddr(&addr_QTreeWidget_setViewportMargins2, "QTreeWidget_setViewportMargins2")
	symaddr(&addr_QTreeWidget_viewportMargins, "QTreeWidget_viewportMargins")
	symaddr(&addr_QTreeWidget_drawFrame, "QTreeWidget_drawFrame")
	symaddr(&addr_QTreeWidget_setTabOrder, "QTreeWidget_setTabOrder")
	symaddr(&addr_QTreeWidget_mouseGrabber, "QTreeWidget_mouseGrabber")
	symaddr(&addr_QTreeWidget_keyboardGrabber, "QTreeWidget_keyboardGrabber")
	symaddr(&addr_QTreeWidget_find, "QTreeWidget_find")
	symaddr(&addr_QTreeWidget_createWindowContainer, "QTreeWidget_createWindowContainer")
	symaddr(&addr_QTreeWidget_updateMicroFocus, "QTreeWidget_updateMicroFocus")
	symaddr(&addr_QTreeWidget_create, "QTreeWidget_create")
	symaddr(&addr_QTreeWidget_destroy, "QTreeWidget_destroy")
	symaddr(&addr_QTreeWidget_focusNextChild, "QTreeWidget_focusNextChild")
	symaddr(&addr_QTreeWidget_focusPreviousChild, "QTreeWidget_focusPreviousChild")
	symaddr(&addr_QTreeWidget_connect, "QTreeWidget_connect")
	symaddr(&addr_QTreeWidget_connect2, "QTreeWidget_connect2")
	symaddr(&addr_QTreeWidget_sender, "QTreeWidget_sender")
	symaddr(&addr_QTreeWidget_senderSignalIndex, "QTreeWidget_senderSignalIndex")
	symaddr(&addr_QTreeWidget_receivers, "QTreeWidget_receivers")
	symaddr(&addr_QTreeWidget_isSignalConnected, "QTreeWidget_isSignalConnected")
	reg(&pg_QTreeWidget_devicePixelRatioFScale, "QTreeWidget_devicePixelRatioFScale")
	reg(&pg_QTreeWidget_encodeMetricF, "QTreeWidget_encodeMetricF")
	reg(&pg_QTreeWidget_getDecodedMetricF, "QTreeWidget_getDecodedMetricF")
	symaddr(&addr_QTreeWidget_tr2, "QTreeWidget_tr2")
	symaddr(&addr_QTreeWidget_tr3, "QTreeWidget_tr3")
	symaddr(&addr_QTreeWidget_editItem2, "QTreeWidget_editItem2")
	symaddr(&addr_QTreeWidget_openPersistentEditor2, "QTreeWidget_openPersistentEditor2")
	symaddr(&addr_QTreeWidget_closePersistentEditor2, "QTreeWidget_closePersistentEditor2")
	symaddr(&addr_QTreeWidget_isPersistentEditorOpen2, "QTreeWidget_isPersistentEditorOpen2")
	symaddr(&addr_QTreeWidget_findItems2, "QTreeWidget_findItems2")
	symaddr(&addr_QTreeWidget_indexFromItem2, "QTreeWidget_indexFromItem2")
	symaddr(&addr_QTreeWidget_scrollToItem2, "QTreeWidget_scrollToItem2")
	symaddr(&addr_QTreeWidget_createWindowContainer2, "QTreeWidget_createWindowContainer2")
	symaddr(&addr_QTreeWidget_createWindowContainer3, "QTreeWidget_createWindowContainer3")
	symaddr(&addr_QTreeWidget_updateMicroFocus2, "QTreeWidget_updateMicroFocus2")
	symaddr(&addr_QTreeWidget_create2, "QTreeWidget_create2")
	symaddr(&addr_QTreeWidget_create3, "QTreeWidget_create3")
	symaddr(&addr_QTreeWidget_create4, "QTreeWidget_create4")
	symaddr(&addr_QTreeWidget_destroy2, "QTreeWidget_destroy2")
	symaddr(&addr_QTreeWidget_destroy3, "QTreeWidget_destroy3")
	symaddr(&addr_QTreeWidget_connect3, "QTreeWidget_connect3")
	symaddr(&addr_QTreeWidget_connect4, "QTreeWidget_connect4")
	symaddr(&addr_QTreeWidget_metaObject_default, "QTreeWidget_metaObject_default")
	symaddr(&addr_QTreeWidget_setSelectionModel_default, "QTreeWidget_setSelectionModel_default")
	symaddr(&addr_QTreeWidget_event_default, "QTreeWidget_event_default")
	symaddr(&addr_QTreeWidget_mimeTypes_default, "QTreeWidget_mimeTypes_default")
	symaddr(&addr_QTreeWidget_dropMimeData_default, "QTreeWidget_dropMimeData_default")
	symaddr(&addr_QTreeWidget_supportedDropActions_default, "QTreeWidget_supportedDropActions_default")
	symaddr(&addr_QTreeWidget_dropEvent_default, "QTreeWidget_dropEvent_default")
	symaddr(&addr_QTreeWidget_setRootIndex_default, "QTreeWidget_setRootIndex_default")
	symaddr(&addr_QTreeWidget_keyboardSearch_default, "QTreeWidget_keyboardSearch_default")
	symaddr(&addr_QTreeWidget_visualRect_default, "QTreeWidget_visualRect_default")
	symaddr(&addr_QTreeWidget_scrollTo_default, "QTreeWidget_scrollTo_default")
	symaddr(&addr_QTreeWidget_indexAt_default, "QTreeWidget_indexAt_default")
	symaddr(&addr_QTreeWidget_doItemsLayout_default, "QTreeWidget_doItemsLayout_default")
	symaddr(&addr_QTreeWidget_reset_default, "QTreeWidget_reset_default")
	symaddr(&addr_QTreeWidget_selectAll_default, "QTreeWidget_selectAll_default")
	symaddr(&addr_QTreeWidget_verticalScrollbarValueChanged_default, "QTreeWidget_verticalScrollbarValueChanged_default")
	symaddr(&addr_QTreeWidget_scrollContentsBy_default, "QTreeWidget_scrollContentsBy_default")
	symaddr(&addr_QTreeWidget_rowsInserted_default, "QTreeWidget_rowsInserted_default")
	symaddr(&addr_QTreeWidget_rowsAboutToBeRemoved_default, "QTreeWidget_rowsAboutToBeRemoved_default")
	symaddr(&addr_QTreeWidget_moveCursor_default, "QTreeWidget_moveCursor_default")
	symaddr(&addr_QTreeWidget_horizontalOffset_default, "QTreeWidget_horizontalOffset_default")
	symaddr(&addr_QTreeWidget_verticalOffset_default, "QTreeWidget_verticalOffset_default")
	symaddr(&addr_QTreeWidget_setSelection_default, "QTreeWidget_setSelection_default")
	symaddr(&addr_QTreeWidget_visualRegionForSelection_default, "QTreeWidget_visualRegionForSelection_default")
	symaddr(&addr_QTreeWidget_selectedIndexes_default, "QTreeWidget_selectedIndexes_default")
	symaddr(&addr_QTreeWidget_changeEvent_default, "QTreeWidget_changeEvent_default")
	symaddr(&addr_QTreeWidget_timerEvent_default, "QTreeWidget_timerEvent_default")
	symaddr(&addr_QTreeWidget_paintEvent_default, "QTreeWidget_paintEvent_default")
	symaddr(&addr_QTreeWidget_drawRow_default, "QTreeWidget_drawRow_default")
	symaddr(&addr_QTreeWidget_drawBranches_default, "QTreeWidget_drawBranches_default")
	symaddr(&addr_QTreeWidget_mousePressEvent_default, "QTreeWidget_mousePressEvent_default")
	symaddr(&addr_QTreeWidget_mouseReleaseEvent_default, "QTreeWidget_mouseReleaseEvent_default")
	symaddr(&addr_QTreeWidget_mouseDoubleClickEvent_default, "QTreeWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QTreeWidget_mouseMoveEvent_default, "QTreeWidget_mouseMoveEvent_default")
	symaddr(&addr_QTreeWidget_keyPressEvent_default, "QTreeWidget_keyPressEvent_default")
	symaddr(&addr_QTreeWidget_dragMoveEvent_default, "QTreeWidget_dragMoveEvent_default")
	symaddr(&addr_QTreeWidget_viewportEvent_default, "QTreeWidget_viewportEvent_default")
	symaddr(&addr_QTreeWidget_updateGeometries_default, "QTreeWidget_updateGeometries_default")
	symaddr(&addr_QTreeWidget_viewportSizeHint_default, "QTreeWidget_viewportSizeHint_default")
	symaddr(&addr_QTreeWidget_sizeHintForColumn_default, "QTreeWidget_sizeHintForColumn_default")
	symaddr(&addr_QTreeWidget_horizontalScrollbarAction_default, "QTreeWidget_horizontalScrollbarAction_default")
	symaddr(&addr_QTreeWidget_isIndexHidden_default, "QTreeWidget_isIndexHidden_default")
	symaddr(&addr_QTreeWidget_selectionChanged_default, "QTreeWidget_selectionChanged_default")
	symaddr(&addr_QTreeWidget_currentChanged_default, "QTreeWidget_currentChanged_default")
	symaddr(&addr_QTreeWidget_sizeHintForRow_default, "QTreeWidget_sizeHintForRow_default")
	symaddr(&addr_QTreeWidget_itemDelegateForIndex_default, "QTreeWidget_itemDelegateForIndex_default")
	symaddr(&addr_QTreeWidget_inputMethodQuery_default, "QTreeWidget_inputMethodQuery_default")
	symaddr(&addr_QTreeWidget_updateEditorData_default, "QTreeWidget_updateEditorData_default")
	symaddr(&addr_QTreeWidget_updateEditorGeometries_default, "QTreeWidget_updateEditorGeometries_default")
	symaddr(&addr_QTreeWidget_verticalScrollbarAction_default, "QTreeWidget_verticalScrollbarAction_default")
	symaddr(&addr_QTreeWidget_horizontalScrollbarValueChanged_default, "QTreeWidget_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QTreeWidget_closeEditor_default, "QTreeWidget_closeEditor_default")
	symaddr(&addr_QTreeWidget_commitData_default, "QTreeWidget_commitData_default")
	symaddr(&addr_QTreeWidget_editorDestroyed_default, "QTreeWidget_editorDestroyed_default")
	symaddr(&addr_QTreeWidget_edit_default, "QTreeWidget_edit_default")
	symaddr(&addr_QTreeWidget_selectionCommand_default, "QTreeWidget_selectionCommand_default")
	symaddr(&addr_QTreeWidget_startDrag_default, "QTreeWidget_startDrag_default")
	symaddr(&addr_QTreeWidget_initViewItemOption_default, "QTreeWidget_initViewItemOption_default")
	symaddr(&addr_QTreeWidget_focusNextPrevChild_default, "QTreeWidget_focusNextPrevChild_default")
	symaddr(&addr_QTreeWidget_dragEnterEvent_default, "QTreeWidget_dragEnterEvent_default")
	symaddr(&addr_QTreeWidget_dragLeaveEvent_default, "QTreeWidget_dragLeaveEvent_default")
	symaddr(&addr_QTreeWidget_focusInEvent_default, "QTreeWidget_focusInEvent_default")
	symaddr(&addr_QTreeWidget_focusOutEvent_default, "QTreeWidget_focusOutEvent_default")
	symaddr(&addr_QTreeWidget_resizeEvent_default, "QTreeWidget_resizeEvent_default")
	symaddr(&addr_QTreeWidget_inputMethodEvent_default, "QTreeWidget_inputMethodEvent_default")
	symaddr(&addr_QTreeWidget_eventFilter_default, "QTreeWidget_eventFilter_default")
	symaddr(&addr_QTreeWidget_minimumSizeHint_default, "QTreeWidget_minimumSizeHint_default")
	symaddr(&addr_QTreeWidget_sizeHint_default, "QTreeWidget_sizeHint_default")
	symaddr(&addr_QTreeWidget_setupViewport_default, "QTreeWidget_setupViewport_default")
	symaddr(&addr_QTreeWidget_wheelEvent_default, "QTreeWidget_wheelEvent_default")
	symaddr(&addr_QTreeWidget_contextMenuEvent_default, "QTreeWidget_contextMenuEvent_default")
	symaddr(&addr_QTreeWidget_initStyleOption_default, "QTreeWidget_initStyleOption_default")
	symaddr(&addr_QTreeWidget_devType_default, "QTreeWidget_devType_default")
	symaddr(&addr_QTreeWidget_setVisible_default, "QTreeWidget_setVisible_default")
	symaddr(&addr_QTreeWidget_heightForWidth_default, "QTreeWidget_heightForWidth_default")
	symaddr(&addr_QTreeWidget_hasHeightForWidth_default, "QTreeWidget_hasHeightForWidth_default")
	symaddr(&addr_QTreeWidget_paintEngine_default, "QTreeWidget_paintEngine_default")
	symaddr(&addr_QTreeWidget_keyReleaseEvent_default, "QTreeWidget_keyReleaseEvent_default")
	symaddr(&addr_QTreeWidget_enterEvent_default, "QTreeWidget_enterEvent_default")
	symaddr(&addr_QTreeWidget_leaveEvent_default, "QTreeWidget_leaveEvent_default")
	symaddr(&addr_QTreeWidget_moveEvent_default, "QTreeWidget_moveEvent_default")
	symaddr(&addr_QTreeWidget_closeEvent_default, "QTreeWidget_closeEvent_default")
	symaddr(&addr_QTreeWidget_tabletEvent_default, "QTreeWidget_tabletEvent_default")
	symaddr(&addr_QTreeWidget_actionEvent_default, "QTreeWidget_actionEvent_default")
	symaddr(&addr_QTreeWidget_showEvent_default, "QTreeWidget_showEvent_default")
	symaddr(&addr_QTreeWidget_hideEvent_default, "QTreeWidget_hideEvent_default")
	symaddr(&addr_QTreeWidget_metric_default, "QTreeWidget_metric_default")
	symaddr(&addr_QTreeWidget_initPainter_default, "QTreeWidget_initPainter_default")
	symaddr(&addr_QTreeWidget_redirected_default, "QTreeWidget_redirected_default")
	symaddr(&addr_QTreeWidget_sharedPainter_default, "QTreeWidget_sharedPainter_default")
	symaddr(&addr_QTreeWidget_childEvent_default, "QTreeWidget_childEvent_default")
	symaddr(&addr_QTreeWidget_customEvent_default, "QTreeWidget_customEvent_default")
	symaddr(&addr_QTreeWidget_connectNotify_default, "QTreeWidget_connectNotify_default")
	symaddr(&addr_QTreeWidget_disconnectNotify_default, "QTreeWidget_disconnectNotify_default")
	reg(&pg_QTreeWidget_connect_itemPressed, "QTreeWidget_connect_itemPressed")
	reg(&pg_QTreeWidget_connect_itemClicked, "QTreeWidget_connect_itemClicked")
	reg(&pg_QTreeWidget_connect_itemDoubleClicked, "QTreeWidget_connect_itemDoubleClicked")
	reg(&pg_QTreeWidget_connect_itemActivated, "QTreeWidget_connect_itemActivated")
	reg(&pg_QTreeWidget_connect_itemEntered, "QTreeWidget_connect_itemEntered")
	reg(&pg_QTreeWidget_connect_itemChanged, "QTreeWidget_connect_itemChanged")
	reg(&pg_QTreeWidget_connect_itemExpanded, "QTreeWidget_connect_itemExpanded")
	reg(&pg_QTreeWidget_connect_itemCollapsed, "QTreeWidget_connect_itemCollapsed")
	reg(&pg_QTreeWidget_connect_currentItemChanged, "QTreeWidget_connect_currentItemChanged")
	reg(&pg_QTreeWidget_connect_itemSelectionChanged, "QTreeWidget_connect_itemSelectionChanged")
	reg(&pg_QTreeWidget_connect_expanded, "QTreeWidget_connect_expanded")
	reg(&pg_QTreeWidget_connect_collapsed, "QTreeWidget_connect_collapsed")
	reg(&pg_QTreeWidget_connect_pressed, "QTreeWidget_connect_pressed")
	reg(&pg_QTreeWidget_connect_clicked, "QTreeWidget_connect_clicked")
	reg(&pg_QTreeWidget_connect_doubleClicked, "QTreeWidget_connect_doubleClicked")
	reg(&pg_QTreeWidget_connect_activated, "QTreeWidget_connect_activated")
	reg(&pg_QTreeWidget_connect_entered, "QTreeWidget_connect_entered")
	reg(&pg_QTreeWidget_connect_viewportEntered, "QTreeWidget_connect_viewportEntered")
	reg(&pg_QTreeWidget_connect_iconSizeChanged, "QTreeWidget_connect_iconSizeChanged")
	reg(&pg_QTreeWidget_connect_windowTitleChanged, "QTreeWidget_connect_windowTitleChanged")
	reg(&pg_QTreeWidget_connect_windowIconChanged, "QTreeWidget_connect_windowIconChanged")
	reg(&pg_QTreeWidget_connect_windowIconTextChanged, "QTreeWidget_connect_windowIconTextChanged")
	reg(&pg_QTreeWidget_connect_customContextMenuRequested, "QTreeWidget_connect_customContextMenuRequested")
	reg(&pg_QTreeWidget_connect_destroyed, "QTreeWidget_connect_destroyed")
	reg(&pg_QTreeWidget_connect_destroyed2, "QTreeWidget_connect_destroyed2")
	reg(&pg_QTreeWidget_connect_objectNameChanged, "QTreeWidget_connect_objectNameChanged")
	reg(&pg_QTreeWidget__disconnect, "QTreeWidget__disconnect")
	reg(&pg_QTreeWidget__set_override, "QTreeWidget__set_override")
	reg(&pg_QTreeWidget_delete, "QTreeWidget_delete")
}

type QTreeWidget struct {
	h uintptr
	*QTreeView
}

func bindQTreeWidget(h uintptr) *QTreeWidget {
	if h == 0 {
		return nil
	}
	var __c [6]uintptr
	purego.SyscallN(addr_QTreeWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTreeWidget{h: h, QTreeView: &QTreeView{h: __c[0], QAbstractItemView: &QAbstractItemView{h: __c[1], QAbstractScrollArea: &QAbstractScrollArea{h: __c[2], QFrame: &QFrame{h: __c[3], QWidget: &QWidget{h: __c[4], QObject: &QObject{h: __c[5]}}}}}}}
}

func newownedQTreeWidget(h uintptr) *QTreeWidget { return bindQTreeWidget(h) }

func (this *QTreeWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTreeWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTreeWidget(h unsafe.Pointer) *QTreeWidget { return bindQTreeWidget(uintptr(h)) }
func UnsafeNewQTreeWidgetOwned(h unsafe.Pointer) *QTreeWidget { return newownedQTreeWidget(uintptr(h)) }

func (this *QTreeWidget) UnsafeRelease() {
	this.h = 0
	if this.QTreeView != nil {
		this.QTreeView.UnsafeRelease()
	}
}

func NewQTreeWidget(parent *QWidget) *QTreeWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTreeWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidget_virtIdx, QTreeWidget_enable)
	return __o
}

func NewQTreeWidget2() *QTreeWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_new2, uintptr(__goid))
	__o := newownedQTreeWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidget_virtIdx, QTreeWidget_enable)
	return __o
}

func QTreeWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTreeWidget) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) SetColumnCount(columns int) {
	purego.SyscallN(addr_QTreeWidget_setColumnCount, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) InvisibleRootItem() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_invisibleRootItem, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) TopLevelItem(index int) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_topLevelItem, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) TopLevelItemCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_topLevelItemCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) InsertTopLevelItem(index int, item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_insertTopLevelItem, this.h, uintptr(index), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) AddTopLevelItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_addTopLevelItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) TakeTopLevelItem(index int) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_takeTopLevelItem, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) IndexOfTopLevelItem(item *QTreeWidgetItem) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_indexOfTopLevelItem, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return int(__r)
}

func (this *QTreeWidget) InsertTopLevelItems(index int, items []*QTreeWidgetItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QTreeWidget_insertTopLevelItems, this.h, uintptr(index), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QTreeWidget) AddTopLevelItems(items []*QTreeWidgetItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QTreeWidget_addTopLevelItems, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QTreeWidget) HeaderItem() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_headerItem, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) SetHeaderItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_setHeaderItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) SetHeaderLabels(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QTreeWidget_setHeaderLabels, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QTreeWidget) SetHeaderLabel(label string) {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	purego.SyscallN(addr_QTreeWidget_setHeaderLabel, this.h, label_p, uintptr(len(label)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
}

func (this *QTreeWidget) CurrentItem() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_currentItem, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) CurrentColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_currentColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) SetCurrentItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_setCurrentItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) SetCurrentItem2(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_setCurrentItem2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) SetCurrentItem3(item *QTreeWidgetItem, column int, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTreeWidget_setCurrentItem3, this.h, uintptr(item.UnsafePointer()), uintptr(column), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemAt(p *QPoint) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemAt, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) ItemAt2(x int, y int) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemAt2, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) VisualItemRect(item *QTreeWidgetItem) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_visualItemRect, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRect(__r)
}

func (this *QTreeWidget) SortColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_sortColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) SortItems(column int, order SortOrder) {
	purego.SyscallN(addr_QTreeWidget_sortItems, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) EditItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_editItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) OpenPersistentEditor(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_openPersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ClosePersistentEditor(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_closePersistentEditor, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) IsPersistentEditorOpen(item *QTreeWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_isPersistentEditorOpen, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QTreeWidget) ItemWidget(item *QTreeWidgetItem, column int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemWidget, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQWidget(__r)
}

func (this *QTreeWidget) SetItemWidget(item *QTreeWidgetItem, column int, widget *QWidget) {
	purego.SyscallN(addr_QTreeWidget_setItemWidget, this.h, uintptr(item.UnsafePointer()), uintptr(column), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(widget)
}

func (this *QTreeWidget) RemoveItemWidget(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_removeItemWidget, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) SelectedItems() []*QTreeWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_selectedItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTreeWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTreeWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeWidget) FindItems(text string, flags MatchFlag) []*QTreeWidgetItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_findItems, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QTreeWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTreeWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeWidget) ItemAbove(item *QTreeWidgetItem) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemAbove, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) ItemBelow(item *QTreeWidgetItem) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemBelow, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) IndexFromItem(item *QTreeWidgetItem) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_indexFromItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQModelIndex(__r)
}

func (this *QTreeWidget) ItemFromIndex(index *QModelIndex) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemFromIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidget) SupportedDragActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_supportedDragActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTreeWidget) SetSupportedDragActions(actions DropAction) {
	purego.SyscallN(addr_QTreeWidget_setSupportedDragActions, this.h, uintptr(actions))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ScrollToItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_scrollToItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ExpandItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_expandItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) CollapseItem(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_collapseItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) Clear() {
	purego.SyscallN(addr_QTreeWidget_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ItemPressed(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemPressed, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemClicked(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemClicked, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemDoubleClicked(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemDoubleClicked, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemActivated(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemActivated, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemEntered(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemEntered, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemChanged(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_itemChanged, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemExpanded(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_itemExpanded, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ItemCollapsed(item *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_itemCollapsed, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) CurrentItemChanged(current *QTreeWidgetItem, previous *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidget_currentItemChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTreeWidget) ItemSelectionChanged() {
	purego.SyscallN(addr_QTreeWidget_itemSelectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ColumnResized(column int, oldSize int, newSize int) {
	purego.SyscallN(addr_QTreeWidget_columnResized, this.h, uintptr(column), uintptr(oldSize), uintptr(newSize))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ColumnCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QTreeWidget_columnCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ColumnMoved() {
	purego.SyscallN(addr_QTreeWidget_columnMoved, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Reexpand() {
	purego.SyscallN(addr_QTreeWidget_reexpand, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) RowsRemoved(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTreeWidget_rowsRemoved, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeWidget) DrawTree(painter *QPainter, region *QRegion) {
	purego.SyscallN(addr_QTreeWidget_drawTree, this.h, uintptr(painter.UnsafePointer()), uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(region)
}

func (this *QTreeWidget) IndexRowSizeHint(index *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_indexRowSizeHint, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return int(__r)
}

func (this *QTreeWidget) RowHeight(index *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_rowHeight, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return int(__r)
}

func (this *QTreeWidget) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QTreeWidget) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QTreeWidget_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QTreeWidget_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QTreeWidget_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QTreeWidget) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QTreeWidget_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QTreeWidget) StartAutoScroll() {
	purego.SyscallN(addr_QTreeWidget_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) StopAutoScroll() {
	purego.SyscallN(addr_QTreeWidget_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) DoAutoScroll() {
	purego.SyscallN(addr_QTreeWidget_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QTreeWidget) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTreeWidget_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTreeWidget_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTreeWidget) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTreeWidget) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTreeWidget_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTreeWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTreeWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTreeWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QTreeWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QTreeWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTreeWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTreeWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QTreeWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Create() {
	purego.SyscallN(addr_QTreeWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Destroy() {
	purego.SyscallN(addr_QTreeWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTreeWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTreeWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTreeWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTreeWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTreeWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTreeWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QTreeWidget_devicePixelRatioFScale()
	return __r
}

func QTreeWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTreeWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTreeWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTreeWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTreeWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTreeWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTreeWidget) EditItem2(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_editItem2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) OpenPersistentEditor2(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_openPersistentEditor2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) ClosePersistentEditor2(item *QTreeWidgetItem, column int) {
	purego.SyscallN(addr_QTreeWidget_closePersistentEditor2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QTreeWidget) IsPersistentEditorOpen2(item *QTreeWidgetItem, column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_isPersistentEditorOpen2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QTreeWidget) FindItems2(text string, flags MatchFlag, column int) []*QTreeWidgetItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_findItems2, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QTreeWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTreeWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeWidget) IndexFromItem2(item *QTreeWidgetItem, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_indexFromItem2, this.h, uintptr(item.UnsafePointer()), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQModelIndex(__r)
}

func (this *QTreeWidget) ScrollToItem2(item *QTreeWidgetItem, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTreeWidget_scrollToItem2, this.h, uintptr(item.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func QTreeWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTreeWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTreeWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTreeWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTreeWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTreeWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTreeWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTreeWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTreeWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTreeWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTreeWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTreeWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTreeWidget) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QTreeWidget_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QTreeWidget) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QTreeWidget) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QTreeWidget) DropMimeDataDefault(parent *QTreeWidgetItem, index int, data *QMimeData, action DropAction) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_dropMimeData_default, this.h, uintptr(parent.UnsafePointer()), uintptr(index), uintptr(data.UnsafePointer()), uintptr(action))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QTreeWidget) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTreeWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTreeWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QTreeWidget_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeWidget) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QTreeWidget_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QTreeWidget) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QTreeWidget) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QTreeWidget_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QTreeWidget) IndexAtDefault(p *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_indexAt_default, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQModelIndex(__r)
}

func (this *QTreeWidget) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QTreeWidget_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ResetDefault() {
	purego.SyscallN(addr_QTreeWidget_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) SelectAllDefault() {
	purego.SyscallN(addr_QTreeWidget_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTreeWidget_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTreeWidget_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTreeWidget_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeWidget) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QTreeWidget_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTreeWidget) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTreeWidget) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QTreeWidget_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QTreeWidget) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QTreeWidget) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeWidget) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeWidget_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTreeWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QTreeWidget_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) DrawRowDefault(painter *QPainter, options *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QTreeWidget_drawRow_default, this.h, uintptr(painter.UnsafePointer()), uintptr(options.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(options)
	runtime.KeepAlive(index)
}

func (this *QTreeWidget) DrawBranchesDefault(painter *QPainter, rect *QRect, index *QModelIndex) {
	purego.SyscallN(addr_QTreeWidget_drawBranches_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(index)
}

func (this *QTreeWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTreeWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTreeWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTreeWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeWidget) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QTreeWidget_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeWidget) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTreeWidget_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTreeWidget) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QTreeWidget_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QTreeWidget) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QTreeWidget_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QTreeWidget) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QTreeWidget) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTreeWidget) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QTreeWidget_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QTreeWidget_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QTreeWidget_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QTreeWidget_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QTreeWidget_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeWidget) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QTreeWidget_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeWidget) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QTreeWidget_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QTreeWidget) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeWidget) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QTreeWidget) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QTreeWidget_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QTreeWidget_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTreeWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTreeWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTreeWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTreeWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTreeWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QTreeWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QTreeWidget_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTreeWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeWidget) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTreeWidget_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTreeWidget) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QTreeWidget_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTreeWidget) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QTreeWidget_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTreeWidget) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTreeWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTreeWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTreeWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTreeWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTreeWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTreeWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTreeWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTreeWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTreeWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTreeWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QTreeWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTreeWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTreeWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTreeWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTreeWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTreeWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTreeWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTreeWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTreeWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTreeWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTreeWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTreeWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTreeWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTreeWidget) OnItemPressed(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemClicked(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemDoubleClicked(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemActivated(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemEntered(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemChanged(fn func(*QTreeWidgetItem, int)) int64 {
	id := registerCallback(func(item uintptr, column int32) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item)), int(column)) })
	pg_QTreeWidget_connect_itemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemExpanded(fn func(*QTreeWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item))) })
	pg_QTreeWidget_connect_itemExpanded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemCollapsed(fn func(*QTreeWidgetItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(item))) })
	pg_QTreeWidget_connect_itemCollapsed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnCurrentItemChanged(fn func(*QTreeWidgetItem, *QTreeWidgetItem)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(current)), UnsafeNewQTreeWidgetItem(unsafe.Pointer(previous))) })
	pg_QTreeWidget_connect_currentItemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnItemSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTreeWidget_connect_itemSelectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnExpanded(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_expanded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnCollapsed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_collapsed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QTreeWidget_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTreeWidget_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QTreeWidget_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTreeWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTreeWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTreeWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTreeWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTreeWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTreeWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTreeWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTreeWidget) Disconnect(handle int64) {
	pg_QTreeWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QTreeWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -4689322347236366780:
		return 1
	case 1237236456312824991:
		return 2
	case 5464467312394730418:
		return 3
	case -9049994684305682212:
		return 4
	case -7784642176186821205:
		return 5
	case 9107058682754591716:
		return 6
	case -1574077111192992763:
		return 7
	case -5041452885022114200:
		return 8
	case -1267696851865372997:
		return 9
	case -4832340528568269947:
		return 10
	case -7252823847285716802:
		return 11
	case -1888629738259754754:
		return 12
	case 1086335023529244512:
		return 13
	case 8758113809152574300:
		return 14
	case 6366953527806415524:
		return 15
	case -6118001150126564745:
		return 16
	case 4472434241583495620:
		return 17
	case -8151049665126058157:
		return 18
	case -158523882194220890:
		return 19
	case -1500310024801917802:
		return 20
	case 5218707810677921016:
		return 21
	case -1539266838375037305:
		return 22
	case 7063795720810552074:
		return 23
	case 740869982805721908:
		return 24
	case 4959376063826165943:
		return 25
	case -5409927581812374498:
		return 26
	case 5290442163167458465:
		return 27
	case -3950604704165452999:
		return 28
	case 6369455307610057755:
		return 29
	case -332090898078543319:
		return 30
	case -4415920049297795213:
		return 31
	case -4029073905697023887:
		return 32
	case -4255494004099806581:
		return 33
	case 1023791977679793185:
		return 34
	case -1781545137352745978:
		return 35
	case -1971465653061861519:
		return 36
	case 1022771470898460576:
		return 37
	case 503579067252048239:
		return 38
	case -3666792239577137638:
		return 39
	case -785551405282638793:
		return 40
	case 6057989667089378859:
		return 41
	case -5892205462145366901:
		return 42
	case -6978583972112316034:
		return 43
	case -4084029023708556876:
		return 44
	case -3811438228839363824:
		return 45
	case 7345853097340709638:
		return 46
	case 9029941737538200621:
		return 47
	case -8726300139121449691:
		return 48
	case -1155555122391250623:
		return 49
	case 1786430177580376614:
		return 50
	case 8648484264628816806:
		return 51
	case -6021383853323633948:
		return 52
	case 6979418566323745651:
		return 53
	case 4206689019978932209:
		return 54
	case -2447490284072398302:
		return 55
	case -8112366604639385491:
		return 56
	case -7450787277202094556:
		return 57
	case 5898328517505784191:
		return 58
	case 206122715919917589:
		return 59
	case -3139904158373725006:
		return 60
	case 975043767033552552:
		return 61
	case -385596010576048993:
		return 62
	case -4302917285655014661:
		return 63
	case 5438819661942843222:
		return 64
	case 5217623514630858297:
		return 65
	case 7049264599837532491:
		return 66
	case 729690261457684101:
		return 67
	case 4531996381736730688:
		return 68
	case 1919988340934511316:
		return 69
	case -6698043222227006015:
		return 70
	case -488693769771308531:
		return 71
	case -8659679893868524044:
		return 72
	case 2110182719375245575:
		return 73
	case 4311534060495737339:
		return 74
	case -8266762397995423915:
		return 75
	case 6398007404756346381:
		return 76
	case 772598475950050299:
		return 77
	case 6316270346668237895:
		return 78
	case -551070009623711640:
		return 79
	case -9161949150489525396:
		return 80
	case -1989937050019464783:
		return 81
	case -212068866860701141:
		return 82
	case 1697444140088555581:
		return 83
	case -6825369174751916:
		return 84
	case -199718707876719129:
		return 85
	case -2039166586246443706:
		return 86
	case 427629649267311777:
		return 87
	case 6477500296980050210:
		return 88
	case 4127240633310655048:
		return 89
	case 5792771328741774991:
		return 90
	case -1185315367526857833:
		return 91
	case 3769868229123868474:
		return 92
	case 7379029982170503008:
		return 93
	case -568369772982215962:
		return 94
	}
	return -1
}

func QTreeWidget_enable(h uintptr, idx int32) { pg_QTreeWidget__set_override(h, idx) }

func (this *QTreeWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTreeWidget) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QTreeWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QTreeWidget) OnMimeTypes(fn func() []string) {
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

func (this *QTreeWidget) OnDropMimeData(fn func(*QTreeWidgetItem, int, *QMimeData, DropAction) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(parent uintptr, index int32, data uintptr, action int32) int32 { return b2i(fn(UnsafeNewQTreeWidgetItem(unsafe.Pointer(parent)), int(index), UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action))) }) // slot "dropMimeData"
}

func (this *QTreeWidget) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QTreeWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTreeWidget) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QTreeWidget) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QTreeWidget) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QTreeWidget) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QTreeWidget) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(p uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(p))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QTreeWidget) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QTreeWidget) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QTreeWidget) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QTreeWidget) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QTreeWidget) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTreeWidget) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QTreeWidget) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QTreeWidget) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QTreeWidget) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QTreeWidget) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QTreeWidget) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QTreeWidget) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QTreeWidget) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QTreeWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QTreeWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTreeWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QTreeWidget) OnDrawRow(fn func(*QPainter, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -3950604704165452999, func(painter uintptr, options uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(options)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "drawRow"
}

func (this *QTreeWidget) OnDrawBranches(fn func(*QPainter, *QRect, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 6369455307610057755, func(painter uintptr, rect uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "drawBranches"
}

func (this *QTreeWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QTreeWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QTreeWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTreeWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QTreeWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QTreeWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTreeWidget) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QTreeWidget) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QTreeWidget) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTreeWidget) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QTreeWidget) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QTreeWidget) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QTreeWidget) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QTreeWidget) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QTreeWidget) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QTreeWidget) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QTreeWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTreeWidget) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QTreeWidget) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QTreeWidget) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QTreeWidget) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QTreeWidget) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QTreeWidget) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QTreeWidget) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QTreeWidget) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QTreeWidget) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QTreeWidget) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QTreeWidget) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QTreeWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTreeWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTreeWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTreeWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTreeWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTreeWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QTreeWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QTreeWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTreeWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTreeWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTreeWidget) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTreeWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QTreeWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QTreeWidget) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTreeWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTreeWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTreeWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTreeWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTreeWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTreeWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTreeWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTreeWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTreeWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTreeWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTreeWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTreeWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTreeWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QTreeWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTreeWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTreeWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTreeWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTreeWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTreeWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTreeWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTreeWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTreeWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTreeWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTreeWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTreeWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
