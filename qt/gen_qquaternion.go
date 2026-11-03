package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QQuaternion_new uintptr
	addr_QQuaternion_new2 uintptr
	pg_QQuaternion_new3 func(float32, float32, float32, float32) uintptr
	pg_QQuaternion_new4 func(float32, uintptr) uintptr
	addr_QQuaternion_new5 uintptr
	addr_QQuaternion_new6 uintptr
	addr_QQuaternion_isNull uintptr
	addr_QQuaternion_isIdentity uintptr
	addr_QQuaternion_vector uintptr
	addr_QQuaternion_setVector uintptr
	pg_QQuaternion_setVector2 func(uintptr, float32, float32, float32)
	pg_QQuaternion_x func(uintptr) float32
	pg_QQuaternion_y func(uintptr) float32
	pg_QQuaternion_z func(uintptr) float32
	pg_QQuaternion_scalar func(uintptr) float32
	pg_QQuaternion_setX func(uintptr, float32)
	pg_QQuaternion_setY func(uintptr, float32)
	pg_QQuaternion_setZ func(uintptr, float32)
	pg_QQuaternion_setScalar func(uintptr, float32)
	pg_QQuaternion_dotProduct func(uintptr, uintptr) float32
	pg_QQuaternion_length func(uintptr) float32
	pg_QQuaternion_lengthSquared func(uintptr) float32
	addr_QQuaternion_normalized uintptr
	addr_QQuaternion_normalize uintptr
	addr_QQuaternion_inverted uintptr
	addr_QQuaternion_conjugated uintptr
	addr_QQuaternion_rotatedVector uintptr
	addr_QQuaternion_OperatorPlusAssign uintptr
	addr_QQuaternion_OperatorMinusAssign uintptr
	pg_QQuaternion_OperatorMultiplyAssign func(uintptr, float32) uintptr
	addr_QQuaternion_OperatorMultiplyAssign2 uintptr
	pg_QQuaternion_OperatorDivideAssign func(uintptr, float32) uintptr
	addr_QQuaternion_toVector4D uintptr
	addr_QQuaternion_ToQVariant uintptr
	addr_QQuaternion_getAxisAndAngle uintptr
	pg_QQuaternion_fromAxisAndAngle func(uintptr, float32) uintptr
	addr_QQuaternion_getAxisAndAngle2 uintptr
	pg_QQuaternion_fromAxisAndAngle2 func(float32, float32, float32, float32) uintptr
	addr_QQuaternion_toEulerAngles uintptr
	addr_QQuaternion_getEulerAngles uintptr
	pg_QQuaternion_fromEulerAngles func(float32, float32, float32) uintptr
	addr_QQuaternion_toAxes uintptr
	addr_QQuaternion_fromAxes uintptr
	addr_QQuaternion_getAxes uintptr
	addr_QQuaternion_fromAxes2 uintptr
	addr_QQuaternion_fromDirection uintptr
	addr_QQuaternion_rotationTo uintptr
	pg_QQuaternion_slerp func(uintptr, uintptr, float32) uintptr
	pg_QQuaternion_nlerp func(uintptr, uintptr, float32) uintptr
	addr_QQuaternion_OperatorAssign uintptr
	addr_QQuaternion_OperatorEqual uintptr
	addr_QQuaternion_OperatorNotEqual uintptr
	addr_QQuaternion_OperatorPlus uintptr
	addr_QQuaternion_OperatorMinus uintptr
	pg_QQuaternion_OperatorMultiply func(uintptr, float32) uintptr
	addr_QQuaternion_OperatorMultiply2 uintptr
	addr_QQuaternion_OperatorMinus2 uintptr
	pg_QQuaternion_OperatorDivide func(uintptr, float32) uintptr
	addr_QQuaternion_OperatorMultiply3 uintptr
	pg_QQuaternion_delete func(uintptr)
)

