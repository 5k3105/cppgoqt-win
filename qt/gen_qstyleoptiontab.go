package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionTab_base func(uintptr) uintptr
	addr_QStyleOptionTab_new uintptr
	addr_QStyleOptionTab_new2 uintptr
	addr_QStyleOptionTab_OperatorAssign uintptr
	addr_QStyleOptionTab_shape_get uintptr
	addr_QStyleOptionTab_shape_set uintptr
	addr_QStyleOptionTab_text_get uintptr
	addr_QStyleOptionTab_text_set uintptr
	addr_QStyleOptionTab_icon_get uintptr
	addr_QStyleOptionTab_icon_set uintptr
	addr_QStyleOptionTab_row_get uintptr
	addr_QStyleOptionTab_row_set uintptr
	addr_QStyleOptionTab_position_get uintptr
	addr_QStyleOptionTab_position_set uintptr
	addr_QStyleOptionTab_selectedPosition_get uintptr
	addr_QStyleOptionTab_selectedPosition_set uintptr
	addr_QStyleOptionTab_cornerWidgets_get uintptr
	addr_QStyleOptionTab_cornerWidgets_set uintptr
	addr_QStyleOptionTab_iconSize_get uintptr
	addr_QStyleOptionTab_iconSize_set uintptr
	addr_QStyleOptionTab_documentMode_get uintptr
	addr_QStyleOptionTab_documentMode_set uintptr
	addr_QStyleOptionTab_leftButtonSize_get uintptr
	addr_QStyleOptionTab_leftButtonSize_set uintptr
	addr_QStyleOptionTab_rightButtonSize_get uintptr
	addr_QStyleOptionTab_rightButtonSize_set uintptr
	addr_QStyleOptionTab_features_get uintptr
	addr_QStyleOptionTab_features_set uintptr
	addr_QStyleOptionTab_tabIndex_get uintptr
	addr_QStyleOptionTab_tabIndex_set uintptr
	pg_QStyleOptionTab_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionTab_base, "QStyleOptionTab_base")
	symaddr(&addr_QStyleOptionTab_new, "QStyleOptionTab_new")
	symaddr(&addr_QStyleOptionTab_new2, "QStyleOptionTab_new2")
	symaddr(&addr_QStyleOptionTab_OperatorAssign, "QStyleOptionTab_OperatorAssign")
	symaddr(&addr_QStyleOptionTab_shape_get, "QStyleOptionTab_shape_get")
	symaddr(&addr_QStyleOptionTab_shape_set, "QStyleOptionTab_shape_set")
	symaddr(&addr_QStyleOptionTab_text_get, "QStyleOptionTab_text_get")
	symaddr(&addr_QStyleOptionTab_text_set, "QStyleOptionTab_text_set")
	symaddr(&addr_QStyleOptionTab_icon_get, "QStyleOptionTab_icon_get")
	symaddr(&addr_QStyleOptionTab_icon_set, "QStyleOptionTab_icon_set")
	symaddr(&addr_QStyleOptionTab_row_get, "QStyleOptionTab_row_get")
	symaddr(&addr_QStyleOptionTab_row_set, "QStyleOptionTab_row_set")
	symaddr(&addr_QStyleOptionTab_position_get, "QStyleOptionTab_position_get")
	symaddr(&addr_QStyleOptionTab_position_set, "QStyleOptionTab_position_set")
	symaddr(&addr_QStyleOptionTab_selectedPosition_get, "QStyleOptionTab_selectedPosition_get")
	symaddr(&addr_QStyleOptionTab_selectedPosition_set, "QStyleOptionTab_selectedPosition_set")
	symaddr(&addr_QStyleOptionTab_cornerWidgets_get, "QStyleOptionTab_cornerWidgets_get")
	symaddr(&addr_QStyleOptionTab_cornerWidgets_set, "QStyleOptionTab_cornerWidgets_set")
	symaddr(&addr_QStyleOptionTab_iconSize_get, "QStyleOptionTab_iconSize_get")
	symaddr(&addr_QStyleOptionTab_iconSize_set, "QStyleOptionTab_iconSize_set")
	symaddr(&addr_QStyleOptionTab_documentMode_get, "QStyleOptionTab_documentMode_get")
	symaddr(&addr_QStyleOptionTab_documentMode_set, "QStyleOptionTab_documentMode_set")
	symaddr(&addr_QStyleOptionTab_leftButtonSize_get, "QStyleOptionTab_leftButtonSize_get")
	symaddr(&addr_QStyleOptionTab_leftButtonSize_set, "QStyleOptionTab_leftButtonSize_set")
	symaddr(&addr_QStyleOptionTab_rightButtonSize_get, "QStyleOptionTab_rightButtonSize_get")
	symaddr(&addr_QStyleOptionTab_rightButtonSize_set, "QStyleOptionTab_rightButtonSize_set")
	symaddr(&addr_QStyleOptionTab_features_get, "QStyleOptionTab_features_get")
	symaddr(&addr_QStyleOptionTab_features_set, "QStyleOptionTab_features_set")
	symaddr(&addr_QStyleOptionTab_tabIndex_get, "QStyleOptionTab_tabIndex_get")
	symaddr(&addr_QStyleOptionTab_tabIndex_set, "QStyleOptionTab_tabIndex_set")
	reg(&pg_QStyleOptionTab_delete, "QStyleOptionTab_delete")
}

