package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWizard_chain uintptr
	addr_QWizard_new uintptr
	addr_QWizard_new2 uintptr
	addr_QWizard_new3 uintptr
	addr_QWizard_tr uintptr
	addr_QWizard_addPage uintptr
	addr_QWizard_setPage uintptr
	addr_QWizard_removePage uintptr
	addr_QWizard_page uintptr
	addr_QWizard_hasVisitedPage uintptr
	addr_QWizard_visitedIds uintptr
	addr_QWizard_pageIds uintptr
	addr_QWizard_setStartId uintptr
	addr_QWizard_startId uintptr
	addr_QWizard_currentPage uintptr
	addr_QWizard_currentId uintptr
	addr_QWizard_validateCurrentPage uintptr
	addr_QWizard_nextId uintptr
	addr_QWizard_setField uintptr
	addr_QWizard_field uintptr
	addr_QWizard_setWizardStyle uintptr
	addr_QWizard_wizardStyle uintptr
	addr_QWizard_setOption uintptr
	addr_QWizard_testOption uintptr
	addr_QWizard_setOptions uintptr
	addr_QWizard_options uintptr
	addr_QWizard_setButtonText uintptr
	addr_QWizard_buttonText uintptr
	addr_QWizard_setButtonLayout uintptr
	addr_QWizard_setButton uintptr
	addr_QWizard_button uintptr
	addr_QWizard_setTitleFormat uintptr
	addr_QWizard_titleFormat uintptr
	addr_QWizard_setSubTitleFormat uintptr
	addr_QWizard_subTitleFormat uintptr
	addr_QWizard_setPixmap uintptr
	addr_QWizard_pixmap uintptr
	addr_QWizard_setSideWidget uintptr
	addr_QWizard_sideWidget uintptr
	addr_QWizard_setDefaultProperty uintptr
	addr_QWizard_currentIdChanged uintptr
	addr_QWizard_helpRequested uintptr
	addr_QWizard_customButtonClicked uintptr
	addr_QWizard_pageAdded uintptr
	addr_QWizard_pageRemoved uintptr
	addr_QWizard_back uintptr
	addr_QWizard_next uintptr
	addr_QWizard_setCurrentId uintptr
	addr_QWizard_restart uintptr
	addr_QWizard_adjustPosition uintptr
	addr_QWizard_setTabOrder uintptr
	addr_QWizard_mouseGrabber uintptr
	addr_QWizard_keyboardGrabber uintptr
	addr_QWizard_find uintptr
	addr_QWizard_createWindowContainer uintptr
	addr_QWizard_updateMicroFocus uintptr
	addr_QWizard_create uintptr
	addr_QWizard_destroy uintptr
	addr_QWizard_focusNextChild uintptr
	addr_QWizard_focusPreviousChild uintptr
	addr_QWizard_connect uintptr
	addr_QWizard_connect2 uintptr
	addr_QWizard_sender uintptr
	addr_QWizard_senderSignalIndex uintptr
	addr_QWizard_receivers uintptr
	addr_QWizard_isSignalConnected uintptr
	pg_QWizard_devicePixelRatioFScale func() float64
	pg_QWizard_encodeMetricF func(int32, float64) int32
	pg_QWizard_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QWizard_tr2 uintptr
	addr_QWizard_tr3 uintptr
	addr_QWizard_setOption2 uintptr
	addr_QWizard_createWindowContainer2 uintptr
	addr_QWizard_createWindowContainer3 uintptr
	addr_QWizard_updateMicroFocus2 uintptr
	addr_QWizard_create2 uintptr
	addr_QWizard_create3 uintptr
	addr_QWizard_create4 uintptr
	addr_QWizard_destroy2 uintptr
	addr_QWizard_destroy3 uintptr
	addr_QWizard_connect3 uintptr
	addr_QWizard_connect4 uintptr
	addr_QWizard_validateCurrentPage_default uintptr
	addr_QWizard_nextId_default uintptr
	addr_QWizard_metaObject_default uintptr
	addr_QWizard_setVisible_default uintptr
	addr_QWizard_sizeHint_default uintptr
	addr_QWizard_event_default uintptr
	addr_QWizard_resizeEvent_default uintptr
	addr_QWizard_paintEvent_default uintptr
	addr_QWizard_done_default uintptr
	addr_QWizard_initializePage_default uintptr
	addr_QWizard_cleanupPage_default uintptr
	addr_QWizard_minimumSizeHint_default uintptr
	addr_QWizard_open_default uintptr
	addr_QWizard_exec_default uintptr
	addr_QWizard_accept_default uintptr
	addr_QWizard_reject_default uintptr
	addr_QWizard_keyPressEvent_default uintptr
	addr_QWizard_closeEvent_default uintptr
	addr_QWizard_showEvent_default uintptr
	addr_QWizard_contextMenuEvent_default uintptr
	addr_QWizard_eventFilter_default uintptr
	addr_QWizard_devType_default uintptr
	addr_QWizard_heightForWidth_default uintptr
	addr_QWizard_hasHeightForWidth_default uintptr
	addr_QWizard_paintEngine_default uintptr
	addr_QWizard_mousePressEvent_default uintptr
	addr_QWizard_mouseReleaseEvent_default uintptr
	addr_QWizard_mouseDoubleClickEvent_default uintptr
	addr_QWizard_mouseMoveEvent_default uintptr
	addr_QWizard_wheelEvent_default uintptr
	addr_QWizard_keyReleaseEvent_default uintptr
	addr_QWizard_focusInEvent_default uintptr
	addr_QWizard_focusOutEvent_default uintptr
	addr_QWizard_enterEvent_default uintptr
	addr_QWizard_leaveEvent_default uintptr
	addr_QWizard_moveEvent_default uintptr
	addr_QWizard_tabletEvent_default uintptr
	addr_QWizard_actionEvent_default uintptr
	addr_QWizard_dragEnterEvent_default uintptr
	addr_QWizard_dragMoveEvent_default uintptr
	addr_QWizard_dragLeaveEvent_default uintptr
	addr_QWizard_dropEvent_default uintptr
	addr_QWizard_hideEvent_default uintptr
	addr_QWizard_changeEvent_default uintptr
	addr_QWizard_metric_default uintptr
	addr_QWizard_initPainter_default uintptr
	addr_QWizard_redirected_default uintptr
	addr_QWizard_sharedPainter_default uintptr
	addr_QWizard_inputMethodEvent_default uintptr
	addr_QWizard_inputMethodQuery_default uintptr
	addr_QWizard_focusNextPrevChild_default uintptr
	addr_QWizard_timerEvent_default uintptr
	addr_QWizard_childEvent_default uintptr
	addr_QWizard_customEvent_default uintptr
	addr_QWizard_connectNotify_default uintptr
	addr_QWizard_disconnectNotify_default uintptr
	pg_QWizard_connect_currentIdChanged func(uintptr, int64)
	pg_QWizard_connect_helpRequested func(uintptr, int64)
	pg_QWizard_connect_customButtonClicked func(uintptr, int64)
	pg_QWizard_connect_pageAdded func(uintptr, int64)
	pg_QWizard_connect_pageRemoved func(uintptr, int64)
	pg_QWizard_connect_finished func(uintptr, int64)
	pg_QWizard_connect_accepted func(uintptr, int64)
	pg_QWizard_connect_rejected func(uintptr, int64)
	pg_QWizard_connect_windowTitleChanged func(uintptr, int64)
	pg_QWizard_connect_windowIconChanged func(uintptr, int64)
	pg_QWizard_connect_windowIconTextChanged func(uintptr, int64)
	pg_QWizard_connect_customContextMenuRequested func(uintptr, int64)
	pg_QWizard_connect_destroyed func(uintptr, int64)
	pg_QWizard_connect_destroyed2 func(uintptr, int64)
	pg_QWizard_connect_objectNameChanged func(uintptr, int64)
	pg_QWizard__disconnect func(int64)
	pg_QWizard__set_override func(uintptr, int32)
	pg_QWizard_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWizard_chain, "QWizard_chain")
	symaddr(&addr_QWizard_new, "QWizard_new")
	symaddr(&addr_QWizard_new2, "QWizard_new2")
	symaddr(&addr_QWizard_new3, "QWizard_new3")
	symaddr(&addr_QWizard_tr, "QWizard_tr")
	symaddr(&addr_QWizard_addPage, "QWizard_addPage")
	symaddr(&addr_QWizard_setPage, "QWizard_setPage")
	symaddr(&addr_QWizard_removePage, "QWizard_removePage")
	symaddr(&addr_QWizard_page, "QWizard_page")
	symaddr(&addr_QWizard_hasVisitedPage, "QWizard_hasVisitedPage")
	symaddr(&addr_QWizard_visitedIds, "QWizard_visitedIds")
	symaddr(&addr_QWizard_pageIds, "QWizard_pageIds")
	symaddr(&addr_QWizard_setStartId, "QWizard_setStartId")
	symaddr(&addr_QWizard_startId, "QWizard_startId")
	symaddr(&addr_QWizard_currentPage, "QWizard_currentPage")
	symaddr(&addr_QWizard_currentId, "QWizard_currentId")
	symaddr(&addr_QWizard_validateCurrentPage, "QWizard_validateCurrentPage")
	symaddr(&addr_QWizard_nextId, "QWizard_nextId")
	symaddr(&addr_QWizard_setField, "QWizard_setField")
	symaddr(&addr_QWizard_field, "QWizard_field")
	symaddr(&addr_QWizard_setWizardStyle, "QWizard_setWizardStyle")
	symaddr(&addr_QWizard_wizardStyle, "QWizard_wizardStyle")
	symaddr(&addr_QWizard_setOption, "QWizard_setOption")
	symaddr(&addr_QWizard_testOption, "QWizard_testOption")
	symaddr(&addr_QWizard_setOptions, "QWizard_setOptions")
	symaddr(&addr_QWizard_options, "QWizard_options")
	symaddr(&addr_QWizard_setButtonText, "QWizard_setButtonText")
	symaddr(&addr_QWizard_buttonText, "QWizard_buttonText")
	symaddr(&addr_QWizard_setButtonLayout, "QWizard_setButtonLayout")
	symaddr(&addr_QWizard_setButton, "QWizard_setButton")
	symaddr(&addr_QWizard_button, "QWizard_button")
	symaddr(&addr_QWizard_setTitleFormat, "QWizard_setTitleFormat")
	symaddr(&addr_QWizard_titleFormat, "QWizard_titleFormat")
	symaddr(&addr_QWizard_setSubTitleFormat, "QWizard_setSubTitleFormat")
	symaddr(&addr_QWizard_subTitleFormat, "QWizard_subTitleFormat")
	symaddr(&addr_QWizard_setPixmap, "QWizard_setPixmap")
	symaddr(&addr_QWizard_pixmap, "QWizard_pixmap")
	symaddr(&addr_QWizard_setSideWidget, "QWizard_setSideWidget")
	symaddr(&addr_QWizard_sideWidget, "QWizard_sideWidget")
	symaddr(&addr_QWizard_setDefaultProperty, "QWizard_setDefaultProperty")
	symaddr(&addr_QWizard_currentIdChanged, "QWizard_currentIdChanged")
	symaddr(&addr_QWizard_helpRequested, "QWizard_helpRequested")
	symaddr(&addr_QWizard_customButtonClicked, "QWizard_customButtonClicked")
	symaddr(&addr_QWizard_pageAdded, "QWizard_pageAdded")
	symaddr(&addr_QWizard_pageRemoved, "QWizard_pageRemoved")
	symaddr(&addr_QWizard_back, "QWizard_back")
	symaddr(&addr_QWizard_next, "QWizard_next")
	symaddr(&addr_QWizard_setCurrentId, "QWizard_setCurrentId")
	symaddr(&addr_QWizard_restart, "QWizard_restart")
	symaddr(&addr_QWizard_adjustPosition, "QWizard_adjustPosition")
	symaddr(&addr_QWizard_setTabOrder, "QWizard_setTabOrder")
	symaddr(&addr_QWizard_mouseGrabber, "QWizard_mouseGrabber")
	symaddr(&addr_QWizard_keyboardGrabber, "QWizard_keyboardGrabber")
	symaddr(&addr_QWizard_find, "QWizard_find")
	symaddr(&addr_QWizard_createWindowContainer, "QWizard_createWindowContainer")
	symaddr(&addr_QWizard_updateMicroFocus, "QWizard_updateMicroFocus")
	symaddr(&addr_QWizard_create, "QWizard_create")
	symaddr(&addr_QWizard_destroy, "QWizard_destroy")
	symaddr(&addr_QWizard_focusNextChild, "QWizard_focusNextChild")
	symaddr(&addr_QWizard_focusPreviousChild, "QWizard_focusPreviousChild")
	symaddr(&addr_QWizard_connect, "QWizard_connect")
	symaddr(&addr_QWizard_connect2, "QWizard_connect2")
	symaddr(&addr_QWizard_sender, "QWizard_sender")
	symaddr(&addr_QWizard_senderSignalIndex, "QWizard_senderSignalIndex")
	symaddr(&addr_QWizard_receivers, "QWizard_receivers")
	symaddr(&addr_QWizard_isSignalConnected, "QWizard_isSignalConnected")
	reg(&pg_QWizard_devicePixelRatioFScale, "QWizard_devicePixelRatioFScale")
	reg(&pg_QWizard_encodeMetricF, "QWizard_encodeMetricF")
	reg(&pg_QWizard_getDecodedMetricF, "QWizard_getDecodedMetricF")
	symaddr(&addr_QWizard_tr2, "QWizard_tr2")
	symaddr(&addr_QWizard_tr3, "QWizard_tr3")
	symaddr(&addr_QWizard_setOption2, "QWizard_setOption2")
	symaddr(&addr_QWizard_createWindowContainer2, "QWizard_createWindowContainer2")
	symaddr(&addr_QWizard_createWindowContainer3, "QWizard_createWindowContainer3")
	symaddr(&addr_QWizard_updateMicroFocus2, "QWizard_updateMicroFocus2")
	symaddr(&addr_QWizard_create2, "QWizard_create2")
	symaddr(&addr_QWizard_create3, "QWizard_create3")
	symaddr(&addr_QWizard_create4, "QWizard_create4")
	symaddr(&addr_QWizard_destroy2, "QWizard_destroy2")
	symaddr(&addr_QWizard_destroy3, "QWizard_destroy3")
	symaddr(&addr_QWizard_connect3, "QWizard_connect3")
	symaddr(&addr_QWizard_connect4, "QWizard_connect4")
	symaddr(&addr_QWizard_validateCurrentPage_default, "QWizard_validateCurrentPage_default")
	symaddr(&addr_QWizard_nextId_default, "QWizard_nextId_default")
	symaddr(&addr_QWizard_metaObject_default, "QWizard_metaObject_default")
	symaddr(&addr_QWizard_setVisible_default, "QWizard_setVisible_default")
	symaddr(&addr_QWizard_sizeHint_default, "QWizard_sizeHint_default")
	symaddr(&addr_QWizard_event_default, "QWizard_event_default")
	symaddr(&addr_QWizard_resizeEvent_default, "QWizard_resizeEvent_default")
	symaddr(&addr_QWizard_paintEvent_default, "QWizard_paintEvent_default")
	symaddr(&addr_QWizard_done_default, "QWizard_done_default")
	symaddr(&addr_QWizard_initializePage_default, "QWizard_initializePage_default")
	symaddr(&addr_QWizard_cleanupPage_default, "QWizard_cleanupPage_default")
	symaddr(&addr_QWizard_minimumSizeHint_default, "QWizard_minimumSizeHint_default")
	symaddr(&addr_QWizard_open_default, "QWizard_open_default")
	symaddr(&addr_QWizard_exec_default, "QWizard_exec_default")
	symaddr(&addr_QWizard_accept_default, "QWizard_accept_default")
	symaddr(&addr_QWizard_reject_default, "QWizard_reject_default")
	symaddr(&addr_QWizard_keyPressEvent_default, "QWizard_keyPressEvent_default")
	symaddr(&addr_QWizard_closeEvent_default, "QWizard_closeEvent_default")
	symaddr(&addr_QWizard_showEvent_default, "QWizard_showEvent_default")
	symaddr(&addr_QWizard_contextMenuEvent_default, "QWizard_contextMenuEvent_default")
	symaddr(&addr_QWizard_eventFilter_default, "QWizard_eventFilter_default")
	symaddr(&addr_QWizard_devType_default, "QWizard_devType_default")
	symaddr(&addr_QWizard_heightForWidth_default, "QWizard_heightForWidth_default")
	symaddr(&addr_QWizard_hasHeightForWidth_default, "QWizard_hasHeightForWidth_default")
	symaddr(&addr_QWizard_paintEngine_default, "QWizard_paintEngine_default")
	symaddr(&addr_QWizard_mousePressEvent_default, "QWizard_mousePressEvent_default")
	symaddr(&addr_QWizard_mouseReleaseEvent_default, "QWizard_mouseReleaseEvent_default")
	symaddr(&addr_QWizard_mouseDoubleClickEvent_default, "QWizard_mouseDoubleClickEvent_default")
	symaddr(&addr_QWizard_mouseMoveEvent_default, "QWizard_mouseMoveEvent_default")
	symaddr(&addr_QWizard_wheelEvent_default, "QWizard_wheelEvent_default")
	symaddr(&addr_QWizard_keyReleaseEvent_default, "QWizard_keyReleaseEvent_default")
	symaddr(&addr_QWizard_focusInEvent_default, "QWizard_focusInEvent_default")
	symaddr(&addr_QWizard_focusOutEvent_default, "QWizard_focusOutEvent_default")
	symaddr(&addr_QWizard_enterEvent_default, "QWizard_enterEvent_default")
	symaddr(&addr_QWizard_leaveEvent_default, "QWizard_leaveEvent_default")
	symaddr(&addr_QWizard_moveEvent_default, "QWizard_moveEvent_default")
	symaddr(&addr_QWizard_tabletEvent_default, "QWizard_tabletEvent_default")
	symaddr(&addr_QWizard_actionEvent_default, "QWizard_actionEvent_default")
	symaddr(&addr_QWizard_dragEnterEvent_default, "QWizard_dragEnterEvent_default")
	symaddr(&addr_QWizard_dragMoveEvent_default, "QWizard_dragMoveEvent_default")
	symaddr(&addr_QWizard_dragLeaveEvent_default, "QWizard_dragLeaveEvent_default")
	symaddr(&addr_QWizard_dropEvent_default, "QWizard_dropEvent_default")
	symaddr(&addr_QWizard_hideEvent_default, "QWizard_hideEvent_default")
	symaddr(&addr_QWizard_changeEvent_default, "QWizard_changeEvent_default")
	symaddr(&addr_QWizard_metric_default, "QWizard_metric_default")
	symaddr(&addr_QWizard_initPainter_default, "QWizard_initPainter_default")
	symaddr(&addr_QWizard_redirected_default, "QWizard_redirected_default")
	symaddr(&addr_QWizard_sharedPainter_default, "QWizard_sharedPainter_default")
	symaddr(&addr_QWizard_inputMethodEvent_default, "QWizard_inputMethodEvent_default")
	symaddr(&addr_QWizard_inputMethodQuery_default, "QWizard_inputMethodQuery_default")
	symaddr(&addr_QWizard_focusNextPrevChild_default, "QWizard_focusNextPrevChild_default")
	symaddr(&addr_QWizard_timerEvent_default, "QWizard_timerEvent_default")
	symaddr(&addr_QWizard_childEvent_default, "QWizard_childEvent_default")
	symaddr(&addr_QWizard_customEvent_default, "QWizard_customEvent_default")
	symaddr(&addr_QWizard_connectNotify_default, "QWizard_connectNotify_default")
	symaddr(&addr_QWizard_disconnectNotify_default, "QWizard_disconnectNotify_default")
	reg(&pg_QWizard_connect_currentIdChanged, "QWizard_connect_currentIdChanged")
	reg(&pg_QWizard_connect_helpRequested, "QWizard_connect_helpRequested")
	reg(&pg_QWizard_connect_customButtonClicked, "QWizard_connect_customButtonClicked")
	reg(&pg_QWizard_connect_pageAdded, "QWizard_connect_pageAdded")
	reg(&pg_QWizard_connect_pageRemoved, "QWizard_connect_pageRemoved")
	reg(&pg_QWizard_connect_finished, "QWizard_connect_finished")
	reg(&pg_QWizard_connect_accepted, "QWizard_connect_accepted")
	reg(&pg_QWizard_connect_rejected, "QWizard_connect_rejected")
	reg(&pg_QWizard_connect_windowTitleChanged, "QWizard_connect_windowTitleChanged")
	reg(&pg_QWizard_connect_windowIconChanged, "QWizard_connect_windowIconChanged")
	reg(&pg_QWizard_connect_windowIconTextChanged, "QWizard_connect_windowIconTextChanged")
	reg(&pg_QWizard_connect_customContextMenuRequested, "QWizard_connect_customContextMenuRequested")
	reg(&pg_QWizard_connect_destroyed, "QWizard_connect_destroyed")
	reg(&pg_QWizard_connect_destroyed2, "QWizard_connect_destroyed2")
	reg(&pg_QWizard_connect_objectNameChanged, "QWizard_connect_objectNameChanged")
	reg(&pg_QWizard__disconnect, "QWizard__disconnect")
	reg(&pg_QWizard__set_override, "QWizard__set_override")
	reg(&pg_QWizard_delete, "QWizard_delete")
}

