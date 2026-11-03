package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QHelpEvent_base func(uintptr) uintptr
	addr_QHelpEvent_new uintptr
	addr_QHelpEvent_x uintptr
	addr_QHelpEvent_y uintptr
	addr_QHelpEvent_globalX uintptr
	addr_QHelpEvent_globalY uintptr
	addr_QHelpEvent_pos uintptr
	addr_QHelpEvent_globalPos uintptr
	addr_QHelpEvent_registerEventType uintptr
	addr_QHelpEvent_registerEventType2 uintptr
	pg_QHelpEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QHelpEvent_base, "QHelpEvent_base")
	symaddr(&addr_QHelpEvent_new, "QHelpEvent_new")
	symaddr(&addr_QHelpEvent_x, "QHelpEvent_x")
	symaddr(&addr_QHelpEvent_y, "QHelpEvent_y")
	symaddr(&addr_QHelpEvent_globalX, "QHelpEvent_globalX")
	symaddr(&addr_QHelpEvent_globalY, "QHelpEvent_globalY")
	symaddr(&addr_QHelpEvent_pos, "QHelpEvent_pos")
	symaddr(&addr_QHelpEvent_globalPos, "QHelpEvent_globalPos")
	symaddr(&addr_QHelpEvent_registerEventType, "QHelpEvent_registerEventType")
	symaddr(&addr_QHelpEvent_registerEventType2, "QHelpEvent_registerEventType2")
	reg(&pg_QHelpEvent_delete, "QHelpEvent_delete")
}

type QHelpEvent struct {
	h uintptr
	*QEvent
}

func bindQHelpEvent(h uintptr) *QHelpEvent {
	if h == 0 {
		return nil
	}
	return &QHelpEvent{h: h, QEvent: bindQEvent(pg_QHelpEvent_base(h))}
}

func newownedQHelpEvent(h uintptr) *QHelpEvent {
	o := bindQHelpEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QHelpEvent).Delete)
	}
	return o
}

func (this *QHelpEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHelpEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHelpEvent(h unsafe.Pointer) *QHelpEvent { return bindQHelpEvent(uintptr(h)) }
func UnsafeNewQHelpEventOwned(h unsafe.Pointer) *QHelpEvent { return newownedQHelpEvent(uintptr(h)) }

func (this *QHelpEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQHelpEvent(type_ QEvent__Type, pos *QPoint, globalPos *QPoint) *QHelpEvent {
	__r, _, _ := purego.SyscallN(addr_QHelpEvent_new, uintptr(type_), uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()))
	__o := newownedQHelpEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	return __o
}

func (this *QHelpEvent) X() int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHelpEvent) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHelpEvent) GlobalX() int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_globalX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHelpEvent) GlobalY() int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_globalY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHelpEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QHelpEvent_pos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QHelpEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QHelpEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func QHelpEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QHelpEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QHelpEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QHelpEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHelpEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
