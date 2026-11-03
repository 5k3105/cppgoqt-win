package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextStream_base func(uintptr) uintptr
	addr_QTextStream_new uintptr
	addr_QTextStream_new2 uintptr
	addr_QTextStream_new3 uintptr
	addr_QTextStream_new4 uintptr
	addr_QTextStream_new5 uintptr
	addr_QTextStream_new6 uintptr
	addr_QTextStream_new7 uintptr
	addr_QTextStream_new8 uintptr
	addr_QTextStream_setEncoding uintptr
	addr_QTextStream_encoding uintptr
	addr_QTextStream_setAutoDetectUnicode uintptr
	addr_QTextStream_autoDetectUnicode uintptr
	addr_QTextStream_setGenerateByteOrderMark uintptr
	addr_QTextStream_generateByteOrderMark uintptr
	addr_QTextStream_setLocale uintptr
	addr_QTextStream_locale uintptr
	addr_QTextStream_setDevice uintptr
	addr_QTextStream_device uintptr
	addr_QTextStream_setString uintptr
	addr_QTextStream_string uintptr
	addr_QTextStream_status uintptr
	addr_QTextStream_setStatus uintptr
	addr_QTextStream_resetStatus uintptr
	addr_QTextStream_atEnd uintptr
	addr_QTextStream_reset uintptr
	addr_QTextStream_flush uintptr
	addr_QTextStream_seek uintptr
	addr_QTextStream_pos uintptr
	addr_QTextStream_skipWhiteSpace uintptr
	addr_QTextStream_readLine uintptr
	addr_QTextStream_readLineInto uintptr
	addr_QTextStream_readAll uintptr
	addr_QTextStream_read uintptr
	addr_QTextStream_setFieldAlignment uintptr
	addr_QTextStream_fieldAlignment uintptr
	addr_QTextStream_setPadChar uintptr
	addr_QTextStream_padChar uintptr
	addr_QTextStream_setFieldWidth uintptr
	addr_QTextStream_fieldWidth uintptr
	addr_QTextStream_setNumberFlags uintptr
	addr_QTextStream_numberFlags uintptr
	addr_QTextStream_setIntegerBase uintptr
	addr_QTextStream_integerBase uintptr
	addr_QTextStream_setRealNumberNotation uintptr
	addr_QTextStream_realNumberNotation uintptr
	addr_QTextStream_setRealNumberPrecision uintptr
	addr_QTextStream_realNumberPrecision uintptr
	addr_QTextStream_OperatorShiftRight uintptr
	addr_QTextStream_OperatorShiftRight2 uintptr
	addr_QTextStream_OperatorShiftRight3 uintptr
	addr_QTextStream_OperatorShiftRight4 uintptr
	addr_QTextStream_OperatorShiftRight5 uintptr
	addr_QTextStream_OperatorShiftRight6 uintptr
	addr_QTextStream_OperatorShiftRight7 uintptr
	addr_QTextStream_OperatorShiftRight8 uintptr
	addr_QTextStream_OperatorShiftRight9 uintptr
	addr_QTextStream_OperatorShiftRight10 uintptr
	addr_QTextStream_OperatorShiftRight11 uintptr
	addr_QTextStream_OperatorShiftRight12 uintptr
	addr_QTextStream_OperatorShiftRight13 uintptr
	addr_QTextStream_OperatorShiftRight14 uintptr
	addr_QTextStream_OperatorShiftRight15 uintptr
	addr_QTextStream_OperatorShiftRight16 uintptr
	addr_QTextStream_OperatorShiftLeft uintptr
	addr_QTextStream_OperatorShiftLeft2 uintptr
	addr_QTextStream_OperatorShiftLeft3 uintptr
	addr_QTextStream_OperatorShiftLeft4 uintptr
	addr_QTextStream_OperatorShiftLeft5 uintptr
	addr_QTextStream_OperatorShiftLeft6 uintptr
	addr_QTextStream_OperatorShiftLeft7 uintptr
	addr_QTextStream_OperatorShiftLeft8 uintptr
	addr_QTextStream_OperatorShiftLeft9 uintptr
	addr_QTextStream_OperatorShiftLeft10 uintptr
	addr_QTextStream_OperatorShiftLeft11 uintptr
	pg_QTextStream_OperatorShiftLeft12 func(uintptr, float32) uintptr
	pg_QTextStream_OperatorShiftLeft13 func(uintptr, float64) uintptr
	addr_QTextStream_OperatorShiftLeft14 uintptr
	addr_QTextStream_OperatorShiftLeft15 uintptr
	addr_QTextStream_OperatorShiftLeft16 uintptr
	addr_QTextStream_OperatorShiftLeft17 uintptr
	addr_QTextStream_OperatorShiftLeft18 uintptr
	addr_QTextStream_ToBool uintptr
	addr_QTextStream_setString2 uintptr
	addr_QTextStream_readLine2 uintptr
	addr_QTextStream_readLineInto2 uintptr
	addr_QTextStream_OperatorShiftLeft19 uintptr
	addr_QTextStream_OperatorShiftRight17 uintptr
	addr_QTextStream_OperatorShiftLeft20 uintptr
	pg_QTextStream_delete func(uintptr)
)

