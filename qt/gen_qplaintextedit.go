package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPlainTextEdit_chain uintptr
	addr_QPlainTextEdit_new uintptr
	addr_QPlainTextEdit_new2 uintptr
	addr_QPlainTextEdit_new3 uintptr
	addr_QPlainTextEdit_new4 uintptr
	addr_QPlainTextEdit_tr uintptr
	addr_QPlainTextEdit_setDocument uintptr
	addr_QPlainTextEdit_document uintptr
	addr_QPlainTextEdit_setPlaceholderText uintptr
	addr_QPlainTextEdit_placeholderText uintptr
	addr_QPlainTextEdit_setTextCursor uintptr
	addr_QPlainTextEdit_textCursor uintptr
	addr_QPlainTextEdit_isReadOnly uintptr
	addr_QPlainTextEdit_setReadOnly uintptr
	addr_QPlainTextEdit_setTextInteractionFlags uintptr
	addr_QPlainTextEdit_textInteractionFlags uintptr
	addr_QPlainTextEdit_mergeCurrentCharFormat uintptr
	addr_QPlainTextEdit_setCurrentCharFormat uintptr
	addr_QPlainTextEdit_currentCharFormat uintptr
	addr_QPlainTextEdit_tabChangesFocus uintptr
	addr_QPlainTextEdit_setTabChangesFocus uintptr
	addr_QPlainTextEdit_setDocumentTitle uintptr
	addr_QPlainTextEdit_documentTitle uintptr
	addr_QPlainTextEdit_isUndoRedoEnabled uintptr
	addr_QPlainTextEdit_setUndoRedoEnabled uintptr
	addr_QPlainTextEdit_setMaximumBlockCount uintptr
	addr_QPlainTextEdit_maximumBlockCount uintptr
	addr_QPlainTextEdit_lineWrapMode uintptr
	addr_QPlainTextEdit_setLineWrapMode uintptr
	addr_QPlainTextEdit_wordWrapMode uintptr
	addr_QPlainTextEdit_setWordWrapMode uintptr
	addr_QPlainTextEdit_setBackgroundVisible uintptr
	addr_QPlainTextEdit_backgroundVisible uintptr
	addr_QPlainTextEdit_setCenterOnScroll uintptr
	addr_QPlainTextEdit_centerOnScroll uintptr
	addr_QPlainTextEdit_find uintptr
	addr_QPlainTextEdit_find2 uintptr
	addr_QPlainTextEdit_toPlainText uintptr
	addr_QPlainTextEdit_ensureCursorVisible uintptr
	addr_QPlainTextEdit_loadResource uintptr
	addr_QPlainTextEdit_createStandardContextMenu uintptr
	addr_QPlainTextEdit_createStandardContextMenu2 uintptr
	addr_QPlainTextEdit_cursorForPosition uintptr
	addr_QPlainTextEdit_cursorRect uintptr
	addr_QPlainTextEdit_cursorRect2 uintptr
	addr_QPlainTextEdit_anchorAt uintptr
	addr_QPlainTextEdit_overwriteMode uintptr
	addr_QPlainTextEdit_setOverwriteMode uintptr
	pg_QPlainTextEdit_tabStopDistance func(uintptr) float64
	pg_QPlainTextEdit_setTabStopDistance func(uintptr, float64)
	addr_QPlainTextEdit_cursorWidth uintptr
	addr_QPlainTextEdit_setCursorWidth uintptr
	addr_QPlainTextEdit_setExtraSelections uintptr
	addr_QPlainTextEdit_extraSelections uintptr
	addr_QPlainTextEdit_moveCursor uintptr
	addr_QPlainTextEdit_canPaste uintptr
	addr_QPlainTextEdit_print uintptr
	addr_QPlainTextEdit_blockCount uintptr
	addr_QPlainTextEdit_inputMethodQuery uintptr
	addr_QPlainTextEdit_inputMethodQuery2 uintptr
	addr_QPlainTextEdit_setPlainText uintptr
	addr_QPlainTextEdit_cut uintptr
	addr_QPlainTextEdit_copy uintptr
	addr_QPlainTextEdit_paste uintptr
	addr_QPlainTextEdit_undo uintptr
	addr_QPlainTextEdit_redo uintptr
	addr_QPlainTextEdit_clear uintptr
	addr_QPlainTextEdit_selectAll uintptr
	addr_QPlainTextEdit_insertPlainText uintptr
	addr_QPlainTextEdit_appendPlainText uintptr
	addr_QPlainTextEdit_appendHtml uintptr
	addr_QPlainTextEdit_centerCursor uintptr
	addr_QPlainTextEdit_zoomIn uintptr
	addr_QPlainTextEdit_zoomOut uintptr
	addr_QPlainTextEdit_textChanged uintptr
	addr_QPlainTextEdit_undoAvailable uintptr
	addr_QPlainTextEdit_redoAvailable uintptr
	addr_QPlainTextEdit_copyAvailable uintptr
	addr_QPlainTextEdit_selectionChanged uintptr
	addr_QPlainTextEdit_cursorPositionChanged uintptr
	addr_QPlainTextEdit_updateRequest uintptr
	addr_QPlainTextEdit_blockCountChanged uintptr
	addr_QPlainTextEdit_modificationChanged uintptr
	addr_QPlainTextEdit_firstVisibleBlock uintptr
	addr_QPlainTextEdit_contentOffset uintptr
	addr_QPlainTextEdit_blockBoundingRect uintptr
	addr_QPlainTextEdit_blockBoundingGeometry uintptr
	addr_QPlainTextEdit_getPaintContext uintptr
	pg_QPlainTextEdit_zoomInF func(uintptr, float32)
	addr_QPlainTextEdit_setViewportMargins uintptr
	addr_QPlainTextEdit_setViewportMargins2 uintptr
	addr_QPlainTextEdit_viewportMargins uintptr
	addr_QPlainTextEdit_drawFrame uintptr
	addr_QPlainTextEdit_setTabOrder uintptr
	addr_QPlainTextEdit_mouseGrabber uintptr
	addr_QPlainTextEdit_keyboardGrabber uintptr
	addr_QPlainTextEdit_createWindowContainer uintptr
	addr_QPlainTextEdit_updateMicroFocus uintptr
	addr_QPlainTextEdit_create uintptr
	addr_QPlainTextEdit_destroy uintptr
	addr_QPlainTextEdit_focusNextChild uintptr
	addr_QPlainTextEdit_focusPreviousChild uintptr
	addr_QPlainTextEdit_connect uintptr
	addr_QPlainTextEdit_connect2 uintptr
	addr_QPlainTextEdit_sender uintptr
	addr_QPlainTextEdit_senderSignalIndex uintptr
	addr_QPlainTextEdit_receivers uintptr
	addr_QPlainTextEdit_isSignalConnected uintptr
	pg_QPlainTextEdit_devicePixelRatioFScale func() float64
	pg_QPlainTextEdit_encodeMetricF func(int32, float64) int32
	pg_QPlainTextEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QPlainTextEdit_tr2 uintptr
	addr_QPlainTextEdit_tr3 uintptr
	addr_QPlainTextEdit_find3 uintptr
	addr_QPlainTextEdit_find4 uintptr
	addr_QPlainTextEdit_moveCursor2 uintptr
	addr_QPlainTextEdit_zoomIn2 uintptr
	addr_QPlainTextEdit_zoomOut2 uintptr
	addr_QPlainTextEdit_createWindowContainer2 uintptr
	addr_QPlainTextEdit_createWindowContainer3 uintptr
	addr_QPlainTextEdit_updateMicroFocus2 uintptr
	addr_QPlainTextEdit_create2 uintptr
	addr_QPlainTextEdit_create3 uintptr
	addr_QPlainTextEdit_create4 uintptr
	addr_QPlainTextEdit_destroy2 uintptr
	addr_QPlainTextEdit_destroy3 uintptr
	addr_QPlainTextEdit_connect3 uintptr
	addr_QPlainTextEdit_connect4 uintptr
	addr_QPlainTextEdit_loadResource_default uintptr
	addr_QPlainTextEdit_inputMethodQuery_default uintptr
	addr_QPlainTextEdit_inputMethodQuery_default2 uintptr
	addr_QPlainTextEdit_metaObject_default uintptr
	addr_QPlainTextEdit_event_default uintptr
	addr_QPlainTextEdit_timerEvent_default uintptr
	addr_QPlainTextEdit_keyPressEvent_default uintptr
	addr_QPlainTextEdit_keyReleaseEvent_default uintptr
	addr_QPlainTextEdit_resizeEvent_default uintptr
	addr_QPlainTextEdit_paintEvent_default uintptr
	addr_QPlainTextEdit_mousePressEvent_default uintptr
	addr_QPlainTextEdit_mouseMoveEvent_default uintptr
	addr_QPlainTextEdit_mouseReleaseEvent_default uintptr
	addr_QPlainTextEdit_mouseDoubleClickEvent_default uintptr
	addr_QPlainTextEdit_focusNextPrevChild_default uintptr
	addr_QPlainTextEdit_contextMenuEvent_default uintptr
	addr_QPlainTextEdit_dragEnterEvent_default uintptr
	addr_QPlainTextEdit_dragLeaveEvent_default uintptr
	addr_QPlainTextEdit_dragMoveEvent_default uintptr
	addr_QPlainTextEdit_dropEvent_default uintptr
	addr_QPlainTextEdit_focusInEvent_default uintptr
	addr_QPlainTextEdit_focusOutEvent_default uintptr
	addr_QPlainTextEdit_showEvent_default uintptr
	addr_QPlainTextEdit_changeEvent_default uintptr
	addr_QPlainTextEdit_wheelEvent_default uintptr
	addr_QPlainTextEdit_createMimeDataFromSelection_default uintptr
	addr_QPlainTextEdit_canInsertFromMimeData_default uintptr
	addr_QPlainTextEdit_insertFromMimeData_default uintptr
	addr_QPlainTextEdit_inputMethodEvent_default uintptr
	addr_QPlainTextEdit_scrollContentsBy_default uintptr
	addr_QPlainTextEdit_doSetTextCursor_default uintptr
	addr_QPlainTextEdit_minimumSizeHint_default uintptr
	addr_QPlainTextEdit_sizeHint_default uintptr
	addr_QPlainTextEdit_setupViewport_default uintptr
	addr_QPlainTextEdit_eventFilter_default uintptr
	addr_QPlainTextEdit_viewportEvent_default uintptr
	addr_QPlainTextEdit_viewportSizeHint_default uintptr
	addr_QPlainTextEdit_initStyleOption_default uintptr
	addr_QPlainTextEdit_devType_default uintptr
	addr_QPlainTextEdit_setVisible_default uintptr
	addr_QPlainTextEdit_heightForWidth_default uintptr
	addr_QPlainTextEdit_hasHeightForWidth_default uintptr
	addr_QPlainTextEdit_paintEngine_default uintptr
	addr_QPlainTextEdit_enterEvent_default uintptr
	addr_QPlainTextEdit_leaveEvent_default uintptr
	addr_QPlainTextEdit_moveEvent_default uintptr
	addr_QPlainTextEdit_closeEvent_default uintptr
	addr_QPlainTextEdit_tabletEvent_default uintptr
	addr_QPlainTextEdit_actionEvent_default uintptr
	addr_QPlainTextEdit_hideEvent_default uintptr
	addr_QPlainTextEdit_metric_default uintptr
	addr_QPlainTextEdit_initPainter_default uintptr
	addr_QPlainTextEdit_redirected_default uintptr
	addr_QPlainTextEdit_sharedPainter_default uintptr
	addr_QPlainTextEdit_childEvent_default uintptr
	addr_QPlainTextEdit_customEvent_default uintptr
	addr_QPlainTextEdit_connectNotify_default uintptr
	addr_QPlainTextEdit_disconnectNotify_default uintptr
	pg_QPlainTextEdit_connect_textChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_undoAvailable func(uintptr, int64)
	pg_QPlainTextEdit_connect_redoAvailable func(uintptr, int64)
	pg_QPlainTextEdit_connect_copyAvailable func(uintptr, int64)
	pg_QPlainTextEdit_connect_selectionChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_cursorPositionChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_updateRequest func(uintptr, int64)
	pg_QPlainTextEdit_connect_blockCountChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_modificationChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QPlainTextEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QPlainTextEdit_connect_destroyed func(uintptr, int64)
	pg_QPlainTextEdit_connect_destroyed2 func(uintptr, int64)
	pg_QPlainTextEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QPlainTextEdit__disconnect func(int64)
	pg_QPlainTextEdit__set_override func(uintptr, int32)
	pg_QPlainTextEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPlainTextEdit_chain, "QPlainTextEdit_chain")
	symaddr(&addr_QPlainTextEdit_new, "QPlainTextEdit_new")
	symaddr(&addr_QPlainTextEdit_new2, "QPlainTextEdit_new2")
	symaddr(&addr_QPlainTextEdit_new3, "QPlainTextEdit_new3")
	symaddr(&addr_QPlainTextEdit_new4, "QPlainTextEdit_new4")
	symaddr(&addr_QPlainTextEdit_tr, "QPlainTextEdit_tr")
	symaddr(&addr_QPlainTextEdit_setDocument, "QPlainTextEdit_setDocument")
	symaddr(&addr_QPlainTextEdit_document, "QPlainTextEdit_document")
	symaddr(&addr_QPlainTextEdit_setPlaceholderText, "QPlainTextEdit_setPlaceholderText")
	symaddr(&addr_QPlainTextEdit_placeholderText, "QPlainTextEdit_placeholderText")
	symaddr(&addr_QPlainTextEdit_setTextCursor, "QPlainTextEdit_setTextCursor")
	symaddr(&addr_QPlainTextEdit_textCursor, "QPlainTextEdit_textCursor")
	symaddr(&addr_QPlainTextEdit_isReadOnly, "QPlainTextEdit_isReadOnly")
	symaddr(&addr_QPlainTextEdit_setReadOnly, "QPlainTextEdit_setReadOnly")
	symaddr(&addr_QPlainTextEdit_setTextInteractionFlags, "QPlainTextEdit_setTextInteractionFlags")
	symaddr(&addr_QPlainTextEdit_textInteractionFlags, "QPlainTextEdit_textInteractionFlags")
	symaddr(&addr_QPlainTextEdit_mergeCurrentCharFormat, "QPlainTextEdit_mergeCurrentCharFormat")
	symaddr(&addr_QPlainTextEdit_setCurrentCharFormat, "QPlainTextEdit_setCurrentCharFormat")
	symaddr(&addr_QPlainTextEdit_currentCharFormat, "QPlainTextEdit_currentCharFormat")
	symaddr(&addr_QPlainTextEdit_tabChangesFocus, "QPlainTextEdit_tabChangesFocus")
	symaddr(&addr_QPlainTextEdit_setTabChangesFocus, "QPlainTextEdit_setTabChangesFocus")
	symaddr(&addr_QPlainTextEdit_setDocumentTitle, "QPlainTextEdit_setDocumentTitle")
	symaddr(&addr_QPlainTextEdit_documentTitle, "QPlainTextEdit_documentTitle")
	symaddr(&addr_QPlainTextEdit_isUndoRedoEnabled, "QPlainTextEdit_isUndoRedoEnabled")
	symaddr(&addr_QPlainTextEdit_setUndoRedoEnabled, "QPlainTextEdit_setUndoRedoEnabled")
	symaddr(&addr_QPlainTextEdit_setMaximumBlockCount, "QPlainTextEdit_setMaximumBlockCount")
	symaddr(&addr_QPlainTextEdit_maximumBlockCount, "QPlainTextEdit_maximumBlockCount")
	symaddr(&addr_QPlainTextEdit_lineWrapMode, "QPlainTextEdit_lineWrapMode")
	symaddr(&addr_QPlainTextEdit_setLineWrapMode, "QPlainTextEdit_setLineWrapMode")
	symaddr(&addr_QPlainTextEdit_wordWrapMode, "QPlainTextEdit_wordWrapMode")
	symaddr(&addr_QPlainTextEdit_setWordWrapMode, "QPlainTextEdit_setWordWrapMode")
	symaddr(&addr_QPlainTextEdit_setBackgroundVisible, "QPlainTextEdit_setBackgroundVisible")
	symaddr(&addr_QPlainTextEdit_backgroundVisible, "QPlainTextEdit_backgroundVisible")
	symaddr(&addr_QPlainTextEdit_setCenterOnScroll, "QPlainTextEdit_setCenterOnScroll")
	symaddr(&addr_QPlainTextEdit_centerOnScroll, "QPlainTextEdit_centerOnScroll")
	symaddr(&addr_QPlainTextEdit_find, "QPlainTextEdit_find")
	symaddr(&addr_QPlainTextEdit_find2, "QPlainTextEdit_find2")
	symaddr(&addr_QPlainTextEdit_toPlainText, "QPlainTextEdit_toPlainText")
	symaddr(&addr_QPlainTextEdit_ensureCursorVisible, "QPlainTextEdit_ensureCursorVisible")
	symaddr(&addr_QPlainTextEdit_loadResource, "QPlainTextEdit_loadResource")
	symaddr(&addr_QPlainTextEdit_createStandardContextMenu, "QPlainTextEdit_createStandardContextMenu")
	symaddr(&addr_QPlainTextEdit_createStandardContextMenu2, "QPlainTextEdit_createStandardContextMenu2")
	symaddr(&addr_QPlainTextEdit_cursorForPosition, "QPlainTextEdit_cursorForPosition")
	symaddr(&addr_QPlainTextEdit_cursorRect, "QPlainTextEdit_cursorRect")
	symaddr(&addr_QPlainTextEdit_cursorRect2, "QPlainTextEdit_cursorRect2")
	symaddr(&addr_QPlainTextEdit_anchorAt, "QPlainTextEdit_anchorAt")
	symaddr(&addr_QPlainTextEdit_overwriteMode, "QPlainTextEdit_overwriteMode")
	symaddr(&addr_QPlainTextEdit_setOverwriteMode, "QPlainTextEdit_setOverwriteMode")
	reg(&pg_QPlainTextEdit_tabStopDistance, "QPlainTextEdit_tabStopDistance")
	reg(&pg_QPlainTextEdit_setTabStopDistance, "QPlainTextEdit_setTabStopDistance")
	symaddr(&addr_QPlainTextEdit_cursorWidth, "QPlainTextEdit_cursorWidth")
	symaddr(&addr_QPlainTextEdit_setCursorWidth, "QPlainTextEdit_setCursorWidth")
	symaddr(&addr_QPlainTextEdit_setExtraSelections, "QPlainTextEdit_setExtraSelections")
	symaddr(&addr_QPlainTextEdit_extraSelections, "QPlainTextEdit_extraSelections")
	symaddr(&addr_QPlainTextEdit_moveCursor, "QPlainTextEdit_moveCursor")
	symaddr(&addr_QPlainTextEdit_canPaste, "QPlainTextEdit_canPaste")
	symaddr(&addr_QPlainTextEdit_print, "QPlainTextEdit_print")
	symaddr(&addr_QPlainTextEdit_blockCount, "QPlainTextEdit_blockCount")
	symaddr(&addr_QPlainTextEdit_inputMethodQuery, "QPlainTextEdit_inputMethodQuery")
	symaddr(&addr_QPlainTextEdit_inputMethodQuery2, "QPlainTextEdit_inputMethodQuery2")
	symaddr(&addr_QPlainTextEdit_setPlainText, "QPlainTextEdit_setPlainText")
	symaddr(&addr_QPlainTextEdit_cut, "QPlainTextEdit_cut")
	symaddr(&addr_QPlainTextEdit_copy, "QPlainTextEdit_copy")
	symaddr(&addr_QPlainTextEdit_paste, "QPlainTextEdit_paste")
	symaddr(&addr_QPlainTextEdit_undo, "QPlainTextEdit_undo")
	symaddr(&addr_QPlainTextEdit_redo, "QPlainTextEdit_redo")
	symaddr(&addr_QPlainTextEdit_clear, "QPlainTextEdit_clear")
	symaddr(&addr_QPlainTextEdit_selectAll, "QPlainTextEdit_selectAll")
	symaddr(&addr_QPlainTextEdit_insertPlainText, "QPlainTextEdit_insertPlainText")
	symaddr(&addr_QPlainTextEdit_appendPlainText, "QPlainTextEdit_appendPlainText")
	symaddr(&addr_QPlainTextEdit_appendHtml, "QPlainTextEdit_appendHtml")
	symaddr(&addr_QPlainTextEdit_centerCursor, "QPlainTextEdit_centerCursor")
	symaddr(&addr_QPlainTextEdit_zoomIn, "QPlainTextEdit_zoomIn")
	symaddr(&addr_QPlainTextEdit_zoomOut, "QPlainTextEdit_zoomOut")
	symaddr(&addr_QPlainTextEdit_textChanged, "QPlainTextEdit_textChanged")
	symaddr(&addr_QPlainTextEdit_undoAvailable, "QPlainTextEdit_undoAvailable")
	symaddr(&addr_QPlainTextEdit_redoAvailable, "QPlainTextEdit_redoAvailable")
	symaddr(&addr_QPlainTextEdit_copyAvailable, "QPlainTextEdit_copyAvailable")
	symaddr(&addr_QPlainTextEdit_selectionChanged, "QPlainTextEdit_selectionChanged")
	symaddr(&addr_QPlainTextEdit_cursorPositionChanged, "QPlainTextEdit_cursorPositionChanged")
	symaddr(&addr_QPlainTextEdit_updateRequest, "QPlainTextEdit_updateRequest")
	symaddr(&addr_QPlainTextEdit_blockCountChanged, "QPlainTextEdit_blockCountChanged")
	symaddr(&addr_QPlainTextEdit_modificationChanged, "QPlainTextEdit_modificationChanged")
	symaddr(&addr_QPlainTextEdit_firstVisibleBlock, "QPlainTextEdit_firstVisibleBlock")
	symaddr(&addr_QPlainTextEdit_contentOffset, "QPlainTextEdit_contentOffset")
	symaddr(&addr_QPlainTextEdit_blockBoundingRect, "QPlainTextEdit_blockBoundingRect")
	symaddr(&addr_QPlainTextEdit_blockBoundingGeometry, "QPlainTextEdit_blockBoundingGeometry")
	symaddr(&addr_QPlainTextEdit_getPaintContext, "QPlainTextEdit_getPaintContext")
	reg(&pg_QPlainTextEdit_zoomInF, "QPlainTextEdit_zoomInF")
	symaddr(&addr_QPlainTextEdit_setViewportMargins, "QPlainTextEdit_setViewportMargins")
	symaddr(&addr_QPlainTextEdit_setViewportMargins2, "QPlainTextEdit_setViewportMargins2")
	symaddr(&addr_QPlainTextEdit_viewportMargins, "QPlainTextEdit_viewportMargins")
	symaddr(&addr_QPlainTextEdit_drawFrame, "QPlainTextEdit_drawFrame")
	symaddr(&addr_QPlainTextEdit_setTabOrder, "QPlainTextEdit_setTabOrder")
	symaddr(&addr_QPlainTextEdit_mouseGrabber, "QPlainTextEdit_mouseGrabber")
	symaddr(&addr_QPlainTextEdit_keyboardGrabber, "QPlainTextEdit_keyboardGrabber")
	symaddr(&addr_QPlainTextEdit_createWindowContainer, "QPlainTextEdit_createWindowContainer")
	symaddr(&addr_QPlainTextEdit_updateMicroFocus, "QPlainTextEdit_updateMicroFocus")
	symaddr(&addr_QPlainTextEdit_create, "QPlainTextEdit_create")
	symaddr(&addr_QPlainTextEdit_destroy, "QPlainTextEdit_destroy")
	symaddr(&addr_QPlainTextEdit_focusNextChild, "QPlainTextEdit_focusNextChild")
	symaddr(&addr_QPlainTextEdit_focusPreviousChild, "QPlainTextEdit_focusPreviousChild")
	symaddr(&addr_QPlainTextEdit_connect, "QPlainTextEdit_connect")
	symaddr(&addr_QPlainTextEdit_connect2, "QPlainTextEdit_connect2")
	symaddr(&addr_QPlainTextEdit_sender, "QPlainTextEdit_sender")
	symaddr(&addr_QPlainTextEdit_senderSignalIndex, "QPlainTextEdit_senderSignalIndex")
	symaddr(&addr_QPlainTextEdit_receivers, "QPlainTextEdit_receivers")
	symaddr(&addr_QPlainTextEdit_isSignalConnected, "QPlainTextEdit_isSignalConnected")
	reg(&pg_QPlainTextEdit_devicePixelRatioFScale, "QPlainTextEdit_devicePixelRatioFScale")
	reg(&pg_QPlainTextEdit_encodeMetricF, "QPlainTextEdit_encodeMetricF")
	reg(&pg_QPlainTextEdit_getDecodedMetricF, "QPlainTextEdit_getDecodedMetricF")
	symaddr(&addr_QPlainTextEdit_tr2, "QPlainTextEdit_tr2")
	symaddr(&addr_QPlainTextEdit_tr3, "QPlainTextEdit_tr3")
	symaddr(&addr_QPlainTextEdit_find3, "QPlainTextEdit_find3")
	symaddr(&addr_QPlainTextEdit_find4, "QPlainTextEdit_find4")
	symaddr(&addr_QPlainTextEdit_moveCursor2, "QPlainTextEdit_moveCursor2")
	symaddr(&addr_QPlainTextEdit_zoomIn2, "QPlainTextEdit_zoomIn2")
	symaddr(&addr_QPlainTextEdit_zoomOut2, "QPlainTextEdit_zoomOut2")
	symaddr(&addr_QPlainTextEdit_createWindowContainer2, "QPlainTextEdit_createWindowContainer2")
	symaddr(&addr_QPlainTextEdit_createWindowContainer3, "QPlainTextEdit_createWindowContainer3")
	symaddr(&addr_QPlainTextEdit_updateMicroFocus2, "QPlainTextEdit_updateMicroFocus2")
	symaddr(&addr_QPlainTextEdit_create2, "QPlainTextEdit_create2")
	symaddr(&addr_QPlainTextEdit_create3, "QPlainTextEdit_create3")
	symaddr(&addr_QPlainTextEdit_create4, "QPlainTextEdit_create4")
	symaddr(&addr_QPlainTextEdit_destroy2, "QPlainTextEdit_destroy2")
	symaddr(&addr_QPlainTextEdit_destroy3, "QPlainTextEdit_destroy3")
	symaddr(&addr_QPlainTextEdit_connect3, "QPlainTextEdit_connect3")
	symaddr(&addr_QPlainTextEdit_connect4, "QPlainTextEdit_connect4")
	symaddr(&addr_QPlainTextEdit_loadResource_default, "QPlainTextEdit_loadResource_default")
	symaddr(&addr_QPlainTextEdit_inputMethodQuery_default, "QPlainTextEdit_inputMethodQuery_default")
	symaddr(&addr_QPlainTextEdit_inputMethodQuery_default2, "QPlainTextEdit_inputMethodQuery_default2")
	symaddr(&addr_QPlainTextEdit_metaObject_default, "QPlainTextEdit_metaObject_default")
	symaddr(&addr_QPlainTextEdit_event_default, "QPlainTextEdit_event_default")
	symaddr(&addr_QPlainTextEdit_timerEvent_default, "QPlainTextEdit_timerEvent_default")
	symaddr(&addr_QPlainTextEdit_keyPressEvent_default, "QPlainTextEdit_keyPressEvent_default")
	symaddr(&addr_QPlainTextEdit_keyReleaseEvent_default, "QPlainTextEdit_keyReleaseEvent_default")
	symaddr(&addr_QPlainTextEdit_resizeEvent_default, "QPlainTextEdit_resizeEvent_default")
	symaddr(&addr_QPlainTextEdit_paintEvent_default, "QPlainTextEdit_paintEvent_default")
	symaddr(&addr_QPlainTextEdit_mousePressEvent_default, "QPlainTextEdit_mousePressEvent_default")
	symaddr(&addr_QPlainTextEdit_mouseMoveEvent_default, "QPlainTextEdit_mouseMoveEvent_default")
	symaddr(&addr_QPlainTextEdit_mouseReleaseEvent_default, "QPlainTextEdit_mouseReleaseEvent_default")
	symaddr(&addr_QPlainTextEdit_mouseDoubleClickEvent_default, "QPlainTextEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QPlainTextEdit_focusNextPrevChild_default, "QPlainTextEdit_focusNextPrevChild_default")
	symaddr(&addr_QPlainTextEdit_contextMenuEvent_default, "QPlainTextEdit_contextMenuEvent_default")
	symaddr(&addr_QPlainTextEdit_dragEnterEvent_default, "QPlainTextEdit_dragEnterEvent_default")
	symaddr(&addr_QPlainTextEdit_dragLeaveEvent_default, "QPlainTextEdit_dragLeaveEvent_default")
	symaddr(&addr_QPlainTextEdit_dragMoveEvent_default, "QPlainTextEdit_dragMoveEvent_default")
	symaddr(&addr_QPlainTextEdit_dropEvent_default, "QPlainTextEdit_dropEvent_default")
	symaddr(&addr_QPlainTextEdit_focusInEvent_default, "QPlainTextEdit_focusInEvent_default")
	symaddr(&addr_QPlainTextEdit_focusOutEvent_default, "QPlainTextEdit_focusOutEvent_default")
	symaddr(&addr_QPlainTextEdit_showEvent_default, "QPlainTextEdit_showEvent_default")
	symaddr(&addr_QPlainTextEdit_changeEvent_default, "QPlainTextEdit_changeEvent_default")
	symaddr(&addr_QPlainTextEdit_wheelEvent_default, "QPlainTextEdit_wheelEvent_default")
	symaddr(&addr_QPlainTextEdit_createMimeDataFromSelection_default, "QPlainTextEdit_createMimeDataFromSelection_default")
	symaddr(&addr_QPlainTextEdit_canInsertFromMimeData_default, "QPlainTextEdit_canInsertFromMimeData_default")
	symaddr(&addr_QPlainTextEdit_insertFromMimeData_default, "QPlainTextEdit_insertFromMimeData_default")
	symaddr(&addr_QPlainTextEdit_inputMethodEvent_default, "QPlainTextEdit_inputMethodEvent_default")
	symaddr(&addr_QPlainTextEdit_scrollContentsBy_default, "QPlainTextEdit_scrollContentsBy_default")
	symaddr(&addr_QPlainTextEdit_doSetTextCursor_default, "QPlainTextEdit_doSetTextCursor_default")
	symaddr(&addr_QPlainTextEdit_minimumSizeHint_default, "QPlainTextEdit_minimumSizeHint_default")
	symaddr(&addr_QPlainTextEdit_sizeHint_default, "QPlainTextEdit_sizeHint_default")
	symaddr(&addr_QPlainTextEdit_setupViewport_default, "QPlainTextEdit_setupViewport_default")
	symaddr(&addr_QPlainTextEdit_eventFilter_default, "QPlainTextEdit_eventFilter_default")
	symaddr(&addr_QPlainTextEdit_viewportEvent_default, "QPlainTextEdit_viewportEvent_default")
	symaddr(&addr_QPlainTextEdit_viewportSizeHint_default, "QPlainTextEdit_viewportSizeHint_default")
	symaddr(&addr_QPlainTextEdit_initStyleOption_default, "QPlainTextEdit_initStyleOption_default")
	symaddr(&addr_QPlainTextEdit_devType_default, "QPlainTextEdit_devType_default")
	symaddr(&addr_QPlainTextEdit_setVisible_default, "QPlainTextEdit_setVisible_default")
	symaddr(&addr_QPlainTextEdit_heightForWidth_default, "QPlainTextEdit_heightForWidth_default")
	symaddr(&addr_QPlainTextEdit_hasHeightForWidth_default, "QPlainTextEdit_hasHeightForWidth_default")
	symaddr(&addr_QPlainTextEdit_paintEngine_default, "QPlainTextEdit_paintEngine_default")
	symaddr(&addr_QPlainTextEdit_enterEvent_default, "QPlainTextEdit_enterEvent_default")
	symaddr(&addr_QPlainTextEdit_leaveEvent_default, "QPlainTextEdit_leaveEvent_default")
	symaddr(&addr_QPlainTextEdit_moveEvent_default, "QPlainTextEdit_moveEvent_default")
	symaddr(&addr_QPlainTextEdit_closeEvent_default, "QPlainTextEdit_closeEvent_default")
	symaddr(&addr_QPlainTextEdit_tabletEvent_default, "QPlainTextEdit_tabletEvent_default")
	symaddr(&addr_QPlainTextEdit_actionEvent_default, "QPlainTextEdit_actionEvent_default")
	symaddr(&addr_QPlainTextEdit_hideEvent_default, "QPlainTextEdit_hideEvent_default")
	symaddr(&addr_QPlainTextEdit_metric_default, "QPlainTextEdit_metric_default")
	symaddr(&addr_QPlainTextEdit_initPainter_default, "QPlainTextEdit_initPainter_default")
	symaddr(&addr_QPlainTextEdit_redirected_default, "QPlainTextEdit_redirected_default")
	symaddr(&addr_QPlainTextEdit_sharedPainter_default, "QPlainTextEdit_sharedPainter_default")
	symaddr(&addr_QPlainTextEdit_childEvent_default, "QPlainTextEdit_childEvent_default")
	symaddr(&addr_QPlainTextEdit_customEvent_default, "QPlainTextEdit_customEvent_default")
	symaddr(&addr_QPlainTextEdit_connectNotify_default, "QPlainTextEdit_connectNotify_default")
	symaddr(&addr_QPlainTextEdit_disconnectNotify_default, "QPlainTextEdit_disconnectNotify_default")
	reg(&pg_QPlainTextEdit_connect_textChanged, "QPlainTextEdit_connect_textChanged")
	reg(&pg_QPlainTextEdit_connect_undoAvailable, "QPlainTextEdit_connect_undoAvailable")
	reg(&pg_QPlainTextEdit_connect_redoAvailable, "QPlainTextEdit_connect_redoAvailable")
	reg(&pg_QPlainTextEdit_connect_copyAvailable, "QPlainTextEdit_connect_copyAvailable")
	reg(&pg_QPlainTextEdit_connect_selectionChanged, "QPlainTextEdit_connect_selectionChanged")
	reg(&pg_QPlainTextEdit_connect_cursorPositionChanged, "QPlainTextEdit_connect_cursorPositionChanged")
	reg(&pg_QPlainTextEdit_connect_updateRequest, "QPlainTextEdit_connect_updateRequest")
	reg(&pg_QPlainTextEdit_connect_blockCountChanged, "QPlainTextEdit_connect_blockCountChanged")
	reg(&pg_QPlainTextEdit_connect_modificationChanged, "QPlainTextEdit_connect_modificationChanged")
	reg(&pg_QPlainTextEdit_connect_windowTitleChanged, "QPlainTextEdit_connect_windowTitleChanged")
	reg(&pg_QPlainTextEdit_connect_windowIconChanged, "QPlainTextEdit_connect_windowIconChanged")
	reg(&pg_QPlainTextEdit_connect_windowIconTextChanged, "QPlainTextEdit_connect_windowIconTextChanged")
	reg(&pg_QPlainTextEdit_connect_customContextMenuRequested, "QPlainTextEdit_connect_customContextMenuRequested")
	reg(&pg_QPlainTextEdit_connect_destroyed, "QPlainTextEdit_connect_destroyed")
	reg(&pg_QPlainTextEdit_connect_destroyed2, "QPlainTextEdit_connect_destroyed2")
	reg(&pg_QPlainTextEdit_connect_objectNameChanged, "QPlainTextEdit_connect_objectNameChanged")
	reg(&pg_QPlainTextEdit__disconnect, "QPlainTextEdit__disconnect")
	reg(&pg_QPlainTextEdit__set_override, "QPlainTextEdit__set_override")
	reg(&pg_QPlainTextEdit_delete, "QPlainTextEdit_delete")
}

