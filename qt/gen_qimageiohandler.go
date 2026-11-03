package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QImageIOHandler_setDevice uintptr
	addr_QImageIOHandler_device uintptr
	addr_QImageIOHandler_setFormat uintptr
	addr_QImageIOHandler_setFormat2 uintptr
	addr_QImageIOHandler_format uintptr
	addr_QImageIOHandler_canRead uintptr
	addr_QImageIOHandler_read uintptr
	addr_QImageIOHandler_write uintptr
	addr_QImageIOHandler_option uintptr
	addr_QImageIOHandler_setOption uintptr
	addr_QImageIOHandler_supportsOption uintptr
	addr_QImageIOHandler_jumpToNextImage uintptr
	addr_QImageIOHandler_jumpToImage uintptr
	addr_QImageIOHandler_loopCount uintptr
	addr_QImageIOHandler_imageCount uintptr
	addr_QImageIOHandler_nextImageDelay uintptr
	addr_QImageIOHandler_currentImageNumber uintptr
	addr_QImageIOHandler_currentImageRect uintptr
	addr_QImageIOHandler_allocateImage uintptr
	addr_QImageIOHandler_write_default uintptr
	addr_QImageIOHandler_option_default uintptr
	addr_QImageIOHandler_setOption_default uintptr
	addr_QImageIOHandler_supportsOption_default uintptr
	addr_QImageIOHandler_jumpToNextImage_default uintptr
	addr_QImageIOHandler_jumpToImage_default uintptr
	addr_QImageIOHandler_loopCount_default uintptr
	addr_QImageIOHandler_imageCount_default uintptr
	addr_QImageIOHandler_nextImageDelay_default uintptr
	addr_QImageIOHandler_currentImageNumber_default uintptr
	addr_QImageIOHandler_currentImageRect_default uintptr
	pg_QImageIOHandler__set_override func(uintptr, int32)
	pg_QImageIOHandler__Go_new func(int64) uintptr
	pg_QImageIOHandler_delete func(uintptr)
)

func init() {
	symaddr(&addr_QImageIOHandler_setDevice, "QImageIOHandler_setDevice")
	symaddr(&addr_QImageIOHandler_device, "QImageIOHandler_device")
	symaddr(&addr_QImageIOHandler_setFormat, "QImageIOHandler_setFormat")
	symaddr(&addr_QImageIOHandler_setFormat2, "QImageIOHandler_setFormat2")
	symaddr(&addr_QImageIOHandler_format, "QImageIOHandler_format")
	symaddr(&addr_QImageIOHandler_canRead, "QImageIOHandler_canRead")
	symaddr(&addr_QImageIOHandler_read, "QImageIOHandler_read")
	symaddr(&addr_QImageIOHandler_write, "QImageIOHandler_write")
	symaddr(&addr_QImageIOHandler_option, "QImageIOHandler_option")
	symaddr(&addr_QImageIOHandler_setOption, "QImageIOHandler_setOption")
	symaddr(&addr_QImageIOHandler_supportsOption, "QImageIOHandler_supportsOption")
	symaddr(&addr_QImageIOHandler_jumpToNextImage, "QImageIOHandler_jumpToNextImage")
	symaddr(&addr_QImageIOHandler_jumpToImage, "QImageIOHandler_jumpToImage")
	symaddr(&addr_QImageIOHandler_loopCount, "QImageIOHandler_loopCount")
	symaddr(&addr_QImageIOHandler_imageCount, "QImageIOHandler_imageCount")
	symaddr(&addr_QImageIOHandler_nextImageDelay, "QImageIOHandler_nextImageDelay")
	symaddr(&addr_QImageIOHandler_currentImageNumber, "QImageIOHandler_currentImageNumber")
	symaddr(&addr_QImageIOHandler_currentImageRect, "QImageIOHandler_currentImageRect")
	symaddr(&addr_QImageIOHandler_allocateImage, "QImageIOHandler_allocateImage")
	symaddr(&addr_QImageIOHandler_write_default, "QImageIOHandler_write_default")
	symaddr(&addr_QImageIOHandler_option_default, "QImageIOHandler_option_default")
	symaddr(&addr_QImageIOHandler_setOption_default, "QImageIOHandler_setOption_default")
	symaddr(&addr_QImageIOHandler_supportsOption_default, "QImageIOHandler_supportsOption_default")
	symaddr(&addr_QImageIOHandler_jumpToNextImage_default, "QImageIOHandler_jumpToNextImage_default")
	symaddr(&addr_QImageIOHandler_jumpToImage_default, "QImageIOHandler_jumpToImage_default")
	symaddr(&addr_QImageIOHandler_loopCount_default, "QImageIOHandler_loopCount_default")
	symaddr(&addr_QImageIOHandler_imageCount_default, "QImageIOHandler_imageCount_default")
	symaddr(&addr_QImageIOHandler_nextImageDelay_default, "QImageIOHandler_nextImageDelay_default")
	symaddr(&addr_QImageIOHandler_currentImageNumber_default, "QImageIOHandler_currentImageNumber_default")
	symaddr(&addr_QImageIOHandler_currentImageRect_default, "QImageIOHandler_currentImageRect_default")
	reg(&pg_QImageIOHandler__set_override, "QImageIOHandler__set_override")
	reg(&pg_QImageIOHandler__Go_new, "QImageIOHandler__Go_new")
	reg(&pg_QImageIOHandler_delete, "QImageIOHandler_delete")
}

