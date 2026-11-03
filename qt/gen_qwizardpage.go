package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWizardPage_chain uintptr
	addr_QWizardPage_new uintptr
	addr_QWizardPage_new2 uintptr
	addr_QWizardPage_tr uintptr
	addr_QWizardPage_setTitle uintptr
	addr_QWizardPage_title uintptr
	addr_QWizardPage_setSubTitle uintptr
	addr_QWizardPage_subTitle uintptr
	addr_QWizardPage_setPixmap uintptr
	addr_QWizardPage_pixmap uintptr
	addr_QWizardPage_setFinalPage uintptr
	addr_QWizardPage_isFinalPage uintptr
	addr_QWizardPage_setCommitPage uintptr
	addr_QWizardPage_isCommitPage uintptr
	addr_QWizardPage_setButtonText uintptr
	addr_QWizardPage_buttonText uintptr
	addr_QWizardPage_initializePage uintptr
	addr_QWizardPage_cleanupPage uintptr
	addr_QWizardPage_validatePage uintptr
	addr_QWizardPage_isComplete uintptr
	addr_QWizardPage_nextId uintptr
	addr_QWizardPage_completeChanged uintptr
	addr_QWizardPage_setField uintptr
	addr_QWizardPage_field uintptr
	addr_QWizardPage_registerField uintptr
	addr_QWizardPage_wizard uintptr
	addr_QWizardPage_setTabOrder uintptr
	addr_QWizardPage_mouseGrabber uintptr
	addr_QWizardPage_keyboardGrabber uintptr
	addr_QWizardPage_find uintptr
	addr_QWizardPage_createWindowContainer uintptr
	addr_QWizardPage_updateMicroFocus uintptr
	addr_QWizardPage_create uintptr
	addr_QWizardPage_destroy uintptr
	addr_QWizardPage_focusNextChild uintptr
	addr_QWizardPage_focusPreviousChild uintptr
	addr_QWizardPage_connect uintptr
	addr_QWizardPage_connect2 uintptr
	addr_QWizardPage_sender uintptr
	addr_QWizardPage_senderSignalIndex uintptr
	addr_QWizardPage_receivers uintptr
	addr_QWizardPage_isSignalConnected uintptr
	pg_QWizardPage_devicePixelRatioFScale func() float64
	pg_QWizardPage_encodeMetricF func(int32, float64) int32
	pg_QWizardPage_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QWizardPage_tr2 uintptr
	addr_QWizardPage_tr3 uintptr
	addr_QWizardPage_registerField2 uintptr
	addr_QWizardPage_registerField3 uintptr
	addr_QWizardPage_createWindowContainer2 uintptr
	addr_QWizardPage_createWindowContainer3 uintptr
	addr_QWizardPage_updateMicroFocus2 uintptr
	addr_QWizardPage_create2 uintptr
	addr_QWizardPage_create3 uintptr
	addr_QWizardPage_create4 uintptr
	addr_QWizardPage_destroy2 uintptr
	addr_QWizardPage_destroy3 uintptr
	addr_QWizardPage_connect3 uintptr
	addr_QWizardPage_connect4 uintptr
	addr_QWizardPage_initializePage_default uintptr
	addr_QWizardPage_cleanupPage_default uintptr
	addr_QWizardPage_validatePage_default uintptr
	addr_QWizardPage_isComplete_default uintptr
	addr_QWizardPage_nextId_default uintptr
	addr_QWizardPage_metaObject_default uintptr
	addr_QWizardPage_devType_default uintptr
	addr_QWizardPage_setVisible_default uintptr
	addr_QWizardPage_sizeHint_default uintptr
	addr_QWizardPage_minimumSizeHint_default uintptr
	addr_QWizardPage_heightForWidth_default uintptr
	addr_QWizardPage_hasHeightForWidth_default uintptr
	addr_QWizardPage_paintEngine_default uintptr
	addr_QWizardPage_event_default uintptr
	addr_QWizardPage_mousePressEvent_default uintptr
	addr_QWizardPage_mouseReleaseEvent_default uintptr
	addr_QWizardPage_mouseDoubleClickEvent_default uintptr
	addr_QWizardPage_mouseMoveEvent_default uintptr
	addr_QWizardPage_wheelEvent_default uintptr
	addr_QWizardPage_keyPressEvent_default uintptr
	addr_QWizardPage_keyReleaseEvent_default uintptr
	addr_QWizardPage_focusInEvent_default uintptr
	addr_QWizardPage_focusOutEvent_default uintptr
	addr_QWizardPage_enterEvent_default uintptr
	addr_QWizardPage_leaveEvent_default uintptr
	addr_QWizardPage_paintEvent_default uintptr
	addr_QWizardPage_moveEvent_default uintptr
	addr_QWizardPage_resizeEvent_default uintptr
	addr_QWizardPage_closeEvent_default uintptr
	addr_QWizardPage_contextMenuEvent_default uintptr
	addr_QWizardPage_tabletEvent_default uintptr
	addr_QWizardPage_actionEvent_default uintptr
	addr_QWizardPage_dragEnterEvent_default uintptr
	addr_QWizardPage_dragMoveEvent_default uintptr
	addr_QWizardPage_dragLeaveEvent_default uintptr
	addr_QWizardPage_dropEvent_default uintptr
	addr_QWizardPage_showEvent_default uintptr
	addr_QWizardPage_hideEvent_default uintptr
	addr_QWizardPage_changeEvent_default uintptr
	addr_QWizardPage_metric_default uintptr
	addr_QWizardPage_initPainter_default uintptr
	addr_QWizardPage_redirected_default uintptr
	addr_QWizardPage_sharedPainter_default uintptr
	addr_QWizardPage_inputMethodEvent_default uintptr
	addr_QWizardPage_inputMethodQuery_default uintptr
	addr_QWizardPage_focusNextPrevChild_default uintptr
	addr_QWizardPage_eventFilter_default uintptr
	addr_QWizardPage_timerEvent_default uintptr
	addr_QWizardPage_childEvent_default uintptr
	addr_QWizardPage_customEvent_default uintptr
	addr_QWizardPage_connectNotify_default uintptr
	addr_QWizardPage_disconnectNotify_default uintptr
	pg_QWizardPage_connect_completeChanged func(uintptr, int64)
	pg_QWizardPage_connect_windowTitleChanged func(uintptr, int64)
	pg_QWizardPage_connect_windowIconChanged func(uintptr, int64)
	pg_QWizardPage_connect_windowIconTextChanged func(uintptr, int64)
	pg_QWizardPage_connect_customContextMenuRequested func(uintptr, int64)
	pg_QWizardPage_connect_destroyed func(uintptr, int64)
	pg_QWizardPage_connect_destroyed2 func(uintptr, int64)
	pg_QWizardPage_connect_objectNameChanged func(uintptr, int64)
	pg_QWizardPage__disconnect func(int64)
	pg_QWizardPage__set_override func(uintptr, int32)
	pg_QWizardPage_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWizardPage_chain, "QWizardPage_chain")
	symaddr(&addr_QWizardPage_new, "QWizardPage_new")
	symaddr(&addr_QWizardPage_new2, "QWizardPage_new2")
	symaddr(&addr_QWizardPage_tr, "QWizardPage_tr")
	symaddr(&addr_QWizardPage_setTitle, "QWizardPage_setTitle")
	symaddr(&addr_QWizardPage_title, "QWizardPage_title")
	symaddr(&addr_QWizardPage_setSubTitle, "QWizardPage_setSubTitle")
	symaddr(&addr_QWizardPage_subTitle, "QWizardPage_subTitle")
	symaddr(&addr_QWizardPage_setPixmap, "QWizardPage_setPixmap")
	symaddr(&addr_QWizardPage_pixmap, "QWizardPage_pixmap")
	symaddr(&addr_QWizardPage_setFinalPage, "QWizardPage_setFinalPage")
	symaddr(&addr_QWizardPage_isFinalPage, "QWizardPage_isFinalPage")
	symaddr(&addr_QWizardPage_setCommitPage, "QWizardPage_setCommitPage")
	symaddr(&addr_QWizardPage_isCommitPage, "QWizardPage_isCommitPage")
	symaddr(&addr_QWizardPage_setButtonText, "QWizardPage_setButtonText")
	symaddr(&addr_QWizardPage_buttonText, "QWizardPage_buttonText")
	symaddr(&addr_QWizardPage_initializePage, "QWizardPage_initializePage")
	symaddr(&addr_QWizardPage_cleanupPage, "QWizardPage_cleanupPage")
	symaddr(&addr_QWizardPage_validatePage, "QWizardPage_validatePage")
	symaddr(&addr_QWizardPage_isComplete, "QWizardPage_isComplete")
	symaddr(&addr_QWizardPage_nextId, "QWizardPage_nextId")
	symaddr(&addr_QWizardPage_completeChanged, "QWizardPage_completeChanged")
	symaddr(&addr_QWizardPage_setField, "QWizardPage_setField")
	symaddr(&addr_QWizardPage_field, "QWizardPage_field")
	symaddr(&addr_QWizardPage_registerField, "QWizardPage_registerField")
	symaddr(&addr_QWizardPage_wizard, "QWizardPage_wizard")
	symaddr(&addr_QWizardPage_setTabOrder, "QWizardPage_setTabOrder")
	symaddr(&addr_QWizardPage_mouseGrabber, "QWizardPage_mouseGrabber")
	symaddr(&addr_QWizardPage_keyboardGrabber, "QWizardPage_keyboardGrabber")
	symaddr(&addr_QWizardPage_find, "QWizardPage_find")
	symaddr(&addr_QWizardPage_createWindowContainer, "QWizardPage_createWindowContainer")
	symaddr(&addr_QWizardPage_updateMicroFocus, "QWizardPage_updateMicroFocus")
	symaddr(&addr_QWizardPage_create, "QWizardPage_create")
	symaddr(&addr_QWizardPage_destroy, "QWizardPage_destroy")
	symaddr(&addr_QWizardPage_focusNextChild, "QWizardPage_focusNextChild")
	symaddr(&addr_QWizardPage_focusPreviousChild, "QWizardPage_focusPreviousChild")
	symaddr(&addr_QWizardPage_connect, "QWizardPage_connect")
	symaddr(&addr_QWizardPage_connect2, "QWizardPage_connect2")
	symaddr(&addr_QWizardPage_sender, "QWizardPage_sender")
	symaddr(&addr_QWizardPage_senderSignalIndex, "QWizardPage_senderSignalIndex")
	symaddr(&addr_QWizardPage_receivers, "QWizardPage_receivers")
	symaddr(&addr_QWizardPage_isSignalConnected, "QWizardPage_isSignalConnected")
	reg(&pg_QWizardPage_devicePixelRatioFScale, "QWizardPage_devicePixelRatioFScale")
	reg(&pg_QWizardPage_encodeMetricF, "QWizardPage_encodeMetricF")
	reg(&pg_QWizardPage_getDecodedMetricF, "QWizardPage_getDecodedMetricF")
	symaddr(&addr_QWizardPage_tr2, "QWizardPage_tr2")
	symaddr(&addr_QWizardPage_tr3, "QWizardPage_tr3")
	symaddr(&addr_QWizardPage_registerField2, "QWizardPage_registerField2")
	symaddr(&addr_QWizardPage_registerField3, "QWizardPage_registerField3")
	symaddr(&addr_QWizardPage_createWindowContainer2, "QWizardPage_createWindowContainer2")
	symaddr(&addr_QWizardPage_createWindowContainer3, "QWizardPage_createWindowContainer3")
	symaddr(&addr_QWizardPage_updateMicroFocus2, "QWizardPage_updateMicroFocus2")
	symaddr(&addr_QWizardPage_create2, "QWizardPage_create2")
	symaddr(&addr_QWizardPage_create3, "QWizardPage_create3")
	symaddr(&addr_QWizardPage_create4, "QWizardPage_create4")
	symaddr(&addr_QWizardPage_destroy2, "QWizardPage_destroy2")
	symaddr(&addr_QWizardPage_destroy3, "QWizardPage_destroy3")
	symaddr(&addr_QWizardPage_connect3, "QWizardPage_connect3")
	symaddr(&addr_QWizardPage_connect4, "QWizardPage_connect4")
	symaddr(&addr_QWizardPage_initializePage_default, "QWizardPage_initializePage_default")
	symaddr(&addr_QWizardPage_cleanupPage_default, "QWizardPage_cleanupPage_default")
	symaddr(&addr_QWizardPage_validatePage_default, "QWizardPage_validatePage_default")
	symaddr(&addr_QWizardPage_isComplete_default, "QWizardPage_isComplete_default")
	symaddr(&addr_QWizardPage_nextId_default, "QWizardPage_nextId_default")
	symaddr(&addr_QWizardPage_metaObject_default, "QWizardPage_metaObject_default")
	symaddr(&addr_QWizardPage_devType_default, "QWizardPage_devType_default")
	symaddr(&addr_QWizardPage_setVisible_default, "QWizardPage_setVisible_default")
	symaddr(&addr_QWizardPage_sizeHint_default, "QWizardPage_sizeHint_default")
	symaddr(&addr_QWizardPage_minimumSizeHint_default, "QWizardPage_minimumSizeHint_default")
	symaddr(&addr_QWizardPage_heightForWidth_default, "QWizardPage_heightForWidth_default")
	symaddr(&addr_QWizardPage_hasHeightForWidth_default, "QWizardPage_hasHeightForWidth_default")
	symaddr(&addr_QWizardPage_paintEngine_default, "QWizardPage_paintEngine_default")
	symaddr(&addr_QWizardPage_event_default, "QWizardPage_event_default")
	symaddr(&addr_QWizardPage_mousePressEvent_default, "QWizardPage_mousePressEvent_default")
	symaddr(&addr_QWizardPage_mouseReleaseEvent_default, "QWizardPage_mouseReleaseEvent_default")
	symaddr(&addr_QWizardPage_mouseDoubleClickEvent_default, "QWizardPage_mouseDoubleClickEvent_default")
	symaddr(&addr_QWizardPage_mouseMoveEvent_default, "QWizardPage_mouseMoveEvent_default")
	symaddr(&addr_QWizardPage_wheelEvent_default, "QWizardPage_wheelEvent_default")
	symaddr(&addr_QWizardPage_keyPressEvent_default, "QWizardPage_keyPressEvent_default")
	symaddr(&addr_QWizardPage_keyReleaseEvent_default, "QWizardPage_keyReleaseEvent_default")
	symaddr(&addr_QWizardPage_focusInEvent_default, "QWizardPage_focusInEvent_default")
	symaddr(&addr_QWizardPage_focusOutEvent_default, "QWizardPage_focusOutEvent_default")
	symaddr(&addr_QWizardPage_enterEvent_default, "QWizardPage_enterEvent_default")
	symaddr(&addr_QWizardPage_leaveEvent_default, "QWizardPage_leaveEvent_default")
	symaddr(&addr_QWizardPage_paintEvent_default, "QWizardPage_paintEvent_default")
	symaddr(&addr_QWizardPage_moveEvent_default, "QWizardPage_moveEvent_default")
	symaddr(&addr_QWizardPage_resizeEvent_default, "QWizardPage_resizeEvent_default")
	symaddr(&addr_QWizardPage_closeEvent_default, "QWizardPage_closeEvent_default")
	symaddr(&addr_QWizardPage_contextMenuEvent_default, "QWizardPage_contextMenuEvent_default")
	symaddr(&addr_QWizardPage_tabletEvent_default, "QWizardPage_tabletEvent_default")
	symaddr(&addr_QWizardPage_actionEvent_default, "QWizardPage_actionEvent_default")
	symaddr(&addr_QWizardPage_dragEnterEvent_default, "QWizardPage_dragEnterEvent_default")
	symaddr(&addr_QWizardPage_dragMoveEvent_default, "QWizardPage_dragMoveEvent_default")
	symaddr(&addr_QWizardPage_dragLeaveEvent_default, "QWizardPage_dragLeaveEvent_default")
	symaddr(&addr_QWizardPage_dropEvent_default, "QWizardPage_dropEvent_default")
	symaddr(&addr_QWizardPage_showEvent_default, "QWizardPage_showEvent_default")
	symaddr(&addr_QWizardPage_hideEvent_default, "QWizardPage_hideEvent_default")
	symaddr(&addr_QWizardPage_changeEvent_default, "QWizardPage_changeEvent_default")
	symaddr(&addr_QWizardPage_metric_default, "QWizardPage_metric_default")
	symaddr(&addr_QWizardPage_initPainter_default, "QWizardPage_initPainter_default")
	symaddr(&addr_QWizardPage_redirected_default, "QWizardPage_redirected_default")
	symaddr(&addr_QWizardPage_sharedPainter_default, "QWizardPage_sharedPainter_default")
	symaddr(&addr_QWizardPage_inputMethodEvent_default, "QWizardPage_inputMethodEvent_default")
	symaddr(&addr_QWizardPage_inputMethodQuery_default, "QWizardPage_inputMethodQuery_default")
	symaddr(&addr_QWizardPage_focusNextPrevChild_default, "QWizardPage_focusNextPrevChild_default")
	symaddr(&addr_QWizardPage_eventFilter_default, "QWizardPage_eventFilter_default")
	symaddr(&addr_QWizardPage_timerEvent_default, "QWizardPage_timerEvent_default")
	symaddr(&addr_QWizardPage_childEvent_default, "QWizardPage_childEvent_default")
	symaddr(&addr_QWizardPage_customEvent_default, "QWizardPage_customEvent_default")
	symaddr(&addr_QWizardPage_connectNotify_default, "QWizardPage_connectNotify_default")
	symaddr(&addr_QWizardPage_disconnectNotify_default, "QWizardPage_disconnectNotify_default")
	reg(&pg_QWizardPage_connect_completeChanged, "QWizardPage_connect_completeChanged")
	reg(&pg_QWizardPage_connect_windowTitleChanged, "QWizardPage_connect_windowTitleChanged")
	reg(&pg_QWizardPage_connect_windowIconChanged, "QWizardPage_connect_windowIconChanged")
	reg(&pg_QWizardPage_connect_windowIconTextChanged, "QWizardPage_connect_windowIconTextChanged")
	reg(&pg_QWizardPage_connect_customContextMenuRequested, "QWizardPage_connect_customContextMenuRequested")
	reg(&pg_QWizardPage_connect_destroyed, "QWizardPage_connect_destroyed")
	reg(&pg_QWizardPage_connect_destroyed2, "QWizardPage_connect_destroyed2")
	reg(&pg_QWizardPage_connect_objectNameChanged, "QWizardPage_connect_objectNameChanged")
	reg(&pg_QWizardPage__disconnect, "QWizardPage__disconnect")
	reg(&pg_QWizardPage__set_override, "QWizardPage__set_override")
	reg(&pg_QWizardPage_delete, "QWizardPage_delete")
}