func init() {
	symaddr(&addr_QQuaternion_new, "QQuaternion_new")
	symaddr(&addr_QQuaternion_new2, "QQuaternion_new2")
	reg(&pg_QQuaternion_new3, "QQuaternion_new3")
	reg(&pg_QQuaternion_new4, "QQuaternion_new4")
	symaddr(&addr_QQuaternion_new5, "QQuaternion_new5")
	symaddr(&addr_QQuaternion_new6, "QQuaternion_new6")
	symaddr(&addr_QQuaternion_isNull, "QQuaternion_isNull")
	symaddr(&addr_QQuaternion_isIdentity, "QQuaternion_isIdentity")
	symaddr(&addr_QQuaternion_vector, "QQuaternion_vector")
	symaddr(&addr_QQuaternion_setVector, "QQuaternion_setVector")
	reg(&pg_QQuaternion_setVector2, "QQuaternion_setVector2")
	reg(&pg_QQuaternion_x, "QQuaternion_x")
	reg(&pg_QQuaternion_y, "QQuaternion_y")
	reg(&pg_QQuaternion_z, "QQuaternion_z")
	reg(&pg_QQuaternion_scalar, "QQuaternion_scalar")
	reg(&pg_QQuaternion_setX, "QQuaternion_setX")
	reg(&pg_QQuaternion_setY, "QQuaternion_setY")
	reg(&pg_QQuaternion_setZ, "QQuaternion_setZ")
	reg(&pg_QQuaternion_setScalar, "QQuaternion_setScalar")
	reg(&pg_QQuaternion_dotProduct, "QQuaternion_dotProduct")
	reg(&pg_QQuaternion_length, "QQuaternion_length")
	reg(&pg_QQuaternion_lengthSquared, "QQuaternion_lengthSquared")
	symaddr(&addr_QQuaternion_normalized, "QQuaternion_normalized")
	symaddr(&addr_QQuaternion_normalize, "QQuaternion_normalize")
	symaddr(&addr_QQuaternion_inverted, "QQuaternion_inverted")
	symaddr(&addr_QQuaternion_conjugated, "QQuaternion_conjugated")
	symaddr(&addr_QQuaternion_rotatedVector, "QQuaternion_rotatedVector")
	symaddr(&addr_QQuaternion_OperatorPlusAssign, "QQuaternion_OperatorPlusAssign")
	symaddr(&addr_QQuaternion_OperatorMinusAssign, "QQuaternion_OperatorMinusAssign")
	reg(&pg_QQuaternion_OperatorMultiplyAssign, "QQuaternion_OperatorMultiplyAssign")
	symaddr(&addr_QQuaternion_OperatorMultiplyAssign2, "QQuaternion_OperatorMultiplyAssign2")
	reg(&pg_QQuaternion_OperatorDivideAssign, "QQuaternion_OperatorDivideAssign")
	symaddr(&addr_QQuaternion_toVector4D, "QQuaternion_toVector4D")
	symaddr(&addr_QQuaternion_ToQVariant, "QQuaternion_ToQVariant")
	symaddr(&addr_QQuaternion_getAxisAndAngle, "QQuaternion_getAxisAndAngle")
	reg(&pg_QQuaternion_fromAxisAndAngle, "QQuaternion_fromAxisAndAngle")
	symaddr(&addr_QQuaternion_getAxisAndAngle2, "QQuaternion_getAxisAndAngle2")
	reg(&pg_QQuaternion_fromAxisAndAngle2, "QQuaternion_fromAxisAndAngle2")
	symaddr(&addr_QQuaternion_toEulerAngles, "QQuaternion_toEulerAngles")
	symaddr(&addr_QQuaternion_getEulerAngles, "QQuaternion_getEulerAngles")
	reg(&pg_QQuaternion_fromEulerAngles, "QQuaternion_fromEulerAngles")
	symaddr(&addr_QQuaternion_toAxes, "QQuaternion_toAxes")
	symaddr(&addr_QQuaternion_fromAxes, "QQuaternion_fromAxes")
	symaddr(&addr_QQuaternion_getAxes, "QQuaternion_getAxes")
	symaddr(&addr_QQuaternion_fromAxes2, "QQuaternion_fromAxes2")
	symaddr(&addr_QQuaternion_fromDirection, "QQuaternion_fromDirection")
	symaddr(&addr_QQuaternion_rotationTo, "QQuaternion_rotationTo")
	reg(&pg_QQuaternion_slerp, "QQuaternion_slerp")
	reg(&pg_QQuaternion_nlerp, "QQuaternion_nlerp")
	symaddr(&addr_QQuaternion_OperatorAssign, "QQuaternion_OperatorAssign")
	symaddr(&addr_QQuaternion_OperatorEqual, "QQuaternion_OperatorEqual")
	symaddr(&addr_QQuaternion_OperatorNotEqual, "QQuaternion_OperatorNotEqual")
	symaddr(&addr_QQuaternion_OperatorPlus, "QQuaternion_OperatorPlus")
	symaddr(&addr_QQuaternion_OperatorMinus, "QQuaternion_OperatorMinus")
	reg(&pg_QQuaternion_OperatorMultiply, "QQuaternion_OperatorMultiply")
	symaddr(&addr_QQuaternion_OperatorMultiply2, "QQuaternion_OperatorMultiply2")
	symaddr(&addr_QQuaternion_OperatorMinus2, "QQuaternion_OperatorMinus2")
	reg(&pg_QQuaternion_OperatorDivide, "QQuaternion_OperatorDivide")
	symaddr(&addr_QQuaternion_OperatorMultiply3, "QQuaternion_OperatorMultiply3")
	reg(&pg_QQuaternion_delete, "QQuaternion_delete")
}

