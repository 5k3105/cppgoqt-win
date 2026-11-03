package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSplashScreen_chain uintptr
	addr_QSplashScreen_new uintptr
	addr_QSplashScreen_new2 uintptr
	addr_QSplashScreen_new3 uintptr
	addr_QSplashScreen_new4 uintptr
	addr_QSplashScreen_new5 uintptr
	addr_QSplashScreen_new6 uintptr
	addr_QSplashScreen_tr uintptr
	addr_QSplashScreen_setPixmap uintptr
	addr_QSplashScreen_pixmap uintptr
	addr_QSplashScreen_finish uintptr
	addr_QSplashScreen_repaint uintptr
	addr_QSplashScreen_message uintptr
	addr_QSplashScreen_showMessage uintptr
	addr_QSplashScreen_clearMessage uintptr
	addr_QSplashScreen_messageChanged uintptr
	addr_QSplashScreen_setTabOrder uintptr
	addr_QSplashScreen_mouseGrabber uintptr
	addr_QSplashScreen_keyboardGrabber uintptr
	addr_QSplashScreen_find uintptr
	addr_QSplashScreen_createWindowContainer uintptr
	addr_QSplashScreen_updateMicroFocus uintptr
	addr_QSplashScreen_create uintptr
	addr_QSplashScreen_destroy uintptr
	addr_QSplashScreen_focusNextChild uintptr
	addr_QSplashScreen_focusPreviousChild uintptr
	addr_QSplashScreen_connect uintptr
	addr_QSplashScreen_connect2 uintptr
	addr_QSplashScreen_sender uintptr
	addr_QSplashScreen_senderSignalIndex uintptr
	addr_QSplashScreen_receivers uintptr
	addr_QSplashScreen_isSignalConnected uintptr
	pg_QSplashScreen_devicePixelRatioFScale func() float64
	pg_QSplashScreen_encodeMetricF func(int32, float64) int32
	pg_QSplashScreen_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QSplashScreen_tr2 uintptr
	addr_QSplashScreen_tr3 uintptr
	addr_QSplashScreen_showMessage2 uintptr
	addr_QSplashScreen_showMessage3 uintptr
	addr_QSplashScreen_createWindowContainer2 uintptr
	addr_QSplashScreen_createWindowContainer3 uintptr
	addr_QSplashScreen_updateMicroFocus2 uintptr
	addr_QSplashScreen_create2 uintptr
	addr_QSplashScreen_create3 uintptr
	addr_QSplashScreen_create4 uintptr
	addr_QSplashScreen_destroy2 uintptr
	addr_QSplashScreen_destroy3 uintptr
	addr_QSplashScreen_connect3 uintptr
	addr_QSplashScreen_connect4 uintptr
	addr_QSplashScreen_metaObject_default uintptr
	addr_QSplashScreen_event_default uintptr
	addr_QSplashScreen_drawContents_default uintptr
	addr_QSplashScreen_mousePressEvent_default uintptr
	addr_QSplashScreen_devType_default uintptr
	addr_QSplashScreen_setVisible_default uintptr
	addr_QSplashScreen_sizeHint_default uintptr
	addr_QSplashScreen_minimumSizeHint_default uintptr
	addr_QSplashScreen_heightForWidth_default uintptr
	addr_QSplashScreen_hasHeightForWidth_default uintptr
	addr_QSplashScreen_paintEngine_default uintptr
	addr_QSplashScreen_mouseReleaseEvent_default uintptr
	addr_QSplashScreen_mouseDoubleClickEvent_default uintptr
	addr_QSplashScreen_mouseMoveEvent_default uintptr
	addr_QSplashScreen_wheelEvent_default uintptr
	addr_QSplashScreen_keyPressEvent_default uintptr
	addr_QSplashScreen_keyReleaseEvent_default uintptr
	addr_QSplashScreen_focusInEvent_default uintptr
	addr_QSplashScreen_focusOutEvent_default uintptr
	addr_QSplashScreen_enterEvent_default uintptr
	addr_QSplashScreen_leaveEvent_default uintptr
	addr_QSplashScreen_paintEvent_default uintptr
	addr_QSplashScreen_moveEvent_default uintptr
	addr_QSplashScreen_resizeEvent_default uintptr
	addr_QSplashScreen_closeEvent_default uintptr
	addr_QSplashScreen_contextMenuEvent_default uintptr
	addr_QSplashScreen_tabletEvent_default uintptr
	addr_QSplashScreen_actionEvent_default uintptr
	addr_QSplashScreen_dragEnterEvent_default uintptr
	addr_QSplashScreen_dragMoveEvent_default uintptr
	addr_QSplashScreen_dragLeaveEvent_default uintptr
	addr_QSplashScreen_dropEvent_default uintptr
	addr_QSplashScreen_showEvent_default uintptr
	addr_QSplashScreen_hideEvent_default uintptr
	addr_QSplashScreen_changeEvent_default uintptr
	addr_QSplashScreen_metric_default uintptr
	addr_QSplashScreen_initPainter_default uintptr
	addr_QSplashScreen_redirected_default uintptr
	addr_QSplashScreen_sharedPainter_default uintptr
	addr_QSplashScreen_inputMethodEvent_default uintptr
	addr_QSplashScreen_inputMethodQuery_default uintptr
	addr_QSplashScreen_focusNextPrevChild_default uintptr
	addr_QSplashScreen_eventFilter_default uintptr
	addr_QSplashScreen_timerEvent_default uintptr
	addr_QSplashScreen_childEvent_default uintptr
	addr_QSplashScreen_customEvent_default uintptr
	addr_QSplashScreen_connectNotify_default uintptr
	addr_QSplashScreen_disconnectNotify_default uintptr
	pg_QSplashScreen_connect_messageChanged func(uintptr, int64)
	pg_QSplashScreen_connect_windowTitleChanged func(uintptr, int64)
	pg_QSplashScreen_connect_windowIconChanged func(uintptr, int64)
	pg_QSplashScreen_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSplashScreen_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSplashScreen_connect_destroyed func(uintptr, int64)
	pg_QSplashScreen_connect_destroyed2 func(uintptr, int64)
	pg_QSplashScreen_connect_objectNameChanged func(uintptr, int64)
	pg_QSplashScreen__disconnect func(int64)
	pg_QSplashScreen__set_override func(uintptr, int32)
	pg_QSplashScreen_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSplashScreen_chain, "QSplashScreen_chain")
	symaddr(&addr_QSplashScreen_new, "QSplashScreen_new")
	symaddr(&addr_QSplashScreen_new2, "QSplashScreen_new2")
	symaddr(&addr_QSplashScreen_new3, "QSplashScreen_new3")
	symaddr(&addr_QSplashScreen_new4, "QSplashScreen_new4")
	symaddr(&addr_QSplashScreen_new5, "QSplashScreen_new5")
	symaddr(&addr_QSplashScreen_new6, "QSplashScreen_new6")
	symaddr(&addr_QSplashScreen_tr, "QSplashScreen_tr")
	symaddr(&addr_QSplashScreen_setPixmap, "QSplashScreen_setPixmap")
	symaddr(&addr_QSplashScreen_pixmap, "QSplashScreen_pixmap")
	symaddr(&addr_QSplashScreen_finish, "QSplashScreen_finish")
	symaddr(&addr_QSplashScreen_repaint, "QSplashScreen_repaint")
	symaddr(&addr_QSplashScreen_message, "QSplashScreen_message")
	symaddr(&addr_QSplashScreen_showMessage, "QSplashScreen_showMessage")
	symaddr(&addr_QSplashScreen_clearMessage, "QSplashScreen_clearMessage")
	symaddr(&addr_QSplashScreen_messageChanged, "QSplashScreen_messageChanged")
	symaddr(&addr_QSplashScreen_setTabOrder, "QSplashScreen_setTabOrder")
	symaddr(&addr_QSplashScreen_mouseGrabber, "QSplashScreen_mouseGrabber")
	symaddr(&addr_QSplashScreen_keyboardGrabber, "QSplashScreen_keyboardGrabber")
	symaddr(&addr_QSplashScreen_find, "QSplashScreen_find")
	symaddr(&addr_QSplashScreen_createWindowContainer, "QSplashScreen_createWindowContainer")
	symaddr(&addr_QSplashScreen_updateMicroFocus, "QSplashScreen_updateMicroFocus")
	symaddr(&addr_QSplashScreen_create, "QSplashScreen_create")
	symaddr(&addr_QSplashScreen_destroy, "QSplashScreen_destroy")
	symaddr(&addr_QSplashScreen_focusNextChild, "QSplashScreen_focusNextChild")
	symaddr(&addr_QSplashScreen_focusPreviousChild, "QSplashScreen_focusPreviousChild")
	symaddr(&addr_QSplashScreen_connect, "QSplashScreen_connect")
	symaddr(&addr_QSplashScreen_connect2, "QSplashScreen_connect2")
	symaddr(&addr_QSplashScreen_sender, "QSplashScreen_sender")
	symaddr(&addr_QSplashScreen_senderSignalIndex, "QSplashScreen_senderSignalIndex")
	symaddr(&addr_QSplashScreen_receivers, "QSplashScreen_receivers")
	symaddr(&addr_QSplashScreen_isSignalConnected, "QSplashScreen_isSignalConnected")
	reg(&pg_QSplashScreen_devicePixelRatioFScale, "QSplashScreen_devicePixelRatioFScale")
	reg(&pg_QSplashScreen_encodeMetricF, "QSplashScreen_encodeMetricF")
	reg(&pg_QSplashScreen_getDecodedMetricF, "QSplashScreen_getDecodedMetricF")
	symaddr(&addr_QSplashScreen_tr2, "QSplashScreen_tr2")
	symaddr(&addr_QSplashScreen_tr3, "QSplashScreen_tr3")
	symaddr(&addr_QSplashScreen_showMessage2, "QSplashScreen_showMessage2")
	symaddr(&addr_QSplashScreen_showMessage3, "QSplashScreen_showMessage3")
	symaddr(&addr_QSplashScreen_createWindowContainer2, "QSplashScreen_createWindowContainer2")
	symaddr(&addr_QSplashScreen_createWindowContainer3, "QSplashScreen_createWindowContainer3")
	symaddr(&addr_QSplashScreen_updateMicroFocus2, "QSplashScreen_updateMicroFocus2")
	symaddr(&addr_QSplashScreen_create2, "QSplashScreen_create2")
	symaddr(&addr_QSplashScreen_create3, "QSplashScreen_create3")
	symaddr(&addr_QSplashScreen_create4, "QSplashScreen_create4")
	symaddr(&addr_QSplashScreen_destroy2, "QSplashScreen_destroy2")
	symaddr(&addr_QSplashScreen_destroy3, "QSplashScreen_destroy3")
	symaddr(&addr_QSplashScreen_connect3, "QSplashScreen_connect3")
	symaddr(&addr_QSplashScreen_connect4, "QSplashScreen_connect4")
	symaddr(&addr_QSplashScreen_metaObject_default, "QSplashScreen_metaObject_default")
	symaddr(&addr_QSplashScreen_event_default, "QSplashScreen_event_default")
	symaddr(&addr_QSplashScreen_drawContents_default, "QSplashScreen_drawContents_default")
	symaddr(&addr_QSplashScreen_mousePressEvent_default, "QSplashScreen_mousePressEvent_default")
	symaddr(&addr_QSplashScreen_devType_default, "QSplashScreen_devType_default")
	symaddr(&addr_QSplashScreen_setVisible_default, "QSplashScreen_setVisible_default")
	symaddr(&addr_QSplashScreen_sizeHint_default, "QSplashScreen_sizeHint_default")
	symaddr(&addr_QSplashScreen_minimumSizeHint_default, "QSplashScreen_minimumSizeHint_default")
	symaddr(&addr_QSplashScreen_heightForWidth_default, "QSplashScreen_heightForWidth_default")
	symaddr(&addr_QSplashScreen_hasHeightForWidth_default, "QSplashScreen_hasHeightForWidth_default")
	symaddr(&addr_QSplashScreen_paintEngine_default, "QSplashScreen_paintEngine_default")
	symaddr(&addr_QSplashScreen_mouseReleaseEvent_default, "QSplashScreen_mouseReleaseEvent_default")
	symaddr(&addr_QSplashScreen_mouseDoubleClickEvent_default, "QSplashScreen_mouseDoubleClickEvent_default")
	symaddr(&addr_QSplashScreen_mouseMoveEvent_default, "QSplashScreen_mouseMoveEvent_default")
	symaddr(&addr_QSplashScreen_wheelEvent_default, "QSplashScreen_wheelEvent_default")
	symaddr(&addr_QSplashScreen_keyPressEvent_default, "QSplashScreen_keyPressEvent_default")
	symaddr(&addr_QSplashScreen_keyReleaseEvent_default, "QSplashScreen_keyReleaseEvent_default")
	symaddr(&addr_QSplashScreen_focusInEvent_default, "QSplashScreen_focusInEvent_default")
	symaddr(&addr_QSplashScreen_focusOutEvent_default, "QSplashScreen_focusOutEvent_default")
	symaddr(&addr_QSplashScreen_enterEvent_default, "QSplashScreen_enterEvent_default")
	symaddr(&addr_QSplashScreen_leaveEvent_default, "QSplashScreen_leaveEvent_default")
	symaddr(&addr_QSplashScreen_paintEvent_default, "QSplashScreen_paintEvent_default")
	symaddr(&addr_QSplashScreen_moveEvent_default, "QSplashScreen_moveEvent_default")
	symaddr(&addr_QSplashScreen_resizeEvent_default, "QSplashScreen_resizeEvent_default")
	symaddr(&addr_QSplashScreen_closeEvent_default, "QSplashScreen_closeEvent_default")
	symaddr(&addr_QSplashScreen_contextMenuEvent_default, "QSplashScreen_contextMenuEvent_default")
	symaddr(&addr_QSplashScreen_tabletEvent_default, "QSplashScreen_tabletEvent_default")
	symaddr(&addr_QSplashScreen_actionEvent_default, "QSplashScreen_actionEvent_default")
	symaddr(&addr_QSplashScreen_dragEnterEvent_default, "QSplashScreen_dragEnterEvent_default")
	symaddr(&addr_QSplashScreen_dragMoveEvent_default, "QSplashScreen_dragMoveEvent_default")
	symaddr(&addr_QSplashScreen_dragLeaveEvent_default, "QSplashScreen_dragLeaveEvent_default")
	symaddr(&addr_QSplashScreen_dropEvent_default, "QSplashScreen_dropEvent_default")
	symaddr(&addr_QSplashScreen_showEvent_default, "QSplashScreen_showEvent_default")
	symaddr(&addr_QSplashScreen_hideEvent_default, "QSplashScreen_hideEvent_default")
	symaddr(&addr_QSplashScreen_changeEvent_default, "QSplashScreen_changeEvent_default")
	symaddr(&addr_QSplashScreen_metric_default, "QSplashScreen_metric_default")
	symaddr(&addr_QSplashScreen_initPainter_default, "QSplashScreen_initPainter_default")
	symaddr(&addr_QSplashScreen_redirected_default, "QSplashScreen_redirected_default")
	symaddr(&addr_QSplashScreen_sharedPainter_default, "QSplashScreen_sharedPainter_default")
	symaddr(&addr_QSplashScreen_inputMethodEvent_default, "QSplashScreen_inputMethodEvent_default")
	symaddr(&addr_QSplashScreen_inputMethodQuery_default, "QSplashScreen_inputMethodQuery_default")
	symaddr(&addr_QSplashScreen_focusNextPrevChild_default, "QSplashScreen_focusNextPrevChild_default")
	symaddr(&addr_QSplashScreen_eventFilter_default, "QSplashScreen_eventFilter_default")
	symaddr(&addr_QSplashScreen_timerEvent_default, "QSplashScreen_timerEvent_default")
	symaddr(&addr_QSplashScreen_childEvent_default, "QSplashScreen_childEvent_default")
	symaddr(&addr_QSplashScreen_customEvent_default, "QSplashScreen_customEvent_default")
	symaddr(&addr_QSplashScreen_connectNotify_default, "QSplashScreen_connectNotify_default")
	symaddr(&addr_QSplashScreen_disconnectNotify_default, "QSplashScreen_disconnectNotify_default")
	reg(&pg_QSplashScreen_connect_messageChanged, "QSplashScreen_connect_messageChanged")
	reg(&pg_QSplashScreen_connect_windowTitleChanged, "QSplashScreen_connect_windowTitleChanged")
	reg(&pg_QSplashScreen_connect_windowIconChanged, "QSplashScreen_connect_windowIconChanged")
	reg(&pg_QSplashScreen_connect_windowIconTextChanged, "QSplashScreen_connect_windowIconTextChanged")
	reg(&pg_QSplashScreen_connect_customContextMenuRequested, "QSplashScreen_connect_customContextMenuRequested")
	reg(&pg_QSplashScreen_connect_destroyed, "QSplashScreen_connect_destroyed")
	reg(&pg_QSplashScreen_connect_destroyed2, "QSplashScreen_connect_destroyed2")
	reg(&pg_QSplashScreen_connect_objectNameChanged, "QSplashScreen_connect_objectNameChanged")
	reg(&pg_QSplashScreen__disconnect, "QSplashScreen__disconnect")
	reg(&pg_QSplashScreen__set_override, "QSplashScreen__set_override")
	reg(&pg_QSplashScreen_delete, "QSplashScreen_delete")
}

