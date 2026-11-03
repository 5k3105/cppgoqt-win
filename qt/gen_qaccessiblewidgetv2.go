package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleWidgetV2_chain uintptr
	addr_QAccessibleWidgetV2_new uintptr
	addr_QAccessibleWidgetV2_new2 uintptr
	addr_QAccessibleWidgetV2_new3 uintptr
	addr_QAccessibleWidgetV2_attributeKeys uintptr
	addr_QAccessibleWidgetV2_attributeValue uintptr
	addr_QAccessibleWidgetV2_tr uintptr
	addr_QAccessibleWidgetV2_pressAction uintptr
	addr_QAccessibleWidgetV2_increaseAction uintptr
	addr_QAccessibleWidgetV2_decreaseAction uintptr
	addr_QAccessibleWidgetV2_showMenuAction uintptr
	addr_QAccessibleWidgetV2_setFocusAction uintptr
	addr_QAccessibleWidgetV2_toggleAction uintptr
	addr_QAccessibleWidgetV2_scrollLeftAction uintptr
	addr_QAccessibleWidgetV2_scrollRightAction uintptr
	addr_QAccessibleWidgetV2_scrollUpAction uintptr
	addr_QAccessibleWidgetV2_scrollDownAction uintptr
	addr_QAccessibleWidgetV2_nextPageAction uintptr
	addr_QAccessibleWidgetV2_previousPageAction uintptr
	addr_QAccessibleWidgetV2_tr2 uintptr
	addr_QAccessibleWidgetV2_tr3 uintptr
	addr_QAccessibleWidgetV2_asQAccessibleAttributesInterface uintptr
)

func init() {
	symaddr(&addr_QAccessibleWidgetV2_chain, "QAccessibleWidgetV2_chain")
	symaddr(&addr_QAccessibleWidgetV2_new, "QAccessibleWidgetV2_new")
	symaddr(&addr_QAccessibleWidgetV2_new2, "QAccessibleWidgetV2_new2")
	symaddr(&addr_QAccessibleWidgetV2_new3, "QAccessibleWidgetV2_new3")
	symaddr(&addr_QAccessibleWidgetV2_attributeKeys, "QAccessibleWidgetV2_attributeKeys")
	symaddr(&addr_QAccessibleWidgetV2_attributeValue, "QAccessibleWidgetV2_attributeValue")
	symaddr(&addr_QAccessibleWidgetV2_tr, "QAccessibleWidgetV2_tr")
	symaddr(&addr_QAccessibleWidgetV2_pressAction, "QAccessibleWidgetV2_pressAction")
	symaddr(&addr_QAccessibleWidgetV2_increaseAction, "QAccessibleWidgetV2_increaseAction")
	symaddr(&addr_QAccessibleWidgetV2_decreaseAction, "QAccessibleWidgetV2_decreaseAction")
	symaddr(&addr_QAccessibleWidgetV2_showMenuAction, "QAccessibleWidgetV2_showMenuAction")
	symaddr(&addr_QAccessibleWidgetV2_setFocusAction, "QAccessibleWidgetV2_setFocusAction")
	symaddr(&addr_QAccessibleWidgetV2_toggleAction, "QAccessibleWidgetV2_toggleAction")
	symaddr(&addr_QAccessibleWidgetV2_scrollLeftAction, "QAccessibleWidgetV2_scrollLeftAction")
	symaddr(&addr_QAccessibleWidgetV2_scrollRightAction, "QAccessibleWidgetV2_scrollRightAction")
	symaddr(&addr_QAccessibleWidgetV2_scrollUpAction, "QAccessibleWidgetV2_scrollUpAction")
	symaddr(&addr_QAccessibleWidgetV2_scrollDownAction, "QAccessibleWidgetV2_scrollDownAction")
	symaddr(&addr_QAccessibleWidgetV2_nextPageAction, "QAccessibleWidgetV2_nextPageAction")
	symaddr(&addr_QAccessibleWidgetV2_previousPageAction, "QAccessibleWidgetV2_previousPageAction")
	symaddr(&addr_QAccessibleWidgetV2_tr2, "QAccessibleWidgetV2_tr2")
	symaddr(&addr_QAccessibleWidgetV2_tr3, "QAccessibleWidgetV2_tr3")
	symaddr(&addr_QAccessibleWidgetV2_asQAccessibleAttributesInterface, "QAccessibleWidgetV2_asQAccessibleAttributesInterface")
}

type QAccessibleWidgetV2 struct {
	h uintptr
	*QAccessibleWidget
}

func bindQAccessibleWidgetV2(h uintptr) *QAccessibleWidgetV2 {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QAccessibleWidgetV2_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAccessibleWidgetV2{h: h, QAccessibleWidget: &QAccessibleWidget{h: __c[0], QAccessibleObject: &QAccessibleObject{h: __c[1], QAccessibleInterface: &QAccessibleInterface{h: __c[2]}}}}
}

func newownedQAccessibleWidgetV2(h uintptr) *QAccessibleWidgetV2 { return bindQAccessibleWidgetV2(h) }

func (this *QAccessibleWidgetV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleWidgetV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleWidgetV2(h unsafe.Pointer) *QAccessibleWidgetV2 { return bindQAccessibleWidgetV2(uintptr(h)) }
func UnsafeNewQAccessibleWidgetV2Owned(h unsafe.Pointer) *QAccessibleWidgetV2 { return newownedQAccessibleWidgetV2(uintptr(h)) }

func (this *QAccessibleWidgetV2) UnsafeRelease() {
	this.h = 0
	if this.QAccessibleWidget != nil {
		this.QAccessibleWidget.UnsafeRelease()
	}
}

func NewQAccessibleWidgetV2(object *QWidget) *QAccessibleWidgetV2 {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_new, uintptr(object.UnsafePointer()))
	__o := newownedQAccessibleWidgetV2(__r)
	runtime.KeepAlive(object)
	return __o
}

func NewQAccessibleWidgetV22(object *QWidget, role QAccessible__Role, name string) *QAccessibleWidgetV2 {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_new2, uintptr(object.UnsafePointer()), uintptr(role), name_p, uintptr(len(name)))
	__o := newownedQAccessibleWidgetV2(__r)
	runtime.KeepAlive(object)
	runtime.KeepAlive(name)
	return __o
}

func NewQAccessibleWidgetV23(object *QWidget, role QAccessible__Role) *QAccessibleWidgetV2 {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_new3, uintptr(object.UnsafePointer()), uintptr(role))
	__o := newownedQAccessibleWidgetV2(__r)
	runtime.KeepAlive(object)
	return __o
}

func (this *QAccessibleWidgetV2) AttributeKeys() []QAccessible__Attribute {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_attributeKeys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]QAccessible__Attribute, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = QAccessible__Attribute(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleWidgetV2) AttributeValue(key QAccessible__Attribute) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_attributeValue, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QAccessibleWidgetV2_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_PressAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_pressAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_IncreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_increaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_DecreaseAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_decreaseAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ShowMenuAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_showMenuAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_SetFocusAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_setFocusAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ToggleAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_toggleAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ScrollLeftAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_scrollLeftAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ScrollRightAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_scrollRightAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ScrollUpAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_scrollUpAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_ScrollDownAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_scrollDownAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_NextPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_nextPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_PreviousPageAction() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_previousPageAction, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QAccessibleWidgetV2_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QAccessibleWidgetV2) AsQAccessibleAttributesInterface() *QAccessibleAttributesInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleWidgetV2_asQAccessibleAttributesInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleAttributesInterface(__r)
}

func (this *QAccessibleWidgetV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
