package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainter__PixmapFragment_new uintptr
	addr_QPainter__PixmapFragment_new2 uintptr
	addr_QPainter__PixmapFragment_create uintptr
	pg_QPainter__PixmapFragment_create2 func(uintptr, uintptr, float64) uintptr
	pg_QPainter__PixmapFragment_create3 func(uintptr, uintptr, float64, float64) uintptr
	pg_QPainter__PixmapFragment_create4 func(uintptr, uintptr, float64, float64, float64) uintptr
	pg_QPainter__PixmapFragment_create5 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QPainter__PixmapFragment_x_get func(uintptr) float64
	pg_QPainter__PixmapFragment_x_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_y_get func(uintptr) float64
	pg_QPainter__PixmapFragment_y_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_sourceLeft_get func(uintptr) float64
	pg_QPainter__PixmapFragment_sourceLeft_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_sourceTop_get func(uintptr) float64
	pg_QPainter__PixmapFragment_sourceTop_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_width_get func(uintptr) float64
	pg_QPainter__PixmapFragment_width_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_height_get func(uintptr) float64
	pg_QPainter__PixmapFragment_height_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_scaleX_get func(uintptr) float64
	pg_QPainter__PixmapFragment_scaleX_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_scaleY_get func(uintptr) float64
	pg_QPainter__PixmapFragment_scaleY_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_rotation_get func(uintptr) float64
	pg_QPainter__PixmapFragment_rotation_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_opacity_get func(uintptr) float64
	pg_QPainter__PixmapFragment_opacity_set func(uintptr, float64)
	pg_QPainter__PixmapFragment_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainter__PixmapFragment_new, "QPainter__PixmapFragment_new")
	symaddr(&addr_QPainter__PixmapFragment_new2, "QPainter__PixmapFragment_new2")
	symaddr(&addr_QPainter__PixmapFragment_create, "QPainter__PixmapFragment_create")
	reg(&pg_QPainter__PixmapFragment_create2, "QPainter__PixmapFragment_create2")
	reg(&pg_QPainter__PixmapFragment_create3, "QPainter__PixmapFragment_create3")
	reg(&pg_QPainter__PixmapFragment_create4, "QPainter__PixmapFragment_create4")
	reg(&pg_QPainter__PixmapFragment_create5, "QPainter__PixmapFragment_create5")
	reg(&pg_QPainter__PixmapFragment_x_get, "QPainter__PixmapFragment_x_get")
	reg(&pg_QPainter__PixmapFragment_x_set, "QPainter__PixmapFragment_x_set")
	reg(&pg_QPainter__PixmapFragment_y_get, "QPainter__PixmapFragment_y_get")
	reg(&pg_QPainter__PixmapFragment_y_set, "QPainter__PixmapFragment_y_set")
	reg(&pg_QPainter__PixmapFragment_sourceLeft_get, "QPainter__PixmapFragment_sourceLeft_get")
	reg(&pg_QPainter__PixmapFragment_sourceLeft_set, "QPainter__PixmapFragment_sourceLeft_set")
	reg(&pg_QPainter__PixmapFragment_sourceTop_get, "QPainter__PixmapFragment_sourceTop_get")
	reg(&pg_QPainter__PixmapFragment_sourceTop_set, "QPainter__PixmapFragment_sourceTop_set")
	reg(&pg_QPainter__PixmapFragment_width_get, "QPainter__PixmapFragment_width_get")
	reg(&pg_QPainter__PixmapFragment_width_set, "QPainter__PixmapFragment_width_set")
	reg(&pg_QPainter__PixmapFragment_height_get, "QPainter__PixmapFragment_height_get")
	reg(&pg_QPainter__PixmapFragment_height_set, "QPainter__PixmapFragment_height_set")
	reg(&pg_QPainter__PixmapFragment_scaleX_get, "QPainter__PixmapFragment_scaleX_get")
	reg(&pg_QPainter__PixmapFragment_scaleX_set, "QPainter__PixmapFragment_scaleX_set")
	reg(&pg_QPainter__PixmapFragment_scaleY_get, "QPainter__PixmapFragment_scaleY_get")
	reg(&pg_QPainter__PixmapFragment_scaleY_set, "QPainter__PixmapFragment_scaleY_set")
	reg(&pg_QPainter__PixmapFragment_rotation_get, "QPainter__PixmapFragment_rotation_get")
	reg(&pg_QPainter__PixmapFragment_rotation_set, "QPainter__PixmapFragment_rotation_set")
	reg(&pg_QPainter__PixmapFragment_opacity_get, "QPainter__PixmapFragment_opacity_get")
	reg(&pg_QPainter__PixmapFragment_opacity_set, "QPainter__PixmapFragment_opacity_set")
	reg(&pg_QPainter__PixmapFragment_delete, "QPainter__PixmapFragment_delete")
}

type QPainter__PixmapFragment struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainter__PixmapFragment) UnsafeGoID() int64     { return this.vid }
func (this *QPainter__PixmapFragment) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainter__PixmapFragment(h uintptr) *QPainter__PixmapFragment {
	if h == 0 {
		return nil
	}
	return &QPainter__PixmapFragment{h: h}
}

