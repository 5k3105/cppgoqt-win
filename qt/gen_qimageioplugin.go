package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QImageIOPlugin_base func(uintptr) uintptr
	addr_QImageIOPlugin_tr uintptr
	addr_QImageIOPlugin_capabilities uintptr
	addr_QImageIOPlugin_create uintptr
	addr_QImageIOPlugin_connect uintptr
	addr_QImageIOPlugin_connect2 uintptr
	addr_QImageIOPlugin_sender uintptr
	addr_QImageIOPlugin_senderSignalIndex uintptr
	addr_QImageIOPlugin_receivers uintptr
	addr_QImageIOPlugin_isSignalConnected uintptr
	addr_QImageIOPlugin_tr2 uintptr
	addr_QImageIOPlugin_tr3 uintptr
	addr_QImageIOPlugin_connect3 uintptr
	addr_QImageIOPlugin_connect4 uintptr
	addr_QImageIOPlugin_metaObject_default uintptr
	addr_QImageIOPlugin_event_default uintptr
	addr_QImageIOPlugin_eventFilter_default uintptr
	addr_QImageIOPlugin_timerEvent_default uintptr
	addr_QImageIOPlugin_childEvent_default uintptr
	addr_QImageIOPlugin_customEvent_default uintptr
	addr_QImageIOPlugin_connectNotify_default uintptr
	addr_QImageIOPlugin_disconnectNotify_default uintptr
	pg_QImageIOPlugin_connect_destroyed func(uintptr, int64)
	pg_QImageIOPlugin_connect_destroyed2 func(uintptr, int64)
	pg_QImageIOPlugin_connect_objectNameChanged func(uintptr, int64)
	pg_QImageIOPlugin__disconnect func(int64)
	pg_QImageIOPlugin__set_override func(uintptr, int32)
	pg_QImageIOPlugin__Go_new func(int64, uintptr) uintptr
	pg_QImageIOPlugin_delete func(uintptr)
)

func init() {
	reg(&pg_QImageIOPlugin_base, "QImageIOPlugin_base")
	symaddr(&addr_QImageIOPlugin_tr, "QImageIOPlugin_tr")
	symaddr(&addr_QImageIOPlugin_capabilities, "QImageIOPlugin_capabilities")
	symaddr(&addr_QImageIOPlugin_create, "QImageIOPlugin_create")
	symaddr(&addr_QImageIOPlugin_connect, "QImageIOPlugin_connect")
	symaddr(&addr_QImageIOPlugin_connect2, "QImageIOPlugin_connect2")
	symaddr(&addr_QImageIOPlugin_sender, "QImageIOPlugin_sender")
	symaddr(&addr_QImageIOPlugin_senderSignalIndex, "QImageIOPlugin_senderSignalIndex")
	symaddr(&addr_QImageIOPlugin_receivers, "QImageIOPlugin_receivers")
	symaddr(&addr_QImageIOPlugin_isSignalConnected, "QImageIOPlugin_isSignalConnected")
	symaddr(&addr_QImageIOPlugin_tr2, "QImageIOPlugin_tr2")
	symaddr(&addr_QImageIOPlugin_tr3, "QImageIOPlugin_tr3")
	symaddr(&addr_QImageIOPlugin_connect3, "QImageIOPlugin_connect3")
	symaddr(&addr_QImageIOPlugin_connect4, "QImageIOPlugin_connect4")
	symaddr(&addr_QImageIOPlugin_metaObject_default, "QImageIOPlugin_metaObject_default")
	symaddr(&addr_QImageIOPlugin_event_default, "QImageIOPlugin_event_default")
	symaddr(&addr_QImageIOPlugin_eventFilter_default, "QImageIOPlugin_eventFilter_default")
	symaddr(&addr_QImageIOPlugin_timerEvent_default, "QImageIOPlugin_timerEvent_default")
	symaddr(&addr_QImageIOPlugin_childEvent_default, "QImageIOPlugin_childEvent_default")
	symaddr(&addr_QImageIOPlugin_customEvent_default, "QImageIOPlugin_customEvent_default")
	symaddr(&addr_QImageIOPlugin_connectNotify_default, "QImageIOPlugin_connectNotify_default")
	symaddr(&addr_QImageIOPlugin_disconnectNotify_default, "QImageIOPlugin_disconnectNotify_default")
	reg(&pg_QImageIOPlugin_connect_destroyed, "QImageIOPlugin_connect_destroyed")
	reg(&pg_QImageIOPlugin_connect_destroyed2, "QImageIOPlugin_connect_destroyed2")
	reg(&pg_QImageIOPlugin_connect_objectNameChanged, "QImageIOPlugin_connect_objectNameChanged")
	reg(&pg_QImageIOPlugin__disconnect, "QImageIOPlugin__disconnect")
	reg(&pg_QImageIOPlugin__set_override, "QImageIOPlugin__set_override")
	reg(&pg_QImageIOPlugin__Go_new, "QImageIOPlugin__Go_new")
	reg(&pg_QImageIOPlugin_delete, "QImageIOPlugin_delete")
}

