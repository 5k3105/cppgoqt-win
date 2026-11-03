#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleWidgetV2_base(void* self);
void QAccessibleWidgetV2_chain(void* self, void** out);
void* QAccessibleWidgetV2_new(void* object);
void* QAccessibleWidgetV2_new2(void* object, int role, const char* name, long long name_len);
void* QAccessibleWidgetV2_new3(void* object, int role);
int* QAccessibleWidgetV2_attributeKeys(void* self, long long* __out_len);
void* QAccessibleWidgetV2_attributeValue(void* self, int key);
char* QAccessibleWidgetV2_tr(const char* sourceText, long long* __out_len);
char* QAccessibleWidgetV2_pressAction(long long* __out_len);
char* QAccessibleWidgetV2_increaseAction(long long* __out_len);
char* QAccessibleWidgetV2_decreaseAction(long long* __out_len);
char* QAccessibleWidgetV2_showMenuAction(long long* __out_len);
char* QAccessibleWidgetV2_setFocusAction(long long* __out_len);
char* QAccessibleWidgetV2_toggleAction(long long* __out_len);
char* QAccessibleWidgetV2_scrollLeftAction(long long* __out_len);
char* QAccessibleWidgetV2_scrollRightAction(long long* __out_len);
char* QAccessibleWidgetV2_scrollUpAction(long long* __out_len);
char* QAccessibleWidgetV2_scrollDownAction(long long* __out_len);
char* QAccessibleWidgetV2_nextPageAction(long long* __out_len);
char* QAccessibleWidgetV2_previousPageAction(long long* __out_len);
char* QAccessibleWidgetV2_tr2(const char* sourceText, const char* disambiguation, long long* __out_len);
char* QAccessibleWidgetV2_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len);
void* QAccessibleWidgetV2_asQAccessibleAttributesInterface(void* self);
void QAccessibleWidgetV2_delete(void* self);

#ifdef __cplusplus
}
#endif
