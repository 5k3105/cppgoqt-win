package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColorTransform_new uintptr
	addr_QColorTransform_new2 uintptr
	addr_QColorTransform_OperatorAssign uintptr
	addr_QColorTransform_swap uintptr
	addr_QColorTransform_isIdentity uintptr
	addr_QColorTransform_map uintptr
	addr_QColorTransform_map2 uintptr
	addr_QColorTransform_map3 uintptr
	addr_QColorTransform_OperatorEqual uintptr
	addr_QColorTransform_OperatorNotEqual uintptr
	pg_QColorTransform_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColorTransform_new, "QColorTransform_new")
	symaddr(&addr_QColorTransform_new2, "QColorTransform_new2")
	symaddr(&addr_QColorTransform_OperatorAssign, "QColorTransform_OperatorAssign")
	symaddr(&addr_QColorTransform_swap, "QColorTransform_swap")
	symaddr(&addr_QColorTransform_isIdentity, "QColorTransform_isIdentity")
	symaddr(&addr_QColorTransform_map, "QColorTransform_map")
	symaddr(&addr_QColorTransform_map2, "QColorTransform_map2")
	symaddr(&addr_QColorTransform_map3, "QColorTransform_map3")
	symaddr(&addr_QColorTransform_OperatorEqual, "QColorTransform_OperatorEqual")
	symaddr(&addr_QColorTransform_OperatorNotEqual, "QColorTransform_OperatorNotEqual")
	reg(&pg_QColorTransform_delete, "QColorTransform_delete")
}

type QColorTransform struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QColorTransform) UnsafeGoID() int64     { return this.vid }
func (this *QColorTransform) UnsafeSetGoID(id int64) { this.vid = id }

func bindQColorTransform(h uintptr) *QColorTransform {
	if h == 0 {
		return nil
	}
	return &QColorTransform{h: h}
}

func newownedQColorTransform(h uintptr) *QColorTransform {
	o := bindQColorTransform(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QColorTransform).Delete)
	}
	return o
}

func (this *QColorTransform) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColorTransform) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColorTransform(h unsafe.Pointer) *QColorTransform { return bindQColorTransform(uintptr(h)) }
func UnsafeNewQColorTransformOwned(h unsafe.Pointer) *QColorTransform { return newownedQColorTransform(uintptr(h)) }

func (this *QColorTransform) UnsafeRelease() { this.h = 0 }

func NewQColorTransform() *QColorTransform {
	__r, _, _ := purego.SyscallN(addr_QColorTransform_new)
	__o := newownedQColorTransform(__r)
	return __o
}

func NewQColorTransform2(colorTransform *QColorTransform) *QColorTransform {
	__r, _, _ := purego.SyscallN(addr_QColorTransform_new2, uintptr(colorTransform.UnsafePointer()))
	__o := newownedQColorTransform(__r)
	runtime.KeepAlive(colorTransform)
	return __o
}

func (this *QColorTransform) OperatorAssign(other *QColorTransform) *QColorTransform {
	__r, _, _ := purego.SyscallN(addr_QColorTransform_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQColorTransform(__r)
}

func (this *QColorTransform) Swap(other *QColorTransform) {
	purego.SyscallN(addr_QColorTransform_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QColorTransform) IsIdentity() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorTransform_isIdentity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorTransform) Map(argb uint) uint {
	__r0, _, _ := purego.SyscallN(addr_QColorTransform_map, this.h, uintptr(argb))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QColorTransform) MapWithRgba64(rgba64 *QRgba64) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QColorTransform_map2, this.h, uintptr(rgba64.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rgba64)
	return newownedQRgba64(__r)
}

func (this *QColorTransform) MapWithColor(color *QColor) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColorTransform_map3, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
	return newownedQColor(__r)
}

func (this *QColorTransform) OperatorEqual(ct2 *QColorTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorTransform_OperatorEqual, this.h, uintptr(ct2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ct2)
	return __r != 0
}

func (this *QColorTransform) OperatorNotEqual(ct2 *QColorTransform) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorTransform_OperatorNotEqual, this.h, uintptr(ct2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ct2)
	return __r != 0
}

func (this *QColorTransform) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColorTransform_delete(this.h)
	}
	this.UnsafeRelease()
}
