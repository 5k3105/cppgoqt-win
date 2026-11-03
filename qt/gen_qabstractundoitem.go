package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractUndoItem_undo uintptr
	addr_QAbstractUndoItem_redo uintptr
	addr_QAbstractUndoItem_OperatorAssign uintptr
	pg_QAbstractUndoItem__set_override func(uintptr, int32)
	pg_QAbstractUndoItem__Go_new func(int64) uintptr
	pg_QAbstractUndoItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractUndoItem_undo, "QAbstractUndoItem_undo")
	symaddr(&addr_QAbstractUndoItem_redo, "QAbstractUndoItem_redo")
	symaddr(&addr_QAbstractUndoItem_OperatorAssign, "QAbstractUndoItem_OperatorAssign")
	reg(&pg_QAbstractUndoItem__set_override, "QAbstractUndoItem__set_override")
	reg(&pg_QAbstractUndoItem__Go_new, "QAbstractUndoItem__Go_new")
	reg(&pg_QAbstractUndoItem_delete, "QAbstractUndoItem_delete")
}

type QAbstractUndoItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractUndoItem) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractUndoItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractUndoItem(h uintptr) *QAbstractUndoItem {
	if h == 0 {
		return nil
	}
	return &QAbstractUndoItem{h: h}
}

func newownedQAbstractUndoItem(h uintptr) *QAbstractUndoItem {
	o := bindQAbstractUndoItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractUndoItem).Delete)
	}
	return o
}

func (this *QAbstractUndoItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractUndoItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractUndoItem(h unsafe.Pointer) *QAbstractUndoItem { return bindQAbstractUndoItem(uintptr(h)) }
func UnsafeNewQAbstractUndoItemOwned(h unsafe.Pointer) *QAbstractUndoItem { return newownedQAbstractUndoItem(uintptr(h)) }

func (this *QAbstractUndoItem) UnsafeRelease() { this.h = 0 }

func (this *QAbstractUndoItem) Undo() {
	purego.SyscallN(addr_QAbstractUndoItem_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractUndoItem) Redo() {
	purego.SyscallN(addr_QAbstractUndoItem_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractUndoItem) OperatorAssign(arg1 *QAbstractUndoItem) *QAbstractUndoItem {
	__r, _, _ := purego.SyscallN(addr_QAbstractUndoItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAbstractUndoItem(__r)
}

func QAbstractUndoItem_virtIdx(slot int64) int32 {
	switch slot {
	case 7259214242580717359:
		return 0
	case 5544063219051948393:
		return 1
	}
	return -1
}

func QAbstractUndoItem_enable(h uintptr, idx int32) { pg_QAbstractUndoItem__set_override(h, idx) }

func NewQAbstractUndoItem() *QAbstractUndoItem {
	id := nextVirtID()
	o := newownedQAbstractUndoItem(pg_QAbstractUndoItem__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractUndoItem_virtIdx, QAbstractUndoItem_enable)
	return o
}

func (this *QAbstractUndoItem) OnUndo(fn func()) {
	registerVirt(this.UnsafeGoID(), 7259214242580717359, func() { fn() }) // slot "undo"
}

func (this *QAbstractUndoItem) OnRedo(fn func()) {
	registerVirt(this.UnsafeGoID(), 5544063219051948393, func() { fn() }) // slot "redo"
}

func (this *QAbstractUndoItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractUndoItem_delete(this.h)
	}
	this.UnsafeRelease()
}
