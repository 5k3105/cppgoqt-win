package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDirIterator_new uintptr
	addr_QDirIterator_new2 uintptr
	addr_QDirIterator_new3 uintptr
	addr_QDirIterator_new4 uintptr
	addr_QDirIterator_new5 uintptr
	addr_QDirIterator_new6 uintptr
	addr_QDirIterator_new7 uintptr
	addr_QDirIterator_new8 uintptr
	addr_QDirIterator_new9 uintptr
	addr_QDirIterator_next uintptr
	addr_QDirIterator_nextFileInfo uintptr
	addr_QDirIterator_hasNext uintptr
	addr_QDirIterator_fileName uintptr
	addr_QDirIterator_filePath uintptr
	addr_QDirIterator_fileInfo uintptr
	addr_QDirIterator_path uintptr
	pg_QDirIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDirIterator_new, "QDirIterator_new")
	symaddr(&addr_QDirIterator_new2, "QDirIterator_new2")
	symaddr(&addr_QDirIterator_new3, "QDirIterator_new3")
	symaddr(&addr_QDirIterator_new4, "QDirIterator_new4")
	symaddr(&addr_QDirIterator_new5, "QDirIterator_new5")
	symaddr(&addr_QDirIterator_new6, "QDirIterator_new6")
	symaddr(&addr_QDirIterator_new7, "QDirIterator_new7")
	symaddr(&addr_QDirIterator_new8, "QDirIterator_new8")
	symaddr(&addr_QDirIterator_new9, "QDirIterator_new9")
	symaddr(&addr_QDirIterator_next, "QDirIterator_next")
	symaddr(&addr_QDirIterator_nextFileInfo, "QDirIterator_nextFileInfo")
	symaddr(&addr_QDirIterator_hasNext, "QDirIterator_hasNext")
	symaddr(&addr_QDirIterator_fileName, "QDirIterator_fileName")
	symaddr(&addr_QDirIterator_filePath, "QDirIterator_filePath")
	symaddr(&addr_QDirIterator_fileInfo, "QDirIterator_fileInfo")
	symaddr(&addr_QDirIterator_path, "QDirIterator_path")
	reg(&pg_QDirIterator_delete, "QDirIterator_delete")
}

type QDirIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDirIterator) UnsafeGoID() int64     { return this.vid }
func (this *QDirIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDirIterator(h uintptr) *QDirIterator {
	if h == 0 {
		return nil
	}
	return &QDirIterator{h: h}
}

func newownedQDirIterator(h uintptr) *QDirIterator {
	o := bindQDirIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDirIterator).Delete)
	}
	return o
}

func (this *QDirIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDirIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDirIterator(h unsafe.Pointer) *QDirIterator { return bindQDirIterator(uintptr(h)) }
func UnsafeNewQDirIteratorOwned(h unsafe.Pointer) *QDirIterator { return newownedQDirIterator(uintptr(h)) }

func (this *QDirIterator) UnsafeRelease() { this.h = 0 }

func NewQDirIterator(dir *QDir) *QDirIterator {
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new, uintptr(dir.UnsafePointer()))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(dir)
	return __o
}

func NewQDirIterator2(path string) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new2, path_p, uintptr(len(path)))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirIterator3(path string, filter QDir__Filter) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new3, path_p, uintptr(len(path)), uintptr(filter))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirIterator4(path string, nameFilters []string) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new4, path_p, uintptr(len(path)), nameFilters_p, uintptr(len(nameFilters)))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	return __o
}

func NewQDirIterator5(dir *QDir, flags QDirIterator__IteratorFlag) *QDirIterator {
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new5, uintptr(dir.UnsafePointer()), uintptr(flags))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(dir)
	return __o
}

func NewQDirIterator6(path string, flags QDirIterator__IteratorFlag) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new6, path_p, uintptr(len(path)), uintptr(flags))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirIterator7(path string, filter QDir__Filter, flags QDirIterator__IteratorFlag) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new7, path_p, uintptr(len(path)), uintptr(filter), uintptr(flags))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	return __o
}

func NewQDirIterator8(path string, nameFilters []string, filters QDir__Filter) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new8, path_p, uintptr(len(path)), nameFilters_p, uintptr(len(nameFilters)), uintptr(filters))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	return __o
}

func NewQDirIterator9(path string, nameFilters []string, filters QDir__Filter, flags QDirIterator__IteratorFlag) *QDirIterator {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	nameFilters_arr := make([]uintptr, len(nameFilters))
	for __i, __e := range nameFilters {
		nameFilters_arr[__i] = pgCString(__e)
	}
	var nameFilters_p uintptr
	if len(nameFilters_arr) > 0 {
		nameFilters_p = uintptr(unsafe.Pointer(&nameFilters_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QDirIterator_new9, path_p, uintptr(len(path)), nameFilters_p, uintptr(len(nameFilters)), uintptr(filters), uintptr(flags))
	__o := newownedQDirIterator(__r)
	runtime.KeepAlive(path)
	for __i := range nameFilters_arr {
		pgFree(nameFilters_arr[__i])
	}
	runtime.KeepAlive(nameFilters_arr)
	return __o
}

func (this *QDirIterator) Next() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirIterator_next, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirIterator) NextFileInfo() *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QDirIterator_nextFileInfo, this.h)
	runtime.KeepAlive(this)
	return newownedQFileInfo(__r)
}

func (this *QDirIterator) HasNext() bool {
	__r0, _, _ := purego.SyscallN(addr_QDirIterator_hasNext, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDirIterator) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirIterator_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirIterator) FilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirIterator_filePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirIterator) FileInfo() *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QDirIterator_fileInfo, this.h)
	runtime.KeepAlive(this)
	return newownedQFileInfo(__r)
}

func (this *QDirIterator) Path() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDirIterator_path, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDirIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDirIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
