package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRegularExpressionMatchIterator_new uintptr
	addr_QRegularExpressionMatchIterator_new2 uintptr
	addr_QRegularExpressionMatchIterator_OperatorAssign uintptr
	addr_QRegularExpressionMatchIterator_swap uintptr
	addr_QRegularExpressionMatchIterator_isValid uintptr
	addr_QRegularExpressionMatchIterator_hasNext uintptr
	addr_QRegularExpressionMatchIterator_next uintptr
	addr_QRegularExpressionMatchIterator_peekNext uintptr
	addr_QRegularExpressionMatchIterator_regularExpression uintptr
	addr_QRegularExpressionMatchIterator_matchType uintptr
	addr_QRegularExpressionMatchIterator_matchOptions uintptr
	pg_QRegularExpressionMatchIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRegularExpressionMatchIterator_new, "QRegularExpressionMatchIterator_new")
	symaddr(&addr_QRegularExpressionMatchIterator_new2, "QRegularExpressionMatchIterator_new2")
	symaddr(&addr_QRegularExpressionMatchIterator_OperatorAssign, "QRegularExpressionMatchIterator_OperatorAssign")
	symaddr(&addr_QRegularExpressionMatchIterator_swap, "QRegularExpressionMatchIterator_swap")
	symaddr(&addr_QRegularExpressionMatchIterator_isValid, "QRegularExpressionMatchIterator_isValid")
	symaddr(&addr_QRegularExpressionMatchIterator_hasNext, "QRegularExpressionMatchIterator_hasNext")
	symaddr(&addr_QRegularExpressionMatchIterator_next, "QRegularExpressionMatchIterator_next")
	symaddr(&addr_QRegularExpressionMatchIterator_peekNext, "QRegularExpressionMatchIterator_peekNext")
	symaddr(&addr_QRegularExpressionMatchIterator_regularExpression, "QRegularExpressionMatchIterator_regularExpression")
	symaddr(&addr_QRegularExpressionMatchIterator_matchType, "QRegularExpressionMatchIterator_matchType")
	symaddr(&addr_QRegularExpressionMatchIterator_matchOptions, "QRegularExpressionMatchIterator_matchOptions")
	reg(&pg_QRegularExpressionMatchIterator_delete, "QRegularExpressionMatchIterator_delete")
}

type QRegularExpressionMatchIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRegularExpressionMatchIterator) UnsafeGoID() int64     { return this.vid }
func (this *QRegularExpressionMatchIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRegularExpressionMatchIterator(h uintptr) *QRegularExpressionMatchIterator {
	if h == 0 {
		return nil
	}
	return &QRegularExpressionMatchIterator{h: h}
}

func newownedQRegularExpressionMatchIterator(h uintptr) *QRegularExpressionMatchIterator {
	o := bindQRegularExpressionMatchIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRegularExpressionMatchIterator).Delete)
	}
	return o
}

func (this *QRegularExpressionMatchIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRegularExpressionMatchIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRegularExpressionMatchIterator(h unsafe.Pointer) *QRegularExpressionMatchIterator { return bindQRegularExpressionMatchIterator(uintptr(h)) }
func UnsafeNewQRegularExpressionMatchIteratorOwned(h unsafe.Pointer) *QRegularExpressionMatchIterator { return newownedQRegularExpressionMatchIterator(uintptr(h)) }

func (this *QRegularExpressionMatchIterator) UnsafeRelease() { this.h = 0 }

func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_new)
	__o := newownedQRegularExpressionMatchIterator(__r)
	return __o
}

func NewQRegularExpressionMatchIterator2(iterator *QRegularExpressionMatchIterator) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_new2, uintptr(iterator.UnsafePointer()))
	__o := newownedQRegularExpressionMatchIterator(__r)
	runtime.KeepAlive(iterator)
	return __o
}

func (this *QRegularExpressionMatchIterator) OperatorAssign(iterator *QRegularExpressionMatchIterator) *QRegularExpressionMatchIterator {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_OperatorAssign, this.h, uintptr(iterator.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iterator)
	return bindQRegularExpressionMatchIterator(__r)
}

func (this *QRegularExpressionMatchIterator) Swap(other *QRegularExpressionMatchIterator) {
	purego.SyscallN(addr_QRegularExpressionMatchIterator_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QRegularExpressionMatchIterator) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatchIterator) HasNext() bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_hasNext, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRegularExpressionMatchIterator) Next() *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_next, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpressionMatchIterator) PeekNext() *QRegularExpressionMatch {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_peekNext, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpressionMatch(__r)
}

func (this *QRegularExpressionMatchIterator) RegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_regularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QRegularExpressionMatchIterator) MatchType() QRegularExpression__MatchType {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_matchType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRegularExpression__MatchType(__r)
}

func (this *QRegularExpressionMatchIterator) MatchOptions() QRegularExpression__MatchOption {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionMatchIterator_matchOptions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRegularExpression__MatchOption(__r)
}

func (this *QRegularExpressionMatchIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRegularExpressionMatchIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
