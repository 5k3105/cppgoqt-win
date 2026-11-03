#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QListWidgetItem__set_override(void* self, int idx);
void* QListWidgetItem_new(long long __goid);
void* QListWidgetItem_new2(long long __goid, const char* text, long long text_len);
void* QListWidgetItem_new3(long long __goid, void* icon, const char* text, long long text_len);
void* QListWidgetItem_new4(long long __goid, void* other);
void* QListWidgetItem_new5(long long __goid, void* listview);
void* QListWidgetItem_new6(long long __goid, void* listview, int type_);
void* QListWidgetItem_new7(long long __goid, const char* text, long long text_len, void* listview);
void* QListWidgetItem_new8(long long __goid, const char* text, long long text_len, void* listview, int type_);
void* QListWidgetItem_new9(long long __goid, void* icon, const char* text, long long text_len, void* listview);
void* QListWidgetItem_new10(long long __goid, void* icon, const char* text, long long text_len, void* listview, int type_);
void* QListWidgetItem_clone(void* self);
void* QListWidgetItem_listWidget(void* self);
void QListWidgetItem_setSelected(void* self, int select_);
int QListWidgetItem_isSelected(void* self);
void QListWidgetItem_setHidden(void* self, int hide);
int QListWidgetItem_isHidden(void* self);
int QListWidgetItem_flags(void* self);
void QListWidgetItem_setFlags(void* self, int flags);
char* QListWidgetItem_text(void* self, long long* __out_len);
void QListWidgetItem_setText(void* self, const char* text, long long text_len);
void* QListWidgetItem_icon(void* self);
void QListWidgetItem_setIcon(void* self, void* icon);
char* QListWidgetItem_statusTip(void* self, long long* __out_len);
void QListWidgetItem_setStatusTip(void* self, const char* statusTip, long long statusTip_len);
char* QListWidgetItem_toolTip(void* self, long long* __out_len);
void QListWidgetItem_setToolTip(void* self, const char* toolTip, long long toolTip_len);
char* QListWidgetItem_whatsThis(void* self, long long* __out_len);
void QListWidgetItem_setWhatsThis(void* self, const char* whatsThis, long long whatsThis_len);
void* QListWidgetItem_font(void* self);
void QListWidgetItem_setFont(void* self, void* font);
int QListWidgetItem_textAlignment(void* self);
void QListWidgetItem_setTextAlignment(void* self, int alignment);
void QListWidgetItem_setTextAlignment2(void* self, int alignment);
void QListWidgetItem_setTextAlignment3(void* self, int alignment);
void* QListWidgetItem_background(void* self);
void QListWidgetItem_setBackground(void* self, void* brush);
void* QListWidgetItem_foreground(void* self);
void QListWidgetItem_setForeground(void* self, void* brush);
int QListWidgetItem_checkState(void* self);
void QListWidgetItem_setCheckState(void* self, int state);
void* QListWidgetItem_sizeHint(void* self);
void QListWidgetItem_setSizeHint(void* self, void* size);
void* QListWidgetItem_data(void* self, int role);
void QListWidgetItem_setData(void* self, int role, void* value);
int QListWidgetItem_OperatorLesser(void* self, void* other);
void QListWidgetItem_read(void* self, void* in);
void QListWidgetItem_write(void* self, void* out);
void* QListWidgetItem_OperatorAssign(void* self, void* other);
int QListWidgetItem_type(void* self);
void* QListWidgetItem_clone_default(void* self);
void* QListWidgetItem_data_default(void* self, int role);
void QListWidgetItem_setData_default(void* self, int role, void* value);
void QListWidgetItem_read_default(void* self, void* in);
void QListWidgetItem_write_default(void* self, void* out);
void* QListWidgetItem__Go_new(long long id);
void QListWidgetItem_delete(void* self);

#ifdef __cplusplus
}
#endif
