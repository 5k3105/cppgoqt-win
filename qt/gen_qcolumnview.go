package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColumnView_chain uintptr
	addr_QColumnView_new uintptr
	addr_QColumnView_new2 uintptr
	addr_QColumnView_tr uintptr
	addr_QColumnView_updatePreviewWidget uintptr
	addr_QColumnView_setResizeGripsVisible uintptr
	addr_QColumnView_resizeGripsVisible uintptr
	addr_QColumnView_setPreviewColumnVisible uintptr
	addr_QColumnView_isPreviewColumnVisible uintptr
	addr_QColumnView_previewWidget uintptr
	addr_QColumnView_setPreviewWidget uintptr
	addr_QColumnView_setColumnWidths uintptr
	addr_QColumnView_columnWidths uintptr
	addr_QColumnView_initializeColumn uintptr
	addr_QColumnView_state uintptr
	addr_QColumnView_scheduleDelayedItemsLayout uintptr
	addr_QColumnView_executeDelayedItemsLayout uintptr
	addr_QColumnView_setDirtyRegion uintptr
	addr_QColumnView_scrollDirtyRegion uintptr
	addr_QColumnView_dirtyRegionOffset uintptr
	addr_QColumnView_startAutoScroll uintptr
	addr_QColumnView_stopAutoScroll uintptr
	addr_QColumnView_doAutoScroll uintptr
	addr_QColumnView_dropIndicatorPosition uintptr
	addr_QColumnView_setViewportMargins uintptr
	addr_QColumnView_setViewportMargins2 uintptr
	addr_QColumnView_viewportMargins uintptr
	addr_QColumnView_drawFrame uintptr
	addr_QColumnView_setTabOrder uintptr
	addr_QColumnView_mouseGrabber uintptr
	addr_QColumnView_keyboardGrabber uintptr
	addr_QColumnView_find uintptr
	addr_QColumnView_createWindowContainer uintptr
	addr_QColumnView_updateMicroFocus uintptr
	addr_QColumnView_create uintptr
	addr_QColumnView_destroy uintptr
	addr_QColumnView_focusNextChild uintptr
	addr_QColumnView_focusPreviousChild uintptr
	addr_QColumnView_connect uintptr
	addr_QColumnView_connect2 uintptr
	addr_QColumnView_sender uintptr
	addr_QColumnView_senderSignalIndex uintptr
	addr_QColumnView_receivers uintptr
	addr_QColumnView_isSignalConnected uintptr
	pg_QColumnView_devicePixelRatioFScale func() float64
	pg_QColumnView_encodeMetricF func(int32, float64) int32
	pg_QColumnView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QColumnView_tr2 uintptr
	addr_QColumnView_tr3 uintptr
	addr_QColumnView_createWindowContainer2 uintptr
	addr_QColumnView_createWindowContainer3 uintptr
	addr_QColumnView_updateMicroFocus2 uintptr
	addr_QColumnView_create2 uintptr
	addr_QColumnView_create3 uintptr
	addr_QColumnView_create4 uintptr
	addr_QColumnView_destroy2 uintptr
	addr_QColumnView_destroy3 uintptr
	addr_QColumnView_connect3 uintptr
	addr_QColumnView_connect4 uintptr
	addr_QColumnView_metaObject_default uintptr
	addr_QColumnView_indexAt_default uintptr
	addr_QColumnView_scrollTo_default uintptr
	addr_QColumnView_sizeHint_default uintptr
	addr_QColumnView_visualRect_default uintptr
	addr_QColumnView_setModel_default uintptr
	addr_QColumnView_setSelectionModel_default uintptr
	addr_QColumnView_setRootIndex_default uintptr
	addr_QColumnView_selectAll_default uintptr
	addr_QColumnView_isIndexHidden_default uintptr
	addr_QColumnView_moveCursor_default uintptr
	addr_QColumnView_resizeEvent_default uintptr
	addr_QColumnView_setSelection_default uintptr
	addr_QColumnView_visualRegionForSelection_default uintptr
	addr_QColumnView_horizontalOffset_default uintptr
	addr_QColumnView_verticalOffset_default uintptr
	addr_QColumnView_rowsInserted_default uintptr
	addr_QColumnView_currentChanged_default uintptr
	addr_QColumnView_scrollContentsBy_default uintptr
	addr_QColumnView_createColumn_default uintptr
	addr_QColumnView_keyboardSearch_default uintptr
	addr_QColumnView_sizeHintForRow_default uintptr
	addr_QColumnView_sizeHintForColumn_default uintptr
	addr_QColumnView_itemDelegateForIndex_default uintptr
	addr_QColumnView_inputMethodQuery_default uintptr
	addr_QColumnView_reset_default uintptr
	addr_QColumnView_doItemsLayout_default uintptr
	addr_QColumnView_rowsAboutToBeRemoved_default uintptr
	addr_QColumnView_selectionChanged_default uintptr
	addr_QColumnView_updateEditorData_default uintptr
	addr_QColumnView_updateEditorGeometries_default uintptr
	addr_QColumnView_updateGeometries_default uintptr
	addr_QColumnView_verticalScrollbarAction_default uintptr
	addr_QColumnView_horizontalScrollbarAction_default uintptr
	addr_QColumnView_verticalScrollbarValueChanged_default uintptr
	addr_QColumnView_horizontalScrollbarValueChanged_default uintptr
	addr_QColumnView_closeEditor_default uintptr
	addr_QColumnView_commitData_default uintptr
	addr_QColumnView_editorDestroyed_default uintptr
	addr_QColumnView_selectedIndexes_default uintptr
	addr_QColumnView_edit_default uintptr
	addr_QColumnView_selectionCommand_default uintptr
	addr_QColumnView_startDrag_default uintptr
	addr_QColumnView_initViewItemOption_default uintptr
	addr_QColumnView_focusNextPrevChild_default uintptr
	addr_QColumnView_event_default uintptr
	addr_QColumnView_viewportEvent_default uintptr
	addr_QColumnView_mousePressEvent_default uintptr
	addr_QColumnView_mouseMoveEvent_default uintptr
	addr_QColumnView_mouseReleaseEvent_default uintptr
	addr_QColumnView_mouseDoubleClickEvent_default uintptr
	addr_QColumnView_dragEnterEvent_default uintptr
	addr_QColumnView_dragMoveEvent_default uintptr
	addr_QColumnView_dragLeaveEvent_default uintptr
	addr_QColumnView_dropEvent_default uintptr
	addr_QColumnView_focusInEvent_default uintptr
	addr_QColumnView_focusOutEvent_default uintptr
	addr_QColumnView_keyPressEvent_default uintptr
	addr_QColumnView_timerEvent_default uintptr
	addr_QColumnView_inputMethodEvent_default uintptr
	addr_QColumnView_eventFilter_default uintptr
	addr_QColumnView_viewportSizeHint_default uintptr
	addr_QColumnView_minimumSizeHint_default uintptr
	addr_QColumnView_setupViewport_default uintptr
	addr_QColumnView_paintEvent_default uintptr
	addr_QColumnView_wheelEvent_default uintptr
	addr_QColumnView_contextMenuEvent_default uintptr
	addr_QColumnView_changeEvent_default uintptr
	addr_QColumnView_initStyleOption_default uintptr
	addr_QColumnView_devType_default uintptr
	addr_QColumnView_setVisible_default uintptr
	addr_QColumnView_heightForWidth_default uintptr
	addr_QColumnView_hasHeightForWidth_default uintptr
	addr_QColumnView_paintEngine_default uintptr
	addr_QColumnView_keyReleaseEvent_default uintptr
	addr_QColumnView_enterEvent_default uintptr
	addr_QColumnView_leaveEvent_default uintptr
	addr_QColumnView_moveEvent_default uintptr
	addr_QColumnView_closeEvent_default uintptr
	addr_QColumnView_tabletEvent_default uintptr
	addr_QColumnView_actionEvent_default uintptr
	addr_QColumnView_showEvent_default uintptr
	addr_QColumnView_hideEvent_default uintptr
	addr_QColumnView_metric_default uintptr
	addr_QColumnView_initPainter_default uintptr
	addr_QColumnView_redirected_default uintptr
	addr_QColumnView_sharedPainter_default uintptr
	addr_QColumnView_childEvent_default uintptr
	addr_QColumnView_customEvent_default uintptr
	addr_QColumnView_connectNotify_default uintptr
	addr_QColumnView_disconnectNotify_default uintptr
	pg_QColumnView_connect_updatePreviewWidget func(uintptr, int64)
	pg_QColumnView_connect_pressed func(uintptr, int64)
	pg_QColumnView_connect_clicked func(uintptr, int64)
	pg_QColumnView_connect_doubleClicked func(uintptr, int64)
	pg_QColumnView_connect_activated func(uintptr, int64)
	pg_QColumnView_connect_entered func(uintptr, int64)
	pg_QColumnView_connect_viewportEntered func(uintptr, int64)
	pg_QColumnView_connect_iconSizeChanged func(uintptr, int64)
	pg_QColumnView_connect_windowTitleChanged func(uintptr, int64)
	pg_QColumnView_connect_windowIconChanged func(uintptr, int64)
	pg_QColumnView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QColumnView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QColumnView_connect_destroyed func(uintptr, int64)
	pg_QColumnView_connect_destroyed2 func(uintptr, int64)
	pg_QColumnView_connect_objectNameChanged func(uintptr, int64)
	pg_QColumnView__disconnect func(int64)
	pg_QColumnView__set_override func(uintptr, int32)
	pg_QColumnView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColumnView_chain, "QColumnView_chain")
	symaddr(&addr_QColumnView_new, "QColumnView_new")
	symaddr(&addr_QColumnView_new2, "QColumnView_new2")
	symaddr(&addr_QColumnView_tr, "QColumnView_tr")
	symaddr(&addr_QColumnView_updatePreviewWidget, "QColumnView_updatePreviewWidget")
	symaddr(&addr_QColumnView_setResizeGripsVisible, "QColumnView_setResizeGripsVisible")
	symaddr(&addr_QColumnView_resizeGripsVisible, "QColumnView_resizeGripsVisible")
	symaddr(&addr_QColumnView_setPreviewColumnVisible, "QColumnView_setPreviewColumnVisible")
	symaddr(&addr_QColumnView_isPreviewColumnVisible, "QColumnView_isPreviewColumnVisible")
	symaddr(&addr_QColumnView_previewWidget, "QColumnView_previewWidget")
	symaddr(&addr_QColumnView_setPreviewWidget, "QColumnView_setPreviewWidget")
	symaddr(&addr_QColumnView_setColumnWidths, "QColumnView_setColumnWidths")
	symaddr(&addr_QColumnView_columnWidths, "QColumnView_columnWidths")
	symaddr(&addr_QColumnView_initializeColumn, "QColumnView_initializeColumn")
	symaddr(&addr_QColumnView_state, "QColumnView_state")
	symaddr(&addr_QColumnView_scheduleDelayedItemsLayout, "QColumnView_scheduleDelayedItemsLayout")
	symaddr(&addr_QColumnView_executeDelayedItemsLayout, "QColumnView_executeDelayedItemsLayout")
	symaddr(&addr_QColumnView_setDirtyRegion, "QColumnView_setDirtyRegion")
	symaddr(&addr_QColumnView_scrollDirtyRegion, "QColumnView_scrollDirtyRegion")
	symaddr(&addr_QColumnView_dirtyRegionOffset, "QColumnView_dirtyRegionOffset")
	symaddr(&addr_QColumnView_startAutoScroll, "QColumnView_startAutoScroll")
	symaddr(&addr_QColumnView_stopAutoScroll, "QColumnView_stopAutoScroll")
	symaddr(&addr_QColumnView_doAutoScroll, "QColumnView_doAutoScroll")
	symaddr(&addr_QColumnView_dropIndicatorPosition, "QColumnView_dropIndicatorPosition")
	symaddr(&addr_QColumnView_setViewportMargins, "QColumnView_setViewportMargins")
	symaddr(&addr_QColumnView_setViewportMargins2, "QColumnView_setViewportMargins2")
	symaddr(&addr_QColumnView_viewportMargins, "QColumnView_viewportMargins")
	symaddr(&addr_QColumnView_drawFrame, "QColumnView_drawFrame")
	symaddr(&addr_QColumnView_setTabOrder, "QColumnView_setTabOrder")
	symaddr(&addr_QColumnView_mouseGrabber, "QColumnView_mouseGrabber")
	symaddr(&addr_QColumnView_keyboardGrabber, "QColumnView_keyboardGrabber")
	symaddr(&addr_QColumnView_find, "QColumnView_find")
	symaddr(&addr_QColumnView_createWindowContainer, "QColumnView_createWindowContainer")
	symaddr(&addr_QColumnView_updateMicroFocus, "QColumnView_updateMicroFocus")
	symaddr(&addr_QColumnView_create, "QColumnView_create")
	symaddr(&addr_QColumnView_destroy, "QColumnView_destroy")
	symaddr(&addr_QColumnView_focusNextChild, "QColumnView_focusNextChild")
	symaddr(&addr_QColumnView_focusPreviousChild, "QColumnView_focusPreviousChild")
	symaddr(&addr_QColumnView_connect, "QColumnView_connect")
	symaddr(&addr_QColumnView_connect2, "QColumnView_connect2")
	symaddr(&addr_QColumnView_sender, "QColumnView_sender")
	symaddr(&addr_QColumnView_senderSignalIndex, "QColumnView_senderSignalIndex")
	symaddr(&addr_QColumnView_receivers, "QColumnView_receivers")
	symaddr(&addr_QColumnView_isSignalConnected, "QColumnView_isSignalConnected")
	reg(&pg_QColumnView_devicePixelRatioFScale, "QColumnView_devicePixelRatioFScale")
	reg(&pg_QColumnView_encodeMetricF, "QColumnView_encodeMetricF")
	reg(&pg_QColumnView_getDecodedMetricF, "QColumnView_getDecodedMetricF")
	symaddr(&addr_QColumnView_tr2, "QColumnView_tr2")
	symaddr(&addr_QColumnView_tr3, "QColumnView_tr3")
	symaddr(&addr_QColumnView_createWindowContainer2, "QColumnView_createWindowContainer2")
	symaddr(&addr_QColumnView_createWindowContainer3, "QColumnView_createWindowContainer3")
	symaddr(&addr_QColumnView_updateMicroFocus2, "QColumnView_updateMicroFocus2")
	symaddr(&addr_QColumnView_create2, "QColumnView_create2")
	symaddr(&addr_QColumnView_create3, "QColumnView_create3")
	symaddr(&addr_QColumnView_create4, "QColumnView_create4")
	symaddr(&addr_QColumnView_destroy2, "QColumnView_destroy2")
	symaddr(&addr_QColumnView_destroy3, "QColumnView_destroy3")
	symaddr(&addr_QColumnView_connect3, "QColumnView_connect3")
	symaddr(&addr_QColumnView_connect4, "QColumnView_connect4")
	symaddr(&addr_QColumnView_metaObject_default, "QColumnView_metaObject_default")
	symaddr(&addr_QColumnView_indexAt_default, "QColumnView_indexAt_default")
	symaddr(&addr_QColumnView_scrollTo_default, "QColumnView_scrollTo_default")
	symaddr(&addr_QColumnView_sizeHint_default, "QColumnView_sizeHint_default")
	symaddr(&addr_QColumnView_visualRect_default, "QColumnView_visualRect_default")
	symaddr(&addr_QColumnView_setModel_default, "QColumnView_setModel_default")
	symaddr(&addr_QColumnView_setSelectionModel_default, "QColumnView_setSelectionModel_default")
	symaddr(&addr_QColumnView_setRootIndex_default, "QColumnView_setRootIndex_default")
	symaddr(&addr_QColumnView_selectAll_default, "QColumnView_selectAll_default")
	symaddr(&addr_QColumnView_isIndexHidden_default, "QColumnView_isIndexHidden_default")
	symaddr(&addr_QColumnView_moveCursor_default, "QColumnView_moveCursor_default")
	symaddr(&addr_QColumnView_resizeEvent_default, "QColumnView_resizeEvent_default")
	symaddr(&addr_QColumnView_setSelection_default, "QColumnView_setSelection_default")
	symaddr(&addr_QColumnView_visualRegionForSelection_default, "QColumnView_visualRegionForSelection_default")
	symaddr(&addr_QColumnView_horizontalOffset_default, "QColumnView_horizontalOffset_default")
	symaddr(&addr_QColumnView_verticalOffset_default, "QColumnView_verticalOffset_default")
	symaddr(&addr_QColumnView_rowsInserted_default, "QColumnView_rowsInserted_default")
	symaddr(&addr_QColumnView_currentChanged_default, "QColumnView_currentChanged_default")
	symaddr(&addr_QColumnView_scrollContentsBy_default, "QColumnView_scrollContentsBy_default")
	symaddr(&addr_QColumnView_createColumn_default, "QColumnView_createColumn_default")
	symaddr(&addr_QColumnView_keyboardSearch_default, "QColumnView_keyboardSearch_default")
	symaddr(&addr_QColumnView_sizeHintForRow_default, "QColumnView_sizeHintForRow_default")
	symaddr(&addr_QColumnView_sizeHintForColumn_default, "QColumnView_sizeHintForColumn_default")
	symaddr(&addr_QColumnView_itemDelegateForIndex_default, "QColumnView_itemDelegateForIndex_default")
	symaddr(&addr_QColumnView_inputMethodQuery_default, "QColumnView_inputMethodQuery_default")
	symaddr(&addr_QColumnView_reset_default, "QColumnView_reset_default")
	symaddr(&addr_QColumnView_doItemsLayout_default, "QColumnView_doItemsLayout_default")
	symaddr(&addr_QColumnView_rowsAboutToBeRemoved_default, "QColumnView_rowsAboutToBeRemoved_default")
	symaddr(&addr_QColumnView_selectionChanged_default, "QColumnView_selectionChanged_default")
	symaddr(&addr_QColumnView_updateEditorData_default, "QColumnView_updateEditorData_default")
	symaddr(&addr_QColumnView_updateEditorGeometries_default, "QColumnView_updateEditorGeometries_default")
	symaddr(&addr_QColumnView_updateGeometries_default, "QColumnView_updateGeometries_default")
	symaddr(&addr_QColumnView_verticalScrollbarAction_default, "QColumnView_verticalScrollbarAction_default")
	symaddr(&addr_QColumnView_horizontalScrollbarAction_default, "QColumnView_horizontalScrollbarAction_default")
	symaddr(&addr_QColumnView_verticalScrollbarValueChanged_default, "QColumnView_verticalScrollbarValueChanged_default")
	symaddr(&addr_QColumnView_horizontalScrollbarValueChanged_default, "QColumnView_horizontalScrollbarValueChanged_default")
	symaddr(&addr_QColumnView_closeEditor_default, "QColumnView_closeEditor_default")
	symaddr(&addr_QColumnView_commitData_default, "QColumnView_commitData_default")
	symaddr(&addr_QColumnView_editorDestroyed_default, "QColumnView_editorDestroyed_default")
	symaddr(&addr_QColumnView_selectedIndexes_default, "QColumnView_selectedIndexes_default")
	symaddr(&addr_QColumnView_edit_default, "QColumnView_edit_default")
	symaddr(&addr_QColumnView_selectionCommand_default, "QColumnView_selectionCommand_default")
	symaddr(&addr_QColumnView_startDrag_default, "QColumnView_startDrag_default")
	symaddr(&addr_QColumnView_initViewItemOption_default, "QColumnView_initViewItemOption_default")
	symaddr(&addr_QColumnView_focusNextPrevChild_default, "QColumnView_focusNextPrevChild_default")
	symaddr(&addr_QColumnView_event_default, "QColumnView_event_default")
	symaddr(&addr_QColumnView_viewportEvent_default, "QColumnView_viewportEvent_default")
	symaddr(&addr_QColumnView_mousePressEvent_default, "QColumnView_mousePressEvent_default")
	symaddr(&addr_QColumnView_mouseMoveEvent_default, "QColumnView_mouseMoveEvent_default")
	symaddr(&addr_QColumnView_mouseReleaseEvent_default, "QColumnView_mouseReleaseEvent_default")
	symaddr(&addr_QColumnView_mouseDoubleClickEvent_default, "QColumnView_mouseDoubleClickEvent_default")
	symaddr(&addr_QColumnView_dragEnterEvent_default, "QColumnView_dragEnterEvent_default")
	symaddr(&addr_QColumnView_dragMoveEvent_default, "QColumnView_dragMoveEvent_default")
	symaddr(&addr_QColumnView_dragLeaveEvent_default, "QColumnView_dragLeaveEvent_default")
	symaddr(&addr_QColumnView_dropEvent_default, "QColumnView_dropEvent_default")
	symaddr(&addr_QColumnView_focusInEvent_default, "QColumnView_focusInEvent_default")
	symaddr(&addr_QColumnView_focusOutEvent_default, "QColumnView_focusOutEvent_default")
	symaddr(&addr_QColumnView_keyPressEvent_default, "QColumnView_keyPressEvent_default")
	symaddr(&addr_QColumnView_timerEvent_default, "QColumnView_timerEvent_default")
	symaddr(&addr_QColumnView_inputMethodEvent_default, "QColumnView_inputMethodEvent_default")
	symaddr(&addr_QColumnView_eventFilter_default, "QColumnView_eventFilter_default")
	symaddr(&addr_QColumnView_viewportSizeHint_default, "QColumnView_viewportSizeHint_default")
	symaddr(&addr_QColumnView_minimumSizeHint_default, "QColumnView_minimumSizeHint_default")
	symaddr(&addr_QColumnView_setupViewport_default, "QColumnView_setupViewport_default")
	symaddr(&addr_QColumnView_paintEvent_default, "QColumnView_paintEvent_default")
	symaddr(&addr_QColumnView_wheelEvent_default, "QColumnView_wheelEvent_default")
	symaddr(&addr_QColumnView_contextMenuEvent_default, "QColumnView_contextMenuEvent_default")
	symaddr(&addr_QColumnView_changeEvent_default, "QColumnView_changeEvent_default")
	symaddr(&addr_QColumnView_initStyleOption_default, "QColumnView_initStyleOption_default")
	symaddr(&addr_QColumnView_devType_default, "QColumnView_devType_default")
	symaddr(&addr_QColumnView_setVisible_default, "QColumnView_setVisible_default")
	symaddr(&addr_QColumnView_heightForWidth_default, "QColumnView_heightForWidth_default")
	symaddr(&addr_QColumnView_hasHeightForWidth_default, "QColumnView_hasHeightForWidth_default")
	symaddr(&addr_QColumnView_paintEngine_default, "QColumnView_paintEngine_default")
	symaddr(&addr_QColumnView_keyReleaseEvent_default, "QColumnView_keyReleaseEvent_default")
	symaddr(&addr_QColumnView_enterEvent_default, "QColumnView_enterEvent_default")
	symaddr(&addr_QColumnView_leaveEvent_default, "QColumnView_leaveEvent_default")
	symaddr(&addr_QColumnView_moveEvent_default, "QColumnView_moveEvent_default")
	symaddr(&addr_QColumnView_closeEvent_default, "QColumnView_closeEvent_default")
	symaddr(&addr_QColumnView_tabletEvent_default, "QColumnView_tabletEvent_default")
	symaddr(&addr_QColumnView_actionEvent_default, "QColumnView_actionEvent_default")
	symaddr(&addr_QColumnView_showEvent_default, "QColumnView_showEvent_default")
	symaddr(&addr_QColumnView_hideEvent_default, "QColumnView_hideEvent_default")
	symaddr(&addr_QColumnView_metric_default, "QColumnView_metric_default")
	symaddr(&addr_QColumnView_initPainter_default, "QColumnView_initPainter_default")
	symaddr(&addr_QColumnView_redirected_default, "QColumnView_redirected_default")
	symaddr(&addr_QColumnView_sharedPainter_default, "QColumnView_sharedPainter_default")
	symaddr(&addr_QColumnView_childEvent_default, "QColumnView_childEvent_default")
	symaddr(&addr_QColumnView_customEvent_default, "QColumnView_customEvent_default")
	symaddr(&addr_QColumnView_connectNotify_default, "QColumnView_connectNotify_default")
	symaddr(&addr_QColumnView_disconnectNotify_default, "QColumnView_disconnectNotify_default")
	reg(&pg_QColumnView_connect_updatePreviewWidget, "QColumnView_connect_updatePreviewWidget")
	reg(&pg_QColumnView_connect_pressed, "QColumnView_connect_pressed")
	reg(&pg_QColumnView_connect_clicked, "QColumnView_connect_clicked")
	reg(&pg_QColumnView_connect_doubleClicked, "QColumnView_connect_doubleClicked")
	reg(&pg_QColumnView_connect_activated, "QColumnView_connect_activated")
	reg(&pg_QColumnView_connect_entered, "QColumnView_connect_entered")
	reg(&pg_QColumnView_connect_viewportEntered, "QColumnView_connect_viewportEntered")
	reg(&pg_QColumnView_connect_iconSizeChanged, "QColumnView_connect_iconSizeChanged")
	reg(&pg_QColumnView_connect_windowTitleChanged, "QColumnView_connect_windowTitleChanged")
	reg(&pg_QColumnView_connect_windowIconChanged, "QColumnView_connect_windowIconChanged")
	reg(&pg_QColumnView_connect_windowIconTextChanged, "QColumnView_connect_windowIconTextChanged")
	reg(&pg_QColumnView_connect_customContextMenuRequested, "QColumnView_connect_customContextMenuRequested")
	reg(&pg_QColumnView_connect_destroyed, "QColumnView_connect_destroyed")
	reg(&pg_QColumnView_connect_destroyed2, "QColumnView_connect_destroyed2")
	reg(&pg_QColumnView_connect_objectNameChanged, "QColumnView_connect_objectNameChanged")
	reg(&pg_QColumnView__disconnect, "QColumnView__disconnect")
	reg(&pg_QColumnView__set_override, "QColumnView__set_override")
	reg(&pg_QColumnView_delete, "QColumnView_delete")
}

