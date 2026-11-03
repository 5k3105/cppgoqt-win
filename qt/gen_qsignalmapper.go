package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSignalMapper_base func(uintptr) uintptr
	addr_QSignalMapper_new uintptr
	addr_QSignalMapper_new2 uintptr
	addr_QSignalMapper_tr uintptr
	addr_QSignalMapper_setMapping uintptr
	addr_QSignalMapper_setMapping2 uintptr
	addr_QSignalMapper_setMapping3 uintptr
	addr_QSignalMapper_removeMappings uintptr
	addr_QSignalMapper_mapping uintptr
	addr_QSignalMapper_mapping2 uintptr
	addr_QSignalMapper_mapping3 uintptr
	addr_QSignalMapper_mappedInt uintptr
	addr_QSignalMapper_mappedString uintptr
	addr_QSignalMapper_mappedObject uintptr
	addr_QSignalMapper_map uintptr
	addr_QSignalMapper_map2 uintptr
	addr_QSignalMapper_connect uintptr
	addr_QSignalMapper_connect2 uintptr
	addr_QSignalMapper_sender uintptr
	addr_QSignalMapper_senderSignalIndex uintptr
	addr_QSignalMapper_receivers uintptr
	addr_QSignalMapper_isSignalConnected uintptr
	addr_QSignalMapper_tr2 uintptr
	addr_QSignalMapper_tr3 uintptr
	addr_QSignalMapper_connect3 uintptr
	addr_QSignalMapper_connect4 uintptr
	addr_QSignalMapper_metaObject_default uintptr
	addr_QSignalMapper_event_default uintptr
	addr_QSignalMapper_eventFilter_default uintptr
	addr_QSignalMapper_timerEvent_default uintptr
	addr_QSignalMapper_childEvent_default uintptr
	addr_QSignalMapper_customEvent_default uintptr
	addr_QSignalMapper_connectNotify_default uintptr
	addr_QSignalMapper_disconnectNotify_default uintptr
	pg_QSignalMapper_connect_mappedInt func(uintptr, int64)
	pg_QSignalMapper_connect_mappedString func(uintptr, int64)
	pg_QSignalMapper_connect_mappedObject func(uintptr, int64)
	pg_QSignalMapper_connect_destroyed func(uintptr, int64)
	pg_QSignalMapper_connect_destroyed2 func(uintptr, int64)
	pg_QSignalMapper_connect_objectNameChanged func(uintptr, int64)
	pg_QSignalMapper__disconnect func(int64)
	pg_QSignalMapper__set_override func(uintptr, int32)
	pg_QSignalMapper_delete func(uintptr)
)

func init() {
	reg(&pg_QSignalMapper_base, "QSignalMapper_base")
	symaddr(&addr_QSignalMapper_new, "QSignalMapper_new")
	symaddr(&addr_QSignalMapper_new2, "QSignalMapper_new2")
	symaddr(&addr_QSignalMapper_tr, "QSignalMapper_tr")
	symaddr(&addr_QSignalMapper_setMapping, "QSignalMapper_setMapping")
	symaddr(&addr_QSignalMapper_setMapping2, "QSignalMapper_setMapping2")
	symaddr(&addr_QSignalMapper_setMapping3, "QSignalMapper_setMapping3")
	symaddr(&addr_QSignalMapper_removeMappings, "QSignalMapper_removeMappings")
	symaddr(&addr_QSignalMapper_mapping, "QSignalMapper_mapping")
	symaddr(&addr_QSignalMapper_mapping2, "QSignalMapper_mapping2")
	symaddr(&addr_QSignalMapper_mapping3, "QSignalMapper_mapping3")
	symaddr(&addr_QSignalMapper_mappedInt, "QSignalMapper_mappedInt")
	symaddr(&addr_QSignalMapper_mappedString, "QSignalMapper_mappedString")
	symaddr(&addr_QSignalMapper_mappedObject, "QSignalMapper_mappedObject")
	symaddr(&addr_QSignalMapper_map, "QSignalMapper_map")
	symaddr(&addr_QSignalMapper_map2, "QSignalMapper_map2")
	symaddr(&addr_QSignalMapper_connect, "QSignalMapper_connect")
	symaddr(&addr_QSignalMapper_connect2, "QSignalMapper_connect2")
	symaddr(&addr_QSignalMapper_sender, "QSignalMapper_sender")
	symaddr(&addr_QSignalMapper_senderSignalIndex, "QSignalMapper_senderSignalIndex")
	symaddr(&addr_QSignalMapper_receivers, "QSignalMapper_receivers")
	symaddr(&addr_QSignalMapper_isSignalConnected, "QSignalMapper_isSignalConnected")
	symaddr(&addr_QSignalMapper_tr2, "QSignalMapper_tr2")
	symaddr(&addr_QSignalMapper_tr3, "QSignalMapper_tr3")
	symaddr(&addr_QSignalMapper_connect3, "QSignalMapper_connect3")
	symaddr(&addr_QSignalMapper_connect4, "QSignalMapper_connect4")
	symaddr(&addr_QSignalMapper_metaObject_default, "QSignalMapper_metaObject_default")
	symaddr(&addr_QSignalMapper_event_default, "QSignalMapper_event_default")
	symaddr(&addr_QSignalMapper_eventFilter_default, "QSignalMapper_eventFilter_default")
	symaddr(&addr_QSignalMapper_timerEvent_default, "QSignalMapper_timerEvent_default")
	symaddr(&addr_QSignalMapper_childEvent_default, "QSignalMapper_childEvent_default")
	symaddr(&addr_QSignalMapper_customEvent_default, "QSignalMapper_customEvent_default")
	symaddr(&addr_QSignalMapper_connectNotify_default, "QSignalMapper_connectNotify_default")
	symaddr(&addr_QSignalMapper_disconnectNotify_default, "QSignalMapper_disconnectNotify_default")
	reg(&pg_QSignalMapper_connect_mappedInt, "QSignalMapper_connect_mappedInt")
	reg(&pg_QSignalMapper_connect_mappedString, "QSignalMapper_connect_mappedString")
	reg(&pg_QSignalMapper_connect_mappedObject, "QSignalMapper_connect_mappedObject")
	reg(&pg_QSignalMapper_connect_destroyed, "QSignalMapper_connect_destroyed")
	reg(&pg_QSignalMapper_connect_destroyed2, "QSignalMapper_connect_destroyed2")
	reg(&pg_QSignalMapper_connect_objectNameChanged, "QSignalMapper_connect_objectNameChanged")
	reg(&pg_QSignalMapper__disconnect, "QSignalMapper__disconnect")
	reg(&pg_QSignalMapper__set_override, "QSignalMapper__set_override")
	reg(&pg_QSignalMapper_delete, "QSignalMapper_delete")
}

