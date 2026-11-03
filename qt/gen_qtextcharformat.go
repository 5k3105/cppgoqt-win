package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextCharFormat_base func(uintptr) uintptr
	addr_QTextCharFormat_new uintptr
	addr_QTextCharFormat_new2 uintptr
	addr_QTextCharFormat_isValid uintptr
	addr_QTextCharFormat_setFont uintptr
	addr_QTextCharFormat_font uintptr
	addr_QTextCharFormat_setFontFamily uintptr
	addr_QTextCharFormat_fontFamily uintptr
	addr_QTextCharFormat_setFontFamilies uintptr
	addr_QTextCharFormat_fontFamilies uintptr
	addr_QTextCharFormat_setFontStyleName uintptr
	addr_QTextCharFormat_fontStyleName uintptr
	pg_QTextCharFormat_setFontPointSize func(uintptr, float64)
	pg_QTextCharFormat_fontPointSize func(uintptr) float64
	addr_QTextCharFormat_setFontWeight uintptr
	addr_QTextCharFormat_fontWeight uintptr
	addr_QTextCharFormat_setFontItalic uintptr
	addr_QTextCharFormat_fontItalic uintptr
	addr_QTextCharFormat_setFontCapitalization uintptr
	addr_QTextCharFormat_fontCapitalization uintptr
	addr_QTextCharFormat_setFontLetterSpacingType uintptr
	addr_QTextCharFormat_fontLetterSpacingType uintptr
	pg_QTextCharFormat_setFontLetterSpacing func(uintptr, float64)
	pg_QTextCharFormat_fontLetterSpacing func(uintptr) float64
	pg_QTextCharFormat_setFontWordSpacing func(uintptr, float64)
	pg_QTextCharFormat_fontWordSpacing func(uintptr) float64
	addr_QTextCharFormat_setFontUnderline uintptr
	addr_QTextCharFormat_fontUnderline uintptr
	addr_QTextCharFormat_setFontOverline uintptr
	addr_QTextCharFormat_fontOverline uintptr
	addr_QTextCharFormat_setFontStrikeOut uintptr
	addr_QTextCharFormat_fontStrikeOut uintptr
	addr_QTextCharFormat_setUnderlineColor uintptr
	addr_QTextCharFormat_underlineColor uintptr
	addr_QTextCharFormat_setFontFixedPitch uintptr
	addr_QTextCharFormat_fontFixedPitch uintptr
	addr_QTextCharFormat_setFontStretch uintptr
	addr_QTextCharFormat_fontStretch uintptr
	addr_QTextCharFormat_setFontStyleHint uintptr
	addr_QTextCharFormat_setFontStyleStrategy uintptr
	addr_QTextCharFormat_fontStyleHint uintptr
	addr_QTextCharFormat_fontStyleStrategy uintptr
	addr_QTextCharFormat_setFontHintingPreference uintptr
	addr_QTextCharFormat_fontHintingPreference uintptr
	addr_QTextCharFormat_setFontKerning uintptr
	addr_QTextCharFormat_fontKerning uintptr
	addr_QTextCharFormat_setUnderlineStyle uintptr
	addr_QTextCharFormat_underlineStyle uintptr
	addr_QTextCharFormat_setVerticalAlignment uintptr
	addr_QTextCharFormat_verticalAlignment uintptr
	addr_QTextCharFormat_setTextOutline uintptr
	addr_QTextCharFormat_textOutline uintptr
	addr_QTextCharFormat_setToolTip uintptr
	addr_QTextCharFormat_toolTip uintptr
	pg_QTextCharFormat_setSuperScriptBaseline func(uintptr, float64)
	pg_QTextCharFormat_superScriptBaseline func(uintptr) float64
	pg_QTextCharFormat_setSubScriptBaseline func(uintptr, float64)
	pg_QTextCharFormat_subScriptBaseline func(uintptr) float64
	pg_QTextCharFormat_setBaselineOffset func(uintptr, float64)
	pg_QTextCharFormat_baselineOffset func(uintptr) float64
	addr_QTextCharFormat_setAnchor uintptr
	addr_QTextCharFormat_isAnchor uintptr
	addr_QTextCharFormat_setAnchorHref uintptr
	addr_QTextCharFormat_anchorHref uintptr
	addr_QTextCharFormat_setAnchorNames uintptr
	addr_QTextCharFormat_anchorNames uintptr
	addr_QTextCharFormat_setTableCellRowSpan uintptr
	addr_QTextCharFormat_tableCellRowSpan uintptr
	addr_QTextCharFormat_setTableCellColumnSpan uintptr
	addr_QTextCharFormat_tableCellColumnSpan uintptr
	addr_QTextCharFormat_OperatorAssign uintptr
	addr_QTextCharFormat_setFont2 uintptr
	addr_QTextCharFormat_setFontStyleHint2 uintptr
	pg_QTextCharFormat_delete func(uintptr)
)

