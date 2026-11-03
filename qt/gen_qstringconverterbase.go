package qt

import (
	"unsafe"
	"runtime"
)

type QStringConverterBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStringConverterBase) UnsafeGoID() int64     { return this.vid }
func (this *QStringConverterBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStringConverterBase(h uintptr) *QStringConverterBase {
	if h == 0 {
		return nil
	}
	return &QStringConverterBase{h: h}
}

func newownedQStringConverterBase(h uintptr) *QStringConverterBase { return bindQStringConverterBase(h) }

func (this *QStringConverterBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringConverterBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringConverterBase(h unsafe.Pointer) *QStringConverterBase { return bindQStringConverterBase(uintptr(h)) }
func UnsafeNewQStringConverterBaseOwned(h unsafe.Pointer) *QStringConverterBase { return newownedQStringConverterBase(uintptr(h)) }

func (this *QStringConverterBase) UnsafeRelease() { this.h = 0 }

func (this *QStringConverterBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
