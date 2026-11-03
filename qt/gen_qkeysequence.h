#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QKeySequence_new();
void* QKeySequence_new2(const char* key, long long key_len);
void* QKeySequence_new3(int k1);
void* QKeySequence_new4(void* k1);
void* QKeySequence_new5(void* ks);
void* QKeySequence_new6(int key);
void* QKeySequence_new7(const char* key, long long key_len, int format);
void* QKeySequence_new8(int k1, int k2);
void* QKeySequence_new9(int k1, int k2, int k3);
void* QKeySequence_new10(int k1, int k2, int k3, int k4);
void* QKeySequence_new11(void* k1, void* k2);
void* QKeySequence_new12(void* k1, void* k2, void* k3);
void* QKeySequence_new13(void* k1, void* k2, void* k3, void* k4);
int QKeySequence_count(void* self);
int QKeySequence_isEmpty(void* self);
char* QKeySequence_toString(void* self, long long* __out_len);
void* QKeySequence_fromString(const char* str, long long str_len);
void** QKeySequence_listFromString(const char* str, long long str_len, long long* __out_len);
char* QKeySequence_listToString(void** list, long long list_len, long long* __out_len);
int QKeySequence_matches(void* self, void* seq);
void* QKeySequence_mnemonic(const char* text, long long text_len);
void** QKeySequence_keyBindings(int key, long long* __out_len);
void* QKeySequence_ToQVariant(void* self);
void* QKeySequence_OperatorSubscript(void* self, unsigned int i);
void* QKeySequence_OperatorAssign(void* self, void* other);
void QKeySequence_swap(void* self, void* other);
int QKeySequence_OperatorEqual(void* self, void* other);
int QKeySequence_OperatorNotEqual(void* self, void* other);
int QKeySequence_OperatorLesser(void* self, void* ks);
int QKeySequence_OperatorGreater(void* self, void* other);
int QKeySequence_OperatorLesserOrEqual(void* self, void* other);
int QKeySequence_OperatorGreaterOrEqual(void* self, void* other);
int QKeySequence_isDetached(void* self);
char* QKeySequence_toString2(void* self, int format, long long* __out_len);
void* QKeySequence_fromString2(const char* str, long long str_len, int format);
void** QKeySequence_listFromString2(const char* str, long long str_len, int format, long long* __out_len);
char* QKeySequence_listToString2(void** list, long long list_len, int format, long long* __out_len);
void* QKeySequence_toQVariant(void* self);
void QKeySequence_delete(void* self);

#ifdef __cplusplus
}
#endif
