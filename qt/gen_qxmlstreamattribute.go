package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamAttribute_new uintptr
	addr_QXmlStreamAttribute_new2 uintptr
	addr_QXmlStreamAttribute_new3 uintptr
	addr_QXmlStreamAttribute_new4 uintptr
	addr_QXmlStreamAttribute_namespaceUri uintptr
	addr_QXmlStreamAttribute_name uintptr
	addr_QXmlStreamAttribute_qualifiedName uintptr
	addr_QXmlStreamAttribute_prefix uintptr
	addr_QXmlStreamAttribute_value uintptr
	addr_QXmlStreamAttribute_isDefault uintptr
	addr_QXmlStreamAttribute_OperatorEqual uintptr
	addr_QXmlStreamAttribute_OperatorNotEqual uintptr
	pg_QXmlStreamAttribute_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamAttribute_new, "QXmlStreamAttribute_new")
	symaddr(&addr_QXmlStreamAttribute_new2, "QXmlStreamAttribute_new2")
	symaddr(&addr_QXmlStreamAttribute_new3, "QXmlStreamAttribute_new3")
	symaddr(&addr_QXmlStreamAttribute_new4, "QXmlStreamAttribute_new4")
	symaddr(&addr_QXmlStreamAttribute_namespaceUri, "QXmlStreamAttribute_namespaceUri")
	symaddr(&addr_QXmlStreamAttribute_name, "QXmlStreamAttribute_name")
	symaddr(&addr_QXmlStreamAttribute_qualifiedName, "QXmlStreamAttribute_qualifiedName")
	symaddr(&addr_QXmlStreamAttribute_prefix, "QXmlStreamAttribute_prefix")
	symaddr(&addr_QXmlStreamAttribute_value, "QXmlStreamAttribute_value")
	symaddr(&addr_QXmlStreamAttribute_isDefault, "QXmlStreamAttribute_isDefault")
	symaddr(&addr_QXmlStreamAttribute_OperatorEqual, "QXmlStreamAttribute_OperatorEqual")
	symaddr(&addr_QXmlStreamAttribute_OperatorNotEqual, "QXmlStreamAttribute_OperatorNotEqual")
	reg(&pg_QXmlStreamAttribute_delete, "QXmlStreamAttribute_delete")
}

type QXmlStreamAttribute struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamAttribute) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamAttribute) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamAttribute(h uintptr) *QXmlStreamAttribute {
	if h == 0 {
		return nil
	}
	return &QXmlStreamAttribute{h: h}
}

func newownedQXmlStreamAttribute(h uintptr) *QXmlStreamAttribute {
	o := bindQXmlStreamAttribute(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamAttribute).Delete)
	}
	return o
}

func (this *QXmlStreamAttribute) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamAttribute) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamAttribute(h unsafe.Pointer) *QXmlStreamAttribute { return bindQXmlStreamAttribute(uintptr(h)) }
func UnsafeNewQXmlStreamAttributeOwned(h unsafe.Pointer) *QXmlStreamAttribute { return newownedQXmlStreamAttribute(uintptr(h)) }

func (this *QXmlStreamAttribute) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_new)
	__o := newownedQXmlStreamAttribute(__r)
	return __o
}

func NewQXmlStreamAttribute2(qualifiedName string, value string) *QXmlStreamAttribute {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_new2, qualifiedName_p, uintptr(len(qualifiedName)), value_p, uintptr(len(value)))
	__o := newownedQXmlStreamAttribute(__r)
	runtime.KeepAlive(qualifiedName)
	runtime.KeepAlive(value)
	return __o
}

func NewQXmlStreamAttribute3(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_new3, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	__o := newownedQXmlStreamAttribute(__r)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
	return __o
}

func NewQXmlStreamAttribute4(arg1 *QXmlStreamAttribute) *QXmlStreamAttribute {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQXmlStreamAttribute(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QXmlStreamAttribute) NamespaceUri() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_namespaceUri, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttribute) Name() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_name, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttribute) QualifiedName() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_qualifiedName, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttribute) Prefix() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_prefix, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttribute) Value() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_value, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamAttribute) IsDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_isDefault, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamAttribute) OperatorEqual(rhs *QXmlStreamAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamAttribute) OperatorNotEqual(rhs *QXmlStreamAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamAttribute_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QXmlStreamAttribute) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamAttribute_delete(this.h)
	}
	this.UnsafeRelease()
}
