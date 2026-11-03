package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIconEngine_paint uintptr
	addr_QIconEngine_actualSize uintptr
	addr_QIconEngine_pixmap uintptr
	addr_QIconEngine_addPixmap uintptr
	addr_QIconEngine_addFile uintptr
	addr_QIconEngine_key uintptr
	addr_QIconEngine_clone uintptr
	addr_QIconEngine_read uintptr
	addr_QIconEngine_write uintptr
	addr_QIconEngine_availableSizes uintptr
	addr_QIconEngine_iconName uintptr
	addr_QIconEngine_isNull uintptr
	pg_QIconEngine_scaledPixmap func(uintptr, uintptr, int32, int32, float64) uintptr
	addr_QIconEngine_actualSize_default uintptr
	addr_QIconEngine_pixmap_default uintptr
	addr_QIconEngine_addPixmap_default uintptr
	addr_QIconEngine_addFile_default uintptr
	addr_QIconEngine_key_default uintptr
	addr_QIconEngine_read_default uintptr
	addr_QIconEngine_write_default uintptr
	addr_QIconEngine_availableSizes_default uintptr
	addr_QIconEngine_iconName_default uintptr
	addr_QIconEngine_isNull_default uintptr
	pg_QIconEngine_scaledPixmap_default func(uintptr, uintptr, int32, int32, float64) uintptr
	pg_QIconEngine__set_override func(uintptr, int32)
	pg_QIconEngine__Go_new func(int64) uintptr
	pg_QIconEngine_delete func(uintptr)
)

func init() {
	symaddr(&addr_QIconEngine_paint, "QIconEngine_paint")
	symaddr(&addr_QIconEngine_actualSize, "QIconEngine_actualSize")
	symaddr(&addr_QIconEngine_pixmap, "QIconEngine_pixmap")
	symaddr(&addr_QIconEngine_addPixmap, "QIconEngine_addPixmap")
	symaddr(&addr_QIconEngine_addFile, "QIconEngine_addFile")
	symaddr(&addr_QIconEngine_key, "QIconEngine_key")
	symaddr(&addr_QIconEngine_clone, "QIconEngine_clone")
	symaddr(&addr_QIconEngine_read, "QIconEngine_read")
	symaddr(&addr_QIconEngine_write, "QIconEngine_write")
	symaddr(&addr_QIconEngine_availableSizes, "QIconEngine_availableSizes")
	symaddr(&addr_QIconEngine_iconName, "QIconEngine_iconName")
	symaddr(&addr_QIconEngine_isNull, "QIconEngine_isNull")
	reg(&pg_QIconEngine_scaledPixmap, "QIconEngine_scaledPixmap")
	symaddr(&addr_QIconEngine_actualSize_default, "QIconEngine_actualSize_default")
	symaddr(&addr_QIconEngine_pixmap_default, "QIconEngine_pixmap_default")
	symaddr(&addr_QIconEngine_addPixmap_default, "QIconEngine_addPixmap_default")
	symaddr(&addr_QIconEngine_addFile_default, "QIconEngine_addFile_default")
	symaddr(&addr_QIconEngine_key_default, "QIconEngine_key_default")
	symaddr(&addr_QIconEngine_read_default, "QIconEngine_read_default")
	symaddr(&addr_QIconEngine_write_default, "QIconEngine_write_default")
	symaddr(&addr_QIconEngine_availableSizes_default, "QIconEngine_availableSizes_default")
	symaddr(&addr_QIconEngine_iconName_default, "QIconEngine_iconName_default")
	symaddr(&addr_QIconEngine_isNull_default, "QIconEngine_isNull_default")
	reg(&pg_QIconEngine_scaledPixmap_default, "QIconEngine_scaledPixmap_default")
	reg(&pg_QIconEngine__set_override, "QIconEngine__set_override")
	reg(&pg_QIconEngine__Go_new, "QIconEngine__Go_new")
	reg(&pg_QIconEngine_delete, "QIconEngine_delete")
}

type QIconEngine struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QIconEngine) UnsafeGoID() int64     { return this.vid }
func (this *QIconEngine) UnsafeSetGoID(id int64) { this.vid = id }

func bindQIconEngine(h uintptr) *QIconEngine {
	if h == 0 {
		return nil
	}
	return &QIconEngine{h: h}
}

