package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMainWindow_chain uintptr
	addr_QMainWindow_new uintptr
	addr_QMainWindow_new2 uintptr
	addr_QMainWindow_new3 uintptr
	addr_QMainWindow_tr uintptr
	addr_QMainWindow_iconSize uintptr
	addr_QMainWindow_setIconSize uintptr
	addr_QMainWindow_toolButtonStyle uintptr
	addr_QMainWindow_setToolButtonStyle uintptr
	addr_QMainWindow_isAnimated uintptr
	addr_QMainWindow_isDockNestingEnabled uintptr
	addr_QMainWindow_documentMode uintptr
	addr_QMainWindow_setDocumentMode uintptr
	addr_QMainWindow_tabShape uintptr
	addr_QMainWindow_setTabShape uintptr
	addr_QMainWindow_tabPosition uintptr
	addr_QMainWindow_setTabPosition uintptr
	addr_QMainWindow_setDockOptions uintptr
	addr_QMainWindow_dockOptions uintptr
	addr_QMainWindow_isSeparator uintptr
	addr_QMainWindow_menuBar uintptr
	addr_QMainWindow_setMenuBar uintptr
	addr_QMainWindow_menuWidget uintptr
	addr_QMainWindow_setMenuWidget uintptr
	addr_QMainWindow_statusBar uintptr
	addr_QMainWindow_setStatusBar uintptr
	addr_QMainWindow_centralWidget uintptr
	addr_QMainWindow_setCentralWidget uintptr
	addr_QMainWindow_takeCentralWidget uintptr
	addr_QMainWindow_setCorner uintptr
	addr_QMainWindow_corner uintptr
	addr_QMainWindow_addToolBarBreak uintptr
	addr_QMainWindow_insertToolBarBreak uintptr
	addr_QMainWindow_addToolBar uintptr
	addr_QMainWindow_addToolBar2 uintptr
	addr_QMainWindow_addToolBar3 uintptr
	addr_QMainWindow_insertToolBar uintptr
	addr_QMainWindow_removeToolBar uintptr
	addr_QMainWindow_removeToolBarBreak uintptr
	addr_QMainWindow_unifiedTitleAndToolBarOnMac uintptr
	addr_QMainWindow_toolBarArea uintptr
	addr_QMainWindow_toolBarBreak uintptr
	addr_QMainWindow_addDockWidget uintptr
	addr_QMainWindow_addDockWidget2 uintptr
	addr_QMainWindow_splitDockWidget uintptr
	addr_QMainWindow_tabifyDockWidget uintptr
	addr_QMainWindow_tabifiedDockWidgets uintptr
	addr_QMainWindow_removeDockWidget uintptr
	addr_QMainWindow_restoreDockWidget uintptr
	addr_QMainWindow_dockWidgetArea uintptr
	addr_QMainWindow_resizeDocks uintptr
	addr_QMainWindow_saveState uintptr
	addr_QMainWindow_restoreState uintptr
	addr_QMainWindow_createPopupMenu uintptr
	addr_QMainWindow_setAnimated uintptr
	addr_QMainWindow_setDockNestingEnabled uintptr
	addr_QMainWindow_setUnifiedTitleAndToolBarOnMac uintptr
	addr_QMainWindow_iconSizeChanged uintptr
	addr_QMainWindow_toolButtonStyleChanged uintptr
	addr_QMainWindow_tabifiedDockWidgetActivated uintptr
	addr_QMainWindow_setTabOrder uintptr
	addr_QMainWindow_mouseGrabber uintptr
	addr_QMainWindow_keyboardGrabber uintptr
	addr_QMainWindow_find uintptr
	addr_QMainWindow_createWindowContainer uintptr
	addr_QMainWindow_updateMicroFocus uintptr
	addr_QMainWindow_create uintptr
	addr_QMainWindow_destroy uintptr
	addr_QMainWindow_focusNextChild uintptr
	addr_QMainWindow_focusPreviousChild uintptr
	addr_QMainWindow_connect uintptr
	addr_QMainWindow_connect2 uintptr
	addr_QMainWindow_sender uintptr
	addr_QMainWindow_senderSignalIndex uintptr
	addr_QMainWindow_receivers uintptr
	addr_QMainWindow_isSignalConnected uintptr
	pg_QMainWindow_devicePixelRatioFScale func() float64
	pg_QMainWindow_encodeMetricF func(int32, float64) int32
	pg_QMainWindow_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMainWindow_tr2 uintptr
	addr_QMainWindow_tr3 uintptr
	addr_QMainWindow_addToolBarBreak2 uintptr
	addr_QMainWindow_saveState2 uintptr
	addr_QMainWindow_restoreState2 uintptr
	addr_QMainWindow_createWindowContainer2 uintptr
	addr_QMainWindow_createWindowContainer3 uintptr
	addr_QMainWindow_updateMicroFocus2 uintptr
	addr_QMainWindow_create2 uintptr
	addr_QMainWindow_create3 uintptr
	addr_QMainWindow_create4 uintptr
	addr_QMainWindow_destroy2 uintptr
	addr_QMainWindow_destroy3 uintptr
	addr_QMainWindow_connect3 uintptr
	addr_QMainWindow_connect4 uintptr
	addr_QMainWindow_createPopupMenu_default uintptr
	addr_QMainWindow_metaObject_default uintptr
	addr_QMainWindow_contextMenuEvent_default uintptr
	addr_QMainWindow_event_default uintptr
	addr_QMainWindow_devType_default uintptr
	addr_QMainWindow_setVisible_default uintptr
	addr_QMainWindow_sizeHint_default uintptr
	addr_QMainWindow_minimumSizeHint_default uintptr
	addr_QMainWindow_heightForWidth_default uintptr
	addr_QMainWindow_hasHeightForWidth_default uintptr
	addr_QMainWindow_paintEngine_default uintptr
	addr_QMainWindow_mousePressEvent_default uintptr
	addr_QMainWindow_mouseReleaseEvent_default uintptr
	addr_QMainWindow_mouseDoubleClickEvent_default uintptr
	addr_QMainWindow_mouseMoveEvent_default uintptr
	addr_QMainWindow_wheelEvent_default uintptr
	addr_QMainWindow_keyPressEvent_default uintptr
	addr_QMainWindow_keyReleaseEvent_default uintptr
	addr_QMainWindow_focusInEvent_default uintptr
	addr_QMainWindow_focusOutEvent_default uintptr
	addr_QMainWindow_enterEvent_default uintptr
	addr_QMainWindow_leaveEvent_default uintptr
	addr_QMainWindow_paintEvent_default uintptr
	addr_QMainWindow_moveEvent_default uintptr
	addr_QMainWindow_resizeEvent_default uintptr
	addr_QMainWindow_closeEvent_default uintptr
	addr_QMainWindow_tabletEvent_default uintptr
	addr_QMainWindow_actionEvent_default uintptr
	addr_QMainWindow_dragEnterEvent_default uintptr
	addr_QMainWindow_dragMoveEvent_default uintptr
	addr_QMainWindow_dragLeaveEvent_default uintptr
	addr_QMainWindow_dropEvent_default uintptr
	addr_QMainWindow_showEvent_default uintptr
	addr_QMainWindow_hideEvent_default uintptr
	addr_QMainWindow_changeEvent_default uintptr
	addr_QMainWindow_metric_default uintptr
	addr_QMainWindow_initPainter_default uintptr
	addr_QMainWindow_redirected_default uintptr
	addr_QMainWindow_sharedPainter_default uintptr
	addr_QMainWindow_inputMethodEvent_default uintptr
	addr_QMainWindow_inputMethodQuery_default uintptr
	addr_QMainWindow_focusNextPrevChild_default uintptr
	addr_QMainWindow_eventFilter_default uintptr
	addr_QMainWindow_timerEvent_default uintptr
	addr_QMainWindow_childEvent_default uintptr
	addr_QMainWindow_customEvent_default uintptr
	addr_QMainWindow_connectNotify_default uintptr
	addr_QMainWindow_disconnectNotify_default uintptr
	pg_QMainWindow_connect_iconSizeChanged func(uintptr, int64)
	pg_QMainWindow_connect_toolButtonStyleChanged func(uintptr, int64)
	pg_QMainWindow_connect_tabifiedDockWidgetActivated func(uintptr, int64)
	pg_QMainWindow_connect_windowTitleChanged func(uintptr, int64)
	pg_QMainWindow_connect_windowIconChanged func(uintptr, int64)
	pg_QMainWindow_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMainWindow_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMainWindow_connect_destroyed func(uintptr, int64)
	pg_QMainWindow_connect_destroyed2 func(uintptr, int64)
	pg_QMainWindow_connect_objectNameChanged func(uintptr, int64)
	pg_QMainWindow__disconnect func(int64)
	pg_QMainWindow__set_override func(uintptr, int32)
	pg_QMainWindow_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMainWindow_chain, "QMainWindow_chain")
	symaddr(&addr_QMainWindow_new, "QMainWindow_new")
	symaddr(&addr_QMainWindow_new2, "QMainWindow_new2")
	symaddr(&addr_QMainWindow_new3, "QMainWindow_new3")
	symaddr(&addr_QMainWindow_tr, "QMainWindow_tr")
	symaddr(&addr_QMainWindow_iconSize, "QMainWindow_iconSize")
	symaddr(&addr_QMainWindow_setIconSize, "QMainWindow_setIconSize")
	symaddr(&addr_QMainWindow_toolButtonStyle, "QMainWindow_toolButtonStyle")
	symaddr(&addr_QMainWindow_setToolButtonStyle, "QMainWindow_setToolButtonStyle")
	symaddr(&addr_QMainWindow_isAnimated, "QMainWindow_isAnimated")
	symaddr(&addr_QMainWindow_isDockNestingEnabled, "QMainWindow_isDockNestingEnabled")
	symaddr(&addr_QMainWindow_documentMode, "QMainWindow_documentMode")
	symaddr(&addr_QMainWindow_setDocumentMode, "QMainWindow_setDocumentMode")
	symaddr(&addr_QMainWindow_tabShape, "QMainWindow_tabShape")
	symaddr(&addr_QMainWindow_setTabShape, "QMainWindow_setTabShape")
	symaddr(&addr_QMainWindow_tabPosition, "QMainWindow_tabPosition")
	symaddr(&addr_QMainWindow_setTabPosition, "QMainWindow_setTabPosition")
	symaddr(&addr_QMainWindow_setDockOptions, "QMainWindow_setDockOptions")
	symaddr(&addr_QMainWindow_dockOptions, "QMainWindow_dockOptions")
	symaddr(&addr_QMainWindow_isSeparator, "QMainWindow_isSeparator")
	symaddr(&addr_QMainWindow_menuBar, "QMainWindow_menuBar")
	symaddr(&addr_QMainWindow_setMenuBar, "QMainWindow_setMenuBar")
	symaddr(&addr_QMainWindow_menuWidget, "QMainWindow_menuWidget")
	symaddr(&addr_QMainWindow_setMenuWidget, "QMainWindow_setMenuWidget")
	symaddr(&addr_QMainWindow_statusBar, "QMainWindow_statusBar")
	symaddr(&addr_QMainWindow_setStatusBar, "QMainWindow_setStatusBar")
	symaddr(&addr_QMainWindow_centralWidget, "QMainWindow_centralWidget")
	symaddr(&addr_QMainWindow_setCentralWidget, "QMainWindow_setCentralWidget")
	symaddr(&addr_QMainWindow_takeCentralWidget, "QMainWindow_takeCentralWidget")
	symaddr(&addr_QMainWindow_setCorner, "QMainWindow_setCorner")
	symaddr(&addr_QMainWindow_corner, "QMainWindow_corner")
	symaddr(&addr_QMainWindow_addToolBarBreak, "QMainWindow_addToolBarBreak")
	symaddr(&addr_QMainWindow_insertToolBarBreak, "QMainWindow_insertToolBarBreak")
	symaddr(&addr_QMainWindow_addToolBar, "QMainWindow_addToolBar")
	symaddr(&addr_QMainWindow_addToolBar2, "QMainWindow_addToolBar2")
	symaddr(&addr_QMainWindow_addToolBar3, "QMainWindow_addToolBar3")
	symaddr(&addr_QMainWindow_insertToolBar, "QMainWindow_insertToolBar")
	symaddr(&addr_QMainWindow_removeToolBar, "QMainWindow_removeToolBar")
	symaddr(&addr_QMainWindow_removeToolBarBreak, "QMainWindow_removeToolBarBreak")
	symaddr(&addr_QMainWindow_unifiedTitleAndToolBarOnMac, "QMainWindow_unifiedTitleAndToolBarOnMac")
	symaddr(&addr_QMainWindow_toolBarArea, "QMainWindow_toolBarArea")
	symaddr(&addr_QMainWindow_toolBarBreak, "QMainWindow_toolBarBreak")
	symaddr(&addr_QMainWindow_addDockWidget, "QMainWindow_addDockWidget")
	symaddr(&addr_QMainWindow_addDockWidget2, "QMainWindow_addDockWidget2")
	symaddr(&addr_QMainWindow_splitDockWidget, "QMainWindow_splitDockWidget")
	symaddr(&addr_QMainWindow_tabifyDockWidget, "QMainWindow_tabifyDockWidget")
	symaddr(&addr_QMainWindow_tabifiedDockWidgets, "QMainWindow_tabifiedDockWidgets")
	symaddr(&addr_QMainWindow_removeDockWidget, "QMainWindow_removeDockWidget")
	symaddr(&addr_QMainWindow_restoreDockWidget, "QMainWindow_restoreDockWidget")
	symaddr(&addr_QMainWindow_dockWidgetArea, "QMainWindow_dockWidgetArea")
	symaddr(&addr_QMainWindow_resizeDocks, "QMainWindow_resizeDocks")
	symaddr(&addr_QMainWindow_saveState, "QMainWindow_saveState")
	symaddr(&addr_QMainWindow_restoreState, "QMainWindow_restoreState")
	symaddr(&addr_QMainWindow_createPopupMenu, "QMainWindow_createPopupMenu")
	symaddr(&addr_QMainWindow_setAnimated, "QMainWindow_setAnimated")
	symaddr(&addr_QMainWindow_setDockNestingEnabled, "QMainWindow_setDockNestingEnabled")
	symaddr(&addr_QMainWindow_setUnifiedTitleAndToolBarOnMac, "QMainWindow_setUnifiedTitleAndToolBarOnMac")
	symaddr(&addr_QMainWindow_iconSizeChanged, "QMainWindow_iconSizeChanged")
	symaddr(&addr_QMainWindow_toolButtonStyleChanged, "QMainWindow_toolButtonStyleChanged")
	symaddr(&addr_QMainWindow_tabifiedDockWidgetActivated, "QMainWindow_tabifiedDockWidgetActivated")
	symaddr(&addr_QMainWindow_setTabOrder, "QMainWindow_setTabOrder")
	symaddr(&addr_QMainWindow_mouseGrabber, "QMainWindow_mouseGrabber")
	symaddr(&addr_QMainWindow_keyboardGrabber, "QMainWindow_keyboardGrabber")
	symaddr(&addr_QMainWindow_find, "QMainWindow_find")
	symaddr(&addr_QMainWindow_createWindowContainer, "QMainWindow_createWindowContainer")
	symaddr(&addr_QMainWindow_updateMicroFocus, "QMainWindow_updateMicroFocus")
	symaddr(&addr_QMainWindow_create, "QMainWindow_create")
	symaddr(&addr_QMainWindow_destroy, "QMainWindow_destroy")
	symaddr(&addr_QMainWindow_focusNextChild, "QMainWindow_focusNextChild")
	symaddr(&addr_QMainWindow_focusPreviousChild, "QMainWindow_focusPreviousChild")
	symaddr(&addr_QMainWindow_connect, "QMainWindow_connect")
	symaddr(&addr_QMainWindow_connect2, "QMainWindow_connect2")
	symaddr(&addr_QMainWindow_sender, "QMainWindow_sender")
	symaddr(&addr_QMainWindow_senderSignalIndex, "QMainWindow_senderSignalIndex")
	symaddr(&addr_QMainWindow_receivers, "QMainWindow_receivers")
	symaddr(&addr_QMainWindow_isSignalConnected, "QMainWindow_isSignalConnected")
	reg(&pg_QMainWindow_devicePixelRatioFScale, "QMainWindow_devicePixelRatioFScale")
	reg(&pg_QMainWindow_encodeMetricF, "QMainWindow_encodeMetricF")
	reg(&pg_QMainWindow_getDecodedMetricF, "QMainWindow_getDecodedMetricF")
	symaddr(&addr_QMainWindow_tr2, "QMainWindow_tr2")
	symaddr(&addr_QMainWindow_tr3, "QMainWindow_tr3")
	symaddr(&addr_QMainWindow_addToolBarBreak2, "QMainWindow_addToolBarBreak2")
	symaddr(&addr_QMainWindow_saveState2, "QMainWindow_saveState2")
	symaddr(&addr_QMainWindow_restoreState2, "QMainWindow_restoreState2")
	symaddr(&addr_QMainWindow_createWindowContainer2, "QMainWindow_createWindowContainer2")
	symaddr(&addr_QMainWindow_createWindowContainer3, "QMainWindow_createWindowContainer3")
	symaddr(&addr_QMainWindow_updateMicroFocus2, "QMainWindow_updateMicroFocus2")
	symaddr(&addr_QMainWindow_create2, "QMainWindow_create2")
	symaddr(&addr_QMainWindow_create3, "QMainWindow_create3")
	symaddr(&addr_QMainWindow_create4, "QMainWindow_create4")
	symaddr(&addr_QMainWindow_destroy2, "QMainWindow_destroy2")
	symaddr(&addr_QMainWindow_destroy3, "QMainWindow_destroy3")
	symaddr(&addr_QMainWindow_connect3, "QMainWindow_connect3")
	symaddr(&addr_QMainWindow_connect4, "QMainWindow_connect4")
	symaddr(&addr_QMainWindow_createPopupMenu_default, "QMainWindow_createPopupMenu_default")
	symaddr(&addr_QMainWindow_metaObject_default, "QMainWindow_metaObject_default")
	symaddr(&addr_QMainWindow_contextMenuEvent_default, "QMainWindow_contextMenuEvent_default")
	symaddr(&addr_QMainWindow_event_default, "QMainWindow_event_default")
	symaddr(&addr_QMainWindow_devType_default, "QMainWindow_devType_default")
	symaddr(&addr_QMainWindow_setVisible_default, "QMainWindow_setVisible_default")
	symaddr(&addr_QMainWindow_sizeHint_default, "QMainWindow_sizeHint_default")
	symaddr(&addr_QMainWindow_minimumSizeHint_default, "QMainWindow_minimumSizeHint_default")
	symaddr(&addr_QMainWindow_heightForWidth_default, "QMainWindow_heightForWidth_default")
	symaddr(&addr_QMainWindow_hasHeightForWidth_default, "QMainWindow_hasHeightForWidth_default")
	symaddr(&addr_QMainWindow_paintEngine_default, "QMainWindow_paintEngine_default")
	symaddr(&addr_QMainWindow_mousePressEvent_default, "QMainWindow_mousePressEvent_default")
	symaddr(&addr_QMainWindow_mouseReleaseEvent_default, "QMainWindow_mouseReleaseEvent_default")
	symaddr(&addr_QMainWindow_mouseDoubleClickEvent_default, "QMainWindow_mouseDoubleClickEvent_default")
	symaddr(&addr_QMainWindow_mouseMoveEvent_default, "QMainWindow_mouseMoveEvent_default")
	symaddr(&addr_QMainWindow_wheelEvent_default, "QMainWindow_wheelEvent_default")
	symaddr(&addr_QMainWindow_keyPressEvent_default, "QMainWindow_keyPressEvent_default")
	symaddr(&addr_QMainWindow_keyReleaseEvent_default, "QMainWindow_keyReleaseEvent_default")
	symaddr(&addr_QMainWindow_focusInEvent_default, "QMainWindow_focusInEvent_default")
	symaddr(&addr_QMainWindow_focusOutEvent_default, "QMainWindow_focusOutEvent_default")
	symaddr(&addr_QMainWindow_enterEvent_default, "QMainWindow_enterEvent_default")
	symaddr(&addr_QMainWindow_leaveEvent_default, "QMainWindow_leaveEvent_default")
	symaddr(&addr_QMainWindow_paintEvent_default, "QMainWindow_paintEvent_default")
	symaddr(&addr_QMainWindow_moveEvent_default, "QMainWindow_moveEvent_default")
	symaddr(&addr_QMainWindow_resizeEvent_default, "QMainWindow_resizeEvent_default")
	symaddr(&addr_QMainWindow_closeEvent_default, "QMainWindow_closeEvent_default")
	symaddr(&addr_QMainWindow_tabletEvent_default, "QMainWindow_tabletEvent_default")
	symaddr(&addr_QMainWindow_actionEvent_default, "QMainWindow_actionEvent_default")
	symaddr(&addr_QMainWindow_dragEnterEvent_default, "QMainWindow_dragEnterEvent_default")
	symaddr(&addr_QMainWindow_dragMoveEvent_default, "QMainWindow_dragMoveEvent_default")
	symaddr(&addr_QMainWindow_dragLeaveEvent_default, "QMainWindow_dragLeaveEvent_default")
	symaddr(&addr_QMainWindow_dropEvent_default, "QMainWindow_dropEvent_default")
	symaddr(&addr_QMainWindow_showEvent_default, "QMainWindow_showEvent_default")
	symaddr(&addr_QMainWindow_hideEvent_default, "QMainWindow_hideEvent_default")
	symaddr(&addr_QMainWindow_changeEvent_default, "QMainWindow_changeEvent_default")
	symaddr(&addr_QMainWindow_metric_default, "QMainWindow_metric_default")
	symaddr(&addr_QMainWindow_initPainter_default, "QMainWindow_initPainter_default")
	symaddr(&addr_QMainWindow_redirected_default, "QMainWindow_redirected_default")
	symaddr(&addr_QMainWindow_sharedPainter_default, "QMainWindow_sharedPainter_default")
	symaddr(&addr_QMainWindow_inputMethodEvent_default, "QMainWindow_inputMethodEvent_default")
	symaddr(&addr_QMainWindow_inputMethodQuery_default, "QMainWindow_inputMethodQuery_default")
	symaddr(&addr_QMainWindow_focusNextPrevChild_default, "QMainWindow_focusNextPrevChild_default")
	symaddr(&addr_QMainWindow_eventFilter_default, "QMainWindow_eventFilter_default")
	symaddr(&addr_QMainWindow_timerEvent_default, "QMainWindow_timerEvent_default")
	symaddr(&addr_QMainWindow_childEvent_default, "QMainWindow_childEvent_default")
	symaddr(&addr_QMainWindow_customEvent_default, "QMainWindow_customEvent_default")
	symaddr(&addr_QMainWindow_connectNotify_default, "QMainWindow_connectNotify_default")
	symaddr(&addr_QMainWindow_disconnectNotify_default, "QMainWindow_disconnectNotify_default")
	reg(&pg_QMainWindow_connect_iconSizeChanged, "QMainWindow_connect_iconSizeChanged")
	reg(&pg_QMainWindow_connect_toolButtonStyleChanged, "QMainWindow_connect_toolButtonStyleChanged")
	reg(&pg_QMainWindow_connect_tabifiedDockWidgetActivated, "QMainWindow_connect_tabifiedDockWidgetActivated")
	reg(&pg_QMainWindow_connect_windowTitleChanged, "QMainWindow_connect_windowTitleChanged")
	reg(&pg_QMainWindow_connect_windowIconChanged, "QMainWindow_connect_windowIconChanged")
	reg(&pg_QMainWindow_connect_windowIconTextChanged, "QMainWindow_connect_windowIconTextChanged")
	reg(&pg_QMainWindow_connect_customContextMenuRequested, "QMainWindow_connect_customContextMenuRequested")
	reg(&pg_QMainWindow_connect_destroyed, "QMainWindow_connect_destroyed")
	reg(&pg_QMainWindow_connect_destroyed2, "QMainWindow_connect_destroyed2")
	reg(&pg_QMainWindow_connect_objectNameChanged, "QMainWindow_connect_objectNameChanged")
	reg(&pg_QMainWindow__disconnect, "QMainWindow__disconnect")
	reg(&pg_QMainWindow__set_override, "QMainWindow__set_override")
	reg(&pg_QMainWindow_delete, "QMainWindow_delete")
}

