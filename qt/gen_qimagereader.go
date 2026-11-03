package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QImageReader_new uintptr
	addr_QImageReader_new2 uintptr
	addr_QImageReader_new3 uintptr
	addr_QImageReader_new4 uintptr
	addr_QImageReader_new5 uintptr
	addr_QImageReader_tr uintptr
	addr_QImageReader_setFormat uintptr
	addr_QImageReader_format uintptr
	addr_QImageReader_setAutoDetectImageFormat uintptr
	addr_QImageReader_autoDetectImageFormat uintptr
	addr_QImageReader_setDecideFormatFromContent uintptr
	addr_QImageReader_decideFormatFromContent uintptr
	addr_QImageReader_setDevice uintptr
	addr_QImageReader_device uintptr
	addr_QImageReader_setFileName uintptr
	addr_QImageReader_fileName uintptr
	addr_QImageReader_size uintptr
	addr_QImageReader_imageFormat uintptr
	addr_QImageReader_textKeys uintptr
	addr_QImageReader_text uintptr
	addr_QImageReader_setClipRect uintptr
	addr_QImageReader_clipRect uintptr
	addr_QImageReader_setScaledSize uintptr
	addr_QImageReader_scaledSize uintptr
	addr_QImageReader_setQuality uintptr
	addr_QImageReader_quality uintptr
	addr_QImageReader_setScaledClipRect uintptr
	addr_QImageReader_scaledClipRect uintptr
	addr_QImageReader_setBackgroundColor uintptr
	addr_QImageReader_backgroundColor uintptr
	addr_QImageReader_supportsAnimation uintptr
	addr_QImageReader_transformation uintptr
	addr_QImageReader_setAutoTransform uintptr
	addr_QImageReader_autoTransform uintptr
	addr_QImageReader_subType uintptr
	addr_QImageReader_supportedSubTypes uintptr
	addr_QImageReader_canRead uintptr
	addr_QImageReader_read uintptr
	addr_QImageReader_read2 uintptr
	addr_QImageReader_jumpToNextImage uintptr
	addr_QImageReader_jumpToImage uintptr
	addr_QImageReader_loopCount uintptr
	addr_QImageReader_imageCount uintptr
	addr_QImageReader_nextImageDelay uintptr
	addr_QImageReader_currentImageNumber uintptr
	addr_QImageReader_currentImageRect uintptr
	addr_QImageReader_error uintptr
	addr_QImageReader_errorString uintptr
	addr_QImageReader_supportsOption uintptr
	addr_QImageReader_imageFormat2 uintptr
	addr_QImageReader_imageFormat3 uintptr
	addr_QImageReader_supportedImageFormats uintptr
	addr_QImageReader_supportedMimeTypes uintptr
	addr_QImageReader_imageFormatsForMimeType uintptr
	addr_QImageReader_allocationLimit uintptr
	addr_QImageReader_setAllocationLimit uintptr
	addr_QImageReader_tr2 uintptr
	addr_QImageReader_tr3 uintptr
	pg_QImageReader_delete func(uintptr)
)

