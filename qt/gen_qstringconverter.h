#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStringConverter_base(void* self);
int QStringConverter_isValid(void* self);
void QStringConverter_resetState(void* self);
int QStringConverter_hasError(void* self);
const char* QStringConverter_name(void* self);
int QStringConverter_encodingForName(const char* name, long long name_len, int* __ok);
const char* QStringConverter_nameForEncoding(int e);
int QStringConverter_encodingForData(void* data, int* __ok);
int QStringConverter_encodingForHtml(void* data, int* __ok);
char** QStringConverter_availableCodecs(long long* __out_len);
int QStringConverter_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok);
void QStringConverter_delete(void* self);

#ifdef __cplusplus
}
#endif
