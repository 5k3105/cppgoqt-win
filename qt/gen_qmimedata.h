#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMimeData_base(void* self);
void QMimeData__set_override(void* self, int idx);
void* QMimeData_fromQObject(void* o);
void* QMimeData_new(long long __goid);
char* QMimeData_tr(const char* s, long long* __out_len);
void** QMimeData_urls(void* self, long long* __out_len);
void QMimeData_setUrls(void* self, void** urls, long long urls_len);
int QMimeData_hasUrls(void* self);
char* QMimeData_text(void* self, long long* __out_len);
void QMimeData_setText(void* self, const char* text, long long text_len);
int QMimeData_hasText(void* self);
char* QMimeData_html(void* self, long long* __out_len);
void QMimeData_setHtml(void* self, const char* html, long long html_len);
int QMimeData_hasHtml(void* self);
void* QMimeData_imageData(void* self);
void QMimeData_setImageData(void* self, void* image);
int QMimeData_hasImage(void* self);
void* QMimeData_colorData(void* self);
void QMimeData_setColorData(void* self, void* color);
int QMimeData_hasColor(void* self);
void* QMimeData_data(void* self, const char* mimetype, long long mimetype_len);
void QMimeData_setData(void* self, const char* mimetype, long long mimetype_len, void* data);
void QMimeData_removeFormat(void* self, const char* mimetype, long long mimetype_len);
int QMimeData_hasFormat(void* self, const char* mimetype, long long mimetype_len);
char** QMimeData_formats(void* self, long long* __out_len);
void QMimeData_clear(void* self);
void* QMimeData_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QMimeData_connect2(void* sender, void* signal, void* receiver, void* method);
void* QMimeData_sender(void* self);
int QMimeData_senderSignalIndex(void* self);
int QMimeData_receivers(void* self, const char* signal);
int QMimeData_isSignalConnected(void* self, void* signal);
char* QMimeData_tr2(const char* s, const char* c, long long* __out_len);
char* QMimeData_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QMimeData_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QMimeData_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
int QMimeData_hasFormat_default(void* self, const char* mimetype, long long mimetype_len);
char** QMimeData_formats_default(void* self, long long* __out_len);
void* QMimeData_metaObject_default(void* self);
void* QMimeData_retrieveData_default(void* self, const char* mimetype, long long mimetype_len, void* preferredType);
int QMimeData_event_default(void* self, void* event);
int QMimeData_eventFilter_default(void* self, void* watched, void* event);
void QMimeData_timerEvent_default(void* self, void* event);
void QMimeData_childEvent_default(void* self, void* event);
void QMimeData_customEvent_default(void* self, void* event);
void QMimeData_connectNotify_default(void* self, void* signal);
void QMimeData_disconnectNotify_default(void* self, void* signal);
void QMimeData_connect_destroyed(void* self, long long id);
void QMimeData_connect_destroyed2(void* self, long long id);
void QMimeData_connect_objectNameChanged(void* self, long long id);
void QMimeData__disconnect(long long id);
void* QMimeData__Go_new(long long id);
void QMimeData_delete(void* self);

#ifdef __cplusplus
}
#endif
