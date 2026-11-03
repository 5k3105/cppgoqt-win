package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPixmap_base func(uintptr) uintptr
	addr_QPixmap_new uintptr
	addr_QPixmap_new2 uintptr
	addr_QPixmap_new3 uintptr
	addr_QPixmap_new4 uintptr
	addr_QPixmap_new5 uintptr
	addr_QPixmap_new6 uintptr
	addr_QPixmap_new7 uintptr
	addr_QPixmap_OperatorAssign uintptr
	addr_QPixmap_swap uintptr
	addr_QPixmap_ToQVariant uintptr
	addr_QPixmap_isNull uintptr
	addr_QPixmap_width uintptr
	addr_QPixmap_height uintptr
	addr_QPixmap_size uintptr
	addr_QPixmap_rect uintptr
	addr_QPixmap_depth uintptr
	addr_QPixmap_defaultDepth uintptr
	addr_QPixmap_fill uintptr
	addr_QPixmap_mask uintptr
	addr_QPixmap_setMask uintptr
	pg_QPixmap_devicePixelRatio func(uintptr) float64
	pg_QPixmap_setDevicePixelRatio func(uintptr, float64)
	addr_QPixmap_deviceIndependentSize uintptr
	addr_QPixmap_hasAlpha uintptr
	addr_QPixmap_hasAlphaChannel uintptr
	addr_QPixmap_createHeuristicMask uintptr
	addr_QPixmap_createMaskFromColor uintptr
	addr_QPixmap_scaled uintptr
	addr_QPixmap_scaled2 uintptr
	addr_QPixmap_scaledToWidth uintptr
	addr_QPixmap_scaledToHeight uintptr
	addr_QPixmap_transformed uintptr
	addr_QPixmap_trueMatrix uintptr
	addr_QPixmap_toImage uintptr
	addr_QPixmap_fromImage uintptr
	addr_QPixmap_fromImageReader uintptr
	addr_QPixmap_load uintptr
	addr_QPixmap_loadFromData uintptr
	addr_QPixmap_loadFromData2 uintptr
	addr_QPixmap_save uintptr
	addr_QPixmap_save2 uintptr
	addr_QPixmap_convertFromImage uintptr
	addr_QPixmap_copy uintptr
	addr_QPixmap_copy2 uintptr
	addr_QPixmap_scroll uintptr
	addr_QPixmap_scroll2 uintptr
	addr_QPixmap_cacheKey uintptr
	addr_QPixmap_isDetached uintptr
	addr_QPixmap_detach uintptr
	addr_QPixmap_isQBitmap uintptr
	addr_QPixmap_OperatorNot uintptr
	pg_QPixmap_devicePixelRatioFScale func() float64
	pg_QPixmap_encodeMetricF func(int32, float64) int32
	pg_QPixmap_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QPixmap_fill2 uintptr
	addr_QPixmap_createHeuristicMask2 uintptr
	addr_QPixmap_createMaskFromColor2 uintptr
	addr_QPixmap_scaled3 uintptr
	addr_QPixmap_scaled4 uintptr
	addr_QPixmap_scaled5 uintptr
	addr_QPixmap_scaled6 uintptr
	addr_QPixmap_scaledToWidth2 uintptr
	addr_QPixmap_scaledToHeight2 uintptr
	addr_QPixmap_transformed2 uintptr
	addr_QPixmap_fromImage2 uintptr
	addr_QPixmap_fromImageReader2 uintptr
	addr_QPixmap_load2 uintptr
	addr_QPixmap_load3 uintptr
	addr_QPixmap_loadFromData3 uintptr
	addr_QPixmap_loadFromData4 uintptr
	addr_QPixmap_loadFromData5 uintptr
	addr_QPixmap_loadFromData6 uintptr
	addr_QPixmap_save3 uintptr
	addr_QPixmap_save4 uintptr
	addr_QPixmap_save5 uintptr
	addr_QPixmap_save6 uintptr
	addr_QPixmap_convertFromImage2 uintptr
	addr_QPixmap_copy3 uintptr
	addr_QPixmap_scroll3 uintptr
	addr_QPixmap_scroll4 uintptr
	addr_QPixmap_toQVariant uintptr
	addr_QPixmap_devType_default uintptr
	addr_QPixmap_paintEngine_default uintptr
	addr_QPixmap_metric_default uintptr
	addr_QPixmap_initPainter_default uintptr
	addr_QPixmap_redirected_default uintptr
	addr_QPixmap_sharedPainter_default uintptr
	pg_QPixmap__set_override func(uintptr, int32)
	pg_QPixmap_delete func(uintptr)
)