type QColumnView struct {
	h uintptr
	*QAbstractItemView
}

func bindQColumnView(h uintptr) *QColumnView {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QColumnView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QColumnView{h: h, QAbstractItemView: &QAbstractItemView{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQColumnView(h uintptr) *QColumnView { return bindQColumnView(h) }

func (this *QColumnView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColumnView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColumnView(h unsafe.Pointer) *QColumnView { return bindQColumnView(uintptr(h)) }
func UnsafeNewQColumnViewOwned(h unsafe.Pointer) *QColumnView { return newownedQColumnView(uintptr(h)) }

func (this *QColumnView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemView != nil {
		this.QAbstractItemView.UnsafeRelease()
	}
}

func NewQColumnView(parent *QWidget) *QColumnView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColumnView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQColumnView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColumnView_virtIdx, QColumnView_enable)
	return __o
}

func NewQColumnView2() *QColumnView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColumnView_new2, uintptr(__goid))
	__o := newownedQColumnView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColumnView_virtIdx, QColumnView_enable)
	return __o
}

func QColumnView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColumnView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QColumnView) UpdatePreviewWidget(index *QModelIndex) {
	purego.SyscallN(addr_QColumnView_updatePreviewWidget, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QColumnView) SetResizeGripsVisible(visible bool) {
	purego.SyscallN(addr_QColumnView_setResizeGripsVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) ResizeGripsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_resizeGripsVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColumnView) SetPreviewColumnVisible(visible bool) {
	purego.SyscallN(addr_QColumnView_setPreviewColumnVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) IsPreviewColumnVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_isPreviewColumnVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColumnView) PreviewWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_previewWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QColumnView) SetPreviewWidget(widget *QWidget) {
	purego.SyscallN(addr_QColumnView_setPreviewWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QColumnView) SetColumnWidths(list []int) {
	list_arr := make([]int32, len(list))
	for __i, __e := range list {
		list_arr[__i] = int32(__e)
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	purego.SyscallN(addr_QColumnView_setColumnWidths, this.h, list_p, uintptr(len(list)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(list_arr)
}

func (this *QColumnView) ColumnWidths() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColumnView_columnWidths, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QColumnView) InitializeColumn(column *QAbstractItemView) {
	purego.SyscallN(addr_QColumnView_initializeColumn, this.h, uintptr(column.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(column)
}

func (this *QColumnView) State() QAbstractItemView__State {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__State(__r)
}

func (this *QColumnView) ScheduleDelayedItemsLayout() {
	purego.SyscallN(addr_QColumnView_scheduleDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) ExecuteDelayedItemsLayout() {
	purego.SyscallN(addr_QColumnView_executeDelayedItemsLayout, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) SetDirtyRegion(region *QRegion) {
	purego.SyscallN(addr_QColumnView_setDirtyRegion, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QColumnView) ScrollDirtyRegion(dx int, dy int) {
	purego.SyscallN(addr_QColumnView_scrollDirtyRegion, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QColumnView) DirtyRegionOffset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QColumnView_dirtyRegionOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QColumnView) StartAutoScroll() {
	purego.SyscallN(addr_QColumnView_startAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) StopAutoScroll() {
	purego.SyscallN(addr_QColumnView_stopAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) DoAutoScroll() {
	purego.SyscallN(addr_QColumnView_doAutoScroll, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) DropIndicatorPosition() QAbstractItemView__DropIndicatorPosition {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_dropIndicatorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractItemView__DropIndicatorPosition(__r)
}

func (this *QColumnView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QColumnView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QColumnView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QColumnView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QColumnView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QColumnView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QColumnView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QColumnView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QColumnView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QColumnView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QColumnView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_mouseGrabber)
	return bindQWidget(__r)
}

func QColumnView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_keyboardGrabber)
	return bindQWidget(__r)
}

func QColumnView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QColumnView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QColumnView) UpdateMicroFocus() {
	purego.SyscallN(addr_QColumnView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) Create() {
	purego.SyscallN(addr_QColumnView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) Destroy() {
	purego.SyscallN(addr_QColumnView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColumnView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QColumnView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QColumnView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QColumnView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QColumnView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QColumnView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QColumnView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QColumnView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QColumnView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QColumnView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QColumnView_DevicePixelRatioFScale() float64 {
	__r := pg_QColumnView_devicePixelRatioFScale()
	return __r
}

func QColumnView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QColumnView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QColumnView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QColumnView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QColumnView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColumnView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QColumnView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColumnView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QColumnView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QColumnView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColumnView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QColumnView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QColumnView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QColumnView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QColumnView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QColumnView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QColumnView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QColumnView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QColumnView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QColumnView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QColumnView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QColumnView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QColumnView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QColumnView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QColumnView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QColumnView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QColumnView) IndexAtDefault(point *QPoint) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QColumnView_indexAt_default, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQModelIndex(__r)
}

func (this *QColumnView) ScrollToDefault(index *QModelIndex, hint QAbstractItemView__ScrollHint) {
	purego.SyscallN(addr_QColumnView_scrollTo_default, this.h, uintptr(index.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QColumnView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QColumnView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QColumnView) VisualRectDefault(index *QModelIndex) *QRect {
	__r, _, _ := purego.SyscallN(addr_QColumnView_visualRect_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QColumnView) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QColumnView_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QColumnView) SetSelectionModelDefault(selectionModel *QItemSelectionModel) {
	purego.SyscallN(addr_QColumnView_setSelectionModel_default, this.h, uintptr(selectionModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectionModel)
}

func (this *QColumnView) SetRootIndexDefault(index *QModelIndex) {
	purego.SyscallN(addr_QColumnView_setRootIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QColumnView) SelectAllDefault() {
	purego.SyscallN(addr_QColumnView_selectAll_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) IsIndexHiddenDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_isIndexHidden_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QColumnView) MoveCursorDefault(cursorAction QAbstractItemView__CursorAction, modifiers KeyboardModifier) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QColumnView_moveCursor_default, this.h, uintptr(cursorAction), uintptr(modifiers))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QColumnView) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QColumnView_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) SetSelectionDefault(rect *QRect, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QColumnView_setSelection_default, this.h, uintptr(rect.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QColumnView) VisualRegionForSelectionDefault(selection *QItemSelection) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QColumnView_visualRegionForSelection_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQRegion(__r)
}

func (this *QColumnView) HorizontalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_horizontalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) VerticalOffsetDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_verticalOffset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) RowsInsertedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QColumnView_rowsInserted_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QColumnView) CurrentChangedDefault(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QColumnView_currentChanged_default, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QColumnView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QColumnView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QColumnView) CreateColumnDefault(rootIndex *QModelIndex) *QAbstractItemView {
	__r, _, _ := purego.SyscallN(addr_QColumnView_createColumn_default, this.h, uintptr(rootIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rootIndex)
	return bindQAbstractItemView(__r)
}

func (this *QColumnView) KeyboardSearchDefault(search string) {
	var search_p uintptr
	if len(search) > 0 {
		search_p = uintptr(unsafe.Pointer(unsafe.StringData(search)))
	}
	purego.SyscallN(addr_QColumnView_keyboardSearch_default, this.h, search_p, uintptr(len(search)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(search)
}

func (this *QColumnView) SizeHintForRowDefault(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_sizeHintForRow_default, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) SizeHintForColumnDefault(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_sizeHintForColumn_default, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) ItemDelegateForIndexDefault(index *QModelIndex) *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QColumnView_itemDelegateForIndex_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQAbstractItemDelegate(__r)
}

func (this *QColumnView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QColumnView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QColumnView) ResetDefault() {
	purego.SyscallN(addr_QColumnView_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) DoItemsLayoutDefault() {
	purego.SyscallN(addr_QColumnView_doItemsLayout_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) RowsAboutToBeRemovedDefault(parent *QModelIndex, start int, end int) {
	purego.SyscallN(addr_QColumnView_rowsAboutToBeRemoved_default, this.h, uintptr(parent.UnsafePointer()), uintptr(start), uintptr(end))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QColumnView) SelectionChangedDefault(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QColumnView_selectionChanged_default, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QColumnView) UpdateEditorDataDefault() {
	purego.SyscallN(addr_QColumnView_updateEditorData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) UpdateEditorGeometriesDefault() {
	purego.SyscallN(addr_QColumnView_updateEditorGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) UpdateGeometriesDefault() {
	purego.SyscallN(addr_QColumnView_updateGeometries_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColumnView) VerticalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QColumnView_verticalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QColumnView) HorizontalScrollbarActionDefault(action int) {
	purego.SyscallN(addr_QColumnView_horizontalScrollbarAction_default, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QColumnView) VerticalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QColumnView_verticalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QColumnView) HorizontalScrollbarValueChangedDefault(value int) {
	purego.SyscallN(addr_QColumnView_horizontalScrollbarValueChanged_default, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QColumnView) CloseEditorDefault(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QColumnView_closeEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QColumnView) CommitDataDefault(editor *QWidget) {
	purego.SyscallN(addr_QColumnView_commitData_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QColumnView) EditorDestroyedDefault(editor *QObject) {
	purego.SyscallN(addr_QColumnView_editorDestroyed_default, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QColumnView) SelectedIndexesDefault() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColumnView_selectedIndexes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QColumnView) EditDefault(index *QModelIndex, trigger QAbstractItemView__EditTrigger, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_edit_default, this.h, uintptr(index.UnsafePointer()), uintptr(trigger), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QColumnView) SelectionCommandDefault(index *QModelIndex, event *QEvent) QItemSelectionModel__SelectionFlag {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_selectionCommand_default, this.h, uintptr(index.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(event)
	return QItemSelectionModel__SelectionFlag(__r)
}

func (this *QColumnView) StartDragDefault(supportedActions DropAction) {
	purego.SyscallN(addr_QColumnView_startDrag_default, this.h, uintptr(supportedActions))
	runtime.KeepAlive(this)
}

func (this *QColumnView) InitViewItemOptionDefault(option *QStyleOptionViewItem) {
	purego.SyscallN(addr_QColumnView_initViewItemOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QColumnView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColumnView) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QColumnView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QColumnView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColumnView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColumnView_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColumnView_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColumnView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QColumnView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QColumnView_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QColumnView_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QColumnView_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QColumnView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QColumnView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QColumnView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QColumnView_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QColumnView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QColumnView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QColumnView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QColumnView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QColumnView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QColumnView) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QColumnView_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QColumnView) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QColumnView_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColumnView) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QColumnView_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColumnView) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QColumnView_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColumnView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QColumnView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColumnView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QColumnView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QColumnView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QColumnView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QColumnView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColumnView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QColumnView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QColumnView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QColumnView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QColumnView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QColumnView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QColumnView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QColumnView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QColumnView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QColumnView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QColumnView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QColumnView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QColumnView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColumnView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QColumnView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QColumnView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QColumnView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QColumnView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QColumnView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QColumnView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QColumnView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QColumnView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColumnView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QColumnView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QColumnView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QColumnView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QColumnView) OnUpdatePreviewWidget(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_updatePreviewWidget(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QColumnView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QColumnView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QColumnView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QColumnView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QColumnView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QColumnView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QColumnView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QColumnView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QColumnView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QColumnView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColumnView) Disconnect(handle int64) {
	pg_QColumnView__disconnect(handle)
	unregisterCallback(handle)
}

func QColumnView_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -7252823847285716802:
		return 1
	case -4832340528568269947:
		return 2
	case 729690261457684101:
		return 3
	case -1267696851865372997:
		return 4
	case 4283922507359547360:
		return 5
	case -4689322347236366780:
		return 6
	case -1574077111192992763:
		return 7
	case 8758113809152574300:
		return 8
	case 6057989667089378859:
		return 9
	case -158523882194220890:
		return 10
	case -4302917285655014661:
		return 11
	case -1539266838375037305:
		return 12
	case 7063795720810552074:
		return 13
	case -1500310024801917802:
		return 14
	case 5218707810677921016:
		return 15
	case 4472434241583495620:
		return 16
	case -6978583972112316034:
		return 17
	case -6118001150126564745:
		return 18
	case -5176237958185448161:
		return 19
	case -5041452885022114200:
		return 20
	case -4084029023708556876:
		return 21
	case -3666792239577137638:
		return 22
	case -3811438228839363824:
		return 23
	case 7345853097340709638:
		return 24
	case 1086335023529244512:
		return 25
	case -1888629738259754754:
		return 26
	case -8151049665126058157:
		return 27
	case -5892205462145366901:
		return 28
	case 9029941737538200621:
		return 29
	case -8726300139121449691:
		return 30
	case 1022771470898460576:
		return 31
	case -1155555122391250623:
		return 32
	case -785551405282638793:
		return 33
	case 6366953527806415524:
		return 34
	case 1786430177580376614:
		return 35
	case 8648484264628816806:
		return 36
	case -6021383853323633948:
		return 37
	case 6979418566323745651:
		return 38
	case 740869982805721908:
		return 39
	case 4206689019978932209:
		return 40
	case -2447490284072398302:
		return 41
	case -8112366604639385491:
		return 42
	case -7450787277202094556:
		return 43
	case 5898328517505784191:
		return 44
	case 1237236456312824991:
		return 45
	case -1971465653061861519:
		return 46
	case -332090898078543319:
		return 47
	case -4255494004099806581:
		return 48
	case -4415920049297795213:
		return 49
	case -4029073905697023887:
		return 50
	case 206122715919917589:
		return 51
	case -1781545137352745978:
		return 52
	case -3139904158373725006:
		return 53
	case 9107058682754591716:
		return 54
	case 975043767033552552:
		return 55
	case -385596010576048993:
		return 56
	case 1023791977679793185:
		return 57
	case -5409927581812374498:
		return 58
	case 5438819661942843222:
		return 59
	case 5217623514630858297:
		return 60
	case 503579067252048239:
		return 61
	case 7049264599837532491:
		return 62
	case 4531996381736730688:
		return 63
	case 5290442163167458465:
		return 64
	case 1919988340934511316:
		return 65
	case -6698043222227006015:
		return 66
	case 4959376063826165943:
		return 67
	case -488693769771308531:
		return 68
	case -8659679893868524044:
		return 69
	case 2110182719375245575:
		return 70
	case 4311534060495737339:
		return 71
	case -8266762397995423915:
		return 72
	case 6398007404756346381:
		return 73
	case 772598475950050299:
		return 74
	case 6316270346668237895:
		return 75
	case -551070009623711640:
		return 76
	case -9161949150489525396:
		return 77
	case -1989937050019464783:
		return 78
	case -212068866860701141:
		return 79
	case 1697444140088555581:
		return 80
	case -6825369174751916:
		return 81
	case -199718707876719129:
		return 82
	case -2039166586246443706:
		return 83
	case 427629649267311777:
		return 84
	case 6477500296980050210:
		return 85
	case 4127240633310655048:
		return 86
	case 5792771328741774991:
		return 87
	case -1185315367526857833:
		return 88
	case 3769868229123868474:
		return 89
	case 7379029982170503008:
		return 90
	case -568369772982215962:
		return 91
	}
	return -1
}

func QColumnView_enable(h uintptr, idx int32) { pg_QColumnView__set_override(h, idx) }

func (this *QColumnView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QColumnView) OnIndexAt(fn func(*QPoint) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -7252823847285716802, func(point uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(point))).UnsafePointer()) }) // slot "indexAt"
}

func (this *QColumnView) OnScrollTo(fn func(*QModelIndex, QAbstractItemView__ScrollHint)) {
	registerVirt(this.UnsafeGoID(), -4832340528568269947, func(index uintptr, hint int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__ScrollHint(hint)) }) // slot "scrollTo"
}

func (this *QColumnView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QColumnView) OnVisualRect(fn func(*QModelIndex) *QRect) {
	registerVirt(this.UnsafeGoID(), -1267696851865372997, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "visualRect"
}

func (this *QColumnView) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QColumnView) OnSetSelectionModel(fn func(*QItemSelectionModel)) {
	registerVirt(this.UnsafeGoID(), -4689322347236366780, func(selectionModel uintptr) { fn(UnsafeNewQItemSelectionModel(unsafe.Pointer(selectionModel))) }) // slot "setSelectionModel"
}

func (this *QColumnView) OnSetRootIndex(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1574077111192992763, func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setRootIndex"
}

func (this *QColumnView) OnSelectAll(fn func()) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() { fn() }) // slot "selectAll"
}

func (this *QColumnView) OnIsIndexHidden(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 6057989667089378859, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "isIndexHidden"
}

func (this *QColumnView) OnMoveCursor(fn func(QAbstractItemView__CursorAction, KeyboardModifier) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -158523882194220890, func(cursorAction int32, modifiers int32) uintptr { return uintptr(fn(QAbstractItemView__CursorAction(cursorAction), KeyboardModifier(modifiers)).UnsafePointer()) }) // slot "moveCursor"
}

func (this *QColumnView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QColumnView) OnSetSelection(fn func(*QRect, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(rect uintptr, command int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setSelection"
}

func (this *QColumnView) OnVisualRegionForSelection(fn func(*QItemSelection) *QRegion) {
	registerVirt(this.UnsafeGoID(), 7063795720810552074, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "visualRegionForSelection"
}

func (this *QColumnView) OnHorizontalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1500310024801917802, func() int32 { return int32(fn()) }) // slot "horizontalOffset"
}

func (this *QColumnView) OnVerticalOffset(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5218707810677921016, func() int32 { return int32(fn()) }) // slot "verticalOffset"
}

func (this *QColumnView) OnRowsInserted(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), 4472434241583495620, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsInserted"
}

func (this *QColumnView) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -6978583972112316034, func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) }) // slot "currentChanged"
}

