package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractEventDispatcherV2_chain uintptr
	addr_QAbstractEventDispatcherV2_tr uintptr
	addr_QAbstractEventDispatcherV2_processEventsWithDeadline uintptr
	addr_QAbstractEventDispatcherV2_instance uintptr
	addr_QAbstractEventDispatcherV2_connect uintptr
	addr_QAbstractEventDispatcherV2_connect2 uintptr
	addr_QAbstractEventDispatcherV2_tr2 uintptr
	addr_QAbstractEventDispatcherV2_tr3 uintptr
	addr_QAbstractEventDispatcherV2_instance2 uintptr
	addr_QAbstractEventDispatcherV2_connect3 uintptr
	addr_QAbstractEventDispatcherV2_connect4 uintptr
	pg_QAbstractEventDispatcherV2_connect_aboutToBlock func(uintptr, int64)
	pg_QAbstractEventDispatcherV2_connect_awake func(uintptr, int64)
	pg_QAbstractEventDispatcherV2_connect_destroyed func(uintptr, int64)
	pg_QAbstractEventDispatcherV2_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractEventDispatcherV2_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractEventDispatcherV2__disconnect func(int64)
	pg_QAbstractEventDispatcherV2_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractEventDispatcherV2_chain, "QAbstractEventDispatcherV2_chain")
	symaddr(&addr_QAbstractEventDispatcherV2_tr, "QAbstractEventDispatcherV2_tr")
	symaddr(&addr_QAbstractEventDispatcherV2_processEventsWithDeadline, "QAbstractEventDispatcherV2_processEventsWithDeadline")
	symaddr(&addr_QAbstractEventDispatcherV2_instance, "QAbstractEventDispatcherV2_instance")
	symaddr(&addr_QAbstractEventDispatcherV2_connect, "QAbstractEventDispatcherV2_connect")
	symaddr(&addr_QAbstractEventDispatcherV2_connect2, "QAbstractEventDispatcherV2_connect2")
	symaddr(&addr_QAbstractEventDispatcherV2_tr2, "QAbstractEventDispatcherV2_tr2")
	symaddr(&addr_QAbstractEventDispatcherV2_tr3, "QAbstractEventDispatcherV2_tr3")
	symaddr(&addr_QAbstractEventDispatcherV2_instance2, "QAbstractEventDispatcherV2_instance2")
	symaddr(&addr_QAbstractEventDispatcherV2_connect3, "QAbstractEventDispatcherV2_connect3")
	symaddr(&addr_QAbstractEventDispatcherV2_connect4, "QAbstractEventDispatcherV2_connect4")
	reg(&pg_QAbstractEventDispatcherV2_connect_aboutToBlock, "QAbstractEventDispatcherV2_connect_aboutToBlock")
	reg(&pg_QAbstractEventDispatcherV2_connect_awake, "QAbstractEventDispatcherV2_connect_awake")
	reg(&pg_QAbstractEventDispatcherV2_connect_destroyed, "QAbstractEventDispatcherV2_connect_destroyed")
	reg(&pg_QAbstractEventDispatcherV2_connect_destroyed2, "QAbstractEventDispatcherV2_connect_destroyed2")
	reg(&pg_QAbstractEventDispatcherV2_connect_objectNameChanged, "QAbstractEventDispatcherV2_connect_objectNameChanged")
	reg(&pg_QAbstractEventDispatcherV2__disconnect, "QAbstractEventDispatcherV2__disconnect")
	reg(&pg_QAbstractEventDispatcherV2_delete, "QAbstractEventDispatcherV2_delete")
}

type QAbstractEventDispatcherV2 struct {
	h uintptr
	*QAbstractEventDispatcher
}

func bindQAbstractEventDispatcherV2(h uintptr) *QAbstractEventDispatcherV2 {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractEventDispatcherV2_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractEventDispatcherV2{h: h, QAbstractEventDispatcher: &QAbstractEventDispatcher{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractEventDispatcherV2(h uintptr) *QAbstractEventDispatcherV2 { return bindQAbstractEventDispatcherV2(h) }

func (this *QAbstractEventDispatcherV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractEventDispatcherV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractEventDispatcherV2(h unsafe.Pointer) *QAbstractEventDispatcherV2 { return bindQAbstractEventDispatcherV2(uintptr(h)) }
func UnsafeNewQAbstractEventDispatcherV2Owned(h unsafe.Pointer) *QAbstractEventDispatcherV2 { return newownedQAbstractEventDispatcherV2(uintptr(h)) }

func (this *QAbstractEventDispatcherV2) UnsafeRelease() {
	this.h = 0
	if this.QAbstractEventDispatcher != nil {
		this.QAbstractEventDispatcher.UnsafeRelease()
	}
}

func QAbstractEventDispatcherV2_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractEventDispatcherV2) ProcessEventsWithDeadline(flags QEventLoop__ProcessEventsFlag, deadline *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_processEventsWithDeadline, this.h, uintptr(flags), uintptr(deadline.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(deadline)
	return __r != 0
}

func QAbstractEventDispatcherV2_Instance() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_instance)
	return bindQAbstractEventDispatcher(__r)
}

func QAbstractEventDispatcherV2_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractEventDispatcherV2_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractEventDispatcherV2_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractEventDispatcherV2_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractEventDispatcherV2_InstanceWithThread(thread *QThread) *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_instance2, uintptr(thread.UnsafePointer()))
	runtime.KeepAlive(thread)
	return bindQAbstractEventDispatcher(__r)
}

func QAbstractEventDispatcherV2_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractEventDispatcherV2_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcherV2_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractEventDispatcherV2) OnAboutToBlock(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcherV2_connect_aboutToBlock(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcherV2) OnAwake(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcherV2_connect_awake(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcherV2) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractEventDispatcherV2_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcherV2) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractEventDispatcherV2_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcherV2) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractEventDispatcherV2_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractEventDispatcherV2) Disconnect(handle int64) {
	pg_QAbstractEventDispatcherV2__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QAbstractEventDispatcherV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractEventDispatcherV2_delete(this.h)
	}
	this.UnsafeRelease()
}
