package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QScrollerProperties_new uintptr
	addr_QScrollerProperties_new2 uintptr
	addr_QScrollerProperties_OperatorAssign uintptr
	addr_QScrollerProperties_OperatorEqual uintptr
	addr_QScrollerProperties_OperatorNotEqual uintptr
	addr_QScrollerProperties_setDefaultScrollerProperties uintptr
	addr_QScrollerProperties_unsetDefaultScrollerProperties uintptr
	addr_QScrollerProperties_scrollMetric uintptr
	addr_QScrollerProperties_setScrollMetric uintptr
	pg_QScrollerProperties_delete func(uintptr)
)

func init() {
	symaddr(&addr_QScrollerProperties_new, "QScrollerProperties_new")
	symaddr(&addr_QScrollerProperties_new2, "QScrollerProperties_new2")
	symaddr(&addr_QScrollerProperties_OperatorAssign, "QScrollerProperties_OperatorAssign")
	symaddr(&addr_QScrollerProperties_OperatorEqual, "QScrollerProperties_OperatorEqual")
	symaddr(&addr_QScrollerProperties_OperatorNotEqual, "QScrollerProperties_OperatorNotEqual")
	symaddr(&addr_QScrollerProperties_setDefaultScrollerProperties, "QScrollerProperties_setDefaultScrollerProperties")
	symaddr(&addr_QScrollerProperties_unsetDefaultScrollerProperties, "QScrollerProperties_unsetDefaultScrollerProperties")
	symaddr(&addr_QScrollerProperties_scrollMetric, "QScrollerProperties_scrollMetric")
	symaddr(&addr_QScrollerProperties_setScrollMetric, "QScrollerProperties_setScrollMetric")
	reg(&pg_QScrollerProperties_delete, "QScrollerProperties_delete")
}

type QScrollerProperties struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QScrollerProperties) UnsafeGoID() int64     { return this.vid }
func (this *QScrollerProperties) UnsafeSetGoID(id int64) { this.vid = id }

func bindQScrollerProperties(h uintptr) *QScrollerProperties {
	if h == 0 {
		return nil
	}
	return &QScrollerProperties{h: h}
}

func newownedQScrollerProperties(h uintptr) *QScrollerProperties {
	o := bindQScrollerProperties(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QScrollerProperties).Delete)
	}
	return o
}

func (this *QScrollerProperties) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScrollerProperties) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScrollerProperties(h unsafe.Pointer) *QScrollerProperties { return bindQScrollerProperties(uintptr(h)) }
func UnsafeNewQScrollerPropertiesOwned(h unsafe.Pointer) *QScrollerProperties { return newownedQScrollerProperties(uintptr(h)) }

func (this *QScrollerProperties) UnsafeRelease() { this.h = 0 }

func NewQScrollerProperties() *QScrollerProperties {
	__r, _, _ := purego.SyscallN(addr_QScrollerProperties_new)
	__o := newownedQScrollerProperties(__r)
	return __o
}

func NewQScrollerProperties2(sp *QScrollerProperties) *QScrollerProperties {
	__r, _, _ := purego.SyscallN(addr_QScrollerProperties_new2, uintptr(sp.UnsafePointer()))
	__o := newownedQScrollerProperties(__r)
	runtime.KeepAlive(sp)
	return __o
}

func (this *QScrollerProperties) OperatorAssign(sp *QScrollerProperties) *QScrollerProperties {
	__r, _, _ := purego.SyscallN(addr_QScrollerProperties_OperatorAssign, this.h, uintptr(sp.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sp)
	return bindQScrollerProperties(__r)
}

func (this *QScrollerProperties) OperatorEqual(sp *QScrollerProperties) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollerProperties_OperatorEqual, this.h, uintptr(sp.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sp)
	return __r != 0
}

func (this *QScrollerProperties) OperatorNotEqual(sp *QScrollerProperties) bool {
	__r0, _, _ := purego.SyscallN(addr_QScrollerProperties_OperatorNotEqual, this.h, uintptr(sp.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sp)
	return __r != 0
}

func QScrollerProperties_SetDefaultScrollerProperties(sp *QScrollerProperties) {
	purego.SyscallN(addr_QScrollerProperties_setDefaultScrollerProperties, uintptr(sp.UnsafePointer()))
	runtime.KeepAlive(sp)
}

func QScrollerProperties_UnsetDefaultScrollerProperties() {
	purego.SyscallN(addr_QScrollerProperties_unsetDefaultScrollerProperties)
}

func (this *QScrollerProperties) ScrollMetric(metric QScrollerProperties__ScrollMetric) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QScrollerProperties_scrollMetric, this.h, uintptr(metric))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QScrollerProperties) SetScrollMetric(metric QScrollerProperties__ScrollMetric, value *QVariant) {
	purego.SyscallN(addr_QScrollerProperties_setScrollMetric, this.h, uintptr(metric), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QScrollerProperties) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScrollerProperties_delete(this.h)
	}
	this.UnsafeRelease()
}
