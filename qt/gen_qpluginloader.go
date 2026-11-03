package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPluginLoader_base func(uintptr) uintptr
	addr_QPluginLoader_new uintptr
	addr_QPluginLoader_new2 uintptr
	addr_QPluginLoader_new3 uintptr
	addr_QPluginLoader_new4 uintptr
	addr_QPluginLoader_tr uintptr
	addr_QPluginLoader_instance uintptr
	addr_QPluginLoader_metaData uintptr
	addr_QPluginLoader_staticInstances uintptr
	addr_QPluginLoader_staticPlugins uintptr
	addr_QPluginLoader_load uintptr
	addr_QPluginLoader_unload uintptr
	addr_QPluginLoader_isLoaded uintptr
	addr_QPluginLoader_setFileName uintptr
	addr_QPluginLoader_fileName uintptr
	addr_QPluginLoader_errorString uintptr
	addr_QPluginLoader_setLoadHints uintptr
	addr_QPluginLoader_loadHints uintptr
	addr_QPluginLoader_connect uintptr
	addr_QPluginLoader_connect2 uintptr
	addr_QPluginLoader_sender uintptr
	addr_QPluginLoader_senderSignalIndex uintptr
	addr_QPluginLoader_receivers uintptr
	addr_QPluginLoader_isSignalConnected uintptr
	addr_QPluginLoader_tr2 uintptr
	addr_QPluginLoader_tr3 uintptr
	addr_QPluginLoader_connect3 uintptr
	addr_QPluginLoader_connect4 uintptr
	addr_QPluginLoader_metaObject_default uintptr
	addr_QPluginLoader_event_default uintptr
	addr_QPluginLoader_eventFilter_default uintptr
	addr_QPluginLoader_timerEvent_default uintptr
	addr_QPluginLoader_childEvent_default uintptr
	addr_QPluginLoader_customEvent_default uintptr
	addr_QPluginLoader_connectNotify_default uintptr
	addr_QPluginLoader_disconnectNotify_default uintptr
	pg_QPluginLoader_connect_destroyed func(uintptr, int64)
	pg_QPluginLoader_connect_destroyed2 func(uintptr, int64)
	pg_QPluginLoader_connect_objectNameChanged func(uintptr, int64)
	pg_QPluginLoader__disconnect func(int64)
	pg_QPluginLoader__set_override func(uintptr, int32)
	pg_QPluginLoader_delete func(uintptr)
)

func init() {
	reg(&pg_QPluginLoader_base, "QPluginLoader_base")
	symaddr(&addr_QPluginLoader_new, "QPluginLoader_new")
	symaddr(&addr_QPluginLoader_new2, "QPluginLoader_new2")
	symaddr(&addr_QPluginLoader_new3, "QPluginLoader_new3")
	symaddr(&addr_QPluginLoader_new4, "QPluginLoader_new4")
	symaddr(&addr_QPluginLoader_tr, "QPluginLoader_tr")
	symaddr(&addr_QPluginLoader_instance, "QPluginLoader_instance")
	symaddr(&addr_QPluginLoader_metaData, "QPluginLoader_metaData")
	symaddr(&addr_QPluginLoader_staticInstances, "QPluginLoader_staticInstances")
	symaddr(&addr_QPluginLoader_staticPlugins, "QPluginLoader_staticPlugins")
	symaddr(&addr_QPluginLoader_load, "QPluginLoader_load")
	symaddr(&addr_QPluginLoader_unload, "QPluginLoader_unload")
	symaddr(&addr_QPluginLoader_isLoaded, "QPluginLoader_isLoaded")
	symaddr(&addr_QPluginLoader_setFileName, "QPluginLoader_setFileName")
	symaddr(&addr_QPluginLoader_fileName, "QPluginLoader_fileName")
	symaddr(&addr_QPluginLoader_errorString, "QPluginLoader_errorString")
	symaddr(&addr_QPluginLoader_setLoadHints, "QPluginLoader_setLoadHints")
	symaddr(&addr_QPluginLoader_loadHints, "QPluginLoader_loadHints")
	symaddr(&addr_QPluginLoader_connect, "QPluginLoader_connect")
	symaddr(&addr_QPluginLoader_connect2, "QPluginLoader_connect2")
	symaddr(&addr_QPluginLoader_sender, "QPluginLoader_sender")
	symaddr(&addr_QPluginLoader_senderSignalIndex, "QPluginLoader_senderSignalIndex")
	symaddr(&addr_QPluginLoader_receivers, "QPluginLoader_receivers")
	symaddr(&addr_QPluginLoader_isSignalConnected, "QPluginLoader_isSignalConnected")
	symaddr(&addr_QPluginLoader_tr2, "QPluginLoader_tr2")
	symaddr(&addr_QPluginLoader_tr3, "QPluginLoader_tr3")
	symaddr(&addr_QPluginLoader_connect3, "QPluginLoader_connect3")
	symaddr(&addr_QPluginLoader_connect4, "QPluginLoader_connect4")
	symaddr(&addr_QPluginLoader_metaObject_default, "QPluginLoader_metaObject_default")
	symaddr(&addr_QPluginLoader_event_default, "QPluginLoader_event_default")
	symaddr(&addr_QPluginLoader_eventFilter_default, "QPluginLoader_eventFilter_default")
	symaddr(&addr_QPluginLoader_timerEvent_default, "QPluginLoader_timerEvent_default")
	symaddr(&addr_QPluginLoader_childEvent_default, "QPluginLoader_childEvent_default")
	symaddr(&addr_QPluginLoader_customEvent_default, "QPluginLoader_customEvent_default")
	symaddr(&addr_QPluginLoader_connectNotify_default, "QPluginLoader_connectNotify_default")
	symaddr(&addr_QPluginLoader_disconnectNotify_default, "QPluginLoader_disconnectNotify_default")
	reg(&pg_QPluginLoader_connect_destroyed, "QPluginLoader_connect_destroyed")
	reg(&pg_QPluginLoader_connect_destroyed2, "QPluginLoader_connect_destroyed2")
	reg(&pg_QPluginLoader_connect_objectNameChanged, "QPluginLoader_connect_objectNameChanged")
	reg(&pg_QPluginLoader__disconnect, "QPluginLoader__disconnect")
	reg(&pg_QPluginLoader__set_override, "QPluginLoader__set_override")
	reg(&pg_QPluginLoader_delete, "QPluginLoader_delete")
}

