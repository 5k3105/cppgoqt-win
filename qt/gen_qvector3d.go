package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVector3D_new uintptr
	addr_QVector3D_new2 uintptr
	pg_QVector3D_new3 func(float32, float32, float32) uintptr
	addr_QVector3D_new4 uintptr
	addr_QVector3D_new5 uintptr
	addr_QVector3D_new6 uintptr
	pg_QVector3D_new7 func(uintptr, float32) uintptr
	addr_QVector3D_new8 uintptr
	addr_QVector3D_new9 uintptr
	addr_QVector3D_isNull uintptr
	pg_QVector3D_x func(uintptr) float32
	pg_QVector3D_y func(uintptr) float32
	pg_QVector3D_z func(uintptr) float32
	pg_QVector3D_setX func(uintptr, float32)
	pg_QVector3D_setY func(uintptr, float32)
	pg_QVector3D_setZ func(uintptr, float32)
	pg_QVector3D_OperatorSubscript func(uintptr, int32) float32
	pg_QVector3D_length func(uintptr) float32
	pg_QVector3D_lengthSquared func(uintptr) float32
	addr_QVector3D_normalized uintptr
	addr_QVector3D_normalize uintptr
	addr_QVector3D_OperatorPlusAssign uintptr
	addr_QVector3D_OperatorMinusAssign uintptr
	pg_QVector3D_OperatorMultiplyAssign func(uintptr, float32) uintptr
	addr_QVector3D_OperatorMultiplyAssign2 uintptr
	pg_QVector3D_OperatorDivideAssign func(uintptr, float32) uintptr
	addr_QVector3D_OperatorDivideAssign2 uintptr
	pg_QVector3D_dotProduct func(uintptr, uintptr) float32
	addr_QVector3D_crossProduct uintptr
	addr_QVector3D_normal uintptr
	addr_QVector3D_normal2 uintptr
	addr_QVector3D_project uintptr
	addr_QVector3D_unproject uintptr
	pg_QVector3D_distanceToPoint func(uintptr, uintptr) float32
	pg_QVector3D_distanceToPlane func(uintptr, uintptr, uintptr) float32
	pg_QVector3D_distanceToPlane2 func(uintptr, uintptr, uintptr, uintptr) float32
	pg_QVector3D_distanceToLine func(uintptr, uintptr, uintptr) float32
	addr_QVector3D_toVector2D uintptr
	addr_QVector3D_toVector4D uintptr
	addr_QVector3D_toPoint uintptr
	addr_QVector3D_toPointF uintptr
	addr_QVector3D_ToQVariant uintptr
	addr_QVector3D_OperatorEqual uintptr
	addr_QVector3D_OperatorNotEqual uintptr
	addr_QVector3D_OperatorPlus uintptr
	addr_QVector3D_OperatorMinus uintptr
	pg_QVector3D_OperatorMultiply func(uintptr, float32) uintptr
	addr_QVector3D_OperatorMultiply2 uintptr
	addr_QVector3D_OperatorMinus2 uintptr
	pg_QVector3D_OperatorDivide func(uintptr, float32) uintptr
	addr_QVector3D_OperatorDivide2 uintptr
	addr_QVector3D_OperatorMultiply3 uintptr
	pg_QVector3D_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVector3D_new, "QVector3D_new")
	symaddr(&addr_QVector3D_new2, "QVector3D_new2")
	reg(&pg_QVector3D_new3, "QVector3D_new3")
	symaddr(&addr_QVector3D_new4, "QVector3D_new4")
	symaddr(&addr_QVector3D_new5, "QVector3D_new5")
	symaddr(&addr_QVector3D_new6, "QVector3D_new6")
	reg(&pg_QVector3D_new7, "QVector3D_new7")
	symaddr(&addr_QVector3D_new8, "QVector3D_new8")
	symaddr(&addr_QVector3D_new9, "QVector3D_new9")
	symaddr(&addr_QVector3D_isNull, "QVector3D_isNull")
	reg(&pg_QVector3D_x, "QVector3D_x")
	reg(&pg_QVector3D_y, "QVector3D_y")
	reg(&pg_QVector3D_z, "QVector3D_z")
	reg(&pg_QVector3D_setX, "QVector3D_setX")
	reg(&pg_QVector3D_setY, "QVector3D_setY")
	reg(&pg_QVector3D_setZ, "QVector3D_setZ")
	reg(&pg_QVector3D_OperatorSubscript, "QVector3D_OperatorSubscript")
	reg(&pg_QVector3D_length, "QVector3D_length")
	reg(&pg_QVector3D_lengthSquared, "QVector3D_lengthSquared")
	symaddr(&addr_QVector3D_normalized, "QVector3D_normalized")
	symaddr(&addr_QVector3D_normalize, "QVector3D_normalize")
	symaddr(&addr_QVector3D_OperatorPlusAssign, "QVector3D_OperatorPlusAssign")
	symaddr(&addr_QVector3D_OperatorMinusAssign, "QVector3D_OperatorMinusAssign")
	reg(&pg_QVector3D_OperatorMultiplyAssign, "QVector3D_OperatorMultiplyAssign")
	symaddr(&addr_QVector3D_OperatorMultiplyAssign2, "QVector3D_OperatorMultiplyAssign2")
	reg(&pg_QVector3D_OperatorDivideAssign, "QVector3D_OperatorDivideAssign")
	symaddr(&addr_QVector3D_OperatorDivideAssign2, "QVector3D_OperatorDivideAssign2")
	reg(&pg_QVector3D_dotProduct, "QVector3D_dotProduct")
	symaddr(&addr_QVector3D_crossProduct, "QVector3D_crossProduct")
	symaddr(&addr_QVector3D_normal, "QVector3D_normal")
	symaddr(&addr_QVector3D_normal2, "QVector3D_normal2")
	symaddr(&addr_QVector3D_project, "QVector3D_project")
	symaddr(&addr_QVector3D_unproject, "QVector3D_unproject")
	reg(&pg_QVector3D_distanceToPoint, "QVector3D_distanceToPoint")
	reg(&pg_QVector3D_distanceToPlane, "QVector3D_distanceToPlane")
	reg(&pg_QVector3D_distanceToPlane2, "QVector3D_distanceToPlane2")
	reg(&pg_QVector3D_distanceToLine, "QVector3D_distanceToLine")
	symaddr(&addr_QVector3D_toVector2D, "QVector3D_toVector2D")
	symaddr(&addr_QVector3D_toVector4D, "QVector3D_toVector4D")
	symaddr(&addr_QVector3D_toPoint, "QVector3D_toPoint")
	symaddr(&addr_QVector3D_toPointF, "QVector3D_toPointF")
	symaddr(&addr_QVector3D_ToQVariant, "QVector3D_ToQVariant")
	symaddr(&addr_QVector3D_OperatorEqual, "QVector3D_OperatorEqual")
	symaddr(&addr_QVector3D_OperatorNotEqual, "QVector3D_OperatorNotEqual")
	symaddr(&addr_QVector3D_OperatorPlus, "QVector3D_OperatorPlus")
	symaddr(&addr_QVector3D_OperatorMinus, "QVector3D_OperatorMinus")
	reg(&pg_QVector3D_OperatorMultiply, "QVector3D_OperatorMultiply")
	symaddr(&addr_QVector3D_OperatorMultiply2, "QVector3D_OperatorMultiply2")
	symaddr(&addr_QVector3D_OperatorMinus2, "QVector3D_OperatorMinus2")
	reg(&pg_QVector3D_OperatorDivide, "QVector3D_OperatorDivide")
	symaddr(&addr_QVector3D_OperatorDivide2, "QVector3D_OperatorDivide2")
	symaddr(&addr_QVector3D_OperatorMultiply3, "QVector3D_OperatorMultiply3")
	reg(&pg_QVector3D_delete, "QVector3D_delete")
}