type QMainWindow struct {
	h uintptr
	*QWidget
}

func bindQMainWindow(h uintptr) *QMainWindow {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QMainWindow_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMainWindow{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQMainWindow(h uintptr) *QMainWindow { return bindQMainWindow(h) }

func (this *QMainWindow) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMainWindow) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMainWindow(h unsafe.Pointer) *QMainWindow { return bindQMainWindow(uintptr(h)) }
func UnsafeNewQMainWindowOwned(h unsafe.Pointer) *QMainWindow { return newownedQMainWindow(uintptr(h)) }

func (this *QMainWindow) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQMainWindow(parent *QWidget) *QMainWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMainWindow_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMainWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMainWindow_virtIdx, QMainWindow_enable)
	return __o
}

func NewQMainWindow2() *QMainWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMainWindow_new2, uintptr(__goid))
	__o := newownedQMainWindow(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMainWindow_virtIdx, QMainWindow_enable)
	return __o
}

func NewQMainWindow3(parent *QWidget, flags WindowType) *QMainWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMainWindow_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQMainWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMainWindow_virtIdx, QMainWindow_enable)
	return __o
}

func QMainWindow_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMainWindow_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMainWindow) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMainWindow) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QMainWindow_setIconSize, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QMainWindow) ToolButtonStyle() ToolButtonStyle {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_toolButtonStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolButtonStyle(__r)
}