func init() {
	reg(&pg_QPixmap_base, "QPixmap_base")
	symaddr(&addr_QPixmap_new, "QPixmap_new")
	symaddr(&addr_QPixmap_new2, "QPixmap_new2")
	symaddr(&addr_QPixmap_new3, "QPixmap_new3")
	symaddr(&addr_QPixmap_new4, "QPixmap_new4")
	symaddr(&addr_QPixmap_new5, "QPixmap_new5")
	symaddr(&addr_QPixmap_new6, "QPixmap_new6")
	symaddr(&addr_QPixmap_new7, "QPixmap_new7")
	symaddr(&addr_QPixmap_OperatorAssign, "QPixmap_OperatorAssign")
	symaddr(&addr_QPixmap_swap, "QPixmap_swap")
	symaddr(&addr_QPixmap_ToQVariant, "QPixmap_ToQVariant")
	symaddr(&addr_QPixmap_isNull, "QPixmap_isNull")
	symaddr(&addr_QPixmap_width, "QPixmap_width")
	symaddr(&addr_QPixmap_height, "QPixmap_height")
	symaddr(&addr_QPixmap_size, "QPixmap_size")
	symaddr(&addr_QPixmap_rect, "QPixmap_rect")
	symaddr(&addr_QPixmap_depth, "QPixmap_depth")
	symaddr(&addr_QPixmap_defaultDepth, "QPixmap_defaultDepth")
	symaddr(&addr_QPixmap_fill, "QPixmap_fill")
	symaddr(&addr_QPixmap_mask, "QPixmap_mask")
	symaddr(&addr_QPixmap_setMask, "QPixmap_setMask")
	reg(&pg_QPixmap_devicePixelRatio, "QPixmap_devicePixelRatio")
	reg(&pg_QPixmap_setDevicePixelRatio, "QPixmap_setDevicePixelRatio")
	symaddr(&addr_QPixmap_deviceIndependentSize, "QPixmap_deviceIndependentSize")
	symaddr(&addr_QPixmap_hasAlpha, "QPixmap_hasAlpha")
	symaddr(&addr_QPixmap_hasAlphaChannel, "QPixmap_hasAlphaChannel")
	symaddr(&addr_QPixmap_createHeuristicMask, "QPixmap_createHeuristicMask")
	symaddr(&addr_QPixmap_createMaskFromColor, "QPixmap_createMaskFromColor")
	symaddr(&addr_QPixmap_scaled, "QPixmap_scaled")
	symaddr(&addr_QPixmap_scaled2, "QPixmap_scaled2")
	symaddr(&addr_QPixmap_scaledToWidth, "QPixmap_scaledToWidth")
	symaddr(&addr_QPixmap_scaledToHeight, "QPixmap_scaledToHeight")
	symaddr(&addr_QPixmap_transformed, "QPixmap_transformed")
	symaddr(&addr_QPixmap_trueMatrix, "QPixmap_trueMatrix")
	symaddr(&addr_QPixmap_toImage, "QPixmap_toImage")
	symaddr(&addr_QPixmap_fromImage, "QPixmap_fromImage")
	symaddr(&addr_QPixmap_fromImageReader, "QPixmap_fromImageReader")
	symaddr(&addr_QPixmap_load, "QPixmap_load")
	symaddr(&addr_QPixmap_loadFromData, "QPixmap_loadFromData")
	symaddr(&addr_QPixmap_loadFromData2, "QPixmap_loadFromData2")
	symaddr(&addr_QPixmap_save, "QPixmap_save")
	symaddr(&addr_QPixmap_save2, "QPixmap_save2")
	symaddr(&addr_QPixmap_convertFromImage, "QPixmap_convertFromImage")
	symaddr(&addr_QPixmap_copy, "QPixmap_copy")
	symaddr(&addr_QPixmap_copy2, "QPixmap_copy2")
	symaddr(&addr_QPixmap_scroll, "QPixmap_scroll")
	symaddr(&addr_QPixmap_scroll2, "QPixmap_scroll2")
	symaddr(&addr_QPixmap_cacheKey, "QPixmap_cacheKey")
	symaddr(&addr_QPixmap_isDetached, "QPixmap_isDetached")
	symaddr(&addr_QPixmap_detach, "QPixmap_detach")
	symaddr(&addr_QPixmap_isQBitmap, "QPixmap_isQBitmap")
	symaddr(&addr_QPixmap_OperatorNot, "QPixmap_OperatorNot")
	reg(&pg_QPixmap_devicePixelRatioFScale, "QPixmap_devicePixelRatioFScale")
	reg(&pg_QPixmap_encodeMetricF, "QPixmap_encodeMetricF")
	reg(&pg_QPixmap_getDecodedMetricF, "QPixmap_getDecodedMetricF")
	symaddr(&addr_QPixmap_fill2, "QPixmap_fill2")
	symaddr(&addr_QPixmap_createHeuristicMask2, "QPixmap_createHeuristicMask2")
	symaddr(&addr_QPixmap_createMaskFromColor2, "QPixmap_createMaskFromColor2")
	symaddr(&addr_QPixmap_scaled3, "QPixmap_scaled3")
	symaddr(&addr_QPixmap_scaled4, "QPixmap_scaled4")
	symaddr(&addr_QPixmap_scaled5, "QPixmap_scaled5")
	symaddr(&addr_QPixmap_scaled6, "QPixmap_scaled6")
	symaddr(&addr_QPixmap_scaledToWidth2, "QPixmap_scaledToWidth2")
	symaddr(&addr_QPixmap_scaledToHeight2, "QPixmap_scaledToHeight2")
	symaddr(&addr_QPixmap_transformed2, "QPixmap_transformed2")
	symaddr(&addr_QPixmap_fromImage2, "QPixmap_fromImage2")
	symaddr(&addr_QPixmap_fromImageReader2, "QPixmap_fromImageReader2")
	symaddr(&addr_QPixmap_load2, "QPixmap_load2")
	symaddr(&addr_QPixmap_load3, "QPixmap_load3")
	symaddr(&addr_QPixmap_loadFromData3, "QPixmap_loadFromData3")
	symaddr(&addr_QPixmap_loadFromData4, "QPixmap_loadFromData4")
	symaddr(&addr_QPixmap_loadFromData5, "QPixmap_loadFromData5")
	symaddr(&addr_QPixmap_loadFromData6, "QPixmap_loadFromData6")
	symaddr(&addr_QPixmap_save3, "QPixmap_save3")
	symaddr(&addr_QPixmap_save4, "QPixmap_save4")
	symaddr(&addr_QPixmap_save5, "QPixmap_save5")
	symaddr(&addr_QPixmap_save6, "QPixmap_save6")
	symaddr(&addr_QPixmap_convertFromImage2, "QPixmap_convertFromImage2")
	symaddr(&addr_QPixmap_copy3, "QPixmap_copy3")
	symaddr(&addr_QPixmap_scroll3, "QPixmap_scroll3")
	symaddr(&addr_QPixmap_scroll4, "QPixmap_scroll4")
	symaddr(&addr_QPixmap_toQVariant, "QPixmap_toQVariant")
	symaddr(&addr_QPixmap_devType_default, "QPixmap_devType_default")
	symaddr(&addr_QPixmap_paintEngine_default, "QPixmap_paintEngine_default")
	symaddr(&addr_QPixmap_metric_default, "QPixmap_metric_default")
	symaddr(&addr_QPixmap_initPainter_default, "QPixmap_initPainter_default")
	symaddr(&addr_QPixmap_redirected_default, "QPixmap_redirected_default")
	symaddr(&addr_QPixmap_sharedPainter_default, "QPixmap_sharedPainter_default")
	reg(&pg_QPixmap__set_override, "QPixmap__set_override")
	reg(&pg_QPixmap_delete, "QPixmap_delete")
}

