package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaObject_new uintptr
	addr_QMetaObject_new2 uintptr
	addr_QMetaObject_className uintptr
	addr_QMetaObject_superClass uintptr
	addr_QMetaObject_inherits uintptr
	addr_QMetaObject_cast uintptr
	addr_QMetaObject_cast2 uintptr
	addr_QMetaObject_tr uintptr
	addr_QMetaObject_metaType uintptr
	addr_QMetaObject_methodOffset uintptr
	addr_QMetaObject_enumeratorOffset uintptr
	addr_QMetaObject_propertyOffset uintptr
	addr_QMetaObject_classInfoOffset uintptr
	addr_QMetaObject_constructorCount uintptr
	addr_QMetaObject_methodCount uintptr
	addr_QMetaObject_enumeratorCount uintptr
	addr_QMetaObject_propertyCount uintptr
	addr_QMetaObject_classInfoCount uintptr
	addr_QMetaObject_indexOfConstructor uintptr
	addr_QMetaObject_indexOfMethod uintptr
	addr_QMetaObject_indexOfSignal uintptr
	addr_QMetaObject_indexOfSlot uintptr
	addr_QMetaObject_indexOfEnumerator uintptr
	addr_QMetaObject_indexOfProperty uintptr
	addr_QMetaObject_indexOfClassInfo uintptr
	addr_QMetaObject_constructor uintptr
	addr_QMetaObject_method uintptr
	addr_QMetaObject_enumerator uintptr
	addr_QMetaObject_property uintptr
	addr_QMetaObject_classInfo uintptr
	addr_QMetaObject_userProperty uintptr
	addr_QMetaObject_checkConnectArgs uintptr
	addr_QMetaObject_checkConnectArgs2 uintptr
	addr_QMetaObject_normalizedSignature uintptr
	addr_QMetaObject_normalizedType uintptr
	addr_QMetaObject_connect uintptr
	addr_QMetaObject_disconnectOne uintptr
	addr_QMetaObject_connectSlotsByName uintptr
	addr_QMetaObject_invokeMethod uintptr
	addr_QMetaObject_invokeMethod2 uintptr
	addr_QMetaObject_invokeMethod3 uintptr
	addr_QMetaObject_invokeMethod4 uintptr
	addr_QMetaObject_newInstance uintptr
	addr_QMetaObject_OperatorAssign uintptr
	addr_QMetaObject_tr2 uintptr
	addr_QMetaObject_connect2 uintptr
	addr_QMetaObject_connect3 uintptr
	addr_QMetaObject_invokeMethod5 uintptr
	addr_QMetaObject_invokeMethod6 uintptr
	addr_QMetaObject_invokeMethod7 uintptr
	addr_QMetaObject_invokeMethod8 uintptr
	addr_QMetaObject_invokeMethod9 uintptr
	addr_QMetaObject_invokeMethod10 uintptr
	addr_QMetaObject_invokeMethod11 uintptr
	addr_QMetaObject_invokeMethod12 uintptr
	addr_QMetaObject_invokeMethod13 uintptr
	addr_QMetaObject_invokeMethod14 uintptr
	addr_QMetaObject_invokeMethod15 uintptr
	addr_QMetaObject_invokeMethod16 uintptr
	addr_QMetaObject_invokeMethod17 uintptr
	addr_QMetaObject_invokeMethod18 uintptr
	addr_QMetaObject_invokeMethod19 uintptr
	addr_QMetaObject_invokeMethod20 uintptr
	addr_QMetaObject_invokeMethod21 uintptr
	addr_QMetaObject_invokeMethod22 uintptr
	addr_QMetaObject_invokeMethod23 uintptr
	addr_QMetaObject_invokeMethod24 uintptr
	addr_QMetaObject_invokeMethod25 uintptr
	addr_QMetaObject_invokeMethod26 uintptr
	addr_QMetaObject_invokeMethod27 uintptr
	addr_QMetaObject_invokeMethod28 uintptr
	addr_QMetaObject_invokeMethod29 uintptr
	addr_QMetaObject_invokeMethod30 uintptr
	addr_QMetaObject_invokeMethod31 uintptr
	addr_QMetaObject_invokeMethod32 uintptr
	addr_QMetaObject_invokeMethod33 uintptr
	addr_QMetaObject_invokeMethod34 uintptr
	addr_QMetaObject_invokeMethod35 uintptr
	addr_QMetaObject_invokeMethod36 uintptr
	addr_QMetaObject_invokeMethod37 uintptr
	addr_QMetaObject_invokeMethod38 uintptr
	addr_QMetaObject_invokeMethod39 uintptr
	addr_QMetaObject_invokeMethod40 uintptr
	addr_QMetaObject_invokeMethod41 uintptr
	addr_QMetaObject_invokeMethod42 uintptr
	addr_QMetaObject_newInstance2 uintptr
	addr_QMetaObject_newInstance3 uintptr
	addr_QMetaObject_newInstance4 uintptr
	addr_QMetaObject_newInstance5 uintptr
	addr_QMetaObject_newInstance6 uintptr
	addr_QMetaObject_newInstance7 uintptr
	addr_QMetaObject_newInstance8 uintptr
	addr_QMetaObject_newInstance9 uintptr
	addr_QMetaObject_newInstance10 uintptr
	addr_QMetaObject_d_get uintptr
	addr_QMetaObject_d_set uintptr
	pg_QMetaObject_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaObject_new, "QMetaObject_new")
	symaddr(&addr_QMetaObject_new2, "QMetaObject_new2")
	symaddr(&addr_QMetaObject_className, "QMetaObject_className")
	symaddr(&addr_QMetaObject_superClass, "QMetaObject_superClass")
	symaddr(&addr_QMetaObject_inherits, "QMetaObject_inherits")
	symaddr(&addr_QMetaObject_cast, "QMetaObject_cast")
	symaddr(&addr_QMetaObject_cast2, "QMetaObject_cast2")
	symaddr(&addr_QMetaObject_tr, "QMetaObject_tr")
	symaddr(&addr_QMetaObject_metaType, "QMetaObject_metaType")
	symaddr(&addr_QMetaObject_methodOffset, "QMetaObject_methodOffset")
	symaddr(&addr_QMetaObject_enumeratorOffset, "QMetaObject_enumeratorOffset")
	symaddr(&addr_QMetaObject_propertyOffset, "QMetaObject_propertyOffset")
	symaddr(&addr_QMetaObject_classInfoOffset, "QMetaObject_classInfoOffset")
	symaddr(&addr_QMetaObject_constructorCount, "QMetaObject_constructorCount")
	symaddr(&addr_QMetaObject_methodCount, "QMetaObject_methodCount")
	symaddr(&addr_QMetaObject_enumeratorCount, "QMetaObject_enumeratorCount")
	symaddr(&addr_QMetaObject_propertyCount, "QMetaObject_propertyCount")
	symaddr(&addr_QMetaObject_classInfoCount, "QMetaObject_classInfoCount")
	symaddr(&addr_QMetaObject_indexOfConstructor, "QMetaObject_indexOfConstructor")
	symaddr(&addr_QMetaObject_indexOfMethod, "QMetaObject_indexOfMethod")
	symaddr(&addr_QMetaObject_indexOfSignal, "QMetaObject_indexOfSignal")
	symaddr(&addr_QMetaObject_indexOfSlot, "QMetaObject_indexOfSlot")
	symaddr(&addr_QMetaObject_indexOfEnumerator, "QMetaObject_indexOfEnumerator")
	symaddr(&addr_QMetaObject_indexOfProperty, "QMetaObject_indexOfProperty")
	symaddr(&addr_QMetaObject_indexOfClassInfo, "QMetaObject_indexOfClassInfo")
	symaddr(&addr_QMetaObject_constructor, "QMetaObject_constructor")
	symaddr(&addr_QMetaObject_method, "QMetaObject_method")
	symaddr(&addr_QMetaObject_enumerator, "QMetaObject_enumerator")
	symaddr(&addr_QMetaObject_property, "QMetaObject_property")
	symaddr(&addr_QMetaObject_classInfo, "QMetaObject_classInfo")
	symaddr(&addr_QMetaObject_userProperty, "QMetaObject_userProperty")
	symaddr(&addr_QMetaObject_checkConnectArgs, "QMetaObject_checkConnectArgs")
	symaddr(&addr_QMetaObject_checkConnectArgs2, "QMetaObject_checkConnectArgs2")
	symaddr(&addr_QMetaObject_normalizedSignature, "QMetaObject_normalizedSignature")
	symaddr(&addr_QMetaObject_normalizedType, "QMetaObject_normalizedType")
	symaddr(&addr_QMetaObject_connect, "QMetaObject_connect")
	symaddr(&addr_QMetaObject_disconnectOne, "QMetaObject_disconnectOne")
	symaddr(&addr_QMetaObject_connectSlotsByName, "QMetaObject_connectSlotsByName")
	symaddr(&addr_QMetaObject_invokeMethod, "QMetaObject_invokeMethod")
	symaddr(&addr_QMetaObject_invokeMethod2, "QMetaObject_invokeMethod2")
	symaddr(&addr_QMetaObject_invokeMethod3, "QMetaObject_invokeMethod3")
	symaddr(&addr_QMetaObject_invokeMethod4, "QMetaObject_invokeMethod4")
	symaddr(&addr_QMetaObject_newInstance, "QMetaObject_newInstance")
	symaddr(&addr_QMetaObject_OperatorAssign, "QMetaObject_OperatorAssign")
	symaddr(&addr_QMetaObject_tr2, "QMetaObject_tr2")
	symaddr(&addr_QMetaObject_connect2, "QMetaObject_connect2")
	symaddr(&addr_QMetaObject_connect3, "QMetaObject_connect3")
	symaddr(&addr_QMetaObject_invokeMethod5, "QMetaObject_invokeMethod5")
	symaddr(&addr_QMetaObject_invokeMethod6, "QMetaObject_invokeMethod6")
	symaddr(&addr_QMetaObject_invokeMethod7, "QMetaObject_invokeMethod7")
	symaddr(&addr_QMetaObject_invokeMethod8, "QMetaObject_invokeMethod8")
	symaddr(&addr_QMetaObject_invokeMethod9, "QMetaObject_invokeMethod9")
	symaddr(&addr_QMetaObject_invokeMethod10, "QMetaObject_invokeMethod10")
	symaddr(&addr_QMetaObject_invokeMethod11, "QMetaObject_invokeMethod11")
	symaddr(&addr_QMetaObject_invokeMethod12, "QMetaObject_invokeMethod12")
	symaddr(&addr_QMetaObject_invokeMethod13, "QMetaObject_invokeMethod13")
	symaddr(&addr_QMetaObject_invokeMethod14, "QMetaObject_invokeMethod14")
	symaddr(&addr_QMetaObject_invokeMethod15, "QMetaObject_invokeMethod15")
	symaddr(&addr_QMetaObject_invokeMethod16, "QMetaObject_invokeMethod16")
	symaddr(&addr_QMetaObject_invokeMethod17, "QMetaObject_invokeMethod17")
	symaddr(&addr_QMetaObject_invokeMethod18, "QMetaObject_invokeMethod18")
	symaddr(&addr_QMetaObject_invokeMethod19, "QMetaObject_invokeMethod19")
	symaddr(&addr_QMetaObject_invokeMethod20, "QMetaObject_invokeMethod20")
	symaddr(&addr_QMetaObject_invokeMethod21, "QMetaObject_invokeMethod21")
	symaddr(&addr_QMetaObject_invokeMethod22, "QMetaObject_invokeMethod22")
	symaddr(&addr_QMetaObject_invokeMethod23, "QMetaObject_invokeMethod23")
	symaddr(&addr_QMetaObject_invokeMethod24, "QMetaObject_invokeMethod24")
	symaddr(&addr_QMetaObject_invokeMethod25, "QMetaObject_invokeMethod25")
	symaddr(&addr_QMetaObject_invokeMethod26, "QMetaObject_invokeMethod26")
	symaddr(&addr_QMetaObject_invokeMethod27, "QMetaObject_invokeMethod27")
	symaddr(&addr_QMetaObject_invokeMethod28, "QMetaObject_invokeMethod28")
	symaddr(&addr_QMetaObject_invokeMethod29, "QMetaObject_invokeMethod29")
	symaddr(&addr_QMetaObject_invokeMethod30, "QMetaObject_invokeMethod30")
	symaddr(&addr_QMetaObject_invokeMethod31, "QMetaObject_invokeMethod31")
	symaddr(&addr_QMetaObject_invokeMethod32, "QMetaObject_invokeMethod32")
	symaddr(&addr_QMetaObject_invokeMethod33, "QMetaObject_invokeMethod33")
	symaddr(&addr_QMetaObject_invokeMethod34, "QMetaObject_invokeMethod34")
	symaddr(&addr_QMetaObject_invokeMethod35, "QMetaObject_invokeMethod35")
	symaddr(&addr_QMetaObject_invokeMethod36, "QMetaObject_invokeMethod36")
	symaddr(&addr_QMetaObject_invokeMethod37, "QMetaObject_invokeMethod37")
	symaddr(&addr_QMetaObject_invokeMethod38, "QMetaObject_invokeMethod38")
	symaddr(&addr_QMetaObject_invokeMethod39, "QMetaObject_invokeMethod39")
	symaddr(&addr_QMetaObject_invokeMethod40, "QMetaObject_invokeMethod40")
	symaddr(&addr_QMetaObject_invokeMethod41, "QMetaObject_invokeMethod41")
	symaddr(&addr_QMetaObject_invokeMethod42, "QMetaObject_invokeMethod42")
	symaddr(&addr_QMetaObject_newInstance2, "QMetaObject_newInstance2")
	symaddr(&addr_QMetaObject_newInstance3, "QMetaObject_newInstance3")
	symaddr(&addr_QMetaObject_newInstance4, "QMetaObject_newInstance4")
	symaddr(&addr_QMetaObject_newInstance5, "QMetaObject_newInstance5")
	symaddr(&addr_QMetaObject_newInstance6, "QMetaObject_newInstance6")
	symaddr(&addr_QMetaObject_newInstance7, "QMetaObject_newInstance7")
	symaddr(&addr_QMetaObject_newInstance8, "QMetaObject_newInstance8")
	symaddr(&addr_QMetaObject_newInstance9, "QMetaObject_newInstance9")
	symaddr(&addr_QMetaObject_newInstance10, "QMetaObject_newInstance10")
	symaddr(&addr_QMetaObject_d_get, "QMetaObject_d_get")
	symaddr(&addr_QMetaObject_d_set, "QMetaObject_d_set")
	reg(&pg_QMetaObject_delete, "QMetaObject_delete")
}

