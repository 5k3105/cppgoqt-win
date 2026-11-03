package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPixelFormat_new uintptr
	addr_QPixelFormat_new2 uintptr
	addr_QPixelFormat_new3 uintptr
	addr_QPixelFormat_new4 uintptr
	addr_QPixelFormat_new5 uintptr
	addr_QPixelFormat_colorModel uintptr
	addr_QPixelFormat_channelCount uintptr
	addr_QPixelFormat_redSize uintptr
	addr_QPixelFormat_greenSize uintptr
	addr_QPixelFormat_blueSize uintptr
	addr_QPixelFormat_cyanSize uintptr
	addr_QPixelFormat_magentaSize uintptr
	addr_QPixelFormat_yellowSize uintptr
	addr_QPixelFormat_blackSize uintptr
	addr_QPixelFormat_hueSize uintptr
	addr_QPixelFormat_saturationSize uintptr
	addr_QPixelFormat_lightnessSize uintptr
	addr_QPixelFormat_brightnessSize uintptr
	addr_QPixelFormat_alphaSize uintptr
	addr_QPixelFormat_bitsPerPixel uintptr
	addr_QPixelFormat_alphaUsage uintptr
	addr_QPixelFormat_alphaPosition uintptr
	addr_QPixelFormat_premultiplied uintptr
	addr_QPixelFormat_typeInterpretation uintptr
	addr_QPixelFormat_byteOrder uintptr
	addr_QPixelFormat_yuvLayout uintptr
	addr_QPixelFormat_subEnum uintptr
	addr_QPixelFormat_OperatorEqual uintptr
	addr_QPixelFormat_OperatorNotEqual uintptr
	pg_QPixelFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPixelFormat_new, "QPixelFormat_new")
	symaddr(&addr_QPixelFormat_new2, "QPixelFormat_new2")
	symaddr(&addr_QPixelFormat_new3, "QPixelFormat_new3")
	symaddr(&addr_QPixelFormat_new4, "QPixelFormat_new4")
	symaddr(&addr_QPixelFormat_new5, "QPixelFormat_new5")
	symaddr(&addr_QPixelFormat_colorModel, "QPixelFormat_colorModel")
	symaddr(&addr_QPixelFormat_channelCount, "QPixelFormat_channelCount")
	symaddr(&addr_QPixelFormat_redSize, "QPixelFormat_redSize")
	symaddr(&addr_QPixelFormat_greenSize, "QPixelFormat_greenSize")
	symaddr(&addr_QPixelFormat_blueSize, "QPixelFormat_blueSize")
	symaddr(&addr_QPixelFormat_cyanSize, "QPixelFormat_cyanSize")
	symaddr(&addr_QPixelFormat_magentaSize, "QPixelFormat_magentaSize")
	symaddr(&addr_QPixelFormat_yellowSize, "QPixelFormat_yellowSize")
	symaddr(&addr_QPixelFormat_blackSize, "QPixelFormat_blackSize")
	symaddr(&addr_QPixelFormat_hueSize, "QPixelFormat_hueSize")
	symaddr(&addr_QPixelFormat_saturationSize, "QPixelFormat_saturationSize")
	symaddr(&addr_QPixelFormat_lightnessSize, "QPixelFormat_lightnessSize")
	symaddr(&addr_QPixelFormat_brightnessSize, "QPixelFormat_brightnessSize")
	symaddr(&addr_QPixelFormat_alphaSize, "QPixelFormat_alphaSize")
	symaddr(&addr_QPixelFormat_bitsPerPixel, "QPixelFormat_bitsPerPixel")
	symaddr(&addr_QPixelFormat_alphaUsage, "QPixelFormat_alphaUsage")
	symaddr(&addr_QPixelFormat_alphaPosition, "QPixelFormat_alphaPosition")
	symaddr(&addr_QPixelFormat_premultiplied, "QPixelFormat_premultiplied")
	symaddr(&addr_QPixelFormat_typeInterpretation, "QPixelFormat_typeInterpretation")
	symaddr(&addr_QPixelFormat_byteOrder, "QPixelFormat_byteOrder")
	symaddr(&addr_QPixelFormat_yuvLayout, "QPixelFormat_yuvLayout")
	symaddr(&addr_QPixelFormat_subEnum, "QPixelFormat_subEnum")
	symaddr(&addr_QPixelFormat_OperatorEqual, "QPixelFormat_OperatorEqual")
	symaddr(&addr_QPixelFormat_OperatorNotEqual, "QPixelFormat_OperatorNotEqual")
	reg(&pg_QPixelFormat_delete, "QPixelFormat_delete")
}

type QPixelFormat struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPixelFormat) UnsafeGoID() int64     { return this.vid }
func (this *QPixelFormat) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPixelFormat(h uintptr) *QPixelFormat {
	if h == 0 {
		return nil
	}
	return &QPixelFormat{h: h}
}

func newownedQPixelFormat(h uintptr) *QPixelFormat {
	o := bindQPixelFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPixelFormat).Delete)
	}
	return o
}

