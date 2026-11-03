package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFileIconProvider_base func(uintptr) uintptr
	addr_QFileIconProvider_new uintptr
	addr_QFileIconProvider_icon_default uintptr
	addr_QFileIconProvider_type_default uintptr
	addr_QFileIconProvider_setOptions_default uintptr
	addr_QFileIconProvider_options_default uintptr
	pg_QFileIconProvider__set_override func(uintptr, int32)
	pg_QFileIconProvider_delete func(uintptr)
)

func init() {
	reg(&pg_QFileIconProvider_base, "QFileIconProvider_base")
	symaddr(&addr_QFileIconProvider_new, "QFileIconProvider_new")
	symaddr(&addr_QFileIconProvider_icon_default, "QFileIconProvider_icon_default")
	symaddr(&addr_QFileIconProvider_type_default, "QFileIconProvider_type_default")
	symaddr(&addr_QFileIconProvider_setOptions_default, "QFileIconProvider_setOptions_default")
	symaddr(&addr_QFileIconProvider_options_default, "QFileIconProvider_options_default")
	reg(&pg_QFileIconProvider__set_override, "QFileIconProvider__set_override")
	reg(&pg_QFileIconProvider_delete, "QFileIconProvider_delete")
}

type QFileIconProvider struct {
	h uintptr
	*QAbstractFileIconProvider
}

func bindQFileIconProvider(h uintptr) *QFileIconProvider {
	if h == 0 {
		return nil
	}
	return &QFileIconProvider{h: h, QAbstractFileIconProvider: bindQAbstractFileIconProvider(pg_QFileIconProvider_base(h))}
}

func newownedQFileIconProvider(h uintptr) *QFileIconProvider {
	o := bindQFileIconProvider(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFileIconProvider).Delete)
	}
	return o
}

func (this *QFileIconProvider) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileIconProvider) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileIconProvider(h unsafe.Pointer) *QFileIconProvider { return bindQFileIconProvider(uintptr(h)) }
func UnsafeNewQFileIconProviderOwned(h unsafe.Pointer) *QFileIconProvider { return newownedQFileIconProvider(uintptr(h)) }

func (this *QFileIconProvider) UnsafeRelease() {
	this.h = 0
	if this.QAbstractFileIconProvider != nil {
		this.QAbstractFileIconProvider.UnsafeRelease()
	}
}

func NewQFileIconProvider() *QFileIconProvider {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileIconProvider_new, uintptr(__goid))
	__o := newownedQFileIconProvider(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileIconProvider_virtIdx, QFileIconProvider_enable)
	return __o
}

func (this *QFileIconProvider) IconDefault(type_ QAbstractFileIconProvider__IconType) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QFileIconProvider_icon_default, this.h, uintptr(type_))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QFileIconProvider) TypeDefault(arg1 *QFileInfo) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileIconProvider_type_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return pgStr(__r, __n)
}

func (this *QFileIconProvider) SetOptionsDefault(arg1 QAbstractFileIconProvider__Option) {
	purego.SyscallN(addr_QFileIconProvider_setOptions_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFileIconProvider) OptionsDefault() QAbstractFileIconProvider__Option {
	__r0, _, _ := purego.SyscallN(addr_QFileIconProvider_options_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractFileIconProvider__Option(__r)
}

func QFileIconProvider_virtIdx(slot int64) int32 {
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

func QFileIconProvider_enable(h uintptr, idx int32) { pg_QFileIconProvider__set_override(h, idx) }

func (this *QFileIconProvider) OnIcon(fn func(QAbstractFileIconProvider__IconType) *QIcon) {
	registerVirt(this.UnsafeGoID(), -2594270449535773968, func(type_ int32) uintptr { return uintptr(fn(QAbstractFileIconProvider__IconType(type_)).UnsafePointer()) }) // slot "icon"
}

func (this *QFileIconProvider) OnIcon2(fn func(*QFileInfo) *QIcon) {
	registerVirt(this.UnsafeGoID(), -909105390709809242, func(info uintptr) uintptr { return uintptr(fn(UnsafeNewQFileInfo(unsafe.Pointer(info))).UnsafePointer()) }) // slot "icon2"
}

func (this *QFileIconProvider) OnType(fn func(*QFileInfo) string) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func(arg1 uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQFileInfo(unsafe.Pointer(arg1)))
		return pgCString(__s), int64(len(__s))
	}) // slot "type"
}

func (this *QFileIconProvider) OnSetOptions(fn func(QAbstractFileIconProvider__Option)) {
	registerVirt(this.UnsafeGoID(), 5783776084322042327, func(arg1 int32) { fn(QAbstractFileIconProvider__Option(arg1)) }) // slot "setOptions"
}

func (this *QFileIconProvider) OnOptions(fn func() QAbstractFileIconProvider__Option) {
	registerVirt(this.UnsafeGoID(), -3871419584579008699, func() int32 { return int32(fn()) }) // slot "options"
}

func (this *QFileIconProvider) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileIconProvider_delete(this.h)
	}
	this.UnsafeRelease()
}
