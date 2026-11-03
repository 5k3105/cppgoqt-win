package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTileRules_new uintptr
	addr_QTileRules_new2 uintptr
	addr_QTileRules_new3 uintptr
	addr_QTileRules_new4 uintptr
	addr_QTileRules_horizontal_get uintptr
	addr_QTileRules_horizontal_set uintptr
	addr_QTileRules_vertical_get uintptr
	addr_QTileRules_vertical_set uintptr
	pg_QTileRules_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTileRules_new, "QTileRules_new")
	symaddr(&addr_QTileRules_new2, "QTileRules_new2")
	symaddr(&addr_QTileRules_new3, "QTileRules_new3")
	symaddr(&addr_QTileRules_new4, "QTileRules_new4")
	symaddr(&addr_QTileRules_horizontal_get, "QTileRules_horizontal_get")
	symaddr(&addr_QTileRules_horizontal_set, "QTileRules_horizontal_set")
	symaddr(&addr_QTileRules_vertical_get, "QTileRules_vertical_get")
	symaddr(&addr_QTileRules_vertical_set, "QTileRules_vertical_set")
	reg(&pg_QTileRules_delete, "QTileRules_delete")
}

type QTileRules struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTileRules) UnsafeGoID() int64     { return this.vid }
func (this *QTileRules) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTileRules(h uintptr) *QTileRules {
	if h == 0 {
		return nil
	}
	return &QTileRules{h: h}
}

func newownedQTileRules(h uintptr) *QTileRules {
	o := bindQTileRules(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTileRules).Delete)
	}
	return o
}

func (this *QTileRules) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTileRules) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTileRules(h unsafe.Pointer) *QTileRules { return bindQTileRules(uintptr(h)) }
func UnsafeNewQTileRulesOwned(h unsafe.Pointer) *QTileRules { return newownedQTileRules(uintptr(h)) }

func (this *QTileRules) UnsafeRelease() { this.h = 0 }

func NewQTileRules(horizontalRule TileRule, verticalRule TileRule) *QTileRules {
	__r, _, _ := purego.SyscallN(addr_QTileRules_new, uintptr(horizontalRule), uintptr(verticalRule))
	__o := newownedQTileRules(__r)
	return __o
}

func NewQTileRules2() *QTileRules {
	__r, _, _ := purego.SyscallN(addr_QTileRules_new2)
	__o := newownedQTileRules(__r)
	return __o
}

func NewQTileRules3(arg1 *QTileRules) *QTileRules {
	__r, _, _ := purego.SyscallN(addr_QTileRules_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQTileRules(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQTileRules4(rule TileRule) *QTileRules {
	__r, _, _ := purego.SyscallN(addr_QTileRules_new4, uintptr(rule))
	__o := newownedQTileRules(__r)
	return __o
}

func (this *QTileRules) Horizontal() TileRule {
	__r0, _, _ := purego.SyscallN(addr_QTileRules_horizontal_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TileRule(__r)
}

func (this *QTileRules) SetHorizontal(horizontal TileRule) {
	purego.SyscallN(addr_QTileRules_horizontal_set, this.h, uintptr(horizontal))
	runtime.KeepAlive(this)
}

func (this *QTileRules) Vertical() TileRule {
	__r0, _, _ := purego.SyscallN(addr_QTileRules_vertical_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TileRule(__r)
}

func (this *QTileRules) SetVertical(vertical TileRule) {
	purego.SyscallN(addr_QTileRules_vertical_set, this.h, uintptr(vertical))
	runtime.KeepAlive(this)
}

func (this *QTileRules) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTileRules_delete(this.h)
	}
	this.UnsafeRelease()
}
