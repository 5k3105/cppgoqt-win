package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborStreamWriter_new uintptr
	addr_QCborStreamWriter_new2 uintptr
	addr_QCborStreamWriter_setDevice uintptr
	addr_QCborStreamWriter_device uintptr
	addr_QCborStreamWriter_append uintptr
	addr_QCborStreamWriter_append2 uintptr
	addr_QCborStreamWriter_append3 uintptr
	addr_QCborStreamWriter_append4 uintptr
	addr_QCborStreamWriter_append5 uintptr
	addr_QCborStreamWriter_append6 uintptr
	addr_QCborStreamWriter_append7 uintptr
	addr_QCborStreamWriter_append8 uintptr
	addr_QCborStreamWriter_append9 uintptr
	addr_QCborStreamWriter_append10 uintptr
	pg_QCborStreamWriter_append11 func(uintptr, float32)
	pg_QCborStreamWriter_append12 func(uintptr, float64)
	addr_QCborStreamWriter_appendByteString uintptr
	addr_QCborStreamWriter_appendTextString uintptr
	addr_QCborStreamWriter_append13 uintptr
	addr_QCborStreamWriter_appendNull uintptr
	addr_QCborStreamWriter_appendUndefined uintptr
	addr_QCborStreamWriter_append14 uintptr
	addr_QCborStreamWriter_append15 uintptr
	addr_QCborStreamWriter_append16 uintptr
	addr_QCborStreamWriter_startArray uintptr
	addr_QCborStreamWriter_startArray2 uintptr
	addr_QCborStreamWriter_endArray uintptr
	addr_QCborStreamWriter_startMap uintptr
	addr_QCborStreamWriter_startMap2 uintptr
	addr_QCborStreamWriter_endMap uintptr
	addr_QCborStreamWriter_append17 uintptr
	pg_QCborStreamWriter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborStreamWriter_new, "QCborStreamWriter_new")
	symaddr(&addr_QCborStreamWriter_new2, "QCborStreamWriter_new2")
	symaddr(&addr_QCborStreamWriter_setDevice, "QCborStreamWriter_setDevice")
	symaddr(&addr_QCborStreamWriter_device, "QCborStreamWriter_device")
	symaddr(&addr_QCborStreamWriter_append, "QCborStreamWriter_append")
	symaddr(&addr_QCborStreamWriter_append2, "QCborStreamWriter_append2")
	symaddr(&addr_QCborStreamWriter_append3, "QCborStreamWriter_append3")
	symaddr(&addr_QCborStreamWriter_append4, "QCborStreamWriter_append4")
	symaddr(&addr_QCborStreamWriter_append5, "QCborStreamWriter_append5")
	symaddr(&addr_QCborStreamWriter_append6, "QCborStreamWriter_append6")
	symaddr(&addr_QCborStreamWriter_append7, "QCborStreamWriter_append7")
	symaddr(&addr_QCborStreamWriter_append8, "QCborStreamWriter_append8")
	symaddr(&addr_QCborStreamWriter_append9, "QCborStreamWriter_append9")
	symaddr(&addr_QCborStreamWriter_append10, "QCborStreamWriter_append10")
	reg(&pg_QCborStreamWriter_append11, "QCborStreamWriter_append11")
	reg(&pg_QCborStreamWriter_append12, "QCborStreamWriter_append12")
	symaddr(&addr_QCborStreamWriter_appendByteString, "QCborStreamWriter_appendByteString")
	symaddr(&addr_QCborStreamWriter_appendTextString, "QCborStreamWriter_appendTextString")
	symaddr(&addr_QCborStreamWriter_append13, "QCborStreamWriter_append13")
	symaddr(&addr_QCborStreamWriter_appendNull, "QCborStreamWriter_appendNull")
	symaddr(&addr_QCborStreamWriter_appendUndefined, "QCborStreamWriter_appendUndefined")
	symaddr(&addr_QCborStreamWriter_append14, "QCborStreamWriter_append14")
	symaddr(&addr_QCborStreamWriter_append15, "QCborStreamWriter_append15")
	symaddr(&addr_QCborStreamWriter_append16, "QCborStreamWriter_append16")
	symaddr(&addr_QCborStreamWriter_startArray, "QCborStreamWriter_startArray")
	symaddr(&addr_QCborStreamWriter_startArray2, "QCborStreamWriter_startArray2")
	symaddr(&addr_QCborStreamWriter_endArray, "QCborStreamWriter_endArray")
	symaddr(&addr_QCborStreamWriter_startMap, "QCborStreamWriter_startMap")
	symaddr(&addr_QCborStreamWriter_startMap2, "QCborStreamWriter_startMap2")
	symaddr(&addr_QCborStreamWriter_endMap, "QCborStreamWriter_endMap")
	symaddr(&addr_QCborStreamWriter_append17, "QCborStreamWriter_append17")
	reg(&pg_QCborStreamWriter_delete, "QCborStreamWriter_delete")
}

