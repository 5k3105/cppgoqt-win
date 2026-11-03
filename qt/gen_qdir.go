package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDir_new uintptr
	addr_QDir_new2 uintptr
	addr_QDir_new3 uintptr
	addr_QDir_new4 uintptr
	addr_QDir_new5 uintptr
	addr_QDir_new6 uintptr
	addr_QDir_OperatorAssign uintptr
	addr_QDir_swap uintptr
	addr_QDir_setPath uintptr
	addr_QDir_path uintptr
	addr_QDir_absolutePath uintptr
	addr_QDir_canonicalPath uintptr
	addr_QDir_setSearchPaths uintptr
	addr_QDir_addSearchPath uintptr
	addr_QDir_searchPaths uintptr
	addr_QDir_dirName uintptr
	addr_QDir_filePath uintptr
	addr_QDir_absoluteFilePath uintptr
	addr_QDir_relativeFilePath uintptr
	addr_QDir_toNativeSeparators uintptr
	addr_QDir_fromNativeSeparators uintptr
	addr_QDir_cd uintptr
	addr_QDir_cdUp uintptr
	addr_QDir_nameFilters uintptr
	addr_QDir_setNameFilters uintptr
	addr_QDir_filter uintptr
	addr_QDir_setFilter uintptr
	addr_QDir_sorting uintptr
	addr_QDir_setSorting uintptr
	addr_QDir_count uintptr
	addr_QDir_isEmpty uintptr
	addr_QDir_OperatorSubscript uintptr
	addr_QDir_nameFiltersFromString uintptr
	addr_QDir_entryList uintptr
	addr_QDir_entryList2 uintptr
	addr_QDir_entryInfoList uintptr
	addr_QDir_entryInfoList2 uintptr
	addr_QDir_mkdir uintptr
	addr_QDir_rmdir uintptr
	addr_QDir_mkpath uintptr
	addr_QDir_rmpath uintptr
	addr_QDir_removeRecursively uintptr
	addr_QDir_isReadable uintptr
	addr_QDir_exists uintptr
	addr_QDir_isRoot uintptr
	addr_QDir_isRelativePath uintptr
	addr_QDir_isAbsolutePath uintptr
	addr_QDir_isRelative uintptr
	addr_QDir_isAbsolute uintptr
	addr_QDir_makeAbsolute uintptr
	addr_QDir_remove uintptr
	addr_QDir_rename uintptr
	addr_QDir_exists2 uintptr
	addr_QDir_drives uintptr
	addr_QDir_listSeparator uintptr
	addr_QDir_separator uintptr
	addr_QDir_setCurrent uintptr
	addr_QDir_current uintptr
	addr_QDir_currentPath uintptr
	addr_QDir_home uintptr
	addr_QDir_homePath uintptr
	addr_QDir_root uintptr
	addr_QDir_rootPath uintptr
	addr_QDir_temp uintptr
	addr_QDir_tempPath uintptr
	addr_QDir_match uintptr
	addr_QDir_match2 uintptr
	addr_QDir_cleanPath uintptr
	addr_QDir_refresh uintptr
	addr_QDir_isEmpty2 uintptr
	addr_QDir_entryList3 uintptr
	addr_QDir_entryList4 uintptr
	addr_QDir_entryList5 uintptr
	addr_QDir_entryList6 uintptr
	addr_QDir_entryInfoList3 uintptr
	addr_QDir_entryInfoList4 uintptr
	addr_QDir_entryInfoList5 uintptr
	addr_QDir_entryInfoList6 uintptr
	addr_QDir_OperatorEqual uintptr
	addr_QDir_OperatorNotEqual uintptr
	pg_QDir_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDir_new, "QDir_new")
	symaddr(&addr_QDir_new2, "QDir_new2")
	symaddr(&addr_QDir_new3, "QDir_new3")
	symaddr(&addr_QDir_new4, "QDir_new4")
	symaddr(&addr_QDir_new5, "QDir_new5")
	symaddr(&addr_QDir_new6, "QDir_new6")
	symaddr(&addr_QDir_OperatorAssign, "QDir_OperatorAssign")
	symaddr(&addr_QDir_swap, "QDir_swap")
	symaddr(&addr_QDir_setPath, "QDir_setPath")
	symaddr(&addr_QDir_path, "QDir_path")
	symaddr(&addr_QDir_absolutePath, "QDir_absolutePath")
	symaddr(&addr_QDir_canonicalPath, "QDir_canonicalPath")
	symaddr(&addr_QDir_setSearchPaths, "QDir_setSearchPaths")
	symaddr(&addr_QDir_addSearchPath, "QDir_addSearchPath")
	symaddr(&addr_QDir_searchPaths, "QDir_searchPaths")
	symaddr(&addr_QDir_dirName, "QDir_dirName")
	symaddr(&addr_QDir_filePath, "QDir_filePath")
	symaddr(&addr_QDir_absoluteFilePath, "QDir_absoluteFilePath")
	symaddr(&addr_QDir_relativeFilePath, "QDir_relativeFilePath")
	symaddr(&addr_QDir_toNativeSeparators, "QDir_toNativeSeparators")
	symaddr(&addr_QDir_fromNativeSeparators, "QDir_fromNativeSeparators")
	symaddr(&addr_QDir_cd, "QDir_cd")
	symaddr(&addr_QDir_cdUp, "QDir_cdUp")
	symaddr(&addr_QDir_nameFilters, "QDir_nameFilters")
	symaddr(&addr_QDir_setNameFilters, "QDir_setNameFilters")
	symaddr(&addr_QDir_filter, "QDir_filter")
	symaddr(&addr_QDir_setFilter, "QDir_setFilter")
	symaddr(&addr_QDir_sorting, "QDir_sorting")
	symaddr(&addr_QDir_setSorting, "QDir_setSorting")
	symaddr(&addr_QDir_count, "QDir_count")
	symaddr(&addr_QDir_isEmpty, "QDir_isEmpty")
	symaddr(&addr_QDir_OperatorSubscript, "QDir_OperatorSubscript")
	symaddr(&addr_QDir_nameFiltersFromString, "QDir_nameFiltersFromString")
	symaddr(&addr_QDir_entryList, "QDir_entryList")
	symaddr(&addr_QDir_entryList2, "QDir_entryList2")
	symaddr(&addr_QDir_entryInfoList, "QDir_entryInfoList")
	symaddr(&addr_QDir_entryInfoList2, "QDir_entryInfoList2")
	symaddr(&addr_QDir_mkdir, "QDir_mkdir")
	symaddr(&addr_QDir_rmdir, "QDir_rmdir")
	symaddr(&addr_QDir_mkpath, "QDir_mkpath")
	symaddr(&addr_QDir_rmpath, "QDir_rmpath")
	symaddr(&addr_QDir_removeRecursively, "QDir_removeRecursively")
	symaddr(&addr_QDir_isReadable, "QDir_isReadable")
	symaddr(&addr_QDir_exists, "QDir_exists")
	symaddr(&addr_QDir_isRoot, "QDir_isRoot")
	symaddr(&addr_QDir_isRelativePath, "QDir_isRelativePath")
	symaddr(&addr_QDir_isAbsolutePath, "QDir_isAbsolutePath")
	symaddr(&addr_QDir_isRelative, "QDir_isRelative")
	symaddr(&addr_QDir_isAbsolute, "QDir_isAbsolute")
	symaddr(&addr_QDir_makeAbsolute, "QDir_makeAbsolute")
	symaddr(&addr_QDir_remove, "QDir_remove")
	symaddr(&addr_QDir_rename, "QDir_rename")
	symaddr(&addr_QDir_exists2, "QDir_exists2")
	symaddr(&addr_QDir_drives, "QDir_drives")
	symaddr(&addr_QDir_listSeparator, "QDir_listSeparator")
	symaddr(&addr_QDir_separator, "QDir_separator")
	symaddr(&addr_QDir_setCurrent, "QDir_setCurrent")
	symaddr(&addr_QDir_current, "QDir_current")
	symaddr(&addr_QDir_currentPath, "QDir_currentPath")
	symaddr(&addr_QDir_home, "QDir_home")
	symaddr(&addr_QDir_homePath, "QDir_homePath")
	symaddr(&addr_QDir_root, "QDir_root")
	symaddr(&addr_QDir_rootPath, "QDir_rootPath")
	symaddr(&addr_QDir_temp, "QDir_temp")
	symaddr(&addr_QDir_tempPath, "QDir_tempPath")
	symaddr(&addr_QDir_match, "QDir_match")
	symaddr(&addr_QDir_match2, "QDir_match2")
	symaddr(&addr_QDir_cleanPath, "QDir_cleanPath")
	symaddr(&addr_QDir_refresh, "QDir_refresh")
	symaddr(&addr_QDir_isEmpty2, "QDir_isEmpty2")
	symaddr(&addr_QDir_entryList3, "QDir_entryList3")
	symaddr(&addr_QDir_entryList4, "QDir_entryList4")
	symaddr(&addr_QDir_entryList5, "QDir_entryList5")
	symaddr(&addr_QDir_entryList6, "QDir_entryList6")
	symaddr(&addr_QDir_entryInfoList3, "QDir_entryInfoList3")
	symaddr(&addr_QDir_entryInfoList4, "QDir_entryInfoList4")
	symaddr(&addr_QDir_entryInfoList5, "QDir_entryInfoList5")
	symaddr(&addr_QDir_entryInfoList6, "QDir_entryInfoList6")
	symaddr(&addr_QDir_OperatorEqual, "QDir_OperatorEqual")
	symaddr(&addr_QDir_OperatorNotEqual, "QDir_OperatorNotEqual")
	reg(&pg_QDir_delete, "QDir_delete")
}

