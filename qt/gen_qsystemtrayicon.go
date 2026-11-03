package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSystemTrayIcon_base func(uintptr) uintptr
	addr_QSystemTrayIcon_new uintptr
	addr_QSystemTrayIcon_new2 uintptr
	addr_QSystemTrayIcon_new3 uintptr
	addr_QSystemTrayIcon_new4 uintptr
	addr_QSystemTrayIcon_tr uintptr
	addr_QSystemTrayIcon_setContextMenu uintptr
	addr_QSystemTrayIcon_contextMenu uintptr
	addr_QSystemTrayIcon_icon uintptr
	addr_QSystemTrayIcon_setIcon uintptr
	addr_QSystemTrayIcon_toolTip uintptr
	addr_QSystemTrayIcon_setToolTip uintptr
	addr_QSystemTrayIcon_isSystemTrayAvailable uintptr
	addr_QSystemTrayIcon_supportsMessages uintptr
	addr_QSystemTrayIcon_geometry uintptr
	addr_QSystemTrayIcon_isVisible uintptr
	addr_QSystemTrayIcon_setVisible uintptr
	addr_QSystemTrayIcon_show uintptr
	addr_QSystemTrayIcon_hide uintptr
	addr_QSystemTrayIcon_showMessage uintptr
	addr_QSystemTrayIcon_showMessage2 uintptr
	addr_QSystemTrayIcon_activated uintptr
	addr_QSystemTrayIcon_messageClicked uintptr
	addr_QSystemTrayIcon_connect uintptr
	addr_QSystemTrayIcon_connect2 uintptr
	addr_QSystemTrayIcon_sender uintptr
	addr_QSystemTrayIcon_senderSignalIndex uintptr
	addr_QSystemTrayIcon_receivers uintptr
	addr_QSystemTrayIcon_isSignalConnected uintptr
	addr_QSystemTrayIcon_tr2 uintptr
	addr_QSystemTrayIcon_tr3 uintptr
	addr_QSystemTrayIcon_showMessage3 uintptr
	addr_QSystemTrayIcon_showMessage4 uintptr
	addr_QSystemTrayIcon_showMessage5 uintptr
	addr_QSystemTrayIcon_connect3 uintptr
	addr_QSystemTrayIcon_connect4 uintptr
	addr_QSystemTrayIcon_metaObject_default uintptr
	addr_QSystemTrayIcon_event_default uintptr
	addr_QSystemTrayIcon_eventFilter_default uintptr
	addr_QSystemTrayIcon_timerEvent_default uintptr
	addr_QSystemTrayIcon_childEvent_default uintptr
	addr_QSystemTrayIcon_customEvent_default uintptr
	addr_QSystemTrayIcon_connectNotify_default uintptr
	addr_QSystemTrayIcon_disconnectNotify_default uintptr
	pg_QSystemTrayIcon_connect_activated func(uintptr, int64)
	pg_QSystemTrayIcon_connect_messageClicked func(uintptr, int64)
	pg_QSystemTrayIcon_connect_destroyed func(uintptr, int64)
	pg_QSystemTrayIcon_connect_destroyed2 func(uintptr, int64)
	pg_QSystemTrayIcon_connect_objectNameChanged func(uintptr, int64)
	pg_QSystemTrayIcon__disconnect func(int64)
	pg_QSystemTrayIcon__set_override func(uintptr, int32)
	pg_QSystemTrayIcon_delete func(uintptr)
)

