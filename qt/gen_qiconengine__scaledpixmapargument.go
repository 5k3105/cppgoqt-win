package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIconEngine__ScaledPixmapArgument_new uintptr
	addr_QIconEngine__ScaledPixmapArgument_OperatorAssign uintptr
	addr_QIconEngine__ScaledPixmapArgument_size_get uintptr
	addr_QIconEngine__ScaledPixmapArgument_size_set uintptr
	addr_QIconEngine__ScaledPixmapArgument_mode_get uintptr
	addr_QIconEngine__ScaledPixmapArgument_mode_set uintptr
	addr_QIconEngine__ScaledPixmapArgument_state_get uintptr
	addr_QIconEngine__ScaledPixmapArgument_state_set uintptr
	pg_QIconEngine__ScaledPixmapArgument_scale_get func(uintptr) float64
	pg_QIconEngine__ScaledPixmapArgument_scale_set func(uintptr, float64)
	addr_QIconEngine__ScaledPixmapArgument_pixmap_get uintptr
	addr_QIconEngine__ScaledPixmapArgument_pixmap_set uintptr
	pg_QIconEngine__ScaledPixmapArgument_delete func(uintptr)
)

func init() {
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_new, "QIconEngine__ScaledPixmapArgument_new")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_OperatorAssign, "QIconEngine__ScaledPixmapArgument_OperatorAssign")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_size_get, "QIconEngine__ScaledPixmapArgument_size_get")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_size_set, "QIconEngine__ScaledPixmapArgument_size_set")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_mode_get, "QIconEngine__ScaledPixmapArgument_mode_get")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_mode_set, "QIconEngine__ScaledPixmapArgument_mode_set")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_state_get, "QIconEngine__ScaledPixmapArgument_state_get")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_state_set, "QIconEngine__ScaledPixmapArgument_state_set")
	reg(&pg_QIconEngine__ScaledPixmapArgument_scale_get, "QIconEngine__ScaledPixmapArgument_scale_get")
	reg(&pg_QIconEngine__ScaledPixmapArgument_scale_set, "QIconEngine__ScaledPixmapArgument_scale_set")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_pixmap_get, "QIconEngine__ScaledPixmapArgument_pixmap_get")
	symaddr(&addr_QIconEngine__ScaledPixmapArgument_pixmap_set, "QIconEngine__ScaledPixmapArgument_pixmap_set")
	reg(&pg_QIconEngine__ScaledPixmapArgument_delete, "QIconEngine__ScaledPixmapArgument_delete")
}

type QIconEngine__ScaledPixmapArgument struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QIconEngine__ScaledPixmapArgument) UnsafeGoID() int64     { return this.vid }
func (this *QIconEngine__ScaledPixmapArgument) UnsafeSetGoID(id int64) { this.vid = id }

func bindQIconEngine__ScaledPixmapArgument(h uintptr) *QIconEngine__ScaledPixmapArgument {
	if h == 0 {
		return nil
	}
	return &QIconEngine__ScaledPixmapArgument{h: h}
}

func newownedQIconEngine__ScaledPixmapArgument(h uintptr) *QIconEngine__ScaledPixmapArgument {
	o := bindQIconEngine__ScaledPixmapArgument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QIconEngine__ScaledPixmapArgument).Delete)
	}
	return o
}

func (this *QIconEngine__ScaledPixmapArgument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIconEngine__ScaledPixmapArgument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIconEngine__ScaledPixmapArgument(h unsafe.Pointer) *QIconEngine__ScaledPixmapArgument { return bindQIconEngine__ScaledPixmapArgument(uintptr(h)) }
func UnsafeNewQIconEngine__ScaledPixmapArgumentOwned(h unsafe.Pointer) *QIconEngine__ScaledPixmapArgument { return newownedQIconEngine__ScaledPixmapArgument(uintptr(h)) }

func (this *QIconEngine__ScaledPixmapArgument) UnsafeRelease() { this.h = 0 }

func NewQIconEngine__ScaledPixmapArgument(arg1 *QIconEngine__ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {
	__r, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQIconEngine__ScaledPixmapArgument(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QIconEngine__ScaledPixmapArgument) OperatorAssign(arg1 *QIconEngine__ScaledPixmapArgument) *QIconEngine__ScaledPixmapArgument {
	__r, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQIconEngine__ScaledPixmapArgument(__r)
}

func (this *QIconEngine__ScaledPixmapArgument) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_size_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QIconEngine__ScaledPixmapArgument) SetSize(size *QSize) {
	purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_size_set, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QIconEngine__ScaledPixmapArgument) Mode() QIcon__Mode {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_mode_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QIcon__Mode(__r)
}

func (this *QIconEngine__ScaledPixmapArgument) SetMode(mode QIcon__Mode) {
	purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_mode_set, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QIconEngine__ScaledPixmapArgument) State() QIcon__State {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_state_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QIcon__State(__r)
}

func (this *QIconEngine__ScaledPixmapArgument) SetState(state QIcon__State) {
	purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_state_set, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QIconEngine__ScaledPixmapArgument) Scale() float64 {
	__r := pg_QIconEngine__ScaledPixmapArgument_scale_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIconEngine__ScaledPixmapArgument) SetScale(scale float64) {
	pg_QIconEngine__ScaledPixmapArgument_scale_set(this.h, scale)
	runtime.KeepAlive(this)
}

func (this *QIconEngine__ScaledPixmapArgument) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_pixmap_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QIconEngine__ScaledPixmapArgument) SetPixmap(pixmap *QPixmap) {
	purego.SyscallN(addr_QIconEngine__ScaledPixmapArgument_pixmap_set, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIconEngine__ScaledPixmapArgument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIconEngine__ScaledPixmapArgument_delete(this.h)
	}
	this.UnsafeRelease()
}
