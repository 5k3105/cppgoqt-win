package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFutureWatcherBase_base func(uintptr) uintptr
	addr_QFutureWatcherBase_tr uintptr
	addr_QFutureWatcherBase_progressValue uintptr
	addr_QFutureWatcherBase_progressMinimum uintptr
	addr_QFutureWatcherBase_progressMaximum uintptr
	addr_QFutureWatcherBase_progressText uintptr
	addr_QFutureWatcherBase_isStarted uintptr
	addr_QFutureWatcherBase_isFinished uintptr
	addr_QFutureWatcherBase_isRunning uintptr
	addr_QFutureWatcherBase_isCanceled uintptr
	addr_QFutureWatcherBase_isPaused uintptr
	addr_QFutureWatcherBase_isSuspending uintptr
	addr_QFutureWatcherBase_isSuspended uintptr
	addr_QFutureWatcherBase_waitForFinished uintptr
	addr_QFutureWatcherBase_setPendingResultsLimit uintptr
	addr_QFutureWatcherBase_started uintptr
	addr_QFutureWatcherBase_finished uintptr
	addr_QFutureWatcherBase_canceled uintptr
	addr_QFutureWatcherBase_paused uintptr
	addr_QFutureWatcherBase_suspending uintptr
	addr_QFutureWatcherBase_suspended uintptr
	addr_QFutureWatcherBase_resumed uintptr
	addr_QFutureWatcherBase_resultReadyAt uintptr
	addr_QFutureWatcherBase_resultsReadyAt uintptr
	addr_QFutureWatcherBase_progressRangeChanged uintptr
	addr_QFutureWatcherBase_progressValueChanged uintptr
	addr_QFutureWatcherBase_progressTextChanged uintptr
	addr_QFutureWatcherBase_cancel uintptr
	addr_QFutureWatcherBase_setSuspended uintptr
	addr_QFutureWatcherBase_suspend uintptr
	addr_QFutureWatcherBase_resume uintptr
	addr_QFutureWatcherBase_toggleSuspended uintptr
	addr_QFutureWatcherBase_setPaused uintptr
	addr_QFutureWatcherBase_pause uintptr
	addr_QFutureWatcherBase_togglePaused uintptr
	addr_QFutureWatcherBase_connect uintptr
	addr_QFutureWatcherBase_connect2 uintptr
	addr_QFutureWatcherBase_tr2 uintptr
	addr_QFutureWatcherBase_tr3 uintptr
	addr_QFutureWatcherBase_connect3 uintptr
	addr_QFutureWatcherBase_connect4 uintptr
	pg_QFutureWatcherBase_connect_started func(uintptr, int64)
	pg_QFutureWatcherBase_connect_finished func(uintptr, int64)
	pg_QFutureWatcherBase_connect_canceled func(uintptr, int64)
	pg_QFutureWatcherBase_connect_paused func(uintptr, int64)
	pg_QFutureWatcherBase_connect_suspending func(uintptr, int64)
	pg_QFutureWatcherBase_connect_suspended func(uintptr, int64)
	pg_QFutureWatcherBase_connect_resumed func(uintptr, int64)
	pg_QFutureWatcherBase_connect_resultReadyAt func(uintptr, int64)
	pg_QFutureWatcherBase_connect_resultsReadyAt func(uintptr, int64)
	pg_QFutureWatcherBase_connect_progressRangeChanged func(uintptr, int64)
	pg_QFutureWatcherBase_connect_progressValueChanged func(uintptr, int64)
	pg_QFutureWatcherBase_connect_progressTextChanged func(uintptr, int64)
	pg_QFutureWatcherBase_connect_destroyed func(uintptr, int64)
	pg_QFutureWatcherBase_connect_destroyed2 func(uintptr, int64)
	pg_QFutureWatcherBase_connect_objectNameChanged func(uintptr, int64)
	pg_QFutureWatcherBase__disconnect func(int64)
	pg_QFutureWatcherBase_delete func(uintptr)
)

