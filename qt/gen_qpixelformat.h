#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPixelFormat_new();
void* QPixelFormat_new2(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation);
void* QPixelFormat_new3(void* arg1);
void* QPixelFormat_new4(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation, int byteOrder);
void* QPixelFormat_new5(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation, int byteOrder, unsigned char subEnum);
int QPixelFormat_colorModel(void* self);
unsigned char QPixelFormat_channelCount(void* self);
unsigned char QPixelFormat_redSize(void* self);
unsigned char QPixelFormat_greenSize(void* self);
unsigned char QPixelFormat_blueSize(void* self);
unsigned char QPixelFormat_cyanSize(void* self);
unsigned char QPixelFormat_magentaSize(void* self);
unsigned char QPixelFormat_yellowSize(void* self);
unsigned char QPixelFormat_blackSize(void* self);
unsigned char QPixelFormat_hueSize(void* self);
unsigned char QPixelFormat_saturationSize(void* self);
unsigned char QPixelFormat_lightnessSize(void* self);
unsigned char QPixelFormat_brightnessSize(void* self);
unsigned char QPixelFormat_alphaSize(void* self);
unsigned char QPixelFormat_bitsPerPixel(void* self);
int QPixelFormat_alphaUsage(void* self);
int QPixelFormat_alphaPosition(void* self);
int QPixelFormat_premultiplied(void* self);
int QPixelFormat_typeInterpretation(void* self);
int QPixelFormat_byteOrder(void* self);
int QPixelFormat_yuvLayout(void* self);
unsigned char QPixelFormat_subEnum(void* self);
int QPixelFormat_OperatorEqual(void* self, void* fmt2);
int QPixelFormat_OperatorNotEqual(void* self, void* fmt2);
void QPixelFormat_delete(void* self);

#ifdef __cplusplus
}
#endif
