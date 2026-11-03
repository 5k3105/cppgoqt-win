package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFileSystemWatcher_base func(uintptr) uintptr
	addr_QFileSystemWatcher_new uintptr
	addr_QFileSystemWatcher_new2 uintptr
	addr_QFileSystemWatcher_new3 uintptr
	addr_QFileSystemWatcher_new4 uintptr
	addr_QFileSystemWatcher_tr uintptr
	addr_QFileSystemWatcher_addPath uintptr
	addr_QFileSystemWatcher_addPaths uintptr
	addr_QFileSystemWatcher_removePath uintptr
	addr_QFileSystemWatcher_removePaths uintptr
	addr_QFileSystemWatcher_files uintptr
	addr_QFileSystemWatcher_directories uintptr
	addr_QFileSystemWatcher_connect uintptr
	addr_QFileSystemWatcher_connect2 uintptr
	addr_QFileSystemWatcher_sender uintptr
	addr_QFileSystemWatcher_senderSignalIndex uintptr
	addr_QFileSystemWatcher_receivers uintptr
	addr_QFileSystemWatcher_isSignalConnected uintptr
	addr_QFileSystemWatcher_tr2 uintptr
	addr_QFileSystemWatcher_tr3 uintptr
	addr_QFileSystemWatcher_connect3 uintptr
	addr_QFileSystemWatcher_connect4 uintptr
	addr_QFileSystemWatcher_metaObject_default uintptr
	addr_QFileSystemWatcher_event_default uintptr
	addr_QFileSystemWatcher_eventFilter_default uintptr
	addr_QFileSystemWatcher_timerEvent_default uintptr
	addr_QFileSystemWatcher_childEvent_default uintptr
	addr_QFileSystemWatcher_customEvent_default uintptr
	addr_QFileSystemWatcher_connectNotify_default uintptr
	addr_QFileSystemWatcher_disconnectNotify_default uintptr
	pg_QFileSystemWatcher_connect_fileChanged func(uintptr, int64)
	pg_QFileSystemWatcher_connect_directoryChanged func(uintptr, int64)
	pg_QFileSystemWatcher_connect_destroyed func(uintptr, int64)
	pg_QFileSystemWatcher_connect_destroyed2 func(uintptr, int64)
	pg_QFileSystemWatcher_connect_objectNameChanged func(uintptr, int64)
	pg_QFileSystemWatcher__disconnect func(int64)
	pg_QFileSystemWatcher__set_override func(uintptr, int32)
	pg_QFileSystemWatcher_delete func(uintptr)
)

