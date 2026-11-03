package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleSelectionInterface_selectedItemCount uintptr
	addr_QAccessibleSelectionInterface_selectedItems uintptr
	addr_QAccessibleSelectionInterface_selectedItem uintptr
	addr_QAccessibleSelectionInterface_isSelected uintptr
	addr_QAccessibleSelectionInterface_select uintptr
	addr_QAccessibleSelectionInterface_unselect uintptr
	addr_QAccessibleSelectionInterface_selectAll uintptr
	addr_QAccessibleSelectionInterface_clear uintptr
	addr_QAccessibleSelectionInterface_OperatorAssign uintptr
	addr_QAccessibleSelectionInterface_selectedItem_default uintptr
	addr_QAccessibleSelectionInterface_isSelected_default uintptr
	pg_QAccessibleSelectionInterface__set_override func(uintptr, int32)
	pg_QAccessibleSelectionInterface__Go_new func(int64) uintptr
	pg_QAccessibleSelectionInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleSelectionInterface_selectedItemCount, "QAccessibleSelectionInterface_selectedItemCount")
	symaddr(&addr_QAccessibleSelectionInterface_selectedItems, "QAccessibleSelectionInterface_selectedItems")
	symaddr(&addr_QAccessibleSelectionInterface_selectedItem, "QAccessibleSelectionInterface_selectedItem")
	symaddr(&addr_QAccessibleSelectionInterface_isSelected, "QAccessibleSelectionInterface_isSelected")
	symaddr(&addr_QAccessibleSelectionInterface_select, "QAccessibleSelectionInterface_select")
	symaddr(&addr_QAccessibleSelectionInterface_unselect, "QAccessibleSelectionInterface_unselect")
	symaddr(&addr_QAccessibleSelectionInterface_selectAll, "QAccessibleSelectionInterface_selectAll")
	symaddr(&addr_QAccessibleSelectionInterface_clear, "QAccessibleSelectionInterface_clear")
	symaddr(&addr_QAccessibleSelectionInterface_OperatorAssign, "QAccessibleSelectionInterface_OperatorAssign")
	symaddr(&addr_QAccessibleSelectionInterface_selectedItem_default, "QAccessibleSelectionInterface_selectedItem_default")
	symaddr(&addr_QAccessibleSelectionInterface_isSelected_default, "QAccessibleSelectionInterface_isSelected_default")
	reg(&pg_QAccessibleSelectionInterface__set_override, "QAccessibleSelectionInterface__set_override")
	reg(&pg_QAccessibleSelectionInterface__Go_new, "QAccessibleSelectionInterface__Go_new")
	reg(&pg_QAccessibleSelectionInterface_delete, "QAccessibleSelectionInterface_delete")
}

type QAccessibleSelectionInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleSelectionInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleSelectionInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleSelectionInterface(h uintptr) *QAccessibleSelectionInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleSelectionInterface{h: h}
}

func newownedQAccessibleSelectionInterface(h uintptr) *QAccessibleSelectionInterface {
	o := bindQAccessibleSelectionInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleSelectionInterface).Delete)
	}
	return o
}

func (this *QAccessibleSelectionInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleSelectionInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleSelectionInterface(h unsafe.Pointer) *QAccessibleSelectionInterface { return bindQAccessibleSelectionInterface(uintptr(h)) }
func UnsafeNewQAccessibleSelectionInterfaceOwned(h unsafe.Pointer) *QAccessibleSelectionInterface { return newownedQAccessibleSelectionInterface(uintptr(h)) }

func (this *QAccessibleSelectionInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleSelectionInterface) SelectedItemCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_selectedItemCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleSelectionInterface) SelectedItems() []*QAccessibleInterface {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_selectedItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAccessibleInterface, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAccessibleInterface(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleSelectionInterface) SelectedItem(selectionIndex int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_selectedItem, this.h, uintptr(selectionIndex))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleSelectionInterface) IsSelected(childItem *QAccessibleInterface) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_isSelected, this.h, uintptr(childItem.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childItem)
	return __r != 0
}

