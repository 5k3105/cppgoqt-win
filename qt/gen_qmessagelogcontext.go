package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMessageLogContext_new uintptr
	addr_QMessageLogContext_new2 uintptr
	addr_QMessageLogContext_version_get uintptr
	addr_QMessageLogContext_version_set uintptr
	addr_QMessageLogContext_line_get uintptr
	addr_QMessageLogContext_line_set uintptr
	addr_QMessageLogContext_file_get uintptr
	addr_QMessageLogContext_file_set uintptr
	addr_QMessageLogContext_function_get uintptr
	addr_QMessageLogContext_function_set uintptr
	addr_QMessageLogContext_category_get uintptr
	addr_QMessageLogContext_category_set uintptr
	pg_QMessageLogContext_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMessageLogContext_new, "QMessageLogContext_new")
	symaddr(&addr_QMessageLogContext_new2, "QMessageLogContext_new2")
	symaddr(&addr_QMessageLogContext_version_get, "QMessageLogContext_version_get")
	symaddr(&addr_QMessageLogContext_version_set, "QMessageLogContext_version_set")
	symaddr(&addr_QMessageLogContext_line_get, "QMessageLogContext_line_get")
	symaddr(&addr_QMessageLogContext_line_set, "QMessageLogContext_line_set")
	symaddr(&addr_QMessageLogContext_file_get, "QMessageLogContext_file_get")
	symaddr(&addr_QMessageLogContext_file_set, "QMessageLogContext_file_set")
	symaddr(&addr_QMessageLogContext_function_get, "QMessageLogContext_function_get")
	symaddr(&addr_QMessageLogContext_function_set, "QMessageLogContext_function_set")
	symaddr(&addr_QMessageLogContext_category_get, "QMessageLogContext_category_get")
	symaddr(&addr_QMessageLogContext_category_set, "QMessageLogContext_category_set")
	reg(&pg_QMessageLogContext_delete, "QMessageLogContext_delete")
}

type QMessageLogContext struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMessageLogContext) UnsafeGoID() int64     { return this.vid }
func (this *QMessageLogContext) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMessageLogContext(h uintptr) *QMessageLogContext {
	if h == 0 {
		return nil
	}
	return &QMessageLogContext{h: h}
}

func newownedQMessageLogContext(h uintptr) *QMessageLogContext {
	o := bindQMessageLogContext(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMessageLogContext).Delete)
	}
	return o
}

func (this *QMessageLogContext) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMessageLogContext) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMessageLogContext(h unsafe.Pointer) *QMessageLogContext { return bindQMessageLogContext(uintptr(h)) }
func UnsafeNewQMessageLogContextOwned(h unsafe.Pointer) *QMessageLogContext { return newownedQMessageLogContext(uintptr(h)) }

func (this *QMessageLogContext) UnsafeRelease() { this.h = 0 }

func NewQMessageLogContext() *QMessageLogContext {
	__r, _, _ := purego.SyscallN(addr_QMessageLogContext_new)
	__o := newownedQMessageLogContext(__r)
	return __o
}

func NewQMessageLogContext2(fileName string, lineNumber int, functionName string, categoryName string) *QMessageLogContext {
	fileName_b := append([]byte(fileName), 0)
	functionName_b := append([]byte(functionName), 0)
	categoryName_b := append([]byte(categoryName), 0)
	__r, _, _ := purego.SyscallN(addr_QMessageLogContext_new2, uintptr(unsafe.Pointer(&fileName_b[0])), uintptr(lineNumber), uintptr(unsafe.Pointer(&functionName_b[0])), uintptr(unsafe.Pointer(&categoryName_b[0])))
	__o := newownedQMessageLogContext(__r)
	runtime.KeepAlive(fileName_b)
	runtime.KeepAlive(functionName_b)
	runtime.KeepAlive(categoryName_b)
	return __o
}

func (this *QMessageLogContext) Version() int {
	__r0, _, _ := purego.SyscallN(addr_QMessageLogContext_version_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageLogContext) SetVersion(version int) {
	purego.SyscallN(addr_QMessageLogContext_version_set, this.h, uintptr(version))
	runtime.KeepAlive(this)
}

func (this *QMessageLogContext) Line() int {
	__r0, _, _ := purego.SyscallN(addr_QMessageLogContext_line_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMessageLogContext) SetLine(line int) {
	purego.SyscallN(addr_QMessageLogContext_line_set, this.h, uintptr(line))
	runtime.KeepAlive(this)
}

func (this *QMessageLogContext) File() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageLogContext_file_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageLogContext) SetFile(file string) {
	file_b := append([]byte(file), 0)
	purego.SyscallN(addr_QMessageLogContext_file_set, this.h, uintptr(unsafe.Pointer(&file_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(file_b)
}

func (this *QMessageLogContext) Function() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageLogContext_function_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageLogContext) SetFunction(function string) {
	function_b := append([]byte(function), 0)
	purego.SyscallN(addr_QMessageLogContext_function_set, this.h, uintptr(unsafe.Pointer(&function_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(function_b)
}

func (this *QMessageLogContext) Category() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMessageLogContext_category_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMessageLogContext) SetCategory(category string) {
	category_b := append([]byte(category), 0)
	purego.SyscallN(addr_QMessageLogContext_category_set, this.h, uintptr(unsafe.Pointer(&category_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(category_b)
}

func (this *QMessageLogContext) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMessageLogContext_delete(this.h)
	}
	this.UnsafeRelease()
}