type QWizardPage struct {
	h uintptr
	*QWidget
}

func bindQWizardPage(h uintptr) *QWizardPage {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QWizardPage_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QWizardPage{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQWizardPage(h uintptr) *QWizardPage { return bindQWizardPage(h) }

func (this *QWizardPage) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWizardPage) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWizardPage(h unsafe.Pointer) *QWizardPage { return bindQWizardPage(uintptr(h)) }
func UnsafeNewQWizardPageOwned(h unsafe.Pointer) *QWizardPage { return newownedQWizardPage(uintptr(h)) }

func (this *QWizardPage) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQWizardPage(parent *QWidget) *QWizardPage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWizardPage_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWizardPage(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWizardPage_virtIdx, QWizardPage_enable)
	return __o
}

func NewQWizardPage2() *QWizardPage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWizardPage_new2, uintptr(__goid))
	__o := newownedQWizardPage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWizardPage_virtIdx, QWizardPage_enable)
	return __o
}

func QWizardPage_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWizardPage) SetTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QWizardPage_setTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QWizardPage) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_title, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWizardPage) SetSubTitle(subTitle string) {
	var subTitle_p uintptr
	if len(subTitle) > 0 {
		subTitle_p = uintptr(unsafe.Pointer(unsafe.StringData(subTitle)))
	}
	purego.SyscallN(addr_QWizardPage_setSubTitle, this.h, subTitle_p, uintptr(len(subTitle)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subTitle)
}

func (this *QWizardPage) SubTitle() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_subTitle, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap *QPixmap) {
	purego.SyscallN(addr_QWizardPage_setPixmap, this.h, uintptr(which), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QWizardPage) Pixmap(which QWizard__WizardPixmap) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_pixmap, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QWizardPage) SetFinalPage(finalPage bool) {
	purego.SyscallN(addr_QWizardPage_setFinalPage, this.h, uintptr(b2i(finalPage)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) IsFinalPage() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_isFinalPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) SetCommitPage(commitPage bool) {
	purego.SyscallN(addr_QWizardPage_setCommitPage, this.h, uintptr(b2i(commitPage)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) IsCommitPage() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_isCommitPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QWizardPage_setButtonText, this.h, uintptr(which), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_buttonText, this.h, uintptr(which), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWizardPage) InitializePage() {
	purego.SyscallN(addr_QWizardPage_initializePage, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) CleanupPage() {
	purego.SyscallN(addr_QWizardPage_cleanupPage, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) ValidatePage() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_validatePage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) IsComplete() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_isComplete, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) NextId() int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_nextId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) CompleteChanged() {
	purego.SyscallN(addr_QWizardPage_completeChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) SetField(name string, value *QVariant) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QWizardPage_setField, this.h, name_p, uintptr(len(name)), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QWizardPage) Field(name string) *QVariant {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QWizardPage_field, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return newownedQVariant(__r)
}

func (this *QWizardPage) RegisterField(name string, widget *QWidget) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QWizardPage_registerField, this.h, name_p, uintptr(len(name)), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(widget)
}