func init() {
	reg(&pg_QTextStream_base, "QTextStream_base")
	symaddr(&addr_QTextStream_new, "QTextStream_new")
	symaddr(&addr_QTextStream_new2, "QTextStream_new2")
	symaddr(&addr_QTextStream_new3, "QTextStream_new3")
	symaddr(&addr_QTextStream_new4, "QTextStream_new4")
	symaddr(&addr_QTextStream_new5, "QTextStream_new5")
	symaddr(&addr_QTextStream_new6, "QTextStream_new6")
	symaddr(&addr_QTextStream_new7, "QTextStream_new7")
	symaddr(&addr_QTextStream_new8, "QTextStream_new8")
	symaddr(&addr_QTextStream_setEncoding, "QTextStream_setEncoding")
	symaddr(&addr_QTextStream_encoding, "QTextStream_encoding")
	symaddr(&addr_QTextStream_setAutoDetectUnicode, "QTextStream_setAutoDetectUnicode")
	symaddr(&addr_QTextStream_autoDetectUnicode, "QTextStream_autoDetectUnicode")
	symaddr(&addr_QTextStream_setGenerateByteOrderMark, "QTextStream_setGenerateByteOrderMark")
	symaddr(&addr_QTextStream_generateByteOrderMark, "QTextStream_generateByteOrderMark")
	symaddr(&addr_QTextStream_setLocale, "QTextStream_setLocale")
	symaddr(&addr_QTextStream_locale, "QTextStream_locale")
	symaddr(&addr_QTextStream_setDevice, "QTextStream_setDevice")
	symaddr(&addr_QTextStream_device, "QTextStream_device")
	symaddr(&addr_QTextStream_setString, "QTextStream_setString")
	symaddr(&addr_QTextStream_string, "QTextStream_string")
	symaddr(&addr_QTextStream_status, "QTextStream_status")
	symaddr(&addr_QTextStream_setStatus, "QTextStream_setStatus")
	symaddr(&addr_QTextStream_resetStatus, "QTextStream_resetStatus")
	symaddr(&addr_QTextStream_atEnd, "QTextStream_atEnd")
	symaddr(&addr_QTextStream_reset, "QTextStream_reset")
	symaddr(&addr_QTextStream_flush, "QTextStream_flush")
	symaddr(&addr_QTextStream_seek, "QTextStream_seek")
	symaddr(&addr_QTextStream_pos, "QTextStream_pos")
	symaddr(&addr_QTextStream_skipWhiteSpace, "QTextStream_skipWhiteSpace")
	symaddr(&addr_QTextStream_readLine, "QTextStream_readLine")
	symaddr(&addr_QTextStream_readLineInto, "QTextStream_readLineInto")
	symaddr(&addr_QTextStream_readAll, "QTextStream_readAll")
	symaddr(&addr_QTextStream_read, "QTextStream_read")
	symaddr(&addr_QTextStream_setFieldAlignment, "QTextStream_setFieldAlignment")
	symaddr(&addr_QTextStream_fieldAlignment, "QTextStream_fieldAlignment")
	symaddr(&addr_QTextStream_setPadChar, "QTextStream_setPadChar")
	symaddr(&addr_QTextStream_padChar, "QTextStream_padChar")
	symaddr(&addr_QTextStream_setFieldWidth, "QTextStream_setFieldWidth")
	symaddr(&addr_QTextStream_fieldWidth, "QTextStream_fieldWidth")
	symaddr(&addr_QTextStream_setNumberFlags, "QTextStream_setNumberFlags")
	symaddr(&addr_QTextStream_numberFlags, "QTextStream_numberFlags")
	symaddr(&addr_QTextStream_setIntegerBase, "QTextStream_setIntegerBase")
	symaddr(&addr_QTextStream_integerBase, "QTextStream_integerBase")
	symaddr(&addr_QTextStream_setRealNumberNotation, "QTextStream_setRealNumberNotation")
	symaddr(&addr_QTextStream_realNumberNotation, "QTextStream_realNumberNotation")
	symaddr(&addr_QTextStream_setRealNumberPrecision, "QTextStream_setRealNumberPrecision")
	symaddr(&addr_QTextStream_realNumberPrecision, "QTextStream_realNumberPrecision")
	symaddr(&addr_QTextStream_OperatorShiftRight, "QTextStream_OperatorShiftRight")
	symaddr(&addr_QTextStream_OperatorShiftRight2, "QTextStream_OperatorShiftRight2")
	symaddr(&addr_QTextStream_OperatorShiftRight3, "QTextStream_OperatorShiftRight3")
	symaddr(&addr_QTextStream_OperatorShiftRight4, "QTextStream_OperatorShiftRight4")
	symaddr(&addr_QTextStream_OperatorShiftRight5, "QTextStream_OperatorShiftRight5")
	symaddr(&addr_QTextStream_OperatorShiftRight6, "QTextStream_OperatorShiftRight6")
	symaddr(&addr_QTextStream_OperatorShiftRight7, "QTextStream_OperatorShiftRight7")
	symaddr(&addr_QTextStream_OperatorShiftRight8, "QTextStream_OperatorShiftRight8")
	symaddr(&addr_QTextStream_OperatorShiftRight9, "QTextStream_OperatorShiftRight9")
	symaddr(&addr_QTextStream_OperatorShiftRight10, "QTextStream_OperatorShiftRight10")
	symaddr(&addr_QTextStream_OperatorShiftRight11, "QTextStream_OperatorShiftRight11")
	symaddr(&addr_QTextStream_OperatorShiftRight12, "QTextStream_OperatorShiftRight12")
	symaddr(&addr_QTextStream_OperatorShiftRight13, "QTextStream_OperatorShiftRight13")
	symaddr(&addr_QTextStream_OperatorShiftRight14, "QTextStream_OperatorShiftRight14")
	symaddr(&addr_QTextStream_OperatorShiftRight15, "QTextStream_OperatorShiftRight15")
	symaddr(&addr_QTextStream_OperatorShiftRight16, "QTextStream_OperatorShiftRight16")
	symaddr(&addr_QTextStream_OperatorShiftLeft, "QTextStream_OperatorShiftLeft")
	symaddr(&addr_QTextStream_OperatorShiftLeft2, "QTextStream_OperatorShiftLeft2")
	symaddr(&addr_QTextStream_OperatorShiftLeft3, "QTextStream_OperatorShiftLeft3")
	symaddr(&addr_QTextStream_OperatorShiftLeft4, "QTextStream_OperatorShiftLeft4")
	symaddr(&addr_QTextStream_OperatorShiftLeft5, "QTextStream_OperatorShiftLeft5")
	symaddr(&addr_QTextStream_OperatorShiftLeft6, "QTextStream_OperatorShiftLeft6")
	symaddr(&addr_QTextStream_OperatorShiftLeft7, "QTextStream_OperatorShiftLeft7")
	symaddr(&addr_QTextStream_OperatorShiftLeft8, "QTextStream_OperatorShiftLeft8")
	symaddr(&addr_QTextStream_OperatorShiftLeft9, "QTextStream_OperatorShiftLeft9")
	symaddr(&addr_QTextStream_OperatorShiftLeft10, "QTextStream_OperatorShiftLeft10")
	symaddr(&addr_QTextStream_OperatorShiftLeft11, "QTextStream_OperatorShiftLeft11")
	reg(&pg_QTextStream_OperatorShiftLeft12, "QTextStream_OperatorShiftLeft12")
	reg(&pg_QTextStream_OperatorShiftLeft13, "QTextStream_OperatorShiftLeft13")
	symaddr(&addr_QTextStream_OperatorShiftLeft14, "QTextStream_OperatorShiftLeft14")
	symaddr(&addr_QTextStream_OperatorShiftLeft15, "QTextStream_OperatorShiftLeft15")
	symaddr(&addr_QTextStream_OperatorShiftLeft16, "QTextStream_OperatorShiftLeft16")
	symaddr(&addr_QTextStream_OperatorShiftLeft17, "QTextStream_OperatorShiftLeft17")
	symaddr(&addr_QTextStream_OperatorShiftLeft18, "QTextStream_OperatorShiftLeft18")
	symaddr(&addr_QTextStream_ToBool, "QTextStream_ToBool")
	symaddr(&addr_QTextStream_setString2, "QTextStream_setString2")
	symaddr(&addr_QTextStream_readLine2, "QTextStream_readLine2")
	symaddr(&addr_QTextStream_readLineInto2, "QTextStream_readLineInto2")
	symaddr(&addr_QTextStream_OperatorShiftLeft19, "QTextStream_OperatorShiftLeft19")
	symaddr(&addr_QTextStream_OperatorShiftRight17, "QTextStream_OperatorShiftRight17")
	symaddr(&addr_QTextStream_OperatorShiftLeft20, "QTextStream_OperatorShiftLeft20")
	reg(&pg_QTextStream_delete, "QTextStream_delete")
}

