package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBuffer_chain uintptr
	addr_QBuffer_new uintptr
	addr_QBuffer_new2 uintptr
	addr_QBuffer_new3 uintptr
	addr_QBuffer_new4 uintptr
	addr_QBuffer_tr uintptr
	addr_QBuffer_buffer uintptr
	addr_QBuffer_buffer2 uintptr
	addr_QBuffer_setBuffer uintptr
	addr_QBuffer_setData uintptr
	addr_QBuffer_setData2 uintptr
	addr_QBuffer_data uintptr
	addr_QBuffer_setOpenMode uintptr
	addr_QBuffer_setErrorString uintptr
	addr_QBuffer_connect uintptr
	addr_QBuffer_connect2 uintptr
	addr_QBuffer_sender uintptr
	addr_QBuffer_senderSignalIndex uintptr
	addr_QBuffer_receivers uintptr
	addr_QBuffer_isSignalConnected uintptr
	addr_QBuffer_tr2 uintptr
	addr_QBuffer_tr3 uintptr
	addr_QBuffer_connect3 uintptr
	addr_QBuffer_connect4 uintptr
	addr_QBuffer_metaObject_default uintptr
	addr_QBuffer_open_default uintptr
	addr_QBuffer_close_default uintptr
	addr_QBuffer_size_default uintptr
	addr_QBuffer_pos_default uintptr
	addr_QBuffer_seek_default uintptr
	addr_QBuffer_atEnd_default uintptr
	addr_QBuffer_canReadLine_default uintptr
	addr_QBuffer_connectNotify_default uintptr
	addr_QBuffer_disconnectNotify_default uintptr
	addr_QBuffer_isSequential_default uintptr
	addr_QBuffer_reset_default uintptr
	addr_QBuffer_bytesAvailable_default uintptr
	addr_QBuffer_bytesToWrite_default uintptr
	addr_QBuffer_waitForReadyRead_default uintptr
	addr_QBuffer_waitForBytesWritten_default uintptr
	addr_QBuffer_skipData_default uintptr
	addr_QBuffer_event_default uintptr
	addr_QBuffer_eventFilter_default uintptr
	addr_QBuffer_timerEvent_default uintptr
	addr_QBuffer_childEvent_default uintptr
	addr_QBuffer_customEvent_default uintptr
	pg_QBuffer_connect_readyRead func(uintptr, int64)
	pg_QBuffer_connect_channelReadyRead func(uintptr, int64)
	pg_QBuffer_connect_bytesWritten func(uintptr, int64)
	pg_QBuffer_connect_channelBytesWritten func(uintptr, int64)
	pg_QBuffer_connect_aboutToClose func(uintptr, int64)
	pg_QBuffer_connect_readChannelFinished func(uintptr, int64)
	pg_QBuffer_connect_destroyed func(uintptr, int64)
	pg_QBuffer_connect_destroyed2 func(uintptr, int64)
	pg_QBuffer_connect_objectNameChanged func(uintptr, int64)
	pg_QBuffer__disconnect func(int64)
	pg_QBuffer__set_override func(uintptr, int32)
	pg_QBuffer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBuffer_chain, "QBuffer_chain")
	symaddr(&addr_QBuffer_new, "QBuffer_new")
	symaddr(&addr_QBuffer_new2, "QBuffer_new2")
	symaddr(&addr_QBuffer_new3, "QBuffer_new3")
	symaddr(&addr_QBuffer_new4, "QBuffer_new4")
	symaddr(&addr_QBuffer_tr, "QBuffer_tr")
	symaddr(&addr_QBuffer_buffer, "QBuffer_buffer")
	symaddr(&addr_QBuffer_buffer2, "QBuffer_buffer2")
	symaddr(&addr_QBuffer_setBuffer, "QBuffer_setBuffer")
	symaddr(&addr_QBuffer_setData, "QBuffer_setData")
	symaddr(&addr_QBuffer_setData2, "QBuffer_setData2")
	symaddr(&addr_QBuffer_data, "QBuffer_data")
	symaddr(&addr_QBuffer_setOpenMode, "QBuffer_setOpenMode")
	symaddr(&addr_QBuffer_setErrorString, "QBuffer_setErrorString")
	symaddr(&addr_QBuffer_connect, "QBuffer_connect")
	symaddr(&addr_QBuffer_connect2, "QBuffer_connect2")
	symaddr(&addr_QBuffer_sender, "QBuffer_sender")
	symaddr(&addr_QBuffer_senderSignalIndex, "QBuffer_senderSignalIndex")
	symaddr(&addr_QBuffer_receivers, "QBuffer_receivers")
	symaddr(&addr_QBuffer_isSignalConnected, "QBuffer_isSignalConnected")
	symaddr(&addr_QBuffer_tr2, "QBuffer_tr2")
	symaddr(&addr_QBuffer_tr3, "QBuffer_tr3")
	symaddr(&addr_QBuffer_connect3, "QBuffer_connect3")
	symaddr(&addr_QBuffer_connect4, "QBuffer_connect4")
	symaddr(&addr_QBuffer_metaObject_default, "QBuffer_metaObject_default")
	symaddr(&addr_QBuffer_open_default, "QBuffer_open_default")
	symaddr(&addr_QBuffer_close_default, "QBuffer_close_default")
	symaddr(&addr_QBuffer_size_default, "QBuffer_size_default")
	symaddr(&addr_QBuffer_pos_default, "QBuffer_pos_default")
	symaddr(&addr_QBuffer_seek_default, "QBuffer_seek_default")
	symaddr(&addr_QBuffer_atEnd_default, "QBuffer_atEnd_default")
	symaddr(&addr_QBuffer_canReadLine_default, "QBuffer_canReadLine_default")
	symaddr(&addr_QBuffer_connectNotify_default, "QBuffer_connectNotify_default")
	symaddr(&addr_QBuffer_disconnectNotify_default, "QBuffer_disconnectNotify_default")
	symaddr(&addr_QBuffer_isSequential_default, "QBuffer_isSequential_default")
	symaddr(&addr_QBuffer_reset_default, "QBuffer_reset_default")
	symaddr(&addr_QBuffer_bytesAvailable_default, "QBuffer_bytesAvailable_default")
	symaddr(&addr_QBuffer_bytesToWrite_default, "QBuffer_bytesToWrite_default")
	symaddr(&addr_QBuffer_waitForReadyRead_default, "QBuffer_waitForReadyRead_default")
	symaddr(&addr_QBuffer_waitForBytesWritten_default, "QBuffer_waitForBytesWritten_default")
	symaddr(&addr_QBuffer_skipData_default, "QBuffer_skipData_default")
	symaddr(&addr_QBuffer_event_default, "QBuffer_event_default")
	symaddr(&addr_QBuffer_eventFilter_default, "QBuffer_eventFilter_default")
	symaddr(&addr_QBuffer_timerEvent_default, "QBuffer_timerEvent_default")
	symaddr(&addr_QBuffer_childEvent_default, "QBuffer_childEvent_default")
	symaddr(&addr_QBuffer_customEvent_default, "QBuffer_customEvent_default")
	reg(&pg_QBuffer_connect_readyRead, "QBuffer_connect_readyRead")
	reg(&pg_QBuffer_connect_channelReadyRead, "QBuffer_connect_channelReadyRead")
	reg(&pg_QBuffer_connect_bytesWritten, "QBuffer_connect_bytesWritten")
	reg(&pg_QBuffer_connect_channelBytesWritten, "QBuffer_connect_channelBytesWritten")
	reg(&pg_QBuffer_connect_aboutToClose, "QBuffer_connect_aboutToClose")
	reg(&pg_QBuffer_connect_readChannelFinished, "QBuffer_connect_readChannelFinished")
	reg(&pg_QBuffer_connect_destroyed, "QBuffer_connect_destroyed")
	reg(&pg_QBuffer_connect_destroyed2, "QBuffer_connect_destroyed2")
	reg(&pg_QBuffer_connect_objectNameChanged, "QBuffer_connect_objectNameChanged")
	reg(&pg_QBuffer__disconnect, "QBuffer__disconnect")
	reg(&pg_QBuffer__set_override, "QBuffer__set_override")
	reg(&pg_QBuffer_delete, "QBuffer_delete")
}

