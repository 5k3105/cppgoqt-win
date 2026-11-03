package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QShortcut_base func(uintptr) uintptr
	addr_QShortcut_new uintptr
	addr_QShortcut_new2 uintptr
	addr_QShortcut_new3 uintptr
	addr_QShortcut_new4 uintptr
	addr_QShortcut_new5 uintptr
	addr_QShortcut_new6 uintptr
	addr_QShortcut_new7 uintptr
	addr_QShortcut_new8 uintptr
	addr_QShortcut_new9 uintptr
	addr_QShortcut_tr uintptr
	addr_QShortcut_setKey uintptr
	addr_QShortcut_key uintptr
	addr_QShortcut_setKeys uintptr
	addr_QShortcut_setKeys2 uintptr
	addr_QShortcut_keys uintptr
	addr_QShortcut_setEnabled uintptr
	addr_QShortcut_isEnabled uintptr
	addr_QShortcut_setContext uintptr
	addr_QShortcut_context uintptr
	addr_QShortcut_setAutoRepeat uintptr
	addr_QShortcut_autoRepeat uintptr
	addr_QShortcut_id uintptr
	addr_QShortcut_setWhatsThis uintptr
	addr_QShortcut_whatsThis uintptr
	addr_QShortcut_activated uintptr
	addr_QShortcut_activatedAmbiguously uintptr
	addr_QShortcut_connect uintptr
	addr_QShortcut_connect2 uintptr
	addr_QShortcut_sender uintptr
	addr_QShortcut_senderSignalIndex uintptr
	addr_QShortcut_receivers uintptr
	addr_QShortcut_isSignalConnected uintptr
	addr_QShortcut_tr2 uintptr
	addr_QShortcut_tr3 uintptr
	addr_QShortcut_connect3 uintptr
	addr_QShortcut_connect4 uintptr
	addr_QShortcut_metaObject_default uintptr
	addr_QShortcut_event_default uintptr
	addr_QShortcut_eventFilter_default uintptr
	addr_QShortcut_timerEvent_default uintptr
	addr_QShortcut_childEvent_default uintptr
	addr_QShortcut_customEvent_default uintptr
	addr_QShortcut_connectNotify_default uintptr
	addr_QShortcut_disconnectNotify_default uintptr
	pg_QShortcut_connect_activated func(uintptr, int64)
	pg_QShortcut_connect_activatedAmbiguously func(uintptr, int64)
	pg_QShortcut_connect_destroyed func(uintptr, int64)
	pg_QShortcut_connect_destroyed2 func(uintptr, int64)
	pg_QShortcut_connect_objectNameChanged func(uintptr, int64)
	pg_QShortcut__disconnect func(int64)
	pg_QShortcut__set_override func(uintptr, int32)
	pg_QShortcut_delete func(uintptr)
)

