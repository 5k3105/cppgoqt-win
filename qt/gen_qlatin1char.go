package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLatin1Char_new uintptr
	addr_QLatin1Char_new2 uintptr
	addr_QLatin1Char_toLatin1 uintptr
	addr_QLatin1Char_unicode uintptr
	addr_QLatin1Char_OperatorEqual uintptr
	addr_QLatin1Char_OperatorNotEqual uintptr
	addr_QLatin1Char_OperatorLesser uintptr
	addr_QLatin1Char_OperatorGreater uintptr
	addr_QLatin1Char_OperatorLesserOrEqual uintptr
	addr_QLatin1Char_OperatorGreaterOrEqual uintptr
	pg_QLatin1Char_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLatin1Char_new, "QLatin1Char_new")
	symaddr(&addr_QLatin1Char_new2, "QLatin1Char_new2")
	symaddr(&addr_QLatin1Char_toLatin1, "QLatin1Char_toLatin1")
	symaddr(&addr_QLatin1Char_unicode, "QLatin1Char_unicode")
	symaddr(&addr_QLatin1Char_OperatorEqual, "QLatin1Char_OperatorEqual")
	symaddr(&addr_QLatin1Char_OperatorNotEqual, "QLatin1Char_OperatorNotEqual")
	symaddr(&addr_QLatin1Char_OperatorLesser, "QLatin1Char_OperatorLesser")
	symaddr(&addr_QLatin1Char_OperatorGreater, "QLatin1Char_OperatorGreater")
	symaddr(&addr_QLatin1Char_OperatorLesserOrEqual, "QLatin1Char_OperatorLesserOrEqual")
	symaddr(&addr_QLatin1Char_OperatorGreaterOrEqual, "QLatin1Char_OperatorGreaterOrEqual")
	reg(&pg_QLatin1Char_delete, "QLatin1Char_delete")
}

type QLatin1Char struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLatin1Char) UnsafeGoID() int64     { return this.vid }
func (this *QLatin1Char) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLatin1Char(h uintptr) *QLatin1Char {
	if h == 0 {
		return nil
	}
	return &QLatin1Char{h: h}
}

func newownedQLatin1Char(h uintptr) *QLatin1Char {
	o := bindQLatin1Char(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLatin1Char).Delete)
	}
	return o
}

func (this *QLatin1Char) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLatin1Char) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLatin1Char(h unsafe.Pointer) *QLatin1Char { return bindQLatin1Char(uintptr(h)) }
func UnsafeNewQLatin1CharOwned(h unsafe.Pointer) *QLatin1Char { return newownedQLatin1Char(uintptr(h)) }

func (this *QLatin1Char) UnsafeRelease() { this.h = 0 }

func NewQLatin1Char(c int8) *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1Char_new, uintptr(c))
	__o := newownedQLatin1Char(__r)
	return __o
}

func NewQLatin1Char2(arg1 *QLatin1Char) *QLatin1Char {
	__r, _, _ := purego.SyscallN(addr_QLatin1Char_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQLatin1Char(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QLatin1Char) ToLatin1() int8 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_toLatin1, this.h)
	__r := int8(uint8(__r0))
	runtime.KeepAlive(this)
	return __r
}

func (this *QLatin1Char) Unicode() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_unicode, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QLatin1Char) OperatorEqual(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) OperatorNotEqual(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) OperatorLesser(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) OperatorGreater(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) OperatorLesserOrEqual(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) OperatorGreaterOrEqual(rhs *QLatin1Char) bool {
	__r0, _, _ := purego.SyscallN(addr_QLatin1Char_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLatin1Char) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLatin1Char_delete(this.h)
	}
	this.UnsafeRelease()
}
