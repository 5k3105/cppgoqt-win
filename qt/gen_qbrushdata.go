package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBrushData_ref_get uintptr
	addr_QBrushData_ref_set uintptr
	addr_QBrushData_style_get uintptr
	addr_QBrushData_style_set uintptr
	addr_QBrushData_color_get uintptr
	addr_QBrushData_color_set uintptr
	addr_QBrushData_transform_get uintptr
	addr_QBrushData_transform_set uintptr
)

func init() {
	symaddr(&addr_QBrushData_ref_get, "QBrushData_ref_get")
	symaddr(&addr_QBrushData_ref_set, "QBrushData_ref_set")
	symaddr(&addr_QBrushData_style_get, "QBrushData_style_get")
	symaddr(&addr_QBrushData_style_set, "QBrushData_style_set")
	symaddr(&addr_QBrushData_color_get, "QBrushData_color_get")
	symaddr(&addr_QBrushData_color_set, "QBrushData_color_set")
	symaddr(&addr_QBrushData_transform_get, "QBrushData_transform_get")
	symaddr(&addr_QBrushData_transform_set, "QBrushData_transform_set")
}

type QBrushData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBrushData) UnsafeGoID() int64     { return this.vid }
func (this *QBrushData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBrushData(h uintptr) *QBrushData {
	if h == 0 {
		return nil
	}
	return &QBrushData{h: h}
}

func newownedQBrushData(h uintptr) *QBrushData { return bindQBrushData(h) }

func (this *QBrushData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBrushData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBrushData(h unsafe.Pointer) *QBrushData { return bindQBrushData(uintptr(h)) }
func UnsafeNewQBrushDataOwned(h unsafe.Pointer) *QBrushData { return newownedQBrushData(uintptr(h)) }

func (this *QBrushData) UnsafeRelease() { this.h = 0 }

func (this *QBrushData) Ref() *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QBrushData_ref_get, this.h)
	runtime.KeepAlive(this)
	return newownedQAtomicInt(__r)
}

func (this *QBrushData) SetRef(ref *QAtomicInt) {
	purego.SyscallN(addr_QBrushData_ref_set, this.h, uintptr(ref.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ref)
}

func (this *QBrushData) Style() BrushStyle {
	__r0, _, _ := purego.SyscallN(addr_QBrushData_style_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return BrushStyle(__r)
}

func (this *QBrushData) SetStyle(style BrushStyle) {
	purego.SyscallN(addr_QBrushData_style_set, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QBrushData) Color() *QColor {
	__r, _, _ := purego.SyscallN(addr_QBrushData_color_get, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QBrushData) SetColor(color *QColor) {
	purego.SyscallN(addr_QBrushData_color_set, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QBrushData) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QBrushData_transform_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QBrushData) SetTransform(transform *QTransform) {
	purego.SyscallN(addr_QBrushData_transform_set, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QBrushData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
