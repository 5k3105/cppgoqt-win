package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRegion_new uintptr
	addr_QRegion_new2 uintptr
	addr_QRegion_new3 uintptr
	addr_QRegion_new4 uintptr
	addr_QRegion_new5 uintptr
	addr_QRegion_new6 uintptr
	addr_QRegion_new7 uintptr
	addr_QRegion_new8 uintptr
	addr_QRegion_new9 uintptr
	addr_QRegion_OperatorAssign uintptr
	addr_QRegion_swap uintptr
	addr_QRegion_isEmpty uintptr
	addr_QRegion_isNull uintptr
	addr_QRegion_begin uintptr
	addr_QRegion_cbegin uintptr
	addr_QRegion_end uintptr
	addr_QRegion_cend uintptr
	addr_QRegion_contains uintptr
	addr_QRegion_contains2 uintptr
	addr_QRegion_translate uintptr
	addr_QRegion_translate2 uintptr
	addr_QRegion_translated uintptr
	addr_QRegion_translated2 uintptr
	addr_QRegion_united uintptr
	addr_QRegion_united2 uintptr
	addr_QRegion_intersected uintptr
	addr_QRegion_intersected2 uintptr
	addr_QRegion_subtracted uintptr
	addr_QRegion_xored uintptr
	addr_QRegion_intersects uintptr
	addr_QRegion_intersects2 uintptr
	addr_QRegion_boundingRect uintptr
	addr_QRegion_setRects uintptr
	addr_QRegion_rectCount uintptr
	addr_QRegion_OperatorBitwiseOr uintptr
	addr_QRegion_OperatorPlus uintptr
	addr_QRegion_OperatorPlus2 uintptr
	addr_QRegion_OperatorBitwiseAnd uintptr
	addr_QRegion_OperatorBitwiseAnd2 uintptr
	addr_QRegion_OperatorMinus uintptr
	addr_QRegion_OperatorBitwiseXor uintptr
	addr_QRegion_OperatorBitwiseOrAssign uintptr
	addr_QRegion_OperatorPlusAssign uintptr
	addr_QRegion_OperatorPlusAssign2 uintptr
	addr_QRegion_OperatorBitwiseAndAssign uintptr
	addr_QRegion_OperatorBitwiseAndAssign2 uintptr
	addr_QRegion_OperatorMinusAssign uintptr
	addr_QRegion_OperatorBitwiseXorAssign uintptr
	addr_QRegion_OperatorEqual uintptr
	addr_QRegion_OperatorNotEqual uintptr
	addr_QRegion_ToQVariant uintptr
	addr_QRegion_OperatorMultiply uintptr
	addr_QRegion_toQVariant uintptr
	pg_QRegion_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRegion_new, "QRegion_new")
	symaddr(&addr_QRegion_new2, "QRegion_new2")
	symaddr(&addr_QRegion_new3, "QRegion_new3")
	symaddr(&addr_QRegion_new4, "QRegion_new4")
	symaddr(&addr_QRegion_new5, "QRegion_new5")
	symaddr(&addr_QRegion_new6, "QRegion_new6")
	symaddr(&addr_QRegion_new7, "QRegion_new7")
	symaddr(&addr_QRegion_new8, "QRegion_new8")
	symaddr(&addr_QRegion_new9, "QRegion_new9")
	symaddr(&addr_QRegion_OperatorAssign, "QRegion_OperatorAssign")
	symaddr(&addr_QRegion_swap, "QRegion_swap")
	symaddr(&addr_QRegion_isEmpty, "QRegion_isEmpty")
	symaddr(&addr_QRegion_isNull, "QRegion_isNull")
	symaddr(&addr_QRegion_begin, "QRegion_begin")
	symaddr(&addr_QRegion_cbegin, "QRegion_cbegin")
	symaddr(&addr_QRegion_end, "QRegion_end")
	symaddr(&addr_QRegion_cend, "QRegion_cend")
	symaddr(&addr_QRegion_contains, "QRegion_contains")
	symaddr(&addr_QRegion_contains2, "QRegion_contains2")
	symaddr(&addr_QRegion_translate, "QRegion_translate")
	symaddr(&addr_QRegion_translate2, "QRegion_translate2")
	symaddr(&addr_QRegion_translated, "QRegion_translated")
	symaddr(&addr_QRegion_translated2, "QRegion_translated2")
	symaddr(&addr_QRegion_united, "QRegion_united")
	symaddr(&addr_QRegion_united2, "QRegion_united2")
	symaddr(&addr_QRegion_intersected, "QRegion_intersected")
	symaddr(&addr_QRegion_intersected2, "QRegion_intersected2")
	symaddr(&addr_QRegion_subtracted, "QRegion_subtracted")
	symaddr(&addr_QRegion_xored, "QRegion_xored")
	symaddr(&addr_QRegion_intersects, "QRegion_intersects")
	symaddr(&addr_QRegion_intersects2, "QRegion_intersects2")
	symaddr(&addr_QRegion_boundingRect, "QRegion_boundingRect")
	symaddr(&addr_QRegion_setRects, "QRegion_setRects")
	symaddr(&addr_QRegion_rectCount, "QRegion_rectCount")
	symaddr(&addr_QRegion_OperatorBitwiseOr, "QRegion_OperatorBitwiseOr")
	symaddr(&addr_QRegion_OperatorPlus, "QRegion_OperatorPlus")
	symaddr(&addr_QRegion_OperatorPlus2, "QRegion_OperatorPlus2")
	symaddr(&addr_QRegion_OperatorBitwiseAnd, "QRegion_OperatorBitwiseAnd")
	symaddr(&addr_QRegion_OperatorBitwiseAnd2, "QRegion_OperatorBitwiseAnd2")
	symaddr(&addr_QRegion_OperatorMinus, "QRegion_OperatorMinus")
	symaddr(&addr_QRegion_OperatorBitwiseXor, "QRegion_OperatorBitwiseXor")
	symaddr(&addr_QRegion_OperatorBitwiseOrAssign, "QRegion_OperatorBitwiseOrAssign")
	symaddr(&addr_QRegion_OperatorPlusAssign, "QRegion_OperatorPlusAssign")
	symaddr(&addr_QRegion_OperatorPlusAssign2, "QRegion_OperatorPlusAssign2")
	symaddr(&addr_QRegion_OperatorBitwiseAndAssign, "QRegion_OperatorBitwiseAndAssign")
	symaddr(&addr_QRegion_OperatorBitwiseAndAssign2, "QRegion_OperatorBitwiseAndAssign2")
	symaddr(&addr_QRegion_OperatorMinusAssign, "QRegion_OperatorMinusAssign")
	symaddr(&addr_QRegion_OperatorBitwiseXorAssign, "QRegion_OperatorBitwiseXorAssign")
	symaddr(&addr_QRegion_OperatorEqual, "QRegion_OperatorEqual")
	symaddr(&addr_QRegion_OperatorNotEqual, "QRegion_OperatorNotEqual")
	symaddr(&addr_QRegion_ToQVariant, "QRegion_ToQVariant")
	symaddr(&addr_QRegion_OperatorMultiply, "QRegion_OperatorMultiply")
	symaddr(&addr_QRegion_toQVariant, "QRegion_toQVariant")
	reg(&pg_QRegion_delete, "QRegion_delete")
}

