package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QScreenOrientationChangeEvent_base func(uintptr) uintptr
	addr_QScreenOrientationChangeEvent_new uintptr
	addr_QScreenOrientationChangeEvent_screen uintptr
	addr_QScreenOrientationChangeEvent_orientation uintptr
	addr_QScreenOrientationChangeEvent_registerEventType uintptr
	addr_QScreenOrientationChangeEvent_registerEventType2 uintptr
	pg_QScreenOrientationChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QScreenOrientationChangeEvent_base, "QScreenOrientationChangeEvent_base")
	symaddr(&addr_QScreenOrientationChangeEvent_new, "QScreenOrientationChangeEvent_new")
	symaddr(&addr_QScreenOrientationChangeEvent_screen, "QScreenOrientationChangeEvent_screen")
	symaddr(&addr_QScreenOrientationChangeEvent_orientation, "QScreenOrientationChangeEvent_orientation")
	symaddr(&addr_QScreenOrientationChangeEvent_registerEventType, "QScreenOrientationChangeEvent_registerEventType")
	symaddr(&addr_QScreenOrientationChangeEvent_registerEventType2, "QScreenOrientationChangeEvent_registerEventType2")
	reg(&pg_QScreenOrientationChangeEvent_delete, "QScreenOrientationChangeEvent_delete")
}

type QScreenOrientationChangeEvent struct {
	h uintptr
	*QEvent
}

func bindQScreenOrientationChangeEvent(h uintptr) *QScreenOrientationChangeEvent {
	if h == 0 {
		return nil
	}
	return &QScreenOrientationChangeEvent{h: h, QEvent: bindQEvent(pg_QScreenOrientationChangeEvent_base(h))}
}

func newownedQScreenOrientationChangeEvent(h uintptr) *QScreenOrientationChangeEvent {
	o := bindQScreenOrientationChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QScreenOrientationChangeEvent).Delete)
	}
	return o
}

func (this *QScreenOrientationChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScreenOrientationChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScreenOrientationChangeEvent(h unsafe.Pointer) *QScreenOrientationChangeEvent { return bindQScreenOrientationChangeEvent(uintptr(h)) }
func UnsafeNewQScreenOrientationChangeEventOwned(h unsafe.Pointer) *QScreenOrientationChangeEvent { return newownedQScreenOrientationChangeEvent(uintptr(h)) }

func (this *QScreenOrientationChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQScreenOrientationChangeEvent(screen *QScreen, orientation ScreenOrientation) *QScreenOrientationChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QScreenOrientationChangeEvent_new, uintptr(screen.UnsafePointer()), uintptr(orientation))
	__o := newownedQScreenOrientationChangeEvent(__r)
	runtime.KeepAlive(screen)
	return __o
}

func (this *QScreenOrientationChangeEvent) Screen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QScreenOrientationChangeEvent_screen, this.h)
	runtime.KeepAlive(this)
	return bindQScreen(__r)
}

func (this *QScreenOrientationChangeEvent) Orientation() ScreenOrientation {
	__r0, _, _ := purego.SyscallN(addr_QScreenOrientationChangeEvent_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScreenOrientation(__r)
}

func QScreenOrientationChangeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QScreenOrientationChangeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QScreenOrientationChangeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QScreenOrientationChangeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QScreenOrientationChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScreenOrientationChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
