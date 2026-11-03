package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLineF_new uintptr
	addr_QLineF_new2 uintptr
	pg_QLineF_new3 func(float64, float64, float64, float64) uintptr
	addr_QLineF_new4 uintptr
	addr_QLineF_new5 uintptr
	pg_QLineF_fromPolar func(float64, float64) uintptr
	addr_QLineF_isNull uintptr
	addr_QLineF_p1 uintptr
	addr_QLineF_p2 uintptr
	pg_QLineF_x1 func(uintptr) float64
	pg_QLineF_y1 func(uintptr) float64
	pg_QLineF_x2 func(uintptr) float64
	pg_QLineF_y2 func(uintptr) float64
	pg_QLineF_dx func(uintptr) float64
	pg_QLineF_dy func(uintptr) float64
	pg_QLineF_length func(uintptr) float64
	pg_QLineF_setLength func(uintptr, float64)
	pg_QLineF_angle func(uintptr) float64
	pg_QLineF_setAngle func(uintptr, float64)
	pg_QLineF_angleTo func(uintptr, uintptr) float64
	addr_QLineF_unitVector uintptr
	addr_QLineF_normalVector uintptr
	addr_QLineF_intersects uintptr
	pg_QLineF_pointAt func(uintptr, float64) uintptr
	addr_QLineF_translate uintptr
	pg_QLineF_translate2 func(uintptr, float64, float64)
	addr_QLineF_translated uintptr
	pg_QLineF_translated2 func(uintptr, float64, float64) uintptr
	addr_QLineF_center uintptr
	addr_QLineF_setP1 uintptr
	addr_QLineF_setP2 uintptr
	addr_QLineF_setPoints uintptr
	pg_QLineF_setLine func(uintptr, float64, float64, float64, float64)
	addr_QLineF_toLine uintptr
	addr_QLineF_intersects2 uintptr
	addr_QLineF_OperatorEqual uintptr
	addr_QLineF_OperatorNotEqual uintptr
	addr_QLineF_OperatorEqual2 uintptr
	addr_QLineF_OperatorNotEqual2 uintptr
	addr_QLineF_OperatorMultiply uintptr
	pg_QLineF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLineF_new, "QLineF_new")
	symaddr(&addr_QLineF_new2, "QLineF_new2")
	reg(&pg_QLineF_new3, "QLineF_new3")
	symaddr(&addr_QLineF_new4, "QLineF_new4")
	symaddr(&addr_QLineF_new5, "QLineF_new5")
	reg(&pg_QLineF_fromPolar, "QLineF_fromPolar")
	symaddr(&addr_QLineF_isNull, "QLineF_isNull")
	symaddr(&addr_QLineF_p1, "QLineF_p1")
	symaddr(&addr_QLineF_p2, "QLineF_p2")
	reg(&pg_QLineF_x1, "QLineF_x1")
	reg(&pg_QLineF_y1, "QLineF_y1")
	reg(&pg_QLineF_x2, "QLineF_x2")
	reg(&pg_QLineF_y2, "QLineF_y2")
	reg(&pg_QLineF_dx, "QLineF_dx")
	reg(&pg_QLineF_dy, "QLineF_dy")
	reg(&pg_QLineF_length, "QLineF_length")
	reg(&pg_QLineF_setLength, "QLineF_setLength")
	reg(&pg_QLineF_angle, "QLineF_angle")
	reg(&pg_QLineF_setAngle, "QLineF_setAngle")
	reg(&pg_QLineF_angleTo, "QLineF_angleTo")
	symaddr(&addr_QLineF_unitVector, "QLineF_unitVector")
	symaddr(&addr_QLineF_normalVector, "QLineF_normalVector")
	symaddr(&addr_QLineF_intersects, "QLineF_intersects")
	reg(&pg_QLineF_pointAt, "QLineF_pointAt")
	symaddr(&addr_QLineF_translate, "QLineF_translate")
	reg(&pg_QLineF_translate2, "QLineF_translate2")
	symaddr(&addr_QLineF_translated, "QLineF_translated")
	reg(&pg_QLineF_translated2, "QLineF_translated2")
	symaddr(&addr_QLineF_center, "QLineF_center")
	symaddr(&addr_QLineF_setP1, "QLineF_setP1")
	symaddr(&addr_QLineF_setP2, "QLineF_setP2")
	symaddr(&addr_QLineF_setPoints, "QLineF_setPoints")
	reg(&pg_QLineF_setLine, "QLineF_setLine")
	symaddr(&addr_QLineF_toLine, "QLineF_toLine")
	symaddr(&addr_QLineF_intersects2, "QLineF_intersects2")
	symaddr(&addr_QLineF_OperatorEqual, "QLineF_OperatorEqual")
	symaddr(&addr_QLineF_OperatorNotEqual, "QLineF_OperatorNotEqual")
	symaddr(&addr_QLineF_OperatorEqual2, "QLineF_OperatorEqual2")
	symaddr(&addr_QLineF_OperatorNotEqual2, "QLineF_OperatorNotEqual2")
	symaddr(&addr_QLineF_OperatorMultiply, "QLineF_OperatorMultiply")
	reg(&pg_QLineF_delete, "QLineF_delete")
}

