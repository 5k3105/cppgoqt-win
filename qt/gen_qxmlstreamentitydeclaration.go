package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamEntityDeclaration_new uintptr
	addr_QXmlStreamEntityDeclaration_new2 uintptr
	addr_QXmlStreamEntityDeclaration_name uintptr
	addr_QXmlStreamEntityDeclaration_notationName uintptr
	addr_QXmlStreamEntityDeclaration_systemId uintptr
	addr_QXmlStreamEntityDeclaration_publicId uintptr
	addr_QXmlStreamEntityDeclaration_value uintptr
	addr_QXmlStreamEntityDeclaration_OperatorEqual uintptr
	addr_QXmlStreamEntityDeclaration_OperatorNotEqual uintptr
	pg_QXmlStreamEntityDeclaration_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamEntityDeclaration_new, "QXmlStreamEntityDeclaration_new")
	symaddr(&addr_QXmlStreamEntityDeclaration_new2, "QXmlStreamEntityDeclaration_new2")
	symaddr(&addr_QXmlStreamEntityDeclaration_name, "QXmlStreamEntityDeclaration_name")
	symaddr(&addr_QXmlStreamEntityDeclaration_notationName, "QXmlStreamEntityDeclaration_notationName")
	symaddr(&addr_QXmlStreamEntityDeclaration_systemId, "QXmlStreamEntityDeclaration_systemId")
	symaddr(&addr_QXmlStreamEntityDeclaration_publicId, "QXmlStreamEntityDeclaration_publicId")
	symaddr(&addr_QXmlStreamEntityDeclaration_value, "QXmlStreamEntityDeclaration_value")
	symaddr(&addr_QXmlStreamEntityDeclaration_OperatorEqual, "QXmlStreamEntityDeclaration_OperatorEqual")
	symaddr(&addr_QXmlStreamEntityDeclaration_OperatorNotEqual, "QXmlStreamEntityDeclaration_OperatorNotEqual")
	reg(&pg_QXmlStreamEntityDeclaration_delete, "QXmlStreamEntityDeclaration_delete")
}

type QXmlStreamEntityDeclaration struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamEntityDeclaration) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamEntityDeclaration) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamEntityDeclaration(h uintptr) *QXmlStreamEntityDeclaration {
	if h == 0 {
		return nil
	}
	return &QXmlStreamEntityDeclaration{h: h}
}

func newownedQXmlStreamEntityDeclaration(h uintptr) *QXmlStreamEntityDeclaration {
	o := bindQXmlStreamEntityDeclaration(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamEntityDeclaration).Delete)
	}
	return o
}

func (this *QXmlStreamEntityDeclaration) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamEntityDeclaration) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamEntityDeclaration(h unsafe.Pointer) *QXmlStreamEntityDeclaration { return bindQXmlStreamEntityDeclaration(uintptr(h)) }
func UnsafeNewQXmlStreamEntityDeclarationOwned(h unsafe.Pointer) *QXmlStreamEntityDeclaration { return newownedQXmlStreamEntityDeclaration(uintptr(h)) }

func (this *QXmlStreamEntityDeclaration) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_new)
	__o := newownedQXmlStreamEntityDeclaration(__r)
	return __o
}

func NewQXmlStreamEntityDeclaration2(arg1 *QXmlStreamEntityDeclaration) *QXmlStreamEntityDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQXmlStreamEntityDeclaration(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QXmlStreamEntityDeclaration) Name() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_name, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamEntityDeclaration) NotationName() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_notationName, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamEntityDeclaration) SystemId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_systemId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamEntityDeclaration) PublicId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_publicId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamEntityDeclaration) Value() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_value, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamEntityDeclaration) OperatorEqual(rhs *QXmlStreamEntityDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamEntityDeclaration) OperatorNotEqual(rhs *QXmlStreamEntityDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamEntityDeclaration_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamEntityDeclaration) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamEntityDeclaration_delete(this.h)
	}
	this.UnsafeRelease()
}