type QImageIOHandler struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QImageIOHandler) UnsafeGoID() int64     { return this.vid }
func (this *QImageIOHandler) UnsafeSetGoID(id int64) { this.vid = id }

func bindQImageIOHandler(h uintptr) *QImageIOHandler {
	if h == 0 {
		return nil
	}
	return &QImageIOHandler{h: h}
}

func newownedQImageIOHandler(h uintptr) *QImageIOHandler {
	o := bindQImageIOHandler(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QImageIOHandler).Delete)
	}
	return o
}

func (this *QImageIOHandler) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QImageIOHandler) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQImageIOHandler(h unsafe.Pointer) *QImageIOHandler { return bindQImageIOHandler(uintptr(h)) }
func UnsafeNewQImageIOHandlerOwned(h unsafe.Pointer) *QImageIOHandler { return newownedQImageIOHandler(uintptr(h)) }

func (this *QImageIOHandler) UnsafeRelease() { this.h = 0 }

func (this *QImageIOHandler) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QImageIOHandler_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QImageIOHandler) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QImageIOHandler) SetFormat(format *QByteArray) {
	purego.SyscallN(addr_QImageIOHandler_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QImageIOHandler) SetFormatWithFormat(format *QByteArray) {
	purego.SyscallN(addr_QImageIOHandler_setFormat2, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QImageIOHandler) Format() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_format, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QImageIOHandler) CanRead() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_canRead, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) Read(image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_read, this.h, uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageIOHandler) Write(image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_write, this.h, uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageIOHandler) Option(option QImageIOHandler__ImageOption) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_option, this.h, uintptr(option))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QImageIOHandler) SetOption(option QImageIOHandler__ImageOption, value *QVariant) {
	purego.SyscallN(addr_QImageIOHandler_setOption, this.h, uintptr(option), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QImageIOHandler) SupportsOption(option QImageIOHandler__ImageOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_supportsOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) JumpToNextImage() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_jumpToNextImage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) JumpToImage(imageNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_jumpToImage, this.h, uintptr(imageNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) LoopCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_loopCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) ImageCount() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_imageCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) NextImageDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_nextImageDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) CurrentImageNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_currentImageNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) CurrentImageRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_currentImageRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func QImageIOHandler_AllocateImage(size *QSize, format QImage__Format, image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_allocateImage, uintptr(size.UnsafePointer()), uintptr(format), uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(size)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageIOHandler) WriteDefault(image *QImage) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_write_default, this.h, uintptr(image.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	return __r != 0
}

