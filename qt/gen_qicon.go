package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIcon_new uintptr
	addr_QIcon_new2 uintptr
	addr_QIcon_new3 uintptr
	addr_QIcon_new4 uintptr
	addr_QIcon_new5 uintptr
	addr_QIcon_OperatorAssign uintptr
	addr_QIcon_swap uintptr
	addr_QIcon_ToQVariant uintptr
	addr_QIcon_pixmap uintptr
	addr_QIcon_pixmap2 uintptr
	addr_QIcon_pixmap3 uintptr
	pg_QIcon_pixmap4 func(uintptr, uintptr, float64) uintptr
	addr_QIcon_pixmap5 uintptr
	addr_QIcon_actualSize uintptr
	addr_QIcon_actualSize2 uintptr
	addr_QIcon_name uintptr
	addr_QIcon_paint uintptr
	addr_QIcon_paint2 uintptr
	addr_QIcon_isNull uintptr
	addr_QIcon_isDetached uintptr
	addr_QIcon_detach uintptr
	addr_QIcon_cacheKey uintptr
	addr_QIcon_addPixmap uintptr
	addr_QIcon_addFile uintptr
	addr_QIcon_availableSizes uintptr
	addr_QIcon_setIsMask uintptr
	addr_QIcon_isMask uintptr
	addr_QIcon_fromTheme uintptr
	addr_QIcon_fromTheme2 uintptr
	addr_QIcon_hasThemeIcon uintptr
	addr_QIcon_fromTheme3 uintptr
	addr_QIcon_fromTheme4 uintptr
	addr_QIcon_hasThemeIcon2 uintptr
	addr_QIcon_themeSearchPaths uintptr
	addr_QIcon_setThemeSearchPaths uintptr
	addr_QIcon_fallbackSearchPaths uintptr
	addr_QIcon_setFallbackSearchPaths uintptr
	addr_QIcon_themeName uintptr
	addr_QIcon_setThemeName uintptr
	addr_QIcon_fallbackThemeName uintptr
	addr_QIcon_setFallbackThemeName uintptr
	addr_QIcon_pixmap6 uintptr
	addr_QIcon_pixmap7 uintptr
	addr_QIcon_pixmap8 uintptr
	addr_QIcon_pixmap9 uintptr
	addr_QIcon_pixmap10 uintptr
	addr_QIcon_pixmap11 uintptr
	pg_QIcon_pixmap12 func(uintptr, uintptr, float64, int32) uintptr
	pg_QIcon_pixmap13 func(uintptr, uintptr, float64, int32, int32) uintptr
	addr_QIcon_pixmap14 uintptr
	addr_QIcon_pixmap15 uintptr
	addr_QIcon_actualSize3 uintptr
	addr_QIcon_actualSize4 uintptr
	addr_QIcon_actualSize5 uintptr
	addr_QIcon_actualSize6 uintptr
	addr_QIcon_paint3 uintptr
	addr_QIcon_paint4 uintptr
	addr_QIcon_paint5 uintptr
	addr_QIcon_paint6 uintptr
	addr_QIcon_paint7 uintptr
	addr_QIcon_paint8 uintptr
	addr_QIcon_addPixmap2 uintptr
	addr_QIcon_addPixmap3 uintptr
	addr_QIcon_addFile2 uintptr
	addr_QIcon_addFile3 uintptr
	addr_QIcon_addFile4 uintptr
	addr_QIcon_availableSizes2 uintptr
	addr_QIcon_availableSizes3 uintptr
	addr_QIcon_toQVariant uintptr
	pg_QIcon_delete func(uintptr)
)

