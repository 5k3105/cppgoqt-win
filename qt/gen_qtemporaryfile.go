package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTemporaryFile_chain uintptr
	addr_QTemporaryFile_new uintptr
	addr_QTemporaryFile_new2 uintptr
	addr_QTemporaryFile_new3 uintptr
	addr_QTemporaryFile_new4 uintptr
	addr_QTemporaryFile_tr uintptr
	addr_QTemporaryFile_autoRemove uintptr
	addr_QTemporaryFile_setAutoRemove uintptr
	addr_QTemporaryFile_open uintptr
	addr_QTemporaryFile_fileTemplate uintptr
	addr_QTemporaryFile_setFileTemplate uintptr
	addr_QTemporaryFile_rename uintptr
	addr_QTemporaryFile_renameOverwrite uintptr
	addr_QTemporaryFile_createNativeFile uintptr
	addr_QTemporaryFile_createNativeFile2 uintptr
	addr_QTemporaryFile_encodeName uintptr
	addr_QTemporaryFile_decodeName uintptr
	addr_QTemporaryFile_decodeName2 uintptr
	addr_QTemporaryFile_exists uintptr
	addr_QTemporaryFile_symLinkTarget uintptr
	addr_QTemporaryFile_remove uintptr
	addr_QTemporaryFile_supportsMoveToTrash uintptr
	addr_QTemporaryFile_moveToTrash uintptr
	addr_QTemporaryFile_link uintptr
	addr_QTemporaryFile_copy uintptr
	addr_QTemporaryFile_resize uintptr
	addr_QTemporaryFile_permissions uintptr
	addr_QTemporaryFile_setPermissions uintptr
	addr_QTemporaryFile_setOpenMode uintptr
	addr_QTemporaryFile_setErrorString uintptr
	addr_QTemporaryFile_connect uintptr
	addr_QTemporaryFile_connect2 uintptr
	addr_QTemporaryFile_sender uintptr
	addr_QTemporaryFile_senderSignalIndex uintptr
	addr_QTemporaryFile_receivers uintptr
	addr_QTemporaryFile_isSignalConnected uintptr
	addr_QTemporaryFile_tr2 uintptr
	addr_QTemporaryFile_tr3 uintptr
	addr_QTemporaryFile_moveToTrash2 uintptr
	addr_QTemporaryFile_connect3 uintptr
	addr_QTemporaryFile_connect4 uintptr
	addr_QTemporaryFile_open_default uintptr
	addr_QTemporaryFile_metaObject_default uintptr
	addr_QTemporaryFile_fileName_default uintptr
	addr_QTemporaryFile_size_default uintptr
	addr_QTemporaryFile_resize_default uintptr
	addr_QTemporaryFile_permissions_default uintptr
	addr_QTemporaryFile_setPermissions_default uintptr
	addr_QTemporaryFile_close_default uintptr
	addr_QTemporaryFile_isSequential_default uintptr
	addr_QTemporaryFile_pos_default uintptr
	addr_QTemporaryFile_seek_default uintptr
	addr_QTemporaryFile_atEnd_default uintptr
	addr_QTemporaryFile_reset_default uintptr
	addr_QTemporaryFile_bytesAvailable_default uintptr
	addr_QTemporaryFile_bytesToWrite_default uintptr
	addr_QTemporaryFile_canReadLine_default uintptr
	addr_QTemporaryFile_waitForReadyRead_default uintptr
	addr_QTemporaryFile_waitForBytesWritten_default uintptr
	addr_QTemporaryFile_skipData_default uintptr
	addr_QTemporaryFile_event_default uintptr
	addr_QTemporaryFile_eventFilter_default uintptr
	addr_QTemporaryFile_timerEvent_default uintptr
	addr_QTemporaryFile_childEvent_default uintptr
	addr_QTemporaryFile_customEvent_default uintptr
	addr_QTemporaryFile_connectNotify_default uintptr
	addr_QTemporaryFile_disconnectNotify_default uintptr
	pg_QTemporaryFile_connect_readyRead func(uintptr, int64)
	pg_QTemporaryFile_connect_channelReadyRead func(uintptr, int64)
	pg_QTemporaryFile_connect_bytesWritten func(uintptr, int64)
	pg_QTemporaryFile_connect_channelBytesWritten func(uintptr, int64)
	pg_QTemporaryFile_connect_aboutToClose func(uintptr, int64)
	pg_QTemporaryFile_connect_readChannelFinished func(uintptr, int64)
	pg_QTemporaryFile_connect_destroyed func(uintptr, int64)
	pg_QTemporaryFile_connect_destroyed2 func(uintptr, int64)
	pg_QTemporaryFile_connect_objectNameChanged func(uintptr, int64)
	pg_QTemporaryFile__disconnect func(int64)
	pg_QTemporaryFile__set_override func(uintptr, int32)
	pg_QTemporaryFile_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTemporaryFile_chain, "QTemporaryFile_chain")
	symaddr(&addr_QTemporaryFile_new, "QTemporaryFile_new")
	symaddr(&addr_QTemporaryFile_new2, "QTemporaryFile_new2")
	symaddr(&addr_QTemporaryFile_new3, "QTemporaryFile_new3")
	symaddr(&addr_QTemporaryFile_new4, "QTemporaryFile_new4")
	symaddr(&addr_QTemporaryFile_tr, "QTemporaryFile_tr")
	symaddr(&addr_QTemporaryFile_autoRemove, "QTemporaryFile_autoRemove")
	symaddr(&addr_QTemporaryFile_setAutoRemove, "QTemporaryFile_setAutoRemove")
	symaddr(&addr_QTemporaryFile_open, "QTemporaryFile_open")
	symaddr(&addr_QTemporaryFile_fileTemplate, "QTemporaryFile_fileTemplate")
	symaddr(&addr_QTemporaryFile_setFileTemplate, "QTemporaryFile_setFileTemplate")
	symaddr(&addr_QTemporaryFile_rename, "QTemporaryFile_rename")
	symaddr(&addr_QTemporaryFile_renameOverwrite, "QTemporaryFile_renameOverwrite")
	symaddr(&addr_QTemporaryFile_createNativeFile, "QTemporaryFile_createNativeFile")
	symaddr(&addr_QTemporaryFile_createNativeFile2, "QTemporaryFile_createNativeFile2")
	symaddr(&addr_QTemporaryFile_encodeName, "QTemporaryFile_encodeName")
	symaddr(&addr_QTemporaryFile_decodeName, "QTemporaryFile_decodeName")
	symaddr(&addr_QTemporaryFile_decodeName2, "QTemporaryFile_decodeName2")
	symaddr(&addr_QTemporaryFile_exists, "QTemporaryFile_exists")
	symaddr(&addr_QTemporaryFile_symLinkTarget, "QTemporaryFile_symLinkTarget")
	symaddr(&addr_QTemporaryFile_remove, "QTemporaryFile_remove")
	symaddr(&addr_QTemporaryFile_supportsMoveToTrash, "QTemporaryFile_supportsMoveToTrash")
	symaddr(&addr_QTemporaryFile_moveToTrash, "QTemporaryFile_moveToTrash")
	symaddr(&addr_QTemporaryFile_link, "QTemporaryFile_link")
	symaddr(&addr_QTemporaryFile_copy, "QTemporaryFile_copy")
	symaddr(&addr_QTemporaryFile_resize, "QTemporaryFile_resize")
	symaddr(&addr_QTemporaryFile_permissions, "QTemporaryFile_permissions")
	symaddr(&addr_QTemporaryFile_setPermissions, "QTemporaryFile_setPermissions")
	symaddr(&addr_QTemporaryFile_setOpenMode, "QTemporaryFile_setOpenMode")
	symaddr(&addr_QTemporaryFile_setErrorString, "QTemporaryFile_setErrorString")
	symaddr(&addr_QTemporaryFile_connect, "QTemporaryFile_connect")
	symaddr(&addr_QTemporaryFile_connect2, "QTemporaryFile_connect2")
	symaddr(&addr_QTemporaryFile_sender, "QTemporaryFile_sender")
	symaddr(&addr_QTemporaryFile_senderSignalIndex, "QTemporaryFile_senderSignalIndex")
	symaddr(&addr_QTemporaryFile_receivers, "QTemporaryFile_receivers")
	symaddr(&addr_QTemporaryFile_isSignalConnected, "QTemporaryFile_isSignalConnected")
	symaddr(&addr_QTemporaryFile_tr2, "QTemporaryFile_tr2")
	symaddr(&addr_QTemporaryFile_tr3, "QTemporaryFile_tr3")
	symaddr(&addr_QTemporaryFile_moveToTrash2, "QTemporaryFile_moveToTrash2")
	symaddr(&addr_QTemporaryFile_connect3, "QTemporaryFile_connect3")
	symaddr(&addr_QTemporaryFile_connect4, "QTemporaryFile_connect4")
	symaddr(&addr_QTemporaryFile_open_default, "QTemporaryFile_open_default")
	symaddr(&addr_QTemporaryFile_metaObject_default, "QTemporaryFile_metaObject_default")
	symaddr(&addr_QTemporaryFile_fileName_default, "QTemporaryFile_fileName_default")
	symaddr(&addr_QTemporaryFile_size_default, "QTemporaryFile_size_default")
	symaddr(&addr_QTemporaryFile_resize_default, "QTemporaryFile_resize_default")
	symaddr(&addr_QTemporaryFile_permissions_default, "QTemporaryFile_permissions_default")
	symaddr(&addr_QTemporaryFile_setPermissions_default, "QTemporaryFile_setPermissions_default")
	symaddr(&addr_QTemporaryFile_close_default, "QTemporaryFile_close_default")
	symaddr(&addr_QTemporaryFile_isSequential_default, "QTemporaryFile_isSequential_default")
	symaddr(&addr_QTemporaryFile_pos_default, "QTemporaryFile_pos_default")
	symaddr(&addr_QTemporaryFile_seek_default, "QTemporaryFile_seek_default")
	symaddr(&addr_QTemporaryFile_atEnd_default, "QTemporaryFile_atEnd_default")
	symaddr(&addr_QTemporaryFile_reset_default, "QTemporaryFile_reset_default")
	symaddr(&addr_QTemporaryFile_bytesAvailable_default, "QTemporaryFile_bytesAvailable_default")
	symaddr(&addr_QTemporaryFile_bytesToWrite_default, "QTemporaryFile_bytesToWrite_default")
	symaddr(&addr_QTemporaryFile_canReadLine_default, "QTemporaryFile_canReadLine_default")
	symaddr(&addr_QTemporaryFile_waitForReadyRead_default, "QTemporaryFile_waitForReadyRead_default")
	symaddr(&addr_QTemporaryFile_waitForBytesWritten_default, "QTemporaryFile_waitForBytesWritten_default")
	symaddr(&addr_QTemporaryFile_skipData_default, "QTemporaryFile_skipData_default")
	symaddr(&addr_QTemporaryFile_event_default, "QTemporaryFile_event_default")
	symaddr(&addr_QTemporaryFile_eventFilter_default, "QTemporaryFile_eventFilter_default")
	symaddr(&addr_QTemporaryFile_timerEvent_default, "QTemporaryFile_timerEvent_default")
	symaddr(&addr_QTemporaryFile_childEvent_default, "QTemporaryFile_childEvent_default")
	symaddr(&addr_QTemporaryFile_customEvent_default, "QTemporaryFile_customEvent_default")
	symaddr(&addr_QTemporaryFile_connectNotify_default, "QTemporaryFile_connectNotify_default")
	symaddr(&addr_QTemporaryFile_disconnectNotify_default, "QTemporaryFile_disconnectNotify_default")
	reg(&pg_QTemporaryFile_connect_readyRead, "QTemporaryFile_connect_readyRead")
	reg(&pg_QTemporaryFile_connect_channelReadyRead, "QTemporaryFile_connect_channelReadyRead")
	reg(&pg_QTemporaryFile_connect_bytesWritten, "QTemporaryFile_connect_bytesWritten")
	reg(&pg_QTemporaryFile_connect_channelBytesWritten, "QTemporaryFile_connect_channelBytesWritten")
	reg(&pg_QTemporaryFile_connect_aboutToClose, "QTemporaryFile_connect_aboutToClose")
	reg(&pg_QTemporaryFile_connect_readChannelFinished, "QTemporaryFile_connect_readChannelFinished")
	reg(&pg_QTemporaryFile_connect_destroyed, "QTemporaryFile_connect_destroyed")
	reg(&pg_QTemporaryFile_connect_destroyed2, "QTemporaryFile_connect_destroyed2")
	reg(&pg_QTemporaryFile_connect_objectNameChanged, "QTemporaryFile_connect_objectNameChanged")
	reg(&pg_QTemporaryFile__disconnect, "QTemporaryFile__disconnect")
	reg(&pg_QTemporaryFile__set_override, "QTemporaryFile__set_override")
	reg(&pg_QTemporaryFile_delete, "QTemporaryFile_delete")
}

