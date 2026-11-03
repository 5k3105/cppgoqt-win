package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QToolBarChangeEvent_base func(uintptr) uintptr
	addr_QToolBarChangeEvent_new uintptr
	addr_QToolBarChangeEvent_toggle uintptr
	addr_QToolBarChangeEvent_registerEventType uintptr
	addr_QToolBarChangeEvent_registerEventType2 uintptr
	pg_QToolBarChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QToolBarChangeEvent_base, "QToolBarChangeEvent_base")
	symaddr(&addr_QToolBarChangeEvent_new, "QToolBarChangeEvent_new")
	symaddr(&addr_QToolBarChangeEvent_toggle, "QToolBarChangeEvent_toggle")
	symaddr(&addr_QToolBarChangeEvent_registerEventType, "QToolBarChangeEvent_registerEventType")
	symaddr(&addr_QToolBarChangeEvent_registerEventType2, "QToolBarChangeEvent_registerEventType2")
	reg(&pg_QToolBarChangeEvent_delete, "QToolBarChangeEvent_delete")
}

type QToolBarChangeEvent struct {
	h uintptr
	*QEvent
}

func bindQToolBarChangeEvent(h uintptr) *QToolBarChangeEvent {
	if h == 0 {
		return nil
	}
	return &QToolBarChangeEvent{h: h, QEvent: bindQEvent(pg_QToolBarChangeEvent_base(h))}
}

func newownedQToolBarChangeEvent(h uintptr) *QToolBarChangeEvent {
	o := bindQToolBarChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QToolBarChangeEvent).Delete)
	}
	return o
}

func (this *QToolBarChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QToolBarChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQToolBarChangeEvent(h unsafe.Pointer) *QToolBarChangeEvent { return bindQToolBarChangeEvent(uintptr(h)) }
func UnsafeNewQToolBarChangeEventOwned(h unsafe.Pointer) *QToolBarChangeEvent { return newownedQToolBarChangeEvent(uintptr(h)) }

func (this *QToolBarChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQToolBarChangeEvent(t bool) *QToolBarChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QToolBarChangeEvent_new, uintptr(b2i(t)))
	__o := newownedQToolBarChangeEvent(__r)
	return __o
}

func (this *QToolBarChangeEvent) Toggle() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBarChangeEvent_toggle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QToolBarChangeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBarChangeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QToolBarChangeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBarChangeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QToolBarChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QToolBarChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