func init() {
	symaddr(&addr_QImageReader_new, "QImageReader_new")
	symaddr(&addr_QImageReader_new2, "QImageReader_new2")
	symaddr(&addr_QImageReader_new3, "QImageReader_new3")
	symaddr(&addr_QImageReader_new4, "QImageReader_new4")
	symaddr(&addr_QImageReader_new5, "QImageReader_new5")
	symaddr(&addr_QImageReader_tr, "QImageReader_tr")
	symaddr(&addr_QImageReader_setFormat, "QImageReader_setFormat")
	symaddr(&addr_QImageReader_format, "QImageReader_format")
	symaddr(&addr_QImageReader_setAutoDetectImageFormat, "QImageReader_setAutoDetectImageFormat")
	symaddr(&addr_QImageReader_autoDetectImageFormat, "QImageReader_autoDetectImageFormat")
	symaddr(&addr_QImageReader_setDecideFormatFromContent, "QImageReader_setDecideFormatFromContent")
	symaddr(&addr_QImageReader_decideFormatFromContent, "QImageReader_decideFormatFromContent")
	symaddr(&addr_QImageReader_setDevice, "QImageReader_setDevice")
	symaddr(&addr_QImageReader_device, "QImageReader_device")
	symaddr(&addr_QImageReader_setFileName, "QImageReader_setFileName")
	symaddr(&addr_QImageReader_fileName, "QImageReader_fileName")
	symaddr(&addr_QImageReader_size, "QImageReader_size")
	symaddr(&addr_QImageReader_imageFormat, "QImageReader_imageFormat")
	symaddr(&addr_QImageReader_textKeys, "QImageReader_textKeys")
	symaddr(&addr_QImageReader_text, "QImageReader_text")
	symaddr(&addr_QImageReader_setClipRect, "QImageReader_setClipRect")
	symaddr(&addr_QImageReader_clipRect, "QImageReader_clipRect")
	symaddr(&addr_QImageReader_setScaledSize, "QImageReader_setScaledSize")
	symaddr(&addr_QImageReader_scaledSize, "QImageReader_scaledSize")
	symaddr(&addr_QImageReader_setQuality, "QImageReader_setQuality")
	symaddr(&addr_QImageReader_quality, "QImageReader_quality")
	symaddr(&addr_QImageReader_setScaledClipRect, "QImageReader_setScaledClipRect")
	symaddr(&addr_QImageReader_scaledClipRect, "QImageReader_scaledClipRect")
	symaddr(&addr_QImageReader_setBackgroundColor, "QImageReader_setBackgroundColor")
	symaddr(&addr_QImageReader_backgroundColor, "QImageReader_backgroundColor")
	symaddr(&addr_QImageReader_supportsAnimation, "QImageReader_supportsAnimation")
	symaddr(&addr_QImageReader_transformation, "QImageReader_transformation")
	symaddr(&addr_QImageReader_setAutoTransform, "QImageReader_setAutoTransform")
	symaddr(&addr_QImageReader_autoTransform, "QImageReader_autoTransform")
	symaddr(&addr_QImageReader_subType, "QImageReader_subType")
	symaddr(&addr_QImageReader_supportedSubTypes, "QImageReader_supportedSubTypes")
	symaddr(&addr_QImageReader_canRead, "QImageReader_canRead")
	symaddr(&addr_QImageReader_read, "QImageReader_read")
	symaddr(&addr_QImageReader_read2, "QImageReader_read2")
	symaddr(&addr_QImageReader_jumpToNextImage, "QImageReader_jumpToNextImage")
	symaddr(&addr_QImageReader_jumpToImage, "QImageReader_jumpToImage")
	symaddr(&addr_QImageReader_loopCount, "QImageReader_loopCount")
	symaddr(&addr_QImageReader_imageCount, "QImageReader_imageCount")
	symaddr(&addr_QImageReader_nextImageDelay, "QImageReader_nextImageDelay")
	symaddr(&addr_QImageReader_currentImageNumber, "QImageReader_currentImageNumber")
	symaddr(&addr_QImageReader_currentImageRect, "QImageReader_currentImageRect")
	symaddr(&addr_QImageReader_error, "QImageReader_error")
	symaddr(&addr_QImageReader_errorString, "QImageReader_errorString")
	symaddr(&addr_QImageReader_supportsOption, "QImageReader_supportsOption")
	symaddr(&addr_QImageReader_imageFormat2, "QImageReader_imageFormat2")
	symaddr(&addr_QImageReader_imageFormat3, "QImageReader_imageFormat3")
	symaddr(&addr_QImageReader_supportedImageFormats, "QImageReader_supportedImageFormats")
	symaddr(&addr_QImageReader_supportedMimeTypes, "QImageReader_supportedMimeTypes")
	symaddr(&addr_QImageReader_imageFormatsForMimeType, "QImageReader_imageFormatsForMimeType")
	symaddr(&addr_QImageReader_allocationLimit, "QImageReader_allocationLimit")
	symaddr(&addr_QImageReader_setAllocationLimit, "QImageReader_setAllocationLimit")
	symaddr(&addr_QImageReader_tr2, "QImageReader_tr2")
	symaddr(&addr_QImageReader_tr3, "QImageReader_tr3")
	reg(&pg_QImageReader_delete, "QImageReader_delete")
}

type QImageReader struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QImageReader) UnsafeGoID() int64     { return this.vid }
func (this *QImageReader) UnsafeSetGoID(id int64) { this.vid = id }

func bindQImageReader(h uintptr) *QImageReader {
	if h == 0 {
		return nil
	}
	return &QImageReader{h: h}
}

func newownedQImageReader(h uintptr) *QImageReader {
	o := bindQImageReader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QImageReader).Delete)
	}
	return o
}