func (this *QWizardPage) Wizard() *QWizard {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_wizard, this.h)
	runtime.KeepAlive(this)
	return bindQWizard(__r)
}

func QWizardPage_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QWizardPage_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QWizardPage_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_mouseGrabber)
	return bindQWidget(__r)
}

func QWizardPage_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_keyboardGrabber)
	return bindQWidget(__r)
}

func QWizardPage_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QWizardPage_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QWizardPage) UpdateMicroFocus() {
	purego.SyscallN(addr_QWizardPage_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) Create() {
	purego.SyscallN(addr_QWizardPage_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) Destroy() {
	purego.SyscallN(addr_QWizardPage_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWizardPage_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWizardPage_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWizardPage_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWizardPage) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWizardPage) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWizardPage) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QWizardPage_DevicePixelRatioFScale() float64 {
	__r := pg_QWizardPage_devicePixelRatioFScale()
	return __r
}

func QWizardPage_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QWizardPage_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QWizardPage) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QWizardPage_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QWizardPage_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWizardPage_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWizardPage_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QWizardPage) RegisterField2(name string, widget *QWidget, property string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	property_b := append([]byte(property), 0)
	purego.SyscallN(addr_QWizardPage_registerField2, this.h, name_p, uintptr(len(name)), uintptr(widget.UnsafePointer()), uintptr(unsafe.Pointer(&property_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(property_b)
}

func (this *QWizardPage) RegisterField3(name string, widget *QWidget, property string, changedSignal string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	property_b := append([]byte(property), 0)
	changedSignal_b := append([]byte(changedSignal), 0)
	purego.SyscallN(addr_QWizardPage_registerField3, this.h, name_p, uintptr(len(name)), uintptr(widget.UnsafePointer()), uintptr(unsafe.Pointer(&property_b[0])), uintptr(unsafe.Pointer(&changedSignal_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(property_b)
	runtime.KeepAlive(changedSignal_b)
}

func QWizardPage_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QWizardPage_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QWizardPage) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QWizardPage_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QWizardPage_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QWizardPage_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QWizardPage_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QWizardPage_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QWizardPage_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QWizardPage_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWizardPage_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWizardPage_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWizardPage) InitializePageDefault() {
	purego.SyscallN(addr_QWizardPage_initializePage_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) CleanupPageDefault() {
	purego.SyscallN(addr_QWizardPage_cleanupPage_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QWizardPage) ValidatePageDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_validatePage_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) IsCompleteDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_isComplete_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) NextIdDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_nextId_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWizardPage) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QWizardPage_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QWizardPage) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWizardPage) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWizardPage) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QWizardPage) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWizardPage) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizardPage_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizardPage_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizardPage_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWizardPage_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QWizardPage_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QWizardPage_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QWizardPage_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWizardPage_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWizardPage_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QWizardPage_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWizardPage_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QWizardPage_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QWizardPage_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QWizardPage_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QWizardPage_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QWizardPage_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QWizardPage_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QWizardPage_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QWizardPage_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QWizardPage_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QWizardPage_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QWizardPage_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QWizardPage_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QWizardPage_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QWizardPage_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizardPage) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWizardPage) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QWizardPage_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QWizardPage) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QWizardPage) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QWizardPage) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QWizardPage_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWizardPage) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QWizardPage_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QWizardPage) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWizardPage) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWizardPage_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWizardPage) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWizardPage_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWizardPage_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWizardPage_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWizardPage) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWizardPage_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWizardPage) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWizardPage_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWizardPage) OnCompleteChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizardPage_connect_completeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QWizardPage_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QWizardPage_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QWizardPage_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QWizardPage_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWizardPage_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWizardPage_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWizardPage_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWizardPage) Disconnect(handle int64) {
	pg_QWizardPage__disconnect(handle)
	unregisterCallback(handle)
}

