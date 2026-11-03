#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTreeWidgetItemIterator_new(void* it);
void* QTreeWidgetItemIterator_new2(void* widget);
void* QTreeWidgetItemIterator_new3(void* item);
void* QTreeWidgetItemIterator_new4(void* widget, int flags);
void* QTreeWidgetItemIterator_new5(void* item, int flags);
void* QTreeWidgetItemIterator_OperatorAssign(void* self, void* it);
void* QTreeWidgetItemIterator_OperatorPlusPlus(void* self);
void* QTreeWidgetItemIterator_OperatorPlusPlus2(void* self, int arg1);
void* QTreeWidgetItemIterator_OperatorPlusAssign(void* self, int n);
void* QTreeWidgetItemIterator_OperatorMinusMinus(void* self);
void* QTreeWidgetItemIterator_OperatorMinusMinus2(void* self, int arg1);
void* QTreeWidgetItemIterator_OperatorMinusAssign(void* self, int n);
void* QTreeWidgetItemIterator_OperatorMultiply(void* self);
void QTreeWidgetItemIterator_delete(void* self);

#ifdef __cplusplus
}
#endif
