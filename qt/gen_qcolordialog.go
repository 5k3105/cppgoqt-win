package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColorDialog_chain uintptr
	addr_QColorDialog_new uintptr
	addr_QColorDialog_new2 uintptr
	addr_QColorDialog_new3 uintptr
	addr_QColorDialog_new4 uintptr
	addr_QColorDialog_tr uintptr
	addr_QColorDialog_setCurrentColor uintptr
	addr_QColorDialog_currentColor uintptr
	addr_QColorDialog_selectedColor uintptr
	addr_QColorDialog_setOption uintptr
	addr_QColorDialog_testOption uintptr
	addr_QColorDialog_setOptions uintptr
	addr_QColorDialog_options uintptr
	addr_QColorDialog_open uintptr
	addr_QColorDialog_getColor uintptr
	addr_QColorDialog_customCount uintptr
	addr_QColorDialog_customColor uintptr
	addr_QColorDialog_setCustomColor uintptr
	addr_QColorDialog_standardColor uintptr
	addr_QColorDialog_setStandardColor uintptr
	addr_QColorDialog_currentColorChanged uintptr
	addr_QColorDialog_colorSelected uintptr
	addr_QColorDialog_adjustPosition uintptr
	addr_QColorDialog_setTabOrder uintptr
	addr_QColorDialog_mouseGrabber uintptr
	addr_QColorDialog_keyboardGrabber uintptr
	addr_QColorDialog_find uintptr
	addr_QColorDialog_createWindowContainer uintptr
	addr_QColorDialog_updateMicroFocus uintptr
	addr_QColorDialog_create uintptr
	addr_QColorDialog_destroy uintptr
	addr_QColorDialog_focusNextChild uintptr
	addr_QColorDialog_focusPreviousChild uintptr
	addr_QColorDialog_connect uintptr
	addr_QColorDialog_connect2 uintptr
	addr_QColorDialog_sender uintptr
	addr_QColorDialog_senderSignalIndex uintptr
	addr_QColorDialog_receivers uintptr
	addr_QColorDialog_isSignalConnected uintptr
	pg_QColorDialog_devicePixelRatioFScale func() float64
	pg_QColorDialog_encodeMetricF func(int32, float64) int32
	pg_QColorDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QColorDialog_tr2 uintptr
	addr_QColorDialog_tr3 uintptr
	addr_QColorDialog_setOption2 uintptr
	addr_QColorDialog_getColor2 uintptr
	addr_QColorDialog_getColor3 uintptr
	addr_QColorDialog_getColor4 uintptr
	addr_QColorDialog_getColor5 uintptr
	addr_QColorDialog_createWindowContainer2 uintptr
	addr_QColorDialog_createWindowContainer3 uintptr
	addr_QColorDialog_updateMicroFocus2 uintptr
	addr_QColorDialog_create2 uintptr
	addr_QColorDialog_create3 uintptr
	addr_QColorDialog_create4 uintptr
	addr_QColorDialog_destroy2 uintptr
	addr_QColorDialog_destroy3 uintptr
	addr_QColorDialog_connect3 uintptr
	addr_QColorDialog_connect4 uintptr
	addr_QColorDialog_metaObject_default uintptr
	addr_QColorDialog_setVisible_default uintptr
	addr_QColorDialog_changeEvent_default uintptr
	addr_QColorDialog_done_default uintptr
	addr_QColorDialog_sizeHint_default uintptr
	addr_QColorDialog_minimumSizeHint_default uintptr
	addr_QColorDialog_exec_default uintptr
	addr_QColorDialog_accept_default uintptr
	addr_QColorDialog_reject_default uintptr
	addr_QColorDialog_keyPressEvent_default uintptr
	addr_QColorDialog_closeEvent_default uintptr
	addr_QColorDialog_showEvent_default uintptr
	addr_QColorDialog_resizeEvent_default uintptr
	addr_QColorDialog_contextMenuEvent_default uintptr
	addr_QColorDialog_eventFilter_default uintptr
	addr_QColorDialog_devType_default uintptr
	addr_QColorDialog_heightForWidth_default uintptr
	addr_QColorDialog_hasHeightForWidth_default uintptr
	addr_QColorDialog_paintEngine_default uintptr
	addr_QColorDialog_event_default uintptr
	addr_QColorDialog_mousePressEvent_default uintptr
	addr_QColorDialog_mouseReleaseEvent_default uintptr
	addr_QColorDialog_mouseDoubleClickEvent_default uintptr
	addr_QColorDialog_mouseMoveEvent_default uintptr
	addr_QColorDialog_wheelEvent_default uintptr
	addr_QColorDialog_keyReleaseEvent_default uintptr
	addr_QColorDialog_focusInEvent_default uintptr
	addr_QColorDialog_focusOutEvent_default uintptr
	addr_QColorDialog_enterEvent_default uintptr
	addr_QColorDialog_leaveEvent_default uintptr
	addr_QColorDialog_paintEvent_default uintptr
	addr_QColorDialog_moveEvent_default uintptr
	addr_QColorDialog_tabletEvent_default uintptr
	addr_QColorDialog_actionEvent_default uintptr
	addr_QColorDialog_dragEnterEvent_default uintptr
	addr_QColorDialog_dragMoveEvent_default uintptr
	addr_QColorDialog_dragLeaveEvent_default uintptr
	addr_QColorDialog_dropEvent_default uintptr
	addr_QColorDialog_hideEvent_default uintptr
	addr_QColorDialog_metric_default uintptr
	addr_QColorDialog_initPainter_default uintptr
	addr_QColorDialog_redirected_default uintptr
	addr_QColorDialog_sharedPainter_default uintptr
	addr_QColorDialog_inputMethodEvent_default uintptr
	addr_QColorDialog_inputMethodQuery_default uintptr
	addr_QColorDialog_focusNextPrevChild_default uintptr
	addr_QColorDialog_timerEvent_default uintptr
	addr_QColorDialog_childEvent_default uintptr
	addr_QColorDialog_customEvent_default uintptr
	addr_QColorDialog_connectNotify_default uintptr
	addr_QColorDialog_disconnectNotify_default uintptr
	pg_QColorDialog_connect_currentColorChanged func(uintptr, int64)
	pg_QColorDialog_connect_colorSelected func(uintptr, int64)
	pg_QColorDialog_connect_finished func(uintptr, int64)
	pg_QColorDialog_connect_accepted func(uintptr, int64)
	pg_QColorDialog_connect_rejected func(uintptr, int64)
	pg_QColorDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QColorDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QColorDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QColorDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QColorDialog_connect_destroyed func(uintptr, int64)
	pg_QColorDialog_connect_destroyed2 func(uintptr, int64)
	pg_QColorDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QColorDialog__disconnect func(int64)
	pg_QColorDialog__set_override func(uintptr, int32)
	pg_QColorDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColorDialog_chain, "QColorDialog_chain")
	symaddr(&addr_QColorDialog_new, "QColorDialog_new")
	symaddr(&addr_QColorDialog_new2, "QColorDialog_new2")
	symaddr(&addr_QColorDialog_new3, "QColorDialog_new3")
	symaddr(&addr_QColorDialog_new4, "QColorDialog_new4")
	symaddr(&addr_QColorDialog_tr, "QColorDialog_tr")
	symaddr(&addr_QColorDialog_setCurrentColor, "QColorDialog_setCurrentColor")
	symaddr(&addr_QColorDialog_currentColor, "QColorDialog_currentColor")
	symaddr(&addr_QColorDialog_selectedColor, "QColorDialog_selectedColor")
	symaddr(&addr_QColorDialog_setOption, "QColorDialog_setOption")
	symaddr(&addr_QColorDialog_testOption, "QColorDialog_testOption")
	symaddr(&addr_QColorDialog_setOptions, "QColorDialog_setOptions")
	symaddr(&addr_QColorDialog_options, "QColorDialog_options")
	symaddr(&addr_QColorDialog_open, "QColorDialog_open")
	symaddr(&addr_QColorDialog_getColor, "QColorDialog_getColor")
	symaddr(&addr_QColorDialog_customCount, "QColorDialog_customCount")
	symaddr(&addr_QColorDialog_customColor, "QColorDialog_customColor")
	symaddr(&addr_QColorDialog_setCustomColor, "QColorDialog_setCustomColor")
	symaddr(&addr_QColorDialog_standardColor, "QColorDialog_standardColor")
	symaddr(&addr_QColorDialog_setStandardColor, "QColorDialog_setStandardColor")
	symaddr(&addr_QColorDialog_currentColorChanged, "QColorDialog_currentColorChanged")
	symaddr(&addr_QColorDialog_colorSelected, "QColorDialog_colorSelected")
	symaddr(&addr_QColorDialog_adjustPosition, "QColorDialog_adjustPosition")
	symaddr(&addr_QColorDialog_setTabOrder, "QColorDialog_setTabOrder")
	symaddr(&addr_QColorDialog_mouseGrabber, "QColorDialog_mouseGrabber")
	symaddr(&addr_QColorDialog_keyboardGrabber, "QColorDialog_keyboardGrabber")
	symaddr(&addr_QColorDialog_find, "QColorDialog_find")
	symaddr(&addr_QColorDialog_createWindowContainer, "QColorDialog_createWindowContainer")
	symaddr(&addr_QColorDialog_updateMicroFocus, "QColorDialog_updateMicroFocus")
	symaddr(&addr_QColorDialog_create, "QColorDialog_create")
	symaddr(&addr_QColorDialog_destroy, "QColorDialog_destroy")
	symaddr(&addr_QColorDialog_focusNextChild, "QColorDialog_focusNextChild")
	symaddr(&addr_QColorDialog_focusPreviousChild, "QColorDialog_focusPreviousChild")
	symaddr(&addr_QColorDialog_connect, "QColorDialog_connect")
	symaddr(&addr_QColorDialog_connect2, "QColorDialog_connect2")
	symaddr(&addr_QColorDialog_sender, "QColorDialog_sender")
	symaddr(&addr_QColorDialog_senderSignalIndex, "QColorDialog_senderSignalIndex")
	symaddr(&addr_QColorDialog_receivers, "QColorDialog_receivers")
	symaddr(&addr_QColorDialog_isSignalConnected, "QColorDialog_isSignalConnected")
	reg(&pg_QColorDialog_devicePixelRatioFScale, "QColorDialog_devicePixelRatioFScale")
	reg(&pg_QColorDialog_encodeMetricF, "QColorDialog_encodeMetricF")
	reg(&pg_QColorDialog_getDecodedMetricF, "QColorDialog_getDecodedMetricF")
	symaddr(&addr_QColorDialog_tr2, "QColorDialog_tr2")
	symaddr(&addr_QColorDialog_tr3, "QColorDialog_tr3")
	symaddr(&addr_QColorDialog_setOption2, "QColorDialog_setOption2")
	symaddr(&addr_QColorDialog_getColor2, "QColorDialog_getColor2")
	symaddr(&addr_QColorDialog_getColor3, "QColorDialog_getColor3")
	symaddr(&addr_QColorDialog_getColor4, "QColorDialog_getColor4")
	symaddr(&addr_QColorDialog_getColor5, "QColorDialog_getColor5")
	symaddr(&addr_QColorDialog_createWindowContainer2, "QColorDialog_createWindowContainer2")
	symaddr(&addr_QColorDialog_createWindowContainer3, "QColorDialog_createWindowContainer3")
	symaddr(&addr_QColorDialog_updateMicroFocus2, "QColorDialog_updateMicroFocus2")
	symaddr(&addr_QColorDialog_create2, "QColorDialog_create2")
	symaddr(&addr_QColorDialog_create3, "QColorDialog_create3")
	symaddr(&addr_QColorDialog_create4, "QColorDialog_create4")
	symaddr(&addr_QColorDialog_destroy2, "QColorDialog_destroy2")
	symaddr(&addr_QColorDialog_destroy3, "QColorDialog_destroy3")
	symaddr(&addr_QColorDialog_connect3, "QColorDialog_connect3")
	symaddr(&addr_QColorDialog_connect4, "QColorDialog_connect4")
	symaddr(&addr_QColorDialog_metaObject_default, "QColorDialog_metaObject_default")
	symaddr(&addr_QColorDialog_setVisible_default, "QColorDialog_setVisible_default")
	symaddr(&addr_QColorDialog_changeEvent_default, "QColorDialog_changeEvent_default")
	symaddr(&addr_QColorDialog_done_default, "QColorDialog_done_default")
	symaddr(&addr_QColorDialog_sizeHint_default, "QColorDialog_sizeHint_default")
	symaddr(&addr_QColorDialog_minimumSizeHint_default, "QColorDialog_minimumSizeHint_default")
	symaddr(&addr_QColorDialog_exec_default, "QColorDialog_exec_default")
	symaddr(&addr_QColorDialog_accept_default, "QColorDialog_accept_default")
	symaddr(&addr_QColorDialog_reject_default, "QColorDialog_reject_default")
	symaddr(&addr_QColorDialog_keyPressEvent_default, "QColorDialog_keyPressEvent_default")
	symaddr(&addr_QColorDialog_closeEvent_default, "QColorDialog_closeEvent_default")
	symaddr(&addr_QColorDialog_showEvent_default, "QColorDialog_showEvent_default")
	symaddr(&addr_QColorDialog_resizeEvent_default, "QColorDialog_resizeEvent_default")
	symaddr(&addr_QColorDialog_contextMenuEvent_default, "QColorDialog_contextMenuEvent_default")
	symaddr(&addr_QColorDialog_eventFilter_default, "QColorDialog_eventFilter_default")
	symaddr(&addr_QColorDialog_devType_default, "QColorDialog_devType_default")
	symaddr(&addr_QColorDialog_heightForWidth_default, "QColorDialog_heightForWidth_default")
	symaddr(&addr_QColorDialog_hasHeightForWidth_default, "QColorDialog_hasHeightForWidth_default")
	symaddr(&addr_QColorDialog_paintEngine_default, "QColorDialog_paintEngine_default")
	symaddr(&addr_QColorDialog_event_default, "QColorDialog_event_default")
	symaddr(&addr_QColorDialog_mousePressEvent_default, "QColorDialog_mousePressEvent_default")
	symaddr(&addr_QColorDialog_mouseReleaseEvent_default, "QColorDialog_mouseReleaseEvent_default")
	symaddr(&addr_QColorDialog_mouseDoubleClickEvent_default, "QColorDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QColorDialog_mouseMoveEvent_default, "QColorDialog_mouseMoveEvent_default")
	symaddr(&addr_QColorDialog_wheelEvent_default, "QColorDialog_wheelEvent_default")
	symaddr(&addr_QColorDialog_keyReleaseEvent_default, "QColorDialog_keyReleaseEvent_default")
	symaddr(&addr_QColorDialog_focusInEvent_default, "QColorDialog_focusInEvent_default")
	symaddr(&addr_QColorDialog_focusOutEvent_default, "QColorDialog_focusOutEvent_default")
	symaddr(&addr_QColorDialog_enterEvent_default, "QColorDialog_enterEvent_default")
	symaddr(&addr_QColorDialog_leaveEvent_default, "QColorDialog_leaveEvent_default")
	symaddr(&addr_QColorDialog_paintEvent_default, "QColorDialog_paintEvent_default")
	symaddr(&addr_QColorDialog_moveEvent_default, "QColorDialog_moveEvent_default")
	symaddr(&addr_QColorDialog_tabletEvent_default, "QColorDialog_tabletEvent_default")
	symaddr(&addr_QColorDialog_actionEvent_default, "QColorDialog_actionEvent_default")
	symaddr(&addr_QColorDialog_dragEnterEvent_default, "QColorDialog_dragEnterEvent_default")
	symaddr(&addr_QColorDialog_dragMoveEvent_default, "QColorDialog_dragMoveEvent_default")
	symaddr(&addr_QColorDialog_dragLeaveEvent_default, "QColorDialog_dragLeaveEvent_default")
	symaddr(&addr_QColorDialog_dropEvent_default, "QColorDialog_dropEvent_default")
	symaddr(&addr_QColorDialog_hideEvent_default, "QColorDialog_hideEvent_default")
	symaddr(&addr_QColorDialog_metric_default, "QColorDialog_metric_default")
	symaddr(&addr_QColorDialog_initPainter_default, "QColorDialog_initPainter_default")
	symaddr(&addr_QColorDialog_redirected_default, "QColorDialog_redirected_default")
	symaddr(&addr_QColorDialog_sharedPainter_default, "QColorDialog_sharedPainter_default")
	symaddr(&addr_QColorDialog_inputMethodEvent_default, "QColorDialog_inputMethodEvent_default")
	symaddr(&addr_QColorDialog_inputMethodQuery_default, "QColorDialog_inputMethodQuery_default")
	symaddr(&addr_QColorDialog_focusNextPrevChild_default, "QColorDialog_focusNextPrevChild_default")
	symaddr(&addr_QColorDialog_timerEvent_default, "QColorDialog_timerEvent_default")
	symaddr(&addr_QColorDialog_childEvent_default, "QColorDialog_childEvent_default")
	symaddr(&addr_QColorDialog_customEvent_default, "QColorDialog_customEvent_default")
	symaddr(&addr_QColorDialog_connectNotify_default, "QColorDialog_connectNotify_default")
	symaddr(&addr_QColorDialog_disconnectNotify_default, "QColorDialog_disconnectNotify_default")
	reg(&pg_QColorDialog_connect_currentColorChanged, "QColorDialog_connect_currentColorChanged")
	reg(&pg_QColorDialog_connect_colorSelected, "QColorDialog_connect_colorSelected")
	reg(&pg_QColorDialog_connect_finished, "QColorDialog_connect_finished")
	reg(&pg_QColorDialog_connect_accepted, "QColorDialog_connect_accepted")
	reg(&pg_QColorDialog_connect_rejected, "QColorDialog_connect_rejected")
	reg(&pg_QColorDialog_connect_windowTitleChanged, "QColorDialog_connect_windowTitleChanged")
	reg(&pg_QColorDialog_connect_windowIconChanged, "QColorDialog_connect_windowIconChanged")
	reg(&pg_QColorDialog_connect_windowIconTextChanged, "QColorDialog_connect_windowIconTextChanged")
	reg(&pg_QColorDialog_connect_customContextMenuRequested, "QColorDialog_connect_customContextMenuRequested")
	reg(&pg_QColorDialog_connect_destroyed, "QColorDialog_connect_destroyed")
	reg(&pg_QColorDialog_connect_destroyed2, "QColorDialog_connect_destroyed2")
	reg(&pg_QColorDialog_connect_objectNameChanged, "QColorDialog_connect_objectNameChanged")
	reg(&pg_QColorDialog__disconnect, "QColorDialog__disconnect")
	reg(&pg_QColorDialog__set_override, "QColorDialog__set_override")
	reg(&pg_QColorDialog_delete, "QColorDialog_delete")
}

