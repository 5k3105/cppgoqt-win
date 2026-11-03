package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontInfo_new uintptr
	addr_QFontInfo_new2 uintptr
	addr_QFontInfo_OperatorAssign uintptr
	addr_QFontInfo_swap uintptr
	addr_QFontInfo_family uintptr
	addr_QFontInfo_styleName uintptr
	addr_QFontInfo_pixelSize uintptr
	addr_QFontInfo_pointSize uintptr
	pg_QFontInfo_pointSizeF func(uintptr) float64
	addr_QFontInfo_italic uintptr
	addr_QFontInfo_style uintptr
	addr_QFontInfo_weight uintptr
	addr_QFontInfo_bold uintptr
	addr_QFontInfo_underline uintptr
	addr_QFontInfo_overline uintptr
	addr_QFontInfo_strikeOut uintptr
	addr_QFontInfo_fixedPitch uintptr
	addr_QFontInfo_styleHint uintptr
	addr_QFontInfo_variableAxes uintptr
	addr_QFontInfo_legacyWeight uintptr
	addr_QFontInfo_exactMatch uintptr
	pg_QFontInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontInfo_new, "QFontInfo_new")
	symaddr(&addr_QFontInfo_new2, "QFontInfo_new2")
	symaddr(&addr_QFontInfo_OperatorAssign, "QFontInfo_OperatorAssign")
	symaddr(&addr_QFontInfo_swap, "QFontInfo_swap")
	symaddr(&addr_QFontInfo_family, "QFontInfo_family")
	symaddr(&addr_QFontInfo_styleName, "QFontInfo_styleName")
	symaddr(&addr_QFontInfo_pixelSize, "QFontInfo_pixelSize")
	symaddr(&addr_QFontInfo_pointSize, "QFontInfo_pointSize")
	reg(&pg_QFontInfo_pointSizeF, "QFontInfo_pointSizeF")
	symaddr(&addr_QFontInfo_italic, "QFontInfo_italic")
	symaddr(&addr_QFontInfo_style, "QFontInfo_style")
	symaddr(&addr_QFontInfo_weight, "QFontInfo_weight")
	symaddr(&addr_QFontInfo_bold, "QFontInfo_bold")
	symaddr(&addr_QFontInfo_underline, "QFontInfo_underline")
	symaddr(&addr_QFontInfo_overline, "QFontInfo_overline")
	symaddr(&addr_QFontInfo_strikeOut, "QFontInfo_strikeOut")
	symaddr(&addr_QFontInfo_fixedPitch, "QFontInfo_fixedPitch")
	symaddr(&addr_QFontInfo_styleHint, "QFontInfo_styleHint")
	symaddr(&addr_QFontInfo_variableAxes, "QFontInfo_variableAxes")
	symaddr(&addr_QFontInfo_legacyWeight, "QFontInfo_legacyWeight")
	symaddr(&addr_QFontInfo_exactMatch, "QFontInfo_exactMatch")
	reg(&pg_QFontInfo_delete, "QFontInfo_delete")
}

type QFontInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFontInfo) UnsafeGoID() int64     { return this.vid }
func (this *QFontInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFontInfo(h uintptr) *QFontInfo {
	if h == 0 {
		return nil
	}
	return &QFontInfo{h: h}
}

func newownedQFontInfo(h uintptr) *QFontInfo {
	o := bindQFontInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFontInfo).Delete)
	}
	return o
}

func (this *QFontInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontInfo(h unsafe.Pointer) *QFontInfo { return bindQFontInfo(uintptr(h)) }
func UnsafeNewQFontInfoOwned(h unsafe.Pointer) *QFontInfo { return newownedQFontInfo(uintptr(h)) }

func (this *QFontInfo) UnsafeRelease() { this.h = 0 }

func NewQFontInfo(arg1 *QFont) *QFontInfo {
	__r, _, _ := purego.SyscallN(addr_QFontInfo_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontInfo(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQFontInfo2(arg1 *QFontInfo) *QFontInfo {
	__r, _, _ := purego.SyscallN(addr_QFontInfo_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQFontInfo(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFontInfo) OperatorAssign(arg1 *QFontInfo) *QFontInfo {
	__r, _, _ := purego.SyscallN(addr_QFontInfo_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFontInfo(__r)
}

func (this *QFontInfo) Swap(other *QFontInfo) {
	purego.SyscallN(addr_QFontInfo_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFontInfo) Family() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontInfo_family, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontInfo) StyleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontInfo_styleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontInfo) PixelSize() int {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_pixelSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontInfo) PointSize() int {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_pointSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontInfo) PointSizeF() float64 {
	__r := pg_QFontInfo_pointSizeF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFontInfo) Italic() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_italic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) Style() QFont__Style {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__Style(__r)
}

func (this *QFontInfo) Weight() int {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_weight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontInfo) Bold() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_bold, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) Underline() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_underline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) Overline() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_overline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) StrikeOut() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_strikeOut, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) FixedPitch() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_fixedPitch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) StyleHint() QFont__StyleHint {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_styleHint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFont__StyleHint(__r)
}

func (this *QFontInfo) VariableAxes() []*QFontVariableAxis {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontInfo_variableAxes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QFontVariableAxis, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQFontVariableAxis(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFontInfo) LegacyWeight() int {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_legacyWeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontInfo) ExactMatch() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontInfo_exactMatch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
