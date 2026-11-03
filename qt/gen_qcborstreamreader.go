package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCborStreamReader_new uintptr
	addr_QCborStreamReader_new2 uintptr
	addr_QCborStreamReader_new3 uintptr
	addr_QCborStreamReader_new4 uintptr
	addr_QCborStreamReader_new5 uintptr
	addr_QCborStreamReader_setDevice uintptr
	addr_QCborStreamReader_device uintptr
	addr_QCborStreamReader_addData uintptr
	addr_QCborStreamReader_addData2 uintptr
	addr_QCborStreamReader_addData3 uintptr
	addr_QCborStreamReader_reparse uintptr
	addr_QCborStreamReader_clear uintptr
	addr_QCborStreamReader_reset uintptr
	addr_QCborStreamReader_lastError uintptr
	addr_QCborStreamReader_currentOffset uintptr
	addr_QCborStreamReader_isValid uintptr
	addr_QCborStreamReader_containerDepth uintptr
	addr_QCborStreamReader_parentContainerType uintptr
	addr_QCborStreamReader_hasNext uintptr
	addr_QCborStreamReader_next uintptr
	addr_QCborStreamReader_type uintptr
	addr_QCborStreamReader_isUnsignedInteger uintptr
	addr_QCborStreamReader_isNegativeInteger uintptr
	addr_QCborStreamReader_isInteger uintptr
	addr_QCborStreamReader_isByteArray uintptr
	addr_QCborStreamReader_isString uintptr
	addr_QCborStreamReader_isArray uintptr
	addr_QCborStreamReader_isMap uintptr
	addr_QCborStreamReader_isTag uintptr
	addr_QCborStreamReader_isSimpleType uintptr
	addr_QCborStreamReader_isFloat16 uintptr
	addr_QCborStreamReader_isFloat uintptr
	addr_QCborStreamReader_isDouble uintptr
	addr_QCborStreamReader_isInvalid uintptr
	addr_QCborStreamReader_isSimpleType2 uintptr
	addr_QCborStreamReader_isFalse uintptr
	addr_QCborStreamReader_isTrue uintptr
	addr_QCborStreamReader_isBool uintptr
	addr_QCborStreamReader_isNull uintptr
	addr_QCborStreamReader_isUndefined uintptr
	addr_QCborStreamReader_isLengthKnown uintptr
	addr_QCborStreamReader_length uintptr
	addr_QCborStreamReader_isContainer uintptr
	addr_QCborStreamReader_enterContainer uintptr
	addr_QCborStreamReader_leaveContainer uintptr
	addr_QCborStreamReader_readAndAppendToString uintptr
	addr_QCborStreamReader_readAndAppendToUtf8String uintptr
	addr_QCborStreamReader_readAndAppendToByteArray uintptr
	addr_QCborStreamReader_currentStringChunkSize uintptr
	addr_QCborStreamReader_toBool uintptr
	addr_QCborStreamReader_toTag uintptr
	addr_QCborStreamReader_toUnsignedInteger uintptr
	addr_QCborStreamReader_toNegativeInteger uintptr
	addr_QCborStreamReader_toSimpleType uintptr
	addr_QCborStreamReader_toFloat16 uintptr
	pg_QCborStreamReader_toFloat func(uintptr) float32
	pg_QCborStreamReader_toDouble func(uintptr) float64
	addr_QCborStreamReader_toInteger uintptr
	addr_QCborStreamReader_readAllString uintptr
	addr_QCborStreamReader_readAllUtf8String uintptr
	addr_QCborStreamReader_readAllByteArray uintptr
	addr_QCborStreamReader_next2 uintptr
	pg_QCborStreamReader_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCborStreamReader_new, "QCborStreamReader_new")
	symaddr(&addr_QCborStreamReader_new2, "QCborStreamReader_new2")
	symaddr(&addr_QCborStreamReader_new3, "QCborStreamReader_new3")
	symaddr(&addr_QCborStreamReader_new4, "QCborStreamReader_new4")
	symaddr(&addr_QCborStreamReader_new5, "QCborStreamReader_new5")
	symaddr(&addr_QCborStreamReader_setDevice, "QCborStreamReader_setDevice")
	symaddr(&addr_QCborStreamReader_device, "QCborStreamReader_device")
	symaddr(&addr_QCborStreamReader_addData, "QCborStreamReader_addData")
	symaddr(&addr_QCborStreamReader_addData2, "QCborStreamReader_addData2")
	symaddr(&addr_QCborStreamReader_addData3, "QCborStreamReader_addData3")
	symaddr(&addr_QCborStreamReader_reparse, "QCborStreamReader_reparse")
	symaddr(&addr_QCborStreamReader_clear, "QCborStreamReader_clear")
	symaddr(&addr_QCborStreamReader_reset, "QCborStreamReader_reset")
	symaddr(&addr_QCborStreamReader_lastError, "QCborStreamReader_lastError")
	symaddr(&addr_QCborStreamReader_currentOffset, "QCborStreamReader_currentOffset")
	symaddr(&addr_QCborStreamReader_isValid, "QCborStreamReader_isValid")
	symaddr(&addr_QCborStreamReader_containerDepth, "QCborStreamReader_containerDepth")
	symaddr(&addr_QCborStreamReader_parentContainerType, "QCborStreamReader_parentContainerType")
	symaddr(&addr_QCborStreamReader_hasNext, "QCborStreamReader_hasNext")
	symaddr(&addr_QCborStreamReader_next, "QCborStreamReader_next")
	symaddr(&addr_QCborStreamReader_type, "QCborStreamReader_type")
	symaddr(&addr_QCborStreamReader_isUnsignedInteger, "QCborStreamReader_isUnsignedInteger")
	symaddr(&addr_QCborStreamReader_isNegativeInteger, "QCborStreamReader_isNegativeInteger")
	symaddr(&addr_QCborStreamReader_isInteger, "QCborStreamReader_isInteger")
	symaddr(&addr_QCborStreamReader_isByteArray, "QCborStreamReader_isByteArray")
	symaddr(&addr_QCborStreamReader_isString, "QCborStreamReader_isString")
	symaddr(&addr_QCborStreamReader_isArray, "QCborStreamReader_isArray")
	symaddr(&addr_QCborStreamReader_isMap, "QCborStreamReader_isMap")
	symaddr(&addr_QCborStreamReader_isTag, "QCborStreamReader_isTag")
	symaddr(&addr_QCborStreamReader_isSimpleType, "QCborStreamReader_isSimpleType")
	symaddr(&addr_QCborStreamReader_isFloat16, "QCborStreamReader_isFloat16")
	symaddr(&addr_QCborStreamReader_isFloat, "QCborStreamReader_isFloat")
	symaddr(&addr_QCborStreamReader_isDouble, "QCborStreamReader_isDouble")
	symaddr(&addr_QCborStreamReader_isInvalid, "QCborStreamReader_isInvalid")
	symaddr(&addr_QCborStreamReader_isSimpleType2, "QCborStreamReader_isSimpleType2")
	symaddr(&addr_QCborStreamReader_isFalse, "QCborStreamReader_isFalse")
	symaddr(&addr_QCborStreamReader_isTrue, "QCborStreamReader_isTrue")
	symaddr(&addr_QCborStreamReader_isBool, "QCborStreamReader_isBool")
	symaddr(&addr_QCborStreamReader_isNull, "QCborStreamReader_isNull")
	symaddr(&addr_QCborStreamReader_isUndefined, "QCborStreamReader_isUndefined")
	symaddr(&addr_QCborStreamReader_isLengthKnown, "QCborStreamReader_isLengthKnown")
	symaddr(&addr_QCborStreamReader_length, "QCborStreamReader_length")
	symaddr(&addr_QCborStreamReader_isContainer, "QCborStreamReader_isContainer")
	symaddr(&addr_QCborStreamReader_enterContainer, "QCborStreamReader_enterContainer")
	symaddr(&addr_QCborStreamReader_leaveContainer, "QCborStreamReader_leaveContainer")
	symaddr(&addr_QCborStreamReader_readAndAppendToString, "QCborStreamReader_readAndAppendToString")
	symaddr(&addr_QCborStreamReader_readAndAppendToUtf8String, "QCborStreamReader_readAndAppendToUtf8String")
	symaddr(&addr_QCborStreamReader_readAndAppendToByteArray, "QCborStreamReader_readAndAppendToByteArray")
	symaddr(&addr_QCborStreamReader_currentStringChunkSize, "QCborStreamReader_currentStringChunkSize")
	symaddr(&addr_QCborStreamReader_toBool, "QCborStreamReader_toBool")
	symaddr(&addr_QCborStreamReader_toTag, "QCborStreamReader_toTag")
	symaddr(&addr_QCborStreamReader_toUnsignedInteger, "QCborStreamReader_toUnsignedInteger")
	symaddr(&addr_QCborStreamReader_toNegativeInteger, "QCborStreamReader_toNegativeInteger")
	symaddr(&addr_QCborStreamReader_toSimpleType, "QCborStreamReader_toSimpleType")
	symaddr(&addr_QCborStreamReader_toFloat16, "QCborStreamReader_toFloat16")
	reg(&pg_QCborStreamReader_toFloat, "QCborStreamReader_toFloat")
	reg(&pg_QCborStreamReader_toDouble, "QCborStreamReader_toDouble")
	symaddr(&addr_QCborStreamReader_toInteger, "QCborStreamReader_toInteger")
	symaddr(&addr_QCborStreamReader_readAllString, "QCborStreamReader_readAllString")
	symaddr(&addr_QCborStreamReader_readAllUtf8String, "QCborStreamReader_readAllUtf8String")
	symaddr(&addr_QCborStreamReader_readAllByteArray, "QCborStreamReader_readAllByteArray")
	symaddr(&addr_QCborStreamReader_next2, "QCborStreamReader_next2")
	reg(&pg_QCborStreamReader_delete, "QCborStreamReader_delete")
}

