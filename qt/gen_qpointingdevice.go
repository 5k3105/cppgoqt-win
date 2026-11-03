package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPointingDevice_chain uintptr
	addr_QPointingDevice_new uintptr
	addr_QPointingDevice_new2 uintptr
	addr_QPointingDevice_new3 uintptr
	addr_QPointingDevice_new4 uintptr
	addr_QPointingDevice_new5 uintptr
	addr_QPointingDevice_new6 uintptr
	addr_QPointingDevice_tr uintptr
	addr_QPointingDevice_setType uintptr
	addr_QPointingDevice_setCapabilities uintptr
	addr_QPointingDevice_setMaximumTouchPoints uintptr
	addr_QPointingDevice_pointerType uintptr
	addr_QPointingDevice_maximumPoints uintptr
	addr_QPointingDevice_buttonCount uintptr
	addr_QPointingDevice_uniqueId uintptr
	addr_QPointingDevice_primaryPointingDevice uintptr
	addr_QPointingDevice_OperatorEqual uintptr
	addr_QPointingDevice_grabChanged uintptr
	addr_QPointingDevice_seatNames uintptr
	addr_QPointingDevice_devices uintptr
	addr_QPointingDevice_primaryKeyboard uintptr
	addr_QPointingDevice_connect uintptr
	addr_QPointingDevice_connect2 uintptr
	addr_QPointingDevice_sender uintptr
	addr_QPointingDevice_senderSignalIndex uintptr
	addr_QPointingDevice_receivers uintptr
	addr_QPointingDevice_isSignalConnected uintptr
	addr_QPointingDevice_tr2 uintptr
	addr_QPointingDevice_tr3 uintptr
	addr_QPointingDevice_primaryPointingDevice2 uintptr
	addr_QPointingDevice_primaryKeyboard2 uintptr
	addr_QPointingDevice_connect3 uintptr
	addr_QPointingDevice_connect4 uintptr
	addr_QPointingDevice_metaObject_default uintptr
	addr_QPointingDevice_event_default uintptr
	addr_QPointingDevice_eventFilter_default uintptr
	addr_QPointingDevice_timerEvent_default uintptr
	addr_QPointingDevice_childEvent_default uintptr
	addr_QPointingDevice_customEvent_default uintptr
	addr_QPointingDevice_connectNotify_default uintptr
	addr_QPointingDevice_disconnectNotify_default uintptr
	pg_QPointingDevice_connect_grabChanged func(uintptr, int64)
	pg_QPointingDevice_connect_availableVirtualGeometryChanged func(uintptr, int64)
	pg_QPointingDevice_connect_capabilitiesChanged func(uintptr, int64)
	pg_QPointingDevice_connect_destroyed func(uintptr, int64)
	pg_QPointingDevice_connect_destroyed2 func(uintptr, int64)
	pg_QPointingDevice_connect_objectNameChanged func(uintptr, int64)
	pg_QPointingDevice__disconnect func(int64)
	pg_QPointingDevice__set_override func(uintptr, int32)
	pg_QPointingDevice_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPointingDevice_chain, "QPointingDevice_chain")
	symaddr(&addr_QPointingDevice_new, "QPointingDevice_new")
	symaddr(&addr_QPointingDevice_new2, "QPointingDevice_new2")
	symaddr(&addr_QPointingDevice_new3, "QPointingDevice_new3")
	symaddr(&addr_QPointingDevice_new4, "QPointingDevice_new4")
	symaddr(&addr_QPointingDevice_new5, "QPointingDevice_new5")
	symaddr(&addr_QPointingDevice_new6, "QPointingDevice_new6")
	symaddr(&addr_QPointingDevice_tr, "QPointingDevice_tr")
	symaddr(&addr_QPointingDevice_setType, "QPointingDevice_setType")
	symaddr(&addr_QPointingDevice_setCapabilities, "QPointingDevice_setCapabilities")
	symaddr(&addr_QPointingDevice_setMaximumTouchPoints, "QPointingDevice_setMaximumTouchPoints")
	symaddr(&addr_QPointingDevice_pointerType, "QPointingDevice_pointerType")
	symaddr(&addr_QPointingDevice_maximumPoints, "QPointingDevice_maximumPoints")
	symaddr(&addr_QPointingDevice_buttonCount, "QPointingDevice_buttonCount")
	symaddr(&addr_QPointingDevice_uniqueId, "QPointingDevice_uniqueId")
	symaddr(&addr_QPointingDevice_primaryPointingDevice, "QPointingDevice_primaryPointingDevice")
	symaddr(&addr_QPointingDevice_OperatorEqual, "QPointingDevice_OperatorEqual")
	symaddr(&addr_QPointingDevice_grabChanged, "QPointingDevice_grabChanged")
	symaddr(&addr_QPointingDevice_seatNames, "QPointingDevice_seatNames")
	symaddr(&addr_QPointingDevice_devices, "QPointingDevice_devices")
	symaddr(&addr_QPointingDevice_primaryKeyboard, "QPointingDevice_primaryKeyboard")
	symaddr(&addr_QPointingDevice_connect, "QPointingDevice_connect")
	symaddr(&addr_QPointingDevice_connect2, "QPointingDevice_connect2")
	symaddr(&addr_QPointingDevice_sender, "QPointingDevice_sender")
	symaddr(&addr_QPointingDevice_senderSignalIndex, "QPointingDevice_senderSignalIndex")
	symaddr(&addr_QPointingDevice_receivers, "QPointingDevice_receivers")
	symaddr(&addr_QPointingDevice_isSignalConnected, "QPointingDevice_isSignalConnected")
	symaddr(&addr_QPointingDevice_tr2, "QPointingDevice_tr2")
	symaddr(&addr_QPointingDevice_tr3, "QPointingDevice_tr3")
	symaddr(&addr_QPointingDevice_primaryPointingDevice2, "QPointingDevice_primaryPointingDevice2")
	symaddr(&addr_QPointingDevice_primaryKeyboard2, "QPointingDevice_primaryKeyboard2")
	symaddr(&addr_QPointingDevice_connect3, "QPointingDevice_connect3")
	symaddr(&addr_QPointingDevice_connect4, "QPointingDevice_connect4")
	symaddr(&addr_QPointingDevice_metaObject_default, "QPointingDevice_metaObject_default")
	symaddr(&addr_QPointingDevice_event_default, "QPointingDevice_event_default")
	symaddr(&addr_QPointingDevice_eventFilter_default, "QPointingDevice_eventFilter_default")
	symaddr(&addr_QPointingDevice_timerEvent_default, "QPointingDevice_timerEvent_default")
	symaddr(&addr_QPointingDevice_childEvent_default, "QPointingDevice_childEvent_default")
	symaddr(&addr_QPointingDevice_customEvent_default, "QPointingDevice_customEvent_default")
	symaddr(&addr_QPointingDevice_connectNotify_default, "QPointingDevice_connectNotify_default")
	symaddr(&addr_QPointingDevice_disconnectNotify_default, "QPointingDevice_disconnectNotify_default")
	reg(&pg_QPointingDevice_connect_grabChanged, "QPointingDevice_connect_grabChanged")
	reg(&pg_QPointingDevice_connect_availableVirtualGeometryChanged, "QPointingDevice_connect_availableVirtualGeometryChanged")
	reg(&pg_QPointingDevice_connect_capabilitiesChanged, "QPointingDevice_connect_capabilitiesChanged")
	reg(&pg_QPointingDevice_connect_destroyed, "QPointingDevice_connect_destroyed")
	reg(&pg_QPointingDevice_connect_destroyed2, "QPointingDevice_connect_destroyed2")
	reg(&pg_QPointingDevice_connect_objectNameChanged, "QPointingDevice_connect_objectNameChanged")
	reg(&pg_QPointingDevice__disconnect, "QPointingDevice__disconnect")
	reg(&pg_QPointingDevice__set_override, "QPointingDevice__set_override")
	reg(&pg_QPointingDevice_delete, "QPointingDevice_delete")
}

