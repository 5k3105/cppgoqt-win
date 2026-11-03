package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleHintReturnVariant_base func(uintptr) uintptr
	addr_QStyleHintReturnVariant_new uintptr
	addr_QStyleHintReturnVariant_new2 uintptr
	addr_QStyleHintReturnVariant_OperatorAssign uintptr
	addr_QStyleHintReturnVariant_variant_get uintptr
	addr_QStyleHintReturnVariant_variant_set uintptr
	pg_QStyleHintReturnVariant_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleHintReturnVariant_base, "QStyleHintReturnVariant_base")
	symaddr(&addr_QStyleHintReturnVariant_new, "QStyleHintReturnVariant_new")
	symaddr(&addr_QStyleHintReturnVariant_new2, "QStyleHintReturnVariant_new2")
	symaddr(&addr_QStyleHintReturnVariant_OperatorAssign, "QStyleHintReturnVariant_OperatorAssign")
	symaddr(&addr_QStyleHintReturnVariant_variant_get, "QStyleHintReturnVariant_variant_get")
	symaddr(&addr_QStyleHintReturnVariant_variant_set, "QStyleHintReturnVariant_variant_set")
	reg(&pg_QStyleHintReturnVariant_delete, "QStyleHintReturnVariant_delete")
}

type QStyleHintReturnVariant struct {
	h uintptr
	*QStyleHintReturn
}

func bindQStyleHintReturnVariant(h uintptr) *QStyleHintReturnVariant {
	if h == 0 {
		return nil
	}
	return &QStyleHintReturnVariant{h: h, QStyleHintReturn: bindQStyleHintReturn(pg_QStyleHintReturnVariant_base(h))}
}

func newownedQStyleHintReturnVariant(h uintptr) *QStyleHintReturnVariant {
	o := bindQStyleHintReturnVariant(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleHintReturnVariant).Delete)
	}
	return o
}

func (this *QStyleHintReturnVariant) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleHintReturnVariant) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleHintReturnVariant(h unsafe.Pointer) *QStyleHintReturnVariant { return bindQStyleHintReturnVariant(uintptr(h)) }
func UnsafeNewQStyleHintReturnVariantOwned(h unsafe.Pointer) *QStyleHintReturnVariant { return newownedQStyleHintReturnVariant(uintptr(h)) }

func (this *QStyleHintReturnVariant) UnsafeRelease() {
	this.h = 0
	if this.QStyleHintReturn != nil {
		this.QStyleHintReturn.UnsafeRelease()
	}
}

func NewQStyleHintReturnVariant() *QStyleHintReturnVariant {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnVariant_new)
	__o := newownedQStyleHintReturnVariant(__r)
	return __o
}

func NewQStyleHintReturnVariant2(arg1 *QStyleHintReturnVariant) *QStyleHintReturnVariant {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnVariant_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQStyleHintReturnVariant(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QStyleHintReturnVariant) OperatorAssign(arg1 *QStyleHintReturnVariant) *QStyleHintReturnVariant {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnVariant_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleHintReturnVariant(__r)
}

func (this *QStyleHintReturnVariant) Variant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnVariant_variant_get, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStyleHintReturnVariant) SetVariant(variant *QVariant) {
	purego.SyscallN(addr_QStyleHintReturnVariant_variant_set, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QStyleHintReturnVariant) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleHintReturnVariant_delete(this.h)
	}
	this.UnsafeRelease()
}