type QQuaternion struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QQuaternion) UnsafeGoID() int64     { return this.vid }
func (this *QQuaternion) UnsafeSetGoID(id int64) { this.vid = id }

func bindQQuaternion(h uintptr) *QQuaternion {
	if h == 0 {
		return nil
	}
	return &QQuaternion{h: h}
}

func newownedQQuaternion(h uintptr) *QQuaternion {
	o := bindQQuaternion(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QQuaternion).Delete)
	}
	return o
}

func (this *QQuaternion) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QQuaternion) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQQuaternion(h unsafe.Pointer) *QQuaternion { return bindQQuaternion(uintptr(h)) }
func UnsafeNewQQuaternionOwned(h unsafe.Pointer) *QQuaternion { return newownedQQuaternion(uintptr(h)) }

func (this *QQuaternion) UnsafeRelease() { this.h = 0 }

func NewQQuaternion() *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_new)
	__o := newownedQQuaternion(__r)
	return __o
}

func NewQQuaternion2(arg1 Initialization) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_new2, uintptr(arg1))
	__o := newownedQQuaternion(__r)
	return __o
}

func NewQQuaternion3(scalar float32, xpos float32, ypos float32, zpos float32) *QQuaternion {
	__o := newownedQQuaternion(pg_QQuaternion_new3(scalar, xpos, ypos, zpos))
	return __o
}

func NewQQuaternion4(scalar float32, vector *QVector3D) *QQuaternion {
	__o := newownedQQuaternion(pg_QQuaternion_new4(scalar, uintptr(vector.UnsafePointer())))
	runtime.KeepAlive(vector)
	return __o
}

func NewQQuaternion5(vector *QVector4D) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_new5, uintptr(vector.UnsafePointer()))
	__o := newownedQQuaternion(__r)
	runtime.KeepAlive(vector)
	return __o
}

