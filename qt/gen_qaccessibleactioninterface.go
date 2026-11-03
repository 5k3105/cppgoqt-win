package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleActionInterface_tr uintptr
	addr_QAccessibleActionInterface_actionNames uintptr
	addr_QAccessibleActionInterface_localizedActionName uintptr
	addr_QAccessibleActionInterface_localizedActionDescription uintptr
	addr_QAccessibleActionInterface_doAction uintptr
	addr_QAccessibleActionInterface_keyBindingsForAction uintptr
	addr_QAccessibleActionInterface_pressAction uintptr
	addr_QAccessibleActionInterface_increaseAction uintptr
	addr_QAccessibleActionInterface_decreaseAction uintptr
	addr_QAccessibleActionInterface_showMenuAction uintptr
	addr_QAccessibleActionInterface_setFocusAction uintptr
	addr_QAccessibleActionInterface_toggleAction uintptr
	addr_QAccessibleActionInterface_scrollLeftAction uintptr
	addr_QAccessibleActionInterface_scrollRightAction uintptr
	addr_QAccessibleActionInterface_scrollUpAction uintptr
	addr_QAccessibleActionInterface_scrollDownAction uintptr
	addr_QAccessibleActionInterface_nextPageAction uintptr
	addr_QAccessibleActionInterface_previousPageAction uintptr
	addr_QAccessibleActionInterface_OperatorAssign uintptr
	addr_QAccessibleActionInterface_tr2 uintptr
	addr_QAccessibleActionInterface_tr3 uintptr
	addr_QAccessibleActionInterface_localizedActionName_default uintptr
	addr_QAccessibleActionInterface_localizedActionDescription_default uintptr
	pg_QAccessibleActionInterface__set_override func(uintptr, int32)
	pg_QAccessibleActionInterface__Go_new func(int64) uintptr
	pg_QAccessibleActionInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleActionInterface_tr, "QAccessibleActionInterface_tr")
	symaddr(&addr_QAccessibleActionInterface_actionNames, "QAccessibleActionInterface_actionNames")
	symaddr(&addr_QAccessibleActionInterface_localizedActionName, "QAccessibleActionInterface_localizedActionName")
	symaddr(&addr_QAccessibleActionInterface_localizedActionDescription, "QAccessibleActionInterface_localizedActionDescription")
	symaddr(&addr_QAccessibleActionInterface_doAction, "QAccessibleActionInterface_doAction")
	symaddr(&addr_QAccessibleActionInterface_keyBindingsForAction, "QAccessibleActionInterface_keyBindingsForAction")
	symaddr(&addr_QAccessibleActionInterface_pressAction, "QAccessibleActionInterface_pressAction")
	symaddr(&addr_QAccessibleActionInterface_increaseAction, "QAccessibleActionInterface_increaseAction")
	symaddr(&addr_QAccessibleActionInterface_decreaseAction, "QAccessibleActionInterface_decreaseAction")
	symaddr(&addr_QAccessibleActionInterface_showMenuAction, "QAccessibleActionInterface_showMenuAction")
	symaddr(&addr_QAccessibleActionInterface_setFocusAction, "QAccessibleActionInterface_setFocusAction")
	symaddr(&addr_QAccessibleActionInterface_toggleAction, "QAccessibleActionInterface_toggleAction")
	symaddr(&addr_QAccessibleActionInterface_scrollLeftAction, "QAccessibleActionInterface_scrollLeftAction")
	symaddr(&addr_QAccessibleActionInterface_scrollRightAction, "QAccessibleActionInterface_scrollRightAction")
	symaddr(&addr_QAccessibleActionInterface_scrollUpAction, "QAccessibleActionInterface_scrollUpAction")
	symaddr(&addr_QAccessibleActionInterface_scrollDownAction, "QAccessibleActionInterface_scrollDownAction")
	symaddr(&addr_QAccessibleActionInterface_nextPageAction, "QAccessibleActionInterface_nextPageAction")
	symaddr(&addr_QAccessibleActionInterface_previousPageAction, "QAccessibleActionInterface_previousPageAction")
	symaddr(&addr_QAccessibleActionInterface_OperatorAssign, "QAccessibleActionInterface_OperatorAssign")
	symaddr(&addr_QAccessibleActionInterface_tr2, "QAccessibleActionInterface_tr2")
	symaddr(&addr_QAccessibleActionInterface_tr3, "QAccessibleActionInterface_tr3")
	symaddr(&addr_QAccessibleActionInterface_localizedActionName_default, "QAccessibleActionInterface_localizedActionName_default")
	symaddr(&addr_QAccessibleActionInterface_localizedActionDescription_default, "QAccessibleActionInterface_localizedActionDescription_default")
	reg(&pg_QAccessibleActionInterface__set_override, "QAccessibleActionInterface__set_override")
	reg(&pg_QAccessibleActionInterface__Go_new, "QAccessibleActionInterface__Go_new")
	reg(&pg_QAccessibleActionInterface_delete, "QAccessibleActionInterface_delete")
}

type QAccessibleActionInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleActionInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleActionInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleActionInterface(h uintptr) *QAccessibleActionInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleActionInterface{h: h}
}

func newownedQAccessibleActionInterface(h uintptr) *QAccessibleActionInterface {
	o := bindQAccessibleActionInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleActionInterface).Delete)
	}
	return o
}

func (this *QAccessibleActionInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleActionInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleActionInterface(h unsafe.Pointer) *QAccessibleActionInterface { return bindQAccessibleActionInterface(uintptr(h)) }
func UnsafeNewQAccessibleActionInterfaceOwned(h unsafe.Pointer) *QAccessibleActionInterface { return newownedQAccessibleActionInterface(uintptr(h)) }

func (this *QAccessibleActionInterface) UnsafeRelease() { this.h = 0 }

func QAccessibleActionInterface_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) ActionNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_actionNames, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAccessibleActionInterface) LocalizedActionName(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_localizedActionName, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) LocalizedActionDescription(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_localizedActionDescription, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) DoAction(actionName string) {
	var actionName_p uintptr
	if len(actionName) > 0 {
		actionName_p = uintptr(unsafe.Pointer(unsafe.StringData(actionName)))
	}
	purego.SyscallN(addr_QAccessibleActionInterface_doAction, this.h, actionName_p, uintptr(len(actionName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(actionName)
}

func (this *QAccessibleActionInterface) KeyBindingsForAction(actionName string) []string {
	var actionName_p uintptr
	if len(actionName) > 0 {
		actionName_p = uintptr(unsafe.Pointer(unsafe.StringData(actionName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_keyBindingsForAction, this.h, actionName_p, uintptr(len(actionName)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(actionName)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QAccessibleActionInterface_PressAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_pressAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_IncreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_increaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_DecreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_decreaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ShowMenuAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_showMenuAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_SetFocusAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_setFocusAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ToggleAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_toggleAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ScrollLeftAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_scrollLeftAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ScrollRightAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_scrollRightAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ScrollUpAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_scrollUpAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_ScrollDownAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_scrollDownAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_NextPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_nextPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_PreviousPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_previousPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) OperatorAssign(arg1 *QAccessibleActionInterface) *QAccessibleActionInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleActionInterface(__r)
}

func QAccessibleActionInterface_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) LocalizedActionNameDefault(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_localizedActionName_default, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QAccessibleActionInterface) LocalizedActionDescriptionDefault(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleActionInterface_localizedActionDescription_default, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func QAccessibleActionInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 5274920687135865149:
		return 0
	case -219369122549130410:
		return 1
	case -9040881633578371181:
		return 2
	case 5320856891804608331:
		return 3
	case -2998875658810589822:
		return 4
	}
	return -1
}

func QAccessibleActionInterface_enable(h uintptr, idx int32) { pg_QAccessibleActionInterface__set_override(h, idx) }

func NewQAccessibleActionInterface() *QAccessibleActionInterface {
	id := nextVirtID()
	o := newownedQAccessibleActionInterface(pg_QAccessibleActionInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleActionInterface_virtIdx, QAccessibleActionInterface_enable)
	return o
}

func (this *QAccessibleActionInterface) OnActionNames(fn func() []string) {
	registerVirt(this.UnsafeGoID(), 5274920687135865149, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "actionNames"
}

func (this *QAccessibleActionInterface) OnDoAction(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -219369122549130410, func(actionName uintptr, actionName_len int64) { fn(pgStrNoFree(actionName, actionName_len)) }) // slot "doAction"
}

func (this *QAccessibleActionInterface) OnKeyBindingsForAction(fn func(string) []string) {
	registerVirt(this.UnsafeGoID(), -9040881633578371181, func(actionName uintptr, actionName_len int64) (uintptr, int64) {
		__rs := fn(pgStrNoFree(actionName, actionName_len))
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "keyBindingsForAction"
}

func (this *QAccessibleActionInterface) OnLocalizedActionName(fn func(string) string) {
	registerVirt(this.UnsafeGoID(), 5320856891804608331, func(name uintptr, name_len int64) (uintptr, int64) {
		__s := fn(pgStrNoFree(name, name_len))
		return pgCString(__s), int64(len(__s))
	}) // slot "localizedActionName"
}

func (this *QAccessibleActionInterface) OnLocalizedActionDescription(fn func(string) string) {
	registerVirt(this.UnsafeGoID(), -2998875658810589822, func(name uintptr, name_len int64) (uintptr, int64) {
		__s := fn(pgStrNoFree(name, name_len))
		return pgCString(__s), int64(len(__s))
	}) // slot "localizedActionDescription"
}

func (this *QAccessibleActionInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleActionInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
