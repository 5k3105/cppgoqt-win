#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextList_base(void* self);
void QTextList_chain(void* self, void** out);
void* QTextList_fromQObject(void* o);
void* QTextList_new(void* doc);
char* QTextList_tr(const char* s, long long* __out_len);
int QTextList_count(void* self);
void* QTextList_item(void* self, int i);
int QTextList_itemNumber(void* self, void* arg1);
char* QTextList_itemText(void* self, void* arg1, long long* __out_len);
void QTextList_removeItem(void* self, int i);
void QTextList_remove(void* self, void* arg1);
void QTextList_add(void* self, void* block);
void QTextList_setFormat(void* self, void* format);
void* QTextList_format(void* self);
void* QTextList_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTextList_connect2(void* sender, void* signal, void* receiver, void* method);
char* QTextList_tr2(const char* s, const char* c, long long* __out_len);
char* QTextList_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTextList_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTextList_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QTextList_connect_destroyed(void* self, long long id);
void QTextList_connect_destroyed2(void* self, long long id);
void QTextList_connect_objectNameChanged(void* self, long long id);
void QTextList__disconnect(long long id);
void QTextList_delete(void* self);

#ifdef __cplusplus
}
#endif