func init() {
	reg(&pg_QShortcut_base, "QShortcut_base")
	symaddr(&addr_QShortcut_new, "QShortcut_new")
	symaddr(&addr_QShortcut_new2, "QShortcut_new2")
	symaddr(&addr_QShortcut_new3, "QShortcut_new3")
	symaddr(&addr_QShortcut_new4, "QShortcut_new4")
	symaddr(&addr_QShortcut_new5, "QShortcut_new5")
	symaddr(&addr_QShortcut_new6, "QShortcut_new6")
	symaddr(&addr_QShortcut_new7, "QShortcut_new7")
	symaddr(&addr_QShortcut_new8, "QShortcut_new8")
	symaddr(&addr_QShortcut_new9, "QShortcut_new9")
	symaddr(&addr_QShortcut_tr, "QShortcut_tr")
	symaddr(&addr_QShortcut_setKey, "QShortcut_setKey")
	symaddr(&addr_QShortcut_key, "QShortcut_key")
	symaddr(&addr_QShortcut_setKeys, "QShortcut_setKeys")
	symaddr(&addr_QShortcut_setKeys2, "QShortcut_setKeys2")
	symaddr(&addr_QShortcut_keys, "QShortcut_keys")
	symaddr(&addr_QShortcut_setEnabled, "QShortcut_setEnabled")
	symaddr(&addr_QShortcut_isEnabled, "QShortcut_isEnabled")
	symaddr(&addr_QShortcut_setContext, "QShortcut_setContext")
	symaddr(&addr_QShortcut_context, "QShortcut_context")
	symaddr(&addr_QShortcut_setAutoRepeat, "QShortcut_setAutoRepeat")
	symaddr(&addr_QShortcut_autoRepeat, "QShortcut_autoRepeat")
	symaddr(&addr_QShortcut_id, "QShortcut_id")
	symaddr(&addr_QShortcut_setWhatsThis, "QShortcut_setWhatsThis")
	symaddr(&addr_QShortcut_whatsThis, "QShortcut_whatsThis")
	symaddr(&addr_QShortcut_activated, "QShortcut_activated")
	symaddr(&addr_QShortcut_activatedAmbiguously, "QShortcut_activatedAmbiguously")
	symaddr(&addr_QShortcut_connect, "QShortcut_connect")
	symaddr(&addr_QShortcut_connect2, "QShortcut_connect2")
	symaddr(&addr_QShortcut_sender, "QShortcut_sender")
	symaddr(&addr_QShortcut_senderSignalIndex, "QShortcut_senderSignalIndex")
	symaddr(&addr_QShortcut_receivers, "QShortcut_receivers")
	symaddr(&addr_QShortcut_isSignalConnected, "QShortcut_isSignalConnected")
	symaddr(&addr_QShortcut_tr2, "QShortcut_tr2")
	symaddr(&addr_QShortcut_tr3, "QShortcut_tr3")
	symaddr(&addr_QShortcut_connect3, "QShortcut_connect3")
	symaddr(&addr_QShortcut_connect4, "QShortcut_connect4")
	symaddr(&addr_QShortcut_metaObject_default, "QShortcut_metaObject_default")
	symaddr(&addr_QShortcut_event_default, "QShortcut_event_default")
	symaddr(&addr_QShortcut_eventFilter_default, "QShortcut_eventFilter_default")
	symaddr(&addr_QShortcut_timerEvent_default, "QShortcut_timerEvent_default")
	symaddr(&addr_QShortcut_childEvent_default, "QShortcut_childEvent_default")
	symaddr(&addr_QShortcut_customEvent_default, "QShortcut_customEvent_default")
	symaddr(&addr_QShortcut_connectNotify_default, "QShortcut_connectNotify_default")
	symaddr(&addr_QShortcut_disconnectNotify_default, "QShortcut_disconnectNotify_default")
	reg(&pg_QShortcut_connect_activated, "QShortcut_connect_activated")
	reg(&pg_QShortcut_connect_activatedAmbiguously, "QShortcut_connect_activatedAmbiguously")
	reg(&pg_QShortcut_connect_destroyed, "QShortcut_connect_destroyed")
	reg(&pg_QShortcut_connect_destroyed2, "QShortcut_connect_destroyed2")
	reg(&pg_QShortcut_connect_objectNameChanged, "QShortcut_connect_objectNameChanged")
	reg(&pg_QShortcut__disconnect, "QShortcut__disconnect")
	reg(&pg_QShortcut__set_override, "QShortcut__set_override")
	reg(&pg_QShortcut_delete, "QShortcut_delete")
}

type QShortcut struct {
	h uintptr
	*QObject
}

func bindQShortcut(h uintptr) *QShortcut {
	if h == 0 {
		return nil
	}
	return &QShortcut{h: h, QObject: bindQObject(pg_QShortcut_base(h))}
}

func newownedQShortcut(h uintptr) *QShortcut { return bindQShortcut(h) }

func (this *QShortcut) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QShortcut) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQShortcut(h unsafe.Pointer) *QShortcut { return bindQShortcut(uintptr(h)) }
func UnsafeNewQShortcutOwned(h unsafe.Pointer) *QShortcut { return newownedQShortcut(uintptr(h)) }

