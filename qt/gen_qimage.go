package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QImage_base func(uintptr) uintptr
	addr_QImage_new uintptr
	addr_QImage_new2 uintptr
	addr_QImage_new3 uintptr
	addr_QImage_new4 uintptr
	addr_QImage_new5 uintptr
	addr_QImage_new6 uintptr
	addr_QImage_new7 uintptr
	addr_QImage_new8 uintptr
	addr_QImage_new9 uintptr
	addr_QImage_new10 uintptr
	addr_QImage_OperatorAssign uintptr
	addr_QImage_swap uintptr
	addr_QImage_isNull uintptr
	addr_QImage_OperatorEqual uintptr
	addr_QImage_OperatorNotEqual uintptr
	addr_QImage_ToQVariant uintptr
	addr_QImage_detach uintptr
	addr_QImage_isDetached uintptr
	addr_QImage_copy uintptr
	addr_QImage_copy2 uintptr
	addr_QImage_format uintptr
	addr_QImage_convertToFormat uintptr
	addr_QImage_convertToFormat2 uintptr
	addr_QImage_convertToFormat3 uintptr
	addr_QImage_reinterpretAsFormat uintptr
	addr_QImage_convertedTo uintptr
	addr_QImage_convertedTo2 uintptr
	addr_QImage_convertTo uintptr
	addr_QImage_width uintptr
	addr_QImage_height uintptr
	addr_QImage_size uintptr
	addr_QImage_rect uintptr
	addr_QImage_depth uintptr
	addr_QImage_colorCount uintptr
	addr_QImage_bitPlaneCount uintptr
	addr_QImage_color uintptr
	addr_QImage_setColor uintptr
	addr_QImage_setColorCount uintptr
	addr_QImage_allGray uintptr
	addr_QImage_isGrayscale uintptr
	addr_QImage_sizeInBytes uintptr
	addr_QImage_bytesPerLine uintptr
	addr_QImage_valid uintptr
	addr_QImage_valid2 uintptr
	addr_QImage_pixelIndex uintptr
	addr_QImage_pixelIndex2 uintptr
	addr_QImage_pixel uintptr
	addr_QImage_pixel2 uintptr
	addr_QImage_setPixel uintptr
	addr_QImage_setPixel2 uintptr
	addr_QImage_pixelColor uintptr
	addr_QImage_pixelColor2 uintptr
	addr_QImage_setPixelColor uintptr
	addr_QImage_setPixelColor2 uintptr
	addr_QImage_colorTable uintptr
	addr_QImage_setColorTable uintptr
	pg_QImage_devicePixelRatio func(uintptr) float64
	pg_QImage_setDevicePixelRatio func(uintptr, float64)
	addr_QImage_deviceIndependentSize uintptr
	addr_QImage_fill uintptr
	addr_QImage_fill2 uintptr
	addr_QImage_fill3 uintptr
	addr_QImage_hasAlphaChannel uintptr
	addr_QImage_setAlphaChannel uintptr
	addr_QImage_createAlphaMask uintptr
	addr_QImage_createHeuristicMask uintptr
	addr_QImage_createMaskFromColor uintptr
	addr_QImage_scaled uintptr
	addr_QImage_scaled2 uintptr
	addr_QImage_scaledToWidth uintptr
	addr_QImage_scaledToHeight uintptr
	addr_QImage_transformed uintptr
	addr_QImage_trueMatrix uintptr
	addr_QImage_mirrored uintptr
	addr_QImage_mirrored2 uintptr
	addr_QImage_mirror uintptr
	addr_QImage_rgbSwapped uintptr
	addr_QImage_rgbSwapped2 uintptr
	addr_QImage_flipped uintptr
	addr_QImage_flipped2 uintptr
	addr_QImage_flip uintptr
	addr_QImage_rgbSwap uintptr
	addr_QImage_invertPixels uintptr
	addr_QImage_colorSpace uintptr
	addr_QImage_convertedToColorSpace uintptr
	addr_QImage_convertedToColorSpace2 uintptr
	addr_QImage_convertedToColorSpace3 uintptr
	addr_QImage_convertToColorSpace uintptr
	addr_QImage_convertToColorSpace2 uintptr
	addr_QImage_setColorSpace uintptr
	addr_QImage_colorTransformed uintptr
	addr_QImage_colorTransformed2 uintptr
	addr_QImage_colorTransformed3 uintptr
	addr_QImage_colorTransformed4 uintptr
	addr_QImage_applyColorTransform uintptr
	addr_QImage_applyColorTransform2 uintptr
	addr_QImage_load uintptr
	addr_QImage_load2 uintptr
	addr_QImage_loadFromData uintptr
	addr_QImage_loadFromData2 uintptr
	addr_QImage_loadFromData3 uintptr
	addr_QImage_save uintptr
	addr_QImage_save2 uintptr
	addr_QImage_fromData uintptr
	addr_QImage_fromData2 uintptr
	addr_QImage_fromData3 uintptr
	addr_QImage_cacheKey uintptr
	addr_QImage_dotsPerMeterX uintptr
	addr_QImage_dotsPerMeterY uintptr
	addr_QImage_setDotsPerMeterX uintptr
	addr_QImage_setDotsPerMeterY uintptr
	addr_QImage_offset uintptr
	addr_QImage_setOffset uintptr
	addr_QImage_textKeys uintptr
	addr_QImage_text uintptr
	addr_QImage_setText uintptr
	addr_QImage_pixelFormat uintptr
	addr_QImage_toPixelFormat uintptr
	addr_QImage_toImageFormat uintptr
	addr_QImage_mirrored_helper uintptr
	addr_QImage_rgbSwapped_helper uintptr
	addr_QImage_mirrored_inplace uintptr
	addr_QImage_rgbSwapped_inplace uintptr
	addr_QImage_convertToFormat_helper uintptr
	addr_QImage_convertToFormat_inplace uintptr
	addr_QImage_smoothScaled uintptr
	addr_QImage_detachMetadata uintptr
	pg_QImage_devicePixelRatioFScale func() float64
	pg_QImage_encodeMetricF func(int32, float64) int32
	pg_QImage_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QImage_copy3 uintptr
	addr_QImage_convertToFormat4 uintptr
	addr_QImage_convertToFormat5 uintptr
	addr_QImage_convertToFormat6 uintptr
	addr_QImage_convertedTo3 uintptr
	addr_QImage_convertedTo4 uintptr
	addr_QImage_convertTo2 uintptr
	addr_QImage_createAlphaMask2 uintptr
	addr_QImage_createHeuristicMask2 uintptr
	addr_QImage_createMaskFromColor2 uintptr
	addr_QImage_scaled3 uintptr
	addr_QImage_scaled4 uintptr
	addr_QImage_scaled5 uintptr
	addr_QImage_scaled6 uintptr
	addr_QImage_scaledToWidth2 uintptr
	addr_QImage_scaledToHeight2 uintptr
	addr_QImage_transformed2 uintptr
	addr_QImage_mirrored3 uintptr
	addr_QImage_mirrored4 uintptr
	addr_QImage_mirrored5 uintptr
	addr_QImage_mirrored6 uintptr
	addr_QImage_mirror2 uintptr
	addr_QImage_mirror3 uintptr
	addr_QImage_flipped3 uintptr
	addr_QImage_flipped4 uintptr
	addr_QImage_flip2 uintptr
	addr_QImage_invertPixels2 uintptr
	addr_QImage_convertedToColorSpace4 uintptr
	addr_QImage_convertedToColorSpace5 uintptr
	addr_QImage_convertToColorSpace3 uintptr
	addr_QImage_colorTransformed5 uintptr
	addr_QImage_colorTransformed6 uintptr
	addr_QImage_applyColorTransform3 uintptr
	addr_QImage_load3 uintptr
	addr_QImage_loadFromData4 uintptr
	addr_QImage_loadFromData5 uintptr
	addr_QImage_loadFromData6 uintptr
	addr_QImage_save3 uintptr
	addr_QImage_save4 uintptr
	addr_QImage_save5 uintptr
	addr_QImage_save6 uintptr
	addr_QImage_fromData4 uintptr
	addr_QImage_fromData5 uintptr
	addr_QImage_fromData6 uintptr
	addr_QImage_text2 uintptr
	addr_QImage_detachMetadata2 uintptr
	addr_QImage_toQVariant uintptr
	addr_QImage_devType_default uintptr
	addr_QImage_paintEngine_default uintptr
	addr_QImage_metric_default uintptr
	addr_QImage_initPainter_default uintptr
	addr_QImage_redirected_default uintptr
	addr_QImage_sharedPainter_default uintptr
	pg_QImage__set_override func(uintptr, int32)
	pg_QImage_delete func(uintptr)
)

