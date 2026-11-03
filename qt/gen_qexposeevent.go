package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QExposeEvent_base func(uintptr) uintptr
	addr_QExposeEvent_new uintptr
	addr_QExposeEvent_region uintptr
	addr_QExposeEvent_registerEventType uintptr
	addr_QExposeEvent_registerEventType2 uintptr
	pg_QExposeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QExposeEvent_base, "QExposeEvent_base")
	symaddr(&addr_QExposeEvent_new, "QExposeEvent_new")
	symaddr(&addr_QExposeEvent_region, "QExposeEvent_region")
	symaddr(&addr_QExposeEvent_registerEventType, "QExposeEvent_registerEventType")
	symaddr(&addr_QExposeEvent_registerEventType2, "QExposeEvent_registerEventType2")
	reg(&pg_QExposeEvent_delete, "QExposeEvent_delete")
}

type QExposeEvent struct {
	h uintptr
	*QEvent
}

func bindQExposeEvent(h uintptr) *QExposeEvent {
	if h == 0 {
		return nil
	}
	return &QExposeEvent{h: h, QEvent: bindQEvent(pg_QExposeEvent_base(h))}
}

func newownedQExposeEvent(h uintptr) *QExposeEvent {
	o := bindQExposeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QExposeEvent).Delete)
	}
	return o
}

func (this *QExposeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QExposeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQExposeEvent(h unsafe.Pointer) *QExposeEvent { return bindQExposeEvent(uintptr(h)) }
func UnsafeNewQExposeEventOwned(h unsafe.Pointer) *QExposeEvent { return newownedQExposeEvent(uintptr(h)) }

func (this *QExposeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQExposeEvent(m_region *QRegion) *QExposeEvent {
	__r, _, _ := purego.SyscallN(addr_QExposeEvent_new, uintptr(m_region.UnsafePointer()))
	__o := newownedQExposeEvent(__r)
	runtime.KeepAlive(m_region)
	return __o
}

func (this *QExposeEvent) Region() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QExposeEvent_region, this.h)
	runtime.KeepAlive(this)
	return bindQRegion(__r)
}

func QExposeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QExposeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QExposeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QExposeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QExposeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QExposeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
