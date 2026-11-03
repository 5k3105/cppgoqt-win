package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVector2D_new uintptr
	addr_QVector2D_new2 uintptr
	pg_QVector2D_new3 func(float32, float32) uintptr
	addr_QVector2D_new4 uintptr
	addr_QVector2D_new5 uintptr
	addr_QVector2D_new6 uintptr
	addr_QVector2D_new7 uintptr
	addr_QVector2D_new8 uintptr
	addr_QVector2D_isNull uintptr
	pg_QVector2D_x func(uintptr) float32
	pg_QVector2D_y func(uintptr) float32
	pg_QVector2D_setX func(uintptr, float32)
	pg_QVector2D_setY func(uintptr, float32)
	pg_QVector2D_OperatorSubscript func(uintptr, int32) float32
	pg_QVector2D_length func(uintptr) float32
	pg_QVector2D_lengthSquared func(uintptr) float32
	addr_QVector2D_normalized uintptr
	addr_QVector2D_normalize uintptr
	pg_QVector2D_distanceToPoint func(uintptr, uintptr) float32
	pg_QVector2D_distanceToLine func(uintptr, uintptr, uintptr) float32
	addr_QVector2D_OperatorPlusAssign uintptr
	addr_QVector2D_OperatorMinusAssign uintptr
	pg_QVector2D_OperatorMultiplyAssign func(uintptr, float32) uintptr
	addr_QVector2D_OperatorMultiplyAssign2 uintptr
	pg_QVector2D_OperatorDivideAssign func(uintptr, float32) uintptr
	addr_QVector2D_OperatorDivideAssign2 uintptr
	pg_QVector2D_dotProduct func(uintptr, uintptr) float32
	addr_QVector2D_toVector3D uintptr
	addr_QVector2D_toVector4D uintptr
	addr_QVector2D_toPoint uintptr
	addr_QVector2D_toPointF uintptr
	addr_QVector2D_ToQVariant uintptr
	addr_QVector2D_OperatorAssign uintptr
	addr_QVector2D_OperatorEqual uintptr
	addr_QVector2D_OperatorNotEqual uintptr
	addr_QVector2D_OperatorPlus uintptr
	addr_QVector2D_OperatorMinus uintptr
	pg_QVector2D_OperatorMultiply func(uintptr, float32) uintptr
	addr_QVector2D_OperatorMultiply2 uintptr
	addr_QVector2D_OperatorMinus2 uintptr
	pg_QVector2D_OperatorDivide func(uintptr, float32) uintptr
	addr_QVector2D_OperatorDivide2 uintptr
	pg_QVector2D_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVector2D_new, "QVector2D_new")
	symaddr(&addr_QVector2D_new2, "QVector2D_new2")
	reg(&pg_QVector2D_new3, "QVector2D_new3")
	symaddr(&addr_QVector2D_new4, "QVector2D_new4")
	symaddr(&addr_QVector2D_new5, "QVector2D_new5")
	symaddr(&addr_QVector2D_new6, "QVector2D_new6")
	symaddr(&addr_QVector2D_new7, "QVector2D_new7")
	symaddr(&addr_QVector2D_new8, "QVector2D_new8")
	symaddr(&addr_QVector2D_isNull, "QVector2D_isNull")
	reg(&pg_QVector2D_x, "QVector2D_x")
	reg(&pg_QVector2D_y, "QVector2D_y")
	reg(&pg_QVector2D_setX, "QVector2D_setX")
	reg(&pg_QVector2D_setY, "QVector2D_setY")
	reg(&pg_QVector2D_OperatorSubscript, "QVector2D_OperatorSubscript")
	reg(&pg_QVector2D_length, "QVector2D_length")
	reg(&pg_QVector2D_lengthSquared, "QVector2D_lengthSquared")
	symaddr(&addr_QVector2D_normalized, "QVector2D_normalized")
	symaddr(&addr_QVector2D_normalize, "QVector2D_normalize")
	reg(&pg_QVector2D_distanceToPoint, "QVector2D_distanceToPoint")
	reg(&pg_QVector2D_distanceToLine, "QVector2D_distanceToLine")
	symaddr(&addr_QVector2D_OperatorPlusAssign, "QVector2D_OperatorPlusAssign")
	symaddr(&addr_QVector2D_OperatorMinusAssign, "QVector2D_OperatorMinusAssign")
	reg(&pg_QVector2D_OperatorMultiplyAssign, "QVector2D_OperatorMultiplyAssign")
	symaddr(&addr_QVector2D_OperatorMultiplyAssign2, "QVector2D_OperatorMultiplyAssign2")
	reg(&pg_QVector2D_OperatorDivideAssign, "QVector2D_OperatorDivideAssign")
	symaddr(&addr_QVector2D_OperatorDivideAssign2, "QVector2D_OperatorDivideAssign2")
	reg(&pg_QVector2D_dotProduct, "QVector2D_dotProduct")
	symaddr(&addr_QVector2D_toVector3D, "QVector2D_toVector3D")
	symaddr(&addr_QVector2D_toVector4D, "QVector2D_toVector4D")
	symaddr(&addr_QVector2D_toPoint, "QVector2D_toPoint")
	symaddr(&addr_QVector2D_toPointF, "QVector2D_toPointF")
	symaddr(&addr_QVector2D_ToQVariant, "QVector2D_ToQVariant")
	symaddr(&addr_QVector2D_OperatorAssign, "QVector2D_OperatorAssign")
	symaddr(&addr_QVector2D_OperatorEqual, "QVector2D_OperatorEqual")
	symaddr(&addr_QVector2D_OperatorNotEqual, "QVector2D_OperatorNotEqual")
	symaddr(&addr_QVector2D_OperatorPlus, "QVector2D_OperatorPlus")
	symaddr(&addr_QVector2D_OperatorMinus, "QVector2D_OperatorMinus")
	reg(&pg_QVector2D_OperatorMultiply, "QVector2D_OperatorMultiply")
	symaddr(&addr_QVector2D_OperatorMultiply2, "QVector2D_OperatorMultiply2")
	symaddr(&addr_QVector2D_OperatorMinus2, "QVector2D_OperatorMinus2")
	reg(&pg_QVector2D_OperatorDivide, "QVector2D_OperatorDivide")
	symaddr(&addr_QVector2D_OperatorDivide2, "QVector2D_OperatorDivide2")
	reg(&pg_QVector2D_delete, "QVector2D_delete")
}

