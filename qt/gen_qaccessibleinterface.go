package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleInterface_isValid uintptr
	addr_QAccessibleInterface_object uintptr
	addr_QAccessibleInterface_window uintptr
	addr_QAccessibleInterface_focusChild uintptr
	addr_QAccessibleInterface_childAt uintptr
	addr_QAccessibleInterface_parent uintptr
	addr_QAccessibleInterface_child uintptr
	addr_QAccessibleInterface_childCount uintptr
	addr_QAccessibleInterface_indexOfChild uintptr
	addr_QAccessibleInterface_text uintptr
	addr_QAccessibleInterface_setText uintptr
	addr_QAccessibleInterface_rect uintptr
	addr_QAccessibleInterface_role uintptr
	addr_QAccessibleInterface_state uintptr
	addr_QAccessibleInterface_foregroundColor uintptr
	addr_QAccessibleInterface_backgroundColor uintptr
	addr_QAccessibleInterface_textInterface uintptr
	addr_QAccessibleInterface_editableTextInterface uintptr
	addr_QAccessibleInterface_valueInterface uintptr
	addr_QAccessibleInterface_actionInterface uintptr
	addr_QAccessibleInterface_imageInterface uintptr
	addr_QAccessibleInterface_tableInterface uintptr
	addr_QAccessibleInterface_tableCellInterface uintptr
	addr_QAccessibleInterface_hyperlinkInterface uintptr
	addr_QAccessibleInterface_selectionInterface uintptr
	addr_QAccessibleInterface_attributesInterface uintptr
	addr_QAccessibleInterface_OperatorAssign uintptr
)

func init() {
	symaddr(&addr_QAccessibleInterface_isValid, "QAccessibleInterface_isValid")
	symaddr(&addr_QAccessibleInterface_object, "QAccessibleInterface_object")
	symaddr(&addr_QAccessibleInterface_window, "QAccessibleInterface_window")
	symaddr(&addr_QAccessibleInterface_focusChild, "QAccessibleInterface_focusChild")
	symaddr(&addr_QAccessibleInterface_childAt, "QAccessibleInterface_childAt")
	symaddr(&addr_QAccessibleInterface_parent, "QAccessibleInterface_parent")
	symaddr(&addr_QAccessibleInterface_child, "QAccessibleInterface_child")
	symaddr(&addr_QAccessibleInterface_childCount, "QAccessibleInterface_childCount")
	symaddr(&addr_QAccessibleInterface_indexOfChild, "QAccessibleInterface_indexOfChild")
	symaddr(&addr_QAccessibleInterface_text, "QAccessibleInterface_text")
	symaddr(&addr_QAccessibleInterface_setText, "QAccessibleInterface_setText")
	symaddr(&addr_QAccessibleInterface_rect, "QAccessibleInterface_rect")
	symaddr(&addr_QAccessibleInterface_role, "QAccessibleInterface_role")
	symaddr(&addr_QAccessibleInterface_state, "QAccessibleInterface_state")
	symaddr(&addr_QAccessibleInterface_foregroundColor, "QAccessibleInterface_foregroundColor")
	symaddr(&addr_QAccessibleInterface_backgroundColor, "QAccessibleInterface_backgroundColor")
	symaddr(&addr_QAccessibleInterface_textInterface, "QAccessibleInterface_textInterface")
	symaddr(&addr_QAccessibleInterface_editableTextInterface, "QAccessibleInterface_editableTextInterface")
	symaddr(&addr_QAccessibleInterface_valueInterface, "QAccessibleInterface_valueInterface")
	symaddr(&addr_QAccessibleInterface_actionInterface, "QAccessibleInterface_actionInterface")
	symaddr(&addr_QAccessibleInterface_imageInterface, "QAccessibleInterface_imageInterface")
	symaddr(&addr_QAccessibleInterface_tableInterface, "QAccessibleInterface_tableInterface")
	symaddr(&addr_QAccessibleInterface_tableCellInterface, "QAccessibleInterface_tableCellInterface")
	symaddr(&addr_QAccessibleInterface_hyperlinkInterface, "QAccessibleInterface_hyperlinkInterface")
	symaddr(&addr_QAccessibleInterface_selectionInterface, "QAccessibleInterface_selectionInterface")
	symaddr(&addr_QAccessibleInterface_attributesInterface, "QAccessibleInterface_attributesInterface")
	symaddr(&addr_QAccessibleInterface_OperatorAssign, "QAccessibleInterface_OperatorAssign")
}

type QAccessibleInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleInterface(h uintptr) *QAccessibleInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleInterface{h: h}
}

func newownedQAccessibleInterface(h uintptr) *QAccessibleInterface { return bindQAccessibleInterface(h) }

func (this *QAccessibleInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleInterface(h unsafe.Pointer) *QAccessibleInterface { return bindQAccessibleInterface(uintptr(h)) }
func UnsafeNewQAccessibleInterfaceOwned(h unsafe.Pointer) *QAccessibleInterface { return newownedQAccessibleInterface(uintptr(h)) }

func (this *QAccessibleInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleInterface) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleInterface_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleInterface) Object() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_object, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAccessibleInterface) Window() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_window, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QAccessibleInterface) FocusChild() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_focusChild, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleInterface) ChildAt(x int, y int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_childAt, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleInterface) Parent() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_parent, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleInterface) Child(index int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_child, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleInterface) ChildCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleInterface_childCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleInterface) IndexOfChild(arg1 *QAccessibleInterface) int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleInterface_indexOfChild, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QAccessibleInterface) Text(t QAccessible__Text) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_text, this.h, uintptr(t), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleInterface) SetText(t QAccessible__Text, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAccessibleInterface_setText, this.h, uintptr(t), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAccessibleInterface) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QAccessibleInterface) Role() QAccessible__Role {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleInterface_role, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAccessible__Role(__r)
}

func (this *QAccessibleInterface) State() *QAccessible__State {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_state, this.h)
	runtime.KeepAlive(this)
	return newownedQAccessible__State(__r)
}

func (this *QAccessibleInterface) ForegroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_foregroundColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QAccessibleInterface) BackgroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_backgroundColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QAccessibleInterface) TextInterface() *QAccessibleTextInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_textInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleTextInterface(__r)
}

func (this *QAccessibleInterface) EditableTextInterface() *QAccessibleEditableTextInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_editableTextInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleEditableTextInterface(__r)
}

func (this *QAccessibleInterface) ValueInterface() *QAccessibleValueInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_valueInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleValueInterface(__r)
}

func (this *QAccessibleInterface) ActionInterface() *QAccessibleActionInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_actionInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleActionInterface(__r)
}

func (this *QAccessibleInterface) ImageInterface() *QAccessibleImageInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_imageInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleImageInterface(__r)
}

func (this *QAccessibleInterface) TableInterface() *QAccessibleTableInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_tableInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleTableInterface(__r)
}

func (this *QAccessibleInterface) TableCellInterface() *QAccessibleTableCellInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_tableCellInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleTableCellInterface(__r)
}

func (this *QAccessibleInterface) HyperlinkInterface() *QAccessibleHyperlinkInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_hyperlinkInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleHyperlinkInterface(__r)
}

func (this *QAccessibleInterface) SelectionInterface() *QAccessibleSelectionInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_selectionInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleSelectionInterface(__r)
}

func (this *QAccessibleInterface) AttributesInterface() *QAccessibleAttributesInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_attributesInterface, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleAttributesInterface(__r)
}

func (this *QAccessibleInterface) OperatorAssign(arg1 *QAccessibleInterface) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
