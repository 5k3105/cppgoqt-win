package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QKeyCombination_new uintptr
	addr_QKeyCombination_new2 uintptr
	addr_QKeyCombination_new3 uintptr
	addr_QKeyCombination_new4 uintptr
	addr_QKeyCombination_new5 uintptr
	addr_QKeyCombination_new6 uintptr
	addr_QKeyCombination_new7 uintptr
	addr_QKeyCombination_keyboardModifiers uintptr
	addr_QKeyCombination_key uintptr
	addr_QKeyCombination_fromCombined uintptr
	addr_QKeyCombination_toCombined uintptr
	addr_QKeyCombination_ToInt uintptr
	addr_QKeyCombination_OperatorEqual uintptr
	addr_QKeyCombination_OperatorNotEqual uintptr
	pg_QKeyCombination_delete func(uintptr)
)

func init() {
	symaddr(&addr_QKeyCombination_new, "QKeyCombination_new")
	symaddr(&addr_QKeyCombination_new2, "QKeyCombination_new2")
	symaddr(&addr_QKeyCombination_new3, "QKeyCombination_new3")
	symaddr(&addr_QKeyCombination_new4, "QKeyCombination_new4")
	symaddr(&addr_QKeyCombination_new5, "QKeyCombination_new5")
	symaddr(&addr_QKeyCombination_new6, "QKeyCombination_new6")
	symaddr(&addr_QKeyCombination_new7, "QKeyCombination_new7")
	symaddr(&addr_QKeyCombination_keyboardModifiers, "QKeyCombination_keyboardModifiers")
	symaddr(&addr_QKeyCombination_key, "QKeyCombination_key")
	symaddr(&addr_QKeyCombination_fromCombined, "QKeyCombination_fromCombined")
	symaddr(&addr_QKeyCombination_toCombined, "QKeyCombination_toCombined")
	symaddr(&addr_QKeyCombination_ToInt, "QKeyCombination_ToInt")
	symaddr(&addr_QKeyCombination_OperatorEqual, "QKeyCombination_OperatorEqual")
	symaddr(&addr_QKeyCombination_OperatorNotEqual, "QKeyCombination_OperatorNotEqual")
	reg(&pg_QKeyCombination_delete, "QKeyCombination_delete")
}

type QKeyCombination struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QKeyCombination) UnsafeGoID() int64     { return this.vid }
func (this *QKeyCombination) UnsafeSetGoID(id int64) { this.vid = id }

func bindQKeyCombination(h uintptr) *QKeyCombination {
	if h == 0 {
		return nil
	}
	return &QKeyCombination{h: h}
}

func newownedQKeyCombination(h uintptr) *QKeyCombination {
	o := bindQKeyCombination(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QKeyCombination).Delete)
	}
	return o
}

func (this *QKeyCombination) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QKeyCombination) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQKeyCombination(h unsafe.Pointer) *QKeyCombination { return bindQKeyCombination(uintptr(h)) }
func UnsafeNewQKeyCombinationOwned(h unsafe.Pointer) *QKeyCombination { return newownedQKeyCombination(uintptr(h)) }

func (this *QKeyCombination) UnsafeRelease() { this.h = 0 }

func NewQKeyCombination() *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new)
	__o := newownedQKeyCombination(__r)
	return __o
}

func NewQKeyCombination2(modifiers Modifier) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new2, uintptr(modifiers))
	__o := newownedQKeyCombination(__r)
	return __o
}

func NewQKeyCombination3(modifiers KeyboardModifier) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new3, uintptr(modifiers))
	__o := newownedQKeyCombination(__r)
	return __o
}

func NewQKeyCombination4(arg1 *QKeyCombination) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQKeyCombination(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQKeyCombination5(key Key) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new5, uintptr(key))
	__o := newownedQKeyCombination(__r)
	return __o
}

func NewQKeyCombination6(modifiers Modifier, key Key) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new6, uintptr(modifiers), uintptr(key))
	__o := newownedQKeyCombination(__r)
	return __o
}

func NewQKeyCombination7(modifiers KeyboardModifier, key Key) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_new7, uintptr(modifiers), uintptr(key))
	__o := newownedQKeyCombination(__r)
	return __o
}

func (this *QKeyCombination) KeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_keyboardModifiers, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return KeyboardModifier(__r)
}

func (this *QKeyCombination) Key() Key {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_key, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Key(__r)
}

func QKeyCombination_FromCombined(combined int) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeyCombination_fromCombined, uintptr(combined))
	return newownedQKeyCombination(__r)
}

func (this *QKeyCombination) ToCombined() int {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_toCombined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeyCombination) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_ToInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeyCombination) OperatorEqual(rhs *QKeyCombination) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QKeyCombination) OperatorNotEqual(rhs *QKeyCombination) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeyCombination_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QKeyCombination) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QKeyCombination_delete(this.h)
	}
	this.UnsafeRelease()
}
