#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QClipboard_base(void* self);
void* QClipboard_fromQObject(void* o);
char* QClipboard_tr(const char* s, long long* __out_len);
void QClipboard_clear(void* self);
int QClipboard_supportsSelection(void* self);
int QClipboard_supportsFindBuffer(void* self);
int QClipboard_ownsSelection(void* self);
int QClipboard_ownsClipboard(void* self);
int QClipboard_ownsFindBuffer(void* self);
char* QClipboard_text(void* self, long long* __out_len);
char* QClipboard_text2(void* self, void* subtype, long long* __out_len);
void QClipboard_setText(void* self, const char* arg1, long long arg1_len);
void* QClipboard_mimeData(void* self);
void QClipboard_setMimeData(void* self, void* data);
void* QClipboard_image(void* self);
void* QClipboard_pixmap(void* self);
void QClipboard_setImage(void* self, void* arg1);
void QClipboard_setPixmap(void* self, void* arg1);
void QClipboard_changed(void* self, int mode);
void QClipboard_selectionChanged(void* self);
void QClipboard_findBufferChanged(void* self);
void QClipboard_dataChanged(void* self);
void* QClipboard_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QClipboard_connect2(void* sender, void* signal, void* receiver, void* method);
char* QClipboard_tr2(const char* s, const char* c, long long* __out_len);
char* QClipboard_tr3(const char* s, const char* c, int n, long long* __out_len);
void QClipboard_clear2(void* self, int mode);
char* QClipboard_text3(void* self, int mode, long long* __out_len);
char* QClipboard_text4(void* self, void* subtype, int mode, long long* __out_len);
void QClipboard_setText2(void* self, const char* arg1, long long arg1_len, int mode);
void* QClipboard_mimeData2(void* self, int mode);
void QClipboard_setMimeData2(void* self, void* data, int mode);
void* QClipboard_image2(void* self, int mode);
void* QClipboard_pixmap2(void* self, int mode);
void QClipboard_setImage2(void* self, void* arg1, int mode);
void QClipboard_setPixmap2(void* self, void* arg1, int mode);
void* QClipboard_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QClipboard_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QClipboard_connect_changed(void* self, long long id);
void QClipboard_connect_selectionChanged(void* self, long long id);
void QClipboard_connect_findBufferChanged(void* self, long long id);
void QClipboard_connect_dataChanged(void* self, long long id);
void QClipboard_connect_destroyed(void* self, long long id);
void QClipboard_connect_destroyed2(void* self, long long id);
void QClipboard_connect_objectNameChanged(void* self, long long id);
void QClipboard__disconnect(long long id);
void QClipboard_delete(void* self);

#ifdef __cplusplus
}
#endif
