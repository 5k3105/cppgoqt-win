#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QGestureRecognizer__set_override(void* self, int idx);
void* QGestureRecognizer_create(void* self, void* target);
int QGestureRecognizer_recognize(void* self, void* state, void* watched, void* event);
void QGestureRecognizer_reset(void* self, void* state);
int QGestureRecognizer_registerRecognizer(void* recognizer);
void QGestureRecognizer_unregisterRecognizer(int type_);
void* QGestureRecognizer_OperatorAssign(void* self, void* arg1);
void* QGestureRecognizer_create_default(void* self, void* target);
void QGestureRecognizer_reset_default(void* self, void* state);
void* QGestureRecognizer__Go_new(long long id);
void QGestureRecognizer_delete(void* self);

#ifdef __cplusplus
}
#endif
