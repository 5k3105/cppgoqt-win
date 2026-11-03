package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPolygon_new uintptr
	addr_QPolygon_new2 uintptr
	addr_QPolygon_new3 uintptr
	addr_QPolygon_new4 uintptr
	addr_QPolygon_new5 uintptr
	addr_QPolygon_swap uintptr
	addr_QPolygon_ToQVariant uintptr
	addr_QPolygon_translate uintptr
	addr_QPolygon_translate2 uintptr
	addr_QPolygon_translated uintptr
	addr_QPolygon_translated2 uintptr
	addr_QPolygon_boundingRect uintptr
	addr_QPolygon_point uintptr
	addr_QPolygon_point2 uintptr
	addr_QPolygon_setPoint uintptr
	addr_QPolygon_setPoint2 uintptr
	addr_QPolygon_setPoints uintptr
	addr_QPolygon_putPoints uintptr
	addr_QPolygon_putPoints2 uintptr
	addr_QPolygon_containsPoint uintptr
	addr_QPolygon_united uintptr
	addr_QPolygon_intersected uintptr
	addr_QPolygon_subtracted uintptr
	addr_QPolygon_intersects uintptr
	addr_QPolygon_toPolygonF uintptr
	addr_QPolygon_maxSize uintptr
	addr_QPolygon_size uintptr
	addr_QPolygon_count uintptr
	addr_QPolygon_length uintptr
	addr_QPolygon_isEmpty uintptr
	addr_QPolygon_resize uintptr
	addr_QPolygon_resize2 uintptr
	addr_QPolygon_resizeForOverwrite uintptr
	addr_QPolygon_capacity uintptr
	addr_QPolygon_reserve uintptr
	addr_QPolygon_squeeze uintptr
	addr_QPolygon_detach uintptr
	addr_QPolygon_isDetached uintptr
	addr_QPolygon_isSharedWith uintptr
	addr_QPolygon_data uintptr
	addr_QPolygon_data2 uintptr
	addr_QPolygon_constData uintptr
	addr_QPolygon_clear uintptr
	addr_QPolygon_at uintptr
	addr_QPolygon_OperatorSubscript uintptr
	addr_QPolygon_OperatorSubscript2 uintptr
	addr_QPolygon_append uintptr
	addr_QPolygon_append2 uintptr
	addr_QPolygon_prepend uintptr
	addr_QPolygon_replace uintptr
	addr_QPolygon_remove uintptr
	addr_QPolygon_removeFirst uintptr
	addr_QPolygon_removeLast uintptr
	addr_QPolygon_takeFirst uintptr
	addr_QPolygon_takeLast uintptr
	addr_QPolygon_removeAt uintptr
	addr_QPolygon_takeAt uintptr
	addr_QPolygon_move uintptr
	addr_QPolygon_first uintptr
	addr_QPolygon_first2 uintptr
	addr_QPolygon_constFirst uintptr
	addr_QPolygon_last uintptr
	addr_QPolygon_last2 uintptr
	addr_QPolygon_constLast uintptr
	addr_QPolygon_startsWith uintptr
	addr_QPolygon_endsWith uintptr
	addr_QPolygon_mid uintptr
	addr_QPolygon_first3 uintptr
	addr_QPolygon_last3 uintptr
	addr_QPolygon_sliced uintptr
	addr_QPolygon_sliced2 uintptr
	addr_QPolygon_value uintptr
	addr_QPolygon_value2 uintptr
	addr_QPolygon_swapItemsAt uintptr
	addr_QPolygon_push_back uintptr
	addr_QPolygon_push_front uintptr
	addr_QPolygon_pop_back uintptr
	addr_QPolygon_pop_front uintptr
	addr_QPolygon_empty uintptr
	addr_QPolygon_front uintptr
	addr_QPolygon_front2 uintptr
	addr_QPolygon_back uintptr
	addr_QPolygon_back2 uintptr
	addr_QPolygon_shrink_to_fit uintptr
	addr_QPolygon_max_size uintptr
	addr_QPolygon_OperatorPlus uintptr
	addr_QPolygon_OperatorPlus2 uintptr
	addr_QPolygon_fromList uintptr
	addr_QPolygon_toList uintptr
	addr_QPolygon_fromVector uintptr
	addr_QPolygon_toVector uintptr
	addr_QPolygon_putPoints3 uintptr
	addr_QPolygon_remove2 uintptr
	addr_QPolygon_mid2 uintptr
	addr_QPolygon_OperatorMultiply uintptr
	addr_QPolygon_toQVariant uintptr
	pg_QPolygon_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPolygon_new, "QPolygon_new")
	symaddr(&addr_QPolygon_new2, "QPolygon_new2")
	symaddr(&addr_QPolygon_new3, "QPolygon_new3")
	symaddr(&addr_QPolygon_new4, "QPolygon_new4")
	symaddr(&addr_QPolygon_new5, "QPolygon_new5")
	symaddr(&addr_QPolygon_swap, "QPolygon_swap")
	symaddr(&addr_QPolygon_ToQVariant, "QPolygon_ToQVariant")
	symaddr(&addr_QPolygon_translate, "QPolygon_translate")
	symaddr(&addr_QPolygon_translate2, "QPolygon_translate2")
	symaddr(&addr_QPolygon_translated, "QPolygon_translated")
	symaddr(&addr_QPolygon_translated2, "QPolygon_translated2")
	symaddr(&addr_QPolygon_boundingRect, "QPolygon_boundingRect")
	symaddr(&addr_QPolygon_point, "QPolygon_point")
	symaddr(&addr_QPolygon_point2, "QPolygon_point2")
	symaddr(&addr_QPolygon_setPoint, "QPolygon_setPoint")
	symaddr(&addr_QPolygon_setPoint2, "QPolygon_setPoint2")
	symaddr(&addr_QPolygon_setPoints, "QPolygon_setPoints")
	symaddr(&addr_QPolygon_putPoints, "QPolygon_putPoints")
	symaddr(&addr_QPolygon_putPoints2, "QPolygon_putPoints2")
	symaddr(&addr_QPolygon_containsPoint, "QPolygon_containsPoint")
	symaddr(&addr_QPolygon_united, "QPolygon_united")
	symaddr(&addr_QPolygon_intersected, "QPolygon_intersected")
	symaddr(&addr_QPolygon_subtracted, "QPolygon_subtracted")
	symaddr(&addr_QPolygon_intersects, "QPolygon_intersects")
	symaddr(&addr_QPolygon_toPolygonF, "QPolygon_toPolygonF")
	symaddr(&addr_QPolygon_maxSize, "QPolygon_maxSize")
	symaddr(&addr_QPolygon_size, "QPolygon_size")
	symaddr(&addr_QPolygon_count, "QPolygon_count")
	symaddr(&addr_QPolygon_length, "QPolygon_length")
	symaddr(&addr_QPolygon_isEmpty, "QPolygon_isEmpty")
	symaddr(&addr_QPolygon_resize, "QPolygon_resize")
	symaddr(&addr_QPolygon_resize2, "QPolygon_resize2")
	symaddr(&addr_QPolygon_resizeForOverwrite, "QPolygon_resizeForOverwrite")
	symaddr(&addr_QPolygon_capacity, "QPolygon_capacity")
	symaddr(&addr_QPolygon_reserve, "QPolygon_reserve")
	symaddr(&addr_QPolygon_squeeze, "QPolygon_squeeze")
	symaddr(&addr_QPolygon_detach, "QPolygon_detach")
	symaddr(&addr_QPolygon_isDetached, "QPolygon_isDetached")
	symaddr(&addr_QPolygon_isSharedWith, "QPolygon_isSharedWith")
	symaddr(&addr_QPolygon_data, "QPolygon_data")
	symaddr(&addr_QPolygon_data2, "QPolygon_data2")
	symaddr(&addr_QPolygon_constData, "QPolygon_constData")
	symaddr(&addr_QPolygon_clear, "QPolygon_clear")
	symaddr(&addr_QPolygon_at, "QPolygon_at")
	symaddr(&addr_QPolygon_OperatorSubscript, "QPolygon_OperatorSubscript")
	symaddr(&addr_QPolygon_OperatorSubscript2, "QPolygon_OperatorSubscript2")
	symaddr(&addr_QPolygon_append, "QPolygon_append")
	symaddr(&addr_QPolygon_append2, "QPolygon_append2")
	symaddr(&addr_QPolygon_prepend, "QPolygon_prepend")
	symaddr(&addr_QPolygon_replace, "QPolygon_replace")
	symaddr(&addr_QPolygon_remove, "QPolygon_remove")
	symaddr(&addr_QPolygon_removeFirst, "QPolygon_removeFirst")
	symaddr(&addr_QPolygon_removeLast, "QPolygon_removeLast")
	symaddr(&addr_QPolygon_takeFirst, "QPolygon_takeFirst")
	symaddr(&addr_QPolygon_takeLast, "QPolygon_takeLast")
	symaddr(&addr_QPolygon_removeAt, "QPolygon_removeAt")
	symaddr(&addr_QPolygon_takeAt, "QPolygon_takeAt")
	symaddr(&addr_QPolygon_move, "QPolygon_move")
	symaddr(&addr_QPolygon_first, "QPolygon_first")
	symaddr(&addr_QPolygon_first2, "QPolygon_first2")
	symaddr(&addr_QPolygon_constFirst, "QPolygon_constFirst")
	symaddr(&addr_QPolygon_last, "QPolygon_last")
	symaddr(&addr_QPolygon_last2, "QPolygon_last2")
	symaddr(&addr_QPolygon_constLast, "QPolygon_constLast")
	symaddr(&addr_QPolygon_startsWith, "QPolygon_startsWith")
	symaddr(&addr_QPolygon_endsWith, "QPolygon_endsWith")
	symaddr(&addr_QPolygon_mid, "QPolygon_mid")
	symaddr(&addr_QPolygon_first3, "QPolygon_first3")
	symaddr(&addr_QPolygon_last3, "QPolygon_last3")
	symaddr(&addr_QPolygon_sliced, "QPolygon_sliced")
	symaddr(&addr_QPolygon_sliced2, "QPolygon_sliced2")
	symaddr(&addr_QPolygon_value, "QPolygon_value")
	symaddr(&addr_QPolygon_value2, "QPolygon_value2")
	symaddr(&addr_QPolygon_swapItemsAt, "QPolygon_swapItemsAt")
	symaddr(&addr_QPolygon_push_back, "QPolygon_push_back")
	symaddr(&addr_QPolygon_push_front, "QPolygon_push_front")
	symaddr(&addr_QPolygon_pop_back, "QPolygon_pop_back")
	symaddr(&addr_QPolygon_pop_front, "QPolygon_pop_front")
	symaddr(&addr_QPolygon_empty, "QPolygon_empty")
	symaddr(&addr_QPolygon_front, "QPolygon_front")
	symaddr(&addr_QPolygon_front2, "QPolygon_front2")
	symaddr(&addr_QPolygon_back, "QPolygon_back")
	symaddr(&addr_QPolygon_back2, "QPolygon_back2")
	symaddr(&addr_QPolygon_shrink_to_fit, "QPolygon_shrink_to_fit")
	symaddr(&addr_QPolygon_max_size, "QPolygon_max_size")
	symaddr(&addr_QPolygon_OperatorPlus, "QPolygon_OperatorPlus")
	symaddr(&addr_QPolygon_OperatorPlus2, "QPolygon_OperatorPlus2")
	symaddr(&addr_QPolygon_fromList, "QPolygon_fromList")
	symaddr(&addr_QPolygon_toList, "QPolygon_toList")
	symaddr(&addr_QPolygon_fromVector, "QPolygon_fromVector")
	symaddr(&addr_QPolygon_toVector, "QPolygon_toVector")
	symaddr(&addr_QPolygon_putPoints3, "QPolygon_putPoints3")
	symaddr(&addr_QPolygon_remove2, "QPolygon_remove2")
	symaddr(&addr_QPolygon_mid2, "QPolygon_mid2")
	symaddr(&addr_QPolygon_OperatorMultiply, "QPolygon_OperatorMultiply")
	symaddr(&addr_QPolygon_toQVariant, "QPolygon_toQVariant")
	reg(&pg_QPolygon_delete, "QPolygon_delete")
}

