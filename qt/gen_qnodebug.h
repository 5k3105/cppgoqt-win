#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QNoDebug_OperatorShiftLeft(void* self, void* arg1);
void* QNoDebug_space(void* self);
void* QNoDebug_nospace(void* self);
void* QNoDebug_maybeSpace(void* self);
void* QNoDebug_quote(void* self);
void* QNoDebug_noquote(void* self);
void* QNoDebug_maybeQuote(void* self);
void* QNoDebug_verbosity(void* self, int arg1);
void* QNoDebug_maybeQuote2(void* self, char arg1);
void QNoDebug_delete(void* self);

#ifdef __cplusplus
}
#endif
