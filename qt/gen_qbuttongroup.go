package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QButtonGroup_base func(uintptr) uintptr
	addr_QButtonGroup_new uintptr
	addr_QButtonGroup_new2 uintptr
	addr_QButtonGroup_tr uintptr
	addr_QButtonGroup_setExclusive uintptr
	addr_QButtonGroup_exclusive uintptr
	addr_QButtonGroup_addButton uintptr
	addr_QButtonGroup_removeButton uintptr
	addr_QButtonGroup_buttons uintptr
	addr_QButtonGroup_checkedButton uintptr
	addr_QButtonGroup_button uintptr
	addr_QButtonGroup_setId uintptr
	addr_QButtonGroup_id uintptr
	addr_QButtonGroup_checkedId uintptr
	addr_QButtonGroup_buttonClicked uintptr
	addr_QButtonGroup_buttonPressed uintptr
	addr_QButtonGroup_buttonReleased uintptr
	addr_QButtonGroup_buttonToggled uintptr
	addr_QButtonGroup_idClicked uintptr
	addr_QButtonGroup_idPressed uintptr
	addr_QButtonGroup_idReleased uintptr
	addr_QButtonGroup_idToggled uintptr
	addr_QButtonGroup_connect uintptr
	addr_QButtonGroup_connect2 uintptr
	addr_QButtonGroup_sender uintptr
	addr_QButtonGroup_senderSignalIndex uintptr
	addr_QButtonGroup_receivers uintptr
	addr_QButtonGroup_isSignalConnected uintptr
	addr_QButtonGroup_tr2 uintptr
	addr_QButtonGroup_tr3 uintptr
	addr_QButtonGroup_addButton2 uintptr
	addr_QButtonGroup_connect3 uintptr
	addr_QButtonGroup_connect4 uintptr
	addr_QButtonGroup_metaObject_default uintptr
	addr_QButtonGroup_event_default uintptr
	addr_QButtonGroup_eventFilter_default uintptr
	addr_QButtonGroup_timerEvent_default uintptr
	addr_QButtonGroup_childEvent_default uintptr
	addr_QButtonGroup_customEvent_default uintptr
	addr_QButtonGroup_connectNotify_default uintptr
	addr_QButtonGroup_disconnectNotify_default uintptr
	pg_QButtonGroup_connect_buttonClicked func(uintptr, int64)
	pg_QButtonGroup_connect_buttonPressed func(uintptr, int64)
	pg_QButtonGroup_connect_buttonReleased func(uintptr, int64)
	pg_QButtonGroup_connect_buttonToggled func(uintptr, int64)
	pg_QButtonGroup_connect_idClicked func(uintptr, int64)
	pg_QButtonGroup_connect_idPressed func(uintptr, int64)
	pg_QButtonGroup_connect_idReleased func(uintptr, int64)
	pg_QButtonGroup_connect_idToggled func(uintptr, int64)
	pg_QButtonGroup_connect_destroyed func(uintptr, int64)
	pg_QButtonGroup_connect_destroyed2 func(uintptr, int64)
	pg_QButtonGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QButtonGroup__disconnect func(int64)
	pg_QButtonGroup__set_override func(uintptr, int32)
	pg_QButtonGroup_delete func(uintptr)
)

