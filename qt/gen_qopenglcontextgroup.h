#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QOpenGLContextGroup_base(void* self);
void* QOpenGLContextGroup_fromQObject(void* o);
char* QOpenGLContextGroup_tr(const char* s, long long* __out_len);
void** QOpenGLContextGroup_shares(void* self, long long* __out_len);
void* QOpenGLContextGroup_currentContextGroup();
void* QOpenGLContextGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QOpenGLContextGroup_connect2(void* sender, void* signal, void* receiver, void* method);
char* QOpenGLContextGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QOpenGLContextGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QOpenGLContextGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QOpenGLContextGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QOpenGLContextGroup_connect_destroyed(void* self, long long id);
void QOpenGLContextGroup_connect_destroyed2(void* self, long long id);
void QOpenGLContextGroup_connect_objectNameChanged(void* self, long long id);
void QOpenGLContextGroup__disconnect(long long id);
void QOpenGLContextGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
