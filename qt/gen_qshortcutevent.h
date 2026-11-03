#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QShortcutEvent_base(void* self);
void* QShortcutEvent_new(void* key, int id);
void* QShortcutEvent_new2(void* key);
void* QShortcutEvent_new3(void* key, int id, int ambiguous);
void* QShortcutEvent_new4(void* key, void* shortcut);
void* QShortcutEvent_new5(void* key, void* shortcut, int ambiguous);
void* QShortcutEvent_key(void* self);
int QShortcutEvent_shortcutId(void* self);
int QShortcutEvent_isAmbiguous(void* self);
int QShortcutEvent_registerEventType();
int QShortcutEvent_registerEventType2(int hint);
void QShortcutEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
