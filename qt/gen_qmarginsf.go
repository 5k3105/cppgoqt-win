package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMarginsF_new uintptr
	pg_QMarginsF_new2 func(float64, float64, float64, float64) uintptr
	addr_QMarginsF_new3 uintptr
	addr_QMarginsF_new4 uintptr
	addr_QMarginsF_isNull uintptr
	pg_QMarginsF_left func(uintptr) float64
	pg_QMarginsF_top func(uintptr) float64
	pg_QMarginsF_right func(uintptr) float64
	pg_QMarginsF_bottom func(uintptr) float64
	pg_QMarginsF_setLeft func(uintptr, float64)
	pg_QMarginsF_setTop func(uintptr, float64)
	pg_QMarginsF_setRight func(uintptr, float64)
	pg_QMarginsF_setBottom func(uintptr, float64)
	addr_QMarginsF_OperatorPlusAssign uintptr
	addr_QMarginsF_OperatorMinusAssign uintptr
	pg_QMarginsF_OperatorPlusAssign2 func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorMinusAssign2 func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorMultiplyAssign func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorDivideAssign func(uintptr, float64) uintptr
	addr_QMarginsF_toMargins uintptr
	addr_QMarginsF_OperatorAssign uintptr
	addr_QMarginsF_OperatorEqual uintptr
	addr_QMarginsF_OperatorNotEqual uintptr
	addr_QMarginsF_OperatorEqual2 uintptr
	addr_QMarginsF_OperatorNotEqual2 uintptr
	addr_QMarginsF_OperatorPlus uintptr
	addr_QMarginsF_OperatorMinus uintptr
	pg_QMarginsF_OperatorPlus2 func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorMinus2 func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorMultiply func(uintptr, float64) uintptr
	pg_QMarginsF_OperatorDivide func(uintptr, float64) uintptr
	addr_QMarginsF_OperatorBitwiseOr uintptr
	addr_QMarginsF_OperatorPlus3 uintptr
	addr_QMarginsF_OperatorMinus3 uintptr
	addr_QMarginsF_OperatorPlus4 uintptr
	pg_QMarginsF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMarginsF_new, "QMarginsF_new")
	reg(&pg_QMarginsF_new2, "QMarginsF_new2")
	symaddr(&addr_QMarginsF_new3, "QMarginsF_new3")
	symaddr(&addr_QMarginsF_new4, "QMarginsF_new4")
	symaddr(&addr_QMarginsF_isNull, "QMarginsF_isNull")
	reg(&pg_QMarginsF_left, "QMarginsF_left")
	reg(&pg_QMarginsF_top, "QMarginsF_top")
	reg(&pg_QMarginsF_right, "QMarginsF_right")
	reg(&pg_QMarginsF_bottom, "QMarginsF_bottom")
	reg(&pg_QMarginsF_setLeft, "QMarginsF_setLeft")
	reg(&pg_QMarginsF_setTop, "QMarginsF_setTop")
	reg(&pg_QMarginsF_setRight, "QMarginsF_setRight")
	reg(&pg_QMarginsF_setBottom, "QMarginsF_setBottom")
	symaddr(&addr_QMarginsF_OperatorPlusAssign, "QMarginsF_OperatorPlusAssign")
	symaddr(&addr_QMarginsF_OperatorMinusAssign, "QMarginsF_OperatorMinusAssign")
	reg(&pg_QMarginsF_OperatorPlusAssign2, "QMarginsF_OperatorPlusAssign2")
	reg(&pg_QMarginsF_OperatorMinusAssign2, "QMarginsF_OperatorMinusAssign2")
	reg(&pg_QMarginsF_OperatorMultiplyAssign, "QMarginsF_OperatorMultiplyAssign")
	reg(&pg_QMarginsF_OperatorDivideAssign, "QMarginsF_OperatorDivideAssign")
	symaddr(&addr_QMarginsF_toMargins, "QMarginsF_toMargins")
	symaddr(&addr_QMarginsF_OperatorAssign, "QMarginsF_OperatorAssign")
	symaddr(&addr_QMarginsF_OperatorEqual, "QMarginsF_OperatorEqual")
	symaddr(&addr_QMarginsF_OperatorNotEqual, "QMarginsF_OperatorNotEqual")
	symaddr(&addr_QMarginsF_OperatorEqual2, "QMarginsF_OperatorEqual2")
	symaddr(&addr_QMarginsF_OperatorNotEqual2, "QMarginsF_OperatorNotEqual2")
	symaddr(&addr_QMarginsF_OperatorPlus, "QMarginsF_OperatorPlus")
	symaddr(&addr_QMarginsF_OperatorMinus, "QMarginsF_OperatorMinus")
	reg(&pg_QMarginsF_OperatorPlus2, "QMarginsF_OperatorPlus2")
	reg(&pg_QMarginsF_OperatorMinus2, "QMarginsF_OperatorMinus2")
	reg(&pg_QMarginsF_OperatorMultiply, "QMarginsF_OperatorMultiply")
	reg(&pg_QMarginsF_OperatorDivide, "QMarginsF_OperatorDivide")
	symaddr(&addr_QMarginsF_OperatorBitwiseOr, "QMarginsF_OperatorBitwiseOr")
	symaddr(&addr_QMarginsF_OperatorPlus3, "QMarginsF_OperatorPlus3")
	symaddr(&addr_QMarginsF_OperatorMinus3, "QMarginsF_OperatorMinus3")
	symaddr(&addr_QMarginsF_OperatorPlus4, "QMarginsF_OperatorPlus4")
	reg(&pg_QMarginsF_delete, "QMarginsF_delete")
}