type QTextStream struct {
	h uintptr
	*QIODeviceBase
}

func bindQTextStream(h uintptr) *QTextStream {
	if h == 0 {
		return nil
	}
	return &QTextStream{h: h, QIODeviceBase: bindQIODeviceBase(pg_QTextStream_base(h))}
}

func newownedQTextStream(h uintptr) *QTextStream {
	o := bindQTextStream(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextStream).Delete)
	}
	return o
}

func (this *QTextStream) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextStream) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextStream(h unsafe.Pointer) *QTextStream { return bindQTextStream(uintptr(h)) }
func UnsafeNewQTextStreamOwned(h unsafe.Pointer) *QTextStream { return newownedQTextStream(uintptr(h)) }

func (this *QTextStream) UnsafeRelease() {
	this.h = 0
	if this.QIODeviceBase != nil {
		this.QIODeviceBase.UnsafeRelease()
	}
}

func NewQTextStream() *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new)
	__o := newownedQTextStream(__r)
	return __o
}

func NewQTextStream2(device *QIODevice) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new2, uintptr(device.UnsafePointer()))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQTextStream3(string_ *QString) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new3, uintptr(string_.UnsafePointer()))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQTextStream4(array *QByteArray) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new4, uintptr(array.UnsafePointer()))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQTextStream5(array *QByteArray) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new5, uintptr(array.UnsafePointer()))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQTextStream6(string_ *QString, openMode QIODeviceBase__OpenModeFlag) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new6, uintptr(string_.UnsafePointer()), uintptr(openMode))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQTextStream7(array *QByteArray, openMode QIODeviceBase__OpenModeFlag) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new7, uintptr(array.UnsafePointer()), uintptr(openMode))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQTextStream8(array *QByteArray, openMode QIODeviceBase__OpenModeFlag) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_new8, uintptr(array.UnsafePointer()), uintptr(openMode))
	__o := newownedQTextStream(__r)
	runtime.KeepAlive(array)
	return __o
}

