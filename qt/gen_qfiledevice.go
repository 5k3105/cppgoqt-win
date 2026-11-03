package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFileDevice_chain uintptr
	addr_QFileDevice_tr uintptr
	addr_QFileDevice_error uintptr
	addr_QFileDevice_unsetError uintptr
	addr_QFileDevice_handle uintptr
	addr_QFileDevice_fileName uintptr
	addr_QFileDevice_flush uintptr
	addr_QFileDevice_resize uintptr
	addr_QFileDevice_permissions uintptr
	addr_QFileDevice_setPermissions uintptr
	addr_QFileDevice_unmap uintptr
	addr_QFileDevice_fileTime uintptr
	addr_QFileDevice_setFileTime uintptr
	addr_QFileDevice_connect uintptr
	addr_QFileDevice_connect2 uintptr
	addr_QFileDevice_tr2 uintptr
	addr_QFileDevice_tr3 uintptr
	addr_QFileDevice_connect3 uintptr
	addr_QFileDevice_connect4 uintptr
	pg_QFileDevice_connect_readyRead func(uintptr, int64)
	pg_QFileDevice_connect_channelReadyRead func(uintptr, int64)
	pg_QFileDevice_connect_bytesWritten func(uintptr, int64)
	pg_QFileDevice_connect_channelBytesWritten func(uintptr, int64)
	pg_QFileDevice_connect_aboutToClose func(uintptr, int64)
	pg_QFileDevice_connect_readChannelFinished func(uintptr, int64)
	pg_QFileDevice_connect_destroyed func(uintptr, int64)
	pg_QFileDevice_connect_destroyed2 func(uintptr, int64)
	pg_QFileDevice_connect_objectNameChanged func(uintptr, int64)
	pg_QFileDevice__disconnect func(int64)
	pg_QFileDevice_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFileDevice_chain, "QFileDevice_chain")
	symaddr(&addr_QFileDevice_tr, "QFileDevice_tr")
	symaddr(&addr_QFileDevice_error, "QFileDevice_error")
	symaddr(&addr_QFileDevice_unsetError, "QFileDevice_unsetError")
	symaddr(&addr_QFileDevice_handle, "QFileDevice_handle")
	symaddr(&addr_QFileDevice_fileName, "QFileDevice_fileName")
	symaddr(&addr_QFileDevice_flush, "QFileDevice_flush")
	symaddr(&addr_QFileDevice_resize, "QFileDevice_resize")
	symaddr(&addr_QFileDevice_permissions, "QFileDevice_permissions")
	symaddr(&addr_QFileDevice_setPermissions, "QFileDevice_setPermissions")
	symaddr(&addr_QFileDevice_unmap, "QFileDevice_unmap")
	symaddr(&addr_QFileDevice_fileTime, "QFileDevice_fileTime")
	symaddr(&addr_QFileDevice_setFileTime, "QFileDevice_setFileTime")
	symaddr(&addr_QFileDevice_connect, "QFileDevice_connect")
	symaddr(&addr_QFileDevice_connect2, "QFileDevice_connect2")
	symaddr(&addr_QFileDevice_tr2, "QFileDevice_tr2")
	symaddr(&addr_QFileDevice_tr3, "QFileDevice_tr3")
	symaddr(&addr_QFileDevice_connect3, "QFileDevice_connect3")
	symaddr(&addr_QFileDevice_connect4, "QFileDevice_connect4")
	reg(&pg_QFileDevice_connect_readyRead, "QFileDevice_connect_readyRead")
	reg(&pg_QFileDevice_connect_channelReadyRead, "QFileDevice_connect_channelReadyRead")
	reg(&pg_QFileDevice_connect_bytesWritten, "QFileDevice_connect_bytesWritten")
	reg(&pg_QFileDevice_connect_channelBytesWritten, "QFileDevice_connect_channelBytesWritten")
	reg(&pg_QFileDevice_connect_aboutToClose, "QFileDevice_connect_aboutToClose")
	reg(&pg_QFileDevice_connect_readChannelFinished, "QFileDevice_connect_readChannelFinished")
	reg(&pg_QFileDevice_connect_destroyed, "QFileDevice_connect_destroyed")
	reg(&pg_QFileDevice_connect_destroyed2, "QFileDevice_connect_destroyed2")
	reg(&pg_QFileDevice_connect_objectNameChanged, "QFileDevice_connect_objectNameChanged")
	reg(&pg_QFileDevice__disconnect, "QFileDevice__disconnect")
	reg(&pg_QFileDevice_delete, "QFileDevice_delete")
}

type QFileDevice struct {
	h uintptr
	*QIODevice
}

func bindQFileDevice(h uintptr) *QFileDevice {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QFileDevice_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFileDevice{h: h, QIODevice: &QIODevice{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQFileDevice(h uintptr) *QFileDevice { return bindQFileDevice(h) }

func (this *QFileDevice) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileDevice) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileDevice(h unsafe.Pointer) *QFileDevice { return bindQFileDevice(uintptr(h)) }
func UnsafeNewQFileDeviceOwned(h unsafe.Pointer) *QFileDevice { return newownedQFileDevice(uintptr(h)) }

func (this *QFileDevice) UnsafeRelease() {
	this.h = 0
	if this.QIODevice != nil {
		this.QIODevice.UnsafeRelease()
	}
}

func QFileDevice_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDevice_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFileDevice) Error() QFileDevice__FileError {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__FileError(__r)
}

func (this *QFileDevice) UnsetError() {
	purego.SyscallN(addr_QFileDevice_unsetError, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileDevice) Handle() int {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_handle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileDevice) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDevice_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileDevice) Flush() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_flush, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDevice) Resize(sz int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_resize, this.h, uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDevice) Permissions() QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_permissions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__Permission(__r)
}

func (this *QFileDevice) SetPermissions(permissionSpec QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_setPermissions, this.h, uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileDevice) Unmap() (bool, byte) {
	var address_out uint8
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_unmap, this.h, uintptr(unsafe.Pointer(&address_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, byte(address_out)
}

func (this *QFileDevice) FileTime(time QFileDevice__FileTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileDevice_fileTime, this.h, uintptr(time))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QFileDevice) SetFileTime(newDate *QDateTime, fileTime QFileDevice__FileTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileDevice_setFileTime, this.h, uintptr(newDate.UnsafePointer()), uintptr(fileTime))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newDate)
	return __r != 0
}

func QFileDevice_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileDevice_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileDevice_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileDevice_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QFileDevice_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDevice_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileDevice_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileDevice_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileDevice_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileDevice_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileDevice_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileDevice_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileDevice) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDevice_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QFileDevice_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QFileDevice_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QFileDevice_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDevice_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDevice_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileDevice_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFileDevice_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFileDevice_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileDevice) Disconnect(handle int64) {
	pg_QFileDevice__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QFileDevice) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileDevice_delete(this.h)
	}
	this.UnsafeRelease()
}
