#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionRubberBand_base(void* self);
void* QStyleOptionRubberBand_new();
void* QStyleOptionRubberBand_new2(void* other);
void* QStyleOptionRubberBand_OperatorAssign(void* self, void* arg1);
int QStyleOptionRubberBand_shape_get(void* self);
void QStyleOptionRubberBand_shape_set(void* self, int shape);
int QStyleOptionRubberBand_opaque_get(void* self);
void QStyleOptionRubberBand_opaque_set(void* self, int opaque);
void QStyleOptionRubberBand_delete(void* self);

#ifdef __cplusplus
}
#endif
