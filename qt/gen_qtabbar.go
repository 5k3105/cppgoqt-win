package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTabBar_chain uintptr
	addr_QTabBar_new uintptr
	addr_QTabBar_new2 uintptr
	addr_QTabBar_tr uintptr
	addr_QTabBar_shape uintptr
	addr_QTabBar_setShape uintptr
	addr_QTabBar_addTab uintptr
	addr_QTabBar_addTab2 uintptr
	addr_QTabBar_insertTab uintptr
	addr_QTabBar_insertTab2 uintptr
	addr_QTabBar_removeTab uintptr
	addr_QTabBar_moveTab uintptr
	addr_QTabBar_isTabEnabled uintptr
	addr_QTabBar_setTabEnabled uintptr
	addr_QTabBar_isTabVisible uintptr
	addr_QTabBar_setTabVisible uintptr
	addr_QTabBar_tabText uintptr
	addr_QTabBar_setTabText uintptr
	addr_QTabBar_tabTextColor uintptr
	addr_QTabBar_setTabTextColor uintptr
	addr_QTabBar_tabIcon uintptr
	addr_QTabBar_setTabIcon uintptr
	addr_QTabBar_elideMode uintptr
	addr_QTabBar_setElideMode uintptr
	addr_QTabBar_setTabToolTip uintptr
	addr_QTabBar_tabToolTip uintptr
	addr_QTabBar_setTabWhatsThis uintptr
	addr_QTabBar_tabWhatsThis uintptr
	addr_QTabBar_setTabData uintptr
	addr_QTabBar_tabData uintptr
	addr_QTabBar_tabRect uintptr
	addr_QTabBar_tabAt uintptr
	addr_QTabBar_currentIndex uintptr
	addr_QTabBar_count uintptr
	addr_QTabBar_setDrawBase uintptr
	addr_QTabBar_drawBase uintptr
	addr_QTabBar_iconSize uintptr
	addr_QTabBar_setIconSize uintptr
	addr_QTabBar_usesScrollButtons uintptr
	addr_QTabBar_setUsesScrollButtons uintptr
	addr_QTabBar_tabsClosable uintptr
	addr_QTabBar_setTabsClosable uintptr
	addr_QTabBar_setTabButton uintptr
	addr_QTabBar_tabButton uintptr
	addr_QTabBar_selectionBehaviorOnRemove uintptr
	addr_QTabBar_setSelectionBehaviorOnRemove uintptr
	addr_QTabBar_expanding uintptr
	addr_QTabBar_setExpanding uintptr
	addr_QTabBar_isMovable uintptr
	addr_QTabBar_setMovable uintptr
	addr_QTabBar_documentMode uintptr
	addr_QTabBar_setDocumentMode uintptr
	addr_QTabBar_autoHide uintptr
	addr_QTabBar_setAutoHide uintptr
	addr_QTabBar_changeCurrentOnDrag uintptr
	addr_QTabBar_setChangeCurrentOnDrag uintptr
	addr_QTabBar_accessibleTabName uintptr
	addr_QTabBar_setAccessibleTabName uintptr
	addr_QTabBar_setCurrentIndex uintptr
	addr_QTabBar_currentChanged uintptr
	addr_QTabBar_tabCloseRequested uintptr
	addr_QTabBar_tabMoved uintptr
	addr_QTabBar_tabBarClicked uintptr
	addr_QTabBar_tabBarDoubleClicked uintptr
	addr_QTabBar_setTabOrder uintptr
	addr_QTabBar_mouseGrabber uintptr
	addr_QTabBar_keyboardGrabber uintptr
	addr_QTabBar_find uintptr
	addr_QTabBar_createWindowContainer uintptr
	addr_QTabBar_updateMicroFocus uintptr
	addr_QTabBar_create uintptr
	addr_QTabBar_destroy uintptr
	addr_QTabBar_focusNextChild uintptr
	addr_QTabBar_focusPreviousChild uintptr
	addr_QTabBar_connect uintptr
	addr_QTabBar_connect2 uintptr
	addr_QTabBar_sender uintptr
	addr_QTabBar_senderSignalIndex uintptr
	addr_QTabBar_receivers uintptr
	addr_QTabBar_isSignalConnected uintptr
	pg_QTabBar_devicePixelRatioFScale func() float64
	pg_QTabBar_encodeMetricF func(int32, float64) int32
	pg_QTabBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QTabBar_tr2 uintptr
	addr_QTabBar_tr3 uintptr
	addr_QTabBar_createWindowContainer2 uintptr
	addr_QTabBar_createWindowContainer3 uintptr
	addr_QTabBar_updateMicroFocus2 uintptr
	addr_QTabBar_create2 uintptr
	addr_QTabBar_create3 uintptr
	addr_QTabBar_create4 uintptr
	addr_QTabBar_destroy2 uintptr
	addr_QTabBar_destroy3 uintptr
	addr_QTabBar_connect3 uintptr
	addr_QTabBar_connect4 uintptr
	addr_QTabBar_metaObject_default uintptr
	addr_QTabBar_sizeHint_default uintptr
	addr_QTabBar_minimumSizeHint_default uintptr
	addr_QTabBar_tabSizeHint_default uintptr
	addr_QTabBar_minimumTabSizeHint_default uintptr
	addr_QTabBar_tabInserted_default uintptr
	addr_QTabBar_tabRemoved_default uintptr
	addr_QTabBar_tabLayoutChange_default uintptr
	addr_QTabBar_event_default uintptr
	addr_QTabBar_resizeEvent_default uintptr
	addr_QTabBar_showEvent_default uintptr
	addr_QTabBar_hideEvent_default uintptr
	addr_QTabBar_paintEvent_default uintptr
	addr_QTabBar_mousePressEvent_default uintptr
	addr_QTabBar_mouseMoveEvent_default uintptr
	addr_QTabBar_mouseReleaseEvent_default uintptr
	addr_QTabBar_mouseDoubleClickEvent_default uintptr
	addr_QTabBar_wheelEvent_default uintptr
	addr_QTabBar_keyPressEvent_default uintptr
	addr_QTabBar_changeEvent_default uintptr
	addr_QTabBar_timerEvent_default uintptr
	addr_QTabBar_initStyleOption_default uintptr
	addr_QTabBar_devType_default uintptr
	addr_QTabBar_setVisible_default uintptr
	addr_QTabBar_heightForWidth_default uintptr
	addr_QTabBar_hasHeightForWidth_default uintptr
	addr_QTabBar_paintEngine_default uintptr
	addr_QTabBar_keyReleaseEvent_default uintptr
	addr_QTabBar_focusInEvent_default uintptr
	addr_QTabBar_focusOutEvent_default uintptr
	addr_QTabBar_enterEvent_default uintptr
	addr_QTabBar_leaveEvent_default uintptr
	addr_QTabBar_moveEvent_default uintptr
	addr_QTabBar_closeEvent_default uintptr
	addr_QTabBar_contextMenuEvent_default uintptr
	addr_QTabBar_tabletEvent_default uintptr
	addr_QTabBar_actionEvent_default uintptr
	addr_QTabBar_dragEnterEvent_default uintptr
	addr_QTabBar_dragMoveEvent_default uintptr
	addr_QTabBar_dragLeaveEvent_default uintptr
	addr_QTabBar_dropEvent_default uintptr
	addr_QTabBar_metric_default uintptr
	addr_QTabBar_initPainter_default uintptr
	addr_QTabBar_redirected_default uintptr
	addr_QTabBar_sharedPainter_default uintptr
	addr_QTabBar_inputMethodEvent_default uintptr
	addr_QTabBar_inputMethodQuery_default uintptr
	addr_QTabBar_focusNextPrevChild_default uintptr
	addr_QTabBar_eventFilter_default uintptr
	addr_QTabBar_childEvent_default uintptr
	addr_QTabBar_customEvent_default uintptr
	addr_QTabBar_connectNotify_default uintptr
	addr_QTabBar_disconnectNotify_default uintptr
	pg_QTabBar_connect_currentChanged func(uintptr, int64)
	pg_QTabBar_connect_tabCloseRequested func(uintptr, int64)
	pg_QTabBar_connect_tabMoved func(uintptr, int64)
	pg_QTabBar_connect_tabBarClicked func(uintptr, int64)
	pg_QTabBar_connect_tabBarDoubleClicked func(uintptr, int64)
	pg_QTabBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QTabBar_connect_windowIconChanged func(uintptr, int64)
	pg_QTabBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QTabBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QTabBar_connect_destroyed func(uintptr, int64)
	pg_QTabBar_connect_destroyed2 func(uintptr, int64)
	pg_QTabBar_connect_objectNameChanged func(uintptr, int64)
	pg_QTabBar__disconnect func(int64)
	pg_QTabBar__set_override func(uintptr, int32)
	pg_QTabBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTabBar_chain, "QTabBar_chain")
	symaddr(&addr_QTabBar_new, "QTabBar_new")
	symaddr(&addr_QTabBar_new2, "QTabBar_new2")
	symaddr(&addr_QTabBar_tr, "QTabBar_tr")
	symaddr(&addr_QTabBar_shape, "QTabBar_shape")
	symaddr(&addr_QTabBar_setShape, "QTabBar_setShape")
	symaddr(&addr_QTabBar_addTab, "QTabBar_addTab")
	symaddr(&addr_QTabBar_addTab2, "QTabBar_addTab2")
	symaddr(&addr_QTabBar_insertTab, "QTabBar_insertTab")
	symaddr(&addr_QTabBar_insertTab2, "QTabBar_insertTab2")
	symaddr(&addr_QTabBar_removeTab, "QTabBar_removeTab")
	symaddr(&addr_QTabBar_moveTab, "QTabBar_moveTab")
	symaddr(&addr_QTabBar_isTabEnabled, "QTabBar_isTabEnabled")
	symaddr(&addr_QTabBar_setTabEnabled, "QTabBar_setTabEnabled")
	symaddr(&addr_QTabBar_isTabVisible, "QTabBar_isTabVisible")
	symaddr(&addr_QTabBar_setTabVisible, "QTabBar_setTabVisible")
	symaddr(&addr_QTabBar_tabText, "QTabBar_tabText")
	symaddr(&addr_QTabBar_setTabText, "QTabBar_setTabText")
	symaddr(&addr_QTabBar_tabTextColor, "QTabBar_tabTextColor")
	symaddr(&addr_QTabBar_setTabTextColor, "QTabBar_setTabTextColor")
	symaddr(&addr_QTabBar_tabIcon, "QTabBar_tabIcon")
	symaddr(&addr_QTabBar_setTabIcon, "QTabBar_setTabIcon")
	symaddr(&addr_QTabBar_elideMode, "QTabBar_elideMode")
	symaddr(&addr_QTabBar_setElideMode, "QTabBar_setElideMode")
	symaddr(&addr_QTabBar_setTabToolTip, "QTabBar_setTabToolTip")
	symaddr(&addr_QTabBar_tabToolTip, "QTabBar_tabToolTip")
	symaddr(&addr_QTabBar_setTabWhatsThis, "QTabBar_setTabWhatsThis")
	symaddr(&addr_QTabBar_tabWhatsThis, "QTabBar_tabWhatsThis")
	symaddr(&addr_QTabBar_setTabData, "QTabBar_setTabData")
	symaddr(&addr_QTabBar_tabData, "QTabBar_tabData")
	symaddr(&addr_QTabBar_tabRect, "QTabBar_tabRect")
	symaddr(&addr_QTabBar_tabAt, "QTabBar_tabAt")
	symaddr(&addr_QTabBar_currentIndex, "QTabBar_currentIndex")
	symaddr(&addr_QTabBar_count, "QTabBar_count")
	symaddr(&addr_QTabBar_setDrawBase, "QTabBar_setDrawBase")
	symaddr(&addr_QTabBar_drawBase, "QTabBar_drawBase")
	symaddr(&addr_QTabBar_iconSize, "QTabBar_iconSize")
	symaddr(&addr_QTabBar_setIconSize, "QTabBar_setIconSize")
	symaddr(&addr_QTabBar_usesScrollButtons, "QTabBar_usesScrollButtons")
	symaddr(&addr_QTabBar_setUsesScrollButtons, "QTabBar_setUsesScrollButtons")
	symaddr(&addr_QTabBar_tabsClosable, "QTabBar_tabsClosable")
	symaddr(&addr_QTabBar_setTabsClosable, "QTabBar_setTabsClosable")
	symaddr(&addr_QTabBar_setTabButton, "QTabBar_setTabButton")
	symaddr(&addr_QTabBar_tabButton, "QTabBar_tabButton")
	symaddr(&addr_QTabBar_selectionBehaviorOnRemove, "QTabBar_selectionBehaviorOnRemove")
	symaddr(&addr_QTabBar_setSelectionBehaviorOnRemove, "QTabBar_setSelectionBehaviorOnRemove")
	symaddr(&addr_QTabBar_expanding, "QTabBar_expanding")
	symaddr(&addr_QTabBar_setExpanding, "QTabBar_setExpanding")
	symaddr(&addr_QTabBar_isMovable, "QTabBar_isMovable")
	symaddr(&addr_QTabBar_setMovable, "QTabBar_setMovable")
	symaddr(&addr_QTabBar_documentMode, "QTabBar_documentMode")
	symaddr(&addr_QTabBar_setDocumentMode, "QTabBar_setDocumentMode")
	symaddr(&addr_QTabBar_autoHide, "QTabBar_autoHide")
	symaddr(&addr_QTabBar_setAutoHide, "QTabBar_setAutoHide")
	symaddr(&addr_QTabBar_changeCurrentOnDrag, "QTabBar_changeCurrentOnDrag")
	symaddr(&addr_QTabBar_setChangeCurrentOnDrag, "QTabBar_setChangeCurrentOnDrag")
	symaddr(&addr_QTabBar_accessibleTabName, "QTabBar_accessibleTabName")
	symaddr(&addr_QTabBar_setAccessibleTabName, "QTabBar_setAccessibleTabName")
	symaddr(&addr_QTabBar_setCurrentIndex, "QTabBar_setCurrentIndex")
	symaddr(&addr_QTabBar_currentChanged, "QTabBar_currentChanged")
	symaddr(&addr_QTabBar_tabCloseRequested, "QTabBar_tabCloseRequested")
	symaddr(&addr_QTabBar_tabMoved, "QTabBar_tabMoved")
	symaddr(&addr_QTabBar_tabBarClicked, "QTabBar_tabBarClicked")
	symaddr(&addr_QTabBar_tabBarDoubleClicked, "QTabBar_tabBarDoubleClicked")
	symaddr(&addr_QTabBar_setTabOrder, "QTabBar_setTabOrder")
	symaddr(&addr_QTabBar_mouseGrabber, "QTabBar_mouseGrabber")
	symaddr(&addr_QTabBar_keyboardGrabber, "QTabBar_keyboardGrabber")
	symaddr(&addr_QTabBar_find, "QTabBar_find")
	symaddr(&addr_QTabBar_createWindowContainer, "QTabBar_createWindowContainer")
	symaddr(&addr_QTabBar_updateMicroFocus, "QTabBar_updateMicroFocus")
	symaddr(&addr_QTabBar_create, "QTabBar_create")
	symaddr(&addr_QTabBar_destroy, "QTabBar_destroy")
	symaddr(&addr_QTabBar_focusNextChild, "QTabBar_focusNextChild")
	symaddr(&addr_QTabBar_focusPreviousChild, "QTabBar_focusPreviousChild")
	symaddr(&addr_QTabBar_connect, "QTabBar_connect")
	symaddr(&addr_QTabBar_connect2, "QTabBar_connect2")
	symaddr(&addr_QTabBar_sender, "QTabBar_sender")
	symaddr(&addr_QTabBar_senderSignalIndex, "QTabBar_senderSignalIndex")
	symaddr(&addr_QTabBar_receivers, "QTabBar_receivers")
	symaddr(&addr_QTabBar_isSignalConnected, "QTabBar_isSignalConnected")
	reg(&pg_QTabBar_devicePixelRatioFScale, "QTabBar_devicePixelRatioFScale")
	reg(&pg_QTabBar_encodeMetricF, "QTabBar_encodeMetricF")
	reg(&pg_QTabBar_getDecodedMetricF, "QTabBar_getDecodedMetricF")
	symaddr(&addr_QTabBar_tr2, "QTabBar_tr2")
	symaddr(&addr_QTabBar_tr3, "QTabBar_tr3")
	symaddr(&addr_QTabBar_createWindowContainer2, "QTabBar_createWindowContainer2")
	symaddr(&addr_QTabBar_createWindowContainer3, "QTabBar_createWindowContainer3")
	symaddr(&addr_QTabBar_updateMicroFocus2, "QTabBar_updateMicroFocus2")
	symaddr(&addr_QTabBar_create2, "QTabBar_create2")
	symaddr(&addr_QTabBar_create3, "QTabBar_create3")
	symaddr(&addr_QTabBar_create4, "QTabBar_create4")
	symaddr(&addr_QTabBar_destroy2, "QTabBar_destroy2")
	symaddr(&addr_QTabBar_destroy3, "QTabBar_destroy3")
	symaddr(&addr_QTabBar_connect3, "QTabBar_connect3")
	symaddr(&addr_QTabBar_connect4, "QTabBar_connect4")
	symaddr(&addr_QTabBar_metaObject_default, "QTabBar_metaObject_default")
	symaddr(&addr_QTabBar_sizeHint_default, "QTabBar_sizeHint_default")
	symaddr(&addr_QTabBar_minimumSizeHint_default, "QTabBar_minimumSizeHint_default")
	symaddr(&addr_QTabBar_tabSizeHint_default, "QTabBar_tabSizeHint_default")
	symaddr(&addr_QTabBar_minimumTabSizeHint_default, "QTabBar_minimumTabSizeHint_default")
	symaddr(&addr_QTabBar_tabInserted_default, "QTabBar_tabInserted_default")
	symaddr(&addr_QTabBar_tabRemoved_default, "QTabBar_tabRemoved_default")
	symaddr(&addr_QTabBar_tabLayoutChange_default, "QTabBar_tabLayoutChange_default")
	symaddr(&addr_QTabBar_event_default, "QTabBar_event_default")
	symaddr(&addr_QTabBar_resizeEvent_default, "QTabBar_resizeEvent_default")
	symaddr(&addr_QTabBar_showEvent_default, "QTabBar_showEvent_default")
	symaddr(&addr_QTabBar_hideEvent_default, "QTabBar_hideEvent_default")
	symaddr(&addr_QTabBar_paintEvent_default, "QTabBar_paintEvent_default")
	symaddr(&addr_QTabBar_mousePressEvent_default, "QTabBar_mousePressEvent_default")
	symaddr(&addr_QTabBar_mouseMoveEvent_default, "QTabBar_mouseMoveEvent_default")
	symaddr(&addr_QTabBar_mouseReleaseEvent_default, "QTabBar_mouseReleaseEvent_default")
	symaddr(&addr_QTabBar_mouseDoubleClickEvent_default, "QTabBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QTabBar_wheelEvent_default, "QTabBar_wheelEvent_default")
	symaddr(&addr_QTabBar_keyPressEvent_default, "QTabBar_keyPressEvent_default")
	symaddr(&addr_QTabBar_changeEvent_default, "QTabBar_changeEvent_default")
	symaddr(&addr_QTabBar_timerEvent_default, "QTabBar_timerEvent_default")
	symaddr(&addr_QTabBar_initStyleOption_default, "QTabBar_initStyleOption_default")
	symaddr(&addr_QTabBar_devType_default, "QTabBar_devType_default")
	symaddr(&addr_QTabBar_setVisible_default, "QTabBar_setVisible_default")
	symaddr(&addr_QTabBar_heightForWidth_default, "QTabBar_heightForWidth_default")
	symaddr(&addr_QTabBar_hasHeightForWidth_default, "QTabBar_hasHeightForWidth_default")
	symaddr(&addr_QTabBar_paintEngine_default, "QTabBar_paintEngine_default")
	symaddr(&addr_QTabBar_keyReleaseEvent_default, "QTabBar_keyReleaseEvent_default")
	symaddr(&addr_QTabBar_focusInEvent_default, "QTabBar_focusInEvent_default")
	symaddr(&addr_QTabBar_focusOutEvent_default, "QTabBar_focusOutEvent_default")
	symaddr(&addr_QTabBar_enterEvent_default, "QTabBar_enterEvent_default")
	symaddr(&addr_QTabBar_leaveEvent_default, "QTabBar_leaveEvent_default")
	symaddr(&addr_QTabBar_moveEvent_default, "QTabBar_moveEvent_default")
	symaddr(&addr_QTabBar_closeEvent_default, "QTabBar_closeEvent_default")
	symaddr(&addr_QTabBar_contextMenuEvent_default, "QTabBar_contextMenuEvent_default")
	symaddr(&addr_QTabBar_tabletEvent_default, "QTabBar_tabletEvent_default")
	symaddr(&addr_QTabBar_actionEvent_default, "QTabBar_actionEvent_default")
	symaddr(&addr_QTabBar_dragEnterEvent_default, "QTabBar_dragEnterEvent_default")
	symaddr(&addr_QTabBar_dragMoveEvent_default, "QTabBar_dragMoveEvent_default")
	symaddr(&addr_QTabBar_dragLeaveEvent_default, "QTabBar_dragLeaveEvent_default")
	symaddr(&addr_QTabBar_dropEvent_default, "QTabBar_dropEvent_default")
	symaddr(&addr_QTabBar_metric_default, "QTabBar_metric_default")
	symaddr(&addr_QTabBar_initPainter_default, "QTabBar_initPainter_default")
	symaddr(&addr_QTabBar_redirected_default, "QTabBar_redirected_default")
	symaddr(&addr_QTabBar_sharedPainter_default, "QTabBar_sharedPainter_default")
	symaddr(&addr_QTabBar_inputMethodEvent_default, "QTabBar_inputMethodEvent_default")
	symaddr(&addr_QTabBar_inputMethodQuery_default, "QTabBar_inputMethodQuery_default")
	symaddr(&addr_QTabBar_focusNextPrevChild_default, "QTabBar_focusNextPrevChild_default")
	symaddr(&addr_QTabBar_eventFilter_default, "QTabBar_eventFilter_default")
	symaddr(&addr_QTabBar_childEvent_default, "QTabBar_childEvent_default")
	symaddr(&addr_QTabBar_customEvent_default, "QTabBar_customEvent_default")
	symaddr(&addr_QTabBar_connectNotify_default, "QTabBar_connectNotify_default")
	symaddr(&addr_QTabBar_disconnectNotify_default, "QTabBar_disconnectNotify_default")
	reg(&pg_QTabBar_connect_currentChanged, "QTabBar_connect_currentChanged")
	reg(&pg_QTabBar_connect_tabCloseRequested, "QTabBar_connect_tabCloseRequested")
	reg(&pg_QTabBar_connect_tabMoved, "QTabBar_connect_tabMoved")
	reg(&pg_QTabBar_connect_tabBarClicked, "QTabBar_connect_tabBarClicked")
	reg(&pg_QTabBar_connect_tabBarDoubleClicked, "QTabBar_connect_tabBarDoubleClicked")
	reg(&pg_QTabBar_connect_windowTitleChanged, "QTabBar_connect_windowTitleChanged")
	reg(&pg_QTabBar_connect_windowIconChanged, "QTabBar_connect_windowIconChanged")
	reg(&pg_QTabBar_connect_windowIconTextChanged, "QTabBar_connect_windowIconTextChanged")
	reg(&pg_QTabBar_connect_customContextMenuRequested, "QTabBar_connect_customContextMenuRequested")
	reg(&pg_QTabBar_connect_destroyed, "QTabBar_connect_destroyed")
	reg(&pg_QTabBar_connect_destroyed2, "QTabBar_connect_destroyed2")
	reg(&pg_QTabBar_connect_objectNameChanged, "QTabBar_connect_objectNameChanged")
	reg(&pg_QTabBar__disconnect, "QTabBar__disconnect")
	reg(&pg_QTabBar__set_override, "QTabBar__set_override")
	reg(&pg_QTabBar_delete, "QTabBar_delete")
}

