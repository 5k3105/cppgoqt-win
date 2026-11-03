package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborError_ToQCborErrorCode uintptr
	addr_QCborError_toString uintptr
	addr_QCborError_c_get uintptr
	addr_QCborError_c_set uintptr
	pg_QCborError_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborError_ToQCborErrorCode, "QCborError_ToQCborErrorCode")
	symaddr(&addr_QCborError_toString, "QCborError_toString")
	symaddr(&addr_QCborError_c_get, "QCborError_c_get")
	symaddr(&addr_QCborError_c_set, "QCborError_c_set")
	reg(&pg_QCborError_delete, "QCborError_delete")
}

type QCborError struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborError) UnsafeGoID() int64     { return this.vid }
func (this *QCborError) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborError(h uintptr) *QCborError {
	if h == 0 {
		return nil
	}
	return &QCborError{h: h}
}

func newownedQCborError(h uintptr) *QCborError {
	o := bindQCborError(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborError).Delete)
	}
	return o
}

func (this *QCborError) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborError) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborError(h unsafe.Pointer) *QCborError { return bindQCborError(uintptr(h)) }
func UnsafeNewQCborErrorOwned(h unsafe.Pointer) *QCborError { return newownedQCborError(uintptr(h)) }

func (this *QCborError) UnsafeRelease() { this.h = 0 }

func (this *QCborError) ToQCborErrorCode() QCborError__Code {
	__r0, _, _ := purego.SyscallN(addr_QCborError_ToQCborErrorCode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborError__Code(__r)
}

func (this *QCborError) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborError_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborError) C() QCborError__Code {
	__r0, _, _ := purego.SyscallN(addr_QCborError_c_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborError__Code(__r)
}

func (this *QCborError) SetC(c QCborError__Code) {
	purego.SyscallN(addr_QCborError_c_set, this.h, uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QCborError) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborError_delete(this.h)
	}
	this.UnsafeRelease()
}