type QSplashScreen struct {
	h uintptr
	*QWidget
}

func bindQSplashScreen(h uintptr) *QSplashScreen {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QSplashScreen_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSplashScreen{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQSplashScreen(h uintptr) *QSplashScreen { return bindQSplashScreen(h) }

func (this *QSplashScreen) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSplashScreen) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSplashScreen(h unsafe.Pointer) *QSplashScreen { return bindQSplashScreen(uintptr(h)) }
func UnsafeNewQSplashScreenOwned(h unsafe.Pointer) *QSplashScreen { return newownedQSplashScreen(uintptr(h)) }

func (this *QSplashScreen) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQSplashScreen() *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new, uintptr(__goid))
	__o := newownedQSplashScreen(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func NewQSplashScreen2(screen *QScreen) *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new2, uintptr(__goid), uintptr(screen.UnsafePointer()))
	__o := newownedQSplashScreen(__r)
	runtime.KeepAlive(screen)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func NewQSplashScreen3(pixmap *QPixmap) *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new3, uintptr(__goid), uintptr(pixmap.UnsafePointer()))
	__o := newownedQSplashScreen(__r)
	runtime.KeepAlive(pixmap)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func NewQSplashScreen4(pixmap *QPixmap, f WindowType) *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new4, uintptr(__goid), uintptr(pixmap.UnsafePointer()), uintptr(f))
	__o := newownedQSplashScreen(__r)
	runtime.KeepAlive(pixmap)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func NewQSplashScreen5(screen *QScreen, pixmap *QPixmap) *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new5, uintptr(__goid), uintptr(screen.UnsafePointer()), uintptr(pixmap.UnsafePointer()))
	__o := newownedQSplashScreen(__r)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(pixmap)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func NewQSplashScreen6(screen *QScreen, pixmap *QPixmap, f WindowType) *QSplashScreen {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_new6, uintptr(__goid), uintptr(screen.UnsafePointer()), uintptr(pixmap.UnsafePointer()), uintptr(f))
	__o := newownedQSplashScreen(__r)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(pixmap)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplashScreen_virtIdx, QSplashScreen_enable)
	return __o
}

