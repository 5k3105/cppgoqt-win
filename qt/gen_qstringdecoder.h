#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStringDecoder_base(void* self);
void QStringDecoder_chain(void* self, void** out);
void* QStringDecoder_new(int encoding);
void* QStringDecoder_new2();
void* QStringDecoder_new3(const char* name, long long name_len);
void* QStringDecoder_new4(int encoding, int flags);
void* QStringDecoder_new5(const char* name, long long name_len, int f);
long long QStringDecoder_requiredSpace(void* self, long long inputLength);
void* QStringDecoder_appendToBuffer(void* self, void* out, void* ba);
int QStringDecoder_encodingForName(const char* name, long long name_len, int* __ok);
const char* QStringDecoder_nameForEncoding(int e);
int QStringDecoder_encodingForData(void* data, int* __ok);
int QStringDecoder_encodingForHtml(void* data, int* __ok);
char** QStringDecoder_availableCodecs(long long* __out_len);
int QStringDecoder_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok);
void QStringDecoder_delete(void* self);

#ifdef __cplusplus
}
#endif
