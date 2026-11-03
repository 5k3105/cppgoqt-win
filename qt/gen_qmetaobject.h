#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaObject_new();
void* QMetaObject_new2(void* arg1);
const char* QMetaObject_className(void* self);
void* QMetaObject_superClass(void* self);
int QMetaObject_inherits(void* self, void* metaObject);
void* QMetaObject_cast(void* self, void* obj);
void* QMetaObject_cast2(void* self, void* obj);
char* QMetaObject_tr(void* self, const char* s, const char* c, long long* __out_len);
void* QMetaObject_metaType(void* self);
int QMetaObject_methodOffset(void* self);
int QMetaObject_enumeratorOffset(void* self);
int QMetaObject_propertyOffset(void* self);
int QMetaObject_classInfoOffset(void* self);
int QMetaObject_constructorCount(void* self);
int QMetaObject_methodCount(void* self);
int QMetaObject_enumeratorCount(void* self);
int QMetaObject_propertyCount(void* self);
int QMetaObject_classInfoCount(void* self);
int QMetaObject_indexOfConstructor(void* self, const char* constructor);
int QMetaObject_indexOfMethod(void* self, const char* method);
int QMetaObject_indexOfSignal(void* self, const char* signal);
int QMetaObject_indexOfSlot(void* self, const char* slot);
int QMetaObject_indexOfEnumerator(void* self, const char* name);
int QMetaObject_indexOfProperty(void* self, const char* name);
int QMetaObject_indexOfClassInfo(void* self, const char* name);
void* QMetaObject_constructor(void* self, int index);
void* QMetaObject_method(void* self, int index);
void* QMetaObject_enumerator(void* self, int index);
void* QMetaObject_property(void* self, int index);
void* QMetaObject_classInfo(void* self, int index);
void* QMetaObject_userProperty(void* self);
int QMetaObject_checkConnectArgs(const char* signal, const char* method);
int QMetaObject_checkConnectArgs2(void* signal, void* method);
void* QMetaObject_normalizedSignature(const char* method);
void* QMetaObject_normalizedType(const char* type_);
void* QMetaObject_connect(void* sender, int signal_index, void* receiver, int method_index);
int QMetaObject_disconnectOne(void* sender, int signal_index, void* receiver, int method_index);
void QMetaObject_connectSlotsByName(void* o);
int QMetaObject_invokeMethod(void* obj, const char* member, int arg3, void* ret);
int QMetaObject_invokeMethod2(void* obj, const char* member, void* ret);
int QMetaObject_invokeMethod3(void* obj, const char* member, int type_, void* val0);
int QMetaObject_invokeMethod4(void* obj, const char* member, void* val0);
void* QMetaObject_newInstance(void* self, void* val0);
void* QMetaObject_OperatorAssign(void* self, void* arg1);
char* QMetaObject_tr2(void* self, const char* s, const char* c, int n, long long* __out_len);
void* QMetaObject_connect2(void* sender, int signal_index, void* receiver, int method_index, int type_);
void* QMetaObject_connect3(void* sender, int signal_index, void* receiver, int method_index, int type_, int* types);
int QMetaObject_invokeMethod5(void* obj, const char* member, int arg3, void* ret, void* val0);
int QMetaObject_invokeMethod6(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1);
int QMetaObject_invokeMethod7(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2);
int QMetaObject_invokeMethod8(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3);
int QMetaObject_invokeMethod9(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaObject_invokeMethod10(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaObject_invokeMethod11(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaObject_invokeMethod12(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaObject_invokeMethod13(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaObject_invokeMethod14(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_invokeMethod15(void* obj, const char* member, void* ret, void* val0);
int QMetaObject_invokeMethod16(void* obj, const char* member, void* ret, void* val0, void* val1);
int QMetaObject_invokeMethod17(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2);
int QMetaObject_invokeMethod18(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3);
int QMetaObject_invokeMethod19(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaObject_invokeMethod20(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaObject_invokeMethod21(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaObject_invokeMethod22(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaObject_invokeMethod23(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaObject_invokeMethod24(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_invokeMethod25(void* obj, const char* member, int type_, void* val0, void* val1);
int QMetaObject_invokeMethod26(void* obj, const char* member, int type_, void* val0, void* val1, void* val2);
int QMetaObject_invokeMethod27(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3);
int QMetaObject_invokeMethod28(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaObject_invokeMethod29(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaObject_invokeMethod30(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaObject_invokeMethod31(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaObject_invokeMethod32(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaObject_invokeMethod33(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaObject_invokeMethod34(void* obj, const char* member, void* val0, void* val1);
int QMetaObject_invokeMethod35(void* obj, const char* member, void* val0, void* val1, void* val2);
int QMetaObject_invokeMethod36(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3);
int QMetaObject_invokeMethod37(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaObject_invokeMethod38(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaObject_invokeMethod39(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaObject_invokeMethod40(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaObject_invokeMethod41(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaObject_invokeMethod42(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
void* QMetaObject_newInstance2(void* self, void* val0, void* val1);
void* QMetaObject_newInstance3(void* self, void* val0, void* val1, void* val2);
void* QMetaObject_newInstance4(void* self, void* val0, void* val1, void* val2, void* val3);
void* QMetaObject_newInstance5(void* self, void* val0, void* val1, void* val2, void* val3, void* val4);
void* QMetaObject_newInstance6(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
void* QMetaObject_newInstance7(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
void* QMetaObject_newInstance8(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
void* QMetaObject_newInstance9(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
void* QMetaObject_newInstance10(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
void* QMetaObject_d_get(void* self);
void QMetaObject_d_set(void* self, void* d);
void QMetaObject_delete(void* self);

#ifdef __cplusplus
}
#endif
