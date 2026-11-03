package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextBrowser_chain uintptr
	addr_QTextBrowser_new uintptr
	addr_QTextBrowser_new2 uintptr
	addr_QTextBrowser_tr uintptr
	addr_QTextBrowser_source uintptr
	addr_QTextBrowser_sourceType uintptr
	addr_QTextBrowser_searchPaths uintptr
	addr_QTextBrowser_setSearchPaths uintptr
	addr_QTextBrowser_isBackwardAvailable uintptr
	addr_QTextBrowser_isForwardAvailable uintptr
	addr_QTextBrowser_clearHistory uintptr
	addr_QTextBrowser_historyTitle uintptr
	addr_QTextBrowser_historyUrl uintptr
	addr_QTextBrowser_backwardHistoryCount uintptr
	addr_QTextBrowser_forwardHistoryCount uintptr
	addr_QTextBrowser_openExternalLinks uintptr
	addr_QTextBrowser_setOpenExternalLinks uintptr
	addr_QTextBrowser_openLinks uintptr
	addr_QTextBrowser_setOpenLinks uintptr
	addr_QTextBrowser_setSource uintptr
	addr_QTextBrowser_backward uintptr
	addr_QTextBrowser_forward uintptr
	addr_QTextBrowser_home uintptr
	addr_QTextBrowser_reload uintptr
	addr_QTextBrowser_backwardAvailable uintptr
	addr_QTextBrowser_forwardAvailable uintptr
	addr_QTextBrowser_historyChanged uintptr
	addr_QTextBrowser_sourceChanged uintptr
	addr_QTextBrowser_highlighted uintptr
	addr_QTextBrowser_anchorClicked uintptr
	pg_QTextBrowser_zoomInF func(uintptr, float32)
	addr_QTextBrowser_setViewportMargins uintptr
	addr_QTextBrowser_setViewportMargins2 uintptr
	addr_QTextBrowser_viewportMargins uintptr
	addr_QTextBrowser_drawFrame uintptr
	addr_QTextBrowser_setTabOrder uintptr
	addr_QTextBrowser_mouseGrabber uintptr
	addr_QTextBrowser_keyboardGrabber uintptr
	addr_QTextBrowser_createWindowContainer uintptr
	addr_QTextBrowser_updateMicroFocus uintptr
	addr_QTextBrowser_create uintptr
	addr_QTextBrowser_destroy uintptr
	addr_QTextBrowser_focusNextChild uintptr
	addr_QTextBrowser_focusPreviousChild uintptr
	addr_QTextBrowser_connect uintptr
	addr_QTextBrowser_connect2 uintptr
	addr_QTextBrowser_sender uintptr
	addr_QTextBrowser_senderSignalIndex uintptr
	addr_QTextBrowser_receivers uintptr
	addr_QTextBrowser_isSignalConnected uintptr
	pg_QTextBrowser_devicePixelRatioFScale func() float64
	pg_QTextBrowser_encodeMetricF func(int32, float64) int32
	pg_QTextBrowser_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTextBrowser_tr2 uintptr
	addr_QTextBrowser_tr3 uintptr
	addr_QTextBrowser_setSource2 uintptr
	addr_QTextBrowser_createWindowContainer2 uintptr
	addr_QTextBrowser_createWindowContainer3 uintptr
	addr_QTextBrowser_updateMicroFocus2 uintptr
	addr_QTextBrowser_create2 uintptr
	addr_QTextBrowser_create3 uintptr
	addr_QTextBrowser_create4 uintptr
	addr_QTextBrowser_destroy2 uintptr
	addr_QTextBrowser_destroy3 uintptr
	addr_QTextBrowser_connect3 uintptr
	addr_QTextBrowser_connect4 uintptr
	addr_QTextBrowser_backward_default uintptr
	addr_QTextBrowser_forward_default uintptr
	addr_QTextBrowser_home_default uintptr
	addr_QTextBrowser_reload_default uintptr
	addr_QTextBrowser_metaObject_default uintptr
	addr_QTextBrowser_loadResource_default uintptr
	addr_QTextBrowser_event_default uintptr
	addr_QTextBrowser_keyPressEvent_default uintptr
	addr_QTextBrowser_mouseMoveEvent_default uintptr
	addr_QTextBrowser_mousePressEvent_default uintptr
	addr_QTextBrowser_mouseReleaseEvent_default uintptr
	addr_QTextBrowser_focusOutEvent_default uintptr
	addr_QTextBrowser_focusNextPrevChild_default uintptr
	addr_QTextBrowser_paintEvent_default uintptr
	addr_QTextBrowser_doSetSource_default uintptr
	addr_QTextBrowser_inputMethodQuery_default uintptr
	addr_QTextBrowser_timerEvent_default uintptr
	addr_QTextBrowser_keyReleaseEvent_default uintptr
	addr_QTextBrowser_resizeEvent_default uintptr
	addr_QTextBrowser_mouseDoubleClickEvent_default uintptr
	addr_QTextBrowser_contextMenuEvent_default uintptr
	addr_QTextBrowser_dragEnterEvent_default uintptr
	addr_QTextBrowser_dragLeaveEvent_default uintptr
	addr_QTextBrowser_dragMoveEvent_default uintptr
	addr_QTextBrowser_dropEvent_default uintptr
	addr_QTextBrowser_focusInEvent_default uintptr
	addr_QTextBrowser_showEvent_default uintptr
	addr_QTextBrowser_changeEvent_default uintptr
	addr_QTextBrowser_wheelEvent_default uintptr
	addr_QTextBrowser_createMimeDataFromSelection_default uintptr
	addr_QTextBrowser_canInsertFromMimeData_default uintptr
	addr_QTextBrowser_insertFromMimeData_default uintptr
	addr_QTextBrowser_inputMethodEvent_default uintptr
	addr_QTextBrowser_scrollContentsBy_default uintptr
	addr_QTextBrowser_doSetTextCursor_default uintptr
	addr_QTextBrowser_minimumSizeHint_default uintptr
	addr_QTextBrowser_sizeHint_default uintptr
	addr_QTextBrowser_setupViewport_default uintptr
	addr_QTextBrowser_eventFilter_default uintptr
	addr_QTextBrowser_viewportEvent_default uintptr
	addr_QTextBrowser_viewportSizeHint_default uintptr
	addr_QTextBrowser_initStyleOption_default uintptr
	addr_QTextBrowser_devType_default uintptr
	addr_QTextBrowser_setVisible_default uintptr
	addr_QTextBrowser_heightForWidth_default uintptr
	addr_QTextBrowser_hasHeightForWidth_default uintptr
	addr_QTextBrowser_paintEngine_default uintptr
	addr_QTextBrowser_enterEvent_default uintptr
	addr_QTextBrowser_leaveEvent_default uintptr
	addr_QTextBrowser_moveEvent_default uintptr
	addr_QTextBrowser_closeEvent_default uintptr
	addr_QTextBrowser_tabletEvent_default uintptr
	addr_QTextBrowser_actionEvent_default uintptr
	addr_QTextBrowser_hideEvent_default uintptr
	addr_QTextBrowser_metric_default uintptr
	addr_QTextBrowser_initPainter_default uintptr
	addr_QTextBrowser_redirected_default uintptr
	addr_QTextBrowser_sharedPainter_default uintptr
	addr_QTextBrowser_childEvent_default uintptr
	addr_QTextBrowser_customEvent_default uintptr
	addr_QTextBrowser_connectNotify_default uintptr
	addr_QTextBrowser_disconnectNotify_default uintptr
	pg_QTextBrowser_connect_backwardAvailable func(uintptr, int64)
	pg_QTextBrowser_connect_forwardAvailable func(uintptr, int64)
	pg_QTextBrowser_connect_historyChanged func(uintptr, int64)
	pg_QTextBrowser_connect_sourceChanged func(uintptr, int64)
	pg_QTextBrowser_connect_highlighted func(uintptr, int64)
	pg_QTextBrowser_connect_anchorClicked func(uintptr, int64)
	pg_QTextBrowser_connect_textChanged func(uintptr, int64)
	pg_QTextBrowser_connect_undoAvailable func(uintptr, int64)
	pg_QTextBrowser_connect_redoAvailable func(uintptr, int64)
	pg_QTextBrowser_connect_currentCharFormatChanged func(uintptr, int64)
	pg_QTextBrowser_connect_copyAvailable func(uintptr, int64)
	pg_QTextBrowser_connect_selectionChanged func(uintptr, int64)
	pg_QTextBrowser_connect_cursorPositionChanged func(uintptr, int64)
	pg_QTextBrowser_connect_windowTitleChanged func(uintptr, int64)
	pg_QTextBrowser_connect_windowIconChanged func(uintptr, int64)
	pg_QTextBrowser_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTextBrowser_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTextBrowser_connect_destroyed func(uintptr, int64)
	pg_QTextBrowser_connect_destroyed2 func(uintptr, int64)
	pg_QTextBrowser_connect_objectNameChanged func(uintptr, int64)
	pg_QTextBrowser__disconnect func(int64)
	pg_QTextBrowser__set_override func(uintptr, int32)
	pg_QTextBrowser_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextBrowser_chain, "QTextBrowser_chain")
	symaddr(&addr_QTextBrowser_new, "QTextBrowser_new")
	symaddr(&addr_QTextBrowser_new2, "QTextBrowser_new2")
	symaddr(&addr_QTextBrowser_tr, "QTextBrowser_tr")
	symaddr(&addr_QTextBrowser_source, "QTextBrowser_source")
	symaddr(&addr_QTextBrowser_sourceType, "QTextBrowser_sourceType")
	symaddr(&addr_QTextBrowser_searchPaths, "QTextBrowser_searchPaths")
	symaddr(&addr_QTextBrowser_setSearchPaths, "QTextBrowser_setSearchPaths")
	symaddr(&addr_QTextBrowser_isBackwardAvailable, "QTextBrowser_isBackwardAvailable")
	symaddr(&addr_QTextBrowser_isForwardAvailable, "QTextBrowser_isForwardAvailable")
	symaddr(&addr_QTextBrowser_clearHistory, "QTextBrowser_clearHistory")
	symaddr(&addr_QTextBrowser_historyTitle, "QTextBrowser_historyTitle")
	symaddr(&addr_QTextBrowser_historyUrl, "QTextBrowser_historyUrl")
	symaddr(&addr_QTextBrowser_backwardHistoryCount, "QTextBrowser_backwardHistoryCount")
	symaddr(&addr_QTextBrowser_forwardHistoryCount, "QTextBrowser_forwardHistoryCount")
	symaddr(&addr_QTextBrowser_openExternalLinks, "QTextBrowser_openExternalLinks")
	symaddr(&addr_QTextBrowser_setOpenExternalLinks, "QTextBrowser_setOpenExternalLinks")
	symaddr(&addr_QTextBrowser_openLinks, "QTextBrowser_openLinks")
	symaddr(&addr_QTextBrowser_setOpenLinks, "QTextBrowser_setOpenLinks")
	symaddr(&addr_QTextBrowser_setSource, "QTextBrowser_setSource")
	symaddr(&addr_QTextBrowser_backward, "QTextBrowser_backward")
	symaddr(&addr_QTextBrowser_forward, "QTextBrowser_forward")
	symaddr(&addr_QTextBrowser_home, "QTextBrowser_home")
	symaddr(&addr_QTextBrowser_reload, "QTextBrowser_reload")
	symaddr(&addr_QTextBrowser_backwardAvailable, "QTextBrowser_backwardAvailable")
	symaddr(&addr_QTextBrowser_forwardAvailable, "QTextBrowser_forwardAvailable")
	symaddr(&addr_QTextBrowser_historyChanged, "QTextBrowser_historyChanged")
	symaddr(&addr_QTextBrowser_sourceChanged, "QTextBrowser_sourceChanged")
	symaddr(&addr_QTextBrowser_highlighted, "QTextBrowser_highlighted")
	symaddr(&addr_QTextBrowser_anchorClicked, "QTextBrowser_anchorClicked")
	reg(&pg_QTextBrowser_zoomInF, "QTextBrowser_zoomInF")
	symaddr(&addr_QTextBrowser_setViewportMargins, "QTextBrowser_setViewportMargins")
	symaddr(&addr_QTextBrowser_setViewportMargins2, "QTextBrowser_setViewportMargins2")
	symaddr(&addr_QTextBrowser_viewportMargins, "QTextBrowser_viewportMargins")
	symaddr(&addr_QTextBrowser_drawFrame, "QTextBrowser_drawFrame")
	symaddr(&addr_QTextBrowser_setTabOrder, "QTextBrowser_setTabOrder")
	symaddr(&addr_QTextBrowser_mouseGrabber, "QTextBrowser_mouseGrabber")
	symaddr(&addr_QTextBrowser_keyboardGrabber, "QTextBrowser_keyboardGrabber")
	symaddr(&addr_QTextBrowser_createWindowContainer, "QTextBrowser_createWindowContainer")
	symaddr(&addr_QTextBrowser_updateMicroFocus, "QTextBrowser_updateMicroFocus")
	symaddr(&addr_QTextBrowser_create, "QTextBrowser_create")
	symaddr(&addr_QTextBrowser_destroy, "QTextBrowser_destroy")
	symaddr(&addr_QTextBrowser_focusNextChild, "QTextBrowser_focusNextChild")
	symaddr(&addr_QTextBrowser_focusPreviousChild, "QTextBrowser_focusPreviousChild")
	symaddr(&addr_QTextBrowser_connect, "QTextBrowser_connect")
	symaddr(&addr_QTextBrowser_connect2, "QTextBrowser_connect2")
	symaddr(&addr_QTextBrowser_sender, "QTextBrowser_sender")
	symaddr(&addr_QTextBrowser_senderSignalIndex, "QTextBrowser_senderSignalIndex")
	symaddr(&addr_QTextBrowser_receivers, "QTextBrowser_receivers")
	symaddr(&addr_QTextBrowser_isSignalConnected, "QTextBrowser_isSignalConnected")
	reg(&pg_QTextBrowser_devicePixelRatioFScale, "QTextBrowser_devicePixelRatioFScale")
	reg(&pg_QTextBrowser_encodeMetricF, "QTextBrowser_encodeMetricF")
	reg(&pg_QTextBrowser_getDecodedMetricF, "QTextBrowser_getDecodedMetricF")
	symaddr(&addr_QTextBrowser_tr2, "QTextBrowser_tr2")
	symaddr(&addr_QTextBrowser_tr3, "QTextBrowser_tr3")
	symaddr(&addr_QTextBrowser_setSource2, "QTextBrowser_setSource2")
	symaddr(&addr_QTextBrowser_createWindowContainer2, "QTextBrowser_createWindowContainer2")
	symaddr(&addr_QTextBrowser_createWindowContainer3, "QTextBrowser_createWindowContainer3")
	symaddr(&addr_QTextBrowser_updateMicroFocus2, "QTextBrowser_updateMicroFocus2")
	symaddr(&addr_QTextBrowser_create2, "QTextBrowser_create2")
	symaddr(&addr_QTextBrowser_create3, "QTextBrowser_create3")
	symaddr(&addr_QTextBrowser_create4, "QTextBrowser_create4")
	symaddr(&addr_QTextBrowser_destroy2, "QTextBrowser_destroy2")
	symaddr(&addr_QTextBrowser_destroy3, "QTextBrowser_destroy3")
	symaddr(&addr_QTextBrowser_connect3, "QTextBrowser_connect3")
	symaddr(&addr_QTextBrowser_connect4, "QTextBrowser_connect4")
	symaddr(&addr_QTextBrowser_backward_default, "QTextBrowser_backward_default")
	symaddr(&addr_QTextBrowser_forward_default, "QTextBrowser_forward_default")
	symaddr(&addr_QTextBrowser_home_default, "QTextBrowser_home_default")
	symaddr(&addr_QTextBrowser_reload_default, "QTextBrowser_reload_default")
	symaddr(&addr_QTextBrowser_metaObject_default, "QTextBrowser_metaObject_default")
	symaddr(&addr_QTextBrowser_loadResource_default, "QTextBrowser_loadResource_default")
	symaddr(&addr_QTextBrowser_event_default, "QTextBrowser_event_default")
	symaddr(&addr_QTextBrowser_keyPressEvent_default, "QTextBrowser_keyPressEvent_default")
	symaddr(&addr_QTextBrowser_mouseMoveEvent_default, "QTextBrowser_mouseMoveEvent_default")
	symaddr(&addr_QTextBrowser_mousePressEvent_default, "QTextBrowser_mousePressEvent_default")
	symaddr(&addr_QTextBrowser_mouseReleaseEvent_default, "QTextBrowser_mouseReleaseEvent_default")
	symaddr(&addr_QTextBrowser_focusOutEvent_default, "QTextBrowser_focusOutEvent_default")
	symaddr(&addr_QTextBrowser_focusNextPrevChild_default, "QTextBrowser_focusNextPrevChild_default")
	symaddr(&addr_QTextBrowser_paintEvent_default, "QTextBrowser_paintEvent_default")
	symaddr(&addr_QTextBrowser_doSetSource_default, "QTextBrowser_doSetSource_default")
	symaddr(&addr_QTextBrowser_inputMethodQuery_default, "QTextBrowser_inputMethodQuery_default")
	symaddr(&addr_QTextBrowser_timerEvent_default, "QTextBrowser_timerEvent_default")
	symaddr(&addr_QTextBrowser_keyReleaseEvent_default, "QTextBrowser_keyReleaseEvent_default")
	symaddr(&addr_QTextBrowser_resizeEvent_default, "QTextBrowser_resizeEvent_default")
	symaddr(&addr_QTextBrowser_mouseDoubleClickEvent_default, "QTextBrowser_mouseDoubleClickEvent_default")
	symaddr(&addr_QTextBrowser_contextMenuEvent_default, "QTextBrowser_contextMenuEvent_default")
	symaddr(&addr_QTextBrowser_dragEnterEvent_default, "QTextBrowser_dragEnterEvent_default")
	symaddr(&addr_QTextBrowser_dragLeaveEvent_default, "QTextBrowser_dragLeaveEvent_default")
	symaddr(&addr_QTextBrowser_dragMoveEvent_default, "QTextBrowser_dragMoveEvent_default")
	symaddr(&addr_QTextBrowser_dropEvent_default, "QTextBrowser_dropEvent_default")
	symaddr(&addr_QTextBrowser_focusInEvent_default, "QTextBrowser_focusInEvent_default")
	symaddr(&addr_QTextBrowser_showEvent_default, "QTextBrowser_showEvent_default")
	symaddr(&addr_QTextBrowser_changeEvent_default, "QTextBrowser_changeEvent_default")
	symaddr(&addr_QTextBrowser_wheelEvent_default, "QTextBrowser_wheelEvent_default")
	symaddr(&addr_QTextBrowser_createMimeDataFromSelection_default, "QTextBrowser_createMimeDataFromSelection_default")
	symaddr(&addr_QTextBrowser_canInsertFromMimeData_default, "QTextBrowser_canInsertFromMimeData_default")
	symaddr(&addr_QTextBrowser_insertFromMimeData_default, "QTextBrowser_insertFromMimeData_default")
	symaddr(&addr_QTextBrowser_inputMethodEvent_default, "QTextBrowser_inputMethodEvent_default")
	symaddr(&addr_QTextBrowser_scrollContentsBy_default, "QTextBrowser_scrollContentsBy_default")
	symaddr(&addr_QTextBrowser_doSetTextCursor_default, "QTextBrowser_doSetTextCursor_default")
	symaddr(&addr_QTextBrowser_minimumSizeHint_default, "QTextBrowser_minimumSizeHint_default")
	symaddr(&addr_QTextBrowser_sizeHint_default, "QTextBrowser_sizeHint_default")
	symaddr(&addr_QTextBrowser_setupViewport_default, "QTextBrowser_setupViewport_default")
	symaddr(&addr_QTextBrowser_eventFilter_default, "QTextBrowser_eventFilter_default")
	symaddr(&addr_QTextBrowser_viewportEvent_default, "QTextBrowser_viewportEvent_default")
	symaddr(&addr_QTextBrowser_viewportSizeHint_default, "QTextBrowser_viewportSizeHint_default")
	symaddr(&addr_QTextBrowser_initStyleOption_default, "QTextBrowser_initStyleOption_default")
	symaddr(&addr_QTextBrowser_devType_default, "QTextBrowser_devType_default")
	symaddr(&addr_QTextBrowser_setVisible_default, "QTextBrowser_setVisible_default")
	symaddr(&addr_QTextBrowser_heightForWidth_default, "QTextBrowser_heightForWidth_default")
	symaddr(&addr_QTextBrowser_hasHeightForWidth_default, "QTextBrowser_hasHeightForWidth_default")
	symaddr(&addr_QTextBrowser_paintEngine_default, "QTextBrowser_paintEngine_default")
	symaddr(&addr_QTextBrowser_enterEvent_default, "QTextBrowser_enterEvent_default")
	symaddr(&addr_QTextBrowser_leaveEvent_default, "QTextBrowser_leaveEvent_default")
	symaddr(&addr_QTextBrowser_moveEvent_default, "QTextBrowser_moveEvent_default")
	symaddr(&addr_QTextBrowser_closeEvent_default, "QTextBrowser_closeEvent_default")
	symaddr(&addr_QTextBrowser_tabletEvent_default, "QTextBrowser_tabletEvent_default")
	symaddr(&addr_QTextBrowser_actionEvent_default, "QTextBrowser_actionEvent_default")
	symaddr(&addr_QTextBrowser_hideEvent_default, "QTextBrowser_hideEvent_default")
	symaddr(&addr_QTextBrowser_metric_default, "QTextBrowser_metric_default")
	symaddr(&addr_QTextBrowser_initPainter_default, "QTextBrowser_initPainter_default")
	symaddr(&addr_QTextBrowser_redirected_default, "QTextBrowser_redirected_default")
	symaddr(&addr_QTextBrowser_sharedPainter_default, "QTextBrowser_sharedPainter_default")
	symaddr(&addr_QTextBrowser_childEvent_default, "QTextBrowser_childEvent_default")
	symaddr(&addr_QTextBrowser_customEvent_default, "QTextBrowser_customEvent_default")
	symaddr(&addr_QTextBrowser_connectNotify_default, "QTextBrowser_connectNotify_default")
	symaddr(&addr_QTextBrowser_disconnectNotify_default, "QTextBrowser_disconnectNotify_default")
	reg(&pg_QTextBrowser_connect_backwardAvailable, "QTextBrowser_connect_backwardAvailable")
	reg(&pg_QTextBrowser_connect_forwardAvailable, "QTextBrowser_connect_forwardAvailable")
	reg(&pg_QTextBrowser_connect_historyChanged, "QTextBrowser_connect_historyChanged")
	reg(&pg_QTextBrowser_connect_sourceChanged, "QTextBrowser_connect_sourceChanged")
	reg(&pg_QTextBrowser_connect_highlighted, "QTextBrowser_connect_highlighted")
	reg(&pg_QTextBrowser_connect_anchorClicked, "QTextBrowser_connect_anchorClicked")
	reg(&pg_QTextBrowser_connect_textChanged, "QTextBrowser_connect_textChanged")
	reg(&pg_QTextBrowser_connect_undoAvailable, "QTextBrowser_connect_undoAvailable")
	reg(&pg_QTextBrowser_connect_redoAvailable, "QTextBrowser_connect_redoAvailable")
	reg(&pg_QTextBrowser_connect_currentCharFormatChanged, "QTextBrowser_connect_currentCharFormatChanged")
	reg(&pg_QTextBrowser_connect_copyAvailable, "QTextBrowser_connect_copyAvailable")
	reg(&pg_QTextBrowser_connect_selectionChanged, "QTextBrowser_connect_selectionChanged")
	reg(&pg_QTextBrowser_connect_cursorPositionChanged, "QTextBrowser_connect_cursorPositionChanged")
	reg(&pg_QTextBrowser_connect_windowTitleChanged, "QTextBrowser_connect_windowTitleChanged")
	reg(&pg_QTextBrowser_connect_windowIconChanged, "QTextBrowser_connect_windowIconChanged")
	reg(&pg_QTextBrowser_connect_windowIconTextChanged, "QTextBrowser_connect_windowIconTextChanged")
	reg(&pg_QTextBrowser_connect_customContextMenuRequested, "QTextBrowser_connect_customContextMenuRequested")
	reg(&pg_QTextBrowser_connect_destroyed, "QTextBrowser_connect_destroyed")
	reg(&pg_QTextBrowser_connect_destroyed2, "QTextBrowser_connect_destroyed2")
	reg(&pg_QTextBrowser_connect_objectNameChanged, "QTextBrowser_connect_objectNameChanged")
	reg(&pg_QTextBrowser__disconnect, "QTextBrowser__disconnect")
	reg(&pg_QTextBrowser__set_override, "QTextBrowser__set_override")
	reg(&pg_QTextBrowser_delete, "QTextBrowser_delete")
}

