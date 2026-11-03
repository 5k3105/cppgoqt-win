package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QIODevice_base func(uintptr) uintptr
	addr_QIODevice_tr uintptr
	addr_QIODevice_openMode uintptr
	addr_QIODevice_setTextModeEnabled uintptr
	addr_QIODevice_isTextModeEnabled uintptr
	addr_QIODevice_isOpen uintptr
	addr_QIODevice_isReadable uintptr
	addr_QIODevice_isWritable uintptr
	addr_QIODevice_isSequential uintptr
	addr_QIODevice_readChannelCount uintptr
	addr_QIODevice_writeChannelCount uintptr
	addr_QIODevice_currentReadChannel uintptr
	addr_QIODevice_setCurrentReadChannel uintptr
	addr_QIODevice_currentWriteChannel uintptr
	addr_QIODevice_setCurrentWriteChannel uintptr
	addr_QIODevice_open uintptr
	addr_QIODevice_close uintptr
	addr_QIODevice_pos uintptr
	addr_QIODevice_size uintptr
	addr_QIODevice_seek uintptr
	addr_QIODevice_atEnd uintptr
	addr_QIODevice_reset uintptr
	addr_QIODevice_bytesAvailable uintptr
	addr_QIODevice_bytesToWrite uintptr
	addr_QIODevice_read uintptr
	addr_QIODevice_read2 uintptr
	addr_QIODevice_readAll uintptr
	addr_QIODevice_readLine uintptr
	addr_QIODevice_readLine2 uintptr
	addr_QIODevice_readLineInto uintptr
	addr_QIODevice_canReadLine uintptr
	addr_QIODevice_startTransaction uintptr
	addr_QIODevice_commitTransaction uintptr
	addr_QIODevice_rollbackTransaction uintptr
	addr_QIODevice_isTransactionStarted uintptr
	addr_QIODevice_write uintptr
	addr_QIODevice_write2 uintptr
	addr_QIODevice_write3 uintptr
	addr_QIODevice_peek uintptr
	addr_QIODevice_peek2 uintptr
	addr_QIODevice_skip uintptr
	addr_QIODevice_waitForReadyRead uintptr
	addr_QIODevice_waitForBytesWritten uintptr
	addr_QIODevice_ungetChar uintptr
	addr_QIODevice_putChar uintptr
	addr_QIODevice_getChar uintptr
	addr_QIODevice_errorString uintptr
	addr_QIODevice_readyRead uintptr
	addr_QIODevice_channelReadyRead uintptr
	addr_QIODevice_bytesWritten uintptr
	addr_QIODevice_channelBytesWritten uintptr
	addr_QIODevice_aboutToClose uintptr
	addr_QIODevice_readChannelFinished uintptr
	addr_QIODevice_connect uintptr
	addr_QIODevice_connect2 uintptr
	addr_QIODevice_tr2 uintptr
	addr_QIODevice_tr3 uintptr
	addr_QIODevice_readLine3 uintptr
	addr_QIODevice_readLineInto2 uintptr
	addr_QIODevice_connect3 uintptr
	addr_QIODevice_connect4 uintptr
	addr_QIODevice_asQIODeviceBase uintptr
	pg_QIODevice_connect_readyRead func(uintptr, int64)
	pg_QIODevice_connect_channelReadyRead func(uintptr, int64)
	pg_QIODevice_connect_bytesWritten func(uintptr, int64)
	pg_QIODevice_connect_channelBytesWritten func(uintptr, int64)
	pg_QIODevice_connect_aboutToClose func(uintptr, int64)
	pg_QIODevice_connect_readChannelFinished func(uintptr, int64)
	pg_QIODevice_connect_destroyed func(uintptr, int64)
	pg_QIODevice_connect_destroyed2 func(uintptr, int64)
	pg_QIODevice_connect_objectNameChanged func(uintptr, int64)
	pg_QIODevice__disconnect func(int64)
	pg_QIODevice_delete func(uintptr)
)

