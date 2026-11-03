package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QModelRoleData_new uintptr
	addr_QModelRoleData_new2 uintptr
	addr_QModelRoleData_role uintptr
	addr_QModelRoleData_data uintptr
	addr_QModelRoleData_data2 uintptr
	addr_QModelRoleData_clearData uintptr
	addr_QModelRoleData_OperatorAssign uintptr
	pg_QModelRoleData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QModelRoleData_new, "QModelRoleData_new")
	symaddr(&addr_QModelRoleData_new2, "QModelRoleData_new2")
	symaddr(&addr_QModelRoleData_role, "QModelRoleData_role")
	symaddr(&addr_QModelRoleData_data, "QModelRoleData_data")
	symaddr(&addr_QModelRoleData_data2, "QModelRoleData_data2")
	symaddr(&addr_QModelRoleData_clearData, "QModelRoleData_clearData")
	symaddr(&addr_QModelRoleData_OperatorAssign, "QModelRoleData_OperatorAssign")
	reg(&pg_QModelRoleData_delete, "QModelRoleData_delete")
}

type QModelRoleData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QModelRoleData) UnsafeGoID() int64     { return this.vid }
func (this *QModelRoleData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQModelRoleData(h uintptr) *QModelRoleData {
	if h == 0 {
		return nil
	}
	return &QModelRoleData{h: h}
}

func newownedQModelRoleData(h uintptr) *QModelRoleData {
	o := bindQModelRoleData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QModelRoleData).Delete)
	}
	return o
}

func (this *QModelRoleData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QModelRoleData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQModelRoleData(h unsafe.Pointer) *QModelRoleData { return bindQModelRoleData(uintptr(h)) }
func UnsafeNewQModelRoleDataOwned(h unsafe.Pointer) *QModelRoleData { return newownedQModelRoleData(uintptr(h)) }

func (this *QModelRoleData) UnsafeRelease() { this.h = 0 }

func NewQModelRoleData(role int) *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleData_new, uintptr(role))
	__o := newownedQModelRoleData(__r)
	return __o
}

func NewQModelRoleData2(arg1 *QModelRoleData) *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleData_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQModelRoleData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QModelRoleData) Role() int {
	__r0, _, _ := purego.SyscallN(addr_QModelRoleData_role, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QModelRoleData) Data() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QModelRoleData_data, this.h)
	runtime.KeepAlive(this)
	return bindQVariant(__r)
}

func (this *QModelRoleData) Data2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QModelRoleData_data2, this.h)
	runtime.KeepAlive(this)
	return bindQVariant(__r)
}

func (this *QModelRoleData) ClearData() {
	purego.SyscallN(addr_QModelRoleData_clearData, this.h)
	runtime.KeepAlive(this)
}

func (this *QModelRoleData) OperatorAssign(arg1 *QModelRoleData) *QModelRoleData {
	__r, _, _ := purego.SyscallN(addr_QModelRoleData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQModelRoleData(__r)
}

func (this *QModelRoleData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QModelRoleData_delete(this.h)
	}
	this.UnsafeRelease()
}
