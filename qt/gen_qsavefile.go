package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSaveFile_chain uintptr
	addr_QSaveFile_new uintptr
	addr_QSaveFile_new2 uintptr
	addr_QSaveFile_new3 uintptr
	addr_QSaveFile_new4 uintptr
	addr_QSaveFile_tr uintptr
	addr_QSaveFile_setFileName uintptr
	addr_QSaveFile_commit uintptr
	addr_QSaveFile_cancelWriting uintptr
	addr_QSaveFile_setDirectWriteFallback uintptr
	addr_QSaveFile_directWriteFallback uintptr
	addr_QSaveFile_setOpenMode uintptr
	addr_QSaveFile_setErrorString uintptr
	addr_QSaveFile_connect uintptr
	addr_QSaveFile_connect2 uintptr
	addr_QSaveFile_sender uintptr
	addr_QSaveFile_senderSignalIndex uintptr
	addr_QSaveFile_receivers uintptr
	addr_QSaveFile_isSignalConnected uintptr
	addr_QSaveFile_tr2 uintptr
	addr_QSaveFile_tr3 uintptr
	addr_QSaveFile_connect3 uintptr
	addr_QSaveFile_connect4 uintptr
	addr_QSaveFile_metaObject_default uintptr
	addr_QSaveFile_fileName_default uintptr
	addr_QSaveFile_open_default uintptr
	addr_QSaveFile_isSequential_default uintptr
	addr_QSaveFile_pos_default uintptr
	addr_QSaveFile_seek_default uintptr
	addr_QSaveFile_atEnd_default uintptr
	addr_QSaveFile_size_default uintptr
	addr_QSaveFile_resize_default uintptr
	addr_QSaveFile_permissions_default uintptr
	addr_QSaveFile_setPermissions_default uintptr
	addr_QSaveFile_reset_default uintptr
	addr_QSaveFile_bytesAvailable_default uintptr
	addr_QSaveFile_bytesToWrite_default uintptr
	addr_QSaveFile_canReadLine_default uintptr
	addr_QSaveFile_waitForReadyRead_default uintptr
	addr_QSaveFile_waitForBytesWritten_default uintptr
	addr_QSaveFile_skipData_default uintptr
	addr_QSaveFile_event_default uintptr
	addr_QSaveFile_eventFilter_default uintptr
	addr_QSaveFile_timerEvent_default uintptr
	addr_QSaveFile_childEvent_default uintptr
	addr_QSaveFile_customEvent_default uintptr
	addr_QSaveFile_connectNotify_default uintptr
	addr_QSaveFile_disconnectNotify_default uintptr
	pg_QSaveFile_connect_readyRead func(uintptr, int64)
	pg_QSaveFile_connect_channelReadyRead func(uintptr, int64)
	pg_QSaveFile_connect_bytesWritten func(uintptr, int64)
	pg_QSaveFile_connect_channelBytesWritten func(uintptr, int64)
	pg_QSaveFile_connect_aboutToClose func(uintptr, int64)
	pg_QSaveFile_connect_readChannelFinished func(uintptr, int64)
	pg_QSaveFile_connect_destroyed func(uintptr, int64)
	pg_QSaveFile_connect_destroyed2 func(uintptr, int64)
	pg_QSaveFile_connect_objectNameChanged func(uintptr, int64)
	pg_QSaveFile__disconnect func(int64)
	pg_QSaveFile__set_override func(uintptr, int32)
	pg_QSaveFile_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSaveFile_chain, "QSaveFile_chain")
	symaddr(&addr_QSaveFile_new, "QSaveFile_new")
	symaddr(&addr_QSaveFile_new2, "QSaveFile_new2")
	symaddr(&addr_QSaveFile_new3, "QSaveFile_new3")
	symaddr(&addr_QSaveFile_new4, "QSaveFile_new4")
	symaddr(&addr_QSaveFile_tr, "QSaveFile_tr")
	symaddr(&addr_QSaveFile_setFileName, "QSaveFile_setFileName")
	symaddr(&addr_QSaveFile_commit, "QSaveFile_commit")
	symaddr(&addr_QSaveFile_cancelWriting, "QSaveFile_cancelWriting")
	symaddr(&addr_QSaveFile_setDirectWriteFallback, "QSaveFile_setDirectWriteFallback")
	symaddr(&addr_QSaveFile_directWriteFallback, "QSaveFile_directWriteFallback")
	symaddr(&addr_QSaveFile_setOpenMode, "QSaveFile_setOpenMode")
	symaddr(&addr_QSaveFile_setErrorString, "QSaveFile_setErrorString")
	symaddr(&addr_QSaveFile_connect, "QSaveFile_connect")
	symaddr(&addr_QSaveFile_connect2, "QSaveFile_connect2")
	symaddr(&addr_QSaveFile_sender, "QSaveFile_sender")
	symaddr(&addr_QSaveFile_senderSignalIndex, "QSaveFile_senderSignalIndex")
	symaddr(&addr_QSaveFile_receivers, "QSaveFile_receivers")
	symaddr(&addr_QSaveFile_isSignalConnected, "QSaveFile_isSignalConnected")
	symaddr(&addr_QSaveFile_tr2, "QSaveFile_tr2")
	symaddr(&addr_QSaveFile_tr3, "QSaveFile_tr3")
	symaddr(&addr_QSaveFile_connect3, "QSaveFile_connect3")
	symaddr(&addr_QSaveFile_connect4, "QSaveFile_connect4")
	symaddr(&addr_QSaveFile_metaObject_default, "QSaveFile_metaObject_default")
	symaddr(&addr_QSaveFile_fileName_default, "QSaveFile_fileName_default")
	symaddr(&addr_QSaveFile_open_default, "QSaveFile_open_default")
	symaddr(&addr_QSaveFile_isSequential_default, "QSaveFile_isSequential_default")
	symaddr(&addr_QSaveFile_pos_default, "QSaveFile_pos_default")
	symaddr(&addr_QSaveFile_seek_default, "QSaveFile_seek_default")
	symaddr(&addr_QSaveFile_atEnd_default, "QSaveFile_atEnd_default")
	symaddr(&addr_QSaveFile_size_default, "QSaveFile_size_default")
	symaddr(&addr_QSaveFile_resize_default, "QSaveFile_resize_default")
	symaddr(&addr_QSaveFile_permissions_default, "QSaveFile_permissions_default")
	symaddr(&addr_QSaveFile_setPermissions_default, "QSaveFile_setPermissions_default")
	symaddr(&addr_QSaveFile_reset_default, "QSaveFile_reset_default")
	symaddr(&addr_QSaveFile_bytesAvailable_default, "QSaveFile_bytesAvailable_default")
	symaddr(&addr_QSaveFile_bytesToWrite_default, "QSaveFile_bytesToWrite_default")
	symaddr(&addr_QSaveFile_canReadLine_default, "QSaveFile_canReadLine_default")
	symaddr(&addr_QSaveFile_waitForReadyRead_default, "QSaveFile_waitForReadyRead_default")
	symaddr(&addr_QSaveFile_waitForBytesWritten_default, "QSaveFile_waitForBytesWritten_default")
	symaddr(&addr_QSaveFile_skipData_default, "QSaveFile_skipData_default")
	symaddr(&addr_QSaveFile_event_default, "QSaveFile_event_default")
	symaddr(&addr_QSaveFile_eventFilter_default, "QSaveFile_eventFilter_default")
	symaddr(&addr_QSaveFile_timerEvent_default, "QSaveFile_timerEvent_default")
	symaddr(&addr_QSaveFile_childEvent_default, "QSaveFile_childEvent_default")
	symaddr(&addr_QSaveFile_customEvent_default, "QSaveFile_customEvent_default")
	symaddr(&addr_QSaveFile_connectNotify_default, "QSaveFile_connectNotify_default")
	symaddr(&addr_QSaveFile_disconnectNotify_default, "QSaveFile_disconnectNotify_default")
	reg(&pg_QSaveFile_connect_readyRead, "QSaveFile_connect_readyRead")
	reg(&pg_QSaveFile_connect_channelReadyRead, "QSaveFile_connect_channelReadyRead")
	reg(&pg_QSaveFile_connect_bytesWritten, "QSaveFile_connect_bytesWritten")
	reg(&pg_QSaveFile_connect_channelBytesWritten, "QSaveFile_connect_channelBytesWritten")
	reg(&pg_QSaveFile_connect_aboutToClose, "QSaveFile_connect_aboutToClose")
	reg(&pg_QSaveFile_connect_readChannelFinished, "QSaveFile_connect_readChannelFinished")
	reg(&pg_QSaveFile_connect_destroyed, "QSaveFile_connect_destroyed")
	reg(&pg_QSaveFile_connect_destroyed2, "QSaveFile_connect_destroyed2")
	reg(&pg_QSaveFile_connect_objectNameChanged, "QSaveFile_connect_objectNameChanged")
	reg(&pg_QSaveFile__disconnect, "QSaveFile__disconnect")
	reg(&pg_QSaveFile__set_override, "QSaveFile__set_override")
	reg(&pg_QSaveFile_delete, "QSaveFile_delete")
}

