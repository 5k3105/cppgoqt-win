package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProcessEnvironment_new uintptr
	addr_QProcessEnvironment_new2 uintptr
	addr_QProcessEnvironment_new3 uintptr
	addr_QProcessEnvironment_OperatorAssign uintptr
	addr_QProcessEnvironment_swap uintptr
	addr_QProcessEnvironment_isEmpty uintptr
	addr_QProcessEnvironment_inheritsFromParent uintptr
	addr_QProcessEnvironment_clear uintptr
	addr_QProcessEnvironment_contains uintptr
	addr_QProcessEnvironment_insert uintptr
	addr_QProcessEnvironment_remove uintptr
	addr_QProcessEnvironment_value uintptr
	addr_QProcessEnvironment_toStringList uintptr
	addr_QProcessEnvironment_keys uintptr
	addr_QProcessEnvironment_insert2 uintptr
	addr_QProcessEnvironment_systemEnvironment uintptr
	addr_QProcessEnvironment_value2 uintptr
	addr_QProcessEnvironment_OperatorEqual uintptr
	addr_QProcessEnvironment_OperatorNotEqual uintptr
	pg_QProcessEnvironment_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProcessEnvironment_new, "QProcessEnvironment_new")
	symaddr(&addr_QProcessEnvironment_new2, "QProcessEnvironment_new2")
	symaddr(&addr_QProcessEnvironment_new3, "QProcessEnvironment_new3")
	symaddr(&addr_QProcessEnvironment_OperatorAssign, "QProcessEnvironment_OperatorAssign")
	symaddr(&addr_QProcessEnvironment_swap, "QProcessEnvironment_swap")
	symaddr(&addr_QProcessEnvironment_isEmpty, "QProcessEnvironment_isEmpty")
	symaddr(&addr_QProcessEnvironment_inheritsFromParent, "QProcessEnvironment_inheritsFromParent")
	symaddr(&addr_QProcessEnvironment_clear, "QProcessEnvironment_clear")
	symaddr(&addr_QProcessEnvironment_contains, "QProcessEnvironment_contains")
	symaddr(&addr_QProcessEnvironment_insert, "QProcessEnvironment_insert")
	symaddr(&addr_QProcessEnvironment_remove, "QProcessEnvironment_remove")
	symaddr(&addr_QProcessEnvironment_value, "QProcessEnvironment_value")
	symaddr(&addr_QProcessEnvironment_toStringList, "QProcessEnvironment_toStringList")
	symaddr(&addr_QProcessEnvironment_keys, "QProcessEnvironment_keys")
	symaddr(&addr_QProcessEnvironment_insert2, "QProcessEnvironment_insert2")
	symaddr(&addr_QProcessEnvironment_systemEnvironment, "QProcessEnvironment_systemEnvironment")
	symaddr(&addr_QProcessEnvironment_value2, "QProcessEnvironment_value2")
	symaddr(&addr_QProcessEnvironment_OperatorEqual, "QProcessEnvironment_OperatorEqual")
	symaddr(&addr_QProcessEnvironment_OperatorNotEqual, "QProcessEnvironment_OperatorNotEqual")
	reg(&pg_QProcessEnvironment_delete, "QProcessEnvironment_delete")
}

type QProcessEnvironment struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QProcessEnvironment) UnsafeGoID() int64     { return this.vid }
func (this *QProcessEnvironment) UnsafeSetGoID(id int64) { this.vid = id }

func bindQProcessEnvironment(h uintptr) *QProcessEnvironment {
	if h == 0 {
		return nil
	}
	return &QProcessEnvironment{h: h}
}

func newownedQProcessEnvironment(h uintptr) *QProcessEnvironment {
	o := bindQProcessEnvironment(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QProcessEnvironment).Delete)
	}
	return o
}

func (this *QProcessEnvironment) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProcessEnvironment) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProcessEnvironment(h unsafe.Pointer) *QProcessEnvironment { return bindQProcessEnvironment(uintptr(h)) }
func UnsafeNewQProcessEnvironmentOwned(h unsafe.Pointer) *QProcessEnvironment { return newownedQProcessEnvironment(uintptr(h)) }

func (this *QProcessEnvironment) UnsafeRelease() { this.h = 0 }

func NewQProcessEnvironment() *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_new)
	__o := newownedQProcessEnvironment(__r)
	return __o
}

func NewQProcessEnvironment2(arg1 QProcessEnvironment__Initialization) *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_new2, uintptr(arg1))
	__o := newownedQProcessEnvironment(__r)
	return __o
}

func NewQProcessEnvironment3(other *QProcessEnvironment) *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_new3, uintptr(other.UnsafePointer()))
	__o := newownedQProcessEnvironment(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QProcessEnvironment) OperatorAssign(other *QProcessEnvironment) *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQProcessEnvironment(__r)
}

func (this *QProcessEnvironment) Swap(other *QProcessEnvironment) {
	purego.SyscallN(addr_QProcessEnvironment_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QProcessEnvironment) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcessEnvironment_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcessEnvironment) InheritsFromParent() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcessEnvironment_inheritsFromParent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcessEnvironment) Clear() {
	purego.SyscallN(addr_QProcessEnvironment_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcessEnvironment) Contains(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcessEnvironment_contains, this.h, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r != 0
}

func (this *QProcessEnvironment) Insert(name string, value string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QProcessEnvironment_insert, this.h, name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QProcessEnvironment) Remove(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QProcessEnvironment_remove, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QProcessEnvironment) Value(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_value, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QProcessEnvironment) ToStringList() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_toStringList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QProcessEnvironment) Keys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_keys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QProcessEnvironment) InsertWithQProcessEnvironment(e *QProcessEnvironment) {
	purego.SyscallN(addr_QProcessEnvironment_insert2, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func QProcessEnvironment_SystemEnvironment() *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_systemEnvironment)
	return newownedQProcessEnvironment(__r)
}

func (this *QProcessEnvironment) Value2(name string, defaultValue string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcessEnvironment_value2, this.h, name_p, uintptr(len(name)), defaultValue_p, uintptr(len(defaultValue)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(defaultValue)
	return pgStr(__r, __n)
}

func (this *QProcessEnvironment) OperatorEqual(rhs *QProcessEnvironment) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcessEnvironment_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QProcessEnvironment) OperatorNotEqual(rhs *QProcessEnvironment) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcessEnvironment_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QProcessEnvironment) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProcessEnvironment_delete(this.h)
	}
	this.UnsafeRelease()
}
