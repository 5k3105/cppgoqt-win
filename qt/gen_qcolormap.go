package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColormap_new uintptr
	addr_QColormap_initialize uintptr
	addr_QColormap_cleanup uintptr
	addr_QColormap_instance uintptr
	addr_QColormap_OperatorAssign uintptr
	addr_QColormap_mode uintptr
	addr_QColormap_depth uintptr
	addr_QColormap_size uintptr
	addr_QColormap_pixel uintptr
	addr_QColormap_colorAt uintptr
	addr_QColormap_colormap uintptr
	addr_QColormap_instance2 uintptr
	pg_QColormap_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColormap_new, "QColormap_new")
	symaddr(&addr_QColormap_initialize, "QColormap_initialize")
	symaddr(&addr_QColormap_cleanup, "QColormap_cleanup")
	symaddr(&addr_QColormap_instance, "QColormap_instance")
	symaddr(&addr_QColormap_OperatorAssign, "QColormap_OperatorAssign")
	symaddr(&addr_QColormap_mode, "QColormap_mode")
	symaddr(&addr_QColormap_depth, "QColormap_depth")
	symaddr(&addr_QColormap_size, "QColormap_size")
	symaddr(&addr_QColormap_pixel, "QColormap_pixel")
	symaddr(&addr_QColormap_colorAt, "QColormap_colorAt")
	symaddr(&addr_QColormap_colormap, "QColormap_colormap")
	symaddr(&addr_QColormap_instance2, "QColormap_instance2")
	reg(&pg_QColormap_delete, "QColormap_delete")
}

type QColormap struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QColormap) UnsafeGoID() int64     { return this.vid }
func (this *QColormap) UnsafeSetGoID(id int64) { this.vid = id }

func bindQColormap(h uintptr) *QColormap {
	if h == 0 {
		return nil
	}
	return &QColormap{h: h}
}

func newownedQColormap(h uintptr) *QColormap {
	o := bindQColormap(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QColormap).Delete)
	}
	return o
}

func (this *QColormap) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColormap) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColormap(h unsafe.Pointer) *QColormap { return bindQColormap(uintptr(h)) }
func UnsafeNewQColormapOwned(h unsafe.Pointer) *QColormap { return newownedQColormap(uintptr(h)) }

func (this *QColormap) UnsafeRelease() { this.h = 0 }

func NewQColormap(colormap *QColormap) *QColormap {
	__r, _, _ := purego.SyscallN(addr_QColormap_new, uintptr(colormap.UnsafePointer()))
	__o := newownedQColormap(__r)
	runtime.KeepAlive(colormap)
	return __o
}

func QColormap_Initialize() {
	purego.SyscallN(addr_QColormap_initialize)
}

func QColormap_Cleanup() {
	purego.SyscallN(addr_QColormap_cleanup)
}

func QColormap_Instance() *QColormap {
	__r, _, _ := purego.SyscallN(addr_QColormap_instance)
	return newownedQColormap(__r)
}

func (this *QColormap) OperatorAssign(colormap *QColormap) *QColormap {
	__r, _, _ := purego.SyscallN(addr_QColormap_OperatorAssign, this.h, uintptr(colormap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colormap)
	return bindQColormap(__r)
}

func (this *QColormap) Mode() QColormap__Mode {
	__r0, _, _ := purego.SyscallN(addr_QColormap_mode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColormap__Mode(__r)
}

func (this *QColormap) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QColormap_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColormap) Size() int {
	__r0, _, _ := purego.SyscallN(addr_QColormap_size, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QColormap) Pixel(color *QColor) uint {
	__r0, _, _ := purego.SyscallN(addr_QColormap_pixel, this.h, uintptr(color.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
	return uint(__r)
}

func (this *QColormap) ColorAt(pixel uint) *QColor {
	__r, _, _ := purego.SyscallN(addr_QColormap_colorAt, this.h, uintptr(pixel))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QColormap) Colormap() []*QColor {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColormap_colormap, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QColor, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQColor(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QColormap_InstanceWithScreen(screen int) *QColormap {
	__r, _, _ := purego.SyscallN(addr_QColormap_instance2, uintptr(screen))
	return newownedQColormap(__r)
}

func (this *QColormap) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColormap_delete(this.h)
	}
	this.UnsafeRelease()
}
