package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QImageWriter_new uintptr
	addr_QImageWriter_new2 uintptr
	addr_QImageWriter_new3 uintptr
	addr_QImageWriter_new4 uintptr
	addr_QImageWriter_tr uintptr
	addr_QImageWriter_setFormat uintptr
	addr_QImageWriter_format uintptr
	addr_QImageWriter_setDevice uintptr
	addr_QImageWriter_device uintptr
	addr_QImageWriter_setFileName uintptr
	addr_QImageWriter_fileName uintptr
	addr_QImageWriter_setQuality uintptr
	addr_QImageWriter_quality uintptr
	addr_QImageWriter_setCompression uintptr
	addr_QImageWriter_compression uintptr
	addr_QImageWriter_setSubType uintptr
	addr_QImageWriter_subType uintptr
	addr_QImageWriter_supportedSubTypes uintptr
	addr_QImageWriter_setOptimizedWrite uintptr
	addr_QImageWriter_optimizedWrite uintptr
	addr_QImageWriter_setProgressiveScanWrite uintptr
	addr_QImageWriter_progressiveScanWrite uintptr
	addr_QImageWriter_transformation uintptr
	addr_QImageWriter_setTransformation uintptr
	addr_QImageWriter_setText uintptr
	addr_QImageWriter_canWrite uintptr
	addr_QImageWriter_write uintptr
	addr_QImageWriter_error uintptr
	addr_QImageWriter_errorString uintptr
	addr_QImageWriter_supportsOption uintptr
	addr_QImageWriter_supportedImageFormats uintptr
	addr_QImageWriter_supportedMimeTypes uintptr
	addr_QImageWriter_imageFormatsForMimeType uintptr
	addr_QImageWriter_tr2 uintptr
	addr_QImageWriter_tr3 uintptr
	pg_QImageWriter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QImageWriter_new, "QImageWriter_new")
	symaddr(&addr_QImageWriter_new2, "QImageWriter_new2")
	symaddr(&addr_QImageWriter_new3, "QImageWriter_new3")
	symaddr(&addr_QImageWriter_new4, "QImageWriter_new4")
	symaddr(&addr_QImageWriter_tr, "QImageWriter_tr")
	symaddr(&addr_QImageWriter_setFormat, "QImageWriter_setFormat")
	symaddr(&addr_QImageWriter_format, "QImageWriter_format")
	symaddr(&addr_QImageWriter_setDevice, "QImageWriter_setDevice")
	symaddr(&addr_QImageWriter_device, "QImageWriter_device")
	symaddr(&addr_QImageWriter_setFileName, "QImageWriter_setFileName")
	symaddr(&addr_QImageWriter_fileName, "QImageWriter_fileName")
	symaddr(&addr_QImageWriter_setQuality, "QImageWriter_setQuality")
	symaddr(&addr_QImageWriter_quality, "QImageWriter_quality")
	symaddr(&addr_QImageWriter_setCompression, "QImageWriter_setCompression")
	symaddr(&addr_QImageWriter_compression, "QImageWriter_compression")
	symaddr(&addr_QImageWriter_setSubType, "QImageWriter_setSubType")
	symaddr(&addr_QImageWriter_subType, "QImageWriter_subType")
	symaddr(&addr_QImageWriter_supportedSubTypes, "QImageWriter_supportedSubTypes")
	symaddr(&addr_QImageWriter_setOptimizedWrite, "QImageWriter_setOptimizedWrite")
	symaddr(&addr_QImageWriter_optimizedWrite, "QImageWriter_optimizedWrite")
	symaddr(&addr_QImageWriter_setProgressiveScanWrite, "QImageWriter_setProgressiveScanWrite")
	symaddr(&addr_QImageWriter_progressiveScanWrite, "QImageWriter_progressiveScanWrite")
	symaddr(&addr_QImageWriter_transformation, "QImageWriter_transformation")
	symaddr(&addr_QImageWriter_setTransformation, "QImageWriter_setTransformation")
	symaddr(&addr_QImageWriter_setText, "QImageWriter_setText")
	symaddr(&addr_QImageWriter_canWrite, "QImageWriter_canWrite")
	symaddr(&addr_QImageWriter_write, "QImageWriter_write")
	symaddr(&addr_QImageWriter_error, "QImageWriter_error")
	symaddr(&addr_QImageWriter_errorString, "QImageWriter_errorString")
	symaddr(&addr_QImageWriter_supportsOption, "QImageWriter_supportsOption")
	symaddr(&addr_QImageWriter_supportedImageFormats, "QImageWriter_supportedImageFormats")
	symaddr(&addr_QImageWriter_supportedMimeTypes, "QImageWriter_supportedMimeTypes")
	symaddr(&addr_QImageWriter_imageFormatsForMimeType, "QImageWriter_imageFormatsForMimeType")
	symaddr(&addr_QImageWriter_tr2, "QImageWriter_tr2")
	symaddr(&addr_QImageWriter_tr3, "QImageWriter_tr3")
	reg(&pg_QImageWriter_delete, "QImageWriter_delete")
}