type QPolygon struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPolygon) UnsafeGoID() int64     { return this.vid }
func (this *QPolygon) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPolygon(h uintptr) *QPolygon {
	if h == 0 {
		return nil
	}
	return &QPolygon{h: h}
}

func newownedQPolygon(h uintptr) *QPolygon {
	o := bindQPolygon(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPolygon).Delete)
	}
	return o
}

func (this *QPolygon) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPolygon) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPolygon(h unsafe.Pointer) *QPolygon { return bindQPolygon(uintptr(h)) }
func UnsafeNewQPolygonOwned(h unsafe.Pointer) *QPolygon { return newownedQPolygon(uintptr(h)) }

func (this *QPolygon) UnsafeRelease() { this.h = 0 }

func NewQPolygon() *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_new)
	__o := newownedQPolygon(__r)
	return __o
}

func NewQPolygon2(v []*QPoint) *QPolygon {
	v_arr := make([]uintptr, len(v))
	for __i, __e := range v {
		v_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var v_p uintptr
	if len(v_arr) > 0 {
		v_p = uintptr(unsafe.Pointer(&v_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QPolygon_new2, v_p, uintptr(len(v)))
	__o := newownedQPolygon(__r)
	runtime.KeepAlive(v)
	runtime.KeepAlive(v_arr)
	return __o
}

func NewQPolygon3(r *QRect) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_new3, uintptr(r.UnsafePointer()))
	__o := newownedQPolygon(__r)
	runtime.KeepAlive(r)
	return __o
}

func NewQPolygon4(arg1 *QPolygon) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQPolygon(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQPolygon5(r *QRect, closed bool) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_new5, uintptr(r.UnsafePointer()), uintptr(b2i(closed)))
	__o := newownedQPolygon(__r)
	runtime.KeepAlive(r)
	return __o
}