func (this *QTextStream) SetEncoding(encoding QStringConverter__Encoding) {
	purego.SyscallN(addr_QTextStream_setEncoding, this.h, uintptr(encoding))
	runtime.KeepAlive(this)
}

func (this *QTextStream) Encoding() QStringConverter__Encoding {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_encoding, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStringConverter__Encoding(__r)
}

func (this *QTextStream) SetAutoDetectUnicode(enabled bool) {
	purego.SyscallN(addr_QTextStream_setAutoDetectUnicode, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QTextStream) AutoDetectUnicode() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_autoDetectUnicode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextStream) SetGenerateByteOrderMark(generate bool) {
	purego.SyscallN(addr_QTextStream_setGenerateByteOrderMark, this.h, uintptr(b2i(generate)))
	runtime.KeepAlive(this)
}

func (this *QTextStream) GenerateByteOrderMark() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_generateByteOrderMark, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextStream) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QTextStream_setLocale, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QTextStream) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QTextStream_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QTextStream) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QTextStream_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QTextStream) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QTextStream_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QTextStream) SetString(string_ *QString) {
	purego.SyscallN(addr_QTextStream_setString, this.h, uintptr(string_.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
}

func (this *QTextStream) String() *QString {
	__r, _, _ := purego.SyscallN(addr_QTextStream_string, this.h)
	runtime.KeepAlive(this)
	return bindQString(__r)
}

func (this *QTextStream) Status() QTextStream__Status {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_status, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextStream__Status(__r)
}

func (this *QTextStream) SetStatus(status QTextStream__Status) {
	purego.SyscallN(addr_QTextStream_setStatus, this.h, uintptr(status))
	runtime.KeepAlive(this)
}

func (this *QTextStream) ResetStatus() {
	purego.SyscallN(addr_QTextStream_resetStatus, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextStream) AtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_atEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextStream) Reset() {
	purego.SyscallN(addr_QTextStream_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextStream) Flush() {
	purego.SyscallN(addr_QTextStream_flush, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextStream) Seek(pos int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_seek, this.h, uintptr(pos))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextStream) Pos() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_pos, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextStream) SkipWhiteSpace() {
	purego.SyscallN(addr_QTextStream_skipWhiteSpace, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextStream) ReadLine() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextStream_readLine, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextStream) ReadLineInto(line *QString) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_readLineInto, this.h, uintptr(line.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
	return __r != 0
}

func (this *QTextStream) ReadAll() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextStream_readAll, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextStream) Read(maxlen int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextStream_read, this.h, uintptr(maxlen), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextStream) SetFieldAlignment(alignment QTextStream__FieldAlignment) {
	purego.SyscallN(addr_QTextStream_setFieldAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTextStream) FieldAlignment() QTextStream__FieldAlignment {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_fieldAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextStream__FieldAlignment(__r)
}

func (this *QTextStream) SetPadChar(ch *QChar) {
	purego.SyscallN(addr_QTextStream_setPadChar, this.h, uintptr(ch.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
}

func (this *QTextStream) PadChar() *QChar {
	__r, _, _ := purego.SyscallN(addr_QTextStream_padChar, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QTextStream) SetFieldWidth(width int) {
	purego.SyscallN(addr_QTextStream_setFieldWidth, this.h, uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QTextStream) FieldWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_fieldWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextStream) SetNumberFlags(flags QTextStream__NumberFlag) {
	purego.SyscallN(addr_QTextStream_setNumberFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextStream) NumberFlags() QTextStream__NumberFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_numberFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextStream__NumberFlag(__r)
}

func (this *QTextStream) SetIntegerBase(base int) {
	purego.SyscallN(addr_QTextStream_setIntegerBase, this.h, uintptr(base))
	runtime.KeepAlive(this)
}

func (this *QTextStream) IntegerBase() int {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_integerBase, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextStream) SetRealNumberNotation(notation QTextStream__RealNumberNotation) {
	purego.SyscallN(addr_QTextStream_setRealNumberNotation, this.h, uintptr(notation))
	runtime.KeepAlive(this)
}

func (this *QTextStream) RealNumberNotation() QTextStream__RealNumberNotation {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_realNumberNotation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextStream__RealNumberNotation(__r)
}

func (this *QTextStream) SetRealNumberPrecision(precision int) {
	purego.SyscallN(addr_QTextStream_setRealNumberPrecision, this.h, uintptr(precision))
	runtime.KeepAlive(this)
}

func (this *QTextStream) RealNumberPrecision() int {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_realNumberPrecision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextStream) OperatorShiftRight(ch *QChar) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight, this.h, uintptr(ch.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftRightWithCh(ch int8) (*QTextStream, int8) {
	ch_io := ch
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight2, this.h, uintptr(unsafe.Pointer(&ch_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), ch_io
}

func (this *QTextStream) OperatorShiftRight2(ch uint16) (*QTextStream, uint16) {
	ch_io := int16(ch)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight3, this.h, uintptr(unsafe.Pointer(&ch_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), uint16(ch_io)
}

func (this *QTextStream) OperatorShiftRightWithShort(i int16) (*QTextStream, int16) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight4, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), i_io
}

func (this *QTextStream) OperatorShiftRightWithUnsignedshort(i uint16) (*QTextStream, uint16) {
	i_io := int16(i)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight5, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), uint16(i_io)
}

func (this *QTextStream) OperatorShiftRightWithInt(i int) (*QTextStream, int) {
	i_io := int32(i)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight6, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), int(i_io)
}

func (this *QTextStream) OperatorShiftRightWithUnsignedint(i uint) (*QTextStream, uint) {
	i_io := int32(i)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight7, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), uint(i_io)
}

func (this *QTextStream) OperatorShiftRightWithLong(i int32) (*QTextStream, int32) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight8, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), i_io
}

func (this *QTextStream) OperatorShiftRightWithUnsignedlong(i uint32) (*QTextStream, uint32) {
	i_io := int32(i)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight9, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), uint32(i_io)
}

