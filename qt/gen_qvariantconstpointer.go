package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVariantConstPointer_new uintptr
	addr_QVariantConstPointer_new2 uintptr
	addr_QVariantConstPointer_OperatorMultiply uintptr
	addr_QVariantConstPointer_OperatorMinusGreater uintptr
	addr_QVariantConstPointer_OperatorAssign uintptr
	pg_QVariantConstPointer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVariantConstPointer_new, "QVariantConstPointer_new")
	symaddr(&addr_QVariantConstPointer_new2, "QVariantConstPointer_new2")
	symaddr(&addr_QVariantConstPointer_OperatorMultiply, "QVariantConstPointer_OperatorMultiply")
	symaddr(&addr_QVariantConstPointer_OperatorMinusGreater, "QVariantConstPointer_OperatorMinusGreater")
	symaddr(&addr_QVariantConstPointer_OperatorAssign, "QVariantConstPointer_OperatorAssign")
	reg(&pg_QVariantConstPointer_delete, "QVariantConstPointer_delete")
}

type QVariantConstPointer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVariantConstPointer) UnsafeGoID() int64     { return this.vid }
func (this *QVariantConstPointer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVariantConstPointer(h uintptr) *QVariantConstPointer {
	if h == 0 {
		return nil
	}
	return &QVariantConstPointer{h: h}
}

func newownedQVariantConstPointer(h uintptr) *QVariantConstPointer {
	o := bindQVariantConstPointer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVariantConstPointer).Delete)
	}
	return o
}

func (this *QVariantConstPointer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVariantConstPointer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVariantConstPointer(h unsafe.Pointer) *QVariantConstPointer { return bindQVariantConstPointer(uintptr(h)) }
func UnsafeNewQVariantConstPointerOwned(h unsafe.Pointer) *QVariantConstPointer { return newownedQVariantConstPointer(uintptr(h)) }

func (this *QVariantConstPointer) UnsafeRelease() { this.h = 0 }

func NewQVariantConstPointer(variant *QVariant) *QVariantConstPointer {
	__r, _, _ := purego.SyscallN(addr_QVariantConstPointer_new, uintptr(variant.UnsafePointer()))
	__o := newownedQVariantConstPointer(__r)
	runtime.KeepAlive(variant)
	return __o
}

func NewQVariantConstPointer2(arg1 *QVariantConstPointer) *QVariantConstPointer {
	__r, _, _ := purego.SyscallN(addr_QVariantConstPointer_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQVariantConstPointer(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QVariantConstPointer) OperatorMultiply() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariantConstPointer_OperatorMultiply, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVariantConstPointer) OperatorMinusGreater() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariantConstPointer_OperatorMinusGreater, this.h)
	runtime.KeepAlive(this)
	return bindQVariant(__r)
}

func (this *QVariantConstPointer) OperatorAssign(arg1 *QVariantConstPointer) *QVariantConstPointer {
	__r, _, _ := purego.SyscallN(addr_QVariantConstPointer_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQVariantConstPointer(__r)
}

func (this *QVariantConstPointer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVariantConstPointer_delete(this.h)
	}
	this.UnsafeRelease()
}