type QWizard struct {
	h uintptr
	*QDialog
}

func bindQWizard(h uintptr) *QWizard {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QWizard_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QWizard{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQWizard(h uintptr) *QWizard { return bindQWizard(h) }

func (this *QWizard) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWizard) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWizard(h unsafe.Pointer) *QWizard { return bindQWizard(uintptr(h)) }
func UnsafeNewQWizardOwned(h unsafe.Pointer) *QWizard { return newownedQWizard(uintptr(h)) }

func (this *QWizard) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQWizard(parent *QWidget) *QWizard {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWizard_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWizard(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWizard_virtIdx, QWizard_enable)
	return __o
}

func NewQWizard2() *QWizard {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWizard_new2, uintptr(__goid))
	__o := newownedQWizard(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWizard_virtIdx, QWizard_enable)
	return __o
}

func NewQWizard3(parent *QWidget, flags WindowType) *QWizard {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWizard_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQWizard(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWizard_virtIdx, QWizard_enable)
	return __o
}

func QWizard_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWizard) AddPage(page *QWizardPage) int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_addPage, this.h, uintptr(page.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(page)
	return int(__r)
}

func (this *QWizard) SetPage(id int, page *QWizardPage) {
	purego.SyscallN(addr_QWizard_setPage, this.h, uintptr(id), uintptr(page.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(page)
}

func (this *QWizard) RemovePage(id int) {
	purego.SyscallN(addr_QWizard_removePage, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) Page(id int) *QWizardPage {
	__r, _, _ := purego.SyscallN(addr_QWizard_page, this.h, uintptr(id))
	runtime.KeepAlive(this)
	return bindQWizardPage(__r)
}

func (this *QWizard) HasVisitedPage(id int) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_hasVisitedPage, this.h, uintptr(id))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) VisitedIds() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_visitedIds, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QWizard) PageIds() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_pageIds, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QWizard) SetStartId(id int) {
	purego.SyscallN(addr_QWizard_setStartId, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) StartId() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_startId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) CurrentPage() *QWizardPage {
	__r, _, _ := purego.SyscallN(addr_QWizard_currentPage, this.h)
	runtime.KeepAlive(this)
	return bindQWizardPage(__r)
}

