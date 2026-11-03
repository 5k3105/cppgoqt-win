package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QInputDevice_base func(uintptr) uintptr
	addr_QInputDevice_new uintptr
	addr_QInputDevice_new2 uintptr
	addr_QInputDevice_new3 uintptr
	addr_QInputDevice_new4 uintptr
	addr_QInputDevice_new5 uintptr
	addr_QInputDevice_tr uintptr
	addr_QInputDevice_name uintptr
	addr_QInputDevice_type uintptr
	addr_QInputDevice_capabilities uintptr
	addr_QInputDevice_hasCapability uintptr
	addr_QInputDevice_systemId uintptr
	addr_QInputDevice_seatName uintptr
	addr_QInputDevice_availableVirtualGeometry uintptr
	addr_QInputDevice_seatNames uintptr
	addr_QInputDevice_devices uintptr
	addr_QInputDevice_primaryKeyboard uintptr
	addr_QInputDevice_OperatorEqual uintptr
	addr_QInputDevice_availableVirtualGeometryChanged uintptr
	addr_QInputDevice_capabilitiesChanged uintptr
	addr_QInputDevice_connect uintptr
	addr_QInputDevice_connect2 uintptr
	addr_QInputDevice_sender uintptr
	addr_QInputDevice_senderSignalIndex uintptr
	addr_QInputDevice_receivers uintptr
	addr_QInputDevice_isSignalConnected uintptr
	addr_QInputDevice_tr2 uintptr
	addr_QInputDevice_tr3 uintptr
	addr_QInputDevice_primaryKeyboard2 uintptr
	addr_QInputDevice_connect3 uintptr
	addr_QInputDevice_connect4 uintptr
	addr_QInputDevice_metaObject_default uintptr
	addr_QInputDevice_event_default uintptr
	addr_QInputDevice_eventFilter_default uintptr
	addr_QInputDevice_timerEvent_default uintptr
	addr_QInputDevice_childEvent_default uintptr
	addr_QInputDevice_customEvent_default uintptr
	addr_QInputDevice_connectNotify_default uintptr
	addr_QInputDevice_disconnectNotify_default uintptr
	pg_QInputDevice_connect_availableVirtualGeometryChanged func(uintptr, int64)
	pg_QInputDevice_connect_capabilitiesChanged func(uintptr, int64)
	pg_QInputDevice_connect_destroyed func(uintptr, int64)
	pg_QInputDevice_connect_destroyed2 func(uintptr, int64)
	pg_QInputDevice_connect_objectNameChanged func(uintptr, int64)
	pg_QInputDevice__disconnect func(int64)
	pg_QInputDevice__set_override func(uintptr, int32)
	pg_QInputDevice_delete func(uintptr)
)

