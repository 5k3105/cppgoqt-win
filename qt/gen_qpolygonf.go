package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPolygonF_new uintptr
	addr_QPolygonF_new2 uintptr
	addr_QPolygonF_new3 uintptr
	addr_QPolygonF_new4 uintptr
	addr_QPolygonF_new5 uintptr
	addr_QPolygonF_swap uintptr
	addr_QPolygonF_ToQVariant uintptr
	pg_QPolygonF_translate func(uintptr, float64, float64)
	addr_QPolygonF_translate2 uintptr
	pg_QPolygonF_translated func(uintptr, float64, float64) uintptr
	addr_QPolygonF_translated2 uintptr
	addr_QPolygonF_toPolygon uintptr
	addr_QPolygonF_isClosed uintptr
	addr_QPolygonF_boundingRect uintptr
	addr_QPolygonF_containsPoint uintptr
	addr_QPolygonF_united uintptr
	addr_QPolygonF_intersected uintptr
	addr_QPolygonF_subtracted uintptr
	addr_QPolygonF_intersects uintptr
	addr_QPolygonF_maxSize uintptr
	addr_QPolygonF_size uintptr
	addr_QPolygonF_count uintptr
	addr_QPolygonF_length uintptr
	addr_QPolygonF_isEmpty uintptr
	addr_QPolygonF_resize uintptr
	addr_QPolygonF_resize2 uintptr
	addr_QPolygonF_resizeForOverwrite uintptr
	addr_QPolygonF_capacity uintptr
	addr_QPolygonF_reserve uintptr
	addr_QPolygonF_squeeze uintptr
	addr_QPolygonF_detach uintptr
	addr_QPolygonF_isDetached uintptr
	addr_QPolygonF_isSharedWith uintptr
	addr_QPolygonF_data uintptr
	addr_QPolygonF_data2 uintptr
	addr_QPolygonF_constData uintptr
	addr_QPolygonF_clear uintptr
	addr_QPolygonF_at uintptr
	addr_QPolygonF_OperatorSubscript uintptr
	addr_QPolygonF_OperatorSubscript2 uintptr
	addr_QPolygonF_append uintptr
	addr_QPolygonF_append2 uintptr
	addr_QPolygonF_prepend uintptr
	addr_QPolygonF_replace uintptr
	addr_QPolygonF_remove uintptr
	addr_QPolygonF_removeFirst uintptr
	addr_QPolygonF_removeLast uintptr
	addr_QPolygonF_takeFirst uintptr
	addr_QPolygonF_takeLast uintptr
	addr_QPolygonF_removeAt uintptr
	addr_QPolygonF_takeAt uintptr
	addr_QPolygonF_move uintptr
	addr_QPolygonF_first uintptr
	addr_QPolygonF_first2 uintptr
	addr_QPolygonF_constFirst uintptr
	addr_QPolygonF_last uintptr
	addr_QPolygonF_last2 uintptr
	addr_QPolygonF_constLast uintptr
	addr_QPolygonF_startsWith uintptr
	addr_QPolygonF_endsWith uintptr
	addr_QPolygonF_mid uintptr
	addr_QPolygonF_first3 uintptr
	addr_QPolygonF_last3 uintptr
	addr_QPolygonF_sliced uintptr
	addr_QPolygonF_sliced2 uintptr
	addr_QPolygonF_value uintptr
	addr_QPolygonF_value2 uintptr
	addr_QPolygonF_swapItemsAt uintptr
	addr_QPolygonF_push_back uintptr
	addr_QPolygonF_push_front uintptr
	addr_QPolygonF_pop_back uintptr
	addr_QPolygonF_pop_front uintptr
	addr_QPolygonF_empty uintptr
	addr_QPolygonF_front uintptr
	addr_QPolygonF_front2 uintptr
	addr_QPolygonF_back uintptr
	addr_QPolygonF_back2 uintptr
	addr_QPolygonF_shrink_to_fit uintptr
	addr_QPolygonF_max_size uintptr
	addr_QPolygonF_OperatorPlus uintptr
	addr_QPolygonF_OperatorPlus2 uintptr
	addr_QPolygonF_fromList uintptr
	addr_QPolygonF_toList uintptr
	addr_QPolygonF_fromVector uintptr
	addr_QPolygonF_toVector uintptr
	addr_QPolygonF_remove2 uintptr
	addr_QPolygonF_mid2 uintptr
	addr_QPolygonF_OperatorMultiply uintptr
	addr_QPolygonF_toQVariant uintptr
	pg_QPolygonF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPolygonF_new, "QPolygonF_new")
	symaddr(&addr_QPolygonF_new2, "QPolygonF_new2")
	symaddr(&addr_QPolygonF_new3, "QPolygonF_new3")
	symaddr(&addr_QPolygonF_new4, "QPolygonF_new4")
	symaddr(&addr_QPolygonF_new5, "QPolygonF_new5")
	symaddr(&addr_QPolygonF_swap, "QPolygonF_swap")
	symaddr(&addr_QPolygonF_ToQVariant, "QPolygonF_ToQVariant")
	reg(&pg_QPolygonF_translate, "QPolygonF_translate")
	symaddr(&addr_QPolygonF_translate2, "QPolygonF_translate2")
	reg(&pg_QPolygonF_translated, "QPolygonF_translated")
	symaddr(&addr_QPolygonF_translated2, "QPolygonF_translated2")
	symaddr(&addr_QPolygonF_toPolygon, "QPolygonF_toPolygon")
	symaddr(&addr_QPolygonF_isClosed, "QPolygonF_isClosed")
	symaddr(&addr_QPolygonF_boundingRect, "QPolygonF_boundingRect")
	symaddr(&addr_QPolygonF_containsPoint, "QPolygonF_containsPoint")
	symaddr(&addr_QPolygonF_united, "QPolygonF_united")
	symaddr(&addr_QPolygonF_intersected, "QPolygonF_intersected")
	symaddr(&addr_QPolygonF_subtracted, "QPolygonF_subtracted")
	symaddr(&addr_QPolygonF_intersects, "QPolygonF_intersects")
	symaddr(&addr_QPolygonF_maxSize, "QPolygonF_maxSize")
	symaddr(&addr_QPolygonF_size, "QPolygonF_size")
	symaddr(&addr_QPolygonF_count, "QPolygonF_count")
	symaddr(&addr_QPolygonF_length, "QPolygonF_length")
	symaddr(&addr_QPolygonF_isEmpty, "QPolygonF_isEmpty")
	symaddr(&addr_QPolygonF_resize, "QPolygonF_resize")
	symaddr(&addr_QPolygonF_resize2, "QPolygonF_resize2")
	symaddr(&addr_QPolygonF_resizeForOverwrite, "QPolygonF_resizeForOverwrite")
	symaddr(&addr_QPolygonF_capacity, "QPolygonF_capacity")
	symaddr(&addr_QPolygonF_reserve, "QPolygonF_reserve")
	symaddr(&addr_QPolygonF_squeeze, "QPolygonF_squeeze")
	symaddr(&addr_QPolygonF_detach, "QPolygonF_detach")
	symaddr(&addr_QPolygonF_isDetached, "QPolygonF_isDetached")
	symaddr(&addr_QPolygonF_isSharedWith, "QPolygonF_isSharedWith")
	symaddr(&addr_QPolygonF_data, "QPolygonF_data")
	symaddr(&addr_QPolygonF_data2, "QPolygonF_data2")
	symaddr(&addr_QPolygonF_constData, "QPolygonF_constData")
	symaddr(&addr_QPolygonF_clear, "QPolygonF_clear")
	symaddr(&addr_QPolygonF_at, "QPolygonF_at")
	symaddr(&addr_QPolygonF_OperatorSubscript, "QPolygonF_OperatorSubscript")
	symaddr(&addr_QPolygonF_OperatorSubscript2, "QPolygonF_OperatorSubscript2")
	symaddr(&addr_QPolygonF_append, "QPolygonF_append")
	symaddr(&addr_QPolygonF_append2, "QPolygonF_append2")
	symaddr(&addr_QPolygonF_prepend, "QPolygonF_prepend")
	symaddr(&addr_QPolygonF_replace, "QPolygonF_replace")
	symaddr(&addr_QPolygonF_remove, "QPolygonF_remove")
	symaddr(&addr_QPolygonF_removeFirst, "QPolygonF_removeFirst")
	symaddr(&addr_QPolygonF_removeLast, "QPolygonF_removeLast")
	symaddr(&addr_QPolygonF_takeFirst, "QPolygonF_takeFirst")
	symaddr(&addr_QPolygonF_takeLast, "QPolygonF_takeLast")
	symaddr(&addr_QPolygonF_removeAt, "QPolygonF_removeAt")
	symaddr(&addr_QPolygonF_takeAt, "QPolygonF_takeAt")
	symaddr(&addr_QPolygonF_move, "QPolygonF_move")
	symaddr(&addr_QPolygonF_first, "QPolygonF_first")
	symaddr(&addr_QPolygonF_first2, "QPolygonF_first2")
	symaddr(&addr_QPolygonF_constFirst, "QPolygonF_constFirst")
	symaddr(&addr_QPolygonF_last, "QPolygonF_last")
	symaddr(&addr_QPolygonF_last2, "QPolygonF_last2")
	symaddr(&addr_QPolygonF_constLast, "QPolygonF_constLast")
	symaddr(&addr_QPolygonF_startsWith, "QPolygonF_startsWith")
	symaddr(&addr_QPolygonF_endsWith, "QPolygonF_endsWith")
	symaddr(&addr_QPolygonF_mid, "QPolygonF_mid")
	symaddr(&addr_QPolygonF_first3, "QPolygonF_first3")
	symaddr(&addr_QPolygonF_last3, "QPolygonF_last3")
	symaddr(&addr_QPolygonF_sliced, "QPolygonF_sliced")
	symaddr(&addr_QPolygonF_sliced2, "QPolygonF_sliced2")
	symaddr(&addr_QPolygonF_value, "QPolygonF_value")
	symaddr(&addr_QPolygonF_value2, "QPolygonF_value2")
	symaddr(&addr_QPolygonF_swapItemsAt, "QPolygonF_swapItemsAt")
	symaddr(&addr_QPolygonF_push_back, "QPolygonF_push_back")
	symaddr(&addr_QPolygonF_push_front, "QPolygonF_push_front")
	symaddr(&addr_QPolygonF_pop_back, "QPolygonF_pop_back")
	symaddr(&addr_QPolygonF_pop_front, "QPolygonF_pop_front")
	symaddr(&addr_QPolygonF_empty, "QPolygonF_empty")
	symaddr(&addr_QPolygonF_front, "QPolygonF_front")
	symaddr(&addr_QPolygonF_front2, "QPolygonF_front2")
	symaddr(&addr_QPolygonF_back, "QPolygonF_back")
	symaddr(&addr_QPolygonF_back2, "QPolygonF_back2")
	symaddr(&addr_QPolygonF_shrink_to_fit, "QPolygonF_shrink_to_fit")
	symaddr(&addr_QPolygonF_max_size, "QPolygonF_max_size")
	symaddr(&addr_QPolygonF_OperatorPlus, "QPolygonF_OperatorPlus")
	symaddr(&addr_QPolygonF_OperatorPlus2, "QPolygonF_OperatorPlus2")
	symaddr(&addr_QPolygonF_fromList, "QPolygonF_fromList")
	symaddr(&addr_QPolygonF_toList, "QPolygonF_toList")
	symaddr(&addr_QPolygonF_fromVector, "QPolygonF_fromVector")
	symaddr(&addr_QPolygonF_toVector, "QPolygonF_toVector")
	symaddr(&addr_QPolygonF_remove2, "QPolygonF_remove2")
	symaddr(&addr_QPolygonF_mid2, "QPolygonF_mid2")
	symaddr(&addr_QPolygonF_OperatorMultiply, "QPolygonF_OperatorMultiply")
	symaddr(&addr_QPolygonF_toQVariant, "QPolygonF_toQVariant")
	reg(&pg_QPolygonF_delete, "QPolygonF_delete")
}