func (this *QWizard) CurrentId() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_currentId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) ValidateCurrentPage() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_validateCurrentPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) NextId() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_nextId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) SetField(name string, value *QVariant) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QWizard_setField, this.h, name_p, uintptr(len(name)), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QWizard) Field(name string) *QVariant {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QWizard_field, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QWizard) SetWizardStyle(style QWizard__WizardStyle) {
	purego.SyscallN(addr_QWizard_setWizardStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QWizard) WizardStyle() QWizard__WizardStyle {
	__r0, _, _ := purego.SyscallN(addr_QWizard_wizardStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QWizard__WizardStyle(__r)
}

func (this *QWizard) SetOption(option QWizard__WizardOption) {
	purego.SyscallN(addr_QWizard_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QWizard) TestOption(option QWizard__WizardOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) SetOptions(options QWizard__WizardOption) {
	purego.SyscallN(addr_QWizard_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QWizard) Options() QWizard__WizardOption {
	__r0, _, _ := purego.SyscallN(addr_QWizard_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QWizard__WizardOption(__r)
}

func (this *QWizard) SetButtonText(which QWizard__WizardButton, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QWizard_setButtonText, this.h, uintptr(which), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QWizard) ButtonText(which QWizard__WizardButton) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_buttonText, this.h, uintptr(which), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWizard) SetButtonLayout(layout []QWizard__WizardButton) {
	layout_arr := make([]int32, len(layout))
	for __i, __e := range layout {
		layout_arr[__i] = int32(__e)
	}
	var layout_p uintptr
	if len(layout_arr) > 0 {
		layout_p = uintptr(unsafe.Pointer(&layout_arr[0]))
	}
	purego.SyscallN(addr_QWizard_setButtonLayout, this.h, layout_p, uintptr(len(layout)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout_arr)
}

func (this *QWizard) SetButton(which QWizard__WizardButton, button *QAbstractButton) {
	purego.SyscallN(addr_QWizard_setButton, this.h, uintptr(which), uintptr(button.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QWizard) Button(which QWizard__WizardButton) *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QWizard_button, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QWizard) SetTitleFormat(format TextFormat) {
	purego.SyscallN(addr_QWizard_setTitleFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QWizard) TitleFormat() TextFormat {
	__r0, _, _ := purego.SyscallN(addr_QWizard_titleFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextFormat(__r)
}

func (this *QWizard) SetSubTitleFormat(format TextFormat) {
	purego.SyscallN(addr_QWizard_setSubTitleFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QWizard) SubTitleFormat() TextFormat {
	__r0, _, _ := purego.SyscallN(addr_QWizard_subTitleFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextFormat(__r)
}

func (this *QWizard) SetPixmap(which QWizard__WizardPixmap, pixmap *QPixmap) {
	purego.SyscallN(addr_QWizard_setPixmap, this.h, uintptr(which), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QWizard) Pixmap(which QWizard__WizardPixmap) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QWizard_pixmap, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QWizard) SetSideWidget(widget *QWidget) {
	purego.SyscallN(addr_QWizard_setSideWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QWizard) SideWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_sideWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	className_b := append([]byte(className), 0)
	property_b := append([]byte(property), 0)
	changedSignal_b := append([]byte(changedSignal), 0)
	purego.SyscallN(addr_QWizard_setDefaultProperty, this.h, uintptr(unsafe.Pointer(&className_b[0])), uintptr(unsafe.Pointer(&property_b[0])), uintptr(unsafe.Pointer(&changedSignal_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(className_b)
	runtime.KeepAlive(property_b)
	runtime.KeepAlive(changedSignal_b)
}

func (this *QWizard) CurrentIdChanged(id int) {
	purego.SyscallN(addr_QWizard_currentIdChanged, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) HelpRequested() {
	purego.SyscallN(addr_QWizard_helpRequested, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) CustomButtonClicked(which int) {
	purego.SyscallN(addr_QWizard_customButtonClicked, this.h, uintptr(which))
	runtime.KeepAlive(this)
}

func (this *QWizard) PageAdded(id int) {
	purego.SyscallN(addr_QWizard_pageAdded, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) PageRemoved(id int) {
	purego.SyscallN(addr_QWizard_pageRemoved, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) Back() {
	purego.SyscallN(addr_QWizard_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) Next() {
	purego.SyscallN(addr_QWizard_next, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) SetCurrentId(id int) {
	purego.SyscallN(addr_QWizard_setCurrentId, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) Restart() {
	purego.SyscallN(addr_QWizard_restart, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QWizard_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QWizard_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QWizard_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QWizard_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_mouseGrabber)
	return bindQWidget(__r)
}

func QWizard_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_keyboardGrabber)
	return bindQWidget(__r)
}

func QWizard_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QWizard_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QWizard) UpdateMicroFocus() {
	purego.SyscallN(addr_QWizard_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) Create() {
	purego.SyscallN(addr_QWizard_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) Destroy() {
	purego.SyscallN(addr_QWizard_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWizard_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWizard_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWizard_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWizard_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWizard) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWizard_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWizard) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWizard_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWizard) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QWizard_DevicePixelRatioFScale() float64 {
	__r := pg_QWizard_devicePixelRatioFScale()
	return __r
}

func QWizard_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QWizard_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QWizard) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QWizard_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QWizard_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWizard_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizard_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QWizard) SetOption2(option QWizard__WizardOption, on bool) {
	purego.SyscallN(addr_QWizard_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QWizard_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QWizard_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizard_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QWizard) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QWizard_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QWizard) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QWizard_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWizard) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QWizard_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizard) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QWizard_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizard) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QWizard_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizard) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QWizard_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QWizard_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWizard_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWizard_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWizard_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWizard) ValidateCurrentPageDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_validateCurrentPage_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) NextIdDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_nextId_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWizard_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWizard) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QWizard_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QWizard) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWizard_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWizard) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWizard) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QWizard_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QWizard_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) DoneDefault(result int) {
	purego.SyscallN(addr_QWizard_done_default, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QWizard) InitializePageDefault(id int) {
	purego.SyscallN(addr_QWizard_initializePage_default, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) CleanupPageDefault(id int) {
	purego.SyscallN(addr_QWizard_cleanupPage_default, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWizard) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWizard_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWizard) OpenDefault() {
	purego.SyscallN(addr_QWizard_open_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) AcceptDefault() {
	purego.SyscallN(addr_QWizard_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) RejectDefault() {
	purego.SyscallN(addr_QWizard_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizard) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QWizard_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QWizard_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QWizard_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QWizard_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QWizard) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QWizard_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QWizard) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizard_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizard_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizard_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizard_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QWizard_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QWizard_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWizard_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWizard_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QWizard_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWizard_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QWizard_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QWizard_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QWizard_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QWizard_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QWizard_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QWizard_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QWizard_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QWizard_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QWizard_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QWizard_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizard) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QWizard_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QWizard) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QWizard_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QWizard) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QWizard_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QWizard) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QWizard_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizard) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QWizard_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QWizard) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizard_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizard) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWizard_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWizard_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWizard_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizard) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWizard_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWizard) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWizard_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWizard) OnCurrentIdChanged(fn func(int)) int64 {
	id := registerCallback(func(id int32) { fn(int(id)) })
	pg_QWizard_connect_currentIdChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnHelpRequested(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizard_connect_helpRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnCustomButtonClicked(fn func(int)) int64 {
	id := registerCallback(func(which int32) { fn(int(which)) })
	pg_QWizard_connect_customButtonClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnPageAdded(fn func(int)) int64 {
	id := registerCallback(func(id int32) { fn(int(id)) })
	pg_QWizard_connect_pageAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnPageRemoved(fn func(int)) int64 {
	id := registerCallback(func(id int32) { fn(int(id)) })
	pg_QWizard_connect_pageRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QWizard_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizard_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizard_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QWizard_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QWizard_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QWizard_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QWizard_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizard_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWizard_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWizard_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizard) Disconnect(handle int64) {
	pg_QWizard__disconnect(handle)
	unregisterCallback(handle)
}

func QWizard_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7851200041405997743:
		return 1
	case 5814903794109731573:
		return 2
	case 2110182719375245575:
		return 3
	case 729690261457684101:
		return 4
	case 1237236456312824991:
		return 5
	case -4302917285655014661:
		return 6
	case 5290442163167458465:
		return 7
	case -2039166586246443706:
		return 8
	case -2570997490807640431:
		return 9
	case 3074997362940195296:
		return 10
	case 5913490871578514792:
		return 11
	case 7049264599837532491:
		return 12
	case -554057428128861975:
		return 13
	case -1745322561779819684:
		return 14
	case 9106131928185084681:
		return 15
	case 2875220963396170150:
		return 16
	case 1023791977679793185:
		return 17
	case -1989937050019464783:
		return 18
	case -6825369174751916:
		return 19
	case -6698043222227006015:
		return 20
	case 5217623514630858297:
		return 21
	case -8659679893868524044:
		return 22
	case 4311534060495737339:
		return 23
	case -8266762397995423915:
		return 24
	case 6398007404756346381:
		return 25
	case -332090898078543319:
		return 26
	case -4415920049297795213:
		return 27
	case -4029073905697023887:
		return 28
	case -4255494004099806581:
		return 29
	case 1919988340934511316:
		return 30
	case 772598475950050299:
		return 31
	case 975043767033552552:
		return 32
	case -385596010576048993:
		return 33
	case 6316270346668237895:
		return 34
	case -551070009623711640:
		return 35
	case -9161949150489525396:
		return 36
	case -212068866860701141:
		return 37
	case 1697444140088555581:
		return 38
	case 206122715919917589:
		return 39
	case -1781545137352745978:
		return 40
	case -3139904158373725006:
		return 41
	case 9107058682754591716:
		return 42
	case -199718707876719129:
		return 43
	case 4959376063826165943:
		return 44
	case 427629649267311777:
		return 45
	case 6477500296980050210:
		return 46
	case 4127240633310655048:
		return 47
	case 5792771328741774991:
		return 48
	case 5438819661942843222:
		return 49
	case 7345853097340709638:
		return 50
	case 5898328517505784191:
		return 51
	case -5409927581812374498:
		return 52
	case -1185315367526857833:
		return 53
	case 3769868229123868474:
		return 54
	case 7379029982170503008:
		return 55
	case -568369772982215962:
		return 56
	}
	return -1
}

func QWizard_enable(h uintptr, idx int32) { pg_QWizard__set_override(h, idx) }

func (this *QWizard) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWizard) OnValidateCurrentPage(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 7851200041405997743, func() int32 { return b2i(fn()) }) // slot "validateCurrentPage"
}

