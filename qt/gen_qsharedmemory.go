package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSharedMemory_base func(uintptr) uintptr
	addr_QSharedMemory_new uintptr
	addr_QSharedMemory_new2 uintptr
	addr_QSharedMemory_new3 uintptr
	addr_QSharedMemory_new4 uintptr
	addr_QSharedMemory_new5 uintptr
	addr_QSharedMemory_new6 uintptr
	addr_QSharedMemory_tr uintptr
	addr_QSharedMemory_setKey uintptr
	addr_QSharedMemory_key uintptr
	addr_QSharedMemory_setNativeKey uintptr
	addr_QSharedMemory_setNativeKey2 uintptr
	addr_QSharedMemory_nativeKey uintptr
	addr_QSharedMemory_nativeIpcKey uintptr
	addr_QSharedMemory_create uintptr
	addr_QSharedMemory_size uintptr
	addr_QSharedMemory_attach uintptr
	addr_QSharedMemory_isAttached uintptr
	addr_QSharedMemory_detach uintptr
	addr_QSharedMemory_lock uintptr
	addr_QSharedMemory_unlock uintptr
	addr_QSharedMemory_error uintptr
	addr_QSharedMemory_errorString uintptr
	addr_QSharedMemory_isKeyTypeSupported uintptr
	addr_QSharedMemory_platformSafeKey uintptr
	addr_QSharedMemory_legacyNativeKey uintptr
	addr_QSharedMemory_connect uintptr
	addr_QSharedMemory_connect2 uintptr
	addr_QSharedMemory_sender uintptr
	addr_QSharedMemory_senderSignalIndex uintptr
	addr_QSharedMemory_receivers uintptr
	addr_QSharedMemory_isSignalConnected uintptr
	addr_QSharedMemory_tr2 uintptr
	addr_QSharedMemory_tr3 uintptr
	addr_QSharedMemory_setNativeKey3 uintptr
	addr_QSharedMemory_create2 uintptr
	addr_QSharedMemory_attach2 uintptr
	addr_QSharedMemory_platformSafeKey2 uintptr
	addr_QSharedMemory_legacyNativeKey2 uintptr
	addr_QSharedMemory_connect3 uintptr
	addr_QSharedMemory_connect4 uintptr
	addr_QSharedMemory_metaObject_default uintptr
	addr_QSharedMemory_event_default uintptr
	addr_QSharedMemory_eventFilter_default uintptr
	addr_QSharedMemory_timerEvent_default uintptr
	addr_QSharedMemory_childEvent_default uintptr
	addr_QSharedMemory_customEvent_default uintptr
	addr_QSharedMemory_connectNotify_default uintptr
	addr_QSharedMemory_disconnectNotify_default uintptr
	pg_QSharedMemory_connect_destroyed func(uintptr, int64)
	pg_QSharedMemory_connect_destroyed2 func(uintptr, int64)
	pg_QSharedMemory_connect_objectNameChanged func(uintptr, int64)
	pg_QSharedMemory__disconnect func(int64)
	pg_QSharedMemory__set_override func(uintptr, int32)
	pg_QSharedMemory_delete func(uintptr)
)

