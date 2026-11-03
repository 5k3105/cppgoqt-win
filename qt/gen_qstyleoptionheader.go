package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionHeader_base func(uintptr) uintptr
	addr_QStyleOptionHeader_new uintptr
	addr_QStyleOptionHeader_new2 uintptr
	addr_QStyleOptionHeader_OperatorAssign uintptr
	addr_QStyleOptionHeader_section_get uintptr
	addr_QStyleOptionHeader_section_set uintptr
	addr_QStyleOptionHeader_text_get uintptr
	addr_QStyleOptionHeader_text_set uintptr
	addr_QStyleOptionHeader_textAlignment_get uintptr
	addr_QStyleOptionHeader_textAlignment_set uintptr
	addr_QStyleOptionHeader_icon_get uintptr
	addr_QStyleOptionHeader_icon_set uintptr
	addr_QStyleOptionHeader_iconAlignment_get uintptr
	addr_QStyleOptionHeader_iconAlignment_set uintptr
	addr_QStyleOptionHeader_position_get uintptr
	addr_QStyleOptionHeader_position_set uintptr
	addr_QStyleOptionHeader_selectedPosition_get uintptr
	addr_QStyleOptionHeader_selectedPosition_set uintptr
	addr_QStyleOptionHeader_sortIndicator_get uintptr
	addr_QStyleOptionHeader_sortIndicator_set uintptr
	addr_QStyleOptionHeader_orientation_get uintptr
	addr_QStyleOptionHeader_orientation_set uintptr
	pg_QStyleOptionHeader_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionHeader_base, "QStyleOptionHeader_base")
	symaddr(&addr_QStyleOptionHeader_new, "QStyleOptionHeader_new")
	symaddr(&addr_QStyleOptionHeader_new2, "QStyleOptionHeader_new2")
	symaddr(&addr_QStyleOptionHeader_OperatorAssign, "QStyleOptionHeader_OperatorAssign")
	symaddr(&addr_QStyleOptionHeader_section_get, "QStyleOptionHeader_section_get")
	symaddr(&addr_QStyleOptionHeader_section_set, "QStyleOptionHeader_section_set")
	symaddr(&addr_QStyleOptionHeader_text_get, "QStyleOptionHeader_text_get")
	symaddr(&addr_QStyleOptionHeader_text_set, "QStyleOptionHeader_text_set")
	symaddr(&addr_QStyleOptionHeader_textAlignment_get, "QStyleOptionHeader_textAlignment_get")
	symaddr(&addr_QStyleOptionHeader_textAlignment_set, "QStyleOptionHeader_textAlignment_set")
	symaddr(&addr_QStyleOptionHeader_icon_get, "QStyleOptionHeader_icon_get")
	symaddr(&addr_QStyleOptionHeader_icon_set, "QStyleOptionHeader_icon_set")
	symaddr(&addr_QStyleOptionHeader_iconAlignment_get, "QStyleOptionHeader_iconAlignment_get")
	symaddr(&addr_QStyleOptionHeader_iconAlignment_set, "QStyleOptionHeader_iconAlignment_set")
	symaddr(&addr_QStyleOptionHeader_position_get, "QStyleOptionHeader_position_get")
	symaddr(&addr_QStyleOptionHeader_position_set, "QStyleOptionHeader_position_set")
	symaddr(&addr_QStyleOptionHeader_selectedPosition_get, "QStyleOptionHeader_selectedPosition_get")
	symaddr(&addr_QStyleOptionHeader_selectedPosition_set, "QStyleOptionHeader_selectedPosition_set")
	symaddr(&addr_QStyleOptionHeader_sortIndicator_get, "QStyleOptionHeader_sortIndicator_get")
	symaddr(&addr_QStyleOptionHeader_sortIndicator_set, "QStyleOptionHeader_sortIndicator_set")
	symaddr(&addr_QStyleOptionHeader_orientation_get, "QStyleOptionHeader_orientation_get")
	symaddr(&addr_QStyleOptionHeader_orientation_set, "QStyleOptionHeader_orientation_set")
	reg(&pg_QStyleOptionHeader_delete, "QStyleOptionHeader_delete")
}

type QStyleOptionHeader struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionHeader(h uintptr) *QStyleOptionHeader {
	if h == 0 {
		return nil
	}
	return &QStyleOptionHeader{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionHeader_base(h))}
}

func newownedQStyleOptionHeader(h uintptr) *QStyleOptionHeader {
	o := bindQStyleOptionHeader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionHeader).Delete)
	}
	return o
}

func (this *QStyleOptionHeader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionHeader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionHeader(h unsafe.Pointer) *QStyleOptionHeader { return bindQStyleOptionHeader(uintptr(h)) }
func UnsafeNewQStyleOptionHeaderOwned(h unsafe.Pointer) *QStyleOptionHeader { return newownedQStyleOptionHeader(uintptr(h)) }

func (this *QStyleOptionHeader) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionHeader() *QStyleOptionHeader {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeader_new)
	__o := newownedQStyleOptionHeader(__r)
	return __o
}

func NewQStyleOptionHeader2(other *QStyleOptionHeader) *QStyleOptionHeader {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeader_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionHeader(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionHeader) OperatorAssign(arg1 *QStyleOptionHeader) *QStyleOptionHeader {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeader_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionHeader(__r)
}

func (this *QStyleOptionHeader) Section() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_section_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOptionHeader) SetSection(section int) {
	purego.SyscallN(addr_QStyleOptionHeader_section_set, this.h, uintptr(section))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeader_text_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStyleOptionHeader) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStyleOptionHeader_text_set, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStyleOptionHeader) TextAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_textAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionHeader) SetTextAlignment(textAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionHeader_textAlignment_set, this.h, uintptr(textAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionHeader_icon_get, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStyleOptionHeader) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStyleOptionHeader_icon_set, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStyleOptionHeader) IconAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_iconAlignment_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStyleOptionHeader) SetIconAlignment(iconAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStyleOptionHeader_iconAlignment_set, this.h, uintptr(iconAlignment))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) Position() QStyleOptionHeader__SectionPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_position_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionHeader__SectionPosition(__r)
}

func (this *QStyleOptionHeader) SetPosition(position QStyleOptionHeader__SectionPosition) {
	purego.SyscallN(addr_QStyleOptionHeader_position_set, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) SelectedPosition() QStyleOptionHeader__SelectedPosition {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_selectedPosition_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionHeader__SelectedPosition(__r)
}

func (this *QStyleOptionHeader) SetSelectedPosition(selectedPosition QStyleOptionHeader__SelectedPosition) {
	purego.SyscallN(addr_QStyleOptionHeader_selectedPosition_set, this.h, uintptr(selectedPosition))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) SortIndicator() QStyleOptionHeader__SortIndicator {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_sortIndicator_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyleOptionHeader__SortIndicator(__r)
}

func (this *QStyleOptionHeader) SetSortIndicator(sortIndicator QStyleOptionHeader__SortIndicator) {
	purego.SyscallN(addr_QStyleOptionHeader_sortIndicator_set, this.h, uintptr(sortIndicator))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionHeader_orientation_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QStyleOptionHeader) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QStyleOptionHeader_orientation_set, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionHeader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionHeader_delete(this.h)
	}
	this.UnsafeRelease()
}
