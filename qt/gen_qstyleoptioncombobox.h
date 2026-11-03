#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyleOptionComboBox_base(void* self);
void QStyleOptionComboBox_chain(void* self, void** out);
void* QStyleOptionComboBox_new();
void* QStyleOptionComboBox_new2(void* other);
void* QStyleOptionComboBox_OperatorAssign(void* self, void* arg1);
int QStyleOptionComboBox_editable_get(void* self);
void QStyleOptionComboBox_editable_set(void* self, int editable);
void* QStyleOptionComboBox_popupRect_get(void* self);
void QStyleOptionComboBox_popupRect_set(void* self, void* popupRect);
int QStyleOptionComboBox_frame_get(void* self);
void QStyleOptionComboBox_frame_set(void* self, int frame);
char* QStyleOptionComboBox_currentText_get(void* self, long long* __out_len);
void QStyleOptionComboBox_currentText_set(void* self, const char* currentText, long long currentText_len);
void* QStyleOptionComboBox_currentIcon_get(void* self);
void QStyleOptionComboBox_currentIcon_set(void* self, void* currentIcon);
void* QStyleOptionComboBox_iconSize_get(void* self);
void QStyleOptionComboBox_iconSize_set(void* self, void* iconSize);
int QStyleOptionComboBox_textAlignment_get(void* self);
void QStyleOptionComboBox_textAlignment_set(void* self, int textAlignment);
void QStyleOptionComboBox_delete(void* self);

#ifdef __cplusplus
}
#endif