func init() {
	reg(&pg_QSharedMemory_base, "QSharedMemory_base")
	symaddr(&addr_QSharedMemory_new, "QSharedMemory_new")
	symaddr(&addr_QSharedMemory_new2, "QSharedMemory_new2")
	symaddr(&addr_QSharedMemory_new3, "QSharedMemory_new3")
	symaddr(&addr_QSharedMemory_new4, "QSharedMemory_new4")
	symaddr(&addr_QSharedMemory_new5, "QSharedMemory_new5")
	symaddr(&addr_QSharedMemory_new6, "QSharedMemory_new6")
	symaddr(&addr_QSharedMemory_tr, "QSharedMemory_tr")
	symaddr(&addr_QSharedMemory_setKey, "QSharedMemory_setKey")
	symaddr(&addr_QSharedMemory_key, "QSharedMemory_key")
	symaddr(&addr_QSharedMemory_setNativeKey, "QSharedMemory_setNativeKey")
	symaddr(&addr_QSharedMemory_setNativeKey2, "QSharedMemory_setNativeKey2")
	symaddr(&addr_QSharedMemory_nativeKey, "QSharedMemory_nativeKey")
	symaddr(&addr_QSharedMemory_nativeIpcKey, "QSharedMemory_nativeIpcKey")
	symaddr(&addr_QSharedMemory_create, "QSharedMemory_create")
	symaddr(&addr_QSharedMemory_size, "QSharedMemory_size")
	symaddr(&addr_QSharedMemory_attach, "QSharedMemory_attach")
	symaddr(&addr_QSharedMemory_isAttached, "QSharedMemory_isAttached")
	symaddr(&addr_QSharedMemory_detach, "QSharedMemory_detach")
	symaddr(&addr_QSharedMemory_lock, "QSharedMemory_lock")
	symaddr(&addr_QSharedMemory_unlock, "QSharedMemory_unlock")
	symaddr(&addr_QSharedMemory_error, "QSharedMemory_error")
	symaddr(&addr_QSharedMemory_errorString, "QSharedMemory_errorString")
	symaddr(&addr_QSharedMemory_isKeyTypeSupported, "QSharedMemory_isKeyTypeSupported")
	symaddr(&addr_QSharedMemory_platformSafeKey, "QSharedMemory_platformSafeKey")
	symaddr(&addr_QSharedMemory_legacyNativeKey, "QSharedMemory_legacyNativeKey")
	symaddr(&addr_QSharedMemory_connect, "QSharedMemory_connect")
	symaddr(&addr_QSharedMemory_connect2, "QSharedMemory_connect2")
	symaddr(&addr_QSharedMemory_sender, "QSharedMemory_sender")
	symaddr(&addr_QSharedMemory_senderSignalIndex, "QSharedMemory_senderSignalIndex")
	symaddr(&addr_QSharedMemory_receivers, "QSharedMemory_receivers")
	symaddr(&addr_QSharedMemory_isSignalConnected, "QSharedMemory_isSignalConnected")
	symaddr(&addr_QSharedMemory_tr2, "QSharedMemory_tr2")
	symaddr(&addr_QSharedMemory_tr3, "QSharedMemory_tr3")
	symaddr(&addr_QSharedMemory_setNativeKey3, "QSharedMemory_setNativeKey3")
	symaddr(&addr_QSharedMemory_create2, "QSharedMemory_create2")
	symaddr(&addr_QSharedMemory_attach2, "QSharedMemory_attach2")
	symaddr(&addr_QSharedMemory_platformSafeKey2, "QSharedMemory_platformSafeKey2")
	symaddr(&addr_QSharedMemory_legacyNativeKey2, "QSharedMemory_legacyNativeKey2")
	symaddr(&addr_QSharedMemory_connect3, "QSharedMemory_connect3")
	symaddr(&addr_QSharedMemory_connect4, "QSharedMemory_connect4")
	symaddr(&addr_QSharedMemory_metaObject_default, "QSharedMemory_metaObject_default")
	symaddr(&addr_QSharedMemory_event_default, "QSharedMemory_event_default")
	symaddr(&addr_QSharedMemory_eventFilter_default, "QSharedMemory_eventFilter_default")
	symaddr(&addr_QSharedMemory_timerEvent_default, "QSharedMemory_timerEvent_default")
	symaddr(&addr_QSharedMemory_childEvent_default, "QSharedMemory_childEvent_default")
	symaddr(&addr_QSharedMemory_customEvent_default, "QSharedMemory_customEvent_default")
	symaddr(&addr_QSharedMemory_connectNotify_default, "QSharedMemory_connectNotify_default")
	symaddr(&addr_QSharedMemory_disconnectNotify_default, "QSharedMemory_disconnectNotify_default")
	reg(&pg_QSharedMemory_connect_destroyed, "QSharedMemory_connect_destroyed")
	reg(&pg_QSharedMemory_connect_destroyed2, "QSharedMemory_connect_destroyed2")
	reg(&pg_QSharedMemory_connect_objectNameChanged, "QSharedMemory_connect_objectNameChanged")
	reg(&pg_QSharedMemory__disconnect, "QSharedMemory__disconnect")
	reg(&pg_QSharedMemory__set_override, "QSharedMemory__set_override")
	reg(&pg_QSharedMemory_delete, "QSharedMemory_delete")
}

type QSharedMemory struct {
	h uintptr
	*QObject
}

func bindQSharedMemory(h uintptr) *QSharedMemory {
	if h == 0 {
		return nil
	}
	return &QSharedMemory{h: h, QObject: bindQObject(pg_QSharedMemory_base(h))}
}

func newownedQSharedMemory(h uintptr) *QSharedMemory { return bindQSharedMemory(h) }

func (this *QSharedMemory) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSharedMemory) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSharedMemory(h unsafe.Pointer) *QSharedMemory { return bindQSharedMemory(uintptr(h)) }
func UnsafeNewQSharedMemoryOwned(h unsafe.Pointer) *QSharedMemory { return newownedQSharedMemory(uintptr(h)) }

