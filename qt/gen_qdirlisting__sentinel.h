#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QDirListing__sentinel_OperatorEqual(void* self, void* arg1);
int QDirListing__sentinel_OperatorNotEqual(void* self, void* arg1);
void QDirListing__sentinel_delete(void* self);

#ifdef __cplusplus
}
#endif
