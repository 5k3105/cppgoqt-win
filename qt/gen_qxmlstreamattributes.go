package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamAttributes_new uintptr
	addr_QXmlStreamAttributes_value uintptr
	addr_QXmlStreamAttributes_value2 uintptr
	addr_QXmlStreamAttributes_append uintptr
	addr_QXmlStreamAttributes_append2 uintptr
	addr_QXmlStreamAttributes_hasAttribute uintptr
	addr_QXmlStreamAttributes_hasAttribute2 uintptr
	addr_QXmlStreamAttributes_maxSize uintptr
	addr_QXmlStreamAttributes_size uintptr
	addr_QXmlStreamAttributes_count uintptr
	addr_QXmlStreamAttributes_length uintptr
	addr_QXmlStreamAttributes_isEmpty uintptr
	addr_QXmlStreamAttributes_resize uintptr
	addr_QXmlStreamAttributes_resize2 uintptr
	addr_QXmlStreamAttributes_resizeForOverwrite uintptr
	addr_QXmlStreamAttributes_capacity uintptr
	addr_QXmlStreamAttributes_reserve uintptr
	addr_QXmlStreamAttributes_squeeze uintptr
	addr_QXmlStreamAttributes_detach uintptr
	addr_QXmlStreamAttributes_isDetached uintptr
	addr_QXmlStreamAttributes_isSharedWith uintptr
	addr_QXmlStreamAttributes_data uintptr
	addr_QXmlStreamAttributes_data2 uintptr
	addr_QXmlStreamAttributes_constData uintptr
	addr_QXmlStreamAttributes_clear uintptr
	addr_QXmlStreamAttributes_at uintptr
	addr_QXmlStreamAttributes_OperatorSubscript uintptr
	addr_QXmlStreamAttributes_OperatorSubscript2 uintptr
	addr_QXmlStreamAttributes_prepend uintptr
	addr_QXmlStreamAttributes_replace uintptr
	addr_QXmlStreamAttributes_remove uintptr
	addr_QXmlStreamAttributes_removeFirst uintptr
	addr_QXmlStreamAttributes_removeLast uintptr
	addr_QXmlStreamAttributes_takeFirst uintptr
	addr_QXmlStreamAttributes_takeLast uintptr
	addr_QXmlStreamAttributes_removeAt uintptr
	addr_QXmlStreamAttributes_takeAt uintptr
	addr_QXmlStreamAttributes_move uintptr
	addr_QXmlStreamAttributes_first uintptr
	addr_QXmlStreamAttributes_first2 uintptr
	addr_QXmlStreamAttributes_constFirst uintptr
	addr_QXmlStreamAttributes_last uintptr
	addr_QXmlStreamAttributes_last2 uintptr
	addr_QXmlStreamAttributes_constLast uintptr
	addr_QXmlStreamAttributes_startsWith uintptr
	addr_QXmlStreamAttributes_endsWith uintptr
	addr_QXmlStreamAttributes_mid uintptr
	addr_QXmlStreamAttributes_first3 uintptr
	addr_QXmlStreamAttributes_last3 uintptr
	addr_QXmlStreamAttributes_sliced uintptr
	addr_QXmlStreamAttributes_sliced2 uintptr
	addr_QXmlStreamAttributes_swapItemsAt uintptr
	addr_QXmlStreamAttributes_push_back uintptr
	addr_QXmlStreamAttributes_push_front uintptr
	addr_QXmlStreamAttributes_pop_back uintptr
	addr_QXmlStreamAttributes_pop_front uintptr
	addr_QXmlStreamAttributes_empty uintptr
	addr_QXmlStreamAttributes_front uintptr
	addr_QXmlStreamAttributes_front2 uintptr
	addr_QXmlStreamAttributes_back uintptr
	addr_QXmlStreamAttributes_back2 uintptr
	addr_QXmlStreamAttributes_shrink_to_fit uintptr
	addr_QXmlStreamAttributes_max_size uintptr
	addr_QXmlStreamAttributes_OperatorPlus uintptr
	addr_QXmlStreamAttributes_OperatorPlus2 uintptr
	addr_QXmlStreamAttributes_fromList uintptr
	addr_QXmlStreamAttributes_toList uintptr
	addr_QXmlStreamAttributes_fromVector uintptr
	addr_QXmlStreamAttributes_toVector uintptr
	addr_QXmlStreamAttributes_remove2 uintptr
	addr_QXmlStreamAttributes_mid2 uintptr
	pg_QXmlStreamAttributes_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamAttributes_new, "QXmlStreamAttributes_new")
	symaddr(&addr_QXmlStreamAttributes_value, "QXmlStreamAttributes_value")
	symaddr(&addr_QXmlStreamAttributes_value2, "QXmlStreamAttributes_value2")
	symaddr(&addr_QXmlStreamAttributes_append, "QXmlStreamAttributes_append")
	symaddr(&addr_QXmlStreamAttributes_append2, "QXmlStreamAttributes_append2")
	symaddr(&addr_QXmlStreamAttributes_hasAttribute, "QXmlStreamAttributes_hasAttribute")
	symaddr(&addr_QXmlStreamAttributes_hasAttribute2, "QXmlStreamAttributes_hasAttribute2")
	symaddr(&addr_QXmlStreamAttributes_maxSize, "QXmlStreamAttributes_maxSize")
	symaddr(&addr_QXmlStreamAttributes_size, "QXmlStreamAttributes_size")
	symaddr(&addr_QXmlStreamAttributes_count, "QXmlStreamAttributes_count")
	symaddr(&addr_QXmlStreamAttributes_length, "QXmlStreamAttributes_length")
	symaddr(&addr_QXmlStreamAttributes_isEmpty, "QXmlStreamAttributes_isEmpty")
	symaddr(&addr_QXmlStreamAttributes_resize, "QXmlStreamAttributes_resize")
	symaddr(&addr_QXmlStreamAttributes_resize2, "QXmlStreamAttributes_resize2")
	symaddr(&addr_QXmlStreamAttributes_resizeForOverwrite, "QXmlStreamAttributes_resizeForOverwrite")
	symaddr(&addr_QXmlStreamAttributes_capacity, "QXmlStreamAttributes_capacity")
	symaddr(&addr_QXmlStreamAttributes_reserve, "QXmlStreamAttributes_reserve")
	symaddr(&addr_QXmlStreamAttributes_squeeze, "QXmlStreamAttributes_squeeze")
	symaddr(&addr_QXmlStreamAttributes_detach, "QXmlStreamAttributes_detach")
	symaddr(&addr_QXmlStreamAttributes_isDetached, "QXmlStreamAttributes_isDetached")
	symaddr(&addr_QXmlStreamAttributes_isSharedWith, "QXmlStreamAttributes_isSharedWith")
	symaddr(&addr_QXmlStreamAttributes_data, "QXmlStreamAttributes_data")
	symaddr(&addr_QXmlStreamAttributes_data2, "QXmlStreamAttributes_data2")
	symaddr(&addr_QXmlStreamAttributes_constData, "QXmlStreamAttributes_constData")
	symaddr(&addr_QXmlStreamAttributes_clear, "QXmlStreamAttributes_clear")
	symaddr(&addr_QXmlStreamAttributes_at, "QXmlStreamAttributes_at")
	symaddr(&addr_QXmlStreamAttributes_OperatorSubscript, "QXmlStreamAttributes_OperatorSubscript")
	symaddr(&addr_QXmlStreamAttributes_OperatorSubscript2, "QXmlStreamAttributes_OperatorSubscript2")
	symaddr(&addr_QXmlStreamAttributes_prepend, "QXmlStreamAttributes_prepend")
	symaddr(&addr_QXmlStreamAttributes_replace, "QXmlStreamAttributes_replace")
	symaddr(&addr_QXmlStreamAttributes_remove, "QXmlStreamAttributes_remove")
	symaddr(&addr_QXmlStreamAttributes_removeFirst, "QXmlStreamAttributes_removeFirst")
	symaddr(&addr_QXmlStreamAttributes_removeLast, "QXmlStreamAttributes_removeLast")
	symaddr(&addr_QXmlStreamAttributes_takeFirst, "QXmlStreamAttributes_takeFirst")
	symaddr(&addr_QXmlStreamAttributes_takeLast, "QXmlStreamAttributes_takeLast")
	symaddr(&addr_QXmlStreamAttributes_removeAt, "QXmlStreamAttributes_removeAt")
	symaddr(&addr_QXmlStreamAttributes_takeAt, "QXmlStreamAttributes_takeAt")
	symaddr(&addr_QXmlStreamAttributes_move, "QXmlStreamAttributes_move")
	symaddr(&addr_QXmlStreamAttributes_first, "QXmlStreamAttributes_first")
	symaddr(&addr_QXmlStreamAttributes_first2, "QXmlStreamAttributes_first2")
	symaddr(&addr_QXmlStreamAttributes_constFirst, "QXmlStreamAttributes_constFirst")
	symaddr(&addr_QXmlStreamAttributes_last, "QXmlStreamAttributes_last")
	symaddr(&addr_QXmlStreamAttributes_last2, "QXmlStreamAttributes_last2")
	symaddr(&addr_QXmlStreamAttributes_constLast, "QXmlStreamAttributes_constLast")
	symaddr(&addr_QXmlStreamAttributes_startsWith, "QXmlStreamAttributes_startsWith")
	symaddr(&addr_QXmlStreamAttributes_endsWith, "QXmlStreamAttributes_endsWith")
	symaddr(&addr_QXmlStreamAttributes_mid, "QXmlStreamAttributes_mid")
	symaddr(&addr_QXmlStreamAttributes_first3, "QXmlStreamAttributes_first3")
	symaddr(&addr_QXmlStreamAttributes_last3, "QXmlStreamAttributes_last3")
	symaddr(&addr_QXmlStreamAttributes_sliced, "QXmlStreamAttributes_sliced")
	symaddr(&addr_QXmlStreamAttributes_sliced2, "QXmlStreamAttributes_sliced2")
	symaddr(&addr_QXmlStreamAttributes_swapItemsAt, "QXmlStreamAttributes_swapItemsAt")
	symaddr(&addr_QXmlStreamAttributes_push_back, "QXmlStreamAttributes_push_back")
	symaddr(&addr_QXmlStreamAttributes_push_front, "QXmlStreamAttributes_push_front")
	symaddr(&addr_QXmlStreamAttributes_pop_back, "QXmlStreamAttributes_pop_back")
	symaddr(&addr_QXmlStreamAttributes_pop_front, "QXmlStreamAttributes_pop_front")
	symaddr(&addr_QXmlStreamAttributes_empty, "QXmlStreamAttributes_empty")
	symaddr(&addr_QXmlStreamAttributes_front, "QXmlStreamAttributes_front")
	symaddr(&addr_QXmlStreamAttributes_front2, "QXmlStreamAttributes_front2")
	symaddr(&addr_QXmlStreamAttributes_back, "QXmlStreamAttributes_back")
	symaddr(&addr_QXmlStreamAttributes_back2, "QXmlStreamAttributes_back2")
	symaddr(&addr_QXmlStreamAttributes_shrink_to_fit, "QXmlStreamAttributes_shrink_to_fit")
	symaddr(&addr_QXmlStreamAttributes_max_size, "QXmlStreamAttributes_max_size")
	symaddr(&addr_QXmlStreamAttributes_OperatorPlus, "QXmlStreamAttributes_OperatorPlus")
	symaddr(&addr_QXmlStreamAttributes_OperatorPlus2, "QXmlStreamAttributes_OperatorPlus2")
	symaddr(&addr_QXmlStreamAttributes_fromList, "QXmlStreamAttributes_fromList")
	symaddr(&addr_QXmlStreamAttributes_toList, "QXmlStreamAttributes_toList")
	symaddr(&addr_QXmlStreamAttributes_fromVector, "QXmlStreamAttributes_fromVector")
	symaddr(&addr_QXmlStreamAttributes_toVector, "QXmlStreamAttributes_toVector")
	symaddr(&addr_QXmlStreamAttributes_remove2, "QXmlStreamAttributes_remove2")
	symaddr(&addr_QXmlStreamAttributes_mid2, "QXmlStreamAttributes_mid2")
	reg(&pg_QXmlStreamAttributes_delete, "QXmlStreamAttributes_delete")
}

