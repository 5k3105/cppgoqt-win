package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPlainTextDocumentLayout_chain uintptr
	addr_QPlainTextDocumentLayout_new uintptr
	addr_QPlainTextDocumentLayout_tr uintptr
	addr_QPlainTextDocumentLayout_ensureBlockLayout uintptr
	addr_QPlainTextDocumentLayout_setCursorWidth uintptr
	addr_QPlainTextDocumentLayout_cursorWidth uintptr
	addr_QPlainTextDocumentLayout_requestUpdate uintptr
	addr_QPlainTextDocumentLayout_connect uintptr
	addr_QPlainTextDocumentLayout_connect2 uintptr
	addr_QPlainTextDocumentLayout_tr2 uintptr
	addr_QPlainTextDocumentLayout_tr3 uintptr
	addr_QPlainTextDocumentLayout_connect3 uintptr
	addr_QPlainTextDocumentLayout_connect4 uintptr
	pg_QPlainTextDocumentLayout_connect_update func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_update2 func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_updateBlock func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_documentSizeChanged func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_pageCountChanged func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_destroyed func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_destroyed2 func(uintptr, int64)
	pg_QPlainTextDocumentLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QPlainTextDocumentLayout__disconnect func(int64)
	pg_QPlainTextDocumentLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPlainTextDocumentLayout_chain, "QPlainTextDocumentLayout_chain")
	symaddr(&addr_QPlainTextDocumentLayout_new, "QPlainTextDocumentLayout_new")
	symaddr(&addr_QPlainTextDocumentLayout_tr, "QPlainTextDocumentLayout_tr")
	symaddr(&addr_QPlainTextDocumentLayout_ensureBlockLayout, "QPlainTextDocumentLayout_ensureBlockLayout")
	symaddr(&addr_QPlainTextDocumentLayout_setCursorWidth, "QPlainTextDocumentLayout_setCursorWidth")
	symaddr(&addr_QPlainTextDocumentLayout_cursorWidth, "QPlainTextDocumentLayout_cursorWidth")
	symaddr(&addr_QPlainTextDocumentLayout_requestUpdate, "QPlainTextDocumentLayout_requestUpdate")
	symaddr(&addr_QPlainTextDocumentLayout_connect, "QPlainTextDocumentLayout_connect")
	symaddr(&addr_QPlainTextDocumentLayout_connect2, "QPlainTextDocumentLayout_connect2")
	symaddr(&addr_QPlainTextDocumentLayout_tr2, "QPlainTextDocumentLayout_tr2")
	symaddr(&addr_QPlainTextDocumentLayout_tr3, "QPlainTextDocumentLayout_tr3")
	symaddr(&addr_QPlainTextDocumentLayout_connect3, "QPlainTextDocumentLayout_connect3")
	symaddr(&addr_QPlainTextDocumentLayout_connect4, "QPlainTextDocumentLayout_connect4")
	reg(&pg_QPlainTextDocumentLayout_connect_update, "QPlainTextDocumentLayout_connect_update")
	reg(&pg_QPlainTextDocumentLayout_connect_update2, "QPlainTextDocumentLayout_connect_update2")
	reg(&pg_QPlainTextDocumentLayout_connect_updateBlock, "QPlainTextDocumentLayout_connect_updateBlock")
	reg(&pg_QPlainTextDocumentLayout_connect_documentSizeChanged, "QPlainTextDocumentLayout_connect_documentSizeChanged")
	reg(&pg_QPlainTextDocumentLayout_connect_pageCountChanged, "QPlainTextDocumentLayout_connect_pageCountChanged")
	reg(&pg_QPlainTextDocumentLayout_connect_destroyed, "QPlainTextDocumentLayout_connect_destroyed")
	reg(&pg_QPlainTextDocumentLayout_connect_destroyed2, "QPlainTextDocumentLayout_connect_destroyed2")
	reg(&pg_QPlainTextDocumentLayout_connect_objectNameChanged, "QPlainTextDocumentLayout_connect_objectNameChanged")
	reg(&pg_QPlainTextDocumentLayout__disconnect, "QPlainTextDocumentLayout__disconnect")
	reg(&pg_QPlainTextDocumentLayout_delete, "QPlainTextDocumentLayout_delete")
}