func init() {
	reg(&pg_QFileSystemWatcher_base, "QFileSystemWatcher_base")
	symaddr(&addr_QFileSystemWatcher_new, "QFileSystemWatcher_new")
	symaddr(&addr_QFileSystemWatcher_new2, "QFileSystemWatcher_new2")
	symaddr(&addr_QFileSystemWatcher_new3, "QFileSystemWatcher_new3")
	symaddr(&addr_QFileSystemWatcher_new4, "QFileSystemWatcher_new4")
	symaddr(&addr_QFileSystemWatcher_tr, "QFileSystemWatcher_tr")
	symaddr(&addr_QFileSystemWatcher_addPath, "QFileSystemWatcher_addPath")
	symaddr(&addr_QFileSystemWatcher_addPaths, "QFileSystemWatcher_addPaths")
	symaddr(&addr_QFileSystemWatcher_removePath, "QFileSystemWatcher_removePath")
	symaddr(&addr_QFileSystemWatcher_removePaths, "QFileSystemWatcher_removePaths")
	symaddr(&addr_QFileSystemWatcher_files, "QFileSystemWatcher_files")
	symaddr(&addr_QFileSystemWatcher_directories, "QFileSystemWatcher_directories")
	symaddr(&addr_QFileSystemWatcher_connect, "QFileSystemWatcher_connect")
	symaddr(&addr_QFileSystemWatcher_connect2, "QFileSystemWatcher_connect2")
	symaddr(&addr_QFileSystemWatcher_sender, "QFileSystemWatcher_sender")
	symaddr(&addr_QFileSystemWatcher_senderSignalIndex, "QFileSystemWatcher_senderSignalIndex")
	symaddr(&addr_QFileSystemWatcher_receivers, "QFileSystemWatcher_receivers")
	symaddr(&addr_QFileSystemWatcher_isSignalConnected, "QFileSystemWatcher_isSignalConnected")
	symaddr(&addr_QFileSystemWatcher_tr2, "QFileSystemWatcher_tr2")
	symaddr(&addr_QFileSystemWatcher_tr3, "QFileSystemWatcher_tr3")
	symaddr(&addr_QFileSystemWatcher_connect3, "QFileSystemWatcher_connect3")
	symaddr(&addr_QFileSystemWatcher_connect4, "QFileSystemWatcher_connect4")
	symaddr(&addr_QFileSystemWatcher_metaObject_default, "QFileSystemWatcher_metaObject_default")
	symaddr(&addr_QFileSystemWatcher_event_default, "QFileSystemWatcher_event_default")
	symaddr(&addr_QFileSystemWatcher_eventFilter_default, "QFileSystemWatcher_eventFilter_default")
	symaddr(&addr_QFileSystemWatcher_timerEvent_default, "QFileSystemWatcher_timerEvent_default")
	symaddr(&addr_QFileSystemWatcher_childEvent_default, "QFileSystemWatcher_childEvent_default")
	symaddr(&addr_QFileSystemWatcher_customEvent_default, "QFileSystemWatcher_customEvent_default")
	symaddr(&addr_QFileSystemWatcher_connectNotify_default, "QFileSystemWatcher_connectNotify_default")
	symaddr(&addr_QFileSystemWatcher_disconnectNotify_default, "QFileSystemWatcher_disconnectNotify_default")
	reg(&pg_QFileSystemWatcher_connect_fileChanged, "QFileSystemWatcher_connect_fileChanged")
	reg(&pg_QFileSystemWatcher_connect_directoryChanged, "QFileSystemWatcher_connect_directoryChanged")
	reg(&pg_QFileSystemWatcher_connect_destroyed, "QFileSystemWatcher_connect_destroyed")
	reg(&pg_QFileSystemWatcher_connect_destroyed2, "QFileSystemWatcher_connect_destroyed2")
	reg(&pg_QFileSystemWatcher_connect_objectNameChanged, "QFileSystemWatcher_connect_objectNameChanged")
	reg(&pg_QFileSystemWatcher__disconnect, "QFileSystemWatcher__disconnect")
	reg(&pg_QFileSystemWatcher__set_override, "QFileSystemWatcher__set_override")
	reg(&pg_QFileSystemWatcher_delete, "QFileSystemWatcher_delete")
}

type QFileSystemWatcher struct {
	h uintptr
	*QObject
}

func bindQFileSystemWatcher(h uintptr) *QFileSystemWatcher {
	if h == 0 {
		return nil
	}
	return &QFileSystemWatcher{h: h, QObject: bindQObject(pg_QFileSystemWatcher_base(h))}
}

func newownedQFileSystemWatcher(h uintptr) *QFileSystemWatcher { return bindQFileSystemWatcher(h) }

func (this *QFileSystemWatcher) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileSystemWatcher) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileSystemWatcher(h unsafe.Pointer) *QFileSystemWatcher { return bindQFileSystemWatcher(uintptr(h)) }
func UnsafeNewQFileSystemWatcherOwned(h unsafe.Pointer) *QFileSystemWatcher { return newownedQFileSystemWatcher(uintptr(h)) }

func (this *QFileSystemWatcher) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQFileSystemWatcher() *QFileSystemWatcher {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_new, uintptr(__goid))
	__o := newownedQFileSystemWatcher(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemWatcher_virtIdx, QFileSystemWatcher_enable)
	return __o
}

func NewQFileSystemWatcher2(paths []string) *QFileSystemWatcher {
	__goid := nextVirtID()
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_new2, uintptr(__goid), paths_p, uintptr(len(paths)))
	__o := newownedQFileSystemWatcher(__r)
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemWatcher_virtIdx, QFileSystemWatcher_enable)
	return __o
}

func NewQFileSystemWatcher3(parent *QObject) *QFileSystemWatcher {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFileSystemWatcher(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemWatcher_virtIdx, QFileSystemWatcher_enable)
	return __o
}

