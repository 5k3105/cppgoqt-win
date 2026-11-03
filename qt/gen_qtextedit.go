package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextEdit_chain uintptr
	addr_QTextEdit_new uintptr
	addr_QTextEdit_new2 uintptr
	addr_QTextEdit_new3 uintptr
	addr_QTextEdit_new4 uintptr
	addr_QTextEdit_tr uintptr
	addr_QTextEdit_setDocument uintptr
	addr_QTextEdit_document uintptr
	addr_QTextEdit_setPlaceholderText uintptr
	addr_QTextEdit_placeholderText uintptr
	addr_QTextEdit_setTextCursor uintptr
	addr_QTextEdit_textCursor uintptr
	addr_QTextEdit_isReadOnly uintptr
	addr_QTextEdit_setReadOnly uintptr
	addr_QTextEdit_setTextInteractionFlags uintptr
	addr_QTextEdit_textInteractionFlags uintptr
	pg_QTextEdit_fontPointSize func(uintptr) float64
	addr_QTextEdit_fontFamily uintptr
	addr_QTextEdit_fontWeight uintptr
	addr_QTextEdit_fontUnderline uintptr
	addr_QTextEdit_fontItalic uintptr
	addr_QTextEdit_textColor uintptr
	addr_QTextEdit_textBackgroundColor uintptr
	addr_QTextEdit_currentFont uintptr
	addr_QTextEdit_alignment uintptr
	addr_QTextEdit_mergeCurrentCharFormat uintptr
	addr_QTextEdit_setCurrentCharFormat uintptr
	addr_QTextEdit_currentCharFormat uintptr
	addr_QTextEdit_autoFormatting uintptr
	addr_QTextEdit_setAutoFormatting uintptr
	addr_QTextEdit_tabChangesFocus uintptr
	addr_QTextEdit_setTabChangesFocus uintptr
	addr_QTextEdit_setDocumentTitle uintptr
	addr_QTextEdit_documentTitle uintptr
	addr_QTextEdit_isUndoRedoEnabled uintptr
	addr_QTextEdit_setUndoRedoEnabled uintptr
	addr_QTextEdit_lineWrapMode uintptr
	addr_QTextEdit_setLineWrapMode uintptr
	addr_QTextEdit_lineWrapColumnOrWidth uintptr
	addr_QTextEdit_setLineWrapColumnOrWidth uintptr
	addr_QTextEdit_wordWrapMode uintptr
	addr_QTextEdit_setWordWrapMode uintptr
	addr_QTextEdit_find uintptr
	addr_QTextEdit_find2 uintptr
	addr_QTextEdit_toPlainText uintptr
	addr_QTextEdit_toHtml uintptr
	addr_QTextEdit_toMarkdown uintptr
	addr_QTextEdit_ensureCursorVisible uintptr
	addr_QTextEdit_loadResource uintptr
	addr_QTextEdit_createStandardContextMenu uintptr
	addr_QTextEdit_createStandardContextMenu2 uintptr
	addr_QTextEdit_cursorForPosition uintptr
	addr_QTextEdit_cursorRect uintptr
	addr_QTextEdit_cursorRect2 uintptr
	addr_QTextEdit_anchorAt uintptr
	addr_QTextEdit_overwriteMode uintptr
	addr_QTextEdit_setOverwriteMode uintptr
	pg_QTextEdit_tabStopDistance func(uintptr) float64
	pg_QTextEdit_setTabStopDistance func(uintptr, float64)
	addr_QTextEdit_cursorWidth uintptr
	addr_QTextEdit_setCursorWidth uintptr
	addr_QTextEdit_acceptRichText uintptr
	addr_QTextEdit_setAcceptRichText uintptr
	addr_QTextEdit_setExtraSelections uintptr
	addr_QTextEdit_extraSelections uintptr
	addr_QTextEdit_moveCursor uintptr
	addr_QTextEdit_canPaste uintptr
	addr_QTextEdit_print uintptr
	addr_QTextEdit_inputMethodQuery uintptr
	addr_QTextEdit_inputMethodQuery2 uintptr
	pg_QTextEdit_setFontPointSize func(uintptr, float64)
	addr_QTextEdit_setFontFamily uintptr
	addr_QTextEdit_setFontWeight uintptr
	addr_QTextEdit_setFontUnderline uintptr
	addr_QTextEdit_setFontItalic uintptr
	addr_QTextEdit_setTextColor uintptr
	addr_QTextEdit_setTextBackgroundColor uintptr
	addr_QTextEdit_setCurrentFont uintptr
	addr_QTextEdit_setAlignment uintptr
	addr_QTextEdit_setPlainText uintptr
	addr_QTextEdit_setHtml uintptr
	addr_QTextEdit_setMarkdown uintptr
	addr_QTextEdit_setText uintptr
	addr_QTextEdit_cut uintptr
	addr_QTextEdit_copy uintptr
	addr_QTextEdit_paste uintptr
	addr_QTextEdit_undo uintptr
	addr_QTextEdit_redo uintptr
	addr_QTextEdit_clear uintptr
	addr_QTextEdit_selectAll uintptr
	addr_QTextEdit_insertPlainText uintptr
	addr_QTextEdit_insertHtml uintptr
	addr_QTextEdit_append uintptr
	addr_QTextEdit_scrollToAnchor uintptr
	addr_QTextEdit_zoomIn uintptr
	addr_QTextEdit_zoomOut uintptr
	addr_QTextEdit_textChanged uintptr
	addr_QTextEdit_undoAvailable uintptr
	addr_QTextEdit_redoAvailable uintptr
	addr_QTextEdit_currentCharFormatChanged uintptr
	addr_QTextEdit_copyAvailable uintptr
	addr_QTextEdit_selectionChanged uintptr
	addr_QTextEdit_cursorPositionChanged uintptr
	pg_QTextEdit_zoomInF func(uintptr, float32)
	addr_QTextEdit_setViewportMargins uintptr
	addr_QTextEdit_setViewportMargins2 uintptr
	addr_QTextEdit_viewportMargins uintptr
	addr_QTextEdit_drawFrame uintptr
	addr_QTextEdit_setTabOrder uintptr
	addr_QTextEdit_mouseGrabber uintptr
	addr_QTextEdit_keyboardGrabber uintptr
	addr_QTextEdit_createWindowContainer uintptr
	addr_QTextEdit_updateMicroFocus uintptr
	addr_QTextEdit_create uintptr
	addr_QTextEdit_destroy uintptr
	addr_QTextEdit_focusNextChild uintptr
	addr_QTextEdit_focusPreviousChild uintptr
	addr_QTextEdit_connect uintptr
	addr_QTextEdit_connect2 uintptr
	addr_QTextEdit_sender uintptr
	addr_QTextEdit_senderSignalIndex uintptr
	addr_QTextEdit_receivers uintptr
	addr_QTextEdit_isSignalConnected uintptr
	pg_QTextEdit_devicePixelRatioFScale func() float64
	pg_QTextEdit_encodeMetricF func(int32, float64) int32
	pg_QTextEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTextEdit_tr2 uintptr
	addr_QTextEdit_tr3 uintptr
	addr_QTextEdit_find3 uintptr
	addr_QTextEdit_find4 uintptr
	addr_QTextEdit_toMarkdown2 uintptr
	addr_QTextEdit_moveCursor2 uintptr
	addr_QTextEdit_zoomIn2 uintptr
	addr_QTextEdit_zoomOut2 uintptr
	addr_QTextEdit_createWindowContainer2 uintptr
	addr_QTextEdit_createWindowContainer3 uintptr
	addr_QTextEdit_updateMicroFocus2 uintptr
	addr_QTextEdit_create2 uintptr
	addr_QTextEdit_create3 uintptr
	addr_QTextEdit_create4 uintptr
	addr_QTextEdit_destroy2 uintptr
	addr_QTextEdit_destroy3 uintptr
	addr_QTextEdit_connect3 uintptr
	addr_QTextEdit_connect4 uintptr
	addr_QTextEdit_loadResource_default uintptr
	addr_QTextEdit_inputMethodQuery_default uintptr
	addr_QTextEdit_inputMethodQuery_default2 uintptr
	addr_QTextEdit_metaObject_default uintptr
	addr_QTextEdit_event_default uintptr
	addr_QTextEdit_timerEvent_default uintptr
	addr_QTextEdit_keyPressEvent_default uintptr
	addr_QTextEdit_keyReleaseEvent_default uintptr
	addr_QTextEdit_resizeEvent_default uintptr
	addr_QTextEdit_paintEvent_default uintptr
	addr_QTextEdit_mousePressEvent_default uintptr
	addr_QTextEdit_mouseMoveEvent_default uintptr
	addr_QTextEdit_mouseReleaseEvent_default uintptr
	addr_QTextEdit_mouseDoubleClickEvent_default uintptr
	addr_QTextEdit_focusNextPrevChild_default uintptr
	addr_QTextEdit_contextMenuEvent_default uintptr
	addr_QTextEdit_dragEnterEvent_default uintptr
	addr_QTextEdit_dragLeaveEvent_default uintptr
	addr_QTextEdit_dragMoveEvent_default uintptr
	addr_QTextEdit_dropEvent_default uintptr
	addr_QTextEdit_focusInEvent_default uintptr
	addr_QTextEdit_focusOutEvent_default uintptr
	addr_QTextEdit_showEvent_default uintptr
	addr_QTextEdit_changeEvent_default uintptr
	addr_QTextEdit_wheelEvent_default uintptr
	addr_QTextEdit_createMimeDataFromSelection_default uintptr
	addr_QTextEdit_canInsertFromMimeData_default uintptr
	addr_QTextEdit_insertFromMimeData_default uintptr
	addr_QTextEdit_inputMethodEvent_default uintptr
	addr_QTextEdit_scrollContentsBy_default uintptr
	addr_QTextEdit_doSetTextCursor_default uintptr
	addr_QTextEdit_minimumSizeHint_default uintptr
	addr_QTextEdit_sizeHint_default uintptr
	addr_QTextEdit_setupViewport_default uintptr
	addr_QTextEdit_eventFilter_default uintptr
	addr_QTextEdit_viewportEvent_default uintptr
	addr_QTextEdit_viewportSizeHint_default uintptr
	addr_QTextEdit_initStyleOption_default uintptr
	addr_QTextEdit_devType_default uintptr
	addr_QTextEdit_setVisible_default uintptr
	addr_QTextEdit_heightForWidth_default uintptr
	addr_QTextEdit_hasHeightForWidth_default uintptr
	addr_QTextEdit_paintEngine_default uintptr
	addr_QTextEdit_enterEvent_default uintptr
	addr_QTextEdit_leaveEvent_default uintptr
	addr_QTextEdit_moveEvent_default uintptr
	addr_QTextEdit_closeEvent_default uintptr
	addr_QTextEdit_tabletEvent_default uintptr
	addr_QTextEdit_actionEvent_default uintptr
	addr_QTextEdit_hideEvent_default uintptr
	addr_QTextEdit_metric_default uintptr
	addr_QTextEdit_initPainter_default uintptr
	addr_QTextEdit_redirected_default uintptr
	addr_QTextEdit_sharedPainter_default uintptr
	addr_QTextEdit_childEvent_default uintptr
	addr_QTextEdit_customEvent_default uintptr
	addr_QTextEdit_connectNotify_default uintptr
	addr_QTextEdit_disconnectNotify_default uintptr
	pg_QTextEdit_connect_textChanged func(uintptr, int64)
	pg_QTextEdit_connect_undoAvailable func(uintptr, int64)
	pg_QTextEdit_connect_redoAvailable func(uintptr, int64)
	pg_QTextEdit_connect_currentCharFormatChanged func(uintptr, int64)
	pg_QTextEdit_connect_copyAvailable func(uintptr, int64)
	pg_QTextEdit_connect_selectionChanged func(uintptr, int64)
	pg_QTextEdit_connect_cursorPositionChanged func(uintptr, int64)
	pg_QTextEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QTextEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QTextEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTextEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTextEdit_connect_destroyed func(uintptr, int64)
	pg_QTextEdit_connect_destroyed2 func(uintptr, int64)
	pg_QTextEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QTextEdit__disconnect func(int64)
	pg_QTextEdit__set_override func(uintptr, int32)
	pg_QTextEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextEdit_chain, "QTextEdit_chain")
	symaddr(&addr_QTextEdit_new, "QTextEdit_new")
	symaddr(&addr_QTextEdit_new2, "QTextEdit_new2")
	symaddr(&addr_QTextEdit_new3, "QTextEdit_new3")
	symaddr(&addr_QTextEdit_new4, "QTextEdit_new4")
	symaddr(&addr_QTextEdit_tr, "QTextEdit_tr")
	symaddr(&addr_QTextEdit_setDocument, "QTextEdit_setDocument")
	symaddr(&addr_QTextEdit_document, "QTextEdit_document")
	symaddr(&addr_QTextEdit_setPlaceholderText, "QTextEdit_setPlaceholderText")
	symaddr(&addr_QTextEdit_placeholderText, "QTextEdit_placeholderText")
	symaddr(&addr_QTextEdit_setTextCursor, "QTextEdit_setTextCursor")
	symaddr(&addr_QTextEdit_textCursor, "QTextEdit_textCursor")
	symaddr(&addr_QTextEdit_isReadOnly, "QTextEdit_isReadOnly")
	symaddr(&addr_QTextEdit_setReadOnly, "QTextEdit_setReadOnly")
	symaddr(&addr_QTextEdit_setTextInteractionFlags, "QTextEdit_setTextInteractionFlags")
	symaddr(&addr_QTextEdit_textInteractionFlags, "QTextEdit_textInteractionFlags")
	reg(&pg_QTextEdit_fontPointSize, "QTextEdit_fontPointSize")
	symaddr(&addr_QTextEdit_fontFamily, "QTextEdit_fontFamily")
	symaddr(&addr_QTextEdit_fontWeight, "QTextEdit_fontWeight")
	symaddr(&addr_QTextEdit_fontUnderline, "QTextEdit_fontUnderline")
	symaddr(&addr_QTextEdit_fontItalic, "QTextEdit_fontItalic")
	symaddr(&addr_QTextEdit_textColor, "QTextEdit_textColor")
	symaddr(&addr_QTextEdit_textBackgroundColor, "QTextEdit_textBackgroundColor")
	symaddr(&addr_QTextEdit_currentFont, "QTextEdit_currentFont")
	symaddr(&addr_QTextEdit_alignment, "QTextEdit_alignment")
	symaddr(&addr_QTextEdit_mergeCurrentCharFormat, "QTextEdit_mergeCurrentCharFormat")
	symaddr(&addr_QTextEdit_setCurrentCharFormat, "QTextEdit_setCurrentCharFormat")
	symaddr(&addr_QTextEdit_currentCharFormat, "QTextEdit_currentCharFormat")
	symaddr(&addr_QTextEdit_autoFormatting, "QTextEdit_autoFormatting")
	symaddr(&addr_QTextEdit_setAutoFormatting, "QTextEdit_setAutoFormatting")
	symaddr(&addr_QTextEdit_tabChangesFocus, "QTextEdit_tabChangesFocus")
	symaddr(&addr_QTextEdit_setTabChangesFocus, "QTextEdit_setTabChangesFocus")
	symaddr(&addr_QTextEdit_setDocumentTitle, "QTextEdit_setDocumentTitle")
	symaddr(&addr_QTextEdit_documentTitle, "QTextEdit_documentTitle")
	symaddr(&addr_QTextEdit_isUndoRedoEnabled, "QTextEdit_isUndoRedoEnabled")
	symaddr(&addr_QTextEdit_setUndoRedoEnabled, "QTextEdit_setUndoRedoEnabled")
	symaddr(&addr_QTextEdit_lineWrapMode, "QTextEdit_lineWrapMode")
	symaddr(&addr_QTextEdit_setLineWrapMode, "QTextEdit_setLineWrapMode")
	symaddr(&addr_QTextEdit_lineWrapColumnOrWidth, "QTextEdit_lineWrapColumnOrWidth")
	symaddr(&addr_QTextEdit_setLineWrapColumnOrWidth, "QTextEdit_setLineWrapColumnOrWidth")
	symaddr(&addr_QTextEdit_wordWrapMode, "QTextEdit_wordWrapMode")
	symaddr(&addr_QTextEdit_setWordWrapMode, "QTextEdit_setWordWrapMode")
	symaddr(&addr_QTextEdit_find, "QTextEdit_find")
	symaddr(&addr_QTextEdit_find2, "QTextEdit_find2")
	symaddr(&addr_QTextEdit_toPlainText, "QTextEdit_toPlainText")
	symaddr(&addr_QTextEdit_toHtml, "QTextEdit_toHtml")
	symaddr(&addr_QTextEdit_toMarkdown, "QTextEdit_toMarkdown")
	symaddr(&addr_QTextEdit_ensureCursorVisible, "QTextEdit_ensureCursorVisible")
	symaddr(&addr_QTextEdit_loadResource, "QTextEdit_loadResource")
	symaddr(&addr_QTextEdit_createStandardContextMenu, "QTextEdit_createStandardContextMenu")
	symaddr(&addr_QTextEdit_createStandardContextMenu2, "QTextEdit_createStandardContextMenu2")
	symaddr(&addr_QTextEdit_cursorForPosition, "QTextEdit_cursorForPosition")
	symaddr(&addr_QTextEdit_cursorRect, "QTextEdit_cursorRect")
	symaddr(&addr_QTextEdit_cursorRect2, "QTextEdit_cursorRect2")
	symaddr(&addr_QTextEdit_anchorAt, "QTextEdit_anchorAt")
	symaddr(&addr_QTextEdit_overwriteMode, "QTextEdit_overwriteMode")
	symaddr(&addr_QTextEdit_setOverwriteMode, "QTextEdit_setOverwriteMode")
	reg(&pg_QTextEdit_tabStopDistance, "QTextEdit_tabStopDistance")
	reg(&pg_QTextEdit_setTabStopDistance, "QTextEdit_setTabStopDistance")
	symaddr(&addr_QTextEdit_cursorWidth, "QTextEdit_cursorWidth")
	symaddr(&addr_QTextEdit_setCursorWidth, "QTextEdit_setCursorWidth")
	symaddr(&addr_QTextEdit_acceptRichText, "QTextEdit_acceptRichText")
	symaddr(&addr_QTextEdit_setAcceptRichText, "QTextEdit_setAcceptRichText")
	symaddr(&addr_QTextEdit_setExtraSelections, "QTextEdit_setExtraSelections")
	symaddr(&addr_QTextEdit_extraSelections, "QTextEdit_extraSelections")
	symaddr(&addr_QTextEdit_moveCursor, "QTextEdit_moveCursor")
	symaddr(&addr_QTextEdit_canPaste, "QTextEdit_canPaste")
	symaddr(&addr_QTextEdit_print, "QTextEdit_print")
	symaddr(&addr_QTextEdit_inputMethodQuery, "QTextEdit_inputMethodQuery")
	symaddr(&addr_QTextEdit_inputMethodQuery2, "QTextEdit_inputMethodQuery2")
	reg(&pg_QTextEdit_setFontPointSize, "QTextEdit_setFontPointSize")
	symaddr(&addr_QTextEdit_setFontFamily, "QTextEdit_setFontFamily")
	symaddr(&addr_QTextEdit_setFontWeight, "QTextEdit_setFontWeight")
	symaddr(&addr_QTextEdit_setFontUnderline, "QTextEdit_setFontUnderline")
	symaddr(&addr_QTextEdit_setFontItalic, "QTextEdit_setFontItalic")
	symaddr(&addr_QTextEdit_setTextColor, "QTextEdit_setTextColor")
	symaddr(&addr_QTextEdit_setTextBackgroundColor, "QTextEdit_setTextBackgroundColor")
	symaddr(&addr_QTextEdit_setCurrentFont, "QTextEdit_setCurrentFont")
	symaddr(&addr_QTextEdit_setAlignment, "QTextEdit_setAlignment")
	symaddr(&addr_QTextEdit_setPlainText, "QTextEdit_setPlainText")
	symaddr(&addr_QTextEdit_setHtml, "QTextEdit_setHtml")
	symaddr(&addr_QTextEdit_setMarkdown, "QTextEdit_setMarkdown")
	symaddr(&addr_QTextEdit_setText, "QTextEdit_setText")
	symaddr(&addr_QTextEdit_cut, "QTextEdit_cut")
	symaddr(&addr_QTextEdit_copy, "QTextEdit_copy")
	symaddr(&addr_QTextEdit_paste, "QTextEdit_paste")
	symaddr(&addr_QTextEdit_undo, "QTextEdit_undo")
	symaddr(&addr_QTextEdit_redo, "QTextEdit_redo")
	symaddr(&addr_QTextEdit_clear, "QTextEdit_clear")
	symaddr(&addr_QTextEdit_selectAll, "QTextEdit_selectAll")
	symaddr(&addr_QTextEdit_insertPlainText, "QTextEdit_insertPlainText")
	symaddr(&addr_QTextEdit_insertHtml, "QTextEdit_insertHtml")
	symaddr(&addr_QTextEdit_append, "QTextEdit_append")
	symaddr(&addr_QTextEdit_scrollToAnchor, "QTextEdit_scrollToAnchor")
	symaddr(&addr_QTextEdit_zoomIn, "QTextEdit_zoomIn")
	symaddr(&addr_QTextEdit_zoomOut, "QTextEdit_zoomOut")
	symaddr(&addr_QTextEdit_textChanged, "QTextEdit_textChanged")
	symaddr(&addr_QTextEdit_undoAvailable, "QTextEdit_undoAvailable")
	symaddr(&addr_QTextEdit_redoAvailable, "QTextEdit_redoAvailable")
	symaddr(&addr_QTextEdit_currentCharFormatChanged, "QTextEdit_currentCharFormatChanged")
	symaddr(&addr_QTextEdit_copyAvailable, "QTextEdit_copyAvailable")
	symaddr(&addr_QTextEdit_selectionChanged, "QTextEdit_selectionChanged")
	symaddr(&addr_QTextEdit_cursorPositionChanged, "QTextEdit_cursorPositionChanged")
	reg(&pg_QTextEdit_zoomInF, "QTextEdit_zoomInF")
	symaddr(&addr_QTextEdit_setViewportMargins, "QTextEdit_setViewportMargins")
	symaddr(&addr_QTextEdit_setViewportMargins2, "QTextEdit_setViewportMargins2")
	symaddr(&addr_QTextEdit_viewportMargins, "QTextEdit_viewportMargins")
	symaddr(&addr_QTextEdit_drawFrame, "QTextEdit_drawFrame")
	symaddr(&addr_QTextEdit_setTabOrder, "QTextEdit_setTabOrder")
	symaddr(&addr_QTextEdit_mouseGrabber, "QTextEdit_mouseGrabber")
	symaddr(&addr_QTextEdit_keyboardGrabber, "QTextEdit_keyboardGrabber")
	symaddr(&addr_QTextEdit_createWindowContainer, "QTextEdit_createWindowContainer")
	symaddr(&addr_QTextEdit_updateMicroFocus, "QTextEdit_updateMicroFocus")
	symaddr(&addr_QTextEdit_create, "QTextEdit_create")
	symaddr(&addr_QTextEdit_destroy, "QTextEdit_destroy")
	symaddr(&addr_QTextEdit_focusNextChild, "QTextEdit_focusNextChild")
	symaddr(&addr_QTextEdit_focusPreviousChild, "QTextEdit_focusPreviousChild")
	symaddr(&addr_QTextEdit_connect, "QTextEdit_connect")
	symaddr(&addr_QTextEdit_connect2, "QTextEdit_connect2")
	symaddr(&addr_QTextEdit_sender, "QTextEdit_sender")
	symaddr(&addr_QTextEdit_senderSignalIndex, "QTextEdit_senderSignalIndex")
	symaddr(&addr_QTextEdit_receivers, "QTextEdit_receivers")
	symaddr(&addr_QTextEdit_isSignalConnected, "QTextEdit_isSignalConnected")
	reg(&pg_QTextEdit_devicePixelRatioFScale, "QTextEdit_devicePixelRatioFScale")
	reg(&pg_QTextEdit_encodeMetricF, "QTextEdit_encodeMetricF")
	reg(&pg_QTextEdit_getDecodedMetricF, "QTextEdit_getDecodedMetricF")
	symaddr(&addr_QTextEdit_tr2, "QTextEdit_tr2")
	symaddr(&addr_QTextEdit_tr3, "QTextEdit_tr3")
	symaddr(&addr_QTextEdit_find3, "QTextEdit_find3")
	symaddr(&addr_QTextEdit_find4, "QTextEdit_find4")
	symaddr(&addr_QTextEdit_toMarkdown2, "QTextEdit_toMarkdown2")
	symaddr(&addr_QTextEdit_moveCursor2, "QTextEdit_moveCursor2")
	symaddr(&addr_QTextEdit_zoomIn2, "QTextEdit_zoomIn2")
	symaddr(&addr_QTextEdit_zoomOut2, "QTextEdit_zoomOut2")
	symaddr(&addr_QTextEdit_createWindowContainer2, "QTextEdit_createWindowContainer2")
	symaddr(&addr_QTextEdit_createWindowContainer3, "QTextEdit_createWindowContainer3")
	symaddr(&addr_QTextEdit_updateMicroFocus2, "QTextEdit_updateMicroFocus2")
	symaddr(&addr_QTextEdit_create2, "QTextEdit_create2")
	symaddr(&addr_QTextEdit_create3, "QTextEdit_create3")
	symaddr(&addr_QTextEdit_create4, "QTextEdit_create4")
	symaddr(&addr_QTextEdit_destroy2, "QTextEdit_destroy2")
	symaddr(&addr_QTextEdit_destroy3, "QTextEdit_destroy3")
	symaddr(&addr_QTextEdit_connect3, "QTextEdit_connect3")
	symaddr(&addr_QTextEdit_connect4, "QTextEdit_connect4")
	symaddr(&addr_QTextEdit_loadResource_default, "QTextEdit_loadResource_default")
	symaddr(&addr_QTextEdit_inputMethodQuery_default, "QTextEdit_inputMethodQuery_default")
	symaddr(&addr_QTextEdit_inputMethodQuery_default2, "QTextEdit_inputMethodQuery_default2")
	symaddr(&addr_QTextEdit_metaObject_default, "QTextEdit_metaObject_default")
	symaddr(&addr_QTextEdit_event_default, "QTextEdit_event_default")
	symaddr(&addr_QTextEdit_timerEvent_default, "QTextEdit_timerEvent_default")
	symaddr(&addr_QTextEdit_keyPressEvent_default, "QTextEdit_keyPressEvent_default")
	symaddr(&addr_QTextEdit_keyReleaseEvent_default, "QTextEdit_keyReleaseEvent_default")
	symaddr(&addr_QTextEdit_resizeEvent_default, "QTextEdit_resizeEvent_default")
	symaddr(&addr_QTextEdit_paintEvent_default, "QTextEdit_paintEvent_default")
	symaddr(&addr_QTextEdit_mousePressEvent_default, "QTextEdit_mousePressEvent_default")
	symaddr(&addr_QTextEdit_mouseMoveEvent_default, "QTextEdit_mouseMoveEvent_default")
	symaddr(&addr_QTextEdit_mouseReleaseEvent_default, "QTextEdit_mouseReleaseEvent_default")
	symaddr(&addr_QTextEdit_mouseDoubleClickEvent_default, "QTextEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QTextEdit_focusNextPrevChild_default, "QTextEdit_focusNextPrevChild_default")
	symaddr(&addr_QTextEdit_contextMenuEvent_default, "QTextEdit_contextMenuEvent_default")
	symaddr(&addr_QTextEdit_dragEnterEvent_default, "QTextEdit_dragEnterEvent_default")
	symaddr(&addr_QTextEdit_dragLeaveEvent_default, "QTextEdit_dragLeaveEvent_default")
	symaddr(&addr_QTextEdit_dragMoveEvent_default, "QTextEdit_dragMoveEvent_default")
	symaddr(&addr_QTextEdit_dropEvent_default, "QTextEdit_dropEvent_default")
	symaddr(&addr_QTextEdit_focusInEvent_default, "QTextEdit_focusInEvent_default")
	symaddr(&addr_QTextEdit_focusOutEvent_default, "QTextEdit_focusOutEvent_default")
	symaddr(&addr_QTextEdit_showEvent_default, "QTextEdit_showEvent_default")
	symaddr(&addr_QTextEdit_changeEvent_default, "QTextEdit_changeEvent_default")
	symaddr(&addr_QTextEdit_wheelEvent_default, "QTextEdit_wheelEvent_default")
	symaddr(&addr_QTextEdit_createMimeDataFromSelection_default, "QTextEdit_createMimeDataFromSelection_default")
	symaddr(&addr_QTextEdit_canInsertFromMimeData_default, "QTextEdit_canInsertFromMimeData_default")
	symaddr(&addr_QTextEdit_insertFromMimeData_default, "QTextEdit_insertFromMimeData_default")
	symaddr(&addr_QTextEdit_inputMethodEvent_default, "QTextEdit_inputMethodEvent_default")
	symaddr(&addr_QTextEdit_scrollContentsBy_default, "QTextEdit_scrollContentsBy_default")
	symaddr(&addr_QTextEdit_doSetTextCursor_default, "QTextEdit_doSetTextCursor_default")
	symaddr(&addr_QTextEdit_minimumSizeHint_default, "QTextEdit_minimumSizeHint_default")
	symaddr(&addr_QTextEdit_sizeHint_default, "QTextEdit_sizeHint_default")
	symaddr(&addr_QTextEdit_setupViewport_default, "QTextEdit_setupViewport_default")
	symaddr(&addr_QTextEdit_eventFilter_default, "QTextEdit_eventFilter_default")
	symaddr(&addr_QTextEdit_viewportEvent_default, "QTextEdit_viewportEvent_default")
	symaddr(&addr_QTextEdit_viewportSizeHint_default, "QTextEdit_viewportSizeHint_default")
	symaddr(&addr_QTextEdit_initStyleOption_default, "QTextEdit_initStyleOption_default")
	symaddr(&addr_QTextEdit_devType_default, "QTextEdit_devType_default")
	symaddr(&addr_QTextEdit_setVisible_default, "QTextEdit_setVisible_default")
	symaddr(&addr_QTextEdit_heightForWidth_default, "QTextEdit_heightForWidth_default")
	symaddr(&addr_QTextEdit_hasHeightForWidth_default, "QTextEdit_hasHeightForWidth_default")
	symaddr(&addr_QTextEdit_paintEngine_default, "QTextEdit_paintEngine_default")
	symaddr(&addr_QTextEdit_enterEvent_default, "QTextEdit_enterEvent_default")
	symaddr(&addr_QTextEdit_leaveEvent_default, "QTextEdit_leaveEvent_default")
	symaddr(&addr_QTextEdit_moveEvent_default, "QTextEdit_moveEvent_default")
	symaddr(&addr_QTextEdit_closeEvent_default, "QTextEdit_closeEvent_default")
	symaddr(&addr_QTextEdit_tabletEvent_default, "QTextEdit_tabletEvent_default")
	symaddr(&addr_QTextEdit_actionEvent_default, "QTextEdit_actionEvent_default")
	symaddr(&addr_QTextEdit_hideEvent_default, "QTextEdit_hideEvent_default")
	symaddr(&addr_QTextEdit_metric_default, "QTextEdit_metric_default")
	symaddr(&addr_QTextEdit_initPainter_default, "QTextEdit_initPainter_default")
	symaddr(&addr_QTextEdit_redirected_default, "QTextEdit_redirected_default")
	symaddr(&addr_QTextEdit_sharedPainter_default, "QTextEdit_sharedPainter_default")
	symaddr(&addr_QTextEdit_childEvent_default, "QTextEdit_childEvent_default")
	symaddr(&addr_QTextEdit_customEvent_default, "QTextEdit_customEvent_default")
	symaddr(&addr_QTextEdit_connectNotify_default, "QTextEdit_connectNotify_default")
	symaddr(&addr_QTextEdit_disconnectNotify_default, "QTextEdit_disconnectNotify_default")
	reg(&pg_QTextEdit_connect_textChanged, "QTextEdit_connect_textChanged")
	reg(&pg_QTextEdit_connect_undoAvailable, "QTextEdit_connect_undoAvailable")
	reg(&pg_QTextEdit_connect_redoAvailable, "QTextEdit_connect_redoAvailable")
	reg(&pg_QTextEdit_connect_currentCharFormatChanged, "QTextEdit_connect_currentCharFormatChanged")
	reg(&pg_QTextEdit_connect_copyAvailable, "QTextEdit_connect_copyAvailable")
	reg(&pg_QTextEdit_connect_selectionChanged, "QTextEdit_connect_selectionChanged")
	reg(&pg_QTextEdit_connect_cursorPositionChanged, "QTextEdit_connect_cursorPositionChanged")
	reg(&pg_QTextEdit_connect_windowTitleChanged, "QTextEdit_connect_windowTitleChanged")
	reg(&pg_QTextEdit_connect_windowIconChanged, "QTextEdit_connect_windowIconChanged")
	reg(&pg_QTextEdit_connect_windowIconTextChanged, "QTextEdit_connect_windowIconTextChanged")
	reg(&pg_QTextEdit_connect_customContextMenuRequested, "QTextEdit_connect_customContextMenuRequested")
	reg(&pg_QTextEdit_connect_destroyed, "QTextEdit_connect_destroyed")
	reg(&pg_QTextEdit_connect_destroyed2, "QTextEdit_connect_destroyed2")
	reg(&pg_QTextEdit_connect_objectNameChanged, "QTextEdit_connect_objectNameChanged")
	reg(&pg_QTextEdit__disconnect, "QTextEdit__disconnect")
	reg(&pg_QTextEdit__set_override, "QTextEdit__set_override")
	reg(&pg_QTextEdit_delete, "QTextEdit_delete")
}

