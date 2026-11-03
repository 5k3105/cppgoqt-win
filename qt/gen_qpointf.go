package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPointF_new uintptr
	addr_QPointF_new2 uintptr
	pg_QPointF_new3 func(float64, float64) uintptr
	addr_QPointF_new4 uintptr
	pg_QPointF_manhattanLength func(uintptr) float64
	addr_QPointF_isNull uintptr
	pg_QPointF_x func(uintptr) float64
	pg_QPointF_y func(uintptr) float64
	pg_QPointF_setX func(uintptr, float64)
	pg_QPointF_setY func(uintptr, float64)
	addr_QPointF_transposed uintptr
	addr_QPointF_OperatorPlusAssign uintptr
	addr_QPointF_OperatorMinusAssign uintptr
	pg_QPointF_OperatorMultiplyAssign func(uintptr, float64) uintptr
	pg_QPointF_OperatorDivideAssign func(uintptr, float64) uintptr
	pg_QPointF_dotProduct func(uintptr, uintptr) float64
	addr_QPointF_toPoint uintptr
	addr_QPointF_OperatorAssign uintptr
	addr_QPointF_OperatorEqual uintptr
	addr_QPointF_OperatorNotEqual uintptr
	addr_QPointF_OperatorEqual2 uintptr
	addr_QPointF_OperatorNotEqual2 uintptr
	addr_QPointF_OperatorPlus uintptr
	addr_QPointF_OperatorMinus uintptr
	pg_QPointF_OperatorMultiply func(uintptr, float64) uintptr
	addr_QPointF_OperatorPlus2 uintptr
	addr_QPointF_OperatorMinus2 uintptr
	pg_QPointF_OperatorDivide func(uintptr, float64) uintptr
	addr_QPointF_OperatorMultiply2 uintptr
	addr_QPointF_OperatorMultiply3 uintptr
	pg_QPointF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPointF_new, "QPointF_new")
	symaddr(&addr_QPointF_new2, "QPointF_new2")
	reg(&pg_QPointF_new3, "QPointF_new3")
	symaddr(&addr_QPointF_new4, "QPointF_new4")
	reg(&pg_QPointF_manhattanLength, "QPointF_manhattanLength")
	symaddr(&addr_QPointF_isNull, "QPointF_isNull")
	reg(&pg_QPointF_x, "QPointF_x")
	reg(&pg_QPointF_y, "QPointF_y")
	reg(&pg_QPointF_setX, "QPointF_setX")
	reg(&pg_QPointF_setY, "QPointF_setY")
	symaddr(&addr_QPointF_transposed, "QPointF_transposed")
	symaddr(&addr_QPointF_OperatorPlusAssign, "QPointF_OperatorPlusAssign")
	symaddr(&addr_QPointF_OperatorMinusAssign, "QPointF_OperatorMinusAssign")
	reg(&pg_QPointF_OperatorMultiplyAssign, "QPointF_OperatorMultiplyAssign")
	reg(&pg_QPointF_OperatorDivideAssign, "QPointF_OperatorDivideAssign")
	reg(&pg_QPointF_dotProduct, "QPointF_dotProduct")
	symaddr(&addr_QPointF_toPoint, "QPointF_toPoint")
	symaddr(&addr_QPointF_OperatorAssign, "QPointF_OperatorAssign")
	symaddr(&addr_QPointF_OperatorEqual, "QPointF_OperatorEqual")
	symaddr(&addr_QPointF_OperatorNotEqual, "QPointF_OperatorNotEqual")
	symaddr(&addr_QPointF_OperatorEqual2, "QPointF_OperatorEqual2")
	symaddr(&addr_QPointF_OperatorNotEqual2, "QPointF_OperatorNotEqual2")
	symaddr(&addr_QPointF_OperatorPlus, "QPointF_OperatorPlus")
	symaddr(&addr_QPointF_OperatorMinus, "QPointF_OperatorMinus")
	reg(&pg_QPointF_OperatorMultiply, "QPointF_OperatorMultiply")
	symaddr(&addr_QPointF_OperatorPlus2, "QPointF_OperatorPlus2")
	symaddr(&addr_QPointF_OperatorMinus2, "QPointF_OperatorMinus2")
	reg(&pg_QPointF_OperatorDivide, "QPointF_OperatorDivide")
	symaddr(&addr_QPointF_OperatorMultiply2, "QPointF_OperatorMultiply2")
	symaddr(&addr_QPointF_OperatorMultiply3, "QPointF_OperatorMultiply3")
	reg(&pg_QPointF_delete, "QPointF_delete")
}