func (this *QMainWindow) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	purego.SyscallN(addr_QMainWindow_setToolButtonStyle, this.h, uintptr(toolButtonStyle))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) IsAnimated() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_isAnimated, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) IsDockNestingEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_isDockNestingEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_documentMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) SetDocumentMode(enabled bool) {
	purego.SyscallN(addr_QMainWindow_setDocumentMode, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) TabShape() QTabWidget__TabShape {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_tabShape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabShape(__r)
}

func (this *QMainWindow) SetTabShape(tabShape QTabWidget__TabShape) {
	purego.SyscallN(addr_QMainWindow_setTabShape, this.h, uintptr(tabShape))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) TabPosition(area DockWidgetArea) QTabWidget__TabPosition {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_tabPosition, this.h, uintptr(area))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabPosition(__r)
}

func (this *QMainWindow) SetTabPosition(areas DockWidgetArea, tabPosition QTabWidget__TabPosition) {
	purego.SyscallN(addr_QMainWindow_setTabPosition, this.h, uintptr(areas), uintptr(tabPosition))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) SetDockOptions(options QMainWindow__DockOption) {
	purego.SyscallN(addr_QMainWindow_setDockOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) DockOptions() QMainWindow__DockOption {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_dockOptions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMainWindow__DockOption(__r)
}

func (this *QMainWindow) IsSeparator(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_isSeparator, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QMainWindow) MenuBar() *QMenuBar {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_menuBar, this.h)
	runtime.KeepAlive(this)
	return bindQMenuBar(__r)
}

