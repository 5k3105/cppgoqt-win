package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColorSpace__PrimaryPoints_fromPrimaries uintptr
	addr_QColorSpace__PrimaryPoints_isValid uintptr
	addr_QColorSpace__PrimaryPoints_whitePoint_get uintptr
	addr_QColorSpace__PrimaryPoints_whitePoint_set uintptr
	addr_QColorSpace__PrimaryPoints_redPoint_get uintptr
	addr_QColorSpace__PrimaryPoints_redPoint_set uintptr
	addr_QColorSpace__PrimaryPoints_greenPoint_get uintptr
	addr_QColorSpace__PrimaryPoints_greenPoint_set uintptr
	addr_QColorSpace__PrimaryPoints_bluePoint_get uintptr
	addr_QColorSpace__PrimaryPoints_bluePoint_set uintptr
	pg_QColorSpace__PrimaryPoints_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColorSpace__PrimaryPoints_fromPrimaries, "QColorSpace__PrimaryPoints_fromPrimaries")
	symaddr(&addr_QColorSpace__PrimaryPoints_isValid, "QColorSpace__PrimaryPoints_isValid")
	symaddr(&addr_QColorSpace__PrimaryPoints_whitePoint_get, "QColorSpace__PrimaryPoints_whitePoint_get")
	symaddr(&addr_QColorSpace__PrimaryPoints_whitePoint_set, "QColorSpace__PrimaryPoints_whitePoint_set")
	symaddr(&addr_QColorSpace__PrimaryPoints_redPoint_get, "QColorSpace__PrimaryPoints_redPoint_get")
	symaddr(&addr_QColorSpace__PrimaryPoints_redPoint_set, "QColorSpace__PrimaryPoints_redPoint_set")
	symaddr(&addr_QColorSpace__PrimaryPoints_greenPoint_get, "QColorSpace__PrimaryPoints_greenPoint_get")
	symaddr(&addr_QColorSpace__PrimaryPoints_greenPoint_set, "QColorSpace__PrimaryPoints_greenPoint_set")
	symaddr(&addr_QColorSpace__PrimaryPoints_bluePoint_get, "QColorSpace__PrimaryPoints_bluePoint_get")
	symaddr(&addr_QColorSpace__PrimaryPoints_bluePoint_set, "QColorSpace__PrimaryPoints_bluePoint_set")
	reg(&pg_QColorSpace__PrimaryPoints_delete, "QColorSpace__PrimaryPoints_delete")
}

type QColorSpace__PrimaryPoints struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QColorSpace__PrimaryPoints) UnsafeGoID() int64     { return this.vid }
func (this *QColorSpace__PrimaryPoints) UnsafeSetGoID(id int64) { this.vid = id }

func bindQColorSpace__PrimaryPoints(h uintptr) *QColorSpace__PrimaryPoints {
	if h == 0 {
		return nil
	}
	return &QColorSpace__PrimaryPoints{h: h}
}

func newownedQColorSpace__PrimaryPoints(h uintptr) *QColorSpace__PrimaryPoints {
	o := bindQColorSpace__PrimaryPoints(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QColorSpace__PrimaryPoints).Delete)
	}
	return o
}

func (this *QColorSpace__PrimaryPoints) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColorSpace__PrimaryPoints) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColorSpace__PrimaryPoints(h unsafe.Pointer) *QColorSpace__PrimaryPoints { return bindQColorSpace__PrimaryPoints(uintptr(h)) }
func UnsafeNewQColorSpace__PrimaryPointsOwned(h unsafe.Pointer) *QColorSpace__PrimaryPoints { return newownedQColorSpace__PrimaryPoints(uintptr(h)) }

func (this *QColorSpace__PrimaryPoints) UnsafeRelease() { this.h = 0 }

func QColorSpace__PrimaryPoints_FromPrimaries(primaries QColorSpace__Primaries) *QColorSpace__PrimaryPoints {
	__r, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_fromPrimaries, uintptr(primaries))
	return newownedQColorSpace__PrimaryPoints(__r)
}

func (this *QColorSpace__PrimaryPoints) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorSpace__PrimaryPoints) WhitePoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_whitePoint_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QColorSpace__PrimaryPoints) SetWhitePoint(whitePoint *QPointF) {
	purego.SyscallN(addr_QColorSpace__PrimaryPoints_whitePoint_set, this.h, uintptr(whitePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whitePoint)
}

func (this *QColorSpace__PrimaryPoints) RedPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_redPoint_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QColorSpace__PrimaryPoints) SetRedPoint(redPoint *QPointF) {
	purego.SyscallN(addr_QColorSpace__PrimaryPoints_redPoint_set, this.h, uintptr(redPoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(redPoint)
}

func (this *QColorSpace__PrimaryPoints) GreenPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_greenPoint_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QColorSpace__PrimaryPoints) SetGreenPoint(greenPoint *QPointF) {
	purego.SyscallN(addr_QColorSpace__PrimaryPoints_greenPoint_set, this.h, uintptr(greenPoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(greenPoint)
}

func (this *QColorSpace__PrimaryPoints) BluePoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QColorSpace__PrimaryPoints_bluePoint_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QColorSpace__PrimaryPoints) SetBluePoint(bluePoint *QPointF) {
	purego.SyscallN(addr_QColorSpace__PrimaryPoints_bluePoint_set, this.h, uintptr(bluePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bluePoint)
}

func (this *QColorSpace__PrimaryPoints) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColorSpace__PrimaryPoints_delete(this.h)
	}
	this.UnsafeRelease()
}