type QSaveFile struct {
	h uintptr
	*QFileDevice
}

func bindQSaveFile(h uintptr) *QSaveFile {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSaveFile_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSaveFile{h: h, QFileDevice: &QFileDevice{h: __c[0], QIODevice: &QIODevice{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSaveFile(h uintptr) *QSaveFile { return bindQSaveFile(h) }

func (this *QSaveFile) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSaveFile) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSaveFile(h unsafe.Pointer) *QSaveFile { return bindQSaveFile(uintptr(h)) }
func UnsafeNewQSaveFileOwned(h unsafe.Pointer) *QSaveFile { return newownedQSaveFile(uintptr(h)) }

func (this *QSaveFile) UnsafeRelease() {
	this.h = 0
	if this.QFileDevice != nil {
		this.QFileDevice.UnsafeRelease()
	}
}

func NewQSaveFile() *QSaveFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSaveFile_new, uintptr(__goid))
	__o := newownedQSaveFile(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSaveFile_virtIdx, QSaveFile_enable)
	return __o
}

func NewQSaveFile2(name string) *QSaveFile {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QSaveFile_new2, uintptr(__goid), name_p, uintptr(len(name)))
	__o := newownedQSaveFile(__r)
	runtime.KeepAlive(name)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSaveFile_virtIdx, QSaveFile_enable)
	return __o
}

