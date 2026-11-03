package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStatusBar_chain uintptr
	addr_QStatusBar_new uintptr
	addr_QStatusBar_new2 uintptr
	addr_QStatusBar_tr uintptr
	addr_QStatusBar_addWidget uintptr
	addr_QStatusBar_insertWidget uintptr
	addr_QStatusBar_addPermanentWidget uintptr
	addr_QStatusBar_insertPermanentWidget uintptr
	addr_QStatusBar_removeWidget uintptr
	addr_QStatusBar_setSizeGripEnabled uintptr
	addr_QStatusBar_isSizeGripEnabled uintptr
	addr_QStatusBar_currentMessage uintptr
	addr_QStatusBar_showMessage uintptr
	addr_QStatusBar_clearMessage uintptr
	addr_QStatusBar_messageChanged uintptr
	addr_QStatusBar_reformat uintptr
	addr_QStatusBar_hideOrShow uintptr
	addr_QStatusBar_setTabOrder uintptr
	addr_QStatusBar_mouseGrabber uintptr
	addr_QStatusBar_keyboardGrabber uintptr
	addr_QStatusBar_find uintptr
	addr_QStatusBar_createWindowContainer uintptr
	addr_QStatusBar_updateMicroFocus uintptr
	addr_QStatusBar_create uintptr
	addr_QStatusBar_destroy uintptr
	addr_QStatusBar_focusNextChild uintptr
	addr_QStatusBar_focusPreviousChild uintptr
	addr_QStatusBar_connect uintptr
	addr_QStatusBar_connect2 uintptr
	addr_QStatusBar_sender uintptr
	addr_QStatusBar_senderSignalIndex uintptr
	addr_QStatusBar_receivers uintptr
	addr_QStatusBar_isSignalConnected uintptr
	pg_QStatusBar_devicePixelRatioFScale func() float64
	pg_QStatusBar_encodeMetricF func(int32, float64) int32
	pg_QStatusBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QStatusBar_tr2 uintptr
	addr_QStatusBar_tr3 uintptr
	addr_QStatusBar_addWidget2 uintptr
	addr_QStatusBar_insertWidget2 uintptr
	addr_QStatusBar_addPermanentWidget2 uintptr
	addr_QStatusBar_insertPermanentWidget2 uintptr
	addr_QStatusBar_showMessage2 uintptr
	addr_QStatusBar_createWindowContainer2 uintptr
	addr_QStatusBar_createWindowContainer3 uintptr
	addr_QStatusBar_updateMicroFocus2 uintptr
	addr_QStatusBar_create2 uintptr
	addr_QStatusBar_create3 uintptr
	addr_QStatusBar_create4 uintptr
	addr_QStatusBar_destroy2 uintptr
	addr_QStatusBar_destroy3 uintptr
	addr_QStatusBar_connect3 uintptr
	addr_QStatusBar_connect4 uintptr
	addr_QStatusBar_metaObject_default uintptr
	addr_QStatusBar_showEvent_default uintptr
	addr_QStatusBar_paintEvent_default uintptr
	addr_QStatusBar_resizeEvent_default uintptr
	addr_QStatusBar_event_default uintptr
	addr_QStatusBar_devType_default uintptr
	addr_QStatusBar_setVisible_default uintptr
	addr_QStatusBar_sizeHint_default uintptr
	addr_QStatusBar_minimumSizeHint_default uintptr
	addr_QStatusBar_heightForWidth_default uintptr
	addr_QStatusBar_hasHeightForWidth_default uintptr
	addr_QStatusBar_paintEngine_default uintptr
	addr_QStatusBar_mousePressEvent_default uintptr
	addr_QStatusBar_mouseReleaseEvent_default uintptr
	addr_QStatusBar_mouseDoubleClickEvent_default uintptr
	addr_QStatusBar_mouseMoveEvent_default uintptr
	addr_QStatusBar_wheelEvent_default uintptr
	addr_QStatusBar_keyPressEvent_default uintptr
	addr_QStatusBar_keyReleaseEvent_default uintptr
	addr_QStatusBar_focusInEvent_default uintptr
	addr_QStatusBar_focusOutEvent_default uintptr
	addr_QStatusBar_enterEvent_default uintptr
	addr_QStatusBar_leaveEvent_default uintptr
	addr_QStatusBar_moveEvent_default uintptr
	addr_QStatusBar_closeEvent_default uintptr
	addr_QStatusBar_contextMenuEvent_default uintptr
	addr_QStatusBar_tabletEvent_default uintptr
	addr_QStatusBar_actionEvent_default uintptr
	addr_QStatusBar_dragEnterEvent_default uintptr
	addr_QStatusBar_dragMoveEvent_default uintptr
	addr_QStatusBar_dragLeaveEvent_default uintptr
	addr_QStatusBar_dropEvent_default uintptr
	addr_QStatusBar_hideEvent_default uintptr
	addr_QStatusBar_changeEvent_default uintptr
	addr_QStatusBar_metric_default uintptr
	addr_QStatusBar_initPainter_default uintptr
	addr_QStatusBar_redirected_default uintptr
	addr_QStatusBar_sharedPainter_default uintptr
	addr_QStatusBar_inputMethodEvent_default uintptr
	addr_QStatusBar_inputMethodQuery_default uintptr
	addr_QStatusBar_focusNextPrevChild_default uintptr
	addr_QStatusBar_eventFilter_default uintptr
	addr_QStatusBar_timerEvent_default uintptr
	addr_QStatusBar_childEvent_default uintptr
	addr_QStatusBar_customEvent_default uintptr
	addr_QStatusBar_connectNotify_default uintptr
	addr_QStatusBar_disconnectNotify_default uintptr
	pg_QStatusBar_connect_messageChanged func(uintptr, int64)
	pg_QStatusBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QStatusBar_connect_windowIconChanged func(uintptr, int64)
	pg_QStatusBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QStatusBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QStatusBar_connect_destroyed func(uintptr, int64)
	pg_QStatusBar_connect_destroyed2 func(uintptr, int64)
	pg_QStatusBar_connect_objectNameChanged func(uintptr, int64)
	pg_QStatusBar__disconnect func(int64)
	pg_QStatusBar__set_override func(uintptr, int32)
	pg_QStatusBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStatusBar_chain, "QStatusBar_chain")
	symaddr(&addr_QStatusBar_new, "QStatusBar_new")
	symaddr(&addr_QStatusBar_new2, "QStatusBar_new2")
	symaddr(&addr_QStatusBar_tr, "QStatusBar_tr")
	symaddr(&addr_QStatusBar_addWidget, "QStatusBar_addWidget")
	symaddr(&addr_QStatusBar_insertWidget, "QStatusBar_insertWidget")
	symaddr(&addr_QStatusBar_addPermanentWidget, "QStatusBar_addPermanentWidget")
	symaddr(&addr_QStatusBar_insertPermanentWidget, "QStatusBar_insertPermanentWidget")
	symaddr(&addr_QStatusBar_removeWidget, "QStatusBar_removeWidget")
	symaddr(&addr_QStatusBar_setSizeGripEnabled, "QStatusBar_setSizeGripEnabled")
	symaddr(&addr_QStatusBar_isSizeGripEnabled, "QStatusBar_isSizeGripEnabled")
	symaddr(&addr_QStatusBar_currentMessage, "QStatusBar_currentMessage")
	symaddr(&addr_QStatusBar_showMessage, "QStatusBar_showMessage")
	symaddr(&addr_QStatusBar_clearMessage, "QStatusBar_clearMessage")
	symaddr(&addr_QStatusBar_messageChanged, "QStatusBar_messageChanged")
	symaddr(&addr_QStatusBar_reformat, "QStatusBar_reformat")
	symaddr(&addr_QStatusBar_hideOrShow, "QStatusBar_hideOrShow")
	symaddr(&addr_QStatusBar_setTabOrder, "QStatusBar_setTabOrder")
	symaddr(&addr_QStatusBar_mouseGrabber, "QStatusBar_mouseGrabber")
	symaddr(&addr_QStatusBar_keyboardGrabber, "QStatusBar_keyboardGrabber")
	symaddr(&addr_QStatusBar_find, "QStatusBar_find")
	symaddr(&addr_QStatusBar_createWindowContainer, "QStatusBar_createWindowContainer")
	symaddr(&addr_QStatusBar_updateMicroFocus, "QStatusBar_updateMicroFocus")
	symaddr(&addr_QStatusBar_create, "QStatusBar_create")
	symaddr(&addr_QStatusBar_destroy, "QStatusBar_destroy")
	symaddr(&addr_QStatusBar_focusNextChild, "QStatusBar_focusNextChild")
	symaddr(&addr_QStatusBar_focusPreviousChild, "QStatusBar_focusPreviousChild")
	symaddr(&addr_QStatusBar_connect, "QStatusBar_connect")
	symaddr(&addr_QStatusBar_connect2, "QStatusBar_connect2")
	symaddr(&addr_QStatusBar_sender, "QStatusBar_sender")
	symaddr(&addr_QStatusBar_senderSignalIndex, "QStatusBar_senderSignalIndex")
	symaddr(&addr_QStatusBar_receivers, "QStatusBar_receivers")
	symaddr(&addr_QStatusBar_isSignalConnected, "QStatusBar_isSignalConnected")
	reg(&pg_QStatusBar_devicePixelRatioFScale, "QStatusBar_devicePixelRatioFScale")
	reg(&pg_QStatusBar_encodeMetricF, "QStatusBar_encodeMetricF")
	reg(&pg_QStatusBar_getDecodedMetricF, "QStatusBar_getDecodedMetricF")
	symaddr(&addr_QStatusBar_tr2, "QStatusBar_tr2")
	symaddr(&addr_QStatusBar_tr3, "QStatusBar_tr3")
	symaddr(&addr_QStatusBar_addWidget2, "QStatusBar_addWidget2")
	symaddr(&addr_QStatusBar_insertWidget2, "QStatusBar_insertWidget2")
	symaddr(&addr_QStatusBar_addPermanentWidget2, "QStatusBar_addPermanentWidget2")
	symaddr(&addr_QStatusBar_insertPermanentWidget2, "QStatusBar_insertPermanentWidget2")
	symaddr(&addr_QStatusBar_showMessage2, "QStatusBar_showMessage2")
	symaddr(&addr_QStatusBar_createWindowContainer2, "QStatusBar_createWindowContainer2")
	symaddr(&addr_QStatusBar_createWindowContainer3, "QStatusBar_createWindowContainer3")
	symaddr(&addr_QStatusBar_updateMicroFocus2, "QStatusBar_updateMicroFocus2")
	symaddr(&addr_QStatusBar_create2, "QStatusBar_create2")
	symaddr(&addr_QStatusBar_create3, "QStatusBar_create3")
	symaddr(&addr_QStatusBar_create4, "QStatusBar_create4")
	symaddr(&addr_QStatusBar_destroy2, "QStatusBar_destroy2")
	symaddr(&addr_QStatusBar_destroy3, "QStatusBar_destroy3")
	symaddr(&addr_QStatusBar_connect3, "QStatusBar_connect3")
	symaddr(&addr_QStatusBar_connect4, "QStatusBar_connect4")
	symaddr(&addr_QStatusBar_metaObject_default, "QStatusBar_metaObject_default")
	symaddr(&addr_QStatusBar_showEvent_default, "QStatusBar_showEvent_default")
	symaddr(&addr_QStatusBar_paintEvent_default, "QStatusBar_paintEvent_default")
	symaddr(&addr_QStatusBar_resizeEvent_default, "QStatusBar_resizeEvent_default")
	symaddr(&addr_QStatusBar_event_default, "QStatusBar_event_default")
	symaddr(&addr_QStatusBar_devType_default, "QStatusBar_devType_default")
	symaddr(&addr_QStatusBar_setVisible_default, "QStatusBar_setVisible_default")
	symaddr(&addr_QStatusBar_sizeHint_default, "QStatusBar_sizeHint_default")
	symaddr(&addr_QStatusBar_minimumSizeHint_default, "QStatusBar_minimumSizeHint_default")
	symaddr(&addr_QStatusBar_heightForWidth_default, "QStatusBar_heightForWidth_default")
	symaddr(&addr_QStatusBar_hasHeightForWidth_default, "QStatusBar_hasHeightForWidth_default")
	symaddr(&addr_QStatusBar_paintEngine_default, "QStatusBar_paintEngine_default")
	symaddr(&addr_QStatusBar_mousePressEvent_default, "QStatusBar_mousePressEvent_default")
	symaddr(&addr_QStatusBar_mouseReleaseEvent_default, "QStatusBar_mouseReleaseEvent_default")
	symaddr(&addr_QStatusBar_mouseDoubleClickEvent_default, "QStatusBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QStatusBar_mouseMoveEvent_default, "QStatusBar_mouseMoveEvent_default")
	symaddr(&addr_QStatusBar_wheelEvent_default, "QStatusBar_wheelEvent_default")
	symaddr(&addr_QStatusBar_keyPressEvent_default, "QStatusBar_keyPressEvent_default")
	symaddr(&addr_QStatusBar_keyReleaseEvent_default, "QStatusBar_keyReleaseEvent_default")
	symaddr(&addr_QStatusBar_focusInEvent_default, "QStatusBar_focusInEvent_default")
	symaddr(&addr_QStatusBar_focusOutEvent_default, "QStatusBar_focusOutEvent_default")
	symaddr(&addr_QStatusBar_enterEvent_default, "QStatusBar_enterEvent_default")
	symaddr(&addr_QStatusBar_leaveEvent_default, "QStatusBar_leaveEvent_default")
	symaddr(&addr_QStatusBar_moveEvent_default, "QStatusBar_moveEvent_default")
	symaddr(&addr_QStatusBar_closeEvent_default, "QStatusBar_closeEvent_default")
	symaddr(&addr_QStatusBar_contextMenuEvent_default, "QStatusBar_contextMenuEvent_default")
	symaddr(&addr_QStatusBar_tabletEvent_default, "QStatusBar_tabletEvent_default")
	symaddr(&addr_QStatusBar_actionEvent_default, "QStatusBar_actionEvent_default")
	symaddr(&addr_QStatusBar_dragEnterEvent_default, "QStatusBar_dragEnterEvent_default")
	symaddr(&addr_QStatusBar_dragMoveEvent_default, "QStatusBar_dragMoveEvent_default")
	symaddr(&addr_QStatusBar_dragLeaveEvent_default, "QStatusBar_dragLeaveEvent_default")
	symaddr(&addr_QStatusBar_dropEvent_default, "QStatusBar_dropEvent_default")
	symaddr(&addr_QStatusBar_hideEvent_default, "QStatusBar_hideEvent_default")
	symaddr(&addr_QStatusBar_changeEvent_default, "QStatusBar_changeEvent_default")
	symaddr(&addr_QStatusBar_metric_default, "QStatusBar_metric_default")
	symaddr(&addr_QStatusBar_initPainter_default, "QStatusBar_initPainter_default")
	symaddr(&addr_QStatusBar_redirected_default, "QStatusBar_redirected_default")
	symaddr(&addr_QStatusBar_sharedPainter_default, "QStatusBar_sharedPainter_default")
	symaddr(&addr_QStatusBar_inputMethodEvent_default, "QStatusBar_inputMethodEvent_default")
	symaddr(&addr_QStatusBar_inputMethodQuery_default, "QStatusBar_inputMethodQuery_default")
	symaddr(&addr_QStatusBar_focusNextPrevChild_default, "QStatusBar_focusNextPrevChild_default")
	symaddr(&addr_QStatusBar_eventFilter_default, "QStatusBar_eventFilter_default")
	symaddr(&addr_QStatusBar_timerEvent_default, "QStatusBar_timerEvent_default")
	symaddr(&addr_QStatusBar_childEvent_default, "QStatusBar_childEvent_default")
	symaddr(&addr_QStatusBar_customEvent_default, "QStatusBar_customEvent_default")
	symaddr(&addr_QStatusBar_connectNotify_default, "QStatusBar_connectNotify_default")
	symaddr(&addr_QStatusBar_disconnectNotify_default, "QStatusBar_disconnectNotify_default")
	reg(&pg_QStatusBar_connect_messageChanged, "QStatusBar_connect_messageChanged")
	reg(&pg_QStatusBar_connect_windowTitleChanged, "QStatusBar_connect_windowTitleChanged")
	reg(&pg_QStatusBar_connect_windowIconChanged, "QStatusBar_connect_windowIconChanged")
	reg(&pg_QStatusBar_connect_windowIconTextChanged, "QStatusBar_connect_windowIconTextChanged")
	reg(&pg_QStatusBar_connect_customContextMenuRequested, "QStatusBar_connect_customContextMenuRequested")
	reg(&pg_QStatusBar_connect_destroyed, "QStatusBar_connect_destroyed")
	reg(&pg_QStatusBar_connect_destroyed2, "QStatusBar_connect_destroyed2")
	reg(&pg_QStatusBar_connect_objectNameChanged, "QStatusBar_connect_objectNameChanged")
	reg(&pg_QStatusBar__disconnect, "QStatusBar__disconnect")
	reg(&pg_QStatusBar__set_override, "QStatusBar__set_override")
	reg(&pg_QStatusBar_delete, "QStatusBar_delete")
}

