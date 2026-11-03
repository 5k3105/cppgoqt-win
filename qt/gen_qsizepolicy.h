#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSizePolicy_new();
void* QSizePolicy_new2(int horizontal, int vertical);
void* QSizePolicy_new3(void* arg1);
void* QSizePolicy_new4(int horizontal, int vertical, int type_);
int QSizePolicy_horizontalPolicy(void* self);
int QSizePolicy_verticalPolicy(void* self);
int QSizePolicy_controlType(void* self);
void QSizePolicy_setHorizontalPolicy(void* self, int d);
void QSizePolicy_setVerticalPolicy(void* self, int d);
void QSizePolicy_setControlType(void* self, int type_);
int QSizePolicy_expandingDirections(void* self);
void QSizePolicy_setHeightForWidth(void* self, int b);
int QSizePolicy_hasHeightForWidth(void* self);
void QSizePolicy_setWidthForHeight(void* self, int b);
int QSizePolicy_hasWidthForHeight(void* self);
int QSizePolicy_OperatorEqual(void* self, void* s);
int QSizePolicy_OperatorNotEqual(void* self, void* s);
void* QSizePolicy_ToQVariant(void* self);
int QSizePolicy_horizontalStretch(void* self);
int QSizePolicy_verticalStretch(void* self);
void QSizePolicy_setHorizontalStretch(void* self, int stretchFactor);
void QSizePolicy_setVerticalStretch(void* self, int stretchFactor);
int QSizePolicy_retainSizeWhenHidden(void* self);
void QSizePolicy_setRetainSizeWhenHidden(void* self, int retainSize);
void QSizePolicy_transpose(void* self);
void* QSizePolicy_transposed(void* self);
void* QSizePolicy_OperatorAssign(void* self, void* arg1);
void* QSizePolicy_toQVariant(void* self);
void QSizePolicy_delete(void* self);

#ifdef __cplusplus
}
#endif