type QPluginLoader struct {
	h uintptr
	*QObject
}

func bindQPluginLoader(h uintptr) *QPluginLoader {
	if h == 0 {
		return nil
	}
	return &QPluginLoader{h: h, QObject: bindQObject(pg_QPluginLoader_base(h))}
}

func newownedQPluginLoader(h uintptr) *QPluginLoader { return bindQPluginLoader(h) }

func (this *QPluginLoader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPluginLoader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPluginLoader(h unsafe.Pointer) *QPluginLoader { return bindQPluginLoader(uintptr(h)) }
func UnsafeNewQPluginLoaderOwned(h unsafe.Pointer) *QPluginLoader { return newownedQPluginLoader(uintptr(h)) }

func (this *QPluginLoader) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQPluginLoader() *QPluginLoader {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_new, uintptr(__goid))
	__o := newownedQPluginLoader(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPluginLoader_virtIdx, QPluginLoader_enable)
	return __o
}

func NewQPluginLoader2(fileName string) *QPluginLoader {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_new2, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQPluginLoader(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPluginLoader_virtIdx, QPluginLoader_enable)
	return __o
}

func NewQPluginLoader3(parent *QObject) *QPluginLoader {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPluginLoader(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPluginLoader_virtIdx, QPluginLoader_enable)
	return __o
}

func NewQPluginLoader4(fileName string, parent *QObject) *QPluginLoader {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_new4, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(parent.UnsafePointer()))
	__o := newownedQPluginLoader(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPluginLoader_virtIdx, QPluginLoader_enable)
	return __o
}

func QPluginLoader_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPluginLoader) Instance() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_instance, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPluginLoader) MetaData() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_metaData, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func QPluginLoader_StaticInstances() []*QObject {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_staticInstances, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QObject, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQObject(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPluginLoader_StaticPlugins() []*QStaticPlugin {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_staticPlugins, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QStaticPlugin, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStaticPlugin(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPluginLoader) Load() bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_load, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPluginLoader) Unload() bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_unload, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPluginLoader) IsLoaded() bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_isLoaded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPluginLoader) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QPluginLoader_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QPluginLoader) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPluginLoader) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPluginLoader) SetLoadHints(loadHints QLibrary__LoadHint) {
	purego.SyscallN(addr_QPluginLoader_setLoadHints, this.h, uintptr(loadHints))
	runtime.KeepAlive(this)
}

func (this *QPluginLoader) LoadHints() QLibrary__LoadHint {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_loadHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLibrary__LoadHint(__r)
}

func QPluginLoader_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPluginLoader_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPluginLoader) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPluginLoader) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPluginLoader) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPluginLoader) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPluginLoader_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPluginLoader_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPluginLoader_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPluginLoader_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPluginLoader) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPluginLoader_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPluginLoader) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPluginLoader) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPluginLoader_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPluginLoader) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPluginLoader_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPluginLoader) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPluginLoader_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPluginLoader) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPluginLoader_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPluginLoader) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPluginLoader_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPluginLoader) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPluginLoader_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPluginLoader) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPluginLoader_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPluginLoader) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPluginLoader_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPluginLoader) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPluginLoader_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPluginLoader) Disconnect(handle int64) {
	pg_QPluginLoader__disconnect(handle)
	unregisterCallback(handle)
}

func QPluginLoader_virtIdx(slot int64) int32 {
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

func QPluginLoader_enable(h uintptr, idx int32) { pg_QPluginLoader__set_override(h, idx) }

func (this *QPluginLoader) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPluginLoader) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QPluginLoader) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPluginLoader) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPluginLoader) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPluginLoader) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPluginLoader) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPluginLoader) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPluginLoader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPluginLoader_delete(this.h)
	}
	this.UnsafeRelease()
}
