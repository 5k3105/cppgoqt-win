package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTransform_new uintptr
	addr_QTransform_new2 uintptr
	pg_QTransform_new3 func(float64, float64, float64, float64, float64, float64, float64, float64, float64) uintptr
	pg_QTransform_new4 func(float64, float64, float64, float64, float64, float64) uintptr
	addr_QTransform_new5 uintptr
	addr_QTransform_OperatorAssign uintptr
	addr_QTransform_isAffine uintptr
	addr_QTransform_isIdentity uintptr
	addr_QTransform_isInvertible uintptr
	addr_QTransform_isScaling uintptr
	addr_QTransform_isRotating uintptr
	addr_QTransform_isTranslating uintptr
	addr_QTransform_type uintptr
	pg_QTransform_determinant func(uintptr) float64
	pg_QTransform_m11 func(uintptr) float64
	pg_QTransform_m12 func(uintptr) float64
	pg_QTransform_m13 func(uintptr) float64
	pg_QTransform_m21 func(uintptr) float64
	pg_QTransform_m22 func(uintptr) float64
	pg_QTransform_m23 func(uintptr) float64
	pg_QTransform_m31 func(uintptr) float64
	pg_QTransform_m32 func(uintptr) float64
	pg_QTransform_m33 func(uintptr) float64
	pg_QTransform_dx func(uintptr) float64
	pg_QTransform_dy func(uintptr) float64
	pg_QTransform_setMatrix func(uintptr, float64, float64, float64, float64, float64, float64, float64, float64, float64)
	addr_QTransform_inverted uintptr
	addr_QTransform_adjoint uintptr
	addr_QTransform_transposed uintptr
	pg_QTransform_translate func(uintptr, float64, float64) uintptr
	pg_QTransform_scale func(uintptr, float64, float64) uintptr
	pg_QTransform_shear func(uintptr, float64, float64) uintptr
	pg_QTransform_rotate func(uintptr, float64, int32, float64) uintptr
	pg_QTransform_rotate2 func(uintptr, float64) uintptr
	pg_QTransform_rotateRadians func(uintptr, float64, int32, float64) uintptr
	pg_QTransform_rotateRadians2 func(uintptr, float64) uintptr
	addr_QTransform_squareToQuad uintptr
	addr_QTransform_quadToSquare uintptr
	addr_QTransform_quadToQuad uintptr
	addr_QTransform_OperatorEqual uintptr
	addr_QTransform_OperatorNotEqual uintptr
	addr_QTransform_OperatorMultiplyAssign uintptr
	addr_QTransform_OperatorMultiply uintptr
	addr_QTransform_ToQVariant uintptr
	addr_QTransform_reset uintptr
	addr_QTransform_map uintptr
	addr_QTransform_map2 uintptr
	addr_QTransform_map3 uintptr
	addr_QTransform_map4 uintptr
	addr_QTransform_map5 uintptr
	addr_QTransform_map6 uintptr
	addr_QTransform_map7 uintptr
	addr_QTransform_map8 uintptr
	addr_QTransform_mapToPolygon uintptr
	addr_QTransform_mapRect uintptr
	addr_QTransform_mapRect2 uintptr
	addr_QTransform_map9 uintptr
	pg_QTransform_map10 func(uintptr, float64, float64, uintptr, uintptr)
	pg_QTransform_OperatorMultiplyAssign2 func(uintptr, float64) uintptr
	pg_QTransform_OperatorDivideAssign func(uintptr, float64) uintptr
	pg_QTransform_OperatorPlusAssign func(uintptr, float64) uintptr
	pg_QTransform_OperatorMinusAssign func(uintptr, float64) uintptr
	pg_QTransform_fromTranslate func(float64, float64) uintptr
	pg_QTransform_fromScale func(float64, float64) uintptr
	addr_QTransform_inverted2 uintptr
	pg_QTransform_rotate3 func(uintptr, float64, int32) uintptr
	pg_QTransform_rotateRadians3 func(uintptr, float64, int32) uintptr
	pg_QTransform_OperatorMultiply2 func(uintptr, float64) uintptr
	pg_QTransform_OperatorDivide func(uintptr, float64) uintptr
	pg_QTransform_OperatorPlus func(uintptr, float64) uintptr
	pg_QTransform_OperatorMinus func(uintptr, float64) uintptr
	addr_QTransform_toQVariant uintptr
	pg_QTransform_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTransform_new, "QTransform_new")
	symaddr(&addr_QTransform_new2, "QTransform_new2")
	reg(&pg_QTransform_new3, "QTransform_new3")
	reg(&pg_QTransform_new4, "QTransform_new4")
	symaddr(&addr_QTransform_new5, "QTransform_new5")
	symaddr(&addr_QTransform_OperatorAssign, "QTransform_OperatorAssign")
	symaddr(&addr_QTransform_isAffine, "QTransform_isAffine")
	symaddr(&addr_QTransform_isIdentity, "QTransform_isIdentity")
	symaddr(&addr_QTransform_isInvertible, "QTransform_isInvertible")
	symaddr(&addr_QTransform_isScaling, "QTransform_isScaling")
	symaddr(&addr_QTransform_isRotating, "QTransform_isRotating")
	symaddr(&addr_QTransform_isTranslating, "QTransform_isTranslating")
	symaddr(&addr_QTransform_type, "QTransform_type")
	reg(&pg_QTransform_determinant, "QTransform_determinant")
	reg(&pg_QTransform_m11, "QTransform_m11")
	reg(&pg_QTransform_m12, "QTransform_m12")
	reg(&pg_QTransform_m13, "QTransform_m13")
	reg(&pg_QTransform_m21, "QTransform_m21")
	reg(&pg_QTransform_m22, "QTransform_m22")
	reg(&pg_QTransform_m23, "QTransform_m23")
	reg(&pg_QTransform_m31, "QTransform_m31")
	reg(&pg_QTransform_m32, "QTransform_m32")
	reg(&pg_QTransform_m33, "QTransform_m33")
	reg(&pg_QTransform_dx, "QTransform_dx")
	reg(&pg_QTransform_dy, "QTransform_dy")
	reg(&pg_QTransform_setMatrix, "QTransform_setMatrix")
	symaddr(&addr_QTransform_inverted, "QTransform_inverted")
	symaddr(&addr_QTransform_adjoint, "QTransform_adjoint")
	symaddr(&addr_QTransform_transposed, "QTransform_transposed")
	reg(&pg_QTransform_translate, "QTransform_translate")
	reg(&pg_QTransform_scale, "QTransform_scale")
	reg(&pg_QTransform_shear, "QTransform_shear")
	reg(&pg_QTransform_rotate, "QTransform_rotate")
	reg(&pg_QTransform_rotate2, "QTransform_rotate2")
	reg(&pg_QTransform_rotateRadians, "QTransform_rotateRadians")
	reg(&pg_QTransform_rotateRadians2, "QTransform_rotateRadians2")
	symaddr(&addr_QTransform_squareToQuad, "QTransform_squareToQuad")
	symaddr(&addr_QTransform_quadToSquare, "QTransform_quadToSquare")
	symaddr(&addr_QTransform_quadToQuad, "QTransform_quadToQuad")
	symaddr(&addr_QTransform_OperatorEqual, "QTransform_OperatorEqual")
	symaddr(&addr_QTransform_OperatorNotEqual, "QTransform_OperatorNotEqual")
	symaddr(&addr_QTransform_OperatorMultiplyAssign, "QTransform_OperatorMultiplyAssign")
	symaddr(&addr_QTransform_OperatorMultiply, "QTransform_OperatorMultiply")
	symaddr(&addr_QTransform_ToQVariant, "QTransform_ToQVariant")
	symaddr(&addr_QTransform_reset, "QTransform_reset")
	symaddr(&addr_QTransform_map, "QTransform_map")
	symaddr(&addr_QTransform_map2, "QTransform_map2")
	symaddr(&addr_QTransform_map3, "QTransform_map3")
	symaddr(&addr_QTransform_map4, "QTransform_map4")
	symaddr(&addr_QTransform_map5, "QTransform_map5")
	symaddr(&addr_QTransform_map6, "QTransform_map6")
	symaddr(&addr_QTransform_map7, "QTransform_map7")
	symaddr(&addr_QTransform_map8, "QTransform_map8")
	symaddr(&addr_QTransform_mapToPolygon, "QTransform_mapToPolygon")
	symaddr(&addr_QTransform_mapRect, "QTransform_mapRect")
	symaddr(&addr_QTransform_mapRect2, "QTransform_mapRect2")
	symaddr(&addr_QTransform_map9, "QTransform_map9")
	reg(&pg_QTransform_map10, "QTransform_map10")
	reg(&pg_QTransform_OperatorMultiplyAssign2, "QTransform_OperatorMultiplyAssign2")
	reg(&pg_QTransform_OperatorDivideAssign, "QTransform_OperatorDivideAssign")
	reg(&pg_QTransform_OperatorPlusAssign, "QTransform_OperatorPlusAssign")
	reg(&pg_QTransform_OperatorMinusAssign, "QTransform_OperatorMinusAssign")
	reg(&pg_QTransform_fromTranslate, "QTransform_fromTranslate")
	reg(&pg_QTransform_fromScale, "QTransform_fromScale")
	symaddr(&addr_QTransform_inverted2, "QTransform_inverted2")
	reg(&pg_QTransform_rotate3, "QTransform_rotate3")
	reg(&pg_QTransform_rotateRadians3, "QTransform_rotateRadians3")
	reg(&pg_QTransform_OperatorMultiply2, "QTransform_OperatorMultiply2")
	reg(&pg_QTransform_OperatorDivide, "QTransform_OperatorDivide")
	reg(&pg_QTransform_OperatorPlus, "QTransform_OperatorPlus")
	reg(&pg_QTransform_OperatorMinus, "QTransform_OperatorMinus")
	symaddr(&addr_QTransform_toQVariant, "QTransform_toQVariant")
	reg(&pg_QTransform_delete, "QTransform_delete")
}

