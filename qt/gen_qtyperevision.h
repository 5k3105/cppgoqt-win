#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTypeRevision_new();
void* QTypeRevision_new2(void* arg1);
void* QTypeRevision_zero();
int QTypeRevision_hasMajorVersion(void* self);
unsigned char QTypeRevision_majorVersion(void* self);
int QTypeRevision_hasMinorVersion(void* self);
unsigned char QTypeRevision_minorVersion(void* self);
int QTypeRevision_isValid(void* self);
int QTypeRevision_OperatorEqual(void* self, void* rhs);
int QTypeRevision_OperatorNotEqual(void* self, void* rhs);
int QTypeRevision_OperatorLesser(void* self, void* rhs);
int QTypeRevision_OperatorGreater(void* self, void* rhs);
int QTypeRevision_OperatorLesserOrEqual(void* self, void* rhs);
int QTypeRevision_OperatorGreaterOrEqual(void* self, void* rhs);
void QTypeRevision_delete(void* self);

#ifdef __cplusplus
}
#endif
