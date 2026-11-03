package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMdiArea_chain uintptr
	addr_QMdiArea_new uintptr
	addr_QMdiArea_new2 uintptr
	addr_QMdiArea_tr uintptr
	addr_QMdiArea_currentSubWindow uintptr
	addr_QMdiArea_activeSubWindow uintptr
	addr_QMdiArea_subWindowList uintptr
	addr_QMdiArea_addSubWindow uintptr
	addr_QMdiArea_removeSubWindow uintptr
	addr_QMdiArea_background uintptr
	addr_QMdiArea_setBackground uintptr
	addr_QMdiArea_activationOrder uintptr
	addr_QMdiArea_setActivationOrder uintptr
	addr_QMdiArea_setOption uintptr
	addr_QMdiArea_testOption uintptr
	addr_QMdiArea_setViewMode uintptr
	addr_QMdiArea_viewMode uintptr
	addr_QMdiArea_documentMode uintptr
	addr_QMdiArea_setDocumentMode uintptr
	addr_QMdiArea_setTabsClosable uintptr
	addr_QMdiArea_tabsClosable uintptr
	addr_QMdiArea_setTabsMovable uintptr
	addr_QMdiArea_tabsMovable uintptr
	addr_QMdiArea_setTabShape uintptr
	addr_QMdiArea_tabShape uintptr
	addr_QMdiArea_setTabPosition uintptr
	addr_QMdiArea_tabPosition uintptr
	addr_QMdiArea_subWindowActivated uintptr
	addr_QMdiArea_setActiveSubWindow uintptr
	addr_QMdiArea_tileSubWindows uintptr
	addr_QMdiArea_cascadeSubWindows uintptr
	addr_QMdiArea_closeActiveSubWindow uintptr
	addr_QMdiArea_closeAllSubWindows uintptr
	addr_QMdiArea_activateNextSubWindow uintptr
	addr_QMdiArea_activatePreviousSubWindow uintptr
	addr_QMdiArea_setViewportMargins uintptr
	addr_QMdiArea_setViewportMargins2 uintptr
	addr_QMdiArea_viewportMargins uintptr
	addr_QMdiArea_drawFrame uintptr
	addr_QMdiArea_setTabOrder uintptr
	addr_QMdiArea_mouseGrabber uintptr
	addr_QMdiArea_keyboardGrabber uintptr
	addr_QMdiArea_find uintptr
	addr_QMdiArea_createWindowContainer uintptr
	addr_QMdiArea_updateMicroFocus uintptr
	addr_QMdiArea_create uintptr
	addr_QMdiArea_destroy uintptr
	addr_QMdiArea_focusNextChild uintptr
	addr_QMdiArea_focusPreviousChild uintptr
	addr_QMdiArea_connect uintptr
	addr_QMdiArea_connect2 uintptr
	addr_QMdiArea_sender uintptr
	addr_QMdiArea_senderSignalIndex uintptr
	addr_QMdiArea_receivers uintptr
	addr_QMdiArea_isSignalConnected uintptr
	pg_QMdiArea_devicePixelRatioFScale func() float64
	pg_QMdiArea_encodeMetricF func(int32, float64) int32
	pg_QMdiArea_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMdiArea_tr2 uintptr
	addr_QMdiArea_tr3 uintptr
	addr_QMdiArea_subWindowList2 uintptr
	addr_QMdiArea_addSubWindow2 uintptr
	addr_QMdiArea_setOption2 uintptr
	addr_QMdiArea_createWindowContainer2 uintptr
	addr_QMdiArea_createWindowContainer3 uintptr
	addr_QMdiArea_updateMicroFocus2 uintptr
	addr_QMdiArea_create2 uintptr
	addr_QMdiArea_create3 uintptr
	addr_QMdiArea_create4 uintptr
	addr_QMdiArea_destroy2 uintptr
	addr_QMdiArea_destroy3 uintptr
	addr_QMdiArea_connect3 uintptr
	addr_QMdiArea_connect4 uintptr
	addr_QMdiArea_metaObject_default uintptr
	addr_QMdiArea_sizeHint_default uintptr
	addr_QMdiArea_minimumSizeHint_default uintptr
	addr_QMdiArea_setupViewport_default uintptr
	addr_QMdiArea_event_default uintptr
	addr_QMdiArea_eventFilter_default uintptr
	addr_QMdiArea_paintEvent_default uintptr
	addr_QMdiArea_childEvent_default uintptr
	addr_QMdiArea_resizeEvent_default uintptr
	addr_QMdiArea_timerEvent_default uintptr
	addr_QMdiArea_showEvent_default uintptr
	addr_QMdiArea_viewportEvent_default uintptr
	addr_QMdiArea_scrollContentsBy_default uintptr
	addr_QMdiArea_mousePressEvent_default uintptr
	addr_QMdiArea_mouseReleaseEvent_default uintptr
	addr_QMdiArea_mouseDoubleClickEvent_default uintptr
	addr_QMdiArea_mouseMoveEvent_default uintptr
	addr_QMdiArea_wheelEvent_default uintptr
	addr_QMdiArea_contextMenuEvent_default uintptr
	addr_QMdiArea_dragEnterEvent_default uintptr
	addr_QMdiArea_dragMoveEvent_default uintptr
	addr_QMdiArea_dragLeaveEvent_default uintptr
	addr_QMdiArea_dropEvent_default uintptr
	addr_QMdiArea_keyPressEvent_default uintptr
	addr_QMdiArea_viewportSizeHint_default uintptr
	addr_QMdiArea_changeEvent_default uintptr
	addr_QMdiArea_initStyleOption_default uintptr
	addr_QMdiArea_devType_default uintptr
	addr_QMdiArea_setVisible_default uintptr
	addr_QMdiArea_heightForWidth_default uintptr
	addr_QMdiArea_hasHeightForWidth_default uintptr
	addr_QMdiArea_paintEngine_default uintptr
	addr_QMdiArea_keyReleaseEvent_default uintptr
	addr_QMdiArea_focusInEvent_default uintptr
	addr_QMdiArea_focusOutEvent_default uintptr
	addr_QMdiArea_enterEvent_default uintptr
	addr_QMdiArea_leaveEvent_default uintptr
	addr_QMdiArea_moveEvent_default uintptr
	addr_QMdiArea_closeEvent_default uintptr
	addr_QMdiArea_tabletEvent_default uintptr
	addr_QMdiArea_actionEvent_default uintptr
	addr_QMdiArea_hideEvent_default uintptr
	addr_QMdiArea_metric_default uintptr
	addr_QMdiArea_initPainter_default uintptr
	addr_QMdiArea_redirected_default uintptr
	addr_QMdiArea_sharedPainter_default uintptr
	addr_QMdiArea_inputMethodEvent_default uintptr
	addr_QMdiArea_inputMethodQuery_default uintptr
	addr_QMdiArea_focusNextPrevChild_default uintptr
	addr_QMdiArea_customEvent_default uintptr
	addr_QMdiArea_connectNotify_default uintptr
	addr_QMdiArea_disconnectNotify_default uintptr
	pg_QMdiArea_connect_subWindowActivated func(uintptr, int64)
	pg_QMdiArea_connect_windowTitleChanged func(uintptr, int64)
	pg_QMdiArea_connect_windowIconChanged func(uintptr, int64)
	pg_QMdiArea_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMdiArea_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMdiArea_connect_destroyed func(uintptr, int64)
	pg_QMdiArea_connect_destroyed2 func(uintptr, int64)
	pg_QMdiArea_connect_objectNameChanged func(uintptr, int64)
	pg_QMdiArea__disconnect func(int64)
	pg_QMdiArea__set_override func(uintptr, int32)
	pg_QMdiArea_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMdiArea_chain, "QMdiArea_chain")
	symaddr(&addr_QMdiArea_new, "QMdiArea_new")
	symaddr(&addr_QMdiArea_new2, "QMdiArea_new2")
	symaddr(&addr_QMdiArea_tr, "QMdiArea_tr")
	symaddr(&addr_QMdiArea_currentSubWindow, "QMdiArea_currentSubWindow")
	symaddr(&addr_QMdiArea_activeSubWindow, "QMdiArea_activeSubWindow")
	symaddr(&addr_QMdiArea_subWindowList, "QMdiArea_subWindowList")
	symaddr(&addr_QMdiArea_addSubWindow, "QMdiArea_addSubWindow")
	symaddr(&addr_QMdiArea_removeSubWindow, "QMdiArea_removeSubWindow")
	symaddr(&addr_QMdiArea_background, "QMdiArea_background")
	symaddr(&addr_QMdiArea_setBackground, "QMdiArea_setBackground")
	symaddr(&addr_QMdiArea_activationOrder, "QMdiArea_activationOrder")
	symaddr(&addr_QMdiArea_setActivationOrder, "QMdiArea_setActivationOrder")
	symaddr(&addr_QMdiArea_setOption, "QMdiArea_setOption")
	symaddr(&addr_QMdiArea_testOption, "QMdiArea_testOption")
	symaddr(&addr_QMdiArea_setViewMode, "QMdiArea_setViewMode")
	symaddr(&addr_QMdiArea_viewMode, "QMdiArea_viewMode")
	symaddr(&addr_QMdiArea_documentMode, "QMdiArea_documentMode")
	symaddr(&addr_QMdiArea_setDocumentMode, "QMdiArea_setDocumentMode")
	symaddr(&addr_QMdiArea_setTabsClosable, "QMdiArea_setTabsClosable")
	symaddr(&addr_QMdiArea_tabsClosable, "QMdiArea_tabsClosable")
	symaddr(&addr_QMdiArea_setTabsMovable, "QMdiArea_setTabsMovable")
	symaddr(&addr_QMdiArea_tabsMovable, "QMdiArea_tabsMovable")
	symaddr(&addr_QMdiArea_setTabShape, "QMdiArea_setTabShape")
	symaddr(&addr_QMdiArea_tabShape, "QMdiArea_tabShape")
	symaddr(&addr_QMdiArea_setTabPosition, "QMdiArea_setTabPosition")
	symaddr(&addr_QMdiArea_tabPosition, "QMdiArea_tabPosition")
	symaddr(&addr_QMdiArea_subWindowActivated, "QMdiArea_subWindowActivated")
	symaddr(&addr_QMdiArea_setActiveSubWindow, "QMdiArea_setActiveSubWindow")
	symaddr(&addr_QMdiArea_tileSubWindows, "QMdiArea_tileSubWindows")
	symaddr(&addr_QMdiArea_cascadeSubWindows, "QMdiArea_cascadeSubWindows")
	symaddr(&addr_QMdiArea_closeActiveSubWindow, "QMdiArea_closeActiveSubWindow")
	symaddr(&addr_QMdiArea_closeAllSubWindows, "QMdiArea_closeAllSubWindows")
	symaddr(&addr_QMdiArea_activateNextSubWindow, "QMdiArea_activateNextSubWindow")
	symaddr(&addr_QMdiArea_activatePreviousSubWindow, "QMdiArea_activatePreviousSubWindow")
	symaddr(&addr_QMdiArea_setViewportMargins, "QMdiArea_setViewportMargins")
	symaddr(&addr_QMdiArea_setViewportMargins2, "QMdiArea_setViewportMargins2")
	symaddr(&addr_QMdiArea_viewportMargins, "QMdiArea_viewportMargins")
	symaddr(&addr_QMdiArea_drawFrame, "QMdiArea_drawFrame")
	symaddr(&addr_QMdiArea_setTabOrder, "QMdiArea_setTabOrder")
	symaddr(&addr_QMdiArea_mouseGrabber, "QMdiArea_mouseGrabber")
	symaddr(&addr_QMdiArea_keyboardGrabber, "QMdiArea_keyboardGrabber")
	symaddr(&addr_QMdiArea_find, "QMdiArea_find")
	symaddr(&addr_QMdiArea_createWindowContainer, "QMdiArea_createWindowContainer")
	symaddr(&addr_QMdiArea_updateMicroFocus, "QMdiArea_updateMicroFocus")
	symaddr(&addr_QMdiArea_create, "QMdiArea_create")
	symaddr(&addr_QMdiArea_destroy, "QMdiArea_destroy")
	symaddr(&addr_QMdiArea_focusNextChild, "QMdiArea_focusNextChild")
	symaddr(&addr_QMdiArea_focusPreviousChild, "QMdiArea_focusPreviousChild")
	symaddr(&addr_QMdiArea_connect, "QMdiArea_connect")
	symaddr(&addr_QMdiArea_connect2, "QMdiArea_connect2")
	symaddr(&addr_QMdiArea_sender, "QMdiArea_sender")
	symaddr(&addr_QMdiArea_senderSignalIndex, "QMdiArea_senderSignalIndex")
	symaddr(&addr_QMdiArea_receivers, "QMdiArea_receivers")
	symaddr(&addr_QMdiArea_isSignalConnected, "QMdiArea_isSignalConnected")
	reg(&pg_QMdiArea_devicePixelRatioFScale, "QMdiArea_devicePixelRatioFScale")
	reg(&pg_QMdiArea_encodeMetricF, "QMdiArea_encodeMetricF")
	reg(&pg_QMdiArea_getDecodedMetricF, "QMdiArea_getDecodedMetricF")
	symaddr(&addr_QMdiArea_tr2, "QMdiArea_tr2")
	symaddr(&addr_QMdiArea_tr3, "QMdiArea_tr3")
	symaddr(&addr_QMdiArea_subWindowList2, "QMdiArea_subWindowList2")
	symaddr(&addr_QMdiArea_addSubWindow2, "QMdiArea_addSubWindow2")
	symaddr(&addr_QMdiArea_setOption2, "QMdiArea_setOption2")
	symaddr(&addr_QMdiArea_createWindowContainer2, "QMdiArea_createWindowContainer2")
	symaddr(&addr_QMdiArea_createWindowContainer3, "QMdiArea_createWindowContainer3")
	symaddr(&addr_QMdiArea_updateMicroFocus2, "QMdiArea_updateMicroFocus2")
	symaddr(&addr_QMdiArea_create2, "QMdiArea_create2")
	symaddr(&addr_QMdiArea_create3, "QMdiArea_create3")
	symaddr(&addr_QMdiArea_create4, "QMdiArea_create4")
	symaddr(&addr_QMdiArea_destroy2, "QMdiArea_destroy2")
	symaddr(&addr_QMdiArea_destroy3, "QMdiArea_destroy3")
	symaddr(&addr_QMdiArea_connect3, "QMdiArea_connect3")
	symaddr(&addr_QMdiArea_connect4, "QMdiArea_connect4")
	symaddr(&addr_QMdiArea_metaObject_default, "QMdiArea_metaObject_default")
	symaddr(&addr_QMdiArea_sizeHint_default, "QMdiArea_sizeHint_default")
	symaddr(&addr_QMdiArea_minimumSizeHint_default, "QMdiArea_minimumSizeHint_default")
	symaddr(&addr_QMdiArea_setupViewport_default, "QMdiArea_setupViewport_default")
	symaddr(&addr_QMdiArea_event_default, "QMdiArea_event_default")
	symaddr(&addr_QMdiArea_eventFilter_default, "QMdiArea_eventFilter_default")
	symaddr(&addr_QMdiArea_paintEvent_default, "QMdiArea_paintEvent_default")
	symaddr(&addr_QMdiArea_childEvent_default, "QMdiArea_childEvent_default")
	symaddr(&addr_QMdiArea_resizeEvent_default, "QMdiArea_resizeEvent_default")
	symaddr(&addr_QMdiArea_timerEvent_default, "QMdiArea_timerEvent_default")
	symaddr(&addr_QMdiArea_showEvent_default, "QMdiArea_showEvent_default")
	symaddr(&addr_QMdiArea_viewportEvent_default, "QMdiArea_viewportEvent_default")
	symaddr(&addr_QMdiArea_scrollContentsBy_default, "QMdiArea_scrollContentsBy_default")
	symaddr(&addr_QMdiArea_mousePressEvent_default, "QMdiArea_mousePressEvent_default")
	symaddr(&addr_QMdiArea_mouseReleaseEvent_default, "QMdiArea_mouseReleaseEvent_default")
	symaddr(&addr_QMdiArea_mouseDoubleClickEvent_default, "QMdiArea_mouseDoubleClickEvent_default")
	symaddr(&addr_QMdiArea_mouseMoveEvent_default, "QMdiArea_mouseMoveEvent_default")
	symaddr(&addr_QMdiArea_wheelEvent_default, "QMdiArea_wheelEvent_default")
	symaddr(&addr_QMdiArea_contextMenuEvent_default, "QMdiArea_contextMenuEvent_default")
	symaddr(&addr_QMdiArea_dragEnterEvent_default, "QMdiArea_dragEnterEvent_default")
	symaddr(&addr_QMdiArea_dragMoveEvent_default, "QMdiArea_dragMoveEvent_default")
	symaddr(&addr_QMdiArea_dragLeaveEvent_default, "QMdiArea_dragLeaveEvent_default")
	symaddr(&addr_QMdiArea_dropEvent_default, "QMdiArea_dropEvent_default")
	symaddr(&addr_QMdiArea_keyPressEvent_default, "QMdiArea_keyPressEvent_default")
	symaddr(&addr_QMdiArea_viewportSizeHint_default, "QMdiArea_viewportSizeHint_default")
	symaddr(&addr_QMdiArea_changeEvent_default, "QMdiArea_changeEvent_default")
	symaddr(&addr_QMdiArea_initStyleOption_default, "QMdiArea_initStyleOption_default")
	symaddr(&addr_QMdiArea_devType_default, "QMdiArea_devType_default")
	symaddr(&addr_QMdiArea_setVisible_default, "QMdiArea_setVisible_default")
	symaddr(&addr_QMdiArea_heightForWidth_default, "QMdiArea_heightForWidth_default")
	symaddr(&addr_QMdiArea_hasHeightForWidth_default, "QMdiArea_hasHeightForWidth_default")
	symaddr(&addr_QMdiArea_paintEngine_default, "QMdiArea_paintEngine_default")
	symaddr(&addr_QMdiArea_keyReleaseEvent_default, "QMdiArea_keyReleaseEvent_default")
	symaddr(&addr_QMdiArea_focusInEvent_default, "QMdiArea_focusInEvent_default")
	symaddr(&addr_QMdiArea_focusOutEvent_default, "QMdiArea_focusOutEvent_default")
	symaddr(&addr_QMdiArea_enterEvent_default, "QMdiArea_enterEvent_default")
	symaddr(&addr_QMdiArea_leaveEvent_default, "QMdiArea_leaveEvent_default")
	symaddr(&addr_QMdiArea_moveEvent_default, "QMdiArea_moveEvent_default")
	symaddr(&addr_QMdiArea_closeEvent_default, "QMdiArea_closeEvent_default")
	symaddr(&addr_QMdiArea_tabletEvent_default, "QMdiArea_tabletEvent_default")
	symaddr(&addr_QMdiArea_actionEvent_default, "QMdiArea_actionEvent_default")
	symaddr(&addr_QMdiArea_hideEvent_default, "QMdiArea_hideEvent_default")
	symaddr(&addr_QMdiArea_metric_default, "QMdiArea_metric_default")
	symaddr(&addr_QMdiArea_initPainter_default, "QMdiArea_initPainter_default")
	symaddr(&addr_QMdiArea_redirected_default, "QMdiArea_redirected_default")
	symaddr(&addr_QMdiArea_sharedPainter_default, "QMdiArea_sharedPainter_default")
	symaddr(&addr_QMdiArea_inputMethodEvent_default, "QMdiArea_inputMethodEvent_default")
	symaddr(&addr_QMdiArea_inputMethodQuery_default, "QMdiArea_inputMethodQuery_default")
	symaddr(&addr_QMdiArea_focusNextPrevChild_default, "QMdiArea_focusNextPrevChild_default")
	symaddr(&addr_QMdiArea_customEvent_default, "QMdiArea_customEvent_default")
	symaddr(&addr_QMdiArea_connectNotify_default, "QMdiArea_connectNotify_default")
	symaddr(&addr_QMdiArea_disconnectNotify_default, "QMdiArea_disconnectNotify_default")
	reg(&pg_QMdiArea_connect_subWindowActivated, "QMdiArea_connect_subWindowActivated")
	reg(&pg_QMdiArea_connect_windowTitleChanged, "QMdiArea_connect_windowTitleChanged")
	reg(&pg_QMdiArea_connect_windowIconChanged, "QMdiArea_connect_windowIconChanged")
	reg(&pg_QMdiArea_connect_windowIconTextChanged, "QMdiArea_connect_windowIconTextChanged")
	reg(&pg_QMdiArea_connect_customContextMenuRequested, "QMdiArea_connect_customContextMenuRequested")
	reg(&pg_QMdiArea_connect_destroyed, "QMdiArea_connect_destroyed")
	reg(&pg_QMdiArea_connect_destroyed2, "QMdiArea_connect_destroyed2")
	reg(&pg_QMdiArea_connect_objectNameChanged, "QMdiArea_connect_objectNameChanged")
	reg(&pg_QMdiArea__disconnect, "QMdiArea__disconnect")
	reg(&pg_QMdiArea__set_override, "QMdiArea__set_override")
	reg(&pg_QMdiArea_delete, "QMdiArea_delete")
}

