package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLineEdit_chain uintptr
	addr_QLineEdit_new uintptr
	addr_QLineEdit_new2 uintptr
	addr_QLineEdit_new3 uintptr
	addr_QLineEdit_new4 uintptr
	addr_QLineEdit_tr uintptr
	addr_QLineEdit_text uintptr
	addr_QLineEdit_displayText uintptr
	addr_QLineEdit_placeholderText uintptr
	addr_QLineEdit_setPlaceholderText uintptr
	addr_QLineEdit_maxLength uintptr
	addr_QLineEdit_setMaxLength uintptr
	addr_QLineEdit_setFrame uintptr
	addr_QLineEdit_hasFrame uintptr
	addr_QLineEdit_setClearButtonEnabled uintptr
	addr_QLineEdit_isClearButtonEnabled uintptr
	addr_QLineEdit_echoMode uintptr
	addr_QLineEdit_setEchoMode uintptr
	addr_QLineEdit_isReadOnly uintptr
	addr_QLineEdit_setReadOnly uintptr
	addr_QLineEdit_setValidator uintptr
	addr_QLineEdit_validator uintptr
	addr_QLineEdit_setCompleter uintptr
	addr_QLineEdit_completer uintptr
	addr_QLineEdit_cursorPosition uintptr
	addr_QLineEdit_setCursorPosition uintptr
	addr_QLineEdit_cursorPositionAt uintptr
	addr_QLineEdit_setAlignment uintptr
	addr_QLineEdit_alignment uintptr
	addr_QLineEdit_cursorForward uintptr
	addr_QLineEdit_cursorBackward uintptr
	addr_QLineEdit_cursorWordForward uintptr
	addr_QLineEdit_cursorWordBackward uintptr
	addr_QLineEdit_backspace uintptr
	addr_QLineEdit_del uintptr
	addr_QLineEdit_home uintptr
	addr_QLineEdit_end uintptr
	addr_QLineEdit_isModified uintptr
	addr_QLineEdit_setModified uintptr
	addr_QLineEdit_setSelection uintptr
	addr_QLineEdit_hasSelectedText uintptr
	addr_QLineEdit_selectedText uintptr
	addr_QLineEdit_selectionStart uintptr
	addr_QLineEdit_selectionEnd uintptr
	addr_QLineEdit_selectionLength uintptr
	addr_QLineEdit_isUndoAvailable uintptr
	addr_QLineEdit_isRedoAvailable uintptr
	addr_QLineEdit_setDragEnabled uintptr
	addr_QLineEdit_dragEnabled uintptr
	addr_QLineEdit_setCursorMoveStyle uintptr
	addr_QLineEdit_cursorMoveStyle uintptr
	addr_QLineEdit_inputMask uintptr
	addr_QLineEdit_setInputMask uintptr
	addr_QLineEdit_hasAcceptableInput uintptr
	addr_QLineEdit_setTextMargins uintptr
	addr_QLineEdit_setTextMargins2 uintptr
	addr_QLineEdit_textMargins uintptr
	addr_QLineEdit_addAction uintptr
	addr_QLineEdit_addAction2 uintptr
	addr_QLineEdit_setText uintptr
	addr_QLineEdit_clear uintptr
	addr_QLineEdit_selectAll uintptr
	addr_QLineEdit_undo uintptr
	addr_QLineEdit_redo uintptr
	addr_QLineEdit_cut uintptr
	addr_QLineEdit_copy uintptr
	addr_QLineEdit_paste uintptr
	addr_QLineEdit_deselect uintptr
	addr_QLineEdit_insert uintptr
	addr_QLineEdit_createStandardContextMenu uintptr
	addr_QLineEdit_textChanged uintptr
	addr_QLineEdit_textEdited uintptr
	addr_QLineEdit_cursorPositionChanged uintptr
	addr_QLineEdit_returnPressed uintptr
	addr_QLineEdit_editingFinished uintptr
	addr_QLineEdit_selectionChanged uintptr
	addr_QLineEdit_inputRejected uintptr
	addr_QLineEdit_inputMethodQuery uintptr
	addr_QLineEdit_inputMethodQuery2 uintptr
	addr_QLineEdit_cursorRect uintptr
	addr_QLineEdit_setTabOrder uintptr
	addr_QLineEdit_mouseGrabber uintptr
	addr_QLineEdit_keyboardGrabber uintptr
	addr_QLineEdit_find uintptr
	addr_QLineEdit_createWindowContainer uintptr
	addr_QLineEdit_updateMicroFocus uintptr
	addr_QLineEdit_create uintptr
	addr_QLineEdit_destroy uintptr
	addr_QLineEdit_focusNextChild uintptr
	addr_QLineEdit_focusPreviousChild uintptr
	addr_QLineEdit_connect uintptr
	addr_QLineEdit_connect2 uintptr
	addr_QLineEdit_sender uintptr
	addr_QLineEdit_senderSignalIndex uintptr
	addr_QLineEdit_receivers uintptr
	addr_QLineEdit_isSignalConnected uintptr
	pg_QLineEdit_devicePixelRatioFScale func() float64
	pg_QLineEdit_encodeMetricF func(int32, float64) int32
	pg_QLineEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QLineEdit_tr2 uintptr
	addr_QLineEdit_tr3 uintptr
	addr_QLineEdit_cursorForward2 uintptr
	addr_QLineEdit_cursorBackward2 uintptr
	addr_QLineEdit_createWindowContainer2 uintptr
	addr_QLineEdit_createWindowContainer3 uintptr
	addr_QLineEdit_updateMicroFocus2 uintptr
	addr_QLineEdit_create2 uintptr
	addr_QLineEdit_create3 uintptr
	addr_QLineEdit_create4 uintptr
	addr_QLineEdit_destroy2 uintptr
	addr_QLineEdit_destroy3 uintptr
	addr_QLineEdit_connect3 uintptr
	addr_QLineEdit_connect4 uintptr
	addr_QLineEdit_inputMethodQuery_default uintptr
	addr_QLineEdit_inputMethodQuery_default2 uintptr
	addr_QLineEdit_metaObject_default uintptr
	addr_QLineEdit_sizeHint_default uintptr
	addr_QLineEdit_minimumSizeHint_default uintptr
	addr_QLineEdit_mousePressEvent_default uintptr
	addr_QLineEdit_mouseMoveEvent_default uintptr
	addr_QLineEdit_mouseReleaseEvent_default uintptr
	addr_QLineEdit_mouseDoubleClickEvent_default uintptr
	addr_QLineEdit_keyPressEvent_default uintptr
	addr_QLineEdit_keyReleaseEvent_default uintptr
	addr_QLineEdit_focusInEvent_default uintptr
	addr_QLineEdit_focusOutEvent_default uintptr
	addr_QLineEdit_paintEvent_default uintptr
	addr_QLineEdit_dragEnterEvent_default uintptr
	addr_QLineEdit_dragMoveEvent_default uintptr
	addr_QLineEdit_dragLeaveEvent_default uintptr
	addr_QLineEdit_dropEvent_default uintptr
	addr_QLineEdit_changeEvent_default uintptr
	addr_QLineEdit_contextMenuEvent_default uintptr
	addr_QLineEdit_inputMethodEvent_default uintptr
	addr_QLineEdit_initStyleOption_default uintptr
	addr_QLineEdit_timerEvent_default uintptr
	addr_QLineEdit_event_default uintptr
	addr_QLineEdit_devType_default uintptr
	addr_QLineEdit_setVisible_default uintptr
	addr_QLineEdit_heightForWidth_default uintptr
	addr_QLineEdit_hasHeightForWidth_default uintptr
	addr_QLineEdit_paintEngine_default uintptr
	addr_QLineEdit_wheelEvent_default uintptr
	addr_QLineEdit_enterEvent_default uintptr
	addr_QLineEdit_leaveEvent_default uintptr
	addr_QLineEdit_moveEvent_default uintptr
	addr_QLineEdit_resizeEvent_default uintptr
	addr_QLineEdit_closeEvent_default uintptr
	addr_QLineEdit_tabletEvent_default uintptr
	addr_QLineEdit_actionEvent_default uintptr
	addr_QLineEdit_showEvent_default uintptr
	addr_QLineEdit_hideEvent_default uintptr
	addr_QLineEdit_metric_default uintptr
	addr_QLineEdit_initPainter_default uintptr
	addr_QLineEdit_redirected_default uintptr
	addr_QLineEdit_sharedPainter_default uintptr
	addr_QLineEdit_focusNextPrevChild_default uintptr
	addr_QLineEdit_eventFilter_default uintptr
	addr_QLineEdit_childEvent_default uintptr
	addr_QLineEdit_customEvent_default uintptr
	addr_QLineEdit_connectNotify_default uintptr
	addr_QLineEdit_disconnectNotify_default uintptr
	pg_QLineEdit_connect_textChanged func(uintptr, int64)
	pg_QLineEdit_connect_textEdited func(uintptr, int64)
	pg_QLineEdit_connect_cursorPositionChanged func(uintptr, int64)
	pg_QLineEdit_connect_returnPressed func(uintptr, int64)
	pg_QLineEdit_connect_editingFinished func(uintptr, int64)
	pg_QLineEdit_connect_selectionChanged func(uintptr, int64)
	pg_QLineEdit_connect_inputRejected func(uintptr, int64)
	pg_QLineEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QLineEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QLineEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QLineEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QLineEdit_connect_destroyed func(uintptr, int64)
	pg_QLineEdit_connect_destroyed2 func(uintptr, int64)
	pg_QLineEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QLineEdit__disconnect func(int64)
	pg_QLineEdit__set_override func(uintptr, int32)
	pg_QLineEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLineEdit_chain, "QLineEdit_chain")
	symaddr(&addr_QLineEdit_new, "QLineEdit_new")
	symaddr(&addr_QLineEdit_new2, "QLineEdit_new2")
	symaddr(&addr_QLineEdit_new3, "QLineEdit_new3")
	symaddr(&addr_QLineEdit_new4, "QLineEdit_new4")
	symaddr(&addr_QLineEdit_tr, "QLineEdit_tr")
	symaddr(&addr_QLineEdit_text, "QLineEdit_text")
	symaddr(&addr_QLineEdit_displayText, "QLineEdit_displayText")
	symaddr(&addr_QLineEdit_placeholderText, "QLineEdit_placeholderText")
	symaddr(&addr_QLineEdit_setPlaceholderText, "QLineEdit_setPlaceholderText")
	symaddr(&addr_QLineEdit_maxLength, "QLineEdit_maxLength")
	symaddr(&addr_QLineEdit_setMaxLength, "QLineEdit_setMaxLength")
	symaddr(&addr_QLineEdit_setFrame, "QLineEdit_setFrame")
	symaddr(&addr_QLineEdit_hasFrame, "QLineEdit_hasFrame")
	symaddr(&addr_QLineEdit_setClearButtonEnabled, "QLineEdit_setClearButtonEnabled")
	symaddr(&addr_QLineEdit_isClearButtonEnabled, "QLineEdit_isClearButtonEnabled")
	symaddr(&addr_QLineEdit_echoMode, "QLineEdit_echoMode")
	symaddr(&addr_QLineEdit_setEchoMode, "QLineEdit_setEchoMode")
	symaddr(&addr_QLineEdit_isReadOnly, "QLineEdit_isReadOnly")
	symaddr(&addr_QLineEdit_setReadOnly, "QLineEdit_setReadOnly")
	symaddr(&addr_QLineEdit_setValidator, "QLineEdit_setValidator")
	symaddr(&addr_QLineEdit_validator, "QLineEdit_validator")
	symaddr(&addr_QLineEdit_setCompleter, "QLineEdit_setCompleter")
	symaddr(&addr_QLineEdit_completer, "QLineEdit_completer")
	symaddr(&addr_QLineEdit_cursorPosition, "QLineEdit_cursorPosition")
	symaddr(&addr_QLineEdit_setCursorPosition, "QLineEdit_setCursorPosition")
	symaddr(&addr_QLineEdit_cursorPositionAt, "QLineEdit_cursorPositionAt")
	symaddr(&addr_QLineEdit_setAlignment, "QLineEdit_setAlignment")
	symaddr(&addr_QLineEdit_alignment, "QLineEdit_alignment")
	symaddr(&addr_QLineEdit_cursorForward, "QLineEdit_cursorForward")
	symaddr(&addr_QLineEdit_cursorBackward, "QLineEdit_cursorBackward")
	symaddr(&addr_QLineEdit_cursorWordForward, "QLineEdit_cursorWordForward")
	symaddr(&addr_QLineEdit_cursorWordBackward, "QLineEdit_cursorWordBackward")
	symaddr(&addr_QLineEdit_backspace, "QLineEdit_backspace")
	symaddr(&addr_QLineEdit_del, "QLineEdit_del")
	symaddr(&addr_QLineEdit_home, "QLineEdit_home")
	symaddr(&addr_QLineEdit_end, "QLineEdit_end")
	symaddr(&addr_QLineEdit_isModified, "QLineEdit_isModified")
	symaddr(&addr_QLineEdit_setModified, "QLineEdit_setModified")
	symaddr(&addr_QLineEdit_setSelection, "QLineEdit_setSelection")
	symaddr(&addr_QLineEdit_hasSelectedText, "QLineEdit_hasSelectedText")
	symaddr(&addr_QLineEdit_selectedText, "QLineEdit_selectedText")
	symaddr(&addr_QLineEdit_selectionStart, "QLineEdit_selectionStart")
	symaddr(&addr_QLineEdit_selectionEnd, "QLineEdit_selectionEnd")
	symaddr(&addr_QLineEdit_selectionLength, "QLineEdit_selectionLength")
	symaddr(&addr_QLineEdit_isUndoAvailable, "QLineEdit_isUndoAvailable")
	symaddr(&addr_QLineEdit_isRedoAvailable, "QLineEdit_isRedoAvailable")
	symaddr(&addr_QLineEdit_setDragEnabled, "QLineEdit_setDragEnabled")
	symaddr(&addr_QLineEdit_dragEnabled, "QLineEdit_dragEnabled")
	symaddr(&addr_QLineEdit_setCursorMoveStyle, "QLineEdit_setCursorMoveStyle")
	symaddr(&addr_QLineEdit_cursorMoveStyle, "QLineEdit_cursorMoveStyle")
	symaddr(&addr_QLineEdit_inputMask, "QLineEdit_inputMask")
	symaddr(&addr_QLineEdit_setInputMask, "QLineEdit_setInputMask")
	symaddr(&addr_QLineEdit_hasAcceptableInput, "QLineEdit_hasAcceptableInput")
	symaddr(&addr_QLineEdit_setTextMargins, "QLineEdit_setTextMargins")
	symaddr(&addr_QLineEdit_setTextMargins2, "QLineEdit_setTextMargins2")
	symaddr(&addr_QLineEdit_textMargins, "QLineEdit_textMargins")
	symaddr(&addr_QLineEdit_addAction, "QLineEdit_addAction")
	symaddr(&addr_QLineEdit_addAction2, "QLineEdit_addAction2")
	symaddr(&addr_QLineEdit_setText, "QLineEdit_setText")
	symaddr(&addr_QLineEdit_clear, "QLineEdit_clear")
	symaddr(&addr_QLineEdit_selectAll, "QLineEdit_selectAll")
	symaddr(&addr_QLineEdit_undo, "QLineEdit_undo")
	symaddr(&addr_QLineEdit_redo, "QLineEdit_redo")
	symaddr(&addr_QLineEdit_cut, "QLineEdit_cut")
	symaddr(&addr_QLineEdit_copy, "QLineEdit_copy")
	symaddr(&addr_QLineEdit_paste, "QLineEdit_paste")
	symaddr(&addr_QLineEdit_deselect, "QLineEdit_deselect")
	symaddr(&addr_QLineEdit_insert, "QLineEdit_insert")
	symaddr(&addr_QLineEdit_createStandardContextMenu, "QLineEdit_createStandardContextMenu")
	symaddr(&addr_QLineEdit_textChanged, "QLineEdit_textChanged")
	symaddr(&addr_QLineEdit_textEdited, "QLineEdit_textEdited")
	symaddr(&addr_QLineEdit_cursorPositionChanged, "QLineEdit_cursorPositionChanged")
	symaddr(&addr_QLineEdit_returnPressed, "QLineEdit_returnPressed")
	symaddr(&addr_QLineEdit_editingFinished, "QLineEdit_editingFinished")
	symaddr(&addr_QLineEdit_selectionChanged, "QLineEdit_selectionChanged")
	symaddr(&addr_QLineEdit_inputRejected, "QLineEdit_inputRejected")
	symaddr(&addr_QLineEdit_inputMethodQuery, "QLineEdit_inputMethodQuery")
	symaddr(&addr_QLineEdit_inputMethodQuery2, "QLineEdit_inputMethodQuery2")
	symaddr(&addr_QLineEdit_cursorRect, "QLineEdit_cursorRect")
	symaddr(&addr_QLineEdit_setTabOrder, "QLineEdit_setTabOrder")
	symaddr(&addr_QLineEdit_mouseGrabber, "QLineEdit_mouseGrabber")
	symaddr(&addr_QLineEdit_keyboardGrabber, "QLineEdit_keyboardGrabber")
	symaddr(&addr_QLineEdit_find, "QLineEdit_find")
	symaddr(&addr_QLineEdit_createWindowContainer, "QLineEdit_createWindowContainer")
	symaddr(&addr_QLineEdit_updateMicroFocus, "QLineEdit_updateMicroFocus")
	symaddr(&addr_QLineEdit_create, "QLineEdit_create")
	symaddr(&addr_QLineEdit_destroy, "QLineEdit_destroy")
	symaddr(&addr_QLineEdit_focusNextChild, "QLineEdit_focusNextChild")
	symaddr(&addr_QLineEdit_focusPreviousChild, "QLineEdit_focusPreviousChild")
	symaddr(&addr_QLineEdit_connect, "QLineEdit_connect")
	symaddr(&addr_QLineEdit_connect2, "QLineEdit_connect2")
	symaddr(&addr_QLineEdit_sender, "QLineEdit_sender")
	symaddr(&addr_QLineEdit_senderSignalIndex, "QLineEdit_senderSignalIndex")
	symaddr(&addr_QLineEdit_receivers, "QLineEdit_receivers")
	symaddr(&addr_QLineEdit_isSignalConnected, "QLineEdit_isSignalConnected")
	reg(&pg_QLineEdit_devicePixelRatioFScale, "QLineEdit_devicePixelRatioFScale")
	reg(&pg_QLineEdit_encodeMetricF, "QLineEdit_encodeMetricF")
	reg(&pg_QLineEdit_getDecodedMetricF, "QLineEdit_getDecodedMetricF")
	symaddr(&addr_QLineEdit_tr2, "QLineEdit_tr2")
	symaddr(&addr_QLineEdit_tr3, "QLineEdit_tr3")
	symaddr(&addr_QLineEdit_cursorForward2, "QLineEdit_cursorForward2")
	symaddr(&addr_QLineEdit_cursorBackward2, "QLineEdit_cursorBackward2")
	symaddr(&addr_QLineEdit_createWindowContainer2, "QLineEdit_createWindowContainer2")
	symaddr(&addr_QLineEdit_createWindowContainer3, "QLineEdit_createWindowContainer3")
	symaddr(&addr_QLineEdit_updateMicroFocus2, "QLineEdit_updateMicroFocus2")
	symaddr(&addr_QLineEdit_create2, "QLineEdit_create2")
	symaddr(&addr_QLineEdit_create3, "QLineEdit_create3")
	symaddr(&addr_QLineEdit_create4, "QLineEdit_create4")
	symaddr(&addr_QLineEdit_destroy2, "QLineEdit_destroy2")
	symaddr(&addr_QLineEdit_destroy3, "QLineEdit_destroy3")
	symaddr(&addr_QLineEdit_connect3, "QLineEdit_connect3")
	symaddr(&addr_QLineEdit_connect4, "QLineEdit_connect4")
	symaddr(&addr_QLineEdit_inputMethodQuery_default, "QLineEdit_inputMethodQuery_default")
	symaddr(&addr_QLineEdit_inputMethodQuery_default2, "QLineEdit_inputMethodQuery_default2")
	symaddr(&addr_QLineEdit_metaObject_default, "QLineEdit_metaObject_default")
	symaddr(&addr_QLineEdit_sizeHint_default, "QLineEdit_sizeHint_default")
	symaddr(&addr_QLineEdit_minimumSizeHint_default, "QLineEdit_minimumSizeHint_default")
	symaddr(&addr_QLineEdit_mousePressEvent_default, "QLineEdit_mousePressEvent_default")
	symaddr(&addr_QLineEdit_mouseMoveEvent_default, "QLineEdit_mouseMoveEvent_default")
	symaddr(&addr_QLineEdit_mouseReleaseEvent_default, "QLineEdit_mouseReleaseEvent_default")
	symaddr(&addr_QLineEdit_mouseDoubleClickEvent_default, "QLineEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QLineEdit_keyPressEvent_default, "QLineEdit_keyPressEvent_default")
	symaddr(&addr_QLineEdit_keyReleaseEvent_default, "QLineEdit_keyReleaseEvent_default")
	symaddr(&addr_QLineEdit_focusInEvent_default, "QLineEdit_focusInEvent_default")
	symaddr(&addr_QLineEdit_focusOutEvent_default, "QLineEdit_focusOutEvent_default")
	symaddr(&addr_QLineEdit_paintEvent_default, "QLineEdit_paintEvent_default")
	symaddr(&addr_QLineEdit_dragEnterEvent_default, "QLineEdit_dragEnterEvent_default")
	symaddr(&addr_QLineEdit_dragMoveEvent_default, "QLineEdit_dragMoveEvent_default")
	symaddr(&addr_QLineEdit_dragLeaveEvent_default, "QLineEdit_dragLeaveEvent_default")
	symaddr(&addr_QLineEdit_dropEvent_default, "QLineEdit_dropEvent_default")
	symaddr(&addr_QLineEdit_changeEvent_default, "QLineEdit_changeEvent_default")
	symaddr(&addr_QLineEdit_contextMenuEvent_default, "QLineEdit_contextMenuEvent_default")
	symaddr(&addr_QLineEdit_inputMethodEvent_default, "QLineEdit_inputMethodEvent_default")
	symaddr(&addr_QLineEdit_initStyleOption_default, "QLineEdit_initStyleOption_default")
	symaddr(&addr_QLineEdit_timerEvent_default, "QLineEdit_timerEvent_default")
	symaddr(&addr_QLineEdit_event_default, "QLineEdit_event_default")
	symaddr(&addr_QLineEdit_devType_default, "QLineEdit_devType_default")
	symaddr(&addr_QLineEdit_setVisible_default, "QLineEdit_setVisible_default")
	symaddr(&addr_QLineEdit_heightForWidth_default, "QLineEdit_heightForWidth_default")
	symaddr(&addr_QLineEdit_hasHeightForWidth_default, "QLineEdit_hasHeightForWidth_default")
	symaddr(&addr_QLineEdit_paintEngine_default, "QLineEdit_paintEngine_default")
	symaddr(&addr_QLineEdit_wheelEvent_default, "QLineEdit_wheelEvent_default")
	symaddr(&addr_QLineEdit_enterEvent_default, "QLineEdit_enterEvent_default")
	symaddr(&addr_QLineEdit_leaveEvent_default, "QLineEdit_leaveEvent_default")
	symaddr(&addr_QLineEdit_moveEvent_default, "QLineEdit_moveEvent_default")
	symaddr(&addr_QLineEdit_resizeEvent_default, "QLineEdit_resizeEvent_default")
	symaddr(&addr_QLineEdit_closeEvent_default, "QLineEdit_closeEvent_default")
	symaddr(&addr_QLineEdit_tabletEvent_default, "QLineEdit_tabletEvent_default")
	symaddr(&addr_QLineEdit_actionEvent_default, "QLineEdit_actionEvent_default")
	symaddr(&addr_QLineEdit_showEvent_default, "QLineEdit_showEvent_default")
	symaddr(&addr_QLineEdit_hideEvent_default, "QLineEdit_hideEvent_default")
	symaddr(&addr_QLineEdit_metric_default, "QLineEdit_metric_default")
	symaddr(&addr_QLineEdit_initPainter_default, "QLineEdit_initPainter_default")
	symaddr(&addr_QLineEdit_redirected_default, "QLineEdit_redirected_default")
	symaddr(&addr_QLineEdit_sharedPainter_default, "QLineEdit_sharedPainter_default")
	symaddr(&addr_QLineEdit_focusNextPrevChild_default, "QLineEdit_focusNextPrevChild_default")
	symaddr(&addr_QLineEdit_eventFilter_default, "QLineEdit_eventFilter_default")
	symaddr(&addr_QLineEdit_childEvent_default, "QLineEdit_childEvent_default")
	symaddr(&addr_QLineEdit_customEvent_default, "QLineEdit_customEvent_default")
	symaddr(&addr_QLineEdit_connectNotify_default, "QLineEdit_connectNotify_default")
	symaddr(&addr_QLineEdit_disconnectNotify_default, "QLineEdit_disconnectNotify_default")
	reg(&pg_QLineEdit_connect_textChanged, "QLineEdit_connect_textChanged")
	reg(&pg_QLineEdit_connect_textEdited, "QLineEdit_connect_textEdited")
	reg(&pg_QLineEdit_connect_cursorPositionChanged, "QLineEdit_connect_cursorPositionChanged")
	reg(&pg_QLineEdit_connect_returnPressed, "QLineEdit_connect_returnPressed")
	reg(&pg_QLineEdit_connect_editingFinished, "QLineEdit_connect_editingFinished")
	reg(&pg_QLineEdit_connect_selectionChanged, "QLineEdit_connect_selectionChanged")
	reg(&pg_QLineEdit_connect_inputRejected, "QLineEdit_connect_inputRejected")
	reg(&pg_QLineEdit_connect_windowTitleChanged, "QLineEdit_connect_windowTitleChanged")
	reg(&pg_QLineEdit_connect_windowIconChanged, "QLineEdit_connect_windowIconChanged")
	reg(&pg_QLineEdit_connect_windowIconTextChanged, "QLineEdit_connect_windowIconTextChanged")
	reg(&pg_QLineEdit_connect_customContextMenuRequested, "QLineEdit_connect_customContextMenuRequested")
	reg(&pg_QLineEdit_connect_destroyed, "QLineEdit_connect_destroyed")
	reg(&pg_QLineEdit_connect_destroyed2, "QLineEdit_connect_destroyed2")
	reg(&pg_QLineEdit_connect_objectNameChanged, "QLineEdit_connect_objectNameChanged")
	reg(&pg_QLineEdit__disconnect, "QLineEdit__disconnect")
	reg(&pg_QLineEdit__set_override, "QLineEdit__set_override")
	reg(&pg_QLineEdit_delete, "QLineEdit_delete")
}

