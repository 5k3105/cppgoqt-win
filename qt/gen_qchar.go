package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QChar_new uintptr
	addr_QChar_new2 uintptr
	addr_QChar_new3 uintptr
	addr_QChar_fromUcs2 uintptr
	addr_QChar_category uintptr
	addr_QChar_direction uintptr
	addr_QChar_joiningType uintptr
	addr_QChar_combiningClass uintptr
	addr_QChar_mirroredChar uintptr
	addr_QChar_hasMirrored uintptr
	addr_QChar_decomposition uintptr
	addr_QChar_decompositionTag uintptr
	addr_QChar_digitValue uintptr
	addr_QChar_toLower uintptr
	addr_QChar_toUpper uintptr
	addr_QChar_toTitleCase uintptr
	addr_QChar_toCaseFolded uintptr
	addr_QChar_script uintptr
	addr_QChar_unicodeVersion uintptr
	addr_QChar_toLatin1 uintptr
	addr_QChar_unicode uintptr
	addr_QChar_fromLatin1 uintptr
	addr_QChar_isNull uintptr
	addr_QChar_isPrint uintptr
	addr_QChar_isSpace uintptr
	addr_QChar_isMark uintptr
	addr_QChar_isPunct uintptr
	addr_QChar_isSymbol uintptr
	addr_QChar_isLetter uintptr
	addr_QChar_isNumber uintptr
	addr_QChar_isLetterOrNumber uintptr
	addr_QChar_isDigit uintptr
	addr_QChar_isLower uintptr
	addr_QChar_isUpper uintptr
	addr_QChar_isTitleCase uintptr
	addr_QChar_isNonCharacter uintptr
	addr_QChar_isHighSurrogate uintptr
	addr_QChar_isLowSurrogate uintptr
	addr_QChar_isSurrogate uintptr
	addr_QChar_cell uintptr
	addr_QChar_row uintptr
	addr_QChar_setCell uintptr
	addr_QChar_setRow uintptr
	addr_QChar_isNonCharacter2 uintptr
	addr_QChar_isHighSurrogate2 uintptr
	addr_QChar_isLowSurrogate2 uintptr
	addr_QChar_isSurrogate2 uintptr
	addr_QChar_requiresSurrogates uintptr
	addr_QChar_surrogateToUcs4 uintptr
	addr_QChar_surrogateToUcs42 uintptr
	addr_QChar_highSurrogate uintptr
	addr_QChar_lowSurrogate uintptr
	addr_QChar_category2 uintptr
	addr_QChar_direction2 uintptr
	addr_QChar_joiningType2 uintptr
	addr_QChar_combiningClass2 uintptr
	addr_QChar_mirroredChar2 uintptr
	addr_QChar_hasMirrored2 uintptr
	addr_QChar_decomposition2 uintptr
	addr_QChar_decompositionTag2 uintptr
	addr_QChar_digitValue2 uintptr
	addr_QChar_toLower2 uintptr
	addr_QChar_toUpper2 uintptr
	addr_QChar_toTitleCase2 uintptr
	addr_QChar_toCaseFolded2 uintptr
	addr_QChar_script2 uintptr
	addr_QChar_unicodeVersion2 uintptr
	addr_QChar_currentUnicodeVersion uintptr
	addr_QChar_isPrint2 uintptr
	addr_QChar_isSpace2 uintptr
	addr_QChar_isMark2 uintptr
	addr_QChar_isPunct2 uintptr
	addr_QChar_isSymbol2 uintptr
	addr_QChar_isLetter2 uintptr
	addr_QChar_isNumber2 uintptr
	addr_QChar_isLetterOrNumber2 uintptr
	addr_QChar_isDigit2 uintptr
	addr_QChar_isLower2 uintptr
	addr_QChar_isUpper2 uintptr
	addr_QChar_isTitleCase2 uintptr
	addr_QChar_OperatorEqual uintptr
	addr_QChar_OperatorNotEqual uintptr
	addr_QChar_OperatorLesser uintptr
	addr_QChar_OperatorGreater uintptr
	addr_QChar_OperatorLesserOrEqual uintptr
	addr_QChar_OperatorGreaterOrEqual uintptr
	addr_QChar_OperatorEqual2 uintptr
	addr_QChar_OperatorNotEqual2 uintptr
	addr_QChar_OperatorLesser2 uintptr
	addr_QChar_OperatorGreater2 uintptr
	addr_QChar_OperatorLesserOrEqual2 uintptr
	addr_QChar_OperatorGreaterOrEqual2 uintptr
	addr_QChar_OperatorEqual3 uintptr
	addr_QChar_OperatorNotEqual3 uintptr
	addr_QChar_OperatorLesser3 uintptr
	addr_QChar_OperatorGreater3 uintptr
	addr_QChar_OperatorLesserOrEqual3 uintptr
	addr_QChar_OperatorGreaterOrEqual3 uintptr
	addr_QChar_OperatorEqual4 uintptr
	addr_QChar_OperatorNotEqual4 uintptr
	addr_QChar_OperatorLesser4 uintptr
	addr_QChar_OperatorGreater4 uintptr
	addr_QChar_OperatorLesserOrEqual4 uintptr
	addr_QChar_OperatorGreaterOrEqual4 uintptr
	addr_QChar_OperatorEqual5 uintptr
	addr_QChar_OperatorNotEqual5 uintptr
	addr_QChar_OperatorLesser5 uintptr
	addr_QChar_OperatorGreater5 uintptr
	addr_QChar_OperatorLesserOrEqual5 uintptr
	addr_QChar_OperatorGreaterOrEqual5 uintptr
	addr_QChar_OperatorPlus uintptr
	pg_QChar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QChar_new, "QChar_new")
	symaddr(&addr_QChar_new2, "QChar_new2")
	symaddr(&addr_QChar_new3, "QChar_new3")
	symaddr(&addr_QChar_fromUcs2, "QChar_fromUcs2")
	symaddr(&addr_QChar_category, "QChar_category")
	symaddr(&addr_QChar_direction, "QChar_direction")
	symaddr(&addr_QChar_joiningType, "QChar_joiningType")
	symaddr(&addr_QChar_combiningClass, "QChar_combiningClass")
	symaddr(&addr_QChar_mirroredChar, "QChar_mirroredChar")
	symaddr(&addr_QChar_hasMirrored, "QChar_hasMirrored")
	symaddr(&addr_QChar_decomposition, "QChar_decomposition")
	symaddr(&addr_QChar_decompositionTag, "QChar_decompositionTag")
	symaddr(&addr_QChar_digitValue, "QChar_digitValue")
	symaddr(&addr_QChar_toLower, "QChar_toLower")
	symaddr(&addr_QChar_toUpper, "QChar_toUpper")
	symaddr(&addr_QChar_toTitleCase, "QChar_toTitleCase")
	symaddr(&addr_QChar_toCaseFolded, "QChar_toCaseFolded")
	symaddr(&addr_QChar_script, "QChar_script")
	symaddr(&addr_QChar_unicodeVersion, "QChar_unicodeVersion")
	symaddr(&addr_QChar_toLatin1, "QChar_toLatin1")
	symaddr(&addr_QChar_unicode, "QChar_unicode")
	symaddr(&addr_QChar_fromLatin1, "QChar_fromLatin1")
	symaddr(&addr_QChar_isNull, "QChar_isNull")
	symaddr(&addr_QChar_isPrint, "QChar_isPrint")
	symaddr(&addr_QChar_isSpace, "QChar_isSpace")
	symaddr(&addr_QChar_isMark, "QChar_isMark")
	symaddr(&addr_QChar_isPunct, "QChar_isPunct")
	symaddr(&addr_QChar_isSymbol, "QChar_isSymbol")
	symaddr(&addr_QChar_isLetter, "QChar_isLetter")
	symaddr(&addr_QChar_isNumber, "QChar_isNumber")
	symaddr(&addr_QChar_isLetterOrNumber, "QChar_isLetterOrNumber")
	symaddr(&addr_QChar_isDigit, "QChar_isDigit")
	symaddr(&addr_QChar_isLower, "QChar_isLower")
	symaddr(&addr_QChar_isUpper, "QChar_isUpper")
	symaddr(&addr_QChar_isTitleCase, "QChar_isTitleCase")
	symaddr(&addr_QChar_isNonCharacter, "QChar_isNonCharacter")
	symaddr(&addr_QChar_isHighSurrogate, "QChar_isHighSurrogate")
	symaddr(&addr_QChar_isLowSurrogate, "QChar_isLowSurrogate")
	symaddr(&addr_QChar_isSurrogate, "QChar_isSurrogate")
	symaddr(&addr_QChar_cell, "QChar_cell")
	symaddr(&addr_QChar_row, "QChar_row")
	symaddr(&addr_QChar_setCell, "QChar_setCell")
	symaddr(&addr_QChar_setRow, "QChar_setRow")
	symaddr(&addr_QChar_isNonCharacter2, "QChar_isNonCharacter2")
	symaddr(&addr_QChar_isHighSurrogate2, "QChar_isHighSurrogate2")
	symaddr(&addr_QChar_isLowSurrogate2, "QChar_isLowSurrogate2")
	symaddr(&addr_QChar_isSurrogate2, "QChar_isSurrogate2")
	symaddr(&addr_QChar_requiresSurrogates, "QChar_requiresSurrogates")
	symaddr(&addr_QChar_surrogateToUcs4, "QChar_surrogateToUcs4")
	symaddr(&addr_QChar_surrogateToUcs42, "QChar_surrogateToUcs42")
	symaddr(&addr_QChar_highSurrogate, "QChar_highSurrogate")
	symaddr(&addr_QChar_lowSurrogate, "QChar_lowSurrogate")
	symaddr(&addr_QChar_category2, "QChar_category2")
	symaddr(&addr_QChar_direction2, "QChar_direction2")
	symaddr(&addr_QChar_joiningType2, "QChar_joiningType2")
	symaddr(&addr_QChar_combiningClass2, "QChar_combiningClass2")
	symaddr(&addr_QChar_mirroredChar2, "QChar_mirroredChar2")
	symaddr(&addr_QChar_hasMirrored2, "QChar_hasMirrored2")
	symaddr(&addr_QChar_decomposition2, "QChar_decomposition2")
	symaddr(&addr_QChar_decompositionTag2, "QChar_decompositionTag2")
	symaddr(&addr_QChar_digitValue2, "QChar_digitValue2")
	symaddr(&addr_QChar_toLower2, "QChar_toLower2")
	symaddr(&addr_QChar_toUpper2, "QChar_toUpper2")
	symaddr(&addr_QChar_toTitleCase2, "QChar_toTitleCase2")
	symaddr(&addr_QChar_toCaseFolded2, "QChar_toCaseFolded2")
	symaddr(&addr_QChar_script2, "QChar_script2")
	symaddr(&addr_QChar_unicodeVersion2, "QChar_unicodeVersion2")
	symaddr(&addr_QChar_currentUnicodeVersion, "QChar_currentUnicodeVersion")
	symaddr(&addr_QChar_isPrint2, "QChar_isPrint2")
	symaddr(&addr_QChar_isSpace2, "QChar_isSpace2")
	symaddr(&addr_QChar_isMark2, "QChar_isMark2")
	symaddr(&addr_QChar_isPunct2, "QChar_isPunct2")
	symaddr(&addr_QChar_isSymbol2, "QChar_isSymbol2")
	symaddr(&addr_QChar_isLetter2, "QChar_isLetter2")
	symaddr(&addr_QChar_isNumber2, "QChar_isNumber2")
	symaddr(&addr_QChar_isLetterOrNumber2, "QChar_isLetterOrNumber2")
	symaddr(&addr_QChar_isDigit2, "QChar_isDigit2")
	symaddr(&addr_QChar_isLower2, "QChar_isLower2")
	symaddr(&addr_QChar_isUpper2, "QChar_isUpper2")
	symaddr(&addr_QChar_isTitleCase2, "QChar_isTitleCase2")
	symaddr(&addr_QChar_OperatorEqual, "QChar_OperatorEqual")
	symaddr(&addr_QChar_OperatorNotEqual, "QChar_OperatorNotEqual")
	symaddr(&addr_QChar_OperatorLesser, "QChar_OperatorLesser")
	symaddr(&addr_QChar_OperatorGreater, "QChar_OperatorGreater")
	symaddr(&addr_QChar_OperatorLesserOrEqual, "QChar_OperatorLesserOrEqual")
	symaddr(&addr_QChar_OperatorGreaterOrEqual, "QChar_OperatorGreaterOrEqual")
	symaddr(&addr_QChar_OperatorEqual2, "QChar_OperatorEqual2")
	symaddr(&addr_QChar_OperatorNotEqual2, "QChar_OperatorNotEqual2")
	symaddr(&addr_QChar_OperatorLesser2, "QChar_OperatorLesser2")
	symaddr(&addr_QChar_OperatorGreater2, "QChar_OperatorGreater2")
	symaddr(&addr_QChar_OperatorLesserOrEqual2, "QChar_OperatorLesserOrEqual2")
	symaddr(&addr_QChar_OperatorGreaterOrEqual2, "QChar_OperatorGreaterOrEqual2")
	symaddr(&addr_QChar_OperatorEqual3, "QChar_OperatorEqual3")
	symaddr(&addr_QChar_OperatorNotEqual3, "QChar_OperatorNotEqual3")
	symaddr(&addr_QChar_OperatorLesser3, "QChar_OperatorLesser3")
	symaddr(&addr_QChar_OperatorGreater3, "QChar_OperatorGreater3")
	symaddr(&addr_QChar_OperatorLesserOrEqual3, "QChar_OperatorLesserOrEqual3")
	symaddr(&addr_QChar_OperatorGreaterOrEqual3, "QChar_OperatorGreaterOrEqual3")
	symaddr(&addr_QChar_OperatorEqual4, "QChar_OperatorEqual4")
	symaddr(&addr_QChar_OperatorNotEqual4, "QChar_OperatorNotEqual4")
	symaddr(&addr_QChar_OperatorLesser4, "QChar_OperatorLesser4")
	symaddr(&addr_QChar_OperatorGreater4, "QChar_OperatorGreater4")
	symaddr(&addr_QChar_OperatorLesserOrEqual4, "QChar_OperatorLesserOrEqual4")
	symaddr(&addr_QChar_OperatorGreaterOrEqual4, "QChar_OperatorGreaterOrEqual4")
	symaddr(&addr_QChar_OperatorEqual5, "QChar_OperatorEqual5")
	symaddr(&addr_QChar_OperatorNotEqual5, "QChar_OperatorNotEqual5")
	symaddr(&addr_QChar_OperatorLesser5, "QChar_OperatorLesser5")
	symaddr(&addr_QChar_OperatorGreater5, "QChar_OperatorGreater5")
	symaddr(&addr_QChar_OperatorLesserOrEqual5, "QChar_OperatorLesserOrEqual5")
	symaddr(&addr_QChar_OperatorGreaterOrEqual5, "QChar_OperatorGreaterOrEqual5")
	symaddr(&addr_QChar_OperatorPlus, "QChar_OperatorPlus")
	reg(&pg_QChar_delete, "QChar_delete")
}

