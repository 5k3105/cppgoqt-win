package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFile_chain uintptr
	addr_QFile_new uintptr
	addr_QFile_new2 uintptr
	addr_QFile_new3 uintptr
	addr_QFile_new4 uintptr
	addr_QFile_tr uintptr
	addr_QFile_setFileName uintptr
	addr_QFile_encodeName uintptr
	addr_QFile_decodeName uintptr
	addr_QFile_decodeName2 uintptr
	addr_QFile_exists uintptr
	addr_QFile_exists2 uintptr
	addr_QFile_symLinkTarget uintptr
	addr_QFile_symLinkTarget2 uintptr
	addr_QFile_remove uintptr
	addr_QFile_remove2 uintptr
	addr_QFile_supportsMoveToTrash uintptr
	addr_QFile_moveToTrash uintptr
	addr_QFile_moveToTrash2 uintptr
	addr_QFile_rename uintptr
	addr_QFile_rename2 uintptr
	addr_QFile_link uintptr
	addr_QFile_link2 uintptr
	addr_QFile_copy uintptr
	addr_QFile_copy2 uintptr
	addr_QFile_open uintptr
	addr_QFile_open2 uintptr
	addr_QFile_open3 uintptr
	addr_QFile_resize uintptr
	addr_QFile_permissions uintptr
	addr_QFile_setPermissions uintptr
	addr_QFile_setOpenMode uintptr
	addr_QFile_setErrorString uintptr
	addr_QFile_connect uintptr
	addr_QFile_connect2 uintptr
	addr_QFile_sender uintptr
	addr_QFile_senderSignalIndex uintptr
	addr_QFile_receivers uintptr
	addr_QFile_isSignalConnected uintptr
	addr_QFile_tr2 uintptr
	addr_QFile_tr3 uintptr
	addr_QFile_moveToTrash3 uintptr
	addr_QFile_open4 uintptr
	addr_QFile_connect3 uintptr
	addr_QFile_connect4 uintptr
	addr_QFile_open_default uintptr
	addr_QFile_open_default2 uintptr
	addr_QFile_open_default3 uintptr
	addr_QFile_open_default4 uintptr
	addr_QFile_metaObject_default uintptr
	addr_QFile_fileName_default uintptr
	addr_QFile_size_default uintptr
	addr_QFile_resize_default uintptr
	addr_QFile_permissions_default uintptr
	addr_QFile_setPermissions_default uintptr
	addr_QFile_close_default uintptr
	addr_QFile_isSequential_default uintptr
	addr_QFile_pos_default uintptr
	addr_QFile_seek_default uintptr
	addr_QFile_atEnd_default uintptr
	addr_QFile_reset_default uintptr
	addr_QFile_bytesAvailable_default uintptr
	addr_QFile_bytesToWrite_default uintptr
	addr_QFile_canReadLine_default uintptr
	addr_QFile_waitForReadyRead_default uintptr
	addr_QFile_waitForBytesWritten_default uintptr
	addr_QFile_skipData_default uintptr
	addr_QFile_event_default uintptr
	addr_QFile_eventFilter_default uintptr
	addr_QFile_timerEvent_default uintptr
	addr_QFile_childEvent_default uintptr
	addr_QFile_customEvent_default uintptr
	addr_QFile_connectNotify_default uintptr
	addr_QFile_disconnectNotify_default uintptr
	pg_QFile_connect_readyRead func(uintptr, int64)
	pg_QFile_connect_channelReadyRead func(uintptr, int64)
	pg_QFile_connect_bytesWritten func(uintptr, int64)
	pg_QFile_connect_channelBytesWritten func(uintptr, int64)
	pg_QFile_connect_aboutToClose func(uintptr, int64)
	pg_QFile_connect_readChannelFinished func(uintptr, int64)
	pg_QFile_connect_destroyed func(uintptr, int64)
	pg_QFile_connect_destroyed2 func(uintptr, int64)
	pg_QFile_connect_objectNameChanged func(uintptr, int64)
	pg_QFile__disconnect func(int64)
	pg_QFile__set_override func(uintptr, int32)
	pg_QFile_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFile_chain, "QFile_chain")
	symaddr(&addr_QFile_new, "QFile_new")
	symaddr(&addr_QFile_new2, "QFile_new2")
	symaddr(&addr_QFile_new3, "QFile_new3")
	symaddr(&addr_QFile_new4, "QFile_new4")
	symaddr(&addr_QFile_tr, "QFile_tr")
	symaddr(&addr_QFile_setFileName, "QFile_setFileName")
	symaddr(&addr_QFile_encodeName, "QFile_encodeName")
	symaddr(&addr_QFile_decodeName, "QFile_decodeName")
	symaddr(&addr_QFile_decodeName2, "QFile_decodeName2")
	symaddr(&addr_QFile_exists, "QFile_exists")
	symaddr(&addr_QFile_exists2, "QFile_exists2")
	symaddr(&addr_QFile_symLinkTarget, "QFile_symLinkTarget")
	symaddr(&addr_QFile_symLinkTarget2, "QFile_symLinkTarget2")
	symaddr(&addr_QFile_remove, "QFile_remove")
	symaddr(&addr_QFile_remove2, "QFile_remove2")
	symaddr(&addr_QFile_supportsMoveToTrash, "QFile_supportsMoveToTrash")
	symaddr(&addr_QFile_moveToTrash, "QFile_moveToTrash")
	symaddr(&addr_QFile_moveToTrash2, "QFile_moveToTrash2")
	symaddr(&addr_QFile_rename, "QFile_rename")
	symaddr(&addr_QFile_rename2, "QFile_rename2")
	symaddr(&addr_QFile_link, "QFile_link")
	symaddr(&addr_QFile_link2, "QFile_link2")
	symaddr(&addr_QFile_copy, "QFile_copy")
	symaddr(&addr_QFile_copy2, "QFile_copy2")
	symaddr(&addr_QFile_open, "QFile_open")
	symaddr(&addr_QFile_open2, "QFile_open2")
	symaddr(&addr_QFile_open3, "QFile_open3")
	symaddr(&addr_QFile_resize, "QFile_resize")
	symaddr(&addr_QFile_permissions, "QFile_permissions")
	symaddr(&addr_QFile_setPermissions, "QFile_setPermissions")
	symaddr(&addr_QFile_setOpenMode, "QFile_setOpenMode")
	symaddr(&addr_QFile_setErrorString, "QFile_setErrorString")
	symaddr(&addr_QFile_connect, "QFile_connect")
	symaddr(&addr_QFile_connect2, "QFile_connect2")
	symaddr(&addr_QFile_sender, "QFile_sender")
	symaddr(&addr_QFile_senderSignalIndex, "QFile_senderSignalIndex")
	symaddr(&addr_QFile_receivers, "QFile_receivers")
	symaddr(&addr_QFile_isSignalConnected, "QFile_isSignalConnected")
	symaddr(&addr_QFile_tr2, "QFile_tr2")
	symaddr(&addr_QFile_tr3, "QFile_tr3")
	symaddr(&addr_QFile_moveToTrash3, "QFile_moveToTrash3")
	symaddr(&addr_QFile_open4, "QFile_open4")
	symaddr(&addr_QFile_connect3, "QFile_connect3")
	symaddr(&addr_QFile_connect4, "QFile_connect4")
	symaddr(&addr_QFile_open_default, "QFile_open_default")
	symaddr(&addr_QFile_open_default2, "QFile_open_default2")
	symaddr(&addr_QFile_open_default3, "QFile_open_default3")
	symaddr(&addr_QFile_open_default4, "QFile_open_default4")
	symaddr(&addr_QFile_metaObject_default, "QFile_metaObject_default")
	symaddr(&addr_QFile_fileName_default, "QFile_fileName_default")
	symaddr(&addr_QFile_size_default, "QFile_size_default")
	symaddr(&addr_QFile_resize_default, "QFile_resize_default")
	symaddr(&addr_QFile_permissions_default, "QFile_permissions_default")
	symaddr(&addr_QFile_setPermissions_default, "QFile_setPermissions_default")
	symaddr(&addr_QFile_close_default, "QFile_close_default")
	symaddr(&addr_QFile_isSequential_default, "QFile_isSequential_default")
	symaddr(&addr_QFile_pos_default, "QFile_pos_default")
	symaddr(&addr_QFile_seek_default, "QFile_seek_default")
	symaddr(&addr_QFile_atEnd_default, "QFile_atEnd_default")
	symaddr(&addr_QFile_reset_default, "QFile_reset_default")
	symaddr(&addr_QFile_bytesAvailable_default, "QFile_bytesAvailable_default")
	symaddr(&addr_QFile_bytesToWrite_default, "QFile_bytesToWrite_default")
	symaddr(&addr_QFile_canReadLine_default, "QFile_canReadLine_default")
	symaddr(&addr_QFile_waitForReadyRead_default, "QFile_waitForReadyRead_default")
	symaddr(&addr_QFile_waitForBytesWritten_default, "QFile_waitForBytesWritten_default")
	symaddr(&addr_QFile_skipData_default, "QFile_skipData_default")
	symaddr(&addr_QFile_event_default, "QFile_event_default")
	symaddr(&addr_QFile_eventFilter_default, "QFile_eventFilter_default")
	symaddr(&addr_QFile_timerEvent_default, "QFile_timerEvent_default")
	symaddr(&addr_QFile_childEvent_default, "QFile_childEvent_default")
	symaddr(&addr_QFile_customEvent_default, "QFile_customEvent_default")
	symaddr(&addr_QFile_connectNotify_default, "QFile_connectNotify_default")
	symaddr(&addr_QFile_disconnectNotify_default, "QFile_disconnectNotify_default")
	reg(&pg_QFile_connect_readyRead, "QFile_connect_readyRead")
	reg(&pg_QFile_connect_channelReadyRead, "QFile_connect_channelReadyRead")
	reg(&pg_QFile_connect_bytesWritten, "QFile_connect_bytesWritten")
	reg(&pg_QFile_connect_channelBytesWritten, "QFile_connect_channelBytesWritten")
	reg(&pg_QFile_connect_aboutToClose, "QFile_connect_aboutToClose")
	reg(&pg_QFile_connect_readChannelFinished, "QFile_connect_readChannelFinished")
	reg(&pg_QFile_connect_destroyed, "QFile_connect_destroyed")
	reg(&pg_QFile_connect_destroyed2, "QFile_connect_destroyed2")
	reg(&pg_QFile_connect_objectNameChanged, "QFile_connect_objectNameChanged")
	reg(&pg_QFile__disconnect, "QFile__disconnect")
	reg(&pg_QFile__set_override, "QFile__set_override")
	reg(&pg_QFile_delete, "QFile_delete")
}