type QLineEdit struct {
	h uintptr
	*QWidget
}

func bindQLineEdit(h uintptr) *QLineEdit {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QLineEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QLineEdit{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQLineEdit(h uintptr) *QLineEdit { return bindQLineEdit(h) }

func (this *QLineEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLineEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLineEdit(h unsafe.Pointer) *QLineEdit { return bindQLineEdit(uintptr(h)) }
func UnsafeNewQLineEditOwned(h unsafe.Pointer) *QLineEdit { return newownedQLineEdit(uintptr(h)) }

func (this *QLineEdit) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQLineEdit(parent *QWidget) *QLineEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLineEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQLineEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLineEdit_virtIdx, QLineEdit_enable)
	return __o
}

func NewQLineEdit2() *QLineEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLineEdit_new2, uintptr(__goid))
	__o := newownedQLineEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLineEdit_virtIdx, QLineEdit_enable)
	return __o
}

func NewQLineEdit3(arg1 string) *QLineEdit {
	__goid := nextVirtID()
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QLineEdit_new3, uintptr(__goid), arg1_p, uintptr(len(arg1)))
	__o := newownedQLineEdit(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLineEdit_virtIdx, QLineEdit_enable)
	return __o
}

func NewQLineEdit4(arg1 string, parent *QWidget) *QLineEdit {
	__goid := nextVirtID()
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QLineEdit_new4, uintptr(__goid), arg1_p, uintptr(len(arg1)), uintptr(parent.UnsafePointer()))
	__o := newownedQLineEdit(__r)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLineEdit_virtIdx, QLineEdit_enable)
	return __o
}

func QLineEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QLineEdit) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLineEdit) DisplayText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_displayText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLineEdit) PlaceholderText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_placeholderText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLineEdit) SetPlaceholderText(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLineEdit_setPlaceholderText, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) MaxLength() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_maxLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) SetMaxLength(arg1 int) {
	purego.SyscallN(addr_QLineEdit_setMaxLength, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SetFrame(arg1 bool) {
	purego.SyscallN(addr_QLineEdit_setFrame, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) HasFrame() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_hasFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetClearButtonEnabled(enable bool) {
	purego.SyscallN(addr_QLineEdit_setClearButtonEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) IsClearButtonEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isClearButtonEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) EchoMode() QLineEdit__EchoMode {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_echoMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLineEdit__EchoMode(__r)
}

func (this *QLineEdit) SetEchoMode(arg1 QLineEdit__EchoMode) {
	purego.SyscallN(addr_QLineEdit_setEchoMode, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetReadOnly(arg1 bool) {
	purego.SyscallN(addr_QLineEdit_setReadOnly, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SetValidator(arg1 *QValidator) {
	purego.SyscallN(addr_QLineEdit_setValidator, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) Validator() *QValidator {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_validator, this.h)
	runtime.KeepAlive(this)
	return bindQValidator(__r)
}

func (this *QLineEdit) SetCompleter(completer *QCompleter) {
	purego.SyscallN(addr_QLineEdit_setCompleter, this.h, uintptr(completer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(completer)
}

func (this *QLineEdit) Completer() *QCompleter {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_completer, this.h)
	runtime.KeepAlive(this)
	return bindQCompleter(__r)
}

func (this *QLineEdit) CursorPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_cursorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) SetCursorPosition(arg1 int) {
	purego.SyscallN(addr_QLineEdit_setCursorPosition, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorPositionAt(pos *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_cursorPositionAt, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return int(__r)
}

func (this *QLineEdit) SetAlignment(flag AlignmentFlag) {
	purego.SyscallN(addr_QLineEdit_setAlignment, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QLineEdit) CursorForward(mark bool) {
	purego.SyscallN(addr_QLineEdit_cursorForward, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorBackward(mark bool) {
	purego.SyscallN(addr_QLineEdit_cursorBackward, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorWordForward(mark bool) {
	purego.SyscallN(addr_QLineEdit_cursorWordForward, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorWordBackward(mark bool) {
	purego.SyscallN(addr_QLineEdit_cursorWordBackward, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Backspace() {
	purego.SyscallN(addr_QLineEdit_backspace, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Del() {
	purego.SyscallN(addr_QLineEdit_del, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Home(mark bool) {
	purego.SyscallN(addr_QLineEdit_home, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) End(mark bool) {
	purego.SyscallN(addr_QLineEdit_end, this.h, uintptr(b2i(mark)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) IsModified() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isModified, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetModified(arg1 bool) {
	purego.SyscallN(addr_QLineEdit_setModified, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SetSelection(arg1 int, arg2 int) {
	purego.SyscallN(addr_QLineEdit_setSelection, this.h, uintptr(arg1), uintptr(arg2))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) HasSelectedText() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_hasSelectedText, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SelectedText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_selectedText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLineEdit) SelectionStart() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_selectionStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) SelectionEnd() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_selectionEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) SelectionLength() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_selectionLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) IsUndoAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isUndoAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) IsRedoAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isRedoAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetDragEnabled(b bool) {
	purego.SyscallN(addr_QLineEdit_setDragEnabled, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) DragEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_dragEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetCursorMoveStyle(style CursorMoveStyle) {
	purego.SyscallN(addr_QLineEdit_setCursorMoveStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorMoveStyle() CursorMoveStyle {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_cursorMoveStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CursorMoveStyle(__r)
}

func (this *QLineEdit) InputMask() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_inputMask, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLineEdit) SetInputMask(inputMask string) {
	var inputMask_p uintptr
	if len(inputMask) > 0 {
		inputMask_p = uintptr(unsafe.Pointer(unsafe.StringData(inputMask)))
	}
	purego.SyscallN(addr_QLineEdit_setInputMask, this.h, inputMask_p, uintptr(len(inputMask)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(inputMask)
}

func (this *QLineEdit) HasAcceptableInput() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_hasAcceptableInput, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) SetTextMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QLineEdit_setTextMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SetTextMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QLineEdit_setTextMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QLineEdit) TextMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_textMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QLineEdit) AddAction(action *QAction, position QLineEdit__ActionPosition) {
	purego.SyscallN(addr_QLineEdit_addAction, this.h, uintptr(action.UnsafePointer()), uintptr(position))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QLineEdit) AddAction2(icon *QIcon, position QLineEdit__ActionPosition) *QAction {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_addAction2, this.h, uintptr(icon.UnsafePointer()), uintptr(position))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	return bindQAction(__r)
}

func (this *QLineEdit) SetText(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLineEdit_setText, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) Clear() {
	purego.SyscallN(addr_QLineEdit_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SelectAll() {
	purego.SyscallN(addr_QLineEdit_selectAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Undo() {
	purego.SyscallN(addr_QLineEdit_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Redo() {
	purego.SyscallN(addr_QLineEdit_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Cut() {
	purego.SyscallN(addr_QLineEdit_cut, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Copy() {
	purego.SyscallN(addr_QLineEdit_copy, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Paste() {
	purego.SyscallN(addr_QLineEdit_paste, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Deselect() {
	purego.SyscallN(addr_QLineEdit_deselect, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Insert(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLineEdit_insert, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) CreateStandardContextMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_createStandardContextMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QLineEdit) TextChanged(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLineEdit_textChanged, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) TextEdited(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLineEdit_textEdited, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) CursorPositionChanged(arg1 int, arg2 int) {
	purego.SyscallN(addr_QLineEdit_cursorPositionChanged, this.h, uintptr(arg1), uintptr(arg2))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) ReturnPressed() {
	purego.SyscallN(addr_QLineEdit_returnPressed, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) EditingFinished() {
	purego.SyscallN(addr_QLineEdit_editingFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) SelectionChanged() {
	purego.SyscallN(addr_QLineEdit_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) InputRejected() {
	purego.SyscallN(addr_QLineEdit_inputRejected, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) InputMethodQuery(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_inputMethodQuery, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QLineEdit) InputMethodQuery2(property InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_inputMethodQuery2, this.h, uintptr(property), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QLineEdit) CursorRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_cursorRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func QLineEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QLineEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QLineEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QLineEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QLineEdit_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QLineEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QLineEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QLineEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Create() {
	purego.SyscallN(addr_QLineEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Destroy() {
	purego.SyscallN(addr_QLineEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QLineEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QLineEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLineEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLineEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLineEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QLineEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QLineEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QLineEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QLineEdit_devicePixelRatioFScale()
	return __r
}

func QLineEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QLineEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QLineEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QLineEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QLineEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLineEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLineEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QLineEdit) CursorForward2(mark bool, steps int) {
	purego.SyscallN(addr_QLineEdit_cursorForward2, this.h, uintptr(b2i(mark)), uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CursorBackward2(mark bool, steps int) {
	purego.SyscallN(addr_QLineEdit_cursorBackward2, this.h, uintptr(b2i(mark)), uintptr(steps))
	runtime.KeepAlive(this)
}

func QLineEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QLineEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QLineEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QLineEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QLineEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QLineEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QLineEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QLineEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QLineEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QLineEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLineEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLineEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLineEdit) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QLineEdit) InputMethodQuery2Default(property InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_inputMethodQuery_default2, this.h, uintptr(property), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QLineEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QLineEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLineEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLineEdit) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QLineEdit_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QLineEdit_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QLineEdit_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QLineEdit_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QLineEdit_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) KeyReleaseEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QLineEdit_keyReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QLineEdit_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QLineEdit_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QLineEdit_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) DragEnterEventDefault(arg1 *QDragEnterEvent) {
	purego.SyscallN(addr_QLineEdit_dragEnterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QLineEdit_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QLineEdit) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QLineEdit_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QLineEdit) DropEventDefault(arg1 *QDropEvent) {
	purego.SyscallN(addr_QLineEdit_dropEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QLineEdit_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QLineEdit_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QLineEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QLineEdit_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QLineEdit) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QLineEdit_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLineEdit) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QLineEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QLineEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QLineEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QLineEdit) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QLineEdit_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QLineEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLineEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QLineEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QLineEdit_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QLineEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QLineEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QLineEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QLineEdit_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QLineEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLineEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QLineEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QLineEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QLineEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QLineEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QLineEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineEdit) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineEdit_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLineEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QLineEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLineEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLineEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLineEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLineEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLineEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLineEdit) OnTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QLineEdit_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnTextEdited(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QLineEdit_connect_textEdited(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnCursorPositionChanged(fn func(int, int)) int64 {
	id := registerCallback(func(arg1 int32, arg2 int32) { fn(int(arg1), int(arg2)) })
	pg_QLineEdit_connect_cursorPositionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLineEdit_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLineEdit_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLineEdit_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnInputRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLineEdit_connect_inputRejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QLineEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QLineEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QLineEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QLineEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLineEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QLineEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QLineEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLineEdit) Disconnect(handle int64) {
	pg_QLineEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QLineEdit_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case -332090898078543319:
		return 3
	case -4255494004099806581:
		return 4
	case -4415920049297795213:
		return 5
	case -4029073905697023887:
		return 6
	case 1023791977679793185:
		return 7
	case 772598475950050299:
		return 8
	case 975043767033552552:
		return 9
	case -385596010576048993:
		return 10
	case 5290442163167458465:
		return 11
	case 206122715919917589:
		return 12
	case -1781545137352745978:
		return 13
	case -3139904158373725006:
		return 14
	case 9107058682754591716:
		return 15
	case 4959376063826165943:
		return 16
	case -6698043222227006015:
		return 17
	case 5438819661942843222:
		return 18
	case -488693769771308531:
		return 19
	case 7345853097340709638:
		return 20
	case -5409927581812374498:
		return 21
	case 1237236456312824991:
		return 22
	case -8659679893868524044:
		return 23
	case 2110182719375245575:
		return 24
	case 4311534060495737339:
		return 25
	case -8266762397995423915:
		return 26
	case 6398007404756346381:
		return 27
	case 1919988340934511316:
		return 28
	case 6316270346668237895:
		return 29
	case -551070009623711640:
		return 30
	case -9161949150489525396:
		return 31
	case -4302917285655014661:
		return 32
	case -1989937050019464783:
		return 33
	case -212068866860701141:
		return 34
	case 1697444140088555581:
		return 35
	case -6825369174751916:
		return 36
	case -199718707876719129:
		return 37
	case -2039166586246443706:
		return 38
	case 427629649267311777:
		return 39
	case 6477500296980050210:
		return 40
	case 4127240633310655048:
		return 41
	case 5792771328741774991:
		return 42
	case 5898328517505784191:
		return 43
	case 5217623514630858297:
		return 44
	case -1185315367526857833:
		return 45
	case 3769868229123868474:
		return 46
	case 7379029982170503008:
		return 47
	case -568369772982215962:
		return 48
	}
	return -1
}

func QLineEdit_enable(h uintptr, idx int32) { pg_QLineEdit__set_override(h, idx) }

func (this *QLineEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QLineEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QLineEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QLineEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QLineEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QLineEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QLineEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QLineEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QLineEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyReleaseEvent"
}

func (this *QLineEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QLineEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QLineEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QLineEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(arg1 uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(arg1))) }) // slot "dragEnterEvent"
}

func (this *QLineEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QLineEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QLineEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(arg1 uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(arg1))) }) // slot "dropEvent"
}

func (this *QLineEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QLineEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QLineEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QLineEdit) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QLineEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QLineEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QLineEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QLineEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QLineEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QLineEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QLineEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QLineEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QLineEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QLineEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QLineEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QLineEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QLineEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QLineEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QLineEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QLineEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QLineEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QLineEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QLineEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QLineEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QLineEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QLineEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QLineEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QLineEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QLineEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QLineEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QLineEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QLineEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QLineEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QLineEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLineEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
