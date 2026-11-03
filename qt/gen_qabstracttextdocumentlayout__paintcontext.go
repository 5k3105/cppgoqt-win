package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractTextDocumentLayout__PaintContext_new uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_new2 uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_OperatorAssign uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_get uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_set uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_palette_get uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_palette_set uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_clip_get uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_clip_set uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_selections_get uintptr
	addr_QAbstractTextDocumentLayout__PaintContext_selections_set uintptr
	pg_QAbstractTextDocumentLayout__PaintContext_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_new, "QAbstractTextDocumentLayout__PaintContext_new")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_new2, "QAbstractTextDocumentLayout__PaintContext_new2")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_OperatorAssign, "QAbstractTextDocumentLayout__PaintContext_OperatorAssign")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_get, "QAbstractTextDocumentLayout__PaintContext_cursorPosition_get")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_set, "QAbstractTextDocumentLayout__PaintContext_cursorPosition_set")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_palette_get, "QAbstractTextDocumentLayout__PaintContext_palette_get")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_palette_set, "QAbstractTextDocumentLayout__PaintContext_palette_set")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_clip_get, "QAbstractTextDocumentLayout__PaintContext_clip_get")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_clip_set, "QAbstractTextDocumentLayout__PaintContext_clip_set")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_selections_get, "QAbstractTextDocumentLayout__PaintContext_selections_get")
	symaddr(&addr_QAbstractTextDocumentLayout__PaintContext_selections_set, "QAbstractTextDocumentLayout__PaintContext_selections_set")
	reg(&pg_QAbstractTextDocumentLayout__PaintContext_delete, "QAbstractTextDocumentLayout__PaintContext_delete")
}

type QAbstractTextDocumentLayout__PaintContext struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractTextDocumentLayout__PaintContext) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractTextDocumentLayout__PaintContext) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractTextDocumentLayout__PaintContext(h uintptr) *QAbstractTextDocumentLayout__PaintContext {
	if h == 0 {
		return nil
	}
	return &QAbstractTextDocumentLayout__PaintContext{h: h}
}

func newownedQAbstractTextDocumentLayout__PaintContext(h uintptr) *QAbstractTextDocumentLayout__PaintContext {
	o := bindQAbstractTextDocumentLayout__PaintContext(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractTextDocumentLayout__PaintContext).Delete)
	}
	return o
}

func (this *QAbstractTextDocumentLayout__PaintContext) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractTextDocumentLayout__PaintContext) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractTextDocumentLayout__PaintContext(h unsafe.Pointer) *QAbstractTextDocumentLayout__PaintContext { return bindQAbstractTextDocumentLayout__PaintContext(uintptr(h)) }
func UnsafeNewQAbstractTextDocumentLayout__PaintContextOwned(h unsafe.Pointer) *QAbstractTextDocumentLayout__PaintContext { return newownedQAbstractTextDocumentLayout__PaintContext(uintptr(h)) }

func (this *QAbstractTextDocumentLayout__PaintContext) UnsafeRelease() { this.h = 0 }

func NewQAbstractTextDocumentLayout__PaintContext() *QAbstractTextDocumentLayout__PaintContext {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_new)
	__o := newownedQAbstractTextDocumentLayout__PaintContext(__r)
	return __o
}

func NewQAbstractTextDocumentLayout__PaintContext2(arg1 *QAbstractTextDocumentLayout__PaintContext) *QAbstractTextDocumentLayout__PaintContext {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQAbstractTextDocumentLayout__PaintContext(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAbstractTextDocumentLayout__PaintContext) OperatorAssign(arg1 *QAbstractTextDocumentLayout__PaintContext) *QAbstractTextDocumentLayout__PaintContext {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAbstractTextDocumentLayout__PaintContext(__r)
}

func (this *QAbstractTextDocumentLayout__PaintContext) CursorPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractTextDocumentLayout__PaintContext) SetCursorPosition(cursorPosition int) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_cursorPosition_set, this.h, uintptr(cursorPosition))
	runtime.KeepAlive(this)
}

func (this *QAbstractTextDocumentLayout__PaintContext) Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_palette_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QAbstractTextDocumentLayout__PaintContext) SetPalette(palette *QPalette) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_palette_set, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QAbstractTextDocumentLayout__PaintContext) Clip() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_clip_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QAbstractTextDocumentLayout__PaintContext) SetClip(clip *QRectF) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_clip_set, this.h, uintptr(clip.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(clip)
}

func (this *QAbstractTextDocumentLayout__PaintContext) Selections() []*QAbstractTextDocumentLayout__Selection {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_selections_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAbstractTextDocumentLayout__Selection, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQAbstractTextDocumentLayout__Selection(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractTextDocumentLayout__PaintContext) SetSelections(selections []*QAbstractTextDocumentLayout__Selection) {
	selections_arr := make([]uintptr, len(selections))
	for __i, __e := range selections {
		selections_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var selections_p uintptr
	if len(selections_arr) > 0 {
		selections_p = uintptr(unsafe.Pointer(&selections_arr[0]))
	}
	purego.SyscallN(addr_QAbstractTextDocumentLayout__PaintContext_selections_set, this.h, selections_p, uintptr(len(selections)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selections)
	runtime.KeepAlive(selections_arr)
}

func (this *QAbstractTextDocumentLayout__PaintContext) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractTextDocumentLayout__PaintContext_delete(this.h)
	}
	this.UnsafeRelease()
}
