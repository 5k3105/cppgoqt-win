package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMenuBar_chain uintptr
	addr_QMenuBar_new uintptr
	addr_QMenuBar_new2 uintptr
	addr_QMenuBar_tr uintptr
	addr_QMenuBar_addMenu uintptr
	addr_QMenuBar_addMenu2 uintptr
	addr_QMenuBar_addMenu3 uintptr
	addr_QMenuBar_addSeparator uintptr
	addr_QMenuBar_insertSeparator uintptr
	addr_QMenuBar_insertMenu uintptr
	addr_QMenuBar_clear uintptr
	addr_QMenuBar_activeAction uintptr
	addr_QMenuBar_setActiveAction uintptr
	addr_QMenuBar_setDefaultUp uintptr
	addr_QMenuBar_isDefaultUp uintptr
	addr_QMenuBar_actionGeometry uintptr
	addr_QMenuBar_actionAt uintptr
	addr_QMenuBar_setCornerWidget uintptr
	addr_QMenuBar_cornerWidget uintptr
	addr_QMenuBar_isNativeMenuBar uintptr
	addr_QMenuBar_setNativeMenuBar uintptr
	addr_QMenuBar_triggered uintptr
	addr_QMenuBar_hovered uintptr
	addr_QMenuBar_setTabOrder uintptr
	addr_QMenuBar_mouseGrabber uintptr
	addr_QMenuBar_keyboardGrabber uintptr
	addr_QMenuBar_find uintptr
	addr_QMenuBar_createWindowContainer uintptr
	addr_QMenuBar_updateMicroFocus uintptr
	addr_QMenuBar_create uintptr
	addr_QMenuBar_destroy uintptr
	addr_QMenuBar_focusNextChild uintptr
	addr_QMenuBar_focusPreviousChild uintptr
	addr_QMenuBar_connect uintptr
	addr_QMenuBar_connect2 uintptr
	addr_QMenuBar_sender uintptr
	addr_QMenuBar_senderSignalIndex uintptr
	addr_QMenuBar_receivers uintptr
	addr_QMenuBar_isSignalConnected uintptr
	pg_QMenuBar_devicePixelRatioFScale func() float64
	pg_QMenuBar_encodeMetricF func(int32, float64) int32
	pg_QMenuBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMenuBar_tr2 uintptr
	addr_QMenuBar_tr3 uintptr
	addr_QMenuBar_setCornerWidget2 uintptr
	addr_QMenuBar_cornerWidget2 uintptr
	addr_QMenuBar_createWindowContainer2 uintptr
	addr_QMenuBar_createWindowContainer3 uintptr
	addr_QMenuBar_updateMicroFocus2 uintptr
	addr_QMenuBar_create2 uintptr
	addr_QMenuBar_create3 uintptr
	addr_QMenuBar_create4 uintptr
	addr_QMenuBar_destroy2 uintptr
	addr_QMenuBar_destroy3 uintptr
	addr_QMenuBar_connect3 uintptr
	addr_QMenuBar_connect4 uintptr
	addr_QMenuBar_metaObject_default uintptr
	addr_QMenuBar_sizeHint_default uintptr
	addr_QMenuBar_minimumSizeHint_default uintptr
	addr_QMenuBar_heightForWidth_default uintptr
	addr_QMenuBar_setVisible_default uintptr
	addr_QMenuBar_changeEvent_default uintptr
	addr_QMenuBar_keyPressEvent_default uintptr
	addr_QMenuBar_mouseReleaseEvent_default uintptr
	addr_QMenuBar_mousePressEvent_default uintptr
	addr_QMenuBar_mouseMoveEvent_default uintptr
	addr_QMenuBar_leaveEvent_default uintptr
	addr_QMenuBar_paintEvent_default uintptr
	addr_QMenuBar_resizeEvent_default uintptr
	addr_QMenuBar_actionEvent_default uintptr
	addr_QMenuBar_focusOutEvent_default uintptr
	addr_QMenuBar_focusInEvent_default uintptr
	addr_QMenuBar_timerEvent_default uintptr
	addr_QMenuBar_eventFilter_default uintptr
	addr_QMenuBar_event_default uintptr
	addr_QMenuBar_initStyleOption_default uintptr
	addr_QMenuBar_devType_default uintptr
	addr_QMenuBar_hasHeightForWidth_default uintptr
	addr_QMenuBar_paintEngine_default uintptr
	addr_QMenuBar_mouseDoubleClickEvent_default uintptr
	addr_QMenuBar_wheelEvent_default uintptr
	addr_QMenuBar_keyReleaseEvent_default uintptr
	addr_QMenuBar_enterEvent_default uintptr
	addr_QMenuBar_moveEvent_default uintptr
	addr_QMenuBar_closeEvent_default uintptr
	addr_QMenuBar_contextMenuEvent_default uintptr
	addr_QMenuBar_tabletEvent_default uintptr
	addr_QMenuBar_dragEnterEvent_default uintptr
	addr_QMenuBar_dragMoveEvent_default uintptr
	addr_QMenuBar_dragLeaveEvent_default uintptr
	addr_QMenuBar_dropEvent_default uintptr
	addr_QMenuBar_showEvent_default uintptr
	addr_QMenuBar_hideEvent_default uintptr
	addr_QMenuBar_metric_default uintptr
	addr_QMenuBar_initPainter_default uintptr
	addr_QMenuBar_redirected_default uintptr
	addr_QMenuBar_sharedPainter_default uintptr
	addr_QMenuBar_inputMethodEvent_default uintptr
	addr_QMenuBar_inputMethodQuery_default uintptr
	addr_QMenuBar_focusNextPrevChild_default uintptr
	addr_QMenuBar_childEvent_default uintptr
	addr_QMenuBar_customEvent_default uintptr
	addr_QMenuBar_connectNotify_default uintptr
	addr_QMenuBar_disconnectNotify_default uintptr
	pg_QMenuBar_connect_triggered func(uintptr, int64)
	pg_QMenuBar_connect_hovered func(uintptr, int64)
	pg_QMenuBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QMenuBar_connect_windowIconChanged func(uintptr, int64)
	pg_QMenuBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMenuBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMenuBar_connect_destroyed func(uintptr, int64)
	pg_QMenuBar_connect_destroyed2 func(uintptr, int64)
	pg_QMenuBar_connect_objectNameChanged func(uintptr, int64)
	pg_QMenuBar__disconnect func(int64)
	pg_QMenuBar__set_override func(uintptr, int32)
	pg_QMenuBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMenuBar_chain, "QMenuBar_chain")
	symaddr(&addr_QMenuBar_new, "QMenuBar_new")
	symaddr(&addr_QMenuBar_new2, "QMenuBar_new2")
	symaddr(&addr_QMenuBar_tr, "QMenuBar_tr")
	symaddr(&addr_QMenuBar_addMenu, "QMenuBar_addMenu")
	symaddr(&addr_QMenuBar_addMenu2, "QMenuBar_addMenu2")
	symaddr(&addr_QMenuBar_addMenu3, "QMenuBar_addMenu3")
	symaddr(&addr_QMenuBar_addSeparator, "QMenuBar_addSeparator")
	symaddr(&addr_QMenuBar_insertSeparator, "QMenuBar_insertSeparator")
	symaddr(&addr_QMenuBar_insertMenu, "QMenuBar_insertMenu")
	symaddr(&addr_QMenuBar_clear, "QMenuBar_clear")
	symaddr(&addr_QMenuBar_activeAction, "QMenuBar_activeAction")
	symaddr(&addr_QMenuBar_setActiveAction, "QMenuBar_setActiveAction")
	symaddr(&addr_QMenuBar_setDefaultUp, "QMenuBar_setDefaultUp")
	symaddr(&addr_QMenuBar_isDefaultUp, "QMenuBar_isDefaultUp")
	symaddr(&addr_QMenuBar_actionGeometry, "QMenuBar_actionGeometry")
	symaddr(&addr_QMenuBar_actionAt, "QMenuBar_actionAt")
	symaddr(&addr_QMenuBar_setCornerWidget, "QMenuBar_setCornerWidget")
	symaddr(&addr_QMenuBar_cornerWidget, "QMenuBar_cornerWidget")
	symaddr(&addr_QMenuBar_isNativeMenuBar, "QMenuBar_isNativeMenuBar")
	symaddr(&addr_QMenuBar_setNativeMenuBar, "QMenuBar_setNativeMenuBar")
	symaddr(&addr_QMenuBar_triggered, "QMenuBar_triggered")
	symaddr(&addr_QMenuBar_hovered, "QMenuBar_hovered")
	symaddr(&addr_QMenuBar_setTabOrder, "QMenuBar_setTabOrder")
	symaddr(&addr_QMenuBar_mouseGrabber, "QMenuBar_mouseGrabber")
	symaddr(&addr_QMenuBar_keyboardGrabber, "QMenuBar_keyboardGrabber")
	symaddr(&addr_QMenuBar_find, "QMenuBar_find")
	symaddr(&addr_QMenuBar_createWindowContainer, "QMenuBar_createWindowContainer")
	symaddr(&addr_QMenuBar_updateMicroFocus, "QMenuBar_updateMicroFocus")
	symaddr(&addr_QMenuBar_create, "QMenuBar_create")
	symaddr(&addr_QMenuBar_destroy, "QMenuBar_destroy")
	symaddr(&addr_QMenuBar_focusNextChild, "QMenuBar_focusNextChild")
	symaddr(&addr_QMenuBar_focusPreviousChild, "QMenuBar_focusPreviousChild")
	symaddr(&addr_QMenuBar_connect, "QMenuBar_connect")
	symaddr(&addr_QMenuBar_connect2, "QMenuBar_connect2")
	symaddr(&addr_QMenuBar_sender, "QMenuBar_sender")
	symaddr(&addr_QMenuBar_senderSignalIndex, "QMenuBar_senderSignalIndex")
	symaddr(&addr_QMenuBar_receivers, "QMenuBar_receivers")
	symaddr(&addr_QMenuBar_isSignalConnected, "QMenuBar_isSignalConnected")
	reg(&pg_QMenuBar_devicePixelRatioFScale, "QMenuBar_devicePixelRatioFScale")
	reg(&pg_QMenuBar_encodeMetricF, "QMenuBar_encodeMetricF")
	reg(&pg_QMenuBar_getDecodedMetricF, "QMenuBar_getDecodedMetricF")
	symaddr(&addr_QMenuBar_tr2, "QMenuBar_tr2")
	symaddr(&addr_QMenuBar_tr3, "QMenuBar_tr3")
	symaddr(&addr_QMenuBar_setCornerWidget2, "QMenuBar_setCornerWidget2")
	symaddr(&addr_QMenuBar_cornerWidget2, "QMenuBar_cornerWidget2")
	symaddr(&addr_QMenuBar_createWindowContainer2, "QMenuBar_createWindowContainer2")
	symaddr(&addr_QMenuBar_createWindowContainer3, "QMenuBar_createWindowContainer3")
	symaddr(&addr_QMenuBar_updateMicroFocus2, "QMenuBar_updateMicroFocus2")
	symaddr(&addr_QMenuBar_create2, "QMenuBar_create2")
	symaddr(&addr_QMenuBar_create3, "QMenuBar_create3")
	symaddr(&addr_QMenuBar_create4, "QMenuBar_create4")
	symaddr(&addr_QMenuBar_destroy2, "QMenuBar_destroy2")
	symaddr(&addr_QMenuBar_destroy3, "QMenuBar_destroy3")
	symaddr(&addr_QMenuBar_connect3, "QMenuBar_connect3")
	symaddr(&addr_QMenuBar_connect4, "QMenuBar_connect4")
	symaddr(&addr_QMenuBar_metaObject_default, "QMenuBar_metaObject_default")
	symaddr(&addr_QMenuBar_sizeHint_default, "QMenuBar_sizeHint_default")
	symaddr(&addr_QMenuBar_minimumSizeHint_default, "QMenuBar_minimumSizeHint_default")
	symaddr(&addr_QMenuBar_heightForWidth_default, "QMenuBar_heightForWidth_default")
	symaddr(&addr_QMenuBar_setVisible_default, "QMenuBar_setVisible_default")
	symaddr(&addr_QMenuBar_changeEvent_default, "QMenuBar_changeEvent_default")
	symaddr(&addr_QMenuBar_keyPressEvent_default, "QMenuBar_keyPressEvent_default")
	symaddr(&addr_QMenuBar_mouseReleaseEvent_default, "QMenuBar_mouseReleaseEvent_default")
	symaddr(&addr_QMenuBar_mousePressEvent_default, "QMenuBar_mousePressEvent_default")
	symaddr(&addr_QMenuBar_mouseMoveEvent_default, "QMenuBar_mouseMoveEvent_default")
	symaddr(&addr_QMenuBar_leaveEvent_default, "QMenuBar_leaveEvent_default")
	symaddr(&addr_QMenuBar_paintEvent_default, "QMenuBar_paintEvent_default")
	symaddr(&addr_QMenuBar_resizeEvent_default, "QMenuBar_resizeEvent_default")
	symaddr(&addr_QMenuBar_actionEvent_default, "QMenuBar_actionEvent_default")
	symaddr(&addr_QMenuBar_focusOutEvent_default, "QMenuBar_focusOutEvent_default")
	symaddr(&addr_QMenuBar_focusInEvent_default, "QMenuBar_focusInEvent_default")
	symaddr(&addr_QMenuBar_timerEvent_default, "QMenuBar_timerEvent_default")
	symaddr(&addr_QMenuBar_eventFilter_default, "QMenuBar_eventFilter_default")
	symaddr(&addr_QMenuBar_event_default, "QMenuBar_event_default")
	symaddr(&addr_QMenuBar_initStyleOption_default, "QMenuBar_initStyleOption_default")
	symaddr(&addr_QMenuBar_devType_default, "QMenuBar_devType_default")
	symaddr(&addr_QMenuBar_hasHeightForWidth_default, "QMenuBar_hasHeightForWidth_default")
	symaddr(&addr_QMenuBar_paintEngine_default, "QMenuBar_paintEngine_default")
	symaddr(&addr_QMenuBar_mouseDoubleClickEvent_default, "QMenuBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QMenuBar_wheelEvent_default, "QMenuBar_wheelEvent_default")
	symaddr(&addr_QMenuBar_keyReleaseEvent_default, "QMenuBar_keyReleaseEvent_default")
	symaddr(&addr_QMenuBar_enterEvent_default, "QMenuBar_enterEvent_default")
	symaddr(&addr_QMenuBar_moveEvent_default, "QMenuBar_moveEvent_default")
	symaddr(&addr_QMenuBar_closeEvent_default, "QMenuBar_closeEvent_default")
	symaddr(&addr_QMenuBar_contextMenuEvent_default, "QMenuBar_contextMenuEvent_default")
	symaddr(&addr_QMenuBar_tabletEvent_default, "QMenuBar_tabletEvent_default")
	symaddr(&addr_QMenuBar_dragEnterEvent_default, "QMenuBar_dragEnterEvent_default")
	symaddr(&addr_QMenuBar_dragMoveEvent_default, "QMenuBar_dragMoveEvent_default")
	symaddr(&addr_QMenuBar_dragLeaveEvent_default, "QMenuBar_dragLeaveEvent_default")
	symaddr(&addr_QMenuBar_dropEvent_default, "QMenuBar_dropEvent_default")
	symaddr(&addr_QMenuBar_showEvent_default, "QMenuBar_showEvent_default")
	symaddr(&addr_QMenuBar_hideEvent_default, "QMenuBar_hideEvent_default")
	symaddr(&addr_QMenuBar_metric_default, "QMenuBar_metric_default")
	symaddr(&addr_QMenuBar_initPainter_default, "QMenuBar_initPainter_default")
	symaddr(&addr_QMenuBar_redirected_default, "QMenuBar_redirected_default")
	symaddr(&addr_QMenuBar_sharedPainter_default, "QMenuBar_sharedPainter_default")
	symaddr(&addr_QMenuBar_inputMethodEvent_default, "QMenuBar_inputMethodEvent_default")
	symaddr(&addr_QMenuBar_inputMethodQuery_default, "QMenuBar_inputMethodQuery_default")
	symaddr(&addr_QMenuBar_focusNextPrevChild_default, "QMenuBar_focusNextPrevChild_default")
	symaddr(&addr_QMenuBar_childEvent_default, "QMenuBar_childEvent_default")
	symaddr(&addr_QMenuBar_customEvent_default, "QMenuBar_customEvent_default")
	symaddr(&addr_QMenuBar_connectNotify_default, "QMenuBar_connectNotify_default")
	symaddr(&addr_QMenuBar_disconnectNotify_default, "QMenuBar_disconnectNotify_default")
	reg(&pg_QMenuBar_connect_triggered, "QMenuBar_connect_triggered")
	reg(&pg_QMenuBar_connect_hovered, "QMenuBar_connect_hovered")
	reg(&pg_QMenuBar_connect_windowTitleChanged, "QMenuBar_connect_windowTitleChanged")
	reg(&pg_QMenuBar_connect_windowIconChanged, "QMenuBar_connect_windowIconChanged")
	reg(&pg_QMenuBar_connect_windowIconTextChanged, "QMenuBar_connect_windowIconTextChanged")
	reg(&pg_QMenuBar_connect_customContextMenuRequested, "QMenuBar_connect_customContextMenuRequested")
	reg(&pg_QMenuBar_connect_destroyed, "QMenuBar_connect_destroyed")
	reg(&pg_QMenuBar_connect_destroyed2, "QMenuBar_connect_destroyed2")
	reg(&pg_QMenuBar_connect_objectNameChanged, "QMenuBar_connect_objectNameChanged")
	reg(&pg_QMenuBar__disconnect, "QMenuBar__disconnect")
	reg(&pg_QMenuBar__set_override, "QMenuBar__set_override")
	reg(&pg_QMenuBar_delete, "QMenuBar_delete")
}