type QRegion struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRegion) UnsafeGoID() int64     { return this.vid }
func (this *QRegion) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRegion(h uintptr) *QRegion {
	if h == 0 {
		return nil
	}
	return &QRegion{h: h}
}

func newownedQRegion(h uintptr) *QRegion {
	o := bindQRegion(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRegion).Delete)
	}
	return o
}

func (this *QRegion) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRegion) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRegion(h unsafe.Pointer) *QRegion { return bindQRegion(uintptr(h)) }
func UnsafeNewQRegionOwned(h unsafe.Pointer) *QRegion { return newownedQRegion(uintptr(h)) }

func (this *QRegion) UnsafeRelease() { this.h = 0 }

func NewQRegion() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new)
	__o := newownedQRegion(__r)
	return __o
}

func NewQRegion2(x int, y int, w int, h int) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new2, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	__o := newownedQRegion(__r)
	return __o
}

func NewQRegion3(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new3, uintptr(r.UnsafePointer()))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(r)
	return __o
}

func NewQRegion4(pa *QPolygon) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new4, uintptr(pa.UnsafePointer()))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(pa)
	return __o
}

func NewQRegion5(region *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new5, uintptr(region.UnsafePointer()))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(region)
	return __o
}

func NewQRegion6(bitmap *QBitmap) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new6, uintptr(bitmap.UnsafePointer()))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(bitmap)
	return __o
}

