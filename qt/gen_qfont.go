package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFont_new uintptr
	addr_QFont_new2 uintptr
	addr_QFont_new3 uintptr
	addr_QFont_new4 uintptr
	addr_QFont_new5 uintptr
	addr_QFont_new6 uintptr
	addr_QFont_new7 uintptr
	addr_QFont_new8 uintptr
	addr_QFont_new9 uintptr
	addr_QFont_new10 uintptr
	addr_QFont_new11 uintptr
	addr_QFont_swap uintptr
	addr_QFont_family uintptr
	addr_QFont_setFamily uintptr
	addr_QFont_families uintptr
	addr_QFont_setFamilies uintptr
	addr_QFont_styleName uintptr
	addr_QFont_setStyleName uintptr
	addr_QFont_pointSize uintptr
	addr_QFont_setPointSize uintptr
	pg_QFont_pointSizeF func(uintptr) float64
	pg_QFont_setPointSizeF func(uintptr, float64)
	addr_QFont_pixelSize uintptr
	addr_QFont_setPixelSize uintptr
	addr_QFont_weight uintptr
	addr_QFont_setWeight uintptr
	addr_QFont_bold uintptr
	addr_QFont_setBold uintptr
	addr_QFont_setStyle uintptr
	addr_QFont_style uintptr
	addr_QFont_italic uintptr
	addr_QFont_setItalic uintptr
	addr_QFont_underline uintptr
	addr_QFont_setUnderline uintptr
	addr_QFont_overline uintptr
	addr_QFont_setOverline uintptr
	addr_QFont_strikeOut uintptr
	addr_QFont_setStrikeOut uintptr
	addr_QFont_fixedPitch uintptr
	addr_QFont_setFixedPitch uintptr
	addr_QFont_kerning uintptr
	addr_QFont_setKerning uintptr
	addr_QFont_styleHint uintptr
	addr_QFont_styleStrategy uintptr
	addr_QFont_setStyleHint uintptr
	addr_QFont_setStyleStrategy uintptr
	addr_QFont_stretch uintptr
	addr_QFont_setStretch uintptr
	pg_QFont_letterSpacing func(uintptr) float64
	addr_QFont_letterSpacingType uintptr
	pg_QFont_setLetterSpacing func(uintptr, int32, float64)
	pg_QFont_wordSpacing func(uintptr) float64
	pg_QFont_setWordSpacing func(uintptr, float64)
	addr_QFont_setCapitalization uintptr
	addr_QFont_capitalization uintptr
	addr_QFont_setHintingPreference uintptr
	addr_QFont_hintingPreference uintptr
	addr_QFont_setFeature uintptr
	addr_QFont_unsetFeature uintptr
	addr_QFont_featureValue uintptr
	addr_QFont_isFeatureSet uintptr
	addr_QFont_featureTags uintptr
	addr_QFont_clearFeatures uintptr
	pg_QFont_setVariableAxis func(uintptr, uintptr, float32)
	addr_QFont_unsetVariableAxis uintptr
	addr_QFont_isVariableAxisSet uintptr
	pg_QFont_variableAxisValue func(uintptr, uintptr) float32
	addr_QFont_clearVariableAxes uintptr
	addr_QFont_variableAxisTags uintptr
	addr_QFont_exactMatch uintptr
	addr_QFont_OperatorAssign uintptr
	addr_QFont_OperatorEqual uintptr
	addr_QFont_OperatorNotEqual uintptr
	addr_QFont_OperatorLesser uintptr
	addr_QFont_ToQVariant uintptr
	addr_QFont_isCopyOf uintptr
	addr_QFont_key uintptr
	addr_QFont_toString uintptr
	addr_QFont_fromString uintptr
	addr_QFont_substitute uintptr
	addr_QFont_substitutes uintptr
	addr_QFont_substitutions uintptr
	addr_QFont_insertSubstitution uintptr
	addr_QFont_insertSubstitutions uintptr
	addr_QFont_removeSubstitutions uintptr
	addr_QFont_initialize uintptr
	addr_QFont_cleanup uintptr
	addr_QFont_cacheStatistics uintptr
	addr_QFont_defaultFamily uintptr
	addr_QFont_resolve uintptr
	addr_QFont_resolveMask uintptr
	addr_QFont_setResolveMask uintptr
	addr_QFont_setLegacyWeight uintptr
	addr_QFont_legacyWeight uintptr
	addr_QFont_setStyleHint2 uintptr
	addr_QFont_toQVariant uintptr
	pg_QFont_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFont_new, "QFont_new")
	symaddr(&addr_QFont_new2, "QFont_new2")
	symaddr(&addr_QFont_new3, "QFont_new3")
	symaddr(&addr_QFont_new4, "QFont_new4")
	symaddr(&addr_QFont_new5, "QFont_new5")
	symaddr(&addr_QFont_new6, "QFont_new6")
	symaddr(&addr_QFont_new7, "QFont_new7")
	symaddr(&addr_QFont_new8, "QFont_new8")
	symaddr(&addr_QFont_new9, "QFont_new9")
	symaddr(&addr_QFont_new10, "QFont_new10")
	symaddr(&addr_QFont_new11, "QFont_new11")
	symaddr(&addr_QFont_swap, "QFont_swap")
	symaddr(&addr_QFont_family, "QFont_family")
	symaddr(&addr_QFont_setFamily, "QFont_setFamily")
	symaddr(&addr_QFont_families, "QFont_families")
	symaddr(&addr_QFont_setFamilies, "QFont_setFamilies")
	symaddr(&addr_QFont_styleName, "QFont_styleName")
	symaddr(&addr_QFont_setStyleName, "QFont_setStyleName")
	symaddr(&addr_QFont_pointSize, "QFont_pointSize")
	symaddr(&addr_QFont_setPointSize, "QFont_setPointSize")
	reg(&pg_QFont_pointSizeF, "QFont_pointSizeF")
	reg(&pg_QFont_setPointSizeF, "QFont_setPointSizeF")
	symaddr(&addr_QFont_pixelSize, "QFont_pixelSize")
	symaddr(&addr_QFont_setPixelSize, "QFont_setPixelSize")
	symaddr(&addr_QFont_weight, "QFont_weight")
	symaddr(&addr_QFont_setWeight, "QFont_setWeight")
	symaddr(&addr_QFont_bold, "QFont_bold")
	symaddr(&addr_QFont_setBold, "QFont_setBold")
	symaddr(&addr_QFont_setStyle, "QFont_setStyle")
	symaddr(&addr_QFont_style, "QFont_style")
	symaddr(&addr_QFont_italic, "QFont_italic")
	symaddr(&addr_QFont_setItalic, "QFont_setItalic")
	symaddr(&addr_QFont_underline, "QFont_underline")
	symaddr(&addr_QFont_setUnderline, "QFont_setUnderline")
	symaddr(&addr_QFont_overline, "QFont_overline")
	symaddr(&addr_QFont_setOverline, "QFont_setOverline")
	symaddr(&addr_QFont_strikeOut, "QFont_strikeOut")
	symaddr(&addr_QFont_setStrikeOut, "QFont_setStrikeOut")
	symaddr(&addr_QFont_fixedPitch, "QFont_fixedPitch")
	symaddr(&addr_QFont_setFixedPitch, "QFont_setFixedPitch")
	symaddr(&addr_QFont_kerning, "QFont_kerning")
	symaddr(&addr_QFont_setKerning, "QFont_setKerning")
	symaddr(&addr_QFont_styleHint, "QFont_styleHint")
	symaddr(&addr_QFont_styleStrategy, "QFont_styleStrategy")
	symaddr(&addr_QFont_setStyleHint, "QFont_setStyleHint")
	symaddr(&addr_QFont_setStyleStrategy, "QFont_setStyleStrategy")
	symaddr(&addr_QFont_stretch, "QFont_stretch")
	symaddr(&addr_QFont_setStretch, "QFont_setStretch")
	reg(&pg_QFont_letterSpacing, "QFont_letterSpacing")
	symaddr(&addr_QFont_letterSpacingType, "QFont_letterSpacingType")
	reg(&pg_QFont_setLetterSpacing, "QFont_setLetterSpacing")
	reg(&pg_QFont_wordSpacing, "QFont_wordSpacing")
	reg(&pg_QFont_setWordSpacing, "QFont_setWordSpacing")
	symaddr(&addr_QFont_setCapitalization, "QFont_setCapitalization")
	symaddr(&addr_QFont_capitalization, "QFont_capitalization")
	symaddr(&addr_QFont_setHintingPreference, "QFont_setHintingPreference")
	symaddr(&addr_QFont_hintingPreference, "QFont_hintingPreference")
	symaddr(&addr_QFont_setFeature, "QFont_setFeature")
	symaddr(&addr_QFont_unsetFeature, "QFont_unsetFeature")
	symaddr(&addr_QFont_featureValue, "QFont_featureValue")
	symaddr(&addr_QFont_isFeatureSet, "QFont_isFeatureSet")
	symaddr(&addr_QFont_featureTags, "QFont_featureTags")
	symaddr(&addr_QFont_clearFeatures, "QFont_clearFeatures")
	reg(&pg_QFont_setVariableAxis, "QFont_setVariableAxis")
	symaddr(&addr_QFont_unsetVariableAxis, "QFont_unsetVariableAxis")
	symaddr(&addr_QFont_isVariableAxisSet, "QFont_isVariableAxisSet")
	reg(&pg_QFont_variableAxisValue, "QFont_variableAxisValue")
	symaddr(&addr_QFont_clearVariableAxes, "QFont_clearVariableAxes")
	symaddr(&addr_QFont_variableAxisTags, "QFont_variableAxisTags")
	symaddr(&addr_QFont_exactMatch, "QFont_exactMatch")
	symaddr(&addr_QFont_OperatorAssign, "QFont_OperatorAssign")
	symaddr(&addr_QFont_OperatorEqual, "QFont_OperatorEqual")
	symaddr(&addr_QFont_OperatorNotEqual, "QFont_OperatorNotEqual")
	symaddr(&addr_QFont_OperatorLesser, "QFont_OperatorLesser")
	symaddr(&addr_QFont_ToQVariant, "QFont_ToQVariant")
	symaddr(&addr_QFont_isCopyOf, "QFont_isCopyOf")
	symaddr(&addr_QFont_key, "QFont_key")
	symaddr(&addr_QFont_toString, "QFont_toString")
	symaddr(&addr_QFont_fromString, "QFont_fromString")
	symaddr(&addr_QFont_substitute, "QFont_substitute")
	symaddr(&addr_QFont_substitutes, "QFont_substitutes")
	symaddr(&addr_QFont_substitutions, "QFont_substitutions")
	symaddr(&addr_QFont_insertSubstitution, "QFont_insertSubstitution")
	symaddr(&addr_QFont_insertSubstitutions, "QFont_insertSubstitutions")
	symaddr(&addr_QFont_removeSubstitutions, "QFont_removeSubstitutions")
	symaddr(&addr_QFont_initialize, "QFont_initialize")
	symaddr(&addr_QFont_cleanup, "QFont_cleanup")
	symaddr(&addr_QFont_cacheStatistics, "QFont_cacheStatistics")
	symaddr(&addr_QFont_defaultFamily, "QFont_defaultFamily")
	symaddr(&addr_QFont_resolve, "QFont_resolve")
	symaddr(&addr_QFont_resolveMask, "QFont_resolveMask")
	symaddr(&addr_QFont_setResolveMask, "QFont_setResolveMask")
	symaddr(&addr_QFont_setLegacyWeight, "QFont_setLegacyWeight")
	symaddr(&addr_QFont_legacyWeight, "QFont_legacyWeight")
	symaddr(&addr_QFont_setStyleHint2, "QFont_setStyleHint2")
	symaddr(&addr_QFont_toQVariant, "QFont_toQVariant")
	reg(&pg_QFont_delete, "QFont_delete")
}

