package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMdiSubWindow_chain uintptr
	addr_QMdiSubWindow_new uintptr
	addr_QMdiSubWindow_new2 uintptr
	addr_QMdiSubWindow_new3 uintptr
	addr_QMdiSubWindow_tr uintptr
	addr_QMdiSubWindow_setWidget uintptr
	addr_QMdiSubWindow_widget uintptr
	addr_QMdiSubWindow_maximizedButtonsWidget uintptr
	addr_QMdiSubWindow_maximizedSystemMenuIconWidget uintptr
	addr_QMdiSubWindow_isShaded uintptr
	addr_QMdiSubWindow_setOption uintptr
	addr_QMdiSubWindow_testOption uintptr
	addr_QMdiSubWindow_setKeyboardSingleStep uintptr
	addr_QMdiSubWindow_keyboardSingleStep uintptr
	addr_QMdiSubWindow_setKeyboardPageStep uintptr
	addr_QMdiSubWindow_keyboardPageStep uintptr
	addr_QMdiSubWindow_setSystemMenu uintptr
	addr_QMdiSubWindow_systemMenu uintptr
	addr_QMdiSubWindow_mdiArea uintptr
	addr_QMdiSubWindow_windowStateChanged uintptr
	addr_QMdiSubWindow_aboutToActivate uintptr
	addr_QMdiSubWindow_showSystemMenu uintptr
	addr_QMdiSubWindow_showShaded uintptr
	addr_QMdiSubWindow_setTabOrder uintptr
	addr_QMdiSubWindow_mouseGrabber uintptr
	addr_QMdiSubWindow_keyboardGrabber uintptr
	addr_QMdiSubWindow_find uintptr
	addr_QMdiSubWindow_createWindowContainer uintptr
	addr_QMdiSubWindow_updateMicroFocus uintptr
	addr_QMdiSubWindow_create uintptr
	addr_QMdiSubWindow_destroy uintptr
	addr_QMdiSubWindow_focusNextChild uintptr
	addr_QMdiSubWindow_focusPreviousChild uintptr
	addr_QMdiSubWindow_connect uintptr
	addr_QMdiSubWindow_connect2 uintptr
	addr_QMdiSubWindow_sender uintptr
	addr_QMdiSubWindow_senderSignalIndex uintptr
	addr_QMdiSubWindow_receivers uintptr
	addr_QMdiSubWindow_isSignalConnected uintptr
	pg_QMdiSubWindow_devicePixelRatioFScale func() float64
	pg_QMdiSubWindow_encodeMetricF func(int32, float64) int32
	pg_QMdiSubWindow_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QMdiSubWindow_tr2 uintptr
	addr_QMdiSubWindow_tr3 uintptr
	addr_QMdiSubWindow_setOption2 uintptr
	addr_QMdiSubWindow_createWindowContainer2 uintptr
	addr_QMdiSubWindow_createWindowContainer3 uintptr
	addr_QMdiSubWindow_updateMicroFocus2 uintptr
	addr_QMdiSubWindow_create2 uintptr
	addr_QMdiSubWindow_create3 uintptr
	addr_QMdiSubWindow_create4 uintptr
	addr_QMdiSubWindow_destroy2 uintptr
	addr_QMdiSubWindow_destroy3 uintptr
	addr_QMdiSubWindow_connect3 uintptr
	addr_QMdiSubWindow_connect4 uintptr
	addr_QMdiSubWindow_metaObject_default uintptr
	addr_QMdiSubWindow_sizeHint_default uintptr
	addr_QMdiSubWindow_minimumSizeHint_default uintptr
	addr_QMdiSubWindow_eventFilter_default uintptr
	addr_QMdiSubWindow_event_default uintptr
	addr_QMdiSubWindow_showEvent_default uintptr
	addr_QMdiSubWindow_hideEvent_default uintptr
	addr_QMdiSubWindow_changeEvent_default uintptr
	addr_QMdiSubWindow_closeEvent_default uintptr
	addr_QMdiSubWindow_leaveEvent_default uintptr
	addr_QMdiSubWindow_resizeEvent_default uintptr
	addr_QMdiSubWindow_timerEvent_default uintptr
	addr_QMdiSubWindow_moveEvent_default uintptr
	addr_QMdiSubWindow_paintEvent_default uintptr
	addr_QMdiSubWindow_mousePressEvent_default uintptr
	addr_QMdiSubWindow_mouseDoubleClickEvent_default uintptr
	addr_QMdiSubWindow_mouseReleaseEvent_default uintptr
	addr_QMdiSubWindow_mouseMoveEvent_default uintptr
	addr_QMdiSubWindow_keyPressEvent_default uintptr
	addr_QMdiSubWindow_contextMenuEvent_default uintptr
	addr_QMdiSubWindow_focusInEvent_default uintptr
	addr_QMdiSubWindow_focusOutEvent_default uintptr
	addr_QMdiSubWindow_childEvent_default uintptr
	addr_QMdiSubWindow_devType_default uintptr
	addr_QMdiSubWindow_setVisible_default uintptr
	addr_QMdiSubWindow_heightForWidth_default uintptr
	addr_QMdiSubWindow_hasHeightForWidth_default uintptr
	addr_QMdiSubWindow_paintEngine_default uintptr
	addr_QMdiSubWindow_wheelEvent_default uintptr
	addr_QMdiSubWindow_keyReleaseEvent_default uintptr
	addr_QMdiSubWindow_enterEvent_default uintptr
	addr_QMdiSubWindow_tabletEvent_default uintptr
	addr_QMdiSubWindow_actionEvent_default uintptr
	addr_QMdiSubWindow_dragEnterEvent_default uintptr
	addr_QMdiSubWindow_dragMoveEvent_default uintptr
	addr_QMdiSubWindow_dragLeaveEvent_default uintptr
	addr_QMdiSubWindow_dropEvent_default uintptr
	addr_QMdiSubWindow_metric_default uintptr
	addr_QMdiSubWindow_initPainter_default uintptr
	addr_QMdiSubWindow_redirected_default uintptr
	addr_QMdiSubWindow_sharedPainter_default uintptr
	addr_QMdiSubWindow_inputMethodEvent_default uintptr
	addr_QMdiSubWindow_inputMethodQuery_default uintptr
	addr_QMdiSubWindow_focusNextPrevChild_default uintptr
	addr_QMdiSubWindow_customEvent_default uintptr
	addr_QMdiSubWindow_connectNotify_default uintptr
	addr_QMdiSubWindow_disconnectNotify_default uintptr
	pg_QMdiSubWindow_connect_windowStateChanged func(uintptr, int64)
	pg_QMdiSubWindow_connect_aboutToActivate func(uintptr, int64)
	pg_QMdiSubWindow_connect_windowTitleChanged func(uintptr, int64)
	pg_QMdiSubWindow_connect_windowIconChanged func(uintptr, int64)
	pg_QMdiSubWindow_connect_windowIconTextChanged func(uintptr, int64)
	pg_QMdiSubWindow_connect_customContextMenuRequested func(uintptr, int64)
	pg_QMdiSubWindow_connect_destroyed func(uintptr, int64)
	pg_QMdiSubWindow_connect_destroyed2 func(uintptr, int64)
	pg_QMdiSubWindow_connect_objectNameChanged func(uintptr, int64)
	pg_QMdiSubWindow__disconnect func(int64)
	pg_QMdiSubWindow__set_override func(uintptr, int32)
	pg_QMdiSubWindow_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMdiSubWindow_chain, "QMdiSubWindow_chain")
	symaddr(&addr_QMdiSubWindow_new, "QMdiSubWindow_new")
	symaddr(&addr_QMdiSubWindow_new2, "QMdiSubWindow_new2")
	symaddr(&addr_QMdiSubWindow_new3, "QMdiSubWindow_new3")
	symaddr(&addr_QMdiSubWindow_tr, "QMdiSubWindow_tr")
	symaddr(&addr_QMdiSubWindow_setWidget, "QMdiSubWindow_setWidget")
	symaddr(&addr_QMdiSubWindow_widget, "QMdiSubWindow_widget")
	symaddr(&addr_QMdiSubWindow_maximizedButtonsWidget, "QMdiSubWindow_maximizedButtonsWidget")
	symaddr(&addr_QMdiSubWindow_maximizedSystemMenuIconWidget, "QMdiSubWindow_maximizedSystemMenuIconWidget")
	symaddr(&addr_QMdiSubWindow_isShaded, "QMdiSubWindow_isShaded")
	symaddr(&addr_QMdiSubWindow_setOption, "QMdiSubWindow_setOption")
	symaddr(&addr_QMdiSubWindow_testOption, "QMdiSubWindow_testOption")
	symaddr(&addr_QMdiSubWindow_setKeyboardSingleStep, "QMdiSubWindow_setKeyboardSingleStep")
	symaddr(&addr_QMdiSubWindow_keyboardSingleStep, "QMdiSubWindow_keyboardSingleStep")
	symaddr(&addr_QMdiSubWindow_setKeyboardPageStep, "QMdiSubWindow_setKeyboardPageStep")
	symaddr(&addr_QMdiSubWindow_keyboardPageStep, "QMdiSubWindow_keyboardPageStep")
	symaddr(&addr_QMdiSubWindow_setSystemMenu, "QMdiSubWindow_setSystemMenu")
	symaddr(&addr_QMdiSubWindow_systemMenu, "QMdiSubWindow_systemMenu")
	symaddr(&addr_QMdiSubWindow_mdiArea, "QMdiSubWindow_mdiArea")
	symaddr(&addr_QMdiSubWindow_windowStateChanged, "QMdiSubWindow_windowStateChanged")
	symaddr(&addr_QMdiSubWindow_aboutToActivate, "QMdiSubWindow_aboutToActivate")
	symaddr(&addr_QMdiSubWindow_showSystemMenu, "QMdiSubWindow_showSystemMenu")
	symaddr(&addr_QMdiSubWindow_showShaded, "QMdiSubWindow_showShaded")
	symaddr(&addr_QMdiSubWindow_setTabOrder, "QMdiSubWindow_setTabOrder")
	symaddr(&addr_QMdiSubWindow_mouseGrabber, "QMdiSubWindow_mouseGrabber")
	symaddr(&addr_QMdiSubWindow_keyboardGrabber, "QMdiSubWindow_keyboardGrabber")
	symaddr(&addr_QMdiSubWindow_find, "QMdiSubWindow_find")
	symaddr(&addr_QMdiSubWindow_createWindowContainer, "QMdiSubWindow_createWindowContainer")
	symaddr(&addr_QMdiSubWindow_updateMicroFocus, "QMdiSubWindow_updateMicroFocus")
	symaddr(&addr_QMdiSubWindow_create, "QMdiSubWindow_create")
	symaddr(&addr_QMdiSubWindow_destroy, "QMdiSubWindow_destroy")
	symaddr(&addr_QMdiSubWindow_focusNextChild, "QMdiSubWindow_focusNextChild")
	symaddr(&addr_QMdiSubWindow_focusPreviousChild, "QMdiSubWindow_focusPreviousChild")
	symaddr(&addr_QMdiSubWindow_connect, "QMdiSubWindow_connect")
	symaddr(&addr_QMdiSubWindow_connect2, "QMdiSubWindow_connect2")
	symaddr(&addr_QMdiSubWindow_sender, "QMdiSubWindow_sender")
	symaddr(&addr_QMdiSubWindow_senderSignalIndex, "QMdiSubWindow_senderSignalIndex")
	symaddr(&addr_QMdiSubWindow_receivers, "QMdiSubWindow_receivers")
	symaddr(&addr_QMdiSubWindow_isSignalConnected, "QMdiSubWindow_isSignalConnected")
	reg(&pg_QMdiSubWindow_devicePixelRatioFScale, "QMdiSubWindow_devicePixelRatioFScale")
	reg(&pg_QMdiSubWindow_encodeMetricF, "QMdiSubWindow_encodeMetricF")
	reg(&pg_QMdiSubWindow_getDecodedMetricF, "QMdiSubWindow_getDecodedMetricF")
	symaddr(&addr_QMdiSubWindow_tr2, "QMdiSubWindow_tr2")
	symaddr(&addr_QMdiSubWindow_tr3, "QMdiSubWindow_tr3")
	symaddr(&addr_QMdiSubWindow_setOption2, "QMdiSubWindow_setOption2")
	symaddr(&addr_QMdiSubWindow_createWindowContainer2, "QMdiSubWindow_createWindowContainer2")
	symaddr(&addr_QMdiSubWindow_createWindowContainer3, "QMdiSubWindow_createWindowContainer3")
	symaddr(&addr_QMdiSubWindow_updateMicroFocus2, "QMdiSubWindow_updateMicroFocus2")
	symaddr(&addr_QMdiSubWindow_create2, "QMdiSubWindow_create2")
	symaddr(&addr_QMdiSubWindow_create3, "QMdiSubWindow_create3")
	symaddr(&addr_QMdiSubWindow_create4, "QMdiSubWindow_create4")
	symaddr(&addr_QMdiSubWindow_destroy2, "QMdiSubWindow_destroy2")
	symaddr(&addr_QMdiSubWindow_destroy3, "QMdiSubWindow_destroy3")
	symaddr(&addr_QMdiSubWindow_connect3, "QMdiSubWindow_connect3")
	symaddr(&addr_QMdiSubWindow_connect4, "QMdiSubWindow_connect4")
	symaddr(&addr_QMdiSubWindow_metaObject_default, "QMdiSubWindow_metaObject_default")
	symaddr(&addr_QMdiSubWindow_sizeHint_default, "QMdiSubWindow_sizeHint_default")
	symaddr(&addr_QMdiSubWindow_minimumSizeHint_default, "QMdiSubWindow_minimumSizeHint_default")
	symaddr(&addr_QMdiSubWindow_eventFilter_default, "QMdiSubWindow_eventFilter_default")
	symaddr(&addr_QMdiSubWindow_event_default, "QMdiSubWindow_event_default")
	symaddr(&addr_QMdiSubWindow_showEvent_default, "QMdiSubWindow_showEvent_default")
	symaddr(&addr_QMdiSubWindow_hideEvent_default, "QMdiSubWindow_hideEvent_default")
	symaddr(&addr_QMdiSubWindow_changeEvent_default, "QMdiSubWindow_changeEvent_default")
	symaddr(&addr_QMdiSubWindow_closeEvent_default, "QMdiSubWindow_closeEvent_default")
	symaddr(&addr_QMdiSubWindow_leaveEvent_default, "QMdiSubWindow_leaveEvent_default")
	symaddr(&addr_QMdiSubWindow_resizeEvent_default, "QMdiSubWindow_resizeEvent_default")
	symaddr(&addr_QMdiSubWindow_timerEvent_default, "QMdiSubWindow_timerEvent_default")
	symaddr(&addr_QMdiSubWindow_moveEvent_default, "QMdiSubWindow_moveEvent_default")
	symaddr(&addr_QMdiSubWindow_paintEvent_default, "QMdiSubWindow_paintEvent_default")
	symaddr(&addr_QMdiSubWindow_mousePressEvent_default, "QMdiSubWindow_mousePressEvent_default")
	symaddr(&addr_QMdiSubWindow_mouseDoubleClickEvent_default, "QMdiSubWindow_mouseDoubleClickEvent_default")
	symaddr(&addr_QMdiSubWindow_mouseReleaseEvent_default, "QMdiSubWindow_mouseReleaseEvent_default")
	symaddr(&addr_QMdiSubWindow_mouseMoveEvent_default, "QMdiSubWindow_mouseMoveEvent_default")
	symaddr(&addr_QMdiSubWindow_keyPressEvent_default, "QMdiSubWindow_keyPressEvent_default")
	symaddr(&addr_QMdiSubWindow_contextMenuEvent_default, "QMdiSubWindow_contextMenuEvent_default")
	symaddr(&addr_QMdiSubWindow_focusInEvent_default, "QMdiSubWindow_focusInEvent_default")
	symaddr(&addr_QMdiSubWindow_focusOutEvent_default, "QMdiSubWindow_focusOutEvent_default")
	symaddr(&addr_QMdiSubWindow_childEvent_default, "QMdiSubWindow_childEvent_default")
	symaddr(&addr_QMdiSubWindow_devType_default, "QMdiSubWindow_devType_default")
	symaddr(&addr_QMdiSubWindow_setVisible_default, "QMdiSubWindow_setVisible_default")
	symaddr(&addr_QMdiSubWindow_heightForWidth_default, "QMdiSubWindow_heightForWidth_default")
	symaddr(&addr_QMdiSubWindow_hasHeightForWidth_default, "QMdiSubWindow_hasHeightForWidth_default")
	symaddr(&addr_QMdiSubWindow_paintEngine_default, "QMdiSubWindow_paintEngine_default")
	symaddr(&addr_QMdiSubWindow_wheelEvent_default, "QMdiSubWindow_wheelEvent_default")
	symaddr(&addr_QMdiSubWindow_keyReleaseEvent_default, "QMdiSubWindow_keyReleaseEvent_default")
	symaddr(&addr_QMdiSubWindow_enterEvent_default, "QMdiSubWindow_enterEvent_default")
	symaddr(&addr_QMdiSubWindow_tabletEvent_default, "QMdiSubWindow_tabletEvent_default")
	symaddr(&addr_QMdiSubWindow_actionEvent_default, "QMdiSubWindow_actionEvent_default")
	symaddr(&addr_QMdiSubWindow_dragEnterEvent_default, "QMdiSubWindow_dragEnterEvent_default")
	symaddr(&addr_QMdiSubWindow_dragMoveEvent_default, "QMdiSubWindow_dragMoveEvent_default")
	symaddr(&addr_QMdiSubWindow_dragLeaveEvent_default, "QMdiSubWindow_dragLeaveEvent_default")
	symaddr(&addr_QMdiSubWindow_dropEvent_default, "QMdiSubWindow_dropEvent_default")
	symaddr(&addr_QMdiSubWindow_metric_default, "QMdiSubWindow_metric_default")
	symaddr(&addr_QMdiSubWindow_initPainter_default, "QMdiSubWindow_initPainter_default")
	symaddr(&addr_QMdiSubWindow_redirected_default, "QMdiSubWindow_redirected_default")
	symaddr(&addr_QMdiSubWindow_sharedPainter_default, "QMdiSubWindow_sharedPainter_default")
	symaddr(&addr_QMdiSubWindow_inputMethodEvent_default, "QMdiSubWindow_inputMethodEvent_default")
	symaddr(&addr_QMdiSubWindow_inputMethodQuery_default, "QMdiSubWindow_inputMethodQuery_default")
	symaddr(&addr_QMdiSubWindow_focusNextPrevChild_default, "QMdiSubWindow_focusNextPrevChild_default")
	symaddr(&addr_QMdiSubWindow_customEvent_default, "QMdiSubWindow_customEvent_default")
	symaddr(&addr_QMdiSubWindow_connectNotify_default, "QMdiSubWindow_connectNotify_default")
	symaddr(&addr_QMdiSubWindow_disconnectNotify_default, "QMdiSubWindow_disconnectNotify_default")
	reg(&pg_QMdiSubWindow_connect_windowStateChanged, "QMdiSubWindow_connect_windowStateChanged")
	reg(&pg_QMdiSubWindow_connect_aboutToActivate, "QMdiSubWindow_connect_aboutToActivate")
	reg(&pg_QMdiSubWindow_connect_windowTitleChanged, "QMdiSubWindow_connect_windowTitleChanged")
	reg(&pg_QMdiSubWindow_connect_windowIconChanged, "QMdiSubWindow_connect_windowIconChanged")
	reg(&pg_QMdiSubWindow_connect_windowIconTextChanged, "QMdiSubWindow_connect_windowIconTextChanged")
	reg(&pg_QMdiSubWindow_connect_customContextMenuRequested, "QMdiSubWindow_connect_customContextMenuRequested")
	reg(&pg_QMdiSubWindow_connect_destroyed, "QMdiSubWindow_connect_destroyed")
	reg(&pg_QMdiSubWindow_connect_destroyed2, "QMdiSubWindow_connect_destroyed2")
	reg(&pg_QMdiSubWindow_connect_objectNameChanged, "QMdiSubWindow_connect_objectNameChanged")
	reg(&pg_QMdiSubWindow__disconnect, "QMdiSubWindow__disconnect")
	reg(&pg_QMdiSubWindow__set_override, "QMdiSubWindow__set_override")
	reg(&pg_QMdiSubWindow_delete, "QMdiSubWindow_delete")
}

