package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleApplication_chain uintptr
	addr_QAccessibleApplication_new uintptr
	addr_QAccessibleApplication_window_default uintptr
	addr_QAccessibleApplication_childCount_default uintptr
	addr_QAccessibleApplication_indexOfChild_default uintptr
	addr_QAccessibleApplication_focusChild_default uintptr
	addr_QAccessibleApplication_parent_default uintptr
	addr_QAccessibleApplication_child_default uintptr
	addr_QAccessibleApplication_text_default uintptr
	addr_QAccessibleApplication_role_default uintptr
	addr_QAccessibleApplication_state_default uintptr
	addr_QAccessibleApplication_isValid_default uintptr
	addr_QAccessibleApplication_object_default uintptr
	addr_QAccessibleApplication_rect_default uintptr
	addr_QAccessibleApplication_setText_default uintptr
	addr_QAccessibleApplication_childAt_default uintptr
	addr_QAccessibleApplication_foregroundColor_default uintptr
	addr_QAccessibleApplication_backgroundColor_default uintptr
	pg_QAccessibleApplication__set_override func(uintptr, int32)
	pg_QAccessibleApplication_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleApplication_chain, "QAccessibleApplication_chain")
	symaddr(&addr_QAccessibleApplication_new, "QAccessibleApplication_new")
	symaddr(&addr_QAccessibleApplication_window_default, "QAccessibleApplication_window_default")
	symaddr(&addr_QAccessibleApplication_childCount_default, "QAccessibleApplication_childCount_default")
	symaddr(&addr_QAccessibleApplication_indexOfChild_default, "QAccessibleApplication_indexOfChild_default")
	symaddr(&addr_QAccessibleApplication_focusChild_default, "QAccessibleApplication_focusChild_default")
	symaddr(&addr_QAccessibleApplication_parent_default, "QAccessibleApplication_parent_default")
	symaddr(&addr_QAccessibleApplication_child_default, "QAccessibleApplication_child_default")
	symaddr(&addr_QAccessibleApplication_text_default, "QAccessibleApplication_text_default")
	symaddr(&addr_QAccessibleApplication_role_default, "QAccessibleApplication_role_default")
	symaddr(&addr_QAccessibleApplication_state_default, "QAccessibleApplication_state_default")
	symaddr(&addr_QAccessibleApplication_isValid_default, "QAccessibleApplication_isValid_default")
	symaddr(&addr_QAccessibleApplication_object_default, "QAccessibleApplication_object_default")
	symaddr(&addr_QAccessibleApplication_rect_default, "QAccessibleApplication_rect_default")
	symaddr(&addr_QAccessibleApplication_setText_default, "QAccessibleApplication_setText_default")
	symaddr(&addr_QAccessibleApplication_childAt_default, "QAccessibleApplication_childAt_default")
	symaddr(&addr_QAccessibleApplication_foregroundColor_default, "QAccessibleApplication_foregroundColor_default")
	symaddr(&addr_QAccessibleApplication_backgroundColor_default, "QAccessibleApplication_backgroundColor_default")
	reg(&pg_QAccessibleApplication__set_override, "QAccessibleApplication__set_override")
	reg(&pg_QAccessibleApplication_delete, "QAccessibleApplication_delete")
}

type QAccessibleApplication struct {
	h uintptr
	*QAccessibleObject
}

func bindQAccessibleApplication(h uintptr) *QAccessibleApplication {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAccessibleApplication_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleApplication{h: h, QAccessibleObject: &QAccessibleObject{h: __c[0], QAccessibleInterface: &QAccessibleInterface{h: __c[1]}}}
}

func newownedQAccessibleApplication(h uintptr) *QAccessibleApplication {
	o := bindQAccessibleApplication(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleApplication).Delete)
	}
	return o
}

func (this *QAccessibleApplication) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleApplication) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleApplication(h unsafe.Pointer) *QAccessibleApplication { return bindQAccessibleApplication(uintptr(h)) }
func UnsafeNewQAccessibleApplicationOwned(h unsafe.Pointer) *QAccessibleApplication { return newownedQAccessibleApplication(uintptr(h)) }

func (this *QAccessibleApplication) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleObject != nil {
		this.QAccessibleObject.UnsafeRelease()
	}
}

func NewQAccessibleApplication() *QAccessibleApplication {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_new, uintptr(__goid))
	__o := newownedQAccessibleApplication(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAccessibleApplication_virtIdx, QAccessibleApplication_enable)
	return __o
}

func (this *QAccessibleApplication) WindowDefault() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_window_default, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QAccessibleApplication) ChildCountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleApplication_childCount_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleApplication) IndexOfChildDefault(arg1 *QAccessibleInterface) int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleApplication_indexOfChild_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QAccessibleApplication) FocusChildDefault() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_focusChild_default, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleApplication) ParentDefault() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_parent_default, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleApplication) ChildDefault(index int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_child_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleApplication) TextDefault(t QAccessible__Text) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_text_default, this.h, uintptr(t), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleApplication) RoleDefault() QAccessible__Role {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleApplication_role_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAccessible__Role(__r)
}

