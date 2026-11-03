package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProcess_chain uintptr
	addr_QProcess_new uintptr
	addr_QProcess_new2 uintptr
	addr_QProcess_tr uintptr
	addr_QProcess_start uintptr
	addr_QProcess_start2 uintptr
	addr_QProcess_startCommand uintptr
	addr_QProcess_startDetached uintptr
	addr_QProcess_program uintptr
	addr_QProcess_setProgram uintptr
	addr_QProcess_arguments uintptr
	addr_QProcess_setArguments uintptr
	addr_QProcess_processChannelMode uintptr
	addr_QProcess_setProcessChannelMode uintptr
	addr_QProcess_inputChannelMode uintptr
	addr_QProcess_setInputChannelMode uintptr
	addr_QProcess_readChannel uintptr
	addr_QProcess_setReadChannel uintptr
	addr_QProcess_closeReadChannel uintptr
	addr_QProcess_closeWriteChannel uintptr
	addr_QProcess_setStandardInputFile uintptr
	addr_QProcess_setStandardOutputFile uintptr
	addr_QProcess_setStandardErrorFile uintptr
	addr_QProcess_setStandardOutputProcess uintptr
	addr_QProcess_nativeArguments uintptr
	addr_QProcess_setNativeArguments uintptr
	addr_QProcess_workingDirectory uintptr
	addr_QProcess_setWorkingDirectory uintptr
	addr_QProcess_setEnvironment uintptr
	addr_QProcess_environment uintptr
	addr_QProcess_setProcessEnvironment uintptr
	addr_QProcess_processEnvironment uintptr
	addr_QProcess_error uintptr
	addr_QProcess_state uintptr
	addr_QProcess_processId uintptr
	addr_QProcess_waitForStarted uintptr
	addr_QProcess_waitForFinished uintptr
	addr_QProcess_readAllStandardOutput uintptr
	addr_QProcess_readAllStandardError uintptr
	addr_QProcess_exitCode uintptr
	addr_QProcess_exitStatus uintptr
	addr_QProcess_execute uintptr
	addr_QProcess_startDetached2 uintptr
	addr_QProcess_systemEnvironment uintptr
	addr_QProcess_nullDevice uintptr
	addr_QProcess_splitCommand uintptr
	addr_QProcess_terminate uintptr
	addr_QProcess_kill uintptr
	addr_QProcess_finished uintptr
	addr_QProcess_errorOccurred uintptr
	addr_QProcess_setProcessState uintptr
	addr_QProcess_setOpenMode uintptr
	addr_QProcess_setErrorString uintptr
	addr_QProcess_connect uintptr
	addr_QProcess_connect2 uintptr
	addr_QProcess_sender uintptr
	addr_QProcess_senderSignalIndex uintptr
	addr_QProcess_receivers uintptr
	addr_QProcess_isSignalConnected uintptr
	addr_QProcess_tr2 uintptr
	addr_QProcess_tr3 uintptr
	addr_QProcess_start3 uintptr
	addr_QProcess_start4 uintptr
	addr_QProcess_start5 uintptr
	addr_QProcess_startCommand2 uintptr
	addr_QProcess_startDetached3 uintptr
	addr_QProcess_setStandardOutputFile2 uintptr
	addr_QProcess_setStandardErrorFile2 uintptr
	addr_QProcess_waitForStarted2 uintptr
	addr_QProcess_waitForFinished2 uintptr
	addr_QProcess_execute2 uintptr
	addr_QProcess_startDetached4 uintptr
	addr_QProcess_startDetached5 uintptr
	addr_QProcess_startDetached6 uintptr
	addr_QProcess_finished2 uintptr
	addr_QProcess_connect3 uintptr
	addr_QProcess_connect4 uintptr
	addr_QProcess_metaObject_default uintptr
	addr_QProcess_open_default uintptr
	addr_QProcess_waitForReadyRead_default uintptr
	addr_QProcess_waitForBytesWritten_default uintptr
	addr_QProcess_bytesToWrite_default uintptr
	addr_QProcess_isSequential_default uintptr
	addr_QProcess_close_default uintptr
	addr_QProcess_pos_default uintptr
	addr_QProcess_size_default uintptr
	addr_QProcess_seek_default uintptr
	addr_QProcess_atEnd_default uintptr
	addr_QProcess_reset_default uintptr
	addr_QProcess_bytesAvailable_default uintptr
	addr_QProcess_canReadLine_default uintptr
	addr_QProcess_skipData_default uintptr
	addr_QProcess_event_default uintptr
	addr_QProcess_eventFilter_default uintptr
	addr_QProcess_timerEvent_default uintptr
	addr_QProcess_childEvent_default uintptr
	addr_QProcess_customEvent_default uintptr
	addr_QProcess_connectNotify_default uintptr
	addr_QProcess_disconnectNotify_default uintptr
	pg_QProcess_connect_started func(uintptr, int64)
	pg_QProcess_connect_finished func(uintptr, int64)
	pg_QProcess_connect_finished2 func(uintptr, int64)
	pg_QProcess_connect_errorOccurred func(uintptr, int64)
	pg_QProcess_connect_stateChanged func(uintptr, int64)
	pg_QProcess_connect_readyReadStandardOutput func(uintptr, int64)
	pg_QProcess_connect_readyReadStandardError func(uintptr, int64)
	pg_QProcess_connect_readyRead func(uintptr, int64)
	pg_QProcess_connect_channelReadyRead func(uintptr, int64)
	pg_QProcess_connect_bytesWritten func(uintptr, int64)
	pg_QProcess_connect_channelBytesWritten func(uintptr, int64)
	pg_QProcess_connect_aboutToClose func(uintptr, int64)
	pg_QProcess_connect_readChannelFinished func(uintptr, int64)
	pg_QProcess_connect_destroyed func(uintptr, int64)
	pg_QProcess_connect_destroyed2 func(uintptr, int64)
	pg_QProcess_connect_objectNameChanged func(uintptr, int64)
	pg_QProcess__disconnect func(int64)
	pg_QProcess__set_override func(uintptr, int32)
	pg_QProcess_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProcess_chain, "QProcess_chain")
	symaddr(&addr_QProcess_new, "QProcess_new")
	symaddr(&addr_QProcess_new2, "QProcess_new2")
	symaddr(&addr_QProcess_tr, "QProcess_tr")
	symaddr(&addr_QProcess_start, "QProcess_start")
	symaddr(&addr_QProcess_start2, "QProcess_start2")
	symaddr(&addr_QProcess_startCommand, "QProcess_startCommand")
	symaddr(&addr_QProcess_startDetached, "QProcess_startDetached")
	symaddr(&addr_QProcess_program, "QProcess_program")
	symaddr(&addr_QProcess_setProgram, "QProcess_setProgram")
	symaddr(&addr_QProcess_arguments, "QProcess_arguments")
	symaddr(&addr_QProcess_setArguments, "QProcess_setArguments")
	symaddr(&addr_QProcess_processChannelMode, "QProcess_processChannelMode")
	symaddr(&addr_QProcess_setProcessChannelMode, "QProcess_setProcessChannelMode")
	symaddr(&addr_QProcess_inputChannelMode, "QProcess_inputChannelMode")
	symaddr(&addr_QProcess_setInputChannelMode, "QProcess_setInputChannelMode")
	symaddr(&addr_QProcess_readChannel, "QProcess_readChannel")
	symaddr(&addr_QProcess_setReadChannel, "QProcess_setReadChannel")
	symaddr(&addr_QProcess_closeReadChannel, "QProcess_closeReadChannel")
	symaddr(&addr_QProcess_closeWriteChannel, "QProcess_closeWriteChannel")
	symaddr(&addr_QProcess_setStandardInputFile, "QProcess_setStandardInputFile")
	symaddr(&addr_QProcess_setStandardOutputFile, "QProcess_setStandardOutputFile")
	symaddr(&addr_QProcess_setStandardErrorFile, "QProcess_setStandardErrorFile")
	symaddr(&addr_QProcess_setStandardOutputProcess, "QProcess_setStandardOutputProcess")
	symaddr(&addr_QProcess_nativeArguments, "QProcess_nativeArguments")
	symaddr(&addr_QProcess_setNativeArguments, "QProcess_setNativeArguments")
	symaddr(&addr_QProcess_workingDirectory, "QProcess_workingDirectory")
	symaddr(&addr_QProcess_setWorkingDirectory, "QProcess_setWorkingDirectory")
	symaddr(&addr_QProcess_setEnvironment, "QProcess_setEnvironment")
	symaddr(&addr_QProcess_environment, "QProcess_environment")
	symaddr(&addr_QProcess_setProcessEnvironment, "QProcess_setProcessEnvironment")
	symaddr(&addr_QProcess_processEnvironment, "QProcess_processEnvironment")
	symaddr(&addr_QProcess_error, "QProcess_error")
	symaddr(&addr_QProcess_state, "QProcess_state")
	symaddr(&addr_QProcess_processId, "QProcess_processId")
	symaddr(&addr_QProcess_waitForStarted, "QProcess_waitForStarted")
	symaddr(&addr_QProcess_waitForFinished, "QProcess_waitForFinished")
	symaddr(&addr_QProcess_readAllStandardOutput, "QProcess_readAllStandardOutput")
	symaddr(&addr_QProcess_readAllStandardError, "QProcess_readAllStandardError")
	symaddr(&addr_QProcess_exitCode, "QProcess_exitCode")
	symaddr(&addr_QProcess_exitStatus, "QProcess_exitStatus")
	symaddr(&addr_QProcess_execute, "QProcess_execute")
	symaddr(&addr_QProcess_startDetached2, "QProcess_startDetached2")
	symaddr(&addr_QProcess_systemEnvironment, "QProcess_systemEnvironment")
	symaddr(&addr_QProcess_nullDevice, "QProcess_nullDevice")
	symaddr(&addr_QProcess_splitCommand, "QProcess_splitCommand")
	symaddr(&addr_QProcess_terminate, "QProcess_terminate")
	symaddr(&addr_QProcess_kill, "QProcess_kill")
	symaddr(&addr_QProcess_finished, "QProcess_finished")
	symaddr(&addr_QProcess_errorOccurred, "QProcess_errorOccurred")
	symaddr(&addr_QProcess_setProcessState, "QProcess_setProcessState")
	symaddr(&addr_QProcess_setOpenMode, "QProcess_setOpenMode")
	symaddr(&addr_QProcess_setErrorString, "QProcess_setErrorString")
	symaddr(&addr_QProcess_connect, "QProcess_connect")
	symaddr(&addr_QProcess_connect2, "QProcess_connect2")
	symaddr(&addr_QProcess_sender, "QProcess_sender")
	symaddr(&addr_QProcess_senderSignalIndex, "QProcess_senderSignalIndex")
	symaddr(&addr_QProcess_receivers, "QProcess_receivers")
	symaddr(&addr_QProcess_isSignalConnected, "QProcess_isSignalConnected")
	symaddr(&addr_QProcess_tr2, "QProcess_tr2")
	symaddr(&addr_QProcess_tr3, "QProcess_tr3")
	symaddr(&addr_QProcess_start3, "QProcess_start3")
	symaddr(&addr_QProcess_start4, "QProcess_start4")
	symaddr(&addr_QProcess_start5, "QProcess_start5")
	symaddr(&addr_QProcess_startCommand2, "QProcess_startCommand2")
	symaddr(&addr_QProcess_startDetached3, "QProcess_startDetached3")
	symaddr(&addr_QProcess_setStandardOutputFile2, "QProcess_setStandardOutputFile2")
	symaddr(&addr_QProcess_setStandardErrorFile2, "QProcess_setStandardErrorFile2")
	symaddr(&addr_QProcess_waitForStarted2, "QProcess_waitForStarted2")
	symaddr(&addr_QProcess_waitForFinished2, "QProcess_waitForFinished2")
	symaddr(&addr_QProcess_execute2, "QProcess_execute2")
	symaddr(&addr_QProcess_startDetached4, "QProcess_startDetached4")
	symaddr(&addr_QProcess_startDetached5, "QProcess_startDetached5")
	symaddr(&addr_QProcess_startDetached6, "QProcess_startDetached6")
	symaddr(&addr_QProcess_finished2, "QProcess_finished2")
	symaddr(&addr_QProcess_connect3, "QProcess_connect3")
	symaddr(&addr_QProcess_connect4, "QProcess_connect4")
	symaddr(&addr_QProcess_metaObject_default, "QProcess_metaObject_default")
	symaddr(&addr_QProcess_open_default, "QProcess_open_default")
	symaddr(&addr_QProcess_waitForReadyRead_default, "QProcess_waitForReadyRead_default")
	symaddr(&addr_QProcess_waitForBytesWritten_default, "QProcess_waitForBytesWritten_default")
	symaddr(&addr_QProcess_bytesToWrite_default, "QProcess_bytesToWrite_default")
	symaddr(&addr_QProcess_isSequential_default, "QProcess_isSequential_default")
	symaddr(&addr_QProcess_close_default, "QProcess_close_default")
	symaddr(&addr_QProcess_pos_default, "QProcess_pos_default")
	symaddr(&addr_QProcess_size_default, "QProcess_size_default")
	symaddr(&addr_QProcess_seek_default, "QProcess_seek_default")
	symaddr(&addr_QProcess_atEnd_default, "QProcess_atEnd_default")
	symaddr(&addr_QProcess_reset_default, "QProcess_reset_default")
	symaddr(&addr_QProcess_bytesAvailable_default, "QProcess_bytesAvailable_default")
	symaddr(&addr_QProcess_canReadLine_default, "QProcess_canReadLine_default")
	symaddr(&addr_QProcess_skipData_default, "QProcess_skipData_default")
	symaddr(&addr_QProcess_event_default, "QProcess_event_default")
	symaddr(&addr_QProcess_eventFilter_default, "QProcess_eventFilter_default")
	symaddr(&addr_QProcess_timerEvent_default, "QProcess_timerEvent_default")
	symaddr(&addr_QProcess_childEvent_default, "QProcess_childEvent_default")
	symaddr(&addr_QProcess_customEvent_default, "QProcess_customEvent_default")
	symaddr(&addr_QProcess_connectNotify_default, "QProcess_connectNotify_default")
	symaddr(&addr_QProcess_disconnectNotify_default, "QProcess_disconnectNotify_default")
	reg(&pg_QProcess_connect_started, "QProcess_connect_started")
	reg(&pg_QProcess_connect_finished, "QProcess_connect_finished")
	reg(&pg_QProcess_connect_finished2, "QProcess_connect_finished2")
	reg(&pg_QProcess_connect_errorOccurred, "QProcess_connect_errorOccurred")
	reg(&pg_QProcess_connect_stateChanged, "QProcess_connect_stateChanged")
	reg(&pg_QProcess_connect_readyReadStandardOutput, "QProcess_connect_readyReadStandardOutput")
	reg(&pg_QProcess_connect_readyReadStandardError, "QProcess_connect_readyReadStandardError")
	reg(&pg_QProcess_connect_readyRead, "QProcess_connect_readyRead")
	reg(&pg_QProcess_connect_channelReadyRead, "QProcess_connect_channelReadyRead")
	reg(&pg_QProcess_connect_bytesWritten, "QProcess_connect_bytesWritten")
	reg(&pg_QProcess_connect_channelBytesWritten, "QProcess_connect_channelBytesWritten")
	reg(&pg_QProcess_connect_aboutToClose, "QProcess_connect_aboutToClose")
	reg(&pg_QProcess_connect_readChannelFinished, "QProcess_connect_readChannelFinished")
	reg(&pg_QProcess_connect_destroyed, "QProcess_connect_destroyed")
	reg(&pg_QProcess_connect_destroyed2, "QProcess_connect_destroyed2")
	reg(&pg_QProcess_connect_objectNameChanged, "QProcess_connect_objectNameChanged")
	reg(&pg_QProcess__disconnect, "QProcess__disconnect")
	reg(&pg_QProcess__set_override, "QProcess__set_override")
	reg(&pg_QProcess_delete, "QProcess_delete")
}