func (this *QWizard) OnNextId(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5814903794109731573, func() int32 { return int32(fn()) }) // slot "nextId"
}

func (this *QWizard) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QWizard) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QWizard) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QWizard) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QWizard) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QWizard) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QWizard) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(result int32) { fn(int(result)) }) // slot "done"
}

func (this *QWizard) OnInitializePage(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 3074997362940195296, func(id int32) { fn(int(id)) }) // slot "initializePage"
}

func (this *QWizard) OnCleanupPage(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 5913490871578514792, func(id int32) { fn(int(id)) }) // slot "cleanupPage"
}

func (this *QWizard) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QWizard) OnOpen(fn func()) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func() { fn() }) // slot "open"
}

func (this *QWizard) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QWizard) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QWizard) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QWizard) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QWizard) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QWizard) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QWizard) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QWizard) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QWizard) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QWizard) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QWizard) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QWizard) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QWizard) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QWizard) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QWizard) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QWizard) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QWizard) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QWizard) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QWizard) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QWizard) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QWizard) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QWizard) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QWizard) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QWizard) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QWizard) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QWizard) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QWizard) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QWizard) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QWizard) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QWizard) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QWizard) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QWizard) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QWizard) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QWizard) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QWizard) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QWizard) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QWizard) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QWizard) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QWizard) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWizard) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWizard) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWizard) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWizard) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWizard) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWizard_delete(this.h)
	}
	this.UnsafeRelease()
}