func init() {
	reg(&pg_QInputDevice_base, "QInputDevice_base")
	symaddr(&addr_QInputDevice_new, "QInputDevice_new")
	symaddr(&addr_QInputDevice_new2, "QInputDevice_new2")
	symaddr(&addr_QInputDevice_new3, "QInputDevice_new3")
	symaddr(&addr_QInputDevice_new4, "QInputDevice_new4")
	symaddr(&addr_QInputDevice_new5, "QInputDevice_new5")
	symaddr(&addr_QInputDevice_tr, "QInputDevice_tr")
	symaddr(&addr_QInputDevice_name, "QInputDevice_name")
	symaddr(&addr_QInputDevice_type, "QInputDevice_type")
	symaddr(&addr_QInputDevice_capabilities, "QInputDevice_capabilities")
	symaddr(&addr_QInputDevice_hasCapability, "QInputDevice_hasCapability")
	symaddr(&addr_QInputDevice_systemId, "QInputDevice_systemId")
	symaddr(&addr_QInputDevice_seatName, "QInputDevice_seatName")
	symaddr(&addr_QInputDevice_availableVirtualGeometry, "QInputDevice_availableVirtualGeometry")
	symaddr(&addr_QInputDevice_seatNames, "QInputDevice_seatNames")
	symaddr(&addr_QInputDevice_devices, "QInputDevice_devices")
	symaddr(&addr_QInputDevice_primaryKeyboard, "QInputDevice_primaryKeyboard")
	symaddr(&addr_QInputDevice_OperatorEqual, "QInputDevice_OperatorEqual")
	symaddr(&addr_QInputDevice_availableVirtualGeometryChanged, "QInputDevice_availableVirtualGeometryChanged")
	symaddr(&addr_QInputDevice_capabilitiesChanged, "QInputDevice_capabilitiesChanged")
	symaddr(&addr_QInputDevice_connect, "QInputDevice_connect")
	symaddr(&addr_QInputDevice_connect2, "QInputDevice_connect2")
	symaddr(&addr_QInputDevice_sender, "QInputDevice_sender")
	symaddr(&addr_QInputDevice_senderSignalIndex, "QInputDevice_senderSignalIndex")
	symaddr(&addr_QInputDevice_receivers, "QInputDevice_receivers")
	symaddr(&addr_QInputDevice_isSignalConnected, "QInputDevice_isSignalConnected")
	symaddr(&addr_QInputDevice_tr2, "QInputDevice_tr2")
	symaddr(&addr_QInputDevice_tr3, "QInputDevice_tr3")
	symaddr(&addr_QInputDevice_primaryKeyboard2, "QInputDevice_primaryKeyboard2")
	symaddr(&addr_QInputDevice_connect3, "QInputDevice_connect3")
	symaddr(&addr_QInputDevice_connect4, "QInputDevice_connect4")
	symaddr(&addr_QInputDevice_metaObject_default, "QInputDevice_metaObject_default")
	symaddr(&addr_QInputDevice_event_default, "QInputDevice_event_default")
	symaddr(&addr_QInputDevice_eventFilter_default, "QInputDevice_eventFilter_default")
	symaddr(&addr_QInputDevice_timerEvent_default, "QInputDevice_timerEvent_default")
	symaddr(&addr_QInputDevice_childEvent_default, "QInputDevice_childEvent_default")
	symaddr(&addr_QInputDevice_customEvent_default, "QInputDevice_customEvent_default")
	symaddr(&addr_QInputDevice_connectNotify_default, "QInputDevice_connectNotify_default")
	symaddr(&addr_QInputDevice_disconnectNotify_default, "QInputDevice_disconnectNotify_default")
	reg(&pg_QInputDevice_connect_availableVirtualGeometryChanged, "QInputDevice_connect_availableVirtualGeometryChanged")
	reg(&pg_QInputDevice_connect_capabilitiesChanged, "QInputDevice_connect_capabilitiesChanged")
	reg(&pg_QInputDevice_connect_destroyed, "QInputDevice_connect_destroyed")
	reg(&pg_QInputDevice_connect_destroyed2, "QInputDevice_connect_destroyed2")
	reg(&pg_QInputDevice_connect_objectNameChanged, "QInputDevice_connect_objectNameChanged")
	reg(&pg_QInputDevice__disconnect, "QInputDevice__disconnect")
	reg(&pg_QInputDevice__set_override, "QInputDevice__set_override")
	reg(&pg_QInputDevice_delete, "QInputDevice_delete")
}

type QInputDevice struct {
	h uintptr
	*QObject
}

func bindQInputDevice(h uintptr) *QInputDevice {
	if h == 0 {
		return nil
	}
	return &QInputDevice{h: h, QObject: bindQObject(pg_QInputDevice_base(h))}
}

func newownedQInputDevice(h uintptr) *QInputDevice { return bindQInputDevice(h) }

func (this *QInputDevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputDevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputDevice(h unsafe.Pointer) *QInputDevice { return bindQInputDevice(uintptr(h)) }
func UnsafeNewQInputDeviceOwned(h unsafe.Pointer) *QInputDevice { return newownedQInputDevice(uintptr(h)) }

func (this *QInputDevice) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQInputDevice() *QInputDevice {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputDevice_new, uintptr(__goid))
	__o := newownedQInputDevice(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDevice_virtIdx, QInputDevice_enable)
	return __o
}

func NewQInputDevice2(name string, systemId int64, type_ QInputDevice__DeviceType) *QInputDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QInputDevice_new2, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(type_))
	__o := newownedQInputDevice(__r)
	runtime.KeepAlive(name)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDevice_virtIdx, QInputDevice_enable)
	return __o
}