func (this *QImageReader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QImageReader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQImageReader(h unsafe.Pointer) *QImageReader { return bindQImageReader(uintptr(h)) }
func UnsafeNewQImageReaderOwned(h unsafe.Pointer) *QImageReader { return newownedQImageReader(uintptr(h)) }

func (this *QImageReader) UnsafeRelease() { this.h = 0 }

func NewQImageReader() *QImageReader {
	__r, _, _ := purego.SyscallN(addr_QImageReader_new)
	__o := newownedQImageReader(__r)
	return __o
}

func NewQImageReader2(device *QIODevice) *QImageReader {
	__r, _, _ := purego.SyscallN(addr_QImageReader_new2, uintptr(device.UnsafePointer()))
	__o := newownedQImageReader(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQImageReader3(fileName string) *QImageReader {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImageReader_new3, fileName_p, uintptr(len(fileName)))
	__o := newownedQImageReader(__r)
	runtime.KeepAlive(fileName)
	return __o
}

func NewQImageReader4(device *QIODevice, format *QByteArray) *QImageReader {
	__r, _, _ := purego.SyscallN(addr_QImageReader_new4, uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	__o := newownedQImageReader(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	return __o
}

func NewQImageReader5(fileName string, format *QByteArray) *QImageReader {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImageReader_new5, fileName_p, uintptr(len(fileName)), uintptr(format.UnsafePointer()))
	__o := newownedQImageReader(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format)
	return __o
}

func QImageReader_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QImageReader) SetFormat(format *QByteArray) {
	purego.SyscallN(addr_QImageReader_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QImageReader) Format() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageReader_format, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QImageReader) SetAutoDetectImageFormat(enabled bool) {
	purego.SyscallN(addr_QImageReader_setAutoDetectImageFormat, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QImageReader) AutoDetectImageFormat() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_autoDetectImageFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) SetDecideFormatFromContent(ignored bool) {
	purego.SyscallN(addr_QImageReader_setDecideFormatFromContent, this.h, uintptr(b2i(ignored)))
	runtime.KeepAlive(this)
}

func (this *QImageReader) DecideFormatFromContent() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_decideFormatFromContent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QImageReader_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QImageReader) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QImageReader_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QImageReader) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QImageReader_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QImageReader) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QImageReader) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QImageReader_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QImageReader) ImageFormat() QImage__Format {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_imageFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImage__Format(__r)
}

func (this *QImageReader) TextKeys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_textKeys, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QImageReader) Text(key string) string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_text, this.h, key_p, uintptr(len(key)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return pgStr(__r, __n)
}

func (this *QImageReader) SetClipRect(rect *QRect) {
	purego.SyscallN(addr_QImageReader_setClipRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QImageReader) ClipRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImageReader_clipRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QImageReader) SetScaledSize(size *QSize) {
	purego.SyscallN(addr_QImageReader_setScaledSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QImageReader) ScaledSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QImageReader_scaledSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QImageReader) SetQuality(quality int) {
	purego.SyscallN(addr_QImageReader_setQuality, this.h, uintptr(quality))
	runtime.KeepAlive(this)
}

func (this *QImageReader) Quality() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_quality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageReader) SetScaledClipRect(rect *QRect) {
	purego.SyscallN(addr_QImageReader_setScaledClipRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QImageReader) ScaledClipRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImageReader_scaledClipRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QImageReader) SetBackgroundColor(color *QColor) {
	purego.SyscallN(addr_QImageReader_setBackgroundColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QImageReader) BackgroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QImageReader_backgroundColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QImageReader) SupportsAnimation() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_supportsAnimation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) Transformation() QImageIOHandler__Transformation {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_transformation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImageIOHandler__Transformation(__r)
}

func (this *QImageReader) SetAutoTransform(enabled bool) {
	purego.SyscallN(addr_QImageReader_setAutoTransform, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QImageReader) AutoTransform() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_autoTransform, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) SubType() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageReader_subType, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QImageReader) SupportedSubTypes() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_supportedSubTypes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QImageReader) CanRead() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_canRead, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) Read() *QImage {
	__r, _, _ := purego.SyscallN(addr_QImageReader_read, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QImageReader) ReadWithImage(image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_read2, this.h, uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageReader) JumpToNextImage() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_jumpToNextImage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) JumpToImage(imageNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_jumpToImage, this.h, uintptr(imageNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageReader) LoopCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_loopCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageReader) ImageCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_imageCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageReader) NextImageDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_nextImageDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageReader) CurrentImageNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_currentImageNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageReader) CurrentImageRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImageReader_currentImageRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QImageReader) Error() QImageReader__ImageReaderError {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImageReader__ImageReaderError(__r)
}

func (this *QImageReader) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QImageReader) SupportsOption(option QImageIOHandler__ImageOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_supportsOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QImageReader_ImageFormat(fileName string) *QByteArray {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImageReader_imageFormat2, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(fileName)
	return newownedQByteArray(__r)
}

func QImageReader_ImageFormatWithDevice(device *QIODevice) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageReader_imageFormat3, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(device)
	return newownedQByteArray(__r)
}

func QImageReader_SupportedImageFormats() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_supportedImageFormats, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageReader_SupportedMimeTypes() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_supportedMimeTypes, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageReader_ImageFormatsForMimeType(mimeType *QByteArray) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_imageFormatsForMimeType, uintptr(mimeType.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(mimeType)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageReader_AllocationLimit() int {
	__r0, _, _ := purego.SyscallN(addr_QImageReader_allocationLimit)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QImageReader_SetAllocationLimit(mbLimit int) {
	purego.SyscallN(addr_QImageReader_setAllocationLimit, uintptr(mbLimit))
}

func QImageReader_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QImageReader_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageReader_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QImageReader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QImageReader_delete(this.h)
	}
	this.UnsafeRelease()
}