type QMdiSubWindow struct {
	h uintptr
	*QWidget
}

func bindQMdiSubWindow(h uintptr) *QMdiSubWindow {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QMdiSubWindow_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QMdiSubWindow{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQMdiSubWindow(h uintptr) *QMdiSubWindow { return bindQMdiSubWindow(h) }

func (this *QMdiSubWindow) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMdiSubWindow) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMdiSubWindow(h unsafe.Pointer) *QMdiSubWindow { return bindQMdiSubWindow(uintptr(h)) }
func UnsafeNewQMdiSubWindowOwned(h unsafe.Pointer) *QMdiSubWindow { return newownedQMdiSubWindow(uintptr(h)) }

func (this *QMdiSubWindow) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQMdiSubWindow(parent *QWidget) *QMdiSubWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMdiSubWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMdiSubWindow_virtIdx, QMdiSubWindow_enable)
	return __o
}

func NewQMdiSubWindow2() *QMdiSubWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_new2, uintptr(__goid))
	__o := newownedQMdiSubWindow(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMdiSubWindow_virtIdx, QMdiSubWindow_enable)
	return __o
}

func NewQMdiSubWindow3(parent *QWidget, flags WindowType) *QMdiSubWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQMdiSubWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMdiSubWindow_virtIdx, QMdiSubWindow_enable)
	return __o
}

