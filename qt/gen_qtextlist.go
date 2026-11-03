package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextList_chain uintptr
	addr_QTextList_new uintptr
	addr_QTextList_tr uintptr
	addr_QTextList_count uintptr
	addr_QTextList_item uintptr
	addr_QTextList_itemNumber uintptr
	addr_QTextList_itemText uintptr
	addr_QTextList_removeItem uintptr
	addr_QTextList_remove uintptr
	addr_QTextList_add uintptr
	addr_QTextList_setFormat uintptr
	addr_QTextList_format uintptr
	addr_QTextList_connect uintptr
	addr_QTextList_connect2 uintptr
	addr_QTextList_tr2 uintptr
	addr_QTextList_tr3 uintptr
	addr_QTextList_connect3 uintptr
	addr_QTextList_connect4 uintptr
	pg_QTextList_connect_destroyed func(uintptr, int64)
	pg_QTextList_connect_destroyed2 func(uintptr, int64)
	pg_QTextList_connect_objectNameChanged func(uintptr, int64)
	pg_QTextList__disconnect func(int64)
	pg_QTextList_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextList_chain, "QTextList_chain")
	symaddr(&addr_QTextList_new, "QTextList_new")
	symaddr(&addr_QTextList_tr, "QTextList_tr")
	symaddr(&addr_QTextList_count, "QTextList_count")
	symaddr(&addr_QTextList_item, "QTextList_item")
	symaddr(&addr_QTextList_itemNumber, "QTextList_itemNumber")
	symaddr(&addr_QTextList_itemText, "QTextList_itemText")
	symaddr(&addr_QTextList_removeItem, "QTextList_removeItem")
	symaddr(&addr_QTextList_remove, "QTextList_remove")
	symaddr(&addr_QTextList_add, "QTextList_add")
	symaddr(&addr_QTextList_setFormat, "QTextList_setFormat")
	symaddr(&addr_QTextList_format, "QTextList_format")
	symaddr(&addr_QTextList_connect, "QTextList_connect")
	symaddr(&addr_QTextList_connect2, "QTextList_connect2")
	symaddr(&addr_QTextList_tr2, "QTextList_tr2")
	symaddr(&addr_QTextList_tr3, "QTextList_tr3")
	symaddr(&addr_QTextList_connect3, "QTextList_connect3")
	symaddr(&addr_QTextList_connect4, "QTextList_connect4")
	reg(&pg_QTextList_connect_destroyed, "QTextList_connect_destroyed")
	reg(&pg_QTextList_connect_destroyed2, "QTextList_connect_destroyed2")
	reg(&pg_QTextList_connect_objectNameChanged, "QTextList_connect_objectNameChanged")
	reg(&pg_QTextList__disconnect, "QTextList__disconnect")
	reg(&pg_QTextList_delete, "QTextList_delete")
}

type QTextList struct {
	h uintptr
	*QTextBlockGroup
}

func bindQTextList(h uintptr) *QTextList {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QTextList_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextList{h: h, QTextBlockGroup: &QTextBlockGroup{h: __c[0], QTextObject: &QTextObject{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQTextList(h uintptr) *QTextList { return bindQTextList(h) }

func (this *QTextList) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextList) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextList(h unsafe.Pointer) *QTextList { return bindQTextList(uintptr(h)) }
func UnsafeNewQTextListOwned(h unsafe.Pointer) *QTextList { return newownedQTextList(uintptr(h)) }

func (this *QTextList) UnsafeRelease() {
	this.h = 0
	if this.QTextBlockGroup != nil {
		this.QTextBlockGroup.UnsafeRelease()
	}
}

func NewQTextList(doc *QTextDocument) *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextList_new, uintptr(doc.UnsafePointer()))
	__o := newownedQTextList(__r)
	runtime.KeepAlive(doc)
	return __o
}

func QTextList_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextList_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTextList) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QTextList_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextList) Item(i int) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextList_item, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextList) ItemNumber(arg1 *QTextBlock) int {
	__r0, _, _ := purego.SyscallN(addr_QTextList_itemNumber, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QTextList) ItemText(arg1 *QTextBlock) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextList_itemText, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func (this *QTextList) RemoveItem(i int) {
	purego.SyscallN(addr_QTextList_removeItem, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QTextList) Remove(arg1 *QTextBlock) {
	purego.SyscallN(addr_QTextList_remove, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QTextList) Add(block *QTextBlock) {
	purego.SyscallN(addr_QTextList_add, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
}

func (this *QTextList) SetFormat(format *QTextListFormat) {
	purego.SyscallN(addr_QTextList_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextList) Format() *QTextListFormat {
	__r, _, _ := purego.SyscallN(addr_QTextList_format, this.h)
	runtime.KeepAlive(this)
	return newownedQTextListFormat(__r)
}

func QTextList_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextList_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextList_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextList_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QTextList_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextList_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextList_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextList_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTextList_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTextList_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTextList_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTextList_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTextList) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTextList_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextList) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTextList_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextList) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTextList_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTextList) Disconnect(handle int64) {
	pg_QTextList__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QTextList) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextList_delete(this.h)
	}
	this.UnsafeRelease()
}
