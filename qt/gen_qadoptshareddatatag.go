package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAdoptSharedDataTag_new uintptr
	pg_QAdoptSharedDataTag_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAdoptSharedDataTag_new, "QAdoptSharedDataTag_new")
	reg(&pg_QAdoptSharedDataTag_delete, "QAdoptSharedDataTag_delete")
}

type QAdoptSharedDataTag struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAdoptSharedDataTag) UnsafeGoID() int64     { return this.vid }
func (this *QAdoptSharedDataTag) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAdoptSharedDataTag(h uintptr) *QAdoptSharedDataTag {
	if h == 0 {
		return nil
	}
	return &QAdoptSharedDataTag{h: h}
}

func newownedQAdoptSharedDataTag(h uintptr) *QAdoptSharedDataTag {
	o := bindQAdoptSharedDataTag(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAdoptSharedDataTag).Delete)
	}
	return o
}

func (this *QAdoptSharedDataTag) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAdoptSharedDataTag) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAdoptSharedDataTag(h unsafe.Pointer) *QAdoptSharedDataTag { return bindQAdoptSharedDataTag(uintptr(h)) }
func UnsafeNewQAdoptSharedDataTagOwned(h unsafe.Pointer) *QAdoptSharedDataTag { return newownedQAdoptSharedDataTag(uintptr(h)) }

func (this *QAdoptSharedDataTag) UnsafeRelease() { this.h = 0 }

func NewQAdoptSharedDataTag() *QAdoptSharedDataTag {
	__r, _, _ := purego.SyscallN(addr_QAdoptSharedDataTag_new)
	__o := newownedQAdoptSharedDataTag(__r)
	return __o
}

func (this *QAdoptSharedDataTag) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAdoptSharedDataTag_delete(this.h)
	}
	this.UnsafeRelease()
}