type QPlainTextEdit struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQPlainTextEdit(h uintptr) *QPlainTextEdit {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QPlainTextEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPlainTextEdit{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQPlainTextEdit(h uintptr) *QPlainTextEdit { return bindQPlainTextEdit(h) }

func (this *QPlainTextEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPlainTextEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPlainTextEdit(h unsafe.Pointer) *QPlainTextEdit { return bindQPlainTextEdit(uintptr(h)) }
func UnsafeNewQPlainTextEditOwned(h unsafe.Pointer) *QPlainTextEdit { return newownedQPlainTextEdit(uintptr(h)) }

func (this *QPlainTextEdit) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func NewQPlainTextEdit(parent *QWidget) *QPlainTextEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPlainTextEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPlainTextEdit_virtIdx, QPlainTextEdit_enable)
	return __o
}

func NewQPlainTextEdit2() *QPlainTextEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_new2, uintptr(__goid))
	__o := newownedQPlainTextEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPlainTextEdit_virtIdx, QPlainTextEdit_enable)
	return __o
}

func NewQPlainTextEdit3(text string) *QPlainTextEdit {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQPlainTextEdit(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPlainTextEdit_virtIdx, QPlainTextEdit_enable)
	return __o
}

func NewQPlainTextEdit4(text string, parent *QWidget) *QPlainTextEdit {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQPlainTextEdit(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPlainTextEdit_virtIdx, QPlainTextEdit_enable)
	return __o
}

func QPlainTextEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) SetDocument(document *QTextDocument) {
	purego.SyscallN(addr_QPlainTextEdit_setDocument, this.h, uintptr(document.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(document)
}

func (this *QPlainTextEdit) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QPlainTextEdit) SetPlaceholderText(placeholderText string) {
	var placeholderText_p uintptr
	if len(placeholderText) > 0 {
		placeholderText_p = uintptr(unsafe.Pointer(unsafe.StringData(placeholderText)))
	}
	purego.SyscallN(addr_QPlainTextEdit_setPlaceholderText, this.h, placeholderText_p, uintptr(len(placeholderText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(placeholderText)
}

func (this *QPlainTextEdit) PlaceholderText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_placeholderText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) SetTextCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QPlainTextEdit_setTextCursor, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QPlainTextEdit) TextCursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_textCursor, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QPlainTextEdit) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) SetReadOnly(ro bool) {
	purego.SyscallN(addr_QPlainTextEdit_setReadOnly, this.h, uintptr(b2i(ro)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetTextInteractionFlags(flags TextInteractionFlag) {
	purego.SyscallN(addr_QPlainTextEdit_setTextInteractionFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) TextInteractionFlags() TextInteractionFlag {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_textInteractionFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextInteractionFlag(__r)
}

func (this *QPlainTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat) {
	purego.SyscallN(addr_QPlainTextEdit_mergeCurrentCharFormat, this.h, uintptr(modifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modifier)
}

func (this *QPlainTextEdit) SetCurrentCharFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QPlainTextEdit_setCurrentCharFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QPlainTextEdit) CurrentCharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_currentCharFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QPlainTextEdit) TabChangesFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_tabChangesFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) SetTabChangesFocus(b bool) {
	purego.SyscallN(addr_QPlainTextEdit_setTabChangesFocus, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetDocumentTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QPlainTextEdit_setDocumentTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QPlainTextEdit) DocumentTitle() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_documentTitle, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) IsUndoRedoEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_isUndoRedoEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) SetUndoRedoEnabled(enable bool) {
	purego.SyscallN(addr_QPlainTextEdit_setUndoRedoEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetMaximumBlockCount(maximum int) {
	purego.SyscallN(addr_QPlainTextEdit_setMaximumBlockCount, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) MaximumBlockCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_maximumBlockCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) LineWrapMode() QPlainTextEdit__LineWrapMode {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_lineWrapMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPlainTextEdit__LineWrapMode(__r)
}

func (this *QPlainTextEdit) SetLineWrapMode(mode QPlainTextEdit__LineWrapMode) {
	purego.SyscallN(addr_QPlainTextEdit_setLineWrapMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) WordWrapMode() QTextOption__WrapMode {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_wordWrapMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextOption__WrapMode(__r)
}

func (this *QPlainTextEdit) SetWordWrapMode(policy QTextOption__WrapMode) {
	purego.SyscallN(addr_QPlainTextEdit_setWordWrapMode, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetBackgroundVisible(visible bool) {
	purego.SyscallN(addr_QPlainTextEdit_setBackgroundVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) BackgroundVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_backgroundVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) SetCenterOnScroll(enabled bool) {
	purego.SyscallN(addr_QPlainTextEdit_setCenterOnScroll, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CenterOnScroll() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_centerOnScroll, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) Find(exp string) bool {
	var exp_p uintptr
	if len(exp) > 0 {
		exp_p = uintptr(unsafe.Pointer(unsafe.StringData(exp)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_find, this.h, exp_p, uintptr(len(exp)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QPlainTextEdit) FindWithExp(exp *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_find2, this.h, uintptr(exp.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QPlainTextEdit) ToPlainText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_toPlainText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) EnsureCursorVisible() {
	purego.SyscallN(addr_QPlainTextEdit_ensureCursorVisible, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) LoadResource(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_loadResource, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) CreateStandardContextMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createStandardContextMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QPlainTextEdit) CreateStandardContextMenuWithPosition(position *QPoint) *QMenu {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createStandardContextMenu2, this.h, uintptr(position.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	return bindQMenu(__r)
}

func (this *QPlainTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_cursorForPosition, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQTextCursor(__r)
}

func (this *QPlainTextEdit) CursorRect(cursor *QTextCursor) *QRect {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_cursorRect, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
	return newownedQRect(__r)
}

func (this *QPlainTextEdit) CursorRect2() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_cursorRect2, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPlainTextEdit) AnchorAt(pos *QPoint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_anchorAt, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) OverwriteMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_overwriteMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) SetOverwriteMode(overwrite bool) {
	purego.SyscallN(addr_QPlainTextEdit_setOverwriteMode, this.h, uintptr(b2i(overwrite)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) TabStopDistance() float64 {
	__r := pg_QPlainTextEdit_tabStopDistance(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPlainTextEdit) SetTabStopDistance(distance float64) {
	pg_QPlainTextEdit_setTabStopDistance(this.h, distance)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CursorWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_cursorWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) SetCursorWidth(width int) {
	purego.SyscallN(addr_QPlainTextEdit_setCursorWidth, this.h, uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetExtraSelections(selections []*QTextEdit__ExtraSelection) {
	selections_arr := make([]uintptr, len(selections))
	for __i, __e := range selections {
		selections_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var selections_p uintptr
	if len(selections_arr) > 0 {
		selections_p = uintptr(unsafe.Pointer(&selections_arr[0]))
	}
	purego.SyscallN(addr_QPlainTextEdit_setExtraSelections, this.h, selections_p, uintptr(len(selections)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selections)
	runtime.KeepAlive(selections_arr)
}

func (this *QPlainTextEdit) ExtraSelections() []*QTextEdit__ExtraSelection {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_extraSelections, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextEdit__ExtraSelection, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextEdit__ExtraSelection(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPlainTextEdit) MoveCursor(operation QTextCursor__MoveOperation) {
	purego.SyscallN(addr_QPlainTextEdit_moveCursor, this.h, uintptr(operation))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CanPaste() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_canPaste, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) Print(printer *QPagedPaintDevice) {
	purego.SyscallN(addr_QPlainTextEdit_print, this.h, uintptr(printer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(printer)
}

func (this *QPlainTextEdit) BlockCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_blockCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) InputMethodQuery(property InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_inputMethodQuery, this.h, uintptr(property))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) InputMethodQuery2(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_inputMethodQuery2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) SetPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPlainTextEdit_setPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QPlainTextEdit) Cut() {
	purego.SyscallN(addr_QPlainTextEdit_cut, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Copy() {
	purego.SyscallN(addr_QPlainTextEdit_copy, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Paste() {
	purego.SyscallN(addr_QPlainTextEdit_paste, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Undo() {
	purego.SyscallN(addr_QPlainTextEdit_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Redo() {
	purego.SyscallN(addr_QPlainTextEdit_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Clear() {
	purego.SyscallN(addr_QPlainTextEdit_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SelectAll() {
	purego.SyscallN(addr_QPlainTextEdit_selectAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) InsertPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPlainTextEdit_insertPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QPlainTextEdit) AppendPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPlainTextEdit_appendPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QPlainTextEdit) AppendHtml(html string) {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	purego.SyscallN(addr_QPlainTextEdit_appendHtml, this.h, html_p, uintptr(len(html)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(html)
}

func (this *QPlainTextEdit) CenterCursor() {
	purego.SyscallN(addr_QPlainTextEdit_centerCursor, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) ZoomIn() {
	purego.SyscallN(addr_QPlainTextEdit_zoomIn, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) ZoomOut() {
	purego.SyscallN(addr_QPlainTextEdit_zoomOut, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) TextChanged() {
	purego.SyscallN(addr_QPlainTextEdit_textChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) UndoAvailable(b bool) {
	purego.SyscallN(addr_QPlainTextEdit_undoAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) RedoAvailable(b bool) {
	purego.SyscallN(addr_QPlainTextEdit_redoAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CopyAvailable(b bool) {
	purego.SyscallN(addr_QPlainTextEdit_copyAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SelectionChanged() {
	purego.SyscallN(addr_QPlainTextEdit_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CursorPositionChanged() {
	purego.SyscallN(addr_QPlainTextEdit_cursorPositionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) UpdateRequest(rect *QRect, dy int) {
	purego.SyscallN(addr_QPlainTextEdit_updateRequest, this.h, uintptr(rect.UnsafePointer()), uintptr(dy))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPlainTextEdit) BlockCountChanged(newBlockCount int) {
	purego.SyscallN(addr_QPlainTextEdit_blockCountChanged, this.h, uintptr(newBlockCount))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) ModificationChanged(arg1 bool) {
	purego.SyscallN(addr_QPlainTextEdit_modificationChanged, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) FirstVisibleBlock() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_firstVisibleBlock, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QPlainTextEdit) ContentOffset() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_contentOffset, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPlainTextEdit) BlockBoundingRect(block *QTextBlock) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_blockBoundingRect, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
	return newownedQRectF(__r)
}

func (this *QPlainTextEdit) BlockBoundingGeometry(block *QTextBlock) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_blockBoundingGeometry, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
	return newownedQRectF(__r)
}

func (this *QPlainTextEdit) GetPaintContext() *QAbstractTextDocumentLayout__PaintContext {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_getPaintContext, this.h)
	runtime.KeepAlive(this)
	return newownedQAbstractTextDocumentLayout__PaintContext(__r)
}

func (this *QPlainTextEdit) ZoomInF(range_ float32) {
	pg_QPlainTextEdit_zoomInF(this.h, range_)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QPlainTextEdit_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QPlainTextEdit_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QPlainTextEdit) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QPlainTextEdit) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QPlainTextEdit_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QPlainTextEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QPlainTextEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QPlainTextEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QPlainTextEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QPlainTextEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QPlainTextEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QPlainTextEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Create() {
	purego.SyscallN(addr_QPlainTextEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Destroy() {
	purego.SyscallN(addr_QPlainTextEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QPlainTextEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPlainTextEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPlainTextEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPlainTextEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPlainTextEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPlainTextEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QPlainTextEdit_devicePixelRatioFScale()
	return __r
}

func QPlainTextEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPlainTextEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPlainTextEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QPlainTextEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QPlainTextEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPlainTextEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QPlainTextEdit) Find2(exp string, options QTextDocument__FindFlag) bool {
	var exp_p uintptr
	if len(exp) > 0 {
		exp_p = uintptr(unsafe.Pointer(unsafe.StringData(exp)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_find3, this.h, exp_p, uintptr(len(exp)), uintptr(options))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QPlainTextEdit) Find3(exp *QRegularExpression, options QTextDocument__FindFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_find4, this.h, uintptr(exp.UnsafePointer()), uintptr(options))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QPlainTextEdit) MoveCursor2(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode) {
	purego.SyscallN(addr_QPlainTextEdit_moveCursor2, this.h, uintptr(operation), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) ZoomInWithRange(range_ int) {
	purego.SyscallN(addr_QPlainTextEdit_zoomIn2, this.h, uintptr(range_))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) ZoomOutWithRange(range_ int) {
	purego.SyscallN(addr_QPlainTextEdit_zoomOut2, this.h, uintptr(range_))
	runtime.KeepAlive(this)
}

func QPlainTextEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QPlainTextEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QPlainTextEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QPlainTextEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QPlainTextEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QPlainTextEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QPlainTextEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QPlainTextEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QPlainTextEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QPlainTextEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPlainTextEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPlainTextEdit) LoadResourceDefault(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_loadResource_default, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) InputMethodQueryDefault(property InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_inputMethodQuery_default, this.h, uintptr(property))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) InputMethodQuery2Default(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_inputMethodQuery_default2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QPlainTextEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPlainTextEdit) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QPlainTextEdit) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QPlainTextEdit_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QPlainTextEdit_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QPlainTextEdit_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QPlainTextEdit_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QPlainTextEdit_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPlainTextEdit_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPlainTextEdit_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPlainTextEdit_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) MouseDoubleClickEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QPlainTextEdit_mouseDoubleClickEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) ContextMenuEventDefault(e *QContextMenuEvent) {
	purego.SyscallN(addr_QPlainTextEdit_contextMenuEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) DragEnterEventDefault(e *QDragEnterEvent) {
	purego.SyscallN(addr_QPlainTextEdit_dragEnterEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QPlainTextEdit_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QPlainTextEdit_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) DropEventDefault(e *QDropEvent) {
	purego.SyscallN(addr_QPlainTextEdit_dropEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QPlainTextEdit_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QPlainTextEdit_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QPlainTextEdit_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPlainTextEdit) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QPlainTextEdit_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QPlainTextEdit_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QPlainTextEdit) CreateMimeDataFromSelectionDefault() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_createMimeDataFromSelection_default, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QPlainTextEdit) CanInsertFromMimeDataDefault(source *QMimeData) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_canInsertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
	return __r != 0
}

func (this *QPlainTextEdit) InsertFromMimeDataDefault(source *QMimeData) {
	purego.SyscallN(addr_QPlainTextEdit_insertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
}

func (this *QPlainTextEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QPlainTextEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPlainTextEdit) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QPlainTextEdit_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) DoSetTextCursorDefault(cursor *QTextCursor) {
	purego.SyscallN(addr_QPlainTextEdit_doSetTextCursor_default, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QPlainTextEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPlainTextEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPlainTextEdit) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QPlainTextEdit_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QPlainTextEdit) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QPlainTextEdit) ViewportEventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_viewportEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QPlainTextEdit) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPlainTextEdit) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QPlainTextEdit_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QPlainTextEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QPlainTextEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QPlainTextEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPlainTextEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPlainTextEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QPlainTextEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPlainTextEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QPlainTextEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QPlainTextEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QPlainTextEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QPlainTextEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QPlainTextEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QPlainTextEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QPlainTextEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QPlainTextEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPlainTextEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QPlainTextEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPlainTextEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPlainTextEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPlainTextEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPlainTextEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPlainTextEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPlainTextEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPlainTextEdit) OnTextChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextEdit_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnUndoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QPlainTextEdit_connect_undoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnRedoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QPlainTextEdit_connect_redoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnCopyAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QPlainTextEdit_connect_copyAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextEdit_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnCursorPositionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextEdit_connect_cursorPositionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnUpdateRequest(fn func(*QRect, int)) int64 {
	id := registerCallback(func(rect uintptr, dy int32) { fn(UnsafeNewQRect(unsafe.Pointer(rect)), int(dy)) })
	pg_QPlainTextEdit_connect_updateRequest(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnBlockCountChanged(fn func(int)) int64 {
	id := registerCallback(func(newBlockCount int32) { fn(int(newBlockCount)) })
	pg_QPlainTextEdit_connect_blockCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnModificationChanged(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QPlainTextEdit_connect_modificationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QPlainTextEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QPlainTextEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QPlainTextEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QPlainTextEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPlainTextEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPlainTextEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextEdit) Disconnect(handle int64) {
	pg_QPlainTextEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QPlainTextEdit_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -519616681208050567:
		return 1
	case 7345853097340709638:
		return 2
	case 1237236456312824991:
		return 3
	case -5409927581812374498:
		return 4
	case 1023791977679793185:
		return 5
	case 772598475950050299:
		return 6
	case -4302917285655014661:
		return 7
	case 5290442163167458465:
		return 8
	case -332090898078543319:
		return 9
	case -4255494004099806581:
		return 10
	case -4415920049297795213:
		return 11
	case -4029073905697023887:
		return 12
	case 5898328517505784191:
		return 13
	case -6698043222227006015:
		return 14
	case 206122715919917589:
		return 15
	case -3139904158373725006:
		return 16
	case -1781545137352745978:
		return 17
	case 9107058682754591716:
		return 18
	case 975043767033552552:
		return 19
	case -385596010576048993:
		return 20
	case -6825369174751916:
		return 21
	case 4959376063826165943:
		return 22
	case 1919988340934511316:
		return 23
	case 3493558106800165317:
		return 24
	case -2605205276815256896:
		return 25
	case -466553387102125056:
		return 26
	case 5438819661942843222:
		return 27
	case -6118001150126564745:
		return 28
	case 4090661491944147209:
		return 29
	case 7049264599837532491:
		return 30
	case 729690261457684101:
		return 31
	case 4531996381736730688:
		return 32
	case 5217623514630858297:
		return 33
	case -1971465653061861519:
		return 34
	case 503579067252048239:
		return 35
	case -488693769771308531:
		return 36
	case -8659679893868524044:
		return 37
	case 2110182719375245575:
		return 38
	case 4311534060495737339:
		return 39
	case -8266762397995423915:
		return 40
	case 6398007404756346381:
		return 41
	case 6316270346668237895:
		return 42
	case -551070009623711640:
		return 43
	case -9161949150489525396:
		return 44
	case -1989937050019464783:
		return 45
	case -212068866860701141:
		return 46
	case 1697444140088555581:
		return 47
	case -199718707876719129:
		return 48
	case -2039166586246443706:
		return 49
	case 427629649267311777:
		return 50
	case 6477500296980050210:
		return 51
	case 4127240633310655048:
		return 52
	case 5792771328741774991:
		return 53
	case -1185315367526857833:
		return 54
	case 3769868229123868474:
		return 55
	case 7379029982170503008:
		return 56
	case -568369772982215962:
		return 57
	}
	return -1
}

func QPlainTextEdit_enable(h uintptr, idx int32) { pg_QPlainTextEdit__set_override(h, idx) }

func (this *QPlainTextEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPlainTextEdit) OnLoadResource(fn func(int, *QUrl) *QVariant) {
	registerVirt(this.UnsafeGoID(), -519616681208050567, func(type_ int32, name uintptr) uintptr { return uintptr(fn(int(type_), UnsafeNewQUrl(unsafe.Pointer(name))).UnsafePointer()) }) // slot "loadResource"
}

func (this *QPlainTextEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(property int32) uintptr { return uintptr(fn(InputMethodQuery(property)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QPlainTextEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QPlainTextEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QPlainTextEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QPlainTextEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QPlainTextEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QPlainTextEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QPlainTextEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QPlainTextEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QPlainTextEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QPlainTextEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseDoubleClickEvent"
}

func (this *QPlainTextEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QPlainTextEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(e uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(e))) }) // slot "contextMenuEvent"
}

func (this *QPlainTextEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(e uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(e))) }) // slot "dragEnterEvent"
}

func (this *QPlainTextEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QPlainTextEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QPlainTextEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(e uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(e))) }) // slot "dropEvent"
}

func (this *QPlainTextEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QPlainTextEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QPlainTextEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QPlainTextEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QPlainTextEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QPlainTextEdit) OnCreateMimeDataFromSelection(fn func() *QMimeData) {
	registerVirt(this.UnsafeGoID(), 3493558106800165317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "createMimeDataFromSelection"
}

func (this *QPlainTextEdit) OnCanInsertFromMimeData(fn func(*QMimeData) bool) {
	registerVirt(this.UnsafeGoID(), -2605205276815256896, func(source uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(source)))) }) // slot "canInsertFromMimeData"
}

func (this *QPlainTextEdit) OnInsertFromMimeData(fn func(*QMimeData)) {
	registerVirt(this.UnsafeGoID(), -466553387102125056, func(source uintptr) { fn(UnsafeNewQMimeData(unsafe.Pointer(source))) }) // slot "insertFromMimeData"
}

func (this *QPlainTextEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QPlainTextEdit) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QPlainTextEdit) OnDoSetTextCursor(fn func(*QTextCursor)) {
	registerVirt(this.UnsafeGoID(), 4090661491944147209, func(cursor uintptr) { fn(UnsafeNewQTextCursor(unsafe.Pointer(cursor))) }) // slot "doSetTextCursor"
}

func (this *QPlainTextEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QPlainTextEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QPlainTextEdit) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QPlainTextEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QPlainTextEdit) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "viewportEvent"
}

func (this *QPlainTextEdit) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QPlainTextEdit) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QPlainTextEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QPlainTextEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QPlainTextEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QPlainTextEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QPlainTextEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QPlainTextEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QPlainTextEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QPlainTextEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QPlainTextEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QPlainTextEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QPlainTextEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QPlainTextEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QPlainTextEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QPlainTextEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QPlainTextEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QPlainTextEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QPlainTextEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QPlainTextEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPlainTextEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPlainTextEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPlainTextEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPlainTextEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPlainTextEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
