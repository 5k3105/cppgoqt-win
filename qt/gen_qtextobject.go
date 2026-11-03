package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextObject_base func(uintptr) uintptr
	addr_QTextObject_tr uintptr
	addr_QTextObject_format uintptr
	addr_QTextObject_formatIndex uintptr
	addr_QTextObject_document uintptr
	addr_QTextObject_objectIndex uintptr
	addr_QTextObject_connect uintptr
	addr_QTextObject_connect2 uintptr
	addr_QTextObject_tr2 uintptr
	addr_QTextObject_tr3 uintptr
	addr_QTextObject_connect3 uintptr
	addr_QTextObject_connect4 uintptr
	pg_QTextObject_connect_destroyed func(uintptr, int64)
	pg_QTextObject_connect_destroyed2 func(uintptr, int64)
	pg_QTextObject_connect_objectNameChanged func(uintptr, int64)
	pg_QTextObject__disconnect func(int64)
)

func init() {
	reg(&pg_QTextObject_base, "QTextObject_base")
	symaddr(&addr_QTextObject_tr, "QTextObject_tr")
	symaddr(&addr_QTextObject_format, "QTextObject_format")
	symaddr(&addr_QTextObject_formatIndex, "QTextObject_formatIndex")
	symaddr(&addr_QTextObject_document, "QTextObject_document")
	symaddr(&addr_QTextObject_objectIndex, "QTextObject_objectIndex")
	symaddr(&addr_QTextObject_connect, "QTextObject_connect")
	symaddr(&addr_QTextObject_connect2, "QTextObject_connect2")
	symaddr(&addr_QTextObject_tr2, "QTextObject_tr2")
	symaddr(&addr_QTextObject_tr3, "QTextObject_tr3")
	symaddr(&addr_QTextObject_connect3, "QTextObject_connect3")
	symaddr(&addr_QTextObject_connect4, "QTextObject_connect4")
	reg(&pg_QTextObject_connect_destroyed, "QTextObject_connect_destroyed")
	reg(&pg_QTextObject_connect_destroyed2, "QTextObject_connect_destroyed2")
	reg(&pg_QTextObject_connect_objectNameChanged, "QTextObject_connect_objectNameChanged")
	reg(&pg_QTextObject__disconnect, "QTextObject__disconnect")
}

type QTextObject struct {
	h uintptr
	*QObject
}

func bindQTextObject(h uintptr) *QTextObject {
	if h == 0 {
		return nil
	}
	return &QTextObject{h: h, QObject: bindQObject(pg_QTextObject_base(h))}
}

func newownedQTextObject(h uintptr) *QTextObject { return bindQTextObject(h) }

func (this *QTextObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextObject(h unsafe.Pointer) *QTextObject { return bindQTextObject(uintptr(h)) }
func UnsafeNewQTextObjectOwned(h unsafe.Pointer) *QTextObject { return newownedQTextObject(uintptr(h)) }

func (this *QTextObject) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QTextObject_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextObject_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextObject) Format() *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QTextObject_format, this.h)
	runtime.KeepAlive(this)
	return newownedQTextFormat(__r)
}

func (this *QTextObject) FormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextObject_formatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextObject) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextObject_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QTextObject) ObjectIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextObject_objectIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QTextObject_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextObject_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextObject_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextObject_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QTextObject_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextObject_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextObject_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextObject_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextObject_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextObject_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextObject_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextObject_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextObject) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextObject_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextObject) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextObject_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextObject) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextObject_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextObject) Disconnect(handle int64) {
	pg_QTextObject__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QTextObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
