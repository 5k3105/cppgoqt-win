#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGridLayout_base(void* self);
void QGridLayout_chain(void* self, void** out);
void QGridLayout__set_override(void* self, int idx);
void* QGridLayout_fromQObject(void* o);
void* QGridLayout_new(long long __goid, void* parent);
void* QGridLayout_new2(long long __goid);
char* QGridLayout_tr(const char* s, long long* __out_len);
void QGridLayout_setHorizontalSpacing(void* self, int spacing);
int QGridLayout_horizontalSpacing(void* self);
void QGridLayout_setVerticalSpacing(void* self, int spacing);
int QGridLayout_verticalSpacing(void* self);
void QGridLayout_setRowStretch(void* self, int row, int stretch);
void QGridLayout_setColumnStretch(void* self, int column, int stretch);
int QGridLayout_rowStretch(void* self, int row);
int QGridLayout_columnStretch(void* self, int column);
void QGridLayout_setRowMinimumHeight(void* self, int row, int minSize);
void QGridLayout_setColumnMinimumWidth(void* self, int column, int minSize);
int QGridLayout_rowMinimumHeight(void* self, int row);
int QGridLayout_columnMinimumWidth(void* self, int column);
int QGridLayout_columnCount(void* self);
int QGridLayout_rowCount(void* self);
void* QGridLayout_cellRect(void* self, int row, int column);
void QGridLayout_addWidget(void* self, void* w);
void QGridLayout_addWidget2(void* self, void* arg1, int row, int column);
void QGridLayout_addWidget3(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_addLayout(void* self, void* arg1, int row, int column);
void QGridLayout_addLayout2(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_setOriginCorner(void* self, int arg1);
int QGridLayout_originCorner(void* self);
void* QGridLayout_itemAtPosition(void* self, int row, int column);
void QGridLayout_addItem(void* self, void* item, int row, int column);
void QGridLayout_setDefaultPositioning(void* self, int n, int orient);
void QGridLayout_getItemPosition(void* self, int idx, int* row, int* column, int* rowSpan, int* columnSpan);
void* QGridLayout_closestAcceptableSize(void* w, void* s);
void QGridLayout_widgetEvent(void* self, void* arg1);
void QGridLayout_addChildLayout(void* self, void* l);
void QGridLayout_addChildWidget(void* self, void* w);
int QGridLayout_adoptLayout(void* self, void* layout);
void* QGridLayout_alignmentRect(void* self, void* arg1);
void* QGridLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QGridLayout_connect2(void* sender, void* signal, void* receiver, void* method);
void* QGridLayout_sender(void* self);
int QGridLayout_senderSignalIndex(void* self);
int QGridLayout_receivers(void* self, const char* signal);
int QGridLayout_isSignalConnected(void* self, void* signal);
char* QGridLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QGridLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void QGridLayout_addWidget4(void* self, void* arg1, int row, int column, int arg4);
void QGridLayout_addWidget5(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan, int arg6);
void QGridLayout_addLayout3(void* self, void* arg1, int row, int column, int arg4);
void QGridLayout_addLayout4(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan, int arg6);
void QGridLayout_addItem2(void* self, void* item, int row, int column, int rowSpan);
void QGridLayout_addItem3(void* self, void* item, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_addItem4(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int arg6);
void* QGridLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QGridLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QGridLayout_addItem_default(void* self, void* item, int row, int column);
void QGridLayout_addItem_default2(void* self, void* item, int row, int column, int rowSpan);
void QGridLayout_addItem_default3(void* self, void* item, int row, int column, int rowSpan, int columnSpan);
void QGridLayout_addItem_default4(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int arg6);
void* QGridLayout_metaObject_default(void* self);
void* QGridLayout_sizeHint_default(void* self);
void* QGridLayout_minimumSize_default(void* self);
void* QGridLayout_maximumSize_default(void* self);
void QGridLayout_setSpacing_default(void* self, int spacing);
int QGridLayout_spacing_default(void* self);
int QGridLayout_hasHeightForWidth_default(void* self);
int QGridLayout_heightForWidth_default(void* self, int arg1);
int QGridLayout_minimumHeightForWidth_default(void* self, int arg1);
int QGridLayout_expandingDirections_default(void* self);
void QGridLayout_invalidate_default(void* self);
void* QGridLayout_itemAt_default(void* self, int index);
void* QGridLayout_takeAt_default(void* self, int index);
int QGridLayout_count_default(void* self);
void QGridLayout_setGeometry_default(void* self, void* arg1);
void* QGridLayout_geometry_default(void* self);
int QGridLayout_indexOf_default(void* self, void* arg1);
int QGridLayout_isEmpty_default(void* self);
int QGridLayout_controlTypes_default(void* self);
void* QGridLayout_replaceWidget_default(void* self, void* from, void* to, int options);
void* QGridLayout_layout_default(void* self);
void QGridLayout_childEvent_default(void* self, void* e);
int QGridLayout_event_default(void* self, void* event);
int QGridLayout_eventFilter_default(void* self, void* watched, void* event);
void QGridLayout_timerEvent_default(void* self, void* event);
void QGridLayout_customEvent_default(void* self, void* event);
void QGridLayout_connectNotify_default(void* self, void* signal);
void QGridLayout_disconnectNotify_default(void* self, void* signal);
void* QGridLayout_widget_default(void* self);
void* QGridLayout_spacerItem_default(void* self);
void QGridLayout_connect_destroyed(void* self, long long id);
void QGridLayout_connect_destroyed2(void* self, long long id);
void QGridLayout_connect_objectNameChanged(void* self, long long id);
void QGridLayout__disconnect(long long id);
void* QGridLayout__Go_new(long long id);
void QGridLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
