package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMenu_chain uintptr
	addr_QMenu_new uintptr
	addr_QMenu_new2 uintptr
	addr_QMenu_new3 uintptr
	addr_QMenu_new4 uintptr
	addr_QMenu_tr uintptr
	addr_QMenu_addAction uintptr
	addr_QMenu_addAction2 uintptr
	addr_QMenu_addMenu uintptr
	addr_QMenu_addMenu2 uintptr
	addr_QMenu_addMenu3 uintptr
	addr_QMenu_addSeparator uintptr
	addr_QMenu_addSection uintptr
	addr_QMenu_addSection2 uintptr
	addr_QMenu_insertMenu uintptr
	addr_QMenu_insertSeparator uintptr
	addr_QMenu_insertSection uintptr
	addr_QMenu_insertSection2 uintptr
	addr_QMenu_isEmpty uintptr
	addr_QMenu_clear uintptr
	addr_QMenu_setTearOffEnabled uintptr
	addr_QMenu_isTearOffEnabled uintptr
	addr_QMenu_isTearOffMenuVisible uintptr
	addr_QMenu_showTearOffMenu uintptr
	addr_QMenu_showTearOffMenu2 uintptr
	addr_QMenu_hideTearOffMenu uintptr
	addr_QMenu_setDefaultAction uintptr
	addr_QMenu_defaultAction uintptr
	addr_QMenu_setActiveAction uintptr
	addr_QMenu_activeAction uintptr
	addr_QMenu_popup uintptr
	addr_QMenu_exec uintptr
	addr_QMenu_exec2 uintptr
	addr_QMenu_exec3 uintptr
	addr_QMenu_actionGeometry uintptr
	addr_QMenu_actionAt uintptr
	addr_QMenu_menuAction uintptr
	addr_QMenu_menuInAction uintptr
	addr_QMenu_title uintptr
	addr_QMenu_setTitle uintptr
	addr_QMenu_icon uintptr
	addr_QMenu_setIcon uintptr
	addr_QMenu_setNoReplayFor uintptr
	addr_QMenu_separatorsCollapsible uintptr
	addr_QMenu_setSeparatorsCollapsible uintptr
	addr_QMenu_toolTipsVisible uintptr
	addr_QMenu_setToolTipsVisible uintptr
	addr_QMenu_aboutToShow uintptr
	addr_QMenu_aboutToHide uintptr
	addr_QMenu_triggered uintptr
	addr_QMenu_hovered uintptr
	addr_QMenu_columnCount uintptr
	addr_QMenu_setTabOrder uintptr
	addr_QMenu_mouseGrabber uintptr
	addr_QMenu_keyboardGrabber uintptr
	addr_QMenu_find uintptr
	addr_QMenu_createWindowContainer uintptr
	addr_QMenu_updateMicroFocus uintptr
	addr_QMenu_create uintptr
	addr_QMenu_destroy uintptr
	addr_QMenu_focusNextChild uintptr
	addr_QMenu_focusPreviousChild uintptr
	addr_QMenu_connect uintptr
	addr_QMenu_connect2 uintptr
	addr_QMenu_sender uintptr
	addr_QMenu_senderSignalIndex uintptr
	addr_QMenu_receivers uintptr
	addr_QMenu_isSignalConnected uintptr
	pg_QMenu_devicePixelRatioFScale func() float64
	pg_QMenu_encodeMetricF func(int32, float64) int32
	pg_QMenu_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMenu_tr2 uintptr
	addr_QMenu_tr3 uintptr
	addr_QMenu_popup2 uintptr
	addr_QMenu_exec4 uintptr
	addr_QMenu_exec5 uintptr
	addr_QMenu_exec6 uintptr
	addr_QMenu_createWindowContainer2 uintptr
	addr_QMenu_createWindowContainer3 uintptr
	addr_QMenu_updateMicroFocus2 uintptr
	addr_QMenu_create2 uintptr
	addr_QMenu_create3 uintptr
	addr_QMenu_create4 uintptr
	addr_QMenu_destroy2 uintptr
	addr_QMenu_destroy3 uintptr
	addr_QMenu_connect3 uintptr
	addr_QMenu_connect4 uintptr
	addr_QMenu_metaObject_default uintptr
	addr_QMenu_sizeHint_default uintptr
	addr_QMenu_changeEvent_default uintptr
	addr_QMenu_keyPressEvent_default uintptr
	addr_QMenu_mouseReleaseEvent_default uintptr
	addr_QMenu_mousePressEvent_default uintptr
	addr_QMenu_mouseMoveEvent_default uintptr
	addr_QMenu_wheelEvent_default uintptr
	addr_QMenu_enterEvent_default uintptr
	addr_QMenu_leaveEvent_default uintptr
	addr_QMenu_hideEvent_default uintptr
	addr_QMenu_paintEvent_default uintptr
	addr_QMenu_actionEvent_default uintptr
	addr_QMenu_timerEvent_default uintptr
	addr_QMenu_event_default uintptr
	addr_QMenu_focusNextPrevChild_default uintptr
	addr_QMenu_initStyleOption_default uintptr
	addr_QMenu_devType_default uintptr
	addr_QMenu_setVisible_default uintptr
	addr_QMenu_minimumSizeHint_default uintptr
	addr_QMenu_heightForWidth_default uintptr
	addr_QMenu_hasHeightForWidth_default uintptr
	addr_QMenu_paintEngine_default uintptr
	addr_QMenu_mouseDoubleClickEvent_default uintptr
	addr_QMenu_keyReleaseEvent_default uintptr
	addr_QMenu_focusInEvent_default uintptr
	addr_QMenu_focusOutEvent_default uintptr
	addr_QMenu_moveEvent_default uintptr
	addr_QMenu_resizeEvent_default uintptr
	addr_QMenu_closeEvent_default uintptr
	addr_QMenu_contextMenuEvent_default uintptr
	addr_QMenu_tabletEvent_default uintptr
	addr_QMenu_dragEnterEvent_default uintptr
	addr_QMenu_dragMoveEvent_default uintptr
	addr_QMenu_dragLeaveEvent_default uintptr
	addr_QMenu_dropEvent_default uintptr
	addr_QMenu_showEvent_default uintptr
	addr_QMenu_metric_default uintptr
	addr_QMenu_initPainter_default uintptr
	addr_QMenu_redirected_default uintptr
	addr_QMenu_sharedPainter_default uintptr
	addr_QMenu_inputMethodEvent_default uintptr
	addr_QMenu_inputMethodQuery_default uintptr
	addr_QMenu_eventFilter_default uintptr
	addr_QMenu_childEvent_default uintptr
	addr_QMenu_customEvent_default uintptr
	addr_QMenu_connectNotify_default uintptr
	addr_QMenu_disconnectNotify_default uintptr
	pg_QMenu_connect_aboutToShow func(uintptr, int64)
	pg_QMenu_connect_aboutToHide func(uintptr, int64)
	pg_QMenu_connect_triggered func(uintptr, int64)
	pg_QMenu_connect_hovered func(uintptr, int64)
	pg_QMenu_connect_windowTitleChanged func(uintptr, int64)
	pg_QMenu_connect_windowIconChanged func(uintptr, int64)
	pg_QMenu_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMenu_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMenu_connect_destroyed func(uintptr, int64)
	pg_QMenu_connect_destroyed2 func(uintptr, int64)
	pg_QMenu_connect_objectNameChanged func(uintptr, int64)
	pg_QMenu__disconnect func(int64)
	pg_QMenu__set_override func(uintptr, int32)
	pg_QMenu_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMenu_chain, "QMenu_chain")
	symaddr(&addr_QMenu_new, "QMenu_new")
	symaddr(&addr_QMenu_new2, "QMenu_new2")
	symaddr(&addr_QMenu_new3, "QMenu_new3")
	symaddr(&addr_QMenu_new4, "QMenu_new4")
	symaddr(&addr_QMenu_tr, "QMenu_tr")
	symaddr(&addr_QMenu_addAction, "QMenu_addAction")
	symaddr(&addr_QMenu_addAction2, "QMenu_addAction2")
	symaddr(&addr_QMenu_addMenu, "QMenu_addMenu")
	symaddr(&addr_QMenu_addMenu2, "QMenu_addMenu2")
	symaddr(&addr_QMenu_addMenu3, "QMenu_addMenu3")
	symaddr(&addr_QMenu_addSeparator, "QMenu_addSeparator")
	symaddr(&addr_QMenu_addSection, "QMenu_addSection")
	symaddr(&addr_QMenu_addSection2, "QMenu_addSection2")
	symaddr(&addr_QMenu_insertMenu, "QMenu_insertMenu")
	symaddr(&addr_QMenu_insertSeparator, "QMenu_insertSeparator")
	symaddr(&addr_QMenu_insertSection, "QMenu_insertSection")
	symaddr(&addr_QMenu_insertSection2, "QMenu_insertSection2")
	symaddr(&addr_QMenu_isEmpty, "QMenu_isEmpty")
	symaddr(&addr_QMenu_clear, "QMenu_clear")
	symaddr(&addr_QMenu_setTearOffEnabled, "QMenu_setTearOffEnabled")
	symaddr(&addr_QMenu_isTearOffEnabled, "QMenu_isTearOffEnabled")
	symaddr(&addr_QMenu_isTearOffMenuVisible, "QMenu_isTearOffMenuVisible")
	symaddr(&addr_QMenu_showTearOffMenu, "QMenu_showTearOffMenu")
	symaddr(&addr_QMenu_showTearOffMenu2, "QMenu_showTearOffMenu2")
	symaddr(&addr_QMenu_hideTearOffMenu, "QMenu_hideTearOffMenu")
	symaddr(&addr_QMenu_setDefaultAction, "QMenu_setDefaultAction")
	symaddr(&addr_QMenu_defaultAction, "QMenu_defaultAction")
	symaddr(&addr_QMenu_setActiveAction, "QMenu_setActiveAction")
	symaddr(&addr_QMenu_activeAction, "QMenu_activeAction")
	symaddr(&addr_QMenu_popup, "QMenu_popup")
	symaddr(&addr_QMenu_exec, "QMenu_exec")
	symaddr(&addr_QMenu_exec2, "QMenu_exec2")
	symaddr(&addr_QMenu_exec3, "QMenu_exec3")
	symaddr(&addr_QMenu_actionGeometry, "QMenu_actionGeometry")
	symaddr(&addr_QMenu_actionAt, "QMenu_actionAt")
	symaddr(&addr_QMenu_menuAction, "QMenu_menuAction")
	symaddr(&addr_QMenu_menuInAction, "QMenu_menuInAction")
	symaddr(&addr_QMenu_title, "QMenu_title")
	symaddr(&addr_QMenu_setTitle, "QMenu_setTitle")
	symaddr(&addr_QMenu_icon, "QMenu_icon")
	symaddr(&addr_QMenu_setIcon, "QMenu_setIcon")
	symaddr(&addr_QMenu_setNoReplayFor, "QMenu_setNoReplayFor")
	symaddr(&addr_QMenu_separatorsCollapsible, "QMenu_separatorsCollapsible")
	symaddr(&addr_QMenu_setSeparatorsCollapsible, "QMenu_setSeparatorsCollapsible")
	symaddr(&addr_QMenu_toolTipsVisible, "QMenu_toolTipsVisible")
	symaddr(&addr_QMenu_setToolTipsVisible, "QMenu_setToolTipsVisible")
	symaddr(&addr_QMenu_aboutToShow, "QMenu_aboutToShow")
	symaddr(&addr_QMenu_aboutToHide, "QMenu_aboutToHide")
	symaddr(&addr_QMenu_triggered, "QMenu_triggered")
	symaddr(&addr_QMenu_hovered, "QMenu_hovered")
	symaddr(&addr_QMenu_columnCount, "QMenu_columnCount")
	symaddr(&addr_QMenu_setTabOrder, "QMenu_setTabOrder")
	symaddr(&addr_QMenu_mouseGrabber, "QMenu_mouseGrabber")
	symaddr(&addr_QMenu_keyboardGrabber, "QMenu_keyboardGrabber")
	symaddr(&addr_QMenu_find, "QMenu_find")
	symaddr(&addr_QMenu_createWindowContainer, "QMenu_createWindowContainer")
	symaddr(&addr_QMenu_updateMicroFocus, "QMenu_updateMicroFocus")
	symaddr(&addr_QMenu_create, "QMenu_create")
	symaddr(&addr_QMenu_destroy, "QMenu_destroy")
	symaddr(&addr_QMenu_focusNextChild, "QMenu_focusNextChild")
	symaddr(&addr_QMenu_focusPreviousChild, "QMenu_focusPreviousChild")
	symaddr(&addr_QMenu_connect, "QMenu_connect")
	symaddr(&addr_QMenu_connect2, "QMenu_connect2")
	symaddr(&addr_QMenu_sender, "QMenu_sender")
	symaddr(&addr_QMenu_senderSignalIndex, "QMenu_senderSignalIndex")
	symaddr(&addr_QMenu_receivers, "QMenu_receivers")
	symaddr(&addr_QMenu_isSignalConnected, "QMenu_isSignalConnected")
	reg(&pg_QMenu_devicePixelRatioFScale, "QMenu_devicePixelRatioFScale")
	reg(&pg_QMenu_encodeMetricF, "QMenu_encodeMetricF")
	reg(&pg_QMenu_getDecodedMetricF, "QMenu_getDecodedMetricF")
	symaddr(&addr_QMenu_tr2, "QMenu_tr2")
	symaddr(&addr_QMenu_tr3, "QMenu_tr3")
	symaddr(&addr_QMenu_popup2, "QMenu_popup2")
	symaddr(&addr_QMenu_exec4, "QMenu_exec4")
	symaddr(&addr_QMenu_exec5, "QMenu_exec5")
	symaddr(&addr_QMenu_exec6, "QMenu_exec6")
	symaddr(&addr_QMenu_createWindowContainer2, "QMenu_createWindowContainer2")
	symaddr(&addr_QMenu_createWindowContainer3, "QMenu_createWindowContainer3")
	symaddr(&addr_QMenu_updateMicroFocus2, "QMenu_updateMicroFocus2")
	symaddr(&addr_QMenu_create2, "QMenu_create2")
	symaddr(&addr_QMenu_create3, "QMenu_create3")
	symaddr(&addr_QMenu_create4, "QMenu_create4")
	symaddr(&addr_QMenu_destroy2, "QMenu_destroy2")
	symaddr(&addr_QMenu_destroy3, "QMenu_destroy3")
	symaddr(&addr_QMenu_connect3, "QMenu_connect3")
	symaddr(&addr_QMenu_connect4, "QMenu_connect4")
	symaddr(&addr_QMenu_metaObject_default, "QMenu_metaObject_default")
	symaddr(&addr_QMenu_sizeHint_default, "QMenu_sizeHint_default")
	symaddr(&addr_QMenu_changeEvent_default, "QMenu_changeEvent_default")
	symaddr(&addr_QMenu_keyPressEvent_default, "QMenu_keyPressEvent_default")
	symaddr(&addr_QMenu_mouseReleaseEvent_default, "QMenu_mouseReleaseEvent_default")
	symaddr(&addr_QMenu_mousePressEvent_default, "QMenu_mousePressEvent_default")
	symaddr(&addr_QMenu_mouseMoveEvent_default, "QMenu_mouseMoveEvent_default")
	symaddr(&addr_QMenu_wheelEvent_default, "QMenu_wheelEvent_default")
	symaddr(&addr_QMenu_enterEvent_default, "QMenu_enterEvent_default")
	symaddr(&addr_QMenu_leaveEvent_default, "QMenu_leaveEvent_default")
	symaddr(&addr_QMenu_hideEvent_default, "QMenu_hideEvent_default")
	symaddr(&addr_QMenu_paintEvent_default, "QMenu_paintEvent_default")
	symaddr(&addr_QMenu_actionEvent_default, "QMenu_actionEvent_default")
	symaddr(&addr_QMenu_timerEvent_default, "QMenu_timerEvent_default")
	symaddr(&addr_QMenu_event_default, "QMenu_event_default")
	symaddr(&addr_QMenu_focusNextPrevChild_default, "QMenu_focusNextPrevChild_default")
	symaddr(&addr_QMenu_initStyleOption_default, "QMenu_initStyleOption_default")
	symaddr(&addr_QMenu_devType_default, "QMenu_devType_default")
	symaddr(&addr_QMenu_setVisible_default, "QMenu_setVisible_default")
	symaddr(&addr_QMenu_minimumSizeHint_default, "QMenu_minimumSizeHint_default")
	symaddr(&addr_QMenu_heightForWidth_default, "QMenu_heightForWidth_default")
	symaddr(&addr_QMenu_hasHeightForWidth_default, "QMenu_hasHeightForWidth_default")
	symaddr(&addr_QMenu_paintEngine_default, "QMenu_paintEngine_default")
	symaddr(&addr_QMenu_mouseDoubleClickEvent_default, "QMenu_mouseDoubleClickEvent_default")
	symaddr(&addr_QMenu_keyReleaseEvent_default, "QMenu_keyReleaseEvent_default")
	symaddr(&addr_QMenu_focusInEvent_default, "QMenu_focusInEvent_default")
	symaddr(&addr_QMenu_focusOutEvent_default, "QMenu_focusOutEvent_default")
	symaddr(&addr_QMenu_moveEvent_default, "QMenu_moveEvent_default")
	symaddr(&addr_QMenu_resizeEvent_default, "QMenu_resizeEvent_default")
	symaddr(&addr_QMenu_closeEvent_default, "QMenu_closeEvent_default")
	symaddr(&addr_QMenu_contextMenuEvent_default, "QMenu_contextMenuEvent_default")
	symaddr(&addr_QMenu_tabletEvent_default, "QMenu_tabletEvent_default")
	symaddr(&addr_QMenu_dragEnterEvent_default, "QMenu_dragEnterEvent_default")
	symaddr(&addr_QMenu_dragMoveEvent_default, "QMenu_dragMoveEvent_default")
	symaddr(&addr_QMenu_dragLeaveEvent_default, "QMenu_dragLeaveEvent_default")
	symaddr(&addr_QMenu_dropEvent_default, "QMenu_dropEvent_default")
	symaddr(&addr_QMenu_showEvent_default, "QMenu_showEvent_default")
	symaddr(&addr_QMenu_metric_default, "QMenu_metric_default")
	symaddr(&addr_QMenu_initPainter_default, "QMenu_initPainter_default")
	symaddr(&addr_QMenu_redirected_default, "QMenu_redirected_default")
	symaddr(&addr_QMenu_sharedPainter_default, "QMenu_sharedPainter_default")
	symaddr(&addr_QMenu_inputMethodEvent_default, "QMenu_inputMethodEvent_default")
	symaddr(&addr_QMenu_inputMethodQuery_default, "QMenu_inputMethodQuery_default")
	symaddr(&addr_QMenu_eventFilter_default, "QMenu_eventFilter_default")
	symaddr(&addr_QMenu_childEvent_default, "QMenu_childEvent_default")
	symaddr(&addr_QMenu_customEvent_default, "QMenu_customEvent_default")
	symaddr(&addr_QMenu_connectNotify_default, "QMenu_connectNotify_default")
	symaddr(&addr_QMenu_disconnectNotify_default, "QMenu_disconnectNotify_default")
	reg(&pg_QMenu_connect_aboutToShow, "QMenu_connect_aboutToShow")
	reg(&pg_QMenu_connect_aboutToHide, "QMenu_connect_aboutToHide")
	reg(&pg_QMenu_connect_triggered, "QMenu_connect_triggered")
	reg(&pg_QMenu_connect_hovered, "QMenu_connect_hovered")
	reg(&pg_QMenu_connect_windowTitleChanged, "QMenu_connect_windowTitleChanged")
	reg(&pg_QMenu_connect_windowIconChanged, "QMenu_connect_windowIconChanged")
	reg(&pg_QMenu_connect_windowIconTextChanged, "QMenu_connect_windowIconTextChanged")
	reg(&pg_QMenu_connect_customContextMenuRequested, "QMenu_connect_customContextMenuRequested")
	reg(&pg_QMenu_connect_destroyed, "QMenu_connect_destroyed")
	reg(&pg_QMenu_connect_destroyed2, "QMenu_connect_destroyed2")
	reg(&pg_QMenu_connect_objectNameChanged, "QMenu_connect_objectNameChanged")
	reg(&pg_QMenu__disconnect, "QMenu__disconnect")
	reg(&pg_QMenu__set_override, "QMenu__set_override")
	reg(&pg_QMenu_delete, "QMenu_delete")
}

