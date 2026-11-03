package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextFormat_new uintptr
	addr_QTextFormat_new2 uintptr
	addr_QTextFormat_new3 uintptr
	addr_QTextFormat_OperatorAssign uintptr
	addr_QTextFormat_swap uintptr
	addr_QTextFormat_merge uintptr
	addr_QTextFormat_isValid uintptr
	addr_QTextFormat_isEmpty uintptr
	addr_QTextFormat_type uintptr
	addr_QTextFormat_objectIndex uintptr
	addr_QTextFormat_setObjectIndex uintptr
	addr_QTextFormat_property uintptr
	addr_QTextFormat_setProperty uintptr
	addr_QTextFormat_clearProperty uintptr
	addr_QTextFormat_hasProperty uintptr
	addr_QTextFormat_boolProperty uintptr
	addr_QTextFormat_intProperty uintptr
	pg_QTextFormat_doubleProperty func(uintptr, int32) float64
	addr_QTextFormat_stringProperty uintptr
	addr_QTextFormat_colorProperty uintptr
	addr_QTextFormat_penProperty uintptr
	addr_QTextFormat_brushProperty uintptr
	addr_QTextFormat_lengthProperty uintptr
	addr_QTextFormat_lengthVectorProperty uintptr
	addr_QTextFormat_setProperty2 uintptr
	addr_QTextFormat_properties uintptr
	addr_QTextFormat_propertyCount uintptr
	addr_QTextFormat_setObjectType uintptr
	addr_QTextFormat_objectType uintptr
	addr_QTextFormat_isCharFormat uintptr
	addr_QTextFormat_isBlockFormat uintptr
	addr_QTextFormat_isListFormat uintptr
	addr_QTextFormat_isFrameFormat uintptr
	addr_QTextFormat_isImageFormat uintptr
	addr_QTextFormat_isTableFormat uintptr
	addr_QTextFormat_isTableCellFormat uintptr
	addr_QTextFormat_toBlockFormat uintptr
	addr_QTextFormat_toCharFormat uintptr
	addr_QTextFormat_toListFormat uintptr
	addr_QTextFormat_toTableFormat uintptr
	addr_QTextFormat_toFrameFormat uintptr
	addr_QTextFormat_toImageFormat uintptr
	addr_QTextFormat_toTableCellFormat uintptr
	addr_QTextFormat_OperatorEqual uintptr
	addr_QTextFormat_OperatorNotEqual uintptr
	addr_QTextFormat_ToQVariant uintptr
	addr_QTextFormat_setLayoutDirection uintptr
	addr_QTextFormat_layoutDirection uintptr
	addr_QTextFormat_setBackground uintptr
	addr_QTextFormat_background uintptr
	addr_QTextFormat_clearBackground uintptr
	addr_QTextFormat_setForeground uintptr
	addr_QTextFormat_foreground uintptr
	addr_QTextFormat_clearForeground uintptr
	addr_QTextFormat_toQVariant uintptr
	pg_QTextFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextFormat_new, "QTextFormat_new")
	symaddr(&addr_QTextFormat_new2, "QTextFormat_new2")
	symaddr(&addr_QTextFormat_new3, "QTextFormat_new3")
	symaddr(&addr_QTextFormat_OperatorAssign, "QTextFormat_OperatorAssign")
	symaddr(&addr_QTextFormat_swap, "QTextFormat_swap")
	symaddr(&addr_QTextFormat_merge, "QTextFormat_merge")
	symaddr(&addr_QTextFormat_isValid, "QTextFormat_isValid")
	symaddr(&addr_QTextFormat_isEmpty, "QTextFormat_isEmpty")
	symaddr(&addr_QTextFormat_type, "QTextFormat_type")
	symaddr(&addr_QTextFormat_objectIndex, "QTextFormat_objectIndex")
	symaddr(&addr_QTextFormat_setObjectIndex, "QTextFormat_setObjectIndex")
	symaddr(&addr_QTextFormat_property, "QTextFormat_property")
	symaddr(&addr_QTextFormat_setProperty, "QTextFormat_setProperty")
	symaddr(&addr_QTextFormat_clearProperty, "QTextFormat_clearProperty")
	symaddr(&addr_QTextFormat_hasProperty, "QTextFormat_hasProperty")
	symaddr(&addr_QTextFormat_boolProperty, "QTextFormat_boolProperty")
	symaddr(&addr_QTextFormat_intProperty, "QTextFormat_intProperty")
	reg(&pg_QTextFormat_doubleProperty, "QTextFormat_doubleProperty")
	symaddr(&addr_QTextFormat_stringProperty, "QTextFormat_stringProperty")
	symaddr(&addr_QTextFormat_colorProperty, "QTextFormat_colorProperty")
	symaddr(&addr_QTextFormat_penProperty, "QTextFormat_penProperty")
	symaddr(&addr_QTextFormat_brushProperty, "QTextFormat_brushProperty")
	symaddr(&addr_QTextFormat_lengthProperty, "QTextFormat_lengthProperty")
	symaddr(&addr_QTextFormat_lengthVectorProperty, "QTextFormat_lengthVectorProperty")
	symaddr(&addr_QTextFormat_setProperty2, "QTextFormat_setProperty2")
	symaddr(&addr_QTextFormat_properties, "QTextFormat_properties")
	symaddr(&addr_QTextFormat_propertyCount, "QTextFormat_propertyCount")
	symaddr(&addr_QTextFormat_setObjectType, "QTextFormat_setObjectType")
	symaddr(&addr_QTextFormat_objectType, "QTextFormat_objectType")
	symaddr(&addr_QTextFormat_isCharFormat, "QTextFormat_isCharFormat")
	symaddr(&addr_QTextFormat_isBlockFormat, "QTextFormat_isBlockFormat")
	symaddr(&addr_QTextFormat_isListFormat, "QTextFormat_isListFormat")
	symaddr(&addr_QTextFormat_isFrameFormat, "QTextFormat_isFrameFormat")
	symaddr(&addr_QTextFormat_isImageFormat, "QTextFormat_isImageFormat")
	symaddr(&addr_QTextFormat_isTableFormat, "QTextFormat_isTableFormat")
	symaddr(&addr_QTextFormat_isTableCellFormat, "QTextFormat_isTableCellFormat")
	symaddr(&addr_QTextFormat_toBlockFormat, "QTextFormat_toBlockFormat")
	symaddr(&addr_QTextFormat_toCharFormat, "QTextFormat_toCharFormat")
	symaddr(&addr_QTextFormat_toListFormat, "QTextFormat_toListFormat")
	symaddr(&addr_QTextFormat_toTableFormat, "QTextFormat_toTableFormat")
	symaddr(&addr_QTextFormat_toFrameFormat, "QTextFormat_toFrameFormat")
	symaddr(&addr_QTextFormat_toImageFormat, "QTextFormat_toImageFormat")
	symaddr(&addr_QTextFormat_toTableCellFormat, "QTextFormat_toTableCellFormat")
	symaddr(&addr_QTextFormat_OperatorEqual, "QTextFormat_OperatorEqual")
	symaddr(&addr_QTextFormat_OperatorNotEqual, "QTextFormat_OperatorNotEqual")
	symaddr(&addr_QTextFormat_ToQVariant, "QTextFormat_ToQVariant")
	symaddr(&addr_QTextFormat_setLayoutDirection, "QTextFormat_setLayoutDirection")
	symaddr(&addr_QTextFormat_layoutDirection, "QTextFormat_layoutDirection")
	symaddr(&addr_QTextFormat_setBackground, "QTextFormat_setBackground")
	symaddr(&addr_QTextFormat_background, "QTextFormat_background")
	symaddr(&addr_QTextFormat_clearBackground, "QTextFormat_clearBackground")
	symaddr(&addr_QTextFormat_setForeground, "QTextFormat_setForeground")
	symaddr(&addr_QTextFormat_foreground, "QTextFormat_foreground")
	symaddr(&addr_QTextFormat_clearForeground, "QTextFormat_clearForeground")
	symaddr(&addr_QTextFormat_toQVariant, "QTextFormat_toQVariant")
	reg(&pg_QTextFormat_delete, "QTextFormat_delete")
}

