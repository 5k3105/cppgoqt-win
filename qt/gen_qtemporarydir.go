package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTemporaryDir_new uintptr
	addr_QTemporaryDir_new2 uintptr
	addr_QTemporaryDir_swap uintptr
	addr_QTemporaryDir_isValid uintptr
	addr_QTemporaryDir_errorString uintptr
	addr_QTemporaryDir_autoRemove uintptr
	addr_QTemporaryDir_setAutoRemove uintptr
	addr_QTemporaryDir_remove uintptr
	addr_QTemporaryDir_path uintptr
	addr_QTemporaryDir_filePath uintptr
	pg_QTemporaryDir_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTemporaryDir_new, "QTemporaryDir_new")
	symaddr(&addr_QTemporaryDir_new2, "QTemporaryDir_new2")
	symaddr(&addr_QTemporaryDir_swap, "QTemporaryDir_swap")
	symaddr(&addr_QTemporaryDir_isValid, "QTemporaryDir_isValid")
	symaddr(&addr_QTemporaryDir_errorString, "QTemporaryDir_errorString")
	symaddr(&addr_QTemporaryDir_autoRemove, "QTemporaryDir_autoRemove")
	symaddr(&addr_QTemporaryDir_setAutoRemove, "QTemporaryDir_setAutoRemove")
	symaddr(&addr_QTemporaryDir_remove, "QTemporaryDir_remove")
	symaddr(&addr_QTemporaryDir_path, "QTemporaryDir_path")
	symaddr(&addr_QTemporaryDir_filePath, "QTemporaryDir_filePath")
	reg(&pg_QTemporaryDir_delete, "QTemporaryDir_delete")
}

type QTemporaryDir struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTemporaryDir) UnsafeGoID() int64     { return this.vid }
func (this *QTemporaryDir) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTemporaryDir(h uintptr) *QTemporaryDir {
	if h == 0 {
		return nil
	}
	return &QTemporaryDir{h: h}
}

func newownedQTemporaryDir(h uintptr) *QTemporaryDir {
	o := bindQTemporaryDir(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTemporaryDir).Delete)
	}
	return o
}

func (this *QTemporaryDir) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTemporaryDir) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTemporaryDir(h unsafe.Pointer) *QTemporaryDir { return bindQTemporaryDir(uintptr(h)) }
func UnsafeNewQTemporaryDirOwned(h unsafe.Pointer) *QTemporaryDir { return newownedQTemporaryDir(uintptr(h)) }

func (this *QTemporaryDir) UnsafeRelease() { this.h = 0 }

func NewQTemporaryDir() *QTemporaryDir {
	__r, _, _ := purego.SyscallN(addr_QTemporaryDir_new)
	__o := newownedQTemporaryDir(__r)
	return __o
}

func NewQTemporaryDir2(templateName string) *QTemporaryDir {
	var templateName_p uintptr
	if len(templateName) > 0 {
		templateName_p = uintptr(unsafe.Pointer(unsafe.StringData(templateName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTemporaryDir_new2, templateName_p, uintptr(len(templateName)))
	__o := newownedQTemporaryDir(__r)
	runtime.KeepAlive(templateName)
	return __o
}

func (this *QTemporaryDir) Swap(other *QTemporaryDir) {
	purego.SyscallN(addr_QTemporaryDir_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTemporaryDir) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryDir_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryDir) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryDir_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTemporaryDir) AutoRemove() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryDir_autoRemove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryDir) SetAutoRemove(b bool) {
	purego.SyscallN(addr_QTemporaryDir_setAutoRemove, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTemporaryDir) Remove() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryDir_remove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryDir) Path() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryDir_path, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTemporaryDir) FilePath(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryDir_filePath, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func (this *QTemporaryDir) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTemporaryDir_delete(this.h)
	}
	this.UnsafeRelease()
}
