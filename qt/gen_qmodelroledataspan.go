package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QModelRoleDataSpan_new uintptr
	addr_QModelRoleDataSpan_new2 uintptr
	addr_QModelRoleDataSpan_new3 uintptr
	addr_QModelRoleDataSpan_new4 uintptr
	addr_QModelRoleDataSpan_size uintptr
	addr_QModelRoleDataSpan_length uintptr
	addr_QModelRoleDataSpan_data uintptr
	addr_QModelRoleDataSpan_begin uintptr
	addr_QModelRoleDataSpan_end uintptr
	addr_QModelRoleDataSpan_OperatorSubscript uintptr
	addr_QModelRoleDataSpan_dataForRole uintptr
	pg_QModelRoleDataSpan_delete func(uintptr)
)

func init() {
	symaddr(&addr_QModelRoleDataSpan_new, "QModelRoleDataSpan_new")
	symaddr(&addr_QModelRoleDataSpan_new2, "QModelRoleDataSpan_new2")
	symaddr(&addr_QModelRoleDataSpan_new3, "QModelRoleDataSpan_new3")
	symaddr(&addr_QModelRoleDataSpan_new4, "QModelRoleDataSpan_new4")
	symaddr(&addr_QModelRoleDataSpan_size, "QModelRoleDataSpan_size")
	symaddr(&addr_QModelRoleDataSpan_length, "QModelRoleDataSpan_length")
	symaddr(&addr_QModelRoleDataSpan_data, "QModelRoleDataSpan_data")
	symaddr(&addr_QModelRoleDataSpan_begin, "QModelRoleDataSpan_begin")
	symaddr(&addr_QModelRoleDataSpan_end, "QModelRoleDataSpan_end")
	symaddr(&addr_QModelRoleDataSpan_OperatorSubscript, "QModelRoleDataSpan_OperatorSubscript")
	symaddr(&addr_QModelRoleDataSpan_dataForRole, "QModelRoleDataSpan_dataForRole")
	reg(&pg_QModelRoleDataSpan_delete, "QModelRoleDataSpan_delete")
}

type QModelRoleDataSpan struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QModelRoleDataSpan) UnsafeGoID() int64     { return this.vid }
func (this *QModelRoleDataSpan) UnsafeSetGoID(id int64) { this.vid = id }

func bindQModelRoleDataSpan(h uintptr) *QModelRoleDataSpan {
	if h == 0 {
		return nil
	}
	return &QModelRoleDataSpan{h: h}
}

func newownedQModelRoleDataSpan(h uintptr) *QModelRoleDataSpan {
	o := bindQModelRoleDataSpan(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QModelRoleDataSpan).Delete)
	}
	return o
}

func (this *QModelRoleDataSpan) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QModelRoleDataSpan) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQModelRoleDataSpan(h unsafe.Pointer) *QModelRoleDataSpan { return bindQModelRoleDataSpan(uintptr(h)) }
func UnsafeNewQModelRoleDataSpanOwned(h unsafe.Pointer) *QModelRoleDataSpan { return newownedQModelRoleDataSpan(uintptr(h)) }

func (this *QModelRoleDataSpan) UnsafeRelease() { this.h = 0 }

func NewQModelRoleDataSpan() *QModelRoleDataSpan {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_new)
	__o := newownedQModelRoleDataSpan(__r)
	return __o
}

func NewQModelRoleDataSpan2(modelRoleData *QModelRoleData) *QModelRoleDataSpan {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_new2, uintptr(modelRoleData.UnsafePointer()))
	__o := newownedQModelRoleDataSpan(__r)
	runtime.KeepAlive(modelRoleData)
	return __o
}

func NewQModelRoleDataSpan3(modelRoleData *QModelRoleData, len_ int64) *QModelRoleDataSpan {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_new3, uintptr(modelRoleData.UnsafePointer()), uintptr(len_))
	__o := newownedQModelRoleDataSpan(__r)
	runtime.KeepAlive(modelRoleData)
	return __o
}

func NewQModelRoleDataSpan4(arg1 *QModelRoleDataSpan) *QModelRoleDataSpan {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQModelRoleDataSpan(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QModelRoleDataSpan) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QModelRoleDataSpan) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QModelRoleDataSpan) Data() *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_data, this.h)
	runtime.KeepAlive(this)
	return bindQModelRoleData(__r)
}

func (this *QModelRoleDataSpan) Begin() *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_begin, this.h)
	runtime.KeepAlive(this)
	return bindQModelRoleData(__r)
}

func (this *QModelRoleDataSpan) End() *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_end, this.h)
	runtime.KeepAlive(this)
	return bindQModelRoleData(__r)
}

func (this *QModelRoleDataSpan) OperatorSubscript(index int64) *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_OperatorSubscript, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQModelRoleData(__r)
}

func (this *QModelRoleDataSpan) DataForRole(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QModelRoleDataSpan_dataForRole, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return bindQVariant(__r)
}

func (this *QModelRoleDataSpan) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QModelRoleDataSpan_delete(this.h)
	}
	this.UnsafeRelease()
}