type QMenuBar struct {
	h uintptr
	*QWidget
}

func bindQMenuBar(h uintptr) *QMenuBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QMenuBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMenuBar{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQMenuBar(h uintptr) *QMenuBar { return bindQMenuBar(h) }

func (this *QMenuBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMenuBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMenuBar(h unsafe.Pointer) *QMenuBar { return bindQMenuBar(uintptr(h)) }
func UnsafeNewQMenuBarOwned(h unsafe.Pointer) *QMenuBar { return newownedQMenuBar(uintptr(h)) }

func (this *QMenuBar) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQMenuBar(parent *QWidget) *QMenuBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMenuBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMenuBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenuBar_virtIdx, QMenuBar_enable)
	return __o
}

func NewQMenuBar2() *QMenuBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMenuBar_new2, uintptr(__goid))
	__o := newownedQMenuBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMenuBar_virtIdx, QMenuBar_enable)
	return __o
}

func QMenuBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenuBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMenuBar) AddMenu(menu *QMenu) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_addMenu, this.h, uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menu)
	return bindQAction(__r)
}

func (this *QMenuBar) AddMenuWithTitle(title string) *QMenu {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenuBar_addMenu2, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	return bindQMenu(__r)
}

func (this *QMenuBar) AddMenu2(icon *QIcon, title string) *QMenu {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QMenuBar_addMenu3, this.h, uintptr(icon.UnsafePointer()), title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(title)
	return bindQMenu(__r)
}