func init() {
	reg(&pg_QButtonGroup_base, "QButtonGroup_base")
	symaddr(&addr_QButtonGroup_new, "QButtonGroup_new")
	symaddr(&addr_QButtonGroup_new2, "QButtonGroup_new2")
	symaddr(&addr_QButtonGroup_tr, "QButtonGroup_tr")
	symaddr(&addr_QButtonGroup_setExclusive, "QButtonGroup_setExclusive")
	symaddr(&addr_QButtonGroup_exclusive, "QButtonGroup_exclusive")
	symaddr(&addr_QButtonGroup_addButton, "QButtonGroup_addButton")
	symaddr(&addr_QButtonGroup_removeButton, "QButtonGroup_removeButton")
	symaddr(&addr_QButtonGroup_buttons, "QButtonGroup_buttons")
	symaddr(&addr_QButtonGroup_checkedButton, "QButtonGroup_checkedButton")
	symaddr(&addr_QButtonGroup_button, "QButtonGroup_button")
	symaddr(&addr_QButtonGroup_setId, "QButtonGroup_setId")
	symaddr(&addr_QButtonGroup_id, "QButtonGroup_id")
	symaddr(&addr_QButtonGroup_checkedId, "QButtonGroup_checkedId")
	symaddr(&addr_QButtonGroup_buttonClicked, "QButtonGroup_buttonClicked")
	symaddr(&addr_QButtonGroup_buttonPressed, "QButtonGroup_buttonPressed")
	symaddr(&addr_QButtonGroup_buttonReleased, "QButtonGroup_buttonReleased")
	symaddr(&addr_QButtonGroup_buttonToggled, "QButtonGroup_buttonToggled")
	symaddr(&addr_QButtonGroup_idClicked, "QButtonGroup_idClicked")
	symaddr(&addr_QButtonGroup_idPressed, "QButtonGroup_idPressed")
	symaddr(&addr_QButtonGroup_idReleased, "QButtonGroup_idReleased")
	symaddr(&addr_QButtonGroup_idToggled, "QButtonGroup_idToggled")
	symaddr(&addr_QButtonGroup_connect, "QButtonGroup_connect")
	symaddr(&addr_QButtonGroup_connect2, "QButtonGroup_connect2")
	symaddr(&addr_QButtonGroup_sender, "QButtonGroup_sender")
	symaddr(&addr_QButtonGroup_senderSignalIndex, "QButtonGroup_senderSignalIndex")
	symaddr(&addr_QButtonGroup_receivers, "QButtonGroup_receivers")
	symaddr(&addr_QButtonGroup_isSignalConnected, "QButtonGroup_isSignalConnected")
	symaddr(&addr_QButtonGroup_tr2, "QButtonGroup_tr2")
	symaddr(&addr_QButtonGroup_tr3, "QButtonGroup_tr3")
	symaddr(&addr_QButtonGroup_addButton2, "QButtonGroup_addButton2")
	symaddr(&addr_QButtonGroup_connect3, "QButtonGroup_connect3")
	symaddr(&addr_QButtonGroup_connect4, "QButtonGroup_connect4")
	symaddr(&addr_QButtonGroup_metaObject_default, "QButtonGroup_metaObject_default")
	symaddr(&addr_QButtonGroup_event_default, "QButtonGroup_event_default")
	symaddr(&addr_QButtonGroup_eventFilter_default, "QButtonGroup_eventFilter_default")
	symaddr(&addr_QButtonGroup_timerEvent_default, "QButtonGroup_timerEvent_default")
	symaddr(&addr_QButtonGroup_childEvent_default, "QButtonGroup_childEvent_default")
	symaddr(&addr_QButtonGroup_customEvent_default, "QButtonGroup_customEvent_default")
	symaddr(&addr_QButtonGroup_connectNotify_default, "QButtonGroup_connectNotify_default")
	symaddr(&addr_QButtonGroup_disconnectNotify_default, "QButtonGroup_disconnectNotify_default")
	reg(&pg_QButtonGroup_connect_buttonClicked, "QButtonGroup_connect_buttonClicked")
	reg(&pg_QButtonGroup_connect_buttonPressed, "QButtonGroup_connect_buttonPressed")
	reg(&pg_QButtonGroup_connect_buttonReleased, "QButtonGroup_connect_buttonReleased")
	reg(&pg_QButtonGroup_connect_buttonToggled, "QButtonGroup_connect_buttonToggled")
	reg(&pg_QButtonGroup_connect_idClicked, "QButtonGroup_connect_idClicked")
	reg(&pg_QButtonGroup_connect_idPressed, "QButtonGroup_connect_idPressed")
	reg(&pg_QButtonGroup_connect_idReleased, "QButtonGroup_connect_idReleased")
	reg(&pg_QButtonGroup_connect_idToggled, "QButtonGroup_connect_idToggled")
	reg(&pg_QButtonGroup_connect_destroyed, "QButtonGroup_connect_destroyed")
	reg(&pg_QButtonGroup_connect_destroyed2, "QButtonGroup_connect_destroyed2")
	reg(&pg_QButtonGroup_connect_objectNameChanged, "QButtonGroup_connect_objectNameChanged")
	reg(&pg_QButtonGroup__disconnect, "QButtonGroup__disconnect")
	reg(&pg_QButtonGroup__set_override, "QButtonGroup__set_override")
	reg(&pg_QButtonGroup_delete, "QButtonGroup_delete")
}

type QButtonGroup struct {
	h uintptr
	*QObject
}

func bindQButtonGroup(h uintptr) *QButtonGroup {
	if h == 0 {
		return nil
	}
	return &QButtonGroup{h: h, QObject: bindQObject(pg_QButtonGroup_base(h))}
}

func newownedQButtonGroup(h uintptr) *QButtonGroup { return bindQButtonGroup(h) }

func (this *QButtonGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QButtonGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQButtonGroup(h unsafe.Pointer) *QButtonGroup { return bindQButtonGroup(uintptr(h)) }
func UnsafeNewQButtonGroupOwned(h unsafe.Pointer) *QButtonGroup { return newownedQButtonGroup(uintptr(h)) }

func (this *QButtonGroup) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQButtonGroup() *QButtonGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_new, uintptr(__goid))
	__o := newownedQButtonGroup(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QButtonGroup_virtIdx, QButtonGroup_enable)
	return __o
}