type QMenu struct {
	h uintptr
	*QWidget
}

func bindQMenu(h uintptr) *QMenu {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QMenu_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMenu{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQMenu(h uintptr) *QMenu { return bindQMenu(h) }

func (this *QMenu) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMenu) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMenu(h unsafe.Pointer) *QMenu { return bindQMenu(uintptr(h)) }
func UnsafeNewQMenuOwned(h unsafe.Pointer) *QMenu { return newownedQMenu(uintptr(h)) }

func (this *QMenu) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQMenu(parent *QWidget) *QMenu {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMenu_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMenu(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenu_virtIdx, QMenu_enable)
	return __o
}

func NewQMenu2() *QMenu {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMenu_new2, uintptr(__goid))
	__o := newownedQMenu(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenu_virtIdx, QMenu_enable)
	return __o
}

func NewQMenu3(title string) *QMenu {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_new3, uintptr(__goid), title_p, uintptr(len(title)))
	__o := newownedQMenu(__r)
	runtime.KeepAlive(title)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenu_virtIdx, QMenu_enable)
	return __o
}

func NewQMenu4(title string, parent *QWidget) *QMenu {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_new4, uintptr(__goid), title_p, uintptr(len(title)), uintptr(parent.UnsafePointer()))
	__o := newownedQMenu(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenu_virtIdx, QMenu_enable)
	return __o
}