func (this *QMainWindow) SetMenuBar(menubar *QMenuBar) {
	purego.SyscallN(addr_QMainWindow_setMenuBar, this.h, uintptr(menubar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menubar)
}

func (this *QMainWindow) MenuWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_menuWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMainWindow) SetMenuWidget(menubar *QWidget) {
	purego.SyscallN(addr_QMainWindow_setMenuWidget, this.h, uintptr(menubar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menubar)
}

func (this *QMainWindow) StatusBar() *QStatusBar {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_statusBar, this.h)
	runtime.KeepAlive(this)
	return bindQStatusBar(__r)
}

func (this *QMainWindow) SetStatusBar(statusbar *QStatusBar) {
	purego.SyscallN(addr_QMainWindow_setStatusBar, this.h, uintptr(statusbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusbar)
}

func (this *QMainWindow) CentralWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_centralWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMainWindow) SetCentralWidget(widget *QWidget) {
	purego.SyscallN(addr_QMainWindow_setCentralWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QMainWindow) TakeCentralWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_takeCentralWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMainWindow) SetCorner(corner Corner, area DockWidgetArea) {
	purego.SyscallN(addr_QMainWindow_setCorner, this.h, uintptr(corner), uintptr(area))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Corner(corner Corner) DockWidgetArea {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_corner, this.h, uintptr(corner))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DockWidgetArea(__r)
}