func NewQButtonGroup2(parent *QObject) *QButtonGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQButtonGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QButtonGroup_virtIdx, QButtonGroup_enable)
	return __o
}

func QButtonGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QButtonGroup) SetExclusive(arg1 bool) {
	purego.SyscallN(addr_QButtonGroup_setExclusive, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QButtonGroup) Exclusive() bool {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_exclusive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QButtonGroup) AddButton(arg1 *QAbstractButton) {
	purego.SyscallN(addr_QButtonGroup_addButton, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) RemoveButton(arg1 *QAbstractButton) {
	purego.SyscallN(addr_QButtonGroup_removeButton, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) Buttons() []*QAbstractButton {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_buttons, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAbstractButton, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAbstractButton(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QButtonGroup) CheckedButton() *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_checkedButton, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QButtonGroup) Button(id int) *QAbstractButton {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_button, this.h, uintptr(id))
	runtime.KeepAlive(this)
	return bindQAbstractButton(__r)
}

func (this *QButtonGroup) SetId(button *QAbstractButton, id int) {
	purego.SyscallN(addr_QButtonGroup_setId, this.h, uintptr(button.UnsafePointer()), uintptr(id))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
}

func (this *QButtonGroup) Id(button *QAbstractButton) int {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_id, this.h, uintptr(button.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(button)
	return int(__r)
}

func (this *QButtonGroup) CheckedId() int {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_checkedId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QButtonGroup) ButtonClicked(arg1 *QAbstractButton) {
	purego.SyscallN(addr_QButtonGroup_buttonClicked, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) ButtonPressed(arg1 *QAbstractButton) {
	purego.SyscallN(addr_QButtonGroup_buttonPressed, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) ButtonReleased(arg1 *QAbstractButton) {
	purego.SyscallN(addr_QButtonGroup_buttonReleased, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) ButtonToggled(arg1 *QAbstractButton, arg2 bool) {
	purego.SyscallN(addr_QButtonGroup_buttonToggled, this.h, uintptr(arg1.UnsafePointer()), uintptr(b2i(arg2)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QButtonGroup) IdClicked(arg1 int) {
	purego.SyscallN(addr_QButtonGroup_idClicked, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QButtonGroup) IdPressed(arg1 int) {
	purego.SyscallN(addr_QButtonGroup_idPressed, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QButtonGroup) IdReleased(arg1 int) {
	purego.SyscallN(addr_QButtonGroup_idReleased, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QButtonGroup) IdToggled(arg1 int, arg2 bool) {
	purego.SyscallN(addr_QButtonGroup_idToggled, this.h, uintptr(arg1), uintptr(b2i(arg2)))
	runtime.KeepAlive(this)
}

func QButtonGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QButtonGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QButtonGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QButtonGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QButtonGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QButtonGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QButtonGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QButtonGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QButtonGroup) AddButton2(arg1 *QAbstractButton, id int) {
	purego.SyscallN(addr_QButtonGroup_addButton2, this.h, uintptr(arg1.UnsafePointer()), uintptr(id))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QButtonGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QButtonGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QButtonGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QButtonGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QButtonGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QButtonGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QButtonGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QButtonGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QButtonGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QButtonGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QButtonGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QButtonGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QButtonGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QButtonGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QButtonGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QButtonGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QButtonGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QButtonGroup) OnButtonClicked(fn func(*QAbstractButton)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(arg1))) })
	pg_QButtonGroup_connect_buttonClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnButtonPressed(fn func(*QAbstractButton)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(arg1))) })
	pg_QButtonGroup_connect_buttonPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnButtonReleased(fn func(*QAbstractButton)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(arg1))) })
	pg_QButtonGroup_connect_buttonReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnButtonToggled(fn func(*QAbstractButton, bool)) int64 {
	id := registerCallback(func(arg1 uintptr, arg2 int32) { fn(UnsafeNewQAbstractButton(unsafe.Pointer(arg1)), arg2 != 0) })
	pg_QButtonGroup_connect_buttonToggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnIdClicked(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QButtonGroup_connect_idClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnIdPressed(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QButtonGroup_connect_idPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnIdReleased(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QButtonGroup_connect_idReleased(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnIdToggled(fn func(int, bool)) int64 {
	id := registerCallback(func(arg1 int32, arg2 int32) { fn(int(arg1), arg2 != 0) })
	pg_QButtonGroup_connect_idToggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QButtonGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QButtonGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QButtonGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QButtonGroup) Disconnect(handle int64) {
	pg_QButtonGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QButtonGroup_virtIdx(slot int64) int32 {
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

func QButtonGroup_enable(h uintptr, idx int32) { pg_QButtonGroup__set_override(h, idx) }

func (this *QButtonGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QButtonGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QButtonGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QButtonGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QButtonGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QButtonGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QButtonGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QButtonGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QButtonGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QButtonGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