type QCborStreamWriter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborStreamWriter) UnsafeGoID() int64     { return this.vid }
func (this *QCborStreamWriter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborStreamWriter(h uintptr) *QCborStreamWriter {
	if h == 0 {
		return nil
	}
	return &QCborStreamWriter{h: h}
}

func newownedQCborStreamWriter(h uintptr) *QCborStreamWriter {
	o := bindQCborStreamWriter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborStreamWriter).Delete)
	}
	return o
}

func (this *QCborStreamWriter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborStreamWriter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborStreamWriter(h unsafe.Pointer) *QCborStreamWriter { return bindQCborStreamWriter(uintptr(h)) }
func UnsafeNewQCborStreamWriterOwned(h unsafe.Pointer) *QCborStreamWriter { return newownedQCborStreamWriter(uintptr(h)) }

func (this *QCborStreamWriter) UnsafeRelease() { this.h = 0 }

func NewQCborStreamWriter(device *QIODevice) *QCborStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QCborStreamWriter_new, uintptr(device.UnsafePointer()))
	__o := newownedQCborStreamWriter(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQCborStreamWriter2(data *QByteArray) *QCborStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QCborStreamWriter_new2, uintptr(data.UnsafePointer()))
	__o := newownedQCborStreamWriter(__r)
	runtime.KeepAlive(data)
	return __o
}

func (this *QCborStreamWriter) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QCborStreamWriter_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QCborStreamWriter) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QCborStreamWriter_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QCborStreamWriter) Append(u uint64) {
	purego.SyscallN(addr_QCborStreamWriter_append, this.h, uintptr(u))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWith(i int64) {
	purego.SyscallN(addr_QCborStreamWriter_append2, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithQCborNegativeInteger(n QCborNegativeInteger) {
	purego.SyscallN(addr_QCborStreamWriter_append3, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithBa(ba *QByteArrayView) {
	purego.SyscallN(addr_QCborStreamWriter_append4, this.h, uintptr(ba.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ba)
}

func (this *QCborStreamWriter) AppendWithStr(str *QLatin1String) {
	purego.SyscallN(addr_QCborStreamWriter_append5, this.h, uintptr(str.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
}

func (this *QCborStreamWriter) Append2(str *QStringView) {
	purego.SyscallN(addr_QCborStreamWriter_append6, this.h, uintptr(str.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
}

func (this *QCborStreamWriter) AppendWithTag(tag QCborTag) {
	purego.SyscallN(addr_QCborStreamWriter_append7, this.h, uintptr(tag))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) Append3(tag QCborKnownTags) {
	purego.SyscallN(addr_QCborStreamWriter_append8, this.h, uintptr(tag))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithSt(st QCborSimpleType) {
	purego.SyscallN(addr_QCborStreamWriter_append9, this.h, uintptr(st))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithQfloat16(f *qfloat16) {
	purego.SyscallN(addr_QCborStreamWriter_append10, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QCborStreamWriter) AppendWithFloat(f float32) {
	pg_QCborStreamWriter_append11(this.h, f)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithDouble(d float64) {
	pg_QCborStreamWriter_append12(this.h, d)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendByteString(data string, len_ int64) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QCborStreamWriter_appendByteString, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QCborStreamWriter) AppendTextString(utf8 string, len_ int64) {
	utf8_b := append([]byte(utf8), 0)
	purego.SyscallN(addr_QCborStreamWriter_appendTextString, this.h, uintptr(unsafe.Pointer(&utf8_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(utf8_b)
}

func (this *QCborStreamWriter) AppendWithBool(b bool) {
	purego.SyscallN(addr_QCborStreamWriter_append13, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendNull() {
	purego.SyscallN(addr_QCborStreamWriter_appendNull, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendUndefined() {
	purego.SyscallN(addr_QCborStreamWriter_appendUndefined, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) AppendWithInt(i int) {
	purego.SyscallN(addr_QCborStreamWriter_append14, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) Append4(u uint) {
	purego.SyscallN(addr_QCborStreamWriter_append15, this.h, uintptr(u))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) Append5(str string) {
	str_b := append([]byte(str), 0)
	purego.SyscallN(addr_QCborStreamWriter_append16, this.h, uintptr(unsafe.Pointer(&str_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
}

func (this *QCborStreamWriter) StartArray() {
	purego.SyscallN(addr_QCborStreamWriter_startArray, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) StartArrayWithCount(count uint64) {
	purego.SyscallN(addr_QCborStreamWriter_startArray2, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) EndArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamWriter_endArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamWriter) StartMap() {
	purego.SyscallN(addr_QCborStreamWriter_startMap, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) StartMapWithCount(count uint64) {
	purego.SyscallN(addr_QCborStreamWriter_startMap2, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QCborStreamWriter) EndMap() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamWriter_endMap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamWriter) Append6(str string, size int64) {
	str_b := append([]byte(str), 0)
	purego.SyscallN(addr_QCborStreamWriter_append17, this.h, uintptr(unsafe.Pointer(&str_b[0])), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
}

func (this *QCborStreamWriter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborStreamWriter_delete(this.h)
	}
	this.UnsafeRelease()
}
