package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractButton_chain uintptr
	addr_QAbstractButton_tr uintptr
	addr_QAbstractButton_setText uintptr
	addr_QAbstractButton_text uintptr
	addr_QAbstractButton_setIcon uintptr
	addr_QAbstractButton_icon uintptr
	addr_QAbstractButton_iconSize uintptr
	addr_QAbstractButton_setShortcut uintptr
	addr_QAbstractButton_shortcut uintptr
	addr_QAbstractButton_setCheckable uintptr
	addr_QAbstractButton_isCheckable uintptr
	addr_QAbstractButton_isChecked uintptr
	addr_QAbstractButton_setDown uintptr
	addr_QAbstractButton_isDown uintptr
	addr_QAbstractButton_setAutoRepeat uintptr
	addr_QAbstractButton_autoRepeat uintptr
	addr_QAbstractButton_setAutoRepeatDelay uintptr
	addr_QAbstractButton_autoRepeatDelay uintptr
	addr_QAbstractButton_setAutoRepeatInterval uintptr
	addr_QAbstractButton_autoRepeatInterval uintptr
	addr_QAbstractButton_setAutoExclusive uintptr
	addr_QAbstractButton_autoExclusive uintptr
	addr_QAbstractButton_group uintptr
	addr_QAbstractButton_setIconSize uintptr
	addr_QAbstractButton_animateClick uintptr
	addr_QAbstractButton_click uintptr
	addr_QAbstractButton_toggle uintptr
	addr_QAbstractButton_setChecked uintptr
	addr_QAbstractButton_pressed uintptr
	addr_QAbstractButton_released uintptr
	addr_QAbstractButton_clicked uintptr
	addr_QAbstractButton_toggled uintptr
	addr_QAbstractButton_setTabOrder uintptr
	addr_QAbstractButton_mouseGrabber uintptr
	addr_QAbstractButton_keyboardGrabber uintptr
	addr_QAbstractButton_find uintptr
	addr_QAbstractButton_createWindowContainer uintptr
	addr_QAbstractButton_updateMicroFocus uintptr
	addr_QAbstractButton_create uintptr
	addr_QAbstractButton_destroy uintptr
	addr_QAbstractButton_focusNextChild uintptr
	addr_QAbstractButton_focusPreviousChild uintptr
	addr_QAbstractButton_connect uintptr
	addr_QAbstractButton_connect2 uintptr
	addr_QAbstractButton_sender uintptr
	addr_QAbstractButton_senderSignalIndex uintptr
	addr_QAbstractButton_receivers uintptr
	addr_QAbstractButton_isSignalConnected uintptr
	pg_QAbstractButton_devicePixelRatioFScale func() float64
	pg_QAbstractButton_encodeMetricF func(int32, float64) int32
	pg_QAbstractButton_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QAbstractButton_tr2 uintptr
	addr_QAbstractButton_tr3 uintptr
	addr_QAbstractButton_clicked2 uintptr
	addr_QAbstractButton_createWindowContainer2 uintptr
	addr_QAbstractButton_createWindowContainer3 uintptr
	addr_QAbstractButton_updateMicroFocus2 uintptr
	addr_QAbstractButton_create2 uintptr
	addr_QAbstractButton_create3 uintptr
	addr_QAbstractButton_create4 uintptr
	addr_QAbstractButton_destroy2 uintptr
	addr_QAbstractButton_destroy3 uintptr
	addr_QAbstractButton_connect3 uintptr
	addr_QAbstractButton_connect4 uintptr
	addr_QAbstractButton_metaObject_default uintptr
	addr_QAbstractButton_hitButton_default uintptr
	addr_QAbstractButton_checkStateSet_default uintptr
	addr_QAbstractButton_nextCheckState_default uintptr
	addr_QAbstractButton_event_default uintptr
	addr_QAbstractButton_keyPressEvent_default uintptr
	addr_QAbstractButton_keyReleaseEvent_default uintptr
	addr_QAbstractButton_mousePressEvent_default uintptr
	addr_QAbstractButton_mouseReleaseEvent_default uintptr
	addr_QAbstractButton_mouseMoveEvent_default uintptr
	addr_QAbstractButton_focusInEvent_default uintptr
	addr_QAbstractButton_focusOutEvent_default uintptr
	addr_QAbstractButton_changeEvent_default uintptr
	addr_QAbstractButton_timerEvent_default uintptr
	addr_QAbstractButton_devType_default uintptr
	addr_QAbstractButton_setVisible_default uintptr
	addr_QAbstractButton_sizeHint_default uintptr
	addr_QAbstractButton_minimumSizeHint_default uintptr
	addr_QAbstractButton_heightForWidth_default uintptr
	addr_QAbstractButton_hasHeightForWidth_default uintptr
	addr_QAbstractButton_paintEngine_default uintptr
	addr_QAbstractButton_mouseDoubleClickEvent_default uintptr
	addr_QAbstractButton_wheelEvent_default uintptr
	addr_QAbstractButton_enterEvent_default uintptr
	addr_QAbstractButton_leaveEvent_default uintptr
	addr_QAbstractButton_moveEvent_default uintptr
	addr_QAbstractButton_resizeEvent_default uintptr
	addr_QAbstractButton_closeEvent_default uintptr
	addr_QAbstractButton_contextMenuEvent_default uintptr
	addr_QAbstractButton_tabletEvent_default uintptr
	addr_QAbstractButton_actionEvent_default uintptr
	addr_QAbstractButton_dragEnterEvent_default uintptr
	addr_QAbstractButton_dragMoveEvent_default uintptr
	addr_QAbstractButton_dragLeaveEvent_default uintptr
	addr_QAbstractButton_dropEvent_default uintptr
	addr_QAbstractButton_showEvent_default uintptr
	addr_QAbstractButton_hideEvent_default uintptr
	addr_QAbstractButton_metric_default uintptr
	addr_QAbstractButton_initPainter_default uintptr
	addr_QAbstractButton_redirected_default uintptr
	addr_QAbstractButton_sharedPainter_default uintptr
	addr_QAbstractButton_inputMethodEvent_default uintptr
	addr_QAbstractButton_inputMethodQuery_default uintptr
	addr_QAbstractButton_focusNextPrevChild_default uintptr
	addr_QAbstractButton_eventFilter_default uintptr
	addr_QAbstractButton_childEvent_default uintptr
	addr_QAbstractButton_customEvent_default uintptr
	addr_QAbstractButton_connectNotify_default uintptr
	addr_QAbstractButton_disconnectNotify_default uintptr
	pg_QAbstractButton_connect_pressed func(uintptr, int64)
	pg_QAbstractButton_connect_released func(uintptr, int64)
	pg_QAbstractButton_connect_clicked func(uintptr, int64)
	pg_QAbstractButton_connect_clicked2 func(uintptr, int64)
	pg_QAbstractButton_connect_toggled func(uintptr, int64)
	pg_QAbstractButton_connect_windowTitleChanged func(uintptr, int64)
	pg_QAbstractButton_connect_windowIconChanged func(uintptr, int64)
	pg_QAbstractButton_connect_windowIconTextChanged func(uintptr, int64)
	pg_QAbstractButton_connect_customContextMenuRequested func(uintptr, int64)
	pg_QAbstractButton_connect_destroyed func(uintptr, int64)
	pg_QAbstractButton_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractButton_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractButton__disconnect func(int64)
	pg_QAbstractButton__set_override func(uintptr, int32)
	pg_QAbstractButton__Go_new func(int64) uintptr
	pg_QAbstractButton_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractButton_chain, "QAbstractButton_chain")
	symaddr(&addr_QAbstractButton_tr, "QAbstractButton_tr")
	symaddr(&addr_QAbstractButton_setText, "QAbstractButton_setText")
	symaddr(&addr_QAbstractButton_text, "QAbstractButton_text")
	symaddr(&addr_QAbstractButton_setIcon, "QAbstractButton_setIcon")
	symaddr(&addr_QAbstractButton_icon, "QAbstractButton_icon")
	symaddr(&addr_QAbstractButton_iconSize, "QAbstractButton_iconSize")
	symaddr(&addr_QAbstractButton_setShortcut, "QAbstractButton_setShortcut")
	symaddr(&addr_QAbstractButton_shortcut, "QAbstractButton_shortcut")
	symaddr(&addr_QAbstractButton_setCheckable, "QAbstractButton_setCheckable")
	symaddr(&addr_QAbstractButton_isCheckable, "QAbstractButton_isCheckable")
	symaddr(&addr_QAbstractButton_isChecked, "QAbstractButton_isChecked")
	symaddr(&addr_QAbstractButton_setDown, "QAbstractButton_setDown")
	symaddr(&addr_QAbstractButton_isDown, "QAbstractButton_isDown")
	symaddr(&addr_QAbstractButton_setAutoRepeat, "QAbstractButton_setAutoRepeat")
	symaddr(&addr_QAbstractButton_autoRepeat, "QAbstractButton_autoRepeat")
	symaddr(&addr_QAbstractButton_setAutoRepeatDelay, "QAbstractButton_setAutoRepeatDelay")
	symaddr(&addr_QAbstractButton_autoRepeatDelay, "QAbstractButton_autoRepeatDelay")
	symaddr(&addr_QAbstractButton_setAutoRepeatInterval, "QAbstractButton_setAutoRepeatInterval")
	symaddr(&addr_QAbstractButton_autoRepeatInterval, "QAbstractButton_autoRepeatInterval")
	symaddr(&addr_QAbstractButton_setAutoExclusive, "QAbstractButton_setAutoExclusive")
	symaddr(&addr_QAbstractButton_autoExclusive, "QAbstractButton_autoExclusive")
	symaddr(&addr_QAbstractButton_group, "QAbstractButton_group")
	symaddr(&addr_QAbstractButton_setIconSize, "QAbstractButton_setIconSize")
	symaddr(&addr_QAbstractButton_animateClick, "QAbstractButton_animateClick")
	symaddr(&addr_QAbstractButton_click, "QAbstractButton_click")
	symaddr(&addr_QAbstractButton_toggle, "QAbstractButton_toggle")
	symaddr(&addr_QAbstractButton_setChecked, "QAbstractButton_setChecked")
	symaddr(&addr_QAbstractButton_pressed, "QAbstractButton_pressed")
	symaddr(&addr_QAbstractButton_released, "QAbstractButton_released")
	symaddr(&addr_QAbstractButton_clicked, "QAbstractButton_clicked")
	symaddr(&addr_QAbstractButton_toggled, "QAbstractButton_toggled")
	symaddr(&addr_QAbstractButton_setTabOrder, "QAbstractButton_setTabOrder")
	symaddr(&addr_QAbstractButton_mouseGrabber, "QAbstractButton_mouseGrabber")
	symaddr(&addr_QAbstractButton_keyboardGrabber, "QAbstractButton_keyboardGrabber")
	symaddr(&addr_QAbstractButton_find, "QAbstractButton_find")
	symaddr(&addr_QAbstractButton_createWindowContainer, "QAbstractButton_createWindowContainer")
	symaddr(&addr_QAbstractButton_updateMicroFocus, "QAbstractButton_updateMicroFocus")
	symaddr(&addr_QAbstractButton_create, "QAbstractButton_create")
	symaddr(&addr_QAbstractButton_destroy, "QAbstractButton_destroy")
	symaddr(&addr_QAbstractButton_focusNextChild, "QAbstractButton_focusNextChild")
	symaddr(&addr_QAbstractButton_focusPreviousChild, "QAbstractButton_focusPreviousChild")
	symaddr(&addr_QAbstractButton_connect, "QAbstractButton_connect")
	symaddr(&addr_QAbstractButton_connect2, "QAbstractButton_connect2")
	symaddr(&addr_QAbstractButton_sender, "QAbstractButton_sender")
	symaddr(&addr_QAbstractButton_senderSignalIndex, "QAbstractButton_senderSignalIndex")
	symaddr(&addr_QAbstractButton_receivers, "QAbstractButton_receivers")
	symaddr(&addr_QAbstractButton_isSignalConnected, "QAbstractButton_isSignalConnected")
	reg(&pg_QAbstractButton_devicePixelRatioFScale, "QAbstractButton_devicePixelRatioFScale")
	reg(&pg_QAbstractButton_encodeMetricF, "QAbstractButton_encodeMetricF")
	reg(&pg_QAbstractButton_getDecodedMetricF, "QAbstractButton_getDecodedMetricF")
	symaddr(&addr_QAbstractButton_tr2, "QAbstractButton_tr2")
	symaddr(&addr_QAbstractButton_tr3, "QAbstractButton_tr3")
	symaddr(&addr_QAbstractButton_clicked2, "QAbstractButton_clicked2")
	symaddr(&addr_QAbstractButton_createWindowContainer2, "QAbstractButton_createWindowContainer2")
	symaddr(&addr_QAbstractButton_createWindowContainer3, "QAbstractButton_createWindowContainer3")
	symaddr(&addr_QAbstractButton_updateMicroFocus2, "QAbstractButton_updateMicroFocus2")
	symaddr(&addr_QAbstractButton_create2, "QAbstractButton_create2")
	symaddr(&addr_QAbstractButton_create3, "QAbstractButton_create3")
	symaddr(&addr_QAbstractButton_create4, "QAbstractButton_create4")
	symaddr(&addr_QAbstractButton_destroy2, "QAbstractButton_destroy2")
	symaddr(&addr_QAbstractButton_destroy3, "QAbstractButton_destroy3")
	symaddr(&addr_QAbstractButton_connect3, "QAbstractButton_connect3")
	symaddr(&addr_QAbstractButton_connect4, "QAbstractButton_connect4")
	symaddr(&addr_QAbstractButton_metaObject_default, "QAbstractButton_metaObject_default")
	symaddr(&addr_QAbstractButton_hitButton_default, "QAbstractButton_hitButton_default")
	symaddr(&addr_QAbstractButton_checkStateSet_default, "QAbstractButton_checkStateSet_default")
	symaddr(&addr_QAbstractButton_nextCheckState_default, "QAbstractButton_nextCheckState_default")
	symaddr(&addr_QAbstractButton_event_default, "QAbstractButton_event_default")
	symaddr(&addr_QAbstractButton_keyPressEvent_default, "QAbstractButton_keyPressEvent_default")
	symaddr(&addr_QAbstractButton_keyReleaseEvent_default, "QAbstractButton_keyReleaseEvent_default")
	symaddr(&addr_QAbstractButton_mousePressEvent_default, "QAbstractButton_mousePressEvent_default")
	symaddr(&addr_QAbstractButton_mouseReleaseEvent_default, "QAbstractButton_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractButton_mouseMoveEvent_default, "QAbstractButton_mouseMoveEvent_default")
	symaddr(&addr_QAbstractButton_focusInEvent_default, "QAbstractButton_focusInEvent_default")
	symaddr(&addr_QAbstractButton_focusOutEvent_default, "QAbstractButton_focusOutEvent_default")
	symaddr(&addr_QAbstractButton_changeEvent_default, "QAbstractButton_changeEvent_default")
	symaddr(&addr_QAbstractButton_timerEvent_default, "QAbstractButton_timerEvent_default")
	symaddr(&addr_QAbstractButton_devType_default, "QAbstractButton_devType_default")
	symaddr(&addr_QAbstractButton_setVisible_default, "QAbstractButton_setVisible_default")
	symaddr(&addr_QAbstractButton_sizeHint_default, "QAbstractButton_sizeHint_default")
	symaddr(&addr_QAbstractButton_minimumSizeHint_default, "QAbstractButton_minimumSizeHint_default")
	symaddr(&addr_QAbstractButton_heightForWidth_default, "QAbstractButton_heightForWidth_default")
	symaddr(&addr_QAbstractButton_hasHeightForWidth_default, "QAbstractButton_hasHeightForWidth_default")
	symaddr(&addr_QAbstractButton_paintEngine_default, "QAbstractButton_paintEngine_default")
	symaddr(&addr_QAbstractButton_mouseDoubleClickEvent_default, "QAbstractButton_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractButton_wheelEvent_default, "QAbstractButton_wheelEvent_default")
	symaddr(&addr_QAbstractButton_enterEvent_default, "QAbstractButton_enterEvent_default")
	symaddr(&addr_QAbstractButton_leaveEvent_default, "QAbstractButton_leaveEvent_default")
	symaddr(&addr_QAbstractButton_moveEvent_default, "QAbstractButton_moveEvent_default")
	symaddr(&addr_QAbstractButton_resizeEvent_default, "QAbstractButton_resizeEvent_default")
	symaddr(&addr_QAbstractButton_closeEvent_default, "QAbstractButton_closeEvent_default")
	symaddr(&addr_QAbstractButton_contextMenuEvent_default, "QAbstractButton_contextMenuEvent_default")
	symaddr(&addr_QAbstractButton_tabletEvent_default, "QAbstractButton_tabletEvent_default")
	symaddr(&addr_QAbstractButton_actionEvent_default, "QAbstractButton_actionEvent_default")
	symaddr(&addr_QAbstractButton_dragEnterEvent_default, "QAbstractButton_dragEnterEvent_default")
	symaddr(&addr_QAbstractButton_dragMoveEvent_default, "QAbstractButton_dragMoveEvent_default")
	symaddr(&addr_QAbstractButton_dragLeaveEvent_default, "QAbstractButton_dragLeaveEvent_default")
	symaddr(&addr_QAbstractButton_dropEvent_default, "QAbstractButton_dropEvent_default")
	symaddr(&addr_QAbstractButton_showEvent_default, "QAbstractButton_showEvent_default")
	symaddr(&addr_QAbstractButton_hideEvent_default, "QAbstractButton_hideEvent_default")
	symaddr(&addr_QAbstractButton_metric_default, "QAbstractButton_metric_default")
	symaddr(&addr_QAbstractButton_initPainter_default, "QAbstractButton_initPainter_default")
	symaddr(&addr_QAbstractButton_redirected_default, "QAbstractButton_redirected_default")
	symaddr(&addr_QAbstractButton_sharedPainter_default, "QAbstractButton_sharedPainter_default")
	symaddr(&addr_QAbstractButton_inputMethodEvent_default, "QAbstractButton_inputMethodEvent_default")
	symaddr(&addr_QAbstractButton_inputMethodQuery_default, "QAbstractButton_inputMethodQuery_default")
	symaddr(&addr_QAbstractButton_focusNextPrevChild_default, "QAbstractButton_focusNextPrevChild_default")
	symaddr(&addr_QAbstractButton_eventFilter_default, "QAbstractButton_eventFilter_default")
	symaddr(&addr_QAbstractButton_childEvent_default, "QAbstractButton_childEvent_default")
	symaddr(&addr_QAbstractButton_customEvent_default, "QAbstractButton_customEvent_default")
	symaddr(&addr_QAbstractButton_connectNotify_default, "QAbstractButton_connectNotify_default")
	symaddr(&addr_QAbstractButton_disconnectNotify_default, "QAbstractButton_disconnectNotify_default")
	reg(&pg_QAbstractButton_connect_pressed, "QAbstractButton_connect_pressed")
	reg(&pg_QAbstractButton_connect_released, "QAbstractButton_connect_released")
	reg(&pg_QAbstractButton_connect_clicked, "QAbstractButton_connect_clicked")
	reg(&pg_QAbstractButton_connect_clicked2, "QAbstractButton_connect_clicked2")
	reg(&pg_QAbstractButton_connect_toggled, "QAbstractButton_connect_toggled")
	reg(&pg_QAbstractButton_connect_windowTitleChanged, "QAbstractButton_connect_windowTitleChanged")
	reg(&pg_QAbstractButton_connect_windowIconChanged, "QAbstractButton_connect_windowIconChanged")
	reg(&pg_QAbstractButton_connect_windowIconTextChanged, "QAbstractButton_connect_windowIconTextChanged")
	reg(&pg_QAbstractButton_connect_customContextMenuRequested, "QAbstractButton_connect_customContextMenuRequested")
	reg(&pg_QAbstractButton_connect_destroyed, "QAbstractButton_connect_destroyed")
	reg(&pg_QAbstractButton_connect_destroyed2, "QAbstractButton_connect_destroyed2")
	reg(&pg_QAbstractButton_connect_objectNameChanged, "QAbstractButton_connect_objectNameChanged")
	reg(&pg_QAbstractButton__disconnect, "QAbstractButton__disconnect")
	reg(&pg_QAbstractButton__set_override, "QAbstractButton__set_override")
	reg(&pg_QAbstractButton__Go_new, "QAbstractButton__Go_new")
	reg(&pg_QAbstractButton_delete, "QAbstractButton_delete")
}

type QAbstractButton struct {
	h uintptr
	*QWidget
}

func bindQAbstractButton(h uintptr) *QAbstractButton {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractButton_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractButton{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractButton(h uintptr) *QAbstractButton { return bindQAbstractButton(h) }

func (this *QAbstractButton) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractButton) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractButton(h unsafe.Pointer) *QAbstractButton { return bindQAbstractButton(uintptr(h)) }
func UnsafeNewQAbstractButtonOwned(h unsafe.Pointer) *QAbstractButton { return newownedQAbstractButton(uintptr(h)) }

func (this *QAbstractButton) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func QAbstractButton_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractButton) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAbstractButton_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAbstractButton) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAbstractButton) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QAbstractButton_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QAbstractButton) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QAbstractButton) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractButton) SetShortcut(key *QKeySequence) {
	purego.SyscallN(addr_QAbstractButton_setShortcut, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QAbstractButton) Shortcut() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_shortcut, this.h)
	runtime.KeepAlive(this)
	return newownedQKeySequence(__r)
}

