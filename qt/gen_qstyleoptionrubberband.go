package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionRubberBand_base func(uintptr) uintptr
	addr_QStyleOptionRubberBand_new uintptr
	addr_QStyleOptionRubberBand_new2 uintptr
	addr_QStyleOptionRubberBand_OperatorAssign uintptr
	addr_QStyleOptionRubberBand_shape_get uintptr
	addr_QStyleOptionRubberBand_shape_set uintptr
	addr_QStyleOptionRubberBand_opaque_get uintptr
	addr_QStyleOptionRubberBand_opaque_set uintptr
	pg_QStyleOptionRubberBand_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionRubberBand_base, "QStyleOptionRubberBand_base")
	symaddr(&addr_QStyleOptionRubberBand_new, "QStyleOptionRubberBand_new")
	symaddr(&addr_QStyleOptionRubberBand_new2, "QStyleOptionRubberBand_new2")
	symaddr(&addr_QStyleOptionRubberBand_OperatorAssign, "QStyleOptionRubberBand_OperatorAssign")
	symaddr(&addr_QStyleOptionRubberBand_shape_get, "QStyleOptionRubberBand_shape_get")
	symaddr(&addr_QStyleOptionRubberBand_shape_set, "QStyleOptionRubberBand_shape_set")
	symaddr(&addr_QStyleOptionRubberBand_opaque_get, "QStyleOptionRubberBand_opaque_get")
	symaddr(&addr_QStyleOptionRubberBand_opaque_set, "QStyleOptionRubberBand_opaque_set")
	reg(&pg_QStyleOptionRubberBand_delete, "QStyleOptionRubberBand_delete")
}

type QStyleOptionRubberBand struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionRubberBand(h uintptr) *QStyleOptionRubberBand {
	if h == 0 {
		return nil
	}
	return &QStyleOptionRubberBand{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionRubberBand_base(h))}
}

func newownedQStyleOptionRubberBand(h uintptr) *QStyleOptionRubberBand {
	o := bindQStyleOptionRubberBand(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionRubberBand).Delete)
	}
	return o
}

func (this *QStyleOptionRubberBand) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionRubberBand) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionRubberBand(h unsafe.Pointer) *QStyleOptionRubberBand { return bindQStyleOptionRubberBand(uintptr(h)) }
func UnsafeNewQStyleOptionRubberBandOwned(h unsafe.Pointer) *QStyleOptionRubberBand { return newownedQStyleOptionRubberBand(uintptr(h)) }

func (this *QStyleOptionRubberBand) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionRubberBand() *QStyleOptionRubberBand {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionRubberBand_new)
	__o := newownedQStyleOptionRubberBand(__r)
	return __o
}

func NewQStyleOptionRubberBand2(other *QStyleOptionRubberBand) *QStyleOptionRubberBand {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionRubberBand_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionRubberBand(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionRubberBand) OperatorAssign(arg1 *QStyleOptionRubberBand) *QStyleOptionRubberBand {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionRubberBand_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionRubberBand(__r)
}

func (this *QStyleOptionRubberBand) Shape() QRubberBand__Shape {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionRubberBand_shape_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRubberBand__Shape(__r)
}

func (this *QStyleOptionRubberBand) SetShape(shape QRubberBand__Shape) {
	purego.SyscallN(addr_QStyleOptionRubberBand_shape_set, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionRubberBand) Opaque() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionRubberBand_opaque_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionRubberBand) SetOpaque(opaque bool) {
	purego.SyscallN(addr_QStyleOptionRubberBand_opaque_set, this.h, uintptr(b2i(opaque)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionRubberBand) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionRubberBand_delete(this.h)
	}
	this.UnsafeRelease()
}