func (this *QAccessibleApplication) StateDefault() *QAccessible__State {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_state_default, this.h)
	runtime.KeepAlive(this)
	return newownedQAccessible__State(__r)
}

func (this *QAccessibleApplication) IsValidDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleApplication_isValid_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleApplication) ObjectDefault() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_object_default, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessibleApplication) RectDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_rect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QAccessibleApplication) SetTextDefault(t QAccessible__Text, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAccessibleApplication_setText_default, this.h, uintptr(t), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAccessibleApplication) ChildAtDefault(x int, y int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_childAt_default, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleApplication) ForegroundColorDefault() *QColor {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_foregroundColor_default, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QAccessibleApplication) BackgroundColorDefault() *QColor {
	__r, _, _ := purego.SyscallN(addr_QAccessibleApplication_backgroundColor_default, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func QAccessibleApplication_virtIdx(slot int64) int32 {
	switch slot {
	case 8210911950352161981:
		return 0
	case 1383730954482307804:
		return 1
	case -9019360399188768292:
		return 2
	case -6961835385172491989:
		return 3
	case 7371339085820260189:
		return 4
	case 7071834996910082205:
		return 5
	case -430950356557152130:
		return 6
	case -6676326825648966727:
		return 7
	case -1268983008813010666:
		return 8
	case 7231157875507698241:
		return 9
	case -8214935597255553030:
		return 10
	case 5542927423540195655:
		return 11
	case -5557558535206160508:
		return 12
	case 3663933090421434080:
		return 13
	case -3317306646323195305:
		return 14
	case 2962551866395559440:
		return 15
	}
	return -1
}

func QAccessibleApplication_enable(h uintptr, idx int32) { pg_QAccessibleApplication__set_override(h, idx) }

func (this *QAccessibleApplication) OnWindow(fn func() *QWindow) {
	registerVirt(this.UnsafeGoID(), 8210911950352161981, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "window"
}

func (this *QAccessibleApplication) OnChildCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), 1383730954482307804, func() int32 { return int32(fn()) }) // slot "childCount"
}

func (this *QAccessibleApplication) OnIndexOfChild(fn func(*QAccessibleInterface) int) {
	registerVirt(this.UnsafeGoID(), -9019360399188768292, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQAccessibleInterface(unsafe.Pointer(arg1)))) }) // slot "indexOfChild"
}

func (this *QAccessibleApplication) OnFocusChild(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -6961835385172491989, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "focusChild"
}

func (this *QAccessibleApplication) OnParent(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "parent"
}

func (this *QAccessibleApplication) OnChild(fn func(int) *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 7071834996910082205, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "child"
}

func (this *QAccessibleApplication) OnText(fn func(QAccessible__Text) string) {
	registerVirt(this.UnsafeGoID(), -430950356557152130, func(t int32) (uintptr, int64) {
		__s := fn(QAccessible__Text(t))
		return pgCString(__s), int64(len(__s))
	}) // slot "text"
}

func (this *QAccessibleApplication) OnRole(fn func() QAccessible__Role) {
	registerVirt(this.UnsafeGoID(), -6676326825648966727, func() int32 { return int32(fn()) }) // slot "role"
}

func (this *QAccessibleApplication) OnState(fn func() *QAccessible__State) {
	registerVirt(this.UnsafeGoID(), -1268983008813010666, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "state"
}

func (this *QAccessibleApplication) OnIsValid(fn func() bool) {
	registerVirt(this.UnsafeGoID(), 7231157875507698241, func() int32 { return b2i(fn()) }) // slot "isValid"
}

func (this *QAccessibleApplication) OnObject(fn func() *QObject) {
	registerVirt(this.UnsafeGoID(), -8214935597255553030, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "object"
}

func (this *QAccessibleApplication) OnRect(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 5542927423540195655, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "rect"
}

func (this *QAccessibleApplication) OnSetText(fn func(QAccessible__Text, string)) {
	registerVirt(this.UnsafeGoID(), -5557558535206160508, func(t int32, text uintptr, text_len int64) { fn(QAccessible__Text(t), pgStrNoFree(text, text_len)) }) // slot "setText"
}

func (this *QAccessibleApplication) OnChildAt(fn func(int, int) *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 3663933090421434080, func(x int32, y int32) uintptr { return uintptr(fn(int(x), int(y)).UnsafePointer()) }) // slot "childAt"
}

func (this *QAccessibleApplication) OnForegroundColor(fn func() *QColor) {
	registerVirt(this.UnsafeGoID(), -3317306646323195305, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "foregroundColor"
}

func (this *QAccessibleApplication) OnBackgroundColor(fn func() *QColor) {
	registerVirt(this.UnsafeGoID(), 2962551866395559440, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "backgroundColor"
}

func (this *QAccessibleApplication) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleApplication_delete(this.h)
	}
	this.UnsafeRelease()
}