type QTextFormat struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextFormat) UnsafeGoID() int64     { return this.vid }
func (this *QTextFormat) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextFormat(h uintptr) *QTextFormat {
	if h == 0 {
		return nil
	}
	return &QTextFormat{h: h}
}

func newownedQTextFormat(h uintptr) *QTextFormat {
	o := bindQTextFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextFormat).Delete)
	}
	return o
}

func (this *QTextFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextFormat(h unsafe.Pointer) *QTextFormat { return bindQTextFormat(uintptr(h)) }
func UnsafeNewQTextFormatOwned(h unsafe.Pointer) *QTextFormat { return newownedQTextFormat(uintptr(h)) }

func (this *QTextFormat) UnsafeRelease() { this.h = 0 }

func NewQTextFormat() *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_new)
	__o := newownedQTextFormat(__r)
	return __o
}

func NewQTextFormat2(type_ int) *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_new2, uintptr(type_))
	__o := newownedQTextFormat(__r)
	return __o
}

func NewQTextFormat3(rhs *QTextFormat) *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_new3, uintptr(rhs.UnsafePointer()))
	__o := newownedQTextFormat(__r)
	runtime.KeepAlive(rhs)
	return __o
}

func (this *QTextFormat) OperatorAssign(rhs *QTextFormat) *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_OperatorAssign, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return bindQTextFormat(__r)
}

