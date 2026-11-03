#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsAnchorLayout_base(void* self);
void QGraphicsAnchorLayout_chain(void* self, void** out);
void QGraphicsAnchorLayout__set_override(void* self, int idx);
void* QGraphicsAnchorLayout_new(long long __goid);
void* QGraphicsAnchorLayout_new2(long long __goid, void* parent);
void* QGraphicsAnchorLayout_addAnchor(void* self, void* firstItem, int firstEdge, void* secondItem, int secondEdge);
void* QGraphicsAnchorLayout_anchor(void* self, void* firstItem, int firstEdge, void* secondItem, int secondEdge);
void QGraphicsAnchorLayout_addCornerAnchors(void* self, void* firstItem, int firstCorner, void* secondItem, int secondCorner);
void QGraphicsAnchorLayout_addAnchors(void* self, void* firstItem, void* secondItem);
void QGraphicsAnchorLayout_setHorizontalSpacing(void* self, double spacing);
void QGraphicsAnchorLayout_setVerticalSpacing(void* self, double spacing);
void QGraphicsAnchorLayout_setSpacing(void* self, double spacing);
double QGraphicsAnchorLayout_horizontalSpacing(void* self);
double QGraphicsAnchorLayout_verticalSpacing(void* self);
void QGraphicsAnchorLayout_setInstantInvalidatePropagation(int enable);
int QGraphicsAnchorLayout_instantInvalidatePropagation();
void QGraphicsAnchorLayout_addChildLayoutItem(void* self, void* layoutItem);
void QGraphicsAnchorLayout_setGraphicsItem(void* self, void* item);
void QGraphicsAnchorLayout_setOwnedByLayout(void* self, int ownedByLayout);
void QGraphicsAnchorLayout_addAnchors2(void* self, void* firstItem, void* secondItem, int orientations);
void QGraphicsAnchorLayout_removeAt_default(void* self, int index);
void QGraphicsAnchorLayout_setGeometry_default(void* self, void* rect);
int QGraphicsAnchorLayout_count_default(void* self);
void* QGraphicsAnchorLayout_itemAt_default(void* self, int index);
void QGraphicsAnchorLayout_invalidate_default(void* self);
void* QGraphicsAnchorLayout_sizeHint_default(void* self, int which, void* constraint);
void QGraphicsAnchorLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom);
void QGraphicsAnchorLayout_updateGeometry_default(void* self);
void QGraphicsAnchorLayout_widgetEvent_default(void* self, void* e);
int QGraphicsAnchorLayout_isEmpty_default(void* self);
void* QGraphicsAnchorLayout__Go_new(long long id);
void QGraphicsAnchorLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
