package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPalette_new uintptr
	addr_QPalette_new2 uintptr
	addr_QPalette_new3 uintptr
	addr_QPalette_new4 uintptr
	addr_QPalette_new5 uintptr
	addr_QPalette_new6 uintptr
	addr_QPalette_new7 uintptr
	addr_QPalette_OperatorAssign uintptr
	addr_QPalette_swap uintptr
	addr_QPalette_ToQVariant uintptr
	addr_QPalette_currentColorGroup uintptr
	addr_QPalette_setCurrentColorGroup uintptr
	addr_QPalette_color uintptr
	addr_QPalette_brush uintptr
	addr_QPalette_setColor uintptr
	addr_QPalette_setColor2 uintptr
	addr_QPalette_setBrush uintptr
	addr_QPalette_isBrushSet uintptr
	addr_QPalette_setBrush2 uintptr
	addr_QPalette_setColorGroup uintptr
	addr_QPalette_isEqual uintptr
	addr_QPalette_color2 uintptr
	addr_QPalette_brush2 uintptr
	addr_QPalette_windowText uintptr
	addr_QPalette_button uintptr
	addr_QPalette_light uintptr
	addr_QPalette_dark uintptr
	addr_QPalette_mid uintptr
	addr_QPalette_text uintptr
	addr_QPalette_base uintptr
	addr_QPalette_alternateBase uintptr
	addr_QPalette_toolTipBase uintptr
	addr_QPalette_toolTipText uintptr
	addr_QPalette_window uintptr
	addr_QPalette_midlight uintptr
	addr_QPalette_brightText uintptr
	addr_QPalette_buttonText uintptr
	addr_QPalette_shadow uintptr
	addr_QPalette_highlight uintptr
	addr_QPalette_highlightedText uintptr
	addr_QPalette_link uintptr
	addr_QPalette_linkVisited uintptr
	addr_QPalette_placeholderText uintptr
	addr_QPalette_accent uintptr
	addr_QPalette_OperatorEqual uintptr
	addr_QPalette_OperatorNotEqual uintptr
	addr_QPalette_isCopyOf uintptr
	addr_QPalette_cacheKey uintptr
	addr_QPalette_resolve uintptr
	addr_QPalette_resolveMask uintptr
	addr_QPalette_setResolveMask uintptr
	addr_QPalette_toQVariant uintptr
	pg_QPalette_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPalette_new, "QPalette_new")
	symaddr(&addr_QPalette_new2, "QPalette_new2")
	symaddr(&addr_QPalette_new3, "QPalette_new3")
	symaddr(&addr_QPalette_new4, "QPalette_new4")
	symaddr(&addr_QPalette_new5, "QPalette_new5")
	symaddr(&addr_QPalette_new6, "QPalette_new6")
	symaddr(&addr_QPalette_new7, "QPalette_new7")
	symaddr(&addr_QPalette_OperatorAssign, "QPalette_OperatorAssign")
	symaddr(&addr_QPalette_swap, "QPalette_swap")
	symaddr(&addr_QPalette_ToQVariant, "QPalette_ToQVariant")
	symaddr(&addr_QPalette_currentColorGroup, "QPalette_currentColorGroup")
	symaddr(&addr_QPalette_setCurrentColorGroup, "QPalette_setCurrentColorGroup")
	symaddr(&addr_QPalette_color, "QPalette_color")
	symaddr(&addr_QPalette_brush, "QPalette_brush")
	symaddr(&addr_QPalette_setColor, "QPalette_setColor")
	symaddr(&addr_QPalette_setColor2, "QPalette_setColor2")
	symaddr(&addr_QPalette_setBrush, "QPalette_setBrush")
	symaddr(&addr_QPalette_isBrushSet, "QPalette_isBrushSet")
	symaddr(&addr_QPalette_setBrush2, "QPalette_setBrush2")
	symaddr(&addr_QPalette_setColorGroup, "QPalette_setColorGroup")
	symaddr(&addr_QPalette_isEqual, "QPalette_isEqual")
	symaddr(&addr_QPalette_color2, "QPalette_color2")
	symaddr(&addr_QPalette_brush2, "QPalette_brush2")
	symaddr(&addr_QPalette_windowText, "QPalette_windowText")
	symaddr(&addr_QPalette_button, "QPalette_button")
	symaddr(&addr_QPalette_light, "QPalette_light")
	symaddr(&addr_QPalette_dark, "QPalette_dark")
	symaddr(&addr_QPalette_mid, "QPalette_mid")
	symaddr(&addr_QPalette_text, "QPalette_text")
	symaddr(&addr_QPalette_base, "QPalette_base")
	symaddr(&addr_QPalette_alternateBase, "QPalette_alternateBase")
	symaddr(&addr_QPalette_toolTipBase, "QPalette_toolTipBase")
	symaddr(&addr_QPalette_toolTipText, "QPalette_toolTipText")
	symaddr(&addr_QPalette_window, "QPalette_window")
	symaddr(&addr_QPalette_midlight, "QPalette_midlight")
	symaddr(&addr_QPalette_brightText, "QPalette_brightText")
	symaddr(&addr_QPalette_buttonText, "QPalette_buttonText")
	symaddr(&addr_QPalette_shadow, "QPalette_shadow")
	symaddr(&addr_QPalette_highlight, "QPalette_highlight")
	symaddr(&addr_QPalette_highlightedText, "QPalette_highlightedText")
	symaddr(&addr_QPalette_link, "QPalette_link")
	symaddr(&addr_QPalette_linkVisited, "QPalette_linkVisited")
	symaddr(&addr_QPalette_placeholderText, "QPalette_placeholderText")
	symaddr(&addr_QPalette_accent, "QPalette_accent")
	symaddr(&addr_QPalette_OperatorEqual, "QPalette_OperatorEqual")
	symaddr(&addr_QPalette_OperatorNotEqual, "QPalette_OperatorNotEqual")
	symaddr(&addr_QPalette_isCopyOf, "QPalette_isCopyOf")
	symaddr(&addr_QPalette_cacheKey, "QPalette_cacheKey")
	symaddr(&addr_QPalette_resolve, "QPalette_resolve")
	symaddr(&addr_QPalette_resolveMask, "QPalette_resolveMask")
	symaddr(&addr_QPalette_setResolveMask, "QPalette_setResolveMask")
	symaddr(&addr_QPalette_toQVariant, "QPalette_toQVariant")
	reg(&pg_QPalette_delete, "QPalette_delete")
}

