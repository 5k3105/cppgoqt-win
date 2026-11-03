package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStackedWidget_chain uintptr
	addr_QStackedWidget_new uintptr
	addr_QStackedWidget_new2 uintptr
	addr_QStackedWidget_tr uintptr
	addr_QStackedWidget_addWidget uintptr
	addr_QStackedWidget_insertWidget uintptr
	addr_QStackedWidget_removeWidget uintptr
	addr_QStackedWidget_currentWidget uintptr
	addr_QStackedWidget_currentIndex uintptr
	addr_QStackedWidget_indexOf uintptr
	addr_QStackedWidget_widget uintptr
	addr_QStackedWidget_count uintptr
	addr_QStackedWidget_setCurrentIndex uintptr
	addr_QStackedWidget_setCurrentWidget uintptr
	addr_QStackedWidget_currentChanged uintptr
	addr_QStackedWidget_widgetRemoved uintptr
	addr_QStackedWidget_widgetAdded uintptr
	addr_QStackedWidget_drawFrame uintptr
	addr_QStackedWidget_setTabOrder uintptr
	addr_QStackedWidget_mouseGrabber uintptr
	addr_QStackedWidget_keyboardGrabber uintptr
	addr_QStackedWidget_find uintptr
	addr_QStackedWidget_createWindowContainer uintptr
	addr_QStackedWidget_updateMicroFocus uintptr
	addr_QStackedWidget_create uintptr
	addr_QStackedWidget_destroy uintptr
	addr_QStackedWidget_focusNextChild uintptr
	addr_QStackedWidget_focusPreviousChild uintptr
	addr_QStackedWidget_connect uintptr
	addr_QStackedWidget_connect2 uintptr
	addr_QStackedWidget_sender uintptr
	addr_QStackedWidget_senderSignalIndex uintptr
	addr_QStackedWidget_receivers uintptr
	addr_QStackedWidget_isSignalConnected uintptr
	pg_QStackedWidget_devicePixelRatioFScale func() float64
	pg_QStackedWidget_encodeMetricF func(int32, float64) int32
	pg_QStackedWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QStackedWidget_tr2 uintptr
	addr_QStackedWidget_tr3 uintptr
	addr_QStackedWidget_createWindowContainer2 uintptr
	addr_QStackedWidget_createWindowContainer3 uintptr
	addr_QStackedWidget_updateMicroFocus2 uintptr
	addr_QStackedWidget_create2 uintptr
	addr_QStackedWidget_create3 uintptr
	addr_QStackedWidget_create4 uintptr
	addr_QStackedWidget_destroy2 uintptr
	addr_QStackedWidget_destroy3 uintptr
	addr_QStackedWidget_connect3 uintptr
	addr_QStackedWidget_connect4 uintptr
	addr_QStackedWidget_metaObject_default uintptr
	addr_QStackedWidget_event_default uintptr
	addr_QStackedWidget_sizeHint_default uintptr
	addr_QStackedWidget_paintEvent_default uintptr
	addr_QStackedWidget_changeEvent_default uintptr
	addr_QStackedWidget_initStyleOption_default uintptr
	addr_QStackedWidget_devType_default uintptr
	addr_QStackedWidget_setVisible_default uintptr
	addr_QStackedWidget_minimumSizeHint_default uintptr
	addr_QStackedWidget_heightForWidth_default uintptr
	addr_QStackedWidget_hasHeightForWidth_default uintptr
	addr_QStackedWidget_paintEngine_default uintptr
	addr_QStackedWidget_mousePressEvent_default uintptr
	addr_QStackedWidget_mouseReleaseEvent_default uintptr
	addr_QStackedWidget_mouseDoubleClickEvent_default uintptr
	addr_QStackedWidget_mouseMoveEvent_default uintptr
	addr_QStackedWidget_wheelEvent_default uintptr
	addr_QStackedWidget_keyPressEvent_default uintptr
	addr_QStackedWidget_keyReleaseEvent_default uintptr
	addr_QStackedWidget_focusInEvent_default uintptr
	addr_QStackedWidget_focusOutEvent_default uintptr
	addr_QStackedWidget_enterEvent_default uintptr
	addr_QStackedWidget_leaveEvent_default uintptr
	addr_QStackedWidget_moveEvent_default uintptr
	addr_QStackedWidget_resizeEvent_default uintptr
	addr_QStackedWidget_closeEvent_default uintptr
	addr_QStackedWidget_contextMenuEvent_default uintptr
	addr_QStackedWidget_tabletEvent_default uintptr
	addr_QStackedWidget_actionEvent_default uintptr
	addr_QStackedWidget_dragEnterEvent_default uintptr
	addr_QStackedWidget_dragMoveEvent_default uintptr
	addr_QStackedWidget_dragLeaveEvent_default uintptr
	addr_QStackedWidget_dropEvent_default uintptr
	addr_QStackedWidget_showEvent_default uintptr
	addr_QStackedWidget_hideEvent_default uintptr
	addr_QStackedWidget_metric_default uintptr
	addr_QStackedWidget_initPainter_default uintptr
	addr_QStackedWidget_redirected_default uintptr
	addr_QStackedWidget_sharedPainter_default uintptr
	addr_QStackedWidget_inputMethodEvent_default uintptr
	addr_QStackedWidget_inputMethodQuery_default uintptr
	addr_QStackedWidget_focusNextPrevChild_default uintptr
	addr_QStackedWidget_eventFilter_default uintptr
	addr_QStackedWidget_timerEvent_default uintptr
	addr_QStackedWidget_childEvent_default uintptr
	addr_QStackedWidget_customEvent_default uintptr
	addr_QStackedWidget_connectNotify_default uintptr
	addr_QStackedWidget_disconnectNotify_default uintptr
	pg_QStackedWidget_connect_currentChanged func(uintptr, int64)
	pg_QStackedWidget_connect_widgetRemoved func(uintptr, int64)
	pg_QStackedWidget_connect_widgetAdded func(uintptr, int64)
	pg_QStackedWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QStackedWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QStackedWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QStackedWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QStackedWidget_connect_destroyed func(uintptr, int64)
	pg_QStackedWidget_connect_destroyed2 func(uintptr, int64)
	pg_QStackedWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QStackedWidget__disconnect func(int64)
	pg_QStackedWidget__set_override func(uintptr, int32)
	pg_QStackedWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStackedWidget_chain, "QStackedWidget_chain")
	symaddr(&addr_QStackedWidget_new, "QStackedWidget_new")
	symaddr(&addr_QStackedWidget_new2, "QStackedWidget_new2")
	symaddr(&addr_QStackedWidget_tr, "QStackedWidget_tr")
	symaddr(&addr_QStackedWidget_addWidget, "QStackedWidget_addWidget")
	symaddr(&addr_QStackedWidget_insertWidget, "QStackedWidget_insertWidget")
	symaddr(&addr_QStackedWidget_removeWidget, "QStackedWidget_removeWidget")
	symaddr(&addr_QStackedWidget_currentWidget, "QStackedWidget_currentWidget")
	symaddr(&addr_QStackedWidget_currentIndex, "QStackedWidget_currentIndex")
	symaddr(&addr_QStackedWidget_indexOf, "QStackedWidget_indexOf")
	symaddr(&addr_QStackedWidget_widget, "QStackedWidget_widget")
	symaddr(&addr_QStackedWidget_count, "QStackedWidget_count")
	symaddr(&addr_QStackedWidget_setCurrentIndex, "QStackedWidget_setCurrentIndex")
	symaddr(&addr_QStackedWidget_setCurrentWidget, "QStackedWidget_setCurrentWidget")
	symaddr(&addr_QStackedWidget_currentChanged, "QStackedWidget_currentChanged")
	symaddr(&addr_QStackedWidget_widgetRemoved, "QStackedWidget_widgetRemoved")
	symaddr(&addr_QStackedWidget_widgetAdded, "QStackedWidget_widgetAdded")
	symaddr(&addr_QStackedWidget_drawFrame, "QStackedWidget_drawFrame")
	symaddr(&addr_QStackedWidget_setTabOrder, "QStackedWidget_setTabOrder")
	symaddr(&addr_QStackedWidget_mouseGrabber, "QStackedWidget_mouseGrabber")
	symaddr(&addr_QStackedWidget_keyboardGrabber, "QStackedWidget_keyboardGrabber")
	symaddr(&addr_QStackedWidget_find, "QStackedWidget_find")
	symaddr(&addr_QStackedWidget_createWindowContainer, "QStackedWidget_createWindowContainer")
	symaddr(&addr_QStackedWidget_updateMicroFocus, "QStackedWidget_updateMicroFocus")
	symaddr(&addr_QStackedWidget_create, "QStackedWidget_create")
	symaddr(&addr_QStackedWidget_destroy, "QStackedWidget_destroy")
	symaddr(&addr_QStackedWidget_focusNextChild, "QStackedWidget_focusNextChild")
	symaddr(&addr_QStackedWidget_focusPreviousChild, "QStackedWidget_focusPreviousChild")
	symaddr(&addr_QStackedWidget_connect, "QStackedWidget_connect")
	symaddr(&addr_QStackedWidget_connect2, "QStackedWidget_connect2")
	symaddr(&addr_QStackedWidget_sender, "QStackedWidget_sender")
	symaddr(&addr_QStackedWidget_senderSignalIndex, "QStackedWidget_senderSignalIndex")
	symaddr(&addr_QStackedWidget_receivers, "QStackedWidget_receivers")
	symaddr(&addr_QStackedWidget_isSignalConnected, "QStackedWidget_isSignalConnected")
	reg(&pg_QStackedWidget_devicePixelRatioFScale, "QStackedWidget_devicePixelRatioFScale")
	reg(&pg_QStackedWidget_encodeMetricF, "QStackedWidget_encodeMetricF")
	reg(&pg_QStackedWidget_getDecodedMetricF, "QStackedWidget_getDecodedMetricF")
	symaddr(&addr_QStackedWidget_tr2, "QStackedWidget_tr2")
	symaddr(&addr_QStackedWidget_tr3, "QStackedWidget_tr3")
	symaddr(&addr_QStackedWidget_createWindowContainer2, "QStackedWidget_createWindowContainer2")
	symaddr(&addr_QStackedWidget_createWindowContainer3, "QStackedWidget_createWindowContainer3")
	symaddr(&addr_QStackedWidget_updateMicroFocus2, "QStackedWidget_updateMicroFocus2")
	symaddr(&addr_QStackedWidget_create2, "QStackedWidget_create2")
	symaddr(&addr_QStackedWidget_create3, "QStackedWidget_create3")
	symaddr(&addr_QStackedWidget_create4, "QStackedWidget_create4")
	symaddr(&addr_QStackedWidget_destroy2, "QStackedWidget_destroy2")
	symaddr(&addr_QStackedWidget_destroy3, "QStackedWidget_destroy3")
	symaddr(&addr_QStackedWidget_connect3, "QStackedWidget_connect3")
	symaddr(&addr_QStackedWidget_connect4, "QStackedWidget_connect4")
	symaddr(&addr_QStackedWidget_metaObject_default, "QStackedWidget_metaObject_default")
	symaddr(&addr_QStackedWidget_event_default, "QStackedWidget_event_default")
	symaddr(&addr_QStackedWidget_sizeHint_default, "QStackedWidget_sizeHint_default")
	symaddr(&addr_QStackedWidget_paintEvent_default, "QStackedWidget_paintEvent_default")
	symaddr(&addr_QStackedWidget_changeEvent_default, "QStackedWidget_changeEvent_default")
	symaddr(&addr_QStackedWidget_initStyleOption_default, "QStackedWidget_initStyleOption_default")
	symaddr(&addr_QStackedWidget_devType_default, "QStackedWidget_devType_default")
	symaddr(&addr_QStackedWidget_setVisible_default, "QStackedWidget_setVisible_default")
	symaddr(&addr_QStackedWidget_minimumSizeHint_default, "QStackedWidget_minimumSizeHint_default")
	symaddr(&addr_QStackedWidget_heightForWidth_default, "QStackedWidget_heightForWidth_default")
	symaddr(&addr_QStackedWidget_hasHeightForWidth_default, "QStackedWidget_hasHeightForWidth_default")
	symaddr(&addr_QStackedWidget_paintEngine_default, "QStackedWidget_paintEngine_default")
	symaddr(&addr_QStackedWidget_mousePressEvent_default, "QStackedWidget_mousePressEvent_default")
	symaddr(&addr_QStackedWidget_mouseReleaseEvent_default, "QStackedWidget_mouseReleaseEvent_default")
	symaddr(&addr_QStackedWidget_mouseDoubleClickEvent_default, "QStackedWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QStackedWidget_mouseMoveEvent_default, "QStackedWidget_mouseMoveEvent_default")
	symaddr(&addr_QStackedWidget_wheelEvent_default, "QStackedWidget_wheelEvent_default")
	symaddr(&addr_QStackedWidget_keyPressEvent_default, "QStackedWidget_keyPressEvent_default")
	symaddr(&addr_QStackedWidget_keyReleaseEvent_default, "QStackedWidget_keyReleaseEvent_default")
	symaddr(&addr_QStackedWidget_focusInEvent_default, "QStackedWidget_focusInEvent_default")
	symaddr(&addr_QStackedWidget_focusOutEvent_default, "QStackedWidget_focusOutEvent_default")
	symaddr(&addr_QStackedWidget_enterEvent_default, "QStackedWidget_enterEvent_default")
	symaddr(&addr_QStackedWidget_leaveEvent_default, "QStackedWidget_leaveEvent_default")
	symaddr(&addr_QStackedWidget_moveEvent_default, "QStackedWidget_moveEvent_default")
	symaddr(&addr_QStackedWidget_resizeEvent_default, "QStackedWidget_resizeEvent_default")
	symaddr(&addr_QStackedWidget_closeEvent_default, "QStackedWidget_closeEvent_default")
	symaddr(&addr_QStackedWidget_contextMenuEvent_default, "QStackedWidget_contextMenuEvent_default")
	symaddr(&addr_QStackedWidget_tabletEvent_default, "QStackedWidget_tabletEvent_default")
	symaddr(&addr_QStackedWidget_actionEvent_default, "QStackedWidget_actionEvent_default")
	symaddr(&addr_QStackedWidget_dragEnterEvent_default, "QStackedWidget_dragEnterEvent_default")
	symaddr(&addr_QStackedWidget_dragMoveEvent_default, "QStackedWidget_dragMoveEvent_default")
	symaddr(&addr_QStackedWidget_dragLeaveEvent_default, "QStackedWidget_dragLeaveEvent_default")
	symaddr(&addr_QStackedWidget_dropEvent_default, "QStackedWidget_dropEvent_default")
	symaddr(&addr_QStackedWidget_showEvent_default, "QStackedWidget_showEvent_default")
	symaddr(&addr_QStackedWidget_hideEvent_default, "QStackedWidget_hideEvent_default")
	symaddr(&addr_QStackedWidget_metric_default, "QStackedWidget_metric_default")
	symaddr(&addr_QStackedWidget_initPainter_default, "QStackedWidget_initPainter_default")
	symaddr(&addr_QStackedWidget_redirected_default, "QStackedWidget_redirected_default")
	symaddr(&addr_QStackedWidget_sharedPainter_default, "QStackedWidget_sharedPainter_default")
	symaddr(&addr_QStackedWidget_inputMethodEvent_default, "QStackedWidget_inputMethodEvent_default")
	symaddr(&addr_QStackedWidget_inputMethodQuery_default, "QStackedWidget_inputMethodQuery_default")
	symaddr(&addr_QStackedWidget_focusNextPrevChild_default, "QStackedWidget_focusNextPrevChild_default")
	symaddr(&addr_QStackedWidget_eventFilter_default, "QStackedWidget_eventFilter_default")
	symaddr(&addr_QStackedWidget_timerEvent_default, "QStackedWidget_timerEvent_default")
	symaddr(&addr_QStackedWidget_childEvent_default, "QStackedWidget_childEvent_default")
	symaddr(&addr_QStackedWidget_customEvent_default, "QStackedWidget_customEvent_default")
	symaddr(&addr_QStackedWidget_connectNotify_default, "QStackedWidget_connectNotify_default")
	symaddr(&addr_QStackedWidget_disconnectNotify_default, "QStackedWidget_disconnectNotify_default")
	reg(&pg_QStackedWidget_connect_currentChanged, "QStackedWidget_connect_currentChanged")
	reg(&pg_QStackedWidget_connect_widgetRemoved, "QStackedWidget_connect_widgetRemoved")
	reg(&pg_QStackedWidget_connect_widgetAdded, "QStackedWidget_connect_widgetAdded")
	reg(&pg_QStackedWidget_connect_windowTitleChanged, "QStackedWidget_connect_windowTitleChanged")
	reg(&pg_QStackedWidget_connect_windowIconChanged, "QStackedWidget_connect_windowIconChanged")
	reg(&pg_QStackedWidget_connect_windowIconTextChanged, "QStackedWidget_connect_windowIconTextChanged")
	reg(&pg_QStackedWidget_connect_customContextMenuRequested, "QStackedWidget_connect_customContextMenuRequested")
	reg(&pg_QStackedWidget_connect_destroyed, "QStackedWidget_connect_destroyed")
	reg(&pg_QStackedWidget_connect_destroyed2, "QStackedWidget_connect_destroyed2")
	reg(&pg_QStackedWidget_connect_objectNameChanged, "QStackedWidget_connect_objectNameChanged")
	reg(&pg_QStackedWidget__disconnect, "QStackedWidget__disconnect")
	reg(&pg_QStackedWidget__set_override, "QStackedWidget__set_override")
	reg(&pg_QStackedWidget_delete, "QStackedWidget_delete")
}