func QWizardPage_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 3074997362940195296:
		return 1
	case 5913490871578514792:
		return 2
	case -4021270481451057600:
		return 3
	case 9144100205243375250:
		return 4
	case 5814903794109731573:
		return 5
	case -8659679893868524044:
		return 6
	case 2110182719375245575:
		return 7
	case 729690261457684101:
		return 8
	case 7049264599837532491:
		return 9
	case 4311534060495737339:
		return 10
	case -8266762397995423915:
		return 11
	case 6398007404756346381:
		return 12
	case 1237236456312824991:
		return 13
	case -332090898078543319:
		return 14
	case -4415920049297795213:
		return 15
	case -4029073905697023887:
		return 16
	case -4255494004099806581:
		return 17
	case 1919988340934511316:
		return 18
	case 1023791977679793185:
		return 19
	case 772598475950050299:
		return 20
	case 975043767033552552:
		return 21
	case -385596010576048993:
		return 22
	case 6316270346668237895:
		return 23
	case -551070009623711640:
		return 24
	case 5290442163167458465:
		return 25
	case -9161949150489525396:
		return 26
	case -4302917285655014661:
		return 27
	case -1989937050019464783:
		return 28
	case -6698043222227006015:
		return 29
	case -212068866860701141:
		return 30
	case 1697444140088555581:
		return 31
	case 206122715919917589:
		return 32
	case -1781545137352745978:
		return 33
	case -3139904158373725006:
		return 34
	case 9107058682754591716:
		return 35
	case -6825369174751916:
		return 36
	case -199718707876719129:
		return 37
	case -2039166586246443706:
		return 38
	case 4959376063826165943:
		return 39
	case 427629649267311777:
		return 40
	case 6477500296980050210:
		return 41
	case 4127240633310655048:
		return 42
	case 5792771328741774991:
		return 43
	case 5438819661942843222:
		return 44
	case 7345853097340709638:
		return 45
	case 5898328517505784191:
		return 46
	case 5217623514630858297:
		return 47
	case -5409927581812374498:
		return 48
	case -1185315367526857833:
		return 49
	case 3769868229123868474:
		return 50
	case 7379029982170503008:
		return 51
	case -568369772982215962:
		return 52
	}
	return -1
}

