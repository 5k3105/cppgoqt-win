package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVector4D_new uintptr
	addr_QVector4D_new2 uintptr
	pg_QVector4D_new3 func(float32, float32, float32, float32) uintptr
	addr_QVector4D_new4 uintptr
	addr_QVector4D_new5 uintptr
	addr_QVector4D_new6 uintptr
	pg_QVector4D_new7 func(uintptr, float32, float32) uintptr
	addr_QVector4D_new8 uintptr
	pg_QVector4D_new9 func(uintptr, float32) uintptr
	addr_QVector4D_new10 uintptr
	addr_QVector4D_isNull uintptr
	pg_QVector4D_x func(uintptr) float32
	pg_QVector4D_y func(uintptr) float32
	pg_QVector4D_z func(uintptr) float32
	pg_QVector4D_w func(uintptr) float32
	pg_QVector4D_setX func(uintptr, float32)
	pg_QVector4D_setY func(uintptr, float32)
	pg_QVector4D_setZ func(uintptr, float32)
	pg_QVector4D_setW func(uintptr, float32)
	pg_QVector4D_OperatorSubscript func(uintptr, int32) float32
	pg_QVector4D_length func(uintptr) float32
	pg_QVector4D_lengthSquared func(uintptr) float32
	addr_QVector4D_normalized uintptr
	addr_QVector4D_normalize uintptr
	addr_QVector4D_OperatorPlusAssign uintptr
	addr_QVector4D_OperatorMinusAssign uintptr
	pg_QVector4D_OperatorMultiplyAssign func(uintptr, float32) uintptr
	addr_QVector4D_OperatorMultiplyAssign2 uintptr
	pg_QVector4D_OperatorDivideAssign func(uintptr, float32) uintptr
	addr_QVector4D_OperatorDivideAssign2 uintptr
	pg_QVector4D_dotProduct func(uintptr, uintptr) float32
	addr_QVector4D_toVector2D uintptr
	addr_QVector4D_toVector2DAffine uintptr
	addr_QVector4D_toVector3D uintptr
	addr_QVector4D_toVector3DAffine uintptr
	addr_QVector4D_toPoint uintptr
	addr_QVector4D_toPointF uintptr
	addr_QVector4D_ToQVariant uintptr
	addr_QVector4D_OperatorEqual uintptr
	addr_QVector4D_OperatorNotEqual uintptr
	addr_QVector4D_OperatorPlus uintptr
	addr_QVector4D_OperatorMinus uintptr
	pg_QVector4D_OperatorMultiply func(uintptr, float32) uintptr
	addr_QVector4D_OperatorMultiply2 uintptr
	addr_QVector4D_OperatorMinus2 uintptr
	pg_QVector4D_OperatorDivide func(uintptr, float32) uintptr
	addr_QVector4D_OperatorDivide2 uintptr
	addr_QVector4D_OperatorMultiply3 uintptr
	pg_QVector4D_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVector4D_new, "QVector4D_new")
	symaddr(&addr_QVector4D_new2, "QVector4D_new2")
	reg(&pg_QVector4D_new3, "QVector4D_new3")
	symaddr(&addr_QVector4D_new4, "QVector4D_new4")
	symaddr(&addr_QVector4D_new5, "QVector4D_new5")
	symaddr(&addr_QVector4D_new6, "QVector4D_new6")
	reg(&pg_QVector4D_new7, "QVector4D_new7")
	symaddr(&addr_QVector4D_new8, "QVector4D_new8")
	reg(&pg_QVector4D_new9, "QVector4D_new9")
	symaddr(&addr_QVector4D_new10, "QVector4D_new10")
	symaddr(&addr_QVector4D_isNull, "QVector4D_isNull")
	reg(&pg_QVector4D_x, "QVector4D_x")
	reg(&pg_QVector4D_y, "QVector4D_y")
	reg(&pg_QVector4D_z, "QVector4D_z")
	reg(&pg_QVector4D_w, "QVector4D_w")
	reg(&pg_QVector4D_setX, "QVector4D_setX")
	reg(&pg_QVector4D_setY, "QVector4D_setY")
	reg(&pg_QVector4D_setZ, "QVector4D_setZ")
	reg(&pg_QVector4D_setW, "QVector4D_setW")
	reg(&pg_QVector4D_OperatorSubscript, "QVector4D_OperatorSubscript")
	reg(&pg_QVector4D_length, "QVector4D_length")
	reg(&pg_QVector4D_lengthSquared, "QVector4D_lengthSquared")
	symaddr(&addr_QVector4D_normalized, "QVector4D_normalized")
	symaddr(&addr_QVector4D_normalize, "QVector4D_normalize")
	symaddr(&addr_QVector4D_OperatorPlusAssign, "QVector4D_OperatorPlusAssign")
	symaddr(&addr_QVector4D_OperatorMinusAssign, "QVector4D_OperatorMinusAssign")
	reg(&pg_QVector4D_OperatorMultiplyAssign, "QVector4D_OperatorMultiplyAssign")
	symaddr(&addr_QVector4D_OperatorMultiplyAssign2, "QVector4D_OperatorMultiplyAssign2")
	reg(&pg_QVector4D_OperatorDivideAssign, "QVector4D_OperatorDivideAssign")
	symaddr(&addr_QVector4D_OperatorDivideAssign2, "QVector4D_OperatorDivideAssign2")
	reg(&pg_QVector4D_dotProduct, "QVector4D_dotProduct")
	symaddr(&addr_QVector4D_toVector2D, "QVector4D_toVector2D")
	symaddr(&addr_QVector4D_toVector2DAffine, "QVector4D_toVector2DAffine")
	symaddr(&addr_QVector4D_toVector3D, "QVector4D_toVector3D")
	symaddr(&addr_QVector4D_toVector3DAffine, "QVector4D_toVector3DAffine")
	symaddr(&addr_QVector4D_toPoint, "QVector4D_toPoint")
	symaddr(&addr_QVector4D_toPointF, "QVector4D_toPointF")
	symaddr(&addr_QVector4D_ToQVariant, "QVector4D_ToQVariant")
	symaddr(&addr_QVector4D_OperatorEqual, "QVector4D_OperatorEqual")
	symaddr(&addr_QVector4D_OperatorNotEqual, "QVector4D_OperatorNotEqual")
	symaddr(&addr_QVector4D_OperatorPlus, "QVector4D_OperatorPlus")
	symaddr(&addr_QVector4D_OperatorMinus, "QVector4D_OperatorMinus")
	reg(&pg_QVector4D_OperatorMultiply, "QVector4D_OperatorMultiply")
	symaddr(&addr_QVector4D_OperatorMultiply2, "QVector4D_OperatorMultiply2")
	symaddr(&addr_QVector4D_OperatorMinus2, "QVector4D_OperatorMinus2")
	reg(&pg_QVector4D_OperatorDivide, "QVector4D_OperatorDivide")
	symaddr(&addr_QVector4D_OperatorDivide2, "QVector4D_OperatorDivide2")
	symaddr(&addr_QVector4D_OperatorMultiply3, "QVector4D_OperatorMultiply3")
	reg(&pg_QVector4D_delete, "QVector4D_delete")
}

