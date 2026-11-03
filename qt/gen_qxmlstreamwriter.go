package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QXmlStreamWriter_new uintptr
	addr_QXmlStreamWriter_new2 uintptr
	addr_QXmlStreamWriter_new3 uintptr
	addr_QXmlStreamWriter_new4 uintptr
	addr_QXmlStreamWriter_setDevice uintptr
	addr_QXmlStreamWriter_device uintptr
	addr_QXmlStreamWriter_setAutoFormatting uintptr
	addr_QXmlStreamWriter_autoFormatting uintptr
	addr_QXmlStreamWriter_setAutoFormattingIndent uintptr
	addr_QXmlStreamWriter_autoFormattingIndent uintptr
	addr_QXmlStreamWriter_setStopWritingOnError uintptr
	addr_QXmlStreamWriter_stopWritingOnError uintptr
	addr_QXmlStreamWriter_writeAttribute uintptr
	addr_QXmlStreamWriter_writeAttribute2 uintptr
	addr_QXmlStreamWriter_writeAttribute3 uintptr
	addr_QXmlStreamWriter_writeAttributes uintptr
	addr_QXmlStreamWriter_writeCDATA uintptr
	addr_QXmlStreamWriter_writeCharacters uintptr
	addr_QXmlStreamWriter_writeComment uintptr
	addr_QXmlStreamWriter_writeDTD uintptr
	addr_QXmlStreamWriter_writeEmptyElement uintptr
	addr_QXmlStreamWriter_writeEmptyElement2 uintptr
	addr_QXmlStreamWriter_writeTextElement uintptr
	addr_QXmlStreamWriter_writeTextElement2 uintptr
	addr_QXmlStreamWriter_writeEndDocument uintptr
	addr_QXmlStreamWriter_writeEndElement uintptr
	addr_QXmlStreamWriter_writeEntityReference uintptr
	addr_QXmlStreamWriter_writeNamespace uintptr
	addr_QXmlStreamWriter_writeDefaultNamespace uintptr
	addr_QXmlStreamWriter_writeProcessingInstruction uintptr
	addr_QXmlStreamWriter_writeStartDocument uintptr
	addr_QXmlStreamWriter_writeStartDocument2 uintptr
	addr_QXmlStreamWriter_writeStartDocument3 uintptr
	addr_QXmlStreamWriter_writeStartElement uintptr
	addr_QXmlStreamWriter_writeStartElement2 uintptr
	addr_QXmlStreamWriter_writeCurrentToken uintptr
	addr_QXmlStreamWriter_raiseError uintptr
	addr_QXmlStreamWriter_errorString uintptr
	addr_QXmlStreamWriter_error uintptr
	addr_QXmlStreamWriter_hasError uintptr
	addr_QXmlStreamWriter_writeNamespace2 uintptr
	addr_QXmlStreamWriter_writeProcessingInstruction2 uintptr
	pg_QXmlStreamWriter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QXmlStreamWriter_new, "QXmlStreamWriter_new")
	symaddr(&addr_QXmlStreamWriter_new2, "QXmlStreamWriter_new2")
	symaddr(&addr_QXmlStreamWriter_new3, "QXmlStreamWriter_new3")
	symaddr(&addr_QXmlStreamWriter_new4, "QXmlStreamWriter_new4")
	symaddr(&addr_QXmlStreamWriter_setDevice, "QXmlStreamWriter_setDevice")
	symaddr(&addr_QXmlStreamWriter_device, "QXmlStreamWriter_device")
	symaddr(&addr_QXmlStreamWriter_setAutoFormatting, "QXmlStreamWriter_setAutoFormatting")
	symaddr(&addr_QXmlStreamWriter_autoFormatting, "QXmlStreamWriter_autoFormatting")
	symaddr(&addr_QXmlStreamWriter_setAutoFormattingIndent, "QXmlStreamWriter_setAutoFormattingIndent")
	symaddr(&addr_QXmlStreamWriter_autoFormattingIndent, "QXmlStreamWriter_autoFormattingIndent")
	symaddr(&addr_QXmlStreamWriter_setStopWritingOnError, "QXmlStreamWriter_setStopWritingOnError")
	symaddr(&addr_QXmlStreamWriter_stopWritingOnError, "QXmlStreamWriter_stopWritingOnError")
	symaddr(&addr_QXmlStreamWriter_writeAttribute, "QXmlStreamWriter_writeAttribute")
	symaddr(&addr_QXmlStreamWriter_writeAttribute2, "QXmlStreamWriter_writeAttribute2")
	symaddr(&addr_QXmlStreamWriter_writeAttribute3, "QXmlStreamWriter_writeAttribute3")
	symaddr(&addr_QXmlStreamWriter_writeAttributes, "QXmlStreamWriter_writeAttributes")
	symaddr(&addr_QXmlStreamWriter_writeCDATA, "QXmlStreamWriter_writeCDATA")
	symaddr(&addr_QXmlStreamWriter_writeCharacters, "QXmlStreamWriter_writeCharacters")
	symaddr(&addr_QXmlStreamWriter_writeComment, "QXmlStreamWriter_writeComment")
	symaddr(&addr_QXmlStreamWriter_writeDTD, "QXmlStreamWriter_writeDTD")
	symaddr(&addr_QXmlStreamWriter_writeEmptyElement, "QXmlStreamWriter_writeEmptyElement")
	symaddr(&addr_QXmlStreamWriter_writeEmptyElement2, "QXmlStreamWriter_writeEmptyElement2")
	symaddr(&addr_QXmlStreamWriter_writeTextElement, "QXmlStreamWriter_writeTextElement")
	symaddr(&addr_QXmlStreamWriter_writeTextElement2, "QXmlStreamWriter_writeTextElement2")
	symaddr(&addr_QXmlStreamWriter_writeEndDocument, "QXmlStreamWriter_writeEndDocument")
	symaddr(&addr_QXmlStreamWriter_writeEndElement, "QXmlStreamWriter_writeEndElement")
	symaddr(&addr_QXmlStreamWriter_writeEntityReference, "QXmlStreamWriter_writeEntityReference")
	symaddr(&addr_QXmlStreamWriter_writeNamespace, "QXmlStreamWriter_writeNamespace")
	symaddr(&addr_QXmlStreamWriter_writeDefaultNamespace, "QXmlStreamWriter_writeDefaultNamespace")
	symaddr(&addr_QXmlStreamWriter_writeProcessingInstruction, "QXmlStreamWriter_writeProcessingInstruction")
	symaddr(&addr_QXmlStreamWriter_writeStartDocument, "QXmlStreamWriter_writeStartDocument")
	symaddr(&addr_QXmlStreamWriter_writeStartDocument2, "QXmlStreamWriter_writeStartDocument2")
	symaddr(&addr_QXmlStreamWriter_writeStartDocument3, "QXmlStreamWriter_writeStartDocument3")
	symaddr(&addr_QXmlStreamWriter_writeStartElement, "QXmlStreamWriter_writeStartElement")
	symaddr(&addr_QXmlStreamWriter_writeStartElement2, "QXmlStreamWriter_writeStartElement2")
	symaddr(&addr_QXmlStreamWriter_writeCurrentToken, "QXmlStreamWriter_writeCurrentToken")
	symaddr(&addr_QXmlStreamWriter_raiseError, "QXmlStreamWriter_raiseError")
	symaddr(&addr_QXmlStreamWriter_errorString, "QXmlStreamWriter_errorString")
	symaddr(&addr_QXmlStreamWriter_error, "QXmlStreamWriter_error")
	symaddr(&addr_QXmlStreamWriter_hasError, "QXmlStreamWriter_hasError")
	symaddr(&addr_QXmlStreamWriter_writeNamespace2, "QXmlStreamWriter_writeNamespace2")
	symaddr(&addr_QXmlStreamWriter_writeProcessingInstruction2, "QXmlStreamWriter_writeProcessingInstruction2")
	reg(&pg_QXmlStreamWriter_delete, "QXmlStreamWriter_delete")
}

