package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColor_new uintptr
	addr_QColor_new2 uintptr
	addr_QColor_new3 uintptr
	addr_QColor_new4 uintptr
	addr_QColor_new5 uintptr
	addr_QColor_new6 uintptr
	addr_QColor_new7 uintptr
	addr_QColor_new8 uintptr
	addr_QColor_new9 uintptr
	addr_QColor_new10 uintptr
	addr_QColor_new11 uintptr
	addr_QColor_new12 uintptr
	addr_QColor_new13 uintptr
	addr_QColor_new14 uintptr
	addr_QColor_fromString uintptr
	addr_QColor_OperatorAssign uintptr
	addr_QColor_isValid uintptr
	addr_QColor_name uintptr
	addr_QColor_setNamedColor uintptr
	addr_QColor_setNamedColor2 uintptr
	addr_QColor_setNamedColor3 uintptr
	addr_QColor_colorNames uintptr
	addr_QColor_spec uintptr
	addr_QColor_alpha uintptr
	addr_QColor_setAlpha uintptr
	pg_QColor_alphaF func(uintptr) float32
	pg_QColor_setAlphaF func(uintptr, float32)
	addr_QColor_red uintptr
	addr_QColor_green uintptr
	addr_QColor_blue uintptr
	addr_QColor_setRed uintptr
	addr_QColor_setGreen uintptr
	addr_QColor_setBlue uintptr
	pg_QColor_redF func(uintptr) float32
	pg_QColor_greenF func(uintptr) float32
	pg_QColor_blueF func(uintptr) float32
	pg_QColor_setRedF func(uintptr, float32)
	pg_QColor_setGreenF func(uintptr, float32)
	pg_QColor_setBlueF func(uintptr, float32)
	addr_QColor_getRgb uintptr
	addr_QColor_setRgb uintptr
	addr_QColor_getRgbF uintptr
	pg_QColor_setRgbF func(uintptr, float32, float32, float32)
	addr_QColor_rgba64 uintptr
	addr_QColor_setRgba64 uintptr
	addr_QColor_rgba uintptr
	addr_QColor_setRgba uintptr
	addr_QColor_rgb uintptr
	addr_QColor_setRgb2 uintptr
	addr_QColor_hue uintptr
	addr_QColor_saturation uintptr
	addr_QColor_hsvHue uintptr
	addr_QColor_hsvSaturation uintptr
	addr_QColor_value uintptr
	pg_QColor_hueF func(uintptr) float32
	pg_QColor_saturationF func(uintptr) float32
	pg_QColor_hsvHueF func(uintptr) float32
	pg_QColor_hsvSaturationF func(uintptr) float32
	pg_QColor_valueF func(uintptr) float32
	addr_QColor_getHsv uintptr
	addr_QColor_setHsv uintptr
	addr_QColor_getHsvF uintptr
	pg_QColor_setHsvF func(uintptr, float32, float32, float32)
	addr_QColor_cyan uintptr
	addr_QColor_magenta uintptr
	addr_QColor_yellow uintptr
	addr_QColor_black uintptr
	pg_QColor_cyanF func(uintptr) float32
	pg_QColor_magentaF func(uintptr) float32
	pg_QColor_yellowF func(uintptr) float32
	pg_QColor_blackF func(uintptr) float32
	addr_QColor_getCmyk uintptr
	addr_QColor_setCmyk uintptr
	addr_QColor_getCmykF uintptr
	pg_QColor_setCmykF func(uintptr, float32, float32, float32, float32)
	addr_QColor_hslHue uintptr
	addr_QColor_hslSaturation uintptr
	addr_QColor_lightness uintptr
	pg_QColor_hslHueF func(uintptr) float32
	pg_QColor_hslSaturationF func(uintptr) float32
	pg_QColor_lightnessF func(uintptr) float32
	addr_QColor_getHsl uintptr
	addr_QColor_setHsl uintptr
	addr_QColor_getHslF uintptr
	pg_QColor_setHslF func(uintptr, float32, float32, float32)
	addr_QColor_toRgb uintptr
	addr_QColor_toHsv uintptr
	addr_QColor_toCmyk uintptr
	addr_QColor_toHsl uintptr
	addr_QColor_toExtendedRgb uintptr
	addr_QColor_convertTo uintptr
	addr_QColor_fromRgb uintptr
	addr_QColor_fromRgba uintptr
	addr_QColor_fromRgb2 uintptr
	pg_QColor_fromRgbF func(float32, float32, float32) uintptr
	addr_QColor_fromRgba64 uintptr
	addr_QColor_fromRgba642 uintptr
	addr_QColor_fromHsv uintptr
	pg_QColor_fromHsvF func(float32, float32, float32) uintptr
	addr_QColor_fromCmyk uintptr
	pg_QColor_fromCmykF func(float32, float32, float32, float32) uintptr
	addr_QColor_fromHsl uintptr
	pg_QColor_fromHslF func(float32, float32, float32) uintptr
	addr_QColor_lighter uintptr
	addr_QColor_darker uintptr
	addr_QColor_OperatorEqual uintptr
	addr_QColor_OperatorNotEqual uintptr
	addr_QColor_ToQVariant uintptr
	addr_QColor_isValidColor uintptr
	addr_QColor_isValidColor2 uintptr
	addr_QColor_isValidColor3 uintptr
	addr_QColor_isValidColorName uintptr
	addr_QColor_OperatorAssign2 uintptr
	addr_QColor_name2 uintptr
	addr_QColor_getRgb2 uintptr
	addr_QColor_setRgb3 uintptr
	addr_QColor_getRgbF2 uintptr
	pg_QColor_setRgbF2 func(uintptr, float32, float32, float32, float32)
	addr_QColor_getHsv2 uintptr
	addr_QColor_setHsv2 uintptr
	addr_QColor_getHsvF2 uintptr
	pg_QColor_setHsvF2 func(uintptr, float32, float32, float32, float32)
	addr_QColor_getCmyk2 uintptr
	addr_QColor_setCmyk2 uintptr
	addr_QColor_getCmykF2 uintptr
	pg_QColor_setCmykF2 func(uintptr, float32, float32, float32, float32, float32)
	addr_QColor_getHsl2 uintptr
	addr_QColor_setHsl2 uintptr
	addr_QColor_getHslF2 uintptr
	pg_QColor_setHslF2 func(uintptr, float32, float32, float32, float32)
	addr_QColor_fromRgb3 uintptr
	pg_QColor_fromRgbF2 func(float32, float32, float32, float32) uintptr
	addr_QColor_fromRgba643 uintptr
	addr_QColor_fromHsv2 uintptr
	pg_QColor_fromHsvF2 func(float32, float32, float32, float32) uintptr
	addr_QColor_fromCmyk2 uintptr
	pg_QColor_fromCmykF2 func(float32, float32, float32, float32, float32) uintptr
	addr_QColor_fromHsl2 uintptr
	pg_QColor_fromHslF2 func(float32, float32, float32, float32) uintptr
	addr_QColor_lighter2 uintptr
	addr_QColor_darker2 uintptr
	addr_QColor_OperatorEqual2 uintptr
	addr_QColor_OperatorNotEqual2 uintptr
	addr_QColor_OperatorEqual3 uintptr
	addr_QColor_OperatorNotEqual3 uintptr
	addr_QColor_toQVariant uintptr
	pg_QColor_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColor_new, "QColor_new")
	symaddr(&addr_QColor_new2, "QColor_new2")
	symaddr(&addr_QColor_new3, "QColor_new3")
	symaddr(&addr_QColor_new4, "QColor_new4")
	symaddr(&addr_QColor_new5, "QColor_new5")
	symaddr(&addr_QColor_new6, "QColor_new6")
	symaddr(&addr_QColor_new7, "QColor_new7")
	symaddr(&addr_QColor_new8, "QColor_new8")
	symaddr(&addr_QColor_new9, "QColor_new9")
	symaddr(&addr_QColor_new10, "QColor_new10")
	symaddr(&addr_QColor_new11, "QColor_new11")
	symaddr(&addr_QColor_new12, "QColor_new12")
	symaddr(&addr_QColor_new13, "QColor_new13")
	symaddr(&addr_QColor_new14, "QColor_new14")
	symaddr(&addr_QColor_fromString, "QColor_fromString")
	symaddr(&addr_QColor_OperatorAssign, "QColor_OperatorAssign")
	symaddr(&addr_QColor_isValid, "QColor_isValid")
	symaddr(&addr_QColor_name, "QColor_name")
	symaddr(&addr_QColor_setNamedColor, "QColor_setNamedColor")
	symaddr(&addr_QColor_setNamedColor2, "QColor_setNamedColor2")
	symaddr(&addr_QColor_setNamedColor3, "QColor_setNamedColor3")
	symaddr(&addr_QColor_colorNames, "QColor_colorNames")
	symaddr(&addr_QColor_spec, "QColor_spec")
	symaddr(&addr_QColor_alpha, "QColor_alpha")
	symaddr(&addr_QColor_setAlpha, "QColor_setAlpha")
	reg(&pg_QColor_alphaF, "QColor_alphaF")
	reg(&pg_QColor_setAlphaF, "QColor_setAlphaF")
	symaddr(&addr_QColor_red, "QColor_red")
	symaddr(&addr_QColor_green, "QColor_green")
	symaddr(&addr_QColor_blue, "QColor_blue")
	symaddr(&addr_QColor_setRed, "QColor_setRed")
	symaddr(&addr_QColor_setGreen, "QColor_setGreen")
	symaddr(&addr_QColor_setBlue, "QColor_setBlue")
	reg(&pg_QColor_redF, "QColor_redF")
	reg(&pg_QColor_greenF, "QColor_greenF")
	reg(&pg_QColor_blueF, "QColor_blueF")
	reg(&pg_QColor_setRedF, "QColor_setRedF")
	reg(&pg_QColor_setGreenF, "QColor_setGreenF")
	reg(&pg_QColor_setBlueF, "QColor_setBlueF")
	symaddr(&addr_QColor_getRgb, "QColor_getRgb")
	symaddr(&addr_QColor_setRgb, "QColor_setRgb")
	symaddr(&addr_QColor_getRgbF, "QColor_getRgbF")
	reg(&pg_QColor_setRgbF, "QColor_setRgbF")
	symaddr(&addr_QColor_rgba64, "QColor_rgba64")
	symaddr(&addr_QColor_setRgba64, "QColor_setRgba64")
	symaddr(&addr_QColor_rgba, "QColor_rgba")
	symaddr(&addr_QColor_setRgba, "QColor_setRgba")
	symaddr(&addr_QColor_rgb, "QColor_rgb")
	symaddr(&addr_QColor_setRgb2, "QColor_setRgb2")
	symaddr(&addr_QColor_hue, "QColor_hue")
	symaddr(&addr_QColor_saturation, "QColor_saturation")
	symaddr(&addr_QColor_hsvHue, "QColor_hsvHue")
	symaddr(&addr_QColor_hsvSaturation, "QColor_hsvSaturation")
	symaddr(&addr_QColor_value, "QColor_value")
	reg(&pg_QColor_hueF, "QColor_hueF")
	reg(&pg_QColor_saturationF, "QColor_saturationF")
	reg(&pg_QColor_hsvHueF, "QColor_hsvHueF")
	reg(&pg_QColor_hsvSaturationF, "QColor_hsvSaturationF")
	reg(&pg_QColor_valueF, "QColor_valueF")
	symaddr(&addr_QColor_getHsv, "QColor_getHsv")
	symaddr(&addr_QColor_setHsv, "QColor_setHsv")
	symaddr(&addr_QColor_getHsvF, "QColor_getHsvF")
	reg(&pg_QColor_setHsvF, "QColor_setHsvF")
	symaddr(&addr_QColor_cyan, "QColor_cyan")
	symaddr(&addr_QColor_magenta, "QColor_magenta")
	symaddr(&addr_QColor_yellow, "QColor_yellow")
	symaddr(&addr_QColor_black, "QColor_black")
	reg(&pg_QColor_cyanF, "QColor_cyanF")
	reg(&pg_QColor_magentaF, "QColor_magentaF")
	reg(&pg_QColor_yellowF, "QColor_yellowF")
	reg(&pg_QColor_blackF, "QColor_blackF")
	symaddr(&addr_QColor_getCmyk, "QColor_getCmyk")
	symaddr(&addr_QColor_setCmyk, "QColor_setCmyk")
	symaddr(&addr_QColor_getCmykF, "QColor_getCmykF")
	reg(&pg_QColor_setCmykF, "QColor_setCmykF")
	symaddr(&addr_QColor_hslHue, "QColor_hslHue")
	symaddr(&addr_QColor_hslSaturation, "QColor_hslSaturation")
	symaddr(&addr_QColor_lightness, "QColor_lightness")
	reg(&pg_QColor_hslHueF, "QColor_hslHueF")
	reg(&pg_QColor_hslSaturationF, "QColor_hslSaturationF")
	reg(&pg_QColor_lightnessF, "QColor_lightnessF")
	symaddr(&addr_QColor_getHsl, "QColor_getHsl")
	symaddr(&addr_QColor_setHsl, "QColor_setHsl")
	symaddr(&addr_QColor_getHslF, "QColor_getHslF")
	reg(&pg_QColor_setHslF, "QColor_setHslF")
	symaddr(&addr_QColor_toRgb, "QColor_toRgb")
	symaddr(&addr_QColor_toHsv, "QColor_toHsv")
	symaddr(&addr_QColor_toCmyk, "QColor_toCmyk")
	symaddr(&addr_QColor_toHsl, "QColor_toHsl")
	symaddr(&addr_QColor_toExtendedRgb, "QColor_toExtendedRgb")
	symaddr(&addr_QColor_convertTo, "QColor_convertTo")
	symaddr(&addr_QColor_fromRgb, "QColor_fromRgb")
	symaddr(&addr_QColor_fromRgba, "QColor_fromRgba")
	symaddr(&addr_QColor_fromRgb2, "QColor_fromRgb2")
	reg(&pg_QColor_fromRgbF, "QColor_fromRgbF")
	symaddr(&addr_QColor_fromRgba64, "QColor_fromRgba64")
	symaddr(&addr_QColor_fromRgba642, "QColor_fromRgba642")
	symaddr(&addr_QColor_fromHsv, "QColor_fromHsv")
	reg(&pg_QColor_fromHsvF, "QColor_fromHsvF")
	symaddr(&addr_QColor_fromCmyk, "QColor_fromCmyk")
	reg(&pg_QColor_fromCmykF, "QColor_fromCmykF")
	symaddr(&addr_QColor_fromHsl, "QColor_fromHsl")
	reg(&pg_QColor_fromHslF, "QColor_fromHslF")
	symaddr(&addr_QColor_lighter, "QColor_lighter")
	symaddr(&addr_QColor_darker, "QColor_darker")
	symaddr(&addr_QColor_OperatorEqual, "QColor_OperatorEqual")
	symaddr(&addr_QColor_OperatorNotEqual, "QColor_OperatorNotEqual")
	symaddr(&addr_QColor_ToQVariant, "QColor_ToQVariant")
	symaddr(&addr_QColor_isValidColor, "QColor_isValidColor")
	symaddr(&addr_QColor_isValidColor2, "QColor_isValidColor2")
	symaddr(&addr_QColor_isValidColor3, "QColor_isValidColor3")
	symaddr(&addr_QColor_isValidColorName, "QColor_isValidColorName")
	symaddr(&addr_QColor_OperatorAssign2, "QColor_OperatorAssign2")
	symaddr(&addr_QColor_name2, "QColor_name2")
	symaddr(&addr_QColor_getRgb2, "QColor_getRgb2")
	symaddr(&addr_QColor_setRgb3, "QColor_setRgb3")
	symaddr(&addr_QColor_getRgbF2, "QColor_getRgbF2")
	reg(&pg_QColor_setRgbF2, "QColor_setRgbF2")
	symaddr(&addr_QColor_getHsv2, "QColor_getHsv2")
	symaddr(&addr_QColor_setHsv2, "QColor_setHsv2")
	symaddr(&addr_QColor_getHsvF2, "QColor_getHsvF2")
	reg(&pg_QColor_setHsvF2, "QColor_setHsvF2")
	symaddr(&addr_QColor_getCmyk2, "QColor_getCmyk2")
	symaddr(&addr_QColor_setCmyk2, "QColor_setCmyk2")
	symaddr(&addr_QColor_getCmykF2, "QColor_getCmykF2")
	reg(&pg_QColor_setCmykF2, "QColor_setCmykF2")
	symaddr(&addr_QColor_getHsl2, "QColor_getHsl2")
	symaddr(&addr_QColor_setHsl2, "QColor_setHsl2")
	symaddr(&addr_QColor_getHslF2, "QColor_getHslF2")
	reg(&pg_QColor_setHslF2, "QColor_setHslF2")
	symaddr(&addr_QColor_fromRgb3, "QColor_fromRgb3")
	reg(&pg_QColor_fromRgbF2, "QColor_fromRgbF2")
	symaddr(&addr_QColor_fromRgba643, "QColor_fromRgba643")
	symaddr(&addr_QColor_fromHsv2, "QColor_fromHsv2")
	reg(&pg_QColor_fromHsvF2, "QColor_fromHsvF2")
	symaddr(&addr_QColor_fromCmyk2, "QColor_fromCmyk2")
	reg(&pg_QColor_fromCmykF2, "QColor_fromCmykF2")
	symaddr(&addr_QColor_fromHsl2, "QColor_fromHsl2")
	reg(&pg_QColor_fromHslF2, "QColor_fromHslF2")
	symaddr(&addr_QColor_lighter2, "QColor_lighter2")
	symaddr(&addr_QColor_darker2, "QColor_darker2")
	symaddr(&addr_QColor_OperatorEqual2, "QColor_OperatorEqual2")
	symaddr(&addr_QColor_OperatorNotEqual2, "QColor_OperatorNotEqual2")
	symaddr(&addr_QColor_OperatorEqual3, "QColor_OperatorEqual3")
	symaddr(&addr_QColor_OperatorNotEqual3, "QColor_OperatorNotEqual3")
	symaddr(&addr_QColor_toQVariant, "QColor_toQVariant")
	reg(&pg_QColor_delete, "QColor_delete")
}

