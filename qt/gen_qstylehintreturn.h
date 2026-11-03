#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleHintReturn_new();
void* QStyleHintReturn_new2(void* arg1);
void* QStyleHintReturn_new3(int version);
void* QStyleHintReturn_new4(int version, int type_);
void* QStyleHintReturn_OperatorAssign(void* self, void* arg1);
int QStyleHintReturn_version_get(void* self);
void QStyleHintReturn_version_set(void* self, int version);
int QStyleHintReturn_type_get(void* self);
void QStyleHintReturn_type_set(void* self, int type_);
void QStyleHintReturn_delete(void* self);

#ifdef __cplusplus
}
#endif
