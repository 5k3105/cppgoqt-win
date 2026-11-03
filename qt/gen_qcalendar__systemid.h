#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCalendar__SystemId_new();
unsigned long long QCalendar__SystemId_index(void* self);
int QCalendar__SystemId_isValid(void* self);
void QCalendar__SystemId_delete(void* self);

#ifdef __cplusplus
}
#endif
