package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPoint_new uintptr
	addr_QPoint_new2 uintptr
	addr_QPoint_new3 uintptr
	addr_QPoint_isNull uintptr
	addr_QPoint_x uintptr
	addr_QPoint_y uintptr
	addr_QPoint_setX uintptr
	addr_QPoint_setY uintptr
	addr_QPoint_manhattanLength uintptr
	addr_QPoint_transposed uintptr
	addr_QPoint_OperatorPlusAssign uintptr
	addr_QPoint_OperatorMinusAssign uintptr
	pg_QPoint_OperatorMultiplyAssign func(uintptr, float32) uintptr
	pg_QPoint_OperatorMultiplyAssign2 func(uintptr, float64) uintptr
	addr_QPoint_OperatorMultiplyAssign3 uintptr
	pg_QPoint_OperatorDivideAssign func(uintptr, float64) uintptr
	addr_QPoint_dotProduct uintptr
	addr_QPoint_toPointF uintptr
	addr_QPoint_OperatorAssign uintptr
	addr_QPoint_OperatorEqual uintptr
	addr_QPoint_OperatorNotEqual uintptr
	addr_QPoint_OperatorPlus uintptr
	addr_QPoint_OperatorMinus uintptr
	pg_QPoint_OperatorMultiply func(uintptr, float32) uintptr
	pg_QPoint_OperatorMultiply2 func(uintptr, float64) uintptr
	addr_QPoint_OperatorMultiply3 uintptr
	addr_QPoint_OperatorPlus2 uintptr
	addr_QPoint_OperatorMinus2 uintptr
	pg_QPoint_OperatorDivide func(uintptr, float64) uintptr
	addr_QPoint_OperatorEqual2 uintptr
	addr_QPoint_OperatorNotEqual2 uintptr
	addr_QPoint_OperatorMultiply4 uintptr
	addr_QPoint_OperatorMultiply5 uintptr
	pg_QPoint_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPoint_new, "QPoint_new")
	symaddr(&addr_QPoint_new2, "QPoint_new2")
	symaddr(&addr_QPoint_new3, "QPoint_new3")
	symaddr(&addr_QPoint_isNull, "QPoint_isNull")
	symaddr(&addr_QPoint_x, "QPoint_x")
	symaddr(&addr_QPoint_y, "QPoint_y")
	symaddr(&addr_QPoint_setX, "QPoint_setX")
	symaddr(&addr_QPoint_setY, "QPoint_setY")
	symaddr(&addr_QPoint_manhattanLength, "QPoint_manhattanLength")
	symaddr(&addr_QPoint_transposed, "QPoint_transposed")
	symaddr(&addr_QPoint_OperatorPlusAssign, "QPoint_OperatorPlusAssign")
	symaddr(&addr_QPoint_OperatorMinusAssign, "QPoint_OperatorMinusAssign")
	reg(&pg_QPoint_OperatorMultiplyAssign, "QPoint_OperatorMultiplyAssign")
	reg(&pg_QPoint_OperatorMultiplyAssign2, "QPoint_OperatorMultiplyAssign2")
	symaddr(&addr_QPoint_OperatorMultiplyAssign3, "QPoint_OperatorMultiplyAssign3")
	reg(&pg_QPoint_OperatorDivideAssign, "QPoint_OperatorDivideAssign")
	symaddr(&addr_QPoint_dotProduct, "QPoint_dotProduct")
	symaddr(&addr_QPoint_toPointF, "QPoint_toPointF")
	symaddr(&addr_QPoint_OperatorAssign, "QPoint_OperatorAssign")
	symaddr(&addr_QPoint_OperatorEqual, "QPoint_OperatorEqual")
	symaddr(&addr_QPoint_OperatorNotEqual, "QPoint_OperatorNotEqual")
	symaddr(&addr_QPoint_OperatorPlus, "QPoint_OperatorPlus")
	symaddr(&addr_QPoint_OperatorMinus, "QPoint_OperatorMinus")
	reg(&pg_QPoint_OperatorMultiply, "QPoint_OperatorMultiply")
	reg(&pg_QPoint_OperatorMultiply2, "QPoint_OperatorMultiply2")
	symaddr(&addr_QPoint_OperatorMultiply3, "QPoint_OperatorMultiply3")
	symaddr(&addr_QPoint_OperatorPlus2, "QPoint_OperatorPlus2")
	symaddr(&addr_QPoint_OperatorMinus2, "QPoint_OperatorMinus2")
	reg(&pg_QPoint_OperatorDivide, "QPoint_OperatorDivide")
	symaddr(&addr_QPoint_OperatorEqual2, "QPoint_OperatorEqual2")
	symaddr(&addr_QPoint_OperatorNotEqual2, "QPoint_OperatorNotEqual2")
	symaddr(&addr_QPoint_OperatorMultiply4, "QPoint_OperatorMultiply4")
	symaddr(&addr_QPoint_OperatorMultiply5, "QPoint_OperatorMultiply5")
	reg(&pg_QPoint_delete, "QPoint_delete")
}

