package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionDockWidget_base func(uintptr) uintptr
	addr_QStyleOptionDockWidget_new uintptr
	addr_QStyleOptionDockWidget_new2 uintptr
	addr_QStyleOptionDockWidget_OperatorAssign uintptr
	addr_QStyleOptionDockWidget_title_get uintptr
	addr_QStyleOptionDockWidget_title_set uintptr
	addr_QStyleOptionDockWidget_closable_get uintptr
	addr_QStyleOptionDockWidget_closable_set uintptr
	addr_QStyleOptionDockWidget_movable_get uintptr
	addr_QStyleOptionDockWidget_movable_set uintptr
	addr_QStyleOptionDockWidget_floatable_get uintptr
	addr_QStyleOptionDockWidget_floatable_set uintptr
	addr_QStyleOptionDockWidget_verticalTitleBar_get uintptr
	addr_QStyleOptionDockWidget_verticalTitleBar_set uintptr
	pg_QStyleOptionDockWidget_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionDockWidget_base, "QStyleOptionDockWidget_base")
	symaddr(&addr_QStyleOptionDockWidget_new, "QStyleOptionDockWidget_new")
	symaddr(&addr_QStyleOptionDockWidget_new2, "QStyleOptionDockWidget_new2")
	symaddr(&addr_QStyleOptionDockWidget_OperatorAssign, "QStyleOptionDockWidget_OperatorAssign")
	symaddr(&addr_QStyleOptionDockWidget_title_get, "QStyleOptionDockWidget_title_get")
	symaddr(&addr_QStyleOptionDockWidget_title_set, "QStyleOptionDockWidget_title_set")
	symaddr(&addr_QStyleOptionDockWidget_closable_get, "QStyleOptionDockWidget_closable_get")
	symaddr(&addr_QStyleOptionDockWidget_closable_set, "QStyleOptionDockWidget_closable_set")
	symaddr(&addr_QStyleOptionDockWidget_movable_get, "QStyleOptionDockWidget_movable_get")
	symaddr(&addr_QStyleOptionDockWidget_movable_set, "QStyleOptionDockWidget_movable_set")
	symaddr(&addr_QStyleOptionDockWidget_floatable_get, "QStyleOptionDockWidget_floatable_get")
	symaddr(&addr_QStyleOptionDockWidget_floatable_set, "QStyleOptionDockWidget_floatable_set")
	symaddr(&addr_QStyleOptionDockWidget_verticalTitleBar_get, "QStyleOptionDockWidget_verticalTitleBar_get")
	symaddr(&addr_QStyleOptionDockWidget_verticalTitleBar_set, "QStyleOptionDockWidget_verticalTitleBar_set")
	reg(&pg_QStyleOptionDockWidget_delete, "QStyleOptionDockWidget_delete")
}

type QStyleOptionDockWidget struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionDockWidget(h uintptr) *QStyleOptionDockWidget {
	if h == 0 {
		return nil
	}
	return &QStyleOptionDockWidget{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionDockWidget_base(h))}
}

func newownedQStyleOptionDockWidget(h uintptr) *QStyleOptionDockWidget {
	o := bindQStyleOptionDockWidget(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionDockWidget).Delete)
	}
	return o
}

func (this *QStyleOptionDockWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionDockWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionDockWidget(h unsafe.Pointer) *QStyleOptionDockWidget { return bindQStyleOptionDockWidget(uintptr(h)) }
func UnsafeNewQStyleOptionDockWidgetOwned(h unsafe.Pointer) *QStyleOptionDockWidget { return newownedQStyleOptionDockWidget(uintptr(h)) }

func (this *QStyleOptionDockWidget) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionDockWidget() *QStyleOptionDockWidget {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_new)
	__o := newownedQStyleOptionDockWidget(__r)
	return __o
}

func NewQStyleOptionDockWidget2(other *QStyleOptionDockWidget) *QStyleOptionDockWidget {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionDockWidget(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionDockWidget) OperatorAssign(arg1 *QStyleOptionDockWidget) *QStyleOptionDockWidget {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionDockWidget(__r)
}

func (this *QStyleOptionDockWidget) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_title_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionDockWidget) SetTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QStyleOptionDockWidget_title_set, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QStyleOptionDockWidget) Closable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_closable_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionDockWidget) SetClosable(closable bool) {
	purego.SyscallN(addr_QStyleOptionDockWidget_closable_set, this.h, uintptr(b2i(closable)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionDockWidget) Movable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_movable_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionDockWidget) SetMovable(movable bool) {
	purego.SyscallN(addr_QStyleOptionDockWidget_movable_set, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionDockWidget) Floatable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_floatable_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionDockWidget) SetFloatable(floatable bool) {
	purego.SyscallN(addr_QStyleOptionDockWidget_floatable_set, this.h, uintptr(b2i(floatable)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionDockWidget) VerticalTitleBar() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionDockWidget_verticalTitleBar_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionDockWidget) SetVerticalTitleBar(verticalTitleBar bool) {
	purego.SyscallN(addr_QStyleOptionDockWidget_verticalTitleBar_set, this.h, uintptr(b2i(verticalTitleBar)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionDockWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionDockWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