type QVector4D struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVector4D) UnsafeGoID() int64     { return this.vid }
func (this *QVector4D) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVector4D(h uintptr) *QVector4D {
	if h == 0 {
		return nil
	}
	return &QVector4D{h: h}
}

func newownedQVector4D(h uintptr) *QVector4D {
	o := bindQVector4D(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVector4D).Delete)
	}
	return o
}

func (this *QVector4D) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVector4D) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVector4D(h unsafe.Pointer) *QVector4D { return bindQVector4D(uintptr(h)) }
func UnsafeNewQVector4DOwned(h unsafe.Pointer) *QVector4D { return newownedQVector4D(uintptr(h)) }

func (this *QVector4D) UnsafeRelease() { this.h = 0 }

func NewQVector4D() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new)
	__o := newownedQVector4D(__r)
	return __o
}

func NewQVector4D2(arg1 Initialization) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new2, uintptr(arg1))
	__o := newownedQVector4D(__r)
	return __o
}

func NewQVector4D3(xpos float32, ypos float32, zpos float32, wpos float32) *QVector4D {
	__o := newownedQVector4D(pg_QVector4D_new3(xpos, ypos, zpos, wpos))
	return __o
}

func NewQVector4D4(point *QPoint) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new4, uintptr(point.UnsafePointer()))
	__o := newownedQVector4D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector4D5(point *QPointF) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new5, uintptr(point.UnsafePointer()))
	__o := newownedQVector4D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector4D6(vector *QVector2D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new6, uintptr(vector.UnsafePointer()))
	__o := newownedQVector4D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector4D7(vector *QVector2D, zpos float32, wpos float32) *QVector4D {
	__o := newownedQVector4D(pg_QVector4D_new7(uintptr(vector.UnsafePointer()), zpos, wpos))
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector4D8(vector *QVector3D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new8, uintptr(vector.UnsafePointer()))
	__o := newownedQVector4D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector4D9(vector *QVector3D, wpos float32) *QVector4D {
	__o := newownedQVector4D(pg_QVector4D_new9(uintptr(vector.UnsafePointer()), wpos))
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector4D10(arg1 *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_new10, uintptr(arg1.UnsafePointer()))
	__o := newownedQVector4D(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QVector4D) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QVector4D_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVector4D) X() float32 {
	__r := pg_QVector4D_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) Y() float32 {
	__r := pg_QVector4D_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) Z() float32 {
	__r := pg_QVector4D_z(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) W() float32 {
	__r := pg_QVector4D_w(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) SetX(x float32) {
	pg_QVector4D_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QVector4D) SetY(y float32) {
	pg_QVector4D_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QVector4D) SetZ(z float32) {
	pg_QVector4D_setZ(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QVector4D) SetW(w float32) {
	pg_QVector4D_setW(this.h, w)
	runtime.KeepAlive(this)
}

func (this *QVector4D) OperatorSubscript(i int) float32 {
	__r := pg_QVector4D_OperatorSubscript(this.h, int32(i))
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) Length() float32 {
	__r := pg_QVector4D_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) LengthSquared() float32 {
	__r := pg_QVector4D_lengthSquared(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector4D) Normalized() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector4D) Normalize() {
	purego.SyscallN(addr_QVector4D_normalize, this.h)
	runtime.KeepAlive(this)
}

func (this *QVector4D) OperatorPlusAssign(vector *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorPlusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector4D(__r)
}

func (this *QVector4D) OperatorMinusAssign(vector *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMinusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector4D(__r)
}

func (this *QVector4D) OperatorMultiplyAssign(factor float32) *QVector4D {
	__r := pg_QVector4D_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQVector4D(__r)
}

func (this *QVector4D) OperatorMultiplyAssignWithVector(vector *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMultiplyAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector4D(__r)
}

func (this *QVector4D) OperatorDivideAssign(divisor float32) *QVector4D {
	__r := pg_QVector4D_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQVector4D(__r)
}

func (this *QVector4D) OperatorDivideAssignWithVector(vector *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorDivideAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector4D(__r)
}

func QVector4D_DotProduct(v1 *QVector4D, v2 *QVector4D) float32 {
	__r := pg_QVector4D_dotProduct(uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return __r
}

func (this *QVector4D) ToVector2D() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toVector2D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector4D) ToVector2DAffine() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toVector2DAffine, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector4D) ToVector3D() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toVector3D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector4D) ToVector3DAffine() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toVector3DAffine, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector4D) ToPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QVector4D) ToPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QVector4D_toPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QVector4D) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVector4D_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVector4D) OperatorEqual(v2 *QVector4D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector4D_OperatorEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector4D) OperatorNotEqual(v2 *QVector4D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector4D_OperatorNotEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector4D) OperatorPlus(v2 *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorPlus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorMinus(v2 *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMinus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorMultiply(factor float32) *QVector4D {
	__r := pg_QVector4D_OperatorMultiply(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorMultiplyWithV2(v2 *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMultiply2, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorMinus2() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorDivide(divisor float32) *QVector4D {
	__r := pg_QVector4D_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorDivideWithDivisor(divisor *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorDivide2, this.h, uintptr(divisor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(divisor)
	return newownedQVector4D(__r)
}

func (this *QVector4D) OperatorMultiplyWithMatrix(matrix *QMatrix4x4) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector4D_OperatorMultiply3, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQVector4D(__r)
}

func (this *QVector4D) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVector4D_delete(this.h)
	}
	this.UnsafeRelease()
}