type QPixmap struct {
	h uintptr
	*QPaintDevice
}

func bindQPixmap(h uintptr) *QPixmap {
	if h == 0 {
		return nil
	}
	return &QPixmap{h: h, QPaintDevice: bindQPaintDevice(pg_QPixmap_base(h))}
}

func newownedQPixmap(h uintptr) *QPixmap {
	o := bindQPixmap(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPixmap).Delete)
	}
	return o
}

func (this *QPixmap) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPixmap) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPixmap(h unsafe.Pointer) *QPixmap { return bindQPixmap(uintptr(h)) }
func UnsafeNewQPixmapOwned(h unsafe.Pointer) *QPixmap { return newownedQPixmap(uintptr(h)) }

func (this *QPixmap) UnsafeRelease() {
	this.h = 0
	if this.QPaintDevice != nil {
		this.QPaintDevice.UnsafeRelease()
	}
}

func NewQPixmap() *QPixmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPixmap_new, uintptr(__goid))
	__o := newownedQPixmap(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap2(w int, h int) *QPixmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPixmap_new2, uintptr(__goid), uintptr(w), uintptr(h))
	__o := newownedQPixmap(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap3(arg1 *QSize) *QPixmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPixmap_new3, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQPixmap(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap4(fileName string) *QPixmap {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QPixmap_new4, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQPixmap(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap5(arg1 *QPixmap) *QPixmap {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QPixmap_new5, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQPixmap(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap6(fileName string, format string) *QPixmap {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QPixmap_new6, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__o := newownedQPixmap(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func NewQPixmap7(fileName string, format string, flags ImageConversionFlag) *QPixmap {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QPixmap_new7, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])), uintptr(flags))
	__o := newownedQPixmap(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QPixmap_virtIdx, QPixmap_enable)
	return __o
}

func (this *QPixmap) OperatorAssign(arg1 *QPixmap) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPixmap(__r)
}

func (this *QPixmap) Swap(other *QPixmap) {
	purego.SyscallN(addr_QPixmap_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPixmap) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPixmap_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPixmap) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmap) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPixmap) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPixmap) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPixmap_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPixmap) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPixmap_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPixmap) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QPixmap_DefaultDepth() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_defaultDepth)
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QPixmap) Fill() {
	purego.SyscallN(addr_QPixmap_fill, this.h)
	runtime.KeepAlive(this)
}

