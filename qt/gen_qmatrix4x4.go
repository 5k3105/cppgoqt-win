package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): NewQMatrix4x43

var (
	addr_QMatrix4x4_new uintptr
	addr_QMatrix4x4_new2 uintptr
	addr_QMatrix4x4_new4 uintptr
	addr_QMatrix4x4_new5 uintptr
	addr_QMatrix4x4_column uintptr
	addr_QMatrix4x4_setColumn uintptr
	addr_QMatrix4x4_row uintptr
	addr_QMatrix4x4_setRow uintptr
	addr_QMatrix4x4_isAffine uintptr
	addr_QMatrix4x4_isIdentity uintptr
	addr_QMatrix4x4_setToIdentity uintptr
	pg_QMatrix4x4_fill func(uintptr, float32)
	pg_QMatrix4x4_determinant func(uintptr) float64
	addr_QMatrix4x4_inverted uintptr
	addr_QMatrix4x4_transposed uintptr
	addr_QMatrix4x4_OperatorPlusAssign uintptr
	addr_QMatrix4x4_OperatorMinusAssign uintptr
	addr_QMatrix4x4_OperatorMultiplyAssign uintptr
	pg_QMatrix4x4_OperatorMultiplyAssign2 func(uintptr, float32) uintptr
	pg_QMatrix4x4_OperatorDivideAssign func(uintptr, float32) uintptr
	addr_QMatrix4x4_OperatorEqual uintptr
	addr_QMatrix4x4_OperatorNotEqual uintptr
	addr_QMatrix4x4_scale uintptr
	addr_QMatrix4x4_translate uintptr
	pg_QMatrix4x4_rotate func(uintptr, float32, uintptr)
	pg_QMatrix4x4_scale2 func(uintptr, float32, float32)
	pg_QMatrix4x4_scale3 func(uintptr, float32, float32, float32)
	pg_QMatrix4x4_scale4 func(uintptr, float32)
	pg_QMatrix4x4_translate2 func(uintptr, float32, float32)
	pg_QMatrix4x4_translate3 func(uintptr, float32, float32, float32)
	pg_QMatrix4x4_rotate2 func(uintptr, float32, float32, float32)
	addr_QMatrix4x4_rotate3 uintptr
	addr_QMatrix4x4_ortho uintptr
	addr_QMatrix4x4_ortho2 uintptr
	pg_QMatrix4x4_ortho3 func(uintptr, float32, float32, float32, float32, float32, float32)
	pg_QMatrix4x4_frustum func(uintptr, float32, float32, float32, float32, float32, float32)
	pg_QMatrix4x4_perspective func(uintptr, float32, float32, float32, float32)
	addr_QMatrix4x4_lookAt uintptr
	addr_QMatrix4x4_viewport uintptr
	pg_QMatrix4x4_viewport2 func(uintptr, float32, float32, float32, float32)
	addr_QMatrix4x4_flipCoordinates uintptr
	addr_QMatrix4x4_copyDataTo uintptr
	addr_QMatrix4x4_toTransform uintptr
	pg_QMatrix4x4_toTransform2 func(uintptr, float32) uintptr
	addr_QMatrix4x4_map uintptr
	addr_QMatrix4x4_map2 uintptr
	addr_QMatrix4x4_map3 uintptr
	addr_QMatrix4x4_mapVector uintptr
	addr_QMatrix4x4_map4 uintptr
	addr_QMatrix4x4_mapRect uintptr
	addr_QMatrix4x4_mapRect2 uintptr
	addr_QMatrix4x4_optimize uintptr
	addr_QMatrix4x4_ToQVariant uintptr
	pg_QMatrix4x4_projectedRotate func(uintptr, float32, float32, float32, float32, float32)
	pg_QMatrix4x4_projectedRotate2 func(uintptr, float32, float32, float32, float32)
	addr_QMatrix4x4_flags uintptr
	addr_QMatrix4x4_inverted2 uintptr
	pg_QMatrix4x4_rotate4 func(uintptr, float32, float32, float32, float32)
	pg_QMatrix4x4_viewport3 func(uintptr, float32, float32, float32, float32, float32)
	pg_QMatrix4x4_viewport4 func(uintptr, float32, float32, float32, float32, float32, float32)
	addr_QMatrix4x4_OperatorMultiply uintptr
	addr_QMatrix4x4_OperatorMultiply2 uintptr
	addr_QMatrix4x4_OperatorPlus uintptr
	addr_QMatrix4x4_OperatorMinus uintptr
	addr_QMatrix4x4_OperatorMultiply3 uintptr
	addr_QMatrix4x4_OperatorMinus2 uintptr
	addr_QMatrix4x4_OperatorMultiply4 uintptr
	addr_QMatrix4x4_OperatorMultiply5 uintptr
	pg_QMatrix4x4_OperatorMultiply6 func(uintptr, float32) uintptr
	pg_QMatrix4x4_OperatorDivide func(uintptr, float32) uintptr
	pg_QMatrix4x4_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMatrix4x4_new, "QMatrix4x4_new")
	symaddr(&addr_QMatrix4x4_new2, "QMatrix4x4_new2")
	symaddr(&addr_QMatrix4x4_new4, "QMatrix4x4_new4")
	symaddr(&addr_QMatrix4x4_new5, "QMatrix4x4_new5")
	symaddr(&addr_QMatrix4x4_column, "QMatrix4x4_column")
	symaddr(&addr_QMatrix4x4_setColumn, "QMatrix4x4_setColumn")
	symaddr(&addr_QMatrix4x4_row, "QMatrix4x4_row")
	symaddr(&addr_QMatrix4x4_setRow, "QMatrix4x4_setRow")
	symaddr(&addr_QMatrix4x4_isAffine, "QMatrix4x4_isAffine")
	symaddr(&addr_QMatrix4x4_isIdentity, "QMatrix4x4_isIdentity")
	symaddr(&addr_QMatrix4x4_setToIdentity, "QMatrix4x4_setToIdentity")
	reg(&pg_QMatrix4x4_fill, "QMatrix4x4_fill")
	reg(&pg_QMatrix4x4_determinant, "QMatrix4x4_determinant")
	symaddr(&addr_QMatrix4x4_inverted, "QMatrix4x4_inverted")
	symaddr(&addr_QMatrix4x4_transposed, "QMatrix4x4_transposed")
	symaddr(&addr_QMatrix4x4_OperatorPlusAssign, "QMatrix4x4_OperatorPlusAssign")
	symaddr(&addr_QMatrix4x4_OperatorMinusAssign, "QMatrix4x4_OperatorMinusAssign")
	symaddr(&addr_QMatrix4x4_OperatorMultiplyAssign, "QMatrix4x4_OperatorMultiplyAssign")
	reg(&pg_QMatrix4x4_OperatorMultiplyAssign2, "QMatrix4x4_OperatorMultiplyAssign2")
	reg(&pg_QMatrix4x4_OperatorDivideAssign, "QMatrix4x4_OperatorDivideAssign")
	symaddr(&addr_QMatrix4x4_OperatorEqual, "QMatrix4x4_OperatorEqual")
	symaddr(&addr_QMatrix4x4_OperatorNotEqual, "QMatrix4x4_OperatorNotEqual")
	symaddr(&addr_QMatrix4x4_scale, "QMatrix4x4_scale")
	symaddr(&addr_QMatrix4x4_translate, "QMatrix4x4_translate")
	reg(&pg_QMatrix4x4_rotate, "QMatrix4x4_rotate")
	reg(&pg_QMatrix4x4_scale2, "QMatrix4x4_scale2")
	reg(&pg_QMatrix4x4_scale3, "QMatrix4x4_scale3")
	reg(&pg_QMatrix4x4_scale4, "QMatrix4x4_scale4")
	reg(&pg_QMatrix4x4_translate2, "QMatrix4x4_translate2")
	reg(&pg_QMatrix4x4_translate3, "QMatrix4x4_translate3")
	reg(&pg_QMatrix4x4_rotate2, "QMatrix4x4_rotate2")
	symaddr(&addr_QMatrix4x4_rotate3, "QMatrix4x4_rotate3")
	symaddr(&addr_QMatrix4x4_ortho, "QMatrix4x4_ortho")
	symaddr(&addr_QMatrix4x4_ortho2, "QMatrix4x4_ortho2")
	reg(&pg_QMatrix4x4_ortho3, "QMatrix4x4_ortho3")
	reg(&pg_QMatrix4x4_frustum, "QMatrix4x4_frustum")
	reg(&pg_QMatrix4x4_perspective, "QMatrix4x4_perspective")
	symaddr(&addr_QMatrix4x4_lookAt, "QMatrix4x4_lookAt")
	symaddr(&addr_QMatrix4x4_viewport, "QMatrix4x4_viewport")
	reg(&pg_QMatrix4x4_viewport2, "QMatrix4x4_viewport2")
	symaddr(&addr_QMatrix4x4_flipCoordinates, "QMatrix4x4_flipCoordinates")
	symaddr(&addr_QMatrix4x4_copyDataTo, "QMatrix4x4_copyDataTo")
	symaddr(&addr_QMatrix4x4_toTransform, "QMatrix4x4_toTransform")
	reg(&pg_QMatrix4x4_toTransform2, "QMatrix4x4_toTransform2")
	symaddr(&addr_QMatrix4x4_map, "QMatrix4x4_map")
	symaddr(&addr_QMatrix4x4_map2, "QMatrix4x4_map2")
	symaddr(&addr_QMatrix4x4_map3, "QMatrix4x4_map3")
	symaddr(&addr_QMatrix4x4_mapVector, "QMatrix4x4_mapVector")
	symaddr(&addr_QMatrix4x4_map4, "QMatrix4x4_map4")
	symaddr(&addr_QMatrix4x4_mapRect, "QMatrix4x4_mapRect")
	symaddr(&addr_QMatrix4x4_mapRect2, "QMatrix4x4_mapRect2")
	symaddr(&addr_QMatrix4x4_optimize, "QMatrix4x4_optimize")
	symaddr(&addr_QMatrix4x4_ToQVariant, "QMatrix4x4_ToQVariant")
	reg(&pg_QMatrix4x4_projectedRotate, "QMatrix4x4_projectedRotate")
	reg(&pg_QMatrix4x4_projectedRotate2, "QMatrix4x4_projectedRotate2")
	symaddr(&addr_QMatrix4x4_flags, "QMatrix4x4_flags")
	symaddr(&addr_QMatrix4x4_inverted2, "QMatrix4x4_inverted2")
	reg(&pg_QMatrix4x4_rotate4, "QMatrix4x4_rotate4")
	reg(&pg_QMatrix4x4_viewport3, "QMatrix4x4_viewport3")
	reg(&pg_QMatrix4x4_viewport4, "QMatrix4x4_viewport4")
	symaddr(&addr_QMatrix4x4_OperatorMultiply, "QMatrix4x4_OperatorMultiply")
	symaddr(&addr_QMatrix4x4_OperatorMultiply2, "QMatrix4x4_OperatorMultiply2")
	symaddr(&addr_QMatrix4x4_OperatorPlus, "QMatrix4x4_OperatorPlus")
	symaddr(&addr_QMatrix4x4_OperatorMinus, "QMatrix4x4_OperatorMinus")
	symaddr(&addr_QMatrix4x4_OperatorMultiply3, "QMatrix4x4_OperatorMultiply3")
	symaddr(&addr_QMatrix4x4_OperatorMinus2, "QMatrix4x4_OperatorMinus2")
	symaddr(&addr_QMatrix4x4_OperatorMultiply4, "QMatrix4x4_OperatorMultiply4")
	symaddr(&addr_QMatrix4x4_OperatorMultiply5, "QMatrix4x4_OperatorMultiply5")
	reg(&pg_QMatrix4x4_OperatorMultiply6, "QMatrix4x4_OperatorMultiply6")
	reg(&pg_QMatrix4x4_OperatorDivide, "QMatrix4x4_OperatorDivide")
	reg(&pg_QMatrix4x4_delete, "QMatrix4x4_delete")
}

