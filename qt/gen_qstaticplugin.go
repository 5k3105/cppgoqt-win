package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStaticPlugin_new uintptr
	addr_QStaticPlugin_metaData uintptr
	pg_QStaticPlugin_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStaticPlugin_new, "QStaticPlugin_new")
	symaddr(&addr_QStaticPlugin_metaData, "QStaticPlugin_metaData")
	reg(&pg_QStaticPlugin_delete, "QStaticPlugin_delete")
}

type QStaticPlugin struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStaticPlugin) UnsafeGoID() int64     { return this.vid }
func (this *QStaticPlugin) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStaticPlugin(h uintptr) *QStaticPlugin {
	if h == 0 {
		return nil
	}
	return &QStaticPlugin{h: h}
}

func newownedQStaticPlugin(h uintptr) *QStaticPlugin {
	o := bindQStaticPlugin(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStaticPlugin).Delete)
	}
	return o
}

func (this *QStaticPlugin) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStaticPlugin) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStaticPlugin(h unsafe.Pointer) *QStaticPlugin { return bindQStaticPlugin(uintptr(h)) }
func UnsafeNewQStaticPluginOwned(h unsafe.Pointer) *QStaticPlugin { return newownedQStaticPlugin(uintptr(h)) }

func (this *QStaticPlugin) UnsafeRelease() { this.h = 0 }

func NewQStaticPlugin(arg1 *QStaticPlugin) *QStaticPlugin {
	__r, _, _ := purego.SyscallN(addr_QStaticPlugin_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQStaticPlugin(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QStaticPlugin) MetaData() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QStaticPlugin_metaData, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QStaticPlugin) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStaticPlugin_delete(this.h)
	}
	this.UnsafeRelease()
}
