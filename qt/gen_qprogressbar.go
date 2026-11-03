package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProgressBar_chain uintptr
	addr_QProgressBar_new uintptr
	addr_QProgressBar_new2 uintptr
	addr_QProgressBar_tr uintptr
	addr_QProgressBar_minimum uintptr
	addr_QProgressBar_maximum uintptr
	addr_QProgressBar_value uintptr
	addr_QProgressBar_text uintptr
	addr_QProgressBar_setTextVisible uintptr
	addr_QProgressBar_isTextVisible uintptr
	addr_QProgressBar_alignment uintptr
	addr_QProgressBar_setAlignment uintptr
	addr_QProgressBar_orientation uintptr
	addr_QProgressBar_setInvertedAppearance uintptr
	addr_QProgressBar_invertedAppearance uintptr
	addr_QProgressBar_setTextDirection uintptr
	addr_QProgressBar_textDirection uintptr
	addr_QProgressBar_setFormat uintptr
	addr_QProgressBar_resetFormat uintptr
	addr_QProgressBar_format uintptr
	addr_QProgressBar_reset uintptr
	addr_QProgressBar_setRange uintptr
	addr_QProgressBar_setMinimum uintptr
	addr_QProgressBar_setMaximum uintptr
	addr_QProgressBar_setValue uintptr
	addr_QProgressBar_setOrientation uintptr
	addr_QProgressBar_valueChanged uintptr
	addr_QProgressBar_setTabOrder uintptr
	addr_QProgressBar_mouseGrabber uintptr
	addr_QProgressBar_keyboardGrabber uintptr
	addr_QProgressBar_find uintptr
	addr_QProgressBar_createWindowContainer uintptr
	addr_QProgressBar_updateMicroFocus uintptr
	addr_QProgressBar_create uintptr
	addr_QProgressBar_destroy uintptr
	addr_QProgressBar_focusNextChild uintptr
	addr_QProgressBar_focusPreviousChild uintptr
	addr_QProgressBar_connect uintptr
	addr_QProgressBar_connect2 uintptr
	addr_QProgressBar_sender uintptr
	addr_QProgressBar_senderSignalIndex uintptr
	addr_QProgressBar_receivers uintptr
	addr_QProgressBar_isSignalConnected uintptr
	pg_QProgressBar_devicePixelRatioFScale func() float64
	pg_QProgressBar_encodeMetricF func(int32, float64) int32
	pg_QProgressBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QProgressBar_tr2 uintptr
	addr_QProgressBar_tr3 uintptr
	addr_QProgressBar_createWindowContainer2 uintptr
	addr_QProgressBar_createWindowContainer3 uintptr
	addr_QProgressBar_updateMicroFocus2 uintptr
	addr_QProgressBar_create2 uintptr
	addr_QProgressBar_create3 uintptr
	addr_QProgressBar_create4 uintptr
	addr_QProgressBar_destroy2 uintptr
	addr_QProgressBar_destroy3 uintptr
	addr_QProgressBar_connect3 uintptr
	addr_QProgressBar_connect4 uintptr
	addr_QProgressBar_text_default uintptr
	addr_QProgressBar_metaObject_default uintptr
	addr_QProgressBar_sizeHint_default uintptr
	addr_QProgressBar_minimumSizeHint_default uintptr
	addr_QProgressBar_event_default uintptr
	addr_QProgressBar_paintEvent_default uintptr
	addr_QProgressBar_initStyleOption_default uintptr
	addr_QProgressBar_devType_default uintptr
	addr_QProgressBar_setVisible_default uintptr
	addr_QProgressBar_heightForWidth_default uintptr
	addr_QProgressBar_hasHeightForWidth_default uintptr
	addr_QProgressBar_paintEngine_default uintptr
	addr_QProgressBar_mousePressEvent_default uintptr
	addr_QProgressBar_mouseReleaseEvent_default uintptr
	addr_QProgressBar_mouseDoubleClickEvent_default uintptr
	addr_QProgressBar_mouseMoveEvent_default uintptr
	addr_QProgressBar_wheelEvent_default uintptr
	addr_QProgressBar_keyPressEvent_default uintptr
	addr_QProgressBar_keyReleaseEvent_default uintptr
	addr_QProgressBar_focusInEvent_default uintptr
	addr_QProgressBar_focusOutEvent_default uintptr
	addr_QProgressBar_enterEvent_default uintptr
	addr_QProgressBar_leaveEvent_default uintptr
	addr_QProgressBar_moveEvent_default uintptr
	addr_QProgressBar_resizeEvent_default uintptr
	addr_QProgressBar_closeEvent_default uintptr
	addr_QProgressBar_contextMenuEvent_default uintptr
	addr_QProgressBar_tabletEvent_default uintptr
	addr_QProgressBar_actionEvent_default uintptr
	addr_QProgressBar_dragEnterEvent_default uintptr
	addr_QProgressBar_dragMoveEvent_default uintptr
	addr_QProgressBar_dragLeaveEvent_default uintptr
	addr_QProgressBar_dropEvent_default uintptr
	addr_QProgressBar_showEvent_default uintptr
	addr_QProgressBar_hideEvent_default uintptr
	addr_QProgressBar_changeEvent_default uintptr
	addr_QProgressBar_metric_default uintptr
	addr_QProgressBar_initPainter_default uintptr
	addr_QProgressBar_redirected_default uintptr
	addr_QProgressBar_sharedPainter_default uintptr
	addr_QProgressBar_inputMethodEvent_default uintptr
	addr_QProgressBar_inputMethodQuery_default uintptr
	addr_QProgressBar_focusNextPrevChild_default uintptr
	addr_QProgressBar_eventFilter_default uintptr
	addr_QProgressBar_timerEvent_default uintptr
	addr_QProgressBar_childEvent_default uintptr
	addr_QProgressBar_customEvent_default uintptr
	addr_QProgressBar_connectNotify_default uintptr
	addr_QProgressBar_disconnectNotify_default uintptr
	pg_QProgressBar_connect_valueChanged func(uintptr, int64)
	pg_QProgressBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QProgressBar_connect_windowIconChanged func(uintptr, int64)
	pg_QProgressBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QProgressBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QProgressBar_connect_destroyed func(uintptr, int64)
	pg_QProgressBar_connect_destroyed2 func(uintptr, int64)
	pg_QProgressBar_connect_objectNameChanged func(uintptr, int64)
	pg_QProgressBar__disconnect func(int64)
	pg_QProgressBar__set_override func(uintptr, int32)
	pg_QProgressBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProgressBar_chain, "QProgressBar_chain")
	symaddr(&addr_QProgressBar_new, "QProgressBar_new")
	symaddr(&addr_QProgressBar_new2, "QProgressBar_new2")
	symaddr(&addr_QProgressBar_tr, "QProgressBar_tr")
	symaddr(&addr_QProgressBar_minimum, "QProgressBar_minimum")
	symaddr(&addr_QProgressBar_maximum, "QProgressBar_maximum")
	symaddr(&addr_QProgressBar_value, "QProgressBar_value")
	symaddr(&addr_QProgressBar_text, "QProgressBar_text")
	symaddr(&addr_QProgressBar_setTextVisible, "QProgressBar_setTextVisible")
	symaddr(&addr_QProgressBar_isTextVisible, "QProgressBar_isTextVisible")
	symaddr(&addr_QProgressBar_alignment, "QProgressBar_alignment")
	symaddr(&addr_QProgressBar_setAlignment, "QProgressBar_setAlignment")
	symaddr(&addr_QProgressBar_orientation, "QProgressBar_orientation")
	symaddr(&addr_QProgressBar_setInvertedAppearance, "QProgressBar_setInvertedAppearance")
	symaddr(&addr_QProgressBar_invertedAppearance, "QProgressBar_invertedAppearance")
	symaddr(&addr_QProgressBar_setTextDirection, "QProgressBar_setTextDirection")
	symaddr(&addr_QProgressBar_textDirection, "QProgressBar_textDirection")
	symaddr(&addr_QProgressBar_setFormat, "QProgressBar_setFormat")
	symaddr(&addr_QProgressBar_resetFormat, "QProgressBar_resetFormat")
	symaddr(&addr_QProgressBar_format, "QProgressBar_format")
	symaddr(&addr_QProgressBar_reset, "QProgressBar_reset")
	symaddr(&addr_QProgressBar_setRange, "QProgressBar_setRange")
	symaddr(&addr_QProgressBar_setMinimum, "QProgressBar_setMinimum")
	symaddr(&addr_QProgressBar_setMaximum, "QProgressBar_setMaximum")
	symaddr(&addr_QProgressBar_setValue, "QProgressBar_setValue")
	symaddr(&addr_QProgressBar_setOrientation, "QProgressBar_setOrientation")
	symaddr(&addr_QProgressBar_valueChanged, "QProgressBar_valueChanged")
	symaddr(&addr_QProgressBar_setTabOrder, "QProgressBar_setTabOrder")
	symaddr(&addr_QProgressBar_mouseGrabber, "QProgressBar_mouseGrabber")
	symaddr(&addr_QProgressBar_keyboardGrabber, "QProgressBar_keyboardGrabber")
	symaddr(&addr_QProgressBar_find, "QProgressBar_find")
	symaddr(&addr_QProgressBar_createWindowContainer, "QProgressBar_createWindowContainer")
	symaddr(&addr_QProgressBar_updateMicroFocus, "QProgressBar_updateMicroFocus")
	symaddr(&addr_QProgressBar_create, "QProgressBar_create")
	symaddr(&addr_QProgressBar_destroy, "QProgressBar_destroy")
	symaddr(&addr_QProgressBar_focusNextChild, "QProgressBar_focusNextChild")
	symaddr(&addr_QProgressBar_focusPreviousChild, "QProgressBar_focusPreviousChild")
	symaddr(&addr_QProgressBar_connect, "QProgressBar_connect")
	symaddr(&addr_QProgressBar_connect2, "QProgressBar_connect2")
	symaddr(&addr_QProgressBar_sender, "QProgressBar_sender")
	symaddr(&addr_QProgressBar_senderSignalIndex, "QProgressBar_senderSignalIndex")
	symaddr(&addr_QProgressBar_receivers, "QProgressBar_receivers")
	symaddr(&addr_QProgressBar_isSignalConnected, "QProgressBar_isSignalConnected")
	reg(&pg_QProgressBar_devicePixelRatioFScale, "QProgressBar_devicePixelRatioFScale")
	reg(&pg_QProgressBar_encodeMetricF, "QProgressBar_encodeMetricF")
	reg(&pg_QProgressBar_getDecodedMetricF, "QProgressBar_getDecodedMetricF")
	symaddr(&addr_QProgressBar_tr2, "QProgressBar_tr2")
	symaddr(&addr_QProgressBar_tr3, "QProgressBar_tr3")
	symaddr(&addr_QProgressBar_createWindowContainer2, "QProgressBar_createWindowContainer2")
	symaddr(&addr_QProgressBar_createWindowContainer3, "QProgressBar_createWindowContainer3")
	symaddr(&addr_QProgressBar_updateMicroFocus2, "QProgressBar_updateMicroFocus2")
	symaddr(&addr_QProgressBar_create2, "QProgressBar_create2")
	symaddr(&addr_QProgressBar_create3, "QProgressBar_create3")
	symaddr(&addr_QProgressBar_create4, "QProgressBar_create4")
	symaddr(&addr_QProgressBar_destroy2, "QProgressBar_destroy2")
	symaddr(&addr_QProgressBar_destroy3, "QProgressBar_destroy3")
	symaddr(&addr_QProgressBar_connect3, "QProgressBar_connect3")
	symaddr(&addr_QProgressBar_connect4, "QProgressBar_connect4")
	symaddr(&addr_QProgressBar_text_default, "QProgressBar_text_default")
	symaddr(&addr_QProgressBar_metaObject_default, "QProgressBar_metaObject_default")
	symaddr(&addr_QProgressBar_sizeHint_default, "QProgressBar_sizeHint_default")
	symaddr(&addr_QProgressBar_minimumSizeHint_default, "QProgressBar_minimumSizeHint_default")
	symaddr(&addr_QProgressBar_event_default, "QProgressBar_event_default")
	symaddr(&addr_QProgressBar_paintEvent_default, "QProgressBar_paintEvent_default")
	symaddr(&addr_QProgressBar_initStyleOption_default, "QProgressBar_initStyleOption_default")
	symaddr(&addr_QProgressBar_devType_default, "QProgressBar_devType_default")
	symaddr(&addr_QProgressBar_setVisible_default, "QProgressBar_setVisible_default")
	symaddr(&addr_QProgressBar_heightForWidth_default, "QProgressBar_heightForWidth_default")
	symaddr(&addr_QProgressBar_hasHeightForWidth_default, "QProgressBar_hasHeightForWidth_default")
	symaddr(&addr_QProgressBar_paintEngine_default, "QProgressBar_paintEngine_default")
	symaddr(&addr_QProgressBar_mousePressEvent_default, "QProgressBar_mousePressEvent_default")
	symaddr(&addr_QProgressBar_mouseReleaseEvent_default, "QProgressBar_mouseReleaseEvent_default")
	symaddr(&addr_QProgressBar_mouseDoubleClickEvent_default, "QProgressBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QProgressBar_mouseMoveEvent_default, "QProgressBar_mouseMoveEvent_default")
	symaddr(&addr_QProgressBar_wheelEvent_default, "QProgressBar_wheelEvent_default")
	symaddr(&addr_QProgressBar_keyPressEvent_default, "QProgressBar_keyPressEvent_default")
	symaddr(&addr_QProgressBar_keyReleaseEvent_default, "QProgressBar_keyReleaseEvent_default")
	symaddr(&addr_QProgressBar_focusInEvent_default, "QProgressBar_focusInEvent_default")
	symaddr(&addr_QProgressBar_focusOutEvent_default, "QProgressBar_focusOutEvent_default")
	symaddr(&addr_QProgressBar_enterEvent_default, "QProgressBar_enterEvent_default")
	symaddr(&addr_QProgressBar_leaveEvent_default, "QProgressBar_leaveEvent_default")
	symaddr(&addr_QProgressBar_moveEvent_default, "QProgressBar_moveEvent_default")
	symaddr(&addr_QProgressBar_resizeEvent_default, "QProgressBar_resizeEvent_default")
	symaddr(&addr_QProgressBar_closeEvent_default, "QProgressBar_closeEvent_default")
	symaddr(&addr_QProgressBar_contextMenuEvent_default, "QProgressBar_contextMenuEvent_default")
	symaddr(&addr_QProgressBar_tabletEvent_default, "QProgressBar_tabletEvent_default")
	symaddr(&addr_QProgressBar_actionEvent_default, "QProgressBar_actionEvent_default")
	symaddr(&addr_QProgressBar_dragEnterEvent_default, "QProgressBar_dragEnterEvent_default")
	symaddr(&addr_QProgressBar_dragMoveEvent_default, "QProgressBar_dragMoveEvent_default")
	symaddr(&addr_QProgressBar_dragLeaveEvent_default, "QProgressBar_dragLeaveEvent_default")
	symaddr(&addr_QProgressBar_dropEvent_default, "QProgressBar_dropEvent_default")
	symaddr(&addr_QProgressBar_showEvent_default, "QProgressBar_showEvent_default")
	symaddr(&addr_QProgressBar_hideEvent_default, "QProgressBar_hideEvent_default")
	symaddr(&addr_QProgressBar_changeEvent_default, "QProgressBar_changeEvent_default")
	symaddr(&addr_QProgressBar_metric_default, "QProgressBar_metric_default")
	symaddr(&addr_QProgressBar_initPainter_default, "QProgressBar_initPainter_default")
	symaddr(&addr_QProgressBar_redirected_default, "QProgressBar_redirected_default")
	symaddr(&addr_QProgressBar_sharedPainter_default, "QProgressBar_sharedPainter_default")
	symaddr(&addr_QProgressBar_inputMethodEvent_default, "QProgressBar_inputMethodEvent_default")
	symaddr(&addr_QProgressBar_inputMethodQuery_default, "QProgressBar_inputMethodQuery_default")
	symaddr(&addr_QProgressBar_focusNextPrevChild_default, "QProgressBar_focusNextPrevChild_default")
	symaddr(&addr_QProgressBar_eventFilter_default, "QProgressBar_eventFilter_default")
	symaddr(&addr_QProgressBar_timerEvent_default, "QProgressBar_timerEvent_default")
	symaddr(&addr_QProgressBar_childEvent_default, "QProgressBar_childEvent_default")
	symaddr(&addr_QProgressBar_customEvent_default, "QProgressBar_customEvent_default")
	symaddr(&addr_QProgressBar_connectNotify_default, "QProgressBar_connectNotify_default")
	symaddr(&addr_QProgressBar_disconnectNotify_default, "QProgressBar_disconnectNotify_default")
	reg(&pg_QProgressBar_connect_valueChanged, "QProgressBar_connect_valueChanged")
	reg(&pg_QProgressBar_connect_windowTitleChanged, "QProgressBar_connect_windowTitleChanged")
	reg(&pg_QProgressBar_connect_windowIconChanged, "QProgressBar_connect_windowIconChanged")
	reg(&pg_QProgressBar_connect_windowIconTextChanged, "QProgressBar_connect_windowIconTextChanged")
	reg(&pg_QProgressBar_connect_customContextMenuRequested, "QProgressBar_connect_customContextMenuRequested")
	reg(&pg_QProgressBar_connect_destroyed, "QProgressBar_connect_destroyed")
	reg(&pg_QProgressBar_connect_destroyed2, "QProgressBar_connect_destroyed2")
	reg(&pg_QProgressBar_connect_objectNameChanged, "QProgressBar_connect_objectNameChanged")
	reg(&pg_QProgressBar__disconnect, "QProgressBar__disconnect")
	reg(&pg_QProgressBar__set_override, "QProgressBar__set_override")
	reg(&pg_QProgressBar_delete, "QProgressBar_delete")
}

