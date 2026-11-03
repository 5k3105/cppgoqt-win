package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringConverterBase__State_new uintptr
	addr_QStringConverterBase__State_new2 uintptr
	addr_QStringConverterBase__State_clear uintptr
	addr_QStringConverterBase__State_reset uintptr
	addr_QStringConverterBase__State_flags_get uintptr
	addr_QStringConverterBase__State_flags_set uintptr
	addr_QStringConverterBase__State_internalState_get uintptr
	addr_QStringConverterBase__State_internalState_set uintptr
	addr_QStringConverterBase__State_remainingChars_get uintptr
	addr_QStringConverterBase__State_remainingChars_set uintptr
	addr_QStringConverterBase__State_invalidChars_get uintptr
	addr_QStringConverterBase__State_invalidChars_set uintptr
	pg_QStringConverterBase__State_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringConverterBase__State_new, "QStringConverterBase__State_new")
	symaddr(&addr_QStringConverterBase__State_new2, "QStringConverterBase__State_new2")
	symaddr(&addr_QStringConverterBase__State_clear, "QStringConverterBase__State_clear")
	symaddr(&addr_QStringConverterBase__State_reset, "QStringConverterBase__State_reset")
	symaddr(&addr_QStringConverterBase__State_flags_get, "QStringConverterBase__State_flags_get")
	symaddr(&addr_QStringConverterBase__State_flags_set, "QStringConverterBase__State_flags_set")
	symaddr(&addr_QStringConverterBase__State_internalState_get, "QStringConverterBase__State_internalState_get")
	symaddr(&addr_QStringConverterBase__State_internalState_set, "QStringConverterBase__State_internalState_set")
	symaddr(&addr_QStringConverterBase__State_remainingChars_get, "QStringConverterBase__State_remainingChars_get")
	symaddr(&addr_QStringConverterBase__State_remainingChars_set, "QStringConverterBase__State_remainingChars_set")
	symaddr(&addr_QStringConverterBase__State_invalidChars_get, "QStringConverterBase__State_invalidChars_get")
	symaddr(&addr_QStringConverterBase__State_invalidChars_set, "QStringConverterBase__State_invalidChars_set")
	reg(&pg_QStringConverterBase__State_delete, "QStringConverterBase__State_delete")
}

type QStringConverterBase__State struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStringConverterBase__State) UnsafeGoID() int64     { return this.vid }
func (this *QStringConverterBase__State) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStringConverterBase__State(h uintptr) *QStringConverterBase__State {
	if h == 0 {
		return nil
	}
	return &QStringConverterBase__State{h: h}
}

func newownedQStringConverterBase__State(h uintptr) *QStringConverterBase__State {
	o := bindQStringConverterBase__State(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStringConverterBase__State).Delete)
	}
	return o
}

func (this *QStringConverterBase__State) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringConverterBase__State) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringConverterBase__State(h unsafe.Pointer) *QStringConverterBase__State { return bindQStringConverterBase__State(uintptr(h)) }
func UnsafeNewQStringConverterBase__StateOwned(h unsafe.Pointer) *QStringConverterBase__State { return newownedQStringConverterBase__State(uintptr(h)) }

func (this *QStringConverterBase__State) UnsafeRelease() { this.h = 0 }

func NewQStringConverterBase__State() *QStringConverterBase__State {
	__r, _, _ := purego.SyscallN(addr_QStringConverterBase__State_new)
	__o := newownedQStringConverterBase__State(__r)
	return __o
}

func NewQStringConverterBase__State2(f QStringConverterBase__Flag) *QStringConverterBase__State {
	__r, _, _ := purego.SyscallN(addr_QStringConverterBase__State_new2, uintptr(f))
	__o := newownedQStringConverterBase__State(__r)
	return __o
}

func (this *QStringConverterBase__State) Clear() {
	purego.SyscallN(addr_QStringConverterBase__State_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) Reset() {
	purego.SyscallN(addr_QStringConverterBase__State_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) Flags() QStringConverterBase__Flag {
	__r0, _, _ := purego.SyscallN(addr_QStringConverterBase__State_flags_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStringConverterBase__Flag(__r)
}

func (this *QStringConverterBase__State) SetFlags(flags QStringConverterBase__Flag) {
	purego.SyscallN(addr_QStringConverterBase__State_flags_set, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) InternalState() int {
	__r0, _, _ := purego.SyscallN(addr_QStringConverterBase__State_internalState_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStringConverterBase__State) SetInternalState(internalState int) {
	purego.SyscallN(addr_QStringConverterBase__State_internalState_set, this.h, uintptr(internalState))
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) RemainingChars() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringConverterBase__State_remainingChars_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringConverterBase__State) SetRemainingChars(remainingChars int64) {
	purego.SyscallN(addr_QStringConverterBase__State_remainingChars_set, this.h, uintptr(remainingChars))
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) InvalidChars() int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringConverterBase__State_invalidChars_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringConverterBase__State) SetInvalidChars(invalidChars int64) {
	purego.SyscallN(addr_QStringConverterBase__State_invalidChars_set, this.h, uintptr(invalidChars))
	runtime.KeepAlive(this)
}

func (this *QStringConverterBase__State) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringConverterBase__State_delete(this.h)
	}
	this.UnsafeRelease()
}
