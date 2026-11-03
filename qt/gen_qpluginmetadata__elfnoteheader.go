package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPluginMetaData__ElfNoteHeader_new uintptr
	addr_QPluginMetaData__ElfNoteHeader_new2 uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_namesz_get uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_namesz_set uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_descsz_get uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_descsz_set uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_type_get uintptr
	addr_QPluginMetaData__ElfNoteHeader_n_type_set uintptr
	addr_QPluginMetaData__ElfNoteHeader_header_get uintptr
	addr_QPluginMetaData__ElfNoteHeader_header_set uintptr
	pg_QPluginMetaData__ElfNoteHeader_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_new, "QPluginMetaData__ElfNoteHeader_new")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_new2, "QPluginMetaData__ElfNoteHeader_new2")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_namesz_get, "QPluginMetaData__ElfNoteHeader_n_namesz_get")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_namesz_set, "QPluginMetaData__ElfNoteHeader_n_namesz_set")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_descsz_get, "QPluginMetaData__ElfNoteHeader_n_descsz_get")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_descsz_set, "QPluginMetaData__ElfNoteHeader_n_descsz_set")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_type_get, "QPluginMetaData__ElfNoteHeader_n_type_get")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_n_type_set, "QPluginMetaData__ElfNoteHeader_n_type_set")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_header_get, "QPluginMetaData__ElfNoteHeader_header_get")
	symaddr(&addr_QPluginMetaData__ElfNoteHeader_header_set, "QPluginMetaData__ElfNoteHeader_header_set")
	reg(&pg_QPluginMetaData__ElfNoteHeader_delete, "QPluginMetaData__ElfNoteHeader_delete")
}

type QPluginMetaData__ElfNoteHeader struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPluginMetaData__ElfNoteHeader) UnsafeGoID() int64     { return this.vid }
func (this *QPluginMetaData__ElfNoteHeader) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPluginMetaData__ElfNoteHeader(h uintptr) *QPluginMetaData__ElfNoteHeader {
	if h == 0 {
		return nil
	}
	return &QPluginMetaData__ElfNoteHeader{h: h}
}

func newownedQPluginMetaData__ElfNoteHeader(h uintptr) *QPluginMetaData__ElfNoteHeader {
	o := bindQPluginMetaData__ElfNoteHeader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPluginMetaData__ElfNoteHeader).Delete)
	}
	return o
}

func (this *QPluginMetaData__ElfNoteHeader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPluginMetaData__ElfNoteHeader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPluginMetaData__ElfNoteHeader(h unsafe.Pointer) *QPluginMetaData__ElfNoteHeader { return bindQPluginMetaData__ElfNoteHeader(uintptr(h)) }
func UnsafeNewQPluginMetaData__ElfNoteHeaderOwned(h unsafe.Pointer) *QPluginMetaData__ElfNoteHeader { return newownedQPluginMetaData__ElfNoteHeader(uintptr(h)) }

func (this *QPluginMetaData__ElfNoteHeader) UnsafeRelease() { this.h = 0 }

func NewQPluginMetaData__ElfNoteHeader(payloadSize uint) *QPluginMetaData__ElfNoteHeader {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_new, uintptr(payloadSize))
	__o := newownedQPluginMetaData__ElfNoteHeader(__r)
	return __o
}

func NewQPluginMetaData__ElfNoteHeader2(arg1 *QPluginMetaData__ElfNoteHeader) *QPluginMetaData__ElfNoteHeader {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPluginMetaData__ElfNoteHeader(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPluginMetaData__ElfNoteHeader) N_namesz() uint {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_namesz_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPluginMetaData__ElfNoteHeader) SetN_namesz(n_namesz uint) {
	purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_namesz_set, this.h, uintptr(n_namesz))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__ElfNoteHeader) N_descsz() uint {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_descsz_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPluginMetaData__ElfNoteHeader) SetN_descsz(n_descsz uint) {
	purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_descsz_set, this.h, uintptr(n_descsz))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__ElfNoteHeader) N_type() uint {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPluginMetaData__ElfNoteHeader) SetN_type(n_type uint) {
	purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_n_type_set, this.h, uintptr(n_type))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__ElfNoteHeader) Header() *QPluginMetaData__Header {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_header_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPluginMetaData__Header(__r)
}

func (this *QPluginMetaData__ElfNoteHeader) SetHeader(header *QPluginMetaData__Header) {
	purego.SyscallN(addr_QPluginMetaData__ElfNoteHeader_header_set, this.h, uintptr(header.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(header)
}

func (this *QPluginMetaData__ElfNoteHeader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPluginMetaData__ElfNoteHeader_delete(this.h)
	}
	this.UnsafeRelease()
}