func NewQQuaternion6(arg1 *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_new6, uintptr(arg1.UnsafePointer()))
	__o := newownedQQuaternion(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QQuaternion) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QQuaternion_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QQuaternion) IsIdentity() bool {
	__r0, _, _ := purego.SyscallN(addr_QQuaternion_isIdentity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QQuaternion) Vector() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_vector, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QQuaternion) SetVector(vector *QVector3D) {
	purego.SyscallN(addr_QQuaternion_setVector, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
}

func (this *QQuaternion) SetVector2(x float32, y float32, z float32) {
	pg_QQuaternion_setVector2(this.h, x, y, z)
	runtime.KeepAlive(this)
}

func (this *QQuaternion) X() float32 {
	__r := pg_QQuaternion_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) Y() float32 {
	__r := pg_QQuaternion_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) Z() float32 {
	__r := pg_QQuaternion_z(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) Scalar() float32 {
	__r := pg_QQuaternion_scalar(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) SetX(x float32) {
	pg_QQuaternion_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QQuaternion) SetY(y float32) {
	pg_QQuaternion_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QQuaternion) SetZ(z float32) {
	pg_QQuaternion_setZ(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QQuaternion) SetScalar(scalar float32) {
	pg_QQuaternion_setScalar(this.h, scalar)
	runtime.KeepAlive(this)
}

func QQuaternion_DotProduct(q1 *QQuaternion, q2 *QQuaternion) float32 {
	__r := pg_QQuaternion_dotProduct(uintptr(q1.UnsafePointer()), uintptr(q2.UnsafePointer()))
	runtime.KeepAlive(q1)
	runtime.KeepAlive(q2)
	return __r
}

func (this *QQuaternion) Length() float32 {
	__r := pg_QQuaternion_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) LengthSquared() float32 {
	__r := pg_QQuaternion_lengthSquared(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion) Normalized() *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) Normalize() {
	purego.SyscallN(addr_QQuaternion_normalize, this.h)
	runtime.KeepAlive(this)
}

func (this *QQuaternion) Inverted() *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_inverted, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) Conjugated() *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_conjugated, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) RotatedVector(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_rotatedVector, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQVector3D(__r)
}

func (this *QQuaternion) OperatorPlusAssign(quaternion *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorPlusAssign, this.h, uintptr(quaternion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(quaternion)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) OperatorMinusAssign(quaternion *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMinusAssign, this.h, uintptr(quaternion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(quaternion)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) OperatorMultiplyAssign(factor float32) *QQuaternion {
	__r := pg_QQuaternion_OperatorMultiplyAssign(this.h, factor)
	runtime.KeepAlive(this)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) OperatorMultiplyAssignWithQuaternion(quaternion *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMultiplyAssign2, this.h, uintptr(quaternion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(quaternion)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) OperatorDivideAssign(divisor float32) *QQuaternion {
	__r := pg_QQuaternion_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) ToVector4D() *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_toVector4D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QQuaternion) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QQuaternion) GetAxisAndAngle(axis *QVector3D) float32 {
	var angle_out float32
	purego.SyscallN(addr_QQuaternion_getAxisAndAngle, this.h, uintptr(axis.UnsafePointer()), uintptr(unsafe.Pointer(&angle_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(axis)
	return angle_out
}

func QQuaternion_FromAxisAndAngle(axis *QVector3D, angle float32) *QQuaternion {
	__r := pg_QQuaternion_fromAxisAndAngle(uintptr(axis.UnsafePointer()), angle)
	runtime.KeepAlive(axis)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) GetAxisAndAngle2() (float32, float32, float32, float32) {
	var x_out float32
	var y_out float32
	var z_out float32
	var angle_out float32
	purego.SyscallN(addr_QQuaternion_getAxisAndAngle2, this.h, uintptr(unsafe.Pointer(&x_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&z_out)), uintptr(unsafe.Pointer(&angle_out)))
	runtime.KeepAlive(this)
	return x_out, y_out, z_out, angle_out
}

func QQuaternion_FromAxisAndAngle2(x float32, y float32, z float32, angle float32) *QQuaternion {
	__r := pg_QQuaternion_fromAxisAndAngle2(x, y, z, angle)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) ToEulerAngles() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_toEulerAngles, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QQuaternion) GetEulerAngles() (float32, float32, float32) {
	var pitch_out float32
	var yaw_out float32
	var roll_out float32
	purego.SyscallN(addr_QQuaternion_getEulerAngles, this.h, uintptr(unsafe.Pointer(&pitch_out)), uintptr(unsafe.Pointer(&yaw_out)), uintptr(unsafe.Pointer(&roll_out)))
	runtime.KeepAlive(this)
	return pitch_out, yaw_out, roll_out
}

func QQuaternion_FromEulerAngles(pitch float32, yaw float32, roll float32) *QQuaternion {
	__r := pg_QQuaternion_fromEulerAngles(pitch, yaw, roll)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) ToAxes() *QQuaternion__Axes {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_toAxes, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion__Axes(__r)
}

func QQuaternion_FromAxes(axes *QQuaternion__Axes) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_fromAxes, uintptr(axes.UnsafePointer()))
	runtime.KeepAlive(axes)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) GetAxes(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) {
	purego.SyscallN(addr_QQuaternion_getAxes, this.h, uintptr(xAxis.UnsafePointer()), uintptr(yAxis.UnsafePointer()), uintptr(zAxis.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(xAxis)
	runtime.KeepAlive(yAxis)
	runtime.KeepAlive(zAxis)
}

func QQuaternion_FromAxes2(xAxis *QVector3D, yAxis *QVector3D, zAxis *QVector3D) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_fromAxes2, uintptr(xAxis.UnsafePointer()), uintptr(yAxis.UnsafePointer()), uintptr(zAxis.UnsafePointer()))
	runtime.KeepAlive(xAxis)
	runtime.KeepAlive(yAxis)
	runtime.KeepAlive(zAxis)
	return newownedQQuaternion(__r)
}

func QQuaternion_FromDirection(direction *QVector3D, up *QVector3D) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_fromDirection, uintptr(direction.UnsafePointer()), uintptr(up.UnsafePointer()))
	runtime.KeepAlive(direction)
	runtime.KeepAlive(up)
	return newownedQQuaternion(__r)
}

func QQuaternion_RotationTo(from *QVector3D, to *QVector3D) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_rotationTo, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return newownedQQuaternion(__r)
}

