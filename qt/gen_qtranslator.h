#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTranslator_base(void* self);
void QTranslator__set_override(void* self, int idx);
void* QTranslator_fromQObject(void* o);
void* QTranslator_new(long long __goid);
void* QTranslator_new2(long long __goid, void* parent);
char* QTranslator_tr(const char* s, long long* __out_len);
char* QTranslator_translate(void* self, const char* context, const char* sourceText, const char* disambiguation, int n, long long* __out_len);
int QTranslator_isEmpty(void* self);
char* QTranslator_language(void* self, long long* __out_len);
char* QTranslator_filePath(void* self, long long* __out_len);
int QTranslator_load(void* self, const char* filename, long long filename_len);
int QTranslator_load2(void* self, void* locale, const char* filename, long long filename_len);
int QTranslator_load3(void* self, const unsigned char* data, int len_);
void* QTranslator_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QTranslator_connect2(void* sender, void* signal, void* receiver, void* method);
void* QTranslator_sender(void* self);
int QTranslator_senderSignalIndex(void* self);
int QTranslator_receivers(void* self, const char* signal);
int QTranslator_isSignalConnected(void* self, void* signal);
char* QTranslator_tr2(const char* s, const char* c, long long* __out_len);
char* QTranslator_tr3(const char* s, const char* c, int n, long long* __out_len);
int QTranslator_load4(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len);
int QTranslator_load5(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len, const char* search_delimiters, long long search_delimiters_len);
int QTranslator_load6(void* self, const char* filename, long long filename_len, const char* directory, long long directory_len, const char* search_delimiters, long long search_delimiters_len, const char* suffix, long long suffix_len);
int QTranslator_load7(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len);
int QTranslator_load8(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len, const char* directory, long long directory_len);
int QTranslator_load9(void* self, void* locale, const char* filename, long long filename_len, const char* prefix, long long prefix_len, const char* directory, long long directory_len, const char* suffix, long long suffix_len);
int QTranslator_load10(void* self, const unsigned char* data, int len_, const char* directory, long long directory_len);
void* QTranslator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QTranslator_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
int QTranslator_isEmpty_default(void* self);
void* QTranslator_metaObject_default(void* self);
int QTranslator_event_default(void* self, void* event);
int QTranslator_eventFilter_default(void* self, void* watched, void* event);
void QTranslator_timerEvent_default(void* self, void* event);
void QTranslator_childEvent_default(void* self, void* event);
void QTranslator_customEvent_default(void* self, void* event);
void QTranslator_connectNotify_default(void* self, void* signal);
void QTranslator_disconnectNotify_default(void* self, void* signal);
void QTranslator_connect_destroyed(void* self, long long id);
void QTranslator_connect_destroyed2(void* self, long long id);
void QTranslator_connect_objectNameChanged(void* self, long long id);
void QTranslator__disconnect(long long id);
void* QTranslator__Go_new(long long id, void* parent);
void QTranslator_delete(void* self);

#ifdef __cplusplus
}
#endif
