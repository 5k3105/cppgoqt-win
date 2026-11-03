package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDateEdit_chain uintptr
	addr_QDateEdit_new uintptr
	addr_QDateEdit_new2 uintptr
	addr_QDateEdit_new3 uintptr
	addr_QDateEdit_new4 uintptr
	addr_QDateEdit_tr uintptr
	addr_QDateEdit_userDateChanged uintptr
	addr_QDateEdit_lineEdit uintptr
	addr_QDateEdit_setLineEdit uintptr
	addr_QDateEdit_setTabOrder uintptr
	addr_QDateEdit_mouseGrabber uintptr
	addr_QDateEdit_keyboardGrabber uintptr
	addr_QDateEdit_find uintptr
	addr_QDateEdit_createWindowContainer uintptr
	addr_QDateEdit_updateMicroFocus uintptr
	addr_QDateEdit_create uintptr
	addr_QDateEdit_destroy uintptr
	addr_QDateEdit_focusNextChild uintptr
	addr_QDateEdit_focusPreviousChild uintptr
	addr_QDateEdit_connect uintptr
	addr_QDateEdit_connect2 uintptr
	addr_QDateEdit_sender uintptr
	addr_QDateEdit_senderSignalIndex uintptr
	addr_QDateEdit_receivers uintptr
	addr_QDateEdit_isSignalConnected uintptr
	pg_QDateEdit_devicePixelRatioFScale func() float64
	pg_QDateEdit_encodeMetricF func(int32, float64) int32
	pg_QDateEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDateEdit_tr2 uintptr
	addr_QDateEdit_tr3 uintptr
	addr_QDateEdit_createWindowContainer2 uintptr
	addr_QDateEdit_createWindowContainer3 uintptr
	addr_QDateEdit_updateMicroFocus2 uintptr
	addr_QDateEdit_create2 uintptr
	addr_QDateEdit_create3 uintptr
	addr_QDateEdit_create4 uintptr
	addr_QDateEdit_destroy2 uintptr
	addr_QDateEdit_destroy3 uintptr
	addr_QDateEdit_connect3 uintptr
	addr_QDateEdit_connect4 uintptr
	addr_QDateEdit_metaObject_default uintptr
	addr_QDateEdit_sizeHint_default uintptr
	addr_QDateEdit_clear_default uintptr
	addr_QDateEdit_stepBy_default uintptr
	addr_QDateEdit_event_default uintptr
	addr_QDateEdit_keyPressEvent_default uintptr
	addr_QDateEdit_wheelEvent_default uintptr
	addr_QDateEdit_focusInEvent_default uintptr
	addr_QDateEdit_focusNextPrevChild_default uintptr
	addr_QDateEdit_validate_default uintptr
	addr_QDateEdit_fixup_default uintptr
	addr_QDateEdit_dateTimeFromText_default uintptr
	addr_QDateEdit_textFromDateTime_default uintptr
	addr_QDateEdit_stepEnabled_default uintptr
	addr_QDateEdit_mousePressEvent_default uintptr
	addr_QDateEdit_paintEvent_default uintptr
	addr_QDateEdit_initStyleOption_default uintptr
	addr_QDateEdit_minimumSizeHint_default uintptr
	addr_QDateEdit_inputMethodQuery_default uintptr
	addr_QDateEdit_resizeEvent_default uintptr
	addr_QDateEdit_keyReleaseEvent_default uintptr
	addr_QDateEdit_focusOutEvent_default uintptr
	addr_QDateEdit_contextMenuEvent_default uintptr
	addr_QDateEdit_changeEvent_default uintptr
	addr_QDateEdit_closeEvent_default uintptr
	addr_QDateEdit_hideEvent_default uintptr
	addr_QDateEdit_mouseReleaseEvent_default uintptr
	addr_QDateEdit_mouseMoveEvent_default uintptr
	addr_QDateEdit_timerEvent_default uintptr
	addr_QDateEdit_showEvent_default uintptr
	addr_QDateEdit_devType_default uintptr
	addr_QDateEdit_setVisible_default uintptr
	addr_QDateEdit_heightForWidth_default uintptr
	addr_QDateEdit_hasHeightForWidth_default uintptr
	addr_QDateEdit_paintEngine_default uintptr
	addr_QDateEdit_mouseDoubleClickEvent_default uintptr
	addr_QDateEdit_enterEvent_default uintptr
	addr_QDateEdit_leaveEvent_default uintptr
	addr_QDateEdit_moveEvent_default uintptr
	addr_QDateEdit_tabletEvent_default uintptr
	addr_QDateEdit_actionEvent_default uintptr
	addr_QDateEdit_dragEnterEvent_default uintptr
	addr_QDateEdit_dragMoveEvent_default uintptr
	addr_QDateEdit_dragLeaveEvent_default uintptr
	addr_QDateEdit_dropEvent_default uintptr
	addr_QDateEdit_metric_default uintptr
	addr_QDateEdit_initPainter_default uintptr
	addr_QDateEdit_redirected_default uintptr
	addr_QDateEdit_sharedPainter_default uintptr
	addr_QDateEdit_inputMethodEvent_default uintptr
	addr_QDateEdit_eventFilter_default uintptr
	addr_QDateEdit_childEvent_default uintptr
	addr_QDateEdit_customEvent_default uintptr
	addr_QDateEdit_connectNotify_default uintptr
	addr_QDateEdit_disconnectNotify_default uintptr
	pg_QDateEdit_connect_userDateChanged func(uintptr, int64)
	pg_QDateEdit_connect_dateTimeChanged func(uintptr, int64)
	pg_QDateEdit_connect_timeChanged func(uintptr, int64)
	pg_QDateEdit_connect_dateChanged func(uintptr, int64)
	pg_QDateEdit_connect_editingFinished func(uintptr, int64)
	pg_QDateEdit_connect_returnPressed func(uintptr, int64)
	pg_QDateEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QDateEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QDateEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDateEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDateEdit_connect_destroyed func(uintptr, int64)
	pg_QDateEdit_connect_destroyed2 func(uintptr, int64)
	pg_QDateEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QDateEdit__disconnect func(int64)
	pg_QDateEdit__set_override func(uintptr, int32)
	pg_QDateEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDateEdit_chain, "QDateEdit_chain")
	symaddr(&addr_QDateEdit_new, "QDateEdit_new")
	symaddr(&addr_QDateEdit_new2, "QDateEdit_new2")
	symaddr(&addr_QDateEdit_new3, "QDateEdit_new3")
	symaddr(&addr_QDateEdit_new4, "QDateEdit_new4")
	symaddr(&addr_QDateEdit_tr, "QDateEdit_tr")
	symaddr(&addr_QDateEdit_userDateChanged, "QDateEdit_userDateChanged")
	symaddr(&addr_QDateEdit_lineEdit, "QDateEdit_lineEdit")
	symaddr(&addr_QDateEdit_setLineEdit, "QDateEdit_setLineEdit")
	symaddr(&addr_QDateEdit_setTabOrder, "QDateEdit_setTabOrder")
	symaddr(&addr_QDateEdit_mouseGrabber, "QDateEdit_mouseGrabber")
	symaddr(&addr_QDateEdit_keyboardGrabber, "QDateEdit_keyboardGrabber")
	symaddr(&addr_QDateEdit_find, "QDateEdit_find")
	symaddr(&addr_QDateEdit_createWindowContainer, "QDateEdit_createWindowContainer")
	symaddr(&addr_QDateEdit_updateMicroFocus, "QDateEdit_updateMicroFocus")
	symaddr(&addr_QDateEdit_create, "QDateEdit_create")
	symaddr(&addr_QDateEdit_destroy, "QDateEdit_destroy")
	symaddr(&addr_QDateEdit_focusNextChild, "QDateEdit_focusNextChild")
	symaddr(&addr_QDateEdit_focusPreviousChild, "QDateEdit_focusPreviousChild")
	symaddr(&addr_QDateEdit_connect, "QDateEdit_connect")
	symaddr(&addr_QDateEdit_connect2, "QDateEdit_connect2")
	symaddr(&addr_QDateEdit_sender, "QDateEdit_sender")
	symaddr(&addr_QDateEdit_senderSignalIndex, "QDateEdit_senderSignalIndex")
	symaddr(&addr_QDateEdit_receivers, "QDateEdit_receivers")
	symaddr(&addr_QDateEdit_isSignalConnected, "QDateEdit_isSignalConnected")
	reg(&pg_QDateEdit_devicePixelRatioFScale, "QDateEdit_devicePixelRatioFScale")
	reg(&pg_QDateEdit_encodeMetricF, "QDateEdit_encodeMetricF")
	reg(&pg_QDateEdit_getDecodedMetricF, "QDateEdit_getDecodedMetricF")
	symaddr(&addr_QDateEdit_tr2, "QDateEdit_tr2")
	symaddr(&addr_QDateEdit_tr3, "QDateEdit_tr3")
	symaddr(&addr_QDateEdit_createWindowContainer2, "QDateEdit_createWindowContainer2")
	symaddr(&addr_QDateEdit_createWindowContainer3, "QDateEdit_createWindowContainer3")
	symaddr(&addr_QDateEdit_updateMicroFocus2, "QDateEdit_updateMicroFocus2")
	symaddr(&addr_QDateEdit_create2, "QDateEdit_create2")
	symaddr(&addr_QDateEdit_create3, "QDateEdit_create3")
	symaddr(&addr_QDateEdit_create4, "QDateEdit_create4")
	symaddr(&addr_QDateEdit_destroy2, "QDateEdit_destroy2")
	symaddr(&addr_QDateEdit_destroy3, "QDateEdit_destroy3")
	symaddr(&addr_QDateEdit_connect3, "QDateEdit_connect3")
	symaddr(&addr_QDateEdit_connect4, "QDateEdit_connect4")
	symaddr(&addr_QDateEdit_metaObject_default, "QDateEdit_metaObject_default")
	symaddr(&addr_QDateEdit_sizeHint_default, "QDateEdit_sizeHint_default")
	symaddr(&addr_QDateEdit_clear_default, "QDateEdit_clear_default")
	symaddr(&addr_QDateEdit_stepBy_default, "QDateEdit_stepBy_default")
	symaddr(&addr_QDateEdit_event_default, "QDateEdit_event_default")
	symaddr(&addr_QDateEdit_keyPressEvent_default, "QDateEdit_keyPressEvent_default")
	symaddr(&addr_QDateEdit_wheelEvent_default, "QDateEdit_wheelEvent_default")
	symaddr(&addr_QDateEdit_focusInEvent_default, "QDateEdit_focusInEvent_default")
	symaddr(&addr_QDateEdit_focusNextPrevChild_default, "QDateEdit_focusNextPrevChild_default")
	symaddr(&addr_QDateEdit_validate_default, "QDateEdit_validate_default")
	symaddr(&addr_QDateEdit_fixup_default, "QDateEdit_fixup_default")
	symaddr(&addr_QDateEdit_dateTimeFromText_default, "QDateEdit_dateTimeFromText_default")
	symaddr(&addr_QDateEdit_textFromDateTime_default, "QDateEdit_textFromDateTime_default")
	symaddr(&addr_QDateEdit_stepEnabled_default, "QDateEdit_stepEnabled_default")
	symaddr(&addr_QDateEdit_mousePressEvent_default, "QDateEdit_mousePressEvent_default")
	symaddr(&addr_QDateEdit_paintEvent_default, "QDateEdit_paintEvent_default")
	symaddr(&addr_QDateEdit_initStyleOption_default, "QDateEdit_initStyleOption_default")
	symaddr(&addr_QDateEdit_minimumSizeHint_default, "QDateEdit_minimumSizeHint_default")
	symaddr(&addr_QDateEdit_inputMethodQuery_default, "QDateEdit_inputMethodQuery_default")
	symaddr(&addr_QDateEdit_resizeEvent_default, "QDateEdit_resizeEvent_default")
	symaddr(&addr_QDateEdit_keyReleaseEvent_default, "QDateEdit_keyReleaseEvent_default")
	symaddr(&addr_QDateEdit_focusOutEvent_default, "QDateEdit_focusOutEvent_default")
	symaddr(&addr_QDateEdit_contextMenuEvent_default, "QDateEdit_contextMenuEvent_default")
	symaddr(&addr_QDateEdit_changeEvent_default, "QDateEdit_changeEvent_default")
	symaddr(&addr_QDateEdit_closeEvent_default, "QDateEdit_closeEvent_default")
	symaddr(&addr_QDateEdit_hideEvent_default, "QDateEdit_hideEvent_default")
	symaddr(&addr_QDateEdit_mouseReleaseEvent_default, "QDateEdit_mouseReleaseEvent_default")
	symaddr(&addr_QDateEdit_mouseMoveEvent_default, "QDateEdit_mouseMoveEvent_default")
	symaddr(&addr_QDateEdit_timerEvent_default, "QDateEdit_timerEvent_default")
	symaddr(&addr_QDateEdit_showEvent_default, "QDateEdit_showEvent_default")
	symaddr(&addr_QDateEdit_devType_default, "QDateEdit_devType_default")
	symaddr(&addr_QDateEdit_setVisible_default, "QDateEdit_setVisible_default")
	symaddr(&addr_QDateEdit_heightForWidth_default, "QDateEdit_heightForWidth_default")
	symaddr(&addr_QDateEdit_hasHeightForWidth_default, "QDateEdit_hasHeightForWidth_default")
	symaddr(&addr_QDateEdit_paintEngine_default, "QDateEdit_paintEngine_default")
	symaddr(&addr_QDateEdit_mouseDoubleClickEvent_default, "QDateEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QDateEdit_enterEvent_default, "QDateEdit_enterEvent_default")
	symaddr(&addr_QDateEdit_leaveEvent_default, "QDateEdit_leaveEvent_default")
	symaddr(&addr_QDateEdit_moveEvent_default, "QDateEdit_moveEvent_default")
	symaddr(&addr_QDateEdit_tabletEvent_default, "QDateEdit_tabletEvent_default")
	symaddr(&addr_QDateEdit_actionEvent_default, "QDateEdit_actionEvent_default")
	symaddr(&addr_QDateEdit_dragEnterEvent_default, "QDateEdit_dragEnterEvent_default")
	symaddr(&addr_QDateEdit_dragMoveEvent_default, "QDateEdit_dragMoveEvent_default")
	symaddr(&addr_QDateEdit_dragLeaveEvent_default, "QDateEdit_dragLeaveEvent_default")
	symaddr(&addr_QDateEdit_dropEvent_default, "QDateEdit_dropEvent_default")
	symaddr(&addr_QDateEdit_metric_default, "QDateEdit_metric_default")
	symaddr(&addr_QDateEdit_initPainter_default, "QDateEdit_initPainter_default")
	symaddr(&addr_QDateEdit_redirected_default, "QDateEdit_redirected_default")
	symaddr(&addr_QDateEdit_sharedPainter_default, "QDateEdit_sharedPainter_default")
	symaddr(&addr_QDateEdit_inputMethodEvent_default, "QDateEdit_inputMethodEvent_default")
	symaddr(&addr_QDateEdit_eventFilter_default, "QDateEdit_eventFilter_default")
	symaddr(&addr_QDateEdit_childEvent_default, "QDateEdit_childEvent_default")
	symaddr(&addr_QDateEdit_customEvent_default, "QDateEdit_customEvent_default")
	symaddr(&addr_QDateEdit_connectNotify_default, "QDateEdit_connectNotify_default")
	symaddr(&addr_QDateEdit_disconnectNotify_default, "QDateEdit_disconnectNotify_default")
	reg(&pg_QDateEdit_connect_userDateChanged, "QDateEdit_connect_userDateChanged")
	reg(&pg_QDateEdit_connect_dateTimeChanged, "QDateEdit_connect_dateTimeChanged")
	reg(&pg_QDateEdit_connect_timeChanged, "QDateEdit_connect_timeChanged")
	reg(&pg_QDateEdit_connect_dateChanged, "QDateEdit_connect_dateChanged")
	reg(&pg_QDateEdit_connect_editingFinished, "QDateEdit_connect_editingFinished")
	reg(&pg_QDateEdit_connect_returnPressed, "QDateEdit_connect_returnPressed")
	reg(&pg_QDateEdit_connect_windowTitleChanged, "QDateEdit_connect_windowTitleChanged")
	reg(&pg_QDateEdit_connect_windowIconChanged, "QDateEdit_connect_windowIconChanged")
	reg(&pg_QDateEdit_connect_windowIconTextChanged, "QDateEdit_connect_windowIconTextChanged")
	reg(&pg_QDateEdit_connect_customContextMenuRequested, "QDateEdit_connect_customContextMenuRequested")
	reg(&pg_QDateEdit_connect_destroyed, "QDateEdit_connect_destroyed")
	reg(&pg_QDateEdit_connect_destroyed2, "QDateEdit_connect_destroyed2")
	reg(&pg_QDateEdit_connect_objectNameChanged, "QDateEdit_connect_objectNameChanged")
	reg(&pg_QDateEdit__disconnect, "QDateEdit__disconnect")
	reg(&pg_QDateEdit__set_override, "QDateEdit__set_override")
	reg(&pg_QDateEdit_delete, "QDateEdit_delete")
}