func QMdiSubWindow_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMdiSubWindow) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QMdiSubWindow_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QMdiSubWindow) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMdiSubWindow) MaximizedButtonsWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_maximizedButtonsWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMdiSubWindow) MaximizedSystemMenuIconWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_maximizedSystemMenuIconWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QMdiSubWindow) IsShaded() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_isShaded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiSubWindow) SetOption(option QMdiSubWindow__SubWindowOption) {
	purego.SyscallN(addr_QMdiSubWindow_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) TestOption(arg1 QMdiSubWindow__SubWindowOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_testOption, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiSubWindow) SetKeyboardSingleStep(step int) {
	purego.SyscallN(addr_QMdiSubWindow_setKeyboardSingleStep, this.h, uintptr(step))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) KeyboardSingleStep() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_keyboardSingleStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) SetKeyboardPageStep(step int) {
	purego.SyscallN(addr_QMdiSubWindow_setKeyboardPageStep, this.h, uintptr(step))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) KeyboardPageStep() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_keyboardPageStep, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) SetSystemMenu(systemMenu *QMenu) {
	purego.SyscallN(addr_QMdiSubWindow_setSystemMenu, this.h, uintptr(systemMenu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(systemMenu)
}

func (this *QMdiSubWindow) SystemMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_systemMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QMdiSubWindow) MdiArea() *QMdiArea {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_mdiArea, this.h)
	runtime.KeepAlive(this)
	return bindQMdiArea(__r)
}