func QSplashScreen_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSplashScreen) SetPixmap(pixmap *QPixmap) {
	purego.SyscallN(addr_QSplashScreen_setPixmap, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QSplashScreen) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_pixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QSplashScreen) Finish(w *QWidget) {
	purego.SyscallN(addr_QSplashScreen_finish, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QSplashScreen) Repaint() {
	purego.SyscallN(addr_QSplashScreen_repaint, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Message() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_message, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSplashScreen) ShowMessage(message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QSplashScreen_showMessage, this.h, message_p, uintptr(len(message)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func (this *QSplashScreen) ClearMessage() {
	purego.SyscallN(addr_QSplashScreen_clearMessage, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) MessageChanged(message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QSplashScreen_messageChanged, this.h, message_p, uintptr(len(message)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func QSplashScreen_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSplashScreen_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSplashScreen_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_mouseGrabber)
	return bindQWidget(__r)
}

func QSplashScreen_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_keyboardGrabber)
	return bindQWidget(__r)
}

func QSplashScreen_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSplashScreen_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QSplashScreen) UpdateMicroFocus() {
	purego.SyscallN(addr_QSplashScreen_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Create() {
	purego.SyscallN(addr_QSplashScreen_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Destroy() {
	purego.SyscallN(addr_QSplashScreen_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplashScreen) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSplashScreen_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplashScreen_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSplashScreen) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSplashScreen) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplashScreen) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSplashScreen) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSplashScreen_DevicePixelRatioFScale() float64 {
	__r := pg_QSplashScreen_devicePixelRatioFScale()
	return __r
}

func QSplashScreen_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSplashScreen_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QSplashScreen) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QSplashScreen_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QSplashScreen_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSplashScreen_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSplashScreen) ShowMessage2(message string, alignment int) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QSplashScreen_showMessage2, this.h, message_p, uintptr(len(message)), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func (this *QSplashScreen) ShowMessage3(message string, alignment int, color *QColor) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QSplashScreen_showMessage3, this.h, message_p, uintptr(len(message)), uintptr(alignment), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
	runtime.KeepAlive(color)
}

