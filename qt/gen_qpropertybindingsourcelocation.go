package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyBindingSourceLocation_new uintptr
	addr_QPropertyBindingSourceLocation_new2 uintptr
	addr_QPropertyBindingSourceLocation_fileName_get uintptr
	addr_QPropertyBindingSourceLocation_fileName_set uintptr
	addr_QPropertyBindingSourceLocation_functionName_get uintptr
	addr_QPropertyBindingSourceLocation_functionName_set uintptr
	addr_QPropertyBindingSourceLocation_line_get uintptr
	addr_QPropertyBindingSourceLocation_line_set uintptr
	addr_QPropertyBindingSourceLocation_column_get uintptr
	addr_QPropertyBindingSourceLocation_column_set uintptr
	pg_QPropertyBindingSourceLocation_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyBindingSourceLocation_new, "QPropertyBindingSourceLocation_new")
	symaddr(&addr_QPropertyBindingSourceLocation_new2, "QPropertyBindingSourceLocation_new2")
	symaddr(&addr_QPropertyBindingSourceLocation_fileName_get, "QPropertyBindingSourceLocation_fileName_get")
	symaddr(&addr_QPropertyBindingSourceLocation_fileName_set, "QPropertyBindingSourceLocation_fileName_set")
	symaddr(&addr_QPropertyBindingSourceLocation_functionName_get, "QPropertyBindingSourceLocation_functionName_get")
	symaddr(&addr_QPropertyBindingSourceLocation_functionName_set, "QPropertyBindingSourceLocation_functionName_set")
	symaddr(&addr_QPropertyBindingSourceLocation_line_get, "QPropertyBindingSourceLocation_line_get")
	symaddr(&addr_QPropertyBindingSourceLocation_line_set, "QPropertyBindingSourceLocation_line_set")
	symaddr(&addr_QPropertyBindingSourceLocation_column_get, "QPropertyBindingSourceLocation_column_get")
	symaddr(&addr_QPropertyBindingSourceLocation_column_set, "QPropertyBindingSourceLocation_column_set")
	reg(&pg_QPropertyBindingSourceLocation_delete, "QPropertyBindingSourceLocation_delete")
}

type QPropertyBindingSourceLocation struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPropertyBindingSourceLocation) UnsafeGoID() int64     { return this.vid }
func (this *QPropertyBindingSourceLocation) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPropertyBindingSourceLocation(h uintptr) *QPropertyBindingSourceLocation {
	if h == 0 {
		return nil
	}
	return &QPropertyBindingSourceLocation{h: h}
}

func newownedQPropertyBindingSourceLocation(h uintptr) *QPropertyBindingSourceLocation {
	o := bindQPropertyBindingSourceLocation(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyBindingSourceLocation).Delete)
	}
	return o
}

func (this *QPropertyBindingSourceLocation) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyBindingSourceLocation) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyBindingSourceLocation(h unsafe.Pointer) *QPropertyBindingSourceLocation { return bindQPropertyBindingSourceLocation(uintptr(h)) }
func UnsafeNewQPropertyBindingSourceLocationOwned(h unsafe.Pointer) *QPropertyBindingSourceLocation { return newownedQPropertyBindingSourceLocation(uintptr(h)) }

func (this *QPropertyBindingSourceLocation) UnsafeRelease() { this.h = 0 }

func NewQPropertyBindingSourceLocation() *QPropertyBindingSourceLocation {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_new)
	__o := newownedQPropertyBindingSourceLocation(__r)
	return __o
}

func NewQPropertyBindingSourceLocation2(arg1 *QPropertyBindingSourceLocation) *QPropertyBindingSourceLocation {
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPropertyBindingSourceLocation(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPropertyBindingSourceLocation) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_fileName_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPropertyBindingSourceLocation) SetFileName(fileName string) {
	fileName_b := append([]byte(fileName), 0)
	purego.SyscallN(addr_QPropertyBindingSourceLocation_fileName_set, this.h, uintptr(unsafe.Pointer(&fileName_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName_b)
}

func (this *QPropertyBindingSourceLocation) FunctionName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_functionName_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPropertyBindingSourceLocation) SetFunctionName(functionName string) {
	functionName_b := append([]byte(functionName), 0)
	purego.SyscallN(addr_QPropertyBindingSourceLocation_functionName_set, this.h, uintptr(unsafe.Pointer(&functionName_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(functionName_b)
}

func (this *QPropertyBindingSourceLocation) Line() uint {
	__r0, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_line_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPropertyBindingSourceLocation) SetLine(line uint) {
	purego.SyscallN(addr_QPropertyBindingSourceLocation_line_set, this.h, uintptr(line))
	runtime.KeepAlive(this)
}

func (this *QPropertyBindingSourceLocation) Column() uint {
	__r0, _, _ := purego.SyscallN(addr_QPropertyBindingSourceLocation_column_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QPropertyBindingSourceLocation) SetColumn(column uint) {
	purego.SyscallN(addr_QPropertyBindingSourceLocation_column_set, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QPropertyBindingSourceLocation) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyBindingSourceLocation_delete(this.h)
	}
	this.UnsafeRelease()
}
