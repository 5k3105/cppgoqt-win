#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QTableWidgetItem__set_override(void* self, int idx);
void* QTableWidgetItem_new(long long __goid);
void* QTableWidgetItem_new2(long long __goid, const char* text, long long text_len);
void* QTableWidgetItem_new3(long long __goid, void* icon, const char* text, long long text_len);
void* QTableWidgetItem_new4(long long __goid, void* other);
void* QTableWidgetItem_new5(long long __goid, int type_);
void* QTableWidgetItem_new6(long long __goid, const char* text, long long text_len, int type_);
void* QTableWidgetItem_new7(long long __goid, void* icon, const char* text, long long text_len, int type_);
void* QTableWidgetItem_clone(void* self);
void* QTableWidgetItem_tableWidget(void* self);
int QTableWidgetItem_row(void* self);
int QTableWidgetItem_column(void* self);
void QTableWidgetItem_setSelected(void* self, int select_);
int QTableWidgetItem_isSelected(void* self);
int QTableWidgetItem_flags(void* self);
void QTableWidgetItem_setFlags(void* self, int flags);
char* QTableWidgetItem_text(void* self, long long* __out_len);
void QTableWidgetItem_setText(void* self, const char* text, long long text_len);
void* QTableWidgetItem_icon(void* self);
void QTableWidgetItem_setIcon(void* self, void* icon);
char* QTableWidgetItem_statusTip(void* self, long long* __out_len);
void QTableWidgetItem_setStatusTip(void* self, const char* statusTip, long long statusTip_len);
char* QTableWidgetItem_toolTip(void* self, long long* __out_len);
void QTableWidgetItem_setToolTip(void* self, const char* toolTip, long long toolTip_len);
char* QTableWidgetItem_whatsThis(void* self, long long* __out_len);
void QTableWidgetItem_setWhatsThis(void* self, const char* whatsThis, long long whatsThis_len);
void* QTableWidgetItem_font(void* self);
void QTableWidgetItem_setFont(void* self, void* font);
int QTableWidgetItem_textAlignment(void* self);
void QTableWidgetItem_setTextAlignment(void* self, int alignment);
void QTableWidgetItem_setTextAlignment2(void* self, int alignment);
void QTableWidgetItem_setTextAlignment3(void* self, int alignment);
void* QTableWidgetItem_background(void* self);
void QTableWidgetItem_setBackground(void* self, void* brush);
void* QTableWidgetItem_foreground(void* self);
void QTableWidgetItem_setForeground(void* self, void* brush);
int QTableWidgetItem_checkState(void* self);
void QTableWidgetItem_setCheckState(void* self, int state);
void* QTableWidgetItem_sizeHint(void* self);
void QTableWidgetItem_setSizeHint(void* self, void* size);
void* QTableWidgetItem_data(void* self, int role);
void QTableWidgetItem_setData(void* self, int role, void* value);
int QTableWidgetItem_OperatorLesser(void* self, void* other);
void QTableWidgetItem_read(void* self, void* in);
void QTableWidgetItem_write(void* self, void* out);
void* QTableWidgetItem_OperatorAssign(void* self, void* other);
int QTableWidgetItem_type(void* self);
void* QTableWidgetItem_clone_default(void* self);
void* QTableWidgetItem_data_default(void* self, int role);
void QTableWidgetItem_setData_default(void* self, int role, void* value);
void QTableWidgetItem_read_default(void* self, void* in);
void QTableWidgetItem_write_default(void* self, void* out);
void* QTableWidgetItem__Go_new(long long id);
void QTableWidgetItem_delete(void* self);

#ifdef __cplusplus
}
#endif
