package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QObjectData_dynamicMetaObject uintptr
	addr_QObjectData_q_ptr_get uintptr
	addr_QObjectData_q_ptr_set uintptr
	addr_QObjectData_parent_get uintptr
	addr_QObjectData_parent_set uintptr
	addr_QObjectData_children_get uintptr
	addr_QObjectData_children_set uintptr
	addr_QObjectData_isWidget_get uintptr
	addr_QObjectData_isWidget_set uintptr
	addr_QObjectData_blockSig_get uintptr
	addr_QObjectData_blockSig_set uintptr
	addr_QObjectData_wasDeleted_get uintptr
	addr_QObjectData_wasDeleted_set uintptr
	addr_QObjectData_isDeletingChildren_get uintptr
	addr_QObjectData_isDeletingChildren_set uintptr
	addr_QObjectData_sendChildEvents_get uintptr
	addr_QObjectData_sendChildEvents_set uintptr
	addr_QObjectData_receiveChildEvents_get uintptr
	addr_QObjectData_receiveChildEvents_set uintptr
	addr_QObjectData_isWindow_get uintptr
	addr_QObjectData_isWindow_set uintptr
	addr_QObjectData_deleteLaterCalled_get uintptr
	addr_QObjectData_deleteLaterCalled_set uintptr
	addr_QObjectData_isQuickItem_get uintptr
	addr_QObjectData_isQuickItem_set uintptr
	addr_QObjectData_willBeWidget_get uintptr
	addr_QObjectData_willBeWidget_set uintptr
	addr_QObjectData_wasWidget_get uintptr
	addr_QObjectData_wasWidget_set uintptr
	addr_QObjectData_receiveParentEvents_get uintptr
	addr_QObjectData_receiveParentEvents_set uintptr
	addr_QObjectData_unused_get uintptr
	addr_QObjectData_unused_set uintptr
	addr_QObjectData_postedEvents_get uintptr
	addr_QObjectData_postedEvents_set uintptr
	addr_QObjectData_bindingStorage_get uintptr
	addr_QObjectData_bindingStorage_set uintptr
	pg_QObjectData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QObjectData_dynamicMetaObject, "QObjectData_dynamicMetaObject")
	symaddr(&addr_QObjectData_q_ptr_get, "QObjectData_q_ptr_get")
	symaddr(&addr_QObjectData_q_ptr_set, "QObjectData_q_ptr_set")
	symaddr(&addr_QObjectData_parent_get, "QObjectData_parent_get")
	symaddr(&addr_QObjectData_parent_set, "QObjectData_parent_set")
	symaddr(&addr_QObjectData_children_get, "QObjectData_children_get")
	symaddr(&addr_QObjectData_children_set, "QObjectData_children_set")
	symaddr(&addr_QObjectData_isWidget_get, "QObjectData_isWidget_get")
	symaddr(&addr_QObjectData_isWidget_set, "QObjectData_isWidget_set")
	symaddr(&addr_QObjectData_blockSig_get, "QObjectData_blockSig_get")
	symaddr(&addr_QObjectData_blockSig_set, "QObjectData_blockSig_set")
	symaddr(&addr_QObjectData_wasDeleted_get, "QObjectData_wasDeleted_get")
	symaddr(&addr_QObjectData_wasDeleted_set, "QObjectData_wasDeleted_set")
	symaddr(&addr_QObjectData_isDeletingChildren_get, "QObjectData_isDeletingChildren_get")
	symaddr(&addr_QObjectData_isDeletingChildren_set, "QObjectData_isDeletingChildren_set")
	symaddr(&addr_QObjectData_sendChildEvents_get, "QObjectData_sendChildEvents_get")
	symaddr(&addr_QObjectData_sendChildEvents_set, "QObjectData_sendChildEvents_set")
	symaddr(&addr_QObjectData_receiveChildEvents_get, "QObjectData_receiveChildEvents_get")
	symaddr(&addr_QObjectData_receiveChildEvents_set, "QObjectData_receiveChildEvents_set")
	symaddr(&addr_QObjectData_isWindow_get, "QObjectData_isWindow_get")
	symaddr(&addr_QObjectData_isWindow_set, "QObjectData_isWindow_set")
	symaddr(&addr_QObjectData_deleteLaterCalled_get, "QObjectData_deleteLaterCalled_get")
	symaddr(&addr_QObjectData_deleteLaterCalled_set, "QObjectData_deleteLaterCalled_set")
	symaddr(&addr_QObjectData_isQuickItem_get, "QObjectData_isQuickItem_get")
	symaddr(&addr_QObjectData_isQuickItem_set, "QObjectData_isQuickItem_set")
	symaddr(&addr_QObjectData_willBeWidget_get, "QObjectData_willBeWidget_get")
	symaddr(&addr_QObjectData_willBeWidget_set, "QObjectData_willBeWidget_set")
	symaddr(&addr_QObjectData_wasWidget_get, "QObjectData_wasWidget_get")
	symaddr(&addr_QObjectData_wasWidget_set, "QObjectData_wasWidget_set")
	symaddr(&addr_QObjectData_receiveParentEvents_get, "QObjectData_receiveParentEvents_get")
	symaddr(&addr_QObjectData_receiveParentEvents_set, "QObjectData_receiveParentEvents_set")
	symaddr(&addr_QObjectData_unused_get, "QObjectData_unused_get")
	symaddr(&addr_QObjectData_unused_set, "QObjectData_unused_set")
	symaddr(&addr_QObjectData_postedEvents_get, "QObjectData_postedEvents_get")
	symaddr(&addr_QObjectData_postedEvents_set, "QObjectData_postedEvents_set")
	symaddr(&addr_QObjectData_bindingStorage_get, "QObjectData_bindingStorage_get")
	symaddr(&addr_QObjectData_bindingStorage_set, "QObjectData_bindingStorage_set")
	reg(&pg_QObjectData_delete, "QObjectData_delete")
}