type QFile struct {
	h uintptr
	*QFileDevice
}

func bindQFile(h uintptr) *QFile {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QFile_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFile{h: h, QFileDevice: &QFileDevice{h: __c[0], QIODevice: &QIODevice{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQFile(h uintptr) *QFile { return bindQFile(h) }

func (this *QFile) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFile) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFile(h unsafe.Pointer) *QFile { return bindQFile(uintptr(h)) }
func UnsafeNewQFileOwned(h unsafe.Pointer) *QFile { return newownedQFile(uintptr(h)) }

func (this *QFile) UnsafeRelease() {
	this.h = 0
	if this.QFileDevice != nil {
		this.QFileDevice.UnsafeRelease()
	}
}

func NewQFile() *QFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFile_new, uintptr(__goid))
	__o := newownedQFile(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFile_virtIdx, QFile_enable)
	return __o
}

func NewQFile2(name string) *QFile {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QFile_new2, uintptr(__goid), name_p, uintptr(len(name)))
	__o := newownedQFile(__r)
	runtime.KeepAlive(name)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFile_virtIdx, QFile_enable)
	return __o
}

func NewQFile3(parent *QObject) *QFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFile_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFile(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFile_virtIdx, QFile_enable)
	return __o
}

func NewQFile4(name string, parent *QObject) *QFile {
	__goid := nextVirtID()
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QFile_new4, uintptr(__goid), name_p, uintptr(len(name)), uintptr(parent.UnsafePointer()))
	__o := newownedQFile(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFile_virtIdx, QFile_enable)
	return __o
}