func (this *QPolygon) Swap(other *QPolygon) {
	purego.SyscallN(addr_QPolygon_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPolygon) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPolygon_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPolygon) Translate(dx int, dy int) {
	purego.SyscallN(addr_QPolygon_translate, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QPolygon) TranslateWithOffset(offset *QPoint) {
	purego.SyscallN(addr_QPolygon_translate2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QPolygon) Translated(dx int, dy int) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_translated, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
	return newownedQPolygon(__r)
}

func (this *QPolygon) TranslatedWithOffset(offset *QPoint) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_translated2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPolygon(__r)
}

func (this *QPolygon) BoundingRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPolygon_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPolygon) Point(i int) (int, int) {
	var x_out int32
	var y_out int32
	purego.SyscallN(addr_QPolygon_point, this.h, uintptr(i), uintptr(unsafe.Pointer(&x_out)), uintptr(unsafe.Pointer(&y_out)))
	runtime.KeepAlive(this)
	return int(x_out), int(y_out)
}

func (this *QPolygon) PointWithInt(i int) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_point2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPolygon) SetPoint(index int, x int, y int) {
	purego.SyscallN(addr_QPolygon_setPoint, this.h, uintptr(index), uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QPolygon) SetPoint2(index int, p *QPoint) {
	purego.SyscallN(addr_QPolygon_setPoint2, this.h, uintptr(index), uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QPolygon) SetPoints(nPoints int, firstx int, firsty int) {
	purego.SyscallN(addr_QPolygon_setPoints, this.h, uintptr(nPoints), uintptr(firstx), uintptr(firsty))
	runtime.KeepAlive(this)
}

func (this *QPolygon) PutPoints(index int, nPoints int, firstx int, firsty int) {
	purego.SyscallN(addr_QPolygon_putPoints, this.h, uintptr(index), uintptr(nPoints), uintptr(firstx), uintptr(firsty))
	runtime.KeepAlive(this)
}

func (this *QPolygon) PutPoints2(index int, nPoints int, from *QPolygon) {
	purego.SyscallN(addr_QPolygon_putPoints2, this.h, uintptr(index), uintptr(nPoints), uintptr(from.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
}

func (this *QPolygon) ContainsPoint(pt *QPoint, fillRule FillRule) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_containsPoint, this.h, uintptr(pt.UnsafePointer()), uintptr(fillRule))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return __r != 0
}

func (this *QPolygon) United(r *QPolygon) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_united, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygon(__r)
}

func (this *QPolygon) Intersected(r *QPolygon) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_intersected, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygon(__r)
}