func NewQSaveFile3(parent *QObject) *QSaveFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSaveFile_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSaveFile(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSaveFile_virtIdx, QSaveFile_enable)
	return __o
}

func NewQSaveFile4(name string, parent *QObject) *QSaveFile {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QSaveFile_new4, uintptr(__goid), name_p, uintptr(len(name)), uintptr(parent.UnsafePointer()))
	__o := newownedQSaveFile(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSaveFile_virtIdx, QSaveFile_enable)
	return __o
}

func QSaveFile_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSaveFile_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSaveFile) SetFileName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QSaveFile_setFileName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QSaveFile) Commit() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_commit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) CancelWriting() {
	purego.SyscallN(addr_QSaveFile_cancelWriting, this.h)
	runtime.KeepAlive(this)
}

func (this *QSaveFile) SetDirectWriteFallback(enabled bool) {
	purego.SyscallN(addr_QSaveFile_setDirectWriteFallback, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QSaveFile) DirectWriteFallback() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_directWriteFallback, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) SetOpenMode(openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QSaveFile_setOpenMode, this.h, uintptr(openMode))
	runtime.KeepAlive(this)
}

func (this *QSaveFile) SetErrorString(errorString string) {
	var errorString_p uintptr
	if len(errorString) > 0 {
		errorString_p = uintptr(unsafe.Pointer(unsafe.StringData(errorString)))
	}
	purego.SyscallN(addr_QSaveFile_setErrorString, this.h, errorString_p, uintptr(len(errorString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(errorString)
}

func QSaveFile_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSaveFile_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSaveFile_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSaveFile_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSaveFile) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSaveFile_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSaveFile) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSaveFile) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSaveFile) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSaveFile_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSaveFile_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSaveFile_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSaveFile_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSaveFile_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSaveFile_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSaveFile_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSaveFile_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSaveFile) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSaveFile_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSaveFile) FileNameDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSaveFile_fileName_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSaveFile) OpenDefault(flags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_open_default, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) IsSequentialDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_isSequential_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) PosDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_pos_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSaveFile) SeekDefault(offset int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_seek_default, this.h, uintptr(offset))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) AtEndDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_atEnd_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) SizeDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_size_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSaveFile) ResizeDefault(sz int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_resize_default, this.h, uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) PermissionsDefault() QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_permissions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__Permission(__r)
}

func (this *QSaveFile) SetPermissionsDefault(permissionSpec QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_setPermissions_default, this.h, uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) ResetDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_reset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) BytesAvailableDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_bytesAvailable_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSaveFile) BytesToWriteDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_bytesToWrite_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSaveFile) CanReadLineDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_canReadLine_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) WaitForReadyReadDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_waitForReadyRead_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) WaitForBytesWrittenDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_waitForBytesWritten_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSaveFile) SkipDataDefault(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_skipData_default, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSaveFile) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSaveFile) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSaveFile_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSaveFile) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSaveFile_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSaveFile) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSaveFile_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSaveFile) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSaveFile_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSaveFile) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSaveFile_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSaveFile) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSaveFile_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSaveFile) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSaveFile_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QSaveFile_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QSaveFile_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QSaveFile_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSaveFile_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSaveFile_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSaveFile_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSaveFile_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSaveFile_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSaveFile) Disconnect(handle int64) {
	pg_QSaveFile__disconnect(handle)
	unregisterCallback(handle)
}

