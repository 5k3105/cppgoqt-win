package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextDocumentWriter_new uintptr
	addr_QTextDocumentWriter_new2 uintptr
	addr_QTextDocumentWriter_new3 uintptr
	addr_QTextDocumentWriter_new4 uintptr
	addr_QTextDocumentWriter_setFormat uintptr
	addr_QTextDocumentWriter_format uintptr
	addr_QTextDocumentWriter_setDevice uintptr
	addr_QTextDocumentWriter_device uintptr
	addr_QTextDocumentWriter_setFileName uintptr
	addr_QTextDocumentWriter_fileName uintptr
	addr_QTextDocumentWriter_write uintptr
	addr_QTextDocumentWriter_write2 uintptr
	addr_QTextDocumentWriter_supportedDocumentFormats uintptr
	pg_QTextDocumentWriter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextDocumentWriter_new, "QTextDocumentWriter_new")
	symaddr(&addr_QTextDocumentWriter_new2, "QTextDocumentWriter_new2")
	symaddr(&addr_QTextDocumentWriter_new3, "QTextDocumentWriter_new3")
	symaddr(&addr_QTextDocumentWriter_new4, "QTextDocumentWriter_new4")
	symaddr(&addr_QTextDocumentWriter_setFormat, "QTextDocumentWriter_setFormat")
	symaddr(&addr_QTextDocumentWriter_format, "QTextDocumentWriter_format")
	symaddr(&addr_QTextDocumentWriter_setDevice, "QTextDocumentWriter_setDevice")
	symaddr(&addr_QTextDocumentWriter_device, "QTextDocumentWriter_device")
	symaddr(&addr_QTextDocumentWriter_setFileName, "QTextDocumentWriter_setFileName")
	symaddr(&addr_QTextDocumentWriter_fileName, "QTextDocumentWriter_fileName")
	symaddr(&addr_QTextDocumentWriter_write, "QTextDocumentWriter_write")
	symaddr(&addr_QTextDocumentWriter_write2, "QTextDocumentWriter_write2")
	symaddr(&addr_QTextDocumentWriter_supportedDocumentFormats, "QTextDocumentWriter_supportedDocumentFormats")
	reg(&pg_QTextDocumentWriter_delete, "QTextDocumentWriter_delete")
}

type QTextDocumentWriter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextDocumentWriter) UnsafeGoID() int64     { return this.vid }
func (this *QTextDocumentWriter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextDocumentWriter(h uintptr) *QTextDocumentWriter {
	if h == 0 {
		return nil
	}
	return &QTextDocumentWriter{h: h}
}

func newownedQTextDocumentWriter(h uintptr) *QTextDocumentWriter {
	o := bindQTextDocumentWriter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextDocumentWriter).Delete)
	}
	return o
}

func (this *QTextDocumentWriter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextDocumentWriter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextDocumentWriter(h unsafe.Pointer) *QTextDocumentWriter { return bindQTextDocumentWriter(uintptr(h)) }
func UnsafeNewQTextDocumentWriterOwned(h unsafe.Pointer) *QTextDocumentWriter { return newownedQTextDocumentWriter(uintptr(h)) }

func (this *QTextDocumentWriter) UnsafeRelease() { this.h = 0 }

func NewQTextDocumentWriter() *QTextDocumentWriter {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_new)
	__o := newownedQTextDocumentWriter(__r)
	return __o
}

func NewQTextDocumentWriter2(device *QIODevice, format *QByteArray) *QTextDocumentWriter {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_new2, uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	__o := newownedQTextDocumentWriter(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	return __o
}

func NewQTextDocumentWriter3(fileName string) *QTextDocumentWriter {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_new3, fileName_p, uintptr(len(fileName)))
	__o := newownedQTextDocumentWriter(__r)
	runtime.KeepAlive(fileName)
	return __o
}

func NewQTextDocumentWriter4(fileName string, format *QByteArray) *QTextDocumentWriter {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_new4, fileName_p, uintptr(len(fileName)), uintptr(format.UnsafePointer()))
	__o := newownedQTextDocumentWriter(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format)
	return __o
}

func (this *QTextDocumentWriter) SetFormat(format *QByteArray) {
	purego.SyscallN(addr_QTextDocumentWriter_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextDocumentWriter) Format() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_format, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QTextDocumentWriter) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QTextDocumentWriter_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QTextDocumentWriter) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QTextDocumentWriter) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QTextDocumentWriter_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QTextDocumentWriter) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocumentWriter) Write(document *QTextDocument) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocumentWriter_write, this.h, uintptr(document.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(document)
	return __r != 0
}

func (this *QTextDocumentWriter) WriteWithFragment(fragment *QTextDocumentFragment) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocumentWriter_write2, this.h, uintptr(fragment.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragment)
	return __r != 0
}

func QTextDocumentWriter_SupportedDocumentFormats() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentWriter_supportedDocumentFormats, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextDocumentWriter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextDocumentWriter_delete(this.h)
	}
	this.UnsafeRelease()
}
