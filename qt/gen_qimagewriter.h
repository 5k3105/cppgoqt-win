#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QImageWriter_new();
void* QImageWriter_new2(void* device, void* format);
void* QImageWriter_new3(const char* fileName, long long fileName_len);
void* QImageWriter_new4(const char* fileName, long long fileName_len, void* format);
char* QImageWriter_tr(const char* sourceText, long long* __out_len);
void QImageWriter_setFormat(void* self, void* format);
void* QImageWriter_format(void* self);
void QImageWriter_setDevice(void* self, void* device);
void* QImageWriter_device(void* self);
void QImageWriter_setFileName(void* self, const char* fileName, long long fileName_len);
char* QImageWriter_fileName(void* self, long long* __out_len);
void QImageWriter_setQuality(void* self, int quality);
int QImageWriter_quality(void* self);
void QImageWriter_setCompression(void* self, int compression);
int QImageWriter_compression(void* self);
void QImageWriter_setSubType(void* self, void* type_);
void* QImageWriter_subType(void* self);
void** QImageWriter_supportedSubTypes(void* self, long long* __out_len);
void QImageWriter_setOptimizedWrite(void* self, int optimize);
int QImageWriter_optimizedWrite(void* self);
void QImageWriter_setProgressiveScanWrite(void* self, int progressive);
int QImageWriter_progressiveScanWrite(void* self);
int QImageWriter_transformation(void* self);
void QImageWriter_setTransformation(void* self, int orientation);
void QImageWriter_setText(void* self, const char* key, long long key_len, const char* text, long long text_len);
int QImageWriter_canWrite(void* self);
int QImageWriter_write(void* self, void* image);
int QImageWriter_error(void* self);
char* QImageWriter_errorString(void* self, long long* __out_len);
int QImageWriter_supportsOption(void* self, int option);
void** QImageWriter_supportedImageFormats(long long* __out_len);
void** QImageWriter_supportedMimeTypes(long long* __out_len);
void** QImageWriter_imageFormatsForMimeType(void* mimeType, long long* __out_len);
char* QImageWriter_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QImageWriter_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void QImageWriter_delete(void* self);

#ifdef __cplusplus
}
#endif
