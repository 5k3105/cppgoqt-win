#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTableModelChangeEvent_base(void* self);
void* QAccessibleTableModelChangeEvent_new(void* obj, int changeType);
void* QAccessibleTableModelChangeEvent_new2(void* iface, int changeType);
void QAccessibleTableModelChangeEvent_setModelChangeType(void* self, int changeType);
int QAccessibleTableModelChangeEvent_modelChangeType(void* self);
void QAccessibleTableModelChangeEvent_setFirstRow(void* self, int row);
void QAccessibleTableModelChangeEvent_setFirstColumn(void* self, int col);
void QAccessibleTableModelChangeEvent_setLastRow(void* self, int row);
void QAccessibleTableModelChangeEvent_setLastColumn(void* self, int col);
int QAccessibleTableModelChangeEvent_firstRow(void* self);
int QAccessibleTableModelChangeEvent_firstColumn(void* self);
int QAccessibleTableModelChangeEvent_lastRow(void* self);
int QAccessibleTableModelChangeEvent_lastColumn(void* self);
void QAccessibleTableModelChangeEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
