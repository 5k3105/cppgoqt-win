#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QImageReader_new();
void* QImageReader_new2(void* device);
void* QImageReader_new3(const char* fileName, long long fileName_len);
void* QImageReader_new4(void* device, void* format);
void* QImageReader_new5(const char* fileName, long long fileName_len, void* format);
char* QImageReader_tr(const char* sourceText, long long* __out_len);
void QImageReader_setFormat(void* self, void* format);
void* QImageReader_format(void* self);
void QImageReader_setAutoDetectImageFormat(void* self, int enabled);
int QImageReader_autoDetectImageFormat(void* self);
void QImageReader_setDecideFormatFromContent(void* self, int ignored);
int QImageReader_decideFormatFromContent(void* self);
void QImageReader_setDevice(void* self, void* device);
void* QImageReader_device(void* self);
void QImageReader_setFileName(void* self, const char* fileName, long long fileName_len);
char* QImageReader_fileName(void* self, long long* __out_len);
void* QImageReader_size(void* self);
int QImageReader_imageFormat(void* self);
char** QImageReader_textKeys(void* self, long long* __out_len);
char* QImageReader_text(void* self, const char* key, long long key_len, long long* __out_len);
void QImageReader_setClipRect(void* self, void* rect);
void* QImageReader_clipRect(void* self);
void QImageReader_setScaledSize(void* self, void* size);
void* QImageReader_scaledSize(void* self);
void QImageReader_setQuality(void* self, int quality);
int QImageReader_quality(void* self);
void QImageReader_setScaledClipRect(void* self, void* rect);
void* QImageReader_scaledClipRect(void* self);
void QImageReader_setBackgroundColor(void* self, void* color);
void* QImageReader_backgroundColor(void* self);
int QImageReader_supportsAnimation(void* self);
int QImageReader_transformation(void* self);
void QImageReader_setAutoTransform(void* self, int enabled);
int QImageReader_autoTransform(void* self);
void* QImageReader_subType(void* self);
void** QImageReader_supportedSubTypes(void* self, long long* __out_len);
int QImageReader_canRead(void* self);
void* QImageReader_read(void* self);
int QImageReader_read2(void* self, void* image);
int QImageReader_jumpToNextImage(void* self);
int QImageReader_jumpToImage(void* self, int imageNumber);
int QImageReader_loopCount(void* self);
int QImageReader_imageCount(void* self);
int QImageReader_nextImageDelay(void* self);
int QImageReader_currentImageNumber(void* self);
void* QImageReader_currentImageRect(void* self);
int QImageReader_error(void* self);
char* QImageReader_errorString(void* self, long long* __out_len);
int QImageReader_supportsOption(void* self, int option);
void* QImageReader_imageFormat2(const char* fileName, long long fileName_len);
void* QImageReader_imageFormat3(void* device);
void** QImageReader_supportedImageFormats(long long* __out_len);
void** QImageReader_supportedMimeTypes(long long* __out_len);
void** QImageReader_imageFormatsForMimeType(void* mimeType, long long* __out_len);
int QImageReader_allocationLimit();
void QImageReader_setAllocationLimit(int mbLimit);
char* QImageReader_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QImageReader_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void QImageReader_delete(void* self);

#ifdef __cplusplus
}
#endif