type QMatrix4x4 struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMatrix4x4) UnsafeGoID() int64     { return this.vid }
func (this *QMatrix4x4) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMatrix4x4(h uintptr) *QMatrix4x4 {
	if h == 0 {
		return nil
	}
	return &QMatrix4x4{h: h}
}

func newownedQMatrix4x4(h uintptr) *QMatrix4x4 {
	o := bindQMatrix4x4(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMatrix4x4).Delete)
	}
	return o
}

func (this *QMatrix4x4) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMatrix4x4) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMatrix4x4(h unsafe.Pointer) *QMatrix4x4 { return bindQMatrix4x4(uintptr(h)) }
func UnsafeNewQMatrix4x4Owned(h unsafe.Pointer) *QMatrix4x4 { return newownedQMatrix4x4(uintptr(h)) }

func (this *QMatrix4x4) UnsafeRelease() { this.h = 0 }

func NewQMatrix4x4() *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_new)
	__o := newownedQMatrix4x4(__r)
	return __o
}

func NewQMatrix4x42(arg1 Initialization) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_new2, uintptr(arg1))
	__o := newownedQMatrix4x4(__r)
	return __o
}

func NewQMatrix4x44(transform *QTransform) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_new4, uintptr(transform.UnsafePointer()))
	__o := newownedQMatrix4x4(__r)
	runtime.KeepAlive(transform)
	return __o
}