func QSaveFile_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 8848495542063941256:
		return 1
	case -554057428128861975:
		return 2
	case 8888098400087344862:
		return 3
	case 8618582476283917417:
		return 4
	case 3241140482992819217:
		return 5
	case -1487528724786208339:
		return 6
	case 5614464919154503228:
		return 7
	case -4890198374319873567:
		return 8
	case -7517927236947282013:
		return 9
	case 420407675554349589:
		return 10
	case 1086335023529244512:
		return 11
	case 4814449630054007943:
		return 12
	case 1024747515116322948:
		return 13
	case -1250682508954504761:
		return 14
	case 8356976193102722774:
		return 15
	case 7726833620054767889:
		return 16
	case 8178738247667652082:
		return 17
	case 1237236456312824991:
		return 18
	case 5217623514630858297:
		return 19
	case -5409927581812374498:
		return 20
	case -1185315367526857833:
		return 21
	case 3769868229123868474:
		return 22
	case 7379029982170503008:
		return 23
	case -568369772982215962:
		return 24
	}
	return -1
}

func QSaveFile_enable(h uintptr, idx int32) { pg_QSaveFile__set_override(h, idx) }

func (this *QSaveFile) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSaveFile) OnFileName(fn func() string) {
	registerVirt(this.UnsafeGoID(), 8848495542063941256, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "fileName"
}

func (this *QSaveFile) OnOpen(fn func(QIODeviceBase__OpenModeFlag) bool) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func(flags int32) int32 { return b2i(fn(QIODeviceBase__OpenModeFlag(flags))) }) // slot "open"
}

func (this *QSaveFile) OnIsSequential(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8888098400087344862, func() int32 { return b2i(fn()) }) // slot "isSequential"
}

func (this *QSaveFile) OnPos(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 8618582476283917417, func() int64 { return fn() }) // slot "pos"
}

func (this *QSaveFile) OnSeek(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), 3241140482992819217, func(offset int64) int32 { return b2i(fn(offset)) }) // slot "seek"
}

func (this *QSaveFile) OnAtEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1487528724786208339, func() int32 { return b2i(fn()) }) // slot "atEnd"
}

func (this *QSaveFile) OnSize(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() int64 { return fn() }) // slot "size"
}

func (this *QSaveFile) OnResize(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), -4890198374319873567, func(sz int64) int32 { return b2i(fn(sz)) }) // slot "resize"
}

func (this *QSaveFile) OnPermissions(fn func() QFileDevice__Permission) {
	registerVirt(this.UnsafeGoID(), -7517927236947282013, func() int32 { return int32(fn()) }) // slot "permissions"
}

func (this *QSaveFile) OnSetPermissions(fn func(QFileDevice__Permission) bool) {
	registerVirt(this.UnsafeGoID(), 420407675554349589, func(permissionSpec int32) int32 { return b2i(fn(QFileDevice__Permission(permissionSpec))) }) // slot "setPermissions"
}

func (this *QSaveFile) OnReset(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() int32 { return b2i(fn()) }) // slot "reset"
}

func (this *QSaveFile) OnBytesAvailable(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 4814449630054007943, func() int64 { return fn() }) // slot "bytesAvailable"
}

func (this *QSaveFile) OnBytesToWrite(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 1024747515116322948, func() int64 { return fn() }) // slot "bytesToWrite"
}

func (this *QSaveFile) OnCanReadLine(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1250682508954504761, func() int32 { return b2i(fn()) }) // slot "canReadLine"
}

func (this *QSaveFile) OnWaitForReadyRead(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 8356976193102722774, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForReadyRead"
}

func (this *QSaveFile) OnWaitForBytesWritten(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 7726833620054767889, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForBytesWritten"
}

func (this *QSaveFile) OnSkipData(fn func(int64) int64) {
	registerVirt(this.UnsafeGoID(), 8178738247667652082, func(maxSize int64) int64 { return fn(maxSize) }) // slot "skipData"
}

func (this *QSaveFile) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSaveFile) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSaveFile) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSaveFile) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSaveFile) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSaveFile) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSaveFile) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSaveFile) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSaveFile_delete(this.h)
	}
	this.UnsafeRelease()
}