type QDir struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDir) UnsafeGoID() int64     { return this.vid }
func (this *QDir) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDir(h uintptr) *QDir {
	if h == 0 {
		return nil
	}
	return &QDir{h: h}
}

func newownedQDir(h uintptr) *QDir {
	o := bindQDir(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDir).Delete)
	}
	return o
}

func (this *QDir) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDir) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDir(h unsafe.Pointer) *QDir { return bindQDir(uintptr(h)) }
func UnsafeNewQDirOwned(h unsafe.Pointer) *QDir { return newownedQDir(uintptr(h)) }

func (this *QDir) UnsafeRelease() { this.h = 0 }

func NewQDir(arg1 *QDir) *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQDir(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQDir2() *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_new2)
	__o := newownedQDir(__r)
	return __o
}

func NewQDir3(path string, nameFilter string) *QDir {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var nameFilter_p uintptr
	if len(nameFilter) > 0 {
		nameFilter_p = uintptr(unsafe.Pointer(unsafe.StringData(nameFilter)))
	}
	__r, _, _ := purego.SyscallN(addr_QDir_new3, path_p, uintptr(len(path)), nameFilter_p, uintptr(len(nameFilter)))
	__o := newownedQDir(__r)
	runtime.KeepAlive(path)
	runtime.KeepAlive(nameFilter)
	return __o
}