type QMdiArea struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQMdiArea(h uintptr) *QMdiArea {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QMdiArea_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMdiArea{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQMdiArea(h uintptr) *QMdiArea { return bindQMdiArea(h) }

func (this *QMdiArea) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMdiArea) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMdiArea(h unsafe.Pointer) *QMdiArea { return bindQMdiArea(uintptr(h)) }
func UnsafeNewQMdiAreaOwned(h unsafe.Pointer) *QMdiArea { return newownedQMdiArea(uintptr(h)) }

func (this *QMdiArea) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func NewQMdiArea(parent *QWidget) *QMdiArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMdiArea_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMdiArea(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMdiArea_virtIdx, QMdiArea_enable)
	return __o
}

func NewQMdiArea2() *QMdiArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMdiArea_new2, uintptr(__goid))
	__o := newownedQMdiArea(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMdiArea_virtIdx, QMdiArea_enable)
	return __o
}

func QMdiArea_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiArea_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMdiArea) CurrentSubWindow() *QMdiSubWindow {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_currentSubWindow, this.h)
	runtime.KeepAlive(this)
	return bindQMdiSubWindow(__r)
}

func (this *QMdiArea) ActiveSubWindow() *QMdiSubWindow {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_activeSubWindow, this.h)
	runtime.KeepAlive(this)
	return bindQMdiSubWindow(__r)
}

