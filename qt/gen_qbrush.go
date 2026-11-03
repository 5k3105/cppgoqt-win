package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBrush_new uintptr
	addr_QBrush_new2 uintptr
	addr_QBrush_new3 uintptr
	addr_QBrush_new4 uintptr
	addr_QBrush_new5 uintptr
	addr_QBrush_new6 uintptr
	addr_QBrush_new7 uintptr
	addr_QBrush_new8 uintptr
	addr_QBrush_new9 uintptr
	addr_QBrush_new10 uintptr
	addr_QBrush_new11 uintptr
	addr_QBrush_new12 uintptr
	addr_QBrush_OperatorAssign uintptr
	addr_QBrush_swap uintptr
	addr_QBrush_OperatorAssign2 uintptr
	addr_QBrush_OperatorAssign3 uintptr
	addr_QBrush_OperatorAssign4 uintptr
	addr_QBrush_ToQVariant uintptr
	addr_QBrush_style uintptr
	addr_QBrush_setStyle uintptr
	addr_QBrush_transform uintptr
	addr_QBrush_setTransform uintptr
	addr_QBrush_texture uintptr
	addr_QBrush_setTexture uintptr
	addr_QBrush_textureImage uintptr
	addr_QBrush_setTextureImage uintptr
	addr_QBrush_color uintptr
	addr_QBrush_setColor uintptr
	addr_QBrush_setColor2 uintptr
	addr_QBrush_gradient uintptr
	addr_QBrush_isOpaque uintptr
	addr_QBrush_OperatorEqual uintptr
	addr_QBrush_OperatorNotEqual uintptr
	addr_QBrush_isDetached uintptr
	addr_QBrush_OperatorEqual2 uintptr
	addr_QBrush_OperatorNotEqual2 uintptr
	addr_QBrush_toQVariant uintptr
	pg_QBrush_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBrush_new, "QBrush_new")
	symaddr(&addr_QBrush_new2, "QBrush_new2")
	symaddr(&addr_QBrush_new3, "QBrush_new3")
	symaddr(&addr_QBrush_new4, "QBrush_new4")
	symaddr(&addr_QBrush_new5, "QBrush_new5")
	symaddr(&addr_QBrush_new6, "QBrush_new6")
	symaddr(&addr_QBrush_new7, "QBrush_new7")
	symaddr(&addr_QBrush_new8, "QBrush_new8")
	symaddr(&addr_QBrush_new9, "QBrush_new9")
	symaddr(&addr_QBrush_new10, "QBrush_new10")
	symaddr(&addr_QBrush_new11, "QBrush_new11")
	symaddr(&addr_QBrush_new12, "QBrush_new12")
	symaddr(&addr_QBrush_OperatorAssign, "QBrush_OperatorAssign")
	symaddr(&addr_QBrush_swap, "QBrush_swap")
	symaddr(&addr_QBrush_OperatorAssign2, "QBrush_OperatorAssign2")
	symaddr(&addr_QBrush_OperatorAssign3, "QBrush_OperatorAssign3")
	symaddr(&addr_QBrush_OperatorAssign4, "QBrush_OperatorAssign4")
	symaddr(&addr_QBrush_ToQVariant, "QBrush_ToQVariant")
	symaddr(&addr_QBrush_style, "QBrush_style")
	symaddr(&addr_QBrush_setStyle, "QBrush_setStyle")
	symaddr(&addr_QBrush_transform, "QBrush_transform")
	symaddr(&addr_QBrush_setTransform, "QBrush_setTransform")
	symaddr(&addr_QBrush_texture, "QBrush_texture")
	symaddr(&addr_QBrush_setTexture, "QBrush_setTexture")
	symaddr(&addr_QBrush_textureImage, "QBrush_textureImage")
	symaddr(&addr_QBrush_setTextureImage, "QBrush_setTextureImage")
	symaddr(&addr_QBrush_color, "QBrush_color")
	symaddr(&addr_QBrush_setColor, "QBrush_setColor")
	symaddr(&addr_QBrush_setColor2, "QBrush_setColor2")
	symaddr(&addr_QBrush_gradient, "QBrush_gradient")
	symaddr(&addr_QBrush_isOpaque, "QBrush_isOpaque")
	symaddr(&addr_QBrush_OperatorEqual, "QBrush_OperatorEqual")
	symaddr(&addr_QBrush_OperatorNotEqual, "QBrush_OperatorNotEqual")
	symaddr(&addr_QBrush_isDetached, "QBrush_isDetached")
	symaddr(&addr_QBrush_OperatorEqual2, "QBrush_OperatorEqual2")
	symaddr(&addr_QBrush_OperatorNotEqual2, "QBrush_OperatorNotEqual2")
	symaddr(&addr_QBrush_toQVariant, "QBrush_toQVariant")
	reg(&pg_QBrush_delete, "QBrush_delete")
}