type QObjectData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QObjectData) UnsafeGoID() int64     { return this.vid }
func (this *QObjectData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQObjectData(h uintptr) *QObjectData {
	if h == 0 {
		return nil
	}
	return &QObjectData{h: h}
}

func newownedQObjectData(h uintptr) *QObjectData {
	o := bindQObjectData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QObjectData).Delete)
	}
	return o
}

func (this *QObjectData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QObjectData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQObjectData(h unsafe.Pointer) *QObjectData { return bindQObjectData(uintptr(h)) }
func UnsafeNewQObjectDataOwned(h unsafe.Pointer) *QObjectData { return newownedQObjectData(uintptr(h)) }

func (this *QObjectData) UnsafeRelease() { this.h = 0 }

func (this *QObjectData) DynamicMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QObjectData_dynamicMetaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QObjectData) Q_ptr() *QObject {
	__r, _, _ := purego.SyscallN(addr_QObjectData_q_ptr_get, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QObjectData) SetQ_ptr(q_ptr *QObject) {
	purego.SyscallN(addr_QObjectData_q_ptr_set, this.h, uintptr(q_ptr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(q_ptr)
}

func (this *QObjectData) Parent() *QObject {
	__r, _, _ := purego.SyscallN(addr_QObjectData_parent_get, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QObjectData) SetParent(parent *QObject) {
	purego.SyscallN(addr_QObjectData_parent_set, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QObjectData) Children() []*QObject {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObjectData_children_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QObject, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQObject(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QObjectData) SetChildren(children []*QObject) {
	children_arr := make([]uintptr, len(children))
	for __i, __e := range children {
		children_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var children_p uintptr
	if len(children_arr) > 0 {
		children_p = uintptr(unsafe.Pointer(&children_arr[0]))
	}
	purego.SyscallN(addr_QObjectData_children_set, this.h, children_p, uintptr(len(children)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(children)
	runtime.KeepAlive(children_arr)
}

func (this *QObjectData) IsWidget() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_isWidget_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetIsWidget(isWidget uint) {
	purego.SyscallN(addr_QObjectData_isWidget_set, this.h, uintptr(isWidget))
	runtime.KeepAlive(this)
}

func (this *QObjectData) BlockSig() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_blockSig_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetBlockSig(blockSig uint) {
	purego.SyscallN(addr_QObjectData_blockSig_set, this.h, uintptr(blockSig))
	runtime.KeepAlive(this)
}

func (this *QObjectData) WasDeleted() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_wasDeleted_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetWasDeleted(wasDeleted uint) {
	purego.SyscallN(addr_QObjectData_wasDeleted_set, this.h, uintptr(wasDeleted))
	runtime.KeepAlive(this)
}

func (this *QObjectData) IsDeletingChildren() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_isDeletingChildren_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetIsDeletingChildren(isDeletingChildren uint) {
	purego.SyscallN(addr_QObjectData_isDeletingChildren_set, this.h, uintptr(isDeletingChildren))
	runtime.KeepAlive(this)
}

func (this *QObjectData) SendChildEvents() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_sendChildEvents_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetSendChildEvents(sendChildEvents uint) {
	purego.SyscallN(addr_QObjectData_sendChildEvents_set, this.h, uintptr(sendChildEvents))
	runtime.KeepAlive(this)
}

func (this *QObjectData) ReceiveChildEvents() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_receiveChildEvents_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetReceiveChildEvents(receiveChildEvents uint) {
	purego.SyscallN(addr_QObjectData_receiveChildEvents_set, this.h, uintptr(receiveChildEvents))
	runtime.KeepAlive(this)
}

func (this *QObjectData) IsWindow() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_isWindow_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetIsWindow(isWindow uint) {
	purego.SyscallN(addr_QObjectData_isWindow_set, this.h, uintptr(isWindow))
	runtime.KeepAlive(this)
}

func (this *QObjectData) DeleteLaterCalled() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_deleteLaterCalled_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetDeleteLaterCalled(deleteLaterCalled uint) {
	purego.SyscallN(addr_QObjectData_deleteLaterCalled_set, this.h, uintptr(deleteLaterCalled))
	runtime.KeepAlive(this)
}

func (this *QObjectData) IsQuickItem() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_isQuickItem_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetIsQuickItem(isQuickItem uint) {
	purego.SyscallN(addr_QObjectData_isQuickItem_set, this.h, uintptr(isQuickItem))
	runtime.KeepAlive(this)
}

func (this *QObjectData) WillBeWidget() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_willBeWidget_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetWillBeWidget(willBeWidget uint) {
	purego.SyscallN(addr_QObjectData_willBeWidget_set, this.h, uintptr(willBeWidget))
	runtime.KeepAlive(this)
}

func (this *QObjectData) WasWidget() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_wasWidget_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetWasWidget(wasWidget uint) {
	purego.SyscallN(addr_QObjectData_wasWidget_set, this.h, uintptr(wasWidget))
	runtime.KeepAlive(this)
}

func (this *QObjectData) ReceiveParentEvents() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_receiveParentEvents_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetReceiveParentEvents(receiveParentEvents uint) {
	purego.SyscallN(addr_QObjectData_receiveParentEvents_set, this.h, uintptr(receiveParentEvents))
	runtime.KeepAlive(this)
}

func (this *QObjectData) Unused() uint {
	__r0, _, _ := purego.SyscallN(addr_QObjectData_unused_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QObjectData) SetUnused(unused uint) {
	purego.SyscallN(addr_QObjectData_unused_set, this.h, uintptr(unused))
	runtime.KeepAlive(this)
}

func (this *QObjectData) PostedEvents() *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QObjectData_postedEvents_get, this.h)
	runtime.KeepAlive(this)
	return newownedQAtomicInt(__r)
}

func (this *QObjectData) SetPostedEvents(postedEvents *QAtomicInt) {
	purego.SyscallN(addr_QObjectData_postedEvents_set, this.h, uintptr(postedEvents.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(postedEvents)
}

func (this *QObjectData) BindingStorage() *QBindingStorage {
	__r, _, _ := purego.SyscallN(addr_QObjectData_bindingStorage_get, this.h)
	runtime.KeepAlive(this)
	return newownedQBindingStorage(__r)
}

func (this *QObjectData) SetBindingStorage(bindingStorage *QBindingStorage) {
	purego.SyscallN(addr_QObjectData_bindingStorage_set, this.h, uintptr(bindingStorage.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bindingStorage)
}

func (this *QObjectData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QObjectData_delete(this.h)
	}
	this.UnsafeRelease()
}