func newownedQIconEngine(h uintptr) *QIconEngine {
	o := bindQIconEngine(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QIconEngine).Delete)
	}
	return o
}

func (this *QIconEngine) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIconEngine) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIconEngine(h unsafe.Pointer) *QIconEngine { return bindQIconEngine(uintptr(h)) }
func UnsafeNewQIconEngineOwned(h unsafe.Pointer) *QIconEngine { return newownedQIconEngine(uintptr(h)) }

func (this *QIconEngine) UnsafeRelease() { this.h = 0 }

func (this *QIconEngine) Paint(painter *QPainter, rect *QRect, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIconEngine_paint, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QIconEngine) ActualSize(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIconEngine_actualSize, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIconEngine) Pixmap(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIconEngine_pixmap, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIconEngine) AddPixmap(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIconEngine_addPixmap, this.h, uintptr(pixmap.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIconEngine) AddFile(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIconEngine_addFile, this.h, fileName_p, uintptr(len(fileName)), uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(size)
}

func (this *QIconEngine) Key() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_key, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIconEngine) Clone() *QIconEngine {
	__r, _, _ := purego.SyscallN(addr_QIconEngine_clone, this.h)
	runtime.KeepAlive(this)
	return bindQIconEngine(__r)
}

func (this *QIconEngine) Read(in *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_read, this.h, uintptr(in.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
	return __r != 0
}

func (this *QIconEngine) Write(out *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_write, this.h, uintptr(out.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
	return __r != 0
}

func (this *QIconEngine) AvailableSizes(mode QIcon__Mode, state QIcon__State) []*QSize {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_availableSizes, this.h, uintptr(mode), uintptr(state), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QSize, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQSize(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QIconEngine) IconName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_iconName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIconEngine) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIconEngine) ScaledPixmap(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap {
	__r := pg_QIconEngine_scaledPixmap(this.h, uintptr(size.UnsafePointer()), int32(mode), int32(state), scale)
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIconEngine) ActualSizeDefault(size *QSize, mode QIcon__Mode, state QIcon__State) *QSize {
	__r, _, _ := purego.SyscallN(addr_QIconEngine_actualSize_default, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQSize(__r)
}

func (this *QIconEngine) PixmapDefault(size *QSize, mode QIcon__Mode, state QIcon__State) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QIconEngine_pixmap_default, this.h, uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func (this *QIconEngine) AddPixmapDefault(pixmap *QPixmap, mode QIcon__Mode, state QIcon__State) {
	purego.SyscallN(addr_QIconEngine_addPixmap_default, this.h, uintptr(pixmap.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
}

func (this *QIconEngine) AddFileDefault(fileName string, size *QSize, mode QIcon__Mode, state QIcon__State) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QIconEngine_addFile_default, this.h, fileName_p, uintptr(len(fileName)), uintptr(size.UnsafePointer()), uintptr(mode), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(size)
}

func (this *QIconEngine) KeyDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_key_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIconEngine) ReadDefault(in *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_read_default, this.h, uintptr(in.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
	return __r != 0
}

func (this *QIconEngine) WriteDefault(out *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_write_default, this.h, uintptr(out.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
	return __r != 0
}

func (this *QIconEngine) AvailableSizesDefault(mode QIcon__Mode, state QIcon__State) []*QSize {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_availableSizes_default, this.h, uintptr(mode), uintptr(state), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QSize, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQSize(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QIconEngine) IconNameDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIconEngine_iconName_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIconEngine) IsNullDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QIconEngine_isNull_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIconEngine) ScaledPixmapDefault(size *QSize, mode QIcon__Mode, state QIcon__State, scale float64) *QPixmap {
	__r := pg_QIconEngine_scaledPixmap_default(this.h, uintptr(size.UnsafePointer()), int32(mode), int32(state), scale)
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
	return newownedQPixmap(__r)
}

func QIconEngine_virtIdx(slot int64) int32 {
	switch slot {
	case 6083584261053556819:
		return 0
	case -7032813281048616478:
		return 1
	case 6795227186376262862:
		return 2
	case -5258308936133250472:
		return 3
	case -3986256266470558159:
		return 4
	case -5081758061965960830:
		return 5
	case 4452171178779021548:
		return 6
	case 5541207787354052101:
		return 7
	case -5099743077423337732:
		return 8
	case 7223428027511831406:
		return 9
	case -3060016504017260269:
		return 10
	case 2193725045587114196:
		return 11
	case -1544658507073703166:
		return 12
	}
	return -1
}

func QIconEngine_enable(h uintptr, idx int32) { pg_QIconEngine__set_override(h, idx) }

func NewQIconEngine() *QIconEngine {
	id := nextVirtID()
	o := newownedQIconEngine(pg_QIconEngine__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QIconEngine_virtIdx, QIconEngine_enable)
	return o
}

func (this *QIconEngine) OnPaint(fn func(*QPainter, *QRect, QIcon__Mode, QIcon__State)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, rect uintptr, mode int32, state int32) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), QIcon__Mode(mode), QIcon__State(state)) }) // slot "paint"
}

