#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAccessibleActionInterface__set_override(void* self, int idx);
char* QAccessibleActionInterface_tr(const char* sourceText, long long* __out_len);
char** QAccessibleActionInterface_actionNames(void* self, long long* __out_len);
char* QAccessibleActionInterface_localizedActionName(void* self, const char* name, long long name_len, long long* __out_len);
char* QAccessibleActionInterface_localizedActionDescription(void* self, const char* name, long long name_len, long long* __out_len);
void QAccessibleActionInterface_doAction(void* self, const char* actionName, long long actionName_len);
char** QAccessibleActionInterface_keyBindingsForAction(void* self, const char* actionName, long long actionName_len, long long* __out_len);
char* QAccessibleActionInterface_pressAction(long long* __out_len);
char* QAccessibleActionInterface_increaseAction(long long* __out_len);
char* QAccessibleActionInterface_decreaseAction(long long* __out_len);
char* QAccessibleActionInterface_showMenuAction(long long* __out_len);
char* QAccessibleActionInterface_setFocusAction(long long* __out_len);
char* QAccessibleActionInterface_toggleAction(long long* __out_len);
char* QAccessibleActionInterface_scrollLeftAction(long long* __out_len);
char* QAccessibleActionInterface_scrollRightAction(long long* __out_len);
char* QAccessibleActionInterface_scrollUpAction(long long* __out_len);
char* QAccessibleActionInterface_scrollDownAction(long long* __out_len);
char* QAccessibleActionInterface_nextPageAction(long long* __out_len);
char* QAccessibleActionInterface_previousPageAction(long long* __out_len);
void* QAccessibleActionInterface_OperatorAssign(void* self, void* arg1);
char* QAccessibleActionInterface_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QAccessibleActionInterface_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
char* QAccessibleActionInterface_localizedActionName_default(void* self, const char* name, long long name_len, long long* __out_len);
char* QAccessibleActionInterface_localizedActionDescription_default(void* self, const char* name, long long name_len, long long* __out_len);
void* QAccessibleActionInterface__Go_new(long long id);
void QAccessibleActionInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
