package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSurface_surfaceClass uintptr
	addr_QSurface_format uintptr
	addr_QSurface_surfaceType uintptr
	addr_QSurface_supportsOpenGL uintptr
	addr_QSurface_size uintptr
	addr_QSurface_OperatorAssign uintptr
	pg_QSurface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSurface_surfaceClass, "QSurface_surfaceClass")
	symaddr(&addr_QSurface_format, "QSurface_format")
	symaddr(&addr_QSurface_surfaceType, "QSurface_surfaceType")
	symaddr(&addr_QSurface_supportsOpenGL, "QSurface_supportsOpenGL")
	symaddr(&addr_QSurface_size, "QSurface_size")
	symaddr(&addr_QSurface_OperatorAssign, "QSurface_OperatorAssign")
	reg(&pg_QSurface_delete, "QSurface_delete")
}

type QSurface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSurface) UnsafeGoID() int64     { return this.vid }
func (this *QSurface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSurface(h uintptr) *QSurface {
	if h == 0 {
		return nil
	}
	return &QSurface{h: h}
}

func newownedQSurface(h uintptr) *QSurface {
	o := bindQSurface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSurface).Delete)
	}
	return o
}

func (this *QSurface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSurface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSurface(h unsafe.Pointer) *QSurface { return bindQSurface(uintptr(h)) }
func UnsafeNewQSurfaceOwned(h unsafe.Pointer) *QSurface { return newownedQSurface(uintptr(h)) }

func (this *QSurface) UnsafeRelease() { this.h = 0 }

func (this *QSurface) SurfaceClass() QSurface__SurfaceClass {
	__r0, _, _ := purego.SyscallN(addr_QSurface_surfaceClass, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceClass(__r)
}

func (this *QSurface) Format() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QSurface_format, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QSurface) SurfaceType() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QSurface_surfaceType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QSurface) SupportsOpenGL() bool {
	__r0, _, _ := purego.SyscallN(addr_QSurface_supportsOpenGL, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSurface) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSurface_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSurface) OperatorAssign(arg1 *QSurface) *QSurface {
	__r, _, _ := purego.SyscallN(addr_QSurface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSurface(__r)
}

func (this *QSurface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSurface_delete(this.h)
	}
	this.UnsafeRelease()
}
