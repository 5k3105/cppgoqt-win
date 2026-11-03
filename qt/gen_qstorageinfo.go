package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStorageInfo_new uintptr
	addr_QStorageInfo_new2 uintptr
	addr_QStorageInfo_new3 uintptr
	addr_QStorageInfo_new4 uintptr
	addr_QStorageInfo_OperatorAssign uintptr
	addr_QStorageInfo_swap uintptr
	addr_QStorageInfo_setPath uintptr
	addr_QStorageInfo_rootPath uintptr
	addr_QStorageInfo_device uintptr
	addr_QStorageInfo_subvolume uintptr
	addr_QStorageInfo_fileSystemType uintptr
	addr_QStorageInfo_name uintptr
	addr_QStorageInfo_displayName uintptr
	addr_QStorageInfo_bytesTotal uintptr
	addr_QStorageInfo_bytesFree uintptr
	addr_QStorageInfo_bytesAvailable uintptr
	addr_QStorageInfo_blockSize uintptr
	addr_QStorageInfo_isRoot uintptr
	addr_QStorageInfo_isReadOnly uintptr
	addr_QStorageInfo_isReady uintptr
	addr_QStorageInfo_isValid uintptr
	addr_QStorageInfo_refresh uintptr
	addr_QStorageInfo_mountedVolumes uintptr
	addr_QStorageInfo_root uintptr
	addr_QStorageInfo_OperatorEqual uintptr
	addr_QStorageInfo_OperatorNotEqual uintptr
	pg_QStorageInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStorageInfo_new, "QStorageInfo_new")
	symaddr(&addr_QStorageInfo_new2, "QStorageInfo_new2")
	symaddr(&addr_QStorageInfo_new3, "QStorageInfo_new3")
	symaddr(&addr_QStorageInfo_new4, "QStorageInfo_new4")
	symaddr(&addr_QStorageInfo_OperatorAssign, "QStorageInfo_OperatorAssign")
	symaddr(&addr_QStorageInfo_swap, "QStorageInfo_swap")
	symaddr(&addr_QStorageInfo_setPath, "QStorageInfo_setPath")
	symaddr(&addr_QStorageInfo_rootPath, "QStorageInfo_rootPath")
	symaddr(&addr_QStorageInfo_device, "QStorageInfo_device")
	symaddr(&addr_QStorageInfo_subvolume, "QStorageInfo_subvolume")
	symaddr(&addr_QStorageInfo_fileSystemType, "QStorageInfo_fileSystemType")
	symaddr(&addr_QStorageInfo_name, "QStorageInfo_name")
	symaddr(&addr_QStorageInfo_displayName, "QStorageInfo_displayName")
	symaddr(&addr_QStorageInfo_bytesTotal, "QStorageInfo_bytesTotal")
	symaddr(&addr_QStorageInfo_bytesFree, "QStorageInfo_bytesFree")
	symaddr(&addr_QStorageInfo_bytesAvailable, "QStorageInfo_bytesAvailable")
	symaddr(&addr_QStorageInfo_blockSize, "QStorageInfo_blockSize")
	symaddr(&addr_QStorageInfo_isRoot, "QStorageInfo_isRoot")
	symaddr(&addr_QStorageInfo_isReadOnly, "QStorageInfo_isReadOnly")
	symaddr(&addr_QStorageInfo_isReady, "QStorageInfo_isReady")
	symaddr(&addr_QStorageInfo_isValid, "QStorageInfo_isValid")
	symaddr(&addr_QStorageInfo_refresh, "QStorageInfo_refresh")
	symaddr(&addr_QStorageInfo_mountedVolumes, "QStorageInfo_mountedVolumes")
	symaddr(&addr_QStorageInfo_root, "QStorageInfo_root")
	symaddr(&addr_QStorageInfo_OperatorEqual, "QStorageInfo_OperatorEqual")
	symaddr(&addr_QStorageInfo_OperatorNotEqual, "QStorageInfo_OperatorNotEqual")
	reg(&pg_QStorageInfo_delete, "QStorageInfo_delete")
}

type QStorageInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStorageInfo) UnsafeGoID() int64     { return this.vid }
func (this *QStorageInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStorageInfo(h uintptr) *QStorageInfo {
	if h == 0 {
		return nil
	}
	return &QStorageInfo{h: h}
}

func newownedQStorageInfo(h uintptr) *QStorageInfo {
	o := bindQStorageInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStorageInfo).Delete)
	}
	return o
}

func (this *QStorageInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStorageInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStorageInfo(h unsafe.Pointer) *QStorageInfo { return bindQStorageInfo(uintptr(h)) }
func UnsafeNewQStorageInfoOwned(h unsafe.Pointer) *QStorageInfo { return newownedQStorageInfo(uintptr(h)) }

func (this *QStorageInfo) UnsafeRelease() { this.h = 0 }

func NewQStorageInfo() *QStorageInfo {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_new)
	__o := newownedQStorageInfo(__r)
	return __o
}

func NewQStorageInfo2(path string) *QStorageInfo {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_new2, path_p, uintptr(len(path)))
	__o := newownedQStorageInfo(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQStorageInfo3(dir *QDir) *QStorageInfo {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_new3, uintptr(dir.UnsafePointer()))
	__o := newownedQStorageInfo(__r)
	runtime.KeepAlive(dir)
	return __o
}

func NewQStorageInfo4(other *QStorageInfo) *QStorageInfo {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_new4, uintptr(other.UnsafePointer()))
	__o := newownedQStorageInfo(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStorageInfo) OperatorAssign(other *QStorageInfo) *QStorageInfo {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQStorageInfo(__r)
}

func (this *QStorageInfo) Swap(other *QStorageInfo) {
	purego.SyscallN(addr_QStorageInfo_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QStorageInfo) SetPath(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QStorageInfo_setPath, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QStorageInfo) RootPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_rootPath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStorageInfo) Device() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_device, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStorageInfo) Subvolume() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_subvolume, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStorageInfo) FileSystemType() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_fileSystemType, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QStorageInfo) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStorageInfo) DisplayName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_displayName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStorageInfo) BytesTotal() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_bytesTotal, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStorageInfo) BytesFree() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_bytesFree, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStorageInfo) BytesAvailable() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_bytesAvailable, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStorageInfo) BlockSize() int {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_blockSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStorageInfo) IsRoot() bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_isRoot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStorageInfo) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStorageInfo) IsReady() bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_isReady, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStorageInfo) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStorageInfo) Refresh() {
	purego.SyscallN(addr_QStorageInfo_refresh, this.h)
	runtime.KeepAlive(this)
}

func QStorageInfo_MountedVolumes() []*QStorageInfo {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_mountedVolumes, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QStorageInfo, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQStorageInfo(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStorageInfo_Root() *QStorageInfo {
	__r, _, _ := purego.SyscallN(addr_QStorageInfo_root)
	return newownedQStorageInfo(__r)
}

func (this *QStorageInfo) OperatorEqual(rhs *QStorageInfo) bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStorageInfo) OperatorNotEqual(rhs *QStorageInfo) bool {
	__r0, _, _ := purego.SyscallN(addr_QStorageInfo_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QStorageInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStorageInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