func init() {
	symaddr(&addr_QIcon_new, "QIcon_new")
	symaddr(&addr_QIcon_new2, "QIcon_new2")
	symaddr(&addr_QIcon_new3, "QIcon_new3")
	symaddr(&addr_QIcon_new4, "QIcon_new4")
	symaddr(&addr_QIcon_new5, "QIcon_new5")
	symaddr(&addr_QIcon_OperatorAssign, "QIcon_OperatorAssign")
	symaddr(&addr_QIcon_swap, "QIcon_swap")
	symaddr(&addr_QIcon_ToQVariant, "QIcon_ToQVariant")
	symaddr(&addr_QIcon_pixmap, "QIcon_pixmap")
	symaddr(&addr_QIcon_pixmap2, "QIcon_pixmap2")
	symaddr(&addr_QIcon_pixmap3, "QIcon_pixmap3")
	reg(&pg_QIcon_pixmap4, "QIcon_pixmap4")
	symaddr(&addr_QIcon_pixmap5, "QIcon_pixmap5")
	symaddr(&addr_QIcon_actualSize, "QIcon_actualSize")
	symaddr(&addr_QIcon_actualSize2, "QIcon_actualSize2")
	symaddr(&addr_QIcon_name, "QIcon_name")
	symaddr(&addr_QIcon_paint, "QIcon_paint")
	symaddr(&addr_QIcon_paint2, "QIcon_paint2")
	symaddr(&addr_QIcon_isNull, "QIcon_isNull")
	symaddr(&addr_QIcon_isDetached, "QIcon_isDetached")
	symaddr(&addr_QIcon_detach, "QIcon_detach")
	symaddr(&addr_QIcon_cacheKey, "QIcon_cacheKey")
	symaddr(&addr_QIcon_addPixmap, "QIcon_addPixmap")
	symaddr(&addr_QIcon_addFile, "QIcon_addFile")
	symaddr(&addr_QIcon_availableSizes, "QIcon_availableSizes")
	symaddr(&addr_QIcon_setIsMask, "QIcon_setIsMask")
	symaddr(&addr_QIcon_isMask, "QIcon_isMask")
	symaddr(&addr_QIcon_fromTheme, "QIcon_fromTheme")
	symaddr(&addr_QIcon_fromTheme2, "QIcon_fromTheme2")
	symaddr(&addr_QIcon_hasThemeIcon, "QIcon_hasThemeIcon")
	symaddr(&addr_QIcon_fromTheme3, "QIcon_fromTheme3")
	symaddr(&addr_QIcon_fromTheme4, "QIcon_fromTheme4")
	symaddr(&addr_QIcon_hasThemeIcon2, "QIcon_hasThemeIcon2")
	symaddr(&addr_QIcon_themeSearchPaths, "QIcon_themeSearchPaths")
	symaddr(&addr_QIcon_setThemeSearchPaths, "QIcon_setThemeSearchPaths")
	symaddr(&addr_QIcon_fallbackSearchPaths, "QIcon_fallbackSearchPaths")
	symaddr(&addr_QIcon_setFallbackSearchPaths, "QIcon_setFallbackSearchPaths")
	symaddr(&addr_QIcon_themeName, "QIcon_themeName")
	symaddr(&addr_QIcon_setThemeName, "QIcon_setThemeName")
	symaddr(&addr_QIcon_fallbackThemeName, "QIcon_fallbackThemeName")
	symaddr(&addr_QIcon_setFallbackThemeName, "QIcon_setFallbackThemeName")
	symaddr(&addr_QIcon_pixmap6, "QIcon_pixmap6")
	symaddr(&addr_QIcon_pixmap7, "QIcon_pixmap7")
	symaddr(&addr_QIcon_pixmap8, "QIcon_pixmap8")
	symaddr(&addr_QIcon_pixmap9, "QIcon_pixmap9")
	symaddr(&addr_QIcon_pixmap10, "QIcon_pixmap10")
	symaddr(&addr_QIcon_pixmap11, "QIcon_pixmap11")
	reg(&pg_QIcon_pixmap12, "QIcon_pixmap12")
	reg(&pg_QIcon_pixmap13, "QIcon_pixmap13")
	symaddr(&addr_QIcon_pixmap14, "QIcon_pixmap14")
	symaddr(&addr_QIcon_pixmap15, "QIcon_pixmap15")
	symaddr(&addr_QIcon_actualSize3, "QIcon_actualSize3")
	symaddr(&addr_QIcon_actualSize4, "QIcon_actualSize4")
	symaddr(&addr_QIcon_actualSize5, "QIcon_actualSize5")
	symaddr(&addr_QIcon_actualSize6, "QIcon_actualSize6")
	symaddr(&addr_QIcon_paint3, "QIcon_paint3")
	symaddr(&addr_QIcon_paint4, "QIcon_paint4")
	symaddr(&addr_QIcon_paint5, "QIcon_paint5")
	symaddr(&addr_QIcon_paint6, "QIcon_paint6")
	symaddr(&addr_QIcon_paint7, "QIcon_paint7")
	symaddr(&addr_QIcon_paint8, "QIcon_paint8")
	symaddr(&addr_QIcon_addPixmap2, "QIcon_addPixmap2")
	symaddr(&addr_QIcon_addPixmap3, "QIcon_addPixmap3")
	symaddr(&addr_QIcon_addFile2, "QIcon_addFile2")
	symaddr(&addr_QIcon_addFile3, "QIcon_addFile3")
	symaddr(&addr_QIcon_addFile4, "QIcon_addFile4")
	symaddr(&addr_QIcon_availableSizes2, "QIcon_availableSizes2")
	symaddr(&addr_QIcon_availableSizes3, "QIcon_availableSizes3")
	symaddr(&addr_QIcon_toQVariant, "QIcon_toQVariant")
	reg(&pg_QIcon_delete, "QIcon_delete")
}