type QDateEdit struct {
	h uintptr
	*QDateTimeEdit
}

func bindQDateEdit(h uintptr) *QDateEdit {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QDateEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDateEdit{h: h, QDateTimeEdit: &QDateTimeEdit{h: __c[0], QAbstractSpinBox: &QAbstractSpinBox{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQDateEdit(h uintptr) *QDateEdit { return bindQDateEdit(h) }

func (this *QDateEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDateEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDateEdit(h unsafe.Pointer) *QDateEdit { return bindQDateEdit(uintptr(h)) }
func UnsafeNewQDateEditOwned(h unsafe.Pointer) *QDateEdit { return newownedQDateEdit(uintptr(h)) }

func (this *QDateEdit) UnsafeRelease() {
	this.h = 0
	if this.QDateTimeEdit != nil {
		this.QDateTimeEdit.UnsafeRelease()
	}
}

func NewQDateEdit(parent *QWidget) *QDateEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDateEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateEdit_virtIdx, QDateEdit_enable)
	return __o
}

func NewQDateEdit2() *QDateEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateEdit_new2, uintptr(__goid))
	__o := newownedQDateEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateEdit_virtIdx, QDateEdit_enable)
	return __o
}

func NewQDateEdit3(date *QDate) *QDateEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateEdit_new3, uintptr(__goid), uintptr(date.UnsafePointer()))
	__o := newownedQDateEdit(__r)
	runtime.KeepAlive(date)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateEdit_virtIdx, QDateEdit_enable)
	return __o
}

