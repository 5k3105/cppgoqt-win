#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPlainTextDocumentLayout_base(void* self);
void QPlainTextDocumentLayout_chain(void* self, void** out);
void* QPlainTextDocumentLayout_fromQObject(void* o);
void* QPlainTextDocumentLayout_new(void* document);
char* QPlainTextDocumentLayout_tr(const char* s, long long* __out_len);
void QPlainTextDocumentLayout_ensureBlockLayout(void* self, void* block);
void QPlainTextDocumentLayout_setCursorWidth(void* self, int width);
int QPlainTextDocumentLayout_cursorWidth(void* self);
void QPlainTextDocumentLayout_requestUpdate(void* self);
void* QPlainTextDocumentLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QPlainTextDocumentLayout_connect2(void* sender, void* signal, void* receiver, void* method);
char* QPlainTextDocumentLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QPlainTextDocumentLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QPlainTextDocumentLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QPlainTextDocumentLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QPlainTextDocumentLayout_connect_update(void* self, long long id);
void QPlainTextDocumentLayout_connect_update2(void* self, long long id);
void QPlainTextDocumentLayout_connect_updateBlock(void* self, long long id);
void QPlainTextDocumentLayout_connect_documentSizeChanged(void* self, long long id);
void QPlainTextDocumentLayout_connect_pageCountChanged(void* self, long long id);
void QPlainTextDocumentLayout_connect_destroyed(void* self, long long id);
void QPlainTextDocumentLayout_connect_destroyed2(void* self, long long id);
void QPlainTextDocumentLayout_connect_objectNameChanged(void* self, long long id);
void QPlainTextDocumentLayout__disconnect(long long id);
void QPlainTextDocumentLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
