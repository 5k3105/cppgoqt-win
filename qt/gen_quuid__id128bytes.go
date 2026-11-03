package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUuid__Id128Bytes_new uintptr
	addr_QUuid__Id128Bytes_new2 uintptr
	addr_QUuid__Id128Bytes_ToQByteArrayView uintptr
	pg_QUuid__Id128Bytes_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUuid__Id128Bytes_new, "QUuid__Id128Bytes_new")
	symaddr(&addr_QUuid__Id128Bytes_new2, "QUuid__Id128Bytes_new2")
	symaddr(&addr_QUuid__Id128Bytes_ToQByteArrayView, "QUuid__Id128Bytes_ToQByteArrayView")
	reg(&pg_QUuid__Id128Bytes_delete, "QUuid__Id128Bytes_delete")
}

type QUuid__Id128Bytes struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUuid__Id128Bytes) UnsafeGoID() int64     { return this.vid }
func (this *QUuid__Id128Bytes) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUuid__Id128Bytes(h uintptr) *QUuid__Id128Bytes {
	if h == 0 {
		return nil
	}
	return &QUuid__Id128Bytes{h: h}
}

func newownedQUuid__Id128Bytes(h uintptr) *QUuid__Id128Bytes {
	o := bindQUuid__Id128Bytes(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUuid__Id128Bytes).Delete)
	}
	return o
}

func (this *QUuid__Id128Bytes) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUuid__Id128Bytes) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUuid__Id128Bytes(h unsafe.Pointer) *QUuid__Id128Bytes { return bindQUuid__Id128Bytes(uintptr(h)) }
func UnsafeNewQUuid__Id128BytesOwned(h unsafe.Pointer) *QUuid__Id128Bytes { return newownedQUuid__Id128Bytes(uintptr(h)) }

func (this *QUuid__Id128Bytes) UnsafeRelease() { this.h = 0 }

func NewQUuid__Id128Bytes() *QUuid__Id128Bytes {
	__r, _, _ := purego.SyscallN(addr_QUuid__Id128Bytes_new)
	__o := newownedQUuid__Id128Bytes(__r)
	return __o
}

func NewQUuid__Id128Bytes2(arg1 *QUuid__Id128Bytes) *QUuid__Id128Bytes {
	__r, _, _ := purego.SyscallN(addr_QUuid__Id128Bytes_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQUuid__Id128Bytes(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QUuid__Id128Bytes) ToQByteArrayView() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QUuid__Id128Bytes_ToQByteArrayView, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QUuid__Id128Bytes) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUuid__Id128Bytes_delete(this.h)
	}
	this.UnsafeRelease()
}
