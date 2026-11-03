package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRect_new uintptr
	addr_QRect_new2 uintptr
	addr_QRect_new3 uintptr
	addr_QRect_new4 uintptr
	addr_QRect_new5 uintptr
	addr_QRect_isNull uintptr
	addr_QRect_isEmpty uintptr
	addr_QRect_isValid uintptr
	addr_QRect_left uintptr
	addr_QRect_top uintptr
	addr_QRect_right uintptr
	addr_QRect_bottom uintptr
	addr_QRect_normalized uintptr
	addr_QRect_x uintptr
	addr_QRect_y uintptr
	addr_QRect_setLeft uintptr
	addr_QRect_setTop uintptr
	addr_QRect_setRight uintptr
	addr_QRect_setBottom uintptr
	addr_QRect_setX uintptr
	addr_QRect_setY uintptr
	addr_QRect_setTopLeft uintptr
	addr_QRect_setBottomRight uintptr
	addr_QRect_setTopRight uintptr
	addr_QRect_setBottomLeft uintptr
	addr_QRect_topLeft uintptr
	addr_QRect_bottomRight uintptr
	addr_QRect_topRight uintptr
	addr_QRect_bottomLeft uintptr
	addr_QRect_center uintptr
	addr_QRect_moveLeft uintptr
	addr_QRect_moveTop uintptr
	addr_QRect_moveRight uintptr
	addr_QRect_moveBottom uintptr
	addr_QRect_moveTopLeft uintptr
	addr_QRect_moveBottomRight uintptr
	addr_QRect_moveTopRight uintptr
	addr_QRect_moveBottomLeft uintptr
	addr_QRect_moveCenter uintptr
	addr_QRect_translate uintptr
	addr_QRect_translate2 uintptr
	addr_QRect_translated uintptr
	addr_QRect_translated2 uintptr
	addr_QRect_transposed uintptr
	addr_QRect_moveTo uintptr
	addr_QRect_moveTo2 uintptr
	addr_QRect_setRect uintptr
	addr_QRect_getRect uintptr
	addr_QRect_setCoords uintptr
	addr_QRect_getCoords uintptr
	addr_QRect_adjust uintptr
	addr_QRect_adjusted uintptr
	addr_QRect_size uintptr
	addr_QRect_width uintptr
	addr_QRect_height uintptr
	addr_QRect_setWidth uintptr
	addr_QRect_setHeight uintptr
	addr_QRect_setSize uintptr
	addr_QRect_OperatorBitwiseOr uintptr
	addr_QRect_OperatorBitwiseAnd uintptr
	addr_QRect_OperatorBitwiseOrAssign uintptr
	addr_QRect_OperatorBitwiseAndAssign uintptr
	addr_QRect_contains uintptr
	addr_QRect_contains2 uintptr
	addr_QRect_contains3 uintptr
	addr_QRect_contains4 uintptr
	addr_QRect_united uintptr
	addr_QRect_intersected uintptr
	addr_QRect_intersects uintptr
	addr_QRect_marginsAdded uintptr
	addr_QRect_marginsRemoved uintptr
	addr_QRect_OperatorPlusAssign uintptr
	addr_QRect_OperatorMinusAssign uintptr
	addr_QRect_span uintptr
	addr_QRect_toRectF uintptr
	addr_QRect_OperatorAssign uintptr
	addr_QRect_contains5 uintptr
	addr_QRect_contains6 uintptr
	addr_QRect_OperatorEqual uintptr
	addr_QRect_OperatorNotEqual uintptr
	addr_QRect_OperatorPlus uintptr
	addr_QRect_OperatorMinus uintptr
	addr_QRect_OperatorEqual2 uintptr
	addr_QRect_OperatorNotEqual2 uintptr
	pg_QRect_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRect_new, "QRect_new")
	symaddr(&addr_QRect_new2, "QRect_new2")
	symaddr(&addr_QRect_new3, "QRect_new3")
	symaddr(&addr_QRect_new4, "QRect_new4")
	symaddr(&addr_QRect_new5, "QRect_new5")
	symaddr(&addr_QRect_isNull, "QRect_isNull")
	symaddr(&addr_QRect_isEmpty, "QRect_isEmpty")
	symaddr(&addr_QRect_isValid, "QRect_isValid")
	symaddr(&addr_QRect_left, "QRect_left")
	symaddr(&addr_QRect_top, "QRect_top")
	symaddr(&addr_QRect_right, "QRect_right")
	symaddr(&addr_QRect_bottom, "QRect_bottom")
	symaddr(&addr_QRect_normalized, "QRect_normalized")
	symaddr(&addr_QRect_x, "QRect_x")
	symaddr(&addr_QRect_y, "QRect_y")
	symaddr(&addr_QRect_setLeft, "QRect_setLeft")
	symaddr(&addr_QRect_setTop, "QRect_setTop")
	symaddr(&addr_QRect_setRight, "QRect_setRight")
	symaddr(&addr_QRect_setBottom, "QRect_setBottom")
	symaddr(&addr_QRect_setX, "QRect_setX")
	symaddr(&addr_QRect_setY, "QRect_setY")
	symaddr(&addr_QRect_setTopLeft, "QRect_setTopLeft")
	symaddr(&addr_QRect_setBottomRight, "QRect_setBottomRight")
	symaddr(&addr_QRect_setTopRight, "QRect_setTopRight")
	symaddr(&addr_QRect_setBottomLeft, "QRect_setBottomLeft")
	symaddr(&addr_QRect_topLeft, "QRect_topLeft")
	symaddr(&addr_QRect_bottomRight, "QRect_bottomRight")
	symaddr(&addr_QRect_topRight, "QRect_topRight")
	symaddr(&addr_QRect_bottomLeft, "QRect_bottomLeft")
	symaddr(&addr_QRect_center, "QRect_center")
	symaddr(&addr_QRect_moveLeft, "QRect_moveLeft")
	symaddr(&addr_QRect_moveTop, "QRect_moveTop")
	symaddr(&addr_QRect_moveRight, "QRect_moveRight")
	symaddr(&addr_QRect_moveBottom, "QRect_moveBottom")
	symaddr(&addr_QRect_moveTopLeft, "QRect_moveTopLeft")
	symaddr(&addr_QRect_moveBottomRight, "QRect_moveBottomRight")
	symaddr(&addr_QRect_moveTopRight, "QRect_moveTopRight")
	symaddr(&addr_QRect_moveBottomLeft, "QRect_moveBottomLeft")
	symaddr(&addr_QRect_moveCenter, "QRect_moveCenter")
	symaddr(&addr_QRect_translate, "QRect_translate")
	symaddr(&addr_QRect_translate2, "QRect_translate2")
	symaddr(&addr_QRect_translated, "QRect_translated")
	symaddr(&addr_QRect_translated2, "QRect_translated2")
	symaddr(&addr_QRect_transposed, "QRect_transposed")
	symaddr(&addr_QRect_moveTo, "QRect_moveTo")
	symaddr(&addr_QRect_moveTo2, "QRect_moveTo2")
	symaddr(&addr_QRect_setRect, "QRect_setRect")
	symaddr(&addr_QRect_getRect, "QRect_getRect")
	symaddr(&addr_QRect_setCoords, "QRect_setCoords")
	symaddr(&addr_QRect_getCoords, "QRect_getCoords")
	symaddr(&addr_QRect_adjust, "QRect_adjust")
	symaddr(&addr_QRect_adjusted, "QRect_adjusted")
	symaddr(&addr_QRect_size, "QRect_size")
	symaddr(&addr_QRect_width, "QRect_width")
	symaddr(&addr_QRect_height, "QRect_height")
	symaddr(&addr_QRect_setWidth, "QRect_setWidth")
	symaddr(&addr_QRect_setHeight, "QRect_setHeight")
	symaddr(&addr_QRect_setSize, "QRect_setSize")
	symaddr(&addr_QRect_OperatorBitwiseOr, "QRect_OperatorBitwiseOr")
	symaddr(&addr_QRect_OperatorBitwiseAnd, "QRect_OperatorBitwiseAnd")
	symaddr(&addr_QRect_OperatorBitwiseOrAssign, "QRect_OperatorBitwiseOrAssign")
	symaddr(&addr_QRect_OperatorBitwiseAndAssign, "QRect_OperatorBitwiseAndAssign")
	symaddr(&addr_QRect_contains, "QRect_contains")
	symaddr(&addr_QRect_contains2, "QRect_contains2")
	symaddr(&addr_QRect_contains3, "QRect_contains3")
	symaddr(&addr_QRect_contains4, "QRect_contains4")
	symaddr(&addr_QRect_united, "QRect_united")
	symaddr(&addr_QRect_intersected, "QRect_intersected")
	symaddr(&addr_QRect_intersects, "QRect_intersects")
	symaddr(&addr_QRect_marginsAdded, "QRect_marginsAdded")
	symaddr(&addr_QRect_marginsRemoved, "QRect_marginsRemoved")
	symaddr(&addr_QRect_OperatorPlusAssign, "QRect_OperatorPlusAssign")
	symaddr(&addr_QRect_OperatorMinusAssign, "QRect_OperatorMinusAssign")
	symaddr(&addr_QRect_span, "QRect_span")
	symaddr(&addr_QRect_toRectF, "QRect_toRectF")
	symaddr(&addr_QRect_OperatorAssign, "QRect_OperatorAssign")
	symaddr(&addr_QRect_contains5, "QRect_contains5")
	symaddr(&addr_QRect_contains6, "QRect_contains6")
	symaddr(&addr_QRect_OperatorEqual, "QRect_OperatorEqual")
	symaddr(&addr_QRect_OperatorNotEqual, "QRect_OperatorNotEqual")
	symaddr(&addr_QRect_OperatorPlus, "QRect_OperatorPlus")
	symaddr(&addr_QRect_OperatorMinus, "QRect_OperatorMinus")
	symaddr(&addr_QRect_OperatorEqual2, "QRect_OperatorEqual2")
	symaddr(&addr_QRect_OperatorNotEqual2, "QRect_OperatorNotEqual2")
	reg(&pg_QRect_delete, "QRect_delete")
}

