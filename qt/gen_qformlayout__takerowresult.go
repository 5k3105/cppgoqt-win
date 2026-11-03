package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFormLayout__TakeRowResult_new uintptr
	addr_QFormLayout__TakeRowResult_new2 uintptr
	addr_QFormLayout__TakeRowResult_labelItem_get uintptr
	addr_QFormLayout__TakeRowResult_labelItem_set uintptr
	addr_QFormLayout__TakeRowResult_fieldItem_get uintptr
	addr_QFormLayout__TakeRowResult_fieldItem_set uintptr
	pg_QFormLayout__TakeRowResult_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFormLayout__TakeRowResult_new, "QFormLayout__TakeRowResult_new")
	symaddr(&addr_QFormLayout__TakeRowResult_new2, "QFormLayout__TakeRowResult_new2")
	symaddr(&addr_QFormLayout__TakeRowResult_labelItem_get, "QFormLayout__TakeRowResult_labelItem_get")
	symaddr(&addr_QFormLayout__TakeRowResult_labelItem_set, "QFormLayout__TakeRowResult_labelItem_set")
	symaddr(&addr_QFormLayout__TakeRowResult_fieldItem_get, "QFormLayout__TakeRowResult_fieldItem_get")
	symaddr(&addr_QFormLayout__TakeRowResult_fieldItem_set, "QFormLayout__TakeRowResult_fieldItem_set")
	reg(&pg_QFormLayout__TakeRowResult_delete, "QFormLayout__TakeRowResult_delete")
}

type QFormLayout__TakeRowResult struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFormLayout__TakeRowResult) UnsafeGoID() int64     { return this.vid }
func (this *QFormLayout__TakeRowResult) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFormLayout__TakeRowResult(h uintptr) *QFormLayout__TakeRowResult {
	if h == 0 {
		return nil
	}
	return &QFormLayout__TakeRowResult{h: h}
}

func newownedQFormLayout__TakeRowResult(h uintptr) *QFormLayout__TakeRowResult {
	o := bindQFormLayout__TakeRowResult(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFormLayout__TakeRowResult).Delete)
	}
	return o
}

func (this *QFormLayout__TakeRowResult) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFormLayout__TakeRowResult) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFormLayout__TakeRowResult(h unsafe.Pointer) *QFormLayout__TakeRowResult { return bindQFormLayout__TakeRowResult(uintptr(h)) }
func UnsafeNewQFormLayout__TakeRowResultOwned(h unsafe.Pointer) *QFormLayout__TakeRowResult { return newownedQFormLayout__TakeRowResult(uintptr(h)) }

func (this *QFormLayout__TakeRowResult) UnsafeRelease() { this.h = 0 }

func NewQFormLayout__TakeRowResult() *QFormLayout__TakeRowResult {
	__r, _, _ := purego.SyscallN(addr_QFormLayout__TakeRowResult_new)
	__o := newownedQFormLayout__TakeRowResult(__r)
	return __o
}

func NewQFormLayout__TakeRowResult2(arg1 *QFormLayout__TakeRowResult) *QFormLayout__TakeRowResult {
	__r, _, _ := purego.SyscallN(addr_QFormLayout__TakeRowResult_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQFormLayout__TakeRowResult(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFormLayout__TakeRowResult) LabelItem() *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout__TakeRowResult_labelItem_get, this.h)
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout__TakeRowResult) SetLabelItem(labelItem *QLayoutItem) {
	purego.SyscallN(addr_QFormLayout__TakeRowResult_labelItem_set, this.h, uintptr(labelItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(labelItem)
}

func (this *QFormLayout__TakeRowResult) FieldItem() *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout__TakeRowResult_fieldItem_get, this.h)
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout__TakeRowResult) SetFieldItem(fieldItem *QLayoutItem) {
	purego.SyscallN(addr_QFormLayout__TakeRowResult_fieldItem_set, this.h, uintptr(fieldItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fieldItem)
}

func (this *QFormLayout__TakeRowResult) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFormLayout__TakeRowResult_delete(this.h)
	}
	this.UnsafeRelease()
}
