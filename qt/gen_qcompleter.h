#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCompleter_base(void* self);
void QCompleter__set_override(void* self, int idx);
void* QCompleter_fromQObject(void* o);
void* QCompleter_new(long long __goid);
void* QCompleter_new2(long long __goid, void* model);
void* QCompleter_new3(long long __goid, char** completions, long long completions_len);
void* QCompleter_new4(long long __goid, void* parent);
void* QCompleter_new5(long long __goid, void* model, void* parent);
void* QCompleter_new6(long long __goid, char** completions, long long completions_len, void* parent);
char* QCompleter_tr(const char* s, long long* __out_len);
void QCompleter_setWidget(void* self, void* widget);
void* QCompleter_widget(void* self);
void QCompleter_setModel(void* self, void* c);
void* QCompleter_model(void* self);
void QCompleter_setCompletionMode(void* self, int mode);
int QCompleter_completionMode(void* self);
void QCompleter_setFilterMode(void* self, int filterMode);
int QCompleter_filterMode(void* self);
void* QCompleter_popup(void* self);
void QCompleter_setPopup(void* self, void* popup);
void QCompleter_setCaseSensitivity(void* self, int caseSensitivity);
int QCompleter_caseSensitivity(void* self);
void QCompleter_setModelSorting(void* self, int sorting);
int QCompleter_modelSorting(void* self);
void QCompleter_setCompletionColumn(void* self, int column);
int QCompleter_completionColumn(void* self);
void QCompleter_setCompletionRole(void* self, int role);
int QCompleter_completionRole(void* self);
int QCompleter_wrapAround(void* self);
int QCompleter_maxVisibleItems(void* self);
void QCompleter_setMaxVisibleItems(void* self, int maxItems);
int QCompleter_completionCount(void* self);
int QCompleter_setCurrentRow(void* self, int row);
int QCompleter_currentRow(void* self);
void* QCompleter_currentIndex(void* self);
char* QCompleter_currentCompletion(void* self, long long* __out_len);
void* QCompleter_completionModel(void* self);
char* QCompleter_completionPrefix(void* self, long long* __out_len);
void QCompleter_setCompletionPrefix(void* self, const char* prefix, long long prefix_len);
void QCompleter_complete(void* self);
void QCompleter_setWrapAround(void* self, int wrap);
char* QCompleter_pathFromIndex(void* self, void* index, long long* __out_len);
char** QCompleter_splitPath(void* self, const char* path, long long path_len, long long* __out_len);
void QCompleter_activated(void* self, const char* text, long long text_len);
void QCompleter_activated2(void* self, void* index);
void QCompleter_highlighted(void* self, const char* text, long long text_len);
void QCompleter_highlighted2(void* self, void* index);
void* QCompleter_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QCompleter_connect2(void* sender, void* signal, void* receiver, void* method);
void* QCompleter_sender(void* self);
int QCompleter_senderSignalIndex(void* self);
int QCompleter_receivers(void* self, const char* signal);
int QCompleter_isSignalConnected(void* self, void* signal);
char* QCompleter_tr2(const char* s, const char* c, long long* __out_len);
char* QCompleter_tr3(const char* s, const char* c, int n, long long* __out_len);
void QCompleter_complete2(void* self, void* rect);
void* QCompleter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QCompleter_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
char* QCompleter_pathFromIndex_default(void* self, void* index, long long* __out_len);
char** QCompleter_splitPath_default(void* self, const char* path, long long path_len, long long* __out_len);
void* QCompleter_metaObject_default(void* self);
int QCompleter_eventFilter_default(void* self, void* o, void* e);
int QCompleter_event_default(void* self, void* arg1);
void QCompleter_timerEvent_default(void* self, void* event);
void QCompleter_childEvent_default(void* self, void* event);
void QCompleter_customEvent_default(void* self, void* event);
void QCompleter_connectNotify_default(void* self, void* signal);
void QCompleter_disconnectNotify_default(void* self, void* signal);
void QCompleter_connect_activated(void* self, long long id);
void QCompleter_connect_activated2(void* self, long long id);
void QCompleter_connect_highlighted(void* self, long long id);
void QCompleter_connect_highlighted2(void* self, long long id);
void QCompleter_connect_destroyed(void* self, long long id);
void QCompleter_connect_destroyed2(void* self, long long id);
void QCompleter_connect_objectNameChanged(void* self, long long id);
void QCompleter__disconnect(long long id);
void* QCompleter__Go_new(long long id, void* parent);
void QCompleter_delete(void* self);

#ifdef __cplusplus
}
#endif
