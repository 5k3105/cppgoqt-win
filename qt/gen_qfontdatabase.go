package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontDatabase_new uintptr
	addr_QFontDatabase_standardSizes uintptr
	addr_QFontDatabase_writingSystems uintptr
	addr_QFontDatabase_writingSystems2 uintptr
	addr_QFontDatabase_families uintptr
	addr_QFontDatabase_styles uintptr
	addr_QFontDatabase_pointSizes uintptr
	addr_QFontDatabase_smoothSizes uintptr
	addr_QFontDatabase_styleString uintptr
	addr_QFontDatabase_styleString2 uintptr
	addr_QFontDatabase_font uintptr
	addr_QFontDatabase_isBitmapScalable uintptr
	addr_QFontDatabase_isSmoothlyScalable uintptr
	addr_QFontDatabase_isScalable uintptr
	addr_QFontDatabase_isFixedPitch uintptr
	addr_QFontDatabase_italic uintptr
	addr_QFontDatabase_bold uintptr
	addr_QFontDatabase_weight uintptr
	addr_QFontDatabase_hasFamily uintptr
	addr_QFontDatabase_isPrivateFamily uintptr
	addr_QFontDatabase_writingSystemName uintptr
	addr_QFontDatabase_writingSystemSample uintptr
	addr_QFontDatabase_addApplicationFont uintptr
	addr_QFontDatabase_addApplicationFontFromData uintptr
	addr_QFontDatabase_applicationFontFamilies uintptr
	addr_QFontDatabase_removeApplicationFont uintptr
	addr_QFontDatabase_removeAllApplicationFonts uintptr
	addr_QFontDatabase_addApplicationFallbackFontFamily uintptr
	addr_QFontDatabase_removeApplicationFallbackFontFamily uintptr
	addr_QFontDatabase_setApplicationFallbackFontFamilies uintptr
	addr_QFontDatabase_applicationFallbackFontFamilies uintptr
	addr_QFontDatabase_addApplicationEmojiFontFamily uintptr
	addr_QFontDatabase_removeApplicationEmojiFontFamily uintptr
	addr_QFontDatabase_setApplicationEmojiFontFamilies uintptr
	addr_QFontDatabase_applicationEmojiFontFamilies uintptr
	addr_QFontDatabase_systemFont uintptr
	addr_QFontDatabase_families2 uintptr
	addr_QFontDatabase_pointSizes2 uintptr
	addr_QFontDatabase_isBitmapScalable2 uintptr
	addr_QFontDatabase_isSmoothlyScalable2 uintptr
	addr_QFontDatabase_isScalable2 uintptr
	addr_QFontDatabase_isFixedPitch2 uintptr
	pg_QFontDatabase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontDatabase_new, "QFontDatabase_new")
	symaddr(&addr_QFontDatabase_standardSizes, "QFontDatabase_standardSizes")
	symaddr(&addr_QFontDatabase_writingSystems, "QFontDatabase_writingSystems")
	symaddr(&addr_QFontDatabase_writingSystems2, "QFontDatabase_writingSystems2")
	symaddr(&addr_QFontDatabase_families, "QFontDatabase_families")
	symaddr(&addr_QFontDatabase_styles, "QFontDatabase_styles")
	symaddr(&addr_QFontDatabase_pointSizes, "QFontDatabase_pointSizes")
	symaddr(&addr_QFontDatabase_smoothSizes, "QFontDatabase_smoothSizes")
	symaddr(&addr_QFontDatabase_styleString, "QFontDatabase_styleString")
	symaddr(&addr_QFontDatabase_styleString2, "QFontDatabase_styleString2")
	symaddr(&addr_QFontDatabase_font, "QFontDatabase_font")
	symaddr(&addr_QFontDatabase_isBitmapScalable, "QFontDatabase_isBitmapScalable")
	symaddr(&addr_QFontDatabase_isSmoothlyScalable, "QFontDatabase_isSmoothlyScalable")
	symaddr(&addr_QFontDatabase_isScalable, "QFontDatabase_isScalable")
	symaddr(&addr_QFontDatabase_isFixedPitch, "QFontDatabase_isFixedPitch")
	symaddr(&addr_QFontDatabase_italic, "QFontDatabase_italic")
	symaddr(&addr_QFontDatabase_bold, "QFontDatabase_bold")
	symaddr(&addr_QFontDatabase_weight, "QFontDatabase_weight")
	symaddr(&addr_QFontDatabase_hasFamily, "QFontDatabase_hasFamily")
	symaddr(&addr_QFontDatabase_isPrivateFamily, "QFontDatabase_isPrivateFamily")
	symaddr(&addr_QFontDatabase_writingSystemName, "QFontDatabase_writingSystemName")
	symaddr(&addr_QFontDatabase_writingSystemSample, "QFontDatabase_writingSystemSample")
	symaddr(&addr_QFontDatabase_addApplicationFont, "QFontDatabase_addApplicationFont")
	symaddr(&addr_QFontDatabase_addApplicationFontFromData, "QFontDatabase_addApplicationFontFromData")
	symaddr(&addr_QFontDatabase_applicationFontFamilies, "QFontDatabase_applicationFontFamilies")
	symaddr(&addr_QFontDatabase_removeApplicationFont, "QFontDatabase_removeApplicationFont")
	symaddr(&addr_QFontDatabase_removeAllApplicationFonts, "QFontDatabase_removeAllApplicationFonts")
	symaddr(&addr_QFontDatabase_addApplicationFallbackFontFamily, "QFontDatabase_addApplicationFallbackFontFamily")
	symaddr(&addr_QFontDatabase_removeApplicationFallbackFontFamily, "QFontDatabase_removeApplicationFallbackFontFamily")
	symaddr(&addr_QFontDatabase_setApplicationFallbackFontFamilies, "QFontDatabase_setApplicationFallbackFontFamilies")
	symaddr(&addr_QFontDatabase_applicationFallbackFontFamilies, "QFontDatabase_applicationFallbackFontFamilies")
	symaddr(&addr_QFontDatabase_addApplicationEmojiFontFamily, "QFontDatabase_addApplicationEmojiFontFamily")
	symaddr(&addr_QFontDatabase_removeApplicationEmojiFontFamily, "QFontDatabase_removeApplicationEmojiFontFamily")
	symaddr(&addr_QFontDatabase_setApplicationEmojiFontFamilies, "QFontDatabase_setApplicationEmojiFontFamilies")
	symaddr(&addr_QFontDatabase_applicationEmojiFontFamilies, "QFontDatabase_applicationEmojiFontFamilies")
	symaddr(&addr_QFontDatabase_systemFont, "QFontDatabase_systemFont")
	symaddr(&addr_QFontDatabase_families2, "QFontDatabase_families2")
	symaddr(&addr_QFontDatabase_pointSizes2, "QFontDatabase_pointSizes2")
	symaddr(&addr_QFontDatabase_isBitmapScalable2, "QFontDatabase_isBitmapScalable2")
	symaddr(&addr_QFontDatabase_isSmoothlyScalable2, "QFontDatabase_isSmoothlyScalable2")
	symaddr(&addr_QFontDatabase_isScalable2, "QFontDatabase_isScalable2")
	symaddr(&addr_QFontDatabase_isFixedPitch2, "QFontDatabase_isFixedPitch2")
	reg(&pg_QFontDatabase_delete, "QFontDatabase_delete")
}

