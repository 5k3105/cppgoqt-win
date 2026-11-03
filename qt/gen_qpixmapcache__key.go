package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPixmapCache__Key_new uintptr
	addr_QPixmapCache__Key_new2 uintptr
	addr_QPixmapCache__Key_OperatorEqual uintptr
	addr_QPixmapCache__Key_OperatorNotEqual uintptr
	addr_QPixmapCache__Key_OperatorAssign uintptr
	addr_QPixmapCache__Key_swap uintptr
	addr_QPixmapCache__Key_isValid uintptr
	pg_QPixmapCache__Key_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPixmapCache__Key_new, "QPixmapCache__Key_new")
	symaddr(&addr_QPixmapCache__Key_new2, "QPixmapCache__Key_new2")
	symaddr(&addr_QPixmapCache__Key_OperatorEqual, "QPixmapCache__Key_OperatorEqual")
	symaddr(&addr_QPixmapCache__Key_OperatorNotEqual, "QPixmapCache__Key_OperatorNotEqual")
	symaddr(&addr_QPixmapCache__Key_OperatorAssign, "QPixmapCache__Key_OperatorAssign")
	symaddr(&addr_QPixmapCache__Key_swap, "QPixmapCache__Key_swap")
	symaddr(&addr_QPixmapCache__Key_isValid, "QPixmapCache__Key_isValid")
	reg(&pg_QPixmapCache__Key_delete, "QPixmapCache__Key_delete")
}

type QPixmapCache__Key struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPixmapCache__Key) UnsafeGoID() int64     { return this.vid }
func (this *QPixmapCache__Key) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPixmapCache__Key(h uintptr) *QPixmapCache__Key {
	if h == 0 {
		return nil
	}
	return &QPixmapCache__Key{h: h}
}

func newownedQPixmapCache__Key(h uintptr) *QPixmapCache__Key {
	o := bindQPixmapCache__Key(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPixmapCache__Key).Delete)
	}
	return o
}

func (this *QPixmapCache__Key) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPixmapCache__Key) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPixmapCache__Key(h unsafe.Pointer) *QPixmapCache__Key { return bindQPixmapCache__Key(uintptr(h)) }
func UnsafeNewQPixmapCache__KeyOwned(h unsafe.Pointer) *QPixmapCache__Key { return newownedQPixmapCache__Key(uintptr(h)) }

func (this *QPixmapCache__Key) UnsafeRelease() { this.h = 0 }

func NewQPixmapCache__Key() *QPixmapCache__Key {
	__r, _, _ := purego.SyscallN(addr_QPixmapCache__Key_new)
	__o := newownedQPixmapCache__Key(__r)
	return __o
}

func NewQPixmapCache__Key2(other *QPixmapCache__Key) *QPixmapCache__Key {
	__r, _, _ := purego.SyscallN(addr_QPixmapCache__Key_new2, uintptr(other.UnsafePointer()))
	__o := newownedQPixmapCache__Key(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QPixmapCache__Key) OperatorEqual(key *QPixmapCache__Key) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache__Key_OperatorEqual, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QPixmapCache__Key) OperatorNotEqual(key *QPixmapCache__Key) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache__Key_OperatorNotEqual, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QPixmapCache__Key) OperatorAssign(other *QPixmapCache__Key) *QPixmapCache__Key {
	__r, _, _ := purego.SyscallN(addr_QPixmapCache__Key_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPixmapCache__Key(__r)
}

func (this *QPixmapCache__Key) Swap(other *QPixmapCache__Key) {
	purego.SyscallN(addr_QPixmapCache__Key_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPixmapCache__Key) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache__Key_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmapCache__Key) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPixmapCache__Key_delete(this.h)
	}
	this.UnsafeRelease()
}
