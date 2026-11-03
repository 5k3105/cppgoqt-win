package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QResource_new uintptr
	addr_QResource_new2 uintptr
	addr_QResource_new3 uintptr
	addr_QResource_setFileName uintptr
	addr_QResource_fileName uintptr
	addr_QResource_absoluteFilePath uintptr
	addr_QResource_setLocale uintptr
	addr_QResource_locale uintptr
	addr_QResource_isValid uintptr
	addr_QResource_compressionAlgorithm uintptr
	addr_QResource_size uintptr
	addr_QResource_uncompressedSize uintptr
	addr_QResource_uncompressedData uintptr
	addr_QResource_lastModified uintptr
	addr_QResource_registerResource uintptr
	addr_QResource_unregisterResource uintptr
	addr_QResource_registerResource2 uintptr
	addr_QResource_unregisterResource2 uintptr
	addr_QResource_registerResource3 uintptr
	addr_QResource_unregisterResource3 uintptr
	addr_QResource_registerResource4 uintptr
	addr_QResource_unregisterResource4 uintptr
	pg_QResource_delete func(uintptr)
)

func init() {
	symaddr(&addr_QResource_new, "QResource_new")
	symaddr(&addr_QResource_new2, "QResource_new2")
	symaddr(&addr_QResource_new3, "QResource_new3")
	symaddr(&addr_QResource_setFileName, "QResource_setFileName")
	symaddr(&addr_QResource_fileName, "QResource_fileName")
	symaddr(&addr_QResource_absoluteFilePath, "QResource_absoluteFilePath")
	symaddr(&addr_QResource_setLocale, "QResource_setLocale")
	symaddr(&addr_QResource_locale, "QResource_locale")
	symaddr(&addr_QResource_isValid, "QResource_isValid")
	symaddr(&addr_QResource_compressionAlgorithm, "QResource_compressionAlgorithm")
	symaddr(&addr_QResource_size, "QResource_size")
	symaddr(&addr_QResource_uncompressedSize, "QResource_uncompressedSize")
	symaddr(&addr_QResource_uncompressedData, "QResource_uncompressedData")
	symaddr(&addr_QResource_lastModified, "QResource_lastModified")
	symaddr(&addr_QResource_registerResource, "QResource_registerResource")
	symaddr(&addr_QResource_unregisterResource, "QResource_unregisterResource")
	symaddr(&addr_QResource_registerResource2, "QResource_registerResource2")
	symaddr(&addr_QResource_unregisterResource2, "QResource_unregisterResource2")
	symaddr(&addr_QResource_registerResource3, "QResource_registerResource3")
	symaddr(&addr_QResource_unregisterResource3, "QResource_unregisterResource3")
	symaddr(&addr_QResource_registerResource4, "QResource_registerResource4")
	symaddr(&addr_QResource_unregisterResource4, "QResource_unregisterResource4")
	reg(&pg_QResource_delete, "QResource_delete")
}

type QResource struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QResource) UnsafeGoID() int64     { return this.vid }
func (this *QResource) UnsafeSetGoID(id int64) { this.vid = id }

func bindQResource(h uintptr) *QResource {
	if h == 0 {
		return nil
	}
	return &QResource{h: h}
}

func newownedQResource(h uintptr) *QResource {
	o := bindQResource(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QResource).Delete)
	}
	return o
}

func (this *QResource) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QResource) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQResource(h unsafe.Pointer) *QResource { return bindQResource(uintptr(h)) }
func UnsafeNewQResourceOwned(h unsafe.Pointer) *QResource { return newownedQResource(uintptr(h)) }

func (this *QResource) UnsafeRelease() { this.h = 0 }

func NewQResource() *QResource {
	__r, _, _ := purego.SyscallN(addr_QResource_new)
	__o := newownedQResource(__r)
	return __o
}

func NewQResource2(file string) *QResource {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r, _, _ := purego.SyscallN(addr_QResource_new2, file_p, uintptr(len(file)))
	__o := newownedQResource(__r)
	runtime.KeepAlive(file)
	return __o
}

func NewQResource3(file string, locale *QLocale) *QResource {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r, _, _ := purego.SyscallN(addr_QResource_new3, file_p, uintptr(len(file)), uintptr(locale.UnsafePointer()))
	__o := newownedQResource(__r)
	runtime.KeepAlive(file)
	runtime.KeepAlive(locale)
	return __o
}

