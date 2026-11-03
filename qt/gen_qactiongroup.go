package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QActionGroup_base func(uintptr) uintptr
	addr_QActionGroup_new uintptr
	addr_QActionGroup_tr uintptr
	addr_QActionGroup_addAction uintptr
	addr_QActionGroup_addAction2 uintptr
	addr_QActionGroup_addAction3 uintptr
	addr_QActionGroup_removeAction uintptr
	addr_QActionGroup_actions uintptr
	addr_QActionGroup_checkedAction uintptr
	addr_QActionGroup_isExclusive uintptr
	addr_QActionGroup_isEnabled uintptr
	addr_QActionGroup_isVisible uintptr
	addr_QActionGroup_exclusionPolicy uintptr
	addr_QActionGroup_setEnabled uintptr
	addr_QActionGroup_setDisabled uintptr
	addr_QActionGroup_setVisible uintptr
	addr_QActionGroup_setExclusive uintptr
	addr_QActionGroup_setExclusionPolicy uintptr
	addr_QActionGroup_triggered uintptr
	addr_QActionGroup_hovered uintptr
	addr_QActionGroup_connect uintptr
	addr_QActionGroup_connect2 uintptr
	addr_QActionGroup_sender uintptr
	addr_QActionGroup_senderSignalIndex uintptr
	addr_QActionGroup_receivers uintptr
	addr_QActionGroup_isSignalConnected uintptr
	addr_QActionGroup_tr2 uintptr
	addr_QActionGroup_tr3 uintptr
	addr_QActionGroup_connect3 uintptr
	addr_QActionGroup_connect4 uintptr
	addr_QActionGroup_metaObject_default uintptr
	addr_QActionGroup_event_default uintptr
	addr_QActionGroup_eventFilter_default uintptr
	addr_QActionGroup_timerEvent_default uintptr
	addr_QActionGroup_childEvent_default uintptr
	addr_QActionGroup_customEvent_default uintptr
	addr_QActionGroup_connectNotify_default uintptr
	addr_QActionGroup_disconnectNotify_default uintptr
	pg_QActionGroup_connect_triggered func(uintptr, int64)
	pg_QActionGroup_connect_hovered func(uintptr, int64)
	pg_QActionGroup_connect_destroyed func(uintptr, int64)
	pg_QActionGroup_connect_destroyed2 func(uintptr, int64)
	pg_QActionGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QActionGroup__disconnect func(int64)
	pg_QActionGroup__set_override func(uintptr, int32)
	pg_QActionGroup_delete func(uintptr)
)

func init() {
	reg(&pg_QActionGroup_base, "QActionGroup_base")
	symaddr(&addr_QActionGroup_new, "QActionGroup_new")
	symaddr(&addr_QActionGroup_tr, "QActionGroup_tr")
	symaddr(&addr_QActionGroup_addAction, "QActionGroup_addAction")
	symaddr(&addr_QActionGroup_addAction2, "QActionGroup_addAction2")
	symaddr(&addr_QActionGroup_addAction3, "QActionGroup_addAction3")
	symaddr(&addr_QActionGroup_removeAction, "QActionGroup_removeAction")
	symaddr(&addr_QActionGroup_actions, "QActionGroup_actions")
	symaddr(&addr_QActionGroup_checkedAction, "QActionGroup_checkedAction")
	symaddr(&addr_QActionGroup_isExclusive, "QActionGroup_isExclusive")
	symaddr(&addr_QActionGroup_isEnabled, "QActionGroup_isEnabled")
	symaddr(&addr_QActionGroup_isVisible, "QActionGroup_isVisible")
	symaddr(&addr_QActionGroup_exclusionPolicy, "QActionGroup_exclusionPolicy")
	symaddr(&addr_QActionGroup_setEnabled, "QActionGroup_setEnabled")
	symaddr(&addr_QActionGroup_setDisabled, "QActionGroup_setDisabled")
	symaddr(&addr_QActionGroup_setVisible, "QActionGroup_setVisible")
	symaddr(&addr_QActionGroup_setExclusive, "QActionGroup_setExclusive")
	symaddr(&addr_QActionGroup_setExclusionPolicy, "QActionGroup_setExclusionPolicy")
	symaddr(&addr_QActionGroup_triggered, "QActionGroup_triggered")
	symaddr(&addr_QActionGroup_hovered, "QActionGroup_hovered")
	symaddr(&addr_QActionGroup_connect, "QActionGroup_connect")
	symaddr(&addr_QActionGroup_connect2, "QActionGroup_connect2")
	symaddr(&addr_QActionGroup_sender, "QActionGroup_sender")
	symaddr(&addr_QActionGroup_senderSignalIndex, "QActionGroup_senderSignalIndex")
	symaddr(&addr_QActionGroup_receivers, "QActionGroup_receivers")
	symaddr(&addr_QActionGroup_isSignalConnected, "QActionGroup_isSignalConnected")
	symaddr(&addr_QActionGroup_tr2, "QActionGroup_tr2")
	symaddr(&addr_QActionGroup_tr3, "QActionGroup_tr3")
	symaddr(&addr_QActionGroup_connect3, "QActionGroup_connect3")
	symaddr(&addr_QActionGroup_connect4, "QActionGroup_connect4")
	symaddr(&addr_QActionGroup_metaObject_default, "QActionGroup_metaObject_default")
	symaddr(&addr_QActionGroup_event_default, "QActionGroup_event_default")
	symaddr(&addr_QActionGroup_eventFilter_default, "QActionGroup_eventFilter_default")
	symaddr(&addr_QActionGroup_timerEvent_default, "QActionGroup_timerEvent_default")
	symaddr(&addr_QActionGroup_childEvent_default, "QActionGroup_childEvent_default")
	symaddr(&addr_QActionGroup_customEvent_default, "QActionGroup_customEvent_default")
	symaddr(&addr_QActionGroup_connectNotify_default, "QActionGroup_connectNotify_default")
	symaddr(&addr_QActionGroup_disconnectNotify_default, "QActionGroup_disconnectNotify_default")
	reg(&pg_QActionGroup_connect_triggered, "QActionGroup_connect_triggered")
	reg(&pg_QActionGroup_connect_hovered, "QActionGroup_connect_hovered")
	reg(&pg_QActionGroup_connect_destroyed, "QActionGroup_connect_destroyed")
	reg(&pg_QActionGroup_connect_destroyed2, "QActionGroup_connect_destroyed2")
	reg(&pg_QActionGroup_connect_objectNameChanged, "QActionGroup_connect_objectNameChanged")
	reg(&pg_QActionGroup__disconnect, "QActionGroup__disconnect")
	reg(&pg_QActionGroup__set_override, "QActionGroup__set_override")
	reg(&pg_QActionGroup_delete, "QActionGroup_delete")
}