func (this *QMdiSubWindow) WindowStateChanged(oldState WindowState, newState WindowState) {
	purego.SyscallN(addr_QMdiSubWindow_windowStateChanged, this.h, uintptr(oldState), uintptr(newState))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) AboutToActivate() {
	purego.SyscallN(addr_QMdiSubWindow_aboutToActivate, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) ShowSystemMenu() {
	purego.SyscallN(addr_QMdiSubWindow_showSystemMenu, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) ShowShaded() {
	purego.SyscallN(addr_QMdiSubWindow_showShaded, this.h)
	runtime.KeepAlive(this)
}

func QMdiSubWindow_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QMdiSubWindow_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QMdiSubWindow_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_mouseGrabber)
	return bindQWidget(__r)
}

func QMdiSubWindow_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_keyboardGrabber)
	return bindQWidget(__r)
}

func QMdiSubWindow_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QMdiSubWindow_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QMdiSubWindow) UpdateMicroFocus() {
	purego.SyscallN(addr_QMdiSubWindow_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) Create() {
	purego.SyscallN(addr_QMdiSubWindow_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) Destroy() {
	purego.SyscallN(addr_QMdiSubWindow_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiSubWindow) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QMdiSubWindow_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMdiSubWindow_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMdiSubWindow) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMdiSubWindow) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMdiSubWindow) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMdiSubWindow_DevicePixelRatioFScale() float64 {
	__r := pg_QMdiSubWindow_devicePixelRatioFScale()
	return __r
}