func (this *QTextStream) OperatorShiftRightWith(i int64) (*QTextStream, int64) {
	i_io := i
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight10, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), i_io
}

func (this *QTextStream) OperatorShiftRight3(i uint64) (*QTextStream, uint64) {
	i_io := int64(i)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight11, this.h, uintptr(unsafe.Pointer(&i_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), uint64(i_io)
}

func (this *QTextStream) OperatorShiftRightWithFloat(f float32) (*QTextStream, float32) {
	f_io := f
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight12, this.h, uintptr(unsafe.Pointer(&f_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), f_io
}

func (this *QTextStream) OperatorShiftRightWithDouble(f float64) (*QTextStream, float64) {
	f_io := f
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight13, this.h, uintptr(unsafe.Pointer(&f_io)))
	runtime.KeepAlive(this)
	return bindQTextStream(__r), f_io
}

func (this *QTextStream) OperatorShiftRightWithQString(s *QString) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight14, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftRightWithArray(array *QByteArray) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight15, this.h, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(array)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftRightWithChar(c []byte) *QTextStream {
	var c_p uintptr
	if len(c) > 0 {
		c_p = uintptr(unsafe.Pointer(&c[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight16, this.h, c_p)
	runtime.KeepAlive(this)
	runtime.KeepAlive(c)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeft(ch *QChar) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft, this.h, uintptr(ch.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ch)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithCh(ch int8) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft2, this.h, uintptr(ch))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeft2(ch uint16) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft3, this.h, uintptr(ch))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithShort(i int16) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft4, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedshort(i uint16) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft5, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithInt(i int) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft6, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedint(i uint) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft7, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithLong(i int32) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft8, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithUnsignedlong(i uint32) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft9, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWith(i int64) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft10, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeft3(i uint64) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft11, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithFloat(f float32) *QTextStream {
	__r := pg_QTextStream_OperatorShiftLeft12(this.h, f)
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithDouble(f float64) *QTextStream {
	__r := pg_QTextStream_OperatorShiftLeft13(this.h, f)
	runtime.KeepAlive(this)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithQString(s string) *QTextStream {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft14, this.h, s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithQStringView(s *QStringView) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft15, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeft4(s *QLatin1String) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft16, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithArray(array *QByteArray) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft17, this.h, uintptr(array.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(array)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithChar(c string) *QTextStream {
	c_b := append([]byte(c), 0)
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft18, this.h, uintptr(unsafe.Pointer(&c_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c_b)
	return bindQTextStream(__r)
}

func (this *QTextStream) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_ToBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextStream) SetString2(string_ *QString, openMode QIODeviceBase__OpenModeFlag) {
	purego.SyscallN(addr_QTextStream_setString2, this.h, uintptr(string_.UnsafePointer()), uintptr(openMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
}

func (this *QTextStream) ReadLineWithMaxlen(maxlen int64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextStream_readLine2, this.h, uintptr(maxlen), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextStream) ReadLineInto2(line *QString, maxlen int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextStream_readLineInto2, this.h, uintptr(line.UnsafePointer()), uintptr(maxlen))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
	return __r != 0
}

func (this *QTextStream) OperatorShiftLeftWithQfloat16(f *qfloat16) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft19, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftRightWithQfloat16(f *qfloat16) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftRight17, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
	return bindQTextStream(__r)
}

func (this *QTextStream) OperatorShiftLeftWithQTextStreamManipulator(m *QTextStreamManipulator) *QTextStream {
	__r, _, _ := purego.SyscallN(addr_QTextStream_OperatorShiftLeft20, this.h, uintptr(m.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(m)
	return bindQTextStream(__r)
}

func (this *QTextStream) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextStream_delete(this.h)
	}
	this.UnsafeRelease()
}