func (this *QColumnView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QColumnView) OnCreateColumn(fn func(*QModelIndex) *QAbstractItemView) {
	registerVirt(this.UnsafeGoID(), -5176237958185448161, func(rootIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(rootIndex))).UnsafePointer()) }) // slot "createColumn"
}

func (this *QColumnView) OnKeyboardSearch(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -5041452885022114200, func(search uintptr, search_len int64) { fn(pgStrNoFree(search, search_len)) }) // slot "keyboardSearch"
}

func (this *QColumnView) OnSizeHintForRow(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -4084029023708556876, func(row int32) int32 { return int32(fn(int(row))) }) // slot "sizeHintForRow"
}

func (this *QColumnView) OnSizeHintForColumn(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), -3666792239577137638, func(column int32) int32 { return int32(fn(int(column))) }) // slot "sizeHintForColumn"
}

func (this *QColumnView) OnItemDelegateForIndex(fn func(*QModelIndex) *QAbstractItemDelegate) {
	registerVirt(this.UnsafeGoID(), -3811438228839363824, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "itemDelegateForIndex"
}

func (this *QColumnView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QColumnView) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QColumnView) OnDoItemsLayout(fn func()) {
	registerVirt(this.UnsafeGoID(), -1888629738259754754, func() { fn() }) // slot "doItemsLayout"
}

