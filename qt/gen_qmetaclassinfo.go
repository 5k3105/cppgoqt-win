package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaClassInfo_new uintptr
	addr_QMetaClassInfo_new2 uintptr
	addr_QMetaClassInfo_name uintptr
	addr_QMetaClassInfo_value uintptr
	addr_QMetaClassInfo_enclosingMetaObject uintptr
	pg_QMetaClassInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaClassInfo_new, "QMetaClassInfo_new")
	symaddr(&addr_QMetaClassInfo_new2, "QMetaClassInfo_new2")
	symaddr(&addr_QMetaClassInfo_name, "QMetaClassInfo_name")
	symaddr(&addr_QMetaClassInfo_value, "QMetaClassInfo_value")
	symaddr(&addr_QMetaClassInfo_enclosingMetaObject, "QMetaClassInfo_enclosingMetaObject")
	reg(&pg_QMetaClassInfo_delete, "QMetaClassInfo_delete")
}

type QMetaClassInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaClassInfo) UnsafeGoID() int64     { return this.vid }
func (this *QMetaClassInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaClassInfo(h uintptr) *QMetaClassInfo {
	if h == 0 {
		return nil
	}
	return &QMetaClassInfo{h: h}
}

func newownedQMetaClassInfo(h uintptr) *QMetaClassInfo {
	o := bindQMetaClassInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaClassInfo).Delete)
	}
	return o
}

func (this *QMetaClassInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaClassInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaClassInfo(h unsafe.Pointer) *QMetaClassInfo { return bindQMetaClassInfo(uintptr(h)) }
func UnsafeNewQMetaClassInfoOwned(h unsafe.Pointer) *QMetaClassInfo { return newownedQMetaClassInfo(uintptr(h)) }

func (this *QMetaClassInfo) UnsafeRelease() { this.h = 0 }

func NewQMetaClassInfo() *QMetaClassInfo {
	__r, _, _ := purego.SyscallN(addr_QMetaClassInfo_new)
	__o := newownedQMetaClassInfo(__r)
	return __o
}

func NewQMetaClassInfo2(arg1 *QMetaClassInfo) *QMetaClassInfo {
	__r, _, _ := purego.SyscallN(addr_QMetaClassInfo_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaClassInfo(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaClassInfo) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaClassInfo_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaClassInfo) Value() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaClassInfo_value, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaClassInfo) EnclosingMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaClassInfo_enclosingMetaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaClassInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaClassInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
