package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMessageAuthenticationCode_new uintptr
	addr_QMessageAuthenticationCode_new2 uintptr
	addr_QMessageAuthenticationCode_swap uintptr
	addr_QMessageAuthenticationCode_reset uintptr
	addr_QMessageAuthenticationCode_setKey uintptr
	addr_QMessageAuthenticationCode_addData uintptr
	addr_QMessageAuthenticationCode_addData2 uintptr
	addr_QMessageAuthenticationCode_addData3 uintptr
	addr_QMessageAuthenticationCode_resultView uintptr
	addr_QMessageAuthenticationCode_result uintptr
	addr_QMessageAuthenticationCode_hash uintptr
	pg_QMessageAuthenticationCode_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMessageAuthenticationCode_new, "QMessageAuthenticationCode_new")
	symaddr(&addr_QMessageAuthenticationCode_new2, "QMessageAuthenticationCode_new2")
	symaddr(&addr_QMessageAuthenticationCode_swap, "QMessageAuthenticationCode_swap")
	symaddr(&addr_QMessageAuthenticationCode_reset, "QMessageAuthenticationCode_reset")
	symaddr(&addr_QMessageAuthenticationCode_setKey, "QMessageAuthenticationCode_setKey")
	symaddr(&addr_QMessageAuthenticationCode_addData, "QMessageAuthenticationCode_addData")
	symaddr(&addr_QMessageAuthenticationCode_addData2, "QMessageAuthenticationCode_addData2")
	symaddr(&addr_QMessageAuthenticationCode_addData3, "QMessageAuthenticationCode_addData3")
	symaddr(&addr_QMessageAuthenticationCode_resultView, "QMessageAuthenticationCode_resultView")
	symaddr(&addr_QMessageAuthenticationCode_result, "QMessageAuthenticationCode_result")
	symaddr(&addr_QMessageAuthenticationCode_hash, "QMessageAuthenticationCode_hash")
	reg(&pg_QMessageAuthenticationCode_delete, "QMessageAuthenticationCode_delete")
}

type QMessageAuthenticationCode struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMessageAuthenticationCode) UnsafeGoID() int64     { return this.vid }
func (this *QMessageAuthenticationCode) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMessageAuthenticationCode(h uintptr) *QMessageAuthenticationCode {
	if h == 0 {
		return nil
	}
	return &QMessageAuthenticationCode{h: h}
}

func newownedQMessageAuthenticationCode(h uintptr) *QMessageAuthenticationCode {
	o := bindQMessageAuthenticationCode(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMessageAuthenticationCode).Delete)
	}
	return o
}

func (this *QMessageAuthenticationCode) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMessageAuthenticationCode) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMessageAuthenticationCode(h unsafe.Pointer) *QMessageAuthenticationCode { return bindQMessageAuthenticationCode(uintptr(h)) }
func UnsafeNewQMessageAuthenticationCodeOwned(h unsafe.Pointer) *QMessageAuthenticationCode { return newownedQMessageAuthenticationCode(uintptr(h)) }

func (this *QMessageAuthenticationCode) UnsafeRelease() { this.h = 0 }

func NewQMessageAuthenticationCode(method QCryptographicHash__Algorithm) *QMessageAuthenticationCode {
	__r, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_new, uintptr(method))
	__o := newownedQMessageAuthenticationCode(__r)
	return __o
}

func NewQMessageAuthenticationCode2(method QCryptographicHash__Algorithm, key *QByteArrayView) *QMessageAuthenticationCode {
	__r, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_new2, uintptr(method), uintptr(key.UnsafePointer()))
	__o := newownedQMessageAuthenticationCode(__r)
	runtime.KeepAlive(key)
	return __o
}

func (this *QMessageAuthenticationCode) Swap(other *QMessageAuthenticationCode) {
	purego.SyscallN(addr_QMessageAuthenticationCode_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QMessageAuthenticationCode) Reset() {
	purego.SyscallN(addr_QMessageAuthenticationCode_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QMessageAuthenticationCode) SetKey(key *QByteArrayView) {
	purego.SyscallN(addr_QMessageAuthenticationCode_setKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QMessageAuthenticationCode) AddData(data string, length int64) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QMessageAuthenticationCode_addData, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(length))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QMessageAuthenticationCode) AddDataWithData(data *QByteArrayView) {
	purego.SyscallN(addr_QMessageAuthenticationCode_addData2, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QMessageAuthenticationCode) AddDataWithDevice(device *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_addData3, this.h, uintptr(device.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	return __r != 0
}

func (this *QMessageAuthenticationCode) ResultView() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_resultView, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QMessageAuthenticationCode) Result() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_result, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QMessageAuthenticationCode_Hash(message *QByteArrayView, key *QByteArrayView, method QCryptographicHash__Algorithm) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMessageAuthenticationCode_hash, uintptr(message.UnsafePointer()), uintptr(key.UnsafePointer()), uintptr(method))
	runtime.KeepAlive(message)
	runtime.KeepAlive(key)
	return newownedQByteArray(__r)
}

func (this *QMessageAuthenticationCode) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMessageAuthenticationCode_delete(this.h)
	}
	this.UnsafeRelease()
}
