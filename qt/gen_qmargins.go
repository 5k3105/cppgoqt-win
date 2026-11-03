package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMargins_new uintptr
	addr_QMargins_new2 uintptr
	addr_QMargins_new3 uintptr
	addr_QMargins_isNull uintptr
	addr_QMargins_left uintptr
	addr_QMargins_top uintptr
	addr_QMargins_right uintptr
	addr_QMargins_bottom uintptr
	addr_QMargins_setLeft uintptr
	addr_QMargins_setTop uintptr
	addr_QMargins_setRight uintptr
	addr_QMargins_setBottom uintptr
	addr_QMargins_OperatorPlusAssign uintptr
	addr_QMargins_OperatorMinusAssign uintptr
	addr_QMargins_OperatorPlusAssign2 uintptr
	addr_QMargins_OperatorMinusAssign2 uintptr
	addr_QMargins_OperatorMultiplyAssign uintptr
	addr_QMargins_OperatorDivideAssign uintptr
	pg_QMargins_OperatorMultiplyAssign2 func(uintptr, float64) uintptr
	pg_QMargins_OperatorDivideAssign2 func(uintptr, float64) uintptr
	addr_QMargins_toMarginsF uintptr
	addr_QMargins_OperatorAssign uintptr
	addr_QMargins_OperatorPlus uintptr
	addr_QMargins_OperatorMinus uintptr
	addr_QMargins_OperatorPlus2 uintptr
	addr_QMargins_OperatorMinus2 uintptr
	addr_QMargins_OperatorMultiply uintptr
	pg_QMargins_OperatorMultiply2 func(uintptr, float64) uintptr
	addr_QMargins_OperatorDivide uintptr
	pg_QMargins_OperatorDivide2 func(uintptr, float64) uintptr
	addr_QMargins_OperatorBitwiseOr uintptr
	addr_QMargins_OperatorEqual uintptr
	addr_QMargins_OperatorNotEqual uintptr
	addr_QMargins_OperatorPlus3 uintptr
	addr_QMargins_OperatorMinus3 uintptr
	addr_QMargins_OperatorEqual2 uintptr
	addr_QMargins_OperatorNotEqual2 uintptr
	addr_QMargins_OperatorPlus4 uintptr
	pg_QMargins_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMargins_new, "QMargins_new")
	symaddr(&addr_QMargins_new2, "QMargins_new2")
	symaddr(&addr_QMargins_new3, "QMargins_new3")
	symaddr(&addr_QMargins_isNull, "QMargins_isNull")
	symaddr(&addr_QMargins_left, "QMargins_left")
	symaddr(&addr_QMargins_top, "QMargins_top")
	symaddr(&addr_QMargins_right, "QMargins_right")
	symaddr(&addr_QMargins_bottom, "QMargins_bottom")
	symaddr(&addr_QMargins_setLeft, "QMargins_setLeft")
	symaddr(&addr_QMargins_setTop, "QMargins_setTop")
	symaddr(&addr_QMargins_setRight, "QMargins_setRight")
	symaddr(&addr_QMargins_setBottom, "QMargins_setBottom")
	symaddr(&addr_QMargins_OperatorPlusAssign, "QMargins_OperatorPlusAssign")
	symaddr(&addr_QMargins_OperatorMinusAssign, "QMargins_OperatorMinusAssign")
	symaddr(&addr_QMargins_OperatorPlusAssign2, "QMargins_OperatorPlusAssign2")
	symaddr(&addr_QMargins_OperatorMinusAssign2, "QMargins_OperatorMinusAssign2")
	symaddr(&addr_QMargins_OperatorMultiplyAssign, "QMargins_OperatorMultiplyAssign")
	symaddr(&addr_QMargins_OperatorDivideAssign, "QMargins_OperatorDivideAssign")
	reg(&pg_QMargins_OperatorMultiplyAssign2, "QMargins_OperatorMultiplyAssign2")
	reg(&pg_QMargins_OperatorDivideAssign2, "QMargins_OperatorDivideAssign2")
	symaddr(&addr_QMargins_toMarginsF, "QMargins_toMarginsF")
	symaddr(&addr_QMargins_OperatorAssign, "QMargins_OperatorAssign")
	symaddr(&addr_QMargins_OperatorPlus, "QMargins_OperatorPlus")
	symaddr(&addr_QMargins_OperatorMinus, "QMargins_OperatorMinus")
	symaddr(&addr_QMargins_OperatorPlus2, "QMargins_OperatorPlus2")
	symaddr(&addr_QMargins_OperatorMinus2, "QMargins_OperatorMinus2")
	symaddr(&addr_QMargins_OperatorMultiply, "QMargins_OperatorMultiply")
	reg(&pg_QMargins_OperatorMultiply2, "QMargins_OperatorMultiply2")
	symaddr(&addr_QMargins_OperatorDivide, "QMargins_OperatorDivide")
	reg(&pg_QMargins_OperatorDivide2, "QMargins_OperatorDivide2")
	symaddr(&addr_QMargins_OperatorBitwiseOr, "QMargins_OperatorBitwiseOr")
	symaddr(&addr_QMargins_OperatorEqual, "QMargins_OperatorEqual")
	symaddr(&addr_QMargins_OperatorNotEqual, "QMargins_OperatorNotEqual")
	symaddr(&addr_QMargins_OperatorPlus3, "QMargins_OperatorPlus3")
	symaddr(&addr_QMargins_OperatorMinus3, "QMargins_OperatorMinus3")
	symaddr(&addr_QMargins_OperatorEqual2, "QMargins_OperatorEqual2")
	symaddr(&addr_QMargins_OperatorNotEqual2, "QMargins_OperatorNotEqual2")
	symaddr(&addr_QMargins_OperatorPlus4, "QMargins_OperatorPlus4")
	reg(&pg_QMargins_delete, "QMargins_delete")
}

