package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSizeF_new uintptr
	addr_QSizeF_new2 uintptr
	pg_QSizeF_new3 func(float64, float64) uintptr
	addr_QSizeF_new4 uintptr
	addr_QSizeF_isNull uintptr
	addr_QSizeF_isEmpty uintptr
	addr_QSizeF_isValid uintptr
	pg_QSizeF_width func(uintptr) float64
	pg_QSizeF_height func(uintptr) float64
	pg_QSizeF_setWidth func(uintptr, float64)
	pg_QSizeF_setHeight func(uintptr, float64)
	addr_QSizeF_transpose uintptr
	addr_QSizeF_transposed uintptr
	pg_QSizeF_scale func(uintptr, float64, float64, int32)
	addr_QSizeF_scale2 uintptr
	pg_QSizeF_scaled func(uintptr, float64, float64, int32) uintptr
	addr_QSizeF_scaled2 uintptr
	addr_QSizeF_expandedTo uintptr
	addr_QSizeF_boundedTo uintptr
	addr_QSizeF_grownBy uintptr
	addr_QSizeF_shrunkBy uintptr
	addr_QSizeF_OperatorPlusAssign uintptr
	addr_QSizeF_OperatorMinusAssign uintptr
	pg_QSizeF_OperatorMultiplyAssign func(uintptr, float64) uintptr
	pg_QSizeF_OperatorDivideAssign func(uintptr, float64) uintptr
	addr_QSizeF_toSize uintptr
	addr_QSizeF_OperatorAssign uintptr
	addr_QSizeF_OperatorEqual uintptr
	addr_QSizeF_OperatorNotEqual uintptr
	addr_QSizeF_OperatorEqual2 uintptr
	addr_QSizeF_OperatorNotEqual2 uintptr
	addr_QSizeF_OperatorPlus uintptr
	addr_QSizeF_OperatorMinus uintptr
	pg_QSizeF_OperatorMultiply func(uintptr, float64) uintptr
	pg_QSizeF_OperatorDivide func(uintptr, float64) uintptr
	pg_QSizeF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSizeF_new, "QSizeF_new")
	symaddr(&addr_QSizeF_new2, "QSizeF_new2")
	reg(&pg_QSizeF_new3, "QSizeF_new3")
	symaddr(&addr_QSizeF_new4, "QSizeF_new4")
	symaddr(&addr_QSizeF_isNull, "QSizeF_isNull")
	symaddr(&addr_QSizeF_isEmpty, "QSizeF_isEmpty")
	symaddr(&addr_QSizeF_isValid, "QSizeF_isValid")
	reg(&pg_QSizeF_width, "QSizeF_width")
	reg(&pg_QSizeF_height, "QSizeF_height")
	reg(&pg_QSizeF_setWidth, "QSizeF_setWidth")
	reg(&pg_QSizeF_setHeight, "QSizeF_setHeight")
	symaddr(&addr_QSizeF_transpose, "QSizeF_transpose")
	symaddr(&addr_QSizeF_transposed, "QSizeF_transposed")
	reg(&pg_QSizeF_scale, "QSizeF_scale")
	symaddr(&addr_QSizeF_scale2, "QSizeF_scale2")
	reg(&pg_QSizeF_scaled, "QSizeF_scaled")
	symaddr(&addr_QSizeF_scaled2, "QSizeF_scaled2")
	symaddr(&addr_QSizeF_expandedTo, "QSizeF_expandedTo")
	symaddr(&addr_QSizeF_boundedTo, "QSizeF_boundedTo")
	symaddr(&addr_QSizeF_grownBy, "QSizeF_grownBy")
	symaddr(&addr_QSizeF_shrunkBy, "QSizeF_shrunkBy")
	symaddr(&addr_QSizeF_OperatorPlusAssign, "QSizeF_OperatorPlusAssign")
	symaddr(&addr_QSizeF_OperatorMinusAssign, "QSizeF_OperatorMinusAssign")
	reg(&pg_QSizeF_OperatorMultiplyAssign, "QSizeF_OperatorMultiplyAssign")
	reg(&pg_QSizeF_OperatorDivideAssign, "QSizeF_OperatorDivideAssign")
	symaddr(&addr_QSizeF_toSize, "QSizeF_toSize")
	symaddr(&addr_QSizeF_OperatorAssign, "QSizeF_OperatorAssign")
	symaddr(&addr_QSizeF_OperatorEqual, "QSizeF_OperatorEqual")
	symaddr(&addr_QSizeF_OperatorNotEqual, "QSizeF_OperatorNotEqual")
	symaddr(&addr_QSizeF_OperatorEqual2, "QSizeF_OperatorEqual2")
	symaddr(&addr_QSizeF_OperatorNotEqual2, "QSizeF_OperatorNotEqual2")
	symaddr(&addr_QSizeF_OperatorPlus, "QSizeF_OperatorPlus")
	symaddr(&addr_QSizeF_OperatorMinus, "QSizeF_OperatorMinus")
	reg(&pg_QSizeF_OperatorMultiply, "QSizeF_OperatorMultiply")
	reg(&pg_QSizeF_OperatorDivide, "QSizeF_OperatorDivide")
	reg(&pg_QSizeF_delete, "QSizeF_delete")
}

