package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDirListing__DirEntry_new uintptr
	addr_QDirListing__DirEntry_new2 uintptr
	addr_QDirListing__DirEntry_fileName uintptr
	addr_QDirListing__DirEntry_baseName uintptr
	addr_QDirListing__DirEntry_completeBaseName uintptr
	addr_QDirListing__DirEntry_suffix uintptr
	addr_QDirListing__DirEntry_bundleName uintptr
	addr_QDirListing__DirEntry_completeSuffix uintptr
	addr_QDirListing__DirEntry_filePath uintptr
	addr_QDirListing__DirEntry_isDir uintptr
	addr_QDirListing__DirEntry_isFile uintptr
	addr_QDirListing__DirEntry_isSymLink uintptr
	addr_QDirListing__DirEntry_exists uintptr
	addr_QDirListing__DirEntry_isHidden uintptr
	addr_QDirListing__DirEntry_isReadable uintptr
	addr_QDirListing__DirEntry_isWritable uintptr
	addr_QDirListing__DirEntry_isExecutable uintptr
	addr_QDirListing__DirEntry_fileInfo uintptr
	addr_QDirListing__DirEntry_canonicalFilePath uintptr
	addr_QDirListing__DirEntry_absoluteFilePath uintptr
	addr_QDirListing__DirEntry_absolutePath uintptr
	addr_QDirListing__DirEntry_size uintptr
	addr_QDirListing__DirEntry_birthTime uintptr
	addr_QDirListing__DirEntry_metadataChangeTime uintptr
	addr_QDirListing__DirEntry_lastModified uintptr
	addr_QDirListing__DirEntry_lastRead uintptr
	addr_QDirListing__DirEntry_fileTime uintptr
	addr_QDirListing__DirEntry_OperatorAssign uintptr
	pg_QDirListing__DirEntry_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDirListing__DirEntry_new, "QDirListing__DirEntry_new")
	symaddr(&addr_QDirListing__DirEntry_new2, "QDirListing__DirEntry_new2")
	symaddr(&addr_QDirListing__DirEntry_fileName, "QDirListing__DirEntry_fileName")
	symaddr(&addr_QDirListing__DirEntry_baseName, "QDirListing__DirEntry_baseName")
	symaddr(&addr_QDirListing__DirEntry_completeBaseName, "QDirListing__DirEntry_completeBaseName")
	symaddr(&addr_QDirListing__DirEntry_suffix, "QDirListing__DirEntry_suffix")
	symaddr(&addr_QDirListing__DirEntry_bundleName, "QDirListing__DirEntry_bundleName")
	symaddr(&addr_QDirListing__DirEntry_completeSuffix, "QDirListing__DirEntry_completeSuffix")
	symaddr(&addr_QDirListing__DirEntry_filePath, "QDirListing__DirEntry_filePath")
	symaddr(&addr_QDirListing__DirEntry_isDir, "QDirListing__DirEntry_isDir")
	symaddr(&addr_QDirListing__DirEntry_isFile, "QDirListing__DirEntry_isFile")
	symaddr(&addr_QDirListing__DirEntry_isSymLink, "QDirListing__DirEntry_isSymLink")
	symaddr(&addr_QDirListing__DirEntry_exists, "QDirListing__DirEntry_exists")
	symaddr(&addr_QDirListing__DirEntry_isHidden, "QDirListing__DirEntry_isHidden")
	symaddr(&addr_QDirListing__DirEntry_isReadable, "QDirListing__DirEntry_isReadable")
	symaddr(&addr_QDirListing__DirEntry_isWritable, "QDirListing__DirEntry_isWritable")
	symaddr(&addr_QDirListing__DirEntry_isExecutable, "QDirListing__DirEntry_isExecutable")
	symaddr(&addr_QDirListing__DirEntry_fileInfo, "QDirListing__DirEntry_fileInfo")
	symaddr(&addr_QDirListing__DirEntry_canonicalFilePath, "QDirListing__DirEntry_canonicalFilePath")
	symaddr(&addr_QDirListing__DirEntry_absoluteFilePath, "QDirListing__DirEntry_absoluteFilePath")
	symaddr(&addr_QDirListing__DirEntry_absolutePath, "QDirListing__DirEntry_absolutePath")
	symaddr(&addr_QDirListing__DirEntry_size, "QDirListing__DirEntry_size")
	symaddr(&addr_QDirListing__DirEntry_birthTime, "QDirListing__DirEntry_birthTime")
	symaddr(&addr_QDirListing__DirEntry_metadataChangeTime, "QDirListing__DirEntry_metadataChangeTime")
	symaddr(&addr_QDirListing__DirEntry_lastModified, "QDirListing__DirEntry_lastModified")
	symaddr(&addr_QDirListing__DirEntry_lastRead, "QDirListing__DirEntry_lastRead")
	symaddr(&addr_QDirListing__DirEntry_fileTime, "QDirListing__DirEntry_fileTime")
	symaddr(&addr_QDirListing__DirEntry_OperatorAssign, "QDirListing__DirEntry_OperatorAssign")
	reg(&pg_QDirListing__DirEntry_delete, "QDirListing__DirEntry_delete")
}

