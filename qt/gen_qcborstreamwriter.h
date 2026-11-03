#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCborStreamWriter_new(void* device);
void* QCborStreamWriter_new2(void* data);
void QCborStreamWriter_setDevice(void* self, void* device);
void* QCborStreamWriter_device(void* self);
void QCborStreamWriter_append(void* self, unsigned long long u);
void QCborStreamWriter_append2(void* self, long long i);
void QCborStreamWriter_append3(void* self, int n);
void QCborStreamWriter_append4(void* self, void* ba);
void QCborStreamWriter_append5(void* self, void* str);
void QCborStreamWriter_append6(void* self, void* str);
void QCborStreamWriter_append7(void* self, int tag);
void QCborStreamWriter_append8(void* self, int tag);
void QCborStreamWriter_append9(void* self, int st);
void QCborStreamWriter_append10(void* self, void* f);
void QCborStreamWriter_append11(void* self, float f);
void QCborStreamWriter_append12(void* self, double d);
void QCborStreamWriter_appendByteString(void* self, const char* data, long long len_);
void QCborStreamWriter_appendTextString(void* self, const char* utf8, long long len_);
void QCborStreamWriter_append13(void* self, int b);
void QCborStreamWriter_appendNull(void* self);
void QCborStreamWriter_appendUndefined(void* self);
void QCborStreamWriter_append14(void* self, int i);
void QCborStreamWriter_append15(void* self, unsigned int u);
void QCborStreamWriter_append16(void* self, const char* str);
void QCborStreamWriter_startArray(void* self);
void QCborStreamWriter_startArray2(void* self, unsigned long long count);
int QCborStreamWriter_endArray(void* self);
void QCborStreamWriter_startMap(void* self);
void QCborStreamWriter_startMap2(void* self, unsigned long long count);
int QCborStreamWriter_endMap(void* self);
void QCborStreamWriter_append17(void* self, const char* str, long long size);
void QCborStreamWriter_delete(void* self);

#ifdef __cplusplus
}
#endif
