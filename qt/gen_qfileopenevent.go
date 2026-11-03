package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFileOpenEvent_base func(uintptr) uintptr
	addr_QFileOpenEvent_new uintptr
	addr_QFileOpenEvent_new2 uintptr
	addr_QFileOpenEvent_file uintptr
	addr_QFileOpenEvent_url uintptr
	addr_QFileOpenEvent_openFile uintptr
	addr_QFileOpenEvent_registerEventType uintptr
	addr_QFileOpenEvent_registerEventType2 uintptr
	pg_QFileOpenEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QFileOpenEvent_base, "QFileOpenEvent_base")
	symaddr(&addr_QFileOpenEvent_new, "QFileOpenEvent_new")
	symaddr(&addr_QFileOpenEvent_new2, "QFileOpenEvent_new2")
	symaddr(&addr_QFileOpenEvent_file, "QFileOpenEvent_file")
	symaddr(&addr_QFileOpenEvent_url, "QFileOpenEvent_url")
	symaddr(&addr_QFileOpenEvent_openFile, "QFileOpenEvent_openFile")
	symaddr(&addr_QFileOpenEvent_registerEventType, "QFileOpenEvent_registerEventType")
	symaddr(&addr_QFileOpenEvent_registerEventType2, "QFileOpenEvent_registerEventType2")
	reg(&pg_QFileOpenEvent_delete, "QFileOpenEvent_delete")
}

type QFileOpenEvent struct {
	h uintptr
	*QEvent
}

func bindQFileOpenEvent(h uintptr) *QFileOpenEvent {
	if h == 0 {
		return nil
	}
	return &QFileOpenEvent{h: h, QEvent: bindQEvent(pg_QFileOpenEvent_base(h))}
}

func newownedQFileOpenEvent(h uintptr) *QFileOpenEvent {
	o := bindQFileOpenEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFileOpenEvent).Delete)
	}
	return o
}

func (this *QFileOpenEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileOpenEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileOpenEvent(h unsafe.Pointer) *QFileOpenEvent { return bindQFileOpenEvent(uintptr(h)) }
func UnsafeNewQFileOpenEventOwned(h unsafe.Pointer) *QFileOpenEvent { return newownedQFileOpenEvent(uintptr(h)) }

func (this *QFileOpenEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQFileOpenEvent(file string) *QFileOpenEvent {
	var file_p uintptr
	if len(file) > 0 {
		file_p = uintptr(unsafe.Pointer(unsafe.StringData(file)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileOpenEvent_new, file_p, uintptr(len(file)))
	__o := newownedQFileOpenEvent(__r)
	runtime.KeepAlive(file)
	return __o
}

func NewQFileOpenEvent2(url *QUrl) *QFileOpenEvent {
	__r, _, _ := purego.SyscallN(addr_QFileOpenEvent_new2, uintptr(url.UnsafePointer()))
	__o := newownedQFileOpenEvent(__r)
	runtime.KeepAlive(url)
	return __o
}

func (this *QFileOpenEvent) File() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileOpenEvent_file, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileOpenEvent) Url() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileOpenEvent_url, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QFileOpenEvent) OpenFile(file *QFile, flags QIODeviceBase__OpenModeFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileOpenEvent_openFile, this.h, uintptr(file.UnsafePointer()), uintptr(flags))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file)
	return __r != 0
}

func QFileOpenEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QFileOpenEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QFileOpenEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QFileOpenEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QFileOpenEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileOpenEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
