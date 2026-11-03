package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaMethod_new uintptr
	addr_QMetaMethod_new2 uintptr
	addr_QMetaMethod_methodSignature uintptr
	addr_QMetaMethod_name uintptr
	addr_QMetaMethod_nameView uintptr
	addr_QMetaMethod_typeName uintptr
	addr_QMetaMethod_returnType uintptr
	addr_QMetaMethod_returnMetaType uintptr
	addr_QMetaMethod_parameterCount uintptr
	addr_QMetaMethod_parameterType uintptr
	addr_QMetaMethod_parameterMetaType uintptr
	addr_QMetaMethod_getParameterTypes uintptr
	addr_QMetaMethod_parameterTypes uintptr
	addr_QMetaMethod_parameterTypeName uintptr
	addr_QMetaMethod_parameterNames uintptr
	addr_QMetaMethod_tag uintptr
	addr_QMetaMethod_access uintptr
	addr_QMetaMethod_methodType uintptr
	addr_QMetaMethod_attributes uintptr
	addr_QMetaMethod_methodIndex uintptr
	addr_QMetaMethod_relativeMethodIndex uintptr
	addr_QMetaMethod_revision uintptr
	addr_QMetaMethod_isConst uintptr
	addr_QMetaMethod_enclosingMetaObject uintptr
	addr_QMetaMethod_invoke uintptr
	addr_QMetaMethod_invoke2 uintptr
	addr_QMetaMethod_invoke3 uintptr
	addr_QMetaMethod_invoke4 uintptr
	addr_QMetaMethod_isValid uintptr
	addr_QMetaMethod_invoke5 uintptr
	addr_QMetaMethod_invoke6 uintptr
	addr_QMetaMethod_invoke7 uintptr
	addr_QMetaMethod_invoke8 uintptr
	addr_QMetaMethod_invoke9 uintptr
	addr_QMetaMethod_invoke10 uintptr
	addr_QMetaMethod_invoke11 uintptr
	addr_QMetaMethod_invoke12 uintptr
	addr_QMetaMethod_invoke13 uintptr
	addr_QMetaMethod_invoke14 uintptr
	addr_QMetaMethod_invoke15 uintptr
	addr_QMetaMethod_invoke16 uintptr
	addr_QMetaMethod_invoke17 uintptr
	addr_QMetaMethod_invoke18 uintptr
	addr_QMetaMethod_invoke19 uintptr
	addr_QMetaMethod_invoke20 uintptr
	addr_QMetaMethod_invoke21 uintptr
	addr_QMetaMethod_invoke22 uintptr
	addr_QMetaMethod_invoke23 uintptr
	addr_QMetaMethod_invoke24 uintptr
	addr_QMetaMethod_invoke25 uintptr
	addr_QMetaMethod_invoke26 uintptr
	addr_QMetaMethod_invoke27 uintptr
	addr_QMetaMethod_invoke28 uintptr
	addr_QMetaMethod_invoke29 uintptr
	addr_QMetaMethod_invoke30 uintptr
	addr_QMetaMethod_invoke31 uintptr
	addr_QMetaMethod_invoke32 uintptr
	addr_QMetaMethod_invoke33 uintptr
	addr_QMetaMethod_invoke34 uintptr
	addr_QMetaMethod_invoke35 uintptr
	addr_QMetaMethod_invoke36 uintptr
	addr_QMetaMethod_invoke37 uintptr
	addr_QMetaMethod_invoke38 uintptr
	addr_QMetaMethod_invoke39 uintptr
	addr_QMetaMethod_invoke40 uintptr
	addr_QMetaMethod_invoke41 uintptr
	addr_QMetaMethod_invoke42 uintptr
	addr_QMetaMethod_OperatorEqual uintptr
	addr_QMetaMethod_OperatorNotEqual uintptr
	pg_QMetaMethod_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaMethod_new, "QMetaMethod_new")
	symaddr(&addr_QMetaMethod_new2, "QMetaMethod_new2")
	symaddr(&addr_QMetaMethod_methodSignature, "QMetaMethod_methodSignature")
	symaddr(&addr_QMetaMethod_name, "QMetaMethod_name")
	symaddr(&addr_QMetaMethod_nameView, "QMetaMethod_nameView")
	symaddr(&addr_QMetaMethod_typeName, "QMetaMethod_typeName")
	symaddr(&addr_QMetaMethod_returnType, "QMetaMethod_returnType")
	symaddr(&addr_QMetaMethod_returnMetaType, "QMetaMethod_returnMetaType")
	symaddr(&addr_QMetaMethod_parameterCount, "QMetaMethod_parameterCount")
	symaddr(&addr_QMetaMethod_parameterType, "QMetaMethod_parameterType")
	symaddr(&addr_QMetaMethod_parameterMetaType, "QMetaMethod_parameterMetaType")
	symaddr(&addr_QMetaMethod_getParameterTypes, "QMetaMethod_getParameterTypes")
	symaddr(&addr_QMetaMethod_parameterTypes, "QMetaMethod_parameterTypes")
	symaddr(&addr_QMetaMethod_parameterTypeName, "QMetaMethod_parameterTypeName")
	symaddr(&addr_QMetaMethod_parameterNames, "QMetaMethod_parameterNames")
	symaddr(&addr_QMetaMethod_tag, "QMetaMethod_tag")
	symaddr(&addr_QMetaMethod_access, "QMetaMethod_access")
	symaddr(&addr_QMetaMethod_methodType, "QMetaMethod_methodType")
	symaddr(&addr_QMetaMethod_attributes, "QMetaMethod_attributes")
	symaddr(&addr_QMetaMethod_methodIndex, "QMetaMethod_methodIndex")
	symaddr(&addr_QMetaMethod_relativeMethodIndex, "QMetaMethod_relativeMethodIndex")
	symaddr(&addr_QMetaMethod_revision, "QMetaMethod_revision")
	symaddr(&addr_QMetaMethod_isConst, "QMetaMethod_isConst")
	symaddr(&addr_QMetaMethod_enclosingMetaObject, "QMetaMethod_enclosingMetaObject")
	symaddr(&addr_QMetaMethod_invoke, "QMetaMethod_invoke")
	symaddr(&addr_QMetaMethod_invoke2, "QMetaMethod_invoke2")
	symaddr(&addr_QMetaMethod_invoke3, "QMetaMethod_invoke3")
	symaddr(&addr_QMetaMethod_invoke4, "QMetaMethod_invoke4")
	symaddr(&addr_QMetaMethod_isValid, "QMetaMethod_isValid")
	symaddr(&addr_QMetaMethod_invoke5, "QMetaMethod_invoke5")
	symaddr(&addr_QMetaMethod_invoke6, "QMetaMethod_invoke6")
	symaddr(&addr_QMetaMethod_invoke7, "QMetaMethod_invoke7")
	symaddr(&addr_QMetaMethod_invoke8, "QMetaMethod_invoke8")
	symaddr(&addr_QMetaMethod_invoke9, "QMetaMethod_invoke9")
	symaddr(&addr_QMetaMethod_invoke10, "QMetaMethod_invoke10")
	symaddr(&addr_QMetaMethod_invoke11, "QMetaMethod_invoke11")
	symaddr(&addr_QMetaMethod_invoke12, "QMetaMethod_invoke12")
	symaddr(&addr_QMetaMethod_invoke13, "QMetaMethod_invoke13")
	symaddr(&addr_QMetaMethod_invoke14, "QMetaMethod_invoke14")
	symaddr(&addr_QMetaMethod_invoke15, "QMetaMethod_invoke15")
	symaddr(&addr_QMetaMethod_invoke16, "QMetaMethod_invoke16")
	symaddr(&addr_QMetaMethod_invoke17, "QMetaMethod_invoke17")
	symaddr(&addr_QMetaMethod_invoke18, "QMetaMethod_invoke18")
	symaddr(&addr_QMetaMethod_invoke19, "QMetaMethod_invoke19")
	symaddr(&addr_QMetaMethod_invoke20, "QMetaMethod_invoke20")
	symaddr(&addr_QMetaMethod_invoke21, "QMetaMethod_invoke21")
	symaddr(&addr_QMetaMethod_invoke22, "QMetaMethod_invoke22")
	symaddr(&addr_QMetaMethod_invoke23, "QMetaMethod_invoke23")
	symaddr(&addr_QMetaMethod_invoke24, "QMetaMethod_invoke24")
	symaddr(&addr_QMetaMethod_invoke25, "QMetaMethod_invoke25")
	symaddr(&addr_QMetaMethod_invoke26, "QMetaMethod_invoke26")
	symaddr(&addr_QMetaMethod_invoke27, "QMetaMethod_invoke27")
	symaddr(&addr_QMetaMethod_invoke28, "QMetaMethod_invoke28")
	symaddr(&addr_QMetaMethod_invoke29, "QMetaMethod_invoke29")
	symaddr(&addr_QMetaMethod_invoke30, "QMetaMethod_invoke30")
	symaddr(&addr_QMetaMethod_invoke31, "QMetaMethod_invoke31")
	symaddr(&addr_QMetaMethod_invoke32, "QMetaMethod_invoke32")
	symaddr(&addr_QMetaMethod_invoke33, "QMetaMethod_invoke33")
	symaddr(&addr_QMetaMethod_invoke34, "QMetaMethod_invoke34")
	symaddr(&addr_QMetaMethod_invoke35, "QMetaMethod_invoke35")
	symaddr(&addr_QMetaMethod_invoke36, "QMetaMethod_invoke36")
	symaddr(&addr_QMetaMethod_invoke37, "QMetaMethod_invoke37")
	symaddr(&addr_QMetaMethod_invoke38, "QMetaMethod_invoke38")
	symaddr(&addr_QMetaMethod_invoke39, "QMetaMethod_invoke39")
	symaddr(&addr_QMetaMethod_invoke40, "QMetaMethod_invoke40")
	symaddr(&addr_QMetaMethod_invoke41, "QMetaMethod_invoke41")
	symaddr(&addr_QMetaMethod_invoke42, "QMetaMethod_invoke42")
	symaddr(&addr_QMetaMethod_OperatorEqual, "QMetaMethod_OperatorEqual")
	symaddr(&addr_QMetaMethod_OperatorNotEqual, "QMetaMethod_OperatorNotEqual")
	reg(&pg_QMetaMethod_delete, "QMetaMethod_delete")
}