type QTextEdit struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQTextEdit(h uintptr) *QTextEdit {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QTextEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextEdit{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQTextEdit(h uintptr) *QTextEdit { return bindQTextEdit(h) }

func (this *QTextEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextEdit(h unsafe.Pointer) *QTextEdit { return bindQTextEdit(uintptr(h)) }
func UnsafeNewQTextEditOwned(h unsafe.Pointer) *QTextEdit { return newownedQTextEdit(uintptr(h)) }

func (this *QTextEdit) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func NewQTextEdit(parent *QWidget) *QTextEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTextEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextEdit_virtIdx, QTextEdit_enable)
	return __o
}

func NewQTextEdit2() *QTextEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextEdit_new2, uintptr(__goid))
	__o := newownedQTextEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextEdit_virtIdx, QTextEdit_enable)
	return __o
}

func NewQTextEdit3(text string) *QTextEdit {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextEdit_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQTextEdit(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextEdit_virtIdx, QTextEdit_enable)
	return __o
}

func NewQTextEdit4(text string, parent *QWidget) *QTextEdit {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextEdit_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQTextEdit(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextEdit_virtIdx, QTextEdit_enable)
	return __o
}

func QTextEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextEdit) SetDocument(document *QTextDocument) {
	purego.SyscallN(addr_QTextEdit_setDocument, this.h, uintptr(document.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(document)
}

func (this *QTextEdit) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QTextEdit) SetPlaceholderText(placeholderText string) {
	var placeholderText_p uintptr
	if len(placeholderText) > 0 {
		placeholderText_p = uintptr(unsafe.Pointer(unsafe.StringData(placeholderText)))
	}
	purego.SyscallN(addr_QTextEdit_setPlaceholderText, this.h, placeholderText_p, uintptr(len(placeholderText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(placeholderText)
}

func (this *QTextEdit) PlaceholderText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_placeholderText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) SetTextCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextEdit_setTextCursor, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextEdit) TextCursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_textCursor, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextEdit) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) SetReadOnly(ro bool) {
	purego.SyscallN(addr_QTextEdit_setReadOnly, this.h, uintptr(b2i(ro)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetTextInteractionFlags(flags TextInteractionFlag) {
	purego.SyscallN(addr_QTextEdit_setTextInteractionFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) TextInteractionFlags() TextInteractionFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_textInteractionFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextInteractionFlag(__r)
}

func (this *QTextEdit) FontPointSize() float64 {
	__r := pg_QTextEdit_fontPointSize(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextEdit) FontFamily() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_fontFamily, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) FontWeight() int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_fontWeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) FontUnderline() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_fontUnderline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) FontItalic() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_fontItalic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) TextColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_textColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QTextEdit) TextBackgroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_textBackgroundColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QTextEdit) CurrentFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_currentFont, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTextEdit) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QTextEdit) MergeCurrentCharFormat(modifier *QTextCharFormat) {
	purego.SyscallN(addr_QTextEdit_mergeCurrentCharFormat, this.h, uintptr(modifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modifier)
}

func (this *QTextEdit) SetCurrentCharFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextEdit_setCurrentCharFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextEdit) CurrentCharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_currentCharFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextEdit) AutoFormatting() QTextEdit__AutoFormattingFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_autoFormatting, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextEdit__AutoFormattingFlag(__r)
}