func init() {
	reg(&pg_QTextCharFormat_base, "QTextCharFormat_base")
	symaddr(&addr_QTextCharFormat_new, "QTextCharFormat_new")
	symaddr(&addr_QTextCharFormat_new2, "QTextCharFormat_new2")
	symaddr(&addr_QTextCharFormat_isValid, "QTextCharFormat_isValid")
	symaddr(&addr_QTextCharFormat_setFont, "QTextCharFormat_setFont")
	symaddr(&addr_QTextCharFormat_font, "QTextCharFormat_font")
	symaddr(&addr_QTextCharFormat_setFontFamily, "QTextCharFormat_setFontFamily")
	symaddr(&addr_QTextCharFormat_fontFamily, "QTextCharFormat_fontFamily")
	symaddr(&addr_QTextCharFormat_setFontFamilies, "QTextCharFormat_setFontFamilies")
	symaddr(&addr_QTextCharFormat_fontFamilies, "QTextCharFormat_fontFamilies")
	symaddr(&addr_QTextCharFormat_setFontStyleName, "QTextCharFormat_setFontStyleName")
	symaddr(&addr_QTextCharFormat_fontStyleName, "QTextCharFormat_fontStyleName")
	reg(&pg_QTextCharFormat_setFontPointSize, "QTextCharFormat_setFontPointSize")
	reg(&pg_QTextCharFormat_fontPointSize, "QTextCharFormat_fontPointSize")
	symaddr(&addr_QTextCharFormat_setFontWeight, "QTextCharFormat_setFontWeight")
	symaddr(&addr_QTextCharFormat_fontWeight, "QTextCharFormat_fontWeight")
	symaddr(&addr_QTextCharFormat_setFontItalic, "QTextCharFormat_setFontItalic")
	symaddr(&addr_QTextCharFormat_fontItalic, "QTextCharFormat_fontItalic")
	symaddr(&addr_QTextCharFormat_setFontCapitalization, "QTextCharFormat_setFontCapitalization")
	symaddr(&addr_QTextCharFormat_fontCapitalization, "QTextCharFormat_fontCapitalization")
	symaddr(&addr_QTextCharFormat_setFontLetterSpacingType, "QTextCharFormat_setFontLetterSpacingType")
	symaddr(&addr_QTextCharFormat_fontLetterSpacingType, "QTextCharFormat_fontLetterSpacingType")
	reg(&pg_QTextCharFormat_setFontLetterSpacing, "QTextCharFormat_setFontLetterSpacing")
	reg(&pg_QTextCharFormat_fontLetterSpacing, "QTextCharFormat_fontLetterSpacing")
	reg(&pg_QTextCharFormat_setFontWordSpacing, "QTextCharFormat_setFontWordSpacing")
	reg(&pg_QTextCharFormat_fontWordSpacing, "QTextCharFormat_fontWordSpacing")
	symaddr(&addr_QTextCharFormat_setFontUnderline, "QTextCharFormat_setFontUnderline")
	symaddr(&addr_QTextCharFormat_fontUnderline, "QTextCharFormat_fontUnderline")
	symaddr(&addr_QTextCharFormat_setFontOverline, "QTextCharFormat_setFontOverline")
	symaddr(&addr_QTextCharFormat_fontOverline, "QTextCharFormat_fontOverline")
	symaddr(&addr_QTextCharFormat_setFontStrikeOut, "QTextCharFormat_setFontStrikeOut")
	symaddr(&addr_QTextCharFormat_fontStrikeOut, "QTextCharFormat_fontStrikeOut")
	symaddr(&addr_QTextCharFormat_setUnderlineColor, "QTextCharFormat_setUnderlineColor")
	symaddr(&addr_QTextCharFormat_underlineColor, "QTextCharFormat_underlineColor")
	symaddr(&addr_QTextCharFormat_setFontFixedPitch, "QTextCharFormat_setFontFixedPitch")
	symaddr(&addr_QTextCharFormat_fontFixedPitch, "QTextCharFormat_fontFixedPitch")
	symaddr(&addr_QTextCharFormat_setFontStretch, "QTextCharFormat_setFontStretch")
	symaddr(&addr_QTextCharFormat_fontStretch, "QTextCharFormat_fontStretch")
	symaddr(&addr_QTextCharFormat_setFontStyleHint, "QTextCharFormat_setFontStyleHint")
	symaddr(&addr_QTextCharFormat_setFontStyleStrategy, "QTextCharFormat_setFontStyleStrategy")
	symaddr(&addr_QTextCharFormat_fontStyleHint, "QTextCharFormat_fontStyleHint")
	symaddr(&addr_QTextCharFormat_fontStyleStrategy, "QTextCharFormat_fontStyleStrategy")
	symaddr(&addr_QTextCharFormat_setFontHintingPreference, "QTextCharFormat_setFontHintingPreference")
	symaddr(&addr_QTextCharFormat_fontHintingPreference, "QTextCharFormat_fontHintingPreference")
	symaddr(&addr_QTextCharFormat_setFontKerning, "QTextCharFormat_setFontKerning")
	symaddr(&addr_QTextCharFormat_fontKerning, "QTextCharFormat_fontKerning")
	symaddr(&addr_QTextCharFormat_setUnderlineStyle, "QTextCharFormat_setUnderlineStyle")
	symaddr(&addr_QTextCharFormat_underlineStyle, "QTextCharFormat_underlineStyle")
	symaddr(&addr_QTextCharFormat_setVerticalAlignment, "QTextCharFormat_setVerticalAlignment")
	symaddr(&addr_QTextCharFormat_verticalAlignment, "QTextCharFormat_verticalAlignment")
	symaddr(&addr_QTextCharFormat_setTextOutline, "QTextCharFormat_setTextOutline")
	symaddr(&addr_QTextCharFormat_textOutline, "QTextCharFormat_textOutline")
	symaddr(&addr_QTextCharFormat_setToolTip, "QTextCharFormat_setToolTip")
	symaddr(&addr_QTextCharFormat_toolTip, "QTextCharFormat_toolTip")
	reg(&pg_QTextCharFormat_setSuperScriptBaseline, "QTextCharFormat_setSuperScriptBaseline")
	reg(&pg_QTextCharFormat_superScriptBaseline, "QTextCharFormat_superScriptBaseline")
	reg(&pg_QTextCharFormat_setSubScriptBaseline, "QTextCharFormat_setSubScriptBaseline")
	reg(&pg_QTextCharFormat_subScriptBaseline, "QTextCharFormat_subScriptBaseline")
	reg(&pg_QTextCharFormat_setBaselineOffset, "QTextCharFormat_setBaselineOffset")
	reg(&pg_QTextCharFormat_baselineOffset, "QTextCharFormat_baselineOffset")
	symaddr(&addr_QTextCharFormat_setAnchor, "QTextCharFormat_setAnchor")
	symaddr(&addr_QTextCharFormat_isAnchor, "QTextCharFormat_isAnchor")
	symaddr(&addr_QTextCharFormat_setAnchorHref, "QTextCharFormat_setAnchorHref")
	symaddr(&addr_QTextCharFormat_anchorHref, "QTextCharFormat_anchorHref")
	symaddr(&addr_QTextCharFormat_setAnchorNames, "QTextCharFormat_setAnchorNames")
	symaddr(&addr_QTextCharFormat_anchorNames, "QTextCharFormat_anchorNames")
	symaddr(&addr_QTextCharFormat_setTableCellRowSpan, "QTextCharFormat_setTableCellRowSpan")
	symaddr(&addr_QTextCharFormat_tableCellRowSpan, "QTextCharFormat_tableCellRowSpan")
	symaddr(&addr_QTextCharFormat_setTableCellColumnSpan, "QTextCharFormat_setTableCellColumnSpan")
	symaddr(&addr_QTextCharFormat_tableCellColumnSpan, "QTextCharFormat_tableCellColumnSpan")
	symaddr(&addr_QTextCharFormat_OperatorAssign, "QTextCharFormat_OperatorAssign")
	symaddr(&addr_QTextCharFormat_setFont2, "QTextCharFormat_setFont2")
	symaddr(&addr_QTextCharFormat_setFontStyleHint2, "QTextCharFormat_setFontStyleHint2")
	reg(&pg_QTextCharFormat_delete, "QTextCharFormat_delete")
}

