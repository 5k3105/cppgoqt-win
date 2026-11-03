package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QInputEvent_base func(uintptr) uintptr
	addr_QInputEvent_new uintptr
	addr_QInputEvent_new2 uintptr
	addr_QInputEvent_device uintptr
	addr_QInputEvent_deviceType uintptr
	addr_QInputEvent_modifiers uintptr
	addr_QInputEvent_setModifiers uintptr
	addr_QInputEvent_timestamp uintptr
	addr_QInputEvent_setTimestamp uintptr
	addr_QInputEvent_registerEventType uintptr
	addr_QInputEvent_registerEventType2 uintptr
	pg_QInputEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QInputEvent_base, "QInputEvent_base")
	symaddr(&addr_QInputEvent_new, "QInputEvent_new")
	symaddr(&addr_QInputEvent_new2, "QInputEvent_new2")
	symaddr(&addr_QInputEvent_device, "QInputEvent_device")
	symaddr(&addr_QInputEvent_deviceType, "QInputEvent_deviceType")
	symaddr(&addr_QInputEvent_modifiers, "QInputEvent_modifiers")
	symaddr(&addr_QInputEvent_setModifiers, "QInputEvent_setModifiers")
	symaddr(&addr_QInputEvent_timestamp, "QInputEvent_timestamp")
	symaddr(&addr_QInputEvent_setTimestamp, "QInputEvent_setTimestamp")
	symaddr(&addr_QInputEvent_registerEventType, "QInputEvent_registerEventType")
	symaddr(&addr_QInputEvent_registerEventType2, "QInputEvent_registerEventType2")
	reg(&pg_QInputEvent_delete, "QInputEvent_delete")
}

type QInputEvent struct {
	h uintptr
	*QEvent
}

func bindQInputEvent(h uintptr) *QInputEvent {
	if h == 0 {
		return nil
	}
	return &QInputEvent{h: h, QEvent: bindQEvent(pg_QInputEvent_base(h))}
}

func newownedQInputEvent(h uintptr) *QInputEvent {
	o := bindQInputEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QInputEvent).Delete)
	}
	return o
}

func (this *QInputEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputEvent(h unsafe.Pointer) *QInputEvent { return bindQInputEvent(uintptr(h)) }
func UnsafeNewQInputEventOwned(h unsafe.Pointer) *QInputEvent { return newownedQInputEvent(uintptr(h)) }

func (this *QInputEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQInputEvent(type_ QEvent__Type, m_dev *QInputDevice) *QInputEvent {
	__r, _, _ := purego.SyscallN(addr_QInputEvent_new, uintptr(type_), uintptr(m_dev.UnsafePointer()))
	__o := newownedQInputEvent(__r)
	runtime.KeepAlive(m_dev)
	return __o
}

func NewQInputEvent2(type_ QEvent__Type, m_dev *QInputDevice, modifiers KeyboardModifier) *QInputEvent {
	__r, _, _ := purego.SyscallN(addr_QInputEvent_new2, uintptr(type_), uintptr(m_dev.UnsafePointer()), uintptr(modifiers))
	__o := newownedQInputEvent(__r)
	runtime.KeepAlive(m_dev)
	return __o
}

func (this *QInputEvent) Device() *QInputDevice {
	__r, _, _ := purego.SyscallN(addr_QInputEvent_device, this.h)
	runtime.KeepAlive(this)
	return bindQInputDevice(__r)
}

func (this *QInputEvent) DeviceType() QInputDevice__DeviceType {
	__r0, _, _ := purego.SyscallN(addr_QInputEvent_deviceType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputDevice__DeviceType(__r)
}

func (this *QInputEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QInputEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QInputEvent) SetModifiers(modifiers KeyboardModifier) {
	purego.SyscallN(addr_QInputEvent_setModifiers, this.h, uintptr(modifiers))
	runtime.KeepAlive(this)
}

func (this *QInputEvent) Timestamp() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QInputEvent_timestamp, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QInputEvent) SetTimestamp(timestamp uint64) {
	purego.SyscallN(addr_QInputEvent_setTimestamp, this.h, uintptr(timestamp))
	runtime.KeepAlive(this)
}

func QInputEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QInputEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QInputEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QInputEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QInputEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