type QStatusBar struct {
	h uintptr
	*QWidget
}

func bindQStatusBar(h uintptr) *QStatusBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStatusBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStatusBar{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQStatusBar(h uintptr) *QStatusBar { return bindQStatusBar(h) }

func (this *QStatusBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStatusBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStatusBar(h unsafe.Pointer) *QStatusBar { return bindQStatusBar(uintptr(h)) }
func UnsafeNewQStatusBarOwned(h unsafe.Pointer) *QStatusBar { return newownedQStatusBar(uintptr(h)) }

func (this *QStatusBar) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQStatusBar(parent *QWidget) *QStatusBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStatusBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStatusBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStatusBar_virtIdx, QStatusBar_enable)
	return __o
}

func NewQStatusBar2() *QStatusBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStatusBar_new2, uintptr(__goid))
	__o := newownedQStatusBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStatusBar_virtIdx, QStatusBar_enable)
	return __o
}

func QStatusBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStatusBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStatusBar) AddWidget(widget *QWidget) {
	purego.SyscallN(addr_QStatusBar_addWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStatusBar) InsertWidget(index int, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_insertWidget, this.h, uintptr(index), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStatusBar) AddPermanentWidget(widget *QWidget) {
	purego.SyscallN(addr_QStatusBar_addPermanentWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStatusBar) InsertPermanentWidget(index int, widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_insertPermanentWidget, this.h, uintptr(index), uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStatusBar) RemoveWidget(widget *QWidget) {
	purego.SyscallN(addr_QStatusBar_removeWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStatusBar) SetSizeGripEnabled(arg1 bool) {
	purego.SyscallN(addr_QStatusBar_setSizeGripEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) IsSizeGripEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_isSizeGripEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStatusBar) CurrentMessage() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStatusBar_currentMessage, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStatusBar) ShowMessage(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStatusBar_showMessage, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStatusBar) ClearMessage() {
	purego.SyscallN(addr_QStatusBar_clearMessage, this.h)
	runtime.KeepAlive(this)
}

func (this *QStatusBar) MessageChanged(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStatusBar_messageChanged, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStatusBar) Reformat() {
	purego.SyscallN(addr_QStatusBar_reformat, this.h)
	runtime.KeepAlive(this)
}

func (this *QStatusBar) HideOrShow() {
	purego.SyscallN(addr_QStatusBar_hideOrShow, this.h)
	runtime.KeepAlive(this)
}

func QStatusBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QStatusBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QStatusBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_mouseGrabber)
	return bindQWidget(__r)
}

func QStatusBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QStatusBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QStatusBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QStatusBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QStatusBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QStatusBar) Create() {
	purego.SyscallN(addr_QStatusBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QStatusBar) Destroy() {
	purego.SyscallN(addr_QStatusBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QStatusBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStatusBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QStatusBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStatusBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStatusBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStatusBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStatusBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStatusBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStatusBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStatusBar_DevicePixelRatioFScale() float64 {
	__r := pg_QStatusBar_devicePixelRatioFScale()
	return __r
}

func QStatusBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QStatusBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QStatusBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QStatusBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QStatusBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStatusBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStatusBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStatusBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QStatusBar) AddWidget2(widget *QWidget, stretch int) {
	purego.SyscallN(addr_QStatusBar_addWidget2, this.h, uintptr(widget.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStatusBar) InsertWidget2(index int, widget *QWidget, stretch int) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_insertWidget2, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(stretch))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStatusBar) AddPermanentWidget2(widget *QWidget, stretch int) {
	purego.SyscallN(addr_QStatusBar_addPermanentWidget2, this.h, uintptr(widget.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QStatusBar) InsertPermanentWidget2(index int, widget *QWidget, stretch int) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_insertPermanentWidget2, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(stretch))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QStatusBar) ShowMessage2(text string, timeout int) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStatusBar_showMessage2, this.h, text_p, uintptr(len(text)), uintptr(timeout))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func QStatusBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QStatusBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QStatusBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QStatusBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QStatusBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QStatusBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QStatusBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QStatusBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QStatusBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QStatusBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStatusBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStatusBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStatusBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStatusBar) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QStatusBar_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStatusBar) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QStatusBar_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStatusBar) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QStatusBar_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStatusBar) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QStatusBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStatusBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QStatusBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QStatusBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStatusBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStatusBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStatusBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStatusBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QStatusBar) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStatusBar_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStatusBar_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStatusBar_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QStatusBar_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QStatusBar_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QStatusBar_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QStatusBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QStatusBar_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QStatusBar_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QStatusBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStatusBar_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QStatusBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QStatusBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QStatusBar_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QStatusBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QStatusBar_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QStatusBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QStatusBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QStatusBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QStatusBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QStatusBar_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QStatusBar_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStatusBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStatusBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QStatusBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QStatusBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QStatusBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QStatusBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QStatusBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStatusBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStatusBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStatusBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStatusBar) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStatusBar_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStatusBar) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStatusBar_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStatusBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStatusBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStatusBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStatusBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStatusBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStatusBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStatusBar) OnMessageChanged(fn func(string)) int64 {
	id := registerCallback(func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) })
	pg_QStatusBar_connect_messageChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QStatusBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QStatusBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QStatusBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QStatusBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStatusBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStatusBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStatusBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStatusBar) Disconnect(handle int64) {
	pg_QStatusBar__disconnect(handle)
	unregisterCallback(handle)
}

func QStatusBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -6825369174751916:
		return 1
	case 5290442163167458465:
		return 2
	case -4302917285655014661:
		return 3
	case 1237236456312824991:
		return 4
	case -8659679893868524044:
		return 5
	case 2110182719375245575:
		return 6
	case 729690261457684101:
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
	case -199718707876719129:
		return 32
	case -2039166586246443706:
		return 33
	case 4959376063826165943:
		return 34
	case 427629649267311777:
		return 35
	case 6477500296980050210:
		return 36
	case 4127240633310655048:
		return 37
	case 5792771328741774991:
		return 38
	case 5438819661942843222:
		return 39
	case 7345853097340709638:
		return 40
	case 5898328517505784191:
		return 41
	case 5217623514630858297:
		return 42
	case -5409927581812374498:
		return 43
	case -1185315367526857833:
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

func QStatusBar_enable(h uintptr, idx int32) { pg_QStatusBar__set_override(h, idx) }

func (this *QStatusBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStatusBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QStatusBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QStatusBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QStatusBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QStatusBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QStatusBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QStatusBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QStatusBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QStatusBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QStatusBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QStatusBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QStatusBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QStatusBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QStatusBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QStatusBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QStatusBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QStatusBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QStatusBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QStatusBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QStatusBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QStatusBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QStatusBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QStatusBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QStatusBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QStatusBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QStatusBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QStatusBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QStatusBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QStatusBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QStatusBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QStatusBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QStatusBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QStatusBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QStatusBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QStatusBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QStatusBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QStatusBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QStatusBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QStatusBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QStatusBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QStatusBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QStatusBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStatusBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStatusBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStatusBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStatusBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStatusBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStatusBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStatusBar_delete(this.h)
	}
	this.UnsafeRelease()
}
