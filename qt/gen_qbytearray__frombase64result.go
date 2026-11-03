package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QByteArray__FromBase64Result_new uintptr
	addr_QByteArray__FromBase64Result_new2 uintptr
	addr_QByteArray__FromBase64Result_swap uintptr
	addr_QByteArray__FromBase64Result_ToBool uintptr
	addr_QByteArray__FromBase64Result_OperatorMultiply uintptr
	addr_QByteArray__FromBase64Result_OperatorMultiply2 uintptr
	addr_QByteArray__FromBase64Result_OperatorAssign uintptr
	addr_QByteArray__FromBase64Result_OperatorEqual uintptr
	addr_QByteArray__FromBase64Result_OperatorNotEqual uintptr
	addr_QByteArray__FromBase64Result_decoded_get uintptr
	addr_QByteArray__FromBase64Result_decoded_set uintptr
	addr_QByteArray__FromBase64Result_decodingStatus_get uintptr
	addr_QByteArray__FromBase64Result_decodingStatus_set uintptr
	pg_QByteArray__FromBase64Result_delete func(uintptr)
)

func init() {
	symaddr(&addr_QByteArray__FromBase64Result_new, "QByteArray__FromBase64Result_new")
	symaddr(&addr_QByteArray__FromBase64Result_new2, "QByteArray__FromBase64Result_new2")
	symaddr(&addr_QByteArray__FromBase64Result_swap, "QByteArray__FromBase64Result_swap")
	symaddr(&addr_QByteArray__FromBase64Result_ToBool, "QByteArray__FromBase64Result_ToBool")
	symaddr(&addr_QByteArray__FromBase64Result_OperatorMultiply, "QByteArray__FromBase64Result_OperatorMultiply")
	symaddr(&addr_QByteArray__FromBase64Result_OperatorMultiply2, "QByteArray__FromBase64Result_OperatorMultiply2")
	symaddr(&addr_QByteArray__FromBase64Result_OperatorAssign, "QByteArray__FromBase64Result_OperatorAssign")
	symaddr(&addr_QByteArray__FromBase64Result_OperatorEqual, "QByteArray__FromBase64Result_OperatorEqual")
	symaddr(&addr_QByteArray__FromBase64Result_OperatorNotEqual, "QByteArray__FromBase64Result_OperatorNotEqual")
	symaddr(&addr_QByteArray__FromBase64Result_decoded_get, "QByteArray__FromBase64Result_decoded_get")
	symaddr(&addr_QByteArray__FromBase64Result_decoded_set, "QByteArray__FromBase64Result_decoded_set")
	symaddr(&addr_QByteArray__FromBase64Result_decodingStatus_get, "QByteArray__FromBase64Result_decodingStatus_get")
	symaddr(&addr_QByteArray__FromBase64Result_decodingStatus_set, "QByteArray__FromBase64Result_decodingStatus_set")
	reg(&pg_QByteArray__FromBase64Result_delete, "QByteArray__FromBase64Result_delete")
}

type QByteArray__FromBase64Result struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QByteArray__FromBase64Result) UnsafeGoID() int64     { return this.vid }
func (this *QByteArray__FromBase64Result) UnsafeSetGoID(id int64) { this.vid = id }

func bindQByteArray__FromBase64Result(h uintptr) *QByteArray__FromBase64Result {
	if h == 0 {
		return nil
	}
	return &QByteArray__FromBase64Result{h: h}
}

func newownedQByteArray__FromBase64Result(h uintptr) *QByteArray__FromBase64Result {
	o := bindQByteArray__FromBase64Result(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QByteArray__FromBase64Result).Delete)
	}
	return o
}

func (this *QByteArray__FromBase64Result) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QByteArray__FromBase64Result) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQByteArray__FromBase64Result(h unsafe.Pointer) *QByteArray__FromBase64Result { return bindQByteArray__FromBase64Result(uintptr(h)) }
func UnsafeNewQByteArray__FromBase64ResultOwned(h unsafe.Pointer) *QByteArray__FromBase64Result { return newownedQByteArray__FromBase64Result(uintptr(h)) }

func (this *QByteArray__FromBase64Result) UnsafeRelease() { this.h = 0 }

func NewQByteArray__FromBase64Result(arg1 *QByteArray__FromBase64Result) *QByteArray__FromBase64Result {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQByteArray__FromBase64Result(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQByteArray__FromBase64Result2() *QByteArray__FromBase64Result {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_new2)
	__o := newownedQByteArray__FromBase64Result(__r)
	return __o
}

func (this *QByteArray__FromBase64Result) Swap(other *QByteArray__FromBase64Result) {
	purego.SyscallN(addr_QByteArray__FromBase64Result_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QByteArray__FromBase64Result) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_ToBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QByteArray__FromBase64Result) OperatorMultiply() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_OperatorMultiply, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray__FromBase64Result) OperatorMultiply2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_OperatorMultiply2, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QByteArray__FromBase64Result) OperatorAssign(arg1 *QByteArray__FromBase64Result) *QByteArray__FromBase64Result {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQByteArray__FromBase64Result(__r)
}

func (this *QByteArray__FromBase64Result) OperatorEqual(rhs *QByteArray__FromBase64Result) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray__FromBase64Result) OperatorNotEqual(rhs *QByteArray__FromBase64Result) bool {
	__r0, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QByteArray__FromBase64Result) Decoded() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_decoded_get, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArray__FromBase64Result) SetDecoded(decoded *QByteArray) {
	purego.SyscallN(addr_QByteArray__FromBase64Result_decoded_set, this.h, uintptr(decoded.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(decoded)
}

func (this *QByteArray__FromBase64Result) DecodingStatus() QByteArray__Base64DecodingStatus {
	__r0, _, _ := purego.SyscallN(addr_QByteArray__FromBase64Result_decodingStatus_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QByteArray__Base64DecodingStatus(__r)
}

func (this *QByteArray__FromBase64Result) SetDecodingStatus(decodingStatus QByteArray__Base64DecodingStatus) {
	purego.SyscallN(addr_QByteArray__FromBase64Result_decodingStatus_set, this.h, uintptr(decodingStatus))
	runtime.KeepAlive(this)
}

func (this *QByteArray__FromBase64Result) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QByteArray__FromBase64Result_delete(this.h)
	}
	this.UnsafeRelease()
}
