package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFutureInterfaceBase_new uintptr
	addr_QFutureInterfaceBase_new2 uintptr
	addr_QFutureInterfaceBase_new3 uintptr
	addr_QFutureInterfaceBase_OperatorAssign uintptr
	addr_QFutureInterfaceBase_reportStarted uintptr
	addr_QFutureInterfaceBase_reportFinished uintptr
	addr_QFutureInterfaceBase_reportCanceled uintptr
	addr_QFutureInterfaceBase_reportException uintptr
	addr_QFutureInterfaceBase_reportResultsReady uintptr
	addr_QFutureInterfaceBase_setRunnable uintptr
	addr_QFutureInterfaceBase_setThreadPool uintptr
	addr_QFutureInterfaceBase_threadPool uintptr
	addr_QFutureInterfaceBase_setFilterMode uintptr
	addr_QFutureInterfaceBase_setProgressRange uintptr
	addr_QFutureInterfaceBase_progressMinimum uintptr
	addr_QFutureInterfaceBase_progressMaximum uintptr
	addr_QFutureInterfaceBase_isProgressUpdateNeeded uintptr
	addr_QFutureInterfaceBase_setProgressValue uintptr
	addr_QFutureInterfaceBase_progressValue uintptr
	addr_QFutureInterfaceBase_setProgressValueAndText uintptr
	addr_QFutureInterfaceBase_progressText uintptr
	addr_QFutureInterfaceBase_setExpectedResultCount uintptr
	addr_QFutureInterfaceBase_expectedResultCount uintptr
	addr_QFutureInterfaceBase_resultCount uintptr
	addr_QFutureInterfaceBase_queryState uintptr
	addr_QFutureInterfaceBase_isRunning uintptr
	addr_QFutureInterfaceBase_isStarted uintptr
	addr_QFutureInterfaceBase_isCanceled uintptr
	addr_QFutureInterfaceBase_isFinished uintptr
	addr_QFutureInterfaceBase_isPaused uintptr
	addr_QFutureInterfaceBase_setPaused uintptr
	addr_QFutureInterfaceBase_togglePaused uintptr
	addr_QFutureInterfaceBase_isSuspending uintptr
	addr_QFutureInterfaceBase_isSuspended uintptr
	addr_QFutureInterfaceBase_isThrottled uintptr
	addr_QFutureInterfaceBase_isResultReadyAt uintptr
	addr_QFutureInterfaceBase_isValid uintptr
	addr_QFutureInterfaceBase_loadState uintptr
	addr_QFutureInterfaceBase_cancel uintptr
	addr_QFutureInterfaceBase_cancelAndFinish uintptr
	addr_QFutureInterfaceBase_cancelChain uintptr
	addr_QFutureInterfaceBase_setAddResultsIfCanceledEnabled uintptr
	addr_QFutureInterfaceBase_isAddResultsIfCanceledEnabled uintptr
	addr_QFutureInterfaceBase_setSuspended uintptr
	addr_QFutureInterfaceBase_toggleSuspended uintptr
	addr_QFutureInterfaceBase_reportSuspended uintptr
	addr_QFutureInterfaceBase_setThrottled uintptr
	addr_QFutureInterfaceBase_waitForFinished uintptr
	addr_QFutureInterfaceBase_waitForNextResult uintptr
	addr_QFutureInterfaceBase_waitForResult uintptr
	addr_QFutureInterfaceBase_waitForResume uintptr
	addr_QFutureInterfaceBase_suspendIfRequested uintptr
	addr_QFutureInterfaceBase_mutex uintptr
	addr_QFutureInterfaceBase_hasException uintptr
	addr_QFutureInterfaceBase_OperatorEqual uintptr
	addr_QFutureInterfaceBase_OperatorNotEqual uintptr
	addr_QFutureInterfaceBase_swap uintptr
	addr_QFutureInterfaceBase_isChainCanceled uintptr
	pg_QFutureInterfaceBase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFutureInterfaceBase_new, "QFutureInterfaceBase_new")
	symaddr(&addr_QFutureInterfaceBase_new2, "QFutureInterfaceBase_new2")
	symaddr(&addr_QFutureInterfaceBase_new3, "QFutureInterfaceBase_new3")
	symaddr(&addr_QFutureInterfaceBase_OperatorAssign, "QFutureInterfaceBase_OperatorAssign")
	symaddr(&addr_QFutureInterfaceBase_reportStarted, "QFutureInterfaceBase_reportStarted")
	symaddr(&addr_QFutureInterfaceBase_reportFinished, "QFutureInterfaceBase_reportFinished")
	symaddr(&addr_QFutureInterfaceBase_reportCanceled, "QFutureInterfaceBase_reportCanceled")
	symaddr(&addr_QFutureInterfaceBase_reportException, "QFutureInterfaceBase_reportException")
	symaddr(&addr_QFutureInterfaceBase_reportResultsReady, "QFutureInterfaceBase_reportResultsReady")
	symaddr(&addr_QFutureInterfaceBase_setRunnable, "QFutureInterfaceBase_setRunnable")
	symaddr(&addr_QFutureInterfaceBase_setThreadPool, "QFutureInterfaceBase_setThreadPool")
	symaddr(&addr_QFutureInterfaceBase_threadPool, "QFutureInterfaceBase_threadPool")
	symaddr(&addr_QFutureInterfaceBase_setFilterMode, "QFutureInterfaceBase_setFilterMode")
	symaddr(&addr_QFutureInterfaceBase_setProgressRange, "QFutureInterfaceBase_setProgressRange")
	symaddr(&addr_QFutureInterfaceBase_progressMinimum, "QFutureInterfaceBase_progressMinimum")
	symaddr(&addr_QFutureInterfaceBase_progressMaximum, "QFutureInterfaceBase_progressMaximum")
	symaddr(&addr_QFutureInterfaceBase_isProgressUpdateNeeded, "QFutureInterfaceBase_isProgressUpdateNeeded")
	symaddr(&addr_QFutureInterfaceBase_setProgressValue, "QFutureInterfaceBase_setProgressValue")
	symaddr(&addr_QFutureInterfaceBase_progressValue, "QFutureInterfaceBase_progressValue")
	symaddr(&addr_QFutureInterfaceBase_setProgressValueAndText, "QFutureInterfaceBase_setProgressValueAndText")
	symaddr(&addr_QFutureInterfaceBase_progressText, "QFutureInterfaceBase_progressText")
	symaddr(&addr_QFutureInterfaceBase_setExpectedResultCount, "QFutureInterfaceBase_setExpectedResultCount")
	symaddr(&addr_QFutureInterfaceBase_expectedResultCount, "QFutureInterfaceBase_expectedResultCount")
	symaddr(&addr_QFutureInterfaceBase_resultCount, "QFutureInterfaceBase_resultCount")
	symaddr(&addr_QFutureInterfaceBase_queryState, "QFutureInterfaceBase_queryState")
	symaddr(&addr_QFutureInterfaceBase_isRunning, "QFutureInterfaceBase_isRunning")
	symaddr(&addr_QFutureInterfaceBase_isStarted, "QFutureInterfaceBase_isStarted")
	symaddr(&addr_QFutureInterfaceBase_isCanceled, "QFutureInterfaceBase_isCanceled")
	symaddr(&addr_QFutureInterfaceBase_isFinished, "QFutureInterfaceBase_isFinished")
	symaddr(&addr_QFutureInterfaceBase_isPaused, "QFutureInterfaceBase_isPaused")
	symaddr(&addr_QFutureInterfaceBase_setPaused, "QFutureInterfaceBase_setPaused")
	symaddr(&addr_QFutureInterfaceBase_togglePaused, "QFutureInterfaceBase_togglePaused")
	symaddr(&addr_QFutureInterfaceBase_isSuspending, "QFutureInterfaceBase_isSuspending")
	symaddr(&addr_QFutureInterfaceBase_isSuspended, "QFutureInterfaceBase_isSuspended")
	symaddr(&addr_QFutureInterfaceBase_isThrottled, "QFutureInterfaceBase_isThrottled")
	symaddr(&addr_QFutureInterfaceBase_isResultReadyAt, "QFutureInterfaceBase_isResultReadyAt")
	symaddr(&addr_QFutureInterfaceBase_isValid, "QFutureInterfaceBase_isValid")
	symaddr(&addr_QFutureInterfaceBase_loadState, "QFutureInterfaceBase_loadState")
	symaddr(&addr_QFutureInterfaceBase_cancel, "QFutureInterfaceBase_cancel")
	symaddr(&addr_QFutureInterfaceBase_cancelAndFinish, "QFutureInterfaceBase_cancelAndFinish")
	symaddr(&addr_QFutureInterfaceBase_cancelChain, "QFutureInterfaceBase_cancelChain")
	symaddr(&addr_QFutureInterfaceBase_setAddResultsIfCanceledEnabled, "QFutureInterfaceBase_setAddResultsIfCanceledEnabled")
	symaddr(&addr_QFutureInterfaceBase_isAddResultsIfCanceledEnabled, "QFutureInterfaceBase_isAddResultsIfCanceledEnabled")
	symaddr(&addr_QFutureInterfaceBase_setSuspended, "QFutureInterfaceBase_setSuspended")
	symaddr(&addr_QFutureInterfaceBase_toggleSuspended, "QFutureInterfaceBase_toggleSuspended")
	symaddr(&addr_QFutureInterfaceBase_reportSuspended, "QFutureInterfaceBase_reportSuspended")
	symaddr(&addr_QFutureInterfaceBase_setThrottled, "QFutureInterfaceBase_setThrottled")
	symaddr(&addr_QFutureInterfaceBase_waitForFinished, "QFutureInterfaceBase_waitForFinished")
	symaddr(&addr_QFutureInterfaceBase_waitForNextResult, "QFutureInterfaceBase_waitForNextResult")
	symaddr(&addr_QFutureInterfaceBase_waitForResult, "QFutureInterfaceBase_waitForResult")
	symaddr(&addr_QFutureInterfaceBase_waitForResume, "QFutureInterfaceBase_waitForResume")
	symaddr(&addr_QFutureInterfaceBase_suspendIfRequested, "QFutureInterfaceBase_suspendIfRequested")
	symaddr(&addr_QFutureInterfaceBase_mutex, "QFutureInterfaceBase_mutex")
	symaddr(&addr_QFutureInterfaceBase_hasException, "QFutureInterfaceBase_hasException")
	symaddr(&addr_QFutureInterfaceBase_OperatorEqual, "QFutureInterfaceBase_OperatorEqual")
	symaddr(&addr_QFutureInterfaceBase_OperatorNotEqual, "QFutureInterfaceBase_OperatorNotEqual")
	symaddr(&addr_QFutureInterfaceBase_swap, "QFutureInterfaceBase_swap")
	symaddr(&addr_QFutureInterfaceBase_isChainCanceled, "QFutureInterfaceBase_isChainCanceled")
	reg(&pg_QFutureInterfaceBase_delete, "QFutureInterfaceBase_delete")
}

type QFutureInterfaceBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFutureInterfaceBase) UnsafeGoID() int64     { return this.vid }
func (this *QFutureInterfaceBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFutureInterfaceBase(h uintptr) *QFutureInterfaceBase {
	if h == 0 {
		return nil
	}
	return &QFutureInterfaceBase{h: h}
}

func newownedQFutureInterfaceBase(h uintptr) *QFutureInterfaceBase {
	o := bindQFutureInterfaceBase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFutureInterfaceBase).Delete)
	}
	return o
}

func (this *QFutureInterfaceBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFutureInterfaceBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFutureInterfaceBase(h unsafe.Pointer) *QFutureInterfaceBase { return bindQFutureInterfaceBase(uintptr(h)) }
func UnsafeNewQFutureInterfaceBaseOwned(h unsafe.Pointer) *QFutureInterfaceBase { return newownedQFutureInterfaceBase(uintptr(h)) }

func (this *QFutureInterfaceBase) UnsafeRelease() { this.h = 0 }

func NewQFutureInterfaceBase() *QFutureInterfaceBase {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_new)
	__o := newownedQFutureInterfaceBase(__r)
	return __o
}

func NewQFutureInterfaceBase2(other *QFutureInterfaceBase) *QFutureInterfaceBase {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_new2, uintptr(other.UnsafePointer()))
	__o := newownedQFutureInterfaceBase(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQFutureInterfaceBase3(initialState QFutureInterfaceBase__State) *QFutureInterfaceBase {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_new3, uintptr(initialState))
	__o := newownedQFutureInterfaceBase(__r)
	return __o
}

