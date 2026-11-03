package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QScrollArea_chain uintptr
	addr_QScrollArea_new uintptr
	addr_QScrollArea_new2 uintptr
	addr_QScrollArea_tr uintptr
	addr_QScrollArea_widget uintptr
	addr_QScrollArea_setWidget uintptr
	addr_QScrollArea_takeWidget uintptr
	addr_QScrollArea_widgetResizable uintptr
	addr_QScrollArea_setWidgetResizable uintptr
	addr_QScrollArea_alignment uintptr
	addr_QScrollArea_setAlignment uintptr
	addr_QScrollArea_ensureVisible uintptr
	addr_QScrollArea_ensureWidgetVisible uintptr
	addr_QScrollArea_setViewportMargins uintptr
	addr_QScrollArea_setViewportMargins2 uintptr
	addr_QScrollArea_viewportMargins uintptr
	addr_QScrollArea_drawFrame uintptr
	addr_QScrollArea_setTabOrder uintptr
	addr_QScrollArea_mouseGrabber uintptr
	addr_QScrollArea_keyboardGrabber uintptr
	addr_QScrollArea_find uintptr
	addr_QScrollArea_createWindowContainer uintptr
	addr_QScrollArea_updateMicroFocus uintptr
	addr_QScrollArea_create uintptr
	addr_QScrollArea_destroy uintptr
	addr_QScrollArea_focusNextChild uintptr
	addr_QScrollArea_focusPreviousChild uintptr
	addr_QScrollArea_connect uintptr
	addr_QScrollArea_connect2 uintptr
	addr_QScrollArea_sender uintptr
	addr_QScrollArea_senderSignalIndex uintptr
	addr_QScrollArea_receivers uintptr
	addr_QScrollArea_isSignalConnected uintptr
	pg_QScrollArea_devicePixelRatioFScale func() float64
	pg_QScrollArea_encodeMetricF func(int32, float64) int32
	pg_QScrollArea_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QScrollArea_tr2 uintptr
	addr_QScrollArea_tr3 uintptr
	addr_QScrollArea_ensureVisible2 uintptr
	addr_QScrollArea_ensureVisible3 uintptr
	addr_QScrollArea_ensureWidgetVisible2 uintptr
	addr_QScrollArea_ensureWidgetVisible3 uintptr
	addr_QScrollArea_createWindowContainer2 uintptr
	addr_QScrollArea_createWindowContainer3 uintptr
	addr_QScrollArea_updateMicroFocus2 uintptr
	addr_QScrollArea_create2 uintptr
	addr_QScrollArea_create3 uintptr
	addr_QScrollArea_create4 uintptr
	addr_QScrollArea_destroy2 uintptr
	addr_QScrollArea_destroy3 uintptr
	addr_QScrollArea_connect3 uintptr
	addr_QScrollArea_connect4 uintptr
	addr_QScrollArea_metaObject_default uintptr
	addr_QScrollArea_sizeHint_default uintptr
	addr_QScrollArea_focusNextPrevChild_default uintptr
	addr_QScrollArea_event_default uintptr
	addr_QScrollArea_eventFilter_default uintptr
	addr_QScrollArea_resizeEvent_default uintptr
	addr_QScrollArea_scrollContentsBy_default uintptr
	addr_QScrollArea_viewportSizeHint_default uintptr
	addr_QScrollArea_minimumSizeHint_default uintptr
	addr_QScrollArea_setupViewport_default uintptr
	addr_QScrollArea_viewportEvent_default uintptr
	addr_QScrollArea_paintEvent_default uintptr
	addr_QScrollArea_mousePressEvent_default uintptr
	addr_QScrollArea_mouseReleaseEvent_default uintptr
	addr_QScrollArea_mouseDoubleClickEvent_default uintptr
	addr_QScrollArea_mouseMoveEvent_default uintptr
	addr_QScrollArea_wheelEvent_default uintptr
	addr_QScrollArea_contextMenuEvent_default uintptr
	addr_QScrollArea_dragEnterEvent_default uintptr
	addr_QScrollArea_dragMoveEvent_default uintptr
	addr_QScrollArea_dragLeaveEvent_default uintptr
	addr_QScrollArea_dropEvent_default uintptr
	addr_QScrollArea_keyPressEvent_default uintptr
	addr_QScrollArea_changeEvent_default uintptr
	addr_QScrollArea_initStyleOption_default uintptr
	addr_QScrollArea_devType_default uintptr
	addr_QScrollArea_setVisible_default uintptr
	addr_QScrollArea_heightForWidth_default uintptr
	addr_QScrollArea_hasHeightForWidth_default uintptr
	addr_QScrollArea_paintEngine_default uintptr
	addr_QScrollArea_keyReleaseEvent_default uintptr
	addr_QScrollArea_focusInEvent_default uintptr
	addr_QScrollArea_focusOutEvent_default uintptr
	addr_QScrollArea_enterEvent_default uintptr
	addr_QScrollArea_leaveEvent_default uintptr
	addr_QScrollArea_moveEvent_default uintptr
	addr_QScrollArea_closeEvent_default uintptr
	addr_QScrollArea_tabletEvent_default uintptr
	addr_QScrollArea_actionEvent_default uintptr
	addr_QScrollArea_showEvent_default uintptr
	addr_QScrollArea_hideEvent_default uintptr
	addr_QScrollArea_metric_default uintptr
	addr_QScrollArea_initPainter_default uintptr
	addr_QScrollArea_redirected_default uintptr
	addr_QScrollArea_sharedPainter_default uintptr
	addr_QScrollArea_inputMethodEvent_default uintptr
	addr_QScrollArea_inputMethodQuery_default uintptr
	addr_QScrollArea_timerEvent_default uintptr
	addr_QScrollArea_childEvent_default uintptr
	addr_QScrollArea_customEvent_default uintptr
	addr_QScrollArea_connectNotify_default uintptr
	addr_QScrollArea_disconnectNotify_default uintptr
	pg_QScrollArea_connect_windowTitleChanged func(uintptr, int64)
	pg_QScrollArea_connect_windowIconChanged func(uintptr, int64)
	pg_QScrollArea_connect_windowIconTextChanged func(uintptr, int64)
	pg_QScrollArea_connect_customContextMenuRequested func(uintptr, int64)
	pg_QScrollArea_connect_destroyed func(uintptr, int64)
	pg_QScrollArea_connect_destroyed2 func(uintptr, int64)
	pg_QScrollArea_connect_objectNameChanged func(uintptr, int64)
	pg_QScrollArea__disconnect func(int64)
	pg_QScrollArea__set_override func(uintptr, int32)
	pg_QScrollArea_delete func(uintptr)
)