func NewQFileSystemWatcher4(paths []string, parent *QObject) *QFileSystemWatcher {
	__goid := nextVirtID()
	paths_arr := make([]uintptr, len(paths))
	for __i, __e := range paths {
		paths_arr[__i] = pgCString(__e)
	}
	var paths_p uintptr
	if len(paths_arr) > 0 {
		paths_p = uintptr(unsafe.Pointer(&paths_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_new4, uintptr(__goid), paths_p, uintptr(len(paths)), uintptr(parent.UnsafePointer()))
	__o := newownedQFileSystemWatcher(__r)
	for __i := range paths_arr {
		pgFree(paths_arr[__i])
	}
	runtime.KeepAlive(paths_arr)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemWatcher_virtIdx, QFileSystemWatcher_enable)
	return __o
}

func QFileSystemWatcher_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFileSystemWatcher) AddPath(file string) bool {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_addPath, this.h, file_p, uintptr(len(file)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
	return __r != 0
}

func (this *QFileSystemWatcher) AddPaths(files []string) []string {
	files_arr := make([]uintptr, len(files))
	for __i, __e := range files {
		files_arr[__i] = pgCString(__e)
	}
	var files_p uintptr
	if len(files_arr) > 0 {
		files_p = uintptr(unsafe.Pointer(&files_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_addPaths, this.h, files_p, uintptr(len(files)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range files_arr {
		pgFree(files_arr[__i])
	}
	runtime.KeepAlive(files_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileSystemWatcher) RemovePath(file string) bool {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_removePath, this.h, file_p, uintptr(len(file)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
	return __r != 0
}

func (this *QFileSystemWatcher) RemovePaths(files []string) []string {
	files_arr := make([]uintptr, len(files))
	for __i, __e := range files {
		files_arr[__i] = pgCString(__e)
	}
	var files_p uintptr
	if len(files_arr) > 0 {
		files_p = uintptr(unsafe.Pointer(&files_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_removePaths, this.h, files_p, uintptr(len(files)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	for __i := range files_arr {
		pgFree(files_arr[__i])
	}
	runtime.KeepAlive(files_arr)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileSystemWatcher) Files() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_files, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileSystemWatcher) Directories() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_directories, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileSystemWatcher_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSystemWatcher_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSystemWatcher) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFileSystemWatcher) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileSystemWatcher) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFileSystemWatcher) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFileSystemWatcher_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileSystemWatcher_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileSystemWatcher_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSystemWatcher_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSystemWatcher) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFileSystemWatcher_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFileSystemWatcher) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSystemWatcher) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemWatcher_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSystemWatcher) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFileSystemWatcher_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemWatcher) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFileSystemWatcher_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemWatcher) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFileSystemWatcher_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemWatcher) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSystemWatcher_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSystemWatcher) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSystemWatcher_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSystemWatcher) OnFileChanged(fn func(string)) int64 {
	id := registerCallback(func(path uintptr, path_len int64) { fn(pgStrNoFree(path, path_len)) })
	pg_QFileSystemWatcher_connect_fileChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemWatcher) OnDirectoryChanged(fn func(string)) int64 {
	id := registerCallback(func(path uintptr, path_len int64) { fn(pgStrNoFree(path, path_len)) })
	pg_QFileSystemWatcher_connect_directoryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemWatcher) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemWatcher_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemWatcher) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFileSystemWatcher_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemWatcher) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFileSystemWatcher_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemWatcher) Disconnect(handle int64) {
	pg_QFileSystemWatcher__disconnect(handle)
	unregisterCallback(handle)
}

func QFileSystemWatcher_virtIdx(slot int64) int32 {
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

func QFileSystemWatcher_enable(h uintptr, idx int32) { pg_QFileSystemWatcher__set_override(h, idx) }

func (this *QFileSystemWatcher) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFileSystemWatcher) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFileSystemWatcher) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFileSystemWatcher) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFileSystemWatcher) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFileSystemWatcher) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFileSystemWatcher) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFileSystemWatcher) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFileSystemWatcher) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileSystemWatcher_delete(this.h)
	}
	this.UnsafeRelease()
}
