package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QKeyEvent_chain uintptr
	addr_QKeyEvent_new uintptr
	addr_QKeyEvent_new2 uintptr
	addr_QKeyEvent_new3 uintptr
	addr_QKeyEvent_new4 uintptr
	addr_QKeyEvent_new5 uintptr
	addr_QKeyEvent_new6 uintptr
	addr_QKeyEvent_new7 uintptr
	addr_QKeyEvent_new8 uintptr
	addr_QKeyEvent_new9 uintptr
	addr_QKeyEvent_key uintptr
	addr_QKeyEvent_matches uintptr
	addr_QKeyEvent_modifiers uintptr
	addr_QKeyEvent_keyCombination uintptr
	addr_QKeyEvent_text uintptr
	addr_QKeyEvent_isAutoRepeat uintptr
	addr_QKeyEvent_count uintptr
	addr_QKeyEvent_nativeScanCode uintptr
	addr_QKeyEvent_nativeVirtualKey uintptr
	addr_QKeyEvent_nativeModifiers uintptr
	addr_QKeyEvent_registerEventType uintptr
	addr_QKeyEvent_registerEventType2 uintptr
	pg_QKeyEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QKeyEvent_chain, "QKeyEvent_chain")
	symaddr(&addr_QKeyEvent_new, "QKeyEvent_new")
	symaddr(&addr_QKeyEvent_new2, "QKeyEvent_new2")
	symaddr(&addr_QKeyEvent_new3, "QKeyEvent_new3")
	symaddr(&addr_QKeyEvent_new4, "QKeyEvent_new4")
	symaddr(&addr_QKeyEvent_new5, "QKeyEvent_new5")
	symaddr(&addr_QKeyEvent_new6, "QKeyEvent_new6")
	symaddr(&addr_QKeyEvent_new7, "QKeyEvent_new7")
	symaddr(&addr_QKeyEvent_new8, "QKeyEvent_new8")
	symaddr(&addr_QKeyEvent_new9, "QKeyEvent_new9")
	symaddr(&addr_QKeyEvent_key, "QKeyEvent_key")
	symaddr(&addr_QKeyEvent_matches, "QKeyEvent_matches")
	symaddr(&addr_QKeyEvent_modifiers, "QKeyEvent_modifiers")
	symaddr(&addr_QKeyEvent_keyCombination, "QKeyEvent_keyCombination")
	symaddr(&addr_QKeyEvent_text, "QKeyEvent_text")
	symaddr(&addr_QKeyEvent_isAutoRepeat, "QKeyEvent_isAutoRepeat")
	symaddr(&addr_QKeyEvent_count, "QKeyEvent_count")
	symaddr(&addr_QKeyEvent_nativeScanCode, "QKeyEvent_nativeScanCode")
	symaddr(&addr_QKeyEvent_nativeVirtualKey, "QKeyEvent_nativeVirtualKey")
	symaddr(&addr_QKeyEvent_nativeModifiers, "QKeyEvent_nativeModifiers")
	symaddr(&addr_QKeyEvent_registerEventType, "QKeyEvent_registerEventType")
	symaddr(&addr_QKeyEvent_registerEventType2, "QKeyEvent_registerEventType2")
	reg(&pg_QKeyEvent_delete, "QKeyEvent_delete")
}

type QKeyEvent struct {
	h uintptr
	*QInputEvent
}

func bindQKeyEvent(h uintptr) *QKeyEvent {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QKeyEvent_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QKeyEvent{h: h, QInputEvent: &QInputEvent{h: __c[0], QEvent: &QEvent{h: __c[1]}}}
}

func newownedQKeyEvent(h uintptr) *QKeyEvent {
	o := bindQKeyEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QKeyEvent).Delete)
	}
	return o
}