func (this *QAbstractButton) SetCheckable(arg1 bool) {
	purego.SyscallN(addr_QAbstractButton_setCheckable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) IsCheckable() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_isCheckable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) IsChecked() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_isChecked, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) SetDown(arg1 bool) {
	purego.SyscallN(addr_QAbstractButton_setDown, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) IsDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_isDown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) SetAutoRepeat(arg1 bool) {
	purego.SyscallN(addr_QAbstractButton_setAutoRepeat, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) AutoRepeat() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_autoRepeat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) SetAutoRepeatDelay(arg1 int) {
	purego.SyscallN(addr_QAbstractButton_setAutoRepeatDelay, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) AutoRepeatDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_autoRepeatDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) SetAutoRepeatInterval(arg1 int) {
	purego.SyscallN(addr_QAbstractButton_setAutoRepeatInterval, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) AutoRepeatInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_autoRepeatInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) SetAutoExclusive(arg1 bool) {
	purego.SyscallN(addr_QAbstractButton_setAutoExclusive, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) AutoExclusive() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_autoExclusive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) Group() *QButtonGroup {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_group, this.h)
	runtime.KeepAlive(this)
	return bindQButtonGroup(__r)
}

func (this *QAbstractButton) SetIconSize(size *QSize) {
	purego.SyscallN(addr_QAbstractButton_setIconSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QAbstractButton) AnimateClick() {
	purego.SyscallN(addr_QAbstractButton_animateClick, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Click() {
	purego.SyscallN(addr_QAbstractButton_click, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Toggle() {
	purego.SyscallN(addr_QAbstractButton_toggle, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) SetChecked(arg1 bool) {
	purego.SyscallN(addr_QAbstractButton_setChecked, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Pressed() {
	purego.SyscallN(addr_QAbstractButton_pressed, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Released() {
	purego.SyscallN(addr_QAbstractButton_released, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Clicked() {
	purego.SyscallN(addr_QAbstractButton_clicked, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Toggled(checked bool) {
	purego.SyscallN(addr_QAbstractButton_toggled, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func QAbstractButton_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QAbstractButton_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QAbstractButton_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_mouseGrabber)
	return bindQWidget(__r)
}

func QAbstractButton_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_keyboardGrabber)
	return bindQWidget(__r)
}

func QAbstractButton_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QAbstractButton_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QAbstractButton) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractButton_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Create() {
	purego.SyscallN(addr_QAbstractButton_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Destroy() {
	purego.SyscallN(addr_QAbstractButton_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAbstractButton_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractButton_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractButton) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractButton) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractButton) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractButton_DevicePixelRatioFScale() float64 {
	__r := pg_QAbstractButton_devicePixelRatioFScale()
	return __r
}

func QAbstractButton_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QAbstractButton_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QAbstractButton) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QAbstractButton_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QAbstractButton_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractButton_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractButton) ClickedWithChecked(checked bool) {
	purego.SyscallN(addr_QAbstractButton_clicked2, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func QAbstractButton_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QAbstractButton_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QAbstractButton) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QAbstractButton_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QAbstractButton_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QAbstractButton_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QAbstractButton_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QAbstractButton_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QAbstractButton_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QAbstractButton_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractButton_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractButton) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractButton) HitButtonDefault(pos *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_hitButton_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return __r != 0
}

func (this *QAbstractButton) CheckStateSetDefault() {
	purego.SyscallN(addr_QAbstractButton_checkStateSet_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) NextCheckStateDefault() {
	purego.SyscallN(addr_QAbstractButton_nextCheckState_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QAbstractButton) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QAbstractButton_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QAbstractButton_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QAbstractButton_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QAbstractButton_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) MouseMoveEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QAbstractButton_mouseMoveEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QAbstractButton_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QAbstractButton_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QAbstractButton_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) TimerEventDefault(e *QTimerEvent) {
	purego.SyscallN(addr_QAbstractButton_timerEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QAbstractButton) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QAbstractButton_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAbstractButton) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractButton) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractButton) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QAbstractButton) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QAbstractButton_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QAbstractButton_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QAbstractButton_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractButton_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QAbstractButton_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QAbstractButton_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QAbstractButton_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QAbstractButton_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QAbstractButton_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QAbstractButton_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QAbstractButton_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QAbstractButton_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QAbstractButton_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QAbstractButton_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QAbstractButton_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QAbstractButton_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractButton) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QAbstractButton_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QAbstractButton) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QAbstractButton) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QAbstractButton) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractButton_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractButton) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractButton_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractButton) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractButton) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractButton_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractButton) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractButton_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractButton_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractButton) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractButton_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractButton) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractButton_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractButton) OnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractButton_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnReleased(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractButton_connect_released(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractButton_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnClickedWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QAbstractButton_connect_clicked2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QAbstractButton_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QAbstractButton_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QAbstractButton_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QAbstractButton_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QAbstractButton_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractButton_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractButton_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractButton_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractButton) Disconnect(handle int64) {
	pg_QAbstractButton__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractButton_virtIdx(slot int64) int32 {
	switch slot {
	case 5290442163167458465:
		return 0
	case 1607084412591819503:
		return 1
	case 5736040353629178930:
		return 2
	case 5747806663224476200:
		return 3
	case 1474647914901213797:
		return 4
	case 1237236456312824991:
		return 5
	case 1023791977679793185:
		return 6
	case 772598475950050299:
		return 7
	case -332090898078543319:
		return 8
	case -4415920049297795213:
		return 9
	case -4255494004099806581:
		return 10
	case 975043767033552552:
		return 11
	case -385596010576048993:
		return 12
	case 4959376063826165943:
		return 13
	case -5409927581812374498:
		return 14
	case -8659679893868524044:
		return 15
	case 2110182719375245575:
		return 16
	case 729690261457684101:
		return 17
	case 7049264599837532491:
		return 18
	case 4311534060495737339:
		return 19
	case -8266762397995423915:
		return 20
	case 6398007404756346381:
		return 21
	case -4029073905697023887:
		return 22
	case 1919988340934511316:
		return 23
	case 6316270346668237895:
		return 24
	case -551070009623711640:
		return 25
	case -9161949150489525396:
		return 26
	case -4302917285655014661:
		return 27
	case -1989937050019464783:
		return 28
	case -6698043222227006015:
		return 29
	case -212068866860701141:
		return 30
	case 1697444140088555581:
		return 31
	case 206122715919917589:
		return 32
	case -1781545137352745978:
		return 33
	case -3139904158373725006:
		return 34
	case 9107058682754591716:
		return 35
	case -6825369174751916:
		return 36
	case -199718707876719129:
		return 37
	case -2039166586246443706:
		return 38
	case 427629649267311777:
		return 39
	case 6477500296980050210:
		return 40
	case 4127240633310655048:
		return 41
	case 5792771328741774991:
		return 42
	case 5438819661942843222:
		return 43
	case 7345853097340709638:
		return 44
	case 5898328517505784191:
		return 45
	case 5217623514630858297:
		return 46
	case -1185315367526857833:
		return 47
	case 3769868229123868474:
		return 48
	case 7379029982170503008:
		return 49
	case -568369772982215962:
		return 50
	}
	return -1
}

func QAbstractButton_enable(h uintptr, idx int32) { pg_QAbstractButton__set_override(h, idx) }

func NewQAbstractButton() *QAbstractButton {
	id := nextVirtID()
	o := newownedQAbstractButton(pg_QAbstractButton__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractButton_virtIdx, QAbstractButton_enable)
	return o
}

func (this *QAbstractButton) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QAbstractButton) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractButton) OnHitButton(fn func(*QPoint) bool) {
	registerVirt(this.UnsafeGoID(), 5736040353629178930, func(pos uintptr) int32 { return b2i(fn(UnsafeNewQPoint(unsafe.Pointer(pos)))) }) // slot "hitButton"
}

