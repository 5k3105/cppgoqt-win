package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionTabBarBase_base func(uintptr) uintptr
	addr_QStyleOptionTabBarBase_new uintptr
	addr_QStyleOptionTabBarBase_new2 uintptr
	addr_QStyleOptionTabBarBase_OperatorAssign uintptr
	addr_QStyleOptionTabBarBase_shape_get uintptr
	addr_QStyleOptionTabBarBase_shape_set uintptr
	addr_QStyleOptionTabBarBase_tabBarRect_get uintptr
	addr_QStyleOptionTabBarBase_tabBarRect_set uintptr
	addr_QStyleOptionTabBarBase_selectedTabRect_get uintptr
	addr_QStyleOptionTabBarBase_selectedTabRect_set uintptr
	addr_QStyleOptionTabBarBase_documentMode_get uintptr
	addr_QStyleOptionTabBarBase_documentMode_set uintptr
	pg_QStyleOptionTabBarBase_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionTabBarBase_base, "QStyleOptionTabBarBase_base")
	symaddr(&addr_QStyleOptionTabBarBase_new, "QStyleOptionTabBarBase_new")
	symaddr(&addr_QStyleOptionTabBarBase_new2, "QStyleOptionTabBarBase_new2")
	symaddr(&addr_QStyleOptionTabBarBase_OperatorAssign, "QStyleOptionTabBarBase_OperatorAssign")
	symaddr(&addr_QStyleOptionTabBarBase_shape_get, "QStyleOptionTabBarBase_shape_get")
	symaddr(&addr_QStyleOptionTabBarBase_shape_set, "QStyleOptionTabBarBase_shape_set")
	symaddr(&addr_QStyleOptionTabBarBase_tabBarRect_get, "QStyleOptionTabBarBase_tabBarRect_get")
	symaddr(&addr_QStyleOptionTabBarBase_tabBarRect_set, "QStyleOptionTabBarBase_tabBarRect_set")
	symaddr(&addr_QStyleOptionTabBarBase_selectedTabRect_get, "QStyleOptionTabBarBase_selectedTabRect_get")
	symaddr(&addr_QStyleOptionTabBarBase_selectedTabRect_set, "QStyleOptionTabBarBase_selectedTabRect_set")
	symaddr(&addr_QStyleOptionTabBarBase_documentMode_get, "QStyleOptionTabBarBase_documentMode_get")
	symaddr(&addr_QStyleOptionTabBarBase_documentMode_set, "QStyleOptionTabBarBase_documentMode_set")
	reg(&pg_QStyleOptionTabBarBase_delete, "QStyleOptionTabBarBase_delete")
}

type QStyleOptionTabBarBase struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionTabBarBase(h uintptr) *QStyleOptionTabBarBase {
	if h == 0 {
		return nil
	}
	return &QStyleOptionTabBarBase{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionTabBarBase_base(h))}
}

func newownedQStyleOptionTabBarBase(h uintptr) *QStyleOptionTabBarBase {
	o := bindQStyleOptionTabBarBase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionTabBarBase).Delete)
	}
	return o
}

func (this *QStyleOptionTabBarBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionTabBarBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionTabBarBase(h unsafe.Pointer) *QStyleOptionTabBarBase { return bindQStyleOptionTabBarBase(uintptr(h)) }
func UnsafeNewQStyleOptionTabBarBaseOwned(h unsafe.Pointer) *QStyleOptionTabBarBase { return newownedQStyleOptionTabBarBase(uintptr(h)) }

func (this *QStyleOptionTabBarBase) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionTabBarBase() *QStyleOptionTabBarBase {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_new)
	__o := newownedQStyleOptionTabBarBase(__r)
	return __o
}

func NewQStyleOptionTabBarBase2(other *QStyleOptionTabBarBase) *QStyleOptionTabBarBase {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionTabBarBase(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionTabBarBase) OperatorAssign(arg1 *QStyleOptionTabBarBase) *QStyleOptionTabBarBase {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionTabBarBase(__r)
}

func (this *QStyleOptionTabBarBase) Shape() QTabBar__Shape {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_shape_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabBar__Shape(__r)
}

func (this *QStyleOptionTabBarBase) SetShape(shape QTabBar__Shape) {
	purego.SyscallN(addr_QStyleOptionTabBarBase_shape_set, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTabBarBase) TabBarRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_tabBarRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionTabBarBase) SetTabBarRect(tabBarRect *QRect) {
	purego.SyscallN(addr_QStyleOptionTabBarBase_tabBarRect_set, this.h, uintptr(tabBarRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabBarRect)
}

func (this *QStyleOptionTabBarBase) SelectedTabRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_selectedTabRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionTabBarBase) SetSelectedTabRect(selectedTabRect *QRect) {
	purego.SyscallN(addr_QStyleOptionTabBarBase_selectedTabRect_set, this.h, uintptr(selectedTabRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectedTabRect)
}

func (this *QStyleOptionTabBarBase) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTabBarBase_documentMode_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionTabBarBase) SetDocumentMode(documentMode bool) {
	purego.SyscallN(addr_QStyleOptionTabBarBase_documentMode_set, this.h, uintptr(b2i(documentMode)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTabBarBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionTabBarBase_delete(this.h)
	}
	this.UnsafeRelease()
}