func NewQMatrix4x45(arg1 *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQMatrix4x4(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMatrix4x4) Column(index int) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_column, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QMatrix4x4) SetColumn(index int, value *QVector4D) {
	purego.SyscallN(addr_QMatrix4x4_setColumn, this.h, uintptr(index), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QMatrix4x4) Row(index int) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_row, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQVector4D(__r)
}

func (this *QMatrix4x4) SetRow(index int, value *QVector4D) {
	purego.SyscallN(addr_QMatrix4x4_setRow, this.h, uintptr(index), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QMatrix4x4) IsAffine() bool {
	__r0, _, _ := purego.SyscallN(addr_QMatrix4x4_isAffine, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMatrix4x4) IsIdentity() bool {
	__r0, _, _ := purego.SyscallN(addr_QMatrix4x4_isIdentity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMatrix4x4) SetToIdentity() {
	purego.SyscallN(addr_QMatrix4x4_setToIdentity, this.h)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Fill(value float32) {
	pg_QMatrix4x4_fill(this.h, value)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Determinant() float64 {
	__r := pg_QMatrix4x4_determinant(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QMatrix4x4) Inverted() *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_inverted, this.h)
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) Transposed() *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorPlusAssign(other *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorPlusAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMinusAssign(other *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMinusAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMultiplyAssign(other *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiplyAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMultiplyAssignWithFactor(factor float32) *QMatrix4x4 {
	__r := pg_QMatrix4x4_OperatorMultiplyAssign2(this.h, factor)
	runtime.KeepAlive(this)
	return bindQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorDivideAssign(divisor float32) *QMatrix4x4 {
	__r := pg_QMatrix4x4_OperatorDivideAssign(this.h, divisor)
	runtime.KeepAlive(this)
	return bindQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorEqual(other *QMatrix4x4) bool {
	__r0, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QMatrix4x4) OperatorNotEqual(other *QMatrix4x4) bool {
	__r0, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QMatrix4x4) Scale(vector *QVector3D) {
	purego.SyscallN(addr_QMatrix4x4_scale, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
}

func (this *QMatrix4x4) Translate(vector *QVector3D) {
	purego.SyscallN(addr_QMatrix4x4_translate, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
}

func (this *QMatrix4x4) Rotate(angle float32, vector *QVector3D) {
	pg_QMatrix4x4_rotate(this.h, angle, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
}

func (this *QMatrix4x4) Scale2(x float32, y float32) {
	pg_QMatrix4x4_scale2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Scale3(x float32, y float32, z float32) {
	pg_QMatrix4x4_scale3(this.h, x, y, z)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) ScaleWithFactor(factor float32) {
	pg_QMatrix4x4_scale4(this.h, factor)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Translate2(x float32, y float32) {
	pg_QMatrix4x4_translate2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Translate3(x float32, y float32, z float32) {
	pg_QMatrix4x4_translate3(this.h, x, y, z)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Rotate2(angle float32, x float32, y float32) {
	pg_QMatrix4x4_rotate2(this.h, angle, x, y)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) RotateWithQuaternion(quaternion *QQuaternion) {
	purego.SyscallN(addr_QMatrix4x4_rotate3, this.h, uintptr(quaternion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(quaternion)
}

func (this *QMatrix4x4) Ortho(rect *QRect) {
	purego.SyscallN(addr_QMatrix4x4_ortho, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QMatrix4x4) OrthoWithRect(rect *QRectF) {
	purego.SyscallN(addr_QMatrix4x4_ortho2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QMatrix4x4) Ortho2(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	pg_QMatrix4x4_ortho3(this.h, left, right, bottom, top, nearPlane, farPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Frustum(left float32, right float32, bottom float32, top float32, nearPlane float32, farPlane float32) {
	pg_QMatrix4x4_frustum(this.h, left, right, bottom, top, nearPlane, farPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Perspective(verticalAngle float32, aspectRatio float32, nearPlane float32, farPlane float32) {
	pg_QMatrix4x4_perspective(this.h, verticalAngle, aspectRatio, nearPlane, farPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) LookAt(eye *QVector3D, center *QVector3D, up *QVector3D) {
	purego.SyscallN(addr_QMatrix4x4_lookAt, this.h, uintptr(eye.UnsafePointer()), uintptr(center.UnsafePointer()), uintptr(up.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(eye)
	runtime.KeepAlive(center)
	runtime.KeepAlive(up)
}

func (this *QMatrix4x4) Viewport(rect *QRectF) {
	purego.SyscallN(addr_QMatrix4x4_viewport, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QMatrix4x4) Viewport2(left float32, bottom float32, width float32, height float32) {
	pg_QMatrix4x4_viewport2(this.h, left, bottom, width, height)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) FlipCoordinates() {
	purego.SyscallN(addr_QMatrix4x4_flipCoordinates, this.h)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) CopyDataTo() float32 {
	var values_out float32
	purego.SyscallN(addr_QMatrix4x4_copyDataTo, this.h, uintptr(unsafe.Pointer(&values_out)))
	runtime.KeepAlive(this)
	return values_out
}

func (this *QMatrix4x4) ToTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_toTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QMatrix4x4) ToTransformWithDistanceToPlane(distanceToPlane float32) *QTransform {
	__r := pg_QMatrix4x4_toTransform2(this.h, distanceToPlane)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QMatrix4x4) Map(point *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_map, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPoint(__r)
}

func (this *QMatrix4x4) MapWithPoint(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_map2, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QMatrix4x4) Map2(point *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_map3, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQVector3D(__r)
}

func (this *QMatrix4x4) MapVector(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_mapVector, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQVector3D(__r)
}

func (this *QMatrix4x4) Map3(point *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_map4, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQVector4D(__r)
}

func (this *QMatrix4x4) MapRect(rect *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_mapRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRect(__r)
}

func (this *QMatrix4x4) MapRectWithRect(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_mapRect2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QMatrix4x4) Optimize() {
	purego.SyscallN(addr_QMatrix4x4_optimize, this.h)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMatrix4x4) ProjectedRotate(angle float32, x float32, y float32, z float32, distanceToPlane float32) {
	pg_QMatrix4x4_projectedRotate(this.h, angle, x, y, z, distanceToPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) ProjectedRotate2(angle float32, x float32, y float32, z float32) {
	pg_QMatrix4x4_projectedRotate2(this.h, angle, x, y, z)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Flags() QMatrix4x4__Flag {
	__r0, _, _ := purego.SyscallN(addr_QMatrix4x4_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMatrix4x4__Flag(__r)
}

func (this *QMatrix4x4) InvertedWithInvertible() (*QMatrix4x4, bool) {
	var invertible_out int32
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_inverted2, this.h, uintptr(unsafe.Pointer(&invertible_out)))
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r), invertible_out != 0
}

func (this *QMatrix4x4) Rotate3(angle float32, x float32, y float32, z float32) {
	pg_QMatrix4x4_rotate4(this.h, angle, x, y, z)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Viewport3(left float32, bottom float32, width float32, height float32, nearPlane float32) {
	pg_QMatrix4x4_viewport3(this.h, left, bottom, width, height, nearPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) Viewport4(left float32, bottom float32, width float32, height float32, nearPlane float32, farPlane float32) {
	pg_QMatrix4x4_viewport4(this.h, left, bottom, width, height, nearPlane, farPlane)
	runtime.KeepAlive(this)
}

func (this *QMatrix4x4) OperatorMultiply(vector *QVector3D) *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiply, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQVector3D(__r)
}

func (this *QMatrix4x4) OperatorMultiplyWithVector(vector *QVector4D) *QVector4D {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiply2, this.h, uintptr(vector.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(vector)
	return newownedQVector4D(__r)
}

func (this *QMatrix4x4) OperatorPlus(m2 *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorPlus, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMinus(m2 *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMinus, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMultiplyWithM2(m2 *QMatrix4x4) *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiply3, this.h, uintptr(m2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m2)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMinus2() *QMatrix4x4 {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMinus2, this.h)
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorMultiplyWithPoint(point *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiply4, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPoint(__r)
}

func (this *QMatrix4x4) OperatorMultiply2(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QMatrix4x4_OperatorMultiply5, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QMatrix4x4) OperatorMultiplyWithFactor(factor float32) *QMatrix4x4 {
	__r := pg_QMatrix4x4_OperatorMultiply6(this.h, factor)
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) OperatorDivide(divisor float32) *QMatrix4x4 {
	__r := pg_QMatrix4x4_OperatorDivide(this.h, divisor)
	runtime.KeepAlive(this)
	return newownedQMatrix4x4(__r)
}

func (this *QMatrix4x4) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMatrix4x4_delete(this.h)
	}
	this.UnsafeRelease()
}