type QTabBar struct {
	h uintptr
	*QWidget
}

func bindQTabBar(h uintptr) *QTabBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTabBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTabBar{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTabBar(h uintptr) *QTabBar { return bindQTabBar(h) }

func (this *QTabBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTabBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTabBar(h unsafe.Pointer) *QTabBar { return bindQTabBar(uintptr(h)) }
func UnsafeNewQTabBarOwned(h unsafe.Pointer) *QTabBar { return newownedQTabBar(uintptr(h)) }

func (this *QTabBar) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQTabBar(parent *QWidget) *QTabBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTabBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTabBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTabBar_virtIdx, QTabBar_enable)
	return __o
}

func NewQTabBar2() *QTabBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTabBar_new2, uintptr(__goid))
	__o := newownedQTabBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTabBar_virtIdx, QTabBar_enable)
	return __o
}

func QTabBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTabBar) Shape() QTabBar__Shape {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_shape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabBar__Shape(__r)
}

func (this *QTabBar) SetShape(shape QTabBar__Shape) {
	purego.SyscallN(addr_QTabBar_setShape, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QTabBar) AddTab(text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabBar_addTab, this.h, text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QTabBar) AddTab2(icon *QIcon, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabBar_addTab2, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QTabBar) InsertTab(index int, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabBar_insertTab, this.h, uintptr(index), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QTabBar) InsertTab2(index int, icon *QIcon, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTabBar_insertTab2, this.h, uintptr(index), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QTabBar) RemoveTab(index int) {
	purego.SyscallN(addr_QTabBar_removeTab, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) MoveTab(from int, to int) {
	purego.SyscallN(addr_QTabBar_moveTab, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QTabBar) IsTabEnabled(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_isTabEnabled, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetTabEnabled(index int, enabled bool) {
	purego.SyscallN(addr_QTabBar_setTabEnabled, this.h, uintptr(index), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) IsTabVisible(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_isTabVisible, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetTabVisible(index int, visible bool) {
	purego.SyscallN(addr_QTabBar_setTabVisible, this.h, uintptr(index), uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabText(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabText, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabBar) SetTabText(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTabBar_setTabText, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTabBar) TabTextColor(index int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabTextColor, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QTabBar) SetTabTextColor(index int, color *QColor) {
	purego.SyscallN(addr_QTabBar_setTabTextColor, this.h, uintptr(index), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QTabBar) TabIcon(index int) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabIcon, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QTabBar) SetTabIcon(index int, icon *QIcon) {
	purego.SyscallN(addr_QTabBar_setTabIcon, this.h, uintptr(index), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QTabBar) ElideMode() TextElideMode {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_elideMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextElideMode(__r)
}

func (this *QTabBar) SetElideMode(mode TextElideMode) {
	purego.SyscallN(addr_QTabBar_setElideMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QTabBar) SetTabToolTip(index int, tip string) {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	purego.SyscallN(addr_QTabBar_setTabToolTip, this.h, uintptr(index), tip_p, uintptr(len(tip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tip)
}

func (this *QTabBar) TabToolTip(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabToolTip, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabBar) SetTabWhatsThis(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTabBar_setTabWhatsThis, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTabBar) TabWhatsThis(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabWhatsThis, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabBar) SetTabData(index int, data *QVariant) {
	purego.SyscallN(addr_QTabBar_setTabData, this.h, uintptr(index), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QTabBar) TabData(index int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabData, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTabBar) TabRect(index int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabRect, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QTabBar) TabAt(pos *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_tabAt, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return int(__r)
}

func (this *QTabBar) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) SetDrawBase(drawTheBase bool) {
	purego.SyscallN(addr_QTabBar_setDrawBase, this.h, uintptr(b2i(drawTheBase)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) DrawBase() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_drawBase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabBar_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabBar) SetIconSize(size *QSize) {
	purego.SyscallN(addr_QTabBar_setIconSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QTabBar) UsesScrollButtons() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_usesScrollButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetUsesScrollButtons(useButtons bool) {
	purego.SyscallN(addr_QTabBar_setUsesScrollButtons, this.h, uintptr(b2i(useButtons)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabsClosable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_tabsClosable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetTabsClosable(closable bool) {
	purego.SyscallN(addr_QTabBar_setTabsClosable, this.h, uintptr(b2i(closable)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) SetTabButton(index int, position QTabBar__ButtonPosition, widget *QWidget) {
	purego.SyscallN(addr_QTabBar_setTabButton, this.h, uintptr(index), uintptr(position), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QTabBar) TabButton(index int, position QTabBar__ButtonPosition) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabButton, this.h, uintptr(index), uintptr(position))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QTabBar) SelectionBehaviorOnRemove() QTabBar__SelectionBehavior {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_selectionBehaviorOnRemove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabBar__SelectionBehavior(__r)
}

func (this *QTabBar) SetSelectionBehaviorOnRemove(behavior QTabBar__SelectionBehavior) {
	purego.SyscallN(addr_QTabBar_setSelectionBehaviorOnRemove, this.h, uintptr(behavior))
	runtime.KeepAlive(this)
}

func (this *QTabBar) Expanding() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_expanding, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetExpanding(enabled bool) {
	purego.SyscallN(addr_QTabBar_setExpanding, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) IsMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_isMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetMovable(movable bool) {
	purego.SyscallN(addr_QTabBar_setMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_documentMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetDocumentMode(set bool) {
	purego.SyscallN(addr_QTabBar_setDocumentMode, this.h, uintptr(b2i(set)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) AutoHide() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_autoHide, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetAutoHide(hide bool) {
	purego.SyscallN(addr_QTabBar_setAutoHide, this.h, uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) ChangeCurrentOnDrag() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_changeCurrentOnDrag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) SetChangeCurrentOnDrag(change bool) {
	purego.SyscallN(addr_QTabBar_setChangeCurrentOnDrag, this.h, uintptr(b2i(change)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) AccessibleTabName(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_accessibleTabName, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTabBar) SetAccessibleTabName(index int, name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTabBar_setAccessibleTabName, this.h, uintptr(index), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTabBar) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QTabBar_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) CurrentChanged(index int) {
	purego.SyscallN(addr_QTabBar_currentChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabCloseRequested(index int) {
	purego.SyscallN(addr_QTabBar_tabCloseRequested, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabMoved(from int, to int) {
	purego.SyscallN(addr_QTabBar_tabMoved, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabBarClicked(index int) {
	purego.SyscallN(addr_QTabBar_tabBarClicked, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabBarDoubleClicked(index int) {
	purego.SyscallN(addr_QTabBar_tabBarDoubleClicked, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func QTabBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QTabBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QTabBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_mouseGrabber)
	return bindQWidget(__r)
}

func QTabBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QTabBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QTabBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QTabBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QTabBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabBar) Create() {
	purego.SyscallN(addr_QTabBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabBar) Destroy() {
	purego.SyscallN(addr_QTabBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTabBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTabBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTabBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTabBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTabBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTabBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTabBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTabBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTabBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTabBar_DevicePixelRatioFScale() float64 {
	__r := pg_QTabBar_devicePixelRatioFScale()
	return __r
}

func QTabBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QTabBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QTabBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QTabBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QTabBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTabBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTabBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTabBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QTabBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QTabBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QTabBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QTabBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QTabBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QTabBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QTabBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QTabBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QTabBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QTabBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QTabBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QTabBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTabBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTabBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTabBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTabBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTabBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTabBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabBar) TabSizeHintDefault(index int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabBar_tabSizeHint_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabBar) MinimumTabSizeHintDefault(index int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QTabBar_minimumTabSizeHint_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTabBar) TabInsertedDefault(index int) {
	purego.SyscallN(addr_QTabBar_tabInserted_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabRemovedDefault(index int) {
	purego.SyscallN(addr_QTabBar_tabRemoved_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QTabBar) TabLayoutChangeDefault() {
	purego.SyscallN(addr_QTabBar_tabLayoutChange_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTabBar) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTabBar) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QTabBar_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QTabBar_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) HideEventDefault(arg1 *QHideEvent) {
	purego.SyscallN(addr_QTabBar_hideEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QTabBar_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QTabBar_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QTabBar_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QTabBar_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QTabBar_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QTabBar_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QTabBar_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QTabBar_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTabBar_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) InitStyleOptionDefault(option *QStyleOptionTab, tabIndex int) {
	purego.SyscallN(addr_QTabBar_initStyleOption_default, this.h, uintptr(option.UnsafePointer()), uintptr(tabIndex))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QTabBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QTabBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QTabBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QTabBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QTabBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QTabBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTabBar_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QTabBar_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QTabBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTabBar_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QTabBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QTabBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QTabBar_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QTabBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QTabBar_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QTabBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QTabBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QTabBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QTabBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTabBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QTabBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QTabBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QTabBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QTabBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QTabBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QTabBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QTabBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTabBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTabBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTabBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTabBar) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTabBar_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTabBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTabBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTabBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTabBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTabBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTabBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTabBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTabBar) OnCurrentChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabBar_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnTabCloseRequested(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabBar_connect_tabCloseRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnTabMoved(fn func(int, int)) int64 {
	id := registerCallback(func(from int32, to int32) { fn(int(from), int(to)) })
	pg_QTabBar_connect_tabMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnTabBarClicked(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabBar_connect_tabBarClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnTabBarDoubleClicked(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QTabBar_connect_tabBarDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QTabBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QTabBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QTabBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QTabBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTabBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTabBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTabBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTabBar) Disconnect(handle int64) {
	pg_QTabBar__disconnect(handle)
	unregisterCallback(handle)
}

func QTabBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case -5008018812442754120:
		return 3
	case -5553900436438809362:
		return 4
	case -2734286415018532480:
		return 5
	case -3307029722643392856:
		return 6
	case 9115710866053612342:
		return 7
	case 1237236456312824991:
		return 8
	case -4302917285655014661:
		return 9
	case -6825369174751916:
		return 10
	case -199718707876719129:
		return 11
	case 5290442163167458465:
		return 12
	case -332090898078543319:
		return 13
	case -4255494004099806581:
		return 14
	case -4415920049297795213:
		return 15
	case -4029073905697023887:
		return 16
	case 1919988340934511316:
		return 17
	case 1023791977679793185:
		return 18
	case 4959376063826165943:
		return 19
	case -5409927581812374498:
		return 20
	case -488693769771308531:
		return 21
	case -8659679893868524044:
		return 22
	case 2110182719375245575:
		return 23
	case 4311534060495737339:
		return 24
	case -8266762397995423915:
		return 25
	case 6398007404756346381:
		return 26
	case 772598475950050299:
		return 27
	case 975043767033552552:
		return 28
	case -385596010576048993:
		return 29
	case 6316270346668237895:
		return 30
	case -551070009623711640:
		return 31
	case -9161949150489525396:
		return 32
	case -1989937050019464783:
		return 33
	case -6698043222227006015:
		return 34
	case -212068866860701141:
		return 35
	case 1697444140088555581:
		return 36
	case 206122715919917589:
		return 37
	case -1781545137352745978:
		return 38
	case -3139904158373725006:
		return 39
	case 9107058682754591716:
		return 40
	case -2039166586246443706:
		return 41
	case 427629649267311777:
		return 42
	case 6477500296980050210:
		return 43
	case 4127240633310655048:
		return 44
	case 5792771328741774991:
		return 45
	case 5438819661942843222:
		return 46
	case 7345853097340709638:
		return 47
	case 5898328517505784191:
		return 48
	case 5217623514630858297:
		return 49
	case -1185315367526857833:
		return 50
	case 3769868229123868474:
		return 51
	case 7379029982170503008:
		return 52
	case -568369772982215962:
		return 53
	}
	return -1
}

func QTabBar_enable(h uintptr, idx int32) { pg_QTabBar__set_override(h, idx) }

func (this *QTabBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTabBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QTabBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QTabBar) OnTabSizeHint(fn func(int) *QSize) {
	registerVirt(this.UnsafeGoID(), -5008018812442754120, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "tabSizeHint"
}

func (this *QTabBar) OnMinimumTabSizeHint(fn func(int) *QSize) {
	registerVirt(this.UnsafeGoID(), -5553900436438809362, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "minimumTabSizeHint"
}

func (this *QTabBar) OnTabInserted(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2734286415018532480, func(index int32) { fn(int(index)) }) // slot "tabInserted"
}

func (this *QTabBar) OnTabRemoved(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -3307029722643392856, func(index int32) { fn(int(index)) }) // slot "tabRemoved"
}

func (this *QTabBar) OnTabLayoutChange(fn func()) {
	registerVirt(this.UnsafeGoID(), 9115710866053612342, func() { fn() }) // slot "tabLayoutChange"
}

func (this *QTabBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QTabBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QTabBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QTabBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(arg1 uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(arg1))) }) // slot "hideEvent"
}

func (this *QTabBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QTabBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QTabBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QTabBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QTabBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QTabBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QTabBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QTabBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QTabBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTabBar) OnInitStyleOption(fn func(*QStyleOptionTab, int)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr, tabIndex int32) { fn(UnsafeNewQStyleOptionTab(unsafe.Pointer(option)), int(tabIndex)) }) // slot "initStyleOption"
}

func (this *QTabBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QTabBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QTabBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QTabBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QTabBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QTabBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QTabBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QTabBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QTabBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QTabBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QTabBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QTabBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QTabBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QTabBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QTabBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QTabBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QTabBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QTabBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QTabBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QTabBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QTabBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QTabBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QTabBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QTabBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QTabBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QTabBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QTabBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QTabBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTabBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTabBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTabBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTabBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTabBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTabBar_delete(this.h)
	}
	this.UnsafeRelease()
}