type QVector3D struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVector3D) UnsafeGoID() int64     { return this.vid }
func (this *QVector3D) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVector3D(h uintptr) *QVector3D {
	if h == 0 {
		return nil
	}
	return &QVector3D{h: h}
}

func newownedQVector3D(h uintptr) *QVector3D {
	o := bindQVector3D(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVector3D).Delete)
	}
	return o
}

func (this *QVector3D) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVector3D) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVector3D(h unsafe.Pointer) *QVector3D { return bindQVector3D(uintptr(h)) }
func UnsafeNewQVector3DOwned(h unsafe.Pointer) *QVector3D { return newownedQVector3D(uintptr(h)) }

func (this *QVector3D) UnsafeRelease() { this.h = 0 }

func NewQVector3D() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new)
	__o := newownedQVector3D(__r)
	return __o
}

func NewQVector3D2(arg1 Initialization) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new2, uintptr(arg1))
	__o := newownedQVector3D(__r)
	return __o
}

func NewQVector3D3(xpos float32, ypos float32, zpos float32) *QVector3D {
	__o := newownedQVector3D(pg_QVector3D_new3(xpos, ypos, zpos))
	return __o
}

func NewQVector3D4(point *QPoint) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new4, uintptr(point.UnsafePointer()))
	__o := newownedQVector3D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector3D5(point *QPointF) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new5, uintptr(point.UnsafePointer()))
	__o := newownedQVector3D(__r)
	runtime.KeepAlive(point)
	return __o
}