func (this *QMenuBar) AddSeparator() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_addSeparator, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenuBar) InsertSeparator(before *QAction) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_insertSeparator, this.h, uintptr(before.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	return bindQAction(__r)
}

func (this *QMenuBar) InsertMenu(before *QAction, menu *QMenu) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_insertMenu, this.h, uintptr(before.UnsafePointer()), uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(menu)
	return bindQAction(__r)
}

func (this *QMenuBar) Clear() {
	purego.SyscallN(addr_QMenuBar_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenuBar) ActiveAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_activeAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QMenuBar) SetActiveAction(action *QAction) {
	purego.SyscallN(addr_QMenuBar_setActiveAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QMenuBar) SetDefaultUp(arg1 bool) {
	purego.SyscallN(addr_QMenuBar_setDefaultUp, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) IsDefaultUp() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_isDefaultUp, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenuBar) ActionGeometry(arg1 *QAction) *QRect {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_actionGeometry, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func (this *QMenuBar) ActionAt(arg1 *QPoint) *QAction {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_actionAt, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAction(__r)
}

func (this *QMenuBar) SetCornerWidget(w *QWidget) {
	purego.SyscallN(addr_QMenuBar_setCornerWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QMenuBar) CornerWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_cornerWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMenuBar) IsNativeMenuBar() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_isNativeMenuBar, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenuBar) SetNativeMenuBar(nativeMenuBar bool) {
	purego.SyscallN(addr_QMenuBar_setNativeMenuBar, this.h, uintptr(b2i(nativeMenuBar)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Triggered(action *QAction) {
	purego.SyscallN(addr_QMenuBar_triggered, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QMenuBar) Hovered(action *QAction) {
	purego.SyscallN(addr_QMenuBar_hovered, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func QMenuBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMenuBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMenuBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_mouseGrabber)
	return bindQWidget(__r)
}

func QMenuBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QMenuBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMenuBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMenuBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QMenuBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Create() {
	purego.SyscallN(addr_QMenuBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Destroy() {
	purego.SyscallN(addr_QMenuBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMenuBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenuBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMenuBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenuBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMenuBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMenuBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMenuBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenuBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMenuBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMenuBar_DevicePixelRatioFScale() float64 {
	__r := pg_QMenuBar_devicePixelRatioFScale()
	return __r
}

func QMenuBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMenuBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMenuBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMenuBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMenuBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenuBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMenuBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMenuBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMenuBar) SetCornerWidget2(w *QWidget, corner Corner) {
	purego.SyscallN(addr_QMenuBar_setCornerWidget2, this.h, uintptr(w.UnsafePointer()), uintptr(corner))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QMenuBar) CornerWidgetWithCorner(corner Corner) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_cornerWidget2, this.h, uintptr(corner))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func QMenuBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMenuBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMenuBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMenuBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMenuBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMenuBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMenuBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMenuBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMenuBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMenuBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMenuBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMenuBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMenuBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMenuBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMenuBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMenuBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenuBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMenuBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMenuBar) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMenuBar_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QMenuBar_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenuBar_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenuBar_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QMenuBar_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) LeaveEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QMenuBar_leaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QMenuBar_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QMenuBar_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) ActionEventDefault(arg1 *QActionEvent) {
	purego.SyscallN(addr_QMenuBar_actionEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QMenuBar_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QMenuBar_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QMenuBar_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QMenuBar) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QMenuBar) InitStyleOptionDefault(option *QStyleOptionMenuItem, action *QAction) {
	purego.SyscallN(addr_QMenuBar_initStyleOption_default, this.h, uintptr(option.UnsafePointer()), uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(action)
}

func (this *QMenuBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenuBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenuBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMenuBar) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QMenuBar_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QMenuBar_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMenuBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QMenuBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QMenuBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QMenuBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QMenuBar_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMenuBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QMenuBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QMenuBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QMenuBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QMenuBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QMenuBar_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QMenuBar_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMenuBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMenuBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMenuBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMenuBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMenuBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMenuBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMenuBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMenuBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMenuBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMenuBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMenuBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMenuBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMenuBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMenuBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMenuBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMenuBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMenuBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMenuBar) OnTriggered(fn func(*QAction)) int64 {
	id := registerCallback(func(action uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(action))) })
	pg_QMenuBar_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnHovered(fn func(*QAction)) int64 {
	id := registerCallback(func(action uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(action))) })
	pg_QMenuBar_connect_hovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMenuBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMenuBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMenuBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMenuBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMenuBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMenuBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMenuBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMenuBar) Disconnect(handle int64) {
	pg_QMenuBar__disconnect(handle)
	unregisterCallback(handle)
}

func QMenuBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 4311534060495737339:
		return 3
	case 2110182719375245575:
		return 4
	case 4959376063826165943:
		return 5
	case 1023791977679793185:
		return 6
	case -4415920049297795213:
		return 7
	case -332090898078543319:
		return 8
	case -4255494004099806581:
		return 9
	case -551070009623711640:
		return 10
	case 5290442163167458465:
		return 11
	case -4302917285655014661:
		return 12
	case 1697444140088555581:
		return 13
	case -385596010576048993:
		return 14
	case 975043767033552552:
		return 15
	case -5409927581812374498:
		return 16
	case 5217623514630858297:
		return 17
	case 1237236456312824991:
		return 18
	case -488693769771308531:
		return 19
	case -8659679893868524044:
		return 20
	case -8266762397995423915:
		return 21
	case 6398007404756346381:
		return 22
	case -4029073905697023887:
		return 23
	case 1919988340934511316:
		return 24
	case 772598475950050299:
		return 25
	case 6316270346668237895:
		return 26
	case -9161949150489525396:
		return 27
	case -1989937050019464783:
		return 28
	case -6698043222227006015:
		return 29
	case -212068866860701141:
		return 30
	case 206122715919917589:
		return 31
	case -1781545137352745978:
		return 32
	case -3139904158373725006:
		return 33
	case 9107058682754591716:
		return 34
	case -6825369174751916:
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

func QMenuBar_enable(h uintptr, idx int32) { pg_QMenuBar__set_override(h, idx) }

func (this *QMenuBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMenuBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMenuBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMenuBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMenuBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMenuBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QMenuBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QMenuBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QMenuBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QMenuBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QMenuBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "leaveEvent"
}

func (this *QMenuBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QMenuBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QMenuBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(arg1 uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(arg1))) }) // slot "actionEvent"
}

func (this *QMenuBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QMenuBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QMenuBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QMenuBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QMenuBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QMenuBar) OnInitStyleOption(fn func(*QStyleOptionMenuItem, *QAction)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr, action uintptr) { fn(UnsafeNewQStyleOptionMenuItem(unsafe.Pointer(option)), UnsafeNewQAction(unsafe.Pointer(action))) }) // slot "initStyleOption"
}

func (this *QMenuBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMenuBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMenuBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMenuBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMenuBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QMenuBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMenuBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QMenuBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QMenuBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QMenuBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QMenuBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMenuBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QMenuBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QMenuBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QMenuBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QMenuBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QMenuBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QMenuBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMenuBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMenuBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMenuBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMenuBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMenuBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMenuBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMenuBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMenuBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMenuBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMenuBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMenuBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMenuBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMenuBar_delete(this.h)
	}
	this.UnsafeRelease()
}
