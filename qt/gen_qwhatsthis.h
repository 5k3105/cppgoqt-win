#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QWhatsThis_enterWhatsThisMode();
int QWhatsThis_inWhatsThisMode();
void QWhatsThis_leaveWhatsThisMode();
void QWhatsThis_showText(void* pos, const char* text, long long text_len);
void QWhatsThis_hideText();
void* QWhatsThis_createAction();
void QWhatsThis_showText2(void* pos, const char* text, long long text_len, void* w);
void* QWhatsThis_createAction2(void* parent);
void QWhatsThis_delete(void* self);

#ifdef __cplusplus
}
#endif