type QMarginsF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMarginsF) UnsafeGoID() int64     { return this.vid }
func (this *QMarginsF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMarginsF(h uintptr) *QMarginsF {
	if h == 0 {
		return nil
	}
	return &QMarginsF{h: h}
}

func newownedQMarginsF(h uintptr) *QMarginsF {
	o := bindQMarginsF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMarginsF).Delete)
	}
	return o
}

func (this *QMarginsF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMarginsF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMarginsF(h unsafe.Pointer) *QMarginsF { return bindQMarginsF(uintptr(h)) }
func UnsafeNewQMarginsFOwned(h unsafe.Pointer) *QMarginsF { return newownedQMarginsF(uintptr(h)) }

func (this *QMarginsF) UnsafeRelease() { this.h = 0 }

func NewQMarginsF() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_new)
	__o := newownedQMarginsF(__r)
	return __o
}

func NewQMarginsF2(left float64, top float64, right float64, bottom float64) *QMarginsF {
	__o := newownedQMarginsF(pg_QMarginsF_new2(left, top, right, bottom))
	return __o
}

func NewQMarginsF3(margins *QMargins) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_new3, uintptr(margins.UnsafePointer()))
	__o := newownedQMarginsF(__r)
	runtime.KeepAlive(margins)
	return __o
}

func NewQMarginsF4(arg1 *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQMarginsF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMarginsF) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QMarginsF_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMarginsF) Left() float64 {
	__r := pg_QMarginsF_left(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMarginsF) Top() float64 {
	__r := pg_QMarginsF_top(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMarginsF) Right() float64 {
	__r := pg_QMarginsF_right(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMarginsF) Bottom() float64 {
	__r := pg_QMarginsF_bottom(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMarginsF) SetLeft(aleft float64) {
	pg_QMarginsF_setLeft(this.h, aleft)
	runtime.KeepAlive(this)
}

func (this *QMarginsF) SetTop(atop float64) {
	pg_QMarginsF_setTop(this.h, atop)
	runtime.KeepAlive(this)
}

func (this *QMarginsF) SetRight(aright float64) {
	pg_QMarginsF_setRight(this.h, aright)
	runtime.KeepAlive(this)
}

func (this *QMarginsF) SetBottom(abottom float64) {
	pg_QMarginsF_setBottom(this.h, abottom)
	runtime.KeepAlive(this)
}

func (this *QMarginsF) OperatorPlusAssign(margins *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorPlusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorMinusAssign(margins *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorMinusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorPlusAssignWithAddend(addend float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorPlusAssign2(this.h, addend)
	runtime.KeepAlive(this)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorMinusAssignWithSubtrahend(subtrahend float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorMinusAssign2(this.h, subtrahend)
	runtime.KeepAlive(this)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorMultiplyAssign(factor float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorDivideAssign(divisor float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) ToMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_toMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QMarginsF) OperatorAssign(arg1 *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQMarginsF(__r)
}

func (this *QMarginsF) OperatorEqual(rhs *QMarginsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QMarginsF_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMarginsF) OperatorNotEqual(rhs *QMarginsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QMarginsF_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMarginsF) OperatorEqualWithRhs(rhs *QMargins) bool {
	__r0, _, _ := purego.SyscallN(addr_QMarginsF_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMarginsF) OperatorNotEqualWithRhs(rhs *QMargins) bool {
	__r0, _, _ := purego.SyscallN(addr_QMarginsF_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMarginsF) OperatorPlus(rhs *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorPlus, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorMinus(rhs *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorMinus, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorPlusWithRhs(rhs float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorPlus2(this.h, rhs)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorMinusWithRhs(rhs float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorMinus2(this.h, rhs)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorMultiply(rhs float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorMultiply(this.h, rhs)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorDivide(divisor float64) *QMarginsF {
	__r := pg_QMarginsF_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorBitwiseOr(m2 *QMarginsF) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorBitwiseOr, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorPlus2() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorPlus3, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorMinus2() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorMinus3, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QMarginsF) OperatorPlus3(rhs *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QMarginsF_OperatorPlus4, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQRectF(__r)
}

func (this *QMarginsF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMarginsF_delete(this.h)
	}
	this.UnsafeRelease()
}