func init() {
	reg(&pg_QFutureWatcherBase_base, "QFutureWatcherBase_base")
	symaddr(&addr_QFutureWatcherBase_tr, "QFutureWatcherBase_tr")
	symaddr(&addr_QFutureWatcherBase_progressValue, "QFutureWatcherBase_progressValue")
	symaddr(&addr_QFutureWatcherBase_progressMinimum, "QFutureWatcherBase_progressMinimum")
	symaddr(&addr_QFutureWatcherBase_progressMaximum, "QFutureWatcherBase_progressMaximum")
	symaddr(&addr_QFutureWatcherBase_progressText, "QFutureWatcherBase_progressText")
	symaddr(&addr_QFutureWatcherBase_isStarted, "QFutureWatcherBase_isStarted")
	symaddr(&addr_QFutureWatcherBase_isFinished, "QFutureWatcherBase_isFinished")
	symaddr(&addr_QFutureWatcherBase_isRunning, "QFutureWatcherBase_isRunning")
	symaddr(&addr_QFutureWatcherBase_isCanceled, "QFutureWatcherBase_isCanceled")
	symaddr(&addr_QFutureWatcherBase_isPaused, "QFutureWatcherBase_isPaused")
	symaddr(&addr_QFutureWatcherBase_isSuspending, "QFutureWatcherBase_isSuspending")
	symaddr(&addr_QFutureWatcherBase_isSuspended, "QFutureWatcherBase_isSuspended")
	symaddr(&addr_QFutureWatcherBase_waitForFinished, "QFutureWatcherBase_waitForFinished")
	symaddr(&addr_QFutureWatcherBase_setPendingResultsLimit, "QFutureWatcherBase_setPendingResultsLimit")
	symaddr(&addr_QFutureWatcherBase_started, "QFutureWatcherBase_started")
	symaddr(&addr_QFutureWatcherBase_finished, "QFutureWatcherBase_finished")
	symaddr(&addr_QFutureWatcherBase_canceled, "QFutureWatcherBase_canceled")
	symaddr(&addr_QFutureWatcherBase_paused, "QFutureWatcherBase_paused")
	symaddr(&addr_QFutureWatcherBase_suspending, "QFutureWatcherBase_suspending")
	symaddr(&addr_QFutureWatcherBase_suspended, "QFutureWatcherBase_suspended")
	symaddr(&addr_QFutureWatcherBase_resumed, "QFutureWatcherBase_resumed")
	symaddr(&addr_QFutureWatcherBase_resultReadyAt, "QFutureWatcherBase_resultReadyAt")
	symaddr(&addr_QFutureWatcherBase_resultsReadyAt, "QFutureWatcherBase_resultsReadyAt")
	symaddr(&addr_QFutureWatcherBase_progressRangeChanged, "QFutureWatcherBase_progressRangeChanged")
	symaddr(&addr_QFutureWatcherBase_progressValueChanged, "QFutureWatcherBase_progressValueChanged")
	symaddr(&addr_QFutureWatcherBase_progressTextChanged, "QFutureWatcherBase_progressTextChanged")
	symaddr(&addr_QFutureWatcherBase_cancel, "QFutureWatcherBase_cancel")
	symaddr(&addr_QFutureWatcherBase_setSuspended, "QFutureWatcherBase_setSuspended")
	symaddr(&addr_QFutureWatcherBase_suspend, "QFutureWatcherBase_suspend")
	symaddr(&addr_QFutureWatcherBase_resume, "QFutureWatcherBase_resume")
	symaddr(&addr_QFutureWatcherBase_toggleSuspended, "QFutureWatcherBase_toggleSuspended")
	symaddr(&addr_QFutureWatcherBase_setPaused, "QFutureWatcherBase_setPaused")
	symaddr(&addr_QFutureWatcherBase_pause, "QFutureWatcherBase_pause")
	symaddr(&addr_QFutureWatcherBase_togglePaused, "QFutureWatcherBase_togglePaused")
	symaddr(&addr_QFutureWatcherBase_connect, "QFutureWatcherBase_connect")
	symaddr(&addr_QFutureWatcherBase_connect2, "QFutureWatcherBase_connect2")
	symaddr(&addr_QFutureWatcherBase_tr2, "QFutureWatcherBase_tr2")
	symaddr(&addr_QFutureWatcherBase_tr3, "QFutureWatcherBase_tr3")
	symaddr(&addr_QFutureWatcherBase_connect3, "QFutureWatcherBase_connect3")
	symaddr(&addr_QFutureWatcherBase_connect4, "QFutureWatcherBase_connect4")
	reg(&pg_QFutureWatcherBase_connect_started, "QFutureWatcherBase_connect_started")
	reg(&pg_QFutureWatcherBase_connect_finished, "QFutureWatcherBase_connect_finished")
	reg(&pg_QFutureWatcherBase_connect_canceled, "QFutureWatcherBase_connect_canceled")
	reg(&pg_QFutureWatcherBase_connect_paused, "QFutureWatcherBase_connect_paused")
	reg(&pg_QFutureWatcherBase_connect_suspending, "QFutureWatcherBase_connect_suspending")
	reg(&pg_QFutureWatcherBase_connect_suspended, "QFutureWatcherBase_connect_suspended")
	reg(&pg_QFutureWatcherBase_connect_resumed, "QFutureWatcherBase_connect_resumed")
	reg(&pg_QFutureWatcherBase_connect_resultReadyAt, "QFutureWatcherBase_connect_resultReadyAt")
	reg(&pg_QFutureWatcherBase_connect_resultsReadyAt, "QFutureWatcherBase_connect_resultsReadyAt")
	reg(&pg_QFutureWatcherBase_connect_progressRangeChanged, "QFutureWatcherBase_connect_progressRangeChanged")
	reg(&pg_QFutureWatcherBase_connect_progressValueChanged, "QFutureWatcherBase_connect_progressValueChanged")
	reg(&pg_QFutureWatcherBase_connect_progressTextChanged, "QFutureWatcherBase_connect_progressTextChanged")
	reg(&pg_QFutureWatcherBase_connect_destroyed, "QFutureWatcherBase_connect_destroyed")
	reg(&pg_QFutureWatcherBase_connect_destroyed2, "QFutureWatcherBase_connect_destroyed2")
	reg(&pg_QFutureWatcherBase_connect_objectNameChanged, "QFutureWatcherBase_connect_objectNameChanged")
	reg(&pg_QFutureWatcherBase__disconnect, "QFutureWatcherBase__disconnect")
	reg(&pg_QFutureWatcherBase_delete, "QFutureWatcherBase_delete")
}

