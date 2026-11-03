package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionTabWidgetFrame_base func(uintptr) uintptr
	addr_QStyleOptionTabWidgetFrame_new uintptr
	addr_QStyleOptionTabWidgetFrame_new2 uintptr
	addr_QStyleOptionTabWidgetFrame_OperatorAssign uintptr
	addr_QStyleOptionTabWidgetFrame_lineWidth_get uintptr
	addr_QStyleOptionTabWidgetFrame_lineWidth_set uintptr
	addr_QStyleOptionTabWidgetFrame_midLineWidth_get uintptr
	addr_QStyleOptionTabWidgetFrame_midLineWidth_set uintptr
	addr_QStyleOptionTabWidgetFrame_shape_get uintptr
	addr_QStyleOptionTabWidgetFrame_shape_set uintptr
	addr_QStyleOptionTabWidgetFrame_tabBarSize_get uintptr
	addr_QStyleOptionTabWidgetFrame_tabBarSize_set uintptr
	addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get uintptr
	addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set uintptr
	addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get uintptr
	addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set uintptr
	addr_QStyleOptionTabWidgetFrame_tabBarRect_get uintptr
	addr_QStyleOptionTabWidgetFrame_tabBarRect_set uintptr
	addr_QStyleOptionTabWidgetFrame_selectedTabRect_get uintptr
	addr_QStyleOptionTabWidgetFrame_selectedTabRect_set uintptr
	pg_QStyleOptionTabWidgetFrame_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionTabWidgetFrame_base, "QStyleOptionTabWidgetFrame_base")
	symaddr(&addr_QStyleOptionTabWidgetFrame_new, "QStyleOptionTabWidgetFrame_new")
	symaddr(&addr_QStyleOptionTabWidgetFrame_new2, "QStyleOptionTabWidgetFrame_new2")
	symaddr(&addr_QStyleOptionTabWidgetFrame_OperatorAssign, "QStyleOptionTabWidgetFrame_OperatorAssign")
	symaddr(&addr_QStyleOptionTabWidgetFrame_lineWidth_get, "QStyleOptionTabWidgetFrame_lineWidth_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_lineWidth_set, "QStyleOptionTabWidgetFrame_lineWidth_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_midLineWidth_get, "QStyleOptionTabWidgetFrame_midLineWidth_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_midLineWidth_set, "QStyleOptionTabWidgetFrame_midLineWidth_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_shape_get, "QStyleOptionTabWidgetFrame_shape_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_shape_set, "QStyleOptionTabWidgetFrame_shape_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_tabBarSize_get, "QStyleOptionTabWidgetFrame_tabBarSize_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_tabBarSize_set, "QStyleOptionTabWidgetFrame_tabBarSize_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get, "QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set, "QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get, "QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set, "QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_tabBarRect_get, "QStyleOptionTabWidgetFrame_tabBarRect_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_tabBarRect_set, "QStyleOptionTabWidgetFrame_tabBarRect_set")
	symaddr(&addr_QStyleOptionTabWidgetFrame_selectedTabRect_get, "QStyleOptionTabWidgetFrame_selectedTabRect_get")
	symaddr(&addr_QStyleOptionTabWidgetFrame_selectedTabRect_set, "QStyleOptionTabWidgetFrame_selectedTabRect_set")
	reg(&pg_QStyleOptionTabWidgetFrame_delete, "QStyleOptionTabWidgetFrame_delete")
}

type QStyleOptionTabWidgetFrame struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionTabWidgetFrame(h uintptr) *QStyleOptionTabWidgetFrame {
	if h == 0 {
		return nil
	}
	return &QStyleOptionTabWidgetFrame{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionTabWidgetFrame_base(h))}
}

func newownedQStyleOptionTabWidgetFrame(h uintptr) *QStyleOptionTabWidgetFrame {
	o := bindQStyleOptionTabWidgetFrame(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionTabWidgetFrame).Delete)
	}
	return o
}

func (this *QStyleOptionTabWidgetFrame) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionTabWidgetFrame) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionTabWidgetFrame(h unsafe.Pointer) *QStyleOptionTabWidgetFrame { return bindQStyleOptionTabWidgetFrame(uintptr(h)) }
func UnsafeNewQStyleOptionTabWidgetFrameOwned(h unsafe.Pointer) *QStyleOptionTabWidgetFrame { return newownedQStyleOptionTabWidgetFrame(uintptr(h)) }

func (this *QStyleOptionTabWidgetFrame) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionTabWidgetFrame() *QStyleOptionTabWidgetFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_new)
	__o := newownedQStyleOptionTabWidgetFrame(__r)
	return __o
}

func NewQStyleOptionTabWidgetFrame2(other *QStyleOptionTabWidgetFrame) *QStyleOptionTabWidgetFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionTabWidgetFrame(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionTabWidgetFrame) OperatorAssign(arg1 *QStyleOptionTabWidgetFrame) *QStyleOptionTabWidgetFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionTabWidgetFrame(__r)
}

func (this *QStyleOptionTabWidgetFrame) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_lineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetLineWidth(lineWidth int) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_lineWidth_set, this.h, uintptr(lineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTabWidgetFrame) MidLineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_midLineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetMidLineWidth(midLineWidth int) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_midLineWidth_set, this.h, uintptr(midLineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTabWidgetFrame) Shape() QTabBar__Shape {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_shape_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabBar__Shape(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetShape(shape QTabBar__Shape) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_shape_set, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTabWidgetFrame) TabBarSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_tabBarSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetTabBarSize(tabBarSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_tabBarSize_set, this.h, uintptr(tabBarSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabBarSize)
}

func (this *QStyleOptionTabWidgetFrame) RightCornerWidgetSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetRightCornerWidgetSize(rightCornerWidgetSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_rightCornerWidgetSize_set, this.h, uintptr(rightCornerWidgetSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rightCornerWidgetSize)
}

func (this *QStyleOptionTabWidgetFrame) LeftCornerWidgetSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetLeftCornerWidgetSize(leftCornerWidgetSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_leftCornerWidgetSize_set, this.h, uintptr(leftCornerWidgetSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(leftCornerWidgetSize)
}

func (this *QStyleOptionTabWidgetFrame) TabBarRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_tabBarRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetTabBarRect(tabBarRect *QRect) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_tabBarRect_set, this.h, uintptr(tabBarRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabBarRect)
}

func (this *QStyleOptionTabWidgetFrame) SelectedTabRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTabWidgetFrame_selectedTabRect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOptionTabWidgetFrame) SetSelectedTabRect(selectedTabRect *QRect) {
	purego.SyscallN(addr_QStyleOptionTabWidgetFrame_selectedTabRect_set, this.h, uintptr(selectedTabRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selectedTabRect)
}

func (this *QStyleOptionTabWidgetFrame) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionTabWidgetFrame_delete(this.h)
	}
	this.UnsafeRelease()
}
