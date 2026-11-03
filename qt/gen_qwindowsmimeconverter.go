package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWindowsMimeConverter_registerMimeType uintptr
	pg_QWindowsMimeConverter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWindowsMimeConverter_registerMimeType, "QWindowsMimeConverter_registerMimeType")
	reg(&pg_QWindowsMimeConverter_delete, "QWindowsMimeConverter_delete")
}

type QWindowsMimeConverter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QWindowsMimeConverter) UnsafeGoID() int64     { return this.vid }
func (this *QWindowsMimeConverter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQWindowsMimeConverter(h uintptr) *QWindowsMimeConverter {
	if h == 0 {
		return nil
	}
	return &QWindowsMimeConverter{h: h}
}

func newownedQWindowsMimeConverter(h uintptr) *QWindowsMimeConverter {
	o := bindQWindowsMimeConverter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWindowsMimeConverter).Delete)
	}
	return o
}

func (this *QWindowsMimeConverter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWindowsMimeConverter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWindowsMimeConverter(h unsafe.Pointer) *QWindowsMimeConverter { return bindQWindowsMimeConverter(uintptr(h)) }
func UnsafeNewQWindowsMimeConverterOwned(h unsafe.Pointer) *QWindowsMimeConverter { return newownedQWindowsMimeConverter(uintptr(h)) }

func (this *QWindowsMimeConverter) UnsafeRelease() { this.h = 0 }

func QWindowsMimeConverter_RegisterMimeType(mimeType string) int {
	var mimeType_p uintptr
	if len(mimeType) > 0 {
		mimeType_p = uintptr(unsafe.Pointer(unsafe.StringData(mimeType)))
	}
	__r0, _, _ := purego.SyscallN(addr_QWindowsMimeConverter_registerMimeType, mimeType_p, uintptr(len(mimeType)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(mimeType)
	return int(__r)
}

func (this *QWindowsMimeConverter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWindowsMimeConverter_delete(this.h)
	}
	this.UnsafeRelease()
}
