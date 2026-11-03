package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGlyphRun_new uintptr
	addr_QGlyphRun_new2 uintptr
	addr_QGlyphRun_OperatorAssign uintptr
	addr_QGlyphRun_swap uintptr
	addr_QGlyphRun_rawFont uintptr
	addr_QGlyphRun_setRawFont uintptr
	addr_QGlyphRun_glyphIndexes uintptr
	addr_QGlyphRun_setGlyphIndexes uintptr
	addr_QGlyphRun_positions uintptr
	addr_QGlyphRun_setPositions uintptr
	addr_QGlyphRun_clear uintptr
	addr_QGlyphRun_OperatorEqual uintptr
	addr_QGlyphRun_OperatorNotEqual uintptr
	addr_QGlyphRun_setOverline uintptr
	addr_QGlyphRun_overline uintptr
	addr_QGlyphRun_setUnderline uintptr
	addr_QGlyphRun_underline uintptr
	addr_QGlyphRun_setStrikeOut uintptr
	addr_QGlyphRun_strikeOut uintptr
	addr_QGlyphRun_setRightToLeft uintptr
	addr_QGlyphRun_isRightToLeft uintptr
	addr_QGlyphRun_setFlag uintptr
	addr_QGlyphRun_setFlags uintptr
	addr_QGlyphRun_flags uintptr
	addr_QGlyphRun_setBoundingRect uintptr
	addr_QGlyphRun_boundingRect uintptr
	addr_QGlyphRun_stringIndexes uintptr
	addr_QGlyphRun_setStringIndexes uintptr
	addr_QGlyphRun_setSourceString uintptr
	addr_QGlyphRun_sourceString uintptr
	addr_QGlyphRun_isEmpty uintptr
	addr_QGlyphRun_setFlag2 uintptr
	pg_QGlyphRun_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGlyphRun_new, "QGlyphRun_new")
	symaddr(&addr_QGlyphRun_new2, "QGlyphRun_new2")
	symaddr(&addr_QGlyphRun_OperatorAssign, "QGlyphRun_OperatorAssign")
	symaddr(&addr_QGlyphRun_swap, "QGlyphRun_swap")
	symaddr(&addr_QGlyphRun_rawFont, "QGlyphRun_rawFont")
	symaddr(&addr_QGlyphRun_setRawFont, "QGlyphRun_setRawFont")
	symaddr(&addr_QGlyphRun_glyphIndexes, "QGlyphRun_glyphIndexes")
	symaddr(&addr_QGlyphRun_setGlyphIndexes, "QGlyphRun_setGlyphIndexes")
	symaddr(&addr_QGlyphRun_positions, "QGlyphRun_positions")
	symaddr(&addr_QGlyphRun_setPositions, "QGlyphRun_setPositions")
	symaddr(&addr_QGlyphRun_clear, "QGlyphRun_clear")
	symaddr(&addr_QGlyphRun_OperatorEqual, "QGlyphRun_OperatorEqual")
	symaddr(&addr_QGlyphRun_OperatorNotEqual, "QGlyphRun_OperatorNotEqual")
	symaddr(&addr_QGlyphRun_setOverline, "QGlyphRun_setOverline")
	symaddr(&addr_QGlyphRun_overline, "QGlyphRun_overline")
	symaddr(&addr_QGlyphRun_setUnderline, "QGlyphRun_setUnderline")
	symaddr(&addr_QGlyphRun_underline, "QGlyphRun_underline")
	symaddr(&addr_QGlyphRun_setStrikeOut, "QGlyphRun_setStrikeOut")
	symaddr(&addr_QGlyphRun_strikeOut, "QGlyphRun_strikeOut")
	symaddr(&addr_QGlyphRun_setRightToLeft, "QGlyphRun_setRightToLeft")
	symaddr(&addr_QGlyphRun_isRightToLeft, "QGlyphRun_isRightToLeft")
	symaddr(&addr_QGlyphRun_setFlag, "QGlyphRun_setFlag")
	symaddr(&addr_QGlyphRun_setFlags, "QGlyphRun_setFlags")
	symaddr(&addr_QGlyphRun_flags, "QGlyphRun_flags")
	symaddr(&addr_QGlyphRun_setBoundingRect, "QGlyphRun_setBoundingRect")
	symaddr(&addr_QGlyphRun_boundingRect, "QGlyphRun_boundingRect")
	symaddr(&addr_QGlyphRun_stringIndexes, "QGlyphRun_stringIndexes")
	symaddr(&addr_QGlyphRun_setStringIndexes, "QGlyphRun_setStringIndexes")
	symaddr(&addr_QGlyphRun_setSourceString, "QGlyphRun_setSourceString")
	symaddr(&addr_QGlyphRun_sourceString, "QGlyphRun_sourceString")
	symaddr(&addr_QGlyphRun_isEmpty, "QGlyphRun_isEmpty")
	symaddr(&addr_QGlyphRun_setFlag2, "QGlyphRun_setFlag2")
	reg(&pg_QGlyphRun_delete, "QGlyphRun_delete")
}

type QGlyphRun struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGlyphRun) UnsafeGoID() int64     { return this.vid }
func (this *QGlyphRun) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGlyphRun(h uintptr) *QGlyphRun {
	if h == 0 {
		return nil
	}
	return &QGlyphRun{h: h}
}

func newownedQGlyphRun(h uintptr) *QGlyphRun {
	o := bindQGlyphRun(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGlyphRun).Delete)
	}
	return o
}

func (this *QGlyphRun) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGlyphRun) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGlyphRun(h unsafe.Pointer) *QGlyphRun { return bindQGlyphRun(uintptr(h)) }
func UnsafeNewQGlyphRunOwned(h unsafe.Pointer) *QGlyphRun { return newownedQGlyphRun(uintptr(h)) }