func NewQInputDevice3(parent *QObject) *QInputDevice {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputDevice_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQInputDevice(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDevice_virtIdx, QInputDevice_enable)
	return __o
}

func NewQInputDevice4(name string, systemId int64, type_ QInputDevice__DeviceType, seatName string) *QInputDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QInputDevice_new4, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(type_), seatName_p, uintptr(len(seatName)))
	__o := newownedQInputDevice(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(seatName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDevice_virtIdx, QInputDevice_enable)
	return __o
}

func NewQInputDevice5(name string, systemId int64, type_ QInputDevice__DeviceType, seatName string, parent *QObject) *QInputDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QInputDevice_new5, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(type_), seatName_p, uintptr(len(seatName)), uintptr(parent.UnsafePointer()))
	__o := newownedQInputDevice(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(seatName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputDevice_virtIdx, QInputDevice_enable)
	return __o
}

func QInputDevice_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QInputDevice) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDevice) Type() QInputDevice__DeviceType {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputDevice__DeviceType(__r)
}

func (this *QInputDevice) Capabilities() QInputDevice__Capability {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_capabilities, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputDevice__Capability(__r)
}

func (this *QInputDevice) HasCapability(cap_ QInputDevice__Capability) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_hasCapability, this.h, uintptr(cap_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputDevice) SystemId() int64 {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_systemId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QInputDevice) SeatName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_seatName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputDevice) AvailableVirtualGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_availableVirtualGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func QInputDevice_SeatNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_seatNames, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QInputDevice_Devices() []*QInputDevice {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_devices, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QInputDevice, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQInputDevice(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QInputDevice_PrimaryKeyboard() *QInputDevice {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_primaryKeyboard)
	return bindQInputDevice(__r)
}

func (this *QInputDevice) OperatorEqual(other *QInputDevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QInputDevice) AvailableVirtualGeometryChanged(area *QRect) {
	purego.SyscallN(addr_QInputDevice_availableVirtualGeometryChanged, this.h, uintptr(area.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(area)
}

func (this *QInputDevice) CapabilitiesChanged(capabilities QInputDevice__Capability) {
	purego.SyscallN(addr_QInputDevice_capabilitiesChanged, this.h, uintptr(capabilities))
	runtime.KeepAlive(this)
}

func QInputDevice_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputDevice_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputDevice_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QInputDevice) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QInputDevice) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputDevice) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QInputDevice) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QInputDevice_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QInputDevice_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputDevice_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QInputDevice_PrimaryKeyboardWithSeatName(seatName string) *QInputDevice {
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QInputDevice_primaryKeyboard2, seatName_p, uintptr(len(seatName)))
	runtime.KeepAlive(seatName)
	return bindQInputDevice(__r)
}

func QInputDevice_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputDevice_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputDevice_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QInputDevice) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QInputDevice_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QInputDevice) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QInputDevice) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputDevice_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QInputDevice) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QInputDevice_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDevice) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QInputDevice_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDevice) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QInputDevice_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QInputDevice) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QInputDevice_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QInputDevice) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QInputDevice_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QInputDevice) OnAvailableVirtualGeometryChanged(fn func(*QRect)) int64 {
	id := registerCallback(func(area uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(area))) })
	pg_QInputDevice_connect_availableVirtualGeometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDevice) OnCapabilitiesChanged(fn func(QInputDevice__Capability)) int64 {
	id := registerCallback(func(capabilities int32) { fn(QInputDevice__Capability(capabilities)) })
	pg_QInputDevice_connect_capabilitiesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDevice) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputDevice_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDevice) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QInputDevice_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDevice) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QInputDevice_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputDevice) Disconnect(handle int64) {
	pg_QInputDevice__disconnect(handle)
	unregisterCallback(handle)
}

func QInputDevice_virtIdx(slot int64) int32 {
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

func QInputDevice_enable(h uintptr, idx int32) { pg_QInputDevice__set_override(h, idx) }

func (this *QInputDevice) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QInputDevice) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QInputDevice) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QInputDevice) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QInputDevice) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QInputDevice) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QInputDevice) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QInputDevice) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QInputDevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputDevice_delete(this.h)
	}
	this.UnsafeRelease()
}
