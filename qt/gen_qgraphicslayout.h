#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsLayout_base(void* self);
void QGraphicsLayout__set_override(void* self, int idx);
void QGraphicsLayout_setContentsMargins(void* self, double left, double top, double right, double bottom);
void QGraphicsLayout_activate(void* self);
int QGraphicsLayout_isActivated(void* self);
void QGraphicsLayout_invalidate(void* self);
void QGraphicsLayout_widgetEvent(void* self, void* e);
int QGraphicsLayout_count(void* self);
void* QGraphicsLayout_itemAt(void* self, int i);
void QGraphicsLayout_removeAt(void* self, int index);
void QGraphicsLayout_setInstantInvalidatePropagation(int enable);
int QGraphicsLayout_instantInvalidatePropagation();
void QGraphicsLayout_addChildLayoutItem(void* self, void* layoutItem);
void QGraphicsLayout_setGraphicsItem(void* self, void* item);
void QGraphicsLayout_setOwnedByLayout(void* self, int ownedByLayout);
void QGraphicsLayout_invalidate_default(void* self);
void QGraphicsLayout_widgetEvent_default(void* self, void* e);
void QGraphicsLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom);
void QGraphicsLayout_updateGeometry_default(void* self);
void QGraphicsLayout_setGeometry_default(void* self, void* rect);
int QGraphicsLayout_isEmpty_default(void* self);
void* QGraphicsLayout__Go_new(long long id);
void QGraphicsLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