type QRect struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRect) UnsafeGoID() int64     { return this.vid }
func (this *QRect) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRect(h uintptr) *QRect {
	if h == 0 {
		return nil
	}
	return &QRect{h: h}
}

func newownedQRect(h uintptr) *QRect {
	o := bindQRect(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRect).Delete)
	}
	return o
}

func (this *QRect) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRect) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRect(h unsafe.Pointer) *QRect { return bindQRect(uintptr(h)) }
func UnsafeNewQRectOwned(h unsafe.Pointer) *QRect { return newownedQRect(uintptr(h)) }

func (this *QRect) UnsafeRelease() { this.h = 0 }

func NewQRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_new)
	__o := newownedQRect(__r)
	return __o
}

func NewQRect2(topleft *QPoint, bottomright *QPoint) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_new2, uintptr(topleft.UnsafePointer()), uintptr(bottomright.UnsafePointer()))
	__o := newownedQRect(__r)
	runtime.KeepAlive(topleft)
	runtime.KeepAlive(bottomright)
	return __o
}

func NewQRect3(topleft *QPoint, size *QSize) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_new3, uintptr(topleft.UnsafePointer()), uintptr(size.UnsafePointer()))
	__o := newownedQRect(__r)
	runtime.KeepAlive(topleft)
	runtime.KeepAlive(size)
	return __o
}

