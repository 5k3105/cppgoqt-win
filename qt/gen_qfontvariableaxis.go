package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontVariableAxis_new uintptr
	addr_QFontVariableAxis_new2 uintptr
	addr_QFontVariableAxis_swap uintptr
	addr_QFontVariableAxis_OperatorAssign uintptr
	addr_QFontVariableAxis_tag uintptr
	addr_QFontVariableAxis_setTag uintptr
	addr_QFontVariableAxis_name uintptr
	addr_QFontVariableAxis_setName uintptr
	pg_QFontVariableAxis_minimumValue func(uintptr) float64
	pg_QFontVariableAxis_setMinimumValue func(uintptr, float64)
	pg_QFontVariableAxis_maximumValue func(uintptr) float64
	pg_QFontVariableAxis_setMaximumValue func(uintptr, float64)
	pg_QFontVariableAxis_defaultValue func(uintptr) float64
	pg_QFontVariableAxis_setDefaultValue func(uintptr, float64)
	pg_QFontVariableAxis_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontVariableAxis_new, "QFontVariableAxis_new")
	symaddr(&addr_QFontVariableAxis_new2, "QFontVariableAxis_new2")
	symaddr(&addr_QFontVariableAxis_swap, "QFontVariableAxis_swap")
	symaddr(&addr_QFontVariableAxis_OperatorAssign, "QFontVariableAxis_OperatorAssign")
	symaddr(&addr_QFontVariableAxis_tag, "QFontVariableAxis_tag")
	symaddr(&addr_QFontVariableAxis_setTag, "QFontVariableAxis_setTag")
	symaddr(&addr_QFontVariableAxis_name, "QFontVariableAxis_name")
	symaddr(&addr_QFontVariableAxis_setName, "QFontVariableAxis_setName")
	reg(&pg_QFontVariableAxis_minimumValue, "QFontVariableAxis_minimumValue")
	reg(&pg_QFontVariableAxis_setMinimumValue, "QFontVariableAxis_setMinimumValue")
	reg(&pg_QFontVariableAxis_maximumValue, "QFontVariableAxis_maximumValue")
	reg(&pg_QFontVariableAxis_setMaximumValue, "QFontVariableAxis_setMaximumValue")
	reg(&pg_QFontVariableAxis_defaultValue, "QFontVariableAxis_defaultValue")
	reg(&pg_QFontVariableAxis_setDefaultValue, "QFontVariableAxis_setDefaultValue")
	reg(&pg_QFontVariableAxis_delete, "QFontVariableAxis_delete")
}

type QFontVariableAxis struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFontVariableAxis) UnsafeGoID() int64     { return this.vid }
func (this *QFontVariableAxis) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFontVariableAxis(h uintptr) *QFontVariableAxis {
	if h == 0 {
		return nil
	}
	return &QFontVariableAxis{h: h}
}

func newownedQFontVariableAxis(h uintptr) *QFontVariableAxis {
	o := bindQFontVariableAxis(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFontVariableAxis).Delete)
	}
	return o
}

func (this *QFontVariableAxis) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontVariableAxis) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontVariableAxis(h unsafe.Pointer) *QFontVariableAxis { return bindQFontVariableAxis(uintptr(h)) }
func UnsafeNewQFontVariableAxisOwned(h unsafe.Pointer) *QFontVariableAxis { return newownedQFontVariableAxis(uintptr(h)) }

func (this *QFontVariableAxis) UnsafeRelease() { this.h = 0 }

func NewQFontVariableAxis() *QFontVariableAxis {
	__r, _, _ := purego.SyscallN(addr_QFontVariableAxis_new)
	__o := newownedQFontVariableAxis(__r)
	return __o
}

func NewQFontVariableAxis2(axis *QFontVariableAxis) *QFontVariableAxis {
	__r, _, _ := purego.SyscallN(addr_QFontVariableAxis_new2, uintptr(axis.UnsafePointer()))
	__o := newownedQFontVariableAxis(__r)
	runtime.KeepAlive(axis)
	return __o
}

func (this *QFontVariableAxis) Swap(other *QFontVariableAxis) {
	purego.SyscallN(addr_QFontVariableAxis_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFontVariableAxis) OperatorAssign(axis *QFontVariableAxis) *QFontVariableAxis {
	__r, _, _ := purego.SyscallN(addr_QFontVariableAxis_OperatorAssign, this.h, uintptr(axis.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(axis)
	return bindQFontVariableAxis(__r)
}

func (this *QFontVariableAxis) Tag() *QFont__Tag {
	__r, _, _ := purego.SyscallN(addr_QFontVariableAxis_tag, this.h)
	runtime.KeepAlive(this)
	return newownedQFont__Tag(__r)
}

func (this *QFontVariableAxis) SetTag(tag *QFont__Tag) {
	purego.SyscallN(addr_QFontVariableAxis_setTag, this.h, uintptr(tag.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tag)
}

func (this *QFontVariableAxis) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontVariableAxis_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontVariableAxis) SetName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QFontVariableAxis_setName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QFontVariableAxis) MinimumValue() float64 {
	__r := pg_QFontVariableAxis_minimumValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontVariableAxis) SetMinimumValue(minimumValue float64) {
	pg_QFontVariableAxis_setMinimumValue(this.h, minimumValue)
	runtime.KeepAlive(this)
}

func (this *QFontVariableAxis) MaximumValue() float64 {
	__r := pg_QFontVariableAxis_maximumValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontVariableAxis) SetMaximumValue(maximumValue float64) {
	pg_QFontVariableAxis_setMaximumValue(this.h, maximumValue)
	runtime.KeepAlive(this)
}

func (this *QFontVariableAxis) DefaultValue() float64 {
	__r := pg_QFontVariableAxis_defaultValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontVariableAxis) SetDefaultValue(defaultValue float64) {
	pg_QFontVariableAxis_setDefaultValue(this.h, defaultValue)
	runtime.KeepAlive(this)
}

func (this *QFontVariableAxis) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontVariableAxis_delete(this.h)
	}
	this.UnsafeRelease()
}
