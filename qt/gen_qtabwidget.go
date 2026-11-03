package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTabWidget_chain uintptr
	addr_QTabWidget_new uintptr
	addr_QTabWidget_new2 uintptr
	addr_QTabWidget_tr uintptr
	addr_QTabWidget_addTab uintptr
	addr_QTabWidget_addTab2 uintptr
	addr_QTabWidget_insertTab uintptr
	addr_QTabWidget_insertTab2 uintptr
	addr_QTabWidget_removeTab uintptr
	addr_QTabWidget_isTabEnabled uintptr
	addr_QTabWidget_setTabEnabled uintptr
	addr_QTabWidget_isTabVisible uintptr
	addr_QTabWidget_setTabVisible uintptr
	addr_QTabWidget_tabText uintptr
	addr_QTabWidget_setTabText uintptr
	addr_QTabWidget_tabIcon uintptr
	addr_QTabWidget_setTabIcon uintptr
	addr_QTabWidget_setTabToolTip uintptr
	addr_QTabWidget_tabToolTip uintptr
	addr_QTabWidget_setTabWhatsThis uintptr
	addr_QTabWidget_tabWhatsThis uintptr
	addr_QTabWidget_currentIndex uintptr
	addr_QTabWidget_currentWidget uintptr
	addr_QTabWidget_widget uintptr
	addr_QTabWidget_indexOf uintptr
	addr_QTabWidget_count uintptr
	addr_QTabWidget_tabPosition uintptr
	addr_QTabWidget_setTabPosition uintptr
	addr_QTabWidget_tabsClosable uintptr
	addr_QTabWidget_setTabsClosable uintptr
	addr_QTabWidget_isMovable uintptr
	addr_QTabWidget_setMovable uintptr
	addr_QTabWidget_tabShape uintptr
	addr_QTabWidget_setTabShape uintptr
	addr_QTabWidget_setCornerWidget uintptr
	addr_QTabWidget_cornerWidget uintptr
	addr_QTabWidget_elideMode uintptr
	addr_QTabWidget_setElideMode uintptr
	addr_QTabWidget_iconSize uintptr
	addr_QTabWidget_setIconSize uintptr
	addr_QTabWidget_usesScrollButtons uintptr
	addr_QTabWidget_setUsesScrollButtons uintptr
	addr_QTabWidget_documentMode uintptr
	addr_QTabWidget_setDocumentMode uintptr
	addr_QTabWidget_tabBarAutoHide uintptr
	addr_QTabWidget_setTabBarAutoHide uintptr
	addr_QTabWidget_clear uintptr
	addr_QTabWidget_tabBar uintptr
	addr_QTabWidget_setCurrentIndex uintptr
	addr_QTabWidget_setCurrentWidget uintptr
	addr_QTabWidget_currentChanged uintptr
	addr_QTabWidget_tabCloseRequested uintptr
	addr_QTabWidget_tabBarClicked uintptr
	addr_QTabWidget_tabBarDoubleClicked uintptr
	addr_QTabWidget_setTabBar uintptr
	addr_QTabWidget_setTabOrder uintptr
	addr_QTabWidget_mouseGrabber uintptr
	addr_QTabWidget_keyboardGrabber uintptr
	addr_QTabWidget_find uintptr
	addr_QTabWidget_createWindowContainer uintptr
	addr_QTabWidget_updateMicroFocus uintptr
	addr_QTabWidget_create uintptr
	addr_QTabWidget_destroy uintptr
	addr_QTabWidget_focusNextChild uintptr
	addr_QTabWidget_focusPreviousChild uintptr
	addr_QTabWidget_connect uintptr
	addr_QTabWidget_connect2 uintptr
	addr_QTabWidget_sender uintptr
	addr_QTabWidget_senderSignalIndex uintptr
	addr_QTabWidget_receivers uintptr
	addr_QTabWidget_isSignalConnected uintptr
	pg_QTabWidget_devicePixelRatioFScale func() float64
	pg_QTabWidget_encodeMetricF func(int32, float64) int32
	pg_QTabWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTabWidget_tr2 uintptr
	addr_QTabWidget_tr3 uintptr
	addr_QTabWidget_setCornerWidget2 uintptr
	addr_QTabWidget_cornerWidget2 uintptr
	addr_QTabWidget_createWindowContainer2 uintptr
	addr_QTabWidget_createWindowContainer3 uintptr
	addr_QTabWidget_updateMicroFocus2 uintptr
	addr_QTabWidget_create2 uintptr
	addr_QTabWidget_create3 uintptr
	addr_QTabWidget_create4 uintptr
	addr_QTabWidget_destroy2 uintptr
	addr_QTabWidget_destroy3 uintptr
	addr_QTabWidget_connect3 uintptr
	addr_QTabWidget_connect4 uintptr
	addr_QTabWidget_metaObject_default uintptr
	addr_QTabWidget_sizeHint_default uintptr
	addr_QTabWidget_minimumSizeHint_default uintptr
	addr_QTabWidget_heightForWidth_default uintptr
	addr_QTabWidget_hasHeightForWidth_default uintptr
	addr_QTabWidget_tabInserted_default uintptr
	addr_QTabWidget_tabRemoved_default uintptr
	addr_QTabWidget_showEvent_default uintptr
	addr_QTabWidget_resizeEvent_default uintptr
	addr_QTabWidget_keyPressEvent_default uintptr
	addr_QTabWidget_paintEvent_default uintptr
	addr_QTabWidget_changeEvent_default uintptr
	addr_QTabWidget_event_default uintptr
	addr_QTabWidget_initStyleOption_default uintptr
	addr_QTabWidget_devType_default uintptr
	addr_QTabWidget_setVisible_default uintptr
	addr_QTabWidget_paintEngine_default uintptr
	addr_QTabWidget_mousePressEvent_default uintptr
	addr_QTabWidget_mouseReleaseEvent_default uintptr
	addr_QTabWidget_mouseDoubleClickEvent_default uintptr
	addr_QTabWidget_mouseMoveEvent_default uintptr
	addr_QTabWidget_wheelEvent_default uintptr
	addr_QTabWidget_keyReleaseEvent_default uintptr
	addr_QTabWidget_focusInEvent_default uintptr
	addr_QTabWidget_focusOutEvent_default uintptr
	addr_QTabWidget_enterEvent_default uintptr
	addr_QTabWidget_leaveEvent_default uintptr
	addr_QTabWidget_moveEvent_default uintptr
	addr_QTabWidget_closeEvent_default uintptr
	addr_QTabWidget_contextMenuEvent_default uintptr
	addr_QTabWidget_tabletEvent_default uintptr
	addr_QTabWidget_actionEvent_default uintptr
	addr_QTabWidget_dragEnterEvent_default uintptr
	addr_QTabWidget_dragMoveEvent_default uintptr
	addr_QTabWidget_dragLeaveEvent_default uintptr
	addr_QTabWidget_dropEvent_default uintptr
	addr_QTabWidget_hideEvent_default uintptr
	addr_QTabWidget_metric_default uintptr
	addr_QTabWidget_initPainter_default uintptr
	addr_QTabWidget_redirected_default uintptr
	addr_QTabWidget_sharedPainter_default uintptr
	addr_QTabWidget_inputMethodEvent_default uintptr
	addr_QTabWidget_inputMethodQuery_default uintptr
	addr_QTabWidget_focusNextPrevChild_default uintptr
	addr_QTabWidget_eventFilter_default uintptr
	addr_QTabWidget_timerEvent_default uintptr
	addr_QTabWidget_childEvent_default uintptr
	addr_QTabWidget_customEvent_default uintptr
	addr_QTabWidget_connectNotify_default uintptr
	addr_QTabWidget_disconnectNotify_default uintptr
	pg_QTabWidget_connect_currentChanged func(uintptr, int64)
	pg_QTabWidget_connect_tabCloseRequested func(uintptr, int64)
	pg_QTabWidget_connect_tabBarClicked func(uintptr, int64)
	pg_QTabWidget_connect_tabBarDoubleClicked func(uintptr, int64)
	pg_QTabWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QTabWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QTabWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTabWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTabWidget_connect_destroyed func(uintptr, int64)
	pg_QTabWidget_connect_destroyed2 func(uintptr, int64)
	pg_QTabWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QTabWidget__disconnect func(int64)
	pg_QTabWidget__set_override func(uintptr, int32)
	pg_QTabWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTabWidget_chain, "QTabWidget_chain")
	symaddr(&addr_QTabWidget_new, "QTabWidget_new")
	symaddr(&addr_QTabWidget_new2, "QTabWidget_new2")
	symaddr(&addr_QTabWidget_tr, "QTabWidget_tr")
	symaddr(&addr_QTabWidget_addTab, "QTabWidget_addTab")
	symaddr(&addr_QTabWidget_addTab2, "QTabWidget_addTab2")
	symaddr(&addr_QTabWidget_insertTab, "QTabWidget_insertTab")
	symaddr(&addr_QTabWidget_insertTab2, "QTabWidget_insertTab2")
	symaddr(&addr_QTabWidget_removeTab, "QTabWidget_removeTab")
	symaddr(&addr_QTabWidget_isTabEnabled, "QTabWidget_isTabEnabled")
	symaddr(&addr_QTabWidget_setTabEnabled, "QTabWidget_setTabEnabled")
	symaddr(&addr_QTabWidget_isTabVisible, "QTabWidget_isTabVisible")
	symaddr(&addr_QTabWidget_setTabVisible, "QTabWidget_setTabVisible")
	symaddr(&addr_QTabWidget_tabText, "QTabWidget_tabText")
	symaddr(&addr_QTabWidget_setTabText, "QTabWidget_setTabText")
	symaddr(&addr_QTabWidget_tabIcon, "QTabWidget_tabIcon")
	symaddr(&addr_QTabWidget_setTabIcon, "QTabWidget_setTabIcon")
	symaddr(&addr_QTabWidget_setTabToolTip, "QTabWidget_setTabToolTip")
	symaddr(&addr_QTabWidget_tabToolTip, "QTabWidget_tabToolTip")
	symaddr(&addr_QTabWidget_setTabWhatsThis, "QTabWidget_setTabWhatsThis")
	symaddr(&addr_QTabWidget_tabWhatsThis, "QTabWidget_tabWhatsThis")
	symaddr(&addr_QTabWidget_currentIndex, "QTabWidget_currentIndex")
	symaddr(&addr_QTabWidget_currentWidget, "QTabWidget_currentWidget")
	symaddr(&addr_QTabWidget_widget, "QTabWidget_widget")
	symaddr(&addr_QTabWidget_indexOf, "QTabWidget_indexOf")
	symaddr(&addr_QTabWidget_count, "QTabWidget_count")
	symaddr(&addr_QTabWidget_tabPosition, "QTabWidget_tabPosition")
	symaddr(&addr_QTabWidget_setTabPosition, "QTabWidget_setTabPosition")
	symaddr(&addr_QTabWidget_tabsClosable, "QTabWidget_tabsClosable")
	symaddr(&addr_QTabWidget_setTabsClosable, "QTabWidget_setTabsClosable")
	symaddr(&addr_QTabWidget_isMovable, "QTabWidget_isMovable")
	symaddr(&addr_QTabWidget_setMovable, "QTabWidget_setMovable")
	symaddr(&addr_QTabWidget_tabShape, "QTabWidget_tabShape")
	symaddr(&addr_QTabWidget_setTabShape, "QTabWidget_setTabShape")
	symaddr(&addr_QTabWidget_setCornerWidget, "QTabWidget_setCornerWidget")
	symaddr(&addr_QTabWidget_cornerWidget, "QTabWidget_cornerWidget")
	symaddr(&addr_QTabWidget_elideMode, "QTabWidget_elideMode")
	symaddr(&addr_QTabWidget_setElideMode, "QTabWidget_setElideMode")
	symaddr(&addr_QTabWidget_iconSize, "QTabWidget_iconSize")
	symaddr(&addr_QTabWidget_setIconSize, "QTabWidget_setIconSize")
	symaddr(&addr_QTabWidget_usesScrollButtons, "QTabWidget_usesScrollButtons")
	symaddr(&addr_QTabWidget_setUsesScrollButtons, "QTabWidget_setUsesScrollButtons")
	symaddr(&addr_QTabWidget_documentMode, "QTabWidget_documentMode")
	symaddr(&addr_QTabWidget_setDocumentMode, "QTabWidget_setDocumentMode")
	symaddr(&addr_QTabWidget_tabBarAutoHide, "QTabWidget_tabBarAutoHide")
	symaddr(&addr_QTabWidget_setTabBarAutoHide, "QTabWidget_setTabBarAutoHide")
	symaddr(&addr_QTabWidget_clear, "QTabWidget_clear")
	symaddr(&addr_QTabWidget_tabBar, "QTabWidget_tabBar")
	symaddr(&addr_QTabWidget_setCurrentIndex, "QTabWidget_setCurrentIndex")
	symaddr(&addr_QTabWidget_setCurrentWidget, "QTabWidget_setCurrentWidget")
	symaddr(&addr_QTabWidget_currentChanged, "QTabWidget_currentChanged")
	symaddr(&addr_QTabWidget_tabCloseRequested, "QTabWidget_tabCloseRequested")
	symaddr(&addr_QTabWidget_tabBarClicked, "QTabWidget_tabBarClicked")
	symaddr(&addr_QTabWidget_tabBarDoubleClicked, "QTabWidget_tabBarDoubleClicked")
	symaddr(&addr_QTabWidget_setTabBar, "QTabWidget_setTabBar")
	symaddr(&addr_QTabWidget_setTabOrder, "QTabWidget_setTabOrder")
	symaddr(&addr_QTabWidget_mouseGrabber, "QTabWidget_mouseGrabber")
	symaddr(&addr_QTabWidget_keyboardGrabber, "QTabWidget_keyboardGrabber")
	symaddr(&addr_QTabWidget_find, "QTabWidget_find")
	symaddr(&addr_QTabWidget_createWindowContainer, "QTabWidget_createWindowContainer")
	symaddr(&addr_QTabWidget_updateMicroFocus, "QTabWidget_updateMicroFocus")
	symaddr(&addr_QTabWidget_create, "QTabWidget_create")
	symaddr(&addr_QTabWidget_destroy, "QTabWidget_destroy")
	symaddr(&addr_QTabWidget_focusNextChild, "QTabWidget_focusNextChild")
	symaddr(&addr_QTabWidget_focusPreviousChild, "QTabWidget_focusPreviousChild")
	symaddr(&addr_QTabWidget_connect, "QTabWidget_connect")
	symaddr(&addr_QTabWidget_connect2, "QTabWidget_connect2")
	symaddr(&addr_QTabWidget_sender, "QTabWidget_sender")
	symaddr(&addr_QTabWidget_senderSignalIndex, "QTabWidget_senderSignalIndex")
	symaddr(&addr_QTabWidget_receivers, "QTabWidget_receivers")
	symaddr(&addr_QTabWidget_isSignalConnected, "QTabWidget_isSignalConnected")
	reg(&pg_QTabWidget_devicePixelRatioFScale, "QTabWidget_devicePixelRatioFScale")
	reg(&pg_QTabWidget_encodeMetricF, "QTabWidget_encodeMetricF")
	reg(&pg_QTabWidget_getDecodedMetricF, "QTabWidget_getDecodedMetricF")
	symaddr(&addr_QTabWidget_tr2, "QTabWidget_tr2")
	symaddr(&addr_QTabWidget_tr3, "QTabWidget_tr3")
	symaddr(&addr_QTabWidget_setCornerWidget2, "QTabWidget_setCornerWidget2")
	symaddr(&addr_QTabWidget_cornerWidget2, "QTabWidget_cornerWidget2")
	symaddr(&addr_QTabWidget_createWindowContainer2, "QTabWidget_createWindowContainer2")
	symaddr(&addr_QTabWidget_createWindowContainer3, "QTabWidget_createWindowContainer3")
	symaddr(&addr_QTabWidget_updateMicroFocus2, "QTabWidget_updateMicroFocus2")
	symaddr(&addr_QTabWidget_create2, "QTabWidget_create2")
	symaddr(&addr_QTabWidget_create3, "QTabWidget_create3")
	symaddr(&addr_QTabWidget_create4, "QTabWidget_create4")
	symaddr(&addr_QTabWidget_destroy2, "QTabWidget_destroy2")
	symaddr(&addr_QTabWidget_destroy3, "QTabWidget_destroy3")
	symaddr(&addr_QTabWidget_connect3, "QTabWidget_connect3")
	symaddr(&addr_QTabWidget_connect4, "QTabWidget_connect4")
	symaddr(&addr_QTabWidget_metaObject_default, "QTabWidget_metaObject_default")
	symaddr(&addr_QTabWidget_sizeHint_default, "QTabWidget_sizeHint_default")
	symaddr(&addr_QTabWidget_minimumSizeHint_default, "QTabWidget_minimumSizeHint_default")
	symaddr(&addr_QTabWidget_heightForWidth_default, "QTabWidget_heightForWidth_default")
	symaddr(&addr_QTabWidget_hasHeightForWidth_default, "QTabWidget_hasHeightForWidth_default")
	symaddr(&addr_QTabWidget_tabInserted_default, "QTabWidget_tabInserted_default")
	symaddr(&addr_QTabWidget_tabRemoved_default, "QTabWidget_tabRemoved_default")
	symaddr(&addr_QTabWidget_showEvent_default, "QTabWidget_showEvent_default")
	symaddr(&addr_QTabWidget_resizeEvent_default, "QTabWidget_resizeEvent_default")
	symaddr(&addr_QTabWidget_keyPressEvent_default, "QTabWidget_keyPressEvent_default")
	symaddr(&addr_QTabWidget_paintEvent_default, "QTabWidget_paintEvent_default")
	symaddr(&addr_QTabWidget_changeEvent_default, "QTabWidget_changeEvent_default")
	symaddr(&addr_QTabWidget_event_default, "QTabWidget_event_default")
	symaddr(&addr_QTabWidget_initStyleOption_default, "QTabWidget_initStyleOption_default")
	symaddr(&addr_QTabWidget_devType_default, "QTabWidget_devType_default")
	symaddr(&addr_QTabWidget_setVisible_default, "QTabWidget_setVisible_default")
	symaddr(&addr_QTabWidget_paintEngine_default, "QTabWidget_paintEngine_default")
	symaddr(&addr_QTabWidget_mousePressEvent_default, "QTabWidget_mousePressEvent_default")
	symaddr(&addr_QTabWidget_mouseReleaseEvent_default, "QTabWidget_mouseReleaseEvent_default")
	symaddr(&addr_QTabWidget_mouseDoubleClickEvent_default, "QTabWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QTabWidget_mouseMoveEvent_default, "QTabWidget_mouseMoveEvent_default")
	symaddr(&addr_QTabWidget_wheelEvent_default, "QTabWidget_wheelEvent_default")
	symaddr(&addr_QTabWidget_keyReleaseEvent_default, "QTabWidget_keyReleaseEvent_default")
	symaddr(&addr_QTabWidget_focusInEvent_default, "QTabWidget_focusInEvent_default")
	symaddr(&addr_QTabWidget_focusOutEvent_default, "QTabWidget_focusOutEvent_default")
	symaddr(&addr_QTabWidget_enterEvent_default, "QTabWidget_enterEvent_default")
	symaddr(&addr_QTabWidget_leaveEvent_default, "QTabWidget_leaveEvent_default")
	symaddr(&addr_QTabWidget_moveEvent_default, "QTabWidget_moveEvent_default")
	symaddr(&addr_QTabWidget_closeEvent_default, "QTabWidget_closeEvent_default")
	symaddr(&addr_QTabWidget_contextMenuEvent_default, "QTabWidget_contextMenuEvent_default")
	symaddr(&addr_QTabWidget_tabletEvent_default, "QTabWidget_tabletEvent_default")
	symaddr(&addr_QTabWidget_actionEvent_default, "QTabWidget_actionEvent_default")
	symaddr(&addr_QTabWidget_dragEnterEvent_default, "QTabWidget_dragEnterEvent_default")
	symaddr(&addr_QTabWidget_dragMoveEvent_default, "QTabWidget_dragMoveEvent_default")
	symaddr(&addr_QTabWidget_dragLeaveEvent_default, "QTabWidget_dragLeaveEvent_default")
	symaddr(&addr_QTabWidget_dropEvent_default, "QTabWidget_dropEvent_default")
	symaddr(&addr_QTabWidget_hideEvent_default, "QTabWidget_hideEvent_default")
	symaddr(&addr_QTabWidget_metric_default, "QTabWidget_metric_default")
	symaddr(&addr_QTabWidget_initPainter_default, "QTabWidget_initPainter_default")
	symaddr(&addr_QTabWidget_redirected_default, "QTabWidget_redirected_default")
	symaddr(&addr_QTabWidget_sharedPainter_default, "QTabWidget_sharedPainter_default")
	symaddr(&addr_QTabWidget_inputMethodEvent_default, "QTabWidget_inputMethodEvent_default")
	symaddr(&addr_QTabWidget_inputMethodQuery_default, "QTabWidget_inputMethodQuery_default")
	symaddr(&addr_QTabWidget_focusNextPrevChild_default, "QTabWidget_focusNextPrevChild_default")
	symaddr(&addr_QTabWidget_eventFilter_default, "QTabWidget_eventFilter_default")
	symaddr(&addr_QTabWidget_timerEvent_default, "QTabWidget_timerEvent_default")
	symaddr(&addr_QTabWidget_childEvent_default, "QTabWidget_childEvent_default")
	symaddr(&addr_QTabWidget_customEvent_default, "QTabWidget_customEvent_default")
	symaddr(&addr_QTabWidget_connectNotify_default, "QTabWidget_connectNotify_default")
	symaddr(&addr_QTabWidget_disconnectNotify_default, "QTabWidget_disconnectNotify_default")
	reg(&pg_QTabWidget_connect_currentChanged, "QTabWidget_connect_currentChanged")
	reg(&pg_QTabWidget_connect_tabCloseRequested, "QTabWidget_connect_tabCloseRequested")
	reg(&pg_QTabWidget_connect_tabBarClicked, "QTabWidget_connect_tabBarClicked")
	reg(&pg_QTabWidget_connect_tabBarDoubleClicked, "QTabWidget_connect_tabBarDoubleClicked")
	reg(&pg_QTabWidget_connect_windowTitleChanged, "QTabWidget_connect_windowTitleChanged")
	reg(&pg_QTabWidget_connect_windowIconChanged, "QTabWidget_connect_windowIconChanged")
	reg(&pg_QTabWidget_connect_windowIconTextChanged, "QTabWidget_connect_windowIconTextChanged")
	reg(&pg_QTabWidget_connect_customContextMenuRequested, "QTabWidget_connect_customContextMenuRequested")
	reg(&pg_QTabWidget_connect_destroyed, "QTabWidget_connect_destroyed")
	reg(&pg_QTabWidget_connect_destroyed2, "QTabWidget_connect_destroyed2")
	reg(&pg_QTabWidget_connect_objectNameChanged, "QTabWidget_connect_objectNameChanged")
	reg(&pg_QTabWidget__disconnect, "QTabWidget__disconnect")
	reg(&pg_QTabWidget__set_override, "QTabWidget__set_override")
	reg(&pg_QTabWidget_delete, "QTabWidget_delete")
}

