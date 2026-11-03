package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRawFont_new uintptr
	pg_QRawFont_new2 func(string, int64, float64) uintptr
	pg_QRawFont_new3 func(uintptr, float64) uintptr
	addr_QRawFont_new4 uintptr
	pg_QRawFont_new5 func(string, int64, float64, int32) uintptr
	pg_QRawFont_new6 func(uintptr, float64, int32) uintptr
	addr_QRawFont_OperatorAssign uintptr
	addr_QRawFont_swap uintptr
	addr_QRawFont_isValid uintptr
	addr_QRawFont_OperatorEqual uintptr
	addr_QRawFont_OperatorNotEqual uintptr
	addr_QRawFont_glyphCount uintptr
	addr_QRawFont_familyName uintptr
	addr_QRawFont_styleName uintptr
	addr_QRawFont_style uintptr
	addr_QRawFont_weight uintptr
	addr_QRawFont_glyphIndexesForString uintptr
	addr_QRawFont_advancesForGlyphIndexes uintptr
	addr_QRawFont_advancesForGlyphIndexes2 uintptr
	addr_QRawFont_glyphIndexesForChars uintptr
	addr_QRawFont_alphaMapForGlyph uintptr
	addr_QRawFont_pathForGlyph uintptr
	addr_QRawFont_boundingRect uintptr
	pg_QRawFont_setPixelSize func(uintptr, float64)
	pg_QRawFont_pixelSize func(uintptr) float64
	addr_QRawFont_hintingPreference uintptr
	pg_QRawFont_ascent func(uintptr) float64
	pg_QRawFont_capHeight func(uintptr) float64
	pg_QRawFont_descent func(uintptr) float64
	pg_QRawFont_leading func(uintptr) float64
	pg_QRawFont_xHeight func(uintptr) float64
	pg_QRawFont_averageCharWidth func(uintptr) float64
	pg_QRawFont_maxCharWidth func(uintptr) float64
	pg_QRawFont_lineThickness func(uintptr) float64
	pg_QRawFont_underlinePosition func(uintptr) float64
	pg_QRawFont_unitsPerEm func(uintptr) float64
	addr_QRawFont_glyphName uintptr
	pg_QRawFont_loadFromFile func(uintptr, string, int64, float64, int32)
	pg_QRawFont_loadFromData func(uintptr, uintptr, float64, int32)
	addr_QRawFont_supportsCharacter uintptr
	addr_QRawFont_supportsCharacter2 uintptr
	addr_QRawFont_supportedWritingSystems uintptr
	addr_QRawFont_fontTable uintptr
	addr_QRawFont_fontTable2 uintptr
	addr_QRawFont_fromFont uintptr
	addr_QRawFont_alphaMapForGlyph2 uintptr
	addr_QRawFont_alphaMapForGlyph3 uintptr
	addr_QRawFont_fromFont2 uintptr
	pg_QRawFont_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRawFont_new, "QRawFont_new")
	reg(&pg_QRawFont_new2, "QRawFont_new2")
	reg(&pg_QRawFont_new3, "QRawFont_new3")
	symaddr(&addr_QRawFont_new4, "QRawFont_new4")
	reg(&pg_QRawFont_new5, "QRawFont_new5")
	reg(&pg_QRawFont_new6, "QRawFont_new6")
	symaddr(&addr_QRawFont_OperatorAssign, "QRawFont_OperatorAssign")
	symaddr(&addr_QRawFont_swap, "QRawFont_swap")
	symaddr(&addr_QRawFont_isValid, "QRawFont_isValid")
	symaddr(&addr_QRawFont_OperatorEqual, "QRawFont_OperatorEqual")
	symaddr(&addr_QRawFont_OperatorNotEqual, "QRawFont_OperatorNotEqual")
	symaddr(&addr_QRawFont_glyphCount, "QRawFont_glyphCount")
	symaddr(&addr_QRawFont_familyName, "QRawFont_familyName")
	symaddr(&addr_QRawFont_styleName, "QRawFont_styleName")
	symaddr(&addr_QRawFont_style, "QRawFont_style")
	symaddr(&addr_QRawFont_weight, "QRawFont_weight")
	symaddr(&addr_QRawFont_glyphIndexesForString, "QRawFont_glyphIndexesForString")
	symaddr(&addr_QRawFont_advancesForGlyphIndexes, "QRawFont_advancesForGlyphIndexes")
	symaddr(&addr_QRawFont_advancesForGlyphIndexes2, "QRawFont_advancesForGlyphIndexes2")
	symaddr(&addr_QRawFont_glyphIndexesForChars, "QRawFont_glyphIndexesForChars")
	symaddr(&addr_QRawFont_alphaMapForGlyph, "QRawFont_alphaMapForGlyph")
	symaddr(&addr_QRawFont_pathForGlyph, "QRawFont_pathForGlyph")
	symaddr(&addr_QRawFont_boundingRect, "QRawFont_boundingRect")
	reg(&pg_QRawFont_setPixelSize, "QRawFont_setPixelSize")
	reg(&pg_QRawFont_pixelSize, "QRawFont_pixelSize")
	symaddr(&addr_QRawFont_hintingPreference, "QRawFont_hintingPreference")
	reg(&pg_QRawFont_ascent, "QRawFont_ascent")
	reg(&pg_QRawFont_capHeight, "QRawFont_capHeight")
	reg(&pg_QRawFont_descent, "QRawFont_descent")
	reg(&pg_QRawFont_leading, "QRawFont_leading")
	reg(&pg_QRawFont_xHeight, "QRawFont_xHeight")
	reg(&pg_QRawFont_averageCharWidth, "QRawFont_averageCharWidth")
	reg(&pg_QRawFont_maxCharWidth, "QRawFont_maxCharWidth")
	reg(&pg_QRawFont_lineThickness, "QRawFont_lineThickness")
	reg(&pg_QRawFont_underlinePosition, "QRawFont_underlinePosition")
	reg(&pg_QRawFont_unitsPerEm, "QRawFont_unitsPerEm")
	symaddr(&addr_QRawFont_glyphName, "QRawFont_glyphName")
	reg(&pg_QRawFont_loadFromFile, "QRawFont_loadFromFile")
	reg(&pg_QRawFont_loadFromData, "QRawFont_loadFromData")
	symaddr(&addr_QRawFont_supportsCharacter, "QRawFont_supportsCharacter")
	symaddr(&addr_QRawFont_supportsCharacter2, "QRawFont_supportsCharacter2")
	symaddr(&addr_QRawFont_supportedWritingSystems, "QRawFont_supportedWritingSystems")
	symaddr(&addr_QRawFont_fontTable, "QRawFont_fontTable")
	symaddr(&addr_QRawFont_fontTable2, "QRawFont_fontTable2")
	symaddr(&addr_QRawFont_fromFont, "QRawFont_fromFont")
	symaddr(&addr_QRawFont_alphaMapForGlyph2, "QRawFont_alphaMapForGlyph2")
	symaddr(&addr_QRawFont_alphaMapForGlyph3, "QRawFont_alphaMapForGlyph3")
	symaddr(&addr_QRawFont_fromFont2, "QRawFont_fromFont2")
	reg(&pg_QRawFont_delete, "QRawFont_delete")
}