func init() {
	reg(&pg_QSystemTrayIcon_base, "QSystemTrayIcon_base")
	symaddr(&addr_QSystemTrayIcon_new, "QSystemTrayIcon_new")
	symaddr(&addr_QSystemTrayIcon_new2, "QSystemTrayIcon_new2")
	symaddr(&addr_QSystemTrayIcon_new3, "QSystemTrayIcon_new3")
	symaddr(&addr_QSystemTrayIcon_new4, "QSystemTrayIcon_new4")
	symaddr(&addr_QSystemTrayIcon_tr, "QSystemTrayIcon_tr")
	symaddr(&addr_QSystemTrayIcon_setContextMenu, "QSystemTrayIcon_setContextMenu")
	symaddr(&addr_QSystemTrayIcon_contextMenu, "QSystemTrayIcon_contextMenu")
	symaddr(&addr_QSystemTrayIcon_icon, "QSystemTrayIcon_icon")
	symaddr(&addr_QSystemTrayIcon_setIcon, "QSystemTrayIcon_setIcon")
	symaddr(&addr_QSystemTrayIcon_toolTip, "QSystemTrayIcon_toolTip")
	symaddr(&addr_QSystemTrayIcon_setToolTip, "QSystemTrayIcon_setToolTip")
	symaddr(&addr_QSystemTrayIcon_isSystemTrayAvailable, "QSystemTrayIcon_isSystemTrayAvailable")
	symaddr(&addr_QSystemTrayIcon_supportsMessages, "QSystemTrayIcon_supportsMessages")
	symaddr(&addr_QSystemTrayIcon_geometry, "QSystemTrayIcon_geometry")
	symaddr(&addr_QSystemTrayIcon_isVisible, "QSystemTrayIcon_isVisible")
	symaddr(&addr_QSystemTrayIcon_setVisible, "QSystemTrayIcon_setVisible")
	symaddr(&addr_QSystemTrayIcon_show, "QSystemTrayIcon_show")
	symaddr(&addr_QSystemTrayIcon_hide, "QSystemTrayIcon_hide")
	symaddr(&addr_QSystemTrayIcon_showMessage, "QSystemTrayIcon_showMessage")
	symaddr(&addr_QSystemTrayIcon_showMessage2, "QSystemTrayIcon_showMessage2")
	symaddr(&addr_QSystemTrayIcon_activated, "QSystemTrayIcon_activated")
	symaddr(&addr_QSystemTrayIcon_messageClicked, "QSystemTrayIcon_messageClicked")
	symaddr(&addr_QSystemTrayIcon_connect, "QSystemTrayIcon_connect")
	symaddr(&addr_QSystemTrayIcon_connect2, "QSystemTrayIcon_connect2")
	symaddr(&addr_QSystemTrayIcon_sender, "QSystemTrayIcon_sender")
	symaddr(&addr_QSystemTrayIcon_senderSignalIndex, "QSystemTrayIcon_senderSignalIndex")
	symaddr(&addr_QSystemTrayIcon_receivers, "QSystemTrayIcon_receivers")
	symaddr(&addr_QSystemTrayIcon_isSignalConnected, "QSystemTrayIcon_isSignalConnected")
	symaddr(&addr_QSystemTrayIcon_tr2, "QSystemTrayIcon_tr2")
	symaddr(&addr_QSystemTrayIcon_tr3, "QSystemTrayIcon_tr3")
	symaddr(&addr_QSystemTrayIcon_showMessage3, "QSystemTrayIcon_showMessage3")
	symaddr(&addr_QSystemTrayIcon_showMessage4, "QSystemTrayIcon_showMessage4")
	symaddr(&addr_QSystemTrayIcon_showMessage5, "QSystemTrayIcon_showMessage5")
	symaddr(&addr_QSystemTrayIcon_connect3, "QSystemTrayIcon_connect3")
	symaddr(&addr_QSystemTrayIcon_connect4, "QSystemTrayIcon_connect4")
	symaddr(&addr_QSystemTrayIcon_metaObject_default, "QSystemTrayIcon_metaObject_default")
	symaddr(&addr_QSystemTrayIcon_event_default, "QSystemTrayIcon_event_default")
	symaddr(&addr_QSystemTrayIcon_eventFilter_default, "QSystemTrayIcon_eventFilter_default")
	symaddr(&addr_QSystemTrayIcon_timerEvent_default, "QSystemTrayIcon_timerEvent_default")
	symaddr(&addr_QSystemTrayIcon_childEvent_default, "QSystemTrayIcon_childEvent_default")
	symaddr(&addr_QSystemTrayIcon_customEvent_default, "QSystemTrayIcon_customEvent_default")
	symaddr(&addr_QSystemTrayIcon_connectNotify_default, "QSystemTrayIcon_connectNotify_default")
	symaddr(&addr_QSystemTrayIcon_disconnectNotify_default, "QSystemTrayIcon_disconnectNotify_default")
	reg(&pg_QSystemTrayIcon_connect_activated, "QSystemTrayIcon_connect_activated")
	reg(&pg_QSystemTrayIcon_connect_messageClicked, "QSystemTrayIcon_connect_messageClicked")
	reg(&pg_QSystemTrayIcon_connect_destroyed, "QSystemTrayIcon_connect_destroyed")
	reg(&pg_QSystemTrayIcon_connect_destroyed2, "QSystemTrayIcon_connect_destroyed2")
	reg(&pg_QSystemTrayIcon_connect_objectNameChanged, "QSystemTrayIcon_connect_objectNameChanged")
	reg(&pg_QSystemTrayIcon__disconnect, "QSystemTrayIcon__disconnect")
	reg(&pg_QSystemTrayIcon__set_override, "QSystemTrayIcon__set_override")
	reg(&pg_QSystemTrayIcon_delete, "QSystemTrayIcon_delete")
}