type QPointF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPointF) UnsafeGoID() int64     { return this.vid }
func (this *QPointF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPointF(h uintptr) *QPointF {
	if h == 0 {
		return nil
	}
	return &QPointF{h: h}
}

func newownedQPointF(h uintptr) *QPointF {
	o := bindQPointF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPointF).Delete)
	}
	return o
}

func (this *QPointF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPointF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPointF(h unsafe.Pointer) *QPointF { return bindQPointF(uintptr(h)) }
func UnsafeNewQPointFOwned(h unsafe.Pointer) *QPointF { return newownedQPointF(uintptr(h)) }

func (this *QPointF) UnsafeRelease() { this.h = 0 }

func NewQPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_new)
	__o := newownedQPointF(__r)
	return __o
}

func NewQPointF2(p *QPoint) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_new2, uintptr(p.UnsafePointer()))
	__o := newownedQPointF(__r)
	runtime.KeepAlive(p)
	return __o
}

func NewQPointF3(xpos float64, ypos float64) *QPointF {
	__o := newownedQPointF(pg_QPointF_new3(xpos, ypos))
	return __o
}

func NewQPointF4(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQPointF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPointF) ManhattanLength() float64 {
	__r := pg_QPointF_manhattanLength(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPointF) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointF_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointF) X() float64 {
	__r := pg_QPointF_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPointF) Y() float64 {
	__r := pg_QPointF_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPointF) SetX(x float64) {
	pg_QPointF_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QPointF) SetY(y float64) {
	pg_QPointF_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QPointF) Transposed() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorPlusAssign(p *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorPlusAssign, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQPointF(__r)
}

func (this *QPointF) OperatorMinusAssign(p *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorMinusAssign, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQPointF(__r)
}

func (this *QPointF) OperatorMultiplyAssign(c float64) *QPointF {
	__r := pg_QPointF_OperatorMultiplyAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPointF) OperatorDivideAssign(c float64) *QPointF {
	__r := pg_QPointF_OperatorDivideAssign(this.h, c)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func QPointF_DotProduct(p1 *QPointF, p2 *QPointF) float64 {
	__r := pg_QPointF_dotProduct(uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
	return __r
}

func (this *QPointF) ToPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPointF_toPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPointF) OperatorAssign(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPointF(__r)
}

func (this *QPointF) OperatorEqual(rhs *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointF_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointF) OperatorNotEqual(rhs *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointF_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointF) OperatorEqualWithRhs(rhs *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointF_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointF) OperatorNotEqualWithRhs(rhs *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointF_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointF) OperatorPlus(p2 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorPlus, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorMinus(p2 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorMinus, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorMultiply(c float64) *QPointF {
	__r := pg_QPointF_OperatorMultiply(this.h, c)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorPlus2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorPlus2, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorMinus2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorDivide(divisor float64) *QPointF {
	__r := pg_QPointF_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorMultiplyWithQTransform(m *QTransform) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorMultiply2, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQPointF(__r)
}

func (this *QPointF) OperatorMultiplyWithMatrix(matrix *QMatrix4x4) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPointF_OperatorMultiply3, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQPointF(__r)
}

func (this *QPointF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPointF_delete(this.h)
	}
	this.UnsafeRelease()
}