func init() {
	reg(&pg_QIODevice_base, "QIODevice_base")
	symaddr(&addr_QIODevice_tr, "QIODevice_tr")
	symaddr(&addr_QIODevice_openMode, "QIODevice_openMode")
	symaddr(&addr_QIODevice_setTextModeEnabled, "QIODevice_setTextModeEnabled")
	symaddr(&addr_QIODevice_isTextModeEnabled, "QIODevice_isTextModeEnabled")
	symaddr(&addr_QIODevice_isOpen, "QIODevice_isOpen")
	symaddr(&addr_QIODevice_isReadable, "QIODevice_isReadable")
	symaddr(&addr_QIODevice_isWritable, "QIODevice_isWritable")
	symaddr(&addr_QIODevice_isSequential, "QIODevice_isSequential")
	symaddr(&addr_QIODevice_readChannelCount, "QIODevice_readChannelCount")
	symaddr(&addr_QIODevice_writeChannelCount, "QIODevice_writeChannelCount")
	symaddr(&addr_QIODevice_currentReadChannel, "QIODevice_currentReadChannel")
	symaddr(&addr_QIODevice_setCurrentReadChannel, "QIODevice_setCurrentReadChannel")
	symaddr(&addr_QIODevice_currentWriteChannel, "QIODevice_currentWriteChannel")
	symaddr(&addr_QIODevice_setCurrentWriteChannel, "QIODevice_setCurrentWriteChannel")
	symaddr(&addr_QIODevice_open, "QIODevice_open")
	symaddr(&addr_QIODevice_close, "QIODevice_close")
	symaddr(&addr_QIODevice_pos, "QIODevice_pos")
	symaddr(&addr_QIODevice_size, "QIODevice_size")
	symaddr(&addr_QIODevice_seek, "QIODevice_seek")
	symaddr(&addr_QIODevice_atEnd, "QIODevice_atEnd")
	symaddr(&addr_QIODevice_reset, "QIODevice_reset")
	symaddr(&addr_QIODevice_bytesAvailable, "QIODevice_bytesAvailable")
	symaddr(&addr_QIODevice_bytesToWrite, "QIODevice_bytesToWrite")
	symaddr(&addr_QIODevice_read, "QIODevice_read")
	symaddr(&addr_QIODevice_read2, "QIODevice_read2")
	symaddr(&addr_QIODevice_readAll, "QIODevice_readAll")
	symaddr(&addr_QIODevice_readLine, "QIODevice_readLine")
	symaddr(&addr_QIODevice_readLine2, "QIODevice_readLine2")
	symaddr(&addr_QIODevice_readLineInto, "QIODevice_readLineInto")
	symaddr(&addr_QIODevice_canReadLine, "QIODevice_canReadLine")
	symaddr(&addr_QIODevice_startTransaction, "QIODevice_startTransaction")
	symaddr(&addr_QIODevice_commitTransaction, "QIODevice_commitTransaction")
	symaddr(&addr_QIODevice_rollbackTransaction, "QIODevice_rollbackTransaction")
	symaddr(&addr_QIODevice_isTransactionStarted, "QIODevice_isTransactionStarted")
	symaddr(&addr_QIODevice_write, "QIODevice_write")
	symaddr(&addr_QIODevice_write2, "QIODevice_write2")
	symaddr(&addr_QIODevice_write3, "QIODevice_write3")
	symaddr(&addr_QIODevice_peek, "QIODevice_peek")
	symaddr(&addr_QIODevice_peek2, "QIODevice_peek2")
	symaddr(&addr_QIODevice_skip, "QIODevice_skip")
	symaddr(&addr_QIODevice_waitForReadyRead, "QIODevice_waitForReadyRead")
	symaddr(&addr_QIODevice_waitForBytesWritten, "QIODevice_waitForBytesWritten")
	symaddr(&addr_QIODevice_ungetChar, "QIODevice_ungetChar")
	symaddr(&addr_QIODevice_putChar, "QIODevice_putChar")
	symaddr(&addr_QIODevice_getChar, "QIODevice_getChar")
	symaddr(&addr_QIODevice_errorString, "QIODevice_errorString")
	symaddr(&addr_QIODevice_readyRead, "QIODevice_readyRead")
	symaddr(&addr_QIODevice_channelReadyRead, "QIODevice_channelReadyRead")
	symaddr(&addr_QIODevice_bytesWritten, "QIODevice_bytesWritten")
	symaddr(&addr_QIODevice_channelBytesWritten, "QIODevice_channelBytesWritten")
	symaddr(&addr_QIODevice_aboutToClose, "QIODevice_aboutToClose")
	symaddr(&addr_QIODevice_readChannelFinished, "QIODevice_readChannelFinished")
	symaddr(&addr_QIODevice_connect, "QIODevice_connect")
	symaddr(&addr_QIODevice_connect2, "QIODevice_connect2")
	symaddr(&addr_QIODevice_tr2, "QIODevice_tr2")
	symaddr(&addr_QIODevice_tr3, "QIODevice_tr3")
	symaddr(&addr_QIODevice_readLine3, "QIODevice_readLine3")
	symaddr(&addr_QIODevice_readLineInto2, "QIODevice_readLineInto2")
	symaddr(&addr_QIODevice_connect3, "QIODevice_connect3")
	symaddr(&addr_QIODevice_connect4, "QIODevice_connect4")
	symaddr(&addr_QIODevice_asQIODeviceBase, "QIODevice_asQIODeviceBase")
	reg(&pg_QIODevice_connect_readyRead, "QIODevice_connect_readyRead")
	reg(&pg_QIODevice_connect_channelReadyRead, "QIODevice_connect_channelReadyRead")
	reg(&pg_QIODevice_connect_bytesWritten, "QIODevice_connect_bytesWritten")
	reg(&pg_QIODevice_connect_channelBytesWritten, "QIODevice_connect_channelBytesWritten")
	reg(&pg_QIODevice_connect_aboutToClose, "QIODevice_connect_aboutToClose")
	reg(&pg_QIODevice_connect_readChannelFinished, "QIODevice_connect_readChannelFinished")
	reg(&pg_QIODevice_connect_destroyed, "QIODevice_connect_destroyed")
	reg(&pg_QIODevice_connect_destroyed2, "QIODevice_connect_destroyed2")
	reg(&pg_QIODevice_connect_objectNameChanged, "QIODevice_connect_objectNameChanged")
	reg(&pg_QIODevice__disconnect, "QIODevice__disconnect")
	reg(&pg_QIODevice_delete, "QIODevice_delete")
}

