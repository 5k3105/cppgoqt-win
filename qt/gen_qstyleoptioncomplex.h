#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionComplex_base(void* self);
void* QStyleOptionComplex_new();
void* QStyleOptionComplex_new2(void* other);
void* QStyleOptionComplex_new3(int version);
void* QStyleOptionComplex_new4(int version, int type_);
void* QStyleOptionComplex_OperatorAssign(void* self, void* arg1);
int QStyleOptionComplex_subControls_get(void* self);
void QStyleOptionComplex_subControls_set(void* self, int subControls);
int QStyleOptionComplex_activeSubControls_get(void* self);
void QStyleOptionComplex_activeSubControls_set(void* self, int activeSubControls);
void QStyleOptionComplex_delete(void* self);

#ifdef __cplusplus
}
#endif
