#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QScroller_base(void* self);
void* QScroller_fromQObject(void* o);
char* QScroller_tr(const char* s, long long* __out_len);
int QScroller_hasScroller(void* target);
void* QScroller_scroller(void* target);
void* QScroller_scroller2(void* target);
int QScroller_grabGesture(void* target);
int QScroller_grabbedGesture(void* target);
void QScroller_ungrabGesture(void* target);
void** QScroller_activeScrollers(long long* __out_len);
void* QScroller_target(void* self);
int QScroller_state(void* self);
int QScroller_handleInput(void* self, int input, void* position);
void QScroller_stop(void* self);
void* QScroller_velocity(void* self);
void* QScroller_finalPosition(void* self);
void* QScroller_pixelPerMeter(void* self);
void* QScroller_scrollerProperties(void* self);
void QScroller_setSnapPositionsX(void* self, double* positions, long long positions_len);
void QScroller_setSnapPositionsX2(void* self, double first, double interval);
void QScroller_setSnapPositionsY(void* self, double* positions, long long positions_len);
void QScroller_setSnapPositionsY2(void* self, double first, double interval);
void QScroller_setScrollerProperties(void* self, void* prop);
void QScroller_scrollTo(void* self, void* pos);
void QScroller_scrollTo2(void* self, void* pos, int scrollTime);
void QScroller_ensureVisible(void* self, void* rect, double xmargin, double ymargin);
void QScroller_ensureVisible2(void* self, void* rect, double xmargin, double ymargin, int scrollTime);
void QScroller_resendPrepareEvent(void* self);
void QScroller_stateChanged(void* self, int newstate);
void QScroller_scrollerPropertiesChanged(void* self, void* arg1);
void* QScroller_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QScroller_connect2(void* sender, void* signal, void* receiver, void* method);
char* QScroller_tr2(const char* s, const char* c, long long* __out_len);
char* QScroller_tr3(const char* s, const char* c, int n, long long* __out_len);
int QScroller_grabGesture2(void* target, int gestureType);
int QScroller_handleInput2(void* self, int input, void* position, long long timestamp);
void* QScroller_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QScroller_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QScroller_connect_stateChanged(void* self, long long id);
void QScroller_connect_scrollerPropertiesChanged(void* self, long long id);
void QScroller_connect_destroyed(void* self, long long id);
void QScroller_connect_destroyed2(void* self, long long id);
void QScroller_connect_objectNameChanged(void* self, long long id);
void QScroller__disconnect(long long id);
void QScroller_delete(void* self);

#ifdef __cplusplus
}
#endif