type QProcess struct {
	h uintptr
	*QIODevice
}

func bindQProcess(h uintptr) *QProcess {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QProcess_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QProcess{h: h, QIODevice: &QIODevice{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQProcess(h uintptr) *QProcess { return bindQProcess(h) }

func (this *QProcess) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProcess) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProcess(h unsafe.Pointer) *QProcess { return bindQProcess(uintptr(h)) }
func UnsafeNewQProcessOwned(h unsafe.Pointer) *QProcess { return newownedQProcess(uintptr(h)) }

func (this *QProcess) UnsafeRelease() {
	this.h = 0
	if this.QIODevice != nil {
		this.QIODevice.UnsafeRelease()
	}
}

func NewQProcess() *QProcess {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProcess_new, uintptr(__goid))
	__o := newownedQProcess(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProcess_virtIdx, QProcess_enable)
	return __o
}

func NewQProcess2(parent *QObject) *QProcess {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QProcess_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQProcess(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QProcess_virtIdx, QProcess_enable)
	return __o
}

func QProcess_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QProcess) Start(program string) {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	purego.SyscallN(addr_QProcess_start, this.h, program_p, uintptr(len(program)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(program)
}

func (this *QProcess) Start2() {
	purego.SyscallN(addr_QProcess_start2, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcess) StartCommand(command string) {
	var command_p uintptr
	if len(command) > 0 {
		command_p = uintptr(unsafe.Pointer(unsafe.StringData(command)))
	}
	purego.SyscallN(addr_QProcess_startCommand, this.h, command_p, uintptr(len(command)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(command)
}

func (this *QProcess) StartDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) Program() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_program, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProcess) SetProgram(program string) {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	purego.SyscallN(addr_QProcess_setProgram, this.h, program_p, uintptr(len(program)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(program)
}

func (this *QProcess) Arguments() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_arguments, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QProcess) SetArguments(arguments []string) {
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	purego.SyscallN(addr_QProcess_setArguments, this.h, arguments_p, uintptr(len(arguments)))
	runtime.KeepAlive(this)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
}

func (this *QProcess) ProcessChannelMode() QProcess__ProcessChannelMode {
	__r0, _, _ := purego.SyscallN(addr_QProcess_processChannelMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__ProcessChannelMode(__r)
}

func (this *QProcess) SetProcessChannelMode(mode QProcess__ProcessChannelMode) {
	purego.SyscallN(addr_QProcess_setProcessChannelMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QProcess) InputChannelMode() QProcess__InputChannelMode {
	__r0, _, _ := purego.SyscallN(addr_QProcess_inputChannelMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__InputChannelMode(__r)
}

func (this *QProcess) SetInputChannelMode(mode QProcess__InputChannelMode) {
	purego.SyscallN(addr_QProcess_setInputChannelMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QProcess) ReadChannel() QProcess__ProcessChannel {
	__r0, _, _ := purego.SyscallN(addr_QProcess_readChannel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__ProcessChannel(__r)
}

func (this *QProcess) SetReadChannel(channel QProcess__ProcessChannel) {
	purego.SyscallN(addr_QProcess_setReadChannel, this.h, uintptr(channel))
	runtime.KeepAlive(this)
}

func (this *QProcess) CloseReadChannel(channel QProcess__ProcessChannel) {
	purego.SyscallN(addr_QProcess_closeReadChannel, this.h, uintptr(channel))
	runtime.KeepAlive(this)
}

func (this *QProcess) CloseWriteChannel() {
	purego.SyscallN(addr_QProcess_closeWriteChannel, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcess) SetStandardInputFile(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QProcess_setStandardInputFile, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QProcess) SetStandardOutputFile(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QProcess_setStandardOutputFile, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QProcess) SetStandardErrorFile(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QProcess_setStandardErrorFile, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QProcess) SetStandardOutputProcess(destination *QProcess) {
	purego.SyscallN(addr_QProcess_setStandardOutputProcess, this.h, uintptr(destination.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(destination)
}

func (this *QProcess) NativeArguments() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_nativeArguments, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProcess) SetNativeArguments(arguments string) {
	var arguments_p uintptr
	if len(arguments) > 0 {
		arguments_p = uintptr(unsafe.Pointer(unsafe.StringData(arguments)))
	}
	purego.SyscallN(addr_QProcess_setNativeArguments, this.h, arguments_p, uintptr(len(arguments)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arguments)
}

func (this *QProcess) WorkingDirectory() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_workingDirectory, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QProcess) SetWorkingDirectory(dir string) {
	var dir_p uintptr
	if len(dir) > 0 {
		dir_p = uintptr(unsafe.Pointer(unsafe.StringData(dir)))
	}
	purego.SyscallN(addr_QProcess_setWorkingDirectory, this.h, dir_p, uintptr(len(dir)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dir)
}

func (this *QProcess) SetEnvironment(environment []string) {
	environment_arr := make([]uintptr, len(environment))
	for __i, __e := range environment {
		environment_arr[__i] = pgCString(__e)
	}
	var environment_p uintptr
	if len(environment_arr) > 0 {
		environment_p = uintptr(unsafe.Pointer(&environment_arr[0]))
	}
	purego.SyscallN(addr_QProcess_setEnvironment, this.h, environment_p, uintptr(len(environment)))
	runtime.KeepAlive(this)
	for __i := range environment_arr {
		pgFree(environment_arr[__i])
	}
	runtime.KeepAlive(environment_arr)
}

func (this *QProcess) Environment() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_environment, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QProcess) SetProcessEnvironment(environment *QProcessEnvironment) {
	purego.SyscallN(addr_QProcess_setProcessEnvironment, this.h, uintptr(environment.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(environment)
}

func (this *QProcess) ProcessEnvironment() *QProcessEnvironment {
	__r, _, _ := purego.SyscallN(addr_QProcess_processEnvironment, this.h)
	runtime.KeepAlive(this)
	return newownedQProcessEnvironment(__r)
}

func (this *QProcess) Error() QProcess__ProcessError {
	__r0, _, _ := purego.SyscallN(addr_QProcess_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__ProcessError(__r)
}

func (this *QProcess) State() QProcess__ProcessState {
	__r0, _, _ := purego.SyscallN(addr_QProcess_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__ProcessState(__r)
}

func (this *QProcess) ProcessId() int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_processId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) WaitForStarted() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForStarted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) WaitForFinished() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForFinished, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) ReadAllStandardOutput() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QProcess_readAllStandardOutput, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QProcess) ReadAllStandardError() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QProcess_readAllStandardError, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QProcess) ExitCode() int {
	__r0, _, _ := purego.SyscallN(addr_QProcess_exitCode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProcess) ExitStatus() QProcess__ExitStatus {
	__r0, _, _ := purego.SyscallN(addr_QProcess_exitStatus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QProcess__ExitStatus(__r)
}

func QProcess_Execute(program string) int {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcess_execute, program_p, uintptr(len(program)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	return int(__r)
}

func QProcess_StartDetached(program string) bool {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached2, program_p, uintptr(len(program)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	return __r != 0
}

func QProcess_SystemEnvironment() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_systemEnvironment, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QProcess_NullDevice() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_nullDevice, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QProcess_SplitCommand(command *QStringView) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_splitCommand, uintptr(command.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(command)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QProcess) Terminate() {
	purego.SyscallN(addr_QProcess_terminate, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcess) Kill() {
	purego.SyscallN(addr_QProcess_kill, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcess) Finished(exitCode int) {
	purego.SyscallN(addr_QProcess_finished, this.h, uintptr(exitCode))
	runtime.KeepAlive(this)
}

func (this *QProcess) ErrorOccurred(error_ QProcess__ProcessError) {
	purego.SyscallN(addr_QProcess_errorOccurred, this.h, uintptr(error_))
	runtime.KeepAlive(this)
}

func (this *QProcess) SetProcessState(state QProcess__ProcessState) {
	purego.SyscallN(addr_QProcess_setProcessState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QProcess) SetOpenMode(openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QProcess_setOpenMode, this.h, uintptr(openMode))
	runtime.KeepAlive(this)
}

func (this *QProcess) SetErrorString(errorString string) {
	var errorString_p uintptr
	if len(errorString) > 0 {
		errorString_p = uintptr(unsafe.Pointer(unsafe.StringData(errorString)))
	}
	purego.SyscallN(addr_QProcess_setErrorString, this.h, errorString_p, uintptr(len(errorString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(errorString)
}

func QProcess_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProcess_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProcess_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProcess_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProcess) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QProcess_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QProcess) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QProcess_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QProcess) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QProcess_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QProcess) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QProcess_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QProcess_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QProcess_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QProcess) Start3(program string, arguments []string) {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	purego.SyscallN(addr_QProcess_start3, this.h, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
}

func (this *QProcess) Start4(program string, arguments []string, mode QIODeviceBase__OpenModeFlag) {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	purego.SyscallN(addr_QProcess_start4, this.h, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
}

func (this *QProcess) StartWithMode(mode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QProcess_start5, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QProcess) StartCommand2(command string, mode QIODeviceBase__OpenModeFlag) {
	var command_p uintptr
	if len(command) > 0 {
		command_p = uintptr(unsafe.Pointer(unsafe.StringData(command)))
	}
	purego.SyscallN(addr_QProcess_startCommand2, this.h, command_p, uintptr(len(command)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(command)
}

func (this *QProcess) StartDetachedWithPid() (bool, int64) {
	var pid_out int64
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached3, this.h, uintptr(unsafe.Pointer(&pid_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, pid_out
}

func (this *QProcess) SetStandardOutputFile2(fileName string, mode QIODeviceBase__OpenModeFlag) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QProcess_setStandardOutputFile2, this.h, fileName_p, uintptr(len(fileName)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QProcess) SetStandardErrorFile2(fileName string, mode QIODeviceBase__OpenModeFlag) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QProcess_setStandardErrorFile2, this.h, fileName_p, uintptr(len(fileName)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QProcess) WaitForStartedWithMsecs(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForStarted2, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) WaitForFinishedWithMsecs(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForFinished2, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QProcess_Execute2(program string, arguments []string) int {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcess_execute2, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
	return int(__r)
}

func QProcess_StartDetached2(program string, arguments []string) bool {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached4, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
	return __r != 0
}

func QProcess_StartDetached3(program string, arguments []string, workingDirectory string) bool {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	var workingDirectory_p uintptr
	if len(workingDirectory) > 0 {
		workingDirectory_p = uintptr(unsafe.Pointer(unsafe.StringData(workingDirectory)))
	}
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached5, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)), workingDirectory_p, uintptr(len(workingDirectory)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
	runtime.KeepAlive(workingDirectory)
	return __r != 0
}

func QProcess_StartDetached4(program string, arguments []string, workingDirectory string) (bool, int64) {
	var program_p uintptr
	if len(program) > 0 {
		program_p = uintptr(unsafe.Pointer(unsafe.StringData(program)))
	}
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	var workingDirectory_p uintptr
	if len(workingDirectory) > 0 {
		workingDirectory_p = uintptr(unsafe.Pointer(unsafe.StringData(workingDirectory)))
	}
	var pid_out int64
	__r0, _, _ := purego.SyscallN(addr_QProcess_startDetached6, program_p, uintptr(len(program)), arguments_p, uintptr(len(arguments)), workingDirectory_p, uintptr(len(workingDirectory)), uintptr(unsafe.Pointer(&pid_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(program)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
	runtime.KeepAlive(workingDirectory)
	return __r != 0, pid_out
}

func (this *QProcess) Finished2(exitCode int, exitStatus QProcess__ExitStatus) {
	purego.SyscallN(addr_QProcess_finished2, this.h, uintptr(exitCode), uintptr(exitStatus))
	runtime.KeepAlive(this)
}

func QProcess_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QProcess_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QProcess_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QProcess_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QProcess) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QProcess_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QProcess) OpenDefault(mode QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_open_default, this.h, uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) WaitForReadyReadDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForReadyRead_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) WaitForBytesWrittenDefault(msecs int) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_waitForBytesWritten_default, this.h, uintptr(msecs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) BytesToWriteDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_bytesToWrite_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) IsSequentialDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_isSequential_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) CloseDefault() {
	purego.SyscallN(addr_QProcess_close_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QProcess) PosDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_pos_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) SizeDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_size_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) SeekDefault(pos int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_seek_default, this.h, uintptr(pos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) AtEndDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_atEnd_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) ResetDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_reset_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) BytesAvailableDefault() int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_bytesAvailable_default, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) CanReadLineDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_canReadLine_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess) SkipDataDefault(maxSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QProcess_skipData_default, this.h, uintptr(maxSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QProcess) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QProcess) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QProcess) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QProcess_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProcess) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QProcess_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProcess) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QProcess_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QProcess) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProcess_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProcess) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QProcess_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QProcess) OnStarted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_started(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnFinished(fn func(int)) int64 {
	id := registerCallback(func(exitCode int32) { fn(int(exitCode)) })
	pg_QProcess_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnFinished2(fn func(int, QProcess__ExitStatus)) int64 {
	id := registerCallback(func(exitCode int32, exitStatus int32) { fn(int(exitCode), QProcess__ExitStatus(exitStatus)) })
	pg_QProcess_connect_finished2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnErrorOccurred(fn func(QProcess__ProcessError)) int64 {
	id := registerCallback(func(error_ int32) { fn(QProcess__ProcessError(error_)) })
	pg_QProcess_connect_errorOccurred(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnStateChanged(fn func(QProcess__ProcessState)) int64 {
	id := registerCallback(func(state int32) { fn(QProcess__ProcessState(state)) })
	pg_QProcess_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnReadyReadStandardOutput(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_readyReadStandardOutput(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnReadyReadStandardError(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_readyReadStandardError(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnReadyRead(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_readyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnChannelReadyRead(fn func(int)) int64 {
	id := registerCallback(func(channel int32) { fn(int(channel)) })
	pg_QProcess_connect_channelReadyRead(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnBytesWritten(fn func(int64)) int64 {
	id := registerCallback(func(bytes int64) { fn(bytes) })
	pg_QProcess_connect_bytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnChannelBytesWritten(fn func(int, int64)) int64 {
	id := registerCallback(func(channel int32, bytes int64) { fn(int(channel), bytes) })
	pg_QProcess_connect_channelBytesWritten(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnAboutToClose(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_aboutToClose(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnReadChannelFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_readChannelFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QProcess_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QProcess_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QProcess_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QProcess) Disconnect(handle int64) {
	pg_QProcess__disconnect(handle)
	unregisterCallback(handle)
}

func QProcess_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -554057428128861975:
		return 1
	case 8356976193102722774:
		return 2
	case 7726833620054767889:
		return 3
	case 1024747515116322948:
		return 4
	case 8888098400087344862:
		return 5
	case -7006921981232612925:
		return 6
	case 8618582476283917417:
		return 7
	case 5614464919154503228:
		return 8
	case 3241140482992819217:
		return 9
	case -1487528724786208339:
		return 10
	case 1086335023529244512:
		return 11
	case 4814449630054007943:
		return 12
	case -1250682508954504761:
		return 13
	case 8178738247667652082:
		return 14
	case 1237236456312824991:
		return 15
	case 5217623514630858297:
		return 16
	case -5409927581812374498:
		return 17
	case -1185315367526857833:
		return 18
	case 3769868229123868474:
		return 19
	case 7379029982170503008:
		return 20
	case -568369772982215962:
		return 21
	}
	return -1
}

func QProcess_enable(h uintptr, idx int32) { pg_QProcess__set_override(h, idx) }

func (this *QProcess) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QProcess) OnOpen(fn func(QIODeviceBase__OpenModeFlag) bool) {
	registerVirt(this.UnsafeGoID(), -554057428128861975, func(mode int32) int32 { return b2i(fn(QIODeviceBase__OpenModeFlag(mode))) }) // slot "open"
}

func (this *QProcess) OnWaitForReadyRead(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 8356976193102722774, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForReadyRead"
}

func (this *QProcess) OnWaitForBytesWritten(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 7726833620054767889, func(msecs int32) int32 { return b2i(fn(int(msecs))) }) // slot "waitForBytesWritten"
}

func (this *QProcess) OnBytesToWrite(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 1024747515116322948, func() int64 { return fn() }) // slot "bytesToWrite"
}

func (this *QProcess) OnIsSequential(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 8888098400087344862, func() int32 { return b2i(fn()) }) // slot "isSequential"
}

func (this *QProcess) OnClose(fn func()) {
	registerVirt(this.UnsafeGoID(), -7006921981232612925, func() { fn() }) // slot "close"
}

func (this *QProcess) OnPos(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 8618582476283917417, func() int64 { return fn() }) // slot "pos"
}

func (this *QProcess) OnSize(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() int64 { return fn() }) // slot "size"
}

func (this *QProcess) OnSeek(fn func(int64) bool) {
	registerVirt(this.UnsafeGoID(), 3241140482992819217, func(pos int64) int32 { return b2i(fn(pos)) }) // slot "seek"
}

func (this *QProcess) OnAtEnd(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1487528724786208339, func() int32 { return b2i(fn()) }) // slot "atEnd"
}

func (this *QProcess) OnReset(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() int32 { return b2i(fn()) }) // slot "reset"
}

func (this *QProcess) OnBytesAvailable(fn func() int64) {
	registerVirt(this.UnsafeGoID(), 4814449630054007943, func() int64 { return fn() }) // slot "bytesAvailable"
}

func (this *QProcess) OnCanReadLine(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -1250682508954504761, func() int32 { return b2i(fn()) }) // slot "canReadLine"
}

func (this *QProcess) OnSkipData(fn func(int64) int64) {
	registerVirt(this.UnsafeGoID(), 8178738247667652082, func(maxSize int64) int64 { return fn(maxSize) }) // slot "skipData"
}

func (this *QProcess) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QProcess) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QProcess) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QProcess) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QProcess) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QProcess) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QProcess) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QProcess) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProcess_delete(this.h)
	}
	this.UnsafeRelease()
}
