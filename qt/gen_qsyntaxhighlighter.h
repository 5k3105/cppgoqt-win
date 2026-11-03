#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSyntaxHighlighter_base(void* self);
void QSyntaxHighlighter__set_override(void* self, int idx);
void* QSyntaxHighlighter_fromQObject(void* o);
char* QSyntaxHighlighter_tr(const char* s, long long* __out_len);
void QSyntaxHighlighter_setDocument(void* self, void* doc);
void* QSyntaxHighlighter_document(void* self);
void QSyntaxHighlighter_rehighlight(void* self);
void QSyntaxHighlighter_rehighlightBlock(void* self, void* block);
void QSyntaxHighlighter_setFormat(void* self, int start, int count, void* format);
void QSyntaxHighlighter_setFormat2(void* self, int start, int count, void* color);
void QSyntaxHighlighter_setFormat3(void* self, int start, int count, void* font);
void* QSyntaxHighlighter_format(void* self, int pos);
int QSyntaxHighlighter_previousBlockState(void* self);
int QSyntaxHighlighter_currentBlockState(void* self);
void QSyntaxHighlighter_setCurrentBlockState(void* self, int newState);
void QSyntaxHighlighter_setCurrentBlockUserData(void* self, void* data);
void* QSyntaxHighlighter_currentBlockUserData(void* self);
void* QSyntaxHighlighter_currentBlock(void* self);
void* QSyntaxHighlighter_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSyntaxHighlighter_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSyntaxHighlighter_sender(void* self);
int QSyntaxHighlighter_senderSignalIndex(void* self);
int QSyntaxHighlighter_receivers(void* self, const char* signal);
int QSyntaxHighlighter_isSignalConnected(void* self, void* signal);
char* QSyntaxHighlighter_tr2(const char* s, const char* c, long long* __out_len);
char* QSyntaxHighlighter_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSyntaxHighlighter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSyntaxHighlighter_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSyntaxHighlighter_metaObject_default(void* self);
int QSyntaxHighlighter_event_default(void* self, void* event);
int QSyntaxHighlighter_eventFilter_default(void* self, void* watched, void* event);
void QSyntaxHighlighter_timerEvent_default(void* self, void* event);
void QSyntaxHighlighter_childEvent_default(void* self, void* event);
void QSyntaxHighlighter_customEvent_default(void* self, void* event);
void QSyntaxHighlighter_connectNotify_default(void* self, void* signal);
void QSyntaxHighlighter_disconnectNotify_default(void* self, void* signal);
void QSyntaxHighlighter_connect_destroyed(void* self, long long id);
void QSyntaxHighlighter_connect_destroyed2(void* self, long long id);
void QSyntaxHighlighter_connect_objectNameChanged(void* self, long long id);
void QSyntaxHighlighter__disconnect(long long id);
void* QSyntaxHighlighter__Go_new(long long id, void* parent);
void QSyntaxHighlighter_delete(void* self);

#ifdef __cplusplus
}
#endif