type QTextCharFormat struct {
	h uintptr
	*QTextFormat
}

func bindQTextCharFormat(h uintptr) *QTextCharFormat {
	if h == 0 {
		return nil
	}
	return &QTextCharFormat{h: h, QTextFormat: bindQTextFormat(pg_QTextCharFormat_base(h))}
}

func newownedQTextCharFormat(h uintptr) *QTextCharFormat {
	o := bindQTextCharFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextCharFormat).Delete)
	}
	return o
}

func (this *QTextCharFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextCharFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextCharFormat(h unsafe.Pointer) *QTextCharFormat { return bindQTextCharFormat(uintptr(h)) }
func UnsafeNewQTextCharFormatOwned(h unsafe.Pointer) *QTextCharFormat { return newownedQTextCharFormat(uintptr(h)) }

func (this *QTextCharFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextFormat != nil {
		this.QTextFormat.UnsafeRelease()
	}
}

func NewQTextCharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_new)
	__o := newownedQTextCharFormat(__r)
	return __o
}

func NewQTextCharFormat2(arg1 *QTextCharFormat) *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextCharFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextCharFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetFont(font *QFont) {
	purego.SyscallN(addr_QTextCharFormat_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QTextCharFormat) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTextCharFormat) SetFontFamily(family string) {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	purego.SyscallN(addr_QTextCharFormat_setFontFamily, this.h, family_p, uintptr(len(family)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(family)
}

func (this *QTextCharFormat) FontFamily() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_fontFamily, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextCharFormat) SetFontFamilies(families []string) {
	families_arr := make([]uintptr, len(families))
	for __i, __e := range families {
		families_arr[__i] = pgCString(__e)
	}
	var families_p uintptr
	if len(families_arr) > 0 {
		families_p = uintptr(unsafe.Pointer(&families_arr[0]))
	}
	purego.SyscallN(addr_QTextCharFormat_setFontFamilies, this.h, families_p, uintptr(len(families)))
	runtime.KeepAlive(this)
	for __i := range families_arr {
		pgFree(families_arr[__i])
	}
	runtime.KeepAlive(families_arr)
}

func (this *QTextCharFormat) FontFamilies() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_fontFamilies, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextCharFormat) SetFontStyleName(styleName string) {
	var styleName_p uintptr
	if len(styleName) > 0 {
		styleName_p = uintptr(unsafe.Pointer(unsafe.StringData(styleName)))
	}
	purego.SyscallN(addr_QTextCharFormat_setFontStyleName, this.h, styleName_p, uintptr(len(styleName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(styleName)
}

func (this *QTextCharFormat) FontStyleName() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_fontStyleName, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextCharFormat) SetFontPointSize(size float64) {
	pg_QTextCharFormat_setFontPointSize(this.h, size)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontPointSize() float64 {
	__r := pg_QTextCharFormat_fontPointSize(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetFontWeight(weight int) {
	purego.SyscallN(addr_QTextCharFormat_setFontWeight, this.h, uintptr(weight))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontWeight() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontWeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCharFormat) SetFontItalic(italic bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontItalic, this.h, uintptr(b2i(italic)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontItalic() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontItalic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetFontCapitalization(capitalization QFont__Capitalization) {
	purego.SyscallN(addr_QTextCharFormat_setFontCapitalization, this.h, uintptr(capitalization))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontCapitalization() QFont__Capitalization {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontCapitalization, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Capitalization(__r)
}

func (this *QTextCharFormat) SetFontLetterSpacingType(letterSpacingType QFont__SpacingType) {
	purego.SyscallN(addr_QTextCharFormat_setFontLetterSpacingType, this.h, uintptr(letterSpacingType))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontLetterSpacingType() QFont__SpacingType {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontLetterSpacingType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__SpacingType(__r)
}

func (this *QTextCharFormat) SetFontLetterSpacing(spacing float64) {
	pg_QTextCharFormat_setFontLetterSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontLetterSpacing() float64 {
	__r := pg_QTextCharFormat_fontLetterSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetFontWordSpacing(spacing float64) {
	pg_QTextCharFormat_setFontWordSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontWordSpacing() float64 {
	__r := pg_QTextCharFormat_fontWordSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetFontUnderline(underline bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontUnderline, this.h, uintptr(b2i(underline)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontUnderline() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontUnderline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetFontOverline(overline bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontOverline, this.h, uintptr(b2i(overline)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontOverline() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontOverline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetFontStrikeOut(strikeOut bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontStrikeOut, this.h, uintptr(b2i(strikeOut)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontStrikeOut() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontStrikeOut, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetUnderlineColor(color *QColor) {
	purego.SyscallN(addr_QTextCharFormat_setUnderlineColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QTextCharFormat) UnderlineColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_underlineColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QTextCharFormat) SetFontFixedPitch(fixedPitch bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontFixedPitch, this.h, uintptr(b2i(fixedPitch)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontFixedPitch() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontFixedPitch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetFontStretch(factor int) {
	purego.SyscallN(addr_QTextCharFormat_setFontStretch, this.h, uintptr(factor))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontStretch() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontStretch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCharFormat) SetFontStyleHint(hint QFont__StyleHint) {
	purego.SyscallN(addr_QTextCharFormat_setFontStyleHint, this.h, uintptr(hint))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) SetFontStyleStrategy(strategy QFont__StyleStrategy) {
	purego.SyscallN(addr_QTextCharFormat_setFontStyleStrategy, this.h, uintptr(strategy))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontStyleHint() QFont__StyleHint {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontStyleHint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__StyleHint(__r)
}

func (this *QTextCharFormat) FontStyleStrategy() QFont__StyleStrategy {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontStyleStrategy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__StyleStrategy(__r)
}

func (this *QTextCharFormat) SetFontHintingPreference(hintingPreference QFont__HintingPreference) {
	purego.SyscallN(addr_QTextCharFormat_setFontHintingPreference, this.h, uintptr(hintingPreference))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontHintingPreference() QFont__HintingPreference {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontHintingPreference, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__HintingPreference(__r)
}

func (this *QTextCharFormat) SetFontKerning(enable bool) {
	purego.SyscallN(addr_QTextCharFormat_setFontKerning, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) FontKerning() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_fontKerning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetUnderlineStyle(style QTextCharFormat__UnderlineStyle) {
	purego.SyscallN(addr_QTextCharFormat_setUnderlineStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) UnderlineStyle() QTextCharFormat__UnderlineStyle {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_underlineStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextCharFormat__UnderlineStyle(__r)
}

func (this *QTextCharFormat) SetVerticalAlignment(alignment QTextCharFormat__VerticalAlignment) {
	purego.SyscallN(addr_QTextCharFormat_setVerticalAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) VerticalAlignment() QTextCharFormat__VerticalAlignment {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_verticalAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextCharFormat__VerticalAlignment(__r)
}

func (this *QTextCharFormat) SetTextOutline(pen *QPen) {
	purego.SyscallN(addr_QTextCharFormat_setTextOutline, this.h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
}

func (this *QTextCharFormat) TextOutline() *QPen {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_textOutline, this.h)
	runtime.KeepAlive(this)
	return newownedQPen(__r)
}

func (this *QTextCharFormat) SetToolTip(tip string) {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	purego.SyscallN(addr_QTextCharFormat_setToolTip, this.h, tip_p, uintptr(len(tip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tip)
}

func (this *QTextCharFormat) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextCharFormat) SetSuperScriptBaseline(baseline float64) {
	pg_QTextCharFormat_setSuperScriptBaseline(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) SuperScriptBaseline() float64 {
	__r := pg_QTextCharFormat_superScriptBaseline(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetSubScriptBaseline(baseline float64) {
	pg_QTextCharFormat_setSubScriptBaseline(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) SubScriptBaseline() float64 {
	__r := pg_QTextCharFormat_subScriptBaseline(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetBaselineOffset(baseline float64) {
	pg_QTextCharFormat_setBaselineOffset(this.h, baseline)
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) BaselineOffset() float64 {
	__r := pg_QTextCharFormat_baselineOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextCharFormat) SetAnchor(anchor bool) {
	purego.SyscallN(addr_QTextCharFormat_setAnchor, this.h, uintptr(b2i(anchor)))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) IsAnchor() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_isAnchor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCharFormat) SetAnchorHref(value string) {
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QTextCharFormat_setAnchorHref, this.h, value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTextCharFormat) AnchorHref() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_anchorHref, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextCharFormat) SetAnchorNames(names []string) {
	names_arr := make([]uintptr, len(names))
	for __i, __e := range names {
		names_arr[__i] = pgCString(__e)
	}
	var names_p uintptr
	if len(names_arr) > 0 {
		names_p = uintptr(unsafe.Pointer(&names_arr[0]))
	}
	purego.SyscallN(addr_QTextCharFormat_setAnchorNames, this.h, names_p, uintptr(len(names)))
	runtime.KeepAlive(this)
	for __i := range names_arr {
		pgFree(names_arr[__i])
	}
	runtime.KeepAlive(names_arr)
}

func (this *QTextCharFormat) AnchorNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_anchorNames, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QTextCharFormat) SetTableCellRowSpan(tableCellRowSpan int) {
	purego.SyscallN(addr_QTextCharFormat_setTableCellRowSpan, this.h, uintptr(tableCellRowSpan))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) TableCellRowSpan() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_tableCellRowSpan, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCharFormat) SetTableCellColumnSpan(tableCellColumnSpan int) {
	purego.SyscallN(addr_QTextCharFormat_setTableCellColumnSpan, this.h, uintptr(tableCellColumnSpan))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) TableCellColumnSpan() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCharFormat_tableCellColumnSpan, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCharFormat) OperatorAssign(arg1 *QTextCharFormat) *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCharFormat_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextCharFormat(__r)
}

func (this *QTextCharFormat) SetFont2(font *QFont, behavior QTextCharFormat__FontPropertiesInheritanceBehavior) {
	purego.SyscallN(addr_QTextCharFormat_setFont2, this.h, uintptr(font.UnsafePointer()), uintptr(behavior))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QTextCharFormat) SetFontStyleHint2(hint QFont__StyleHint, strategy QFont__StyleStrategy) {
	purego.SyscallN(addr_QTextCharFormat_setFontStyleHint2, this.h, uintptr(hint), uintptr(strategy))
	runtime.KeepAlive(this)
}

func (this *QTextCharFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextCharFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