type QIcon struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QIcon) UnsafeGoID() int64     { return this.vid }
func (this *QIcon) UnsafeSetGoID(id int64) { this.vid = id }

func bindQIcon(h uintptr) *QIcon {
	if h == 0 {
		return nil
	}
	return &QIcon{h: h}
}

func newownedQIcon(h uintptr) *QIcon {
	o := bindQIcon(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QIcon).Delete)
	}
	return o
}

func (this *QIcon) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIcon) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIcon(h unsafe.Pointer) *QIcon { return bindQIcon(uintptr(h)) }
func UnsafeNewQIconOwned(h unsafe.Pointer) *QIcon { return newownedQIcon(uintptr(h)) }

func (this *QIcon) UnsafeRelease() { this.h = 0 }

func NewQIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_new)
	__o := newownedQIcon(__r)
	return __o
}

func NewQIcon2(pixmap *QPixmap) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_new2, uintptr(pixmap.UnsafePointer()))
	__o := newownedQIcon(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQIcon3(other *QIcon) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_new3, uintptr(other.UnsafePointer()))
	__o := newownedQIcon(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQIcon4(fileName string) *QIcon {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QIcon_new4, fileName_p, uintptr(len(fileName)))
	__o := newownedQIcon(__r)
	runtime.KeepAlive(fileName)
	return __o
}

func NewQIcon5(engine *QIconEngine) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_new5, uintptr(engine.UnsafePointer()))
	__o := newownedQIcon(__r)
	runtime.KeepAlive(engine)
	return __o
}

func (this *QIcon) OperatorAssign(other *QIcon) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQIcon(__r)
}

func (this *QIcon) Swap(other *QIcon) {
	purego.SyscallN(addr_QIcon_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QIcon) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QIcon_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QIcon) Pixmap(size *QSize) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap2(w int, h int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap2, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) PixmapWithExtent(extent int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap3, this.h, uintptr(extent))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap3(size *QSize, devicePixelRatio float64) *QPixmap {
	__r := pg_QIcon_pixmap4(this.h, uintptr(size.UnsafePointer()), devicePixelRatio)
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap4(window *QWindow, size *QSize) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap5, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) ActualSize(size *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) ActualSize2(window *QWindow, size *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize2, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIcon) Paint(painter *QPainter, rect *QRect) {
	purego.SyscallN(addr_QIcon_paint, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QIcon) Paint2(painter *QPainter, x int, y int, w int, h int) {
	purego.SyscallN(addr_QIcon_paint2, this.h, uintptr(painter.UnsafePointer()), uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QIcon) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QIcon_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIcon) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QIcon_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIcon) Detach() {
	purego.SyscallN(addr_QIcon_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QIcon) CacheKey() int64 {
	__r0, _, _ := purego.SyscallN(addr_QIcon_cacheKey, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIcon) AddPixmap(pixmap *QPixmap) {
	purego.SyscallN(addr_QIcon_addPixmap, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIcon) AddFile(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIcon_addFile, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QIcon) AvailableSizes() []*QSize {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_availableSizes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QSize, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQSize(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QIcon) SetIsMask(isMask bool) {
	purego.SyscallN(addr_QIcon_setIsMask, this.h, uintptr(b2i(isMask)))
	runtime.KeepAlive(this)
}

func (this *QIcon) IsMask() bool {
	__r0, _, _ := purego.SyscallN(addr_QIcon_isMask, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QIcon_FromTheme(name string) *QIcon {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QIcon_fromTheme, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
	return newownedQIcon(__r)
}

func QIcon_FromTheme2(name string, fallback *QIcon) *QIcon {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QIcon_fromTheme2, name_p, uintptr(len(name)), uintptr(fallback.UnsafePointer()))
	runtime.KeepAlive(name)
	runtime.KeepAlive(fallback)
	return newownedQIcon(__r)
}

func QIcon_HasThemeIcon(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QIcon_hasThemeIcon, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name)
	return __r != 0
}

func QIcon_FromThemeWithIcon(icon QIcon__ThemeIcon) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_fromTheme3, uintptr(icon))
	return newownedQIcon(__r)
}

func QIcon_FromTheme3(icon QIcon__ThemeIcon, fallback *QIcon) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QIcon_fromTheme4, uintptr(icon), uintptr(fallback.UnsafePointer()))
	runtime.KeepAlive(fallback)
	return newownedQIcon(__r)
}