func (this *QPixelFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPixelFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPixelFormat(h unsafe.Pointer) *QPixelFormat { return bindQPixelFormat(uintptr(h)) }
func UnsafeNewQPixelFormatOwned(h unsafe.Pointer) *QPixelFormat { return newownedQPixelFormat(uintptr(h)) }

func (this *QPixelFormat) UnsafeRelease() { this.h = 0 }

func NewQPixelFormat() *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QPixelFormat_new)
	__o := newownedQPixelFormat(__r)
	return __o
}

func NewQPixelFormat2(colorModel QPixelFormat__ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage QPixelFormat__AlphaUsage, alphaPosition QPixelFormat__AlphaPosition, premultiplied QPixelFormat__AlphaPremultiplied, typeInterpretation QPixelFormat__TypeInterpretation) *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QPixelFormat_new2, uintptr(colorModel), uintptr(firstSize), uintptr(secondSize), uintptr(thirdSize), uintptr(fourthSize), uintptr(fifthSize), uintptr(alphaSize), uintptr(alphaUsage), uintptr(alphaPosition), uintptr(premultiplied), uintptr(typeInterpretation))
	__o := newownedQPixelFormat(__r)
	return __o
}

func NewQPixelFormat3(arg1 *QPixelFormat) *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QPixelFormat_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQPixelFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQPixelFormat4(colorModel QPixelFormat__ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage QPixelFormat__AlphaUsage, alphaPosition QPixelFormat__AlphaPosition, premultiplied QPixelFormat__AlphaPremultiplied, typeInterpretation QPixelFormat__TypeInterpretation, byteOrder QPixelFormat__ByteOrder) *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QPixelFormat_new4, uintptr(colorModel), uintptr(firstSize), uintptr(secondSize), uintptr(thirdSize), uintptr(fourthSize), uintptr(fifthSize), uintptr(alphaSize), uintptr(alphaUsage), uintptr(alphaPosition), uintptr(premultiplied), uintptr(typeInterpretation), uintptr(byteOrder))
	__o := newownedQPixelFormat(__r)
	return __o
}

func NewQPixelFormat5(colorModel QPixelFormat__ColorModel, firstSize byte, secondSize byte, thirdSize byte, fourthSize byte, fifthSize byte, alphaSize byte, alphaUsage QPixelFormat__AlphaUsage, alphaPosition QPixelFormat__AlphaPosition, premultiplied QPixelFormat__AlphaPremultiplied, typeInterpretation QPixelFormat__TypeInterpretation, byteOrder QPixelFormat__ByteOrder, subEnum byte) *QPixelFormat {
	__r, _, _ := purego.SyscallN(addr_QPixelFormat_new5, uintptr(colorModel), uintptr(firstSize), uintptr(secondSize), uintptr(thirdSize), uintptr(fourthSize), uintptr(fifthSize), uintptr(alphaSize), uintptr(alphaUsage), uintptr(alphaPosition), uintptr(premultiplied), uintptr(typeInterpretation), uintptr(byteOrder), uintptr(subEnum))
	__o := newownedQPixelFormat(__r)
	return __o
}

func (this *QPixelFormat) ColorModel() QPixelFormat__ColorModel {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_colorModel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__ColorModel(__r)
}

func (this *QPixelFormat) ChannelCount() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_channelCount, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) RedSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_redSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) GreenSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_greenSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) BlueSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_blueSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) CyanSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_cyanSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) MagentaSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_magentaSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) YellowSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_yellowSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) BlackSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_blackSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) HueSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_hueSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) SaturationSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_saturationSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) LightnessSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_lightnessSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) BrightnessSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_brightnessSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) AlphaSize() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_alphaSize, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) BitsPerPixel() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_bitsPerPixel, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) AlphaUsage() QPixelFormat__AlphaUsage {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_alphaUsage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__AlphaUsage(__r)
}

func (this *QPixelFormat) AlphaPosition() QPixelFormat__AlphaPosition {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_alphaPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__AlphaPosition(__r)
}

func (this *QPixelFormat) Premultiplied() QPixelFormat__AlphaPremultiplied {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_premultiplied, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__AlphaPremultiplied(__r)
}

func (this *QPixelFormat) TypeInterpretation() QPixelFormat__TypeInterpretation {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_typeInterpretation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__TypeInterpretation(__r)
}

func (this *QPixelFormat) ByteOrder() QPixelFormat__ByteOrder {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_byteOrder, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__ByteOrder(__r)
}

func (this *QPixelFormat) YuvLayout() QPixelFormat__YUVLayout {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_yuvLayout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPixelFormat__YUVLayout(__r)
}

func (this *QPixelFormat) SubEnum() byte {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_subEnum, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPixelFormat) OperatorEqual(fmt2 *QPixelFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_OperatorEqual, this.h, uintptr(fmt2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fmt2)
	return __r != 0
}

func (this *QPixelFormat) OperatorNotEqual(fmt2 *QPixelFormat) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixelFormat_OperatorNotEqual, this.h, uintptr(fmt2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fmt2)
	return __r != 0
}

func (this *QPixelFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPixelFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