func (this *QPixmap) Mask() *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_mask, this.h)
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QPixmap) SetMask(arg1 *QBitmap) {
	purego.SyscallN(addr_QPixmap_setMask, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPixmap) DevicePixelRatio() float64 {
	__r := pg_QPixmap_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPixmap) SetDevicePixelRatio(scaleFactor float64) {
	pg_QPixmap_setDevicePixelRatio(this.h, scaleFactor)
	runtime.KeepAlive(this)
}

func (this *QPixmap) DeviceIndependentSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QPixmap_deviceIndependentSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QPixmap) HasAlpha() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_hasAlpha, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmap) HasAlphaChannel() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_hasAlphaChannel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmap) CreateHeuristicMask() *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_createHeuristicMask, this.h)
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QPixmap) CreateMaskFromColor(maskColor *QColor) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_createMaskFromColor, this.h, uintptr(maskColor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(maskColor)
	return newownedQBitmap(__r)
}

func (this *QPixmap) Scaled(w int, h int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) ScaledWithQSize(s *QSize) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled2, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQPixmap(__r)
}

func (this *QPixmap) ScaledToWidth(w int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaledToWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) ScaledToHeight(h int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaledToHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Transformed(arg1 *QTransform) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_transformed, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPixmap(__r)
}

func QPixmap_TrueMatrix(m *QTransform, w int, h int) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPixmap_trueMatrix, uintptr(m.UnsafePointer()), uintptr(w), uintptr(h))
	runtime.KeepAlive(m)
	return newownedQTransform(__r)
}