func QIcon_HasThemeIconWithIcon(icon QIcon__ThemeIcon) bool {
	__r0, _, _ := purego.SyscallN(addr_QIcon_hasThemeIcon2, uintptr(icon))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QIcon_ThemeSearchPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_themeSearchPaths, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QIcon_SetThemeSearchPaths(searchpath []string) {
	searchpath_arr := make([]uintptr, len(searchpath))
	for __i, __e := range searchpath {
		searchpath_arr[__i] = pgCString(__e)
	}
	var searchpath_p uintptr
	if len(searchpath_arr) > 0 {
		searchpath_p = uintptr(unsafe.Pointer(&searchpath_arr[0]))
	}
	purego.SyscallN(addr_QIcon_setThemeSearchPaths, searchpath_p, uintptr(len(searchpath)))
	for __i := range searchpath_arr {
		pgFree(searchpath_arr[__i])
	}
	runtime.KeepAlive(searchpath_arr)
}

func QIcon_FallbackSearchPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_fallbackSearchPaths, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QIcon_SetFallbackSearchPaths(paths []string) {
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	purego.SyscallN(addr_QIcon_setFallbackSearchPaths, paths_p, uintptr(len(paths)))
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
}

func QIcon_ThemeName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_themeName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QIcon_SetThemeName(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QIcon_setThemeName, path_p, uintptr(len(path)))
	runtime.KeepAlive(path)
}

func QIcon_FallbackThemeName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_fallbackThemeName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QIcon_SetFallbackThemeName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QIcon_setFallbackThemeName, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
}

func (this *QIcon) Pixmap5(size *QSize, mode QIcon__Mode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap6, this.h, uintptr(size.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap6(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap7, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap7(w int, h int, mode QIcon__Mode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap8, this.h, uintptr(w), uintptr(h), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap8(w int, h int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap9, this.h, uintptr(w), uintptr(h), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap9(extent int, mode QIcon__Mode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap10, this.h, uintptr(extent), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap10(extent int, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap11, this.h, uintptr(extent), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap11(size *QSize, devicePixelRatio float64, mode QIcon__Mode) *QPixmap {
	__r := pg_QIcon_pixmap12(this.h, uintptr(size.UnsafePointer()), devicePixelRatio, int32(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap12(size *QSize, devicePixelRatio float64, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r := pg_QIcon_pixmap13(this.h, uintptr(size.UnsafePointer()), devicePixelRatio, int32(mode), int32(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap13(window *QWindow, size *QSize, mode QIcon__Mode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap14, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) Pixmap14(window *QWindow, size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIcon_pixmap15, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIcon) ActualSize3(size *QSize, mode QIcon__Mode) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize3, this.h, uintptr(size.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) ActualSize4(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize4, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) ActualSize5(window *QWindow, size *QSize, mode QIcon__Mode) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize5, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) ActualSize6(window *QWindow, size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIcon_actualSize6, this.h, uintptr(window.UnsafePointer()), uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIcon) Paint3(painter *QPainter, rect *QRect, alignment AlignmentFlag) {
	purego.SyscallN(addr_QIcon_paint3, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QIcon) Paint4(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode QIcon__Mode) {
	purego.SyscallN(addr_QIcon_paint4, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QIcon) Paint5(painter *QPainter, rect *QRect, alignment AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIcon_paint5, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(alignment), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QIcon) Paint6(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QIcon_paint6, this.h, uintptr(painter.UnsafePointer()), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QIcon) Paint7(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode QIcon__Mode) {
	purego.SyscallN(addr_QIcon_paint7, this.h, uintptr(painter.UnsafePointer()), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(alignment), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QIcon) Paint8(painter *QPainter, x int, y int, w int, h int, alignment AlignmentFlag, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIcon_paint8, this.h, uintptr(painter.UnsafePointer()), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(alignment), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QIcon) AddPixmap2(pixmap *QPixmap, mode QIcon__Mode) {
	purego.SyscallN(addr_QIcon_addPixmap2, this.h, uintptr(pixmap.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIcon) AddPixmap3(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIcon_addPixmap3, this.h, uintptr(pixmap.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIcon) AddFile2(fileName string, size *QSize) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIcon_addFile2, this.h, fileName_p, uintptr(len(fileName)), uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(size)
}

func (this *QIcon) AddFile3(fileName string, size *QSize, mode QIcon__Mode) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIcon_addFile3, this.h, fileName_p, uintptr(len(fileName)), uintptr(size.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(size)
}

func (this *QIcon) AddFile4(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIcon_addFile4, this.h, fileName_p, uintptr(len(fileName)), uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(size)
}

func (this *QIcon) AvailableSizesWithMode(mode QIcon__Mode) []*QSize {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_availableSizes2, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QSize, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQSize(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QIcon) AvailableSizes2(mode QIcon__Mode, state QIcon__State) []*QSize {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIcon_availableSizes3, this.h, uintptr(mode), uintptr(state), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QSize, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQSize(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QIcon) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QIcon_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QIcon) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIcon_delete(this.h)
	}
	this.UnsafeRelease()
}