func (this *QColumnView) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) {
	registerVirt(this.UnsafeGoID(), -8151049665126058157, func(parent uintptr, start int32, end int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(start), int(end)) }) // slot "rowsAboutToBeRemoved"
}

func (this *QColumnView) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) {
	registerVirt(this.UnsafeGoID(), -5892205462145366901, func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) }) // slot "selectionChanged"
}

func (this *QColumnView) OnUpdateEditorData(fn func()) {
	registerVirt(this.UnsafeGoID(), 9029941737538200621, func() { fn() }) // slot "updateEditorData"
}

func (this *QColumnView) OnUpdateEditorGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), -8726300139121449691, func() { fn() }) // slot "updateEditorGeometries"
}

func (this *QColumnView) OnUpdateGeometries(fn func()) {
	registerVirt(this.UnsafeGoID(), 1022771470898460576, func() { fn() }) // slot "updateGeometries"
}

func (this *QColumnView) OnVerticalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -1155555122391250623, func(action int32) { fn(int(action)) }) // slot "verticalScrollbarAction"
}

func (this *QColumnView) OnHorizontalScrollbarAction(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -785551405282638793, func(action int32) { fn(int(action)) }) // slot "horizontalScrollbarAction"
}

func (this *QColumnView) OnVerticalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 6366953527806415524, func(value int32) { fn(int(value)) }) // slot "verticalScrollbarValueChanged"
}

