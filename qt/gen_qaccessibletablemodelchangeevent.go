package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAccessibleTableModelChangeEvent_base func(uintptr) uintptr
	addr_QAccessibleTableModelChangeEvent_new uintptr
	addr_QAccessibleTableModelChangeEvent_new2 uintptr
	addr_QAccessibleTableModelChangeEvent_setModelChangeType uintptr
	addr_QAccessibleTableModelChangeEvent_modelChangeType uintptr
	addr_QAccessibleTableModelChangeEvent_setFirstRow uintptr
	addr_QAccessibleTableModelChangeEvent_setFirstColumn uintptr
	addr_QAccessibleTableModelChangeEvent_setLastRow uintptr
	addr_QAccessibleTableModelChangeEvent_setLastColumn uintptr
	addr_QAccessibleTableModelChangeEvent_firstRow uintptr
	addr_QAccessibleTableModelChangeEvent_firstColumn uintptr
	addr_QAccessibleTableModelChangeEvent_lastRow uintptr
	addr_QAccessibleTableModelChangeEvent_lastColumn uintptr
	pg_QAccessibleTableModelChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QAccessibleTableModelChangeEvent_base, "QAccessibleTableModelChangeEvent_base")
	symaddr(&addr_QAccessibleTableModelChangeEvent_new, "QAccessibleTableModelChangeEvent_new")
	symaddr(&addr_QAccessibleTableModelChangeEvent_new2, "QAccessibleTableModelChangeEvent_new2")
	symaddr(&addr_QAccessibleTableModelChangeEvent_setModelChangeType, "QAccessibleTableModelChangeEvent_setModelChangeType")
	symaddr(&addr_QAccessibleTableModelChangeEvent_modelChangeType, "QAccessibleTableModelChangeEvent_modelChangeType")
	symaddr(&addr_QAccessibleTableModelChangeEvent_setFirstRow, "QAccessibleTableModelChangeEvent_setFirstRow")
	symaddr(&addr_QAccessibleTableModelChangeEvent_setFirstColumn, "QAccessibleTableModelChangeEvent_setFirstColumn")
	symaddr(&addr_QAccessibleTableModelChangeEvent_setLastRow, "QAccessibleTableModelChangeEvent_setLastRow")
	symaddr(&addr_QAccessibleTableModelChangeEvent_setLastColumn, "QAccessibleTableModelChangeEvent_setLastColumn")
	symaddr(&addr_QAccessibleTableModelChangeEvent_firstRow, "QAccessibleTableModelChangeEvent_firstRow")
	symaddr(&addr_QAccessibleTableModelChangeEvent_firstColumn, "QAccessibleTableModelChangeEvent_firstColumn")
	symaddr(&addr_QAccessibleTableModelChangeEvent_lastRow, "QAccessibleTableModelChangeEvent_lastRow")
	symaddr(&addr_QAccessibleTableModelChangeEvent_lastColumn, "QAccessibleTableModelChangeEvent_lastColumn")
	reg(&pg_QAccessibleTableModelChangeEvent_delete, "QAccessibleTableModelChangeEvent_delete")
}

type QAccessibleTableModelChangeEvent struct {
	h uintptr
	*QAccessibleEvent
}

func bindQAccessibleTableModelChangeEvent(h uintptr) *QAccessibleTableModelChangeEvent {
	if h == 0 {
		return nil
	}
	return &QAccessibleTableModelChangeEvent{h: h, QAccessibleEvent: bindQAccessibleEvent(pg_QAccessibleTableModelChangeEvent_base(h))}
}

func newownedQAccessibleTableModelChangeEvent(h uintptr) *QAccessibleTableModelChangeEvent {
	o := bindQAccessibleTableModelChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTableModelChangeEvent).Delete)
	}
	return o
}

func (this *QAccessibleTableModelChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTableModelChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTableModelChangeEvent(h unsafe.Pointer) *QAccessibleTableModelChangeEvent { return bindQAccessibleTableModelChangeEvent(uintptr(h)) }
func UnsafeNewQAccessibleTableModelChangeEventOwned(h unsafe.Pointer) *QAccessibleTableModelChangeEvent { return newownedQAccessibleTableModelChangeEvent(uintptr(h)) }

func (this *QAccessibleTableModelChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleEvent != nil {
		this.QAccessibleEvent.UnsafeRelease()
	}
}

func NewQAccessibleTableModelChangeEvent(obj *QObject, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_new, uintptr(obj.UnsafePointer()), uintptr(changeType))
	__o := newownedQAccessibleTableModelChangeEvent(__r)
	runtime.KeepAlive(obj)
	return __o
}

func NewQAccessibleTableModelChangeEvent2(iface *QAccessibleInterface, changeType QAccessibleTableModelChangeEvent__ModelChangeType) *QAccessibleTableModelChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_new2, uintptr(iface.UnsafePointer()), uintptr(changeType))
	__o := newownedQAccessibleTableModelChangeEvent(__r)
	runtime.KeepAlive(iface)
	return __o
}

func (this *QAccessibleTableModelChangeEvent) SetModelChangeType(changeType QAccessibleTableModelChangeEvent__ModelChangeType) {
	purego.SyscallN(addr_QAccessibleTableModelChangeEvent_setModelChangeType, this.h, uintptr(changeType))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTableModelChangeEvent) ModelChangeType() QAccessibleTableModelChangeEvent__ModelChangeType {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_modelChangeType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAccessibleTableModelChangeEvent__ModelChangeType(__r)
}

func (this *QAccessibleTableModelChangeEvent) SetFirstRow(row int) {
	purego.SyscallN(addr_QAccessibleTableModelChangeEvent_setFirstRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTableModelChangeEvent) SetFirstColumn(col int) {
	purego.SyscallN(addr_QAccessibleTableModelChangeEvent_setFirstColumn, this.h, uintptr(col))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTableModelChangeEvent) SetLastRow(row int) {
	purego.SyscallN(addr_QAccessibleTableModelChangeEvent_setLastRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTableModelChangeEvent) SetLastColumn(col int) {
	purego.SyscallN(addr_QAccessibleTableModelChangeEvent_setLastColumn, this.h, uintptr(col))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTableModelChangeEvent) FirstRow() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_firstRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableModelChangeEvent) FirstColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_firstColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableModelChangeEvent) LastRow() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_lastRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableModelChangeEvent) LastColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableModelChangeEvent_lastColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableModelChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTableModelChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