type QMetaObject struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaObject) UnsafeGoID() int64     { return this.vid }
func (this *QMetaObject) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaObject(h uintptr) *QMetaObject {
	if h == 0 {
		return nil
	}
	return &QMetaObject{h: h}
}

func newownedQMetaObject(h uintptr) *QMetaObject {
	o := bindQMetaObject(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaObject).Delete)
	}
	return o
}

func (this *QMetaObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaObject(h unsafe.Pointer) *QMetaObject { return bindQMetaObject(uintptr(h)) }
func UnsafeNewQMetaObjectOwned(h unsafe.Pointer) *QMetaObject { return newownedQMetaObject(uintptr(h)) }

func (this *QMetaObject) UnsafeRelease() { this.h = 0 }

func NewQMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_new)
	__o := newownedQMetaObject(__r)
	return __o
}

func NewQMetaObject2(arg1 *QMetaObject) *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaObject(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaObject) ClassName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaObject_className, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaObject) SuperClass() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_superClass, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaObject) Inherits(metaObject *QMetaObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_inherits, this.h, uintptr(metaObject.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(metaObject)
	return __r != 0
}

func (this *QMetaObject) Cast(obj *QObject) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_cast, this.h, uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
	return bindQObject(__r)
}

func (this *QMetaObject) CastWithObj(obj *QObject) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_cast2, this.h, uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
	return bindQObject(__r)
}