type QStackedWidget struct {
	h uintptr
	*QFrame
}

func bindQStackedWidget(h uintptr) *QStackedWidget {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QStackedWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStackedWidget{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQStackedWidget(h uintptr) *QStackedWidget { return bindQStackedWidget(h) }

func (this *QStackedWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStackedWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStackedWidget(h unsafe.Pointer) *QStackedWidget { return bindQStackedWidget(uintptr(h)) }
func UnsafeNewQStackedWidgetOwned(h unsafe.Pointer) *QStackedWidget { return newownedQStackedWidget(uintptr(h)) }

func (this *QStackedWidget) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQStackedWidget(parent *QWidget) *QStackedWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStackedWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStackedWidget_virtIdx, QStackedWidget_enable)
	return __o
}

func NewQStackedWidget2() *QStackedWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_new2, uintptr(__goid))
	__o := newownedQStackedWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStackedWidget_virtIdx, QStackedWidget_enable)
	return __o
}

func QStackedWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStackedWidget) AddWidget(w *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_addWidget, this.h, uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return int(__r)
}

func (this *QStackedWidget) InsertWidget(index int, w *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_insertWidget, this.h, uintptr(index), uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return int(__r)
}

func (this *QStackedWidget) RemoveWidget(w *QWidget) {
	purego.SyscallN(addr_QStackedWidget_removeWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QStackedWidget) CurrentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_currentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QStackedWidget) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) IndexOf(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_indexOf, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QStackedWidget) Widget(arg1 int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_widget, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QStackedWidget) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QStackedWidget_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) SetCurrentWidget(w *QWidget) {
	purego.SyscallN(addr_QStackedWidget_setCurrentWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QStackedWidget) CurrentChanged(arg1 int) {
	purego.SyscallN(addr_QStackedWidget_currentChanged, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) WidgetRemoved(index int) {
	purego.SyscallN(addr_QStackedWidget_widgetRemoved, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) WidgetAdded(index int) {
	purego.SyscallN(addr_QStackedWidget_widgetAdded, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QStackedWidget_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QStackedWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QStackedWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QStackedWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QStackedWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QStackedWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QStackedWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QStackedWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QStackedWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) Create() {
	purego.SyscallN(addr_QStackedWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) Destroy() {
	purego.SyscallN(addr_QStackedWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStackedWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QStackedWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStackedWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStackedWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStackedWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStackedWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStackedWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QStackedWidget_devicePixelRatioFScale()
	return __r
}

func QStackedWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QStackedWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QStackedWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QStackedWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QStackedWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStackedWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStackedWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QStackedWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QStackedWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QStackedWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QStackedWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QStackedWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QStackedWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QStackedWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QStackedWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QStackedWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStackedWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStackedWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStackedWidget) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QStackedWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStackedWidget) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QStackedWidget_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStackedWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QStackedWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStackedWidget) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QStackedWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QStackedWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QStackedWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QStackedWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStackedWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStackedWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QStackedWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStackedWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStackedWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStackedWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStackedWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QStackedWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QStackedWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QStackedWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QStackedWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QStackedWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QStackedWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStackedWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QStackedWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QStackedWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QStackedWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QStackedWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QStackedWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QStackedWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QStackedWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QStackedWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QStackedWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QStackedWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QStackedWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QStackedWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QStackedWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QStackedWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QStackedWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QStackedWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QStackedWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStackedWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStackedWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStackedWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStackedWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStackedWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStackedWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStackedWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStackedWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStackedWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStackedWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStackedWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStackedWidget) OnCurrentChanged(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QStackedWidget_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnWidgetRemoved(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QStackedWidget_connect_widgetRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnWidgetAdded(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QStackedWidget_connect_widgetAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QStackedWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QStackedWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QStackedWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QStackedWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStackedWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStackedWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStackedWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedWidget) Disconnect(handle int64) {
	pg_QStackedWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QStackedWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 729690261457684101:
		return 2
	case 5290442163167458465:
		return 3
	case 4959376063826165943:
		return 4
	case -488693769771308531:
		return 5
	case -8659679893868524044:
		return 6
	case 2110182719375245575:
		return 7
	case 7049264599837532491:
		return 8
	case 4311534060495737339:
		return 9
	case -8266762397995423915:
		return 10
	case 6398007404756346381:
		return 11
	case -332090898078543319:
		return 12
	case -4415920049297795213:
		return 13
	case -4029073905697023887:
		return 14
	case -4255494004099806581:
		return 15
	case 1919988340934511316:
		return 16
	case 1023791977679793185:
		return 17
	case 772598475950050299:
		return 18
	case 975043767033552552:
		return 19
	case -385596010576048993:
		return 20
	case 6316270346668237895:
		return 21
	case -551070009623711640:
		return 22
	case -9161949150489525396:
		return 23
	case -4302917285655014661:
		return 24
	case -1989937050019464783:
		return 25
	case -6698043222227006015:
		return 26
	case -212068866860701141:
		return 27
	case 1697444140088555581:
		return 28
	case 206122715919917589:
		return 29
	case -1781545137352745978:
		return 30
	case -3139904158373725006:
		return 31
	case 9107058682754591716:
		return 32
	case -6825369174751916:
		return 33
	case -199718707876719129:
		return 34
	case -2039166586246443706:
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

func QStackedWidget_enable(h uintptr, idx int32) { pg_QStackedWidget__set_override(h, idx) }

func (this *QStackedWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStackedWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QStackedWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QStackedWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QStackedWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QStackedWidget) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QStackedWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QStackedWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QStackedWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QStackedWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QStackedWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QStackedWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QStackedWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QStackedWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QStackedWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QStackedWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QStackedWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QStackedWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QStackedWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QStackedWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QStackedWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QStackedWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QStackedWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QStackedWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QStackedWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QStackedWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QStackedWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QStackedWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QStackedWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QStackedWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QStackedWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QStackedWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QStackedWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QStackedWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QStackedWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QStackedWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QStackedWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QStackedWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QStackedWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QStackedWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QStackedWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QStackedWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QStackedWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QStackedWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStackedWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStackedWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStackedWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStackedWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStackedWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStackedWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStackedWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
