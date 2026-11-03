#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyle_base(void* self);
void QStyle__set_override(void* self, int idx);
void* QStyle_fromQObject(void* o);
char* QStyle_tr(const char* s, long long* __out_len);
char* QStyle_name(void* self, long long* __out_len);
void QStyle_polish(void* self, void* widget);
void QStyle_unpolish(void* self, void* widget);
void QStyle_polish2(void* self, void* application);
void QStyle_unpolish2(void* self, void* application);
void QStyle_polish3(void* self, void* palette);
void* QStyle_itemTextRect(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len);
void* QStyle_itemPixmapRect(void* self, void* r, int flags, void* pixmap);
void QStyle_drawItemText(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole);
void QStyle_drawItemPixmap(void* self, void* painter, void* rect, int alignment, void* pixmap);
void* QStyle_standardPalette(void* self);
void QStyle_drawPrimitive(void* self, int pe, void* opt, void* p, void* w);
void QStyle_drawControl(void* self, int element, void* opt, void* p, void* w);
void* QStyle_subElementRect(void* self, int subElement, void* option, void* widget);
void QStyle_drawComplexControl(void* self, int cc, void* opt, void* p, void* widget);
int QStyle_hitTestComplexControl(void* self, int cc, void* opt, void* pt, void* widget);
void* QStyle_subControlRect(void* self, int cc, void* opt, int sc, void* widget);
int QStyle_pixelMetric(void* self, int metric, void* option, void* widget);
void* QStyle_sizeFromContents(void* self, int ct, void* opt, void* contentsSize, void* w);
int QStyle_styleHint(void* self, int stylehint, void* opt, void* widget, void* returnData);
void* QStyle_standardPixmap(void* self, int standardPixmap, void* opt, void* widget);
void* QStyle_standardIcon(void* self, int standardIcon, void* option, void* widget);
void* QStyle_generatedIconPixmap(void* self, int iconMode, void* pixmap, void* opt);
void* QStyle_visualRect(int direction, void* boundingRect, void* logicalRect);
void* QStyle_visualPos(int direction, void* boundingRect, void* logicalPos);
int QStyle_sliderPositionFromValue(int min, int max, int val, int space);
int QStyle_sliderValueFromPosition(int min, int max, int pos, int space);
int QStyle_visualAlignment(int direction, int alignment);
void* QStyle_alignedRect(int direction, int alignment, void* size, void* rectangle);
int QStyle_layoutSpacing(void* self, int control1, int control2, int orientation, void* option, void* widget);
int QStyle_combinedLayoutSpacing(void* self, int controls1, int controls2, int orientation);
void* QStyle_proxy(void* self);
void* QStyle_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QStyle_connect2(void* sender, void* signal, void* receiver, void* method);
void* QStyle_sender(void* self);
int QStyle_senderSignalIndex(void* self);
int QStyle_receivers(void* self, const char* signal);
int QStyle_isSignalConnected(void* self, void* signal);
char* QStyle_tr2(const char* s, const char* c, long long* __out_len);
char* QStyle_tr3(const char* s, const char* c, int n, long long* __out_len);
int QStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown);
int QStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown);
int QStyle_combinedLayoutSpacing2(void* self, int controls1, int controls2, int orientation, void* option);
int QStyle_combinedLayoutSpacing3(void* self, int controls1, int controls2, int orientation, void* option, void* widget);
void* QStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QStyle_polish_default(void* self, void* widget);
void QStyle_unpolish_default(void* self, void* widget);
void QStyle_polish_default2(void* self, void* application);
void QStyle_unpolish_default2(void* self, void* application);
void QStyle_polish_default3(void* self, void* palette);
void* QStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len);
void* QStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap);
void QStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole);
void QStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap);
void* QStyle_standardPalette_default(void* self);
void* QStyle_metaObject_default(void* self);
int QStyle_event_default(void* self, void* event);
int QStyle_eventFilter_default(void* self, void* watched, void* event);
void QStyle_timerEvent_default(void* self, void* event);
void QStyle_childEvent_default(void* self, void* event);
void QStyle_customEvent_default(void* self, void* event);
void QStyle_connectNotify_default(void* self, void* signal);
void QStyle_disconnectNotify_default(void* self, void* signal);
void QStyle_connect_destroyed(void* self, long long id);
void QStyle_connect_destroyed2(void* self, long long id);
void QStyle_connect_objectNameChanged(void* self, long long id);
void QStyle__disconnect(long long id);
void* QStyle__Go_new(long long id);
void QStyle_delete(void* self);

#ifdef __cplusplus
}
#endif