func (this *QColumnView) OnHorizontalScrollbarValueChanged(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1786430177580376614, func(value int32) { fn(int(value)) }) // slot "horizontalScrollbarValueChanged"
}

func (this *QColumnView) OnCloseEditor(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) {
	registerVirt(this.UnsafeGoID(), 8648484264628816806, func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) }) // slot "closeEditor"
}

func (this *QColumnView) OnCommitData(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6021383853323633948, func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) }) // slot "commitData"
}

func (this *QColumnView) OnEditorDestroyed(fn func(*QObject)) {
	registerVirt(this.UnsafeGoID(), 6979418566323745651, func(editor uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(editor))) }) // slot "editorDestroyed"
}

func (this *QColumnView) OnSelectedIndexes(fn func() []*QModelIndex) {
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

func (this *QColumnView) OnEdit(fn func(*QModelIndex, QAbstractItemView__EditTrigger, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 4206689019978932209, func(index uintptr, trigger int32, event uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QAbstractItemView__EditTrigger(trigger), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "edit"
}

func (this *QColumnView) OnSelectionCommand(fn func(*QModelIndex, *QEvent) QItemSelectionModel__SelectionFlag) {
	registerVirt(this.UnsafeGoID(), -2447490284072398302, func(index uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "selectionCommand"
}

func (this *QColumnView) OnStartDrag(fn func(DropAction)) {
	registerVirt(this.UnsafeGoID(), -8112366604639385491, func(supportedActions int32) { fn(DropAction(supportedActions)) }) // slot "startDrag"
}

func (this *QColumnView) OnInitViewItemOption(fn func(*QStyleOptionViewItem)) {
	registerVirt(this.UnsafeGoID(), -7450787277202094556, func(option uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option))) }) // slot "initViewItemOption"
}

func (this *QColumnView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QColumnView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QColumnView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QColumnView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QColumnView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QColumnView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QColumnView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QColumnView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QColumnView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QColumnView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QColumnView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QColumnView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QColumnView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QColumnView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QColumnView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QColumnView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QColumnView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QColumnView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QColumnView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QColumnView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QColumnView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QColumnView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QColumnView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QColumnView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QColumnView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QColumnView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QColumnView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QColumnView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QColumnView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QColumnView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QColumnView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QColumnView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QColumnView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QColumnView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QColumnView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QColumnView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QColumnView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QColumnView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QColumnView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QColumnView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QColumnView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QColumnView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QColumnView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QColumnView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QColumnView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QColumnView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QColumnView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QColumnView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QColumnView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColumnView_delete(this.h)
	}
	this.UnsafeRelease()
}