func (this *QSharedMemory) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQSharedMemory() *QSharedMemory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new, uintptr(__goid))
	__o := newownedQSharedMemory(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func NewQSharedMemory2(key *QNativeIpcKey) *QSharedMemory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new2, uintptr(__goid), uintptr(key.UnsafePointer()))
	__o := newownedQSharedMemory(__r)
	runtime.KeepAlive(key)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func NewQSharedMemory3(key string) *QSharedMemory {
	__goid := nextVirtID()
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new3, uintptr(__goid), key_p, uintptr(len(key)))
	__o := newownedQSharedMemory(__r)
	runtime.KeepAlive(key)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func NewQSharedMemory4(parent *QObject) *QSharedMemory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSharedMemory(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func NewQSharedMemory5(key *QNativeIpcKey, parent *QObject) *QSharedMemory {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new5, uintptr(__goid), uintptr(key.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQSharedMemory(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func NewQSharedMemory6(key string, parent *QObject) *QSharedMemory {
	__goid := nextVirtID()
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_new6, uintptr(__goid), key_p, uintptr(len(key)), uintptr(parent.UnsafePointer()))
	__o := newownedQSharedMemory(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSharedMemory_virtIdx, QSharedMemory_enable)
	return __o
}

func QSharedMemory_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSharedMemory) SetKey(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSharedMemory_setKey, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSharedMemory) Key() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_key, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSharedMemory) SetNativeKey(key *QNativeIpcKey) {
	purego.SyscallN(addr_QSharedMemory_setNativeKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSharedMemory) SetNativeKeyWithKey(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSharedMemory_setNativeKey2, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSharedMemory) NativeKey() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_nativeKey, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSharedMemory) NativeIpcKey() *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_nativeIpcKey, this.h)
	runtime.KeepAlive(this)
	return newownedQNativeIpcKey(__r)
}

func (this *QSharedMemory) Create(size int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_create, this.h, uintptr(size))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSharedMemory) Attach() bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_attach, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) IsAttached() bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_isAttached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) Detach() bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_detach, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) Lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) Unlock() bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_unlock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) Error() QSharedMemory__SharedMemoryError {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSharedMemory__SharedMemoryError(__r)
}

func (this *QSharedMemory) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QSharedMemory_IsKeyTypeSupported(type_ QNativeIpcKey__Type) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_isKeyTypeSupported, uintptr(type_))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QSharedMemory_PlatformSafeKey(key string) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_platformSafeKey, key_p, uintptr(len(key)))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSharedMemory_LegacyNativeKey(key string) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_legacyNativeKey, key_p, uintptr(len(key)))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSharedMemory_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSharedMemory_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSharedMemory) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSharedMemory) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSharedMemory) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSharedMemory) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSharedMemory_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSharedMemory_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSharedMemory) SetNativeKey2(key string, type_ QNativeIpcKey__Type) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSharedMemory_setNativeKey3, this.h, key_p, uintptr(len(key)), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSharedMemory) Create2(size int64, mode QSharedMemory__AccessMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_create2, this.h, uintptr(size), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSharedMemory) AttachWithMode(mode QSharedMemory__AccessMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_attach2, this.h, uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSharedMemory_PlatformSafeKey2(key string, type_ QNativeIpcKey__Type) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_platformSafeKey2, key_p, uintptr(len(key)), uintptr(type_))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSharedMemory_LegacyNativeKey2(key string, type_ QNativeIpcKey__Type) *QNativeIpcKey {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_legacyNativeKey2, key_p, uintptr(len(key)), uintptr(type_))
	runtime.KeepAlive(key)
	return newownedQNativeIpcKey(__r)
}

func QSharedMemory_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSharedMemory_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSharedMemory) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSharedMemory_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSharedMemory) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSharedMemory) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSharedMemory_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSharedMemory) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSharedMemory_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSharedMemory) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSharedMemory_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSharedMemory) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSharedMemory_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSharedMemory) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSharedMemory_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSharedMemory) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSharedMemory_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSharedMemory) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSharedMemory_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSharedMemory) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSharedMemory_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSharedMemory) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSharedMemory_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSharedMemory) Disconnect(handle int64) {
	pg_QSharedMemory__disconnect(handle)
	unregisterCallback(handle)
}

func QSharedMemory_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QSharedMemory_enable(h uintptr, idx int32) { pg_QSharedMemory__set_override(h, idx) }

func (this *QSharedMemory) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSharedMemory) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSharedMemory) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSharedMemory) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSharedMemory) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSharedMemory) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSharedMemory) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSharedMemory) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSharedMemory) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSharedMemory_delete(this.h)
	}
	this.UnsafeRelease()
}