type QColor struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QColor) UnsafeGoID() int64     { return this.vid }
func (this *QColor) UnsafeSetGoID(id int64) { this.vid = id }

func bindQColor(h uintptr) *QColor {
	if h == 0 {
		return nil
	}
	return &QColor{h: h}
}

func newownedQColor(h uintptr) *QColor {
	o := bindQColor(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QColor).Delete)
	}
	return o
}

func (this *QColor) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColor) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColor(h unsafe.Pointer) *QColor { return bindQColor(uintptr(h)) }
func UnsafeNewQColorOwned(h unsafe.Pointer) *QColor { return newownedQColor(uintptr(h)) }

func (this *QColor) UnsafeRelease() { this.h = 0 }

func NewQColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new)
	__o := newownedQColor(__r)
	return __o
}

func NewQColor2(color GlobalColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new2, uintptr(color))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor3(r int, g int, b int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new3, uintptr(r), uintptr(g), uintptr(b))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor4(rgb uint) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new4, uintptr(rgb))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor5(rgba64 *QRgba64) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new5, uintptr(rgba64.UnsafePointer()))
	__o := newownedQColor(__r)
	runtime.KeepAlive(rgba64)
	return __o
}

func NewQColor6(name string) *QColor {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QColor_new6, name_p, uintptr(len(name)))
	__o := newownedQColor(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQColor7(name *QStringView) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new7, uintptr(name.UnsafePointer()))
	__o := newownedQColor(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQColor8(aname string) *QColor {
	aname_b := append([]byte(aname), 0)
	__r, _, _ := purego.SyscallN(addr_QColor_new8, uintptr(unsafe.Pointer(&aname_b[0])))
	__o := newownedQColor(__r)
	runtime.KeepAlive(aname_b)
	return __o
}

func NewQColor9(name *QLatin1String) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new9, uintptr(name.UnsafePointer()))
	__o := newownedQColor(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQColor10(spec QColor__Spec) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new10, uintptr(spec))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor11(spec QColor__Spec, a1 uint16, a2 uint16, a3 uint16, a4 uint16) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new11, uintptr(spec), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor12(arg1 *QColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new12, uintptr(arg1.UnsafePointer()))
	__o := newownedQColor(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQColor13(r int, g int, b int, a int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new13, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	__o := newownedQColor(__r)
	return __o
}

func NewQColor14(spec QColor__Spec, a1 uint16, a2 uint16, a3 uint16, a4 uint16, a5 uint16) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_new14, uintptr(spec), uintptr(a1), uintptr(a2), uintptr(a3), uintptr(a4), uintptr(a5))
	__o := newownedQColor(__r)
	return __o
}

