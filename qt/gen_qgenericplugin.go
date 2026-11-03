package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGenericPlugin_base func(uintptr) uintptr
	addr_QGenericPlugin_tr uintptr
	addr_QGenericPlugin_create uintptr
	addr_QGenericPlugin_connect uintptr
	addr_QGenericPlugin_connect2 uintptr
	addr_QGenericPlugin_sender uintptr
	addr_QGenericPlugin_senderSignalIndex uintptr
	addr_QGenericPlugin_receivers uintptr
	addr_QGenericPlugin_isSignalConnected uintptr
	addr_QGenericPlugin_tr2 uintptr
	addr_QGenericPlugin_tr3 uintptr
	addr_QGenericPlugin_connect3 uintptr
	addr_QGenericPlugin_connect4 uintptr
	addr_QGenericPlugin_metaObject_default uintptr
	addr_QGenericPlugin_event_default uintptr
	addr_QGenericPlugin_eventFilter_default uintptr
	addr_QGenericPlugin_timerEvent_default uintptr
	addr_QGenericPlugin_childEvent_default uintptr
	addr_QGenericPlugin_customEvent_default uintptr
	addr_QGenericPlugin_connectNotify_default uintptr
	addr_QGenericPlugin_disconnectNotify_default uintptr
	pg_QGenericPlugin_connect_destroyed func(uintptr, int64)
	pg_QGenericPlugin_connect_destroyed2 func(uintptr, int64)
	pg_QGenericPlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QGenericPlugin__disconnect func(int64)
	pg_QGenericPlugin__set_override func(uintptr, int32)
	pg_QGenericPlugin__Go_new func(int64, uintptr) uintptr
	pg_QGenericPlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QGenericPlugin_base, "QGenericPlugin_base")
	symaddr(&addr_QGenericPlugin_tr, "QGenericPlugin_tr")
	symaddr(&addr_QGenericPlugin_create, "QGenericPlugin_create")
	symaddr(&addr_QGenericPlugin_connect, "QGenericPlugin_connect")
	symaddr(&addr_QGenericPlugin_connect2, "QGenericPlugin_connect2")
	symaddr(&addr_QGenericPlugin_sender, "QGenericPlugin_sender")
	symaddr(&addr_QGenericPlugin_senderSignalIndex, "QGenericPlugin_senderSignalIndex")
	symaddr(&addr_QGenericPlugin_receivers, "QGenericPlugin_receivers")
	symaddr(&addr_QGenericPlugin_isSignalConnected, "QGenericPlugin_isSignalConnected")
	symaddr(&addr_QGenericPlugin_tr2, "QGenericPlugin_tr2")
	symaddr(&addr_QGenericPlugin_tr3, "QGenericPlugin_tr3")
	symaddr(&addr_QGenericPlugin_connect3, "QGenericPlugin_connect3")
	symaddr(&addr_QGenericPlugin_connect4, "QGenericPlugin_connect4")
	symaddr(&addr_QGenericPlugin_metaObject_default, "QGenericPlugin_metaObject_default")
	symaddr(&addr_QGenericPlugin_event_default, "QGenericPlugin_event_default")
	symaddr(&addr_QGenericPlugin_eventFilter_default, "QGenericPlugin_eventFilter_default")
	symaddr(&addr_QGenericPlugin_timerEvent_default, "QGenericPlugin_timerEvent_default")
	symaddr(&addr_QGenericPlugin_childEvent_default, "QGenericPlugin_childEvent_default")
	symaddr(&addr_QGenericPlugin_customEvent_default, "QGenericPlugin_customEvent_default")
	symaddr(&addr_QGenericPlugin_connectNotify_default, "QGenericPlugin_connectNotify_default")
	symaddr(&addr_QGenericPlugin_disconnectNotify_default, "QGenericPlugin_disconnectNotify_default")
	reg(&pg_QGenericPlugin_connect_destroyed, "QGenericPlugin_connect_destroyed")
	reg(&pg_QGenericPlugin_connect_destroyed2, "QGenericPlugin_connect_destroyed2")
	reg(&pg_QGenericPlugin_connect_objectNameChanged, "QGenericPlugin_connect_objectNameChanged")
	reg(&pg_QGenericPlugin__disconnect, "QGenericPlugin__disconnect")
	reg(&pg_QGenericPlugin__set_override, "QGenericPlugin__set_override")
	reg(&pg_QGenericPlugin__Go_new, "QGenericPlugin__Go_new")
	reg(&pg_QGenericPlugin_delete, "QGenericPlugin_delete")
}