func (this *QMdiArea) SubWindowList() []*QMdiSubWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiArea_subWindowList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QMdiSubWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQMdiSubWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMdiArea) AddSubWindow(widget *QWidget) *QMdiSubWindow {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_addSubWindow, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQMdiSubWindow(__r)
}

func (this *QMdiArea) RemoveSubWindow(widget *QWidget) {
	purego.SyscallN(addr_QMdiArea_removeSubWindow, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QMdiArea) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_background, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QMdiArea) SetBackground(background *QBrush) {
	purego.SyscallN(addr_QMdiArea_setBackground, this.h, uintptr(background.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(background)
}

func (this *QMdiArea) ActivationOrder() QMdiArea__WindowOrder {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_activationOrder, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMdiArea__WindowOrder(__r)
}

func (this *QMdiArea) SetActivationOrder(order QMdiArea__WindowOrder) {
	purego.SyscallN(addr_QMdiArea_setActivationOrder, this.h, uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) SetOption(option QMdiArea__AreaOption) {
	purego.SyscallN(addr_QMdiArea_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) TestOption(opton QMdiArea__AreaOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_testOption, this.h, uintptr(opton))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) SetViewMode(mode QMdiArea__ViewMode) {
	purego.SyscallN(addr_QMdiArea_setViewMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) ViewMode() QMdiArea__ViewMode {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_viewMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMdiArea__ViewMode(__r)
}

func (this *QMdiArea) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_documentMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) SetDocumentMode(enabled bool) {
	purego.SyscallN(addr_QMdiArea_setDocumentMode, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) SetTabsClosable(closable bool) {
	purego.SyscallN(addr_QMdiArea_setTabsClosable, this.h, uintptr(b2i(closable)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) TabsClosable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_tabsClosable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) SetTabsMovable(movable bool) {
	purego.SyscallN(addr_QMdiArea_setTabsMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) TabsMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_tabsMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) SetTabShape(shape QTabWidget__TabShape) {
	purego.SyscallN(addr_QMdiArea_setTabShape, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) TabShape() QTabWidget__TabShape {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_tabShape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabShape(__r)
}

func (this *QMdiArea) SetTabPosition(position QTabWidget__TabPosition) {
	purego.SyscallN(addr_QMdiArea_setTabPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) TabPosition() QTabWidget__TabPosition {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_tabPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabPosition(__r)
}

func (this *QMdiArea) SubWindowActivated(arg1 *QMdiSubWindow) {
	purego.SyscallN(addr_QMdiArea_subWindowActivated, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) SetActiveSubWindow(window *QMdiSubWindow) {
	purego.SyscallN(addr_QMdiArea_setActiveSubWindow, this.h, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
}

func (this *QMdiArea) TileSubWindows() {
	purego.SyscallN(addr_QMdiArea_tileSubWindows, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) CascadeSubWindows() {
	purego.SyscallN(addr_QMdiArea_cascadeSubWindows, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) CloseActiveSubWindow() {
	purego.SyscallN(addr_QMdiArea_closeActiveSubWindow, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) CloseAllSubWindows() {
	purego.SyscallN(addr_QMdiArea_closeAllSubWindows, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) ActivateNextSubWindow() {
	purego.SyscallN(addr_QMdiArea_activateNextSubWindow, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) ActivatePreviousSubWindow() {
	purego.SyscallN(addr_QMdiArea_activatePreviousSubWindow, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QMdiArea_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QMdiArea_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QMdiArea) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMdiArea) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QMdiArea_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QMdiArea_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMdiArea_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMdiArea_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_mouseGrabber)
	return bindQWidget(__r)
}

func QMdiArea_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_keyboardGrabber)
	return bindQWidget(__r)
}

func QMdiArea_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMdiArea_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMdiArea) UpdateMicroFocus() {
	purego.SyscallN(addr_QMdiArea_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) Create() {
	purego.SyscallN(addr_QMdiArea_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) Destroy() {
	purego.SyscallN(addr_QMdiArea_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiArea) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMdiArea_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMdiArea_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMdiArea_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMdiArea) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMdiArea) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiArea) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMdiArea) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMdiArea_DevicePixelRatioFScale() float64 {
	__r := pg_QMdiArea_devicePixelRatioFScale()
	return __r
}

func QMdiArea_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMdiArea_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMdiArea) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMdiArea_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMdiArea_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiArea_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMdiArea_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiArea_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMdiArea) SubWindowListWithOrder(order QMdiArea__WindowOrder) []*QMdiSubWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiArea_subWindowList2, this.h, uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QMdiSubWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQMdiSubWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMdiArea) AddSubWindow2(widget *QWidget, flags WindowType) *QMdiSubWindow {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_addSubWindow2, this.h, uintptr(widget.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQMdiSubWindow(__r)
}

func (this *QMdiArea) SetOption2(option QMdiArea__AreaOption, on bool) {
	purego.SyscallN(addr_QMdiArea_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QMdiArea_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMdiArea_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMdiArea) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMdiArea_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMdiArea_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMdiArea_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMdiArea_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMdiArea_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMdiArea_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMdiArea_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMdiArea_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMdiArea_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMdiArea) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMdiArea) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMdiArea) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMdiArea) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QMdiArea_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QMdiArea) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMdiArea) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMdiArea) PaintEventDefault(paintEvent *QPaintEvent) {
	purego.SyscallN(addr_QMdiArea_paintEvent_default, this.h, uintptr(paintEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(paintEvent)
}

func (this *QMdiArea) ChildEventDefault(childEvent *QChildEvent) {
	purego.SyscallN(addr_QMdiArea_childEvent_default, this.h, uintptr(childEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childEvent)
}

func (this *QMdiArea) ResizeEventDefault(resizeEvent *QResizeEvent) {
	purego.SyscallN(addr_QMdiArea_resizeEvent_default, this.h, uintptr(resizeEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(resizeEvent)
}

func (this *QMdiArea) TimerEventDefault(timerEvent *QTimerEvent) {
	purego.SyscallN(addr_QMdiArea_timerEvent_default, this.h, uintptr(timerEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timerEvent)
}

func (this *QMdiArea) ShowEventDefault(showEvent *QShowEvent) {
	purego.SyscallN(addr_QMdiArea_showEvent_default, this.h, uintptr(showEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(showEvent)
}

func (this *QMdiArea) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMdiArea) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QMdiArea_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMdiArea_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMdiArea_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMdiArea_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMdiArea_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QMdiArea_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QMdiArea_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) DragEnterEventDefault(arg1 *QDragEnterEvent) {
	purego.SyscallN(addr_QMdiArea_dragEnterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) DragMoveEventDefault(arg1 *QDragMoveEvent) {
	purego.SyscallN(addr_QMdiArea_dragMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) DragLeaveEventDefault(arg1 *QDragLeaveEvent) {
	purego.SyscallN(addr_QMdiArea_dragLeaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) DropEventDefault(arg1 *QDropEvent) {
	purego.SyscallN(addr_QMdiArea_dropEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QMdiArea_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMdiArea) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMdiArea_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QMdiArea_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QMdiArea) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiArea) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMdiArea_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMdiArea) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiArea) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMdiArea) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMdiArea_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMdiArea_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMdiArea_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QMdiArea_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMdiArea_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QMdiArea_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QMdiArea_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMdiArea_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QMdiArea_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QMdiArea_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiArea) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMdiArea_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMdiArea) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMdiArea) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMdiArea) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMdiArea_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiArea) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMdiArea_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMdiArea) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiArea_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiArea) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMdiArea_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiArea) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMdiArea_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMdiArea) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMdiArea_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMdiArea) OnSubWindowActivated(fn func(*QMdiSubWindow)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQMdiSubWindow(unsafe.Pointer(arg1))) })
	pg_QMdiArea_connect_subWindowActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMdiArea_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMdiArea_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMdiArea_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMdiArea_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMdiArea_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMdiArea_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMdiArea_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiArea) Disconnect(handle int64) {
	pg_QMdiArea__disconnect(handle)
	unregisterCallback(handle)
}