type QRawFont struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRawFont) UnsafeGoID() int64     { return this.vid }
func (this *QRawFont) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRawFont(h uintptr) *QRawFont {
	if h == 0 {
		return nil
	}
	return &QRawFont{h: h}
}

func newownedQRawFont(h uintptr) *QRawFont {
	o := bindQRawFont(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRawFont).Delete)
	}
	return o
}

func (this *QRawFont) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRawFont) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRawFont(h unsafe.Pointer) *QRawFont { return bindQRawFont(uintptr(h)) }
func UnsafeNewQRawFontOwned(h unsafe.Pointer) *QRawFont { return newownedQRawFont(uintptr(h)) }

func (this *QRawFont) UnsafeRelease() { this.h = 0 }

func NewQRawFont() *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QRawFont_new)
	__o := newownedQRawFont(__r)
	return __o
}

func NewQRawFont2(fileName string, pixelSize float64) *QRawFont {
	__o := newownedQRawFont(pg_QRawFont_new2(fileName, int64(len(fileName)), pixelSize))
	return __o
}

func NewQRawFont3(fontData *QByteArray, pixelSize float64) *QRawFont {
	__o := newownedQRawFont(pg_QRawFont_new3(uintptr(fontData.UnsafePointer()), pixelSize))
	runtime.KeepAlive(fontData)
	return __o
}

