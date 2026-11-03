package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRunnable_run uintptr
	addr_QRunnable_autoDelete uintptr
	addr_QRunnable_setAutoDelete uintptr
	pg_QRunnable__set_override func(uintptr, int32)
	pg_QRunnable__Go_new func(int64) uintptr
	pg_QRunnable_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRunnable_run, "QRunnable_run")
	symaddr(&addr_QRunnable_autoDelete, "QRunnable_autoDelete")
	symaddr(&addr_QRunnable_setAutoDelete, "QRunnable_setAutoDelete")
	reg(&pg_QRunnable__set_override, "QRunnable__set_override")
	reg(&pg_QRunnable__Go_new, "QRunnable__Go_new")
	reg(&pg_QRunnable_delete, "QRunnable_delete")
}

type QRunnable struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRunnable) UnsafeGoID() int64     { return this.vid }
func (this *QRunnable) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRunnable(h uintptr) *QRunnable {
	if h == 0 {
		return nil
	}
	return &QRunnable{h: h}
}

func newownedQRunnable(h uintptr) *QRunnable {
	o := bindQRunnable(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRunnable).Delete)
	}
	return o
}

func (this *QRunnable) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRunnable) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRunnable(h unsafe.Pointer) *QRunnable { return bindQRunnable(uintptr(h)) }
func UnsafeNewQRunnableOwned(h unsafe.Pointer) *QRunnable { return newownedQRunnable(uintptr(h)) }

func (this *QRunnable) UnsafeRelease() { this.h = 0 }

func (this *QRunnable) Run() {
	purego.SyscallN(addr_QRunnable_run, this.h)
	runtime.KeepAlive(this)
}

func (this *QRunnable) AutoDelete() bool {
	__r0, _, _ := purego.SyscallN(addr_QRunnable_autoDelete, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRunnable) SetAutoDelete(autoDelete bool) {
	purego.SyscallN(addr_QRunnable_setAutoDelete, this.h, uintptr(b2i(autoDelete)))
	runtime.KeepAlive(this)
}

func QRunnable_virtIdx(slot int64) int32 {
	switch slot {
	case -8524566317707732406:
		return 0
	}
	return -1
}

func QRunnable_enable(h uintptr, idx int32) { pg_QRunnable__set_override(h, idx) }

func NewQRunnable() *QRunnable {
	id := nextVirtID()
	o := newownedQRunnable(pg_QRunnable__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QRunnable_virtIdx, QRunnable_enable)
	return o
}

func (this *QRunnable) OnRun(fn func()) {
	registerVirt(this.UnsafeGoID(), -8524566317707732406, func() { fn() }) // slot "run"
}

func (this *QRunnable) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRunnable_delete(this.h)
	}
	this.UnsafeRelease()
}
