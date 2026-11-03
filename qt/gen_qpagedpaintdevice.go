package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPagedPaintDevice_base func(uintptr) uintptr
	addr_QPagedPaintDevice_newPage uintptr
	addr_QPagedPaintDevice_setPageLayout uintptr
	addr_QPagedPaintDevice_setPageSize uintptr
	addr_QPagedPaintDevice_setPageOrientation uintptr
	addr_QPagedPaintDevice_setPageMargins uintptr
	addr_QPagedPaintDevice_pageLayout uintptr
	addr_QPagedPaintDevice_setPageRanges uintptr
	addr_QPagedPaintDevice_pageRanges uintptr
	pg_QPagedPaintDevice_devicePixelRatioFScale func() float64
	pg_QPagedPaintDevice_encodeMetricF func(int32, float64) int32
	pg_QPagedPaintDevice_delete func(uintptr)
)

func init() {
	reg(&pg_QPagedPaintDevice_base, "QPagedPaintDevice_base")
	symaddr(&addr_QPagedPaintDevice_newPage, "QPagedPaintDevice_newPage")
	symaddr(&addr_QPagedPaintDevice_setPageLayout, "QPagedPaintDevice_setPageLayout")
	symaddr(&addr_QPagedPaintDevice_setPageSize, "QPagedPaintDevice_setPageSize")
	symaddr(&addr_QPagedPaintDevice_setPageOrientation, "QPagedPaintDevice_setPageOrientation")
	symaddr(&addr_QPagedPaintDevice_setPageMargins, "QPagedPaintDevice_setPageMargins")
	symaddr(&addr_QPagedPaintDevice_pageLayout, "QPagedPaintDevice_pageLayout")
	symaddr(&addr_QPagedPaintDevice_setPageRanges, "QPagedPaintDevice_setPageRanges")
	symaddr(&addr_QPagedPaintDevice_pageRanges, "QPagedPaintDevice_pageRanges")
	reg(&pg_QPagedPaintDevice_devicePixelRatioFScale, "QPagedPaintDevice_devicePixelRatioFScale")
	reg(&pg_QPagedPaintDevice_encodeMetricF, "QPagedPaintDevice_encodeMetricF")
	reg(&pg_QPagedPaintDevice_delete, "QPagedPaintDevice_delete")
}

type QPagedPaintDevice struct {
	h uintptr
	*QPaintDevice
}

func bindQPagedPaintDevice(h uintptr) *QPagedPaintDevice {
	if h == 0 {
		return nil
	}
	return &QPagedPaintDevice{h: h, QPaintDevice: bindQPaintDevice(pg_QPagedPaintDevice_base(h))}
}

func newownedQPagedPaintDevice(h uintptr) *QPagedPaintDevice {
	o := bindQPagedPaintDevice(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPagedPaintDevice).Delete)
	}
	return o
}

func (this *QPagedPaintDevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPagedPaintDevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPagedPaintDevice(h unsafe.Pointer) *QPagedPaintDevice { return bindQPagedPaintDevice(uintptr(h)) }
func UnsafeNewQPagedPaintDeviceOwned(h unsafe.Pointer) *QPagedPaintDevice { return newownedQPagedPaintDevice(uintptr(h)) }

func (this *QPagedPaintDevice) UnsafeRelease() {
	this.h = 0
	if this.QPaintDevice != nil {
		this.QPaintDevice.UnsafeRelease()
	}
}

func (this *QPagedPaintDevice) NewPage() bool {
	__r0, _, _ := purego.SyscallN(addr_QPagedPaintDevice_newPage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPagedPaintDevice) SetPageLayout(pageLayout *QPageLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QPagedPaintDevice_setPageLayout, this.h, uintptr(pageLayout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageLayout)
	return __r != 0
}

func (this *QPagedPaintDevice) SetPageSize(pageSize *QPageSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QPagedPaintDevice_setPageSize, this.h, uintptr(pageSize.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageSize)
	return __r != 0
}

func (this *QPagedPaintDevice) SetPageOrientation(orientation QPageLayout__Orientation) bool {
	__r0, _, _ := purego.SyscallN(addr_QPagedPaintDevice_setPageOrientation, this.h, uintptr(orientation))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPagedPaintDevice) SetPageMargins(margins *QMarginsF, units QPageLayout__Unit) bool {
	__r0, _, _ := purego.SyscallN(addr_QPagedPaintDevice_setPageMargins, this.h, uintptr(margins.UnsafePointer()), uintptr(units))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return __r != 0
}

func (this *QPagedPaintDevice) PageLayout() *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPagedPaintDevice_pageLayout, this.h)
	runtime.KeepAlive(this)
	return newownedQPageLayout(__r)
}

func (this *QPagedPaintDevice) SetPageRanges(ranges *QPageRanges) {
	purego.SyscallN(addr_QPagedPaintDevice_setPageRanges, this.h, uintptr(ranges.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ranges)
}

func (this *QPagedPaintDevice) PageRanges() *QPageRanges {
	__r, _, _ := purego.SyscallN(addr_QPagedPaintDevice_pageRanges, this.h)
	runtime.KeepAlive(this)
	return newownedQPageRanges(__r)
}

func QPagedPaintDevice_DevicePixelRatioFScale() float64 {
	__r := pg_QPagedPaintDevice_devicePixelRatioFScale()
	return __r
}

func QPagedPaintDevice_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPagedPaintDevice_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPagedPaintDevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPagedPaintDevice_delete(this.h)
	}
	this.UnsafeRelease()
}
