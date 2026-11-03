package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSizePolicy_new uintptr
	addr_QSizePolicy_new2 uintptr
	addr_QSizePolicy_new3 uintptr
	addr_QSizePolicy_new4 uintptr
	addr_QSizePolicy_horizontalPolicy uintptr
	addr_QSizePolicy_verticalPolicy uintptr
	addr_QSizePolicy_controlType uintptr
	addr_QSizePolicy_setHorizontalPolicy uintptr
	addr_QSizePolicy_setVerticalPolicy uintptr
	addr_QSizePolicy_setControlType uintptr
	addr_QSizePolicy_expandingDirections uintptr
	addr_QSizePolicy_setHeightForWidth uintptr
	addr_QSizePolicy_hasHeightForWidth uintptr
	addr_QSizePolicy_setWidthForHeight uintptr
	addr_QSizePolicy_hasWidthForHeight uintptr
	addr_QSizePolicy_OperatorEqual uintptr
	addr_QSizePolicy_OperatorNotEqual uintptr
	addr_QSizePolicy_ToQVariant uintptr
	addr_QSizePolicy_horizontalStretch uintptr
	addr_QSizePolicy_verticalStretch uintptr
	addr_QSizePolicy_setHorizontalStretch uintptr
	addr_QSizePolicy_setVerticalStretch uintptr
	addr_QSizePolicy_retainSizeWhenHidden uintptr
	addr_QSizePolicy_setRetainSizeWhenHidden uintptr
	addr_QSizePolicy_transpose uintptr
	addr_QSizePolicy_transposed uintptr
	addr_QSizePolicy_OperatorAssign uintptr
	addr_QSizePolicy_toQVariant uintptr
	pg_QSizePolicy_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSizePolicy_new, "QSizePolicy_new")
	symaddr(&addr_QSizePolicy_new2, "QSizePolicy_new2")
	symaddr(&addr_QSizePolicy_new3, "QSizePolicy_new3")
	symaddr(&addr_QSizePolicy_new4, "QSizePolicy_new4")
	symaddr(&addr_QSizePolicy_horizontalPolicy, "QSizePolicy_horizontalPolicy")
	symaddr(&addr_QSizePolicy_verticalPolicy, "QSizePolicy_verticalPolicy")
	symaddr(&addr_QSizePolicy_controlType, "QSizePolicy_controlType")
	symaddr(&addr_QSizePolicy_setHorizontalPolicy, "QSizePolicy_setHorizontalPolicy")
	symaddr(&addr_QSizePolicy_setVerticalPolicy, "QSizePolicy_setVerticalPolicy")
	symaddr(&addr_QSizePolicy_setControlType, "QSizePolicy_setControlType")
	symaddr(&addr_QSizePolicy_expandingDirections, "QSizePolicy_expandingDirections")
	symaddr(&addr_QSizePolicy_setHeightForWidth, "QSizePolicy_setHeightForWidth")
	symaddr(&addr_QSizePolicy_hasHeightForWidth, "QSizePolicy_hasHeightForWidth")
	symaddr(&addr_QSizePolicy_setWidthForHeight, "QSizePolicy_setWidthForHeight")
	symaddr(&addr_QSizePolicy_hasWidthForHeight, "QSizePolicy_hasWidthForHeight")
	symaddr(&addr_QSizePolicy_OperatorEqual, "QSizePolicy_OperatorEqual")
	symaddr(&addr_QSizePolicy_OperatorNotEqual, "QSizePolicy_OperatorNotEqual")
	symaddr(&addr_QSizePolicy_ToQVariant, "QSizePolicy_ToQVariant")
	symaddr(&addr_QSizePolicy_horizontalStretch, "QSizePolicy_horizontalStretch")
	symaddr(&addr_QSizePolicy_verticalStretch, "QSizePolicy_verticalStretch")
	symaddr(&addr_QSizePolicy_setHorizontalStretch, "QSizePolicy_setHorizontalStretch")
	symaddr(&addr_QSizePolicy_setVerticalStretch, "QSizePolicy_setVerticalStretch")
	symaddr(&addr_QSizePolicy_retainSizeWhenHidden, "QSizePolicy_retainSizeWhenHidden")
	symaddr(&addr_QSizePolicy_setRetainSizeWhenHidden, "QSizePolicy_setRetainSizeWhenHidden")
	symaddr(&addr_QSizePolicy_transpose, "QSizePolicy_transpose")
	symaddr(&addr_QSizePolicy_transposed, "QSizePolicy_transposed")
	symaddr(&addr_QSizePolicy_OperatorAssign, "QSizePolicy_OperatorAssign")
	symaddr(&addr_QSizePolicy_toQVariant, "QSizePolicy_toQVariant")
	reg(&pg_QSizePolicy_delete, "QSizePolicy_delete")
}

