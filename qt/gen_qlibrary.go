package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QLibrary_base func(uintptr) uintptr
	addr_QLibrary_new uintptr
	addr_QLibrary_new2 uintptr
	addr_QLibrary_new3 uintptr
	addr_QLibrary_new4 uintptr
	addr_QLibrary_new5 uintptr
	addr_QLibrary_new6 uintptr
	addr_QLibrary_new7 uintptr
	addr_QLibrary_new8 uintptr
	addr_QLibrary_tr uintptr
	addr_QLibrary_load uintptr
	addr_QLibrary_unload uintptr
	addr_QLibrary_isLoaded uintptr
	addr_QLibrary_isLibrary uintptr
	addr_QLibrary_setFileName uintptr
	addr_QLibrary_fileName uintptr
	addr_QLibrary_setFileNameAndVersion uintptr
	addr_QLibrary_setFileNameAndVersion2 uintptr
	addr_QLibrary_errorString uintptr
	addr_QLibrary_setLoadHints uintptr
	addr_QLibrary_loadHints uintptr
	addr_QLibrary_connect uintptr
	addr_QLibrary_connect2 uintptr
	addr_QLibrary_sender uintptr
	addr_QLibrary_senderSignalIndex uintptr
	addr_QLibrary_receivers uintptr
	addr_QLibrary_isSignalConnected uintptr
	addr_QLibrary_tr2 uintptr
	addr_QLibrary_tr3 uintptr
	addr_QLibrary_connect3 uintptr
	addr_QLibrary_connect4 uintptr
	addr_QLibrary_metaObject_default uintptr
	addr_QLibrary_event_default uintptr
	addr_QLibrary_eventFilter_default uintptr
	addr_QLibrary_timerEvent_default uintptr
	addr_QLibrary_childEvent_default uintptr
	addr_QLibrary_customEvent_default uintptr
	addr_QLibrary_connectNotify_default uintptr
	addr_QLibrary_disconnectNotify_default uintptr
	pg_QLibrary_connect_destroyed func(uintptr, int64)
	pg_QLibrary_connect_destroyed2 func(uintptr, int64)
	pg_QLibrary_connect_objectNameChanged func(uintptr, int64)
	pg_QLibrary__disconnect func(int64)
	pg_QLibrary__set_override func(uintptr, int32)
	pg_QLibrary_delete func(uintptr)
)

func init() {
	reg(&pg_QLibrary_base, "QLibrary_base")
	symaddr(&addr_QLibrary_new, "QLibrary_new")
	symaddr(&addr_QLibrary_new2, "QLibrary_new2")
	symaddr(&addr_QLibrary_new3, "QLibrary_new3")
	symaddr(&addr_QLibrary_new4, "QLibrary_new4")
	symaddr(&addr_QLibrary_new5, "QLibrary_new5")
	symaddr(&addr_QLibrary_new6, "QLibrary_new6")
	symaddr(&addr_QLibrary_new7, "QLibrary_new7")
	symaddr(&addr_QLibrary_new8, "QLibrary_new8")
	symaddr(&addr_QLibrary_tr, "QLibrary_tr")
	symaddr(&addr_QLibrary_load, "QLibrary_load")
	symaddr(&addr_QLibrary_unload, "QLibrary_unload")
	symaddr(&addr_QLibrary_isLoaded, "QLibrary_isLoaded")
	symaddr(&addr_QLibrary_isLibrary, "QLibrary_isLibrary")
	symaddr(&addr_QLibrary_setFileName, "QLibrary_setFileName")
	symaddr(&addr_QLibrary_fileName, "QLibrary_fileName")
	symaddr(&addr_QLibrary_setFileNameAndVersion, "QLibrary_setFileNameAndVersion")
	symaddr(&addr_QLibrary_setFileNameAndVersion2, "QLibrary_setFileNameAndVersion2")
	symaddr(&addr_QLibrary_errorString, "QLibrary_errorString")
	symaddr(&addr_QLibrary_setLoadHints, "QLibrary_setLoadHints")
	symaddr(&addr_QLibrary_loadHints, "QLibrary_loadHints")
	symaddr(&addr_QLibrary_connect, "QLibrary_connect")
	symaddr(&addr_QLibrary_connect2, "QLibrary_connect2")
	symaddr(&addr_QLibrary_sender, "QLibrary_sender")
	symaddr(&addr_QLibrary_senderSignalIndex, "QLibrary_senderSignalIndex")
	symaddr(&addr_QLibrary_receivers, "QLibrary_receivers")
	symaddr(&addr_QLibrary_isSignalConnected, "QLibrary_isSignalConnected")
	symaddr(&addr_QLibrary_tr2, "QLibrary_tr2")
	symaddr(&addr_QLibrary_tr3, "QLibrary_tr3")
	symaddr(&addr_QLibrary_connect3, "QLibrary_connect3")
	symaddr(&addr_QLibrary_connect4, "QLibrary_connect4")
	symaddr(&addr_QLibrary_metaObject_default, "QLibrary_metaObject_default")
	symaddr(&addr_QLibrary_event_default, "QLibrary_event_default")
	symaddr(&addr_QLibrary_eventFilter_default, "QLibrary_eventFilter_default")
	symaddr(&addr_QLibrary_timerEvent_default, "QLibrary_timerEvent_default")
	symaddr(&addr_QLibrary_childEvent_default, "QLibrary_childEvent_default")
	symaddr(&addr_QLibrary_customEvent_default, "QLibrary_customEvent_default")
	symaddr(&addr_QLibrary_connectNotify_default, "QLibrary_connectNotify_default")
	symaddr(&addr_QLibrary_disconnectNotify_default, "QLibrary_disconnectNotify_default")
	reg(&pg_QLibrary_connect_destroyed, "QLibrary_connect_destroyed")
	reg(&pg_QLibrary_connect_destroyed2, "QLibrary_connect_destroyed2")
	reg(&pg_QLibrary_connect_objectNameChanged, "QLibrary_connect_objectNameChanged")
	reg(&pg_QLibrary__disconnect, "QLibrary__disconnect")
	reg(&pg_QLibrary__set_override, "QLibrary__set_override")
	reg(&pg_QLibrary_delete, "QLibrary_delete")
}