func init() {
	reg(&pg_QImage_base, "QImage_base")
	symaddr(&addr_QImage_new, "QImage_new")
	symaddr(&addr_QImage_new2, "QImage_new2")
	symaddr(&addr_QImage_new3, "QImage_new3")
	symaddr(&addr_QImage_new4, "QImage_new4")
	symaddr(&addr_QImage_new5, "QImage_new5")
	symaddr(&addr_QImage_new6, "QImage_new6")
	symaddr(&addr_QImage_new7, "QImage_new7")
	symaddr(&addr_QImage_new8, "QImage_new8")
	symaddr(&addr_QImage_new9, "QImage_new9")
	symaddr(&addr_QImage_new10, "QImage_new10")
	symaddr(&addr_QImage_OperatorAssign, "QImage_OperatorAssign")
	symaddr(&addr_QImage_swap, "QImage_swap")
	symaddr(&addr_QImage_isNull, "QImage_isNull")
	symaddr(&addr_QImage_OperatorEqual, "QImage_OperatorEqual")
	symaddr(&addr_QImage_OperatorNotEqual, "QImage_OperatorNotEqual")
	symaddr(&addr_QImage_ToQVariant, "QImage_ToQVariant")
	symaddr(&addr_QImage_detach, "QImage_detach")
	symaddr(&addr_QImage_isDetached, "QImage_isDetached")
	symaddr(&addr_QImage_copy, "QImage_copy")
	symaddr(&addr_QImage_copy2, "QImage_copy2")
	symaddr(&addr_QImage_format, "QImage_format")
	symaddr(&addr_QImage_convertToFormat, "QImage_convertToFormat")
	symaddr(&addr_QImage_convertToFormat2, "QImage_convertToFormat2")
	symaddr(&addr_QImage_convertToFormat3, "QImage_convertToFormat3")
	symaddr(&addr_QImage_reinterpretAsFormat, "QImage_reinterpretAsFormat")
	symaddr(&addr_QImage_convertedTo, "QImage_convertedTo")
	symaddr(&addr_QImage_convertedTo2, "QImage_convertedTo2")
	symaddr(&addr_QImage_convertTo, "QImage_convertTo")
	symaddr(&addr_QImage_width, "QImage_width")
	symaddr(&addr_QImage_height, "QImage_height")
	symaddr(&addr_QImage_size, "QImage_size")
	symaddr(&addr_QImage_rect, "QImage_rect")
	symaddr(&addr_QImage_depth, "QImage_depth")
	symaddr(&addr_QImage_colorCount, "QImage_colorCount")
	symaddr(&addr_QImage_bitPlaneCount, "QImage_bitPlaneCount")
	symaddr(&addr_QImage_color, "QImage_color")
	symaddr(&addr_QImage_setColor, "QImage_setColor")
	symaddr(&addr_QImage_setColorCount, "QImage_setColorCount")
	symaddr(&addr_QImage_allGray, "QImage_allGray")
	symaddr(&addr_QImage_isGrayscale, "QImage_isGrayscale")
	symaddr(&addr_QImage_sizeInBytes, "QImage_sizeInBytes")
	symaddr(&addr_QImage_bytesPerLine, "QImage_bytesPerLine")
	symaddr(&addr_QImage_valid, "QImage_valid")
	symaddr(&addr_QImage_valid2, "QImage_valid2")
	symaddr(&addr_QImage_pixelIndex, "QImage_pixelIndex")
	symaddr(&addr_QImage_pixelIndex2, "QImage_pixelIndex2")
	symaddr(&addr_QImage_pixel, "QImage_pixel")
	symaddr(&addr_QImage_pixel2, "QImage_pixel2")
	symaddr(&addr_QImage_setPixel, "QImage_setPixel")
	symaddr(&addr_QImage_setPixel2, "QImage_setPixel2")
	symaddr(&addr_QImage_pixelColor, "QImage_pixelColor")
	symaddr(&addr_QImage_pixelColor2, "QImage_pixelColor2")
	symaddr(&addr_QImage_setPixelColor, "QImage_setPixelColor")
	symaddr(&addr_QImage_setPixelColor2, "QImage_setPixelColor2")
	symaddr(&addr_QImage_colorTable, "QImage_colorTable")
	symaddr(&addr_QImage_setColorTable, "QImage_setColorTable")
	reg(&pg_QImage_devicePixelRatio, "QImage_devicePixelRatio")
	reg(&pg_QImage_setDevicePixelRatio, "QImage_setDevicePixelRatio")
	symaddr(&addr_QImage_deviceIndependentSize, "QImage_deviceIndependentSize")
	symaddr(&addr_QImage_fill, "QImage_fill")
	symaddr(&addr_QImage_fill2, "QImage_fill2")
	symaddr(&addr_QImage_fill3, "QImage_fill3")
	symaddr(&addr_QImage_hasAlphaChannel, "QImage_hasAlphaChannel")
	symaddr(&addr_QImage_setAlphaChannel, "QImage_setAlphaChannel")
	symaddr(&addr_QImage_createAlphaMask, "QImage_createAlphaMask")
	symaddr(&addr_QImage_createHeuristicMask, "QImage_createHeuristicMask")
	symaddr(&addr_QImage_createMaskFromColor, "QImage_createMaskFromColor")
	symaddr(&addr_QImage_scaled, "QImage_scaled")
	symaddr(&addr_QImage_scaled2, "QImage_scaled2")
	symaddr(&addr_QImage_scaledToWidth, "QImage_scaledToWidth")
	symaddr(&addr_QImage_scaledToHeight, "QImage_scaledToHeight")
	symaddr(&addr_QImage_transformed, "QImage_transformed")
	symaddr(&addr_QImage_trueMatrix, "QImage_trueMatrix")
	symaddr(&addr_QImage_mirrored, "QImage_mirrored")
	symaddr(&addr_QImage_mirrored2, "QImage_mirrored2")
	symaddr(&addr_QImage_mirror, "QImage_mirror")
	symaddr(&addr_QImage_rgbSwapped, "QImage_rgbSwapped")
	symaddr(&addr_QImage_rgbSwapped2, "QImage_rgbSwapped2")
	symaddr(&addr_QImage_flipped, "QImage_flipped")
	symaddr(&addr_QImage_flipped2, "QImage_flipped2")
	symaddr(&addr_QImage_flip, "QImage_flip")
	symaddr(&addr_QImage_rgbSwap, "QImage_rgbSwap")
	symaddr(&addr_QImage_invertPixels, "QImage_invertPixels")
	symaddr(&addr_QImage_colorSpace, "QImage_colorSpace")
	symaddr(&addr_QImage_convertedToColorSpace, "QImage_convertedToColorSpace")
	symaddr(&addr_QImage_convertedToColorSpace2, "QImage_convertedToColorSpace2")
	symaddr(&addr_QImage_convertedToColorSpace3, "QImage_convertedToColorSpace3")
	symaddr(&addr_QImage_convertToColorSpace, "QImage_convertToColorSpace")
	symaddr(&addr_QImage_convertToColorSpace2, "QImage_convertToColorSpace2")
	symaddr(&addr_QImage_setColorSpace, "QImage_setColorSpace")
	symaddr(&addr_QImage_colorTransformed, "QImage_colorTransformed")
	symaddr(&addr_QImage_colorTransformed2, "QImage_colorTransformed2")
	symaddr(&addr_QImage_colorTransformed3, "QImage_colorTransformed3")
	symaddr(&addr_QImage_colorTransformed4, "QImage_colorTransformed4")
	symaddr(&addr_QImage_applyColorTransform, "QImage_applyColorTransform")
	symaddr(&addr_QImage_applyColorTransform2, "QImage_applyColorTransform2")
	symaddr(&addr_QImage_load, "QImage_load")
	symaddr(&addr_QImage_load2, "QImage_load2")
	symaddr(&addr_QImage_loadFromData, "QImage_loadFromData")
	symaddr(&addr_QImage_loadFromData2, "QImage_loadFromData2")
	symaddr(&addr_QImage_loadFromData3, "QImage_loadFromData3")
	symaddr(&addr_QImage_save, "QImage_save")
	symaddr(&addr_QImage_save2, "QImage_save2")
	symaddr(&addr_QImage_fromData, "QImage_fromData")
	symaddr(&addr_QImage_fromData2, "QImage_fromData2")
	symaddr(&addr_QImage_fromData3, "QImage_fromData3")
	symaddr(&addr_QImage_cacheKey, "QImage_cacheKey")
	symaddr(&addr_QImage_dotsPerMeterX, "QImage_dotsPerMeterX")
	symaddr(&addr_QImage_dotsPerMeterY, "QImage_dotsPerMeterY")
	symaddr(&addr_QImage_setDotsPerMeterX, "QImage_setDotsPerMeterX")
	symaddr(&addr_QImage_setDotsPerMeterY, "QImage_setDotsPerMeterY")
	symaddr(&addr_QImage_offset, "QImage_offset")
	symaddr(&addr_QImage_setOffset, "QImage_setOffset")
	symaddr(&addr_QImage_textKeys, "QImage_textKeys")
	symaddr(&addr_QImage_text, "QImage_text")
	symaddr(&addr_QImage_setText, "QImage_setText")
	symaddr(&addr_QImage_pixelFormat, "QImage_pixelFormat")
	symaddr(&addr_QImage_toPixelFormat, "QImage_toPixelFormat")
	symaddr(&addr_QImage_toImageFormat, "QImage_toImageFormat")
	symaddr(&addr_QImage_mirrored_helper, "QImage_mirrored_helper")
	symaddr(&addr_QImage_rgbSwapped_helper, "QImage_rgbSwapped_helper")
	symaddr(&addr_QImage_mirrored_inplace, "QImage_mirrored_inplace")
	symaddr(&addr_QImage_rgbSwapped_inplace, "QImage_rgbSwapped_inplace")
	symaddr(&addr_QImage_convertToFormat_helper, "QImage_convertToFormat_helper")
	symaddr(&addr_QImage_convertToFormat_inplace, "QImage_convertToFormat_inplace")
	symaddr(&addr_QImage_smoothScaled, "QImage_smoothScaled")
	symaddr(&addr_QImage_detachMetadata, "QImage_detachMetadata")
	reg(&pg_QImage_devicePixelRatioFScale, "QImage_devicePixelRatioFScale")
	reg(&pg_QImage_encodeMetricF, "QImage_encodeMetricF")
	reg(&pg_QImage_getDecodedMetricF, "QImage_getDecodedMetricF")
	symaddr(&addr_QImage_copy3, "QImage_copy3")
	symaddr(&addr_QImage_convertToFormat4, "QImage_convertToFormat4")
	symaddr(&addr_QImage_convertToFormat5, "QImage_convertToFormat5")
	symaddr(&addr_QImage_convertToFormat6, "QImage_convertToFormat6")
	symaddr(&addr_QImage_convertedTo3, "QImage_convertedTo3")
	symaddr(&addr_QImage_convertedTo4, "QImage_convertedTo4")
	symaddr(&addr_QImage_convertTo2, "QImage_convertTo2")
	symaddr(&addr_QImage_createAlphaMask2, "QImage_createAlphaMask2")
	symaddr(&addr_QImage_createHeuristicMask2, "QImage_createHeuristicMask2")
	symaddr(&addr_QImage_createMaskFromColor2, "QImage_createMaskFromColor2")
	symaddr(&addr_QImage_scaled3, "QImage_scaled3")
	symaddr(&addr_QImage_scaled4, "QImage_scaled4")
	symaddr(&addr_QImage_scaled5, "QImage_scaled5")
	symaddr(&addr_QImage_scaled6, "QImage_scaled6")
	symaddr(&addr_QImage_scaledToWidth2, "QImage_scaledToWidth2")
	symaddr(&addr_QImage_scaledToHeight2, "QImage_scaledToHeight2")
	symaddr(&addr_QImage_transformed2, "QImage_transformed2")
	symaddr(&addr_QImage_mirrored3, "QImage_mirrored3")
	symaddr(&addr_QImage_mirrored4, "QImage_mirrored4")
	symaddr(&addr_QImage_mirrored5, "QImage_mirrored5")
	symaddr(&addr_QImage_mirrored6, "QImage_mirrored6")
	symaddr(&addr_QImage_mirror2, "QImage_mirror2")
	symaddr(&addr_QImage_mirror3, "QImage_mirror3")
	symaddr(&addr_QImage_flipped3, "QImage_flipped3")
	symaddr(&addr_QImage_flipped4, "QImage_flipped4")
	symaddr(&addr_QImage_flip2, "QImage_flip2")
	symaddr(&addr_QImage_invertPixels2, "QImage_invertPixels2")
	symaddr(&addr_QImage_convertedToColorSpace4, "QImage_convertedToColorSpace4")
	symaddr(&addr_QImage_convertedToColorSpace5, "QImage_convertedToColorSpace5")
	symaddr(&addr_QImage_convertToColorSpace3, "QImage_convertToColorSpace3")
	symaddr(&addr_QImage_colorTransformed5, "QImage_colorTransformed5")
	symaddr(&addr_QImage_colorTransformed6, "QImage_colorTransformed6")
	symaddr(&addr_QImage_applyColorTransform3, "QImage_applyColorTransform3")
	symaddr(&addr_QImage_load3, "QImage_load3")
	symaddr(&addr_QImage_loadFromData4, "QImage_loadFromData4")
	symaddr(&addr_QImage_loadFromData5, "QImage_loadFromData5")
	symaddr(&addr_QImage_loadFromData6, "QImage_loadFromData6")
	symaddr(&addr_QImage_save3, "QImage_save3")
	symaddr(&addr_QImage_save4, "QImage_save4")
	symaddr(&addr_QImage_save5, "QImage_save5")
	symaddr(&addr_QImage_save6, "QImage_save6")
	symaddr(&addr_QImage_fromData4, "QImage_fromData4")
	symaddr(&addr_QImage_fromData5, "QImage_fromData5")
	symaddr(&addr_QImage_fromData6, "QImage_fromData6")
	symaddr(&addr_QImage_text2, "QImage_text2")
	symaddr(&addr_QImage_detachMetadata2, "QImage_detachMetadata2")
	symaddr(&addr_QImage_toQVariant, "QImage_toQVariant")
	symaddr(&addr_QImage_devType_default, "QImage_devType_default")
	symaddr(&addr_QImage_paintEngine_default, "QImage_paintEngine_default")
	symaddr(&addr_QImage_metric_default, "QImage_metric_default")
	symaddr(&addr_QImage_initPainter_default, "QImage_initPainter_default")
	symaddr(&addr_QImage_redirected_default, "QImage_redirected_default")
	symaddr(&addr_QImage_sharedPainter_default, "QImage_sharedPainter_default")
	reg(&pg_QImage__set_override, "QImage__set_override")
	reg(&pg_QImage_delete, "QImage_delete")
}

