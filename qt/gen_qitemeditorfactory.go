package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QItemEditorFactory_new uintptr
	addr_QItemEditorFactory_new2 uintptr
	addr_QItemEditorFactory_createEditor uintptr
	addr_QItemEditorFactory_valuePropertyName uintptr
	addr_QItemEditorFactory_registerEditor uintptr
	addr_QItemEditorFactory_defaultFactory uintptr
	addr_QItemEditorFactory_setDefaultFactory uintptr
	addr_QItemEditorFactory_OperatorAssign uintptr
	addr_QItemEditorFactory_createEditor_default uintptr
	addr_QItemEditorFactory_valuePropertyName_default uintptr
	pg_QItemEditorFactory__set_override func(uintptr, int32)
	pg_QItemEditorFactory_delete func(uintptr)
)

func init() {
	symaddr(&addr_QItemEditorFactory_new, "QItemEditorFactory_new")
	symaddr(&addr_QItemEditorFactory_new2, "QItemEditorFactory_new2")
	symaddr(&addr_QItemEditorFactory_createEditor, "QItemEditorFactory_createEditor")
	symaddr(&addr_QItemEditorFactory_valuePropertyName, "QItemEditorFactory_valuePropertyName")
	symaddr(&addr_QItemEditorFactory_registerEditor, "QItemEditorFactory_registerEditor")
	symaddr(&addr_QItemEditorFactory_defaultFactory, "QItemEditorFactory_defaultFactory")
	symaddr(&addr_QItemEditorFactory_setDefaultFactory, "QItemEditorFactory_setDefaultFactory")
	symaddr(&addr_QItemEditorFactory_OperatorAssign, "QItemEditorFactory_OperatorAssign")
	symaddr(&addr_QItemEditorFactory_createEditor_default, "QItemEditorFactory_createEditor_default")
	symaddr(&addr_QItemEditorFactory_valuePropertyName_default, "QItemEditorFactory_valuePropertyName_default")
	reg(&pg_QItemEditorFactory__set_override, "QItemEditorFactory__set_override")
	reg(&pg_QItemEditorFactory_delete, "QItemEditorFactory_delete")
}

type QItemEditorFactory struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QItemEditorFactory) UnsafeGoID() int64     { return this.vid }
func (this *QItemEditorFactory) UnsafeSetGoID(id int64) { this.vid = id }

func bindQItemEditorFactory(h uintptr) *QItemEditorFactory {
	if h == 0 {
		return nil
	}
	return &QItemEditorFactory{h: h}
}

func newownedQItemEditorFactory(h uintptr) *QItemEditorFactory {
	o := bindQItemEditorFactory(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QItemEditorFactory).Delete)
	}
	return o
}

func (this *QItemEditorFactory) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemEditorFactory) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemEditorFactory(h unsafe.Pointer) *QItemEditorFactory { return bindQItemEditorFactory(uintptr(h)) }
func UnsafeNewQItemEditorFactoryOwned(h unsafe.Pointer) *QItemEditorFactory { return newownedQItemEditorFactory(uintptr(h)) }

func (this *QItemEditorFactory) UnsafeRelease() { this.h = 0 }

func NewQItemEditorFactory() *QItemEditorFactory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_new, uintptr(__goid))
	__o := newownedQItemEditorFactory(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemEditorFactory_virtIdx, QItemEditorFactory_enable)
	return __o
}

func NewQItemEditorFactory2(arg1 *QItemEditorFactory) *QItemEditorFactory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_new2, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQItemEditorFactory(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemEditorFactory_virtIdx, QItemEditorFactory_enable)
	return __o
}

func (this *QItemEditorFactory) CreateEditor(userType int, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_createEditor, this.h, uintptr(userType), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QItemEditorFactory) ValuePropertyName(userType int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_valuePropertyName, this.h, uintptr(userType))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QItemEditorFactory) RegisterEditor(userType int, creator *QItemEditorCreatorBase) {
	purego.SyscallN(addr_QItemEditorFactory_registerEditor, this.h, uintptr(userType), uintptr(creator.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(creator)
}

func QItemEditorFactory_DefaultFactory() *QItemEditorFactory {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_defaultFactory)
	return bindQItemEditorFactory(__r)
}

func QItemEditorFactory_SetDefaultFactory(factory *QItemEditorFactory) {
	purego.SyscallN(addr_QItemEditorFactory_setDefaultFactory, uintptr(factory.UnsafePointer()))
	runtime.KeepAlive(factory)
}

func (this *QItemEditorFactory) OperatorAssign(arg1 *QItemEditorFactory) *QItemEditorFactory {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQItemEditorFactory(__r)
}

func (this *QItemEditorFactory) CreateEditorDefault(userType int, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_createEditor_default, this.h, uintptr(userType), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QItemEditorFactory) ValuePropertyNameDefault(userType int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QItemEditorFactory_valuePropertyName_default, this.h, uintptr(userType))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QItemEditorFactory_virtIdx(slot int64) int32 {
	switch slot {
	case -2118980437705941456:
		return 0
	case -3548500303342436100:
		return 1
	}
	return -1
}

func QItemEditorFactory_enable(h uintptr, idx int32) { pg_QItemEditorFactory__set_override(h, idx) }

func (this *QItemEditorFactory) OnCreateEditor(fn func(int, *QWidget) *QWidget) {
	registerVirt(this.UnsafeGoID(), -2118980437705941456, func(userType int32, parent uintptr) uintptr { return uintptr(fn(int(userType), UnsafeNewQWidget(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "createEditor"
}

func (this *QItemEditorFactory) OnValuePropertyName(fn func(int) *QByteArray) {
	registerVirt(this.UnsafeGoID(), -3548500303342436100, func(userType int32) uintptr { return uintptr(fn(int(userType)).UnsafePointer()) }) // slot "valuePropertyName"
}

func (this *QItemEditorFactory) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemEditorFactory_delete(this.h)
	}
	this.UnsafeRelease()
}
