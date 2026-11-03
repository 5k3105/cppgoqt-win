package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRectF_new uintptr
	addr_QRectF_new2 uintptr
	addr_QRectF_new3 uintptr
	pg_QRectF_new4 func(float64, float64, float64, float64) uintptr
	addr_QRectF_new5 uintptr
	addr_QRectF_new6 uintptr
	addr_QRectF_isNull uintptr
	addr_QRectF_isEmpty uintptr
	addr_QRectF_isValid uintptr
	addr_QRectF_normalized uintptr
	pg_QRectF_left func(uintptr) float64
	pg_QRectF_top func(uintptr) float64
	pg_QRectF_right func(uintptr) float64
	pg_QRectF_bottom func(uintptr) float64
	pg_QRectF_x func(uintptr) float64
	pg_QRectF_y func(uintptr) float64
	pg_QRectF_setLeft func(uintptr, float64)
	pg_QRectF_setTop func(uintptr, float64)
	pg_QRectF_setRight func(uintptr, float64)
	pg_QRectF_setBottom func(uintptr, float64)
	pg_QRectF_setX func(uintptr, float64)
	pg_QRectF_setY func(uintptr, float64)
	addr_QRectF_topLeft uintptr
	addr_QRectF_bottomRight uintptr
	addr_QRectF_topRight uintptr
	addr_QRectF_bottomLeft uintptr
	addr_QRectF_center uintptr
	addr_QRectF_setTopLeft uintptr
	addr_QRectF_setBottomRight uintptr
	addr_QRectF_setTopRight uintptr
	addr_QRectF_setBottomLeft uintptr
	pg_QRectF_moveLeft func(uintptr, float64)
	pg_QRectF_moveTop func(uintptr, float64)
	pg_QRectF_moveRight func(uintptr, float64)
	pg_QRectF_moveBottom func(uintptr, float64)
	addr_QRectF_moveTopLeft uintptr
	addr_QRectF_moveBottomRight uintptr
	addr_QRectF_moveTopRight uintptr
	addr_QRectF_moveBottomLeft uintptr
	addr_QRectF_moveCenter uintptr
	pg_QRectF_translate func(uintptr, float64, float64)
	addr_QRectF_translate2 uintptr
	pg_QRectF_translated func(uintptr, float64, float64) uintptr
	addr_QRectF_translated2 uintptr
	addr_QRectF_transposed uintptr
	pg_QRectF_moveTo func(uintptr, float64, float64)
	addr_QRectF_moveTo2 uintptr
	pg_QRectF_setRect func(uintptr, float64, float64, float64, float64)
	addr_QRectF_getRect uintptr
	pg_QRectF_setCoords func(uintptr, float64, float64, float64, float64)
	addr_QRectF_getCoords uintptr
	pg_QRectF_adjust func(uintptr, float64, float64, float64, float64)
	pg_QRectF_adjusted func(uintptr, float64, float64, float64, float64) uintptr
	addr_QRectF_size uintptr
	pg_QRectF_width func(uintptr) float64
	pg_QRectF_height func(uintptr) float64
	pg_QRectF_setWidth func(uintptr, float64)
	pg_QRectF_setHeight func(uintptr, float64)
	addr_QRectF_setSize uintptr
	addr_QRectF_OperatorBitwiseOr uintptr
	addr_QRectF_OperatorBitwiseAnd uintptr
	addr_QRectF_OperatorBitwiseOrAssign uintptr
	addr_QRectF_OperatorBitwiseAndAssign uintptr
	addr_QRectF_contains uintptr
	addr_QRectF_contains2 uintptr
	pg_QRectF_contains3 func(uintptr, float64, float64) int32
	addr_QRectF_united uintptr
	addr_QRectF_intersected uintptr
	addr_QRectF_intersects uintptr
	addr_QRectF_marginsAdded uintptr
	addr_QRectF_marginsRemoved uintptr
	addr_QRectF_OperatorPlusAssign uintptr
	addr_QRectF_OperatorMinusAssign uintptr
	addr_QRectF_toRect uintptr
	addr_QRectF_toAlignedRect uintptr
	addr_QRectF_OperatorAssign uintptr
	addr_QRectF_OperatorEqual uintptr
	addr_QRectF_OperatorNotEqual uintptr
	addr_QRectF_OperatorEqual2 uintptr
	addr_QRectF_OperatorNotEqual2 uintptr
	addr_QRectF_OperatorPlus uintptr
	addr_QRectF_OperatorMinus uintptr
	pg_QRectF_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRectF_new, "QRectF_new")
	symaddr(&addr_QRectF_new2, "QRectF_new2")
	symaddr(&addr_QRectF_new3, "QRectF_new3")
	reg(&pg_QRectF_new4, "QRectF_new4")
	symaddr(&addr_QRectF_new5, "QRectF_new5")
	symaddr(&addr_QRectF_new6, "QRectF_new6")
	symaddr(&addr_QRectF_isNull, "QRectF_isNull")
	symaddr(&addr_QRectF_isEmpty, "QRectF_isEmpty")
	symaddr(&addr_QRectF_isValid, "QRectF_isValid")
	symaddr(&addr_QRectF_normalized, "QRectF_normalized")
	reg(&pg_QRectF_left, "QRectF_left")
	reg(&pg_QRectF_top, "QRectF_top")
	reg(&pg_QRectF_right, "QRectF_right")
	reg(&pg_QRectF_bottom, "QRectF_bottom")
	reg(&pg_QRectF_x, "QRectF_x")
	reg(&pg_QRectF_y, "QRectF_y")
	reg(&pg_QRectF_setLeft, "QRectF_setLeft")
	reg(&pg_QRectF_setTop, "QRectF_setTop")
	reg(&pg_QRectF_setRight, "QRectF_setRight")
	reg(&pg_QRectF_setBottom, "QRectF_setBottom")
	reg(&pg_QRectF_setX, "QRectF_setX")
	reg(&pg_QRectF_setY, "QRectF_setY")
	symaddr(&addr_QRectF_topLeft, "QRectF_topLeft")
	symaddr(&addr_QRectF_bottomRight, "QRectF_bottomRight")
	symaddr(&addr_QRectF_topRight, "QRectF_topRight")
	symaddr(&addr_QRectF_bottomLeft, "QRectF_bottomLeft")
	symaddr(&addr_QRectF_center, "QRectF_center")
	symaddr(&addr_QRectF_setTopLeft, "QRectF_setTopLeft")
	symaddr(&addr_QRectF_setBottomRight, "QRectF_setBottomRight")
	symaddr(&addr_QRectF_setTopRight, "QRectF_setTopRight")
	symaddr(&addr_QRectF_setBottomLeft, "QRectF_setBottomLeft")
	reg(&pg_QRectF_moveLeft, "QRectF_moveLeft")
	reg(&pg_QRectF_moveTop, "QRectF_moveTop")
	reg(&pg_QRectF_moveRight, "QRectF_moveRight")
	reg(&pg_QRectF_moveBottom, "QRectF_moveBottom")
	symaddr(&addr_QRectF_moveTopLeft, "QRectF_moveTopLeft")
	symaddr(&addr_QRectF_moveBottomRight, "QRectF_moveBottomRight")
	symaddr(&addr_QRectF_moveTopRight, "QRectF_moveTopRight")
	symaddr(&addr_QRectF_moveBottomLeft, "QRectF_moveBottomLeft")
	symaddr(&addr_QRectF_moveCenter, "QRectF_moveCenter")
	reg(&pg_QRectF_translate, "QRectF_translate")
	symaddr(&addr_QRectF_translate2, "QRectF_translate2")
	reg(&pg_QRectF_translated, "QRectF_translated")
	symaddr(&addr_QRectF_translated2, "QRectF_translated2")
	symaddr(&addr_QRectF_transposed, "QRectF_transposed")
	reg(&pg_QRectF_moveTo, "QRectF_moveTo")
	symaddr(&addr_QRectF_moveTo2, "QRectF_moveTo2")
	reg(&pg_QRectF_setRect, "QRectF_setRect")
	symaddr(&addr_QRectF_getRect, "QRectF_getRect")
	reg(&pg_QRectF_setCoords, "QRectF_setCoords")
	symaddr(&addr_QRectF_getCoords, "QRectF_getCoords")
	reg(&pg_QRectF_adjust, "QRectF_adjust")
	reg(&pg_QRectF_adjusted, "QRectF_adjusted")
	symaddr(&addr_QRectF_size, "QRectF_size")
	reg(&pg_QRectF_width, "QRectF_width")
	reg(&pg_QRectF_height, "QRectF_height")
	reg(&pg_QRectF_setWidth, "QRectF_setWidth")
	reg(&pg_QRectF_setHeight, "QRectF_setHeight")
	symaddr(&addr_QRectF_setSize, "QRectF_setSize")
	symaddr(&addr_QRectF_OperatorBitwiseOr, "QRectF_OperatorBitwiseOr")
	symaddr(&addr_QRectF_OperatorBitwiseAnd, "QRectF_OperatorBitwiseAnd")
	symaddr(&addr_QRectF_OperatorBitwiseOrAssign, "QRectF_OperatorBitwiseOrAssign")
	symaddr(&addr_QRectF_OperatorBitwiseAndAssign, "QRectF_OperatorBitwiseAndAssign")
	symaddr(&addr_QRectF_contains, "QRectF_contains")
	symaddr(&addr_QRectF_contains2, "QRectF_contains2")
	reg(&pg_QRectF_contains3, "QRectF_contains3")
	symaddr(&addr_QRectF_united, "QRectF_united")
	symaddr(&addr_QRectF_intersected, "QRectF_intersected")
	symaddr(&addr_QRectF_intersects, "QRectF_intersects")
	symaddr(&addr_QRectF_marginsAdded, "QRectF_marginsAdded")
	symaddr(&addr_QRectF_marginsRemoved, "QRectF_marginsRemoved")
	symaddr(&addr_QRectF_OperatorPlusAssign, "QRectF_OperatorPlusAssign")
	symaddr(&addr_QRectF_OperatorMinusAssign, "QRectF_OperatorMinusAssign")
	symaddr(&addr_QRectF_toRect, "QRectF_toRect")
	symaddr(&addr_QRectF_toAlignedRect, "QRectF_toAlignedRect")
	symaddr(&addr_QRectF_OperatorAssign, "QRectF_OperatorAssign")
	symaddr(&addr_QRectF_OperatorEqual, "QRectF_OperatorEqual")
	symaddr(&addr_QRectF_OperatorNotEqual, "QRectF_OperatorNotEqual")
	symaddr(&addr_QRectF_OperatorEqual2, "QRectF_OperatorEqual2")
	symaddr(&addr_QRectF_OperatorNotEqual2, "QRectF_OperatorNotEqual2")
	symaddr(&addr_QRectF_OperatorPlus, "QRectF_OperatorPlus")
	symaddr(&addr_QRectF_OperatorMinus, "QRectF_OperatorMinus")
	reg(&pg_QRectF_delete, "QRectF_delete")
}