func (this *QTextFormat) Swap(other *QTextFormat) {
	purego.SyscallN(addr_QTextFormat_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTextFormat) Merge(other *QTextFormat) {
	purego.SyscallN(addr_QTextFormat_merge, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTextFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFormat) ObjectIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_objectIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFormat) SetObjectIndex(object int) {
	purego.SyscallN(addr_QTextFormat_setObjectIndex, this.h, uintptr(object))
	runtime.KeepAlive(this)
}

func (this *QTextFormat) Property(propertyId int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_property, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextFormat) SetProperty(propertyId int, value *QVariant) {
	purego.SyscallN(addr_QTextFormat_setProperty, this.h, uintptr(propertyId), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTextFormat) ClearProperty(propertyId int) {
	purego.SyscallN(addr_QTextFormat_clearProperty, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
}

func (this *QTextFormat) HasProperty(propertyId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_hasProperty, this.h, uintptr(propertyId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) BoolProperty(propertyId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_boolProperty, this.h, uintptr(propertyId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IntProperty(propertyId int) int {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_intProperty, this.h, uintptr(propertyId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFormat) DoubleProperty(propertyId int) float64 {
	__r := pg_QTextFormat_doubleProperty(this.h, int32(propertyId))
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextFormat) StringProperty(propertyId int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFormat_stringProperty, this.h, uintptr(propertyId), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextFormat) ColorProperty(propertyId int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_colorProperty, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QTextFormat) PenProperty(propertyId int) *QPen {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_penProperty, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
	return newownedQPen(__r)
}

func (this *QTextFormat) BrushProperty(propertyId int) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_brushProperty, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextFormat) LengthProperty(propertyId int) *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_lengthProperty, this.h, uintptr(propertyId))
	runtime.KeepAlive(this)
	return newownedQTextLength(__r)
}

func (this *QTextFormat) LengthVectorProperty(propertyId int) []*QTextLength {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFormat_lengthVectorProperty, this.h, uintptr(propertyId), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextLength, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextLength(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextFormat) SetProperty2(propertyId int, lengths []*QTextLength) {
	lengths_arr := make([]uintptr, len(lengths))
	for __i, __e := range lengths {
		lengths_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var lengths_p uintptr
	if len(lengths_arr) > 0 {
		lengths_p = uintptr(unsafe.Pointer(&lengths_arr[0]))
	}
	purego.SyscallN(addr_QTextFormat_setProperty2, this.h, uintptr(propertyId), lengths_p, uintptr(len(lengths)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lengths)
	runtime.KeepAlive(lengths_arr)
}

func (this *QTextFormat) Properties() map[int]*QVariant {
	var __vp uintptr
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFormat_properties, this.h, uintptr(unsafe.Pointer(&__vp)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__ks := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	__vs := unsafe.Slice((*uintptr)(unsafe.Pointer(__vp)), int(__n))
	__out := make(map[int]*QVariant, int(__n))
	for __i := 0; __i < int(__n); __i++ {
		__out[int(__ks[__i])] = newownedQVariant(__vs[__i])
	}
	pgFree(__r)
	pgFree(__vp)
	return __out
}

func (this *QTextFormat) PropertyCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_propertyCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFormat) SetObjectType(type_ int) {
	purego.SyscallN(addr_QTextFormat_setObjectType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QTextFormat) ObjectType() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_objectType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFormat) IsCharFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isCharFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsBlockFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isBlockFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsListFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isListFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsFrameFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isFrameFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsImageFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isImageFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsTableFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isTableFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) IsTableCellFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_isTableCellFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextFormat) ToBlockFormat() *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toBlockFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlockFormat(__r)
}

func (this *QTextFormat) ToCharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toCharFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextFormat) ToListFormat() *QTextListFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toListFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextListFormat(__r)
}

func (this *QTextFormat) ToTableFormat() *QTextTableFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toTableFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextTableFormat(__r)
}

func (this *QTextFormat) ToFrameFormat() *QTextFrameFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toFrameFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextFrameFormat(__r)
}

func (this *QTextFormat) ToImageFormat() *QTextImageFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toImageFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextImageFormat(__r)
}

func (this *QTextFormat) ToTableCellFormat() *QTextTableCellFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toTableCellFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextTableCellFormat(__r)
}

func (this *QTextFormat) OperatorEqual(rhs *QTextFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextFormat) OperatorNotEqual(rhs *QTextFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextFormat) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextFormat) SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QTextFormat_setLayoutDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QTextFormat) LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QTextFormat_layoutDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QTextFormat) SetBackground(brush *QBrush) {
	purego.SyscallN(addr_QTextFormat_setBackground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextFormat) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_background, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextFormat) ClearBackground() {
	purego.SyscallN(addr_QTextFormat_clearBackground, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextFormat) SetForeground(brush *QBrush) {
	purego.SyscallN(addr_QTextFormat_setForeground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTextFormat) Foreground() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_foreground, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTextFormat) ClearForeground() {
	purego.SyscallN(addr_QTextFormat_clearForeground, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextFormat) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextFormat_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
