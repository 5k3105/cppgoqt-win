package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMoveEvent_base func(uintptr) uintptr
	addr_QMoveEvent_new uintptr
	addr_QMoveEvent_pos uintptr
	addr_QMoveEvent_oldPos uintptr
	addr_QMoveEvent_registerEventType uintptr
	addr_QMoveEvent_registerEventType2 uintptr
	pg_QMoveEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QMoveEvent_base, "QMoveEvent_base")
	symaddr(&addr_QMoveEvent_new, "QMoveEvent_new")
	symaddr(&addr_QMoveEvent_pos, "QMoveEvent_pos")
	symaddr(&addr_QMoveEvent_oldPos, "QMoveEvent_oldPos")
	symaddr(&addr_QMoveEvent_registerEventType, "QMoveEvent_registerEventType")
	symaddr(&addr_QMoveEvent_registerEventType2, "QMoveEvent_registerEventType2")
	reg(&pg_QMoveEvent_delete, "QMoveEvent_delete")
}

type QMoveEvent struct {
	h uintptr
	*QEvent
}

func bindQMoveEvent(h uintptr) *QMoveEvent {
	if h == 0 {
		return nil
	}
	return &QMoveEvent{h: h, QEvent: bindQEvent(pg_QMoveEvent_base(h))}
}

func newownedQMoveEvent(h uintptr) *QMoveEvent {
	o := bindQMoveEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMoveEvent).Delete)
	}
	return o
}

func (this *QMoveEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMoveEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMoveEvent(h unsafe.Pointer) *QMoveEvent { return bindQMoveEvent(uintptr(h)) }
func UnsafeNewQMoveEventOwned(h unsafe.Pointer) *QMoveEvent { return newownedQMoveEvent(uintptr(h)) }

func (this *QMoveEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQMoveEvent(pos *QPoint, oldPos *QPoint) *QMoveEvent {
	__r, _, _ := purego.SyscallN(addr_QMoveEvent_new, uintptr(pos.UnsafePointer()), uintptr(oldPos.UnsafePointer()))
	__o := newownedQMoveEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(oldPos)
	return __o
}

func (this *QMoveEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMoveEvent_pos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QMoveEvent) OldPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMoveEvent_oldPos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func QMoveEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QMoveEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QMoveEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QMoveEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QMoveEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMoveEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