type QMargins struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMargins) UnsafeGoID() int64     { return this.vid }
func (this *QMargins) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMargins(h uintptr) *QMargins {
	if h == 0 {
		return nil
	}
	return &QMargins{h: h}
}

func newownedQMargins(h uintptr) *QMargins {
	o := bindQMargins(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMargins).Delete)
	}
	return o
}

func (this *QMargins) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMargins) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMargins(h unsafe.Pointer) *QMargins { return bindQMargins(uintptr(h)) }
func UnsafeNewQMarginsOwned(h unsafe.Pointer) *QMargins { return newownedQMargins(uintptr(h)) }

func (this *QMargins) UnsafeRelease() { this.h = 0 }

func NewQMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_new)
	__o := newownedQMargins(__r)
	return __o
}

func NewQMargins2(left int, top int, right int, bottom int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_new2, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	__o := newownedQMargins(__r)
	return __o
}

func NewQMargins3(arg1 *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQMargins(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMargins) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QMargins_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMargins) Left() int {
	__r0, _, _ := purego.SyscallN(addr_QMargins_left, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMargins) Top() int {
	__r0, _, _ := purego.SyscallN(addr_QMargins_top, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMargins) Right() int {
	__r0, _, _ := purego.SyscallN(addr_QMargins_right, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMargins) Bottom() int {
	__r0, _, _ := purego.SyscallN(addr_QMargins_bottom, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMargins) SetLeft(left int) {
	purego.SyscallN(addr_QMargins_setLeft, this.h, uintptr(left))
	runtime.KeepAlive(this)
}

func (this *QMargins) SetTop(top int) {
	purego.SyscallN(addr_QMargins_setTop, this.h, uintptr(top))
	runtime.KeepAlive(this)
}

func (this *QMargins) SetRight(right int) {
	purego.SyscallN(addr_QMargins_setRight, this.h, uintptr(right))
	runtime.KeepAlive(this)
}

func (this *QMargins) SetBottom(bottom int) {
	purego.SyscallN(addr_QMargins_setBottom, this.h, uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QMargins) OperatorPlusAssign(margins *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorMinusAssign(margins *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMinusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorPlusAssignWithInt(arg1 int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlusAssign2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorMinusAssignWithInt(arg1 int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMinusAssign2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorMultiplyAssign(arg1 int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMultiplyAssign, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorDivideAssign(arg1 int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorDivideAssign, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorMultiplyAssignWith(arg1 float64) *QMargins {
	__r := pg_QMargins_OperatorMultiplyAssign2(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorDivideAssignWith(arg1 float64) *QMargins {
	__r := pg_QMargins_OperatorDivideAssign2(this.h, arg1)
	runtime.KeepAlive(this)
	return bindQMargins(__r)
}

func (this *QMargins) ToMarginsF() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMargins_toMarginsF, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMargins) OperatorAssign(arg1 *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQMargins(__r)
}

func (this *QMargins) OperatorPlus(m2 *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlus, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorMinus(m2 *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMinus, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorPlusWithRhs(rhs int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlus2, this.h, uintptr(rhs))
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorMinusWithRhs(rhs int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMinus2, this.h, uintptr(rhs))
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorMultiply(factor int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMultiply, this.h, uintptr(factor))
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorMultiplyWithFactor(factor float64) *QMargins {
	__r := pg_QMargins_OperatorMultiply2(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorDivide(divisor int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorDivide, this.h, uintptr(divisor))
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorDivideWithDivisor(divisor float64) *QMargins {
	__r := pg_QMargins_OperatorDivide2(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorBitwiseOr(m2 *QMargins) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorBitwiseOr, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorEqual(rhs *QMargins) bool {
	__r0, _, _ := purego.SyscallN(addr_QMargins_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMargins) OperatorNotEqual(rhs *QMargins) bool {
	__r0, _, _ := purego.SyscallN(addr_QMargins_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMargins) OperatorPlus2() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlus3, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorMinus2() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorMinus3, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMargins) OperatorEqualWithRhs(rhs *QMarginsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QMargins_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMargins) OperatorNotEqualWithRhs(rhs *QMarginsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QMargins_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMargins) OperatorPlusWithRectangle(rectangle *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QMargins_OperatorPlus4, this.h, uintptr(rectangle.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rectangle)
	return newownedQRect(__r)
}

func (this *QMargins) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMargins_delete(this.h)
	}
	this.UnsafeRelease()
}
