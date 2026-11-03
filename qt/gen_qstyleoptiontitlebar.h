#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionTitleBar_base(void* self);
void QStyleOptionTitleBar_chain(void* self, void** out);
void* QStyleOptionTitleBar_new();
void* QStyleOptionTitleBar_new2(void* other);
void* QStyleOptionTitleBar_OperatorAssign(void* self, void* arg1);
char* QStyleOptionTitleBar_text_get(void* self, long long* __out_len);
void QStyleOptionTitleBar_text_set(void* self, const char* text, long long text_len);
void* QStyleOptionTitleBar_icon_get(void* self);
void QStyleOptionTitleBar_icon_set(void* self, void* icon);
int QStyleOptionTitleBar_titleBarState_get(void* self);
void QStyleOptionTitleBar_titleBarState_set(void* self, int titleBarState);
int QStyleOptionTitleBar_titleBarFlags_get(void* self);
void QStyleOptionTitleBar_titleBarFlags_set(void* self, int titleBarFlags);
void QStyleOptionTitleBar_delete(void* self);

#ifdef __cplusplus
}
#endif
