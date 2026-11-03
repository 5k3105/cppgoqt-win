package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringDecoder_chain uintptr
	addr_QStringDecoder_new uintptr
	addr_QStringDecoder_new2 uintptr
	addr_QStringDecoder_new3 uintptr
	addr_QStringDecoder_new4 uintptr
	addr_QStringDecoder_new5 uintptr
	addr_QStringDecoder_requiredSpace uintptr
	addr_QStringDecoder_appendToBuffer uintptr
	addr_QStringDecoder_encodingForName uintptr
	addr_QStringDecoder_nameForEncoding uintptr
	addr_QStringDecoder_encodingForData uintptr
	addr_QStringDecoder_encodingForHtml uintptr
	addr_QStringDecoder_availableCodecs uintptr
	addr_QStringDecoder_encodingForData2 uintptr
	pg_QStringDecoder_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringDecoder_chain, "QStringDecoder_chain")
	symaddr(&addr_QStringDecoder_new, "QStringDecoder_new")
	symaddr(&addr_QStringDecoder_new2, "QStringDecoder_new2")
	symaddr(&addr_QStringDecoder_new3, "QStringDecoder_new3")
	symaddr(&addr_QStringDecoder_new4, "QStringDecoder_new4")
	symaddr(&addr_QStringDecoder_new5, "QStringDecoder_new5")
	symaddr(&addr_QStringDecoder_requiredSpace, "QStringDecoder_requiredSpace")
	symaddr(&addr_QStringDecoder_appendToBuffer, "QStringDecoder_appendToBuffer")
	symaddr(&addr_QStringDecoder_encodingForName, "QStringDecoder_encodingForName")
	symaddr(&addr_QStringDecoder_nameForEncoding, "QStringDecoder_nameForEncoding")
	symaddr(&addr_QStringDecoder_encodingForData, "QStringDecoder_encodingForData")
	symaddr(&addr_QStringDecoder_encodingForHtml, "QStringDecoder_encodingForHtml")
	symaddr(&addr_QStringDecoder_availableCodecs, "QStringDecoder_availableCodecs")
	symaddr(&addr_QStringDecoder_encodingForData2, "QStringDecoder_encodingForData2")
	reg(&pg_QStringDecoder_delete, "QStringDecoder_delete")
}

type QStringDecoder struct {
	h uintptr
	*QStringConverter
}

func bindQStringDecoder(h uintptr) *QStringDecoder {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStringDecoder_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStringDecoder{h: h, QStringConverter: &QStringConverter{h: __c[0], QStringConverterBase: &QStringConverterBase{h: __c[1]}}}
}

func newownedQStringDecoder(h uintptr) *QStringDecoder {
	o := bindQStringDecoder(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStringDecoder).Delete)
	}
	return o
}

func (this *QStringDecoder) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringDecoder) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringDecoder(h unsafe.Pointer) *QStringDecoder { return bindQStringDecoder(uintptr(h)) }
func UnsafeNewQStringDecoderOwned(h unsafe.Pointer) *QStringDecoder { return newownedQStringDecoder(uintptr(h)) }

func (this *QStringDecoder) UnsafeRelease() {
	this.h = 0
	if this.QStringConverter != nil {
		this.QStringConverter.UnsafeRelease()
	}
}

func NewQStringDecoder(encoding QStringConverter__Encoding) *QStringDecoder {
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_new, uintptr(encoding))
	__o := newownedQStringDecoder(__r)
	return __o
}

func NewQStringDecoder2() *QStringDecoder {
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_new2)
	__o := newownedQStringDecoder(__r)
	return __o
}

func NewQStringDecoder3(name string) *QStringDecoder {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_new3, name_p, uintptr(len(name)))
	__o := newownedQStringDecoder(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQStringDecoder4(encoding QStringConverter__Encoding, flags QStringConverterBase__Flag) *QStringDecoder {
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_new4, uintptr(encoding), uintptr(flags))
	__o := newownedQStringDecoder(__r)
	return __o
}

func NewQStringDecoder5(name string, f QStringConverterBase__Flag) *QStringDecoder {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_new5, name_p, uintptr(len(name)), uintptr(f))
	__o := newownedQStringDecoder(__r)
	runtime.KeepAlive(name)
	return __o
}

func (this *QStringDecoder) RequiredSpace(inputLength int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringDecoder_requiredSpace, this.h, uintptr(inputLength))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStringDecoder) AppendToBuffer(out *QChar, ba *QByteArrayView) *QChar {
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_appendToBuffer, this.h, uintptr(out.UnsafePointer()), uintptr(ba.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
	runtime.KeepAlive(ba)
	return bindQChar(__r)
}

func QStringDecoder_EncodingForName(name string) (QStringConverter__Encoding, bool) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringDecoder_encodingForName, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringDecoder_NameForEncoding(e QStringConverter__Encoding) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_nameForEncoding, uintptr(e), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QStringDecoder_EncodingForData(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringDecoder_encodingForData, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringDecoder_EncodingForHtml(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringDecoder_encodingForHtml, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringDecoder_AvailableCodecs() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringDecoder_availableCodecs, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStringDecoder_EncodingForData2(data *QByteArrayView, expectedFirstCharacter uint16) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringDecoder_encodingForData2, uintptr(data.UnsafePointer()), uintptr(expectedFirstCharacter), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func (this *QStringDecoder) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringDecoder_delete(this.h)
	}
	this.UnsafeRelease()
}