type QXmlStreamWriter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QXmlStreamWriter) UnsafeGoID() int64     { return this.vid }
func (this *QXmlStreamWriter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQXmlStreamWriter(h uintptr) *QXmlStreamWriter {
	if h == 0 {
		return nil
	}
	return &QXmlStreamWriter{h: h}
}

func newownedQXmlStreamWriter(h uintptr) *QXmlStreamWriter {
	o := bindQXmlStreamWriter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QXmlStreamWriter).Delete)
	}
	return o
}

func (this *QXmlStreamWriter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QXmlStreamWriter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQXmlStreamWriter(h unsafe.Pointer) *QXmlStreamWriter { return bindQXmlStreamWriter(uintptr(h)) }
func UnsafeNewQXmlStreamWriterOwned(h unsafe.Pointer) *QXmlStreamWriter { return newownedQXmlStreamWriter(uintptr(h)) }

func (this *QXmlStreamWriter) UnsafeRelease() { this.h = 0 }

func NewQXmlStreamWriter() *QXmlStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_new)
	__o := newownedQXmlStreamWriter(__r)
	return __o
}

func NewQXmlStreamWriter2(device *QIODevice) *QXmlStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_new2, uintptr(device.UnsafePointer()))
	__o := newownedQXmlStreamWriter(__r)
	runtime.KeepAlive(device)
	return __o
}

