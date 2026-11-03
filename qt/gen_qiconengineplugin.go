package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QIconEnginePlugin_base func(uintptr) uintptr
	addr_QIconEnginePlugin_tr uintptr
	addr_QIconEnginePlugin_create uintptr
	addr_QIconEnginePlugin_connect uintptr
	addr_QIconEnginePlugin_connect2 uintptr
	addr_QIconEnginePlugin_sender uintptr
	addr_QIconEnginePlugin_senderSignalIndex uintptr
	addr_QIconEnginePlugin_receivers uintptr
	addr_QIconEnginePlugin_isSignalConnected uintptr
	addr_QIconEnginePlugin_tr2 uintptr
	addr_QIconEnginePlugin_tr3 uintptr
	addr_QIconEnginePlugin_connect3 uintptr
	addr_QIconEnginePlugin_connect4 uintptr
	addr_QIconEnginePlugin_metaObject_default uintptr
	addr_QIconEnginePlugin_event_default uintptr
	addr_QIconEnginePlugin_eventFilter_default uintptr
	addr_QIconEnginePlugin_timerEvent_default uintptr
	addr_QIconEnginePlugin_childEvent_default uintptr
	addr_QIconEnginePlugin_customEvent_default uintptr
	addr_QIconEnginePlugin_connectNotify_default uintptr
	addr_QIconEnginePlugin_disconnectNotify_default uintptr
	pg_QIconEnginePlugin_connect_destroyed func(uintptr, int64)
	pg_QIconEnginePlugin_connect_destroyed2 func(uintptr, int64)
	pg_QIconEnginePlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QIconEnginePlugin__disconnect func(int64)
	pg_QIconEnginePlugin__set_override func(uintptr, int32)
	pg_QIconEnginePlugin__Go_new func(int64, uintptr) uintptr
	pg_QIconEnginePlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QIconEnginePlugin_base, "QIconEnginePlugin_base")
	symaddr(&addr_QIconEnginePlugin_tr, "QIconEnginePlugin_tr")
	symaddr(&addr_QIconEnginePlugin_create, "QIconEnginePlugin_create")
	symaddr(&addr_QIconEnginePlugin_connect, "QIconEnginePlugin_connect")
	symaddr(&addr_QIconEnginePlugin_connect2, "QIconEnginePlugin_connect2")
	symaddr(&addr_QIconEnginePlugin_sender, "QIconEnginePlugin_sender")
	symaddr(&addr_QIconEnginePlugin_senderSignalIndex, "QIconEnginePlugin_senderSignalIndex")
	symaddr(&addr_QIconEnginePlugin_receivers, "QIconEnginePlugin_receivers")
	symaddr(&addr_QIconEnginePlugin_isSignalConnected, "QIconEnginePlugin_isSignalConnected")
	symaddr(&addr_QIconEnginePlugin_tr2, "QIconEnginePlugin_tr2")
	symaddr(&addr_QIconEnginePlugin_tr3, "QIconEnginePlugin_tr3")
	symaddr(&addr_QIconEnginePlugin_connect3, "QIconEnginePlugin_connect3")
	symaddr(&addr_QIconEnginePlugin_connect4, "QIconEnginePlugin_connect4")
	symaddr(&addr_QIconEnginePlugin_metaObject_default, "QIconEnginePlugin_metaObject_default")
	symaddr(&addr_QIconEnginePlugin_event_default, "QIconEnginePlugin_event_default")
	symaddr(&addr_QIconEnginePlugin_eventFilter_default, "QIconEnginePlugin_eventFilter_default")
	symaddr(&addr_QIconEnginePlugin_timerEvent_default, "QIconEnginePlugin_timerEvent_default")
	symaddr(&addr_QIconEnginePlugin_childEvent_default, "QIconEnginePlugin_childEvent_default")
	symaddr(&addr_QIconEnginePlugin_customEvent_default, "QIconEnginePlugin_customEvent_default")
	symaddr(&addr_QIconEnginePlugin_connectNotify_default, "QIconEnginePlugin_connectNotify_default")
	symaddr(&addr_QIconEnginePlugin_disconnectNotify_default, "QIconEnginePlugin_disconnectNotify_default")
	reg(&pg_QIconEnginePlugin_connect_destroyed, "QIconEnginePlugin_connect_destroyed")
	reg(&pg_QIconEnginePlugin_connect_destroyed2, "QIconEnginePlugin_connect_destroyed2")
	reg(&pg_QIconEnginePlugin_connect_objectNameChanged, "QIconEnginePlugin_connect_objectNameChanged")
	reg(&pg_QIconEnginePlugin__disconnect, "QIconEnginePlugin__disconnect")
	reg(&pg_QIconEnginePlugin__set_override, "QIconEnginePlugin__set_override")
	reg(&pg_QIconEnginePlugin__Go_new, "QIconEnginePlugin__Go_new")
	reg(&pg_QIconEnginePlugin_delete, "QIconEnginePlugin_delete")
}

