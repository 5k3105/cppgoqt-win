package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_qfloat16_new uintptr
	addr_qfloat16_new2 uintptr
	pg_qfloat16_new3 func(float32) uintptr
	addr_qfloat16_new4 uintptr
	pg_qfloat16_ToFloat func(uintptr) float32
	addr_qfloat16_isInf uintptr
	addr_qfloat16_isNaN uintptr
	addr_qfloat16_isFinite uintptr
	addr_qfloat16_fpClassify uintptr
	addr_qfloat16_copySign uintptr
	addr_qfloat16__limit_epsilon uintptr
	addr_qfloat16__limit_min uintptr
	addr_qfloat16__limit_denorm_min uintptr
	addr_qfloat16__limit_max uintptr
	addr_qfloat16__limit_lowest uintptr
	addr_qfloat16__limit_infinity uintptr
	addr_qfloat16__limit_quiet_NaN uintptr
	addr_qfloat16__limit_signaling_NaN uintptr
	addr_qfloat16_isNormal uintptr
	addr_qfloat16_OperatorMinus uintptr
	addr_qfloat16_OperatorPlus uintptr
	addr_qfloat16_OperatorMinus2 uintptr
	addr_qfloat16_OperatorMultiply uintptr
	addr_qfloat16_OperatorDivide uintptr
	pg_qfloat16_OperatorPlus2 func(uintptr, float64) float64
	pg_qfloat16_OperatorMinus3 func(uintptr, float64) float64
	pg_qfloat16_OperatorMultiply2 func(uintptr, float64) float64
	pg_qfloat16_OperatorDivide2 func(uintptr, float64) float64
	pg_qfloat16_OperatorPlusAssign func(uintptr, float64) uintptr
	pg_qfloat16_OperatorMinusAssign func(uintptr, float64) uintptr
	pg_qfloat16_OperatorMultiplyAssign func(uintptr, float64) uintptr
	pg_qfloat16_OperatorDivideAssign func(uintptr, float64) uintptr
	pg_qfloat16_OperatorPlus3 func(uintptr, float32) float32
	pg_qfloat16_OperatorMinus4 func(uintptr, float32) float32
	pg_qfloat16_OperatorMultiply3 func(uintptr, float32) float32
	pg_qfloat16_OperatorDivide3 func(uintptr, float32) float32
	pg_qfloat16_OperatorPlusAssign2 func(uintptr, float32) uintptr
	pg_qfloat16_OperatorMinusAssign2 func(uintptr, float32) uintptr
	pg_qfloat16_OperatorMultiplyAssign2 func(uintptr, float32) uintptr
	pg_qfloat16_OperatorDivideAssign2 func(uintptr, float32) uintptr
	pg_qfloat16_OperatorPlus4 func(uintptr, int32) float64
	pg_qfloat16_OperatorMinus5 func(uintptr, int32) float64
	pg_qfloat16_OperatorMultiply4 func(uintptr, int32) float64
	pg_qfloat16_OperatorDivide4 func(uintptr, int32) float64
	addr_qfloat16_OperatorEqual uintptr
	addr_qfloat16_OperatorNotEqual uintptr
	addr_qfloat16_OperatorLesser uintptr
	addr_qfloat16_OperatorGreater uintptr
	addr_qfloat16_OperatorLesserOrEqual uintptr
	addr_qfloat16_OperatorGreaterOrEqual uintptr
	pg_qfloat16_delete func(uintptr)
)

