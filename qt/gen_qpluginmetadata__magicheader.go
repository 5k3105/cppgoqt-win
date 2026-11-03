package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPluginMetaData__MagicHeader_new uintptr
	addr_QPluginMetaData__MagicHeader_header_get uintptr
	addr_QPluginMetaData__MagicHeader_header_set uintptr
	pg_QPluginMetaData__MagicHeader_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPluginMetaData__MagicHeader_new, "QPluginMetaData__MagicHeader_new")
	symaddr(&addr_QPluginMetaData__MagicHeader_header_get, "QPluginMetaData__MagicHeader_header_get")
	symaddr(&addr_QPluginMetaData__MagicHeader_header_set, "QPluginMetaData__MagicHeader_header_set")
	reg(&pg_QPluginMetaData__MagicHeader_delete, "QPluginMetaData__MagicHeader_delete")
}

type QPluginMetaData__MagicHeader struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPluginMetaData__MagicHeader) UnsafeGoID() int64     { return this.vid }
func (this *QPluginMetaData__MagicHeader) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPluginMetaData__MagicHeader(h uintptr) *QPluginMetaData__MagicHeader {
	if h == 0 {
		return nil
	}
	return &QPluginMetaData__MagicHeader{h: h}
}

func newownedQPluginMetaData__MagicHeader(h uintptr) *QPluginMetaData__MagicHeader {
	o := bindQPluginMetaData__MagicHeader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPluginMetaData__MagicHeader).Delete)
	}
	return o
}

func (this *QPluginMetaData__MagicHeader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPluginMetaData__MagicHeader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPluginMetaData__MagicHeader(h unsafe.Pointer) *QPluginMetaData__MagicHeader { return bindQPluginMetaData__MagicHeader(uintptr(h)) }
func UnsafeNewQPluginMetaData__MagicHeaderOwned(h unsafe.Pointer) *QPluginMetaData__MagicHeader { return newownedQPluginMetaData__MagicHeader(uintptr(h)) }

func (this *QPluginMetaData__MagicHeader) UnsafeRelease() { this.h = 0 }

func NewQPluginMetaData__MagicHeader() *QPluginMetaData__MagicHeader {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__MagicHeader_new)
	__o := newownedQPluginMetaData__MagicHeader(__r)
	return __o
}

func (this *QPluginMetaData__MagicHeader) Header() *QPluginMetaData__Header {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__MagicHeader_header_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPluginMetaData__Header(__r)
}

func (this *QPluginMetaData__MagicHeader) SetHeader(header *QPluginMetaData__Header) {
	purego.SyscallN(addr_QPluginMetaData__MagicHeader_header_set, this.h, uintptr(header.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(header)
}

func (this *QPluginMetaData__MagicHeader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPluginMetaData__MagicHeader_delete(this.h)
	}
	this.UnsafeRelease()
}
