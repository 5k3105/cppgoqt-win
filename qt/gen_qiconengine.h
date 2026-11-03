#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QIconEngine__set_override(void* self, int idx);
void QIconEngine_paint(void* self, void* painter, void* rect, int mode, int state);
void* QIconEngine_actualSize(void* self, void* size, int mode, int state);
void* QIconEngine_pixmap(void* self, void* size, int mode, int state);
void QIconEngine_addPixmap(void* self, void* pixmap, int mode, int state);
void QIconEngine_addFile(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state);
char* QIconEngine_key(void* self, long long* __out_len);
void* QIconEngine_clone(void* self);
int QIconEngine_read(void* self, void* in);
int QIconEngine_write(void* self, void* out);
void** QIconEngine_availableSizes(void* self, int mode, int state, long long* __out_len);
char* QIconEngine_iconName(void* self, long long* __out_len);
int QIconEngine_isNull(void* self);
void* QIconEngine_scaledPixmap(void* self, void* size, int mode, int state, double scale);
void* QIconEngine_actualSize_default(void* self, void* size, int mode, int state);
void* QIconEngine_pixmap_default(void* self, void* size, int mode, int state);
void QIconEngine_addPixmap_default(void* self, void* pixmap, int mode, int state);
void QIconEngine_addFile_default(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state);
char* QIconEngine_key_default(void* self, long long* __out_len);
int QIconEngine_read_default(void* self, void* in);
int QIconEngine_write_default(void* self, void* out);
void** QIconEngine_availableSizes_default(void* self, int mode, int state, long long* __out_len);
char* QIconEngine_iconName_default(void* self, long long* __out_len);
int QIconEngine_isNull_default(void* self);
void* QIconEngine_scaledPixmap_default(void* self, void* size, int mode, int state, double scale);
void* QIconEngine__Go_new(long long id);
void QIconEngine_delete(void* self);

#ifdef __cplusplus
}
#endif
