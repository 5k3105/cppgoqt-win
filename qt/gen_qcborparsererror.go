package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborParserError_errorString uintptr
	addr_QCborParserError_offset_get uintptr
	addr_QCborParserError_offset_set uintptr
	addr_QCborParserError_error_get uintptr
	addr_QCborParserError_error_set uintptr
	pg_QCborParserError_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborParserError_errorString, "QCborParserError_errorString")
	symaddr(&addr_QCborParserError_offset_get, "QCborParserError_offset_get")
	symaddr(&addr_QCborParserError_offset_set, "QCborParserError_offset_set")
	symaddr(&addr_QCborParserError_error_get, "QCborParserError_error_get")
	symaddr(&addr_QCborParserError_error_set, "QCborParserError_error_set")
	reg(&pg_QCborParserError_delete, "QCborParserError_delete")
}

type QCborParserError struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborParserError) UnsafeGoID() int64     { return this.vid }
func (this *QCborParserError) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborParserError(h uintptr) *QCborParserError {
	if h == 0 {
		return nil
	}
	return &QCborParserError{h: h}
}

func newownedQCborParserError(h uintptr) *QCborParserError {
	o := bindQCborParserError(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborParserError).Delete)
	}
	return o
}

func (this *QCborParserError) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborParserError) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborParserError(h unsafe.Pointer) *QCborParserError { return bindQCborParserError(uintptr(h)) }
func UnsafeNewQCborParserErrorOwned(h unsafe.Pointer) *QCborParserError { return newownedQCborParserError(uintptr(h)) }

func (this *QCborParserError) UnsafeRelease() { this.h = 0 }

func (this *QCborParserError) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborParserError_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborParserError) Offset() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborParserError_offset_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborParserError) SetOffset(offset int64) {
	purego.SyscallN(addr_QCborParserError_offset_set, this.h, uintptr(offset))
	runtime.KeepAlive(this)
}

func (this *QCborParserError) Error() *QCborError {
	__r, _, _ := purego.SyscallN(addr_QCborParserError_error_get, this.h)
	runtime.KeepAlive(this)
	return newownedQCborError(__r)
}

func (this *QCborParserError) SetError(error_ *QCborError) {
	purego.SyscallN(addr_QCborParserError_error_set, this.h, uintptr(error_.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(error_)
}

func (this *QCborParserError) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborParserError_delete(this.h)
	}
	this.UnsafeRelease()
}
