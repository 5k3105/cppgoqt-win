package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGroupBox_chain uintptr
	addr_QGroupBox_new uintptr
	addr_QGroupBox_new2 uintptr
	addr_QGroupBox_new3 uintptr
	addr_QGroupBox_new4 uintptr
	addr_QGroupBox_tr uintptr
	addr_QGroupBox_title uintptr
	addr_QGroupBox_setTitle uintptr
	addr_QGroupBox_alignment uintptr
	addr_QGroupBox_setAlignment uintptr
	addr_QGroupBox_isFlat uintptr
	addr_QGroupBox_setFlat uintptr
	addr_QGroupBox_isCheckable uintptr
	addr_QGroupBox_setCheckable uintptr
	addr_QGroupBox_isChecked uintptr
	addr_QGroupBox_setChecked uintptr
	addr_QGroupBox_clicked uintptr
	addr_QGroupBox_toggled uintptr
	addr_QGroupBox_setTabOrder uintptr
	addr_QGroupBox_mouseGrabber uintptr
	addr_QGroupBox_keyboardGrabber uintptr
	addr_QGroupBox_find uintptr
	addr_QGroupBox_createWindowContainer uintptr
	addr_QGroupBox_updateMicroFocus uintptr
	addr_QGroupBox_create uintptr
	addr_QGroupBox_destroy uintptr
	addr_QGroupBox_focusNextChild uintptr
	addr_QGroupBox_focusPreviousChild uintptr
	addr_QGroupBox_connect uintptr
	addr_QGroupBox_connect2 uintptr
	addr_QGroupBox_sender uintptr
	addr_QGroupBox_senderSignalIndex uintptr
	addr_QGroupBox_receivers uintptr
	addr_QGroupBox_isSignalConnected uintptr
	pg_QGroupBox_devicePixelRatioFScale func() float64
	pg_QGroupBox_encodeMetricF func(int32, float64) int32
	pg_QGroupBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QGroupBox_tr2 uintptr
	addr_QGroupBox_tr3 uintptr
	addr_QGroupBox_clicked2 uintptr
	addr_QGroupBox_createWindowContainer2 uintptr
	addr_QGroupBox_createWindowContainer3 uintptr
	addr_QGroupBox_updateMicroFocus2 uintptr
	addr_QGroupBox_create2 uintptr
	addr_QGroupBox_create3 uintptr
	addr_QGroupBox_create4 uintptr
	addr_QGroupBox_destroy2 uintptr
	addr_QGroupBox_destroy3 uintptr
	addr_QGroupBox_connect3 uintptr
	addr_QGroupBox_connect4 uintptr
	addr_QGroupBox_metaObject_default uintptr
	addr_QGroupBox_minimumSizeHint_default uintptr
	addr_QGroupBox_event_default uintptr
	addr_QGroupBox_childEvent_default uintptr
	addr_QGroupBox_resizeEvent_default uintptr
	addr_QGroupBox_paintEvent_default uintptr
	addr_QGroupBox_focusInEvent_default uintptr
	addr_QGroupBox_changeEvent_default uintptr
	addr_QGroupBox_mousePressEvent_default uintptr
	addr_QGroupBox_mouseMoveEvent_default uintptr
	addr_QGroupBox_mouseReleaseEvent_default uintptr
	addr_QGroupBox_initStyleOption_default uintptr
	addr_QGroupBox_devType_default uintptr
	addr_QGroupBox_setVisible_default uintptr
	addr_QGroupBox_sizeHint_default uintptr
	addr_QGroupBox_heightForWidth_default uintptr
	addr_QGroupBox_hasHeightForWidth_default uintptr
	addr_QGroupBox_paintEngine_default uintptr
	addr_QGroupBox_mouseDoubleClickEvent_default uintptr
	addr_QGroupBox_wheelEvent_default uintptr
	addr_QGroupBox_keyPressEvent_default uintptr
	addr_QGroupBox_keyReleaseEvent_default uintptr
	addr_QGroupBox_focusOutEvent_default uintptr
	addr_QGroupBox_enterEvent_default uintptr
	addr_QGroupBox_leaveEvent_default uintptr
	addr_QGroupBox_moveEvent_default uintptr
	addr_QGroupBox_closeEvent_default uintptr
	addr_QGroupBox_contextMenuEvent_default uintptr
	addr_QGroupBox_tabletEvent_default uintptr
	addr_QGroupBox_actionEvent_default uintptr
	addr_QGroupBox_dragEnterEvent_default uintptr
	addr_QGroupBox_dragMoveEvent_default uintptr
	addr_QGroupBox_dragLeaveEvent_default uintptr
	addr_QGroupBox_dropEvent_default uintptr
	addr_QGroupBox_showEvent_default uintptr
	addr_QGroupBox_hideEvent_default uintptr
	addr_QGroupBox_metric_default uintptr
	addr_QGroupBox_initPainter_default uintptr
	addr_QGroupBox_redirected_default uintptr
	addr_QGroupBox_sharedPainter_default uintptr
	addr_QGroupBox_inputMethodEvent_default uintptr
	addr_QGroupBox_inputMethodQuery_default uintptr
	addr_QGroupBox_focusNextPrevChild_default uintptr
	addr_QGroupBox_eventFilter_default uintptr
	addr_QGroupBox_timerEvent_default uintptr
	addr_QGroupBox_customEvent_default uintptr
	addr_QGroupBox_connectNotify_default uintptr
	addr_QGroupBox_disconnectNotify_default uintptr
	pg_QGroupBox_connect_clicked func(uintptr, int64)
	pg_QGroupBox_connect_clicked2 func(uintptr, int64)
	pg_QGroupBox_connect_toggled func(uintptr, int64)
	pg_QGroupBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QGroupBox_connect_windowIconChanged func(uintptr, int64)
	pg_QGroupBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QGroupBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QGroupBox_connect_destroyed func(uintptr, int64)
	pg_QGroupBox_connect_destroyed2 func(uintptr, int64)
	pg_QGroupBox_connect_objectNameChanged func(uintptr, int64)
	pg_QGroupBox__disconnect func(int64)
	pg_QGroupBox__set_override func(uintptr, int32)
	pg_QGroupBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGroupBox_chain, "QGroupBox_chain")
	symaddr(&addr_QGroupBox_new, "QGroupBox_new")
	symaddr(&addr_QGroupBox_new2, "QGroupBox_new2")
	symaddr(&addr_QGroupBox_new3, "QGroupBox_new3")
	symaddr(&addr_QGroupBox_new4, "QGroupBox_new4")
	symaddr(&addr_QGroupBox_tr, "QGroupBox_tr")
	symaddr(&addr_QGroupBox_title, "QGroupBox_title")
	symaddr(&addr_QGroupBox_setTitle, "QGroupBox_setTitle")
	symaddr(&addr_QGroupBox_alignment, "QGroupBox_alignment")
	symaddr(&addr_QGroupBox_setAlignment, "QGroupBox_setAlignment")
	symaddr(&addr_QGroupBox_isFlat, "QGroupBox_isFlat")
	symaddr(&addr_QGroupBox_setFlat, "QGroupBox_setFlat")
	symaddr(&addr_QGroupBox_isCheckable, "QGroupBox_isCheckable")
	symaddr(&addr_QGroupBox_setCheckable, "QGroupBox_setCheckable")
	symaddr(&addr_QGroupBox_isChecked, "QGroupBox_isChecked")
	symaddr(&addr_QGroupBox_setChecked, "QGroupBox_setChecked")
	symaddr(&addr_QGroupBox_clicked, "QGroupBox_clicked")
	symaddr(&addr_QGroupBox_toggled, "QGroupBox_toggled")
	symaddr(&addr_QGroupBox_setTabOrder, "QGroupBox_setTabOrder")
	symaddr(&addr_QGroupBox_mouseGrabber, "QGroupBox_mouseGrabber")
	symaddr(&addr_QGroupBox_keyboardGrabber, "QGroupBox_keyboardGrabber")
	symaddr(&addr_QGroupBox_find, "QGroupBox_find")
	symaddr(&addr_QGroupBox_createWindowContainer, "QGroupBox_createWindowContainer")
	symaddr(&addr_QGroupBox_updateMicroFocus, "QGroupBox_updateMicroFocus")
	symaddr(&addr_QGroupBox_create, "QGroupBox_create")
	symaddr(&addr_QGroupBox_destroy, "QGroupBox_destroy")
	symaddr(&addr_QGroupBox_focusNextChild, "QGroupBox_focusNextChild")
	symaddr(&addr_QGroupBox_focusPreviousChild, "QGroupBox_focusPreviousChild")
	symaddr(&addr_QGroupBox_connect, "QGroupBox_connect")
	symaddr(&addr_QGroupBox_connect2, "QGroupBox_connect2")
	symaddr(&addr_QGroupBox_sender, "QGroupBox_sender")
	symaddr(&addr_QGroupBox_senderSignalIndex, "QGroupBox_senderSignalIndex")
	symaddr(&addr_QGroupBox_receivers, "QGroupBox_receivers")
	symaddr(&addr_QGroupBox_isSignalConnected, "QGroupBox_isSignalConnected")
	reg(&pg_QGroupBox_devicePixelRatioFScale, "QGroupBox_devicePixelRatioFScale")
	reg(&pg_QGroupBox_encodeMetricF, "QGroupBox_encodeMetricF")
	reg(&pg_QGroupBox_getDecodedMetricF, "QGroupBox_getDecodedMetricF")
	symaddr(&addr_QGroupBox_tr2, "QGroupBox_tr2")
	symaddr(&addr_QGroupBox_tr3, "QGroupBox_tr3")
	symaddr(&addr_QGroupBox_clicked2, "QGroupBox_clicked2")
	symaddr(&addr_QGroupBox_createWindowContainer2, "QGroupBox_createWindowContainer2")
	symaddr(&addr_QGroupBox_createWindowContainer3, "QGroupBox_createWindowContainer3")
	symaddr(&addr_QGroupBox_updateMicroFocus2, "QGroupBox_updateMicroFocus2")
	symaddr(&addr_QGroupBox_create2, "QGroupBox_create2")
	symaddr(&addr_QGroupBox_create3, "QGroupBox_create3")
	symaddr(&addr_QGroupBox_create4, "QGroupBox_create4")
	symaddr(&addr_QGroupBox_destroy2, "QGroupBox_destroy2")
	symaddr(&addr_QGroupBox_destroy3, "QGroupBox_destroy3")
	symaddr(&addr_QGroupBox_connect3, "QGroupBox_connect3")
	symaddr(&addr_QGroupBox_connect4, "QGroupBox_connect4")
	symaddr(&addr_QGroupBox_metaObject_default, "QGroupBox_metaObject_default")
	symaddr(&addr_QGroupBox_minimumSizeHint_default, "QGroupBox_minimumSizeHint_default")
	symaddr(&addr_QGroupBox_event_default, "QGroupBox_event_default")
	symaddr(&addr_QGroupBox_childEvent_default, "QGroupBox_childEvent_default")
	symaddr(&addr_QGroupBox_resizeEvent_default, "QGroupBox_resizeEvent_default")
	symaddr(&addr_QGroupBox_paintEvent_default, "QGroupBox_paintEvent_default")
	symaddr(&addr_QGroupBox_focusInEvent_default, "QGroupBox_focusInEvent_default")
	symaddr(&addr_QGroupBox_changeEvent_default, "QGroupBox_changeEvent_default")
	symaddr(&addr_QGroupBox_mousePressEvent_default, "QGroupBox_mousePressEvent_default")
	symaddr(&addr_QGroupBox_mouseMoveEvent_default, "QGroupBox_mouseMoveEvent_default")
	symaddr(&addr_QGroupBox_mouseReleaseEvent_default, "QGroupBox_mouseReleaseEvent_default")
	symaddr(&addr_QGroupBox_initStyleOption_default, "QGroupBox_initStyleOption_default")
	symaddr(&addr_QGroupBox_devType_default, "QGroupBox_devType_default")
	symaddr(&addr_QGroupBox_setVisible_default, "QGroupBox_setVisible_default")
	symaddr(&addr_QGroupBox_sizeHint_default, "QGroupBox_sizeHint_default")
	symaddr(&addr_QGroupBox_heightForWidth_default, "QGroupBox_heightForWidth_default")
	symaddr(&addr_QGroupBox_hasHeightForWidth_default, "QGroupBox_hasHeightForWidth_default")
	symaddr(&addr_QGroupBox_paintEngine_default, "QGroupBox_paintEngine_default")
	symaddr(&addr_QGroupBox_mouseDoubleClickEvent_default, "QGroupBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QGroupBox_wheelEvent_default, "QGroupBox_wheelEvent_default")
	symaddr(&addr_QGroupBox_keyPressEvent_default, "QGroupBox_keyPressEvent_default")
	symaddr(&addr_QGroupBox_keyReleaseEvent_default, "QGroupBox_keyReleaseEvent_default")
	symaddr(&addr_QGroupBox_focusOutEvent_default, "QGroupBox_focusOutEvent_default")
	symaddr(&addr_QGroupBox_enterEvent_default, "QGroupBox_enterEvent_default")
	symaddr(&addr_QGroupBox_leaveEvent_default, "QGroupBox_leaveEvent_default")
	symaddr(&addr_QGroupBox_moveEvent_default, "QGroupBox_moveEvent_default")
	symaddr(&addr_QGroupBox_closeEvent_default, "QGroupBox_closeEvent_default")
	symaddr(&addr_QGroupBox_contextMenuEvent_default, "QGroupBox_contextMenuEvent_default")
	symaddr(&addr_QGroupBox_tabletEvent_default, "QGroupBox_tabletEvent_default")
	symaddr(&addr_QGroupBox_actionEvent_default, "QGroupBox_actionEvent_default")
	symaddr(&addr_QGroupBox_dragEnterEvent_default, "QGroupBox_dragEnterEvent_default")
	symaddr(&addr_QGroupBox_dragMoveEvent_default, "QGroupBox_dragMoveEvent_default")
	symaddr(&addr_QGroupBox_dragLeaveEvent_default, "QGroupBox_dragLeaveEvent_default")
	symaddr(&addr_QGroupBox_dropEvent_default, "QGroupBox_dropEvent_default")
	symaddr(&addr_QGroupBox_showEvent_default, "QGroupBox_showEvent_default")
	symaddr(&addr_QGroupBox_hideEvent_default, "QGroupBox_hideEvent_default")
	symaddr(&addr_QGroupBox_metric_default, "QGroupBox_metric_default")
	symaddr(&addr_QGroupBox_initPainter_default, "QGroupBox_initPainter_default")
	symaddr(&addr_QGroupBox_redirected_default, "QGroupBox_redirected_default")
	symaddr(&addr_QGroupBox_sharedPainter_default, "QGroupBox_sharedPainter_default")
	symaddr(&addr_QGroupBox_inputMethodEvent_default, "QGroupBox_inputMethodEvent_default")
	symaddr(&addr_QGroupBox_inputMethodQuery_default, "QGroupBox_inputMethodQuery_default")
	symaddr(&addr_QGroupBox_focusNextPrevChild_default, "QGroupBox_focusNextPrevChild_default")
	symaddr(&addr_QGroupBox_eventFilter_default, "QGroupBox_eventFilter_default")
	symaddr(&addr_QGroupBox_timerEvent_default, "QGroupBox_timerEvent_default")
	symaddr(&addr_QGroupBox_customEvent_default, "QGroupBox_customEvent_default")
	symaddr(&addr_QGroupBox_connectNotify_default, "QGroupBox_connectNotify_default")
	symaddr(&addr_QGroupBox_disconnectNotify_default, "QGroupBox_disconnectNotify_default")
	reg(&pg_QGroupBox_connect_clicked, "QGroupBox_connect_clicked")
	reg(&pg_QGroupBox_connect_clicked2, "QGroupBox_connect_clicked2")
	reg(&pg_QGroupBox_connect_toggled, "QGroupBox_connect_toggled")
	reg(&pg_QGroupBox_connect_windowTitleChanged, "QGroupBox_connect_windowTitleChanged")
	reg(&pg_QGroupBox_connect_windowIconChanged, "QGroupBox_connect_windowIconChanged")
	reg(&pg_QGroupBox_connect_windowIconTextChanged, "QGroupBox_connect_windowIconTextChanged")
	reg(&pg_QGroupBox_connect_customContextMenuRequested, "QGroupBox_connect_customContextMenuRequested")
	reg(&pg_QGroupBox_connect_destroyed, "QGroupBox_connect_destroyed")
	reg(&pg_QGroupBox_connect_destroyed2, "QGroupBox_connect_destroyed2")
	reg(&pg_QGroupBox_connect_objectNameChanged, "QGroupBox_connect_objectNameChanged")
	reg(&pg_QGroupBox__disconnect, "QGroupBox__disconnect")
	reg(&pg_QGroupBox__set_override, "QGroupBox__set_override")
	reg(&pg_QGroupBox_delete, "QGroupBox_delete")
}