func init() {
	symaddr(&addr_QScrollArea_chain, "QScrollArea_chain")
	symaddr(&addr_QScrollArea_new, "QScrollArea_new")
	symaddr(&addr_QScrollArea_new2, "QScrollArea_new2")
	symaddr(&addr_QScrollArea_tr, "QScrollArea_tr")
	symaddr(&addr_QScrollArea_widget, "QScrollArea_widget")
	symaddr(&addr_QScrollArea_setWidget, "QScrollArea_setWidget")
	symaddr(&addr_QScrollArea_takeWidget, "QScrollArea_takeWidget")
	symaddr(&addr_QScrollArea_widgetResizable, "QScrollArea_widgetResizable")
	symaddr(&addr_QScrollArea_setWidgetResizable, "QScrollArea_setWidgetResizable")
	symaddr(&addr_QScrollArea_alignment, "QScrollArea_alignment")
	symaddr(&addr_QScrollArea_setAlignment, "QScrollArea_setAlignment")
	symaddr(&addr_QScrollArea_ensureVisible, "QScrollArea_ensureVisible")
	symaddr(&addr_QScrollArea_ensureWidgetVisible, "QScrollArea_ensureWidgetVisible")
	symaddr(&addr_QScrollArea_setViewportMargins, "QScrollArea_setViewportMargins")
	symaddr(&addr_QScrollArea_setViewportMargins2, "QScrollArea_setViewportMargins2")
	symaddr(&addr_QScrollArea_viewportMargins, "QScrollArea_viewportMargins")
	symaddr(&addr_QScrollArea_drawFrame, "QScrollArea_drawFrame")
	symaddr(&addr_QScrollArea_setTabOrder, "QScrollArea_setTabOrder")
	symaddr(&addr_QScrollArea_mouseGrabber, "QScrollArea_mouseGrabber")
	symaddr(&addr_QScrollArea_keyboardGrabber, "QScrollArea_keyboardGrabber")
	symaddr(&addr_QScrollArea_find, "QScrollArea_find")
	symaddr(&addr_QScrollArea_createWindowContainer, "QScrollArea_createWindowContainer")
	symaddr(&addr_QScrollArea_updateMicroFocus, "QScrollArea_updateMicroFocus")
	symaddr(&addr_QScrollArea_create, "QScrollArea_create")
	symaddr(&addr_QScrollArea_destroy, "QScrollArea_destroy")
	symaddr(&addr_QScrollArea_focusNextChild, "QScrollArea_focusNextChild")
	symaddr(&addr_QScrollArea_focusPreviousChild, "QScrollArea_focusPreviousChild")
	symaddr(&addr_QScrollArea_connect, "QScrollArea_connect")
	symaddr(&addr_QScrollArea_connect2, "QScrollArea_connect2")
	symaddr(&addr_QScrollArea_sender, "QScrollArea_sender")
	symaddr(&addr_QScrollArea_senderSignalIndex, "QScrollArea_senderSignalIndex")
	symaddr(&addr_QScrollArea_receivers, "QScrollArea_receivers")
	symaddr(&addr_QScrollArea_isSignalConnected, "QScrollArea_isSignalConnected")
	reg(&pg_QScrollArea_devicePixelRatioFScale, "QScrollArea_devicePixelRatioFScale")
	reg(&pg_QScrollArea_encodeMetricF, "QScrollArea_encodeMetricF")
	reg(&pg_QScrollArea_getDecodedMetricF, "QScrollArea_getDecodedMetricF")
	symaddr(&addr_QScrollArea_tr2, "QScrollArea_tr2")
	symaddr(&addr_QScrollArea_tr3, "QScrollArea_tr3")
	symaddr(&addr_QScrollArea_ensureVisible2, "QScrollArea_ensureVisible2")
	symaddr(&addr_QScrollArea_ensureVisible3, "QScrollArea_ensureVisible3")
	symaddr(&addr_QScrollArea_ensureWidgetVisible2, "QScrollArea_ensureWidgetVisible2")
	symaddr(&addr_QScrollArea_ensureWidgetVisible3, "QScrollArea_ensureWidgetVisible3")
	symaddr(&addr_QScrollArea_createWindowContainer2, "QScrollArea_createWindowContainer2")
	symaddr(&addr_QScrollArea_createWindowContainer3, "QScrollArea_createWindowContainer3")
	symaddr(&addr_QScrollArea_updateMicroFocus2, "QScrollArea_updateMicroFocus2")
	symaddr(&addr_QScrollArea_create2, "QScrollArea_create2")
	symaddr(&addr_QScrollArea_create3, "QScrollArea_create3")
	symaddr(&addr_QScrollArea_create4, "QScrollArea_create4")
	symaddr(&addr_QScrollArea_destroy2, "QScrollArea_destroy2")
	symaddr(&addr_QScrollArea_destroy3, "QScrollArea_destroy3")
	symaddr(&addr_QScrollArea_connect3, "QScrollArea_connect3")
	symaddr(&addr_QScrollArea_connect4, "QScrollArea_connect4")
	symaddr(&addr_QScrollArea_metaObject_default, "QScrollArea_metaObject_default")
	symaddr(&addr_QScrollArea_sizeHint_default, "QScrollArea_sizeHint_default")
	symaddr(&addr_QScrollArea_focusNextPrevChild_default, "QScrollArea_focusNextPrevChild_default")
	symaddr(&addr_QScrollArea_event_default, "QScrollArea_event_default")
	symaddr(&addr_QScrollArea_eventFilter_default, "QScrollArea_eventFilter_default")
	symaddr(&addr_QScrollArea_resizeEvent_default, "QScrollArea_resizeEvent_default")
	symaddr(&addr_QScrollArea_scrollContentsBy_default, "QScrollArea_scrollContentsBy_default")
	symaddr(&addr_QScrollArea_viewportSizeHint_default, "QScrollArea_viewportSizeHint_default")
	symaddr(&addr_QScrollArea_minimumSizeHint_default, "QScrollArea_minimumSizeHint_default")
	symaddr(&addr_QScrollArea_setupViewport_default, "QScrollArea_setupViewport_default")
	symaddr(&addr_QScrollArea_viewportEvent_default, "QScrollArea_viewportEvent_default")
	symaddr(&addr_QScrollArea_paintEvent_default, "QScrollArea_paintEvent_default")
	symaddr(&addr_QScrollArea_mousePressEvent_default, "QScrollArea_mousePressEvent_default")
	symaddr(&addr_QScrollArea_mouseReleaseEvent_default, "QScrollArea_mouseReleaseEvent_default")
	symaddr(&addr_QScrollArea_mouseDoubleClickEvent_default, "QScrollArea_mouseDoubleClickEvent_default")
	symaddr(&addr_QScrollArea_mouseMoveEvent_default, "QScrollArea_mouseMoveEvent_default")
	symaddr(&addr_QScrollArea_wheelEvent_default, "QScrollArea_wheelEvent_default")
	symaddr(&addr_QScrollArea_contextMenuEvent_default, "QScrollArea_contextMenuEvent_default")
	symaddr(&addr_QScrollArea_dragEnterEvent_default, "QScrollArea_dragEnterEvent_default")
	symaddr(&addr_QScrollArea_dragMoveEvent_default, "QScrollArea_dragMoveEvent_default")
	symaddr(&addr_QScrollArea_dragLeaveEvent_default, "QScrollArea_dragLeaveEvent_default")
	symaddr(&addr_QScrollArea_dropEvent_default, "QScrollArea_dropEvent_default")
	symaddr(&addr_QScrollArea_keyPressEvent_default, "QScrollArea_keyPressEvent_default")
	symaddr(&addr_QScrollArea_changeEvent_default, "QScrollArea_changeEvent_default")
	symaddr(&addr_QScrollArea_initStyleOption_default, "QScrollArea_initStyleOption_default")
	symaddr(&addr_QScrollArea_devType_default, "QScrollArea_devType_default")
	symaddr(&addr_QScrollArea_setVisible_default, "QScrollArea_setVisible_default")
	symaddr(&addr_QScrollArea_heightForWidth_default, "QScrollArea_heightForWidth_default")
	symaddr(&addr_QScrollArea_hasHeightForWidth_default, "QScrollArea_hasHeightForWidth_default")
	symaddr(&addr_QScrollArea_paintEngine_default, "QScrollArea_paintEngine_default")
	symaddr(&addr_QScrollArea_keyReleaseEvent_default, "QScrollArea_keyReleaseEvent_default")
	symaddr(&addr_QScrollArea_focusInEvent_default, "QScrollArea_focusInEvent_default")
	symaddr(&addr_QScrollArea_focusOutEvent_default, "QScrollArea_focusOutEvent_default")
	symaddr(&addr_QScrollArea_enterEvent_default, "QScrollArea_enterEvent_default")
	symaddr(&addr_QScrollArea_leaveEvent_default, "QScrollArea_leaveEvent_default")
	symaddr(&addr_QScrollArea_moveEvent_default, "QScrollArea_moveEvent_default")
	symaddr(&addr_QScrollArea_closeEvent_default, "QScrollArea_closeEvent_default")
	symaddr(&addr_QScrollArea_tabletEvent_default, "QScrollArea_tabletEvent_default")
	symaddr(&addr_QScrollArea_actionEvent_default, "QScrollArea_actionEvent_default")
	symaddr(&addr_QScrollArea_showEvent_default, "QScrollArea_showEvent_default")
	symaddr(&addr_QScrollArea_hideEvent_default, "QScrollArea_hideEvent_default")
	symaddr(&addr_QScrollArea_metric_default, "QScrollArea_metric_default")
	symaddr(&addr_QScrollArea_initPainter_default, "QScrollArea_initPainter_default")
	symaddr(&addr_QScrollArea_redirected_default, "QScrollArea_redirected_default")
	symaddr(&addr_QScrollArea_sharedPainter_default, "QScrollArea_sharedPainter_default")
	symaddr(&addr_QScrollArea_inputMethodEvent_default, "QScrollArea_inputMethodEvent_default")
	symaddr(&addr_QScrollArea_inputMethodQuery_default, "QScrollArea_inputMethodQuery_default")
	symaddr(&addr_QScrollArea_timerEvent_default, "QScrollArea_timerEvent_default")
	symaddr(&addr_QScrollArea_childEvent_default, "QScrollArea_childEvent_default")
	symaddr(&addr_QScrollArea_customEvent_default, "QScrollArea_customEvent_default")
	symaddr(&addr_QScrollArea_connectNotify_default, "QScrollArea_connectNotify_default")
	symaddr(&addr_QScrollArea_disconnectNotify_default, "QScrollArea_disconnectNotify_default")
	reg(&pg_QScrollArea_connect_windowTitleChanged, "QScrollArea_connect_windowTitleChanged")
	reg(&pg_QScrollArea_connect_windowIconChanged, "QScrollArea_connect_windowIconChanged")
	reg(&pg_QScrollArea_connect_windowIconTextChanged, "QScrollArea_connect_windowIconTextChanged")
	reg(&pg_QScrollArea_connect_customContextMenuRequested, "QScrollArea_connect_customContextMenuRequested")
	reg(&pg_QScrollArea_connect_destroyed, "QScrollArea_connect_destroyed")
	reg(&pg_QScrollArea_connect_destroyed2, "QScrollArea_connect_destroyed2")
	reg(&pg_QScrollArea_connect_objectNameChanged, "QScrollArea_connect_objectNameChanged")
	reg(&pg_QScrollArea__disconnect, "QScrollArea__disconnect")
	reg(&pg_QScrollArea__set_override, "QScrollArea__set_override")
	reg(&pg_QScrollArea_delete, "QScrollArea_delete")
}