type QTemporaryFile struct {
	h uintptr
	*QFile
}

func bindQTemporaryFile(h uintptr) *QTemporaryFile {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QTemporaryFile_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTemporaryFile{h: h, QFile: &QFile{h: __c[0], QFileDevice: &QFileDevice{h: __c[1], QIODevice: &QIODevice{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQTemporaryFile(h uintptr) *QTemporaryFile { return bindQTemporaryFile(h) }

func (this *QTemporaryFile) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTemporaryFile) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTemporaryFile(h unsafe.Pointer) *QTemporaryFile { return bindQTemporaryFile(uintptr(h)) }
func UnsafeNewQTemporaryFileOwned(h unsafe.Pointer) *QTemporaryFile { return newownedQTemporaryFile(uintptr(h)) }

func (this *QTemporaryFile) UnsafeRelease() {
	this.h = 0
	if this.QFile != nil {
		this.QFile.UnsafeRelease()
	}
}

func NewQTemporaryFile() *QTemporaryFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_new, uintptr(__goid))
	__o := newownedQTemporaryFile(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTemporaryFile_virtIdx, QTemporaryFile_enable)
	return __o
}

func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	__goid := nextVirtID()
	var templateName_p uintptr
	if len(templateName) > 0 {
		templateName_p = uintptr(unsafe.Pointer(unsafe.StringData(templateName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_new2, uintptr(__goid), templateName_p, uintptr(len(templateName)))
	__o := newownedQTemporaryFile(__r)
	runtime.KeepAlive(templateName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTemporaryFile_virtIdx, QTemporaryFile_enable)
	return __o
}

func NewQTemporaryFile3(parent *QObject) *QTemporaryFile {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTemporaryFile(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTemporaryFile_virtIdx, QTemporaryFile_enable)
	return __o
}

func NewQTemporaryFile4(templateName string, parent *QObject) *QTemporaryFile {
	__goid := nextVirtID()
	var templateName_p uintptr
	if len(templateName) > 0 {
		templateName_p = uintptr(unsafe.Pointer(unsafe.StringData(templateName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_new4, uintptr(__goid), templateName_p, uintptr(len(templateName)), uintptr(parent.UnsafePointer()))
	__o := newownedQTemporaryFile(__r)
	runtime.KeepAlive(templateName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTemporaryFile_virtIdx, QTemporaryFile_enable)
	return __o
}

func QTemporaryFile_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTemporaryFile) AutoRemove() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_autoRemove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) SetAutoRemove(b bool) {
	purego.SyscallN(addr_QTemporaryFile_setAutoRemove, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTemporaryFile) Open() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_open, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) FileTemplate() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_fileTemplate, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTemporaryFile) SetFileTemplate(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTemporaryFile_setFileTemplate, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTemporaryFile) Rename(newName string) bool {
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_rename, this.h, newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newName)
	return __r != 0
}

func (this *QTemporaryFile) RenameOverwrite(newName string) bool {
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_renameOverwrite, this.h, newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QTemporaryFile_CreateNativeFile(fileName string) *QTemporaryFile {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_createNativeFile, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(fileName)
	return bindQTemporaryFile(__r)
}

func QTemporaryFile_CreateNativeFileWithFile(file *QFile) *QTemporaryFile {
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_createNativeFile2, uintptr(file.UnsafePointer()))
	runtime.KeepAlive(file)
	return bindQTemporaryFile(__r)
}

func QTemporaryFile_EncodeName(fileName string) *QByteArray {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_encodeName, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(fileName)
	return newownedQByteArray(__r)
}

func QTemporaryFile_DecodeName(localFileName *QByteArray) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_decodeName, uintptr(localFileName.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(localFileName)
	return pgStr(__r, __n)
}

func QTemporaryFile_DecodeNameWithLocalFileName(localFileName string) string {
	localFileName_b := append([]byte(localFileName), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_decodeName2, uintptr(unsafe.Pointer(&localFileName_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(localFileName_b)
	return pgStr(__r, __n)
}

func QTemporaryFile_Exists(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_exists, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QTemporaryFile_SymLinkTarget(fileName string) string {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_symLinkTarget, fileName_p, uintptr(len(fileName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(fileName)
	return pgStr(__r, __n)
}

func QTemporaryFile_Remove(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_remove, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QTemporaryFile_SupportsMoveToTrash() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_supportsMoveToTrash)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QTemporaryFile_MoveToTrash(fileName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_moveToTrash, fileName_p, uintptr(len(fileName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	return __r != 0
}

func QTemporaryFile_Link(fileName string, newName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_link, fileName_p, uintptr(len(fileName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QTemporaryFile_Copy(fileName string, newName string) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_copy, fileName_p, uintptr(len(fileName)), newName_p, uintptr(len(newName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(newName)
	return __r != 0
}

func QTemporaryFile_Resize(filename string, sz int64) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_resize, filename_p, uintptr(len(filename)), uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return __r != 0
}

func QTemporaryFile_Permissions(filename string) QFileDevice__Permission {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_permissions, filename_p, uintptr(len(filename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return QFileDevice__Permission(__r)
}

func QTemporaryFile_SetPermissions(filename string, permissionSpec QFileDevice__Permission) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_setPermissions, filename_p, uintptr(len(filename)), uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(filename)
	return __r != 0
}

func (this *QTemporaryFile) SetOpenMode(openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QTemporaryFile_setOpenMode, this.h, uintptr(openMode))
	runtime.KeepAlive(this)
}

func (this *QTemporaryFile) SetErrorString(errorString string) {
	var errorString_p uintptr
	if len(errorString) > 0 {
		errorString_p = uintptr(unsafe.Pointer(unsafe.StringData(errorString)))
	}
	purego.SyscallN(addr_QTemporaryFile_setErrorString, this.h, errorString_p, uintptr(len(errorString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(errorString)
}

func QTemporaryFile_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTemporaryFile_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTemporaryFile) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTemporaryFile) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTemporaryFile) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTemporaryFile) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTemporaryFile_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTemporaryFile_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTemporaryFile_MoveToTrash2(fileName string, pathInTrash *QString) bool {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_moveToTrash2, fileName_p, uintptr(len(fileName)), uintptr(pathInTrash.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(pathInTrash)
	return __r != 0
}

func QTemporaryFile_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTemporaryFile_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTemporaryFile) OpenDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_open_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTemporaryFile) FileNameDefault() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTemporaryFile_fileName_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTemporaryFile) SizeDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_size_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTemporaryFile) ResizeDefault(sz int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_resize_default, this.h, uintptr(sz))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) PermissionsDefault() QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_permissions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileDevice__Permission(__r)
}

func (this *QTemporaryFile) SetPermissionsDefault(permissionSpec QFileDevice__Permission) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_setPermissions_default, this.h, uintptr(permissionSpec))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) CloseDefault() {
	purego.SyscallN(addr_QTemporaryFile_close_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTemporaryFile) IsSequentialDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_isSequential_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) PosDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_pos_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTemporaryFile) SeekDefault(offset int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_seek_default, this.h, uintptr(offset))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) AtEndDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_atEnd_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) ResetDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_reset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) BytesAvailableDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_bytesAvailable_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTemporaryFile) BytesToWriteDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_bytesToWrite_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTemporaryFile) CanReadLineDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_canReadLine_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) WaitForReadyReadDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_waitForReadyRead_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) WaitForBytesWrittenDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_waitForBytesWritten_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTemporaryFile) SkipDataDefault(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_skipData_default, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTemporaryFile) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTemporaryFile) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTemporaryFile_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTemporaryFile) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTemporaryFile_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTemporaryFile) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTemporaryFile_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTemporaryFile) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTemporaryFile_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTemporaryFile) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTemporaryFile_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTemporaryFile) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTemporaryFile_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTemporaryFile) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTemporaryFile_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QTemporaryFile_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QTemporaryFile_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QTemporaryFile_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTemporaryFile_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTemporaryFile_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTemporaryFile_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTemporaryFile_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTemporaryFile_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTemporaryFile) Disconnect(handle int64) {
	pg_QTemporaryFile__disconnect(handle)
	unregisterCallback(handle)
}

