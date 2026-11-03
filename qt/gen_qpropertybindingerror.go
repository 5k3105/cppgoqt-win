package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyBindingError_new uintptr
	addr_QPropertyBindingError_new2 uintptr
	addr_QPropertyBindingError_new3 uintptr
	addr_QPropertyBindingError_new4 uintptr
	addr_QPropertyBindingError_OperatorAssign uintptr
	addr_QPropertyBindingError_hasError uintptr
	addr_QPropertyBindingError_type uintptr
	addr_QPropertyBindingError_description uintptr
	pg_QPropertyBindingError_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyBindingError_new, "QPropertyBindingError_new")
	symaddr(&addr_QPropertyBindingError_new2, "QPropertyBindingError_new2")
	symaddr(&addr_QPropertyBindingError_new3, "QPropertyBindingError_new3")
	symaddr(&addr_QPropertyBindingError_new4, "QPropertyBindingError_new4")
	symaddr(&addr_QPropertyBindingError_OperatorAssign, "QPropertyBindingError_OperatorAssign")
	symaddr(&addr_QPropertyBindingError_hasError, "QPropertyBindingError_hasError")
	symaddr(&addr_QPropertyBindingError_type, "QPropertyBindingError_type")
	symaddr(&addr_QPropertyBindingError_description, "QPropertyBindingError_description")
	reg(&pg_QPropertyBindingError_delete, "QPropertyBindingError_delete")
}

type QPropertyBindingError struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPropertyBindingError) UnsafeGoID() int64     { return this.vid }
func (this *QPropertyBindingError) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPropertyBindingError(h uintptr) *QPropertyBindingError {
	if h == 0 {
		return nil
	}
	return &QPropertyBindingError{h: h}
}

func newownedQPropertyBindingError(h uintptr) *QPropertyBindingError {
	o := bindQPropertyBindingError(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyBindingError).Delete)
	}
	return o
}

func (this *QPropertyBindingError) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyBindingError) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyBindingError(h unsafe.Pointer) *QPropertyBindingError { return bindQPropertyBindingError(uintptr(h)) }
func UnsafeNewQPropertyBindingErrorOwned(h unsafe.Pointer) *QPropertyBindingError { return newownedQPropertyBindingError(uintptr(h)) }

func (this *QPropertyBindingError) UnsafeRelease() { this.h = 0 }

func NewQPropertyBindingError() *QPropertyBindingError {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_new)
	__o := newownedQPropertyBindingError(__r)
	return __o
}

func NewQPropertyBindingError2(type_ QPropertyBindingError__Type) *QPropertyBindingError {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_new2, uintptr(type_))
	__o := newownedQPropertyBindingError(__r)
	return __o
}

func NewQPropertyBindingError3(other *QPropertyBindingError) *QPropertyBindingError {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_new3, uintptr(other.UnsafePointer()))
	__o := newownedQPropertyBindingError(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQPropertyBindingError4(type_ QPropertyBindingError__Type, description string) *QPropertyBindingError {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_new4, uintptr(type_), description_p, uintptr(len(description)))
	__o := newownedQPropertyBindingError(__r)
	runtime.KeepAlive(description)
	return __o
}

func (this *QPropertyBindingError) OperatorAssign(other *QPropertyBindingError) *QPropertyBindingError {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPropertyBindingError(__r)
}

func (this *QPropertyBindingError) HasError() bool {
	__r0, _, _ := purego.SyscallN(addr_QPropertyBindingError_hasError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPropertyBindingError) Type() QPropertyBindingError__Type {
	__r0, _, _ := purego.SyscallN(addr_QPropertyBindingError_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPropertyBindingError__Type(__r)
}

func (this *QPropertyBindingError) Description() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingError_description, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPropertyBindingError) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyBindingError_delete(this.h)
	}
	this.UnsafeRelease()
}
