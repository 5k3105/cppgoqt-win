package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QItemEditorCreatorBase_createWidget uintptr
	addr_QItemEditorCreatorBase_valuePropertyName uintptr
	addr_QItemEditorCreatorBase_OperatorAssign uintptr
	pg_QItemEditorCreatorBase__set_override func(uintptr, int32)
	pg_QItemEditorCreatorBase__Go_new func(int64) uintptr
	pg_QItemEditorCreatorBase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QItemEditorCreatorBase_createWidget, "QItemEditorCreatorBase_createWidget")
	symaddr(&addr_QItemEditorCreatorBase_valuePropertyName, "QItemEditorCreatorBase_valuePropertyName")
	symaddr(&addr_QItemEditorCreatorBase_OperatorAssign, "QItemEditorCreatorBase_OperatorAssign")
	reg(&pg_QItemEditorCreatorBase__set_override, "QItemEditorCreatorBase__set_override")
	reg(&pg_QItemEditorCreatorBase__Go_new, "QItemEditorCreatorBase__Go_new")
	reg(&pg_QItemEditorCreatorBase_delete, "QItemEditorCreatorBase_delete")
}

type QItemEditorCreatorBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QItemEditorCreatorBase) UnsafeGoID() int64     { return this.vid }
func (this *QItemEditorCreatorBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQItemEditorCreatorBase(h uintptr) *QItemEditorCreatorBase {
	if h == 0 {
		return nil
	}
	return &QItemEditorCreatorBase{h: h}
}

func newownedQItemEditorCreatorBase(h uintptr) *QItemEditorCreatorBase {
	o := bindQItemEditorCreatorBase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QItemEditorCreatorBase).Delete)
	}
	return o
}

func (this *QItemEditorCreatorBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemEditorCreatorBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemEditorCreatorBase(h unsafe.Pointer) *QItemEditorCreatorBase { return bindQItemEditorCreatorBase(uintptr(h)) }
func UnsafeNewQItemEditorCreatorBaseOwned(h unsafe.Pointer) *QItemEditorCreatorBase { return newownedQItemEditorCreatorBase(uintptr(h)) }

func (this *QItemEditorCreatorBase) UnsafeRelease() { this.h = 0 }

func (this *QItemEditorCreatorBase) CreateWidget(parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QItemEditorCreatorBase_createWidget, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QItemEditorCreatorBase) ValuePropertyName() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QItemEditorCreatorBase_valuePropertyName, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QItemEditorCreatorBase) OperatorAssign(arg1 *QItemEditorCreatorBase) *QItemEditorCreatorBase {
	__r, _, _ := purego.SyscallN(addr_QItemEditorCreatorBase_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQItemEditorCreatorBase(__r)
}

func QItemEditorCreatorBase_virtIdx(slot int64) int32 {
	switch slot {
	case -8470300370605365395:
		return 0
	case -3548500303342436100:
		return 1
	}
	return -1
}

func QItemEditorCreatorBase_enable(h uintptr, idx int32) { pg_QItemEditorCreatorBase__set_override(h, idx) }

func NewQItemEditorCreatorBase() *QItemEditorCreatorBase {
	id := nextVirtID()
	o := newownedQItemEditorCreatorBase(pg_QItemEditorCreatorBase__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QItemEditorCreatorBase_virtIdx, QItemEditorCreatorBase_enable)
	return o
}

func (this *QItemEditorCreatorBase) OnCreateWidget(fn func(*QWidget) *QWidget) {
	registerVirt(this.UnsafeGoID(), -8470300370605365395, func(parent uintptr) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "createWidget"
}

func (this *QItemEditorCreatorBase) OnValuePropertyName(fn func() *QByteArray) {
	registerVirt(this.UnsafeGoID(), -3548500303342436100, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "valuePropertyName"
}

func (this *QItemEditorCreatorBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemEditorCreatorBase_delete(this.h)
	}
	this.UnsafeRelease()
}
