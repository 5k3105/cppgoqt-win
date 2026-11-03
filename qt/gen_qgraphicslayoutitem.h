#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QGraphicsLayoutItem__set_override(void* self, int idx);
void QGraphicsLayoutItem_setSizePolicy(void* self, void* policy);
void QGraphicsLayoutItem_setSizePolicy2(void* self, int hPolicy, int vPolicy);
void* QGraphicsLayoutItem_sizePolicy(void* self);
void QGraphicsLayoutItem_setMinimumSize(void* self, void* size);
void QGraphicsLayoutItem_setMinimumSize2(void* self, double w, double h);
void* QGraphicsLayoutItem_minimumSize(void* self);
void QGraphicsLayoutItem_setMinimumWidth(void* self, double width);
double QGraphicsLayoutItem_minimumWidth(void* self);
void QGraphicsLayoutItem_setMinimumHeight(void* self, double height);
double QGraphicsLayoutItem_minimumHeight(void* self);
void QGraphicsLayoutItem_setPreferredSize(void* self, void* size);
void QGraphicsLayoutItem_setPreferredSize2(void* self, double w, double h);
void* QGraphicsLayoutItem_preferredSize(void* self);
void QGraphicsLayoutItem_setPreferredWidth(void* self, double width);
double QGraphicsLayoutItem_preferredWidth(void* self);
void QGraphicsLayoutItem_setPreferredHeight(void* self, double height);
double QGraphicsLayoutItem_preferredHeight(void* self);
void QGraphicsLayoutItem_setMaximumSize(void* self, void* size);
void QGraphicsLayoutItem_setMaximumSize2(void* self, double w, double h);
void* QGraphicsLayoutItem_maximumSize(void* self);
void QGraphicsLayoutItem_setMaximumWidth(void* self, double width);
double QGraphicsLayoutItem_maximumWidth(void* self);
void QGraphicsLayoutItem_setMaximumHeight(void* self, double height);
double QGraphicsLayoutItem_maximumHeight(void* self);
void QGraphicsLayoutItem_setGeometry(void* self, void* rect);
void* QGraphicsLayoutItem_geometry(void* self);
void QGraphicsLayoutItem_getContentsMargins(void* self, double* left, double* top, double* right, double* bottom);
void* QGraphicsLayoutItem_contentsRect(void* self);
void* QGraphicsLayoutItem_effectiveSizeHint(void* self, int which);
void QGraphicsLayoutItem_updateGeometry(void* self);
int QGraphicsLayoutItem_isEmpty(void* self);
void* QGraphicsLayoutItem_parentLayoutItem(void* self);
void QGraphicsLayoutItem_setParentLayoutItem(void* self, void* parent);
int QGraphicsLayoutItem_isLayout(void* self);
void* QGraphicsLayoutItem_graphicsItem(void* self);
int QGraphicsLayoutItem_ownedByLayout(void* self);
void QGraphicsLayoutItem_setGraphicsItem(void* self, void* item);
void QGraphicsLayoutItem_setOwnedByLayout(void* self, int ownedByLayout);
void QGraphicsLayoutItem_setSizePolicy3(void* self, int hPolicy, int vPolicy, int controlType);
void* QGraphicsLayoutItem_effectiveSizeHint2(void* self, int which, void* constraint);
void QGraphicsLayoutItem_setGeometry_default(void* self, void* rect);
void QGraphicsLayoutItem_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom);
void QGraphicsLayoutItem_updateGeometry_default(void* self);
int QGraphicsLayoutItem_isEmpty_default(void* self);
void* QGraphicsLayoutItem__Go_new(long long id);
void QGraphicsLayoutItem_delete(void* self);

#ifdef __cplusplus
}
#endif
