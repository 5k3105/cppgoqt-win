package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontDialog_chain uintptr
	addr_QFontDialog_new uintptr
	addr_QFontDialog_new2 uintptr
	addr_QFontDialog_new3 uintptr
	addr_QFontDialog_new4 uintptr
	addr_QFontDialog_tr uintptr
	addr_QFontDialog_setCurrentFont uintptr
	addr_QFontDialog_currentFont uintptr
	addr_QFontDialog_selectedFont uintptr
	addr_QFontDialog_setOption uintptr
	addr_QFontDialog_testOption uintptr
	addr_QFontDialog_setOptions uintptr
	addr_QFontDialog_options uintptr
	addr_QFontDialog_open uintptr
	addr_QFontDialog_getFont uintptr
	addr_QFontDialog_getFont2 uintptr
	addr_QFontDialog_currentFontChanged uintptr
	addr_QFontDialog_fontSelected uintptr
	addr_QFontDialog_adjustPosition uintptr
	addr_QFontDialog_setTabOrder uintptr
	addr_QFontDialog_mouseGrabber uintptr
	addr_QFontDialog_keyboardGrabber uintptr
	addr_QFontDialog_find uintptr
	addr_QFontDialog_createWindowContainer uintptr
	addr_QFontDialog_updateMicroFocus uintptr
	addr_QFontDialog_create uintptr
	addr_QFontDialog_destroy uintptr
	addr_QFontDialog_focusNextChild uintptr
	addr_QFontDialog_focusPreviousChild uintptr
	addr_QFontDialog_connect uintptr
	addr_QFontDialog_connect2 uintptr
	addr_QFontDialog_sender uintptr
	addr_QFontDialog_senderSignalIndex uintptr
	addr_QFontDialog_receivers uintptr
	addr_QFontDialog_isSignalConnected uintptr
	pg_QFontDialog_devicePixelRatioFScale func() float64
	pg_QFontDialog_encodeMetricF func(int32, float64) int32
	pg_QFontDialog_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QFontDialog_tr2 uintptr
	addr_QFontDialog_tr3 uintptr
	addr_QFontDialog_setOption2 uintptr
	addr_QFontDialog_getFont3 uintptr
	addr_QFontDialog_getFont4 uintptr
	addr_QFontDialog_getFont5 uintptr
	addr_QFontDialog_getFont6 uintptr
	addr_QFontDialog_createWindowContainer2 uintptr
	addr_QFontDialog_createWindowContainer3 uintptr
	addr_QFontDialog_updateMicroFocus2 uintptr
	addr_QFontDialog_create2 uintptr
	addr_QFontDialog_create3 uintptr
	addr_QFontDialog_create4 uintptr
	addr_QFontDialog_destroy2 uintptr
	addr_QFontDialog_destroy3 uintptr
	addr_QFontDialog_connect3 uintptr
	addr_QFontDialog_connect4 uintptr
	addr_QFontDialog_metaObject_default uintptr
	addr_QFontDialog_setVisible_default uintptr
	addr_QFontDialog_changeEvent_default uintptr
	addr_QFontDialog_done_default uintptr
	addr_QFontDialog_eventFilter_default uintptr
	addr_QFontDialog_sizeHint_default uintptr
	addr_QFontDialog_minimumSizeHint_default uintptr
	addr_QFontDialog_exec_default uintptr
	addr_QFontDialog_accept_default uintptr
	addr_QFontDialog_reject_default uintptr
	addr_QFontDialog_keyPressEvent_default uintptr
	addr_QFontDialog_closeEvent_default uintptr
	addr_QFontDialog_showEvent_default uintptr
	addr_QFontDialog_resizeEvent_default uintptr
	addr_QFontDialog_contextMenuEvent_default uintptr
	addr_QFontDialog_devType_default uintptr
	addr_QFontDialog_heightForWidth_default uintptr
	addr_QFontDialog_hasHeightForWidth_default uintptr
	addr_QFontDialog_paintEngine_default uintptr
	addr_QFontDialog_event_default uintptr
	addr_QFontDialog_mousePressEvent_default uintptr
	addr_QFontDialog_mouseReleaseEvent_default uintptr
	addr_QFontDialog_mouseDoubleClickEvent_default uintptr
	addr_QFontDialog_mouseMoveEvent_default uintptr
	addr_QFontDialog_wheelEvent_default uintptr
	addr_QFontDialog_keyReleaseEvent_default uintptr
	addr_QFontDialog_focusInEvent_default uintptr
	addr_QFontDialog_focusOutEvent_default uintptr
	addr_QFontDialog_enterEvent_default uintptr
	addr_QFontDialog_leaveEvent_default uintptr
	addr_QFontDialog_paintEvent_default uintptr
	addr_QFontDialog_moveEvent_default uintptr
	addr_QFontDialog_tabletEvent_default uintptr
	addr_QFontDialog_actionEvent_default uintptr
	addr_QFontDialog_dragEnterEvent_default uintptr
	addr_QFontDialog_dragMoveEvent_default uintptr
	addr_QFontDialog_dragLeaveEvent_default uintptr
	addr_QFontDialog_dropEvent_default uintptr
	addr_QFontDialog_hideEvent_default uintptr
	addr_QFontDialog_metric_default uintptr
	addr_QFontDialog_initPainter_default uintptr
	addr_QFontDialog_redirected_default uintptr
	addr_QFontDialog_sharedPainter_default uintptr
	addr_QFontDialog_inputMethodEvent_default uintptr
	addr_QFontDialog_inputMethodQuery_default uintptr
	addr_QFontDialog_focusNextPrevChild_default uintptr
	addr_QFontDialog_timerEvent_default uintptr
	addr_QFontDialog_childEvent_default uintptr
	addr_QFontDialog_customEvent_default uintptr
	addr_QFontDialog_connectNotify_default uintptr
	addr_QFontDialog_disconnectNotify_default uintptr
	pg_QFontDialog_connect_currentFontChanged func(uintptr, int64)
	pg_QFontDialog_connect_fontSelected func(uintptr, int64)
	pg_QFontDialog_connect_finished func(uintptr, int64)
	pg_QFontDialog_connect_accepted func(uintptr, int64)
	pg_QFontDialog_connect_rejected func(uintptr, int64)
	pg_QFontDialog_connect_windowTitleChanged func(uintptr, int64)
	pg_QFontDialog_connect_windowIconChanged func(uintptr, int64)
	pg_QFontDialog_connect_windowIconTextChanged func(uintptr, int64)
	pg_QFontDialog_connect_customContextMenuRequested func(uintptr, int64)
	pg_QFontDialog_connect_destroyed func(uintptr, int64)
	pg_QFontDialog_connect_destroyed2 func(uintptr, int64)
	pg_QFontDialog_connect_objectNameChanged func(uintptr, int64)
	pg_QFontDialog__disconnect func(int64)
	pg_QFontDialog__set_override func(uintptr, int32)
	pg_QFontDialog_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontDialog_chain, "QFontDialog_chain")
	symaddr(&addr_QFontDialog_new, "QFontDialog_new")
	symaddr(&addr_QFontDialog_new2, "QFontDialog_new2")
	symaddr(&addr_QFontDialog_new3, "QFontDialog_new3")
	symaddr(&addr_QFontDialog_new4, "QFontDialog_new4")
	symaddr(&addr_QFontDialog_tr, "QFontDialog_tr")
	symaddr(&addr_QFontDialog_setCurrentFont, "QFontDialog_setCurrentFont")
	symaddr(&addr_QFontDialog_currentFont, "QFontDialog_currentFont")
	symaddr(&addr_QFontDialog_selectedFont, "QFontDialog_selectedFont")
	symaddr(&addr_QFontDialog_setOption, "QFontDialog_setOption")
	symaddr(&addr_QFontDialog_testOption, "QFontDialog_testOption")
	symaddr(&addr_QFontDialog_setOptions, "QFontDialog_setOptions")
	symaddr(&addr_QFontDialog_options, "QFontDialog_options")
	symaddr(&addr_QFontDialog_open, "QFontDialog_open")
	symaddr(&addr_QFontDialog_getFont, "QFontDialog_getFont")
	symaddr(&addr_QFontDialog_getFont2, "QFontDialog_getFont2")
	symaddr(&addr_QFontDialog_currentFontChanged, "QFontDialog_currentFontChanged")
	symaddr(&addr_QFontDialog_fontSelected, "QFontDialog_fontSelected")
	symaddr(&addr_QFontDialog_adjustPosition, "QFontDialog_adjustPosition")
	symaddr(&addr_QFontDialog_setTabOrder, "QFontDialog_setTabOrder")
	symaddr(&addr_QFontDialog_mouseGrabber, "QFontDialog_mouseGrabber")
	symaddr(&addr_QFontDialog_keyboardGrabber, "QFontDialog_keyboardGrabber")
	symaddr(&addr_QFontDialog_find, "QFontDialog_find")
	symaddr(&addr_QFontDialog_createWindowContainer, "QFontDialog_createWindowContainer")
	symaddr(&addr_QFontDialog_updateMicroFocus, "QFontDialog_updateMicroFocus")
	symaddr(&addr_QFontDialog_create, "QFontDialog_create")
	symaddr(&addr_QFontDialog_destroy, "QFontDialog_destroy")
	symaddr(&addr_QFontDialog_focusNextChild, "QFontDialog_focusNextChild")
	symaddr(&addr_QFontDialog_focusPreviousChild, "QFontDialog_focusPreviousChild")
	symaddr(&addr_QFontDialog_connect, "QFontDialog_connect")
	symaddr(&addr_QFontDialog_connect2, "QFontDialog_connect2")
	symaddr(&addr_QFontDialog_sender, "QFontDialog_sender")
	symaddr(&addr_QFontDialog_senderSignalIndex, "QFontDialog_senderSignalIndex")
	symaddr(&addr_QFontDialog_receivers, "QFontDialog_receivers")
	symaddr(&addr_QFontDialog_isSignalConnected, "QFontDialog_isSignalConnected")
	reg(&pg_QFontDialog_devicePixelRatioFScale, "QFontDialog_devicePixelRatioFScale")
	reg(&pg_QFontDialog_encodeMetricF, "QFontDialog_encodeMetricF")
	reg(&pg_QFontDialog_getDecodedMetricF, "QFontDialog_getDecodedMetricF")
	symaddr(&addr_QFontDialog_tr2, "QFontDialog_tr2")
	symaddr(&addr_QFontDialog_tr3, "QFontDialog_tr3")
	symaddr(&addr_QFontDialog_setOption2, "QFontDialog_setOption2")
	symaddr(&addr_QFontDialog_getFont3, "QFontDialog_getFont3")
	symaddr(&addr_QFontDialog_getFont4, "QFontDialog_getFont4")
	symaddr(&addr_QFontDialog_getFont5, "QFontDialog_getFont5")
	symaddr(&addr_QFontDialog_getFont6, "QFontDialog_getFont6")
	symaddr(&addr_QFontDialog_createWindowContainer2, "QFontDialog_createWindowContainer2")
	symaddr(&addr_QFontDialog_createWindowContainer3, "QFontDialog_createWindowContainer3")
	symaddr(&addr_QFontDialog_updateMicroFocus2, "QFontDialog_updateMicroFocus2")
	symaddr(&addr_QFontDialog_create2, "QFontDialog_create2")
	symaddr(&addr_QFontDialog_create3, "QFontDialog_create3")
	symaddr(&addr_QFontDialog_create4, "QFontDialog_create4")
	symaddr(&addr_QFontDialog_destroy2, "QFontDialog_destroy2")
	symaddr(&addr_QFontDialog_destroy3, "QFontDialog_destroy3")
	symaddr(&addr_QFontDialog_connect3, "QFontDialog_connect3")
	symaddr(&addr_QFontDialog_connect4, "QFontDialog_connect4")
	symaddr(&addr_QFontDialog_metaObject_default, "QFontDialog_metaObject_default")
	symaddr(&addr_QFontDialog_setVisible_default, "QFontDialog_setVisible_default")
	symaddr(&addr_QFontDialog_changeEvent_default, "QFontDialog_changeEvent_default")
	symaddr(&addr_QFontDialog_done_default, "QFontDialog_done_default")
	symaddr(&addr_QFontDialog_eventFilter_default, "QFontDialog_eventFilter_default")
	symaddr(&addr_QFontDialog_sizeHint_default, "QFontDialog_sizeHint_default")
	symaddr(&addr_QFontDialog_minimumSizeHint_default, "QFontDialog_minimumSizeHint_default")
	symaddr(&addr_QFontDialog_exec_default, "QFontDialog_exec_default")
	symaddr(&addr_QFontDialog_accept_default, "QFontDialog_accept_default")
	symaddr(&addr_QFontDialog_reject_default, "QFontDialog_reject_default")
	symaddr(&addr_QFontDialog_keyPressEvent_default, "QFontDialog_keyPressEvent_default")
	symaddr(&addr_QFontDialog_closeEvent_default, "QFontDialog_closeEvent_default")
	symaddr(&addr_QFontDialog_showEvent_default, "QFontDialog_showEvent_default")
	symaddr(&addr_QFontDialog_resizeEvent_default, "QFontDialog_resizeEvent_default")
	symaddr(&addr_QFontDialog_contextMenuEvent_default, "QFontDialog_contextMenuEvent_default")
	symaddr(&addr_QFontDialog_devType_default, "QFontDialog_devType_default")
	symaddr(&addr_QFontDialog_heightForWidth_default, "QFontDialog_heightForWidth_default")
	symaddr(&addr_QFontDialog_hasHeightForWidth_default, "QFontDialog_hasHeightForWidth_default")
	symaddr(&addr_QFontDialog_paintEngine_default, "QFontDialog_paintEngine_default")
	symaddr(&addr_QFontDialog_event_default, "QFontDialog_event_default")
	symaddr(&addr_QFontDialog_mousePressEvent_default, "QFontDialog_mousePressEvent_default")
	symaddr(&addr_QFontDialog_mouseReleaseEvent_default, "QFontDialog_mouseReleaseEvent_default")
	symaddr(&addr_QFontDialog_mouseDoubleClickEvent_default, "QFontDialog_mouseDoubleClickEvent_default")
	symaddr(&addr_QFontDialog_mouseMoveEvent_default, "QFontDialog_mouseMoveEvent_default")
	symaddr(&addr_QFontDialog_wheelEvent_default, "QFontDialog_wheelEvent_default")
	symaddr(&addr_QFontDialog_keyReleaseEvent_default, "QFontDialog_keyReleaseEvent_default")
	symaddr(&addr_QFontDialog_focusInEvent_default, "QFontDialog_focusInEvent_default")
	symaddr(&addr_QFontDialog_focusOutEvent_default, "QFontDialog_focusOutEvent_default")
	symaddr(&addr_QFontDialog_enterEvent_default, "QFontDialog_enterEvent_default")
	symaddr(&addr_QFontDialog_leaveEvent_default, "QFontDialog_leaveEvent_default")
	symaddr(&addr_QFontDialog_paintEvent_default, "QFontDialog_paintEvent_default")
	symaddr(&addr_QFontDialog_moveEvent_default, "QFontDialog_moveEvent_default")
	symaddr(&addr_QFontDialog_tabletEvent_default, "QFontDialog_tabletEvent_default")
	symaddr(&addr_QFontDialog_actionEvent_default, "QFontDialog_actionEvent_default")
	symaddr(&addr_QFontDialog_dragEnterEvent_default, "QFontDialog_dragEnterEvent_default")
	symaddr(&addr_QFontDialog_dragMoveEvent_default, "QFontDialog_dragMoveEvent_default")
	symaddr(&addr_QFontDialog_dragLeaveEvent_default, "QFontDialog_dragLeaveEvent_default")
	symaddr(&addr_QFontDialog_dropEvent_default, "QFontDialog_dropEvent_default")
	symaddr(&addr_QFontDialog_hideEvent_default, "QFontDialog_hideEvent_default")
	symaddr(&addr_QFontDialog_metric_default, "QFontDialog_metric_default")
	symaddr(&addr_QFontDialog_initPainter_default, "QFontDialog_initPainter_default")
	symaddr(&addr_QFontDialog_redirected_default, "QFontDialog_redirected_default")
	symaddr(&addr_QFontDialog_sharedPainter_default, "QFontDialog_sharedPainter_default")
	symaddr(&addr_QFontDialog_inputMethodEvent_default, "QFontDialog_inputMethodEvent_default")
	symaddr(&addr_QFontDialog_inputMethodQuery_default, "QFontDialog_inputMethodQuery_default")
	symaddr(&addr_QFontDialog_focusNextPrevChild_default, "QFontDialog_focusNextPrevChild_default")
	symaddr(&addr_QFontDialog_timerEvent_default, "QFontDialog_timerEvent_default")
	symaddr(&addr_QFontDialog_childEvent_default, "QFontDialog_childEvent_default")
	symaddr(&addr_QFontDialog_customEvent_default, "QFontDialog_customEvent_default")
	symaddr(&addr_QFontDialog_connectNotify_default, "QFontDialog_connectNotify_default")
	symaddr(&addr_QFontDialog_disconnectNotify_default, "QFontDialog_disconnectNotify_default")
	reg(&pg_QFontDialog_connect_currentFontChanged, "QFontDialog_connect_currentFontChanged")
	reg(&pg_QFontDialog_connect_fontSelected, "QFontDialog_connect_fontSelected")
	reg(&pg_QFontDialog_connect_finished, "QFontDialog_connect_finished")
	reg(&pg_QFontDialog_connect_accepted, "QFontDialog_connect_accepted")
	reg(&pg_QFontDialog_connect_rejected, "QFontDialog_connect_rejected")
	reg(&pg_QFontDialog_connect_windowTitleChanged, "QFontDialog_connect_windowTitleChanged")
	reg(&pg_QFontDialog_connect_windowIconChanged, "QFontDialog_connect_windowIconChanged")
	reg(&pg_QFontDialog_connect_windowIconTextChanged, "QFontDialog_connect_windowIconTextChanged")
	reg(&pg_QFontDialog_connect_customContextMenuRequested, "QFontDialog_connect_customContextMenuRequested")
	reg(&pg_QFontDialog_connect_destroyed, "QFontDialog_connect_destroyed")
	reg(&pg_QFontDialog_connect_destroyed2, "QFontDialog_connect_destroyed2")
	reg(&pg_QFontDialog_connect_objectNameChanged, "QFontDialog_connect_objectNameChanged")
	reg(&pg_QFontDialog__disconnect, "QFontDialog__disconnect")
	reg(&pg_QFontDialog__set_override, "QFontDialog__set_override")
	reg(&pg_QFontDialog_delete, "QFontDialog_delete")
}

