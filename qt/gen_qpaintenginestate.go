package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPaintEngineState_state uintptr
	addr_QPaintEngineState_pen uintptr
	addr_QPaintEngineState_brush uintptr
	addr_QPaintEngineState_brushOrigin uintptr
	addr_QPaintEngineState_backgroundBrush uintptr
	addr_QPaintEngineState_backgroundMode uintptr
	addr_QPaintEngineState_font uintptr
	addr_QPaintEngineState_transform uintptr
	addr_QPaintEngineState_clipOperation uintptr
	addr_QPaintEngineState_clipRegion uintptr
	addr_QPaintEngineState_clipPath uintptr
	addr_QPaintEngineState_isClipEnabled uintptr
	addr_QPaintEngineState_renderHints uintptr
	addr_QPaintEngineState_compositionMode uintptr
	pg_QPaintEngineState_opacity func(uintptr) float64
	addr_QPaintEngineState_painter uintptr
	addr_QPaintEngineState_brushNeedsResolving uintptr
	addr_QPaintEngineState_penNeedsResolving uintptr
	pg_QPaintEngineState_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPaintEngineState_state, "QPaintEngineState_state")
	symaddr(&addr_QPaintEngineState_pen, "QPaintEngineState_pen")
	symaddr(&addr_QPaintEngineState_brush, "QPaintEngineState_brush")
	symaddr(&addr_QPaintEngineState_brushOrigin, "QPaintEngineState_brushOrigin")
	symaddr(&addr_QPaintEngineState_backgroundBrush, "QPaintEngineState_backgroundBrush")
	symaddr(&addr_QPaintEngineState_backgroundMode, "QPaintEngineState_backgroundMode")
	symaddr(&addr_QPaintEngineState_font, "QPaintEngineState_font")
	symaddr(&addr_QPaintEngineState_transform, "QPaintEngineState_transform")
	symaddr(&addr_QPaintEngineState_clipOperation, "QPaintEngineState_clipOperation")
	symaddr(&addr_QPaintEngineState_clipRegion, "QPaintEngineState_clipRegion")
	symaddr(&addr_QPaintEngineState_clipPath, "QPaintEngineState_clipPath")
	symaddr(&addr_QPaintEngineState_isClipEnabled, "QPaintEngineState_isClipEnabled")
	symaddr(&addr_QPaintEngineState_renderHints, "QPaintEngineState_renderHints")
	symaddr(&addr_QPaintEngineState_compositionMode, "QPaintEngineState_compositionMode")
	reg(&pg_QPaintEngineState_opacity, "QPaintEngineState_opacity")
	symaddr(&addr_QPaintEngineState_painter, "QPaintEngineState_painter")
	symaddr(&addr_QPaintEngineState_brushNeedsResolving, "QPaintEngineState_brushNeedsResolving")
	symaddr(&addr_QPaintEngineState_penNeedsResolving, "QPaintEngineState_penNeedsResolving")
	reg(&pg_QPaintEngineState_delete, "QPaintEngineState_delete")
}

type QPaintEngineState struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPaintEngineState) UnsafeGoID() int64     { return this.vid }
func (this *QPaintEngineState) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPaintEngineState(h uintptr) *QPaintEngineState {
	if h == 0 {
		return nil
	}
	return &QPaintEngineState{h: h}
}

func newownedQPaintEngineState(h uintptr) *QPaintEngineState {
	o := bindQPaintEngineState(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPaintEngineState).Delete)
	}
	return o
}

func (this *QPaintEngineState) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPaintEngineState) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPaintEngineState(h unsafe.Pointer) *QPaintEngineState { return bindQPaintEngineState(uintptr(h)) }
func UnsafeNewQPaintEngineStateOwned(h unsafe.Pointer) *QPaintEngineState { return newownedQPaintEngineState(uintptr(h)) }

func (this *QPaintEngineState) UnsafeRelease() { this.h = 0 }

func (this *QPaintEngineState) State() QPaintEngine__DirtyFlag {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPaintEngine__DirtyFlag(__r)
}

func (this *QPaintEngineState) Pen() *QPen {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_pen, this.h)
	runtime.KeepAlive(this)
	return newownedQPen(__r)
}

func (this *QPaintEngineState) Brush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_brush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QPaintEngineState) BrushOrigin() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_brushOrigin, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPaintEngineState) BackgroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_backgroundBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QPaintEngineState) BackgroundMode() BGMode {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_backgroundMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return BGMode(__r)
}

func (this *QPaintEngineState) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QPaintEngineState) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_transform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QPaintEngineState) ClipOperation() ClipOperation {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_clipOperation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ClipOperation(__r)
}

func (this *QPaintEngineState) ClipRegion() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_clipRegion, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QPaintEngineState) ClipPath() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_clipPath, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPaintEngineState) IsClipEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_isClipEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngineState) RenderHints() QPainter__RenderHint {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_renderHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainter__RenderHint(__r)
}

func (this *QPaintEngineState) CompositionMode() QPainter__CompositionMode {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_compositionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainter__CompositionMode(__r)
}

func (this *QPaintEngineState) Opacity() float64 {
	__r := pg_QPaintEngineState_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPaintEngineState) Painter() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPaintEngineState_painter, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QPaintEngineState) BrushNeedsResolving() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_brushNeedsResolving, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngineState) PenNeedsResolving() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintEngineState_penNeedsResolving, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintEngineState) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPaintEngineState_delete(this.h)
	}
	this.UnsafeRelease()
}