type QBrush struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBrush) UnsafeGoID() int64     { return this.vid }
func (this *QBrush) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBrush(h uintptr) *QBrush {
	if h == 0 {
		return nil
	}
	return &QBrush{h: h}
}

func newownedQBrush(h uintptr) *QBrush {
	o := bindQBrush(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBrush).Delete)
	}
	return o
}

func (this *QBrush) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBrush) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBrush(h unsafe.Pointer) *QBrush { return bindQBrush(uintptr(h)) }
func UnsafeNewQBrushOwned(h unsafe.Pointer) *QBrush { return newownedQBrush(uintptr(h)) }

func (this *QBrush) UnsafeRelease() { this.h = 0 }

func NewQBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new)
	__o := newownedQBrush(__r)
	return __o
}

func NewQBrush2(bs BrushStyle) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new2, uintptr(bs))
	__o := newownedQBrush(__r)
	return __o
}

func NewQBrush3(color *QColor) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new3, uintptr(color.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(color)
	return __o
}

func NewQBrush4(color GlobalColor) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new4, uintptr(color))
	__o := newownedQBrush(__r)
	return __o
}

func NewQBrush5(color *QColor, pixmap *QPixmap) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new5, uintptr(color.UnsafePointer()), uintptr(pixmap.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(color)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQBrush6(color GlobalColor, pixmap *QPixmap) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new6, uintptr(color), uintptr(pixmap.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQBrush7(pixmap *QPixmap) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new7, uintptr(pixmap.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQBrush8(image *QImage) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new8, uintptr(image.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(image)
	return __o
}

func NewQBrush9(brush *QBrush) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new9, uintptr(brush.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(brush)
	return __o
}

func NewQBrush10(gradient *QGradient) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new10, uintptr(gradient.UnsafePointer()))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(gradient)
	return __o
}

func NewQBrush11(color *QColor, bs BrushStyle) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new11, uintptr(color.UnsafePointer()), uintptr(bs))
	__o := newownedQBrush(__r)
	runtime.KeepAlive(color)
	return __o
}

func NewQBrush12(color GlobalColor, bs BrushStyle) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_new12, uintptr(color), uintptr(bs))
	__o := newownedQBrush(__r)
	return __o
}

func (this *QBrush) OperatorAssign(brush *QBrush) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_OperatorAssign, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
	return bindQBrush(__r)
}

func (this *QBrush) Swap(other *QBrush) {
	purego.SyscallN(addr_QBrush_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QBrush) OperatorAssignWithStyle(style BrushStyle) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_OperatorAssign2, this.h, uintptr(style))
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QBrush) OperatorAssignWithColor(color *QColor) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_OperatorAssign3, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
	return bindQBrush(__r)
}

func (this *QBrush) OperatorAssign2(color GlobalColor) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QBrush_OperatorAssign4, this.h, uintptr(color))
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QBrush) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QBrush_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QBrush) Style() BrushStyle {
	__r0, _, _ := purego.SyscallN(addr_QBrush_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return BrushStyle(__r)
}

func (this *QBrush) SetStyle(arg1 BrushStyle) {
	purego.SyscallN(addr_QBrush_setStyle, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QBrush) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QBrush_transform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QBrush) SetTransform(arg1 *QTransform) {
	purego.SyscallN(addr_QBrush_setTransform, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBrush) Texture() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QBrush_texture, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QBrush) SetTexture(pixmap *QPixmap) {
	purego.SyscallN(addr_QBrush_setTexture, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QBrush) TextureImage() *QImage {
	__r, _, _ := purego.SyscallN(addr_QBrush_textureImage, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QBrush) SetTextureImage(image *QImage) {
	purego.SyscallN(addr_QBrush_setTextureImage, this.h, uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QBrush) Color() *QColor {
	__r, _, _ := purego.SyscallN(addr_QBrush_color, this.h)
	runtime.KeepAlive(this)
	return bindQColor(__r)
}

func (this *QBrush) SetColor(color *QColor) {
	purego.SyscallN(addr_QBrush_setColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QBrush) SetColorWithColor(color GlobalColor) {
	purego.SyscallN(addr_QBrush_setColor2, this.h, uintptr(color))
	runtime.KeepAlive(this)
}

func (this *QBrush) Gradient() *QGradient {
	__r, _, _ := purego.SyscallN(addr_QBrush_gradient, this.h)
	runtime.KeepAlive(this)
	return bindQGradient(__r)
}

func (this *QBrush) IsOpaque() bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_isOpaque, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBrush) OperatorEqual(b *QBrush) bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_OperatorEqual, this.h, uintptr(b.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return __r != 0
}

func (this *QBrush) OperatorNotEqual(b *QBrush) bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_OperatorNotEqual, this.h, uintptr(b.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(b)
	return __r != 0
}

func (this *QBrush) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBrush) OperatorEqualWithRhs(rhs *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QBrush) OperatorNotEqualWithRhs(rhs *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QBrush_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QBrush) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QBrush_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QBrush) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBrush_delete(this.h)
	}
	this.UnsafeRelease()
}