func NewQRect4(left int, top int, width int, height int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_new4, uintptr(left), uintptr(top), uintptr(width), uintptr(height))
	__o := newownedQRect(__r)
	return __o
}

func NewQRect5(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQRect(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QRect) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRect) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRect) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRect) Left() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_left, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Top() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_top, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Right() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_right, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Bottom() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_bottom, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Normalized() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRect) X() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) SetLeft(pos int) {
	purego.SyscallN(addr_QRect_setLeft, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) SetTop(pos int) {
	purego.SyscallN(addr_QRect_setTop, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) SetRight(pos int) {
	purego.SyscallN(addr_QRect_setRight, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) SetBottom(pos int) {
	purego.SyscallN(addr_QRect_setBottom, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) SetX(x int) {
	purego.SyscallN(addr_QRect_setX, this.h, uintptr(x))
	runtime.KeepAlive(this)
}

func (this *QRect) SetY(y int) {
	purego.SyscallN(addr_QRect_setY, this.h, uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QRect) SetTopLeft(p *QPoint) {
	purego.SyscallN(addr_QRect_setTopLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) SetBottomRight(p *QPoint) {
	purego.SyscallN(addr_QRect_setBottomRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) SetTopRight(p *QPoint) {
	purego.SyscallN(addr_QRect_setTopRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) SetBottomLeft(p *QPoint) {
	purego.SyscallN(addr_QRect_setBottomLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) TopLeft() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QRect_topLeft, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QRect) BottomRight() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QRect_bottomRight, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QRect) TopRight() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QRect_topRight, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QRect) BottomLeft() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QRect_bottomLeft, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QRect) Center() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QRect_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QRect) MoveLeft(pos int) {
	purego.SyscallN(addr_QRect_moveLeft, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) MoveTop(pos int) {
	purego.SyscallN(addr_QRect_moveTop, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) MoveRight(pos int) {
	purego.SyscallN(addr_QRect_moveRight, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) MoveBottom(pos int) {
	purego.SyscallN(addr_QRect_moveBottom, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QRect) MoveTopLeft(p *QPoint) {
	purego.SyscallN(addr_QRect_moveTopLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) MoveBottomRight(p *QPoint) {
	purego.SyscallN(addr_QRect_moveBottomRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) MoveTopRight(p *QPoint) {
	purego.SyscallN(addr_QRect_moveTopRight, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) MoveBottomLeft(p *QPoint) {
	purego.SyscallN(addr_QRect_moveBottomLeft, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) MoveCenter(p *QPoint) {
	purego.SyscallN(addr_QRect_moveCenter, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) Translate(dx int, dy int) {
	purego.SyscallN(addr_QRect_translate, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QRect) TranslateWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QRect_translate2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) Translated(dx int, dy int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_translated, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRect) TranslatedWithQPoint(p *QPoint) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_translated2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQRect(__r)
}

func (this *QRect) Transposed() *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRect) MoveTo(x int, t int) {
	purego.SyscallN(addr_QRect_moveTo, this.h, uintptr(x), uintptr(t))
	runtime.KeepAlive(this)
}

func (this *QRect) MoveToWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QRect_moveTo2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRect) SetRect(x int, y int, w int, h int) {
	purego.SyscallN(addr_QRect_setRect, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QRect) GetRect() (int, int, int, int) {
	var x_out int32
	var y_out int32
	var w_out int32
	var h_out int32
	purego.SyscallN(addr_QRect_getRect, this.h, uintptr(unsafe.Pointer(&x_out)), uintptr(unsafe.Pointer(&y_out)), uintptr(unsafe.Pointer(&w_out)), uintptr(unsafe.Pointer(&h_out)))
	runtime.KeepAlive(this)
	return int(x_out), int(y_out), int(w_out), int(h_out)
}

func (this *QRect) SetCoords(x1 int, y1 int, x2 int, y2 int) {
	purego.SyscallN(addr_QRect_setCoords, this.h, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	runtime.KeepAlive(this)
}

func (this *QRect) GetCoords() (int, int, int, int) {
	var x1_out int32
	var y1_out int32
	var x2_out int32
	var y2_out int32
	purego.SyscallN(addr_QRect_getCoords, this.h, uintptr(unsafe.Pointer(&x1_out)), uintptr(unsafe.Pointer(&y1_out)), uintptr(unsafe.Pointer(&x2_out)), uintptr(unsafe.Pointer(&y2_out)))
	runtime.KeepAlive(this)
	return int(x1_out), int(y1_out), int(x2_out), int(y2_out)
}

func (this *QRect) Adjust(x1 int, y1 int, x2 int, y2 int) {
	purego.SyscallN(addr_QRect_adjust, this.h, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	runtime.KeepAlive(this)
}

func (this *QRect) Adjusted(x1 int, y1 int, x2 int, y2 int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_adjusted, this.h, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QRect) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QRect_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QRect) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QRect_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRect) SetWidth(w int) {
	purego.SyscallN(addr_QRect_setWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QRect) SetHeight(h int) {
	purego.SyscallN(addr_QRect_setHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QRect) SetSize(s *QSize) {
	purego.SyscallN(addr_QRect_setSize, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QRect) OperatorBitwiseOr(r *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorBitwiseOr, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRect(__r)
}

func (this *QRect) OperatorBitwiseAnd(r *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorBitwiseAnd, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return newownedQRect(__r)
}

func (this *QRect) OperatorBitwiseOrAssign(r *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorBitwiseOrAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRect(__r)
}

func (this *QRect) OperatorBitwiseAndAssign(r *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorBitwiseAndAssign, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return bindQRect(__r)
}

func (this *QRect) Contains(r *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRect) ContainsWithQPoint(p *QPoint) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains2, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QRect) Contains2(x int, y int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains3, this.h, uintptr(x), uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRect) Contains3(x int, y int, proper bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains4, this.h, uintptr(x), uintptr(y), uintptr(b2i(proper)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRect) United(other *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_united, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQRect(__r)
}

func (this *QRect) Intersected(other *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_intersected, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQRect(__r)
}

func (this *QRect) Intersects(r *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_intersects, this.h, uintptr(r.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRect) MarginsAdded(margins *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_marginsAdded, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return newownedQRect(__r)
}

func (this *QRect) MarginsRemoved(margins *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_marginsRemoved, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return newownedQRect(__r)
}

func (this *QRect) OperatorPlusAssign(margins *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorPlusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQRect(__r)
}

func (this *QRect) OperatorMinusAssign(margins *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorMinusAssign, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return bindQRect(__r)
}

func QRect_Span(p1 *QPoint, p2 *QPoint) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_span, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
	return newownedQRect(__r)
}

func (this *QRect) ToRectF() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QRect_toRectF, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QRect) OperatorAssign(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQRect(__r)
}

func (this *QRect) Contains4(r *QRect, proper bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains5, this.h, uintptr(r.UnsafePointer()), uintptr(b2i(proper)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	return __r != 0
}

func (this *QRect) Contains5(p *QPoint, proper bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_contains6, this.h, uintptr(p.UnsafePointer()), uintptr(b2i(proper)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QRect) OperatorEqual(rhs *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRect) OperatorNotEqual(rhs *QRect) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRect) OperatorPlus(margins *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorPlus, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return newownedQRect(__r)
}

func (this *QRect) OperatorMinus(rhs *QMargins) *QRect {
	__r, _, _ := purego.SyscallN(addr_QRect_OperatorMinus, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return newownedQRect(__r)
}

func (this *QRect) OperatorEqualWithRhs(rhs *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRect) OperatorNotEqualWithRhs(rhs *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QRect_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QRect) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRect_delete(this.h)
	}
	this.UnsafeRelease()
}