func (this *QImageIOHandler) OptionDefault(option QImageIOHandler__ImageOption) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_option_default, this.h, uintptr(option))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QImageIOHandler) SetOptionDefault(option QImageIOHandler__ImageOption, value *QVariant) {
	purego.SyscallN(addr_QImageIOHandler_setOption_default, this.h, uintptr(option), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QImageIOHandler) SupportsOptionDefault(option QImageIOHandler__ImageOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_supportsOption_default, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) JumpToNextImageDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_jumpToNextImage_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) JumpToImageDefault(imageNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_jumpToImage_default, this.h, uintptr(imageNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QImageIOHandler) LoopCountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_loopCount_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) ImageCountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_imageCount_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) NextImageDelayDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_nextImageDelay_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) CurrentImageNumberDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QImageIOHandler_currentImageNumber_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QImageIOHandler) CurrentImageRectDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QImageIOHandler_currentImageRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func QImageIOHandler_virtIdx(slot int64) int32 {
	switch slot {
	case -5697610016496362555:
		return 0
	case 5541207787354052101:
		return 1
	case -5099743077423337732:
		return 2
	case 270525795752552276:
		return 3
	case 5527147985269015870:
		return 4
	case 89707390584546352:
		return 5
	case 3906762202460275946:
		return 6
	case 732703795301504057:
		return 7
	case 60103087788910490:
		return 8
	case -7262609269034271663:
		return 9
	case -5074503007600271320:
		return 10
	case 2054658236031792398:
		return 11
	case 2007698496588377553:
		return 12
	}
	return -1
}

func QImageIOHandler_enable(h uintptr, idx int32) { pg_QImageIOHandler__set_override(h, idx) }

func NewQImageIOHandler() *QImageIOHandler {
	id := nextVirtID()
	o := newownedQImageIOHandler(pg_QImageIOHandler__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QImageIOHandler_virtIdx, QImageIOHandler_enable)
	return o
}

func (this *QImageIOHandler) OnCanRead(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -5697610016496362555, func() int32 { return b2i(fn()) }) // slot "canRead"
}

func (this *QImageIOHandler) OnRead(fn func(*QImage) bool) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(image uintptr) int32 { return b2i(fn(UnsafeNewQImage(unsafe.Pointer(image)))) }) // slot "read"
}

func (this *QImageIOHandler) OnWrite(fn func(*QImage) bool) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(image uintptr) int32 { return b2i(fn(UnsafeNewQImage(unsafe.Pointer(image)))) }) // slot "write"
}

func (this *QImageIOHandler) OnOption(fn func(QImageIOHandler__ImageOption) *QVariant) {
	registerVirt(this.UnsafeGoID(), 270525795752552276, func(option int32) uintptr { return uintptr(fn(QImageIOHandler__ImageOption(option)).UnsafePointer()) }) // slot "option"
}

func (this *QImageIOHandler) OnSetOption(fn func(QImageIOHandler__ImageOption, *QVariant)) {
	registerVirt(this.UnsafeGoID(), 5527147985269015870, func(option int32, value uintptr) { fn(QImageIOHandler__ImageOption(option), UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "setOption"
}

func (this *QImageIOHandler) OnSupportsOption(fn func(QImageIOHandler__ImageOption) bool) {
	registerVirt(this.UnsafeGoID(), 89707390584546352, func(option int32) int32 { return b2i(fn(QImageIOHandler__ImageOption(option))) }) // slot "supportsOption"
}

func (this *QImageIOHandler) OnJumpToNextImage(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 3906762202460275946, func() int32 { return b2i(fn()) }) // slot "jumpToNextImage"
}

func (this *QImageIOHandler) OnJumpToImage(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 732703795301504057, func(imageNumber int32) int32 { return b2i(fn(int(imageNumber))) }) // slot "jumpToImage"
}

func (this *QImageIOHandler) OnLoopCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), 60103087788910490, func() int32 { return int32(fn()) }) // slot "loopCount"
}

func (this *QImageIOHandler) OnImageCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -7262609269034271663, func() int32 { return int32(fn()) }) // slot "imageCount"
}

func (this *QImageIOHandler) OnNextImageDelay(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5074503007600271320, func() int32 { return int32(fn()) }) // slot "nextImageDelay"
}

func (this *QImageIOHandler) OnCurrentImageNumber(fn func() int) {
	registerVirt(this.UnsafeGoID(), 2054658236031792398, func() int32 { return int32(fn()) }) // slot "currentImageNumber"
}

func (this *QImageIOHandler) OnCurrentImageRect(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 2007698496588377553, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "currentImageRect"
}

func (this *QImageIOHandler) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QImageIOHandler_delete(this.h)
	}
	this.UnsafeRelease()
}