type QRectF struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRectF) UnsafeGoID() int64     { return this.vid }
func (this *QRectF) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRectF(h uintptr) *QRectF {
	if h == 0 {
		return nil
	}
	return &QRectF{h: h}
}

func newownedQRectF(h uintptr) *QRectF {
	o := bindQRectF(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRectF).Delete)
	}
	return o
}

func (this *QRectF) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRectF) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRectF(h unsafe.Pointer) *QRectF { return bindQRectF(uintptr(h)) }
func UnsafeNewQRectFOwned(h unsafe.Pointer) *QRectF { return newownedQRectF(uintptr(h)) }

func (this *QRectF) UnsafeRelease() { this.h = 0 }

func NewQRectF() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_new)
	__o := newownedQRectF(__r)
	return __o
}

func NewQRectF2(topleft *QPointF, size *QSizeF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_new2, uintptr(topleft.UnsafePointer()), uintptr(size.UnsafePointer()))
	__o := newownedQRectF(__r)
	runtime.KeepAlive(topleft)
	runtime.KeepAlive(size)
	return __o
}

func NewQRectF3(topleft *QPointF, bottomRight *QPointF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_new3, uintptr(topleft.UnsafePointer()), uintptr(bottomRight.UnsafePointer()))
	__o := newownedQRectF(__r)
	runtime.KeepAlive(topleft)
	runtime.KeepAlive(bottomRight)
	return __o
}