func NewQDateEdit4(date *QDate, parent *QWidget) *QDateEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateEdit_new4, uintptr(__goid), uintptr(date.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQDateEdit(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateEdit_virtIdx, QDateEdit_enable)
	return __o
}

func QDateEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDateEdit) UserDateChanged(date *QDate) {
	purego.SyscallN(addr_QDateEdit_userDateChanged, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QDateEdit) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QDateEdit) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QDateEdit_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func QDateEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDateEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDateEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QDateEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QDateEdit_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDateEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDateEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QDateEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateEdit) Create() {
	purego.SyscallN(addr_QDateEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateEdit) Destroy() {
	purego.SyscallN(addr_QDateEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDateEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDateEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDateEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDateEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDateEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDateEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDateEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QDateEdit_devicePixelRatioFScale()
	return __r
}

func QDateEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDateEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDateEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDateEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDateEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDateEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDateEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDateEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDateEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDateEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDateEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDateEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDateEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDateEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDateEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDateEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDateEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDateEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDateEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDateEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDateEdit) ClearDefault() {
	purego.SyscallN(addr_QDateEdit_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateEdit) StepByDefault(steps int) {
	purego.SyscallN(addr_QDateEdit_stepBy_default, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDateEdit) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDateEdit_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDateEdit_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDateEdit_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateEdit) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QDateEdit) FixupDefault(input *QString) {
	purego.SyscallN(addr_QDateEdit_fixup_default, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QDateEdit) DateTimeFromTextDefault(text string) *QDateTime {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateEdit_dateTimeFromText_default, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQDateTime(__r)
}

func (this *QDateEdit) TextFromDateTimeDefault(dt *QDateTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateEdit_textFromDateTime_default, this.h, uintptr(dt.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dt)
	return pgStr(__r, __n)
}

func (this *QDateEdit) StepEnabledDefault() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_stepEnabled_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QDateEdit) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateEdit_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDateEdit_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) InitStyleOptionDefault(option *QStyleOptionSpinBox) {
	purego.SyscallN(addr_QDateEdit_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QDateEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDateEdit) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDateEdit) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QDateEdit_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDateEdit_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDateEdit_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDateEdit_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateEdit_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDateEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDateEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateEdit_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateEdit_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDateEdit_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDateEdit_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDateEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDateEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDateEdit) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateEdit_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDateEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDateEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDateEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDateEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDateEdit_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDateEdit_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDateEdit_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDateEdit_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDateEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDateEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDateEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDateEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDateEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDateEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDateEdit) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateEdit_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDateEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDateEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDateEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDateEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDateEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDateEdit) OnUserDateChanged(fn func(*QDate)) int64 {
	id := registerCallback(func(date uintptr) { fn(UnsafeNewQDate(unsafe.Pointer(date))) })
	pg_QDateEdit_connect_userDateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnDateTimeChanged(fn func(*QDateTime)) int64 {
	id := registerCallback(func(dateTime uintptr) { fn(UnsafeNewQDateTime(unsafe.Pointer(dateTime))) })
	pg_QDateEdit_connect_dateTimeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnTimeChanged(fn func(*QTime)) int64 {
	id := registerCallback(func(time uintptr) { fn(UnsafeNewQTime(unsafe.Pointer(time))) })
	pg_QDateEdit_connect_timeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnDateChanged(fn func(*QDate)) int64 {
	id := registerCallback(func(date uintptr) { fn(UnsafeNewQDate(unsafe.Pointer(date))) })
	pg_QDateEdit_connect_dateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateEdit_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateEdit_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDateEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDateEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDateEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDateEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDateEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDateEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateEdit) Disconnect(handle int64) {
	pg_QDateEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QDateEdit_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case -778567940572639070:
		return 2
	case 7085718643589954634:
		return 3
	case 1237236456312824991:
		return 4
	case 1023791977679793185:
		return 5
	case 1919988340934511316:
		return 6
	case 975043767033552552:
		return 7
	case 5898328517505784191:
		return 8
	case 6784179835071232555:
		return 9
	case -4855206404746966715:
		return 10
	case -50441425836614729:
		return 11
	case -3778148144965928937:
		return 12
	case 5257289574516412244:
		return 13
	case -332090898078543319:
		return 14
	case 5290442163167458465:
		return 15
	case -488693769771308531:
		return 16
	case 7049264599837532491:
		return 17
	case 7345853097340709638:
		return 18
	case -4302917285655014661:
		return 19
	case 772598475950050299:
		return 20
	case -385596010576048993:
		return 21
	case -6698043222227006015:
		return 22
	case 4959376063826165943:
		return 23
	case -1989937050019464783:
		return 24
	case -199718707876719129:
		return 25
	case -4415920049297795213:
		return 26
	case -4255494004099806581:
		return 27
	case -5409927581812374498:
		return 28
	case -6825369174751916:
		return 29
	case -8659679893868524044:
		return 30
	case 2110182719375245575:
		return 31
	case 4311534060495737339:
		return 32
	case -8266762397995423915:
		return 33
	case 6398007404756346381:
		return 34
	case -4029073905697023887:
		return 35
	case 6316270346668237895:
		return 36
	case -551070009623711640:
		return 37
	case -9161949150489525396:
		return 38
	case -212068866860701141:
		return 39
	case 1697444140088555581:
		return 40
	case 206122715919917589:
		return 41
	case -1781545137352745978:
		return 42
	case -3139904158373725006:
		return 43
	case 9107058682754591716:
		return 44
	case -2039166586246443706:
		return 45
	case 427629649267311777:
		return 46
	case 6477500296980050210:
		return 47
	case 4127240633310655048:
		return 48
	case 5792771328741774991:
		return 49
	case 5438819661942843222:
		return 50
	case 5217623514630858297:
		return 51
	case -1185315367526857833:
		return 52
	case 3769868229123868474:
		return 53
	case 7379029982170503008:
		return 54
	case -568369772982215962:
		return 55
	}
	return -1
}

func QDateEdit_enable(h uintptr, idx int32) { pg_QDateEdit__set_override(h, idx) }

func (this *QDateEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDateEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDateEdit) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QDateEdit) OnStepBy(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 7085718643589954634, func(steps int32) { fn(int(steps)) }) // slot "stepBy"
}

