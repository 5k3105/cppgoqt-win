package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPaintEvent_base func(uintptr) uintptr
	addr_QPaintEvent_new uintptr
	addr_QPaintEvent_new2 uintptr
	addr_QPaintEvent_rect uintptr
	addr_QPaintEvent_region uintptr
	addr_QPaintEvent_registerEventType uintptr
	addr_QPaintEvent_registerEventType2 uintptr
	pg_QPaintEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QPaintEvent_base, "QPaintEvent_base")
	symaddr(&addr_QPaintEvent_new, "QPaintEvent_new")
	symaddr(&addr_QPaintEvent_new2, "QPaintEvent_new2")
	symaddr(&addr_QPaintEvent_rect, "QPaintEvent_rect")
	symaddr(&addr_QPaintEvent_region, "QPaintEvent_region")
	symaddr(&addr_QPaintEvent_registerEventType, "QPaintEvent_registerEventType")
	symaddr(&addr_QPaintEvent_registerEventType2, "QPaintEvent_registerEventType2")
	reg(&pg_QPaintEvent_delete, "QPaintEvent_delete")
}

type QPaintEvent struct {
	h uintptr
	*QEvent
}

func bindQPaintEvent(h uintptr) *QPaintEvent {
	if h == 0 {
		return nil
	}
	return &QPaintEvent{h: h, QEvent: bindQEvent(pg_QPaintEvent_base(h))}
}

func newownedQPaintEvent(h uintptr) *QPaintEvent {
	o := bindQPaintEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPaintEvent).Delete)
	}
	return o
}

func (this *QPaintEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPaintEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPaintEvent(h unsafe.Pointer) *QPaintEvent { return bindQPaintEvent(uintptr(h)) }
func UnsafeNewQPaintEventOwned(h unsafe.Pointer) *QPaintEvent { return newownedQPaintEvent(uintptr(h)) }

func (this *QPaintEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQPaintEvent(paintRegion *QRegion) *QPaintEvent {
	__r, _, _ := purego.SyscallN(addr_QPaintEvent_new, uintptr(paintRegion.UnsafePointer()))
	__o := newownedQPaintEvent(__r)
	runtime.KeepAlive(paintRegion)
	return __o
}

func NewQPaintEvent2(paintRect *QRect) *QPaintEvent {
	__r, _, _ := purego.SyscallN(addr_QPaintEvent_new2, uintptr(paintRect.UnsafePointer()))
	__o := newownedQPaintEvent(__r)
	runtime.KeepAlive(paintRect)
	return __o
}

func (this *QPaintEvent) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPaintEvent_rect, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QPaintEvent) Region() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QPaintEvent_region, this.h)
	runtime.KeepAlive(this)
	return bindQRegion(__r)
}

func QPaintEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QPaintEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QPaintEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QPaintEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPaintEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
