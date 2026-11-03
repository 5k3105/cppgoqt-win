package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSystemSemaphore_new uintptr
	addr_QSystemSemaphore_new2 uintptr
	addr_QSystemSemaphore_new3 uintptr
	addr_QSystemSemaphore_new4 uintptr
	addr_QSystemSemaphore_new5 uintptr
	addr_QSystemSemaphore_new6 uintptr
	addr_QSystemSemaphore_tr uintptr
	addr_QSystemSemaphore_setNativeKey uintptr
	addr_QSystemSemaphore_setNativeKey2 uintptr
	addr_QSystemSemaphore_nativeIpcKey uintptr
	addr_QSystemSemaphore_setKey uintptr
	addr_QSystemSemaphore_key uintptr
	addr_QSystemSemaphore_acquire uintptr
	addr_QSystemSemaphore_release uintptr
	addr_QSystemSemaphore_error uintptr
	addr_QSystemSemaphore_errorString uintptr
	addr_QSystemSemaphore_isKeyTypeSupported uintptr
	addr_QSystemSemaphore_platformSafeKey uintptr
	addr_QSystemSemaphore_legacyNativeKey uintptr
	addr_QSystemSemaphore_tr2 uintptr
	addr_QSystemSemaphore_tr3 uintptr
	addr_QSystemSemaphore_setNativeKey3 uintptr
	addr_QSystemSemaphore_setNativeKey4 uintptr
	addr_QSystemSemaphore_setNativeKey5 uintptr
	addr_QSystemSemaphore_setNativeKey6 uintptr
	addr_QSystemSemaphore_setNativeKey7 uintptr
	addr_QSystemSemaphore_setKey2 uintptr
	addr_QSystemSemaphore_setKey3 uintptr
	addr_QSystemSemaphore_release2 uintptr
	addr_QSystemSemaphore_platformSafeKey2 uintptr
	addr_QSystemSemaphore_legacyNativeKey2 uintptr
	pg_QSystemSemaphore_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSystemSemaphore_new, "QSystemSemaphore_new")
	symaddr(&addr_QSystemSemaphore_new2, "QSystemSemaphore_new2")
	symaddr(&addr_QSystemSemaphore_new3, "QSystemSemaphore_new3")
	symaddr(&addr_QSystemSemaphore_new4, "QSystemSemaphore_new4")
	symaddr(&addr_QSystemSemaphore_new5, "QSystemSemaphore_new5")
	symaddr(&addr_QSystemSemaphore_new6, "QSystemSemaphore_new6")
	symaddr(&addr_QSystemSemaphore_tr, "QSystemSemaphore_tr")
	symaddr(&addr_QSystemSemaphore_setNativeKey, "QSystemSemaphore_setNativeKey")
	symaddr(&addr_QSystemSemaphore_setNativeKey2, "QSystemSemaphore_setNativeKey2")
	symaddr(&addr_QSystemSemaphore_nativeIpcKey, "QSystemSemaphore_nativeIpcKey")
	symaddr(&addr_QSystemSemaphore_setKey, "QSystemSemaphore_setKey")
	symaddr(&addr_QSystemSemaphore_key, "QSystemSemaphore_key")
	symaddr(&addr_QSystemSemaphore_acquire, "QSystemSemaphore_acquire")
	symaddr(&addr_QSystemSemaphore_release, "QSystemSemaphore_release")
	symaddr(&addr_QSystemSemaphore_error, "QSystemSemaphore_error")
	symaddr(&addr_QSystemSemaphore_errorString, "QSystemSemaphore_errorString")
	symaddr(&addr_QSystemSemaphore_isKeyTypeSupported, "QSystemSemaphore_isKeyTypeSupported")
	symaddr(&addr_QSystemSemaphore_platformSafeKey, "QSystemSemaphore_platformSafeKey")
	symaddr(&addr_QSystemSemaphore_legacyNativeKey, "QSystemSemaphore_legacyNativeKey")
	symaddr(&addr_QSystemSemaphore_tr2, "QSystemSemaphore_tr2")
	symaddr(&addr_QSystemSemaphore_tr3, "QSystemSemaphore_tr3")
	symaddr(&addr_QSystemSemaphore_setNativeKey3, "QSystemSemaphore_setNativeKey3")
	symaddr(&addr_QSystemSemaphore_setNativeKey4, "QSystemSemaphore_setNativeKey4")
	symaddr(&addr_QSystemSemaphore_setNativeKey5, "QSystemSemaphore_setNativeKey5")
	symaddr(&addr_QSystemSemaphore_setNativeKey6, "QSystemSemaphore_setNativeKey6")
	symaddr(&addr_QSystemSemaphore_setNativeKey7, "QSystemSemaphore_setNativeKey7")
	symaddr(&addr_QSystemSemaphore_setKey2, "QSystemSemaphore_setKey2")
	symaddr(&addr_QSystemSemaphore_setKey3, "QSystemSemaphore_setKey3")
	symaddr(&addr_QSystemSemaphore_release2, "QSystemSemaphore_release2")
	symaddr(&addr_QSystemSemaphore_platformSafeKey2, "QSystemSemaphore_platformSafeKey2")
	symaddr(&addr_QSystemSemaphore_legacyNativeKey2, "QSystemSemaphore_legacyNativeKey2")
	reg(&pg_QSystemSemaphore_delete, "QSystemSemaphore_delete")
}