type QFont struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFont) UnsafeGoID() int64     { return this.vid }
func (this *QFont) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFont(h uintptr) *QFont {
	if h == 0 {
		return nil
	}
	return &QFont{h: h}
}

func newownedQFont(h uintptr) *QFont {
	o := bindQFont(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFont).Delete)
	}
	return o
}

func (this *QFont) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFont) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFont(h unsafe.Pointer) *QFont { return bindQFont(uintptr(h)) }
func UnsafeNewQFontOwned(h unsafe.Pointer) *QFont { return newownedQFont(uintptr(h)) }

func (this *QFont) UnsafeRelease() { this.h = 0 }

func NewQFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QFont_new)
	__o := newownedQFont(__r)
	return __o
}

func NewQFont2(family string) *QFont {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new2, family_p, uintptr(len(family)))
	__o := newownedQFont(__r)
	runtime.KeepAlive(family)
	return __o
}

func NewQFont3(families []string) *QFont {
	families_arr := make([]uintptr, len(families))
	for __i, __e := range families {
		families_arr[__i] = pgCString(__e)
	}
	var families_p uintptr
	if len(families_arr) > 0 {
		families_p = uintptr(unsafe.Pointer(&families_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new3, families_p, uintptr(len(families)))
	__o := newownedQFont(__r)
	for __i := range families_arr {
		pgFree(families_arr[__i])
	}
	runtime.KeepAlive(families_arr)
	return __o
}

func NewQFont4(font *QFont, pd *QPaintDevice) *QFont {
	__r, _, _ := purego.SyscallN(addr_QFont_new4, uintptr(font.UnsafePointer()), uintptr(pd.UnsafePointer()))
	__o := newownedQFont(__r)
	runtime.KeepAlive(font)
	runtime.KeepAlive(pd)
	return __o
}

func NewQFont5(font *QFont) *QFont {
	__r, _, _ := purego.SyscallN(addr_QFont_new5, uintptr(font.UnsafePointer()))
	__o := newownedQFont(__r)
	runtime.KeepAlive(font)
	return __o
}

func NewQFont6(family string, pointSize int) *QFont {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new6, family_p, uintptr(len(family)), uintptr(pointSize))
	__o := newownedQFont(__r)
	runtime.KeepAlive(family)
	return __o
}

func NewQFont7(family string, pointSize int, weight int) *QFont {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new7, family_p, uintptr(len(family)), uintptr(pointSize), uintptr(weight))
	__o := newownedQFont(__r)
	runtime.KeepAlive(family)
	return __o
}