func (this *QMetaObject) Tr(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaObject_tr, this.h, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QMetaObject) MetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_metaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaObject) MethodOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_methodOffset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) EnumeratorOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_enumeratorOffset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) PropertyOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_propertyOffset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) ClassInfoOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_classInfoOffset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) ConstructorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_constructorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) MethodCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_methodCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) EnumeratorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_enumeratorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) PropertyCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_propertyCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) ClassInfoCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_classInfoCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaObject) IndexOfConstructor(constructor string) int {
	constructor_b := append([]byte(constructor), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfConstructor, this.h, uintptr(unsafe.Pointer(&constructor_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constructor_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfMethod(method string) int {
	method_b := append([]byte(method), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfMethod, this.h, uintptr(unsafe.Pointer(&method_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(method_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfSignal(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfSignal, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfSlot(slot string) int {
	slot_b := append([]byte(slot), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfSlot, this.h, uintptr(unsafe.Pointer(&slot_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(slot_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfEnumerator(name string) int {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfEnumerator, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfProperty(name string) int {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfProperty, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return int(__r)
}

func (this *QMetaObject) IndexOfClassInfo(name string) int {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_indexOfClassInfo, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return int(__r)
}

func (this *QMetaObject) Constructor(index int) *QMetaMethod {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_constructor, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaMethod(__r)
}

func (this *QMetaObject) Method(index int) *QMetaMethod {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_method, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaMethod(__r)
}

func (this *QMetaObject) Enumerator(index int) *QMetaEnum {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_enumerator, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaEnum(__r)
}

func (this *QMetaObject) Property(index int) *QMetaProperty {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_property, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaProperty(__r)
}

func (this *QMetaObject) ClassInfo(index int) *QMetaClassInfo {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_classInfo, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQMetaClassInfo(__r)
}

func (this *QMetaObject) UserProperty() *QMetaProperty {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_userProperty, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaProperty(__r)
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	signal_b := append([]byte(signal), 0)
	method_b := append([]byte(method), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_checkConnectArgs, uintptr(unsafe.Pointer(&signal_b[0])), uintptr(unsafe.Pointer(&method_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(method_b)
	return __r != 0
}

func QMetaObject_CheckConnectArgs2(signal *QMetaMethod, method *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_checkConnectArgs2, uintptr(signal.UnsafePointer()), uintptr(method.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(signal)
	runtime.KeepAlive(method)
	return __r != 0
}

func QMetaObject_NormalizedSignature(method string) *QByteArray {
	method_b := append([]byte(method), 0)
	__r, _, _ := purego.SyscallN(addr_QMetaObject_normalizedSignature, uintptr(unsafe.Pointer(&method_b[0])))
	runtime.KeepAlive(method_b)
	return newownedQByteArray(__r)
}

func QMetaObject_NormalizedType(type_ string) *QByteArray {
	type__b := append([]byte(type_), 0)
	__r, _, _ := purego.SyscallN(addr_QMetaObject_normalizedType, uintptr(unsafe.Pointer(&type__b[0])))
	runtime.KeepAlive(type__b)
	return newownedQByteArray(__r)
}

func QMetaObject_Connect(sender *QObject, signal_index int, receiver *QObject, method_index int) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_connect, uintptr(sender.UnsafePointer()), uintptr(signal_index), uintptr(receiver.UnsafePointer()), uintptr(method_index))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(receiver)
	return newownedQMetaObject__Connection(__r)
}

func QMetaObject_DisconnectOne(sender *QObject, signal_index int, receiver *QObject, method_index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_disconnectOne, uintptr(sender.UnsafePointer()), uintptr(signal_index), uintptr(receiver.UnsafePointer()), uintptr(method_index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(receiver)
	return __r != 0
}

func QMetaObject_ConnectSlotsByName(o *QObject) {
	purego.SyscallN(addr_QMetaObject_connectSlotsByName, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(o)
}

func QMetaObject_InvokeMethod(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	return __r != 0
}

func QMetaObject_InvokeMethod2(obj *QObject, member string, ret *QGenericReturnArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod2, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	return __r != 0
}

func QMetaObject_InvokeMethod3(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod3, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	return __r != 0
}

func QMetaObject_InvokeMethod4(obj *QObject, member string, val0 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod4, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	return __r != 0
}

func (this *QMetaObject) NewInstance(val0 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance, this.h, uintptr(val0.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	return bindQObject(__r)
}

func (this *QMetaObject) OperatorAssign(arg1 *QMetaObject) *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQMetaObject(__r)
}

func (this *QMetaObject) Tr2(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaObject_tr2, this.h, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMetaObject_Connect2(sender *QObject, signal_index int, receiver *QObject, method_index int, type_ int) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_connect2, uintptr(sender.UnsafePointer()), uintptr(signal_index), uintptr(receiver.UnsafePointer()), uintptr(method_index), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(receiver)
	return newownedQMetaObject__Connection(__r)
}

func QMetaObject_Connect3(sender *QObject, signal_index int, receiver *QObject, method_index int, type_ int) (*QMetaObject__Connection, int) {
	var types_out int32
	__r, _, _ := purego.SyscallN(addr_QMetaObject_connect3, uintptr(sender.UnsafePointer()), uintptr(signal_index), uintptr(receiver.UnsafePointer()), uintptr(method_index), uintptr(type_), uintptr(unsafe.Pointer(&types_out)))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(receiver)
	return newownedQMetaObject__Connection(__r), int(types_out)
}

func QMetaObject_InvokeMethod5(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod5, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	return __r != 0
}

func QMetaObject_InvokeMethod6(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod6, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func QMetaObject_InvokeMethod7(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod7, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func QMetaObject_InvokeMethod8(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod8, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func QMetaObject_InvokeMethod9(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod9, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func QMetaObject_InvokeMethod10(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod10, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func QMetaObject_InvokeMethod11(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod11, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func QMetaObject_InvokeMethod12(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod12, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod13(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod13, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod14(obj *QObject, member string, arg3 ConnectionType, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod14, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg3), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod15(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod15, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	return __r != 0
}

func QMetaObject_InvokeMethod16(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod16, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func QMetaObject_InvokeMethod17(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod17, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func QMetaObject_InvokeMethod18(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod18, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func QMetaObject_InvokeMethod19(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod19, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func QMetaObject_InvokeMethod20(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod20, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func QMetaObject_InvokeMethod21(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod21, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func QMetaObject_InvokeMethod22(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod22, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod23(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod23, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod24(obj *QObject, member string, ret *QGenericReturnArgument, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod24, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(ret.UnsafePointer()), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ret)
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

func QMetaObject_InvokeMethod25(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod25, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func QMetaObject_InvokeMethod26(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod26, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func QMetaObject_InvokeMethod27(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod27, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func QMetaObject_InvokeMethod28(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod28, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func QMetaObject_InvokeMethod29(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod29, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func QMetaObject_InvokeMethod30(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod30, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func QMetaObject_InvokeMethod31(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod31, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func QMetaObject_InvokeMethod32(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod32, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func QMetaObject_InvokeMethod33(obj *QObject, member string, type_ ConnectionType, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod33, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func QMetaObject_InvokeMethod34(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod34, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return __r != 0
}

func QMetaObject_InvokeMethod35(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod35, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return __r != 0
}

func QMetaObject_InvokeMethod36(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod36, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return __r != 0
}

func QMetaObject_InvokeMethod37(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod37, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return __r != 0
}

func QMetaObject_InvokeMethod38(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod38, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return __r != 0
}

func QMetaObject_InvokeMethod39(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod39, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return __r != 0
}

func QMetaObject_InvokeMethod40(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod40, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func QMetaObject_InvokeMethod41(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod41, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func QMetaObject_InvokeMethod42(obj *QObject, member string, val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) bool {
	member_b := append([]byte(member), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaObject_invokeMethod42, uintptr(obj.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(obj)
	runtime.KeepAlive(member_b)
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

func (this *QMetaObject) NewInstance2(val0 *QGenericArgument, val1 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance2, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance3(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance3, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance4(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance4, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance5(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance5, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance6(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance6, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance7(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance7, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance8(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance8, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance9(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance9, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val0)
	runtime.KeepAlive(val1)
	runtime.KeepAlive(val2)
	runtime.KeepAlive(val3)
	runtime.KeepAlive(val4)
	runtime.KeepAlive(val5)
	runtime.KeepAlive(val6)
	runtime.KeepAlive(val7)
	runtime.KeepAlive(val8)
	return bindQObject(__r)
}

func (this *QMetaObject) NewInstance10(val0 *QGenericArgument, val1 *QGenericArgument, val2 *QGenericArgument, val3 *QGenericArgument, val4 *QGenericArgument, val5 *QGenericArgument, val6 *QGenericArgument, val7 *QGenericArgument, val8 *QGenericArgument, val9 *QGenericArgument) *QObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_newInstance10, this.h, uintptr(val0.UnsafePointer()), uintptr(val1.UnsafePointer()), uintptr(val2.UnsafePointer()), uintptr(val3.UnsafePointer()), uintptr(val4.UnsafePointer()), uintptr(val5.UnsafePointer()), uintptr(val6.UnsafePointer()), uintptr(val7.UnsafePointer()), uintptr(val8.UnsafePointer()), uintptr(val9.UnsafePointer()))
	runtime.KeepAlive(this)
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
	return bindQObject(__r)
}

func (this *QMetaObject) D() *QMetaObject__Data {
	__r, _, _ := purego.SyscallN(addr_QMetaObject_d_get, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaObject__Data(__r)
}

func (this *QMetaObject) SetD(d *QMetaObject__Data) {
	purego.SyscallN(addr_QMetaObject_d_set, this.h, uintptr(d.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(d)
}

func (this *QMetaObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaObject_delete(this.h)
	}
	this.UnsafeRelease()
}