type QFontDialog struct {
	h uintptr
	*QDialog
}

func bindQFontDialog(h uintptr) *QFontDialog {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QFontDialog_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFontDialog{h: h, QDialog: &QDialog{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQFontDialog(h uintptr) *QFontDialog { return bindQFontDialog(h) }

func (this *QFontDialog) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontDialog) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontDialog(h unsafe.Pointer) *QFontDialog { return bindQFontDialog(uintptr(h)) }
func UnsafeNewQFontDialogOwned(h unsafe.Pointer) *QFontDialog { return newownedQFontDialog(uintptr(h)) }

func (this *QFontDialog) UnsafeRelease() {
	this.h = 0
	if this.QDialog != nil {
		this.QDialog.UnsafeRelease()
	}
}

func NewQFontDialog(parent *QWidget) *QFontDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontDialog_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFontDialog(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontDialog_virtIdx, QFontDialog_enable)
	return __o
}

func NewQFontDialog2() *QFontDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontDialog_new2, uintptr(__goid))
	__o := newownedQFontDialog(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontDialog_virtIdx, QFontDialog_enable)
	return __o
}

func NewQFontDialog3(initial *QFont) *QFontDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontDialog_new3, uintptr(__goid), uintptr(initial.UnsafePointer()))
	__o := newownedQFontDialog(__r)
	runtime.KeepAlive(initial)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontDialog_virtIdx, QFontDialog_enable)
	return __o
}