type QChar struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QChar) UnsafeGoID() int64     { return this.vid }
func (this *QChar) UnsafeSetGoID(id int64) { this.vid = id }

func bindQChar(h uintptr) *QChar {
	if h == 0 {
		return nil
	}
	return &QChar{h: h}
}

func newownedQChar(h uintptr) *QChar {
	o := bindQChar(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QChar).Delete)
	}
	return o
}

func (this *QChar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QChar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQChar(h unsafe.Pointer) *QChar { return bindQChar(uintptr(h)) }
func UnsafeNewQCharOwned(h unsafe.Pointer) *QChar { return newownedQChar(uintptr(h)) }

func (this *QChar) UnsafeRelease() { this.h = 0 }

func NewQChar() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_new)
	__o := newownedQChar(__r)
	return __o
}

func NewQChar2(c byte, r byte) *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_new2, uintptr(c), uintptr(r))
	__o := newownedQChar(__r)
	return __o
}

func NewQChar3(arg1 *QChar) *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQChar(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QChar_FromUcs2(c uint16) *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_fromUcs2, uintptr(c))
	return newownedQChar(__r)
}

func (this *QChar) Category() QChar__Category {
	__r0, _, _ := purego.SyscallN(addr_QChar_category, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__Category(__r)
}

func (this *QChar) Direction() QChar__Direction {
	__r0, _, _ := purego.SyscallN(addr_QChar_direction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__Direction(__r)
}

func (this *QChar) JoiningType() QChar__JoiningType {
	__r0, _, _ := purego.SyscallN(addr_QChar_joiningType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__JoiningType(__r)
}

func (this *QChar) CombiningClass() byte {
	__r0, _, _ := purego.SyscallN(addr_QChar_combiningClass, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QChar) MirroredChar() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_mirroredChar, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QChar) HasMirrored() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_hasMirrored, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) Decomposition() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChar_decomposition, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QChar) DecompositionTag() QChar__Decomposition {
	__r0, _, _ := purego.SyscallN(addr_QChar_decompositionTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__Decomposition(__r)
}

func (this *QChar) DigitValue() int {
	__r0, _, _ := purego.SyscallN(addr_QChar_digitValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QChar) ToLower() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_toLower, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QChar) ToUpper() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_toUpper, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QChar) ToTitleCase() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_toTitleCase, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QChar) ToCaseFolded() *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_toCaseFolded, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QChar) Script() QChar__Script {
	__r0, _, _ := purego.SyscallN(addr_QChar_script, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__Script(__r)
}

func (this *QChar) UnicodeVersion() QChar__UnicodeVersion {
	__r0, _, _ := purego.SyscallN(addr_QChar_unicodeVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QChar__UnicodeVersion(__r)
}

func (this *QChar) ToLatin1() int8 {
	__r0, _, _ := purego.SyscallN(addr_QChar_toLatin1, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QChar) Unicode() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QChar_unicode, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func QChar_FromLatin1(c int8) *QChar {
	__r, _, _ := purego.SyscallN(addr_QChar_fromLatin1, uintptr(c))
	return newownedQChar(__r)
}

func (this *QChar) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsPrint() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isPrint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsSpace() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSpace, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsMark() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isMark, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsPunct() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isPunct, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsSymbol() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSymbol, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsLetter() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLetter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsNumber() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsLetterOrNumber() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLetterOrNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsDigit() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isDigit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsLower() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLower, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsUpper() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isUpper, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsTitleCase() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isTitleCase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsNonCharacter() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isNonCharacter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsHighSurrogate() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isHighSurrogate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsLowSurrogate() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLowSurrogate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) IsSurrogate() bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSurrogate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChar) Cell() byte {
	__r0, _, _ := purego.SyscallN(addr_QChar_cell, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QChar) Row() byte {
	__r0, _, _ := purego.SyscallN(addr_QChar_row, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QChar) SetCell(acell byte) {
	purego.SyscallN(addr_QChar_setCell, this.h, uintptr(acell))
	runtime.KeepAlive(this)
}

func (this *QChar) SetRow(arow byte) {
	purego.SyscallN(addr_QChar_setRow, this.h, uintptr(arow))
	runtime.KeepAlive(this)
}

func QChar_IsNonCharacter(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isNonCharacter2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsHighSurrogate(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isHighSurrogate2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsLowSurrogate(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLowSurrogate2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsSurrogate(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSurrogate2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_RequiresSurrogates(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_requiresSurrogates, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_SurrogateToUcs4(high uint16, low uint16) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_surrogateToUcs4, uintptr(high), uintptr(low))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_SurrogateToUcs42(high *QChar, low *QChar) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_surrogateToUcs42, uintptr(high.UnsafePointer()), uintptr(low.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(high)
	runtime.KeepAlive(low)
	return uint32(__r)
}

func QChar_HighSurrogate(ucs4 uint32) uint16 {
	__r0, _, _ := purego.SyscallN(addr_QChar_highSurrogate, uintptr(ucs4))
	__r := int16(uint16(__r0))
	return uint16(__r)
}

func QChar_LowSurrogate(ucs4 uint32) uint16 {
	__r0, _, _ := purego.SyscallN(addr_QChar_lowSurrogate, uintptr(ucs4))
	__r := int16(uint16(__r0))
	return uint16(__r)
}

func QChar_Category(ucs4 uint32) QChar__Category {
	__r0, _, _ := purego.SyscallN(addr_QChar_category2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__Category(__r)
}

func QChar_Direction(ucs4 uint32) QChar__Direction {
	__r0, _, _ := purego.SyscallN(addr_QChar_direction2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__Direction(__r)
}

func QChar_JoiningType(ucs4 uint32) QChar__JoiningType {
	__r0, _, _ := purego.SyscallN(addr_QChar_joiningType2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__JoiningType(__r)
}

func QChar_CombiningClass(ucs4 uint32) byte {
	__r0, _, _ := purego.SyscallN(addr_QChar_combiningClass2, uintptr(ucs4))
	__r := uint8(__r0)
	return byte(__r)
}

func QChar_MirroredChar(ucs4 uint32) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_mirroredChar2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_HasMirrored(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_hasMirrored2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_Decomposition(ucs4 uint32) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChar_decomposition2, uintptr(ucs4), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QChar_DecompositionTag(ucs4 uint32) QChar__Decomposition {
	__r0, _, _ := purego.SyscallN(addr_QChar_decompositionTag2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__Decomposition(__r)
}

func QChar_DigitValue(ucs4 uint32) int {
	__r0, _, _ := purego.SyscallN(addr_QChar_digitValue2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QChar_ToLower(ucs4 uint32) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_toLower2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_ToUpper(ucs4 uint32) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_toUpper2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_ToTitleCase(ucs4 uint32) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_toTitleCase2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_ToCaseFolded(ucs4 uint32) uint32 {
	__r0, _, _ := purego.SyscallN(addr_QChar_toCaseFolded2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return uint32(__r)
}

func QChar_Script(ucs4 uint32) QChar__Script {
	__r0, _, _ := purego.SyscallN(addr_QChar_script2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__Script(__r)
}

func QChar_UnicodeVersion(ucs4 uint32) QChar__UnicodeVersion {
	__r0, _, _ := purego.SyscallN(addr_QChar_unicodeVersion2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return QChar__UnicodeVersion(__r)
}

func QChar_CurrentUnicodeVersion() QChar__UnicodeVersion {
	__r0, _, _ := purego.SyscallN(addr_QChar_currentUnicodeVersion)
	__r := int32(uint32(__r0))
	return QChar__UnicodeVersion(__r)
}

func QChar_IsPrint(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isPrint2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsSpace(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSpace2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsMark(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isMark2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsPunct(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isPunct2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsSymbol(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isSymbol2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsLetter(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLetter2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsNumber(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isNumber2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsLetterOrNumber(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLetterOrNumber2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsDigit(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isDigit2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsLower(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isLower2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsUpper(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isUpper2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QChar_IsTitleCase(ucs4 uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_isTitleCase2, uintptr(ucs4))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QChar) OperatorEqual(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorNotEqual(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesser(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreater(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserOrEqual(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterOrEqual(rhs *QChar) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorEqualWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorNotEqualWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesser2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreater2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserOrEqualWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesserOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterOrEqualWithRhs(rhs *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreaterOrEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorEqual2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorNotEqual2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorNotEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesser2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesser3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreater2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreater3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserOrEqual2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesserOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterOrEqual2(rhs *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreaterOrEqual3, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorEqual3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorNotEqual3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorNotEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesser3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesser4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreater3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreater4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserOrEqual3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesserOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterOrEqual3(rhs *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreaterOrEqual4, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorEqual4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorEqual5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorNotEqual4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorNotEqual5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesser4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesser5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreater4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreater5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorLesserOrEqual4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorLesserOrEqual5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorGreaterOrEqual4(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QChar_OperatorGreaterOrEqual5, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QChar) OperatorPlus(s2 string) string {
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QChar_OperatorPlus, this.h, s2_p, uintptr(len(s2)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s2)
	return pgStr(__r, __n)
}

func (this *QChar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QChar_delete(this.h)
	}
	this.UnsafeRelease()
}