type QTransform struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTransform) UnsafeGoID() int64     { return this.vid }
func (this *QTransform) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTransform(h uintptr) *QTransform {
	if h == 0 {
		return nil
	}
	return &QTransform{h: h}
}

func newownedQTransform(h uintptr) *QTransform {
	o := bindQTransform(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTransform).Delete)
	}
	return o
}

func (this *QTransform) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTransform) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTransform(h unsafe.Pointer) *QTransform { return bindQTransform(uintptr(h)) }
func UnsafeNewQTransformOwned(h unsafe.Pointer) *QTransform { return newownedQTransform(uintptr(h)) }

func (this *QTransform) UnsafeRelease() { this.h = 0 }

func NewQTransform(arg1 Initialization) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_new, uintptr(arg1))
	__o := newownedQTransform(__r)
	return __o
}

func NewQTransform2() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_new2)
	__o := newownedQTransform(__r)
	return __o
}

func NewQTransform3(h11 float64, h12 float64, h13 float64, h21 float64, h22 float64, h23 float64, h31 float64, h32 float64, h33 float64) *QTransform {
	__o := newownedQTransform(pg_QTransform_new3(h11, h12, h13, h21, h22, h23, h31, h32, h33))
	return __o
}

func NewQTransform4(h11 float64, h12 float64, h21 float64, h22 float64, dx float64, dy float64) *QTransform {
	__o := newownedQTransform(pg_QTransform_new4(h11, h12, h21, h22, dx, dy))
	return __o
}