type QSignalMapper struct {
	h uintptr
	*QObject
}

func bindQSignalMapper(h uintptr) *QSignalMapper {
	if h == 0 {
		return nil
	}
	return &QSignalMapper{h: h, QObject: bindQObject(pg_QSignalMapper_base(h))}
}

func newownedQSignalMapper(h uintptr) *QSignalMapper { return bindQSignalMapper(h) }

func (this *QSignalMapper) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSignalMapper) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSignalMapper(h unsafe.Pointer) *QSignalMapper { return bindQSignalMapper(uintptr(h)) }
func UnsafeNewQSignalMapperOwned(h unsafe.Pointer) *QSignalMapper { return newownedQSignalMapper(uintptr(h)) }

func (this *QSignalMapper) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQSignalMapper() *QSignalMapper {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_new, uintptr(__goid))
	__o := newownedQSignalMapper(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSignalMapper_virtIdx, QSignalMapper_enable)
	return __o
}

func NewQSignalMapper2(parent *QObject) *QSignalMapper {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSignalMapper(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSignalMapper_virtIdx, QSignalMapper_enable)
	return __o
}

func QSignalMapper_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSignalMapper) SetMapping(sender *QObject, id int) {
	purego.SyscallN(addr_QSignalMapper_setMapping, this.h, uintptr(sender.UnsafePointer()), uintptr(id))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
}

func (this *QSignalMapper) SetMapping2(sender *QObject, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QSignalMapper_setMapping2, this.h, uintptr(sender.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
	runtime.KeepAlive(text)
}

func (this *QSignalMapper) SetMapping3(sender *QObject, object *QObject) {
	purego.SyscallN(addr_QSignalMapper_setMapping3, this.h, uintptr(sender.UnsafePointer()), uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
	runtime.KeepAlive(object)
}

func (this *QSignalMapper) RemoveMappings(sender *QObject) {
	purego.SyscallN(addr_QSignalMapper_removeMappings, this.h, uintptr(sender.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
}

func (this *QSignalMapper) Mapping(id int) *QObject {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_mapping, this.h, uintptr(id))
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSignalMapper) MappingWithText(text string) *QObject {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_mapping2, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQObject(__r)
}

func (this *QSignalMapper) MappingWithObject(object *QObject) *QObject {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_mapping3, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return bindQObject(__r)
}

func (this *QSignalMapper) MappedInt(arg1 int) {
	purego.SyscallN(addr_QSignalMapper_mappedInt, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSignalMapper) MappedString(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QSignalMapper_mappedString, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSignalMapper) MappedObject(arg1 *QObject) {
	purego.SyscallN(addr_QSignalMapper_mappedObject, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSignalMapper) Map() {
	purego.SyscallN(addr_QSignalMapper_map, this.h)
	runtime.KeepAlive(this)
}

func (this *QSignalMapper) MapWithSender(sender *QObject) {
	purego.SyscallN(addr_QSignalMapper_map2, this.h, uintptr(sender.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
}

func QSignalMapper_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSignalMapper_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSignalMapper) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSignalMapper) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSignalMapper_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSignalMapper) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSignalMapper_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSignalMapper) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSignalMapper_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSignalMapper_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSignalMapper_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSignalMapper_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSignalMapper_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSignalMapper) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSignalMapper_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSignalMapper) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSignalMapper_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSignalMapper) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSignalMapper_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSignalMapper) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSignalMapper_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSignalMapper) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSignalMapper_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSignalMapper) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSignalMapper_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSignalMapper) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSignalMapper_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSignalMapper) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSignalMapper_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSignalMapper) OnMappedInt(fn func(int)) int64 {
	id := registerCallback(func(arg1 int32) { fn(int(arg1)) })
	pg_QSignalMapper_connect_mappedInt(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) OnMappedString(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QSignalMapper_connect_mappedString(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) OnMappedObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSignalMapper_connect_mappedObject(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSignalMapper_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSignalMapper_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSignalMapper_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSignalMapper) Disconnect(handle int64) {
	pg_QSignalMapper__disconnect(handle)
	unregisterCallback(handle)
}

func QSignalMapper_virtIdx(slot int64) int32 {
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

func QSignalMapper_enable(h uintptr, idx int32) { pg_QSignalMapper__set_override(h, idx) }

func (this *QSignalMapper) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSignalMapper) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSignalMapper) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSignalMapper) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSignalMapper) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSignalMapper) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSignalMapper) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSignalMapper) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSignalMapper) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSignalMapper_delete(this.h)
	}
	this.UnsafeRelease()
}
