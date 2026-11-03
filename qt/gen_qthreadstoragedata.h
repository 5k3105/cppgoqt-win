#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QThreadStorageData_id_get(void* self);
void QThreadStorageData_id_set(void* self, int id);
void QThreadStorageData_delete(void* self);

#ifdef __cplusplus
}
#endif
