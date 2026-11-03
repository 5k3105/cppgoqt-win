#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QLayoutItem__set_override(void* self, int idx);
void* QLayoutItem_sizeHint(void* self);
void* QLayoutItem_minimumSize(void* self);
void* QLayoutItem_maximumSize(void* self);
int QLayoutItem_expandingDirections(void* self);
void QLayoutItem_setGeometry(void* self, void* arg1);
void* QLayoutItem_geometry(void* self);
int QLayoutItem_isEmpty(void* self);
int QLayoutItem_hasHeightForWidth(void* self);
int QLayoutItem_heightForWidth(void* self, int arg1);
int QLayoutItem_minimumHeightForWidth(void* self, int arg1);
void QLayoutItem_invalidate(void* self);
void* QLayoutItem_widget(void* self);
void* QLayoutItem_layout(void* self);
void* QLayoutItem_spacerItem(void* self);
int QLayoutItem_alignment(void* self);
void QLayoutItem_setAlignment(void* self, int a);
int QLayoutItem_controlTypes(void* self);
void* QLayoutItem_OperatorAssign(void* self, void* arg1);
int QLayoutItem_hasHeightForWidth_default(void* self);
int QLayoutItem_heightForWidth_default(void* self, int arg1);
int QLayoutItem_minimumHeightForWidth_default(void* self, int arg1);
void QLayoutItem_invalidate_default(void* self);
void* QLayoutItem_widget_default(void* self);
void* QLayoutItem_layout_default(void* self);
void* QLayoutItem_spacerItem_default(void* self);
int QLayoutItem_controlTypes_default(void* self);
void* QLayoutItem__Go_new(long long id);
void QLayoutItem_delete(void* self);

#ifdef __cplusplus
}
#endif
