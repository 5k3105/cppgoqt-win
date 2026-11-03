#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QImageIOHandler__set_override(void* self, int idx);
void QImageIOHandler_setDevice(void* self, void* device);
void* QImageIOHandler_device(void* self);
void QImageIOHandler_setFormat(void* self, void* format);
void QImageIOHandler_setFormat2(void* self, void* format);
void* QImageIOHandler_format(void* self);
int QImageIOHandler_canRead(void* self);
int QImageIOHandler_read(void* self, void* image);
int QImageIOHandler_write(void* self, void* image);
void* QImageIOHandler_option(void* self, int option);
void QImageIOHandler_setOption(void* self, int option, void* value);
int QImageIOHandler_supportsOption(void* self, int option);
int QImageIOHandler_jumpToNextImage(void* self);
int QImageIOHandler_jumpToImage(void* self, int imageNumber);
int QImageIOHandler_loopCount(void* self);
int QImageIOHandler_imageCount(void* self);
int QImageIOHandler_nextImageDelay(void* self);
int QImageIOHandler_currentImageNumber(void* self);
void* QImageIOHandler_currentImageRect(void* self);
int QImageIOHandler_allocateImage(void* size, int format, void* image);
int QImageIOHandler_write_default(void* self, void* image);
void* QImageIOHandler_option_default(void* self, int option);
void QImageIOHandler_setOption_default(void* self, int option, void* value);
int QImageIOHandler_supportsOption_default(void* self, int option);
int QImageIOHandler_jumpToNextImage_default(void* self);
int QImageIOHandler_jumpToImage_default(void* self, int imageNumber);
int QImageIOHandler_loopCount_default(void* self);
int QImageIOHandler_imageCount_default(void* self);
int QImageIOHandler_nextImageDelay_default(void* self);
int QImageIOHandler_currentImageNumber_default(void* self);
void* QImageIOHandler_currentImageRect_default(void* self);
void* QImageIOHandler__Go_new(long long id);
void QImageIOHandler_delete(void* self);

#ifdef __cplusplus
}
#endif