func QFile_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFile) SetFileName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QFile_setFileName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func QFile_EncodeName(fileName string) *QByteArray {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QFile_encodeName, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(fileName)
	return newownedQByteArray(__r)
}

func QFile_DecodeName(localFileName *QByteArray) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_decodeName, uintptr(localFileName.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(localFileName)
	return pgStr(__r, __n)
}

func QFile_DecodeNameWithLocalFileName(localFileName string) string {
	localFileName_b := append([]byte(localFileName), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_decodeName2, uintptr(unsafe.Pointer(&localFileName_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(localFileName_b)
	return pgStr(__r, __n)
}

func (this *QFile) Exists() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_exists, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFile_Exists(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_exists2, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QFile) SymLinkTarget() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_symLinkTarget, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QFile_SymLinkTarget(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_symLinkTarget2, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func (this *QFile) Remove() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_remove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFile_Remove(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_remove2, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QFile_SupportsMoveToTrash() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_supportsMoveToTrash)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QFile) MoveToTrash() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_moveToTrash, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFile_MoveToTrash(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_moveToTrash2, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func (this *QFile) Rename(newName string) bool {
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_rename, this.h, newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QFile_Rename(oldName string, newName string) bool {
	var oldName_p uintptr
	if len(oldName) > 0 {
		oldName_p = uintptr(unsafe.Pointer(unsafe.StringData(oldName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_rename2, oldName_p, uintptr(len(oldName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(oldName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func (this *QFile) Link(newName string) bool {
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_link, this.h, newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QFile_Link(fileName string, newName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_link2, fileName_p, uintptr(len(fileName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func (this *QFile) Copy(newName string) bool {
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_copy, this.h, newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QFile_Copy(fileName string, newName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_copy2, fileName_p, uintptr(len(fileName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func (this *QFile) Open(flags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) Open2(flags QIODeviceBase__OpenModeFlag, permissions QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open2, this.h, uintptr(flags), uintptr(permissions))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) Open3(fd int, ioFlags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open3, this.h, uintptr(fd), uintptr(ioFlags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFile_Resize(filename string, sz int64) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_resize, filename_p, uintptr(len(filename)), uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return __r != 0
}

func QFile_Permissions(filename string) QFileDevice__Permission {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_permissions, filename_p, uintptr(len(filename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return QFileDevice__Permission(__r)
}

func QFile_SetPermissions(filename string, permissionSpec QFileDevice__Permission) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_setPermissions, filename_p, uintptr(len(filename)), uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return __r != 0
}

func (this *QFile) SetOpenMode(openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QFile_setOpenMode, this.h, uintptr(openMode))
	runtime.KeepAlive(this)
}

func (this *QFile) SetErrorString(errorString string) {
	var errorString_p uintptr
	if len(errorString) > 0 {
		errorString_p = uintptr(unsafe.Pointer(unsafe.StringData(errorString)))
	}
	purego.SyscallN(addr_QFile_setErrorString, this.h, errorString_p, uintptr(len(errorString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(errorString)
}

func QFile_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFile_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFile_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFile_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFile) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFile_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFile) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFile_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFile) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFile_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFile) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFile_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFile_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFile_MoveToTrash2(fileName string, pathInTrash *QString) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QFile_moveToTrash3, fileName_p, uintptr(len(fileName)), uintptr(pathInTrash.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(pathInTrash)
	return __r != 0
}

func (this *QFile) Open4(fd int, ioFlags QIODeviceBase__OpenModeFlag, handleFlags QFileDevice__FileHandleFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open4, this.h, uintptr(fd), uintptr(ioFlags), uintptr(handleFlags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFile_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFile_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFile_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFile_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFile) OpenDefault(flags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open_default, this.h, uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) Open2Default(flags QIODeviceBase__OpenModeFlag, permissions QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open_default2, this.h, uintptr(flags), uintptr(permissions))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) Open3Default(fd int, ioFlags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open_default3, this.h, uintptr(fd), uintptr(ioFlags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) Open4Default(fd int, ioFlags QIODeviceBase__OpenModeFlag, handleFlags QFileDevice__FileHandleFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_open_default4, this.h, uintptr(fd), uintptr(ioFlags), uintptr(handleFlags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFile_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFile) FileNameDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFile_fileName_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFile) SizeDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QFile_size_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFile) ResizeDefault(sz int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_resize_default, this.h, uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) PermissionsDefault() QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QFile_permissions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__Permission(__r)
}

func (this *QFile) SetPermissionsDefault(permissionSpec QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_setPermissions_default, this.h, uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) CloseDefault() {
	purego.SyscallN(addr_QFile_close_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFile) IsSequentialDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_isSequential_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) PosDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QFile_pos_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFile) SeekDefault(offset int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_seek_default, this.h, uintptr(offset))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) AtEndDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_atEnd_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) ResetDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_reset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) BytesAvailableDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QFile_bytesAvailable_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFile) BytesToWriteDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QFile_bytesToWrite_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFile) CanReadLineDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_canReadLine_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) WaitForReadyReadDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_waitForReadyRead_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) WaitForBytesWrittenDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_waitForBytesWritten_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFile) SkipDataDefault(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QFile_skipData_default, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QFile) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFile) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFile_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFile) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFile_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFile) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFile_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFile) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFile_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFile) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFile_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFile) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFile_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFile) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFile_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QFile_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QFile_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QFile_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFile_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFile_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFile_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFile_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFile_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFile) Disconnect(handle int64) {
	pg_QFile__disconnect(handle)
	unregisterCallback(handle)
}

func QFile_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 8848495542063941256:
		return 1
	case -554057428128861975:
		return 2
	case 5614464919154503228:
		return 3
	case -4890198374319873567:
		return 4
	case -7517927236947282013:
		return 5
	case 420407675554349589:
		return 6
	case -7006921981232612925:
		return 7
	case 8888098400087344862:
		return 8
	case 8618582476283917417:
		return 9
	case 3241140482992819217:
		return 10
	case -1487528724786208339:
		return 11
	case 1086335023529244512:
		return 12
	case 4814449630054007943:
		return 13
	case 1024747515116322948:
		return 14
	case -1250682508954504761:
		return 15
	case 8356976193102722774:
		return 16
	case 7726833620054767889:
		return 17
	case 8178738247667652082:
		return 18
	case 1237236456312824991:
		return 19
	case 5217623514630858297:
		return 20
	case -5409927581812374498:
		return 21
	case -1185315367526857833:
		return 22
	case 3769868229123868474:
		return 23
	case 7379029982170503008:
		return 24
	case -568369772982215962:
		return 25
	}
	return -1
}