type QGroupBox struct {
	h uintptr
	*QWidget
}

func bindQGroupBox(h uintptr) *QGroupBox {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGroupBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGroupBox{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGroupBox(h uintptr) *QGroupBox { return bindQGroupBox(h) }

func (this *QGroupBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGroupBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGroupBox(h unsafe.Pointer) *QGroupBox { return bindQGroupBox(uintptr(h)) }
func UnsafeNewQGroupBoxOwned(h unsafe.Pointer) *QGroupBox { return newownedQGroupBox(uintptr(h)) }

func (this *QGroupBox) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQGroupBox(parent *QWidget) *QGroupBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGroupBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGroupBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGroupBox_virtIdx, QGroupBox_enable)
	return __o
}

func NewQGroupBox2() *QGroupBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGroupBox_new2, uintptr(__goid))
	__o := newownedQGroupBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGroupBox_virtIdx, QGroupBox_enable)
	return __o
}

func NewQGroupBox3(title string) *QGroupBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QGroupBox_new3, uintptr(__goid), title_p, uintptr(len(title)))
	__o := newownedQGroupBox(__r)
	runtime.KeepAlive(title)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGroupBox_virtIdx, QGroupBox_enable)
	return __o
}

func NewQGroupBox4(title string, parent *QWidget) *QGroupBox {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QGroupBox_new4, uintptr(__goid), title_p, uintptr(len(title)), uintptr(parent.UnsafePointer()))
	__o := newownedQGroupBox(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGroupBox_virtIdx, QGroupBox_enable)
	return __o
}

func QGroupBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGroupBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGroupBox) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGroupBox_title, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGroupBox) SetTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QGroupBox_setTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QGroupBox) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QGroupBox) SetAlignment(alignment int) {
	purego.SyscallN(addr_QGroupBox_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) IsFlat() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_isFlat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) SetFlat(flat bool) {
	purego.SyscallN(addr_QGroupBox_setFlat, this.h, uintptr(b2i(flat)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) IsCheckable() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_isCheckable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) SetCheckable(checkable bool) {
	purego.SyscallN(addr_QGroupBox_setCheckable, this.h, uintptr(b2i(checkable)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) IsChecked() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_isChecked, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) SetChecked(checked bool) {
	purego.SyscallN(addr_QGroupBox_setChecked, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Clicked() {
	purego.SyscallN(addr_QGroupBox_clicked, this.h)
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Toggled(arg1 bool) {
	purego.SyscallN(addr_QGroupBox_toggled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func QGroupBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QGroupBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QGroupBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_mouseGrabber)
	return bindQWidget(__r)
}

func QGroupBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QGroupBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QGroupBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QGroupBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QGroupBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Create() {
	purego.SyscallN(addr_QGroupBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Destroy() {
	purego.SyscallN(addr_QGroupBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QGroupBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGroupBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGroupBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGroupBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGroupBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGroupBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGroupBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGroupBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGroupBox_DevicePixelRatioFScale() float64 {
	__r := pg_QGroupBox_devicePixelRatioFScale()
	return __r
}

func QGroupBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QGroupBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QGroupBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QGroupBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QGroupBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGroupBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGroupBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGroupBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGroupBox) ClickedWithChecked(checked bool) {
	purego.SyscallN(addr_QGroupBox_clicked2, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func QGroupBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QGroupBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QGroupBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QGroupBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QGroupBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QGroupBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QGroupBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QGroupBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QGroupBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QGroupBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGroupBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGroupBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGroupBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGroupBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGroupBox) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGroupBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGroupBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QGroupBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QGroupBox_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGroupBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGroupBox_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGroupBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGroupBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGroupBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) InitStyleOptionDefault(option *QStyleOptionGroupBox) {
	purego.SyscallN(addr_QGroupBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QGroupBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGroupBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QGroupBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QGroupBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGroupBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGroupBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QGroupBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGroupBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QGroupBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGroupBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGroupBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGroupBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QGroupBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGroupBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QGroupBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QGroupBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QGroupBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QGroupBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QGroupBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QGroupBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QGroupBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QGroupBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QGroupBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QGroupBox_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QGroupBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGroupBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QGroupBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGroupBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QGroupBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QGroupBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QGroupBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGroupBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGroupBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGroupBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGroupBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGroupBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGroupBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGroupBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGroupBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGroupBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGroupBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGroupBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGroupBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGroupBox) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGroupBox_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QGroupBox_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QGroupBox_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QGroupBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QGroupBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QGroupBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QGroupBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGroupBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGroupBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGroupBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGroupBox) Disconnect(handle int64) {
	pg_QGroupBox__disconnect(handle)
	unregisterCallback(handle)
}

func QGroupBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7049264599837532491:
		return 1
	case 1237236456312824991:
		return 2
	case -1185315367526857833:
		return 3
	case -4302917285655014661:
		return 4
	case 5290442163167458465:
		return 5
	case 975043767033552552:
		return 6
	case 4959376063826165943:
		return 7
	case -332090898078543319:
		return 8
	case -4255494004099806581:
		return 9
	case -4415920049297795213:
		return 10
	case -488693769771308531:
		return 11
	case -8659679893868524044:
		return 12
	case 2110182719375245575:
		return 13
	case 729690261457684101:
		return 14
	case 4311534060495737339:
		return 15
	case -8266762397995423915:
		return 16
	case 6398007404756346381:
		return 17
	case -4029073905697023887:
		return 18
	case 1919988340934511316:
		return 19
	case 1023791977679793185:
		return 20
	case 772598475950050299:
		return 21
	case -385596010576048993:
		return 22
	case 6316270346668237895:
		return 23
	case -551070009623711640:
		return 24
	case -9161949150489525396:
		return 25
	case -1989937050019464783:
		return 26
	case -6698043222227006015:
		return 27
	case -212068866860701141:
		return 28
	case 1697444140088555581:
		return 29
	case 206122715919917589:
		return 30
	case -1781545137352745978:
		return 31
	case -3139904158373725006:
		return 32
	case 9107058682754591716:
		return 33
	case -6825369174751916:
		return 34
	case -199718707876719129:
		return 35
	case -2039166586246443706:
		return 36
	case 427629649267311777:
		return 37
	case 6477500296980050210:
		return 38
	case 4127240633310655048:
		return 39
	case 5792771328741774991:
		return 40
	case 5438819661942843222:
		return 41
	case 7345853097340709638:
		return 42
	case 5898328517505784191:
		return 43
	case 5217623514630858297:
		return 44
	case -5409927581812374498:
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

func QGroupBox_enable(h uintptr, idx int32) { pg_QGroupBox__set_override(h, idx) }

func (this *QGroupBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGroupBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QGroupBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGroupBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGroupBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QGroupBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QGroupBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGroupBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QGroupBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGroupBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGroupBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGroupBox) OnInitStyleOption(fn func(*QStyleOptionGroupBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionGroupBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QGroupBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QGroupBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QGroupBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGroupBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QGroupBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QGroupBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QGroupBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGroupBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGroupBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGroupBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGroupBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGroupBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QGroupBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QGroupBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QGroupBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QGroupBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGroupBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QGroupBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QGroupBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGroupBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGroupBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGroupBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGroupBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QGroupBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QGroupBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QGroupBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QGroupBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QGroupBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QGroupBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QGroupBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QGroupBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGroupBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QGroupBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGroupBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGroupBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGroupBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGroupBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGroupBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGroupBox_delete(this.h)
	}
	this.UnsafeRelease()
}