type QSizePolicy struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSizePolicy) UnsafeGoID() int64     { return this.vid }
func (this *QSizePolicy) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSizePolicy(h uintptr) *QSizePolicy {
	if h == 0 {
		return nil
	}
	return &QSizePolicy{h: h}
}

func newownedQSizePolicy(h uintptr) *QSizePolicy {
	o := bindQSizePolicy(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSizePolicy).Delete)
	}
	return o
}

func (this *QSizePolicy) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSizePolicy) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSizePolicy(h unsafe.Pointer) *QSizePolicy { return bindQSizePolicy(uintptr(h)) }
func UnsafeNewQSizePolicyOwned(h unsafe.Pointer) *QSizePolicy { return newownedQSizePolicy(uintptr(h)) }

func (this *QSizePolicy) UnsafeRelease() { this.h = 0 }

func NewQSizePolicy() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_new)
	__o := newownedQSizePolicy(__r)
	return __o
}

func NewQSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_new2, uintptr(horizontal), uintptr(vertical))
	__o := newownedQSizePolicy(__r)
	return __o
}

func NewQSizePolicy3(arg1 *QSizePolicy) *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQSizePolicy(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQSizePolicy4(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy, type_ QSizePolicy__ControlType) *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_new4, uintptr(horizontal), uintptr(vertical), uintptr(type_))
	__o := newownedQSizePolicy(__r)
	return __o
}

func (this *QSizePolicy) HorizontalPolicy() QSizePolicy__Policy {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_horizontalPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__Policy(__r)
}

func (this *QSizePolicy) VerticalPolicy() QSizePolicy__Policy {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_verticalPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__Policy(__r)
}

func (this *QSizePolicy) ControlType() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_controlType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QSizePolicy) SetHorizontalPolicy(d QSizePolicy__Policy) {
	purego.SyscallN(addr_QSizePolicy_setHorizontalPolicy, this.h, uintptr(d))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) SetVerticalPolicy(d QSizePolicy__Policy) {
	purego.SyscallN(addr_QSizePolicy_setVerticalPolicy, this.h, uintptr(d))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) SetControlType(type_ QSizePolicy__ControlType) {
	purego.SyscallN(addr_QSizePolicy_setControlType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) ExpandingDirections() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_expandingDirections, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QSizePolicy) SetHeightForWidth(b bool) {
	purego.SyscallN(addr_QSizePolicy_setHeightForWidth, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) HasHeightForWidth() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_hasHeightForWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizePolicy) SetWidthForHeight(b bool) {
	purego.SyscallN(addr_QSizePolicy_setWidthForHeight, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) HasWidthForHeight() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_hasWidthForHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizePolicy) OperatorEqual(s *QSizePolicy) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_OperatorEqual, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QSizePolicy) OperatorNotEqual(s *QSizePolicy) bool {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_OperatorNotEqual, this.h, uintptr(s.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return __r != 0
}

func (this *QSizePolicy) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSizePolicy) HorizontalStretch() int {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_horizontalStretch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSizePolicy) VerticalStretch() int {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_verticalStretch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	purego.SyscallN(addr_QSizePolicy_setHorizontalStretch, this.h, uintptr(stretchFactor))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	purego.SyscallN(addr_QSizePolicy_setVerticalStretch, this.h, uintptr(stretchFactor))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) RetainSizeWhenHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QSizePolicy_retainSizeWhenHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	purego.SyscallN(addr_QSizePolicy_setRetainSizeWhenHidden, this.h, uintptr(b2i(retainSize)))
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) Transpose() {
	purego.SyscallN(addr_QSizePolicy_transpose, this.h)
	runtime.KeepAlive(this)
}

func (this *QSizePolicy) Transposed() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_transposed, this.h)
	runtime.KeepAlive(this)
	return newownedQSizePolicy(__r)
}

func (this *QSizePolicy) OperatorAssign(arg1 *QSizePolicy) *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQSizePolicy(__r)
}

func (this *QSizePolicy) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSizePolicy_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSizePolicy) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSizePolicy_delete(this.h)
	}
	this.UnsafeRelease()
}