func QMenu_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenu_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMenu) AddAction(text string, receiver *QObject, member string, shortcut *QKeySequence) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenu_addAction, this.h, text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(shortcut.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(shortcut)
	return bindQAction(__r)
}

func (this *QMenu) AddAction2(icon *QIcon, text string, receiver *QObject, member string, shortcut *QKeySequence) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenu_addAction2, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(shortcut.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(shortcut)
	return bindQAction(__r)
}

func (this *QMenu) AddMenu(menu *QMenu) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_addMenu, this.h, uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menu)
	return bindQAction(__r)
}

func (this *QMenu) AddMenuWithTitle(title string) *QMenu {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_addMenu2, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	return bindQMenu(__r)
}

func (this *QMenu) AddMenu2(icon *QIcon, title string) *QMenu {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_addMenu3, this.h, uintptr(icon.UnsafePointer()), title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(title)
	return bindQMenu(__r)
}

func (this *QMenu) AddSeparator() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_addSeparator, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenu) AddSection(text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_addSection, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QMenu) AddSection2(icon *QIcon, text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_addSection2, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QMenu) InsertMenu(before *QAction, menu *QMenu) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_insertMenu, this.h, uintptr(before.UnsafePointer()), uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(menu)
	return bindQAction(__r)
}

func (this *QMenu) InsertSeparator(before *QAction) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_insertSeparator, this.h, uintptr(before.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	return bindQAction(__r)
}

