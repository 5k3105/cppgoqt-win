#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QKeyEvent_base(void* self);
void QKeyEvent_chain(void* self, void** out);
void* QKeyEvent_new(int type_, int key, int modifiers);
void* QKeyEvent_new2(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers);
void* QKeyEvent_new3(int type_, int key, int modifiers, const char* text, long long text_len);
void* QKeyEvent_new4(int type_, int key, int modifiers, const char* text, long long text_len, int autorep);
void* QKeyEvent_new5(int type_, int key, int modifiers, const char* text, long long text_len, int autorep, unsigned short count);
void* QKeyEvent_new6(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len);
void* QKeyEvent_new7(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep);
void* QKeyEvent_new8(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep, unsigned short count);
void* QKeyEvent_new9(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep, unsigned short count, void* device);
int QKeyEvent_key(void* self);
int QKeyEvent_matches(void* self, int key);
int QKeyEvent_modifiers(void* self);
void* QKeyEvent_keyCombination(void* self);
char* QKeyEvent_text(void* self, long long* __out_len);
int QKeyEvent_isAutoRepeat(void* self);
int QKeyEvent_count(void* self);
unsigned int QKeyEvent_nativeScanCode(void* self);
unsigned int QKeyEvent_nativeVirtualKey(void* self);
unsigned int QKeyEvent_nativeModifiers(void* self);
int QKeyEvent_registerEventType();
int QKeyEvent_registerEventType2(int hint);
void QKeyEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