func (this *QTextEdit) SetAutoFormatting(features QTextEdit__AutoFormattingFlag) {
	purego.SyscallN(addr_QTextEdit_setAutoFormatting, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) TabChangesFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_tabChangesFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) SetTabChangesFocus(b bool) {
	purego.SyscallN(addr_QTextEdit_setTabChangesFocus, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetDocumentTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QTextEdit_setDocumentTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QTextEdit) DocumentTitle() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_documentTitle, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) IsUndoRedoEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_isUndoRedoEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) SetUndoRedoEnabled(enable bool) {
	purego.SyscallN(addr_QTextEdit_setUndoRedoEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) LineWrapMode() QTextEdit__LineWrapMode {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_lineWrapMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextEdit__LineWrapMode(__r)
}

func (this *QTextEdit) SetLineWrapMode(mode QTextEdit__LineWrapMode) {
	purego.SyscallN(addr_QTextEdit_setLineWrapMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) LineWrapColumnOrWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_lineWrapColumnOrWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) SetLineWrapColumnOrWidth(w int) {
	purego.SyscallN(addr_QTextEdit_setLineWrapColumnOrWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) WordWrapMode() QTextOption__WrapMode {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_wordWrapMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextOption__WrapMode(__r)
}

func (this *QTextEdit) SetWordWrapMode(policy QTextOption__WrapMode) {
	purego.SyscallN(addr_QTextEdit_setWordWrapMode, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Find(exp string) bool {
	var exp_p uintptr
	if len(exp) > 0 {
		exp_p = uintptr(unsafe.Pointer(unsafe.StringData(exp)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_find, this.h, exp_p, uintptr(len(exp)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QTextEdit) FindWithExp(exp *QRegularExpression) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_find2, this.h, uintptr(exp.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QTextEdit) ToPlainText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_toPlainText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) ToHtml() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_toHtml, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) ToMarkdown() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_toMarkdown, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) EnsureCursorVisible() {
	purego.SyscallN(addr_QTextEdit_ensureCursorVisible, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) LoadResource(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_loadResource, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QTextEdit) CreateStandardContextMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createStandardContextMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QTextEdit) CreateStandardContextMenuWithPosition(position *QPoint) *QMenu {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createStandardContextMenu2, this.h, uintptr(position.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	return bindQMenu(__r)
}

func (this *QTextEdit) CursorForPosition(pos *QPoint) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_cursorForPosition, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQTextCursor(__r)
}

func (this *QTextEdit) CursorRect(cursor *QTextCursor) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_cursorRect, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
	return newownedQRect(__r)
}

func (this *QTextEdit) CursorRect2() *QRect {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_cursorRect2, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QTextEdit) AnchorAt(pos *QPoint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_anchorAt, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return pgStr(__r, __n)
}

func (this *QTextEdit) OverwriteMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_overwriteMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) SetOverwriteMode(overwrite bool) {
	purego.SyscallN(addr_QTextEdit_setOverwriteMode, this.h, uintptr(b2i(overwrite)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) TabStopDistance() float64 {
	__r := pg_QTextEdit_tabStopDistance(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextEdit) SetTabStopDistance(distance float64) {
	pg_QTextEdit_setTabStopDistance(this.h, distance)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) CursorWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_cursorWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) SetCursorWidth(width int) {
	purego.SyscallN(addr_QTextEdit_setCursorWidth, this.h, uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) AcceptRichText() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_acceptRichText, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) SetAcceptRichText(accept bool) {
	purego.SyscallN(addr_QTextEdit_setAcceptRichText, this.h, uintptr(b2i(accept)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetExtraSelections(selections []*QTextEdit__ExtraSelection) {
	selections_arr := make([]uintptr, len(selections))
	for __i, __e := range selections {
		selections_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var selections_p uintptr
	if len(selections_arr) > 0 {
		selections_p = uintptr(unsafe.Pointer(&selections_arr[0]))
	}
	purego.SyscallN(addr_QTextEdit_setExtraSelections, this.h, selections_p, uintptr(len(selections)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selections)
	runtime.KeepAlive(selections_arr)
}

func (this *QTextEdit) ExtraSelections() []*QTextEdit__ExtraSelection {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_extraSelections, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextEdit__ExtraSelection, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextEdit__ExtraSelection(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextEdit) MoveCursor(operation QTextCursor__MoveOperation) {
	purego.SyscallN(addr_QTextEdit_moveCursor, this.h, uintptr(operation))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) CanPaste() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_canPaste, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) Print(printer *QPagedPaintDevice) {
	purego.SyscallN(addr_QTextEdit_print, this.h, uintptr(printer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(printer)
}

func (this *QTextEdit) InputMethodQuery(property InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_inputMethodQuery, this.h, uintptr(property))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextEdit) InputMethodQuery2(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_inputMethodQuery2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QTextEdit) SetFontPointSize(s float64) {
	pg_QTextEdit_setFontPointSize(this.h, s)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetFontFamily(fontFamily string) {
	var fontFamily_p uintptr
	if len(fontFamily) > 0 {
		fontFamily_p = uintptr(unsafe.Pointer(unsafe.StringData(fontFamily)))
	}
	purego.SyscallN(addr_QTextEdit_setFontFamily, this.h, fontFamily_p, uintptr(len(fontFamily)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontFamily)
}

func (this *QTextEdit) SetFontWeight(w int) {
	purego.SyscallN(addr_QTextEdit_setFontWeight, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetFontUnderline(b bool) {
	purego.SyscallN(addr_QTextEdit_setFontUnderline, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetFontItalic(b bool) {
	purego.SyscallN(addr_QTextEdit_setFontItalic, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetTextColor(c *QColor) {
	purego.SyscallN(addr_QTextEdit_setTextColor, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QTextEdit) SetTextBackgroundColor(c *QColor) {
	purego.SyscallN(addr_QTextEdit_setTextBackgroundColor, this.h, uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QTextEdit) SetCurrentFont(f *QFont) {
	purego.SyscallN(addr_QTextEdit_setCurrentFont, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QTextEdit) SetAlignment(a AlignmentFlag) {
	purego.SyscallN(addr_QTextEdit_setAlignment, this.h, uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_setPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) SetHtml(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_setHtml, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) SetMarkdown(markdown string) {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	purego.SyscallN(addr_QTextEdit_setMarkdown, this.h, markdown_p, uintptr(len(markdown)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(markdown)
}

func (this *QTextEdit) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) Cut() {
	purego.SyscallN(addr_QTextEdit_cut, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Copy() {
	purego.SyscallN(addr_QTextEdit_copy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Paste() {
	purego.SyscallN(addr_QTextEdit_paste, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Undo() {
	purego.SyscallN(addr_QTextEdit_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Redo() {
	purego.SyscallN(addr_QTextEdit_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Clear() {
	purego.SyscallN(addr_QTextEdit_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SelectAll() {
	purego.SyscallN(addr_QTextEdit_selectAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) InsertPlainText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_insertPlainText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) InsertHtml(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_insertHtml, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) Append(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextEdit_append, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextEdit) ScrollToAnchor(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTextEdit_scrollToAnchor, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTextEdit) ZoomIn() {
	purego.SyscallN(addr_QTextEdit_zoomIn, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) ZoomOut() {
	purego.SyscallN(addr_QTextEdit_zoomOut, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) TextChanged() {
	purego.SyscallN(addr_QTextEdit_textChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) UndoAvailable(b bool) {
	purego.SyscallN(addr_QTextEdit_undoAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) RedoAvailable(b bool) {
	purego.SyscallN(addr_QTextEdit_redoAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) CurrentCharFormatChanged(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextEdit_currentCharFormatChanged, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextEdit) CopyAvailable(b bool) {
	purego.SyscallN(addr_QTextEdit_copyAvailable, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SelectionChanged() {
	purego.SyscallN(addr_QTextEdit_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) CursorPositionChanged() {
	purego.SyscallN(addr_QTextEdit_cursorPositionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) ZoomInF(range_ float32) {
	pg_QTextEdit_zoomInF(this.h, range_)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTextEdit_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTextEdit_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTextEdit) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTextEdit) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTextEdit_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTextEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTextEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTextEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QTextEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QTextEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTextEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QTextEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Create() {
	purego.SyscallN(addr_QTextEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Destroy() {
	purego.SyscallN(addr_QTextEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTextEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTextEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTextEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTextEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QTextEdit_devicePixelRatioFScale()
	return __r
}

func QTextEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTextEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTextEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTextEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTextEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTextEdit) Find2(exp string, options QTextDocument__FindFlag) bool {
	var exp_p uintptr
	if len(exp) > 0 {
		exp_p = uintptr(unsafe.Pointer(unsafe.StringData(exp)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_find3, this.h, exp_p, uintptr(len(exp)), uintptr(options))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QTextEdit) Find3(exp *QRegularExpression, options QTextDocument__FindFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_find4, this.h, uintptr(exp.UnsafePointer()), uintptr(options))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exp)
	return __r != 0
}

func (this *QTextEdit) ToMarkdownWithFeatures(features QTextDocument__MarkdownFeature) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextEdit_toMarkdown2, this.h, uintptr(features), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextEdit) MoveCursor2(operation QTextCursor__MoveOperation, mode QTextCursor__MoveMode) {
	purego.SyscallN(addr_QTextEdit_moveCursor2, this.h, uintptr(operation), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) ZoomInWithRange(range_ int) {
	purego.SyscallN(addr_QTextEdit_zoomIn2, this.h, uintptr(range_))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) ZoomOutWithRange(range_ int) {
	purego.SyscallN(addr_QTextEdit_zoomOut2, this.h, uintptr(range_))
	runtime.KeepAlive(this)
}

func QTextEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTextEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTextEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTextEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTextEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTextEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTextEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTextEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTextEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTextEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextEdit) LoadResourceDefault(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_loadResource_default, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QTextEdit) InputMethodQueryDefault(property InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_inputMethodQuery_default, this.h, uintptr(property))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextEdit) InputMethodQuery2Default(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_inputMethodQuery_default2, this.h, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QTextEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTextEdit) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QTextEdit) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QTextEdit_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QTextEdit_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QTextEdit_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QTextEdit_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QTextEdit_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QTextEdit_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QTextEdit_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QTextEdit_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) MouseDoubleClickEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QTextEdit_mouseDoubleClickEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) ContextMenuEventDefault(e *QContextMenuEvent) {
	purego.SyscallN(addr_QTextEdit_contextMenuEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) DragEnterEventDefault(e *QDragEnterEvent) {
	purego.SyscallN(addr_QTextEdit_dragEnterEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QTextEdit_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QTextEdit_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) DropEventDefault(e *QDropEvent) {
	purego.SyscallN(addr_QTextEdit_dropEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QTextEdit_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QTextEdit_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QTextEdit_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextEdit) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QTextEdit_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QTextEdit_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextEdit) CreateMimeDataFromSelectionDefault() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_createMimeDataFromSelection_default, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QTextEdit) CanInsertFromMimeDataDefault(source *QMimeData) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_canInsertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
	return __r != 0
}

func (this *QTextEdit) InsertFromMimeDataDefault(source *QMimeData) {
	purego.SyscallN(addr_QTextEdit_insertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
}

func (this *QTextEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QTextEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextEdit) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTextEdit_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) DoSetTextCursorDefault(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextEdit_doSetTextCursor_default, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextEdit) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTextEdit_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTextEdit) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QTextEdit) ViewportEventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_viewportEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTextEdit) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextEdit) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTextEdit_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTextEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTextEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTextEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTextEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTextEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTextEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTextEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTextEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTextEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTextEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTextEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTextEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTextEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTextEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTextEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTextEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTextEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTextEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTextEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextEdit) OnTextChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextEdit_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnUndoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextEdit_connect_undoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnRedoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextEdit_connect_redoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnCurrentCharFormatChanged(fn func(*QTextCharFormat)) int64 {
	id := registerCallback(func(format uintptr) { fn(UnsafeNewQTextCharFormat(unsafe.Pointer(format))) })
	pg_QTextEdit_connect_currentCharFormatChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnCopyAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextEdit_connect_copyAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextEdit_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnCursorPositionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextEdit_connect_cursorPositionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTextEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTextEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTextEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTextEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextEdit) Disconnect(handle int64) {
	pg_QTextEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QTextEdit_virtIdx(slot int64) int32 {
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

func QTextEdit_enable(h uintptr, idx int32) { pg_QTextEdit__set_override(h, idx) }

func (this *QTextEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTextEdit) OnLoadResource(fn func(int, *QUrl) *QVariant) {
	registerVirt(this.UnsafeGoID(), -519616681208050567, func(type_ int32, name uintptr) uintptr { return uintptr(fn(int(type_), UnsafeNewQUrl(unsafe.Pointer(name))).UnsafePointer()) }) // slot "loadResource"
}

func (this *QTextEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(property int32) uintptr { return uintptr(fn(InputMethodQuery(property)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTextEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QTextEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QTextEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QTextEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QTextEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QTextEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QTextEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QTextEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QTextEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QTextEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTextEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTextEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(e uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(e))) }) // slot "contextMenuEvent"
}

func (this *QTextEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(e uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(e))) }) // slot "dragEnterEvent"
}

func (this *QTextEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QTextEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QTextEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(e uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(e))) }) // slot "dropEvent"
}

func (this *QTextEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QTextEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QTextEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QTextEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QTextEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QTextEdit) OnCreateMimeDataFromSelection(fn func() *QMimeData) {
	registerVirt(this.UnsafeGoID(), 3493558106800165317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "createMimeDataFromSelection"
}

func (this *QTextEdit) OnCanInsertFromMimeData(fn func(*QMimeData) bool) {
	registerVirt(this.UnsafeGoID(), -2605205276815256896, func(source uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(source)))) }) // slot "canInsertFromMimeData"
}

func (this *QTextEdit) OnInsertFromMimeData(fn func(*QMimeData)) {
	registerVirt(this.UnsafeGoID(), -466553387102125056, func(source uintptr) { fn(UnsafeNewQMimeData(unsafe.Pointer(source))) }) // slot "insertFromMimeData"
}

func (this *QTextEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QTextEdit) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTextEdit) OnDoSetTextCursor(fn func(*QTextCursor)) {
	registerVirt(this.UnsafeGoID(), 4090661491944147209, func(cursor uintptr) { fn(UnsafeNewQTextCursor(unsafe.Pointer(cursor))) }) // slot "doSetTextCursor"
}

func (this *QTextEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTextEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTextEdit) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTextEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QTextEdit) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "viewportEvent"
}

func (this *QTextEdit) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTextEdit) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTextEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTextEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTextEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTextEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTextEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTextEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTextEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTextEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTextEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTextEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTextEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTextEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTextEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTextEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTextEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTextEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTextEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTextEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTextEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTextEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTextEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTextEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
