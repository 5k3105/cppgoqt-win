package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCursor_new uintptr
	addr_QCursor_new2 uintptr
	addr_QCursor_new3 uintptr
	addr_QCursor_new4 uintptr
	addr_QCursor_new5 uintptr
	addr_QCursor_new6 uintptr
	addr_QCursor_new7 uintptr
	addr_QCursor_new8 uintptr
	addr_QCursor_new9 uintptr
	addr_QCursor_OperatorAssign uintptr
	addr_QCursor_swap uintptr
	addr_QCursor_ToQVariant uintptr
	addr_QCursor_shape uintptr
	addr_QCursor_setShape uintptr
	addr_QCursor_bitmap uintptr
	addr_QCursor_mask uintptr
	addr_QCursor_bitmap2 uintptr
	addr_QCursor_mask2 uintptr
	addr_QCursor_pixmap uintptr
	addr_QCursor_hotSpot uintptr
	addr_QCursor_pos uintptr
	addr_QCursor_pos2 uintptr
	addr_QCursor_setPos uintptr
	addr_QCursor_setPos2 uintptr
	addr_QCursor_setPos3 uintptr
	addr_QCursor_setPos4 uintptr
	addr_QCursor_OperatorEqual uintptr
	addr_QCursor_OperatorNotEqual uintptr
	addr_QCursor_toQVariant uintptr
	pg_QCursor_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCursor_new, "QCursor_new")
	symaddr(&addr_QCursor_new2, "QCursor_new2")
	symaddr(&addr_QCursor_new3, "QCursor_new3")
	symaddr(&addr_QCursor_new4, "QCursor_new4")
	symaddr(&addr_QCursor_new5, "QCursor_new5")
	symaddr(&addr_QCursor_new6, "QCursor_new6")
	symaddr(&addr_QCursor_new7, "QCursor_new7")
	symaddr(&addr_QCursor_new8, "QCursor_new8")
	symaddr(&addr_QCursor_new9, "QCursor_new9")
	symaddr(&addr_QCursor_OperatorAssign, "QCursor_OperatorAssign")
	symaddr(&addr_QCursor_swap, "QCursor_swap")
	symaddr(&addr_QCursor_ToQVariant, "QCursor_ToQVariant")
	symaddr(&addr_QCursor_shape, "QCursor_shape")
	symaddr(&addr_QCursor_setShape, "QCursor_setShape")
	symaddr(&addr_QCursor_bitmap, "QCursor_bitmap")
	symaddr(&addr_QCursor_mask, "QCursor_mask")
	symaddr(&addr_QCursor_bitmap2, "QCursor_bitmap2")
	symaddr(&addr_QCursor_mask2, "QCursor_mask2")
	symaddr(&addr_QCursor_pixmap, "QCursor_pixmap")
	symaddr(&addr_QCursor_hotSpot, "QCursor_hotSpot")
	symaddr(&addr_QCursor_pos, "QCursor_pos")
	symaddr(&addr_QCursor_pos2, "QCursor_pos2")
	symaddr(&addr_QCursor_setPos, "QCursor_setPos")
	symaddr(&addr_QCursor_setPos2, "QCursor_setPos2")
	symaddr(&addr_QCursor_setPos3, "QCursor_setPos3")
	symaddr(&addr_QCursor_setPos4, "QCursor_setPos4")
	symaddr(&addr_QCursor_OperatorEqual, "QCursor_OperatorEqual")
	symaddr(&addr_QCursor_OperatorNotEqual, "QCursor_OperatorNotEqual")
	symaddr(&addr_QCursor_toQVariant, "QCursor_toQVariant")
	reg(&pg_QCursor_delete, "QCursor_delete")
}

type QCursor struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCursor) UnsafeGoID() int64     { return this.vid }
func (this *QCursor) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCursor(h uintptr) *QCursor {
	if h == 0 {
		return nil
	}
	return &QCursor{h: h}
}

func newownedQCursor(h uintptr) *QCursor {
	o := bindQCursor(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCursor).Delete)
	}
	return o
}

func (this *QCursor) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCursor) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCursor(h unsafe.Pointer) *QCursor { return bindQCursor(uintptr(h)) }
func UnsafeNewQCursorOwned(h unsafe.Pointer) *QCursor { return newownedQCursor(uintptr(h)) }

