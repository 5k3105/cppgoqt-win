package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QJsonParseError_errorString uintptr
	addr_QJsonParseError_offset_get uintptr
	addr_QJsonParseError_offset_set uintptr
	addr_QJsonParseError_error_get uintptr
	addr_QJsonParseError_error_set uintptr
	pg_QJsonParseError_delete func(uintptr)
)

func init() {
	symaddr(&addr_QJsonParseError_errorString, "QJsonParseError_errorString")
	symaddr(&addr_QJsonParseError_offset_get, "QJsonParseError_offset_get")
	symaddr(&addr_QJsonParseError_offset_set, "QJsonParseError_offset_set")
	symaddr(&addr_QJsonParseError_error_get, "QJsonParseError_error_get")
	symaddr(&addr_QJsonParseError_error_set, "QJsonParseError_error_set")
	reg(&pg_QJsonParseError_delete, "QJsonParseError_delete")
}

type QJsonParseError struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QJsonParseError) UnsafeGoID() int64     { return this.vid }
func (this *QJsonParseError) UnsafeSetGoID(id int64) { this.vid = id }

func bindQJsonParseError(h uintptr) *QJsonParseError {
	if h == 0 {
		return nil
	}
	return &QJsonParseError{h: h}
}

func newownedQJsonParseError(h uintptr) *QJsonParseError {
	o := bindQJsonParseError(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QJsonParseError).Delete)
	}
	return o
}

func (this *QJsonParseError) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QJsonParseError) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQJsonParseError(h unsafe.Pointer) *QJsonParseError { return bindQJsonParseError(uintptr(h)) }
func UnsafeNewQJsonParseErrorOwned(h unsafe.Pointer) *QJsonParseError { return newownedQJsonParseError(uintptr(h)) }

func (this *QJsonParseError) UnsafeRelease() { this.h = 0 }

func (this *QJsonParseError) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QJsonParseError_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QJsonParseError) Offset() int {
	__r0, _, _ := purego.SyscallN(addr_QJsonParseError_offset_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QJsonParseError) SetOffset(offset int) {
	purego.SyscallN(addr_QJsonParseError_offset_set, this.h, uintptr(offset))
	runtime.KeepAlive(this)
}

func (this *QJsonParseError) Error() QJsonParseError__ParseError {
	__r0, _, _ := purego.SyscallN(addr_QJsonParseError_error_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QJsonParseError__ParseError(__r)
}

func (this *QJsonParseError) SetError(error_ QJsonParseError__ParseError) {
	purego.SyscallN(addr_QJsonParseError_error_set, this.h, uintptr(error_))
	runtime.KeepAlive(this)
}

func (this *QJsonParseError) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QJsonParseError_delete(this.h)
	}
	this.UnsafeRelease()
}