func (this *QPixmap) ToImage() *QImage {
	__r, _, _ := purego.SyscallN(addr_QPixmap_toImage, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func QPixmap_FromImage(image *QImage) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_fromImage, uintptr(image.UnsafePointer()))
	runtime.KeepAlive(image)
	return newownedQPixmap(__r)
}

func QPixmap_FromImageReader(imageReader *QImageReader) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_fromImageReader, uintptr(imageReader.UnsafePointer()))
	runtime.KeepAlive(imageReader)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Load(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPixmap_load, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QPixmap) LoadFromData(buf []byte, len_ uint) bool {
	var buf_p uintptr
	if len(buf) > 0 {
		buf_p = uintptr(unsafe.Pointer(&buf[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData, this.h, buf_p, uintptr(len_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(buf)
	return __r != 0
}

func (this *QPixmap) LoadFromDataWithData(data *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData2, this.h, uintptr(data.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QPixmap) Save(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QPixmap) SaveWithDevice(device *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save2, this.h, uintptr(device.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	return __r != 0
}

func (this *QPixmap) ConvertFromImage(img *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_convertFromImage, this.h, uintptr(img.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(img)
	return __r != 0
}

func (this *QPixmap) Copy(x int, y int, width int, height int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_copy, this.h, uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Copy2() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_copy2, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Scroll(dx int, dy int, x int, y int, width int, height int) {
	purego.SyscallN(addr_QPixmap_scroll, this.h, uintptr(dx), uintptr(dy), uintptr(x), uintptr(y), uintptr(width), uintptr(height))
	runtime.KeepAlive(this)
}

func (this *QPixmap) Scroll2(dx int, dy int, rect *QRect) {
	purego.SyscallN(addr_QPixmap_scroll2, this.h, uintptr(dx), uintptr(dy), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPixmap) CacheKey() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_cacheKey, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPixmap) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmap) Detach() {
	purego.SyscallN(addr_QPixmap_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QPixmap) IsQBitmap() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_isQBitmap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPixmap) OperatorNot() bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_OperatorNot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QPixmap_DevicePixelRatioFScale() float64 {
	__r := pg_QPixmap_devicePixelRatioFScale()
	return __r
}

func QPixmap_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPixmap_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QPixmap) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QPixmap_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func (this *QPixmap) FillWithFillColor(fillColor *QColor) {
	purego.SyscallN(addr_QPixmap_fill2, this.h, uintptr(fillColor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fillColor)
}

func (this *QPixmap) CreateHeuristicMaskWithClipTight(clipTight bool) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_createHeuristicMask2, this.h, uintptr(b2i(clipTight)))
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QPixmap) CreateMaskFromColor2(maskColor *QColor, mode MaskMode) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_createMaskFromColor2, this.h, uintptr(maskColor.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(maskColor)
	return newownedQBitmap(__r)
}

func (this *QPixmap) Scaled2(w int, h int, aspectMode AspectRatioMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled3, this.h, uintptr(w), uintptr(h), uintptr(aspectMode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Scaled3(w int, h int, aspectMode AspectRatioMode, mode TransformationMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled4, this.h, uintptr(w), uintptr(h), uintptr(aspectMode), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Scaled4(s *QSize, aspectMode AspectRatioMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled5, this.h, uintptr(s.UnsafePointer()), uintptr(aspectMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Scaled5(s *QSize, aspectMode AspectRatioMode, mode TransformationMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaled6, this.h, uintptr(s.UnsafePointer()), uintptr(aspectMode), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQPixmap(__r)
}

func (this *QPixmap) ScaledToWidth2(w int, mode TransformationMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaledToWidth2, this.h, uintptr(w), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) ScaledToHeight2(h int, mode TransformationMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_scaledToHeight2, this.h, uintptr(h), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Transformed2(arg1 *QTransform, mode TransformationMode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_transformed2, this.h, uintptr(arg1.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPixmap(__r)
}

func QPixmap_FromImage2(image *QImage, flags ImageConversionFlag) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_fromImage2, uintptr(image.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(image)
	return newownedQPixmap(__r)
}

func QPixmap_FromImageReader2(imageReader *QImageReader, flags ImageConversionFlag) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_fromImageReader2, uintptr(imageReader.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(imageReader)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Load2(fileName string, format string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_load2, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) Load3(fileName string, format string, flags ImageConversionFlag) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_load3, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) LoadFromData2(buf []byte, len_ uint, format string) bool {
	var buf_p uintptr
	if len(buf) > 0 {
		buf_p = uintptr(unsafe.Pointer(&buf[0]))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData3, this.h, buf_p, uintptr(len_), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) LoadFromData3(buf []byte, len_ uint, format string, flags ImageConversionFlag) bool {
	var buf_p uintptr
	if len(buf) > 0 {
		buf_p = uintptr(unsafe.Pointer(&buf[0]))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData4, this.h, buf_p, uintptr(len_), uintptr(unsafe.Pointer(&format_b[0])), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) LoadFromData4(data *QByteArray, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData5, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) LoadFromData5(data *QByteArray, format string, flags ImageConversionFlag) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_loadFromData6, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) Save2(fileName string, format string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save3, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) Save3(fileName string, format string, quality int) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save4, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])), uintptr(quality))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) Save4(device *QIODevice, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save5, this.h, uintptr(device.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) Save5(device *QIODevice, format string, quality int) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QPixmap_save6, this.h, uintptr(device.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])), uintptr(quality))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QPixmap) ConvertFromImage2(img *QImage, flags ImageConversionFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_convertFromImage2, this.h, uintptr(img.UnsafePointer()), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(img)
	return __r != 0
}

func (this *QPixmap) CopyWithRect(rect *QRect) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QPixmap_copy3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPixmap(__r)
}

func (this *QPixmap) Scroll3(dx int, dy int, x int, y int, width int, height int, exposed *QRegion) {
	purego.SyscallN(addr_QPixmap_scroll3, this.h, uintptr(dx), uintptr(dy), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(exposed.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(exposed)
}

func (this *QPixmap) Scroll4(dx int, dy int, rect *QRect, exposed *QRegion) {
	purego.SyscallN(addr_QPixmap_scroll4, this.h, uintptr(dx), uintptr(dy), uintptr(rect.UnsafePointer()), uintptr(exposed.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(exposed)
}

func (this *QPixmap) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPixmap_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPixmap) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPixmap) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPixmap_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPixmap) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QPixmap_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPixmap) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QPixmap_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QPixmap) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPixmap_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QPixmap) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPixmap_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func QPixmap_virtIdx(slot int64) int32 {
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

func QPixmap_enable(h uintptr, idx int32) { pg_QPixmap__set_override(h, idx) }

func (this *QPixmap) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QPixmap) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QPixmap) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QPixmap) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QPixmap) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QPixmap) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QPixmap) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPixmap_delete(this.h)
	}
	this.UnsafeRelease()
}
