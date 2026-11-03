package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QToolTip_showText uintptr
	addr_QToolTip_hideText uintptr
	addr_QToolTip_isVisible uintptr
	addr_QToolTip_text uintptr
	addr_QToolTip_palette uintptr
	addr_QToolTip_setPalette uintptr
	addr_QToolTip_font uintptr
	addr_QToolTip_setFont uintptr
	addr_QToolTip_showText2 uintptr
	addr_QToolTip_showText3 uintptr
	addr_QToolTip_showText4 uintptr
	pg_QToolTip_delete func(uintptr)
)

func init() {
	symaddr(&addr_QToolTip_showText, "QToolTip_showText")
	symaddr(&addr_QToolTip_hideText, "QToolTip_hideText")
	symaddr(&addr_QToolTip_isVisible, "QToolTip_isVisible")
	symaddr(&addr_QToolTip_text, "QToolTip_text")
	symaddr(&addr_QToolTip_palette, "QToolTip_palette")
	symaddr(&addr_QToolTip_setPalette, "QToolTip_setPalette")
	symaddr(&addr_QToolTip_font, "QToolTip_font")
	symaddr(&addr_QToolTip_setFont, "QToolTip_setFont")
	symaddr(&addr_QToolTip_showText2, "QToolTip_showText2")
	symaddr(&addr_QToolTip_showText3, "QToolTip_showText3")
	symaddr(&addr_QToolTip_showText4, "QToolTip_showText4")
	reg(&pg_QToolTip_delete, "QToolTip_delete")
}

type QToolTip struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QToolTip) UnsafeGoID() int64     { return this.vid }
func (this *QToolTip) UnsafeSetGoID(id int64) { this.vid = id }

func bindQToolTip(h uintptr) *QToolTip {
	if h == 0 {
		return nil
	}
	return &QToolTip{h: h}
}

func newownedQToolTip(h uintptr) *QToolTip {
	o := bindQToolTip(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QToolTip).Delete)
	}
	return o
}

func (this *QToolTip) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QToolTip) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQToolTip(h unsafe.Pointer) *QToolTip { return bindQToolTip(uintptr(h)) }
func UnsafeNewQToolTipOwned(h unsafe.Pointer) *QToolTip { return newownedQToolTip(uintptr(h)) }

func (this *QToolTip) UnsafeRelease() { this.h = 0 }

func QToolTip_ShowText(pos *QPoint, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QToolTip_showText, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
}

func QToolTip_HideText() {
	purego.SyscallN(addr_QToolTip_hideText)
}

func QToolTip_IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolTip_isVisible)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QToolTip_Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolTip_text, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QToolTip_Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QToolTip_palette)
	return newownedQPalette(__r)
}

func QToolTip_SetPalette(arg1 *QPalette) {
	purego.SyscallN(addr_QToolTip_setPalette, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QToolTip_Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QToolTip_font)
	return newownedQFont(__r)
}

func QToolTip_SetFont(arg1 *QFont) {
	purego.SyscallN(addr_QToolTip_setFont, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QToolTip_ShowText2(pos *QPoint, text string, w *QWidget) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QToolTip_showText2, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
	runtime.KeepAlive(w)
}

func QToolTip_ShowText3(pos *QPoint, text string, w *QWidget, rect *QRect) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QToolTip_showText3, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)), uintptr(w.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
	runtime.KeepAlive(w)
	runtime.KeepAlive(rect)
}

func QToolTip_ShowText4(pos *QPoint, text string, w *QWidget, rect *QRect, msecShowTime int) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QToolTip_showText4, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)), uintptr(w.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(msecShowTime))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
	runtime.KeepAlive(w)
	runtime.KeepAlive(rect)
}

func (this *QToolTip) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QToolTip_delete(this.h)
	}
	this.UnsafeRelease()
}