type QStyleOptionTab struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionTab(h uintptr) *QStyleOptionTab {
	if h == 0 {
		return nil
	}
	return &QStyleOptionTab{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionTab_base(h))}
}

func newownedQStyleOptionTab(h uintptr) *QStyleOptionTab {
	o := bindQStyleOptionTab(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionTab).Delete)
	}
	return o
}

func (this *QStyleOptionTab) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionTab) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionTab(h unsafe.Pointer) *QStyleOptionTab { return bindQStyleOptionTab(uintptr(h)) }
func UnsafeNewQStyleOptionTabOwned(h unsafe.Pointer) *QStyleOptionTab { return newownedQStyleOptionTab(uintptr(h)) }

func (this *QStyleOptionTab) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionTab() *QStyleOptionTab {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_new)
	__o := newownedQStyleOptionTab(__r)
	return __o
}

func NewQStyleOptionTab2(other *QStyleOptionTab) *QStyleOptionTab {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionTab(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionTab) OperatorAssign(arg1 *QStyleOptionTab) *QStyleOptionTab {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionTab(__r)
}

func (this *QStyleOptionTab) Shape() QTabBar__Shape {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_shape_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTabBar__Shape(__r)
}

func (this *QStyleOptionTab) SetShape(shape QTabBar__Shape) {
	purego.SyscallN(addr_QStyleOptionTab_shape_set, this.h, uintptr(shape))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionTab) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionTab_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionTab) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionTab) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionTab_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionTab) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_row_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionTab) SetRow(row int) {
	purego.SyscallN(addr_QStyleOptionTab_row_set, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) Position() QStyleOptionTab__TabPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_position_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionTab__TabPosition(__r)
}

func (this *QStyleOptionTab) SetPosition(position QStyleOptionTab__TabPosition) {
	purego.SyscallN(addr_QStyleOptionTab_position_set, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) SelectedPosition() QStyleOptionTab__SelectedPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_selectedPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionTab__SelectedPosition(__r)
}

func (this *QStyleOptionTab) SetSelectedPosition(selectedPosition QStyleOptionTab__SelectedPosition) {
	purego.SyscallN(addr_QStyleOptionTab_selectedPosition_set, this.h, uintptr(selectedPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) CornerWidgets() QStyleOptionTab__CornerWidget {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_cornerWidgets_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionTab__CornerWidget(__r)
}

func (this *QStyleOptionTab) SetCornerWidgets(cornerWidgets QStyleOptionTab__CornerWidget) {
	purego.SyscallN(addr_QStyleOptionTab_cornerWidgets_set, this.h, uintptr(cornerWidgets))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_iconSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTab) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTab_iconSize_set, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QStyleOptionTab) DocumentMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_documentMode_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleOptionTab) SetDocumentMode(documentMode bool) {
	purego.SyscallN(addr_QStyleOptionTab_documentMode_set, this.h, uintptr(b2i(documentMode)))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) LeftButtonSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_leftButtonSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTab) SetLeftButtonSize(leftButtonSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTab_leftButtonSize_set, this.h, uintptr(leftButtonSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(leftButtonSize)
}

func (this *QStyleOptionTab) RightButtonSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionTab_rightButtonSize_get, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStyleOptionTab) SetRightButtonSize(rightButtonSize *QSize) {
	purego.SyscallN(addr_QStyleOptionTab_rightButtonSize_set, this.h, uintptr(rightButtonSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rightButtonSize)
}

func (this *QStyleOptionTab) Features() QStyleOptionTab__TabFeature {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_features_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionTab__TabFeature(__r)
}

func (this *QStyleOptionTab) SetFeatures(features QStyleOptionTab__TabFeature) {
	purego.SyscallN(addr_QStyleOptionTab_features_set, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) TabIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionTab_tabIndex_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionTab) SetTabIndex(tabIndex int) {
	purego.SyscallN(addr_QStyleOptionTab_tabIndex_set, this.h, uintptr(tabIndex))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionTab) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionTab_delete(this.h)
	}
	this.UnsafeRelease()
}