type QFutureWatcherBase struct {
	h uintptr
	*QObject
}

func bindQFutureWatcherBase(h uintptr) *QFutureWatcherBase {
	if h == 0 {
		return nil
	}
	return &QFutureWatcherBase{h: h, QObject: bindQObject(pg_QFutureWatcherBase_base(h))}
}

func newownedQFutureWatcherBase(h uintptr) *QFutureWatcherBase { return bindQFutureWatcherBase(h) }

func (this *QFutureWatcherBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFutureWatcherBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFutureWatcherBase(h unsafe.Pointer) *QFutureWatcherBase { return bindQFutureWatcherBase(uintptr(h)) }
func UnsafeNewQFutureWatcherBaseOwned(h unsafe.Pointer) *QFutureWatcherBase { return newownedQFutureWatcherBase(uintptr(h)) }

func (this *QFutureWatcherBase) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QFutureWatcherBase_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFutureWatcherBase) ProgressValue() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_progressValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureWatcherBase) ProgressMinimum() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_progressMinimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureWatcherBase) ProgressMaximum() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_progressMaximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureWatcherBase) ProgressText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_progressText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFutureWatcherBase) IsStarted() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isStarted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsFinished() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isFinished, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsRunning() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isRunning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsCanceled() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isCanceled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsPaused() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isPaused, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsSuspending() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isSuspending, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) IsSuspended() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureWatcherBase_isSuspended, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureWatcherBase) WaitForFinished() {
	purego.SyscallN(addr_QFutureWatcherBase_waitForFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) SetPendingResultsLimit(limit int) {
	purego.SyscallN(addr_QFutureWatcherBase_setPendingResultsLimit, this.h, uintptr(limit))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Started() {
	purego.SyscallN(addr_QFutureWatcherBase_started, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Finished() {
	purego.SyscallN(addr_QFutureWatcherBase_finished, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Canceled() {
	purego.SyscallN(addr_QFutureWatcherBase_canceled, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Paused() {
	purego.SyscallN(addr_QFutureWatcherBase_paused, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Suspending() {
	purego.SyscallN(addr_QFutureWatcherBase_suspending, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Suspended() {
	purego.SyscallN(addr_QFutureWatcherBase_suspended, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Resumed() {
	purego.SyscallN(addr_QFutureWatcherBase_resumed, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ResultReadyAt(resultIndex int) {
	purego.SyscallN(addr_QFutureWatcherBase_resultReadyAt, this.h, uintptr(resultIndex))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ResultsReadyAt(beginIndex int, endIndex int) {
	purego.SyscallN(addr_QFutureWatcherBase_resultsReadyAt, this.h, uintptr(beginIndex), uintptr(endIndex))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ProgressRangeChanged(minimum int, maximum int) {
	purego.SyscallN(addr_QFutureWatcherBase_progressRangeChanged, this.h, uintptr(minimum), uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ProgressValueChanged(progressValue int) {
	purego.SyscallN(addr_QFutureWatcherBase_progressValueChanged, this.h, uintptr(progressValue))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ProgressTextChanged(progressText string) {
	var progressText_p uintptr
	if len(progressText) > 0 {
		progressText_p = uintptr(unsafe.Pointer(unsafe.StringData(progressText)))
	}
	purego.SyscallN(addr_QFutureWatcherBase_progressTextChanged, this.h, progressText_p, uintptr(len(progressText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(progressText)
}

func (this *QFutureWatcherBase) Cancel() {
	purego.SyscallN(addr_QFutureWatcherBase_cancel, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) SetSuspended(suspend bool) {
	purego.SyscallN(addr_QFutureWatcherBase_setSuspended, this.h, uintptr(b2i(suspend)))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Suspend() {
	purego.SyscallN(addr_QFutureWatcherBase_suspend, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Resume() {
	purego.SyscallN(addr_QFutureWatcherBase_resume, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) ToggleSuspended() {
	purego.SyscallN(addr_QFutureWatcherBase_toggleSuspended, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) SetPaused(paused bool) {
	purego.SyscallN(addr_QFutureWatcherBase_setPaused, this.h, uintptr(b2i(paused)))
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) Pause() {
	purego.SyscallN(addr_QFutureWatcherBase_pause, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureWatcherBase) TogglePaused() {
	purego.SyscallN(addr_QFutureWatcherBase_togglePaused, this.h)
	runtime.KeepAlive(this)
}

func QFutureWatcherBase_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFutureWatcherBase_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QFutureWatcherBase_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFutureWatcherBase_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFutureWatcherBase_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFutureWatcherBase_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFutureWatcherBase_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFutureWatcherBase) OnStarted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_started(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnCanceled(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_canceled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnPaused(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_paused(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnSuspending(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_suspending(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnSuspended(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_suspended(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnResumed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_resumed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnResultReadyAt(fn func(int)) int64 {
	id := registerCallback(func(resultIndex int32) { fn(int(resultIndex)) })
	pg_QFutureWatcherBase_connect_resultReadyAt(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnResultsReadyAt(fn func(int, int)) int64 {
	id := registerCallback(func(beginIndex int32, endIndex int32) { fn(int(beginIndex), int(endIndex)) })
	pg_QFutureWatcherBase_connect_resultsReadyAt(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnProgressRangeChanged(fn func(int, int)) int64 {
	id := registerCallback(func(minimum int32, maximum int32) { fn(int(minimum), int(maximum)) })
	pg_QFutureWatcherBase_connect_progressRangeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnProgressValueChanged(fn func(int)) int64 {
	id := registerCallback(func(progressValue int32) { fn(int(progressValue)) })
	pg_QFutureWatcherBase_connect_progressValueChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnProgressTextChanged(fn func(string)) int64 {
	id := registerCallback(func(progressText uintptr, progressText_len int64) { fn(pgStrNoFree(progressText, progressText_len)) })
	pg_QFutureWatcherBase_connect_progressTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFutureWatcherBase_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFutureWatcherBase_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFutureWatcherBase_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFutureWatcherBase) Disconnect(handle int64) {
	pg_QFutureWatcherBase__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QFutureWatcherBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFutureWatcherBase_delete(this.h)
	}
	this.UnsafeRelease()
}