func NewQRegion7(x int, y int, w int, h int, t QRegion__RegionType) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new7, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(t))
	__o := newownedQRegion(__r)
	return __o
}

func NewQRegion8(r *QRect, t QRegion__RegionType) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new8, uintptr(r.UnsafePointer()), uintptr(t))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(r)
	return __o
}

func NewQRegion9(pa *QPolygon, fillRule FillRule) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_new9, uintptr(pa.UnsafePointer()), uintptr(fillRule))
	__o := newownedQRegion(__r)
	runtime.KeepAlive(pa)
	return __o
}

func (this *QRegion) OperatorAssign(arg1 *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQRegion(__r)
}

func (this *QRegion) Swap(other *QRegion) {
	purego.SyscallN(addr_QRegion_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QRegion) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegion) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegion) Begin() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRegion_begin, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QRegion) Cbegin() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRegion_cbegin, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QRegion) End() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRegion_end, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QRegion) Cend() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRegion_cend, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QRegion) Contains(p *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_contains, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QRegion) ContainsWithQRect(r *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_contains2, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRegion) Translate(dx int, dy int) {
	purego.SyscallN(addr_QRegion_translate, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QRegion) TranslateWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QRegion_translate2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRegion) Translated(dx int, dy int) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_translated, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QRegion) TranslatedWithQPoint(p *QPoint) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_translated2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQRegion(__r)
}

func (this *QRegion) United(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_united, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) UnitedWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_united2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) Intersected(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_intersected, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) IntersectedWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_intersected2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) Subtracted(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_subtracted, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) Xored(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_xored, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) Intersects(r *QRegion) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_intersects, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRegion) IntersectsWithQRect(r *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_intersects2, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRegion) BoundingRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRegion_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRegion) SetRects(rect *QRect, num int) {
	purego.SyscallN(addr_QRegion_setRects, this.h, uintptr(rect.UnsafePointer()), uintptr(num))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QRegion) RectCount() int {
	__r0, _, _ := purego.SyscallN(addr_QRegion_rectCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRegion) OperatorBitwiseOr(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseOr, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorPlus(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorPlus, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorPlusWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorPlus2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorBitwiseAnd(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseAnd, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorBitwiseAndWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseAnd2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorMinus(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorMinus, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorBitwiseXor(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseXor, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRegion(__r)
}

func (this *QRegion) OperatorBitwiseOrAssign(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseOrAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorPlusAssign(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorPlusAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorPlusAssignWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorPlusAssign2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorBitwiseAndAssign(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseAndAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorBitwiseAndAssignWithQRect(r *QRect) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseAndAssign2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorMinusAssign(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorMinusAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorBitwiseXorAssign(r *QRegion) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorBitwiseXorAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRegion(__r)
}

func (this *QRegion) OperatorEqual(r *QRegion) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_OperatorEqual, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRegion) OperatorNotEqual(r *QRegion) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegion_OperatorNotEqual, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRegion) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QRegion_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QRegion) OperatorMultiply(m *QTransform) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QRegion_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQRegion(__r)
}

func (this *QRegion) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QRegion_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QRegion) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRegion_delete(this.h)
	}
	this.UnsafeRelease()
}
