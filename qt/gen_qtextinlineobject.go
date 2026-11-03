package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextInlineObject_new uintptr
	addr_QTextInlineObject_isValid uintptr
	addr_QTextInlineObject_rect uintptr
	pg_QTextInlineObject_width func(uintptr) float64
	pg_QTextInlineObject_ascent func(uintptr) float64
	pg_QTextInlineObject_descent func(uintptr) float64
	pg_QTextInlineObject_height func(uintptr) float64
	addr_QTextInlineObject_textDirection uintptr
	pg_QTextInlineObject_setWidth func(uintptr, float64)
	pg_QTextInlineObject_setAscent func(uintptr, float64)
	pg_QTextInlineObject_setDescent func(uintptr, float64)
	addr_QTextInlineObject_textPosition uintptr
	addr_QTextInlineObject_formatIndex uintptr
	addr_QTextInlineObject_format uintptr
	pg_QTextInlineObject_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextInlineObject_new, "QTextInlineObject_new")
	symaddr(&addr_QTextInlineObject_isValid, "QTextInlineObject_isValid")
	symaddr(&addr_QTextInlineObject_rect, "QTextInlineObject_rect")
	reg(&pg_QTextInlineObject_width, "QTextInlineObject_width")
	reg(&pg_QTextInlineObject_ascent, "QTextInlineObject_ascent")
	reg(&pg_QTextInlineObject_descent, "QTextInlineObject_descent")
	reg(&pg_QTextInlineObject_height, "QTextInlineObject_height")
	symaddr(&addr_QTextInlineObject_textDirection, "QTextInlineObject_textDirection")
	reg(&pg_QTextInlineObject_setWidth, "QTextInlineObject_setWidth")
	reg(&pg_QTextInlineObject_setAscent, "QTextInlineObject_setAscent")
	reg(&pg_QTextInlineObject_setDescent, "QTextInlineObject_setDescent")
	symaddr(&addr_QTextInlineObject_textPosition, "QTextInlineObject_textPosition")
	symaddr(&addr_QTextInlineObject_formatIndex, "QTextInlineObject_formatIndex")
	symaddr(&addr_QTextInlineObject_format, "QTextInlineObject_format")
	reg(&pg_QTextInlineObject_delete, "QTextInlineObject_delete")
}

type QTextInlineObject struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextInlineObject) UnsafeGoID() int64     { return this.vid }
func (this *QTextInlineObject) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextInlineObject(h uintptr) *QTextInlineObject {
	if h == 0 {
		return nil
	}
	return &QTextInlineObject{h: h}
}

func newownedQTextInlineObject(h uintptr) *QTextInlineObject {
	o := bindQTextInlineObject(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextInlineObject).Delete)
	}
	return o
}

func (this *QTextInlineObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextInlineObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextInlineObject(h unsafe.Pointer) *QTextInlineObject { return bindQTextInlineObject(uintptr(h)) }
func UnsafeNewQTextInlineObjectOwned(h unsafe.Pointer) *QTextInlineObject { return newownedQTextInlineObject(uintptr(h)) }

func (this *QTextInlineObject) UnsafeRelease() { this.h = 0 }

func NewQTextInlineObject() *QTextInlineObject {
	__r, _, _ := purego.SyscallN(addr_QTextInlineObject_new)
	__o := newownedQTextInlineObject(__r)
	return __o
}

func (this *QTextInlineObject) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextInlineObject_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextInlineObject) Rect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QTextInlineObject_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QTextInlineObject) Width() float64 {
	__r := pg_QTextInlineObject_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextInlineObject) Ascent() float64 {
	__r := pg_QTextInlineObject_ascent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextInlineObject) Descent() float64 {
	__r := pg_QTextInlineObject_descent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextInlineObject) Height() float64 {
	__r := pg_QTextInlineObject_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextInlineObject) TextDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QTextInlineObject_textDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QTextInlineObject) SetWidth(w float64) {
	pg_QTextInlineObject_setWidth(this.h, w)
	runtime.KeepAlive(this)
}

func (this *QTextInlineObject) SetAscent(a float64) {
	pg_QTextInlineObject_setAscent(this.h, a)
	runtime.KeepAlive(this)
}

func (this *QTextInlineObject) SetDescent(d float64) {
	pg_QTextInlineObject_setDescent(this.h, d)
	runtime.KeepAlive(this)
}

func (this *QTextInlineObject) TextPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextInlineObject_textPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextInlineObject) FormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextInlineObject_formatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextInlineObject) Format() *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextInlineObject_format, this.h)
	runtime.KeepAlive(this)
	return newownedQTextFormat(__r)
}

func (this *QTextInlineObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextInlineObject_delete(this.h)
	}
	this.UnsafeRelease()
}