func newownedQPainter__PixmapFragment(h uintptr) *QPainter__PixmapFragment {
	o := bindQPainter__PixmapFragment(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainter__PixmapFragment).Delete)
	}
	return o
}

func (this *QPainter__PixmapFragment) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainter__PixmapFragment) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainter__PixmapFragment(h unsafe.Pointer) *QPainter__PixmapFragment { return bindQPainter__PixmapFragment(uintptr(h)) }
func UnsafeNewQPainter__PixmapFragmentOwned(h unsafe.Pointer) *QPainter__PixmapFragment { return newownedQPainter__PixmapFragment(uintptr(h)) }

func (this *QPainter__PixmapFragment) UnsafeRelease() { this.h = 0 }

func NewQPainter__PixmapFragment() *QPainter__PixmapFragment {
	__r, _, _ := purego.SyscallN(addr_QPainter__PixmapFragment_new)
	__o := newownedQPainter__PixmapFragment(__r)
	return __o
}

func NewQPainter__PixmapFragment2(arg1 *QPainter__PixmapFragment) *QPainter__PixmapFragment {
	__r, _, _ := purego.SyscallN(addr_QPainter__PixmapFragment_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPainter__PixmapFragment(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QPainter__PixmapFragment_Create(pos *QPointF, sourceRect *QRectF) *QPainter__PixmapFragment {
	__r, _, _ := purego.SyscallN(addr_QPainter__PixmapFragment_create, uintptr(pos.UnsafePointer()), uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(sourceRect)
	return newownedQPainter__PixmapFragment(__r)
}

func QPainter__PixmapFragment_Create2(pos *QPointF, sourceRect *QRectF, scaleX float64) *QPainter__PixmapFragment {
	__r := pg_QPainter__PixmapFragment_create2(uintptr(pos.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), scaleX)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(sourceRect)
	return newownedQPainter__PixmapFragment(__r)
}

func QPainter__PixmapFragment_Create3(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64) *QPainter__PixmapFragment {
	__r := pg_QPainter__PixmapFragment_create3(uintptr(pos.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), scaleX, scaleY)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(sourceRect)
	return newownedQPainter__PixmapFragment(__r)
}

func QPainter__PixmapFragment_Create4(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64, rotation float64) *QPainter__PixmapFragment {
	__r := pg_QPainter__PixmapFragment_create4(uintptr(pos.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), scaleX, scaleY, rotation)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(sourceRect)
	return newownedQPainter__PixmapFragment(__r)
}

func QPainter__PixmapFragment_Create5(pos *QPointF, sourceRect *QRectF, scaleX float64, scaleY float64, rotation float64, opacity float64) *QPainter__PixmapFragment {
	__r := pg_QPainter__PixmapFragment_create5(uintptr(pos.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), scaleX, scaleY, rotation, opacity)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(sourceRect)
	return newownedQPainter__PixmapFragment(__r)
}

func (this *QPainter__PixmapFragment) X() float64 {
	__r := pg_QPainter__PixmapFragment_x_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetX(x float64) {
	pg_QPainter__PixmapFragment_x_set(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Y() float64 {
	__r := pg_QPainter__PixmapFragment_y_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetY(y float64) {
	pg_QPainter__PixmapFragment_y_set(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) SourceLeft() float64 {
	__r := pg_QPainter__PixmapFragment_sourceLeft_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetSourceLeft(sourceLeft float64) {
	pg_QPainter__PixmapFragment_sourceLeft_set(this.h, sourceLeft)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) SourceTop() float64 {
	__r := pg_QPainter__PixmapFragment_sourceTop_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetSourceTop(sourceTop float64) {
	pg_QPainter__PixmapFragment_sourceTop_set(this.h, sourceTop)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Width() float64 {
	__r := pg_QPainter__PixmapFragment_width_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetWidth(width float64) {
	pg_QPainter__PixmapFragment_width_set(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Height() float64 {
	__r := pg_QPainter__PixmapFragment_height_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetHeight(height float64) {
	pg_QPainter__PixmapFragment_height_set(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) ScaleX() float64 {
	__r := pg_QPainter__PixmapFragment_scaleX_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetScaleX(scaleX float64) {
	pg_QPainter__PixmapFragment_scaleX_set(this.h, scaleX)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) ScaleY() float64 {
	__r := pg_QPainter__PixmapFragment_scaleY_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetScaleY(scaleY float64) {
	pg_QPainter__PixmapFragment_scaleY_set(this.h, scaleY)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Rotation() float64 {
	__r := pg_QPainter__PixmapFragment_rotation_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetRotation(rotation float64) {
	pg_QPainter__PixmapFragment_rotation_set(this.h, rotation)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Opacity() float64 {
	__r := pg_QPainter__PixmapFragment_opacity_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter__PixmapFragment) SetOpacity(opacity float64) {
	pg_QPainter__PixmapFragment_opacity_set(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QPainter__PixmapFragment) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainter__PixmapFragment_delete(this.h)
	}
	this.UnsafeRelease()
}