func QMdiSubWindow_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QMdiSubWindow_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QMdiSubWindow) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QMdiSubWindow_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QMdiSubWindow_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMdiSubWindow_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMdiSubWindow) SetOption2(option QMdiSubWindow__SubWindowOption, on bool) {
	purego.SyscallN(addr_QMdiSubWindow_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QMdiSubWindow_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QMdiSubWindow_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QMdiSubWindow) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QMdiSubWindow_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QMdiSubWindow_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QMdiSubWindow_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QMdiSubWindow_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QMdiSubWindow_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QMdiSubWindow_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QMdiSubWindow_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMdiSubWindow_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMdiSubWindow) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMdiSubWindow) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMdiSubWindow) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMdiSubWindow) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMdiSubWindow) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMdiSubWindow) ShowEventDefault(showEvent *QShowEvent) {
	purego.SyscallN(addr_QMdiSubWindow_showEvent_default, this.h, uintptr(showEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(showEvent)
}

func (this *QMdiSubWindow) HideEventDefault(hideEvent *QHideEvent) {
	purego.SyscallN(addr_QMdiSubWindow_hideEvent_default, this.h, uintptr(hideEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(hideEvent)
}

func (this *QMdiSubWindow) ChangeEventDefault(changeEvent *QEvent) {
	purego.SyscallN(addr_QMdiSubWindow_changeEvent_default, this.h, uintptr(changeEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(changeEvent)
}

func (this *QMdiSubWindow) CloseEventDefault(closeEvent *QCloseEvent) {
	purego.SyscallN(addr_QMdiSubWindow_closeEvent_default, this.h, uintptr(closeEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(closeEvent)
}

func (this *QMdiSubWindow) LeaveEventDefault(leaveEvent *QEvent) {
	purego.SyscallN(addr_QMdiSubWindow_leaveEvent_default, this.h, uintptr(leaveEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(leaveEvent)
}

func (this *QMdiSubWindow) ResizeEventDefault(resizeEvent *QResizeEvent) {
	purego.SyscallN(addr_QMdiSubWindow_resizeEvent_default, this.h, uintptr(resizeEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(resizeEvent)
}

func (this *QMdiSubWindow) TimerEventDefault(timerEvent *QTimerEvent) {
	purego.SyscallN(addr_QMdiSubWindow_timerEvent_default, this.h, uintptr(timerEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timerEvent)
}

func (this *QMdiSubWindow) MoveEventDefault(moveEvent *QMoveEvent) {
	purego.SyscallN(addr_QMdiSubWindow_moveEvent_default, this.h, uintptr(moveEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(moveEvent)
}

func (this *QMdiSubWindow) PaintEventDefault(paintEvent *QPaintEvent) {
	purego.SyscallN(addr_QMdiSubWindow_paintEvent_default, this.h, uintptr(paintEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(paintEvent)
}

func (this *QMdiSubWindow) MousePressEventDefault(mouseEvent *QMouseEvent) {
	purego.SyscallN(addr_QMdiSubWindow_mousePressEvent_default, this.h, uintptr(mouseEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mouseEvent)
}

func (this *QMdiSubWindow) MouseDoubleClickEventDefault(mouseEvent *QMouseEvent) {
	purego.SyscallN(addr_QMdiSubWindow_mouseDoubleClickEvent_default, this.h, uintptr(mouseEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mouseEvent)
}

func (this *QMdiSubWindow) MouseReleaseEventDefault(mouseEvent *QMouseEvent) {
	purego.SyscallN(addr_QMdiSubWindow_mouseReleaseEvent_default, this.h, uintptr(mouseEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mouseEvent)
}

func (this *QMdiSubWindow) MouseMoveEventDefault(mouseEvent *QMouseEvent) {
	purego.SyscallN(addr_QMdiSubWindow_mouseMoveEvent_default, this.h, uintptr(mouseEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mouseEvent)
}

func (this *QMdiSubWindow) KeyPressEventDefault(keyEvent *QKeyEvent) {
	purego.SyscallN(addr_QMdiSubWindow_keyPressEvent_default, this.h, uintptr(keyEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keyEvent)
}

func (this *QMdiSubWindow) ContextMenuEventDefault(contextMenuEvent *QContextMenuEvent) {
	purego.SyscallN(addr_QMdiSubWindow_contextMenuEvent_default, this.h, uintptr(contextMenuEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(contextMenuEvent)
}

func (this *QMdiSubWindow) FocusInEventDefault(focusInEvent *QFocusEvent) {
	purego.SyscallN(addr_QMdiSubWindow_focusInEvent_default, this.h, uintptr(focusInEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(focusInEvent)
}

func (this *QMdiSubWindow) FocusOutEventDefault(focusOutEvent *QFocusEvent) {
	purego.SyscallN(addr_QMdiSubWindow_focusOutEvent_default, this.h, uintptr(focusOutEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(focusOutEvent)
}

func (this *QMdiSubWindow) ChildEventDefault(childEvent *QChildEvent) {
	purego.SyscallN(addr_QMdiSubWindow_childEvent_default, this.h, uintptr(childEvent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childEvent)
}

func (this *QMdiSubWindow) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QMdiSubWindow_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QMdiSubWindow) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiSubWindow) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QMdiSubWindow) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QMdiSubWindow_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QMdiSubWindow_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QMdiSubWindow_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QMdiSubWindow_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QMdiSubWindow_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QMdiSubWindow_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QMdiSubWindow_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QMdiSubWindow_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QMdiSubWindow_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMdiSubWindow) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QMdiSubWindow_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QMdiSubWindow) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QMdiSubWindow) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QMdiSubWindow) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QMdiSubWindow_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QMdiSubWindow) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMdiSubWindow_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMdiSubWindow) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QMdiSubWindow_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMdiSubWindow) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMdiSubWindow_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMdiSubWindow) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMdiSubWindow_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMdiSubWindow) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMdiSubWindow_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMdiSubWindow) OnWindowStateChanged(fn func(WindowState, WindowState)) int64 {
	id := registerCallback(func(oldState int32, newState int32) { fn(WindowState(oldState), WindowState(newState)) })
	pg_QMdiSubWindow_connect_windowStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnAboutToActivate(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMdiSubWindow_connect_aboutToActivate(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QMdiSubWindow_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QMdiSubWindow_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QMdiSubWindow_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QMdiSubWindow_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMdiSubWindow_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMdiSubWindow_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMdiSubWindow_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMdiSubWindow) Disconnect(handle int64) {
	pg_QMdiSubWindow__disconnect(handle)
	unregisterCallback(handle)
}

func QMdiSubWindow_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 5217623514630858297:
		return 3
	case 1237236456312824991:
		return 4
	case -6825369174751916:
		return 5
	case -199718707876719129:
		return 6
	case 4959376063826165943:
		return 7
	case -1989937050019464783:
		return 8
	case -551070009623711640:
		return 9
	case -4302917285655014661:
		return 10
	case -5409927581812374498:
		return 11
	case -9161949150489525396:
		return 12
	case 5290442163167458465:
		return 13
	case -332090898078543319:
		return 14
	case -4029073905697023887:
		return 15
	case -4415920049297795213:
		return 16
	case -4255494004099806581:
		return 17
	case 1023791977679793185:
		return 18
	case -6698043222227006015:
		return 19
	case 975043767033552552:
		return 20
	case -385596010576048993:
		return 21
	case -1185315367526857833:
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
	case 772598475950050299:
		return 29
	case 6316270346668237895:
		return 30
	case -212068866860701141:
		return 31
	case 1697444140088555581:
		return 32
	case 206122715919917589:
		return 33
	case -1781545137352745978:
		return 34
	case -3139904158373725006:
		return 35
	case 9107058682754591716:
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
	case 3769868229123868474:
		return 45
	case 7379029982170503008:
		return 46
	case -568369772982215962:
		return 47
	}
	return -1
}

func QMdiSubWindow_enable(h uintptr, idx int32) { pg_QMdiSubWindow__set_override(h, idx) }

func (this *QMdiSubWindow) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMdiSubWindow) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QMdiSubWindow) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QMdiSubWindow) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMdiSubWindow) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QMdiSubWindow) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(showEvent uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(showEvent))) }) // slot "showEvent"
}

func (this *QMdiSubWindow) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(hideEvent uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(hideEvent))) }) // slot "hideEvent"
}

func (this *QMdiSubWindow) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(changeEvent uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(changeEvent))) }) // slot "changeEvent"
}

