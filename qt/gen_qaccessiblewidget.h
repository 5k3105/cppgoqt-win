#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleWidget_base(void* self);
void QAccessibleWidget_chain(void* self, void** out);
void* QAccessibleWidget_new(void* o);
void* QAccessibleWidget_new2(void* o, int r, const char* name, long long name_len);
void* QAccessibleWidget_new3(void* o, int r);
char** QAccessibleWidget_actionNames(void* self, long long* __out_len);
void QAccessibleWidget_doAction(void* self, const char* actionName, long long actionName_len);
char** QAccessibleWidget_keyBindingsForAction(void* self, const char* actionName, long long actionName_len, long long* __out_len);
char* QAccessibleWidget_tr(const char* sourceText, long long* __out_len);
char* QAccessibleWidget_localizedActionName(void* self, const char* name, long long name_len, long long* __out_len);
char* QAccessibleWidget_localizedActionDescription(void* self, const char* name, long long name_len, long long* __out_len);
char* QAccessibleWidget_pressAction(long long* __out_len);
char* QAccessibleWidget_increaseAction(long long* __out_len);
char* QAccessibleWidget_decreaseAction(long long* __out_len);
char* QAccessibleWidget_showMenuAction(long long* __out_len);
char* QAccessibleWidget_setFocusAction(long long* __out_len);
char* QAccessibleWidget_toggleAction(long long* __out_len);
char* QAccessibleWidget_scrollLeftAction(long long* __out_len);
char* QAccessibleWidget_scrollRightAction(long long* __out_len);
char* QAccessibleWidget_scrollUpAction(long long* __out_len);
char* QAccessibleWidget_scrollDownAction(long long* __out_len);
char* QAccessibleWidget_nextPageAction(long long* __out_len);
char* QAccessibleWidget_previousPageAction(long long* __out_len);
char* QAccessibleWidget_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QAccessibleWidget_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void* QAccessibleWidget_asQAccessibleActionInterface(void* self);
void QAccessibleWidget_delete(void* self);

#ifdef __cplusplus
}
#endif
