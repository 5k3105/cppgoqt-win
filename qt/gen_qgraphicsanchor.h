#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsAnchor_base(void* self);
void* QGraphicsAnchor_fromQObject(void* o);
char* QGraphicsAnchor_tr(const char* s, long long* __out_len);
void QGraphicsAnchor_setSpacing(void* self, double spacing);
void QGraphicsAnchor_unsetSpacing(void* self);
double QGraphicsAnchor_spacing(void* self);
void QGraphicsAnchor_setSizePolicy(void* self, int policy);
int QGraphicsAnchor_sizePolicy(void* self);
void* QGraphicsAnchor_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGraphicsAnchor_connect2(void* sender, void* signal, void* receiver, void* method);
char* QGraphicsAnchor_tr2(const char* s, const char* c, long long* __out_len);
char* QGraphicsAnchor_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QGraphicsAnchor_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGraphicsAnchor_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QGraphicsAnchor_connect_destroyed(void* self, long long id);
void QGraphicsAnchor_connect_destroyed2(void* self, long long id);
void QGraphicsAnchor_connect_objectNameChanged(void* self, long long id);
void QGraphicsAnchor__disconnect(long long id);
void QGraphicsAnchor_delete(void* self);

#ifdef __cplusplus
}
#endif