func (this *QMdiSubWindow) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(closeEvent uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(closeEvent))) }) // slot "closeEvent"
}

func (this *QMdiSubWindow) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(leaveEvent uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(leaveEvent))) }) // slot "leaveEvent"
}

func (this *QMdiSubWindow) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(resizeEvent uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(resizeEvent))) }) // slot "resizeEvent"
}

func (this *QMdiSubWindow) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(timerEvent uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(timerEvent))) }) // slot "timerEvent"
}

func (this *QMdiSubWindow) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(moveEvent uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(moveEvent))) }) // slot "moveEvent"
}

func (this *QMdiSubWindow) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(paintEvent uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(paintEvent))) }) // slot "paintEvent"
}

func (this *QMdiSubWindow) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(mouseEvent uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(mouseEvent))) }) // slot "mousePressEvent"
}

func (this *QMdiSubWindow) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(mouseEvent uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(mouseEvent))) }) // slot "mouseDoubleClickEvent"
}

func (this *QMdiSubWindow) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(mouseEvent uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(mouseEvent))) }) // slot "mouseReleaseEvent"
}

func (this *QMdiSubWindow) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(mouseEvent uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(mouseEvent))) }) // slot "mouseMoveEvent"
}

func (this *QMdiSubWindow) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(keyEvent uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(keyEvent))) }) // slot "keyPressEvent"
}