type QVector2D struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVector2D) UnsafeGoID() int64     { return this.vid }
func (this *QVector2D) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVector2D(h uintptr) *QVector2D {
	if h == 0 {
		return nil
	}
	return &QVector2D{h: h}
}

func newownedQVector2D(h uintptr) *QVector2D {
	o := bindQVector2D(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVector2D).Delete)
	}
	return o
}

func (this *QVector2D) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVector2D) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVector2D(h unsafe.Pointer) *QVector2D { return bindQVector2D(uintptr(h)) }
func UnsafeNewQVector2DOwned(h unsafe.Pointer) *QVector2D { return newownedQVector2D(uintptr(h)) }

func (this *QVector2D) UnsafeRelease() { this.h = 0 }

func NewQVector2D() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new)
	__o := newownedQVector2D(__r)
	return __o
}

func NewQVector2D2(arg1 Initialization) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new2, uintptr(arg1))
	__o := newownedQVector2D(__r)
	return __o
}

func NewQVector2D3(xpos float32, ypos float32) *QVector2D {
	__o := newownedQVector2D(pg_QVector2D_new3(xpos, ypos))
	return __o
}

func NewQVector2D4(point *QPoint) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new4, uintptr(point.UnsafePointer()))
	__o := newownedQVector2D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector2D5(point *QPointF) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new5, uintptr(point.UnsafePointer()))
	__o := newownedQVector2D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector2D6(vector *QVector3D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new6, uintptr(vector.UnsafePointer()))
	__o := newownedQVector2D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector2D7(vector *QVector4D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new7, uintptr(vector.UnsafePointer()))
	__o := newownedQVector2D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector2D8(arg1 *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_new8, uintptr(arg1.UnsafePointer()))
	__o := newownedQVector2D(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QVector2D) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QVector2D_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVector2D) X() float32 {
	__r := pg_QVector2D_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector2D) Y() float32 {
	__r := pg_QVector2D_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector2D) SetX(x float32) {
	pg_QVector2D_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QVector2D) SetY(y float32) {
	pg_QVector2D_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QVector2D) OperatorSubscript(i int) float32 {
	__r := pg_QVector2D_OperatorSubscript(this.h, int32(i))
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector2D) Length() float32 {
	__r := pg_QVector2D_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector2D) LengthSquared() float32 {
	__r := pg_QVector2D_lengthSquared(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector2D) Normalized() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector2D) Normalize() {
	purego.SyscallN(addr_QVector2D_normalize, this.h)
	runtime.KeepAlive(this)
}

func (this *QVector2D) DistanceToPoint(point *QVector2D) float32 {
	__r := pg_QVector2D_distanceToPoint(this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r
}

func (this *QVector2D) DistanceToLine(point *QVector2D, direction *QVector2D) float32 {
	__r := pg_QVector2D_distanceToLine(this.h, uintptr(point.UnsafePointer()), uintptr(direction.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(direction)
	return __r
}

func (this *QVector2D) OperatorPlusAssign(vector *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorPlusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorMinusAssign(vector *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorMinusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorMultiplyAssign(factor float32) *QVector2D {
	__r := pg_QVector2D_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorMultiplyAssignWithVector(vector *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorMultiplyAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorDivideAssign(divisor float32) *QVector2D {
	__r := pg_QVector2D_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorDivideAssignWithVector(vector *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorDivideAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector2D(__r)
}

func QVector2D_DotProduct(v1 *QVector2D, v2 *QVector2D) float32 {
	__r := pg_QVector2D_dotProduct(uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return __r
}

func (this *QVector2D) ToVector3D() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_toVector3D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector2D) ToVector4D() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_toVector4D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector2D) ToPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QVector2D_toPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QVector2D) ToPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QVector2D_toPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QVector2D) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVector2D_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVector2D) OperatorAssign(arg1 *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQVector2D(__r)
}

func (this *QVector2D) OperatorEqual(v2 *QVector2D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector2D_OperatorEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector2D) OperatorNotEqual(v2 *QVector2D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector2D_OperatorNotEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector2D) OperatorPlus(v2 *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorPlus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorMinus(v2 *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorMinus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorMultiply(factor float32) *QVector2D {
	__r := pg_QVector2D_OperatorMultiply(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorMultiplyWithV2(v2 *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorMultiply2, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorMinus2() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorDivide(divisor float32) *QVector2D {
	__r := pg_QVector2D_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector2D) OperatorDivideWithDivisor(divisor *QVector2D) *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector2D_OperatorDivide2, this.h, uintptr(divisor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(divisor)
	return newownedQVector2D(__r)
}

func (this *QVector2D) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVector2D_delete(this.h)
	}
	this.UnsafeRelease()
}
