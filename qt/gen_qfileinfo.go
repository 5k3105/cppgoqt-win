package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFileInfo_new uintptr
	addr_QFileInfo_new2 uintptr
	addr_QFileInfo_new3 uintptr
	addr_QFileInfo_new4 uintptr
	addr_QFileInfo_new5 uintptr
	addr_QFileInfo_OperatorAssign uintptr
	addr_QFileInfo_swap uintptr
	addr_QFileInfo_setFile uintptr
	addr_QFileInfo_setFile2 uintptr
	addr_QFileInfo_setFile3 uintptr
	addr_QFileInfo_exists uintptr
	addr_QFileInfo_exists2 uintptr
	addr_QFileInfo_refresh uintptr
	addr_QFileInfo_filePath uintptr
	addr_QFileInfo_absoluteFilePath uintptr
	addr_QFileInfo_canonicalFilePath uintptr
	addr_QFileInfo_fileName uintptr
	addr_QFileInfo_baseName uintptr
	addr_QFileInfo_completeBaseName uintptr
	addr_QFileInfo_suffix uintptr
	addr_QFileInfo_bundleName uintptr
	addr_QFileInfo_completeSuffix uintptr
	addr_QFileInfo_path uintptr
	addr_QFileInfo_absolutePath uintptr
	addr_QFileInfo_canonicalPath uintptr
	addr_QFileInfo_dir uintptr
	addr_QFileInfo_absoluteDir uintptr
	addr_QFileInfo_isReadable uintptr
	addr_QFileInfo_isWritable uintptr
	addr_QFileInfo_isExecutable uintptr
	addr_QFileInfo_isHidden uintptr
	addr_QFileInfo_isNativePath uintptr
	addr_QFileInfo_isRelative uintptr
	addr_QFileInfo_isAbsolute uintptr
	addr_QFileInfo_makeAbsolute uintptr
	addr_QFileInfo_isFile uintptr
	addr_QFileInfo_isDir uintptr
	addr_QFileInfo_isSymLink uintptr
	addr_QFileInfo_isSymbolicLink uintptr
	addr_QFileInfo_isOther uintptr
	addr_QFileInfo_isShortcut uintptr
	addr_QFileInfo_isAlias uintptr
	addr_QFileInfo_isJunction uintptr
	addr_QFileInfo_isRoot uintptr
	addr_QFileInfo_isBundle uintptr
	addr_QFileInfo_symLinkTarget uintptr
	addr_QFileInfo_readSymLink uintptr
	addr_QFileInfo_junctionTarget uintptr
	addr_QFileInfo_owner uintptr
	addr_QFileInfo_ownerId uintptr
	addr_QFileInfo_group uintptr
	addr_QFileInfo_groupId uintptr
	addr_QFileInfo_permission uintptr
	addr_QFileInfo_permissions uintptr
	addr_QFileInfo_size uintptr
	addr_QFileInfo_birthTime uintptr
	addr_QFileInfo_metadataChangeTime uintptr
	addr_QFileInfo_lastModified uintptr
	addr_QFileInfo_lastRead uintptr
	addr_QFileInfo_fileTime uintptr
	addr_QFileInfo_birthTime2 uintptr
	addr_QFileInfo_metadataChangeTime2 uintptr
	addr_QFileInfo_lastModified2 uintptr
	addr_QFileInfo_lastRead2 uintptr
	addr_QFileInfo_fileTime2 uintptr
	addr_QFileInfo_caching uintptr
	addr_QFileInfo_setCaching uintptr
	addr_QFileInfo_stat uintptr
	addr_QFileInfo_OperatorEqual uintptr
	addr_QFileInfo_OperatorNotEqual uintptr
	pg_QFileInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFileInfo_new, "QFileInfo_new")
	symaddr(&addr_QFileInfo_new2, "QFileInfo_new2")
	symaddr(&addr_QFileInfo_new3, "QFileInfo_new3")
	symaddr(&addr_QFileInfo_new4, "QFileInfo_new4")
	symaddr(&addr_QFileInfo_new5, "QFileInfo_new5")
	symaddr(&addr_QFileInfo_OperatorAssign, "QFileInfo_OperatorAssign")
	symaddr(&addr_QFileInfo_swap, "QFileInfo_swap")
	symaddr(&addr_QFileInfo_setFile, "QFileInfo_setFile")
	symaddr(&addr_QFileInfo_setFile2, "QFileInfo_setFile2")
	symaddr(&addr_QFileInfo_setFile3, "QFileInfo_setFile3")
	symaddr(&addr_QFileInfo_exists, "QFileInfo_exists")
	symaddr(&addr_QFileInfo_exists2, "QFileInfo_exists2")
	symaddr(&addr_QFileInfo_refresh, "QFileInfo_refresh")
	symaddr(&addr_QFileInfo_filePath, "QFileInfo_filePath")
	symaddr(&addr_QFileInfo_absoluteFilePath, "QFileInfo_absoluteFilePath")
	symaddr(&addr_QFileInfo_canonicalFilePath, "QFileInfo_canonicalFilePath")
	symaddr(&addr_QFileInfo_fileName, "QFileInfo_fileName")
	symaddr(&addr_QFileInfo_baseName, "QFileInfo_baseName")
	symaddr(&addr_QFileInfo_completeBaseName, "QFileInfo_completeBaseName")
	symaddr(&addr_QFileInfo_suffix, "QFileInfo_suffix")
	symaddr(&addr_QFileInfo_bundleName, "QFileInfo_bundleName")
	symaddr(&addr_QFileInfo_completeSuffix, "QFileInfo_completeSuffix")
	symaddr(&addr_QFileInfo_path, "QFileInfo_path")
	symaddr(&addr_QFileInfo_absolutePath, "QFileInfo_absolutePath")
	symaddr(&addr_QFileInfo_canonicalPath, "QFileInfo_canonicalPath")
	symaddr(&addr_QFileInfo_dir, "QFileInfo_dir")
	symaddr(&addr_QFileInfo_absoluteDir, "QFileInfo_absoluteDir")
	symaddr(&addr_QFileInfo_isReadable, "QFileInfo_isReadable")
	symaddr(&addr_QFileInfo_isWritable, "QFileInfo_isWritable")
	symaddr(&addr_QFileInfo_isExecutable, "QFileInfo_isExecutable")
	symaddr(&addr_QFileInfo_isHidden, "QFileInfo_isHidden")
	symaddr(&addr_QFileInfo_isNativePath, "QFileInfo_isNativePath")
	symaddr(&addr_QFileInfo_isRelative, "QFileInfo_isRelative")
	symaddr(&addr_QFileInfo_isAbsolute, "QFileInfo_isAbsolute")
	symaddr(&addr_QFileInfo_makeAbsolute, "QFileInfo_makeAbsolute")
	symaddr(&addr_QFileInfo_isFile, "QFileInfo_isFile")
	symaddr(&addr_QFileInfo_isDir, "QFileInfo_isDir")
	symaddr(&addr_QFileInfo_isSymLink, "QFileInfo_isSymLink")
	symaddr(&addr_QFileInfo_isSymbolicLink, "QFileInfo_isSymbolicLink")
	symaddr(&addr_QFileInfo_isOther, "QFileInfo_isOther")
	symaddr(&addr_QFileInfo_isShortcut, "QFileInfo_isShortcut")
	symaddr(&addr_QFileInfo_isAlias, "QFileInfo_isAlias")
	symaddr(&addr_QFileInfo_isJunction, "QFileInfo_isJunction")
	symaddr(&addr_QFileInfo_isRoot, "QFileInfo_isRoot")
	symaddr(&addr_QFileInfo_isBundle, "QFileInfo_isBundle")
	symaddr(&addr_QFileInfo_symLinkTarget, "QFileInfo_symLinkTarget")
	symaddr(&addr_QFileInfo_readSymLink, "QFileInfo_readSymLink")
	symaddr(&addr_QFileInfo_junctionTarget, "QFileInfo_junctionTarget")
	symaddr(&addr_QFileInfo_owner, "QFileInfo_owner")
	symaddr(&addr_QFileInfo_ownerId, "QFileInfo_ownerId")
	symaddr(&addr_QFileInfo_group, "QFileInfo_group")
	symaddr(&addr_QFileInfo_groupId, "QFileInfo_groupId")
	symaddr(&addr_QFileInfo_permission, "QFileInfo_permission")
	symaddr(&addr_QFileInfo_permissions, "QFileInfo_permissions")
	symaddr(&addr_QFileInfo_size, "QFileInfo_size")
	symaddr(&addr_QFileInfo_birthTime, "QFileInfo_birthTime")
	symaddr(&addr_QFileInfo_metadataChangeTime, "QFileInfo_metadataChangeTime")
	symaddr(&addr_QFileInfo_lastModified, "QFileInfo_lastModified")
	symaddr(&addr_QFileInfo_lastRead, "QFileInfo_lastRead")
	symaddr(&addr_QFileInfo_fileTime, "QFileInfo_fileTime")
	symaddr(&addr_QFileInfo_birthTime2, "QFileInfo_birthTime2")
	symaddr(&addr_QFileInfo_metadataChangeTime2, "QFileInfo_metadataChangeTime2")
	symaddr(&addr_QFileInfo_lastModified2, "QFileInfo_lastModified2")
	symaddr(&addr_QFileInfo_lastRead2, "QFileInfo_lastRead2")
	symaddr(&addr_QFileInfo_fileTime2, "QFileInfo_fileTime2")
	symaddr(&addr_QFileInfo_caching, "QFileInfo_caching")
	symaddr(&addr_QFileInfo_setCaching, "QFileInfo_setCaching")
	symaddr(&addr_QFileInfo_stat, "QFileInfo_stat")
	symaddr(&addr_QFileInfo_OperatorEqual, "QFileInfo_OperatorEqual")
	symaddr(&addr_QFileInfo_OperatorNotEqual, "QFileInfo_OperatorNotEqual")
	reg(&pg_QFileInfo_delete, "QFileInfo_delete")
}

type QFileInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFileInfo) UnsafeGoID() int64     { return this.vid }
func (this *QFileInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFileInfo(h uintptr) *QFileInfo {
	if h == 0 {
		return nil
	}
	return &QFileInfo{h: h}
}

func newownedQFileInfo(h uintptr) *QFileInfo {
	o := bindQFileInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFileInfo).Delete)
	}
	return o
}

func (this *QFileInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileInfo(h unsafe.Pointer) *QFileInfo { return bindQFileInfo(uintptr(h)) }
func UnsafeNewQFileInfoOwned(h unsafe.Pointer) *QFileInfo { return newownedQFileInfo(uintptr(h)) }

func (this *QFileInfo) UnsafeRelease() { this.h = 0 }

func NewQFileInfo() *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_new)
	__o := newownedQFileInfo(__r)
	return __o
}

func NewQFileInfo2(file string) *QFileInfo {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileInfo_new2, file_p, uintptr(len(file)))
	__o := newownedQFileInfo(__r)
	runtime.KeepAlive(file)
	return __o
}

func NewQFileInfo3(file *QFileDevice) *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_new3, uintptr(file.UnsafePointer()))
	__o := newownedQFileInfo(__r)
	runtime.KeepAlive(file)
	return __o
}

func NewQFileInfo4(dir *QDir, file string) *QFileInfo {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileInfo_new4, uintptr(dir.UnsafePointer()), file_p, uintptr(len(file)))
	__o := newownedQFileInfo(__r)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(file)
	return __o
}