type QImageWriter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QImageWriter) UnsafeGoID() int64     { return this.vid }
func (this *QImageWriter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQImageWriter(h uintptr) *QImageWriter {
	if h == 0 {
		return nil
	}
	return &QImageWriter{h: h}
}

func newownedQImageWriter(h uintptr) *QImageWriter {
	o := bindQImageWriter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QImageWriter).Delete)
	}
	return o
}

func (this *QImageWriter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QImageWriter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQImageWriter(h unsafe.Pointer) *QImageWriter { return bindQImageWriter(uintptr(h)) }
func UnsafeNewQImageWriterOwned(h unsafe.Pointer) *QImageWriter { return newownedQImageWriter(uintptr(h)) }

func (this *QImageWriter) UnsafeRelease() { this.h = 0 }

func NewQImageWriter() *QImageWriter {
	__r, _, _ := purego.SyscallN(addr_QImageWriter_new)
	__o := newownedQImageWriter(__r)
	return __o
}

func NewQImageWriter2(device *QIODevice, format *QByteArray) *QImageWriter {
	__r, _, _ := purego.SyscallN(addr_QImageWriter_new2, uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	__o := newownedQImageWriter(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	return __o
}

func NewQImageWriter3(fileName string) *QImageWriter {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImageWriter_new3, fileName_p, uintptr(len(fileName)))
	__o := newownedQImageWriter(__r)
	runtime.KeepAlive(fileName)
	return __o
}

func NewQImageWriter4(fileName string, format *QByteArray) *QImageWriter {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QImageWriter_new4, fileName_p, uintptr(len(fileName)), uintptr(format.UnsafePointer()))
	__o := newownedQImageWriter(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format)
	return __o
}

func QImageWriter_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QImageWriter) SetFormat(format *QByteArray) {
	purego.SyscallN(addr_QImageWriter_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QImageWriter) Format() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageWriter_format, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QImageWriter) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QImageWriter_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QImageWriter) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QImageWriter_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QImageWriter) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QImageWriter_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QImageWriter) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QImageWriter) SetQuality(quality int) {
	purego.SyscallN(addr_QImageWriter_setQuality, this.h, uintptr(quality))
	runtime.KeepAlive(this)
}

func (this *QImageWriter) Quality() int {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_quality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageWriter) SetCompression(compression int) {
	purego.SyscallN(addr_QImageWriter_setCompression, this.h, uintptr(compression))
	runtime.KeepAlive(this)
}

func (this *QImageWriter) Compression() int {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_compression, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageWriter) SetSubType(type_ *QByteArray) {
	purego.SyscallN(addr_QImageWriter_setSubType, this.h, uintptr(type_.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(type_)
}

func (this *QImageWriter) SubType() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageWriter_subType, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QImageWriter) SupportedSubTypes() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_supportedSubTypes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QImageWriter) SetOptimizedWrite(optimize bool) {
	purego.SyscallN(addr_QImageWriter_setOptimizedWrite, this.h, uintptr(b2i(optimize)))
	runtime.KeepAlive(this)
}

func (this *QImageWriter) OptimizedWrite() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_optimizedWrite, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageWriter) SetProgressiveScanWrite(progressive bool) {
	purego.SyscallN(addr_QImageWriter_setProgressiveScanWrite, this.h, uintptr(b2i(progressive)))
	runtime.KeepAlive(this)
}

func (this *QImageWriter) ProgressiveScanWrite() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_progressiveScanWrite, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageWriter) Transformation() QImageIOHandler__Transformation {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_transformation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImageIOHandler__Transformation(__r)
}

func (this *QImageWriter) SetTransformation(orientation QImageIOHandler__Transformation) {
	purego.SyscallN(addr_QImageWriter_setTransformation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QImageWriter) SetText(key string, text string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QImageWriter_setText, this.h, key_p, uintptr(len(key)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(text)
}

func (this *QImageWriter) CanWrite() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_canWrite, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageWriter) Write(image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_write, this.h, uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageWriter) Error() QImageWriter__ImageWriterError {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImageWriter__ImageWriterError(__r)
}

func (this *QImageWriter) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QImageWriter) SupportsOption(option QImageIOHandler__ImageOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageWriter_supportsOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QImageWriter_SupportedImageFormats() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_supportedImageFormats, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageWriter_SupportedMimeTypes() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_supportedMimeTypes, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageWriter_ImageFormatsForMimeType(mimeType *QByteArray) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_imageFormatsForMimeType, uintptr(mimeType.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(mimeType)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QImageWriter_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QImageWriter_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QImageWriter_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QImageWriter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QImageWriter_delete(this.h)
	}
	this.UnsafeRelease()
}
