#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleHyperlinkInterface__set_override(void* self, int idx);
char* QAccessibleHyperlinkInterface_anchor(void* self, long long* __out_len);
char* QAccessibleHyperlinkInterface_anchorTarget(void* self, long long* __out_len);
int QAccessibleHyperlinkInterface_startIndex(void* self);
int QAccessibleHyperlinkInterface_endIndex(void* self);
int QAccessibleHyperlinkInterface_isValid(void* self);
void* QAccessibleHyperlinkInterface_OperatorAssign(void* self, void* arg1);
void* QAccessibleHyperlinkInterface__Go_new(long long id);
void QAccessibleHyperlinkInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