type QIODevice struct {
	h uintptr
	*QObject
}

func bindQIODevice(h uintptr) *QIODevice {
	if h == 0 {
		return nil
	}
	return &QIODevice{h: h, QObject: bindQObject(pg_QIODevice_base(h))}
}

func newownedQIODevice(h uintptr) *QIODevice { return bindQIODevice(h) }

func (this *QIODevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIODevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIODevice(h unsafe.Pointer) *QIODevice { return bindQIODevice(uintptr(h)) }
func UnsafeNewQIODeviceOwned(h unsafe.Pointer) *QIODevice { return newownedQIODevice(uintptr(h)) }

func (this *QIODevice) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QIODevice_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIODevice_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QIODevice) OpenMode() QIODeviceBase__OpenModeFlag {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_openMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QIODeviceBase__OpenModeFlag(__r)
}

func (this *QIODevice) SetTextModeEnabled(enabled bool) {
	purego.SyscallN(addr_QIODevice_setTextModeEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QIODevice) IsTextModeEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isTextModeEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) IsOpen() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isOpen, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) IsReadable() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isReadable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) IsWritable() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isWritable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) IsSequential() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isSequential, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) ReadChannelCount() int {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_readChannelCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIODevice) WriteChannelCount() int {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_writeChannelCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIODevice) CurrentReadChannel() int {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_currentReadChannel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIODevice) SetCurrentReadChannel(channel int) {
	purego.SyscallN(addr_QIODevice_setCurrentReadChannel, this.h, uintptr(channel))
	runtime.KeepAlive(this)
}

func (this *QIODevice) CurrentWriteChannel() int {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_currentWriteChannel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIODevice) SetCurrentWriteChannel(channel int) {
	purego.SyscallN(addr_QIODevice_setCurrentWriteChannel, this.h, uintptr(channel))
	runtime.KeepAlive(this)
}