type QPalette struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPalette) UnsafeGoID() int64     { return this.vid }
func (this *QPalette) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPalette(h uintptr) *QPalette {
	if h == 0 {
		return nil
	}
	return &QPalette{h: h}
}

func newownedQPalette(h uintptr) *QPalette {
	o := bindQPalette(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPalette).Delete)
	}
	return o
}

func (this *QPalette) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPalette) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPalette(h unsafe.Pointer) *QPalette { return bindQPalette(uintptr(h)) }
func UnsafeNewQPaletteOwned(h unsafe.Pointer) *QPalette { return newownedQPalette(uintptr(h)) }

func (this *QPalette) UnsafeRelease() { this.h = 0 }

func NewQPalette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new)
	__o := newownedQPalette(__r)
	return __o
}

func NewQPalette2(button *QColor) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new2, uintptr(button.UnsafePointer()))
	__o := newownedQPalette(__r)
	runtime.KeepAlive(button)
	return __o
}

func NewQPalette3(button GlobalColor) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new3, uintptr(button))
	__o := newownedQPalette(__r)
	return __o
}

func NewQPalette4(button *QColor, window *QColor) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new4, uintptr(button.UnsafePointer()), uintptr(window.UnsafePointer()))
	__o := newownedQPalette(__r)
	runtime.KeepAlive(button)
	runtime.KeepAlive(window)
	return __o
}

func NewQPalette5(windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new5, uintptr(windowText.UnsafePointer()), uintptr(button.UnsafePointer()), uintptr(light.UnsafePointer()), uintptr(dark.UnsafePointer()), uintptr(mid.UnsafePointer()), uintptr(text.UnsafePointer()), uintptr(bright_text.UnsafePointer()), uintptr(base.UnsafePointer()), uintptr(window.UnsafePointer()))
	__o := newownedQPalette(__r)
	runtime.KeepAlive(windowText)
	runtime.KeepAlive(button)
	runtime.KeepAlive(light)
	runtime.KeepAlive(dark)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(text)
	runtime.KeepAlive(bright_text)
	runtime.KeepAlive(base)
	runtime.KeepAlive(window)
	return __o
}

func NewQPalette6(windowText *QColor, window *QColor, light *QColor, dark *QColor, mid *QColor, text *QColor, base *QColor) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new6, uintptr(windowText.UnsafePointer()), uintptr(window.UnsafePointer()), uintptr(light.UnsafePointer()), uintptr(dark.UnsafePointer()), uintptr(mid.UnsafePointer()), uintptr(text.UnsafePointer()), uintptr(base.UnsafePointer()))
	__o := newownedQPalette(__r)
	runtime.KeepAlive(windowText)
	runtime.KeepAlive(window)
	runtime.KeepAlive(light)
	runtime.KeepAlive(dark)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(text)
	runtime.KeepAlive(base)
	return __o
}

func NewQPalette7(palette *QPalette) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_new7, uintptr(palette.UnsafePointer()))
	__o := newownedQPalette(__r)
	runtime.KeepAlive(palette)
	return __o
}

func (this *QPalette) OperatorAssign(palette *QPalette) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_OperatorAssign, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
	return bindQPalette(__r)
}