func NewQFont8(family string, pointSize int, weight int, italic bool) *QFont {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new8, family_p, uintptr(len(family)), uintptr(pointSize), uintptr(weight), uintptr(b2i(italic)))
	__o := newownedQFont(__r)
	runtime.KeepAlive(family)
	return __o
}

func NewQFont9(families []string, pointSize int) *QFont {
	families_arr := make([]uintptr, len(families))
	for __i, __e := range families {
		families_arr[__i] = pgCString(__e)
	}
	var families_p uintptr
	if len(families_arr) > 0 {
		families_p = uintptr(unsafe.Pointer(&families_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new9, families_p, uintptr(len(families)), uintptr(pointSize))
	__o := newownedQFont(__r)
	for __i := range families_arr {
		pgFree(families_arr[__i])
	}
	runtime.KeepAlive(families_arr)
	return __o
}

func NewQFont10(families []string, pointSize int, weight int) *QFont {
	families_arr := make([]uintptr, len(families))
	for __i, __e := range families {
		families_arr[__i] = pgCString(__e)
	}
	var families_p uintptr
	if len(families_arr) > 0 {
		families_p = uintptr(unsafe.Pointer(&families_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new10, families_p, uintptr(len(families)), uintptr(pointSize), uintptr(weight))
	__o := newownedQFont(__r)
	for __i := range families_arr {
		pgFree(families_arr[__i])
	}
	runtime.KeepAlive(families_arr)
	return __o
}

func NewQFont11(families []string, pointSize int, weight int, italic bool) *QFont {
	families_arr := make([]uintptr, len(families))
	for __i, __e := range families {
		families_arr[__i] = pgCString(__e)
	}
	var families_p uintptr
	if len(families_arr) > 0 {
		families_p = uintptr(unsafe.Pointer(&families_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFont_new11, families_p, uintptr(len(families)), uintptr(pointSize), uintptr(weight), uintptr(b2i(italic)))
	__o := newownedQFont(__r)
	for __i := range families_arr {
		pgFree(families_arr[__i])
	}
	runtime.KeepAlive(families_arr)
	return __o
}

func (this *QFont) Swap(other *QFont) {
	purego.SyscallN(addr_QFont_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFont) Family() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_family, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFont) SetFamily(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QFont_setFamily, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFont) Families() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_families, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFont) SetFamilies(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QFont_setFamilies, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func (this *QFont) StyleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_styleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFont) SetStyleName(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QFont_setStyleName, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFont) PointSize() int {
	__r0, _, _ := purego.SyscallN(addr_QFont_pointSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFont) SetPointSize(arg1 int) {
	purego.SyscallN(addr_QFont_setPointSize, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFont) PointSizeF() float64 {
	__r := pg_QFont_pointSizeF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFont) SetPointSizeF(arg1 float64) {
	pg_QFont_setPointSizeF(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QFont) PixelSize() int {
	__r0, _, _ := purego.SyscallN(addr_QFont_pixelSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFont) SetPixelSize(arg1 int) {
	purego.SyscallN(addr_QFont_setPixelSize, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFont) Weight() QFont__Weight {
	__r0, _, _ := purego.SyscallN(addr_QFont_weight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Weight(__r)
}

func (this *QFont) SetWeight(weight QFont__Weight) {
	purego.SyscallN(addr_QFont_setWeight, this.h, uintptr(weight))
	runtime.KeepAlive(this)
}

func (this *QFont) Bold() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_bold, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetBold(arg1 bool) {
	purego.SyscallN(addr_QFont_setBold, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) SetStyle(style QFont__Style) {
	purego.SyscallN(addr_QFont_setStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QFont) Style() QFont__Style {
	__r0, _, _ := purego.SyscallN(addr_QFont_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Style(__r)
}

func (this *QFont) Italic() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_italic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetItalic(b bool) {
	purego.SyscallN(addr_QFont_setItalic, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QFont) Underline() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_underline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetUnderline(arg1 bool) {
	purego.SyscallN(addr_QFont_setUnderline, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) Overline() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_overline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetOverline(arg1 bool) {
	purego.SyscallN(addr_QFont_setOverline, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) StrikeOut() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_strikeOut, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetStrikeOut(arg1 bool) {
	purego.SyscallN(addr_QFont_setStrikeOut, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) FixedPitch() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_fixedPitch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetFixedPitch(arg1 bool) {
	purego.SyscallN(addr_QFont_setFixedPitch, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) Kerning() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_kerning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) SetKerning(arg1 bool) {
	purego.SyscallN(addr_QFont_setKerning, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QFont) StyleHint() QFont__StyleHint {
	__r0, _, _ := purego.SyscallN(addr_QFont_styleHint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__StyleHint(__r)
}

func (this *QFont) StyleStrategy() QFont__StyleStrategy {
	__r0, _, _ := purego.SyscallN(addr_QFont_styleStrategy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__StyleStrategy(__r)
}

func (this *QFont) SetStyleHint(arg1 QFont__StyleHint) {
	purego.SyscallN(addr_QFont_setStyleHint, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFont) SetStyleStrategy(s QFont__StyleStrategy) {
	purego.SyscallN(addr_QFont_setStyleStrategy, this.h, uintptr(s))
	runtime.KeepAlive(this)
}

func (this *QFont) Stretch() int {
	__r0, _, _ := purego.SyscallN(addr_QFont_stretch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFont) SetStretch(arg1 int) {
	purego.SyscallN(addr_QFont_setStretch, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFont) LetterSpacing() float64 {
	__r := pg_QFont_letterSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFont) LetterSpacingType() QFont__SpacingType {
	__r0, _, _ := purego.SyscallN(addr_QFont_letterSpacingType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__SpacingType(__r)
}

func (this *QFont) SetLetterSpacing(type_ QFont__SpacingType, spacing float64) {
	pg_QFont_setLetterSpacing(this.h, int32(type_), spacing)
	runtime.KeepAlive(this)
}

func (this *QFont) WordSpacing() float64 {
	__r := pg_QFont_wordSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFont) SetWordSpacing(spacing float64) {
	pg_QFont_setWordSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QFont) SetCapitalization(arg1 QFont__Capitalization) {
	purego.SyscallN(addr_QFont_setCapitalization, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFont) Capitalization() QFont__Capitalization {
	__r0, _, _ := purego.SyscallN(addr_QFont_capitalization, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Capitalization(__r)
}

func (this *QFont) SetHintingPreference(hintingPreference QFont__HintingPreference) {
	purego.SyscallN(addr_QFont_setHintingPreference, this.h, uintptr(hintingPreference))
	runtime.KeepAlive(this)
}

func (this *QFont) HintingPreference() QFont__HintingPreference {
	__r0, _, _ := purego.SyscallN(addr_QFont_hintingPreference, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__HintingPreference(__r)
}

func (this *QFont) SetFeature(tag *QFont__Tag, value uint) {
	purego.SyscallN(addr_QFont_setFeature, this.h, uintptr(tag.UnsafePointer()), uintptr(value))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
}

func (this *QFont) UnsetFeature(tag *QFont__Tag) {
	purego.SyscallN(addr_QFont_unsetFeature, this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
}

func (this *QFont) FeatureValue(tag *QFont__Tag) uint {
	__r0, _, _ := purego.SyscallN(addr_QFont_featureValue, this.h, uintptr(tag.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return uint(__r)
}

func (this *QFont) IsFeatureSet(tag *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_isFeatureSet, this.h, uintptr(tag.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return __r != 0
}

func (this *QFont) FeatureTags() []*QFont__Tag {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_featureTags, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFont__Tag, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFont__Tag(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFont) ClearFeatures() {
	purego.SyscallN(addr_QFont_clearFeatures, this.h)
	runtime.KeepAlive(this)
}

func (this *QFont) SetVariableAxis(tag *QFont__Tag, value float32) {
	pg_QFont_setVariableAxis(this.h, uintptr(tag.UnsafePointer()), value)
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
}

func (this *QFont) UnsetVariableAxis(tag *QFont__Tag) {
	purego.SyscallN(addr_QFont_unsetVariableAxis, this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
}

func (this *QFont) IsVariableAxisSet(tag *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_isVariableAxisSet, this.h, uintptr(tag.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return __r != 0
}

func (this *QFont) VariableAxisValue(tag *QFont__Tag) float32 {
	__r := pg_QFont_variableAxisValue(this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return __r
}

func (this *QFont) ClearVariableAxes() {
	purego.SyscallN(addr_QFont_clearVariableAxes, this.h)
	runtime.KeepAlive(this)
}

func (this *QFont) VariableAxisTags() []*QFont__Tag {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_variableAxisTags, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFont__Tag, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFont__Tag(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFont) ExactMatch() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_exactMatch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont) OperatorAssign(arg1 *QFont) *QFont {
	__r, _, _ := purego.SyscallN(addr_QFont_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFont(__r)
}

func (this *QFont) OperatorEqual(arg1 *QFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFont) OperatorNotEqual(arg1 *QFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFont) OperatorLesser(arg1 *QFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_OperatorLesser, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFont) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFont_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFont) IsCopyOf(arg1 *QFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont_isCopyOf, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QFont) Key() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_key, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFont) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFont) FromString(arg1 string) bool {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFont_fromString, this.h, arg1_p, uintptr(len(arg1)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func QFont_Substitute(arg1 string) string {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_substitute, arg1_p, uintptr(len(arg1)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func QFont_Substitutes(arg1 string) []string {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_substitutes, arg1_p, uintptr(len(arg1)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(arg1)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFont_Substitutions() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_substitutions, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFont_InsertSubstitution(arg1 string, arg2 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	var arg2_p uintptr
	if len(arg2) > 0 {
		arg2_p = uintptr(unsafe.Pointer(unsafe.StringData(arg2)))
	}
	purego.SyscallN(addr_QFont_insertSubstitution, arg1_p, uintptr(len(arg1)), arg2_p, uintptr(len(arg2)))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFont_InsertSubstitutions(arg1 string, arg2 []string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	arg2_arr := make([]uintptr, len(arg2))
	for __i, __e := range arg2 {
		arg2_arr[__i] = pgCString(__e)
	}
	var arg2_p uintptr
	if len(arg2_arr) > 0 {
		arg2_p = uintptr(unsafe.Pointer(&arg2_arr[0]))
	}
	purego.SyscallN(addr_QFont_insertSubstitutions, arg1_p, uintptr(len(arg1)), arg2_p, uintptr(len(arg2)))
	runtime.KeepAlive(arg1)
	for __i := range arg2_arr {
		pgFree(arg2_arr[__i])
	}
	runtime.KeepAlive(arg2_arr)
}

func QFont_RemoveSubstitutions(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QFont_removeSubstitutions, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QFont_Initialize() {
	purego.SyscallN(addr_QFont_initialize)
}

func QFont_Cleanup() {
	purego.SyscallN(addr_QFont_cleanup)
}

func QFont_CacheStatistics() {
	purego.SyscallN(addr_QFont_cacheStatistics)
}

func (this *QFont) DefaultFamily() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFont_defaultFamily, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFont) Resolve(arg1 *QFont) *QFont {
	__r, _, _ := purego.SyscallN(addr_QFont_resolve, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQFont(__r)
}

func (this *QFont) ResolveMask() uint {
	__r0, _, _ := purego.SyscallN(addr_QFont_resolveMask, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QFont) SetResolveMask(mask uint) {
	purego.SyscallN(addr_QFont_setResolveMask, this.h, uintptr(mask))
	runtime.KeepAlive(this)
}

func (this *QFont) SetLegacyWeight(legacyWeight int) {
	purego.SyscallN(addr_QFont_setLegacyWeight, this.h, uintptr(legacyWeight))
	runtime.KeepAlive(this)
}

func (this *QFont) LegacyWeight() int {
	__r0, _, _ := purego.SyscallN(addr_QFont_legacyWeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFont) SetStyleHint2(arg1 QFont__StyleHint, arg2 QFont__StyleStrategy) {
	purego.SyscallN(addr_QFont_setStyleHint2, this.h, uintptr(arg1), uintptr(arg2))
	runtime.KeepAlive(this)
}

func (this *QFont) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFont_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFont) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFont_delete(this.h)
	}
	this.UnsafeRelease()
}