type QFontDatabase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFontDatabase) UnsafeGoID() int64     { return this.vid }
func (this *QFontDatabase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFontDatabase(h uintptr) *QFontDatabase {
	if h == 0 {
		return nil
	}
	return &QFontDatabase{h: h}
}

func newownedQFontDatabase(h uintptr) *QFontDatabase {
	o := bindQFontDatabase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFontDatabase).Delete)
	}
	return o
}

func (this *QFontDatabase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontDatabase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontDatabase(h unsafe.Pointer) *QFontDatabase { return bindQFontDatabase(uintptr(h)) }
func UnsafeNewQFontDatabaseOwned(h unsafe.Pointer) *QFontDatabase { return newownedQFontDatabase(uintptr(h)) }

func (this *QFontDatabase) UnsafeRelease() { this.h = 0 }

func NewQFontDatabase() *QFontDatabase {
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_new)
	__o := newownedQFontDatabase(__r)
	return __o
}

func QFontDatabase_StandardSizes() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_standardSizes, uintptr(unsafe.Pointer(&__n)))
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_WritingSystems() []QFontDatabase__WritingSystem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_writingSystems, uintptr(unsafe.Pointer(&__n)))
	__out := make([]QFontDatabase__WritingSystem, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QFontDatabase__WritingSystem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_WritingSystemsWithFamily(family string) []QFontDatabase__WritingSystem {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_writingSystems2, family_p, uintptr(len(family)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(family)
	__out := make([]QFontDatabase__WritingSystem, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QFontDatabase__WritingSystem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_Families() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_families, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_Styles(family string) []string {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_styles, family_p, uintptr(len(family)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(family)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_PointSizes(family string) []int {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_pointSizes, family_p, uintptr(len(family)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(family)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_SmoothSizes(family string, style string) []int {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_smoothSizes, family_p, uintptr(len(family)), style_p, uintptr(len(style)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_StyleString(font *QFont) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_styleString, uintptr(font.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(font)
	return pgStr(__r, __n)
}

func QFontDatabase_StyleStringWithFontInfo(fontInfo *QFontInfo) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_styleString2, uintptr(fontInfo.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fontInfo)
	return pgStr(__r, __n)
}

func QFontDatabase_Font(family string, style string, pointSize int) *QFont {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_font, family_p, uintptr(len(family)), style_p, uintptr(len(style)), uintptr(pointSize))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return newownedQFont(__r)
}

func QFontDatabase_IsBitmapScalable(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isBitmapScalable, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_IsSmoothlyScalable(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isSmoothlyScalable, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_IsScalable(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isScalable, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_IsFixedPitch(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isFixedPitch, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_Italic(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_italic, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func QFontDatabase_Bold(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_bold, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func QFontDatabase_Weight(family string, style string) int {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_weight, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return int(__r)
}

func QFontDatabase_HasFamily(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_hasFamily, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_IsPrivateFamily(family string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isPrivateFamily, family_p, uintptr(len(family)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	return __r != 0
}

func QFontDatabase_WritingSystemName(writingSystem QFontDatabase__WritingSystem) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_writingSystemName, uintptr(writingSystem), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QFontDatabase_WritingSystemSample(writingSystem QFontDatabase__WritingSystem) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_writingSystemSample, uintptr(writingSystem), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QFontDatabase_AddApplicationFont(fileName string) int {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_addApplicationFont, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return int(__r)
}

func QFontDatabase_AddApplicationFontFromData(fontData *QByteArray) int {
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_addApplicationFontFromData, uintptr(fontData.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fontData)
	return int(__r)
}

func QFontDatabase_ApplicationFontFamilies(id int) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_applicationFontFamilies, uintptr(id), uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_RemoveApplicationFont(id int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_removeApplicationFont, uintptr(id))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QFontDatabase_RemoveAllApplicationFonts() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_removeAllApplicationFonts)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QFontDatabase_AddApplicationFallbackFontFamily(script QChar__Script, familyName string) {
	var familyName_p uintptr
	if len(familyName) > 0 {
		familyName_p = uintptr(unsafe.Pointer(unsafe.StringData(familyName)))
	}
	purego.SyscallN(addr_QFontDatabase_addApplicationFallbackFontFamily, uintptr(script), familyName_p, uintptr(len(familyName)))
	runtime.KeepAlive(familyName)
}

func QFontDatabase_RemoveApplicationFallbackFontFamily(script QChar__Script, familyName string) bool {
	var familyName_p uintptr
	if len(familyName) > 0 {
		familyName_p = uintptr(unsafe.Pointer(unsafe.StringData(familyName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_removeApplicationFallbackFontFamily, uintptr(script), familyName_p, uintptr(len(familyName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(familyName)
	return __r != 0
}

func QFontDatabase_SetApplicationFallbackFontFamilies(arg1 QChar__Script, familyNames []string) {
	familyNames_arr := make([]uintptr, len(familyNames))
	for __i, __e := range familyNames {
		familyNames_arr[__i] = pgCString(__e)
	}
	var familyNames_p uintptr
	if len(familyNames_arr) > 0 {
		familyNames_p = uintptr(unsafe.Pointer(&familyNames_arr[0]))
	}
	purego.SyscallN(addr_QFontDatabase_setApplicationFallbackFontFamilies, uintptr(arg1), familyNames_p, uintptr(len(familyNames)))
	for __i := range familyNames_arr {
		pgFree(familyNames_arr[__i])
	}
	runtime.KeepAlive(familyNames_arr)
}

func QFontDatabase_ApplicationFallbackFontFamilies(script QChar__Script) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_applicationFallbackFontFamilies, uintptr(script), uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_AddApplicationEmojiFontFamily(familyName string) {
	var familyName_p uintptr
	if len(familyName) > 0 {
		familyName_p = uintptr(unsafe.Pointer(unsafe.StringData(familyName)))
	}
	purego.SyscallN(addr_QFontDatabase_addApplicationEmojiFontFamily, familyName_p, uintptr(len(familyName)))
	runtime.KeepAlive(familyName)
}

func QFontDatabase_RemoveApplicationEmojiFontFamily(familyName string) bool {
	var familyName_p uintptr
	if len(familyName) > 0 {
		familyName_p = uintptr(unsafe.Pointer(unsafe.StringData(familyName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_removeApplicationEmojiFontFamily, familyName_p, uintptr(len(familyName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(familyName)
	return __r != 0
}

func QFontDatabase_SetApplicationEmojiFontFamilies(familyNames []string) {
	familyNames_arr := make([]uintptr, len(familyNames))
	for __i, __e := range familyNames {
		familyNames_arr[__i] = pgCString(__e)
	}
	var familyNames_p uintptr
	if len(familyNames_arr) > 0 {
		familyNames_p = uintptr(unsafe.Pointer(&familyNames_arr[0]))
	}
	purego.SyscallN(addr_QFontDatabase_setApplicationEmojiFontFamilies, familyNames_p, uintptr(len(familyNames)))
	for __i := range familyNames_arr {
		pgFree(familyNames_arr[__i])
	}
	runtime.KeepAlive(familyNames_arr)
}

func QFontDatabase_ApplicationEmojiFontFamilies() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_applicationEmojiFontFamilies, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_SystemFont(type_ QFontDatabase__SystemFont) *QFont {
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_systemFont, uintptr(type_))
	return newownedQFont(__r)
}

func QFontDatabase_FamiliesWithWritingSystem(writingSystem QFontDatabase__WritingSystem) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_families2, uintptr(writingSystem), uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_PointSizes2(family string, style string) []int {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontDatabase_pointSizes2, family_p, uintptr(len(family)), style_p, uintptr(len(style)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFontDatabase_IsBitmapScalable2(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isBitmapScalable2, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func QFontDatabase_IsSmoothlyScalable2(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isSmoothlyScalable2, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func QFontDatabase_IsScalable2(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isScalable2, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func QFontDatabase_IsFixedPitch2(family string, style string) bool {
	var family_p uintptr
	if len(family) > 0 {
		family_p = uintptr(unsafe.Pointer(unsafe.StringData(family)))
	}
	var style_p uintptr
	if len(style) > 0 {
		style_p = uintptr(unsafe.Pointer(unsafe.StringData(style)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFontDatabase_isFixedPitch2, family_p, uintptr(len(family)), style_p, uintptr(len(style)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(family)
	runtime.KeepAlive(style)
	return __r != 0
}

func (this *QFontDatabase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontDatabase_delete(this.h)
	}
	this.UnsafeRelease()
}