func NewQVector3D6(vector *QVector2D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new6, uintptr(vector.UnsafePointer()))
	__o := newownedQVector3D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector3D7(vector *QVector2D, zpos float32) *QVector3D {
	__o := newownedQVector3D(pg_QVector3D_new7(uintptr(vector.UnsafePointer()), zpos))
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector3D8(vector *QVector4D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new8, uintptr(vector.UnsafePointer()))
	__o := newownedQVector3D(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQVector3D9(arg1 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_new9, uintptr(arg1.UnsafePointer()))
	__o := newownedQVector3D(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QVector3D) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QVector3D_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVector3D) X() float32 {
	__r := pg_QVector3D_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) Y() float32 {
	__r := pg_QVector3D_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) Z() float32 {
	__r := pg_QVector3D_z(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) SetX(x float32) {
	pg_QVector3D_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QVector3D) SetY(y float32) {
	pg_QVector3D_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QVector3D) SetZ(z float32) {
	pg_QVector3D_setZ(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QVector3D) OperatorSubscript(i int) float32 {
	__r := pg_QVector3D_OperatorSubscript(this.h, int32(i))
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) Length() float32 {
	__r := pg_QVector3D_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) LengthSquared() float32 {
	__r := pg_QVector3D_lengthSquared(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVector3D) Normalized() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector3D) Normalize() {
	purego.SyscallN(addr_QVector3D_normalize, this.h)
	runtime.KeepAlive(this)
}

func (this *QVector3D) OperatorPlusAssign(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorPlusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector3D(__r)
}

func (this *QVector3D) OperatorMinusAssign(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMinusAssign, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector3D(__r)
}

func (this *QVector3D) OperatorMultiplyAssign(factor float32) *QVector3D {
	__r := pg_QVector3D_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQVector3D(__r)
}

func (this *QVector3D) OperatorMultiplyAssignWithVector(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMultiplyAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector3D(__r)
}

func (this *QVector3D) OperatorDivideAssign(divisor float32) *QVector3D {
	__r := pg_QVector3D_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQVector3D(__r)
}

func (this *QVector3D) OperatorDivideAssignWithVector(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorDivideAssign2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return bindQVector3D(__r)
}

func QVector3D_DotProduct(v1 *QVector3D, v2 *QVector3D) float32 {
	__r := pg_QVector3D_dotProduct(uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return __r
}

func QVector3D_CrossProduct(v1 *QVector3D, v2 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_crossProduct, uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return newownedQVector3D(__r)
}

func QVector3D_Normal(v1 *QVector3D, v2 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_normal, uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return newownedQVector3D(__r)
}

func QVector3D_Normal2(v1 *QVector3D, v2 *QVector3D, v3 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_normal2, uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()), uintptr(v3.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	runtime.KeepAlive(v3)
	return newownedQVector3D(__r)
}

func (this *QVector3D) Project(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_project, this.h, uintptr(modelView.UnsafePointer()), uintptr(projection.UnsafePointer()), uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modelView)
	runtime.KeepAlive(projection)
	runtime.KeepAlive(viewport)
	return newownedQVector3D(__r)
}

func (this *QVector3D) Unproject(modelView *QMatrix4x4, projection *QMatrix4x4, viewport *QRect) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_unproject, this.h, uintptr(modelView.UnsafePointer()), uintptr(projection.UnsafePointer()), uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modelView)
	runtime.KeepAlive(projection)
	runtime.KeepAlive(viewport)
	return newownedQVector3D(__r)
}

func (this *QVector3D) DistanceToPoint(point *QVector3D) float32 {
	__r := pg_QVector3D_distanceToPoint(this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r
}

func (this *QVector3D) DistanceToPlane(plane *QVector3D, normal *QVector3D) float32 {
	__r := pg_QVector3D_distanceToPlane(this.h, uintptr(plane.UnsafePointer()), uintptr(normal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(plane)
	runtime.KeepAlive(normal)
	return __r
}

func (this *QVector3D) DistanceToPlane2(plane1 *QVector3D, plane2 *QVector3D, plane3 *QVector3D) float32 {
	__r := pg_QVector3D_distanceToPlane2(this.h, uintptr(plane1.UnsafePointer()), uintptr(plane2.UnsafePointer()), uintptr(plane3.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(plane1)
	runtime.KeepAlive(plane2)
	runtime.KeepAlive(plane3)
	return __r
}

func (this *QVector3D) DistanceToLine(point *QVector3D, direction *QVector3D) float32 {
	__r := pg_QVector3D_distanceToLine(this.h, uintptr(point.UnsafePointer()), uintptr(direction.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	runtime.KeepAlive(direction)
	return __r
}

func (this *QVector3D) ToVector2D() *QVector2D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_toVector2D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector2D(__r)
}

func (this *QVector3D) ToVector4D() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_toVector4D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QVector3D) ToPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QVector3D_toPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QVector3D) ToPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QVector3D_toPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QVector3D) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVector3D_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QVector3D) OperatorEqual(v2 *QVector3D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector3D_OperatorEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector3D) OperatorNotEqual(v2 *QVector3D) bool {
	__r0, _, _ := purego.SyscallN(addr_QVector3D_OperatorNotEqual, this.h, uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return __r != 0
}

func (this *QVector3D) OperatorPlus(v2 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorPlus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorMinus(v2 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMinus, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorMultiply(factor float32) *QVector3D {
	__r := pg_QVector3D_OperatorMultiply(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorMultiplyWithV2(v2 *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMultiply2, this.h, uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(v2)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorMinus2() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorDivide(divisor float32) *QVector3D {
	__r := pg_QVector3D_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorDivideWithDivisor(divisor *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorDivide2, this.h, uintptr(divisor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(divisor)
	return newownedQVector3D(__r)
}

func (this *QVector3D) OperatorMultiplyWithMatrix(matrix *QMatrix4x4) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QVector3D_OperatorMultiply3, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQVector3D(__r)
}

func (this *QVector3D) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVector3D_delete(this.h)
	}
	this.UnsafeRelease()
}