func QSplashScreen_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSplashScreen_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QSplashScreen) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QSplashScreen_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QSplashScreen_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QSplashScreen_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QSplashScreen_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QSplashScreen_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QSplashScreen_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QSplashScreen_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplashScreen_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSplashScreen) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSplashScreen) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QSplashScreen) DrawContentsDefault(painter *QPainter) {
	purego.SyscallN(addr_QSplashScreen_drawContents_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QSplashScreen) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QSplashScreen_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplashScreen) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplashScreen) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QSplashScreen_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QSplashScreen) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSplashScreen) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSplashScreen) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplashScreen) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplashScreen) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QSplashScreen) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplashScreen_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplashScreen_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplashScreen_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QSplashScreen_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSplashScreen_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSplashScreen_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSplashScreen_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSplashScreen_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QSplashScreen_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSplashScreen_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QSplashScreen_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QSplashScreen_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QSplashScreen_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QSplashScreen_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QSplashScreen_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QSplashScreen_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QSplashScreen_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QSplashScreen_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QSplashScreen_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QSplashScreen_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QSplashScreen_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QSplashScreen_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QSplashScreen_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QSplashScreen_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplashScreen) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplashScreen) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QSplashScreen_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QSplashScreen) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QSplashScreen) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QSplashScreen) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QSplashScreen_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplashScreen) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSplashScreen_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSplashScreen) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplashScreen) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplashScreen_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSplashScreen) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSplashScreen_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSplashScreen_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSplashScreen_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplashScreen) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSplashScreen_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSplashScreen) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSplashScreen_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSplashScreen) OnMessageChanged(fn func(string)) int64 {
	id := registerCallback(func(message uintptr, message_len int64) { fn(pgStrNoFree(message, message_len)) })
	pg_QSplashScreen_connect_messageChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSplashScreen_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSplashScreen_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSplashScreen_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSplashScreen_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSplashScreen_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSplashScreen_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSplashScreen_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplashScreen) Disconnect(handle int64) {
	pg_QSplashScreen__disconnect(handle)
	unregisterCallback(handle)
}

