package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCryptographicHash_new uintptr
	addr_QCryptographicHash_swap uintptr
	addr_QCryptographicHash_reset uintptr
	addr_QCryptographicHash_algorithm uintptr
	addr_QCryptographicHash_addData uintptr
	addr_QCryptographicHash_addData2 uintptr
	addr_QCryptographicHash_addData3 uintptr
	addr_QCryptographicHash_result uintptr
	addr_QCryptographicHash_resultView uintptr
	addr_QCryptographicHash_hash uintptr
	addr_QCryptographicHash_hashLength uintptr
	addr_QCryptographicHash_supportsAlgorithm uintptr
	pg_QCryptographicHash_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCryptographicHash_new, "QCryptographicHash_new")
	symaddr(&addr_QCryptographicHash_swap, "QCryptographicHash_swap")
	symaddr(&addr_QCryptographicHash_reset, "QCryptographicHash_reset")
	symaddr(&addr_QCryptographicHash_algorithm, "QCryptographicHash_algorithm")
	symaddr(&addr_QCryptographicHash_addData, "QCryptographicHash_addData")
	symaddr(&addr_QCryptographicHash_addData2, "QCryptographicHash_addData2")
	symaddr(&addr_QCryptographicHash_addData3, "QCryptographicHash_addData3")
	symaddr(&addr_QCryptographicHash_result, "QCryptographicHash_result")
	symaddr(&addr_QCryptographicHash_resultView, "QCryptographicHash_resultView")
	symaddr(&addr_QCryptographicHash_hash, "QCryptographicHash_hash")
	symaddr(&addr_QCryptographicHash_hashLength, "QCryptographicHash_hashLength")
	symaddr(&addr_QCryptographicHash_supportsAlgorithm, "QCryptographicHash_supportsAlgorithm")
	reg(&pg_QCryptographicHash_delete, "QCryptographicHash_delete")
}

type QCryptographicHash struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCryptographicHash) UnsafeGoID() int64     { return this.vid }
func (this *QCryptographicHash) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCryptographicHash(h uintptr) *QCryptographicHash {
	if h == 0 {
		return nil
	}
	return &QCryptographicHash{h: h}
}

func newownedQCryptographicHash(h uintptr) *QCryptographicHash {
	o := bindQCryptographicHash(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCryptographicHash).Delete)
	}
	return o
}

func (this *QCryptographicHash) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCryptographicHash) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCryptographicHash(h unsafe.Pointer) *QCryptographicHash { return bindQCryptographicHash(uintptr(h)) }
func UnsafeNewQCryptographicHashOwned(h unsafe.Pointer) *QCryptographicHash { return newownedQCryptographicHash(uintptr(h)) }

func (this *QCryptographicHash) UnsafeRelease() { this.h = 0 }

func NewQCryptographicHash(method QCryptographicHash__Algorithm) *QCryptographicHash {
	__r, _, _ := purego.SyscallN(addr_QCryptographicHash_new, uintptr(method))
	__o := newownedQCryptographicHash(__r)
	return __o
}

func (this *QCryptographicHash) Swap(other *QCryptographicHash) {
	purego.SyscallN(addr_QCryptographicHash_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCryptographicHash) Reset() {
	purego.SyscallN(addr_QCryptographicHash_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QCryptographicHash) Algorithm() QCryptographicHash__Algorithm {
	__r0, _, _ := purego.SyscallN(addr_QCryptographicHash_algorithm, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCryptographicHash__Algorithm(__r)
}

func (this *QCryptographicHash) AddData(data string, length int64) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QCryptographicHash_addData, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(length))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QCryptographicHash) AddDataWithData(data *QByteArrayView) {
	purego.SyscallN(addr_QCryptographicHash_addData2, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QCryptographicHash) AddDataWithDevice(device *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QCryptographicHash_addData3, this.h, uintptr(device.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	return __r != 0
}

func (this *QCryptographicHash) Result() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCryptographicHash_result, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCryptographicHash) ResultView() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QCryptographicHash_resultView, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func QCryptographicHash_Hash(data *QByteArrayView, method QCryptographicHash__Algorithm) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCryptographicHash_hash, uintptr(data.UnsafePointer()), uintptr(method))
	runtime.KeepAlive(data)
	return newownedQByteArray(__r)
}

func QCryptographicHash_HashLength(method QCryptographicHash__Algorithm) int {
	__r0, _, _ := purego.SyscallN(addr_QCryptographicHash_hashLength, uintptr(method))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCryptographicHash_SupportsAlgorithm(method QCryptographicHash__Algorithm) bool {
	__r0, _, _ := purego.SyscallN(addr_QCryptographicHash_supportsAlgorithm, uintptr(method))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QCryptographicHash) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCryptographicHash_delete(this.h)
	}
	this.UnsafeRelease()
}
