package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamNamespaceDeclaration_new uintptr
	addr_QXmlStreamNamespaceDeclaration_new2 uintptr
	addr_QXmlStreamNamespaceDeclaration_new3 uintptr
	addr_QXmlStreamNamespaceDeclaration_prefix uintptr
	addr_QXmlStreamNamespaceDeclaration_namespaceUri uintptr
	addr_QXmlStreamNamespaceDeclaration_OperatorEqual uintptr
	addr_QXmlStreamNamespaceDeclaration_OperatorNotEqual uintptr
	pg_QXmlStreamNamespaceDeclaration_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamNamespaceDeclaration_new, "QXmlStreamNamespaceDeclaration_new")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_new2, "QXmlStreamNamespaceDeclaration_new2")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_new3, "QXmlStreamNamespaceDeclaration_new3")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_prefix, "QXmlStreamNamespaceDeclaration_prefix")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_namespaceUri, "QXmlStreamNamespaceDeclaration_namespaceUri")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_OperatorEqual, "QXmlStreamNamespaceDeclaration_OperatorEqual")
	symaddr(&addr_QXmlStreamNamespaceDeclaration_OperatorNotEqual, "QXmlStreamNamespaceDeclaration_OperatorNotEqual")
	reg(&pg_QXmlStreamNamespaceDeclaration_delete, "QXmlStreamNamespaceDeclaration_delete")
}

type QXmlStreamNamespaceDeclaration struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamNamespaceDeclaration) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamNamespaceDeclaration) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamNamespaceDeclaration(h uintptr) *QXmlStreamNamespaceDeclaration {
	if h == 0 {
		return nil
	}
	return &QXmlStreamNamespaceDeclaration{h: h}
}

func newownedQXmlStreamNamespaceDeclaration(h uintptr) *QXmlStreamNamespaceDeclaration {
	o := bindQXmlStreamNamespaceDeclaration(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamNamespaceDeclaration).Delete)
	}
	return o
}

func (this *QXmlStreamNamespaceDeclaration) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamNamespaceDeclaration) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamNamespaceDeclaration(h unsafe.Pointer) *QXmlStreamNamespaceDeclaration { return bindQXmlStreamNamespaceDeclaration(uintptr(h)) }
func UnsafeNewQXmlStreamNamespaceDeclarationOwned(h unsafe.Pointer) *QXmlStreamNamespaceDeclaration { return newownedQXmlStreamNamespaceDeclaration(uintptr(h)) }

func (this *QXmlStreamNamespaceDeclaration) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_new)
	__o := newownedQXmlStreamNamespaceDeclaration(__r)
	return __o
}

func NewQXmlStreamNamespaceDeclaration2(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_new2, prefix_p, uintptr(len(prefix)), namespaceUri_p, uintptr(len(namespaceUri)))
	__o := newownedQXmlStreamNamespaceDeclaration(__r)
	runtime.KeepAlive(prefix)
	runtime.KeepAlive(namespaceUri)
	return __o
}

func NewQXmlStreamNamespaceDeclaration3(arg1 *QXmlStreamNamespaceDeclaration) *QXmlStreamNamespaceDeclaration {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQXmlStreamNamespaceDeclaration(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QXmlStreamNamespaceDeclaration) Prefix() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_prefix, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_namespaceUri, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamNamespaceDeclaration) OperatorEqual(rhs *QXmlStreamNamespaceDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamNamespaceDeclaration) OperatorNotEqual(rhs *QXmlStreamNamespaceDeclaration) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamNamespaceDeclaration_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamNamespaceDeclaration) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamNamespaceDeclaration_delete(this.h)
	}
	this.UnsafeRelease()
}