type QPoint struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPoint) UnsafeGoID() int64     { return this.vid }
func (this *QPoint) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPoint(h uintptr) *QPoint {
	if h == 0 {
		return nil
	}
	return &QPoint{h: h}
}

func newownedQPoint(h uintptr) *QPoint {
	o := bindQPoint(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPoint).Delete)
	}
	return o
}

func (this *QPoint) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPoint) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPoint(h unsafe.Pointer) *QPoint { return bindQPoint(uintptr(h)) }
func UnsafeNewQPointOwned(h unsafe.Pointer) *QPoint { return newownedQPoint(uintptr(h)) }

func (this *QPoint) UnsafeRelease() { this.h = 0 }

func NewQPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_new)
	__o := newownedQPoint(__r)
	return __o
}

func NewQPoint2(xpos int, ypos int) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_new2, uintptr(xpos), uintptr(ypos))
	__o := newownedQPoint(__r)
	return __o
}

func NewQPoint3(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQPoint(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPoint) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QPoint_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPoint) X() int {
	__r0, _, _ := purego.SyscallN(addr_QPoint_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPoint) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QPoint_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPoint) SetX(x int) {
	purego.SyscallN(addr_QPoint_setX, this.h, uintptr(x))
	runtime.KeepAlive(this)
}

func (this *QPoint) SetY(y int) {
	purego.SyscallN(addr_QPoint_setY, this.h, uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QPoint) ManhattanLength() int {
	__r0, _, _ := purego.SyscallN(addr_QPoint_manhattanLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPoint) Transposed() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorPlusAssign(p *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorPlusAssign, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorMinusAssign(p *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMinusAssign, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorMultiplyAssign(factor float32) *QPoint {
	__r := pg_QPoint_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorMultiplyAssignWithFactor(factor float64) *QPoint {
	__r := pg_QPoint_OperatorMultiplyAssign2(this.h, factor)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorMultiplyAssign2(factor int) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMultiplyAssign3, this.h, uintptr(factor))
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorDivideAssign(divisor float64) *QPoint {
	__r := pg_QPoint_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func QPoint_DotProduct(p1 *QPoint, p2 *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QPoint_dotProduct, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
	return int(__r)
}

func (this *QPoint) ToPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPoint_toPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPoint) OperatorAssign(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPoint(__r)
}

func (this *QPoint) OperatorEqual(rhs *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPoint_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPoint) OperatorNotEqual(rhs *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPoint_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPoint) OperatorPlus(p2 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorPlus, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMinus(p2 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMinus, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMultiply(factor float32) *QPoint {
	__r := pg_QPoint_OperatorMultiply(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMultiplyWithFactor(factor float64) *QPoint {
	__r := pg_QPoint_OperatorMultiply2(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMultiply2(factor int) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMultiply3, this.h, uintptr(factor))
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorPlus2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorPlus2, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMinus2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorDivide(c float64) *QPoint {
	__r := pg_QPoint_OperatorDivide(this.h, c)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorEqualWithRhs(rhs *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPoint_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPoint) OperatorNotEqualWithRhs(rhs *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPoint_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPoint) OperatorMultiplyWithQTransform(m *QTransform) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMultiply4, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQPoint(__r)
}

func (this *QPoint) OperatorMultiplyWithMatrix(matrix *QMatrix4x4) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPoint_OperatorMultiply5, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQPoint(__r)
}

func (this *QPoint) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPoint_delete(this.h)
	}
	this.UnsafeRelease()
}
