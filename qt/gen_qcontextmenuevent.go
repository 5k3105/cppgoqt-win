package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QContextMenuEvent_chain uintptr
	addr_QContextMenuEvent_new uintptr
	addr_QContextMenuEvent_new2 uintptr
	addr_QContextMenuEvent_new3 uintptr
	addr_QContextMenuEvent_x uintptr
	addr_QContextMenuEvent_y uintptr
	addr_QContextMenuEvent_globalX uintptr
	addr_QContextMenuEvent_globalY uintptr
	addr_QContextMenuEvent_pos uintptr
	addr_QContextMenuEvent_globalPos uintptr
	addr_QContextMenuEvent_reason uintptr
	addr_QContextMenuEvent_registerEventType uintptr
	addr_QContextMenuEvent_registerEventType2 uintptr
	pg_QContextMenuEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QContextMenuEvent_chain, "QContextMenuEvent_chain")
	symaddr(&addr_QContextMenuEvent_new, "QContextMenuEvent_new")
	symaddr(&addr_QContextMenuEvent_new2, "QContextMenuEvent_new2")
	symaddr(&addr_QContextMenuEvent_new3, "QContextMenuEvent_new3")
	symaddr(&addr_QContextMenuEvent_x, "QContextMenuEvent_x")
	symaddr(&addr_QContextMenuEvent_y, "QContextMenuEvent_y")
	symaddr(&addr_QContextMenuEvent_globalX, "QContextMenuEvent_globalX")
	symaddr(&addr_QContextMenuEvent_globalY, "QContextMenuEvent_globalY")
	symaddr(&addr_QContextMenuEvent_pos, "QContextMenuEvent_pos")
	symaddr(&addr_QContextMenuEvent_globalPos, "QContextMenuEvent_globalPos")
	symaddr(&addr_QContextMenuEvent_reason, "QContextMenuEvent_reason")
	symaddr(&addr_QContextMenuEvent_registerEventType, "QContextMenuEvent_registerEventType")
	symaddr(&addr_QContextMenuEvent_registerEventType2, "QContextMenuEvent_registerEventType2")
	reg(&pg_QContextMenuEvent_delete, "QContextMenuEvent_delete")
}

type QContextMenuEvent struct {
	h uintptr
	*QInputEvent
}

func bindQContextMenuEvent(h uintptr) *QContextMenuEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QContextMenuEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QContextMenuEvent{h: h, QInputEvent: &QInputEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQContextMenuEvent(h uintptr) *QContextMenuEvent {
	o := bindQContextMenuEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QContextMenuEvent).Delete)
	}
	return o
}

func (this *QContextMenuEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QContextMenuEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQContextMenuEvent(h unsafe.Pointer) *QContextMenuEvent { return bindQContextMenuEvent(uintptr(h)) }
func UnsafeNewQContextMenuEventOwned(h unsafe.Pointer) *QContextMenuEvent { return newownedQContextMenuEvent(uintptr(h)) }

func (this *QContextMenuEvent) UnsafeRelease() {
	this.h = 0
	if this.QInputEvent != nil {
		this.QInputEvent.UnsafeRelease()
	}
}

func NewQContextMenuEvent(reason QContextMenuEvent__Reason, pos *QPoint, globalPos *QPoint) *QContextMenuEvent {
	__r, _, _ := purego.SyscallN(addr_QContextMenuEvent_new, uintptr(reason), uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()))
	__o := newownedQContextMenuEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	return __o
}

func NewQContextMenuEvent2(reason QContextMenuEvent__Reason, pos *QPoint) *QContextMenuEvent {
	__r, _, _ := purego.SyscallN(addr_QContextMenuEvent_new2, uintptr(reason), uintptr(pos.UnsafePointer()))
	__o := newownedQContextMenuEvent(__r)
	runtime.KeepAlive(pos)
	return __o
}

func NewQContextMenuEvent3(reason QContextMenuEvent__Reason, pos *QPoint, globalPos *QPoint, modifiers KeyboardModifier) *QContextMenuEvent {
	__r, _, _ := purego.SyscallN(addr_QContextMenuEvent_new3, uintptr(reason), uintptr(pos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(modifiers))
	__o := newownedQContextMenuEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(globalPos)
	return __o
}

func (this *QContextMenuEvent) X() int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QContextMenuEvent) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QContextMenuEvent) GlobalX() int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_globalX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QContextMenuEvent) GlobalY() int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_globalY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QContextMenuEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QContextMenuEvent_pos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QContextMenuEvent) GlobalPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QContextMenuEvent_globalPos, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QContextMenuEvent) Reason() QContextMenuEvent__Reason {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_reason, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QContextMenuEvent__Reason(__r)
}

func QContextMenuEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QContextMenuEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QContextMenuEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QContextMenuEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QContextMenuEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