type QMetaMethod struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaMethod) UnsafeGoID() int64     { return this.vid }
func (this *QMetaMethod) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaMethod(h uintptr) *QMetaMethod {
	if h == 0 {
		return nil
	}
	return &QMetaMethod{h: h}
}

func newownedQMetaMethod(h uintptr) *QMetaMethod {
	o := bindQMetaMethod(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaMethod).Delete)
	}
	return o
}

func (this *QMetaMethod) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaMethod) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaMethod(h unsafe.Pointer) *QMetaMethod { return bindQMetaMethod(uintptr(h)) }
func UnsafeNewQMetaMethodOwned(h unsafe.Pointer) *QMetaMethod { return newownedQMetaMethod(uintptr(h)) }

func (this *QMetaMethod) UnsafeRelease() { this.h = 0 }

func NewQMetaMethod() *QMetaMethod {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_new)
	__o := newownedQMetaMethod(__r)
	return __o
}

func NewQMetaMethod2(arg1 *QMetaMethod) *QMetaMethod {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaMethod(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaMethod) MethodSignature() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_methodSignature, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMetaMethod) Name() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_name, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMetaMethod) NameView() *QByteArrayView {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_nameView, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArrayView(__r)
}

func (this *QMetaMethod) TypeName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_typeName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaMethod) ReturnType() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_returnType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) ReturnMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_returnMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaMethod) ParameterCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_parameterCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) ParameterType(index int) int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_parameterType, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) ParameterMetaType(index int) *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_parameterMetaType, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaMethod) GetParameterTypes() int {
	var types_out int32
	purego.SyscallN(addr_QMetaMethod_getParameterTypes, this.h, uintptr(unsafe.Pointer(&types_out)))
	runtime.KeepAlive(this)
	return int(types_out)
}