func QTemporaryFile_virtIdx(slot int64) int32 {
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

func QTemporaryFile_enable(h uintptr, idx int32) { pg_QTemporaryFile__set_override(h, idx) }

func (this *QTemporaryFile) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTemporaryFile) OnFileName(fn func() string) {
	registerVirt(this.UnsafeGoID(), 8848495542063941256, func() (uintptr, int64) {
		__s := fn()
		return pgCString(__s), int64(len(__s))
	}) // slot "fileName"
}

func (this *QTemporaryFile) OnOpen(fn func(QIODeviceBase__OpenModeFlag) bool) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func(flags int32) int32 { return b2i(fn(QIODeviceBase__OpenModeFlag(flags))) }) // slot "open"
}

func (this *QTemporaryFile) OnSize(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() int64 { return fn() }) // slot "size"
}

func (this *QTemporaryFile) OnResize(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), -4890198374319873567, func(sz int64) int32 { return b2i(fn(sz)) }) // slot "resize"
}

func (this *QTemporaryFile) OnPermissions(fn func() QFileDevice__Permission) {
	registerVirt(this.UnsafeGoID(), -7517927236947282013, func() int32 { return int32(fn()) }) // slot "permissions"
}

func (this *QTemporaryFile) OnSetPermissions(fn func(QFileDevice__Permission) bool) {
	registerVirt(this.UnsafeGoID(), 420407675554349589, func(permissionSpec int32) int32 { return b2i(fn(QFileDevice__Permission(permissionSpec))) }) // slot "setPermissions"
}

