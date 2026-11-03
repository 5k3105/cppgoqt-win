#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QMetaMethod_new();
void* QMetaMethod_new2(void* arg1);
void* QMetaMethod_methodSignature(void* self);
void* QMetaMethod_name(void* self);
void* QMetaMethod_nameView(void* self);
const char* QMetaMethod_typeName(void* self);
int QMetaMethod_returnType(void* self);
void* QMetaMethod_returnMetaType(void* self);
int QMetaMethod_parameterCount(void* self);
int QMetaMethod_parameterType(void* self, int index);
void* QMetaMethod_parameterMetaType(void* self, int index);
void QMetaMethod_getParameterTypes(void* self, int* types);
void** QMetaMethod_parameterTypes(void* self, long long* __out_len);
void* QMetaMethod_parameterTypeName(void* self, int index);
void** QMetaMethod_parameterNames(void* self, long long* __out_len);
const char* QMetaMethod_tag(void* self);
int QMetaMethod_access(void* self);
int QMetaMethod_methodType(void* self);
int QMetaMethod_attributes(void* self);
int QMetaMethod_methodIndex(void* self);
int QMetaMethod_relativeMethodIndex(void* self);
int QMetaMethod_revision(void* self);
int QMetaMethod_isConst(void* self);
void* QMetaMethod_enclosingMetaObject(void* self);
int QMetaMethod_invoke(void* self, void* object, int connectionType, void* returnValue);
int QMetaMethod_invoke2(void* self, void* object, void* returnValue);
int QMetaMethod_invoke3(void* self, void* object, int connectionType, void* val0);
int QMetaMethod_invoke4(void* self, void* object, void* val0);
int QMetaMethod_isValid(void* self);
int QMetaMethod_invoke5(void* self, void* object, int connectionType, void* returnValue, void* val0);
int QMetaMethod_invoke6(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1);
int QMetaMethod_invoke7(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2);
int QMetaMethod_invoke8(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3);
int QMetaMethod_invoke9(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaMethod_invoke10(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaMethod_invoke11(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaMethod_invoke12(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaMethod_invoke13(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaMethod_invoke14(void* self, void* object, int connectionType, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_invoke15(void* self, void* object, void* returnValue, void* val0);
int QMetaMethod_invoke16(void* self, void* object, void* returnValue, void* val0, void* val1);
int QMetaMethod_invoke17(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2);
int QMetaMethod_invoke18(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3);
int QMetaMethod_invoke19(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaMethod_invoke20(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaMethod_invoke21(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaMethod_invoke22(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaMethod_invoke23(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaMethod_invoke24(void* self, void* object, void* returnValue, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_invoke25(void* self, void* object, int connectionType, void* val0, void* val1);
int QMetaMethod_invoke26(void* self, void* object, int connectionType, void* val0, void* val1, void* val2);
int QMetaMethod_invoke27(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3);
int QMetaMethod_invoke28(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaMethod_invoke29(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaMethod_invoke30(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaMethod_invoke31(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaMethod_invoke32(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaMethod_invoke33(void* self, void* object, int connectionType, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_invoke34(void* self, void* object, void* val0, void* val1);
int QMetaMethod_invoke35(void* self, void* object, void* val0, void* val1, void* val2);
int QMetaMethod_invoke36(void* self, void* object, void* val0, void* val1, void* val2, void* val3);
int QMetaMethod_invoke37(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4);
int QMetaMethod_invoke38(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5);
int QMetaMethod_invoke39(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6);
int QMetaMethod_invoke40(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7);
int QMetaMethod_invoke41(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8);
int QMetaMethod_invoke42(void* self, void* object, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9);
int QMetaMethod_OperatorEqual(void* self, void* rhs);
int QMetaMethod_OperatorNotEqual(void* self, void* rhs);
void QMetaMethod_delete(void* self);

#ifdef __cplusplus
}
#endif
