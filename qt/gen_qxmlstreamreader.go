package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamReader_new uintptr
	addr_QXmlStreamReader_new2 uintptr
	addr_QXmlStreamReader_new3 uintptr
	addr_QXmlStreamReader_setDevice uintptr
	addr_QXmlStreamReader_device uintptr
	addr_QXmlStreamReader_addData uintptr
	addr_QXmlStreamReader_clear uintptr
	addr_QXmlStreamReader_atEnd uintptr
	addr_QXmlStreamReader_readNext uintptr
	addr_QXmlStreamReader_readNextStartElement uintptr
	addr_QXmlStreamReader_skipCurrentElement uintptr
	addr_QXmlStreamReader_readRawInnerData uintptr
	addr_QXmlStreamReader_tokenType uintptr
	addr_QXmlStreamReader_tokenString uintptr
	addr_QXmlStreamReader_setNamespaceProcessing uintptr
	addr_QXmlStreamReader_namespaceProcessing uintptr
	addr_QXmlStreamReader_isStartDocument uintptr
	addr_QXmlStreamReader_isEndDocument uintptr
	addr_QXmlStreamReader_isStartElement uintptr
	addr_QXmlStreamReader_isEndElement uintptr
	addr_QXmlStreamReader_isCharacters uintptr
	addr_QXmlStreamReader_isWhitespace uintptr
	addr_QXmlStreamReader_isCDATA uintptr
	addr_QXmlStreamReader_isComment uintptr
	addr_QXmlStreamReader_isDTD uintptr
	addr_QXmlStreamReader_isEntityReference uintptr
	addr_QXmlStreamReader_isProcessingInstruction uintptr
	addr_QXmlStreamReader_isStandaloneDocument uintptr
	addr_QXmlStreamReader_hasStandaloneDeclaration uintptr
	addr_QXmlStreamReader_documentVersion uintptr
	addr_QXmlStreamReader_documentEncoding uintptr
	addr_QXmlStreamReader_lineNumber uintptr
	addr_QXmlStreamReader_columnNumber uintptr
	addr_QXmlStreamReader_characterOffset uintptr
	addr_QXmlStreamReader_attributes uintptr
	addr_QXmlStreamReader_readElementText uintptr
	addr_QXmlStreamReader_name uintptr
	addr_QXmlStreamReader_namespaceUri uintptr
	addr_QXmlStreamReader_qualifiedName uintptr
	addr_QXmlStreamReader_prefix uintptr
	addr_QXmlStreamReader_processingInstructionTarget uintptr
	addr_QXmlStreamReader_processingInstructionData uintptr
	addr_QXmlStreamReader_text uintptr
	addr_QXmlStreamReader_namespaceDeclarations uintptr
	addr_QXmlStreamReader_addExtraNamespaceDeclaration uintptr
	addr_QXmlStreamReader_addExtraNamespaceDeclarations uintptr
	addr_QXmlStreamReader_notationDeclarations uintptr
	addr_QXmlStreamReader_entityDeclarations uintptr
	addr_QXmlStreamReader_dtdName uintptr
	addr_QXmlStreamReader_dtdPublicId uintptr
	addr_QXmlStreamReader_dtdSystemId uintptr
	addr_QXmlStreamReader_entityExpansionLimit uintptr
	addr_QXmlStreamReader_setEntityExpansionLimit uintptr
	addr_QXmlStreamReader_raiseError uintptr
	addr_QXmlStreamReader_errorString uintptr
	addr_QXmlStreamReader_error uintptr
	addr_QXmlStreamReader_hasError uintptr
	addr_QXmlStreamReader_setEntityResolver uintptr
	addr_QXmlStreamReader_entityResolver uintptr
	addr_QXmlStreamReader_readElementText2 uintptr
	addr_QXmlStreamReader_raiseError2 uintptr
	pg_QXmlStreamReader_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamReader_new, "QXmlStreamReader_new")
	symaddr(&addr_QXmlStreamReader_new2, "QXmlStreamReader_new2")
	symaddr(&addr_QXmlStreamReader_new3, "QXmlStreamReader_new3")
	symaddr(&addr_QXmlStreamReader_setDevice, "QXmlStreamReader_setDevice")
	symaddr(&addr_QXmlStreamReader_device, "QXmlStreamReader_device")
	symaddr(&addr_QXmlStreamReader_addData, "QXmlStreamReader_addData")
	symaddr(&addr_QXmlStreamReader_clear, "QXmlStreamReader_clear")
	symaddr(&addr_QXmlStreamReader_atEnd, "QXmlStreamReader_atEnd")
	symaddr(&addr_QXmlStreamReader_readNext, "QXmlStreamReader_readNext")
	symaddr(&addr_QXmlStreamReader_readNextStartElement, "QXmlStreamReader_readNextStartElement")
	symaddr(&addr_QXmlStreamReader_skipCurrentElement, "QXmlStreamReader_skipCurrentElement")
	symaddr(&addr_QXmlStreamReader_readRawInnerData, "QXmlStreamReader_readRawInnerData")
	symaddr(&addr_QXmlStreamReader_tokenType, "QXmlStreamReader_tokenType")
	symaddr(&addr_QXmlStreamReader_tokenString, "QXmlStreamReader_tokenString")
	symaddr(&addr_QXmlStreamReader_setNamespaceProcessing, "QXmlStreamReader_setNamespaceProcessing")
	symaddr(&addr_QXmlStreamReader_namespaceProcessing, "QXmlStreamReader_namespaceProcessing")
	symaddr(&addr_QXmlStreamReader_isStartDocument, "QXmlStreamReader_isStartDocument")
	symaddr(&addr_QXmlStreamReader_isEndDocument, "QXmlStreamReader_isEndDocument")
	symaddr(&addr_QXmlStreamReader_isStartElement, "QXmlStreamReader_isStartElement")
	symaddr(&addr_QXmlStreamReader_isEndElement, "QXmlStreamReader_isEndElement")
	symaddr(&addr_QXmlStreamReader_isCharacters, "QXmlStreamReader_isCharacters")
	symaddr(&addr_QXmlStreamReader_isWhitespace, "QXmlStreamReader_isWhitespace")
	symaddr(&addr_QXmlStreamReader_isCDATA, "QXmlStreamReader_isCDATA")
	symaddr(&addr_QXmlStreamReader_isComment, "QXmlStreamReader_isComment")
	symaddr(&addr_QXmlStreamReader_isDTD, "QXmlStreamReader_isDTD")
	symaddr(&addr_QXmlStreamReader_isEntityReference, "QXmlStreamReader_isEntityReference")
	symaddr(&addr_QXmlStreamReader_isProcessingInstruction, "QXmlStreamReader_isProcessingInstruction")
	symaddr(&addr_QXmlStreamReader_isStandaloneDocument, "QXmlStreamReader_isStandaloneDocument")
	symaddr(&addr_QXmlStreamReader_hasStandaloneDeclaration, "QXmlStreamReader_hasStandaloneDeclaration")
	symaddr(&addr_QXmlStreamReader_documentVersion, "QXmlStreamReader_documentVersion")
	symaddr(&addr_QXmlStreamReader_documentEncoding, "QXmlStreamReader_documentEncoding")
	symaddr(&addr_QXmlStreamReader_lineNumber, "QXmlStreamReader_lineNumber")
	symaddr(&addr_QXmlStreamReader_columnNumber, "QXmlStreamReader_columnNumber")
	symaddr(&addr_QXmlStreamReader_characterOffset, "QXmlStreamReader_characterOffset")
	symaddr(&addr_QXmlStreamReader_attributes, "QXmlStreamReader_attributes")
	symaddr(&addr_QXmlStreamReader_readElementText, "QXmlStreamReader_readElementText")
	symaddr(&addr_QXmlStreamReader_name, "QXmlStreamReader_name")
	symaddr(&addr_QXmlStreamReader_namespaceUri, "QXmlStreamReader_namespaceUri")
	symaddr(&addr_QXmlStreamReader_qualifiedName, "QXmlStreamReader_qualifiedName")
	symaddr(&addr_QXmlStreamReader_prefix, "QXmlStreamReader_prefix")
	symaddr(&addr_QXmlStreamReader_processingInstructionTarget, "QXmlStreamReader_processingInstructionTarget")
	symaddr(&addr_QXmlStreamReader_processingInstructionData, "QXmlStreamReader_processingInstructionData")
	symaddr(&addr_QXmlStreamReader_text, "QXmlStreamReader_text")
	symaddr(&addr_QXmlStreamReader_namespaceDeclarations, "QXmlStreamReader_namespaceDeclarations")
	symaddr(&addr_QXmlStreamReader_addExtraNamespaceDeclaration, "QXmlStreamReader_addExtraNamespaceDeclaration")
	symaddr(&addr_QXmlStreamReader_addExtraNamespaceDeclarations, "QXmlStreamReader_addExtraNamespaceDeclarations")
	symaddr(&addr_QXmlStreamReader_notationDeclarations, "QXmlStreamReader_notationDeclarations")
	symaddr(&addr_QXmlStreamReader_entityDeclarations, "QXmlStreamReader_entityDeclarations")
	symaddr(&addr_QXmlStreamReader_dtdName, "QXmlStreamReader_dtdName")
	symaddr(&addr_QXmlStreamReader_dtdPublicId, "QXmlStreamReader_dtdPublicId")
	symaddr(&addr_QXmlStreamReader_dtdSystemId, "QXmlStreamReader_dtdSystemId")
	symaddr(&addr_QXmlStreamReader_entityExpansionLimit, "QXmlStreamReader_entityExpansionLimit")
	symaddr(&addr_QXmlStreamReader_setEntityExpansionLimit, "QXmlStreamReader_setEntityExpansionLimit")
	symaddr(&addr_QXmlStreamReader_raiseError, "QXmlStreamReader_raiseError")
	symaddr(&addr_QXmlStreamReader_errorString, "QXmlStreamReader_errorString")
	symaddr(&addr_QXmlStreamReader_error, "QXmlStreamReader_error")
	symaddr(&addr_QXmlStreamReader_hasError, "QXmlStreamReader_hasError")
	symaddr(&addr_QXmlStreamReader_setEntityResolver, "QXmlStreamReader_setEntityResolver")
	symaddr(&addr_QXmlStreamReader_entityResolver, "QXmlStreamReader_entityResolver")
	symaddr(&addr_QXmlStreamReader_readElementText2, "QXmlStreamReader_readElementText2")
	symaddr(&addr_QXmlStreamReader_raiseError2, "QXmlStreamReader_raiseError2")
	reg(&pg_QXmlStreamReader_delete, "QXmlStreamReader_delete")
}

