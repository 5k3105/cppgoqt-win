package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBitmap_chain uintptr
	addr_QBitmap_new uintptr
	addr_QBitmap_new2 uintptr
	addr_QBitmap_new3 uintptr
	addr_QBitmap_new4 uintptr
	addr_QBitmap_new5 uintptr
	addr_QBitmap_new6 uintptr
	addr_QBitmap_new7 uintptr
	addr_QBitmap_OperatorAssign uintptr
	addr_QBitmap_swap uintptr
	addr_QBitmap_ToQVariant uintptr
	addr_QBitmap_clear uintptr
	addr_QBitmap_fromImage uintptr
	addr_QBitmap_fromData uintptr
	addr_QBitmap_fromPixmap uintptr
	addr_QBitmap_transformed uintptr
	addr_QBitmap_OperatorAssign2 uintptr
	addr_QBitmap_defaultDepth uintptr
	addr_QBitmap_trueMatrix uintptr
	addr_QBitmap_fromImageReader uintptr
	pg_QBitmap_devicePixelRatioFScale func() float64
	pg_QBitmap_encodeMetricF func(int32, float64) int32
	pg_QBitmap_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QBitmap_fromImage2 uintptr
	addr_QBitmap_fromData2 uintptr
	addr_QBitmap_fromImageReader2 uintptr
	addr_QBitmap_toQVariant uintptr
	addr_QBitmap_devType_default uintptr
	addr_QBitmap_paintEngine_default uintptr
	addr_QBitmap_metric_default uintptr
	addr_QBitmap_initPainter_default uintptr
	addr_QBitmap_redirected_default uintptr
	addr_QBitmap_sharedPainter_default uintptr
	pg_QBitmap__set_override func(uintptr, int32)
	pg_QBitmap_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBitmap_chain, "QBitmap_chain")
	symaddr(&addr_QBitmap_new, "QBitmap_new")
	symaddr(&addr_QBitmap_new2, "QBitmap_new2")
	symaddr(&addr_QBitmap_new3, "QBitmap_new3")
	symaddr(&addr_QBitmap_new4, "QBitmap_new4")
	symaddr(&addr_QBitmap_new5, "QBitmap_new5")
	symaddr(&addr_QBitmap_new6, "QBitmap_new6")
	symaddr(&addr_QBitmap_new7, "QBitmap_new7")
	symaddr(&addr_QBitmap_OperatorAssign, "QBitmap_OperatorAssign")
	symaddr(&addr_QBitmap_swap, "QBitmap_swap")
	symaddr(&addr_QBitmap_ToQVariant, "QBitmap_ToQVariant")
	symaddr(&addr_QBitmap_clear, "QBitmap_clear")
	symaddr(&addr_QBitmap_fromImage, "QBitmap_fromImage")
	symaddr(&addr_QBitmap_fromData, "QBitmap_fromData")
	symaddr(&addr_QBitmap_fromPixmap, "QBitmap_fromPixmap")
	symaddr(&addr_QBitmap_transformed, "QBitmap_transformed")
	symaddr(&addr_QBitmap_OperatorAssign2, "QBitmap_OperatorAssign2")
	symaddr(&addr_QBitmap_defaultDepth, "QBitmap_defaultDepth")
	symaddr(&addr_QBitmap_trueMatrix, "QBitmap_trueMatrix")
	symaddr(&addr_QBitmap_fromImageReader, "QBitmap_fromImageReader")
	reg(&pg_QBitmap_devicePixelRatioFScale, "QBitmap_devicePixelRatioFScale")
	reg(&pg_QBitmap_encodeMetricF, "QBitmap_encodeMetricF")
	reg(&pg_QBitmap_getDecodedMetricF, "QBitmap_getDecodedMetricF")
	symaddr(&addr_QBitmap_fromImage2, "QBitmap_fromImage2")
	symaddr(&addr_QBitmap_fromData2, "QBitmap_fromData2")
	symaddr(&addr_QBitmap_fromImageReader2, "QBitmap_fromImageReader2")
	symaddr(&addr_QBitmap_toQVariant, "QBitmap_toQVariant")
	symaddr(&addr_QBitmap_devType_default, "QBitmap_devType_default")
	symaddr(&addr_QBitmap_paintEngine_default, "QBitmap_paintEngine_default")
	symaddr(&addr_QBitmap_metric_default, "QBitmap_metric_default")
	symaddr(&addr_QBitmap_initPainter_default, "QBitmap_initPainter_default")
	symaddr(&addr_QBitmap_redirected_default, "QBitmap_redirected_default")
	symaddr(&addr_QBitmap_sharedPainter_default, "QBitmap_sharedPainter_default")
	reg(&pg_QBitmap__set_override, "QBitmap__set_override")
	reg(&pg_QBitmap_delete, "QBitmap_delete")
}

type QBitmap struct {
	h uintptr
	*QPixmap
}

func bindQBitmap(h uintptr) *QBitmap {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QBitmap_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QBitmap{h: h, QPixmap: &QPixmap{h: __c[0], QPaintDevice: &QPaintDevice{h: __c[1]}}}
}