func (this *QDateEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDateEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QDateEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDateEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDateEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDateEdit) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QDateEdit) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(input uintptr) { fn(UnsafeNewQString(unsafe.Pointer(input))) }) // slot "fixup"
}

func (this *QDateEdit) OnDateTimeFromText(fn func(string) *QDateTime) {
	registerVirt(this.UnsafeGoID(), -50441425836614729, func(text uintptr, text_len int64) uintptr { return uintptr(fn(pgStrNoFree(text, text_len)).UnsafePointer()) }) // slot "dateTimeFromText"
}

func (this *QDateEdit) OnTextFromDateTime(fn func(*QDateTime) string) {
	registerVirt(this.UnsafeGoID(), -3778148144965928937, func(dt uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQDateTime(unsafe.Pointer(dt)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textFromDateTime"
}

func (this *QDateEdit) OnStepEnabled(fn func() QAbstractSpinBox__StepEnabledFlag) {
	registerVirt(this.UnsafeGoID(), 5257289574516412244, func() int32 { return int32(fn()) }) // slot "stepEnabled"
}

func (this *QDateEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDateEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDateEdit) OnInitStyleOption(fn func(*QStyleOptionSpinBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSpinBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QDateEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDateEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDateEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QDateEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDateEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDateEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDateEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QDateEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDateEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDateEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDateEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDateEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDateEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDateEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDateEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDateEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDateEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDateEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDateEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDateEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDateEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDateEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDateEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDateEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDateEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDateEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDateEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDateEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDateEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDateEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDateEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDateEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDateEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDateEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDateEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDateEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDateEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDateEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDateEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDateEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDateEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
