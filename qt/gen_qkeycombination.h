#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QKeyCombination_new();
void* QKeyCombination_new2(int modifiers);
void* QKeyCombination_new3(int modifiers);
void* QKeyCombination_new4(void* arg1);
void* QKeyCombination_new5(int key);
void* QKeyCombination_new6(int modifiers, int key);
void* QKeyCombination_new7(int modifiers, int key);
int QKeyCombination_keyboardModifiers(void* self);
int QKeyCombination_key(void* self);
void* QKeyCombination_fromCombined(int combined);
int QKeyCombination_toCombined(void* self);
int QKeyCombination_ToInt(void* self);
int QKeyCombination_OperatorEqual(void* self, void* rhs);
int QKeyCombination_OperatorNotEqual(void* self, void* rhs);
void QKeyCombination_delete(void* self);

#ifdef __cplusplus
}
#endif
