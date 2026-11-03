package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QNoDebug_OperatorShiftLeft uintptr
	addr_QNoDebug_space uintptr
	addr_QNoDebug_nospace uintptr
	addr_QNoDebug_maybeSpace uintptr
	addr_QNoDebug_quote uintptr
	addr_QNoDebug_noquote uintptr
	addr_QNoDebug_maybeQuote uintptr
	addr_QNoDebug_verbosity uintptr
	addr_QNoDebug_maybeQuote2 uintptr
	pg_QNoDebug_delete func(uintptr)
)

func init() {
	symaddr(&addr_QNoDebug_OperatorShiftLeft, "QNoDebug_OperatorShiftLeft")
	symaddr(&addr_QNoDebug_space, "QNoDebug_space")
	symaddr(&addr_QNoDebug_nospace, "QNoDebug_nospace")
	symaddr(&addr_QNoDebug_maybeSpace, "QNoDebug_maybeSpace")
	symaddr(&addr_QNoDebug_quote, "QNoDebug_quote")
	symaddr(&addr_QNoDebug_noquote, "QNoDebug_noquote")
	symaddr(&addr_QNoDebug_maybeQuote, "QNoDebug_maybeQuote")
	symaddr(&addr_QNoDebug_verbosity, "QNoDebug_verbosity")
	symaddr(&addr_QNoDebug_maybeQuote2, "QNoDebug_maybeQuote2")
	reg(&pg_QNoDebug_delete, "QNoDebug_delete")
}

type QNoDebug struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QNoDebug) UnsafeGoID() int64     { return this.vid }
func (this *QNoDebug) UnsafeSetGoID(id int64) { this.vid = id }

func bindQNoDebug(h uintptr) *QNoDebug {
	if h == 0 {
		return nil
	}
	return &QNoDebug{h: h}
}

func newownedQNoDebug(h uintptr) *QNoDebug {
	o := bindQNoDebug(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QNoDebug).Delete)
	}
	return o
}

func (this *QNoDebug) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QNoDebug) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQNoDebug(h unsafe.Pointer) *QNoDebug { return bindQNoDebug(uintptr(h)) }
func UnsafeNewQNoDebugOwned(h unsafe.Pointer) *QNoDebug { return newownedQNoDebug(uintptr(h)) }

func (this *QNoDebug) UnsafeRelease() { this.h = 0 }

func (this *QNoDebug) OperatorShiftLeft(arg1 *QTextStreamManipulator) *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_OperatorShiftLeft, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Space() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_space, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Nospace() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_nospace, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) MaybeSpace() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_maybeSpace, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Quote() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_quote, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Noquote() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_noquote, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) MaybeQuote() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_maybeQuote, this.h)
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Verbosity(arg1 int) *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_verbosity, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) MaybeQuoteWithChar(arg1 int8) *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QNoDebug_maybeQuote2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQNoDebug(__r)
}

func (this *QNoDebug) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QNoDebug_delete(this.h)
	}
	this.UnsafeRelease()
}
