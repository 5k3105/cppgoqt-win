#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScrollerProperties_new();
void* QScrollerProperties_new2(void* sp);
void* QScrollerProperties_OperatorAssign(void* self, void* sp);
int QScrollerProperties_OperatorEqual(void* self, void* sp);
int QScrollerProperties_OperatorNotEqual(void* self, void* sp);
void QScrollerProperties_setDefaultScrollerProperties(void* sp);
void QScrollerProperties_unsetDefaultScrollerProperties();
void* QScrollerProperties_scrollMetric(void* self, int metric);
void QScrollerProperties_setScrollMetric(void* self, int metric, void* value);
void QScrollerProperties_delete(void* self);

#ifdef __cplusplus
}
#endif
