#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSpacerItem_base(void* self);
void* QSpacerItem_new(int w, int h);
void* QSpacerItem_new2(void* arg1);
void* QSpacerItem_new3(int w, int h, int hData);
void* QSpacerItem_new4(int w, int h, int hData, int vData);
void QSpacerItem_changeSize(void* self, int w, int h);
void* QSpacerItem_sizePolicy(void* self);
void* QSpacerItem_OperatorAssign(void* self, void* arg1);
void QSpacerItem_changeSize2(void* self, int w, int h, int hData);
void QSpacerItem_changeSize3(void* self, int w, int h, int hData, int vData);
void QSpacerItem_delete(void* self);

#ifdef __cplusplus
}
#endif