type QSystemTrayIcon struct {
	h uintptr
	*QObject
}

func bindQSystemTrayIcon(h uintptr) *QSystemTrayIcon {
	if h == 0 {
		return nil
	}
	return &QSystemTrayIcon{h: h, QObject: bindQObject(pg_QSystemTrayIcon_base(h))}
}

func newownedQSystemTrayIcon(h uintptr) *QSystemTrayIcon { return bindQSystemTrayIcon(h) }

func (this *QSystemTrayIcon) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSystemTrayIcon) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSystemTrayIcon(h unsafe.Pointer) *QSystemTrayIcon { return bindQSystemTrayIcon(uintptr(h)) }
func UnsafeNewQSystemTrayIconOwned(h unsafe.Pointer) *QSystemTrayIcon { return newownedQSystemTrayIcon(uintptr(h)) }

func (this *QSystemTrayIcon) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQSystemTrayIcon() *QSystemTrayIcon {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_new, uintptr(__goid))
	__o := newownedQSystemTrayIcon(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSystemTrayIcon_virtIdx, QSystemTrayIcon_enable)
	return __o
}

func NewQSystemTrayIcon2(icon *QIcon) *QSystemTrayIcon {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_new2, uintptr(__goid), uintptr(icon.UnsafePointer()))
	__o := newownedQSystemTrayIcon(__r)
	runtime.KeepAlive(icon)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSystemTrayIcon_virtIdx, QSystemTrayIcon_enable)
	return __o
}

func NewQSystemTrayIcon3(parent *QObject) *QSystemTrayIcon {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSystemTrayIcon(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSystemTrayIcon_virtIdx, QSystemTrayIcon_enable)
	return __o
}

