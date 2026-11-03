#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionHeader_base(void* self);
void* QStyleOptionHeader_new();
void* QStyleOptionHeader_new2(void* other);
void* QStyleOptionHeader_OperatorAssign(void* self, void* arg1);
int QStyleOptionHeader_section_get(void* self);
void QStyleOptionHeader_section_set(void* self, int section);
char* QStyleOptionHeader_text_get(void* self, long long* __out_len);
void QStyleOptionHeader_text_set(void* self, const char* text, long long text_len);
int QStyleOptionHeader_textAlignment_get(void* self);
void QStyleOptionHeader_textAlignment_set(void* self, int textAlignment);
void* QStyleOptionHeader_icon_get(void* self);
void QStyleOptionHeader_icon_set(void* self, void* icon);
int QStyleOptionHeader_iconAlignment_get(void* self);
void QStyleOptionHeader_iconAlignment_set(void* self, int iconAlignment);
int QStyleOptionHeader_position_get(void* self);
void QStyleOptionHeader_position_set(void* self, int position);
int QStyleOptionHeader_selectedPosition_get(void* self);
void QStyleOptionHeader_selectedPosition_set(void* self, int selectedPosition);
int QStyleOptionHeader_sortIndicator_get(void* self);
void QStyleOptionHeader_sortIndicator_set(void* self, int sortIndicator);
int QStyleOptionHeader_orientation_get(void* self);
void QStyleOptionHeader_orientation_set(void* self, int orientation);
void QStyleOptionHeader_delete(void* self);

#ifdef __cplusplus
}
#endif
