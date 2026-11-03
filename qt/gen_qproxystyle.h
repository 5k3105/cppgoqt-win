#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QProxyStyle_base(void* self);
void QProxyStyle_chain(void* self, void** out);
void QProxyStyle__set_override(void* self, int idx);
void* QProxyStyle_fromQObject(void* o);
void* QProxyStyle_new(long long __goid);
void* QProxyStyle_new2(long long __goid, const char* key, long long key_len);
void* QProxyStyle_new3(long long __goid, void* style);
char* QProxyStyle_tr(const char* s, long long* __out_len);
void* QProxyStyle_baseStyle(void* self);
void QProxyStyle_setBaseStyle(void* self, void* style);
void* QProxyStyle_visualRect(int direction, void* boundingRect, void* logicalRect);
void* QProxyStyle_visualPos(int direction, void* boundingRect, void* logicalPos);
int QProxyStyle_sliderPositionFromValue(int min, int max, int val, int space);
int QProxyStyle_sliderValueFromPosition(int min, int max, int pos, int space);
int QProxyStyle_visualAlignment(int direction, int alignment);
void* QProxyStyle_alignedRect(int direction, int alignment, void* size, void* rectangle);
void* QProxyStyle_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QProxyStyle_connect2(void* sender, void* signal, void* receiver, void* method);
void* QProxyStyle_sender(void* self);
int QProxyStyle_senderSignalIndex(void* self);
int QProxyStyle_receivers(void* self, const char* signal);
int QProxyStyle_isSignalConnected(void* self, void* signal);
char* QProxyStyle_tr2(const char* s, const char* c, long long* __out_len);
char* QProxyStyle_tr3(const char* s, const char* c, int n, long long* __out_len);
int QProxyStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown);
int QProxyStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown);
void* QProxyStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QProxyStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QProxyStyle_metaObject_default(void* self);
void QProxyStyle_drawPrimitive_default(void* self, int element, void* option, void* painter, void* widget);
void QProxyStyle_drawControl_default(void* self, int element, void* option, void* painter, void* widget);
void QProxyStyle_drawComplexControl_default(void* self, int control, void* option, void* painter, void* widget);
void QProxyStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole);
void QProxyStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap);
void* QProxyStyle_sizeFromContents_default(void* self, int type_, void* option, void* size, void* widget);
void* QProxyStyle_subElementRect_default(void* self, int element, void* option, void* widget);
void* QProxyStyle_subControlRect_default(void* self, int cc, void* opt, int sc, void* widget);
void* QProxyStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len);
void* QProxyStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap);
int QProxyStyle_hitTestComplexControl_default(void* self, int control, void* option, void* pos, void* widget);
int QProxyStyle_styleHint_default(void* self, int hint, void* option, void* widget, void* returnData);
int QProxyStyle_pixelMetric_default(void* self, int metric, void* option, void* widget);
int QProxyStyle_layoutSpacing_default(void* self, int control1, int control2, int orientation, void* option, void* widget);
void* QProxyStyle_standardIcon_default(void* self, int standardIcon, void* option, void* widget);
void* QProxyStyle_standardPixmap_default(void* self, int standardPixmap, void* opt, void* widget);
void* QProxyStyle_generatedIconPixmap_default(void* self, int iconMode, void* pixmap, void* opt);
void* QProxyStyle_standardPalette_default(void* self);
void QProxyStyle_polish_default(void* self, void* widget);
void QProxyStyle_unpolish_default(void* self, void* widget);
int QProxyStyle_event_default(void* self, void* e);
int QProxyStyle_eventFilter_default(void* self, void* watched, void* event);
void QProxyStyle_timerEvent_default(void* self, void* event);
void QProxyStyle_childEvent_default(void* self, void* event);
void QProxyStyle_customEvent_default(void* self, void* event);
void QProxyStyle_connectNotify_default(void* self, void* signal);
void QProxyStyle_disconnectNotify_default(void* self, void* signal);
void QProxyStyle_connect_destroyed(void* self, long long id);
void QProxyStyle_connect_destroyed2(void* self, long long id);
void QProxyStyle_connect_objectNameChanged(void* self, long long id);
void QProxyStyle__disconnect(long long id);
void* QProxyStyle__Go_new(long long id);
void QProxyStyle_delete(void* self);

#ifdef __cplusplus
}
#endif