func QColor_FromString(name string) *QColor {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QColor_fromString, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
	return newownedQColor(__r)
}

func (this *QColor) OperatorAssign(color GlobalColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_OperatorAssign, this.h, uintptr(color))
	runtime.KeepAlive(this)
	return bindQColor(__r)
}

func (this *QColor) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColor) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColor_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QColor) SetNamedColor(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QColor_setNamedColor, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QColor) SetNamedColorWithName(name *QStringView) {
	purego.SyscallN(addr_QColor_setNamedColor2, this.h, uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QColor) SetNamedColor2(name *QLatin1String) {
	purego.SyscallN(addr_QColor_setNamedColor3, this.h, uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func QColor_ColorNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColor_colorNames, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QColor) Spec() QColor__Spec {
	__r0, _, _ := purego.SyscallN(addr_QColor_spec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColor__Spec(__r)
}

func (this *QColor) Alpha() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_alpha, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) SetAlpha(alpha int) {
	purego.SyscallN(addr_QColor_setAlpha, this.h, uintptr(alpha))
	runtime.KeepAlive(this)
}

func (this *QColor) AlphaF() float32 {
	__r := pg_QColor_alphaF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) SetAlphaF(alpha float32) {
	pg_QColor_setAlphaF(this.h, alpha)
	runtime.KeepAlive(this)
}

func (this *QColor) Red() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_red, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Green() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_green, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Blue() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_blue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) SetRed(red int) {
	purego.SyscallN(addr_QColor_setRed, this.h, uintptr(red))
	runtime.KeepAlive(this)
}

func (this *QColor) SetGreen(green int) {
	purego.SyscallN(addr_QColor_setGreen, this.h, uintptr(green))
	runtime.KeepAlive(this)
}

func (this *QColor) SetBlue(blue int) {
	purego.SyscallN(addr_QColor_setBlue, this.h, uintptr(blue))
	runtime.KeepAlive(this)
}

func (this *QColor) RedF() float32 {
	__r := pg_QColor_redF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) GreenF() float32 {
	__r := pg_QColor_greenF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) BlueF() float32 {
	__r := pg_QColor_blueF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) SetRedF(red float32) {
	pg_QColor_setRedF(this.h, red)
	runtime.KeepAlive(this)
}

func (this *QColor) SetGreenF(green float32) {
	pg_QColor_setGreenF(this.h, green)
	runtime.KeepAlive(this)
}

func (this *QColor) SetBlueF(blue float32) {
	pg_QColor_setBlueF(this.h, blue)
	runtime.KeepAlive(this)
}

func (this *QColor) GetRgb() (int, int, int) {
	var r_out int32
	var g_out int32
	var b_out int32
	purego.SyscallN(addr_QColor_getRgb, this.h, uintptr(unsafe.Pointer(&r_out)), uintptr(unsafe.Pointer(&g_out)), uintptr(unsafe.Pointer(&b_out)))
	runtime.KeepAlive(this)
	return int(r_out), int(g_out), int(b_out)
}

func (this *QColor) SetRgb(r int, g int, b int) {
	purego.SyscallN(addr_QColor_setRgb, this.h, uintptr(r), uintptr(g), uintptr(b))
	runtime.KeepAlive(this)
}

func (this *QColor) GetRgbF() (float32, float32, float32) {
	var r_out float32
	var g_out float32
	var b_out float32
	purego.SyscallN(addr_QColor_getRgbF, this.h, uintptr(unsafe.Pointer(&r_out)), uintptr(unsafe.Pointer(&g_out)), uintptr(unsafe.Pointer(&b_out)))
	runtime.KeepAlive(this)
	return r_out, g_out, b_out
}

func (this *QColor) SetRgbF(r float32, g float32, b float32) {
	pg_QColor_setRgbF(this.h, r, g, b)
	runtime.KeepAlive(this)
}

func (this *QColor) Rgba64() *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QColor_rgba64, this.h)
	runtime.KeepAlive(this)
	return newownedQRgba64(__r)
}

