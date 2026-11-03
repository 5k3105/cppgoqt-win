package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSysInfo_buildCpuArchitecture uintptr
	addr_QSysInfo_currentCpuArchitecture uintptr
	addr_QSysInfo_buildAbi uintptr
	addr_QSysInfo_kernelType uintptr
	addr_QSysInfo_kernelVersion uintptr
	addr_QSysInfo_productType uintptr
	addr_QSysInfo_productVersion uintptr
	addr_QSysInfo_prettyProductName uintptr
	addr_QSysInfo_machineHostName uintptr
	addr_QSysInfo_machineUniqueId uintptr
	addr_QSysInfo_bootUniqueId uintptr
	pg_QSysInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSysInfo_buildCpuArchitecture, "QSysInfo_buildCpuArchitecture")
	symaddr(&addr_QSysInfo_currentCpuArchitecture, "QSysInfo_currentCpuArchitecture")
	symaddr(&addr_QSysInfo_buildAbi, "QSysInfo_buildAbi")
	symaddr(&addr_QSysInfo_kernelType, "QSysInfo_kernelType")
	symaddr(&addr_QSysInfo_kernelVersion, "QSysInfo_kernelVersion")
	symaddr(&addr_QSysInfo_productType, "QSysInfo_productType")
	symaddr(&addr_QSysInfo_productVersion, "QSysInfo_productVersion")
	symaddr(&addr_QSysInfo_prettyProductName, "QSysInfo_prettyProductName")
	symaddr(&addr_QSysInfo_machineHostName, "QSysInfo_machineHostName")
	symaddr(&addr_QSysInfo_machineUniqueId, "QSysInfo_machineUniqueId")
	symaddr(&addr_QSysInfo_bootUniqueId, "QSysInfo_bootUniqueId")
	reg(&pg_QSysInfo_delete, "QSysInfo_delete")
}

type QSysInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSysInfo) UnsafeGoID() int64     { return this.vid }
func (this *QSysInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSysInfo(h uintptr) *QSysInfo {
	if h == 0 {
		return nil
	}
	return &QSysInfo{h: h}
}

func newownedQSysInfo(h uintptr) *QSysInfo {
	o := bindQSysInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSysInfo).Delete)
	}
	return o
}

func (this *QSysInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSysInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSysInfo(h unsafe.Pointer) *QSysInfo { return bindQSysInfo(uintptr(h)) }
func UnsafeNewQSysInfoOwned(h unsafe.Pointer) *QSysInfo { return newownedQSysInfo(uintptr(h)) }

func (this *QSysInfo) UnsafeRelease() { this.h = 0 }

func QSysInfo_BuildCpuArchitecture() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_buildCpuArchitecture, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_CurrentCpuArchitecture() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_currentCpuArchitecture, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_BuildAbi() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_buildAbi, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_KernelType() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_kernelType, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_KernelVersion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_kernelVersion, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_ProductType() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_productType, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_ProductVersion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_productVersion, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_PrettyProductName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_prettyProductName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_MachineHostName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSysInfo_machineHostName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QSysInfo_MachineUniqueId() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QSysInfo_machineUniqueId)
	return newownedQByteArray(__r)
}

func QSysInfo_BootUniqueId() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QSysInfo_bootUniqueId)
	return newownedQByteArray(__r)
}

func (this *QSysInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSysInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
