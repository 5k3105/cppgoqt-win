package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractFileIconProvider_new uintptr
	addr_QAbstractFileIconProvider_icon uintptr
	addr_QAbstractFileIconProvider_icon2 uintptr
	addr_QAbstractFileIconProvider_type uintptr
	addr_QAbstractFileIconProvider_setOptions uintptr
	addr_QAbstractFileIconProvider_options uintptr
	addr_QAbstractFileIconProvider_icon_default uintptr
	addr_QAbstractFileIconProvider_icon_default2 uintptr
	addr_QAbstractFileIconProvider_type_default uintptr
	addr_QAbstractFileIconProvider_setOptions_default uintptr
	addr_QAbstractFileIconProvider_options_default uintptr
	pg_QAbstractFileIconProvider__set_override func(uintptr, int32)
	pg_QAbstractFileIconProvider_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractFileIconProvider_new, "QAbstractFileIconProvider_new")
	symaddr(&addr_QAbstractFileIconProvider_icon, "QAbstractFileIconProvider_icon")
	symaddr(&addr_QAbstractFileIconProvider_icon2, "QAbstractFileIconProvider_icon2")
	symaddr(&addr_QAbstractFileIconProvider_type, "QAbstractFileIconProvider_type")
	symaddr(&addr_QAbstractFileIconProvider_setOptions, "QAbstractFileIconProvider_setOptions")
	symaddr(&addr_QAbstractFileIconProvider_options, "QAbstractFileIconProvider_options")
	symaddr(&addr_QAbstractFileIconProvider_icon_default, "QAbstractFileIconProvider_icon_default")
	symaddr(&addr_QAbstractFileIconProvider_icon_default2, "QAbstractFileIconProvider_icon_default2")
	symaddr(&addr_QAbstractFileIconProvider_type_default, "QAbstractFileIconProvider_type_default")
	symaddr(&addr_QAbstractFileIconProvider_setOptions_default, "QAbstractFileIconProvider_setOptions_default")
	symaddr(&addr_QAbstractFileIconProvider_options_default, "QAbstractFileIconProvider_options_default")
	reg(&pg_QAbstractFileIconProvider__set_override, "QAbstractFileIconProvider__set_override")
	reg(&pg_QAbstractFileIconProvider_delete, "QAbstractFileIconProvider_delete")
}

type QAbstractFileIconProvider struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractFileIconProvider) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractFileIconProvider) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractFileIconProvider(h uintptr) *QAbstractFileIconProvider {
	if h == 0 {
		return nil
	}
	return &QAbstractFileIconProvider{h: h}
}

func newownedQAbstractFileIconProvider(h uintptr) *QAbstractFileIconProvider {
	o := bindQAbstractFileIconProvider(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractFileIconProvider).Delete)
	}
	return o
}

func (this *QAbstractFileIconProvider) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractFileIconProvider) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractFileIconProvider(h unsafe.Pointer) *QAbstractFileIconProvider { return bindQAbstractFileIconProvider(uintptr(h)) }
func UnsafeNewQAbstractFileIconProviderOwned(h unsafe.Pointer) *QAbstractFileIconProvider { return newownedQAbstractFileIconProvider(uintptr(h)) }

func (this *QAbstractFileIconProvider) UnsafeRelease() { this.h = 0 }

func NewQAbstractFileIconProvider() *QAbstractFileIconProvider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_new, uintptr(__goid))
	__o := newownedQAbstractFileIconProvider(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractFileIconProvider_virtIdx, QAbstractFileIconProvider_enable)
	return __o
}

func (this *QAbstractFileIconProvider) Icon(arg1 QAbstractFileIconProvider__IconType) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_icon, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QAbstractFileIconProvider) IconWithQFileInfo(arg1 *QFileInfo) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_icon2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQIcon(__r)
}

func (this *QAbstractFileIconProvider) Type(arg1 *QFileInfo) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_type, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func (this *QAbstractFileIconProvider) SetOptions(arg1 QAbstractFileIconProvider__Option) {
	purego.SyscallN(addr_QAbstractFileIconProvider_setOptions, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractFileIconProvider) Options() QAbstractFileIconProvider__Option {
	__r0, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractFileIconProvider__Option(__r)
}

func (this *QAbstractFileIconProvider) IconDefault(arg1 QAbstractFileIconProvider__IconType) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_icon_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QAbstractFileIconProvider) IconWithQFileInfoDefault(arg1 *QFileInfo) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_icon_default2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQIcon(__r)
}

func (this *QAbstractFileIconProvider) TypeDefault(arg1 *QFileInfo) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_type_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func (this *QAbstractFileIconProvider) SetOptionsDefault(arg1 QAbstractFileIconProvider__Option) {
	purego.SyscallN(addr_QAbstractFileIconProvider_setOptions_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractFileIconProvider) OptionsDefault() QAbstractFileIconProvider__Option {
	__r0, _, _ := purego.SyscallN(addr_QAbstractFileIconProvider_options_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractFileIconProvider__Option(__r)
}

func QAbstractFileIconProvider_virtIdx(slot int64) int32 {
	switch slot {
	case -2594270449535773968:
		return 0
	case -909105390709809242:
		return 1
	case -6371403872082420691:
		return 2
	case 5783776084322042327:
		return 3
	case -3871419584579008699:
		return 4
	}
	return -1
}

func QAbstractFileIconProvider_enable(h uintptr, idx int32) { pg_QAbstractFileIconProvider__set_override(h, idx) }

func (this *QAbstractFileIconProvider) OnIcon(fn func(QAbstractFileIconProvider__IconType) *QIcon) {
	registerVirt(this.UnsafeGoID(), -2594270449535773968, func(arg1 int32) uintptr { return uintptr(fn(QAbstractFileIconProvider__IconType(arg1)).UnsafePointer()) }) // slot "icon"
}

func (this *QAbstractFileIconProvider) OnIcon2(fn func(*QFileInfo) *QIcon) {
	registerVirt(this.UnsafeGoID(), -909105390709809242, func(arg1 uintptr) uintptr { return uintptr(fn(UnsafeNewQFileInfo(unsafe.Pointer(arg1))).UnsafePointer()) }) // slot "icon2"
}

func (this *QAbstractFileIconProvider) OnType(fn func(*QFileInfo) string) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func(arg1 uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQFileInfo(unsafe.Pointer(arg1)))
		return pgCString(__s), int64(len(__s))
	}) // slot "type"
}

func (this *QAbstractFileIconProvider) OnSetOptions(fn func(QAbstractFileIconProvider__Option)) {
	registerVirt(this.UnsafeGoID(), 5783776084322042327, func(arg1 int32) { fn(QAbstractFileIconProvider__Option(arg1)) }) // slot "setOptions"
}

func (this *QAbstractFileIconProvider) OnOptions(fn func() QAbstractFileIconProvider__Option) {
	registerVirt(this.UnsafeGoID(), -3871419584579008699, func() int32 { return int32(fn()) }) // slot "options"
}

func (this *QAbstractFileIconProvider) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractFileIconProvider_delete(this.h)
	}
	this.UnsafeRelease()
}