type QScrollArea struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQScrollArea(h uintptr) *QScrollArea {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QScrollArea_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QScrollArea{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQScrollArea(h uintptr) *QScrollArea { return bindQScrollArea(h) }

func (this *QScrollArea) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScrollArea) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScrollArea(h unsafe.Pointer) *QScrollArea { return bindQScrollArea(uintptr(h)) }
func UnsafeNewQScrollAreaOwned(h unsafe.Pointer) *QScrollArea { return newownedQScrollArea(uintptr(h)) }

func (this *QScrollArea) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func NewQScrollArea(parent *QWidget) *QScrollArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollArea_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQScrollArea(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollArea_virtIdx, QScrollArea_enable)
	return __o
}

func NewQScrollArea2() *QScrollArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QScrollArea_new2, uintptr(__goid))
	__o := newownedQScrollArea(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QScrollArea_virtIdx, QScrollArea_enable)
	return __o
}

func QScrollArea_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollArea_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QScrollArea) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QScrollArea) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QScrollArea_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QScrollArea) TakeWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_takeWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QScrollArea) WidgetResizable() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_widgetResizable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollArea) SetWidgetResizable(resizable bool) {
	purego.SyscallN(addr_QScrollArea_setWidgetResizable, this.h, uintptr(b2i(resizable)))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QScrollArea) SetAlignment(arg1 AlignmentFlag) {
	purego.SyscallN(addr_QScrollArea_setAlignment, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) EnsureVisible(x int, y int) {
	purego.SyscallN(addr_QScrollArea_ensureVisible, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) EnsureWidgetVisible(childWidget *QWidget) {
	purego.SyscallN(addr_QScrollArea_ensureWidgetVisible, this.h, uintptr(childWidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childWidget)
}

func (this *QScrollArea) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QScrollArea_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QScrollArea_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QScrollArea) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QScrollArea) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QScrollArea_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QScrollArea_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QScrollArea_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QScrollArea_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_mouseGrabber)
	return bindQWidget(__r)
}