func QMdiArea_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 4531996381736730688:
		return 3
	case 1237236456312824991:
		return 4
	case 5217623514630858297:
		return 5
	case 5290442163167458465:
		return 6
	case -1185315367526857833:
		return 7
	case -4302917285655014661:
		return 8
	case -5409927581812374498:
		return 9
	case -6825369174751916:
		return 10
	case -1971465653061861519:
		return 11
	case -6118001150126564745:
		return 12
	case -332090898078543319:
		return 13
	case -4415920049297795213:
		return 14
	case -4029073905697023887:
		return 15
	case -4255494004099806581:
		return 16
	case 1919988340934511316:
		return 17
	case -6698043222227006015:
		return 18
	case 206122715919917589:
		return 19
	case -1781545137352745978:
		return 20
	case -3139904158373725006:
		return 21
	case 9107058682754591716:
		return 22
	case 1023791977679793185:
		return 23
	case 503579067252048239:
		return 24
	case 4959376063826165943:
		return 25
	case -488693769771308531:
		return 26
	case -8659679893868524044:
		return 27
	case 2110182719375245575:
		return 28
	case 4311534060495737339:
		return 29
	case -8266762397995423915:
		return 30
	case 6398007404756346381:
		return 31
	case 772598475950050299:
		return 32
	case 975043767033552552:
		return 33
	case -385596010576048993:
		return 34
	case 6316270346668237895:
		return 35
	case -551070009623711640:
		return 36
	case -9161949150489525396:
		return 37
	case -1989937050019464783:
		return 38
	case -212068866860701141:
		return 39
	case 1697444140088555581:
		return 40
	case -199718707876719129:
		return 41
	case -2039166586246443706:
		return 42
	case 427629649267311777:
		return 43
	case 6477500296980050210:
		return 44
	case 4127240633310655048:
		return 45
	case 5792771328741774991:
		return 46
	case 5438819661942843222:
		return 47
	case 7345853097340709638:
		return 48
	case 5898328517505784191:
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

