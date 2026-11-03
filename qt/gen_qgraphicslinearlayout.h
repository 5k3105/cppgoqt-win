#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsLinearLayout_base(void* self);
void QGraphicsLinearLayout_chain(void* self, void** out);
void QGraphicsLinearLayout__set_override(void* self, int idx);
void* QGraphicsLinearLayout_new(long long __goid);
void* QGraphicsLinearLayout_new2(long long __goid, int orientation);
void* QGraphicsLinearLayout_new3(long long __goid, void* parent);
void* QGraphicsLinearLayout_new4(long long __goid, int orientation, void* parent);
void QGraphicsLinearLayout_setOrientation(void* self, int orientation);
int QGraphicsLinearLayout_orientation(void* self);
void QGraphicsLinearLayout_addItem(void* self, void* item);
void QGraphicsLinearLayout_addStretch(void* self);
void QGraphicsLinearLayout_insertItem(void* self, int index, void* item);
void QGraphicsLinearLayout_insertStretch(void* self, int index);
void QGraphicsLinearLayout_removeItem(void* self, void* item);
void QGraphicsLinearLayout_setSpacing(void* self, double spacing);
double QGraphicsLinearLayout_spacing(void* self);
void QGraphicsLinearLayout_setItemSpacing(void* self, int index, double spacing);
double QGraphicsLinearLayout_itemSpacing(void* self, int index);
void QGraphicsLinearLayout_setStretchFactor(void* self, void* item, int stretch);
int QGraphicsLinearLayout_stretchFactor(void* self, void* item);
void QGraphicsLinearLayout_setAlignment(void* self, void* item, int alignment);
int QGraphicsLinearLayout_alignment(void* self, void* item);
void QGraphicsLinearLayout_dump(void* self);
void QGraphicsLinearLayout_setInstantInvalidatePropagation(int enable);
int QGraphicsLinearLayout_instantInvalidatePropagation();
void QGraphicsLinearLayout_addChildLayoutItem(void* self, void* layoutItem);
void QGraphicsLinearLayout_setGraphicsItem(void* self, void* item);
void QGraphicsLinearLayout_setOwnedByLayout(void* self, int ownedByLayout);
void QGraphicsLinearLayout_addStretch2(void* self, int stretch);
void QGraphicsLinearLayout_insertStretch2(void* self, int index, int stretch);
void QGraphicsLinearLayout_dump2(void* self, int indent);
void QGraphicsLinearLayout_removeAt_default(void* self, int index);
void QGraphicsLinearLayout_setGeometry_default(void* self, void* rect);
int QGraphicsLinearLayout_count_default(void* self);
void* QGraphicsLinearLayout_itemAt_default(void* self, int index);
void QGraphicsLinearLayout_invalidate_default(void* self);
void* QGraphicsLinearLayout_sizeHint_default(void* self, int which, void* constraint);
void QGraphicsLinearLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom);
void QGraphicsLinearLayout_updateGeometry_default(void* self);
void QGraphicsLinearLayout_widgetEvent_default(void* self, void* e);
int QGraphicsLinearLayout_isEmpty_default(void* self);
void* QGraphicsLinearLayout__Go_new(long long id);
void QGraphicsLinearLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