func NewQFontDialog4(initial *QFont, parent *QWidget) *QFontDialog {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontDialog_new4, uintptr(__goid), uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQFontDialog(__r)
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontDialog_virtIdx, QFontDialog_enable)
	return __o
}

func QFontDialog_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDialog_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFontDialog) SetCurrentFont(font *QFont) {
	purego.SyscallN(addr_QFontDialog_setCurrentFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QFontDialog) CurrentFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_currentFont, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QFontDialog) SelectedFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_selectedFont, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QFontDialog) SetOption(option QFontDialog__FontDialogOption) {
	purego.SyscallN(addr_QFontDialog_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) TestOption(option QFontDialog__FontDialogOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontDialog) SetOptions(options QFontDialog__FontDialogOption) {
	purego.SyscallN(addr_QFontDialog_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Options() QFontDialog__FontDialogOption {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFontDialog__FontDialogOption(__r)
}

func (this *QFontDialog) Open(receiver *QObject, member string) {
	member_b := append([]byte(member), 0)
	purego.SyscallN(addr_QFontDialog_open, this.h, uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
}

func QFontDialog_GetFont() (*QFont, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont, uintptr(unsafe.Pointer(&ok_out)))
	return newownedQFont(__r), ok_out != 0
}

func QFontDialog_GetFont2(initial *QFont) (*QFont, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont2, uintptr(unsafe.Pointer(&ok_out)), uintptr(initial.UnsafePointer()))
	runtime.KeepAlive(initial)
	return newownedQFont(__r), ok_out != 0
}

func (this *QFontDialog) CurrentFontChanged(font *QFont) {
	purego.SyscallN(addr_QFontDialog_currentFontChanged, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QFontDialog) FontSelected(font *QFont) {
	purego.SyscallN(addr_QFontDialog_fontSelected, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QFontDialog) AdjustPosition(arg1 *QWidget) {
	purego.SyscallN(addr_QFontDialog_adjustPosition, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QFontDialog_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QFontDialog_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFontDialog_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_mouseGrabber)
	return bindQWidget(__r)
}

func QFontDialog_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_keyboardGrabber)
	return bindQWidget(__r)
}

func QFontDialog_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QFontDialog_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QFontDialog) UpdateMicroFocus() {
	purego.SyscallN(addr_QFontDialog_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Create() {
	purego.SyscallN(addr_QFontDialog_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Destroy() {
	purego.SyscallN(addr_QFontDialog_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontDialog) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontDialog) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFontDialog_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFontDialog_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFontDialog_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFontDialog) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFontDialog) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontDialog) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFontDialog) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFontDialog_DevicePixelRatioFScale() float64 {
	__r := pg_QFontDialog_devicePixelRatioFScale()
	return __r
}

func QFontDialog_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QFontDialog_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QFontDialog) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QFontDialog_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QFontDialog_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDialog_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFontDialog_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDialog_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QFontDialog) SetOption2(option QFontDialog__FontDialogOption, on bool) {
	purego.SyscallN(addr_QFontDialog_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QFontDialog_GetFont3(parent *QWidget) (*QFont, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont3, uintptr(unsafe.Pointer(&ok_out)), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
	return newownedQFont(__r), ok_out != 0
}

func QFontDialog_GetFont4(initial *QFont, parent *QWidget) (*QFont, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont4, uintptr(unsafe.Pointer(&ok_out)), uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	return newownedQFont(__r), ok_out != 0
}

func QFontDialog_GetFont5(initial *QFont, parent *QWidget, title string) (*QFont, bool) {
	var ok_out int32
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont5, uintptr(unsafe.Pointer(&ok_out)), uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	return newownedQFont(__r), ok_out != 0
}

func QFontDialog_GetFont6(initial *QFont, parent *QWidget, title string, options QFontDialog__FontDialogOption) (*QFont, bool) {
	var ok_out int32
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontDialog_getFont6, uintptr(unsafe.Pointer(&ok_out)), uintptr(initial.UnsafePointer()), uintptr(parent.UnsafePointer()), title_p, uintptr(len(title)), uintptr(options))
	runtime.KeepAlive(initial)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(title)
	return newownedQFont(__r), ok_out != 0
}

func QFontDialog_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QFontDialog_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QFontDialog) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QFontDialog_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QFontDialog_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QFontDialog_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QFontDialog_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QFontDialog_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QFontDialog_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QFontDialog_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFontDialog_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFontDialog_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFontDialog) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFontDialog) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QFontDialog_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFontDialog_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) DoneDefault(result int) {
	purego.SyscallN(addr_QFontDialog_done_default, this.h, uintptr(result))
	runtime.KeepAlive(this)
}