func NewQFileInfo5(fileinfo *QFileInfo) *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_new5, uintptr(fileinfo.UnsafePointer()))
	__o := newownedQFileInfo(__r)
	runtime.KeepAlive(fileinfo)
	return __o
}

func (this *QFileInfo) OperatorAssign(fileinfo *QFileInfo) *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_OperatorAssign, this.h, uintptr(fileinfo.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileinfo)
	return bindQFileInfo(__r)
}

func (this *QFileInfo) Swap(other *QFileInfo) {
	purego.SyscallN(addr_QFileInfo_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFileInfo) SetFile(file string) {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	purego.SyscallN(addr_QFileInfo_setFile, this.h, file_p, uintptr(len(file)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
}

func (this *QFileInfo) SetFileWithFile(file *QFileDevice) {
	purego.SyscallN(addr_QFileInfo_setFile2, this.h, uintptr(file.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
}

func (this *QFileInfo) SetFile2(dir *QDir, file string) {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	purego.SyscallN(addr_QFileInfo_setFile3, this.h, uintptr(dir.UnsafePointer()), file_p, uintptr(len(file)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dir)
	runtime.KeepAlive(file)
}

func (this *QFileInfo) Exists() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_exists, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFileInfo_Exists(file string) bool {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_exists2, file_p, uintptr(len(file)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(file)
	return __r != 0
}

func (this *QFileInfo) Refresh() {
	purego.SyscallN(addr_QFileInfo_refresh, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileInfo) FilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_filePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) AbsoluteFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_absoluteFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) CanonicalFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_canonicalFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) BaseName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_baseName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) CompleteBaseName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_completeBaseName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) Suffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_suffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) BundleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_bundleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) CompleteSuffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_completeSuffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) Path() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_path, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) AbsolutePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_absolutePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) CanonicalPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_canonicalPath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) Dir() *QDir {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_dir, this.h)
	runtime.KeepAlive(this)
	return newownedQDir(__r)
}

