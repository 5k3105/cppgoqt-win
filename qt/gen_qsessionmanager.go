package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSessionManager_base func(uintptr) uintptr
	addr_QSessionManager_tr uintptr
	addr_QSessionManager_sessionId uintptr
	addr_QSessionManager_sessionKey uintptr
	addr_QSessionManager_allowsInteraction uintptr
	addr_QSessionManager_allowsErrorInteraction uintptr
	addr_QSessionManager_release uintptr
	addr_QSessionManager_cancel uintptr
	addr_QSessionManager_setRestartHint uintptr
	addr_QSessionManager_restartHint uintptr
	addr_QSessionManager_setRestartCommand uintptr
	addr_QSessionManager_restartCommand uintptr
	addr_QSessionManager_setDiscardCommand uintptr
	addr_QSessionManager_discardCommand uintptr
	addr_QSessionManager_setManagerProperty uintptr
	addr_QSessionManager_setManagerProperty2 uintptr
	addr_QSessionManager_isPhase2 uintptr
	addr_QSessionManager_requestPhase2 uintptr
	addr_QSessionManager_connect uintptr
	addr_QSessionManager_connect2 uintptr
	addr_QSessionManager_tr2 uintptr
	addr_QSessionManager_tr3 uintptr
	addr_QSessionManager_connect3 uintptr
	addr_QSessionManager_connect4 uintptr
	pg_QSessionManager_connect_destroyed func(uintptr, int64)
	pg_QSessionManager_connect_destroyed2 func(uintptr, int64)
	pg_QSessionManager_connect_objectNameChanged func(uintptr, int64)
	pg_QSessionManager__disconnect func(int64)
)

func init() {
	reg(&pg_QSessionManager_base, "QSessionManager_base")
	symaddr(&addr_QSessionManager_tr, "QSessionManager_tr")
	symaddr(&addr_QSessionManager_sessionId, "QSessionManager_sessionId")
	symaddr(&addr_QSessionManager_sessionKey, "QSessionManager_sessionKey")
	symaddr(&addr_QSessionManager_allowsInteraction, "QSessionManager_allowsInteraction")
	symaddr(&addr_QSessionManager_allowsErrorInteraction, "QSessionManager_allowsErrorInteraction")
	symaddr(&addr_QSessionManager_release, "QSessionManager_release")
	symaddr(&addr_QSessionManager_cancel, "QSessionManager_cancel")
	symaddr(&addr_QSessionManager_setRestartHint, "QSessionManager_setRestartHint")
	symaddr(&addr_QSessionManager_restartHint, "QSessionManager_restartHint")
	symaddr(&addr_QSessionManager_setRestartCommand, "QSessionManager_setRestartCommand")
	symaddr(&addr_QSessionManager_restartCommand, "QSessionManager_restartCommand")
	symaddr(&addr_QSessionManager_setDiscardCommand, "QSessionManager_setDiscardCommand")
	symaddr(&addr_QSessionManager_discardCommand, "QSessionManager_discardCommand")
	symaddr(&addr_QSessionManager_setManagerProperty, "QSessionManager_setManagerProperty")
	symaddr(&addr_QSessionManager_setManagerProperty2, "QSessionManager_setManagerProperty2")
	symaddr(&addr_QSessionManager_isPhase2, "QSessionManager_isPhase2")
	symaddr(&addr_QSessionManager_requestPhase2, "QSessionManager_requestPhase2")
	symaddr(&addr_QSessionManager_connect, "QSessionManager_connect")
	symaddr(&addr_QSessionManager_connect2, "QSessionManager_connect2")
	symaddr(&addr_QSessionManager_tr2, "QSessionManager_tr2")
	symaddr(&addr_QSessionManager_tr3, "QSessionManager_tr3")
	symaddr(&addr_QSessionManager_connect3, "QSessionManager_connect3")
	symaddr(&addr_QSessionManager_connect4, "QSessionManager_connect4")
	reg(&pg_QSessionManager_connect_destroyed, "QSessionManager_connect_destroyed")
	reg(&pg_QSessionManager_connect_destroyed2, "QSessionManager_connect_destroyed2")
	reg(&pg_QSessionManager_connect_objectNameChanged, "QSessionManager_connect_objectNameChanged")
	reg(&pg_QSessionManager__disconnect, "QSessionManager__disconnect")
}