type QActionGroup struct {
	h uintptr
	*QObject
}

func bindQActionGroup(h uintptr) *QActionGroup {
	if h == 0 {
		return nil
	}
	return &QActionGroup{h: h, QObject: bindQObject(pg_QActionGroup_base(h))}
}

func newownedQActionGroup(h uintptr) *QActionGroup { return bindQActionGroup(h) }

func (this *QActionGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QActionGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQActionGroup(h unsafe.Pointer) *QActionGroup { return bindQActionGroup(uintptr(h)) }
func UnsafeNewQActionGroupOwned(h unsafe.Pointer) *QActionGroup { return newownedQActionGroup(uintptr(h)) }

func (this *QActionGroup) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQActionGroup(parent *QObject) *QActionGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QActionGroup_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQActionGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QActionGroup_virtIdx, QActionGroup_enable)
	return __o
}

func QActionGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QActionGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QActionGroup) AddAction(a *QAction) *QAction {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_addAction, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return bindQAction(__r)
}

func (this *QActionGroup) AddActionWithText(text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QActionGroup_addAction2, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QActionGroup) AddAction2(icon *QIcon, text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QActionGroup_addAction3, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QActionGroup) RemoveAction(a *QAction) {
	purego.SyscallN(addr_QActionGroup_removeAction, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QActionGroup) Actions() []*QAction {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QActionGroup_actions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAction, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAction(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QActionGroup) CheckedAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_checkedAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QActionGroup) IsExclusive() bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_isExclusive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QActionGroup) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QActionGroup) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QActionGroup) ExclusionPolicy() QActionGroup__ExclusionPolicy {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_exclusionPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QActionGroup__ExclusionPolicy(__r)
}

func (this *QActionGroup) SetEnabled(arg1 bool) {
	purego.SyscallN(addr_QActionGroup_setEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QActionGroup) SetDisabled(b bool) {
	purego.SyscallN(addr_QActionGroup_setDisabled, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QActionGroup) SetVisible(arg1 bool) {
	purego.SyscallN(addr_QActionGroup_setVisible, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QActionGroup) SetExclusive(arg1 bool) {
	purego.SyscallN(addr_QActionGroup_setExclusive, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QActionGroup) SetExclusionPolicy(policy QActionGroup__ExclusionPolicy) {
	purego.SyscallN(addr_QActionGroup_setExclusionPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QActionGroup) Triggered(arg1 *QAction) {
	purego.SyscallN(addr_QActionGroup_triggered, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QActionGroup) Hovered(arg1 *QAction) {
	purego.SyscallN(addr_QActionGroup_hovered, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QActionGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QActionGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QActionGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QActionGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QActionGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QActionGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QActionGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QActionGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QActionGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QActionGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QActionGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QActionGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QActionGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QActionGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QActionGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QActionGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QActionGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QActionGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QActionGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QActionGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QActionGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QActionGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QActionGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QActionGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QActionGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QActionGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QActionGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QActionGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QActionGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QActionGroup) OnTriggered(fn func(*QAction)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(arg1))) })
	pg_QActionGroup_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QActionGroup) OnHovered(fn func(*QAction)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(arg1))) })
	pg_QActionGroup_connect_hovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QActionGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QActionGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QActionGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QActionGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QActionGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QActionGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QActionGroup) Disconnect(handle int64) {
	pg_QActionGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QActionGroup_virtIdx(slot int64) int32 {
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

func QActionGroup_enable(h uintptr, idx int32) { pg_QActionGroup__set_override(h, idx) }

func (this *QActionGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QActionGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QActionGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QActionGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QActionGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QActionGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QActionGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QActionGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QActionGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QActionGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
