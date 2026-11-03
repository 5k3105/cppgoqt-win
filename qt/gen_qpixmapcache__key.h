#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPixmapCache__Key_new();
void* QPixmapCache__Key_new2(void* other);
int QPixmapCache__Key_OperatorEqual(void* self, void* key);
int QPixmapCache__Key_OperatorNotEqual(void* self, void* key);
void* QPixmapCache__Key_OperatorAssign(void* self, void* other);
void QPixmapCache__Key_swap(void* self, void* other);
int QPixmapCache__Key_isValid(void* self);
void QPixmapCache__Key_delete(void* self);

#ifdef __cplusplus
}
#endif
