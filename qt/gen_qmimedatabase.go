package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMimeDatabase_new uintptr
	addr_QMimeDatabase_mimeTypeForName uintptr
	addr_QMimeDatabase_mimeTypeForFile uintptr
	addr_QMimeDatabase_mimeTypeForFile2 uintptr
	addr_QMimeDatabase_mimeTypesForFileName uintptr
	addr_QMimeDatabase_mimeTypeForData uintptr
	addr_QMimeDatabase_mimeTypeForData2 uintptr
	addr_QMimeDatabase_mimeTypeForUrl uintptr
	addr_QMimeDatabase_mimeTypeForFileNameAndData uintptr
	addr_QMimeDatabase_mimeTypeForFileNameAndData2 uintptr
	addr_QMimeDatabase_suffixForFileName uintptr
	addr_QMimeDatabase_allMimeTypes uintptr
	addr_QMimeDatabase_mimeTypeForFile3 uintptr
	addr_QMimeDatabase_mimeTypeForFile4 uintptr
	pg_QMimeDatabase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMimeDatabase_new, "QMimeDatabase_new")
	symaddr(&addr_QMimeDatabase_mimeTypeForName, "QMimeDatabase_mimeTypeForName")
	symaddr(&addr_QMimeDatabase_mimeTypeForFile, "QMimeDatabase_mimeTypeForFile")
	symaddr(&addr_QMimeDatabase_mimeTypeForFile2, "QMimeDatabase_mimeTypeForFile2")
	symaddr(&addr_QMimeDatabase_mimeTypesForFileName, "QMimeDatabase_mimeTypesForFileName")
	symaddr(&addr_QMimeDatabase_mimeTypeForData, "QMimeDatabase_mimeTypeForData")
	symaddr(&addr_QMimeDatabase_mimeTypeForData2, "QMimeDatabase_mimeTypeForData2")
	symaddr(&addr_QMimeDatabase_mimeTypeForUrl, "QMimeDatabase_mimeTypeForUrl")
	symaddr(&addr_QMimeDatabase_mimeTypeForFileNameAndData, "QMimeDatabase_mimeTypeForFileNameAndData")
	symaddr(&addr_QMimeDatabase_mimeTypeForFileNameAndData2, "QMimeDatabase_mimeTypeForFileNameAndData2")
	symaddr(&addr_QMimeDatabase_suffixForFileName, "QMimeDatabase_suffixForFileName")
	symaddr(&addr_QMimeDatabase_allMimeTypes, "QMimeDatabase_allMimeTypes")
	symaddr(&addr_QMimeDatabase_mimeTypeForFile3, "QMimeDatabase_mimeTypeForFile3")
	symaddr(&addr_QMimeDatabase_mimeTypeForFile4, "QMimeDatabase_mimeTypeForFile4")
	reg(&pg_QMimeDatabase_delete, "QMimeDatabase_delete")
}

type QMimeDatabase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMimeDatabase) UnsafeGoID() int64     { return this.vid }
func (this *QMimeDatabase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMimeDatabase(h uintptr) *QMimeDatabase {
	if h == 0 {
		return nil
	}
	return &QMimeDatabase{h: h}
}

func newownedQMimeDatabase(h uintptr) *QMimeDatabase {
	o := bindQMimeDatabase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMimeDatabase).Delete)
	}
	return o
}

func (this *QMimeDatabase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMimeDatabase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMimeDatabase(h unsafe.Pointer) *QMimeDatabase { return bindQMimeDatabase(uintptr(h)) }
func UnsafeNewQMimeDatabaseOwned(h unsafe.Pointer) *QMimeDatabase { return newownedQMimeDatabase(uintptr(h)) }

func (this *QMimeDatabase) UnsafeRelease() { this.h = 0 }

func NewQMimeDatabase() *QMimeDatabase {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_new)
	__o := newownedQMimeDatabase(__r)
	return __o
}

func (this *QMimeDatabase) MimeTypeForName(nameOrAlias string) *QMimeType {
	var nameOrAlias_p uintptr
	if len(nameOrAlias) > 0 {
		nameOrAlias_p = uintptr(unsafe.Pointer(unsafe.StringData(nameOrAlias)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForName, this.h, nameOrAlias_p, uintptr(len(nameOrAlias)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(nameOrAlias)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForFile(fileName string) *QMimeType {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFile, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForFileWithFileInfo(fileInfo *QFileInfo) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFile2, this.h, uintptr(fileInfo.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileInfo)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypesForFileName(fileName string) []*QMimeType {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypesForFileName, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	__out := make([]*QMimeType, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQMimeType(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeDatabase) MimeTypeForData(data *QByteArray) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForDataWithDevice(device *QIODevice) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForData2, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForUrl(url *QUrl) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForUrl, this.h, uintptr(url.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(url)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName string, device *QIODevice) *QMimeType {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFileNameAndData, this.h, fileName_p, uintptr(len(fileName)), uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(device)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForFileNameAndData2(fileName string, data *QByteArray) *QMimeType {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFileNameAndData2, this.h, fileName_p, uintptr(len(fileName)), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(data)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) SuffixForFileName(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_suffixForFileName, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func (this *QMimeDatabase) AllMimeTypes() []*QMimeType {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_allMimeTypes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QMimeType, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQMimeType(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeDatabase) MimeTypeForFile2(fileName string, mode QMimeDatabase__MatchMode) *QMimeType {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFile3, this.h, fileName_p, uintptr(len(fileName)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) MimeTypeForFile3(fileInfo *QFileInfo, mode QMimeDatabase__MatchMode) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeDatabase_mimeTypeForFile4, this.h, uintptr(fileInfo.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileInfo)
	return newownedQMimeType(__r)
}

func (this *QMimeDatabase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMimeDatabase_delete(this.h)
	}
	this.UnsafeRelease()
}