type QLibrary struct {
	h uintptr
	*QObject
}

func bindQLibrary(h uintptr) *QLibrary {
	if h == 0 {
		return nil
	}
	return &QLibrary{h: h, QObject: bindQObject(pg_QLibrary_base(h))}
}

func newownedQLibrary(h uintptr) *QLibrary { return bindQLibrary(h) }

func (this *QLibrary) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLibrary) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLibrary(h unsafe.Pointer) *QLibrary { return bindQLibrary(uintptr(h)) }
func UnsafeNewQLibraryOwned(h unsafe.Pointer) *QLibrary { return newownedQLibrary(uintptr(h)) }

func (this *QLibrary) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQLibrary() *QLibrary {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLibrary_new, uintptr(__goid))
	__o := newownedQLibrary(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary2(fileName string) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new2, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary3(fileName string, verNum int) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new3, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(verNum))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary4(fileName string, version string) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new4, uintptr(__goid), fileName_p, uintptr(len(fileName)), version_p, uintptr(len(version)))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(version)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary5(parent *QObject) *QLibrary {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLibrary_new5, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary6(fileName string, parent *QObject) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new6, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(parent.UnsafePointer()))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary7(fileName string, verNum int, parent *QObject) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new7, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(verNum), uintptr(parent.UnsafePointer()))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func NewQLibrary8(fileName string, version string, parent *QObject) *QLibrary {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	__r, _, _ := purego.SyscallN(addr_QLibrary_new8, uintptr(__goid), fileName_p, uintptr(len(fileName)), version_p, uintptr(len(version)), uintptr(parent.UnsafePointer()))
	__o := newownedQLibrary(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(version)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLibrary_virtIdx, QLibrary_enable)
	return __o
}

func QLibrary_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibrary_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QLibrary) Load() bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_load, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLibrary) Unload() bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_unload, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLibrary) IsLoaded() bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_isLoaded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QLibrary_IsLibrary(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QLibrary_isLibrary, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QLibrary) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QLibrary_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QLibrary) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibrary_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLibrary) SetFileNameAndVersion(fileName string, verNum int) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QLibrary_setFileNameAndVersion, this.h, fileName_p, uintptr(len(fileName)), uintptr(verNum))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QLibrary) SetFileNameAndVersion2(fileName string, version string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QLibrary_setFileNameAndVersion2, this.h, fileName_p, uintptr(len(fileName)), version_p, uintptr(len(version)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(version)
}

func (this *QLibrary) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibrary_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLibrary) SetLoadHints(hints QLibrary__LoadHint) {
	purego.SyscallN(addr_QLibrary_setLoadHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QLibrary) LoadHints() QLibrary__LoadHint {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_loadHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLibrary__LoadHint(__r)
}

func QLibrary_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLibrary_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLibrary_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLibrary_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLibrary) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QLibrary_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QLibrary) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLibrary) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QLibrary_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QLibrary) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QLibrary_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibrary_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLibrary_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLibrary_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLibrary_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLibrary_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLibrary_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLibrary_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLibrary) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QLibrary_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QLibrary) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLibrary) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLibrary_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLibrary) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QLibrary_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLibrary) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QLibrary_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLibrary) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLibrary_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLibrary) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLibrary_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLibrary) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLibrary_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLibrary) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLibrary_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLibrary) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QLibrary_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLibrary) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QLibrary_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLibrary) Disconnect(handle int64) {
	pg_QLibrary__disconnect(handle)
	unregisterCallback(handle)
}

func QLibrary_virtIdx(slot int64) int32 {
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

func QLibrary_enable(h uintptr, idx int32) { pg_QLibrary__set_override(h, idx) }

func (this *QLibrary) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QLibrary) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QLibrary) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QLibrary) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QLibrary) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QLibrary) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QLibrary) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QLibrary) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QLibrary) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLibrary_delete(this.h)
	}
	this.UnsafeRelease()
}