func QMdiArea_enable(h uintptr, idx int32) { pg_QMdiArea__set_override(h, idx) }

func (this *QMdiArea) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMdiArea) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMdiArea) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMdiArea) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QMdiArea) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QMdiArea) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMdiArea) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(paintEvent uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(paintEvent))) }) // slot "paintEvent"
}

func (this *QMdiArea) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(childEvent uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(childEvent))) }) // slot "childEvent"
}

func (this *QMdiArea) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(resizeEvent uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(resizeEvent))) }) // slot "resizeEvent"
}

func (this *QMdiArea) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(timerEvent uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(timerEvent))) }) // slot "timerEvent"
}

func (this *QMdiArea) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(showEvent uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(showEvent))) }) // slot "showEvent"
}

func (this *QMdiArea) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QMdiArea) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QMdiArea) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QMdiArea) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QMdiArea) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMdiArea) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QMdiArea) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QMdiArea) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QMdiArea) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(arg1 uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(arg1))) }) // slot "dragEnterEvent"
}

func (this *QMdiArea) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(arg1 uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(arg1))) }) // slot "dragMoveEvent"
}

func (this *QMdiArea) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(arg1 uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(arg1))) }) // slot "dragLeaveEvent"
}

func (this *QMdiArea) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(arg1 uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(arg1))) }) // slot "dropEvent"
}

func (this *QMdiArea) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QMdiArea) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QMdiArea) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QMdiArea) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QMdiArea) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMdiArea) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMdiArea) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMdiArea) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMdiArea) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMdiArea) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMdiArea) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QMdiArea) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QMdiArea) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QMdiArea) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QMdiArea) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QMdiArea) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QMdiArea) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMdiArea) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QMdiArea) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QMdiArea) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMdiArea) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMdiArea) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMdiArea) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMdiArea) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMdiArea) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMdiArea) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMdiArea) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMdiArea) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMdiArea) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMdiArea) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMdiArea) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMdiArea_delete(this.h)
	}
	this.UnsafeRelease()
}