func (this *QColor) SetRgba64(rgba *QRgba64) {
	purego.SyscallN(addr_QColor_setRgba64, this.h, uintptr(rgba.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rgba)
}

func (this *QColor) Rgba() uint {
	__r0, _, _ := purego.SyscallN(addr_QColor_rgba, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QColor) SetRgba(rgba uint) {
	purego.SyscallN(addr_QColor_setRgba, this.h, uintptr(rgba))
	runtime.KeepAlive(this)
}

func (this *QColor) Rgb() uint {
	__r0, _, _ := purego.SyscallN(addr_QColor_rgb, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QColor) SetRgbWithRgb(rgb uint) {
	purego.SyscallN(addr_QColor_setRgb2, this.h, uintptr(rgb))
	runtime.KeepAlive(this)
}

func (this *QColor) Hue() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_hue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Saturation() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_saturation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) HsvHue() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_hsvHue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) HsvSaturation() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_hsvSaturation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Value() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) HueF() float32 {
	__r := pg_QColor_hueF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) SaturationF() float32 {
	__r := pg_QColor_saturationF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) HsvHueF() float32 {
	__r := pg_QColor_hsvHueF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) HsvSaturationF() float32 {
	__r := pg_QColor_hsvSaturationF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) ValueF() float32 {
	__r := pg_QColor_valueF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) GetHsv() (int, int, int) {
	var h_out int32
	var s_out int32
	var v_out int32
	purego.SyscallN(addr_QColor_getHsv, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&v_out)))
	runtime.KeepAlive(this)
	return int(h_out), int(s_out), int(v_out)
}

