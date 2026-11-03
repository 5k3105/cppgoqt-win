#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStringEncoder_base(void* self);
void QStringEncoder_chain(void* self, void** out);
void* QStringEncoder_new();
void* QStringEncoder_new2(int encoding);
void* QStringEncoder_new3(const char* name, long long name_len);
void* QStringEncoder_new4(int encoding, int flags);
void* QStringEncoder_new5(const char* name, long long name_len, int flags);
long long QStringEncoder_requiredSpace(void* self, long long inputLength);
int QStringEncoder_encodingForName(const char* name, long long name_len, int* __ok);
const char* QStringEncoder_nameForEncoding(int e);
int QStringEncoder_encodingForData(void* data, int* __ok);
int QStringEncoder_encodingForHtml(void* data, int* __ok);
char** QStringEncoder_availableCodecs(long long* __out_len);
int QStringEncoder_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok);
void QStringEncoder_delete(void* self);

#ifdef __cplusplus
}
#endif