func NewQTransform5(other *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_new5, uintptr(other.UnsafePointer()))
	__o := newownedQTransform(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QTransform) OperatorAssign(arg1 *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTransform(__r)
}

func (this *QTransform) IsAffine() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isAffine, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) IsIdentity() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isIdentity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) IsInvertible() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isInvertible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) IsScaling() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isScaling, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) IsRotating() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isRotating, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) IsTranslating() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_isTranslating, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransform) Type() QTransform__TransformationType {
	__r0, _, _ := purego.SyscallN(addr_QTransform_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTransform__TransformationType(__r)
}

func (this *QTransform) Determinant() float64 {
	__r := pg_QTransform_determinant(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M11() float64 {
	__r := pg_QTransform_m11(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M12() float64 {
	__r := pg_QTransform_m12(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M13() float64 {
	__r := pg_QTransform_m13(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M21() float64 {
	__r := pg_QTransform_m21(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M22() float64 {
	__r := pg_QTransform_m22(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M23() float64 {
	__r := pg_QTransform_m23(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M31() float64 {
	__r := pg_QTransform_m31(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M32() float64 {
	__r := pg_QTransform_m32(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) M33() float64 {
	__r := pg_QTransform_m33(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) Dx() float64 {
	__r := pg_QTransform_dx(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) Dy() float64 {
	__r := pg_QTransform_dy(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTransform) SetMatrix(m11 float64, m12 float64, m13 float64, m21 float64, m22 float64, m23 float64, m31 float64, m32 float64, m33 float64) {
	pg_QTransform_setMatrix(this.h, m11, m12, m13, m21, m22, m23, m31, m32, m33)
	runtime.KeepAlive(this)
}

func (this *QTransform) Inverted() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_inverted, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) Adjoint() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_adjoint, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) Transposed() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) Translate(dx float64, dy float64) *QTransform {
	__r := pg_QTransform_translate(this.h, dx, dy)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) Scale(sx float64, sy float64) *QTransform {
	__r := pg_QTransform_scale(this.h, sx, sy)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) Shear(sh float64, sv float64) *QTransform {
	__r := pg_QTransform_shear(this.h, sh, sv)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) Rotate(a float64, axis Axis, distanceToPlane float64) *QTransform {
	__r := pg_QTransform_rotate(this.h, a, int32(axis), distanceToPlane)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) RotateWith(a float64) *QTransform {
	__r := pg_QTransform_rotate2(this.h, a)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) RotateRadians(a float64, axis Axis, distanceToPlane float64) *QTransform {
	__r := pg_QTransform_rotateRadians(this.h, a, int32(axis), distanceToPlane)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) RotateRadiansWith(a float64) *QTransform {
	__r := pg_QTransform_rotateRadians2(this.h, a)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func QTransform_SquareToQuad(square *QPolygonF, result *QTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_squareToQuad, uintptr(square.UnsafePointer()), uintptr(result.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(square)
	runtime.KeepAlive(result)
	return __r != 0
}

func QTransform_QuadToSquare(quad *QPolygonF, result *QTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_quadToSquare, uintptr(quad.UnsafePointer()), uintptr(result.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(quad)
	runtime.KeepAlive(result)
	return __r != 0
}

func QTransform_QuadToQuad(one *QPolygonF, two *QPolygonF, result *QTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_quadToQuad, uintptr(one.UnsafePointer()), uintptr(two.UnsafePointer()), uintptr(result.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(one)
	runtime.KeepAlive(two)
	runtime.KeepAlive(result)
	return __r != 0
}

func (this *QTransform) OperatorEqual(arg1 *QTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTransform) OperatorNotEqual(arg1 *QTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransform_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QTransform) OperatorMultiplyAssign(arg1 *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_OperatorMultiplyAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTransform(__r)
}

func (this *QTransform) OperatorMultiply(o *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QTransform_OperatorMultiply, this.h, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return newownedQTransform(__r)
}

func (this *QTransform) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTransform_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTransform) Reset() {
	purego.SyscallN(addr_QTransform_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransform) Map(p *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QTransform_map, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQPoint(__r)
}

func (this *QTransform) MapWithQPointF(p *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QTransform_map2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQPointF(__r)
}

func (this *QTransform) MapWithQLine(l *QLine) *QLine {
	__r, _, _ := purego.SyscallN(addr_QTransform_map3, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return newownedQLine(__r)
}

func (this *QTransform) MapWithQLineF(l *QLineF) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QTransform_map4, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return newownedQLineF(__r)
}

func (this *QTransform) MapWithQPolygonF(a *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QTransform_map5, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return newownedQPolygonF(__r)
}

func (this *QTransform) MapWithQPolygon(a *QPolygon) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QTransform_map6, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
	return newownedQPolygon(__r)
}

func (this *QTransform) MapWithQRegion(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QTransform_map7, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QTransform) MapWithQPainterPath(p *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QTransform_map8, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQPainterPath(__r)
}

func (this *QTransform) MapToPolygon(r *QRect) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QTransform_mapToPolygon, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygon(__r)
}

func (this *QTransform) MapRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QTransform_mapRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func (this *QTransform) MapRectWithQRectF(arg1 *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QTransform_mapRect2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRectF(__r)
}

func (this *QTransform) Map2(x int, y int) (int, int) {
	var tx_out int32
	var ty_out int32
	purego.SyscallN(addr_QTransform_map9, this.h, uintptr(x), uintptr(y), uintptr(unsafe.Pointer(&tx_out)), uintptr(unsafe.Pointer(&ty_out)))
	runtime.KeepAlive(this)
	return int(tx_out), int(ty_out)
}

func (this *QTransform) Map3(x float64, y float64) (float64, float64) {
	var tx_out float64
	var ty_out float64
	pg_QTransform_map10(this.h, x, y, uintptr(unsafe.Pointer(&tx_out)), uintptr(unsafe.Pointer(&ty_out)))
	runtime.KeepAlive(this)
	return tx_out, ty_out
}

func (this *QTransform) OperatorMultiplyAssignWithDiv(div float64) *QTransform {
	__r := pg_QTransform_OperatorMultiplyAssign2(this.h, div)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) OperatorDivideAssign(div float64) *QTransform {
	__r := pg_QTransform_OperatorDivideAssign(this.h, div)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) OperatorPlusAssign(div float64) *QTransform {
	__r := pg_QTransform_OperatorPlusAssign(this.h, div)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) OperatorMinusAssign(div float64) *QTransform {
	__r := pg_QTransform_OperatorMinusAssign(this.h, div)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func QTransform_FromTranslate(dx float64, dy float64) *QTransform {
	__r := pg_QTransform_fromTranslate(dx, dy)
	return newownedQTransform(__r)
}

func QTransform_FromScale(dx float64, dy float64) *QTransform {
	__r := pg_QTransform_fromScale(dx, dy)
	return newownedQTransform(__r)
}

func (this *QTransform) InvertedWithInvertible() (*QTransform, bool) {
	var invertible_out int32
	__r, _, _ := purego.SyscallN(addr_QTransform_inverted2, this.h, uintptr(unsafe.Pointer(&invertible_out)))
	runtime.KeepAlive(this)
	return newownedQTransform(__r), invertible_out != 0
}

func (this *QTransform) Rotate2(a float64, axis Axis) *QTransform {
	__r := pg_QTransform_rotate3(this.h, a, int32(axis))
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) RotateRadians2(a float64, axis Axis) *QTransform {
	__r := pg_QTransform_rotateRadians3(this.h, a, int32(axis))
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QTransform) OperatorMultiplyWith(n float64) *QTransform {
	__r := pg_QTransform_OperatorMultiply2(this.h, n)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) OperatorDivide(n float64) *QTransform {
	__r := pg_QTransform_OperatorDivide(this.h, n)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) OperatorPlus(n float64) *QTransform {
	__r := pg_QTransform_OperatorPlus(this.h, n)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) OperatorMinus(n float64) *QTransform {
	__r := pg_QTransform_OperatorMinus(this.h, n)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QTransform) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTransform_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTransform) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTransform_delete(this.h)
	}
	this.UnsafeRelease()
}