func (this *QColor) SetHsv(h int, s int, v int) {
	purego.SyscallN(addr_QColor_setHsv, this.h, uintptr(h), uintptr(s), uintptr(v))
	runtime.KeepAlive(this)
}

func (this *QColor) GetHsvF() (float32, float32, float32) {
	var h_out float32
	var s_out float32
	var v_out float32
	purego.SyscallN(addr_QColor_getHsvF, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&v_out)))
	runtime.KeepAlive(this)
	return h_out, s_out, v_out
}

func (this *QColor) SetHsvF(h float32, s float32, v float32) {
	pg_QColor_setHsvF(this.h, h, s, v)
	runtime.KeepAlive(this)
}

func (this *QColor) Cyan() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_cyan, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Magenta() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_magenta, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Yellow() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_yellow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Black() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_black, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) CyanF() float32 {
	__r := pg_QColor_cyanF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) MagentaF() float32 {
	__r := pg_QColor_magentaF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) YellowF() float32 {
	__r := pg_QColor_yellowF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) BlackF() float32 {
	__r := pg_QColor_blackF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) GetCmyk() (int, int, int, int) {
	var c_out int32
	var m_out int32
	var y_out int32
	var k_out int32
	purego.SyscallN(addr_QColor_getCmyk, this.h, uintptr(unsafe.Pointer(&c_out)), uintptr(unsafe.Pointer(&m_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&k_out)))
	runtime.KeepAlive(this)
	return int(c_out), int(m_out), int(y_out), int(k_out)
}