type QPlainTextDocumentLayout struct {
	h uintptr
	*QAbstractTextDocumentLayout
}

func bindQPlainTextDocumentLayout(h uintptr) *QPlainTextDocumentLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPlainTextDocumentLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPlainTextDocumentLayout{h: h, QAbstractTextDocumentLayout: &QAbstractTextDocumentLayout{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPlainTextDocumentLayout(h uintptr) *QPlainTextDocumentLayout { return bindQPlainTextDocumentLayout(h) }

func (this *QPlainTextDocumentLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPlainTextDocumentLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPlainTextDocumentLayout(h unsafe.Pointer) *QPlainTextDocumentLayout { return bindQPlainTextDocumentLayout(uintptr(h)) }
func UnsafeNewQPlainTextDocumentLayoutOwned(h unsafe.Pointer) *QPlainTextDocumentLayout { return newownedQPlainTextDocumentLayout(uintptr(h)) }

func (this *QPlainTextDocumentLayout) UnsafeRelease() {
	this.h = 0
	if this.QAbstractTextDocumentLayout != nil {
		this.QAbstractTextDocumentLayout.UnsafeRelease()
	}
}

func NewQPlainTextDocumentLayout(document *QTextDocument) *QPlainTextDocumentLayout {
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_new, uintptr(document.UnsafePointer()))
	__o := newownedQPlainTextDocumentLayout(__r)
	runtime.KeepAlive(document)
	return __o
}

func QPlainTextDocumentLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPlainTextDocumentLayout) EnsureBlockLayout(block *QTextBlock) {
	purego.SyscallN(addr_QPlainTextDocumentLayout_ensureBlockLayout, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
}

func (this *QPlainTextDocumentLayout) SetCursorWidth(width int) {
	purego.SyscallN(addr_QPlainTextDocumentLayout_setCursorWidth, this.h, uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QPlainTextDocumentLayout) CursorWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_cursorWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPlainTextDocumentLayout) RequestUpdate() {
	purego.SyscallN(addr_QPlainTextDocumentLayout_requestUpdate, this.h)
	runtime.KeepAlive(this)
}

func QPlainTextDocumentLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPlainTextDocumentLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QPlainTextDocumentLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPlainTextDocumentLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPlainTextDocumentLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPlainTextDocumentLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPlainTextDocumentLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPlainTextDocumentLayout) OnUpdate(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextDocumentLayout_connect_update(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnUpdateWithQRectF(fn func(*QRectF)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(arg1))) })
	pg_QPlainTextDocumentLayout_connect_update2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnUpdateBlock(fn func(*QTextBlock)) int64 {
	id := registerCallback(func(block uintptr) { fn(UnsafeNewQTextBlock(unsafe.Pointer(block))) })
	pg_QPlainTextDocumentLayout_connect_updateBlock(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnDocumentSizeChanged(fn func(*QSizeF)) int64 {
	id := registerCallback(func(newSize uintptr) { fn(UnsafeNewQSizeF(unsafe.Pointer(newSize))) })
	pg_QPlainTextDocumentLayout_connect_documentSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnPageCountChanged(fn func(int)) int64 {
	id := registerCallback(func(newPages int32) { fn(int(newPages)) })
	pg_QPlainTextDocumentLayout_connect_pageCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPlainTextDocumentLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPlainTextDocumentLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPlainTextDocumentLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPlainTextDocumentLayout) Disconnect(handle int64) {
	pg_QPlainTextDocumentLayout__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QPlainTextDocumentLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPlainTextDocumentLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
