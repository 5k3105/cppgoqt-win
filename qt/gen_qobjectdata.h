#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QObjectData_dynamicMetaObject(void* self);
void* QObjectData_q_ptr_get(void* self);
void QObjectData_q_ptr_set(void* self, void* q_ptr);
void* QObjectData_parent_get(void* self);
void QObjectData_parent_set(void* self, void* parent);
void** QObjectData_children_get(void* self, long long* __out_len);
void QObjectData_children_set(void* self, void** children, long long children_len);
unsigned int QObjectData_isWidget_get(void* self);
void QObjectData_isWidget_set(void* self, unsigned int isWidget);
unsigned int QObjectData_blockSig_get(void* self);
void QObjectData_blockSig_set(void* self, unsigned int blockSig);
unsigned int QObjectData_wasDeleted_get(void* self);
void QObjectData_wasDeleted_set(void* self, unsigned int wasDeleted);
unsigned int QObjectData_isDeletingChildren_get(void* self);
void QObjectData_isDeletingChildren_set(void* self, unsigned int isDeletingChildren);
unsigned int QObjectData_sendChildEvents_get(void* self);
void QObjectData_sendChildEvents_set(void* self, unsigned int sendChildEvents);
unsigned int QObjectData_receiveChildEvents_get(void* self);
void QObjectData_receiveChildEvents_set(void* self, unsigned int receiveChildEvents);
unsigned int QObjectData_isWindow_get(void* self);
void QObjectData_isWindow_set(void* self, unsigned int isWindow);
unsigned int QObjectData_deleteLaterCalled_get(void* self);
void QObjectData_deleteLaterCalled_set(void* self, unsigned int deleteLaterCalled);
unsigned int QObjectData_isQuickItem_get(void* self);
void QObjectData_isQuickItem_set(void* self, unsigned int isQuickItem);
unsigned int QObjectData_willBeWidget_get(void* self);
void QObjectData_willBeWidget_set(void* self, unsigned int willBeWidget);
unsigned int QObjectData_wasWidget_get(void* self);
void QObjectData_wasWidget_set(void* self, unsigned int wasWidget);
unsigned int QObjectData_receiveParentEvents_get(void* self);
void QObjectData_receiveParentEvents_set(void* self, unsigned int receiveParentEvents);
unsigned int QObjectData_unused_get(void* self);
void QObjectData_unused_set(void* self, unsigned int unused);
void* QObjectData_postedEvents_get(void* self);
void QObjectData_postedEvents_set(void* self, void* postedEvents);
void* QObjectData_bindingStorage_get(void* self);
void QObjectData_bindingStorage_set(void* self, void* bindingStorage);
void QObjectData_delete(void* self);

#ifdef __cplusplus
}
#endif
