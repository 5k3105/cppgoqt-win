package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QShortcutEvent_base func(uintptr) uintptr
	addr_QShortcutEvent_new uintptr
	addr_QShortcutEvent_new2 uintptr
	addr_QShortcutEvent_new3 uintptr
	addr_QShortcutEvent_new4 uintptr
	addr_QShortcutEvent_new5 uintptr
	addr_QShortcutEvent_key uintptr
	addr_QShortcutEvent_shortcutId uintptr
	addr_QShortcutEvent_isAmbiguous uintptr
	addr_QShortcutEvent_registerEventType uintptr
	addr_QShortcutEvent_registerEventType2 uintptr
	pg_QShortcutEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QShortcutEvent_base, "QShortcutEvent_base")
	symaddr(&addr_QShortcutEvent_new, "QShortcutEvent_new")
	symaddr(&addr_QShortcutEvent_new2, "QShortcutEvent_new2")
	symaddr(&addr_QShortcutEvent_new3, "QShortcutEvent_new3")
	symaddr(&addr_QShortcutEvent_new4, "QShortcutEvent_new4")
	symaddr(&addr_QShortcutEvent_new5, "QShortcutEvent_new5")
	symaddr(&addr_QShortcutEvent_key, "QShortcutEvent_key")
	symaddr(&addr_QShortcutEvent_shortcutId, "QShortcutEvent_shortcutId")
	symaddr(&addr_QShortcutEvent_isAmbiguous, "QShortcutEvent_isAmbiguous")
	symaddr(&addr_QShortcutEvent_registerEventType, "QShortcutEvent_registerEventType")
	symaddr(&addr_QShortcutEvent_registerEventType2, "QShortcutEvent_registerEventType2")
	reg(&pg_QShortcutEvent_delete, "QShortcutEvent_delete")
}

type QShortcutEvent struct {
	h uintptr
	*QEvent
}

func bindQShortcutEvent(h uintptr) *QShortcutEvent {
	if h == 0 {
		return nil
	}
	return &QShortcutEvent{h: h, QEvent: bindQEvent(pg_QShortcutEvent_base(h))}
}

func newownedQShortcutEvent(h uintptr) *QShortcutEvent {
	o := bindQShortcutEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QShortcutEvent).Delete)
	}
	return o
}

func (this *QShortcutEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QShortcutEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQShortcutEvent(h unsafe.Pointer) *QShortcutEvent { return bindQShortcutEvent(uintptr(h)) }
func UnsafeNewQShortcutEventOwned(h unsafe.Pointer) *QShortcutEvent { return newownedQShortcutEvent(uintptr(h)) }

func (this *QShortcutEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQShortcutEvent(key *QKeySequence, id int) *QShortcutEvent {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_new, uintptr(key.UnsafePointer()), uintptr(id))
	__o := newownedQShortcutEvent(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQShortcutEvent2(key *QKeySequence) *QShortcutEvent {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_new2, uintptr(key.UnsafePointer()))
	__o := newownedQShortcutEvent(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQShortcutEvent3(key *QKeySequence, id int, ambiguous bool) *QShortcutEvent {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_new3, uintptr(key.UnsafePointer()), uintptr(id), uintptr(b2i(ambiguous)))
	__o := newownedQShortcutEvent(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQShortcutEvent4(key *QKeySequence, shortcut *QShortcut) *QShortcutEvent {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_new4, uintptr(key.UnsafePointer()), uintptr(shortcut.UnsafePointer()))
	__o := newownedQShortcutEvent(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(shortcut)
	return __o
}

func NewQShortcutEvent5(key *QKeySequence, shortcut *QShortcut, ambiguous bool) *QShortcutEvent {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_new5, uintptr(key.UnsafePointer()), uintptr(shortcut.UnsafePointer()), uintptr(b2i(ambiguous)))
	__o := newownedQShortcutEvent(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(shortcut)
	return __o
}

func (this *QShortcutEvent) Key() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QShortcutEvent_key, this.h)
	runtime.KeepAlive(this)
	return bindQKeySequence(__r)
}

func (this *QShortcutEvent) ShortcutId() int {
	__r0, _, _ := purego.SyscallN(addr_QShortcutEvent_shortcutId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QShortcutEvent) IsAmbiguous() bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcutEvent_isAmbiguous, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QShortcutEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QShortcutEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QShortcutEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QShortcutEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QShortcutEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QShortcutEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