func QWizardPage_enable(h uintptr, idx int32) { pg_QWizardPage__set_override(h, idx) }

func (this *QWizardPage) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWizardPage) OnInitializePage(fn func()) {
	registerVirt(this.UnsafeGoID(), 3074997362940195296, func() { fn() }) // slot "initializePage"
}

func (this *QWizardPage) OnCleanupPage(fn func()) {
	registerVirt(this.UnsafeGoID(), 5913490871578514792, func() { fn() }) // slot "cleanupPage"
}

func (this *QWizardPage) OnValidatePage(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -4021270481451057600, func() int32 { return b2i(fn()) }) // slot "validatePage"
}

func (this *QWizardPage) OnIsComplete(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 9144100205243375250, func() int32 { return b2i(fn()) }) // slot "isComplete"
}

func (this *QWizardPage) OnNextId(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5814903794109731573, func() int32 { return int32(fn()) }) // slot "nextId"
}

func (this *QWizardPage) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QWizardPage) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QWizardPage) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QWizardPage) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QWizardPage) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QWizardPage) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QWizardPage) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QWizardPage) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QWizardPage) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QWizardPage) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QWizardPage) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QWizardPage) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QWizardPage) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QWizardPage) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QWizardPage) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QWizardPage) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QWizardPage) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QWizardPage) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QWizardPage) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QWizardPage) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QWizardPage) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QWizardPage) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QWizardPage) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QWizardPage) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QWizardPage) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QWizardPage) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QWizardPage) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QWizardPage) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QWizardPage) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QWizardPage) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QWizardPage) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QWizardPage) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QWizardPage) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QWizardPage) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QWizardPage) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QWizardPage) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QWizardPage) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QWizardPage) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QWizardPage) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QWizardPage) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QWizardPage) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QWizardPage) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QWizardPage) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWizardPage) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWizardPage) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWizardPage) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWizardPage) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWizardPage) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWizardPage_delete(this.h)
	}
	this.UnsafeRelease()
}