func (this *QMdiSubWindow) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(contextMenuEvent uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(contextMenuEvent))) }) // slot "contextMenuEvent"
}

func (this *QMdiSubWindow) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(focusInEvent uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(focusInEvent))) }) // slot "focusInEvent"
}

func (this *QMdiSubWindow) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(focusOutEvent uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(focusOutEvent))) }) // slot "focusOutEvent"
}

func (this *QMdiSubWindow) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(childEvent uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(childEvent))) }) // slot "childEvent"
}

func (this *QMdiSubWindow) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QMdiSubWindow) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QMdiSubWindow) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QMdiSubWindow) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QMdiSubWindow) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QMdiSubWindow) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QMdiSubWindow) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QMdiSubWindow) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QMdiSubWindow) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QMdiSubWindow) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QMdiSubWindow) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QMdiSubWindow) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QMdiSubWindow) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QMdiSubWindow) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QMdiSubWindow) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QMdiSubWindow) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QMdiSubWindow) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QMdiSubWindow) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QMdiSubWindow) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QMdiSubWindow) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QMdiSubWindow) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QMdiSubWindow) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QMdiSubWindow) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMdiSubWindow) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMdiSubWindow) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMdiSubWindow) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMdiSubWindow_delete(this.h)
	}
	this.UnsafeRelease()
}