type QSessionManager struct {
	h uintptr
	*QObject
}

func bindQSessionManager(h uintptr) *QSessionManager {
	if h == 0 {
		return nil
	}
	return &QSessionManager{h: h, QObject: bindQObject(pg_QSessionManager_base(h))}
}

func newownedQSessionManager(h uintptr) *QSessionManager { return bindQSessionManager(h) }

func (this *QSessionManager) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSessionManager) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSessionManager(h unsafe.Pointer) *QSessionManager { return bindQSessionManager(uintptr(h)) }
func UnsafeNewQSessionManagerOwned(h unsafe.Pointer) *QSessionManager { return newownedQSessionManager(uintptr(h)) }

func (this *QSessionManager) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QSessionManager_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSessionManager) SessionId() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_sessionId, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSessionManager) SessionKey() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_sessionKey, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSessionManager) AllowsInteraction() bool {
	__r0, _, _ := purego.SyscallN(addr_QSessionManager_allowsInteraction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSessionManager) AllowsErrorInteraction() bool {
	__r0, _, _ := purego.SyscallN(addr_QSessionManager_allowsErrorInteraction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSessionManager) Release() {
	purego.SyscallN(addr_QSessionManager_release, this.h)
	runtime.KeepAlive(this)
}

func (this *QSessionManager) Cancel() {
	purego.SyscallN(addr_QSessionManager_cancel, this.h)
	runtime.KeepAlive(this)
}

func (this *QSessionManager) SetRestartHint(arg1 QSessionManager__RestartHint) {
	purego.SyscallN(addr_QSessionManager_setRestartHint, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSessionManager) RestartHint() QSessionManager__RestartHint {
	__r0, _, _ := purego.SyscallN(addr_QSessionManager_restartHint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSessionManager__RestartHint(__r)
}

func (this *QSessionManager) SetRestartCommand(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QSessionManager_setRestartCommand, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func (this *QSessionManager) RestartCommand() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_restartCommand, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QSessionManager) SetDiscardCommand(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QSessionManager_setDiscardCommand, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func (this *QSessionManager) DiscardCommand() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_discardCommand, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QSessionManager) SetManagerProperty(name string, value string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QSessionManager_setManagerProperty, this.h, name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QSessionManager) SetManagerProperty2(name string, value []string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	value_arr := make([]uintptr, len(value))
	for __i, __e := range value {
		value_arr[__i] = pgCString(__e)
	}
	var value_p uintptr
	if len(value_arr) > 0 {
		value_p = uintptr(unsafe.Pointer(&value_arr[0]))
	}
	purego.SyscallN(addr_QSessionManager_setManagerProperty2, this.h, name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	for __i := range value_arr {
		pgFree(value_arr[__i])
	}
	runtime.KeepAlive(value_arr)
}

func (this *QSessionManager) IsPhase2() bool {
	__r0, _, _ := purego.SyscallN(addr_QSessionManager_isPhase2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSessionManager) RequestPhase2() {
	purego.SyscallN(addr_QSessionManager_requestPhase2, this.h)
	runtime.KeepAlive(this)
}

func QSessionManager_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSessionManager_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSessionManager_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSessionManager_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QSessionManager_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSessionManager_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSessionManager_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSessionManager_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSessionManager_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSessionManager_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSessionManager_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSessionManager) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSessionManager_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSessionManager) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSessionManager_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSessionManager) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSessionManager_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSessionManager) Disconnect(handle int64) {
	pg_QSessionManager__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QSessionManager) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
