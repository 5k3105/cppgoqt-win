package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleWidget_chain uintptr
	addr_QAccessibleWidget_new uintptr
	addr_QAccessibleWidget_new2 uintptr
	addr_QAccessibleWidget_new3 uintptr
	addr_QAccessibleWidget_actionNames uintptr
	addr_QAccessibleWidget_doAction uintptr
	addr_QAccessibleWidget_keyBindingsForAction uintptr
	addr_QAccessibleWidget_tr uintptr
	addr_QAccessibleWidget_localizedActionName uintptr
	addr_QAccessibleWidget_localizedActionDescription uintptr
	addr_QAccessibleWidget_pressAction uintptr
	addr_QAccessibleWidget_increaseAction uintptr
	addr_QAccessibleWidget_decreaseAction uintptr
	addr_QAccessibleWidget_showMenuAction uintptr
	addr_QAccessibleWidget_setFocusAction uintptr
	addr_QAccessibleWidget_toggleAction uintptr
	addr_QAccessibleWidget_scrollLeftAction uintptr
	addr_QAccessibleWidget_scrollRightAction uintptr
	addr_QAccessibleWidget_scrollUpAction uintptr
	addr_QAccessibleWidget_scrollDownAction uintptr
	addr_QAccessibleWidget_nextPageAction uintptr
	addr_QAccessibleWidget_previousPageAction uintptr
	addr_QAccessibleWidget_tr2 uintptr
	addr_QAccessibleWidget_tr3 uintptr
	addr_QAccessibleWidget_asQAccessibleActionInterface uintptr
)

func init() {
	symaddr(&addr_QAccessibleWidget_chain, "QAccessibleWidget_chain")
	symaddr(&addr_QAccessibleWidget_new, "QAccessibleWidget_new")
	symaddr(&addr_QAccessibleWidget_new2, "QAccessibleWidget_new2")
	symaddr(&addr_QAccessibleWidget_new3, "QAccessibleWidget_new3")
	symaddr(&addr_QAccessibleWidget_actionNames, "QAccessibleWidget_actionNames")
	symaddr(&addr_QAccessibleWidget_doAction, "QAccessibleWidget_doAction")
	symaddr(&addr_QAccessibleWidget_keyBindingsForAction, "QAccessibleWidget_keyBindingsForAction")
	symaddr(&addr_QAccessibleWidget_tr, "QAccessibleWidget_tr")
	symaddr(&addr_QAccessibleWidget_localizedActionName, "QAccessibleWidget_localizedActionName")
	symaddr(&addr_QAccessibleWidget_localizedActionDescription, "QAccessibleWidget_localizedActionDescription")
	symaddr(&addr_QAccessibleWidget_pressAction, "QAccessibleWidget_pressAction")
	symaddr(&addr_QAccessibleWidget_increaseAction, "QAccessibleWidget_increaseAction")
	symaddr(&addr_QAccessibleWidget_decreaseAction, "QAccessibleWidget_decreaseAction")
	symaddr(&addr_QAccessibleWidget_showMenuAction, "QAccessibleWidget_showMenuAction")
	symaddr(&addr_QAccessibleWidget_setFocusAction, "QAccessibleWidget_setFocusAction")
	symaddr(&addr_QAccessibleWidget_toggleAction, "QAccessibleWidget_toggleAction")
	symaddr(&addr_QAccessibleWidget_scrollLeftAction, "QAccessibleWidget_scrollLeftAction")
	symaddr(&addr_QAccessibleWidget_scrollRightAction, "QAccessibleWidget_scrollRightAction")
	symaddr(&addr_QAccessibleWidget_scrollUpAction, "QAccessibleWidget_scrollUpAction")
	symaddr(&addr_QAccessibleWidget_scrollDownAction, "QAccessibleWidget_scrollDownAction")
	symaddr(&addr_QAccessibleWidget_nextPageAction, "QAccessibleWidget_nextPageAction")
	symaddr(&addr_QAccessibleWidget_previousPageAction, "QAccessibleWidget_previousPageAction")
	symaddr(&addr_QAccessibleWidget_tr2, "QAccessibleWidget_tr2")
	symaddr(&addr_QAccessibleWidget_tr3, "QAccessibleWidget_tr3")
	symaddr(&addr_QAccessibleWidget_asQAccessibleActionInterface, "QAccessibleWidget_asQAccessibleActionInterface")
}

type QAccessibleWidget struct {
	h uintptr
	*QAccessibleObject
}

func bindQAccessibleWidget(h uintptr) *QAccessibleWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleWidget{h: h, QAccessibleObject: &QAccessibleObject{h: __c[0], QAccessibleInterface: &QAccessibleInterface{h: __c[1]}}}
}

func newownedQAccessibleWidget(h uintptr) *QAccessibleWidget { return bindQAccessibleWidget(h) }

func (this *QAccessibleWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleWidget(h unsafe.Pointer) *QAccessibleWidget { return bindQAccessibleWidget(uintptr(h)) }
func UnsafeNewQAccessibleWidgetOwned(h unsafe.Pointer) *QAccessibleWidget { return newownedQAccessibleWidget(uintptr(h)) }

func (this *QAccessibleWidget) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleObject != nil {
		this.QAccessibleObject.UnsafeRelease()
	}
}

func NewQAccessibleWidget(o *QWidget) *QAccessibleWidget {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_new, uintptr(o.UnsafePointer()))
	__o := newownedQAccessibleWidget(__r)
	runtime.KeepAlive(o)
	return __o
}

func NewQAccessibleWidget2(o *QWidget, r QAccessible__Role, name string) *QAccessibleWidget {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_new2, uintptr(o.UnsafePointer()), uintptr(r), name_p, uintptr(len(name)))
	__o := newownedQAccessibleWidget(__r)
	runtime.KeepAlive(o)
	runtime.KeepAlive(name)
	return __o
}

func NewQAccessibleWidget3(o *QWidget, r QAccessible__Role) *QAccessibleWidget {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_new3, uintptr(o.UnsafePointer()), uintptr(r))
	__o := newownedQAccessibleWidget(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QAccessibleWidget) ActionNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_actionNames, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAccessibleWidget) DoAction(actionName string) {
	var actionName_p uintptr
	if len(actionName) > 0 {
		actionName_p = uintptr(unsafe.Pointer(unsafe.StringData(actionName)))
	}
	purego.SyscallN(addr_QAccessibleWidget_doAction, this.h, actionName_p, uintptr(len(actionName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(actionName)
}

func (this *QAccessibleWidget) KeyBindingsForAction(actionName string) []string {
	var actionName_p uintptr
	if len(actionName) > 0 {
		actionName_p = uintptr(unsafe.Pointer(unsafe.StringData(actionName)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_keyBindingsForAction, this.h, actionName_p, uintptr(len(actionName)), uintptr(unsafe.Pointer(&__n)))
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

func QAccessibleWidget_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleWidget) LocalizedActionName(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_localizedActionName, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QAccessibleWidget) LocalizedActionDescription(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_localizedActionDescription, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func QAccessibleWidget_PressAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_pressAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_IncreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_increaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_DecreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_decreaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ShowMenuAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_showMenuAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_SetFocusAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_setFocusAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ToggleAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_toggleAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ScrollLeftAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_scrollLeftAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ScrollRightAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_scrollRightAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ScrollUpAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_scrollUpAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_ScrollDownAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_scrollDownAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_NextPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_nextPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_PreviousPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_previousPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidget_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QAccessibleWidget_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleWidget) AsQAccessibleActionInterface() *QAccessibleActionInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidget_asQAccessibleActionInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleActionInterface(__r)
}

func (this *QAccessibleWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