type QCborStreamReader struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCborStreamReader) UnsafeGoID() int64     { return this.vid }
func (this *QCborStreamReader) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCborStreamReader(h uintptr) *QCborStreamReader {
	if h == 0 {
		return nil
	}
	return &QCborStreamReader{h: h}
}

func newownedQCborStreamReader(h uintptr) *QCborStreamReader {
	o := bindQCborStreamReader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCborStreamReader).Delete)
	}
	return o
}

func (this *QCborStreamReader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCborStreamReader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCborStreamReader(h unsafe.Pointer) *QCborStreamReader { return bindQCborStreamReader(uintptr(h)) }
func UnsafeNewQCborStreamReaderOwned(h unsafe.Pointer) *QCborStreamReader { return newownedQCborStreamReader(uintptr(h)) }

func (this *QCborStreamReader) UnsafeRelease() { this.h = 0 }

func NewQCborStreamReader() *QCborStreamReader {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_new)
	__o := newownedQCborStreamReader(__r)
	return __o
}

func NewQCborStreamReader2(data string, len_ int64) *QCborStreamReader {
	data_b := append([]byte(data), 0)
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_new2, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	__o := newownedQCborStreamReader(__r)
	runtime.KeepAlive(data_b)
	return __o
}

func NewQCborStreamReader3(data []byte, len_ int64) *QCborStreamReader {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_new3, data_p, uintptr(len_))
	__o := newownedQCborStreamReader(__r)
	runtime.KeepAlive(data)
	return __o
}