type QBuffer struct {
	h uintptr
	*QIODevice
}

func bindQBuffer(h uintptr) *QBuffer {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QBuffer_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QBuffer{h: h, QIODevice: &QIODevice{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQBuffer(h uintptr) *QBuffer { return bindQBuffer(h) }

func (this *QBuffer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBuffer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBuffer(h unsafe.Pointer) *QBuffer { return bindQBuffer(uintptr(h)) }
func UnsafeNewQBufferOwned(h unsafe.Pointer) *QBuffer { return newownedQBuffer(uintptr(h)) }

func (this *QBuffer) UnsafeRelease() {
	this.h = 0
	if this.QIODevice != nil {
		this.QIODevice.UnsafeRelease()
	}
}

func NewQBuffer() *QBuffer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBuffer_new, uintptr(__goid))
	__o := newownedQBuffer(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBuffer_virtIdx, QBuffer_enable)
	return __o
}

func NewQBuffer2(buf *QByteArray) *QBuffer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBuffer_new2, uintptr(__goid), uintptr(buf.UnsafePointer()))
	__o := newownedQBuffer(__r)
	runtime.KeepAlive(buf)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBuffer_virtIdx, QBuffer_enable)
	return __o
}

func NewQBuffer3(parent *QObject) *QBuffer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBuffer_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQBuffer(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBuffer_virtIdx, QBuffer_enable)
	return __o
}

