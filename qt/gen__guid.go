package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr__GUID_new uintptr
	addr__GUID_new2 uintptr
	addr__GUID_Data1_get uintptr
	addr__GUID_Data1_set uintptr
	addr__GUID_Data2_get uintptr
	addr__GUID_Data2_set uintptr
	addr__GUID_Data3_get uintptr
	addr__GUID_Data3_set uintptr
	pg__GUID_delete func(uintptr)
)

func init() {
	symaddr(&addr__GUID_new, "_GUID_new")
	symaddr(&addr__GUID_new2, "_GUID_new2")
	symaddr(&addr__GUID_Data1_get, "_GUID_Data1_get")
	symaddr(&addr__GUID_Data1_set, "_GUID_Data1_set")
	symaddr(&addr__GUID_Data2_get, "_GUID_Data2_get")
	symaddr(&addr__GUID_Data2_set, "_GUID_Data2_set")
	symaddr(&addr__GUID_Data3_get, "_GUID_Data3_get")
	symaddr(&addr__GUID_Data3_set, "_GUID_Data3_set")
	reg(&pg__GUID_delete, "_GUID_delete")
}

type _GUID struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *_GUID) UnsafeGoID() int64     { return this.vid }
func (this *_GUID) UnsafeSetGoID(id int64) { this.vid = id }

func bind_GUID(h uintptr) *_GUID {
	if h == 0 {
		return nil
	}
	return &_GUID{h: h}
}

func newowned_GUID(h uintptr) *_GUID {
	o := bind_GUID(h)
	if o != nil {
		runtime.SetFinalizer(o, (*_GUID).Delete)
	}
	return o
}

func (this *_GUID) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *_GUID) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNew_GUID(h unsafe.Pointer) *_GUID { return bind_GUID(uintptr(h)) }
func UnsafeNew_GUIDOwned(h unsafe.Pointer) *_GUID { return newowned_GUID(uintptr(h)) }

func (this *_GUID) UnsafeRelease() { this.h = 0 }

func New_GUID() *_GUID {
	__r, _, _ := purego.SyscallN(addr__GUID_new)
	__o := newowned_GUID(__r)
	return __o
}

func New_GUID2(arg1 *_GUID) *_GUID {
	__r, _, _ := purego.SyscallN(addr__GUID_new2, uintptr(arg1.UnsafePointer()))
	__o := newowned_GUID(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *_GUID) Data1() uint32 {
	__r0, _, _ := purego.SyscallN(addr__GUID_Data1_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *_GUID) SetData1(Data1 uint32) {
	purego.SyscallN(addr__GUID_Data1_set, this.h, uintptr(Data1))
	runtime.KeepAlive(this)
}

func (this *_GUID) Data2() uint16 {
	__r0, _, _ := purego.SyscallN(addr__GUID_Data2_get, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *_GUID) SetData2(Data2 uint16) {
	purego.SyscallN(addr__GUID_Data2_set, this.h, uintptr(Data2))
	runtime.KeepAlive(this)
}

func (this *_GUID) Data3() uint16 {
	__r0, _, _ := purego.SyscallN(addr__GUID_Data3_get, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *_GUID) SetData3(Data3 uint16) {
	purego.SyscallN(addr__GUID_Data3_set, this.h, uintptr(Data3))
	runtime.KeepAlive(this)
}

func (this *_GUID) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg__GUID_delete(this.h)
	}
	this.UnsafeRelease()
}