type QPolygonF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPolygonF) UnsafeGoID() int64     { return this.vid }
func (this *QPolygonF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPolygonF(h uintptr) *QPolygonF {
	if h == 0 {
		return nil
	}
	return &QPolygonF{h: h}
}

func newownedQPolygonF(h uintptr) *QPolygonF {
	o := bindQPolygonF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPolygonF).Delete)
	}
	return o
}

func (this *QPolygonF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPolygonF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPolygonF(h unsafe.Pointer) *QPolygonF { return bindQPolygonF(uintptr(h)) }
func UnsafeNewQPolygonFOwned(h unsafe.Pointer) *QPolygonF { return newownedQPolygonF(uintptr(h)) }

func (this *QPolygonF) UnsafeRelease() { this.h = 0 }

func NewQPolygonF() *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_new)
	__o := newownedQPolygonF(__r)
	return __o
}

func NewQPolygonF2(v []*QPointF) *QPolygonF {
	v_arr := make([]uintptr, len(v))
	for __i, __e := range v {
		v_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var v_p uintptr
	if len(v_arr) > 0 {
		v_p = uintptr(unsafe.Pointer(&v_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QPolygonF_new2, v_p, uintptr(len(v)))
	__o := newownedQPolygonF(__r)
	runtime.KeepAlive(v)
	runtime.KeepAlive(v_arr)
	return __o
}

func NewQPolygonF3(r *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_new3, uintptr(r.UnsafePointer()))
	__o := newownedQPolygonF(__r)
	runtime.KeepAlive(r)
	return __o
}

func NewQPolygonF4(a *QPolygon) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_new4, uintptr(a.UnsafePointer()))
	__o := newownedQPolygonF(__r)
	runtime.KeepAlive(a)
	return __o
}