func NewQDir4(path string) *QDir {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDir_new4, path_p, uintptr(len(path)))
	__o := newownedQDir(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDir5(path string, nameFilter string, sort QDir__SortFlag) *QDir {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var nameFilter_p uintptr
	if len(nameFilter) > 0 {
		nameFilter_p = uintptr(unsafe.Pointer(unsafe.StringData(nameFilter)))
	}
	__r, _, _ := purego.SyscallN(addr_QDir_new5, path_p, uintptr(len(path)), nameFilter_p, uintptr(len(nameFilter)), uintptr(sort))
	__o := newownedQDir(__r)
	runtime.KeepAlive(path)
	runtime.KeepAlive(nameFilter)
	return __o
}

func NewQDir6(path string, nameFilter string, sort QDir__SortFlag, filter QDir__Filter) *QDir {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var nameFilter_p uintptr
	if len(nameFilter) > 0 {
		nameFilter_p = uintptr(unsafe.Pointer(unsafe.StringData(nameFilter)))
	}
	__r, _, _ := purego.SyscallN(addr_QDir_new6, path_p, uintptr(len(path)), nameFilter_p, uintptr(len(nameFilter)), uintptr(sort), uintptr(filter))
	__o := newownedQDir(__r)
	runtime.KeepAlive(path)
	runtime.KeepAlive(nameFilter)
	return __o
}

func (this *QDir) OperatorAssign(arg1 *QDir) *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDir(__r)
}

func (this *QDir) Swap(other *QDir) {
	purego.SyscallN(addr_QDir_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QDir) SetPath(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QDir_setPath, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QDir) Path() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_path, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDir) AbsolutePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_absolutePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDir) CanonicalPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_canonicalPath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QDir_SetSearchPaths(prefix string, searchPaths []string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	searchPaths_arr := make([]uintptr, len(searchPaths))
	for __i, __e := range searchPaths {
		searchPaths_arr[__i] = pgCString(__e)
	}
	var searchPaths_p uintptr
	if len(searchPaths_arr) > 0 {
		searchPaths_p = uintptr(unsafe.Pointer(&searchPaths_arr[0]))
	}
	purego.SyscallN(addr_QDir_setSearchPaths, prefix_p, uintptr(len(prefix)), searchPaths_p, uintptr(len(searchPaths)))
	runtime.KeepAlive(prefix)
	for __i := range searchPaths_arr {
		pgFree(searchPaths_arr[__i])
	}
	runtime.KeepAlive(searchPaths_arr)
}

func QDir_AddSearchPath(prefix string, path string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QDir_addSearchPath, prefix_p, uintptr(len(prefix)), path_p, uintptr(len(path)))
	runtime.KeepAlive(prefix)
	runtime.KeepAlive(path)
}