func (this *QIconEngine) OnClone(fn func() *QIconEngine) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QIconEngine) OnActualSize(fn func(*QSize, QIcon__Mode, QIcon__State) *QSize) {
	registerVirt(this.UnsafeGoID(), 6795227186376262862, func(size uintptr, mode int32, state int32) uintptr { return uintptr(fn(UnsafeNewQSize(unsafe.Pointer(size)), QIcon__Mode(mode), QIcon__State(state)).UnsafePointer()) }) // slot "actualSize"
}

func (this *QIconEngine) OnPixmap(fn func(*QSize, QIcon__Mode, QIcon__State) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -5258308936133250472, func(size uintptr, mode int32, state int32) uintptr { return uintptr(fn(UnsafeNewQSize(unsafe.Pointer(size)), QIcon__Mode(mode), QIcon__State(state)).UnsafePointer()) }) // slot "pixmap"
}

func (this *QIconEngine) OnAddPixmap(fn func(*QPixmap, QIcon__Mode, QIcon__State)) {
	registerVirt(this.UnsafeGoID(), -3986256266470558159, func(pixmap uintptr, mode int32, state int32) { fn(UnsafeNewQPixmap(unsafe.Pointer(pixmap)), QIcon__Mode(mode), QIcon__State(state)) }) // slot "addPixmap"
}

func (this *QIconEngine) OnAddFile(fn func(string, *QSize, QIcon__Mode, QIcon__State)) {
	registerVirt(this.UnsafeGoID(), -5081758061965960830, func(fileName uintptr, fileName_len int64, size uintptr, mode int32, state int32) { fn(pgStrNoFree(fileName, fileName_len), UnsafeNewQSize(unsafe.Pointer(size)), QIcon__Mode(mode), QIcon__State(state)) }) // slot "addFile"
}

func (this *QIconEngine) OnKey(fn func() string) {
	registerVirt(this.UnsafeGoID(), 4452171178779021548, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "key"
}

func (this *QIconEngine) OnRead(fn func(*QDataStream) bool) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(in uintptr) int32 { return b2i(fn(UnsafeNewQDataStream(unsafe.Pointer(in)))) }) // slot "read"
}

func (this *QIconEngine) OnWrite(fn func(*QDataStream) bool) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(out uintptr) int32 { return b2i(fn(UnsafeNewQDataStream(unsafe.Pointer(out)))) }) // slot "write"
}

func (this *QIconEngine) OnAvailableSizes(fn func(QIcon__Mode, QIcon__State) []*QSize) {
	registerVirt(this.UnsafeGoID(), 7223428027511831406, func(mode int32, state int32) (uintptr, int64) {
		__rs := fn(QIcon__Mode(mode), QIcon__State(state))
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "availableSizes"
}

func (this *QIconEngine) OnIconName(fn func() string) {
	registerVirt(this.UnsafeGoID(), -3060016504017260269, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "iconName"
}

func (this *QIconEngine) OnIsNull(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 2193725045587114196, func() int32 { return b2i(fn()) }) // slot "isNull"
}

func (this *QIconEngine) OnScaledPixmap(fn func(*QSize, QIcon__Mode, QIcon__State, float64) *QPixmap) {
	registerVirt(this.UnsafeGoID(), -1544658507073703166, func(size uintptr, mode int32, state int32, scale uintptr) uintptr { return uintptr(fn(UnsafeNewQSize(unsafe.Pointer(size)), QIcon__Mode(mode), QIcon__State(state), math.Float64frombits(uint64(scale))).UnsafePointer()) }) // slot "scaledPixmap"
}

func (this *QIconEngine) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIconEngine_delete(this.h)
	}
	this.UnsafeRelease()
}
