package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStringConverter_base func(uintptr) uintptr
	addr_QStringConverter_isValid uintptr
	addr_QStringConverter_resetState uintptr
	addr_QStringConverter_hasError uintptr
	addr_QStringConverter_name uintptr
	addr_QStringConverter_encodingForName uintptr
	addr_QStringConverter_nameForEncoding uintptr
	addr_QStringConverter_encodingForData uintptr
	addr_QStringConverter_encodingForHtml uintptr
	addr_QStringConverter_availableCodecs uintptr
	addr_QStringConverter_encodingForData2 uintptr
)

func init() {
	reg(&pg_QStringConverter_base, "QStringConverter_base")
	symaddr(&addr_QStringConverter_isValid, "QStringConverter_isValid")
	symaddr(&addr_QStringConverter_resetState, "QStringConverter_resetState")
	symaddr(&addr_QStringConverter_hasError, "QStringConverter_hasError")
	symaddr(&addr_QStringConverter_name, "QStringConverter_name")
	symaddr(&addr_QStringConverter_encodingForName, "QStringConverter_encodingForName")
	symaddr(&addr_QStringConverter_nameForEncoding, "QStringConverter_nameForEncoding")
	symaddr(&addr_QStringConverter_encodingForData, "QStringConverter_encodingForData")
	symaddr(&addr_QStringConverter_encodingForHtml, "QStringConverter_encodingForHtml")
	symaddr(&addr_QStringConverter_availableCodecs, "QStringConverter_availableCodecs")
	symaddr(&addr_QStringConverter_encodingForData2, "QStringConverter_encodingForData2")
}

type QStringConverter struct {
	h uintptr
	*QStringConverterBase
}

func bindQStringConverter(h uintptr) *QStringConverter {
	if h == 0 {
		return nil
	}
	return &QStringConverter{h: h, QStringConverterBase: bindQStringConverterBase(pg_QStringConverter_base(h))}
}

func newownedQStringConverter(h uintptr) *QStringConverter { return bindQStringConverter(h) }

func (this *QStringConverter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringConverter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringConverter(h unsafe.Pointer) *QStringConverter { return bindQStringConverter(uintptr(h)) }
func UnsafeNewQStringConverterOwned(h unsafe.Pointer) *QStringConverter { return newownedQStringConverter(uintptr(h)) }

func (this *QStringConverter) UnsafeRelease() {
	this.h = 0
	if this.QStringConverterBase != nil {
		this.QStringConverterBase.UnsafeRelease()
	}
}

func (this *QStringConverter) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringConverter) ResetState() {
	purego.SyscallN(addr_QStringConverter_resetState, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringConverter) HasError() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_hasError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringConverter) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringConverter_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QStringConverter_EncodingForName(name string) (QStringConverter__Encoding, bool) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_encodingForName, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringConverter_NameForEncoding(e QStringConverter__Encoding) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringConverter_nameForEncoding, uintptr(e), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QStringConverter_EncodingForData(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_encodingForData, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringConverter_EncodingForHtml(data *QByteArrayView) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_encodingForHtml, uintptr(data.UnsafePointer()), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func QStringConverter_AvailableCodecs() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringConverter_availableCodecs, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStringConverter_EncodingForData2(data *QByteArrayView, expectedFirstCharacter uint16) (QStringConverter__Encoding, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QStringConverter_encodingForData2, uintptr(data.UnsafePointer()), uintptr(expectedFirstCharacter), uintptr(unsafe.Pointer(&__b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(data)
	return QStringConverter__Encoding(__r), __b != 0
}

func (this *QStringConverter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