func (this *QMetaMethod) ParameterTypes() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_parameterTypes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMetaMethod) ParameterTypeName(index int) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_parameterTypeName, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMetaMethod) ParameterNames() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_parameterNames, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMetaMethod) Tag() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_tag, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaMethod) Access() QMetaMethod__Access {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_access, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMetaMethod__Access(__r)
}

func (this *QMetaMethod) MethodType() QMetaMethod__MethodType {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_methodType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMetaMethod__MethodType(__r)
}

func (this *QMetaMethod) Attributes() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_attributes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) MethodIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_methodIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) RelativeMethodIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_relativeMethodIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) Revision() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_revision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaMethod) IsConst() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_isConst, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaMethod) EnclosingMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaMethod_enclosingMetaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaMethod) Invoke(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	return __r != 0
}

func (this *QMetaMethod) Invoke2(object *QObject, returnValue *QGenericReturnArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke2, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	return __r != 0
}

func (this *QMetaMethod) Invoke3(object *QObject, connectionType ConnectionType, val0 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke3, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	return __r != 0
}

func (this *QMetaMethod) Invoke4(object *QObject, val0 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke4, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	return __r != 0
}

func (this *QMetaMethod) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaMethod) Invoke5(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke5, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	return __r != 0
}

func (this *QMetaMethod) Invoke6(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke6, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func (this *QMetaMethod) Invoke7(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke7, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func (this *QMetaMethod) Invoke8(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke8, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func (this *QMetaMethod) Invoke9(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke9, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func (this *QMetaMethod) Invoke10(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke10, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func (this *QMetaMethod) Invoke11(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke11, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func (this *QMetaMethod) Invoke12(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke12, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	return __r != 0
}

func (this *QMetaMethod) Invoke13(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke13, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	return __r != 0
}

func (this *QMetaMethod) Invoke14(object *QObject, connectionType ConnectionType, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke14, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	runtime.KeepAlive(val9)
	return __r != 0
}

func (this *QMetaMethod) Invoke15(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke15, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	return __r != 0
}

func (this *QMetaMethod) Invoke16(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke16, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func (this *QMetaMethod) Invoke17(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke17, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func (this *QMetaMethod) Invoke18(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke18, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func (this *QMetaMethod) Invoke19(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke19, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func (this *QMetaMethod) Invoke20(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke20, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func (this *QMetaMethod) Invoke21(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke21, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func (this *QMetaMethod) Invoke22(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke22, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	return __r != 0
}

func (this *QMetaMethod) Invoke23(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke23, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	return __r != 0
}

func (this *QMetaMethod) Invoke24(object *QObject, returnValue *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke24, this.h, uintptr(object.UnsafePointer()), uintptr(returnValue.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(returnValue)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	runtime.KeepAlive(val9)
	return __r != 0
}

func (this *QMetaMethod) Invoke25(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke25, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func (this *QMetaMethod) Invoke26(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke26, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func (this *QMetaMethod) Invoke27(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke27, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func (this *QMetaMethod) Invoke28(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke28, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func (this *QMetaMethod) Invoke29(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke29, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func (this *QMetaMethod) Invoke30(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke30, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func (this *QMetaMethod) Invoke31(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke31, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	return __r != 0
}

func (this *QMetaMethod) Invoke32(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke32, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	return __r != 0
}

func (this *QMetaMethod) Invoke33(object *QObject, connectionType ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke33, this.h, uintptr(object.UnsafePointer()), uintptr(connectionType), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	runtime.KeepAlive(val9)
	return __r != 0
}

func (this *QMetaMethod) Invoke34(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke34, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func (this *QMetaMethod) Invoke35(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke35, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func (this *QMetaMethod) Invoke36(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke36, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func (this *QMetaMethod) Invoke37(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke37, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func (this *QMetaMethod) Invoke38(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke38, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func (this *QMetaMethod) Invoke39(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke39, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func (this *QMetaMethod) Invoke40(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke40, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	return __r != 0
}

func (this *QMetaMethod) Invoke41(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke41, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	return __r != 0
}

func (this *QMetaMethod) Invoke42(object *QObject, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_invoke42, this.h, uintptr(object.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	runtime.KeepAlive(val9)
	return __r != 0
}

func (this *QMetaMethod) OperatorEqual(rhs *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaMethod) OperatorNotEqual(rhs *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaMethod_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMetaMethod) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaMethod_delete(this.h)
	}
	this.UnsafeRelease()
}