type QPointingDevice struct {
	h uintptr
	*QInputDevice
}

func bindQPointingDevice(h uintptr) *QPointingDevice {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPointingDevice_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPointingDevice{h: h, QInputDevice: &QInputDevice{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPointingDevice(h uintptr) *QPointingDevice { return bindQPointingDevice(h) }

func (this *QPointingDevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPointingDevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPointingDevice(h unsafe.Pointer) *QPointingDevice { return bindQPointingDevice(uintptr(h)) }
func UnsafeNewQPointingDeviceOwned(h unsafe.Pointer) *QPointingDevice { return newownedQPointingDevice(uintptr(h)) }

func (this *QPointingDevice) UnsafeRelease() {
	this.h = 0
	if this.QInputDevice != nil {
		this.QInputDevice.UnsafeRelease()
	}
}

func NewQPointingDevice() *QPointingDevice {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new, uintptr(__goid))
	__o := newownedQPointingDevice(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func NewQPointingDevice2(name string, systemId int64, devType QInputDevice__DeviceType, pType QPointingDevice__PointerType, caps QInputDevice__Capability, maxPoints int, buttonCount int) *QPointingDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new2, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(devType), uintptr(pType), uintptr(caps), uintptr(maxPoints), uintptr(buttonCount))
	__o := newownedQPointingDevice(__r)
	runtime.KeepAlive(name)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func NewQPointingDevice3(parent *QObject) *QPointingDevice {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQPointingDevice(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func NewQPointingDevice4(name string, systemId int64, devType QInputDevice__DeviceType, pType QPointingDevice__PointerType, caps QInputDevice__Capability, maxPoints int, buttonCount int, seatName string) *QPointingDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new4, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(devType), uintptr(pType), uintptr(caps), uintptr(maxPoints), uintptr(buttonCount), seatName_p, uintptr(len(seatName)))
	__o := newownedQPointingDevice(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(seatName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func NewQPointingDevice5(name string, systemId int64, devType QInputDevice__DeviceType, pType QPointingDevice__PointerType, caps QInputDevice__Capability, maxPoints int, buttonCount int, seatName string, uniqueId *QPointingDeviceUniqueId) *QPointingDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new5, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(devType), uintptr(pType), uintptr(caps), uintptr(maxPoints), uintptr(buttonCount), seatName_p, uintptr(len(seatName)), uintptr(uniqueId.UnsafePointer()))
	__o := newownedQPointingDevice(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(seatName)
	runtime.KeepAlive(uniqueId)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func NewQPointingDevice6(name string, systemId int64, devType QInputDevice__DeviceType, pType QPointingDevice__PointerType, caps QInputDevice__Capability, maxPoints int, buttonCount int, seatName string, uniqueId *QPointingDeviceUniqueId, parent *QObject) *QPointingDevice {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_new6, uintptr(__goid), name_p, uintptr(len(name)), uintptr(systemId), uintptr(devType), uintptr(pType), uintptr(caps), uintptr(maxPoints), uintptr(buttonCount), seatName_p, uintptr(len(seatName)), uintptr(uniqueId.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQPointingDevice(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(seatName)
	runtime.KeepAlive(uniqueId)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPointingDevice_virtIdx, QPointingDevice_enable)
	return __o
}

func QPointingDevice_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPointingDevice) SetType(devType QInputDevice__DeviceType) {
	purego.SyscallN(addr_QPointingDevice_setType, this.h, uintptr(devType))
	runtime.KeepAlive(this)
}

func (this *QPointingDevice) SetCapabilities(caps QInputDevice__Capability) {
	purego.SyscallN(addr_QPointingDevice_setCapabilities, this.h, uintptr(caps))
	runtime.KeepAlive(this)
}

func (this *QPointingDevice) SetMaximumTouchPoints(c int) {
	purego.SyscallN(addr_QPointingDevice_setMaximumTouchPoints, this.h, uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QPointingDevice) PointerType() QPointingDevice__PointerType {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_pointerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPointingDevice__PointerType(__r)
}

func (this *QPointingDevice) MaximumPoints() int {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_maximumPoints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPointingDevice) ButtonCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_buttonCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPointingDevice) UniqueId() *QPointingDeviceUniqueId {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_uniqueId, this.h)
	runtime.KeepAlive(this)
	return newownedQPointingDeviceUniqueId(__r)
}

func QPointingDevice_PrimaryPointingDevice() *QPointingDevice {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_primaryPointingDevice)
	return bindQPointingDevice(__r)
}

func (this *QPointingDevice) OperatorEqual(other *QPointingDevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QPointingDevice) GrabChanged(grabber *QObject, transition QPointingDevice__GrabTransition, event *QPointerEvent, point *QEventPoint) {
	purego.SyscallN(addr_QPointingDevice_grabChanged, this.h, uintptr(grabber.UnsafePointer()), uintptr(transition), uintptr(event.UnsafePointer()), uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(grabber)
	runtime.KeepAlive(event)
	runtime.KeepAlive(point)
}

func QPointingDevice_SeatNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_seatNames, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPointingDevice_Devices() []*QInputDevice {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_devices, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QInputDevice, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQInputDevice(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPointingDevice_PrimaryKeyboard() *QInputDevice {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_primaryKeyboard)
	return bindQInputDevice(__r)
}

func QPointingDevice_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPointingDevice_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPointingDevice) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QPointingDevice) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPointingDevice) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QPointingDevice) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QPointingDevice_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPointingDevice_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPointingDevice_PrimaryPointingDeviceWithSeatName(seatName string) *QPointingDevice {
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_primaryPointingDevice2, seatName_p, uintptr(len(seatName)))
	runtime.KeepAlive(seatName)
	return bindQPointingDevice(__r)
}

func QPointingDevice_PrimaryKeyboardWithSeatName(seatName string) *QInputDevice {
	var seatName_p uintptr
	if len(seatName) > 0 {
		seatName_p = uintptr(unsafe.Pointer(unsafe.StringData(seatName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_primaryKeyboard2, seatName_p, uintptr(len(seatName)))
	runtime.KeepAlive(seatName)
	return bindQInputDevice(__r)
}

func QPointingDevice_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPointingDevice_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPointingDevice) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QPointingDevice_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QPointingDevice) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPointingDevice) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDevice_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QPointingDevice) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QPointingDevice_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPointingDevice) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QPointingDevice_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPointingDevice) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QPointingDevice_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QPointingDevice) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPointingDevice_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPointingDevice) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QPointingDevice_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QPointingDevice) OnGrabChanged(fn func(*QObject, QPointingDevice__GrabTransition, *QPointerEvent, *QEventPoint)) int64 {
	id := registerCallback(func(grabber uintptr, transition int32, event uintptr, point uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(grabber)), QPointingDevice__GrabTransition(transition), UnsafeNewQPointerEvent(unsafe.Pointer(event)), UnsafeNewQEventPoint(unsafe.Pointer(point))) })
	pg_QPointingDevice_connect_grabChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) OnAvailableVirtualGeometryChanged(fn func(*QRect)) int64 {
	id := registerCallback(func(area uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(area))) })
	pg_QPointingDevice_connect_availableVirtualGeometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) OnCapabilitiesChanged(fn func(QInputDevice__Capability)) int64 {
	id := registerCallback(func(capabilities int32) { fn(QInputDevice__Capability(capabilities)) })
	pg_QPointingDevice_connect_capabilitiesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPointingDevice_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPointingDevice_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPointingDevice_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPointingDevice) Disconnect(handle int64) {
	pg_QPointingDevice__disconnect(handle)
	unregisterCallback(handle)
}

func QPointingDevice_virtIdx(slot int64) int32 {
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

func QPointingDevice_enable(h uintptr, idx int32) { pg_QPointingDevice__set_override(h, idx) }

func (this *QPointingDevice) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QPointingDevice) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QPointingDevice) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QPointingDevice) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QPointingDevice) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QPointingDevice) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QPointingDevice) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QPointingDevice) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QPointingDevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPointingDevice_delete(this.h)
	}
	this.UnsafeRelease()
}
