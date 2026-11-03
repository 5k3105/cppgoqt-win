package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextFrame_chain uintptr
	addr_QTextFrame_new uintptr
	addr_QTextFrame_tr uintptr
	addr_QTextFrame_setFrameFormat uintptr
	addr_QTextFrame_frameFormat uintptr
	addr_QTextFrame_firstCursorPosition uintptr
	addr_QTextFrame_lastCursorPosition uintptr
	addr_QTextFrame_firstPosition uintptr
	addr_QTextFrame_lastPosition uintptr
	addr_QTextFrame_layoutData uintptr
	addr_QTextFrame_setLayoutData uintptr
	addr_QTextFrame_childFrames uintptr
	addr_QTextFrame_parentFrame uintptr
	addr_QTextFrame_connect uintptr
	addr_QTextFrame_connect2 uintptr
	addr_QTextFrame_tr2 uintptr
	addr_QTextFrame_tr3 uintptr
	addr_QTextFrame_connect3 uintptr
	addr_QTextFrame_connect4 uintptr
	pg_QTextFrame_connect_destroyed func(uintptr, int64)
	pg_QTextFrame_connect_destroyed2 func(uintptr, int64)
	pg_QTextFrame_connect_objectNameChanged func(uintptr, int64)
	pg_QTextFrame__disconnect func(int64)
	pg_QTextFrame_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextFrame_chain, "QTextFrame_chain")
	symaddr(&addr_QTextFrame_new, "QTextFrame_new")
	symaddr(&addr_QTextFrame_tr, "QTextFrame_tr")
	symaddr(&addr_QTextFrame_setFrameFormat, "QTextFrame_setFrameFormat")
	symaddr(&addr_QTextFrame_frameFormat, "QTextFrame_frameFormat")
	symaddr(&addr_QTextFrame_firstCursorPosition, "QTextFrame_firstCursorPosition")
	symaddr(&addr_QTextFrame_lastCursorPosition, "QTextFrame_lastCursorPosition")
	symaddr(&addr_QTextFrame_firstPosition, "QTextFrame_firstPosition")
	symaddr(&addr_QTextFrame_lastPosition, "QTextFrame_lastPosition")
	symaddr(&addr_QTextFrame_layoutData, "QTextFrame_layoutData")
	symaddr(&addr_QTextFrame_setLayoutData, "QTextFrame_setLayoutData")
	symaddr(&addr_QTextFrame_childFrames, "QTextFrame_childFrames")
	symaddr(&addr_QTextFrame_parentFrame, "QTextFrame_parentFrame")
	symaddr(&addr_QTextFrame_connect, "QTextFrame_connect")
	symaddr(&addr_QTextFrame_connect2, "QTextFrame_connect2")
	symaddr(&addr_QTextFrame_tr2, "QTextFrame_tr2")
	symaddr(&addr_QTextFrame_tr3, "QTextFrame_tr3")
	symaddr(&addr_QTextFrame_connect3, "QTextFrame_connect3")
	symaddr(&addr_QTextFrame_connect4, "QTextFrame_connect4")
	reg(&pg_QTextFrame_connect_destroyed, "QTextFrame_connect_destroyed")
	reg(&pg_QTextFrame_connect_destroyed2, "QTextFrame_connect_destroyed2")
	reg(&pg_QTextFrame_connect_objectNameChanged, "QTextFrame_connect_objectNameChanged")
	reg(&pg_QTextFrame__disconnect, "QTextFrame__disconnect")
	reg(&pg_QTextFrame_delete, "QTextFrame_delete")
}

type QTextFrame struct {
	h uintptr
	*QTextObject
}

func bindQTextFrame(h uintptr) *QTextFrame {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTextFrame_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextFrame{h: h, QTextObject: &QTextObject{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQTextFrame(h uintptr) *QTextFrame { return bindQTextFrame(h) }

func (this *QTextFrame) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextFrame) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextFrame(h unsafe.Pointer) *QTextFrame { return bindQTextFrame(uintptr(h)) }
func UnsafeNewQTextFrameOwned(h unsafe.Pointer) *QTextFrame { return newownedQTextFrame(uintptr(h)) }

func (this *QTextFrame) UnsafeRelease() {
	this.h = 0
	if this.QTextObject != nil {
		this.QTextObject.UnsafeRelease()
	}
}

func NewQTextFrame(doc *QTextDocument) *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_new, uintptr(doc.UnsafePointer()))
	__o := newownedQTextFrame(__r)
	runtime.KeepAlive(doc)
	return __o
}

func QTextFrame_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFrame_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextFrame) SetFrameFormat(format *QTextFrameFormat) {
	purego.SyscallN(addr_QTextFrame_setFrameFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextFrame) FrameFormat() *QTextFrameFormat {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_frameFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextFrameFormat(__r)
}

func (this *QTextFrame) FirstCursorPosition() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_firstCursorPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextFrame) LastCursorPosition() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_lastCursorPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextFrame) FirstPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFrame_firstPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFrame) LastPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextFrame_lastPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextFrame) LayoutData() *QTextFrameLayoutData {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_layoutData, this.h)
	runtime.KeepAlive(this)
	return bindQTextFrameLayoutData(__r)
}

func (this *QTextFrame) SetLayoutData(data *QTextFrameLayoutData) {
	purego.SyscallN(addr_QTextFrame_setLayoutData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QTextFrame) ChildFrames() []*QTextFrame {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFrame_childFrames, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextFrame, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTextFrame(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextFrame) ParentFrame() *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_parentFrame, this.h)
	runtime.KeepAlive(this)
	return bindQTextFrame(__r)
}

func QTextFrame_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextFrame_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextFrame_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QTextFrame_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFrame_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextFrame_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextFrame_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextFrame_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextFrame_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextFrame_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextFrame_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextFrame) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextFrame_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextFrame) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextFrame_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextFrame) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextFrame_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextFrame) Disconnect(handle int64) {
	pg_QTextFrame__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QTextFrame) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextFrame_delete(this.h)
	}
	this.UnsafeRelease()
}
