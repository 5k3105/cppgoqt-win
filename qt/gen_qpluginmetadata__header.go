package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPluginMetaData__Header_new uintptr
	addr_QPluginMetaData__Header_version_get uintptr
	addr_QPluginMetaData__Header_version_set uintptr
	addr_QPluginMetaData__Header_qt_major_version_get uintptr
	addr_QPluginMetaData__Header_qt_major_version_set uintptr
	addr_QPluginMetaData__Header_qt_minor_version_get uintptr
	addr_QPluginMetaData__Header_qt_minor_version_set uintptr
	addr_QPluginMetaData__Header_plugin_arch_requirements_get uintptr
	addr_QPluginMetaData__Header_plugin_arch_requirements_set uintptr
	pg_QPluginMetaData__Header_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPluginMetaData__Header_new, "QPluginMetaData__Header_new")
	symaddr(&addr_QPluginMetaData__Header_version_get, "QPluginMetaData__Header_version_get")
	symaddr(&addr_QPluginMetaData__Header_version_set, "QPluginMetaData__Header_version_set")
	symaddr(&addr_QPluginMetaData__Header_qt_major_version_get, "QPluginMetaData__Header_qt_major_version_get")
	symaddr(&addr_QPluginMetaData__Header_qt_major_version_set, "QPluginMetaData__Header_qt_major_version_set")
	symaddr(&addr_QPluginMetaData__Header_qt_minor_version_get, "QPluginMetaData__Header_qt_minor_version_get")
	symaddr(&addr_QPluginMetaData__Header_qt_minor_version_set, "QPluginMetaData__Header_qt_minor_version_set")
	symaddr(&addr_QPluginMetaData__Header_plugin_arch_requirements_get, "QPluginMetaData__Header_plugin_arch_requirements_get")
	symaddr(&addr_QPluginMetaData__Header_plugin_arch_requirements_set, "QPluginMetaData__Header_plugin_arch_requirements_set")
	reg(&pg_QPluginMetaData__Header_delete, "QPluginMetaData__Header_delete")
}

type QPluginMetaData__Header struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPluginMetaData__Header) UnsafeGoID() int64     { return this.vid }
func (this *QPluginMetaData__Header) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPluginMetaData__Header(h uintptr) *QPluginMetaData__Header {
	if h == 0 {
		return nil
	}
	return &QPluginMetaData__Header{h: h}
}

func newownedQPluginMetaData__Header(h uintptr) *QPluginMetaData__Header {
	o := bindQPluginMetaData__Header(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPluginMetaData__Header).Delete)
	}
	return o
}

func (this *QPluginMetaData__Header) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPluginMetaData__Header) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPluginMetaData__Header(h unsafe.Pointer) *QPluginMetaData__Header { return bindQPluginMetaData__Header(uintptr(h)) }
func UnsafeNewQPluginMetaData__HeaderOwned(h unsafe.Pointer) *QPluginMetaData__Header { return newownedQPluginMetaData__Header(uintptr(h)) }

func (this *QPluginMetaData__Header) UnsafeRelease() { this.h = 0 }

func NewQPluginMetaData__Header(arg1 *QPluginMetaData__Header) *QPluginMetaData__Header {
	__r, _, _ := purego.SyscallN(addr_QPluginMetaData__Header_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQPluginMetaData__Header(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPluginMetaData__Header) Version() byte {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__Header_version_get, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPluginMetaData__Header) SetVersion(version byte) {
	purego.SyscallN(addr_QPluginMetaData__Header_version_set, this.h, uintptr(version))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__Header) Qt_major_version() byte {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__Header_qt_major_version_get, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPluginMetaData__Header) SetQt_major_version(qt_major_version byte) {
	purego.SyscallN(addr_QPluginMetaData__Header_qt_major_version_set, this.h, uintptr(qt_major_version))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__Header) Qt_minor_version() byte {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__Header_qt_minor_version_get, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPluginMetaData__Header) SetQt_minor_version(qt_minor_version byte) {
	purego.SyscallN(addr_QPluginMetaData__Header_qt_minor_version_set, this.h, uintptr(qt_minor_version))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__Header) Plugin_arch_requirements() byte {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData__Header_plugin_arch_requirements_get, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QPluginMetaData__Header) SetPlugin_arch_requirements(plugin_arch_requirements byte) {
	purego.SyscallN(addr_QPluginMetaData__Header_plugin_arch_requirements_set, this.h, uintptr(plugin_arch_requirements))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData__Header) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPluginMetaData__Header_delete(this.h)
	}
	this.UnsafeRelease()
}
