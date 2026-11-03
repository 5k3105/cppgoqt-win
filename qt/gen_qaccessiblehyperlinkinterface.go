package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleHyperlinkInterface_anchor uintptr
	addr_QAccessibleHyperlinkInterface_anchorTarget uintptr
	addr_QAccessibleHyperlinkInterface_startIndex uintptr
	addr_QAccessibleHyperlinkInterface_endIndex uintptr
	addr_QAccessibleHyperlinkInterface_isValid uintptr
	addr_QAccessibleHyperlinkInterface_OperatorAssign uintptr
	pg_QAccessibleHyperlinkInterface__set_override func(uintptr, int32)
	pg_QAccessibleHyperlinkInterface__Go_new func(int64) uintptr
	pg_QAccessibleHyperlinkInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleHyperlinkInterface_anchor, "QAccessibleHyperlinkInterface_anchor")
	symaddr(&addr_QAccessibleHyperlinkInterface_anchorTarget, "QAccessibleHyperlinkInterface_anchorTarget")
	symaddr(&addr_QAccessibleHyperlinkInterface_startIndex, "QAccessibleHyperlinkInterface_startIndex")
	symaddr(&addr_QAccessibleHyperlinkInterface_endIndex, "QAccessibleHyperlinkInterface_endIndex")
	symaddr(&addr_QAccessibleHyperlinkInterface_isValid, "QAccessibleHyperlinkInterface_isValid")
	symaddr(&addr_QAccessibleHyperlinkInterface_OperatorAssign, "QAccessibleHyperlinkInterface_OperatorAssign")
	reg(&pg_QAccessibleHyperlinkInterface__set_override, "QAccessibleHyperlinkInterface__set_override")
	reg(&pg_QAccessibleHyperlinkInterface__Go_new, "QAccessibleHyperlinkInterface__Go_new")
	reg(&pg_QAccessibleHyperlinkInterface_delete, "QAccessibleHyperlinkInterface_delete")
}

type QAccessibleHyperlinkInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleHyperlinkInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleHyperlinkInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleHyperlinkInterface(h uintptr) *QAccessibleHyperlinkInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleHyperlinkInterface{h: h}
}

func newownedQAccessibleHyperlinkInterface(h uintptr) *QAccessibleHyperlinkInterface {
	o := bindQAccessibleHyperlinkInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleHyperlinkInterface).Delete)
	}
	return o
}

func (this *QAccessibleHyperlinkInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleHyperlinkInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleHyperlinkInterface(h unsafe.Pointer) *QAccessibleHyperlinkInterface { return bindQAccessibleHyperlinkInterface(uintptr(h)) }
func UnsafeNewQAccessibleHyperlinkInterfaceOwned(h unsafe.Pointer) *QAccessibleHyperlinkInterface { return newownedQAccessibleHyperlinkInterface(uintptr(h)) }

func (this *QAccessibleHyperlinkInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleHyperlinkInterface) Anchor() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_anchor, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleHyperlinkInterface) AnchorTarget() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_anchorTarget, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleHyperlinkInterface) StartIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_startIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleHyperlinkInterface) EndIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_endIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleHyperlinkInterface) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleHyperlinkInterface) OperatorAssign(arg1 *QAccessibleHyperlinkInterface) *QAccessibleHyperlinkInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleHyperlinkInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleHyperlinkInterface(__r)
}

func QAccessibleHyperlinkInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 2999672799903939988:
		return 0
	case -8101286066007612963:
		return 1
	case -1674626578571099579:
		return 2
	case 4812684365611001022:
		return 3
	case 7231157875507698241:
		return 4
	}
	return -1
}

func QAccessibleHyperlinkInterface_enable(h uintptr, idx int32) { pg_QAccessibleHyperlinkInterface__set_override(h, idx) }

func NewQAccessibleHyperlinkInterface() *QAccessibleHyperlinkInterface {
	id := nextVirtID()
	o := newownedQAccessibleHyperlinkInterface(pg_QAccessibleHyperlinkInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleHyperlinkInterface_virtIdx, QAccessibleHyperlinkInterface_enable)
	return o
}

func (this *QAccessibleHyperlinkInterface) OnAnchor(fn func() string) {
	registerVirt(this.UnsafeGoID(), 2999672799903939988, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "anchor"
}

func (this *QAccessibleHyperlinkInterface) OnAnchorTarget(fn func() string) {
	registerVirt(this.UnsafeGoID(), -8101286066007612963, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "anchorTarget"
}

func (this *QAccessibleHyperlinkInterface) OnStartIndex(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1674626578571099579, func() int32 { return int32(fn()) }) // slot "startIndex"
}

func (this *QAccessibleHyperlinkInterface) OnEndIndex(fn func() int) {
	registerVirt(this.UnsafeGoID(), 4812684365611001022, func() int32 { return int32(fn()) }) // slot "endIndex"
}

func (this *QAccessibleHyperlinkInterface) OnIsValid(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 7231157875507698241, func() int32 { return b2i(fn()) }) // slot "isValid"
}

func (this *QAccessibleHyperlinkInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleHyperlinkInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
