#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextObject_base(void* self);
void* QTextObject_fromQObject(void* o);
char* QTextObject_tr(const char* s, long long* __out_len);
void* QTextObject_format(void* self);
int QTextObject_formatIndex(void* self);
void* QTextObject_document(void* self);
int QTextObject_objectIndex(void* self);
void* QTextObject_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTextObject_connect2(void* sender, void* signal, void* receiver, void* method);
char* QTextObject_tr2(const char* s, const char* c, long long* __out_len);
char* QTextObject_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTextObject_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTextObject_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QTextObject_connect_destroyed(void* self, long long id);
void QTextObject_connect_destroyed2(void* self, long long id);
void QTextObject_connect_objectNameChanged(void* self, long long id);
void QTextObject__disconnect(long long id);
void QTextObject_delete(void* self);

#ifdef __cplusplus
}
#endif