type QImage struct {
	h uintptr
	*QPaintDevice
}

func bindQImage(h uintptr) *QImage {
	if h == 0 {
		return nil
	}
	return &QImage{h: h, QPaintDevice: bindQPaintDevice(pg_QImage_base(h))}
}

func newownedQImage(h uintptr) *QImage {
	o := bindQImage(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QImage).Delete)
	}
	return o
}

func (this *QImage) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QImage) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQImage(h unsafe.Pointer) *QImage { return bindQImage(uintptr(h)) }
func UnsafeNewQImageOwned(h unsafe.Pointer) *QImage { return newownedQImage(uintptr(h)) }

func (this *QImage) UnsafeRelease() {
	this.h = 0
	if this.QPaintDevice != nil {
		this.QPaintDevice.UnsafeRelease()
	}
}

func NewQImage() *QImage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QImage_new, uintptr(__goid))
	__o := newownedQImage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage2(size *QSize, format QImage__Format) *QImage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QImage_new2, uintptr(__goid), uintptr(size.UnsafePointer()), uintptr(format))
	__o := newownedQImage(__r)
	runtime.KeepAlive(size)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage3(width int, height int, format QImage__Format) *QImage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QImage_new3, uintptr(__goid), uintptr(width), uintptr(height), uintptr(format))
	__o := newownedQImage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage4(width int, height int, format QImage__Format) *QImage {
	__goid := nextVirtID()
	var data_out uint8
	__r, _, _ := purego.SyscallN(addr_QImage_new4, uintptr(__goid), uintptr(unsafe.Pointer(&data_out)), uintptr(width), uintptr(height), uintptr(format))
	__o := newownedQImage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage5(data []byte, width int, height int, format QImage__Format) *QImage {
	__goid := nextVirtID()
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_new5, uintptr(__goid), data_p, uintptr(width), uintptr(height), uintptr(format))
	__o := newownedQImage(__r)
	runtime.KeepAlive(data)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage6(width int, height int, bytesPerLine int64, format QImage__Format) *QImage {
	__goid := nextVirtID()
	var data_out uint8
	__r, _, _ := purego.SyscallN(addr_QImage_new6, uintptr(__goid), uintptr(unsafe.Pointer(&data_out)), uintptr(width), uintptr(height), uintptr(bytesPerLine), uintptr(format))
	__o := newownedQImage(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage7(data []byte, width int, height int, bytesPerLine int64, format QImage__Format) *QImage {
	__goid := nextVirtID()
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_new7, uintptr(__goid), data_p, uintptr(width), uintptr(height), uintptr(bytesPerLine), uintptr(format))
	__o := newownedQImage(__r)
	runtime.KeepAlive(data)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage8(fileName string) *QImage {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_new8, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQImage(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage9(arg1 *QImage) *QImage {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QImage_new9, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQImage(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func NewQImage10(fileName string, format string) *QImage {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QImage_new10, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__o := newownedQImage(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QImage_virtIdx, QImage_enable)
	return __o
}

func (this *QImage) OperatorAssign(arg1 *QImage) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQImage(__r)
}

func (this *QImage) Swap(other *QImage) {
	purego.SyscallN(addr_QImage_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QImage) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) OperatorEqual(arg1 *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QImage) OperatorNotEqual(arg1 *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QImage) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QImage_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QImage) Detach() {
	purego.SyscallN(addr_QImage_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) Copy() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_copy, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Copy2(x int, y int, w int, h int) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_copy2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Format() QImage__Format {
	__r0, _, _ := purego.SyscallN(addr_QImage_format, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImage__Format(__r)
}

func (this *QImage) ConvertToFormat(f QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormatWithQImageFormat(f QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat2, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormat2(f QImage__Format, colorTable []uint) *QImage {
	colorTable_arr := make([]int32, len(colorTable))
	for __i, __e := range colorTable {
		colorTable_arr[__i] = int32(__e)
	}
	var colorTable_p uintptr
	if len(colorTable_arr) > 0 {
		colorTable_p = uintptr(unsafe.Pointer(&colorTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat3, this.h, uintptr(f), colorTable_p, uintptr(len(colorTable)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorTable_arr)
	return newownedQImage(__r)
}

func (this *QImage) ReinterpretAsFormat(f QImage__Format) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_reinterpretAsFormat, this.h, uintptr(f))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) ConvertedTo(f QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedTo, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedToWithQImageFormat(f QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedTo2, this.h, uintptr(f))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertTo(f QImage__Format) {
	purego.SyscallN(addr_QImage_convertTo, this.h, uintptr(f))
	runtime.KeepAlive(this)
}

func (this *QImage) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QImage_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QImage) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImage_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QImage) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) ColorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_colorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) BitPlaneCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_bitPlaneCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) Color(i int) uint {
	__r0, _, _ := purego.SyscallN(addr_QImage_color, this.h, uintptr(i))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QImage) SetColor(i int, c uint) {
	purego.SyscallN(addr_QImage_setColor, this.h, uintptr(i), uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QImage) SetColorCount(arg1 int) {
	purego.SyscallN(addr_QImage_setColorCount, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QImage) AllGray() bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_allGray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) IsGrayscale() bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_isGrayscale, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) SizeInBytes() int64 {
	__r0, _, _ := purego.SyscallN(addr_QImage_sizeInBytes, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QImage) BytesPerLine() int64 {
	__r0, _, _ := purego.SyscallN(addr_QImage_bytesPerLine, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QImage) Valid(x int, y int) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_valid, this.h, uintptr(x), uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) ValidWithPt(pt *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_valid2, this.h, uintptr(pt.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return __r != 0
}

func (this *QImage) PixelIndex(x int, y int) int {
	__r0, _, _ := purego.SyscallN(addr_QImage_pixelIndex, this.h, uintptr(x), uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) PixelIndexWithPt(pt *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QImage_pixelIndex2, this.h, uintptr(pt.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return int(__r)
}

func (this *QImage) Pixel(x int, y int) uint {
	__r0, _, _ := purego.SyscallN(addr_QImage_pixel, this.h, uintptr(x), uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QImage) PixelWithPt(pt *QPoint) uint {
	__r0, _, _ := purego.SyscallN(addr_QImage_pixel2, this.h, uintptr(pt.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return uint(__r)
}

func (this *QImage) SetPixel(x int, y int, index_or_rgb uint) {
	purego.SyscallN(addr_QImage_setPixel, this.h, uintptr(x), uintptr(y), uintptr(index_or_rgb))
	runtime.KeepAlive(this)
}

func (this *QImage) SetPixel2(pt *QPoint, index_or_rgb uint) {
	purego.SyscallN(addr_QImage_setPixel2, this.h, uintptr(pt.UnsafePointer()), uintptr(index_or_rgb))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
}

func (this *QImage) PixelColor(x int, y int) *QColor {
	__r, _, _ := purego.SyscallN(addr_QImage_pixelColor, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QImage) PixelColorWithPt(pt *QPoint) *QColor {
	__r, _, _ := purego.SyscallN(addr_QImage_pixelColor2, this.h, uintptr(pt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	return newownedQColor(__r)
}

func (this *QImage) SetPixelColor(x int, y int, c *QColor) {
	purego.SyscallN(addr_QImage_setPixelColor, this.h, uintptr(x), uintptr(y), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
}

func (this *QImage) SetPixelColor2(pt *QPoint, c *QColor) {
	purego.SyscallN(addr_QImage_setPixelColor2, this.h, uintptr(pt.UnsafePointer()), uintptr(c.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
	runtime.KeepAlive(c)
}

func (this *QImage) ColorTable() []uint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImage_colorTable, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]uint, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = uint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QImage) SetColorTable(colors []uint) {
	colors_arr := make([]int32, len(colors))
	for __i, __e := range colors {
		colors_arr[__i] = int32(__e)
	}
	var colors_p uintptr
	if len(colors_arr) > 0 {
		colors_p = uintptr(unsafe.Pointer(&colors_arr[0]))
	}
	purego.SyscallN(addr_QImage_setColorTable, this.h, colors_p, uintptr(len(colors)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colors_arr)
}

func (this *QImage) DevicePixelRatio() float64 {
	__r := pg_QImage_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QImage) SetDevicePixelRatio(scaleFactor float64) {
	pg_QImage_setDevicePixelRatio(this.h, scaleFactor)
	runtime.KeepAlive(this)
}

func (this *QImage) DeviceIndependentSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QImage_deviceIndependentSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QImage) Fill(pixel uint) {
	purego.SyscallN(addr_QImage_fill, this.h, uintptr(pixel))
	runtime.KeepAlive(this)
}

func (this *QImage) FillWithColor(color *QColor) {
	purego.SyscallN(addr_QImage_fill2, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QImage) Fill2(color GlobalColor) {
	purego.SyscallN(addr_QImage_fill3, this.h, uintptr(color))
	runtime.KeepAlive(this)
}

func (this *QImage) HasAlphaChannel() bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_hasAlphaChannel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) SetAlphaChannel(alphaChannel *QImage) {
	purego.SyscallN(addr_QImage_setAlphaChannel, this.h, uintptr(alphaChannel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(alphaChannel)
}

func (this *QImage) CreateAlphaMask() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createAlphaMask, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) CreateHeuristicMask() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createHeuristicMask, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) CreateMaskFromColor(color uint) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createMaskFromColor, this.h, uintptr(color))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Scaled(w int, h int) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ScaledWithQSize(s *QSize) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled2, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQImage(__r)
}

func (this *QImage) ScaledToWidth(w int) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaledToWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ScaledToHeight(h int) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaledToHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Transformed(matrix *QTransform) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_transformed, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQImage(__r)
}

func QImage_TrueMatrix(arg1 *QTransform, w int, h int) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QImage_trueMatrix, uintptr(arg1.UnsafePointer()), uintptr(w), uintptr(h))
	runtime.KeepAlive(arg1)
	return newownedQTransform(__r)
}

func (this *QImage) Mirrored() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirrored2() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored2, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirror() {
	purego.SyscallN(addr_QImage_mirror, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) RgbSwapped() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_rgbSwapped, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) RgbSwapped2() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_rgbSwapped2, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Flipped() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_flipped, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Flipped2() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_flipped2, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Flip() {
	purego.SyscallN(addr_QImage_flip, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) RgbSwap() {
	purego.SyscallN(addr_QImage_rgbSwap, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) InvertPixels() {
	purego.SyscallN(addr_QImage_invertPixels, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) ColorSpace() *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QImage_colorSpace, this.h)
	runtime.KeepAlive(this)
	return newownedQColorSpace(__r)
}

func (this *QImage) ConvertedToColorSpace(colorSpace *QColorSpace) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedToColorSpace, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedToColorSpace2(colorSpace *QColorSpace, format QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedToColorSpace2, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedToColorSpace3(colorSpace *QColorSpace, format QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedToColorSpace3, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToColorSpace(colorSpace *QColorSpace) {
	purego.SyscallN(addr_QImage_convertToColorSpace, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QImage) ConvertToColorSpace2(colorSpace *QColorSpace, format QImage__Format) {
	purego.SyscallN(addr_QImage_convertToColorSpace2, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QImage) SetColorSpace(colorSpace *QColorSpace) {
	purego.SyscallN(addr_QImage_setColorSpace, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QImage) ColorTransformed(transform *QColorTransform) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ColorTransformed2(transform *QColorTransform, format QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed2, this.h, uintptr(transform.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ColorTransformedWithTransform(transform *QColorTransform) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed3, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ColorTransformed3(transform *QColorTransform, format QImage__Format) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed4, this.h, uintptr(transform.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ApplyColorTransform(transform *QColorTransform) {
	purego.SyscallN(addr_QImage_applyColorTransform, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QImage) ApplyColorTransform2(transform *QColorTransform, format QImage__Format) {
	purego.SyscallN(addr_QImage_applyColorTransform2, this.h, uintptr(transform.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QImage) Load(device *QIODevice, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_load, this.h, uintptr(device.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) LoadWithFileName(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QImage_load2, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QImage) LoadFromData(data *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData, this.h, uintptr(data.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QImage) LoadFromData2(buf []byte, len_ int) bool {
	var buf_p uintptr
	if len(buf) > 0 {
		buf_p = uintptr(unsafe.Pointer(&buf[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData2, this.h, buf_p, uintptr(len_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(buf)
	return __r != 0
}

func (this *QImage) LoadFromDataWithData(data *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData3, this.h, uintptr(data.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func (this *QImage) Save(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QImage_save, this.h, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QImage) SaveWithDevice(device *QIODevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_save2, this.h, uintptr(device.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	return __r != 0
}

func QImage_FromData(data *QByteArrayView) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_fromData, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(data)
	return newownedQImage(__r)
}

func QImage_FromData2(data []byte, size int) *QImage {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_fromData2, data_p, uintptr(size))
	runtime.KeepAlive(data)
	return newownedQImage(__r)
}

func QImage_FromDataWithData(data *QByteArray) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_fromData3, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(data)
	return newownedQImage(__r)
}

func (this *QImage) CacheKey() int64 {
	__r0, _, _ := purego.SyscallN(addr_QImage_cacheKey, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QImage) DotsPerMeterX() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_dotsPerMeterX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) DotsPerMeterY() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_dotsPerMeterY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) SetDotsPerMeterX(arg1 int) {
	purego.SyscallN(addr_QImage_setDotsPerMeterX, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QImage) SetDotsPerMeterY(arg1 int) {
	purego.SyscallN(addr_QImage_setDotsPerMeterY, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QImage) Offset() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QImage_offset, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QImage) SetOffset(arg1 *QPoint) {
	purego.SyscallN(addr_QImage_setOffset, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QImage) TextKeys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImage_textKeys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QImage) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImage_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QImage) SetText(key string, value string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QImage_setText, this.h, key_p, uintptr(len(key)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)
}

func (this *QImage) PixelFormat() *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QImage_pixelFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQPixelFormat(__r)
}

func QImage_ToPixelFormat(format QImage__Format) *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QImage_toPixelFormat, uintptr(format))
	return newownedQPixelFormat(__r)
}

func QImage_ToImageFormat(format *QPixelFormat) QImage__Format {
	__r0, _, _ := purego.SyscallN(addr_QImage_toImageFormat, uintptr(format.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(format)
	return QImage__Format(__r)
}

func (this *QImage) Mirrored_helper(horizontal bool, vertical bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored_helper, this.h, uintptr(b2i(horizontal)), uintptr(b2i(vertical)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) RgbSwapped_helper() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_rgbSwapped_helper, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirrored_inplace(horizontal bool, vertical bool) {
	purego.SyscallN(addr_QImage_mirrored_inplace, this.h, uintptr(b2i(horizontal)), uintptr(b2i(vertical)))
	runtime.KeepAlive(this)
}

func (this *QImage) RgbSwapped_inplace() {
	purego.SyscallN(addr_QImage_rgbSwapped_inplace, this.h)
	runtime.KeepAlive(this)
}

func (this *QImage) ConvertToFormat_helper(format QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat_helper, this.h, uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormat_inplace(format QImage__Format, flags ImageConversionFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QImage_convertToFormat_inplace, this.h, uintptr(format), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImage) SmoothScaled(w int, h int) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_smoothScaled, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) DetachMetadata() {
	purego.SyscallN(addr_QImage_detachMetadata, this.h)
	runtime.KeepAlive(this)
}

func QImage_DevicePixelRatioFScale() float64 {
	__r := pg_QImage_devicePixelRatioFScale()
	return __r
}

func QImage_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QImage_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QImage) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QImage_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func (this *QImage) CopyWithRect(rect *QRect) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_copy3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormat3(f QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat4, this.h, uintptr(f), uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormat4(f QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat5, this.h, uintptr(f), uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToFormat5(f QImage__Format, colorTable []uint, flags ImageConversionFlag) *QImage {
	colorTable_arr := make([]int32, len(colorTable))
	for __i, __e := range colorTable {
		colorTable_arr[__i] = int32(__e)
	}
	var colorTable_p uintptr
	if len(colorTable_arr) > 0 {
		colorTable_p = uintptr(unsafe.Pointer(&colorTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QImage_convertToFormat6, this.h, uintptr(f), colorTable_p, uintptr(len(colorTable)), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorTable_arr)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedTo2(f QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedTo3, this.h, uintptr(f), uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedTo3(f QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedTo4, this.h, uintptr(f), uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ConvertTo2(f QImage__Format, flags ImageConversionFlag) {
	purego.SyscallN(addr_QImage_convertTo2, this.h, uintptr(f), uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QImage) CreateAlphaMaskWithFlags(flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createAlphaMask2, this.h, uintptr(flags))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) CreateHeuristicMaskWithClipTight(clipTight bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createHeuristicMask2, this.h, uintptr(b2i(clipTight)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) CreateMaskFromColor2(color uint, mode MaskMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_createMaskFromColor2, this.h, uintptr(color), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Scaled2(w int, h int, aspectMode AspectRatioMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled3, this.h, uintptr(w), uintptr(h), uintptr(aspectMode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Scaled3(w int, h int, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled4, this.h, uintptr(w), uintptr(h), uintptr(aspectMode), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Scaled4(s *QSize, aspectMode AspectRatioMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled5, this.h, uintptr(s.UnsafePointer()), uintptr(aspectMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQImage(__r)
}

func (this *QImage) Scaled5(s *QSize, aspectMode AspectRatioMode, mode TransformationMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaled6, this.h, uintptr(s.UnsafePointer()), uintptr(aspectMode), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return newownedQImage(__r)
}

func (this *QImage) ScaledToWidth2(w int, mode TransformationMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaledToWidth2, this.h, uintptr(w), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) ScaledToHeight2(h int, mode TransformationMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_scaledToHeight2, this.h, uintptr(h), uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Transformed2(matrix *QTransform, mode TransformationMode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_transformed2, this.h, uintptr(matrix.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
	return newownedQImage(__r)
}

func (this *QImage) MirroredWithHorizontally(horizontally bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored3, this.h, uintptr(b2i(horizontally)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirrored3(horizontally bool, vertically bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored4, this.h, uintptr(b2i(horizontally)), uintptr(b2i(vertically)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirrored4(horizontally bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored5, this.h, uintptr(b2i(horizontally)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Mirrored5(horizontally bool, vertically bool) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_mirrored6, this.h, uintptr(b2i(horizontally)), uintptr(b2i(vertically)))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) MirrorWithHorizontally(horizontally bool) {
	purego.SyscallN(addr_QImage_mirror2, this.h, uintptr(b2i(horizontally)))
	runtime.KeepAlive(this)
}

func (this *QImage) Mirror2(horizontally bool, vertically bool) {
	purego.SyscallN(addr_QImage_mirror3, this.h, uintptr(b2i(horizontally)), uintptr(b2i(vertically)))
	runtime.KeepAlive(this)
}

func (this *QImage) FlippedWithOrient(orient Orientation) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_flipped3, this.h, uintptr(orient))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) Flipped3(orient Orientation) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_flipped4, this.h, uintptr(orient))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImage) FlipWithOrient(orient Orientation) {
	purego.SyscallN(addr_QImage_flip2, this.h, uintptr(orient))
	runtime.KeepAlive(this)
}

func (this *QImage) InvertPixelsWithQImageInvertMode(arg1 QImage__InvertMode) {
	purego.SyscallN(addr_QImage_invertPixels2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QImage) ConvertedToColorSpace4(colorSpace *QColorSpace, format QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedToColorSpace4, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQImage(__r)
}

func (this *QImage) ConvertedToColorSpace5(colorSpace *QColorSpace, format QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_convertedToColorSpace5, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return newownedQImage(__r)
}

func (this *QImage) ConvertToColorSpace3(colorSpace *QColorSpace, format QImage__Format, flags ImageConversionFlag) {
	purego.SyscallN(addr_QImage_convertToColorSpace3, this.h, uintptr(colorSpace.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QImage) ColorTransformed4(transform *QColorTransform, format QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed5, this.h, uintptr(transform.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ColorTransformed5(transform *QColorTransform, format QImage__Format, flags ImageConversionFlag) *QImage {
	__r, _, _ := purego.SyscallN(addr_QImage_colorTransformed6, this.h, uintptr(transform.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
	return newownedQImage(__r)
}

func (this *QImage) ApplyColorTransform3(transform *QColorTransform, format QImage__Format, flags ImageConversionFlag) {
	purego.SyscallN(addr_QImage_applyColorTransform3, this.h, uintptr(transform.UnsafePointer()), uintptr(format), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QImage) Load2(fileName string, format string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_load3, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) LoadFromData3(data *QByteArrayView, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData4, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) LoadFromData4(buf []byte, len_ int, format string) bool {
	var buf_p uintptr
	if len(buf) > 0 {
		buf_p = uintptr(unsafe.Pointer(&buf[0]))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData5, this.h, buf_p, uintptr(len_), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) LoadFromData5(data *QByteArray, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_loadFromData6, this.h, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) Save2(fileName string, format string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_save3, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) Save3(fileName string, format string, quality int) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_save4, this.h, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&format_b[0])), uintptr(quality))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) Save4(device *QIODevice, format string) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_save5, this.h, uintptr(device.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func (this *QImage) Save5(device *QIODevice, format string, quality int) bool {
	format_b := append([]byte(format), 0)
	__r0, _, _ := purego.SyscallN(addr_QImage_save6, this.h, uintptr(device.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])), uintptr(quality))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format_b)
	return __r != 0
}

func QImage_FromData3(data *QByteArrayView, format string) *QImage {
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QImage_fromData4, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return newownedQImage(__r)
}

func QImage_FromData4(data []byte, size int, format string) *QImage {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QImage_fromData5, data_p, uintptr(size), uintptr(unsafe.Pointer(&format_b[0])))
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return newownedQImage(__r)
}

func QImage_FromData5(data *QByteArray, format string) *QImage {
	format_b := append([]byte(format), 0)
	__r, _, _ := purego.SyscallN(addr_QImage_fromData6, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&format_b[0])))
	runtime.KeepAlive(data)
	runtime.KeepAlive(format_b)
	return newownedQImage(__r)
}

func (this *QImage) TextWithKey(key string) string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImage_text2, this.h, key_p, uintptr(len(key)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return pgStr(__r, __n)
}

func (this *QImage) DetachMetadataWithInvalidateCache(invalidateCache bool) {
	purego.SyscallN(addr_QImage_detachMetadata2, this.h, uintptr(b2i(invalidateCache)))
	runtime.KeepAlive(this)
}

func (this *QImage) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QImage_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QImage) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QImage_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QImage_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QImage) MetricDefault(metric QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QImage_metric_default, this.h, uintptr(metric))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImage) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QImage_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QImage) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QImage_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QImage) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QImage_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func QImage_virtIdx(slot int64) int32 {
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

func QImage_enable(h uintptr, idx int32) { pg_QImage__set_override(h, idx) }

func (this *QImage) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QImage) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QImage) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(metric int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(metric))) }) // slot "metric"
}

func (this *QImage) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QImage) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QImage) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QImage) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QImage_delete(this.h)
	}
	this.UnsafeRelease()
}