type QProgressBar struct {
	h uintptr
	*QWidget
}

func bindQProgressBar(h uintptr) *QProgressBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QProgressBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QProgressBar{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQProgressBar(h uintptr) *QProgressBar { return bindQProgressBar(h) }

func (this *QProgressBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProgressBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProgressBar(h unsafe.Pointer) *QProgressBar { return bindQProgressBar(uintptr(h)) }
func UnsafeNewQProgressBarOwned(h unsafe.Pointer) *QProgressBar { return newownedQProgressBar(uintptr(h)) }

func (this *QProgressBar) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQProgressBar(parent *QWidget) *QProgressBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProgressBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQProgressBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressBar_virtIdx, QProgressBar_enable)
	return __o
}

func NewQProgressBar2() *QProgressBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProgressBar_new2, uintptr(__goid))
	__o := newownedQProgressBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProgressBar_virtIdx, QProgressBar_enable)
	return __o
}

func QProgressBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QProgressBar) Minimum() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_minimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) Maximum() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_maximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) Value() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProgressBar) SetTextVisible(visible bool) {
	purego.SyscallN(addr_QProgressBar_setTextVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) IsTextVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_isTextVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressBar) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QProgressBar) SetAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QProgressBar_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QProgressBar) SetInvertedAppearance(invert bool) {
	purego.SyscallN(addr_QProgressBar_setInvertedAppearance, this.h, uintptr(b2i(invert)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) InvertedAppearance() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_invertedAppearance, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressBar) SetTextDirection(textDirection QProgressBar__Direction) {
	purego.SyscallN(addr_QProgressBar_setTextDirection, this.h, uintptr(textDirection))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) TextDirection() QProgressBar__Direction {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_textDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProgressBar__Direction(__r)
}

func (this *QProgressBar) SetFormat(format string) {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	purego.SyscallN(addr_QProgressBar_setFormat, this.h, format_p, uintptr(len(format)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QProgressBar) ResetFormat() {
	purego.SyscallN(addr_QProgressBar_resetFormat, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Format() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_format, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProgressBar) Reset() {
	purego.SyscallN(addr_QProgressBar_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressBar) SetRange(minimum int, maximum int) {
	purego.SyscallN(addr_QProgressBar_setRange, this.h, uintptr(minimum), uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) SetMinimum(minimum int) {
	purego.SyscallN(addr_QProgressBar_setMinimum, this.h, uintptr(minimum))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) SetMaximum(maximum int) {
	purego.SyscallN(addr_QProgressBar_setMaximum, this.h, uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) SetValue(value int) {
	purego.SyscallN(addr_QProgressBar_setValue, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) SetOrientation(arg1 Orientation) {
	purego.SyscallN(addr_QProgressBar_setOrientation, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) ValueChanged(value int) {
	purego.SyscallN(addr_QProgressBar_valueChanged, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func QProgressBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QProgressBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QProgressBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_mouseGrabber)
	return bindQWidget(__r)
}

func QProgressBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QProgressBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QProgressBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QProgressBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QProgressBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Create() {
	purego.SyscallN(addr_QProgressBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Destroy() {
	purego.SyscallN(addr_QProgressBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QProgressBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QProgressBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProgressBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProgressBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProgressBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QProgressBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QProgressBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QProgressBar_DevicePixelRatioFScale() float64 {
	__r := pg_QProgressBar_devicePixelRatioFScale()
	return __r
}

func QProgressBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QProgressBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QProgressBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QProgressBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QProgressBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProgressBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProgressBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QProgressBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QProgressBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QProgressBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QProgressBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QProgressBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QProgressBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QProgressBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QProgressBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QProgressBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProgressBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProgressBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProgressBar) TextDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProgressBar_text_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProgressBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QProgressBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QProgressBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QProgressBar) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QProgressBar) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QProgressBar_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressBar) InitStyleOptionDefault(option *QStyleOptionProgressBar) {
	purego.SyscallN(addr_QProgressBar_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QProgressBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QProgressBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QProgressBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QProgressBar) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressBar_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressBar_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressBar_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QProgressBar_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QProgressBar_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QProgressBar_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QProgressBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QProgressBar_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QProgressBar_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QProgressBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProgressBar_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QProgressBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QProgressBar_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QProgressBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QProgressBar_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QProgressBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QProgressBar_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QProgressBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QProgressBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QProgressBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QProgressBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QProgressBar_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QProgressBar_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QProgressBar_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProgressBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QProgressBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QProgressBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QProgressBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QProgressBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QProgressBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QProgressBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QProgressBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QProgressBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProgressBar) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProgressBar_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QProgressBar) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QProgressBar_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QProgressBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProgressBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProgressBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProgressBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProgressBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProgressBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProgressBar) OnValueChanged(fn func(int)) int64 {
	id := registerCallback(func(value int32) { fn(int(value)) })
	pg_QProgressBar_connect_valueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QProgressBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QProgressBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QProgressBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QProgressBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProgressBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QProgressBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QProgressBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProgressBar) Disconnect(handle int64) {
	pg_QProgressBar__disconnect(handle)
	unregisterCallback(handle)
}

func QProgressBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -430950356557152130:
		return 1
	case 729690261457684101:
		return 2
	case 7049264599837532491:
		return 3
	case 1237236456312824991:
		return 4
	case 5290442163167458465:
		return 5
	case -488693769771308531:
		return 6
	case -8659679893868524044:
		return 7
	case 2110182719375245575:
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
	case 4959376063826165943:
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
	case -1185315367526857833:
		return 46
	case 3769868229123868474:
		return 47
	case 7379029982170503008:
		return 48
	case -568369772982215962:
		return 49
	}
	return -1
}

func QProgressBar_enable(h uintptr, idx int32) { pg_QProgressBar__set_override(h, idx) }

func (this *QProgressBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QProgressBar) OnText(fn func() string) {
	registerVirt(this.UnsafeGoID(), -430950356557152130, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "text"
}

func (this *QProgressBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QProgressBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QProgressBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QProgressBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QProgressBar) OnInitStyleOption(fn func(*QStyleOptionProgressBar)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionProgressBar(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QProgressBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QProgressBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QProgressBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QProgressBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QProgressBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QProgressBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QProgressBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QProgressBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QProgressBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QProgressBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QProgressBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QProgressBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QProgressBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QProgressBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QProgressBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QProgressBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QProgressBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QProgressBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QProgressBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QProgressBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QProgressBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QProgressBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QProgressBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QProgressBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QProgressBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QProgressBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QProgressBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QProgressBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QProgressBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QProgressBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QProgressBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QProgressBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QProgressBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QProgressBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QProgressBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QProgressBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QProgressBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QProgressBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QProgressBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QProgressBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QProgressBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QProgressBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QProgressBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QProgressBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProgressBar_delete(this.h)
	}
	this.UnsafeRelease()
}
