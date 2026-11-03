package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionTitleBar_chain uintptr
	addr_QStyleOptionTitleBar_new uintptr
	addr_QStyleOptionTitleBar_new2 uintptr
	addr_QStyleOptionTitleBar_OperatorAssign uintptr
	addr_QStyleOptionTitleBar_text_get uintptr
	addr_QStyleOptionTitleBar_text_set uintptr
	addr_QStyleOptionTitleBar_icon_get uintptr
	addr_QStyleOptionTitleBar_icon_set uintptr
	addr_QStyleOptionTitleBar_titleBarState_get uintptr
	addr_QStyleOptionTitleBar_titleBarState_set uintptr
	addr_QStyleOptionTitleBar_titleBarFlags_get uintptr
	addr_QStyleOptionTitleBar_titleBarFlags_set uintptr
	pg_QStyleOptionTitleBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionTitleBar_chain, "QStyleOptionTitleBar_chain")
	symaddr(&addr_QStyleOptionTitleBar_new, "QStyleOptionTitleBar_new")
	symaddr(&addr_QStyleOptionTitleBar_new2, "QStyleOptionTitleBar_new2")
	symaddr(&addr_QStyleOptionTitleBar_OperatorAssign, "QStyleOptionTitleBar_OperatorAssign")
	symaddr(&addr_QStyleOptionTitleBar_text_get, "QStyleOptionTitleBar_text_get")
	symaddr(&addr_QStyleOptionTitleBar_text_set, "QStyleOptionTitleBar_text_set")
	symaddr(&addr_QStyleOptionTitleBar_icon_get, "QStyleOptionTitleBar_icon_get")
	symaddr(&addr_QStyleOptionTitleBar_icon_set, "QStyleOptionTitleBar_icon_set")
	symaddr(&addr_QStyleOptionTitleBar_titleBarState_get, "QStyleOptionTitleBar_titleBarState_get")
	symaddr(&addr_QStyleOptionTitleBar_titleBarState_set, "QStyleOptionTitleBar_titleBarState_set")
	symaddr(&addr_QStyleOptionTitleBar_titleBarFlags_get, "QStyleOptionTitleBar_titleBarFlags_get")
	symaddr(&addr_QStyleOptionTitleBar_titleBarFlags_set, "QStyleOptionTitleBar_titleBarFlags_set")
	reg(&pg_QStyleOptionTitleBar_delete, "QStyleOptionTitleBar_delete")
}

type QStyleOptionTitleBar struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionTitleBar(h uintptr) *QStyleOptionTitleBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionTitleBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionTitleBar{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionTitleBar(h uintptr) *QStyleOptionTitleBar {
	o := bindQStyleOptionTitleBar(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionTitleBar).Delete)
	}
	return o
}

func (this *QStyleOptionTitleBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionTitleBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionTitleBar(h unsafe.Pointer) *QStyleOptionTitleBar { return bindQStyleOptionTitleBar(uintptr(h)) }
func UnsafeNewQStyleOptionTitleBarOwned(h unsafe.Pointer) *QStyleOptionTitleBar { return newownedQStyleOptionTitleBar(uintptr(h)) }

func (this *QStyleOptionTitleBar) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionTitleBar() *QStyleOptionTitleBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_new)
	__o := newownedQStyleOptionTitleBar(__r)
	return __o
}

func NewQStyleOptionTitleBar2(other *QStyleOptionTitleBar) *QStyleOptionTitleBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionTitleBar(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionTitleBar) OperatorAssign(arg1 *QStyleOptionTitleBar) *QStyleOptionTitleBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionTitleBar(__r)
}

func (this *QStyleOptionTitleBar) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionTitleBar) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionTitleBar_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionTitleBar) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionTitleBar) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionTitleBar_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionTitleBar) TitleBarState() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_titleBarState_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionTitleBar) SetTitleBarState(titleBarState int) {
	purego.SyscallN(addr_QStyleOptionTitleBar_titleBarState_set, this.h, uintptr(titleBarState))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTitleBar) TitleBarFlags() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTitleBar_titleBarFlags_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QStyleOptionTitleBar) SetTitleBarFlags(titleBarFlags WindowType) {
	purego.SyscallN(addr_QStyleOptionTitleBar_titleBarFlags_set, this.h, uintptr(titleBarFlags))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTitleBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionTitleBar_delete(this.h)
	}
	this.UnsafeRelease()
}