func (this *QMenu) InsertSection(before *QAction, text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_insertSection, this.h, uintptr(before.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QMenu) InsertSection2(before *QAction, icon *QIcon, text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_insertSection2, this.h, uintptr(before.UnsafePointer()), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QMenu) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) Clear() {
	purego.SyscallN(addr_QMenu_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) SetTearOffEnabled(arg1 bool) {
	purego.SyscallN(addr_QMenu_setTearOffEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QMenu) IsTearOffEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_isTearOffEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) IsTearOffMenuVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_isTearOffMenuVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) ShowTearOffMenu() {
	purego.SyscallN(addr_QMenu_showTearOffMenu, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) ShowTearOffMenuWithPos(pos *QPoint) {
	purego.SyscallN(addr_QMenu_showTearOffMenu2, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QMenu) HideTearOffMenu() {
	purego.SyscallN(addr_QMenu_hideTearOffMenu, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) SetDefaultAction(arg1 *QAction) {
	purego.SyscallN(addr_QMenu_setDefaultAction, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) DefaultAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_defaultAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenu) SetActiveAction(act *QAction) {
	purego.SyscallN(addr_QMenu_setActiveAction, this.h, uintptr(act.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(act)
}

func (this *QMenu) ActiveAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_activeAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenu) Popup(pos *QPoint) {
	purego.SyscallN(addr_QMenu_popup, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QMenu) Exec() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_exec, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenu) ExecWithPos(pos *QPoint) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_exec2, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return bindQAction(__r)
}

func QMenu_Exec(actions []*QAction, pos *QPoint) *QAction {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_exec3, actions_p, uintptr(len(actions)), uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
	runtime.KeepAlive(pos)
	return bindQAction(__r)
}

func (this *QMenu) ActionGeometry(arg1 *QAction) *QRect {
	__r, _, _ := purego.SyscallN(addr_QMenu_actionGeometry, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func (this *QMenu) ActionAt(arg1 *QPoint) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_actionAt, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAction(__r)
}

func (this *QMenu) MenuAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_menuAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func QMenu_MenuInAction(action *QAction) *QMenu {
	__r, _, _ := purego.SyscallN(addr_QMenu_menuInAction, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(action)
	return bindQMenu(__r)
}

func (this *QMenu) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenu_title, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMenu) SetTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QMenu_setTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QMenu) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QMenu_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QMenu) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QMenu_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QMenu) SetNoReplayFor(widget *QWidget) {
	purego.SyscallN(addr_QMenu_setNoReplayFor, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QMenu) SeparatorsCollapsible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_separatorsCollapsible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	purego.SyscallN(addr_QMenu_setSeparatorsCollapsible, this.h, uintptr(b2i(collapse)))
	runtime.KeepAlive(this)
}

func (this *QMenu) ToolTipsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_toolTipsVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) SetToolTipsVisible(visible bool) {
	purego.SyscallN(addr_QMenu_setToolTipsVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMenu) AboutToShow() {
	purego.SyscallN(addr_QMenu_aboutToShow, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) AboutToHide() {
	purego.SyscallN(addr_QMenu_aboutToHide, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) Triggered(action *QAction) {
	purego.SyscallN(addr_QMenu_triggered, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QMenu) Hovered(action *QAction) {
	purego.SyscallN(addr_QMenu_hovered, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QMenu) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMenu_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QMenu_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMenu_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMenu_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_mouseGrabber)
	return bindQWidget(__r)
}

func QMenu_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_keyboardGrabber)
	return bindQWidget(__r)
}

func QMenu_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMenu_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMenu) UpdateMicroFocus() {
	purego.SyscallN(addr_QMenu_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) Create() {
	purego.SyscallN(addr_QMenu_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) Destroy() {
	purego.SyscallN(addr_QMenu_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenu) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMenu_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenu_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMenu_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMenu_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMenu) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMenu_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMenu) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMenu_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenu) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMenu_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMenu) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMenu_DevicePixelRatioFScale() float64 {
	__r := pg_QMenu_devicePixelRatioFScale()
	return __r
}

