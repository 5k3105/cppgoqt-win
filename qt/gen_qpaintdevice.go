package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPaintDevice_devType uintptr
	addr_QPaintDevice_paintingActive uintptr
	addr_QPaintDevice_paintEngine uintptr
	addr_QPaintDevice_width uintptr
	addr_QPaintDevice_height uintptr
	addr_QPaintDevice_widthMM uintptr
	addr_QPaintDevice_heightMM uintptr
	addr_QPaintDevice_logicalDpiX uintptr
	addr_QPaintDevice_logicalDpiY uintptr
	addr_QPaintDevice_physicalDpiX uintptr
	addr_QPaintDevice_physicalDpiY uintptr
	pg_QPaintDevice_devicePixelRatio func(uintptr) float64
	pg_QPaintDevice_devicePixelRatioF func(uintptr) float64
	addr_QPaintDevice_colorCount uintptr
	addr_QPaintDevice_depth uintptr
	pg_QPaintDevice_devicePixelRatioFScale func() float64
	pg_QPaintDevice_encodeMetricF func(int32, float64) int32
	pg_QPaintDevice_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPaintDevice_devType, "QPaintDevice_devType")
	symaddr(&addr_QPaintDevice_paintingActive, "QPaintDevice_paintingActive")
	symaddr(&addr_QPaintDevice_paintEngine, "QPaintDevice_paintEngine")
	symaddr(&addr_QPaintDevice_width, "QPaintDevice_width")
	symaddr(&addr_QPaintDevice_height, "QPaintDevice_height")
	symaddr(&addr_QPaintDevice_widthMM, "QPaintDevice_widthMM")
	symaddr(&addr_QPaintDevice_heightMM, "QPaintDevice_heightMM")
	symaddr(&addr_QPaintDevice_logicalDpiX, "QPaintDevice_logicalDpiX")
	symaddr(&addr_QPaintDevice_logicalDpiY, "QPaintDevice_logicalDpiY")
	symaddr(&addr_QPaintDevice_physicalDpiX, "QPaintDevice_physicalDpiX")
	symaddr(&addr_QPaintDevice_physicalDpiY, "QPaintDevice_physicalDpiY")
	reg(&pg_QPaintDevice_devicePixelRatio, "QPaintDevice_devicePixelRatio")
	reg(&pg_QPaintDevice_devicePixelRatioF, "QPaintDevice_devicePixelRatioF")
	symaddr(&addr_QPaintDevice_colorCount, "QPaintDevice_colorCount")
	symaddr(&addr_QPaintDevice_depth, "QPaintDevice_depth")
	reg(&pg_QPaintDevice_devicePixelRatioFScale, "QPaintDevice_devicePixelRatioFScale")
	reg(&pg_QPaintDevice_encodeMetricF, "QPaintDevice_encodeMetricF")
	reg(&pg_QPaintDevice_delete, "QPaintDevice_delete")
}

type QPaintDevice struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPaintDevice) UnsafeGoID() int64     { return this.vid }
func (this *QPaintDevice) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPaintDevice(h uintptr) *QPaintDevice {
	if h == 0 {
		return nil
	}
	return &QPaintDevice{h: h}
}

func newownedQPaintDevice(h uintptr) *QPaintDevice {
	o := bindQPaintDevice(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPaintDevice).Delete)
	}
	return o
}

func (this *QPaintDevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPaintDevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPaintDevice(h unsafe.Pointer) *QPaintDevice { return bindQPaintDevice(uintptr(h)) }
func UnsafeNewQPaintDeviceOwned(h unsafe.Pointer) *QPaintDevice { return newownedQPaintDevice(uintptr(h)) }

func (this *QPaintDevice) UnsafeRelease() { this.h = 0 }

func (this *QPaintDevice) DevType() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_devType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) PaintingActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_paintingActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintDevice) PaintEngine() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPaintDevice_paintEngine, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPaintDevice) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) WidthMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_widthMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) HeightMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_heightMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) LogicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_logicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) LogicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_logicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) PhysicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_physicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) PhysicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_physicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) DevicePixelRatio() float64 {
	__r := pg_QPaintDevice_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPaintDevice) DevicePixelRatioF() float64 {
	__r := pg_QPaintDevice_devicePixelRatioF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPaintDevice) ColorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_colorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDevice) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDevice_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QPaintDevice_DevicePixelRatioFScale() float64 {
	__r := pg_QPaintDevice_devicePixelRatioFScale()
	return __r
}

func QPaintDevice_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPaintDevice_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPaintDevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPaintDevice_delete(this.h)
	}
	this.UnsafeRelease()
}
