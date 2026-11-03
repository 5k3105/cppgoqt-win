package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBackingStore_new uintptr
	addr_QBackingStore_window uintptr
	addr_QBackingStore_paintDevice uintptr
	addr_QBackingStore_flush uintptr
	addr_QBackingStore_resize uintptr
	addr_QBackingStore_size uintptr
	addr_QBackingStore_scroll uintptr
	addr_QBackingStore_beginPaint uintptr
	addr_QBackingStore_endPaint uintptr
	addr_QBackingStore_setStaticContents uintptr
	addr_QBackingStore_staticContents uintptr
	addr_QBackingStore_hasStaticContents uintptr
	addr_QBackingStore_flush2 uintptr
	addr_QBackingStore_flush3 uintptr
	pg_QBackingStore_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBackingStore_new, "QBackingStore_new")
	symaddr(&addr_QBackingStore_window, "QBackingStore_window")
	symaddr(&addr_QBackingStore_paintDevice, "QBackingStore_paintDevice")
	symaddr(&addr_QBackingStore_flush, "QBackingStore_flush")
	symaddr(&addr_QBackingStore_resize, "QBackingStore_resize")
	symaddr(&addr_QBackingStore_size, "QBackingStore_size")
	symaddr(&addr_QBackingStore_scroll, "QBackingStore_scroll")
	symaddr(&addr_QBackingStore_beginPaint, "QBackingStore_beginPaint")
	symaddr(&addr_QBackingStore_endPaint, "QBackingStore_endPaint")
	symaddr(&addr_QBackingStore_setStaticContents, "QBackingStore_setStaticContents")
	symaddr(&addr_QBackingStore_staticContents, "QBackingStore_staticContents")
	symaddr(&addr_QBackingStore_hasStaticContents, "QBackingStore_hasStaticContents")
	symaddr(&addr_QBackingStore_flush2, "QBackingStore_flush2")
	symaddr(&addr_QBackingStore_flush3, "QBackingStore_flush3")
	reg(&pg_QBackingStore_delete, "QBackingStore_delete")
}

type QBackingStore struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBackingStore) UnsafeGoID() int64     { return this.vid }
func (this *QBackingStore) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBackingStore(h uintptr) *QBackingStore {
	if h == 0 {
		return nil
	}
	return &QBackingStore{h: h}
}

func newownedQBackingStore(h uintptr) *QBackingStore {
	o := bindQBackingStore(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBackingStore).Delete)
	}
	return o
}

func (this *QBackingStore) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBackingStore) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBackingStore(h unsafe.Pointer) *QBackingStore { return bindQBackingStore(uintptr(h)) }
func UnsafeNewQBackingStoreOwned(h unsafe.Pointer) *QBackingStore { return newownedQBackingStore(uintptr(h)) }

func (this *QBackingStore) UnsafeRelease() { this.h = 0 }

func NewQBackingStore(window *QWindow) *QBackingStore {
	__r, _, _ := purego.SyscallN(addr_QBackingStore_new, uintptr(window.UnsafePointer()))
	__o := newownedQBackingStore(__r)
	runtime.KeepAlive(window)
	return __o
}

func (this *QBackingStore) Window() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QBackingStore_window, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QBackingStore) PaintDevice() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QBackingStore_paintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QBackingStore) Flush(region *QRegion) {
	purego.SyscallN(addr_QBackingStore_flush, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QBackingStore) Resize(size *QSize) {
	purego.SyscallN(addr_QBackingStore_resize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QBackingStore) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QBackingStore_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QBackingStore) Scroll(area *QRegion, dx int, dy int) bool {
	__r0, _, _ := purego.SyscallN(addr_QBackingStore_scroll, this.h, uintptr(area.UnsafePointer()), uintptr(dx), uintptr(dy))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(area)
	return __r != 0
}

func (this *QBackingStore) BeginPaint(arg1 *QRegion) {
	purego.SyscallN(addr_QBackingStore_beginPaint, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBackingStore) EndPaint() {
	purego.SyscallN(addr_QBackingStore_endPaint, this.h)
	runtime.KeepAlive(this)
}

func (this *QBackingStore) SetStaticContents(region *QRegion) {
	purego.SyscallN(addr_QBackingStore_setStaticContents, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QBackingStore) StaticContents() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QBackingStore_staticContents, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QBackingStore) HasStaticContents() bool {
	__r0, _, _ := purego.SyscallN(addr_QBackingStore_hasStaticContents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBackingStore) Flush2(region *QRegion, window *QWindow) {
	purego.SyscallN(addr_QBackingStore_flush2, this.h, uintptr(region.UnsafePointer()), uintptr(window.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
	runtime.KeepAlive(window)
}

func (this *QBackingStore) Flush3(region *QRegion, window *QWindow, offset *QPoint) {
	purego.SyscallN(addr_QBackingStore_flush3, this.h, uintptr(region.UnsafePointer()), uintptr(window.UnsafePointer()), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
	runtime.KeepAlive(window)
	runtime.KeepAlive(offset)
}

func (this *QBackingStore) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBackingStore_delete(this.h)
	}
	this.UnsafeRelease()
}