func QSplashScreen_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case -2902866686677005435:
		return 2
	case -332090898078543319:
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
	case -4415920049297795213:
		return 11
	case -4029073905697023887:
		return 12
	case -4255494004099806581:
		return 13
	case 1919988340934511316:
		return 14
	case 1023791977679793185:
		return 15
	case 772598475950050299:
		return 16
	case 975043767033552552:
		return 17
	case -385596010576048993:
		return 18
	case 6316270346668237895:
		return 19
	case -551070009623711640:
		return 20
	case 5290442163167458465:
		return 21
	case -9161949150489525396:
		return 22
	case -4302917285655014661:
		return 23
	case -1989937050019464783:
		return 24
	case -6698043222227006015:
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

func QSplashScreen_enable(h uintptr, idx int32) { pg_QSplashScreen__set_override(h, idx) }

func (this *QSplashScreen) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSplashScreen) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QSplashScreen) OnDrawContents(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), -2902866686677005435, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "drawContents"
}

func (this *QSplashScreen) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QSplashScreen) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QSplashScreen) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QSplashScreen) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QSplashScreen) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QSplashScreen) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QSplashScreen) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QSplashScreen) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QSplashScreen) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QSplashScreen) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QSplashScreen) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QSplashScreen) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QSplashScreen) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QSplashScreen) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QSplashScreen) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QSplashScreen) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QSplashScreen) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QSplashScreen) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QSplashScreen) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QSplashScreen) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QSplashScreen) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QSplashScreen) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QSplashScreen) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QSplashScreen) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QSplashScreen) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QSplashScreen) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QSplashScreen) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QSplashScreen) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QSplashScreen) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QSplashScreen) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QSplashScreen) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QSplashScreen) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QSplashScreen) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QSplashScreen) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QSplashScreen) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QSplashScreen) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QSplashScreen) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QSplashScreen) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QSplashScreen) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QSplashScreen) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QSplashScreen) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSplashScreen) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSplashScreen) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSplashScreen) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSplashScreen) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSplashScreen) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSplashScreen) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSplashScreen_delete(this.h)
	}
	this.UnsafeRelease()
}