func NewQXmlStreamWriter3(array *QByteArray) *QXmlStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_new3, uintptr(array.UnsafePointer()))
	__o := newownedQXmlStreamWriter(__r)
	runtime.KeepAlive(array)
	return __o
}

func NewQXmlStreamWriter4(string_ *QString) *QXmlStreamWriter {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_new4, uintptr(string_.UnsafePointer()))
	__o := newownedQXmlStreamWriter(__r)
	runtime.KeepAlive(string_)
	return __o
}

func (this *QXmlStreamWriter) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QXmlStreamWriter_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QXmlStreamWriter) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QXmlStreamWriter) SetAutoFormatting(arg1 bool) {
	purego.SyscallN(addr_QXmlStreamWriter_setAutoFormatting, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) AutoFormatting() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamWriter_autoFormatting, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamWriter) SetAutoFormattingIndent(spacesOrTabs int) {
	purego.SyscallN(addr_QXmlStreamWriter_setAutoFormattingIndent, this.h, uintptr(spacesOrTabs))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) AutoFormattingIndent() int {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamWriter_autoFormattingIndent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QXmlStreamWriter) SetStopWritingOnError(stop bool) {
	purego.SyscallN(addr_QXmlStreamWriter_setStopWritingOnError, this.h, uintptr(b2i(stop)))
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) StopWritingOnError() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamWriter_stopWritingOnError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamWriter) WriteAttribute(qualifiedName string, value string) {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeAttribute, this.h, qualifiedName_p, uintptr(len(qualifiedName)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
	runtime.KeepAlive(value)
}

func (this *QXmlStreamWriter) WriteAttribute2(namespaceUri string, name string, value string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeAttribute2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

func (this *QXmlStreamWriter) WriteAttributeWithAttribute(attribute *QXmlStreamAttribute) {
	purego.SyscallN(addr_QXmlStreamWriter_writeAttribute3, this.h, uintptr(attribute.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(attribute)
}

func (this *QXmlStreamWriter) WriteAttributes(attributes *QXmlStreamAttributes) {
	purego.SyscallN(addr_QXmlStreamWriter_writeAttributes, this.h, uintptr(attributes.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(attributes)
}

func (this *QXmlStreamWriter) WriteCDATA(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeCDATA, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QXmlStreamWriter) WriteCharacters(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeCharacters, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QXmlStreamWriter) WriteComment(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeComment, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QXmlStreamWriter) WriteDTD(dtd string) {
	var dtd_p uintptr
	if len(dtd) > 0 {
		dtd_p = uintptr(unsafe.Pointer(unsafe.StringData(dtd)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeDTD, this.h, dtd_p, uintptr(len(dtd)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dtd)
}

func (this *QXmlStreamWriter) WriteEmptyElement(qualifiedName string) {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeEmptyElement, this.h, qualifiedName_p, uintptr(len(qualifiedName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
}

func (this *QXmlStreamWriter) WriteEmptyElement2(namespaceUri string, name string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeEmptyElement2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
}

func (this *QXmlStreamWriter) WriteTextElement(qualifiedName string, text string) {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeTextElement, this.h, qualifiedName_p, uintptr(len(qualifiedName)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
	runtime.KeepAlive(text)
}

func (this *QXmlStreamWriter) WriteTextElement2(namespaceUri string, name string, text string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeTextElement2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
	runtime.KeepAlive(text)
}

func (this *QXmlStreamWriter) WriteEndDocument() {
	purego.SyscallN(addr_QXmlStreamWriter_writeEndDocument, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) WriteEndElement() {
	purego.SyscallN(addr_QXmlStreamWriter_writeEndElement, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) WriteEntityReference(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeEntityReference, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QXmlStreamWriter) WriteNamespace(namespaceUri string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeNamespace, this.h, namespaceUri_p, uintptr(len(namespaceUri)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
}

func (this *QXmlStreamWriter) WriteDefaultNamespace(namespaceUri string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeDefaultNamespace, this.h, namespaceUri_p, uintptr(len(namespaceUri)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
}

func (this *QXmlStreamWriter) WriteProcessingInstruction(target string) {
	var target_p uintptr
	if len(target) > 0 {
		target_p = uintptr(unsafe.Pointer(unsafe.StringData(target)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeProcessingInstruction, this.h, target_p, uintptr(len(target)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
}

func (this *QXmlStreamWriter) WriteStartDocument() {
	purego.SyscallN(addr_QXmlStreamWriter_writeStartDocument, this.h)
	runtime.KeepAlive(this)
}

func (this *QXmlStreamWriter) WriteStartDocumentWithVersion(version string) {
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeStartDocument2, this.h, version_p, uintptr(len(version)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(version)
}

func (this *QXmlStreamWriter) WriteStartDocument2(version string, standalone bool) {
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeStartDocument3, this.h, version_p, uintptr(len(version)), uintptr(b2i(standalone)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(version)
}

func (this *QXmlStreamWriter) WriteStartElement(qualifiedName string) {
	var qualifiedName_p uintptr
	if len(qualifiedName) > 0 {
		qualifiedName_p = uintptr(unsafe.Pointer(unsafe.StringData(qualifiedName)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeStartElement, this.h, qualifiedName_p, uintptr(len(qualifiedName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(qualifiedName)
}

func (this *QXmlStreamWriter) WriteStartElement2(namespaceUri string, name string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeStartElement2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(name)
}

func (this *QXmlStreamWriter) WriteCurrentToken(reader *QXmlStreamReader) {
	purego.SyscallN(addr_QXmlStreamWriter_writeCurrentToken, this.h, uintptr(reader.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(reader)
}

func (this *QXmlStreamWriter) RaiseError(message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_raiseError, this.h, message_p, uintptr(len(message)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(message)
}

func (this *QXmlStreamWriter) ErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QXmlStreamWriter_errorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QXmlStreamWriter) Error() QXmlStreamWriter__Error {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamWriter_error, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QXmlStreamWriter__Error(__r)
}

func (this *QXmlStreamWriter) HasError() bool {
	__r0, _, _ := purego.SyscallN(addr_QXmlStreamWriter_hasError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QXmlStreamWriter) WriteNamespace2(namespaceUri string, prefix string) {
	var namespaceUri_p uintptr
	if len(namespaceUri) > 0 {
		namespaceUri_p = uintptr(unsafe.Pointer(unsafe.StringData(namespaceUri)))
	}
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeNamespace2, this.h, namespaceUri_p, uintptr(len(namespaceUri)), prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(namespaceUri)
	runtime.KeepAlive(prefix)
}

func (this *QXmlStreamWriter) WriteProcessingInstruction2(target string, data string) {
	var target_p uintptr
	if len(target) > 0 {
		target_p = uintptr(unsafe.Pointer(unsafe.StringData(target)))
	}
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(unsafe.StringData(data)))
	}
	purego.SyscallN(addr_QXmlStreamWriter_writeProcessingInstruction2, this.h, target_p, uintptr(len(target)), data_p, uintptr(len(data)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	runtime.KeepAlive(data)
}

func (this *QXmlStreamWriter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QXmlStreamWriter_delete(this.h)
	}
	this.UnsafeRelease()
}