type QTextBrowser struct {
	h uintptr
	*QTextEdit
}

func bindQTextBrowser(h uintptr) *QTextBrowser {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QTextBrowser_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextBrowser{h: h, QTextEdit: &QTextEdit{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQTextBrowser(h uintptr) *QTextBrowser { return bindQTextBrowser(h) }

func (this *QTextBrowser) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBrowser) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBrowser(h unsafe.Pointer) *QTextBrowser { return bindQTextBrowser(uintptr(h)) }
func UnsafeNewQTextBrowserOwned(h unsafe.Pointer) *QTextBrowser { return newownedQTextBrowser(uintptr(h)) }

func (this *QTextBrowser) UnsafeRelease() {
	this.h = 0
	if this.QTextEdit != nil {
		this.QTextEdit.UnsafeRelease()
	}
}

func NewQTextBrowser(parent *QWidget) *QTextBrowser {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTextBrowser(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextBrowser_virtIdx, QTextBrowser_enable)
	return __o
}

func NewQTextBrowser2() *QTextBrowser {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_new2, uintptr(__goid))
	__o := newownedQTextBrowser(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTextBrowser_virtIdx, QTextBrowser_enable)
	return __o
}

func QTextBrowser_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextBrowser) Source() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_source, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QTextBrowser) SourceType() QTextDocument__ResourceType {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_sourceType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextDocument__ResourceType(__r)
}

