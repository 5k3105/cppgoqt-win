#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamWriter_new();
void* QXmlStreamWriter_new2(void* device);
void* QXmlStreamWriter_new3(void* array);
void* QXmlStreamWriter_new4(void* string_);
void QXmlStreamWriter_setDevice(void* self, void* device);
void* QXmlStreamWriter_device(void* self);
void QXmlStreamWriter_setAutoFormatting(void* self, int arg1);
int QXmlStreamWriter_autoFormatting(void* self);
void QXmlStreamWriter_setAutoFormattingIndent(void* self, int spacesOrTabs);
int QXmlStreamWriter_autoFormattingIndent(void* self);
void QXmlStreamWriter_setStopWritingOnError(void* self, int stop);
int QXmlStreamWriter_stopWritingOnError(void* self);
void QXmlStreamWriter_writeAttribute(void* self, const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len);
void QXmlStreamWriter_writeAttribute2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len);
void QXmlStreamWriter_writeAttribute3(void* self, void* attribute);
void QXmlStreamWriter_writeAttributes(void* self, void* attributes);
void QXmlStreamWriter_writeCDATA(void* self, const char* text, long long text_len);
void QXmlStreamWriter_writeCharacters(void* self, const char* text, long long text_len);
void QXmlStreamWriter_writeComment(void* self, const char* text, long long text_len);
void QXmlStreamWriter_writeDTD(void* self, const char* dtd, long long dtd_len);
void QXmlStreamWriter_writeEmptyElement(void* self, const char* qualifiedName, long long qualifiedName_len);
void QXmlStreamWriter_writeEmptyElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len);
void QXmlStreamWriter_writeTextElement(void* self, const char* qualifiedName, long long qualifiedName_len, const char* text, long long text_len);
void QXmlStreamWriter_writeTextElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* text, long long text_len);
void QXmlStreamWriter_writeEndDocument(void* self);
void QXmlStreamWriter_writeEndElement(void* self);
void QXmlStreamWriter_writeEntityReference(void* self, const char* name, long long name_len);
void QXmlStreamWriter_writeNamespace(void* self, const char* namespaceUri, long long namespaceUri_len);
void QXmlStreamWriter_writeDefaultNamespace(void* self, const char* namespaceUri, long long namespaceUri_len);
void QXmlStreamWriter_writeProcessingInstruction(void* self, const char* target, long long target_len);
void QXmlStreamWriter_writeStartDocument(void* self);
void QXmlStreamWriter_writeStartDocument2(void* self, const char* version, long long version_len);
void QXmlStreamWriter_writeStartDocument3(void* self, const char* version, long long version_len, int standalone);
void QXmlStreamWriter_writeStartElement(void* self, const char* qualifiedName, long long qualifiedName_len);
void QXmlStreamWriter_writeStartElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len);
void QXmlStreamWriter_writeCurrentToken(void* self, void* reader);
void QXmlStreamWriter_raiseError(void* self, const char* message, long long message_len);
char* QXmlStreamWriter_errorString(void* self, long long* __out_len);
int QXmlStreamWriter_error(void* self);
int QXmlStreamWriter_hasError(void* self);
void QXmlStreamWriter_writeNamespace2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* prefix, long long prefix_len);
void QXmlStreamWriter_writeProcessingInstruction2(void* self, const char* target, long long target_len, const char* data, long long data_len);
void QXmlStreamWriter_delete(void* self);

#ifdef __cplusplus
}
#endif