func (this *QCursor) UnsafeRelease() { this.h = 0 }

func NewQCursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new)
	__o := newownedQCursor(__r)
	return __o
}

func NewQCursor2(shape CursorShape) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new2, uintptr(shape))
	__o := newownedQCursor(__r)
	return __o
}

func NewQCursor3(bitmap *QBitmap, mask *QBitmap) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new3, uintptr(bitmap.UnsafePointer()), uintptr(mask.UnsafePointer()))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(bitmap)
	runtime.KeepAlive(mask)
	return __o
}

func NewQCursor4(pixmap *QPixmap) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new4, uintptr(pixmap.UnsafePointer()))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQCursor5(cursor *QCursor) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new5, uintptr(cursor.UnsafePointer()))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(cursor)
	return __o
}

func NewQCursor6(bitmap *QBitmap, mask *QBitmap, hotX int) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new6, uintptr(bitmap.UnsafePointer()), uintptr(mask.UnsafePointer()), uintptr(hotX))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(bitmap)
	runtime.KeepAlive(mask)
	return __o
}

func NewQCursor7(bitmap *QBitmap, mask *QBitmap, hotX int, hotY int) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new7, uintptr(bitmap.UnsafePointer()), uintptr(mask.UnsafePointer()), uintptr(hotX), uintptr(hotY))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(bitmap)
	runtime.KeepAlive(mask)
	return __o
}

func NewQCursor8(pixmap *QPixmap, hotX int) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new8, uintptr(pixmap.UnsafePointer()), uintptr(hotX))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func NewQCursor9(pixmap *QPixmap, hotX int, hotY int) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_new9, uintptr(pixmap.UnsafePointer()), uintptr(hotX), uintptr(hotY))
	__o := newownedQCursor(__r)
	runtime.KeepAlive(pixmap)
	return __o
}

func (this *QCursor) OperatorAssign(cursor *QCursor) *QCursor {
	__r, _, _ := purego.SyscallN(addr_QCursor_OperatorAssign, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
	return bindQCursor(__r)
}

func (this *QCursor) Swap(other *QCursor) {
	purego.SyscallN(addr_QCursor_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCursor) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCursor_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCursor) Shape() CursorShape {
	__r0, _, _ := purego.SyscallN(addr_QCursor_shape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CursorShape(__r)
}

func (this *QCursor) SetShape(newShape CursorShape) {
	purego.SyscallN(addr_QCursor_setShape, this.h, uintptr(newShape))
	runtime.KeepAlive(this)
}

func (this *QCursor) Bitmap(arg1 ReturnByValueConstant) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QCursor_bitmap, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QCursor) Mask(arg1 ReturnByValueConstant) *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QCursor_mask, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QCursor) Bitmap2() *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QCursor_bitmap2, this.h)
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QCursor) Mask2() *QBitmap {
	__r, _, _ := purego.SyscallN(addr_QCursor_mask2, this.h)
	runtime.KeepAlive(this)
	return newownedQBitmap(__r)
}

func (this *QCursor) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QCursor_pixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QCursor) HotSpot() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QCursor_hotSpot, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func QCursor_Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QCursor_pos)
	return newownedQPoint(__r)
}

func QCursor_PosWithScreen(screen *QScreen) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QCursor_pos2, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(screen)
	return newownedQPoint(__r)
}

func QCursor_SetPos(x int, y int) {
	purego.SyscallN(addr_QCursor_setPos, uintptr(x), uintptr(y))
}

func QCursor_SetPos2(screen *QScreen, x int, y int) {
	purego.SyscallN(addr_QCursor_setPos2, uintptr(screen.UnsafePointer()), uintptr(x), uintptr(y))
	runtime.KeepAlive(screen)
}

func QCursor_SetPosWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QCursor_setPos3, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(p)
}

func QCursor_SetPos3(screen *QScreen, p *QPoint) {
	purego.SyscallN(addr_QCursor_setPos4, uintptr(screen.UnsafePointer()), uintptr(p.UnsafePointer()))
	runtime.KeepAlive(screen)
	runtime.KeepAlive(p)
}

func (this *QCursor) OperatorEqual(rhs *QCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QCursor_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCursor) OperatorNotEqual(rhs *QCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QCursor_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCursor) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCursor_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCursor) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCursor_delete(this.h)
	}
	this.UnsafeRelease()
}