func (this *QMainWindow) AddToolBarBreak() {
	purego.SyscallN(addr_QMainWindow_addToolBarBreak, this.h)
	runtime.KeepAlive(this)
}

func (this *QMainWindow) InsertToolBarBreak(before *QToolBar) {
	purego.SyscallN(addr_QMainWindow_insertToolBarBreak, this.h, uintptr(before.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
}

func (this *QMainWindow) AddToolBar(area ToolBarArea, toolbar *QToolBar) {
	purego.SyscallN(addr_QMainWindow_addToolBar, this.h, uintptr(area), uintptr(toolbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolbar)
}

func (this *QMainWindow) AddToolBarWithToolbar(toolbar *QToolBar) {
	purego.SyscallN(addr_QMainWindow_addToolBar2, this.h, uintptr(toolbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolbar)
}

func (this *QMainWindow) AddToolBarWithTitle(title string) *QToolBar {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMainWindow_addToolBar3, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	return bindQToolBar(__r)
}

func (this *QMainWindow) InsertToolBar(before *QToolBar, toolbar *QToolBar) {
	purego.SyscallN(addr_QMainWindow_insertToolBar, this.h, uintptr(before.UnsafePointer()), uintptr(toolbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(toolbar)
}

func (this *QMainWindow) RemoveToolBar(toolbar *QToolBar) {
	purego.SyscallN(addr_QMainWindow_removeToolBar, this.h, uintptr(toolbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolbar)
}

func (this *QMainWindow) RemoveToolBarBreak(before *QToolBar) {
	purego.SyscallN(addr_QMainWindow_removeToolBarBreak, this.h, uintptr(before.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
}

func (this *QMainWindow) UnifiedTitleAndToolBarOnMac() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_unifiedTitleAndToolBarOnMac, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) ToolBarArea(toolbar *QToolBar) ToolBarArea {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_toolBarArea, this.h, uintptr(toolbar.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolbar)
	return ToolBarArea(__r)
}

func (this *QMainWindow) ToolBarBreak(toolbar *QToolBar) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_toolBarBreak, this.h, uintptr(toolbar.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolbar)
	return __r != 0
}

func (this *QMainWindow) AddDockWidget(area DockWidgetArea, dockwidget *QDockWidget) {
	purego.SyscallN(addr_QMainWindow_addDockWidget, this.h, uintptr(area), uintptr(dockwidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
}

func (this *QMainWindow) AddDockWidget2(area DockWidgetArea, dockwidget *QDockWidget, orientation Orientation) {
	purego.SyscallN(addr_QMainWindow_addDockWidget2, this.h, uintptr(area), uintptr(dockwidget.UnsafePointer()), uintptr(orientation))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
}

func (this *QMainWindow) SplitDockWidget(after *QDockWidget, dockwidget *QDockWidget, orientation Orientation) {
	purego.SyscallN(addr_QMainWindow_splitDockWidget, this.h, uintptr(after.UnsafePointer()), uintptr(dockwidget.UnsafePointer()), uintptr(orientation))
	runtime.KeepAlive(this)
	runtime.KeepAlive(after)
	runtime.KeepAlive(dockwidget)
}

func (this *QMainWindow) TabifyDockWidget(first *QDockWidget, second *QDockWidget) {
	purego.SyscallN(addr_QMainWindow_tabifyDockWidget, this.h, uintptr(first.UnsafePointer()), uintptr(second.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)
}

func (this *QMainWindow) TabifiedDockWidgets(dockwidget *QDockWidget) []*QDockWidget {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMainWindow_tabifiedDockWidgets, this.h, uintptr(dockwidget.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
	__out := make([]*QDockWidget, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQDockWidget(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMainWindow) RemoveDockWidget(dockwidget *QDockWidget) {
	purego.SyscallN(addr_QMainWindow_removeDockWidget, this.h, uintptr(dockwidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
}

func (this *QMainWindow) RestoreDockWidget(dockwidget *QDockWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_restoreDockWidget, this.h, uintptr(dockwidget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
	return __r != 0
}

func (this *QMainWindow) DockWidgetArea(dockwidget *QDockWidget) DockWidgetArea {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_dockWidgetArea, this.h, uintptr(dockwidget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockwidget)
	return DockWidgetArea(__r)
}

func (this *QMainWindow) ResizeDocks(docks []*QDockWidget, sizes []int, orientation Orientation) {
	docks_arr := make([]uintptr, len(docks))
	for __i, __e := range docks {
		docks_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var docks_p uintptr
	if len(docks_arr) > 0 {
		docks_p = uintptr(unsafe.Pointer(&docks_arr[0]))
	}
	sizes_arr := make([]int32, len(sizes))
	for __i, __e := range sizes {
		sizes_arr[__i] = int32(__e)
	}
	var sizes_p uintptr
	if len(sizes_arr) > 0 {
		sizes_p = uintptr(unsafe.Pointer(&sizes_arr[0]))
	}
	purego.SyscallN(addr_QMainWindow_resizeDocks, this.h, docks_p, uintptr(len(docks)), sizes_p, uintptr(len(sizes)), uintptr(orientation))
	runtime.KeepAlive(this)
	runtime.KeepAlive(docks)
	runtime.KeepAlive(docks_arr)
	runtime.KeepAlive(sizes_arr)
}

func (this *QMainWindow) SaveState() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_saveState, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMainWindow) RestoreState(state *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_restoreState, this.h, uintptr(state.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	return __r != 0
}

func (this *QMainWindow) CreatePopupMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_createPopupMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QMainWindow) SetAnimated(enabled bool) {
	purego.SyscallN(addr_QMainWindow_setAnimated, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) SetDockNestingEnabled(enabled bool) {
	purego.SyscallN(addr_QMainWindow_setDockNestingEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) SetUnifiedTitleAndToolBarOnMac(set bool) {
	purego.SyscallN(addr_QMainWindow_setUnifiedTitleAndToolBarOnMac, this.h, uintptr(b2i(set)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) IconSizeChanged(iconSize *QSize) {
	purego.SyscallN(addr_QMainWindow_iconSizeChanged, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QMainWindow) ToolButtonStyleChanged(toolButtonStyle ToolButtonStyle) {
	purego.SyscallN(addr_QMainWindow_toolButtonStyleChanged, this.h, uintptr(toolButtonStyle))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) TabifiedDockWidgetActivated(dockWidget *QDockWidget) {
	purego.SyscallN(addr_QMainWindow_tabifiedDockWidgetActivated, this.h, uintptr(dockWidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dockWidget)
}

func QMainWindow_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMainWindow_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMainWindow_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_mouseGrabber)
	return bindQWidget(__r)
}

func QMainWindow_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_keyboardGrabber)
	return bindQWidget(__r)
}

func QMainWindow_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMainWindow_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMainWindow) UpdateMicroFocus() {
	purego.SyscallN(addr_QMainWindow_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Create() {
	purego.SyscallN(addr_QMainWindow_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Destroy() {
	purego.SyscallN(addr_QMainWindow_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMainWindow) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMainWindow_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMainWindow_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMainWindow_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMainWindow) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMainWindow) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMainWindow) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMainWindow) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMainWindow_DevicePixelRatioFScale() float64 {
	__r := pg_QMainWindow_devicePixelRatioFScale()
	return __r
}

func QMainWindow_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMainWindow_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMainWindow) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMainWindow_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMainWindow_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMainWindow_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMainWindow_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMainWindow_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMainWindow) AddToolBarBreakWithArea(area ToolBarArea) {
	purego.SyscallN(addr_QMainWindow_addToolBarBreak2, this.h, uintptr(area))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) SaveStateWithVersion(version int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_saveState2, this.h, uintptr(version))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMainWindow) RestoreState2(state *QByteArray, version int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_restoreState2, this.h, uintptr(state.UnsafePointer()), uintptr(version))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	return __r != 0
}

func QMainWindow_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMainWindow_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMainWindow) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMainWindow_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMainWindow_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMainWindow_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMainWindow_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMainWindow_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMainWindow_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMainWindow_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMainWindow_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMainWindow_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMainWindow) CreatePopupMenuDefault() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_createPopupMenu_default, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QMainWindow) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMainWindow) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QMainWindow_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMainWindow) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMainWindow) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMainWindow_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMainWindow) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMainWindow) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMainWindow) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMainWindow) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMainWindow) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMainWindow_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMainWindow_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMainWindow_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMainWindow_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QMainWindow_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMainWindow_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMainWindow_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMainWindow_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMainWindow_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QMainWindow_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMainWindow_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QMainWindow_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QMainWindow_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QMainWindow_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QMainWindow_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMainWindow_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QMainWindow_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QMainWindow_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QMainWindow_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QMainWindow_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QMainWindow_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QMainWindow_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QMainWindow_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMainWindow_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMainWindow) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMainWindow) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMainWindow_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMainWindow) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMainWindow) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMainWindow) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMainWindow_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMainWindow) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMainWindow_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMainWindow) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMainWindow) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMainWindow_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMainWindow) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QMainWindow_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMainWindow_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMainWindow_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMainWindow) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMainWindow_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMainWindow) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMainWindow_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMainWindow) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(iconSize uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(iconSize))) })
	pg_QMainWindow_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnToolButtonStyleChanged(fn func(ToolButtonStyle)) int64 {
	id := registerCallback(func(toolButtonStyle int32) { fn(ToolButtonStyle(toolButtonStyle)) })
	pg_QMainWindow_connect_toolButtonStyleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnTabifiedDockWidgetActivated(fn func(*QDockWidget)) int64 {
	id := registerCallback(func(dockWidget uintptr) { fn(UnsafeNewQDockWidget(unsafe.Pointer(dockWidget))) })
	pg_QMainWindow_connect_tabifiedDockWidgetActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMainWindow_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMainWindow_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMainWindow_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMainWindow_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMainWindow_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMainWindow_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMainWindow_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMainWindow) Disconnect(handle int64) {
	pg_QMainWindow__disconnect(handle)
	unregisterCallback(handle)
}

func QMainWindow_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1078400071543754474:
		return 1
	case -6698043222227006015:
		return 2
	case 1237236456312824991:
		return 3
	case -8659679893868524044:
		return 4
	case 2110182719375245575:
		return 5
	case 729690261457684101:
		return 6
	case 7049264599837532491:
		return 7
	case 4311534060495737339:
		return 8
	case -8266762397995423915:
		return 9
	case 6398007404756346381:
		return 10
	case -332090898078543319:
		return 11
	case -4415920049297795213:
		return 12
	case -4029073905697023887:
		return 13
	case -4255494004099806581:
		return 14
	case 1919988340934511316:
		return 15
	case 1023791977679793185:
		return 16
	case 772598475950050299:
		return 17
	case 975043767033552552:
		return 18
	case -385596010576048993:
		return 19
	case 6316270346668237895:
		return 20
	case -551070009623711640:
		return 21
	case 5290442163167458465:
		return 22
	case -9161949150489525396:
		return 23
	case -4302917285655014661:
		return 24
	case -1989937050019464783:
		return 25
	case -212068866860701141:
		return 26
	case 1697444140088555581:
		return 27
	case 206122715919917589:
		return 28
	case -1781545137352745978:
		return 29
	case -3139904158373725006:
		return 30
	case 9107058682754591716:
		return 31
	case -6825369174751916:
		return 32
	case -199718707876719129:
		return 33
	case -2039166586246443706:
		return 34
	case 4959376063826165943:
		return 35
	case 427629649267311777:
		return 36
	case 6477500296980050210:
		return 37
	case 4127240633310655048:
		return 38
	case 5792771328741774991:
		return 39
	case 5438819661942843222:
		return 40
	case 7345853097340709638:
		return 41
	case 5898328517505784191:
		return 42
	case 5217623514630858297:
		return 43
	case -5409927581812374498:
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

func QMainWindow_enable(h uintptr, idx int32) { pg_QMainWindow__set_override(h, idx) }

func (this *QMainWindow) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMainWindow) OnCreatePopupMenu(fn func() *QMenu) {
	registerVirt(this.UnsafeGoID(), 1078400071543754474, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "createPopupMenu"
}

func (this *QMainWindow) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QMainWindow) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QMainWindow) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMainWindow) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMainWindow) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMainWindow) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMainWindow) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMainWindow) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMainWindow) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMainWindow) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QMainWindow) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QMainWindow) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMainWindow) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QMainWindow) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QMainWindow) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QMainWindow) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMainWindow) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QMainWindow) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QMainWindow) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QMainWindow) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QMainWindow) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QMainWindow) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QMainWindow) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QMainWindow) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QMainWindow) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMainWindow) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QMainWindow) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QMainWindow) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QMainWindow) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QMainWindow) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QMainWindow) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QMainWindow) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QMainWindow) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMainWindow) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QMainWindow) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMainWindow) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMainWindow) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMainWindow) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMainWindow) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMainWindow) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMainWindow) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMainWindow) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMainWindow) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QMainWindow) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMainWindow) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMainWindow) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMainWindow) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMainWindow) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMainWindow_delete(this.h)
	}
	this.UnsafeRelease()
}