func (this *QTextBrowser) SearchPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_searchPaths, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QTextBrowser) SetSearchPaths(paths []string) {
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	purego.SyscallN(addr_QTextBrowser_setSearchPaths, this.h, paths_p, uintptr(len(paths)))
	runtime.KeepAlive(this)
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
}

func (this *QTextBrowser) IsBackwardAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_isBackwardAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) IsForwardAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_isForwardAvailable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) ClearHistory() {
	purego.SyscallN(addr_QTextBrowser_clearHistory, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) HistoryTitle(arg1 int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_historyTitle, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextBrowser) HistoryUrl(arg1 int) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_historyUrl, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QTextBrowser) BackwardHistoryCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_backwardHistoryCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) ForwardHistoryCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_forwardHistoryCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) OpenExternalLinks() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_openExternalLinks, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	purego.SyscallN(addr_QTextBrowser_setOpenExternalLinks, this.h, uintptr(b2i(open)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) OpenLinks() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_openLinks, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) SetOpenLinks(open bool) {
	purego.SyscallN(addr_QTextBrowser_setOpenLinks, this.h, uintptr(b2i(open)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) SetSource(name *QUrl) {
	purego.SyscallN(addr_QTextBrowser_setSource, this.h, uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTextBrowser) Backward() {
	purego.SyscallN(addr_QTextBrowser_backward, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Forward() {
	purego.SyscallN(addr_QTextBrowser_forward, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Home() {
	purego.SyscallN(addr_QTextBrowser_home, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Reload() {
	purego.SyscallN(addr_QTextBrowser_reload, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) BackwardAvailable(arg1 bool) {
	purego.SyscallN(addr_QTextBrowser_backwardAvailable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) ForwardAvailable(arg1 bool) {
	purego.SyscallN(addr_QTextBrowser_forwardAvailable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) HistoryChanged() {
	purego.SyscallN(addr_QTextBrowser_historyChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) SourceChanged(arg1 *QUrl) {
	purego.SyscallN(addr_QTextBrowser_sourceChanged, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextBrowser) Highlighted(arg1 *QUrl) {
	purego.SyscallN(addr_QTextBrowser_highlighted, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextBrowser) AnchorClicked(arg1 *QUrl) {
	purego.SyscallN(addr_QTextBrowser_anchorClicked, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextBrowser) ZoomInF(range_ float32) {
	pg_QTextBrowser_zoomInF(this.h, range_)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QTextBrowser_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QTextBrowser_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QTextBrowser) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QTextBrowser) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QTextBrowser_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTextBrowser_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTextBrowser_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTextBrowser_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_mouseGrabber)
	return bindQWidget(__r)
}

func QTextBrowser_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_keyboardGrabber)
	return bindQWidget(__r)
}

func QTextBrowser_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTextBrowser) UpdateMicroFocus() {
	purego.SyscallN(addr_QTextBrowser_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Create() {
	purego.SyscallN(addr_QTextBrowser_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Destroy() {
	purego.SyscallN(addr_QTextBrowser_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTextBrowser_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextBrowser_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextBrowser) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTextBrowser) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTextBrowser) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTextBrowser_DevicePixelRatioFScale() float64 {
	__r := pg_QTextBrowser_devicePixelRatioFScale()
	return __r
}

func QTextBrowser_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTextBrowser_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTextBrowser) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTextBrowser_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTextBrowser_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextBrowser_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTextBrowser) SetSource2(name *QUrl, type_ QTextDocument__ResourceType) {
	purego.SyscallN(addr_QTextBrowser_setSource2, this.h, uintptr(name.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func QTextBrowser_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTextBrowser_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTextBrowser) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTextBrowser_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTextBrowser_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTextBrowser_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTextBrowser_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTextBrowser_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTextBrowser_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTextBrowser_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextBrowser_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextBrowser) BackwardDefault() {
	purego.SyscallN(addr_QTextBrowser_backward_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) ForwardDefault() {
	purego.SyscallN(addr_QTextBrowser_forward_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) HomeDefault() {
	purego.SyscallN(addr_QTextBrowser_home_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) ReloadDefault() {
	purego.SyscallN(addr_QTextBrowser_reload_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTextBrowser) LoadResourceDefault(type_ int, name *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_loadResource_default, this.h, uintptr(type_), uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QTextBrowser) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QTextBrowser) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QTextBrowser_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QTextBrowser) MouseMoveEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QTextBrowser_mouseMoveEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QTextBrowser) MousePressEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QTextBrowser_mousePressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QTextBrowser) MouseReleaseEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QTextBrowser_mouseReleaseEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QTextBrowser) FocusOutEventDefault(ev *QFocusEvent) {
	purego.SyscallN(addr_QTextBrowser_focusOutEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QTextBrowser) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QTextBrowser_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) DoSetSourceDefault(name *QUrl, type_ QTextDocument__ResourceType) {
	purego.SyscallN(addr_QTextBrowser_doSetSource_default, this.h, uintptr(name.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTextBrowser) InputMethodQueryDefault(property InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_inputMethodQuery_default, this.h, uintptr(property))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextBrowser) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QTextBrowser_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QTextBrowser_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QTextBrowser_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) MouseDoubleClickEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QTextBrowser_mouseDoubleClickEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) ContextMenuEventDefault(e *QContextMenuEvent) {
	purego.SyscallN(addr_QTextBrowser_contextMenuEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) DragEnterEventDefault(e *QDragEnterEvent) {
	purego.SyscallN(addr_QTextBrowser_dragEnterEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) DragLeaveEventDefault(e *QDragLeaveEvent) {
	purego.SyscallN(addr_QTextBrowser_dragLeaveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) DragMoveEventDefault(e *QDragMoveEvent) {
	purego.SyscallN(addr_QTextBrowser_dragMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) DropEventDefault(e *QDropEvent) {
	purego.SyscallN(addr_QTextBrowser_dropEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QTextBrowser_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QTextBrowser_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextBrowser) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QTextBrowser_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QTextBrowser_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QTextBrowser) CreateMimeDataFromSelectionDefault() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_createMimeDataFromSelection_default, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QTextBrowser) CanInsertFromMimeDataDefault(source *QMimeData) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_canInsertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
	return __r != 0
}

func (this *QTextBrowser) InsertFromMimeDataDefault(source *QMimeData) {
	purego.SyscallN(addr_QTextBrowser_insertFromMimeData_default, this.h, uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source)
}

func (this *QTextBrowser) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QTextBrowser_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextBrowser) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QTextBrowser_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) DoSetTextCursorDefault(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextBrowser_doSetTextCursor_default, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextBrowser) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextBrowser) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextBrowser) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QTextBrowser_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QTextBrowser) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QTextBrowser) ViewportEventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_viewportEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTextBrowser) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTextBrowser) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QTextBrowser_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTextBrowser) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTextBrowser_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTextBrowser) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBrowser) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTextBrowser) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTextBrowser_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTextBrowser_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTextBrowser_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTextBrowser_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTextBrowser_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTextBrowser_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTextBrowser_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTextBrowser_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextBrowser) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTextBrowser_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTextBrowser) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTextBrowser) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTextBrowser_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTextBrowser) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTextBrowser_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTextBrowser_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTextBrowser) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextBrowser_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextBrowser) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTextBrowser_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTextBrowser) OnBackwardAvailable(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QTextBrowser_connect_backwardAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnForwardAvailable(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QTextBrowser_connect_forwardAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnHistoryChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBrowser_connect_historyChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnSourceChanged(fn func(*QUrl)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(arg1))) })
	pg_QTextBrowser_connect_sourceChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnHighlighted(fn func(*QUrl)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(arg1))) })
	pg_QTextBrowser_connect_highlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnAnchorClicked(fn func(*QUrl)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQUrl(unsafe.Pointer(arg1))) })
	pg_QTextBrowser_connect_anchorClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnTextChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBrowser_connect_textChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnUndoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextBrowser_connect_undoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnRedoAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextBrowser_connect_redoAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnCurrentCharFormatChanged(fn func(*QTextCharFormat)) int64 {
	id := registerCallback(func(format uintptr) { fn(UnsafeNewQTextCharFormat(unsafe.Pointer(format))) })
	pg_QTextBrowser_connect_currentCharFormatChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnCopyAvailable(fn func(bool)) int64 {
	id := registerCallback(func(b int32) { fn(b != 0) })
	pg_QTextBrowser_connect_copyAvailable(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBrowser_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnCursorPositionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBrowser_connect_cursorPositionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTextBrowser_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTextBrowser_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTextBrowser_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTextBrowser_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBrowser_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextBrowser_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextBrowser_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBrowser) Disconnect(handle int64) {
	pg_QTextBrowser__disconnect(handle)
	unregisterCallback(handle)
}

func QTextBrowser_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -519616681208050567:
		return 1
	case 2479296857741040516:
		return 2
	case 2190319719396888442:
		return 3
	case 4624382957487889774:
		return 4
	case -1731273083434011732:
		return 5
	case 1237236456312824991:
		return 6
	case 1023791977679793185:
		return 7
	case -4255494004099806581:
		return 8
	case -332090898078543319:
		return 9
	case -4415920049297795213:
		return 10
	case -385596010576048993:
		return 11
	case 5898328517505784191:
		return 12
	case 5290442163167458465:
		return 13
	case 3615073669235633897:
		return 14
	case 7345853097340709638:
		return 15
	case -5409927581812374498:
		return 16
	case 772598475950050299:
		return 17
	case -4302917285655014661:
		return 18
	case -4029073905697023887:
		return 19
	case -6698043222227006015:
		return 20
	case 206122715919917589:
		return 21
	case -3139904158373725006:
		return 22
	case -1781545137352745978:
		return 23
	case 9107058682754591716:
		return 24
	case 975043767033552552:
		return 25
	case -6825369174751916:
		return 26
	case 4959376063826165943:
		return 27
	case 1919988340934511316:
		return 28
	case 3493558106800165317:
		return 29
	case -2605205276815256896:
		return 30
	case -466553387102125056:
		return 31
	case 5438819661942843222:
		return 32
	case -6118001150126564745:
		return 33
	case 4090661491944147209:
		return 34
	case 7049264599837532491:
		return 35
	case 729690261457684101:
		return 36
	case 4531996381736730688:
		return 37
	case 5217623514630858297:
		return 38
	case -1971465653061861519:
		return 39
	case 503579067252048239:
		return 40
	case -488693769771308531:
		return 41
	case -8659679893868524044:
		return 42
	case 2110182719375245575:
		return 43
	case 4311534060495737339:
		return 44
	case -8266762397995423915:
		return 45
	case 6398007404756346381:
		return 46
	case 6316270346668237895:
		return 47
	case -551070009623711640:
		return 48
	case -9161949150489525396:
		return 49
	case -1989937050019464783:
		return 50
	case -212068866860701141:
		return 51
	case 1697444140088555581:
		return 52
	case -199718707876719129:
		return 53
	case -2039166586246443706:
		return 54
	case 427629649267311777:
		return 55
	case 6477500296980050210:
		return 56
	case 4127240633310655048:
		return 57
	case 5792771328741774991:
		return 58
	case -1185315367526857833:
		return 59
	case 3769868229123868474:
		return 60
	case 7379029982170503008:
		return 61
	case -568369772982215962:
		return 62
	}
	return -1
}

