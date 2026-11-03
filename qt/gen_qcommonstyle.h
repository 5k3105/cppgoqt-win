#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCommonStyle_base(void* self);
void QCommonStyle_chain(void* self, void** out);
void QCommonStyle__set_override(void* self, int idx);
void* QCommonStyle_fromQObject(void* o);
void* QCommonStyle_new(long long __goid);
char* QCommonStyle_tr(const char* s, long long* __out_len);
void* QCommonStyle_visualRect(int direction, void* boundingRect, void* logicalRect);
void* QCommonStyle_visualPos(int direction, void* boundingRect, void* logicalPos);
int QCommonStyle_sliderPositionFromValue(int min, int max, int val, int space);
int QCommonStyle_sliderValueFromPosition(int min, int max, int pos, int space);
int QCommonStyle_visualAlignment(int direction, int alignment);
void* QCommonStyle_alignedRect(int direction, int alignment, void* size, void* rectangle);
void* QCommonStyle_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QCommonStyle_connect2(void* sender, void* signal, void* receiver, void* method);
void* QCommonStyle_sender(void* self);
int QCommonStyle_senderSignalIndex(void* self);
int QCommonStyle_receivers(void* self, const char* signal);
int QCommonStyle_isSignalConnected(void* self, void* signal);
char* QCommonStyle_tr2(const char* s, const char* c, long long* __out_len);
char* QCommonStyle_tr3(const char* s, const char* c, int n, long long* __out_len);
int QCommonStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown);
int QCommonStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown);
void* QCommonStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QCommonStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QCommonStyle_metaObject_default(void* self);
void QCommonStyle_drawPrimitive_default(void* self, int pe, void* opt, void* p, void* w);
void QCommonStyle_drawControl_default(void* self, int element, void* opt, void* p, void* w);
void* QCommonStyle_subElementRect_default(void* self, int r, void* opt, void* widget);
void QCommonStyle_drawComplexControl_default(void* self, int cc, void* opt, void* p, void* w);
int QCommonStyle_hitTestComplexControl_default(void* self, int cc, void* opt, void* pt, void* w);
void* QCommonStyle_subControlRect_default(void* self, int cc, void* opt, int sc, void* w);
void* QCommonStyle_sizeFromContents_default(void* self, int ct, void* opt, void* contentsSize, void* widget);
int QCommonStyle_pixelMetric_default(void* self, int m, void* opt, void* widget);
int QCommonStyle_styleHint_default(void* self, int sh, void* opt, void* w, void* shret);
void* QCommonStyle_standardIcon_default(void* self, int standardIcon, void* opt, void* widget);
void* QCommonStyle_standardPixmap_default(void* self, int sp, void* opt, void* widget);
void* QCommonStyle_generatedIconPixmap_default(void* self, int iconMode, void* pixmap, void* opt);
int QCommonStyle_layoutSpacing_default(void* self, int control1, int control2, int orientation, void* option, void* widget);
void QCommonStyle_polish_default(void* self, void* arg1);
void QCommonStyle_unpolish_default(void* self, void* widget);
void* QCommonStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len);
void* QCommonStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap);
void QCommonStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole);
void QCommonStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap);
void* QCommonStyle_standardPalette_default(void* self);
int QCommonStyle_event_default(void* self, void* event);
int QCommonStyle_eventFilter_default(void* self, void* watched, void* event);
void QCommonStyle_timerEvent_default(void* self, void* event);
void QCommonStyle_childEvent_default(void* self, void* event);
void QCommonStyle_customEvent_default(void* self, void* event);
void QCommonStyle_connectNotify_default(void* self, void* signal);
void QCommonStyle_disconnectNotify_default(void* self, void* signal);
void QCommonStyle_connect_destroyed(void* self, long long id);
void QCommonStyle_connect_destroyed2(void* self, long long id);
void QCommonStyle_connect_objectNameChanged(void* self, long long id);
void QCommonStyle__disconnect(long long id);
void* QCommonStyle__Go_new(long long id);
void QCommonStyle_delete(void* self);

#ifdef __cplusplus
}
#endif
