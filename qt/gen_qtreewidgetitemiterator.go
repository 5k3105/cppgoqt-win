package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTreeWidgetItemIterator_new uintptr
	addr_QTreeWidgetItemIterator_new2 uintptr
	addr_QTreeWidgetItemIterator_new3 uintptr
	addr_QTreeWidgetItemIterator_new4 uintptr
	addr_QTreeWidgetItemIterator_new5 uintptr
	addr_QTreeWidgetItemIterator_OperatorAssign uintptr
	addr_QTreeWidgetItemIterator_OperatorPlusPlus uintptr
	addr_QTreeWidgetItemIterator_OperatorPlusPlus2 uintptr
	addr_QTreeWidgetItemIterator_OperatorPlusAssign uintptr
	addr_QTreeWidgetItemIterator_OperatorMinusMinus uintptr
	addr_QTreeWidgetItemIterator_OperatorMinusMinus2 uintptr
	addr_QTreeWidgetItemIterator_OperatorMinusAssign uintptr
	addr_QTreeWidgetItemIterator_OperatorMultiply uintptr
	pg_QTreeWidgetItemIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTreeWidgetItemIterator_new, "QTreeWidgetItemIterator_new")
	symaddr(&addr_QTreeWidgetItemIterator_new2, "QTreeWidgetItemIterator_new2")
	symaddr(&addr_QTreeWidgetItemIterator_new3, "QTreeWidgetItemIterator_new3")
	symaddr(&addr_QTreeWidgetItemIterator_new4, "QTreeWidgetItemIterator_new4")
	symaddr(&addr_QTreeWidgetItemIterator_new5, "QTreeWidgetItemIterator_new5")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorAssign, "QTreeWidgetItemIterator_OperatorAssign")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorPlusPlus, "QTreeWidgetItemIterator_OperatorPlusPlus")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorPlusPlus2, "QTreeWidgetItemIterator_OperatorPlusPlus2")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorPlusAssign, "QTreeWidgetItemIterator_OperatorPlusAssign")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorMinusMinus, "QTreeWidgetItemIterator_OperatorMinusMinus")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorMinusMinus2, "QTreeWidgetItemIterator_OperatorMinusMinus2")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorMinusAssign, "QTreeWidgetItemIterator_OperatorMinusAssign")
	symaddr(&addr_QTreeWidgetItemIterator_OperatorMultiply, "QTreeWidgetItemIterator_OperatorMultiply")
	reg(&pg_QTreeWidgetItemIterator_delete, "QTreeWidgetItemIterator_delete")
}

type QTreeWidgetItemIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTreeWidgetItemIterator) UnsafeGoID() int64     { return this.vid }
func (this *QTreeWidgetItemIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTreeWidgetItemIterator(h uintptr) *QTreeWidgetItemIterator {
	if h == 0 {
		return nil
	}
	return &QTreeWidgetItemIterator{h: h}
}

func newownedQTreeWidgetItemIterator(h uintptr) *QTreeWidgetItemIterator {
	o := bindQTreeWidgetItemIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTreeWidgetItemIterator).Delete)
	}
	return o
}

func (this *QTreeWidgetItemIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTreeWidgetItemIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTreeWidgetItemIterator(h unsafe.Pointer) *QTreeWidgetItemIterator { return bindQTreeWidgetItemIterator(uintptr(h)) }
func UnsafeNewQTreeWidgetItemIteratorOwned(h unsafe.Pointer) *QTreeWidgetItemIterator { return newownedQTreeWidgetItemIterator(uintptr(h)) }

func (this *QTreeWidgetItemIterator) UnsafeRelease() { this.h = 0 }

func NewQTreeWidgetItemIterator(it *QTreeWidgetItemIterator) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_new, uintptr(it.UnsafePointer()))
	__o := newownedQTreeWidgetItemIterator(__r)
	runtime.KeepAlive(it)
	return __o
}

func NewQTreeWidgetItemIterator2(widget *QTreeWidget) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_new2, uintptr(widget.UnsafePointer()))
	__o := newownedQTreeWidgetItemIterator(__r)
	runtime.KeepAlive(widget)
	return __o
}

func NewQTreeWidgetItemIterator3(item *QTreeWidgetItem) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_new3, uintptr(item.UnsafePointer()))
	__o := newownedQTreeWidgetItemIterator(__r)
	runtime.KeepAlive(item)
	return __o
}

func NewQTreeWidgetItemIterator4(widget *QTreeWidget, flags QTreeWidgetItemIterator__IteratorFlag) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_new4, uintptr(widget.UnsafePointer()), uintptr(flags))
	__o := newownedQTreeWidgetItemIterator(__r)
	runtime.KeepAlive(widget)
	return __o
}

func NewQTreeWidgetItemIterator5(item *QTreeWidgetItem, flags QTreeWidgetItemIterator__IteratorFlag) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_new5, uintptr(item.UnsafePointer()), uintptr(flags))
	__o := newownedQTreeWidgetItemIterator(__r)
	runtime.KeepAlive(item)
	return __o
}

func (this *QTreeWidgetItemIterator) OperatorAssign(it *QTreeWidgetItemIterator) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorAssign, this.h, uintptr(it.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(it)
	return bindQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorPlusPlus() *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorPlusPlus, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorPlusPlusWithInt(arg1 int) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorPlusPlus2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorPlusAssign(n int) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorPlusAssign, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorMinusMinus() *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorMinusMinus, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorMinusMinusWithInt(arg1 int) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorMinusMinus2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorMinusAssign(n int) *QTreeWidgetItemIterator {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorMinusAssign, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItemIterator(__r)
}

func (this *QTreeWidgetItemIterator) OperatorMultiply() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItemIterator_OperatorMultiply, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItemIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTreeWidgetItemIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
