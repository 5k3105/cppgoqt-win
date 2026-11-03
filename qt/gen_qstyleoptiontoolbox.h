#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionToolBox_base(void* self);
void* QStyleOptionToolBox_new();
void* QStyleOptionToolBox_new2(void* other);
void* QStyleOptionToolBox_OperatorAssign(void* self, void* arg1);
char* QStyleOptionToolBox_text_get(void* self, long long* __out_len);
void QStyleOptionToolBox_text_set(void* self, const char* text, long long text_len);
void* QStyleOptionToolBox_icon_get(void* self);
void QStyleOptionToolBox_icon_set(void* self, void* icon);
int QStyleOptionToolBox_position_get(void* self);
void QStyleOptionToolBox_position_set(void* self, int position);
int QStyleOptionToolBox_selectedPosition_get(void* self);
void QStyleOptionToolBox_selectedPosition_set(void* self, int selectedPosition);
void QStyleOptionToolBox_delete(void* self);

#ifdef __cplusplus
}
#endif