func (this *QTemporaryFile) OnClose(fn func()) {
	registerVirt(this.UnsafeGoID(), -7006921981232612925, func() { fn() }) // slot "close"
}

func (this *QTemporaryFile) OnIsSequential(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8888098400087344862, func() int32 { return b2i(fn()) }) // slot "isSequential"
}

func (this *QTemporaryFile) OnPos(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 8618582476283917417, func() int64 { return fn() }) // slot "pos"
}

func (this *QTemporaryFile) OnSeek(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), 3241140482992819217, func(offset int64) int32 { return b2i(fn(offset)) }) // slot "seek"
}

func (this *QTemporaryFile) OnAtEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1487528724786208339, func() int32 { return b2i(fn()) }) // slot "atEnd"
}

func (this *QTemporaryFile) OnReset(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() int32 { return b2i(fn()) }) // slot "reset"
}

func (this *QTemporaryFile) OnBytesAvailable(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 4814449630054007943, func() int64 { return fn() }) // slot "bytesAvailable"
}

func (this *QTemporaryFile) OnBytesToWrite(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 1024747515116322948, func() int64 { return fn() }) // slot "bytesToWrite"
}

func (this *QTemporaryFile) OnCanReadLine(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1250682508954504761, func() int32 { return b2i(fn()) }) // slot "canReadLine"
}

func (this *QTemporaryFile) OnWaitForReadyRead(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 8356976193102722774, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForReadyRead"
}

func (this *QTemporaryFile) OnWaitForBytesWritten(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 7726833620054767889, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForBytesWritten"
}

func (this *QTemporaryFile) OnSkipData(fn func(int64) int64) {
	registerVirt(this.UnsafeGoID(), 8178738247667652082, func(maxSize int64) int64 { return fn(maxSize) }) // slot "skipData"
}

func (this *QTemporaryFile) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTemporaryFile) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTemporaryFile) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTemporaryFile) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTemporaryFile) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTemporaryFile) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTemporaryFile) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTemporaryFile) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTemporaryFile_delete(this.h)
	}
	this.UnsafeRelease()
}
