package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCalendar__SystemId_new uintptr
	addr_QCalendar__SystemId_index uintptr
	addr_QCalendar__SystemId_isValid uintptr
	pg_QCalendar__SystemId_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCalendar__SystemId_new, "QCalendar__SystemId_new")
	symaddr(&addr_QCalendar__SystemId_index, "QCalendar__SystemId_index")
	symaddr(&addr_QCalendar__SystemId_isValid, "QCalendar__SystemId_isValid")
	reg(&pg_QCalendar__SystemId_delete, "QCalendar__SystemId_delete")
}

type QCalendar__SystemId struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCalendar__SystemId) UnsafeGoID() int64     { return this.vid }
func (this *QCalendar__SystemId) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCalendar__SystemId(h uintptr) *QCalendar__SystemId {
	if h == 0 {
		return nil
	}
	return &QCalendar__SystemId{h: h}
}

func newownedQCalendar__SystemId(h uintptr) *QCalendar__SystemId {
	o := bindQCalendar__SystemId(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCalendar__SystemId).Delete)
	}
	return o
}

func (this *QCalendar__SystemId) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCalendar__SystemId) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCalendar__SystemId(h unsafe.Pointer) *QCalendar__SystemId { return bindQCalendar__SystemId(uintptr(h)) }
func UnsafeNewQCalendar__SystemIdOwned(h unsafe.Pointer) *QCalendar__SystemId { return newownedQCalendar__SystemId(uintptr(h)) }

func (this *QCalendar__SystemId) UnsafeRelease() { this.h = 0 }

func NewQCalendar__SystemId() *QCalendar__SystemId {
	__r, _, _ := purego.SyscallN(addr_QCalendar__SystemId_new)
	__o := newownedQCalendar__SystemId(__r)
	return __o
}

func (this *QCalendar__SystemId) Index() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__SystemId_index, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QCalendar__SystemId) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__SystemId_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar__SystemId) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCalendar__SystemId_delete(this.h)
	}
	this.UnsafeRelease()
}