func (this *QKeyEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QKeyEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQKeyEvent(h unsafe.Pointer) *QKeyEvent { return bindQKeyEvent(uintptr(h)) }
func UnsafeNewQKeyEventOwned(h unsafe.Pointer) *QKeyEvent { return newownedQKeyEvent(uintptr(h)) }

func (this *QKeyEvent) UnsafeRelease() {
	this.h = 0
	if this.QInputEvent != nil {
		this.QInputEvent.UnsafeRelease()
	}
}

func NewQKeyEvent(type_ QEvent__Type, key int, modifiers KeyboardModifier) *QKeyEvent {
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new, uintptr(type_), uintptr(key), uintptr(modifiers))
	__o := newownedQKeyEvent(__r)
	return __o
}

func NewQKeyEvent2(type_ QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint) *QKeyEvent {
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new2, uintptr(type_), uintptr(key), uintptr(modifiers), uintptr(nativeScanCode), uintptr(nativeVirtualKey), uintptr(nativeModifiers))
	__o := newownedQKeyEvent(__r)
	return __o
}

func NewQKeyEvent3(type_ QEvent__Type, key int, modifiers KeyboardModifier, text string) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new3, uintptr(type_), uintptr(key), uintptr(modifiers), text_p, uintptr(len(text)))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent4(type_ QEvent__Type, key int, modifiers KeyboardModifier, text string, autorep bool) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new4, uintptr(type_), uintptr(key), uintptr(modifiers), text_p, uintptr(len(text)), uintptr(b2i(autorep)))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent5(type_ QEvent__Type, key int, modifiers KeyboardModifier, text string, autorep bool, count uint16) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new5, uintptr(type_), uintptr(key), uintptr(modifiers), text_p, uintptr(len(text)), uintptr(b2i(autorep)), uintptr(count))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent6(type_ QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new6, uintptr(type_), uintptr(key), uintptr(modifiers), uintptr(nativeScanCode), uintptr(nativeVirtualKey), uintptr(nativeModifiers), text_p, uintptr(len(text)))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent7(type_ QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new7, uintptr(type_), uintptr(key), uintptr(modifiers), uintptr(nativeScanCode), uintptr(nativeVirtualKey), uintptr(nativeModifiers), text_p, uintptr(len(text)), uintptr(b2i(autorep)))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent8(type_ QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new8, uintptr(type_), uintptr(key), uintptr(modifiers), uintptr(nativeScanCode), uintptr(nativeVirtualKey), uintptr(nativeModifiers), text_p, uintptr(len(text)), uintptr(b2i(autorep)), uintptr(count))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	return __o
}

func NewQKeyEvent9(type_ QEvent__Type, key int, modifiers KeyboardModifier, nativeScanCode uint, nativeVirtualKey uint, nativeModifiers uint, text string, autorep bool, count uint16, device *QInputDevice) *QKeyEvent {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_new9, uintptr(type_), uintptr(key), uintptr(modifiers), uintptr(nativeScanCode), uintptr(nativeVirtualKey), uintptr(nativeModifiers), text_p, uintptr(len(text)), uintptr(b2i(autorep)), uintptr(count), uintptr(device.UnsafePointer()))
	__o := newownedQKeyEvent(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(device)
	return __o
}

func (this *QKeyEvent) Key() int {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_key, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeyEvent) Matches(key QKeySequence__StandardKey) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_matches, this.h, uintptr(key))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeyEvent) Modifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_modifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QKeyEvent) KeyCombination() *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_keyCombination, this.h)
	runtime.KeepAlive(this)
	return newownedQKeyCombination(__r)
}

func (this *QKeyEvent) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeyEvent_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QKeyEvent) IsAutoRepeat() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_isAutoRepeat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeyEvent) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeyEvent) NativeScanCode() uint {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_nativeScanCode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QKeyEvent) NativeVirtualKey() uint {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_nativeVirtualKey, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QKeyEvent) NativeModifiers() uint {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_nativeModifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func QKeyEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QKeyEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QKeyEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QKeyEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QKeyEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
