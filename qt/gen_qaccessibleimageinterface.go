package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleImageInterface_imageDescription uintptr
	addr_QAccessibleImageInterface_imageSize uintptr
	addr_QAccessibleImageInterface_imagePosition uintptr
	addr_QAccessibleImageInterface_OperatorAssign uintptr
	pg_QAccessibleImageInterface__set_override func(uintptr, int32)
	pg_QAccessibleImageInterface__Go_new func(int64) uintptr
	pg_QAccessibleImageInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleImageInterface_imageDescription, "QAccessibleImageInterface_imageDescription")
	symaddr(&addr_QAccessibleImageInterface_imageSize, "QAccessibleImageInterface_imageSize")
	symaddr(&addr_QAccessibleImageInterface_imagePosition, "QAccessibleImageInterface_imagePosition")
	symaddr(&addr_QAccessibleImageInterface_OperatorAssign, "QAccessibleImageInterface_OperatorAssign")
	reg(&pg_QAccessibleImageInterface__set_override, "QAccessibleImageInterface__set_override")
	reg(&pg_QAccessibleImageInterface__Go_new, "QAccessibleImageInterface__Go_new")
	reg(&pg_QAccessibleImageInterface_delete, "QAccessibleImageInterface_delete")
}

type QAccessibleImageInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleImageInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleImageInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleImageInterface(h uintptr) *QAccessibleImageInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleImageInterface{h: h}
}

func newownedQAccessibleImageInterface(h uintptr) *QAccessibleImageInterface {
	o := bindQAccessibleImageInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleImageInterface).Delete)
	}
	return o
}

func (this *QAccessibleImageInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleImageInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleImageInterface(h unsafe.Pointer) *QAccessibleImageInterface { return bindQAccessibleImageInterface(uintptr(h)) }
func UnsafeNewQAccessibleImageInterfaceOwned(h unsafe.Pointer) *QAccessibleImageInterface { return newownedQAccessibleImageInterface(uintptr(h)) }

func (this *QAccessibleImageInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleImageInterface) ImageDescription() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleImageInterface_imageDescription, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleImageInterface) ImageSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAccessibleImageInterface_imageSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAccessibleImageInterface) ImagePosition() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QAccessibleImageInterface_imagePosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QAccessibleImageInterface) OperatorAssign(arg1 *QAccessibleImageInterface) *QAccessibleImageInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleImageInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleImageInterface(__r)
}

func QAccessibleImageInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 503321968110705180:
		return 0
	case -5595411163605219293:
		return 1
	case -7694988315546962491:
		return 2
	}
	return -1
}

func QAccessibleImageInterface_enable(h uintptr, idx int32) { pg_QAccessibleImageInterface__set_override(h, idx) }

func NewQAccessibleImageInterface() *QAccessibleImageInterface {
	id := nextVirtID()
	o := newownedQAccessibleImageInterface(pg_QAccessibleImageInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleImageInterface_virtIdx, QAccessibleImageInterface_enable)
	return o
}

func (this *QAccessibleImageInterface) OnImageDescription(fn func() string) {
	registerVirt(this.UnsafeGoID(), 503321968110705180, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "imageDescription"
}

func (this *QAccessibleImageInterface) OnImageSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -5595411163605219293, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "imageSize"
}

func (this *QAccessibleImageInterface) OnImagePosition(fn func() *QPoint) {
	registerVirt(this.UnsafeGoID(), -7694988315546962491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "imagePosition"
}

func (this *QAccessibleImageInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleImageInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