func NewQSystemTrayIcon4(icon *QIcon, parent *QObject) *QSystemTrayIcon {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_new4, uintptr(__goid), uintptr(icon.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQSystemTrayIcon(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSystemTrayIcon_virtIdx, QSystemTrayIcon_enable)
	return __o
}

func QSystemTrayIcon_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSystemTrayIcon) SetContextMenu(menu *QMenu) {
	purego.SyscallN(addr_QSystemTrayIcon_setContextMenu, this.h, uintptr(menu.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(menu)
}

func (this *QSystemTrayIcon) ContextMenu() *QMenu {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_contextMenu, this.h)
	runtime.KeepAlive(this)
	return bindQMenu(__r)
}

func (this *QSystemTrayIcon) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QSystemTrayIcon) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QSystemTrayIcon_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QSystemTrayIcon) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSystemTrayIcon) SetToolTip(tip string) {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_setToolTip, this.h, tip_p, uintptr(len(tip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tip)
}

func QSystemTrayIcon_IsSystemTrayAvailable() bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_isSystemTrayAvailable)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QSystemTrayIcon_SupportsMessages() bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_supportsMessages)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QSystemTrayIcon) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QSystemTrayIcon) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSystemTrayIcon) SetVisible(visible bool) {
	purego.SyscallN(addr_QSystemTrayIcon_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QSystemTrayIcon) Show() {
	purego.SyscallN(addr_QSystemTrayIcon_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QSystemTrayIcon) Hide() {
	purego.SyscallN(addr_QSystemTrayIcon_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QSystemTrayIcon) ShowMessage(title string, msg string, icon *QIcon) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var msg_p uintptr
	if len(msg) > 0 {
		msg_p = uintptr(unsafe.Pointer(unsafe.StringData(msg)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_showMessage, this.h, title_p, uintptr(len(title)), msg_p, uintptr(len(msg)), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(icon)
}

func (this *QSystemTrayIcon) ShowMessage2(title string, msg string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var msg_p uintptr
	if len(msg) > 0 {
		msg_p = uintptr(unsafe.Pointer(unsafe.StringData(msg)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_showMessage2, this.h, title_p, uintptr(len(title)), msg_p, uintptr(len(msg)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	runtime.KeepAlive(msg)
}

func (this *QSystemTrayIcon) Activated(reason QSystemTrayIcon__ActivationReason) {
	purego.SyscallN(addr_QSystemTrayIcon_activated, this.h, uintptr(reason))
	runtime.KeepAlive(this)
}

func (this *QSystemTrayIcon) MessageClicked() {
	purego.SyscallN(addr_QSystemTrayIcon_messageClicked, this.h)
	runtime.KeepAlive(this)
}

func QSystemTrayIcon_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSystemTrayIcon_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSystemTrayIcon) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSystemTrayIcon) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSystemTrayIcon) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSystemTrayIcon) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSystemTrayIcon_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSystemTrayIcon_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSystemTrayIcon) ShowMessage3(title string, msg string, icon *QIcon, msecs int) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var msg_p uintptr
	if len(msg) > 0 {
		msg_p = uintptr(unsafe.Pointer(unsafe.StringData(msg)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_showMessage3, this.h, title_p, uintptr(len(title)), msg_p, uintptr(len(msg)), uintptr(icon.UnsafePointer()), uintptr(msecs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	runtime.KeepAlive(msg)
	runtime.KeepAlive(icon)
}

func (this *QSystemTrayIcon) ShowMessage4(title string, msg string, icon QSystemTrayIcon__MessageIcon) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var msg_p uintptr
	if len(msg) > 0 {
		msg_p = uintptr(unsafe.Pointer(unsafe.StringData(msg)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_showMessage4, this.h, title_p, uintptr(len(title)), msg_p, uintptr(len(msg)), uintptr(icon))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	runtime.KeepAlive(msg)
}

func (this *QSystemTrayIcon) ShowMessage5(title string, msg string, icon QSystemTrayIcon__MessageIcon, msecs int) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	var msg_p uintptr
	if len(msg) > 0 {
		msg_p = uintptr(unsafe.Pointer(unsafe.StringData(msg)))
	}
	purego.SyscallN(addr_QSystemTrayIcon_showMessage5, this.h, title_p, uintptr(len(title)), msg_p, uintptr(len(msg)), uintptr(icon), uintptr(msecs))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
	runtime.KeepAlive(msg)
}

func QSystemTrayIcon_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSystemTrayIcon_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSystemTrayIcon) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSystemTrayIcon_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSystemTrayIcon) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSystemTrayIcon) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemTrayIcon_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSystemTrayIcon) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSystemTrayIcon_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSystemTrayIcon) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSystemTrayIcon_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSystemTrayIcon) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSystemTrayIcon_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSystemTrayIcon) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSystemTrayIcon_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSystemTrayIcon) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSystemTrayIcon_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSystemTrayIcon) OnActivated(fn func(QSystemTrayIcon__ActivationReason)) int64 {
	id := registerCallback(func(reason int32) { fn(QSystemTrayIcon__ActivationReason(reason)) })
	pg_QSystemTrayIcon_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSystemTrayIcon) OnMessageClicked(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSystemTrayIcon_connect_messageClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSystemTrayIcon) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSystemTrayIcon_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSystemTrayIcon) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSystemTrayIcon_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSystemTrayIcon) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSystemTrayIcon_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSystemTrayIcon) Disconnect(handle int64) {
	pg_QSystemTrayIcon__disconnect(handle)
	unregisterCallback(handle)
}

func QSystemTrayIcon_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QSystemTrayIcon_enable(h uintptr, idx int32) { pg_QSystemTrayIcon__set_override(h, idx) }

func (this *QSystemTrayIcon) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSystemTrayIcon) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSystemTrayIcon) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSystemTrayIcon) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSystemTrayIcon) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSystemTrayIcon) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSystemTrayIcon) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSystemTrayIcon) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSystemTrayIcon) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSystemTrayIcon_delete(this.h)
	}
	this.UnsafeRelease()
}
