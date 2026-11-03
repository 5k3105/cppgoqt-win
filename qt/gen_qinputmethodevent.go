package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QInputMethodEvent_base func(uintptr) uintptr
	addr_QInputMethodEvent_new uintptr
	addr_QInputMethodEvent_new2 uintptr
	addr_QInputMethodEvent_setCommitString uintptr
	addr_QInputMethodEvent_attributes uintptr
	addr_QInputMethodEvent_preeditString uintptr
	addr_QInputMethodEvent_commitString uintptr
	addr_QInputMethodEvent_replacementStart uintptr
	addr_QInputMethodEvent_replacementLength uintptr
	addr_QInputMethodEvent_registerEventType uintptr
	addr_QInputMethodEvent_setCommitString2 uintptr
	addr_QInputMethodEvent_setCommitString3 uintptr
	addr_QInputMethodEvent_registerEventType2 uintptr
	addr_QInputMethodEvent_clone_default uintptr
	addr_QInputMethodEvent_setAccepted_default uintptr
	pg_QInputMethodEvent__set_override func(uintptr, int32)
	pg_QInputMethodEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QInputMethodEvent_base, "QInputMethodEvent_base")
	symaddr(&addr_QInputMethodEvent_new, "QInputMethodEvent_new")
	symaddr(&addr_QInputMethodEvent_new2, "QInputMethodEvent_new2")
	symaddr(&addr_QInputMethodEvent_setCommitString, "QInputMethodEvent_setCommitString")
	symaddr(&addr_QInputMethodEvent_attributes, "QInputMethodEvent_attributes")
	symaddr(&addr_QInputMethodEvent_preeditString, "QInputMethodEvent_preeditString")
	symaddr(&addr_QInputMethodEvent_commitString, "QInputMethodEvent_commitString")
	symaddr(&addr_QInputMethodEvent_replacementStart, "QInputMethodEvent_replacementStart")
	symaddr(&addr_QInputMethodEvent_replacementLength, "QInputMethodEvent_replacementLength")
	symaddr(&addr_QInputMethodEvent_registerEventType, "QInputMethodEvent_registerEventType")
	symaddr(&addr_QInputMethodEvent_setCommitString2, "QInputMethodEvent_setCommitString2")
	symaddr(&addr_QInputMethodEvent_setCommitString3, "QInputMethodEvent_setCommitString3")
	symaddr(&addr_QInputMethodEvent_registerEventType2, "QInputMethodEvent_registerEventType2")
	symaddr(&addr_QInputMethodEvent_clone_default, "QInputMethodEvent_clone_default")
	symaddr(&addr_QInputMethodEvent_setAccepted_default, "QInputMethodEvent_setAccepted_default")
	reg(&pg_QInputMethodEvent__set_override, "QInputMethodEvent__set_override")
	reg(&pg_QInputMethodEvent_delete, "QInputMethodEvent_delete")
}

type QInputMethodEvent struct {
	h uintptr
	*QEvent
}

func bindQInputMethodEvent(h uintptr) *QInputMethodEvent {
	if h == 0 {
		return nil
	}
	return &QInputMethodEvent{h: h, QEvent: bindQEvent(pg_QInputMethodEvent_base(h))}
}

func newownedQInputMethodEvent(h uintptr) *QInputMethodEvent {
	o := bindQInputMethodEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QInputMethodEvent).Delete)
	}
	return o
}

func (this *QInputMethodEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputMethodEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputMethodEvent(h unsafe.Pointer) *QInputMethodEvent { return bindQInputMethodEvent(uintptr(h)) }
func UnsafeNewQInputMethodEventOwned(h unsafe.Pointer) *QInputMethodEvent { return newownedQInputMethodEvent(uintptr(h)) }

func (this *QInputMethodEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQInputMethodEvent() *QInputMethodEvent {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_new, uintptr(__goid))
	__o := newownedQInputMethodEvent(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputMethodEvent_virtIdx, QInputMethodEvent_enable)
	return __o
}

func NewQInputMethodEvent2(preeditText string, attributes []*QInputMethodEvent__Attribute) *QInputMethodEvent {
	__goid := nextVirtID()
	var preeditText_p uintptr
	if len(preeditText) > 0 {
		preeditText_p = uintptr(unsafe.Pointer(unsafe.StringData(preeditText)))
	}
	attributes_arr := make([]uintptr, len(attributes))
	for __i, __e := range attributes {
		attributes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var attributes_p uintptr
	if len(attributes_arr) > 0 {
		attributes_p = uintptr(unsafe.Pointer(&attributes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_new2, uintptr(__goid), preeditText_p, uintptr(len(preeditText)), attributes_p, uintptr(len(attributes)))
	__o := newownedQInputMethodEvent(__r)
	runtime.KeepAlive(preeditText)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(attributes_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QInputMethodEvent_virtIdx, QInputMethodEvent_enable)
	return __o
}

func (this *QInputMethodEvent) SetCommitString(commitString string) {
	var commitString_p uintptr
	if len(commitString) > 0 {
		commitString_p = uintptr(unsafe.Pointer(unsafe.StringData(commitString)))
	}
	purego.SyscallN(addr_QInputMethodEvent_setCommitString, this.h, commitString_p, uintptr(len(commitString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(commitString)
}

func (this *QInputMethodEvent) Attributes() []*QInputMethodEvent__Attribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_attributes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QInputMethodEvent__Attribute, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQInputMethodEvent__Attribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QInputMethodEvent) PreeditString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_preeditString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputMethodEvent) CommitString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_commitString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QInputMethodEvent) ReplacementStart() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent_replacementStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputMethodEvent) ReplacementLength() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent_replacementLength, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QInputMethodEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QInputMethodEvent) SetCommitString2(commitString string, replaceFrom int) {
	var commitString_p uintptr
	if len(commitString) > 0 {
		commitString_p = uintptr(unsafe.Pointer(unsafe.StringData(commitString)))
	}
	purego.SyscallN(addr_QInputMethodEvent_setCommitString2, this.h, commitString_p, uintptr(len(commitString)), uintptr(replaceFrom))
	runtime.KeepAlive(this)
	runtime.KeepAlive(commitString)
}

func (this *QInputMethodEvent) SetCommitString3(commitString string, replaceFrom int, replaceLength int) {
	var commitString_p uintptr
	if len(commitString) > 0 {
		commitString_p = uintptr(unsafe.Pointer(unsafe.StringData(commitString)))
	}
	purego.SyscallN(addr_QInputMethodEvent_setCommitString3, this.h, commitString_p, uintptr(len(commitString)), uintptr(replaceFrom), uintptr(replaceLength))
	runtime.KeepAlive(this)
	runtime.KeepAlive(commitString)
}

func QInputMethodEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QInputMethodEvent) CloneDefault() *QInputMethodEvent {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQInputMethodEvent(__r)
}

func (this *QInputMethodEvent) SetAcceptedDefault(accepted bool) {
	purego.SyscallN(addr_QInputMethodEvent_setAccepted_default, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func QInputMethodEvent_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case 6367896784077321574:
		return 1
	}
	return -1
}

func QInputMethodEvent_enable(h uintptr, idx int32) { pg_QInputMethodEvent__set_override(h, idx) }

func (this *QInputMethodEvent) OnClone(fn func() *QInputMethodEvent) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QInputMethodEvent) OnSetAccepted(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 6367896784077321574, func(accepted int32) { fn(accepted != 0) }) // slot "setAccepted"
}

func (this *QInputMethodEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputMethodEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
