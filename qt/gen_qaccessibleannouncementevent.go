package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleAnnouncementEvent_base func(uintptr) uintptr
	addr_QAccessibleAnnouncementEvent_new uintptr
	addr_QAccessibleAnnouncementEvent_new2 uintptr
	addr_QAccessibleAnnouncementEvent_message uintptr
	addr_QAccessibleAnnouncementEvent_politeness uintptr
	addr_QAccessibleAnnouncementEvent_setPoliteness uintptr
	pg_QAccessibleAnnouncementEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleAnnouncementEvent_base, "QAccessibleAnnouncementEvent_base")
	symaddr(&addr_QAccessibleAnnouncementEvent_new, "QAccessibleAnnouncementEvent_new")
	symaddr(&addr_QAccessibleAnnouncementEvent_new2, "QAccessibleAnnouncementEvent_new2")
	symaddr(&addr_QAccessibleAnnouncementEvent_message, "QAccessibleAnnouncementEvent_message")
	symaddr(&addr_QAccessibleAnnouncementEvent_politeness, "QAccessibleAnnouncementEvent_politeness")
	symaddr(&addr_QAccessibleAnnouncementEvent_setPoliteness, "QAccessibleAnnouncementEvent_setPoliteness")
	reg(&pg_QAccessibleAnnouncementEvent_delete, "QAccessibleAnnouncementEvent_delete")
}

type QAccessibleAnnouncementEvent struct {
	h uintptr
	*QAccessibleEvent
}

func bindQAccessibleAnnouncementEvent(h uintptr) *QAccessibleAnnouncementEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleAnnouncementEvent{h: h, QAccessibleEvent: bindQAccessibleEvent(pg_QAccessibleAnnouncementEvent_base(h))}
}

func newownedQAccessibleAnnouncementEvent(h uintptr) *QAccessibleAnnouncementEvent {
	o := bindQAccessibleAnnouncementEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleAnnouncementEvent).Delete)
	}
	return o
}

func (this *QAccessibleAnnouncementEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleAnnouncementEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleAnnouncementEvent(h unsafe.Pointer) *QAccessibleAnnouncementEvent { return bindQAccessibleAnnouncementEvent(uintptr(h)) }
func UnsafeNewQAccessibleAnnouncementEventOwned(h unsafe.Pointer) *QAccessibleAnnouncementEvent { return newownedQAccessibleAnnouncementEvent(uintptr(h)) }

func (this *QAccessibleAnnouncementEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleEvent != nil {
		this.QAccessibleEvent.UnsafeRelease()
	}
}

func NewQAccessibleAnnouncementEvent(object *QObject, message string) *QAccessibleAnnouncementEvent {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleAnnouncementEvent_new, uintptr(object.UnsafePointer()), message_p, uintptr(len(message)))
	__o := newownedQAccessibleAnnouncementEvent(__r)
	runtime.KeepAlive(object)
	runtime.KeepAlive(message)
	return __o
}

func NewQAccessibleAnnouncementEvent2(iface *QAccessibleInterface, message string) *QAccessibleAnnouncementEvent {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleAnnouncementEvent_new2, uintptr(iface.UnsafePointer()), message_p, uintptr(len(message)))
	__o := newownedQAccessibleAnnouncementEvent(__r)
	runtime.KeepAlive(iface)
	runtime.KeepAlive(message)
	return __o
}

func (this *QAccessibleAnnouncementEvent) Message() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleAnnouncementEvent_message, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleAnnouncementEvent) Politeness() QAccessible__AnnouncementPoliteness {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleAnnouncementEvent_politeness, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAccessible__AnnouncementPoliteness(__r)
}

func (this *QAccessibleAnnouncementEvent) SetPoliteness(politeness QAccessible__AnnouncementPoliteness) {
	purego.SyscallN(addr_QAccessibleAnnouncementEvent_setPoliteness, this.h, uintptr(politeness))
	runtime.KeepAlive(this)
}

func (this *QAccessibleAnnouncementEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleAnnouncementEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