func (this *QGlyphRun) UnsafeRelease() { this.h = 0 }

func NewQGlyphRun() *QGlyphRun {
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_new)
	__o := newownedQGlyphRun(__r)
	return __o
}

func NewQGlyphRun2(other *QGlyphRun) *QGlyphRun {
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_new2, uintptr(other.UnsafePointer()))
	__o := newownedQGlyphRun(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QGlyphRun) OperatorAssign(other *QGlyphRun) *QGlyphRun {
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQGlyphRun(__r)
}

func (this *QGlyphRun) Swap(other *QGlyphRun) {
	purego.SyscallN(addr_QGlyphRun_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QGlyphRun) RawFont() *QRawFont {
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_rawFont, this.h)
	runtime.KeepAlive(this)
	return newownedQRawFont(__r)
}

func (this *QGlyphRun) SetRawFont(rawFont *QRawFont) {
	purego.SyscallN(addr_QGlyphRun_setRawFont, this.h, uintptr(rawFont.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rawFont)
}

func (this *QGlyphRun) GlyphIndexes() []uint {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_glyphIndexes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]uint, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = uint(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGlyphRun) SetGlyphIndexes(glyphIndexes []uint) {
	glyphIndexes_arr := make([]int32, len(glyphIndexes))
	for __i, __e := range glyphIndexes {
		glyphIndexes_arr[__i] = int32(__e)
	}
	var glyphIndexes_p uintptr
	if len(glyphIndexes_arr) > 0 {
		glyphIndexes_p = uintptr(unsafe.Pointer(&glyphIndexes_arr[0]))
	}
	purego.SyscallN(addr_QGlyphRun_setGlyphIndexes, this.h, glyphIndexes_p, uintptr(len(glyphIndexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(glyphIndexes_arr)
}

func (this *QGlyphRun) Positions() []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_positions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGlyphRun) SetPositions(positions []*QPointF) {
	positions_arr := make([]uintptr, len(positions))
	for __i, __e := range positions {
		positions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var positions_p uintptr
	if len(positions_arr) > 0 {
		positions_p = uintptr(unsafe.Pointer(&positions_arr[0]))
	}
	purego.SyscallN(addr_QGlyphRun_setPositions, this.h, positions_p, uintptr(len(positions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(positions)
	runtime.KeepAlive(positions_arr)
}

func (this *QGlyphRun) Clear() {
	purego.SyscallN(addr_QGlyphRun_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) OperatorEqual(other *QGlyphRun) bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGlyphRun) OperatorNotEqual(other *QGlyphRun) bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGlyphRun) SetOverline(overline bool) {
	purego.SyscallN(addr_QGlyphRun_setOverline, this.h, uintptr(b2i(overline)))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) Overline() bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_overline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGlyphRun) SetUnderline(underline bool) {
	purego.SyscallN(addr_QGlyphRun_setUnderline, this.h, uintptr(b2i(underline)))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) Underline() bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_underline, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGlyphRun) SetStrikeOut(strikeOut bool) {
	purego.SyscallN(addr_QGlyphRun_setStrikeOut, this.h, uintptr(b2i(strikeOut)))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) StrikeOut() bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_strikeOut, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGlyphRun) SetRightToLeft(on bool) {
	purego.SyscallN(addr_QGlyphRun_setRightToLeft, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_isRightToLeft, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGlyphRun) SetFlag(flag QGlyphRun__GlyphRunFlag) {
	purego.SyscallN(addr_QGlyphRun_setFlag, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) SetFlags(flags QGlyphRun__GlyphRunFlag) {
	purego.SyscallN(addr_QGlyphRun_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) Flags() QGlyphRun__GlyphRunFlag {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGlyphRun__GlyphRunFlag(__r)
}

func (this *QGlyphRun) SetBoundingRect(boundingRect *QRectF) {
	purego.SyscallN(addr_QGlyphRun_setBoundingRect, this.h, uintptr(boundingRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(boundingRect)
}

func (this *QGlyphRun) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGlyphRun) StringIndexes() []int64 {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_stringIndexes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int64, __n)
	__sl := unsafe.Slice((*int64)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = __sl[__i]
	}
	pgFree(__r)
	return __out
}

func (this *QGlyphRun) SetStringIndexes(stringIndexes []int64) {
	stringIndexes_arr := make([]int64, len(stringIndexes))
	for __i, __e := range stringIndexes {
		stringIndexes_arr[__i] = __e
	}
	var stringIndexes_p uintptr
	if len(stringIndexes_arr) > 0 {
		stringIndexes_p = uintptr(unsafe.Pointer(&stringIndexes_arr[0]))
	}
	purego.SyscallN(addr_QGlyphRun_setStringIndexes, this.h, stringIndexes_p, uintptr(len(stringIndexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stringIndexes_arr)
}

func (this *QGlyphRun) SetSourceString(sourceString string) {
	var sourceString_p uintptr
	if len(sourceString) > 0 {
		sourceString_p = uintptr(unsafe.Pointer(unsafe.StringData(sourceString)))
	}
	purego.SyscallN(addr_QGlyphRun_setSourceString, this.h, sourceString_p, uintptr(len(sourceString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceString)
}

func (this *QGlyphRun) SourceString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGlyphRun_sourceString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGlyphRun) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QGlyphRun_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGlyphRun) SetFlag2(flag QGlyphRun__GlyphRunFlag, enabled bool) {
	purego.SyscallN(addr_QGlyphRun_setFlag2, this.h, uintptr(flag), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGlyphRun) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGlyphRun_delete(this.h)
	}
	this.UnsafeRelease()
}
