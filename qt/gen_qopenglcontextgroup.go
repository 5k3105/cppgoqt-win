package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QOpenGLContextGroup_base func(uintptr) uintptr
	addr_QOpenGLContextGroup_tr uintptr
	addr_QOpenGLContextGroup_shares uintptr
	addr_QOpenGLContextGroup_currentContextGroup uintptr
	addr_QOpenGLContextGroup_connect uintptr
	addr_QOpenGLContextGroup_connect2 uintptr
	addr_QOpenGLContextGroup_tr2 uintptr
	addr_QOpenGLContextGroup_tr3 uintptr
	addr_QOpenGLContextGroup_connect3 uintptr
	addr_QOpenGLContextGroup_connect4 uintptr
	pg_QOpenGLContextGroup_connect_destroyed func(uintptr, int64)
	pg_QOpenGLContextGroup_connect_destroyed2 func(uintptr, int64)
	pg_QOpenGLContextGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QOpenGLContextGroup__disconnect func(int64)
	pg_QOpenGLContextGroup_delete func(uintptr)
)

func init() {
	reg(&pg_QOpenGLContextGroup_base, "QOpenGLContextGroup_base")
	symaddr(&addr_QOpenGLContextGroup_tr, "QOpenGLContextGroup_tr")
	symaddr(&addr_QOpenGLContextGroup_shares, "QOpenGLContextGroup_shares")
	symaddr(&addr_QOpenGLContextGroup_currentContextGroup, "QOpenGLContextGroup_currentContextGroup")
	symaddr(&addr_QOpenGLContextGroup_connect, "QOpenGLContextGroup_connect")
	symaddr(&addr_QOpenGLContextGroup_connect2, "QOpenGLContextGroup_connect2")
	symaddr(&addr_QOpenGLContextGroup_tr2, "QOpenGLContextGroup_tr2")
	symaddr(&addr_QOpenGLContextGroup_tr3, "QOpenGLContextGroup_tr3")
	symaddr(&addr_QOpenGLContextGroup_connect3, "QOpenGLContextGroup_connect3")
	symaddr(&addr_QOpenGLContextGroup_connect4, "QOpenGLContextGroup_connect4")
	reg(&pg_QOpenGLContextGroup_connect_destroyed, "QOpenGLContextGroup_connect_destroyed")
	reg(&pg_QOpenGLContextGroup_connect_destroyed2, "QOpenGLContextGroup_connect_destroyed2")
	reg(&pg_QOpenGLContextGroup_connect_objectNameChanged, "QOpenGLContextGroup_connect_objectNameChanged")
	reg(&pg_QOpenGLContextGroup__disconnect, "QOpenGLContextGroup__disconnect")
	reg(&pg_QOpenGLContextGroup_delete, "QOpenGLContextGroup_delete")
}

type QOpenGLContextGroup struct {
	h uintptr
	*QObject
}

func bindQOpenGLContextGroup(h uintptr) *QOpenGLContextGroup {
	if h == 0 {
		return nil
	}
	return &QOpenGLContextGroup{h: h, QObject: bindQObject(pg_QOpenGLContextGroup_base(h))}
}

func newownedQOpenGLContextGroup(h uintptr) *QOpenGLContextGroup { return bindQOpenGLContextGroup(h) }

func (this *QOpenGLContextGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOpenGLContextGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOpenGLContextGroup(h unsafe.Pointer) *QOpenGLContextGroup { return bindQOpenGLContextGroup(uintptr(h)) }
func UnsafeNewQOpenGLContextGroupOwned(h unsafe.Pointer) *QOpenGLContextGroup { return newownedQOpenGLContextGroup(uintptr(h)) }

func (this *QOpenGLContextGroup) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QOpenGLContextGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QOpenGLContextGroup) Shares() []*QOpenGLContext {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_shares, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QOpenGLContext, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQOpenGLContext(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QOpenGLContextGroup_CurrentContextGroup() *QOpenGLContextGroup {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_currentContextGroup)
	return bindQOpenGLContextGroup(__r)
}

func QOpenGLContextGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOpenGLContextGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QOpenGLContextGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOpenGLContextGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOpenGLContextGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOpenGLContextGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContextGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QOpenGLContextGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QOpenGLContextGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContextGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QOpenGLContextGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContextGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QOpenGLContextGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContextGroup) Disconnect(handle int64) {
	pg_QOpenGLContextGroup__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QOpenGLContextGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOpenGLContextGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
