package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringTokenizerBaseBase_new uintptr
)

func init() {
	symaddr(&addr_QStringTokenizerBaseBase_new, "QStringTokenizerBaseBase_new")
}

type QStringTokenizerBaseBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStringTokenizerBaseBase) UnsafeGoID() int64     { return this.vid }
func (this *QStringTokenizerBaseBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStringTokenizerBaseBase(h uintptr) *QStringTokenizerBaseBase {
	if h == 0 {
		return nil
	}
	return &QStringTokenizerBaseBase{h: h}
}

func newownedQStringTokenizerBaseBase(h uintptr) *QStringTokenizerBaseBase { return bindQStringTokenizerBaseBase(h) }

func (this *QStringTokenizerBaseBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringTokenizerBaseBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringTokenizerBaseBase(h unsafe.Pointer) *QStringTokenizerBaseBase { return bindQStringTokenizerBaseBase(uintptr(h)) }
func UnsafeNewQStringTokenizerBaseBaseOwned(h unsafe.Pointer) *QStringTokenizerBaseBase { return newownedQStringTokenizerBaseBase(uintptr(h)) }

func (this *QStringTokenizerBaseBase) UnsafeRelease() { this.h = 0 }

func NewQStringTokenizerBaseBase(arg1 *QStringTokenizerBaseBase) *QStringTokenizerBaseBase {
	__r, _, _ := purego.SyscallN(addr_QStringTokenizerBaseBase_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQStringTokenizerBaseBase(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QStringTokenizerBaseBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