func (this *QColor) SetCmyk(c int, m int, y int, k int) {
	purego.SyscallN(addr_QColor_setCmyk, this.h, uintptr(c), uintptr(m), uintptr(y), uintptr(k))
	runtime.KeepAlive(this)
}

func (this *QColor) GetCmykF() (float32, float32, float32, float32) {
	var c_out float32
	var m_out float32
	var y_out float32
	var k_out float32
	purego.SyscallN(addr_QColor_getCmykF, this.h, uintptr(unsafe.Pointer(&c_out)), uintptr(unsafe.Pointer(&m_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&k_out)))
	runtime.KeepAlive(this)
	return c_out, m_out, y_out, k_out
}

func (this *QColor) SetCmykF(c float32, m float32, y float32, k float32) {
	pg_QColor_setCmykF(this.h, c, m, y, k)
	runtime.KeepAlive(this)
}

func (this *QColor) HslHue() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_hslHue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) HslSaturation() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_hslSaturation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) Lightness() int {
	__r0, _, _ := purego.SyscallN(addr_QColor_lightness, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColor) HslHueF() float32 {
	__r := pg_QColor_hslHueF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) HslSaturationF() float32 {
	__r := pg_QColor_hslSaturationF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) LightnessF() float32 {
	__r := pg_QColor_lightnessF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColor) GetHsl() (int, int, int) {
	var h_out int32
	var s_out int32
	var l_out int32
	purego.SyscallN(addr_QColor_getHsl, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&l_out)))
	runtime.KeepAlive(this)
	return int(h_out), int(s_out), int(l_out)
}

func (this *QColor) SetHsl(h int, s int, l int) {
	purego.SyscallN(addr_QColor_setHsl, this.h, uintptr(h), uintptr(s), uintptr(l))
	runtime.KeepAlive(this)
}

func (this *QColor) GetHslF() (float32, float32, float32) {
	var h_out float32
	var s_out float32
	var l_out float32
	purego.SyscallN(addr_QColor_getHslF, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&l_out)))
	runtime.KeepAlive(this)
	return h_out, s_out, l_out
}

func (this *QColor) SetHslF(h float32, s float32, l float32) {
	pg_QColor_setHslF(this.h, h, s, l)
	runtime.KeepAlive(this)
}