func (this *QAbstractButton) OnCheckStateSet(fn func()) {
	registerVirt(this.UnsafeGoID(), 5747806663224476200, func() { fn() }) // slot "checkStateSet"
}

func (this *QAbstractButton) OnNextCheckState(fn func()) {
	registerVirt(this.UnsafeGoID(), 1474647914901213797, func() { fn() }) // slot "nextCheckState"
}

func (this *QAbstractButton) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QAbstractButton) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QAbstractButton) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractButton) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QAbstractButton) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractButton) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractButton) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QAbstractButton) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QAbstractButton) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QAbstractButton) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(e uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(e))) }) // slot "timerEvent"
}

func (this *QAbstractButton) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QAbstractButton) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QAbstractButton) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractButton) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QAbstractButton) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QAbstractButton) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QAbstractButton) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QAbstractButton) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractButton) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QAbstractButton) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QAbstractButton) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QAbstractButton) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QAbstractButton) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QAbstractButton) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QAbstractButton) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QAbstractButton) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QAbstractButton) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QAbstractButton) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QAbstractButton) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QAbstractButton) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractButton) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QAbstractButton) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QAbstractButton) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QAbstractButton) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QAbstractButton) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QAbstractButton) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QAbstractButton) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QAbstractButton) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QAbstractButton) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QAbstractButton) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractButton) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QAbstractButton) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractButton) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractButton) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractButton) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractButton) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractButton) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractButton_delete(this.h)
	}
	this.UnsafeRelease()
}