func NewQBuffer4(buf *QByteArray, parent *QObject) *QBuffer {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QBuffer_new4, uintptr(__goid), uintptr(buf.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQBuffer(__r)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QBuffer_virtIdx, QBuffer_enable)
	return __o
}

func QBuffer_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBuffer_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QBuffer) Buffer() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QBuffer_buffer, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QBuffer) Buffer2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QBuffer_buffer2, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QBuffer) SetBuffer(a *QByteArray) {
	purego.SyscallN(addr_QBuffer_setBuffer, this.h, uintptr(a.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a)
}

func (this *QBuffer) SetData(data *QByteArray) {
	purego.SyscallN(addr_QBuffer_setData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QBuffer) SetData2(data string, len_ int64) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QBuffer_setData2, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QBuffer) Data() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QBuffer_data, this.h)
	runtime.KeepAlive(this)
	return bindQByteArray(__r)
}

func (this *QBuffer) SetOpenMode(openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QBuffer_setOpenMode, this.h, uintptr(openMode))
	runtime.KeepAlive(this)
}

func (this *QBuffer) SetErrorString(errorString string) {
	var errorString_p uintptr
	if len(errorString) > 0 {
		errorString_p = uintptr(unsafe.Pointer(unsafe.StringData(errorString)))
	}
	purego.SyscallN(addr_QBuffer_setErrorString, this.h, errorString_p, uintptr(len(errorString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(errorString)
}

func QBuffer_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QBuffer_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QBuffer_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QBuffer_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QBuffer) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QBuffer_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QBuffer) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QBuffer) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QBuffer_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QBuffer) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QBuffer_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBuffer_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QBuffer_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBuffer_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QBuffer_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QBuffer_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QBuffer_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QBuffer_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QBuffer) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QBuffer_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QBuffer) OpenDefault(openMode QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_open_default, this.h, uintptr(openMode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) CloseDefault() {
	purego.SyscallN(addr_QBuffer_close_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QBuffer) SizeDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_size_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBuffer) PosDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_pos_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBuffer) SeekDefault(off int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_seek_default, this.h, uintptr(off))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) AtEndDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_atEnd_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) CanReadLineDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_canReadLine_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) ConnectNotifyDefault(arg1 *QMetaMethod) {
	purego.SyscallN(addr_QBuffer_connectNotify_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBuffer) DisconnectNotifyDefault(arg1 *QMetaMethod) {
	purego.SyscallN(addr_QBuffer_disconnectNotify_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBuffer) IsSequentialDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_isSequential_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) ResetDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_reset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) BytesAvailableDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_bytesAvailable_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBuffer) BytesToWriteDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_bytesToWrite_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBuffer) WaitForReadyReadDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_waitForReadyRead_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) WaitForBytesWrittenDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_waitForBytesWritten_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBuffer) SkipDataDefault(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_skipData_default, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBuffer) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QBuffer) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QBuffer_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QBuffer) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QBuffer_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QBuffer) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QBuffer_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QBuffer) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QBuffer_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QBuffer) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QBuffer_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QBuffer_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QBuffer_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QBuffer_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QBuffer_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QBuffer_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QBuffer_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QBuffer_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QBuffer_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBuffer) Disconnect(handle int64) {
	pg_QBuffer__disconnect(handle)
	unregisterCallback(handle)
}