func NewQPolygonF5(arg1 *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQPolygonF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPolygonF) Swap(other *QPolygonF) {
	purego.SyscallN(addr_QPolygonF_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPolygonF) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPolygonF) Translate(dx float64, dy float64) {
	pg_QPolygonF_translate(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) TranslateWithOffset(offset *QPointF) {
	purego.SyscallN(addr_QPolygonF_translate2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QPolygonF) Translated(dx float64, dy float64) *QPolygonF {
	__r := pg_QPolygonF_translated(this.h, dx, dy)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) TranslatedWithOffset(offset *QPointF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_translated2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) ToPolygon() *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_toPolygon, this.h)
	runtime.KeepAlive(this)
	return newownedQPolygon(__r)
}

func (this *QPolygonF) IsClosed() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_isClosed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygonF) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPolygonF) ContainsPoint(pt *QPointF, fillRule FillRule) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_containsPoint, this.h, uintptr(pt.UnsafePointer()), uintptr(fillRule))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return __r != 0
}

func (this *QPolygonF) United(r *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_united, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) Intersected(r *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_intersected, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) Subtracted(r *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_subtracted, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) Intersects(r *QPolygonF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_intersects, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func QPolygonF_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QPolygonF) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygonF) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygonF) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygonF) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygonF) Resize(size int64) {
	purego.SyscallN(addr_QPolygonF_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Resize2(size int64, c *QPointF) {
	purego.SyscallN(addr_QPolygonF_resize2, this.h, uintptr(size), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QPolygonF) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QPolygonF_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygonF) Reserve(size int64) {
	purego.SyscallN(addr_QPolygonF_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Squeeze() {
	purego.SyscallN(addr_QPolygonF_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Detach() {
	purego.SyscallN(addr_QPolygonF_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygonF) IsSharedWith(other []*QPointF) bool {
	other_arr := make([]uintptr, len(other))
	for __i, __e := range other {
		other_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var other_p uintptr
	if len(other_arr) > 0 {
		other_p = uintptr(unsafe.Pointer(&other_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_isSharedWith, this.h, other_p, uintptr(len(other)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	runtime.KeepAlive(other_arr)
	return __r != 0
}

func (this *QPolygonF) Data() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_data, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Data2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_data2, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) ConstData() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_constData, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Clear() {
	purego.SyscallN(addr_QPolygonF_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) At(i int64) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) OperatorSubscript(i int64) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) OperatorSubscriptWith(i int64) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Append(t *QPointF) {
	purego.SyscallN(addr_QPolygonF_append, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygonF) AppendWith(l []*QPointF) {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	purego.SyscallN(addr_QPolygonF_append2, this.h, l_p, uintptr(len(l)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
}

func (this *QPolygonF) Prepend(t *QPointF) {
	purego.SyscallN(addr_QPolygonF_prepend, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygonF) Replace(i int64, t *QPointF) {
	purego.SyscallN(addr_QPolygonF_replace, this.h, uintptr(i), uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygonF) Remove(i int64) {
	purego.SyscallN(addr_QPolygonF_remove, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) RemoveFirst() {
	purego.SyscallN(addr_QPolygonF_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) RemoveLast() {
	purego.SyscallN(addr_QPolygonF_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) TakeFirst() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_takeFirst, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPolygonF) TakeLast() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_takeLast, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPolygonF) RemoveAt(i int64) {
	purego.SyscallN(addr_QPolygonF_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) TakeAt(i int64) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPolygonF) Move(from int64, to int64) {
	purego.SyscallN(addr_QPolygonF_move, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) First() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_first, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) First2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_first2, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) ConstFirst() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_constFirst, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Last() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_last, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Last2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_last2, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) ConstLast() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_constLast, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) StartsWith(t *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_startsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QPolygonF) EndsWith(t *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_endsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QPolygonF) Mid(pos int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_mid, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) FirstWith(n int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_first3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) LastWith(n int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_last3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) Sliced(pos int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_sliced, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) Sliced2(pos int64, n int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_sliced2, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) Value(i int64) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_value, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPolygonF) Value2(i int64, defaultValue *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_value2, this.h, uintptr(i), uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
	return newownedQPointF(__r)
}

func (this *QPolygonF) SwapItemsAt(i int64, j int64) {
	purego.SyscallN(addr_QPolygonF_swapItemsAt, this.h, uintptr(i), uintptr(j))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Push_back(t *QPointF) {
	purego.SyscallN(addr_QPolygonF_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygonF) Push_front(t *QPointF) {
	purego.SyscallN(addr_QPolygonF_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QPolygonF) Pop_back() {
	purego.SyscallN(addr_QPolygonF_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Pop_front() {
	purego.SyscallN(addr_QPolygonF_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPolygonF) Front() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_front, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Front2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_front2, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Back() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_back, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Back2() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_back2, this.h)
	runtime.KeepAlive(this)
	return bindQPointF(__r)
}

func (this *QPolygonF) Shrink_to_fit() {
	purego.SyscallN(addr_QPolygonF_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPolygonF_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPolygonF) OperatorPlus(l []*QPointF) []*QPointF {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_OperatorPlus, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) OperatorPlusWith(l []*QPointF) []*QPointF {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_OperatorPlus2, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPolygonF_FromList(list []*QPointF) []*QPointF {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_fromList, list_p, uintptr(len(list)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) ToList() []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_toList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QPolygonF_FromVector(vector []*QPointF) []*QPointF {
	vector_arr := make([]uintptr, len(vector))
	for __i, __e := range vector {
		vector_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var vector_p uintptr
	if len(vector_arr) > 0 {
		vector_p = uintptr(unsafe.Pointer(&vector_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_fromVector, vector_p, uintptr(len(vector)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(vector)
	runtime.KeepAlive(vector_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) ToVector() []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_toVector, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) Remove2(i int64, n int64) {
	purego.SyscallN(addr_QPolygonF_remove2, this.h, uintptr(i), uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QPolygonF) Mid2(pos int64, len_ int64) []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPolygonF_mid2, this.h, uintptr(pos), uintptr(len_), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QPolygonF) OperatorMultiply(m *QTransform) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQPolygonF(__r)
}

func (this *QPolygonF) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPolygonF_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPolygonF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPolygonF_delete(this.h)
	}
	this.UnsafeRelease()
}
