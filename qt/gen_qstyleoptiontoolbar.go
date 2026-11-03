package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionToolBar_base func(uintptr) uintptr
	addr_QStyleOptionToolBar_new uintptr
	addr_QStyleOptionToolBar_new2 uintptr
	addr_QStyleOptionToolBar_OperatorAssign uintptr
	addr_QStyleOptionToolBar_positionOfLine_get uintptr
	addr_QStyleOptionToolBar_positionOfLine_set uintptr
	addr_QStyleOptionToolBar_positionWithinLine_get uintptr
	addr_QStyleOptionToolBar_positionWithinLine_set uintptr
	addr_QStyleOptionToolBar_toolBarArea_get uintptr
	addr_QStyleOptionToolBar_toolBarArea_set uintptr
	addr_QStyleOptionToolBar_features_get uintptr
	addr_QStyleOptionToolBar_features_set uintptr
	addr_QStyleOptionToolBar_lineWidth_get uintptr
	addr_QStyleOptionToolBar_lineWidth_set uintptr
	addr_QStyleOptionToolBar_midLineWidth_get uintptr
	addr_QStyleOptionToolBar_midLineWidth_set uintptr
	pg_QStyleOptionToolBar_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionToolBar_base, "QStyleOptionToolBar_base")
	symaddr(&addr_QStyleOptionToolBar_new, "QStyleOptionToolBar_new")
	symaddr(&addr_QStyleOptionToolBar_new2, "QStyleOptionToolBar_new2")
	symaddr(&addr_QStyleOptionToolBar_OperatorAssign, "QStyleOptionToolBar_OperatorAssign")
	symaddr(&addr_QStyleOptionToolBar_positionOfLine_get, "QStyleOptionToolBar_positionOfLine_get")
	symaddr(&addr_QStyleOptionToolBar_positionOfLine_set, "QStyleOptionToolBar_positionOfLine_set")
	symaddr(&addr_QStyleOptionToolBar_positionWithinLine_get, "QStyleOptionToolBar_positionWithinLine_get")
	symaddr(&addr_QStyleOptionToolBar_positionWithinLine_set, "QStyleOptionToolBar_positionWithinLine_set")
	symaddr(&addr_QStyleOptionToolBar_toolBarArea_get, "QStyleOptionToolBar_toolBarArea_get")
	symaddr(&addr_QStyleOptionToolBar_toolBarArea_set, "QStyleOptionToolBar_toolBarArea_set")
	symaddr(&addr_QStyleOptionToolBar_features_get, "QStyleOptionToolBar_features_get")
	symaddr(&addr_QStyleOptionToolBar_features_set, "QStyleOptionToolBar_features_set")
	symaddr(&addr_QStyleOptionToolBar_lineWidth_get, "QStyleOptionToolBar_lineWidth_get")
	symaddr(&addr_QStyleOptionToolBar_lineWidth_set, "QStyleOptionToolBar_lineWidth_set")
	symaddr(&addr_QStyleOptionToolBar_midLineWidth_get, "QStyleOptionToolBar_midLineWidth_get")
	symaddr(&addr_QStyleOptionToolBar_midLineWidth_set, "QStyleOptionToolBar_midLineWidth_set")
	reg(&pg_QStyleOptionToolBar_delete, "QStyleOptionToolBar_delete")
}

type QStyleOptionToolBar struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionToolBar(h uintptr) *QStyleOptionToolBar {
	if h == 0 {
		return nil
	}
	return &QStyleOptionToolBar{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionToolBar_base(h))}
}

func newownedQStyleOptionToolBar(h uintptr) *QStyleOptionToolBar {
	o := bindQStyleOptionToolBar(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionToolBar).Delete)
	}
	return o
}

func (this *QStyleOptionToolBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionToolBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionToolBar(h unsafe.Pointer) *QStyleOptionToolBar { return bindQStyleOptionToolBar(uintptr(h)) }
func UnsafeNewQStyleOptionToolBarOwned(h unsafe.Pointer) *QStyleOptionToolBar { return newownedQStyleOptionToolBar(uintptr(h)) }

func (this *QStyleOptionToolBar) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionToolBar() *QStyleOptionToolBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_new)
	__o := newownedQStyleOptionToolBar(__r)
	return __o
}

func NewQStyleOptionToolBar2(other *QStyleOptionToolBar) *QStyleOptionToolBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionToolBar(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionToolBar) OperatorAssign(arg1 *QStyleOptionToolBar) *QStyleOptionToolBar {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionToolBar(__r)
}

func (this *QStyleOptionToolBar) PositionOfLine() QStyleOptionToolBar__ToolBarPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_positionOfLine_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolBar__ToolBarPosition(__r)
}

func (this *QStyleOptionToolBar) SetPositionOfLine(positionOfLine QStyleOptionToolBar__ToolBarPosition) {
	purego.SyscallN(addr_QStyleOptionToolBar_positionOfLine_set, this.h, uintptr(positionOfLine))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) PositionWithinLine() QStyleOptionToolBar__ToolBarPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_positionWithinLine_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolBar__ToolBarPosition(__r)
}

func (this *QStyleOptionToolBar) SetPositionWithinLine(positionWithinLine QStyleOptionToolBar__ToolBarPosition) {
	purego.SyscallN(addr_QStyleOptionToolBar_positionWithinLine_set, this.h, uintptr(positionWithinLine))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) ToolBarArea() ToolBarArea {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_toolBarArea_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolBarArea(__r)
}

func (this *QStyleOptionToolBar) SetToolBarArea(toolBarArea ToolBarArea) {
	purego.SyscallN(addr_QStyleOptionToolBar_toolBarArea_set, this.h, uintptr(toolBarArea))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) Features() QStyleOptionToolBar__ToolBarFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionToolBar__ToolBarFeature(__r)
}

func (this *QStyleOptionToolBar) SetFeatures(features QStyleOptionToolBar__ToolBarFeature) {
	purego.SyscallN(addr_QStyleOptionToolBar_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_lineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionToolBar) SetLineWidth(lineWidth int) {
	purego.SyscallN(addr_QStyleOptionToolBar_lineWidth_set, this.h, uintptr(lineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) MidLineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionToolBar_midLineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionToolBar) SetMidLineWidth(midLineWidth int) {
	purego.SyscallN(addr_QStyleOptionToolBar_midLineWidth_set, this.h, uintptr(midLineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionToolBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionToolBar_delete(this.h)
	}
	this.UnsafeRelease()
}
