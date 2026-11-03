package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainterPath__Element_new uintptr
	addr_QPainterPath__Element_new2 uintptr
	addr_QPainterPath__Element_isMoveTo uintptr
	addr_QPainterPath__Element_isLineTo uintptr
	addr_QPainterPath__Element_isCurveTo uintptr
	addr_QPainterPath__Element_ToQPointF uintptr
	addr_QPainterPath__Element_OperatorEqual uintptr
	addr_QPainterPath__Element_OperatorNotEqual uintptr
	pg_QPainterPath__Element_x_get func(uintptr) float64
	pg_QPainterPath__Element_x_set func(uintptr, float64)
	pg_QPainterPath__Element_y_get func(uintptr) float64
	pg_QPainterPath__Element_y_set func(uintptr, float64)
	addr_QPainterPath__Element_type_get uintptr
	addr_QPainterPath__Element_type_set uintptr
	pg_QPainterPath__Element_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainterPath__Element_new, "QPainterPath__Element_new")
	symaddr(&addr_QPainterPath__Element_new2, "QPainterPath__Element_new2")
	symaddr(&addr_QPainterPath__Element_isMoveTo, "QPainterPath__Element_isMoveTo")
	symaddr(&addr_QPainterPath__Element_isLineTo, "QPainterPath__Element_isLineTo")
	symaddr(&addr_QPainterPath__Element_isCurveTo, "QPainterPath__Element_isCurveTo")
	symaddr(&addr_QPainterPath__Element_ToQPointF, "QPainterPath__Element_ToQPointF")
	symaddr(&addr_QPainterPath__Element_OperatorEqual, "QPainterPath__Element_OperatorEqual")
	symaddr(&addr_QPainterPath__Element_OperatorNotEqual, "QPainterPath__Element_OperatorNotEqual")
	reg(&pg_QPainterPath__Element_x_get, "QPainterPath__Element_x_get")
	reg(&pg_QPainterPath__Element_x_set, "QPainterPath__Element_x_set")
	reg(&pg_QPainterPath__Element_y_get, "QPainterPath__Element_y_get")
	reg(&pg_QPainterPath__Element_y_set, "QPainterPath__Element_y_set")
	symaddr(&addr_QPainterPath__Element_type_get, "QPainterPath__Element_type_get")
	symaddr(&addr_QPainterPath__Element_type_set, "QPainterPath__Element_type_set")
	reg(&pg_QPainterPath__Element_delete, "QPainterPath__Element_delete")
}

type QPainterPath__Element struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainterPath__Element) UnsafeGoID() int64     { return this.vid }
func (this *QPainterPath__Element) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainterPath__Element(h uintptr) *QPainterPath__Element {
	if h == 0 {
		return nil
	}
	return &QPainterPath__Element{h: h}
}

func newownedQPainterPath__Element(h uintptr) *QPainterPath__Element {
	o := bindQPainterPath__Element(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainterPath__Element).Delete)
	}
	return o
}

func (this *QPainterPath__Element) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainterPath__Element) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainterPath__Element(h unsafe.Pointer) *QPainterPath__Element { return bindQPainterPath__Element(uintptr(h)) }
func UnsafeNewQPainterPath__ElementOwned(h unsafe.Pointer) *QPainterPath__Element { return newownedQPainterPath__Element(uintptr(h)) }

func (this *QPainterPath__Element) UnsafeRelease() { this.h = 0 }

func NewQPainterPath__Element() *QPainterPath__Element {
	__r, _, _ := purego.SyscallN(addr_QPainterPath__Element_new)
	__o := newownedQPainterPath__Element(__r)
	return __o
}

func NewQPainterPath__Element2(arg1 *QPainterPath__Element) *QPainterPath__Element {
	__r, _, _ := purego.SyscallN(addr_QPainterPath__Element_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPainterPath__Element(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPainterPath__Element) IsMoveTo() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_isMoveTo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainterPath__Element) IsLineTo() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_isLineTo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainterPath__Element) IsCurveTo() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_isCurveTo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainterPath__Element) ToQPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPainterPath__Element_ToQPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPainterPath__Element) OperatorEqual(e *QPainterPath__Element) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_OperatorEqual, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QPainterPath__Element) OperatorNotEqual(e *QPainterPath__Element) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_OperatorNotEqual, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QPainterPath__Element) X() float64 {
	__r := pg_QPainterPath__Element_x_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath__Element) SetX(x float64) {
	pg_QPainterPath__Element_x_set(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QPainterPath__Element) Y() float64 {
	__r := pg_QPainterPath__Element_y_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPath__Element) SetY(y float64) {
	pg_QPainterPath__Element_y_set(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QPainterPath__Element) Type() QPainterPath__ElementType {
	__r0, _, _ := purego.SyscallN(addr_QPainterPath__Element_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainterPath__ElementType(__r)
}

func (this *QPainterPath__Element) SetType(type_ QPainterPath__ElementType) {
	purego.SyscallN(addr_QPainterPath__Element_type_set, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QPainterPath__Element) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainterPath__Element_delete(this.h)
	}
	this.UnsafeRelease()
}