func (this *QPolygon) Subtracted(r *QPolygon) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_subtracted, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygon(__r)
}

func (this *QPolygon) Intersects(r *QPolygon) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_intersects, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QPolygon) ToPolygonF() *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygon_toPolygonF, this.h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func QPolygon_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QPolygon) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygon) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygon) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygon) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygon) Resize(size int64) {
	purego.SyscallN(addr_QPolygon_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygon) Resize2(size int64, c *QPoint) {
	purego.SyscallN(addr_QPolygon_resize2, this.h, uintptr(size), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QPolygon) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QPolygon_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygon) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygon) Reserve(size int64) {
	purego.SyscallN(addr_QPolygon_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygon) Squeeze() {
	purego.SyscallN(addr_QPolygon_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) Detach() {
	purego.SyscallN(addr_QPolygon_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygon) IsSharedWith(other []*QPoint) bool {
	other_arr := make([]uintptr, len(other))
	for __i, __e := range other {
		other_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var other_p uintptr
	if len(other_arr) > 0 {
		other_p = uintptr(unsafe.Pointer(&other_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QPolygon_isSharedWith, this.h, other_p, uintptr(len(other)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	runtime.KeepAlive(other_arr)
	return __r != 0
}

func (this *QPolygon) Data() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_data, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Data2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_data2, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) ConstData() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_constData, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Clear() {
	purego.SyscallN(addr_QPolygon_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) At(i int64) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) OperatorSubscript(i int64) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) OperatorSubscriptWith(i int64) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Append(t *QPoint) {
	purego.SyscallN(addr_QPolygon_append, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygon) AppendWith(l []*QPoint) {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	purego.SyscallN(addr_QPolygon_append2, this.h, l_p, uintptr(len(l)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
}

func (this *QPolygon) Prepend(t *QPoint) {
	purego.SyscallN(addr_QPolygon_prepend, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygon) Replace(i int64, t *QPoint) {
	purego.SyscallN(addr_QPolygon_replace, this.h, uintptr(i), uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygon) Remove(i int64) {
	purego.SyscallN(addr_QPolygon_remove, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QPolygon) RemoveFirst() {
	purego.SyscallN(addr_QPolygon_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) RemoveLast() {
	purego.SyscallN(addr_QPolygon_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) TakeFirst() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_takeFirst, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPolygon) TakeLast() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_takeLast, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPolygon) RemoveAt(i int64) {
	purego.SyscallN(addr_QPolygon_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QPolygon) TakeAt(i int64) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPolygon) Move(from int64, to int64) {
	purego.SyscallN(addr_QPolygon_move, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QPolygon) First() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_first, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) First2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_first2, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) ConstFirst() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_constFirst, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Last() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_last, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Last2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_last2, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) ConstLast() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_constLast, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) StartsWith(t *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_startsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QPolygon) EndsWith(t *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_endsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QPolygon) Mid(pos int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_mid, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) FirstWith(n int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_first3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) LastWith(n int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_last3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) Sliced(pos int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_sliced, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) Sliced2(pos int64, n int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_sliced2, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) Value(i int64) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_value, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPolygon) Value2(i int64, defaultValue *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_value2, this.h, uintptr(i), uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQPoint(__r)
}

func (this *QPolygon) SwapItemsAt(i int64, j int64) {
	purego.SyscallN(addr_QPolygon_swapItemsAt, this.h, uintptr(i), uintptr(j))
	runtime.KeepAlive(this)
}

func (this *QPolygon) Push_back(t *QPoint) {
	purego.SyscallN(addr_QPolygon_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygon) Push_front(t *QPoint) {
	purego.SyscallN(addr_QPolygon_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygon) Pop_back() {
	purego.SyscallN(addr_QPolygon_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) Pop_front() {
	purego.SyscallN(addr_QPolygon_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygon) Front() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_front, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Front2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_front2, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Back() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_back, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Back2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPolygon_back2, this.h)
	runtime.KeepAlive(this)
	return bindQPoint(__r)
}

func (this *QPolygon) Shrink_to_fit() {
	purego.SyscallN(addr_QPolygon_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygon) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygon_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygon) OperatorPlus(l []*QPoint) []*QPoint {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_OperatorPlus, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) OperatorPlusWith(l []*QPoint) []*QPoint {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_OperatorPlus2, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPolygon_FromList(list []*QPoint) []*QPoint {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_fromList, list_p, uintptr(len(list)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) ToList() []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_toList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPolygon_FromVector(vector []*QPoint) []*QPoint {
	vector_arr := make([]uintptr, len(vector))
	for __i, __e := range vector {
		vector_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var vector_p uintptr
	if len(vector_arr) > 0 {
		vector_p = uintptr(unsafe.Pointer(&vector_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_fromVector, vector_p, uintptr(len(vector)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(vector)
	runtime.KeepAlive(vector_arr)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) ToVector() []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_toVector, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) PutPoints3(index int, nPoints int, from *QPolygon, fromIndex int) {
	purego.SyscallN(addr_QPolygon_putPoints3, this.h, uintptr(index), uintptr(nPoints), uintptr(from.UnsafePointer()), uintptr(fromIndex))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
}

func (this *QPolygon) Remove2(i int64, n int64) {
	purego.SyscallN(addr_QPolygon_remove2, this.h, uintptr(i), uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QPolygon) Mid2(pos int64, len_ int64) []*QPoint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygon_mid2, this.h, uintptr(pos), uintptr(len_), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPoint, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPoint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygon) OperatorMultiply(m *QTransform) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygon_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQPolygon(__r)
}

func (this *QPolygon) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPolygon_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPolygon) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPolygon_delete(this.h)
	}
	this.UnsafeRelease()
}
