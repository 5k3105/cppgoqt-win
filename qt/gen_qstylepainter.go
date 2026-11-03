package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStylePainter_base func(uintptr) uintptr
	addr_QStylePainter_new uintptr
	addr_QStylePainter_new2 uintptr
	addr_QStylePainter_new3 uintptr
	addr_QStylePainter_begin uintptr
	addr_QStylePainter_begin2 uintptr
	addr_QStylePainter_drawPrimitive uintptr
	addr_QStylePainter_drawControl uintptr
	addr_QStylePainter_drawComplexControl uintptr
	addr_QStylePainter_drawItemText uintptr
	addr_QStylePainter_drawItemPixmap uintptr
	addr_QStylePainter_style uintptr
	addr_QStylePainter_drawItemText2 uintptr
	pg_QStylePainter_delete func(uintptr)
)

func init() {
	reg(&pg_QStylePainter_base, "QStylePainter_base")
	symaddr(&addr_QStylePainter_new, "QStylePainter_new")
	symaddr(&addr_QStylePainter_new2, "QStylePainter_new2")
	symaddr(&addr_QStylePainter_new3, "QStylePainter_new3")
	symaddr(&addr_QStylePainter_begin, "QStylePainter_begin")
	symaddr(&addr_QStylePainter_begin2, "QStylePainter_begin2")
	symaddr(&addr_QStylePainter_drawPrimitive, "QStylePainter_drawPrimitive")
	symaddr(&addr_QStylePainter_drawControl, "QStylePainter_drawControl")
	symaddr(&addr_QStylePainter_drawComplexControl, "QStylePainter_drawComplexControl")
	symaddr(&addr_QStylePainter_drawItemText, "QStylePainter_drawItemText")
	symaddr(&addr_QStylePainter_drawItemPixmap, "QStylePainter_drawItemPixmap")
	symaddr(&addr_QStylePainter_style, "QStylePainter_style")
	symaddr(&addr_QStylePainter_drawItemText2, "QStylePainter_drawItemText2")
	reg(&pg_QStylePainter_delete, "QStylePainter_delete")
}

type QStylePainter struct {
	h uintptr
	*QPainter
}

func bindQStylePainter(h uintptr) *QStylePainter {
	if h == 0 {
		return nil
	}
	return &QStylePainter{h: h, QPainter: bindQPainter(pg_QStylePainter_base(h))}
}

func newownedQStylePainter(h uintptr) *QStylePainter {
	o := bindQStylePainter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStylePainter).Delete)
	}
	return o
}

func (this *QStylePainter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStylePainter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStylePainter(h unsafe.Pointer) *QStylePainter { return bindQStylePainter(uintptr(h)) }
func UnsafeNewQStylePainterOwned(h unsafe.Pointer) *QStylePainter { return newownedQStylePainter(uintptr(h)) }

func (this *QStylePainter) UnsafeRelease() {
	this.h = 0
	if this.QPainter != nil {
		this.QPainter.UnsafeRelease()
	}
}

func NewQStylePainter(w *QWidget) *QStylePainter {
	__r, _, _ := purego.SyscallN(addr_QStylePainter_new, uintptr(w.UnsafePointer()))
	__o := newownedQStylePainter(__r)
	runtime.KeepAlive(w)
	return __o
}

func NewQStylePainter2() *QStylePainter {
	__r, _, _ := purego.SyscallN(addr_QStylePainter_new2)
	__o := newownedQStylePainter(__r)
	return __o
}

func NewQStylePainter3(pd *QPaintDevice, w *QWidget) *QStylePainter {
	__r, _, _ := purego.SyscallN(addr_QStylePainter_new3, uintptr(pd.UnsafePointer()), uintptr(w.UnsafePointer()))
	__o := newownedQStylePainter(__r)
	runtime.KeepAlive(pd)
	runtime.KeepAlive(w)
	return __o
}

func (this *QStylePainter) Begin(w *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QStylePainter_begin, this.h, uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return __r != 0
}

func (this *QStylePainter) Begin2(pd *QPaintDevice, w *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QStylePainter_begin2, this.h, uintptr(pd.UnsafePointer()), uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pd)
	runtime.KeepAlive(w)
	return __r != 0
}

func (this *QStylePainter) DrawPrimitive(pe QStyle__PrimitiveElement, opt *QStyleOption) {
	purego.SyscallN(addr_QStylePainter_drawPrimitive, this.h, uintptr(pe), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
}

func (this *QStylePainter) DrawControl(ce QStyle__ControlElement, opt *QStyleOption) {
	purego.SyscallN(addr_QStylePainter_drawControl, this.h, uintptr(ce), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
}

func (this *QStylePainter) DrawComplexControl(cc QStyle__ComplexControl, opt *QStyleOptionComplex) {
	purego.SyscallN(addr_QStylePainter_drawComplexControl, this.h, uintptr(cc), uintptr(opt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(opt)
}

func (this *QStylePainter) DrawItemText(r *QRect, flags int, pal *QPalette, enabled bool, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStylePainter_drawItemText, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QStylePainter) DrawItemPixmap(r *QRect, flags int, pixmap *QPixmap) {
	purego.SyscallN(addr_QStylePainter_drawItemPixmap, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pixmap)
}

func (this *QStylePainter) Style() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QStylePainter_style, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func (this *QStylePainter) DrawItemText2(r *QRect, flags int, pal *QPalette, enabled bool, text string, textRole QPalette__ColorRole) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStylePainter_drawItemText2, this.h, uintptr(r.UnsafePointer()), uintptr(flags), uintptr(pal.UnsafePointer()), uintptr(b2i(enabled)), text_p, uintptr(len(text)), uintptr(textRole))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pal)
	runtime.KeepAlive(text)
}

func (this *QStylePainter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStylePainter_delete(this.h)
	}
	this.UnsafeRelease()
}