func (this *QFontDialog) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFontDialog) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFontDialog) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFontDialog) ExecDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_exec_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontDialog) AcceptDefault() {
	purego.SyscallN(addr_QFontDialog_accept_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontDialog) RejectDefault() {
	purego.SyscallN(addr_QFontDialog_reject_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontDialog) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QFontDialog_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QFontDialog_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QFontDialog_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QFontDialog_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QFontDialog_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontDialog) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontDialog) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontDialog) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QFontDialog) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFontDialog) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontDialog_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontDialog_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontDialog_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontDialog_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QFontDialog_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QFontDialog_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFontDialog_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QFontDialog_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QFontDialog_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFontDialog_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QFontDialog_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QFontDialog_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QFontDialog_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QFontDialog_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QFontDialog_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QFontDialog_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QFontDialog_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QFontDialog_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QFontDialog_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontDialog) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QFontDialog_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QFontDialog) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QFontDialog) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QFontDialog) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QFontDialog_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontDialog) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFontDialog_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFontDialog) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDialog_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontDialog) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFontDialog_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFontDialog_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFontDialog_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontDialog) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFontDialog_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFontDialog) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFontDialog_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFontDialog) OnCurrentFontChanged(fn func(*QFont)) int64 {
	id := registerCallback(func(font uintptr) { fn(UnsafeNewQFont(unsafe.Pointer(font))) })
	pg_QFontDialog_connect_currentFontChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnFontSelected(fn func(*QFont)) int64 {
	id := registerCallback(func(font uintptr) { fn(UnsafeNewQFont(unsafe.Pointer(font))) })
	pg_QFontDialog_connect_fontSelected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(result int32) { fn(int(result)) })
	pg_QFontDialog_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnAccepted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFontDialog_connect_accepted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnRejected(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFontDialog_connect_rejected(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QFontDialog_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QFontDialog_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QFontDialog_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QFontDialog_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFontDialog_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFontDialog_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFontDialog_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontDialog) Disconnect(handle int64) {
	pg_QFontDialog__disconnect(handle)
	unregisterCallback(handle)
}

func QFontDialog_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 2110182719375245575:
		return 1
	case 4959376063826165943:
		return 2
	case -2570997490807640431:
		return 3
	case 5217623514630858297:
		return 4
	case 729690261457684101:
		return 5
	case 7049264599837532491:
		return 6
	case -1745322561779819684:
		return 7
	case 9106131928185084681:
		return 8
	case 2875220963396170150:
		return 9
	case 1023791977679793185:
		return 10
	case -1989937050019464783:
		return 11
	case -6825369174751916:
		return 12
	case -4302917285655014661:
		return 13
	case -6698043222227006015:
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

func QFontDialog_enable(h uintptr, idx int32) { pg_QFontDialog__set_override(h, idx) }

func (this *QFontDialog) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFontDialog) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QFontDialog) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QFontDialog) OnDone(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2570997490807640431, func(result int32) { fn(int(result)) }) // slot "done"
}

