package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSize_new uintptr
	addr_QSize_new2 uintptr
	addr_QSize_new3 uintptr
	addr_QSize_isNull uintptr
	addr_QSize_isEmpty uintptr
	addr_QSize_isValid uintptr
	addr_QSize_width uintptr
	addr_QSize_height uintptr
	addr_QSize_setWidth uintptr
	addr_QSize_setHeight uintptr
	addr_QSize_transpose uintptr
	addr_QSize_transposed uintptr
	addr_QSize_scale uintptr
	addr_QSize_scale2 uintptr
	addr_QSize_scaled uintptr
	addr_QSize_scaled2 uintptr
	addr_QSize_expandedTo uintptr
	addr_QSize_boundedTo uintptr
	addr_QSize_grownBy uintptr
	addr_QSize_shrunkBy uintptr
	addr_QSize_OperatorPlusAssign uintptr
	addr_QSize_OperatorMinusAssign uintptr
	pg_QSize_OperatorMultiplyAssign func(uintptr, float64) uintptr
	pg_QSize_OperatorDivideAssign func(uintptr, float64) uintptr
	addr_QSize_toSizeF uintptr
	addr_QSize_OperatorAssign uintptr
	addr_QSize_OperatorEqual uintptr
	addr_QSize_OperatorNotEqual uintptr
	addr_QSize_OperatorPlus uintptr
	addr_QSize_OperatorMinus uintptr
	pg_QSize_OperatorMultiply func(uintptr, float64) uintptr
	pg_QSize_OperatorDivide func(uintptr, float64) uintptr
	addr_QSize_OperatorEqual2 uintptr
	addr_QSize_OperatorNotEqual2 uintptr
	pg_QSize_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSize_new, "QSize_new")
	symaddr(&addr_QSize_new2, "QSize_new2")
	symaddr(&addr_QSize_new3, "QSize_new3")
	symaddr(&addr_QSize_isNull, "QSize_isNull")
	symaddr(&addr_QSize_isEmpty, "QSize_isEmpty")
	symaddr(&addr_QSize_isValid, "QSize_isValid")
	symaddr(&addr_QSize_width, "QSize_width")
	symaddr(&addr_QSize_height, "QSize_height")
	symaddr(&addr_QSize_setWidth, "QSize_setWidth")
	symaddr(&addr_QSize_setHeight, "QSize_setHeight")
	symaddr(&addr_QSize_transpose, "QSize_transpose")
	symaddr(&addr_QSize_transposed, "QSize_transposed")
	symaddr(&addr_QSize_scale, "QSize_scale")
	symaddr(&addr_QSize_scale2, "QSize_scale2")
	symaddr(&addr_QSize_scaled, "QSize_scaled")
	symaddr(&addr_QSize_scaled2, "QSize_scaled2")
	symaddr(&addr_QSize_expandedTo, "QSize_expandedTo")
	symaddr(&addr_QSize_boundedTo, "QSize_boundedTo")
	symaddr(&addr_QSize_grownBy, "QSize_grownBy")
	symaddr(&addr_QSize_shrunkBy, "QSize_shrunkBy")
	symaddr(&addr_QSize_OperatorPlusAssign, "QSize_OperatorPlusAssign")
	symaddr(&addr_QSize_OperatorMinusAssign, "QSize_OperatorMinusAssign")
	reg(&pg_QSize_OperatorMultiplyAssign, "QSize_OperatorMultiplyAssign")
	reg(&pg_QSize_OperatorDivideAssign, "QSize_OperatorDivideAssign")
	symaddr(&addr_QSize_toSizeF, "QSize_toSizeF")
	symaddr(&addr_QSize_OperatorAssign, "QSize_OperatorAssign")
	symaddr(&addr_QSize_OperatorEqual, "QSize_OperatorEqual")
	symaddr(&addr_QSize_OperatorNotEqual, "QSize_OperatorNotEqual")
	symaddr(&addr_QSize_OperatorPlus, "QSize_OperatorPlus")
	symaddr(&addr_QSize_OperatorMinus, "QSize_OperatorMinus")
	reg(&pg_QSize_OperatorMultiply, "QSize_OperatorMultiply")
	reg(&pg_QSize_OperatorDivide, "QSize_OperatorDivide")
	symaddr(&addr_QSize_OperatorEqual2, "QSize_OperatorEqual2")
	symaddr(&addr_QSize_OperatorNotEqual2, "QSize_OperatorNotEqual2")
	reg(&pg_QSize_delete, "QSize_delete")
}