type QColorDialog struct {
	h uintptr
	*QDialog
}

func bindQColorDialog(h uintptr) *QColorDialog {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QColorDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QColorDialog{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQColorDialog(h uintptr) *QColorDialog { return bindQColorDialog(h) }

func (this *QColorDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColorDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColorDialog(h unsafe.Pointer) *QColorDialog { return bindQColorDialog(uintptr(h)) }
func UnsafeNewQColorDialogOwned(h unsafe.Pointer) *QColorDialog { return newownedQColorDialog(uintptr(h)) }

func (this *QColorDialog) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQColorDialog(parent *QWidget) *QColorDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColorDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQColorDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColorDialog_virtIdx, QColorDialog_enable)
	return __o
}

func NewQColorDialog2() *QColorDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColorDialog_new2, uintptr(__goid))
	__o := newownedQColorDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColorDialog_virtIdx, QColorDialog_enable)
	return __o
}

func NewQColorDialog3(initial *QColor) *QColorDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColorDialog_new3, uintptr(__goid), uintptr(initial.UnsafePointer()))
	__o := newownedQColorDialog(__r)
	runtime.KeepAlive(initial)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColorDialog_virtIdx, QColorDialog_enable)
	return __o
}

func NewQColorDialog4(initial *QColor, parent *QWidget) *QColorDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QColorDialog_new4, uintptr(__goid), uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQColorDialog(__r)
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QColorDialog_virtIdx, QColorDialog_enable)
	return __o
}

func QColorDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColorDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QColorDialog) SetCurrentColor(color *QColor) {
	purego.SyscallN(addr_QColorDialog_setCurrentColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QColorDialog) CurrentColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_currentColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColorDialog) SelectedColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_selectedColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColorDialog) SetOption(option QColorDialog__ColorDialogOption) {
	purego.SyscallN(addr_QColorDialog_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) TestOption(option QColorDialog__ColorDialogOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorDialog) SetOptions(options QColorDialog__ColorDialogOption) {
	purego.SyscallN(addr_QColorDialog_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Options() QColorDialog__ColorDialogOption {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColorDialog__ColorDialogOption(__r)
}

func (this *QColorDialog) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QColorDialog_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func QColorDialog_GetColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_getColor)
	return newownedQColor(__r)
}

func QColorDialog_CustomCount() int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_customCount)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QColorDialog_CustomColor(index int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_customColor, uintptr(index))
	return newownedQColor(__r)
}

func QColorDialog_SetCustomColor(index int, color *QColor) {
	purego.SyscallN(addr_QColorDialog_setCustomColor, uintptr(index), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(color)
}

func QColorDialog_StandardColor(index int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_standardColor, uintptr(index))
	return newownedQColor(__r)
}

func QColorDialog_SetStandardColor(index int, color *QColor) {
	purego.SyscallN(addr_QColorDialog_setStandardColor, uintptr(index), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(color)
}

func (this *QColorDialog) CurrentColorChanged(color *QColor) {
	purego.SyscallN(addr_QColorDialog_currentColorChanged, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QColorDialog) ColorSelected(color *QColor) {
	purego.SyscallN(addr_QColorDialog_colorSelected, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QColorDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QColorDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QColorDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QColorDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QColorDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QColorDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QColorDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QColorDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QColorDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QColorDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Create() {
	purego.SyscallN(addr_QColorDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Destroy() {
	purego.SyscallN(addr_QColorDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QColorDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QColorDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QColorDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QColorDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QColorDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColorDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QColorDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QColorDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QColorDialog_devicePixelRatioFScale()
	return __r
}

func QColorDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QColorDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QColorDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QColorDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QColorDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColorDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QColorDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColorDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QColorDialog) SetOption2(option QColorDialog__ColorDialogOption, on bool) {
	purego.SyscallN(addr_QColorDialog_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QColorDialog_GetColorWithInitial(initial *QColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_getColor2, uintptr(initial.UnsafePointer()))
	runtime.KeepAlive(initial)
	return newownedQColor(__r)
}

func QColorDialog_GetColor2(initial *QColor, parent *QWidget) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_getColor3, uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	return newownedQColor(__r)
}

func QColorDialog_GetColor3(initial *QColor, parent *QWidget, title string) *QColor {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QColorDialog_getColor4, uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	return newownedQColor(__r)
}

func QColorDialog_GetColor4(initial *QColor, parent *QWidget, title string, options QColorDialog__ColorDialogOption) *QColor {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QColorDialog_getColor5, uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), uintptr(options))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	return newownedQColor(__r)
}

func QColorDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QColorDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QColorDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QColorDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QColorDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QColorDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QColorDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QColorDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QColorDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QColorDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QColorDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QColorDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QColorDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QColorDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QColorDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QColorDialog_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) DoneDefault(result int) {
	purego.SyscallN(addr_QColorDialog_done_default, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QColorDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QColorDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QColorDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColorDialog) AcceptDefault() {
	purego.SyscallN(addr_QColorDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorDialog) RejectDefault() {
	purego.SyscallN(addr_QColorDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QColorDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QColorDialog_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QColorDialog_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QColorDialog_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QColorDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QColorDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColorDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColorDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QColorDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QColorDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColorDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColorDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColorDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QColorDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QColorDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QColorDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QColorDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QColorDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QColorDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QColorDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QColorDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QColorDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QColorDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QColorDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QColorDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QColorDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QColorDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QColorDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QColorDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColorDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QColorDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QColorDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QColorDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QColorDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QColorDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QColorDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QColorDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QColorDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QColorDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QColorDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QColorDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QColorDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QColorDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QColorDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QColorDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QColorDialog) OnCurrentColorChanged(fn func(*QColor)) int64 {
	id := registerCallback(func(color uintptr) { fn(UnsafeNewQColor(unsafe.Pointer(color))) })
	pg_QColorDialog_connect_currentColorChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnColorSelected(fn func(*QColor)) int64 {
	id := registerCallback(func(color uintptr) { fn(UnsafeNewQColor(unsafe.Pointer(color))) })
	pg_QColorDialog_connect_colorSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QColorDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QColorDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QColorDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QColorDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QColorDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QColorDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QColorDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QColorDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QColorDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QColorDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QColorDialog) Disconnect(handle int64) {
	pg_QColorDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QColorDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 2110182719375245575:
		return 1
	case 4959376063826165943:
		return 2
	case -2570997490807640431:
		return 3
	case 729690261457684101:
		return 4
	case 7049264599837532491:
		return 5
	case -1745322561779819684:
		return 6
	case 9106131928185084681:
		return 7
	case 2875220963396170150:
		return 8
	case 1023791977679793185:
		return 9
	case -1989937050019464783:
		return 10
	case -6825369174751916:
		return 11
	case -4302917285655014661:
		return 12
	case -6698043222227006015:
		return 13
	case 5217623514630858297:
		return 14
	case -8659679893868524044:
		return 15
	case 4311534060495737339:
		return 16
	case -8266762397995423915:
		return 17
	case 6398007404756346381:
		return 18
	case 1237236456312824991:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case -4255494004099806581:
		return 23
	case 1919988340934511316:
		return 24
	case 772598475950050299:
		return 25
	case 975043767033552552:
		return 26
	case -385596010576048993:
		return 27
	case 6316270346668237895:
		return 28
	case -551070009623711640:
		return 29
	case 5290442163167458465:
		return 30
	case -9161949150489525396:
		return 31
	case -212068866860701141:
		return 32
	case 1697444140088555581:
		return 33
	case 206122715919917589:
		return 34
	case -1781545137352745978:
		return 35
	case -3139904158373725006:
		return 36
	case 9107058682754591716:
		return 37
	case -199718707876719129:
		return 38
	case -2039166586246443706:
		return 39
	case 427629649267311777:
		return 40
	case 6477500296980050210:
		return 41
	case 4127240633310655048:
		return 42
	case 5792771328741774991:
		return 43
	case 5438819661942843222:
		return 44
	case 7345853097340709638:
		return 45
	case 5898328517505784191:
		return 46
	case -5409927581812374498:
		return 47
	case -1185315367526857833:
		return 48
	case 3769868229123868474:
		return 49
	case 7379029982170503008:
		return 50
	case -568369772982215962:
		return 51
	}
	return -1
}

func QColorDialog_enable(h uintptr, idx int32) { pg_QColorDialog__set_override(h, idx) }

func (this *QColorDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QColorDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QColorDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QColorDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(result int32) { fn(int(result)) }) // slot "done"
}

func (this *QColorDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QColorDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QColorDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QColorDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QColorDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QColorDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QColorDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QColorDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QColorDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QColorDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QColorDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QColorDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QColorDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QColorDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QColorDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QColorDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QColorDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QColorDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QColorDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QColorDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QColorDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QColorDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QColorDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QColorDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QColorDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QColorDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QColorDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QColorDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QColorDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QColorDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QColorDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QColorDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QColorDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QColorDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QColorDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QColorDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QColorDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QColorDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QColorDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QColorDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QColorDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QColorDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QColorDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QColorDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QColorDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QColorDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QColorDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QColorDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QColorDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColorDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