func (this *QPalette) Swap(other *QPalette) {
	purego.SyscallN(addr_QPalette_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPalette) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPalette_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPalette) CurrentColorGroup() QPalette__ColorGroup {
	__r0, _, _ := purego.SyscallN(addr_QPalette_currentColorGroup, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPalette__ColorGroup(__r)
}

func (this *QPalette) SetCurrentColorGroup(cg QPalette__ColorGroup) {
	purego.SyscallN(addr_QPalette_setCurrentColorGroup, this.h, uintptr(cg))
	runtime.KeepAlive(this)
}

func (this *QPalette) Color(cg QPalette__ColorGroup, cr QPalette__ColorRole) *QColor {
	__r, _, _ := purego.SyscallN(addr_QPalette_color, this.h, uintptr(cg), uintptr(cr))
	runtime.KeepAlive(this)
	return bindQColor(__r)
}

func (this *QPalette) Brush(cg QPalette__ColorGroup, cr QPalette__ColorRole) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_brush, this.h, uintptr(cg), uintptr(cr))
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) SetColor(cg QPalette__ColorGroup, cr QPalette__ColorRole, color *QColor) {
	purego.SyscallN(addr_QPalette_setColor, this.h, uintptr(cg), uintptr(cr), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPalette) SetColor2(cr QPalette__ColorRole, color *QColor) {
	purego.SyscallN(addr_QPalette_setColor2, this.h, uintptr(cr), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPalette) SetBrush(cr QPalette__ColorRole, brush *QBrush) {
	purego.SyscallN(addr_QPalette_setBrush, this.h, uintptr(cr), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QPalette) IsBrushSet(cg QPalette__ColorGroup, cr QPalette__ColorRole) bool {
	__r0, _, _ := purego.SyscallN(addr_QPalette_isBrushSet, this.h, uintptr(cg), uintptr(cr))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPalette) SetBrush2(cg QPalette__ColorGroup, cr QPalette__ColorRole, brush *QBrush) {
	purego.SyscallN(addr_QPalette_setBrush2, this.h, uintptr(cg), uintptr(cr), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QPalette) SetColorGroup(cr QPalette__ColorGroup, windowText *QBrush, button *QBrush, light *QBrush, dark *QBrush, mid *QBrush, text *QBrush, bright_text *QBrush, base *QBrush, window *QBrush) {
	purego.SyscallN(addr_QPalette_setColorGroup, this.h, uintptr(cr), uintptr(windowText.UnsafePointer()), uintptr(button.UnsafePointer()), uintptr(light.UnsafePointer()), uintptr(dark.UnsafePointer()), uintptr(mid.UnsafePointer()), uintptr(text.UnsafePointer()), uintptr(bright_text.UnsafePointer()), uintptr(base.UnsafePointer()), uintptr(window.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(windowText)
	runtime.KeepAlive(button)
	runtime.KeepAlive(light)
	runtime.KeepAlive(dark)
	runtime.KeepAlive(mid)
	runtime.KeepAlive(text)
	runtime.KeepAlive(bright_text)
	runtime.KeepAlive(base)
	runtime.KeepAlive(window)
}

func (this *QPalette) IsEqual(cr1 QPalette__ColorGroup, cr2 QPalette__ColorGroup) bool {
	__r0, _, _ := purego.SyscallN(addr_QPalette_isEqual, this.h, uintptr(cr1), uintptr(cr2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPalette) ColorWithCr(cr QPalette__ColorRole) *QColor {
	__r, _, _ := purego.SyscallN(addr_QPalette_color2, this.h, uintptr(cr))
	runtime.KeepAlive(this)
	return bindQColor(__r)
}

func (this *QPalette) BrushWithCr(cr QPalette__ColorRole) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_brush2, this.h, uintptr(cr))
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) WindowText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_windowText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Button() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_button, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Light() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_light, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Dark() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_dark, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Mid() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_mid, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Text() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_text, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Base() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_base, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) AlternateBase() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_alternateBase, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) ToolTipBase() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_toolTipBase, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) ToolTipText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_toolTipText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Window() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_window, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Midlight() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_midlight, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) BrightText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_brightText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) ButtonText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_buttonText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Shadow() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_shadow, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Highlight() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_highlight, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) HighlightedText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_highlightedText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Link() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_link, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) LinkVisited() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_linkVisited, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) PlaceholderText() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_placeholderText, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) Accent() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPalette_accent, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPalette) OperatorEqual(p *QPalette) bool {
	__r0, _, _ := purego.SyscallN(addr_QPalette_OperatorEqual, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPalette) OperatorNotEqual(p *QPalette) bool {
	__r0, _, _ := purego.SyscallN(addr_QPalette_OperatorNotEqual, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPalette) IsCopyOf(p *QPalette) bool {
	__r0, _, _ := purego.SyscallN(addr_QPalette_isCopyOf, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPalette) CacheKey() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPalette_cacheKey, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPalette) Resolve(other *QPalette) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QPalette_resolve, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQPalette(__r)
}

func (this *QPalette) ResolveMask() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QPalette_resolveMask, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QPalette) SetResolveMask(mask uint64) {
	purego.SyscallN(addr_QPalette_setResolveMask, this.h, uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QPalette) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPalette_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPalette) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPalette_delete(this.h)
	}
	this.UnsafeRelease()
}
