package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessible_installActivationObserver uintptr
	addr_QAccessible_removeActivationObserver uintptr
	addr_QAccessible_queryAccessibleInterface uintptr
	addr_QAccessible_uniqueId uintptr
	addr_QAccessible_accessibleInterface uintptr
	addr_QAccessible_registerAccessibleInterface uintptr
	addr_QAccessible_deleteAccessibleInterface uintptr
	addr_QAccessible_updateAccessibility uintptr
	addr_QAccessible_isActive uintptr
	addr_QAccessible_setActive uintptr
	addr_QAccessible_setRootObject uintptr
	addr_QAccessible_cleanup uintptr
	addr_QAccessible_qAccessibleTextBoundaryHelper uintptr
	pg_QAccessible_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessible_installActivationObserver, "QAccessible_installActivationObserver")
	symaddr(&addr_QAccessible_removeActivationObserver, "QAccessible_removeActivationObserver")
	symaddr(&addr_QAccessible_queryAccessibleInterface, "QAccessible_queryAccessibleInterface")
	symaddr(&addr_QAccessible_uniqueId, "QAccessible_uniqueId")
	symaddr(&addr_QAccessible_accessibleInterface, "QAccessible_accessibleInterface")
	symaddr(&addr_QAccessible_registerAccessibleInterface, "QAccessible_registerAccessibleInterface")
	symaddr(&addr_QAccessible_deleteAccessibleInterface, "QAccessible_deleteAccessibleInterface")
	symaddr(&addr_QAccessible_updateAccessibility, "QAccessible_updateAccessibility")
	symaddr(&addr_QAccessible_isActive, "QAccessible_isActive")
	symaddr(&addr_QAccessible_setActive, "QAccessible_setActive")
	symaddr(&addr_QAccessible_setRootObject, "QAccessible_setRootObject")
	symaddr(&addr_QAccessible_cleanup, "QAccessible_cleanup")
	symaddr(&addr_QAccessible_qAccessibleTextBoundaryHelper, "QAccessible_qAccessibleTextBoundaryHelper")
	reg(&pg_QAccessible_delete, "QAccessible_delete")
}

type QAccessible struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessible) UnsafeGoID() int64     { return this.vid }
func (this *QAccessible) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessible(h uintptr) *QAccessible {
	if h == 0 {
		return nil
	}
	return &QAccessible{h: h}
}

func newownedQAccessible(h uintptr) *QAccessible {
	o := bindQAccessible(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessible).Delete)
	}
	return o
}

func (this *QAccessible) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessible) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessible(h unsafe.Pointer) *QAccessible { return bindQAccessible(uintptr(h)) }
func UnsafeNewQAccessibleOwned(h unsafe.Pointer) *QAccessible { return newownedQAccessible(uintptr(h)) }

func (this *QAccessible) UnsafeRelease() { this.h = 0 }

func QAccessible_InstallActivationObserver(arg1 *QAccessible__ActivationObserver) {
	purego.SyscallN(addr_QAccessible_installActivationObserver, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QAccessible_RemoveActivationObserver(arg1 *QAccessible__ActivationObserver) {
	purego.SyscallN(addr_QAccessible_removeActivationObserver, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QAccessible_QueryAccessibleInterface(arg1 *QObject) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessible_queryAccessibleInterface, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
	return bindQAccessibleInterface(__r)
}

func QAccessible_UniqueId(iface *QAccessibleInterface) uint {
	__r0, _, _ := purego.SyscallN(addr_QAccessible_uniqueId, uintptr(iface.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(iface)
	return uint(__r)
}

func QAccessible_AccessibleInterface(uniqueId uint) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessible_accessibleInterface, uintptr(uniqueId))
	return bindQAccessibleInterface(__r)
}

func QAccessible_RegisterAccessibleInterface(iface *QAccessibleInterface) uint {
	__r0, _, _ := purego.SyscallN(addr_QAccessible_registerAccessibleInterface, uintptr(iface.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(iface)
	return uint(__r)
}

func QAccessible_DeleteAccessibleInterface(uniqueId uint) {
	purego.SyscallN(addr_QAccessible_deleteAccessibleInterface, uintptr(uniqueId))
}

func QAccessible_UpdateAccessibility(event *QAccessibleEvent) {
	purego.SyscallN(addr_QAccessible_updateAccessibility, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(event)
}

func QAccessible_IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessible_isActive)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QAccessible_SetActive(active bool) {
	purego.SyscallN(addr_QAccessible_setActive, uintptr(b2i(active)))
}

func QAccessible_SetRootObject(object *QObject) {
	purego.SyscallN(addr_QAccessible_setRootObject, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(object)
}

func QAccessible_Cleanup() {
	purego.SyscallN(addr_QAccessible_cleanup)
}

func QAccessible_QAccessibleTextBoundaryHelper(cursor *QTextCursor, boundaryType QAccessible__TextBoundaryType) (int, int) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QAccessible_qAccessibleTextBoundaryHelper, uintptr(cursor.UnsafePointer()), uintptr(boundaryType), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(cursor)
	return int(__r), int(__b)
}

func (this *QAccessible) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessible_delete(this.h)
	}
	this.UnsafeRelease()
}
