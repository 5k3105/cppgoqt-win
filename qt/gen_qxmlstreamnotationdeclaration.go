package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamNotationDeclaration_new uintptr
	addr_QXmlStreamNotationDeclaration_new2 uintptr
	addr_QXmlStreamNotationDeclaration_name uintptr
	addr_QXmlStreamNotationDeclaration_systemId uintptr
	addr_QXmlStreamNotationDeclaration_publicId uintptr
	addr_QXmlStreamNotationDeclaration_OperatorEqual uintptr
	addr_QXmlStreamNotationDeclaration_OperatorNotEqual uintptr
	pg_QXmlStreamNotationDeclaration_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamNotationDeclaration_new, "QXmlStreamNotationDeclaration_new")
	symaddr(&addr_QXmlStreamNotationDeclaration_new2, "QXmlStreamNotationDeclaration_new2")
	symaddr(&addr_QXmlStreamNotationDeclaration_name, "QXmlStreamNotationDeclaration_name")
	symaddr(&addr_QXmlStreamNotationDeclaration_systemId, "QXmlStreamNotationDeclaration_systemId")
	symaddr(&addr_QXmlStreamNotationDeclaration_publicId, "QXmlStreamNotationDeclaration_publicId")
	symaddr(&addr_QXmlStreamNotationDeclaration_OperatorEqual, "QXmlStreamNotationDeclaration_OperatorEqual")
	symaddr(&addr_QXmlStreamNotationDeclaration_OperatorNotEqual, "QXmlStreamNotationDeclaration_OperatorNotEqual")
	reg(&pg_QXmlStreamNotationDeclaration_delete, "QXmlStreamNotationDeclaration_delete")
}

type QXmlStreamNotationDeclaration struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamNotationDeclaration) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamNotationDeclaration) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamNotationDeclaration(h uintptr) *QXmlStreamNotationDeclaration {
	if h == 0 {
		return nil
	}
	return &QXmlStreamNotationDeclaration{h: h}
}

func newownedQXmlStreamNotationDeclaration(h uintptr) *QXmlStreamNotationDeclaration {
	o := bindQXmlStreamNotationDeclaration(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamNotationDeclaration).Delete)
	}
	return o
}

func (this *QXmlStreamNotationDeclaration) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamNotationDeclaration) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamNotationDeclaration(h unsafe.Pointer) *QXmlStreamNotationDeclaration { return bindQXmlStreamNotationDeclaration(uintptr(h)) }
func UnsafeNewQXmlStreamNotationDeclarationOwned(h unsafe.Pointer) *QXmlStreamNotationDeclaration { return newownedQXmlStreamNotationDeclaration(uintptr(h)) }

func (this *QXmlStreamNotationDeclaration) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamNotationDeclaration() *QXmlStreamNotationDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_new)
	__o := newownedQXmlStreamNotationDeclaration(__r)
	return __o
}

func NewQXmlStreamNotationDeclaration2(arg1 *QXmlStreamNotationDeclaration) *QXmlStreamNotationDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQXmlStreamNotationDeclaration(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QXmlStreamNotationDeclaration) Name() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_name, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamNotationDeclaration) SystemId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_systemId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamNotationDeclaration) PublicId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_publicId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamNotationDeclaration) OperatorEqual(rhs *QXmlStreamNotationDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamNotationDeclaration) OperatorNotEqual(rhs *QXmlStreamNotationDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamNotationDeclaration_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamNotationDeclaration) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamNotationDeclaration_delete(this.h)
	}
	this.UnsafeRelease()
}