func newownedQBitmap(h uintptr) *QBitmap {
	o := bindQBitmap(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBitmap).Delete)
	}
	return o
}

func (this *QBitmap) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBitmap) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBitmap(h unsafe.Pointer) *QBitmap { return bindQBitmap(uintptr(h)) }
func UnsafeNewQBitmapOwned(h unsafe.Pointer) *QBitmap { return newownedQBitmap(uintptr(h)) }

func (this *QBitmap) UnsafeRelease() {
	this.h = 0
	if this.QPixmap != nil {
		this.QPixmap.UnsafeRelease()
	}
}

func NewQBitmap() *QBitmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBitmap_new, uintptr(__goid))
	__o := newownedQBitmap(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap2(arg1 *QPixmap) *QBitmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBitmap_new2, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQBitmap(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap3(w int, h int) *QBitmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBitmap_new3, uintptr(__goid), uintptr(w), uintptr(h))
	__o := newownedQBitmap(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap4(arg1 *QSize) *QBitmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBitmap_new4, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQBitmap(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap5(fileName string) *QBitmap {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QBitmap_new5, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQBitmap(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap6(arg1 *QBitmap) *QBitmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBitmap_new6, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQBitmap(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func NewQBitmap7(fileName string, format string) *QBitmap {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QBitmap_new7, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__o := newownedQBitmap(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBitmap_virtIdx, QBitmap_enable)
	return __o
}

func (this *QBitmap) OperatorAssign(arg1 *QPixmap) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQBitmap(__r)
}

func (this *QBitmap) Swap(other *QBitmap) {
	purego.SyscallN(addr_QBitmap_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QBitmap) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QBitmap_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QBitmap) Clear() {
	purego.SyscallN(addr_QBitmap_clear, this.h)
	runtime.KeepAlive(this)
}

func QBitmap_FromImage(image *QImage) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromImage, uintptr(image.UnsafePointer()))
	runtime.KeepAlive(image)
	return newownedQBitmap(__r)
}

func QBitmap_FromData(size *QSize, bits []byte) *QBitmap {
	var bits_p uintptr
	if len(bits) > 0 {
		bits_p = uintptr(unsafe.Pointer(&bits[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromData, uintptr(size.UnsafePointer()), bits_p)
	runtime.KeepAlive(size)
	runtime.KeepAlive(bits)
	return newownedQBitmap(__r)
}

func QBitmap_FromPixmap(pixmap *QPixmap) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromPixmap, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(pixmap)
	return newownedQBitmap(__r)
}

func (this *QBitmap) Transformed(matrix *QTransform) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_transformed, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQBitmap(__r)
}

func (this *QBitmap) OperatorAssignWithQBitmap(arg1 *QBitmap) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_OperatorAssign2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQBitmap(__r)
}

func QBitmap_DefaultDepth() int {
	__r0, _, _ := purego.SyscallN(addr_QBitmap_defaultDepth)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QBitmap_TrueMatrix(m *QTransform, w int, h int) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QBitmap_trueMatrix, uintptr(m.UnsafePointer()), uintptr(w), uintptr(h))
	runtime.KeepAlive(m)
	return newownedQTransform(__r)
}

func QBitmap_FromImageReader(imageReader *QImageReader) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromImageReader, uintptr(imageReader.UnsafePointer()))
	runtime.KeepAlive(imageReader)
	return newownedQPixmap(__r)
}

func QBitmap_DevicePixelRatioFScale() float64 {
	__r := pg_QBitmap_devicePixelRatioFScale()
	return __r
}

func QBitmap_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QBitmap_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QBitmap) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QBitmap_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QBitmap_FromImage2(image *QImage, flags ImageConversionFlag) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromImage2, uintptr(image.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(image)
	return newownedQBitmap(__r)
}

func QBitmap_FromData2(size *QSize, bits []byte, monoFormat QImage__Format) *QBitmap {
	var bits_p uintptr
	if len(bits) > 0 {
		bits_p = uintptr(unsafe.Pointer(&bits[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromData2, uintptr(size.UnsafePointer()), bits_p, uintptr(monoFormat))
	runtime.KeepAlive(size)
	runtime.KeepAlive(bits)
	return newownedQBitmap(__r)
}

func QBitmap_FromImageReader2(imageReader *QImageReader, flags ImageConversionFlag) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QBitmap_fromImageReader2, uintptr(imageReader.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(imageReader)
	return newownedQPixmap(__r)
}

func (this *QBitmap) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QBitmap_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QBitmap) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QBitmap_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QBitmap) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QBitmap_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QBitmap) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QBitmap_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QBitmap) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QBitmap_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QBitmap) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QBitmap_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QBitmap) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QBitmap_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func QBitmap_virtIdx(slot int64) int32 {
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

func QBitmap_enable(h uintptr, idx int32) { pg_QBitmap__set_override(h, idx) }

func (this *QBitmap) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QBitmap) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QBitmap) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QBitmap) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QBitmap) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QBitmap) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QBitmap) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBitmap_delete(this.h)
	}
	this.UnsafeRelease()
}
