package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextOption__Tab_new uintptr
	pg_QTextOption__Tab_new2 func(float64, int32) uintptr
	pg_QTextOption__Tab_new3 func(float64, int32, uintptr) uintptr
	addr_QTextOption__Tab_OperatorEqual uintptr
	addr_QTextOption__Tab_OperatorNotEqual uintptr
	pg_QTextOption__Tab_position_get func(uintptr) float64
	pg_QTextOption__Tab_position_set func(uintptr, float64)
	addr_QTextOption__Tab_type_get uintptr
	addr_QTextOption__Tab_type_set uintptr
	addr_QTextOption__Tab_delimiter_get uintptr
	addr_QTextOption__Tab_delimiter_set uintptr
	pg_QTextOption__Tab_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextOption__Tab_new, "QTextOption__Tab_new")
	reg(&pg_QTextOption__Tab_new2, "QTextOption__Tab_new2")
	reg(&pg_QTextOption__Tab_new3, "QTextOption__Tab_new3")
	symaddr(&addr_QTextOption__Tab_OperatorEqual, "QTextOption__Tab_OperatorEqual")
	symaddr(&addr_QTextOption__Tab_OperatorNotEqual, "QTextOption__Tab_OperatorNotEqual")
	reg(&pg_QTextOption__Tab_position_get, "QTextOption__Tab_position_get")
	reg(&pg_QTextOption__Tab_position_set, "QTextOption__Tab_position_set")
	symaddr(&addr_QTextOption__Tab_type_get, "QTextOption__Tab_type_get")
	symaddr(&addr_QTextOption__Tab_type_set, "QTextOption__Tab_type_set")
	symaddr(&addr_QTextOption__Tab_delimiter_get, "QTextOption__Tab_delimiter_get")
	symaddr(&addr_QTextOption__Tab_delimiter_set, "QTextOption__Tab_delimiter_set")
	reg(&pg_QTextOption__Tab_delete, "QTextOption__Tab_delete")
}

type QTextOption__Tab struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextOption__Tab) UnsafeGoID() int64     { return this.vid }
func (this *QTextOption__Tab) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextOption__Tab(h uintptr) *QTextOption__Tab {
	if h == 0 {
		return nil
	}
	return &QTextOption__Tab{h: h}
}

func newownedQTextOption__Tab(h uintptr) *QTextOption__Tab {
	o := bindQTextOption__Tab(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextOption__Tab).Delete)
	}
	return o
}

func (this *QTextOption__Tab) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextOption__Tab) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextOption__Tab(h unsafe.Pointer) *QTextOption__Tab { return bindQTextOption__Tab(uintptr(h)) }
func UnsafeNewQTextOption__TabOwned(h unsafe.Pointer) *QTextOption__Tab { return newownedQTextOption__Tab(uintptr(h)) }

func (this *QTextOption__Tab) UnsafeRelease() { this.h = 0 }

func NewQTextOption__Tab() *QTextOption__Tab {
	__r, _, _ := purego.SyscallN(addr_QTextOption__Tab_new)
	__o := newownedQTextOption__Tab(__r)
	return __o
}

func NewQTextOption__Tab2(pos float64, tabType QTextOption__TabType) *QTextOption__Tab {
	__o := newownedQTextOption__Tab(pg_QTextOption__Tab_new2(pos, int32(tabType)))
	return __o
}

func NewQTextOption__Tab3(pos float64, tabType QTextOption__TabType, delim *QChar) *QTextOption__Tab {
	__o := newownedQTextOption__Tab(pg_QTextOption__Tab_new3(pos, int32(tabType), uintptr(delim.UnsafePointer())))
	runtime.KeepAlive(delim)
	return __o
}

func (this *QTextOption__Tab) OperatorEqual(other *QTextOption__Tab) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextOption__Tab_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextOption__Tab) OperatorNotEqual(other *QTextOption__Tab) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextOption__Tab_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextOption__Tab) Position() float64 {
	__r := pg_QTextOption__Tab_position_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextOption__Tab) SetPosition(position float64) {
	pg_QTextOption__Tab_position_set(this.h, position)
	runtime.KeepAlive(this)
}

func (this *QTextOption__Tab) Type() QTextOption__TabType {
	__r0, _, _ := purego.SyscallN(addr_QTextOption__Tab_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextOption__TabType(__r)
}

func (this *QTextOption__Tab) SetType(type_ QTextOption__TabType) {
	purego.SyscallN(addr_QTextOption__Tab_type_set, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QTextOption__Tab) Delimiter() *QChar {
	__r, _, _ := purego.SyscallN(addr_QTextOption__Tab_delimiter_get, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QTextOption__Tab) SetDelimiter(delimiter *QChar) {
	purego.SyscallN(addr_QTextOption__Tab_delimiter_set, this.h, uintptr(delimiter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delimiter)
}

func (this *QTextOption__Tab) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextOption__Tab_delete(this.h)
	}
	this.UnsafeRelease()
}