func (this *QFontDialog) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFontDialog) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFontDialog) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QFontDialog) OnExec(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1745322561779819684, func() int32 { return int32(fn()) }) // slot "exec"
}

func (this *QFontDialog) OnAccept(fn func()) {
	registerVirt(this.UnsafeGoID(), 9106131928185084681, func() { fn() }) // slot "accept"
}

func (this *QFontDialog) OnReject(fn func()) {
	registerVirt(this.UnsafeGoID(), 2875220963396170150, func() { fn() }) // slot "reject"
}

func (this *QFontDialog) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QFontDialog) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QFontDialog) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QFontDialog) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QFontDialog) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QFontDialog) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QFontDialog) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QFontDialog) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFontDialog) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QFontDialog) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFontDialog) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QFontDialog) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QFontDialog) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QFontDialog) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QFontDialog) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QFontDialog) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QFontDialog) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QFontDialog) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QFontDialog) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QFontDialog) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QFontDialog) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QFontDialog) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QFontDialog) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QFontDialog) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QFontDialog) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QFontDialog) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QFontDialog) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QFontDialog) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QFontDialog) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QFontDialog) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QFontDialog) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QFontDialog) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QFontDialog) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QFontDialog) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QFontDialog) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QFontDialog) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QFontDialog) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QFontDialog) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFontDialog) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFontDialog) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFontDialog) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFontDialog) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFontDialog) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontDialog_delete(this.h)
	}
	this.UnsafeRelease()
}
