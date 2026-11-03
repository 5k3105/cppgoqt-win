#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextBlockGroup_base(void* self);
void QTextBlockGroup_chain(void* self, void** out);
void* QTextBlockGroup_fromQObject(void* o);
char* QTextBlockGroup_tr(const char* s, long long* __out_len);
void* QTextBlockGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTextBlockGroup_connect2(void* sender, void* signal, void* receiver, void* method);
char* QTextBlockGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QTextBlockGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTextBlockGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTextBlockGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QTextBlockGroup_connect_destroyed(void* self, long long id);
void QTextBlockGroup_connect_destroyed2(void* self, long long id);
void QTextBlockGroup_connect_objectNameChanged(void* self, long long id);
void QTextBlockGroup__disconnect(long long id);
void QTextBlockGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
