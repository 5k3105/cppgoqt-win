#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QXmlStreamReader_new();
void* QXmlStreamReader_new2(void* device);
void* QXmlStreamReader_new3(const char* data, long long data_len);
void QXmlStreamReader_setDevice(void* self, void* device);
void* QXmlStreamReader_device(void* self);
void QXmlStreamReader_addData(void* self, const char* data, long long data_len);
void QXmlStreamReader_clear(void* self);
int QXmlStreamReader_atEnd(void* self);
int QXmlStreamReader_readNext(void* self);
int QXmlStreamReader_readNextStartElement(void* self);
void QXmlStreamReader_skipCurrentElement(void* self);
char* QXmlStreamReader_readRawInnerData(void* self, long long* __out_len);
int QXmlStreamReader_tokenType(void* self);
char* QXmlStreamReader_tokenString(void* self, long long* __out_len);
void QXmlStreamReader_setNamespaceProcessing(void* self, int arg1);
int QXmlStreamReader_namespaceProcessing(void* self);
int QXmlStreamReader_isStartDocument(void* self);
int QXmlStreamReader_isEndDocument(void* self);
int QXmlStreamReader_isStartElement(void* self);
int QXmlStreamReader_isEndElement(void* self);
int QXmlStreamReader_isCharacters(void* self);
int QXmlStreamReader_isWhitespace(void* self);
int QXmlStreamReader_isCDATA(void* self);
int QXmlStreamReader_isComment(void* self);
int QXmlStreamReader_isDTD(void* self);
int QXmlStreamReader_isEntityReference(void* self);
int QXmlStreamReader_isProcessingInstruction(void* self);
int QXmlStreamReader_isStandaloneDocument(void* self);
int QXmlStreamReader_hasStandaloneDeclaration(void* self);
void* QXmlStreamReader_documentVersion(void* self);
void* QXmlStreamReader_documentEncoding(void* self);
long long QXmlStreamReader_lineNumber(void* self);
long long QXmlStreamReader_columnNumber(void* self);
long long QXmlStreamReader_characterOffset(void* self);
void* QXmlStreamReader_attributes(void* self);
char* QXmlStreamReader_readElementText(void* self, long long* __out_len);
void* QXmlStreamReader_name(void* self);
void* QXmlStreamReader_namespaceUri(void* self);
void* QXmlStreamReader_qualifiedName(void* self);
void* QXmlStreamReader_prefix(void* self);
void* QXmlStreamReader_processingInstructionTarget(void* self);
void* QXmlStreamReader_processingInstructionData(void* self);
void* QXmlStreamReader_text(void* self);
void** QXmlStreamReader_namespaceDeclarations(void* self, long long* __out_len);
void QXmlStreamReader_addExtraNamespaceDeclaration(void* self, void* extraNamespaceDeclaraction);
void QXmlStreamReader_addExtraNamespaceDeclarations(void* self, void** extraNamespaceDeclaractions, long long extraNamespaceDeclaractions_len);
void** QXmlStreamReader_notationDeclarations(void* self, long long* __out_len);
void** QXmlStreamReader_entityDeclarations(void* self, long long* __out_len);
void* QXmlStreamReader_dtdName(void* self);
void* QXmlStreamReader_dtdPublicId(void* self);
void* QXmlStreamReader_dtdSystemId(void* self);
int QXmlStreamReader_entityExpansionLimit(void* self);
void QXmlStreamReader_setEntityExpansionLimit(void* self, int limit);
void QXmlStreamReader_raiseError(void* self);
char* QXmlStreamReader_errorString(void* self, long long* __out_len);
int QXmlStreamReader_error(void* self);
int QXmlStreamReader_hasError(void* self);
void QXmlStreamReader_setEntityResolver(void* self, void* resolver);
void* QXmlStreamReader_entityResolver(void* self);
char* QXmlStreamReader_readElementText2(void* self, int behaviour, long long* __out_len);
void QXmlStreamReader_raiseError2(void* self, const char* message, long long message_len);
void QXmlStreamReader_delete(void* self);

#ifdef __cplusplus
}
#endif
