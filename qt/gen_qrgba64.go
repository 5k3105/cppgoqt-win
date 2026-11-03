package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRgba64_new uintptr
	addr_QRgba64_new2 uintptr
	addr_QRgba64_fromRgba64 uintptr
	addr_QRgba64_fromRgba642 uintptr
	addr_QRgba64_fromRgba uintptr
	addr_QRgba64_fromArgb32 uintptr
	addr_QRgba64_isOpaque uintptr
	addr_QRgba64_isTransparent uintptr
	addr_QRgba64_red uintptr
	addr_QRgba64_green uintptr
	addr_QRgba64_blue uintptr
	addr_QRgba64_alpha uintptr
	addr_QRgba64_setRed uintptr
	addr_QRgba64_setGreen uintptr
	addr_QRgba64_setBlue uintptr
	addr_QRgba64_setAlpha uintptr
	addr_QRgba64_red8 uintptr
	addr_QRgba64_green8 uintptr
	addr_QRgba64_blue8 uintptr
	addr_QRgba64_alpha8 uintptr
	addr_QRgba64_toArgb32 uintptr
	addr_QRgba64_toRgb16 uintptr
	addr_QRgba64_premultiplied uintptr
	addr_QRgba64_unpremultiplied uintptr
	addr_QRgba64_ToQuint64 uintptr
	addr_QRgba64_OperatorAssign uintptr
	pg_QRgba64_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRgba64_new, "QRgba64_new")
	symaddr(&addr_QRgba64_new2, "QRgba64_new2")
	symaddr(&addr_QRgba64_fromRgba64, "QRgba64_fromRgba64")
	symaddr(&addr_QRgba64_fromRgba642, "QRgba64_fromRgba642")
	symaddr(&addr_QRgba64_fromRgba, "QRgba64_fromRgba")
	symaddr(&addr_QRgba64_fromArgb32, "QRgba64_fromArgb32")
	symaddr(&addr_QRgba64_isOpaque, "QRgba64_isOpaque")
	symaddr(&addr_QRgba64_isTransparent, "QRgba64_isTransparent")
	symaddr(&addr_QRgba64_red, "QRgba64_red")
	symaddr(&addr_QRgba64_green, "QRgba64_green")
	symaddr(&addr_QRgba64_blue, "QRgba64_blue")
	symaddr(&addr_QRgba64_alpha, "QRgba64_alpha")
	symaddr(&addr_QRgba64_setRed, "QRgba64_setRed")
	symaddr(&addr_QRgba64_setGreen, "QRgba64_setGreen")
	symaddr(&addr_QRgba64_setBlue, "QRgba64_setBlue")
	symaddr(&addr_QRgba64_setAlpha, "QRgba64_setAlpha")
	symaddr(&addr_QRgba64_red8, "QRgba64_red8")
	symaddr(&addr_QRgba64_green8, "QRgba64_green8")
	symaddr(&addr_QRgba64_blue8, "QRgba64_blue8")
	symaddr(&addr_QRgba64_alpha8, "QRgba64_alpha8")
	symaddr(&addr_QRgba64_toArgb32, "QRgba64_toArgb32")
	symaddr(&addr_QRgba64_toRgb16, "QRgba64_toRgb16")
	symaddr(&addr_QRgba64_premultiplied, "QRgba64_premultiplied")
	symaddr(&addr_QRgba64_unpremultiplied, "QRgba64_unpremultiplied")
	symaddr(&addr_QRgba64_ToQuint64, "QRgba64_ToQuint64")
	symaddr(&addr_QRgba64_OperatorAssign, "QRgba64_OperatorAssign")
	reg(&pg_QRgba64_delete, "QRgba64_delete")
}

type QRgba64 struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRgba64) UnsafeGoID() int64     { return this.vid }
func (this *QRgba64) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRgba64(h uintptr) *QRgba64 {
	if h == 0 {
		return nil
	}
	return &QRgba64{h: h}
}

func newownedQRgba64(h uintptr) *QRgba64 {
	o := bindQRgba64(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRgba64).Delete)
	}
	return o
}

func (this *QRgba64) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRgba64) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRgba64(h unsafe.Pointer) *QRgba64 { return bindQRgba64(uintptr(h)) }
func UnsafeNewQRgba64Owned(h unsafe.Pointer) *QRgba64 { return newownedQRgba64(uintptr(h)) }

func (this *QRgba64) UnsafeRelease() { this.h = 0 }

func NewQRgba64() *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_new)
	__o := newownedQRgba64(__r)
	return __o
}

func NewQRgba642(arg1 *QRgba64) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQRgba64(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QRgba64_FromRgba64(c uint64) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_fromRgba64, uintptr(c))
	return newownedQRgba64(__r)
}

func QRgba64_FromRgba642(red uint16, green uint16, blue uint16, alpha uint16) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_fromRgba642, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha))
	return newownedQRgba64(__r)
}

func QRgba64_FromRgba(red byte, green byte, blue byte, alpha byte) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_fromRgba, uintptr(red), uintptr(green), uintptr(blue), uintptr(alpha))
	return newownedQRgba64(__r)
}

func QRgba64_FromArgb32(rgb uint) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_fromArgb32, uintptr(rgb))
	return newownedQRgba64(__r)
}

func (this *QRgba64) IsOpaque() bool {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_isOpaque, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRgba64) IsTransparent() bool {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_isTransparent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRgba64) Red() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_red, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QRgba64) Green() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_green, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QRgba64) Blue() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_blue, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QRgba64) Alpha() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_alpha, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QRgba64) SetRed(_red uint16) {
	purego.SyscallN(addr_QRgba64_setRed, this.h, uintptr(_red))
	runtime.KeepAlive(this)
}

func (this *QRgba64) SetGreen(_green uint16) {
	purego.SyscallN(addr_QRgba64_setGreen, this.h, uintptr(_green))
	runtime.KeepAlive(this)
}

func (this *QRgba64) SetBlue(_blue uint16) {
	purego.SyscallN(addr_QRgba64_setBlue, this.h, uintptr(_blue))
	runtime.KeepAlive(this)
}

func (this *QRgba64) SetAlpha(_alpha uint16) {
	purego.SyscallN(addr_QRgba64_setAlpha, this.h, uintptr(_alpha))
	runtime.KeepAlive(this)
}

func (this *QRgba64) Red8() byte {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_red8, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QRgba64) Green8() byte {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_green8, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QRgba64) Blue8() byte {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_blue8, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QRgba64) Alpha8() byte {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_alpha8, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QRgba64) ToArgb32() uint {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_toArgb32, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QRgba64) ToRgb16() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_toRgb16, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QRgba64) Premultiplied() *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_premultiplied, this.h)
	runtime.KeepAlive(this)
	return newownedQRgba64(__r)
}

func (this *QRgba64) Unpremultiplied() *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_unpremultiplied, this.h)
	runtime.KeepAlive(this)
	return newownedQRgba64(__r)
}

func (this *QRgba64) ToQuint64() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRgba64_ToQuint64, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRgba64) OperatorAssign(_rgba uint64) *QRgba64 {
	__r, _, _ := purego.SyscallN(addr_QRgba64_OperatorAssign, this.h, uintptr(_rgba))
	runtime.KeepAlive(this)
	return bindQRgba64(__r)
}

func (this *QRgba64) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRgba64_delete(this.h)
	}
	this.UnsafeRelease()
}