func QFile_enable(h uintptr, idx int32) { pg_QFile__set_override(h, idx) }

func (this *QFile) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFile) OnFileName(fn func() string) {
	registerVirt(this.UnsafeGoID(), 8848495542063941256, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "fileName"
}

func (this *QFile) OnOpen(fn func(QIODeviceBase__OpenModeFlag) bool) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func(flags int32) int32 { return b2i(fn(QIODeviceBase__OpenModeFlag(flags))) }) // slot "open"
}

func (this *QFile) OnSize(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() int64 { return fn() }) // slot "size"
}

func (this *QFile) OnResize(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), -4890198374319873567, func(sz int64) int32 { return b2i(fn(sz)) }) // slot "resize"
}

func (this *QFile) OnPermissions(fn func() QFileDevice__Permission) {
	registerVirt(this.UnsafeGoID(), -7517927236947282013, func() int32 { return int32(fn()) }) // slot "permissions"
}

func (this *QFile) OnSetPermissions(fn func(QFileDevice__Permission) bool) {
	registerVirt(this.UnsafeGoID(), 420407675554349589, func(permissionSpec int32) int32 { return b2i(fn(QFileDevice__Permission(permissionSpec))) }) // slot "setPermissions"
}

func (this *QFile) OnClose(fn func()) {
	registerVirt(this.UnsafeGoID(), -7006921981232612925, func() { fn() }) // slot "close"
}