func QScrollArea_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_keyboardGrabber)
	return bindQWidget(__r)
}

func QScrollArea_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QScrollArea_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QScrollArea) UpdateMicroFocus() {
	purego.SyscallN(addr_QScrollArea_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Create() {
	purego.SyscallN(addr_QScrollArea_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Destroy() {
	purego.SyscallN(addr_QScrollArea_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QScrollArea) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollArea) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QScrollArea_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScrollArea_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScrollArea_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScrollArea) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QScrollArea) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollArea) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QScrollArea) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QScrollArea_DevicePixelRatioFScale() float64 {
	__r := pg_QScrollArea_devicePixelRatioFScale()
	return __r
}

func QScrollArea_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QScrollArea_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QScrollArea) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QScrollArea_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QScrollArea_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollArea_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QScrollArea_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScrollArea_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QScrollArea) EnsureVisible2(x int, y int, xmargin int) {
	purego.SyscallN(addr_QScrollArea_ensureVisible2, this.h, uintptr(x), uintptr(y), uintptr(xmargin))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) EnsureVisible3(x int, y int, xmargin int, ymargin int) {
	purego.SyscallN(addr_QScrollArea_ensureVisible3, this.h, uintptr(x), uintptr(y), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) EnsureWidgetVisible2(childWidget *QWidget, xmargin int) {
	purego.SyscallN(addr_QScrollArea_ensureWidgetVisible2, this.h, uintptr(childWidget.UnsafePointer()), uintptr(xmargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childWidget)
}

func (this *QScrollArea) EnsureWidgetVisible3(childWidget *QWidget, xmargin int, ymargin int) {
	purego.SyscallN(addr_QScrollArea_ensureWidgetVisible3, this.h, uintptr(childWidget.UnsafePointer()), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childWidget)
}

func QScrollArea_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QScrollArea_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QScrollArea) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QScrollArea_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QScrollArea_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QScrollArea_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QScrollArea_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QScrollArea_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QScrollArea_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QScrollArea_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScrollArea_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScrollArea_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScrollArea) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QScrollArea) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScrollArea) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollArea) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QScrollArea) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QScrollArea) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QScrollArea_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QScrollArea_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScrollArea) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScrollArea) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QScrollArea_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QScrollArea) ViewportEventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_viewportEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QScrollArea) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QScrollArea_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollArea_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollArea_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollArea_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QScrollArea_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QScrollArea_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QScrollArea_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) DragEnterEventDefault(arg1 *QDragEnterEvent) {
	purego.SyscallN(addr_QScrollArea_dragEnterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) DragMoveEventDefault(arg1 *QDragMoveEvent) {
	purego.SyscallN(addr_QScrollArea_dragMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) DragLeaveEventDefault(arg1 *QDragLeaveEvent) {
	purego.SyscallN(addr_QScrollArea_dragLeaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) DropEventDefault(arg1 *QDropEvent) {
	purego.SyscallN(addr_QScrollArea_dropEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QScrollArea_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QScrollArea_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QScrollArea_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QScrollArea) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollArea) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QScrollArea_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QScrollArea) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollArea) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScrollArea) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QScrollArea) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QScrollArea_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QScrollArea_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QScrollArea_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QScrollArea_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QScrollArea_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QScrollArea_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QScrollArea_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QScrollArea_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QScrollArea_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QScrollArea_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QScrollArea_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollArea_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScrollArea) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QScrollArea_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QScrollArea) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QScrollArea) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QScrollArea) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QScrollArea_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QScrollArea) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QScrollArea_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QScrollArea) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QScrollArea_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QScrollArea_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QScrollArea_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QScrollArea) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QScrollArea_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QScrollArea) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QScrollArea_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QScrollArea) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QScrollArea_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QScrollArea_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QScrollArea_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QScrollArea_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScrollArea_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QScrollArea_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QScrollArea_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScrollArea) Disconnect(handle int64) {
	pg_QScrollArea__disconnect(handle)
	unregisterCallback(handle)
}

