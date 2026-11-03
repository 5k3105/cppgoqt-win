package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleAttributesInterface_attributeKeys uintptr
	addr_QAccessibleAttributesInterface_attributeValue uintptr
	addr_QAccessibleAttributesInterface_OperatorAssign uintptr
	pg_QAccessibleAttributesInterface__set_override func(uintptr, int32)
	pg_QAccessibleAttributesInterface__Go_new func(int64) uintptr
	pg_QAccessibleAttributesInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleAttributesInterface_attributeKeys, "QAccessibleAttributesInterface_attributeKeys")
	symaddr(&addr_QAccessibleAttributesInterface_attributeValue, "QAccessibleAttributesInterface_attributeValue")
	symaddr(&addr_QAccessibleAttributesInterface_OperatorAssign, "QAccessibleAttributesInterface_OperatorAssign")
	reg(&pg_QAccessibleAttributesInterface__set_override, "QAccessibleAttributesInterface__set_override")
	reg(&pg_QAccessibleAttributesInterface__Go_new, "QAccessibleAttributesInterface__Go_new")
	reg(&pg_QAccessibleAttributesInterface_delete, "QAccessibleAttributesInterface_delete")
}

type QAccessibleAttributesInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleAttributesInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleAttributesInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleAttributesInterface(h uintptr) *QAccessibleAttributesInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleAttributesInterface{h: h}
}

func newownedQAccessibleAttributesInterface(h uintptr) *QAccessibleAttributesInterface {
	o := bindQAccessibleAttributesInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleAttributesInterface).Delete)
	}
	return o
}

func (this *QAccessibleAttributesInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleAttributesInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleAttributesInterface(h unsafe.Pointer) *QAccessibleAttributesInterface { return bindQAccessibleAttributesInterface(uintptr(h)) }
func UnsafeNewQAccessibleAttributesInterfaceOwned(h unsafe.Pointer) *QAccessibleAttributesInterface { return newownedQAccessibleAttributesInterface(uintptr(h)) }

func (this *QAccessibleAttributesInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleAttributesInterface) AttributeKeys() []QAccessible__Attribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleAttributesInterface_attributeKeys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]QAccessible__Attribute, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QAccessible__Attribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleAttributesInterface) AttributeValue(key QAccessible__Attribute) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleAttributesInterface_attributeValue, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAccessibleAttributesInterface) OperatorAssign(arg1 *QAccessibleAttributesInterface) *QAccessibleAttributesInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleAttributesInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleAttributesInterface(__r)
}

func QAccessibleAttributesInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 9162027173711402047:
		return 0
	case 4705686874656021092:
		return 1
	}
	return -1
}

func QAccessibleAttributesInterface_enable(h uintptr, idx int32) { pg_QAccessibleAttributesInterface__set_override(h, idx) }

func NewQAccessibleAttributesInterface() *QAccessibleAttributesInterface {
	id := nextVirtID()
	o := newownedQAccessibleAttributesInterface(pg_QAccessibleAttributesInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleAttributesInterface_virtIdx, QAccessibleAttributesInterface_enable)
	return o
}

func (this *QAccessibleAttributesInterface) OnAttributeKeys(fn func() []QAccessible__Attribute) {
	registerVirt(this.UnsafeGoID(), 9162027173711402047, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(int32(0)))
		__sl := unsafe.Slice((*int32)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = int32(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "attributeKeys"
}

func (this *QAccessibleAttributesInterface) OnAttributeValue(fn func(QAccessible__Attribute) *QVariant) {
	registerVirt(this.UnsafeGoID(), 4705686874656021092, func(key int32) uintptr { return uintptr(fn(QAccessible__Attribute(key)).UnsafePointer()) }) // slot "attributeValue"
}

func (this *QAccessibleAttributesInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleAttributesInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