func init() {
	symaddr(&addr_qfloat16_new, "qfloat16_new")
	symaddr(&addr_qfloat16_new2, "qfloat16_new2")
	reg(&pg_qfloat16_new3, "qfloat16_new3")
	symaddr(&addr_qfloat16_new4, "qfloat16_new4")
	reg(&pg_qfloat16_ToFloat, "qfloat16_ToFloat")
	symaddr(&addr_qfloat16_isInf, "qfloat16_isInf")
	symaddr(&addr_qfloat16_isNaN, "qfloat16_isNaN")
	symaddr(&addr_qfloat16_isFinite, "qfloat16_isFinite")
	symaddr(&addr_qfloat16_fpClassify, "qfloat16_fpClassify")
	symaddr(&addr_qfloat16_copySign, "qfloat16_copySign")
	symaddr(&addr_qfloat16__limit_epsilon, "qfloat16__limit_epsilon")
	symaddr(&addr_qfloat16__limit_min, "qfloat16__limit_min")
	symaddr(&addr_qfloat16__limit_denorm_min, "qfloat16__limit_denorm_min")
	symaddr(&addr_qfloat16__limit_max, "qfloat16__limit_max")
	symaddr(&addr_qfloat16__limit_lowest, "qfloat16__limit_lowest")
	symaddr(&addr_qfloat16__limit_infinity, "qfloat16__limit_infinity")
	symaddr(&addr_qfloat16__limit_quiet_NaN, "qfloat16__limit_quiet_NaN")
	symaddr(&addr_qfloat16__limit_signaling_NaN, "qfloat16__limit_signaling_NaN")
	symaddr(&addr_qfloat16_isNormal, "qfloat16_isNormal")
	symaddr(&addr_qfloat16_OperatorMinus, "qfloat16_OperatorMinus")
	symaddr(&addr_qfloat16_OperatorPlus, "qfloat16_OperatorPlus")
	symaddr(&addr_qfloat16_OperatorMinus2, "qfloat16_OperatorMinus2")
	symaddr(&addr_qfloat16_OperatorMultiply, "qfloat16_OperatorMultiply")
	symaddr(&addr_qfloat16_OperatorDivide, "qfloat16_OperatorDivide")
	reg(&pg_qfloat16_OperatorPlus2, "qfloat16_OperatorPlus2")
	reg(&pg_qfloat16_OperatorMinus3, "qfloat16_OperatorMinus3")
	reg(&pg_qfloat16_OperatorMultiply2, "qfloat16_OperatorMultiply2")
	reg(&pg_qfloat16_OperatorDivide2, "qfloat16_OperatorDivide2")
	reg(&pg_qfloat16_OperatorPlusAssign, "qfloat16_OperatorPlusAssign")
	reg(&pg_qfloat16_OperatorMinusAssign, "qfloat16_OperatorMinusAssign")
	reg(&pg_qfloat16_OperatorMultiplyAssign, "qfloat16_OperatorMultiplyAssign")
	reg(&pg_qfloat16_OperatorDivideAssign, "qfloat16_OperatorDivideAssign")
	reg(&pg_qfloat16_OperatorPlus3, "qfloat16_OperatorPlus3")
	reg(&pg_qfloat16_OperatorMinus4, "qfloat16_OperatorMinus4")
	reg(&pg_qfloat16_OperatorMultiply3, "qfloat16_OperatorMultiply3")
	reg(&pg_qfloat16_OperatorDivide3, "qfloat16_OperatorDivide3")
	reg(&pg_qfloat16_OperatorPlusAssign2, "qfloat16_OperatorPlusAssign2")
	reg(&pg_qfloat16_OperatorMinusAssign2, "qfloat16_OperatorMinusAssign2")
	reg(&pg_qfloat16_OperatorMultiplyAssign2, "qfloat16_OperatorMultiplyAssign2")
	reg(&pg_qfloat16_OperatorDivideAssign2, "qfloat16_OperatorDivideAssign2")
	reg(&pg_qfloat16_OperatorPlus4, "qfloat16_OperatorPlus4")
	reg(&pg_qfloat16_OperatorMinus5, "qfloat16_OperatorMinus5")
	reg(&pg_qfloat16_OperatorMultiply4, "qfloat16_OperatorMultiply4")
	reg(&pg_qfloat16_OperatorDivide4, "qfloat16_OperatorDivide4")
	symaddr(&addr_qfloat16_OperatorEqual, "qfloat16_OperatorEqual")
	symaddr(&addr_qfloat16_OperatorNotEqual, "qfloat16_OperatorNotEqual")
	symaddr(&addr_qfloat16_OperatorLesser, "qfloat16_OperatorLesser")
	symaddr(&addr_qfloat16_OperatorGreater, "qfloat16_OperatorGreater")
	symaddr(&addr_qfloat16_OperatorLesserOrEqual, "qfloat16_OperatorLesserOrEqual")
	symaddr(&addr_qfloat16_OperatorGreaterOrEqual, "qfloat16_OperatorGreaterOrEqual")
	reg(&pg_qfloat16_delete, "qfloat16_delete")
}

type qfloat16 struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *qfloat16) UnsafeGoID() int64     { return this.vid }
func (this *qfloat16) UnsafeSetGoID(id int64) { this.vid = id }

func bindqfloat16(h uintptr) *qfloat16 {
	if h == 0 {
		return nil
	}
	return &qfloat16{h: h}
}

func newownedqfloat16(h uintptr) *qfloat16 {
	o := bindqfloat16(h)
	if o != nil {
		runtime.SetFinalizer(o, (*qfloat16).Delete)
	}
	return o
}

func (this *qfloat16) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *qfloat16) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewqfloat16(h unsafe.Pointer) *qfloat16 { return bindqfloat16(uintptr(h)) }
func UnsafeNewqfloat16Owned(h unsafe.Pointer) *qfloat16 { return newownedqfloat16(uintptr(h)) }

func (this *qfloat16) UnsafeRelease() { this.h = 0 }

func Newqfloat16() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_new)
	__o := newownedqfloat16(__r)
	return __o
}

func Newqfloat162(arg1 Initialization) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_new2, uintptr(arg1))
	__o := newownedqfloat16(__r)
	return __o
}

func Newqfloat163(f float32) *qfloat16 {
	__o := newownedqfloat16(pg_qfloat16_new3(f))
	return __o
}