func (this *QColor) ToRgb() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_toRgb, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) ToHsv() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_toHsv, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) ToCmyk() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_toCmyk, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) ToHsl() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_toHsl, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) ToExtendedRgb() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_toExtendedRgb, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) ConvertTo(colorSpec QColor__Spec) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_convertTo, this.h, uintptr(colorSpec))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func QColor_FromRgb(rgb uint) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgb, uintptr(rgb))
	return newownedQColor(__r)
}

func QColor_FromRgba(rgba uint) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgba, uintptr(rgba))
	return newownedQColor(__r)
}

func QColor_FromRgb2(r int, g int, b int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgb2, uintptr(r), uintptr(g), uintptr(b))
	return newownedQColor(__r)
}

func QColor_FromRgbF(r float32, g float32, b float32) *QColor {
	__r := pg_QColor_fromRgbF(r, g, b)
	return newownedQColor(__r)
}

func QColor_FromRgba64(r uint16, g uint16, b uint16) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgba64, uintptr(r), uintptr(g), uintptr(b))
	return newownedQColor(__r)
}

func QColor_FromRgba64WithRgba(rgba *QRgba64) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgba642, uintptr(rgba.UnsafePointer()))
	runtime.KeepAlive(rgba)
	return newownedQColor(__r)
}

func QColor_FromHsv(h int, s int, v int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromHsv, uintptr(h), uintptr(s), uintptr(v))
	return newownedQColor(__r)
}

func QColor_FromHsvF(h float32, s float32, v float32) *QColor {
	__r := pg_QColor_fromHsvF(h, s, v)
	return newownedQColor(__r)
}

func QColor_FromCmyk(c int, m int, y int, k int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromCmyk, uintptr(c), uintptr(m), uintptr(y), uintptr(k))
	return newownedQColor(__r)
}

func QColor_FromCmykF(c float32, m float32, y float32, k float32) *QColor {
	__r := pg_QColor_fromCmykF(c, m, y, k)
	return newownedQColor(__r)
}

func QColor_FromHsl(h int, s int, l int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromHsl, uintptr(h), uintptr(s), uintptr(l))
	return newownedQColor(__r)
}

func QColor_FromHslF(h float32, s float32, l float32) *QColor {
	__r := pg_QColor_fromHslF(h, s, l)
	return newownedQColor(__r)
}

func (this *QColor) Lighter() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_lighter, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) Darker() *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_darker, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) OperatorEqual(c *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorEqual, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QColor) OperatorNotEqual(c *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorNotEqual, this.h, uintptr(c.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QColor) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QColor_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QColor_IsValidColor(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QColor_isValidColor, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name)
	return __r != 0
}

func QColor_IsValidColorWithQStringView(arg1 *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_isValidColor2, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(arg1)
	return __r != 0
}

func QColor_IsValidColorWith(arg1 *QLatin1String) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_isValidColor3, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(arg1)
	return __r != 0
}