func QQuaternion_Slerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion {
	__r := pg_QQuaternion_slerp(uintptr(q1.UnsafePointer()), uintptr(q2.UnsafePointer()), t)
	runtime.KeepAlive(q1)
	runtime.KeepAlive(q2)
	return newownedQQuaternion(__r)
}

func QQuaternion_Nlerp(q1 *QQuaternion, q2 *QQuaternion, t float32) *QQuaternion {
	__r := pg_QQuaternion_nlerp(uintptr(q1.UnsafePointer()), uintptr(q2.UnsafePointer()), t)
	runtime.KeepAlive(q1)
	runtime.KeepAlive(q2)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorAssign(arg1 *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQQuaternion(__r)
}

func (this *QQuaternion) OperatorEqual(q2 *QQuaternion) bool {
	__r0, _, _ := purego.SyscallN(addr_QQuaternion_OperatorEqual, this.h, uintptr(q2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q2)
	return __r != 0
}

func (this *QQuaternion) OperatorNotEqual(q2 *QQuaternion) bool {
	__r0, _, _ := purego.SyscallN(addr_QQuaternion_OperatorNotEqual, this.h, uintptr(q2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q2)
	return __r != 0
}

func (this *QQuaternion) OperatorPlus(q2 *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorPlus, this.h, uintptr(q2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q2)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorMinus(q2 *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMinus, this.h, uintptr(q2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q2)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorMultiply(factor float32) *QQuaternion {
	__r := pg_QQuaternion_OperatorMultiply(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorMultiplyWithQ2(q2 *QQuaternion) *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMultiply2, this.h, uintptr(q2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q2)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorMinus2() *QQuaternion {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorDivide(divisor float32) *QQuaternion {
	__r := pg_QQuaternion_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQQuaternion(__r)
}

func (this *QQuaternion) OperatorMultiplyWithVec(vec *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion_OperatorMultiply3, this.h, uintptr(vec.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vec)
	return newownedQVector3D(__r)
}

func (this *QQuaternion) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QQuaternion_delete(this.h)
	}
	this.UnsafeRelease()
}