type QTabWidget struct {
	h uintptr
	*QWidget
}

func bindQTabWidget(h uintptr) *QTabWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTabWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTabWidget{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTabWidget(h uintptr) *QTabWidget { return bindQTabWidget(h) }

func (this *QTabWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTabWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTabWidget(h unsafe.Pointer) *QTabWidget { return bindQTabWidget(uintptr(h)) }
func UnsafeNewQTabWidgetOwned(h unsafe.Pointer) *QTabWidget { return newownedQTabWidget(uintptr(h)) }

func (this *QTabWidget) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQTabWidget(parent *QWidget) *QTabWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTabWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTabWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTabWidget_virtIdx, QTabWidget_enable)
	return __o
}

func NewQTabWidget2() *QTabWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTabWidget_new2, uintptr(__goid))
	__o := newownedQTabWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTabWidget_virtIdx, QTabWidget_enable)
	return __o
}

func QTabWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTabWidget) AddTab(widget *QWidget, arg2 string) int {
	var arg2_p uintptr
	if len(arg2) > 0 {
		arg2_p = uintptr(unsafe.Pointer(unsafe.StringData(arg2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_addTab, this.h, uintptr(widget.UnsafePointer()), arg2_p, uintptr(len(arg2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(arg2)
	return int(__r)
}

func (this *QTabWidget) AddTab2(widget *QWidget, icon *QIcon, label string) int {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_addTab2, this.h, uintptr(widget.UnsafePointer()), uintptr(icon.UnsafePointer()), label_p, uintptr(len(label)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(label)
	return int(__r)
}

func (this *QTabWidget) InsertTab(index int, widget *QWidget, arg3 string) int {
	var arg3_p uintptr
	if len(arg3) > 0 {
		arg3_p = uintptr(unsafe.Pointer(unsafe.StringData(arg3)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_insertTab, this.h, uintptr(index), uintptr(widget.UnsafePointer()), arg3_p, uintptr(len(arg3)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(arg3)
	return int(__r)
}

func (this *QTabWidget) InsertTab2(index int, widget *QWidget, icon *QIcon, label string) int {
	var label_p uintptr
	if len(label) > 0 {
		label_p = uintptr(unsafe.Pointer(unsafe.StringData(label)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_insertTab2, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(icon.UnsafePointer()), label_p, uintptr(len(label)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(label)
	return int(__r)
}

func (this *QTabWidget) RemoveTab(index int) {
	purego.SyscallN(addr_QTabWidget_removeTab, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) IsTabEnabled(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_isTabEnabled, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetTabEnabled(index int, enabled bool) {
	purego.SyscallN(addr_QTabWidget_setTabEnabled, this.h, uintptr(index), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) IsTabVisible(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_isTabVisible, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetTabVisible(index int, visible bool) {
	purego.SyscallN(addr_QTabWidget_setTabVisible, this.h, uintptr(index), uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabText(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tabText, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabWidget) SetTabText(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTabWidget_setTabText, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTabWidget) TabIcon(index int) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tabIcon, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QTabWidget) SetTabIcon(index int, icon *QIcon) {
	purego.SyscallN(addr_QTabWidget_setTabIcon, this.h, uintptr(index), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QTabWidget) SetTabToolTip(index int, tip string) {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	purego.SyscallN(addr_QTabWidget_setTabToolTip, this.h, uintptr(index), tip_p, uintptr(len(tip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tip)
}

func (this *QTabWidget) TabToolTip(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tabToolTip, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabWidget) SetTabWhatsThis(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTabWidget_setTabWhatsThis, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTabWidget) TabWhatsThis(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tabWhatsThis, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabWidget) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) CurrentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_currentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QTabWidget) Widget(index int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_widget, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QTabWidget) IndexOf(widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_indexOf, this.h, uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QTabWidget) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) TabPosition() QTabWidget__TabPosition {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_tabPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabPosition(__r)
}

func (this *QTabWidget) SetTabPosition(position QTabWidget__TabPosition) {
	purego.SyscallN(addr_QTabWidget_setTabPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabsClosable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_tabsClosable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetTabsClosable(closeable bool) {
	purego.SyscallN(addr_QTabWidget_setTabsClosable, this.h, uintptr(b2i(closeable)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) IsMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_isMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetMovable(movable bool) {
	purego.SyscallN(addr_QTabWidget_setMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabShape() QTabWidget__TabShape {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_tabShape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabWidget__TabShape(__r)
}

func (this *QTabWidget) SetTabShape(s QTabWidget__TabShape) {
	purego.SyscallN(addr_QTabWidget_setTabShape, this.h, uintptr(s))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) SetCornerWidget(w *QWidget) {
	purego.SyscallN(addr_QTabWidget_setCornerWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QTabWidget) CornerWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_cornerWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QTabWidget) ElideMode() TextElideMode {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_elideMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextElideMode(__r)
}

func (this *QTabWidget) SetElideMode(mode TextElideMode) {
	purego.SyscallN(addr_QTabWidget_setElideMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabWidget) SetIconSize(size *QSize) {
	purego.SyscallN(addr_QTabWidget_setIconSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QTabWidget) UsesScrollButtons() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_usesScrollButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	purego.SyscallN(addr_QTabWidget_setUsesScrollButtons, this.h, uintptr(b2i(useButtons)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_documentMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetDocumentMode(set bool) {
	purego.SyscallN(addr_QTabWidget_setDocumentMode, this.h, uintptr(b2i(set)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabBarAutoHide() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_tabBarAutoHide, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	purego.SyscallN(addr_QTabWidget_setTabBarAutoHide, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Clear() {
	purego.SyscallN(addr_QTabWidget_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabBar() *QTabBar {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tabBar, this.h)
	runtime.KeepAlive(this)
	return bindQTabBar(__r)
}

func (this *QTabWidget) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QTabWidget_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) SetCurrentWidget(widget *QWidget) {
	purego.SyscallN(addr_QTabWidget_setCurrentWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QTabWidget) CurrentChanged(index int) {
	purego.SyscallN(addr_QTabWidget_currentChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabCloseRequested(index int) {
	purego.SyscallN(addr_QTabWidget_tabCloseRequested, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabBarClicked(index int) {
	purego.SyscallN(addr_QTabWidget_tabBarClicked, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabBarDoubleClicked(index int) {
	purego.SyscallN(addr_QTabWidget_tabBarDoubleClicked, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) SetTabBar(arg1 *QTabBar) {
	purego.SyscallN(addr_QTabWidget_setTabBar, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QTabWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTabWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTabWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QTabWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QTabWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTabWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTabWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QTabWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Create() {
	purego.SyscallN(addr_QTabWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Destroy() {
	purego.SyscallN(addr_QTabWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTabWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTabWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTabWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTabWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTabWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTabWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTabWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QTabWidget_devicePixelRatioFScale()
	return __r
}

func QTabWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTabWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTabWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTabWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTabWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTabWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTabWidget) SetCornerWidget2(w *QWidget, corner Corner) {
	purego.SyscallN(addr_QTabWidget_setCornerWidget2, this.h, uintptr(w.UnsafePointer()), uintptr(corner))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QTabWidget) CornerWidgetWithCorner(corner Corner) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_cornerWidget2, this.h, uintptr(corner))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func QTabWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTabWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTabWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTabWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTabWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTabWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTabWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTabWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTabWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTabWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTabWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTabWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTabWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTabWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabWidget) HeightForWidthDefault(width int) int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_heightForWidth_default, this.h, uintptr(width))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) TabInsertedDefault(index int) {
	purego.SyscallN(addr_QTabWidget_tabInserted_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) TabRemovedDefault(index int) {
	purego.SyscallN(addr_QTabWidget_tabRemoved_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QTabWidget_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QTabWidget_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QTabWidget_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QTabWidget_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QTabWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTabWidget) InitStyleOptionDefault(option *QStyleOptionTabWidgetFrame) {
	purego.SyscallN(addr_QTabWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTabWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTabWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTabWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTabWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTabWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTabWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTabWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QTabWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QTabWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTabWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTabWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTabWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTabWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTabWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTabWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTabWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QTabWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTabWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTabWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTabWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTabWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTabWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTabWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QTabWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTabWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTabWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTabWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTabWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QTabWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTabWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTabWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTabWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTabWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTabWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTabWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTabWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTabWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTabWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTabWidget) OnCurrentChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabWidget_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnTabCloseRequested(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabWidget_connect_tabCloseRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnTabBarClicked(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabWidget_connect_tabBarClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnTabBarDoubleClicked(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabWidget_connect_tabBarDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTabWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTabWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTabWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTabWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTabWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTabWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTabWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabWidget) Disconnect(handle int64) {
	pg_QTabWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QTabWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 4311534060495737339:
		return 3
	case -8266762397995423915:
		return 4
	case -2734286415018532480:
		return 5
	case -3307029722643392856:
		return 6
	case -6825369174751916:
		return 7
	case -4302917285655014661:
		return 8
	case 1023791977679793185:
		return 9
	case 5290442163167458465:
		return 10
	case 4959376063826165943:
		return 11
	case 1237236456312824991:
		return 12
	case -488693769771308531:
		return 13
	case -8659679893868524044:
		return 14
	case 2110182719375245575:
		return 15
	case 6398007404756346381:
		return 16
	case -332090898078543319:
		return 17
	case -4415920049297795213:
		return 18
	case -4029073905697023887:
		return 19
	case -4255494004099806581:
		return 20
	case 1919988340934511316:
		return 21
	case 772598475950050299:
		return 22
	case 975043767033552552:
		return 23
	case -385596010576048993:
		return 24
	case 6316270346668237895:
		return 25
	case -551070009623711640:
		return 26
	case -9161949150489525396:
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
	case -199718707876719129:
		return 36
	case -2039166586246443706:
		return 37
	case 427629649267311777:
		return 38
	case 6477500296980050210:
		return 39
	case 4127240633310655048:
		return 40
	case 5792771328741774991:
		return 41
	case 5438819661942843222:
		return 42
	case 7345853097340709638:
		return 43
	case 5898328517505784191:
		return 44
	case 5217623514630858297:
		return 45
	case -5409927581812374498:
		return 46
	case -1185315367526857833:
		return 47
	case 3769868229123868474:
		return 48
	case 7379029982170503008:
		return 49
	case -568369772982215962:
		return 50
	}
	return -1
}

func QTabWidget_enable(h uintptr, idx int32) { pg_QTabWidget__set_override(h, idx) }

func (this *QTabWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTabWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTabWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTabWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(width int32) int32 { return int32(fn(int(width))) }) // slot "heightForWidth"
}

func (this *QTabWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTabWidget) OnTabInserted(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2734286415018532480, func(index int32) { fn(int(index)) }) // slot "tabInserted"
}

func (this *QTabWidget) OnTabRemoved(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -3307029722643392856, func(index int32) { fn(int(index)) }) // slot "tabRemoved"
}

func (this *QTabWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QTabWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QTabWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QTabWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QTabWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QTabWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QTabWidget) OnInitStyleOption(fn func(*QStyleOptionTabWidgetFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionTabWidgetFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QTabWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTabWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTabWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTabWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QTabWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QTabWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTabWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QTabWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QTabWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTabWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTabWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTabWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTabWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTabWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTabWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTabWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QTabWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTabWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTabWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTabWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTabWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTabWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTabWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QTabWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTabWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTabWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTabWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTabWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTabWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QTabWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTabWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTabWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTabWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTabWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTabWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTabWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTabWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTabWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTabWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
