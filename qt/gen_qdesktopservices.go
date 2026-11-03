package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDesktopServices_openUrl uintptr
	addr_QDesktopServices_setUrlHandler uintptr
	addr_QDesktopServices_unsetUrlHandler uintptr
	pg_QDesktopServices_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDesktopServices_openUrl, "QDesktopServices_openUrl")
	symaddr(&addr_QDesktopServices_setUrlHandler, "QDesktopServices_setUrlHandler")
	symaddr(&addr_QDesktopServices_unsetUrlHandler, "QDesktopServices_unsetUrlHandler")
	reg(&pg_QDesktopServices_delete, "QDesktopServices_delete")
}

type QDesktopServices struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDesktopServices) UnsafeGoID() int64     { return this.vid }
func (this *QDesktopServices) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDesktopServices(h uintptr) *QDesktopServices {
	if h == 0 {
		return nil
	}
	return &QDesktopServices{h: h}
}

func newownedQDesktopServices(h uintptr) *QDesktopServices {
	o := bindQDesktopServices(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDesktopServices).Delete)
	}
	return o
}

func (this *QDesktopServices) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDesktopServices) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDesktopServices(h unsafe.Pointer) *QDesktopServices { return bindQDesktopServices(uintptr(h)) }
func UnsafeNewQDesktopServicesOwned(h unsafe.Pointer) *QDesktopServices { return newownedQDesktopServices(uintptr(h)) }

func (this *QDesktopServices) UnsafeRelease() { this.h = 0 }

func QDesktopServices_OpenUrl(url *QUrl) bool {
	__r0, _, _ := purego.SyscallN(addr_QDesktopServices_openUrl, uintptr(url.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(url)
	return __r != 0
}

func QDesktopServices_SetUrlHandler(scheme string, receiver *QObject, method string) {
	var scheme_p uintptr
	if len(scheme) > 0 {
		scheme_p = uintptr(unsafe.Pointer(unsafe.StringData(scheme)))
	}
	method_b := append([]byte(method), 0)
	purego.SyscallN(addr_QDesktopServices_setUrlHandler, scheme_p, uintptr(len(scheme)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&method_b[0])))
	runtime.KeepAlive(scheme)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method_b)
}

func QDesktopServices_UnsetUrlHandler(scheme string) {
	var scheme_p uintptr
	if len(scheme) > 0 {
		scheme_p = uintptr(unsafe.Pointer(unsafe.StringData(scheme)))
	}
	purego.SyscallN(addr_QDesktopServices_unsetUrlHandler, scheme_p, uintptr(len(scheme)))
	runtime.KeepAlive(scheme)
}

func (this *QDesktopServices) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDesktopServices_delete(this.h)
	}
	this.UnsafeRelease()
}