func (this *QFileInfo) AbsoluteDir() *QDir {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_absoluteDir, this.h)
	runtime.KeepAlive(this)
	return newownedQDir(__r)
}

func (this *QFileInfo) IsReadable() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isReadable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsWritable() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isWritable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsExecutable() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isExecutable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsNativePath() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isNativePath, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsRelative() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isRelative, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsAbsolute() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isAbsolute, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) MakeAbsolute() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_makeAbsolute, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsFile() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isFile, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsDir() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isDir, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsSymLink() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isSymLink, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsSymbolicLink() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isSymbolicLink, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsOther() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isOther, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsShortcut() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isShortcut, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsAlias() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isAlias, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsJunction() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isJunction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsRoot() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isRoot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) IsBundle() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_isBundle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) SymLinkTarget() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_symLinkTarget, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) ReadSymLink() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_readSymLink, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) JunctionTarget() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_junctionTarget, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) Owner() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_owner, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) OwnerId() uint {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_ownerId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QFileInfo) Group() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileInfo_group, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileInfo) GroupId() uint {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_groupId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QFileInfo) Permission(permissions QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_permission, this.h, uintptr(permissions))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) Permissions() QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_permissions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__Permission(__r)
}

func (this *QFileInfo) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFileInfo) BirthTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_birthTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) MetadataChangeTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_metadataChangeTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) LastModified() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_lastModified, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) LastRead() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_lastRead, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) FileTime(time QFileDevice__FileTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_fileTime, this.h, uintptr(time))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) BirthTimeWithTz(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_birthTime2, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) MetadataChangeTimeWithTz(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_metadataChangeTime2, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) LastModifiedWithTz(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_lastModified2, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) LastReadWithTz(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_lastRead2, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) FileTime2(time QFileDevice__FileTime, tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileInfo_fileTime2, this.h, uintptr(time), uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileInfo) Caching() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_caching, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileInfo) SetCaching(on bool) {
	purego.SyscallN(addr_QFileInfo_setCaching, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QFileInfo) Stat() {
	purego.SyscallN(addr_QFileInfo_stat, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileInfo) OperatorEqual(rhs *QFileInfo) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFileInfo) OperatorNotEqual(rhs *QFileInfo) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileInfo_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFileInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
