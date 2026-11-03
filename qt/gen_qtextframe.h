#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextFrame_base(void* self);
void QTextFrame_chain(void* self, void** out);
void* QTextFrame_fromQObject(void* o);
void* QTextFrame_new(void* doc);
char* QTextFrame_tr(const char* s, long long* __out_len);
void QTextFrame_setFrameFormat(void* self, void* format);
void* QTextFrame_frameFormat(void* self);
void* QTextFrame_firstCursorPosition(void* self);
void* QTextFrame_lastCursorPosition(void* self);
int QTextFrame_firstPosition(void* self);
int QTextFrame_lastPosition(void* self);
void* QTextFrame_layoutData(void* self);
void QTextFrame_setLayoutData(void* self, void* data);
void** QTextFrame_childFrames(void* self, long long* __out_len);
void* QTextFrame_parentFrame(void* self);
void* QTextFrame_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTextFrame_connect2(void* sender, void* signal, void* receiver, void* method);
char* QTextFrame_tr2(const char* s, const char* c, long long* __out_len);
char* QTextFrame_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QTextFrame_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTextFrame_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QTextFrame_connect_destroyed(void* self, long long id);
void QTextFrame_connect_destroyed2(void* self, long long id);
void QTextFrame_connect_objectNameChanged(void* self, long long id);
void QTextFrame__disconnect(long long id);
void QTextFrame_delete(void* self);

#ifdef __cplusplus
}
#endif
