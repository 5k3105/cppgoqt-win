package qt

import (
	"unsafe"
	"runtime"
)

var (
	pg_QAbstractNativeEventFilter__set_override func(uintptr, int32)
	pg_QAbstractNativeEventFilter__Go_new func(int64) uintptr
	pg_QAbstractNativeEventFilter_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractNativeEventFilter__set_override, "QAbstractNativeEventFilter__set_override")
	reg(&pg_QAbstractNativeEventFilter__Go_new, "QAbstractNativeEventFilter__Go_new")
	reg(&pg_QAbstractNativeEventFilter_delete, "QAbstractNativeEventFilter_delete")
}

type QAbstractNativeEventFilter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractNativeEventFilter) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractNativeEventFilter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractNativeEventFilter(h uintptr) *QAbstractNativeEventFilter {
	if h == 0 {
		return nil
	}
	return &QAbstractNativeEventFilter{h: h}
}

func newownedQAbstractNativeEventFilter(h uintptr) *QAbstractNativeEventFilter {
	o := bindQAbstractNativeEventFilter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractNativeEventFilter).Delete)
	}
	return o
}

func (this *QAbstractNativeEventFilter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractNativeEventFilter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractNativeEventFilter(h unsafe.Pointer) *QAbstractNativeEventFilter { return bindQAbstractNativeEventFilter(uintptr(h)) }
func UnsafeNewQAbstractNativeEventFilterOwned(h unsafe.Pointer) *QAbstractNativeEventFilter { return newownedQAbstractNativeEventFilter(uintptr(h)) }

func (this *QAbstractNativeEventFilter) UnsafeRelease() { this.h = 0 }

func QAbstractNativeEventFilter_virtIdx(slot int64) int32 {
	switch slot {
	case -3445417634659622720:
		return 0
	}
	return -1
}

func QAbstractNativeEventFilter_enable(h uintptr, idx int32) { pg_QAbstractNativeEventFilter__set_override(h, idx) }

func NewQAbstractNativeEventFilter() *QAbstractNativeEventFilter {
	id := nextVirtID()
	o := newownedQAbstractNativeEventFilter(pg_QAbstractNativeEventFilter__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractNativeEventFilter_virtIdx, QAbstractNativeEventFilter_enable)
	return o
}

func (this *QAbstractNativeEventFilter) OnNativeEventFilter(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -3445417634659622720, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEventFilter"
}

func (this *QAbstractNativeEventFilter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractNativeEventFilter_delete(this.h)
	}
	this.UnsafeRelease()
}