func QScrollArea_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 5898328517505784191:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
		return 4
	case -4302917285655014661:
		return 5
	case -6118001150126564745:
		return 6
	case 503579067252048239:
		return 7
	case 7049264599837532491:
		return 8
	case 4531996381736730688:
		return 9
	case -1971465653061861519:
		return 10
	case 5290442163167458465:
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
	case -6698043222227006015:
		return 17
	case 206122715919917589:
		return 18
	case -1781545137352745978:
		return 19
	case -3139904158373725006:
		return 20
	case 9107058682754591716:
		return 21
	case 1023791977679793185:
		return 22
	case 4959376063826165943:
		return 23
	case -488693769771308531:
		return 24
	case -8659679893868524044:
		return 25
	case 2110182719375245575:
		return 26
	case 4311534060495737339:
		return 27
	case -8266762397995423915:
		return 28
	case 6398007404756346381:
		return 29
	case 772598475950050299:
		return 30
	case 975043767033552552:
		return 31
	case -385596010576048993:
		return 32
	case 6316270346668237895:
		return 33
	case -551070009623711640:
		return 34
	case -9161949150489525396:
		return 35
	case -1989937050019464783:
		return 36
	case -212068866860701141:
		return 37
	case 1697444140088555581:
		return 38
	case -6825369174751916:
		return 39
	case -199718707876719129:
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
	case -5409927581812374498:
		return 48
	case -1185315367526857833:
		return 49
	case 3769868229123868474:
		return 50
	case 7379029982170503008:
		return 51
	case -568369772982215962:
		return 52
	}
	return -1
}