type QXmlStreamAttributes struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamAttributes) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamAttributes) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamAttributes(h uintptr) *QXmlStreamAttributes {
	if h == 0 {
		return nil
	}
	return &QXmlStreamAttributes{h: h}
}

func newownedQXmlStreamAttributes(h uintptr) *QXmlStreamAttributes {
	o := bindQXmlStreamAttributes(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamAttributes).Delete)
	}
	return o
}

func (this *QXmlStreamAttributes) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamAttributes) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamAttributes(h unsafe.Pointer) *QXmlStreamAttributes { return bindQXmlStreamAttributes(uintptr(h)) }
func UnsafeNewQXmlStreamAttributesOwned(h unsafe.Pointer) *QXmlStreamAttributes { return newownedQXmlStreamAttributes(uintptr(h)) }

func (this *QXmlStreamAttributes) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_new)
	__o := newownedQXmlStreamAttributes(__r)
	return __o
}

func (this *QXmlStreamAttributes) Value(namespaceUri string, name string) *QStringView {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_value, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttributes) ValueWithQualifiedName(qualifiedName string) *QStringView {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_value2, this.h, qualifiedName_p, uintptr(len(qualifiedName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttributes) Append(namespaceUri string, name string, value string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QXmlStreamAttributes_append, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QXmlStreamAttributes) Append2(qualifiedName string, value string) {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QXmlStreamAttributes_append2, this.h, qualifiedName_p, uintptr(len(qualifiedName)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
	runtime.KeepAlive(value)
}

func (this *QXmlStreamAttributes) HasAttribute(qualifiedName string) bool {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_hasAttribute, this.h, qualifiedName_p, uintptr(len(qualifiedName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
	return __r != 0
}

func (this *QXmlStreamAttributes) HasAttribute2(namespaceUri string, name string) bool {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_hasAttribute2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	return __r != 0
}

func QXmlStreamAttributes_MaxSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_maxSize)
	__r := int64(__r0)
	return __r
}

func (this *QXmlStreamAttributes) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamAttributes) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamAttributes) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamAttributes) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamAttributes) Resize(size int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Resize2(size int64, c *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamAttributes_resize2, this.h, uintptr(size), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QXmlStreamAttributes) ResizeForOverwrite(size int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_resizeForOverwrite, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamAttributes) Reserve(size int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_reserve, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Squeeze() {
	purego.SyscallN(addr_QXmlStreamAttributes_squeeze, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Detach() {
	purego.SyscallN(addr_QXmlStreamAttributes_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamAttributes) IsSharedWith(other []*QXmlStreamAttribute) bool {
	other_arr := make([]uintptr, len(other))
	for __i, __e := range other {
		other_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var other_p uintptr
	if len(other_arr) > 0 {
		other_p = uintptr(unsafe.Pointer(&other_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_isSharedWith, this.h, other_p, uintptr(len(other)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	runtime.KeepAlive(other_arr)
	return __r != 0
}

func (this *QXmlStreamAttributes) Data() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_data, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Data2() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_data2, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) ConstData() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_constData, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Clear() {
	purego.SyscallN(addr_QXmlStreamAttributes_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) At(i int64) *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_at, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) OperatorSubscript(i int64) *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) OperatorSubscriptWith(i int64) *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_OperatorSubscript2, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Prepend(t *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamAttributes_prepend, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QXmlStreamAttributes) Replace(i int64, t *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamAttributes_replace, this.h, uintptr(i), uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QXmlStreamAttributes) Remove(i int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_remove, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) RemoveFirst() {
	purego.SyscallN(addr_QXmlStreamAttributes_removeFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) RemoveLast() {
	purego.SyscallN(addr_QXmlStreamAttributes_removeLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) TakeFirst() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_takeFirst, this.h)
	runtime.KeepAlive(this)
	return newownedQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) TakeLast() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_takeLast, this.h)
	runtime.KeepAlive(this)
	return newownedQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) RemoveAt(i int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_removeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) TakeAt(i int64) *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_takeAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Move(from int64, to int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_move, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) First() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_first, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) First2() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_first2, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) ConstFirst() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_constFirst, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Last() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_last, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Last2() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_last2, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) ConstLast() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_constLast, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) StartsWith(t *QXmlStreamAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_startsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QXmlStreamAttributes) EndsWith(t *QXmlStreamAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_endsWith, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return __r != 0
}

func (this *QXmlStreamAttributes) Mid(pos int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_mid, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) FirstWith(n int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_first3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) LastWith(n int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_last3, this.h, uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) Sliced(pos int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_sliced, this.h, uintptr(pos), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) Sliced2(pos int64, n int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_sliced2, this.h, uintptr(pos), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) SwapItemsAt(i int64, j int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_swapItemsAt, this.h, uintptr(i), uintptr(j))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Push_back(t *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamAttributes_push_back, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QXmlStreamAttributes) Push_front(t *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamAttributes_push_front, this.h, uintptr(t.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
}

func (this *QXmlStreamAttributes) Pop_back() {
	purego.SyscallN(addr_QXmlStreamAttributes_pop_back, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Pop_front() {
	purego.SyscallN(addr_QXmlStreamAttributes_pop_front, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamAttributes) Front() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_front, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Front2() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_front2, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Back() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_back, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Back2() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_back2, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamAttribute(__r)
}

func (this *QXmlStreamAttributes) Shrink_to_fit() {
	purego.SyscallN(addr_QXmlStreamAttributes_shrink_to_fit, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamAttributes) OperatorPlus(l []*QXmlStreamAttribute) []*QXmlStreamAttribute {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_OperatorPlus, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) OperatorPlusWith(l []*QXmlStreamAttribute) []*QXmlStreamAttribute {
	l_arr := make([]uintptr, len(l))
	for __i, __e := range l {
		l_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var l_p uintptr
	if len(l_arr) > 0 {
		l_p = uintptr(unsafe.Pointer(&l_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_OperatorPlus2, this.h, l_p, uintptr(len(l)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	runtime.KeepAlive(l_arr)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QXmlStreamAttributes_FromList(list []*QXmlStreamAttribute) []*QXmlStreamAttribute {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_fromList, list_p, uintptr(len(list)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) ToList() []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_toList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QXmlStreamAttributes_FromVector(vector []*QXmlStreamAttribute) []*QXmlStreamAttribute {
	vector_arr := make([]uintptr, len(vector))
	for __i, __e := range vector {
		vector_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var vector_p uintptr
	if len(vector_arr) > 0 {
		vector_p = uintptr(unsafe.Pointer(&vector_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_fromVector, vector_p, uintptr(len(vector)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(vector)
	runtime.KeepAlive(vector_arr)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) ToVector() []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_toVector, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) Remove2(i int64, n int64) {
	purego.SyscallN(addr_QXmlStreamAttributes_remove2, this.h, uintptr(i), uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamAttributes) Mid2(pos int64, len_ int64) []*QXmlStreamAttribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttributes_mid2, this.h, uintptr(pos), uintptr(len_), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamAttribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamAttribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamAttributes) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamAttributes_delete(this.h)
	}
	this.UnsafeRelease()
}