func (this *QAccessibleSelectionInterface) Select(childItem *QAccessibleInterface) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_select, this.h, uintptr(childItem.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childItem)
	return __r != 0
}

func (this *QAccessibleSelectionInterface) Unselect(childItem *QAccessibleInterface) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_unselect, this.h, uintptr(childItem.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childItem)
	return __r != 0
}

func (this *QAccessibleSelectionInterface) SelectAll() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_selectAll, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleSelectionInterface) Clear() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_clear, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleSelectionInterface) OperatorAssign(arg1 *QAccessibleSelectionInterface) *QAccessibleSelectionInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleSelectionInterface(__r)
}

func (this *QAccessibleSelectionInterface) SelectedItemDefault(selectionIndex int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_selectedItem_default, this.h, uintptr(selectionIndex))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleSelectionInterface) IsSelectedDefault(childItem *QAccessibleInterface) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleSelectionInterface_isSelected_default, this.h, uintptr(childItem.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(childItem)
	return __r != 0
}

func QAccessibleSelectionInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 1835560489624661370:
		return 0
	case -5807941332178582040:
		return 1
	case -4794627110275867635:
		return 2
	case 5098092392905213826:
		return 3
	case 8758113809152574300:
		return 4
	case -778567940572639070:
		return 5
	case -864810461788103157:
		return 6
	case -3350427747592150876:
		return 7
	}
	return -1
}

func QAccessibleSelectionInterface_enable(h uintptr, idx int32) { pg_QAccessibleSelectionInterface__set_override(h, idx) }

func NewQAccessibleSelectionInterface() *QAccessibleSelectionInterface {
	id := nextVirtID()
	o := newownedQAccessibleSelectionInterface(pg_QAccessibleSelectionInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleSelectionInterface_virtIdx, QAccessibleSelectionInterface_enable)
	return o
}

func (this *QAccessibleSelectionInterface) OnSelectedItemCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), 1835560489624661370, func() int32 { return int32(fn()) }) // slot "selectedItemCount"
}

func (this *QAccessibleSelectionInterface) OnSelectedItems(fn func() []*QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -5807941332178582040, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "selectedItems"
}

func (this *QAccessibleSelectionInterface) OnSelect(fn func(*QAccessibleInterface) bool) {
	registerVirt(this.UnsafeGoID(), -4794627110275867635, func(childItem uintptr) int32 { return b2i(fn(UnsafeNewQAccessibleInterface(unsafe.Pointer(childItem)))) }) // slot "select"
}

func (this *QAccessibleSelectionInterface) OnUnselect(fn func(*QAccessibleInterface) bool) {
	registerVirt(this.UnsafeGoID(), 5098092392905213826, func(childItem uintptr) int32 { return b2i(fn(UnsafeNewQAccessibleInterface(unsafe.Pointer(childItem)))) }) // slot "unselect"
}

func (this *QAccessibleSelectionInterface) OnSelectAll(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8758113809152574300, func() int32 { return b2i(fn()) }) // slot "selectAll"
}

func (this *QAccessibleSelectionInterface) OnClear(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() int32 { return b2i(fn()) }) // slot "clear"
}

func (this *QAccessibleSelectionInterface) OnSelectedItem(fn func(int) *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -864810461788103157, func(selectionIndex int32) uintptr { return uintptr(fn(int(selectionIndex)).UnsafePointer()) }) // slot "selectedItem"
}

func (this *QAccessibleSelectionInterface) OnIsSelected(fn func(*QAccessibleInterface) bool) {
	registerVirt(this.UnsafeGoID(), -3350427747592150876, func(childItem uintptr) int32 { return b2i(fn(UnsafeNewQAccessibleInterface(unsafe.Pointer(childItem)))) }) // slot "isSelected"
}

func (this *QAccessibleSelectionInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleSelectionInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