type QLineF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLineF) UnsafeGoID() int64     { return this.vid }
func (this *QLineF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLineF(h uintptr) *QLineF {
	if h == 0 {
		return nil
	}
	return &QLineF{h: h}
}

func newownedQLineF(h uintptr) *QLineF {
	o := bindQLineF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLineF).Delete)
	}
	return o
}

func (this *QLineF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLineF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLineF(h unsafe.Pointer) *QLineF { return bindQLineF(uintptr(h)) }
func UnsafeNewQLineFOwned(h unsafe.Pointer) *QLineF { return newownedQLineF(uintptr(h)) }

func (this *QLineF) UnsafeRelease() { this.h = 0 }

func NewQLineF() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_new)
	__o := newownedQLineF(__r)
	return __o
}

func NewQLineF2(pt1 *QPointF, pt2 *QPointF) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_new2, uintptr(pt1.UnsafePointer()), uintptr(pt2.UnsafePointer()))
	__o := newownedQLineF(__r)
	runtime.KeepAlive(pt1)
	runtime.KeepAlive(pt2)
	return __o
}

func NewQLineF3(x1 float64, y1 float64, x2 float64, y2 float64) *QLineF {
	__o := newownedQLineF(pg_QLineF_new3(x1, y1, x2, y2))
	return __o
}

func NewQLineF4(line *QLine) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_new4, uintptr(line.UnsafePointer()))
	__o := newownedQLineF(__r)
	runtime.KeepAlive(line)
	return __o
}

func NewQLineF5(arg1 *QLineF) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQLineF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QLineF_FromPolar(length float64, angle float64) *QLineF {
	__r := pg_QLineF_fromPolar(length, angle)
	return newownedQLineF(__r)
}

func (this *QLineF) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QLineF_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLineF) P1() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QLineF_p1, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLineF) P2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QLineF_p2, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLineF) X1() float64 {
	__r := pg_QLineF_x1(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) Y1() float64 {
	__r := pg_QLineF_y1(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) X2() float64 {
	__r := pg_QLineF_x2(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) Y2() float64 {
	__r := pg_QLineF_y2(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) Dx() float64 {
	__r := pg_QLineF_dx(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) Dy() float64 {
	__r := pg_QLineF_dy(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) Length() float64 {
	__r := pg_QLineF_length(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) SetLength(len_ float64) {
	pg_QLineF_setLength(this.h, len_)
	runtime.KeepAlive(this)
}

func (this *QLineF) Angle() float64 {
	__r := pg_QLineF_angle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLineF) SetAngle(angle float64) {
	pg_QLineF_setAngle(this.h, angle)
	runtime.KeepAlive(this)
}

func (this *QLineF) AngleTo(l *QLineF) float64 {
	__r := pg_QLineF_angleTo(this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return __r
}

func (this *QLineF) UnitVector() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_unitVector, this.h)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QLineF) NormalVector() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_normalVector, this.h)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QLineF) Intersects(l *QLineF) QLineF__IntersectionType {
	__r0, _, _ := purego.SyscallN(addr_QLineF_intersects, this.h, uintptr(l.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return QLineF__IntersectionType(__r)
}

func (this *QLineF) PointAt(t float64) *QPointF {
	__r := pg_QLineF_pointAt(this.h, t)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLineF) Translate(p *QPointF) {
	purego.SyscallN(addr_QLineF_translate, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QLineF) Translate2(dx float64, dy float64) {
	pg_QLineF_translate2(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QLineF) Translated(p *QPointF) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_translated, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQLineF(__r)
}

func (this *QLineF) Translated2(dx float64, dy float64) *QLineF {
	__r := pg_QLineF_translated2(this.h, dx, dy)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QLineF) Center() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QLineF_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLineF) SetP1(p1 *QPointF) {
	purego.SyscallN(addr_QLineF_setP1, this.h, uintptr(p1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
}

func (this *QLineF) SetP2(p2 *QPointF) {
	purego.SyscallN(addr_QLineF_setP2, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
}

func (this *QLineF) SetPoints(p1 *QPointF, p2 *QPointF) {
	purego.SyscallN(addr_QLineF_setPoints, this.h, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
}

func (this *QLineF) SetLine(x1 float64, y1 float64, x2 float64, y2 float64) {
	pg_QLineF_setLine(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
}

func (this *QLineF) ToLine() *QLine {
	__r, _, _ := purego.SyscallN(addr_QLineF_toLine, this.h)
	runtime.KeepAlive(this)
	return newownedQLine(__r)
}

func (this *QLineF) Intersects2(l *QLineF, intersectionPoint *QPointF) QLineF__IntersectionType {
	__r0, _, _ := purego.SyscallN(addr_QLineF_intersects2, this.h, uintptr(l.UnsafePointer()), uintptr(intersectionPoint.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(intersectionPoint)
	return QLineF__IntersectionType(__r)
}

func (this *QLineF) OperatorEqual(rhs *QLineF) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineF_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLineF) OperatorNotEqual(rhs *QLineF) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineF_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLineF) OperatorEqualWithRhs(rhs *QLine) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineF_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLineF) OperatorNotEqualWithRhs(rhs *QLine) bool {
	__r0, _, _ := purego.SyscallN(addr_QLineF_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLineF) OperatorMultiply(m *QTransform) *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLineF_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQLineF(__r)
}

func (this *QLineF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLineF_delete(this.h)
	}
	this.UnsafeRelease()
}
