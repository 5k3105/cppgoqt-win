package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOption_new uintptr
	addr_QStyleOption_new2 uintptr
	addr_QStyleOption_new3 uintptr
	addr_QStyleOption_new4 uintptr
	addr_QStyleOption_initFrom uintptr
	addr_QStyleOption_OperatorAssign uintptr
	addr_QStyleOption_version_get uintptr
	addr_QStyleOption_version_set uintptr
	addr_QStyleOption_type_get uintptr
	addr_QStyleOption_type_set uintptr
	addr_QStyleOption_state_get uintptr
	addr_QStyleOption_state_set uintptr
	addr_QStyleOption_direction_get uintptr
	addr_QStyleOption_direction_set uintptr
	addr_QStyleOption_rect_get uintptr
	addr_QStyleOption_rect_set uintptr
	addr_QStyleOption_fontMetrics_get uintptr
	addr_QStyleOption_fontMetrics_set uintptr
	addr_QStyleOption_palette_get uintptr
	addr_QStyleOption_palette_set uintptr
	addr_QStyleOption_styleObject_get uintptr
	addr_QStyleOption_styleObject_set uintptr
	pg_QStyleOption_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOption_new, "QStyleOption_new")
	symaddr(&addr_QStyleOption_new2, "QStyleOption_new2")
	symaddr(&addr_QStyleOption_new3, "QStyleOption_new3")
	symaddr(&addr_QStyleOption_new4, "QStyleOption_new4")
	symaddr(&addr_QStyleOption_initFrom, "QStyleOption_initFrom")
	symaddr(&addr_QStyleOption_OperatorAssign, "QStyleOption_OperatorAssign")
	symaddr(&addr_QStyleOption_version_get, "QStyleOption_version_get")
	symaddr(&addr_QStyleOption_version_set, "QStyleOption_version_set")
	symaddr(&addr_QStyleOption_type_get, "QStyleOption_type_get")
	symaddr(&addr_QStyleOption_type_set, "QStyleOption_type_set")
	symaddr(&addr_QStyleOption_state_get, "QStyleOption_state_get")
	symaddr(&addr_QStyleOption_state_set, "QStyleOption_state_set")
	symaddr(&addr_QStyleOption_direction_get, "QStyleOption_direction_get")
	symaddr(&addr_QStyleOption_direction_set, "QStyleOption_direction_set")
	symaddr(&addr_QStyleOption_rect_get, "QStyleOption_rect_get")
	symaddr(&addr_QStyleOption_rect_set, "QStyleOption_rect_set")
	symaddr(&addr_QStyleOption_fontMetrics_get, "QStyleOption_fontMetrics_get")
	symaddr(&addr_QStyleOption_fontMetrics_set, "QStyleOption_fontMetrics_set")
	symaddr(&addr_QStyleOption_palette_get, "QStyleOption_palette_get")
	symaddr(&addr_QStyleOption_palette_set, "QStyleOption_palette_set")
	symaddr(&addr_QStyleOption_styleObject_get, "QStyleOption_styleObject_get")
	symaddr(&addr_QStyleOption_styleObject_set, "QStyleOption_styleObject_set")
	reg(&pg_QStyleOption_delete, "QStyleOption_delete")
}

type QStyleOption struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStyleOption) UnsafeGoID() int64     { return this.vid }
func (this *QStyleOption) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStyleOption(h uintptr) *QStyleOption {
	if h == 0 {
		return nil
	}
	return &QStyleOption{h: h}
}

func newownedQStyleOption(h uintptr) *QStyleOption {
	o := bindQStyleOption(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOption).Delete)
	}
	return o
}

func (this *QStyleOption) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOption) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOption(h unsafe.Pointer) *QStyleOption { return bindQStyleOption(uintptr(h)) }
func UnsafeNewQStyleOptionOwned(h unsafe.Pointer) *QStyleOption { return newownedQStyleOption(uintptr(h)) }

func (this *QStyleOption) UnsafeRelease() { this.h = 0 }

func NewQStyleOption() *QStyleOption {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_new)
	__o := newownedQStyleOption(__r)
	return __o
}

func NewQStyleOption2(other *QStyleOption) *QStyleOption {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOption(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQStyleOption3(version int) *QStyleOption {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_new3, uintptr(version))
	__o := newownedQStyleOption(__r)
	return __o
}

func NewQStyleOption4(version int, type_ int) *QStyleOption {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_new4, uintptr(version), uintptr(type_))
	__o := newownedQStyleOption(__r)
	return __o
}

func (this *QStyleOption) InitFrom(w *QWidget) {
	purego.SyscallN(addr_QStyleOption_initFrom, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QStyleOption) OperatorAssign(other *QStyleOption) *QStyleOption {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQStyleOption(__r)
}

func (this *QStyleOption) Version() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOption_version_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOption) SetVersion(version int) {
	purego.SyscallN(addr_QStyleOption_version_set, this.h, uintptr(version))
	runtime.KeepAlive(this)
}

func (this *QStyleOption) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleOption_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleOption) SetType(type_ int) {
	purego.SyscallN(addr_QStyleOption_type_set, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QStyleOption) State() QStyle__StateFlag {
	__r0, _, _ := purego.SyscallN(addr_QStyleOption_state_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyle__StateFlag(__r)
}

func (this *QStyleOption) SetState(state QStyle__StateFlag) {
	purego.SyscallN(addr_QStyleOption_state_set, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QStyleOption) Direction() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QStyleOption_direction_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QStyleOption) SetDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QStyleOption_direction_set, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QStyleOption) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_rect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStyleOption) SetRect(rect *QRect) {
	purego.SyscallN(addr_QStyleOption_rect_set, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QStyleOption) FontMetrics() *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_fontMetrics_get, this.h)
	runtime.KeepAlive(this)
	return newownedQFontMetrics(__r)
}

func (this *QStyleOption) SetFontMetrics(fontMetrics *QFontMetrics) {
	purego.SyscallN(addr_QStyleOption_fontMetrics_set, this.h, uintptr(fontMetrics.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontMetrics)
}

func (this *QStyleOption) Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_palette_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QStyleOption) SetPalette(palette *QPalette) {
	purego.SyscallN(addr_QStyleOption_palette_set, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QStyleOption) StyleObject() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStyleOption_styleObject_get, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStyleOption) SetStyleObject(styleObject *QObject) {
	purego.SyscallN(addr_QStyleOption_styleObject_set, this.h, uintptr(styleObject.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(styleObject)
}

func (this *QStyleOption) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOption_delete(this.h)
	}
	this.UnsafeRelease()
}