func NewQRawFont4(other *QRawFont) *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QRawFont_new4, uintptr(other.UnsafePointer()))
	__o := newownedQRawFont(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQRawFont5(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	__o := newownedQRawFont(pg_QRawFont_new5(fileName, int64(len(fileName)), pixelSize, int32(hintingPreference)))
	return __o
}

func NewQRawFont6(fontData *QByteArray, pixelSize float64, hintingPreference QFont__HintingPreference) *QRawFont {
	__o := newownedQRawFont(pg_QRawFont_new6(uintptr(fontData.UnsafePointer()), pixelSize, int32(hintingPreference)))
	runtime.KeepAlive(fontData)
	return __o
}

func (this *QRawFont) OperatorAssign(other *QRawFont) *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QRawFont_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQRawFont(__r)
}

func (this *QRawFont) Swap(other *QRawFont) {
	purego.SyscallN(addr_QRawFont_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QRawFont) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRawFont) OperatorEqual(other *QRawFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QRawFont) OperatorNotEqual(other *QRawFont) bool {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QRawFont) GlyphCount() uint {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_glyphCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRawFont) FamilyName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_familyName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRawFont) StyleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_styleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRawFont) Style() QFont__Style {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Style(__r)
}

func (this *QRawFont) Weight() int {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_weight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRawFont) GlyphIndexesForString(text string) []uint {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_glyphIndexesForString, this.h, text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]uint, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = uint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QRawFont) AdvancesForGlyphIndexes(glyphIndexes []uint) []*QPointF {
	glyphIndexes_arr := make([]int32, len(glyphIndexes))
	for __i, __e := range glyphIndexes {
		glyphIndexes_arr[__i] = int32(__e)
	}
	var glyphIndexes_p uintptr
	if len(glyphIndexes_arr) > 0 {
		glyphIndexes_p = uintptr(unsafe.Pointer(&glyphIndexes_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_advancesForGlyphIndexes, this.h, glyphIndexes_p, uintptr(len(glyphIndexes)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(glyphIndexes_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QRawFont) AdvancesForGlyphIndexes2(glyphIndexes []uint, layoutFlags QRawFont__LayoutFlag) []*QPointF {
	glyphIndexes_arr := make([]int32, len(glyphIndexes))
	for __i, __e := range glyphIndexes {
		glyphIndexes_arr[__i] = int32(__e)
	}
	var glyphIndexes_p uintptr
	if len(glyphIndexes_arr) > 0 {
		glyphIndexes_p = uintptr(unsafe.Pointer(&glyphIndexes_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_advancesForGlyphIndexes2, this.h, glyphIndexes_p, uintptr(len(glyphIndexes)), uintptr(layoutFlags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(glyphIndexes_arr)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QRawFont) GlyphIndexesForChars(chars *QChar, numChars int) (bool, uint, int) {
	var glyphIndexes_out int32
	var numGlyphs_out int32
	__r0, _, _ := purego.SyscallN(addr_QRawFont_glyphIndexesForChars, this.h, uintptr(chars.UnsafePointer()), uintptr(numChars), uintptr(unsafe.Pointer(&glyphIndexes_out)), uintptr(unsafe.Pointer(&numGlyphs_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(chars)
	return __r != 0, uint(glyphIndexes_out), int(numGlyphs_out)
}

func (this *QRawFont) AlphaMapForGlyph(glyphIndex uint) *QImage {
	__r, _, _ := purego.SyscallN(addr_QRawFont_alphaMapForGlyph, this.h, uintptr(glyphIndex))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QRawFont) PathForGlyph(glyphIndex uint) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QRawFont_pathForGlyph, this.h, uintptr(glyphIndex))
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QRawFont) BoundingRect(glyphIndex uint) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRawFont_boundingRect, this.h, uintptr(glyphIndex))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRawFont) SetPixelSize(pixelSize float64) {
	pg_QRawFont_setPixelSize(this.h, pixelSize)
	runtime.KeepAlive(this)
}

func (this *QRawFont) PixelSize() float64 {
	__r := pg_QRawFont_pixelSize(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) HintingPreference() QFont__HintingPreference {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_hintingPreference, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__HintingPreference(__r)
}

func (this *QRawFont) Ascent() float64 {
	__r := pg_QRawFont_ascent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) CapHeight() float64 {
	__r := pg_QRawFont_capHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) Descent() float64 {
	__r := pg_QRawFont_descent(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) Leading() float64 {
	__r := pg_QRawFont_leading(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) XHeight() float64 {
	__r := pg_QRawFont_xHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) AverageCharWidth() float64 {
	__r := pg_QRawFont_averageCharWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) MaxCharWidth() float64 {
	__r := pg_QRawFont_maxCharWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) LineThickness() float64 {
	__r := pg_QRawFont_lineThickness(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) UnderlinePosition() float64 {
	__r := pg_QRawFont_underlinePosition(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) UnitsPerEm() float64 {
	__r := pg_QRawFont_unitsPerEm(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRawFont) GlyphName(glyphIndex uint) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_glyphName, this.h, uintptr(glyphIndex), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QRawFont) LoadFromFile(fileName string, pixelSize float64, hintingPreference QFont__HintingPreference) {
	pg_QRawFont_loadFromFile(this.h, fileName, int64(len(fileName)), pixelSize, int32(hintingPreference))
	runtime.KeepAlive(this)
}

func (this *QRawFont) LoadFromData(fontData *QByteArray, pixelSize float64, hintingPreference QFont__HintingPreference) {
	pg_QRawFont_loadFromData(this.h, uintptr(fontData.UnsafePointer()), pixelSize, int32(hintingPreference))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontData)
}

func (this *QRawFont) SupportsCharacter(ucs4 uint) bool {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_supportsCharacter, this.h, uintptr(ucs4))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRawFont) SupportsCharacterWithCharacter(character *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QRawFont_supportsCharacter2, this.h, uintptr(character.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(character)
	return __r != 0
}

func (this *QRawFont) SupportedWritingSystems() []QFontDatabase__WritingSystem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRawFont_supportedWritingSystems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]QFontDatabase__WritingSystem, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QFontDatabase__WritingSystem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QRawFont) FontTable(tagName string) *QByteArray {
	tagName_b := append([]byte(tagName), 0)
	__r, _, _ := purego.SyscallN(addr_QRawFont_fontTable, this.h, uintptr(unsafe.Pointer(&tagName_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tagName_b)
	return newownedQByteArray(__r)
}

func (this *QRawFont) FontTableWithTag(tag *QFont__Tag) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QRawFont_fontTable2, this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
	return newownedQByteArray(__r)
}

func QRawFont_FromFont(font *QFont) *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QRawFont_fromFont, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(font)
	return newownedQRawFont(__r)
}

func (this *QRawFont) AlphaMapForGlyph2(glyphIndex uint, antialiasingType QRawFont__AntialiasingType) *QImage {
	__r, _, _ := purego.SyscallN(addr_QRawFont_alphaMapForGlyph2, this.h, uintptr(glyphIndex), uintptr(antialiasingType))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QRawFont) AlphaMapForGlyph3(glyphIndex uint, antialiasingType QRawFont__AntialiasingType, transform *QTransform) *QImage {
	__r, _, _ := purego.SyscallN(addr_QRawFont_alphaMapForGlyph3, this.h, uintptr(glyphIndex), uintptr(antialiasingType), uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func QRawFont_FromFont2(font *QFont, writingSystem QFontDatabase__WritingSystem) *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QRawFont_fromFont2, uintptr(font.UnsafePointer()), uintptr(writingSystem))
	runtime.KeepAlive(font)
	return newownedQRawFont(__r)
}

func (this *QRawFont) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRawFont_delete(this.h)
	}
	this.UnsafeRelease()
}