type QImageIOPlugin struct {
	h uintptr
	*QObject
}

func bindQImageIOPlugin(h uintptr) *QImageIOPlugin {
	if h == 0 {
		return nil
	}
	return &QImageIOPlugin{h: h, QObject: bindQObject(pg_QImageIOPlugin_base(h))}
}

func newownedQImageIOPlugin(h uintptr) *QImageIOPlugin { return bindQImageIOPlugin(h) }

func (this *QImageIOPlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QImageIOPlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQImageIOPlugin(h unsafe.Pointer) *QImageIOPlugin { return bindQImageIOPlugin(uintptr(h)) }
func UnsafeNewQImageIOPluginOwned(h unsafe.Pointer) *QImageIOPlugin { return newownedQImageIOPlugin(uintptr(h)) }

func (this *QImageIOPlugin) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QImageIOPlugin_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QImageIOPlugin) Capabilities(device *QIODevice, format *QByteArray) QImageIOPlugin__Capability {
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_capabilities, this.h, uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	return QImageIOPlugin__Capability(__r)
}

func (this *QImageIOPlugin) Create(device *QIODevice, format *QByteArray) *QImageIOHandler {
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_create, this.h, uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	return bindQImageIOHandler(__r)
}

func QImageIOPlugin_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QImageIOPlugin_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QImageIOPlugin) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QImageIOPlugin) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOPlugin) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QImageIOPlugin) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QImageIOPlugin_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QImageIOPlugin_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QImageIOPlugin_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QImageIOPlugin_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QImageIOPlugin) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QImageIOPlugin_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QImageIOPlugin) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QImageIOPlugin) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOPlugin_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QImageIOPlugin) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QImageIOPlugin_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QImageIOPlugin) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QImageIOPlugin_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QImageIOPlugin) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QImageIOPlugin_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QImageIOPlugin) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QImageIOPlugin_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QImageIOPlugin) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QImageIOPlugin_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QImageIOPlugin) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QImageIOPlugin_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QImageIOPlugin) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QImageIOPlugin_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QImageIOPlugin) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QImageIOPlugin_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QImageIOPlugin) Disconnect(handle int64) {
	pg_QImageIOPlugin__disconnect(handle)
	unregisterCallback(handle)
}

func QImageIOPlugin_virtIdx(slot int64) int32 {
	switch slot {
	case 5572960902616693279:
		return 0
	case -626457835795418115:
		return 1
	case 1607084412591819503:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
		return 4
	case -5409927581812374498:
		return 5
	case -1185315367526857833:
		return 6
	case 3769868229123868474:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	}
	return -1
}

func QImageIOPlugin_enable(h uintptr, idx int32) { pg_QImageIOPlugin__set_override(h, idx) }

func NewQImageIOPlugin() *QImageIOPlugin {
	id := nextVirtID()
	o := newownedQImageIOPlugin(pg_QImageIOPlugin__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QImageIOPlugin_virtIdx, QImageIOPlugin_enable)
	return o
}

func NewQImageIOPlugin2(parent *QObject) *QImageIOPlugin {
	id := nextVirtID()
	o := newownedQImageIOPlugin(pg_QImageIOPlugin__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QImageIOPlugin_virtIdx, QImageIOPlugin_enable)
	return o
}

func (this *QImageIOPlugin) OnCapabilities(fn func(*QIODevice, *QByteArray) QImageIOPlugin__Capability) {
	registerVirt(this.UnsafeGoID(), 5572960902616693279, func(device uintptr, format uintptr) int32 { return int32(fn(UnsafeNewQIODevice(unsafe.Pointer(device)), UnsafeNewQByteArray(unsafe.Pointer(format)))) }) // slot "capabilities"
}

func (this *QImageIOPlugin) OnCreate(fn func(*QIODevice, *QByteArray) *QImageIOHandler) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(device uintptr, format uintptr) uintptr { return uintptr(fn(UnsafeNewQIODevice(unsafe.Pointer(device)), UnsafeNewQByteArray(unsafe.Pointer(format))).UnsafePointer()) }) // slot "create"
}

func (this *QImageIOPlugin) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QImageIOPlugin) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QImageIOPlugin) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QImageIOPlugin) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QImageIOPlugin) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QImageIOPlugin) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QImageIOPlugin) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QImageIOPlugin) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QImageIOPlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QImageIOPlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
