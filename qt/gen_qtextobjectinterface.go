package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextObjectInterface_intrinsicSize uintptr
	addr_QTextObjectInterface_drawObject uintptr
	addr_QTextObjectInterface_OperatorAssign uintptr
	pg_QTextObjectInterface__set_override func(uintptr, int32)
	pg_QTextObjectInterface__Go_new func(int64) uintptr
	pg_QTextObjectInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextObjectInterface_intrinsicSize, "QTextObjectInterface_intrinsicSize")
	symaddr(&addr_QTextObjectInterface_drawObject, "QTextObjectInterface_drawObject")
	symaddr(&addr_QTextObjectInterface_OperatorAssign, "QTextObjectInterface_OperatorAssign")
	reg(&pg_QTextObjectInterface__set_override, "QTextObjectInterface__set_override")
	reg(&pg_QTextObjectInterface__Go_new, "QTextObjectInterface__Go_new")
	reg(&pg_QTextObjectInterface_delete, "QTextObjectInterface_delete")
}

type QTextObjectInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextObjectInterface) UnsafeGoID() int64     { return this.vid }
func (this *QTextObjectInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextObjectInterface(h uintptr) *QTextObjectInterface {
	if h == 0 {
		return nil
	}
	return &QTextObjectInterface{h: h}
}

func newownedQTextObjectInterface(h uintptr) *QTextObjectInterface {
	o := bindQTextObjectInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextObjectInterface).Delete)
	}
	return o
}

func (this *QTextObjectInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextObjectInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextObjectInterface(h unsafe.Pointer) *QTextObjectInterface { return bindQTextObjectInterface(uintptr(h)) }
func UnsafeNewQTextObjectInterfaceOwned(h unsafe.Pointer) *QTextObjectInterface { return newownedQTextObjectInterface(uintptr(h)) }

func (this *QTextObjectInterface) UnsafeRelease() { this.h = 0 }

func (this *QTextObjectInterface) IntrinsicSize(doc *QTextDocument, posInDocument int, format *QTextFormat) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QTextObjectInterface_intrinsicSize, this.h, uintptr(doc.UnsafePointer()), uintptr(posInDocument), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(doc)
	runtime.KeepAlive(format)
	return newownedQSizeF(__r)
}

func (this *QTextObjectInterface) DrawObject(painter *QPainter, rect *QRectF, doc *QTextDocument, posInDocument int, format *QTextFormat) {
	purego.SyscallN(addr_QTextObjectInterface_drawObject, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(doc.UnsafePointer()), uintptr(posInDocument), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(doc)
	runtime.KeepAlive(format)
}

func (this *QTextObjectInterface) OperatorAssign(arg1 *QTextObjectInterface) *QTextObjectInterface {
	__r, _, _ := purego.SyscallN(addr_QTextObjectInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextObjectInterface(__r)
}

func QTextObjectInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 5173984003908732465:
		return 0
	case 8290039120595127048:
		return 1
	}
	return -1
}

func QTextObjectInterface_enable(h uintptr, idx int32) { pg_QTextObjectInterface__set_override(h, idx) }

func NewQTextObjectInterface() *QTextObjectInterface {
	id := nextVirtID()
	o := newownedQTextObjectInterface(pg_QTextObjectInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QTextObjectInterface_virtIdx, QTextObjectInterface_enable)
	return o
}

func (this *QTextObjectInterface) OnIntrinsicSize(fn func(*QTextDocument, int, *QTextFormat) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 5173984003908732465, func(doc uintptr, posInDocument int32, format uintptr) uintptr { return uintptr(fn(UnsafeNewQTextDocument(unsafe.Pointer(doc)), int(posInDocument), UnsafeNewQTextFormat(unsafe.Pointer(format))).UnsafePointer()) }) // slot "intrinsicSize"
}

func (this *QTextObjectInterface) OnDrawObject(fn func(*QPainter, *QRectF, *QTextDocument, int, *QTextFormat)) {
	registerVirt(this.UnsafeGoID(), 8290039120595127048, func(painter uintptr, rect uintptr, doc uintptr, posInDocument int32, format uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRectF(unsafe.Pointer(rect)), UnsafeNewQTextDocument(unsafe.Pointer(doc)), int(posInDocument), UnsafeNewQTextFormat(unsafe.Pointer(format))) }) // slot "drawObject"
}

func (this *QTextObjectInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextObjectInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
