#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPainterStateGuard_new(void* painter);
void* QPainterStateGuard_new2(void* painter, int state);
void QPainterStateGuard_swap(void* self, void* other);
void QPainterStateGuard_save(void* self);
void QPainterStateGuard_restore(void* self);
void QPainterStateGuard_delete(void* self);

#ifdef __cplusplus
}
#endif