func QColor_IsValidColorName(arg1 string) bool {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QColor_isValidColorName, arg1_p, uintptr(len(arg1)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QColor) OperatorAssignWithQColor(arg1 *QColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_OperatorAssign2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQColor(__r)
}

func (this *QColor) NameWithFormat(format QColor__NameFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColor_name2, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QColor) GetRgb2() (int, int, int, int) {
	var r_out int32
	var g_out int32
	var b_out int32
	var a_out int32
	purego.SyscallN(addr_QColor_getRgb2, this.h, uintptr(unsafe.Pointer(&r_out)), uintptr(unsafe.Pointer(&g_out)), uintptr(unsafe.Pointer(&b_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return int(r_out), int(g_out), int(b_out), int(a_out)
}

func (this *QColor) SetRgb2(r int, g int, b int, a int) {
	purego.SyscallN(addr_QColor_setRgb3, this.h, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QColor) GetRgbF2() (float32, float32, float32, float32) {
	var r_out float32
	var g_out float32
	var b_out float32
	var a_out float32
	purego.SyscallN(addr_QColor_getRgbF2, this.h, uintptr(unsafe.Pointer(&r_out)), uintptr(unsafe.Pointer(&g_out)), uintptr(unsafe.Pointer(&b_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return r_out, g_out, b_out, a_out
}

func (this *QColor) SetRgbF2(r float32, g float32, b float32, a float32) {
	pg_QColor_setRgbF2(this.h, r, g, b, a)
	runtime.KeepAlive(this)
}

func (this *QColor) GetHsv2() (int, int, int, int) {
	var h_out int32
	var s_out int32
	var v_out int32
	var a_out int32
	purego.SyscallN(addr_QColor_getHsv2, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&v_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return int(h_out), int(s_out), int(v_out), int(a_out)
}

func (this *QColor) SetHsv2(h int, s int, v int, a int) {
	purego.SyscallN(addr_QColor_setHsv2, this.h, uintptr(h), uintptr(s), uintptr(v), uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QColor) GetHsvF2() (float32, float32, float32, float32) {
	var h_out float32
	var s_out float32
	var v_out float32
	var a_out float32
	purego.SyscallN(addr_QColor_getHsvF2, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&v_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return h_out, s_out, v_out, a_out
}

func (this *QColor) SetHsvF2(h float32, s float32, v float32, a float32) {
	pg_QColor_setHsvF2(this.h, h, s, v, a)
	runtime.KeepAlive(this)
}

func (this *QColor) GetCmyk2() (int, int, int, int, int) {
	var c_out int32
	var m_out int32
	var y_out int32
	var k_out int32
	var a_out int32
	purego.SyscallN(addr_QColor_getCmyk2, this.h, uintptr(unsafe.Pointer(&c_out)), uintptr(unsafe.Pointer(&m_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&k_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return int(c_out), int(m_out), int(y_out), int(k_out), int(a_out)
}

func (this *QColor) SetCmyk2(c int, m int, y int, k int, a int) {
	purego.SyscallN(addr_QColor_setCmyk2, this.h, uintptr(c), uintptr(m), uintptr(y), uintptr(k), uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QColor) GetCmykF2() (float32, float32, float32, float32, float32) {
	var c_out float32
	var m_out float32
	var y_out float32
	var k_out float32
	var a_out float32
	purego.SyscallN(addr_QColor_getCmykF2, this.h, uintptr(unsafe.Pointer(&c_out)), uintptr(unsafe.Pointer(&m_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&k_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return c_out, m_out, y_out, k_out, a_out
}

func (this *QColor) SetCmykF2(c float32, m float32, y float32, k float32, a float32) {
	pg_QColor_setCmykF2(this.h, c, m, y, k, a)
	runtime.KeepAlive(this)
}

func (this *QColor) GetHsl2() (int, int, int, int) {
	var h_out int32
	var s_out int32
	var l_out int32
	var a_out int32
	purego.SyscallN(addr_QColor_getHsl2, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&l_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return int(h_out), int(s_out), int(l_out), int(a_out)
}

func (this *QColor) SetHsl2(h int, s int, l int, a int) {
	purego.SyscallN(addr_QColor_setHsl2, this.h, uintptr(h), uintptr(s), uintptr(l), uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QColor) GetHslF2() (float32, float32, float32, float32) {
	var h_out float32
	var s_out float32
	var l_out float32
	var a_out float32
	purego.SyscallN(addr_QColor_getHslF2, this.h, uintptr(unsafe.Pointer(&h_out)), uintptr(unsafe.Pointer(&s_out)), uintptr(unsafe.Pointer(&l_out)), uintptr(unsafe.Pointer(&a_out)))
	runtime.KeepAlive(this)
	return h_out, s_out, l_out, a_out
}

func (this *QColor) SetHslF2(h float32, s float32, l float32, a float32) {
	pg_QColor_setHslF2(this.h, h, s, l, a)
	runtime.KeepAlive(this)
}

func QColor_FromRgb3(r int, g int, b int, a int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgb3, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	return newownedQColor(__r)
}

func QColor_FromRgbF2(r float32, g float32, b float32, a float32) *QColor {
	__r := pg_QColor_fromRgbF2(r, g, b, a)
	return newownedQColor(__r)
}

func QColor_FromRgba642(r uint16, g uint16, b uint16, a uint16) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromRgba643, uintptr(r), uintptr(g), uintptr(b), uintptr(a))
	return newownedQColor(__r)
}

func QColor_FromHsv2(h int, s int, v int, a int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromHsv2, uintptr(h), uintptr(s), uintptr(v), uintptr(a))
	return newownedQColor(__r)
}

func QColor_FromHsvF2(h float32, s float32, v float32, a float32) *QColor {
	__r := pg_QColor_fromHsvF2(h, s, v, a)
	return newownedQColor(__r)
}

func QColor_FromCmyk2(c int, m int, y int, k int, a int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromCmyk2, uintptr(c), uintptr(m), uintptr(y), uintptr(k), uintptr(a))
	return newownedQColor(__r)
}

func QColor_FromCmykF2(c float32, m float32, y float32, k float32, a float32) *QColor {
	__r := pg_QColor_fromCmykF2(c, m, y, k, a)
	return newownedQColor(__r)
}

func QColor_FromHsl2(h int, s int, l int, a int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_fromHsl2, uintptr(h), uintptr(s), uintptr(l), uintptr(a))
	return newownedQColor(__r)
}

func QColor_FromHslF2(h float32, s float32, l float32, a float32) *QColor {
	__r := pg_QColor_fromHslF2(h, s, l, a)
	return newownedQColor(__r)
}

func (this *QColor) LighterWithInt(f int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_lighter2, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) DarkerWithInt(f int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColor_darker2, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColor) OperatorEqualWithRhs(rhs *QBrush) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QColor) OperatorNotEqualWithRhs(rhs *QBrush) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QColor) OperatorEqual2(rhs *QPen) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QColor) OperatorNotEqual2(rhs *QPen) bool {
	__r0, _, _ := purego.SyscallN(addr_QColor_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QColor) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QColor_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QColor) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColor_delete(this.h)
	}
	this.UnsafeRelease()
}