func Newqfloat164(arg1 *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedqfloat16(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *qfloat16) ToFloat() float32 {
	__r := pg_qfloat16_ToFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) IsInf() bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_isInf, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *qfloat16) IsNaN() bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_isNaN, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *qfloat16) IsFinite() bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_isFinite, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *qfloat16) FpClassify() int {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_fpClassify, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *qfloat16) CopySign(sign *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_copySign, this.h, uintptr(sign.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sign)
	return newownedqfloat16(__r)
}

func qfloat16__limit_epsilon() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_epsilon)
	return newownedqfloat16(__r)
}

func qfloat16__limit_min() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_min)
	return newownedqfloat16(__r)
}

func qfloat16__limit_denorm_min() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_denorm_min)
	return newownedqfloat16(__r)
}

func qfloat16__limit_max() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_max)
	return newownedqfloat16(__r)
}

func qfloat16__limit_lowest() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_lowest)
	return newownedqfloat16(__r)
}

func qfloat16__limit_infinity() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_infinity)
	return newownedqfloat16(__r)
}

func qfloat16__limit_quiet_NaN() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_quiet_NaN)
	return newownedqfloat16(__r)
}

func qfloat16__limit_signaling_NaN() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16__limit_signaling_NaN)
	return newownedqfloat16(__r)
}

func (this *qfloat16) IsNormal() bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_isNormal, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *qfloat16) OperatorMinus() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_OperatorMinus, this.h)
	runtime.KeepAlive(this)
	return newownedqfloat16(__r)
}

func (this *qfloat16) OperatorPlus(b *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_OperatorPlus, this.h, uintptr(b.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return newownedqfloat16(__r)
}

func (this *qfloat16) OperatorMinusWithQfloat16(b *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_OperatorMinus2, this.h, uintptr(b.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return newownedqfloat16(__r)
}

func (this *qfloat16) OperatorMultiply(b *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_OperatorMultiply, this.h, uintptr(b.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return newownedqfloat16(__r)
}

func (this *qfloat16) OperatorDivide(b *qfloat16) *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_qfloat16_OperatorDivide, this.h, uintptr(b.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return newownedqfloat16(__r)
}

func (this *qfloat16) OperatorPlusWithRhs(rhs float64) float64 {
	__r := pg_qfloat16_OperatorPlus2(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMinusWithRhs(rhs float64) float64 {
	__r := pg_qfloat16_OperatorMinus3(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMultiplyWithRhs(rhs float64) float64 {
	__r := pg_qfloat16_OperatorMultiply2(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorDivideWithRhs(rhs float64) float64 {
	__r := pg_qfloat16_OperatorDivide2(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorPlusAssign(rhs float64) *qfloat16 {
	__r := pg_qfloat16_OperatorPlusAssign(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorMinusAssign(rhs float64) *qfloat16 {
	__r := pg_qfloat16_OperatorMinusAssign(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorMultiplyAssign(rhs float64) *qfloat16 {
	__r := pg_qfloat16_OperatorMultiplyAssign(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorDivideAssign(rhs float64) *qfloat16 {
	__r := pg_qfloat16_OperatorDivideAssign(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorPlus2(rhs float32) float32 {
	__r := pg_qfloat16_OperatorPlus3(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMinus2(rhs float32) float32 {
	__r := pg_qfloat16_OperatorMinus4(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMultiply2(rhs float32) float32 {
	__r := pg_qfloat16_OperatorMultiply3(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorDivide2(rhs float32) float32 {
	__r := pg_qfloat16_OperatorDivide3(this.h, rhs)
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorPlusAssignWithRhs(rhs float32) *qfloat16 {
	__r := pg_qfloat16_OperatorPlusAssign2(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorMinusAssignWithRhs(rhs float32) *qfloat16 {
	__r := pg_qfloat16_OperatorMinusAssign2(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorMultiplyAssignWithRhs(rhs float32) *qfloat16 {
	__r := pg_qfloat16_OperatorMultiplyAssign2(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorDivideAssignWithRhs(rhs float32) *qfloat16 {
	__r := pg_qfloat16_OperatorDivideAssign2(this.h, rhs)
	runtime.KeepAlive(this)
	return bindqfloat16(__r)
}

func (this *qfloat16) OperatorPlus3(rhs int) float64 {
	__r := pg_qfloat16_OperatorPlus4(this.h, int32(rhs))
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMinus3(rhs int) float64 {
	__r := pg_qfloat16_OperatorMinus5(this.h, int32(rhs))
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorMultiply3(rhs int) float64 {
	__r := pg_qfloat16_OperatorMultiply4(this.h, int32(rhs))
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorDivide3(rhs int) float64 {
	__r := pg_qfloat16_OperatorDivide4(this.h, int32(rhs))
	runtime.KeepAlive(this)
	return __r
}

func (this *qfloat16) OperatorEqual(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) OperatorNotEqual(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) OperatorLesser(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) OperatorGreater(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) OperatorLesserOrEqual(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) OperatorGreaterOrEqual(rhs *qfloat16) bool {
	__r0, _, _ := purego.SyscallN(addr_qfloat16_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *qfloat16) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_qfloat16_delete(this.h)
	}
	this.UnsafeRelease()
}
