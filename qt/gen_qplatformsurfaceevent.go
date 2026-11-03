package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPlatformSurfaceEvent_base func(uintptr) uintptr
	addr_QPlatformSurfaceEvent_new uintptr
	addr_QPlatformSurfaceEvent_surfaceEventType uintptr
	addr_QPlatformSurfaceEvent_registerEventType uintptr
	addr_QPlatformSurfaceEvent_registerEventType2 uintptr
	pg_QPlatformSurfaceEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QPlatformSurfaceEvent_base, "QPlatformSurfaceEvent_base")
	symaddr(&addr_QPlatformSurfaceEvent_new, "QPlatformSurfaceEvent_new")
	symaddr(&addr_QPlatformSurfaceEvent_surfaceEventType, "QPlatformSurfaceEvent_surfaceEventType")
	symaddr(&addr_QPlatformSurfaceEvent_registerEventType, "QPlatformSurfaceEvent_registerEventType")
	symaddr(&addr_QPlatformSurfaceEvent_registerEventType2, "QPlatformSurfaceEvent_registerEventType2")
	reg(&pg_QPlatformSurfaceEvent_delete, "QPlatformSurfaceEvent_delete")
}

type QPlatformSurfaceEvent struct {
	h uintptr
	*QEvent
}

func bindQPlatformSurfaceEvent(h uintptr) *QPlatformSurfaceEvent {
	if h == 0 {
		return nil
	}
	return &QPlatformSurfaceEvent{h: h, QEvent: bindQEvent(pg_QPlatformSurfaceEvent_base(h))}
}

func newownedQPlatformSurfaceEvent(h uintptr) *QPlatformSurfaceEvent {
	o := bindQPlatformSurfaceEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPlatformSurfaceEvent).Delete)
	}
	return o
}

func (this *QPlatformSurfaceEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPlatformSurfaceEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPlatformSurfaceEvent(h unsafe.Pointer) *QPlatformSurfaceEvent { return bindQPlatformSurfaceEvent(uintptr(h)) }
func UnsafeNewQPlatformSurfaceEventOwned(h unsafe.Pointer) *QPlatformSurfaceEvent { return newownedQPlatformSurfaceEvent(uintptr(h)) }

func (this *QPlatformSurfaceEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQPlatformSurfaceEvent(surfaceEventType QPlatformSurfaceEvent__SurfaceEventType) *QPlatformSurfaceEvent {
	__r, _, _ := purego.SyscallN(addr_QPlatformSurfaceEvent_new, uintptr(surfaceEventType))
	__o := newownedQPlatformSurfaceEvent(__r)
	return __o
}

func (this *QPlatformSurfaceEvent) SurfaceEventType() QPlatformSurfaceEvent__SurfaceEventType {
	__r0, _, _ := purego.SyscallN(addr_QPlatformSurfaceEvent_surfaceEventType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPlatformSurfaceEvent__SurfaceEventType(__r)
}

func QPlatformSurfaceEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QPlatformSurfaceEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QPlatformSurfaceEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QPlatformSurfaceEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QPlatformSurfaceEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPlatformSurfaceEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
