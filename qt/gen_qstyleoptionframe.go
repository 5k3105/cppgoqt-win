package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionFrame_base func(uintptr) uintptr
	addr_QStyleOptionFrame_new uintptr
	addr_QStyleOptionFrame_new2 uintptr
	addr_QStyleOptionFrame_OperatorAssign uintptr
	addr_QStyleOptionFrame_lineWidth_get uintptr
	addr_QStyleOptionFrame_lineWidth_set uintptr
	addr_QStyleOptionFrame_midLineWidth_get uintptr
	addr_QStyleOptionFrame_midLineWidth_set uintptr
	addr_QStyleOptionFrame_features_get uintptr
	addr_QStyleOptionFrame_features_set uintptr
	addr_QStyleOptionFrame_frameShape_get uintptr
	addr_QStyleOptionFrame_frameShape_set uintptr
	pg_QStyleOptionFrame_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionFrame_base, "QStyleOptionFrame_base")
	symaddr(&addr_QStyleOptionFrame_new, "QStyleOptionFrame_new")
	symaddr(&addr_QStyleOptionFrame_new2, "QStyleOptionFrame_new2")
	symaddr(&addr_QStyleOptionFrame_OperatorAssign, "QStyleOptionFrame_OperatorAssign")
	symaddr(&addr_QStyleOptionFrame_lineWidth_get, "QStyleOptionFrame_lineWidth_get")
	symaddr(&addr_QStyleOptionFrame_lineWidth_set, "QStyleOptionFrame_lineWidth_set")
	symaddr(&addr_QStyleOptionFrame_midLineWidth_get, "QStyleOptionFrame_midLineWidth_get")
	symaddr(&addr_QStyleOptionFrame_midLineWidth_set, "QStyleOptionFrame_midLineWidth_set")
	symaddr(&addr_QStyleOptionFrame_features_get, "QStyleOptionFrame_features_get")
	symaddr(&addr_QStyleOptionFrame_features_set, "QStyleOptionFrame_features_set")
	symaddr(&addr_QStyleOptionFrame_frameShape_get, "QStyleOptionFrame_frameShape_get")
	symaddr(&addr_QStyleOptionFrame_frameShape_set, "QStyleOptionFrame_frameShape_set")
	reg(&pg_QStyleOptionFrame_delete, "QStyleOptionFrame_delete")
}

type QStyleOptionFrame struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionFrame(h uintptr) *QStyleOptionFrame {
	if h == 0 {
		return nil
	}
	return &QStyleOptionFrame{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionFrame_base(h))}
}

func newownedQStyleOptionFrame(h uintptr) *QStyleOptionFrame {
	o := bindQStyleOptionFrame(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionFrame).Delete)
	}
	return o
}

func (this *QStyleOptionFrame) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionFrame) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionFrame(h unsafe.Pointer) *QStyleOptionFrame { return bindQStyleOptionFrame(uintptr(h)) }
func UnsafeNewQStyleOptionFrameOwned(h unsafe.Pointer) *QStyleOptionFrame { return newownedQStyleOptionFrame(uintptr(h)) }

func (this *QStyleOptionFrame) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionFrame() *QStyleOptionFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFrame_new)
	__o := newownedQStyleOptionFrame(__r)
	return __o
}

func NewQStyleOptionFrame2(other *QStyleOptionFrame) *QStyleOptionFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFrame_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionFrame(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionFrame) OperatorAssign(arg1 *QStyleOptionFrame) *QStyleOptionFrame {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionFrame_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionFrame(__r)
}

func (this *QStyleOptionFrame) LineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionFrame_lineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionFrame) SetLineWidth(lineWidth int) {
	purego.SyscallN(addr_QStyleOptionFrame_lineWidth_set, this.h, uintptr(lineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionFrame) MidLineWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionFrame_midLineWidth_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionFrame) SetMidLineWidth(midLineWidth int) {
	purego.SyscallN(addr_QStyleOptionFrame_midLineWidth_set, this.h, uintptr(midLineWidth))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionFrame) Features() QStyleOptionFrame__FrameFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionFrame_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionFrame__FrameFeature(__r)
}

func (this *QStyleOptionFrame) SetFeatures(features QStyleOptionFrame__FrameFeature) {
	purego.SyscallN(addr_QStyleOptionFrame_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionFrame) FrameShape() QFrame__Shape {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionFrame_frameShape_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFrame__Shape(__r)
}

func (this *QStyleOptionFrame) SetFrameShape(frameShape QFrame__Shape) {
	purego.SyscallN(addr_QStyleOptionFrame_frameShape_set, this.h, uintptr(frameShape))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionFrame) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionFrame_delete(this.h)
	}
	this.UnsafeRelease()
}