type QXmlStreamReader struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamReader) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamReader) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamReader(h uintptr) *QXmlStreamReader {
	if h == 0 {
		return nil
	}
	return &QXmlStreamReader{h: h}
}

func newownedQXmlStreamReader(h uintptr) *QXmlStreamReader {
	o := bindQXmlStreamReader(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamReader).Delete)
	}
	return o
}

func (this *QXmlStreamReader) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamReader) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamReader(h unsafe.Pointer) *QXmlStreamReader { return bindQXmlStreamReader(uintptr(h)) }
func UnsafeNewQXmlStreamReaderOwned(h unsafe.Pointer) *QXmlStreamReader { return newownedQXmlStreamReader(uintptr(h)) }

func (this *QXmlStreamReader) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamReader() *QXmlStreamReader {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_new)
	__o := newownedQXmlStreamReader(__r)
	return __o
}

func NewQXmlStreamReader2(device *QIODevice) *QXmlStreamReader {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_new2, uintptr(device.UnsafePointer()))
	__o := newownedQXmlStreamReader(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQXmlStreamReader3(data string) *QXmlStreamReader {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(unsafe.StringData(data)))
	}
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_new3, data_p, uintptr(len(data)))
	__o := newownedQXmlStreamReader(__r)
	runtime.KeepAlive(data)
	return __o
}