type QGenericPlugin struct {
	h uintptr
	*QObject
}

func bindQGenericPlugin(h uintptr) *QGenericPlugin {
	if h == 0 {
		return nil
	}
	return &QGenericPlugin{h: h, QObject: bindQObject(pg_QGenericPlugin_base(h))}
}

func newownedQGenericPlugin(h uintptr) *QGenericPlugin { return bindQGenericPlugin(h) }

func (this *QGenericPlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGenericPlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGenericPlugin(h unsafe.Pointer) *QGenericPlugin { return bindQGenericPlugin(uintptr(h)) }
func UnsafeNewQGenericPluginOwned(h unsafe.Pointer) *QGenericPlugin { return newownedQGenericPlugin(uintptr(h)) }

func (this *QGenericPlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QGenericPlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGenericPlugin) Create(name string, spec string) *QObject {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var spec_p uintptr
	if len(spec) > 0 {
		spec_p = uintptr(unsafe.Pointer(unsafe.StringData(spec)))
	}
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_create, this.h, name_p, uintptr(len(name)), spec_p, uintptr(len(spec)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(spec)
	return bindQObject(__r)
}

func QGenericPlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGenericPlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGenericPlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGenericPlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGenericPlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGenericPlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGenericPlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGenericPlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGenericPlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGenericPlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGenericPlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGenericPlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGenericPlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGenericPlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGenericPlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGenericPlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGenericPlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGenericPlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGenericPlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGenericPlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGenericPlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGenericPlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGenericPlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGenericPlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGenericPlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGenericPlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGenericPlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGenericPlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGenericPlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGenericPlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGenericPlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGenericPlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGenericPlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGenericPlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGenericPlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGenericPlugin) Disconnect(handle int64) {
	pg_QGenericPlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QGenericPlugin_virtIdx(slot int64) int32 {
	switch slot {
	case -626457835795418115:
		return 0
	case 1607084412591819503:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
		return 3
	case -5409927581812374498:
		return 4
	case -1185315367526857833:
		return 5
	case 3769868229123868474:
		return 6
	case 7379029982170503008:
		return 7
	case -568369772982215962:
		return 8
	}
	return -1
}

func QGenericPlugin_enable(h uintptr, idx int32) { pg_QGenericPlugin__set_override(h, idx) }

func NewQGenericPlugin() *QGenericPlugin {
	id := nextVirtID()
	o := newownedQGenericPlugin(pg_QGenericPlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGenericPlugin_virtIdx, QGenericPlugin_enable)
	return o
}

func NewQGenericPlugin2(parent *QObject) *QGenericPlugin {
	id := nextVirtID()
	o := newownedQGenericPlugin(pg_QGenericPlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGenericPlugin_virtIdx, QGenericPlugin_enable)
	return o
}

func (this *QGenericPlugin) OnCreate(fn func(string, string) *QObject) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(name uintptr, name_len int64, spec uintptr, spec_len int64) uintptr { return uintptr(fn(pgStrNoFree(name, name_len), pgStrNoFree(spec, spec_len)).UnsafePointer()) }) // slot "create"
}

func (this *QGenericPlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGenericPlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGenericPlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGenericPlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGenericPlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGenericPlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGenericPlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGenericPlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGenericPlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGenericPlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
