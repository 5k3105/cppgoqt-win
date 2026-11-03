package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextBlockGroup_chain uintptr
	addr_QTextBlockGroup_tr uintptr
	addr_QTextBlockGroup_connect uintptr
	addr_QTextBlockGroup_connect2 uintptr
	addr_QTextBlockGroup_tr2 uintptr
	addr_QTextBlockGroup_tr3 uintptr
	addr_QTextBlockGroup_connect3 uintptr
	addr_QTextBlockGroup_connect4 uintptr
	pg_QTextBlockGroup_connect_destroyed func(uintptr, int64)
	pg_QTextBlockGroup_connect_destroyed2 func(uintptr, int64)
	pg_QTextBlockGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QTextBlockGroup__disconnect func(int64)
)

func init() {
	symaddr(&addr_QTextBlockGroup_chain, "QTextBlockGroup_chain")
	symaddr(&addr_QTextBlockGroup_tr, "QTextBlockGroup_tr")
	symaddr(&addr_QTextBlockGroup_connect, "QTextBlockGroup_connect")
	symaddr(&addr_QTextBlockGroup_connect2, "QTextBlockGroup_connect2")
	symaddr(&addr_QTextBlockGroup_tr2, "QTextBlockGroup_tr2")
	symaddr(&addr_QTextBlockGroup_tr3, "QTextBlockGroup_tr3")
	symaddr(&addr_QTextBlockGroup_connect3, "QTextBlockGroup_connect3")
	symaddr(&addr_QTextBlockGroup_connect4, "QTextBlockGroup_connect4")
	reg(&pg_QTextBlockGroup_connect_destroyed, "QTextBlockGroup_connect_destroyed")
	reg(&pg_QTextBlockGroup_connect_destroyed2, "QTextBlockGroup_connect_destroyed2")
	reg(&pg_QTextBlockGroup_connect_objectNameChanged, "QTextBlockGroup_connect_objectNameChanged")
	reg(&pg_QTextBlockGroup__disconnect, "QTextBlockGroup__disconnect")
}

type QTextBlockGroup struct {
	h uintptr
	*QTextObject
}

func bindQTextBlockGroup(h uintptr) *QTextBlockGroup {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTextBlockGroup_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextBlockGroup{h: h, QTextObject: &QTextObject{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTextBlockGroup(h uintptr) *QTextBlockGroup { return bindQTextBlockGroup(h) }

func (this *QTextBlockGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBlockGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBlockGroup(h unsafe.Pointer) *QTextBlockGroup { return bindQTextBlockGroup(uintptr(h)) }
func UnsafeNewQTextBlockGroupOwned(h unsafe.Pointer) *QTextBlockGroup { return newownedQTextBlockGroup(uintptr(h)) }

func (this *QTextBlockGroup) UnsafeRelease() {
	this.h = 0
	if this.QTextObject != nil {
		this.QTextObject.UnsafeRelease()
	}
}

func QTextBlockGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QTextBlockGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextBlockGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QTextBlockGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextBlockGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextBlockGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextBlockGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextBlockGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextBlockGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextBlockGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBlockGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextBlockGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBlockGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextBlockGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextBlockGroup) Disconnect(handle int64) {
	pg_QTextBlockGroup__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QTextBlockGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