func (this *QXmlStreamReader) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QXmlStreamReader_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QXmlStreamReader) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QXmlStreamReader) AddData(data string) {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(unsafe.StringData(data)))
	}
	purego.SyscallN(addr_QXmlStreamReader_addData, this.h, data_p, uintptr(len(data)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QXmlStreamReader) Clear() {
	purego.SyscallN(addr_QXmlStreamReader_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamReader) AtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_atEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) ReadNext() QXmlStreamReader__TokenType {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_readNext, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QXmlStreamReader__TokenType(__r)
}

func (this *QXmlStreamReader) ReadNextStartElement() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_readNextStartElement, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) SkipCurrentElement() {
	purego.SyscallN(addr_QXmlStreamReader_skipCurrentElement, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamReader) ReadRawInnerData() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_readRawInnerData, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamReader) TokenType() QXmlStreamReader__TokenType {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_tokenType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QXmlStreamReader__TokenType(__r)
}

func (this *QXmlStreamReader) TokenString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_tokenString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamReader) SetNamespaceProcessing(arg1 bool) {
	purego.SyscallN(addr_QXmlStreamReader_setNamespaceProcessing, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamReader) NamespaceProcessing() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_namespaceProcessing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsStartDocument() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isStartDocument, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsEndDocument() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isEndDocument, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsStartElement() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isStartElement, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsEndElement() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isEndElement, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsCharacters() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isCharacters, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsWhitespace() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isWhitespace, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsCDATA() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isCDATA, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsComment() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isComment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsDTD() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isDTD, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsEntityReference() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isEntityReference, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsProcessingInstruction() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isProcessingInstruction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) IsStandaloneDocument() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_isStandaloneDocument, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) HasStandaloneDeclaration() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_hasStandaloneDeclaration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) DocumentVersion() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_documentVersion, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) DocumentEncoding() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_documentEncoding, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) LineNumber() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_lineNumber, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamReader) ColumnNumber() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_columnNumber, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamReader) CharacterOffset() int64 {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_characterOffset, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QXmlStreamReader) Attributes() *QXmlStreamAttributes {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_attributes, this.h)
	runtime.KeepAlive(this)
	return newownedQXmlStreamAttributes(__r)
}