func (this *QFile) OnIsSequential(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8888098400087344862, func() int32 { return b2i(fn()) }) // slot "isSequential"
}

func (this *QFile) OnPos(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 8618582476283917417, func() int64 { return fn() }) // slot "pos"
}

func (this *QFile) OnSeek(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), 3241140482992819217, func(offset int64) int32 { return b2i(fn(offset)) }) // slot "seek"
}

func (this *QFile) OnAtEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1487528724786208339, func() int32 { return b2i(fn()) }) // slot "atEnd"
}

func (this *QFile) OnReset(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() int32 { return b2i(fn()) }) // slot "reset"
}

func (this *QFile) OnBytesAvailable(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 4814449630054007943, func() int64 { return fn() }) // slot "bytesAvailable"
}

func (this *QFile) OnBytesToWrite(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 1024747515116322948, func() int64 { return fn() }) // slot "bytesToWrite"
}

func (this *QFile) OnCanReadLine(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1250682508954504761, func() int32 { return b2i(fn()) }) // slot "canReadLine"
}

func (this *QFile) OnWaitForReadyRead(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 8356976193102722774, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForReadyRead"
}

func (this *QFile) OnWaitForBytesWritten(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 7726833620054767889, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForBytesWritten"
}

func (this *QFile) OnSkipData(fn func(int64) int64) {
	registerVirt(this.UnsafeGoID(), 8178738247667652082, func(maxSize int64) int64 { return fn(maxSize) }) // slot "skipData"
}

func (this *QFile) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFile) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFile) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFile) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFile) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFile) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFile) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFile) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFile_delete(this.h)
	}
	this.UnsafeRelease()
}
