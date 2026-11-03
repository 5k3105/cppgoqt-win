package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSocketNotifier_base func(uintptr) uintptr
	addr_QSocketNotifier_new uintptr
	addr_QSocketNotifier_new2 uintptr
	addr_QSocketNotifier_new3 uintptr
	addr_QSocketNotifier_new4 uintptr
	addr_QSocketNotifier_tr uintptr
	addr_QSocketNotifier_setSocket uintptr
	addr_QSocketNotifier_socket uintptr
	addr_QSocketNotifier_type uintptr
	addr_QSocketNotifier_isValid uintptr
	addr_QSocketNotifier_isEnabled uintptr
	addr_QSocketNotifier_setEnabled uintptr
	addr_QSocketNotifier_connect uintptr
	addr_QSocketNotifier_connect2 uintptr
	addr_QSocketNotifier_tr2 uintptr
	addr_QSocketNotifier_tr3 uintptr
	addr_QSocketNotifier_connect3 uintptr
	addr_QSocketNotifier_connect4 uintptr
	pg_QSocketNotifier_connect_activated func(uintptr, int64)
	pg_QSocketNotifier_connect_destroyed func(uintptr, int64)
	pg_QSocketNotifier_connect_destroyed2 func(uintptr, int64)
	pg_QSocketNotifier_connect_objectNameChanged func(uintptr, int64)
	pg_QSocketNotifier__disconnect func(int64)
	pg_QSocketNotifier_delete func(uintptr)
)

func init() {
	reg(&pg_QSocketNotifier_base, "QSocketNotifier_base")
	symaddr(&addr_QSocketNotifier_new, "QSocketNotifier_new")
	symaddr(&addr_QSocketNotifier_new2, "QSocketNotifier_new2")
	symaddr(&addr_QSocketNotifier_new3, "QSocketNotifier_new3")
	symaddr(&addr_QSocketNotifier_new4, "QSocketNotifier_new4")
	symaddr(&addr_QSocketNotifier_tr, "QSocketNotifier_tr")
	symaddr(&addr_QSocketNotifier_setSocket, "QSocketNotifier_setSocket")
	symaddr(&addr_QSocketNotifier_socket, "QSocketNotifier_socket")
	symaddr(&addr_QSocketNotifier_type, "QSocketNotifier_type")
	symaddr(&addr_QSocketNotifier_isValid, "QSocketNotifier_isValid")
	symaddr(&addr_QSocketNotifier_isEnabled, "QSocketNotifier_isEnabled")
	symaddr(&addr_QSocketNotifier_setEnabled, "QSocketNotifier_setEnabled")
	symaddr(&addr_QSocketNotifier_connect, "QSocketNotifier_connect")
	symaddr(&addr_QSocketNotifier_connect2, "QSocketNotifier_connect2")
	symaddr(&addr_QSocketNotifier_tr2, "QSocketNotifier_tr2")
	symaddr(&addr_QSocketNotifier_tr3, "QSocketNotifier_tr3")
	symaddr(&addr_QSocketNotifier_connect3, "QSocketNotifier_connect3")
	symaddr(&addr_QSocketNotifier_connect4, "QSocketNotifier_connect4")
	reg(&pg_QSocketNotifier_connect_activated, "QSocketNotifier_connect_activated")
	reg(&pg_QSocketNotifier_connect_destroyed, "QSocketNotifier_connect_destroyed")
	reg(&pg_QSocketNotifier_connect_destroyed2, "QSocketNotifier_connect_destroyed2")
	reg(&pg_QSocketNotifier_connect_objectNameChanged, "QSocketNotifier_connect_objectNameChanged")
	reg(&pg_QSocketNotifier__disconnect, "QSocketNotifier__disconnect")
	reg(&pg_QSocketNotifier_delete, "QSocketNotifier_delete")
}

type QSocketNotifier struct {
	h uintptr
	*QObject
}

func bindQSocketNotifier(h uintptr) *QSocketNotifier {
	if h == 0 {
		return nil
	}
	return &QSocketNotifier{h: h, QObject: bindQObject(pg_QSocketNotifier_base(h))}
}

func newownedQSocketNotifier(h uintptr) *QSocketNotifier { return bindQSocketNotifier(h) }

func (this *QSocketNotifier) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSocketNotifier) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSocketNotifier(h unsafe.Pointer) *QSocketNotifier { return bindQSocketNotifier(uintptr(h)) }
func UnsafeNewQSocketNotifierOwned(h unsafe.Pointer) *QSocketNotifier { return newownedQSocketNotifier(uintptr(h)) }

func (this *QSocketNotifier) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQSocketNotifier(arg1 QSocketNotifier__Type) *QSocketNotifier {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_new, uintptr(arg1))
	__o := newownedQSocketNotifier(__r)
	return __o
}

func NewQSocketNotifier2(socket uintptr, arg2 QSocketNotifier__Type) *QSocketNotifier {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_new2, uintptr(socket), uintptr(arg2))
	__o := newownedQSocketNotifier(__r)
	return __o
}

func NewQSocketNotifier3(arg1 QSocketNotifier__Type, parent *QObject) *QSocketNotifier {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_new3, uintptr(arg1), uintptr(parent.UnsafePointer()))
	__o := newownedQSocketNotifier(__r)
	runtime.KeepAlive(parent)
	return __o
}

func NewQSocketNotifier4(socket uintptr, arg2 QSocketNotifier__Type, parent *QObject) *QSocketNotifier {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_new4, uintptr(socket), uintptr(arg2), uintptr(parent.UnsafePointer()))
	__o := newownedQSocketNotifier(__r)
	runtime.KeepAlive(parent)
	return __o
}

func QSocketNotifier_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSocketNotifier) SetSocket(socket uintptr) {
	purego.SyscallN(addr_QSocketNotifier_setSocket, this.h, uintptr(socket))
	runtime.KeepAlive(this)
}

func (this *QSocketNotifier) Socket() uintptr {
	__r0, _, _ := purego.SyscallN(addr_QSocketNotifier_socket, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uintptr(__r)
}

func (this *QSocketNotifier) Type() QSocketNotifier__Type {
	__r0, _, _ := purego.SyscallN(addr_QSocketNotifier_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSocketNotifier__Type(__r)
}

func (this *QSocketNotifier) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QSocketNotifier_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSocketNotifier) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QSocketNotifier_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSocketNotifier) SetEnabled(arg1 bool) {
	purego.SyscallN(addr_QSocketNotifier_setEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func QSocketNotifier_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSocketNotifier_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QSocketNotifier_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSocketNotifier_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSocketNotifier_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSocketNotifier_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSocketNotifier_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSocketNotifier) OnActivated(fn func(*QSocketDescriptor, QSocketNotifier__Type)) int64 {
	id := registerCallback(func(socket uintptr, activationEvent int32) { fn(UnsafeNewQSocketDescriptor(unsafe.Pointer(socket)), QSocketNotifier__Type(activationEvent)) })
	pg_QSocketNotifier_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSocketNotifier) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSocketNotifier_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSocketNotifier) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSocketNotifier_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSocketNotifier) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSocketNotifier_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSocketNotifier) Disconnect(handle int64) {
	pg_QSocketNotifier__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QSocketNotifier) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSocketNotifier_delete(this.h)
	}
	this.UnsafeRelease()
}