func QTextBrowser_enable(h uintptr, idx int32) { pg_QTextBrowser__set_override(h, idx) }

func (this *QTextBrowser) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTextBrowser) OnLoadResource(fn func(int, *QUrl) *QVariant) {
	registerVirt(this.UnsafeGoID(), -519616681208050567, func(type_ int32, name uintptr) uintptr { return uintptr(fn(int(type_), UnsafeNewQUrl(unsafe.Pointer(name))).UnsafePointer()) }) // slot "loadResource"
}

func (this *QTextBrowser) OnBackward(fn func()) {
	registerVirt(this.UnsafeGoID(), 2479296857741040516, func() { fn() }) // slot "backward"
}

func (this *QTextBrowser) OnForward(fn func()) {
	registerVirt(this.UnsafeGoID(), 2190319719396888442, func() { fn() }) // slot "forward"
}

func (this *QTextBrowser) OnHome(fn func()) {
	registerVirt(this.UnsafeGoID(), 4624382957487889774, func() { fn() }) // slot "home"
}

func (this *QTextBrowser) OnReload(fn func()) {
	registerVirt(this.UnsafeGoID(), -1731273083434011732, func() { fn() }) // slot "reload"
}

func (this *QTextBrowser) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QTextBrowser) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QTextBrowser) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseMoveEvent"
}