func NewQRectF4(left float64, top float64, width float64, height float64) *QRectF {
	__o := newownedQRectF(pg_QRectF_new4(left, top, width, height))
	return __o
}

func NewQRectF5(rect *QRect) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_new5, uintptr(rect.UnsafePointer()))
	__o := newownedQRectF(__r)
	runtime.KeepAlive(rect)
	return __o
}

func NewQRectF6(arg1 *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_new6, uintptr(arg1.UnsafePointer()))
	__o := newownedQRectF(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QRectF) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRectF) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRectF) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRectF) Normalized() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRectF) Left() float64 {
	__r := pg_QRectF_left(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) Top() float64 {
	__r := pg_QRectF_top(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) Right() float64 {
	__r := pg_QRectF_right(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) Bottom() float64 {
	__r := pg_QRectF_bottom(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) X() float64 {
	__r := pg_QRectF_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) Y() float64 {
	__r := pg_QRectF_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) SetLeft(pos float64) {
	pg_QRectF_setLeft(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetTop(pos float64) {
	pg_QRectF_setTop(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetRight(pos float64) {
	pg_QRectF_setRight(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetBottom(pos float64) {
	pg_QRectF_setBottom(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetX(pos float64) {
	pg_QRectF_setX(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetY(pos float64) {
	pg_QRectF_setY(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) TopLeft() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRectF_topLeft, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRectF) BottomRight() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRectF_bottomRight, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRectF) TopRight() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRectF_topRight, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRectF) BottomLeft() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRectF_bottomLeft, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRectF) Center() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRectF_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRectF) SetTopLeft(p *QPointF) {
	purego.SyscallN(addr_QRectF_setTopLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) SetBottomRight(p *QPointF) {
	purego.SyscallN(addr_QRectF_setBottomRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) SetTopRight(p *QPointF) {
	purego.SyscallN(addr_QRectF_setTopRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) SetBottomLeft(p *QPointF) {
	purego.SyscallN(addr_QRectF_setBottomLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) MoveLeft(pos float64) {
	pg_QRectF_moveLeft(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) MoveTop(pos float64) {
	pg_QRectF_moveTop(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) MoveRight(pos float64) {
	pg_QRectF_moveRight(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) MoveBottom(pos float64) {
	pg_QRectF_moveBottom(this.h, pos)
	runtime.KeepAlive(this)
}

func (this *QRectF) MoveTopLeft(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveTopLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) MoveBottomRight(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveBottomRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) MoveTopRight(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveTopRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) MoveBottomLeft(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveBottomLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) MoveCenter(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveCenter, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) Translate(dx float64, dy float64) {
	pg_QRectF_translate(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QRectF) TranslateWithQPointF(p *QPointF) {
	purego.SyscallN(addr_QRectF_translate2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) Translated(dx float64, dy float64) *QRectF {
	__r := pg_QRectF_translated(this.h, dx, dy)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRectF) TranslatedWithQPointF(p *QPointF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_translated2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQRectF(__r)
}

func (this *QRectF) Transposed() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRectF) MoveTo(x float64, y float64) {
	pg_QRectF_moveTo(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QRectF) MoveToWithQPointF(p *QPointF) {
	purego.SyscallN(addr_QRectF_moveTo2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRectF) SetRect(x float64, y float64, w float64, h float64) {
	pg_QRectF_setRect(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QRectF) GetRect() (float64, float64, float64, float64) {
	var x_out float64
	var y_out float64
	var w_out float64
	var h_out float64
	purego.SyscallN(addr_QRectF_getRect, this.h, uintptr(unsafe.Pointer(&x_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&w_out)), uintptr(unsafe.Pointer(&h_out)))
	runtime.KeepAlive(this)
	return x_out, y_out, w_out, h_out
}

func (this *QRectF) SetCoords(x1 float64, y1 float64, x2 float64, y2 float64) {
	pg_QRectF_setCoords(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
}

func (this *QRectF) GetCoords() (float64, float64, float64, float64) {
	var x1_out float64
	var y1_out float64
	var x2_out float64
	var y2_out float64
	purego.SyscallN(addr_QRectF_getCoords, this.h, uintptr(unsafe.Pointer(&x1_out)), uintptr(unsafe.Pointer(&y1_out)), uintptr(unsafe.Pointer(&x2_out)), uintptr(unsafe.Pointer(&y2_out)))
	runtime.KeepAlive(this)
	return x1_out, y1_out, x2_out, y2_out
}

func (this *QRectF) Adjust(x1 float64, y1 float64, x2 float64, y2 float64) {
	pg_QRectF_adjust(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
}

func (this *QRectF) Adjusted(x1 float64, y1 float64, x2 float64, y2 float64) *QRectF {
	__r := pg_QRectF_adjusted(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRectF) Size() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QRectF_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QRectF) Width() float64 {
	__r := pg_QRectF_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) Height() float64 {
	__r := pg_QRectF_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRectF) SetWidth(w float64) {
	pg_QRectF_setWidth(this.h, w)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetHeight(h float64) {
	pg_QRectF_setHeight(this.h, h)
	runtime.KeepAlive(this)
}

func (this *QRectF) SetSize(s *QSizeF) {
	purego.SyscallN(addr_QRectF_setSize, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QRectF) OperatorBitwiseOr(r *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorBitwiseOr, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRectF(__r)
}

func (this *QRectF) OperatorBitwiseAnd(r *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorBitwiseAnd, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRectF(__r)
}

func (this *QRectF) OperatorBitwiseOrAssign(r *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorBitwiseOrAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRectF(__r)
}

func (this *QRectF) OperatorBitwiseAndAssign(r *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorBitwiseAndAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRectF(__r)
}

func (this *QRectF) Contains(r *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_contains, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRectF) ContainsWithQPointF(p *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_contains2, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QRectF) Contains2(x float64, y float64) bool {
	__r := pg_QRectF_contains3(this.h, x, y)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRectF) United(other *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_united, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQRectF(__r)
}

func (this *QRectF) Intersected(other *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_intersected, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQRectF(__r)
}

func (this *QRectF) Intersects(r *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_intersects, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRectF) MarginsAdded(margins *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_marginsAdded, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return newownedQRectF(__r)
}

func (this *QRectF) MarginsRemoved(margins *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_marginsRemoved, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return newownedQRectF(__r)
}

func (this *QRectF) OperatorPlusAssign(margins *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorPlusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQRectF(__r)
}

func (this *QRectF) OperatorMinusAssign(margins *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorMinusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQRectF(__r)
}

func (this *QRectF) ToRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRectF_toRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRectF) ToAlignedRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRectF_toAlignedRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRectF) OperatorAssign(arg1 *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQRectF(__r)
}

func (this *QRectF) OperatorEqual(rhs *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRectF) OperatorNotEqual(rhs *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRectF) OperatorEqualWithRhs(rhs *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRectF) OperatorNotEqualWithRhs(rhs *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRectF_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRectF) OperatorPlus(rhs *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorPlus, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQRectF(__r)
}

func (this *QRectF) OperatorMinus(rhs *QMarginsF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRectF_OperatorMinus, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQRectF(__r)
}

func (this *QRectF) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRectF_delete(this.h)
	}
	this.UnsafeRelease()
}
