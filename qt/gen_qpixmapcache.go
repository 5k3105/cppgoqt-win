package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPixmapCache_cacheLimit uintptr
	addr_QPixmapCache_setCacheLimit uintptr
	addr_QPixmapCache_find uintptr
	addr_QPixmapCache_find2 uintptr
	addr_QPixmapCache_insert uintptr
	addr_QPixmapCache_insert2 uintptr
	addr_QPixmapCache_replace uintptr
	addr_QPixmapCache_remove uintptr
	addr_QPixmapCache_remove2 uintptr
	addr_QPixmapCache_clear uintptr
	pg_QPixmapCache_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPixmapCache_cacheLimit, "QPixmapCache_cacheLimit")
	symaddr(&addr_QPixmapCache_setCacheLimit, "QPixmapCache_setCacheLimit")
	symaddr(&addr_QPixmapCache_find, "QPixmapCache_find")
	symaddr(&addr_QPixmapCache_find2, "QPixmapCache_find2")
	symaddr(&addr_QPixmapCache_insert, "QPixmapCache_insert")
	symaddr(&addr_QPixmapCache_insert2, "QPixmapCache_insert2")
	symaddr(&addr_QPixmapCache_replace, "QPixmapCache_replace")
	symaddr(&addr_QPixmapCache_remove, "QPixmapCache_remove")
	symaddr(&addr_QPixmapCache_remove2, "QPixmapCache_remove2")
	symaddr(&addr_QPixmapCache_clear, "QPixmapCache_clear")
	reg(&pg_QPixmapCache_delete, "QPixmapCache_delete")
}

type QPixmapCache struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPixmapCache) UnsafeGoID() int64     { return this.vid }
func (this *QPixmapCache) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPixmapCache(h uintptr) *QPixmapCache {
	if h == 0 {
		return nil
	}
	return &QPixmapCache{h: h}
}

func newownedQPixmapCache(h uintptr) *QPixmapCache {
	o := bindQPixmapCache(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPixmapCache).Delete)
	}
	return o
}

func (this *QPixmapCache) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPixmapCache) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPixmapCache(h unsafe.Pointer) *QPixmapCache { return bindQPixmapCache(uintptr(h)) }
func UnsafeNewQPixmapCacheOwned(h unsafe.Pointer) *QPixmapCache { return newownedQPixmapCache(uintptr(h)) }

func (this *QPixmapCache) UnsafeRelease() { this.h = 0 }

func QPixmapCache_CacheLimit() int {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache_cacheLimit)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QPixmapCache_SetCacheLimit(arg1 int) {
	purego.SyscallN(addr_QPixmapCache_setCacheLimit, uintptr(arg1))
}

func QPixmapCache_Find(key string, pixmap *QPixmap) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache_find, key_p, uintptr(len(key)), uintptr(pixmap.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(key)
	runtime.KeepAlive(pixmap)
	return __r != 0
}

func QPixmapCache_Find2(key *QPixmapCache__Key, pixmap *QPixmap) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache_find2, uintptr(key.UnsafePointer()), uintptr(pixmap.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(key)
	runtime.KeepAlive(pixmap)
	return __r != 0
}

func QPixmapCache_Insert(key string, pixmap *QPixmap) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache_insert, key_p, uintptr(len(key)), uintptr(pixmap.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(key)
	runtime.KeepAlive(pixmap)
	return __r != 0
}

func QPixmapCache_InsertWithPixmap(pixmap *QPixmap) *QPixmapCache__Key {
	__r, _, _ := purego.SyscallN(addr_QPixmapCache_insert2, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(pixmap)
	return newownedQPixmapCache__Key(__r)
}

func QPixmapCache_Replace(key *QPixmapCache__Key, pixmap *QPixmap) bool {
	__r0, _, _ := purego.SyscallN(addr_QPixmapCache_replace, uintptr(key.UnsafePointer()), uintptr(pixmap.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(key)
	runtime.KeepAlive(pixmap)
	return __r != 0
}

func QPixmapCache_Remove(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QPixmapCache_remove, key_p, uintptr(len(key)))
	runtime.KeepAlive(key)
}

func QPixmapCache_RemoveWithKey(key *QPixmapCache__Key) {
	purego.SyscallN(addr_QPixmapCache_remove2, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(key)
}

func QPixmapCache_Clear() {
	purego.SyscallN(addr_QPixmapCache_clear)
}

func (this *QPixmapCache) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPixmapCache_delete(this.h)
	}
	this.UnsafeRelease()
}
