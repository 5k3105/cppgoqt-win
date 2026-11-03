package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsAnchor_base func(uintptr) uintptr
	addr_QGraphicsAnchor_tr uintptr
	pg_QGraphicsAnchor_setSpacing func(uintptr, float64)
	addr_QGraphicsAnchor_unsetSpacing uintptr
	pg_QGraphicsAnchor_spacing func(uintptr) float64
	addr_QGraphicsAnchor_setSizePolicy uintptr
	addr_QGraphicsAnchor_sizePolicy uintptr
	addr_QGraphicsAnchor_connect uintptr
	addr_QGraphicsAnchor_connect2 uintptr
	addr_QGraphicsAnchor_tr2 uintptr
	addr_QGraphicsAnchor_tr3 uintptr
	addr_QGraphicsAnchor_connect3 uintptr
	addr_QGraphicsAnchor_connect4 uintptr
	pg_QGraphicsAnchor_connect_destroyed func(uintptr, int64)
	pg_QGraphicsAnchor_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsAnchor_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsAnchor__disconnect func(int64)
	pg_QGraphicsAnchor_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsAnchor_base, "QGraphicsAnchor_base")
	symaddr(&addr_QGraphicsAnchor_tr, "QGraphicsAnchor_tr")
	reg(&pg_QGraphicsAnchor_setSpacing, "QGraphicsAnchor_setSpacing")
	symaddr(&addr_QGraphicsAnchor_unsetSpacing, "QGraphicsAnchor_unsetSpacing")
	reg(&pg_QGraphicsAnchor_spacing, "QGraphicsAnchor_spacing")
	symaddr(&addr_QGraphicsAnchor_setSizePolicy, "QGraphicsAnchor_setSizePolicy")
	symaddr(&addr_QGraphicsAnchor_sizePolicy, "QGraphicsAnchor_sizePolicy")
	symaddr(&addr_QGraphicsAnchor_connect, "QGraphicsAnchor_connect")
	symaddr(&addr_QGraphicsAnchor_connect2, "QGraphicsAnchor_connect2")
	symaddr(&addr_QGraphicsAnchor_tr2, "QGraphicsAnchor_tr2")
	symaddr(&addr_QGraphicsAnchor_tr3, "QGraphicsAnchor_tr3")
	symaddr(&addr_QGraphicsAnchor_connect3, "QGraphicsAnchor_connect3")
	symaddr(&addr_QGraphicsAnchor_connect4, "QGraphicsAnchor_connect4")
	reg(&pg_QGraphicsAnchor_connect_destroyed, "QGraphicsAnchor_connect_destroyed")
	reg(&pg_QGraphicsAnchor_connect_destroyed2, "QGraphicsAnchor_connect_destroyed2")
	reg(&pg_QGraphicsAnchor_connect_objectNameChanged, "QGraphicsAnchor_connect_objectNameChanged")
	reg(&pg_QGraphicsAnchor__disconnect, "QGraphicsAnchor__disconnect")
	reg(&pg_QGraphicsAnchor_delete, "QGraphicsAnchor_delete")
}

type QGraphicsAnchor struct {
	h uintptr
	*QObject
}

func bindQGraphicsAnchor(h uintptr) *QGraphicsAnchor {
	if h == 0 {
		return nil
	}
	return &QGraphicsAnchor{h: h, QObject: bindQObject(pg_QGraphicsAnchor_base(h))}
}

func newownedQGraphicsAnchor(h uintptr) *QGraphicsAnchor { return bindQGraphicsAnchor(h) }

func (this *QGraphicsAnchor) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsAnchor) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsAnchor(h unsafe.Pointer) *QGraphicsAnchor { return bindQGraphicsAnchor(uintptr(h)) }
func UnsafeNewQGraphicsAnchorOwned(h unsafe.Pointer) *QGraphicsAnchor { return newownedQGraphicsAnchor(uintptr(h)) }

func (this *QGraphicsAnchor) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QGraphicsAnchor_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsAnchor) SetSpacing(spacing float64) {
	pg_QGraphicsAnchor_setSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchor) UnsetSpacing() {
	purego.SyscallN(addr_QGraphicsAnchor_unsetSpacing, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchor) Spacing() float64 {
	__r := pg_QGraphicsAnchor_spacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsAnchor) SetSizePolicy(policy QSizePolicy__Policy) {
	purego.SyscallN(addr_QGraphicsAnchor_setSizePolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchor) SizePolicy() QSizePolicy__Policy {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsAnchor_sizePolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__Policy(__r)
}

func QGraphicsAnchor_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsAnchor_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsAnchor_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsAnchor_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsAnchor_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsAnchor_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchor_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsAnchor) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsAnchor_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsAnchor) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsAnchor_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsAnchor) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsAnchor_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsAnchor) Disconnect(handle int64) {
	pg_QGraphicsAnchor__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QGraphicsAnchor) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsAnchor_delete(this.h)
	}
	this.UnsafeRelease()
}