type QSystemSemaphore struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSystemSemaphore) UnsafeGoID() int64     { return this.vid }
func (this *QSystemSemaphore) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSystemSemaphore(h uintptr) *QSystemSemaphore {
	if h == 0 {
		return nil
	}
	return &QSystemSemaphore{h: h}
}

func newownedQSystemSemaphore(h uintptr) *QSystemSemaphore {
	o := bindQSystemSemaphore(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSystemSemaphore).Delete)
	}
	return o
}

func (this *QSystemSemaphore) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSystemSemaphore) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSystemSemaphore(h unsafe.Pointer) *QSystemSemaphore { return bindQSystemSemaphore(uintptr(h)) }
func UnsafeNewQSystemSemaphoreOwned(h unsafe.Pointer) *QSystemSemaphore { return newownedQSystemSemaphore(uintptr(h)) }

func (this *QSystemSemaphore) UnsafeRelease() { this.h = 0 }

func NewQSystemSemaphore(key *QNativeIpcKey) *QSystemSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new, uintptr(key.UnsafePointer()))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQSystemSemaphore2(key string) *QSystemSemaphore {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new2, key_p, uintptr(len(key)))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQSystemSemaphore3(key *QNativeIpcKey, initialValue int) *QSystemSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new3, uintptr(key.UnsafePointer()), uintptr(initialValue))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQSystemSemaphore4(key *QNativeIpcKey, initialValue int, arg3 QSystemSemaphore__AccessMode) *QSystemSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new4, uintptr(key.UnsafePointer()), uintptr(initialValue), uintptr(arg3))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQSystemSemaphore5(key string, initialValue int) *QSystemSemaphore {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new5, key_p, uintptr(len(key)), uintptr(initialValue))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQSystemSemaphore6(key string, initialValue int, mode QSystemSemaphore__AccessMode) *QSystemSemaphore {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_new6, key_p, uintptr(len(key)), uintptr(initialValue), uintptr(mode))
	__o := newownedQSystemSemaphore(__r)
	runtime.KeepAlive(key)
	return __o
}

func QSystemSemaphore_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QSystemSemaphore) SetNativeKey(key *QNativeIpcKey) {
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetNativeKeyWithKey(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey2, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) NativeIpcKey() *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_nativeIpcKey, this.h)
	runtime.KeepAlive(this)
	return newownedQNativeIpcKey(__r)
}

func (this *QSystemSemaphore) SetKey(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setKey, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) Key() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_key, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSystemSemaphore) Acquire() bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemSemaphore_acquire, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSystemSemaphore) Release() bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemSemaphore_release, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSystemSemaphore) Error() QSystemSemaphore__SystemSemaphoreError {
	__r0, _, _ := purego.SyscallN(addr_QSystemSemaphore_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSystemSemaphore__SystemSemaphoreError(__r)
}

func (this *QSystemSemaphore) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QSystemSemaphore_IsKeyTypeSupported(type_ QNativeIpcKey__Type) bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemSemaphore_isKeyTypeSupported, uintptr(type_))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QSystemSemaphore_PlatformSafeKey(key string) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_platformSafeKey, key_p, uintptr(len(key)))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSystemSemaphore_LegacyNativeKey(key string) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_legacyNativeKey, key_p, uintptr(len(key)))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSystemSemaphore_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QSystemSemaphore_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QSystemSemaphore) SetNativeKey2(key *QNativeIpcKey, initialValue int) {
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey3, this.h, uintptr(key.UnsafePointer()), uintptr(initialValue))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetNativeKey3(key *QNativeIpcKey, initialValue int, arg3 QSystemSemaphore__AccessMode) {
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey4, this.h, uintptr(key.UnsafePointer()), uintptr(initialValue), uintptr(arg3))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetNativeKey4(key string, initialValue int) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey5, this.h, key_p, uintptr(len(key)), uintptr(initialValue))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetNativeKey5(key string, initialValue int, mode QSystemSemaphore__AccessMode) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey6, this.h, key_p, uintptr(len(key)), uintptr(initialValue), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetNativeKey6(key string, initialValue int, mode QSystemSemaphore__AccessMode, type_ QNativeIpcKey__Type) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setNativeKey7, this.h, key_p, uintptr(len(key)), uintptr(initialValue), uintptr(mode), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetKey2(key string, initialValue int) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setKey2, this.h, key_p, uintptr(len(key)), uintptr(initialValue))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) SetKey3(key string, initialValue int, mode QSystemSemaphore__AccessMode) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSystemSemaphore_setKey3, this.h, key_p, uintptr(len(key)), uintptr(initialValue), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSystemSemaphore) ReleaseWithInt(n int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSystemSemaphore_release2, this.h, uintptr(n))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSystemSemaphore_PlatformSafeKey2(key string, type_ QNativeIpcKey__Type) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_platformSafeKey2, key_p, uintptr(len(key)), uintptr(type_))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSystemSemaphore_LegacyNativeKey2(key string, type_ QNativeIpcKey__Type) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSystemSemaphore_legacyNativeKey2, key_p, uintptr(len(key)), uintptr(type_))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func (this *QSystemSemaphore) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSystemSemaphore_delete(this.h)
	}
	this.UnsafeRelease()
}
