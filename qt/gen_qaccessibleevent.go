package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleEvent_new uintptr
	addr_QAccessibleEvent_new2 uintptr
	addr_QAccessibleEvent_type uintptr
	addr_QAccessibleEvent_object uintptr
	addr_QAccessibleEvent_uniqueId uintptr
	addr_QAccessibleEvent_setChild uintptr
	addr_QAccessibleEvent_child uintptr
	addr_QAccessibleEvent_accessibleInterface uintptr
	pg_QAccessibleEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleEvent_new, "QAccessibleEvent_new")
	symaddr(&addr_QAccessibleEvent_new2, "QAccessibleEvent_new2")
	symaddr(&addr_QAccessibleEvent_type, "QAccessibleEvent_type")
	symaddr(&addr_QAccessibleEvent_object, "QAccessibleEvent_object")
	symaddr(&addr_QAccessibleEvent_uniqueId, "QAccessibleEvent_uniqueId")
	symaddr(&addr_QAccessibleEvent_setChild, "QAccessibleEvent_setChild")
	symaddr(&addr_QAccessibleEvent_child, "QAccessibleEvent_child")
	symaddr(&addr_QAccessibleEvent_accessibleInterface, "QAccessibleEvent_accessibleInterface")
	reg(&pg_QAccessibleEvent_delete, "QAccessibleEvent_delete")
}

type QAccessibleEvent struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleEvent) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleEvent) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleEvent(h uintptr) *QAccessibleEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleEvent{h: h}
}

func newownedQAccessibleEvent(h uintptr) *QAccessibleEvent {
	o := bindQAccessibleEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleEvent).Delete)
	}
	return o
}

func (this *QAccessibleEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleEvent(h unsafe.Pointer) *QAccessibleEvent { return bindQAccessibleEvent(uintptr(h)) }
func UnsafeNewQAccessibleEventOwned(h unsafe.Pointer) *QAccessibleEvent { return newownedQAccessibleEvent(uintptr(h)) }

func (this *QAccessibleEvent) UnsafeRelease() { this.h = 0 }

func NewQAccessibleEvent(obj *QObject, typ QAccessible__Event) *QAccessibleEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleEvent_new, uintptr(obj.UnsafePointer()), uintptr(typ))
	__o := newownedQAccessibleEvent(__r)
	runtime.KeepAlive(obj)
	return __o
}

func NewQAccessibleEvent2(iface *QAccessibleInterface, typ QAccessible__Event) *QAccessibleEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleEvent_new2, uintptr(iface.UnsafePointer()), uintptr(typ))
	__o := newownedQAccessibleEvent(__r)
	runtime.KeepAlive(iface)
	return __o
}

func (this *QAccessibleEvent) Type() QAccessible__Event {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleEvent_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAccessible__Event(__r)
}

func (this *QAccessibleEvent) Object() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibleEvent_object, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessibleEvent) UniqueId() uint {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleEvent_uniqueId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QAccessibleEvent) SetChild(chld int) {
	purego.SyscallN(addr_QAccessibleEvent_setChild, this.h, uintptr(chld))
	runtime.KeepAlive(this)
}

func (this *QAccessibleEvent) Child() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleEvent_child, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleEvent_accessibleInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