func (this *QIODevice) Open(mode QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_open, this.h, uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) Close() {
	purego.SyscallN(addr_QIODevice_close, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) Pos() int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_pos, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIODevice) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIODevice) Seek(pos int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_seek, this.h, uintptr(pos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) AtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_atEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) Reset() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_reset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) BytesAvailable() int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_bytesAvailable, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIODevice) BytesToWrite() int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_bytesToWrite, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIODevice) Read(data []byte, maxlen int64) int64 {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QIODevice_read, this.h, data_p, uintptr(maxlen))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QIODevice) ReadWithMaxlen(maxlen int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QIODevice_read2, this.h, uintptr(maxlen))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QIODevice) ReadAll() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QIODevice_readAll, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QIODevice) ReadLine(data []byte, maxlen int64) int64 {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QIODevice_readLine, this.h, data_p, uintptr(maxlen))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QIODevice) ReadLine2() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QIODevice_readLine2, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QIODevice) ReadLineInto(result *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_readLineInto, this.h, uintptr(result.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(result)
	return __r != 0
}

func (this *QIODevice) CanReadLine() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_canReadLine, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) StartTransaction() {
	purego.SyscallN(addr_QIODevice_startTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) CommitTransaction() {
	purego.SyscallN(addr_QIODevice_commitTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) RollbackTransaction() {
	purego.SyscallN(addr_QIODevice_rollbackTransaction, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) IsTransactionStarted() bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_isTransactionStarted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) Write(data string, len_ int64) int64 {
	data_b := append([]byte(data), 0)
	__r0, _, _ := purego.SyscallN(addr_QIODevice_write, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
	return __r
}

func (this *QIODevice) WriteWithData(data string) int64 {
	data_b := append([]byte(data), 0)
	__r0, _, _ := purego.SyscallN(addr_QIODevice_write2, this.h, uintptr(unsafe.Pointer(&data_b[0])))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
	return __r
}

func (this *QIODevice) Write2(data *QByteArray) int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_write3, this.h, uintptr(data.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QIODevice) Peek(data []byte, maxlen int64) int64 {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QIODevice_peek, this.h, data_p, uintptr(maxlen))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QIODevice) PeekWithMaxlen(maxlen int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QIODevice_peek2, this.h, uintptr(maxlen))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QIODevice) Skip(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_skip, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QIODevice) WaitForReadyRead(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_waitForReadyRead, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) WaitForBytesWritten(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_waitForBytesWritten, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) UngetChar(c int8) {
	purego.SyscallN(addr_QIODevice_ungetChar, this.h, uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QIODevice) PutChar(c int8) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_putChar, this.h, uintptr(c))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QIODevice) GetChar(c []byte) bool {
	var c_p uintptr
	if len(c) > 0 {
		c_p = uintptr(unsafe.Pointer(&c[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QIODevice_getChar, this.h, c_p)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return __r != 0
}

func (this *QIODevice) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIODevice_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QIODevice) ReadyRead() {
	purego.SyscallN(addr_QIODevice_readyRead, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) ChannelReadyRead(channel int) {
	purego.SyscallN(addr_QIODevice_channelReadyRead, this.h, uintptr(channel))
	runtime.KeepAlive(this)
}

func (this *QIODevice) BytesWritten(bytes int64) {
	purego.SyscallN(addr_QIODevice_bytesWritten, this.h, uintptr(bytes))
	runtime.KeepAlive(this)
}

func (this *QIODevice) ChannelBytesWritten(channel int, bytes int64) {
	purego.SyscallN(addr_QIODevice_channelBytesWritten, this.h, uintptr(channel), uintptr(bytes))
	runtime.KeepAlive(this)
}

func (this *QIODevice) AboutToClose() {
	purego.SyscallN(addr_QIODevice_aboutToClose, this.h)
	runtime.KeepAlive(this)
}

func (this *QIODevice) ReadChannelFinished() {
	purego.SyscallN(addr_QIODevice_readChannelFinished, this.h)
	runtime.KeepAlive(this)
}

func QIODevice_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIODevice_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIODevice_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIODevice_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QIODevice_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIODevice_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QIODevice_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIODevice_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QIODevice) ReadLineWithMaxlen(maxlen int64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QIODevice_readLine3, this.h, uintptr(maxlen))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QIODevice) ReadLineInto2(result *QByteArray, maxlen int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QIODevice_readLineInto2, this.h, uintptr(result.UnsafePointer()), uintptr(maxlen))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(result)
	return __r != 0
}

func QIODevice_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIODevice_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIODevice_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIODevice_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QIODevice) AsQIODeviceBase() *QIODeviceBase {
	__r, _, _ := purego.SyscallN(addr_QIODevice_asQIODeviceBase, this.h)
	runtime.KeepAlive(this)
	return bindQIODeviceBase(__r)
}

func (this *QIODevice) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIODevice_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QIODevice_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QIODevice_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QIODevice_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIODevice_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIODevice_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIODevice_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QIODevice_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QIODevice_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIODevice) Disconnect(handle int64) {
	pg_QIODevice__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QIODevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIODevice_delete(this.h)
	}
	this.UnsafeRelease()
}