func QMenu_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMenu_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMenu) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMenu_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMenu_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenu_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMenu_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenu_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMenu) Popup2(pos *QPoint, at *QAction) {
	purego.SyscallN(addr_QMenu_popup2, this.h, uintptr(pos.UnsafePointer()), uintptr(at.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(at)
}

func (this *QMenu) Exec2(pos *QPoint, at *QAction) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenu_exec4, this.h, uintptr(pos.UnsafePointer()), uintptr(at.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(at)
	return bindQAction(__r)
}

func QMenu_Exec2(actions []*QAction, pos *QPoint, at *QAction) *QAction {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_exec5, actions_p, uintptr(len(actions)), uintptr(pos.UnsafePointer()), uintptr(at.UnsafePointer()))
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(at)
	return bindQAction(__r)
}

func QMenu_Exec3(actions []*QAction, pos *QPoint, at *QAction, parent *QWidget) *QAction {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QMenu_exec6, actions_p, uintptr(len(actions)), uintptr(pos.UnsafePointer()), uintptr(at.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(at)
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func QMenu_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMenu_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenu_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMenu) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMenu_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMenu) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMenu_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMenu) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMenu_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenu) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMenu_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenu) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMenu_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenu) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMenu_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMenu_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenu_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMenu_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMenu_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMenu) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMenu_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMenu) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMenu_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMenu) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMenu_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QMenu_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenu_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenu_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenu_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QMenu_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) EnterEventDefault(arg1 *QEnterEvent) {
	purego.SyscallN(addr_QMenu_enterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) LeaveEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMenu_leaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) HideEventDefault(arg1 *QHideEvent) {
	purego.SyscallN(addr_QMenu_hideEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QMenu_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) ActionEventDefault(arg1 *QActionEvent) {
	purego.SyscallN(addr_QMenu_actionEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QMenu_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QMenu) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) InitStyleOptionDefault(option *QStyleOptionMenuItem, action *QAction) {
	purego.SyscallN(addr_QMenu_initStyleOption_default, this.h, uintptr(option.UnsafePointer()), uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(action)
}

func (this *QMenu) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMenu_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenu) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMenu_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMenu) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMenu_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMenu) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMenu_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenu) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenu) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMenu_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMenu) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMenu_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMenu_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMenu_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QMenu_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QMenu_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QMenu_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QMenu_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QMenu_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMenu_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QMenu_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QMenu_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QMenu_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QMenu_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QMenu_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMenu_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenu) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMenu_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMenu) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMenu_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMenu) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMenu_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMenu) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMenu_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenu) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMenu_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMenu) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenu_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMenu) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMenu_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMenu_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenu) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMenu_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMenu) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMenu_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMenu) OnAboutToShow(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMenu_connect_aboutToShow(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnAboutToHide(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMenu_connect_aboutToHide(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnTriggered(fn func(*QAction)) int64 {
	id := registerCallback(func(action uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(action))) })
	pg_QMenu_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnHovered(fn func(*QAction)) int64 {
	id := registerCallback(func(action uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(action))) })
	pg_QMenu_connect_hovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMenu_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMenu_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMenu_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMenu_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMenu_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMenu_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMenu_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenu) Disconnect(handle int64) {
	pg_QMenu__disconnect(handle)
	unregisterCallback(handle)
}

func QMenu_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 4959376063826165943:
		return 2
	case 1023791977679793185:
		return 3
	case -4415920049297795213:
		return 4
	case -332090898078543319:
		return 5
	case -4255494004099806581:
		return 6
	case 1919988340934511316:
		return 7
	case 6316270346668237895:
		return 8
	case -551070009623711640:
		return 9
	case -199718707876719129:
		return 10
	case 5290442163167458465:
		return 11
	case 1697444140088555581:
		return 12
	case -5409927581812374498:
		return 13
	case 1237236456312824991:
		return 14
	case 5898328517505784191:
		return 15
	case -488693769771308531:
		return 16
	case -8659679893868524044:
		return 17
	case 2110182719375245575:
		return 18
	case 7049264599837532491:
		return 19
	case 4311534060495737339:
		return 20
	case -8266762397995423915:
		return 21
	case 6398007404756346381:
		return 22
	case -4029073905697023887:
		return 23
	case 772598475950050299:
		return 24
	case 975043767033552552:
		return 25
	case -385596010576048993:
		return 26
	case -9161949150489525396:
		return 27
	case -4302917285655014661:
		return 28
	case -1989937050019464783:
		return 29
	case -6698043222227006015:
		return 30
	case -212068866860701141:
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

func QMenu_enable(h uintptr, idx int32) { pg_QMenu__set_override(h, idx) }

func (this *QMenu) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMenu) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMenu) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QMenu) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QMenu) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QMenu) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QMenu) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QMenu) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QMenu) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(arg1 uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(arg1))) }) // slot "enterEvent"
}

func (this *QMenu) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "leaveEvent"
}

func (this *QMenu) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(arg1 uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(arg1))) }) // slot "hideEvent"
}

func (this *QMenu) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QMenu) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(arg1 uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(arg1))) }) // slot "actionEvent"
}

func (this *QMenu) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QMenu) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QMenu) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMenu) OnInitStyleOption(fn func(*QStyleOptionMenuItem, *QAction)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr, action uintptr) { fn(UnsafeNewQStyleOptionMenuItem(unsafe.Pointer(option)), UnsafeNewQAction(unsafe.Pointer(action))) }) // slot "initStyleOption"
}

func (this *QMenu) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMenu) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMenu) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMenu) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMenu) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMenu) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMenu) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMenu) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMenu) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QMenu) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QMenu) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QMenu) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QMenu) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QMenu) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QMenu) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMenu) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QMenu) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QMenu) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QMenu) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QMenu) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QMenu) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMenu) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMenu) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMenu) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMenu) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMenu) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMenu) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMenu) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMenu) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMenu) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMenu) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMenu) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMenu) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMenu_delete(this.h)
	}
	this.UnsafeRelease()
}