type QDirListing__DirEntry struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDirListing__DirEntry) UnsafeGoID() int64     { return this.vid }
func (this *QDirListing__DirEntry) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDirListing__DirEntry(h uintptr) *QDirListing__DirEntry {
	if h == 0 {
		return nil
	}
	return &QDirListing__DirEntry{h: h}
}

func newownedQDirListing__DirEntry(h uintptr) *QDirListing__DirEntry {
	o := bindQDirListing__DirEntry(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDirListing__DirEntry).Delete)
	}
	return o
}

func (this *QDirListing__DirEntry) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDirListing__DirEntry) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDirListing__DirEntry(h unsafe.Pointer) *QDirListing__DirEntry { return bindQDirListing__DirEntry(uintptr(h)) }
func UnsafeNewQDirListing__DirEntryOwned(h unsafe.Pointer) *QDirListing__DirEntry { return newownedQDirListing__DirEntry(uintptr(h)) }

func (this *QDirListing__DirEntry) UnsafeRelease() { this.h = 0 }

func NewQDirListing__DirEntry(arg1 *QDirListing__DirEntry) *QDirListing__DirEntry {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQDirListing__DirEntry(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQDirListing__DirEntry2() *QDirListing__DirEntry {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_new2)
	__o := newownedQDirListing__DirEntry(__r)
	return __o
}

func (this *QDirListing__DirEntry) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) BaseName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_baseName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) CompleteBaseName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_completeBaseName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) Suffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_suffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) BundleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_bundleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) CompleteSuffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_completeSuffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) FilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_filePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) IsDir() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isDir, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsFile() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isFile, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsSymLink() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isSymLink, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) Exists() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_exists, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsReadable() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isReadable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsWritable() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isWritable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) IsExecutable() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_isExecutable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirListing__DirEntry) FileInfo() *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_fileInfo, this.h)
	runtime.KeepAlive(this)
	return newownedQFileInfo(__r)
}

func (this *QDirListing__DirEntry) CanonicalFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_canonicalFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) AbsoluteFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_absoluteFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) AbsolutePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_absolutePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirListing__DirEntry) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDirListing__DirEntry) BirthTime(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_birthTime, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QDirListing__DirEntry) MetadataChangeTime(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_metadataChangeTime, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QDirListing__DirEntry) LastModified(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_lastModified, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QDirListing__DirEntry) LastRead(tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_lastRead, this.h, uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QDirListing__DirEntry) FileTime(type_ QFileDevice__FileTime, tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_fileTime, this.h, uintptr(type_), uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QDirListing__DirEntry) OperatorAssign(arg1 *QDirListing__DirEntry) *QDirListing__DirEntry {
	__r, _, _ := purego.SyscallN(addr_QDirListing__DirEntry_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQDirListing__DirEntry(__r)
}

func (this *QDirListing__DirEntry) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDirListing__DirEntry_delete(this.h)
	}
	this.UnsafeRelease()
}
