package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QHoverEvent_chain uintptr
	addr_QHoverEvent_new uintptr
	addr_QHoverEvent_new2 uintptr
	addr_QHoverEvent_new3 uintptr
	addr_QHoverEvent_new4 uintptr
	addr_QHoverEvent_new5 uintptr
	addr_QHoverEvent_new6 uintptr
	addr_QHoverEvent_pos uintptr
	addr_QHoverEvent_posF uintptr
	addr_QHoverEvent_oldPos uintptr
	addr_QHoverEvent_oldPosF uintptr
	addr_QHoverEvent_registerEventType uintptr
	addr_QHoverEvent_registerEventType2 uintptr
	pg_QHoverEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QHoverEvent_chain, "QHoverEvent_chain")
	symaddr(&addr_QHoverEvent_new, "QHoverEvent_new")
	symaddr(&addr_QHoverEvent_new2, "QHoverEvent_new2")
	symaddr(&addr_QHoverEvent_new3, "QHoverEvent_new3")
	symaddr(&addr_QHoverEvent_new4, "QHoverEvent_new4")
	symaddr(&addr_QHoverEvent_new5, "QHoverEvent_new5")
	symaddr(&addr_QHoverEvent_new6, "QHoverEvent_new6")
	symaddr(&addr_QHoverEvent_pos, "QHoverEvent_pos")
	symaddr(&addr_QHoverEvent_posF, "QHoverEvent_posF")
	symaddr(&addr_QHoverEvent_oldPos, "QHoverEvent_oldPos")
	symaddr(&addr_QHoverEvent_oldPosF, "QHoverEvent_oldPosF")
	symaddr(&addr_QHoverEvent_registerEventType, "QHoverEvent_registerEventType")
	symaddr(&addr_QHoverEvent_registerEventType2, "QHoverEvent_registerEventType2")
	reg(&pg_QHoverEvent_delete, "QHoverEvent_delete")
}

type QHoverEvent struct {
	h uintptr
	*QSinglePointEvent
}

func bindQHoverEvent(h uintptr) *QHoverEvent {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QHoverEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QHoverEvent{h: h, QSinglePointEvent: &QSinglePointEvent{h: __c[0], QPointerEvent: &QPointerEvent{h: __c[1], QInputEvent: &QInputEvent{h: __c[2], QEvent: &QEvent{h: __c[3]}}}}}
}

func newownedQHoverEvent(h uintptr) *QHoverEvent {
	o := bindQHoverEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QHoverEvent).Delete)
	}
	return o
}

func (this *QHoverEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHoverEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHoverEvent(h unsafe.Pointer) *QHoverEvent { return bindQHoverEvent(uintptr(h)) }
func UnsafeNewQHoverEventOwned(h unsafe.Pointer) *QHoverEvent { return newownedQHoverEvent(uintptr(h)) }

func (this *QHoverEvent) UnsafeRelease() {
	this.h = 0
	if this.QSinglePointEvent != nil {
		this.QSinglePointEvent.UnsafeRelease()
	}
}

func NewQHoverEvent(type_ QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new, uintptr(type_), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(oldPos.UnsafePointer()))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(oldPos)
	return __o
}

func NewQHoverEvent2(type_ QEvent__Type, pos *QPointF, oldPos *QPointF) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new2, uintptr(type_), uintptr(pos.UnsafePointer()), uintptr(oldPos.UnsafePointer()))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(oldPos)
	return __o
}

func NewQHoverEvent3(type_ QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new3, uintptr(type_), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(oldPos.UnsafePointer()), uintptr(modifiers))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(oldPos)
	return __o
}

func NewQHoverEvent4(type_ QEvent__Type, scenePos *QPointF, globalPos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new4, uintptr(type_), uintptr(scenePos.UnsafePointer()), uintptr(globalPos.UnsafePointer()), uintptr(oldPos.UnsafePointer()), uintptr(modifiers), uintptr(device.UnsafePointer()))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(scenePos)
	runtime.KeepAlive(globalPos)
	runtime.KeepAlive(oldPos)
	runtime.KeepAlive(device)
	return __o
}

func NewQHoverEvent5(type_ QEvent__Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new5, uintptr(type_), uintptr(pos.UnsafePointer()), uintptr(oldPos.UnsafePointer()), uintptr(modifiers))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(oldPos)
	return __o
}

func NewQHoverEvent6(type_ QEvent__Type, pos *QPointF, oldPos *QPointF, modifiers KeyboardModifier, device *QPointingDevice) *QHoverEvent {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_new6, uintptr(type_), uintptr(pos.UnsafePointer()), uintptr(oldPos.UnsafePointer()), uintptr(modifiers), uintptr(device.UnsafePointer()))
	__o := newownedQHoverEvent(__r)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(oldPos)
	runtime.KeepAlive(device)
	return __o
}

func (this *QHoverEvent) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QHoverEvent) PosF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_posF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QHoverEvent) OldPos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_oldPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QHoverEvent) OldPosF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QHoverEvent_oldPosF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func QHoverEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QHoverEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QHoverEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QHoverEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QHoverEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHoverEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
