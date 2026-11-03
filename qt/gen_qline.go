package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLine_new uintptr
	addr_QLine_new2 uintptr
	addr_QLine_new3 uintptr
	addr_QLine_new4 uintptr
	addr_QLine_isNull uintptr
	addr_QLine_p1 uintptr
	addr_QLine_p2 uintptr
	addr_QLine_x1 uintptr
	addr_QLine_y1 uintptr
	addr_QLine_x2 uintptr
	addr_QLine_y2 uintptr
	addr_QLine_dx uintptr
	addr_QLine_dy uintptr
	addr_QLine_translate uintptr
	addr_QLine_translate2 uintptr
	addr_QLine_translated uintptr
	addr_QLine_translated2 uintptr
	addr_QLine_center uintptr
	addr_QLine_setP1 uintptr
	addr_QLine_setP2 uintptr
	addr_QLine_setPoints uintptr
	addr_QLine_setLine uintptr
	addr_QLine_toLineF uintptr
	addr_QLine_OperatorEqual uintptr
	addr_QLine_OperatorNotEqual uintptr
	addr_QLine_OperatorEqual2 uintptr
	addr_QLine_OperatorNotEqual2 uintptr
	addr_QLine_OperatorMultiply uintptr
	pg_QLine_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLine_new, "QLine_new")
	symaddr(&addr_QLine_new2, "QLine_new2")
	symaddr(&addr_QLine_new3, "QLine_new3")
	symaddr(&addr_QLine_new4, "QLine_new4")
	symaddr(&addr_QLine_isNull, "QLine_isNull")
	symaddr(&addr_QLine_p1, "QLine_p1")
	symaddr(&addr_QLine_p2, "QLine_p2")
	symaddr(&addr_QLine_x1, "QLine_x1")
	symaddr(&addr_QLine_y1, "QLine_y1")
	symaddr(&addr_QLine_x2, "QLine_x2")
	symaddr(&addr_QLine_y2, "QLine_y2")
	symaddr(&addr_QLine_dx, "QLine_dx")
	symaddr(&addr_QLine_dy, "QLine_dy")
	symaddr(&addr_QLine_translate, "QLine_translate")
	symaddr(&addr_QLine_translate2, "QLine_translate2")
	symaddr(&addr_QLine_translated, "QLine_translated")
	symaddr(&addr_QLine_translated2, "QLine_translated2")
	symaddr(&addr_QLine_center, "QLine_center")
	symaddr(&addr_QLine_setP1, "QLine_setP1")
	symaddr(&addr_QLine_setP2, "QLine_setP2")
	symaddr(&addr_QLine_setPoints, "QLine_setPoints")
	symaddr(&addr_QLine_setLine, "QLine_setLine")
	symaddr(&addr_QLine_toLineF, "QLine_toLineF")
	symaddr(&addr_QLine_OperatorEqual, "QLine_OperatorEqual")
	symaddr(&addr_QLine_OperatorNotEqual, "QLine_OperatorNotEqual")
	symaddr(&addr_QLine_OperatorEqual2, "QLine_OperatorEqual2")
	symaddr(&addr_QLine_OperatorNotEqual2, "QLine_OperatorNotEqual2")
	symaddr(&addr_QLine_OperatorMultiply, "QLine_OperatorMultiply")
	reg(&pg_QLine_delete, "QLine_delete")
}

type QLine struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLine) UnsafeGoID() int64     { return this.vid }
func (this *QLine) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLine(h uintptr) *QLine {
	if h == 0 {
		return nil
	}
	return &QLine{h: h}
}

func newownedQLine(h uintptr) *QLine {
	o := bindQLine(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLine).Delete)
	}
	return o
}

func (this *QLine) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLine) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLine(h unsafe.Pointer) *QLine { return bindQLine(uintptr(h)) }
func UnsafeNewQLineOwned(h unsafe.Pointer) *QLine { return newownedQLine(uintptr(h)) }

func (this *QLine) UnsafeRelease() { this.h = 0 }

func NewQLine() *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_new)
	__o := newownedQLine(__r)
	return __o
}

func NewQLine2(pt1 *QPoint, pt2 *QPoint) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_new2, uintptr(pt1.UnsafePointer()), uintptr(pt2.UnsafePointer()))
	__o := newownedQLine(__r)
	runtime.KeepAlive(pt1)
	runtime.KeepAlive(pt2)
	return __o
}

func NewQLine3(x1 int, y1 int, x2 int, y2 int) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_new3, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	__o := newownedQLine(__r)
	return __o
}

func NewQLine4(arg1 *QLine) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQLine(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QLine) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QLine_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLine) P1() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QLine_p1, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QLine) P2() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QLine_p2, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QLine) X1() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_x1, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) Y1() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_y1, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) X2() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_x2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) Y2() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_y2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) Dx() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_dx, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) Dy() int {
	__r0, _, _ := purego.SyscallN(addr_QLine_dy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLine) Translate(p *QPoint) {
	purego.SyscallN(addr_QLine_translate, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QLine) Translate2(dx int, dy int) {
	purego.SyscallN(addr_QLine_translate2, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QLine) Translated(p *QPoint) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_translated, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return newownedQLine(__r)
}

func (this *QLine) Translated2(dx int, dy int) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_translated2, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
	return newownedQLine(__r)
}

func (this *QLine) Center() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QLine_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QLine) SetP1(p1 *QPoint) {
	purego.SyscallN(addr_QLine_setP1, this.h, uintptr(p1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
}

func (this *QLine) SetP2(p2 *QPoint) {
	purego.SyscallN(addr_QLine_setP2, this.h, uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p2)
}

func (this *QLine) SetPoints(p1 *QPoint, p2 *QPoint) {
	purego.SyscallN(addr_QLine_setPoints, this.h, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
}

func (this *QLine) SetLine(x1 int, y1 int, x2 int, y2 int) {
	purego.SyscallN(addr_QLine_setLine, this.h, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	runtime.KeepAlive(this)
}

func (this *QLine) ToLineF() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QLine_toLineF, this.h)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QLine) OperatorEqual(rhs *QLine) bool {
	__r0, _, _ := purego.SyscallN(addr_QLine_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLine) OperatorNotEqual(rhs *QLine) bool {
	__r0, _, _ := purego.SyscallN(addr_QLine_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLine) OperatorEqualWithRhs(rhs *QLineF) bool {
	__r0, _, _ := purego.SyscallN(addr_QLine_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLine) OperatorNotEqualWithRhs(rhs *QLineF) bool {
	__r0, _, _ := purego.SyscallN(addr_QLine_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QLine) OperatorMultiply(m *QTransform) *QLine {
	__r, _, _ := purego.SyscallN(addr_QLine_OperatorMultiply, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return newownedQLine(__r)
}

func (this *QLine) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLine_delete(this.h)
	}
	this.UnsafeRelease()
}