func (this *QShortcut) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQShortcut(parent *QObject) *QShortcut {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QShortcut_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut2(key *QKeySequence, parent *QObject) *QShortcut {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QShortcut_new2, uintptr(__goid), uintptr(key.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut3(key QKeySequence__StandardKey, parent *QObject) *QShortcut {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QShortcut_new3, uintptr(__goid), uintptr(key), uintptr(parent.UnsafePointer()))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut4(key *QKeySequence, parent *QObject, member string) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new4, uintptr(__goid), uintptr(key.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut5(key *QKeySequence, parent *QObject, member string, ambiguousMember string) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	ambiguousMember_b := append([]byte(ambiguousMember), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new5, uintptr(__goid), uintptr(key.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(unsafe.Pointer(&ambiguousMember_b[0])))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ambiguousMember_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut6(key *QKeySequence, parent *QObject, member string, ambiguousMember string, context ShortcutContext) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	ambiguousMember_b := append([]byte(ambiguousMember), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new6, uintptr(__goid), uintptr(key.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(unsafe.Pointer(&ambiguousMember_b[0])), uintptr(context))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(key)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ambiguousMember_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut7(key QKeySequence__StandardKey, parent *QObject, member string) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new7, uintptr(__goid), uintptr(key), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut8(key QKeySequence__StandardKey, parent *QObject, member string, ambiguousMember string) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	ambiguousMember_b := append([]byte(ambiguousMember), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new8, uintptr(__goid), uintptr(key), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(unsafe.Pointer(&ambiguousMember_b[0])))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ambiguousMember_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func NewQShortcut9(key QKeySequence__StandardKey, parent *QObject, member string, ambiguousMember string, context ShortcutContext) *QShortcut {
	__goid := nextVirtID()
	member_b := append([]byte(member), 0)
	ambiguousMember_b := append([]byte(ambiguousMember), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_new9, uintptr(__goid), uintptr(key), uintptr(parent.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(unsafe.Pointer(&ambiguousMember_b[0])), uintptr(context))
	__o := newownedQShortcut(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(member_b)
	runtime.KeepAlive(ambiguousMember_b)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QShortcut_virtIdx, QShortcut_enable)
	return __o
}

func QShortcut_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QShortcut_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QShortcut) SetKey(key *QKeySequence) {
	purego.SyscallN(addr_QShortcut_setKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QShortcut) Key() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QShortcut_key, this.h)
	runtime.KeepAlive(this)
	return newownedQKeySequence(__r)
}

func (this *QShortcut) SetKeys(key QKeySequence__StandardKey) {
	purego.SyscallN(addr_QShortcut_setKeys, this.h, uintptr(key))
	runtime.KeepAlive(this)
}

func (this *QShortcut) SetKeysWithKeys(keys []*QKeySequence) {
	keys_arr := make([]uintptr, len(keys))
	for __i, __e := range keys {
		keys_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var keys_p uintptr
	if len(keys_arr) > 0 {
		keys_p = uintptr(unsafe.Pointer(&keys_arr[0]))
	}
	purego.SyscallN(addr_QShortcut_setKeys2, this.h, keys_p, uintptr(len(keys)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keys)
	runtime.KeepAlive(keys_arr)
}

func (this *QShortcut) Keys() []*QKeySequence {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QShortcut_keys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QKeySequence, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeySequence(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QShortcut) SetEnabled(enable bool) {
	purego.SyscallN(addr_QShortcut_setEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QShortcut) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QShortcut) SetContext(context ShortcutContext) {
	purego.SyscallN(addr_QShortcut_setContext, this.h, uintptr(context))
	runtime.KeepAlive(this)
}

func (this *QShortcut) Context() ShortcutContext {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_context, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ShortcutContext(__r)
}

func (this *QShortcut) SetAutoRepeat(on bool) {
	purego.SyscallN(addr_QShortcut_setAutoRepeat, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QShortcut) AutoRepeat() bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_autoRepeat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QShortcut) Id() int {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QShortcut) SetWhatsThis(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QShortcut_setWhatsThis, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QShortcut) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QShortcut_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QShortcut) Activated() {
	purego.SyscallN(addr_QShortcut_activated, this.h)
	runtime.KeepAlive(this)
}

func (this *QShortcut) ActivatedAmbiguously() {
	purego.SyscallN(addr_QShortcut_activatedAmbiguously, this.h)
	runtime.KeepAlive(this)
}

func QShortcut_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QShortcut_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QShortcut_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QShortcut) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QShortcut_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QShortcut) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QShortcut) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QShortcut_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QShortcut) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QShortcut_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QShortcut_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QShortcut_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QShortcut_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QShortcut_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QShortcut_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QShortcut_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QShortcut_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QShortcut) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QShortcut_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QShortcut) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QShortcut) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QShortcut_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QShortcut) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QShortcut_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QShortcut) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QShortcut_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QShortcut) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QShortcut_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QShortcut) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QShortcut_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QShortcut) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QShortcut_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QShortcut) OnActivated(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QShortcut_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QShortcut) OnActivatedAmbiguously(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QShortcut_connect_activatedAmbiguously(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QShortcut) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QShortcut_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QShortcut) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QShortcut_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QShortcut) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QShortcut_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QShortcut) Disconnect(handle int64) {
	pg_QShortcut__disconnect(handle)
	unregisterCallback(handle)
}

func QShortcut_virtIdx(slot int64) int32 {
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

func QShortcut_enable(h uintptr, idx int32) { pg_QShortcut__set_override(h, idx) }

func (this *QShortcut) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QShortcut) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QShortcut) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QShortcut) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QShortcut) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QShortcut) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QShortcut) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QShortcut) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QShortcut) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QShortcut_delete(this.h)
	}
	this.UnsafeRelease()
}
