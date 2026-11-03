package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextItem_new uintptr
	addr_QTextItem_new2 uintptr
	pg_QTextItem_descent func(uintptr) float64
	pg_QTextItem_ascent func(uintptr) float64
	pg_QTextItem_width func(uintptr) float64
	addr_QTextItem_renderFlags uintptr
	addr_QTextItem_text uintptr
	addr_QTextItem_font uintptr
	pg_QTextItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextItem_new, "QTextItem_new")
	symaddr(&addr_QTextItem_new2, "QTextItem_new2")
	reg(&pg_QTextItem_descent, "QTextItem_descent")
	reg(&pg_QTextItem_ascent, "QTextItem_ascent")
	reg(&pg_QTextItem_width, "QTextItem_width")
	symaddr(&addr_QTextItem_renderFlags, "QTextItem_renderFlags")
	symaddr(&addr_QTextItem_text, "QTextItem_text")
	symaddr(&addr_QTextItem_font, "QTextItem_font")
	reg(&pg_QTextItem_delete, "QTextItem_delete")
}

type QTextItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextItem) UnsafeGoID() int64     { return this.vid }
func (this *QTextItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextItem(h uintptr) *QTextItem {
	if h == 0 {
		return nil
	}
	return &QTextItem{h: h}
}

func newownedQTextItem(h uintptr) *QTextItem {
	o := bindQTextItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextItem).Delete)
	}
	return o
}

func (this *QTextItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextItem(h unsafe.Pointer) *QTextItem { return bindQTextItem(uintptr(h)) }
func UnsafeNewQTextItemOwned(h unsafe.Pointer) *QTextItem { return newownedQTextItem(uintptr(h)) }

func (this *QTextItem) UnsafeRelease() { this.h = 0 }

func NewQTextItem() *QTextItem {
	__r, _, _ := purego.SyscallN(addr_QTextItem_new)
	__o := newownedQTextItem(__r)
	return __o
}

func NewQTextItem2(arg1 *QTextItem) *QTextItem {
	__r, _, _ := purego.SyscallN(addr_QTextItem_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextItem(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextItem) Descent() float64 {
	__r := pg_QTextItem_descent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextItem) Ascent() float64 {
	__r := pg_QTextItem_ascent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextItem) Width() float64 {
	__r := pg_QTextItem_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextItem) RenderFlags() QTextItem__RenderFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextItem_renderFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextItem__RenderFlag(__r)
}

func (this *QTextItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextItem_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTextItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTextItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextItem_delete(this.h)
	}
	this.UnsafeRelease()
}