func (this *QTextBrowser) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mousePressEvent"
}

func (this *QTextBrowser) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseReleaseEvent"
}

func (this *QTextBrowser) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(ev uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(ev))) }) // slot "focusOutEvent"
}

func (this *QTextBrowser) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTextBrowser) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QTextBrowser) OnDoSetSource(fn func(*QUrl, QTextDocument__ResourceType)) {
	registerVirt(this.UnsafeGoID(), 3615073669235633897, func(name uintptr, type_ int32) { fn(UnsafeNewQUrl(unsafe.Pointer(name)), QTextDocument__ResourceType(type_)) }) // slot "doSetSource"
}

func (this *QTextBrowser) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(property int32) uintptr { return uintptr(fn(InputMethodQuery(property)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTextBrowser) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QTextBrowser) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QTextBrowser) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QTextBrowser) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTextBrowser) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(e uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(e))) }) // slot "contextMenuEvent"
}

func (this *QTextBrowser) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(e uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(e))) }) // slot "dragEnterEvent"
}

func (this *QTextBrowser) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(e uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(e))) }) // slot "dragLeaveEvent"
}

func (this *QTextBrowser) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(e uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(e))) }) // slot "dragMoveEvent"
}

func (this *QTextBrowser) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(e uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(e))) }) // slot "dropEvent"
}

