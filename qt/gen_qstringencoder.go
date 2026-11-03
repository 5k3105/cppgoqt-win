package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringEncoder_chain uintptr
	addr_QStringEncoder_new uintptr
	addr_QStringEncoder_new2 uintptr
	addr_QStringEncoder_new3 uintptr
	addr_QStringEncoder_new4 uintptr
	addr_QStringEncoder_new5 uintptr
	addr_QStringEncoder_requiredSpace uintptr
	addr_QStringEncoder_encodingForName uintptr
	addr_QStringEncoder_nameForEncoding uintptr
	addr_QStringEncoder_encodingForData uintptr
	addr_QStringEncoder_encodingForHtml uintptr
	addr_QStringEncoder_availableCodecs uintptr
	addr_QStringEncoder_encodingForData2 uintptr
	pg_QStringEncoder_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringEncoder_chain, "QStringEncoder_chain")
	symaddr(&addr_QStringEncoder_new, "QStringEncoder_new")
	symaddr(&addr_QStringEncoder_new2, "QStringEncoder_new2")
	symaddr(&addr_QStringEncoder_new3, "QStringEncoder_new3")
	symaddr(&addr_QStringEncoder_new4, "QStringEncoder_new4")
	symaddr(&addr_QStringEncoder_new5, "QStringEncoder_new5")
	symaddr(&addr_QStringEncoder_requiredSpace, "QStringEncoder_requiredSpace")
	symaddr(&addr_QStringEncoder_encodingForName, "QStringEncoder_encodingForName")
	symaddr(&addr_QStringEncoder_nameForEncoding, "QStringEncoder_nameForEncoding")
	symaddr(&addr_QStringEncoder_encodingForData, "QStringEncoder_encodingForData")
	symaddr(&addr_QStringEncoder_encodingForHtml, "QStringEncoder_encodingForHtml")
	symaddr(&addr_QStringEncoder_availableCodecs, "QStringEncoder_availableCodecs")
	symaddr(&addr_QStringEncoder_encodingForData2, "QStringEncoder_encodingForData2")
	reg(&pg_QStringEncoder_delete, "QStringEncoder_delete")
}

type QStringEncoder struct {
	h uintptr
	*QStringConverter
}

func bindQStringEncoder(h uintptr) *QStringEncoder {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStringEncoder_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStringEncoder{h: h, QStringConverter: &QStringConverter{h: __c[0], QStringConverterBase: &QStringConverterBase{h: __c[1]}}}
}

func newownedQStringEncoder(h uintptr) *QStringEncoder {
	o := bindQStringEncoder(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStringEncoder).Delete)
	}
	return o
}

func (this *QStringEncoder) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringEncoder) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringEncoder(h unsafe.Pointer) *QStringEncoder { return bindQStringEncoder(uintptr(h)) }
func UnsafeNewQStringEncoderOwned(h unsafe.Pointer) *QStringEncoder { return newownedQStringEncoder(uintptr(h)) }

func (this *QStringEncoder) UnsafeRelease() {
	this.h = 0
	if this.QStringConverter != nil {
		this.QStringConverter.UnsafeRelease()
	}
}

func NewQStringEncoder() *QStringEncoder {
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_new)
	__o := newownedQStringEncoder(__r)
	return __o
}

func NewQStringEncoder2(encoding QStringConverter__Encoding) *QStringEncoder {
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_new2, uintptr(encoding))
	__o := newownedQStringEncoder(__r)
	return __o
}

func NewQStringEncoder3(name string) *QStringEncoder {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_new3, name_p, uintptr(len(name)))
	__o := newownedQStringEncoder(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQStringEncoder4(encoding QStringConverter__Encoding, flags QStringConverterBase__Flag) *QStringEncoder {
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_new4, uintptr(encoding), uintptr(flags))
	__o := newownedQStringEncoder(__r)
	return __o
}

func NewQStringEncoder5(name string, flags QStringConverterBase__Flag) *QStringEncoder {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_new5, name_p, uintptr(len(name)), uintptr(flags))
	__o := newownedQStringEncoder(__r)
	runtime.KeepAlive(name)
	return __o
}

func (this *QStringEncoder) RequiredSpace(inputLength int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringEncoder_requiredSpace, this.h, uintptr(inputLength))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QStringEncoder_EncodingForName(name string) (QStringConverter__Encoding, bool) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringEncoder_encodingForName, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringEncoder_NameForEncoding(e QStringConverter__Encoding) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_nameForEncoding, uintptr(e), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QStringEncoder_EncodingForData(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringEncoder_encodingForData, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringEncoder_EncodingForHtml(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringEncoder_encodingForHtml, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringEncoder_AvailableCodecs() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringEncoder_availableCodecs, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStringEncoder_EncodingForData2(data *QByteArrayView, expectedFirstCharacter uint16) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringEncoder_encodingForData2, uintptr(data.UnsafePointer()), uintptr(expectedFirstCharacter), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func (this *QStringEncoder) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringEncoder_delete(this.h)
	}
	this.UnsafeRelease()
}