func (this *QXmlStreamReader) ReadElementText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_readElementText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamReader) Name() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_name, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) NamespaceUri() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_namespaceUri, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) QualifiedName() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_qualifiedName, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) Prefix() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_prefix, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) ProcessingInstructionTarget() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_processingInstructionTarget, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) ProcessingInstructionData() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_processingInstructionData, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) Text() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_text, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) NamespaceDeclarations() []*QXmlStreamNamespaceDeclaration {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_namespaceDeclarations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamNamespaceDeclaration, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamNamespaceDeclaration(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamReader) AddExtraNamespaceDeclaration(extraNamespaceDeclaraction *QXmlStreamNamespaceDeclaration) {
	purego.SyscallN(addr_QXmlStreamReader_addExtraNamespaceDeclaration, this.h, uintptr(extraNamespaceDeclaraction.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(extraNamespaceDeclaraction)
}

func (this *QXmlStreamReader) AddExtraNamespaceDeclarations(extraNamespaceDeclaractions []*QXmlStreamNamespaceDeclaration) {
	extraNamespaceDeclaractions_arr := make([]uintptr, len(extraNamespaceDeclaractions))
	for __i, __e := range extraNamespaceDeclaractions {
		extraNamespaceDeclaractions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var extraNamespaceDeclaractions_p uintptr
	if len(extraNamespaceDeclaractions_arr) > 0 {
		extraNamespaceDeclaractions_p = uintptr(unsafe.Pointer(&extraNamespaceDeclaractions_arr[0]))
	}
	purego.SyscallN(addr_QXmlStreamReader_addExtraNamespaceDeclarations, this.h, extraNamespaceDeclaractions_p, uintptr(len(extraNamespaceDeclaractions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(extraNamespaceDeclaractions)
	runtime.KeepAlive(extraNamespaceDeclaractions_arr)
}

func (this *QXmlStreamReader) NotationDeclarations() []*QXmlStreamNotationDeclaration {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_notationDeclarations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamNotationDeclaration, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamNotationDeclaration(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamReader) EntityDeclarations() []*QXmlStreamEntityDeclaration {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_entityDeclarations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QXmlStreamEntityDeclaration, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQXmlStreamEntityDeclaration(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QXmlStreamReader) DtdName() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_dtdName, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) DtdPublicId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_dtdPublicId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) DtdSystemId() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_dtdSystemId, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QXmlStreamReader) EntityExpansionLimit() int {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_entityExpansionLimit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QXmlStreamReader) SetEntityExpansionLimit(limit int) {
	purego.SyscallN(addr_QXmlStreamReader_setEntityExpansionLimit, this.h, uintptr(limit))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamReader) RaiseError() {
	purego.SyscallN(addr_QXmlStreamReader_raiseError, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamReader) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamReader) Error() QXmlStreamReader__Error {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QXmlStreamReader__Error(__r)
}

func (this *QXmlStreamReader) HasError() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamReader_hasError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamReader) SetEntityResolver(resolver *QXmlStreamEntityResolver) {
	purego.SyscallN(addr_QXmlStreamReader_setEntityResolver, this.h, uintptr(resolver.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(resolver)
}

func (this *QXmlStreamReader) EntityResolver() *QXmlStreamEntityResolver {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_entityResolver, this.h)
	runtime.KeepAlive(this)
	return bindQXmlStreamEntityResolver(__r)
}

func (this *QXmlStreamReader) ReadElementTextWithBehaviour(behaviour QXmlStreamReader__ReadElementTextBehaviour) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamReader_readElementText2, this.h, uintptr(behaviour), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamReader) RaiseErrorWithMessage(message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QXmlStreamReader_raiseError2, this.h, message_p, uintptr(len(message)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func (this *QXmlStreamReader) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamReader_delete(this.h)
	}
	this.UnsafeRelease()
}