type QIconEnginePlugin struct {
	h uintptr
	*QObject
}

func bindQIconEnginePlugin(h uintptr) *QIconEnginePlugin {
	if h == 0 {
		return nil
	}
	return &QIconEnginePlugin{h: h, QObject: bindQObject(pg_QIconEnginePlugin_base(h))}
}

func newownedQIconEnginePlugin(h uintptr) *QIconEnginePlugin { return bindQIconEnginePlugin(h) }

func (this *QIconEnginePlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIconEnginePlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIconEnginePlugin(h unsafe.Pointer) *QIconEnginePlugin { return bindQIconEnginePlugin(uintptr(h)) }
func UnsafeNewQIconEnginePluginOwned(h unsafe.Pointer) *QIconEnginePlugin { return newownedQIconEnginePlugin(uintptr(h)) }

func (this *QIconEnginePlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QIconEnginePlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QIconEnginePlugin) Create(filename string) *QIconEngine {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_create, this.h, filename_p, uintptr(len(filename)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
	return bindQIconEngine(__r)
}

func QIconEnginePlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIconEnginePlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QIconEnginePlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QIconEnginePlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QIconEnginePlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIconEnginePlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QIconEnginePlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QIconEnginePlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEnginePlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QIconEnginePlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QIconEnginePlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QIconEnginePlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIconEnginePlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QIconEnginePlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QIconEnginePlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QIconEnginePlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEnginePlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QIconEnginePlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEnginePlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QIconEnginePlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QIconEnginePlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIconEnginePlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QIconEnginePlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIconEnginePlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QIconEnginePlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIconEnginePlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QIconEnginePlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QIconEnginePlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QIconEnginePlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QIconEnginePlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIconEnginePlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIconEnginePlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QIconEnginePlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIconEnginePlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QIconEnginePlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIconEnginePlugin) Disconnect(handle int64) {
	pg_QIconEnginePlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QIconEnginePlugin_virtIdx(slot int64) int32 {
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

func QIconEnginePlugin_enable(h uintptr, idx int32) { pg_QIconEnginePlugin__set_override(h, idx) }

func NewQIconEnginePlugin() *QIconEnginePlugin {
	id := nextVirtID()
	o := newownedQIconEnginePlugin(pg_QIconEnginePlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QIconEnginePlugin_virtIdx, QIconEnginePlugin_enable)
	return o
}

func NewQIconEnginePlugin2(parent *QObject) *QIconEnginePlugin {
	id := nextVirtID()
	o := newownedQIconEnginePlugin(pg_QIconEnginePlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QIconEnginePlugin_virtIdx, QIconEnginePlugin_enable)
	return o
}

func (this *QIconEnginePlugin) OnCreate(fn func(string) *QIconEngine) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(filename uintptr, filename_len int64) uintptr { return uintptr(fn(pgStrNoFree(filename, filename_len)).UnsafePointer()) }) // slot "create"
}

func (this *QIconEnginePlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QIconEnginePlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QIconEnginePlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QIconEnginePlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QIconEnginePlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QIconEnginePlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QIconEnginePlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QIconEnginePlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QIconEnginePlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIconEnginePlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