func NewQCborStreamReader4(data *QByteArray) *QCborStreamReader {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_new4, uintptr(data.UnsafePointer()))
	__o := newownedQCborStreamReader(__r)
	runtime.KeepAlive(data)
	return __o
}

func NewQCborStreamReader5(device *QIODevice) *QCborStreamReader {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_new5, uintptr(device.UnsafePointer()))
	__o := newownedQCborStreamReader(__r)
	runtime.KeepAlive(device)
	return __o
}

func (this *QCborStreamReader) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QCborStreamReader_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QCborStreamReader) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QCborStreamReader) AddData(data *QByteArray) {
	purego.SyscallN(addr_QCborStreamReader_addData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QCborStreamReader) AddData2(data string, len_ int64) {
	data_b := append([]byte(data), 0)
	purego.SyscallN(addr_QCborStreamReader_addData2, this.h, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data_b)
}

func (this *QCborStreamReader) AddData3(data []byte, len_ int64) {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	purego.SyscallN(addr_QCborStreamReader_addData3, this.h, data_p, uintptr(len_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QCborStreamReader) Reparse() {
	purego.SyscallN(addr_QCborStreamReader_reparse, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamReader) Clear() {
	purego.SyscallN(addr_QCborStreamReader_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamReader) Reset() {
	purego.SyscallN(addr_QCborStreamReader_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QCborStreamReader) LastError() *QCborError {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_lastError, this.h)
	runtime.KeepAlive(this)
	return newownedQCborError(__r)
}

func (this *QCborStreamReader) CurrentOffset() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_currentOffset, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborStreamReader) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) ContainerDepth() int {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_containerDepth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCborStreamReader) ParentContainerType() QCborStreamReader__Type {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_parentContainerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborStreamReader__Type(__r)
}

func (this *QCborStreamReader) HasNext() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_hasNext, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) Next() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_next, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) Type() QCborStreamReader__Type {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborStreamReader__Type(__r)
}

func (this *QCborStreamReader) IsUnsignedInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isUnsignedInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsNegativeInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isNegativeInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsInteger() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsByteArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isByteArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsString() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isString, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsArray() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isArray, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsMap() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isMap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsTag() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsSimpleType() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsFloat16() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isFloat16, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsFloat() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isFloat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsDouble() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isDouble, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsInvalid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isInvalid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsSimpleTypeWithSt(st QCborSimpleType) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isSimpleType2, this.h, uintptr(st))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsFalse() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isFalse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsTrue() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isTrue, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsUndefined() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isUndefined, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) IsLengthKnown() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isLengthKnown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) Length() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QCborStreamReader) IsContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_isContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) EnterContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_enterContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) LeaveContainer() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_leaveContainer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) ReadAndAppendToString(dst *QString) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_readAndAppendToString, this.h, uintptr(dst.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dst)
	return __r != 0
}