func (this *QTextBrowser) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QTextBrowser) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QTextBrowser) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QTextBrowser) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QTextBrowser) OnCreateMimeDataFromSelection(fn func() *QMimeData) {
	registerVirt(this.UnsafeGoID(), 3493558106800165317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "createMimeDataFromSelection"
}

func (this *QTextBrowser) OnCanInsertFromMimeData(fn func(*QMimeData) bool) {
	registerVirt(this.UnsafeGoID(), -2605205276815256896, func(source uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(source)))) }) // slot "canInsertFromMimeData"
}

func (this *QTextBrowser) OnInsertFromMimeData(fn func(*QMimeData)) {
	registerVirt(this.UnsafeGoID(), -466553387102125056, func(source uintptr) { fn(UnsafeNewQMimeData(unsafe.Pointer(source))) }) // slot "insertFromMimeData"
}

func (this *QTextBrowser) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QTextBrowser) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QTextBrowser) OnDoSetTextCursor(fn func(*QTextCursor)) {
	registerVirt(this.UnsafeGoID(), 4090661491944147209, func(cursor uintptr) { fn(UnsafeNewQTextCursor(unsafe.Pointer(cursor))) }) // slot "doSetTextCursor"
}

func (this *QTextBrowser) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTextBrowser) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTextBrowser) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QTextBrowser) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QTextBrowser) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "viewportEvent"
}

func (this *QTextBrowser) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QTextBrowser) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTextBrowser) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTextBrowser) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTextBrowser) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTextBrowser) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTextBrowser) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTextBrowser) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTextBrowser) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTextBrowser) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTextBrowser) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTextBrowser) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTextBrowser) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTextBrowser) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTextBrowser) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTextBrowser) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTextBrowser) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTextBrowser) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTextBrowser) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTextBrowser) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTextBrowser) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTextBrowser) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTextBrowser) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTextBrowser) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextBrowser_delete(this.h)
	}
	this.UnsafeRelease()
}
