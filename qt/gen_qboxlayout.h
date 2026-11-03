#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QBoxLayout_base(void* self);
void QBoxLayout_chain(void* self, void** out);
void* QBoxLayout_fromQObject(void* o);
void* QBoxLayout_new(int arg1);
void* QBoxLayout_new2(int arg1, void* parent);
char* QBoxLayout_tr(const char* s, long long* __out_len);
int QBoxLayout_direction(void* self);
void QBoxLayout_setDirection(void* self, int arg1);
void QBoxLayout_addSpacing(void* self, int size);
void QBoxLayout_addStretch(void* self);
void QBoxLayout_addSpacerItem(void* self, void* spacerItem);
void QBoxLayout_addWidget(void* self, void* arg1);
void QBoxLayout_addLayout(void* self, void* layout);
void QBoxLayout_addStrut(void* self, int arg1);
void QBoxLayout_insertSpacing(void* self, int index, int size);
void QBoxLayout_insertStretch(void* self, int index);
void QBoxLayout_insertSpacerItem(void* self, int index, void* spacerItem);
void QBoxLayout_insertWidget(void* self, int index, void* widget);
void QBoxLayout_insertLayout(void* self, int index, void* layout);
void QBoxLayout_insertItem(void* self, int index, void* arg2);
int QBoxLayout_setStretchFactor(void* self, void* w, int stretch);
int QBoxLayout_setStretchFactor2(void* self, void* l, int stretch);
void QBoxLayout_setStretch(void* self, int index, int stretch);
int QBoxLayout_stretch(void* self, int index);
void* QBoxLayout_closestAcceptableSize(void* w, void* s);
void* QBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method);
char* QBoxLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void QBoxLayout_addStretch2(void* self, int stretch);
void QBoxLayout_addWidget2(void* self, void* arg1, int stretch);
void QBoxLayout_addWidget3(void* self, void* arg1, int stretch, int alignment);
void QBoxLayout_addLayout2(void* self, void* layout, int stretch);
void QBoxLayout_insertStretch2(void* self, int index, int stretch);
void QBoxLayout_insertWidget2(void* self, int index, void* widget, int stretch);
void QBoxLayout_insertWidget3(void* self, int index, void* widget, int stretch, int alignment);
void QBoxLayout_insertLayout2(void* self, int index, void* layout, int stretch);
void* QBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QBoxLayout_connect_destroyed(void* self, long long id);
void QBoxLayout_connect_destroyed2(void* self, long long id);
void QBoxLayout_connect_objectNameChanged(void* self, long long id);
void QBoxLayout__disconnect(long long id);
void QBoxLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