func QScrollArea_enable(h uintptr, idx int32) { pg_QScrollArea__set_override(h, idx) }

func (this *QScrollArea) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QScrollArea) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QScrollArea) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QScrollArea) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QScrollArea) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QScrollArea) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QScrollArea) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QScrollArea) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QScrollArea) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QScrollArea) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QScrollArea) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "viewportEvent"
}

func (this *QScrollArea) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QScrollArea) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QScrollArea) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QScrollArea) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QScrollArea) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QScrollArea) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QScrollArea) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QScrollArea) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(arg1 uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(arg1))) }) // slot "dragEnterEvent"
}

func (this *QScrollArea) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(arg1 uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(arg1))) }) // slot "dragMoveEvent"
}

func (this *QScrollArea) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(arg1 uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(arg1))) }) // slot "dragLeaveEvent"
}

func (this *QScrollArea) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(arg1 uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(arg1))) }) // slot "dropEvent"
}

func (this *QScrollArea) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QScrollArea) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QScrollArea) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QScrollArea) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QScrollArea) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QScrollArea) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QScrollArea) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QScrollArea) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QScrollArea) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QScrollArea) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QScrollArea) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QScrollArea) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QScrollArea) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QScrollArea) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QScrollArea) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QScrollArea) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QScrollArea) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QScrollArea) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QScrollArea) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QScrollArea) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QScrollArea) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QScrollArea) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QScrollArea) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QScrollArea) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QScrollArea) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QScrollArea) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QScrollArea) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QScrollArea) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QScrollArea) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QScrollArea) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QScrollArea) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QScrollArea) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScrollArea_delete(this.h)
	}
	this.UnsafeRelease()
}
