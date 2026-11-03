#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QTreeWidgetItem__set_override(void* self, int idx);
void* QTreeWidgetItem_new(long long __goid);
void* QTreeWidgetItem_new2(long long __goid, char** strings, long long strings_len);
void* QTreeWidgetItem_new3(long long __goid, void* treeview);
void* QTreeWidgetItem_new4(long long __goid, void* treeview, char** strings, long long strings_len);
void* QTreeWidgetItem_new5(long long __goid, void* treeview, void* after);
void* QTreeWidgetItem_new6(long long __goid, void* parent);
void* QTreeWidgetItem_new7(long long __goid, void* parent, char** strings, long long strings_len);
void* QTreeWidgetItem_new8(long long __goid, void* parent, void* after);
void* QTreeWidgetItem_new9(long long __goid, void* other);
void* QTreeWidgetItem_new10(long long __goid, int type_);
void* QTreeWidgetItem_new11(long long __goid, char** strings, long long strings_len, int type_);
void* QTreeWidgetItem_new12(long long __goid, void* treeview, int type_);
void* QTreeWidgetItem_new13(long long __goid, void* treeview, char** strings, long long strings_len, int type_);
void* QTreeWidgetItem_new14(long long __goid, void* treeview, void* after, int type_);
void* QTreeWidgetItem_new15(long long __goid, void* parent, int type_);
void* QTreeWidgetItem_new16(long long __goid, void* parent, char** strings, long long strings_len, int type_);
void* QTreeWidgetItem_new17(long long __goid, void* parent, void* after, int type_);
void* QTreeWidgetItem_clone(void* self);
void* QTreeWidgetItem_treeWidget(void* self);
void QTreeWidgetItem_setSelected(void* self, int select_);
int QTreeWidgetItem_isSelected(void* self);
void QTreeWidgetItem_setHidden(void* self, int hide);
int QTreeWidgetItem_isHidden(void* self);
void QTreeWidgetItem_setExpanded(void* self, int expand);
int QTreeWidgetItem_isExpanded(void* self);
void QTreeWidgetItem_setFirstColumnSpanned(void* self, int span);
int QTreeWidgetItem_isFirstColumnSpanned(void* self);
void QTreeWidgetItem_setDisabled(void* self, int disabled);
int QTreeWidgetItem_isDisabled(void* self);
void QTreeWidgetItem_setChildIndicatorPolicy(void* self, int policy);
int QTreeWidgetItem_childIndicatorPolicy(void* self);
int QTreeWidgetItem_flags(void* self);
void QTreeWidgetItem_setFlags(void* self, int flags);
char* QTreeWidgetItem_text(void* self, int column, long long* __out_len);
void QTreeWidgetItem_setText(void* self, int column, const char* text, long long text_len);
void* QTreeWidgetItem_icon(void* self, int column);
void QTreeWidgetItem_setIcon(void* self, int column, void* icon);
char* QTreeWidgetItem_statusTip(void* self, int column, long long* __out_len);
void QTreeWidgetItem_setStatusTip(void* self, int column, const char* statusTip, long long statusTip_len);
char* QTreeWidgetItem_toolTip(void* self, int column, long long* __out_len);
void QTreeWidgetItem_setToolTip(void* self, int column, const char* toolTip, long long toolTip_len);
char* QTreeWidgetItem_whatsThis(void* self, int column, long long* __out_len);
void QTreeWidgetItem_setWhatsThis(void* self, int column, const char* whatsThis, long long whatsThis_len);
void* QTreeWidgetItem_font(void* self, int column);
void QTreeWidgetItem_setFont(void* self, int column, void* font);
int QTreeWidgetItem_textAlignment(void* self, int column);
void QTreeWidgetItem_setTextAlignment(void* self, int column, int alignment);
void QTreeWidgetItem_setTextAlignment2(void* self, int column, int alignment);
void QTreeWidgetItem_setTextAlignment3(void* self, int column, int alignment);
void* QTreeWidgetItem_background(void* self, int column);
void QTreeWidgetItem_setBackground(void* self, int column, void* brush);
void* QTreeWidgetItem_foreground(void* self, int column);
void QTreeWidgetItem_setForeground(void* self, int column, void* brush);
int QTreeWidgetItem_checkState(void* self, int column);
void QTreeWidgetItem_setCheckState(void* self, int column, int state);
void* QTreeWidgetItem_sizeHint(void* self, int column);
void QTreeWidgetItem_setSizeHint(void* self, int column, void* size);
void* QTreeWidgetItem_data(void* self, int column, int role);
void QTreeWidgetItem_setData(void* self, int column, int role, void* value);
int QTreeWidgetItem_OperatorLesser(void* self, void* other);
void QTreeWidgetItem_read(void* self, void* in);
void QTreeWidgetItem_write(void* self, void* out);
void* QTreeWidgetItem_OperatorAssign(void* self, void* other);
void* QTreeWidgetItem_parent(void* self);
void* QTreeWidgetItem_child(void* self, int index);
int QTreeWidgetItem_childCount(void* self);
int QTreeWidgetItem_columnCount(void* self);
int QTreeWidgetItem_indexOfChild(void* self, void* child);
void QTreeWidgetItem_addChild(void* self, void* child);
void QTreeWidgetItem_insertChild(void* self, int index, void* child);
void QTreeWidgetItem_removeChild(void* self, void* child);
void* QTreeWidgetItem_takeChild(void* self, int index);
void QTreeWidgetItem_addChildren(void* self, void** children, long long children_len);
void QTreeWidgetItem_insertChildren(void* self, int index, void** children, long long children_len);
void** QTreeWidgetItem_takeChildren(void* self, long long* __out_len);
int QTreeWidgetItem_type(void* self);
void QTreeWidgetItem_sortChildren(void* self, int column, int order);
void QTreeWidgetItem_emitDataChanged(void* self);
void* QTreeWidgetItem_clone_default(void* self);
void* QTreeWidgetItem_data_default(void* self, int column, int role);
void QTreeWidgetItem_setData_default(void* self, int column, int role, void* value);
void QTreeWidgetItem_read_default(void* self, void* in);
void QTreeWidgetItem_write_default(void* self, void* out);
void* QTreeWidgetItem__Go_new(long long id);
void QTreeWidgetItem_delete(void* self);

#ifdef __cplusplus
}
#endif
