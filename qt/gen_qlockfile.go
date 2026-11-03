package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLockFile_new uintptr
	addr_QLockFile_fileName uintptr
	addr_QLockFile_lock uintptr
	addr_QLockFile_tryLock uintptr
	addr_QLockFile_unlock uintptr
	addr_QLockFile_setStaleLockTime uintptr
	addr_QLockFile_staleLockTime uintptr
	addr_QLockFile_tryLock2 uintptr
	addr_QLockFile_setStaleLockTime2 uintptr
	addr_QLockFile_staleLockTimeAsDuration uintptr
	addr_QLockFile_isLocked uintptr
	addr_QLockFile_getLockInfo uintptr
	addr_QLockFile_removeStaleLockFile uintptr
	addr_QLockFile_error uintptr
	addr_QLockFile_tryLock3 uintptr
	pg_QLockFile_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLockFile_new, "QLockFile_new")
	symaddr(&addr_QLockFile_fileName, "QLockFile_fileName")
	symaddr(&addr_QLockFile_lock, "QLockFile_lock")
	symaddr(&addr_QLockFile_tryLock, "QLockFile_tryLock")
	symaddr(&addr_QLockFile_unlock, "QLockFile_unlock")
	symaddr(&addr_QLockFile_setStaleLockTime, "QLockFile_setStaleLockTime")
	symaddr(&addr_QLockFile_staleLockTime, "QLockFile_staleLockTime")
	symaddr(&addr_QLockFile_tryLock2, "QLockFile_tryLock2")
	symaddr(&addr_QLockFile_setStaleLockTime2, "QLockFile_setStaleLockTime2")
	symaddr(&addr_QLockFile_staleLockTimeAsDuration, "QLockFile_staleLockTimeAsDuration")
	symaddr(&addr_QLockFile_isLocked, "QLockFile_isLocked")
	symaddr(&addr_QLockFile_getLockInfo, "QLockFile_getLockInfo")
	symaddr(&addr_QLockFile_removeStaleLockFile, "QLockFile_removeStaleLockFile")
	symaddr(&addr_QLockFile_error, "QLockFile_error")
	symaddr(&addr_QLockFile_tryLock3, "QLockFile_tryLock3")
	reg(&pg_QLockFile_delete, "QLockFile_delete")
}

type QLockFile struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLockFile) UnsafeGoID() int64     { return this.vid }
func (this *QLockFile) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLockFile(h uintptr) *QLockFile {
	if h == 0 {
		return nil
	}
	return &QLockFile{h: h}
}

func newownedQLockFile(h uintptr) *QLockFile {
	o := bindQLockFile(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLockFile).Delete)
	}
	return o
}

func (this *QLockFile) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLockFile) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLockFile(h unsafe.Pointer) *QLockFile { return bindQLockFile(uintptr(h)) }
func UnsafeNewQLockFileOwned(h unsafe.Pointer) *QLockFile { return newownedQLockFile(uintptr(h)) }

func (this *QLockFile) UnsafeRelease() { this.h = 0 }

func NewQLockFile(fileName string) *QLockFile {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QLockFile_new, fileName_p, uintptr(len(fileName)))
	__o := newownedQLockFile(__r)
	runtime.KeepAlive(fileName)
	return __o
}

func (this *QLockFile) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLockFile_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLockFile) Lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) TryLock(timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_tryLock, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) Unlock() {
	purego.SyscallN(addr_QLockFile_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QLockFile) SetStaleLockTime(arg1 int) {
	purego.SyscallN(addr_QLockFile_setStaleLockTime, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLockFile) StaleLockTime() int {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_staleLockTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLockFile) TryLock2() bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_tryLock2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) SetStaleLockTimeWithValue(value int64) {
	purego.SyscallN(addr_QLockFile_setStaleLockTime2, this.h, uintptr(value))
	runtime.KeepAlive(this)
}

func (this *QLockFile) StaleLockTimeAsDuration() int64 {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_staleLockTimeAsDuration, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QLockFile) IsLocked() bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_isLocked, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) GetLockInfo(hostname *QString, appname *QString) (bool, int64) {
	var pid_out int64
	__r0, _, _ := purego.SyscallN(addr_QLockFile_getLockInfo, this.h, uintptr(unsafe.Pointer(&pid_out)), uintptr(hostname.UnsafePointer()), uintptr(appname.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(hostname)
	runtime.KeepAlive(appname)
	return __r != 0, pid_out
}

func (this *QLockFile) RemoveStaleLockFile() bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_removeStaleLockFile, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) Error() QLockFile__LockError {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLockFile__LockError(__r)
}

func (this *QLockFile) TryLockWithTimeout(timeout int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QLockFile_tryLock3, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLockFile) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLockFile_delete(this.h)
	}
	this.UnsafeRelease()
}
