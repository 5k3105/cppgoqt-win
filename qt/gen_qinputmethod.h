#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethod_base(void* self);
void* QInputMethod_fromQObject(void* o);
char* QInputMethod_tr(const char* s, long long* __out_len);
void* QInputMethod_inputItemTransform(void* self);
void QInputMethod_setInputItemTransform(void* self, void* transform);
void* QInputMethod_inputItemRectangle(void* self);
void QInputMethod_setInputItemRectangle(void* self, void* rect);
void* QInputMethod_cursorRectangle(void* self);
void* QInputMethod_anchorRectangle(void* self);
void* QInputMethod_keyboardRectangle(void* self);
void* QInputMethod_inputItemClipRectangle(void* self);
int QInputMethod_isVisible(void* self);
void QInputMethod_setVisible(void* self, int visible);
int QInputMethod_isAnimating(void* self);
void* QInputMethod_locale(void* self);
int QInputMethod_inputDirection(void* self);
void* QInputMethod_queryFocusObject(int query, void* argument);
void QInputMethod_show(void* self);
void QInputMethod_hide(void* self);
void QInputMethod_update(void* self, int queries);
void QInputMethod_reset(void* self);
void QInputMethod_commit(void* self);
void QInputMethod_invokeAction(void* self, int a, int cursorPosition);
void QInputMethod_cursorRectangleChanged(void* self);
void QInputMethod_anchorRectangleChanged(void* self);
void QInputMethod_keyboardRectangleChanged(void* self);
void QInputMethod_inputItemClipRectangleChanged(void* self);
void QInputMethod_visibleChanged(void* self);
void QInputMethod_animatingChanged(void* self);
void QInputMethod_localeChanged(void* self);
void QInputMethod_inputDirectionChanged(void* self, int newDirection);
void* QInputMethod_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QInputMethod_connect2(void* sender, void* signal, void* receiver, void* method);
char* QInputMethod_tr2(const char* s, const char* c, long long* __out_len);
char* QInputMethod_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QInputMethod_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QInputMethod_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QInputMethod_connect_cursorRectangleChanged(void* self, long long id);
void QInputMethod_connect_anchorRectangleChanged(void* self, long long id);
void QInputMethod_connect_keyboardRectangleChanged(void* self, long long id);
void QInputMethod_connect_inputItemClipRectangleChanged(void* self, long long id);
void QInputMethod_connect_visibleChanged(void* self, long long id);
void QInputMethod_connect_animatingChanged(void* self, long long id);
void QInputMethod_connect_localeChanged(void* self, long long id);
void QInputMethod_connect_inputDirectionChanged(void* self, long long id);
void QInputMethod_connect_destroyed(void* self, long long id);
void QInputMethod_connect_destroyed2(void* self, long long id);
void QInputMethod_connect_objectNameChanged(void* self, long long id);
void QInputMethod__disconnect(long long id);
void QInputMethod_delete(void* self);

#ifdef __cplusplus
}
#endif