func (this *QCborStreamReader) ReadAndAppendToUtf8String(dst *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_readAndAppendToUtf8String, this.h, uintptr(dst.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dst)
	return __r != 0
}

func (this *QCborStreamReader) ReadAndAppendToByteArray(dst *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_readAndAppendToByteArray, this.h, uintptr(dst.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dst)
	return __r != 0
}

func (this *QCborStreamReader) CurrentStringChunkSize() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_currentStringChunkSize, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborStreamReader) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) ToTag() QCborTag {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toTag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborTag(__r)
}

func (this *QCborStreamReader) ToUnsignedInteger() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toUnsignedInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QCborStreamReader) ToNegativeInteger() QCborNegativeInteger {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toNegativeInteger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborNegativeInteger(__r)
}

func (this *QCborStreamReader) ToSimpleType() QCborSimpleType {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toSimpleType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCborSimpleType(__r)
}

func (this *QCborStreamReader) ToFloat16() *qfloat16 {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_toFloat16, this.h)
	runtime.KeepAlive(this)
	return newownedqfloat16(__r)
}

func (this *QCborStreamReader) ToFloat() float32 {
	__r := pg_QCborStreamReader_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborStreamReader) ToDouble() float64 {
	__r := pg_QCborStreamReader_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborStreamReader) ToInteger() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_toInteger, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QCborStreamReader) ReadAllString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_readAllString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCborStreamReader) ReadAllUtf8String() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_readAllUtf8String, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborStreamReader) ReadAllByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QCborStreamReader_readAllByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QCborStreamReader) NextWithMaxRecursion(maxRecursion int) bool {
	__r0, _, _ := purego.SyscallN(addr_QCborStreamReader_next2, this.h, uintptr(maxRecursion))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCborStreamReader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCborStreamReader_delete(this.h)
	}
	this.UnsafeRelease()
}