func (this *QResource) SetFileName(file string) {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	purego.SyscallN(addr_QResource_setFileName, this.h, file_p, uintptr(len(file)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
}

func (this *QResource) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QResource_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QResource) AbsoluteFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QResource_absoluteFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QResource) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QResource_setLocale, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QResource) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QResource_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QResource) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QResource_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QResource) CompressionAlgorithm() QResource__Compression {
	__r0, _, _ := purego.SyscallN(addr_QResource_compressionAlgorithm, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QResource__Compression(__r)
}

func (this *QResource) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QResource_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QResource) UncompressedSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QResource_uncompressedSize, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QResource) UncompressedData() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QResource_uncompressedData, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QResource) LastModified() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QResource_lastModified, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func QResource_RegisterResource(rccFilename string) bool {
	var rccFilename_p uintptr
	if len(rccFilename) > 0 {
		rccFilename_p = uintptr(unsafe.Pointer(unsafe.StringData(rccFilename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_registerResource, rccFilename_p, uintptr(len(rccFilename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccFilename)
	return __r != 0
}

func QResource_UnregisterResource(rccFilename string) bool {
	var rccFilename_p uintptr
	if len(rccFilename) > 0 {
		rccFilename_p = uintptr(unsafe.Pointer(unsafe.StringData(rccFilename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_unregisterResource, rccFilename_p, uintptr(len(rccFilename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccFilename)
	return __r != 0
}

func QResource_RegisterResourceWithRccData(rccData []byte) bool {
	var rccData_p uintptr
	if len(rccData) > 0 {
		rccData_p = uintptr(unsafe.Pointer(&rccData[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_registerResource2, rccData_p)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccData)
	return __r != 0
}

func QResource_UnregisterResourceWithRccData(rccData []byte) bool {
	var rccData_p uintptr
	if len(rccData) > 0 {
		rccData_p = uintptr(unsafe.Pointer(&rccData[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_unregisterResource2, rccData_p)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccData)
	return __r != 0
}

func QResource_RegisterResource2(rccFilename string, resourceRoot string) bool {
	var rccFilename_p uintptr
	if len(rccFilename) > 0 {
		rccFilename_p = uintptr(unsafe.Pointer(unsafe.StringData(rccFilename)))
	}
	var resourceRoot_p uintptr
	if len(resourceRoot) > 0 {
		resourceRoot_p = uintptr(unsafe.Pointer(unsafe.StringData(resourceRoot)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_registerResource3, rccFilename_p, uintptr(len(rccFilename)), resourceRoot_p, uintptr(len(resourceRoot)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccFilename)
	runtime.KeepAlive(resourceRoot)
	return __r != 0
}

func QResource_UnregisterResource2(rccFilename string, resourceRoot string) bool {
	var rccFilename_p uintptr
	if len(rccFilename) > 0 {
		rccFilename_p = uintptr(unsafe.Pointer(unsafe.StringData(rccFilename)))
	}
	var resourceRoot_p uintptr
	if len(resourceRoot) > 0 {
		resourceRoot_p = uintptr(unsafe.Pointer(unsafe.StringData(resourceRoot)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_unregisterResource3, rccFilename_p, uintptr(len(rccFilename)), resourceRoot_p, uintptr(len(resourceRoot)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccFilename)
	runtime.KeepAlive(resourceRoot)
	return __r != 0
}

func QResource_RegisterResource3(rccData []byte, resourceRoot string) bool {
	var rccData_p uintptr
	if len(rccData) > 0 {
		rccData_p = uintptr(unsafe.Pointer(&rccData[0]))
	}
	var resourceRoot_p uintptr
	if len(resourceRoot) > 0 {
		resourceRoot_p = uintptr(unsafe.Pointer(unsafe.StringData(resourceRoot)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_registerResource4, rccData_p, resourceRoot_p, uintptr(len(resourceRoot)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccData)
	runtime.KeepAlive(resourceRoot)
	return __r != 0
}

func QResource_UnregisterResource3(rccData []byte, resourceRoot string) bool {
	var rccData_p uintptr
	if len(rccData) > 0 {
		rccData_p = uintptr(unsafe.Pointer(&rccData[0]))
	}
	var resourceRoot_p uintptr
	if len(resourceRoot) > 0 {
		resourceRoot_p = uintptr(unsafe.Pointer(unsafe.StringData(resourceRoot)))
	}
	__r0, _, _ := purego.SyscallN(addr_QResource_unregisterResource4, rccData_p, resourceRoot_p, uintptr(len(resourceRoot)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(rccData)
	runtime.KeepAlive(resourceRoot)
	return __r != 0
}

func (this *QResource) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QResource_delete(this.h)
	}
	this.UnsafeRelease()
}
