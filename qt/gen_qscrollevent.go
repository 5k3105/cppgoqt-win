package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QScrollEvent_base func(uintptr) uintptr
	addr_QScrollEvent_new uintptr
	addr_QScrollEvent_contentPos uintptr
	addr_QScrollEvent_overshootDistance uintptr
	addr_QScrollEvent_scrollState uintptr
	addr_QScrollEvent_registerEventType uintptr
	addr_QScrollEvent_registerEventType2 uintptr
	pg_QScrollEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QScrollEvent_base, "QScrollEvent_base")
	symaddr(&addr_QScrollEvent_new, "QScrollEvent_new")
	symaddr(&addr_QScrollEvent_contentPos, "QScrollEvent_contentPos")
	symaddr(&addr_QScrollEvent_overshootDistance, "QScrollEvent_overshootDistance")
	symaddr(&addr_QScrollEvent_scrollState, "QScrollEvent_scrollState")
	symaddr(&addr_QScrollEvent_registerEventType, "QScrollEvent_registerEventType")
	symaddr(&addr_QScrollEvent_registerEventType2, "QScrollEvent_registerEventType2")
	reg(&pg_QScrollEvent_delete, "QScrollEvent_delete")
}

type QScrollEvent struct {
	h uintptr
	*QEvent
}

func bindQScrollEvent(h uintptr) *QScrollEvent {
	if h == 0 {
		return nil
	}
	return &QScrollEvent{h: h, QEvent: bindQEvent(pg_QScrollEvent_base(h))}
}

func newownedQScrollEvent(h uintptr) *QScrollEvent {
	o := bindQScrollEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QScrollEvent).Delete)
	}
	return o
}

func (this *QScrollEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScrollEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScrollEvent(h unsafe.Pointer) *QScrollEvent { return bindQScrollEvent(uintptr(h)) }
func UnsafeNewQScrollEventOwned(h unsafe.Pointer) *QScrollEvent { return newownedQScrollEvent(uintptr(h)) }

func (this *QScrollEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQScrollEvent(contentPos *QPointF, overshoot *QPointF, scrollState QScrollEvent__ScrollState) *QScrollEvent {
	__r, _, _ := purego.SyscallN(addr_QScrollEvent_new, uintptr(contentPos.UnsafePointer()), uintptr(overshoot.UnsafePointer()), uintptr(scrollState))
	__o := newownedQScrollEvent(__r)
	runtime.KeepAlive(contentPos)
	runtime.KeepAlive(overshoot)
	return __o
}

func (this *QScrollEvent) ContentPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScrollEvent_contentPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScrollEvent) OvershootDistance() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScrollEvent_overshootDistance, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScrollEvent) ScrollState() QScrollEvent__ScrollState {
	__r0, _, _ := purego.SyscallN(addr_QScrollEvent_scrollState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QScrollEvent__ScrollState(__r)
}

func QScrollEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QScrollEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QScrollEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScrollEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
