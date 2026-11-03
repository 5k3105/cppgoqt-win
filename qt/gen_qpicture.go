package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPicture_base func(uintptr) uintptr
	addr_QPicture_new uintptr
	addr_QPicture_new2 uintptr
	addr_QPicture_new3 uintptr
	addr_QPicture_isNull uintptr
	addr_QPicture_size uintptr
	addr_QPicture_data uintptr
	addr_QPicture_setData uintptr
	addr_QPicture_play uintptr
	addr_QPicture_load uintptr
	addr_QPicture_load2 uintptr
	addr_QPicture_save uintptr
	addr_QPicture_save2 uintptr
	addr_QPicture_boundingRect uintptr
	addr_QPicture_setBoundingRect uintptr
	addr_QPicture_OperatorAssign uintptr
	addr_QPicture_swap uintptr
	addr_QPicture_detach uintptr
	addr_QPicture_isDetached uintptr
	pg_QPicture_devicePixelRatioFScale func() float64
	pg_QPicture_encodeMetricF func(int32, float64) int32
	pg_QPicture_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QPicture_devType_default uintptr
	addr_QPicture_paintEngine_default uintptr
	addr_QPicture_metric_default uintptr
	addr_QPicture_initPainter_default uintptr
	addr_QPicture_redirected_default uintptr
	addr_QPicture_sharedPainter_default uintptr
	pg_QPicture__set_override func(uintptr, int32)
	pg_QPicture_delete func(uintptr)
)

func init() {
	reg(&pg_QPicture_base, "QPicture_base")
	symaddr(&addr_QPicture_new, "QPicture_new")
	symaddr(&addr_QPicture_new2, "QPicture_new2")
	symaddr(&addr_QPicture_new3, "QPicture_new3")
	symaddr(&addr_QPicture_isNull, "QPicture_isNull")
	symaddr(&addr_QPicture_size, "QPicture_size")
	symaddr(&addr_QPicture_data, "QPicture_data")
	symaddr(&addr_QPicture_setData, "QPicture_setData")
	symaddr(&addr_QPicture_play, "QPicture_play")
	symaddr(&addr_QPicture_load, "QPicture_load")
	symaddr(&addr_QPicture_load2, "QPicture_load2")
	symaddr(&addr_QPicture_save, "QPicture_save")
	symaddr(&addr_QPicture_save2, "QPicture_save2")
	symaddr(&addr_QPicture_boundingRect, "QPicture_boundingRect")
	symaddr(&addr_QPicture_setBoundingRect, "QPicture_setBoundingRect")
	symaddr(&addr_QPicture_OperatorAssign, "QPicture_OperatorAssign")
	symaddr(&addr_QPicture_swap, "QPicture_swap")
	symaddr(&addr_QPicture_detach, "QPicture_detach")
	symaddr(&addr_QPicture_isDetached, "QPicture_isDetached")
	reg(&pg_QPicture_devicePixelRatioFScale, "QPicture_devicePixelRatioFScale")
	reg(&pg_QPicture_encodeMetricF, "QPicture_encodeMetricF")
	reg(&pg_QPicture_getDecodedMetricF, "QPicture_getDecodedMetricF")
	symaddr(&addr_QPicture_devType_default, "QPicture_devType_default")
	symaddr(&addr_QPicture_paintEngine_default, "QPicture_paintEngine_default")
	symaddr(&addr_QPicture_metric_default, "QPicture_metric_default")
	symaddr(&addr_QPicture_initPainter_default, "QPicture_initPainter_default")
	symaddr(&addr_QPicture_redirected_default, "QPicture_redirected_default")
	symaddr(&addr_QPicture_sharedPainter_default, "QPicture_sharedPainter_default")
	reg(&pg_QPicture__set_override, "QPicture__set_override")
	reg(&pg_QPicture_delete, "QPicture_delete")
}

type QPicture struct {
	h uintptr
	*QPaintDevice
}

func bindQPicture(h uintptr) *QPicture {
	if h == 0 {
		return nil
	}
	return &QPicture{h: h, QPaintDevice: bindQPaintDevice(pg_QPicture_base(h))}
}

func newownedQPicture(h uintptr) *QPicture {
	o := bindQPicture(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPicture).Delete)
	}
	return o
}

func (this *QPicture) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPicture) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPicture(h unsafe.Pointer) *QPicture { return bindQPicture(uintptr(h)) }
func UnsafeNewQPictureOwned(h unsafe.Pointer) *QPicture { return newownedQPicture(uintptr(h)) }

func (this *QPicture) UnsafeRelease() {
	this.h = 0
	if this.QPaintDevice != nil {
		this.QPaintDevice.UnsafeRelease()
	}
}

func NewQPicture() *QPicture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPicture_new, uintptr(__goid))
	__o := newownedQPicture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPicture_virtIdx, QPicture_enable)
	return __o
}

func NewQPicture2(arg1 *QPicture) *QPicture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPicture_new2, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQPicture(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPicture_virtIdx, QPicture_enable)
	return __o
}

func NewQPicture3(formatVersion int) *QPicture {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPicture_new3, uintptr(__goid), uintptr(formatVersion))
	__o := newownedQPicture(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPicture_virtIdx, QPicture_enable)
	return __o
}

func (this *QPicture) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QPicture_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPicture) Size() uint {
	__r0, _, _ := purego.SyscallN(addr_QPicture_size, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPicture) Data() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPicture_data, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPicture) SetData(data string, size uint) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QPicture_setData, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QPicture) Play(p *QPainter) bool {
	__r0, _, _ := purego.SyscallN(addr_QPicture_play, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPicture) Load(dev *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPicture_load, this.h, uintptr(dev.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dev)
	return __r != 0
}

func (this *QPicture) LoadWithFileName(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPicture_load2, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QPicture) Save(dev *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPicture_save, this.h, uintptr(dev.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dev)
	return __r != 0
}

func (this *QPicture) SaveWithFileName(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPicture_save2, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QPicture) BoundingRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPicture_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPicture) SetBoundingRect(r *QRect) {
	purego.SyscallN(addr_QPicture_setBoundingRect, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPicture) OperatorAssign(p *QPicture) *QPicture {
	__r, _, _ := purego.SyscallN(addr_QPicture_OperatorAssign, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQPicture(__r)
}

func (this *QPicture) Swap(other *QPicture) {
	purego.SyscallN(addr_QPicture_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPicture) Detach() {
	purego.SyscallN(addr_QPicture_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QPicture) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QPicture_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QPicture_DevicePixelRatioFScale() float64 {
	__r := pg_QPicture_devicePixelRatioFScale()
	return __r
}

func QPicture_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPicture_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPicture) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QPicture_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func (this *QPicture) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPicture_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPicture) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPicture_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPicture) MetricDefault(m QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QPicture_metric_default, this.h, uintptr(m))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPicture) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QPicture_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QPicture) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPicture_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QPicture) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPicture_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func QPicture_virtIdx(slot int64) int32 {
	switch slot {
	case -8659679893868524044:
		return 0
	case 6398007404756346381:
		return 1
	case 427629649267311777:
		return 2
	case 6477500296980050210:
		return 3
	case 4127240633310655048:
		return 4
	case 5792771328741774991:
		return 5
	}
	return -1
}

func QPicture_enable(h uintptr, idx int32) { pg_QPicture__set_override(h, idx) }

func (this *QPicture) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QPicture) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QPicture) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(m int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(m))) }) // slot "metric"
}

func (this *QPicture) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QPicture) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QPicture) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QPicture) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPicture_delete(this.h)
	}
	this.UnsafeRelease()
}