type QSizeF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSizeF) UnsafeGoID() int64     { return this.vid }
func (this *QSizeF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSizeF(h uintptr) *QSizeF {
	if h == 0 {
		return nil
	}
	return &QSizeF{h: h}
}

func newownedQSizeF(h uintptr) *QSizeF {
	o := bindQSizeF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSizeF).Delete)
	}
	return o
}

func (this *QSizeF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSizeF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSizeF(h unsafe.Pointer) *QSizeF { return bindQSizeF(uintptr(h)) }
func UnsafeNewQSizeFOwned(h unsafe.Pointer) *QSizeF { return newownedQSizeF(uintptr(h)) }

func (this *QSizeF) UnsafeRelease() { this.h = 0 }

func NewQSizeF() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_new)
	__o := newownedQSizeF(__r)
	return __o
}

func NewQSizeF2(sz *QSize) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_new2, uintptr(sz.UnsafePointer()))
	__o := newownedQSizeF(__r)
	runtime.KeepAlive(sz)
	return __o
}

func NewQSizeF3(w float64, h float64) *QSizeF {
	__o := newownedQSizeF(pg_QSizeF_new3(w, h))
	return __o
}

func NewQSizeF4(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQSizeF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSizeF) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizeF) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizeF) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizeF) Width() float64 {
	__r := pg_QSizeF_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSizeF) Height() float64 {
	__r := pg_QSizeF_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSizeF) SetWidth(w float64) {
	pg_QSizeF_setWidth(this.h, w)
	runtime.KeepAlive(this)
}

func (this *QSizeF) SetHeight(h float64) {
	pg_QSizeF_setHeight(this.h, h)
	runtime.KeepAlive(this)
}

func (this *QSizeF) Transpose() {
	purego.SyscallN(addr_QSizeF_transpose, this.h)
	runtime.KeepAlive(this)
}

func (this *QSizeF) Transposed() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QSizeF) Scale(w float64, h float64, mode AspectRatioMode) {
	pg_QSizeF_scale(this.h, w, h, int32(mode))
	runtime.KeepAlive(this)
}

func (this *QSizeF) Scale2(s *QSizeF, mode AspectRatioMode) {
	purego.SyscallN(addr_QSizeF_scale2, this.h, uintptr(s.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QSizeF) Scaled(w float64, h float64, mode AspectRatioMode) *QSizeF {
	__r := pg_QSizeF_scaled(this.h, w, h, int32(mode))
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QSizeF) Scaled2(s *QSizeF, mode AspectRatioMode) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_scaled2, this.h, uintptr(s.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQSizeF(__r)
}

func (this *QSizeF) ExpandedTo(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_expandedTo, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQSizeF(__r)
}

func (this *QSizeF) BoundedTo(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_boundedTo, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQSizeF(__r)
}

func (this *QSizeF) GrownBy(m *QMarginsF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_grownBy, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQSizeF(__r)
}

func (this *QSizeF) ShrunkBy(m *QMarginsF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_shrunkBy, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQSizeF(__r)
}

func (this *QSizeF) OperatorPlusAssign(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_OperatorPlusAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSizeF(__r)
}

func (this *QSizeF) OperatorMinusAssign(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_OperatorMinusAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSizeF(__r)
}

func (this *QSizeF) OperatorMultiplyAssign(c float64) *QSizeF {
	__r := pg_QSizeF_OperatorMultiplyAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQSizeF(__r)
}

func (this *QSizeF) OperatorDivideAssign(c float64) *QSizeF {
	__r := pg_QSizeF_OperatorDivideAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQSizeF(__r)
}

func (this *QSizeF) ToSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSizeF_toSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSizeF) OperatorAssign(arg1 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSizeF(__r)
}

func (this *QSizeF) OperatorEqual(rhs *QSizeF) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSizeF) OperatorNotEqual(rhs *QSizeF) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSizeF) OperatorEqualWithRhs(rhs *QSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSizeF) OperatorNotEqualWithRhs(rhs *QSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizeF_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSizeF) OperatorPlus(s2 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_OperatorPlus, this.h, uintptr(s2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return newownedQSizeF(__r)
}

func (this *QSizeF) OperatorMinus(s2 *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSizeF_OperatorMinus, this.h, uintptr(s2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return newownedQSizeF(__r)
}

func (this *QSizeF) OperatorMultiply(c float64) *QSizeF {
	__r := pg_QSizeF_OperatorMultiply(this.h, c)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QSizeF) OperatorDivide(c float64) *QSizeF {
	__r := pg_QSizeF_OperatorDivide(this.h, c)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QSizeF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSizeF_delete(this.h)
	}
	this.UnsafeRelease()
}