func (this *QFutureInterfaceBase) OperatorAssign(other *QFutureInterfaceBase) *QFutureInterfaceBase {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQFutureInterfaceBase(__r)
}

func (this *QFutureInterfaceBase) ReportStarted() {
	purego.SyscallN(addr_QFutureInterfaceBase_reportStarted, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ReportFinished() {
	purego.SyscallN(addr_QFutureInterfaceBase_reportFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ReportCanceled() {
	purego.SyscallN(addr_QFutureInterfaceBase_reportCanceled, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ReportException(e *QException) {
	purego.SyscallN(addr_QFutureInterfaceBase_reportException, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFutureInterfaceBase) ReportResultsReady(beginIndex int, endIndex int) {
	purego.SyscallN(addr_QFutureInterfaceBase_reportResultsReady, this.h, uintptr(beginIndex), uintptr(endIndex))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) SetRunnable(runnable *QRunnable) {
	purego.SyscallN(addr_QFutureInterfaceBase_setRunnable, this.h, uintptr(runnable.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(runnable)
}

func (this *QFutureInterfaceBase) SetThreadPool(pool *QThreadPool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setThreadPool, this.h, uintptr(pool.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pool)
}

func (this *QFutureInterfaceBase) ThreadPool() *QThreadPool {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_threadPool, this.h)
	runtime.KeepAlive(this)
	return bindQThreadPool(__r)
}

func (this *QFutureInterfaceBase) SetFilterMode(enable bool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setFilterMode, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) SetProgressRange(minimum int, maximum int) {
	purego.SyscallN(addr_QFutureInterfaceBase_setProgressRange, this.h, uintptr(minimum), uintptr(maximum))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ProgressMinimum() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_progressMinimum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) ProgressMaximum() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_progressMaximum, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) IsProgressUpdateNeeded() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isProgressUpdateNeeded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) SetProgressValue(progressValue int) {
	purego.SyscallN(addr_QFutureInterfaceBase_setProgressValue, this.h, uintptr(progressValue))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ProgressValue() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_progressValue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) SetProgressValueAndText(progressValue int, progressText string) {
	var progressText_p uintptr
	if len(progressText) > 0 {
		progressText_p = uintptr(unsafe.Pointer(unsafe.StringData(progressText)))
	}
	purego.SyscallN(addr_QFutureInterfaceBase_setProgressValueAndText, this.h, uintptr(progressValue), progressText_p, uintptr(len(progressText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(progressText)
}

func (this *QFutureInterfaceBase) ProgressText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_progressText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFutureInterfaceBase) SetExpectedResultCount(resultCount int) {
	purego.SyscallN(addr_QFutureInterfaceBase_setExpectedResultCount, this.h, uintptr(resultCount))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ExpectedResultCount() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_expectedResultCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) ResultCount() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_resultCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) QueryState(state QFutureInterfaceBase__State) bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_queryState, this.h, uintptr(state))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsRunning() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isRunning, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsStarted() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isStarted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsCanceled() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isCanceled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsFinished() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isFinished, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsPaused() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isPaused, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) SetPaused(paused bool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setPaused, this.h, uintptr(b2i(paused)))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) TogglePaused() {
	purego.SyscallN(addr_QFutureInterfaceBase_togglePaused, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) IsSuspending() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isSuspending, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsSuspended() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isSuspended, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsThrottled() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isThrottled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsResultReadyAt(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isResultReadyAt, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) LoadState() int {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_loadState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFutureInterfaceBase) Cancel() {
	purego.SyscallN(addr_QFutureInterfaceBase_cancel, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) CancelAndFinish() {
	purego.SyscallN(addr_QFutureInterfaceBase_cancelAndFinish, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) CancelChain() {
	purego.SyscallN(addr_QFutureInterfaceBase_cancelChain, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) SetAddResultsIfCanceledEnabled(enable bool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setAddResultsIfCanceledEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) IsAddResultsIfCanceledEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isAddResultsIfCanceledEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) SetSuspended(suspend bool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setSuspended, this.h, uintptr(b2i(suspend)))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ToggleSuspended() {
	purego.SyscallN(addr_QFutureInterfaceBase_toggleSuspended, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) ReportSuspended() {
	purego.SyscallN(addr_QFutureInterfaceBase_reportSuspended, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) SetThrottled(enable bool) {
	purego.SyscallN(addr_QFutureInterfaceBase_setThrottled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) WaitForFinished() {
	purego.SyscallN(addr_QFutureInterfaceBase_waitForFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) WaitForNextResult() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_waitForNextResult, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) WaitForResult(resultIndex int) {
	purego.SyscallN(addr_QFutureInterfaceBase_waitForResult, this.h, uintptr(resultIndex))
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) WaitForResume() {
	purego.SyscallN(addr_QFutureInterfaceBase_waitForResume, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) SuspendIfRequested() {
	purego.SyscallN(addr_QFutureInterfaceBase_suspendIfRequested, this.h)
	runtime.KeepAlive(this)
}

func (this *QFutureInterfaceBase) Mutex() *QMutex {
	__r, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_mutex, this.h)
	runtime.KeepAlive(this)
	return bindQMutex(__r)
}

func (this *QFutureInterfaceBase) HasException() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_hasException, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) OperatorEqual(other *QFutureInterfaceBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFutureInterfaceBase) OperatorNotEqual(other *QFutureInterfaceBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QFutureInterfaceBase) Swap(other *QFutureInterfaceBase) {
	purego.SyscallN(addr_QFutureInterfaceBase_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QFutureInterfaceBase) IsChainCanceled() bool {
	__r0, _, _ := purego.SyscallN(addr_QFutureInterfaceBase_isChainCanceled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFutureInterfaceBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFutureInterfaceBase_delete(this.h)
	}
	this.UnsafeRelease()
}