type QSize struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSize) UnsafeGoID() int64     { return this.vid }
func (this *QSize) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSize(h uintptr) *QSize {
	if h == 0 {
		return nil
	}
	return &QSize{h: h}
}

func newownedQSize(h uintptr) *QSize {
	o := bindQSize(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSize).Delete)
	}
	return o
}

func (this *QSize) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSize) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSize(h unsafe.Pointer) *QSize { return bindQSize(uintptr(h)) }
func UnsafeNewQSizeOwned(h unsafe.Pointer) *QSize { return newownedQSize(uintptr(h)) }

func (this *QSize) UnsafeRelease() { this.h = 0 }

func NewQSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_new)
	__o := newownedQSize(__r)
	return __o
}

func NewQSize2(w int, h int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_new2, uintptr(w), uintptr(h))
	__o := newownedQSize(__r)
	return __o
}

func NewQSize3(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQSize(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSize) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSize) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSize) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSize) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QSize_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSize) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QSize_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSize) SetWidth(w int) {
	purego.SyscallN(addr_QSize_setWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QSize) SetHeight(h int) {
	purego.SyscallN(addr_QSize_setHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QSize) Transpose() {
	purego.SyscallN(addr_QSize_transpose, this.h)
	runtime.KeepAlive(this)
}

func (this *QSize) Transposed() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSize) Scale(w int, h int, mode AspectRatioMode) {
	purego.SyscallN(addr_QSize_scale, this.h, uintptr(w), uintptr(h), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QSize) Scale2(s *QSize, mode AspectRatioMode) {
	purego.SyscallN(addr_QSize_scale2, this.h, uintptr(s.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QSize) Scaled(w int, h int, mode AspectRatioMode) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_scaled, this.h, uintptr(w), uintptr(h), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSize) Scaled2(s *QSize, mode AspectRatioMode) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_scaled2, this.h, uintptr(s.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QSize) ExpandedTo(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_expandedTo, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQSize(__r)
}

func (this *QSize) BoundedTo(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_boundedTo, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQSize(__r)
}

func (this *QSize) GrownBy(m *QMargins) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_grownBy, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQSize(__r)
}

func (this *QSize) ShrunkBy(m *QMargins) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_shrunkBy, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQSize(__r)
}

func (this *QSize) OperatorPlusAssign(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_OperatorPlusAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSize(__r)
}

func (this *QSize) OperatorMinusAssign(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_OperatorMinusAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSize(__r)
}

func (this *QSize) OperatorMultiplyAssign(c float64) *QSize {
	__r := pg_QSize_OperatorMultiplyAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQSize(__r)
}

func (this *QSize) OperatorDivideAssign(c float64) *QSize {
	__r := pg_QSize_OperatorDivideAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQSize(__r)
}

func (this *QSize) ToSizeF() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QSize_toSizeF, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QSize) OperatorAssign(arg1 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSize(__r)
}

func (this *QSize) OperatorEqual(rhs *QSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSize) OperatorNotEqual(rhs *QSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSize) OperatorPlus(s2 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_OperatorPlus, this.h, uintptr(s2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return newownedQSize(__r)
}

func (this *QSize) OperatorMinus(s2 *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSize_OperatorMinus, this.h, uintptr(s2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return newownedQSize(__r)
}

func (this *QSize) OperatorMultiply(c float64) *QSize {
	__r := pg_QSize_OperatorMultiply(this.h, c)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSize) OperatorDivide(c float64) *QSize {
	__r := pg_QSize_OperatorDivide(this.h, c)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSize) OperatorEqualWithRhs(rhs *QSizeF) bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSize) OperatorNotEqualWithRhs(rhs *QSizeF) bool {
	__r0, _, _ := purego.SyscallN(addr_QSize_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSize) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSize_delete(this.h)
	}
	this.UnsafeRelease()
}