func QBuffer_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -554057428128861975:
		return 1
	case -7006921981232612925:
		return 2
	case 5614464919154503228:
		return 3
	case 8618582476283917417:
		return 4
	case 3241140482992819217:
		return 5
	case -1487528724786208339:
		return 6
	case -1250682508954504761:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	case 8888098400087344862:
		return 10
	case 1086335023529244512:
		return 11
	case 4814449630054007943:
		return 12
	case 1024747515116322948:
		return 13
	case 8356976193102722774:
		return 14
	case 7726833620054767889:
		return 15
	case 8178738247667652082:
		return 16
	case 1237236456312824991:
		return 17
	case 5217623514630858297:
		return 18
	case -5409927581812374498:
		return 19
	case -1185315367526857833:
		return 20
	case 3769868229123868474:
		return 21
	}
	return -1
}

func QBuffer_enable(h uintptr, idx int32) { pg_QBuffer__set_override(h, idx) }

func (this *QBuffer) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QBuffer) OnOpen(fn func(QIODeviceBase__OpenModeFlag) bool) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func(openMode int32) int32 { return b2i(fn(QIODeviceBase__OpenModeFlag(openMode))) }) // slot "open"
}

func (this *QBuffer) OnClose(fn func()) {
	registerVirt(this.UnsafeGoID(), -7006921981232612925, func() { fn() }) // slot "close"
}

func (this *QBuffer) OnSize(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() int64 { return fn() }) // slot "size"
}

func (this *QBuffer) OnPos(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 8618582476283917417, func() int64 { return fn() }) // slot "pos"
}

func (this *QBuffer) OnSeek(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), 3241140482992819217, func(off int64) int32 { return b2i(fn(off)) }) // slot "seek"
}

func (this *QBuffer) OnAtEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1487528724786208339, func() int32 { return b2i(fn()) }) // slot "atEnd"
}

func (this *QBuffer) OnCanReadLine(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1250682508954504761, func() int32 { return b2i(fn()) }) // slot "canReadLine"
}

func (this *QBuffer) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(arg1 uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(arg1))) }) // slot "connectNotify"
}

func (this *QBuffer) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(arg1 uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(arg1))) }) // slot "disconnectNotify"
}

func (this *QBuffer) OnIsSequential(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8888098400087344862, func() int32 { return b2i(fn()) }) // slot "isSequential"
}

func (this *QBuffer) OnReset(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() int32 { return b2i(fn()) }) // slot "reset"
}

func (this *QBuffer) OnBytesAvailable(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 4814449630054007943, func() int64 { return fn() }) // slot "bytesAvailable"
}

func (this *QBuffer) OnBytesToWrite(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 1024747515116322948, func() int64 { return fn() }) // slot "bytesToWrite"
}

func (this *QBuffer) OnWaitForReadyRead(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 8356976193102722774, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForReadyRead"
}

func (this *QBuffer) OnWaitForBytesWritten(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 7726833620054767889, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForBytesWritten"
}

func (this *QBuffer) OnSkipData(fn func(int64) int64) {
	registerVirt(this.UnsafeGoID(), 8178738247667652082, func(maxSize int64) int64 { return fn(maxSize) }) // slot "skipData"
}

func (this *QBuffer) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QBuffer) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QBuffer) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QBuffer) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QBuffer) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QBuffer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBuffer_delete(this.h)
	}
	this.UnsafeRelease()
}