func QDir_SearchPaths(prefix string) []string {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_searchPaths, prefix_p, uintptr(len(prefix)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(prefix)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) DirName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_dirName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDir) FilePath(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_filePath, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func (this *QDir) AbsoluteFilePath(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_absoluteFilePath, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func (this *QDir) RelativeFilePath(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_relativeFilePath, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func QDir_ToNativeSeparators(pathName string) string {
	var pathName_p uintptr
	if len(pathName) > 0 {
		pathName_p = uintptr(unsafe.Pointer(unsafe.StringData(pathName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_toNativeSeparators, pathName_p, uintptr(len(pathName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(pathName)
	return pgStr(__r, __n)
}

func QDir_FromNativeSeparators(pathName string) string {
	var pathName_p uintptr
	if len(pathName) > 0 {
		pathName_p = uintptr(unsafe.Pointer(unsafe.StringData(pathName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_fromNativeSeparators, pathName_p, uintptr(len(pathName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(pathName)
	return pgStr(__r, __n)
}

func (this *QDir) Cd(dirName string) bool {
	var dirName_p uintptr
	if len(dirName) > 0 {
		dirName_p = uintptr(unsafe.Pointer(unsafe.StringData(dirName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_cd, this.h, dirName_p, uintptr(len(dirName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dirName)
	return __r != 0
}

func (this *QDir) CdUp() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_cdUp, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) NameFilters() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_nameFilters, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QDir) SetNameFilters(nameFilters []string) {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	purego.SyscallN(addr_QDir_setNameFilters, this.h, nameFilters_p, uintptr(len(nameFilters)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
}

func (this *QDir) Filter() QDir__Filter {
	__r0, _, _ := purego.SyscallN(addr_QDir_filter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDir__Filter(__r)
}

func (this *QDir) SetFilter(filter QDir__Filter) {
	purego.SyscallN(addr_QDir_setFilter, this.h, uintptr(filter))
	runtime.KeepAlive(this)
}

func (this *QDir) Sorting() QDir__SortFlag {
	__r0, _, _ := purego.SyscallN(addr_QDir_sorting, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDir__SortFlag(__r)
}

func (this *QDir) SetSorting(sort QDir__SortFlag) {
	purego.SyscallN(addr_QDir_setSorting, this.h, uintptr(sort))
	runtime.KeepAlive(this)
}

func (this *QDir) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDir_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDir) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) OperatorSubscript(arg1 int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_OperatorSubscript, this.h, uintptr(arg1), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QDir_NameFiltersFromString(nameFilter string) []string {
	var nameFilter_p uintptr
	if len(nameFilter) > 0 {
		nameFilter_p = uintptr(unsafe.Pointer(unsafe.StringData(nameFilter)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_nameFiltersFromString, nameFilter_p, uintptr(len(nameFilter)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(nameFilter)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryList() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QDir) EntryListWithNameFilters(nameFilters []string) []string {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList2, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoList() []*QFileInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoListWithNameFilters(nameFilters []string) []*QFileInfo {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList2, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) Mkdir(dirName string) bool {
	var dirName_p uintptr
	if len(dirName) > 0 {
		dirName_p = uintptr(unsafe.Pointer(unsafe.StringData(dirName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_mkdir, this.h, dirName_p, uintptr(len(dirName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dirName)
	return __r != 0
}

func (this *QDir) Rmdir(dirName string) bool {
	var dirName_p uintptr
	if len(dirName) > 0 {
		dirName_p = uintptr(unsafe.Pointer(unsafe.StringData(dirName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_rmdir, this.h, dirName_p, uintptr(len(dirName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dirName)
	return __r != 0
}

func (this *QDir) Mkpath(dirPath string) bool {
	var dirPath_p uintptr
	if len(dirPath) > 0 {
		dirPath_p = uintptr(unsafe.Pointer(unsafe.StringData(dirPath)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_mkpath, this.h, dirPath_p, uintptr(len(dirPath)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dirPath)
	return __r != 0
}

func (this *QDir) Rmpath(dirPath string) bool {
	var dirPath_p uintptr
	if len(dirPath) > 0 {
		dirPath_p = uintptr(unsafe.Pointer(unsafe.StringData(dirPath)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_rmpath, this.h, dirPath_p, uintptr(len(dirPath)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dirPath)
	return __r != 0
}

func (this *QDir) RemoveRecursively() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_removeRecursively, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) IsReadable() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isReadable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) Exists() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_exists, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) IsRoot() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isRoot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDir_IsRelativePath(path string) bool {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_isRelativePath, path_p, uintptr(len(path)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(path)
	return __r != 0
}

func QDir_IsAbsolutePath(path string) bool {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_isAbsolutePath, path_p, uintptr(len(path)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QDir) IsRelative() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isRelative, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) IsAbsolute() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isAbsolute, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) MakeAbsolute() bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_makeAbsolute, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) Remove(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_remove, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QDir) Rename(oldName string, newName string) bool {
	var oldName_p uintptr
	if len(oldName) > 0 {
		oldName_p = uintptr(unsafe.Pointer(unsafe.StringData(oldName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_rename, this.h, oldName_p, uintptr(len(oldName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(oldName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func (this *QDir) ExistsWithName(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_exists2, this.h, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r != 0
}

func QDir_Drives() []*QFileInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_drives, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QDir_ListSeparator() *QChar {
	__r, _, _ := purego.SyscallN(addr_QDir_listSeparator)
	return newownedQChar(__r)
}

func QDir_Separator() *QChar {
	__r, _, _ := purego.SyscallN(addr_QDir_separator)
	return newownedQChar(__r)
}

func QDir_SetCurrent(path string) bool {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_setCurrent, path_p, uintptr(len(path)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(path)
	return __r != 0
}

func QDir_Current() *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_current)
	return newownedQDir(__r)
}

func QDir_CurrentPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_currentPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QDir_Home() *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_home)
	return newownedQDir(__r)
}

func QDir_HomePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_homePath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QDir_Root() *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_root)
	return newownedQDir(__r)
}

func QDir_RootPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_rootPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QDir_Temp() *QDir {
	__r, _, _ := purego.SyscallN(addr_QDir_temp)
	return newownedQDir(__r)
}

func QDir_TempPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_tempPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QDir_Match(filters []string, fileName string) bool {
	filters_arr := make([]uintptr, len(filters))
	for __i, __e := range filters {
		filters_arr[__i] = pgCString(__e)
	}
	var filters_p uintptr
	if len(filters_arr) > 0 {
		filters_p = uintptr(unsafe.Pointer(&filters_arr[0]))
	}
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_match, filters_p, uintptr(len(filters)), fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	for __i := range filters_arr {
		pgFree(filters_arr[__i])
	}
	runtime.KeepAlive(filters_arr)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QDir_Match2(filter string, fileName string) bool {
	var filter_p uintptr
	if len(filter) > 0 {
		filter_p = uintptr(unsafe.Pointer(unsafe.StringData(filter)))
	}
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QDir_match2, filter_p, uintptr(len(filter)), fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filter)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QDir_CleanPath(path string) string {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_cleanPath, path_p, uintptr(len(path)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(path)
	return pgStr(__r, __n)
}

func (this *QDir) Refresh() {
	purego.SyscallN(addr_QDir_refresh, this.h)
	runtime.KeepAlive(this)
}

func (this *QDir) IsEmptyWithFilters(filters QDir__Filter) bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_isEmpty2, this.h, uintptr(filters))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDir) EntryListWithFilters(filters QDir__Filter) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList3, this.h, uintptr(filters), uintptr(unsafe.Pointer(&__n)))
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

func (this *QDir) EntryList2(filters QDir__Filter, sort QDir__SortFlag) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList4, this.h, uintptr(filters), uintptr(sort), uintptr(unsafe.Pointer(&__n)))
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

func (this *QDir) EntryList3(nameFilters []string, filters QDir__Filter) []string {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList5, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(filters), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryList4(nameFilters []string, filters QDir__Filter, sort QDir__SortFlag) []string {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryList6, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(filters), uintptr(sort), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoListWithFilters(filters QDir__Filter) []*QFileInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList3, this.h, uintptr(filters), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoList2(filters QDir__Filter, sort QDir__SortFlag) []*QFileInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList4, this.h, uintptr(filters), uintptr(sort), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoList3(nameFilters []string, filters QDir__Filter) []*QFileInfo {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList5, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(filters), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) EntryInfoList4(nameFilters []string, filters QDir__Filter, sort QDir__SortFlag) []*QFileInfo {
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDir_entryInfoList6, this.h, nameFilters_p, uintptr(len(nameFilters)), uintptr(filters), uintptr(sort), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	__out := make([]*QFileInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFileInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QDir) OperatorEqual(rhs *QDir) bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDir) OperatorNotEqual(rhs *QDir) bool {
	__r0, _, _ := purego.SyscallN(addr_QDir_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDir) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDir_delete(this.h)
	}
	this.UnsafeRelease()
}
